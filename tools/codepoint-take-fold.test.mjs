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
import {optimize,describe,validDescriptor} from '../output/Gopurs.CodePointFold/index.js';
import {runtimeGoCode} from '../output/Gopurs.Runtime.Source/index.js';
import {runtimeSimdFiles} from '../output/Gopurs.Runtime.Simd/index.js';
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

import {emit} from '../output/Gopurs.CodePointTakeFold/index.js';
import {analyze as analyzeTco} from '../output/PureScript.Backend.Optimizer.Codegen.Tco/index.js';
const analyze=x=>analyzeTco([])(x);
const take=(count=int(2),input=local(4,C.String.value),head=ref('Data.String.CodePoints','take'))=>typed(C.String.value,app(head,count,input));

test('codepoint descriptors preserve the exact closed callback and every TAST layer',()=>{
 for(const seed of [0,7,2147483647,4294967295]){
  const d=describe(transformed(compose(sum,int(seed))));assert.ok(d instanceof Just);assert.equal(d.value0.seed,seed);assert.deepEqual(d.value0.callback,sum);assert.equal(validDescriptor(d.value0),true);
 }
 const explicit=typed(stringFold,abs([4],typed(C.Int.value,app(fold,sum,int(7),typed(array,app(convert,local(4,C.String.value)))))));
 assert.ok(describe(explicit) instanceof Just);
 for(const badHead of [typed(C.String.value,fold),typed(new C.Func([cbtype,C.Int.value,array],C.Int.value),typed(C.String.value,fold))]) assert.deepEqual(describe(typed(stringFold,abs([4],app(badHead,sum,int(7),app(convert,local(4,C.String.value)))))),Nothing.value);
 const captured=typed(cbtype,abs([0,1],add(local(0),local(9,C.Int.value))));
 const contradictory=typed(cbtype,abs([0,1],add(local(0),typed(C.Int.value,typed(C.String.value,local(1))))));
 const wrongSuffix=typed(cbtype,abs([0],typed(new C.Func([C.String.value],C.Int.value),abs([1],add(local(0),local(1))))));
 for(const callback of [captured,contradictory,wrongSuffix,typed(cbtype,abs([0,1],app(ref('Foreign','call'),local(1))))])assert.equal(validDescriptor({callback,seed:0}),false);
 for(const bad of [typed(new C.Func([C.Int.value],C.Int.value),explicit),typed(stringFold,abs([4],app(fold,sum,int(7),app(convert,local(9,C.String.value))))),typed(stringFold,abs([4],app(fold,sum,local(9,C.Int.value),app(convert,local(4,C.String.value))))),typed(stringFold,abs([4],typed(C.String.value,app(fold,sum,int(7),app(convert,local(4,C.String.value))))))])assert.deepEqual(describe(bad),Nothing.value);
});

test('bounded codepoint folds reject foreign names, stages and contradictory inputs before translation',()=>{
 const never=()=>{throw Error('rejected candidate translated');};
 const reject=(p,d={callback:sum,seed:0},recVars=[])=>assert.deepEqual(emit(never)({recVars})(0)(d)(analyze(p)),Nothing.value);
 reject(take(int(2),local(4,C.String.value),ref('Trap','take')));
 reject(take(typed(C.String.value,int(2))));reject(take(int(2),local(4,C.Int.value)));
 reject(typed(C.Int.value,take()));reject(app(ref('Data.String.CodePoints','take'),int(2)));
 reject(take(int(2),local(4,C.String.value),typed(new C.Func([C.String.value,C.String.value],C.String.value),ref('Data.String.CodePoints','take'))));
 reject(take(int(2),local(4,C.String.value),e(new S.UncurriedApp(ref('Data.String.CodePoints','take'),[]))));
 reject(take(),{callback:sum,seed:0},['recursive']);
 reject(take(),{callback:typed(cbtype,abs([0,1],add(local(0),local(9,C.Int.value)))),seed:0});
});

function generated(){
 const consumer=(name,callback,seed)=>[name,compose(callback,int(seed))];
 const caller=name=>[name+'Take',typed(new C.Func([C.Int.value,C.String.value],C.Int.value),abs([3,4],typed(C.Int.value,app(ref('CodePointFixture',name),take(local(3,C.Int.value),local(4,C.String.value))))))];
 const weighted=typed(cbtype,abs([0,1],add(operation(S.OpMultiply.value,local(0),int(3)),local(1))));
 const wideProduct=typed(cbtype,abs([0,1],add(operation(S.OpMultiply.value,local(0),int(2147483647)),local(1))));
 return CodeGen.translate(metadata)(moduleOf([consumer('sum',sum,0),consumer('weighted',weighted,17),consumer('wide',sum,4294967295),consumer('wideProduct',wideProduct,2147483647),caller('wideProduct'),caller('sum'),caller('weighted'),caller('wide')]));
}
test('generated bounded workers preserve order, empty seeds and native UTF8/WTF8 semantics',t=>{
 const code=generated();assert.match(code,/codePointTake_limit/);assert.doesNotMatch(code,/Get_Data_String_CodePoints_take|Get_Data_String_CodePoints_toCodePointArray|make\(\[\]/);
 const dir=mkdtempSync(join(tmpdir(),'gopurs-codepoint-take-'));t.after(()=>rmSync(dir,{recursive:true,force:true}));mkdirSync(join(dir,'gopurs_runtime'));writeFileSync(join(dir,'go.mod'),'module gopurs/output\n\ngo 1.22\n');writeFileSync(join(dir,'gopurs_runtime/runtime.go'),runtimeGoCode);for(const f of runtimeSimdFiles)writeFileSync(join(dir,'gopurs_runtime',f.name),f.content);writeFileSync(join(dir,'fixture.go'),code.replace('package purescript','package main'));
 writeFileSync(join(dir,'main.go'),`package main
import("fmt";"strings";"math/rand";r "gopurs/output/gopurs_runtime")
func want(s string,n,seed,mult int64)int64{for i:=int64(0);i<n&&len(s)>0;i++{v,w:=r.DecodeStringCodePoint(s);s=s[w:];seed=r.IntAdd(r.IntMul(seed,mult),int64(v))};return seed}
func check(got,want int64){if got!=want{panic(fmt.Sprintf("got%d want%d",got,want))}}
func main(){checks:=0;ss:=[]string{"","abc","a😀é",strings.Repeat("a😀é",1000),"\\xed\\xa0\\xbd\\xed\\xb8\\x80","\\xed\\xa0\\x80x\\xed\\xb0\\x80","\\xff\\xed\\xa0"};rng:=rand.New(rand.NewSource(772));for i:=0;i<100;i++{v:=make([]byte,i%70);for j:=range v{v[j]=byte(rng.Intn(256))};ss=append(ss,string(v))};for _,s:=range ss{for _,n:=range []int64{-1<<63,-10,-1,0,1,2,3,4,7,8,15,16,69,70,100,1499,1500,1501,3000,1<<31-1,1<<62}{
check(Call_CodePointFixture_sumTake(n,s),want(s,n,0,1));check(r.Apply2(Get_CodePointFixture_sumTake(),r.Int(n),r.Str(s)).IntVal,want(s,n,0,1));check(Call_CodePointFixture_weightedTake(n,s),want(s,n,17,3));check(Call_CodePointFixture_wideTake(n,s),want(s,n,4294967295,1));check(Call_CodePointFixture_wideProductTake(n,s),want(s,n,2147483647,2147483647));checks+=5}};fmt.Printf("%d bounded native/public checks: ok\\n",checks)}
`);
 const result=spawnSync('go',['run','.'],{cwd:dir,encoding:'utf8',timeout:60000,env:{...process.env,GOWORK:'off'}});assert.ifError(result.error);assert.equal(result.status,0,result.stdout+result.stderr);assert.match(result.stdout,/11235 bounded native\/public checks: ok/);
});
