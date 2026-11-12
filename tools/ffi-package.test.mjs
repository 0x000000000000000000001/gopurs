import assert from "node:assert/strict";
import { execFileSync, spawnSync } from "node:child_process";
import { mkdtempSync, readFileSync, rmSync } from "node:fs";
import { tmpdir } from "node:os";
import { join } from "node:path";
import { fileURLToPath } from "node:url";
import test from "node:test";
import { findTypedCompiler } from "./build/native-workspace.mjs";
import { findOfficialCompiler } from "./release/official-compiler.mjs";

const root = fileURLToPath(new URL("../", import.meta.url));

test("installed tarball preserves idiomatic FFI, migration and cache recovery", { timeout: 600000 }, t => {
  const work = mkdtempSync(join(tmpdir(), "gopurs-ffi-package-"));
  let passed = false;
  t.after(() => {
    if (passed && process.env.GOPURS_TEST_KEEP_WORKSPACE !== "1") rmSync(work, { recursive: true, force: true });
    else console.log(`Package artifacts retained: ${work}`);
  });
  const [manifest] = JSON.parse(execFileSync("npm", ["pack", "--ignore-scripts", "--json", "--pack-destination", work], {
    cwd: root, encoding: "utf8", timeout: 60000,
  }));
  const selected = process.env.GOPURS_TEST_MODES?.split(",") ?? ["js", "native-parallel"];
  const hosts = [
    ...(selected.includes("js") ? ["javascript"] : []),
    ...(selected.some(mode => mode.startsWith("native")) ? ["native"] : []),
  ];
  assert.ok(hosts.length, "GOPURS_TEST_MODES must include js or native");
  const report = join(work, "report.json");
  const args = [
    join(root, "tools/release/release-smoke.mjs"),
    "--package", join(work, manifest.filename),
    "--purs", findTypedCompiler(root, process.env.GOPURS_PURS),
    "--official-purs", findOfficialCompiler(root),
    "--hosts", hosts.join(","),
    "--report", report,
  ];
  if (hosts.includes("native")) args.push("--native", join(root, "bin/gopurs-native"));
  if (process.env.GOPURS_TEST_KEEP_WORKSPACE === "1") args.push("--keep-workspace");
  const result = spawnSync(process.execPath, args, { cwd: root, env: process.env, encoding: "utf8", timeout: 540000, maxBuffer: 16 * 1024 * 1024 });
  process.stdout.write(result.stdout ?? "");
  process.stderr.write(result.stderr ?? "");
  assert.ifError(result.error);
  assert.equal(result.status, 0, "installed package smoke failed");
  const evidence = JSON.parse(readFileSync(report, "utf8"));
  assert.deepEqual(evidence.tests.map(result => result.host), hosts);
  assert.ok(evidence.tests.every(result => result.passed && result.stages.includes("recovery")));
  console.log(`Package ${evidence.version}, ${evidence.go}, ${evidence.node}: ${hosts.join(", ")}`);
  passed = true;
});
