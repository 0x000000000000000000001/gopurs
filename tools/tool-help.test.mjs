import assert from "node:assert/strict";
import { spawnSync } from "node:child_process";
import test from "node:test";

const tools = [
  "build/build-native.mjs",
  "build/build-rust.mjs",
  "release/release-benchmark.mjs",
  "release/release-prepare.mjs",
  "release/release-quality-report.mjs",
  "release/release-quality.mjs",
  "release/release-smoke.mjs",
  "bootstrap-checksums.mjs",
  "check-docs.mjs",
  "check-output-identity.mjs",
  "test-native-pbo.mjs",
];

test("manual tools provide a usable --help entry point", () => {
  for (const tool of tools) {
    const result = spawnSync(process.execPath, [`tools/${tool}`, "--help"], { encoding: "utf8" });
    assert.equal(result.status, 0, `${tool}: ${result.stderr}`);
    assert.match(result.stdout + result.stderr, /Usage/i, tool);
  }
});
