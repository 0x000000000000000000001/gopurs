import assert from 'node:assert/strict';
import test from 'node:test';
import { mkdirSync, writeFileSync } from 'node:fs';
import { join } from 'node:path';
import { withReboxFields } from './codegen-metadata.mjs';
import { empty as emptyMap, insert, lookup } from '../output/Data.Map/index.js';
import { Just, Nothing } from '../output/Data.Maybe/index.js';
import { ordString } from '../output/Data.Ord/index.js';
import { Tuple } from '../output/Data.Tuple/index.js';
import { empty as emptySet } from '../output/Data.Set/index.js';
import * as Ref from '../output/Effect.Ref/index.js';
import * as C from '../output/PureScript.Backend.Optimizer.CoreFn/index.js';
import * as S from '../output/PureScript.Backend.Optimizer.Syntax/index.js';
import { TcoExpr } from '../output/PureScript.Backend.Optimizer.Codegen.Tco/index.js';
import { localId } from '../output/PureScript.Backend.Optimizer.FreeVars/index.js';
import * as Go from '../output/Gopurs.GoAst/index.js';
import { StmtEmpty, flattenStmts } from '../output/Gopurs.ExprContext/index.js';
import { printGoExpr } from '../output/Gopurs.Printer/index.js';
import { emit } from '../output/Gopurs.ArrayRangeFold/index.js';
import * as ArrayIntrinsics from '../output/Gopurs.ArrayIntrinsics/index.js';
import { annotated } from '../output/Gopurs.TypedExprs/index.js';
import { runtimeGoCode } from '../output/Gopurs.Runtime.Source/index.js';
import { goModule, runGo, writeGoFile } from './lib/go.mjs';

const int = C.Int.value, array = new C.Array(int), predicateType = new C.Func([int], C.Boolean.value);
const expr = syntax => new TcoExpr(null, syntax);
const typed = (ty, value) => expr(new S.Typed(ty, value));
const variable = (module, name) => expr(new S.Var(new C.Qualified(new Just(module), name)));
const factory = (name, ty) => typed(ty, expr(new S.App(variable('Probe', name), [typed(int, expr(new S.Lit(new C.LitInt(0))))])));
const uncurried = (name, args, ty = array, module = 'Data.Array') => typed(ty, expr(new S.UncurriedApp(variable(module, name), args)));
const strip = value => value.value1 instanceof S.Typed ? strip(value.value1.value1) : value.value1;
const integer = value => typed(int, expr(new S.Lit(new C.LitInt(value))));
const local = (name, level) => typed(int, expr(new S.Local(new Just(name), level)));
const element = local('element', 7);
const binary = (operation, left, right, type = int) => typed(type, expr(new S.PrimOp(new S.Op2(operation, left, right))));
const modulo = (left, right) => binary(new S.OpIntNum(S.OpMod.value), left, right);
const equal = (left, right) => binary(new S.OpIntOrd(S.OpEq.value), left, right, C.Boolean.value);
function purePredicate({ modulus = 2, residue = 0, body, commute = false } = {}) {
    const remainder = modulo(element, integer(modulus)), selected = integer(residue);
    return typed(predicateType, expr(new S.Abs([new Tuple(new Just('element'), 7)],
        body ?? (commute ? equal(selected, remainder) : equal(remainder, selected)))));
}

function fixture(options = {}) {
    const callback = options.callback ?? variable('Data.Semiring', 'intAdd');
    const seed = options.seed ?? factory('seed', int);
    const predicate = options.predicate ?? factory('predicate', predicateType);
    const start = options.start ?? factory('start', int), end = options.end ?? factory('end', int);
    const range = options.range ?? uncurried('rangeImpl', options.rangeArgs ?? [start, end], array, options.rangeOwner);
    const filtered = options.filtered ?? uncurried('filterImpl', options.filterArgs ?? [predicate, range], options.arrayType ?? array, options.filterOwner);
    return { fn: options.fold ?? variable('Data.Foldable', 'foldlArray'),
        args: options.args ?? [callback, seed, filtered], values: [callback, seed, predicate, start, end] };
}

function context() {
    return {
        metadata: withReboxFields({ pointerAdtPaths: emptyMap, pointerAdtNodes: emptySet, pointerAdtLeaves: emptyMap,
            enumAdts: emptySet, enumCtors: emptySet, elidedCtors: emptySet, ctorTypes: emptyMap,
            classDeclsFields: emptyMap, globalTypes: emptyMap, globalFunctions: emptyMap }),
        codegenStateRef: Ref.new({ declarations: [], globalId: 0, reboxPairs: emptySet })(),
        depth: 0, modNameStr: 'Probe', recVars: [], moduleFunctions: emptyMap, bound: emptyMap,
        tcoIdent: Nothing.value, loopCtx: [], options: { isTail: false, inEffectBlock: false },
        mbExpectedExprType: Nothing.value, mbExpectedGoType: Nothing.value,
    };
}

for (const [label, options] of [
    ['an unrelated fold', { fold: variable('Other', 'foldlArray') }],
    ['an opaque fold callback', { callback: factory('callback', new C.Func([int, int], int)) }],
    ['a similarly named callback', { callback: variable('Other', 'intAdd') }],
    ['a non-Int seed', { seed: factory('seed', C.Number.value) }],
    ['a non-Int bound', { end: factory('end', C.Number.value) }],
    ['an erased array annotation', { arrayType: C.Any.value }],
    ['a quantified predicate', { predicate: factory('predicate', new C.ForAll(['a'], predicateType)) }],
    ['a differently typed predicate', { predicate: factory('predicate', new C.Func([C.Number.value], C.Boolean.value)) }],
    ['an unrelated range', { rangeOwner: 'Other' }],
    ['an unrelated filter', { filterOwner: 'Other' }],
    ['a shared range result', { range: typed(array, expr(new S.Local(new Just('shared'), 0))) }],
    ['a shared filtered array', { filtered: typed(array, expr(new S.Local(new Just('shared'), 0))) }],
    ['partial folds', { args: [variable('Data.Semiring', 'intAdd')] }],
    ['partial ranges', { rangeArgs: [factory('start', int)] }],
    ['partial filters', { filterArgs: [factory('predicate', predicateType)] }],
]) {
    test(`range/filter/fold rejects ${label} before translating arguments`, () => {
        const source = fixture(options);
        assert.equal(emit(() => assert.fail('translated a rejected pipeline'))(context())(0)(source.fn)(source.args), Nothing.value);
    });
}

// Keep foreign factories observable, but run real annotation coercion and
// array-filter emission in the fallback. This catches native callback ABIs
// being lost while the captured references are translated.
function generated(variant, options = {}, ctx = context()) {
    const source = fixture(options), types = [Go.TypeValue.value, Go.TypeInt64.value,
        variant === 'Boxed' ? Go.TypeValue.value : variant === 'NativeInt'
            ? new Go.TypeFunc([Go.TypeInt64.value], Go.TypeBool.value)
            : new Go.TypeFunc([Go.TypeValue.value], Go.TypeValue.value),
        Go.TypeInt64.value, Go.TypeInt64.value];
    const factories = ['callback', 'seed', `predicate${variant}`, 'start', 'end'];
    const translation = ctx => nextId => value => {
        const index = source.values.indexOf(value);
        if (index >= 0) return { stmts: StmtEmpty.value, expr: new Go.GoCall(new Go.GoVar(factories[index]), []), exprType: types[index], nextId };
        if (value.value1 instanceof S.Typed) return annotated(translation)(ctx)(nextId)(value.value1.value0)(value.value1.value1);
        const syntax = strip(value);
        if (syntax instanceof S.Local) {
            const binding = lookup(ordString)(localId(syntax.value0)(syntax.value1))(ctx.bound);
            assert.ok(binding instanceof Just, 'fallback references must use captured values');
            return { stmts: StmtEmpty.value, expr: new Go.GoVar(binding.value0.name), exprType: binding.value0.goType, nextId };
        }
        if (syntax instanceof S.Lit) return {
            stmts: StmtEmpty.value, expr: new Go.GoInt(syntax.value0.value0), exprType: Go.TypeInt64.value, nextId };
        if (syntax instanceof S.PrimOp) {
            const operation = syntax.value0;
            assert.ok(operation instanceof S.Op2);
            const left = translation(ctx)(nextId)(operation.value1), right = translation(ctx)(nextId)(operation.value2);
            let emitted, type = Go.TypeInt64.value;
            if (operation.value0 instanceof S.OpIntNum) {
                const name = operation.value0.value0 === S.OpMod.value ? 'IntMod' : 'IntAdd';
                assert.ok(operation.value0.value0 === S.OpMod.value || operation.value0.value0 === S.OpAdd.value);
                emitted = new Go.GoCall(new Go.GoSelector(new Go.GoVar('gopurs_runtime'), name), [left.expr, right.expr]);
            } else {
                assert.ok(operation.value0 instanceof S.OpIntOrd);
                const comparison = operation.value0.value0 === S.OpEq.value ? '==' : '!=';
                assert.ok(operation.value0.value0 === S.OpEq.value || operation.value0.value0 === S.OpNotEq.value);
                emitted = new Go.GoBinOp(comparison, left.expr, right.expr);
                type = Go.TypeBool.value;
            }
            return { stmts: StmtEmpty.value, expr: emitted, exprType: type, nextId };
        }
        assert.ok(syntax instanceof S.UncurriedApp);
        const target = strip(syntax.value0).value0.value1;
        if (target === 'filterImpl') {
            const intrinsic = ArrayIntrinsics.recognize(ArrayIntrinsics.Uncurried.value)('Probe')(
                new Just({ mbMod: new Just('Data.Array'), name: 'filterImpl' }))(2);
            assert.ok(intrinsic instanceof Just);
            return ArrayIntrinsics.emitUncurried(translation)(ctx)(nextId)(syntax.value1)(intrinsic.value0);
        }
        const arguments_ = syntax.value1.map(arg => translation(ctx)(nextId)(arg));
        assert.equal(target, 'rangeImpl');
        return { stmts: StmtEmpty.value, expr: new Go.GoCall(new Go.GoVar('ordinaryRange'),
            arguments_.map(arg => arg.expr)), exprType: new Go.TypeNativeArray(Go.TypeInt64.value), nextId };
    };
    const result = emit(translation)(ctx)(0)(source.fn)(source.args);
    assert.ok(result instanceof Just);
    return [...flattenStmts(result.value0.stmts), new Go.GoReturn(result.value0.expr)].map(printGoExpr).join('\n');
}

test('unannotated literals and TAST-derived native local bounds are accepted', () => {
    const ctx = context();
    ctx.bound = insert(ordString)('n_0')({ name: 'n', goType: Go.TypeInt64.value })(ctx.bound);
    const body = generated('Boxed', { seed: expr(new S.Lit(new C.LitInt(0))),
        start: expr(new S.Lit(new C.LitInt(1))), end: expr(new S.Local(new Just('n'), 0)) }, ctx);
    assert.match(body, /range_filter_fold/);
});

for (const [label, type] of [['absent', null], ['Number', Go.TypeFloat64.value], ['erased', Go.TypeValue.value]]) {
    test(`unannotated locals with ${label} native bindings are rejected`, () => {
        const ctx = context();
        if (type) ctx.bound = insert(ordString)('n_0')({ name: 'n', goType: type })(ctx.bound);
        const source = fixture({ end: expr(new S.Local(new Just('n'), 0)) });
        assert.equal(emit(() => assert.fail('translated an unknown local'))(ctx)(0)(source.fn)(source.args), Nothing.value);
    });
}

test('explicit non-Int annotations cannot be overridden by literal syntax or a native local binding', () => {
    const ctx = context();
    ctx.bound = insert(ordString)('n_0')({ name: 'n', goType: Go.TypeInt64.value })(ctx.bound);
    for (const value of [expr(new S.Lit(new C.LitInt(3))), expr(new S.Local(new Just('n'), 0))]) {
        const source = fixture({ end: typed(C.Number.value, value) });
        assert.equal(emit(() => assert.fail('ignored an explicit Number annotation'))(ctx)(0)(source.fn)(source.args), Nothing.value);
    }
});

test('range fusion preserves capture order, inclusive endpoints and actual typed-filter fallback for all callback ABIs', t => {
    const boxed = generated('Boxed'), native = generated('Native'), nativeInt = generated('NativeInt');
    for (const body of [boxed, native, nativeInt]) {
        for (const name of ['callback', 'seed', 'start', 'end']) assert.equal((body.match(new RegExp(`\\b${name}\\(\\)`, 'g')) ?? []).length, 1);
        assert.match(body, /ordinaryRange\(/);
        assert.ok(body.indexOf('current) == (last') < body.indexOf('current = (current) + (step)'), 'endpoint checked before increment');
    }
    const directory = goModule(t, 'gopurs-array-range-fold-ir-', { runtime: runtimeGoCode });
    writeGoFile(directory, 'fusion_test.go', `package main
import ("testing"; "slices"; "gopurs/output/gopurs_runtime")
var trace []int64
var lower,upper,initial int64
func callback() gopurs_runtime.Value { trace=append(trace,-10001); return gopurs_runtime.Value{} }
func seed() int64 { trace=append(trace,-10002);return initial }
func start() int64 { trace=append(trace,-10004);return lower }
func end() int64 { trace=append(trace,-10005);return upper }
func predicateNative() func(gopurs_runtime.Value) gopurs_runtime.Value {
 trace=append(trace,-10003)
 return func(v gopurs_runtime.Value) gopurs_runtime.Value { trace=append(trace,v.IntVal);return gopurs_runtime.Bool(v.IntVal%2==0) }
}
func predicateBoxed() gopurs_runtime.Value { return gopurs_runtime.Func(predicateNative()) }
func predicateNativeInt() func(int64) bool {
 trace=append(trace,-10003)
 return func(v int64) bool { trace=append(trace,v);return v%2==0 }
}
func ordinaryRange(start,end int64) []int64 {
 step:=int64(1);if start>end {step=-1};values:=make([]int64,(end-start)*step+1)
 for i:=range values {values[i]=start+int64(i)*step};return values
}
func fusedBoxed() int64 {${boxed}}
func fusedNative() int64 {${native}}
func fusedNativeInt() int64 {${nativeInt}}
func TestFusion(t *testing.T) {
 ranges:=[][2]int64{{1,900},{3,-2},{-3,2},{0,0},{1,1},{2147483645,2147483647},{-2147483646,-2147483648},
 {2147483648,2147483650},{-2147483649,-2147483651},{9223372036854775805,9223372036854775807},{-9223372036854775806,-9223372036854775808}}
 for _,bounds:=range ranges {for _,s:=range []int64{0,-17,2147483647,-2147483648} {
  lower,upper,initial=bounds[0],bounds[1],s
  expectedTrace:=[]int64{-10001,-10002,-10003,-10004,-10005};want:=s
  for _,v:=range ordinaryRange(lower,upper) {expectedTrace=append(expectedTrace,v);if v%2==0 {want=int64(int32(want)+int32(v))}}
  for _,fn:=range []func()int64{fusedBoxed,fusedNative,fusedNativeInt} {trace=nil;got:=fn();if got!=want || !slices.Equal(trace,expectedTrace) {t.Fatalf("bounds=%v seed=%d got=%d want=%d trace=%v",bounds,s,got,want,trace)}}
 }}
}
`);
    const result = runGo(directory, ['test', '.'], { timeout: 60_000 });
    assert.equal(result.status, 0, result.stdout + result.stderr);
});

test('direct pure congruence lambdas admit a guarded sum with either equality order', () => {
    for (const options of [{}, { commute: true }, { modulus: 1 }, { modulus: 5, residue: 3 },
        { modulus: 2147483647, residue: 2147483646 }]) {
        const body = generated('NativeInt', { predicate: purePredicate(options) });
        assert.match(body, /func\(lower int64, upper int64, seed int64\) int64/);
        assert.match(body, /swap := lower/);
        assert.match(body, /first :=/);
        assert.match(body, /range_filter_fold/);
        assert.match(body, /ordinaryRange\(/);
        const formula = body.slice(body.indexOf('func(lower int64'), body.indexOf('range_filter_fold'));
        assert.doesNotMatch(formula, /IntMul/);
        for (const name of ['callback', 'seed', 'start', 'end']) {
            assert.equal((body.match(new RegExp(`\\b${name}\\(\\)`, 'g')) ?? []).length, 1);
            assert.ok(body.indexOf(`${name}()`) < body.indexOf('func(lower int64'));
        }
    }
});

test('pure-sum shortcut rejects unsupported constants, operations and other local bindings', () => {
    const ctx = context();
    for (const [name, level] of [['element', 8], ['divisor', 9]]) {
        ctx.bound = insert(ordString)(localId(new Just(name))(level))({ name: 'external', goType: Go.TypeInt64.value })(ctx.bound);
    }
    const options = [
        ...[0, -2, 2147483648].map(modulus => ({ modulus })),
        ...[-1, 2, 2147483648].map(residue => ({ residue })),
        { body: equal(modulo(local('element', 8), integer(2)), integer(0)) },
        { body: equal(modulo(element, local('divisor', 9)), integer(0)) },
        { body: equal(modulo(binary(new S.OpIntNum(S.OpAdd.value), element, integer(1)), integer(2)), integer(0)) },
        { body: binary(new S.OpIntOrd(S.OpNotEq.value), modulo(element, integer(2)), integer(0), C.Boolean.value) },
    ];
    for (const option of options) {
        const body = generated('NativeInt', { predicate: purePredicate(option) }, ctx);
        assert.doesNotMatch(body, /func\(lower int64, upper int64, seed int64\)/);
        assert.match(body, /range_filter_fold/);
    }
    for (const variant of ['Boxed', 'Native', 'NativeInt']) {
        assert.doesNotMatch(generated(variant), /func\(lower int64, upper int64, seed int64\)/,
            'opaque or effectful predicate factories cannot use the sum shortcut');
    }
});

test('generated pure sums preserve bidirectional ranges, exact arithmetic, captures and all fallbacks', t => {
    const configs = [{ modulus: 2, residue: 0 }, { modulus: 5, residue: 3 },
        { modulus: 1, residue: 0 }, { modulus: 7, residue: 6 },
        { modulus: 2147483647, residue: 2147483646 }];
    const bodies = configs.map(config => generated('NativeInt', { predicate: purePredicate(config) }));
    const largeCases = [];
    for (const [index, { modulus, residue }] of configs.entries()) {
        for (const [lower, upper] of [[1, 2147483647], [2147483647, 1], [2147483639, 2147483647],
            [2147483647, 2147483639], [2147483647, 2147483647], [2147483646, 2147483646]]) {
            const lo = BigInt(Math.min(lower, upper)), hi = BigInt(Math.max(lower, upper));
            const m = BigInt(modulus), r = BigInt(residue);
            const begin = lo > r ? (lo - r + m - 1n) / m : 0n;
            const end = hi >= r ? (hi - r) / m : -1n;
            const count = end >= begin ? end - begin + 1n : 0n;
            const sum = m * (end * (end + 1n) - (begin - 1n) * begin) / 2n + count * r;
            for (const seed of [-2147483648, -17, 0, 2147483647]) {
                const want = count === 0n ? BigInt(seed) : BigInt.asIntN(32, BigInt(seed) + sum);
                largeCases.push(`{fused${index}, ${lower}, ${upper}, ${seed}, ${want}}`);
            }
        }
    }
    const source = `package main
import ("testing"; "slices"; "gopurs/output/gopurs_runtime")
var trace []int64
var lower,upper,initial int64
var rangeCalls int
func callback() gopurs_runtime.Value {trace=append(trace,1);return gopurs_runtime.Value{}}
func seed() int64 {trace=append(trace,2);return initial}
func start() int64 {trace=append(trace,3);return lower}
func end() int64 {trace=append(trace,4);return upper}
func ordinaryRange(start,end int64) []int64 {
 rangeCalls++;step:=int64(1);if start>end {step=-1};values:=make([]int64,(end-start)*step+1)
 for i:=range values {values[i]=start+int64(i)*step};return values
}
${bodies.map((body, index) => `func fused${index}() int64 {${body}}`).join('\n')}
func TestPureSums(t *testing.T) {
 ranges:=[][2]int64{{1,900},{900,1},{2,2},{1,1},{3,4},{4,3},{-3,2},{2,-3},{0,0},
 {2147483645,2147483647},{2147483647,2147483645},{-2147483646,-2147483648},
 {2147483648,2147483650},{2147483650,2147483648},{-2147483649,-2147483651},
 {9223372036854775805,9223372036854775807},{-9223372036854775806,-9223372036854775808}}
 checks:=0
 for _,c:=range []struct{f func()int64;m,r int64}{${configs.map((c, i) => `{fused${i},${c.modulus},${c.residue}}`).join(',')}} {
  for _,bounds:=range ranges {for _,s:=range []int64{-9223372036854775808,-2147483649,-2147483648,0,2147483647,2147483648,9223372036854775807} {
   lower,upper,initial=bounds[0],bounds[1],s;want:=s
   for _,v:=range ordinaryRange(lower,upper) {mod:=v%c.m;if mod<0 {mod+=c.m};if mod==c.r {want=int64(int32(want)+int32(v))}}
   trace=nil;rangeCalls=0;got:=c.f();checks++
   if got!=want || !slices.Equal(trace,[]int64{1,2,3,4}) {t.Fatalf("bounds=%v seed=%d m=%d r=%d got=%d want=%d trace=%v",bounds,s,c.m,c.r,got,want,trace)}
   wide:=lower!=int64(int32(lower)) || upper!=int64(int32(upper))
   if (rangeCalls==1)!=wide {t.Fatalf("wide range fallback: bounds=%v calls=%d",bounds,rangeCalls)}
  }}
 }
 for _,c:=range []struct{f func()int64;lo,hi,seed,want int64}{${largeCases.join(',\n')}} {
  lower,upper,initial=c.lo,c.hi,c.seed;trace=nil;rangeCalls=0;got:=c.f();checks++
  if got!=c.want || rangeCalls!=0 || !slices.Equal(trace,[]int64{1,2,3,4}) {t.Fatalf("large bounds=%d,%d seed=%d got=%d want=%d",lower,upper,initial,got,c.want)}
 }
 t.Logf("%d bidirectional, large, wide and capture-order cases passed",checks)
}
`;
    const directory = goModule(t, 'gopurs-array-pure-sum-ir-', { runtime: runtimeGoCode });
    writeGoFile(directory, 'pure_sum_test.go', source);
    const result = runGo(directory, ['test', '-count=1', '-v', '.'], { timeout: 60_000 });
    if (process.env.GOPURS_ARRAY_SUM_TEST_OUTPUT) {
        mkdirSync(process.env.GOPURS_ARRAY_SUM_TEST_OUTPUT, { recursive: true });
        writeFileSync(join(process.env.GOPURS_ARRAY_SUM_TEST_OUTPUT, 'pure_sum_test.go'), source);
        writeFileSync(join(process.env.GOPURS_ARRAY_SUM_TEST_OUTPUT, 'go-test.log'), result.stdout + result.stderr);
    }
});
