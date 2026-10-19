import assert from "node:assert/strict";
import { existsSync, mkdirSync, readFileSync, writeFileSync } from "node:fs";
import { join, resolve } from "node:path";

assert(process.argv.length === 3, "Usage: node tools/preparation-analyze.mjs CAMPAIGN");
const campaign = resolve(process.argv[2]);
const destination = join(campaign, "analysis");
assert(!existsSync(destination), "Preserve previous analyses");
const rows = JSON.parse(readFileSync(join(campaign, "measurement/results.json")));
assert.equal(rows.length, 36);
assert(rows.every(row => row.status === 0 && row.byteExact));
const median = values => {
  const ordered = [...values].sort((a, b) => a - b);
  assert(ordered.length);
  const mid = Math.floor(ordered.length / 2);
  return ordered.length % 2 ? ordered[mid] : (ordered[mid - 1] + ordered[mid]) / 2;
};
const distribution = values => ({ median: median(values), min: Math.min(...values), max: Math.max(...values), samples: values });
const sum = values => values.reduce((total, value) => total + value, 0);
const topLevel = ["cache lookup + freeze", "load TAST + sort", "prepare + monomorphize", "freeze FFI", "runtime",
  "optimize + emit", "Go exports", "Go packages", "entry points", "cache publish"];
const samples = rows.map(row => {
  const phases = {}, count = {};
  for (const { phase, seconds } of row.metrics) {
    phases[phase] = (phases[phase] ?? 0) + seconds;
    count[phase] = (count[phase] ?? 0) + 1;
  }
  const phase = name => phases[name] ?? 0;
  const prepareExclusive = Object.keys(phases).filter(name => name.startsWith("prepare/") && name !== "prepare/transitive jobs");
  const avoidableCeiling = phase("load TAST + sort") + phase("prepare + monomorphize") + phase("module cache keys");
  return {
    scenario: row.scenario, pair: row.pair, compiler: row.compiler, process: row.seconds,
    peakRssBytes: row.peakRssBytes, objectsBytes: row.objectsBytes, generatedGoBytes: row.generatedGoBytes, phases,
    processEnvelope: row.seconds - phase("backend total"),
    backendResidual: phase("backend total") - sum(topLevel.map(phase)),
    optimizeResidual: phase("optimize + emit") - sum(["module cache keys", "module cache all-hit lookup", "PBO producer", "cached module producer", "emitter drain"].map(phase)),
    prepareResidual: row.compiler === "diagnostic" && row.scenario !== "identical"
      ? phase("prepare + monomorphize") - phase("transitive specializations") - sum(prepareExclusive.map(phase)) : null,
    transitiveRounds: count["prepare/transitive jobs"] ?? null,
    transitiveControl: row.compiler === "diagnostic" && row.scenario !== "identical"
      ? phase("transitive specializations") - phase("prepare/transitive jobs") : null,
    avoidableCeiling, zeroCostPreparationResidual: row.seconds - avoidableCeiling,
    note: "Ceiling includes the entire source load and module-key phase; restoration, invalidation and cold recomputation are not free.",
  };
});
const scenarios = {};
for (const scenario of ["identical", "leaf", "shared", "ffi", "cold-app", "cold-complete"]) {
  const modes = {};
  for (const compiler of ["reference", "diagnostic"]) {
    const selected = samples.filter(row => row.scenario === scenario && row.compiler === compiler);
    assert.equal(selected.length, 3);
    assert.deepEqual(selected.map(row => row.pair).sort(), [1, 2, 3]);
    const phases = [...new Set(selected.flatMap(row => Object.keys(row.phases)))].sort();
    const summary = Object.fromEntries(["process", "peakRssBytes", "objectsBytes", "generatedGoBytes", "processEnvelope", "backendResidual", "optimizeResidual", "avoidableCeiling", "zeroCostPreparationResidual"]
      .map(key => [key, distribution(selected.map(row => row[key]))]));
    for (const key of ["prepareResidual", "transitiveRounds", "transitiveControl"]) {
      const values = selected.map(row => row[key]).filter(value => value !== null);
      if (values.length) summary[key] = distribution(values);
    }
    modes[compiler] = { ...summary, phases: Object.fromEntries(phases.map(phase => [phase, distribution(selected.map(row => row.phases[phase] ?? 0))])) };
  }
  const pairDeltas = [1, 2, 3].map(pair => {
    const selected = samples.filter(row => row.scenario === scenario && row.pair === pair);
    const before = selected.find(row => row.compiler === "reference").process;
    const after = selected.find(row => row.compiler === "diagnostic").process;
    return { pair, seconds: after - before, percent: (after / before - 1) * 100 };
  });
  const overheadPercent = (modes.diagnostic.process.median / modes.reference.process.median - 1) * 100;
  scenarios[scenario] = { ...modes, pairDeltas, instrumentationOverheadPercent: overheadPercent,
    instrumentationWithinFivePercent: Math.abs(overheadPercent) <= 5 };
}
mkdirSync(destination);
writeFileSync(join(destination, "timings.json"), JSON.stringify({ schema: 1, scope: "backend process, no application Go runner", scenarios, samples }, null, 2) + "\n");
const phaseNames = [...new Set(samples.flatMap(row => Object.keys(row.phases)))].sort();
const columns = ["scenario", "pair", "compiler", "process", "peakRssBytes", "objectsBytes", "generatedGoBytes", "processEnvelope", "backendResidual", "optimizeResidual", "prepareResidual", "transitiveRounds", "transitiveControl", "avoidableCeiling", "zeroCostPreparationResidual"];
const quote = value => JSON.stringify(value === null || value === undefined ? "" : value);
writeFileSync(join(destination, "timings.csv"), [columns.concat(phaseNames).map(quote).join(","),
  ...samples.map(row => [...columns.map(key => row[key]), ...phaseNames.map(name => row.phases[name])].map(quote).join(","))].join("\n") + "\n");
for (const [scenario, value] of Object.entries(scenarios)) console.log(`${scenario}: ${value.reference.process.median.toFixed(3)} s; instrumentation ${value.instrumentationOverheadPercent.toFixed(2)}%; zero-cost bound ${value.reference.zeroCostPreparationResidual.median.toFixed(3)} s`);
