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
import * as T from '../output/Gopurs.DecoderSchemas.Types/index.js';
import { supported } from '../output/Gopurs.DecoderSchemas.Stream/index.js';
import { specializeDecoderSchemas } from '../output/Gopurs.DecoderSchemas/index.js';
import * as G from '../output/Gopurs.GoAst/index.js';
import { collectImports } from '../output/Gopurs.GoImports/index.js';
import { instantiateGenericGoType } from '../output/Gopurs.GoTypes/index.js';
import { printGoDecl } from '../output/Gopurs.Printer/index.js';
import { translate } from '../output/Gopurs.CodeGen/index.js';
import { empty as emptySet } from '../output/Data.Set/index.js';
import { withGlobalTypes, withReboxFields } from './codegen-metadata.mjs';

const global=(m,n)=>new E(new S.Var(new C.Qualified(m===null?Nothing.value:new Just(m),n)));
const cls=n=>global('Data.Argonaut.Decode.Class',n);
const app=(fn,...args)=>new E(new S.App(fn,args));
const erased=new E(S.PrimUndefined.value);
const literal=s=>new E(new S.Lit(new C.LitString(s)));
const local=i=>new E(new S.Local(Nothing.value,i));
const symbol=s=>new E(new S.Lit(new C.LitRecord([new C.Prop('reflectSymbol',new E(new S.Abs([new Tuple(Nothing.value,0)],literal(s))))])));
const int=cls('decodeJsonInt'),str=cls('decodeJsonString'),num=cls('decodeJsonNumber');
const field=(key,decoder,tail=cls('gDecodeJsonNil'))=>app(cls('gDecodeJsonCons'),app(cls('decodeFieldId'),decoder),tail,symbol(key),erased,erased);
const record=row=>app(cls('decodeRecord'),row,erased);
const method=dict=>app(cls('decodeJson'),dict);
const metadata=(text=true,stream=true,streamType=C.Int.value)=>{
 let types=Map.insert(ordString)('Data.Argonaut.Decode.Internal.Record.schemaDecoderABI1')(C.Int.value)(Map.empty);
 if(text)types=Map.insert(ordString)('Data.Argonaut.Decode.Parser.textDecoderABI1')(C.Int.value)(types);
 if(stream)types=Map.insert(ordString)('Data.Argonaut.Decode.Parser.streamDecoderABI1')(streamType)(types);
 return {globalTypes:types};
};
function run(dict,extra=[],meta=metadata()) {
 const mod={name:'Probe',bindings:[{recursive:false,bindings:[...extra,new Tuple('decode',method(dict))]}]};
 const snapshot=JSON.stringify(mod);const result=specializeDecoderSchemas(withGlobalTypes(meta))(mod);
 assert.equal(JSON.stringify(mod),snapshot);
 return {...result,code:result.declarations.map(printGoDecl).join('\n')};
}
const streamFunctions=result=>result.declarations.map(printGoDecl).filter(x=>/^func \w+_stream\(/.test(x)).join('\n');

test('multiple Go results retain their structured types, imports and generic instantiation',()=>{
 const result=new G.TypeResults([G.TypeValue.value,G.TypeBool.value]);
 assert.equal(G.goTypeToStr(result),'(gopurs_runtime.Value, bool)');
 assert.equal(G.goTypeIdentity(result),'(gopurs_runtime.Value, bool)');
 const declaration=new G.GoFunctionDecl({name:{owner:'generated/Probe',name:'pair'},params:[],result,body:G.rawGo('panic("unused")')});
 assert.deepEqual(collectImports([[declaration]]),['gopurs/output/gopurs_runtime']);
 assert.match(printGoDecl(declaration),/func pair\(\) \(gopurs_runtime.Value, bool\)/);
 const generic=new G.TypeFunc([],new G.TypeResults([new G.TypeGenericParam('a'),G.TypeBool.value]));
 const env=Map.insert(ordString)('a')(G.TypeString.value)(Map.empty);
 assert.equal(G.goTypeToStr(instantiateGenericGoType(env)(generic)),'func() (string, bool)');
});

test('stream ABI is independent and requires the existing text ABI and typed markers',()=>{
 const schema=record(field('count',int));
 for(const [text,stream] of [[false,false],[false,true],[true,false]]) {
  const result=run(schema,[],metadata(text,stream));
  assert.doesNotMatch(result.code,/argonautCompileStreamSchema|argonautStreamParser/);
  assert.equal(result.code.includes('argonautCompileTextSchema'),text);
 }
 assert.doesNotMatch(run(schema,[],metadata(true,true,C.Any.value)).code,/argonautCompileStreamSchema/);
 const result=run(schema);
 assert.match(result.code,/argonautCompileStreamSchema\(argonautCompileTextSchema\(/);
 assert.equal((result.code.match(/Get_Probe___json_schema_0_source\(\)/g)??[]).length,1,'source dictionary is evaluated once');
 assert.match(result.code,/raw\.document\.tokens\[at\]\.next/,'text ABI1 traversal stays unchanged');
});

test('streaming preserves typed applications and module TAST declarations',()=>{
 const source=new E(new S.Typed(C.Any.value,new E(new S.TypeApp(method(record(field('value',int))),C.Any.value))));
 const mod={name:'Probe',bindings:[{recursive:false,bindings:[new Tuple('decode',source)]}],
  dataDecls:[{name:'Retained',vars:[],constructors:[]}],classDecls:[{name:'RetainedClass',members:[]}]};
 const result=specializeDecoderSchemas(withGlobalTypes(metadata()))(mod);
 assert.equal(result.module.dataDecls,mod.dataDecls);assert.equal(result.module.classDecls,mod.classDecls);
 const value=result.module.bindings[0].bindings[0].value1;
 assert.ok(value instanceof S.Typed);assert.ok(value.value1 instanceof S.TypeApp);
 assert.match(result.declarations.map(printGoDecl).join('\n'),/argonautCompileStreamSchema/);
});

test('stream admission refuses opaque leaves, Json, duplicate schemas and opaque successful computations',()=>{
 for(const schema of [T.Custom.value,new T.Scalar('Json'),new T.Scalar('unknown'),new T.Sequence(T.Custom.value),
  new T.ObjectSchema([new Tuple('x',new T.Scalar('Int')),new Tuple('x',new T.Scalar('Int'))]),
  new T.Derived(new T.ReturnValue(app(global('Opaque','callback'),literal('x'))))]) assert.equal(supported(schema),false);
 for(const dict of [record(field('value',cls('decodeJsonJson'))),record(field('value',global('Opaque','decoder'))),record(field('same',int,field('same',str)))])
  assert.doesNotMatch(run(dict).code,/argonautCompileStreamSchema/);
 assert.equal(supported(new T.Derived(new T.TypeMismatch('rejected'))),true);
});

test('derived slots require structurally identical schemas and distinguish nullable reads',()=>{
 const ret=new T.ReturnValue(local(2));
 const read=(schema,next=ret,nullable=false)=>new T.ReadField(2,'shared',true,nullable,schema,next);
 const choose=(a,b)=>new T.Derived(new T.Choices([new Tuple(new Tuple(1,'left'),a)],b));
 assert.equal(supported(choose(read(new T.Scalar('Int')),read(new T.Scalar('String')))),false);
 assert.equal(supported(choose(read(new T.Sequence(new T.Scalar('Int'))),read(new T.Sequence(new T.Scalar('Number'))))),false);
 assert.equal(supported(choose(read(new T.ObjectSchema([new Tuple('nested',new T.Scalar('Int'))])),read(new T.ObjectSchema([new Tuple('nested',new T.Scalar('Int'))])))),true);
 assert.equal(supported(choose(read(new T.Scalar('Int')),read(new T.Scalar('Int'),ret,true))),true);
});

test('repeated keys are rejected along one path but remain admitted in exclusive branches',()=>{
 const schema=new T.Sequence(new T.Scalar('Int'));
 const ret=new T.ReturnValue(local(2));
 const read=(next=ret)=>new T.ReadField(2,'shared',false,false,schema,next);
 const branches=new T.Choices([new Tuple(new Tuple(1,'left'),read())],read());
 assert.equal(supported(new T.Derived(read(read()))),false);
 assert.equal(supported(new T.Derived(branches)),true);
 assert.equal(supported(new T.Derived(read(branches))),false);
});

test('generated streaming workers use owned scalar paths, conservative object slots and success-only depth release',()=>{
 const result=run(record(field('seed',int,field('note',app(cls('decodeJsonMaybe'),str),field('data',app(cls('decodeArray'),num))))));
 const code=streamFunctions(result);
 assert.match(code,/p\.string\(true\)/);
 assert.match(code,/if seen\[0\]/);
 assert.match(code,/default: return gopurs_runtime.Value\{\}, false/);
 assert.match(code,/if !seen\[1\] \{ value1 = plan\.mkNothing\(\)/);
 assert.match(code,/if len\(values\) == 0 \{ values = make\(\[\]gopurs_runtime.Value, 0, 4\)/);
 assert.match(code,/p\.leave\(\); return gopurs_runtime.Array\(values\), true/);
 assert.doesNotMatch(code,/defer|mkTypeMismatch|mkAtKey|mkAtIndex|mkNamed|materialize|argonautSchemaCustom/);
});

const either=n=>new C.Qualified(new Just('Data.Either'),n);
const ctor=(name,...fields)=>new E(new S.CtorSaturated(either(name),C.SumType.value,'Either',name,fields.map((v,i)=>new Tuple('value'+i,v))));
const access=(name,i)=>new E(new S.Accessor(local(i),new S.GetCtorField(either(name),C.SumType.value,'Either',name,'value0',0)));
const is=(name,i)=>new E(new S.PrimOp(new S.Op1(new S.OpIsTag(either(name)),local(i))));
const bind=(i,value,next)=>new E(new S.Let(Nothing.value,i,value,new E(new S.Branch([
 new S.Pair(is('Left',i),ctor('Left',access('Left',i))),new S.Pair(is('Right',i),next)
],new E(new S.Fail('Failed pattern match'))))));
const dictionary=body=>new E(new S.Lit(new C.LitRecord([new C.Prop('decodeJson',new E(new S.Abs([new Tuple(Nothing.value,0)],body)))])));
const borrow=app(global('Data.Argonaut.Decode.Internal.Record','borrowObject'),new E(new S.Accessor(app(cls('decodeForeignObject'),cls('decodeJsonJson')),new S.GetProp('decodeJson'))),local(0));
const getField=(key,decoder=str,reader='getField')=>app(global('Data.Argonaut.Decode.Decoders',reader),method(decoder),access('Right',1),literal(key));
const equal=(left,right)=>new E(new S.PrimOp(new S.Op2(new S.OpStringOrd(S.OpEq.value),left,right)));
const mismatch=message=>ctor('Left',new E(new S.CtorSaturated(new C.Qualified(new Just('Data.Argonaut.Decode.Error'),'TypeMismatch'),C.SumType.value,'JsonDecodeError','TypeMismatch',[new Tuple('value0',literal(message))])));
function derivedResult(conflict=false) {
 const read=(decoder,reader='getField')=>bind(4,getField('payload',decoder,reader),ctor('Right',access('Right',4)));
 const body=bind(9,getField('kind'),new E(new S.Branch([
  new S.Pair(equal(access('Right',9),literal('strict')),read(int)),
  new S.Pair(equal(access('Right',9),literal('strict')),ctor('Right',literal('unreachable'))),
  new S.Pair(equal(access('Right',9),literal('optional')),read(int,'getFieldOptional')),
  new S.Pair(equal(access('Right',9),literal('nullable')),read(conflict?str:int,"getFieldOptional'")),
 ],mismatch('unknown kind'))));
 return run(app(cls('decodeArray'),global('Probe','custom')),[new Tuple('custom',dictionary(bind(1,borrow,body)))]);
}
test('derived emission shares original constructors and keeps choices, nullable flags and conflicts conservative',()=>{
 const result=derivedResult();const code=streamFunctions(result);
 assert.match(code,/var seen, nulls \[2\]bool/);
 assert.match(code,/!seen\[1\] \|\| nulls\[1\]/);
 assert.match(code,/if nulls\[1\] \{ return gopurs_runtime.Value\{\}, false/);
 assert.equal((code.match(/case "payload"/g)??[]).length,1);
 assert.doesNotMatch(code,/mkTypeMismatch|mkAtKey|argonautSchemaResult/);
 assert.ok(result.module.bindings.at(-1).bindings.some(p=>p.value0.endsWith('_construct')));
 assert.doesNotMatch(derivedResult(true).code,/argonautCompileStreamSchema/);
});

function constructors(result) {
 const compilerMetadata=withReboxFields({elidedCtors:emptySet,ctorTypes:Map.empty,pointerAdtPaths:Map.empty,pointerAdtNodes:emptySet,
  pointerAdtLeaves:Map.empty,enumAdts:emptySet,enumCtors:emptySet,globalFunctions:Map.empty,globalTypes:Map.empty,classDeclsFields:Map.empty});
 return translate(compilerMetadata)({name:'Probe',bindings:[{recursive:false,bindings:result.module.bindings.at(-1).bindings.filter(p=>p.value0.endsWith('_construct'))}],
  comments:[],imports:emptySet,exports:emptySet,reExports:emptySet,dataTypes:Map.empty,dataDecls:[],classDecls:[],foreign:Map.empty,implementations:Map.empty,directives:Map.empty});
}
function checkEmitted(result,testCode,constructorCode='',includePortTests=false) {
 const work=mkdtempSync(join(tmpdir(),'gopurs-stream-emission-'));
 try {
  mkdirSync(join(work,'output/gopurs_runtime'),{recursive:true});mkdirSync(join(work,'record'));
  writeFileSync(join(work,'go.mod'),'module gopurs\n\ngo 1.22\n');
  for (const rtFile of ['runtime.go', 'simd_asm.go', 'simd_generic.go', 'simd_arm64.s', 'simd_amd64.s'])
   writeFileSync(join(work, `output/gopurs_runtime/${rtFile}`), readFileSync(new URL(`../runtime/${rtFile}`, import.meta.url)));
  for(const [name,path] of [['record','../../gopurs-argonaut-codecs/src/Data/Argonaut/Decode/Internal/Record.go'],
   ['parser','../../gopurs-argonaut-core/src/Data/Argonaut/Parser.go'],['text','../../gopurs-argonaut-codecs/src/Data/Argonaut/Decode/Parser.go']])
   writeFileSync(join(work,`record/${name}.go`),readFileSync(new URL(path,import.meta.url),'utf8').replace('package Parser','package Record'));
  writeFileSync(join(work,'record/record_test.go'),readFileSync(new URL('../../gopurs-argonaut-codecs/test/record-plan_test.go',import.meta.url)));
  if(includePortTests)writeFileSync(join(work,'record/stream_owned_test.go'),readFileSync(new URL('../../gopurs-argonaut-codecs/test/stream-owned_test.go',import.meta.url)));
  writeFileSync(join(work,'record/workers.go'),'package Record\nimport "gopurs/output/gopurs_runtime"\n'+result.declarations.map(printGoDecl).filter(x=>x.startsWith('func ')).join('\n'));
  if(constructorCode)writeFileSync(join(work,'record/constructors.go'),constructorCode.replace('package purescript','package Record'));
  writeFileSync(join(work,'record/emitted_test.go'),testCode);
  writeFileSync(join(work,'record/comparison_test.go'),`package Record
import "gopurs/output/gopurs_runtime"
func streamComparable(v gopurs_runtime.Value) any {
 switch v.Type {case gopurs_runtime.TypeString:return v.StrVal();case gopurs_runtime.TypeInt:return v.IntVal;case gopurs_runtime.TypeFloat:return gopurs_runtime.Unbox[float64](v);case gopurs_runtime.TypeBool:return v.BoolVal();case gopurs_runtime.TypeArray:
 out:=make([]any,gopurs_runtime.ArrayLength(v));for i:=range out{out[i]=streamComparable(gopurs_runtime.ArrayAccess(v,i))};return out}
 out:=map[string]any{};for k,x:=range gopurs_runtime.RecordToMap(v){out[k]=streamComparable(x)};return out
}
`);
  const checked=spawnSync('go',['test','-race','-count=1','-run','^Test(EmittedStream|Stream)','./record'],{cwd:work,encoding:'utf8',timeout:90000,env:{...process.env,GOWORK:'off',GOMAXPROCS:'2'}});
  assert.ifError(checked.error);assert.equal(checked.status,0,checked.stdout+checked.stderr);
 } finally {rmSync(work,{recursive:true,force:true});}
}

test('generated workers preserve exact text fallback, source ownership and cold ABI compatibility in Go',()=>{
 const result=run(record(field('seed',int,field('note',app(cls('decodeJsonMaybe'),str),field('data',app(cls('decodeArray'),num))))));
 checkEmitted(result,`package Record
import("testing";"reflect";"unsafe";"gopurs/output/gopurs_runtime")
func TestEmittedStreamRecord(t *testing.T){
 plan:=newErrorPlan(testErrorSupport);plan.fields=[]recordDecodeField{{kind:&fieldKind{tag:kindInt}},{kind:&fieldKind{tag:kindMaybe,inner:&fieldKind{tag:kindString}}},{kind:&fieldKind{tag:kindArray,inner:&fieldKind{tag:kindNumber}}}}
 kind:=&fieldKind{tag:kindRecord,plan:plan};original:=gopurs_runtime.WithFunctionData(gopurs_runtime.Func(func(_ gopurs_runtime.Value)gopurs_runtime.Value{panic("opaque original")}),&typedKind{kind:kind,errors:plan,worker:Probe___json_schema_0_decode})
 plain:=argonautCompileTextSchema(original,Probe___json_schema_0_decode_text,Probe___json_schema_0_decode_text_accepts);compiled:=argonautCompileStreamSchema(plain,Probe___json_schema_0_decode_stream)
 fallback:=gopurs_runtime.Func(func(text gopurs_runtime.Value)gopurs_runtime.Value{return gopurs_runtime.Str("fallback:"+text.StrVal())})
 cases:=[]struct{text string;accepted bool}{
 {\`{"seed":7,"note":"owned","data":[1,2.25,-0]}\`,true},{\`{"data":[],"seed":0}\`,true},{\`{"note":null,"data":[3],"seed":1}\`,true},
 {\`{"seed":0,"seed":7,"data":[]}\`,false},{\`{"seed":7,"data":[],"ignored":true}\`,false},{\`{"seed":false,"data":[]}\`,false},{\`{"seed":7}\`,false},
 {\`{"seed":7,"data":[1e999]}\`,false},{\`{"seed":7,"data":[]} trailing\`,false},{\`{"seed":7,"data":[],}\`,false},
 }
 for _,tc:=range cases {_,ok:=argonautTryStreaming(Probe___json_schema_0_decode_stream,plan,kind,tc.text);if ok!=tc.accepted{t.Fatal("admission",tc.text)}
 a,b:=DecodeJsonStringImpl(1,fallback,plain,tc.text),DecodeJsonStringImpl(1,fallback,compiled,tc.text);if !reflect.DeepEqual(streamComparable(a),streamComparable(b)){t.Fatal("fallback changed",tc.text)}}
 data:=[]byte(\`{"seed":7,"note":"own\\ncopy-é","data":[]}\`);text:=unsafe.String(unsafe.SliceData(data),len(data))
 a,ok:=argonautTryStreaming(Probe___json_schema_0_decode_stream,plan,kind,text);if !ok{t.Fatal("owned")};b,ok:=argonautTryStreaming(Probe___json_schema_0_decode_stream,plan,kind,text);if !ok{t.Fatal("owned twice")}
 get:=func(v gopurs_runtime.Value)string{return gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v,"note"),"just").StrVal()};x,y:=get(a),get(b);if unsafe.StringData(x)==unsafe.StringData(y){t.Fatal("cross-call sharing")};for i:=range data{data[i]='x'};if get(a)!="own\\ncopy-é"||get(b)!="own\\ncopy-é"{t.Fatal("source alias")}
}
`,'',true);
});

test('generated derived workers preserve ordered choices and separate absent, null and decoded values',()=>{
 const result=derivedResult();checkEmitted(result,`package Record
import("testing";"reflect";"gopurs/output/gopurs_runtime")
func TestEmittedStreamDerived(t *testing.T){
 plan:=newErrorPlan(testErrorSupport);kind:=&fieldKind{tag:kindArray}
 cases:=[]struct{text string;accepted bool}{
 {\`[{"kind":"strict","payload":7}]\`,true},{\`[{"payload":8,"kind":"strict"}]\`,true},{\`[{"kind":"optional"},{"kind":"nullable","payload":null}]\`,true},
 {\`[{"kind":"optional","payload":9}]\`,true},{\`[{"kind":"optional","payload":null}]\`,false},{\`[{"kind":"strict","payload":null}]\`,false},
 {\`[{"kind":"strict"}]\`,false},{\`[{"kind":"other","payload":7}]\`,false},{\`[{"kind":"strict","payload":1,"payload":9}]\`,false},{\`[{"kind":"strict","payload":7,"unknown":0}]\`,false},
 }
 for _,tc:=range cases {fast,ok:=argonautTryStreaming(Probe___json_schema_0_decode_stream,plan,kind,tc.text);if ok!=tc.accepted{t.Fatal("admission",tc.text)};cursor,valid:=argonautTextIndex(tc.text);if !valid{t.Fatal("index")};base:=Probe___json_schema_0_decode_text(plan,kind,cursor);if ok&&(!base.ok||!reflect.DeepEqual(streamComparable(fast),streamComparable(base.value))){t.Fatal("value",tc.text)}}
 // A failed speculative read must not call any error constructor.
 plan.atKey=gopurs_runtime.Func2(func(_, _ gopurs_runtime.Value)gopurs_runtime.Value{panic("speculative error callback")});plan.typeMismatch=gopurs_runtime.Func(func(_ gopurs_runtime.Value)gopurs_runtime.Value{panic("speculative error callback")})
 for _,text:=range []string{\`[{}]\`,\`[{"kind":"other"}]\`,\`[{"kind":"strict","payload":false}]\`}{if _,ok:=argonautTryStreaming(Probe___json_schema_0_decode_stream,plan,kind,text);ok{t.Fatal("unexpected admission")}}
}
`,constructors(result));
});

test('repeated derived container reads preserve distinct result storage',()=>{
 const pair=new E(new S.CtorSaturated(new C.Qualified(new Just('Probe'),'Pair'),C.ProductType.value,'Pair','Pair',[
  new Tuple('value0',access('Right',2)),new Tuple('value1',access('Right',3)),
 ]));
 const arrayInt=app(cls('decodeArray'),int);
 const body=bind(1,borrow,bind(2,getField('shared',arrayInt),bind(3,getField('shared',arrayInt),ctor('Right',pair))));
 const result=run(app(cls('decodeArray'),global('Probe','custom')),[new Tuple('custom',dictionary(body))]);
 const construct=result.module.bindings.at(-1).bindings.find(p=>p.value0.endsWith('_construct')).value0;
 const worker=`Probe_${result.module.bindings[0].bindings.at(-1).value1.value0.value1}_decode`;
 assert.doesNotMatch(streamFunctions(result),new RegExp(`func ${worker}_stream\\(`));
 checkEmitted(result,`package Record
import("testing";"unsafe";"gopurs/output/gopurs_runtime")
type streamPair struct{first,second gopurs_runtime.Value}
// A pure two-field ADT constructor: it stores its arguments independently and
// never copies or shares their array backing stores itself.
func Get_Probe_${construct}()gopurs_runtime.Value{return gopurs_runtime.Func2(func(a,b gopurs_runtime.Value)gopurs_runtime.Value{return gopurs_runtime.Value{Type:gopurs_runtime.TypeConstructor,UnsafePtr:unsafe.Pointer(&streamPair{a,b})}})}
func TestEmittedStreamRepeatedContainers(t *testing.T){
 plan:=newErrorPlan(testErrorSupport);kind:=&fieldKind{tag:kindArray};text:=\`[{"shared":[1,2]}]\`
 cursor,ok:=argonautTextIndex(text);if !ok{t.Fatal("index")}
 original:=Probe___json_schema_0_decode_text(plan,kind,cursor);if !original.ok{t.Fatal("baseline")}
 decoder:=gopurs_runtime.WithFunctionData(gopurs_runtime.Func(func(_ gopurs_runtime.Value)gopurs_runtime.Value{panic("unexpected opaque decoder")}),&typedKind{kind:kind,errors:plan,worker:Probe___json_schema_0_decode})
 decoder=argonautCompileTextSchema(decoder,Probe___json_schema_0_decode_text,Probe___json_schema_0_decode_text_accepts)
 fallback:=gopurs_runtime.Func(func(_ gopurs_runtime.Value)gopurs_runtime.Value{panic("valid text must retain text worker")})
 decoded:=DecodeJsonStringImpl(1,fallback,decoder,text);if !gopurs_runtime.RecordGet(decoded,"ok").BoolVal(){t.Fatal("text fallback")};streamed:=gopurs_runtime.RecordGet(decoded,"value")
 oldPair:=(*streamPair)(gopurs_runtime.ArrayAccess(original.value,0).UnsafePtr);newPair:=(*streamPair)(gopurs_runtime.ArrayAccess(streamed,0).UnsafePtr)
 oldLeft,oldRight:=gopurs_runtime.Unbox[[]gopurs_runtime.Value](oldPair.first),gopurs_runtime.Unbox[[]gopurs_runtime.Value](oldPair.second)
 newLeft,newRight:=gopurs_runtime.Unbox[[]gopurs_runtime.Value](newPair.first),gopurs_runtime.Unbox[[]gopurs_runtime.Value](newPair.second)
 oldAlias,newAlias:=unsafe.SliceData(oldLeft)==unsafe.SliceData(oldRight),unsafe.SliceData(newLeft)==unsafe.SliceData(newRight)
 oldLeft[0]=gopurs_runtime.Int(99);newLeft[0]=gopurs_runtime.Int(99)
 t.Logf("baseline alias=%t other[0]=%d; retained text path alias=%t other[0]=%d",oldAlias,oldRight[0].IntVal,newAlias,newRight[0].IntVal)
 if oldAlias||oldRight[0].IntVal!=1{t.Fatal("invalid baseline oracle")}
 if newAlias||newRight[0].IntVal!=1{t.Fatal("fallback introduced an observable alias between two independent decoded arrays")}
}
`.replaceAll('Probe___json_schema_0_decode',worker));
});
