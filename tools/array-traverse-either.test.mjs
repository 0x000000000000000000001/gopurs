import { withReboxFields } from './codegen-metadata.mjs';
import assert from 'node:assert/strict';
import { spawnSync } from 'node:child_process';
import { mkdirSync, mkdtempSync, readFileSync, rmSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join } from 'node:path';
import test from 'node:test';
import { empty as emptyMap, insert } from '../output/Data.Map/index.js';
import { Just, Nothing } from '../output/Data.Maybe/index.js';
import { ordString } from '../output/Data.Ord/index.js';
import { empty as emptySet } from '../output/Data.Set/index.js';
import { Tuple } from '../output/Data.Tuple/index.js';
import * as Ref from '../output/Effect.Ref/index.js';
import { emit } from '../output/Gopurs.ArrayTraverse/index.js';
import * as CodeGen from '../output/Gopurs.CodeGen/index.js';
import { StmtEmpty, StmtLeaf, flattenStmts } from '../output/Gopurs.ExprContext/index.js';
import * as Go from '../output/Gopurs.GoAst/index.js';
import { printGoExpr } from '../output/Gopurs.Printer/index.js';
import { runtimeGoCode } from '../output/Gopurs.Runtime/index.js';
import { TcoExpr } from '../output/PureScript.Backend.Optimizer.Codegen.Tco/index.js';
import * as C from '../output/PureScript.Backend.Optimizer.CoreFn/index.js';
import { hashString } from '../output/PureScript.Backend.Optimizer.FfiSupport/index.js';
import { NeutralExpr } from '../output/PureScript.Backend.Optimizer.Semantics/index.js';
import * as S from '../output/PureScript.Backend.Optimizer.Syntax/index.js';

const constructors = insert(ordString)('Data.Either.Left')({ vars: ['a', 'b'], fields: [new C.TypeVar('a')] })(
    insert(ordString)('Data.Either.Right')({ vars: ['a', 'b'], fields: [new C.TypeVar('b')] })(emptyMap));
const metadata = withReboxFields({
    elidedCtors: emptySet, ctorTypes: constructors, pointerAdtPaths: emptyMap,
    pointerAdtNodes: emptySet, pointerAdtLeaves: emptyMap, enumAdts: emptySet,
    enumCtors: emptySet, globalTypes: emptyMap, globalFunctions: emptyMap,
    classDeclsFields: emptyMap,
});
const context = () => ({
    metadata, codegenStateRef: Ref.new({ declarations: [], globalId: 0, reboxPairs: emptySet })(),
    depth: 0, modNameStr: 'Probe', recVars: [], moduleFunctions: emptyMap, bound: emptyMap,
    tcoIdent: Nothing.value, loopCtx: [], options: { isTail: false, inEffectBlock: false },
    mbExpectedExprType: Nothing.value,
});
const variable = (module, name) => new TcoExpr(null, new S.Var(new C.Qualified(new Just(module), name)));
const target = new Just({ mbMod: new Just('Data.TraversableWithIndex'), name: 'traverseWithIndexDefault' });
const args = [variable('Data.TraversableWithIndex', 'traversableWithIndexArray'),
    variable('Data.Either', 'applicativeEither'), variable('Probe', 'callback'), variable('Probe', 'array')];

test('Either traversal fusion requires the standard qualified Array and Either dictionaries', () => {
    const forbidden = () => { throw new Error('unexpected translation'); };
    const reject = (candidate, supplied) => assert.deepEqual(
        emit(forbidden)(context())(0)(candidate)(supplied), Nothing.value);
    for (const candidate of [Nothing.value,
        new Just({ mbMod: new Just('Other'), name: 'traverseWithIndexDefault' }),
        new Just({ mbMod: Nothing.value, name: 'traverseWithIndexDefault' }),
        new Just({ mbMod: new Just('Data.TraversableWithIndex'), name: 'forWithIndex' })]) {
        reject(candidate, args);
    }
    reject(target, args.slice(0, 2));
    reject(target, [...args, args[3]]);
    reject(target, [variable('Other', 'traversableWithIndexArray'), ...args.slice(1)]);
    reject(target, [args[0], variable('Other', 'applicativeEither'), ...args.slice(2)]);
    reject(target, [args[0], variable('Data.Maybe', 'applicativeMaybe'), ...args.slice(2)]);
    reject(target, [args[0], new TcoExpr(null, new S.Local(new Just('dictionary'), 0)), ...args.slice(2)]);
});

function generatedFixture() {
    const neutral = syntax => new NeutralExpr(syntax);
    const typed = (type, expression) => neutral(new S.Typed(type, expression));
    const ref = (module, name) => neutral(new S.Var(new C.Qualified(new Just(module), name)));
    const either = new C.ADT('Data.Either.Either', ['Data', 'Either', 'Either'], [C.Any.value, C.Any.value]);
    const array = new C.Array(C.Any.value);
    const callback = new C.Func([C.Int.value, C.Any.value], either);
    const local = (name, level, type) => typed(type, neutral(new S.Local(new Just(name), level)));
    const call = (partial) => typed(partial ? new C.Func([array], either) : either, neutral(new S.App(
        ref('Data.TraversableWithIndex', 'traverseWithIndexDefault'), [
            ref('Data.TraversableWithIndex', 'traversableWithIndexArray'),
            ref('Data.Either', 'applicativeEither'), local('callback', 0, callback),
            ...(partial ? [] : [local('array', 1, array)]),
        ])));
    const binding = (name, partial) => new Tuple(name,
        typed(new C.Func(partial ? [callback] : [callback, array], partial ? new C.Func([array], either) : either),
            neutral(new S.Abs([new Tuple(new Just('callback'), 0),
                ...(partial ? [] : [new Tuple(new Just('array'), 1)])], call(partial)))));
    return CodeGen.translate(metadata)({
        name: 'EitherTraversalFixture', bindings: [binding('partial', true), binding('direct', false)]
            .map(value => ({ recursive: false, bindings: [value] })),
        comments: [], imports: emptySet, exports: emptySet, reExports: emptySet,
        dataTypes: emptyMap, dataDecls: [], classDecls: [], foreign: emptyMap,
        implementations: emptyMap, directives: emptyMap,
    });
}

test('generated loops match strict indexed traversal, including every callback after the first error', t => {
    const modes = [
        { name: 'directBoxed', type: Go.TypeValue.value, parameter: 'gopurs_runtime.Value', partial: false },
        { name: 'directNative', type: new Go.TypeNativeArray(Go.TypeInt64.value), parameter: '[]int64', partial: false },
        { name: 'partial', type: Go.TypeValue.value, parameter: 'gopurs_runtime.Value', partial: true },
    ];
    const declarations = modes.map(mode => {
        const translate = _context => nextId => expression => {
            const name = expression.value1.value0.value1;
            const label = name === 'traversableWithIndexArray' ? 'traversal'
                : name === 'applicativeEither' ? 'applicative' : name;
            const value = label === 'traversal' || label === 'applicative' ? 'gopurs_runtime.Value{}' : label;
            return {
                stmts: new StmtLeaf(Go.rawGo(`markStatement("${label}")`)),
                expr: Go.rawGo(`markExpression("${label}", ${value})`),
                exprType: label === 'array' ? mode.type : Go.TypeValue.value, nextId,
            };
        };
        const result = emit(translate)(context())(0)(target)(mode.partial ? args.slice(0, 3) : args);
        assert.ok(result instanceof Just);
        const body = [...flattenStmts(result.value0.stmts), new Go.GoReturn(result.value0.expr)].map(printGoExpr).join('\n');
        return `func ${mode.name}(callback gopurs_runtime.Value${mode.partial ? '' : `, array ${mode.parameter}`}) gopurs_runtime.Value {\n${body}\n}`;
    });
    const fixture = generatedFixture();
    assert.match(fixture, /traverseEither_output_/, 'the partial shape used by real decodeArray must hit the intrinsic');
    assert.doesNotMatch(fixture, /Get_Data_TraversableWithIndex_traverseWithIndexDefault\(/);
    const directory = mkdtempSync(join(tmpdir(), 'gopurs-array-traverse-either-'));
    t.after(() => rmSync(directory, { recursive: true, force: true }));
    mkdirSync(join(directory, 'gopurs_runtime'));
    writeFileSync(join(directory, 'go.mod'), 'module gopurs/output\n\ngo 1.22\n');
    writeFileSync(join(directory, 'gopurs_runtime/runtime.go'), runtimeGoCode);
    writeFileSync(join(directory, 'fixture.go'), fixture.replace('package purescript', 'package main'));
    writeFileSync(join(directory, 'traversable.go'), readFileSync(
        new URL('../../gopurs-foldable-traversable/src/Data/Traversable.go', import.meta.url), 'utf8',
    ).replace('package Data_Traversable', 'package main'));
    writeFileSync(join(directory, 'main.go'), `package main
import ("fmt"; "reflect"; "unsafe"; "gopurs/output/gopurs_runtime")
type Constructor_Data_Either_Left[A,B any] struct { V0 A }
type Constructor_Data_Either_Right[A,B any] struct { V0 B }
const leftTag = ${hashString('Data_Data_Either_Left')}
const rightTag = ${hashString('Data_Data_Either_Right')}
var trace []string
func markStatement(name string) { trace=append(trace,name+":statement") }
func markExpression[T any](name string,value T) T { trace=append(trace,name+":expression");return value }
func Get_Data_TraversableWithIndex_traversableWithIndexArray() gopurs_runtime.Value { return gopurs_runtime.Value{} }
func Get_Data_Either_applicativeEither() gopurs_runtime.Value { return gopurs_runtime.Value{} }
func right(v gopurs_runtime.Value) gopurs_runtime.Value { return gopurs_runtime.Value{Type:9,IntVal:rightTag,UnsafePtr:unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value,gopurs_runtime.Value]{v})} }
func left(v gopurs_runtime.Value) gopurs_runtime.Value { return gopurs_runtime.Value{Type:9,IntVal:leftTag,UnsafePtr:unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value,gopurs_runtime.Value]{v})} }
func payload(v gopurs_runtime.Value) gopurs_runtime.Value {
    if v.IntVal==leftTag { return (*Constructor_Data_Either_Left[gopurs_runtime.Value,gopurs_runtime.Value])(v.UnsafePtr).V0 }
    return (*Constructor_Data_Either_Right[gopurs_runtime.Value,gopurs_runtime.Value])(v.UnsafePtr).V0
}
func original(callback gopurs_runtime.Value, values []gopurs_runtime.Value) gopurs_runtime.Value {
    mapped:=make([]gopurs_runtime.Value,len(values))
    for i,value:=range values { mapped[i]=gopurs_runtime.Apply2(callback,gopurs_runtime.Int(int64(i)),value) }
    apply:=func(a,b gopurs_runtime.Value)gopurs_runtime.Value {
        if a.IntVal==leftTag{return a};if b.IntVal==leftTag{return b};return right(gopurs_runtime.Apply(payload(a),payload(b)))
    }
    mapFn:=func(f func(gopurs_runtime.Value)gopurs_runtime.Value,a gopurs_runtime.Value)gopurs_runtime.Value {
        if a.IntVal==leftTag{return a};return right(f(payload(a)))
    }
    concat:=func(a gopurs_runtime.Value)func(gopurs_runtime.Value)gopurs_runtime.Value {
        return func(b gopurs_runtime.Value)gopurs_runtime.Value {
            av,bv:=*(*[]gopurs_runtime.Value)(a.UnsafePtr),*(*[]gopurs_runtime.Value)(b.UnsafePtr)
            result:=make([]gopurs_runtime.Value,len(av)+len(bv));copy(result,av);copy(result[len(av):],bv)
            return gopurs_runtime.Array(result)
        }
    }
    return TraverseArrayImpl(apply,mapFn,right,concat,func(v gopurs_runtime.Value)gopurs_runtime.Value{return v},mapped)
}
func normalize(result gopurs_runtime.Value) string {
    if result.IntVal==leftTag { return "Left:"+payload(result).StrVal() }
    values:=*(*[]gopurs_runtime.Value)(payload(result).UnsafePtr);out:=make([]int64,len(values))
    for i,v:=range values {out[i]=v.IntVal};return fmt.Sprint(out)
}
${declarations.join('\n')}
func main() {
    cases:=0
    for n:=0;n<=65;n++ { for bad:=-1;bad<n;bad++ {
        values:=make([]gopurs_runtime.Value,n);native:=make([]int64,n)
        for i:=range values {values[i]=gopurs_runtime.Int(int64(i));native[i]=int64(i)}
        before:=append([]gopurs_runtime.Value{},values...)
        callback:=gopurs_runtime.Func2(func(index,value gopurs_runtime.Value)gopurs_runtime.Value {
            trace=append(trace,fmt.Sprintf("call:%d",index.IntVal))
            if bad>=0 && int(index.IntVal)>=bad { return left(gopurs_runtime.Str(fmt.Sprintf("bad-%d",index.IntVal))) }
            return right(gopurs_runtime.Int(value.IntVal*2+index.IntVal))
        })
        trace=nil;expected:=normalize(original(callback,values));expectedCalls:=append([]string{},trace...)
        probes:=[]func()gopurs_runtime.Value{
            func()gopurs_runtime.Value{return directBoxed(callback,gopurs_runtime.Array(values))},
            func()gopurs_runtime.Value{return directNative(callback,native)},
            func()gopurs_runtime.Value{return gopurs_runtime.Apply(partial(callback),gopurs_runtime.Array(values))},
            func()gopurs_runtime.Value{return gopurs_runtime.Apply(Call_EitherTraversalFixture_partial(callback),gopurs_runtime.Array(values))},
            func()gopurs_runtime.Value{return gopurs_runtime.Apply2(Get_EitherTraversalFixture_direct(),callback,gopurs_runtime.Array(values))},
        }
        for mode,probe:=range probes {
            trace=[]string{};actual:=normalize(probe());prefix:=[]string{}
            if mode<3 {
                labels:=[]string{"traversal","applicative","callback"};if mode<2 {labels=append(labels,"array")}
                for _,label:=range labels {prefix=append(prefix,label+":statement",label+":expression")}
            }
            expectedTrace:=append(prefix,expectedCalls...)
            if actual!=expected || !reflect.DeepEqual(trace,expectedTrace) {panic(fmt.Sprintf("mode=%d n=%d bad=%d got=%s want=%s trace=%v wantTrace=%v",mode,n,bad,actual,expected,trace,expectedTrace))}
            if !reflect.DeepEqual(values,before) {panic("input array changed")}
            cases++
        }
    } }
    // Reusing a partial application must allocate fresh result storage each time.
    callback:=gopurs_runtime.Func2(func(i,v gopurs_runtime.Value)gopurs_runtime.Value{return right(v)})
    reused:=partial(callback);input:=gopurs_runtime.Array([]gopurs_runtime.Value{gopurs_runtime.Int(7)})
    first:=gopurs_runtime.Apply(reused,input);second:=gopurs_runtime.Apply(reused,input)
    (*(*[]gopurs_runtime.Value)(payload(first).UnsafePtr))[0]=gopurs_runtime.Int(99)
    if normalize(second)!="[7]" || (*(*[]gopurs_runtime.Value)(input.UnsafePtr))[0].IntVal!=7 {panic("result aliases input or another call")}
    fmt.Printf("%d strict traversal comparisons; fresh reusable closures: ok\\n",cases)
}
`);
    const result = spawnSync('go', ['run', '.'], {
        cwd: directory, encoding: 'utf8', timeout: 30_000,
        env: { ...process.env, GOWORK: 'off' },
    });
    assert.ifError(result.error);
    assert.equal(result.status, 0, result.stdout + result.stderr);
    assert.match(result.stdout, /^11055 strict traversal comparisons; fresh reusable closures: ok\n$/);
});
