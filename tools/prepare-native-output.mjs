import { mkdirSync, readFileSync, writeFileSync } from "node:fs";
import { resolve, join } from "node:path";
import { fileURLToPath } from "node:url";

// The native compiler links the same parser sources as the WASM runner. No
// parser executable, JavaScript runtime or source lookup is needed at runtime.
const output = process.argv[2];
if (!output || process.argv.length !== 3) {
  throw new Error("Usage: node tools/prepare-native-output.mjs <generated-output-directory>");
}
const sourceDirectory = fileURLToPath(new URL("./ffi-gen/", import.meta.url));
const destination = join(resolve(output), "gopurs_ffi_parser");
const sources = ["parser.go", "types.go", "native_api.go"].map(name => {
  const source = readFileSync(join(sourceDirectory, name), "utf8");
  if (!source.startsWith("package main\n")) {
    throw new Error(`Unexpected parser package in ${name}`);
  }
  return [name, source.replace(/^package main\n/, "package gopurs_ffi_parser\n")];
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
