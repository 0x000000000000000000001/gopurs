import assert from 'node:assert/strict';
import test from 'node:test';
import { spawnSync } from 'node:child_process';
import { mkdirSync, mkdtempSync, rmSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join } from 'node:path';
import { empty as emptyMap, insert } from '../output/Data.Map/index.js';
import { ordString } from '../output/Data.Ord/index.js';
import { empty as emptySet } from '../output/Data.Set/index.js';
import { Just, Nothing } from '../output/Data.Maybe/index.js';
import { Tuple } from '../output/Data.Tuple/index.js';
import * as Ref from '../output/Effect.Ref/index.js';
import * as G from '../output/Gopurs.GoAst/index.js';
import { SymbolReference, generated } from '../output/Gopurs.GoSymbol/index.js';
import * as CodeGen from '../output/Gopurs.CodeGen/index.js';
import * as Fold from '../output/Gopurs.ClosedIndexedFolds/index.js';
import * as Fusion from '../output/Gopurs.ArrayTraverseFold/index.js';
import { StmtEmpty } from '../output/Gopurs.ExprContext/index.js';
import * as C from '../output/PureScript.Backend.Optimizer.CoreFn/index.js';
import * as S from '../output/PureScript.Backend.Optimizer.Syntax/index.js';
import { NeutralExpr } from '../output/PureScript.Backend.Optimizer.Semantics/index.js';
import { analyze } from '../output/PureScript.Backend.Optimizer.Codegen.Tco/index.js';
import { runtimeGoCode } from '../output/Gopurs.Runtime.Source/index.js';
import { withReboxFields } from './codegen-metadata.mjs';

const n = syntax => new NeutralExpr(syntax);
const typed = (type, value) => n(new S.Typed(type, value));
const int = value => typed(C.Int.value, n(new S.Lit(new C.LitInt(value))));
const array = new C.Array(C.Int.value);
const maybe = element => new C.ADT('Data.Maybe.Maybe', ['Data', 'Maybe', 'Maybe'], [element]);
const local = (level, type = C.Int.value) => typed(type, n(new S.Local(new Just('value'), level)));
const qual = (owner, name) => new C.Qualified(new Just(owner), name);
const ref = (owner, name) => n(new S.Var(qual(owner, name)));
const apply = (fn, args, type) => typed(type, n(new S.App(fn, args)));
const arithmetic = (op, a, b) => typed(C.Int.value, n(new S.PrimOp(new S.Op2(new S.OpIntNum(op), a, b))));
const lambda = (levels, ret, body, types = levels.map(() => C.Int.value)) => typed(new C.Func(types, ret), n(new S.Abs(levels.map(level => new Tuple(new Just('value'), level)), body)));
const just = (value, owner = 'Data.Maybe') => typed(maybe(C.Int.value), n(new S.CtorSaturated(qual(owner, 'Just'), C.SumType.value, 'Maybe', 'Just', [new Tuple('value0', value)])));
const mapper = (indexed = false) => lambda(indexed ? [30, 31] : [31], maybe(C.Int.value), just(arithmetic(S.OpAdd.value, local(31), indexed ? local(30) : int(1))));
const foldCallback = lambda([0, 1, 2], C.Int.value, arithmetic(S.OpSubtract.value,
    arithmetic(S.OpMultiply.value, local(1), int(3)), arithmetic(S.OpAdd.value, local(0), arithmetic(S.OpMultiply.value, local(2), int(3)))));
const consumer = lambda([4], C.Int.value, apply(ref('Data.FoldableWithIndex', 'foldlWithIndexImpl'), [foldCallback, int(17), local(4, array)], C.Int.value), [array]);
const producer = (callback = mapper(), indexed = false, input = local(20, array), owner = 'Data.Maybe') => apply(
    ref(indexed ? 'Data.TraversableWithIndex' : 'Data.Traversable', indexed ? 'traverseWithIndex' : 'traverse'),
    [ref(indexed ? 'Data.TraversableWithIndex' : 'Data.Traversable', indexed ? 'traversableWithIndexArray' : 'traversableArray'), ref(owner, 'applicativeMaybe'), callback, input], maybe(array));
const tagged = (tag, level = 40) => n(new S.PrimOp(new S.Op1(new S.OpIsTag(qual('Data.Maybe', tag)), local(level, maybe(array)))));
const payload = (level = 40) => typed(array, n(new S.Accessor(local(level, maybe(array)), new S.GetCtorField(qual('Data.Maybe', 'Just'), C.SumType.value, 'Maybe', 'Just', 'value0', 0))));
const branch = (value = payload(), consume = ref('TraverseFoldFixture', 'consume')) => typed(C.Int.value, n(new S.Branch([
    new S.Pair(tagged('Nothing'), int(-1)),
    new S.Pair(tagged('Just'), apply(consume, [value], C.Int.value)),
], n(new S.Fail('Failed pattern match')))));
const whole = (produced, body = branch()) => typed(C.Int.value, n(new S.Let(new Just('result'), 40, produced, body)));
const info = (name, fArgs, fRet, descriptor = Nothing.value) => ({ reference: new SymbolReference(generated('Factories')(`Call_Factories_${name}`)), fArgs, fRet, arity: fArgs.length, arrayFold: descriptor });
const put = (key, value, map) => insert(ordString)(key)(value)(map);
const metadata = () => withReboxFields({ elidedCtors: emptySet, ctorTypes: emptyMap, pointerAdtPaths: emptyMap, pointerAdtNodes: emptySet,
    pointerAdtLeaves: emptyMap, enumAdts: emptySet, enumCtors: emptySet, globalTypes: emptyMap,
    globalFunctions: put('Factories.input', info('input', [new G.TypeNativeArray(G.TypeInt64.value)], new G.TypeNativeArray(G.TypeInt64.value)), emptyMap), classDeclsFields: emptyMap });
const context = () => ({ metadata: metadata(), codegenStateRef: Ref.new({ declarations: [], globalId: 0, reboxPairs: emptySet })(),
    depth: 0, modNameStr: 'TraverseFoldFixture', recVars: [], moduleFunctions: put('consume', info('consume', [new G.TypeNativeArray(G.TypeInt64.value)], G.TypeInt64.value, Fold.describe(consumer)), emptyMap),
    bound: emptyMap, tcoIdent: Nothing.value, loopCtx: [], options: { isTail: false, inEffectBlock: false }, mbExpectedExprType: Nothing.value, mbExpectedGoType: Nothing.value });
const emit = (translate, ctx, produced, body) => Fusion.emit(translate)(ctx)(0)(new Just('result'))(40)(analyze([])(produced))(analyze([])(body));

test('Just mapper proves closure, canonical constructor and every type layer', () => {
    for (const indexed of [false, true]) assert.ok(Fusion.justMapper(indexed ? 2 : 1)(analyze([])(mapper(indexed))) instanceof Just);
    const reject = value => assert.deepEqual(Fusion.justMapper(1)(analyze([])(value)), Nothing.value);
    reject(lambda([31], maybe(C.Int.value), just(local(99))));
    reject(lambda([31], maybe(C.Int.value), just(int(1), 'Other.Maybe')));
    reject(lambda([31], maybe(C.Int.value), typed(maybe(C.String.value), just(int(1)))));
    reject(lambda([31], maybe(C.Int.value), just(typed(C.String.value, int(1)))));
    reject(lambda([31], maybe(C.Int.value), just(apply(ref('Foreign', 'unknown'), [local(31)], C.Int.value))));
    reject(lambda([31], maybe(C.Int.value), just(arithmetic(S.OpDivide.value, int(1), int(0)))));
    reject(lambda([31], maybe(C.Int.value), n(new S.CtorSaturated(qual('Data.Maybe', 'Nothing'), C.SumType.value, 'Maybe', 'Nothing', []))));
    reject(typed(new C.Func([C.Int.value], maybe(C.Int.value)), lambda([31], maybe(C.Int.value), just(int(1)), [C.String.value])));
    reject(typed(new C.Func([C.Int.value], maybe(C.Int.value)), typed(new C.Func([new C.TypeVar('a')], new C.TypeVar('a')), mapper())));
});

test('fusion keeps dictionary, effect, stage and observed-result barriers', () => {
    const never = () => { throw Error('rejected candidate must not translate'); };
    const reject = (produced, body = branch(), ctx = context()) => assert.deepEqual(emit(never, ctx, produced, body), Nothing.value);
    reject(producer(mapper(), false, local(20, array), 'Other'));
    reject(typed(maybe(new C.Array(C.String.value)), producer()));
    reject(producer(mapper(), false, typed(new C.Array(C.String.value), local(20, array))));
    reject(producer(), branch(payload(99)));
    reject(producer(), branch(payload(), typed(new C.Func([array], C.String.value), ref('TraverseFoldFixture', 'consume'))));
    reject(producer(), branch(typed(array, n(new S.Accessor(typed(C.String.value, local(40, maybe(array))), new S.GetCtorField(qual('Data.Maybe', 'Just'), C.SumType.value, 'Maybe', 'Just', 'value0', 0))))));
    reject(producer(), n(new S.Let(new Just('unused'), 42, apply(ref('Foreign', 'mark'), [int(1)], C.Int.value), branch())));
    reject(producer(), n(new S.EffectPure(branch())));
    reject(producer(), branch(), { ...context(), recVars: ['cell'] });
    const raw = producer().value1;
    reject(typed(maybe(array), n(new S.App(n(new S.UncurriedApp(raw.value0, [])), raw.value1))));
    reject(producer(), branch(typed(array, n(new S.Accessor(local(40, maybe(array)), new S.GetCtorField(qual('Other.Maybe', 'Just'), C.SumType.value, 'Maybe', 'Just', 'value0', 0))))));
    const dummy = _ctx => nextId => _expr => ({ stmts: StmtEmpty.value, expr: new G.GoInt(0), exprType: G.TypeInt64.value, nextId });
    assert.ok(emit(dummy, context(), producer(), branch()) instanceof Just);
    assert.ok(emit(dummy, { ...context(), options: { isTail: false, inEffectBlock: true } }, producer(), n(new S.EffectPure(branch()))) instanceof Just);
});

function generatedFixture() {
    return CodeGen.translate(metadata())({ name: 'TraverseFoldFixture', bindings: [{ recursive: false, bindings: [
        new Tuple('consume', consumer),
        new Tuple('direct', lambda([20], C.Int.value, whole(producer()), [array])),
        new Tuple('indexed', lambda([20], C.Int.value, whole(producer(mapper(true), true)), [array])),
        new Tuple('identity', lambda([20], C.Int.value, whole(producer(lambda([31], maybe(C.Int.value), just(local(31))))), [array])),
        new Tuple('factory', lambda([20], C.Int.value, whole(producer(mapper(), false, apply(ref('Factories', 'input'), [local(20, array)], array))), [array])),
    ] }], comments: [], imports: emptySet, exports: emptySet, reExports: emptySet, dataTypes: emptyMap, dataDecls: [], classDecls: [], foreign: emptyMap, implementations: emptyMap, directives: emptyMap });
}

test('generated traverse folds preserve indices, wide arithmetic, seed and external alias mutation', t => {
    const directory = mkdtempSync(join(tmpdir(), 'gopurs-traverse-fold-'));
    t.after(() => rmSync(directory, { recursive: true, force: true }));
    mkdirSync(join(directory, 'gopurs_runtime'));
    writeFileSync(join(directory, 'go.mod'), 'module gopurs/output\n\ngo 1.22\n');
    writeFileSync(join(directory, 'gopurs_runtime/runtime.go'), runtimeGoCode);
    const source = generatedFixture();
    assert.match(source, /traverseFold_fold_/);
    for (const name of ['direct', 'indexed', 'identity', 'factory']) {
        const start = source.indexOf(`func Call_TraverseFoldFixture_${name}(`);
        const end = source.indexOf('\nfunc ', start + 1);
        const worker = source.slice(start, end < 0 ? undefined : end);
        assert.match(worker, /traverseFold_fold_/);
        assert.doesNotMatch(worker, /traverseTotal_output_|make\(\[\]int64/);
    }
    writeFileSync(join(directory, 'fixture.go'), source.replace('package purescript', 'package main'));
    writeFileSync(join(directory, 'main.go'), `package main
import("fmt";"math";"reflect";r "gopurs/output/gopurs_runtime")
var trace []string
func Get_Data_Traversable_traversableArray()r.Value{trace=append(trace,"traversable");return r.Value{}}
func Get_Data_TraversableWithIndex_traversableWithIndexArray()r.Value{trace=append(trace,"indexed");return r.Value{}}
func Get_Data_Maybe_applicativeMaybe()r.Value{trace=append(trace,"applicative");return r.Value{}}
func Call_Factories_input(xs []int64)[]int64{trace=append(trace,"input");if len(xs)>0{xs[0]=r.IntAdd(xs[0],17)};return xs}
func expected(xs []int64,mode int)int64{acc:=int64(17);for i,v:=range xs{if mode<2{delta:=int64(1);if mode==1{delta=int64(i)};v=r.IntAdd(v,delta)};acc=r.IntSub(r.IntMul(acc,3),r.IntAdd(int64(i),r.IntMul(v,3)))};return acc}
func main(){checks:=0;values:=[]int64{0,1,-1,math.MaxInt32,math.MinInt32,1<<52,(1<<52)+1,-(1<<52)-1,math.MaxInt64,math.MinInt64};for n:=0;n<37;n++{for offset:=range values{xs:=make([]int64,n);for i:=range xs{xs[i]=values[(i+offset)%len(values)]};before:=make([]int64,n);copy(before,xs)
trace=nil;if got,want:=Call_TraverseFoldFixture_direct(xs),expected(xs,0);got!=want{panic(fmt.Sprint("direct",n,offset,got,want))};if !reflect.DeepEqual(trace,[]string{"traversable","applicative"})||!reflect.DeepEqual(xs,before){panic("direct order or mutation")};checks++
trace=nil;if got,want:=Call_TraverseFoldFixture_indexed(xs),expected(xs,1);got!=want{panic(fmt.Sprint("indexed",n,offset,got,want))};if !reflect.DeepEqual(trace,[]string{"indexed","applicative"})||!reflect.DeepEqual(xs,before){panic("indexed order or mutation")};checks++
trace=nil;if got,want:=Call_TraverseFoldFixture_identity(xs),expected(xs,2);got!=want{panic(fmt.Sprint("identity wide multiply",n,offset,got,want))};if !reflect.DeepEqual(trace,[]string{"traversable","applicative"})||!reflect.DeepEqual(xs,before){panic("identity order or mutation")};checks++
changed:=make([]int64,n);copy(changed,xs);if n>0{changed[0]=r.IntAdd(changed[0],17)};trace=nil;if got,want:=Call_TraverseFoldFixture_factory(xs),expected(changed,0);got!=want{panic(fmt.Sprint("factory",n,offset,got,want))};if !reflect.DeepEqual(trace,[]string{"traversable","applicative","input"})||!reflect.DeepEqual(xs,changed){panic("factory phase or mutation")};checks++
}};fmt.Println(checks,"traverse fold controls passed")}
`);
    const result = spawnSync('go', ['run', '.'], { cwd: directory, encoding: 'utf8', env: { ...process.env, GOWORK: 'off' } });
    assert.equal(result.status, 0, result.stdout + result.stderr);
    assert.equal(result.stdout.trim(), '1480 traverse fold controls passed');
});
