import assert from "node:assert/strict";
import { createHash } from "node:crypto";
import { observe } from "./preparation-observe.mjs";
import { constants, copyFileSync, existsSync, mkdirSync, readFileSync, readdirSync, realpathSync, unlinkSync, writeFileSync } from "node:fs";
import { dirname, join, resolve } from "node:path";
import { fileURLToPath } from "node:url";

assert(process.argv.length === 3, "Usage: node tools/preparation-codec-measure.mjs CAMPAIGN");
const root = fileURLToPath(new URL("../", import.meta.url));
const campaign = realpathSync(resolve(process.argv[2]));
const destination = join(campaign, "codec-measurement");
assert(!existsSync(destination), "Preserve earlier codec measurements");
assert(existsSync(join(campaign, "codec/build-logs/summary.json")), "Build the private codec probe first");
mkdirSync(destination); mkdirSync(join(destination, "logs"));
const hash = bytes => createHash("sha256").update(bytes).digest("hex");
const load = path => JSON.parse(readFileSync(path));
const save = (path, value) => { mkdirSync(dirname(path), { recursive: true }); writeFileSync(path, JSON.stringify(value, null, 2) + "\n"); };
const env = { ...process.env };
for (const key of Object.keys(env)) if (key.startsWith("GOPURS_") || ["PPROF", "GOFLAGS", "GOCACHE", "GOGC", "GOMEMLIMIT", "GOMAXPROCS", "NODE_OPTIONS"].includes(key)) delete env[key];
Object.assign(env, { GOWORK: "off", GOCACHE: join(destination, "go-cache"), GOGC: "off", GOMEMLIMIT: "10GiB", TMPDIR: join(campaign, "tmp") });
const probe = join(campaign, "codec/gopurs/gopurs/bin/gopurs-native");
const results = [], sets = {}, commands = [];
const businessOracles = load(join(campaign, "measurement/business-oracles.json"));

function command(label, executable, args, extra = {}, input) {
  const started = process.hrtime.bigint();
  const result = observe("/usr/bin/time", ["-l", executable, ...args], {
    cwd: destination, env: { ...env, ...extra }, input, encoding: "utf8", timeout: 1800000, maxBuffer: 64 * 1024 * 1024,
  }, join(destination, "logs", label + ".resources.jsonl"));
  const record = { label, executable, args, extra, seconds: Number(process.hrtime.bigint() - started) / 1e9,
    status: result.status, signal: result.signal, error: result.error?.message,
    peakRssBytes: Number(result.stderr?.match(/(\d+)\s+maximum resident set size/)?.[1] ?? 0),
    metrics: [...(result.stderr ?? "").matchAll(/^\[gopurs\] (.*): (\d+) ms$/gm)].map(([, phase, value]) => ({ phase, seconds: Number(value) / 1000 })) };
  writeFileSync(join(destination, "logs", label + ".stdout"), result.stdout ?? "");
  writeFileSync(join(destination, "logs", label + ".stderr"), result.stderr ?? "");
  commands.push(record); save(join(destination, "commands.json"), commands);
  assert.ifError(result.error);
  assert.equal(result.status, 0, label + ": see retained codec logs");
  console.log(`${label}: ${record.seconds.toFixed(3)} s`);
  return { record, stdout: result.stdout };
}

const storeSource = join(destination, "store-source"); mkdirSync(storeSource);
const storeFiles = ["contract.go", "store.go", "session.go", "compiler.go", "modules.go", "packages.go"];
for (const name of storeFiles) copyFileSync(join(campaign, "tree/gopurs/gopurs/tools/build-cache", name), join(storeSource, name));
copyFileSync(join(root, "tools/preparation-probe/store.go"), join(storeSource, "main.go"));
const storeProbe = join(destination, "store-probe");
command("build-store-probe", "go", ["build", "-trimpath", "-o", storeProbe, ...storeFiles.concat("main.go").map(name => join(storeSource, name))]);

const categories = ["source", "prepared", "metadata", "direct", "callers", "transitive", "contributions"];
const states = load(join(campaign, "measurement/states.json"));
function inspect(directory, state, repetition) {
  const current = {};
  for (const kind of categories) {
    const entries = load(join(directory, kind, "index.json")).map(([name, [sha256, bytes]]) => {
      const path = join(directory, kind, hash(name) + ".json"), value = readFileSync(path);
      assert.equal(hash(value), sha256); assert.equal(value.length, bytes);
      return { name, source: path, object: { sha256, bytes } };
    });
    const base = sets.base?.[kind]?.entries;
    const before = new Map((base ?? []).map(value => [value.name, value.object.sha256]));
    const after = new Map(entries.map(value => [value.name, value.object.sha256]));
    const changed = entries.filter(value => before.get(value.name) !== value.object.sha256);
    current[kind] = { entries, bytes: entries.reduce((sum, value) => sum + value.object.bytes, 0), count: entries.length,
      changed: changed.map(value => value.name), changedBytes: changed.reduce((sum, value) => sum + value.object.bytes, 0),
      removed: [...before.keys()].filter(name => !after.has(name)) };
    if (kind === "prepared") {
      const expected = new Map(businessOracles[state.name].modules.map(mod => [mod.name, mod.prepared]));
      assert.equal(entries.length, expected.size);
      for (const entry of entries) assert.equal(entry.object.sha256, expected.get(entry.name), "Probe/production prepared bytes: " + entry.name);
    }
    if (state.name === "ffi") {
      assert.deepEqual(current[kind].changed, [], "FFI-only preparation is identical: " + kind);
      assert.deepEqual(current[kind].removed, []);
    }
    if (state.name === "base" && repetition > 1) assert.deepEqual(entries.map(({ name, object }) => ({ name, object })), base.map(({ name, object }) => ({ name, object })));
  }
  if (repetition === 1) sets[state.name] = current;
  save(join(directory, "inventory.json"), current);
  return current;
}

function encode(state, repetition) {
  const objects = join(destination, "encoded", state.name + "-" + repetition);
  mkdirSync(objects, { recursive: true });
  for (const action of ["encode-source", "encode-prepared", "encode-metadata", "encode-demands"]) {
    const { record } = command(`${state.name}-${repetition}-${action}`, probe, [], {
      P00_ACTION: action, P00_INPUT: join(campaign, "inputs", state.name, "tast"), P00_COUNT: String(Object.keys(state.tast).length), P00_OBJECTS: objects,
    });
    results.push({ state: state.name, repetition, action, ...record });
  }
  inspect(objects, state, repetition);
  return objects;
}

function restore(objects, repetition) {
  for (const kind of categories) {
    const { record } = command(`base-${repetition}-restore-${kind}`, probe, [], { P00_ACTION: "restore-" + kind, P00_OBJECTS: objects });
    results.push({ state: "base", repetition, action: "restore-" + kind, ...record });
  }
}

let first;
for (let repetition = 1; repetition <= 3; repetition++) {
  if (repetition === 2) restore(first, repetition);
  const objects = encode(states[0], repetition); first ??= objects;
  if (repetition !== 2) restore(objects, repetition);
}
for (const state of states.slice(1)) encode(state, 1);
save(join(destination, "sets.json"), sets);

const changedEncoding = [];
for (const state of states.slice(1)) {
  const subset = join(destination, "changed", state.name);
  for (const kind of categories) {
    const names = new Set(sets[state.name][kind].changed);
    const entries = sets[state.name][kind].entries.filter(entry => names.has(entry.name));
    const directory = join(subset, kind); mkdirSync(directory, { recursive: true });
    for (const entry of entries) copyFileSync(entry.source, join(directory, hash(entry.name) + ".json"), constants.COPYFILE_FICLONE);
    save(join(directory, "index.json"), entries.map(entry => [entry.name, [entry.object.sha256, entry.object.bytes]]));
  }
  for (let repetition = 1; repetition <= 3; repetition++) {
    const target = join(destination, "reencoded", state.name + "-" + repetition); mkdirSync(target, { recursive: true });
    for (const kind of categories) {
      const { record } = command(`${state.name}-${repetition}-changed-${kind}`, probe, [], {
        P00_ACTION: "restore-" + kind, P00_OBJECTS: subset, P00_REENCODE: target,
      });
      changedEncoding.push({ state: state.name, repetition, kind, ...record });
      const entries = load(join(target, kind, "index.json"));
      for (const [name, [digest, bytes]] of entries) {
        const value = readFileSync(join(target, kind, hash(name) + ".json"));
        assert.equal(hash(value), digest); assert.equal(value.length, bytes);
      }
    }
    save(join(destination, "changed-encoding.json"), changedEncoding);
  }
}

const publication = [];
for (const kind of categories) for (let repetition = 1; repetition <= 3; repetition++) {
  const store = join(destination, "stores", `${kind}-${repetition}`);
  const request = { root: store, previous: [], entries: sets.base[kind].entries };
  const base = command(`store-${kind}-${repetition}-base`, storeProbe, [], {}, JSON.stringify(request));
  const initial = JSON.parse(base.stdout);
  publication.push({ kind, repetition, state: "base", seconds: initial.seconds, process: base.record.seconds, bytes: initial.bytes });
  const initialIndex = readFileSync(join(store, ".gopurs-cache/probe-index.json"));
  for (const state of states.slice(1)) {
    // Only the base objects are admitted as prior evidence. Extra immutable CAS
    // objects are removed after retaining their identities and the original
    // encoded bytes, so a previous edition never prewarms a later one.
    const baseDigests = new Set(initial.artifacts.map(value => value.object.sha256));
    const objects = join(store, ".gopurs-cache/objects");
    for (const name of readdirSync(objects)) if (!baseDigests.has(name)) {
      const retained = join(destination, "store-retired", name);
      mkdirSync(dirname(retained), { recursive: true });
      if (!existsSync(retained)) copyFileSync(join(objects, name), retained, constants.COPYFILE_FICLONE);
      unlinkSync(join(objects, name));
    }
    writeFileSync(join(store, ".gopurs-cache/probe-index.json"), initialIndex);
    const result = command(`store-${kind}-${repetition}-${state.name}`, storeProbe, [], {}, JSON.stringify({ root: store, previous: initial.artifacts, entries: sets[state.name][kind].entries }));
    const actual = JSON.parse(result.stdout);
    publication.push({ kind, repetition, state: state.name, seconds: actual.seconds, process: result.record.seconds, bytes: actual.bytes });
  }
  save(join(destination, "publication.json"), publication);
}
save(join(destination, "results.json"), results);
save(join(destination, "summary.json"), { completed: true, repetitions: 3, states: states.map(state => state.name), categories,
  scope: "Codec microbenchmarks; staging uses the frozen production store. No incremental preparation implementation or full transaction proof overhead is included.",
  artifacts: { codec: hash(readFileSync(probe)), store: hash(readFileSync(storeProbe)) } });
