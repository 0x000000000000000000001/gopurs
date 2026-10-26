import assert from "node:assert/strict";
import { createHash } from "node:crypto";
import { spawnSync } from "node:child_process";
import { copyFileSync, cpSync, mkdtempSync, readFileSync, readdirSync, rmSync, writeFileSync } from "node:fs";
import { tmpdir } from "node:os";
import { join } from "node:path";
import { fileURLToPath } from "node:url";
import { test } from "node:test";

const root = fileURLToPath(new URL("../", import.meta.url));
const hash = path => createHash("sha256").update(readFileSync(path)).digest("hex");
const inventory = directory => Object.fromEntries(readdirSync(directory, { withFileTypes: true })
  .filter(entry => entry.isFile()).map(entry => [entry.name, hash(join(directory, entry.name))]));

test("output directory barriers survive failed synchronization and process death", t => {
  const workspace = mkdtempSync(join(tmpdir(), "gopurs-cache-output-"));
  let passed = false;
  t.after(() => {
    if (passed && process.env.GOPURS_TEST_KEEP_WORKSPACE !== "1") rmSync(workspace, { recursive: true, force: true });
    else console.log(`Output barrier workspace retained: ${workspace}`);
  });
  const source = join(root, "tools/build-cache"), before = inventory(source);
  const fixtures = join(root, "tools/build-cache-fixtures");
  // build-cache/go.mod replaces gopurs/c1c with ../c1c. Keep that module next
  // to the private copy so the test build resolves it outside the checkout.
  cpSync(join(root, "tools/c1c"), join(workspace, "c1c"), { recursive: true });
  const env = { ...process.env, GOWORK: "off" };
  delete env.GOPURS_TEST_CHILD_REQUEST;
  delete env.GOPURS_TEST_PAUSE_DIRECTORY;
  const records = [];
  for (const [fixture, names] of [
    ["barriers_test.go", ["OutputDirectoryBarriers", "OutputBarrierFailureRecovery"]],
    ["crash_test.go", ["OutputBarrierProcessDeath"]],
  ]) {
    const directory = join(workspace, fixture.replace("_test.go", ""));
    cpSync(source, directory, { recursive: true });
    assert.deepEqual(inventory(directory), before);
    // Wrap one filesystem boundary in this private test build only. All
    // publication code remains the exact production source; there is no
    // production-global fault-injection switch or callback.
    const path = join(directory, "store.go"), text = readFileSync(path, "utf8");
    const signature = "func syncDirectory(path string) error {";
    assert.equal(text.split(signature).length, 2);
    writeFileSync(path, text.replace(signature, "func untracedSyncDirectory(path string) error {"));
    copyFileSync(join(fixtures, fixture), join(directory, "output_probe_test.go"));
    const args = ["test", "-race", "-count=1", "-timeout=90s", "-v", "-run", "^Test(" + names.join("|") + ")$", "./..."];
    const result = spawnSync("go", args, { cwd: directory, env,
      encoding: "utf8", timeout: 120000, maxBuffer: 16 * 1024 * 1024 });
    writeFileSync(join(directory, "test.log"), (result.stdout ?? "") + (result.stderr ?? ""));
    records.push({ fixture, args, status: result.status, signal: result.signal, error: result.error?.message,
      source: before, wrappedStore: hash(path), fixtureSHA256: hash(join(fixtures, fixture)), logSHA256: hash(join(directory, "test.log")) });
    writeFileSync(join(workspace, "records.json"), JSON.stringify(records, null, 2) + "\n");
    assert.ifError(result.error);
    assert.equal(result.status, 0, `${fixture}: ${result.stdout}\n${result.stderr}`);
    for (const name of names) assert.ok(result.stdout.includes(`--- PASS: Test${name} `), result.stdout);
  }
  assert.deepEqual(inventory(source), before);
  passed = true;
});
