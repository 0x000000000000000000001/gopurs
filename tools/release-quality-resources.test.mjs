import assert from "node:assert/strict";
import { execFileSync, spawnSync } from "node:child_process";
import { createHash } from "node:crypto";
import { existsSync, mkdirSync, mkdtempSync, readFileSync, realpathSync, rmSync, symlinkSync, utimesSync, writeFileSync } from "node:fs";
import { tmpdir } from "node:os";
import { join } from "node:path";
import { fileURLToPath } from "node:url";
import { setImmediate } from "node:timers/promises";
import test from "node:test";
import { fixtureBatches, fixtureWorkers, QualityGoCache, readFixtureReport, runFixtureWaves } from "./release-quality-resources.mjs";
import { coreGateNames, fixtureShardCount, mergeQualityReports, qualityInventory, qualityOptions, shardBatches } from "./release-quality-report.mjs";
import { moduleGateName } from "./release-quality-modules.mjs";

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

test("partial phases require full intent and an explicit valid fixture shard", () => {
  assert.deepEqual(qualityOptions([]), { full: false, phase: "all", shard: null, report: undefined });
  assert.equal(qualityOptions(["--full"]).phase, "all");
  assert.equal(qualityOptions(["--full", "--phase", "core"]).shard, null);
  assert.deepEqual(qualityOptions(["--full", "--phase", "fixtures", "--shard", "3", "--shard-count", "4"]).shard, { index: 3, count: 4 });
  for (const args of [
    ["--phase", "core"], ["--full", "--phase", "unknown"], ["--full", "--shard", "0", "--shard-count", "4"],
    ["--full", "--phase", "fixtures"], ["--full", "--phase", "fixtures", "--shard", "4", "--shard-count", "4"],
    ["--full", "--phase", "fixtures", "--shard", "0", "--shard-count", "0"],
    ["--full", "--phase", "fixtures", "--shard", "0.5", "--shard-count", "4"],
  ]) assert.throws(() => qualityOptions(args));
});

function partialQualityReports() {
  const inventory = qualityInventory(fileURLToPath(new URL("../", import.meta.url)));
  const identity = { source: { commit: "a".repeat(40), dirty: false },
    compiler: { javascriptSha256: "b".repeat(64), nativeSha256: "c".repeat(64) }, target: "linux-amd64",
    github: { repository: "owner/gopurs", runId: "123", runAttempt: "1" } };
  const common = { schema: 1, requestedFull: true, full: false, status: "passed", identity };
  const gate = (name, index) => ({ index: index + 1, name, passed: true, exitCode: 0, signal: null });
  const core = { ...common, phase: "core", shard: null, gates: coreGateNames(inventory).map(gate) };
  const startedAt = "2026-10-30T00:00:00Z", finishedAt = "2026-10-30T00:01:00Z";
  for (const profile of ["release", "dev"]) for (const library of inventory.modules) {
    core.gates.find(gate => gate.name === moduleGateName(profile, library.name)).modules = {
      profile, inputs: [library], file: `module-reports/${profile}/${library.name}.json`,
      report: { version: 1, kind: "modules", root: "/checkout", startedAt, finishedAt,
        resumedFrom: null, buildProfile: profile, updateSnapshots: false, skipped: [],
        targets: [{ target: library.name, status: "passed", directory: `/tmp/modules/${library.name}`, startedAt, finishedAt }] },
      logs: [{ target: library.name, file: `module-reports/${profile}/${library.name}.log`, sha256: "f".repeat(64) }],
    };
  }
  const fixtures = Array.from({ length: fixtureShardCount }, (_, index) => {
    const shard = { index, count: fixtureShardCount }, batches = shardBatches(inventory.batches, shard);
    return { ...common, phase: "fixtures", shard,
      fixtures: { selectedPerProfile: inventory.targets.length, passed: Object.fromEntries(["release", "dev"].map(profile =>
        [profile, batches.filter(batch => batch.profile === profile).reduce((sum, batch) => sum + batch.targets.length, 0)])) },
      gates: batches.map((batch, i) => ({ ...gate(batch.name, i), fixtures: { batchIndex: batch.batchIndex, profile: batch.profile,
        passed: batch.targets.length, targets: batch.targets, file: `fixture-reports/${batch.batchIndex}-${batch.profile}.json` } })),
    };
  });
  const build = { schema: 1, version: "0.1.0-rc.2", target: identity.target, source: { ...identity.source },
    artifacts: [{ name: "gopurs-0.1.0-rc.2-linux-amd64", sha256: identity.compiler.nativeSha256 }] };
  return structuredClone({ inventory, core, fixtures, build });
}

test("four shards preserve every original batch and only their verified union is full", () => {
  const { inventory, core, fixtures, build } = partialQualityReports();
  assert.equal(inventory.nodeTests.length, 100);
  assert.equal(inventory.targets.length, 401);
  assert.equal(inventory.modules.length, 51);
  assert.equal(core.full, false);
  assert(fixtures.every(report => report.full === false));
  assert.deepEqual(fixtures.map(report => report.gates.length), [41, 41, 40, 40]);
  for (const report of fixtures) {
    const release = report.gates.filter(gate => gate.fixtures.profile === "release").length;
    const dev = report.gates.filter(gate => gate.fixtures.profile === "dev").length;
    assert(release > 0 && dev > 0 && Math.abs(release - dev) <= 1,
      "every shard must balance Release and Dev instead of isolating a profile");
  }
  assert.deepEqual(fixtures.flatMap(report => report.gates.map(gate => gate.fixtures.batchIndex)).sort((a, b) => a - b),
    inventory.batches.map(batch => batch.batchIndex));
  const merged = mergeQualityReports(core, fixtures.toReversed(), inventory, build);
  assert.equal(merged.full, true);
  assert.equal(merged.phase, "all");
  assert.equal(merged.status, "passed");
  assert.equal(merged.gates.length, 369);
  assert.deepEqual(merged.fixtures.passed, { release: 401, dev: 401 });
  assert.deepEqual(merged.modules, { selectedPerProfile: 51, passed: { release: 51, dev: 51 } });
  assert.deepEqual(merged.gates.map(gate => gate.index), Array.from({ length: 369 }, (_, index) => index + 1));
  for (const profile of ["release", "dev"]) assert.deepEqual(
    merged.gates.filter(gate => gate.fixtures?.profile === profile).flatMap(gate => gate.fixtures.targets), inventory.targets);
  assert.equal(merged.gates.at(-1).name, "npm package surface");
});

test("aggregation rejects incomplete, duplicate, failed or mismatched qualification evidence", () => {
  const changes = [
    value => { value.fixtures.pop(); },
    value => { value.fixtures[1] = structuredClone(value.fixtures[0]); },
    value => { value.fixtures[0].shard.count = 5; },
    value => { value.fixtures[0].shard.index = 4; },
    value => { value.fixtures[0].gates.pop(); },
    value => { value.fixtures[0].gates[1] = structuredClone(value.fixtures[0].gates[0]); },
    value => { value.fixtures[0].gates[0].fixtures.targets[0] = "tests/passing/omitted.purs"; },
    value => { value.fixtures[0].gates[0].fixtures.profile = "dev"; },
    value => { value.fixtures[0].gates[0].fixtures.batchIndex = 1; },
    value => { value.fixtures[0].fixtures.passed.release--; },
    value => { value.fixtures[0].identity.source.commit = "d".repeat(40); },
    value => { value.fixtures[0].identity.compiler.javascriptSha256 = "d".repeat(64); },
    value => { value.fixtures[0].identity.compiler.nativeSha256 = "d".repeat(64); },
    value => { value.fixtures[0].identity.github.runAttempt = "2"; },
    value => { value.fixtures[0].identity.target = "linux-arm64"; },
    value => { value.core.identity.source.dirty = true; },
    value => { value.core.gates.pop(); },
    value => { value.core.gates[0].passed = false; },
    value => { value.core.gates[0].exitCode = 1; },
    value => { delete value.core.gates.find(gate => gate.modules).modules; },
    value => { value.core.gates.find(gate => gate.modules).modules.report.targets.pop(); },
    value => { value.core.gates.find(gate => gate.modules).modules.report.targets[0].status = "failed"; },
    value => { value.core.gates.find(gate => gate.modules).modules.profile = "dev"; },
    value => { value.core.gates.find(gate => gate.modules).modules.inputs[0].rev = "d".repeat(40); },
    value => { value.core.gates.find(gate => gate.modules).modules.logs.pop(); },
    value => { value.fixtures[0].full = true; },
    value => { value.fixtures[0].status = "running"; },
    value => { value.core.requestedFull = false; },
    value => { value.build.source.commit = "d".repeat(40); },
    value => { value.build.artifacts[0].sha256 = "d".repeat(64); },
    value => { value.build.target = "linux-arm64"; },
  ];
  for (const change of changes) {
    const value = partialQualityReports();
    // Each input is independent, as it is after loading separate JSON files.
    value.fixtures = value.fixtures.map(report => structuredClone(report));
    value.core = structuredClone(value.core);
    change(value);
    assert.throws(() => mergeQualityReports(value.core, value.fixtures, value.inventory, value.build), change.toString());
  }
});

test("merge CLI checks stable fixture evidence and original bootstrap hashes before writing", t => {
  const root = workspace(t), { core, fixtures, build } = partialQualityReports();
  const commit = execFileSync("git", ["rev-parse", "HEAD"], { cwd: fileURLToPath(new URL("../", import.meta.url)), encoding: "utf8" }).trim();
  for (const report of [core, ...fixtures]) report.identity.source.commit = commit;
  build.source.commit = commit;
  const json = (file, value) => writeFileSync(file, JSON.stringify(value));
  const coreFile = join(root, "core.json"), buildFile = join(root, "build.json"), output = join(root, "merged.json");
  json(coreFile, core);
  const bootstrap = `${core.identity.compiler.javascriptSha256}  bin/gopurs.js\n`;
  const bootstrapName = "gopurs-0.1.0-rc.2-linux-amd64-bootstrap-checksums.txt";
  writeFileSync(join(root, bootstrapName), bootstrap);
  build.artifacts.push({ name: bootstrapName, sha256: createHash("sha256").update(bootstrap).digest("hex") });
  json(buildFile, build);
  const fixtureFiles = fixtures.map(report => {
    const directory = join(root, `shard-${report.shard.index}`);
    mkdirSync(join(directory, "fixture-reports"), { recursive: true });
    for (const gate of report.gates) json(join(directory, gate.fixtures.file), {
      version: 1, kind: "fixtures", root: "/original-checkout", buildProfile: gate.fixtures.profile,
      updateSnapshots: false, skipped: [], finishedAt: "2026-10-28T00:00:00Z",
      targets: gate.fixtures.targets.map(target => ({ target: join("/original-checkout", target), status: "passed" })),
    });
    const file = join(directory, "quality.json");
    json(file, report);
    return file;
  });
  const env = { ...process.env };
  for (const key of ["GITHUB_RUN_ID", "GITHUB_RUN_ATTEMPT", "GITHUB_REPOSITORY", "GITHUB_SHA"]) delete env[key];
  const args = [fileURLToPath(new URL("./release-quality-report.mjs", import.meta.url)), "--core", coreFile,
    ...fixtureFiles.flatMap(file => ["--fixtures", file]), "--build", buildFile, "--out", output];
  let result = spawnSync(process.execPath, args, { env, encoding: "utf8" });
  assert.equal(result.status, 0, result.stderr);
  assert.equal(JSON.parse(readFileSync(output, "utf8")).gates.length, 369);
  rmSync(output);
  const signingEnv = { ...env, CI: "true", GITHUB_ACTIONS: "true", RUNNER_ENVIRONMENT: "github-hosted",
    GITHUB_SERVER_URL: "https://github.com", GITHUB_EVENT_NAME: "push", GITHUB_REF_TYPE: "tag",
    GITHUB_SHA: commit, GITHUB_REPOSITORY: "owner/gopurs", GITHUB_RUN_ID: "123", GITHUB_RUN_ATTEMPT: "2",
    GITHUB_REF: "refs/tags/v0.1.0-rc.2", GITHUB_REF_NAME: "v0.1.0-rc.2", GITHUB_WORKFLOW_SHA: commit,
    GITHUB_WORKFLOW_REF: "owner/gopurs/.github/workflows/ci.yml@refs/tags/v0.1.0-rc.2",
    GOPURS_RELEASE_TARGET: "linux-amd64" };
  const jobs = ["linux-amd64", "linux-arm64", "darwin-arm64"].flatMap(target =>
    [`build and test (${target})`, ...[0, 1, 2, 3].map(shard => `fixtures (${target}, shard ${shard})`)])
    .map((name, index) => ({ id: index + 1, name, attempt: "1" }));
  const context = { schema: 1, source: { commit, tag: "v0.1.0-rc.2", tagObject: "e".repeat(40) },
    target: "linux-amd64", github: { repository: "owner/gopurs", runId: "123", qualificationAttempt: "1", finalizationAttempt: "2" },
    jobs, artifacts: ["candidate-linux-amd64", ...[0, 1, 2, 3].map(shard => `fixtures-linux-amd64-${shard}`)]
      .map((name, index) => ({ id: 100 + index, name, producerJobId: index + 1, attempt: "1", size: 123, digest: `sha256:${"f".repeat(64)}` })),
    qualificationIdentity: core.identity };
  const contextFile = join(root, "finalization.json"); json(contextFile, context);
  result = spawnSync(process.execPath, [...args, "--finalization-context", contextFile], { env: signingEnv, encoding: "utf8" });
  assert.equal(result.status, 0, result.stderr);
  const signed = JSON.parse(readFileSync(output, "utf8"));
  assert.deepEqual(signed.identity, core.identity);
  assert.deepEqual(signed.finalization, context);
  rmSync(output);
  result = spawnSync(process.execPath, args, { env: signingEnv, encoding: "utf8" });
  assert.notEqual(result.status, 0, "a later attempt requires its verified finalization context");
  assert.equal(existsSync(output), false);
  context.qualificationIdentity = structuredClone(core.identity);
  context.qualificationIdentity.compiler.nativeSha256 = "f".repeat(64); json(contextFile, context);
  result = spawnSync(process.execPath, [...args, "--finalization-context", contextFile], { env: signingEnv, encoding: "utf8" });
  assert.notEqual(result.status, 0, "a different original compiler cannot be rebound to the report");
  assert.equal(existsSync(output), false);
  result = spawnSync(process.execPath, args, { env: { ...env, GITHUB_SHA: "d".repeat(40) }, encoding: "utf8" });
  assert.notEqual(result.status, 0);
  assert.match(result.stderr, /Quality source differs from GITHUB_SHA/);
  assert.equal(existsSync(output), false);
  const firstEvidence = join(root, "shard-0", fixtures[0].gates[0].fixtures.file);
  const goodEvidence = readFileSync(firstEvidence, "utf8");
  const broken = JSON.parse(goodEvidence);
  broken.targets[0].status = "failed";
  json(firstEvidence, broken);
  result = spawnSync(process.execPath, args, { env, encoding: "utf8" });
  assert.notEqual(result.status, 0);
  assert.match(result.stderr, /Stable fixture evidence differs/);
  assert.equal(existsSync(output), false);
  writeFileSync(firstEvidence, goodEvidence);
  writeFileSync(join(root, bootstrapName), "changed manifest\n");
  result = spawnSync(process.execPath, args, { env, encoding: "utf8" });
  assert.notEqual(result.status, 0);
  assert.match(result.stderr, /Original bootstrap checksums differ/);
  assert.equal(existsSync(output), false);
});
