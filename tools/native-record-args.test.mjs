import assert from 'node:assert/strict';
import test from 'node:test';
import { empty as emptyMap } from '../output/Data.Map/index.js';
import { empty as emptySet } from '../output/Data.Set/index.js';
import { Just, Nothing } from '../output/Data.Maybe/index.js';
import { Tuple } from '../output/Data.Tuple/index.js';
import * as Go from '../output/Gopurs.GoAst/index.js';
import { exprTypeToGoType } from '../output/Gopurs.GoTypes/index.js';
import { candidateToShare, projectedArgument, workerArguments } from '../output/Gopurs.NativeRecordArgs/index.js';
import * as C from '../output/PureScript.Backend.Optimizer.CoreFn/index.js';
import * as S from '../output/PureScript.Backend.Optimizer.Syntax/index.js';
import { NeutralExpr } from '../output/PureScript.Backend.Optimizer.Semantics/index.js';
import { analyze } from '../output/PureScript.Backend.Optimizer.Codegen.Tco/index.js';

const int = C.Int.value;
const toGoType = exprTypeToGoType(emptyMap)(emptySet)(emptySet)('Probe');
const record = (fields, tail = new Just(new C.TypeVar('r'))) => new C.Record(new C.Row(
  fields.map(([name, type]) => new Tuple(name, type)), tail));
const open = record([['id', int]]);
const projected = new Go.TypeRecord([new Tuple('id', Go.TypeInt64.value)]);
const signature = (args = [open], result = int, vars = ['r']) => new C.ForAll(vars, new C.Func(args, result));
const ann = type => ({span:C.emptySpan, meta:Nothing.value, type, sourceUsage:Nothing.value});
const none = ann(Nothing.value);
const ref = (name = 'row', module = null) => new C.ExprVar(none, new C.Qualified(module === null ? Nothing.value : new Just(module), name));
const read = (value = ref(), field = 'id') => new C.ExprAccessor(none, value, field);
const literal = value => new C.ExprLit(none, new C.LitInt(value));
const lambda = (names, body) => names.reduceRight((inner, name) => new C.ExprAbs(none, name, inner), body);
const app = (fn, ...args) => args.reduce((fn, arg) => new C.ExprApp(none, fn, arg), fn);
const binding = (name, value) => new C.Binding(none, name, value);
const letIn = (bindings, body) => new C.ExprLet(none, bindings, body);
const candidate = (body, type = signature(), names = ['row']) => new C.Binding(ann(new Just(type)), 'reader', lambda(names, body));
function shares(source) {
  const before = JSON.stringify(source);
  const result = candidateToShare(source);
  assert.equal(JSON.stringify(source), before);
  return result;
}

test('projection selects first visible scalar fields, sorts labels and retains their original spelling', () => {
  const fields = [['z', C.Char.value], ['id', int], ['flag', C.Boolean.value], ['amount', C.Number.value], ['text', C.String.value], ['id', C.Any.value]];
  assert.deepEqual(projectedArgument(toGoType)(record(fields)), new Just(new Go.TypeRecord([
    new Tuple('amount', Go.TypeFloat64.value), new Tuple('flag', Go.TypeBool.value),
    new Tuple('id', Go.TypeInt64.value), new Tuple('text', Go.TypeString.value), new Tuple('z', Go.TypeString.value),
  ])));
  const escaped = [['', int], ['0key', int], ['a-b', int], ['type', int]];
  assert.deepEqual(projectedArgument(toGoType)(record(escaped)),
    new Just(new Go.TypeRecord(escaped.map(([label]) => new Tuple(label, Go.TypeInt64.value)))));
  for (const fields of [[['_', int]], [['a b', int]], [['é', int]],
    [['a-b', int], ['a_minus_b', int]], [['type', int], ['go__type', int]],
    [['', int], ['X_empty', int]], [['0key', int], ['X_0key', int]]]) {
    assert.deepEqual(projectedArgument(toGoType)(record(fields)), Nothing.value);
  }
});

test('only nonempty scalar rows with a variable tail qualify for projection', () => {
  for (const type of [record([]), record([['id', C.Any.value], ['id', int]]),
    record([['id', new C.TypeVar('a')]]), record([['id', new C.Array(int)]]),
    record([['id', int]], Nothing.value), record([['id', int]], new Just(C.Any.value)),
    record([['id', int]], new Just(new C.Row([], Nothing.value))), new C.ForAll(['r'], open)]) {
    assert.deepEqual(projectedArgument(toGoType)(type), Nothing.value);
  }
  assert.deepEqual(projectedArgument(toGoType)(open), new Just(projected));
});

test('binding annotations take priority, including explicit Any; only an absent one permits fallback', () => {
  const root = new C.ExprAbs(ann(new Just(signature())), 'row', read());
  assert.equal(shares(new C.Binding(none, 'reader', root)), true);
  assert.equal(shares(new C.Binding(ann(new Just(C.Any.value)), 'reader', root)), false);
  assert.equal(shares(new C.Binding(none, 'reader', lambda(['row'], read()))), false);
  const unknownRoot = new C.ExprAbs(ann(new Just(C.Any.value)), 'row', read());
  assert.equal(shares(new C.Binding(ann(new Just(signature())), 'reader', unknownRoot)), true);
});

test('source signatures permit row-tail quantifiers and monomorphic companion arguments only', () => {
  const other = record([['name', C.String.value]], new Just(new C.TypeVar('s')));
  const nested = new C.ForAll(['r'], new C.ForAll(['s'], new C.Func([open, other], int)));
  assert.equal(shares(candidate(read(), nested, ['row', 'other'])), true);
  for (const companion of [int, new C.Array(int), new C.Func([int], int), record([['extra', int]], Nothing.value)]) {
    assert.equal(shares(candidate(read(), signature([open, companion]), ['row', 'other'])), true);
  }
  for (const companion of [C.Any.value, new C.TypeVar('a'), new C.Array(C.Any.value),
    new C.ForAll(['a'], new C.Func([new C.TypeVar('a')], int)), record([['id', int]], new Just(C.Any.value))]) {
    assert.equal(shares(candidate(read(), signature([open, companion]), ['row', 'other'])), false);
  }
  for (const type of [signature([open], int, ['r', 'unused']), signature([record([['id', int]], Nothing.value)]),
    new C.ConstrainedType([], signature()), new C.TypeApp(signature(), [])]) {
    assert.equal(shares(candidate(read(), type)), false);
  }
  assert.equal(shares(candidate(read(), new C.Func([open], int))), true);
});

test('source sharing requires an exact leading lambda spine with distinct parameter names', () => {
  const pair = signature([open, int]);
  for (const names of [['row'], ['row', 'other', 'extra'], ['row', 'row']]) {
    assert.equal(shares(candidate(read(), pair, names)), false);
  }
  const interrupted = new C.Binding(ann(new Just(pair)), 'reader', lambda(['row'], letIn([], lambda(['other'], read()))));
  assert.equal(shares(interrupted), false);
  const wrapped = new C.Binding(ann(new Just(signature())), 'reader', new C.ExprTypeApp(none, lambda(['row'], read()), int));
  assert.equal(shares(wrapped), false);
  assert.equal(shares(candidate(read(), pair, ['row', 'other'])), true);
});

const nativeResults = [
  new C.ADT('Data.Maybe.Maybe', ['Data', 'Maybe', 'Maybe'], [int]),
  new C.ADT('Data.Either.Either', ['Data', 'Either', 'Either'], [C.String.value, record([['id', int]], Nothing.value)]),
  new C.TypeApp(new C.ADT('Data.Tuple.Tuple', ['Data', 'Tuple', 'Tuple'], []), [int, C.Any.value]),
];
test('source sharing also supports existing native Maybe, Either and Tuple result layouts', () => {
  for (const result of [int, C.Number.value, C.String.value, C.Char.value, C.Boolean.value, ...nativeResults]) {
    assert.equal(shares(candidate(read(), signature([open], result))), true);
  }
  for (const result of [C.Any.value, C.Unit.value, open, new C.Array(int), new C.Func([int], int),
    new C.ADT('Other.Maybe', ['Other', 'Maybe'], [int])]) {
    assert.equal(shares(candidate(read(), signature([open], result))), false);
  }
});

test('source proof admits only direct known-field reads and rejects escapes even through wrappers', () => {
  for (const body of [ref(), read(ref(), 'extra'), app(ref('consume', 'Other'), ref()),
    new C.ExprLit(none, new C.LitArray([ref()])),
    new C.ExprUpdate(none, ref(), [new C.Prop('id', literal(1))]),
    letIn([new C.NonRec(binding('alias', ref()))], read(ref('alias'))),
    read(new C.ExprTypeApp(none, ref(), open)),
    app(ref('delay', 'Other'), lambda(['other'], read()))]) {
    assert.equal(shares(candidate(body)), false);
  }
  assert.equal(shares(candidate(new C.ExprTypeApp(none, read(), int))), true);
  assert.equal(shares(candidate(app(ref('use', 'Other'), read()))), true);
  assert.equal(shares(candidate(app(ref('delay', 'Other'), lambda(['row'], ref())))), true);
  assert.equal(shares(candidate(ref('row', 'Other'))), true, 'qualified globals are not the source parameter');
});

test('source lets check initializers before shadowing and recursive binders shadow the whole group', () => {
  assert.equal(shares(candidate(letIn([new C.NonRec(binding('row', read()))], ref()))), true);
  assert.equal(shares(candidate(letIn([new C.NonRec(binding('row', ref()))], literal(0)))), false);
  assert.equal(shares(candidate(letIn([new C.NonRec(binding('saved', read()))], read()))), true);
  assert.equal(shares(candidate(letIn([new C.Rec([binding('row', ref()), binding('other', ref())])], ref()))), true);
  assert.equal(shares(candidate(letIn([new C.Rec([binding('other', ref())])], literal(0)))), false);
  assert.equal(shares(candidate(letIn([new C.Rec([binding('other', literal(0))])], ref()))), false);
});

test('case shadowing applies to nested patterns and guards, but never to scrutinees or sibling branches', () => {
  const wildcard = new C.BinderNull(none);
  const guarded = (binders, condition, value) => new C.CaseAlternative(binders,
    new C.Guarded([new C.Guard(condition, value)]));
  const caseOf = (value, branches) => new C.ExprCase(none, [value], branches);
  const patterns = [new C.BinderNamed(none, 'row', wildcard),
    new C.BinderLit(none, new C.LitRecord([new C.Prop('nested', new C.BinderVar(none, 'row'))])),
    new C.BinderConstructor(none, new C.Qualified(new Just('Box'), 'Box'),
      new C.Qualified(new Just('Box'), 'Box'), [new C.BinderVar(none, 'row')])];
  for (const pattern of patterns) {
    const shadowed = guarded([pattern], ref(), ref());
    assert.equal(shares(candidate(caseOf(literal(1), [shadowed]))), true);
    assert.equal(shares(candidate(caseOf(ref(), [shadowed]))), false);
    assert.equal(shares(candidate(caseOf(literal(1), [shadowed, guarded([wildcard], literal(1), ref())]))), false);
  }
  assert.equal(shares(candidate(caseOf(read(), [guarded([wildcard], read(), read())]))), true);
  for (const branch of [guarded([wildcard], ref(), read()), guarded([wildcard], read(), ref())]) {
    assert.equal(shares(candidate(caseOf(literal(1), [branch]))), false);
  }
});

const ir = syntax => new NeutralExpr(syntax);
const local = (name = 'row', level = 0) => ir(new S.Local(name === null ? Nothing.value : new Just(name), level));
const typed = (type, value) => ir(new S.Typed(type, value));
const get = (value = local(), field = 'id') => ir(new S.Accessor(value, new S.GetProp(field)));
const worker = (body, {names = ['row_0'], types = [open], result = int} = {}) =>
  workerArguments(toGoType)(names)(analyze([])(body))(types)(result);

test('final TCO proof independently rejects stale source eligibility and selects each argument separately', () => {
  assert.equal(shares(candidate(read())), true);
  assert.deepEqual(worker(get(typed(open, typed(open, local())))), [projected]);
  for (const body of [local(), get(local(), 'extra'), ir(new S.App(ir(new S.Var(new C.Qualified(new Just('Other'), 'consume'))), [local()])),
    ir(new S.Let(new Just('alias'), 1, local(), get(local('alias', 1)))),
    get(ir(new S.TypeApp(local(), open)))]) {
    assert.deepEqual(worker(body), [Go.TypeValue.value]);
  }
  const uses = ir(new S.Lit(new C.LitArray([get(), local('other', 1)])));
  assert.deepEqual(worker(uses, {names:['row_0', 'other_1', 'count_2'], types:[open, open, int]}),
    [projected, Go.TypeValue.value, Go.TypeInt64.value]);
  assert.deepEqual(worker(get(), {names:[], types:[open, int]}), [Go.TypeValue.value, Go.TypeInt64.value]);
});

test('TCO captures use emitted local identities and reject deferred reads even under a rebound identifier', () => {
  for (const scope of [
    value => ir(new S.Abs([new Tuple(new Just('row'), 0)], value)),
    value => ir(new S.UncurriedAbs([new Tuple(new Just('row'), 0)], value)),
    value => ir(new S.UncurriedEffectAbs([new Tuple(new Just('row'), 0)], value)),
    value => ir(new S.EffectDefer(value)),
  ]) {
    assert.deepEqual(worker(scope(get())), [Go.TypeValue.value]);
    assert.deepEqual(worker(scope(get(local('row', 1)))), [projected]);
  }
  assert.deepEqual(worker(get(local("row'", 2)), {names:['row_prime__2']}), [projected]);
  assert.deepEqual(worker(local("row'", 2), {names:['row_prime__2']}), [Go.TypeValue.value]);
  assert.deepEqual(worker(get(local(null, 3)), {names:['__local_var_3']}), [projected]);
  assert.deepEqual(worker(local(null, 3), {names:['__local_var_3']}), [Go.TypeValue.value]);
});

test('final result policy permits native ADTs and preserves the ordinary ABI for unsupported results and rows', () => {
  for (const result of nativeResults) assert.deepEqual(worker(get(), {result}), [projected]);
  for (const result of [open, C.Any.value, new C.Func([int], int)]) {
    assert.deepEqual(worker(get(), {result}), [Go.TypeValue.value]);
  }
  const closed = record([['id', int]], Nothing.value);
  assert.deepEqual(worker(get(), {types:[closed], result:open}), [toGoType(closed)]);
  assert.deepEqual(worker(get(), {types:[record([['id', C.Any.value]])]}), [Go.TypeValue.value]);
});
