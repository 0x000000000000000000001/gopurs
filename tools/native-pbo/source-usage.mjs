// Adapted from purust/purust/tools/test-native-source-usage.mjs; see that file
// for the differential contract (native source-usage validator vs the
// PureScript oracle on synthetic modules and a frozen TAST corpus). The
// candidate is the gopurs PBO CoreFn/Usage.rs injected at NATIVE_FFI; the
// oracles come from the retained Rust-hosted gopurs workspace. Only paths,
// logging and the shared cargo target directory differ.
//
// Usage: node tools/native-pbo/source-usage.mjs GENERATED_RUST [TAST_CORPUS]
import assert from "node:assert/strict";
import { existsSync, readFileSync, readdirSync, statSync, writeFileSync } from "node:fs";
import { join } from "node:path";
import {
  createWorkspace, fixtureSource, generatedCrateSource, loadThreadedRust, pboDir, pboSource,
  resolveCaseArgs, runCase, writeCrate,
} from "./shared.mjs";

const { rust, corpus } = resolveCaseArgs({ corpus: true });
const threadedRust = await loadThreadedRust();

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

await runCase("source-usage", rust, async ctx => {
  const usageGenerated = generatedCrateSource(rust, "Purs_PureScript_Backend_Optimizer_CoreFn_Usage");
  for (const needle of [
    "PureScript_Backend_Optimizer_CoreFn_Usage_validateSourceUsageModulePS",
    "PureScript_Backend_Optimizer_CoreFn_Usage_validateSourceUsageModuleImpl",
    "Native source-usage validator",
  ]) {
    assert.ok(usageGenerated.includes(needle),
      `GENERATED_RUST is stale: Usage must contain ${JSON.stringify(needle)}`);
  }
  const jsonGenerated = generatedCrateSource(rust, "Purs_PureScript_Backend_Optimizer_CoreFn_Json");
  for (const needle of [
    "PureScript_Backend_Optimizer_CoreFn_Json_decodeModulePS",
    "PureScript_Backend_Optimizer_CoreFn_Json_decodeModuleImpl",
  ]) {
    assert.ok(jsonGenerated.includes(needle),
      `GENERATED_RUST is stale: Json must contain ${JSON.stringify(needle)}`);
  }

  const corpusModules = corpusEntries(corpus).map(entry => entry.value);
  assert.ok(corpusModules.length > 0, `no corefn modules found in ${corpus}`);

  const directory = createWorkspace("gopurs-native-pbo-source-usage");
  ctx.state.workspace = directory;
  ctx.note(`fixture workspace ${directory}`);
  writeFileSync(join(directory, "modules.ndjson"),
    corpusModules.map(module => JSON.stringify(module)).join("\n") + "\n");

  const modules = [
    "purust_core", "perceus_ptr", "Purs_Data_Argonaut_Core", "Purs_Data_Argonaut_Decode_Error",
    "Purs_Data_Either", "Purs_Data_Maybe", "Purs_PureScript_Backend_Optimizer_CoreFn",
    "Purs_PureScript_Backend_Optimizer_CoreFn_Json", "Purs_PureScript_Backend_Optimizer_CoreFn_Usage",
  ];
  writeCrate(directory, "purust_native_source_usage_test", modules, rust);

  ctx.note(`injecting gopurs ${join(pboDir, "src/PureScript/Backend/Optimizer/CoreFn/Usage.rs")}`);
  const ffi = threadedRust(pboSource("PureScript/Backend/Optimizer/CoreFn/Usage.rs"));
  assert.ok(ffi.includes("Native source-usage validator"),
    "gopurs PBO CoreFn/Usage.rs lost the source-usage validator marker");
  writeFileSync(join(directory, "src/main.rs"),
    fixtureSource("test-native-source-usage.rs").replace(/^\s*\/\/ NATIVE_FFI$/m, () => ffi));
  ctx.note(`${corpusModules.length} frozen modules from ${corpus}`);

  const executable = ctx.cargoBuild(directory, "purust_native_source_usage_test");
  ctx.spawn("run", executable, [join(directory, "modules.ndjson")]);
});
