import assert from "node:assert/strict";
import { createHash } from "node:crypto";
import { spawnSync } from "node:child_process";
import {
  copyFileSync, cpSync, existsSync, mkdirSync, readFileSync, readdirSync,
  realpathSync, statSync, writeFileSync,
} from "node:fs";
import { dirname, join, relative, resolve } from "node:path";
import { fileURLToPath } from "node:url";

const root = fileURLToPath(new URL("../", import.meta.url));
const htdocs = resolve(root, "../..");
const [operation, directory] = process.argv.slice(2);
const b8xRevision = process.argv[4];
assert(directory && ["freeze", "repair-offline-git", "repair-app-inputs", "repair-go-inputs", "frontend", "seed"].includes(operation),
  "Usage: node tools/preparation-baseline.mjs freeze|repair-offline-git|repair-app-inputs|repair-go-inputs|frontend|seed CAMPAIGN");
const campaign = resolve(directory);
const hash = bytes => createHash("sha256").update(bytes).digest("hex");
const load = path => JSON.parse(readFileSync(path, "utf8"));
const save = (path, value) => writeFileSync(path, JSON.stringify(value, null, 2) + "\n");

function inventory(directory, prefix = "") {
  const files = {};
  for (const entry of readdirSync(directory, { withFileTypes: true }).sort((a, b) => a.name.localeCompare(b.name))) {
    const path = join(directory, entry.name), name = prefix + entry.name;
    if (entry.isDirectory()) Object.assign(files, inventory(path, name + "/"));
    else if (entry.isFile()) files[name] = { bytes: statSync(path).size, sha256: hash(readFileSync(path)) };
    else throw new Error("Unfrozen symbolic/special file: " + path);
  }
  return files;
}

function copy(source, destination) {
  assert(!existsSync(destination), "Destination already exists: " + destination);
  mkdirSync(dirname(destination), { recursive: true });
  cpSync(source, destination, { recursive: true, dereference: true });
}

function tastInventory(output, requireTyped = true) {
  const result = {};
  const untyped = [];
  let types = 0;
  for (const name of readdirSync(output).sort()) {
    const file = join(output, name, "corefn.json");
    if (!existsSync(file)) continue;
    const bytes = readFileSync(file), value = JSON.parse(bytes);
    if (!Array.isArray(value.typeTable)) {
      assert(!requireTyped, "Untyped CoreFn: " + file);
      untyped.push(name);
    } else types += value.typeTable.length;
    result[name] = { sha256: hash(bytes), bytes: bytes.length, source: value.modulePath };
  }
  return { modules: Object.keys(result).length, types, untyped, files: result };
}

function checked(command, args, cwd) {
  const result = spawnSync(command, args, { cwd, encoding: "utf8", timeout: 120000 });
  assert.ifError(result.error);
  assert.equal(result.status, 0, result.stderr);
  return result.stdout.trim();
}

function environment() {
  const env = { ...process.env };
  for (const key of Object.keys(env)) {
    if (key.startsWith("GOPURS_") || ["PPROF", "GOFLAGS", "GOCACHE", "GOGC", "GOMEMLIMIT", "GOMAXPROCS", "NODE_OPTIONS"].includes(key)) delete env[key];
  }
  return {
    ...env, PATH: join(campaign, "toolchain") + ":" + process.env.PATH,
    TMPDIR: join(campaign, "tmp"), GOWORK: "off", GOCACHE: join(campaign, "host-go-cache"),
    GOPURS_GO_PACKAGES: "1", GOPURS_PREPARE_JOBS: "8", GOPURS_PBO_JOBS: "8",
    GOPURS_EMIT_JOBS: "8", GOPURS_PIPELINE: "1", GOGC: "off", GOMEMLIMIT: "10GiB",
  };
}

function command(label, executable, args, cwd, env = environment()) {
  const initialLabel = label;
  for (let attempt = 2; existsSync(join(campaign, "logs", label + ".json")); attempt++) label = initialLabel + "-attempt-" + attempt;
  const log = join(campaign, "logs", label);
  assert(!existsSync(log + ".json"), "Command already recorded: " + label);
  const started = process.hrtime.bigint();
  const result = spawnSync(executable, args, { cwd, env, encoding: "utf8", timeout: 1800000, maxBuffer: 64 * 1024 * 1024 });
  const seconds = Number(process.hrtime.bigint() - started) / 1e9;
  writeFileSync(log + ".stdout", result.stdout ?? "");
  writeFileSync(log + ".stderr", result.stderr ?? "");
  const metrics = [...(result.stderr ?? "").matchAll(/^\[gopurs\] (.*): (\d+) ms(?: \(failed\))?$/gm)]
    .map(([, phase, milliseconds]) => ({ phase, milliseconds: Number(milliseconds) }));
  save(log + ".json", { label, executable, args, cwd, seconds, metrics,
    environment: Object.fromEntries(Object.entries(env).filter(([key]) => /^(GOPURS_|GO|PPROF|NODE_OPTIONS|TMPDIR|PATH)/.test(key))),
    status: result.status, signal: result.signal, error: result.error?.message });
  assert.ifError(result.error);
  assert.equal(result.status, 0, `${label} failed; complete output retained in ${log}.stderr`);
  console.log(`${label}: ${seconds.toFixed(3)} s`);
  return result;
}

const tree = join(campaign, "tree");
const app = join(tree, "b8x");
const compiler = join(tree, "gopurs/gopurs");

if (operation === "freeze") {
  assert(!existsSync(campaign), "Campaign must be new");
  mkdirSync(campaign, { recursive: true });
  for (const name of ["logs", "tmp", "toolchain", "inputs/original"]) mkdirSync(join(campaign, name), { recursive: true });
  const repositories = {};
  for (const name of ["gopurs/gopurs", "purescript-backend-optimizer-gopurs", "b8x"]) {
    const source = join(htdocs, name);
    repositories[name] = {
      commit: checked("git", ["rev-parse", "HEAD"], source),
      status: checked("git", ["status", "--short"], source),
    };
  }
  // Copy source inputs by bytes. No package points back to a live checkout.
  for (const name of readdirSync(join(htdocs, "gopurs")).sort()) {
    if (!name.startsWith("gopurs-") && name !== "gopurs") continue;
    const source = join(htdocs, "gopurs", name);
    if (!existsSync(join(source, "spago.yaml"))) continue;
    const target = join(tree, "gopurs", name);
    for (const part of ["src", "spago.yaml", "spago.lock"]) if (existsSync(join(source, part))) copy(join(source, part), join(target, part));
  }
  for (const part of ["src", "spago.yaml", "spago.lock"]) copy(join(htdocs, "purescript-backend-optimizer-gopurs", part), join(tree, "purescript-backend-optimizer-gopurs", part));
  for (const part of ["bin", "tools", "runtime", "package.json"]) copy(join(root, part), join(compiler, part));
  const goRoot = join(htdocs, "b8x/run/bak/go");
  copy(join(goRoot, ".spago"), join(app, ".spago"));
  copy(join(goRoot, "spago.lock"), join(app, "spago.lock"));
  if (b8xRevision) {
    const revision = checked("git", ["rev-parse", b8xRevision + "^{commit}"], join(htdocs, "b8x"));
    const archive = spawnSync("git", ["archive", revision, "src", "test", "run/bak/go/spago.go.yaml"],
      { cwd: join(htdocs, "b8x"), maxBuffer: 64 * 1024 * 1024 });
    assert.ifError(archive.error); assert.equal(archive.status, 0);
    const path = join(campaign, "inputs/b8x-sources.tar"); writeFileSync(path, archive.stdout);
    checked("tar", ["-xf", path, "-C", app], app);
    copy(join(app, "run/bak/go/spago.go.yaml"), join(campaign, "inputs/original/spago.go.yaml"));
    repositories.b8x.sourceRevision = revision;
    repositories.b8x.sourceArchiveSHA256 = hash(archive.stdout);
  } else {
    for (const part of ["src", "test"]) copy(join(htdocs, "b8x", part), join(app, part));
    copy(join(goRoot, "spago.go.yaml"), join(campaign, "inputs/original/spago.go.yaml"));
  }
  const config = readFileSync(join(campaign, "inputs/original/spago.go.yaml"), "utf8")
    .replace(/^    math:\n      path: .*\n/m, "")
    .replace('cmd: "bin/_gopurs"', 'cmd: "/usr/bin/true"');
  assert(config.includes('cmd: "/usr/bin/true"'));
  writeFileSync(join(app, "spago.yaml"), config);
  const liveOutput = join(goRoot, "output");
  for (const name of ["go.mod", "go.sum"]) copy(join(liveOutput, name), join(campaign, "inputs/original", name));
  const original = tastInventory(liveOutput, false);
  for (const name of Object.keys(original.files)) copy(join(liveOutput, name, "corefn.json"), join(campaign, "inputs/original/tast", name, "corefn.json"));
  assert.deepEqual(tastInventory(liveOutput, false), original, "Available frontend output changed while freezing");
  save(join(campaign, "inputs/original/tast.json"), original);
  const typed = realpathSync(checked("which", ["purs"], root));
  copy(typed, join(campaign, "toolchain/purs"));
  const spago = realpathSync(checked("which", ["spago"], root));
  copy(spago, join(campaign, "toolchain/spago.bundle.js"));
  const snapshot = inventory(tree);
  save(join(campaign, "source-inventory.json"), snapshot);
  save(join(campaign, "provenance.json"), {
    schema: 1, created: new Date().toISOString(), repositories,
    historicalArtifactsAvailable: false, baseline: "reconstituted-current-native",
    original, sourceFiles: Object.keys(snapshot).length, toolchain: inventory(join(campaign, "toolchain")),
    tools: { go: checked("go", ["version"], root), node: process.version, purs: checked(typed, ["--version"], root), spago: checked("spago", ["--version"], root) },
  });
  copy(fileURLToPath(import.meta.url), join(campaign, "harness/preparation-baseline.mjs"));
  console.log(`Frozen ${Object.keys(snapshot).length} files; available ${original.modules} CoreFn / ${original.types} types / ${original.untyped.length} untyped in ${campaign}`);
} else if (operation === "repair-offline-git") {
  const lock = load(join(app, "spago.lock"));
  const repairs = [];
  for (const [name, value] of Object.entries(lock.packages)) {
    if (value.type !== "git") continue;
    const source = join(app, ".spago/g", value.url.replaceAll(":", "%c").replaceAll("/", "%s"));
    assert.equal(checked("git", ["rev-parse", "HEAD"], source), value.rev);
    assert.equal(checked("git", ["status", "--porcelain"], source), "");
    const refs = new Set([value.rev, lock.workspace.extra_packages[name].ref]);
    for (const ref of refs) {
      assert(/^[a-z0-9-]+$/.test(ref), "Unsupported offline ref: " + ref);
      const target = join(app, ".spago/p", name, ref);
      if (!existsSync(target)) copy(source, target);
      repairs.push({ name, ref, revision: value.rev, target: relative(campaign, target), files: inventory(join(target, "src")) });
    }
  }
  save(join(campaign, "offline-git-repair.json"), repairs);
  console.log(`Restored ${repairs.length} private Git cache locations from locked revisions`);
} else if (operation === "repair-app-inputs") {
  const provenance = load(join(campaign, "provenance.json"));
  assert.equal(checked("git", ["rev-parse", "HEAD"], join(htdocs, "b8x")), provenance.repositories.b8x.commit);
  assert.equal(checked("git", ["status", "--short"], join(htdocs, "b8x")), provenance.repositories.b8x.status);
  copy(join(htdocs, "b8x/test"), join(app, "test"));
  save(join(campaign, "app-inputs-repair.json"), inventory(join(app, "test")));
  console.log("Private b8x test sources restored from the recorded revision");
} else if (operation === "repair-go-inputs") {
  const recorded = {};
  for (const name of ["go.mod", "go.sum"]) {
    const source = join(htdocs, "b8x/output", name), target = join(app, "output", name);
    copy(source, join(campaign, "inputs/original", name));
    if (existsSync(target)) copy(target, join(campaign, "inputs/pre-go-repair", name));
    copyFileSync(source, target);
    recorded[name] = { bytes: statSync(target).size, sha256: hash(readFileSync(target)) };
  }
  save(join(campaign, "go-inputs-repair.json"), recorded);
  console.log("Private Go module inputs restored and fingerprinted");
} else if (operation === "frontend") {
  assert(existsSync(join(campaign, "provenance.json")));
  command("frontend-base", process.execPath, [join(campaign, "toolchain/spago.bundle.js"), "build", "--offline", "--monochrome"], app);
  const actual = tastInventory(join(app, "output"));
  for (const entry of Object.values(actual.files)) assert(existsSync(resolve(app, entry.source)), "Missing frozen source: " + entry.source);
  const original = load(join(campaign, "inputs/original/tast.json"));
  const changed = [...new Set([...Object.keys(original.files), ...Object.keys(actual.files)])]
    .filter(name => original.files[name]?.sha256 !== actual.files[name]?.sha256);
  save(join(campaign, "frontend-base.json"), { ...actual, changedFromAvailableOutput: changed });
  for (const name of Object.keys(actual.files)) copy(join(app, "output", name, "corefn.json"), join(campaign, "inputs/base/tast", name, "corefn.json"));
  for (const name of ["go.mod", "go.sum"]) if (existsSync(join(campaign, "inputs/original", name))) copy(join(campaign, "inputs/original", name), join(app, "output", name));
  console.log(`Frontend: ${actual.modules} TAST, ${changed.length} changed from available b8x output`);
} else if (operation === "seed") {
  assert(existsSync(join(campaign, "frontend-base.json")));
  command("native-base", join(compiler, "bin/gopurs"), [], app);
  const output = join(app, "output");
  const cache = join(output, ".gopurs-cache");
  const current = load(join(cache, "current.json"));
  const manifestBytes = readFileSync(join(cache, "objects", current.manifest.sha256));
  assert.equal(hash(manifestBytes), current.manifest.sha256);
  const manifest = JSON.parse(manifestBytes);
  for (const entry of manifest.outputs) assert.equal(hash(readFileSync(join(output, entry.path))), entry.object.sha256);
  save(join(campaign, "native-base-manifest.json"), manifest);
  command("native-identical", join(compiler, "bin/gopurs"), [], app);
  assert.match(readFileSync(join(campaign, "logs/native-identical.stderr"), "utf8"), /cache hit:/);
  console.log(`Seed validated: ${manifest.modules.length} modules, ${manifest.outputs.length} owned outputs`);
}
