import { mkdirSync, readFileSync, writeFileSync } from "node:fs";
import { resolve, join } from "node:path";
import { fileURLToPath } from "node:url";

// The native compiler links the same parser and cache sources as the other
// hosts. No parser executable, JavaScript runtime or source lookup is needed
// at runtime. Shared C1c code is copied as a subpackage so the generated
// gopurs/output module can resolve the build cache's import.
const output = process.argv[2];
if (!output || process.argv.length !== 3) {
  throw new Error("Usage: node tools/prepare-native-output.mjs <generated-output-directory>");
}
const generatedModule = "gopurs/output";
const buildCache = ["contract.go", "store.go", "session.go", "inputpack.go", "compiler.go", "modules.go", "capsules.go", "packages.go", "overlay.go", "overlay_sections.go", "packs.go"];
const c1c = ["manifest.go", "overlay.go", "qualify.go", "rewrite.go", "partition.go", "literals.go"];
const groups = [
  { directory: "ffi-gen", packageName: "gopurs_ffi_parser", destination: "gopurs_ffi_parser", sourcePackage: "main", files: ["parser.go", "types.go", "native_api.go", "bindings.go", "packages.go", "workspace.go"] },
  { directory: "build-cache", packageName: "gopurs_build_cache", destination: "gopurs_build_cache", sourcePackage: "main", files: buildCache,
    replace: [["\"gopurs/c1c\"", `"${generatedModule}/gopurs_c1c"`]] },
  { directory: "c1c", packageName: "c1c", destination: "gopurs_c1c", sourcePackage: "c1c", files: c1c },
];
for (const group of groups) {
  const sourceDirectory = fileURLToPath(new URL(`./${group.directory}/`, import.meta.url));
  const destination = join(resolve(output), group.destination);
  const sources = group.files.map(name => {
    const source = readFileSync(join(sourceDirectory, name), "utf8");
    const declaration = new RegExp(`^package ${group.sourcePackage}\\n`, "m");
    if (!declaration.test(source)) {
      throw new Error(`Unexpected package in ${group.directory}/${name}`);
    }
    let rewritten = source.replace(declaration, `package ${group.packageName}\n`);
    for (const [from, to] of group.replace ?? []) rewritten = rewritten.split(from).join(to);
    return [name, rewritten];
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
