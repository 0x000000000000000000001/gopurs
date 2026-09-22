import assert from "node:assert/strict";
import { spawn } from "node:child_process";
import { cpSync, existsSync, mkdirSync, mkdtempSync, readFileSync, readdirSync, realpathSync, rmSync, writeFileSync } from "node:fs";
import { tmpdir } from "node:os";
import { join } from "node:path";
import { fileURLToPath } from "node:url";
import { test } from "node:test";

const repository = fileURLToPath(new URL("../", import.meta.url));
const fakeTool = `#!/usr/bin/env node
import { appendFileSync, existsSync, mkdirSync, readFileSync, writeFileSync } from "node:fs";
import { spawn } from "node:child_process";
import { basename, dirname, join } from "node:path";
const tool = basename(process.argv[1]);
const cwd = process.cwd();
const fixture = basename(cwd) === "output" ? dirname(cwd) : cwd;
const source = join(fixture, "src/Main.purs");
const content = existsSync(source) ? readFileSync(source, "utf8") : "";
const phase = tool === "go" ? (process.argv[2] === "mod" ? "tidy" : "build") : tool;
appendFileSync(process.env.RUNNER_CALLS, JSON.stringify({ phase, cwd }) + "\\n");
if (content.includes("fail=" + phase) || process.env.RUNNER_FAIL === phase) process.exit(23);
if (process.env.RUNNER_PAUSE && phase === "spago") {
  const child = spawn(process.execPath, ["-e", "setInterval(() => {}, 1000)"], { stdio: "ignore" });
  writeFileSync(process.env.RUNNER_PAUSE, JSON.stringify({ pid: process.pid, child: child.pid }));
  setInterval(() => {}, 1000);
} else if (phase === "spago") {
  mkdirSync("output", { recursive: true });
  writeFileSync("spago.lock", "fixture-only lock");
  writeFileSync("output/module.json", JSON.stringify({ source: content }));
} else if (phase === "gopurs") {
  mkdirSync("output/purescript", { recursive: true });
  mkdirSync("output/main", { recursive: true });
  writeFileSync("output/purescript/Main.go", "package purescript\\n");
  writeFileSync("output/purescript/Main_ffi.go", "package purescript // ffi\\n");
  writeFileSync("output/main/main.go", "package main\\n");
  writeFileSync("output/go.mod", "module gopurs/output\\n\\ngo 1.22\\n");
} else if (phase === "build") {
  const text = content.includes("fail=text") ? "Fail" : "Done";
  const code = content.includes("fail=execute") ? 9 : 0;
  writeFileSync("gopurs_main", "#!/usr/bin/env node\\nconsole.log(" + JSON.stringify(text) + "); process.exit(" + code + ");\\n", { mode: 0o755 });
}
`;

function fixture(t) {
  const base = mkdtempSync(join(tmpdir(), "gopurs-runner-contract-"));
  t.after(() => rmSync(base, { recursive: true, force: true }));
  const root = join(base, "project with spaces");
  for (const dir of ["bin", "tools", "tests/passing", "tests/passing-snapshots", "tests/runner"]) mkdirSync(join(root, dir), { recursive: true });
  for (const dir of ["mockbin", "tmp", "gopurs-prelude"]) mkdirSync(join(base, dir));
  writeFileSync(join(base, "package.json"), '{"type":"module"}');
  for (const name of ["test", "modtest"]) cpSync(join(repository, "bin", name), join(root, "bin", name));
  for (const name of ["test-selection", "test-process", "test-workspace", "test-snapshots", "test-runner", "modtest-runner"]) cpSync(join(repository, "tools", name + ".mjs"), join(root, "tools", name + ".mjs"));
  writeFileSync(join(root, "bin/pkg"), "CORE_PACKAGES=(prelude)\n");
  for (const name of ["spago", "gofmt", "go", "npm"]) writeFileSync(join(base, "mockbin", name), fakeTool, { mode: 0o755 });
  writeFileSync(join(root, "bin/gopurs"), fakeTool, { mode: 0o755 });
  for (const name of ["Alpha", "Beta", "DerivingClause"]) writeFileSync(join(root, "tests/passing", name + ".purs"), `module Main where\n-- ${name}\n`);
  for (const name of ["Alpha", "Beta"]) writeFileSync(join(root, "tests/passing-snapshots", name + ".go"), "package purescript\n");
  writeFileSync(join(root, "tests/runner/spago.yaml"), "shared config sentinel");
  writeFileSync(join(root, "tests/runner/spago.lock"), "shared lock sentinel");
  const calls = join(base, "calls.jsonl");
  const env = { ...process.env, TMPDIR: join(base, "tmp"), PATH: join(base, "mockbin") + ":" + process.env.PATH, RUNNER_CALLS: calls, UPDATE_SNAPSHOTS: "0" };
  const start = (args, extraEnv = {}, command = "test") => {
    const child = spawn("bash", [join(root, "bin", command), ...args], { cwd: base, env: { ...env, ...extraEnv } });
    let output = "";
    child.stdout.on("data", chunk => output += chunk);
    child.stderr.on("data", chunk => output += chunk);
    const completed = new Promise((resolve, reject) => {
      child.on("error", reject);
      child.on("close", (code, signal) => resolve({ code, signal, output }));
    });
    return { child, completed };
  };
  return { base, root, calls, env, start, run: (...args) => start(...args).completed, source: (name, text) => writeFileSync(join(root, "tests/passing", name + ".purs"), text) };
}

function assertSharedUntouched(f) {
  assert.equal(readFileSync(join(f.root, "tests/runner/spago.yaml"), "utf8"), "shared config sentinel");
  assert.equal(readFileSync(join(f.root, "tests/runner/spago.lock"), "utf8"), "shared lock sentinel");
}

test("selection is read-only, ordered, inclusive and rejects unknown names", async t => {
  const f = fixture(t);
  const all = await f.run(["--list"]);
  assert.equal(all.code, 0, all.output);
  assert.deepEqual(all.output.trim().split("\n").map(line => line.split("/").at(-1)), ["Alpha.purs", "Beta.purs"]);
  const ordered = await f.run(["Beta", "Alpha", "--list", "--skip-before", "Alpha"]);
  assert.equal(ordered.code, 0, ordered.output);
  assert.match(ordered.output, /Alpha.purs\n$/);
  assert.doesNotMatch(ordered.output, /Beta.purs/);
  for (const args of [["Missing"], ["Alpha", "Missing"], ["Alpha", "--skip-before=Missing"], ["--unknown"], ["DerivingClause"], ["--all", "Alpha"]]) {
    const result = await f.run(args);
    assert.equal(result.code, 2, result.output);
  }
  assert(!existsSync(f.calls));
  assert.deepEqual(readdirSync(join(f.base, "tmp")), []);
  assertSharedUntouched(f);
});

test("fixtures own their sources, companions, config, lock and output", async t => {
  const f = fixture(t);
  f.source("Alpha", "-- @dependencies: prelude\nmodule Main where\n");
  writeFileSync(join(f.root, "tests/passing/Alpha.go"), "// companion");
  mkdirSync(join(f.root, "tests/passing/Alpha"));
  writeFileSync(join(f.root, "tests/passing/Alpha/Extra.purs"), "module Extra where\n");
  const result = await f.run(["Alpha", "Beta", "--keep-workspace"]);
  assert.equal(result.code, 0, result.output);
  const workspace = result.output.match(/Kept workspace: (.+)/)[1];
  assert(existsSync(join(workspace, "0-Alpha/src/Extra.purs")));
  assert(!existsSync(join(workspace, "1-Beta/src/Extra.purs")));
  assert(!existsSync(join(workspace, "1-Beta/src/Main.go")));
  assert.equal(readFileSync(join(workspace, "0-Alpha/src/Main.purs"), "utf8"), readFileSync(join(f.root, "tests/passing/Alpha.purs"), "utf8"));
  assert.match(readFileSync(join(workspace, "0-Alpha/spago.yaml"), "utf8"), new RegExp(f.base));
  assert.match(readFileSync(join(workspace, "0-Alpha/output/go.mod"), "utf8"), /go 1.22/);
  assertSharedUntouched(f);
  const simultaneous = await Promise.all([f.run(["Alpha"]), f.run(["Beta"])]);
  for (const run of simultaneous) {
    assert.equal(run.code, 0, run.output);
    assert(!existsSync(run.output.match(/Workspace: (.+)/)[1]));
  }
  assert.notEqual(...simultaneous.map(run => run.output.match(/Workspace: (.+)/)[1]));
});

test("explicit dependencies use native checkouts beyond the core package set", async t => {
  const f = fixture(t);
  f.source("Alpha", "-- @dependencies: prelude native-extra registry-only\nmodule Main where\n");
  mkdirSync(join(f.base, "gopurs-native-extra"));
  writeFileSync(join(f.base, "gopurs-native-extra/spago.yaml"), "package:\n  name: native-extra\n");
  const result = await f.run(["Alpha", "--keep-workspace"]);
  assert.equal(result.code, 0, result.output);
  const workspace = result.output.match(/Kept workspace: (.+)/)[1];
  const config = readFileSync(join(workspace, "0-Alpha/spago.yaml"), "utf8");
  assert.match(config, /    - registry-only/);
  assert.match(config, /    native-extra:\n      path: .*gopurs-native-extra/);
  assert.doesNotMatch(config, /    registry-only:/);
  assert.equal((config.match(/    prelude:/g) ?? []).length, 1);
});

test("snapshot creation and replacement require opt-in, including FFI", async t => {
  const f = fixture(t);
  f.source("Alpha", "-- @snapshot-ffi\nmodule Main where\n");
  const missing = await f.run(["Alpha"]);
  assert.equal(missing.code, 1, missing.output);
  assert.match(missing.output, /Missing snapshot/);
  const ffi = join(f.root, "tests/passing-snapshots/Alpha_ffi.go");
  assert(!existsSync(ffi));
  const update = await f.run(["Alpha", "--update-snapshots"]);
  assert.equal(update.code, 0, update.output);
  assert.equal(readFileSync(ffi, "utf8"), "package purescript // ffi\n");
  const expected = join(f.root, "tests/passing-snapshots/Alpha.go");
  writeFileSync(expected, "different snapshot\n");
  const mismatch = await f.run(["Alpha"]);
  assert.equal(mismatch.code, 1, mismatch.output);
  assert.match(mismatch.output, /snapshot diff/);
  assert.equal(readFileSync(expected, "utf8"), "different snapshot\n");
  const legacy = await f.run(["Alpha"], { UPDATE_SNAPSHOTS: "1" });
  assert.equal(legacy.code, 0, legacy.output);
  assert.equal(readFileSync(expected, "utf8"), "package purescript\n");
});

test("every failed phase stops the run and keeps diagnostics without updating snapshots", async t => {
  const f = fixture(t);
  const expected = join(f.root, "tests/passing-snapshots/Alpha.go");
  writeFileSync(expected, "original snapshot\n");
  for (const phase of ["spago", "gopurs", "gofmt", "tidy", "build", "execute", "text"]) {
    f.source("Alpha", `module Main where\n-- fail=${phase}\n`);
    const result = await f.run(["Alpha", "Beta", "--update-snapshots"]);
    assert.equal(result.code, 1, result.output);
    assert.doesNotMatch(result.output, /=> Testing Beta/);
    const workspace = result.output.match(/Kept workspace: (.+)/)[1];
    assert(existsSync(join(workspace, "0-Alpha/logs")));
    assert.equal(readFileSync(expected, "utf8"), "original snapshot\n");
    assertSharedUntouched(f);
  }
  const build = await f.run(["Alpha", "-c"], { RUNNER_FAIL: "npm" });
  assert.equal(build.code, 1, build.output);
  assert.doesNotMatch(build.output, /=> Testing Alpha/);
});

async function waitFor(check) {
  for (let i = 0; i < 200; i++) {
    if (check()) return;
    await new Promise(resolve => setTimeout(resolve, 25));
  }
  assert.fail("Timed out waiting for subprocess state");
}

for (const [signal, code] of [["SIGINT", 130], ["SIGTERM", 143]]) test(`${signal} stops compiler descendants and preserves the shared workspace`, async t => {
  const f = fixture(t);
  const pause = join(f.base, "pause.json");
  const running = f.start(["Alpha"], { RUNNER_PAUSE: pause });
  await waitFor(() => existsSync(pause));
  const pids = JSON.parse(readFileSync(pause, "utf8"));
  running.child.kill(signal);
  const result = await running.completed;
  assert.equal(result.code, code, result.output);
  for (const pid of [pids.pid, pids.child]) await waitFor(() => {
    try { process.kill(pid, 0); return false; } catch (error) { return error.code === "ESRCH"; }
  });
  assert.match(result.output, /Kept workspace/);
  assertSharedUntouched(f);
});

test("module campaigns list all eligible siblings, resume explicitly and stop on failure", async t => {
  const f = fixture(t);
  for (const name of ["alpha", "beta", "strings"]) {
    const bin = join(f.base, "gopurs-" + name, "bin");
    mkdirSync(bin, { recursive: true });
    writeFileSync(join(bin, "test"), `#!/usr/bin/env bash\necho module-${name}\n${name === "beta" ? "exit 7" : "exit 0"}\n`, { mode: 0o755 });
  }
  const all = await f.run(["--list"], {}, "modtest");
  assert.equal(all.code, 0, all.output);
  assert.equal(all.output, "gopurs-alpha\ngopurs-beta\ngopurs-strings\n");
  const resume = await f.run(["--list", "skip_before=strings"], {}, "modtest");
  assert.equal(resume.output, "gopurs-strings\n");
  const unknown = await f.run(["--skip-before=missing"], {}, "modtest");
  assert.equal(unknown.code, 2, unknown.output);
  const run = await f.run(["alpha", "beta", "strings"], {}, "modtest");
  assert.equal(run.code, 1, run.output);
  assert.match(run.output, /module-alpha/);
  assert.doesNotMatch(run.output, /module-strings/);
  const clean = await f.run(["alpha", "-c"], {}, "modtest");
  assert.equal(clean.code, 0, clean.output);
  assert.equal(realpathSync(JSON.parse(readFileSync(f.calls, "utf8").trim()).cwd), realpathSync(f.root));
});
