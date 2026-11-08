import assert from "node:assert/strict";
import { createHash } from "node:crypto";
import { spawnSync } from "node:child_process";
import { constants, copyFileSync, cpSync, existsSync, mkdirSync, readFileSync, readdirSync, realpathSync, statSync, symlinkSync, unlinkSync, writeFileSync } from "node:fs";
import { dirname, join, resolve } from "node:path";
import { fileURLToPath } from "node:url";

assert(process.argv.length === 3, "Usage: node --expose-gc tools/preparation-validate.mjs CAMPAIGN");
const repository = fileURLToPath(new URL("../", import.meta.url));
const campaign = realpathSync(resolve(process.argv[2]));
const destination = join(campaign, "validation");
assert(!existsSync(destination), "Preserve previous validation attempts");
mkdirSync(destination); mkdirSync(join(destination, "logs")); mkdirSync(join(destination, "harness"));
const compiler = join(campaign, "tree/gopurs/gopurs/bin/gopurs");
const hash = bytes => createHash("sha256").update(bytes).digest("hex");
const load = path => JSON.parse(readFileSync(path));
const save = (path, value) => { mkdirSync(dirname(path), { recursive: true }); writeFileSync(path, JSON.stringify(value, null, 2) + "\n"); };
const states = load(join(campaign, "measurement/states.json"));
const expected = load(join(campaign, "measurement/business-oracles.json"));
const commands = [], comparisons = [];
const env = { ...process.env };
for (const key of Object.keys(env)) if (key.startsWith("GOPURS_") || ["PPROF", "GOFLAGS", "GOCACHE", "GOGC", "GOMEMLIMIT", "GOMAXPROCS", "NODE_OPTIONS"].includes(key)) delete env[key];
Object.assign(env, { TMPDIR: join(campaign, "tmp"), GOWORK: "off", GOPURS_GO_PACKAGES: "1", GOGC: "off", GOMEMLIMIT: "10GiB" });
const compilerHashes = Object.fromEntries(["gopurs", "gopurs-native", "gopurs.js", "gopurs-rust"].map(name => [name, hash(readFileSync(join(dirname(compiler), name)))]));

function command(label, executable, args, cwd, extra = {}) {
  global.gc?.();
  const started = process.hrtime.bigint();
  const result = spawnSync(executable, args, { cwd, env: { ...env, ...extra }, encoding: "utf8", timeout: 1800000, maxBuffer: 64 * 1024 * 1024 });
  writeFileSync(join(destination, "logs", label + ".stdout"), result.stdout ?? "");
  writeFileSync(join(destination, "logs", label + ".stderr"), result.stderr ?? "");
  commands.push({ label, executable, args, cwd, extra, seconds: Number(process.hrtime.bigint() - started) / 1e9,
    status: result.status, signal: result.signal, error: result.error?.message });
  save(join(destination, "commands.json"), commands);
  assert.ifError(result.error);
  assert.equal(result.status, 0, label + ": see retained validation logs");
  assert.doesNotMatch(result.stderr, /ERROR: Rebox missing!/);
  console.log(label + ": passed");
  return result;
}

function files(directory, prefix = "") {
  return readdirSync(directory, { withFileTypes: true }).flatMap(entry => entry.isDirectory()
    ? files(join(directory, entry.name), prefix + entry.name + "/") : [prefix + entry.name]);
}
function copy(source, target) {
  assert(!existsSync(target)); mkdirSync(dirname(target), { recursive: true });
  cpSync(source, target, { recursive: true, dereference: true, mode: constants.COPYFILE_FICLONE,
    filter: path => !path.endsWith("/write.lock") });
}
function install(app, state) {
  for (const [path, text] of Object.entries(state.sources)) writeFileSync(join(app, path), text);
  for (const [name, digest] of Object.entries(state.tast)) {
    const target = join(app, "output", name, "corefn.json"); mkdirSync(dirname(target), { recursive: true });
    copyFileSync(join(campaign, "inputs", state.name, "tast", name, "corefn.json"), target);
    assert.equal(hash(readFileSync(target)), digest);
  }
}
function business(app, label) {
  const output = join(app, "output"), cache = join(output, ".gopurs-cache"), pointer = load(join(cache, "current.json"));
  const object = digest => { const value = readFileSync(join(cache, "objects", digest)); assert.equal(hash(value), digest); return value; };
  const manifest = JSON.parse(object(pointer.manifest.sha256));
  assert.equal(manifest.snapshot.output, realpathSync(output));
  for (const file of manifest.outputs) assert.equal(hash(readFileSync(join(output, file.path))), file.object.sha256);
  for (const mod of manifest.modules) for (const part of [...mod.parts, ...(mod.sources ?? [])]) object(part.object.sha256);
  const cas = join(destination, "objects"); mkdirSync(cas, { recursive: true });
  for (const name of readdirSync(join(cache, "objects"))) if (!existsSync(join(cas, name))) copyFileSync(join(cache, "objects", name), join(cas, name), constants.COPYFILE_FICLONE);
  save(join(destination, "manifests", label + ".json"), manifest);
  return { outputs: manifest.outputs, modules: manifest.modules.map(mod => ({ name: mod.name, prepared: mod.prepared,
    callerDemands: mod.callerDemands, parts: mod.parts, sources: mod.sources })) };
}
function restore(app, seed) {
  const output = join(app, "output"), lockName = ".gopurs-cache/write.lock";
  const inode = statSync(join(output, lockName), { bigint: true }).ino, source = new Set(files(seed));
  for (const name of files(output)) if (name !== lockName && !source.has(name)) unlinkSync(join(output, name));
  for (const name of source) if (name !== lockName) {
    mkdirSync(dirname(join(output, name)), { recursive: true });
    copyFileSync(join(seed, name), join(output, name), constants.COPYFILE_FICLONE);
  }
  assert.equal(statSync(join(output, lockName), { bigint: true }).ino, inode);
}

const image = "sha256:41601ad5efc57caa9e7ff557ecfd47c2304eae185de9af36eaaa0d617c5fa45c";
copyFileSync(resolve(repository, "../../b8x/bin/_go-run"), join(destination, "harness/go-run"));
const moduleCache = command("go-module-cache", "go", ["env", "GOMODCACHE"], destination).stdout.trim();
copy(moduleCache, join(destination, "runner-modules"));
save(join(destination, "runner.json"), { image, limits: { GOMAXPROCS: 2, packages: 1, GOMEMLIMIT: "12GiB" },
  helper: hash(readFileSync(join(destination, "harness/go-run"))) });
function execute(app, state, actual) {
  const symbols = ["Get_Util_Type_String_String_caseToCamel", "Util_Type_String_String_RemoveAccents", "Get_Inter_Cli_Ping_Main_main"];
  const owners = symbols.map(name => {
    const found = actual.outputs.filter(file => file.path.endsWith("/0_gopurs_package.go") && readFileSync(join(app, "output", file.path), "utf8").includes("func " + name + "("));
    assert.equal(found.length, 1, "Unique oracle symbol: " + name);
    return dirname(found[0].path);
  });
  const packages = [...new Set(owners)], alias = owner => "p" + packages.indexOf(owner);
  const source = `package main\nimport (\n"fmt"\nr "gopurs/output/gopurs_runtime"\n_ "gopurs/output/purescript/init_all"\n${packages.map((path, index) => `p${index} "gopurs/output/${path}"`).join("\n")}\n)\nfunc main() {\nfmt.Printf("camel=%q\\n", r.Apply(${alias(owners[0])}.${symbols[0]}(), r.Str("")).StrVal())\nfmt.Printf("accent=%q\\n", ${alias(owners[1])}.${symbols[1]}("é"))\nr.Apply(${alias(owners[2])}.${symbols[2]}(), r.Value{})\nr.EventLoopWait()\n}\n`;
  const entry = join(destination, "oracle", state.name + ".go"); mkdirSync(dirname(entry), { recursive: true }); writeFileSync(entry, source);
  const relativeApp = app.slice(destination.length);
  const args = ["run", "--rm", "--network=none", "-v", destination + ":/work", "-w", "/work" + relativeApp + "/output",
    "-e", "PATH=/usr/local/go/bin:/usr/local/bin:/usr/bin:/bin", "-e", "GOTOOLCHAIN=local", "-e", "GOPROXY=off", "-e", "GOSUMDB=off",
    "-e", "ROOT_DIR=/work/runner", "-e", "GOCACHE=/work/runner-cache", "-e", "GOMODCACHE=/work/runner-modules",
    "-e", "B8X_GO_BUILD_PROCS=2", "-e", "B8X_GO_BUILD_PARALLELISM=1", "-e", "B8X_GO_BUILD_MEMORY=12GiB", image,
    "bash", "/work/harness/go-run", "/work" + relativeApp + "/output", "/work/oracle/" + state.name + ".go"];
  const result = command("execute-" + state.name, "docker", args, destination);
  assert.equal(result.stdout, `camel=${JSON.stringify(state.name === "shared" ? "-lot00" : "")}\naccent=${JSON.stringify(state.name === "ffi" ? "e!" : "e")}\n${state.name === "leaf" ? "pong-lot00" : "pong"}\n`);
  assert.match(result.stderr, /GOMAXPROCS=2, packages=-p 1, GOMEMLIMIT=12GiB/);
  const after = business(app, "executed-" + state.name);
  assert.deepEqual(after, actual);
}

for (const mode of [
  { name: "native-parallel", jobs: "8", pipeline: "1", states },
  { name: "native-sequential", jobs: "1", pipeline: "0", states: states.slice(0, 1) },
  { name: "js", jobs: "1", pipeline: "0", states: states.slice(0, 1) },
  { name: "rust", jobs: "8", pipeline: "1", states: states.slice(0, 1) },
]) {
  const parent = join(destination, mode.name), app = join(parent, "b8x"), output = join(app, "output");
  mkdirSync(output, { recursive: true });
  symlinkSync(join(campaign, "tree/gopurs"), join(parent, "gopurs"));
  for (const part of ["src", "test", ".spago", "spago.yaml", "spago.lock"]) copy(join(campaign, "tree/b8x", part), join(app, part));
  for (const name of ["go.mod", "go.sum"]) copyFileSync(join(campaign, "inputs/original", name), join(output, name));
  const extra = { GOPURS_JS: mode.name === "js" ? "1" : "0", GOPURS_RUST: mode.name === "rust" ? "1" : "0",
    GOPURS_PREPARE_JOBS: mode.jobs, GOPURS_PBO_JOBS: mode.jobs, GOPURS_EMIT_JOBS: mode.jobs, GOPURS_PIPELINE: mode.pipeline,
    GOCACHE: join(parent, "go-cache"), NODE_OPTIONS: "--max-old-space-size=24576" };
  const seed = join(parent, "seed");
  for (const state of mode.states) {
    if (state.name !== "base") restore(app, seed);
    install(app, state);
    const label = mode.name + "-" + state.name;
    command(label, compiler, [], app, extra);
    const actual = business(app, label);
    assert.deepEqual(actual, expected[state.name], label + ": measured native/cross-host business equality");
    const dates = Object.fromEntries(actual.outputs.map(file => [file.path, String(statSync(join(output, file.path), { bigint: true }).mtimeNs)]));
    const hit = command(label + "-hit", compiler, [], app, extra);
    assert.match(hit.stderr, /cache hit:/);
    for (const [path, date] of Object.entries(dates)) assert.equal(String(statSync(join(output, path), { bigint: true }).mtimeNs), date);
    const forced = command(label + "-forced", compiler, ["--force"], app, extra);
    assert.match(forced.stderr, /cache miss: forced/);
    assert.deepEqual(business(app, label + "-forced"), actual, label + ": forced/cached equality");
    for (const [name, digest] of Object.entries(state.tast)) assert.equal(hash(readFileSync(join(output, name, "corefn.json"))), digest);
    if (state.name === "base") copy(output, seed);
    if (mode.name === "native-parallel") execute(app, state, actual);
    comparisons.push({ host: mode.name, state: state.name, sameTast: true, byteExact: true, forced: true, hit: true, executed: mode.name === "native-parallel" });
    save(join(destination, "comparisons.json"), comparisons);
  }
}
assert.deepEqual(Object.fromEntries(Object.keys(compilerHashes).map(name => [name, hash(readFileSync(join(dirname(compiler), name)))])), compilerHashes);
save(join(destination, "summary.json"), { completed: true, compilerHashes, comparisons, executionOracles: 4, commands: commands.length,
  scope: "Four-host b8x base; all b8x edits native-parallel; dedicated preparation fixtures cover every transition in every host." });
