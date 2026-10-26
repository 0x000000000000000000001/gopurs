#!/usr/bin/env node
import { execFileSync } from "node:child_process";
import { chmodSync, copyFileSync, mkdirSync, mkdtempSync, readFileSync, rmSync, writeFileSync } from "node:fs";
import { tmpdir } from "node:os";
import { delimiter, join, resolve } from "node:path";

const options = {};
for (let index = 2; index < process.argv.length; index++) {
  const arg = process.argv[index];
  if (arg === "--help") {
    console.log("Usage: node tools/release-smoke.mjs --package FILE.tgz --native FILE [--purs FILE] [--report FILE] [--keep-workspace]\nInstalls only the tarball, then compiles and runs a fresh TAST/Go FFI application with its JS and native hosts.");
    process.exit(0);
  }
  if (arg === "--keep-workspace") options.keep = true;
  else if (["--package", "--native", "--purs", "--report"].includes(arg) && process.argv[index + 1]) {
    options[arg.slice(2)] = resolve(process.argv[++index]);
  } else throw new Error(`Unknown or incomplete option: ${arg}`);
}
if (!options.package || !options.native) throw new Error("--package and --native are required");
const compiler = options.purs ?? process.env.GOPURS_PURS ?? "purs";
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
let passed = false;
const results = [];
function run(command, args, cwd) {
  return execFileSync(command, args, { cwd, env: environment, encoding: "utf8", stdio: ["ignore", "pipe", "pipe"], maxBuffer: 8 * 1024 * 1024, timeout: 300_000 });
}
try {
  writeFileSync(join(workspace, "package.json"), JSON.stringify({ name: "gopurs-release-smoke", version: "1.0.0", private: true }));
  run("npm", ["install", "--offline", "--ignore-scripts", "--no-audit", "--no-fund", options.package], workspace);
  const installed = join(workspace, "node_modules/gopurs");
  const metadata = JSON.parse(readFileSync(join(installed, "package.json"), "utf8"));
  const native = join(installed, "bin/gopurs-native");
  copyFileSync(options.native, native);
  chmodSync(native, 0o755);
  environment.PATH = join(workspace, "node_modules/.bin") + delimiter + environment.PATH;
  for (const host of ["javascript", "native"]) {
    const application = join(workspace, host);
    const source = join(application, "src");
    mkdirSync(source, { recursive: true });
    // Minimal modules make this an offline compiler/FFI test, independent of
    // registry caches and sibling library checkouts. Library gates are separate.
    writeFileSync(join(source, "Effect.purs"), "module Effect where\nforeign import data Effect :: Type -> Type\n");
    writeFileSync(join(source, "Unit.purs"), "module Data.Unit where\ndata Unit = Unit\n");
    writeFileSync(join(source, "Main.purs"), 'module Main where\nimport Data.Unit (Unit)\nimport Effect (Effect)\nforeign import report :: String -> Int -> Effect Unit\nmain :: Effect Unit\nmain = report "gopurs release smoke" 42\n');
    writeFileSync(join(source, "Main.go"), 'package Main\nimport "fmt"\nfunc Report(message string, value int32, _ any) any { fmt.Printf("%s: %d\\n", message, value); return nil }\n');
    run(compiler, ["compile", "src/Effect.purs", "src/Unit.purs", "src/Main.purs", "--codegen", "corefn"], application);
    const tast = JSON.parse(readFileSync(join(application, "output/Main/corefn.json"), "utf8"));
    for (const field of ["typeTable", "dataDecls", "classDecls"]) {
      if (!Array.isArray(tast[field])) throw new Error(`The application compiler did not emit TAST ${field}`);
    }
    if (host === "javascript") run(join(workspace, "node_modules/.bin/gopurs"), ["--main", "Main"], application);
    else run(join(installed, "bin/gopurs"), ["--main", "Main"], application);
    const output = join(application, "output");
    const binary = join(application, "hello");
    run("go", ["build", "-trimpath", "-o", binary, "./main"], output);
    const actual = run(binary, [], application).trim();
    if (actual !== "gopurs release smoke: 42") throw new Error(`${host}: unexpected application output: ${JSON.stringify(actual)}`);
    results.push({ host, output: actual, passed: true });
    console.log(`${host}: installed gopurs ${metadata.version}, TAST → Go → FFI execution passed`);
  }
  if (options.report) writeFileSync(options.report, JSON.stringify({ version: metadata.version, npmInstall: "offline, empty npm cache, lifecycle scripts disabled", tests: results }, null, 2) + "\n");
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
