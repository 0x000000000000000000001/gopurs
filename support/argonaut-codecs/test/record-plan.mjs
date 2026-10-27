import assert from "node:assert/strict";
import { mkdtempSync, mkdirSync, readFileSync, writeFileSync, rmSync } from "node:fs";
import { tmpdir } from "node:os";
import { dirname, resolve, join } from "node:path";
import { fileURLToPath } from "node:url";
import { spawnSync } from "node:child_process";

const root = resolve(dirname(fileURLToPath(import.meta.url)), "..");
const helper = join(root, "src/Data/Argonaut/Decode/Internal/Record");
const source = readFileSync(`${helper}.js`, "utf8");
const js = await import(`data:text/javascript;base64,${Buffer.from(source).toString("base64")}`);
const right = value => ({ right: value });
assert.deepEqual(js.recordNilImpl(null)(null)(right)({})(null), { right: {} });
const fallback = object => proxy => ({ object, proxy });
const selected = js.recordConsImpl(null)(null)(null)(fallback)(null)(null)(null)(null)(null)(null)(null);
assert.equal(selected, fallback);
assert.deepEqual(selected(12)(34), { object: 12, proxy: 34 });

const work = mkdtempSync(join(tmpdir(), "gopurs-record-plan-"));
try {
  mkdirSync(join(work, "output/gopurs_runtime"), { recursive: true });
  mkdirSync(join(work, "record"));
  const compiler = resolve(root, "../..");
  writeFileSync(join(work, "go.mod"), "module gopurs\n\ngo 1.27.0\n");
  writeFileSync(join(work, "output/gopurs_runtime/runtime.go"), readFileSync(join(compiler, "runtime/runtime.go")));
  writeFileSync(join(work, "record/record.go"), readFileSync(`${helper}.go`));
  writeFileSync(join(work, "record/record_test.go"), readFileSync(join(root, "test/record-plan_test.go")));
  const result = spawnSync("go", ["test", "-race", "-count=1", "./record"], { cwd: work, stdio: "inherit", env: { ...process.env, GOWORK: "off" } });
  if (result.error) throw result.error;
  if (result.status !== 0) process.exitCode = result.status ?? 1;
} finally {
  rmSync(work, { recursive: true, force: true });
}
