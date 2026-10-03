// Persistent native maps must interoperate with generated generic operations.
import assert from "node:assert/strict";
import { writeFileSync } from "node:fs";
import { join } from "node:path";
import {
  createWorkspace, fixtureSource, generatedCrateSource, loadThreadedRust,
  pboSource, resolveCaseArgs, runCase, writeCrate,
} from "./shared.mjs";

const { rust } = resolveCaseArgs();
const threadedRust = await loadThreadedRust();
await runCase("maps", rust, async ctx => {
  const ffi = threadedRust(pboSource("PureScript/Backend/Optimizer/NativeMaps.rs"));
  assert.ok(generatedCrateSource(rust, "Purs_PureScript_Backend_Optimizer_NativeMaps").includes(ffi.trim()));
  const directory = createWorkspace("gopurs-native-pbo-maps");
  ctx.state.workspace = directory;
  writeCrate(directory, "purust_native_maps_test", [
    "purust_core", "Purs_Data_Map_Internal", "Purs_Data_Maybe", "Purs_Data_Ord", "Purs_Data_Ordering",
    "Purs_PureScript_Backend_Optimizer_CoreFn", "Purs_PureScript_Backend_Optimizer_NativeMaps",
  ], rust);
  writeFileSync(join(directory, "src/main.rs"),
    fixtureSource("test-native-maps.rs").replace(/^\s*\/\/ NATIVE_FFI$/m, () => ffi));
  const binary = ctx.cargoBuild(directory, "purust_native_maps_test");
  ctx.spawn("run", binary, []);
});
