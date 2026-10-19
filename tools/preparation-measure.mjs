import assert from "node:assert/strict";
import { createHash } from "node:crypto";
import { observe } from "./preparation-observe.mjs";
import {
  constants, copyFileSync, cpSync, existsSync, mkdirSync, readFileSync, readdirSync,
  realpathSync, statSync, unlinkSync, writeFileSync,
} from "node:fs";
import { dirname, join, resolve } from "node:path";

const [operation, directory] = process.argv.slice(2);
assert(directory && ["prepare", "pairs"].includes(operation), "Usage: node --expose-gc tools/preparation-measure.mjs prepare|pairs CAMPAIGN");
const campaign = realpathSync(resolve(directory));
const measurement = join(campaign, "measurement");
const hash = bytes => createHash("sha256").update(bytes).digest("hex");
const load = path => JSON.parse(readFileSync(path, "utf8"));
const save = (path, data) => { mkdirSync(dirname(path), { recursive: true }); writeFileSync(path, JSON.stringify(data, null, 2) + "\n"); };
const lockName = ".gopurs-cache/write.lock";
const modes = {
  reference: { app: join(campaign, "tree/b8x"), compiler: join(campaign, "tree/gopurs/gopurs/bin/gopurs"), cache: join(campaign, "host-go-cache") },
  diagnostic: { app: join(campaign, "diagnostic/b8x"), compiler: join(campaign, "diagnostic/gopurs/gopurs/bin/gopurs"), cache: join(measurement, "diagnostic-go-cache") },
};
const mutationPaths = ["src/Inter/Cli/Ping/Main.purs", "src/Util/Type/String/String.purs", "src/Util/Type/String/String.go"];

function files(directory, prefix = "") {
  if (!existsSync(directory)) return [];
  return readdirSync(directory, { withFileTypes: true }).sort((a, b) => a.name.localeCompare(b.name)).flatMap(entry => {
    const path = join(directory, entry.name), name = prefix + entry.name;
    assert(!entry.isSymbolicLink(), "Private snapshot contains a link: " + path);
    return entry.isDirectory() ? files(path, name + "/") : [name];
  });
}

function copyTree(source, target) {
  assert(!existsSync(target), "Snapshot exists: " + target);
  mkdirSync(dirname(target), { recursive: true });
  cpSync(source, target, { recursive: true, dereference: true, mode: constants.COPYFILE_FICLONE,
    filter: path => !path.endsWith("/write.lock") });
}

function environment(mode) {
  const env = { ...process.env };
  for (const key of Object.keys(env)) if (key.startsWith("GOPURS_") || ["PPROF", "GOFLAGS", "GOCACHE", "GOGC", "GOMEMLIMIT", "GOMAXPROCS", "NODE_OPTIONS"].includes(key)) delete env[key];
  return { ...env, PATH: join(campaign, "toolchain") + ":" + process.env.PATH, TMPDIR: join(campaign, "tmp"),
    GOWORK: "off", GOCACHE: mode.cache, GOPURS_GO_PACKAGES: "1", GOPURS_PREPARE_JOBS: "8",
    GOPURS_PBO_JOBS: "8", GOPURS_EMIT_JOBS: "8", GOPURS_PIPELINE: "1", GOGC: "off", GOMEMLIMIT: "10GiB" };
}

function command(label, mode, executable, args) {
  const base = join(measurement, "logs", label);
  assert(!existsSync(base + ".json"), "Do not overwrite an earlier attempt: " + label);
  mkdirSync(dirname(base), { recursive: true });
  global.gc?.();
  const start = process.hrtime.bigint();
  const result = observe("/usr/bin/time", ["-l", executable, ...args], { cwd: mode.app, env: environment(mode),
    encoding: "utf8", timeout: 1800000, maxBuffer: 64 * 1024 * 1024 }, base + ".resources.jsonl");
  const seconds = Number(process.hrtime.bigint() - start) / 1e9;
  writeFileSync(base + ".stdout", result.stdout ?? "");
  writeFileSync(base + ".stderr", result.stderr ?? "");
  const metrics = [...(result.stderr ?? "").matchAll(/^\[gopurs\] (.*): (\d+) ms(?: \(failed\))?$/gm)]
    .map(([, phase, ms]) => ({ phase, seconds: Number(ms) / 1000 }));
  const record = { label, executable, args, cwd: mode.app, started: new Date(Date.now() - seconds * 1000).toISOString(), seconds, metrics, status: result.status, signal: result.signal,
    error: result.error?.message, peakRssBytes: Number(result.stderr?.match(/(\d+)\s+maximum resident set size/)?.[1] ?? 0) };
  save(base + ".json", record);
  assert.ifError(result.error);
  assert.equal(result.status, 0, `${label} failed; see ${base}.stderr`);
  console.log(`${label}: ${seconds.toFixed(3)} s`);
  return record;
}

function frontend(label, mode) {
  return command(label, mode, process.execPath, [join(campaign, "toolchain/spago.bundle.js"), "build", "--offline", "--monochrome"]);
}

function tast(output) {
  return Object.fromEntries(readdirSync(output).sort().filter(name => existsSync(join(output, name, "corefn.json")))
    .map(name => [name, hash(readFileSync(join(output, name, "corefn.json")))]));
}

function installState(mode, state) {
  for (const [name, text] of Object.entries(state.sources)) {
    const path = join(mode.app, name);
    if (readFileSync(path, "utf8") !== text) writeFileSync(path, text);
  }
  const output = join(mode.app, "output");
  for (const name of Object.keys(state.tast)) {
    const target = join(output, name, "corefn.json");
    mkdirSync(dirname(target), { recursive: true });
    copyFileSync(join(campaign, "inputs", state.name, "tast", name, "corefn.json"), target);
  }
  assert.deepEqual(tast(output), state.tast);
}

function retain(mode, label) {
  const output = join(mode.app, "output"), cache = join(output, ".gopurs-cache");
  assert.deepEqual(readdirSync(join(cache, "stages")), [], "Unfinished publication");
  const pointer = load(join(cache, "current.json"));
  const data = readFileSync(join(cache, "objects", pointer.manifest.sha256));
  assert.equal(hash(data), pointer.manifest.sha256);
  const manifest = JSON.parse(data);
  assert.equal(manifest.snapshot.output, realpathSync(output), "Never activate a transplanted journal");
  const cas = join(measurement, "objects");
  mkdirSync(cas, { recursive: true });
  let bytes = 0;
  for (const name of readdirSync(join(cache, "objects"))) {
    const source = join(cache, "objects", name), value = readFileSync(source);
    assert.equal(hash(value), name);
    bytes += value.length;
    if (!existsSync(join(cas, name))) copyFileSync(source, join(cas, name), constants.COPYFILE_FICLONE);
  }
  for (const item of manifest.outputs) assert.equal(hash(readFileSync(join(output, item.path))), item.object.sha256);
  save(join(measurement, "manifests", label + ".json"), { pointer, manifest, bytes });
  for (const name of files(output).filter(name => name.startsWith(".gopurs") && !name.startsWith(".gopurs-cache/"))) {
    const target = join(measurement, "journals", label, name);
    mkdirSync(dirname(target), { recursive: true });
    copyFileSync(join(output, name), target);
  }
  return { outputs: manifest.outputs, modules: manifest.modules.map(mod => ({ name: mod.name, prepared: mod.prepared,
    callerDemands: mod.callerDemands, parts: mod.parts, sources: mod.sources })), bytes };
}

function business(value) {
  return { outputs: value.outputs, modules: value.modules };
}

// Snapshots are restored only to the original private output path. Its lock
// inode remains in place, and archived journals are never used at archive paths.
function restore(source, target, preserveLock = false) {
  assert(target.startsWith(campaign + "/"));
  const wanted = new Set(files(source));
  const lock = preserveLock ? statSync(join(target, lockName), { bigint: true }).ino : null;
  for (const name of files(target)) if (!wanted.has(name) && !(preserveLock && name === lockName)) unlinkSync(join(target, name));
  for (const name of wanted) {
    if (preserveLock && name === lockName) continue;
    mkdirSync(dirname(join(target, name)), { recursive: true });
    copyFileSync(join(source, name), join(target, name), constants.COPYFILE_FICLONE);
  }
  if (preserveLock) assert.equal(statSync(join(target, lockName), { bigint: true }).ino, lock);
}

if (operation === "prepare") {
  assert(!existsSync(join(measurement, "states.json")), "Preparation already recorded");
  mkdirSync(measurement, { recursive: true });
  const mode = modes.reference, original = Object.fromEntries(mutationPaths.map(path => [path, readFileSync(join(mode.app, path), "utf8")]));
  const base = Object.fromEntries(Object.entries(load(join(campaign, "frontend-base.json")).files).map(([name, value]) => [name, value.sha256]));
  assert.deepEqual(tast(join(mode.app, "output")), base);
  const states = [{ name: "base", sources: original, tast: base }];
  const edits = [
    ["leaf", mutationPaths[0], 'log "pong"', 'log "pong-lot00"'],
    ["shared", mutationPaths[1], "caseToCamel = lowerCaseFirst ◁ caseToPascal", 'caseToCamel input = (lowerCaseFirst ◁ caseToPascal) input <> "-lot00"'],
    ["ffi", mutationPaths[2], "return result", 'return result + "!"'],
  ];
  for (const [name, file, before, after] of edits) {
    assert.equal(original[file].split(before).length, 2, "Unique mutation target");
    const sources = { ...original, [file]: original[file].replace(before, after) };
    for (const [path, text] of Object.entries(sources)) if (readFileSync(join(mode.app, path), "utf8") !== text) writeFileSync(join(mode.app, path), text);
    frontend("frontend-" + name, mode);
    const current = tast(join(mode.app, "output"));
    assert.deepEqual(Object.keys(current), Object.keys(base));
    const changed = Object.keys(base).filter(name => base[name] !== current[name]);
    assert.deepEqual(changed, name === "leaf" ? ["Inter.Cli.Ping.Main"] : name === "shared" ? ["Util.Type.String.String"] : []);
    for (const module of Object.keys(current)) {
      const target = join(campaign, "inputs", name, "tast", module, "corefn.json");
      assert(!existsSync(target)); mkdirSync(dirname(target), { recursive: true });
      copyFileSync(join(mode.app, "output", module, "corefn.json"), target);
    }
    states.push({ name, sources, tast: current, changed });
  }
  for (const [path, text] of Object.entries(original)) if (readFileSync(join(mode.app, path), "utf8") !== text) writeFileSync(join(mode.app, path), text);
  frontend("frontend-restored", mode);
  assert.deepEqual(tast(join(mode.app, "output")), base);
  command("reference-base-hit", mode, mode.compiler, []);
  const reference = retain(mode, "reference-base");
  assert(!existsSync(modes.diagnostic.app)); mkdirSync(modes.diagnostic.app);
  for (const part of ["src", "test", ".spago", "spago.yaml", "spago.lock"]) copyTree(join(mode.app, part), join(modes.diagnostic.app, part));
  mkdirSync(join(modes.diagnostic.app, "output"));
  installState(modes.diagnostic, states[0]);
  for (const name of ["go.mod", "go.sum"]) copyFileSync(join(campaign, "inputs/original", name), join(modes.diagnostic.app, "output", name));
  command("diagnostic-base", modes.diagnostic, modes.diagnostic.compiler, []);
  const diagnostic = retain(modes.diagnostic, "diagnostic-base");
  assert.deepEqual(business(diagnostic), business(reference), "Instrumented preparation changed business objects or Go");
  for (const [name, value] of Object.entries(modes)) {
    copyTree(join(value.app, "output"), join(measurement, "seeds", name, "output"));
    copyTree(value.cache, join(measurement, "seeds", name, "go-cache"));
  }
  save(join(measurement, "states.json"), states);
  save(join(measurement, "base.json"), { reference, diagnostic });
  console.log("All four TAST states frozen; instrumented/reference outputs and prepared/business objects are byte-exact");
} else {
  const states = load(join(measurement, "states.json"));
  assert(!existsSync(join(measurement, "results.json")), "Preserve previous campaign results");
  const results = [], expected = new Map([["base", business(load(join(measurement, "base.json")).reference)]]);
  for (const scenario of ["identical", "leaf", "shared", "ffi", "cold-app", "cold-complete"]) {
    const state = states.find(value => value.name === (states.some(value => value.name === scenario) ? scenario : "base"));
    for (let pair = 1; pair <= 3; pair++) for (const name of pair % 2 ? ["reference", "diagnostic"] : ["diagnostic", "reference"]) {
      const mode = modes[name], output = join(mode.app, "output"), label = `${scenario}-${pair}-${name}`;
      restore(join(measurement, "seeds", name, "output"), output, true);
      restore(join(measurement, "seeds", name, "go-cache"), mode.cache);
      installState(mode, state);
      let frontendResult = null;
      if (scenario.startsWith("cold-")) {
        const preserve = new Set([lockName, "go.mod", "go.sum"]);
        if (scenario === "cold-app") for (const file of files(output)) if (!file.startsWith(".gopurs") && !file.endsWith(".go")) preserve.add(file);
        for (const file of files(output)) if (!preserve.has(file)) unlinkSync(join(output, file));
        if (scenario === "cold-complete") {
          for (const file of files(mode.cache)) unlinkSync(join(mode.cache, file));
          frontendResult = frontend(label + "-frontend", mode);
          assert.deepEqual(tast(output), states[0].tast);
        }
      }
      const dates = Object.fromEntries(files(output).filter(path => path.endsWith(".go")).map(path => [path, String(statSync(join(output, path), { bigint: true }).mtimeNs)]));
      const result = command(label, mode, mode.compiler, []);
      assert.deepEqual(tast(output), state.tast);
      const actual = retain(mode, label);
      if (expected.has(state.name)) assert.deepEqual(business(actual), expected.get(state.name), "Paired/repeated outputs differ: " + label);
      else expected.set(state.name, business(actual));
      const stderr = readFileSync(join(measurement, "logs", label + ".stderr"), "utf8");
      if (scenario === "identical") {
        assert.match(stderr, /cache hit:/);
        for (const [path, date] of Object.entries(dates)) assert.equal(String(statSync(join(output, path), { bigint: true }).mtimeNs), date);
      } else assert.match(stderr, /cache miss:/);
      results.push({ ...result, scenario, pair, compiler: name, state: state.name, frontend: frontendResult,
        objectsBytes: actual.bytes, generatedGoBytes: actual.outputs.reduce((sum, file) => sum + file.object.bytes, 0), byteExact: true });
      save(join(measurement, "results.json"), results);
    }
  }
  save(join(measurement, "business-oracles.json"), Object.fromEntries(expected));
  console.log(`${results.length} backend samples completed and byte-exact`);
}
