import { withReboxFields } from './codegen-metadata.mjs';
import assert from 'node:assert/strict';
import { spawnSync } from 'node:child_process';
import { mkdirSync, mkdtempSync, rmSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join } from 'node:path';
import test from 'node:test';
import { empty as emptyMap, singleton } from '../output/Data.Map/index.js';
import { empty as emptySet } from '../output/Data.Set/index.js';
import { Just, Nothing } from '../output/Data.Maybe/index.js';
import { Tuple } from '../output/Data.Tuple/index.js';
import * as CodeGen from '../output/Gopurs.CodeGen/index.js';
import * as Go from '../output/Gopurs.GoAst/index.js';
import { LocalReference } from '../output/Gopurs.GoSymbol/index.js';
import { runtimeGoCode } from '../output/Gopurs.Runtime.Source/index.js';
import * as C from '../output/PureScript.Backend.Optimizer.CoreFn/index.js';
import * as S from '../output/PureScript.Backend.Optimizer.Syntax/index.js';
import { NeutralExpr } from '../output/PureScript.Backend.Optimizer.Semantics/index.js';

const metadata = withReboxFields({elidedCtors:emptySet,ctorTypes:emptyMap,pointerAdtPaths:emptyMap,
  pointerAdtNodes:emptySet,pointerAdtLeaves:emptyMap,enumAdts:emptySet,enumCtors:emptySet,
  globalTypes:emptyMap,globalFunctions:emptyMap,classDeclsFields:emptyMap});
const int=C.Int.value, boolean=C.Boolean.value;
const expr=x=>new NeutralExpr(x), typed=(ty,x)=>expr(new S.Typed(ty,x));
const literal=n=>typed(int,expr(new S.Lit(new C.LitInt(n))));
const local=(name,level,ty)=>typed(ty,expr(new S.Local(new Just(name),level)));
const bin=(op,a,b,ty=int)=>typed(ty,expr(new S.PrimOp(new S.Op2(op,a,b))));
const add=(a,b)=>bin(new S.OpIntNum(S.OpAdd.value),a,b);
const sub=(a,b)=>bin(new S.OpIntNum(S.OpSubtract.value),a,b);
const rec=fields=>new C.Record(new C.Row(fields.map(([k,v])=>new Tuple(k,v)),Nothing.value));
const row=rec([['borrowed',new C.Array(int)],['stamp',C.String.value],['total',int]]);
const single=rec([['total',int]]);
const nested=rec([['inner',single],['total',int]]);
const access=(x,k,ty=int)=>typed(ty,expr(new S.Accessor(x,new S.GetProp(k))));
const update=(x,props,ty)=>typed(ty,expr(new S.Update(x,props.map(([k,v])=>new C.Prop(k,v)))));
const fnType=ty=>new C.Func([int,ty],ty);
const lambda=(body,ty=row)=>typed(fnType(ty),expr(new S.Abs([new Tuple(new Just('left'),0),new Tuple(new Just('snapshot'),1)],body)));
function recurrence(name,delta,options={}) {
  const ty=options.ty??row,n=local('left',0,int),state=local('snapshot',1,ty);
  const next=options.change?.(state,n)??update(state,[['total',add(access(state,'total'),literal(delta))]],ty);
  const call=typed(ty,expr(new S.App(typed(fnType(ty),expr(new S.Var(new C.Qualified(new Just('HotRecords'),name)))),[sub(n,literal(1)),next])));
  const body=typed(ty,expr(new S.Branch([new S.Pair(bin(new S.OpIntOrd(S.OpEq.value),n,literal(0),boolean),state)],call)));
  return {recursive:true,bindings:[new Tuple(name,lambda(body,ty))]};
}
const moduleOf=bindings=>({name:'HotRecords',bindings,comments:[],imports:emptySet,exports:emptySet,reExports:emptySet,
  dataTypes:emptyMap,dataDecls:[],classDecls:[],foreign:emptyMap,implementations:emptyMap,directives:emptyMap});
const generate=(bindings,meta=metadata)=>CodeGen.translateWithFunctions(meta)(moduleOf(bindings));
function worker(code,name){const start=code.indexOf(`func Call_HotRecords_${name}(`);assert.ok(start>=0,name);const end=code.indexOf('\nfunc ',start+1);return code.slice(start,end<0?undefined:end);}

function eligible() {return generate([recurrence('increase',1),recurrence('decrease',-1),recurrence('normalize',0),recurrence('tiny',1,{ty:single})]);}

test('module workers split only proved small affine updates and retain their public signature',()=>{
  const generated=eligible();
  for(const name of ['increase','decrease','normalize','tiny']){
    const body=worker(generated.code,name),cold=worker(generated.code,'__gopurs_affine_cold_'+name);
    assert.match(body,/uint64\(/);assert.match(body,/2147483647/);assert.match(body,/int64\(int32\(/);
    assert.doesNotMatch(body,/for \{/);assert.match(cold,/for \{/);
    assert.match(body,new RegExp(`return Call_HotRecords___gopurs_affine_cold_${name}\\(`));
    assert.equal(body.slice(0,body.indexOf(' {')),cold.replace('__gopurs_affine_cold_','').slice(0,cold.replace('__gopurs_affine_cold_','').indexOf(' {')));
  }
  assert.doesNotMatch(worker(generated.code,'normalize'),/total\)\s*[+-]/,'delta zero must still narrow without introducing an arithmetic dependency');
});

test('larger deltas, nested writes, cross-field dependencies and reserved helper names keep the original worker',()=>{
  const bindings=[recurrence('multiply',2),recurrence('nested',1,{ty:nested,change:s=>update(s,[['inner',update(access(s,'inner',single),[['total',add(access(access(s,'inner',single),'total'),literal(1))]],single)]],nested)}),
    recurrence('dependent',1,{ty:rec([['other',int],['total',int]]),change:s=>update(s,[['total',add(access(s,'total'),access(s,'other'))]],rec([['other',int],['total',int]]))}),
    recurrence('reserved',1),{recursive:false,bindings:[new Tuple('__gopurs_affine_cold_reserved',lambda(local('snapshot',1,row)))]}];
  const code=generate(bindings).code;
  for(const name of ['multiply','nested','dependent','reserved'])assert.doesNotMatch(worker(code,name),/uint64\(/,name);
  assert.equal((code.match(/func Call_HotRecords___gopurs_affine_cold_reserved\(/g)||[]).length,1,'existing helper-like user binding remains unique');
  const foreign={arrayFold:Nothing.value,reference:new LocalReference('Call_HotRecords___gopurs_affine_cold_external'),fArgs:[],fRet:Go.TypeInt64.value,arity:0};
  const ffi=generate([recurrence('external',1)],{...metadata,ffiFunctions:singleton('foreignName')(foreign)});
  assert.doesNotMatch(worker(ffi.code,'external'),/uint64\(/,'an existing raw FFI worker name is reserved');
});

test('real generated wrappers agree with the original workers for wide seeds and short cold paths',t=>{
  const generated=eligible();
  const directory=mkdtempSync(join(tmpdir(),'gopurs-record-affine-hot-'));t.after(()=>rmSync(directory,{recursive:true,force:true}));
  mkdirSync(join(directory,'purescript'));mkdirSync(join(directory,'gopurs_runtime'));
  writeFileSync(join(directory,'go.mod'),'module gopurs/output\n\ngo 1.22\n');
  writeFileSync(join(directory,'gopurs_runtime/runtime.go'),runtimeGoCode);
  writeFileSync(join(directory,'purescript/Generated.go'),generated.code);
  writeFileSync(join(directory,'purescript/probe_test.go'),`package purescript
import("testing";"reflect")
type state=struct{borrowed []int64;stamp string;total int64}
func TestWrappers(t *testing.T){checks:=0
 cases:=[]struct{hot,cold func(int64,state)state;delta int64}{
  {Call_HotRecords_increase,Call_HotRecords___gopurs_affine_cold_increase,1},
  {Call_HotRecords_decrease,Call_HotRecords___gopurs_affine_cold_decrease,-1},
  {Call_HotRecords_normalize,Call_HotRecords___gopurs_affine_cold_normalize,0},
 }
 for _,c:=range cases{for _,n:=range []int64{0,1,2,7,1000,2147483647,4294967297,4294967298,4294967303,-4294967295,-4294967294}{for _,seed:=range []int64{-9223372036854775808,-2147483649,-2147483648,-1,0,2147483647,2147483648,4294967295,9223372036854775807}{
  input:=state{[]int64{7,11},"kept",seed};got,want:=c.hot(n,input),c.cold(n,input)
  if !reflect.DeepEqual(got,want)||input.total!=seed||got.stamp!="kept"||&got.borrowed[0]!=&input.borrowed[0]||input.borrowed[0]!=7{t.Fatalf("delta=%d n=%d seed=%d got=%v want=%v",c.delta,n,seed,got,want)}
  if n==0&&got.total!=seed{t.Fatal("zero changed noncanonical seed")};if n>0&&n<=2147483647&&got.total!=int64(int32(seed+n*c.delta)){t.Fatal("mod32 result")};checks++
 }}}
 for _,n:=range []int64{0,1,2,2147483647,4294967297}{input:=struct{total int64}{4294967295};if Call_HotRecords_tiny(n,input)!=Call_HotRecords___gopurs_affine_cold_tiny(n,input){t.Fatal("single field")};checks++}
 // Check the exact interval identity on boundary values and a deterministic
 // full-width sample without executing any potentially long cold loops.
 samples:=[]int64{-9223372036854775808,-4294967296,-2147483648,-1,0,1,2147483646,2147483647,2147483648,4294967295,4294967296,9223372036854775807};var random uint64=4321987
 for i:=0;i<10000;i++{random^=random<<13;random^=random>>7;random^=random<<17;samples=append(samples,int64(random))}
 for _,n:=range samples{if (uint64(n)-1<2147483647)!=(n>0&&n<=2147483647){t.Fatal("interval mismatch",n)};checks++}
 t.Logf("%d wrapper, alias, modulo and interval comparisons passed",checks)
}`);
  const result=spawnSync('go',['test','-count=1','-v','./purescript'],{cwd:directory,encoding:'utf8',timeout:60000,env:{...process.env,GOWORK:'off'}});
  assert.ifError(result.error);assert.equal(result.status,0,result.stdout+result.stderr);
  console.log(result.stdout.trim());
});
