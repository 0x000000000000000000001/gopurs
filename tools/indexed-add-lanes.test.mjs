import assert from 'node:assert/strict';
import test from 'node:test';
import {spawnSync} from 'node:child_process';
import {mkdtempSync,mkdirSync,readFileSync,writeFileSync,rmSync} from 'node:fs';
import {tmpdir} from 'node:os';
import {join} from 'node:path';
import {Just} from '../output/Data.Maybe/index.js';
import {Tuple} from '../output/Data.Tuple/index.js';
import {eligible,emit} from '../output/Gopurs.IndexedAddLanes/index.js';
import {printGoExpr} from '../output/Gopurs.Printer/index.js';
import * as G from '../output/Gopurs.GoAst/index.js';
import * as C from '../output/PureScript.Backend.Optimizer.CoreFn/index.js';
import * as S from '../output/PureScript.Backend.Optimizer.Syntax/index.js';
import {NeutralExpr} from '../output/PureScript.Backend.Optimizer.Semantics/index.js';
import {analyze} from '../output/PureScript.Backend.Optimizer.Codegen.Tco/index.js';
const n=x=>new NeutralExpr(x),typed=(ty,x)=>n(new S.Typed(ty,x));
const int=x=>typed(C.Int.value,n(new S.Lit(new C.LitInt(x))));
const local=l=>typed(C.Int.value,n(new S.Local(new Just('v'),l)));
const op=(o,a,b)=>typed(C.Int.value,n(new S.PrimOp(new S.Op2(new S.OpIntNum(o),a,b))));
const add=(a,b)=>op(S.OpAdd.value,a,b),mul=(a,b)=>op(S.OpMultiply.value,a,b);
const lambda=(levels,body)=>typed(new C.Func(levels.map(()=>C.Int.value),C.Int.value),n(new S.Abs(levels.map(l=>new Tuple(new Just('v'),l)),body)));
const mapper=lambda([10],add(local(10),int(1)));
const contribution=mul(add(int(1),op(S.OpMod.value,local(0),int(2))),local(2));
const folder=(body=add(local(1),contribution))=>lambda([0,1,2],body);
const accepts=(m,f)=>eligible(analyze([])(m))(analyze([])(f));
test('lanes require a normalizing closed mapper and an independent additive contribution',()=>{
 assert.equal(accepts(mapper,folder()),true);
 assert.equal(accepts(lambda([10,11],add(local(10),local(11))),folder()),true);
 for(const bad of [lambda([10],local(10)),lambda([10],add(local(10),local(99))),typed(new C.Func([C.String.value],C.Int.value),mapper)])assert.equal(accepts(bad,folder()),false);
 const hiddenLet=typed(C.Int.value,n(new S.Let(new Just('hidden'),7,local(1),add(local(7),local(2)))));
 const hiddenBranch=typed(C.Int.value,n(new S.Branch([new S.Pair(n(new S.Lit(new C.LitBoolean(true))),local(1))],local(2))));
 for(const bad of [folder(add(local(1),add(local(1),local(2)))),folder(add(local(1),hiddenLet)),folder(add(local(1),hiddenBranch)),folder(add(mul(local(1),int(3)),local(2))),folder(add(local(1),op(S.OpDivide.value,local(2),int(3)))),folder(add(local(1),op(S.OpMod.value,local(2),local(0)))),folder(add(local(1),op(S.OpMod.value,local(2),int(0)))),folder(add(local(1),local(99)))])assert.equal(accepts(mapper,bad),false);
});

test('generated lanes preserve scalar order, tails, wide seeds and exact IntMul',t=>{
 const dir=mkdtempSync(join(tmpdir(),'gopurs-indexed-lanes-'));t.after(()=>rmSync(dir,{recursive:true,force:true}));mkdirSync(join(dir,'gopurs_runtime'));
 writeFileSync(join(dir,'go.mod'),'module gopurs/output\n\ngo 1.22\n');writeFileSync(join(dir,'gopurs_runtime/runtime.go'),readFileSync(new URL('../runtime/runtime.go',import.meta.url)));
 const functions=[];
 for(const boxed of [false,true]){
  const name=boxed?'boxed':'native', value=x=>boxed?new G.GoSelector(x,'IntVal'):x;
  const step=(i,a,v)=>new G.GoCall(G.rawGo('step'),[new G.GoCall(G.rawGo('int64'),[i]),a,value(v)]);
  const fallback=new G.GoForRange('i, value := range xs',[new G.GoMutate('result',step(G.rawGo('i'),G.rawGo('result'),G.rawGo('value')))]);
  const code=printGoExpr(emit({source:'xs',prefix:'lanes',result:'result'})(i=>a=>v=>step(i,a,v))(fallback)).replaceAll('gopurs_runtime.','r.');
  assert.match(code,/len\(xs\) < 8/);assert.match(code,/<= len\(xs\)-8/);
  functions.push(`func ${name}(seed int64,xs []${boxed?'r.Value':'int64'})int64{result:=seed;${code};return result}`);
 }
 writeFileSync(join(dir,'main.go'),`package main
import("fmt";"reflect";"math/rand";r "gopurs/output/gopurs_runtime")
var mode int
var trace []int64
func step(i,acc,value int64)int64{trace=append(trace,i,value);mapped:=r.IntAdd(value,1);if mode==1{mapped=r.IntAdd(value,i)};factor:=r.IntAdd(1,i&1);if mode==2{factor=2147483647};return r.IntAdd(acc,r.IntMul(factor,mapped))}
func reference(seed int64,xs []int64)int64{for i,x:=range xs{seed=step(int64(i),seed,x)};return seed}
${functions.join('\n')}
func main(){rng:=rand.New(rand.NewSource(1337));checks:=0;lengths:=[]int{};for i:=0;i<=17;i++{lengths=append(lengths,i)};lengths=append(lengths,31,32,33,1000,10000)
 for _,length:=range lengths{for kind:=0;kind<4;kind++{xs:=make([]int64,length);for i:=range xs{switch kind{case 0:xs[i]=int64(i);case 1:xs[i]=int64(rng.Uint64());case 2:xs[i]=-1<<63;case 3:xs[i]=1<<63-1}};values:=make([]r.Value,length);for i,x:=range xs{values[i]=r.Int(x);if i%3==0{values[i].Type=r.TypeAny}};saved:=append([]int64{},xs...)
 for mode=0;mode<3;mode++{for _,seed:=range []int64{0,17,4294967295,-1<<63,1<<63-1}{trace=nil;want:=reference(seed,xs);order:=append([]int64{},trace...);trace=nil;got:=native(seed,xs);if got!=want||!reflect.DeepEqual(append([]int64{},trace...),order){panic("native result/order")};trace=nil;got=boxed(seed,values);if got!=want||!reflect.DeepEqual(append([]int64{},trace...),order){panic("boxed result/order")};checks+=2}}
 for i,x:=range xs{if x!=saved[i]||values[i].IntVal!=x{panic("input mutated")}}}}
 fmt.Printf("%d generated lane comparisons: ok\\n",checks)}
`);
 const result=spawnSync('go',['run','.'],{cwd:dir,encoding:'utf8',timeout:60000,env:{...process.env,GOWORK:'off'}});assert.ifError(result.error);assert.equal(result.status,0,result.stdout+result.stderr);assert.match(result.stdout,/2760 generated lane comparisons: ok/);
});
