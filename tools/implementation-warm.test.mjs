import assert from "node:assert/strict";
import { mkdirSync, mkdtempSync, readFileSync, rmSync, writeFileSync, existsSync } from "node:fs";
import { tmpdir } from "node:os";
import { join } from "node:path";
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
import { Just, Nothing } from "../output/Data.Maybe/index.js";
import { Tuple } from "../output/Data.Tuple/index.js";
import { monadEffectEffect } from "../output/Effect.Class/index.js";
import * as C from "../output/PureScript.Backend.Optimizer.CoreFn/index.js";
import * as B from "../output/PureScript.Backend.Optimizer.Builder/index.js";
import * as Cache from "../output/Gopurs.Cache.Modules/index.js";
import * as Wire from "../output/Gopurs.Cache.Wire/index.js";
const nothing = Nothing.value;
const run = value => new Promise((resolve, reject) => Aff.runAff(result => () => result instanceof Either.Left ? reject(result.value0) : resolve(result.value0))(value)());
const table = entries => Map.fromFoldable(Ord.ordString)(F.foldableArray)(entries.map(([name, value]) => new Tuple(name, value)));
const set = names => Set.fromFoldable(F.foldableArray)(Ord.ordString)(names);
const names = value => Set.toUnfoldable(U.unfoldableArray)(Map.keys(value));
const codec = Wire.codecMap(C.ordQualified(C.ordIdent))(Wire.keyOrderQualified(Wire.keyOrderIdent))(Wire.codecQualified(Wire.codecIdent))(Wire.codecTuple(Wire.codecBackendAnalysis)(Wire.codecExternImpl));
const ann = { span: C.emptySpan, type: new Just(C.Int.value), meta: nothing, sourceUsage: nothing };
const moduleOf = name => ({ name, path: name + ".purs", span: C.emptySpan, imports: [], exports: ["value"], reExports: [], dataDecls: [], classDecls: [], foreign: Map.empty, comments: [], decls: [new C.NonRec(new C.Binding(ann, "value", new C.ExprLit(ann, new C.LitInt(17))))] });

test("implementation warming remains private, ordered by eligibility and atomic on failure", async t => {
  const directory = mkdtempSync(join(tmpdir(), "gopurs-implementation-warm-"));
  const previousCwd = process.cwd(), previousSnapshot = process.env.GOPURS_COMPILER_SNAPSHOT;
  process.chdir(directory); process.env.GOPURS_COMPILER_SNAPSHOT = directory;
  t.after(() => { process.chdir(previousCwd); if (previousSnapshot === undefined) delete process.env.GOPURS_COMPILER_SNAPSHOT; else process.env.GOPURS_COMPILER_SNAPSHOT = previousSnapshot; rmSync(directory, { recursive: true, force: true }); });
  const implementations = {};
  B.buildModules(monadEffectEffect)({ directives: Map.empty, analyzeCustom: () => () => nothing, foreignSemantics: Map.empty, traceIdents: Set.empty, rewriteLimit: 10000,
    onPrepareModule: () => mod => () => mod, onSkipModule: () => () => () => nothing,
    onCodegenModule: () => mod => backend => () => () => { implementations[mod.name] = backend.implementations; },
  })(List.fromFoldable(F.foldableArray)(["Alpha", "Beta", "Late"].map(moduleOf)))();
  mkdirSync(join(directory, "tools"));
  writeFileSync(join(directory, "tools/build-cache-helper"), `#!/usr/bin/env node
const fs = require("node:fs"), path = require("node:path"), assert = require("node:assert/strict");
const request = JSON.parse(fs.readFileSync(0, "utf8"));
assert.equal(request.operation, "materialize"); assert.deepEqual(request.parts, ["1"]);
fs.appendFileSync(path.join(request.stage, "requests.jsonl"), JSON.stringify(request.owners) + "\\n");
const response = JSON.parse(fs.readFileSync(path.join(request.stage, "session-response.json"), "utf8"));
if (path.basename(request.stage) === "materialize-error") { response.status = "error"; response.reason = "corrupt pinned implementation"; }
console.log(JSON.stringify(response));
`, { mode: 0o755 });
  const ranks = table(["Alpha", "Beta", "Fresh", "Late", "Reader", "Future"].map((name, index) => [name, index]));
  function fixture(mode) {
    const stage = join(directory, mode); mkdirSync(stage);
    const session = { status: "ready", reason: "", stage, root: directory, input: stage, inputLayout: "", output: stage, count: 6, ffi: [], context: "same", sameTast: false, packageRequest: "" };
    writeFileSync(join(stage, "session-response.json"), JSON.stringify(session));
    for (const name of ["Alpha", "Beta", "Late"]) {
      mkdirSync(join(stage, "prior", name), { recursive: true });
      const content = name === "Beta" && mode === "wrong-owner" ? Wire.write(codec)(implementations.Alpha)() : name === "Beta" && mode === "invalid" ? "{}" : Wire.write(codec)(implementations[name])();
      if (!(name === "Beta" && mode === "missing")) writeFileSync(join(stage, "prior", name, "1.json"), content);
    }
    const cache = { session, lazyContributions: true, priorImplementations: Ref.new(Map.empty)(), restoredPBO: Ref.new(Set.empty)(), backends: Ref.new(Map.empty)(),
      headerSafe: set(["Alpha", "Beta", "Late", "Reader", "Future", "Outside"]), schedulingReads: table([["Reader", ["Alpha", "Beta", "Alpha", "Fresh", "Reader", "Future", "Outside"]]]) };
    return { cache, requests: () => existsSync(join(stage, "requests.jsonl")) ? readFileSync(join(stage, "requests.jsonl"), "utf8").trim().split("\n").map(JSON.parse) : [] };
  }
  for (const jobs of [1, 8]) {
    const f = fixture("valid-" + jobs), { cache } = f;
    assert.deepEqual(Cache.implementationOwners(cache)(ranks)("Reader"), ["Alpha", "Beta"]);
    const restored = Ref.read(cache.restoredPBO)(), backends = Ref.read(cache.backends)();
    await run(Cache.warmImplementations(jobs)(cache)(ranks)("Reader"));
    assert.deepEqual(names(Ref.read(cache.priorImplementations)()), ["Alpha", "Beta"]);
    assert.equal(Ref.read(cache.restoredPBO)(), restored); assert.equal(Ref.read(cache.backends)(), backends);
    assert.deepEqual(f.requests(), [["Alpha", "Beta"]]);
    for (const name of ["Alpha", "Beta"]) assert.equal(Wire.write(codec)(Cache.readImplementationsSync(cache)(name)())(), Wire.write(codec)(implementations[name])());
    await run(Cache.warmImplementations(jobs)(cache)(ranks)("Reader"));
    assert.deepEqual(f.requests(), [["Alpha", "Beta"]], "memoized contributions do not materialize twice");
    assert.deepEqual(Cache.readImplementationsSync(cache)("Late")(), implementations.Late, "new reads retain the lazy path");
    assert.deepEqual(f.requests(), [["Alpha", "Beta"], ["Late"]]);
  }
  const cold = fixture("cold"); cold.cache.lazyContributions = false;
  await run(Cache.warmImplementations(8)(cold.cache)(ranks)("Reader"));
  assert.deepEqual(cold.requests(), []); assert.deepEqual(names(Ref.read(cold.cache.priorImplementations)()), []);
  for (const mode of ["invalid", "wrong-owner", "missing", "materialize-error"]) {
    const f = fixture(mode);
    await assert.rejects(run(Cache.warmImplementations(8)(f.cache)(ranks)("Reader")), /incremental cache:/, mode);
    assert.deepEqual(names(Ref.read(f.cache.priorImplementations)()), [], mode + ": failed batches publish no partial memo");
    assert.deepEqual(Set.toUnfoldable(U.unfoldableArray)(Ref.read(f.cache.restoredPBO)()), []);
    assert.deepEqual(names(Ref.read(f.cache.backends)()), []);
  }
});
