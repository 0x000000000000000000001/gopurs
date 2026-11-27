import assert from 'node:assert/strict';
import { spawnSync } from 'node:child_process';
import { mkdirSync, readFileSync, readdirSync, rmSync, symlinkSync, writeFileSync } from 'node:fs';
import { delimiter, join } from 'node:path';
import { fileURLToPath } from 'node:url';
import test from 'node:test';
import { corePackages, createWorkspace, prepareFixture } from './test-workspace.mjs';
import { findTypedCompiler } from './build/native-workspace.mjs';

const root = fileURLToPath(new URL('../', import.meta.url));
const modes = [{ name: 'js', GOPURS_JS: '1' }, { name: 'native-parallel', GOPURS_JS: '0' }]
    .filter(mode => !process.env.GOPURS_TEST_MODES || process.env.GOPURS_TEST_MODES.split(',').includes(mode.name));
const profiles = (process.env.GOPURS_TEST_PROFILES || 'release,dev').split(',');
const fixture = join(root, 'tools/fixtures/array-traverse-fold.purs');
const inputs = [[], [1, 2, 3], [2147483647, -2147483648, -1, 0]];
const parse = stdout => stdout.trim().split('\n').map(line => JSON.parse(line));
function expected(seed = 7, weight = 1, trace = false) {
    const out = [];
    const consume = values => values.reduce((total, value, index) => {
        if (trace) out.push(100 + index);
        return (total + (((index + weight) * value) | 0)) | 0;
    }, seed);
    for (const values of inputs) {
        const a = consume(values.map(x => (x + 1) | 0));
        const b = consume(values.map((x, i) => (x + i) | 0));
        const c = consume(values.map(x => (x + 3) | 0));
        const d = (consume(values.map(x => (x + 1) | 0)) + values.length) | 0;
        const e = values.includes(2) ? -1 : consume(values.map(x => (x + 1) | 0));
        out.push([a, b, c, d, e], values, values.toReversed());
    }
    out.push(consume([2, 3, 4]), [1, 2, 3], 'mapper failure', [1, 2]);
    out.push(consume([2, 3, 4]), [1], consume([]), []);
    out.push(consume([2, 3, 4]), [9]);
    return out;
}
function generatedText(directory) {
    return readdirSync(directory, { withFileTypes: true }).filter(entry => entry.name !== '.gopurs-cache')
        .map(entry => { const path = join(directory, entry.name); return entry.isDirectory() ? generatedText(path) : entry.name.endsWith('.go') ? readFileSync(path, 'utf8') : ''; }).join('\n');
}
function worker(source, name) {
    const start = source.indexOf(`func ${name}(`); if (start < 0) return '';
    const end = source.indexOf('\nfunc ', start + 1); return source.slice(start, end < 0 ? undefined : end);
}
function snapshots(source) {
    return Object.fromEntries(['Get_Main_direct', 'Call_Main_direct', 'Get_Main_indexed', 'Call_Main_indexed', 'Get_Main_captured', 'Call_Main_captured',
        'Get_Main_withUnknown', 'Call_Main_withUnknown', 'Get_Main_withFactory', 'Call_Main_withFactory', 'Get_Main_withBarrier', 'Call_Main_withBarrier', 'Get_Main_mayFail', 'Call_Main_mayFail', 'Call_Subject_consume']
        .map(name => [name, worker(source, name)]));
}
function abi(source) { const match = worker(source, 'Call_Subject_consume').match(/^func Call_Subject_consume\([^\n]+/); assert.ok(match, 'native consumer worker remains present'); return match[0]; }

// Each transition keeps consume :: Array Int -> Int and the same native worker ABI.
// The cache must depend on descriptor contents and on its later absence.
function variants(original) {
    const weighted = original.replace('(index + 1)', '(index + 2)');
    const seeded = weighted.replace(' 7 values', ' 19 values');
    const unknown = '-- @inline Subject.step never\n' + seeded
        .replace('import Prelude', 'import Prelude\nimport Effect.Console (logShow)\nimport Effect.Unsafe (unsafePerformEffect)')
        .replace('foldlWithIndex (\\index total value -> total + (index + 2) * value) 19 values', 'foldlWithIndex step 19 values')
        + '\nstep :: Int -> Int -> Int -> Int\nstep index total value = unsafePerformEffect do\n  logShow (100 + index)\n  pure (total + (index + 2) * value)\n';
    return [
        { name: 'callback', source: weighted, expected: expected(7, 2), eligible: true },
        { name: 'seed', source: seeded, expected: expected(19, 2), eligible: true },
        { name: 'unknown', source: unknown, expected: expected(19, 2, true), eligible: false },
    ];
}

test('traverse folds preserve source results, phases and imported descriptor cache transitions', async t => {
    assert.ok(modes.length); const workspace = createWorkspace(); console.log(`Array traverse fold fixture workspace: ${workspace}`);
    let completed = 0; t.after(() => { if (completed === modes.length && process.env.GOPURS_TEST_KEEP_WORKSPACE !== '1') rmSync(workspace, { recursive: true, force: true }); });
    const typedBin = join(workspace, 'typed-bin'); mkdirSync(typedBin); symlinkSync(findTypedCompiler(root, process.env.GOPURS_PURS), join(typedBin, 'purs'));
    const environment = { ...process.env, PATH: [typedBin, join(root, 'node_modules/.bin'), process.env.PATH].join(delimiter), GOWORK: 'off', GOPURS_DAEMON: '0', GOPURS_RESIDENT: '0',
        GOPURS_RUST: '0', GOPURS_GO_PACKAGES: '1', GOPURS_GO_PACKAGE_OVERLAY: '1', GOPURS_BUILD_MANIFEST: '1', GOPURS_PREPARE_JOBS: '2', GOPURS_PBO_JOBS: '2', GOPURS_EMIT_JOBS: '2', GOPURS_PIPELINE: '1', GOPURS_FORCE_BUILD: '1' };
    for (const key of ['GOPURS_ALLOC_PROFILE', 'GOPURS_COMPILER_SNAPSHOT', 'PPROF']) delete environment[key];
    for (const [hostIndex, mode] of modes.entries()) await t.test(mode.name, () => {
        const directory = prepareFixture(root, workspace, fixture, hostIndex, corePackages(root)).directory;
        const subject = join(directory, 'src/Subject.purs'); const original = readFileSync(subject, 'utf8'); let sequence = 0;
        const run = (label, command, args, cwd = directory, overrides = {}) => {
            console.log(`${mode.name}: ${label}`); const result = spawnSync(command, args, { cwd, env: { ...environment, GOPURS_JS: mode.GOPURS_JS, ...overrides }, encoding: 'utf8', timeout: 240000, maxBuffer: 16 * 1024 * 1024 });
            writeFileSync(join(directory, 'logs', `${sequence++}-${label}.log`), (result.stdout || '') + (result.stderr || '')); assert.ifError(result.error); assert.equal(result.status, 0, `${label}\n${result.stdout}\n${result.stderr}`); return result.stdout;
        };
        const compiler = process.env.GOPURS_TEST_COMPILER || join(root, 'bin/gopurs');
        const javascript = cwd => parse(run('javascript', 'node', ['--input-type=module', '-e', 'import {main} from "./output/Main/index.js";main();'], cwd));
        run('frontend', 'spago', ['build', '-q', '--offline']); const oracle = javascript(directory); assert.deepEqual(oracle, expected());
        if (process.env.GOPURS_TEST_FRONTEND_ONLY !== '1') for (const [profileIndex, profile] of profiles.entries()) {
            writeFileSync(subject, original); run(`${profile}-frontend`, 'spago', ['build', '-q', '--offline']);
            run(profile, compiler, [`--${profile}`, '--main', 'Main']);
            const baselineSource = generatedText(join(directory, 'output/purescript')); const before = snapshots(baselineSource);
            writeFileSync(join(directory, 'logs', `${profile}-workers.json`), JSON.stringify(before, null, 2) + '\n');
            const baselineAbi = abi(baselineSource);
            if (profile === 'release') for (const name of ['direct', 'indexed']) assert.match(before[`Get_Main_${name}`] + before[`Call_Main_${name}`], /traverseFold_fold_/, `${name}: actual imported descriptor must activate`);
            for (const name of ['captured', 'withUnknown', 'mayFail', 'withBarrier']) assert.doesNotMatch(before[`Get_Main_${name}`] + before[`Call_Main_${name}`], /traverseFold_fold_/, `${name}: unknown or captured mapper remains a barrier`);
            assert.deepEqual(parse(run(`${profile}-go`, 'go', ['run', './main'], join(directory, 'output'))), oracle);
            if (process.env.GOPURS_TEST_SKIP_INCREMENTAL === '1') continue;
            for (const [variantIndex, variant] of variants(original).entries()) {
                writeFileSync(subject, variant.source); run(`${profile}-${variant.name}-frontend`, 'spago', ['build', '-q', '--offline']);
                const changedOracle = javascript(directory); assert.deepEqual(changedOracle, variant.expected, `${variant.name}: independent source oracle`);
                run(`${profile}-${variant.name}-warm`, compiler, [`--${profile}`, '--main', 'Main'], directory, { GOPURS_FORCE_BUILD: '0' });
                const warmSource = generatedText(join(directory, 'output/purescript')); assert.equal(abi(warmSource), baselineAbi, 'descriptor transitions retain native ABI');
                if (profile === 'release') {
                    const mapped = worker(warmSource, 'Get_Main_direct') + worker(warmSource, 'Call_Main_direct');
                    if (variant.eligible) assert.match(mapped, /traverseFold_fold_/); else assert.doesNotMatch(mapped, /traverseFold_fold_/);
                }
                const warm = parse(run(`${profile}-${variant.name}-warm-go`, 'go', ['run', './main'], join(directory, 'output'))); assert.deepEqual(warm, changedOracle);
                const clean = prepareFixture(root, workspace, fixture, 100 + hostIndex * 20 + profileIndex * 3 + variantIndex, corePackages(root)).directory;
                writeFileSync(join(clean, 'src/Subject.purs'), variant.source); run(`${profile}-${variant.name}-clean-frontend`, 'spago', ['build', '-q', '--offline'], clean);
                run(`${profile}-${variant.name}-clean`, compiler, [`--${profile}`, '--main', 'Main'], clean);
                const rebuilt = parse(run(`${profile}-${variant.name}-clean-go`, 'go', ['run', './main'], join(clean, 'output'))); assert.deepEqual(rebuilt, warm, 'changed descriptor: incremental equals clean');
                const cleanSource = generatedText(join(clean, 'output/purescript')); assert.equal(abi(cleanSource), baselineAbi);
                assert.deepEqual(snapshots(warmSource), snapshots(cleanSource), 'changed consumer metadata must regenerate the same caller workers as a clean build');
                writeFileSync(join(directory, 'logs', `${profile}-${variant.name}-cache.json`), JSON.stringify({ transition: variant.name, eligible: variant.eligible, same_native_abi: baselineAbi, warm_equals_clean_outputs: true, warm_equals_clean_workers: true, observations: warm.length }, null, 2) + '\n');
            }
        }
        completed++;
    });
});
