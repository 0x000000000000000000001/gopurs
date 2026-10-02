import assert from 'node:assert/strict';
import test from 'node:test';
import { empty, insert, lookup } from '../output/Data.Map/index.js';
import { Just, Nothing } from '../output/Data.Maybe/index.js';
import { ordString } from '../output/Data.Ord/index.js';
import { Tuple } from '../output/Data.Tuple/index.js';
import * as C from '../output/PureScript.Backend.Optimizer.CoreFn/index.js';
import { buildGlobalTypes } from '../output/Gopurs.GlobalTypes/index.js';

const ann=type=>({span:C.emptySpan,meta:Nothing.value,sourceUsage:Nothing.value,type});
const none=ann(Nothing.value),int=C.Int.value,str=C.String.value;
const ref=type=>new C.ExprVar(ann(type),new C.Qualified(new Just('External'),'value'));
const binding=(name,type,expr)=>new C.Binding(ann(type),name,expr);
const moduleOf=(name,bindings,foreign=empty)=>({name,decls:bindings,foreign});
const get=(types,name)=>lookup(ordString)(name)(types);

test('binding annotations precede expression annotations, preserving explicit Any at either level',()=>{
 const mod=moduleOf('Types',[
  new C.NonRec(binding('binding',new Just(int),ref(new Just(str)))),
  new C.NonRec(binding('dynamicBinding',new Just(C.Any.value),ref(new Just(int)))),
  new C.NonRec(binding('expression',Nothing.value,ref(new Just(str)))),
  new C.NonRec(binding('dynamicExpression',Nothing.value,ref(new Just(C.Any.value)))),
  new C.NonRec(binding('unknown',Nothing.value,ref(Nothing.value))),
 ]);
 const before=JSON.stringify(mod),types=buildGlobalTypes([mod]);
 assert.equal(JSON.stringify(mod),before);
 for(const [name,type] of [['binding',int],['dynamicBinding',C.Any.value],['expression',str],['dynamicExpression',C.Any.value]]){
  assert.deepEqual(get(types,`Types.${name}`),new Just(type));
 }
 assert.deepEqual(get(types,'Types.unknown'),Nothing.value);
});

test('recursive and dotted module names keep exact source identity; annotated foreign declarations win',()=>{
 const definitions=[new C.Rec([binding('same',new Just(int),ref(Nothing.value)),binding('other',new Just(str),ref(Nothing.value))])];
 const foreign=insert(ordString)('same')(new Just(str))(insert(ordString)('other')(Nothing.value)(empty));
 const types=buildGlobalTypes([moduleOf('A.B',definitions,foreign),moduleOf('A_B',[new C.NonRec(binding('same',new Just(int),ref(Nothing.value)))])]);
 assert.deepEqual(get(types,'A.B.same'),new Just(str));
 assert.deepEqual(get(types,'A.B.other'),new Just(str),'an unannotated FFI does not erase a known definition');
 assert.deepEqual(get(types,'A_B.same'),new Just(int));
});

test('application fallback unwraps quantifiers and constraints and follows nested function results',()=>{
 const fn=new C.ForAll(['a'],new C.ConstrainedType([new Tuple(['Class','C'],[int])],new C.Func([int],new C.Func([str],int))));
 const first=new C.ExprApp(none,ref(new Just(fn)),ref(Nothing.value));
 const second=new C.ExprApp(none,first,ref(Nothing.value));
 const wrapped=new C.ExprTypeApp(none,second,str);
 const types=buildGlobalTypes([moduleOf('Fallback',[new C.NonRec(binding('first',Nothing.value,first)),
  new C.NonRec(binding('second',Nothing.value,second)),new C.NonRec(binding('wrapped',Nothing.value,wrapped))])]);
 assert.deepEqual(get(types,'Fallback.first'),new Just(new C.Func([str],int)));
 assert.deepEqual(get(types,'Fallback.second'),new Just(int));
 assert.deepEqual(get(types,'Fallback.wrapped'),new Just(int));
});

test('application fallback does not infer arbitrary expressions or recover through explicit unknown function types',()=>{
 const blocked=new C.ExprApp(none,ref(new Just(C.Any.value)),ref(new Just(int)));
 const unknown=new C.ExprApp(none,ref(Nothing.value),ref(new Just(int)));
 const access=new C.ExprAccessor(none,ref(new Just(new C.Record(new C.Row([new Tuple('value',int)],Nothing.value)))),'value');
 const typedApp=new C.ExprApp(ann(new Just(C.Any.value)),ref(new Just(new C.Func([int],str))),ref(Nothing.value));
 const types=buildGlobalTypes([moduleOf('Fallback',[...['blocked','unknown','access'].map((name,i)=>new C.NonRec(binding(name,Nothing.value,[blocked,unknown,access][i]))),
  new C.NonRec(binding('explicit',Nothing.value,typedApp))])]);
 for(const name of ['blocked','unknown','access'])assert.deepEqual(get(types,`Fallback.${name}`),Nothing.value);
 assert.deepEqual(get(types,'Fallback.explicit'),new Just(C.Any.value));
});
