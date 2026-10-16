import assert from "node:assert/strict";
import { spawnSync } from "node:child_process";
import { createHash } from "node:crypto";
import { copyFileSync, existsSync, mkdirSync, mkdtempSync, readFileSync, readdirSync, rmSync, statfsSync, symlinkSync, writeFileSync } from "node:fs";
import { tmpdir } from "node:os";
import { dirname, join } from "node:path";
import { fileURLToPath } from "node:url";
import test from "node:test";
import { findTypedCompiler } from "./native-workspace.mjs";
import { formatSnapshots, snapshotFiles } from "./test-snapshots.mjs";
import { corePackages, prepareFixture } from "./test-workspace.mjs";

const root = fileURLToPath(new URL("../", import.meta.url));
const hash = bytes => createHash("sha256").update(bytes).digest("hex");
const fixtures = ["NativeRecordWorkers", "FFIIntegerReturns", "StringEdgeCases", "NativeRecordBoxing",
  "NativeRecordReturns", "NativeArrayReboxing", "NativeTraverseCallback", "JsonRecordPlan", "B8xNativeGeneration"];

function inventory(output) {
  const go = {}, tast = {};
  function visit(directory, prefix = "") {
    for (const entry of readdirSync(directory, { withFileTypes: true })) {
      if (entry.name.startsWith(".")) continue;
      const name = prefix + entry.name, path = join(directory, entry.name);
      if (entry.isDirectory()) visit(path, name + "/");
      else if (name.endsWith(".go")) go[name] = hash(readFileSync(path));
      else if (entry.name === "corefn.json") tast[name] = hash(readFileSync(path));
    }
  }
  visit(output);
  return { go, tast };
}

test("real package consumers preserve the monolithic ABI and execution across compiler hosts", async t => {
  const workspace = mkdtempSync(join(process.env.GOPURS_TEST_EVIDENCE ?? tmpdir(), "package-consumers-"));
  const modes = (process.env.GOPURS_TEST_MODES ?? "js").split(",");
  const compiler = process.env.GOPURS_TEST_COMPILER ?? join(root, "bin/gopurs");
  let passed = false;
  t.after(() => {
    if (passed && !process.env.GOPURS_TEST_EVIDENCE) rmSync(workspace, { recursive: true, force: true });
    else console.log("Package consumer evidence: " + workspace);
  });
  mkdirSync(join(workspace, "typed-bin"));
  symlinkSync(findTypedCompiler(root), join(workspace, "typed-bin/purs"));
  const env = { ...process.env, GOWORK: "off", GOMAXPROCS: "2", GOMEMLIMIT: "12GiB",
    PATH: join(workspace, "typed-bin") + ":" + join(root, "node_modules/.bin") + ":" + process.env.PATH };
  for (const name of ["GOPURS_FORCE_BUILD", "GOPURS_COMPILER_SNAPSHOT", "GOPURS_ALLOC_PROFILE", "PPROF"]) delete env[name];
  const commands = [], results = [];
  function run(label, command, args, cwd, overrides = {}) {
    const disk = statfsSync(workspace);
    assert(disk.bavail * disk.bsize > 8 * 1024 ** 3, "disk floor");
    const result = spawnSync(command, args, { cwd, env: { ...env, ...overrides }, encoding: "utf8", timeout: 600000, maxBuffer: 32 * 1024 * 1024 });
    const log = commands.length + "-" + label + ".log";
    writeFileSync(join(workspace, log), (result.stdout ?? "") + (result.stderr ?? ""));
    commands.push({ label, command, args, cwd, overrides, status: result.status, log });
    writeFileSync(join(workspace, "commands.json"), JSON.stringify(commands, null, 2));
    assert.ifError(result.error);
    assert.equal(result.status, 0, label + ": " + log);
    return result;
  }
  for (const [index, name] of fixtures.entries()) await t.test(name, async () => {
    const fixture = prepareFixture(root, workspace, join(root, "tests/passing", name + ".purs"), index, corePackages(root));
    const project = fixture.directory, output = join(project, "output");
    run(name + "-frontend", "spago", ["build", "-q", "--offline"], project);
    const tast = inventory(output).tast;
    const sequential = { GOPURS_JS: "1", GOPURS_RUST: "0", GOPURS_GO_PACKAGES: "0",
      GOPURS_PREPARE_JOBS: "1", GOPURS_PBO_JOBS: "1", GOPURS_EMIT_JOBS: "1", GOPURS_PIPELINE: "0" };
    run(name + "-monolithic", compiler, ["--main", "Main", "--force"], project, sequential);
    run(name + "-tidy", "go", ["mod", "tidy"], output);
    const binary = join(project, "regression");
    run(name + "-monolithic-build", "go", ["build", "-p", "1", "-o", binary, "./main/main.go"], output);
    const oracle = run(name + "-monolithic-execute", binary, [], output).stdout;
    assert.match(oracle, /Done/);
    assert.doesNotMatch(oracle, /Fail/);
    const monolithic = run(name + "-monolithic-snapshot", "gofmt", [join(output, "purescript/Main.go")], output).stdout;
    writeFileSync(join(project, "monolithic-main.go"), monolithic);
    let reference;
    for (const mode of modes) {
      const flags = { GOPURS_GO_PACKAGES: "1", GOPURS_JS: mode === "js" ? "1" : "0", GOPURS_RUST: mode === "rust" ? "1" : "0",
        GOPURS_PREPARE_JOBS: mode === "native-parallel" ? "8" : "1", GOPURS_PBO_JOBS: mode === "native-parallel" ? "8" : "1",
        GOPURS_EMIT_JOBS: mode === "native-parallel" ? "8" : "1", GOPURS_PIPELINE: mode === "native-parallel" ? "1" : "0" };
      run(name + "-" + mode + "-packages", compiler, ["--main", "Main", "--force"], project, flags);
      const actual = inventory(output);
      assert.deepEqual(actual.tast, tast, "the same TAST must be used by every host");
      if (reference) assert.deepEqual(actual.go, reference, name + ": host parity");
      else reference = actual.go;
      writeFileSync(join(project, mode + "-inventory.json"), JSON.stringify(actual, null, 2));
      run(name + "-" + mode + "-dag", "go", ["list", "-deps", "-json", "./main"], output);
      run(name + "-" + mode + "-build", "go", ["build", "-p", "1", "-o", binary, "./main/main.go"], output);
      assert.equal(run(name + "-" + mode + "-execute", binary, [], output).stdout, oracle);
    }
    // Format a review copy: the raw emitted inventory stays available for replay.
    const view = { ...fixture, directory: join(project, "snapshot-view") };
    mkdirSync(join(view.directory, "logs"), { recursive: true });
    const sourceFiles = snapshotFiles(root, fixture, { GOPURS_GO_PACKAGES: "1" })[0].sources;
    for (const name of sourceFiles) {
      const target = join(view.directory, "output", name);
      mkdirSync(dirname(target), { recursive: true });
      copyFileSync(join(output, name), target);
    }
    const snapshots = snapshotFiles(root, view, { GOPURS_GO_PACKAGES: "1" });
    await formatSnapshots(snapshots, { run: async (label, command, args, options) => run(name + "-" + label, command, args, options.cwd) }, view, env);
    const row = { name, modes, tastModules: Object.keys(tast).length, files: Object.keys(reference).length,
      oracle, snapshot: snapshots[0].generated, expected: snapshots[0].expected, snapshotVerified: false };
    results.push(row);
    writeFileSync(join(workspace, "results.json"), JSON.stringify(results, null, 2));
    assert(existsSync(row.expected), "Missing package snapshot; review " + row.snapshot);
    assert.equal(readFileSync(row.snapshot, "utf8"), readFileSync(row.expected, "utf8"), name + ": strict package snapshot");
    row.snapshotVerified = true;
    writeFileSync(join(workspace, "results.json"), JSON.stringify(results, null, 2));
  });
  passed = results.length === fixtures.length && results.every(row => row.snapshotVerified);
  writeFileSync(join(workspace, "summary.json"), JSON.stringify({ completed: passed, fixtures: results.length, modes, commands: commands.length }, null, 2));
});
