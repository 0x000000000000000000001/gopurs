// Adapted from purust/purust/tools/test-native-tast-module.mjs. The boundary
// matrix is reused from the Purust harness (its main() is guarded, so importing
// only defines its exports); only the injected native source (gopurs PBO
// CoreFn/Json.rs), the shared cargo target directory and logging differ. The
// PureScript oracles come from the retained Rust-hosted gopurs workspace.
//
// Usage: node tools/native-pbo/tast-module.mjs GENERATED_RUST [TAST_CORPUS]
import assert from "node:assert/strict";
import { existsSync, readFileSync, readdirSync, statSync, writeFileSync } from "node:fs";
import { join } from "node:path";
import { pathToFileURL } from "node:url";
import {
  createWorkspace, fixtureSource, loadThreadedRust, pboDir, pboSource,
  purustTools, resolveCaseArgs, runCase, writeCrate,
} from "./shared.mjs";

const { rust, corpus } = resolveCaseArgs({ corpus: true });
const threadedRust = await loadThreadedRust();

// Reuse the exported boundary table from the Purust harness instead of copying
// it: the module defines buildBoundaries and only runs its own main() when it
// is the process entry point.
const purustHarness = pathToFileURL(join(purustTools, "test-native-tast-module.mjs")).href;
const { buildBoundaries } = await import(purustHarness);

// Frozen corpus: a directory of <Module>/corefn.json trees or the packed
// [{ name, contents }] diagnostic corpus.
function corpusEntries(source) {
  if (statSync(source).isFile()) {
    return JSON.parse(readFileSync(source, "utf8"))
      .map(entry => ({ name: entry.name, value: JSON.parse(entry.contents) }));
  }
  const entries = [];
  for (const entry of readdirSync(source, { withFileTypes: true })) {
    if (!entry.isDirectory()) continue;
    const file = join(source, entry.name, "corefn.json");
    if (!existsSync(file)) continue;
    entries.push({ name: entry.name, value: JSON.parse(readFileSync(file, "utf8")) });
  }
  return entries;
}

await runCase("tast-module", rust, async ctx => {
  const corpusModules = corpusEntries(corpus).map(entry => entry.value);
  assert.ok(corpusModules.length > 0, `no corefn modules found in ${corpus}`);
  const boundaries = buildBoundaries();

  const directory = createWorkspace("gopurs-native-pbo-tast-module");
  ctx.state.workspace = directory;
  ctx.note(`fixture workspace ${directory}`);
  writeFileSync(join(directory, "modules.ndjson"),
    corpusModules.map(module => JSON.stringify(module)).join("\n") + "\n");
  writeFileSync(join(directory, "boundaries.ndjson"),
    boundaries.map(([mode, name, value]) => `${mode}\t${name}\t${JSON.stringify(value)}`).join("\n") + "\n");

  const modules = [
    "purust_core", "perceus_ptr", "Purs_Data_Argonaut_Core", "Purs_Data_Argonaut_Decode_Error",
    "Purs_Data_Either", "Purs_Data_Maybe", "Purs_Data_Tuple", "Purs_Data_Unfoldable", "Purs_Foreign_Object",
    "Purs_Data_Map_Internal", "Purs_Data_Ord", "Purs_Data_Foldable",
    "Purs_PureScript_Backend_Optimizer_CoreFn", "Purs_PureScript_Backend_Optimizer_CoreFn_Json",
    "Purs_PureScript_Backend_Optimizer_CoreFn_TypeTable", "Purs_PureScript_Backend_Optimizer_CoreFn_Usage",
  ];
  writeCrate(directory, "purust_native_tast_module_test", modules, rust);

  ctx.note(`injecting gopurs ${join(pboDir, "src/PureScript/Backend/Optimizer/CoreFn/Json.rs")}`);
  const ffi = threadedRust(pboSource("PureScript/Backend/Optimizer/CoreFn/Json.rs"));
  writeFileSync(join(directory, "src/main.rs"),
    fixtureSource("test-native-tast-module.rs").replace(/^\s*\/\/ NATIVE_FFI$/m, () => ffi));
  ctx.note(`${corpusModules.length} frozen modules / ${boundaries.length} boundaries from ${corpus}`);

  const executable = ctx.cargoBuild(directory, "purust_native_tast_module_test");
  ctx.spawn("run", executable, [join(directory, "modules.ndjson"), join(directory, "boundaries.ndjson")]);
});
