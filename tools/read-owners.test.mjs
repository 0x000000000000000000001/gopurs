import assert from "node:assert/strict";
import { mkdirSync, mkdtempSync, rmSync, writeFileSync } from "node:fs";
import { tmpdir } from "node:os";
import { join } from "node:path";
import { test } from "node:test";
import * as Aff from "../output/Effect.Aff/index.js";
import * as Either from "../output/Data.Either/index.js";
import * as Map from "../output/Data.Map/index.js";
import * as Set from "../output/Data.Set/index.js";
import * as F from "../output/Data.Foldable/index.js";
import * as U from "../output/Data.Unfoldable/index.js";
import { ordString } from "../output/Data.Ord/index.js";
import { Just } from "../output/Data.Maybe/index.js";
import * as Cache from "../output/Gopurs.Cache.Modules/index.js";

const runAff = value => new Promise((resolve, reject) => Aff.runAff(result => () => result instanceof Either.Left ? reject(result.value0) : resolve(result.value0))(value)());
const entries = value => Object.fromEntries(Map.toUnfoldable(U.unfoldableArray)(value).map(entry => [entry.value0, entry.value1 instanceof Just ? entry.value1.value0 : null]));

test("grouped read proofs preserve empty, unknown, missing and requested subsets", async t => {
  const directory = mkdtempSync(join(tmpdir(), "gopurs-read-owners-"));
  t.after(() => rmSync(directory, { recursive: true, force: true }));
  for (const [name, text] of [["Known", "Legacy\n"], ["Unknown", ""], ["Absent", ""]]) {
    mkdirSync(join(directory, name));
    writeFileSync(join(directory, name, "reads.json"), text);
  }
  const load = names => runAff(Cache.loadReads(directory)(names));
  assert.deepEqual(entries(await load(["Known", "Unknown", "Never"])), { Known: ["Legacy"], Unknown: [], Never: null }, "an absent grouped index uses the legacy private files");
  writeFileSync(join(directory, "owners.json"), JSON.stringify([
    { name: "Known", text: "Provider\n\nOther\n" }, { name: "Empty", text: "" }, { name: "Unknown", text: null },
    { name: "Reader", text: "Unknown\n" }, { name: "UnusedExtra", text: "UnusedOwner\n" },
  ]));
  const proof = await load(["Known", "Empty", "Unknown", "Reader", "Absent"]);
  assert.deepEqual(entries(proof), { Known: ["Provider", "Other"], Empty: [], Unknown: null, Reader: ["Unknown"], Absent: null }, "the grouped index is authoritative, including unknown or absent entries");
  assert.deepEqual(entries(await load(["Empty", "Known"])), { Empty: [], Known: ["Provider", "Other"] }, "the exact-prefix filter may request a strict subset of the manifest");
  assert.deepEqual(entries(await load([])), {});
  const names = Set.fromFoldable(F.foldableArray)(ordString)(["Empty", "Unknown", "Reader", "Absent"]);
  assert.deepEqual(Set.toUnfoldable(U.unfoldableArray)(Cache.headerSafeFromReads(Set.empty)(proof)(names)), ["Empty"], "unknown proofs invalidate their readers while a known empty proof stays safe");
});

test("malformed grouped proofs never fall back to apparently valid individual files", async t => {
  const directory = mkdtempSync(join(tmpdir(), "gopurs-invalid-owners-"));
  t.after(() => rmSync(directory, { recursive: true, force: true }));
  mkdirSync(join(directory, "Requested"));
  writeFileSync(join(directory, "Requested/reads.json"), "");
  const path = join(directory, "owners.json");
  const valid = { name: "Requested", text: "" };
  for (const [label, value] of [
    ["not JSON", "{"], ["object", "{}"], ["null", "null"],
    ["duplicate requested", JSON.stringify([valid, valid])],
    ["duplicate extra", JSON.stringify([valid, { name: "Extra", text: null }, { name: "Extra", text: "" }])],
    ["missing text", JSON.stringify([{ name: "Requested" }])],
    ["invalid text", JSON.stringify([{ name: "Requested", text: [] }])],
    ["invalid name", JSON.stringify([{ name: 1, text: "" }])],
    ["empty name", JSON.stringify([{ name: "", text: "" }])],
  ]) {
    writeFileSync(path, value);
    await assert.rejects(runAff(Cache.loadReads(directory)(["Requested"])), /incremental cache:/, label);
  }
  rmSync(path);
  mkdirSync(path);
  await assert.rejects(runAff(Cache.loadReads(directory)(["Requested"])), /incremental cache:/, "an existing unreadable index is not a legacy session");
});
