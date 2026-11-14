import assert from 'node:assert/strict';
import { spawnSync } from 'node:child_process';
import { copyFileSync, mkdirSync, mkdtempSync, rmSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join } from 'node:path';
import test from 'node:test';
import { Just, Nothing } from '../output/Data.Maybe/index.js';
import { Tuple } from '../output/Data.Tuple/index.js';
import * as C from '../output/PureScript.Backend.Optimizer.CoreFn/index.js';
import * as Structured from '../output/Gopurs.FfiBridge.Structured/index.js';
import { printTypeNode, resolveNewtype } from '../output/Gopurs.FfiBridge.TypeSupport/index.js';
import { generateFfiBridge, ffiProblems, ffiProblemsWithTypes, ffiFunctionInfos, ffiValueWorkers } from '../output/Gopurs.FfiBridge/index.js';
import * as Map from '../output/Data.Map/index.js';
import { prepareFfi } from '../output/Gopurs.FfiSupport/index.js';

const record = entries => new C.Record(new C.Row(entries.map(([name, type]) => new Tuple(name, type)), Nothing.value));
const nullable = type => new C.ADT('Data.Nullable.Nullable', ['Data', 'Nullable', 'Nullable'], [type]);
const foreign = new C.ADT('Foreign.Foreign', ['Foreign', 'Foreign'], []);
const child = record([['label', C.String.value], ['active', C.Boolean.value]]);
const payload = record([
    ['children', new C.Array(child)], ['count', C.Int.value], ['child', child],
    ['ratio', C.Number.value], ['maybe', nullable(child)],
]);
const source = `package fixture
type Child struct { Active bool; Label string }
type Payload struct {
 Count int64
 Ratio float64
 Child Child
 Children []Child
 Maybe *Child
}
func RoundTrip(value Payload) Payload { value.Count++; return value }
func Children(value []Child) []Child { return value }
func Nullable(value *Child) *Child { return value }
func Empty(value struct{}) struct{} { return value }
type DynamicAny = any
type Dynamic struct { Value DynamicAny }
func DynamicValue(value Dynamic) Dynamic { values := value.Value.(map[string]any); if _, ok := values["count"].(int64); !ok { panic("Foreign field leaked Value") }; return value }
func MutateNullable(value *Child) *Child {
 if value != nil { value.Label += "!"; value.Active = true }
 return value
}
`;
const parse = source => prepareFfi({ moduleName: 'Structured', path: '/fixture/structured.go' })('Structured_')(source)();
const unbox = go => tast => value => tast === foreign ? `gopurs_runtime.ForeignValue(${value})` : `gopurs_runtime.Unbox[${printTypeNode(go)}](${value})`;
const box = _ => _ => value => `gopurs_runtime.Box(${value})`;

// The helper's emitted expressions are exercised separately from the wrapper
// dispatcher so a dispatch regression can be distinguished from layout bugs.
test('closed records cross ordinary Go structs, arrays and Nullable pointers without reflection', t => {
    const parsed = parse(source);
    const contracts = [payload, new C.Array(child), nullable(child), record([]), record([['value', foreign]])];
    let bridges = '';
    for (const [i, tast] of contracts.entries()) {
        const declaration = parsed.decls[i];
        assert.ok(Structured.needsConversion(declaration.args[0])(tast));
        assert.equal(Structured.problem(declaration.args[0])(tast), Nothing.value);
        const from = Structured.fromValue(unbox)(declaration.args[0])(tast)('value');
        const to = Structured.toValue(box)(declaration.results[0])(tast)('result');
        bridges += `func Bridge${i}(value gopurs_runtime.Value) gopurs_runtime.Value { native := ${from}; result := ${declaration.name}(native); return ${to} }\n`;
    }
    const foreigns = ['roundTrip', 'children', 'nullable', 'empty', 'dynamicValue'].map((name, i) => new Tuple(name, new Just(new C.Func([contracts[i]], contracts[i]))));
    foreigns.push(new Tuple('mutateNullable', new Just(new C.Func([nullable(child)], nullable(child)))));
    assert.deepEqual(ffiProblems('Structured')(parsed.decls)(foreigns), []);
    const declarations = [{ name: 'Wrapped', vars: ['a'], constructors: [{ name: 'Wrapped', fields: [new C.TypeVar('a')] }] }];
    const wrapped = new C.ADT('Structured.Wrapped', ['Structured', 'Wrapped'], [payload]);
    assert.deepEqual(resolveNewtype(declarations)(wrapped), payload);
    const genericForeigns = [new Tuple('roundTrip', new Just(new C.Func([wrapped], wrapped)))];
    assert.deepEqual(ffiProblemsWithTypes('Structured')(declarations)(parsed.decls)(genericForeigns), []);
    bridges += generateFfiBridge('Structured')(declarations)(parsed.decls)([genericForeigns[0], ...foreigns.slice(1)]);
    assert.doesNotMatch(bridges, /reflect\.|json\.|RecordToMap/);
    const directory = mkdtempSync(join(tmpdir(), 'gopurs-ffi-structured-'));
    t.after(() => rmSync(directory, { recursive: true, force: true }));
    writeFileSync(join(directory, 'go.mod'), 'module gopurs/output\ngo 1.27.0\n');
    mkdirSync(join(directory, 'ordinary'));
    writeFileSync(join(directory, 'ordinary/fixture.go'), source);
    writeFileSync(join(directory, 'ordinary/fixture_test.go'), `package fixture
import "testing"
func TestOrdinary(t *testing.T) { if RoundTrip(Payload{Count: 41}).Count != 42 || Nullable(nil) != nil { t.Fatal("ordinary Go behavior") } }
`);
    mkdirSync(join(directory, 'gopurs_runtime'));
    copyFileSync(new URL('../runtime/runtime.go', import.meta.url), join(directory, 'gopurs_runtime/runtime.go'));
    writeFileSync(join(directory, 'bridge.go'), 'package fixture\nimport "gopurs/output/gopurs_runtime"\n' + parsed.content + '\n' + bridges);
    writeFileSync(join(directory, 'bridge_test.go'), `package fixture
import (
 "math"
 "testing"
 rt "gopurs/output/gopurs_runtime"
)
func child(label string, active bool) rt.Value { return rt.RecordDict2("label", "active", rt.Str(label), rt.Box(active)) }
func TestRoundTrip(t *testing.T) {
 originalChild := child("", false)
 value := rt.RecordDict5("count", "ratio", "child", "children", "maybe", rt.Int(41), rt.Box(math.Copysign(0, -1)), originalChild, rt.Array([]rt.Value{child("array", true)}), rt.Any(nil))
 result := Bridge0(value)
 viaGenerated := rt.Apply(_Gopurs_Structured_RoundTrip, value)
 if rt.Unbox[int64](rt.RecordGet(viaGenerated, "count")) != 42 { t.Fatal("FFI wrapper dispatch") }
 if rt.Apply(_Gopurs_Structured_Nullable, rt.Any(nil)).PtrVal() != nil { t.Fatal("generated nullable wrapper") }
 if rt.ArrayLength(rt.Apply(_Gopurs_Structured_Children, rt.Array(nil))) != 0 { t.Fatal("generated array wrapper") }
 if len(rt.RecordToMap(rt.Apply(_Gopurs_Structured_Empty, rt.RecordDict0()))) != 0 { t.Fatal("generated empty wrapper") }
 if rt.Unbox[int64](rt.RecordGet(result, "count")) != 42 { t.Fatal("count") }
 if !math.Signbit(rt.Unbox[float64](rt.RecordGet(result, "ratio"))) { t.Fatal("negative zero") }
 nested := rt.RecordGet(result, "child")
 if rt.Unbox[string](rt.RecordGet(nested, "label")) != "" || rt.Unbox[bool](rt.RecordGet(nested, "active")) { t.Fatal("empty string/false") }
 children := rt.RecordGet(result, "children")
 if rt.ArrayLength(children) != 1 || rt.Unbox[string](rt.RecordGet(rt.ArrayAccess(children, 0), "label")) != "array" { t.Fatal("array of structs") }
 if rt.RecordGet(result, "maybe").PtrVal() != nil { t.Fatal("nullable field") }
 if rt.Unbox[int64](rt.RecordGet(value, "count")) != 41 { t.Fatal("input record mutated") }
 for _, input := range []rt.Value{rt.Array(nil), rt.Array([]rt.Value{})} { if rt.ArrayLength(Bridge1(input)) != 0 { t.Fatal("empty array") } }
 for _, input := range []rt.Value{rt.Any(nil), rt.Any((*Child)(nil))} { result := Bridge2(input); if result.Type != rt.TypeAny || result.PtrVal() != nil { t.Fatal("nil pointer") } }
 if rt.Unbox[string](rt.RecordGet(Bridge2(child("present", true)), "label")) != "present" { t.Fatal("nullable record value") }
 mutationInput := child("before", false)
 mutationResult := rt.Apply(_Gopurs_Structured_MutateNullable, mutationInput)
 if rt.Unbox[string](rt.RecordGet(mutationResult, "label")) != "before!" || !rt.Unbox[bool](rt.RecordGet(mutationResult, "active")) {
  t.Fatal("returning the native input pointer lost record mutations")
 }
 if rt.Unbox[string](rt.RecordGet(mutationInput, "label")) != "before" || rt.Unbox[bool](rt.RecordGet(mutationInput, "active")) {
  t.Fatal("native pointer mutation changed the original PureScript record")
 }
 if rt.Apply(_Gopurs_Structured_MutateNullable, rt.Any(nil)).PtrVal() != nil { t.Fatal("mutating nullable wrapper lost nil") }
 dynamic := rt.RecordDict1("value", rt.RecordDict2("count", "nullable", rt.Int(42), rt.Any(nil)))
 resultDynamic := rt.Apply(_Gopurs_Structured_DynamicValue, dynamic)
 fields := rt.ForeignValue(rt.RecordGet(resultDynamic, "value")).(map[string]any)
 if fields["count"] != int64(42) || fields["nullable"] != nil { t.Fatal("Foreign dynamic field", fields) }
 if len(rt.RecordToMap(Bridge3(rt.RecordDict0()))) != 0 { t.Fatal("empty record") }
}
`);
    const result = spawnSync('go', ['test', '-race', '-count=1', './...'], { cwd: directory, encoding: 'utf8', timeout: 60000, env: { ...process.env, GOWORK: 'off' } });
    assert.ifError(result.error);
    assert.equal(result.status, 0, result.stdout + result.stderr + '\n' + bridges);
});

test('struct shape incompatibilities are diagnosed before Go code generation', () => {
    const cases = [
        ['type R struct{ Name string }; func Read(R) {}', record([['name', C.Int.value]]), /Field name:.*incompatible/],
        ['type R struct{ Count any }; func Read(R) {}', record([['count', C.Int.value]]), /Field count:.*requires Foreign.Foreign/],
        ['type R struct{ Count []any }; func Read(R) {}', record([['count', new C.Array(C.Int.value)]]), /Field count:.*requires Foreign.Foreign/],
        ['type R struct{ Count map[string]any }; func Read(R) {}', record([['count', record([['value', C.Int.value]])]]), /Map fields.*explicit conversion/],
        ['type R struct{ Name string }; func Read(R) {}', record([['count', C.Int.value]]), /Missing Go struct field Count/],
        ['type R struct{ Name string; Count int }; func Read(R) {}', record([['name', C.String.value]]), /Unexpected Go struct field Count/],
        ['type R struct{ name string }; func Read(R) {}', record([['name', C.String.value]]), /exported and non-embedded/],
        ['type Embedded struct{ Name string }; type R struct{ Embedded }; func Read(R) {}', record([['embedded', record([['name', C.String.value]])]]), /exported and non-embedded/],
        ['type R struct{ Name string }; func Read(R) {}', new C.Record(new C.Row([new Tuple('name', C.String.value)], new Just(new C.TypeVar('r')))), /Open PureScript record/],
        ['type R struct{ Name string }; func Read(*R) {}', record([['name', C.String.value]]), /non-nullable/],
        ['type R struct{ X string }; func Read(R) {}', record([['x', C.String.value], ['X', C.String.value]]), /collide/],
        ['import sdk "example/sdk"; func Read(sdk.Record) {}', record([['name', C.String.value]]), /locally declared Go struct/],
        ['type R struct{ Count float64 }; func Read(R) {}', record([['count', C.Int.value]]), /Field count:.*incompatible/],
    ];
    for (const [source, tast, expected] of cases) {
        const declaration = parse(source).decls[0];
        const error = Structured.problem(declaration.args[0])(tast);
        assert.ok(error instanceof Just, source);
        assert.match(error.value0, expected);
        const errors = ffiProblems('Structured')([declaration])([new Tuple('read', new Just(new C.Func([tast], C.Unit.value)))]);
        assert.equal(errors.length, 1);
        assert.match(errors[0], /read at .*Go.*PureScript/);
        assert.match(errors[0], expected);
    }
    const open = new C.Record(new C.Row([], new Just(new C.TypeVar('r'))));
    for (const type of ['any', 'map[string]any', 'gopurs_runtime.Value']) {
        const legacy = parse(`func Read(value ${type}) {}`).decls[0].args[0];
        assert.equal(Structured.problem(legacy)(open), Nothing.value, `legacy ${type} must retain open-row support`);
    }
    // Serialization tags belong to the application, not the FFI field mapping.
    const tagged = parse('type R struct { Name string `json:"business_name"` }; func Read(R) {}').decls[0].args[0];
    assert.equal(Structured.problem(tagged)(record([['name', C.String.value]])), Nothing.value);
    const opaque = new C.ADT('Fixture.Handle', ['Fixture', 'Handle'], []);
    assert.equal(Structured.needsConversion(tagged)(opaque), false);
    assert.equal(Structured.problem(tagged)(opaque), Nothing.value);
});

test('recursive elided declarations terminate and type parameters substitute capture-safely', () => {
    const recursive = new C.ADT('Recursive', ['Recursive'], []);
    const declarations = [{ name: 'Recursive', vars: [], constructors: [{ name: 'Recursive', fields: [recursive] }] }];
    assert.deepEqual(resolveNewtype(declarations)(recursive), recursive);
    const wrapped = [{ name: 'Box', vars: ['a'], constructors: [{ name: 'Box', fields: [new C.TypeVar('a')] }] }];
    assert.deepEqual(resolveNewtype(wrapped)(new C.ADT('Box', ['Box'], [C.String.value])), C.String.value);
});


test('record array results keep field conversion when a callback would otherwise introduce a native worker', () => {
    const parsed = parse('type Payload struct { Count int64 }; func Read(callback func(int64) int64) []Payload { return []Payload{{Count: callback(41)}} }');
    const type = new C.Func([new C.Func([C.Int.value], C.Int.value)], new C.Array(record([['count', C.Int.value]])));
    const foreigns = [new Tuple('read', new Just(type))];
    assert.equal(Map.size(ffiFunctionInfos('Structured')(foreigns)(parsed.decls)), 0);
    assert.equal(ffiValueWorkers('Structured')([])(foreigns)(parsed.decls), '');
    assert.match(generateFfiBridge('Structured')([])(parsed.decls)(foreigns), /RecordDict1/);
});
