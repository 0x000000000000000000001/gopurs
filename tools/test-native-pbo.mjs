#!/usr/bin/env node
// Native-PBO differential qualification for the Rust-hosted gopurs compiler.
//
// Runs eight native FFI differential contracts against a retained Rust-hosted
// gopurs workspace (the rust/ directory produced by `npm run build:rust` with
// GOPURS_KEEP_WORKSPACE=1 or `-- --keep-workspace`):
//
//   memo, directives, qualified, source-usage, tast, tast-module, type-table, maps
//
// Candidate native sources are injected from the local gopurs PBO checkout
// (GOPURS_PBO_DIR, default ../../purescript-backend-optimizer-gopurs);
// PureScript oracles and the embedded FFI come from the generated crates in
// GENERATED_RUST. Case scripts live in tools/native-pbo/; fixtures come from
// PURUST_DIR/tools, except the qualified fixture which ships next to its case
// script. All builds share one cargo target
// directory and force profile.release {opt-level=3, debug=false, lto=false}.
//
// Usage:
//   node tools/test-native-pbo.mjs GENERATED_RUST NEW_LOG_DIR [TAST_CORPUS] [options]
//
// Options:
//   --target-dir DIR   Shared cargo target dir (default <GENERATED_RUST>/../target)
//   --only a,b         Run only the named cases
//   --list             Print the case list and exit
//   --keep             Retain successful fixture workspaces (PBO_NATIVE_KEEP=1)
//   --require-corpus   Fail instead of skipping when no TAST corpus is available
//   --help, -h
//
// The TAST corpus defaults to <GENERATED_RUST>/../output, the typed output of
// the retained build. Four cases need it; memo, directives and qualified run
// without it.
// Failed fixture workspaces and all logs stay under NEW_LOG_DIR.
import assert from "node:assert/strict";
import { existsSync, mkdirSync, readFileSync, readdirSync, statSync, writeFileSync } from "node:fs";
import { dirname, join, resolve } from "node:path";
import { spawnSync } from "node:child_process";
import { createHash } from "node:crypto";
import {
  GOPURS_FORK, PURUST_FORK, gopursRoot, isDirectory, pboDir,
  purustDir, purustTools, requireDirectory, loadThreadedRust,
} from "./native-pbo/shared.mjs";

const CASES = [
  { name: "memo", script: "memo.mjs", corpus: false, injects: [] },
  { name: "directives", script: "directives.mjs", corpus: false, injects: [] },
  { name: "source-usage", script: "source-usage.mjs", corpus: true, injects: ["PureScript/Backend/Optimizer/CoreFn/Usage.rs"] },
  { name: "tast", script: "tast.mjs", corpus: true, injects: ["PureScript/Backend/Optimizer/CoreFn/Json.rs"] },
  { name: "tast-module", script: "tast-module.mjs", corpus: true, injects: ["PureScript/Backend/Optimizer/CoreFn/Json.rs"] },
  { name: "type-table", script: "type-table.mjs", corpus: true, injects: ["PureScript/Backend/Optimizer/CoreFn/Json.rs"] },
  { name: "maps", script: "maps.mjs", corpus: false, injects: ["PureScript/Backend/Optimizer/NativeMaps.rs"] },
  { name: "qualified", script: "qualified.mjs", corpus: false, injects: ["PureScript/Backend/Optimizer/CoreFn.rs"] },
];

function usage() {
  return [
    "Usage: node tools/test-native-pbo.mjs GENERATED_RUST NEW_LOG_DIR [TAST_CORPUS] [options]",
    "",
    "  GENERATED_RUST   retained rust/ output of `npm run build:rust` (GOPURS_KEEP_WORKSPACE=1)",
    "  NEW_LOG_DIR      campaign log directory; failed workspaces stay in NEW_LOG_DIR/workspaces",
    "  TAST_CORPUS      typed output directory (default <GENERATED_RUST>/../output)",
    "",
    "Options:",
    "  --target-dir DIR   shared cargo target dir (default <GENERATED_RUST>/../target)",
    "  --only a,b         run only the named cases",
    "  --list             print the case list and exit",
    "  --keep             retain successful fixture workspaces",
    "  --require-corpus   fail instead of skipping when no TAST corpus is available",
  ].join("\n");
}

function parse(argv) {
  const options = { targetDir: null, only: null, keep: false, requireCorpus: false, list: false, help: false };
  const positionals = [];
  for (let index = 0; index < argv.length; index += 1) {
    const argument = argv[index];
    if (argument === "--help" || argument === "-h") options.help = true;
    else if (argument === "--list") options.list = true;
    else if (argument === "--keep") options.keep = true;
    else if (argument === "--require-corpus") options.requireCorpus = true;
    else if (argument === "--target-dir") {
      options.targetDir = argv[index + 1];
      assert.ok(options.targetDir, "--target-dir needs a directory");
      index += 1;
    } else if (argument.startsWith("--target-dir=")) options.targetDir = argument.slice("--target-dir=".length);
    else if (argument === "--only") {
      options.only = argv[index + 1];
      assert.ok(options.only, "--only needs a case list");
      index += 1;
    } else if (argument.startsWith("--only=")) options.only = argument.slice("--only=".length);
    else if (argument.startsWith("-")) throw new Error(`Unknown option: ${argument}\n\n${usage()}`);
    else positionals.push(argument);
  }
  return { options, positionals };
}

// Read modulePath values from either a typed output directory of
// <Module>/corefn.json trees or the packed [{ name, contents }] corpus.
function corpusModulePaths(corpus) {
  const stat = statSync(corpus, { throwIfNoEntry: false });
  if (!stat) return [];
  if (stat.isFile()) {
    const paths = [];
    let entries;
    try {
      entries = JSON.parse(readFileSync(corpus, "utf8"));
    } catch (error) {
      throw new Error(`TAST corpus ${corpus} is neither a directory nor a packed [{ name, contents }] JSON file`);
    }
    if (!Array.isArray(entries)) {
      throw new Error(`TAST corpus ${corpus} is not a packed [{ name, contents }] array`);
    }
    for (const entry of entries) {
      try {
        const value = JSON.parse(entry.contents);
        if (typeof value.modulePath === "string") paths.push(value.modulePath);
      } catch (error) { /* not a module entry */ }
    }
    return paths;
  }
  const paths = [];
  for (const entry of readdirSync(corpus, { withFileTypes: true })) {
    if (!entry.isDirectory()) continue;
    const file = join(corpus, entry.name, "corefn.json");
    if (!existsSync(file)) continue;
    try {
      const value = JSON.parse(readFileSync(file, "utf8"));
      if (typeof value.modulePath === "string") paths.push(value.modulePath);
    } catch (error) { /* ignore malformed ignored members */ }
  }
  return paths;
}

// The oracles must be generated from the gopurs PBO fork, not the Purust fork.
function verifyProvenance(corpus) {
  const paths = corpusModulePaths(corpus);
  const gopurs = paths.filter(path => path.includes(GOPURS_FORK)).length;
  const purust = paths.filter(path => path.includes(PURUST_FORK)).length;
  if (paths.length === 0) return `unverified (no modulePath found in ${corpus})`;
  if (gopurs === 0 && purust > 0) {
    throw new Error(
      `TAST corpus ${corpus} names ${purust} Purust module paths and no ${GOPURS_FORK} path; ` +
      "pass the typed output of the retained gopurs build instead");
  }
  if (gopurs === 0) return `unverified (${paths.length} module paths, none naming ${GOPURS_FORK})`;
  return `gopurs (${gopurs}/${paths.length} module paths)`;
}

async function main() {
  const { options, positionals } = parse(process.argv.slice(2));
  if (options.help) { console.log(usage()); return; }
  if (options.list) {
    for (const item of CASES) console.log(`${item.name}${item.corpus ? " (TAST_CORPUS)" : ""}`);
    return;
  }

  const [rustArg, logArg, corpusArg] = positionals;
  assert.ok(rustArg && logArg, usage());
  const rust = resolve(rustArg);
  const logDir = resolve(logArg);
  requireDirectory(rust, "GENERATED_RUST workspace");
  assert.ok(!existsSync(logDir), `Use a new log directory: ${logDir}`);
  mkdirSync(logDir, { recursive: true });

  let corpus = null;
  if (corpusArg) {
    corpus = resolve(corpusArg);
    assert.ok(existsSync(corpus), `Missing TAST corpus: ${corpus}`);
  } else if (isDirectory(join(dirname(rust), "output"))) {
    corpus = join(dirname(rust), "output");
  }

  let selected = CASES;
  if (options.only) {
    const names = options.only.split(",").map(name => name.trim()).filter(Boolean);
    const unknown = names.filter(name => !CASES.some(item => item.name === name));
    assert.equal(unknown.length, 0, `Unknown case(s): ${unknown.join(", ")}`);
    selected = CASES.filter(item => names.includes(item.name));
  }

  const targetDir = resolve(options.targetDir ?? process.env.PBO_NATIVE_TARGET_DIR ?? join(dirname(rust), "target"));
  const workspaceParent = resolve(process.env.PBO_NATIVE_TMPDIR ?? join(logDir, "workspaces"));
  mkdirSync(workspaceParent, { recursive: true });

  requireDirectory(purustTools, "Purust tools fixture directory");
  assert.ok(existsSync(join(purustDir, "src/Purust/Threading.js")), `Missing Purust threading helper under ${purustDir}`);
  for (const item of selected) {
    assert.ok(existsSync(join(gopursRoot, "tools", "native-pbo", item.script)), `Missing case script: ${item.script}`);
    for (const relative of item.injects) {
      assert.ok(existsSync(join(pboDir, "src", relative)), `Missing gopurs PBO source ${relative} under ${pboDir}`);
    }
  }

  // Differential injection must test the very FFI embedded in the compiler.
  // Check full contents, not just a function-name/staleness marker.
  const threadedRust = await loadThreadedRust();
  const ffiModules = new Set(selected.flatMap(item => item.injects.length ? item.injects
    : [`PureScript/Backend/Optimizer/${item.name === "memo" ? "BoundedMemo" : "Directives"}.rs`]));
  const nativeSources = [];
  for (const relative of ffiModules) {
    const source = join(pboDir, "src", relative);
    const crate = "Purs_" + relative.replace(/\.rs$/, "").replaceAll("/", "_");
    const generated = join(rust, crate, "src/lib.rs");
    const bytes = readFileSync(source);
    assert.ok(readFileSync(generated, "utf8").includes(threadedRust(bytes.toString()).trim()),
      `${source} differs from the FFI embedded in ${generated}`);
    nativeSources.push({ source, generated, sha256: createHash("sha256").update(bytes).digest("hex") });
  }

  let provenance = "not available (no corpus)";
  let provenanceError = null;
  if (corpus) {
    try {
      provenance = verifyProvenance(corpus);
    } catch (error) {
      provenance = `failed: ${error.message}`;
      provenanceError = error;
    }
  }
  const environment = {
    startedAt: new Date().toISOString(),
    node: process.version,
    platform: process.platform,
    generatedRust: rust,
    logDir,
    corpus,
    targetDir,
    workspaceParent,
    gopursPboDir: pboDir,
    purustDir,
    purustFixtures: purustTools,
    provenance,
    nativeSources,
  };
  writeFileSync(join(logDir, "environment.json"), JSON.stringify(environment, null, 2) + "\n");
  console.log(JSON.stringify(environment, null, 2));

  if (provenanceError) {
    const summary = {
      ...environment,
      finishedAt: new Date().toISOString(),
      status: "failed",
      failed: 1,
      skipped: 0,
      results: [{ name: "provenance", status: "failed", reason: provenanceError.message }],
    };
    writeFileSync(join(logDir, "summary.json"), JSON.stringify(summary, null, 2) + "\n");
    console.error(`[native-pbo] FAILED: ${provenanceError.message}`);
    process.exitCode = 1;
    return;
  }

  const env = { ...process.env,
    GOPURS_PBO_DIR: pboDir,
    PURUST_DIR: purustDir,
    PBO_NATIVE_LOG_DIR: logDir,
    PBO_NATIVE_TARGET_DIR: targetDir,
    PBO_NATIVE_TMPDIR: workspaceParent,
  };
  if (options.keep) env.PBO_NATIVE_KEEP = "1";

  const results = [];
  for (const item of selected) {
    if (item.corpus && !corpus) {
      const reason = "no TAST corpus available (pass TAST_CORPUS or use a retained workspace with output/)";
      const status = options.requireCorpus ? "failed" : "skipped";
      results.push({ name: item.name, status, reason });
      console.error(`[native-pbo] ${item.name}: ${status.toUpperCase()} (${reason})`);
      continue;
    }
    const args = [join(gopursRoot, "tools", "native-pbo", item.script), rust];
    if (item.corpus) args.push(corpus);
    const started = Date.now();
    const child = spawnSync(process.execPath, args, { cwd: gopursRoot, env, encoding: "utf8", maxBuffer: 64 * 1024 * 1024 });
    const seconds = Number(((Date.now() - started) / 1000).toFixed(1));
    const consoleLog = join(logDir, `${item.name}.console.log`);
    writeFileSync(consoleLog, `$ node ${args.join(" ")}\n\n${child.stdout ?? ""}${child.stderr ?? ""}`);
    const result = {
      name: item.name,
      status: child.status === 0 ? "ok" : "failed",
      exitCode: child.status,
      seconds,
      log: consoleLog,
      caseLog: join(logDir, `${item.name}.log`),
      injects: item.injects.map(relative => join(pboDir, "src", relative)),
    };
    if (child.error) result.error = child.error.message;
    if (child.signal) result.signal = child.signal;
    results.push(result);
    console.log(`[native-pbo] ${item.name}: ${result.status.toUpperCase()} (${seconds}s) -> ${consoleLog}`);
  }

  const failed = results.filter(result => result.status === "failed").length;
  const skipped = results.filter(result => result.status === "skipped").length;
  const summary = {
    ...environment,
    finishedAt: new Date().toISOString(),
    status: failed ? "failed" : skipped ? "partial" : "ok",
    failed,
    skipped,
    results,
  };
  writeFileSync(join(logDir, "summary.json"), JSON.stringify(summary, null, 2) + "\n");
  console.log(`[native-pbo] ${summary.status.toUpperCase()}: ` +
    `${results.length - failed - skipped} ok / ${failed} failed / ${skipped} skipped`);
  if (failed) process.exitCode = 1;
}

try {
  await main();
} catch (error) {
  console.error(`[native-pbo] ${error?.message ?? error}`);
  if (error?.stack) console.error(error.stack);
  process.exitCode = 1;
}
