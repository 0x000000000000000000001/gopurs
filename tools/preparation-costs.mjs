import assert from "node:assert/strict";
import { existsSync, readFileSync, writeFileSync } from "node:fs";
import { join, resolve } from "node:path";

assert(process.argv.length === 3, "Usage: node tools/preparation-costs.mjs CAMPAIGN");
const campaign = resolve(process.argv[2]);
const load = path => JSON.parse(readFileSync(join(campaign, path)));
const timings = load("analysis/timings.json");
assert(load("codec-measurement/summary.json").completed);
const records = load("codec-measurement/results.json");
const publication = load("codec-measurement/publication.json");
const changed = load("codec-measurement/changed-encoding.json");
const sets = load("codec-measurement/sets.json");
const median = values => {
  const ordered = [...values].sort((a, b) => a - b), middle = Math.floor(ordered.length / 2);
  assert(ordered.length);
  return ordered.length % 2 ? ordered[middle] : (ordered[middle - 1] + ordered[middle]) / 2;
};
const distribution = values => ({ median: median(values), min: Math.min(...values), max: Math.max(...values), samples: values });
const metric = (row, label) => {
  const values = row.metrics.filter(item => item.phase === label);
  assert.equal(values.length, 1, row.label + ": " + label);
  return values[0].seconds;
};
const categories = {};
for (const kind of Object.keys(sets.base)) {
  const action = ["direct", "callers", "transitive", "contributions"].includes(kind) ? "encode-demands" : "encode-" + kind;
  const encodes = records.filter(row => row.state === "base" && row.action === action);
  const restores = records.filter(row => row.state === "base" && row.action === "restore-" + kind);
  assert.equal(encodes.length, 3); assert.equal(restores.length, 3);
  const unique = new Map(sets.base[kind].entries.map(entry => [entry.object.sha256, entry.object.bytes]));
  const edits = {};
  for (const state of ["leaf", "shared", "ffi"]) {
    const encode = changed.filter(row => row.state === state && row.kind === kind);
    const stage = publication.filter(row => row.state === state && row.kind === kind);
    assert.equal(encode.length, 3); assert.equal(stage.length, 3);
    edits[state] = { changed: sets[state][kind].changed.length, removed: sets[state][kind].removed.length,
      changedBytes: sets[state][kind].changedBytes,
      encodeWrite: distribution(encode.map(row => metric(row, "probe/reencode-write/" + kind))),
      stage: distribution(stage.map(row => row.seconds)), stageProcess: distribution(stage.map(row => row.process)) };
  }
  categories[kind] = {
    count: sets.base[kind].count, rawBytes: sets.base[kind].bytes, uniqueBytes: [...unique.values()].reduce((sum, value) => sum + value, 0),
    encodeWrite: distribution(encodes.map(row => metric(row, "probe/encode-write/" + kind))),
    restore: distribution(restores.map(row => metric(row, "probe/read-validate-decode/" + kind))),
    canonicalRoundtrip: distribution(restores.map(row => metric(row, "probe/verify-roundtrip/" + kind))),
    stageCold: distribution(publication.filter(row => row.state === "base" && row.kind === kind).map(row => row.seconds)),
    stageColdProcess: distribution(publication.filter(row => row.state === "base" && row.kind === kind).map(row => row.process)),
    edits,
  };
}

// These are accounting bounds, not predictions of an implemented incremental
// collector. The remaining budget explicitly belongs to dependency proofs and
// affected-region recomputation, neither of which this probe implements.
const models = {};
for (const [name, kinds, gross] of [
  ["source-only", ["source"], row => row.phases["load TAST + sort"] ?? 0],
  ["prepared-and-metadata", ["prepared", "metadata"], row => row.avoidableCeiling],
  ["prepared-with-contributions", ["prepared", "metadata", "callers", "contributions"], row => row.avoidableCeiling],
  ["caller-and-transitive-contributions", ["callers", "contributions"], row =>
    (row.phases["prepare/direct demands"] ?? 0) + (row.phases["transitive specializations"] ?? 0)],
  ["final-instantiations", ["transitive"], row =>
    (row.phases["prepare/direct demands"] ?? 0) + (row.phases["transitive specializations"] ?? 0)],
]) {
  const unique = new Map(kinds.flatMap(kind => sets.base[kind].entries.map(entry => [entry.object.sha256, entry.object.bytes])));
  const edits = {};
  for (const scenario of ["leaf", "shared", "ffi"]) {
    const observer = name === "source-only" || kinds.includes("prepared") ? "reference" : "diagnostic";
    const grossSamples = timings.samples.filter(row => row.scenario === scenario && row.compiler === observer);
    const grossValues = grossSamples.map(gross);
    const restoration = kinds.reduce((total, kind) => total + categories[kind].restore.median, 0);
    const serialization = kinds.reduce((total, kind) => total + categories[kind].edits[scenario].encodeWrite.median, 0);
    const staging = kinds.reduce((total, kind) => total + categories[kind].edits[scenario].stageProcess.median, 0);
    const knownCosts = restoration + serialization + staging;
    const accounted = grossSamples.map(row => {
      const baseline = timings.samples.find(sample => sample.scenario === scenario && sample.compiler === "reference" && sample.pair === row.pair);
      assert(baseline, "Matching reference sample");
      const headroom = gross(row) - knownCosts;
      return { pair: row.pair, process: baseline.process, gross: gross(row), headroom,
        percent: 100 * headroom / baseline.process, residual: baseline.process - headroom,
        tenPercentBudget: headroom - baseline.process * 0.10 };
    });
    const headroom = median(accounted.map(row => row.headroom));
    const residual = median(accounted.map(row => row.residual));
    edits[scenario] = { gross: distribution(grossValues), restoration, changedSerialization: serialization,
      marginalStagingProcesses: staging, knownCosts, headroomBeforeProofAndRecomputation: headroom,
      headroomPercent: median(accounted.map(row => row.percent)), residualBeforeProofAndRecomputation: residual,
      remainingBudgetForTenPercent: median(accounted.map(row => row.tenPercentBudget)), remainingBudgetForThirtySeconds: 30 - residual,
      accounted, residualDistribution: distribution(accounted.map(row => row.residual)),
      detailedObserverRepresentative: timings.scenarios[scenario].instrumentationWithinFivePercent,
    };
  }
  models[name] = { kinds, bytes: [...unique.values()].reduce((total, bytes) => total + bytes, 0),
    coldEncodeAndStage: kinds.reduce((total, kind) => total + categories[kind].encodeWrite.median + categories[kind].stageColdProcess.median, 0),
    coldStageOnly: kinds.reduce((total, kind) => total + categories[kind].stageCold.median, 0), edits,
    caveats: ["No new dependency proofs, input index or affected-region computation is implemented.",
      "Separate probe processes do not measure GC interaction with the complete backend.",
      "Per-family staging processes are additive estimates; shared transactions may amortize startup and index costs.",
      "Object-cost medians are additive assumptions, not an integrated performance trial; residuals use individual backend samples and their paired observer.",
      "Final instantiations alone do not supply the provenance needed to retract contributions after source edits."] };
}
const target = join(campaign, "analysis/costs.json");
assert(!existsSync(target), "Preserve earlier cost analyses");
writeFileSync(target, JSON.stringify({ schema: 1, categories, models, acceptance: false,
  scope: "Measured object costs and zero-proof headroom; a decision aid before P01, not optimization acceptance." }, null, 2) + "\n");
for (const [name, value] of Object.entries(categories)) console.log(`${name}: ${(value.uniqueBytes / 1048576).toFixed(2)} MiB, restore ${value.restore.median.toFixed(3)} s, initial encode+stage ${(value.encodeWrite.median + value.stageColdProcess.median).toFixed(3)} s`);
for (const [name, value] of Object.entries(models)) console.log(name + ": " + Object.entries(value.edits).map(([scenario, data]) => `${scenario} headroom ${data.headroomBeforeProofAndRecomputation.toFixed(3)} s`).join("; "));
