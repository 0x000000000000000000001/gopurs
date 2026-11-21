import assert from 'node:assert/strict';
import test from 'node:test';
import { spawnSync } from 'node:child_process';
import { mkdirSync, mkdtempSync, rmSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join } from 'node:path';
import { empty as emptyMap, insert } from '../output/Data.Map/index.js';
import { Just, Nothing } from '../output/Data.Maybe/index.js';
import { ordString } from '../output/Data.Ord/index.js';
import { empty as emptySet } from '../output/Data.Set/index.js';
import { Tuple } from '../output/Data.Tuple/index.js';
import { NeutralExpr } from '../output/PureScript.Backend.Optimizer.Semantics/index.js';
import * as S from '../output/PureScript.Backend.Optimizer.Syntax/index.js';
import * as C from '../output/PureScript.Backend.Optimizer.CoreFn/index.js';
import * as CodeGen from '../output/Gopurs.CodeGen/index.js';
import * as Go from '../output/Gopurs.GoAst/index.js';
import { SymbolReference, generated } from '../output/Gopurs.GoSymbol/index.js';
import { runtimeGoCode } from '../output/Gopurs.Runtime.Source/index.js';
import { withReboxFields } from './codegen-metadata.mjs';

const metadata = withReboxFields({ elidedCtors: emptySet, ctorTypes: emptyMap,
    pointerAdtPaths: emptyMap, pointerAdtNodes: emptySet, pointerAdtLeaves: emptyMap,
    enumAdts: emptySet, enumCtors: emptySet, globalTypes: emptyMap, classDeclsFields: emptyMap,
    globalFunctions: insert(ordString)('Imported.increment')({reference:new SymbolReference(generated('Imported')('Call_Imported_increment')),
        arity:1,fArgs:[Go.TypeInt64.value],fRet:Go.TypeInt64.value})(emptyMap) });
const ir = syntax => new NeutralExpr(syntax);
const typed = (type, value) => ir(new S.Typed(type,value));
const arrayType = new C.Array(C.Int.value);
const ref = (module,name) => ir(new S.Var(new C.Qualified(new Just(module),name)));
const local = (name,level,type) => typed(type,ir(new S.Local(new Just(name),level)));
const literal = value => typed(C.Int.value,ir(new S.Lit(new C.LitInt(value))));
const value = local('value',1,C.Int.value);
const lambda = (type,body) => typed(new C.Func([C.Int.value],type),ir(new S.Abs([new Tuple(new Just('value'),1)],body)));
const increment = lambda(C.Int.value,typed(C.Int.value,ir(new S.PrimOp(new S.Op2(new S.OpIntNum(S.OpAdd.value),value,literal(1))))));
const positive = lambda(C.Boolean.value,typed(C.Boolean.value,ir(new S.PrimOp(new S.Op2(new S.OpIntOrd(S.OpGte.value),value,literal(0))))));
function fixture() {
    const bindings=[['mapped','Data.Functor','arrayMap',increment],['filtered','Data.Array','filterImpl',positive],
        ['imported','Data.Functor','arrayMap',ref('Imported','increment')],
        ['dynamic','Data.Array','filterImpl',ref('Imported','predicate')]].map(([name,module,fn,callback])=>{
        // The source is boxed, with a concrete Array Int view at this call site.
        const input=typed(arrayType,local('input',0,C.Any.value));
        const call=typed(arrayType,ir(name==='mapped'||name==='imported'
            ?new S.App(ref(module,fn),[callback,input]):new S.UncurriedApp(ref(module,fn),[callback,input])));
        return {recursive:false,bindings:[new Tuple(name,typed(new C.Func([C.Any.value],arrayType),ir(new S.Abs([new Tuple(new Just('input'),0)],call))))]};
    });
    return CodeGen.translate(metadata)({name:'NativeArrayFixture',bindings,comments:[],imports:emptySet,exports:emptySet,reExports:emptySet,
        dataTypes:emptyMap,dataDecls:[],classDecls:[],foreign:emptyMap,implementations:emptyMap,directives:emptyMap});
}

test('typed native callbacks and imported workers avoid temporary source copies while normalizing scalar views', t => {
    const code=fixture();
    assert.doesNotMatch(code,/unboxed := make\(\[\]int64/,'read-only scalar inputs must not be materialized');
    for(const name of ['mapped','filtered']){
        const worker=code.slice(code.indexOf(`func Call_NativeArrayFixture_${name}(`)).split('\nfunc ')[0];
        assert.doesNotMatch(worker,/gopurs_runtime.Apply\(/,'local lambda has a native ABI');
        assert.match(worker,/func\(value_1 int64\)/);
    }
    assert.match(code,/Call_Imported_increment\(/,'globalFunctions supplies imported native workers');
    const directory=mkdtempSync(join(tmpdir(),'gopurs-native-arrays-'));t.after(()=>rmSync(directory,{recursive:true,force:true}));
    mkdirSync(join(directory,'gopurs_runtime'));writeFileSync(join(directory,'go.mod'),'module gopurs/output\n\ngo 1.22\n');
    writeFileSync(join(directory,'gopurs_runtime/runtime.go'),runtimeGoCode);writeFileSync(join(directory,'fixture.go'),code.replace('package purescript','package main'));
    writeFileSync(join(directory,'main.go'),`package main
import("fmt";"reflect";"unsafe";"testing";r "gopurs/output/gopurs_runtime")
var emptySink []int64
var captures int
var calls []int64
func Get_Imported_increment()r.Value{captures++;return r.Value{}}
func Call_Imported_increment(v int64)int64{calls=append(calls,v);return r.IntAdd(v,1)}
func Get_Imported_predicate()r.Value{captures++;return r.Func(func(v r.Value)r.Value{
    if v.Type!=r.TypeInt||v.UnsafePtr!=nil{panic("scalar input was not normalized")};calls=append(calls,v.IntVal);return r.Bool(v.IntVal>=0)
})}
func check(ok bool){if !ok{panic("array result or alias failure")}}
func main(){
    marker:=int64(42)
    for _,values:=range [][]r.Value{{},{r.Value{Type:r.TypeString,IntVal:-2,UnsafePtr:unsafe.Pointer(&marker)},r.Int(0),r.Int(3),r.Int(2147483647)}}{
        before:=append([]r.Value{},values...);input:=r.Array(values);wantMap,wantFilter:=[]int64{},[]int64{}
        for _,v:=range values{wantMap=append(wantMap,r.IntAdd(v.IntVal,1));if v.IntVal>=0{wantFilter=append(wantFilter,v.IntVal)}}
        for name,run:=range map[string]func(r.Value)[]int64{"map":Call_NativeArrayFixture_mapped,"filter":Call_NativeArrayFixture_filtered,"import":Call_NativeArrayFixture_imported,"dynamic":Call_NativeArrayFixture_dynamic}{
            captures=0;calls=nil;out:=run(input);want:=wantMap;if name=="filter"||name=="dynamic"{want=wantFilter};check(reflect.DeepEqual(out,want));check(reflect.DeepEqual(values,before))
            if name=="import"||name=="dynamic"{check(captures==1&&len(calls)==len(values));for i,v:=range values{check(calls[i]==v.IntVal)}}
            second:=run(input);if len(out)>0{out[0]=99};check(reflect.DeepEqual(second,want));check(reflect.DeepEqual(values,before))
        }
    }
    negatives:=make([]r.Value,10000);for i:=range negatives{negatives[i]=r.Int(-1)};negativeInput:=r.Array(negatives)
    if allocs:=testing.AllocsPerRun(10,func(){emptySink=Call_NativeArrayFixture_filtered(negativeInput)});allocs!=0{panic(fmt.Sprint("empty filter allocated a buffer: ",allocs))}
    fmt.Println("native callbacks, imported workers, normalization, strict order and fresh arrays: ok")
}
`);
    const run=spawnSync('go',['run','.'],{cwd:directory,encoding:'utf8',timeout:60000,env:{...process.env,GOWORK:'off'}});
    assert.ifError(run.error);assert.equal(run.status,0,run.stdout+run.stderr);assert.match(run.stdout,/native callbacks, imported workers/);
});
