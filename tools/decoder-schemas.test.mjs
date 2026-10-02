import assert from 'node:assert/strict';
import test from 'node:test';
import { mkdtempSync, mkdirSync, readFileSync, writeFileSync, rmSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join } from 'node:path';
import { spawnSync } from 'node:child_process';
import * as Map from '../output/Data.Map/index.js';
import { ordString } from '../output/Data.Ord/index.js';
import { Nothing, Just } from '../output/Data.Maybe/index.js';
import { Tuple } from '../output/Data.Tuple/index.js';
import * as C from '../output/PureScript.Backend.Optimizer.CoreFn/index.js';
import * as S from '../output/PureScript.Backend.Optimizer.Syntax/index.js';
import { NeutralExpr as E } from '../output/PureScript.Backend.Optimizer.Semantics/index.js';
import { specializeDecoderSchemas } from '../output/Gopurs.DecoderSchemas/index.js';
import { printGoDecl } from '../output/Gopurs.Printer/index.js';
import { empty as emptySet } from '../output/Data.Set/index.js';
import { translate } from '../output/Gopurs.CodeGen/index.js';
import { withReboxFields } from './codegen-metadata.mjs';

const global=(m,n)=>new E(new S.Var(new C.Qualified(m===null?Nothing.value:new Just(m),n)));
const cls=n=>global('Data.Argonaut.Decode.Class',n);
const app=(fn,...args)=>new E(new S.App(fn,args));
const erased=new E(S.PrimUndefined.value);
const literal=s=>new E(new S.Lit(new C.LitString(s)));
const symbol=s=>new E(new S.Lit(new C.LitRecord([new C.Prop('reflectSymbol',new E(new S.Abs([new Tuple(Nothing.value,0)],literal(s))))])));
const int=cls('decodeJsonInt'), str=cls('decodeJsonString');
const field=(key,decoder,tail=cls('gDecodeJsonNil'))=>app(cls('gDecodeJsonCons'),app(cls('decodeFieldId'),decoder),tail,symbol(key),erased,erased);
const record=row=>app(cls('decodeRecord'),row,erased);
const method=dict=>app(cls('decodeJson'),dict);
const metadata={globalTypes:Map.insert(ordString)('Data.Argonaut.Decode.Internal.Record.schemaDecoderABI1')(C.Int.value)(Map.empty)};
const textMetadata={globalTypes:Map.insert(ordString)('Data.Argonaut.Decode.Parser.textDecoderABI1')(C.Int.value)(metadata.globalTypes)};
function run(dict,extra=[],recursive=false,meta=metadata){
 const mod={name:'Probe',bindings:[{recursive,bindings:[...extra,new Tuple('decode',method(dict))]}]};
 const snapshot=JSON.stringify(mod);const result=specializeDecoderSchemas(meta)(mod);
 assert.equal(JSON.stringify(mod),snapshot);
 return {...result,code:result.declarations.map(printGoDecl).join('\n')};
}
test('resolved standard dictionaries emit direct workers and one original source getter',()=>{
 const dict=record(field('x',int,field('y',app(cls('decodeArray'),str))));
 const result=run(dict);
 assert.match(result.code,/object.Lookup\("x"\)/);
 assert.match(result.code,/intFromNumber/);
 assert.match(result.code,/RecordDict2\("y", "x", value1, value0\)/);
 assert.match(result.code,/argonautCompileSchema/);
 assert.equal(result.module.bindings.at(-1).bindings.length,1);
 assert.ok(JSON.stringify(result.module.bindings.at(-1)).includes('decodeJsonInt'));
});
test('text workers require their ABI and share the original schema proof and constructor sources',()=>{
 const dict=record(field('rows',app(cls('decodeArray'),record(field('text',str)))));
 assert.doesNotMatch(run(dict).code,/argonautCompileTextSchema|argonautTextCursor/);
 const result=run(dict,[],false,textMetadata);
 assert.match(result.code,/argonautCompileTextSchema/);
 assert.match(result.code,/raw argonautTextCursor/);
 assert.match(result.code,/at = raw.document.tokens\[at\].next/);
 assert.equal(result.module.bindings.at(-1).bindings.length,1);
});
test('aliases and partial standard dictionary applications resolve without changing construction',()=>{
 const result=run(global('Probe','dict'),[
  new Tuple('cons',app(cls('gDecodeJsonCons'),app(cls('decodeFieldId'),int))),
  new Tuple('key',symbol('actual-label')),
  new Tuple('dict',record(app(global('Probe','cons'),cls('gDecodeJsonNil'),global('Probe','key'),erased,erased)))
 ]);
 assert.match(result.code,/object.Lookup\("actual-label"\)/);
});
test('constant duplicate labels preserve reverse insertion and head precedence',()=>{
 const result=run(record(field('same',int,field('same',str))));
 assert.match(result.code,/RecordDict1\("same", value0\)/);
});
test('custom container methods retain ordinary dispatch, no type-derived decoder',()=>{
 const result=run(record(field('events',app(cls('decodeArray'),global('Opaque','custom')))));
 assert.match(result.code,/argonautSchemaElement/);
 assert.match(result.code,/argonautSchemaCustom/);
 assert.match(result.code,/argonautSchemaField/);
 assert.match(result.code,/_accepts/);
 const text=run(record(field('events',app(cls('decodeArray'),global('Opaque','custom')))),[],false,textMetadata);
 assert.doesNotMatch(text.code,/argonautCompileTextSchema|argonautTextCursor/);
});

const local=i=>new E(new S.Local(Nothing.value,i));
const either=n=>new C.Qualified(new Just('Data.Either'),n);
const ctor=(name,...fields)=>new E(new S.CtorSaturated(either(name),C.SumType.value,'Either',name,fields.map((v,i)=>new Tuple('value'+i,v))));
const access=(name,i)=>new E(new S.Accessor(local(i),new S.GetCtorField(either(name),C.SumType.value,'Either',name,'value0',0)));
const is=(name,i)=>new E(new S.PrimOp(new S.Op1(new S.OpIsTag(either(name)),local(i))));
const bind=(i,value,next)=>new E(new S.Let(Nothing.value,i,value,new E(new S.Branch([
 new S.Pair(is('Left',i),ctor('Left',access('Left',i))),new S.Pair(is('Right',i),next)
],new E(new S.Fail('Failed pattern match'))))));
const typed=(type,value)=>new E(new S.Typed(type,value));
const array=values=>new E(new S.Lit(new C.LitArray(values)));
const dictionary=body=>new E(new S.Lit(new C.LitRecord([
 new C.Prop('decodeJson',new E(new S.Abs([new Tuple(Nothing.value,0)],body))),
])));
const borrow=app(global('Data.Argonaut.Decode.Internal.Record','borrowObject'),
 new E(new S.Accessor(app(cls('decodeForeignObject'),cls('decodeJsonJson')),new S.GetProp('decodeJson'))),local(0));
const getField=(key,decoder=str,reader='getField')=>app(global('Data.Argonaut.Decode.Decoders',reader),method(decoder),access('Right',1),literal(key));
const derived=(continuation,meta=textMetadata)=>run(app(cls('decodeArray'),global('Probe','custom')),
 [new Tuple('custom',dictionary(bind(1,borrow,continuation)))],false,meta);
const hasProgram=result=>result.module.bindings.at(-1).bindings.some(pair=>pair.value0.endsWith('_construct'));
test('custom bodies specialize only proven reads, forwarding errors and actual constructors',()=>{
 const object=app(cls('decodeForeignObject'),cls('decodeJsonJson'));
 const borrow=app(global('Data.Argonaut.Decode.Internal.Record','borrowObject'),new E(new S.Accessor(object,new S.GetProp('decodeJson'))),local(0));
 const get=app(global('Data.Argonaut.Decode.Decoders','getField'),method(str),access('Right',1),literal('payload'));
 const body=bind(1,borrow,bind(2,get,ctor('Right',access('Right',2))));
 const dictionary=b=>new E(new S.Lit(new C.LitRecord([new C.Prop('decodeJson',new E(new S.Abs([new Tuple(Nothing.value,0)],b)))])));
 const runBody=b=>run(record(field('entries',app(cls('decodeArray'),global('Probe','custom')))),[new Tuple('custom',dictionary(b))]);
 const result=runBody(body);
 assert.match(result.code,/object.Lookup\("payload"\)/);
 assert.match(result.code,/_construct\(\)/);
 assert.equal(result.module.bindings.at(-1).bindings.length,2);
 // Replacing the Left continuation with an opaque call invalidates the proof.
 // Use the real IR classes when reconstructing the changed outer continuation.
 const broken=new E(new S.Let(Nothing.value,1,borrow,new E(new S.Branch([
   new S.Pair(is('Left',1),app(global('Opaque','changeError'),access('Left',1))),new S.Pair(is('Right',1),bind(2,get,ctor('Right',access('Right',2))))
 ],new E(new S.Fail('Failed pattern match'))))));
 assert.doesNotMatch(runBody(broken).code,/object.Lookup\("payload"\)/);
});
test('dynamic symbol callbacks, unknown builders, local captures and recursive scopes are not specialized',()=>{
 const dynamic=new E(new S.Lit(new C.LitRecord([new C.Prop('reflectSymbol',global('Probe','reflect'))])));
 const dynamicRow=app(cls('gDecodeJsonCons'),app(cls('decodeFieldId'),int),cls('gDecodeJsonNil'),dynamic,erased,erased);
 const invalid=[record(dynamicRow),app(global('Other','decodeRecord'),field('x',int),erased),
  app(global(null,'decodeRecord'),field('x',int),erased),app(cls('decodeArray'),app(global('Probe','builder'),erased)),
  app(cls('decodeArray'),new E(new S.Local(Nothing.value,0)))];
 for(const dict of invalid)assert.equal(run(dict).declarations.length,0);
 assert.equal(run(record(field('x',int)),[],true).declarations.length,0);
 assert.equal(run(record(field('x',int)),[],false,{globalTypes:Map.empty}).declarations.length,0);
 assert.equal(run(record(field('x',global('Probe','alias'))),[new Tuple('alias',global('Probe','recursive'))]).declarations.length,0);
});

test('annotations, type applications and generated-name collisions are preserved',()=>{
 const source=new E(new S.Typed(C.Any.value,new E(new S.TypeApp(method(record(field('x',int))),C.Any.value))));
 const result=specializeDecoderSchemas(metadata)({name:'Probe',bindings:[{recursive:false,bindings:[new Tuple('__json_schema_0_decode',literal('user')),new Tuple('decode',source)]}]});
 assert.match(result.declarations.map(printGoDecl).join('\n'),/__json_schema_1_decode/);
 const replacement=result.module.bindings[0].bindings[1].value1;
 assert.ok(replacement instanceof S.Typed);
 assert.ok(replacement.value1 instanceof S.TypeApp);
  assert.equal(result.module.bindings[0].bindings[0].value0,'__json_schema_0_decode');
});

test('only an exact qualified unary decodeJson method is admitted, even with the text ABI',()=>{
 const dict=record(field('x',int));
 const textOnly={globalTypes:Map.insert(ordString)('Data.Argonaut.Decode.Parser.textDecoderABI1')(C.Int.value)(Map.empty)};
 assert.equal(run(dict,[],false,textOnly).declarations.length,0);
 for(const expr of [app(global('Other','decodeJson'),dict),app(global(null,'decodeJson'),dict),
  app(cls('decodeJson'),dict,erased),new E(new S.UncurriedApp(cls('decodeJson'),[dict]))]){
  const mod={name:'Probe',bindings:[{recursive:false,bindings:[new Tuple('decode',expr)]}]};
  const result=specializeDecoderSchemas(textMetadata)(mod);
  assert.deepEqual(result.module.bindings,mod.bindings);
  assert.deepEqual(result.declarations,[]);
 }
});

test('record proofs require erased proxies and a single constant symbol method',()=>{
 const row=(key=symbol('x'),cons=erased,lacks=erased)=>app(cls('gDecodeJsonCons'),app(cls('decodeFieldId'),int),cls('gDecodeJsonNil'),key,cons,lacks);
 const twoArgs=new E(new S.Lit(new C.LitRecord([new C.Prop('reflectSymbol',new E(new S.Abs([
  new Tuple(Nothing.value,0),new Tuple(Nothing.value,1),
 ],literal('x'))))])));
 const extraMember=new E(new S.Lit(new C.LitRecord([
  new C.Prop('reflectSymbol',new E(new S.Abs([new Tuple(Nothing.value,0)],literal('x')))),new C.Prop('extra',erased),
 ])));
 for(const dict of [app(cls('decodeRecord'),row(),literal('proxy')),record(row(symbol('x'),literal('cons'))),
  record(row(symbol('x'),erased,literal('lacks'))),record(row(twoArgs)),record(row(extraMember))]){
  assert.equal(run(dict).declarations.length,0);
 }
 assert.match(run(app(cls('decodeRecord'),row(typed(C.Any.value,symbol('x'))),typed(C.Any.value,erased))).code,/object.Lookup\("x"\)/);
});

test('schema budgets keep scalar, untagged and oversized roots on their original path',()=>{
 for(const dict of [int,record(cls('gDecodeJsonNil')),global('Opaque','custom'),app(cls('decodeArray'),global('Opaque','custom'))]){
  assert.equal(run(dict).declarations.length,0);
 }
 const tree=depth=>depth===0?int:record(field('left',tree(depth-1),field('right',tree(depth-1))));
 const boundary=app(cls('decodeJsonMaybe'),tree(6));
 assert.ok(run(boundary).declarations.length>0,'128 weighted schema nodes remain admissible');
 assert.equal(run(app(cls('decodeJsonMaybe'),boundary)).declarations.length,0);
});

test('alias resolution is bounded, qualified, and excludes missing or cyclic local definitions',()=>{
 const aliases=count=>Array.from({length:count},(_,i)=>new Tuple(`alias${i}`,i+1===count?int:global('Probe',`alias${i+1}`)));
 const dict=app(cls('decodeArray'),global('Probe','alias0'));
 assert.ok(run(dict,aliases(62)).declarations.length>0);
 assert.equal(run(dict,aliases(63)).declarations.length,0);
 assert.equal(run(dict,[new Tuple('alias0',global('Probe','alias1')),new Tuple('alias1',global('Probe','alias0'))]).declarations.length,0);
 assert.equal(run(app(cls('decodeArray'),global(null,'alias0')),aliases(1)).declarations.length,0);
});

test('rejected scopes allow independent children while recursive groups and LetRec stop rewriting',()=>{
 const source=method(record(field('x',int)));
 const recursive=new E(new S.LetRec(0,[new Tuple(new Just('loop'),source)],source));
 const effect=new E(new S.EffectPure(source));
 const mod={name:'Probe',bindings:[
  {recursive:true,bindings:[new Tuple('recursive',source),new Tuple('recursiveDict',record(field('x',int)))]},
  {recursive:false,bindings:[new Tuple('localRec',recursive),new Tuple('effect',effect),
   new Tuple('missing',method(global('Probe','recursiveDict')))]},
 ]};
 const result=specializeDecoderSchemas(metadata)(mod);
 assert.deepEqual(result.module.bindings[0],mod.bindings[0]);
 assert.deepEqual(result.module.bindings[1].bindings[0],mod.bindings[1].bindings[0]);
 assert.deepEqual(result.module.bindings[1].bindings[1].value1,new E(new S.EffectPure(global('Probe','__json_schema_0'))));
 assert.deepEqual(result.module.bindings[1].bindings[2],mod.bindings[1].bindings[2]);
 assert.deepEqual(result.module.bindings.at(-1).bindings,[new Tuple('__json_schema_0_source',typed(C.Any.value,source))]);
});

test('worker families reserve descendant names in all groups and publish untouched sources in visit order',()=>{
 const source=typed(C.Any.value,new E(new S.TypeApp(method(record(field('x',int))),C.Int.value)));
 const mod={name:'Probe',bindings:[
  {recursive:true,bindings:[new Tuple('__json_schema_0_decode_item_text',literal('reserved'))]},
  {recursive:false,bindings:[new Tuple('pair',array([source,source])),new Tuple('__json_schema_2_source',erased)]},
 ]};
 const result=specializeDecoderSchemas(textMetadata)(mod);
 assert.deepEqual(result.module.bindings[0],mod.bindings[0]);
 assert.deepEqual(result.module.bindings[1].bindings[0].value1,array([1,3].map(n=>typed(C.Any.value,
  new E(new S.TypeApp(global('Probe',`__json_schema_${n}`),C.Int.value))))));
 assert.deepEqual(result.module.bindings.at(-1).bindings,[1,3].map(n=>new Tuple(`__json_schema_${n}_source`,typed(C.Any.value,source))));
});

test('derived bodies retain required, optional and nullable reader policies and numeric specializations',()=>{
 for(const reader of ['getField','getField__17','getFieldOptional','getFieldOptional__2',"getFieldOptional'","getFieldOptional'__3"]){
  const result=derived(bind(2,getField('payload',str,reader),ctor('Right',access('Right',2))));
  assert.ok(hasProgram(result),reader);
  assert.match(result.code,/argonautCompileTextSchema/);
  assert.equal(result.code.includes('field_2 = plan.mkNothing()'),reader.startsWith('getFieldOptional'));
  assert.equal(result.code.includes('!present || domNull(input)'),reader.startsWith("getFieldOptional'"));
 }
 for(const reader of ['getField__','getField__x','getField__2x','other']){
  const result=derived(bind(2,getField('payload',str,reader),ctor('Right',access('Right',2))));
  assert.equal(hasProgram(result),false,reader);
  assert.doesNotMatch(result.code,/argonautCompileTextSchema/);
 }
});

test('derived reads require complete dictionaries, literal keys and the proven borrowed object',()=>{
 const changedObject=app(global('Data.Argonaut.Decode.Decoders','getField'),method(str),local(0),literal('payload'));
 const dynamicKey=app(global('Data.Argonaut.Decode.Decoders','getField'),method(str),access('Right',1),global('Probe','key'));
 for(const read of [changedObject,dynamicKey,getField('payload',global('Opaque','custom')),
  getField('payload',app(cls('decodeArray'),global('Opaque','custom')))]){
  const result=derived(bind(2,read,ctor('Right',access('Right',2))));
  assert.equal(hasProgram(result),false);
  assert.equal(result.declarations.length,0);
  assert.doesNotMatch(result.code,/argonautCompileTextSchema/);
 }
});

test('each derived read must forward its exact Left payload, including through aliases',()=>{
 const read=getField('payload');
 const branch=failure=>new E(new S.Branch([
  new S.Pair(is('Left',2),failure),new S.Pair(is('Right',2),ctor('Right',access('Right',2))),
 ],new E(new S.Fail('unreachable'))));
 const alias=new E(new S.Let(Nothing.value,5,access('Left',2),ctor('Left',local(5))));
 for(const failure of [local(2),ctor('Left',access('Left',2)),alias]){
  assert.ok(hasProgram(derived(new E(new S.Let(Nothing.value,2,read,branch(failure))))));
 }
 for(const failure of [ctor('Left',literal('changed')),ctor('Left',access('Left',1)),
  app(global('Opaque','changeError'),access('Left',2))]){
  assert.equal(hasProgram(derived(new E(new S.Let(Nothing.value,2,read,branch(failure))))),false);
 }
});

test('derived producers cannot shadow proven levels or return opaque computations',()=>{
 assert.equal(hasProgram(derived(bind(1,getField('payload'),ctor('Right',access('Right',1))))),false);
 for(const value of [app(global('Opaque','construct'),access('Right',2)),local(0),access('Right',1)]){
  assert.equal(hasProgram(derived(bind(2,getField('payload'),ctor('Right',value)))),false);
 }
});

test('derived constructors retain annotations and capture distinct decoded levels in level order',()=>{
 const pair=new E(new S.CtorSaturated(new C.Qualified(new Just('Probe'),'Payload'),C.ProductType.value,'Payload','Payload',[
  new Tuple('value0',typed(C.String.value,access('Right',7))),
  new Tuple('value1',new E(new S.TypeApp(access('Right',2),C.Any.value))),
  new Tuple('value2',access('Right',7)),
 ]));
 const result=derived(bind(7,getField('first'),bind(2,getField('second'),ctor('Right',pair))));
 const construction=result.module.bindings.at(-1).bindings.find(pair=>pair.value0.endsWith('_construct'));
 assert.ok(construction);
 assert.deepEqual(construction.value1,typed(C.Any.value,new E(new S.Abs([
  new Tuple(Nothing.value,2),new Tuple(Nothing.value,7),
 ],new E(new S.CtorSaturated(new C.Qualified(new Just('Probe'),'Payload'),C.ProductType.value,'Payload','Payload',[
  new Tuple('value0',typed(C.String.value,local(7))),new Tuple('value1',new E(new S.TypeApp(local(2),C.Any.value))),new Tuple('value2',local(7)),
 ]))))));
 assert.match(result.code,/Apply2\(Get_.*_construct\(\), field_2, field_7\)/);
 assert.equal(result.module.bindings.at(-1).bindings.filter(pair=>pair.value0.endsWith('_construct')).length,1,'DOM and text share constructors');
});

function checkEmitted(result,testCode,constructors=''){
  const declarations=result.declarations.map(printGoDecl).filter(text=>text.startsWith('func ')).join('\n');
 const work=mkdtempSync(join(tmpdir(),'gopurs-schema-emission-'));
 try {
  mkdirSync(join(work,'output/gopurs_runtime'),{recursive:true});mkdirSync(join(work,'record'));
  writeFileSync(join(work,'go.mod'),'module gopurs\n\ngo 1.22\n');
  writeFileSync(join(work,'output/gopurs_runtime/runtime.go'),readFileSync(new URL('../runtime/runtime.go',import.meta.url)));
   writeFileSync(join(work,'record/record.go'),readFileSync(new URL('../../gopurs-argonaut-codecs/src/Data/Argonaut/Decode/Internal/Record.go',import.meta.url)));
   for(const [name,path] of [['parser','../../gopurs-argonaut-core/src/Data/Argonaut/Parser.go'],['text','../../gopurs-argonaut-codecs/src/Data/Argonaut/Decode/Parser.go']])
     writeFileSync(join(work,`record/${name}.go`),readFileSync(new URL(path,import.meta.url),'utf8').replace('package Parser','package Record'));
   writeFileSync(join(work,'record/record_test.go'),readFileSync(new URL('../../gopurs-argonaut-codecs/test/record-plan_test.go',import.meta.url)));
    writeFileSync(join(work,'record/workers.go'),`package Record\nimport "gopurs/output/gopurs_runtime"\n${declarations}`);
    if(constructors)writeFileSync(join(work,'record/constructors.go'),constructors.replace('package purescript','package Record'));
    writeFileSync(join(work,'record/emitted_test.go'),testCode);
    writeFileSync(join(work,'record/comparison_test.go'),`package Record
import "gopurs/output/gopurs_runtime"
func comparable(v gopurs_runtime.Value) any {
 switch v.Type {
 case gopurs_runtime.TypeString: return v.StrVal()
 case gopurs_runtime.TypeInt: return v.IntVal
 case gopurs_runtime.TypeBool: return v.BoolVal()
 case gopurs_runtime.TypeArray:
  out:=make([]any,gopurs_runtime.ArrayLength(v));for i:=range out {out[i]=comparable(gopurs_runtime.ArrayAccess(v,i))};return out
 }
 out:=map[string]any{};for k,x:=range gopurs_runtime.RecordToMap(v) {out[k]=comparable(x)};return out
}
`);
   const checked=spawnSync('go',['test','-v','-race','-count=1','-run','^Test(Emitted|CompiledSchema)','./record'],{cwd:work,encoding:'utf8',timeout:60_000,env:{...process.env,GOWORK:'off',GOMAXPROCS:'2'}});
   assert.ifError(checked.error);
   assert.equal(checked.status,0,(checked.stdout??'')+(checked.stderr??''));
   assert.match(checked.stdout,/--- PASS: TestEmitted/);
  } finally {rmSync(work,{recursive:true,force:true});}
}

test('emitted DOM and text workers preserve duplicate labels, exact errors and owned output',()=>{
  const result=run(record(field('same',int,field('same',app(cls('decodeJsonMaybe'),int),field('text',str)))),[],false,textMetadata);
  checkEmitted(result,`package Record
import ("testing"; "reflect"; "unsafe"; "gopurs/output/gopurs_runtime")
func TestEmittedDuplicateLabel(t *testing.T) {
 plan:=newErrorPlan(testErrorSupport)
 plan.fields=[]recordDecodeField{{kind:&fieldKind{tag:kindInt}},{kind:&fieldKind{tag:kindMaybe,inner:&fieldKind{tag:kindInt}}},{kind:&fieldKind{tag:kindString}}}
 kind:=&fieldKind{tag:kindRecord,plan:plan}
 if !Probe___json_schema_0_decode_accepts(kind) {t.Fatal("schema guard")}
 if !Probe___json_schema_0_decode_text_accepts(kind) {t.Fatal("text schema guard")}
 plan.fields[0].kind=nil
 if Probe___json_schema_0_decode_text_accepts(kind) {t.Fatal("text guard accepted an opaque replacement")}
 original:=gopurs_runtime.WithFunctionData(gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {panic("direct worker must not run")}),&typedKind{kind:kind,errors:plan,worker:Probe___json_schema_0_decode})
 rejected:=argonautCompileTextSchema(original,Probe___json_schema_0_decode_text,Probe___json_schema_0_decode_text_accepts)
 fallbackCalls:=0
 fallback:=gopurs_runtime.Func(func(text gopurs_runtime.Value) gopurs_runtime.Value {fallbackCalls++;if text.StrVal()!="original input" {t.Fatal("fallback input changed")};return gopurs_runtime.Int(42)})
 if got:=DecodeJsonStringImpl(1,fallback,rejected,"original input");got.IntVal!=42 || fallbackCalls!=1 {t.Fatal("opaque replacement did not use the whole original composition")}
 plan.fields[0].kind=&fieldKind{tag:kindInt}
 result:=Probe___json_schema_0_decode(plan,kind,map[string]any{"same":float64(9),"text":"owned"})
 if !result.ok || gopurs_runtime.RecordGet(result.value,"same").IntVal!=9 {t.Fatal("duplicate-label precedence")}
 for _,text:=range []string{
  \`{"same":0,"same":9,"text":"owned"}\`, \`{"same":null,"text":false}\`, \`{"text":false}\`,
  \`{"same":false,"text":"x"}\`, \`{"same":2147483648,"text":"x"}\`,
  \`{"same":9,"text":"é🙂"}\`, \`{"same":9,"text":"x","ignored":{"a":[1,true,null]}}\`,
 } {
  parsed,err:=argonautParseJSON(text);if err!=nil {t.Fatal(err)}
  input,ok:=argonautTextIndex(text);if !ok {t.Fatal("index")}
  want,got:=Probe___json_schema_0_decode(plan,kind,parsed),Probe___json_schema_0_decode_text(plan,kind,input)
  if want.ok!=got.ok {t.Fatal(text)}
  if want.ok {if !reflect.DeepEqual(comparable(want.value),comparable(got.value)) {t.Fatal("value",text)}} else if !reflect.DeepEqual(comparable(want.err),comparable(got.err)) {t.Fatal("error",text)}
 }
 bytes:=[]byte(\`{"same":9,"text":"owned-é🙂"}\`)
 input,ok:=argonautTextIndex(unsafe.String(unsafe.SliceData(bytes),len(bytes)));if !ok {t.Fatal("owned index")}
 saved:=Probe___json_schema_0_decode_text(plan,kind,input)
 for i:=range bytes {bytes[i]='x'}
 if !saved.ok || gopurs_runtime.RecordGet(saved.value,"text").StrVal()!="owned-é🙂" {t.Fatal("borrowed final string")}
}
`);
});

const equal=(left,right)=>new E(new S.PrimOp(new S.Op2(new S.OpStringOrd(S.OpEq.value),left,right)));
const mismatch=message=>ctor('Left',new E(new S.CtorSaturated(
 new C.Qualified(new Just('Data.Argonaut.Decode.Error'),'TypeMismatch'),C.SumType.value,'JsonDecodeError','TypeMismatch',
 [new Tuple('value0',literal(message))],
)));

test('derived choices follow actual ordered string comparisons and explicit mismatch constructors',()=>{
 const choose=condition=>bind(9,getField('kind'),new E(new S.Branch([
  new S.Pair(condition,ctor('Right',literal('first'))),
  new S.Pair(equal(access('Right',9),literal('same')),ctor('Right',literal('second'))),
 ],mismatch('unknown kind'))));
 const result=derived(choose(equal(access('Right',9),literal('same'))));
 assert.ok(hasProgram(result));
 assert.match(result.code,/field_9.StrVal\(\) == "same"/);
 assert.match(result.code,/plan.mkTypeMismatch\("unknown kind"\)/);
 const sources=result.module.bindings.at(-1).bindings.slice(1);
 assert.deepEqual(sources.map(pair=>pair.value0),[
  '__json_schema_0_decode_item_next_case0_construct','__json_schema_0_decode_item_next_case1_construct',
 ]);
 assert.deepEqual(sources.map(pair=>pair.value1),['first','second'].map(value=>typed(C.Any.value,literal(value))));
 for(const condition of [equal(literal('same'),access('Right',9)),equal(access('Right',9),global('Probe','label')),
  new E(new S.PrimOp(new S.Op2(new S.OpStringOrd(S.OpNotEq.value),access('Right',9),literal('same'))))]){
  assert.equal(hasProgram(derived(choose(condition))),false);
 }
});

test('emitted derived programs execute real constructor getters, ordered choices, optional reads and exact errors',()=>{
 const read=(decoder,reader='getField')=>bind(4,getField('payload',decoder,reader),ctor('Right',access('Right',4)));
 const result=derived(bind(9,getField('kind'),new E(new S.Branch([
  new S.Pair(equal(access('Right',9),literal('int')),read(int)),
  new S.Pair(equal(access('Right',9),literal('int')),ctor('Right',literal('unreachable'))),
  new S.Pair(equal(access('Right',9),literal('optional')),read(int,'getFieldOptional')),
  new S.Pair(equal(access('Right',9),literal('nullable')),read(int,"getFieldOptional'")),
  new S.Pair(equal(access('Right',9),literal('record')),read(record(field('name',str)))),
 ],mismatch('unknown kind')))));
 assert.ok(hasProgram(result));
 const compilerMetadata=withReboxFields({
  elidedCtors:emptySet,ctorTypes:Map.empty,pointerAdtPaths:Map.empty,pointerAdtNodes:emptySet,
  pointerAdtLeaves:Map.empty,enumAdts:emptySet,enumCtors:emptySet,globalFunctions:Map.empty,
  globalTypes:Map.empty,classDeclsFields:Map.empty,
 });
 const constructors=translate(compilerMetadata)({
  name:'Probe',bindings:[{recursive:false,bindings:result.module.bindings.at(-1).bindings.filter(pair=>pair.value0.endsWith('_construct'))}],
  comments:[],imports:emptySet,exports:emptySet,reExports:emptySet,dataTypes:Map.empty,dataDecls:[],
  classDecls:[],foreign:Map.empty,implementations:Map.empty,directives:Map.empty,
 });
 const worker=`Probe_${result.module.bindings[0].bindings.at(-1).value1.value0.value1}_decode`;
 checkEmitted(result,`package Record
import ("testing"; "reflect"; "gopurs/output/gopurs_runtime")
func TestEmittedDerived(t *testing.T) {
 plan:=newErrorPlan(testErrorSupport)
 kind:=&fieldKind{tag:kindArray}
 if !Probe___json_schema_0_decode_accepts(kind) || !Probe___json_schema_0_decode_text_accepts(kind) {t.Fatal("derived guard needs no opaque child tag")}
 tests:=[]struct {text string; ok bool; want any}{
  {\`[{"kind":"int","payload":7}]\`,true,[]any{int64(7)}},
  {\`[{"kind":"ignored","kind":"int","payload":1,"payload":9}]\`,true,[]any{int64(9)}},
  {\`[{"kind":"optional"},{"kind":"nullable","payload":null}]\`,true,[]any{comparable(plan.mkNothing()),comparable(plan.mkNothing())}},
  {\`[{"kind":"optional","payload":8}]\`,true,[]any{comparable(plan.mkJust(gopurs_runtime.Int(8)))}},
  {\`[{"kind":"record","payload":{"name":"é🙂"}}]\`,true,[]any{map[string]any{"name":"é🙂"}}},
 }
 for _,tc:=range tests {
  parsed,err:=argonautParseJSON(tc.text);if err!=nil {t.Fatal(err)}
  cursor,ok:=argonautTextIndex(tc.text);if !ok {t.Fatal("index")}
  for _,got:=range []argonautSchemaResult{Probe___json_schema_0_decode(plan,kind,parsed),Probe___json_schema_0_decode_text(plan,kind,cursor)} {
   if got.ok!=tc.ok || !reflect.DeepEqual(comparable(got.value),tc.want) {t.Fatalf("%s: %#v",tc.text,comparable(got.value))}
  }
 }
 errors:=[]struct {text string; err gopurs_runtime.Value}{
  {\`[{}]\`,plan.mkAtKey("kind",plan.missingValue)},
  {\`[{"kind":"other"}]\`,plan.mkTypeMismatch("unknown kind")},
  {\`[{"kind":"int"}]\`,plan.mkAtKey("payload",plan.missingValue)},
  {\`[{"kind":"optional","payload":null}]\`,plan.mkAtKey("payload",plan.mkTypeMismatch("Number"))},
  {\`[{"kind":"int","payload":1.5}]\`,plan.mkAtKey("payload",plan.mkTypeMismatch("Integer"))},
  {\`[{"kind":"record","payload":{}}]\`,plan.mkAtKey("payload",plan.mkAtKey("name",plan.missingValue))},
  {\`[false,{"kind":"int"}]\`,plan.mkTypeMismatch("Object")},
 }
 for _,tc:=range errors {
  parsed,err:=argonautParseJSON(tc.text);if err!=nil {t.Fatal(err)}
  cursor,ok:=argonautTextIndex(tc.text);if !ok {t.Fatal("index")}
  want:=comparable(plan.mkNamed("Array",plan.mkAtIndex(0,tc.err)))
  for _,got:=range []argonautSchemaResult{Probe___json_schema_0_decode(plan,kind,parsed),Probe___json_schema_0_decode_text(plan,kind,cursor)} {
   if got.ok || !reflect.DeepEqual(comparable(got.err),want) {t.Fatalf("%s: error %#v, want %#v",tc.text,comparable(got.err),want)}
  }
 }
}
`.replaceAll('Probe___json_schema_0_decode',worker),constructors);
});
