import assert from "node:assert/strict";
import { test } from "node:test";
import * as Aff from "../output/Effect.Aff/index.js";
import * as Either from "../output/Data.Either/index.js";
import * as Ref from "../output/Effect.Ref/index.js";
import * as Map from "../output/Data.Map/index.js";
import * as Set from "../output/Data.Set/index.js";
import * as List from "../output/Data.List/index.js";
import * as F from "../output/Data.Foldable/index.js";
import * as U from "../output/Data.Unfoldable/index.js";
import * as Ord from "../output/Data.Ord/index.js";
import { Just } from "../output/Data.Maybe/index.js";
import { Tuple } from "../output/Data.Tuple/index.js";
import * as C from "../output/PureScript.Backend.Optimizer.CoreFn/index.js";
import * as S from "../output/PureScript.Backend.Optimizer.Semantics/index.js";
import * as Cache from "../output/Gopurs.Cache.Modules/index.js";
import * as Wire from "../output/Gopurs.Cache.Wire/index.js";

const run = value => new Promise((resolve, reject) => Aff.runAff(result => () => result instanceof Either.Left ? reject(result.value0) : resolve(result.value0))(value)());
const table = entries => Map.fromFoldable(Ord.ordString)(F.foldableArray)(entries.map(([name, value]) => new Tuple(name, value)));
const names = map => Set.toUnfoldable(U.unfoldableArray)(Map.keys(map));
const set = values => Set.fromFoldable(F.foldableArray)(Ord.ordString)(values);
const codec = Wire.codecMap(S.ordEvalRef)(Wire.keyOrderEvalRef)(Wire.codecEvalRef)(Wire.codecMap(S.ordInlineAccessor)(Wire.keyOrderInlineAccessor)(Wire.codecInlineAccessor)(Wire.codecInlineDirective));

test("batch headers require the same owner, prepared, FFI and business-codec proofs as individual lookup", async () => {
  const all = ["Alpha", "Beta", "WrongOwner", "Invalid", "Changed", "FfiChanged", "Unsafe"];
  const modules = all.map(name => ({ name, foreign: Map.empty }));
  const list = List.fromFoldable(F.foldableArray)(modules);
  const keys = name => ({ prepared: name, callerDemands: name + "-callers", optimizerEnvironment: "prefix" });
  const foreignDirective = Map.singleton(new S.EvalExtern(new C.Qualified(new Just("Alpha"), "value")))(Map.empty);
  const texts = table(all.map(name => [name, table([["2", name === "Invalid" ? "{}" : Wire.write(codec)(name === "WrongOwner" ? foreignDirective : Map.empty)()]])]));
  for (const jobs of [1, 8]) {
    const cache = {
      lazyContributions: true,
      headerSafe: set(all.filter(name => name !== "Unsafe")),
      candidates: table(all.map(name => [name, { ...keys(name), ffiDigest: new Just("no-foreign-v1") }])),
      keys: table(all.map(name => [name, { ...keys(name), prepared: name === "Changed" ? "changed" : name }])),
      ffiDigests: Map.singleton("FfiChanged")("changed-ffi"), priorFfiDigests: Map.empty,
      priorPartTexts: texts, restoredPBO: Ref.new(Set.empty)(), invalidations: Ref.new(table(all.map(name => [name, "old"])))(),
      priorImplementations: Ref.new(Map.empty)(), priorWorkers: Ref.new(Map.empty)(), backends: Ref.new(Map.empty)(),
    };
    const restored = await run(Cache.lookupHeaders(jobs)(cache)(list));
    assert.deepEqual(names(restored), ["Alpha", "Beta"]);
    assert.deepEqual(Set.toUnfoldable(U.unfoldableArray)(Ref.read(cache.restoredPBO)()), ["Alpha", "Beta"]);
    assert.deepEqual(names(Ref.read(cache.invalidations)()), ["Changed", "FfiChanged", "Invalid", "Unsafe", "WrongOwner"]);
    assert.deepEqual(names(Ref.read(cache.priorImplementations)()), [], "batch headers never load implementation bodies");
    assert.deepEqual(names(Ref.read(cache.priorWorkers)()), [], "batch headers never publish worker signatures");
    assert.deepEqual(names(Ref.read(cache.backends)()), [], "restored headers remain separate from emitted modules");
    for (const { value0: name, value1: backend } of Map.toUnfoldable(U.unfoldableArray)(restored)) {
      assert.equal(backend.name, name);
      assert.deepEqual(backend.bindings, []);
      assert.deepEqual(names(backend.implementations), []);
    }
    assert.deepEqual(names(await run(Cache.lookupHeaders(jobs)({ ...cache, lazyContributions: false })(list))), [], "eager/jobs=1 callers retain the ordinary path");
  }
});
