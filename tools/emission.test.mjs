// After npm run build: node --test tools/emission.test.mjs
import assert from "node:assert/strict";
import { test } from "node:test";
import { createEmitter, createPipelinedEmitter } from "../output/Gopurs.Emission/index.js";
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

const nextTurn = () => new Promise(setImmediate);

function suspendedEmission() {
  const batches = [];
  const published = [];
  const cancelled = [];
  const emit = values => {
    // Snapshot at construction deliberately detects eager emit calls before join.
    const snapshot = published.slice();
    return Aff.makeAff(done => () => {
      const batch = {
        values, snapshot,
        complete: () => {
          published.push(...values);
          done(new Either.Right(undefined))();
        },
        fail: error => done(new Either.Left(error))(),
      };
      batches.push(batch);
      return Aff.effectCanceler(() => { cancelled.push(values); });
    });
  };
  return { batches, published, cancelled, emit };
}

test("pipelined jobs at or below one retain immediate emission", async () => {
  for (const jobs of [-1, 0, 1]) {
    const batches = [];
    const emitter = createPipelinedEmitter(jobs)(batch => effect(() => { batches.push(batch); }))();
    await runAff(emitter.enqueue(entry("A")));
    assert.deepEqual(batches, [["A"]]);
    await runAff(emitter.finish);
    await runAff(emitter.cancel);
    assert.deepEqual(batches, [["A"]]);
  }
});

test("pipelined emission orders batches and bounds batches in flight", async () => {
  const probe = suspendedEmission();
  const emitter = createPipelinedEmitter(2)(probe.emit)(); // cap = 4 batches
  for (const name of ["A", "B", "C", "D", "E", "F", "G", "H", "I"]) await runAff(emitter.enqueue(entry(name)));
  await nextTurn();
  // Four batches are in flight (only the first has started emitting); the
  // fifth launch blocks until the chain drains.
  let tenth = false;
  const enqueueJ = runAff(emitter.enqueue(entry("J"))).then(() => { tenth = true; });
  await nextTurn();
  assert.equal(tenth, false);
  assert.equal(probe.batches.length, 1);
  assert.deepEqual(probe.batches[0].values, ["A", "B"]);
  // Batches start strictly in order, each seeing the previous publications.
  probe.batches[0].complete();
  await nextTurn();
  assert.deepEqual(probe.batches[1].values, ["C", "D"]);
  assert.deepEqual(probe.batches[1].snapshot, ["A", "B"]);
  probe.batches[1].complete();
  await nextTurn();
  assert.deepEqual(probe.batches[2].snapshot, ["A", "B", "C", "D"]);
  probe.batches[2].complete();
  await nextTurn();
  assert.deepEqual(probe.batches[3].snapshot, ["A", "B", "C", "D", "E", "F"]);
  probe.batches[3].complete();
  await enqueueJ;
  await nextTurn();
  assert.deepEqual(probe.batches[4].values, ["I", "J"]);
  assert.deepEqual(probe.batches[4].snapshot, ["A", "B", "C", "D", "E", "F", "G", "H"]);
  const finish = runAff(emitter.finish);
  await nextTurn();
  probe.batches[4].complete();
  await finish;
  await nextTurn();
  assert.deepEqual(probe.published, ["A", "B", "C", "D", "E", "F", "G", "H", "I", "J"]);
});

test("pipelined dependency barriers keep dependent modules in separate ordered batches", async () => {
  const probe = suspendedEmission();
  const emitter = createPipelinedEmitter(4)(probe.emit)();
  await runAff(emitter.enqueue(entry("A")));
  await runAff(emitter.enqueue(entry("B", ["A"])));
  await nextTurn();
  assert.deepEqual(probe.batches.map(batch => batch.values), [["A"]]);
  const finish = runAff(emitter.finish);
  await nextTurn();
  assert.equal(probe.batches.length, 1);
  probe.batches[0].complete();
  await nextTurn();
  assert.deepEqual(probe.batches[1].values, ["B"]);
  assert.deepEqual(probe.batches[1].snapshot, ["A"]);
  probe.batches[1].complete();
  await finish;
});

test("pipelined worker failure is sticky and never starts the queued batch", async () => {
  const probe = suspendedEmission();
  const emitter = createPipelinedEmitter(2)(probe.emit)();
  await runAff(emitter.enqueue(entry("A")));
  await runAff(emitter.enqueue(entry("B")));
  await nextTurn();
  await runAff(emitter.enqueue(entry("C")));
  const failure = new Error("worker failed");
  probe.batches[0].fail(failure);
  await nextTurn();
  assert.equal(probe.batches.length, 1, "the chained batch must not start");
  await assert.rejects(runAff(emitter.enqueue(entry("D"))), error => error === failure);
  await runAff(emitter.cancel);
  await runAff(emitter.finish);
  assert.equal(probe.batches.length, 1);
});

test("pipelined finish propagates the final worker failure", async () => {
  const probe = suspendedEmission();
  const emitter = createPipelinedEmitter(2)(probe.emit)();
  await runAff(emitter.enqueue(entry("A")));
  const failure = new Error("final worker failed");
  const rejected = assert.rejects(runAff(emitter.finish), error => error === failure);
  await nextTurn();
  probe.batches[0].fail(failure);
  await rejected;
  await runAff(emitter.cancel);
});

test("pipelined cancellation drains non-cancellable work and discards the pending batch", async () => {
  const events = [];
  let complete;
  const emit = values => Aff.makeAff(done => () => {
    events.push(["start", values]);
    complete = () => {
      events.push(["completed", values]);
      done(new Either.Right(undefined))();
    };
    return Aff.nonCanceler;
  });
  const emitter = createPipelinedEmitter(2)(emit)();
  await runAff(emitter.enqueue(entry("A")));
  await runAff(emitter.enqueue(entry("B")));
  await runAff(emitter.enqueue(entry("C")));
  let cancelled = false;
  const cancel = runAff(emitter.cancel).then(() => { cancelled = true; });
  await nextTurn();
  assert.deepEqual(events, [["start", ["A", "B"]]]);
  assert.equal(cancelled, false);
  complete();
  await cancel;
  await runAff(emitter.cancel);
  await runAff(emitter.finish);
  await assert.rejects(runAff(emitter.enqueue(entry("D"))), /Go emission cancelled/);
  assert.deepEqual(events, [
    ["start", ["A", "B"]], ["completed", ["A", "B"]],
  ]);
});
