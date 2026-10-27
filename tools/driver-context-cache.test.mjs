import assert from "node:assert/strict";
import { createHash } from "node:crypto";
import { spawnSync } from "node:child_process";
import { copyFileSync, cpSync, existsSync, mkdirSync, readFileSync, readdirSync, renameSync, rmSync, statSync, symlinkSync, writeFileSync } from "node:fs";
import { delimiter, dirname, join } from "node:path";
import { fileURLToPath } from "node:url";
import { test } from "node:test";
import { prepareFixture, corePackages, createWorkspace } from "./test-workspace.mjs";
import { findTypedCompiler, verifyTypedOutput } from "./native-workspace.mjs";
import { sourceVariants } from "./cache-source-variants.mjs";

const root = fileURLToPath(new URL("../", import.meta.url));
const hash = bytes => createHash("sha256").update(bytes).digest("hex");
const load = path => JSON.parse(readFileSync(path));
const save = (path, value) => writeFileSync(path, JSON.stringify(value, null, 2) + "\n");

test("compiler context, options and entry transitions equal fresh generation in every selected host", t => {
  const selected = process.env.GOPURS_TEST_MODES?.split(",") ?? ["js", "native-sequential", "native-parallel",
    ...(process.env.GOPURS_TEST_RUST === "1" ? ["rust"] : [])];
  assert.ok(selected.length > 0 && new Set(selected).size === selected.length);
  const modes = [
    { name: "js", js: "1", jobs: "1", pipeline: "0" },
    { name: "native-sequential", js: "0", jobs: "1", pipeline: "0" },
    { name: "native-parallel", js: "0", jobs: "8", pipeline: "1" },
    { name: "rust", js: "0", jobs: "4", pipeline: "1" },
  ].filter(mode => selected.includes(mode.name));
  assert.equal(modes.length, selected.length, "unknown host mode");
  const catalogPath = process.env.GOPURS_TEST_CONTEXT_VARIANTS;
  let catalog;
  if (catalogPath) {
    catalog = load(catalogPath);
    assert.deepEqual(Object.keys(catalog).sort(), ["alternate", "baseline", "runtime"]);
  } else {
    const compiler = process.env.GOPURS_TEST_COMPILER ?? join(root, "bin/gopurs");
    const rust = process.env.GOPURS_TEST_RUST_BINARY ?? join(root, "bin/gopurs-rust");
    const compilerRoot = dirname(dirname(compiler));
    const sourcePath = join(compilerRoot, "src/Gopurs/Runtime.js");
    const embedded = JSON.parse(readFileSync(sourcePath, "utf8").split("export const runtimeGoCode = ")[1].trim().slice(0, -1));
    const files = new Set([sourcePath]);
    for (const mode of modes) {
      if (mode.name === "rust") files.add(rust);
      else {
        files.add(compiler);
        files.add(join(dirname(compiler), mode.js === "1" ? "gopurs.js" : "gopurs-native"));
      }
    }
    catalog = { baseline: { compiler, rust, runtimeSHA256: hash(embedded),
      files: Object.fromEntries([...files].map(file => [file, hash(readFileSync(file))])) } };
  }
  for (const kind of Object.keys(catalog)) {
    for (const key of ["compiler", "rust", "runtimeSHA256"]) assert.equal(typeof catalog[kind]?.[key], "string", `${kind}.${key}`);
    for (const [file, digest] of Object.entries(catalog[kind].files)) assert.equal(hash(readFileSync(file)), digest, file);
  }
  if (catalogPath) {
    assert.equal(catalog.baseline.runtimeSHA256, catalog.alternate.runtimeSHA256);
    assert.notEqual(catalog.baseline.runtimeSHA256, catalog.runtime.runtimeSHA256);
  }
  const workspace = createWorkspace();
  let passed = false;
  t.after(() => {
    if (passed && process.env.GOPURS_TEST_KEEP_WORKSPACE !== "1") rmSync(workspace, { recursive: true, force: true });
    else console.log(`Context transition workspace retained: ${workspace}`);
  });
  save(join(workspace, "variants.json"), catalog);
  copyFileSync(new URL(import.meta.url), join(workspace, "driver-context-cache.test.mjs"));
  const fixtureRoot = join(root, "tests/host/CacheTransitions");
  const state = sourceVariants(fixtureRoot)[0];
  const { directory } = prepareFixture(root, workspace, join(fixtureRoot, "Main.purs"), 0, corePackages(root));
  const output = join(directory, "output"), logs = join(workspace, "logs"), tastRoot = join(workspace, "tast");
  mkdirSync(logs);
  mkdirSync(join(workspace, "results"));
  mkdirSync(join(workspace, "executables"));
  mkdirSync(join(workspace, "hosts"));
  mkdirSync(join(directory, "alternate-ffi"));
  for (const [name, text] of Object.entries(state.sources)) {
    mkdirSync(dirname(join(directory, "src", name)), { recursive: true });
    writeFileSync(join(directory, "src", name), text);
  }
  cpSync(join(directory, "src"), join(workspace, "source"), { recursive: true });
  const typedCompiler = findTypedCompiler(root, process.env.GOPURS_PURS), typedHash = hash(readFileSync(typedCompiler));
  const typedBin = join(workspace, "typed-bin");
  mkdirSync(typedBin);
  symlinkSync(typedCompiler, join(typedBin, "purs"));
  const baseEnv = { ...process.env, GOWORK: "off", GOPURS_PACK_PUBLICATION: "0", GOPURS_GO_PACKAGES: "0", PPROF: "0", PATH: [typedBin, join(root, "node_modules/.bin"), process.env.PATH].join(delimiter) };
  for (const key of ["GOPURS_ALLOC_PROFILE", "GOPURS_FORCE_BUILD", "GOPURS_COMPILER_SNAPSHOT", "GOGC", "GOMEMLIMIT"]) delete baseEnv[key];
  let sequence = 0;
  const commands = [], results = [];
  function command(label, executable, args, env = baseEnv, cwd = directory) {
    const started = process.hrtime.bigint();
    const result = spawnSync(executable, args, { cwd, env, encoding: "utf8", timeout: 180000, maxBuffer: 32 * 1024 * 1024 });
    const log = `${sequence++}-${label}.log`;
    writeFileSync(join(logs, log), (result.stdout ?? "") + (result.stderr ?? ""));
    commands.push({ label, executable, args, cwd, log, status: result.status, error: result.error?.message,
      signal: result.signal, seconds: Number(process.hrtime.bigint() - started) / 1e9 });
    save(join(workspace, "commands.json"), commands);
    assert.ifError(result.error);
    assert.equal(result.status, 0, `${label}: ${result.stdout}\n${result.stderr}`);
    return result;
  }
  command("frontend", "spago", ["build", "-q"]);
  const typed = verifyTypedOutput(output, typedCompiler);
  function tastInventory(path) {
    return Object.fromEntries(readdirSync(path, { withFileTypes: true })
      .filter(entry => entry.isDirectory() && existsSync(join(path, entry.name, "corefn.json")))
      .map(entry => [entry.name, hash(readFileSync(join(path, entry.name, "corefn.json")))]));
  }
  const tast = tastInventory(output), moduleCount = Object.keys(tast).length;
  assert.equal(moduleCount, 66);
  mkdirSync(tastRoot);
  for (const name of Object.keys(tast)) {
    mkdirSync(join(tastRoot, name));
    copyFileSync(join(output, name, "corefn.json"), join(tastRoot, name, "corefn.json"));
  }
  save(join(workspace, "frontend.json"), { compiler: typedCompiler, sha256: typedHash, typed, tast });
  function installTast() { cpSync(tastRoot, output, { recursive: true }); }
  function generated(path = output, prefix = "") {
    return Object.fromEntries(readdirSync(path, { withFileTypes: true }).filter(entry => !entry.name.startsWith("."))
      .flatMap(entry => {
        const name = prefix + entry.name, file = join(path, entry.name);
        if (entry.isDirectory()) return Object.entries(generated(file, name + "/"));
        return name.endsWith(".go") ? [[name, hash(readFileSync(file))]] : [];
      }));
  }
  function timestamps(files) {
    return Object.fromEntries(Object.keys(files).map(file => [file, String(statSync(join(output, file), { bigint: true }).mtimeNs)]));
  }
  function business() {
    const cache = join(output, ".gopurs-cache");
    const readObject = object => {
      const bytes = readFileSync(join(cache, "objects", object.sha256));
      assert.equal(bytes.length, object.bytes);
      assert.equal(hash(bytes), object.sha256);
      return bytes;
    };
    const manifest = JSON.parse(readObject(load(join(cache, "current.json")).manifest));
    const files = generated(), parts = {}, keys = {}, environments = {};
    const identity = file => {
      const stats = statSync(join(output, file), { bigint: true });
      return `${stats.dev}/${stats.ino}`;
    };
    const physical = new Map(Object.keys(files).map(file => [identity(file), file])), owned = new Set();
    assert.equal(physical.size, Object.keys(files).length);
    for (const item of manifest.outputs) {
      const id = identity(item.path), file = physical.get(id);
      assert.ok(file && !owned.has(id), "bijective physical ownership: " + item.path);
      owned.add(id);
      assert.equal(files[file], item.object.sha256);
      assert.equal(statSync(join(output, item.path)).size, item.object.bytes);
    }
    assert.equal(owned.size, physical.size);
    assert.equal(manifest.modules.length, moduleCount);
    for (const module of manifest.modules) {
      assert.equal(module.parts.length, 4);
      for (const part of module.parts) readObject(part.object);
      parts[module.name] = module.parts;
      keys[module.name] = [module.prepared, module.callerDemands];
      environments[module.name] = [module.optimizerEnvironment, module.emitterEnvironment];
    }
    return { manifest, portable: { files, parts, keys }, environments };
  }
  const references = new Map();
  function transitions(mode) {
    const differentJobs = mode.jobs === "1" ? "2" : "1";
    const states = [{ label: "initial" }, { label: "rewrite-limit", args: ["--rewrite-limit", "9999"] }, { label: "rewrite-restored" }];
    for (const [label, name] of [["prepare", "GOPURS_PREPARE_JOBS"], ["optimizer", "GOPURS_PBO_JOBS"], ["emission", "GOPURS_EMIT_JOBS"]])
      states.push({ label: label + "-jobs", env: { [name]: differentJobs } }, { label: label + "-restored" });
    states.push({ label: "pipeline", env: { GOPURS_PIPELINE: mode.pipeline === "0" ? "1" : "0" } }, { label: "pipeline-restored" },
      { label: "ffi-root", args: ["--ffi", "alternate-ffi"] }, { label: "ffi-root-restored" },
      { label: "secondary-entry", args: ["--main", "Cache.Move"], main: "Cache.Move" },
      { label: "main-entry", args: ["--main", "Main"], main: "Main" }, { label: "all-entries-restored" });
    if (catalogPath) states.push({ label: "compiler-version", compiler: "alternate" }, { label: "compiler-restored", identityChanged: true },
      { label: "runtime-version", compiler: "runtime" }, { label: "runtime-restored", identityChanged: true });
    return states;
  }
  const allMiss = new RegExp(`module cache: optimized=${moduleCount}, reused=0, emitted=${moduleCount}, restored=0`);
  for (const mode of modes) {
    rmSync(output, { recursive: true });
    installTast();
    let previous, previousDates, initial;
    for (const state of transitions(mode)) {
      const label = mode.name + "-" + state.label, variant = catalog[state.compiler ?? "baseline"];
      const compiler = mode.name === "rust" ? variant.rust : variant.compiler;
      const env = { ...baseEnv, GOPURS_JS: mode.js, GOPURS_RUST: mode.name === "rust" ? "1" : "0",
        GOPURS_PREPARE_JOBS: mode.jobs, GOPURS_PBO_JOBS: mode.jobs, GOPURS_EMIT_JOBS: mode.jobs,
        GOPURS_PIPELINE: mode.pipeline, ...state.env };
      const invoke = suffix => command(label + suffix, compiler, state.args ?? [], env);
      const incremental = invoke("");
      assert.match(incremental.stderr, previous ? /cache miss: inputs-changed/ : /cache miss:/);
      assert.match(incremental.stderr, allMiss);
      const prepared = /prepared fingerprints: computed=(\d+), reused=(\d+)/.exec(incremental.stderr);
      assert.ok(prepared, "prepared fingerprint diagnostic");
      assert.equal(Number(prepared[1]) + Number(prepared[2]), moduleCount, "every prepared module accounted for");
      assert.doesNotMatch(incremental.stderr, /optimizer not started/);
      const actual = business(), dates = timestamps(actual.portable.files);
      const rewritten = Object.keys(dates).filter(file => dates[file] !== previousDates?.[file]);
      if (previous) {
        for (const [file, digest] of Object.entries(previous.portable.files)) {
          if (actual.portable.files[file] === digest) assert.equal(dates[file], previousDates[file], label + ": unchanged output rewritten " + file);
        }
        assert.notEqual(actual.manifest.snapshot.key, previous.manifest.snapshot.key);
        const changedIdentity = state.compiler !== undefined || state.identityChanged === true;
        if (changedIdentity) assert.notEqual(actual.manifest.snapshot.compiler, previous.manifest.snapshot.compiler);
        else assert.equal(actual.manifest.snapshot.compiler, previous.manifest.snapshot.compiler);
        for (const name of Object.keys(actual.environments)) assert.notDeepEqual(actual.environments[name], previous.environments[name]);
      }
      assert.equal(actual.portable.files["gopurs_runtime/runtime.go"], variant.runtimeSHA256);
      const entries = actual.manifest.outputs.filter(item => item.kind === "entry").map(item => [item.module, item.path]);
      if (state.main) assert.deepEqual(entries, [[state.main, state.main + "/main/main.go"], [state.main, "main/main.go"]]);
      else assert.deepEqual(entries.map(item => item[1]).sort(), ["Cache.Move/main/main.go", "Main/main/main.go", "main/main.go"].sort());
      if (references.has(state.label)) assert.deepEqual(actual.portable, references.get(state.label), label + ": cross-host bytes");
      else references.set(state.label, actual.portable);
      if (initial && !state.main) {
        const expected = structuredClone(initial.portable);
        expected.files["gopurs_runtime/runtime.go"] = variant.runtimeSHA256;
        assert.deepEqual(actual.portable, expected, label + ": only runtime bytes can change in this state");
      }
      const hit = invoke("-identical");
      assert.match(hit.stderr, /cache hit: .*0 modules optimized\/emitted/);
      assert.deepEqual(timestamps(actual.portable.files), dates);
      const saved = join(workspace, "saved-output");
      assert.ok(!existsSync(saved));
      renameSync(output, saved);
      installTast();
      for (const name of ["go.mod", "go.sum"]) if (existsSync(join(saved, name))) copyFileSync(join(saved, name), join(output, name));
      assert.match(invoke("-fresh").stderr, allMiss);
      const fresh = business();
      assert.deepEqual(fresh.portable, actual.portable);
      assert.deepEqual(fresh.environments, actual.environments);
      assert.deepEqual(fresh.manifest.snapshot, actual.manifest.snapshot);
      rmSync(output, { recursive: true });
      renameSync(saved, output);
      assert.deepEqual(timestamps(actual.portable.files), dates);
      const binary = join(workspace, "executables", label);
      command(label + "-build", "go", ["build", "-trimpath", "-o", binary, "./main/main.go"], env, output);
      const execution = command(label + "-execute", binary, [], env, output);
      assert.equal(execution.stdout, state.main === "Cache.Move" ? "secondary\n" : "14\n5\n8\n6\n5\n9\nDone\n");
      assert.equal(execution.stderr, "");
      assert.deepEqual(tastInventory(output), tast);
      assert.deepEqual(generated(), actual.portable.files);
      save(join(workspace, "results", label + ".json"), actual);
      results.push({ mode: mode.name, state: state.label, compiler: actual.manifest.snapshot.compiler,
        snapshot: actual.manifest.snapshot.key, sameTast: true, sameFresh: true, identicalHit: true, executed: true,
        rewritten, counters: incremental.stderr.match(/^\[gopurs\] module cache:.*$/m)?.[0] });
      save(join(workspace, "results.json"), results);
      initial ??= actual;
      previous = actual;
      previousDates = dates;
    }
    renameSync(output, join(workspace, "hosts", mode.name));
    mkdirSync(output);
  }
  for (const variant of Object.values(catalog)) for (const [file, digest] of Object.entries(variant.files))
    assert.equal(hash(readFileSync(file)), digest, file);
  assert.equal(hash(readFileSync(typedCompiler)), typedHash);
  save(join(workspace, "summary.json"), { completed: true, modes: modes.map(mode => mode.name),
    statesPerMode: transitions(modes[0]).length, comparisons: results.length, successfulCommands: commands.length,
    compilerAndRuntimeVariants: Boolean(catalogPath), sameTast: true, samePartsAndKeys: true, sameGo: true, executed: true });
  passed = true;
});
