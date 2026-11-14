#!/usr/bin/env node
import assert from "node:assert/strict";
import { spawnSync } from "node:child_process";
import { createHash } from "node:crypto";
import {
  chmodSync,
  copyFileSync,
  existsSync,
  mkdirSync,
  mkdtempSync,
  readFileSync,
  readdirSync,
  rmSync,
  writeFileSync,
} from "node:fs";
import { tmpdir } from "node:os";
import { delimiter, join, resolve } from "node:path";
import { fileURLToPath } from "node:url";
import { findOfficialCompiler } from "./official-compiler.mjs";

const options = {};
for (let index = 2; index < process.argv.length; index++) {
  const arg = process.argv[index];
  if (arg === "--help") {
    console.log(
      "Usage: node tools/release/release-smoke.mjs --package FILE.tgz [--native FILE] [--hosts javascript,native] [--purs FILE] [--official-purs FILE] [--report FILE] [--keep-workspace]\nInstalls only the tarball, then compiles and runs a fresh TAST/Go FFI application with its JS and native hosts.",
    );
    process.exit(0);
  }
  if (arg === "--keep-workspace") options.keep = true;
  else if (arg === "--hosts" && process.argv[index + 1]) options.hosts = process.argv[++index].split(",");
  else if (
    ["--package", "--native", "--purs", "--official-purs", "--report"].includes(arg) &&
    process.argv[index + 1]
  ) {
    options[arg.slice(2)] = resolve(process.argv[++index]);
  } else throw new Error(`Unknown or incomplete option: ${arg}`);
}
const hosts = options.hosts ?? ["javascript", "native"];
if (!hosts.length || hosts.some((host) => !["javascript", "native"].includes(host)))
  throw new Error("--hosts accepts javascript,native");
if (!options.package || (hosts.includes("native") && !options.native))
  throw new Error("--package is required; native testing also requires --native");
const compiler = options.purs ?? process.env.GOPURS_PURS ?? "purs";
const officialCompiler = findOfficialCompiler(
  fileURLToPath(new URL("../../", import.meta.url)),
  options["official-purs"],
);
const workspace = mkdtempSync(join(tmpdir(), "gopurs-release-smoke-"));
const environment = { ...process.env, GOWORK: "off", GOTOOLCHAIN: "local" };
// Exercise the installed package without inheriting a developer's backend or
// cache selection. Its npm cache and generated application directories are new.
for (const key of Object.keys(environment)) {
  if (/^(GOPURS_|NODE_PATH$|NODE_OPTIONS$)/.test(key)) delete environment[key];
}
environment.npm_config_cache = join(workspace, "npm-cache");
environment.GOCACHE = join(workspace, "go-cache");
environment.GOMODCACHE = join(workspace, "go-mod-cache");
// Keep the two hosts on the same layout so the test driver can inspect the ABI.
environment.GOPURS_GO_PACKAGES = "0";
environment.GOPURS_PREPARE_JOBS = "2";
environment.GOPURS_PBO_JOBS = "1";
environment.GOPURS_EMIT_JOBS = "2";
let passed = false;
const results = [];
const commands = [];
mkdirSync(join(workspace, "logs"));
function run(command, args, cwd, extra = {}) {
  const started = performance.now();
  const result = spawnSync(command, args, {
    cwd,
    env: environment,
    encoding: "utf8",
    stdio: ["pipe", "pipe", "pipe"],
    maxBuffer: 8 * 1024 * 1024,
    timeout: 300_000,
    ...extra,
  });
  const log = "logs/" + commands.length + ".log";
  writeFileSync(join(workspace, log), (result.stdout ?? "") + (result.stderr ?? ""));
  commands.push({ command, args, cwd, log, status: result.status, seconds: (performance.now() - started) / 1000 });
  writeFileSync(join(workspace, "commands.json"), JSON.stringify(commands, null, 2) + "\n");
  if (result.error || result.status !== 0) {
    const failure = result.error ?? new Error(`${command} exited with status ${result.status}`);
    failure.stdout = result.stdout;
    failure.stderr = result.stderr;
    throw failure;
  }
  return result.stdout;
}
const hash = (bytes) => createHash("sha256").update(bytes).digest("hex");
function generatedHashes(directory, prefix = "") {
  const files = {};
  for (const entry of readdirSync(directory, { withFileTypes: true })) {
    if (entry.name.startsWith(".")) continue;
    const name = prefix + entry.name,
      path = join(directory, entry.name);
    if (entry.isDirectory()) Object.assign(files, generatedHashes(path, name + "/"));
    else if (name.endsWith(".go") && !name.endsWith("_test.go")) files[name] = hash(readFileSync(path));
  }
  return files;
}
try {
  writeFileSync(
    join(workspace, "package.json"),
    JSON.stringify({ name: "gopurs-release-smoke", version: "1.0.0", private: true }),
  );
  run("npm", ["install", "--offline", "--ignore-scripts", "--no-audit", "--no-fund", options.package], workspace);
  const installed = join(workspace, "node_modules/gopurs");
  const metadata = JSON.parse(readFileSync(join(installed, "package.json"), "utf8"));
  const native = join(installed, "bin/gopurs-native");
  if (hosts.includes("native")) {
    copyFileSync(options.native, native);
    chmodSync(native, 0o755);
  }
  for (const forbidden of ["src", "runtime", "output", "tests", "node_modules"])
    assert.equal(existsSync(join(installed, forbidden)), false, "tarball unexpectedly depends on " + forbidden);
  verifyPackagedHelpers(installed);
  environment.PATH = join(workspace, "node_modules/.bin") + delimiter + environment.PATH;
  for (const host of hosts) {
    const application = join(workspace, host);
    const source = join(application, "src");
    mkdirSync(source, { recursive: true });
    installApplication(source);
    writeFileSync(join(application, "package.json"), '{"type":"module"}\n');
    const sources = [
      "src/Effect.purs",
      "src/Unit.purs",
      "src/Uncurried.purs",
      "src/Exception.purs",
      "src/Compat.purs",
      "src/Main.purs",
    ];
    run(officialCompiler, ["compile", ...sources, "--codegen", "js,corefn", "-o", "official-output"], application);
    const officialTast = JSON.parse(readFileSync(join(application, "official-output/Main/corefn.json"), "utf8"));
    assert.equal(officialTast.typeTable, undefined, "--official-purs must select the standard compiler");
    const reference = run(
      "node",
      ["--input-type=module", "-e", 'import { main } from "./official-output/Main/index.js"; main();'],
      application,
    ).trim();
    assert.equal(reference, "gopurs release smoke: 42");
    const oracle = run(
      "node",
      [
        "--input-type=module",
        "-e",
        `
      import { probe, calls } from "./src/Main.js";
      const run = action => new Promise((resolve, reject) => action(reject, resolve));
      const action = probe(false);
      if (calls() !== 0) throw new Error("eager action");
      console.log(await run(action)); console.log(await run(action)); console.log(calls());
      try { await run(probe(true)); } catch (error) { console.log(error.message); }
    `,
      ],
      application,
    ).trim();
    assert.equal(oracle, "packaged result\npackaged result\n2\npackaged error");
    run(compiler, ["compile", ...sources, "--codegen", "corefn"], application);
    const tast = JSON.parse(readFileSync(join(application, "output/Main/corefn.json"), "utf8"));
    for (const field of ["typeTable", "dataDecls", "classDecls"])
      assert.ok(Array.isArray(tast[field]), "missing TAST " + field);
    const driver = host === "javascript" ? join(workspace, "node_modules/.bin/gopurs") : join(installed, "bin/gopurs");
    const output = join(application, "output");
    const execute = (label, expectedResult, expectsError) => {
      const binary = join(application, "hello");
      run("go", ["build", "-trimpath", "-o", binary, "./main"], output);
      assert.equal(run(binary, [], application).trim(), reference);
      writeFileSync(join(output, "purescript/ffi_package_test.go"), bridgeTest(expectedResult, expectsError));
      const tested = run("go", ["test", "-race", "-count=1", "-timeout=20s", "./purescript"], output);
      writeFileSync(join(workspace, "logs", host + "-" + label + ".log"), tested);
    };
    run(driver, ["--main", "Main"], application);
    assert.match(
      readFileSync(join(workspace, commands.at(-1).log), "utf8"),
      /cache miss/,
      "fresh application unexpectedly reused a cache",
    );
    execute("cold", "packaged result", true);
    const cold = generatedHashes(output);
    run(driver, ["--main", "Main"], application);
    assert.match(
      readFileSync(join(workspace, commands.at(-1).log), "utf8"),
      /cache hit/,
      "unchanged application did not reuse its cache",
    );
    assert.deepEqual(generatedHashes(output), cold, "warm cache changed emitted Go");
    execute("warm", "packaged result", true);
    writeFileSync(join(source, "Main.go"), ordinarySource("changed"));
    run(driver, ["--main", "Main"], application);
    assert.notDeepEqual(generatedHashes(output), cold, "signature edit reused stale Go");
    execute("signature-change", "changed result", false);
    const committed = readFileSync(join(output, ".gopurs-cache/current.json"), "utf8");
    writeFileSync(join(source, "Main.go"), ordinarySource("invalid"));
    let diagnostic = "";
    try {
      run(driver, ["--main", "Main"], application);
    } catch (error) {
      diagnostic = String(error.stdout ?? "") + String(error.stderr ?? "");
    }
    assert.match(diagnostic, /Main/, "incompatible signature did not identify its module");
    assert.match(diagnostic, /Main\.go/, "incompatible signature did not identify the handwritten file");
    assert.match(diagnostic, /results|result|signature/, "missing signature diagnostic");
    assert.equal(
      readFileSync(join(output, ".gopurs-cache/current.json"), "utf8"),
      committed,
      "failed build published a cache entry",
    );
    writeFileSync(join(workspace, "logs", host + "-invalid-signature.log"), diagnostic);
    writeFileSync(join(source, "Main.go"), ordinarySource("original"));
    run(driver, ["--main", "Main"], application);
    execute("recovery", "packaged result", true);
    assert.deepEqual(generatedHashes(output), cold, "recovery did not restore the original emitted Go");
    results.push({
      host,
      output: reference,
      oracle,
      stages: ["official-javascript", "cold", "warm", "signature-change", "invalid-signature", "recovery"],
      passed: true,
    });
    console.log(
      `${host}: installed gopurs ${metadata.version}, migration, idiomatic effects, Aff and cache transitions passed`,
    );
  }
  if (options.report)
    writeFileSync(
      options.report,
      JSON.stringify(
        {
          version: metadata.version,
          npmInstall: "offline, empty npm cache, lifecycle scripts disabled",
          tests: results,
          node: process.version,
          go: run("go", ["version"], workspace).trim(),
          typedCompiler: { path: compiler, version: run(compiler, ["--version"], workspace).trim() },
          officialCompiler: { path: officialCompiler, version: run(officialCompiler, ["--version"], workspace).trim() },
          packageSHA256: hash(readFileSync(options.package)),
          commands,
        },
        null,
        2,
      ) + "\n",
    );
  passed = true;
} catch (error) {
  process.stderr.write(error.stdout ?? "");
  process.stderr.write(error.stderr ?? "");
  console.error(`Release smoke failed: ${error.message}\nWorkspace retained: ${workspace}`);
  process.exitCode = 1;
} finally {
  if (passed && !options.keep) rmSync(workspace, { recursive: true, force: true });
  else if (passed) console.log(`Workspace retained: ${workspace}`);
}

function verifyPackagedHelpers(installed) {
  const source =
    'package fixture\nimport ctx "context"\ntype Failure = error\nfunc Probe(context ctx.Context) (value string, err Failure) { return "", nil }\n';
  const wasm = JSON.parse(run("node", [join(installed, "tools/ffi-runner.mjs")], workspace, { input: source }));
  const probe = wasm.find((declaration) => declaration.name === "Probe");
  assert.equal(probe.schemaVersion, 2);
  assert.equal(probe.args[0].identity, "context.Context");
  assert.equal(probe.results.length, 2);
  assert.equal(probe.results[1].identity, "builtin.error");
  const helper = join(installed, "tools/ffi-gen");
  writeFileSync(join(helper, "schema-input.txt"), source);
  writeFileSync(
    join(helper, "schema_smoke_test.go"),
    `package main
import ("os"; "testing")
func TestPackagedSchema(t *testing.T) {
  source, err := os.ReadFile("schema-input.txt"); if err != nil { t.Fatal(err) }
  result, err := Extract(string(source)); if err != nil { t.Fatal(err) }
  if err := os.WriteFile("native-schema.json", []byte(result), 0600); err != nil { t.Fatal(err) }
}
`,
  );
  run("go", ["test", "-count=1", "."], helper);
  assert.deepEqual(
    JSON.parse(readFileSync(join(helper, "native-schema.json"), "utf8")),
    wasm,
    "packaged native and WASM helpers disagree on the versioned signature contract",
  );
  console.log("installed helpers: native/WASM schema v2 parity passed");
}

function ordinarySource(variant) {
  const probe =
    variant === "original"
      ? 'func Probe(fail bool) (string, error) { calls.Add(1); if fail { return "", errors.New("packaged error") }; return "packaged result", nil }'
      : variant === "changed"
        ? 'func Probe(fail bool) string { calls.Add(1); return "changed result" }'
        : 'func Probe(fail bool) (string, int32, error) { return "invalid", 0, nil }';
  return `package Main
import ("fmt"; "sync/atomic"${variant === "original" ? '; "errors"' : ""})
var calls atomic.Int32
func Pure(value int32) int32 { return value + 1 }
func Report(message string, value int32) { fmt.Printf("%s: %d\\n", message, value) }
func Calls() int32 { return calls.Load() }
${probe}
`;
}

function installApplication(source) {
  const files = {
    "Effect.purs": "module Effect where\nforeign import data Effect :: Type -> Type\n",
    "Unit.purs": "module Data.Unit where\ndata Unit = Unit\n",
    "Uncurried.purs":
      "module Effect.Uncurried where\nforeign import data EffectFn1 :: Type -> Type -> Type\nforeign import data EffectFn2 :: Type -> Type -> Type -> Type\nforeign import data EffectFn3 :: Type -> Type -> Type -> Type -> Type\n",
    "Exception.purs": "module Effect.Exception where\nforeign import data Error :: Type\n",
    "Compat.purs": `module Effect.Aff.Compat where
import Data.Unit (Unit)
import Effect.Exception (Error)
import Effect.Uncurried (EffectFn1, EffectFn2, EffectFn3)
type EffectFnCb a = EffectFn1 a Unit
newtype EffectFnAff a = EffectFnAff (EffectFn2 (EffectFnCb Error) (EffectFnCb a) EffectFnCanceler)
newtype EffectFnCanceler = EffectFnCanceler (EffectFn3 Error (EffectFnCb Error) (EffectFnCb Unit) Unit)
`,
    "Main.purs": `module Main where
import Data.Unit (Unit)
import Effect (Effect)
import Effect.Aff.Compat (EffectFnAff)
foreign import pure :: Int -> Int
foreign import report :: String -> Int -> Effect Unit
foreign import probe :: Boolean -> EffectFnAff String
foreign import calls :: Effect Int
main :: Effect Unit
main = report "gopurs release smoke" (pure 41)
`,
    "Main.go": ordinarySource("original"),
    "Main.js": `let count = 0;
export const pure = value => value + 1;
export const report = message => value => () => console.log(message + ": " + value);
export const calls = () => count;
export const probe = fail => (onError, onSuccess) => {
  count++;
  if (fail) onError(new Error("packaged error")); else onSuccess("packaged result");
  return (_reason, _onError, onSuccess) => onSuccess();
};
`,
  };
  assert.doesNotMatch(files["Main.go"], /gopurs_runtime/);
  for (const [name, content] of Object.entries(files)) writeFileSync(join(source, name), content);
}

function bridgeTest(expectedResult, expectsError) {
  return `package purescript
import ("testing"; "time"; rt "gopurs/output/gopurs_runtime")
func TestPackagedIdiomaticAff(t *testing.T) {
  calls := func() int64 { return rt.Apply(_Gopurs_Main_Calls, rt.Value{}).IntVal }
  action := rt.Apply(_Gopurs_Main_Probe, rt.Bool(false))
  if calls() != 0 { t.Fatal("action executed before registration") }
  type outcome struct { value rt.Value; failed bool }
  await := func(action rt.Value) outcome {
    done := make(chan outcome, 1)
    rt.Apply2(action, rt.Func(func(value rt.Value) rt.Value { done <- outcome{value, true}; return rt.Value{} }),
      rt.Func(func(value rt.Value) rt.Value { done <- outcome{value, false}; return rt.Value{} }))
    select { case result := <-done: return result; case <-time.After(time.Second): t.Fatal("action did not settle"); return outcome{} }
  }
  for range 2 { result := await(action); if result.failed || result.value.StrVal() != ${JSON.stringify(expectedResult)} { t.Fatal(result) } }
  if calls() != 2 { t.Fatal("action was memoized") }
  ${expectsError ? 'failure := await(rt.Apply(_Gopurs_Main_Probe, rt.Bool(true))); if !failure.failed || failure.value.AnyVal().(error).Error() != "packaged error" { t.Fatal("error identity was lost") }' : ""}
}
`;
}
