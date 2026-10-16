import assert from "node:assert/strict";
import { createHash } from "node:crypto";
import { copyFileSync, mkdirSync, mkdtempSync, readFileSync, rmSync } from "node:fs";
import { tmpdir } from "node:os";
import { join } from "node:path";
import { test } from "node:test";
import { withReboxFields } from "./codegen-metadata.mjs";
import * as Aff from "../output/Effect.Aff/index.js";
import * as Either from "../output/Data.Either/index.js";
import * as Map from "../output/Data.Map/index.js";
import { ordString } from "../output/Data.Ord/index.js";
import { Just } from "../output/Data.Maybe/index.js";
import { empty as emptySet } from "../output/Data.Set/index.js";
import { Tuple } from "../output/Data.Tuple/index.js";
import * as Ref from "../output/Effect.Ref/index.js";
import * as Cache from "../output/Gopurs.Cache.Modules/index.js";
import * as CodeGen from "../output/Gopurs.CodeGen/index.js";
import * as Wire from "../output/Gopurs.Cache.Wire/index.js";
import * as G from "../output/Gopurs.GoAst/index.js";
import { SymbolReference, generated } from "../output/Gopurs.GoSymbol/index.js";
import * as C from "../output/PureScript.Backend.Optimizer.CoreFn/index.js";
import { NeutralExpr } from "../output/PureScript.Backend.Optimizer.Semantics/index.js";
import * as S from "../output/PureScript.Backend.Optimizer.Syntax/index.js";

const runAff = aff => new Promise((resolve, reject) => {
  Aff.runAff(result => () => result instanceof Either.Left ? reject(result.value0) : resolve(result.value0))(aff)();
});
const hash = bytes => createHash("sha256").update(bytes).digest("hex");
const expr = syntax => new NeutralExpr(syntax);
const int = C.Int.value, unary = new C.Func([int], int);
const typed = (type, value) => expr(new S.Typed(type, value));
const backend = {
  name: "Consumer", comments: [], imports: emptySet, exports: emptySet, reExports: emptySet,
  dataTypes: Map.empty, dataDecls: [], classDecls: [], foreign: Map.empty, implementations: Map.empty, directives: Map.empty,
  bindings: [{ recursive: false, bindings: [new Tuple("use", typed(unary, expr(new S.Abs(
    [new Tuple(new Just("x"), 0)], typed(int, expr(new S.App(
      typed(unary, expr(new S.Var(new C.Qualified(new Just("Producer"), "f")))),
      [typed(int, expr(new S.Local(new Just("x"), 0)))],
    ))),
  ))))] }],
};
const metadata = withReboxFields({
  elidedCtors: emptySet, ctorTypes: Map.empty, pointerAdtPaths: Map.empty,
  pointerAdtNodes: emptySet, pointerAdtLeaves: Map.empty, enumAdts: emptySet, enumCtors: emptySet,
  globalTypes: Map.empty, globalFunctions: Map.empty, classDeclsFields: Map.empty,
});
const worker = (unit, name) => new SymbolReference(generated(unit)(name));
const info = { reference: worker("Producer", "Call_Producer_f"), fArgs: [G.TypeInt64.value], fRet: G.TypeInt64.value, arity: 1 };
const insert = Map.insert(ordString);
const present = insert("Producer.f")(info)(Map.empty);

test("emission keys cover referenced worker signatures and absences without unrelated workers", async t => {
  const directory = mkdtempSync(join(tmpdir(), "gopurs-emitter-worker-keys-"));
  t.after(() => rmSync(directory, { recursive: true, force: true }));
  const keys = { prepared: hash("prepared"), optimizerEnvironment: hash("optimizer"), callerDemands: hash("callers") };
  const codecs = ["pbo/backend-v2", "pbo/implementations-v1", "pbo/directives-v1", "gopurs/workers-v2"];
  async function emit(label, globalFunctions, previous, metadataKey = "source-metadata", sourceBackend = backend, sourceMetadata = metadata) {
    const stage = join(directory, label), output = join(stage, "output");
    for (const child of ["modules", "prior/Consumer", "output/purescript"]) mkdirSync(join(stage, child), { recursive: true });
    if (previous) {
      for (let i = 0; i < 4; i++) copyFileSync(join(previous.stage, `modules/Consumer/${i}.json`), join(stage, `prior/Consumer/${i}.json`));
      copyFileSync(join(previous.stage, "output/purescript/Consumer.go"), join(stage, "prior/Consumer/module.go"));
    }
    const cache = {
      session: { stage, output, context: hash("compiler-options") },
      candidates: previous ? Map.singleton("Consumer")(previous.memo) : Map.empty,
      keys: Map.singleton("Consumer")(keys), metadata: hash(metadataKey),
      invalidations: Ref.new(Map.empty)(), restored: Ref.new(0)(),
      reads: Ref.new(Map.empty)(),
    };
    // Different all-worker prefixes previously invalidated unrelated consumers.
    const prefix = Wire.hash(JSON.stringify(globalFunctions));
    // Observe actual preparation/emission metadata access without replacing a
    // production function. A proven restore needs only the worker snapshot.
    let preparationReads = 0;
    const observedMetadata = new Proxy({ ...sourceMetadata, globalFunctions }, {
      get(target, key) {
        if (key !== "globalFunctions") preparationReads++;
        return target[key];
      },
    });
    await runAff(Cache.emit(cache)(Map.empty)(observedMetadata)(prefix)(Boolean(previous))({})(sourceBackend));
    const memo = JSON.parse(readFileSync(join(stage, "modules/Consumer/memo.json")));
    memo.parts = codecs.map((codec, i) => {
      const bytes = readFileSync(join(stage, `modules/Consumer/${i}.json`));
      return { codec, object: { sha256: hash(bytes), bytes: bytes.length } };
    });
    return { stage, memo, preparationReads, restored: Ref.read(cache.restored)(), code: readFileSync(join(output, "purescript/Consumer.go"), "utf8") };
  }
  const first = await emit("first", present);
  assert.ok(first.preparationReads > 0, "a fresh module must be prepared");
  assert.match(first.code, /Call_Producer_f\(/);
  const unrelated = insert("Producer.unused")({ ...info, reference: worker("Producer", "Call_Producer_unused") })(present);
  const kept = await emit("unrelated", unrelated, first);
  assert.equal(kept.code, first.code);
  assert.equal(kept.restored, 1, "an unrelated worker invalidated emission");
  assert.equal(kept.preparationReads, 0, "a proven emission hit repeated local preparation");
  const oldWorkers = { ...first, memo: { ...first.memo,
    parts: first.memo.parts.map((part, i) => i === 3 ? { ...part, codec: "gopurs/workers-v1" } : part) } };
  const migrated = await emit("old-worker-codec", present, oldWorkers);
  assert.equal(migrated.restored, 0, "old worker schema was reused despite its codec tag");
  assert.equal(migrated.code, first.code, "worker migration changed monolithic output");
  const metadataChanged = await emit("metadata", present, first, "changed-source-metadata");
  assert.equal(metadataChanged.restored, 0, "changed source metadata reused an old preparation proof");
  assert.ok(metadataChanged.preparationReads > 0);
  const metadataFresh = await emit("metadata-fresh", present, undefined, "changed-source-metadata");
  assert.equal(metadataChanged.code, metadataFresh.code, "metadata: incremental/fresh bytes");
  const ownerChanged = await emit("owner", insert("Producer.f")({ ...info,
    reference: worker("OtherOwner", "Call_Producer_f") })(present), first);
  assert.equal(ownerChanged.restored, 0, "a changed symbol owner reused the previous emission proof");
  assert.equal(ownerChanged.code, first.code, "monolithic names do not depend on placement");
  for (const [label, functions] of [
    ["signature", insert("Producer.f")({ ...info, fArgs: [G.TypeValue.value], fRet: G.TypeValue.value })(present)],
    ["removed", Map.empty],
  ]) {
    const changed = await emit(label, functions, first);
    assert.equal(changed.restored, 0, label);
    assert.notEqual(changed.code, first.code, label);
    const fresh = await emit(label + "-fresh", functions);
    assert.equal(changed.code, fresh.code, label + ": incremental/fresh bytes");
    if (label === "removed") {
      const added = await emit("added", present, changed);
      assert.equal(added.restored, 0, "a previously absent worker appeared");
      assert.equal(added.code, first.code);
    }
  }

  // Borrowing introduces a helper absent from the input graph. This real pass
  // exercises the late proof, including present/missing imported signatures.
  const helper = "Data.Argonaut.Decode.Internal.Record.borrowObject";
  const helperType = new C.Func([C.Any.value, C.Any.value], C.Any.value);
  const global = (owner, name) => expr(new S.Var(new C.Qualified(new Just(owner), name)));
  const app = (fn, ...args) => expr(new S.App(fn, args));
  const identity = app(global("Data.Argonaut.Decode.Class", "decodeForeignObject"), global("Data.Argonaut.Decode.Class", "decodeJsonJson"));
  const method = expr(new S.Accessor(global("Consumer", "objectDecoder"), new S.GetProp("decodeJson")));
  const borrowingBackend = { ...backend, bindings: [{ recursive: false, bindings: [
    new Tuple("objectDecoder", identity),
    new Tuple("use", typed(new C.Func([C.Any.value], int), expr(new S.Abs(
      [new Tuple(new Just("json"), 0)], typed(int, expr(new S.Let(new Just("decoded"), 1,
        app(method, expr(new S.Local(new Just("json"), 0))), expr(new S.Lit(new C.LitInt(1))),
      ))),
    )))),
  ] }] };
  const borrowingMetadata = { ...metadata, globalTypes: insert(helper)(helperType)(Map.empty) };
  const helperInfo = { reference: worker("Data_Argonaut_Decode_Internal_Record", "Call_Data_Argonaut_Decode_Internal_Record_borrowObject"), fArgs: [G.TypeValue.value, G.TypeValue.value], fRet: G.TypeValue.value, arity: 2 };
  const helperPresent = insert(helper)(helperInfo)(Map.empty);
  const prepared = CodeGen.prepareModule({ ...borrowingMetadata, globalFunctions: helperPresent })(borrowingBackend);
  assert.equal(Map.member(ordString)(helper)(CodeGen.workerInputs(helperPresent)(borrowingBackend)), false);
  assert.equal(Map.member(ordString)(helper)(prepared.workers), true, "fixture must introduce the external helper");
  const borrowEmit = (label, functions, previous) => emit("borrow-" + label, functions, previous, "borrowing-metadata", borrowingBackend, borrowingMetadata);
  const borrowed = await borrowEmit("first", helperPresent);
  assert.match(borrowed.code, /Call_Data_Argonaut_Decode_Internal_Record_borrowObject\(/);
  const borrowedHit = await borrowEmit("hit", helperPresent, borrowed);
  assert.equal(borrowedHit.restored, 1);
  assert.ok(borrowedHit.preparationReads > 0, "an introduced reference must use the post-preparation proof");
  assert.equal(borrowedHit.code, borrowed.code);
  for (const [label, functions] of [
    ["changed", insert(helper)({ ...helperInfo, reference: worker("Data_Argonaut_Decode_Internal_Record", "Call_BorrowObject_changed") })(Map.empty)],
    ["removed", Map.empty],
  ]) {
    const changed = await borrowEmit(label, functions, borrowed);
    assert.equal(changed.restored, 0, "introduced helper: " + label);
    assert.notEqual(changed.code, borrowed.code);
    const fresh = await borrowEmit(label + "-fresh", functions);
    assert.equal(changed.code, fresh.code, "introduced helper incremental/fresh: " + label);
    if (label === "removed") {
      const added = await borrowEmit("added", helperPresent, changed);
      assert.equal(added.restored, 0, "a previously absent introduced helper appeared");
      assert.equal(added.code, borrowed.code);
    }
  }
});
