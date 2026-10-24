import assert from "node:assert/strict";
import { createHash } from "node:crypto";
import { execFileSync } from "node:child_process";
import { cpSync, existsSync, mkdirSync, mkdtempSync, readFileSync, realpathSync, rmSync, writeFileSync } from "node:fs";
import { tmpdir } from "node:os";
import { join } from "node:path";
import { fileURLToPath } from "node:url";
import { test } from "node:test";

const root = fileURLToPath(new URL("../", import.meta.url));
const hash = bytes => createHash("sha256").update(bytes).digest("hex");
const env = { ...process.env, GOWORK: "off" };
delete env.GOPURS_COMPILER_SNAPSHOT;
const run = (command, args, cwd, extra = {}) => execFileSync(command, args, { cwd, env: { ...env, ...extra },
  encoding: "utf8", timeout: 120000, maxBuffer: 16 * 1024 * 1024 });

test("JS cache helper identity changes with each multi-package overlay source", t => {
  const workspace = realpathSync(mkdtempSync(join(tmpdir(), "gopurs-snapshot-overlay-")));
  t.after(() => rmSync(workspace, { recursive: true, force: true }));
  mkdirSync(join(workspace, "tools"));
  for (const name of ["build-cache", "c1c", "build-cache.mjs"])
    cpSync(join(root, "tools", name), join(workspace, "tools", name), { recursive: true });
  const helper = () => run(process.execPath, [join(workspace, "tools/build-cache.mjs"), "--helper"], workspace).trim();
  let previous = helper();
  assert.equal(helper(), previous, "unchanged helper should be reused");
  for (const file of ["build-cache/inputpack.go", "build-cache/overlay_sections.go", "build-cache/capsules.go", "c1c/partition.go", "c1c/literals.go"]) {
    const path = join(workspace, "tools", file);
    writeFileSync(path, readFileSync(path, "utf8") + "\n// changed source for compiler snapshot regression\n");
    const current = helper();
    assert.notEqual(current, previous, `${file} must invalidate the captured helper`);
    assert.ok(existsSync(current));
    previous = current;
  }
});

test("JS executes and identifies the captured bundle even if installation changes after load", t => {
  const workspace = realpathSync(mkdtempSync(join(tmpdir(), "gopurs-snapshot-js-")));
  t.after(() => rmSync(workspace, { recursive: true, force: true }));
  mkdirSync(join(workspace, "bin"));
  mkdirSync(join(workspace, "tools"));
  for (const name of ["build-cache", "c1c", "build-cache.mjs", "ffi-runner.mjs", "ffi_gen.wasm", "wasm_exec.js"])
    cpSync(join(root, "tools", name), join(workspace, "tools", name), { recursive: true });
  writeFileSync(join(workspace, "package.json"), '{"type":"module"}\n');
  const snapshot = readFileSync(join(root, "src/Gopurs/CompilerSnapshot.js"), "utf8");
  const harness = `
import { createHash } from 'node:crypto';
import { renameSync } from 'node:fs';
const marker = 'old';
if (!process.env.GOPURS_COMPILER_SNAPSHOT) renameSync(process.env.REPLACEMENT, bundle);
withSnapshot(() => {
  const root = process.env.GOPURS_COMPILER_SNAPSHOT;
  const manifest = JSON.parse(readFileSync(join(root, 'compiler.json')));
  const resources = Object.fromEntries(manifest.resources.map(input => [input.name,
    createHash('sha256').update(readFileSync(join(root, input.path))).digest('hex')]));
  console.log(JSON.stringify({ marker, root, resources }));
  release();
})();
`;
  const old = snapshot + harness, next = old.replace("const marker = 'old'", "const marker = 'new'");
  writeFileSync(join(workspace, "bin/gopurs.js"), old);
  writeFileSync(join(workspace, "bin/next.js"), next);
  const result = JSON.parse(run(process.execPath, [join(workspace, "bin/gopurs.js")], workspace, { REPLACEMENT: join(workspace, "bin/next.js") }));
  assert.equal(result.marker, "new", "must re-execute the captured code, not use its old in-memory closure");
  assert.equal(result.resources.compiler, hash(next));
  for (const name of ["ffi-runner.mjs", "ffi_gen.wasm", "wasm_exec.js"]) assert.equal(result.resources[name], hash(readFileSync(join(root, "tools", name))));
  assert.ok(!existsSync(result.root), "snapshot lifecycle releases its resources");
});

test("native re-exec identifies the binary actually executed after installation replacement", t => {
  const workspace = realpathSync(mkdtempSync(join(tmpdir(), "gopurs-snapshot-go-")));
  t.after(() => rmSync(workspace, { recursive: true, force: true }));
  writeFileSync(join(workspace, "go.mod"), "module gopurs/output\n\ngo 1.22\n");
  run(process.execPath, [join(root, "tools/prepare-native-output.mjs"), workspace], root);
  writeFileSync(join(workspace, "main.go"), `package main
import ("crypto/sha256"; "encoding/hex"; "encoding/json"; "fmt"; "os"; "runtime/pprof"; "gopurs/output/gopurs_build_cache")
var marker = "old"
func main() {
  if os.Getenv("PPROF") == "1" {
    file, err := os.Create("cpu.prof"); if err != nil { panic(err) }
    if err := pprof.StartCPUProfile(file); err != nil { panic(err) }
    defer pprof.StopCPUProfile()
  }
  if os.Getenv("GOPURS_COMPILER_SNAPSHOT") == "" {
    executable, err := os.Executable(); if err != nil { panic(err) }
    if err := os.Rename(os.Getenv("REPLACEMENT"), executable); err != nil { panic(err) }
  }
  if err := gopurs_build_cache.FreezeNativeCompiler(); err != nil { panic(err) }
  inputs, err := gopurs_build_cache.CompilerInputs(); if err != nil { panic(err) }
  bytes, err := os.ReadFile(inputs[0].Path); if err != nil { panic(err) }
  sum := sha256.Sum256(bytes)
  result, _ := json.Marshal(map[string]string{"marker": marker, "root": os.Getenv("GOPURS_COMPILER_SNAPSHOT"), "sha256": hex.EncodeToString(sum[:])})
  gopurs_build_cache.ReleaseCompilerSnapshot()
  fmt.Println(string(result))
}
`);
  const binary = join(workspace, "compiler"), next = join(workspace, "next");
  run("go", ["build", "-o", binary, "."], workspace);
  run("go", ["build", "-ldflags=-X main.marker=new", "-o", next, "."], workspace);
  const expected = hash(readFileSync(next));
  const result = JSON.parse(run(binary, [], workspace, { REPLACEMENT: next }));
  assert.equal(result.marker, "new");
  assert.equal(result.sha256, expected);
  assert.ok(!existsSync(result.root));
  const image = readFileSync(binary);
  for (let i = 0; i < 12; i++) {
    const profiled = join(workspace, "profiled-" + i), replacement = profiled + "-next";
    writeFileSync(profiled, image, { mode: 0o755 });
    writeFileSync(replacement, image, { mode: 0o755 });
    const result = JSON.parse(run(profiled, [], workspace, { REPLACEMENT: replacement, PPROF: "1" }));
    assert.equal(result.sha256, expected);
    assert.ok(!existsSync(result.root));
  }
  run("go", ["tool", "pprof", "-raw", join(workspace, "cpu.prof")], workspace);
});
