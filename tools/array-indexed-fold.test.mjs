import assert from 'node:assert/strict';
import { spawnSync } from 'node:child_process';
import { mkdirSync, mkdtempSync, rmSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join } from 'node:path';
import test from 'node:test';
import { empty as emptyMap } from '../output/Data.Map/index.js';
import { Just, Nothing } from '../output/Data.Maybe/index.js';
import { empty as emptySet } from '../output/Data.Set/index.js';
import { Tuple } from '../output/Data.Tuple/index.js';
import * as CodeGen from '../output/Gopurs.CodeGen/index.js';
import { emit, integerCallback, readOnlyScalar } from '../output/Gopurs.ArrayIndexedFold/index.js';
import { runtimeGoCode } from '../output/Gopurs.Runtime.Source/index.js';
import { TcoExpr } from '../output/PureScript.Backend.Optimizer.Codegen.Tco/index.js';
import * as C from '../output/PureScript.Backend.Optimizer.CoreFn/index.js';
import { NeutralExpr } from '../output/PureScript.Backend.Optimizer.Semantics/index.js';
import * as S from '../output/PureScript.Backend.Optimizer.Syntax/index.js';
import { withReboxFields } from './codegen-metadata.mjs';

const e = syntax => new TcoExpr(null, syntax);
const typed = (type, value) => e(new S.Typed(type, value));
const int = value => e(new S.Lit(new C.LitInt(value)));
const bool = value => e(new S.Lit(new C.LitBoolean(value)));
const local = (level, type) => {
    const value = e(new S.Local(new Just('x'), level));
    return type ? typed(type, value) : value;
};
const ref = (module, name) => e(new S.Var(new C.Qualified(new Just(module), name)));
const binary = (op, left, right) => e(new S.PrimOp(new S.Op2(op, left, right)));
const num = (op, a, b) => binary(new S.OpIntNum(op), a, b);
const choose = (condition, yes, no) => e(new S.Branch([new S.Pair(condition, yes)], no));
const abs = (levels, body) => e(new S.Abs(levels.map(level => new Tuple(new Just('x'), level)), body));
const callbackType = new C.Func([C.Int.value, C.Int.value, C.Int.value], C.Int.value);
const callback = body => typed(callbackType, abs([0, 1, 2], body));
const suffixCallback = (type = new C.Func([C.Int.value, C.Int.value], C.Int.value)) => abs([0], typed(type, abs([1, 2], local(1))));

test('indexed callback types require a full TAST signature or the exact typed suffix after index', () => {
    assert.equal(integerCallback(callback(local(1))), true);
    assert.equal(integerCallback(suffixCallback()), true);
    for (const type of [new C.Func([C.Any.value, C.Int.value], C.Int.value), new C.Func([C.Int.value, C.Int.value], C.Any.value), new C.Func([C.Int.value], C.Int.value)]) {
        assert.equal(integerCallback(suffixCallback(type)), false);
    }
    for (const type of [C.Any.value, C.Int.value, new C.Func([C.Boolean.value, C.Int.value, C.Int.value], C.Int.value)]) {
        assert.equal(integerCallback(typed(type, suffixCallback())), false);
    }
    assert.equal(integerCallback(abs([0, 1, 2], local(1))), false);
});

test('read-only primitive proof preserves partial operations and rejects effects, aliases and inconsistent types', () => {
    const accept = body => assert.equal(readOnlyScalar([0, 1, 2])(body), true);
    const reject = body => assert.equal(readOnlyScalar([0, 1, 2])(body), false);
    for (const op of [S.OpAdd, S.OpSubtract, S.OpMultiply, S.OpDivide, S.OpMod]) accept(num(op.value, local(1), local(2)));
    accept(num(S.OpAdd.value, local(1), local(9, C.Int.value)));
    accept(choose(local(9, C.Boolean.value), local(1), local(2)));
    accept(e(new S.Let(new Just('flag'), 3, binary(new S.OpIntOrd(S.OpEq.value), local(0), int(0)), choose(local(3), local(1), local(2)))));
    reject(binary(S.OpArrayIndex.value, local(9), int(0)));
    reject(e(new S.App(ref('Foreign', 'callback'), [local(2)])));
    reject(ref('Foreign', 'integer')); reject(local(9));
    reject(num(S.OpAdd.value, bool(true), local(2)));
    reject(choose(local(0), local(1), local(2)));
    reject(choose(bool(true), local(1), bool(false)));
    reject(typed(C.Boolean.value, num(S.OpAdd.value, local(1), local(2))));
    reject(local(1, C.Boolean.value));
    reject(num(S.OpAdd.value, local(1), local(9, C.String.value)));
});

test('indexed fold admission refuses unknown targets, stages and recursive scopes before translation', () => {
    const never = () => { throw Error('must not translate a rejected candidate'); };
    const target = ref('Data.FoldableWithIndex', 'foldlWithIndexImpl');
    const supplied = [callback(local(1)), int(7), local(5)];
    const reject = (fn, args, recVars = []) => assert.deepEqual(emit(never)({ recVars })(0)(fn)(args), Nothing.value);
    reject(ref('Other', 'foldlWithIndexImpl'), supplied);
    reject(ref('Data.FoldableWithIndex', 'foldrWithIndexImpl'), supplied);
    reject(target, supplied.slice(0, 1)); reject(target, [...supplied, int(1)]);
    reject(target, supplied, ['deferredRecursiveValue']);
    reject(target, [typed(callbackType, local(5)), ...supplied.slice(1)]);
    reject(target, [typed(callbackType, abs([0], e(new S.Let(new Just('saved'), 6, int(1), abs([1, 2], local(1)))))), ...supplied.slice(1)]);
});

function generatedFixture() {
    const n = syntax => new NeutralExpr(syntax);
    const type = (ty, value) => n(new S.Typed(ty, value));
    const value = (level, ty = C.Int.value) => type(ty, n(new S.Local(new Just('x'), level)));
    const integer = number => n(new S.Lit(new C.LitInt(number)));
    const operation = (op, a, b) => type(C.Int.value, n(new S.PrimOp(new S.Op2(new S.OpIntNum(op), a, b))));
    const array = new C.Array(C.Int.value);
    const body = operation(S.OpSubtract.value,
        operation(S.OpMultiply.value, value(1), integer(3)),
        operation(S.OpAdd.value, value(0), value(2)));
    // Deliberately retain only the typed suffix, matching the observed real IR.
    const fn = n(new S.Abs([new Tuple(new Just('x'), 0)], type(new C.Func([C.Int.value, C.Int.value], C.Int.value),
        n(new S.Abs([new Tuple(new Just('x'), 1), new Tuple(new Just('x'), 2)], body)))));
    const fold = n(new S.Var(new C.Qualified(new Just('Data.FoldableWithIndex'), 'foldlWithIndexImpl')));
    const partial = type(new C.Func([array], C.Int.value), n(new S.App(fold, [fn, integer(17)])));
    const complete = type(new C.Func([array], C.Int.value), n(new S.Abs([new Tuple(new Just('x'), 4)],
        type(C.Int.value, n(new S.App(fold, [fn, integer(17), value(4, array)]))))));
    const metadata = withReboxFields({ elidedCtors: emptySet, ctorTypes: emptyMap, pointerAdtPaths: emptyMap,
        pointerAdtNodes: emptySet, pointerAdtLeaves: emptyMap, enumAdts: emptySet, enumCtors: emptySet,
        globalTypes: emptyMap, globalFunctions: emptyMap, classDeclsFields: emptyMap });
    return CodeGen.translate(metadata)({ name: 'IndexedFoldFixture', bindings: [{ recursive: false,
        bindings: [new Tuple('partial', partial), new Tuple('complete', complete)] }], comments: [], imports: emptySet,
        exports: emptySet, reExports: emptySet, dataTypes: emptyMap, dataDecls: [], classDecls: [], foreign: emptyMap,
        implementations: emptyMap, directives: emptyMap });
}

test('generated complete and reusable partial loops preserve ascending order, seed, overflow and input storage', t => {
    const source = generatedFixture();
    assert.match(source, /indexedFold_values_/);
    assert.doesNotMatch(source, /Get_Data_FoldableWithIndex_foldlWithIndexImpl|gopurs_runtime.Apply3/);
    const directory = mkdtempSync(join(tmpdir(), 'gopurs-indexed-fold-'));
    t.after(() => rmSync(directory, { recursive: true, force: true }));
    mkdirSync(join(directory, 'gopurs_runtime'));
    writeFileSync(join(directory, 'go.mod'), 'module gopurs/output\n\ngo 1.22\n');
    writeFileSync(join(directory, 'gopurs_runtime/runtime.go'), runtimeGoCode);
    writeFileSync(join(directory, 'fixture.go'), source.replace('package purescript', 'package main'));
    writeFileSync(join(directory, 'main.go'), `package main
import("fmt";"reflect";r "gopurs/output/gopurs_runtime")
func main(){cases:=0;for _,n:=range []int{0,1,2,3,7,17,65,1000}{for _,seed:=range []int64{0,-1,2147483647,-2147483648}{
xs:=make([]r.Value,n);want:=int64(17);for i:=range xs{xs[i]=r.Int(r.IntAdd(seed,int64(i)));want=r.IntSub(r.IntMul(want,3),r.IntAdd(int64(i),xs[i].IntVal))};before:=append([]r.Value{},xs...)
for _,f:=range []r.Value{Get_IndexedFoldFixture_partial(),Get_IndexedFoldFixture_complete()}{for replay:=0;replay<3;replay++{got:=r.Apply(f,r.Array(xs)).IntVal;if got!=want{panic(fmt.Sprint(n,seed,replay,got,want))};cases++}}
if !reflect.DeepEqual(xs,before){panic("input mutation")}}};fmt.Printf("%d indexed folds: ok\\n",cases)}
`);
    const result = spawnSync('go', ['run', '.'], { cwd: directory, encoding: 'utf8', env: { ...process.env, GOWORK: 'off' } });
    assert.equal(result.status, 0, result.stdout + result.stderr);
    assert.equal(result.stdout.trim(), '192 indexed folds: ok');
});
