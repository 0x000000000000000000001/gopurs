import assert from 'node:assert/strict';
import test from 'node:test';
import { Just } from '../output/Data.Maybe/index.js';
import { Tuple } from '../output/Data.Tuple/index.js';
import { ffiProblems } from '../output/Gopurs.FfiBridge/index.js';
import { prepareFfi } from '../output/Gopurs.FfiSupport/index.js';
import * as C from '../output/PureScript.Backend.Optimizer.CoreFn/index.js';

const adt = (name, args = []) => new C.ADT(name, name.split('.'), args);
const fn = (args, result) => new C.Func(args, result);
const str = C.String.value;
const unit = C.Unit.value;
const error = adt('Effect.Exception.Error');
const erasedCallback = value => new C.TypeApp(C.Any.value, [value]);

// Captured from the typed frontend before EffectFnAff was retained nominally.
// Explicit imports can deliberately have this callback shape too.
const callbackAction = adt('Effect.Uncurried.EffectFn2', [
  erasedCallback(error),
  erasedCallback(str),
  adt('Effect.Uncurried.EffectFn3', [error, erasedCallback(error), erasedCallback(unit), unit]),
]);
const expanded = fn([C.Boolean.value], callbackAction);
const nominal = fn([C.Boolean.value], adt('Effect.Aff.Compat.EffectFnAff', [str]));

function problems(source, type = expanded) {
  const parsed = prepareFfi({ moduleName: 'Fixture', path: '/fixture/fixture.go' })('Fixture_')(source)();
  return ffiProblems('Fixture')(parsed.decls)([new Tuple('probe', new Just(type))]);
}

test('incompatible callback-shaped results identify stale typed frontend output', () => {
  for (const result of ['string', '(string, error)', '*Handle', '[]string', 'map[string]string', '']) {
    for (const type of [expanded, new C.ForAll(['a'], expanded)]) {
      const errors = problems(`func Probe(value bool) ${result} { panic("fixture") }`, type);
      assert.equal(errors.length, 1);
      assert.match(errors[0], /probe at 1:6.*Go.*PureScript/);
      assert.match(errors[0], /EffectFn2\/EffectFn3 callback/);
      assert.match(errors[0], /If the source declares EffectFnAff/);
      assert.match(errors[0], /typed purs frontend on PATH \(including node_modules\/\.bin\)/);
      assert.match(errors[0], /clean frontend output, preserving go\.mod\/go\.sum/);
      assert.match(errors[0], /nominal Effect\.Aff\.Compat\.EffectFnAff/);
    }
  }
});

test('the nominal EffectFnAff contract still accepts ordinary Go actions', () => {
  assert.deepEqual(problems('func Probe(value bool) (string, error) { panic("fixture") }', nominal), []);
});

test('intentional native callback implementations retain their explicit ABI', () => {
  const explicit = fn([C.Boolean.value], adt('Effect.Uncurried.EffectFn2', [
    fn([error], unit),
    fn([str], unit),
    adt('Effect.Uncurried.EffectFn3', [error, fn([error], unit), fn([unit], unit), unit]),
  ]));
  const callback = 'func(func(error), func(string)) func(error, func(error), func())';
  for (const [source, type] of [
    [`func Probe(value bool) ${callback} { panic("fixture") }`, explicit],
    [`type Action ${callback}\nfunc Probe(value bool) Action { panic("fixture") }`, explicit],
    [`type Action = ${callback}\nfunc Probe(value bool) Action { panic("fixture") }`, explicit],
    [`func Probe(value bool) ${callback} { panic("fixture") }`, expanded],
  ]) {
    assert.deepEqual(problems(source, type), []);
  }
});

test('legacy runtime implementations remain accepted with callback-shaped metadata', () => {
  for (const result of ['gopurs_runtime.Value', 'any', 'interface{}']) {
    assert.deepEqual(problems(`func Probe(value bool) ${result} { panic("fixture") }`), []);
  }
});

test('other EffectFn2 result contracts do not receive an EffectFnAff diagnostic', () => {
  const other = adt('Effect.Uncurried.EffectFn2', [str, C.Int.value, C.Boolean.value]);
  assert.deepEqual(problems('func Probe(value string, count int64) bool { return true }', other), []);
});
