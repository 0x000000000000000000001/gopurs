import assert from 'node:assert/strict';
import {spawnSync} from 'node:child_process';
import {mkdirSync,mkdtempSync,rmSync,writeFileSync} from 'node:fs';
import {tmpdir} from 'node:os';
import {join} from 'node:path';
import test from 'node:test';
import {empty as emptyMap} from '../output/Data.Map/index.js';
import {Just,Nothing} from '../output/Data.Maybe/index.js';
import {empty as emptySet} from '../output/Data.Set/index.js';
import {Tuple} from '../output/Data.Tuple/index.js';
import * as CodeGen from '../output/Gopurs.CodeGen/index.js';
import {optimize,emit} from '../output/Gopurs.CodePointFold/index.js';
import {runtimeGoCode} from '../output/Gopurs.Runtime.Source/index.js';
import {TcoExpr} from '../output/PureScript.Backend.Optimizer.Codegen.Tco/index.js';
import * as C from '../output/PureScript.Backend.Optimizer.CoreFn/index.js';
import {NeutralExpr} from '../output/PureScript.Backend.Optimizer.Semantics/index.js';
import * as S from '../output/PureScript.Backend.Optimizer.Syntax/index.js';
import {withReboxFields} from './codegen-metadata.mjs';
const e=syntax=>new NeutralExpr(syntax);
const typed=(ty,x)=>e(new S.Typed(ty,x));
const int=x=>e(new S.Lit(new C.LitInt(x)));
const local=(level,ty)=>{const x=e(new S.Local(new Just('x'),level));return ty?typed(ty,x):x;};
const ref=(module,name)=>e(new S.Var(new C.Qualified(new Just(module),name)));
const app=(fn,...args)=>e(new S.App(fn,args));
const abs=(levels,body)=>e(new S.Abs(levels.map(l=>new Tuple(new Just('x'),l)),body));
const operation=(op,a,b)=>typed(C.Int.value,e(new S.PrimOp(new S.Op2(new S.OpIntNum(op),a,b))));
const add=(a,b)=>operation(S.OpAdd.value,a,b);
const array=new C.Array(C.Int.value), cbtype=new C.Func([C.Int.value,C.Int.value],C.Int.value);
const sum=typed(cbtype,abs([0,1],add(local(0),local(1))));
const fold=ref('Data.Foldable','foldlArray'),convert=ref('Data.String.CodePoints','toCodePointArray');
const stringFold=new C.Func([C.String.value],C.Int.value);
const appliedFunction=(arg,ret)=>new C.TypeApp(new C.ADT("Prim.Function",["Prim","Function"],[]),[arg,ret]);
const compose=(callback=sum,seed=int(0),conversion=convert,dictionary=ref('Control.Semigroupoid','semigroupoidFn'))=>typed(stringFold,app(ref('Control.Semigroupoid','compose'),dictionary,app(fold,callback,seed),conversion));
const moduleOf=(bindings,recursive=false)=>({name:'CodePointFixture',bindings:[{recursive,bindings:bindings.map(([n,x])=>new Tuple(n,x))}],comments:[],imports:emptySet,exports:emptySet,reExports:emptySet,dataTypes:emptyMap,dataDecls:[],classDecls:[],foreign:emptyMap,implementations:emptyMap,directives:emptyMap});
const transformed=(x,recursive=false)=>optimize(moduleOf([['fold',x]],recursive)).bindings[0].bindings[0].value1;
const metadata=withReboxFields({elidedCtors:emptySet,ctorTypes:emptyMap,pointerAdtPaths:emptyMap,pointerAdtNodes:emptySet,pointerAdtLeaves:emptyMap,enumAdts:emptySet,enumCtors:emptySet,globalTypes:emptyMap,globalFunctions:emptyMap,classDeclsFields:emptyMap});

test('only a closed primitive canonical composition exposes a native String worker',()=>{
 for(const x of [compose(),typed(stringFold,typed(appliedFunction(C.String.value,C.Int.value),compose())),compose(sum,int(37)),compose(sum,int(4294967295)),typed(C.Any.value,compose()),typed(new C.Func([new C.TypeVar('a')],new C.TypeVar('b')),compose())]){
  const result=transformed(x);assert.notDeepEqual(result,x);assert.ok(result instanceof S.Typed);assert.deepEqual(result.value0,stringFold);assert.ok(result.value1 instanceof S.Abs);assert.equal(result.value1.value0[0].value1,2);
 }
 const captured=typed(cbtype,abs([0,1],add(local(0),local(9,C.Int.value))));
 const called=typed(cbtype,abs([0,1],app(ref('Foreign','call'),local(1))));
 const staged=typed(cbtype,abs([0],e(new S.Let(new Just('saved'),9,int(2),abs([1],add(local(0),local(1)))))));
 const suffix=abs([0],typed(new C.Func([C.Boolean.value],C.Int.value),abs([1],add(local(0),local(1)))));
 for(const x of [typed(appliedFunction(C.Int.value,C.Int.value),compose()),typed(new C.TypeApp(new C.ADT("Other.Function",["Other","Function"],[]),[C.String.value,C.Int.value]),compose()),compose(captured),compose(called),compose(staged),compose(suffix),compose(sum,ref('Foreign','seed')),compose(sum,typed(C.Boolean.value,int(0))),compose(sum,int(0),ref('Other','toCodePointArray')),compose(sum,int(0),convert,ref('Other','semigroupoidFn')),compose(sum,int(0),typed(new C.Func([C.Int.value],array),convert)),typed(new C.Func([C.Int.value],C.Int.value),compose()),typed(new C.Func([new C.TypeVar('a')],new C.TypeVar('a')),compose()),compose(typed(new C.Func([C.Boolean.value,C.Int.value],C.Int.value),sum))])assert.deepEqual(transformed(x),x);
 assert.deepEqual(transformed(compose(),true),compose());
});

// Use the real TCO analyzer so Typed, TypeApp and lambda suffixes have the same
// shape seen by the emitter, rather than reconstructing analysis records.
const tco=await import('../output/PureScript.Backend.Optimizer.Codegen.Tco/index.js');
const analyze=x=>tco.analyze([])(x);
test('unknown, staged, effectful and incompatible immediate folds are rejected before translation',()=>{
 const never=()=>{throw Error('rejected candidates must not translate');};
 const reject=(fn,args,recVars=[])=>assert.deepEqual(emit(never)({recVars})(0)(analyze(fn))(args.map(analyze)),Nothing.value);
 const input=typed(array,app(convert,local(3,C.String.value)));
 reject(ref('Other','foldlArray'),[sum,int(0),input]);
 reject(fold,[sum,int(0)]);reject(fold,[sum,int(0),input,int(1)]);reject(fold,[sum,int(0),input],['recursive']);
 reject(fold,[local(7,cbtype),int(0),input]);
 reject(fold,[sum,typed(C.Boolean.value,int(0)),input]);
 reject(fold,[sum,int(0),typed(new C.Array(C.Boolean.value),input)]);
 reject(fold,[sum,int(0),typed(array,app(convert,local(3,C.Int.value)))]);
 reject(fold,[sum,int(0),typed(array,app(ref('Other','toCodePointArray'),local(3,C.String.value)))]);
 reject(fold,[sum,int(0),typed(array,app(convert,local(3,C.String.value),int(9)))]);
 const wrongSuffix=abs([0],typed(new C.Func([C.Boolean.value],C.Int.value),abs([1],add(local(0),local(1)))));
 reject(fold,[wrongSuffix,int(0),input]);
 const badBody=typed(cbtype,abs([0,1],app(ref('Foreign','call'),local(1))));reject(fold,[badBody,int(0),input]);
});

function generated(){
 const input=local(4,C.String.value),seed=local(5,C.Int.value);
 const direct=typed(new C.Func([C.Int.value,C.String.value],C.Int.value),abs([5,4],typed(C.Int.value,app(fold,sum,seed,typed(array,app(convert,input))))));
 const weighted=typed(cbtype,abs([0,1],add(operation(S.OpMultiply.value,local(0),int(3)),local(1))));
 return CodeGen.translate(metadata)(moduleOf([['fold',compose()],['seeded',direct],['weighted',compose(weighted,int(17))],['wideEmpty',compose(sum,int(4294967295))]]));
}
test('generated native and public workers fold directly with Unicode decoding and unchanged Int arithmetic',t=>{
 const code=generated();assert.match(code,/func Call_CodePointFixture_fold\([^\n]*string/);assert.match(code,/DecodeStringCodePoint/);assert.doesNotMatch(code,/Get_Data_String_CodePoints_toCodePointArray|Get_Data_Foldable_foldlArray|Get_Control_Semigroupoid_compose|make\(\[\]/);
 const dir=mkdtempSync(join(tmpdir(),'gopurs-codepoint-fold-'));t.after(()=>rmSync(dir,{recursive:true,force:true}));mkdirSync(join(dir,'gopurs_runtime'));writeFileSync(join(dir,'go.mod'),'module gopurs/output\n\ngo 1.22\n');writeFileSync(join(dir,'gopurs_runtime/runtime.go'),runtimeGoCode);writeFileSync(join(dir,'fixture.go'),code.replace('package purescript','package main'));
 writeFileSync(join(dir,'main.go'),`package main
import("fmt";"strings";"testing";r "gopurs/output/gopurs_runtime")
var sink int64
func want(s string,seed,mult int64)int64{for len(s)>0{v,n:=r.DecodeStringCodePoint(s);s=s[n:];seed=r.IntAdd(r.IntMul(seed,mult),int64(v))};return seed}
func check(ok bool){if !ok{panic("fold mismatch or allocation")}}
func main(){checks:=0;for _,s:=range []string{"","a","a😀é","\\xed\\xa0\\xbd\\xed\\xb8\\x80","\\xed\\xa0\\x80x\\xed\\xb0\\x80","\\xff\\xed\\xa0",strings.Repeat("😀",10000)}{for _,seed:=range []int64{0,17,-1,2147483647,-2147483648,4294967295,1<<52+1}{
check(Call_CodePointFixture_seeded(seed,s)==want(s,seed,1));check(r.Apply2(Get_CodePointFixture_seeded(),r.Int(seed),r.Str(s)).IntVal==want(s,seed,1));checks+=2}
check(Call_CodePointFixture_fold(s)==want(s,0,1));check(r.Apply(Get_CodePointFixture_fold(),r.Str(s)).IntVal==want(s,0,1));check(Call_CodePointFixture_weighted(s)==want(s,17,3));check(Call_CodePointFixture_wideEmpty(s)==want(s,4294967295,1));checks+=4}
s:=strings.Repeat("a😀é",500);getter:=Get_CodePointFixture_fold();check(testing.AllocsPerRun(10,func(){sink=r.Apply(getter,r.Str(s)).IntVal})==0);fmt.Printf("%d native/public checks: ok\\n",checks)}
`);
 const result=spawnSync('go',['run','.'],{cwd:dir,encoding:'utf8',timeout:60000,env:{...process.env,GOWORK:'off'}});assert.ifError(result.error);assert.equal(result.status,0,result.stdout+result.stderr);assert.match(result.stdout,/126 native\/public checks: ok/);
});
