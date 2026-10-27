#!/usr/bin/env node
import { existsSync, mkdirSync, mkdtempSync, readFileSync, readdirSync, renameSync, rmSync, statfsSync, writeFileSync } from "node:fs";
import { tmpdir } from "node:os";
import { join, resolve } from "node:path";
import { fileURLToPath } from "node:url";
import { CommandRunner } from "./command-runner.mjs";
import { parseOptions, selectFixtures } from "./test-selection.mjs";
import { fixtureBatches, fixtureWorkers, QualityGoCache, readFixtureReport, runFixtureWaves } from "./release-quality-resources.mjs";

const root = fileURLToPath(new URL("../", import.meta.url));
let full = false;
let report;
for (let index = 2; index < process.argv.length; index++) {
  const arg = process.argv[index];
  if (arg === "--help") {
    console.log("Usage: node tools/release-quality.mjs [--full] [--report DIRECTORY]\nRequires a fresh JS/native build, typed purs and locked library checkouts. --full runs all Node tests, Go race tests, and Release/Dev fixtures. Uses a private bounded Go cache; retains logs and reports. Rust remains outside this gate.");
    process.exit(0);
  }
  if (arg === "--full") full = true;
  else if (arg === "--report" && process.argv[index + 1]) report = resolve(process.argv[++index]);
  else throw new Error(`Unknown or incomplete option: ${arg}`);
}
if (!existsSync(join(root, "bin/gopurs-native"))) throw new Error("Build the native compiler before running release quality gates");
if (full && !existsSync(join(process.env.GOPURS_NATIVE_OUTPUT ?? "/missing-native-output", "purescript"))) {
  throw new Error("--full requires GOPURS_NATIVE_OUTPUT from build:native -- --keep-workspace, so native tests cannot silently skip");
}
report ??= mkdtempSync(join(tmpdir(), "gopurs-quality-"));
mkdirSync(report, { recursive: true });
if (existsSync(join(report, "quality.json"))) throw new Error(`Quality report already exists: ${report}`);
const workspaces = join(report, "workspaces");
mkdirSync(workspaces, { recursive: true });
const cache = new QualityGoCache();
const workers = fixtureWorkers();
const results = [], cacheTrims = [];
let sequence = 0, status = "running", failure, fixtureCount;
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
  writeFileSync(file + ".tmp", JSON.stringify({ full, status, error: failure, fixtureWorkers: workers,
    fixtures: fixtureCount === undefined ? undefined : {
      selectedPerProfile: fixtureCount,
      passed: Object.fromEntries(["release", "dev"].map(profile => [profile,
        results.reduce((count, gate) => count + (gate.passed && gate.fixtures?.profile === profile ? gate.fixtures.passed : 0), 0)])),
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
      if (verify) gate.fixtures = verify(log);
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

console.log(`Release quality evidence: ${report}`);
save();
try {
  const quick = ["release-quality-resources", "release-npm-provenance", "ffi-errors", "launcher", "test-runner", "build-native", "bootstrap-inputs", "cli", "build-profiles", "driver-cache"].map(name => `${name}.test.mjs`);
  const tests = full ? readdirSync(join(root, "tools")).filter(name => name.endsWith(".test.mjs")).sort() : quick;
  // node --test isolates files in child processes too. Separate invocations let
  // us bound the shared private Go cache between those same file boundaries.
  for (const name of tests) await runSingle(`Node ${name}`, process.execPath, ["--test", "--test-concurrency=1", `tools/${name}`]);
  for (const module of ["build-cache", "c1c", "ffi-gen"]) await runSingle(`Go ${module}${full ? " (race)" : ""}`, "go", ["test", ...(full ? ["-race"] : []), "-count=1", "./..."], join(root, "tools", module));
  if (full) {
    const emissionTest = join(process.env.GOPURS_NATIVE_OUTPUT, "purescript/emission_native_test.go");
    writeFileSync(emissionTest, readFileSync(join(root, "tools/emission-native_test.go")), { flag: "wx" });
    try {
      await runSingle("Native emission cancellation and concurrency (race)", "go", ["test", "-race", "-count=1", "-run", "^TestPipelineNative", "-timeout", "30s", "./purescript"], process.env.GOPURS_NATIVE_OUTPUT);
    } finally {
      rmSync(emissionTest, { force: true });
    }
    const { selected, skipped } = selectFixtures(root, parseOptions(["--all"], { env: { UPDATE_SNAPSHOTS: "0" } }));
    if (skipped.length) throw new Error("Release quality requires every fixture, without exclusions");
    fixtureCount = selected.length;
    save();
    console.log(`Full fixture selection: ${selected.length} per profile, ${workers} concurrent batches of at most 5.`);
    await runFixtureWaves(fixtureBatches(selected), workers, batch => {
      const name = `${batch.profile} fixtures ${batch.start + 1}-${batch.start + batch.targets.length}`;
      return run(name, "bash", ["bin/test", ...batch.targets, "--keep-going", ...(batch.profile === "dev" ? ["--dev"] : [])], root, fixtureEnv,
        log => readFixtureReport(log, root, batch));
    }, trimCache, wave => {
      const disk = statfsSync(workspaces);
      const required = (3 + wave.length * 0.75) * 1024 ** 3;
      if (disk.bavail * disk.bsize < required) throw new Error(`Fixture wave needs ${(required / 1024 ** 3).toFixed(2)} GiB free; reports retained at ${report}`);
    });
  } else {
    await runSingle("Representative executable fixtures", "bash", ["bin/test", "FFIIntegerReturns", "TCO", "ThunkFusion", "Int32Loops", "StringEdgeCases"], root, fixtureEnv);
  }
  await runSingle("npm package surface", "npm", ["run", "test:package"]);
  status = "passed";
  console.log(`Release quality: ${results.length} gates passed (${full ? "full" : "quick"}).`);
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
