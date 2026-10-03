// Adapted from purust/purust/tools/test-native-memo.mjs; see that file for the
// differential contract (inner ExprType/BackendSyntax memo keying, FIFO
// eviction, reentrancy, independent effects). Only logging, the shared cargo
// target directory and workspace retention differ. The generated crates come
// from the retained Rust-hosted gopurs workspace, so the embedded native
// source under test is the gopurs PBO checkout.
//
// Usage: node tools/native-pbo/memo.mjs GENERATED_RUST
import assert from "node:assert/strict";
import { writeFileSync } from "node:fs";
import { join } from "node:path";
import { createWorkspace, fixtureSource, generatedCrateSource, resolveCaseArgs, runCase, writeCrate } from "./shared.mjs";

const { rust } = resolveCaseArgs();

await runCase("memo", rust, async ctx => {
  const sources = new Map();
  const generatedSource = (module, required) => {
    if (!sources.has(module)) sources.set(module, generatedCrateSource(rust, module));
    const content = sources.get(module);
    for (const needle of [].concat(required)) {
      assert.ok(content.includes(needle),
        `GENERATED_RUST is stale: ${module} must contain ${JSON.stringify(needle)}`);
    }
    return content;
  };

  generatedSource("Purs_PureScript_Backend_Optimizer_BoundedMemo", [
    "PureScript_Backend_Optimizer_BoundedMemo_createBoundedMemo",
    "PureScript_Backend_Optimizer_BoundedMemo_createStringMemo",
    "downcast_ref::<Rc<Purs_PureScript_Backend_Optimizer_CoreFn::ExprType>>()",
    "downcast_ref::<Rc<Purs_PureScript_Backend_Optimizer_Syntax::BackendSyntax>>()",
    "Pointer(3,",
    "Pointer(4,",
  ]);
  generatedSource("Purs_PureScript_Backend_Optimizer_CoreFn", "pub enum ExprType");
  generatedSource("Purs_PureScript_Backend_Optimizer_Syntax", ["pub enum BackendSyntax", "PrimUndefined"]);
  // The Rust-hosted gopurs workspace is generated with --threaded; the test
  // source mirrors the Arc model with `use std::sync::Arc as Rc`.
  generatedSource("purust_core", "std::sync::Arc");

  const directory = createWorkspace("gopurs-native-pbo-memo");
  ctx.state.workspace = directory;
  ctx.note(`fixture workspace ${directory}`);
  const modules = [
    "purust_core", "Purs_PureScript_Backend_Optimizer_CoreFn",
    "Purs_PureScript_Backend_Optimizer_Syntax", "Purs_PureScript_Backend_Optimizer_BoundedMemo",
  ];
  writeCrate(directory, "purust_native_memo_test", modules, rust);
  writeFileSync(join(directory, "src/main.rs"), fixtureSource("test-native-memo.rs"));
  const executable = ctx.cargoBuild(directory, "purust_native_memo_test");
  ctx.spawn("run", executable, []);
});
