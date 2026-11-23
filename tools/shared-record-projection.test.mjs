import assert from "node:assert/strict";
import { spawnSync } from "node:child_process";
import { mkdirSync, mkdtempSync, rmSync, writeFileSync } from "node:fs";
import { tmpdir } from "node:os";
import { join } from "node:path";
import { fileURLToPath } from "node:url";
import test from "node:test";
import { Left } from "../output/Data.Either/index.js";
import { Cons } from "../output/Data.List.Types/index.js";
import * as Map from "../output/Data.Map/index.js";
import { Just, Nothing } from "../output/Data.Maybe/index.js";
import * as Set from "../output/Data.Set/index.js";
import { unfoldableArray } from "../output/Data.Unfoldable/index.js";
import * as Aff from "../output/Effect.Aff/index.js";
import * as Ref from "../output/Effect.Ref/index.js";
import * as Cache from "../output/Gopurs.Cache.Modules/index.js";
import * as LookupTable from "../output/Gopurs.Cache.LookupTable/index.js";
import * as Prepare from "../output/Gopurs.Driver.Prepare/index.js";
import * as Projection from "../output/Gopurs.Driver.Prepare.Projection/index.js";
import * as Proof from "../output/Gopurs.Driver.Prepare.Proof/index.js";
import { findTypedCompiler } from "./build/native-workspace.mjs";

const root = fileURLToPath(new URL("../", import.meta.url));
const runAff = (action) => new Promise((resolve, reject) =>
  Aff.runAff((result) => () => result instanceof Left ? reject(result.value0) : resolve(result.value0))(action)());
const array = (list) => {
  const values = [];
  for (; list instanceof Cons; list = list.value1) values.push(list.value0);
  return values;
};
const selected = (cache) => Set.toUnfoldable(unfoldableArray)(cache.prepared.metadata.sharedRecords);
const comparableMetadata = (cache) => {
  const { sharedRecords, ...metadata } = cache.prepared.metadata;
  return { ...metadata, globalTypes: LookupTable.completeOrFail(metadata.globalTypes) };
};
const moduleProof = (cache, name) => Proof.moduleProof(array(cache.sourceModules).find((module) => module.name === name))();

test("actual preparation and cache keys track the last callback and protected recursion", async (t) => {
  const workspace = mkdtempSync(join(tmpdir(), "gopurs-shared-record-proof-"));
  t.after(() => rmSync(workspace, { recursive: true, force: true }));
  const source = join(workspace, "src"), output = join(workspace, "output");
  mkdirSync(source);
  const compiler = findTypedCompiler(root, process.env.GOPURS_PURS);
  const names = ["Generic", "Shape", "Caller"];
  writeFileSync(join(source, "Generic.purs"), `module Generic (retain) where
retain :: forall a. a -> a
retain value = value
`);
  function compile(callback, recursion) {
    writeFileSync(join(source, "Shape.purs"), `module Shape (Payload(..), Packet, handler, loop) where
data Payload a = Value a | Text String
type Packet =
  { enabled :: Boolean
  , part :: { payload :: Payload Int, count :: Int }
  }
handler :: Packet -> Packet
handler packet = packet
loop :: Int -> Packet -> Packet
loop count packet = ${recursion ? "loop count packet" : "packet"}
`);
    writeFileSync(join(source, "Caller.purs"), `module Caller (saved) where
import Generic (retain)
import Shape (Packet, handler)
saved :: Packet -> Packet
${callback ? "saved = retain handler" : "saved packet = retain (handler packet)"}
`);
    const result = spawnSync(compiler, ["compile", "--codegen", "corefn", "--output", output,
      ...names.map((name) => join(source, name + ".purs"))],
    { cwd: workspace, encoding: "utf8", timeout: 30000 });
    assert.ifError(result.error);
    assert.equal(result.status, 0, result.stdout + result.stderr);
  }
  const full = () => runAff(Prepare.prepareModulesWithCache(output)(names.length)(1)(Nothing.value)(Map.empty));
  const delta = (previous) => runAff(Prepare.prepareModulesDelta(output)(names.length)(1)(Nothing.value)(Map.empty)(previous));
  async function metadataKey(label, cache) {
    const stage = join(workspace, label);
    mkdirSync(join(stage, "prior"), { recursive: true });
    writeFileSync(join(stage, "prior/index.json"), "[]");
    const session = { root: workspace, cacheRoot: join(workspace, ".gopurs-cache"), stage,
      context: "shared-record-proof", sameTast: false };
    const initialized = await runAff(Cache.initialize(1)(session)(cache.prepared)(Set.empty)
      (Ref.new(Map.empty)())(Map.empty)(Map.empty)(Set.empty));
    return initialized.metadata;
  }
  async function edited(previous, callback, recursion) {
    compile(callback, recursion);
    const current = await delta(previous), expected = await full();
    assert.deepEqual(current.prepared.metadata, expected.prepared.metadata,
      "resident preparation agrees with a full preparation after the body edit");
    assert.deepEqual(comparableMetadata(current), comparableMetadata(previous),
      "declarations and published types are unchanged; only representation evidence changes");
    return current;
  }

  compile(false, false);
  const native = await full();
  assert.deepEqual(selected(native), []);
  const nativeKey = await metadataKey("native", native);

  const callback = await edited(native, true, false);
  assert.equal(selected(callback).length, 1);
  const callbackKey = await metadataKey("callback", callback);
  assert.notEqual(callbackKey, nativeKey, "the actual global metadata fingerprint includes selection");
  assert.equal(Proof.permitsLiteralEdit(Set.empty)(moduleProof(native, "Caller"))(moduleProof(callback, "Caller")), false,
    "a callback edit cannot reuse the literal-only preparation proof");

  const projection = { modules: callback.prepared.modules, metadata: callback.prepared.metadata,
    interfaces: callback.sourceInterfaces, mainModules: ["Caller"] };
  const fragments = Projection.split(projection);
  assert.equal(fragments.filter((fragment) => !Set.isEmpty(fragment.value.metadata.sharedRecords)).length, 1);
  const merged = Projection.merge(fragments);
  assert.ok(merged instanceof Just);
  assert.deepEqual(merged.value0.metadata.sharedRecords, callback.prepared.metadata.sharedRecords,
    "an effectively selected family survives the projection transport");

  const protectedLoop = await edited(callback, true, true);
  assert.deepEqual(selected(protectedLoop), []);
  assert.equal(await metadataKey("protected", protectedLoop), nativeKey,
    "a protected recursive accumulator restores the native representation fingerprint");
  assert.equal(Proof.permitsLiteralEdit(Set.empty)(moduleProof(callback, "Shape"))(moduleProof(protectedLoop, "Shape")), false,
    "a new Rec group cannot reuse the literal-only preparation proof");

  const loopRemoved = await edited(protectedLoop, true, false);
  assert.deepEqual(selected(loopRemoved), selected(callback));
  assert.equal(await metadataKey("loop-removed", loopRemoved), callbackKey);

  const callbackRemoved = await edited(loopRemoved, false, false);
  assert.deepEqual(selected(callbackRemoved), []);
  assert.equal(await metadataKey("callback-removed", callbackRemoved), nativeKey);
});
