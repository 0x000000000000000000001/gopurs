#!/usr/bin/env node
import { execFileSync } from "node:child_process";

const output = execFileSync("npm", ["pack", "--dry-run", "--json", "--ignore-scripts"], {
  encoding: "utf8",
  maxBuffer: 16 * 1024 * 1024,
});
const [manifest] = JSON.parse(output);
const paths = manifest.files.map(file => file.path);
const required = [
  "CHANGELOG.md",
  "CODE_OF_CONDUCT.md",
  "CONTRIBUTING.md",
  "LICENSE",
  "NOTICE",
  "README.md",
  "SECURITY.md",
  "THIRD_PARTY_NOTICES.md",
  "bin/gopurs",
  "bin/gopurs.js",
  "docs/architecture-overview.md",
  "docs/benchmark-protocol.md",
  "docs/bootstrap.md",
  "docs/cli.md",
  "docs/limitations.md",
  "docs/migration.md",
  "docs/publication-scope.md",
  "docs/quickstart.md",
  "docs/troubleshooting.md",
  "package.json",
  "tools/ffi_gen.wasm",
  "tools/ffi-runner.mjs",
  "tools/wasm_exec.js",
];
const forbidden = [
  /^tests\//,
  /^runtime\//,
  /^src\//,
  /^\.github\//,
  /^bin\/gopurs-(?:native|rust)$/,
  /^tools\/.*\.test\.mjs$/,
  /^docs\/build-/,
  /^docs\/testing\.md$/,
  /^docs\/plan-/,
];

const missing = required.filter(path => !paths.includes(path));
if (missing.length) {
  console.error("Package is missing required files:");
  for (const path of missing) console.error(`  ${path}`);
  process.exit(1);
}
const excluded = paths.filter(path => forbidden.some(pattern => pattern.test(path)));
if (excluded.length) {
  console.error("Package contains files excluded from the publication surface:");
  for (const path of excluded) console.error(`  ${path}`);
  process.exit(1);
}
if (manifest.entryCount > 64) {
  console.error(`Package contains ${manifest.entryCount} files; the limit is 64.`);
  process.exit(1);
}
if (manifest.unpackedSize > 25 * 1024 * 1024) {
  console.error(`Package unpacks to ${manifest.unpackedSize} bytes; the limit is 25 MiB.`);
  process.exit(1);
}
console.log(`Package surface OK: ${manifest.entryCount} files, ${manifest.size} bytes packed, ${manifest.unpackedSize} bytes unpacked.`);
