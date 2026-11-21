import assert from 'node:assert/strict';
import { spawnSync } from 'node:child_process';
import { mkdirSync, mkdtempSync, readFileSync, rmSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join } from 'node:path';
import test from 'node:test';
import { empty as emptyMap, insert } from '../output/Data.Map/index.js';
import { Just, Nothing } from '../output/Data.Maybe/index.js';
import { ordString } from '../output/Data.Ord/index.js';
import { empty as emptySet, singleton } from '../output/Data.Set/index.js';
import { Tuple } from '../output/Data.Tuple/index.js';
import * as Ref from '../output/Effect.Ref/index.js';
import { emit } from '../output/Gopurs.ArrayTraverse/index.js';
import * as CodeGen from '../output/Gopurs.CodeGen/index.js';
import { StmtEmpty, StmtLeaf, flattenStmts } from '../output/Gopurs.ExprContext/index.js';
import * as Go from '../output/Gopurs.GoAst/index.js';
import { printGoExpr } from '../output/Gopurs.Printer/index.js';
import { runtimeGoCode } from '../output/Gopurs.Runtime.Source/index.js';
import { TcoExpr } from '../output/PureScript.Backend.Optimizer.Codegen.Tco/index.js';
import * as C from '../output/PureScript.Backend.Optimizer.CoreFn/index.js';
import { hashString } from '../output/PureScript.Backend.Optimizer.FfiSupport/index.js';
import { NeutralExpr } from '../output/PureScript.Backend.Optimizer.Semantics/index.js';
import * as S from '../output/PureScript.Backend.Optimizer.Syntax/index.js';
import { withReboxFields } from './codegen-metadata.mjs';

const put = (key, value, map) => insert(ordString)(key)(value)(map);
const metadata = withReboxFields({
    elidedCtors: emptySet,
    ctorTypes: put('Data.Maybe.Just', { vars: ['a'], fields: [new C.TypeVar('a')] },
        put('Data.Maybe.Nothing', { vars: ['a'], fields: [] }, emptyMap)),
    pointerAdtPaths: emptyMap, pointerAdtNodes: singleton('Data_Data_Maybe_Just'),
    pointerAdtLeaves: put('Data_Data_Maybe_Nothing', { nodeBaseStruct: 'Data_Data_Maybe_Just', nodeCtor: 'Just' }, emptyMap),
    enumAdts: emptySet, enumCtors: emptySet, globalTypes: emptyMap,
    globalFunctions: emptyMap, classDeclsFields: emptyMap,
});
const context = () => ({ metadata, codegenStateRef: Ref.new({ declarations: [], globalId: 0, reboxPairs: emptySet })(),
    depth: 0, modNameStr: 'Probe', recVars: [], moduleFunctions: emptyMap, bound: emptyMap,
    tcoIdent: Nothing.value, loopCtx: [], options: { isTail: false, inEffectBlock: false },
    mbExpectedExprType: Nothing.value, mbExpectedGoType: Nothing.value });
const variable = (module, name) => new TcoExpr(null, new S.Var(new C.Qualified(new Just(module), name)));
const traversal = variable('Data.Traversable', 'traversableArray');
const target = new TcoExpr(null, new S.Accessor(traversal, new S.GetProp('traverse')));
const applicative = variable('Data.Maybe', 'applicativeMaybe');
const args = [applicative, variable('Probe', 'callback'), variable('Probe', 'array')];
const translate = _context => nextId => _expression => ({ stmts: StmtEmpty.value,
    expr: Go.rawGo('placeholder'), exprType: Go.TypeValue.value, nextId });

test('Maybe traversal accepts only canonical dictionaries, methods and exact arities', () => {
    const accept = (fn, supplied) => assert.ok(emit(translate)(context())(0)(fn)(supplied) instanceof Just);
    const reject = (fn, supplied) => assert.deepEqual(emit(() => { throw Error('unexpected translation'); })(context())(0)(fn)(supplied), Nothing.value);
    for (const count of [1, 2, 3]) accept(target, args.slice(0, count));
    accept(variable('Data.Traversable', 'traverse'), [traversal, ...args]);
    accept(variable('Data.TraversableWithIndex', 'traverseWithIndex'),
        [variable('Data.TraversableWithIndex', 'traversableWithIndexArray'), ...args]);
    reject(target, []); reject(target, [...args, args[2]]);
    reject(target, [variable('Other', 'applicativeMaybe'), ...args.slice(1)]);
    reject(target, [new TcoExpr(null, new S.Local(new Just('dictionary'), 0)), ...args.slice(1)]);
    reject(new TcoExpr(null, new S.Accessor(variable('Other', 'traversableArray'), new S.GetProp('traverse'))), args);
    const method = (module, name, dictModule, dictionary) => new TcoExpr(null,
        new S.App(variable(module, name), [variable(dictModule, dictionary)]));
    const ffi = variable('Data.Traversable', 'traverseArrayImpl');
    const prefix = [method('Control.Apply', 'apply', 'Data.Maybe', 'applyMaybe'),
        method('Data.Functor', 'map', 'Data.Maybe', 'functorMaybe'),
        method('Control.Applicative', 'pure', 'Data.Maybe', 'applicativeMaybe'),
        method('Data.Semigroup', 'append', 'Data.Semigroup', 'semigroupArray')];
    for (const count of [4, 5, 6]) accept(ffi, [...prefix, ...args.slice(1)].slice(0, count));
    for (let index = 0; index < prefix.length; index++) {
        const changed = [...prefix, ...args.slice(1)]; changed[index] = variable('Other', 'method'); reject(ffi, changed);
    }
    reject(ffi, [method('Control.Apply', 'apply', 'Data.Either', 'applyEither'), ...prefix.slice(1), ...args.slice(1)]);
    reject(ffi, [...prefix, ...args.slice(1), args[2]]);
});

function generatedFixture() {
    const neutral = syntax => new NeutralExpr(syntax);
    const typed = (type, expr) => neutral(new S.Typed(type, expr));
    const ref = (module, name) => neutral(new S.Var(new C.Qualified(new Just(module), name)));
    const maybe = new C.ADT('Data.Maybe.Maybe', ['Data', 'Maybe', 'Maybe'], [C.Any.value]);
    const array = new C.Array(C.Any.value), callback = new C.Func([C.Any.value], maybe);
    const local = (name, level, type) => typed(type, neutral(new S.Local(new Just(name), level)));
    const body = typed(maybe, neutral(new S.App(neutral(new S.Accessor(ref('Data.Traversable', 'traversableArray'), new S.GetProp('traverse'))),
        [ref('Data.Maybe', 'applicativeMaybe'), local('callback', 0, callback), local('array', 1, array)])));
    const binding = new Tuple('direct', typed(new C.Func([callback, array], maybe), neutral(new S.Abs(
        [new Tuple(new Just('callback'), 0), new Tuple(new Just('array'), 1)], body))));
    return CodeGen.translate(metadata)({ name: 'MaybeTraversalFixture', bindings: [{ recursive: false, bindings: [binding] }],
        comments: [], imports: emptySet, exports: emptySet, reExports: emptySet, dataTypes: emptyMap,
        dataDecls: [], classDecls: [], foreign: emptyMap, implementations: emptyMap, directives: emptyMap });
}

test('generated Maybe loops preserve strict callbacks, failures, replay and fresh results', t => {
    const modes = [
        { name: 'boxed', type: Go.TypeValue.value, parameter: 'gopurs_runtime.Value', partial: false },
        { name: 'native', type: new Go.TypeNativeArray(Go.TypeInt64.value), parameter: '[]int64', partial: false },
        { name: 'partial', type: Go.TypeValue.value, parameter: '', partial: true },
    ];
    const declarations = modes.map(mode => {
        const translator = _context => nextId => expression => {
            const name = expression.value1.value0.value1;
            const label = name === 'traversableArray' ? 'traversal' : name === 'applicativeMaybe' ? 'applicative' : name;
            const value = ['traversal', 'applicative'].includes(label) ? 'gopurs_runtime.Value{}' : label;
            return { stmts: new StmtLeaf(Go.rawGo(`mark("${label}:statement")`)), expr: Go.rawGo(`capture("${label}:expression", ${value})`),
                exprType: label === 'array' ? mode.type : Go.TypeValue.value, nextId };
        };
        const result = emit(translator)(context())(0)(target)(mode.partial ? args.slice(0, 2) : args);
        assert.ok(result instanceof Just);
        const body = [...flattenStmts(result.value0.stmts), new Go.GoReturn(result.value0.expr)].map(printGoExpr).join('\n');
        return `func ${mode.name}(callback gopurs_runtime.Value${mode.partial ? '' : `, array ${mode.parameter}`}) gopurs_runtime.Value {\n${body}\n}`;
    });
    const fixture = generatedFixture(); assert.match(fixture, /traverseMaybe_output_/);
    assert.doesNotMatch(fixture, /Get_Data_Traversable_traverse\(/);
    const directory = mkdtempSync(join(tmpdir(), 'gopurs-traverse-maybe-'));
    t.after(() => rmSync(directory, { recursive: true, force: true }));
    mkdirSync(join(directory, 'gopurs_runtime'));
    writeFileSync(join(directory, 'go.mod'), 'module gopurs/output\n\ngo 1.22\n');
    writeFileSync(join(directory, 'gopurs_runtime/runtime.go'), runtimeGoCode);
    writeFileSync(join(directory, 'fixture.go'), fixture.replace('package purescript', 'package main'));
    writeFileSync(join(directory, 'traversable.go'), readFileSync(new URL('../../gopurs-foldable-traversable/src/Data/Traversable.go', import.meta.url), 'utf8').replace('package Data_Traversable', 'package main'));
    writeFileSync(join(directory, 'main.go'), `package main
import ("fmt"; "reflect"; "unsafe"; "gopurs/output/gopurs_runtime")
type Constructor_Data_Maybe_Just[A any] struct { Rc uint32; V0 A }
var trace []string
func mark(s string) { trace=append(trace,s) }
func capture[T any](s string,v T)T {mark(s);return v}
func Get_Data_Traversable_traversableArray() gopurs_runtime.Value{return gopurs_runtime.Value{}}
func Get_Data_Maybe_applicativeMaybe() gopurs_runtime.Value{return gopurs_runtime.Value{}}
func just(v gopurs_runtime.Value)gopurs_runtime.Value{return gopurs_runtime.Value{Type:9,IntVal:${hashString('Data_Data_Maybe_Just')},UnsafePtr:unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1,v})}}
func nothing()gopurs_runtime.Value{return gopurs_runtime.Value{Type:9,IntVal:${hashString('Data_Data_Maybe_Just')}}}
func payload(v gopurs_runtime.Value)gopurs_runtime.Value{return (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v.UnsafePtr).V0}
func original(callback gopurs_runtime.Value, values []gopurs_runtime.Value)gopurs_runtime.Value {
    apply:=func(f,a gopurs_runtime.Value)gopurs_runtime.Value{if f.UnsafePtr==nil||a.UnsafePtr==nil{return nothing()};return just(gopurs_runtime.Apply(payload(f),payload(a)))}
    mapFn:=func(f func(gopurs_runtime.Value)gopurs_runtime.Value,a gopurs_runtime.Value)gopurs_runtime.Value{if a.UnsafePtr==nil{return nothing()};return just(f(payload(a)))}
    concat:=func(a gopurs_runtime.Value)func(gopurs_runtime.Value)gopurs_runtime.Value{return func(b gopurs_runtime.Value)gopurs_runtime.Value{
        av,bv:=*(*[]gopurs_runtime.Value)(a.UnsafePtr),*(*[]gopurs_runtime.Value)(b.UnsafePtr);out:=make([]gopurs_runtime.Value,len(av)+len(bv));copy(out,av);copy(out[len(av):],bv);return gopurs_runtime.Array(out)}}
    return TraverseArrayImpl(apply,mapFn,just,concat,func(v gopurs_runtime.Value)gopurs_runtime.Value{return gopurs_runtime.Apply(callback,v)},values)
}
func normalize(v gopurs_runtime.Value)string{if v.UnsafePtr==nil{return "Nothing"};a:=*(*[]gopurs_runtime.Value)(payload(v).UnsafePtr);out:=make([]int64,len(a));for i,v:=range a{out[i]=v.IntVal};return fmt.Sprint(out)}
${declarations.join('\n')}
func main(){
    cases:=0
    for _,n:=range []int{0,1,2,3,4,5,17,65,1000}{for _,bad:=range []int{-1,0,n/2,n-1}{
        values:=make([]gopurs_runtime.Value,n);nativeValues:=make([]int64,n);for i:=range values{values[i]=gopurs_runtime.Int(int64(i));nativeValues[i]=int64(i)}
        before:=append([]gopurs_runtime.Value{},values...)
        callback:=gopurs_runtime.Func(func(v gopurs_runtime.Value)gopurs_runtime.Value{mark(fmt.Sprint("call:",v.IntVal));if int(v.IntVal)==bad{return nothing()};return just(gopurs_runtime.Int(v.IntVal+7))})
        trace=nil;expected:=normalize(original(callback,values));calls:=append([]string{},trace...)
        probes:=[]func()gopurs_runtime.Value{func()gopurs_runtime.Value{return boxed(callback,gopurs_runtime.Array(values))},func()gopurs_runtime.Value{return native(callback,nativeValues)},func()gopurs_runtime.Value{return gopurs_runtime.Apply(partial(callback),gopurs_runtime.Array(values))},func()gopurs_runtime.Value{return gopurs_runtime.Apply2(Get_MaybeTraversalFixture_direct(),callback,gopurs_runtime.Array(values))}}
        for mode,probe:=range probes{trace=[]string{};actual:=normalize(probe());prefix:=[]string{};if mode<3{labels:=[]string{"traversal","applicative","callback"};if mode<2{labels=append(labels,"array")};for _,label:=range labels{prefix=append(prefix,label+":statement",label+":expression")}}
            if actual!=expected||!reflect.DeepEqual(trace,append(prefix,calls...))||!reflect.DeepEqual(values,before){panic(fmt.Sprint("differential failure ",mode,n,bad,actual,expected,trace))};cases++}
    }}
    callback:=gopurs_runtime.Func(func(v gopurs_runtime.Value)gopurs_runtime.Value{return just(v)});reused:=partial(callback);input:=gopurs_runtime.Array([]gopurs_runtime.Value{gopurs_runtime.Int(7)})
    first:=gopurs_runtime.Apply(reused,input);second:=gopurs_runtime.Apply(reused,input);(*(*[]gopurs_runtime.Value)(payload(first).UnsafePtr))[0]=gopurs_runtime.Int(99)
    if normalize(second)!="[7]"||(*(*[]gopurs_runtime.Value)(input.UnsafePtr))[0].IntVal!=7{panic("shared result")}
    for _,run:=range []func(gopurs_runtime.Value,gopurs_runtime.Value)gopurs_runtime.Value{boxed,func(f,a gopurs_runtime.Value)gopurs_runtime.Value{return gopurs_runtime.Apply(partial(f),a)}}{
        func(){defer func(){if recover()!="later callback"{panic("late exception lost")}}();run(gopurs_runtime.Func(func(v gopurs_runtime.Value)gopurs_runtime.Value{if v.IntVal==1{panic("later callback")};return nothing()}),gopurs_runtime.Array([]gopurs_runtime.Value{gopurs_runtime.Int(0),gopurs_runtime.Int(1)}));panic("callback skipped")}()
    }
    fmt.Printf("%d strict Maybe comparisons; replay, fresh storage and exceptions: ok\\n",cases)
}
`);
    const run = spawnSync('go', ['run', '.'], { cwd: directory, encoding: 'utf8', timeout: 60000, env: { ...process.env, GOWORK: 'off' } });
    assert.ifError(run.error); assert.equal(run.status, 0, run.stdout + run.stderr);
    assert.match(run.stdout, /^144 strict Maybe comparisons; replay, fresh storage and exceptions: ok\n$/);
});
