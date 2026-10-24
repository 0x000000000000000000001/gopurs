import assert from "node:assert/strict";
import { mkdirSync, mkdtempSync, readFileSync, rmSync, writeFileSync, existsSync } from "node:fs";
import { tmpdir } from "node:os";
import { join } from "node:path";
import { test } from "node:test";
import { withReboxFields } from "./codegen-metadata.mjs";
import * as Aff from "../output/Effect.Aff/index.js";
import * as Either from "../output/Data.Either/index.js";
import * as Ref from "../output/Effect.Ref/index.js";
import * as Map from "../output/Data.Map/index.js";
import * as Set from "../output/Data.Set/index.js";
import * as F from "../output/Data.Foldable/index.js";
import * as U from "../output/Data.Unfoldable/index.js";
import { ordString } from "../output/Data.Ord/index.js";
import { Just, Nothing } from "../output/Data.Maybe/index.js";
import { Tuple } from "../output/Data.Tuple/index.js";
import * as C from "../output/PureScript.Backend.Optimizer.CoreFn/index.js";
import { NeutralExpr } from "../output/PureScript.Backend.Optimizer.Semantics/index.js";
import * as S from "../output/PureScript.Backend.Optimizer.Syntax/index.js";
import * as Cache from "../output/Gopurs.Cache.Modules/index.js";

const run = aff => new Promise((resolve, reject) => Aff.runAff(result => () => result instanceof Either.Left ? reject(result.value0) : resolve(result.value0))(aff)());
const table = entries => Map.fromFoldable(ordString)(F.foldableArray)(entries.map(([name, value]) => new Tuple(name, value)));
const names = map => Set.toUnfoldable(U.unfoldableArray)(Map.keys(map));
const expr = syntax => new NeutralExpr(syntax);
const backend = owners => ({ name: "Reader", bindings: [{ recursive: false, bindings: owners.map(owner => new Tuple(owner, expr(new S.Var(new C.Qualified(new Just(owner), "f"))))) }] });

test("worker warming decodes only visible cached owners and commits complete batches", async t => {
  const directory = mkdtempSync(join(tmpdir(), "gopurs-worker-warm-"));
  const previous = process.env.GOPURS_COMPILER_SNAPSHOT;
  process.env.GOPURS_COMPILER_SNAPSHOT = directory;
  t.after(() => { if (previous === undefined) delete process.env.GOPURS_COMPILER_SNAPSHOT; else process.env.GOPURS_COMPILER_SNAPSHOT = previous; rmSync(directory, { recursive: true, force: true }); });
  const metadata = withReboxFields({ elidedCtors: Set.empty, ctorTypes: Map.empty, pointerAdtPaths: Map.empty, pointerAdtNodes: Set.empty, pointerAdtLeaves: Map.empty,
    enumAdts: Set.empty, enumCtors: Set.empty, globalTypes: Map.empty, globalFunctions: Map.empty, classDeclsFields: Map.empty });
  const payloads = {}, expected = {};
  const keys = { prepared: "prepared", optimizerEnvironment: "optimizer", callerDemands: "callers" };
  // Use the real emitter and worker codec, including non-empty signatures.
  for (const owner of ["Alpha", "Beta", "Late"]) {
    const stage = join(directory, "produce-" + owner), output = join(stage, "output");
    for (const child of ["modules", "output/purescript"]) mkdirSync(join(stage, child), { recursive: true });
    const cache = { session: { stage, output, context: "context" }, candidates: Map.empty, keys: Map.singleton(owner)(keys), metadata: "metadata", metadataDemands: Nothing.value,
      lazyContributions: false, reads: Ref.new(Map.empty)(), invalidations: Ref.new(Map.empty)(), restored: Ref.new(0)(), backends: Ref.new(Map.empty)(), ffiDigests: Map.empty, priorFfiDigests: Map.empty };
    const mod = { name: owner, comments: [], imports: Set.empty, exports: Set.empty, reExports: Set.empty, dataTypes: Map.empty, dataDecls: [], classDecls: [], foreign: Map.empty,
      implementations: Map.empty, directives: Map.empty,
      bindings: [{ recursive: false, bindings: [new Tuple("f", expr(new S.Typed(new C.Func([C.Int.value], C.Int.value), expr(new S.Abs([new Tuple(new Just("x"), 0)], expr(new S.Typed(C.Int.value, expr(new S.Local(new Just("x"), 0)))))))))] }] };
    expected[owner] = (await run(Cache.emit(cache)(Map.empty)(metadata)("workers")(false)({})(mod))).functions;
    assert.ok(Map.size(expected[owner]) > 0);
    payloads[owner] = readFileSync(join(stage, "modules", owner, "3.json"), "utf8");
  }
  mkdirSync(join(directory, "tools"));
  writeFileSync(join(directory, "tools/build-cache-helper"), `#!/usr/bin/env node
const fs = require("node:fs"), path = require("node:path"), assert = require("node:assert/strict");
const request = JSON.parse(fs.readFileSync(0, "utf8")), mode = path.basename(request.stage);
fs.appendFileSync(path.join(request.stage, "requests.jsonl"), JSON.stringify({operation:request.operation,owners:request.owners}) + "\\n");
assert.deepEqual(request.parts, ["3"]);
if (request.operation === "materialize") {
  assert.deepEqual(request.owners, ["Late"]);
  console.log(JSON.stringify({status:"ready",reason:"",stage:request.stage,root:request.stage,cacheRoot:request.stage+"/.gopurs-cache",input:request.stage,inputLayout:"",output:request.stage,count:3,ffi:[],context:"same",sameTast:false,packageRequest:""}));process.exit(0);
}
assert.equal(request.operation, "read-parts");
const contributions=request.owners.map(name=>({name,part:"3",text:fs.readFileSync(path.join(request.stage,"prior",name,"3.json"),"utf8")}));
if(mode==="missing")contributions.pop();
if(mode==="duplicate")contributions[1]=contributions[0];
if(mode==="outside")contributions[0].name="Future";
if(mode==="wrong-part")contributions[0].part="2";
if(mode==="invalid")contributions[1].text="{}";
if(mode==="wrong-owner")contributions[1].text=contributions[0].text;
console.log(JSON.stringify(mode==="integrity-error"?{status:"error",reason:"corrupt pinned worker"}:{status:"ready",reason:"",contributions}));
`, { mode: 0o755 });
  function fixture(mode) {
    const stage = join(directory, mode); mkdirSync(stage);
    for (const owner of ["Alpha", "Beta", "Late"]) { mkdirSync(join(stage, "prior", owner), { recursive: true }); writeFileSync(join(stage, "prior", owner, "3.json"), payloads[owner]); }
    const parts = [{}, {}, {}, { codec: mode === "old-codec" ? "gopurs/workers-v1" : "gopurs/workers-v2" }];
    return { session: { stage }, lazyContributions: true, priorWorkers: Ref.new(Map.empty)(), priorHeaderDirectives: Map.empty, priorPartTexts: Map.empty, candidates: table(["Alpha", "Beta", "Late"].map(owner => [owner, { parts }])) };
  }
  const requests = cache => existsSync(join(cache.session.stage, "requests.jsonl")) ? readFileSync(join(cache.session.stage, "requests.jsonl"), "utf8").trim().split("\n").map(JSON.parse) : [];
  const snapshot = table([["Alpha", Nothing.value], ["Beta", Nothing.value], ["Current", new Just(Map.empty)], ["Unused", Nothing.value]]);
  const reader = backend(["Alpha", "Beta", "Alpha", "Current", "Reader", "Future"]);
  for (const jobs of [1, 8]) {
    const cache = fixture("valid-" + jobs);
    await run(Cache.warmWorkerSnapshot(jobs)(cache)(snapshot)([reader, reader]));
    assert.deepEqual(names(Ref.read(cache.priorWorkers)()), ["Alpha", "Beta"]);
    assert.deepEqual(requests(cache), [{ operation: "read-parts", owners: ["Alpha", "Beta"] }]);
    for (const owner of ["Alpha", "Beta"]) rmSync(join(cache.session.stage, "prior", owner, "3.json"));
    await run(Cache.warmWorkerSnapshot(jobs)(cache)(snapshot)([reader]));
    const resolve = Cache.resolveWorkerSnapshot(Cache.readWorkerContribution(cache))(snapshot);
    assert.deepEqual(names(await run(resolve(reader))), ["Alpha.f", "Beta.f"]);
    assert.equal(requests(cache).length, 1, "shared owners decode once, with no future/current/unused load");
    assert.deepEqual(await run(Cache.readWorkerContribution(cache)("Late")), expected.Late, "unpredicted owners retain the individual loader");
    assert.deepEqual(requests(cache)[1], { operation: "materialize", owners: ["Late"] });
    const fresh = table([["Alpha", new Just(Map.empty)]]);
    assert.deepEqual(names(await run(Cache.resolveWorkerSnapshot(Cache.readWorkerContribution(cache))(fresh)(backend(["Alpha"])))), [], "fresh signatures override private cached values");
  }
  for (const mode of ["missing", "duplicate", "outside", "wrong-part", "invalid", "wrong-owner", "integrity-error", "old-codec"]) {
    const cache = fixture(mode);
    await assert.rejects(run(Cache.warmWorkerSnapshot(8)(cache)(snapshot)([reader])), /incremental cache:/, mode);
    assert.deepEqual(names(Ref.read(cache.priorWorkers)()), [], mode + ": no partial memo publication");
  }
  for (const mode of ["empty-batch", "eager"]) {
    const cache = fixture(mode); if (mode === "eager") cache.lazyContributions = false;
    await run(Cache.warmWorkerSnapshot(8)(cache)(snapshot)(mode === "empty-batch" ? [] : [reader]));
    assert.deepEqual(requests(cache), []);
  }
});
