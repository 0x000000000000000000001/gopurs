import assert from 'node:assert/strict';
import test from 'node:test';
import { empty as emptySet } from '../output/Data.Set/index.js';
import * as Ref from '../output/Effect.Ref/index.js';
import * as Go from '../output/Gopurs.GoAst/index.js';
import { boxGoExpr, unboxGoExpr } from '../output/Gopurs.GoConversions/index.js';
import { printGoExpr } from '../output/Gopurs.Printer/index.js';
import { runtimeGoCode } from '../output/Gopurs.Runtime.Source/index.js';
import { goModule, runGo, writeGoFile } from './lib/go.mjs';

test('immediate Int array conversion copies fuse without exposing aliases or skipping normalization', t => {
    const state = Ref.new({ declarations: [], globalId: 0, reboxPairs: emptySet })();
    const ints = new Go.TypeNativeArray(Go.TypeInt64.value);
    const values = new Go.TypeNativeArray(Go.TypeValue.value);
    const box = (expression, type) => boxGoExpr(state)('Probe')(expression)(type);
    const unbox = (expression, type) => unboxGoExpr(state)('Probe')(expression)(Go.TypeValue.value)(type);
    const read = name => new Go.GoCall(new Go.GoVar(name), [new Go.GoVar('input')]);
    const native = read('readNative');
    const nativeCopy = printGoExpr(unbox(box(native, ints), ints));
    const fullChain = printGoExpr(box(unbox(box(unbox(box(native, ints), values), values), ints), ints));
    const normalizeValue = printGoExpr(box(unbox(read('readValue'), values), values));
    const normalizeInts = printGoExpr(box(unbox(read('readValue'), ints), ints));
    const stored = printGoExpr(unbox(new Go.GoVar('stored'), ints));
    assert.doesNotMatch(nativeCopy, /boxed :=|unboxed :=/);
    assert.match(nativeCopy, /copy\(arrayCopy_output, arrayCopy_input\)/);
    assert.equal((fullChain.match(/boxed :=/g) ?? []).length, 1);
    assert.doesNotMatch(fullChain, /unboxed :=|arrayCopy_/);
    assert.match(normalizeValue, /gopurs_runtime.Array\(/);
    assert.match(normalizeInts, /unboxed :=/);
    assert.match(stored, /unboxed :=/);
    for (const expression of [nativeCopy, fullChain]) {
        assert.equal((expression.match(/readNative\(input\)/g) ?? []).length, 1);
    }
    const directory = goModule(t, 'gopurs-array-conversion-fusion-', { runtime: runtimeGoCode });
    writeGoFile(directory, 'conversions_test.go', `package main
import ("testing"; "slices"; "gopurs/output/gopurs_runtime")
var reads int
func readNative(input []int64) []int64 { reads++; return input }
func readValue(input gopurs_runtime.Value) gopurs_runtime.Value { reads++; return input }
func copyNative(input []int64) []int64 { return ${nativeCopy} }
func fullChain(input []int64) gopurs_runtime.Value { return ${fullChain} }
func normalizeValue(input gopurs_runtime.Value) gopurs_runtime.Value { return ${normalizeValue} }
func normalizeInts(input gopurs_runtime.Value) gopurs_runtime.Value { return ${normalizeInts} }
func valueSlice(value gopurs_runtime.Value) []gopurs_runtime.Value { return *(*[]gopurs_runtime.Value)(value.UnsafePtr) }
func TestFusion(t *testing.T) {
    for _, input := range [][]int64{nil, {}, {0}, {-1, -2147483648, 2147483648, 4294967295, -9223372036854775808, 9223372036854775807}} {
        original := slices.Clone(input)
        reads = 0
        copied := copyNative(input)
        if reads != 1 || copied == nil || !slices.Equal(copied, input) { t.Fatal("copy changed values, nil handling or evaluation") }
        reads = 0
        boxed := fullChain(input)
        if reads != 1 || boxed.Type != gopurs_runtime.TypeArray || boxed.IntVal != 0 { t.Fatal("chain changed evaluation or array normalization") }
        output := valueSlice(boxed)
        if output == nil || len(output) != len(input) { t.Fatal("chain changed nil/empty handling") }
        for i, scalar := range input {
            if output[i] != gopurs_runtime.Int(scalar) { t.Fatalf("scalar normalization changed at %d", i) }
        }
        if len(input) > 0 {
            copied[0] = 700
            output[0] = gopurs_runtime.Int(800)
            if !slices.Equal(input, original) { t.Fatal("output aliases input") }
            input[0] = 900
            if copied[0] != 700 || output[0].IntVal != 800 { t.Fatal("input aliases output") }
        }
    }
}
func TestOpaqueNormalization(t *testing.T) {
    stringValue := gopurs_runtime.StrWithCodeUnits("a😀é", 4)
    for _, input := range [][]gopurs_runtime.Value{nil, {}, {stringValue, gopurs_runtime.Bool(true), gopurs_runtime.Int(4294967295)}} {
        original := slices.Clone(input)
        value := gopurs_runtime.Array(input)
        value.Type = gopurs_runtime.TypeInt
        value.IntVal = 99
        reads = 0
        normalized := normalizeValue(value)
        if reads != 1 || normalized.Type != gopurs_runtime.TypeArray || normalized.IntVal != 0 { t.Fatal("opaque array tag normalization skipped") }
        normalizedSlice := valueSlice(normalized)
        if (normalizedSlice == nil) != (input == nil) || !slices.Equal(normalizedSlice, input) { t.Fatal("Value view changed data/metadata/nil") }
        reads = 0
        normalized = normalizeInts(value)
        if reads != 1 || normalized.Type != gopurs_runtime.TypeArray || normalized.IntVal != 0 { t.Fatal("opaque Int normalization changed") }
        normalizedSlice = valueSlice(normalized)
        if normalizedSlice == nil || len(normalizedSlice) != len(input) { t.Fatal("Int view changed nil/empty handling") }
        for i, item := range input {
            if normalizedSlice[i] != gopurs_runtime.Int(item.IntVal) { t.Fatal("opaque scalar tag/metadata normalization skipped") }
        }
        if len(input) > 0 {
            normalizedSlice[0] = gopurs_runtime.Int(7)
            if !slices.Equal(input, original) { t.Fatal("normalization mutated original Value buffer") }
        }
    }
}
`);
    runGo(directory, ['test', '-count=1', '.'], { timeout: 30_000 });
});
