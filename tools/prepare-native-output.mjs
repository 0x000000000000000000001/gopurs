import { mkdirSync, readFileSync, writeFileSync } from "node:fs";
import { resolve, join } from "node:path";
import { fileURLToPath } from "node:url";

// The native compiler links the same parser sources as the WASM runner. No
// parser executable, JavaScript runtime or source lookup is needed at runtime.
const output = process.argv[2];
if (!output || process.argv.length !== 3) {
  throw new Error("Usage: node tools/prepare-native-output.mjs <generated-output-directory>");
}
for (const [directory, packageName, files] of [
  ["ffi-gen", "gopurs_ffi_parser", ["parser.go", "types.go", "native_api.go"]],
  ["build-cache", "gopurs_build_cache", ["contract.go", "store.go"]],
]) {
  const sourceDirectory = fileURLToPath(new URL(`./${directory}/`, import.meta.url));
  const destination = join(resolve(output), packageName);
  const sources = files.map(name => {
    const source = readFileSync(join(sourceDirectory, name), "utf8");
    if (!source.startsWith("package main\n")) {
      throw new Error(`Unexpected package in ${directory}/${name}`);
    }
    return [name, source.replace(/^package main\n/, `package ${packageName}\n`)];
  });
  mkdirSync(destination, { recursive: true });
  for (const [name, source] of sources) {
    const target = join(destination, name);
    let previous;
    try {
      previous = readFileSync(target, "utf8");
    } catch (error) {
      if (error.code !== "ENOENT") throw error;
    }
    if (previous !== source) writeFileSync(target, source);
  }
}
