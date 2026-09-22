import { withReboxFields } from './codegen-metadata.mjs';
import assert from 'node:assert/strict';
import { spawnSync } from 'node:child_process';
import { mkdirSync, mkdtempSync, rmSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join } from 'node:path';
import test from 'node:test';
import { empty as emptyMap, insert } from '../output/Data.Map/index.js';
import { Just } from '../output/Data.Maybe/index.js';
import { ordString } from '../output/Data.Ord/index.js';
import { empty as emptySet, singleton } from '../output/Data.Set/index.js';
import { Tuple } from '../output/Data.Tuple/index.js';
import * as Ref from '../output/Effect.Ref/index.js';
import * as Adt from '../output/Gopurs.AdtExprs/index.js';
import * as CodeGen from '../output/Gopurs.CodeGen/index.js';
import * as Go from '../output/Gopurs.GoAst/index.js';
import { printGoExpr } from '../output/Gopurs.Printer/index.js';
import { runtimeGoCode } from '../output/Gopurs.Runtime/index.js';
import * as C from '../output/PureScript.Backend.Optimizer.CoreFn/index.js';
import { NeutralExpr } from '../output/PureScript.Backend.Optimizer.Semantics/index.js';
import * as S from '../output/PureScript.Backend.Optimizer.Syntax/index.js';

const map = entries => entries.reduce((result, [key, value]) => insert(ordString)(key)(value)(result), emptyMap);
const cases = [
    { name: 'Maybe', module: 'Data.Maybe', type: new C.ADT('Data.Maybe.Maybe', ['Data', 'Maybe', 'Maybe'], [C.Int.value]),
        signature: [Go.TypeValue.value, Go.TypeBool.value], fields: [['Just', 0, 0]] },
    { name: 'Either', module: 'Data.Either', type: new C.ADT('Data.Either.Either', ['Data', 'Either', 'Either'], [C.Int.value, C.Int.value]),
        signature: [Go.TypeValue.value, Go.TypeValue.value, Go.TypeBool.value], fields: [['Left', 0, 0], ['Right', 0, 1]] },
    { name: 'Tuple', module: 'Data.Tuple', type: new C.ADT('Data.Tuple.Tuple', ['Data', 'Tuple', 'Tuple'], [C.Int.value, C.Int.value]),
        signature: [Go.TypeValue.value, Go.TypeValue.value], fields: [['Tuple', 0, 0], ['Tuple', 1, 1]] },
];
const nativeType = fixture => new Go.TypeStructValue(`${fixture.module}.${fixture.name}`, fixture.signature);
const metadata = withReboxFields({
    elidedCtors: emptySet, ctorTypes: emptyMap, pointerAdtPaths: emptyMap,
    pointerAdtNodes: singleton('Data_Data_Maybe_Just'),
    pointerAdtLeaves: map([['Data_Data_Maybe_Nothing', { nodeBaseStruct: 'Data_Data_Maybe_Just', nodeCtor: 'Just' }]]),
    enumAdts: emptySet, enumCtors: emptySet, globalTypes: emptyMap, classDeclsFields: emptyMap,
    globalFunctions: map(cases.map(fixture => [`Producer.${fixture.name}`, {
        fullName: `Call_Producer_${fixture.name}`, arity: 1, fArgs: [Go.TypeBool.value], fRet: nativeType(fixture),
    }])),
});
const expr = syntax => new NeutralExpr(syntax);
const typed = (type, value) => expr(new S.Typed(type, value));
const local = (name, level) => expr(new S.Local(new Just(name), level));
const integer = value => typed(C.Int.value, expr(new S.Lit(new C.LitInt(value))));
const qualified = (module, name) => new C.Qualified(new Just(module), name);
const value = fixture => typed(fixture.type, local('result', 1));
const field = (fixture, ctor, index) => typed(C.Int.value, expr(new S.Accessor(value(fixture),
    new S.GetCtorField(qualified(fixture.module, ctor), C.SumType.value, fixture.name, ctor, `value${index}`, index))));
const tag = (fixture, ctor) => typed(C.Boolean.value, expr(new S.PrimOp(new S.Op1(
    new S.OpIsTag(qualified(fixture.module, ctor)), value(fixture)))));
const call = fixture => typed(fixture.type, expr(new S.App(
    typed(new C.Func([C.Boolean.value], fixture.type), expr(new S.Var(qualified('Producer', fixture.name)))),
    [typed(C.Boolean.value, local('input', 0))],
)));
const lambda = (types, result, body) => typed(new C.Func(types, result), expr(new S.Abs(
    types.map((_, level) => new Tuple(new Just(level === 0 ? 'input' : 'consume'), level)), body)));

for (const fixture of cases) {
    test(`native ${fixture.name} constructor fields use their actual result slots`, () => {
        const ref = Ref.new({ declarations: [], globalId: 0, reboxPairs: emptySet })();
        for (const [ctorName, index, slot] of fixture.fields) {
            const result = Adt.getField(metadata)(ref)('Consumer')({ moduleName: new Just(fixture.module), ctorName, index })({
                expr: new Go.GoVar('native'), exprType: nativeType(fixture), sourceType: fixture.type,
            });
            assert.equal(printGoExpr(result.expr), `native.V${slot}`);
            assert.deepEqual(result.exprType, Go.TypeValue.value);
        }
    });
}

function consumerCode(codegenMetadata = metadata) {
    const bindings = [];
    for (const fixture of cases) {
        const body = fixture.name === 'Maybe'
            ? typed(C.Int.value, expr(new S.Branch([new S.Pair(tag(fixture, 'Just'), field(fixture, 'Just', 0))], integer(-1))))
            : fixture.name === 'Either'
                ? typed(C.Int.value, expr(new S.Branch([new S.Pair(tag(fixture, 'Right'), field(fixture, 'Right', 0))], field(fixture, 'Left', 0))))
                : typed(C.Int.value, expr(new S.PrimOp(new S.Op2(new S.OpIntNum(S.OpAdd.value), field(fixture, 'Tuple', 0), field(fixture, 'Tuple', 1)))));
        bindings.push([`read${fixture.name}`, lambda([C.Boolean.value], C.Int.value,
            expr(new S.Let(new Just('result'), 1, call(fixture), body)))]);
        bindings.push([`return${fixture.name}`, lambda([C.Boolean.value], fixture.type, call(fixture))]);
        bindings.push([`boxed${fixture.name}`, lambda([C.Boolean.value, C.Any.value], C.Int.value,
            typed(C.Int.value, expr(new S.App(local('consume', 1), [call(fixture)]))))]);
    }
    return CodeGen.translate(codegenMetadata)({
        name: 'SumUse', bindings: bindings.map(([name, body]) => ({ recursive: false, bindings: [new Tuple(name, body)] })),
        comments: [], imports: emptySet, exports: emptySet, reExports: emptySet, dataTypes: emptyMap,
        dataDecls: [], classDecls: [], foreign: emptyMap, implementations: emptyMap, directives: emptyMap,
    });
}

test('annotated native sums survive worker calls and field reads, with boxing at dynamic boundaries', t => {
    const code = consumerCode();
    for (const fixture of cases) {
        const name = `Call_SumUse_read${fixture.name}`;
        const worker = code.slice(code.indexOf(`func ${name}(`)).split(/\nfunc /)[0];
        assert.match(worker, new RegExp(`Call_Producer_${fixture.name}\\(`));
        assert.doesNotMatch(worker, /UnsafePtr|CoerceToStruct|&Constructor_/,
            `${fixture.name} typed annotations must not rebox native worker results`);
    }
    runConsumers(t, code);
});

test('annotations retain typed pointer payloads for native Maybe and Tuple results', t => {
    const code = consumerCode(withReboxFields({
        ...metadata,
        pointerAdtPaths: map([
            ['Data.Maybe.Maybe', { ctorName: 'Just', arity: 1 }],
            ['Data.Tuple.Tuple', { ctorName: 'Tuple', arity: 2 }],
        ]),
        ctorTypes: map([
            ['Data_Maybe.Just', { vars: ['a'], fields: [new C.TypeVar('a')] }],
            ['Data_Tuple.Tuple', { vars: ['a', 'b'], fields: [new C.TypeVar('a'), new C.TypeVar('b')] }],
        ]),
    }));
    for (const name of ['Maybe', 'Tuple']) {
        const worker = code.slice(code.indexOf(`func Call_SumUse_read${name}(`)).split(/\nfunc /)[0];
        assert.match(worker, /Rebox_SumUse_\d+_\d+\(/,
            `${name} annotation must convert generic Value payloads to the expected typed pointer`);
        assert.doesNotMatch(worker, /result\w*\)?\.V\d\.IntVal/,
            `${name} field reads must use the pointer's native int64 payloads`);
    }
    runConsumers(t, code);
});

function runConsumers(t, code) {
    const directory = mkdtempSync(join(tmpdir(), 'gopurs-native-sum-results-'));
    t.after(() => rmSync(directory, { recursive: true, force: true }));
    mkdirSync(join(directory, 'purescript'));
    mkdirSync(join(directory, 'gopurs_runtime'));
    writeFileSync(join(directory, 'go.mod'), 'module gopurs/output\n\ngo 1.22\n');
    writeFileSync(join(directory, 'gopurs_runtime/runtime.go'), runtimeGoCode);
    writeFileSync(join(directory, 'purescript/SumUse.go'), code);
    writeFileSync(join(directory, 'purescript/Producer.go'), `package purescript
import "gopurs/output/gopurs_runtime"
type Constructor_Data_Maybe_Just[A any] struct { Rc uint32; V0 A }
type Constructor_Data_Either_Left[A, B any] struct { Rc uint32; V0 A }
type Constructor_Data_Either_Right[A, B any] struct { Rc uint32; V0 B }
type Constructor_Data_Tuple_Tuple[A, B any] struct { Rc uint32; V0 A; V1 B }
var MaybeCalls, EitherCalls, TupleCalls int
func Call_Producer_Maybe(input bool) struct { V0 gopurs_runtime.Value; V1 bool } {
    MaybeCalls++
    return struct { V0 gopurs_runtime.Value; V1 bool }{gopurs_runtime.Int(42), input}
}
func Call_Producer_Either(input bool) struct { V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool } {
    EitherCalls++
    return struct { V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool }{gopurs_runtime.Int(21), gopurs_runtime.Int(42), input}
}
func Call_Producer_Tuple(input bool) struct { V0 gopurs_runtime.Value; V1 gopurs_runtime.Value } {
    TupleCalls++
    return struct { V0 gopurs_runtime.Value; V1 gopurs_runtime.Value }{gopurs_runtime.Int(21), gopurs_runtime.Int(42)}
}
`);
    writeFileSync(join(directory, 'main.go'), `package main
import (
    "fmt"
    "gopurs/output/purescript"
    "gopurs/output/gopurs_runtime"
)
func main() {
    maybe := gopurs_runtime.Func(func(v gopurs_runtime.Value) gopurs_runtime.Value {
        if v.Type != 9 || v.IntVal != 930809136 { panic("Maybe tag lost") }
        if v.UnsafePtr == nil { return gopurs_runtime.Int(-1) }
        return (*purescript.Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v.UnsafePtr).V0
    })
    either := gopurs_runtime.Func(func(v gopurs_runtime.Value) gopurs_runtime.Value {
        if v.Type != 9 { panic("Either tag lost") }
        switch v.IntVal {
        case 3711209382: return (*purescript.Constructor_Data_Either_Left[gopurs_runtime.Value,gopurs_runtime.Value])(v.UnsafePtr).V0
        case 2465973597: return (*purescript.Constructor_Data_Either_Right[gopurs_runtime.Value,gopurs_runtime.Value])(v.UnsafePtr).V0
        default: panic("unexpected Either constructor")
        }
    })
    tuple := gopurs_runtime.Func(func(v gopurs_runtime.Value) gopurs_runtime.Value {
        if v.Type != 9 { panic("Tuple tag lost") }
        t := (*purescript.Constructor_Data_Tuple_Tuple[gopurs_runtime.Value,gopurs_runtime.Value])(v.UnsafePtr)
        return gopurs_runtime.Int(t.V0.IntVal+t.V1.IntVal)
    })
    for _, input := range []bool{false,true} {
        fmt.Printf("read %d %d %d\\n", purescript.Call_SumUse_readMaybe(input), purescript.Call_SumUse_readEither(input), purescript.Call_SumUse_readTuple(input))
        fmt.Printf("dynamic %d %d %d\\n", purescript.Call_SumUse_boxedMaybe(input,maybe), purescript.Call_SumUse_boxedEither(input,either), purescript.Call_SumUse_boxedTuple(input,tuple))
        boxed := gopurs_runtime.Bool(input)
        fmt.Printf("wrapper %d %d %d\\n",
            gopurs_runtime.Apply(maybe,gopurs_runtime.Apply(purescript.Get_SumUse_returnMaybe(),boxed)).IntVal,
            gopurs_runtime.Apply(either,gopurs_runtime.Apply(purescript.Get_SumUse_returnEither(),boxed)).IntVal,
            gopurs_runtime.Apply(tuple,gopurs_runtime.Apply(purescript.Get_SumUse_returnTuple(),boxed)).IntVal)
    }
    fmt.Printf("calls %d %d %d\\n",purescript.MaybeCalls,purescript.EitherCalls,purescript.TupleCalls)
}
`);
    const result = spawnSync('go', ['run', '.'], {
        cwd: directory, encoding: 'utf8', timeout: 30_000, env: { ...process.env, GOWORK: 'off' },
    });
    assert.ifError(result.error);
    assert.equal(result.status, 0, result.stdout + result.stderr);
    assert.equal(result.stdout,
        'read -1 21 63\ndynamic -1 21 63\nwrapper -1 21 63\n'
        + 'read 42 42 63\ndynamic 42 42 63\nwrapper 42 42 63\ncalls 6 6 6\n');
}
