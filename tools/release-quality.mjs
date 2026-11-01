#!/usr/bin/env node
import { copyFileSync, existsSync, mkdirSync, mkdtempSync, readFileSync, readdirSync, renameSync, rmSync, statfsSync, writeFileSync } from "node:fs";
import { tmpdir } from "node:os";
import { join, relative } from "node:path";
import { fileURLToPath } from "node:url";
import { isDeepStrictEqual } from "node:util";
import { CommandRunner } from "./command-runner.mjs";
import { parseOptions, selectFixtures } from "./test-selection.mjs";
import { corePackages, prepareFixture } from "./test-workspace.mjs";
import { fixtureBatches, fixtureWorkers, QualityGoCache, readFixtureReport, runFixtureWaves } from "./release-quality-resources.mjs";
import { emissionGateName, packageGateName, qualityIdentity, qualityOptions, shardBatches } from "./release-quality-report.mjs";
import { collectModuleEvidence, moduleGateName, moduleInventory, verifyModuleInputs } from "./release-quality-modules.mjs";

const root = fileURLToPath(new URL("../", import.meta.url));
if (process.argv.includes("--help")) {
  console.log("Usage: node tools/release-quality.mjs [--full [--phase core|fixtures --shard INDEX --shard-count COUNT]] [--report DIRECTORY]\nRequires a fresh JS/native build, typed purs and locked libraries. Core/all phases also require GOPURS_NATIVE_OUTPUT from --keep-workspace. --full without --phase runs the entire campaign. Partial phases never claim full qualification. Rust remains outside this gate.");
  process.exit(0);
}
const { full, phase, shard, report: reportOption } = qualityOptions(process.argv.slice(2));
let report = reportOption;
if (!existsSync(join(root, "bin/gopurs-native"))) throw new Error("Build the native compiler before running release quality gates");
if (phase !== "fixtures" && !existsSync(join(process.env.GOPURS_NATIVE_OUTPUT ?? "/missing-native-output", "purescript"))) {
  throw new Error("Release quality requires GOPURS_NATIVE_OUTPUT from build:native -- --keep-workspace, so native tests cannot silently skip");
}
const identity = qualityIdentity(root);
report ??= mkdtempSync(join(tmpdir(), "gopurs-quality-"));
mkdirSync(report, { recursive: true });
if (existsSync(join(report, "quality.json"))) throw new Error(`Quality report already exists: ${report}`);
const workspaces = join(report, "workspaces");
mkdirSync(workspaces, { recursive: true });
const cache = new QualityGoCache();
const workers = fixtureWorkers();
const results = [], cacheTrims = [];
let sequence = 0, status = "running", failure, fixtureCount, fixturePreparation, moduleCount;
const baseEnv = { ...process.env, GOWORK: "off", GOPURS_GO_PACKAGES: "0", GOPURS_TEST_RUST: "0",
  GOPURS_TEST_MODES: "js,native-sequential,native-parallel",
  GOMAXPROCS: "2", GOGC: "100", GOMEMLIMIT: "4GiB", UPDATE_SNAPSHOTS: "0",
  TMPDIR: workspaces, TMP: workspaces, TEMP: workspaces };
delete baseEnv.GOPURS_COMPILER_SNAPSHOT;
// Host-parity tests must select this checkout's launcher, then switch hosts.
// A caller's direct native executable would bypass their GOPURS_JS setting.
delete baseEnv.GOPURS_TEST_COMPILER;
const fixtureEnv = { GOPURS_JS: "0", GOPURS_RUST: "0", GOPURS_TEST_COMPILER: join(root, "bin/gopurs-native"),
  GOPURS_PREPARE_JOBS: "1", GOPURS_PBO_JOBS: "1", GOPURS_EMIT_JOBS: "2", GOFLAGS: "-trimpath" };
function save() {
  const file = join(report, "quality.json");
  writeFileSync(file + ".tmp", JSON.stringify({ schema: 1, requestedFull: full, full: full && phase === "all", phase, shard,
    identity, status, error: failure, fixtureWorkers: workers, fixturePreparation,
    fixtures: fixtureCount === undefined ? undefined : {
      selectedPerProfile: fixtureCount,
      passed: Object.fromEntries(["release", "dev"].map(profile => [profile,
        results.reduce((count, gate) => count + (gate.passed && gate.fixtures?.profile === profile ? gate.fixtures.passed : 0), 0)])),
    },
    modules: moduleCount === undefined ? undefined : {
      selectedPerProfile: moduleCount,
      passed: Object.fromEntries(["release", "dev"].map(profile => [profile,
        results.reduce((count, gate) => count + (gate.passed && gate.modules?.profile === profile ? gate.modules.report.targets.length : 0), 0)])),
    },
    goCache: { path: cache.path, limitBytes: cache.limitBytes, trims: cacheTrims },
    gates: [...results].sort((a, b) => a.index - b.index) }, null, 2) + "\n");
  renameSync(file + ".tmp", file);
}
function trimCache() {
  cacheTrims.push(cache.trim());
  save();
}
async function run(name, command, args, cwd = root, extraEnv = {}, verify) {
  const index = ++sequence;
  const log = join(report, `${String(index).padStart(3, "0")}-${name.replace(/[^a-zA-Z0-9_-]+/g, "-")}.log`);
  console.log(`\n[release quality] ${name} — ${log}`);
  const started = Date.now();
  const runner = new CommandRunner();
  const gate = { index, name, command: [command, ...args], log, passed: false };
  try {
    await cache.withProcess(async () => {
      const result = await runner.run(command, args, { cwd, log, env: cache.environment({ ...baseEnv, ...extraEnv }) });
      Object.assign(gate, { exitCode: result.code, signal: result.signal });
      runner.checkInterrupted();
      if (result.code !== 0) throw new Error(`${name} failed: ${result.signal ?? result.code}; see ${log}`);
      if (verify) Object.assign(gate, verify(log));
      gate.passed = true;
    });
  } catch (error) {
    gate.error = error.message;
    throw error;
  } finally {
    runner.dispose();
    gate.milliseconds = Date.now() - started;
    results.push(gate);
    save();
  }
}
async function runSingle(...args) {
  try { await run(...args); }
  finally { trimCache(); }
}

async function initializeFixtureDependencies(target) {
  const log = join(report, "fixture-preparation.log");
  const started = Date.now(), runner = new CommandRunner();
  let workspace;
  fixturePreparation = { command: ["spago", "fetch"], target: relative(root, target), log, passed: false };
  save();
  try {
    workspace = mkdtempSync(join(workspaces, "spago-initialize-"));
    const { directory } = prepareFixture(root, workspace, target, 0, corePackages(root));
    const temporary = join(directory, "tmp");
    mkdirSync(temporary);
    console.log(`Initialize shared Spago dependencies before fixture waves — ${log}`);
    const result = await runner.run("spago", ["fetch"], { cwd: directory, log,
      env: { ...baseEnv, ...fixtureEnv, TMPDIR: temporary, TMP: temporary, TEMP: temporary } });
    Object.assign(fixturePreparation, { exitCode: result.code, signal: result.signal });
    runner.checkInterrupted();
    if (result.code !== 0) {
      process.stderr.write(readFileSync(log, "utf8"));
      throw new Error(`Spago fixture preparation failed: ${result.signal ?? result.code}; see ${log}`);
    }
    fixturePreparation.passed = true;
  } catch (error) {
    fixturePreparation.error = error.message;
    throw error;
  } finally {
    runner.dispose();
    if (workspace) rmSync(workspace, { recursive: true, force: true });
    fixturePreparation.milliseconds = Date.now() - started;
    save();
  }
}

console.log(`Release quality evidence: ${report}`);
save();
try {
  const quick = ["release-quality-resources", "release-npm-provenance", "release-sbom", "emission", "ffi-errors", "launcher", "test-runner", "build-native", "bootstrap-inputs", "cli", "build-profiles", "manifest-prepared-cache", "driver-cache", "native-go-code", "preparation-native", "check-output-identity", "check-docs", "embed-runtime", "compiler-snapshot"].map(name => `${name}.test.mjs`);
  const tests = full ? readdirSync(join(root, "tools")).filter(name => name.endsWith(".test.mjs")).sort() : quick;
  // node --test isolates files in child processes too. Separate invocations let
  // us bound the shared private Go cache between those same file boundaries.
  if (phase !== "fixtures") {
    for (const name of tests) await runSingle(`Node ${name}`, process.execPath, ["--test", "--test-concurrency=1", `tools/${name}`]);
    for (const module of ["build-cache", "c1c", "ffi-gen"]) await runSingle(`Go ${module}${full ? " (race)" : ""}`, "go", ["test", ...(full ? ["-race"] : []), "-count=1", "./..."], join(root, "tools", module));
  }
  if (full && phase !== "fixtures") {
    const emissionTest = join(process.env.GOPURS_NATIVE_OUTPUT, "purescript/emission_native_test.go");
    writeFileSync(emissionTest, readFileSync(join(root, "tools/emission-native_test.go")), { flag: "wx" });
    try {
      // Match the bounded functional-probe profile; preserve race detection
      // and the runtime deadline. The release bootstrap remains optimized.
      await runSingle(emissionGateName, "go", ["test", "-race", "-gcflags=gopurs/output/purescript=-N -l", "-count=1", "-run", "^TestPipelineNative", "-timeout", "30s", "./purescript"], process.env.GOPURS_NATIVE_OUTPUT);
    } finally {
      rmSync(emissionTest, { force: true });
    }
    const modules = moduleInventory(root);
    verifyModuleInputs(root, modules);
    moduleCount = modules.length;
    save();
    for (const profile of ["release", "dev"]) {
      // Keep the complete module inventory, but trim the private Go cache
      // between libraries instead of accumulating an entire campaign's cache.
      for (const library of modules) await runSingle(moduleGateName(profile, library.name), "bash",
        ["bin/modtest", library.name, ...(profile === "dev" ? ["--dev"] : [])], root, fixtureEnv,
        log => ({ modules: collectModuleEvidence(log, root, profile, [library], report) }));
    }
    verifyModuleInputs(root, modules);
  }
  if (full && phase !== "core") {
    const { selected, skipped } = selectFixtures(root, parseOptions(["--all"], { env: { UPDATE_SNAPSHOTS: "0" } }));
    if (skipped.length) throw new Error("Release quality requires every fixture, without exclusions");
    fixtureCount = selected.length;
    save();
    console.log(`Full fixture selection: ${selected.length} per profile, ${workers} concurrent batches of at most 5.`);
    const batches = fixtureBatches(selected).map((batch, batchIndex) => ({ ...batch, batchIndex }));
    const assignedBatches = shard ? shardBatches(batches, shard) : batches;
    // Spago's cold shared registry must be initialized before concurrent clients.
    // Fetch does not compile; the unchanged gates still test every fixture.
    if (assignedBatches.length) await initializeFixtureDependencies(assignedBatches[0].targets[0]);
    mkdirSync(join(report, "fixture-reports"), { recursive: true });
    await runFixtureWaves(assignedBatches, workers, batch => {
      const name = `${batch.profile} fixtures ${batch.start + 1}-${batch.start + batch.targets.length}`;
      return run(name, "bash", ["bin/test", ...batch.targets, "--keep-going", ...(batch.profile === "dev" ? ["--dev"] : [])], root, fixtureEnv,
        log => {
          const evidence = readFixtureReport(log, root, batch);
          const file = `fixture-reports/${String(batch.batchIndex + 1).padStart(3, "0")}-${batch.profile}-${batch.start + 1}-${batch.start + batch.targets.length}.json`;
          copyFileSync(evidence.file, join(report, file));
          return { fixtures: { ...evidence, file, batchIndex: batch.batchIndex, targets: batch.targets.map(target => relative(root, target)) } };
        });
    }, trimCache, wave => {
      const disk = statfsSync(workspaces);
      const required = (3 + wave.length * 0.75) * 1024 ** 3;
      if (disk.bavail * disk.bsize < required) throw new Error(`Fixture wave needs ${(required / 1024 ** 3).toFixed(2)} GiB free; reports retained at ${report}`);
    });
  } else if (!full) {
    await runSingle("Representative executable fixtures", "bash", ["bin/test", "FFIIntegerReturns", "TCO", "ThunkFusion", "Int32Loops", "StringEdgeCases"], root, fixtureEnv);
  }
  if (phase !== "fixtures") await runSingle(packageGateName, "npm", ["run", "test:package"]);
  if (!isDeepStrictEqual(qualityIdentity(root), identity)) throw new Error("Source or compiler identity changed during quality gates");
  status = "passed";
  console.log(`Release quality: ${results.length} gates passed (${full ? phase === "all" ? "full" : `partial ${phase}` : "quick"}).`);
} catch (error) {
  status = "failed";
  failure = error.message;
  console.error(error.message);
  process.exitCode = error.exitCode ?? 1;
} finally {
  cache.dispose();
  save();
  console.log(`Retained quality logs and reports: ${report}`);
}
