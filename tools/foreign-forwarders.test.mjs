import assert from 'node:assert/strict';
import test from 'node:test';
import { empty as emptyMap, insert } from '../output/Data.Map/index.js';
import { Cons, Nil } from '../output/Data.List.Types/index.js';
import { Just, Nothing } from '../output/Data.Maybe/index.js';
import { ordString } from '../output/Data.Ord/index.js';
import { member } from '../output/Data.Set/index.js';
import * as C from '../output/PureScript.Backend.Optimizer.CoreFn/index.js';
import { collectForeignForwarders } from '../output/Gopurs.Monomorphization/index.js';

const ann = {span: {path: 'Forwarders.purs', start: {line: 1, column: 1}, end: {line: 1, column: 1}},
  meta: Nothing.value, type: Nothing.value, sourceUsage: Nothing.value};
const ref = (name, module = null) => new C.ExprVar(ann, new C.Qualified(module === null ? Nothing.value : new Just(module), name));
const app = (fn, ...args) => args.reduce((f, arg) => new C.ExprApp(ann, f, arg), fn);
const lambda = (params, body) => params.reduceRight((inner, param) => new C.ExprAbs(ann, param, inner), body);
const coerce = (arg, module = 'Unsafe.Coerce') => app(new C.ExprTypeApp(ann, ref('unsafeCoerce', module), C.Any.value), arg);
function recognized(body, {recursive = false, module = 'Forwarders'} = {}) {
  const binding = new C.Binding(ann, 'forward', body);
  const mod = {
    name: module, decls: [recursive ? new C.Rec([binding]) : new C.NonRec(binding)],
    foreign: insert(ordString)('ffi')(Nothing.value)(emptyMap),
  };
  const result = collectForeignForwarders(new Cons(mod, Nil.value));
  return member(ordString)(module + '.forward')(result);
}

test('recognizes exact eta forwarding with qualified or unqualified foreign references', () => {
  for (const module of ['Forwarders', null]) {
    assert.equal(recognized(lambda(['x', 'y'], app(ref('ffi', module), ref('x'), ref('y')))), true);
  }
});

test('recognizes only the qualified Unsafe.Coerce intrinsic through type applications', () => {
  const body = lambda(['x', 'y'], coerce(app(ref('ffi', 'Forwarders'), coerce(ref('x')), coerce(ref('y')))));
  assert.equal(recognized(body), true);
  for (const module of ['Other', null]) {
    assert.equal(recognized(lambda(['x'], app(ref('ffi'), coerce(ref('x'), module)))), false);
    assert.equal(recognized(lambda(['x'], coerce(app(ref('ffi'), ref('x')), module))), false);
  }
});

test('qualification and local shadowing cannot masquerade as foreign calls', () => {
  assert.equal(recognized(lambda(['x'], app(ref('ffi', 'Other'), ref('x')))), false);
  assert.equal(recognized(lambda(['ffi'], app(ref('ffi'), ref('ffi')))), false);
  assert.equal(recognized(lambda(['ffi'], app(ref('ffi', 'Forwarders'), ref('ffi')))), true);
  assert.equal(recognized(lambda(['x', 'x'], app(ref('ffi'), ref('x'), ref('x')))), false);
  assert.equal(recognized(lambda(['unsafeCoerce'], app(ref('ffi'), coerce(ref('unsafeCoerce'), null)))), false);
});

test('captures, reordered, duplicated, computed, missing and extra arguments are rejected', () => {
  for (const args of [
    [ref('y'), ref('x')], [ref('x'), ref('x')], [ref('x'), ref('captured')],
    [ref('x'), ref('y', 'Global')], [ref('x'), app(ref('effect'), ref('y'))],
    [ref('x')], [ref('x'), ref('y'), ref('extra')],
  ]) {
    assert.equal(recognized(lambda(['x', 'y'], app(ref('ffi'), ...args))), false);
  }
  assert.equal(recognized(ref('ffi')), false, 'aliases are not eta wrappers');
});

test('recursive group membership does not hide a real recursive call or extra work', () => {
  assert.equal(recognized(lambda(['x'], app(ref('ffi'), ref('x'))), {recursive: true}), true);
  assert.equal(recognized(lambda(['x'], app(ref('forward', 'Forwarders'), ref('x'))), {recursive: true}), false);
  const recursiveArg = app(ref('forward'), ref('x'));
  assert.equal(recognized(lambda(['x'], app(ref('ffi'), recursiveArg)), {recursive: true}), false);
});
