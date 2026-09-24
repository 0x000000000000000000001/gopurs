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
});

const local=i=>new E(new S.Local(Nothing.value,i));
const either=n=>new C.Qualified(new Just('Data.Either'),n);
const ctor=(name,...fields)=>new E(new S.CtorSaturated(either(name),C.SumType.value,'Either',name,fields.map((v,i)=>new Tuple('value'+i,v))));
const access=(name,i)=>new E(new S.Accessor(local(i),new S.GetCtorField(either(name),C.SumType.value,'Either',name,'value0',0)));
const is=(name,i)=>new E(new S.PrimOp(new S.Op1(new S.OpIsTag(either(name)),local(i))));
const bind=(i,value,next)=>new E(new S.Let(Nothing.value,i,value,new E(new S.Branch([
 new S.Pair(is('Left',i),ctor('Left',access('Left',i))),new S.Pair(is('Right',i),next)
],new E(new S.Fail('Failed pattern match'))))));
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

test('emitted duplicate-label workers compile and preserve the head value',()=>{
 const result=run(record(field('same',int,field('same',app(cls('decodeJsonMaybe'),int)))));
 const declarations=result.declarations.map(printGoDecl).filter(text=>text.startsWith('func ')).join('\n');
 const work=mkdtempSync(join(tmpdir(),'gopurs-schema-emission-'));
 try {
  mkdirSync(join(work,'output/gopurs_runtime'),{recursive:true});mkdirSync(join(work,'record'));
  writeFileSync(join(work,'go.mod'),'module gopurs\n\ngo 1.22\n');
  writeFileSync(join(work,'output/gopurs_runtime/runtime.go'),readFileSync(new URL('../runtime/runtime.go',import.meta.url)));
  writeFileSync(join(work,'record/record.go'),readFileSync(new URL('../../gopurs-argonaut-codecs/src/Data/Argonaut/Decode/Internal/Record.go',import.meta.url)));
  writeFileSync(join(work,'record/record_test.go'),readFileSync(new URL('../../gopurs-argonaut-codecs/test/record-plan_test.go',import.meta.url)));
  writeFileSync(join(work,'record/emitted_test.go'),`package Record
import ("testing"; "gopurs/output/gopurs_runtime")
${declarations}
func TestEmittedDuplicateLabel(t *testing.T) {
 plan:=newErrorPlan(testErrorSupport)
 plan.fields=[]recordDecodeField{{kind:&fieldKind{tag:kindInt}},{kind:&fieldKind{tag:kindMaybe,inner:&fieldKind{tag:kindInt}}}}
 kind:=&fieldKind{tag:kindRecord,plan:plan}
 if !Probe___json_schema_0_decode_accepts(kind) {t.Fatal("schema guard")}
 result:=Probe___json_schema_0_decode(plan,kind,map[string]any{"same":float64(9)})
 if !result.ok || gopurs_runtime.RecordGet(result.value,"same").IntVal!=9 {t.Fatal("duplicate-label precedence")}
}
`);
  const checked=spawnSync('go',['test','-race','-run','^TestEmittedDuplicateLabel$','./record'],{cwd:work,encoding:'utf8',env:{...process.env,GOWORK:'off',GOMAXPROCS:'2'}});
  assert.equal(checked.status,0,(checked.stdout??'')+(checked.stderr??''));
 } finally {rmSync(work,{recursive:true,force:true});}
});
