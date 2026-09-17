import assert from 'node:assert/strict';
import { spawnSync } from 'node:child_process';
import { mkdirSync, mkdtempSync, rmSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join } from 'node:path';
import test from 'node:test';
import { empty as emptySet } from '../output/Data.Set/index.js';
import * as Ref from '../output/Effect.Ref/index.js';
import * as Go from '../output/Gopurs.GoAst/index.js';
import { unboxGoExpr } from '../output/Gopurs.GoConversions/index.js';
import { printGoExpr } from '../output/Gopurs.Printer/index.js';
import { runtimeGoCode } from '../output/Gopurs.Runtime/index.js';

test('Value arrays reuse immutable storage while other element representations are converted', t => {
    const state = Ref.new({ declarations: [], globalId: 0, reboxPairs: emptySet })();
    const unbox = expression => from => to => printGoExpr(unboxGoExpr(state)('Probe')(expression)(from)(to));
    const values = new Go.TypeNativeArray(Go.TypeValue.value);
    const ints = new Go.TypeNativeArray(Go.TypeInt64.value);
    const strings = new Go.TypeNativeArray(Go.TypeString.value);
    const sharedNode = unboxGoExpr(state)('Probe')
        (new Go.GoCall(new Go.GoVar('readSource'), [new Go.GoVar('source')]))(Go.TypeValue.value)(values);
    const shared = printGoExpr(sharedNode);
    const indexed = printGoExpr(new Go.GoIndex(sharedNode, new Go.GoInt(0)));
    const intConversion = unbox(new Go.GoVar('source'))(Go.TypeValue.value)(ints);
    const stringConversion = unbox(new Go.GoVar('source'))(Go.TypeValue.value)(strings);
    const nativeConversion = unbox(new Go.GoVar('source'))(ints)(values);
    const directory = mkdtempSync(join(tmpdir(), 'gopurs-value-array-unboxing-'));
    t.after(() => rmSync(directory, { recursive: true, force: true }));
    mkdirSync(join(directory, 'gopurs_runtime'));
    writeFileSync(join(directory, 'go.mod'), 'module gopurs/output\n\ngo 1.22\n');
    writeFileSync(join(directory, 'gopurs_runtime/runtime.go'), runtimeGoCode);
    writeFileSync(join(directory, 'main.go'), `package main
import ("fmt"; "testing"; "gopurs/output/gopurs_runtime")
var reads int
var sink []gopurs_runtime.Value
func readSource(source gopurs_runtime.Value) gopurs_runtime.Value { reads++; return source }
//go:noinline
func readValues(source gopurs_runtime.Value) []gopurs_runtime.Value { return ${shared} }
func readFirstStructured(source gopurs_runtime.Value) gopurs_runtime.Value { return ${indexed} }
func readFirstPostfix(source gopurs_runtime.Value) gopurs_runtime.Value { return ${shared}[0] }
func readTail(source gopurs_runtime.Value) []gopurs_runtime.Value { return ${shared}[1:] }
func readInts(source gopurs_runtime.Value) []int64 { return ${intConversion} }
func readStrings(source gopurs_runtime.Value) []string { return ${stringConversion} }
func convertNativeInts(source []int64) []gopurs_runtime.Value { return ${nativeConversion} }
func main() {
    sourceValues := []gopurs_runtime.Value{gopurs_runtime.Int(-7), gopurs_runtime.Int(42)}
    source := gopurs_runtime.Array(sourceValues)
    values := readValues(source)
    if reads != 1 || len(values) != 2 || &values[0] != &sourceValues[0] { panic("Value array copied or source evaluated more than once") }
    if values[0].IntVal != -7 || values[1].IntVal != 42 { panic("Value payload changed") }
    if len(readValues(gopurs_runtime.Array(nil))) != 0 { panic("empty Value array") }
    reads = 0
    if readFirstStructured(source).IntVal != -7 || reads != 1 { panic("structured indexing precedence or evaluation") }
    reads = 0
    if readFirstPostfix(source).IntVal != -7 || reads != 1 { panic("postfix indexing precedence or evaluation") }
    reads = 0
    tail := readTail(source)
    if reads != 1 || len(tail) != 1 || tail[0].IntVal != 42 || &tail[0] != &sourceValues[1] { panic("slicing precedence or evaluation") }

    converted := readInts(source)
    if converted[0] != -7 || converted[1] != 42 { panic("Int conversion skipped") }
    converted[0] = 99
    if values[0].IntVal != -7 { panic("Int conversion modified the source") }
    convertedStrings := readStrings(gopurs_runtime.Array([]gopurs_runtime.Value{gopurs_runtime.Str("hello"), gopurs_runtime.Str("😀")}))
    if convertedStrings[0] != "hello" || convertedStrings[1] != "😀" { panic("String conversion skipped") }
    nativeInts := []int64{-5, 17}
    boxedInts := convertNativeInts(nativeInts)
    if boxedInts[0].IntVal != -5 || boxedInts[1].IntVal != 17 { panic("native Int boxing skipped") }
    boxedInts[0] = gopurs_runtime.Int(123)
    if nativeInts[0] != -5 { panic("native Int boxing aliases different representations") }

    for _, size := range []int{16, 32768} {
        source := gopurs_runtime.Array(make([]gopurs_runtime.Value, size))
        allocations := testing.AllocsPerRun(100, func() { sink = readValues(source) })
        if allocations != 0 { panic(fmt.Sprintf("size=%d allocations=%v", size, allocations)) }
        fmt.Printf("size=%d allocations=%v\\n", size, allocations)
    }
    fmt.Println("source evaluated once; immutable storage shared; typed conversions preserved")
}
`);
    const result = spawnSync('go', ['run', '.'], {
        cwd: directory, encoding: 'utf8', timeout: 30_000,
        env: { ...process.env, GOWORK: 'off' },
    });
    assert.ifError(result.error);
    assert.equal(result.status, 0, result.stderr);
    assert.match(result.stdout, /size=16 allocations=0\nsize=32768 allocations=0/);
    assert.match(result.stdout, /typed conversions preserved/);
});
