import assert from "node:assert/strict";
import { execFileSync } from "node:child_process";
import { cpSync, mkdirSync, mkdtempSync, readFileSync, rmSync, writeFileSync } from "node:fs";
import { tmpdir } from "node:os";
import { join } from "node:path";
import { fileURLToPath } from "node:url";
import { test } from "node:test";
import { escapeGoStringImpl } from "../src/Gopurs/Printer.js";
import { newBuilderImpl, pushImpl, toStringImpl } from "../src/Gopurs/Printer/Builder.js";
import { memoizeName } from "../src/Gopurs/GoAst.js";
import { concatStringArrays } from "../src/Gopurs/GoImports.js";
import { referencedImportsImpl } from "../src/Gopurs/GoCode.js";
import { referencedImports } from "../output/Gopurs.GoCode/index.js";

const root = fileURLToPath(new URL("../", import.meta.url));
const run = (command, args, cwd) => execFileSync(command, args, {
  cwd, encoding: "utf8", env: { ...process.env, GOWORK: "off" },
  stdio: ["ignore", "pipe", "pipe"],
});

function wtf8Hex(value) {
  const bytes = [];
  for (const character of value) {
    const code = character.codePointAt(0);
    if (code >= 0xd800 && code <= 0xdfff) {
      bytes.push(0xe0 | (code >> 12), 0x80 | ((code >> 6) & 63), 0x80 | (code & 63));
    } else {
      bytes.push(...Buffer.from(character));
    }
  }
  return Buffer.from(bytes).toString("hex");
}

const importArrays = [[], [[]], [[], ["math", ""], [], ["math", "sync", "é😀"], []]];

test("JS printer builders have private storage and preserve completed strings", () => {
  const first = newBuilderImpl();
  const second = newBuilderImpl();
  assert.equal(toStringImpl(first), "");
  for (const chunk of ["a", "", "é😀", "\ud800", "\0"]) {
    assert.equal(pushImpl(first)(chunk), first);
  }
  pushImpl(second)("other");
  const completed = toStringImpl(first);
  assert.equal(completed, "aé😀\ud800\0");
  pushImpl(first)("x".repeat(8192));
  assert.equal(completed, "aé😀\ud800\0");
  assert.equal(toStringImpl(second), "other");
  assert.equal(toStringImpl(first), completed + "x".repeat(8192));
});

test("JS name caches belong to each partial application, including empty results", () => {
  const calls = [];
  const memo = memoizeName(name => { calls.push(name); return name.toUpperCase(); });
  for (let i = 0; i < 3; i++) {
    assert.equal(memo("abc"), "ABC");
    assert.equal(memo(""), "");
  }
  assert.deepEqual(calls, ["abc", ""]);
  assert.equal(memoizeName(name => `other:${name}`)("abc"), "other:abc");
});

test("JS import FFI preserves ordered duplicates, fresh output and the caller's scanner", () => {
  for (const arrays of importArrays) {
    const frozen = Object.freeze(arrays.map(values => Object.freeze([...values])));
    const result = concatStringArrays(frozen);
    assert.deepEqual(result, arrays.flat());
    result.push("private");
    assert.deepEqual(concatStringArrays(frozen), arrays.flat());
  }
  const calls = [];
  const result = ["chosen-by-caller"];
  assert.equal(referencedImportsImpl(text => { calls.push(text); return result; })("math.Abs"), result);
  assert.deepEqual(calls, ["math.Abs"]);
});

test("native compiler FFI preserves JS contracts and value lifetimes under the race detector", t => {
  const directory = mkdtempSync(join(tmpdir(), "gopurs-native-ffi-"));
  t.after(() => rmSync(directory, { recursive: true, force: true }));
  run(process.execPath, [join(root, "tools/embed-runtime.mjs")], root);
  run(process.execPath, [join(root, "tools/prepare-native-output.mjs"), directory], root);
  for (const name of ["GoAst", "Printer", "Printer/Builder", "Runtime", "FfiSupport", "GoCode", "GoImports", "Metrics"]) {
    const source = readFileSync(join(root, `src/Gopurs/${name}.go`), "utf8");
    writeFileSync(join(directory, `${name.replaceAll("/", "_")}.go`), source.replace(/^package Gopurs_\w+$/m, "package main"));
  }
  cpSync(join(root, "tools/native-ffi/ffi_test.go"), join(directory, "ffi_test.go"));
  cpSync(join(root, "runtime/runtime.go"), join(directory, "runtime.txt"));
  mkdirSync(join(directory, "gopurs_runtime"));
  cpSync(join(root, "runtime/runtime.go"), join(directory, "gopurs_runtime/runtime.go"));
  const values = ["", "plain", "quotes\" slash\\", "\0\b\t\n\r\x1f\x7f", "é漢€", "💻𝄞", "\u2028\u2029",
    "\ud800", "\udfff", "a\ud800z", "\ud800\ud800\udc00\udfff"];
  writeFileSync(join(directory, "cases.json"), JSON.stringify(values.map(value => ({
    input: wtf8Hex(value), expected: escapeGoStringImpl(value),
  }))));
  const fragments = ["", "math", "math.", "/", "/*", "\"math.Abs\\", "'\\", "`sync.Once",
    "math.Abs(1); unsafe.Pointer(nil); sync.Once{}; gopurs_runtime.Value{}; math.Abs(2)",
    "0math.Abs; _math.Abs; math .Abs; math/* */.Abs", "// math.Abs\nsync.Once{}",
    ...["ÿ", "😀", "\ud800", "\udfff"].flatMap(prefix => [
      `${prefix}math.Abs;`, `\"${prefix} math.Abs\"; unsafe.Pointer(nil)`,
    ]),
    "/* unsafe.Pointer */ \"math.Abs\"; ".repeat(1000) + "sync.Once{}",
  ];
  writeFileSync(join(directory, "imports.json"), JSON.stringify(fragments.map(input => ({
    input: wtf8Hex(input), expected: referencedImports(input),
  }))));
  writeFileSync(join(directory, "arrays.json"), JSON.stringify(importArrays.map(input => ({
    input, expected: concatStringArrays(input),
  }))));
  writeFileSync(join(directory, "go.mod"), "module gopurs/output\n\ngo 1.22\n");
  const events = run("go", ["test", "-json", "-race", "-count=1", "-timeout=45s", "./..."], directory)
    .trim().split("\n").map(line => JSON.parse(line));
  const passed = new Set(events.filter(event => event.Action === "pass").map(event => event.Test));
  for (const name of ["Escaping", "Memoization", "RuntimeBytes", "NativeBridge", "BuilderLifetime",
    "ImportScanner", "ImportArrayOwnership", "Metrics"]) {
    assert.ok(passed.has(`Test${name}`), `${name} did not pass`);
  }
});
