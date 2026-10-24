import assert from 'node:assert/strict';
import test from 'node:test';
import * as Table from '../output/Gopurs.Cache.LookupTable/index.js';
import * as Wire from '../output/Gopurs.Cache.Wire/index.js';
import * as Lazy from '../output/Data.Lazy/index.js';
import * as C from '../output/PureScript.Backend.Optimizer.CoreFn/index.js';
import * as Map from '../output/Data.Map/index.js';
import * as Set from '../output/Data.Set/index.js';
import { ordString } from '../output/Data.Ord/index.js';
import { foldableArray } from '../output/Data.Foldable/index.js';
import { Tuple } from '../output/Data.Tuple/index.js';
import { Just, Nothing } from '../output/Data.Maybe/index.js';
import { Left, Right } from '../output/Data.Either/index.js';
import * as Aff from '../output/Effect.Aff/index.js';
import { bind } from '../output/Control.Bind/index.js';
import { pure } from '../output/Control.Applicative/index.js';
import { liftEffect } from '../output/Effect.Class/index.js';
import * as Preparation from '../output/Gopurs.Preparation/index.js';
import * as Emission from '../output/Gopurs.Emission/index.js';

const codec = Wire.codecExprType;
const mapCodec = Wire.codecMap(ordString)(Wire.keyOrderString)(Wire.codecString)(codec);
const rows = [new Tuple('A.present', C.Int.value), new Tuple('B.new', new C.ForAll(['a'], new C.Func([new C.TypeVar('a')], new C.Record(new C.Row([new Tuple('z', C.String.value), new Tuple('a', C.Int.value)], new Just(new C.TypeVar('r')))))))];
const full = Map.fromFoldable(ordString)(foldableArray)(rows);
const coverage = Set.fromFoldable(foldableArray)(ordString)(['A.present', 'A.absent']);
const right = result => { assert(result instanceof Right, result.value0); return result.value0; };
const partial = Map.fromFoldable(ordString)(foldableArray)([rows[0]]);
const table = Table.covered(coverage)(partial)(Lazy.defer(_ => new Right(full))).value0;
const wireMap = values => Wire.write(mapCodec)(values)();
const wireTable = values => Wire.write(Table.codecLookupTable(codec))(values)();
const run = action => new Promise((resolve, reject) => Aff.runAff(result => () => result instanceof Left ? reject(result.value0) : resolve(result.value0))(action)());
const flatMap = bind(Aff.bindAff), succeed = pure(Aff.applicativeAff), effect = liftEffect(Aff.monadEffectAff);

function corruptTable() {
  return Table.covered(coverage)(partial)(Lazy.defer(_ => new Left('complete metadata codec'))).value0;
}

test('covered and new present/absent reads preserve exact structural types', () => {
  for (const key of ['A.present', 'A.absent', 'B.new', 'B.absent']) {
    const expected = Map.lookup(ordString)(key)(full);
    assert.deepEqual(right(Table.lookupChecked(key)(table)), expected);
    assert.deepEqual(Table.lookup(key)(table), expected);
    assert.equal(Table.member(key)(table), expected instanceof Just);
  }
});

test('opaque codec retains exact complete Map bytes and decoding returns a complete table', () => {
  assert.equal(wireTable(table), wireMap(full));
  assert.equal(wireTable(Table.fromMap(full)), wireMap(full));
  const decoded = Wire.read(Table.codecLookupTable(codec))(wireMap(full));
  assert(decoded instanceof Just);
  assert.equal(wireTable(decoded.value0), wireMap(full));
  assert.deepEqual(Table.lookup('B.new')(decoded.value0), new Just(rows[1].value1));
});

test('values outside the coverage domain are rejected', () => {
  assert(Table.covered(coverage)(full)(Lazy.defer(_ => new Right(full))) instanceof Nothing);
});

test('a repeated malformed complete decode remains an explicit error; covered reads stay lazy', () => {
  const bad = corruptTable();
  assert.deepEqual(Table.lookup('A.present')(bad), new Just(C.Int.value));
  assert.deepEqual(Table.lookup('A.absent')(bad), Nothing.value);
  for (let i = 0; i < 3; i++) {
    const failure = Table.lookupChecked('B.new')(bad);
    assert(failure instanceof Left);
    assert.equal(failure.value0, 'complete metadata codec');
    assert.throws(() => Table.lookup('B.new')(bad), /incremental cache: complete metadata codec/);
    assert.throws(() => wireTable(bad), /incremental cache: complete metadata codec/);
  }
});

test('actual preparation jobs and emitter catch before a single fallback/publication', async () => {
  for (const jobs of [1, 8]) for (const pipelined of [false, true]) {
    const bad = corruptTable();
    let active = 0, completed = 0, fallbacks = 0, publications = 0;
    const build = (current, force) => {
      const emit = values => flatMap(Preparation.runPreparationJobs(jobs)(values.map(value => () => {
        active++;
        try { return Table.lookup(value === 0 ? 'B.new' : 'A.present')(current); }
        finally { active--; completed++; }
      })))(_ => succeed(undefined));
      const produce = enqueue => {
        let action = succeed(undefined);
        for (let value = 0; value < 8; value++) action = flatMap(action)(_ => enqueue({ name: 'Module' + value, imports: Set.empty, value }));
        return action;
      };
      return flatMap(Aff.attempt(Emission.withEmitter({ jobs, pipelined })(emit)(produce)))(result => {
        assert.equal(active, 0, 'no computation may survive into fallback/publication');
        if (result instanceof Right) return effect(() => { publications++; });
        assert.match(result.value0.message, /incremental cache:/);
        if (force) throw result.value0;
        fallbacks++;
        return build(Table.fromMap(full), true);
      });
    };
    await run(build(bad, false));
    assert.equal(fallbacks, 1); assert.equal(publications, 1); assert(completed >= 9);
  }
});

