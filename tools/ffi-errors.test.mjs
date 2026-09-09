import assert from 'node:assert/strict';
import { spawnSync } from 'node:child_process';
import { copyFileSync, mkdirSync, mkdtempSync, readFileSync, rmSync, symlinkSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { dirname, join } from 'node:path';
import test from 'node:test';
import { fileURLToPath, pathToFileURL } from 'node:url';
import { Left, Right } from '../output/Data.Either/index.js';
import { decodeFfiDecls, extractFfiDecls } from '../output/Gopurs.FfiSupport/index.js';
import { TFunc } from '../output/Gopurs.FfiTypes/index.js';

const root = dirname(dirname(fileURLToPath(import.meta.url)));
const runner = join(root, 'tools/ffi-runner.mjs');
const source = { moduleName: 'Test.Invalid', path: '/project/ffi/Test.Invalid.go' };

function runNode(args, input = '') {
    const result = spawnSync(process.execPath, args, {
        input, encoding: 'utf8', timeout: 10_000, maxBuffer: 1024 * 1024,
    });
    assert.ifError(result.error);
    assert.equal(result.signal, null);
    return result;
}

function sandbox(t) {
    const directory = mkdtempSync(join(tmpdir(), 'gopurs-ffi-errors-'));
    t.after(() => rmSync(directory, { recursive: true, force: true }));
    writeFileSync(join(directory, 'package.json'), '{"type":"module"}');
    mkdirSync(join(directory, 'tools'));
    return directory;
}

function isolatedApi(directory) {
    const original = join(root, 'output/Gopurs.FfiSupport');
    const output = join(directory, 'output');
    const target = join(output, 'Gopurs.FfiSupport');
    mkdirSync(target, { recursive: true });
    for (const file of ['index.js', 'foreign.js']) {
        copyFileSync(join(original, file), join(target, file));
    }
    // Preserve the real compiled dependencies while giving this API its own runner.
    const imports = readFileSync(join(target, 'index.js'), 'utf8');
    const dependencies = new Set(Array.from(imports.matchAll(/from ["']\.\.\/([^/"']+)\/index\.js["']/g), match => match[1]));
    for (const dependency of dependencies) {
        symlinkSync(join(root, 'output', dependency), join(output, dependency), 'dir');
    }
    return join(target, 'index.js');
}

function caughtError(modulePath, expression) {
    return runNode(['--input-type=module', '-e', `
        const api = await import(${JSON.stringify(pathToFileURL(modulePath).href)});
        try {
            ${expression};
            process.exitCode = 2;
        } catch (error) {
            process.stdout.write(error.message);
        }
    `]);
}

function assertContextualFailure(result) {
    assert.equal(result.status, 0, result.stderr);
    assert.equal(result.stderr, '', 'the caught error must not also be logged');
    assert.ok(result.stdout.includes(source.moduleName), result.stdout);
    assert.ok(result.stdout.includes(source.path), result.stdout);
}

test('valid empty Go and hidden declarations remain successful empty results', () => {
    for (const content of ['', 'package ffi\nfunc hidden() {}\nvar hiddenValue int']) {
        const result = runNode([runner], content);
        assert.equal(result.status, 0, result.stderr);
        assert.equal(result.stderr, '');
        assert.deepEqual(JSON.parse(result.stdout), []);
        assert.deepEqual(extractFfiDecls(source)(content)(), []);
    }
});

test('valid declarations retain structured callback types through the compiled API', () => {
    const declarations = extractFfiDecls(source)('func Call(cb func(int) string) string { return cb(1) }')();
    assert.equal(declarations.length, 1);
    assert.equal(declarations[0].name, 'Call');
    assert.equal(declarations[0].isVar, false);
    assert.equal(declarations[0].args.length, 1);
    assert.ok(declarations[0].args[0] instanceof TFunc);
});

test('invalid Go fails without partial JSON and keeps original source positions', () => {
    for (const [content, line] of [
        ['func Valid() {}\nfunc Broken(', 2],
        ['package ffi\nfunc Valid() {}\nfunc Broken(', 3],
    ]) {
        const result = runNode([runner], content);
        assert.equal(result.status, 1);
        assert.equal(result.stdout, '');
        assert.match(result.stderr, new RegExp(`(?:^|\\s)${line}:13:`));
        assert.throws(() => extractFfiDecls(source)(content)(), error =>
            error.message.includes(source.moduleName) && error.message.includes(source.path));
    }
});

test('the pure decoder distinguishes empty declarations, malformed JSON and invalid shapes', () => {
    const empty = decodeFfiDecls('[]');
    assert.ok(empty instanceof Right);
    assert.deepEqual(empty.value0, []);
    for (const invalid of ['{', '{}', '[{"name":"Incomplete"}]']) {
        assert.ok(decodeFfiDecls(invalid) instanceof Left, invalid);
    }
});

for (const [label, response] of [['malformed JSON', '{'], ['invalid declaration shape', '[{"name":"Incomplete"}]']]) {
    test(`compiled API contextualizes ${label} without logging twice`, t => {
        const directory = sandbox(t);
        const api = isolatedApi(directory);
        writeFileSync(join(directory, 'tools/ffi-runner.mjs'), `process.stdout.write(${JSON.stringify(response)});`);
        const result = caughtError(api, `api.extractFfiDecls(${JSON.stringify(source)})("")()`);
        assertContextualFailure(result);
    });
}

for (const failure of ['missing', 'corrupt']) {
    test(`${failure} WASM makes the runner fail without emitting JSON`, t => {
        const directory = sandbox(t);
        for (const file of ['ffi-runner.mjs', 'wasm_exec.js']) {
            copyFileSync(join(root, 'tools', file), join(directory, 'tools', file));
        }
        if (failure === 'corrupt') writeFileSync(join(directory, 'tools/ffi_gen.wasm'), 'invalid wasm');
        const result = runNode([join(directory, 'tools/ffi-runner.mjs')]);
        assert.equal(result.status, 1);
        assert.equal(result.stdout, '');
        assert.notEqual(result.stderr.trim(), '');
    });

    test(`${failure} runner is contextualized once by the compiled API`, t => {
        const directory = sandbox(t);
        const api = isolatedApi(directory);
        if (failure === 'corrupt') writeFileSync(join(directory, 'tools/ffi-runner.mjs'), 'const = invalid;');
        const result = caughtError(api, `api.extractFfiDecls(${JSON.stringify(source)})("")()`);
        assertContextualFailure(result);
    });
}

test('raw transport propagates a child diagnostic without writing it to stderr itself', t => {
    const directory = sandbox(t);
    const support = join(directory, 'src/Gopurs');
    mkdirSync(support, { recursive: true });
    copyFileSync(join(root, 'src/Gopurs/FfiSupport.js'), join(support, 'FfiSupport.js'));
    writeFileSync(join(directory, 'tools/ffi-runner.mjs'), 'process.stderr.write("parser failed once"); process.exitCode = 1;');
    const result = caughtError(join(support, 'FfiSupport.js'), 'api.extractFfiAstImpl("")()');
    assert.equal(result.status, 0, result.stderr);
    assert.equal(result.stderr, '');
    assert.match(result.stdout, /parser failed once/);
});
