import assert from "node:assert/strict";
import { createHash } from "node:crypto";
import { spawn, spawnSync } from "node:child_process";
import { closeSync, cpSync, existsSync, mkdirSync, openSync, readFileSync, readSync, readdirSync, renameSync, rmSync, statSync, symlinkSync, utimesSync, writeFileSync, writeSync } from "node:fs";
import { delimiter, dirname, join } from "node:path";
import { fileURLToPath } from "node:url";
import { test } from "node:test";
import { corePackages, createWorkspace, prepareFixture } from "./test-workspace.mjs";
import { findTypedCompiler } from "./native-workspace.mjs";
import { packageSourceSections } from "./test-snapshots.mjs";

const root = fileURLToPath(new URL("../", import.meta.url));
const baseEnvironment = { ...process.env, GOWORK: "off", GOPURS_RUST: "0",
  PATH: join(root, "node_modules/.bin") + delimiter + process.env.PATH };
for (const key of ["GOPURS_ALLOC_PROFILE", "GOPURS_FORCE_BUILD", "GOPURS_COMPILER_SNAPSHOT", "PPROF"]) delete baseEnvironment[key];
const hash = bytes => createHash("sha256").update(bytes).digest("hex");
const packages = process.env.GOPURS_GO_PACKAGES === "1";
const sectionsByHash = new Map();
const moduleFiles = (files, module, ffi = false) => {
  const source = module.replaceAll(".", "_") + (ffi ? "_ffi.go" : ".go");
  return Object.fromEntries(Object.entries(files).flatMap(([path, digest]) => {
    if (path.endsWith("/" + source)) return [[path, digest]];
    const section = sectionsByHash.get(digest)?.find(section => section.name === source);
    return section ? [[path, hash(section.content)]] : [];
  }));
};
function generated(output) {
  const files = {};
  function visit(directory, prefix = "") {
    for (const entry of readdirSync(directory, { withFileTypes: true })) {
      if (entry.name.startsWith(".")) continue;
      const name = prefix + entry.name, path = join(directory, entry.name);
      if (entry.isDirectory()) visit(path, name + "/");
      else if (name.endsWith(".go") || name === "go.mod") {
        const data = readFileSync(path), digest = hash(data);
        files[name] = digest;
        if (name.endsWith("/0_gopurs_package.go")) sectionsByHash.set(digest, packageSourceSections(data.toString("utf8")).sections);
      }
    }
  }
  visit(output);
  return files;
}
function timestamps(output, files = generated(output)) {
  return Object.fromEntries(Object.keys(files).map(name => [name, String(statSync(join(output, name), { bigint: true }).mtimeNs)]));
}

// An indexed pack is authoritative even if a loose object with the same SHA
// exists. Corrupt exactly the referenced object's bytes, preserving its neighbors.
function corruptCacheObject(cache, object) {
  const indexPath = join(cache, "packs/index.json");
  const entry = existsSync(indexPath) ? JSON.parse(readFileSync(indexPath)).objects[object.sha256] : undefined;
  if (!entry) {
    const path = join(cache, "objects", object.sha256), bytes = readFileSync(path);
    assert.equal(bytes.length, object.bytes);
    assert.equal(hash(bytes), object.sha256);
    bytes[0] ^= 1;
    writeFileSync(path, bytes);
    return;
  }
  assert.equal(entry.bytes, object.bytes);
  assert.ok(entry.bytes > 0);
  const fd = openSync(join(cache, "packs", entry.pack + ".pack"), "r+");
  try {
    const bytes = Buffer.alloc(entry.bytes);
    assert.equal(readSync(fd, bytes, 0, bytes.length, entry.offset), bytes.length);
    assert.equal(hash(bytes), object.sha256);
    bytes[0] ^= 1;
    assert.equal(writeSync(fd, bytes, 0, 1, entry.offset), 1);
  } finally { closeSync(fd); }
}

test("driver cache transitions equal clean generation in every supported host mode", async t => {
  const workspace = createWorkspace();
  let passed = false;
  t.after(() => {
    if (passed && process.env.GOPURS_TEST_KEEP_WORKSPACE !== "1") rmSync(workspace, { recursive: true, force: true });
    else console.log(`Driver cache workspace retained: ${workspace}`);
  });
  const typedBin = join(workspace, "typed-bin");
  mkdirSync(typedBin);
  symlinkSync(findTypedCompiler(root, process.env.GOPURS_PURS), join(typedBin, "purs"));
  const fixture = prepareFixture(root, workspace, join(root, "tests/passing/FFIIntegerReturns.purs"), 0, corePackages(root));
  const directory = fixture.directory, output = join(directory, "output"), logs = join(directory, "logs");
  writeFileSync(join(directory, "src/Extra.purs"), 'module Extra where\nimport Prelude (Unit)\nimport Effect (Effect)\nimport Effect.Console (log)\nimport Main (returnInt64)\nforeign import echo :: String -> String\nforeignConsumer :: Int -> Int\nforeignConsumer value = returnInt64 value\nmain :: Effect Unit\nmain = log (echo "extra")\n');
  writeFileSync(join(directory, "src/Extra.js"), 'export const echo = value => value;\n');
  writeFileSync(join(directory, "src/Extra.go"), 'func Echo(value string) string { return value }\n');
  const front = spawnSync("spago", ["build", "-q"], { cwd: directory, encoding: "utf8", timeout: 120000,
    env: { ...baseEnvironment, PATH: typedBin + delimiter + baseEnvironment.PATH } });
  assert.equal(front.status, 0, front.stdout + front.stderr);
  const tast = join(workspace, "tast");
  cpSync(output, tast, { recursive: true });
  const moduleCount = readdirSync(tast, { withFileTypes: true }).filter(entry =>
    entry.isDirectory() && existsSync(join(tast, entry.name, "corefn.json"))).length;
  assert.ok(moduleCount > 1);
  const ffiPath = join(directory, "src/Main.go"), originalFfi = readFileSync(ffiPath, "utf8");
  const references = new Map();
  const codecReferences = new Map();
  const keyReferences = new Map();
  let sequence = 0;
  const modes = [
    { name: "js", js: "1", jobs: "1", pipeline: "0" },
    { name: "native-sequential", js: "0", jobs: "1", pipeline: "0" },
    { name: "native-parallel", js: "0", jobs: "8", pipeline: "1" },
    ...(process.env.GOPURS_TEST_RUST === "1" ? [{ name: "rust", js: "0", rust: "1", jobs: "4", pipeline: "1" }] : []),
  ];
  for (const mode of modes.filter(mode => !process.env.GOPURS_TEST_MODES || process.env.GOPURS_TEST_MODES.split(",").includes(mode.name))) await t.test(mode.name, async () => {
    rmSync(output, { recursive: true, force: true });
    cpSync(tast, output, { recursive: true });
    writeFileSync(ffiPath, originalFfi);
    rmSync(join(directory, "ffi"), { recursive: true, force: true });
    const env = { ...baseEnvironment, GOPURS_JS: mode.js, GOPURS_RUST: mode.rust ?? "0", GOPURS_PREPARE_JOBS: mode.jobs,
      GOPURS_PBO_JOBS: mode.jobs, GOPURS_EMIT_JOBS: mode.jobs, GOPURS_PIPELINE: mode.pipeline };
    const compiler = mode.rust && process.env.GOPURS_TEST_RUST_BINARY || process.env.GOPURS_TEST_COMPILER || join(root, "bin/gopurs");
    const invoke = (label, args = [], extra = {}, status = 0) => {
      const result = spawnSync(compiler, args, { cwd: directory, env: { ...env, ...extra }, encoding: "utf8", timeout: 60000, maxBuffer: 8 * 1024 * 1024 });
      writeFileSync(join(logs, `${sequence++}-${mode.name}-${label}.log`), (result.stdout ?? "") + (result.stderr ?? ""));
      assert.ifError(result.error);
      assert.equal(result.status, status, result.stdout + result.stderr);
      return result;
    };
    const counters = result => {
      const matches = [...result.stderr.matchAll(/module cache: optimized=(\d+), reused=(\d+), emitted=(\d+), restored=(\d+)/g)];
      assert.equal(matches.length, 1, "one completed module-cache report");
      const [optimized, reused, emitted, restored] = matches[0].slice(1).map(Number);
      assert.equal(optimized + reused, moduleCount, "every backend is accounted for");
      assert.equal(emitted + restored, moduleCount, "every emitted module is accounted for");
      return { optimized, reused, emitted, restored };
    };
    const fingerprints = result => {
      const matches = [...result.stderr.matchAll(/prepared fingerprints: computed=(\d+), reused=(\d+)/g)];
      assert.equal(matches.length, 1, "one completed prepared-fingerprint report");
      const [computed, reused] = matches[0].slice(1).map(Number);
      assert.equal(computed + reused, moduleCount, "every prepared fingerprint is accounted for");
      return { computed, reused };
    };
    const unchangedBackends = result => {
      const counts = counters(result);
      assert.equal(counts.optimized, 0);
      assert.equal(counts.reused, moduleCount);
      assert.deepEqual(fingerprints(result), { computed: 0, reused: moduleCount });
      assert.match(result.stderr, /pbo module attempts: \d+, codegen: 0\b/);
      return counts;
    };
    const hit = (label, args = []) => {
      const before = timestamps(output), result = invoke(label, args);
      assert.match(result.stderr, /cache hit: .*0 modules optimized\/emitted/);
      assert.doesNotMatch(result.stderr, /load TAST \+ sort|prepare \+ monomorphize|optimize \+ emit|cache publish/);
      assert.deepEqual(timestamps(output), before);
    };
    const compare = (label, args = []) => {
      const actual = generated(output), dates = timestamps(output);
      if (references.has(label)) assert.deepEqual(actual, references.get(label), `${label}: cross-host bytes`);
      else references.set(label, actual);
      const cache = join(output, ".gopurs-cache");
      const pointer = JSON.parse(readFileSync(join(cache, "current.json")));
      const manifest = JSON.parse(readFileSync(join(cache, "objects", pointer.manifest.sha256)));
      const parts = Object.fromEntries(manifest.modules.map(mod => [mod.name, mod.parts]));
      if (codecReferences.has(label)) assert.deepEqual(parts, codecReferences.get(label), `${label}: portable business artifacts`);
      else codecReferences.set(label, parts);
      const keys = Object.fromEntries(manifest.modules.map(mod => [mod.name, [mod.prepared, mod.callerDemands]]));
      if (keyReferences.has(label)) assert.deepEqual(keys, keyReferences.get(label), `${label}: portable prepared keys`);
      else keyReferences.set(label, keys);
      const saved = join(workspace, "saved-output");
      renameSync(output, saved);
      mkdirSync(output);
      // Feed the identical TAST to a fresh generation, with no recipe/journal.
      for (const entry of readdirSync(saved, { withFileTypes: true })) {
        if (entry.isDirectory() && existsSync(join(saved, entry.name, "corefn.json"))) {
          mkdirSync(join(output, entry.name));
          cpSync(join(saved, entry.name, "corefn.json"), join(output, entry.name, "corefn.json"));
        }
      }
      cpSync(join(saved, "go.mod"), join(output, "go.mod"));
      invoke(label + "-clean", args);
      assert.deepEqual(generated(output), actual, `${label}: incremental/clean bytes`);
      rmSync(output, { recursive: true });
      renameSync(saved, output);
      assert.deepEqual(timestamps(output), dates);
    };
    invoke("first");
    compare("first");
    hit("identical");
    const first = generated(output);
    const initialTimes = timestamps(output);
    const forced = invoke("forced", ["--force"]);
    assert.deepEqual(counters(forced), { optimized: moduleCount, reused: 0, emitted: moduleCount, restored: 0 });
    assert.deepEqual(fingerprints(forced), { computed: moduleCount, reused: 0 });
    assert.deepEqual(generated(output), first);
    assert.deepEqual(timestamps(output), initialTimes, "forced generation elides identical writes");

    const goMod = "module gopurs/output\n\ngo 1.22\n// external tidy changes\n";
    writeFileSync(join(output, "go.mod"), goMod);
    const goSum = "example.invalid/external v1.0.0 h1:AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=\n";
    writeFileSync(join(output, "go.sum"), goSum);
    writeFileSync(join(output, "purescript/user.txt"), "external asset\n");
    if (packages) assert.match(invoke("tidy-inputs").stderr, /cache miss: inputs-changed/);
    hit("tidy");
    for (const name of Object.keys(moduleFiles(generated(output), "Main"))) rmSync(join(output, name));
    writeFileSync(join(output, "gopurs_runtime/runtime.go"), "corrupt bytes");
    invoke("repair");
    assert.equal(readFileSync(join(output, "go.mod"), "utf8"), goMod);
    assert.equal(readFileSync(join(output, "go.sum"), "utf8"), goSum);
    assert.equal(readFileSync(join(output, "purescript/user.txt"), "utf8"), "external asset\n");
    compare("repair");

    const oldFfiTime = statSync(ffiPath);
    writeFileSync(ffiPath, originalFfi.replace("return value", "return value + 0"));
    utimesSync(ffiPath, oldFfiTime.atime, oldFfiTime.mtime);
    const ffiEdit = invoke("ffi-edit");
    assert.match(ffiEdit.stderr, /cache miss: inputs-changed/);
    assert.deepEqual(unchangedBackends(ffiEdit),
      { optimized: 0, reused: moduleCount, emitted: 1, restored: moduleCount - 1 });
    // Changed FFI invalidates the header-only shortcut. The ordinary path
    // must still reuse every backend; compare below checks the full output.
    assert.notDeepEqual(moduleFiles(generated(output), "Main", true), moduleFiles(first, "Main", true));
    compare("ffi-edit");
    hit("ffi-identical");
    // A Go signature can change the native worker ABI while the PureScript
    // signature and every TAST stay identical. The backend-reuse path must still
    // parse this FFI and invalidate emission according to the new signatures.
    const changedSignature = originalFfi.replace("func ReturnInt64(value int64) int64", "func ReturnInt64(value int64) any");
    assert.notEqual(changedSignature, originalFfi);
    writeFileSync(ffiPath, changedSignature);
    const ffiSignature = invoke("ffi-signature");
    assert.ok(unchangedBackends(ffiSignature).emitted >= 2, "the FFI owner and its consumer must be emitted");
    assert.notDeepEqual(moduleFiles(generated(output), "Main"), moduleFiles(first, "Main"));
    assert.notDeepEqual(moduleFiles(generated(output), "Extra"), moduleFiles(first, "Extra"), "foreign worker ABI invalidates its cross-module consumer");
    compare("ffi-signature");
    const externalSum = readFileSync(join(output, "go.sum"));
    rmSync(join(output, "go.sum"));
    const signatureExecution = spawnSync("go", ["run", "./Main/main/main.go"], { cwd: output, env, encoding: "utf8", timeout: 120000 });
    writeFileSync(join(output, "go.sum"), externalSum);
    assert.equal(signatureExecution.status, 0, signatureExecution.stdout + signatureExecution.stderr);
    assert.equal(signatureExecution.stdout.trim(), "Done");
    const cacheDirectory = join(output, ".gopurs-cache");
    const cachedPointer = JSON.parse(readFileSync(join(cacheDirectory, "current.json")));
    const cachedManifest = JSON.parse(readFileSync(join(cacheDirectory, "objects", cachedPointer.manifest.sha256)));
    const mainBody = cachedManifest.modules.find(mod => mod.name === "Main").parts[0].object;
    corruptCacheObject(cacheDirectory, mainBody);
    const corruptBackend = invoke("corrupt-cached-backend");
    assert.doesNotMatch(corruptBackend.stderr, /optimizer not started/);
    assert.doesNotMatch(corruptBackend.stderr, /cache hit:/);
    assert.ok(counters(corruptBackend).optimized > 0, "corrupt backend must be rebuilt");
    compare("corrupt-cached-backend");
    if (packages) {
      const pointer = JSON.parse(readFileSync(join(cacheDirectory, "current.json")));
      const manifest = JSON.parse(readFileSync(join(cacheDirectory, "objects", pointer.manifest.sha256)));
      const source = manifest.modules.find(mod => mod.name === "Main").sources[0];
      corruptCacheObject(cacheDirectory, source.object);
      const repaired = invoke("corrupt-cached-source");
      assert.doesNotMatch(repaired.stderr, /optimizer not started/);
      assert.doesNotMatch(repaired.stderr, /cache hit:/);
      assert.ok(counters(repaired).optimized > 0, "corrupt source must be rebuilt");
      compare("corrupt-cached-source");
    }
    const beforeTastEdit = generated(output);
    const oldExtraTime = statSync(join(output, "Extra/corefn.json"));
    const extraText = readFileSync(join(output, "Extra/corefn.json"), "utf8");
    assert.ok(extraText.includes('"extra"'));
    writeFileSync(join(output, "Extra/corefn.json"), extraText.replaceAll('"extra"', '"edited"'));
    utimesSync(join(output, "Extra/corefn.json"), oldExtraTime.atime, oldExtraTime.mtime);
    const tastEdit = invoke("tast-edit");
    assert.match(tastEdit.stderr, /cache miss: inputs-changed/);
    // A committed preparation proof can retain the unchanged fingerprints.
    // The edited module must still be fingerprinted and its output regenerated.
    const tastFingerprints = fingerprints(tastEdit), tastCounts = counters(tastEdit);
    assert.ok(tastFingerprints.computed > 0, "the edited TAST must be fingerprinted");
    assert.ok(tastCounts.optimized > 0 && tastCounts.reused > 0);
    assert.ok(tastCounts.emitted > 0 && tastCounts.restored > 0);
    assert.notDeepEqual(moduleFiles(generated(output), "Extra"), moduleFiles(beforeTastEdit, "Extra"));
    compare("tast-edit");
    const beforeRemoval = generated(output);
    rmSync(join(output, "Extra/corefn.json"));
    invoke("module-removed");
    assert.deepEqual(moduleFiles(generated(output), "Extra"), {});
    assert.deepEqual(moduleFiles(generated(output), "Extra", true), {});
    assert.ok(!existsSync(join(output, "Extra/main/main.go")));
    compare("module-removed");
    cpSync(join(tast, "Extra/corefn.json"), join(output, "Extra/corefn.json"));
    const removedPath = Object.keys(beforeRemoval).find(path => !existsSync(join(output, path)));
    assert(removedPath, "removal must release an owned output");
    const external = join(output, removedPath);
    mkdirSync(dirname(external), { recursive: true });
    writeFileSync(external, "// external bytes, after removal from ownership\n");
    assert.match(invoke("new-output-collision", [], {}, 1).stderr, /unowned output differs/);
    assert.equal(readFileSync(external, "utf8"), "// external bytes, after removal from ownership\n");
    rmSync(external);
    invoke("module-added");
    compare("module-added");
    const entrySelected = invoke("entry-selected", ["--main", "Extra"]);
    assert.match(entrySelected.stderr, /cache miss: inputs-changed/);
    // Selecting an entry point changes the optimizer context, not the TAST.
    assert.deepEqual(fingerprints(entrySelected), { computed: 0, reused: moduleCount });
    assert.deepEqual(counters(entrySelected), { optimized: moduleCount, reused: 0, emitted: moduleCount, restored: 0 });
    assert.ok(!existsSync(join(output, "Main/main/main.go")));
    compare("entry-selected", ["--main", "Extra"]);

    writeFileSync(ffiPath, originalFfi);
    mkdirSync(join(directory, "ffi"));
    renameSync(ffiPath, join(directory, "ffi/Main.go"));
    const ffiArgs = ["--main", "Main", "--ffi", "ffi"];
    invoke("ffi-fallback", ffiArgs);
    compare("ffi-fallback", ffiArgs);
    hit("ffi-fallback-hit", ffiArgs);
    writeFileSync(ffiPath, originalFfi.replace("return value", "return value + 0"));
    assert.match(invoke("ffi-priority", ffiArgs).stderr, /cache miss: inputs-changed/);
    compare("ffi-priority", ffiArgs);
    rmSync(ffiPath);
    rmSync(join(directory, "ffi/Main.go"));
    invoke("ffi-absent", ffiArgs);
    compare("ffi-absent", ffiArgs);
    writeFileSync(ffiPath, originalFfi);
    invoke("ffi-added", ffiArgs);
    compare("ffi-added", ffiArgs);

    const beforeFailure = generated(output), beforeFailureTimes = timestamps(output);
    const pointerPath = join(output, ".gopurs-cache/current.json"), pointer = readFileSync(pointerPath);
    writeFileSync(ffiPath, "func Broken(");
    assert.match(invoke("failed", ffiArgs, {}, 1).stderr, /FFI error in module Main/);
    assert.deepEqual(generated(output), beforeFailure);
    assert.deepEqual(timestamps(output), beforeFailureTimes);
    assert.deepEqual(readFileSync(pointerPath), pointer);
    writeFileSync(ffiPath, originalFfi);
    hit("failure-recovery", ffiArgs);

    const extraTast = readFileSync(join(output, "Extra/corefn.json"));
    writeFileSync(join(output, "Extra/corefn.json"), "invalid TAST");
    assert.match(invoke("invalid-tast", ffiArgs, {}, 1).stderr, /Failed to decode/);
    assert.deepEqual(generated(output), beforeFailure);
    assert.deepEqual(readFileSync(pointerPath), pointer);
    writeFileSync(join(output, "Extra/corefn.json"), extraTast);
    hit("tast-error-recovery", ffiArgs);

    // Stop a real driver after it acquired its private staging, before publish.
    const interrupted = await new Promise((resolve, reject) => {
      const child = spawn(compiler, [...ffiArgs, "--force"], { cwd: directory, env, stdio: ["ignore", "pipe", "pipe"] });
      let text = "", killed = false;
      const timer = setTimeout(() => { child.kill("SIGKILL"); reject(new Error("driver did not reach cache miss")); }, 30000);
      child.stderr.on("data", chunk => {
        text += chunk;
        if (!killed && text.includes("cache miss:")) { killed = true; child.kill("SIGTERM"); }
      });
      child.on("error", reject);
      child.on("close", (code, signal) => { clearTimeout(timer); resolve({ text, code, signal, killed }); });
    });
    writeFileSync(join(logs, `${mode.name}-interrupted.json`), JSON.stringify(interrupted, null, 2));
    assert.equal(interrupted.killed, true);
    assert.notEqual(interrupted.code, 0);
    assert.deepEqual(generated(output), beforeFailure);
    assert.deepEqual(readFileSync(pointerPath), pointer);
    hit("interruption-recovery", ffiArgs);

    const concurrent = await Promise.all([1, 2].map(index => new Promise((resolve, reject) => {
      const child = spawn(compiler, [...ffiArgs, "--force"], { cwd: directory, env, stdio: ["ignore", "pipe", "pipe"] });
      let text = "";
      child.stdout.on("data", chunk => text += chunk);
      child.stderr.on("data", chunk => text += chunk);
      child.on("error", reject);
      child.on("close", (code, signal) => {
        writeFileSync(join(logs, `${mode.name}-concurrent-${index}.log`), text);
        resolve({ code, signal, text });
      });
    })));
    for (const result of concurrent) assert.equal(result.code, 0, result.text);
    assert.deepEqual(generated(output), beforeFailure);
    hit("concurrent-recovery", ffiArgs);

    const profile = join(workspace, `${mode.name}.alloc`);
    assert.match(invoke("profile", ffiArgs, { GOPURS_ALLOC_PROFILE: profile }).stderr, /cache miss: forced/);
    if (mode.js === "0" && !mode.rust) assert.ok(existsSync(profile));
    // Disposable metadata is repairable without losing the ownership journal.
    writeFileSync(join(output, ".gopurs-cache/recipe.json"), "broken");
    invoke("recipe-recovery", ffiArgs);
    writeFileSync(pointerPath, "broken");
    invoke("manifest-recovery", ffiArgs);
    rmSync(join(output, "go.mod"));
    invoke("module-file-recovery", ffiArgs);
    hit("final-hit", ffiArgs);
    assert.equal(readFileSync(join(output, "purescript/user.txt"), "utf8"), "external asset\n");
    // The source-map-independent program remains executable after all repairs.
    rmSync(join(output, "go.sum"));
    const execution = spawnSync("go", ["run", "./main/main.go"], { cwd: output, env, encoding: "utf8", timeout: 120000 });
    assert.equal(execution.status, 0, execution.stdout + execution.stderr);
    assert.equal(execution.stdout.trim(), "Done");
  });
  writeFileSync(join(workspace, "cross-host-inventories.json"), JSON.stringify(Object.fromEntries(references), null, 2));
  writeFileSync(join(workspace, "cross-host-codecs.json"), JSON.stringify(Object.fromEntries(codecReferences), null, 2));
  writeFileSync(join(workspace, "cross-host-keys.json"), JSON.stringify(Object.fromEntries(keyReferences), null, 2));
  passed = true;
});
