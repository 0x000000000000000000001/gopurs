import assert from "node:assert/strict";
import { createHash } from "node:crypto";
import { copyFileSync, existsSync, readFileSync, readdirSync, writeFileSync } from "node:fs";
import { join, resolve } from "node:path";
import { fileURLToPath } from "node:url";

assert(process.argv.length === 3, "Usage: node tools/preparation-publish.mjs CAMPAIGN");
const campaign = resolve(process.argv[2]);
const repository = fileURLToPath(new URL("../", import.meta.url));
const load = path => JSON.parse(readFileSync(join(campaign, path), "utf8"));
const hash = bytes => createHash("sha256").update(bytes).digest("hex");
const measurements = load("measurement/results.json");
assert.equal(measurements.length, 36);
assert(measurements.every(row => row.status === 0 && row.byteExact));
const codecs = load("codec-measurement/summary.json");
const validation = load("validation/summary.json");
const costs = load("analysis/costs.json");
assert(codecs.completed && validation.completed);
assert.equal(validation.comparisons.length, 7);
assert(validation.comparisons.every(row => row.sameTast && row.byteExact && row.hit && row.forced));
assert.equal(validation.executionOracles, 4);
assert.equal(costs.acceptance, false, "P00 accounting is not optimization acceptance");

const fixtures = [];
for (const run of readdirSync(campaign).filter(name => /^run-\d+$/.test(name))) {
  const directory = join(campaign, run, "workspaces");
  if (!existsSync(directory)) continue;
  for (const workspace of readdirSync(directory)) {
    const path = join(run, "workspaces", workspace, "summary.json");
    if (!existsSync(join(campaign, path))) continue;
    const summary = load(path);
    if (summary.completed && summary.states === 7 && summary.comparisons === 28) fixtures.push({ path, summary });
  }
}
assert.equal(fixtures.length, 1, "Identify the completed preparation fixture campaign");

const traces = {};
for (const directory of ["measurement/logs", "codec-measurement/logs"]) {
  const names = readdirSync(join(campaign, directory)).filter(name => name.endsWith(".resources.jsonl"));
  assert(names.length > 0);
  for (const name of names) {
    const records = readFileSync(join(campaign, directory, name), "utf8").trim().split("\n").map(line => JSON.parse(line));
    const result = records.at(-1);
    assert(result.completed && result.status === 0 && !result.contaminated && result.errors.length === 0, name);
    assert(records.every(row => !row.competing?.length), name);
  }
  traces[directory] = names.length;
}

const outputs = [
  ["analysis/timings.csv", "docs/build-preparation-p00-timings.csv"],
  ["analysis/costs.json", "docs/build-preparation-p00-costs.json"],
];
const manifestPath = join(repository, "docs/build-preparation-p00-evidence.json");
for (const [, target] of outputs) assert(!existsSync(join(repository, target)), "Preserve prior published evidence: " + target);
assert(!existsSync(manifestPath));
const retained = ["provenance.json", "source-inventory.json", "frontend-base.json", "measurement/results.json",
  "measurement/business-oracles.json", "analysis/timings.json", "analysis/timings.csv", "analysis/costs.json",
  "codec/build-logs/summary.json", "codec-measurement/results.json", "codec-measurement/sets.json",
  "codec-measurement/changed-encoding.json", "codec-measurement/publication.json", "codec-measurement/summary.json",
  "validation/summary.json", "validation/commands.json", "validation/runner.json", fixtures[0].path];
const files = Object.fromEntries(retained.map(path => {
  const bytes = readFileSync(join(campaign, path));
  return [path, { bytes: bytes.length, sha256: hash(bytes) }];
}));
const provenance = load("provenance.json");
for (const [source, target] of outputs) copyFileSync(join(campaign, source), join(repository, target));
writeFileSync(manifestPath, JSON.stringify({ schema: 1, campaign, created: new Date().toISOString(),
  scope: "P00 reconstituted baseline and codec accounting; no incremental preparation optimization accepted.",
  repositories: provenance.repositories, toolchain: provenance.tools, files, traces, fixtures,
  codecs, validation, exports: Object.fromEntries(outputs) }, null, 2) + "\n");
console.log("P00: published 36 timings, object-cost accounting and evidence fingerprints");
