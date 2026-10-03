// Adapted from purust/purust/tools/test-native-directives.mjs; see that file
// for the ASCII fast-path contract and exact-fallback probes. The generated
// Directives crate comes from the retained Rust-hosted gopurs workspace, so
// the embedded native source under test is the gopurs PBO checkout. No source
// is injected here.
//
// Usage: node tools/native-pbo/directives.mjs GENERATED_RUST
import assert from "node:assert/strict";
import { writeFileSync } from "node:fs";
import { join } from "node:path";
import { createWorkspace, fixtureSource, generatedCrateSource, resolveCaseArgs, runCase, writeCrate } from "./shared.mjs";

const { rust } = resolveCaseArgs();

await runCase("directives", rust, async ctx => {
  const source = generatedCrateSource(rust, "Purs_PureScript_Backend_Optimizer_Directives");
  assert.ok(source.includes("parseDirectiveLinePS") && source.includes("purust_directive_ascii"),
    "GENERATED_RUST is stale: Directives must embed the gopurs PBO native fast path");

  const directory = createWorkspace("gopurs-native-pbo-directives");
  ctx.state.workspace = directory;
  ctx.note(`fixture workspace ${directory}`);
  const modules = [
    "purust_core", "Purs_Data_Either", "Purs_Data_Maybe", "Purs_Data_Tuple",
    "Purs_PureScript_Backend_Optimizer_CoreFn", "Purs_PureScript_Backend_Optimizer_Semantics",
    "Purs_PureScript_Backend_Optimizer_Directives", "Purs_PureScript_Backend_Optimizer_Directives_Defaults",
  ];
  writeCrate(directory, "purust_native_directives_test", modules, rust);
  writeFileSync(join(directory, "src/main.rs"), fixtureSource("test-native-directives.rs"));
  const executable = ctx.cargoBuild(directory, "purust_native_directives_test");
  ctx.spawn("run", executable, []);
});
