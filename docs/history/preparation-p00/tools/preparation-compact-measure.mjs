import assert from "node:assert/strict";
import { createHash } from "node:crypto";
import { constants, copyFileSync, existsSync, mkdirSync, readFileSync, readdirSync, statSync, unlinkSync, writeFileSync } from "node:fs";
import { dirname, join, resolve } from "node:path";
import { fileURLToPath } from "node:url";
import { observe } from "./preparation-observe.mjs";

const [operation, directory] = process.argv.slice(2);
assert(directory && ["project", "packs"].includes(operation), "Usage: node tools/preparation-compact-measure.mjs project|packs STUDY");
const study = resolve(directory), repository = fileURLToPath(new URL("../", import.meta.url));
const provenance = JSON.parse(readFileSync(join(study, "provenance.json")));
const parent = provenance.parent;
const hash = bytes => createHash("sha256").update(bytes).digest("hex");
const load = path => JSON.parse(readFileSync(path, "utf8"));
const save = (path, value) => { mkdirSync(dirname(path), { recursive: true }); writeFileSync(path, JSON.stringify(value, null, 2) + "\n"); };
for (const [name, digest] of Object.entries(provenance.inputs)) assert.equal(hash(readFileSync(join(parent, name))), digest);
const destination = join(study, operation);
assert(!existsSync(destination), "Preserve earlier study measurements");
mkdirSync(join(destination, "logs"), { recursive: true });
copyFileSync(fileURLToPath(import.meta.url), join(destination, "harness.mjs"));
const env = { ...process.env };
for (const key of Object.keys(env)) if (/^(GOPURS_|P00_|P01_)/.test(key) || ["PPROF", "GOFLAGS", "GOCACHE", "GOGC", "GOMEMLIMIT", "GOMAXPROCS", "NODE_OPTIONS"].includes(key)) delete env[key];
Object.assign(env, { GOWORK: "off", GOCACHE: join(study, "probe-go-cache"), TMPDIR: join(study, "tmp"), GOGC: "off", GOMEMLIMIT: "10GiB" });
const native = join(study, "compact/gopurs/gopurs/bin/gopurs-native");
const bundle = join(study, "compact/gopurs/gopurs/bin/gopurs.js");
const originalProbe = join(parent, "codec/gopurs/gopurs/bin/gopurs-native");
const states = load(join(parent, "measurement/states.json"));
const original = load(join(parent, "codec-measurement/sets.json"));
const commands = [], observations = [];

function command(label, executable, args = [], extra = {}, input, failure = false) {
  const base = join(destination, "logs", label);
  assert(!existsSync(base + ".stdout"));
  const start = process.hrtime.bigint();
  const result = observe("/usr/bin/time", ["-l", executable, ...args], { cwd: destination, env: { ...env, ...extra },
    input: input === undefined ? undefined : JSON.stringify(input), encoding: "utf8", timeout: 1800000, maxBuffer: 64 * 1024 * 1024 }, base + ".resources.jsonl");
  const record = { label, executable, args, extra, seconds: Number(process.hrtime.bigint() - start) / 1e9, status: result.status,
    signal: result.signal, error: result.error?.message, expectedFailure: failure,
    peakRssBytes: Number(result.stderr?.match(/(\d+)\s+maximum resident set size/)?.[1] ?? 0),
    metrics: [...(result.stderr ?? "").matchAll(/^\[gopurs\] (.*): (\d+) ms$/gm)].map(([, phase, ms]) => ({ phase, seconds: Number(ms) / 1000 })) };
  writeFileSync(base + ".stdout", result.stdout ?? ""); writeFileSync(base + ".stderr", result.stderr ?? "");
  if (input !== undefined) save(base + ".request.json", input);
  commands.push(record); save(join(destination, "commands.json"), commands);
  assert.ifError(result.error);
  const trace = readFileSync(base + ".resources.jsonl", "utf8").trim().split("\n").map(line => JSON.parse(line));
  assert(trace.at(-1).completed && !trace.at(-1).contaminated && trace.at(-1).errors.length === 0, "Unusable resource trace: " + label);
  assert.equal(result.status, failure ? 1 : 0, label + ": see retained study logs");
  console.log(`${label}: ${record.seconds.toFixed(3)} s${failure ? " (expected rejection)" : ""}`);
  return { record, stdout: result.stdout };
}

function inventory(root, kind) {
  return load(join(root, kind, "index.json")).map(([name, [sha256, bytes]]) => {
    const source = join(root, kind, hash(name) + ".json"), content = readFileSync(source);
    assert.equal(hash(content), sha256); assert.equal(content.length, bytes);
    return { kind, name, source, object: { sha256, bytes } };
  });
}
const identities = entries => entries.map(({ kind, name, object }) => ({ kind, name, object })).sort((a, b) => a.kind.localeCompare(b.kind) || a.name.localeCompare(b.name));
const context = hash(JSON.stringify({ schema: "whole-preparation-v1", compiler: load(join(parent, "validation/summary.json")).compilerHashes,
  projection: hash(readFileSync(native)),
  options: { packages: true, prepare: 8, pbo: 8, emit: 8, pipeline: true }, inputs: provenance.inputs["provenance.json"] }));
const fence = tast => hash(JSON.stringify([context, Object.entries(tast).sort(([a], [b]) => a.localeCompare(b))]));

if (operation === "project") {
  const sets = {};
  function project(state, repetition, source = join(parent, "codec-measurement/encoded", state + "-1"), javascript = false) {
    const label = `${state}-${repetition}${javascript ? "-js" : ""}`, target = join(destination, label);
    mkdirSync(target);
    const result = command(label + "-project", javascript ? process.execPath : native,
      javascript ? ["--stack-size=65536", "--max-old-space-size=24576", bundle] : [], { P01_ACTION: "project", P00_INPUT: source, P00_OBJECTS: target });
    const entries = inventory(target, "compact");
    observations.push({ state, repetition, action: "project", javascript, ...result.record });
    if (!state.endsWith("-changed")) {
      if (sets[state]) assert.deepEqual(identities(entries), identities(sets[state]));
      else sets[state] = entries;
    } else {
      const expected = new Map(sets[state.replace("-changed", "")].map(entry => [entry.name, entry.object]));
      for (const entry of entries) assert.deepEqual(entry.object, expected.get(entry.name));
    }
    return target;
  }
  for (let pair = 1; pair <= 3; pair++) {
    const target = project("base", pair);
    for (const kind of pair % 2 ? ["contributions", "compact"] : ["compact", "contributions"]) {
      const result = command(`base-${pair}-restore-${kind}`, native, [], { P01_ACTION: kind === "compact" ? "restore" : "restore-full",
        P00_OBJECTS: kind === "compact" ? target : join(parent, "codec-measurement/encoded/base-1") });
      observations.push({ state: "base", pair, action: "restore-" + kind, ...result.record });
    }
  }
  for (const state of states.slice(1)) project(state.name, 1);
  project("base", 1, undefined, true);
  for (const state of states.slice(1)) for (let repetition = 1; repetition <= 3; repetition++) {
    project(state.name + "-changed", repetition, join(parent, "codec-measurement/changed", state.name));
  }
  assert.deepEqual(identities(sets.base), identities(sets.ffi), "FFI-only compact preparation");
  save(join(destination, "sets.json"), sets);
  save(join(destination, "results.json"), observations);
  save(join(destination, "summary.json"), { completed: true, context, scope: "Compact projection under a global fence, not selective inter-build reuse", artifacts: { native: hash(readFileSync(native)), js: hash(readFileSync(bundle)) } });
} else {
  assert(load(join(study, "project/summary.json")).completed);
  const compact = load(join(study, "project/sets.json"));
  const source = join(destination, "store-source"); mkdirSync(source);
  const names = ["contract.go", "store.go", "session.go", "compiler.go", "modules.go", "packages.go"];
  for (const name of names) copyFileSync(join(parent, "tree/gopurs/gopurs/tools/build-cache", name), join(source, name));
  copyFileSync(join(repository, "tools/preparation-probe/pack.go"), join(source, "main.go"));
  const helper = join(destination, "pack-probe");
  command("build-pack-probe", "go", ["build", "-trimpath", "-o", helper, ...names.concat("main.go").map(name => join(source, name))]);
  const datasets = Object.fromEntries(["full", "compact"].map(model => [model, Object.fromEntries(states.map(state => [state.name,
    ["prepared", "metadata", "callers"].flatMap(kind => original[state.name][kind].entries.map(entry => ({ kind, ...entry })))
      .concat(model === "full" ? original[state.name].contributions.entries.map(entry => ({ kind: "contributions", ...entry })) : compact[state.name])]))]));
  function manifest(root) {
    const cache = join(root, ".gopurs-cache"), pointer = load(join(cache, "pack-current.json"));
    const bytes = readFileSync(join(cache, "objects", pointer.sha256)); assert.equal(hash(bytes), pointer.sha256);
    return JSON.parse(bytes);
  }
  function restore(root, target, inputFence, label, entries, miss = false, failure = false) {
    const value = command(label, helper, [], {}, { action: "restore", root, target, fence: inputFence }, failure);
    if (failure) return value.record;
    const result = JSON.parse(value.stdout);
    assert.equal(result.status, miss ? "miss" : "ok");
    if (miss) assert(!existsSync(target));
    else {
      const actual = [...new Set(entries.map(entry => entry.kind))].flatMap(kind => inventory(target, kind));
      assert.deepEqual(identities(actual), identities(entries), label + ": original Wire bytes");
    }
    return { ...value.record, internal: result.seconds };
  }
  function publish(root, layout, state, entries, label, failBeforeCommit = false) {
    const result = command(label, helper, [], {}, { action: "publish", root, layout, fence: fence(state.tast), entries, failBeforeCommit }, failBeforeCommit);
    if (failBeforeCommit) return result.record;
    const cache = join(root, ".gopurs-cache");
    const physicalBytes = readdirSync(join(cache, "objects")).reduce((sum, name) => sum + statSync(join(cache, "objects", name)).size, 0)
      + statSync(join(cache, "pack-current.json")).size;
    return { ...result.record, ...JSON.parse(result.stdout), process: result.record.seconds, physicalBytes };
  }
  function measurePacks() {
    const expected = new Map();
    for (const model of ["full", "compact"]) for (let pair = 1; pair <= 3; pair++) {
      for (const layout of pair % 2 ? ["raw", "packed64"] : ["packed64", "raw"]) {
        const root = join(destination, "stores", `${model}-${pair}-${layout}`);
        const base = publish(root, layout, states[0], datasets[model].base, `${model}-${pair}-${layout}-base`);
        observations.push({ model, pair, layout, state: "base", action: "publish", ...base });
        const cache = join(root, ".gopurs-cache"), index = readFileSync(join(cache, "pack-current.json"));
        const initial = new Set(readdirSync(join(cache, "objects")));
        const lock = statSync(join(cache, "write.lock"), { bigint: true }).ino;
        for (const state of states) {
          const label = `${model}-${pair}-${layout}-${state.name}`;
          if (state.name !== "base") {
            for (const name of readdirSync(join(cache, "objects"))) if (!initial.has(name)) {
              const retained = join(destination, "retired", name); mkdirSync(dirname(retained), { recursive: true });
              if (!existsSync(retained)) copyFileSync(join(cache, "objects", name), retained, constants.COPYFILE_FICLONE);
              unlinkSync(join(cache, "objects", name));
            }
            writeFileSync(join(cache, "pack-current.json"), index);
            assert.equal(statSync(join(cache, "write.lock"), { bigint: true }).ino, lock);
            observations.push({ model, pair, layout, state: state.name, action: "publish", ...publish(root, layout, state, datasets[model][state.name], label) });
          }
          const current = manifest(root), key = `${model}-${layout}-${state.name}`;
          if (expected.has(key)) assert.deepEqual(current, expected.get(key), "Deterministic packed bytes and index");
          else expected.set(key, current);
          save(join(destination, "manifests", label + ".json"), current);
          const target = join(destination, "restored", label); mkdirSync(dirname(target), { recursive: true });
          observations.push({ model, pair, layout, state: state.name, action: "restore", ...restore(root, target, fence(state.tast), label + "-restore", datasets[model][state.name]) });
          if (state.name === "base" && model === "compact" && layout === "packed64" && pair === 1) {
            for (const kind of ["prepared", "metadata", "callers", "compact"]) {
              const compactKind = kind === "compact";
              command("typed-packed-" + kind, compactKind ? native : originalProbe, [], { P01_ACTION: "restore", P00_ACTION: "restore-" + kind, P00_OBJECTS: target });
            }
            command("typed-packed-compact-js", process.execPath, ["--stack-size=65536", "--max-old-space-size=24576", bundle], { P01_ACTION: "restore", P00_OBJECTS: target });
          }
          save(join(destination, "results.json"), observations);
        }
      }
    }
  }
  // Protocol failures precede the long campaign and use a small retained
  // fixture. No failure overwrites a previous success or publishes a partial index.
  const fixture = join(destination, "protocol"), entries = datasets.compact.base.filter(entry => entry.kind === "compact").slice(0, 3);
  publish(fixture, "packed64", states[0], entries, "protocol-seed");
  const cache = join(fixture, ".gopurs-cache"), pointer = readFileSync(join(cache, "pack-current.json"));
  restore(fixture, join(destination, "miss"), fence(states[1].tast), "protocol-wrong-fence", entries, true);
  publish(fixture, "packed64", states[1], entries, "protocol-interrupted", true);
  assert.deepEqual(readFileSync(join(cache, "pack-current.json")), pointer);
  restore(fixture, join(destination, "after-interruption"), fence(states[0].tast), "protocol-after-interruption", entries);
  const group = manifest(fixture).groups[0], path = join(cache, "objects", group.body.sha256), correct = readFileSync(path);
  copyFileSync(path, join(destination, "protocol-original-object"));
  const corrupt = Buffer.from(correct); corrupt[Math.floor(corrupt.length / 2)] ^= 1; writeFileSync(path, corrupt);
  restore(fixture, join(destination, "corrupt-result"), fence(states[0].tast), "protocol-corrupt-pack", entries, false, true);
  copyFileSync(path, join(destination, "protocol-corrupt-object"));
  unlinkSync(path);
  restore(fixture, join(destination, "missing-result"), fence(states[0].tast), "protocol-missing-pack", entries, false, true);
  writeFileSync(path, correct);
  restore(fixture, join(destination, "recovered"), fence(states[0].tast), "protocol-recovered", entries);
  writeFileSync(join(cache, "pack-current.json"), "{\"sha256\":");
  restore(fixture, join(destination, "bad-index-result"), fence(states[0].tast), "protocol-corrupt-index", entries, false, true);
  writeFileSync(join(cache, "pack-current.json"), pointer);
  const fixtureRoot = join(parent, "run-3/workspaces/gopurs-tests-gfsQP6/states");
  const labels = ["initial", "cycle-demand-added", "cycle-demand-replaced", "cycle-demand-retracted", "dictionary-provider-changed", "ffi-only", "all-restored"];
  let previous;
  const proofs = [];
  for (const label of labels) {
    const input = load(join(fixtureRoot, label, "input.json"));
    const expected = previous !== undefined && fence(previous) === fence(input.tast);
    if (previous) restore(fixture, join(destination, "fence-" + label), fence(input.tast), "fence-" + label, entries, !expected);
    publish(fixture, "packed64", input, entries, "fence-publish-" + label);
    proofs.push({ label, fence: fence(input.tast), sameAsPrevious: expected }); previous = input.tast;
  }
  assert.equal(proofs.find(value => value.label === "ffi-only").sameAsPrevious, true);
  assert(proofs.filter(value => value.label !== "ffi-only").every(value => !value.sameAsPrevious));
  save(join(destination, "proofs.json"), proofs);
  measurePacks();
  save(join(destination, "summary.json"), { completed: true, candidate: provenance.candidate, context, measurements: observations.length,
    proof: "Whole-input fence only; no reuse after TAST edits is authorized by compact audit digests", helper: hash(readFileSync(helper)) });
}
