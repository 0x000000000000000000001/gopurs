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
import * as CodeGen from '../output/Gopurs.CodeGen/index.js';
import * as Go from '../output/Gopurs.GoAst/index.js';
import { runtimeGoCode } from '../output/Gopurs.Runtime/index.js';
import * as C from '../output/PureScript.Backend.Optimizer.CoreFn/index.js';
import { NeutralExpr } from '../output/PureScript.Backend.Optimizer.Semantics/index.js';
import * as S from '../output/PureScript.Backend.Optimizer.Syntax/index.js';

const map = entries => entries.reduce((result, [key, value]) =>
    insert(ordString)(key)(value)(result), emptyMap);
const expr = syntax => new NeutralExpr(syntax);
const typed = (type, value) => expr(new S.Typed(type, value));
const bool = C.Boolean.value;
const maybe = new C.ADT('Data.Maybe.Maybe', ['Data', 'Maybe', 'Maybe'], [C.Int.value]);
const either = new C.ADT('Data.Either.Either', ['Data', 'Either', 'Either'], [C.Int.value, C.Int.value]);
const tuple = new C.ADT('Data.Tuple.Tuple', ['Data', 'Tuple', 'Tuple'], [C.Int.value, C.Int.value]);
const unary = new C.Func([bool], bool);

test('constructor tests evaluate imported native Maybe, Either and Tuple workers once', t => {
    const metadata = withReboxFields({
        elidedCtors: emptySet, ctorTypes: emptyMap, pointerAdtPaths: emptyMap,
        pointerAdtNodes: singleton('Data_Data_Maybe_Just'),
        pointerAdtLeaves: map([['Data_Data_Maybe_Nothing', {
            nodeBaseStruct: 'Data_Data_Maybe_Just', nodeCtor: 'Just',
        }]]),
        enumAdts: emptySet, enumCtors: emptySet, globalTypes: emptyMap,
        classDeclsFields: emptyMap,
        globalFunctions: map([
            ['Producer.maybe', { fullName: 'Call_Producer_maybe', arity: 1,
                fArgs: [Go.TypeBool.value],
                fRet: new Go.TypeStructValue('Data.Maybe.Maybe', [Go.TypeValue.value, Go.TypeBool.value]) }],
            ['Producer.either', { fullName: 'Call_Producer_either', arity: 1,
                fArgs: [Go.TypeBool.value],
                fRet: new Go.TypeStructValue('Data.Either.Either',
                    [Go.TypeValue.value, Go.TypeValue.value, Go.TypeBool.value]) }],
            ['Producer.tuple', { fullName: 'Call_Producer_tuple', arity: 1,
                fArgs: [Go.TypeBool.value],
                fRet: new Go.TypeStructValue('Data.Tuple.Tuple',
                    [Go.TypeValue.value, Go.TypeValue.value]) }],
        ]),
    });
    const bindings = [
        ['isJust', 'maybe', maybe, 'Data.Maybe', 'Just'],
        ['isNothing', 'maybe', maybe, 'Data.Maybe', 'Nothing'],
        ['isLeft', 'either', either, 'Data.Either', 'Left'],
        ['isRight', 'either', either, 'Data.Either', 'Right'],
        ['isTuple', 'tuple', tuple, 'Data.Tuple', 'Tuple'],
    ].map(([name, worker, resultType, module, tag]) => {
        const parameter = typed(bool, expr(new S.Local(new Just('input'), 0)));
        const head = typed(new C.Func([bool], resultType),
            expr(new S.Var(new C.Qualified(new Just('Producer'), worker))));
        // The imported worker registry supplies the emitted native return type.
        const call = expr(new S.App(head, [parameter]));
        const check = typed(bool, expr(new S.PrimOp(new S.Op1(
            new S.OpIsTag(new C.Qualified(new Just(module), tag)), call))));
        return { recursive: false, bindings: [new Tuple(name, typed(unary,
            expr(new S.Abs([new Tuple(new Just('input'), 0)], check))))] };
    });
    const code = CodeGen.translate(metadata)({
        name: 'Consumer', bindings, comments: [], imports: emptySet,
        exports: emptySet, reExports: emptySet, dataTypes: emptyMap,
        dataDecls: [], classDecls: [], foreign: emptyMap,
        implementations: emptyMap, directives: emptyMap,
    });
    const directory = mkdtempSync(join(tmpdir(), 'gopurs-native-constructor-tags-'));
    t.after(() => rmSync(directory, { recursive: true, force: true }));
    mkdirSync(join(directory, 'purescript'));
    mkdirSync(join(directory, 'gopurs_runtime'));
    writeFileSync(join(directory, 'go.mod'), 'module gopurs/output\n\ngo 1.22\n');
    writeFileSync(join(directory, 'gopurs_runtime/runtime.go'), runtimeGoCode);
    writeFileSync(join(directory, 'purescript/Consumer.go'), code);
    writeFileSync(join(directory, 'purescript/Producer.go'), `package purescript
import "gopurs/output/gopurs_runtime"
type Constructor_Data_Maybe_Just[A any] struct { Rc uint32; V0 A }
type Constructor_Data_Either_Left[A, B any] struct { Rc uint32; V0 A }
type Constructor_Data_Either_Right[A, B any] struct { Rc uint32; V0 B }
type Constructor_Data_Tuple_Tuple[A, B any] struct { Rc uint32; V0 A; V1 B }
var MaybeCalls, EitherCalls, TupleCalls int
func Call_Producer_maybe(input bool) struct { V0 gopurs_runtime.Value; V1 bool } {
    MaybeCalls++
    return struct { V0 gopurs_runtime.Value; V1 bool }{gopurs_runtime.Int(42), input}
}
func Call_Producer_either(input bool) struct { V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool } {
    EitherCalls++
    return struct { V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool }{
        gopurs_runtime.Int(21), gopurs_runtime.Int(42), input,
    }
}
func Call_Producer_tuple(input bool) struct { V0 gopurs_runtime.Value; V1 gopurs_runtime.Value } {
    TupleCalls++
    return struct { V0 gopurs_runtime.Value; V1 gopurs_runtime.Value }{
        gopurs_runtime.Int(21), gopurs_runtime.Int(42),
    }
}
`);
    writeFileSync(join(directory, 'main.go'), `package main
import (
    "fmt"
    "gopurs/output/purescript"
)
func main() {
    for _, input := range []bool{false, true} {
        fmt.Printf("%t %t %t %t %t\\n", purescript.Call_Consumer_isJust(input),
            purescript.Call_Consumer_isNothing(input), purescript.Call_Consumer_isLeft(input),
            purescript.Call_Consumer_isRight(input), purescript.Call_Consumer_isTuple(input))
    }
    fmt.Printf("calls %d %d %d", purescript.MaybeCalls, purescript.EitherCalls, purescript.TupleCalls)
}
`);
    const result = spawnSync('go', ['run', '.'], {
        cwd: directory, encoding: 'utf8', timeout: 30_000,
        env: { ...process.env, GOWORK: 'off' },
    });
    assert.ifError(result.error);
    assert.equal(result.status, 0, result.stderr);
    assert.equal(result.stdout, 'false true true false true\ntrue false false true true\ncalls 4 4 2');
});
