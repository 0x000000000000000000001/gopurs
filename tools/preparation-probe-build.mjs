import assert from "node:assert/strict";
import { createHash } from "node:crypto";
import { spawnSync } from "node:child_process";
import { constants, copyFileSync, cpSync, existsSync, mkdirSync, readFileSync, symlinkSync, writeFileSync } from "node:fs";
import { join, resolve } from "node:path";
import { fileURLToPath, pathToFileURL } from "node:url";

const root = fileURLToPath(new URL("../", import.meta.url));
assert(process.argv.length >= 3 && process.argv.length <= 4, "Usage: node tools/preparation-probe-build.mjs CAMPAIGN [timings|codecs|compact]");
const campaign = resolve(process.argv[2]);
const kind = process.argv[3] ?? "timings";
assert(["timings", "codecs", "compact"].includes(kind));
const family = kind === "compact" ? "compact" : kind === "codecs" ? "codec" : "diagnostic";
if (kind === "timings" && !existsSync(join(campaign, family))) {
  mkdirSync(join(campaign, family));
  for (const name of ["gopurs", "purescript-backend-optimizer-gopurs"]) cpSync(join(campaign, "tree", name), join(campaign, family, name),
    { recursive: true, dereference: true, mode: constants.COPYFILE_FICLONE });
  const patch = spawnSync("git", ["apply", join(root, "tools/preparation-probe/instrumentation.patch")],
    { cwd: join(campaign, family, "gopurs/gopurs"), encoding: "utf8" });
  assert.ifError(patch.error);
  assert.equal(patch.status, 0, patch.stderr);
}
if (kind === "codecs" || kind === "compact") {
  assert(!existsSync(join(campaign, family)), "Codec probe destination must be new");
  mkdirSync(join(campaign, family));
  for (const name of ["gopurs", "purescript-backend-optimizer-gopurs"]) cpSync(join(campaign, "diagnostic", name), join(campaign, family, name),
    { recursive: true, dereference: true, mode: constants.COPYFILE_FICLONE });
  copyFileSync(join(root, "tools/preparation-probe/Probe.purs"), join(campaign, family, "gopurs/gopurs/src/PreparationProbe.purs"));
  if (kind === "compact") {
    const helpers = join(campaign, family, "gopurs/gopurs/src/PreparationProbe.purs");
    const text = readFileSync(helpers, "utf8");
    assert(text.startsWith("module PreparationProbe (main) where"));
    writeFileSync(helpers, text.replace("module PreparationProbe (main) where", "module PreparationProbe (main, mapJobs, readIndex, restoreSet, writeSet) where"));
    copyFileSync(join(root, "tools/preparation-probe/CompactProbe.purs"), join(campaign, family, "gopurs/gopurs/src/PreparationCompactProbe.purs"));
  }
}
const compiler = join(campaign, family, "gopurs/gopurs");
const bootstrap = join(campaign, family, "bootstrap");
const reference = join(campaign, "tree/gopurs/gopurs");
const logs = join(campaign, family, "build-logs");
assert(existsSync(join(campaign, "native-base-manifest.json")), "Validate the reference seed first");
assert(!existsSync(logs), "Build destination must be new; preserve earlier failures");
mkdirSync(logs, { recursive: true });
mkdirSync(bootstrap);
mkdirSync(join(campaign, family, "prior-artifacts"));
for (const name of ["gopurs.js", "gopurs-native"]) copyFileSync(join(compiler, "bin", name), join(campaign, family, "prior-artifacts", name));

const env = { ...process.env };
for (const key of Object.keys(env)) if (key.startsWith("GOPURS_") || ["PPROF", "GOFLAGS", "GOCACHE", "GOGC", "GOMEMLIMIT", "GOMAXPROCS", "NODE_OPTIONS"].includes(key)) delete env[key];
Object.assign(env, {
  PATH: [join(campaign, "toolchain"), join(root, "node_modules/.bin"), process.env.PATH].join(":"),
  TMPDIR: join(campaign, "tmp"), GOWORK: "off", GOCACHE: join(campaign, family, "go-cache"),
  GOPURS_GO_PACKAGES: "1", GOPURS_PREPARE_JOBS: "8", GOPURS_PBO_JOBS: "8",
  GOPURS_EMIT_JOBS: "8", GOPURS_PIPELINE: "1", NODE_OPTIONS: "--max-old-space-size=24576",
});
const spago = join(campaign, "toolchain/spago.bundle.js");
const commands = [];
function command(label, executable, args, cwd, overrides = {}) {
  const started = process.hrtime.bigint();
  const result = spawnSync(executable, args, { cwd, env: { ...env, ...overrides }, encoding: "utf8", timeout: 1800000, maxBuffer: 64 * 1024 * 1024 });
  writeFileSync(join(logs, label + ".stdout"), result.stdout ?? "");
  writeFileSync(join(logs, label + ".stderr"), result.stderr ?? "");
  commands.push({ label, executable, args, cwd, seconds: Number(process.hrtime.bigint() - started) / 1e9,
    status: result.status, signal: result.signal, error: result.error?.message });
  const recordedEnvironment = Object.fromEntries(Object.entries(env).filter(([key]) => /^(GOPURS_|GO|PPROF|NODE_OPTIONS|TMPDIR|PATH)/.test(key)));
  writeFileSync(join(logs, "commands.json"), JSON.stringify({ environment: recordedEnvironment, commands }, null, 2) + "\n");
  assert.ifError(result.error);
  assert.equal(result.status, 0, `${label} failed; see ${join(logs, label + ".stderr")}`);
  console.log(label + ": passed");
}

const entry = kind === "compact" ? "PreparationCompactProbe" : kind === "codecs" ? "PreparationProbe" : "Main";
command("js-bundle", process.execPath, [spago, "bundle", "--offline", "--monochrome", "--module", entry, "--platform", "node", "--outfile", "bin/gopurs.js", "--bundle-type", "app", "--force"], compiler);
command("preparation-contracts", process.execPath, ["--test", "--test-concurrency=1", "tools/monomorphization.test.mjs", "tools/preparation.test.mjs", "tools/global-types.test.mjs", "tools/rebox-metadata.test.mjs"], compiler);

const { nativeWorkspaceConfig, verifyTypedOutput } = await import(pathToFileURL(join(compiler, "tools/native-workspace.mjs")));
writeFileSync(join(bootstrap, "spago.yaml"), nativeWorkspaceConfig(compiler) + '  backend:\n    cmd: "/usr/bin/true"\n');
symlinkSync(join(compiler, "src"), join(bootstrap, "src"));
command("bootstrap-frontend", process.execPath, [spago, "build", "--offline", "--monochrome"], bootstrap);
const typed = verifyTypedOutput(join(bootstrap, "output"), join(campaign, "toolchain/purs"));
command("native-bridges", process.execPath, [join(compiler, "tools/prepare-native-output.mjs"), join(bootstrap, "output")], compiler);
if (kind !== "timings") command("bootstrap-reference-generator", join(reference, "bin/gopurs"), ["--main", entry], bootstrap, { GOGC: "off", GOMEMLIMIT: "10GiB" });
else command("bootstrap-reference-generator", process.execPath, ["--stack-size=65536", join(reference, "bin/gopurs.js"), "--main", entry], bootstrap);
command("go-build", "go", ["build", "-p", "2", "-trimpath", "-o", join(compiler, "bin/gopurs-native"), "./main/main.go"], join(bootstrap, "output"), { GOMAXPROCS: "8", GOMEMLIMIT: "12GiB" });
const hash = path => createHash("sha256").update(readFileSync(path)).digest("hex");
writeFileSync(join(logs, "summary.json"), JSON.stringify({ kind, entry, typed, artifacts: Object.fromEntries(["gopurs.js", "gopurs-native"].map(name => [name, hash(join(compiler, "bin", name))])),
  sources: Object.fromEntries(["Gopurs/Driver/Prepare.purs", "Gopurs/Monomorphization.purs", ...(kind !== "timings" ? ["PreparationProbe.purs"] : []), ...(kind === "compact" ? ["PreparationCompactProbe.purs"] : [])]
    .map(name => [name, hash(join(compiler, "src", name))])) }, null, 2) + "\n");
console.log(`Private ${kind} probe built and source contracts verified`);
