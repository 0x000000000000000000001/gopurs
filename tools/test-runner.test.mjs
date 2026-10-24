import assert from "node:assert/strict";
import { spawn, spawnSync } from "node:child_process";
import { cpSync, existsSync, mkdirSync, mkdtempSync, readFileSync, readdirSync, realpathSync, rmSync, symlinkSync, writeFileSync } from "node:fs";
import { tmpdir } from "node:os";
import { join } from "node:path";
import { fileURLToPath } from "node:url";
import { test } from "node:test";
import { formatSnapshots, packageSourceSections, snapshotFiles } from "./test-snapshots.mjs";
import { parseOptions } from "./test-selection.mjs";

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
const phase = tool === "go" ? (process.argv[2] === "mod" ? "tidy" : "build") : /^gopurs(?:-native|-rust|\\.js)?$/.test(tool) ? "gopurs" : tool;
appendFileSync(process.env.RUNNER_CALLS, JSON.stringify({ phase, tool, args: process.argv.slice(2), cwd, tmp: process.env.TMPDIR }) + "\\n");
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
  if (process.env.GOPURS_GO_PACKAGES === "1") {
    const group = process.env.RUNNER_PACKAGE_SPLIT ? "p_0123456789abcdef0123456789abcdef" : "p_root";
    for (const dir of [group, "init_all", "init_external"]) mkdirSync("output/purescript/" + dir, { recursive: true });
    writeFileSync("output/purescript/" + group + "/Main.go", "package " + group + "\\n");
    writeFileSync("output/purescript/" + group + "/Main_ffi.go", "package " + group + " // ffi\\n");
    writeFileSync("output/purescript/" + group + "/zz_gopurs_order.go", "package " + group + "\\n");
    for (const name of ["init_all", "init_external"]) writeFileSync("output/purescript/" + name + "/init.go", "package " + name + "\\n");
  } else {
    writeFileSync("output/purescript/Main.go", "package purescript\\n");
    writeFileSync("output/purescript/Main_ffi.go", "package purescript // ffi\\n");
  }
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
  for (const name of ["command-runner", "test-selection", "test-process", "test-workspace", "test-snapshots", "test-campaign", "test-compiler", "test-runner", "modtest-runner"]) cpSync(join(repository, "tools", name + ".mjs"), join(root, "tools", name + ".mjs"));
  writeFileSync(join(root, "bin/pkg"), "CORE_PACKAGES=(prelude)\n");
  for (const name of ["spago", "gofmt", "go", "npm"]) writeFileSync(join(base, "mockbin", name), fakeTool, { mode: 0o755 });
  writeFileSync(join(root, "bin/gopurs"), fakeTool, { mode: 0o755 });
  for (const name of ["Alpha", "Beta"]) writeFileSync(join(root, "tests/passing", name + ".purs"), `module Main where\n-- ${name}\n`);
  for (const name of ["Alpha", "Beta"]) writeFileSync(join(root, "tests/passing-snapshots", name + ".go"), "package purescript\n");
  writeFileSync(join(root, "tests/runner/spago.yaml"), "shared config sentinel");
  writeFileSync(join(root, "tests/runner/spago.lock"), "shared lock sentinel");
  const calls = join(base, "calls.jsonl");
  const env = { ...process.env, TMPDIR: join(base, "tmp"), PATH: join(base, "mockbin") + ":" + process.env.PATH, RUNNER_CALLS: calls, UPDATE_SNAPSHOTS: "0", GOPURS_JS: "0", GOPURS_RUST: "0", GOPURS_DAEMON: "0" };
  const start = (args, extraEnv = {}, command = "test") => {
    const child = spawn("bash", [join(root, "bin", command), ...args], { cwd: base, env: { ...env, ...extraEnv } });
    let output = "";
    child.stdout.on("data", chunk => output += chunk);
    child.stderr.on("data", chunk => output += chunk);
    const completed = new Promise((resolve, reject) => {
      child.on("error", reject);
      child.on("close", (code, signal) => resolve({ code, signal, output }));
    });
    return { child, completed, get output() { return output; } };
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
  for (const args of [["Missing"], ["Alpha", "Missing"], ["Alpha", "--skip-before=Missing"], ["--unknown"], ["--all", "Alpha"]]) {
    const result = await f.run(args);
    assert.equal(result.code, 2, result.output);
  }
  assert(!existsSync(f.calls));
  assert.deepEqual(readdirSync(join(f.base, "tmp")), []);
  assertSharedUntouched(f);
});

test("Dev aliases preserve selection, clean options and the option terminator", () => {
  for (const modules of [false, true]) for (const flag of ["-d", "--dev"]) {
    const options = parseOptions(["Beta", flag, "-c", "--skip-before", "Alpha", "Alpha"], { modules, env: {} });
    assert.equal(options.dev, true);
    assert.equal(options.clean, true);
    assert.equal(options.resume, "Alpha");
    assert.deepEqual(options.targets, ["Beta", "Alpha"]);
    const literal = parseOptions(["--", flag, "-c"], { modules, env: {} });
    assert.equal(literal.dev, false);
    assert.equal(literal.clean, false);
    assert.deepEqual(literal.targets, [flag, "-c"]);
  }
  assert.equal(parseOptions([], { env: {} }).dev, false);
});

test("Dev fixtures forward only the backend flag and preserve Release snapshots", async t => {
  const f = fixture(t);
  const expected = join(f.root, "tests/passing-snapshots/Alpha.go");
  writeFileSync(expected, "Release snapshot sentinel\n");
  for (const flag of ["-d", "--dev"]) {
    const result = await f.run(["Beta", "Alpha", flag, "-c", "--skip-before", "Alpha"]);
    assert.equal(result.code, 0, result.output);
    assert.match(result.output, /without comparing Release snapshots/);
    assert.doesNotMatch(result.output, /=> Testing Beta/);
    assert.equal(reportFor(result).buildProfile, "dev");
  }
  const calls = readFileSync(f.calls, "utf8").trim().split("\n").map(JSON.parse);
  assert.deepEqual(calls.filter(call => call.phase === "gopurs").map(call => call.args), [
    ["--main", "Main", "--dev"], ["--main", "Main", "--dev"],
  ]);
  assert.deepEqual(calls.filter(call => call.phase === "npm").map(call => call.args), [
    ["run", "build:native", "--silent"], ["run", "build:native", "--silent"],
  ]);
  assert(calls.filter(call => call.phase !== "gopurs").every(call => !call.args.includes("--dev")));
  assert(!calls.some(call => call.phase === "gofmt"));
  assert.equal(readFileSync(expected, "utf8"), "Release snapshot sentinel\n");
  for (const [args, env] of [[["--dev", "--update-snapshots"], {}], [["-d"], { UPDATE_SNAPSHOTS: "1" }]]) {
    const result = await f.run(["Alpha", ...args], env);
    assert.equal(result.code, 2, result.output);
    assert.match(result.output, /cannot update Release snapshots/);
  }
  assert.equal(readFileSync(f.calls, "utf8").trim().split("\n").length, calls.length);
  assertSharedUntouched(f);
});

test("Dev flags reach native and JS compiler hosts through the public launcher", async t => {
  const f = fixture(t);
  cpSync(join(repository, "bin/gopurs"), join(f.root, "bin/gopurs"));
  for (const tool of ["gopurs-native", "gopurs.js"]) writeFileSync(join(f.root, "bin", tool), fakeTool, { mode: 0o755 });
  for (const [host, flag] of [["0", "-d"], ["1", "--dev"]]) {
    const result = await f.run(["Alpha", flag], { GOPURS_JS: host, GOPURS_RUST: "0", GOPURS_DAEMON: "0" });
    assert.equal(result.code, 0, result.output);
  }
  const calls = readFileSync(f.calls, "utf8").trim().split("\n").map(JSON.parse).filter(call => call.phase === "gopurs");
  assert.deepEqual(calls.map(call => [call.tool, call.args]), [
    ["gopurs-native", ["--main", "Main", "--dev"]],
    ["gopurs.js", ["--main", "Main", "--dev"]],
  ]);
});

test("clean campaigns rebuild the selected compiler host before running it", async t => {
  for (const runner of ["test", "modtest"]) for (const dev of [false, true]) {
    for (const [host, script, tool, env] of [
      ["native", "build:native", "gopurs-native", {}],
      ["JS", "build", "gopurs.js", { GOPURS_JS: "1" }],
      ["Rust", "build:rust", "gopurs-rust", { GOPURS_RUST: "1" }],
    ]) await t.test(`${runner}, ${dev ? "Dev" : "Release"}, ${host}`, async t => {
      const f = fixture(t);
      cpSync(join(repository, "bin/gopurs"), join(f.root, "bin/gopurs"));
      for (const name of ["gopurs-native", "gopurs.js", "gopurs-rust"]) {
        writeFileSync(join(f.root, "bin", name), fakeTool, { mode: 0o755 });
      }
      const bin = join(f.base, "gopurs-alpha/bin");
      mkdirSync(bin, { recursive: true });
      writeFileSync(join(bin, "test"), "#!/usr/bin/env bash\nset -e\n../gopurs/bin/gopurs --main Test.Main\n", { mode: 0o755 });
      const result = await f.run([runner === "test" ? "Alpha" : "alpha", "-c", ...(dev ? ["--dev"] : [])], { GOPURS_GO_PACKAGES: "0", ...env }, runner);
      assert.equal(result.code, 0, result.output);
      const calls = readFileSync(f.calls, "utf8").trim().split("\n").map(JSON.parse);
      assert.deepEqual(calls.filter(call => call.phase === "npm").map(call => call.args), [["run", script, "--silent"]]);
      assert.equal(calls[0].phase, "npm", "rebuild precedes every test phase");
      const compilerCalls = calls.filter(call => call.phase === "gopurs");
      assert.equal(compilerCalls.length, 1);
      assert.equal(compilerCalls[0].tool, tool);
      assert.equal(compilerCalls[0].args.includes("--dev"), dev);
      assert.equal(reportFor(result).buildProfile, dev ? "dev" : "release");
    });
  }
});

test("clean campaigns reject conflicting compiler hosts before building", async t => {
  for (const runner of ["test", "modtest"]) await t.test(runner, async t => {
    const f = fixture(t);
    const bin = join(f.base, "gopurs-alpha/bin");
    mkdirSync(bin, { recursive: true });
    writeFileSync(join(bin, "test"), "#!/usr/bin/env bash\nexit 0\n", { mode: 0o755 });
    const result = await f.run([runner === "test" ? "Alpha" : "alpha", "-c", "--dev"], { GOPURS_JS: "1", GOPURS_RUST: "1" }, runner);
    assert.equal(result.code, 2, result.output);
    assert.match(result.output, /[Ss]elect only one compiler host: GOPURS_RUST=1 or GOPURS_JS=1/);
    assert(!existsSync(f.calls), "no compiler, build or test may start");
    assert.deepEqual(readdirSync(join(f.base, "tmp")), []);
  });
});

test("Dev module campaigns profile each backend call without changing script or app arguments", async t => {
  const f = fixture(t);
  const bin = join(f.base, "gopurs-alpha/bin");
  mkdirSync(bin, { recursive: true });
  writeFileSync(join(bin, "test"), `#!/usr/bin/env bash
set -e
test "$#" -eq 0
../gopurs/bin/gopurs --main 'Test.Main With Spaces'
../gopurs/bin/gopurs --main Other --output 'output with spaces'
`, { mode: 0o755 });
  mkdirSync(join(f.root, "runtime"));
  writeFileSync(join(f.root, "runtime/sentinel"), "compiler resource");
  const originalScript = readFileSync(join(bin, "test"), "utf8");
  for (const flag of ["-d", "--dev"]) {
    const result = await f.run(["alpha", flag, "-c", "--keep-workspace"], {}, "modtest");
    assert.equal(result.code, 0, result.output);
    const report = reportFor(result);
    assert.equal(report.buildProfile, "dev");
    assert.equal(readFileSync(join(report.targets[0].directory, "packages/gopurs/runtime/sentinel"), "utf8"), "compiler resource");
  }
  const release = await f.run(["alpha"], {}, "modtest");
  assert.equal(release.code, 0, release.output);
  assert.equal(reportFor(release).buildProfile, "release");
  const calls = readFileSync(f.calls, "utf8").trim().split("\n").map(JSON.parse);
  assert.deepEqual(calls.filter(call => call.phase === "gopurs").map(call => call.args), [
    ["--dev", "--main", "Test.Main With Spaces"], ["--dev", "--main", "Other", "--output", "output with spaces"],
    ["--dev", "--main", "Test.Main With Spaces"], ["--dev", "--main", "Other", "--output", "output with spaces"],
    ["--main", "Test.Main With Spaces"], ["--main", "Other", "--output", "output with spaces"],
  ]);
  assert.equal(calls.filter(call => call.phase === "npm").length, 2, "clean rebuilds only once per campaign");
  assert.equal(readFileSync(join(bin, "test"), "utf8"), originalScript);
  assert.equal(readFileSync(join(f.root, "bin/gopurs"), "utf8"), fakeTool);
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
    const report = reportFor(run);
    assert.equal(report.targets[0].status, "passed");
    assert(!existsSync(join(report.targets[0].directory, "output")));
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

test("package snapshots cover fragments, FFI, order and entry paths with strict replays", async t => {
  const f = fixture(t);
  f.source("Alpha", "-- @snapshot-ffi\nmodule Main where\n");
  const env = { GOPURS_GO_PACKAGES: "1" };
  const missing = await f.run(["Alpha"], env);
  assert.equal(missing.code, 1, missing.output);
  assert.match(missing.output, /Missing snapshot/);
  const snapshot = join(f.root, "tests/passing-package-snapshots/Alpha.txt");
  assert(!existsSync(snapshot));
  const update = await f.run(["Alpha", "--update-snapshots"], env);
  assert.equal(update.code, 0, update.output);
  const contents = readFileSync(snapshot, "utf8");
  for (const path of ["main/main.go", "purescript/init_all/init.go", "purescript/init_external/init.go",
    "purescript/p_root/Main.go", "purescript/p_root/Main_ffi.go", "purescript/p_root/zz_gopurs_order.go"]) {
    assert(contents.includes("// gopurs-package-file: " + path + "\n"), path);
  }
  const replay = await f.run(["Alpha"], env);
  assert.equal(replay.code, 0, replay.output);
  const moved = await f.run(["Alpha"], { ...env, RUNNER_PACKAGE_SPLIT: "1" });
  assert.equal(moved.code, 1, moved.output);
  assert.match(moved.output, /snapshot diff/);
  assert.equal(readFileSync(snapshot, "utf8"), contents);
  const failed = await f.run(["Alpha", "--update-snapshots"], { ...env, RUNNER_PACKAGE_SPLIT: "1", RUNNER_FAIL: "build" });
  assert.equal(failed.code, 1, failed.output);
  assert.equal(readFileSync(snapshot, "utf8"), contents);
  assert.equal(readFileSync(join(f.root, "tests/passing-snapshots/Alpha.go"), "utf8"), "package purescript\n");
});

test("coalesced snapshots select source sections without treating Go literals as markers", async t => {
  const f = fixture(t), directory = join(f.base, "projection"), group = "p_11111111111111111111111111111111";
  const output = join(directory, "output"), bundle = `purescript/${group}/0_gopurs_package.go`;
  for (const name of [`purescript/${group}`, "purescript/init_all", "purescript/init_external", "main"])
    mkdirSync(join(output, name), { recursive: true });
  mkdirSync(join(directory, "logs"));
  for (const [name, pkg] of [["main/main.go", "main"], ["purescript/init_all/init.go", "init_all"],
    ["purescript/init_external/init.go", "init_external"], [`purescript/${group}/zz_gopurs_order.go`, group]])
    writeFileSync(join(output, name), `package ${pkg}\n`);
  const section = (name, code) => `// gopurs-source-v1: ${Buffer.from(name).toString("hex")}\n${code}\n// gopurs-source-end-v1\n`;
  const code = `package ${group}\n\n` + section("Extra.go", "var Excluded = 17") +
    section("Main.go", 'var Literal = `\n// gopurs-source-v1: 66616b65\n// gopurs-source-end-v1\n`') +
    section("Main_ffi.go", "func Native() int { return Excluded }");
  writeFileSync(join(output, bundle), code);
  assert.deepEqual(packageSourceSections(code).sections.map(section => section.name), ["Extra.go", "Main.go", "Main_ffi.go"]);
  for (const snapshotFfi of [false, true]) {
    const selected = { directory, name: "Alpha", snapshotFfi }, env = { ...process.env, GOPURS_GO_PACKAGES: "1" };
    const files = snapshotFiles(f.root, selected, env);
    await formatSnapshots(files, { run: async (_, command, args, options) => {
      const result = spawnSync(command, args, { ...options, encoding: "utf8" });
      assert.equal(result.status, 0, result.stderr);
    } }, selected, env);
    const view = readFileSync(files[0].generated, "utf8");
    assert(view.includes("// gopurs-package-file: " + bundle));
    assert(view.includes("var Literal = `"));
    assert(!view.includes("var Excluded"));
    assert.equal(view.includes("func Native()"), snapshotFfi);
    assert.equal(readFileSync(join(output, bundle), "utf8").includes("var Excluded"), true);
  }
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
  assert.equal(reportFor(result).targets[0].status, "interrupted");
  assertSharedUntouched(f);
});

test("module campaigns stream both outputs before completion and retain logs without replay", async t => {
  const f = fixture(t);
  const bin = join(f.base, "gopurs-alpha/bin");
  mkdirSync(bin, { recursive: true });
  writeFileSync(join(bin, "test"), `#!/usr/bin/env node
import { existsSync } from "node:fs";
console.log("module-live-stdout");
console.error("module-live-stderr");
const timer = setInterval(() => {
  if (existsSync(process.env.RUNNER_RELEASE)) {
    clearInterval(timer);
    process.exitCode = Number(process.env.RUNNER_STATUS);
  }
}, 10);
`, { mode: 0o755 });
  for (const status of [0, 7]) {
    const release = join(f.base, `release-${status}`);
    const running = f.start(["alpha", "--dev"], { RUNNER_RELEASE: release, RUNNER_STATUS: String(status) }, "modtest");
    try {
      await waitFor(() => running.output.includes("module-live-stdout") && running.output.includes("module-live-stderr"));
      assert.equal(running.child.exitCode, null, "both messages must be visible while the module is still running");
      const inProgress = reportFor({ output: running.output });
      assert.equal(inProgress.targets[0].status, "running");
      const log = join(inProgress.targets[0].directory, "logs/test.log");
      for (const message of ["module-live-stdout", "module-live-stderr"]) assert.match(readFileSync(log, "utf8"), new RegExp(message));
      writeFileSync(release, "continue");
      const result = await running.completed;
      assert.equal(result.code, status === 0 ? 0 : 1, result.output);
      const report = reportFor(result);
      assert.equal(report.targets[0].status, status === 0 ? "passed" : "failed");
      for (const message of ["module-live-stdout", "module-live-stderr"]) {
        assert.equal(result.output.split(message).length - 1, 1, "streamed output must not be replayed");
        assert.equal(readFileSync(log, "utf8").split(message).length - 1, 1, "the log must retain one copy");
      }
    } finally {
      if (running.child.exitCode === null) running.child.kill("SIGTERM");
      await running.completed;
    }
  }
});

test("a closed terminal pipe stops the streamed command and preserves its log", async t => {
  const f = fixture(t);
  const log = join(f.base, "pipe.log");
  const resultFile = join(f.base, "pipe-result.json");
  const childFile = join(f.base, "pipe-child.json");
  const command = `import { writeFileSync } from "node:fs";
writeFileSync(${JSON.stringify(childFile)}, JSON.stringify({ pid: process.pid }));
console.log("module-before-pipe-error");
setInterval(() => {}, 1000);`;
  const driver = `import { writeFileSync } from "node:fs";
import { CommandRunner } from ${JSON.stringify(join(f.root, "tools/command-runner.mjs"))};
const commands = new CommandRunner();
try {
  await commands.run(process.execPath, ["--input-type=module", "-e", ${JSON.stringify(command)}], { log: ${JSON.stringify(log)}, stream: true });
  process.exitCode = 1;
} catch (error) {
  writeFileSync(${JSON.stringify(resultFile)}, JSON.stringify({ code: error.code, message: error.message }));
} finally { commands.dispose(); }`;
  const child = spawn(process.execPath, ["--input-type=module", "-e", driver]);
  let stderr = "";
  child.stderr.on("data", chunk => stderr += chunk);
  const completed = new Promise((resolve, reject) => {
    child.once("error", reject);
    child.once("close", code => resolve(code));
  });
  child.stdout.destroy();
  try {
    await waitFor(() => existsSync(resultFile));
    assert.equal(await completed, 0, stderr);
    const failure = JSON.parse(readFileSync(resultFile, "utf8"));
    assert.equal(failure.code, "EPIPE", failure.message);
    assert.match(readFileSync(log, "utf8"), /module-before-pipe-error/);
    const { pid } = JSON.parse(readFileSync(childFile, "utf8"));
    await waitFor(() => {
      try { process.kill(pid, 0); return false; } catch (error) { return error.code === "ESRCH"; }
    });
  } finally {
    if (child.exitCode === null) child.kill("SIGTERM");
    await completed;
  }
});

for (const [signal, code] of [["SIGINT", 130], ["SIGTERM", 143]]) test(`${signal} stops streamed module descendants and retains their log`, async t => {
  const f = fixture(t);
  const bin = join(f.base, "gopurs-alpha/bin");
  mkdirSync(bin, { recursive: true });
  writeFileSync(join(bin, "test"), `#!/usr/bin/env node
import { spawn } from "node:child_process";
import { writeFileSync } from "node:fs";
const child = spawn(process.execPath, ["-e", "setInterval(() => {}, 1000)"], { stdio: "inherit" });
writeFileSync(process.env.RUNNER_PAUSE, JSON.stringify({ pid: process.pid, child: child.pid }));
console.error("module-stream-ready");
setInterval(() => {}, 1000);
`, { mode: 0o755 });
  const pause = join(f.base, "pause.json");
  const running = f.start(["alpha", "--dev"], { RUNNER_PAUSE: pause }, "modtest");
  try {
    await waitFor(() => running.output.includes("module-stream-ready") && existsSync(pause));
    const pids = JSON.parse(readFileSync(pause, "utf8"));
    running.child.kill(signal);
    const result = await running.completed;
    assert.equal(result.code, code, result.output);
    for (const pid of [pids.pid, pids.child]) await waitFor(() => {
      try { process.kill(pid, 0); return false; } catch (error) { return error.code === "ESRCH"; }
    });
    const report = reportFor(result);
    assert.equal(report.targets[0].status, "interrupted");
    assert.equal(result.output.split("module-stream-ready").length - 1, 1);
    assert.match(readFileSync(join(report.targets[0].directory, "logs/test.log"), "utf8"), /module-stream-ready/);
    assertSharedUntouched(f);
  } finally {
    if (running.child.exitCode === null) running.child.kill("SIGTERM");
    await running.completed;
  }
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

function reportFor(result) {
  const match = result.output.match(/^Results: (.+)$/m);
  assert(match, result.output);
  return { file: match[1], ...JSON.parse(readFileSync(match[1], "utf8")) };
}

test("concurrent campaigns give every subprocess a private temporary directory", async t => {
  const f = fixture(t);
  const runs = await Promise.all([f.run(["Alpha", "--keep-workspace"]), f.run(["Beta", "--keep-workspace"])]);
  for (const run of runs) assert.equal(run.code, 0, run.output);
  const calls = readFileSync(f.calls, "utf8").trim().split("\n").map(JSON.parse);
  const spago = calls.filter(call => call.phase === "spago");
  assert.equal(spago.length, 2);
  assert.notEqual(spago[0].tmp, spago[1].tmp);
  for (const call of spago) {
    assert.equal(realpathSync(call.tmp), realpathSync(join(call.cwd, "tmp")));
    assert(calls.filter(other => other.cwd === call.cwd || other.cwd === join(call.cwd, "output")).every(other => other.tmp === call.tmp));
  }
});

test("complete reports retain failures and resume only unsuccessful targets", async t => {
  const f = fixture(t);
  f.source("Alpha", "module Main where\n-- fail=spago\n");
  const first = await f.run(["Alpha", "Beta", "--keep-going"]);
  assert.equal(first.code, 1, first.output);
  const report = reportFor(first);
  assert.equal(report.kind, "fixtures");
  assert.deepEqual(report.targets.map(row => row.status), ["failed", "passed"]);
  assert.match(report.targets[0].error, /purescript failed/);
  assert(existsSync(join(report.targets[0].directory, "logs/purescript.log")));
  assert(existsSync(join(report.targets[1].directory, "logs/execute.log")));
  assert(!existsSync(join(report.targets[1].directory, "output")));
  assert.match(first.output, /Summary: 1 passed, 1 failed, 0 pending/);
  f.source("Alpha", "module Main where\n");
  const resumed = await f.run(["--resume-failed", report.file]);
  assert.equal(resumed.code, 0, resumed.output);
  assert.doesNotMatch(resumed.output, /=> Testing Beta/);
  const retried = reportFor(resumed);
  assert.equal(retried.resumedFrom, report.file);
  assert.deepEqual(retried.targets.map(row => [row.target.split("/").at(-1), row.status]), [["Alpha.purs", "passed"]]);
  assert.deepEqual(JSON.parse(readFileSync(report.file, "utf8")).targets.map(row => row.status), ["failed", "passed"]);
  const nothing = await f.run(["--resume-failed", retried.file, "--list"]);
  assert.equal(nothing.code, 0, nothing.output);
  assert.equal(nothing.output, "");
});

test("fail-fast reports include unattempted targets and reject invalid resume reports read-only", async t => {
  const f = fixture(t);
  f.source("Alpha", "module Main where\n-- fail=build\n");
  const first = await f.run(["Alpha", "Beta"]);
  assert.equal(first.code, 1, first.output);
  const report = reportFor(first);
  assert.deepEqual(report.targets.map(row => row.status), ["failed", "pending"]);
  const listed = await f.run(["--resume-failed", report.file, "--list"]);
  assert.equal(listed.code, 0, listed.output);
  assert.match(listed.output, /Alpha.purs\n.*Beta.purs\n$/);
  const before = readFileSync(f.calls, "utf8");
  const invalid = join(f.base, "invalid.json");
  for (const data of ["{", JSON.stringify({ ...report, version: 99 }), JSON.stringify({ ...report, kind: "modules" }), JSON.stringify({ ...report, root: f.base }), JSON.stringify({ ...report, targets: [{ target: "Alpha", status: "unknown" }] })]) {
    writeFileSync(invalid, data);
    const result = await f.run(["--resume-failed", invalid]);
    assert.equal(result.code, 2, result.output);
  }
  for (const args of [["Alpha"], ["--all"], ["--skip-before", "Alpha"]]) {
    const result = await f.run(["--resume-failed", report.file, ...args]);
    assert.equal(result.code, 2, result.output);
  }
  assert.equal(readFileSync(f.calls, "utf8"), before);
});

test("module scripts run in fresh copies and cannot delete another checkout or failed output", async t => {
  const f = fixture(t);
  for (const name of ["alpha", "beta"]) {
    const module = join(f.base, "gopurs-" + name);
    mkdirSync(join(module, "bin"), { recursive: true });
    for (const cache of ["output", ".spago", ".cache"]) {
      mkdirSync(join(module, cache));
      writeFileSync(join(module, cache, "sentinel"), name);
    }
    writeFileSync(join(module, "spago.go.yaml"), "source config\n");
    symlinkSync("spago.go.yaml", join(module, "spago.yaml"));
    writeFileSync(join(module, "bin/test"), `#!/usr/bin/env bash
set -e
test ! -e output/sentinel
test -L spago.yaml
test -e ../gopurs/bin/gopurs
rm -rf ../gopurs-*/output ../gopurs-*/.spago ../gopurs-*/.cache
mkdir -p output
echo retained-${name} > output/evidence
echo module-${name}
${name === "alpha" ? "exit 7" : "exit 0"}
`, { mode: 0o755 });
  }
  const result = await f.run(["alpha", "beta", "--keep-going"], {}, "modtest");
  assert.equal(result.code, 1, result.output);
  const report = reportFor(result);
  assert.deepEqual(report.targets.map(row => row.status), ["failed", "passed"]);
  assert.match(result.output, /module-beta/);
  assert.equal(readFileSync(join(report.targets[0].directory, "packages/gopurs-alpha/output/evidence"), "utf8"), "retained-alpha\n");
  for (const name of ["alpha", "beta"]) for (const cache of ["output", ".spago", ".cache"]) {
    assert.equal(readFileSync(join(f.base, "gopurs-" + name, cache, "sentinel"), "utf8"), name);
  }
  const resume = await f.run(["--resume-failed", report.file, "--list"], {}, "modtest");
  assert.equal(resume.code, 0, resume.output);
  assert.equal(resume.output, "gopurs-alpha\n");
});

test("installed library runners limit direct cleanup to their own checkout", async t => {
  const parent = join(repository, "..");
  const libraries = readdirSync(parent).filter(name => name.startsWith("gopurs-") && existsSync(join(parent, name, "bin/test")));
  assert(libraries.length, "requires the local library checkouts used by the runners");
  for (const name of libraries) await t.test(name, () => {
    const base = mkdtempSync(join(tmpdir(), "gopurs-library-cleanup-"));
    try {
      mkdirSync(join(base, name, "bin"), { recursive: true });
      cpSync(join(parent, name, "bin/test"), join(base, name, "bin/test"));
      writeFileSync(join(base, name, "spago.go.yaml"), "package: {}\n");
      mkdirSync(join(base, "mockbin"));
      writeFileSync(join(base, "mockbin/spago"), '#!/usr/bin/env bash\necho reached-spago\nexit 23\n', { mode: 0o755 });
      for (const cache of ["output", ".spago", ".cache"]) {
        const directory = join(base, "gopurs-neighbor", cache);
        mkdirSync(directory, { recursive: true });
        writeFileSync(join(directory, "sentinel"), "neighbor data");
      }
      const result = spawnSync("bash", [join(base, name, "bin/test")], {
        env: { ...process.env, GOPURS_RUST: "0", GOPURS_JS: "0", PATH: join(base, "mockbin") + ":" + process.env.PATH },
        encoding: "utf8", timeout: 10_000,
      });
      assert.ifError(result.error);
      assert.equal(result.status, 1, result.stdout + result.stderr);
      assert.match(result.stdout, /reached-spago/);
      for (const cache of ["output", ".spago", ".cache"]) {
        assert.equal(readFileSync(join(base, "gopurs-neighbor", cache, "sentinel"), "utf8"), "neighbor data");
      }
    } finally {
      rmSync(base, { recursive: true, force: true });
    }
  });
});
