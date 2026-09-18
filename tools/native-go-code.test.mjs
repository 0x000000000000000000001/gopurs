// Run after a native bootstrap:
// GOPURS_NATIVE_OUTPUT=/path/to/bootstrap/output node --test tools/native-go-code.test.mjs
import assert from 'node:assert/strict';
import { spawnSync } from 'node:child_process';
import { mkdtempSync, rmSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join, resolve } from 'node:path';
import test from 'node:test';
import { referencedImports } from '../output/Gopurs.GoCode/index.js';

const nativeOutput = process.env.GOPURS_NATIVE_OUTPUT;
const boundaryCases = [
    ['/', []],
    ['math.Abs(1)/', ['math']],
    ['math', []],
    ['math.', ['math']],
    ['/*', []],
    ['/* math.Abs*', []],
    ['// unsafe.Pointer', []],
    ['"math.Abs', []],
    ['"math.Abs\\', []],
    ["'\\", []],
    ['`unsafe.Pointer', []],
];

test('Go import scanning handles lookahead and escapes at the end of input', () => {
    for (const [source, expected] of boundaryCases) {
        assert.deepEqual(referencedImports(source), expected, JSON.stringify(source));
    }
});

test('native Go import scanning matches JavaScript at boundaries and uses bounded stack', {
    skip: nativeOutput ? false : 'requires GOPURS_NATIVE_OUTPUT from a native bootstrap',
}, t => {
    const cases = [
        ['', []],
        ...boundaryCases,
        ['math.Abs(1); unsafe.Pointer(nil); sync.Once{}; gopurs_runtime.Value{}',
            ['gopurs/output/gopurs_runtime', 'math', 'sync', 'unsafe']],
        ['"math.Abs sync.Once"; `unsafe.Pointer`; /* gopurs_runtime.Value */', []],
        ['"escaped \\" math.Abs"; // unsafe.Pointer\nsync.Once{}', ['sync']],
        ['ÿmath.Abs(1); ≠unsafe.Pointer(nil); ·sync.Once{}; 😀gopurs_runtime.Value{}', []],
        ['x '.repeat(1000) + 'math.Abs(1)', ['math']],
        ['x; '.repeat(10_000) + 'unsafe.Pointer(nil)', ['unsafe']],
        ['/* 😀 math.Abs */ "ÿ unsafe.Pointer"; '.repeat(1000) + 'sync.Once{}', ['sync']],
        ['"' + 'x'.repeat(40_000), []],
        ['/*' + 'x'.repeat(40_000), []],
        ['// ' + 'x'.repeat(40_000), []],
        ['x'.repeat(40_000) + '.Abs(1)', []],
    ];
    const directory = mkdtempSync(join(tmpdir(), 'gopurs-native-go-code-'));
    t.after(() => rmSync(directory, { recursive: true, force: true }));
    writeFileSync(join(directory, 'go.mod'), `module scannerprobe

go 1.22

require gopurs/output v0.0.0
replace gopurs/output => ${JSON.stringify(resolve(nativeOutput))}
`);
    writeFileSync(join(directory, 'main.go'), `package main
import (
    "encoding/json"
    "os"
    "runtime/debug"
    "gopurs/output/purescript"
)
func main() {
    // A two-kilobyte identifier sequence exhausted this budget before the fix.
    // Avoid relying on a machine's memory or on Go's default 1 GiB stack limit.
    debug.SetMaxStack(1 << 20)
    sources := []string{${cases.map(([source]) => JSON.stringify(source)).join(',\n')}}
    results := make([][]string, 0, len(sources))
    for _, source := range sources {
        imports := purescript.Call_Gopurs_GoCode_referencedImports(source)
        if imports == nil { imports = []string{} }
        results = append(results, imports)
    }
    json.NewEncoder(os.Stdout).Encode(results)
}
`);
    const result = spawnSync('go', ['run', '-trimpath', '.'], {
        cwd: directory, encoding: 'utf8', timeout: 60_000,
        env: { ...process.env, GOWORK: 'off' }, maxBuffer: 1024 * 1024,
    });
    assert.ifError(result.error);
    assert.equal(result.status, 0, result.stderr);
    const nativeResults = JSON.parse(result.stdout);
    const jsResults = cases.map(([source]) => referencedImports(source));
    assert.deepEqual(jsResults, cases.map(([, expected]) => expected));
    assert.deepEqual(nativeResults, jsResults);
});
