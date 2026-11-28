import assert from 'node:assert/strict';
import test from 'node:test';
import { spawnSync } from 'node:child_process';
import { mkdirSync, mkdtempSync, readFileSync, writeFileSync, rmSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join } from 'node:path';
import { empty as emptyMap } from '../output/Data.Map/index.js';
import { empty as emptySet } from '../output/Data.Set/index.js';
import { Just, Nothing } from '../output/Data.Maybe/index.js';
import { Tuple } from '../output/Data.Tuple/index.js';
import { prepare } from '../output/Gopurs.ListProgressionFold/index.js';
import { printGoFile } from '../output/Gopurs.Printer/index.js';
import * as C from '../output/PureScript.Backend.Optimizer.CoreFn/index.js';
import { NeutralExpr } from '../output/PureScript.Backend.Optimizer.Semantics/index.js';
import * as S from '../output/PureScript.Backend.Optimizer.Syntax/index.js';

const int = C.Int.value, bool = C.Boolean.value;
const expr = value => new NeutralExpr(value);
const typed = (ty, value) => expr(new S.Typed(ty, value));
const integer = value => typed(int, expr(new S.Lit(new C.LitInt(value))));
const local = (name, level, ty) => typed(ty, expr(new S.Local(name === null ? Nothing.value : new Just(name), level)));
const qualified = (owner, name) => new C.Qualified(new Just(owner), name);
const variable = (owner, name, ty) => typed(ty, expr(new S.Var(qualified(owner, name))));
const apply = (ty, fn, ...args) => typed(ty, expr(new S.App(fn, args)));
const lambda = (ty, args, body) => typed(ty, expr(new S.Abs(args.map(([name, level]) => new Tuple(new Just(name), level)), body)));
const binary = (ty, op, a, b) => typed(ty, expr(new S.PrimOp(new S.Op2(op, a, b))));
const number = (op, a, b) => binary(int, new S.OpIntNum(op), a, b);
const order = (op, a, b) => binary(bool, new S.OpIntOrd(op), a, b);
const branch = (ty, condition, yes, no) => typed(ty, expr(new S.Branch([new S.Pair(condition, yes)], no)));

function fixture(options = {}) {
  const owner = options.owner ?? 'ProgressionProbe', nil = options.nil ?? 'Stop', cons = options.cons ?? 'Cell';
  const filterName = options.filterName ?? 'retain', recurseName = options.recurseName ?? 'exclude';
  const list = new C.ADT(`${owner}.Trail`, [owner, 'Trail'], [int]);
  const generic = new C.ADT(`${owner}.Trail`, [owner, 'Trail'], [new C.TypeVar('a')]);
  const unary = new C.Func([list], list), callback = new C.Func([int], bool);
  const filterType = new C.Func([callback, list], list), traversal = new C.Func([list, list], list);
  const foldType = new C.Func([list, int], int), producerType = new C.Func([int, list], list);
  const ctor = (name, ...fields) => typed(list, expr(new S.CtorSaturated(qualified(owner, name), C.SumType.value, 'Trail', name,
    fields.map((value, index) => new Tuple(`value${index}`, value)))));
  const empty = () => ctor(nil);
  const field = (input, index) => typed(index === 0 ? int : list, expr(new S.Accessor(input,
    new S.GetCtorField(qualified(owner, cons), C.SumType.value, 'Trail', cons, `value${index}`, index))));
  const tag = (name, input) => typed(bool, expr(new S.PrimOp(new S.Op1(new S.OpIsTag(qualified(owner, name)), input))));
  const cases = (ty, input, yes, no) => typed(ty, expr(new S.Branch([
    new S.Pair(tag(nil, input), yes), new S.Pair(tag(cons, input), no),
  ], expr(new S.Fail('incomplete')))));
  const loop = (ty, level, name, value, body, extras = []) => typed(ty, expr(new S.LetRec(level,
    [new Tuple(name, value), ...extras], body)));

  const pred = local('predicate', 0, callback), source = local('source', 1, list);
  const input = local('input', 3, list), acc = local('acc', 4, list), self = local('scan', 2, traversal);
  const revInput = local('r', 6, list), revAcc = local('a', 7, list), revSelf = local('reverse', 5, traversal);
  const reverse = loop(list, 5, 'reverse', lambda(traversal, [['r', 6], ['a', 7]],
    cases(list, revInput, revAcc, apply(list, revSelf, field(revInput, 1), ctor(cons, field(revInput, 0), revAcc)))),
  apply(list, revSelf, acc, empty()));
  const filter = lambda(filterType, [['predicate', 0], ['source', 1]], loop(list, 2, 'scan',
    lambda(traversal, [['input', 3], ['acc', 4]], cases(list, input, options.noReverse ? acc : reverse,
      branch(list, apply(bool, pred, options.filterArgument ? field(input, 1) : field(input, 0)),
        apply(list, self, field(input, 1), ctor(cons, field(input, 0), acc)),
        apply(list, self, field(input, 1), acc)))), apply(list, self, source, empty())));

  const recursiveInput = local('items', 0, list), head = local(null, 1, int), item = local('x', 2, int);
  const modulus = options.capture ? local('other', 99, int) : head;
  const modulo = number(options.moduloOp ?? S.OpMod.value, item, modulus);
  const equal = order(S.OpEq.value, modulo, integer(options.residue ?? 0));
  const predicate = lambda(callback, [['x', 2]], options.opaquePredicate
    ? apply(bool, variable('Foreign', 'test', callback), item)
    : typed(bool, expr(new S.PrimOp(new S.Op1(S.OpBooleanNot.value, equal)))));
  const recursive = lambda(unary, [['items', 0]], cases(list, recursiveInput, empty(),
    typed(list, expr(new S.Let(Nothing.value, 1, field(recursiveInput, 0), ctor(cons,
      options.changedHead ? integer(1) : head,
      apply(list, variable(owner, recurseName, unary), apply(list, variable(owner, filterName, filterType), predicate,
        options.wholeInput ? recursiveInput : field(recursiveInput, 1)))))))));

  const lower = local('lower', 0, int), upper = local('upper', 1, int), seed = local('seed', 2, int);
  const sumSelf = local('sum', 3, foldType), sumInput = local('values', 4, list), sumAcc = local('total', 5, int);
  const folder = lambda(foldType, [['values', 4], ['total', 5]], cases(int, sumInput, sumAcc,
    apply(int, sumSelf, field(sumInput, 1), number(options.sumOp ?? S.OpAdd.value, sumAcc, field(sumInput, 0)))));
  const buildSelf = local('build', 6, producerType), counter = local('current', 7, int), built = local('built', 8, list);
  const bound = options.varyingBound ? counter : options.boundCall ? apply(int, variable('Foreign', 'lower', new C.Func([int], int)), lower) : lower;
  const producer = lambda(producerType, [['current', 7], ['built', 8]],
    branch(list, order(S.OpLt.value, counter, bound), built,
      apply(list, buildSelf, number(S.OpSubtract.value, counter, integer(options.stride ?? 1)), ctor(cons, counter, built))));
  const extras = options.effectfulSibling ? [new Tuple('extra', apply(int, variable('Foreign', 'effect', new C.Func([int], int)), integer(1)))]
    : [new Tuple('unused', producer)];
  const produced = options.borrowed ? local('borrowed', 99, list) : loop(list, 6, 'build', producer,
    apply(list, buildSelf, upper, options.nonemptySeed ? ctor(cons, integer(4), empty()) : empty()), extras);
  const pipeline = apply(int, sumSelf, apply(list, variable(owner, recurseName, unary), produced),
    options.seedCall ? apply(int, variable('Foreign', 'seed', new C.Func([int], int)), seed) : seed);
  const body = options.shadow ? expr(new options.shadow([new Tuple(new Just('sum'), 3)], pipeline)) : pipeline;
  const consume = lambda(new C.Func([int, int, int], int), [['lower', 0], ['upper', 1], ['seed', 2]],
    loop(int, 3, 'sum', folder, options.effectEnvelope ? expr(new S.EffectPure(body)) : body,
      [new Tuple('unusedSum', folder)]));
  const mod = {
    name: owner, bindings: [
      { recursive: false, bindings: [new Tuple(filterName, filter)] },
      { recursive: true, bindings: [new Tuple(recurseName, recursive)] },
      { recursive: false, bindings: [new Tuple('consume', consume)] },
    ],
    dataDecls: [{ name: 'Trail', vars: ['a'], constructors: [
      { name: nil, fields: [] }, { name: cons, fields: options.wrongLayout ? [int, int] : [new C.TypeVar('a'), generic] },
    ] }],
    classDecls: [], dataTypes: emptyMap, foreign: emptyMap, implementations: emptyMap,
    directives: emptyMap, comments: [], imports: emptySet, exports: emptySet, reExports: emptySet,
  };
  if (options.nameCollision) mod.bindings.push({ recursive: false, bindings: [new Tuple('__gopurs_list_progression_fold_0', integer(17))] });
  return mod;
}

function inspect(value, visit) {
  if (!value || typeof value !== 'object') return;
  visit(value);
  for (const child of Object.values(value)) inspect(child, visit);
}

test('recognizes complete renamed recurrences under local scopes and EffectPure', () => {
  for (const options of [{}, { owner: 'Renamed', nil: 'End', cons: 'Link', filterName: 'select', recurseName: 'peel', effectEnvelope: true }]) {
    const source = fixture(options), before = JSON.stringify(source), result = prepare(source);
    assert.equal(JSON.stringify(source), before, 'source module is immutable');
    assert.equal(result.declarations.length, 2);
    assert.deepEqual(result.module.bindings.slice(0, 2), source.bindings.slice(0, 2), 'ordinary functions remain available');
    const fastBranches = [];
    inspect(result.module.bindings[2], value => {
      if (value instanceof S.Branch && JSON.stringify(value.value0).includes('__gopurs_list_progression_fold_')) fastBranches.push(value);
    });
    assert.equal(fastBranches.length, 1);
    const fallback = fastBranches[0].value1;
    const originals = [];
    inspect(source.bindings[2], value => { if (value instanceof S.App) originals.push(value); });
    assert.ok(originals.some(value => JSON.stringify(fallback).includes(JSON.stringify(value))), 'fallback retains the original pipeline');
    const constants = [];
    inspect(fastBranches[0].value0, value => { if (value instanceof C.LitInt) constants.push(value.value0); });
    for (const required of [2, 2147483647, -2147483648, 1048575]) assert.ok(constants.includes(required));
  }
});

for (const [name, options] of [
  ['opaque predicate', { opaquePredicate: true }],
  ['different captured divisor', { capture: true }],
  ['different residue', { residue: 1 }],
  ['division instead of modulo', { moduloOp: S.OpDivide.value }],
  ['changed emitted head', { changedHead: true }],
  ['filtering the complete input again', { wholeInput: true }],
  ['noncommutative folder', { sumOp: S.OpSubtract.value }],
  ['filter without reverse', { noReverse: true }],
  ['predicate receives tail', { filterArgument: true }],
  ['shared producer', { borrowed: true }],
  ['nonempty range seed', { nonemptySeed: true }],
  ['different range stride', { stride: 2 }],
  ['varying range bound', { varyingBound: true }],
  ['opaque bound initializer', { boundCall: true }],
  ['opaque seed initializer', { seedCall: true }],
  ['executed sibling initializer', { effectfulSibling: true }],
  ['unproved list layout', { wrongLayout: true }],
  ['folder masked by a lambda', { shadow: S.Abs }],
  ['folder masked by an uncurried lambda', { shadow: S.UncurriedAbs }],
  ['folder masked by an uncurried effect lambda', { shadow: S.UncurriedEffectAbs }],
]) test(`rejects ${name}`, () => {
  const source = fixture(options), result = prepare(source);
  assert.equal(result.declarations.length, 0);
  assert.deepEqual(result.module, source);
});

test('reserves helper names against existing bindings and emits the guarded bitmap algorithm', () => {
  const result = prepare(fixture({ nameCollision: true }));
  const code = printGoFile({ packageName: 'probe', imports: [], declarationGroups: [result.declarations] });
  assert.match(code, /__gopurs_list_progression_fold_1_general/);
  assert.match(code, /__gopurs_list_progression_fold_1\(/);
  assert.match(code, /make\(\[\]bool, upper-lower\+1\)/);
  assert.match(code, /for value := lower; value <= upper; value\+\+/);
  assert.match(code, /multiple := value \+ value/);
  assert.match(code, /int64\(int32\(/);
  assert.match(code, /lower\) == \(int64\(2\)\)/);
  assert.match(code, /gopurs_runtime\.SumOddProgression\(upper, seed\)/);
  assert.doesNotMatch(code, /pool|Constructor|IntMod/);
});

test('generated prefix and general kernels agree with ordered elimination and retain bounded scratch allocations', () => {
  const code = printGoFile({ packageName: 'probe', imports: [], declarationGroups: [prepare(fixture()).declarations] })
    .replace('import (\n\n)', 'import (\n\tgopurs_runtime "probe/gopurs_runtime"\n)');
  const folder = mkdtempSync(join(tmpdir(), 'gopurs-progression-kernel-'));
  try {
    mkdirSync(join(folder, 'gopurs_runtime'));
    for (const file of ['runtime.go', 'simd_asm.go', 'simd_generic.go', 'simd_arm64.s', 'simd_amd64.s']) {
      writeFileSync(join(folder, 'gopurs_runtime', file), readFileSync(new URL(`../runtime/${file}`, import.meta.url)));
    }
    writeFileSync(join(folder, 'go.mod'), 'module probe\n\ngo 1.24\n');
    writeFileSync(join(folder, 'kernel.go'), code);
    writeFileSync(join(folder, 'kernel_test.go'), `package probe
import("testing";"math/rand")
var sink int64
func oracle(lower,upper,seed int64) int64 {
 heads:=[]int64{};total:=seed
 for value:=lower;value<=upper;value++ {
  keep:=true;for _,head:=range heads {if value%head==0 {keep=false;break}}
  if keep {heads=append(heads,value);total=int64(int32(total+value))}
 }
 return total
}
func TestGenerated(t *testing.T) {
 count:=0
 check:=func(lower,upper,seed int64) {
  got:=Call_ProgressionProbe___gopurs_list_progression_fold_0(lower,upper,seed)
  if want:=oracle(lower,upper,seed);got!=want {t.Fatalf("%d..%d seed%d got%d want%d",lower,upper,seed,got,want)}
  count++
 }
 for _,upper:=range []int64{2,3,4,9,24,25,26,48,49,50,500,1022,1023,1024,1025,2048,20000} {
  for _,seed:=range []int64{-2147483648,-17,0,23,2147483647} {check(2,upper,seed)}
 }
 rng:=rand.New(rand.NewSource(53819))
 for i:=0;i<1000;i++ {
  lower:=int64(2+rng.Intn(100));if i%3==0 {lower=2};if i%7==0 {lower=2147482000+int64(rng.Intn(600))}
  check(lower,lower+int64(rng.Intn(1000)),int64(int32(rng.Uint32())))
 }
 for _,c:=range []struct{lower,upper int64;allocations float64}{{2,2,0},{2,500,0},{2,1023,0},{2,1024,0},{2,1025,1},{4,20,1},{1000,2023,1}} {
  got:=testing.AllocsPerRun(100,func(){sink=Call_ProgressionProbe___gopurs_list_progression_fold_0(c.lower,c.upper,0)})
  // Recent Go versions may also stack-allocate a short dynamic fallback slice.
  if got>c.allocations {t.Fatalf("%+v allocated %g",c,got)}
 }
 t.Logf("%d independent ordered-elimination cases and 7 allocation cases",count)
}
`);
    const result = spawnSync('go', ['test', '-v', './...'], { cwd: folder, encoding: 'utf8' });
    assert.equal(result.status, 0, result.stdout + result.stderr);
  } finally {
    rmSync(folder, { recursive: true, force: true });
  }
});
