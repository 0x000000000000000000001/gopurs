import assert from 'node:assert/strict';
import test from 'node:test';
import { spawnSync } from 'node:child_process';
import { mkdirSync, mkdtempSync, rmSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join } from 'node:path';
import { empty as emptyMap, insert } from '../output/Data.Map/index.js';
import { ordString } from '../output/Data.Ord/index.js';
import * as G from '../output/Gopurs.GoAst/index.js';
import { SymbolReference, generated } from '../output/Gopurs.GoSymbol/index.js';
import { empty as emptySet } from '../output/Data.Set/index.js';
import { Just, Nothing } from '../output/Data.Maybe/index.js';
import { Tuple } from '../output/Data.Tuple/index.js';
import * as CodeGen from '../output/Gopurs.CodeGen/index.js';
import * as Fold from '../output/Gopurs.ClosedIndexedFolds/index.js';
import * as Fusion from '../output/Gopurs.ArrayProducerFold/index.js';
import * as C from '../output/PureScript.Backend.Optimizer.CoreFn/index.js';
import * as S from '../output/PureScript.Backend.Optimizer.Syntax/index.js';
import { NeutralExpr } from '../output/PureScript.Backend.Optimizer.Semantics/index.js';
import { analyze } from '../output/PureScript.Backend.Optimizer.Codegen.Tco/index.js';
import { runtimeGoCode } from '../output/Gopurs.Runtime.Source/index.js';
import { withReboxFields } from './codegen-metadata.mjs';

const n = syntax => new NeutralExpr(syntax);
const typed = (type, value) => n(new S.Typed(type, value));
const integer = value => typed(C.Int.value, n(new S.Lit(new C.LitInt(value))));
const array = new C.Array(C.Int.value);
const local = (level, type = C.Int.value) => typed(type, n(new S.Local(new Just('arg'), level)));
const ref = (owner, name) => n(new S.Var(new C.Qualified(new Just(owner), name)));
const apply = (fn, args, type) => typed(type, n(new S.App(fn, args)));
const arithmetic = (op, a, b) => typed(C.Int.value, n(new S.PrimOp(new S.Op2(new S.OpIntNum(op), a, b))));
const lambda = (levels, args, ret, body) => typed(new C.Func(args, ret), n(new S.Abs(levels.map(level => new Tuple(new Just('arg'), level)), body)));
const foldCallback = lambda([0, 1, 2], [C.Int.value, C.Int.value, C.Int.value], C.Int.value,
    arithmetic(S.OpSubtract.value, arithmetic(S.OpMultiply.value, local(1), integer(3)), arithmetic(S.OpAdd.value, local(0), arithmetic(S.OpMultiply.value, local(2), integer(3)))));
const foldRef = ref('Data.FoldableWithIndex', 'foldlWithIndexImpl');
const consumer = (seed = integer(17), cb = foldCallback, values = local(4, array)) => lambda([4], [array], C.Int.value, apply(foldRef, [cb, seed, values], C.Int.value));
const mapper = lambda([10], [C.Int.value], C.Int.value, arithmetic(S.OpAdd.value, local(10), integer(1)));
const mapProducer = (callback = mapper, owner = 'Data.Functor', values = local(20, array)) => apply(ref(owner, 'arrayMap'), [callback, values], array);
const concatProducer = (left = local(20, array), right = local(21, array)) => apply(ref('Data.Semigroup', 'concatArray'), [left, right], array);
const factoryInfo = name => ({ arrayFold: Nothing.value, codePointFold: Nothing.value, reference: new SymbolReference(generated('Factories')(`Call_Factories_${name}`)), fArgs: [new G.TypeNativeArray(G.TypeInt64.value)], fRet: new G.TypeNativeArray(G.TypeInt64.value), arity: 1 });
const factories = insert(ordString)('Factories.right')(factoryInfo('right'))(insert(ordString)('Factories.left')(factoryInfo('left'))(emptyMap));
const metadata = () => withReboxFields({ elidedCtors: emptySet, ctorTypes: emptyMap, pointerAdtPaths: emptyMap, pointerAdtNodes: emptySet,
    pointerAdtLeaves: emptyMap, enumAdts: emptySet, enumCtors: emptySet, globalTypes: emptyMap, globalFunctions: factories, classDeclsFields: emptyMap });
const module = bindings => ({ name: 'ArrayProducerFixture', bindings: [{ recursive: false, bindings }], comments: [], imports: emptySet,
    exports: emptySet, reExports: emptySet, dataTypes: emptyMap, dataDecls: [], classDecls: [], foreign: emptyMap, implementations: emptyMap, directives: emptyMap });

export function generatedFixture() {
    const consume = expression => apply(ref('ArrayProducerFixture', 'consume'), [expression], C.Int.value);
    return CodeGen.translate(metadata())(module([
        new Tuple('consume', consumer()),
        new Tuple('mapped', lambda([20], [array], C.Int.value, consume(mapProducer()))),
        new Tuple('joined', lambda([20, 21], [array, array], C.Int.value, consume(concatProducer()))),
        new Tuple('factoryJoined', lambda([20, 21], [array, array], C.Int.value, consume(concatProducer(
            apply(ref('Factories', 'left'), [local(20, array)], array), apply(ref('Factories', 'right'), [local(21, array)], array))))),
    ]));
}

test('fold descriptors retain closed callbacks, seed and exact primitive annotations', () => {
    assert.ok(Fold.describe(consumer()) instanceof Just);
    const reject = value => assert.deepEqual(Fold.describe(value), Nothing.value);
    reject(consumer(typed(C.String.value, integer(0))));
    reject(consumer(integer(0), typed(new C.Func([C.Int.value, C.Int.value, C.Int.value], C.Int.value),
        n(new S.Abs([new Tuple(new Just('arg'), 0)], lambda([1, 2], [C.String.value, C.Int.value], C.Int.value, local(1)))))));
    reject(consumer(integer(0), lambda([0, 1, 2], [C.Int.value, C.Int.value, C.Int.value], C.Int.value,
        typed(C.Int.value, typed(C.String.value, arithmetic(S.OpAdd.value, local(1), local(2)))))));
    reject(consumer(integer(0), foldCallback, typed(new C.Array(C.Number.value), local(4, array))));
    reject(consumer(integer(0), lambda([0, 1, 2], [C.Int.value, C.Int.value, C.Int.value], C.Int.value, local(99))));
    reject(consumer(integer(0), lambda([0, 1, 2], [C.Int.value, C.Int.value, C.Int.value], C.Int.value, apply(ref('Foreign', 'unknown'), [local(2)], C.Int.value))));
});

test('fusion rejects staged or foreign producers and callbacks with captures', () => {
    const descriptor = Fold.describe(consumer()).value0;
    const never = () => { throw Error('rejected candidate must not translate'); };
    const reject = value => assert.deepEqual(Fusion.emit(never)({ recVars: [] })(0)(descriptor)(analyze([])(value)), Nothing.value);
    reject(mapProducer(mapper, 'Trap'));
    reject(mapProducer(lambda([10], [C.Int.value], C.Int.value, arithmetic(S.OpAdd.value, local(10), local(99)))));
    reject(mapProducer(ref('Foreign', 'mapper')));
    reject(apply(n(new S.UncurriedApp(ref('Data.Functor', 'arrayMap'), [])), [mapper, local(20, array)], array));
    reject(apply(n(new S.UncurriedApp(ref('Data.Semigroup', 'concatArray'), [])), [local(20, array), local(21, array)], array));
});

test('generated map and append consumers use inline loops without intermediate arrays', () => {
    const source = generatedFixture();
    assert.match(source, /arrayProducer_fold_/);
    assert.doesNotMatch(source.slice(source.indexOf('func Call_ArrayProducerFixture_consume(')), /make\(\[\]int64/);
    assert.doesNotMatch(source, /Get_Data_Functor_arrayMap|Get_Data_Semigroup_concatArray/);
});

test('generated fused consumers preserve seed, order, wide multiply semantics and input aliases', t => {
    const directory = mkdtempSync(join(tmpdir(), 'gopurs-producer-fold-'));
    t.after(() => rmSync(directory, { recursive: true, force: true }));
    mkdirSync(join(directory, 'gopurs_runtime'));
    writeFileSync(join(directory, 'go.mod'), 'module gopurs/output\n\ngo 1.22\n');
    writeFileSync(join(directory, 'gopurs_runtime/runtime.go'), runtimeGoCode);
    writeFileSync(join(directory, 'fixture.go'), generatedFixture().replace('package purescript', 'package main'));
    writeFileSync(join(directory, 'main.go'), `package main
import("fmt";"math";"reflect";r "gopurs/output/gopurs_runtime")
var factoryTrace []string
var factorySaved []int64
func Call_Factories_left(xs []int64)[]int64{factoryTrace=append(factoryTrace,"left");factorySaved=xs;return xs}
func Call_Factories_right(ys []int64)[]int64{factoryTrace=append(factoryTrace,"right");if len(factorySaved)>0{factorySaved[0]=r.IntAdd(factorySaved[0],17)};return ys}
func expected(xs []int64)int64{acc:=int64(17);for i,v:=range xs{acc=r.IntSub(r.IntMul(acc,3),r.IntAdd(int64(i),r.IntMul(v,3)))};return acc}
func main(){checks:=0;values:=[]int64{0,1,-1,math.MaxInt32,math.MinInt32,1<<52,(1<<52)+1,-(1<<52)-1,math.MaxInt64,math.MinInt64};for n:=0;n<37;n++{for offset:=range values{xs:=make([]int64,n);for i:=range xs{xs[i]=values[(i+offset)%len(values)]};before:=make([]int64,n);copy(before,xs);mapped:=make([]int64,n);for i,v:=range xs{mapped[i]=r.IntAdd(v,1)}
if got,want:=Call_ArrayProducerFixture_mapped(xs),expected(mapped);got!=want{panic(fmt.Sprint("map",n,offset,got,want))};checks++
for _,split:=range []int{0,n/2,n}{if got,want:=Call_ArrayProducerFixture_joined(xs[:split],xs[split:]),expected(xs);got!=want{panic(fmt.Sprint("append",n,split,got,want))};checks++};if !reflect.DeepEqual(xs,before){panic("input mutation")}
aliased:=make([]int64,n);copy(aliased,xs);wantAlias:=make([]int64,n);copy(wantAlias,xs);if n>0{wantAlias[0]=r.IntAdd(wantAlias[0],17)};wantJoined:=append(append([]int64{},wantAlias...),wantAlias...);factoryTrace=nil
if got,want:=Call_ArrayProducerFixture_factoryJoined(aliased,aliased),expected(wantJoined);got!=want{panic(fmt.Sprint("factory alias",n,offset,got,want))};if !reflect.DeepEqual(aliased,wantAlias)||!reflect.DeepEqual(factoryTrace,[]string{"left","right"}){panic("factory phase order or alias")};checks++
}};fmt.Println(checks,"fused producer controls passed")}
`);
    const result = spawnSync('go', ['run', '.'], { cwd: directory, encoding: 'utf8', env: { ...process.env, GOWORK: 'off' } });
    assert.equal(result.status, 0, result.stdout + result.stderr);
    assert.equal(result.stdout.trim(), '1850 fused producer controls passed');
});
