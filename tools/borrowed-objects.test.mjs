import assert from 'node:assert/strict';
import test from 'node:test';
import * as Map from '../output/Data.Map/index.js';
import { ordString } from '../output/Data.Ord/index.js';
import { Nothing, Just } from '../output/Data.Maybe/index.js';
import { Tuple } from '../output/Data.Tuple/index.js';
import * as C from '../output/PureScript.Backend.Optimizer.CoreFn/index.js';
import * as S from '../output/PureScript.Backend.Optimizer.Syntax/index.js';
import { NeutralExpr as E } from '../output/PureScript.Backend.Optimizer.Semantics/index.js';
import { borrowReadOnlyObjects } from '../output/Gopurs.BorrowedObjects/index.js';

const q = (m,n) => new C.Qualified(m === null ? Nothing.value : new Just(m), n);
const global = (m,n) => new E(new S.Var(q(m,n)));
const local = n => new E(new S.Local(Nothing.value,n));
const app = (f,...args) => new E(new S.App(f,args));
const letIn = (n,value,body) => new E(new S.Let(Nothing.value,n,value,body));
const field = (ctor, source) => new E(new S.Accessor(source,new S.GetCtorField(q('Data.Either',ctor),C.SumType.value,'Either',ctor,'value0',0)));
const lambda = body => new E(new S.Abs([new Tuple(Nothing.value,9)],body));
const reader = (obj=local(2),m='Data.Argonaut.Decode.Decoders',n='getField__1234') => app(global(m,n),global('Probe','custom'),obj,new E(new S.Lit(new C.LitString('x'))));
const identity = app(global('Data.Argonaut.Decode.Class','decodeForeignObject'),global('Data.Argonaut.Decode.Class','decodeJsonJson'));
const method = new E(new S.Accessor(global('Probe','objectDecoder'),new S.GetProp('decodeJson')));
const read = reader();
const helper = 'Data.Argonaut.Decode.Internal.Record.borrowObject';
const metadata = {globalTypes: Map.insert(ordString)(helper)(new C.Func([C.Any.value,C.Any.value],C.Any.value))(Map.empty)};
function fixture(body, producer=identity, recursive=false) {
  return {name:'Probe',bindings:[{recursive,bindings:[new Tuple('objectDecoder',producer),new Tuple('probe',letIn(1,app(method,local(0)),body))]}]};
}
function borrows(mod, meta=metadata) {
  const before=JSON.stringify(mod), result=borrowReadOnlyObjects(meta)(mod);
  assert.equal(JSON.stringify(mod),before);
  return JSON.stringify(result).includes('borrowObject');
}
test('only synchronous read-only object uses permit borrowing', () => {
  assert.ok(borrows(fixture(letIn(2,field('Right',local(1)),read))));
  assert.ok(borrows(fixture(reader(field('Right',local(1))))));
  assert.ok(borrows(fixture(letIn(2,field('Right',local(1)),letIn(3,local(2),reader(local(3)))))));
  const branches = new E(new S.Branch([
    new S.Pair(new E(new S.PrimOp(new S.Op1(new S.OpIsTag(q('Data.Either','Left')),local(1)))),field('Left',local(1)))
  ],letIn(2,field('Right',local(1)),read)));
  assert.ok(borrows(fixture(branches)));
});
test('returned, captured, stored, mutated and opaque consumers keep owned copies', () => {
  const escapes=[local(2),lambda(read),new E(new S.EffectDefer(read)),
    app(global('Foreign.Object','insert'),local(2)),
    reader(local(2),'Other','getField'),reader(local(2),null,'getField'),
    reader(local(2),'Data.Argonaut.Decode.Decoders','getField__fake'),
    new E(new S.Lit(new C.LitArray([local(2)]))),
    app(global('Probe','unknown'),local(2)),
    app(global('Data.Argonaut.Decode.Decoders','getField'),local(2)),
    new E(new S.LetRec(3,[new Tuple('rec',read)],local(3)))];
  for (const escape of escapes) assert.equal(borrows(fixture(letIn(2,field('Right',local(1)),escape))),false,JSON.stringify(escape));
  assert.equal(borrows(fixture(local(1))),false,'Either envelope escape');
});
test('custom/ambiguous dictionaries, absent helper and recursive groups do not match', () => {
  for(const producer of [global('Probe','unknown'), app(global('Data.Argonaut.Decode.Class','decodeForeignObject'),global('Probe','custom')),app(global(null,'decodeForeignObject'),global('Data.Argonaut.Decode.Class','decodeJsonJson'))]) {
    assert.equal(borrows(fixture(letIn(2,field('Right',local(1)),read),producer)),false);
  }
  const mod=fixture(letIn(2,field('Right',local(1)),read));
  assert.equal(borrows(mod,{globalTypes:Map.empty}),false);
  assert.equal(borrows(fixture(letIn(2,field('Right',local(1)),read),identity,true)),false);
});
