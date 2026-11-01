// Handwritten FFI for Gopurs.Host.Snapshot; keep in sync with the .purs foreign imports.
import { execFileSync, spawn } from "node:child_process";
import { existsSync, mkdirSync, mkdtempSync, readFileSync, realpathSync, rmSync, writeFileSync } from "node:fs";
import { tmpdir } from "node:os";
import { dirname, join, resolve } from "node:path";
import { fileURLToPath } from "node:url";

const bundle = fileURLToPath(import.meta.url);
const variable = "GOPURS_COMPILER_SNAPSHOT";

export const withSnapshot = run => () => {
  if (process.env[variable]) {
    if (bundle !== join(process.env[variable], "bin/gopurs.js")) throw new Error("Compiler is not executing its frozen bundle");
    run();
    return;
  }
  const tools = [
    join(dirname(bundle), "../tools"),
    join(dirname(bundle), "../../tools"),
    join(dirname(bundle), "../../../tools"),
  ].map(path => resolve(path)).find(path => existsSync(join(path, "build-cache.mjs")));
  if (!tools) throw new Error("Cannot locate compiler resources");
  const helper = execFileSync(process.execPath, [join(tools, "build-cache.mjs"), "--helper"], { encoding: "utf8" }).trim();
  const packages = process.env.GOPURS_GO_PACKAGES === "1"
    ? execFileSync(process.execPath, [join(tools, "go-packages.mjs")], { encoding: "utf8" }).trim() : null;
  const root = realpathSync(mkdtempSync(join(tmpdir(), "gopurs-compiler-")));
  const resources = [];
  const copy = (name, source, path, mode = 0o400) => {
    const target = join(root, path);
    mkdirSync(dirname(target), { recursive: true });
    writeFileSync(target, readFileSync(source), { mode });
    resources.push({ name, kind: "file", path });
  };
  try {
    copy("compiler", bundle, "bin/gopurs.js");
    copy("cache-helper", helper, "tools/build-cache-helper", 0o500);
    if (packages) copy("go-packages-helper", packages, "tools/go-packages-helper", 0o500);
    for (const name of ["ffi-runner.mjs", "ffi_gen.wasm", "wasm_exec.js"]) copy(name, join(tools, name), "tools/" + name);
    writeFileSync(join(root, "package.json"), '{"type":"module"}\n', { mode: 0o400 });
    resources.push({ name: "module-format", kind: "file", path: "package.json" });
    writeFileSync(join(root, "compiler.json"), JSON.stringify({ schema: 1, host: "js", resources }), { mode: 0o400 });
    const child = spawn(process.execPath, [...process.execArgv, join(root, "bin/gopurs.js"), ...process.argv.slice(2)], {
      stdio: "inherit", env: { ...process.env, [variable]: root, GOPURS_PURMETA_DIRECTORY: join(root, ".purmeta") },
    });
    const signals = ["SIGINT", "SIGTERM", "SIGHUP"];
    const forward = Object.fromEntries(signals.map(signal => [signal, () => child.kill(signal)]));
    for (const signal of signals) process.on(signal, forward[signal]);
    child.once("error", error => { rmSync(root, { recursive: true, force: true }); console.error(error); process.exitCode = 1; });
    child.once("exit", (status, signal) => {
      for (const name of signals) process.off(name, forward[name]);
      rmSync(root, { recursive: true, force: true });
      if (signal) process.kill(process.pid, signal);
      else process.exitCode = status ?? 1;
    });
  } catch (error) { rmSync(root, { recursive: true, force: true }); throw error; }
};

export const release = () => {
  const root = process.env[variable];
  if (root && bundle === join(root, "bin/gopurs.js")) rmSync(root, { recursive: true, force: true });
};
