// Shared plumbing for the native-PBO differential harness in
// tools/native-pbo. Candidate native sources are injected from the local
// gopurs PBO checkout (GOPURS_PBO_DIR); PureScript oracles and the embedded
// FFI under test come from the retained Rust-hosted gopurs workspace
// (GENERATED_RUST). Fixture programs are reused from the Purust tools checkout
// (PURUST_DIR/tools; override with PBO_NATIVE_FIXTURES).
//
// Environment:
//   GOPURS_PBO_DIR        PBO checkout (default ../../purescript-backend-optimizer-gopurs)
//   PURUST_DIR            Purust checkout (default ../../purust/purust)
//   PBO_NATIVE_FIXTURES   fixture directory (default PURUST_DIR/tools)
//   PBO_NATIVE_LOG_DIR    per-case log directory (set by the orchestrator)
//   PBO_NATIVE_TARGET_DIR shared cargo target dir (default <GENERATED_RUST>/../target)
//   PBO_NATIVE_TMPDIR     fixture workspace parent (default <log>/workspaces or os tmpdir)
//   PBO_NATIVE_KEEP=1     retain successful fixture workspaces
import assert from "node:assert/strict";
import { appendFileSync, existsSync, mkdirSync, mkdtempSync, readFileSync, rmSync, statSync, writeFileSync } from "node:fs";
import { tmpdir } from "node:os";
import { dirname, join, resolve } from "node:path";
import { fileURLToPath, pathToFileURL } from "node:url";
import { spawnSync } from "node:child_process";

export const gopursRoot = fileURLToPath(new URL("../../", import.meta.url));
export const purustDir = resolve(process.env.PURUST_DIR ?? join(gopursRoot, "../../purust/purust"));
export const purustTools = resolve(process.env.PBO_NATIVE_FIXTURES ?? join(purustDir, "tools"));
export const pboDir = resolve(process.env.GOPURS_PBO_DIR ?? join(gopursRoot, "../../purescript-backend-optimizer-gopurs"));

export const PURUST_FORK = "purescript-backend-optimizer-purust";
export const GOPURS_FORK = "purescript-backend-optimizer-gopurs";

export function isDirectory(path) {
  return statSync(path, { throwIfNoEntry: false })?.isDirectory() === true;
}

export function requireDirectory(path, label) {
  assert.ok(isDirectory(path), `Missing ${label}: ${path}`);
}

// Reuse the Purust fixture programs (.rs) verbatim. Only the injected native
// source and the generated crates are specific to this checkout.
export function fixturePath(name) {
  const path = join(purustTools, name);
  assert.ok(existsSync(path), `Missing reused Purust fixture ${name} under ${purustTools} (set PBO_NATIVE_FIXTURES)`);
  return path;
}

export function fixtureSource(name) {
  return readFileSync(fixturePath(name), "utf8");
}

export function pboSource(relative) {
  const path = join(pboDir, "src", relative);
  assert.ok(existsSync(path), `Missing gopurs PBO source ${relative} under ${pboDir} (set GOPURS_PBO_DIR)`);
  return readFileSync(path, "utf8");
}

export async function loadThreadedRust() {
  const module = join(purustDir, "src/Purust/Threading.js");
  assert.ok(existsSync(module), `Missing Purust threading helper: ${module}`);
  const { threadedRust } = await import(pathToFileURL(module).href);
  return threadedRust;
}

export function targetDirFor(rust) {
  return resolve(process.env.PBO_NATIVE_TARGET_DIR ?? join(dirname(rust), "target"));
}

export function defaultCorpus(rust) {
  const candidate = join(dirname(rust), "output");
  return isDirectory(candidate) ? candidate : null;
}

export function resolveCaseArgs({ corpus = false } = {}) {
  const rustArg = process.argv[2];
  assert.ok(rustArg, `Usage: node native-pbo/<case>.mjs GENERATED_RUST${corpus ? " [TAST_CORPUS]" : ""}`);
  const rust = resolve(rustArg);
  requireDirectory(rust, "GENERATED_RUST workspace");
  if (!corpus) return { rust, corpus: null };
  const corpusArg = process.argv[3] ? resolve(process.argv[3]) : defaultCorpus(rust);
  assert.ok(corpusArg, `No TAST corpus given and no ${join(dirname(rust), "output")} directory exists`);
  assert.ok(existsSync(corpusArg), `Missing TAST corpus: ${corpusArg}`);
  return { rust, corpus: corpusArg };
}

export function createWorkspace(prefix) {
  const parent = resolve(process.env.PBO_NATIVE_TMPDIR ?? process.env.GOPURS_NATIVE_TMPDIR ?? tmpdir());
  mkdirSync(parent, { recursive: true });
  const directory = mkdtempSync(join(parent, `${prefix}-`));
  mkdirSync(join(directory, "src"));
  return directory;
}

export function writeCrate(directory, packageName, modules, rust) {
  const manifest =
    `[package]\nname = ${JSON.stringify(packageName)}\nversion = "0.0.0"\nedition = "2021"\n\n` +
    `[profile.release]\nopt-level = 3\ndebug = false\nlto = false\n\n[dependencies]\n` +
    modules.map(name => `${name} = { path = ${JSON.stringify(join(rust, name))} }\n`).join("");
  writeFileSync(join(directory, "Cargo.toml"), manifest);
}

export function generatedCrateSource(rust, module) {
  const path = join(rust, module, "src/lib.rs");
  assert.ok(existsSync(path),
    `GENERATED_RUST lacks ${module}/src/lib.rs; rebuild the Rust-hosted gopurs with a retained workspace first`);
  return readFileSync(path, "utf8");
}

// One case = one differential contract. Logs are appended as they are emitted
// so an interrupted campaign still keeps partial evidence; a failed fixture
// workspace is retained next to the logs, a successful one is removed unless
// PBO_NATIVE_KEEP=1.
export async function runCase(name, rust, body) {
  const logDir = process.env.PBO_NATIVE_LOG_DIR ? resolve(process.env.PBO_NATIVE_LOG_DIR) : null;
  if (logDir) mkdirSync(logDir, { recursive: true });
  const logFile = logDir ? join(logDir, `${name}.log`) : null;
  let started = false;
  const emit = text => {
    const line = String(text).replace(/\n?$/, "\n");
    if (logFile) {
      if (started) appendFileSync(logFile, line);
      else { writeFileSync(logFile, line); started = true; }
    }
    process.stdout.write(line);
  };

  const state = { workspace: null };
  const ctx = {
    name,
    state,
    log: emit,
    note: message => emit(`[${name}] ${message}`),
    spawn: (stage, command, args, options = {}) => {
      emit(`[${name}] ${stage}: ${command} ${args.join(" ")}`);
      const result = spawnSync(command, args, {
        encoding: "utf8",
        timeout: options.timeout ?? 3_600_000,
        maxBuffer: options.maxBuffer ?? 64 * 1024 * 1024,
      });
      if (result.stdout) emit(result.stdout);
      if (result.stderr) emit(result.stderr);
      assert.equal(result.status, 0,
        `[${name}] ${stage} failed: ${result.error?.message ?? result.stderr ?? `exit ${result.status}`}`);
      return result;
    },
    cargoBuild: (directory, packageName) => {
      const target = targetDirFor(rust);
      ctx.spawn("build", "cargo", [
        "build", "--offline", "--release", "--quiet",
        "--manifest-path", join(directory, "Cargo.toml"),
        "--target-dir", target,
        "--config", "profile.release.debug=false",
        "--config", "profile.release.lto=false",
        "--config", "profile.release.opt-level=3",
      ], { timeout: 1_800_000 });
      return join(target, "release", packageName);
    },
  };

  try {
    await body(ctx);
    ctx.note("OK");
    if (state.workspace && process.env.PBO_NATIVE_KEEP !== "1") {
      rmSync(state.workspace, { recursive: true, force: true });
      ctx.note(`removed workspace ${state.workspace}`);
    } else if (state.workspace) {
      ctx.note(`retained workspace ${state.workspace}`);
    }
  } catch (error) {
    ctx.note(`FAILED: ${error?.stack ?? error}`);
    if (state.workspace) ctx.note(`retained workspace ${state.workspace}`);
    process.exitCode = 1;
  }
}
