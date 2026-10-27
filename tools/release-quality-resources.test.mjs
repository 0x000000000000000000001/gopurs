import assert from "node:assert/strict";
import { existsSync, mkdirSync, mkdtempSync, readFileSync, realpathSync, rmSync, symlinkSync, utimesSync, writeFileSync } from "node:fs";
import { tmpdir } from "node:os";
import { join } from "node:path";
import { setImmediate } from "node:timers/promises";
import test from "node:test";
import { fixtureBatches, fixtureWorkers, QualityGoCache, readFixtureReport, runFixtureWaves } from "./release-quality-resources.mjs";

function workspace(t) {
  const path = realpathSync(mkdtempSync(join(tmpdir(), "gopurs-quality-unit-")));
  t.after(() => rmSync(path, { recursive: true, force: true }));
  return path;
}

test("all 401 fixtures occur exactly once in each profile, including the last partial batch", () => {
  const targets = Array.from({ length: 401 }, (_, index) => `/fixtures/${index}.purs`);
  const batches = fixtureBatches(targets);
  assert.equal(batches.length, 162);
  for (const profile of ["release", "dev"]) {
    const selected = batches.filter(batch => batch.profile === profile);
    assert.deepEqual(selected.flatMap(batch => batch.targets), targets);
    assert(selected.every(batch => batch.targets.length <= 5));
    assert.equal(selected.at(-1).targets.length, 1);
  }
  assert.throws(() => fixtureBatches([]), /nonempty/);
  assert.throws(() => fixtureBatches(["same", "same"]), /unique/);
  assert.deepEqual([1, 2, 3, 4, 14].map(fixtureWorkers), [1, 1, 2, 3, 6]);
});

test("fixture waves bound concurrency and trim only after every process settles", async () => {
  let active = 0, maximum = 0, trims = 0;
  const seen = [];
  await runFixtureWaves([0, 1, 2, 3, 4], 2, async batch => {
    active++;
    maximum = Math.max(maximum, active);
    await setImmediate();
    seen.push(batch);
    active--;
  }, () => { assert.equal(active, 0); trims++; });
  assert.equal(maximum, 2);
  assert.equal(trims, 3);
  assert.deepEqual(seen, [0, 1, 2, 3, 4]);
});

for (const synchronous of [true, false]) test(`failed wave preserves active work and stops scheduling (${synchronous ? "sync" : "async"} failure)`, async () => {
  const seen = [];
  let siblingFinished = false, trimmed = false;
  const failure = new Error("mock compiler failed");
  const run = batch => {
    seen.push(batch);
    if (batch === 0) {
      if (synchronous) throw failure;
      return Promise.reject(failure);
    }
    return setImmediate().then(() => { siblingFinished = true; });
  };
  await assert.rejects(runFixtureWaves([0, 1, 2], 2, run, () => {
    assert(siblingFinished);
    trimmed = true;
  }), error => error === failure);
  assert.deepEqual(seen, [0, 1]);
  assert(trimmed);
});

test("interruption and disk guards prevent later waves without changing the error", async () => {
  const interruption = Object.assign(new Error("Interrupted by SIGTERM"), { exitCode: 143 });
  await assert.rejects(runFixtureWaves([0, 1], 1, () => Promise.reject(interruption), () => {}), error => error === interruption);
  let launched = false;
  await assert.rejects(runFixtureWaves([0], 1, () => { launched = true; }, () => {}, () => { throw new Error("disk floor"); }), /disk floor/);
  assert.equal(launched, false);
  await assert.rejects(runFixtureWaves([], 7, () => {}, () => {}), /between 1 and 6/);
});

test("private cache eviction preserves the caller cache and ignores symlink targets", t => {
  const root = workspace(t);
  const caller = join(root, "caller-cache");
  mkdirSync(caller);
  writeFileSync(join(caller, "sentinel"), "never remove");
  const cache = new QualityGoCache({ parent: root, limitBytes: 5 });
  const sourceEnv = { GOCACHE: caller, GOCACHEPROG: "caller-cache-program", KEEP: "yes" };
  assert.deepEqual(cache.environment(sourceEnv), { GOCACHE: cache.path, KEEP: "yes" });
  assert.equal(sourceEnv.GOCACHE, caller);
  assert.equal(sourceEnv.GOCACHEPROG, "caller-cache-program");
  mkdirSync(join(cache.path, "aa"));
  const old = join(cache.path, "aa/old"), recent = join(cache.path, "aa/recent");
  writeFileSync(old, "1234567");
  writeFileSync(recent, "123");
  utimesSync(old, 1, 1);
  utimesSync(recent, 2, 2);
  symlinkSync(caller, join(cache.path, "other-cache"));
  symlinkSync(join(caller, "sentinel"), join(cache.path, "aa/linked-object"));
  assert.deepEqual(cache.trim(), { beforeBytes: 10, afterBytes: 3, deletedObjects: 1 });
  assert.equal(existsSync(old), false);
  assert.equal(readFileSync(recent, "utf8"), "123");
  cache.dispose();
  assert.equal(existsSync(cache.path), false);
  assert.equal(readFileSync(join(caller, "sentinel"), "utf8"), "never remove");
});

test("cache cannot be trimmed or disposed during commands, including a failing command", async t => {
  const cache = new QualityGoCache({ parent: workspace(t), limitBytes: 0 });
  await assert.rejects(cache.withProcess(async () => {
    assert.throws(() => cache.trim(), /commands are active/);
    assert.throws(() => cache.dispose(), /commands are active/);
    await setImmediate();
    throw new Error("mock process failure");
  }), /mock process failure/);
  writeFileSync(join(cache.path, "object"), "cache bytes");
  assert.equal(cache.trim().afterBytes, 0);
  cache.dispose();
});

test("fixture reports must prove the exact profile, selection, completion and snapshot policy", t => {
  const root = workspace(t), file = join(root, "results.json"), log = join(root, "batch.log");
  const batch = { profile: "dev", targets: ["/fixtures/A.purs", "/fixtures/B.purs"] };
  const good = { version: 1, kind: "fixtures", root, buildProfile: "dev", updateSnapshots: false,
    skipped: [], finishedAt: "2026-10-27T00:00:00Z", targets: batch.targets.map(target => ({ target, status: "passed" })) };
  writeFileSync(log, `Results: ${file}\n`);
  writeFileSync(file, JSON.stringify(good));
  assert.deepEqual(readFixtureReport(log, root, batch), { file, profile: "dev", passed: 2 });
  const variants = [
    { updateSnapshots: true }, { buildProfile: "release" }, { root: "/different-root" },
    { skipped: ["A"] }, { finishedAt: null }, { targets: good.targets.slice(0, 1) },
    { targets: [...good.targets].reverse() },
    { targets: good.targets.map(row => ({ ...row, status: "failed" })) },
    { targets: [good.targets[0], good.targets[0]] },
  ];
  for (const change of variants) {
    const contents = JSON.stringify({ ...good, ...change });
    writeFileSync(file, contents);
    assert.throws(() => readFixtureReport(log, root, batch), /mismatched fixture report/);
    assert.equal(readFileSync(file, "utf8"), contents, "failed evidence must be preserved");
  }
  writeFileSync(log, "No fixture report\n");
  assert.throws(() => readFixtureReport(log, root, batch), /Missing or ambiguous/);
});
