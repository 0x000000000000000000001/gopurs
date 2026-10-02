import assert from 'node:assert/strict';
import test from 'node:test';
import { monadIdentity } from '../output/Data.Identity/index.js';
import { Cons, Nil } from '../output/Data.List.Types/index.js';
import * as Map from '../output/Data.Map/index.js';
import { Just, Nothing } from '../output/Data.Maybe/index.js';
import { Left } from '../output/Data.Either/index.js';
import { ordString } from '../output/Data.Ord/index.js';
import { Tuple } from '../output/Data.Tuple/index.js';
import { unfoldableArray } from '../output/Data.Unfoldable/index.js';
import * as Aff from '../output/Effect.Aff/index.js';
import * as C from '../output/PureScript.Backend.Optimizer.CoreFn/index.js';
import * as PBO from '../output/PureScript.Backend.Optimizer.Monomorphize/index.js';
import { buildGlobalTypes } from '../output/Gopurs.GlobalTypes/index.js';
import { monomorphizeModules, monomorphizeModulesWith } from '../output/Gopurs.Monomorphization/index.js';
import { runPreparationJobs } from '../output/Gopurs.Preparation/index.js';

const mapOf=entries=>entries.reduce((map,[key,value])=>Map.insert(ordString)(key)(value)(map),Map.empty);
const keys=map=>Map.toUnfoldable(unfoldableArray)(map).map(pair=>pair.value0);
const list=values=>values.reduceRight((tail,value)=>new Cons(value,tail),Nil.value);
const array=values=>values instanceof Cons?[values.value0,...array(values.value1)]:[];
const int=C.Int.value, variable=new C.TypeVar('a');
const mono=new C.Func([int],int), poly=new C.ForAll(['a'],new C.Func([variable],variable));
const ann=type=>({span:C.emptySpan,meta:Nothing.value,sourceUsage:Nothing.value,type:type===null?Nothing.value:new Just(type)});
const ref=(owner,name,type=poly)=>new C.ExprVar(ann(type),new C.Qualified(owner===null?Nothing.value:new Just(owner),name));
const literal=n=>new C.ExprLit(ann(int),new C.LitInt(n));
const app=(fn,value,type=int)=>new C.ExprApp(ann(type),fn,value);
const lambda=(name,body,type=poly)=>new C.ExprAbs(ann(type),name,body);
const binding=(name,type,body)=>new C.Binding(ann(type),name,body);
const nonrec=(name,type,body)=>new C.NonRec(binding(name,type,body));
const moduleOf=(name,decls,foreign=[])=>({
 name,path:`${name}.purs`,span:C.emptySpan,imports:[],comments:[],decls,
 exports:decls.flatMap(group=>group instanceof C.NonRec?[group.value0.value1]:group.value0.map(b=>b.value1)),
 reExports:Map.empty,dataDecls:[],classDecls:[],foreign:mapOf(foreign),
});
const identity=()=>nonrec('identity',poly,lambda('value',ref(null,'value',variable)));
const program=()=>[
 moduleOf('Generic',[identity()]),
 moduleOf('Wrapper',[nonrec('wrap',mono,lambda('value',app(ref('Generic','identity'),ref(null,'value',int)),mono))]),
 moduleOf('Caller',[nonrec('answer',int,app(ref('Wrapper','wrap',mono),literal(42)))]),
];
const declarations=modules=>modules.flatMap(mod=>mod.decls.flatMap(group=>
 (group instanceof C.NonRec?[group.value0]:group.value0).map(b=>`${mod.name}.${b.value1}`)));
const specialized=(modules,name)=>declarations(modules).filter(value=>value.startsWith(name+'__'));
const facts=value=>value&&typeof value==='object'?
 [...(Object.hasOwn(value,'sourceUsage')?[value.sourceUsage]:[]),...Object.values(value).flatMap(facts)]:[];
function run(modules,types=buildGlobalTypes(modules),collector=PBO.transitiveCollect){
 const before=JSON.stringify(modules);
 const result=array(monomorphizeModulesWith(monadIdentity)(collector)(types)(list(modules)));
 assert.equal(JSON.stringify(modules),before,'source modules remain immutable');
 assert.deepEqual(result.map(mod=>mod.name),modules.map(mod=>mod.name),'caller-supplied module order is preserved');
 return result;
}
const runAff=aff=>new Promise((resolve,reject)=>Aff.runAff(result=>()=>result instanceof Left?reject(result.value0):resolve(result.value0))(aff)());

test('source usage identities are invalidated before collection even when no specialization survives',()=>{
 const original=program().slice(0,1);
 const body=original[0].decls[0].value0.value2;
 const source={moduleName:'Generic',bindingId:4};
 body.value0.sourceUsage=new Just({bindingUsage:new Just({binding:source,maxUses:new Just(1),hasEscapingUseContext:new Just(false)}),variableUse:Nothing.value});
 body.value2.value0.sourceUsage=new Just({bindingUsage:Nothing.value,variableUse:new Just({binding:source,lastLocalUse:new Just(true)})});
 let calls=0;
 const result=run(original,Map.empty,ast=>raw=>{
  calls++;
  assert.ok(facts(ast).every(value=>value instanceof Nothing));
  assert.deepEqual(Map.lookup(ordString)('Generic.identity')(ast).value0.value0,original[0].decls[0].value0.value0);
  assert.deepEqual(keys(raw),[]);
  return Map.empty;
 });
 assert.equal(calls,1);
 assert.ok(facts(result).every(value=>value instanceof Nothing));
 assert.ok(facts(original).some(value=>value instanceof Just));
 assert.deepEqual(result[0].decls[0].value0.value2.value0.type,new Just(poly));
});

test('monomorphic definitions stay visible during transitive collection but do not receive clones',()=>{
 const modules=program();
 let observed=false;
 const result=run(modules,buildGlobalTypes(modules),ast=>raw=>{
  observed=true;
  assert.ok(keys(ast).includes('Wrapper.wrap'));
  assert.ok(keys(raw).includes('Wrapper.wrap'));
  const transitive=PBO.transitiveCollect(ast)(raw);
  assert.ok(keys(transitive).includes('Generic.identity'));
  return transitive;
 });
 assert.equal(observed,true);
 assert.ok(specialized(result,'Generic.identity').length>0);
 assert.deepEqual(specialized(result,'Wrapper.wrap'),[]);
 const wrapper=result.find(mod=>mod.name==='Wrapper').decls[0].value0;
 assert.equal(wrapper.value1,'wrap');
 assert.deepEqual(wrapper.value0.type,new Just(mono),'the original wrapper ABI is retained');
});

test('post-collection admission uses original global types, including explicit Any and opaque runtime names',()=>{
 const modules=program();
 const admitted=[poly,new C.Array(variable),new C.Record(new C.Row([],new Just(new C.TypeVar('r')))),
  new C.ConstrainedType([new Tuple(['Example','Class'],[variable])],mono),new C.ADT('Example.Box',['Example','Box'],[variable])];
 const rejected=[C.Any.value,mono,new C.ForAll(['unused'],mono),new C.TypeVar('Opaque'),new C.TypeVar('gopurs_runtime.Value')];
 for(const type of [...admitted,...rejected]){
  let collected=false;
  const result=run(modules,mapOf([['Generic.identity',type]]),ast=>raw=>{
   collected=true;
   assert.ok(keys(raw).includes('Generic.identity'));
   return PBO.transitiveCollect(ast)(raw);
  });
  assert.equal(collected,true);
  assert.equal(specialized(result,'Generic.identity').length>0,admitted.includes(type));
 }
 assert.deepEqual(specialized(run(modules,Map.empty),'Generic.identity'),[],'a missing source type never justifies cloning');
});

test('the negate intrinsic is the exact early AST barrier while other same-named globals remain available',()=>{
 const negate=()=>nonrec('negate',poly,lambda('value',ref(null,'value',variable)));
 const modules=[moduleOf('Data.Ring',[negate()]),moduleOf('Other',[negate()]),moduleOf('Caller',[
  nonrec('ring',int,app(ref('Data.Ring','negate'),literal(1))),
  nonrec('other',int,app(ref('Other','negate'),literal(1))),
 ])];
 const result=run(modules,buildGlobalTypes(modules),ast=>raw=>{
  assert.equal(keys(ast).includes('Data.Ring.negate'),false);
  assert.ok(keys(ast).includes('Other.negate'));
  assert.ok(keys(raw).includes('Data.Ring.negate'),'its typed calls are still collected');
  return PBO.transitiveCollect(ast)(raw);
 });
 assert.deepEqual(specialized(result,'Data.Ring.negate'),[]);
 assert.ok(specialized(result,'Other.negate').length>0);
});

test('foreign wrappers and row-only readers retain their bodies for collection and share their original workers',()=>{
 const row=new C.Record(new C.Row([new Tuple('id',int)],new Just(new C.TypeVar('r'))));
 const readerType=new C.ForAll(['r'],new C.Func([row],int));
 const concrete=new C.Record(new C.Row([new Tuple('id',int)],Nothing.value));
 const modules=[moduleOf('Foreign',[
  nonrec('forward',poly,lambda('value',app(ref('Foreign','ffi'),ref(null,'value',variable),variable))),
  nonrec('read',readerType,lambda('row',new C.ExprAccessor(ann(int),ref(null,'row',row),'id'),readerType)),
 ],[['ffi',new Just(poly)]]),moduleOf('Caller',[
  nonrec('ffi',int,app(ref('Foreign','forward'),literal(7))),
  nonrec('row',int,app(ref('Foreign','read',readerType),new C.ExprLit(ann(concrete),new C.LitRecord([new C.Prop('id',literal(8))])))),
 ])];
 const result=run(modules,buildGlobalTypes(modules),ast=>raw=>{
  assert.ok(keys(ast).includes('Foreign.forward'));
  assert.ok(keys(ast).includes('Foreign.read'));
  assert.equal(keys(ast).includes('Foreign.ffi'),false);
  for(const name of ['Foreign.forward','Foreign.read','Foreign.ffi'])assert.ok(keys(raw).includes(name),name);
  return PBO.transitiveCollect(ast)(raw);
 });
 for(const name of ['Foreign.forward','Foreign.read','Foreign.ffi'])assert.deepEqual(specialized(result,name),[]);
});

test('recursive groups are indexed and rewritten without reordering modules',()=>{
 const modules=program();
 modules[0].decls=[new C.Rec([modules[0].decls[0].value0])];
 const ordered=[modules[2],modules[0],modules[1]];
 const result=run(ordered,buildGlobalTypes(ordered),ast=>raw=>{
  assert.ok(keys(ast).includes('Generic.identity'));
  return PBO.transitiveCollect(ast)(raw);
 });
 assert.ok(specialized(result,'Generic.identity').length>0);
});

test('pure and Aff preparation produce identical specialized modules at every supported job bound',async()=>{
 const modules=program(),types=buildGlobalTypes(modules),input=list(modules);
 const expected=array(monomorphizeModules(types)(input));
 for(const jobs of [-1,0,1,3,8,99]){
  const collect=ast=>raw=>PBO.transitiveCollectWith(Aff.monadRecAff)(runPreparationJobs(jobs))(ast)(raw);
  const result=await runAff(monomorphizeModulesWith(Aff.monadAff)(collect)(types)(input));
  assert.deepEqual(array(result),expected,`jobs=${jobs}`);
 }
});
