import assert from "node:assert/strict";
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
import * as Ord from "../output/Data.Ord/index.js";
import { Just, Nothing } from "../output/Data.Maybe/index.js";
import * as C from "../output/PureScript.Backend.Optimizer.CoreFn/index.js";
import * as B from "../output/PureScript.Backend.Optimizer.Builder/index.js";
import * as Convert from "../output/PureScript.Backend.Optimizer.Convert/index.js";
import * as CodeGen from "../output/Gopurs.CodeGen/index.js";

const nothing = Nothing.value, effect = Aff.monadEffectAff.liftEffect;
const runAff = value => new Promise((resolve, reject) => Aff.runAff(result => () => result instanceof Either.Left ? reject(result.value0) : resolve(result.value0))(value)());
const asyncAff = run => Aff.makeAff(callback => () => {
  Promise.resolve().then(run).then(value => callback(new Either.Right(value))(), error => callback(new Either.Left(error))());
  return Aff.nonCanceler;
});
const ann = { span: C.emptySpan, type: new Just(C.Int.value), meta: nothing, sourceUsage: nothing };
const lit = value => new C.ExprLit(ann, new C.LitInt(value));
const variable = (owner, name) => new C.ExprVar(ann, new C.Qualified(new Just(owner), name));
const moduleOf = (name, bindings, comments = []) => ({ name, path: name + ".purs", span: C.emptySpan, imports: [],
  exports: bindings.map(([name]) => name), reExports: [], dataDecls: [], classDecls: [], foreign: Map.empty, comments,
  decls: bindings.map(([name, expr]) => new C.NonRec(new C.Binding(ann, name, expr))) });
const metadata = withReboxFields({ elidedCtors: Set.empty, ctorTypes: Map.empty, pointerAdtPaths: Map.empty,
  pointerAdtNodes: Set.empty, pointerAdtLeaves: Map.empty, enumAdts: Set.empty, enumCtors: Set.empty,
  globalTypes: Map.empty, globalFunctions: Map.empty, classDeclsFields: Map.empty });
const fixtures = (reference = ["Provider", "value"], never = false) => [
  moduleOf("Unused", [["value", lit(999)]]),
  moduleOf("Provider", [["value", lit(10)], ["spare", lit(20)]], never ? [new C.LineComment("@inline export value never")] : []),
  moduleOf("Caller", [["value", variable(...reference)]]),
  moduleOf("Future", [["value", lit(500)]]),
];
const header = backend => ({ ...backend, implementations: Map.empty, bindings: [], imports: Set.empty, exports: Set.empty });
async function build(modules, mode, cached = {}, { prepareFresh = mod => mod, cacheViaObservation = false, schedulingHints = {}, restoredHeaders = null } = {}) {
  const queue = [], reused = [], emitted = [], published = [], environments = [], loaded = [], prepared = [], attempts = [], observations = {}, result = {}, code = {}, publishedModules = {}, availability = {};
  let globals = Map.empty, stats;
  const publish = (mod, backend, fromCache, env) => effect(() => {
    const complete = fromCache ? cached[mod.name] : backend;
    result[mod.name] = complete;
    published.push(mod.name);
    environments.push({ name: mod.name, moduleCount: env.moduleCount, moduleIndex: env.moduleIndex });
    publishedModules[mod.name] = mod;
    (fromCache ? reused : emitted).push(mod.name);
    const translated = CodeGen.translateWithFunctions({ ...metadata, globalFunctions: globals })(complete);
    code[mod.name] = translated.code;
    globals = Map.union(Ord.ordString)(translated.functions)(globals);
  });
  const skip = mod => effect(() => cached[mod.name] ? new Just(mode === "lazy" ? header(cached[mod.name]) : cached[mod.name]) : nothing);
  const options = { directives: Map.empty, analyzeCustom: () => () => nothing, foreignSemantics: Map.empty,
    traceIdents: Set.empty, rewriteLimit: 10000, onPrepareModule: () => mod => effect(() => { attempts.push(mod.name); return mod; }),
    onSkipModule: () => mod => cacheViaObservation ? effect(() => nothing) : skip(mod),
    onCodegenModule: env => mod => backend => () => publish(mod, backend, false, env) };
  const observer = new Just({ onSkipModule: () => mod => view => {
    const before = loaded.length;
    (availability[mod.name] ??= []).push(Object.fromEntries([...modules.map(value => value.name), "Outside"].map(name => [name, view.availableOwner(name)])));
    assert.equal(loaded.length, before, "availability must not force a contribution");
    return skip(mod);
  }, onReads: () => mod => reads => effect(() => { observations[mod.name] = reads; }) });
  const onReuse = env => mod => backend => publish(mod, backend, true, env);
  const list = List.fromFoldable(F.foldableArray)(modules);
  if (mode === "full") await runAff(B.buildModulesWithObservation(Aff.monadEffectAff)(observer)(onReuse)(options)(list));
  else {
    const runner = { jobs: 8, scheduler: {
      fork: job => effect(() => { queue.push(job); }),
      await: asyncAff(() => runAff(queue.pop()(undefined))),
    }, onStats: new Just(value => effect(() => { stats = value; })) };
    const loader = new Just(name => () => { loaded.push(name); assert.ok(cached[name], "only finalized cache owners may load"); return cached[name].implementations; });
    const prepare = mod => effect(() => { prepared.push(mod.name); return prepareFresh(mod); });
    const hints = Map.fromFoldable(Ord.ordString)(F.foldableArray)(Object.entries(schedulingHints).map(([value0, value1]) => ({ value0, value1 })));
    const restored = restoredHeaders === null ? null : Map.fromFoldable(Ord.ordString)(F.foldableArray)(Object.entries(restoredHeaders).map(([value0, backend]) => ({ value0, value1: header(backend) })));
    const parallel = restored !== null
      ? B.buildModulesParallelWithRestoredHeaders(Aff.monadEffectAff)(restored)(hints)
      : Object.keys(schedulingHints).length === 0
      ? B.buildModulesParallelWithLazyPreparation(Aff.monadEffectAff)
      : B.buildModulesParallelWithSchedulingHints(Aff.monadEffectAff)(hints);
    await runAff(parallel(prepare)(mode === "lazy" ? loader : nothing)(observer)(onReuse)(runner)(options)(list));
  }
  return { result, code, reused, emitted, published, environments, loaded, prepared, attempts, observations, stats, publishedModules, availability };
}

test("lazy contributions reproduce fresh and eager output, including absent/future reads and directives", async () => {
  for (const never of [false, true]) {
    const base = await build(fixtures(undefined, never), "full");
    const cached = Object.fromEntries(Object.entries(base.result).filter(([name]) => name !== "Caller"));
    for (const reference of [["Provider", "value"], ["Provider", "spare"], ["Provider", "missing"], ["Outside", "missing"], ["Future", "value"]]) {
      const modules = fixtures(reference, never);
      const full = await build(modules, "full"), eager = await build(modules, "eager", cached), lazy = await build(modules, "lazy", cached);
      assert.deepEqual(lazy.code, full.code, `${reference}/${never}: lazy versus fresh Go`);
      assert.deepEqual(lazy.code, eager.code, `${reference}/${never}: lazy versus eager Go`);
      assert.deepEqual(lazy.observations.Caller, full.observations.Caller, `${reference}/${never}: exact reads`);
      assert.deepEqual(lazy.reused, ["Unused", "Provider", "Future"], "reuse publication remains canonical under LIFO scheduling");
      assert.deepEqual(lazy.emitted, ["Caller"]);
      assert.ok(!lazy.loaded.includes("Unused"), "unread predecessor must never decode");
      assert.ok(!lazy.loaded.includes("Future"), "future rank must never invoke its loader");
      if (reference[0] === "Provider" && !(never && reference[1] === "value")) assert.ok(lazy.loaded.includes("Provider"));
      if (reference[0] !== "Provider") assert.deepEqual(lazy.loaded, []);
      for (const visible of lazy.availability.Caller) {
        assert.equal(visible.Caller, false, "current owner remains invisible");
        assert.equal(visible.Future, false, "future owner remains invisible even if its worker completed");
        assert.equal(visible.Outside, false, "unknown owner remains invisible");
      }
    }
  }
});

test("rank lookup preserves absent versus pending and finality", async () => {
  const index = Map.fromFoldable(Ord.ordString)(F.foldableArray)([
    { value0: "Earlier", value1: 0 }, { value0: "Current", value1: 1 }, { value0: "Future", value1: 2 },
  ]);
  const pending = Ref.new(Set.empty)();
  const lookup = await runAff(B.createRankLookup(Aff.monadEffectAff)({ indexByName: index, currentIndex: 1, finalized: Map.empty })(pending));
  assert.ok(lookup("Outside")("value") instanceof Convert.ExternMissing);
  assert.ok(lookup("Current")("value") instanceof Convert.ExternMissing);
  assert.ok(lookup("Future")("value") instanceof Convert.ExternMissing);
  assert.deepEqual(Set.toUnfoldable(U.unfoldableArray)(Ref.read(pending)()), []);
  assert.ok(lookup("Earlier")("value") instanceof Convert.ExternPending);
  assert.deepEqual(Set.toUnfoldable(U.unfoldableArray)(Ref.read(pending)()), [0]);
  const finalized = Map.singleton(0)(Map.empty);
  const complete = await runAff(B.createRankLookup(Aff.monadEffectAff)({ indexByName: index, currentIndex: 1, finalized })(Ref.new(Set.empty)()));
  assert.ok(complete("Earlier")("value") instanceof Convert.ExternMissing);
});

test("a lookup introduced by a cached bridge defers until its earlier owner is finalized", async () => {
  // Bridge cannot see Hidden at its canonical rank, so its reusable optimized
  // contribution still references Hidden. Caller discovers that indirect read
  // while the LIFO scheduler has Hidden in flight; this must not mean absent.
  const modules = [
    moduleOf("Unused", [["value", lit(999)]]),
    moduleOf("Bridge", [["value", variable("Hidden", "value")]]),
    moduleOf("Hidden", [["value", lit(17)]]),
    moduleOf("Caller", [["value", variable("Bridge", "value")]]),
  ];
  const full = await build(modules, "full");
  const cached = Object.fromEntries(Object.entries(full.result).filter(([name]) => name !== "Caller"));
  const eager = await build(modules, "eager", cached), lazy = await build(modules, "lazy", cached);
  assert.deepEqual(lazy.code, full.code);
  assert.deepEqual(lazy.code, eager.code);
  assert.ok(lazy.stats.deferredAttempts > 0, "pending indirect predecessor must replay the conversion");
  assert.ok(lazy.loaded.includes("Bridge"));
  assert.ok(lazy.loaded.includes("Hidden"));
  assert.ok(!lazy.loaded.includes("Unused"));
  assert.deepEqual(lazy.reused, ["Unused", "Bridge", "Hidden"]);
  assert.deepEqual(lazy.emitted, ["Caller"], "deferred attempts publish no partial result");
  assert.equal(lazy.availability.Caller[0].Bridge, true, "the committed cached bridge is visible without loading");
  assert.equal(lazy.availability.Caller[0].Hidden, false, "an earlier owner still in flight is not available");
  assert.equal(lazy.availability.Caller.at(-1).Hidden, true, "replay sees the newly finalized owner");
});

test("lazy preparation leaves cache-hit placeholders inert and converts the restored AST on misses and replay", async () => {
  const modules = [
    moduleOf("Unused", [["value", lit(999)]]),
    moduleOf("Bridge", [["value", variable("Hidden", "value")]]),
    moduleOf("Hidden", [["value", lit(17)]]),
    moduleOf("Caller", [["value", variable("Bridge", "value")]]),
  ];
  const full = await build(modules, "full");
  const cached = Object.fromEntries(Object.entries(full.result).filter(([name]) => name !== "Caller"));
  const placeholders = modules.map(mod => moduleOf(mod.name, [["value", lit(-999)]]));
  for (const cacheViaObservation of [false, true]) {
    const lazy = await build(placeholders, "lazy", cached, {
      cacheViaObservation,
      prepareFresh: placeholder => {
        assert.equal(placeholder.name, "Caller", "neither exact nor observed cache hits prepare their inert AST");
        return modules.find(mod => mod.name === placeholder.name);
      },
    });
    assert.deepEqual(lazy.code, full.code, "conversion must use the callback's actual AST, not the placeholder");
    assert.ok(lazy.stats.deferredAttempts > 0, "restored indirect lookup still waits for its predecessor");
    assert.ok(lazy.prepared.length > 1, "the deferred miss is retried");
    assert.ok(lazy.prepared.every(name => name === "Caller"));
    assert.deepEqual(lazy.publishedModules.Caller, modules[3], "publication receives the actual prepared module");
    assert.deepEqual(lazy.reused, ["Unused", "Bridge", "Hidden"]);
    assert.deepEqual(lazy.emitted, ["Caller"], "replayed attempts publish only once");
  }
});


test("committed read hints avoid placeholder replay without changing canonical visibility", async () => {
  const modules = [
    moduleOf("Unused", [["value", lit(999)]]),
    moduleOf("Bridge", [["value", variable("Hidden", "value")]]),
    moduleOf("Hidden", [["value", lit(17)]]),
    moduleOf("Caller", [["value", variable("Bridge", "value")]]),
    moduleOf("Future", [["value", lit(500)]]),
  ];
  const full = await build(modules, "full");
  const cached = Object.fromEntries(Object.entries(full.result).filter(([name]) => name !== "Caller"));
  const placeholders = modules.map(mod => moduleOf(mod.name, [["value", lit(-999)]]));
  const prepareFresh = placeholder => modules.find(mod => mod.name === placeholder.name);
  const unhinted = await build(placeholders, "lazy", cached, { prepareFresh });
  assert.ok(unhinted.stats.deferredAttempts > 0, "the missing placeholder edges must exercise replay");
  for (const owners of [["Bridge", "Hidden"], ["Bridge", "Hidden", "Hidden", "Unused", "Caller", "Future", "Outside"]]) {
    const hinted = await build(placeholders, "lazy", cached, { prepareFresh, schedulingHints: { Caller: owners, Outside: ["Caller"] } });
    assert.equal(hinted.stats.deferredAttempts, 0, "hints wait for observed predecessor owners before the first attempt");
    assert.deepEqual(hinted.prepared, ["Caller"], "one conversion after all predecessors finalize");
    assert.deepEqual(hinted.code, full.code);
    assert.deepEqual(hinted.observations.Caller, full.observations.Caller);
    assert.deepEqual(hinted.reused, ["Unused", "Bridge", "Hidden", "Future"]);
    assert.deepEqual(hinted.emitted, ["Caller"]);
    assert.ok(!hinted.loaded.includes("Unused"), "a stale edge may delay but must not load a contribution");
    assert.ok(!hinted.loaded.includes("Future"), "a future edge never grants visibility");
    assert.equal(hinted.availability.Caller[0].Hidden, true);
    assert.equal(hinted.availability.Caller[0].Future, false);
    assert.equal(hinted.availability.Caller[0].Outside, false);
  }
  for (const owner of ["Hidden", "Future", "Outside"]) {
    const source = modules.map(mod => mod.name === "Caller" ? moduleOf("Caller", [["value", variable(owner, "missing")]]) : mod);
    const fresh = await build(source, "full");
    const hinted = await build(placeholders, "lazy", cached, {
      prepareFresh: mod => source.find(value => value.name === mod.name), schedulingHints: { Caller: [owner] },
    });
    assert.deepEqual(hinted.code, fresh.code, `${owner}: hints preserve missing binding/owner results`);
    assert.deepEqual(hinted.observations.Caller, fresh.observations.Caller);
    assert.ok(!hinted.loaded.includes("Future"));
  }
});

test("batch-decoded headers finalize only at the canonical front and preserve fresh/eager output", async () => {
  for (const never of [false, true]) {
    const base = await build(fixtures(undefined, never), "full");
    const cached = Object.fromEntries(Object.entries(base.result).filter(([name]) => name !== "Caller"));
    for (const reference of [["Provider", "value"], ["Provider", "missing"], ["Outside", "missing"], ["Future", "value"]]) {
      const modules = fixtures(reference, never);
      const full = await build(modules, "full"), lazy = await build(modules, "lazy", cached);
      const restored = await build(modules, "lazy", cached, { restoredHeaders: cached });
      assert.deepEqual(restored.code, full.code);
      assert.deepEqual(restored.code, lazy.code);
      assert.deepEqual(restored.observations.Caller, full.observations.Caller);
      assert.deepEqual(restored.reused, ["Unused", "Provider", "Future"]);
      assert.deepEqual(restored.emitted, ["Caller"]);
      assert.deepEqual(restored.published, modules.map(mod => mod.name));
      assert.deepEqual(restored.environments, modules.map((mod, index) => ({ name: mod.name, moduleCount: modules.length, moduleIndex: index })));
      assert.equal(restored.stats.dispatched + restored.stats.fallbackDispatched, 1);
      assert.deepEqual(restored.attempts, modules.map(mod => mod.name), "each cached preparation occurs at its canonical front");
      assert.deepEqual(restored.prepared, ["Caller"]);
      assert.ok(!restored.loaded.includes("Unused") && !restored.loaded.includes("Future"));
      assert.equal(restored.availability.Caller[0].Future, false, "a batch-decoded future header is not finalized");
      assert.equal(restored.availability.Caller[0].Provider, true);
    }
  }
});

test("mixed restored ranges preserve pending replay and publish once with five uncertain modules", async () => {
  const modules = [
    moduleOf("Unused", [["value", lit(999)]]),
    moduleOf("Bridge", [["value", variable("Hidden", "value")]]),
    moduleOf("Hidden", [["value", lit(17)]]),
    moduleOf("CachedGap", [["value", lit(23)]]),
    moduleOf("Caller", [["value", variable("Bridge", "value")]]),
    moduleOf("Fresh2", [["value", variable("Caller", "value")]]),
    moduleOf("Fresh3", [["value", variable("Future", "value")]]),
    moduleOf("Fresh4", [["value", variable("Outside", "missing")]]),
    moduleOf("Future", [["value", lit(500)]]),
  ];
  const full = await build(modules, "full");
  const names = ["Unused", "Bridge", "CachedGap", "Future"];
  const cached = Object.fromEntries(names.map(name => [name, full.result[name]]));
  const restored = await build(modules, "lazy", cached, { restoredHeaders: cached });
  assert.deepEqual(restored.code, full.code);
  assert.deepEqual(restored.observations, Object.fromEntries(Object.entries(full.observations).filter(([name]) => !names.includes(name))));
  assert.deepEqual(restored.reused, names);
  assert.deepEqual(restored.emitted, ["Hidden", "Caller", "Fresh2", "Fresh3", "Fresh4"]);
  assert.deepEqual(restored.published, modules.map(mod => mod.name));
  assert.deepEqual(restored.environments, modules.map((mod, index) => ({ name: mod.name, moduleCount: modules.length, moduleIndex: index })));
  assert.ok(restored.stats.deferredAttempts > 0, "LIFO worker discovery of an unfinished predecessor still replays");
  assert.ok(!restored.loaded.includes("Future"));
  assert.ok(restored.availability.Caller.some(view => !view.Hidden), "a fresh owner in flight remains pending");
  assert.equal(restored.availability.Caller.at(-1).Hidden, true);
});
