import assert from "node:assert/strict";
import { createHash } from "node:crypto";
import { execFileSync } from "node:child_process";
import { copyFileSync, cpSync, existsSync, mkdirSync, readFileSync, readdirSync, statSync, symlinkSync, writeFileSync } from "node:fs";
import { dirname, join, resolve } from "node:path";
import { transformSync } from "esbuild";
import { CommandRunner, Interrupted } from "./command-runner.mjs";

if (process.argv.includes("--help")) {
  console.log("Usage: node tools/cache-context-runtime.mjs config.json\nCreate an audited compiler/runtime catalog for test:cache-context.\nConfig: baseline/alternate { compiler, rust }, nativeOutput, rustOutput, destination.\nUse retained Go/Rust bootstrap outputs matching baseline; destination must be new.");
  process.exit(0);
}
assert.equal(process.argv.length, 3, "expected a runtime-variant configuration file");
const config = JSON.parse(readFileSync(process.argv[2], "utf8"));
for (const key of ["nativeOutput", "rustOutput", "destination"]) {
  assert.equal(typeof config[key], "string", key);
  config[key] = resolve(config[key]);
}
for (const name of ["baseline", "alternate"]) for (const field of ["compiler", "rust"]) {
  assert.equal(typeof config[name]?.[field], "string", `${name}.${field}`);
  config[name][field] = resolve(config[name][field]);
}
const hash = bytes => createHash("sha256").update(bytes).digest("hex");
const fileHash = path => hash(readFileSync(path));
const save = (path, value) => writeFileSync(path, JSON.stringify(value, null, 2) + "\n", { flag: "wx" });
function inventory(directory, prefix = "") {
  return Object.fromEntries(readdirSync(directory, { withFileTypes: true }).flatMap(entry => {
    if (entry.name === "node_modules") return [];
    const path = join(directory, entry.name), name = prefix + entry.name;
    if (entry.isDirectory()) return Object.entries(inventory(path, name + "/"));
    return entry.isFile() ? [[name, fileHash(path)]] : [];
  }));
}
function compilerFiles(compiler) {
  const directory = dirname(dirname(compiler));
  const files = { "package.json": fileHash(join(directory, "package.json")) };
  for (const name of ["src", "tools", "bin"]) Object.assign(files, inventory(join(directory, name), name + "/"));
  return { directory, files };
}
const baseline = compilerFiles(config.baseline.compiler), alternate = compilerFiles(config.alternate.compiler);
const baselineRustHash = fileHash(config.baseline.rust), alternateRustHash = fileHash(config.alternate.rust);
const destination = config.destination;
assert.ok(!existsSync(destination), "destination already exists; retain it and use a new attempt");
mkdirSync(destination, { recursive: true });
save(join(destination, "config.json"), config);
save(join(destination, "baseline-files.json"), baseline.files);
save(join(destination, "alternate-files.json"), alternate.files);
const candidate = join(destination, "runtime-compiler");
mkdirSync(candidate);
for (const name of ["src", "tools", "bin"]) cpSync(join(baseline.directory, name), join(candidate, name), { recursive: true });
copyFileSync(join(baseline.directory, "package.json"), join(candidate, "package.json"));
symlinkSync(join(baseline.directory, "node_modules"), join(candidate, "node_modules"), "dir");
const runtimeSource = readFileSync(join(baseline.directory, "src/Gopurs/Runtime/Source.js"), "utf8");
assert.equal(readFileSync(join(alternate.directory, "src/Gopurs/Runtime/Source.js"), "utf8"), runtimeSource);
const original = JSON.parse(runtimeSource.split("export const runtimeGoCode = ")[1].trim().slice(0, -1));
const marker = "\n// lot05 runtime context transition\n";
assert.ok(!original.includes(marker));
const modified = original + marker;
writeFileSync(join(destination, "runtime-original.go"), original, { flag: "wx" });
writeFileSync(join(destination, "runtime-modified.go"), modified, { flag: "wx" });
function replaceOnce(file, before, after) {
  const text = readFileSync(file, "utf8");
  assert.equal(text.split(before).length, 2, `${file}: expected one embedded runtime`);
  const result = text.replace(before, after);
  assert.equal(result.replace(after, before), text, "substitution must be reversible");
  writeFileSync(file, result);
}
for (const extension of ["go", "js", "rs"]) replaceOnce(join(candidate, `src/Gopurs/Runtime/Source.${extension}`),
  extension === "rs" ? original : JSON.stringify(original), extension === "rs" ? modified : JSON.stringify(modified));
const declaration = source => transformSync("var runtimeGoCode = " + JSON.stringify(source) + ";", { format: "esm", target: "esnext" }).code;
replaceOnce(join(candidate, "bin/gopurs.js"), declaration(original), declaration(modified));
const native = join(destination, "native-output"), rust = join(destination, "rust-output");
const nativeBefore = inventory(config.nativeOutput), rustBefore = inventory(config.rustOutput);
cpSync(config.nativeOutput, native, { recursive: true });
cpSync(config.rustOutput, rust, { recursive: true });
replaceOnce(join(native, "purescript/Gopurs_Runtime_Source_ffi.go"), JSON.stringify(original), JSON.stringify(modified));
replaceOnce(join(rust, "Purs_Gopurs_Runtime_Source/src/lib.rs"), original, modified);
const nativeAfter = inventory(native), rustAfter = inventory(rust);
function changed(before, after) {
  assert.deepEqual(Object.keys(after).sort(), Object.keys(before).sort());
  return Object.keys(before).filter(name => before[name] !== after[name]).sort();
}
assert.deepEqual(changed(nativeBefore, nativeAfter), ["purescript/Gopurs_Runtime_Source_ffi.go"]);
assert.deepEqual(changed(rustBefore, rustAfter), ["Purs_Gopurs_Runtime_Source/src/lib.rs"]);
save(join(destination, "generated-audit.json"), { nativeBefore, nativeAfter, rustBefore, rustAfter });
const commands = new CommandRunner(), records = [];
const env = { ...process.env, GOWORK: "off", PPROF: "0" };
for (const key of ["GOPURS_ALLOC_PROFILE", "GOPURS_FORCE_BUILD", "GOPURS_COMPILER_SNAPSHOT"]) delete env[key];
async function run(label, executable, args, cwd) {
  const log = join(destination, label + ".log"), start = process.hrtime.bigint();
  assert.ok(!existsSync(log));
  const result = await commands.run(executable, args, { cwd, env, log });
  records.push({ label, executable, args, cwd, log, status: result.code, signal: result.signal,
    seconds: Number(process.hrtime.bigint() - start) / 1e9, logSHA256: fileHash(log) });
  writeFileSync(join(destination, "commands.json"), JSON.stringify(records, null, 2) + "\n");
  commands.checkInterrupted();
  assert.equal(result.code, 0, `${label} failed; see ${log}`);
}
try {
  await run("build-native", "go", ["build", "-trimpath", "-o", join(candidate, "bin/gopurs-native"), "./main/main.go"], native);
  const target = join(destination, "rust-target"), linkArgs = [];
  if (process.platform === "darwin") {
    const rustc = process.env.RUSTC ?? "rustc";
    const sysroot = execFileSync(rustc, ["--print", "sysroot"], { env, encoding: "utf8" }).trim();
    const host = execFileSync(rustc, ["-vV"], { env, encoding: "utf8" }).match(/^host: (.+)$/m)?.[1];
    assert.ok(host);
    const linker = join(sysroot, "lib/rustlib", host, "bin/gcc-ld/ld64.lld");
    assert.ok(statSync(linker).isFile());
    linkArgs.push("--bin", "purust_output", "--", "-C", "link-arg=-fuse-ld=" + linker);
  }
  await run("build-rust", "cargo", [linkArgs.length ? "rustc" : "build", "--release", "--target-dir", target,
    "--config", 'profile.release.lto="thin"', "--config", "profile.release.opt-level=3", "--config", "profile.release.debug=false",
    "--manifest-path", join(rust, "Cargo.toml"), ...linkArgs], destination);
  const rustBinary = join(destination, "gopurs-rust-runtime");
  copyFileSync(join(target, "release/purust_output"), rustBinary);
  const current = compilerFiles(join(candidate, "bin/gopurs"));
  assert.deepEqual(changed(baseline.files, current.files), ["bin/gopurs-native", "bin/gopurs.js", "src/Gopurs/Runtime/Source.go", "src/Gopurs/Runtime/Source.js", "src/Gopurs/Runtime/Source.rs"]);
  assert.deepEqual(inventory(config.nativeOutput), nativeBefore);
  assert.deepEqual(inventory(config.rustOutput), rustBefore);
  assert.deepEqual(inventory(native), nativeAfter);
  assert.deepEqual(inventory(rust), rustAfter);
  assert.deepEqual(compilerFiles(config.baseline.compiler).files, baseline.files);
  assert.deepEqual(compilerFiles(config.alternate.compiler).files, alternate.files);
  assert.equal(fileHash(config.baseline.rust), baselineRustHash);
  assert.equal(fileHash(config.alternate.rust), alternateRustHash);
  const describe = (compiler, rust, source, runtime) => ({ compiler, rust, runtimeSHA256: hash(runtime),
    files: { ...Object.fromEntries(Object.entries(source.files).map(([name, digest]) => [join(source.directory, name), digest])), [rust]: fileHash(rust) } });
  const catalog = {
    baseline: describe(config.baseline.compiler, config.baseline.rust, baseline, original),
    alternate: describe(config.alternate.compiler, config.alternate.rust, alternate, original),
    runtime: describe(join(candidate, "bin/gopurs"), rustBinary, current, modified),
  };
  commands.checkInterrupted();
  save(join(destination, "variants.json"), catalog);
  console.log(`Runtime variant catalog: ${join(destination, "variants.json")}`);
} catch (error) {
  console.error(error.message);
  process.exitCode = commands.signal ? new Interrupted(commands.signal).exitCode : 1;
} finally {
  commands.dispose();
  console.log(`Runtime variant workspace retained: ${destination}`);
}
