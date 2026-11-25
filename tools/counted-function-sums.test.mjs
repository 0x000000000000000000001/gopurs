import assert from 'node:assert/strict';
import { spawnSync } from 'node:child_process';
import { mkdirSync, mkdtempSync, rmSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join } from 'node:path';
import test from 'node:test';
import { empty as emptyMap } from '../output/Data.Map/index.js';
import { Just, Nothing } from '../output/Data.Maybe/index.js';
import { empty as emptySet, singleton } from '../output/Data.Set/index.js';
import { Tuple } from '../output/Data.Tuple/index.js';
import * as CodeGen from '../output/Gopurs.CodeGen/index.js';
import { optimize } from '../output/Gopurs.CountedFunctionSums/index.js';
import { runtimeGoCode } from '../output/Gopurs.Runtime.Source/index.js';
import { withReboxFields } from './codegen-metadata.mjs';
import * as C from '../output/PureScript.Backend.Optimizer.CoreFn/index.js';
import { NeutralExpr } from '../output/PureScript.Backend.Optimizer.Semantics/index.js';
import * as S from '../output/PureScript.Backend.Optimizer.Syntax/index.js';

const owner = 'IteratorTranslationFixture';
const int = C.Int.value;
const unary = new C.Func([int], int);
const iterator = new C.Func([unary, int], int);
const producer = new C.Func([int, unary, int], int);
const expr = syntax => new NeutralExpr(syntax);
const typed = (type, node) => expr(new S.Typed(type, node));
const local = (type, level) => typed(type, expr(new S.Local(Nothing.value, level)));
const lit = n => typed(int, expr(new S.Lit(new C.LitInt(n))));
const variable = (type, name, module = owner) => typed(type,
    expr(new S.Var(new C.Qualified(new Just(module), name))));
const call = (type, fn, ...args) => typed(type, expr(new S.App(fn, args)));
const lambda = (type, levels, body) => typed(type,
    expr(new S.Abs(levels.map(n => new Tuple(Nothing.value, n)), body)));
const op = (operator, a, b, type = int) => typed(type, expr(new S.PrimOp(new S.Op2(operator, a, b))));
const plus = (a, b) => op(new S.OpIntNum(S.OpAdd.value), a, b);
const minus = (a, b) => op(new S.OpIntNum(S.OpSubtract.value), a, b);
const repeat = n => call(iterator, variable(producer, 'assemble'), n);
const addBinding = (module, name, body) => module.bindings.push({ recursive: false,
    bindings: [new Tuple(name, body)] });
const find = (module, name) => module.bindings.flatMap(g => g.bindings).find(p => p.value0 === name).value1;
const translate = module => optimize(singleton('assemble'))(module);
const metadata = withReboxFields({
    elidedCtors: emptySet, ctorTypes: emptyMap, pointerAdtPaths: emptyMap,
    pointerAdtNodes: emptySet, pointerAdtLeaves: emptyMap, enumAdts: emptySet,
    enumCtors: emptySet, globalTypes: emptyMap, globalFunctions: emptyMap,
    classDeclsFields: emptyMap,
});

function fixture() {
    const n = local(int, 0), fn = local(unary, 2), seed = local(int, 3);
    const base = lambda(iterator, [10, 11], local(int, 11));
    const recursive = call(iterator, variable(producer, 'assemble'), minus(n, lit(1)));
    const step = typed(iterator, expr(new S.Let(Nothing.value, 1, recursive,
        lambda(iterator, [2, 3], call(int, fn, call(int, local(iterator, 1), fn, seed))))));
    const condition = op(new S.OpIntOrd(S.OpEq.value), n, lit(0), C.Boolean.value);
    return {
        name: owner,
        bindings: [{ recursive: true, bindings: [new Tuple('assemble',
            lambda(producer, [0], typed(iterator, expr(new S.Branch([new S.Pair(condition, base)], step)))))] }],
        comments: [], imports: emptySet, exports: emptySet, reExports: emptySet,
        dataTypes: emptyMap, dataDecls: [], classDecls: [], foreign: emptyMap,
        implementations: emptyMap, directives: emptyMap,
    };
}

function consumer(callback = lambda(unary, [2], plus(local(int, 2), lit(1))), seed = lit(0)) {
    return lambda(new C.Func([int], int), [0], call(int, repeat(local(int, 0)), callback, seed));
}

test('translation reduction uses structure and retains a guarded original fallback', () => {
    const source = fixture();
    addBinding(source, 'sum', consumer());
    const serialized = JSON.stringify(source);
    const result = translate(source);
    assert.equal(JSON.stringify(source), serialized);
    assert.notDeepEqual(find(result, 'sum'), find(source, 'sum'));
    const body = find(result, 'sum').value1.value1;
    assert.ok(body.value1 instanceof S.Branch);
    assert.deepEqual(body.value1.value1, find(source, 'sum').value1.value1);
    assert.deepEqual(find(result, 'assemble'), find(source, 'assemble'));
});

for (const [label, callback] of [
    ['an opaque callback', local(unary, 7)],
    ['an imported function', variable(unary, 'callback', 'External')],
    ['a foreign-looking unbound function', variable(unary, 'callback')],
    ['a duplicated argument', lambda(unary, [2], plus(local(int, 2), local(int, 2)))],
    ['a discarded argument', lambda(unary, [2], lit(9))],
    ['multiplication rather than a translation', lambda(unary, [2], op(new S.OpIntNum(S.OpMultiply.value), local(int, 2), lit(2)))],
    ['a failing callback', lambda(unary, [2], typed(int, expr(new S.Fail('deferred'))))],
    ['a callback performing an unknown call', lambda(unary, [2], call(int, variable(unary, 'effect'), local(int, 2)))],
    ['a Number annotation on integer-looking arithmetic', lambda(unary, [2], typed(C.Number.value, plus(local(int, 2), lit(1))))],
    ['incompatible nested annotations on a callback argument', lambda(unary, [2], plus(typed(int, typed(C.Number.value, local(int, 2))), lit(1)))],
]) {
    test(`translation reduction retains ${label}`, () => {
        const source = fixture();
        addBinding(source, 'sum', consumer(callback));
        assert.deepEqual(find(translate(source), 'sum'), find(source, 'sum'));
    });
}

test('an effectful or failing argument cannot disappear when its parameter is unused', () => {
    const source = fixture();
    const ignored = lambda(new C.Func([int], int), [8], call(int, repeat(local(int, 0)),
        lambda(unary, [2], plus(local(int, 2), lit(3))), lit(0)));
    addBinding(source, 'sum', lambda(unary, [0], call(int, ignored,
        call(int, variable(unary, 'unknown'), local(int, 0)))));
    const optimized = translate(source);
    assert.ok(JSON.stringify(find(optimized, 'sum')).includes('unknown'));
});

test('negative counts keep the original divergent producer behind the runtime guard', () => {
    const source = fixture();
    const original = call(int, repeat(lit(-1)), lambda(unary, [2], plus(local(int, 2), lit(1))), lit(0));
    addBinding(source, 'negative', original);
    const result = find(translate(source), 'negative');
    assert.ok(result.value1 instanceof S.Branch);
    assert.deepEqual(result.value1.value1, original);
    const condition = result.value1.value0[0].value0.value1.value0;
    assert.deepEqual(condition.value0, new S.OpIntOrd(S.OpGte.value));
    assert.deepEqual(condition.value1, lit(-1));
    assert.deepEqual(condition.value2, lit(0));
});

function runtimeModule() {
    const source = fixture();
    addBinding(source, 'sum', consumer());
    addBinding(source, 'down', consumer(lambda(unary, [2], minus(local(int, 2), lit(3))), lit(2147483647)));
    addBinding(source, 'largeStep', consumer(lambda(unary, [2], plus(local(int, 2), lit(2147483647))), lit(-2147483648)));
    const dual = new C.Func([int, int], int);
    addBinding(source, 'seeded', lambda(dual, [0, 1], call(int, repeat(local(int, 0)),
        lambda(unary, [2], plus(local(int, 2), lit(1))), local(int, 1))));
    // Arbitrary renamed wrappers, nesting and reusable partially applied values.
    const compose = (left, right) => lambda(iterator, [3, 4], call(int, left,
        call(unary, right, local(unary, 3)), local(int, 4)));
    addBinding(source, 'square', lambda(producer, [0], compose(repeat(local(int, 0)), repeat(local(int, 0)))));
    addBinding(source, 'fourth', lambda(producer, [0], compose(
        call(iterator, variable(producer, 'square'), local(int, 0)),
        call(iterator, variable(producer, 'square'), local(int, 0)))));
    addBinding(source, 'fifth', lambda(producer, [0], compose(
        call(iterator, variable(producer, 'fourth'), local(int, 0)), repeat(local(int, 0)))));
    addBinding(source, 'composed', lambda(unary, [0], call(int,
        call(iterator, variable(producer, 'fifth'), local(int, 0)),
        lambda(unary, [8], plus(local(int, 8), lit(1))), lit(0))));
    addBinding(source, 'added', lambda(unary, [0], call(int, repeat(local(int, 0)),
        lambda(unary, [8], plus(local(int, 8), lit(2))),
        call(int, repeat(local(int, 0)), lambda(unary, [8], plus(local(int, 8), lit(3))), lit(7)))));
    return source;
}

test('generated arithmetic preserves counts, composition, signed overflow and wide-seed fallback', t => {
    const source = runtimeModule();
    const code = CodeGen.translate(metadata)(source);
    assert.ok(code.includes('countedFactor'), 'actual emitter must receive modular products');
    const directory = mkdtempSync(join(tmpdir(), 'gopurs-counted-sums-'));
    t.after(() => rmSync(directory, { recursive: true, force: true }));
    mkdirSync(join(directory, 'purescript'));
    mkdirSync(join(directory, 'gopurs_runtime'));
    writeFileSync(join(directory, 'go.mod'), 'module gopurs/output\n\ngo 1.22\n');
    writeFileSync(join(directory, 'gopurs_runtime/runtime.go'), runtimeGoCode);
    writeFileSync(join(directory, 'purescript/fixture.go'), code);
    writeFileSync(join(directory, 'main.go'), `package main
import("fmt"; p "gopurs/output/purescript")
func check(label string, got,want int64){if got!=want{panic(fmt.Sprintf("%s: got %d; want %d",label,got,want))}}
func main(){
  for _,n:=range []int64{0,1,2,3,10,17,65537,2147483647}{
    check("sum",p.Call_${owner}_sum(n),n)
    check("down",p.Call_${owner}_down(n),int64(int32(2147483647-3*n)))
    check("largeStep",p.Call_${owner}_largeStep(n),int64(int32(-2147483648+n*2147483647)))
    check("added",p.Call_${owner}_added(n),int64(int32(7+5*n)))
    result:=uint32(n); result=result*result*result*result*result
    check("composed",p.Call_${owner}_composed(n),int64(int32(result)))
  }
  for _,n:=range []int64{0,1,2,3,10}{for _,seed:=range []int64{-2147483648,0,2147483647,4294967296}{
    want:=seed; for i:=int64(0);i<n;i++{want=int64(int32(want+1))}
    check("seeded",p.Call_${owner}_seeded(n,seed),want)
  }}
  fmt.Println("counted translations verified")
}
`);
    const result = spawnSync('go', ['run', '.'], { cwd: directory, encoding: 'utf8', timeout: 30_000,
        env: { ...process.env, GOWORK: 'off', GOCACHE: process.env.GOCACHE ?? join(tmpdir(), 'gopurs-counted-tests-go-cache') } });
    assert.ifError(result.error);
    assert.equal(result.status, 0, result.stdout + result.stderr);
    assert.equal(result.stdout, 'counted translations verified\n');
});
