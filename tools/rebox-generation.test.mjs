import assert from 'node:assert/strict';
import { spawnSync } from 'node:child_process';
import { mkdirSync, mkdtempSync, rmSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join } from 'node:path';
import test from 'node:test';
import { empty as emptyMap, insert } from '../output/Data.Map/index.js';
import { ordString } from '../output/Data.Ord/index.js';
import { empty as emptySet, toUnfoldable } from '../output/Data.Set/index.js';
import { unfoldableArray } from '../output/Data.Unfoldable/index.js';
import * as Ref from '../output/Effect.Ref/index.js';
import * as Go from '../output/Gopurs.GoAst/index.js';
import { coerceGoExpr, generateReboxFunctions, unboxGoExpr } from '../output/Gopurs.GoConversions/index.js';
import { printGoDecl, printGoExpr } from '../output/Gopurs.Printer/index.js';
import { runtimeGoCode } from '../output/Gopurs.Runtime/index.js';
import * as C from '../output/PureScript.Backend.Optimizer.CoreFn/index.js';
import { withReboxFields } from './codegen-metadata.mjs';

const map = entries => entries.reduce((result, [key, value]) => insert(ordString)(key)(value)(result), emptyMap);
const a = new C.TypeVar('a');
const adt = name => new C.ADT(`Probe.${name}`, ['Probe', name], [a]);
const metadata = withReboxFields({
    ctorTypes: map([
        ['Probe.Root', { vars: ['a'], fields: [adt('Branch'), adt('Branch'), adt('Phantom')] }],
        ['Probe.Branch', { vars: ['a'], fields: [a, adt('Branch')] }],
        ['Probe.Phantom', { vars: ['a'], fields: [C.Int.value] }],
    ]),
    pointerAdtPaths: map(['Root', 'Branch', 'Phantom'].map(name => [`Probe.${name}`, { ctorName: name, arity: 1 }])),
    pointerAdtNodes: emptySet, pointerAdtLeaves: emptyMap, elidedCtors: emptySet,
    enumAdts: emptySet, enumCtors: emptySet, classDeclsFields: emptyMap,
    globalTypes: emptyMap, globalFunctions: emptyMap,
});
const pointer = (name, argument) => Go.structPointer({
    baseStructName: `Data_Probe_${name}`, fullName: `Probe.${name}`, structName: `Constructor_Probe_${name}`,
})([argument]);
const boxedRoot = pointer('Root', Go.TypeValue.value);
const nativeRoot = pointer('Root', Go.TypeInt64.value);
const newState = () => Ref.new({ declarations: [], globalId: 0, reboxPairs: emptySet })();
const pairs = state => toUnfoldable(unfoldableArray)(Ref.read(state)().reboxPairs);
const convert = state => expr => from => to => coerceGoExpr(state)('Consumer')(expr)(from)(to);
const generate = state => generateReboxFunctions(metadata)(state)('Consumer')();

test('Rebox emits the directed transitive closure once, including recursive and phantom dependencies', () => {
    const forward = [boxedRoot, nativeRoot];
    const backward = [nativeRoot, boxedRoot];
    const outputs = [];
    for (const requests of [[forward, backward], [backward, forward]]) {
        const state = newState();
        for (const [from, to] of [...requests, ...requests]) {
            convert(state)(new Go.GoVar('root'))(from)(to);
        }
        assert.equal(pairs(state).length, 2, 'conversion requests are directed and deduplicated');
        const declarations = generate(state);
        assert.equal(declarations.length, 6, 'root helpers must discover both nested layouts in both directions');
        assert.equal(pairs(state).length, 6, 'recursive branch fields must converge to the existing pair');
        const names = declarations.map(declaration => declaration.value0.name);
        assert.equal(new Set(names).size, names.length);
        assert.deepEqual(names, [...names].sort(), 'emission order is determined by helper names');
        const text = declarations.map(printGoDecl);
        assert.deepEqual(generate(state).map(printGoDecl), text, 'emitting again must not discover further helpers');
        assert.equal(pairs(state).length, 6);
        outputs.push(text);
    }
    assert.deepEqual(outputs[0], outputs[1], 'request insertion order must not affect generated Go');
});

test('generated Rebox helpers preserve nested payloads, nil, phantom identity and single evaluation', t => {
    const state = newState();
    const convertRoot = convert(state)(new Go.GoCall(new Go.GoVar('source'), []))(boxedRoot)(nativeRoot);
    const roundTrip = convert(state)(new Go.GoVar('native'))(nativeRoot)(boxedRoot);
    const nilRoot = convert(state)(Go.rawGo('(*Constructor_Probe_Root[gopurs_runtime.Value])(nil)'))(boxedRoot)(nativeRoot);
    const arrayConversion = unboxGoExpr(state)('Consumer')(new Go.GoCall(new Go.GoVar('arraySource'), []))
        (new Go.TypeNativeArray(pointer('Branch', Go.TypeInt64.value)))
        (new Go.TypeNativeArray(pointer('Branch', Go.TypeValue.value)));
    const declarations = generate(state).map(printGoDecl);
    assert.equal(declarations.length, 6);
    const directory = mkdtempSync(join(tmpdir(), 'gopurs-rebox-generation-'));
    t.after(() => rmSync(directory, { recursive: true, force: true }));
    mkdirSync(join(directory, 'gopurs_runtime'));
    writeFileSync(join(directory, 'go.mod'), 'module gopurs/output\n\ngo 1.22\n');
    writeFileSync(join(directory, 'gopurs_runtime/runtime.go'), runtimeGoCode);
    writeFileSync(join(directory, 'main.go'), `package main
import ("fmt"; "unsafe"; "gopurs/output/gopurs_runtime")
type Constructor_Probe_Branch[A any] struct { Rc uint32; V0 A; V1 *Constructor_Probe_Branch[A] }
type Constructor_Probe_Phantom[A any] struct { Rc uint32; V0 int64 }
type Constructor_Probe_Root[A any] struct {
    Rc uint32
    V0 *Constructor_Probe_Branch[A]
    V1 *Constructor_Probe_Branch[A]
    V2 *Constructor_Probe_Phantom[A]
}
${declarations.join('\n')}
var input *Constructor_Probe_Root[gopurs_runtime.Value]
var nativeArray []*Constructor_Probe_Branch[int64]
var reads, arrayReads int
func source() *Constructor_Probe_Root[gopurs_runtime.Value] { reads++; return input }
func arraySource() []*Constructor_Probe_Branch[int64] { arrayReads++; return nativeArray }
func main() {
    input = &Constructor_Probe_Root[gopurs_runtime.Value]{Rc: 1,
        V0: &Constructor_Probe_Branch[gopurs_runtime.Value]{Rc: 3, V0: gopurs_runtime.Int(7),
            V1: &Constructor_Probe_Branch[gopurs_runtime.Value]{Rc: 4, V0: gopurs_runtime.Int(-4)}},
        V2: &Constructor_Probe_Phantom[gopurs_runtime.Value]{Rc: 5, V0: 42},
    }
    native := ${printGoExpr(convertRoot)}
    boxed := ${printGoExpr(roundTrip)}
    if reads != 1 || native.V0.V0 != 7 || native.V0.V1.V0 != -4 { panic("evaluation or native payload changed") }
    if native.V1 != nil || native.V0.V1.V1 != nil || ${printGoExpr(nilRoot)} != nil { panic("nil changed") }
    if unsafe.Pointer(native.V2) != unsafe.Pointer(input.V2) || native.V2.V0 != 42 { panic("phantom layout must retain pointer identity") }
    if boxed.V0.V0.IntVal != 7 || boxed.V0.V1.V0.IntVal != -4 || boxed.V2 != input.V2 { panic("roundtrip changed payloads") }
    if unsafe.Pointer(native.V0) == unsafe.Pointer(input.V0) || input.V0.Rc != 3 || input.V0.V0.IntVal != 7 { panic("typed conversion must preserve its source") }
    nativeArray = []*Constructor_Probe_Branch[int64]{native.V0, nil}
    convertedArray := ${printGoExpr(arrayConversion)}
    if arrayReads != 1 || len(convertedArray) != 2 || convertedArray[0].V0.IntVal != 7 || convertedArray[1] != nil { panic("native array conversion changed") }
    input = &Constructor_Probe_Root[gopurs_runtime.Value]{}
    empty := ${printGoExpr(convertRoot)}
    if empty.V0 != nil || empty.V1 != nil || empty.V2 != nil { panic("nested nil changed") }
    fmt.Println(native.V0.V0, native.V0.V1.V0, boxed.V0.V0.IntVal, reads, arrayReads)
}
`);
    const result = spawnSync('go', ['run', '.'], {
        cwd: directory, encoding: 'utf8', timeout: 30_000,
        env: { ...process.env, GOWORK: 'off' },
    });
    assert.ifError(result.error);
    assert.equal(result.status, 0, result.stdout + result.stderr);
    assert.equal(result.stdout, '7 -4 7 2 1\n');
});

test('missing Rebox field metadata retains its diagnostic and omits the helper', () => {
    const state = newState();
    const call = convert(state)(new Go.GoVar('unknown'))(pointer('Unknown', Go.TypeValue.value))(pointer('Unknown', Go.TypeInt64.value));
    assert.match(printGoExpr(call), /^Rebox_Consumer_\d+_\d+\(unknown\)$/);
    const messages = [];
    const originalLog = console.log;
    let declarations;
    try {
        console.log = message => messages.push(message);
        declarations = generate(state);
    } finally {
        console.log = originalLog;
    }
    assert.deepEqual(declarations, []);
    assert.deepEqual(messages, ['ERROR: Rebox missing! b1=Data_Probe_Unknown keysCtor: Probe.Branch, Probe.Phantom, Probe.Root']);
    assert.equal(pairs(state).length, 1, 'missing metadata does not erase the conversion request');
});
