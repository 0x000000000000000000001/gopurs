#!/usr/bin/env node
import { createHash } from "node:crypto";
import { existsSync, readFileSync } from "node:fs";
import { join, relative, resolve } from "node:path";
import { fileURLToPath } from "node:url";

const root = fileURLToPath(new URL("../", import.meta.url));
const defaultArtifacts = [
  "bin/gopurs.js",
  "bin/gopurs-native",
  "bin/gopurs-rust",
  "tools/ffi_gen.wasm",
  "tools/wasm_exec.js",
  "package-lock.json",
  "spago.lock",
];
const requiredArtifacts = [
  "bin/gopurs.js",
  "tools/ffi_gen.wasm",
  "tools/wasm_exec.js",
  "package-lock.json",
  "spago.lock",
];

const args = process.argv.slice(2);
if (args.includes("--help") || args.includes("-h")) {
  console.log("Usage: node tools/bootstrap-checksums.mjs [--require-all] [FILE...]");
  console.log("");
  console.log("Print SHA-256 checksums for the bootstrap artifacts.");
  console.log("With no FILE arguments, existing default artifacts are included.");
  console.log("--require-all fails unless the native and Rust compiler binaries exist too.");
  process.exit(0);
}

const requireAll = args.includes("--require-all");
const unknown = args.find(arg => arg.startsWith("-") && arg !== "--require-all");
if (unknown) {
  console.error(`Unknown option: ${unknown}`);
  process.exit(2);
}

const requested = args.filter(arg => arg !== "--require-all");
const paths = (requested.length
  ? requested.map(path => resolve(root, path))
  : defaultArtifacts.map(path => join(root, path)).filter(path => existsSync(path))).sort();

if (!requested.length) {
  const missing = requiredArtifacts.map(path => join(root, path)).filter(path => !existsSync(path));
  if (missing.length) {
    console.error("Missing required bootstrap artifacts:");
    for (const path of missing) console.error(`  ${relative(root, path)}`);
    process.exit(1);
  }
  if (requireAll) {
    for (const path of ["bin/gopurs-native", "bin/gopurs-rust"]) {
      if (!existsSync(join(root, path))) {
        console.error(`Missing required release artifact: ${path}`);
        process.exit(1);
      }
    }
  }
}

if (requireAll && requested.length) {
  console.error("--require-all cannot be combined with explicit FILE arguments");
  process.exit(2);
}

for (const path of paths) {
  if (!existsSync(path)) {
    console.error(`Missing file: ${path}`);
    process.exit(1);
  }
  const hash = createHash("sha256").update(readFileSync(path)).digest("hex");
  console.log(`${hash}  ${relative(root, path)}`);
}
