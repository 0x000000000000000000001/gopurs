import assert from 'node:assert/strict';
import test from 'node:test';
import { Cons, Nil } from '../output/Data.List.Types/index.js';
import * as Map from '../output/Data.Map/index.js';
import * as MI from '../output/Data.Map.Internal/index.js';
import { Just, Nothing } from '../output/Data.Maybe/index.js';
import { ordString } from '../output/Data.Ord/index.js';
import { Tuple } from '../output/Data.Tuple/index.js';
import * as Set from '../output/Data.Set/index.js';
import { unfoldableArray } from '../output/Data.Unfoldable/index.js';
import * as C from '../output/PureScript.Backend.Optimizer.CoreFn/index.js';
import * as Mono from '../output/PureScript.Backend.Optimizer.Monomorphize/index.js';
import { buildGlobalAstMap } from '../output/Gopurs.Monomorphization/index.js';

const ann = type => ({ span: C.emptySpan, meta: Nothing.value, sourceUsage: Nothing.value, type: new Just(type) });
const variable = new C.TypeVar('a'), integer = C.Int.value;
const poly = new C.ForAll(['a'], new C.Func([variable], integer));
const ref = (owner, name, type) => new C.ExprVar(ann(type), new C.Qualified(owner === null ? Nothing.value : new Just(owner), name));
const literal = value => new C.ExprLit(ann(integer), new C.LitInt(value));
const bind = (name, type, value) => new C.NonRec(new C.Binding(ann(type), name, value));
const moduleOf = (name, decls) => ({ name, path: name + '.purs', span: C.emptySpan, imports: [], comments: [], decls, exports: decls.map(d => d.value0.value1), reExports: Map.empty, dataDecls: [], classDecls: [], foreign: Map.empty });
const list = values => values.reduceRight((tail, value) => new Cons(value, tail), Nil.value);
const entries = map => Map.toUnfoldable(unfoldableArray)(map).map(t => [t.value0, t.value1]);

function fixture(value) {
  const owner = moduleOf('Generic', [bind('ignore', poly, new C.ExprAbs(ann(poly), 'value', literal(42)))]);
  const caller = moduleOf('Caller', [bind('answer', integer, new C.ExprApp(ann(integer), ref('Generic', 'ignore', new C.Func([integer], integer)), literal(value)))]);
  const modules = [owner, caller], ast = buildGlobalAstMap(list(modules));
  const direct = modules.reduce((map, mod) => Mono.collectInstantiations(ast)(map)(mod), Map.empty);
  return { modules, ast, instantiations: Mono.transitiveCollect(ast)(direct) };
}

test('global key lookup plus only own payloads reproduces full monomorphization', () => {
  const { modules, ast, instantiations } = fixture(1);
  const index = MI.functorMap.map(Map.keys)(instantiations);
  assert(entries(index).some(([name]) => name === 'Generic.ignore'));
  for (const mod of modules) {
    const own = Map.filterKeys(ordString)(name => name.startsWith(mod.name + '.'))(instantiations);
    const actual = Mono.monomorphizeWith(ast)(Mono.specializationIndexLookup(index))(own)(mod);
    const full = Mono.monomorphize(ast)(instantiations)(mod);
    assert.deepEqual(actual, full);
    const normalized = MI.functorMap.map(MI.functorMap.map(Mono.normalizeInstantiation))(own);
    assert.deepEqual(Mono.monomorphizeWith(ast)(Mono.specializationIndexLookup(index))(normalized)(mod), full);
  }
  assert.equal(Map.size(Map.filterKeys(ordString)(name => name.startsWith('Caller.'))(instantiations)), 0);
});

test('compact global membership retains missing-key behavior', () => {
  const { modules, ast, instantiations } = fixture(1);
  const index = MI.functorMap.map(Map.keys)(instantiations);
  const withoutOwner = Map.delete(ordString)('Generic.ignore')(index);
  const original = Mono.monomorphizeWith(ast)(Mono.specializationIndexLookup(index))(Map.empty)(modules[1]);
  const incomplete = Mono.monomorphizeWith(ast)(Mono.specializationIndexLookup(withoutOwner))(Map.empty)(modules[1]);
  assert.notDeepEqual(original, incomplete, 'missing global membership changes calls, so the complete index is required');
});


test('persisted dynamic argument projection preserves dictionary eta types and static expressions', () => {
  const fn = new C.Func([integer], integer);
  const dictionary = new C.Record(new C.Row([new Tuple('method', fn)], Nothing.value));
  const staticArg = ref('Helper', 'fun', fn);
  const dynamicDict = new C.ExprLit(ann(dictionary), new C.LitRecord([new C.Prop('method', staticArg)]));
  const constrained = new C.ConstrainedType([new Tuple(['Example', 'Class'], [integer])], fn);
  const body = ref('Definition', 'method', constrained);
  const info = { instType: constrained, dictArgs: [dynamicDict], normalArgs: [staticArg], callers: Set.empty, subst: Map.empty };
  const normalized = Mono.normalizeInstantiation(info);
  assert.notEqual(normalized.dictArgs[0], dynamicDict);
  assert.equal(normalized.normalArgs[0], staticArg, 'static expressions remain exact immutable inputs');
  assert.deepEqual(normalized.dictArgs[0].value0.type, new Just(dictionary));
  assert.deepEqual(Mono.applyStaticArgs(info.dictArgs)([])(body), Mono.applyStaticArgs(normalized.dictArgs)([])(body));
  const untyped = new C.ExprVar({ ...normalized.dictArgs[0].value0, type: Nothing.value }, new C.Qualified(Nothing.value, '__dynamic_argument'));
  assert.notDeepEqual(Mono.applyStaticArgs(info.dictArgs)([])(body), Mono.applyStaticArgs([untyped])([])(body), 'dropping the dynamic dictionary type would alter eta expansion');
});
