import assert from "node:assert/strict";
import { execFileSync } from "node:child_process";
import { cpSync, mkdtempSync, readFileSync, rmSync, writeFileSync } from "node:fs";
import { tmpdir } from "node:os";
import { join } from "node:path";
import { fileURLToPath } from "node:url";
import { test } from "node:test";
import { escapeGoStringImpl } from "../src/Gopurs/Printer.js";

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

test("native FFI preserves JS escaping, isolated memoization, runtime bytes and parser behavior", t => {
  const directory = mkdtempSync(join(tmpdir(), "gopurs-native-ffi-"));
  t.after(() => rmSync(directory, { recursive: true, force: true }));
  run(process.execPath, [join(root, "tools/embed-runtime.mjs")], root);
  run(process.execPath, [join(root, "tools/prepare-native-output.mjs"), directory], root);
  for (const name of ["GoAst", "Printer", "Runtime", "FfiSupport"]) {
    const source = readFileSync(join(root, `src/Gopurs/${name}.go`), "utf8");
    writeFileSync(join(directory, `${name}.go`), source.replace(/^package Gopurs_\w+$/m, "package main"));
  }
  cpSync(join(root, "tools/native-ffi/ffi_test.go"), join(directory, "ffi_test.go"));
  cpSync(join(root, "runtime/runtime.go"), join(directory, "runtime.txt"));
  const values = ["", "plain", "quotes\" slash\\", "\0\b\t\n\r\x1f\x7f", "é漢€", "💻𝄞", "\u2028\u2029",
    "\ud800", "\udfff", "a\ud800z", "\ud800\ud800\udc00\udfff"];
  writeFileSync(join(directory, "cases.json"), JSON.stringify(values.map(value => ({
    input: wtf8Hex(value), expected: escapeGoStringImpl(value),
  }))));
  writeFileSync(join(directory, "go.mod"), "module gopurs/output\n\ngo 1.22\n");
  assert.match(run("go", ["test", "-race", "-count=1", "./..."], directory), /^ok\s/m);
});
