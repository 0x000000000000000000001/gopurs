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
import { emit } from '../output/Gopurs.ObjectTraverse/index.js';
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
const target = new Just({ mbMod: new Just('Data.TraversableWithIndex'), name: 'traverseWithIndex' });
const args = [variable('Foreign.Object', 'traversableWithIndexObject'),
    variable('Data.Either', 'applicativeEither'), variable('Probe', 'callback'), variable('Probe', 'object')];

test('Object traversal fusion only accepts the standard Object and Either dictionaries', () => {
    const forbidden = () => { throw new Error('unexpected translation'); };
    const reject = (candidate, supplied) => assert.deepEqual(emit(forbidden)(context())(0)(candidate)(supplied), Nothing.value);
    reject(Nothing.value, args);
    reject(new Just({ mbMod: new Just('Data.TraversableWithIndex'), name: 'traverseWithIndexDefault' }), args);
    reject(new Just({ mbMod: new Just('Other'), name: 'traverseWithIndex' }), args);
    reject(new Just({ mbMod: Nothing.value, name: 'traverseWithIndex' }), args);
    reject(new Just({ mbMod: new Just('Data.TraversableWithIndex'), name: 'forWithIndex' }), args);
    reject(target, args.slice(0, 1));
    reject(target, [...args, args[3]]);
    reject(target, [variable('Other', 'traversableWithIndexObject'), ...args.slice(1)]);
    reject(target, [variable('Data.TraversableWithIndex', 'traversableWithIndexArray'), ...args.slice(1)]);
    reject(target, [args[0], variable('Data.Maybe', 'applicativeMaybe'), ...args.slice(2)]);
    reject(target, [args[0], variable('Other', 'applicativeEither'), ...args.slice(2)]);
    reject(target, [args[0], new TcoExpr(null, new S.Local(new Just('dictionary'), 0)), ...args.slice(2)]);

});

function generatedFixture() {
    const neutral = syntax => new NeutralExpr(syntax);
    const typed = (type, expression) => neutral(new S.Typed(type, expression));
    const ref = (module, name) => neutral(new S.Var(new C.Qualified(new Just(module), name)));
    const either = new C.ADT('Data.Either.Either', ['Data', 'Either', 'Either'], [C.Any.value, C.Any.value]);
    const object = new C.ADT('Foreign.Object.Object', ['Foreign', 'Object', 'Object'], [C.Any.value]);
    const callback = new C.Func([C.String.value, C.Any.value], either);
    const local = (name, level, type) => typed(type, neutral(new S.Local(new Just(name), level)));
    const binding = count => {
        const remaining = count === 2 ? [callback, object] : count === 3 ? [object] : [];
        const resultType = remaining.length ? new C.Func(remaining, either) : either;
        const supplied = count === 2 ? [] : count === 3 ? [callback] : [callback, object];
        const call = typed(resultType, neutral(new S.App(ref('Data.TraversableWithIndex', 'traverseWithIndex'), [
            ref('Foreign.Object', 'traversableWithIndexObject'), ref('Data.Either', 'applicativeEither'),
            ...(count >= 3 ? [local('callback', 0, callback)] : []),
            ...(count === 4 ? [local('object', 1, object)] : []),
        ])));
        const body = supplied.length ? typed(new C.Func(supplied, resultType), neutral(new S.Abs([
            new Tuple(new Just('callback'), 0), ...(count === 4 ? [new Tuple(new Just('object'), 1)] : []),
        ], call))) : call;
        return new Tuple(`stage${count}`, body);
    };
    return CodeGen.translate(metadata)({
        name: 'ObjectEitherFixture', bindings: [2, 3, 4].map(count => ({ recursive: false, bindings: [binding(count)] })),
        comments: [], imports: emptySet, exports: emptySet, reExports: emptySet,
        dataTypes: emptyMap, dataDecls: [], classDecls: [], foreign: emptyMap,
        implementations: emptyMap, directives: emptyMap,
    });
}

test('generated Object loops preserve strict callbacks, Go key order, first error and fresh storage', t => {
    const declarations = [2, 3, 4].map(count => {
        const translate = _context => nextId => expression => {
            const name = expression.value1.value0.value1;
            const label = name === 'traversableWithIndexObject' ? 'traversal' : name === 'applicativeEither' ? 'applicative' : name;
            const value = ['traversal', 'applicative'].includes(label) ? 'gopurs_runtime.Value{}' : label;
            return { stmts: new StmtLeaf(Go.rawGo(`markStatement("${label}")`)),
                expr: Go.rawGo(`markExpression("${label}", ${value})`), exprType: Go.TypeValue.value, nextId };
        };
        const result = emit(translate)(context())(0)(target)(args.slice(0, count));
        assert.ok(result instanceof Just);
        const body = [...flattenStmts(result.value0.stmts), new Go.GoReturn(result.value0.expr)].map(printGoExpr).join('\n');
        const params = count === 2 ? '' : count === 3 ? 'callback gopurs_runtime.Value' : 'callback, object gopurs_runtime.Value';
        return `func stage${count}(${params}) gopurs_runtime.Value {\n${body}\n}`;
    });
    const fixture = generatedFixture();
    assert.match(fixture, /traverseObjectEither_output_/);
    assert.match(fixture, /"sort"/, 'generated Go must own its sort import');
    assert.doesNotMatch(fixture, /Get_Data_TraversableWithIndex_traverseWithIndex\(/);
    const directory = mkdtempSync(join(tmpdir(), 'gopurs-object-traverse-either-'));
    t.after(() => rmSync(directory, { recursive: true, force: true }));
    mkdirSync(join(directory, 'gopurs_runtime'));
    writeFileSync(join(directory, 'go.mod'), 'module gopurs/output\n\ngo 1.22\n');
    writeFileSync(join(directory, 'gopurs_runtime/runtime.go'), runtimeGoCode);
    writeFileSync(join(directory, 'fixture.go'), fixture.replace('package purescript', 'package main'));
    // Use the actual Go backend's fold implementation as the ordering reference.
    writeFileSync(join(directory, 'object.go'), readFileSync(
        new URL('../../gopurs-foreign-object/src/Foreign/Object.go', import.meta.url), 'utf8',
    ).replace('package Object', 'package main'));
    writeFileSync(join(directory, 'main.go'), `package main
import ("fmt"; "reflect"; "sort"; "unsafe"; "gopurs/output/gopurs_runtime")
type Constructor_Data_Either_Left[A,B any] struct { V0 A }
type Constructor_Data_Either_Right[A,B any] struct { V0 B }
const leftTag = ${hashString('Data_Data_Either_Left')}
const rightTag = ${hashString('Data_Data_Either_Right')}
var trace []string
func markStatement(name string) { trace=append(trace,name+":statement") }
func markExpression[T any](name string,value T) T { trace=append(trace,name+":expression");return value }
func Get_Foreign_Object_traversableWithIndexObject() gopurs_runtime.Value {return gopurs_runtime.Value{}}
func Get_Data_Either_applicativeEither() gopurs_runtime.Value {return gopurs_runtime.Value{}}
func right(v gopurs_runtime.Value) gopurs_runtime.Value {return gopurs_runtime.Value{Type:9,IntVal:rightTag,UnsafePtr:unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value,gopurs_runtime.Value]{v})}}
func left(v gopurs_runtime.Value) gopurs_runtime.Value {return gopurs_runtime.Value{Type:9,IntVal:leftTag,UnsafePtr:unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value,gopurs_runtime.Value]{v})}}
func payload(v gopurs_runtime.Value)gopurs_runtime.Value {if v.IntVal==leftTag{return (*Constructor_Data_Either_Left[gopurs_runtime.Value,gopurs_runtime.Value])(v.UnsafePtr).V0};return (*Constructor_Data_Either_Right[gopurs_runtime.Value,gopurs_runtime.Value])(v.UnsafePtr).V0}
func original(callback gopurs_runtime.Value,values map[string]any)gopurs_runtime.Value {
    bind:=func(acc any)func(any)any{return func(f any)any{return f.(func(any)any)(acc)}}
    step:=func(acc any)func(string)func(any)any{return func(key string)func(any)any{return func(value any)any {
        prior:=acc.(gopurs_runtime.Value)
        result:=gopurs_runtime.Apply2(callback,gopurs_runtime.Str(key),gopurs_runtime.Box(value))
        if prior.IntVal==leftTag{return prior};if result.IntVal==leftTag{return result}
        copied:=_CopyST(gopurs_runtime.UnboxObject(payload(prior)))(nil).(map[string]any)
        copied[key]=payload(result);return right(gopurs_runtime.Any(copied))
    }}}
    return _FoldM(bind)(step)(right(gopurs_runtime.Any(map[string]any{})))(values).(gopurs_runtime.Value)
}
func normalize(result gopurs_runtime.Value)string {
    if result.IntVal==leftTag{return "Left:"+payload(result).StrVal()}
    values:=gopurs_runtime.UnboxObject(payload(result));keys:=[]string{};for k:=range values{keys=append(keys,k)};sort.Strings(keys)
    out:=[]string{};for _,key:=range keys{out=append(out,fmt.Sprintf("%s=%d",key,gopurs_runtime.Box(values[key]).IntVal))};return fmt.Sprint(out)
}
${declarations.join('\n')}
func main(){
    keys:=[]string{"2","10","a","__proto__","é","😀","01","0"};cases:=0
    for n:=0;n<=len(keys);n++{for bad:=-1;bad<n;bad++{
        values:=map[string]any{};before:=map[string]any{};for i,key:=range keys[:n]{values[key]=int64(i);before[key]=int64(i)}
        callback:=gopurs_runtime.Func2(func(key,value gopurs_runtime.Value)gopurs_runtime.Value{
            trace=append(trace,"call:"+key.StrVal());if bad>=0&&value.IntVal>=int64(bad){return left(gopurs_runtime.Str(key.StrVal()))};return right(gopurs_runtime.Int(value.IntVal*2))
        })
        trace=nil;expected:=normalize(original(callback,values));calls:=append([]string{},trace...)
        input:=gopurs_runtime.Any(values)
        probes:=[]func()gopurs_runtime.Value{
            func()gopurs_runtime.Value{return gopurs_runtime.Apply2(stage2(),callback,input)},
            func()gopurs_runtime.Value{return gopurs_runtime.Apply(stage3(callback),input)},
            func()gopurs_runtime.Value{return stage4(callback,input)},
            func()gopurs_runtime.Value{return gopurs_runtime.Apply2(Get_ObjectEitherFixture_stage2(),callback,input)},
            func()gopurs_runtime.Value{return gopurs_runtime.Apply2(Get_ObjectEitherFixture_stage3(),callback,input)},
            func()gopurs_runtime.Value{return gopurs_runtime.Apply2(Get_ObjectEitherFixture_stage4(),callback,input)},
        }
        for mode,probe:=range probes{
            trace=[]string{};actual:=normalize(probe());prefix:=[]string{}
            if mode<3{labels:=[]string{"traversal","applicative"};if mode>=1{labels=append(labels,"callback")};if mode==2{labels=append(labels,"object")};for _,label:=range labels{prefix=append(prefix,label+":statement",label+":expression")}}
            wanted:=append(prefix,calls...);if actual!=expected||!reflect.DeepEqual(trace,wanted){panic(fmt.Sprintf("mode=%d n=%d bad=%d got=%s want=%s trace=%v wanted=%v",mode,n,bad,actual,expected,trace,wanted))};if !reflect.DeepEqual(values,before){panic("input changed")};cases++
        }
    }}
    callback:=gopurs_runtime.Func2(func(k,v gopurs_runtime.Value)gopurs_runtime.Value{return right(v)})
    for _,inputMap:=range []map[string]any{{},{"a":int64(7)}} {
        input:=gopurs_runtime.Any(inputMap);binary:=stage2();unary:=gopurs_runtime.Apply(binary,callback)
        first:=gopurs_runtime.Apply(unary,input);second:=gopurs_runtime.Apply(unary,input)
        expected:=normalize(second);gopurs_runtime.UnboxObject(payload(first))["a"]=gopurs_runtime.Int(99)
        if normalize(second)!=expected{panic("results alias")};if v,ok:=inputMap["a"];ok&&v!=int64(7){panic("result aliases input")}
    }
    fmt.Printf("%d strict object comparisons; fresh empty and reused outputs: ok\\n",cases)
}
`);
    const result = spawnSync('go', ['run', '.'], { cwd: directory, encoding: 'utf8', timeout: 30_000,
        env: { ...process.env, GOWORK: 'off' } });
    assert.ifError(result.error);
    assert.equal(result.status, 0, result.stdout + result.stderr);
    assert.match(result.stdout, /^270 strict object comparisons; fresh empty and reused outputs: ok\n$/);
});
