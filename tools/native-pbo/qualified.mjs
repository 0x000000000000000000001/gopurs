// Differential contract for the borrowed `Qualified Ident` comparison.
// Usage: node tools/native-pbo/qualified.mjs GENERATED_RUST
//
// The native Rust source is injected from the gopurs PBO checkout and must be
// byte-identical to the FFI embedded in the retained CoreFn crate. The fixture
// then links the real generated functions: the monomorphic PureScript oracle,
// the public wrappers and the untouched generic `ordQualified` instance.
import assert from "node:assert/strict";
import { readFileSync, writeFileSync } from "node:fs";
import { dirname, join } from "node:path";
import { fileURLToPath } from "node:url";
import {
  createWorkspace, generatedCrateSource, loadThreadedRust, pboSource,
  resolveCaseArgs, runCase, writeCrate,
} from "./shared.mjs";

const { rust } = resolveCaseArgs();
const fixtures = join(dirname(fileURLToPath(import.meta.url)), "fixtures");

await runCase("qualified", rust, async ctx => {
  const threadedRust = await loadThreadedRust();
  const ffi = threadedRust(pboSource("PureScript/Backend/Optimizer/CoreFn.rs")).trim();
  const generated = generatedCrateSource(rust, "Purs_PureScript_Backend_Optimizer_CoreFn");
  assert.ok(generated.includes(ffi),
    "GENERATED_RUST is stale: CoreFn must embed the gopurs PBO native qualified comparison");
  for (const name of ["compareQualifiedIdentPS", "eqQualifiedIdentPS",
    "compareQualifiedIdentImpl", "eqQualifiedIdentImpl",
    "PureScript_Backend_Optimizer_CoreFn_compareQualifiedIdent(",
    "PureScript_Backend_Optimizer_CoreFn_eqQualifiedIdent("]) {
    assert.ok(generated.includes(name),
      `GENERATED_RUST is stale: CoreFn is missing ${name}`);
  }

  const directory = createWorkspace("gopurs-native-pbo-qualified");
  ctx.state.workspace = directory;
  ctx.note(`fixture workspace ${directory}`);
  const modules = [
    "purust_core", "Purs_Data_Maybe", "Purs_Data_Ordering", "Purs_Data_Ord", "Purs_Data_Eq",
    "Purs_PureScript_Backend_Optimizer_CoreFn",
  ];
  writeCrate(directory, "purust_native_qualified_test", modules, rust);
  writeFileSync(join(directory, "src/main.rs"),
    readFileSync(join(fixtures, "test-native-qualified.rs"), "utf8"));
  const executable = ctx.cargoBuild(directory, "purust_native_qualified_test");
  ctx.spawn("run", executable, []);
});
