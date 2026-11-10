import { existsSync } from "node:fs";
import { join } from "node:path";
import { fileURLToPath } from "node:url";

// Single source of truth for the quick release-quality gate. The full campaign
// discovers every tools/*.test.mjs file directly.
export const QUICK_TESTS = Object.freeze([
  "release-quality-resources",
  "release-npm-provenance",
  "release-sbom",
  "emission",
  "ffi-errors",
  "launcher",
  "test-runner",
  "build-native",
  "bootstrap-inputs",
  "cli",
  "build-profiles",
  "manifest-prepared-cache",
  "driver-cache",
  "native-go-code",
  "preparation-native",
  "check-output-identity",
  "check-docs",
  "embed-runtime",
  "compiler-snapshot",
  "tool-help",
  "tools-inventory",
]);

export const QUICK_TEST_FILES = Object.freeze(QUICK_TESTS.map((name) => `${name}.test.mjs`));
export const NODE_GATES = Object.freeze(QUICK_TEST_FILES.map((name) => `Node ${name}`));

const toolsDirectory = fileURLToPath(new URL("../", import.meta.url));
for (const name of QUICK_TEST_FILES) {
  if (!existsSync(join(toolsDirectory, name))) {
    throw new Error(`Quick test is missing: tools/${name}`);
  }
}
