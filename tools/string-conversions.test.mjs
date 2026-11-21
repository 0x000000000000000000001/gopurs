import assert from 'node:assert/strict';
import test from 'node:test';
import { empty as emptySet } from '../output/Data.Set/index.js';
import * as Ref from '../output/Effect.Ref/index.js';
import * as Go from '../output/Gopurs.GoAst/index.js';
import { boxGoExpr, unboxGoExpr } from '../output/Gopurs.GoConversions/index.js';
import { printGoExpr } from '../output/Gopurs.Printer/index.js';
import { runtimeGoCode } from '../output/Gopurs.Runtime.Source/index.js';
import { goModule, runGo, writeGoFile } from './lib/go.mjs';

test('generated String conversions preserve metadata and compute literal lengths at compile time', t => {
    const state = Ref.new({ declarations: [], globalId: 0, reboxPairs: emptySet })();
    const box = expression => boxGoExpr(state)('Probe')(expression)(Go.TypeString.value);
    const input = new Go.GoCall(new Go.GoVar('readValue'), [new Go.GoVar('value')]);
    const unboxed = unboxGoExpr(state)('Probe')(input)(Go.TypeValue.value)(Go.TypeString.value);
    const roundtrip = printGoExpr(box(unboxed));
    const native = printGoExpr(box(new Go.GoCall(new Go.GoVar('readNative'), [new Go.GoVar('value')])));
    const strings = ['', 'abc', '\0', '"\\\n', 'é', 'a😀é', '𝌆', '\ud800', '\udfff',
        'a\ud800b', '\ud800\udfff', '\ud834\udf06'];
    const fixtures = strings.map(text => ({ text, expression: printGoExpr(box(new Go.GoString(text))), units: text.length }));
    assert.match(roundtrip, /NormalizeStringValue\(readValue\(value\)\)/);
    assert.match(native, /Str\(readNative\(value\)\)/);
    const quote = text => printGoExpr(new Go.GoString(text));
    const directory = goModule(t, 'gopurs-string-conversions-', { runtime: runtimeGoCode });
    writeGoFile(directory, 'conversions_test.go', `package main
import ("testing"; "gopurs/output/gopurs_runtime")
var reads int
func readValue(value gopurs_runtime.Value) gopurs_runtime.Value { reads++; return value }
func readNative(value string) string { reads++; return value }
func roundtrip(value gopurs_runtime.Value) gopurs_runtime.Value { return ${roundtrip} }
func nativeBox(value string) gopurs_runtime.Value { return ${native} }
func TestStringConversions(t *testing.T) {
    fixtures := []struct { value gopurs_runtime.Value; text string; units int } {
        ${fixtures.map(({ expression, text, units }) => `{${expression}, ${quote(text)}, ${units}}`).join(',\n')},
    }
    for _, fixture := range fixtures {
        if fixture.value.StrVal() != fixture.text { t.Fatalf("literal bytes changed: %x", fixture.text) }
        if got := gopurs_runtime.StringCodeUnitsLength(fixture.value); got != fixture.units { t.Fatalf("literal %x: length %d, want %d", fixture.text, got, fixture.units) }
        if got, cached := gopurs_runtime.CachedStringCodeUnits(fixture.value); fixture.text != "" && (!cached || got != fixture.units) { t.Fatalf("literal %x lost its known length", fixture.text) }
        reads = 0
        if value := roundtrip(fixture.value); value != fixture.value || reads != 1 { t.Fatal("String roundtrip lost metadata or repeated evaluation") }
        reads = 0
        value := nativeBox(fixture.text)
        if value.StrVal() != fixture.text || reads != 1 { t.Fatal("native String boxing changed content or evaluation") }
        if _, cached := gopurs_runtime.CachedStringCodeUnits(value); fixture.text != "" && cached { t.Fatal("ordinary native String boxing acquired eager metadata") }
    }
    for _, value := range []gopurs_runtime.Value{gopurs_runtime.Int(0), {Type: gopurs_runtime.TypeBool}} {
        reads = 0
        if got, want := roundtrip(value), gopurs_runtime.Str(value.StrVal()); got != want || reads != 1 { t.Fatal("non-String normalization changed") }
    }
    value := gopurs_runtime.Str("bytes")
    value.Type = gopurs_runtime.TypeInt
    reads = 0
    if got, want := roundtrip(value), gopurs_runtime.Str(value.StrVal()); got != want || reads != 1 { t.Fatal("existing non-String byte view was not normalized") }
}
`);
    runGo(directory, ['test', '-count=1', '.'], { timeout: 30_000 });
});
