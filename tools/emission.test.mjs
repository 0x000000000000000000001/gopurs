// After npm run build: node --test tools/emission.test.mjs
import assert from "node:assert/strict";
import { test } from "node:test";
import { createEmitter } from "../output/Gopurs.Emission/index.js";
import * as Aff from "../output/Effect.Aff/index.js";
import * as Either from "../output/Data.Either/index.js";
import * as Set from "../output/Data.Set/index.js";
import { foldableArray } from "../output/Data.Foldable/index.js";
import { ordModuleName } from "../output/PureScript.Backend.Optimizer.CoreFn/index.js";

const imports = Set.fromFoldable(foldableArray)(ordModuleName);
const entry = (name, dependencies = []) => ({ name, imports: imports(dependencies), value: name });
const runAff = aff => new Promise((resolve, reject) => {
  Aff.runAff(result => () => result instanceof Either.Left
    ? reject(result.value0) : resolve(result.value0))(aff)();
});
const effect = Aff.monadEffectAff.liftEffect;

test("jobs at or below one emit each module immediately", async () => {
  for (const jobs of [-1, 0, 1]) {
    const batches = [];
    const emitter = createEmitter(jobs)(batch => effect(() => { batches.push(batch); }))();
    await runAff(emitter.enqueue(entry("A")));
    assert.deepEqual(batches, [["A"]]);
    await runAff(emitter.enqueue(entry("B", ["A"])));
    assert.deepEqual(batches, [["A"], ["B"]]);
    await runAff(emitter.finish);
    assert.deepEqual(batches, [["A"], ["B"]]);
  }
});

test("capacity bounds batches and finish emits the partial batch in order", async () => {
  const batches = [];
  const emitter = createEmitter(3)(batch => effect(() => { batches.push(batch); }))();
  for (const name of ["A", "B"]) await runAff(emitter.enqueue(entry(name)));
  assert.deepEqual(batches, []);
  for (const name of ["C", "D", "E", "F", "G"]) await runAff(emitter.enqueue(entry(name)));
  assert.deepEqual(batches, [["A", "B", "C"], ["D", "E", "F"]]);
  await runAff(emitter.finish);
  await runAff(emitter.finish);
  assert.deepEqual(batches, [["A", "B", "C"], ["D", "E", "F"], ["G"]]);
});

test("a pending import waits for its batch before the dependent module is queued", async () => {
  const batches = [];
  let release;
  const emitter = createEmitter(4)(batch => Aff.makeAff(done => () => {
    batches.push(batch);
    release = () => done(new Either.Right(undefined))();
    return Aff.nonCanceler;
  }))();
  await runAff(emitter.enqueue(entry("A")));
  await runAff(emitter.enqueue(entry("B", ["AlreadyEmitted"])));
  let completed = false;
  const enqueueDependent = runAff(emitter.enqueue(entry("C", ["B"]))).then(() => { completed = true; });
  await new Promise(setImmediate);
  assert.deepEqual(batches, [["A", "B"]]);
  assert.equal(completed, false);
  release();
  await enqueueDependent;
  await runAff(emitter.enqueue(entry("D")));
  const finish = runAff(emitter.finish);
  assert.deepEqual(batches, [["A", "B"], ["C", "D"]]);
  release();
  await finish;
});

test("a failed dependency flush rejects without emitting or queueing the successor", async () => {
  const batches = [];
  const failure = new Error("emission failed");
  let fail = true;
  const emitter = createEmitter(4)(batch => effect(() => {
    batches.push(batch);
    if (fail) throw failure;
  }))();
  await runAff(emitter.enqueue(entry("A")));
  await runAff(emitter.enqueue(entry("B")));
  await assert.rejects(runAff(emitter.enqueue(entry("C", ["A"]))), error => error === failure);
  assert.deepEqual(batches, [["A", "B"]]);
  fail = false;
  await runAff(emitter.finish);
  assert.deepEqual(batches, [["A", "B"], ["A", "B"]]);
  await runAff(emitter.finish);
  assert.equal(batches.length, 2);
});
