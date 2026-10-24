import assert from "node:assert/strict";
import { test } from "node:test";
import * as Capsule from "../output/Gopurs.Cache.BindingCapsule/index.js";
import * as Wire from "../output/Gopurs.Cache.Wire/index.js";
import * as Map from "../output/Data.Map/index.js";
import * as Set from "../output/Data.Set/index.js";
import * as List from "../output/Data.List/index.js";
import * as F from "../output/Data.Foldable/index.js";
import * as U from "../output/Data.Unfoldable/index.js";
import * as Ord from "../output/Data.Ord/index.js";
import { Just, Nothing } from "../output/Data.Maybe/index.js";
import { Tuple } from "../output/Data.Tuple/index.js";
import { monadEffectEffect } from "../output/Effect.Class/index.js";
import * as C from "../output/PureScript.Backend.Optimizer.CoreFn/index.js";
import * as B from "../output/PureScript.Backend.Optimizer.Builder/index.js";

const nothing = Nothing.value;
const table = entries => Map.fromFoldable(Ord.ordString)(F.foldableArray)(entries.map(([k, v]) => new Tuple(k, v)));
const qualified = (owner, ident) => new C.Qualified(new Just(owner), ident);
const qmap = entries => Map.fromFoldable(C.ordQualified(C.ordIdent))(F.foldableArray)(entries.map(([owner, ident, value]) => new Tuple(qualified(owner, ident), value)));
const keys = value => Map.toUnfoldable(U.unfoldableArray)(value).map(row => [row.value0.value0.value0, row.value0.value1]);
const set = names => Set.fromFoldable(F.foldableArray)(Ord.ordString)(names);
const implCodec = Wire.codecMap(C.ordQualified(C.ordIdent))(Wire.keyOrderQualified(Wire.keyOrderIdent))(Wire.codecQualified(Wire.codecIdent))(Wire.codecTuple(Wire.codecBackendAnalysis)(Wire.codecExternImpl));
const ann = { span: C.emptySpan, type: new Just(C.Int.value), meta: nothing, sourceUsage: nothing };
function implementations() {
  const result = {};
  const module = name => ({ name, path: name + ".purs", span: C.emptySpan, imports: [], exports: ["value"], reExports: [], dataDecls: [], classDecls: [], foreign: Map.empty, comments: [],
    decls: [new C.NonRec(new C.Binding(ann, "value", new C.ExprLit(ann, new C.LitInt(17))))] });
  B.buildModules(monadEffectEffect)({ directives: Map.empty, analyzeCustom: () => () => nothing, foreignSemantics: Map.empty, traceIdents: Set.empty, rewriteLimit: 10000,
    onPrepareModule: () => value => () => value, onSkipModule: () => () => () => nothing,
    onCodegenModule: () => mod => backend => () => () => { result[mod.name] = backend.implementations; },
  })(List.fromFoldable(F.foldableArray)(["Owner", "Future", "Fresh"].map(module)))();
  return result;
}

test("binding capsules distinguish unknown queries, verified absence and false future observations", () => {
  const maps = implementations(), anchors = table(Object.entries(maps).map(([name, value]) => [name, Wire.hash(Wire.write(implCodec)(value)())]));
  const complete = table(Object.entries(maps).map(([name, value]) => [name, new Just(value)]));
  const reads = { implementations: qmap([["Owner", "value", true], ["Owner", "missing", false], ["Future", "value", false], ["Fresh", "value", true], ["Outside", "missing", false]]), directives: Map.empty };
  const entries = Capsule.select(set(["Owner", "Future"]))(anchors)(complete)(nothing)(reads);
  assert.deepEqual(keys(entries), [["Owner", "missing"], ["Owner", "value"]], "fresh owners and false-but-present future bindings stay outside the capsule");
  const encoded = Capsule.encode("context")("Reader")(anchors)(entries)();
  assert.ok(encoded instanceof Just);
  const decoded = Capsule.decode("context")("Reader")(anchors)(encoded.value0);
  assert.ok(decoded instanceof Just);
  assert.equal(Capsule.lookup(qualified("Owner", "unknown"))(decoded.value0), nothing, "unknown is not absence");
  assert.deepEqual(Capsule.lookup(qualified("Owner", "missing"))(decoded.value0), new Just(nothing));
  assert.ok(Capsule.lookup(qualified("Owner", "value"))(decoded.value0).value0 instanceof Just);
  assert.deepEqual(Capsule.select(set(["Owner"]))(anchors)(Map.empty)(decoded)(reads), entries, "verified coverage can republish without loading complete owners");
  assert.equal(Capsule.encode("context")("Reader")(anchors)(Map.empty)(), nothing, "no empty capsule objects at cold full builds");
});

test("binding capsules reject mismatched identities, anchors, coverage and payload before exposing a value", () => {
  const maps = implementations(), anchors = table(Object.entries(maps).map(([name, value]) => [name, Wire.hash(Wire.write(implCodec)(value)())]));
  const complete = table(Object.entries(maps).map(([name, value]) => [name, new Just(value)]));
  const reads = { implementations: qmap([["Owner", "value", true], ["Owner", "missing", false], ["Future", "value", true]]), directives: Map.empty };
  const entries = Capsule.select(set(["Owner", "Future"]))(anchors)(complete)(nothing)(reads);
  const original = JSON.parse(Capsule.encode("context")("Reader")(anchors)(entries)().value0);
  const mutations = {
    codec: value => { value.codec += "other"; }, context: value => { value.context += "other"; }, reader: value => { value.reader = "Other"; },
    anchorHash: value => { value.anchors[0].sha256 = Wire.hash("wrong"); },
    anchorOwner: value => { value.anchors[0].owner = "Other"; },
    swapAnchors: value => { [value.anchors[0].sha256, value.anchors[1].sha256] = [value.anchors[1].sha256, value.anchors[0].sha256]; },
    duplicateAnchor: value => { value.anchors.push(value.anchors[0]); }, reverseAnchors: value => { value.anchors.reverse(); },
    duplicateCoverage: value => { value.coverage.push(value.coverage[0]); }, reverseCoverage: value => { value.coverage.reverse(); },
    missingCoverage: value => { value.coverage.pop(); }, missingAnchor: value => { value.anchors.pop(); },
    falsePresence: value => { value.coverage.find(row => row.present).present = false; },
    missingAbsentPayload: value => { value.coverage.find(row => !row.present).present = true; },
    payloadHash: value => { value.payload += " "; },
    invalidWire: value => { value.payload = "[]"; value.payloadSha256 = Wire.hash(value.payload); },
    wrongOwner: value => { value.payload = Wire.write(implCodec)(maps.Fresh)(); value.payloadSha256 = Wire.hash(value.payload); },
  };
  for (const [label, mutate] of Object.entries(mutations)) {
    const value = structuredClone(original); mutate(value);
    assert.equal(Capsule.decode("context")("Reader")(anchors)(JSON.stringify(value)), nothing, label);
  }
  assert.equal(Capsule.decode("context")("Reader")(table([["Owner", Wire.hash("changed")]]))(JSON.stringify(original)), nothing);
  assert.equal(Capsule.decode("context")("Reader")(anchors)("{"), nothing);
});
