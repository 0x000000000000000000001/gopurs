import assert from "node:assert/strict";
import { spawnSync } from "node:child_process";
import { mkdirSync, mkdtempSync, rmSync, writeFileSync } from "node:fs";
import { tmpdir } from "node:os";
import { dirname, join } from "node:path";

/** Create a temporary directory owned by a node:test context. */
export function temporaryDirectory(t, prefix) {
  const directory = mkdtempSync(join(tmpdir(), prefix));
  t.after(() => rmSync(directory, { recursive: true, force: true }));
  return directory;
}

/** Write a Go module with the canonical runtime and a go.mod. */
export function goModule(t, prefix, { module = "gopurs/output", go = "1.22", runtime } = {}) {
  const directory = temporaryDirectory(t, prefix);
  if (runtime !== undefined) {
    mkdirSync(join(directory, "gopurs_runtime"));
    writeFileSync(join(directory, "gopurs_runtime/runtime.go"), runtime);
  }
  writeFileSync(join(directory, "go.mod"), `module ${module}\n\ngo ${go}\n`);
  return directory;
}

/** Write one Go source file below a fixture directory. */
export function writeGoFile(directory, name, source) {
  const file = join(directory, name);
  mkdirSync(dirname(file), { recursive: true });
  writeFileSync(file, source);
  return file;
}

/** Run a Go command in a fixture and fail with its complete output. */
export function runGo(directory, args, { timeout = 120_000, ...options } = {}) {
  const result = spawnSync("go", args, {
    cwd: directory,
    encoding: "utf8",
    timeout,
    env: { ...process.env, GOWORK: "off" },
    ...options,
  });
  assert.ifError(result.error);
  assert.equal(result.status, 0, result.stdout + result.stderr);
  return result;
}
