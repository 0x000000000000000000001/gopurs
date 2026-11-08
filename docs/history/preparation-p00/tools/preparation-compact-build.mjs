import assert from "node:assert/strict";
import { createHash } from "node:crypto";
import { constants, cpSync, existsSync, mkdirSync, readFileSync, writeFileSync } from "node:fs";
import { join, resolve } from "node:path";
import { fileURLToPath } from "node:url";
import { observe } from "./preparation-observe.mjs";

assert(process.argv.length === 4, "Usage: node tools/preparation-compact-build.mjs P00 NEW_STUDY");
const parent = resolve(process.argv[2]), study = resolve(process.argv[3]);
const root = fileURLToPath(new URL("../", import.meta.url));
assert(!existsSync(study), "Preserve prior study attempts");
assert(JSON.parse(readFileSync(join(parent, "validation/summary.json"))).completed);
mkdirSync(study, { recursive: true });
for (const name of ["tmp", "logs", "harness"]) mkdirSync(join(study, name));
for (const name of ["tree/gopurs", "tree/purescript-backend-optimizer-gopurs", "diagnostic/gopurs", "diagnostic/purescript-backend-optimizer-gopurs", "toolchain", "native-base-manifest.json"]) {
  cpSync(join(parent, name), join(study, name), { recursive: true, dereference: true, mode: constants.COPYFILE_FICLONE });
}
const hash = path => createHash("sha256").update(readFileSync(path)).digest("hex");
const inputs = ["provenance.json", "measurement/states.json", "measurement/business-oracles.json", "codec-measurement/sets.json", "codec-measurement/summary.json", "validation/summary.json"];
writeFileSync(join(study, "provenance.json"), JSON.stringify({ schema: 1, parent, created: new Date().toISOString(),
  inputs: Object.fromEntries(inputs.map(name => [name, hash(join(parent, name))])), candidate: { buckets: 64, gzip: "BestSpeed", proof: "whole-preparation-inputs" } }, null, 2) + "\n");
for (const name of ["preparation-compact-build.mjs", "preparation-probe-build.mjs", "preparation-probe/Probe.purs", "preparation-probe/CompactProbe.purs"]) {
  const target = join(study, "harness", name.replaceAll("/", "-"));
  cpSync(join(root, "tools", name), target);
}
const result = observe(process.execPath, [join(root, "tools/preparation-probe-build.mjs"), study, "compact"],
  { cwd: root, env: process.env, encoding: "utf8", timeout: 1800000, maxBuffer: 64 * 1024 * 1024 }, join(study, "logs/build.resources.jsonl"));
writeFileSync(join(study, "logs/build.stdout"), result.stdout ?? "");
writeFileSync(join(study, "logs/build.stderr"), result.stderr ?? "");
assert.ifError(result.error);
assert.equal(result.status, 0, "See retained study build logs");
console.log(result.stdout);
