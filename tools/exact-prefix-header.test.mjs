import assert from "node:assert/strict";
import { createHash } from "node:crypto";
import { mkdirSync, mkdtempSync, writeFileSync, rmSync } from "node:fs";
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
import * as Ord from "../output/Data.Ord/index.js";
import { Just, Nothing } from "../output/Data.Maybe/index.js";
import * as C from "../output/PureScript.Backend.Optimizer.CoreFn/index.js";
import * as B from "../output/PureScript.Backend.Optimizer.Builder/index.js";
import * as CodeGen from "../output/Gopurs.CodeGen/index.js";
import { monadEffectEffect } from "../output/Effect.Class/index.js";
import * as Cache from "../output/Gopurs.Cache.Modules/index.js";

const nothing = Nothing.value;
const hash = value => createHash("sha256").update(value).digest("hex");
const runAff = value => new Promise((resolve, reject) => Aff.runAff(result => () => result instanceof Either.Left ? reject(result.value0) : resolve(result.value0))(value)());
const names = ["Anchor", "Earlier", "Future", "Reader"];
const namesSet = Set.fromFoldable(F.foldableArray)(Ord.ordString)(names);
const ann = { span: C.emptySpan, type: new Just(C.Int.value), meta: nothing, sourceUsage: nothing };
const lit = value => new C.ExprLit(ann, new C.LitInt(value));
const ref = name => new C.ExprVar(ann, new C.Qualified(new Just(name), "value"));
const mod = (name, value) => ({ name, path: name + ".purs", span: C.emptySpan, imports: [], exports: ["value"],
  reExports: [], dataDecls: [], classDecls: [], foreign: Map.empty, comments: [],
  decls: [new C.NonRec(new C.Binding(ann, "value", value))] });
const modules = value => [mod("Anchor", lit(1)), mod("Earlier", ref("Future")), mod("Future", lit(value)), mod("Reader", ref("Future"))];
const metadata = withReboxFields({ elidedCtors: Set.empty, ctorTypes: Map.empty, pointerAdtPaths: Map.empty,
  pointerAdtNodes: Set.empty, pointerAdtLeaves: Map.empty, enumAdts: Set.empty, enumCtors: Set.empty,
  globalTypes: Map.empty, globalFunctions: Map.empty, classDeclsFields: Map.empty });
const array = set => Set.toUnfoldable(U.unfoldableArray)(set);
const reads = Map.fromFoldable(Ord.ordString)(F.foldableArray)(names.map(name => ({
  value0: name, value1: new Just(name === "Earlier" || name === "Reader" ? ["Future"] : []),
})));

test("exact canonical prefixes restore earlier absent readers independently of read-closure eligibility", async t => {
  const stage = mkdtempSync(join(tmpdir(), "gopurs-exact-prefix-"));
  t.after(() => rmSync(stage, { recursive: true, force: true }));
  mkdirSync(join(stage, "prior"), { recursive: true });
  writeFileSync(join(stage, "prior/index.json"), "[]");
  const oldFfi = Map.singleton("Anchor")(hash("ffi-base"));
  const session = { stage, output: join(stage, "output"), context: hash("compiler/options"), sameTast: false };
  const prepared = (ms, meta = metadata, metadataUnchanged = false) => ({ modules: List.fromFoldable(F.foldableArray)(ms), metadata: meta, metadataUnchanged, directives: Map.empty });
  const initialize = (ms, settings = {}) => runAff(Cache.initialize(1)({ ...session, ...settings.session })
    (prepared(ms, settings.metadata ?? metadata, settings.metadataUnchanged ?? false))(Set.empty)(Ref.new(Map.empty)())
    (settings.ffi ?? oldFfi)(oldFfi)(settings.headerCandidates ?? namesSet));
  const base = await initialize(modules(10));
  const memos = Map.toUnfoldable(U.unfoldableArray)(base.keys).map(({ value0: name, value1: keys }) => ({
    name, ...keys, emitterEnvironment: hash("emission/" + name), metadata: base.metadata,
    ffiDigest: name === "Anchor" ? hash("ffi-base") : "no-foreign-v1", parts: [],
  }));
  writeFileSync(join(stage, "prior/index.json"), JSON.stringify(memos));
  for (const name of names) {
    mkdirSync(join(stage, "prior", name));
    writeFileSync(join(stage, "prior", name, "reads.json"), name === "Earlier" || name === "Reader" ? "Future\n" : "");
  }
  assert.deepEqual(array(Cache.headerSafeFromReads(Set.singleton("Future"))(reads)(namesSet)), ["Anchor"]);
  const changed = await initialize(modules(11));
  assert.deepEqual(array(changed.headerSafe), ["Anchor", "Earlier"]);

  const alreadyProven = { ...metadata };
  Object.defineProperty(alreadyProven, "globalTypes", { get() { throw new Error("unchanged metadata must not be encoded"); } });
  const trusted = await initialize(modules(11), { metadata: alreadyProven, metadataUnchanged: true });
  assert.equal(trusted.metadata, base.metadata);
  assert.deepEqual(array(trusted.headerSafe), ["Anchor", "Earlier"]);

  // Exact keys also prove unchanged output without a usable read-owner proof.
  rmSync(join(stage, "prior/Earlier/reads.json"));
  assert.deepEqual(array((await initialize(modules(11), { headerCandidates: Set.empty })).headerSafe), ["Anchor", "Earlier"]);

  // A matching optimizer prefix does not cover a changed emitter environment.
  const changedMetadata = { ...metadata, globalTypes: Map.singleton("Anchor.value")(C.String.value) };
  assert.deepEqual(array((await initialize(modules(11), { metadata: changedMetadata })).headerSafe), []);
  assert.deepEqual(array((await initialize(modules(11), { ffi: Map.singleton("Anchor")(hash("different-ffi")) })).headerSafe), []);
  assert.deepEqual(array((await initialize(modules(11), { session: { context: hash("other-compiler/options") } })).headerSafe), []);

  // Every complete optimizer key is required, including caller demands.
  const stale = memos.map(m => m.name === "Earlier" ? { ...m, callerDemands: hash("other-demands") } : m);
  writeFileSync(join(stage, "prior/index.json"), JSON.stringify(stale));
  assert.deepEqual(array((await initialize(modules(11), { headerCandidates: Set.empty })).headerSafe), ["Anchor"]);

  // A preparation proof cannot repair missing or inconsistent committed keys.
  for (const metadataKey of [undefined, hash("other-metadata")]) {
    writeFileSync(join(stage, "prior/index.json"), JSON.stringify(memos.map(m => m.name === "Earlier" ? { ...m, metadata: metadataKey } : m)));
    const fallback = await initialize(modules(11), { metadataUnchanged: true });
    assert.equal(fallback.metadata, base.metadata, "recompute the real metadata hash when committed keys disagree");
    assert.deepEqual(array(fallback.headerSafe), []);
  }
});

test("read-owner closure is complete with fanout, diamonds, cycles, missing proofs and duplicate edges", () => {
  const all = ["Change", "Left", "Right", "Join", "CycleA", "CycleB", "Missing", "MissingReader", "Independent"];
  const dependencies = { Change: [], Left: ["Change", "Change"], Right: ["Change"], Join: ["Left", "Right"],
    CycleA: ["Join", "CycleB"], CycleB: ["CycleA"], Missing: null, MissingReader: ["Missing"], Independent: [] };
  const proof = Map.fromFoldable(Ord.ordString)(F.foldableArray)(all.map(name => ({
    value0: name, value1: dependencies[name] === null ? nothing : new Just(dependencies[name]),
  })));
  const allSet = Set.fromFoldable(F.foldableArray)(Ord.ordString)(all);
  assert.deepEqual(array(Cache.headerSafeFromReads(Set.singleton("Change"))(proof)(allSet)), ["Independent"]);
});

test("an exact contribution stops transitive invalidation while uncertain owners remain conservative", async t => {
  const stage = mkdtempSync(join(tmpdir(), "gopurs-exact-closure-"));
  t.after(() => rmSync(stage, { recursive: true, force: true }));
  mkdirSync(join(stage, "prior"));
  writeFileSync(join(stage, "prior/index.json"), "[]");
  const corpus = value => [
    { ...mod("Exact", ref("Change")), comments: [new C.LineComment("@inline export value never")] },
    mod("Change", lit(value)), mod("Reader", ref("Exact")), mod("Missing", lit(0)), mod("MissingReader", ref("Missing")),
  ];
  const optimize = (source, cached = {}) => {
    const backends = {}, proofs = {};
    const publish = () => module => backend => () => { backends[module.name] = backend; };
    const observer = new Just({ onSkipModule: () => module => () => () => cached[module.name] ? new Just(cached[module.name]) : nothing,
      onReads: () => module => observed => () => {
        proofs[module.name] = [...new globalThis.Set([...Map.toUnfoldable(U.unfoldableArray)(observed.implementations), ...Map.toUnfoldable(U.unfoldableArray)(observed.directives)].map(entry => entry.value0.value0.value0))];
      } });
    B.buildModulesWithObservation(monadEffectEffect)(observer)(publish)({
      directives: Map.empty, analyzeCustom: () => () => nothing, foreignSemantics: Map.empty, traceIdents: Set.empty, rewriteLimit: 10000,
      onPrepareModule: () => module => () => module, onSkipModule: () => () => () => nothing,
      onCodegenModule: env => module => backend => () => publish(env)(module)(backend),
    })(List.fromFoldable(F.foldableArray)(source))();
    return { backends, proofs, code: Object.fromEntries(Object.entries(backends).map(([name, backend]) => [name, CodeGen.translateWithFunctions(metadata)(backend).code])) };
  };
  const original = corpus(10), edited = corpus(11), prior = optimize(original), fresh = optimize(edited);
  assert.ok(prior.proofs.Exact.includes("Change"));
  assert.ok(prior.proofs.Reader.includes("Exact"));
  assert.ok(!prior.proofs.Reader.includes("Change"), "a forbidden inline keeps the reader independent of the future value");
  const currentNames = Set.fromFoldable(F.foldableArray)(Ord.ordString)(original.map(module => module.name));
  const initialize = (source, settings = {}) => runAff(Cache.initialize(1)({ stage, context: settings.context ?? "same-context", sameTast: false })
    ({ modules: List.fromFoldable(F.foldableArray)(source), metadata: settings.metadata ?? metadata, metadataUnchanged: false, directives: Map.empty })
    (Set.empty)(Ref.new(Map.empty)())(settings.ffi ?? Map.empty)(Map.empty)(currentNames));
  const baseline = await initialize(original);
  const memos = Map.toUnfoldable(U.unfoldableArray)(baseline.keys).map(({ value0: name, value1: keys }) => ({
    name, ...keys, metadata: baseline.metadata, ffiDigest: "no-foreign-v1", emitterEnvironment: hash(name), parts: [],
  }));
  writeFileSync(join(stage, "prior/index.json"), JSON.stringify(memos));
  for (const name of original.map(module => module.name)) {
    mkdirSync(join(stage, "prior", name));
    if (name !== "Missing") writeFileSync(join(stage, "prior", name, "reads.json"), prior.proofs[name].join("\n"));
  }
  const changed = await initialize(edited);
  assert.deepEqual(array(changed.headerSafe), ["Exact", "Reader"], "the exact owner stops the changed→exact→reader chain; missing proof still propagates");
  const cached = Object.fromEntries(array(changed.headerSafe).map(name => [name, prior.backends[name]]));
  assert.deepEqual(optimize(edited, cached).code, fresh.code, "using the newly eligible headers reproduces fresh generated Go");
  rmSync(join(stage, "prior/Exact/reads.json"));
  assert.deepEqual(array((await initialize(edited)).headerSafe), ["Exact", "Reader"], "the exact owner does not need its weaker read proof");
  rmSync(join(stage, "prior/Reader/reads.json"));
  assert.deepEqual(array((await initialize(edited)).headerSafe), ["Exact"], "a missing reader proof never becomes a hit");
  for (const settings of [{ context: "other-context" }, { ffi: Map.singleton("Exact")("changed-ffi") },
    { metadata: { ...metadata, globalTypes: Map.singleton("Exact.value")(C.String.value) } }]) {
    assert.deepEqual(array((await initialize(edited, settings)).headerSafe), [], "context, FFI and metadata still gate the exact barrier");
  }
  // Remove exactness without changing prepared bodies: the changed seed must
  // reach Reader through Exact again, even though Reader itself is unchanged.
  writeFileSync(join(stage, "prior/Exact/reads.json"), "Change\n");
  writeFileSync(join(stage, "prior/Reader/reads.json"), "Exact\n");
  writeFileSync(join(stage, "prior/index.json"), JSON.stringify(memos.map(memo => memo.name === "Exact" ? { ...memo, callerDemands: hash("unknown-demand") } : memo)));
  assert.deepEqual(array((await initialize(edited)).headerSafe), [], "without exactness, the full transitive invalidation remains necessary");
});
