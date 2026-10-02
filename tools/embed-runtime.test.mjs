import assert from "node:assert/strict";
import { execFileSync } from "node:child_process";
import { cpSync, mkdirSync, mkdtempSync, readFileSync, rmSync, statSync, utimesSync, writeFileSync } from "node:fs";
import { tmpdir } from "node:os";
import { join } from "node:path";
import { pathToFileURL } from "node:url";
import { test } from "node:test";
import { runtimeGoCode } from "../src/Gopurs/Runtime.js";

test("both embedded runtimes equal the canonical Go source", () => {
  const source = readFileSync(new URL("../runtime/runtime.go", import.meta.url), "utf8");
  assert.equal(runtimeGoCode, source);
  const native = readFileSync(new URL("../src/Gopurs/Runtime.go", import.meta.url), "utf8");
  assert.equal(JSON.parse(native.match(/^var RuntimeGoCode = (.*)$/m)[1]), source);
});

test("embedding resolves from the script, preserves unchanged timestamps and needs no runtime I/O", async t => {
  const directory = mkdtempSync(join(tmpdir(), "gopurs-embed-runtime-"));
  t.after(() => rmSync(directory, { recursive: true, force: true }));
  for (const name of ["tools", "runtime", "src/Gopurs"]) mkdirSync(join(directory, name), { recursive: true });
  writeFileSync(join(directory, "package.json"), '{"type":"module"}\n');
  cpSync(new URL("./embed-runtime.mjs", import.meta.url), join(directory, "tools/embed-runtime.mjs"));
  const canonical = readFileSync(new URL("../runtime/runtime.go", import.meta.url), "utf8");
  const runtime = join(directory, "runtime/runtime.go");
  writeFileSync(runtime, canonical);
  const run = () => execFileSync(process.execPath, [join(directory, "tools/embed-runtime.mjs")], {
    cwd: join(directory, "src"), stdio: "pipe",
  });
  run();
  const generated = ["js", "go"].map(extension => join(directory, `src/Gopurs/Runtime.${extension}`));
  for (const path of generated) utimesSync(path, 1, 1);
  run();
  for (const path of generated) assert.equal(statSync(path).mtimeMs, 1000);
  const changed = canonical + "\n// embedded quote: \" ; slash: \\\n";
  writeFileSync(runtime, changed);
  run();
  for (const path of generated) assert.notEqual(statSync(path).mtimeMs, 1000);
  rmSync(join(directory, "runtime"), { recursive: true });
  assert.equal((await import(pathToFileURL(generated[0]).href)).runtimeGoCode, changed);
  assert.equal(JSON.parse(readFileSync(generated[1], "utf8").match(/^var RuntimeGoCode = (.*)$/m)[1]), changed);
});
