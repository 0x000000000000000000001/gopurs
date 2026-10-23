import assert from "node:assert/strict";
import { createHash } from "node:crypto";
import { mkdirSync, mkdtempSync, rmSync, writeFileSync } from "node:fs";
import { tmpdir } from "node:os";
import { join } from "node:path";
import test from "node:test";
import * as Aff from "../output/Effect.Aff/index.js";
import { monadEffectEffect } from "../output/Effect.Class/index.js";
import * as Ref from "../output/Effect.Ref/index.js";
import { Left } from "../output/Data.Either/index.js";
import * as Foldable from "../output/Data.Foldable/index.js";
import * as List from "../output/Data.List/index.js";
import * as Map from "../output/Data.Map/index.js";
import { Just, Nothing } from "../output/Data.Maybe/index.js";
import { ordString } from "../output/Data.Ord/index.js";
import * as Set from "../output/Data.Set/index.js";
import { unfoldableArray } from "../output/Data.Unfoldable/index.js";
import * as C from "../output/PureScript.Backend.Optimizer.CoreFn/index.js";
import * as Builder from "../output/PureScript.Backend.Optimizer.Builder/index.js";
import * as Cache from "../output/Gopurs.Cache.Modules/index.js";
import * as CodeGen from "../output/Gopurs.CodeGen/index.js";
import * as Manifest from "../output/Gopurs.Driver.Manifest/index.js";
import { buildGlobalTypes } from "../output/Gopurs.GlobalTypes/index.js";
import { monomorphizeModules } from "../output/Gopurs.Monomorphization/index.js";
import { withReboxFields } from "./codegen-metadata.mjs";

const nothing = Nothing.value;
const list = List.fromFoldable(Foldable.foldableArray);
const array = List.toUnfoldable(unfoldableArray);
const set = Set.fromFoldable(Foldable.foldableArray)(ordString);
const setValues = Set.toUnfoldable(unfoldableArray);
const entries = value => Map.toUnfoldable(unfoldableArray)(value).map(pair => [pair.value0, pair.value1]);
const map = values => values.reduce((result, [name, value]) => Map.insert(ordString)(name)(value)(result), Map.empty);
const hash = value => createHash("sha256").update(JSON.stringify(value)).digest("hex");
const runAff = value => new Promise((resolve, reject) => Aff.runAff(result => () => result instanceof Left ? reject(result.value0) : resolve(result.value0))(value)());
const int = C.Int.value, string = C.String.value;
const poly = new C.ForAll(["a"], new C.Func([new C.TypeVar("a")], int));
const ann = type => ({ span: C.emptySpan, meta: nothing, sourceUsage: nothing, type: new Just(type) });
const literal = (value, type = int) => new C.ExprLit(ann(type), type === int ? new C.LitInt(value) : new C.LitString(value));
const binding = (name, type, body) => new C.NonRec(new C.Binding(ann(type), name, body));
const moduleOf = (name, decls, imports = []) => ({ name, path: name + ".purs", span: C.emptySpan,
  imports: imports.map(owner => new C.Import(ann(int), owner)), comments: [], decls,
  exports: decls.map(group => group.value0.value1), reExports: [], dataDecls: [], classDecls: [], foreign: Map.empty });
const call = (value, type = int) => new C.ExprApp(ann(int), new C.ExprVar(ann(poly), new C.Qualified(new Just("Generic"), "ignore")), literal(value, type));
const fixture = stringCall => [
  moduleOf("Generic", [binding("ignore", poly, new C.ExprAbs(ann(poly), "value", literal(42)))]),
  moduleOf("Caller", [binding("answer", int, stringCall ? call("hello", string) : call(1))], ["Generic"]),
  moduleOf("Downstream", [binding("answer", int, call(2))], ["Generic"]),
  moduleOf("Unrelated", [binding("answer", int, literal(7))]),
];
const prepare = modules => array(monomorphizeModules(buildGlobalTypes(modules))(list(modules)));
const manifest = modules => Manifest.buildManifest("")("same-context")(Map.empty)(modules)(map(modules.map(mod => [mod.name, hash(mod)])))();
const metadataFor = modules => withReboxFields({ elidedCtors: Set.empty, ctorTypes: Map.empty,
  pointerAdtPaths: Map.empty, pointerAdtNodes: Set.empty, pointerAdtLeaves: Map.empty, enumAdts: Set.empty,
  enumCtors: Set.empty, globalTypes: buildGlobalTypes(modules), globalFunctions: Map.empty, classDeclsFields: Map.empty });

test("manifest interface reuse requires the same known source and compiler context", () => {
  const original = fixture(false), saved = manifest(original);
  const cached = (previous, modules, hashes = map(modules.map(mod => [mod.name, hash(mod)])), context = "same-context") =>
    Manifest.buildManifestWithCache(new Just(previous))("")(context)(Map.empty)(modules)(hashes)();
  const typedEdit = [...original.slice(0, 3), moduleOf("Unrelated", [binding("answer", string, literal("new", string))])];
  for (const modules of [original, fixture(true), typedEdit]) {
    assert.deepEqual(cached(saved, modules), manifest(modules));
  }
  const hashes = map(original.map(mod => [mod.name, hash(mod)]));
  const unread = original.map(mod => ({ ...mod, get decls() { throw new Error("unchanged interface re-encoded"); } }));
  assert.deepEqual(cached(saved, unread, hashes), saved);
  const poisoned = { ...saved, modules: map(entries(saved.modules).map(([name, entry]) => [name, { ...entry, interface: "stale" }])) };
  for (const previous of [{ ...poisoned, schema: 0 }, { ...poisoned, context: "other" }, { ...poisoned, directives: "other" }]) {
    assert.deepEqual(cached(previous, original), saved);
  }
  const unknown = Manifest.buildManifest("")("same-context")(Map.empty)(original)(Map.empty)();
  assert.deepEqual(cached({ ...poisoned, modules: map(entries(poisoned.modules).map(([name, entry]) => [name, { ...entry, source: "" }])) }, original, Map.empty), unknown);
});

function optimize(modules) {
  const backends = [], reads = [];
  const observer = new Just({ onSkipModule: () => () => () => () => nothing,
    onReads: () => mod => values => () => {
      const owners = new globalThis.Set([...entries(values.implementations), ...entries(values.directives)]
        .map(([name]) => name.value0.value0));
      reads.push([mod.name, [...owners]]);
    } });
  Builder.buildModulesWithObservation(monadEffectEffect)(observer)(() => () => () => () => {})({
    directives: Map.empty, analyzeCustom: () => () => nothing, foreignSemantics: Map.empty,
    traceIdents: Set.empty, rewriteLimit: 10000, onPrepareModule: () => mod => () => mod,
    onSkipModule: () => () => () => nothing, onCodegenModule: () => () => backend => () => () => backends.push(backend),
  })(list(modules))();
  return { backends, reads };
}

test("a caller body edit invalidates changed specializations and their PBO readers", async t => {
  const directory = mkdtempSync(join(tmpdir(), "gopurs-prepared-manifest-"));
  t.after(() => rmSync(directory, { recursive: true, force: true }));
  const original = fixture(false), edited = fixture(true), before = prepare(original), after = prepare(edited);
  assert.deepEqual(entries(Manifest.diffManifest(manifest(original))(manifest(edited))), [["Caller", "source"]]);
  assert.deepEqual(original[0], edited[0], "the provider's TAST is unchanged");
  assert.notDeepEqual(before[0], after[0], "the caller adds a provider specialization");
  assert.deepEqual(before[2], after[2], "the other caller's prepared body stays unchanged");
  const previous = optimize(before), current = optimize(after), metadata = metadataFor(original);
  assert.deepEqual(previous.reads, [["Generic", []], ["Caller", ["Generic"]], ["Downstream", ["Generic"]], ["Unrelated", []]]);
  assert.notEqual(CodeGen.translate(metadata)(previous.backends[0]), CodeGen.translate(metadata)(current.backends[0]), "restoring the provider would emit stale Go");

  async function initialize(label, modules, candidates, allowed, context = "same-context", ffiDigests = Map.empty) {
    const stage = join(directory, label);
    mkdirSync(join(stage, "prior"), { recursive: true });
    writeFileSync(join(stage, "prior/index.json"), JSON.stringify(candidates));
    for (const candidate of candidates) {
      const owners = previous.reads.find(([name]) => name === candidate.name)[1];
      mkdirSync(join(stage, "prior", candidate.name));
      writeFileSync(join(stage, "prior", candidate.name, "reads.json"), owners.join("\n"));
    }
    const session = { root: directory, stage, context, sameTast: false };
    const prepared = { directives: Map.empty, modules: list(modules), metadata, mainModules: [], sources: [] };
    return runAff(Cache.initialize(1)(session)(prepared)(Set.empty)(Ref.new(Map.empty)())(ffiDigests)(ffiDigests)(allowed));
  }
  const first = await initialize("initial", before, [], Set.empty);
  const candidates = entries(first.keys).map(([name, keys]) => ({ name, ...keys, emitterEnvironment: hash(name), metadata: first.metadata, ffiDigest: "no-foreign-v1", parts: [] }));
  const readsDirectory = join(directory, ".gopurs-cache/reads-v1");
  mkdirSync(readsDirectory, { recursive: true });
  for (const [name] of previous.reads) writeFileSync(join(readsDirectory, name + ".json"), "");
  const next = await initialize("edited", after, candidates, set(after.map(mod => mod.name)));
  const beforeKeys = Object.fromEntries(entries(first.keys)), afterKeys = Object.fromEntries(entries(next.keys));
  assert.notEqual(afterKeys.Generic.prepared, beforeKeys.Generic.prepared, "source equality cannot reuse a prepared key");
  assert.equal(afterKeys.Downstream.prepared, beforeKeys.Downstream.prepared);
  assert.deepEqual(setValues(next.headerSafe), ["Unrelated"], "readers of every changed prepared module must be invalidated");
  const changedContext = await initialize("context-changed", after, candidates, set(after.map(mod => mod.name)), "new-context");
  assert.deepEqual(setValues(changedContext.headerSafe), [], "an auxiliary manifest cannot authorize reuse from another committed context");
  const changedFfi = await initialize("ffi-changed", after, candidates, set(after.map(mod => mod.name)), "same-context", map([["Generic", "new-ffi-digest"]]));
  assert.deepEqual(setValues(changedFfi.headerSafe), [], "matching current and sidecar FFI cannot override different committed FFI");
  const changedMetadata = await initialize("metadata-changed", after,
    candidates.map(candidate => ({ ...candidate, metadata: "older-metadata" })), set(after.map(mod => mod.name)));
  assert.deepEqual(setValues(changedMetadata.headerSafe), [], "committed metadata must agree with the current preparation");
  const legacy = await initialize("legacy-proof-missing", after,
    candidates.map(({ metadata, ffiDigest, ...candidate }) => candidate), set(after.map(mod => mod.name)));
  assert.deepEqual(setValues(legacy.headerSafe), [], "older memo schemas fall back without claiming a header proof");
});

test("missing read proofs invalidate their transitive readers, including absent map entries", () => {
  const names = set(["Source", "Middle", "Reader", "Other"]);
  for (const entriesForMiddle of [[], [["Middle", nothing]]]) {
    const reads = map([["Source", new Just([])], ...entriesForMiddle, ["Reader", new Just(["Middle"])], ["Other", new Just([])]]);
    assert.deepEqual(setValues(Cache.headerSafeFromReads(set(["Source"]))(reads)(names)), ["Other"]);
  }
});
