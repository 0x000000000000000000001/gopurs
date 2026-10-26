#!/usr/bin/env node
import { spawnSync } from "node:child_process";
import { existsSync, mkdirSync, readFileSync, readdirSync, rmSync, writeFileSync } from "node:fs";
import { join, resolve } from "node:path";
import { fileURLToPath } from "node:url";

const root = fileURLToPath(new URL("../", import.meta.url));
let full = false;
let report;
for (let index = 2; index < process.argv.length; index++) {
  const arg = process.argv[index];
  if (arg === "--help") {
    console.log("Usage: node tools/release-quality.mjs [--full] [--report DIRECTORY]\nRequires a fresh JS/native build, typed purs and locked library checkouts. --full runs all Node tests, Go race tests, and Release/Dev fixtures. Rust remains outside this gate.");
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
if (report) mkdirSync(report, { recursive: true });
const results = [];
function run(name, command, args, cwd = root, extraEnv = {}) {
  console.log(`\n[release quality] ${name}`);
  const started = Date.now();
  const result = spawnSync(command, args, { cwd, stdio: "inherit", env: { ...process.env, GOWORK: "off", GOPURS_GO_PACKAGES: "0", GOPURS_TEST_RUST: "0", ...extraEnv } });
  results.push({ name, command: [command, ...args], passed: result.status === 0, exitCode: result.status, signal: result.signal, milliseconds: Date.now() - started });
  if (report) writeFileSync(join(report, "quality.json"), JSON.stringify({ full, gates: results }, null, 2) + "\n");
  if (result.error || result.status !== 0) throw new Error(`${name} failed: ${result.error?.message ?? result.signal ?? result.status}`);
}
const quick = ["ffi-errors", "launcher", "test-runner", "build-native", "bootstrap-inputs", "cli", "build-profiles", "driver-cache"].map(name => `${name}.test.mjs`);
const tests = full ? readdirSync(join(root, "tools")).filter(name => name.endsWith(".test.mjs")).sort() : quick;
run("Node compiler and tooling tests", process.execPath, ["--test", "--test-concurrency=1", ...tests.map(name => `tools/${name}`)], root,
  full ? { GOPURS_TEST_MODES: "js,native-sequential,native-parallel" } : {});
for (const module of ["build-cache", "c1c", "ffi-gen"]) run(`Go ${module}${full ? " (race)" : ""}`, "go", ["test", ...(full ? ["-race"] : []), "-count=1", "./..."], join(root, "tools", module));
if (full) {
  const emissionTest = join(process.env.GOPURS_NATIVE_OUTPUT, "purescript/emission_native_test.go");
  writeFileSync(emissionTest, readFileSync(join(root, "tools/emission-native_test.go")), { flag: "wx" });
  try {
    run("Native emission cancellation and concurrency (race)", "go", ["test", "-race", "-count=1", "-run", "^TestPipelineNative", "-timeout", "30s", "./purescript"], process.env.GOPURS_NATIVE_OUTPUT);
  } finally {
    rmSync(emissionTest, { force: true });
  }
  run("All Release fixtures", "bash", ["bin/test", "--all", "--keep-going"], root, { GOPURS_JS: "0", GOPURS_RUST: "0" });
  run("All Dev fixtures", "bash", ["bin/test", "--all", "--keep-going", "--dev"], root, { GOPURS_JS: "0", GOPURS_RUST: "0" });
} else run("Representative executable fixtures", "bash", ["bin/test", "FFIIntegerReturns", "TCO", "ThunkFusion", "Int32Loops", "StringEdgeCases"], root, { GOPURS_JS: "0", GOPURS_RUST: "0" });
run("npm package surface", "npm", ["run", "test:package"]);
console.log(`Release quality: ${results.length} gates passed (${full ? "full" : "quick"}).`);
