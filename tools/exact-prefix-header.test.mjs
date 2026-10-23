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
  const prepared = (ms, meta = metadata) => ({ modules: List.fromFoldable(F.foldableArray)(ms), metadata: meta, directives: Map.empty });
  const initialize = (ms, settings = {}) => runAff(Cache.initialize(1)({ ...session, ...settings.session })
    (prepared(ms, settings.metadata ?? metadata))(Set.empty)(Ref.new(Map.empty)())
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
