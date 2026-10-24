import assert from "node:assert/strict";
import { createHash } from "node:crypto";
import { copyFileSync, mkdirSync, mkdtempSync, readFileSync, rmSync, writeFileSync } from "node:fs";
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
import * as Foldable from "../output/Data.Foldable/index.js";
import { ordString } from "../output/Data.Ord/index.js";
import { monadEffectEffect } from "../output/Effect.Class/index.js";
import { Just, Nothing } from "../output/Data.Maybe/index.js";
import { Tuple } from "../output/Data.Tuple/index.js";
import * as C from "../output/PureScript.Backend.Optimizer.CoreFn/index.js";
import * as S from "../output/PureScript.Backend.Optimizer.Semantics/index.js";
import * as B from "../output/PureScript.Backend.Optimizer.Builder/index.js";
import * as Convert from "../output/PureScript.Backend.Optimizer.Convert/index.js";
import * as Cache from "../output/Gopurs.Cache.Modules/index.js";

const hash = bytes => createHash("sha256").update(bytes).digest("hex");
const runAff = value => new Promise((resolve, reject) => Aff.runAff(result => () => result instanceof Either.Left ? reject(result.value0) : resolve(result.value0))(value)());
const nothing = Nothing.value;
const qualified = name => new C.Qualified(new Just("Producer"), name);
const qi = C.ordQualified(C.ordIdent);
const insert = Map.insert(ordString);
const ann = { span: C.emptySpan, type: new Just(C.Int.value), meta: nothing, sourceUsage: nothing };
const moduleOf = (name, expr) => ({ name, path: name + ".purs", span: C.emptySpan, imports: [], exports: ["f"], reExports: [], dataDecls: [], classDecls: [], foreign: Map.empty, comments: [], decls: [new C.NonRec(new C.Binding(ann, "f", expr))] });
const modules = [moduleOf("Producer", new C.ExprLit(ann, new C.LitInt(10))), moduleOf("Consumer", new C.ExprVar(ann, qualified("f")))];
const metadata = withReboxFields({ elidedCtors: Set.empty, ctorTypes: Map.empty, pointerAdtPaths: Map.empty, pointerAdtNodes: Set.empty, pointerAdtLeaves: Map.empty, enumAdts: Set.empty, enumCtors: Set.empty, globalTypes: Map.empty, globalFunctions: Map.empty, classDeclsFields: Map.empty });
const reads = { implementations: Map.insert(qi)(qualified("missing"))(false)(Map.singleton(qualified("f"))(true)), directives: Map.singleton(qualified("f"))(nothing) };

test("optimizer reuse checks observed values and absences instead of unrelated predecessors", async t => {
  const directory = mkdtempSync(join(tmpdir(), "gopurs-optimizer-reads-"));
  const cwd = process.cwd();
  process.chdir(directory);
  t.after(() => { process.chdir(cwd); rmSync(directory, { recursive: true, force: true }); });
  const backends = [];
  B.buildModules(monadEffectEffect)({
    directives: Map.empty, analyzeCustom: () => () => nothing, foreignSemantics: Map.empty, traceIdents: Set.empty, rewriteLimit: 10000,
    onPrepareModule: () => mod => () => mod, onSkipModule: () => () => () => nothing,
    onCodegenModule: () => () => backend => () => () => { backends.push(backend); },
  })(List.fromFoldable(Foldable.foldableArray)(modules))();
  const implementation = Map.lookup(qi)(qualified("f"))(backends[0].implementations).value0;
  const key = name => ({ prepared: hash(name), callerDemands: hash("demands/" + name), optimizerEnvironment: hash("prefix/" + name) });
  const keys = insert("Consumer")(key("Consumer"))(Map.singleton("Producer")(key("Producer")));
  const state = () => ({ sameContext: true, reads: Ref.new(Map.empty)(), restoredPBO: Ref.new(Set.empty)(), priorImplementations: Ref.new(Map.empty)(), priorHeaderDirectives: Map.empty, priorPartTexts: Map.empty, lazyContributions: false, reused: Ref.new(Map.empty)(), priorWorkers: Ref.new(Map.empty)(), currentDigests: Ref.new(Map.empty)(), priorDigests: Ref.new(Map.empty)(), invalidations: Ref.new(Map.empty)(), restored: Ref.new(0)(), backends: Ref.new(Map.empty)(), ffiDigests: Map.empty, priorFfiDigests: Map.empty, headerSafe: Set.empty });
  const makeCache = stage => ({ ...state(), session: { stage, output: join(stage, "output"), context: hash("compiler/options") }, metadata: hash("metadata"), metadataDemands: nothing, keys, candidates: Map.empty });
  const fresh = makeCache(join(directory, "fresh"));
  for (const path of ["modules", "output/purescript"]) mkdirSync(join(fresh.session.stage, path), { recursive: true });
  Ref.write(Map.singleton("Consumer")(reads))(fresh.reads)();
  const memos = [];
  for (let i = 0; i < modules.length; i++) {
    const name = modules[i].name;
    await runAff(Cache.emit(fresh)(Map.empty)(metadata)("unused")(false)(modules[i])(backends[i]));
    const path = join(fresh.session.stage, "modules", name);
    const memo = JSON.parse(readFileSync(join(path, "memo.json")));
    memo.parts = ["pbo/backend-v2", "pbo/implementations-v1", "pbo/directives-v1", "gopurs/workers-v2"].map((codec, j) => {
      const bytes = readFileSync(join(path, j + ".json"));
      return { codec, object: { sha256: hash(bytes), bytes: bytes.length } };
    });
    memos.push(memo);
  }
  async function lookup(label, settings = {}) {
    let cache = makeCache(join(directory, label));
    cache.sameContext = settings.sameContext ?? true;
    if (settings.restored) cache.restoredPBO = Ref.new(Set.singleton("Producer"))();
    cache.headerSafe = settings.headerSafe ? Set.singleton("Consumer") : Set.empty;
    if (settings.memoryHeader) cache.backends = Ref.new(Map.singleton("Consumer")({ ...backends[1], bindings: [] }))();
    cache.keys = insert("Consumer")({ ...key("Consumer"), optimizerEnvironment: hash("changed-prefix"), ...settings.keys })(keys);
    for (const memo of memos) {
      const path = join(cache.session.stage, "prior", memo.name);
      mkdirSync(path, { recursive: true });
      for (let i = 0; i < 4; i++) copyFileSync(join(fresh.session.stage, "modules", memo.name, i + ".json"), join(path, i + ".json"));
      cache.candidates = insert(memo.name)({ ...memo, metadata: new Just(memo.metadata), ffiDigest: new Just(memo.ffiDigest) })(cache.candidates);
    }
    if (settings.directHeaders) {
      cache.lazyContributions = true;
      cache.headerSafe = Set.fromFoldable(Foldable.foldableArray)(ordString)(modules.map(mod => mod.name));
      cache = await runAff(Cache.loadHeaderTexts(cache));
      for (const memo of memos) rmSync(join(cache.session.stage, "prior", memo.name, "2.json"));
    }
    const view = { implementation: owner => ident => {
      if (owner === "Producer" && ident === "f") {
        if (settings.forbidPresentLookup) assert.fail("an identical finalized contribution must not decode its implementation");
        return settings.implementation ?? new Convert.ExternFound(implementation);
      }
      return settings.absent ?? Convert.ExternMissing.value;
    }, availableOwner: owner => owner === "Producer" && (settings.available ?? false),
      directive: () => settings.directive ?? nothing, pending: () => settings.pending ?? false };
    if (settings.pendingFirst) {
      const pending = { ...view, implementation: () => () => Convert.ExternPending.value, pending: () => true };
      assert.equal(await runAff(Cache.lookupBackendWithReads(cache)(modules[1])(pending)), nothing);
    }
    const result = await runAff(Cache.lookupBackendWithReads(cache)(modules[1])(view));
    if (settings.directHeaders && result instanceof Just) {
      await runAff(Cache.emit(cache)(Map.empty)(metadata)("unused")(true)(modules[1])(result.value0));
      assert.ok(Map.lookup(ordString)("Consumer")(Ref.read(cache.reused)()) instanceof Just, "header bytes and the worker fallback must support restoration");
    }
    if (settings.pendingFirst && result instanceof Just) {
      assert.equal(Map.lookup(ordString)("Consumer")(Ref.read(cache.invalidations)()), nothing, "a finalized reuse retained its speculative invalidation");
    }
    return result;
  }
  assert.ok(await lookup("unrelated") instanceof Just, "an unrelated predecessor invalidated optimization");
  assert.ok(await lookup("exact-key-body-fallback", { keys: key("Consumer") }) instanceof Just, "an exact key without a header proof must still load its verified body");
  assert.deepEqual((await lookup("resident-header-body-fallback", { keys: key("Consumer"), memoryHeader: true })).value0.bindings, backends[1].bindings, "a prior resident header cannot replace the complete body without a current header proof");
  assert.ok(await lookup("deferred-then-reused", { pendingFirst: true }) instanceof Just);
  assert.ok(await lookup("restored-visible", { restored: true, available: true, forbidPresentLookup: true }) instanceof Just);
  assert.equal(await lookup("restored-future", { restored: true, implementation: Convert.ExternMissing.value }), nothing, "a restored future owner remains invisible");
  assert.equal(await lookup("restored-pending", { restored: true, implementation: Convert.ExternPending.value, pending: true }), nothing, "a restored but unfinalized owner must still defer");
  assert.equal(await lookup("restored-absence-appeared", { restored: true, available: true, absent: new Convert.ExternFound(implementation) }), nothing, "availability cannot prove a former absence");
  assert.equal(await lookup("missing", { implementation: Convert.ExternMissing.value }), nothing);
  assert.equal(await lookup("pending", { implementation: Convert.ExternPending.value, pending: true }), nothing);
  assert.equal(await lookup("appeared", { absent: new Convert.ExternFound(implementation) }), nothing, "a formerly absent binding appeared");
  const changed = new Tuple({ ...implementation.value0, size: implementation.value0.size + 1 }, implementation.value1);
  assert.equal(await lookup("available-changed", { available: true, implementation: new Convert.ExternFound(changed) }), nothing, "visibility alone does not prove an unchanged contribution");
  assert.equal(await lookup("analysis", { implementation: new Convert.ExternFound(changed) }), nothing, "analysis changes belong to the implementation proof");
  assert.equal(await lookup("directive", { directive: new Just(Map.singleton(S.InlineRef.value)(S.InlineNever.value)) }), nothing);
  assert.equal(await lookup("directive-pending", { pending: true }), nothing);
  assert.equal(await lookup("context", { sameContext: false }), nothing);
  assert.equal(await lookup("source", { keys: { prepared: hash("changed-body") } }), nothing);
  assert.equal(await lookup("callers", { keys: { callerDemands: hash("changed-demands") } }), nothing);
  assert.equal(await lookup("unsafe-header-source", { headerSafe: true, keys: { prepared: hash("changed-body") } }), nothing);
  assert.equal(await lookup("unsafe-header-callers", { headerSafe: true, keys: { callerDemands: hash("changed-demands") } }), nothing);

  // Exercise the real Driver.Cache protocol and Wire decoders. A helper
  // fixture supplies directive bytes; workers retain the individual fallback.
  const snapshot = process.env.GOPURS_COMPILER_SNAPSHOT;
  t.after(() => { if (snapshot === undefined) delete process.env.GOPURS_COMPILER_SNAPSHOT; else process.env.GOPURS_COMPILER_SNAPSHOT = snapshot; });
  process.env.GOPURS_COMPILER_SNAPSHOT = directory;
  mkdirSync(join(directory, "tools"));
  writeFileSync(join(directory, "tools/build-cache-helper"), `#!/usr/bin/env node
const fs = require("node:fs"), path = require("node:path"), assert = require("node:assert/strict");
const request = JSON.parse(fs.readFileSync(0, "utf8")), mode = path.basename(request.stage);
if (request.operation === "materialize") {
  assert.deepEqual(request.parts, ["3"]); assert.deepEqual(request.owners, ["Consumer"]);
  if (mode === "bulk-workers") fs.writeFileSync(path.join(request.stage, "prior/Consumer/3.json"), "invalid");
  console.log(JSON.stringify({ status: "ready", reason: "", stage: request.stage, root: request.stage, input: request.stage, inputLayout: "", output: request.stage, count: 2, ffi: [], context: "same", sameTast: false, packageRequest: "" }));
  process.exit(0);
}
assert.equal(request.operation, "read-parts");
assert.deepEqual(request.parts, ["2"]);
assert.deepEqual(request.owners, ["Consumer", "Producer"]);
const contributions = request.owners.flatMap(name => request.parts.map(part => ({ name, part, text: fs.readFileSync(path.join(request.stage, "prior", name, part + ".json"), "utf8") })));
if (mode === "bulk-missing") contributions.pop();
if (mode === "bulk-duplicate") contributions[1] = contributions[0];
if (mode === "bulk-owner") contributions[0].name = "Outside";
if (mode === "bulk-part") contributions[0].part = "1";
if (mode === "bulk-directives") contributions[0].text = "invalid";
console.log(JSON.stringify(mode === "bulk-failure" ? { status: "error", reason: "broken private object" } : { status: "ready", reason: "", contributions }));
`, { mode: 0o755 });
  assert.ok(await lookup("bulk-ok", { directHeaders: true }) instanceof Just);
  for (const label of ["bulk-missing", "bulk-duplicate", "bulk-owner", "bulk-part", "bulk-failure"]) {
    await assert.rejects(lookup(label, { directHeaders: true }), /incremental cache:/, label);
  }
  assert.equal(await lookup("bulk-directives", { directHeaders: true }), nothing, "malformed directive bytes are not a header hit");
  await assert.rejects(lookup("bulk-workers", { directHeaders: true }), /incremental cache: unreadable workers/, "worker bytes still pass their business decoder");
});
