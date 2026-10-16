import assert from 'node:assert/strict';
import { createHash } from 'node:crypto';
import { spawnSync } from 'node:child_process';
import { cpSync, existsSync, mkdirSync, mkdtempSync, readFileSync, readdirSync, renameSync, rmSync, statSync, symlinkSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { dirname, join } from 'node:path';
import { fileURLToPath } from 'node:url';
import test from 'node:test';
import { corePackages, prepareFixture } from './test-workspace.mjs';
import { findTypedCompiler } from './native-workspace.mjs';

const root = fileURLToPath(new URL('../', import.meta.url));
const hash = data => createHash('sha256').update(data).digest('hex');
const put = (root, name, content) => {
    mkdirSync(dirname(join(root, name)), { recursive: true });
    writeFileSync(join(root, name), content);
};
function inventory(root) {
    const files = {};
    function visit(directory, prefix) {
        for (const entry of readdirSync(directory, { withFileTypes: true })) {
            if (entry.name.startsWith('.')) continue;
            const name = prefix + entry.name, path = join(directory, entry.name);
            if (entry.isDirectory()) visit(path, name + '/');
            else if (name.endsWith('.go') || name === 'go.mod' || name === 'go.sum') files[name] = hash(readFileSync(path));
        }
    }
    visit(root, '');
    return files;
}
const dates = directory => Object.fromEntries(Object.keys(inventory(directory)).map(name =>
    [name, String(statSync(join(directory, name), { bigint: true }).mtimeNs)]));

test('the CLI publishes the typed package layout with frozen native dependencies and selected entries', async t => {
    const workspace = mkdtempSync(join(process.env.GOPURS_TEST_EVIDENCE ?? tmpdir(), 'driver-packages-'));
    let passed = false;
    t.after(() => {
        if (passed && !process.env.GOPURS_TEST_EVIDENCE) rmSync(workspace, { recursive: true, force: true });
        else console.log('Driver packages evidence: ' + workspace);
    });
    mkdirSync(join(workspace, 'typed-bin'));
    symlinkSync(findTypedCompiler(root), join(workspace, 'typed-bin/purs'));
    const fixture = prepareFixture(root, workspace, join(root, 'tests/passing/FFIIntegerReturns.purs'), 0, corePackages(root));
    const directory = fixture.directory, output = join(directory, 'output');
    const env = { ...process.env, GOWORK: 'off', GOMAXPROCS: '2', GOMEMLIMIT: '12GiB',
        GOPURS_RUST: '0', GOPURS_GO_PACKAGES: '0', PATH: join(workspace, 'typed-bin') + ':' + join(root, 'node_modules/.bin') + ':' + process.env.PATH };
    for (const key of ['GOPURS_ALLOC_PROFILE', 'GOPURS_FORCE_BUILD', 'GOPURS_COMPILER_SNAPSHOT', 'PPROF']) delete env[key];
    let sequence = 0;
    const commands = [];
    function run(label, command, args, cwd, overrides = {}, statuses = [0]) {
        const result = spawnSync(command, args, { cwd, env: { ...env, ...overrides }, encoding: 'utf8',
            timeout: 180_000, maxBuffer: 16 * 1024 * 1024 });
        put(workspace, `${sequence++}-${label}.log`, (result.stdout ?? '') + (result.stderr ?? ''));
        commands.push({ label, command, args, cwd, overrides, status: result.status });
        put(workspace, 'commands.json', JSON.stringify(commands, null, 2));
        assert.ifError(result.error);
        assert(statuses.includes(result.status), result.stdout + result.stderr);
        return result;
    }
    run('frontend', 'spago', ['build', '-q', '--offline'], directory);
    renameSync(output, join(workspace, 'tast'));
    const ffi = readFileSync(join(directory, 'src/Main.go'), 'utf8');
    const nativeFfi = 'import "gopurs/output/native"\n' + ffi.replace('return value', 'return value + native.Offset()');
    put(directory, 'src/Main.go', nativeFfi);
    const native = 'package native\nimport "example/offset"\nfunc Offset() int64 {return int64(offset.Value())}\n';
    const original = 'package offset\nfunc Value() int64 {return 0}\n';
    put(directory, 'local-module/go.mod', 'module example/offset\n\ngo 1.22\n');
    const goMod = 'module gopurs/output\n\ngo 1.22\nrequire example/offset v0.0.0\nreplace example/offset => ../local-module\n';
    const expectedByStep = new Map();
    let expectedDiagnostic;
    const modes = (process.env.GOPURS_TEST_MODES ?? 'js').split(',');
    for (const mode of modes) await t.test(mode, async () => {
        if (existsSync(output)) renameSync(output, join(directory, 'output-before-' + mode));
        cpSync(join(workspace, 'tast'), output, { recursive: true });
        put(directory, 'src/Main.go', nativeFfi);
        put(directory, 'local-module/offset.go', original);
        put(output, 'native/native.go', native);
        put(output, 'go.mod', goMod);
        put(output, 'purescript/external.txt', 'external asset\n');
        const flags = { GOPURS_JS: mode === 'js' ? '1' : '0', GOPURS_RUST: mode === 'rust' ? '1' : '0',
            GOPURS_PREPARE_JOBS: mode === 'native-parallel' ? '8' : '1', GOPURS_PBO_JOBS: '1',
            GOPURS_EMIT_JOBS: mode === 'native-parallel' ? '8' : '1', GOPURS_PIPELINE: mode === 'native-parallel' ? '1' : '0' };
        const compiler = process.env.GOPURS_TEST_COMPILER ?? join(root, 'bin/gopurs');
        const invoke = (label, packages, extra = {}, statuses = [0]) => run(mode + '-' + label, compiler, ['--main', 'Main'], directory,
            { ...flags, GOPURS_GO_PACKAGES: packages ? '1' : '0', ...extra }, statuses);
        invoke('monolithic', false);
        const before = run(mode + '-monolithic-execute', 'go', ['run', './main/main.go'], output).stdout;
        assert.equal(before.trim(), 'Done');
        const lock = statSync(join(output, '.gopurs-cache/write.lock'), { bigint: true }).ino;
        invoke('packages', true);
        const compare = step => {
            const files = inventory(output);
            if (expectedByStep.has(step)) assert.deepEqual(files, expectedByStep.get(step), step + ': host parity');
            else expectedByStep.set(step, files);
            assert.equal(readFileSync(join(output, 'go.mod'), 'utf8'), goMod);
            assert.equal(readFileSync(join(output, 'native/native.go'), 'utf8'), native);
            assert.equal(readFileSync(join(output, 'purescript/external.txt'), 'utf8'), 'external asset\n');
            assert.equal(statSync(join(output, '.gopurs-cache/write.lock'), { bigint: true }).ino, lock);
            return files;
        };
        assert.ok(!existsSync(join(output, 'purescript/Main.go')));
        assert.ok(readdirSync(join(output, 'purescript')).some(name => name.startsWith('p_')));
        const files = compare('packages');
        run(mode + '-dag', 'go', ['list', '-deps', '-json', './main'], output);
        const binary = join(directory, 'regression-' + mode);
        run(mode + '-build', 'go', ['build', '-p', '1', '-x', '-o', binary, './main/main.go'], output);
        assert.equal(run(mode + '-execute', binary, [], output).stdout, before);
        assert.equal(run(mode + '-selected-entry', 'go', ['run', './Main/main/main.go'], output).stdout, before);
        const timestamps = dates(output);
        assert.match(invoke('identical', true).stderr, /cache hit:/);
        assert.deepEqual(dates(output), timestamps);
        const build = run(mode + '-identical-build', 'go', ['build', '-p', '1', '-x', '-o', binary, './main/main.go'], output);
        assert.doesNotMatch(build.stderr, /\/compile\b|\/link\b/);
        invoke('forced', true, { GOPURS_FORCE_BUILD: '1' });
        assert.deepEqual(compare('forced'), files);
        assert.deepEqual(dates(output), timestamps);
        put(directory, 'local-module/offset.go', 'package offset\nfunc Value() int {return 0}\n');
        assert.match(invoke('native-signature', true).stderr, /cache miss:/);
        compare('native-signature');
        assert.equal(run(mode + '-native-signature-execute', 'go', ['run', './main/main.go'], output).stdout, before);
        assert.match(invoke('native-signature-identical', true).stderr, /cache hit:/);
        const pointer = readFileSync(join(output, '.gopurs-cache/current.json'));
        const beforeError = inventory(output);
        put(directory, 'src/Main.go', 'func Broken(');
        const syntax = invoke('ffi-syntax-error', true, {}, [1]);
        assert.match(syntax.stderr, /FFI error in module Main \(/);
        assert(syntax.stderr.includes(join(directory, 'src/Main.go')));
        const diagnostic = syntax.stderr.match(/Go FFI parse failed: ([^\n]+)/)?.[1];
        assert.match(diagnostic, /expected '\)', found 'EOF'/);
        if (expectedDiagnostic) assert.equal(diagnostic, expectedDiagnostic);
        else expectedDiagnostic = diagnostic;
        assert.deepEqual(readFileSync(join(output, '.gopurs-cache/current.json')), pointer);
        assert.deepEqual(inventory(output), beforeError);
        put(directory, 'src/Main.go', nativeFfi);
        assert.match(invoke('ffi-error-recovery', true).stderr, /cache hit:/);
        assert.deepEqual(compare('ffi-error-recovery'), beforeError);
        renameSync(join(directory, 'src/Main.go'), join(workspace, mode + '-ffi.go'));
        invoke('ffi-missing', true);
        compare('ffi-missing');
        run(mode + '-ffi-missing-build', 'go', ['build', '-p', '1', '-o', binary, './main/main.go'], output);
        const missing = run(mode + '-ffi-missing-execute', binary, [], output, {}, [2]);
        assert.match(missing.stderr, /FFI not implemented:/);
        renameSync(join(workspace, mode + '-ffi.go'), join(directory, 'src/Main.go'));
        invoke('ffi-returned', true);
        assert.deepEqual(compare('ffi-returned'), beforeError);
        assert.equal(run(mode + '-ffi-returned-execute', 'go', ['run', './main/main.go'], output).stdout, before);
        invoke('monolithic-restored', false);
        assert.ok(existsSync(join(output, 'purescript/Main.go')));
        assert.equal(run(mode + '-restored-execute', 'go', ['run', './main/main.go'], output).stdout, before);
        compare('restored');
        renameSync(output, join(directory, 'output-' + mode));
    });
    put(workspace, 'summary.json', JSON.stringify({ completed: true, modes, commands: commands.length }));
    passed = true;
});
