import assert from "node:assert/strict";
import { createHash } from "node:crypto";
import { copyFileSync, mkdirSync, mkdtempSync, readFileSync, rmSync } from "node:fs";
import { tmpdir } from "node:os";
import { join } from "node:path";
import { test } from "node:test";
import { withReboxFields } from "./codegen-metadata.mjs";
import * as Aff from "../output/Effect.Aff/index.js";
import * as Either from "../output/Data.Either/index.js";
import * as Ref from "../output/Effect.Ref/index.js";
import * as Map from "../output/Data.Map/index.js";
import * as Set from "../output/Data.Set/index.js";
import * as List from "../output/Data.List/index.js";
import * as F from "../output/Data.Foldable/index.js";
import * as U from "../output/Data.Unfoldable/index.js";
import { ordString } from "../output/Data.Ord/index.js";
import { Just, Nothing } from "../output/Data.Maybe/index.js";
import * as C from "../output/PureScript.Backend.Optimizer.CoreFn/index.js";
import * as B from "../output/PureScript.Backend.Optimizer.Builder/index.js";
import * as Cache from "../output/Gopurs.Cache.Modules/index.js";

const hash = bytes => createHash("sha256").update(bytes).digest("hex");
const runAff = value => new Promise((resolve, reject) => Aff.runAff(result => () => result instanceof Either.Left ? reject(result.value0) : resolve(result.value0))(value)());
const asyncAff = run => Aff.makeAff(callback => () => {
  Promise.resolve().then(run).then(value => callback(new Either.Right(value))(), error => callback(new Either.Left(error))());
  return Aff.nonCanceler;
});
const nothing = Nothing.value, effect = Aff.monadEffectAff.liftEffect;
const insert = Map.insert(ordString);
const ann = { span: C.emptySpan, type: new Just(C.Int.value), meta: nothing, sourceUsage: nothing };
const lit = value => new C.ExprLit(ann, new C.LitInt(value));
const variable = (module, name) => new C.ExprVar(ann, new C.Qualified(new Just(module), name));
const moduleOf = (name, bindings, comments = []) => ({ name, path: name + ".purs", span: C.emptySpan, imports: [],
  exports: bindings.map(([name]) => name), reExports: [], dataDecls: [], classDecls: [], foreign: Map.empty, comments,
  decls: bindings.map(([name, expr]) => new C.NonRec(new C.Binding(ann, name, expr))) });
const fixture = ({ body = 10, spare = 1, other = 41, directive = false, missing = false, earlier = false } = {}) => {
  const modules = [
    moduleOf("Anchor", [["value", lit(0)]]),
    moduleOf("Ahead", [["value", variable("Later", "value")]]),
    moduleOf("Producer", [["value", lit(body)], ["spare", lit(spare)], ...(missing ? [["missing", lit(42)]] : [])],
      directive ? [new C.LineComment("@inline export value never")] : []),
    moduleOf("Other", [["value", lit(other)]]),
    moduleOf("Alias", [["value", variable("Producer", "value")]]),
    moduleOf("Consumer", [["value", variable("Alias", "value")]]),
    moduleOf("Absent", [["value", variable("Producer", "missing")]]),
    moduleOf("Later", [["value", lit(9)]]),
  ];
  if (earlier) modules.splice(1, 0, modules.pop());
  return modules;
};
const metadata = withReboxFields({ elidedCtors: Set.empty, ctorTypes: Map.empty, pointerAdtPaths: Map.empty,
  pointerAdtNodes: Set.empty, pointerAdtLeaves: Map.empty, enumAdts: Set.empty, enumCtors: Set.empty,
  globalTypes: Map.empty, globalFunctions: Map.empty, classDeclsFields: Map.empty });
const codecs = ["pbo/backend-v2", "pbo/implementations-v1", "pbo/directives-v1", "gopurs/workers-v2"];

test("observed PBO reuse matches fresh generation through body, directive, absence and rank transitions", async t => {
  const directory = mkdtempSync(join(tmpdir(), "gopurs-observed-build-")), cwd = process.cwd();
  process.chdir(directory);
  t.after(() => { process.chdir(cwd); rmSync(directory, { recursive: true, force: true }); });
  let sequence = 0;
  async function build(modules, order, previous) {
    const stage = join(directory, String(sequence++)), output = join(stage, "output");
    for (const path of ["modules", "output/purescript"]) mkdirSync(join(stage, path), { recursive: true });
    const cache = { session: { stage, output, context: hash("compiler/options") }, metadata: hash("metadata"),
      keys: Map.empty, candidates: Map.empty, sameContext: true, reads: Ref.new(Map.empty)(),
      restoredPBO: Ref.new(Set.empty)(), priorImplementations: Ref.new(Map.empty)(), currentDigests: Ref.new(Map.empty)(),
      priorDigests: Ref.new(Map.empty)(), invalidations: Ref.new(Map.empty)(), restored: Ref.new(0)() };
    let prefix = hash("compiler/options/directives");
    for (const mod of modules) {
      const prepared = hash(JSON.stringify(mod));
      cache.keys = insert(mod.name)({ prepared, callerDemands: hash("demands/" + prepared), optimizerEnvironment: prefix })(cache.keys);
      prefix = hash(prefix + "/" + prepared);
    }
    if (previous) for (const memo of previous.memos) {
      const path = join(stage, "prior", memo.name);
      mkdirSync(path, { recursive: true });
      for (let i = 0; i < 4; i++) copyFileSync(join(previous.stage, "modules", memo.name, i + ".json"), join(path, i + ".json"));
      copyFileSync(join(previous.stage, "output/purescript", memo.name + ".go"), join(path, "module.go"));
      cache.candidates = insert(memo.name)(memo)(cache.candidates);
    }
    const reused = [], emitted = [], queue = [];
    const publish = (mod, backend, cached) => asyncAff(async () => {
      (cached ? reused : emitted).push(mod.name);
      await runAff(Cache.emit(cache)(Map.empty)(metadata)("unused")(cached)(mod)(backend));
    });
    const observer = new Just({ onSkipModule: () => mod => view => Cache.lookupBackendWithReads(cache)(mod)(view),
      onReads: () => mod => reads => Cache.captureReads(cache)(mod)(reads) });
    const options = { directives: Map.empty, analyzeCustom: () => () => nothing, foreignSemantics: Map.empty,
      traceIdents: Set.empty, rewriteLimit: 10000, onPrepareModule: () => mod => effect(() => mod),
      onSkipModule: () => () => effect(() => nothing), onCodegenModule: () => mod => backend => () => publish(mod, backend, false) };
    const onReuse = () => mod => backend => publish(mod, backend, true), list = List.fromFoldable(F.foldableArray)(modules);
    if (order === "sequential") await runAff(B.buildModulesWithObservation(Aff.monadEffectAff)(observer)(onReuse)(options)(list));
    else await runAff(B.buildModulesParallelWithObservation(Aff.monadEffectAff)(observer)(onReuse)({ jobs: 8,
      scheduler: { fork: job => effect(() => { queue.push(job); }), await: asyncAff(() => runAff(queue.pop()(undefined))) }, onStats: nothing,
    })(options)(list));
    const memos = modules.map(mod => {
      const path = join(stage, "modules", mod.name), memo = JSON.parse(readFileSync(join(path, "memo.json")));
      memo.parts = codecs.map((codec, i) => { const bytes = readFileSync(join(path, i + ".json")); return { codec, object: { sha256: hash(bytes), bytes: bytes.length } }; });
      return memo;
    });
    const invalidations = Map.toUnfoldable(U.unfoldableArray)(Ref.read(cache.invalidations)()).map(pair => pair.value0);
    return { stage, memos, reused, emitted, invalidations,
      parts: Object.fromEntries(memos.map(memo => [memo.name, memo.parts])),
      go: Object.fromEntries(modules.map(mod => [mod.name, hash(readFileSync(join(output, "purescript", mod.name + ".go")))])) };
  }
  for (const order of ["sequential", "lifo"]) {
    let previous = await build(fixture(), order);
    for (const [label, settings, hits, misses] of [
      ["unrelated", { other: 43 }, ["Alias", "Consumer", "Absent"], ["Other"]],
      ["other-binding", { spare: 2 }, ["Alias", "Consumer", "Absent"], ["Producer"]],
      ["body", { body: 11 }, [], ["Producer", "Alias", "Consumer"]],
      ["directive", { body: 11, directive: true }, [], ["Producer", "Alias", "Consumer"]],
      ["appeared", { body: 11, directive: true, missing: true }, ["Alias", "Consumer"], ["Producer", "Absent"]],
      ["disappeared", { body: 11, directive: true }, ["Alias", "Consumer"], ["Producer", "Absent"]],
      ["rank", { body: 11, directive: true, earlier: true }, [], ["Ahead"]],
    ]) {
      const modules = fixture(settings), incremental = await build(modules, order, previous), fresh = await build(modules, "sequential");
      assert.deepEqual(incremental.go, fresh.go, order + "/" + label + ": Go");
      assert.deepEqual(incremental.parts, fresh.parts, order + "/" + label + ": business parts and read proofs");
      assert.deepEqual(incremental.invalidations.toSorted(), incremental.emitted.toSorted(), order + "/" + label + ": only final misses are reported");
      for (const name of hits) assert.ok(incremental.reused.includes(name), order + "/" + label + ": expected hit " + name);
      for (const name of misses) assert.ok(incremental.emitted.includes(name), order + "/" + label + ": expected miss " + name);
      previous = incremental;
    }
  }
});
