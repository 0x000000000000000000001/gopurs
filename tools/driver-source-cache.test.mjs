import assert from "node:assert/strict";
import { createHash } from "node:crypto";
import { spawnSync } from "node:child_process";
import { copyFileSync, cpSync, existsSync, mkdirSync, readFileSync, readdirSync, renameSync, rmSync, statSync, symlinkSync, unlinkSync, writeFileSync } from "node:fs";
import { delimiter, dirname, join } from "node:path";
import { fileURLToPath } from "node:url";
import { test } from "node:test";
import { gunzipSync } from "node:zlib";
import { corePackages, createWorkspace, prepareFixture } from "./test-workspace.mjs";
import { findTypedCompiler, verifyTypedOutput } from "./native-workspace.mjs";
import { sourceVariants } from "./cache-source-variants.mjs";
import { packageSourceSections } from "./test-snapshots.mjs";

const root = fileURLToPath(new URL("../", import.meta.url));
const hash = bytes => createHash("sha256").update(bytes).digest("hex");
const load = path => JSON.parse(readFileSync(path));
const save = (path, value) => writeFileSync(path, JSON.stringify(value, null, 2) + "\n");

test("source edits, ABI layouts, specialization retraction and renames equal fresh builds in every host", async t => {
  const workspace = createWorkspace();
  let passed = false;
  t.after(() => {
    if (passed && process.env.GOPURS_TEST_KEEP_WORKSPACE !== "1") rmSync(workspace, { recursive: true, force: true });
    else console.log(`Source transition workspace retained: ${workspace}`);
  });
  const fixtureDirectory = join(root, "tests/host/CacheTransitions");
  mkdirSync(join(workspace, "harness"));
  for (const name of ["driver-source-cache.test.mjs", "cache-source-variants.mjs"])
    copyFileSync(join(root, "tools", name), join(workspace, "harness", name));
  cpSync(fixtureDirectory, join(workspace, "harness/fixture"), { recursive: true });
  const fixture = prepareFixture(root, workspace, join(fixtureDirectory, "Main.purs"), 0, corePackages(root));
  const directory = fixture.directory, output = join(directory, "output"), sourceDirectory = join(directory, "src");
  const logs = join(workspace, "logs");
  mkdirSync(logs);
  mkdirSync(join(workspace, "executables"));
  const typedCompiler = findTypedCompiler(root, process.env.GOPURS_PURS), typedHash = hash(readFileSync(typedCompiler));
  const typedBin = join(workspace, "typed-bin");
  mkdirSync(typedBin);
  symlinkSync(typedCompiler, join(typedBin, "purs"));
  const baseEnv = { ...process.env, GOWORK: "off", PATH: [typedBin, join(root, "node_modules/.bin"), process.env.PATH].join(delimiter) };
  for (const key of ["GOPURS_ALLOC_PROFILE", "GOPURS_FORCE_BUILD", "GOPURS_COMPILER_SNAPSHOT", "PPROF", "GOGC", "GOMEMLIMIT"]) delete baseEnv[key];
  const commands = [], results = [], pathAliases = new Map();
  let sequence = 0, installedSources = { "Main.purs": readFileSync(join(sourceDirectory, "Main.purs"), "utf8") };

  function command(label, executable, args, { cwd = directory, env = baseEnv, timeout = 120000 } = {}) {
    const started = process.hrtime.bigint();
    const result = spawnSync(executable, args, { cwd, env, encoding: "utf8", timeout, maxBuffer: 32 * 1024 * 1024 });
    const log = `${sequence++}-${label}.log`;
    writeFileSync(join(logs, log), (result.stdout ?? "") + (result.stderr ?? ""));
    commands.push({ label, executable, args, cwd, log, status: result.status, signal: result.signal, error: result.error?.message,
      seconds: Number(process.hrtime.bigint() - started) / 1e9 });
    save(join(workspace, "commands.json"), commands);
    assert.ifError(result.error);
    assert.equal(result.status, 0, `${label}: ${result.stdout}\n${result.stderr}`);
    return result;
  }

  function installSources(sources) {
    for (const name of Object.keys(installedSources)) if (!(name in sources)) unlinkSync(join(sourceDirectory, name));
    for (const [name, text] of Object.entries(sources)) {
      const target = join(sourceDirectory, name);
      mkdirSync(dirname(target), { recursive: true });
      if (!existsSync(target) || readFileSync(target, "utf8") !== text) writeFileSync(target, text);
    }
    installedSources = { ...sources };
  }

  function tastInventory(path) {
    return Object.fromEntries(readdirSync(path, { withFileTypes: true }).filter(entry => entry.isDirectory() && existsSync(join(path, entry.name, "corefn.json")))
      .map(entry => [entry.name, hash(readFileSync(join(path, entry.name, "corefn.json")))]));
  }

  function installTast(state, target = output) {
    mkdirSync(target, { recursive: true });
    for (const name of Object.keys(tastInventory(target))) if (!(name in state.tast)) unlinkSync(join(target, name, "corefn.json"));
    for (const name of Object.keys(state.tast)) {
      mkdirSync(join(target, name), { recursive: true });
      copyFileSync(join(state.snapshot, name, "corefn.json"), join(target, name, "corefn.json"));
    }
    assert.deepEqual(tastInventory(target), state.tast);
  }

  function generated(path = output, prefix = "") {
    const result = {};
    for (const entry of readdirSync(path, { withFileTypes: true })) {
      if (entry.name.startsWith(".")) continue;
      const file = join(path, entry.name), name = prefix + entry.name;
      if (entry.isDirectory()) Object.assign(result, generated(file, name + "/"));
      else if (name.endsWith(".go")) result[name] = hash(readFileSync(file));
    }
    return result;
  }

  function timestamps(files = generated()) {
    return Object.fromEntries(Object.keys(files).map(name => [name, String(statSync(join(output, name), { bigint: true }).mtimeNs)]));
  }

  function business() {
    const cache = join(output, ".gopurs-cache"), pointer = load(join(cache, "current.json"));
    const readObject = object => {
      const bytes = readFileSync(join(cache, "objects", object.sha256));
      assert.equal(bytes.length, object.bytes);
      assert.equal(hash(bytes), object.sha256);
      return bytes;
    };
    const manifest = JSON.parse(readObject(pointer.manifest));
    const parts = {}, keys = {}, sources = {}, moduleOutputs = {};
    for (const module of manifest.modules) {
      for (const part of module.parts) readObject(part.object);
      assert.equal(module.parts.length, 4);
      parts[module.name] = module.parts;
      keys[module.name] = [module.prepared, module.callerDemands];
      if (process.env.GOPURS_GO_PACKAGES === "1") {
        assert(module.sources?.length >= 1, module.name + ": unqualified source required");
        sources[module.name] = module.sources;
        for (const source of module.sources) {
          assert.equal(source.codec, "gopurs/go-source-gzip-v1");
          const raw = gunzipSync(readObject(source.object));
          assert.equal(raw.length, source.raw.bytes);
          assert.equal(hash(raw), source.raw.sha256);
        }
      }
    }
    const files = generated();
    const identity = path => {
      const stat = statSync(join(output, path), { bigint: true });
      assert.ok(stat.isFile(), path + ": not a regular output");
      return `${stat.dev}/${stat.ino}`;
    };
    // Main/ can also contain main/main.go on a case-insensitive filesystem.
    // Require a bijection with the manifest, including hashes and sizes.
    const byIdentity = new Map(Object.keys(files).map(path => [identity(path), path]));
    assert.equal(byIdentity.size, Object.keys(files).length, "distinct physical outputs");
    const owned = new Set();
    const sourceOwners = new Map(manifest.modules.flatMap(module => [
      [module.name.replaceAll(".", "_") + ".go", { module: module.name, kind: "package" }],
      [module.name.replaceAll(".", "_") + "_ffi.go", { module: module.name, kind: "ffi-package" }],
    ]));
    for (const item of manifest.outputs) {
      const key = identity(item.path), physical = byIdentity.get(key);
      assert.ok(physical !== undefined, item.path + ": unlisted physical output");
      assert.ok(!owned.has(key), item.path + ": multiply owned output");
      owned.add(key);
      assert.equal(files[physical], item.object.sha256, item.path + ": output hash");
      assert.equal(statSync(join(output, item.path)).size, item.object.bytes, item.path + ": output size");
      if (physical !== item.path) pathAliases.set(item.path, physical);
      if (item.module) (moduleOutputs[item.module] ??= []).push({ path: item.path, kind: item.kind, object: item.object });
      if (item.path.endsWith("/0_gopurs_package.go")) {
        const { sections } = packageSourceSections(readFileSync(join(output, item.path), "utf8"));
        assert(sections.length > 0, "coalesced package has no source provenance");
        for (const section of sections) {
          const owner = sourceOwners.get(section.name);
          assert(owner, "stale source in package: " + section.name);
          (moduleOutputs[owner.module] ??= []).push({ path: item.path, source: section.name, kind: owner.kind,
            object: { sha256: hash(section.content), bytes: Buffer.byteLength(section.content) } });
        }
      }
    }
    assert.equal(owned.size, byIdentity.size, "no unowned generated Go");
    save(join(workspace, "path-aliases.json"), Object.fromEntries(pathAliases));
    return { files, parts, keys, sources, moduleOutputs };
  }

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
  const states = sourceVariants(fixtureDirectory);
  for (const [index, state] of states.entries()) {
    installSources(state.sources);
    rmSync(output, { recursive: true, force: true });
    command("frontend-" + state.label, "spago", ["build", "-q"]);
    const typed = verifyTypedOutput(output, typedCompiler);
    state.tast = tastInventory(output);
    state.snapshot = join(workspace, "states", state.label, "tast");
    mkdirSync(state.snapshot, { recursive: true });
    for (const name of Object.keys(state.tast)) {
      mkdirSync(join(state.snapshot, name));
      copyFileSync(join(output, name, "corefn.json"), join(state.snapshot, name, "corefn.json"));
    }
    for (const name of state.sameTast ?? []) assert.equal(state.tast[name], states[index - 1].tast[name], state.label + ": unchanged provider TAST " + name);
    if (state.restoredAll) assert.deepEqual(state.tast, states[0].tast, "restored sources reproduce the initial TAST");
    const sources = join(workspace, "states", state.label, "src");
    for (const [name, text] of Object.entries(state.sources)) {
      mkdirSync(dirname(join(sources, name)), { recursive: true });
      writeFileSync(join(sources, name), text);
    }
    save(join(workspace, "states", state.label, "input.json"), { label: state.label, stdout: state.stdout, typed, tast: state.tast,
      sources: Object.fromEntries(Object.entries(state.sources).map(([name, text]) => [name, hash(text)])) });
  }
  assert.equal(hash(readFileSync(typedCompiler)), typedHash);
  save(join(workspace, "frontend.json"), { compiler: typedCompiler, sha256: typedHash, states: states.map(state => state.label) });

  const references = new Map();
  let completedModes = 0;
  for (const mode of modes) await t.test(mode.name, () => {
    rmSync(output, { recursive: true, force: true });
    const compiler = mode.name === "rust" ? process.env.GOPURS_TEST_RUST_BINARY ?? join(root, "bin/gopurs-rust")
      : process.env.GOPURS_TEST_COMPILER ?? join(root, "bin/gopurs");
    const compilerFiles = mode.name === "rust" ? [compiler] : [compiler, join(dirname(compiler), mode.js === "1" ? "gopurs.js" : "gopurs-native")];
    const compilerHashes = Object.fromEntries(compilerFiles.map(path => [path, hash(readFileSync(path))]));
    const env = { ...baseEnv, GOPURS_JS: mode.js, GOPURS_RUST: mode.name === "rust" ? "1" : "0",
      GOPURS_PREPARE_JOBS: mode.jobs, GOPURS_PBO_JOBS: mode.jobs, GOPURS_EMIT_JOBS: mode.jobs, GOPURS_PIPELINE: mode.pipeline };
    const invoke = label => command(mode.name + "-" + label, compiler, [], { env });
    let previous, initial;
    for (const state of states) {
      installSources(state.sources);
      installTast(state);
      const incremental = invoke(state.label);
      assert.match(incremental.stderr, /cache miss:/);
      const actual = business(), dates = timestamps(actual.files);
      if (references.has(state.label)) assert.deepEqual(actual, references.get(state.label), state.label + ": cross-host outputs, parts and prepared keys");
      else references.set(state.label, actual);
      for (const name of state.changedModules ?? []) assert.notDeepEqual(actual.keys[name], previous.keys[name], state.label + ": changed prepared input " + name);
      for (const name of state.changedPrepared ?? []) {
        assert.notDeepEqual(actual.keys[name], previous.keys[name], state.label + ": specialization provider must change " + name);
        assert.notDeepEqual(actual.moduleOutputs[name], previous.moduleOutputs[name], state.label + ": provider output must change " + name);
      }
      for (const name of state.restoredPrepared ?? []) assert.deepEqual(actual.keys[name], initial.keys[name], state.label + ": removed demands restore " + name);
      for (const name of state.removedModules ?? []) {
        assert.ok(!(name in actual.keys), state.label + ": stale memo " + name);
        assert.ok(!(name in actual.moduleOutputs), state.label + ": stale owned output " + name);
      }
      for (const name of state.addedModules ?? []) {
        assert.ok(name in actual.keys);
        for (const kind of process.env.GOPURS_GO_PACKAGES === "1" ? ["package", "ffi-package", "entry"] : ["module", "ffi", "entry"])
          assert(actual.moduleOutputs[name]?.some(file => file.kind === kind), state.label + ": new " + kind + " output " + name);
      }
      if (state.restoredAll) assert.deepEqual(actual, initial, "full return reproduces original outputs, parts and prepared keys");
      const hit = invoke(state.label + "-identical");
      assert.match(hit.stderr, /cache hit: .*0 modules optimized\/emitted/);
      assert.deepEqual(timestamps(actual.files), dates);
      const saved = join(workspace, "saved-output");
      assert.ok(!existsSync(saved));
      renameSync(output, saved);
      installTast(state);
      for (const file of ["go.mod", "go.sum"]) if (existsSync(join(saved, file))) copyFileSync(join(saved, file), join(output, file));
      const fresh = invoke(state.label + "-fresh");
      assert.match(fresh.stderr, /module cache: optimized=[1-9]\d*, reused=0, emitted=[1-9]\d*, restored=0/);
      assert.deepEqual(business(), actual, state.label + ": incremental/fresh outputs, parts and prepared keys");
      rmSync(output, { recursive: true });
      renameSync(saved, output);
      assert.deepEqual(timestamps(actual.files), dates);
      const binary = join(workspace, "executables", mode.name + "-" + state.label);
      command(mode.name + "-" + state.label + "-build", "go", ["build", "-p", "1", "-x", "-trimpath", "-o", binary, "./Main/main/main.go"], { cwd: output, env });
      const identicalBuild = command(mode.name + "-" + state.label + "-identical-build", "go", ["build", "-p", "1", "-x", "-trimpath", "-o", binary, "./Main/main/main.go"], { cwd: output, env });
      assert.doesNotMatch(identicalBuild.stderr, /\/compile\b|\/link\b/);
      const execution = command(mode.name + "-" + state.label + "-execute", binary, [], { cwd: output, env });
      assert.equal(execution.stdout.replaceAll("\r\n", "\n"), state.stdout, state.label + ": executable oracle");
      assert.equal(execution.stderr, "", state.label + ": unexpected runtime diagnostic");
      assert.deepEqual(generated(), actual.files, "Go execution preserves generated sources");
      assert.deepEqual(tastInventory(output), state.tast);
      save(join(workspace, "states", state.label, mode.name + "-result.json"), actual);
      results.push({ mode: mode.name, state: state.label, byteExact: true, executed: true,
        counters: incremental.stderr.match(/^\[gopurs\] module cache:.*$/m)?.[0] });
      save(join(workspace, "results.json"), results);
      initial ??= actual;
      previous = actual;
    }
    assert.deepEqual(Object.fromEntries(compilerFiles.map(path => [path, hash(readFileSync(path))])), compilerHashes);
    const archive = join(workspace, "hosts", mode.name);
    mkdirSync(archive, { recursive: true });
    renameSync(output, join(archive, "output"));
    save(join(archive, "compiler.json"), { files: compilerHashes });
    completedModes++;
  });
  assert.equal(completedModes, modes.length, "every host completed its audit");
  assert.equal(results.length, states.length * modes.length);
  save(join(workspace, "summary.json"), { completed: true, states: states.length, modes: modes.map(mode => mode.name),
    comparisons: results.length, sameTast: true, samePartsAndPrepared: true, sameGo: true, executed: true });
  passed = true;
});
