import {withReboxFields} from './codegen-metadata.mjs';
import assert from 'node:assert/strict';
import {spawnSync} from 'node:child_process';
import {mkdirSync,mkdtempSync,rmSync,writeFileSync} from 'node:fs';
import {tmpdir} from 'node:os';
import {join} from 'node:path';
import test from 'node:test';
import {empty as emptyMap,singleton,insert} from '../output/Data.Map/index.js';
import {ordString} from '../output/Data.Ord/index.js';
import * as Go from '../output/Gopurs.GoAst/index.js';
import {emit as emitConsumer} from '../output/Gopurs.RecordConsumer/index.js';
import {TcoExpr} from '../output/PureScript.Backend.Optimizer.Codegen.Tco/index.js';
import {localId} from '../output/PureScript.Backend.Optimizer.FreeVars/index.js';
import {empty as emptySet} from '../output/Data.Set/index.js';
import {Just,Nothing} from '../output/Data.Maybe/index.js';
import {Tuple} from '../output/Data.Tuple/index.js';
import * as CodeGen from '../output/Gopurs.CodeGen/index.js';
import {runtimeGoCode} from '../output/Gopurs.Runtime.Source/index.js';
import * as C from '../output/PureScript.Backend.Optimizer.CoreFn/index.js';
import * as S from '../output/PureScript.Backend.Optimizer.Syntax/index.js';
import {NeutralExpr} from '../output/PureScript.Backend.Optimizer.Semantics/index.js';
const metadata=withReboxFields({elidedCtors:emptySet,ctorTypes:emptyMap,pointerAdtPaths:emptyMap,pointerAdtNodes:emptySet,pointerAdtLeaves:emptyMap,enumAdts:emptySet,enumCtors:emptySet,globalTypes:emptyMap,globalFunctions:emptyMap,classDeclsFields:emptyMap});
const int=C.Int.value,bool=C.Boolean.value,e=s=>new NeutralExpr(s),typed=(t,x)=>e(new S.Typed(t,x));
const lit=n=>typed(int,e(new S.Lit(new C.LitInt(n)))),local=(n,l,t)=>typed(t,e(new S.Local(new Just(n),l)));
const binary=(o,a,b,t=int)=>typed(t,e(new S.PrimOp(new S.Op2(o,a,b))));
const add=(a,b)=>binary(new S.OpIntNum(S.OpAdd.value),a,b),sub=(a,b)=>binary(new S.OpIntNum(S.OpSubtract.value),a,b);
const effect=t=>new C.ADT('Effect.Effect',['Effect','Effect'],[t]);
const row=keys=>new C.Record(new C.Row(keys.map(k=>new Tuple(k,int)),Nothing.value));
const access=(r,k)=>typed(int,e(new S.Accessor(r,new S.GetProp(k))));
const variable=(name,t)=>typed(t,e(new S.Var(new C.Qualified(new Just('VariableNames'),name))));
const call=(name,args,t,ret)=>typed(ret,e(new S.App(variable(name,t),args)));
const binding=(name,args,t,body,recursive=false)=>({recursive,bindings:[new Tuple(name,typed(t,e(new S.Abs(args.map(([n,l])=>new Tuple(new Just(n),l)),body))))]});
function recurrence(name,keys,deltas,options={}){const r=row(keys),n=local('remaining',0,int),original=local('source',1,r),fn=new C.Func(options.recordFirst?[r,int]:[int,r],r);
 let update=typed(r,e(new S.Update(original,deltas.map(([k,d])=>new C.Prop(k,options.cross?add(access(original,keys[0]),lit(d)):add(access(original,k),lit(d)))))));
 if(options.badUpdate)update=typed(r,typed(C.String.value,update));
 const next=call(name,options.recordFirst?[update,sub(n,lit(1))]:[sub(n,lit(1)),update],fn,r),condition=binary(new S.OpIntOrd(S.OpEq.value),n,lit(0),bool);
 return binding(name,options.recordFirst?[['source',1],['remaining',0]]:[['remaining',0],['source',1]],fn,typed(r,e(new S.Branch([new S.Pair(condition,original)],next))),true);
}
const identity=binding('identity',[['value',0]],new C.Func([int],int),local('value',0,int));
function consume(name,step,keys,options={}){const r=row(keys),n=local('count',options.boxedCount?4:0,int),pending=local('pending',1,effect(r)),original=local('original',2,r),changed=local('changed',3,r);
 let leaves=[...keys.map(k=>access(changed,k)),...keys.map(k=>access(original,k))];
 if(options.subset)leaves.pop();if(options.duplicate)leaves[leaves.length-1]=access(original,keys[0]);if(options.reversed)leaves.reverse();
 let sum=leaves.reduce(add);if(options.constant)sum=add(sum,lit(7));if(options.reference)sum=call('identity',[sum],new C.Func([int],int),int);
 let result=typed(effect(int),e(new S.EffectPure(sum)));
 // The real optimizer preserves a generic Effect a wrapper under Effect Int.
 result=typed(effect(int),typed(effect(new C.TypeVar('a$scope27')),result));
 if(options.badEffect)result=typed(effect(C.String.value),result);
 if(options.intermediate)result=e(new S.Let(new Just('seen'),4,access(changed,keys[0]),result));
 if(options.escaping)result=e(new S.Let(new Just('saved'),4,result,local('saved',4,effect(int))));
 const updated=call(step,options.recordFirst?[original,n]:[n,original],new C.Func(options.recordFirst?[r,int]:[int,r],r),r);
 let body=typed(effect(int),e(new S.EffectBind(new Just('original'),2,pending,e(new S.Let(new Just('changed'),3,updated,result)))));
 if(options.boxedCount)body=typed(effect(int),e(new S.EffectBind(new Just('count'),4,local('countSource',0,effect(int)),body)));
 return binding(name,[[options.boxedCount?'countSource':'count',0],['pending',1]],new C.Func([options.boxedCount?effect(int):int,effect(r)],effect(int)),body);
}
const moduleOf=bindings=>({name:'VariableNames',bindings,comments:[],imports:emptySet,exports:emptySet,reExports:emptySet,dataTypes:emptyMap,dataDecls:[],classDecls:[],foreign:emptyMap,implementations:emptyMap,directives:emptyMap});
const generate=bindings=>CodeGen.translateWithFunctions(metadata)(moduleOf(bindings)).code;
function worker(code,name){const a=code.indexOf(`func Call_VariableNames_${name}(`);assert.ok(a>=0,name);const b=code.indexOf('\nfunc ',a+1);return code.slice(a,b<0?undefined:b);}
const keys=['alpha','beta','delta','gamma','iota','omega','theta','zeta'];
function fixtures(){const specs=[['up',keys,[['theta',1]]],['down',keys,[['beta',-1]]],['zero',keys,[['gamma',0]]],['several',keys,[['alpha',3],['zeta',-2]]],['single',['only'],[['only',0]]],['pair',['first','last'],[['first',-3],['last',2]]],['flipped',keys,[['omega',3]],{recordFirst:true}],['huge',['first','last'],[['first',2147483647]]]];
 const bindings=[identity];for(const[name,fields,deltas,options={}]of specs)bindings.push(recurrence(name,fields,deltas,options),consume('fused_'+name,name,fields,options),consume('reference_'+name,name,fields,{...options,reference:true}));bindings.push(consume('boxed','up',keys,{boxedCount:true}),consume('boxedReference','up',keys,{boxedCount:true,reference:true}));return{specs,bindings};}

test('intramodule affine consumers are generic in names, deltas and field counts',()=>{
 const {specs,bindings}=fixtures(),code=generate(bindings);
 for(const[name]of specs){assert.match(worker(code,'fused_'+name),/recordConsumer_/);assert.doesNotMatch(worker(code,'reference_'+name),/recordConsumer_/);}
 assert.match(code,/2147483647/);assert.match(worker(code,'boxed'),/recordConsumer_/);assert.doesNotMatch(worker(code,'boxedReference'),/recordConsumer_/);
});

test('incomplete projections, extra uses, intermediate bindings and unknown calls retain their path',()=>{
 const bindings=[identity,recurrence('step',keys,[['theta',1]])];for(const[name,opt]of Object.entries({subset:{subset:true},duplicate:{duplicate:true},constant:{constant:true},unknown:{reference:true},intermediate:{intermediate:true},escaping:{escaping:true},badEffect:{badEffect:true}}))bindings.push(consume(name,'step',keys,opt));
 bindings.push(recurrence('badUpdate',keys,[['theta',1]],{badUpdate:true}),consume('badUpdateConsumer','badUpdate',keys));
 const code=generate(bindings);assert.doesNotMatch(worker(code,'badUpdateConsumer'),/recordConsumer_/);for(const name of ['subset','duplicate','constant','unknown','intermediate','escaping','badEffect'])assert.doesNotMatch(worker(code,name),/recordConsumer_/,name);
});

test('saved EffectPure, foreign proofs, staged inputs and recursive contexts are refused before translation',()=>{
 const r=row(keys),count=local('count',0,int),original=local('original',2,r),changed=local('changed',3,r);
 const sum=[...keys.map(k=>access(changed,k)),...keys.map(k=>access(original,k))].reduce(add);
 const pending=typed(effect(int),e(new S.EffectPure(sum))),binding=call('step',[count,original],new C.Func([int,r],r),r);
 const tco=x=>new TcoExpr(null,S.functorBackendSyntax.map(tco)(x));
 const bound=insert(ordString)(localId(new Just('original'))(2))({name:'originalBox',goType:Go.TypeValue.value})(singleton(localId(new Just('count'))(0))({name:'countNative',goType:Go.TypeInt64.value}));
 const context={modNameStr:'VariableNames',recVars:[],bound,options:{inEffectBlock:false,isTail:false}};
 const plans=singleton('step')({keys,deltas:[1],counterFirst:true,recordType:r});
 const translate=()=>{throw new Error('a rejected candidate must not translate or consume IDs');};
 const attempt=(value= binding,body=pending,ctx=context)=>emitConsumer(plans)(translate)(ctx)(0)(new Just('changed'))(3)(tco(value))(tco(body));
 assert.equal(attempt(),Nothing.value,'construction of a saved effect must remain outside its execution');
 assert.equal(attempt(binding,sum,{...context,recVars:['recursive']}),Nothing.value);
 const factory=call('identity',[count],new C.Func([int],int),int);
 assert.equal(attempt(call('step',[factory,original],new C.Func([int,r],r),r),sum),Nothing.value);
 assert.equal(attempt(typed(r,e(new S.UncurriedApp(variable('step',new C.Func([int,r],r)),[count,original]))),sum),Nothing.value);
 const remote=typed(r,e(new S.App(typed(new C.Func([int,r],r),e(new S.Var(new C.Qualified(new Just('External'),'step')))),[count,original])));
 assert.equal(attempt(remote,sum),Nothing.value);
});

test('generated guarded consumers preserve wide fields, malformed shapes, getter order and replay',t=>{
 const {specs,bindings}=fixtures(),code=generate(bindings),dir=mkdtempSync(join(tmpdir(),'gopurs-record-consumer-'));t.after(()=>rmSync(dir,{recursive:true,force:true}));mkdirSync(join(dir,'purescript'));mkdirSync(join(dir,'gopurs_runtime'));writeFileSync(join(dir,'go.mod'),'module gopurs/output\n\ngo 1.22\n');writeFileSync(join(dir,'gopurs_runtime/runtime.go'),runtimeGoCode);writeFileSync(join(dir,'purescript/generated.go'),code);
 const cases=specs.map(([name,fields])=>`{${JSON.stringify(name)},[]string{${fields.map(JSON.stringify).join(',')}},Call_VariableNames_fused_${name},Call_VariableNames_reference_${name}}`).join(',\n')+',\n'+`{"boxed",[]string{${keys.map(JSON.stringify).join(',')}},func(n int64,p r.Value)r.Value{return Call_VariableNames_boxed(r.Func(func(r.Value)r.Value{trace=append(trace,"count");return r.Int(n)}),p)},func(n int64,p r.Value)r.Value{return Call_VariableNames_boxedReference(r.Func(func(r.Value)r.Value{trace=append(trace,"count");return r.Int(n)}),p)}}`;
 writeFileSync(join(dir,'purescript/proof_test.go'),`package purescript
import("testing";"fmt";"reflect";r "gopurs/output/gopurs_runtime")
var trace []string
var sentinel=&struct{}{}
type getters struct{values map[string]r.Value;fail string;mutate bool}
func(g *getters)JSONLookup(key string)(any,bool){trace=append(trace,key);if key==g.fail{panic(sentinel)};v:=g.values[key];if g.mutate{g.values[key]=r.Int(v.IntVal+1)};return v,true}
func(g *getters)JSONLength()int{return len(g.values)}
func(g *getters)JSONEntry(index int)(string,any){panic("unused enumeration")}
func makeRecord(keys []string,seed int64,shape int)(r.Value,*getters){labels:=append([]string(nil),keys...);values:=make([]r.Value,len(keys));m:=map[string]r.Value{};for i,k:=range keys{values[i]=r.Int(seed+int64(i));m[k]=values[i]};switch shape{case1:for i,j:=0,len(keys)-1;i<j;i,j=i+1,j-1{labels[i],labels[j]=labels[j],labels[i];values[i],values[j]=values[j],values[i]};case2:return r.Record(m),nil;case3:labels[len(labels)-1]=labels[0];case4:labels[len(labels)-1]="missing";case5:values=values[:len(values)-1];case6:values=nil;case7:g:=&getters{values:m};return r.Box(g),g;case8:g:=&getters{values:m,fail:keys[len(keys)-1]};return r.Box(g),g;case9:g:=&getters{values:m,mutate:true};return r.Box(g),g};if shape==0 {if len(keys)==1{return r.RecordDict1(keys[0],values[0]),nil};if len(keys)==2{return r.RecordDict2(keys[0],keys[1],values[0],values[1]),nil}};return r.RecordDict(labels,values),nil}
func observe(fn func()r.Value)(value int64,panicked bool,panicText string){defer func(){if p:=recover();p!=nil{panicked=true;if p==sentinel{panicText="sentinel"}else{panicText=fmt.Sprint(p)}}}();value=fn().IntVal;return}
func TestConsumers(t *testing.T){checks:=0;cases:=[]struct{name string;keys []string;fused,reference func(int64,r.Value)r.Value}{${cases}}
 for _,c:=range cases{for _,seed:=range []int64{-9223372036854775808,-2147483649,-2147483648,0,2147483647,4294967296,9223372036854775807}{for shape:=0;shape<10;shape++{for _,count:=range []int64{0,1,2,7,1000,2147483647,4294967297,4294967299,-4294967295}{
  left,_:=makeRecord(c.keys,seed,shape);right,_:=makeRecord(c.keys,seed,shape);if shape<7{right=left};callsL,callsR:=0,0
  effectL:=r.Func(func(r.Value)r.Value{callsL++;trace=append(trace,"input");return left});effectR:=r.Func(func(r.Value)r.Value{callsR++;trace=append(trace,"input");return right})
  fused:=c.fused(count,effectL);reference:=c.reference(count,effectR)
  for replay:=0;replay<2;replay++{trace=nil;a,ap,at:=observe(func()r.Value{return r.Apply(fused,r.Value{})});ta:=append([]string(nil),trace...);trace=nil;b,bp,bt:=observe(func()r.Value{return r.Apply(reference,r.Value{})});if a!=b||ap!=bp||at!=bt||!reflect.DeepEqual(ta,trace){t.Fatalf("%s count%d seed%d shape%d got(%d,%v,%q,%v) want(%d,%v,%q,%v)",c.name,count,seed,shape,a,ap,at,ta,b,bp,bt,trace)};checks++}
  if callsL!=2||callsR!=2{t.Fatal("saved action did not replay input exactly once")}
 }}}};t.Logf("%d paired generated consumer observations",checks)
}
` .replaceAll('case1:','case 1:').replaceAll('case2:','case 2:').replaceAll('case3:','case 3:').replaceAll('case4:','case 4:').replaceAll('case5:','case 5:').replaceAll('case6:','case 6:').replaceAll('case7:','case 7:').replaceAll('case8:','case 8:').replaceAll('case9:','case 9:'));
 const result=spawnSync('go',['test','-count=1','-timeout=20s','-v','./purescript'],{cwd:dir,encoding:'utf8',timeout:90000,env:{...process.env,GOWORK:'off'}});assert.ifError(result.error);assert.equal(result.status,0,result.stdout+result.stderr);console.log(result.stdout.trim());
});
