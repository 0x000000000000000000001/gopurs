import assert from "node:assert/strict";
import { createHash } from "node:crypto";
import { existsSync, mkdirSync, readFileSync, statSync, writeFileSync } from "node:fs";
import { join, resolve } from "node:path";
import { fileURLToPath } from "node:url";

assert(process.argv.length === 3, "Usage: node tools/preparation-compact-analyze.mjs STUDY");
const study = resolve(process.argv[2]), repository = fileURLToPath(new URL("../", import.meta.url));
const load = path => JSON.parse(readFileSync(path));
const hash = bytes => createHash("sha256").update(bytes).digest("hex");
const provenance = load(join(study, "provenance.json")), parent = provenance.parent;
const destination = join(study, "analysis");
assert(!existsSync(destination), "Preserve earlier analyses");
const projectSummary = load(join(study, "project/summary.json")), packsSummary = load(join(study, "packs/summary.json"));
assert(projectSummary.completed && packsSummary.completed);
assert.equal(projectSummary.context, packsSummary.context);
for (const [name, digest] of Object.entries(provenance.inputs)) assert.equal(hash(readFileSync(join(parent, name))), digest, name);
for (const [name, digest] of Object.entries(projectSummary.artifacts)) {
  assert.equal(hash(readFileSync(join(study, "compact/gopurs/gopurs/bin", name === "native" ? "gopurs-native" : "gopurs.js"))), digest);
}
assert.equal(hash(readFileSync(join(study, "packs/pack-probe"))), packsSummary.helper);
const projections = load(join(study, "project/results.json")), rows = load(join(study, "packs/results.json"));
assert.equal(rows.length, 96);
const baseline = load(join(parent, "analysis/timings.json")), oldCosts = load(join(parent, "analysis/costs.json"));
const median = values => { const ordered = [...values].sort((a, b) => a - b); assert(ordered.length); const mid = Math.floor(ordered.length / 2); return ordered.length % 2 ? ordered[mid] : (ordered[mid - 1] + ordered[mid]) / 2; };
const distribution = values => ({ median: median(values), min: Math.min(...values), max: Math.max(...values), samples: values });
const metric = (record, phase) => { const selected = record.metrics.filter(row => row.phase === phase); assert.equal(selected.length, 1, phase); return selected[0].seconds; };
const commands = ["project", "packs"].flatMap(kind => load(join(study, kind, "commands.json")).map(row => ({ kind, ...row })));
const commandEvidence = [];
for (const row of commands) {
  assert.equal(row.status, row.expectedFailure ? 1 : 0, row.label);
  const prefix = join(row.kind, "logs", row.label);
  const paths = [".stdout", ".stderr", ".resources.jsonl", ".request.json"].map(suffix => prefix + suffix).filter(name => existsSync(join(study, name)));
  const files = Object.fromEntries(paths.map(name => { const bytes = readFileSync(join(study, name)); return [name, { bytes: bytes.length, sha256: hash(bytes) }]; }));
  const trace = readFileSync(join(study, prefix + ".resources.jsonl"), "utf8").trim().split("\n").map(line => JSON.parse(line));
  const last = trace.at(-1);
  assert(last.completed && last.status === row.status && !last.contaminated && last.errors.length === 0, row.label);
  assert(!readFileSync(join(study, prefix + ".stderr"), "utf8").includes("ERROR: Rebox missing! b1="), row.label);
  commandEvidence.push({ label: row.label, kind: row.kind, status: row.status, expectedFailure: row.expectedFailure, files });
}

const projection = {};
for (const [name, action, phase] of [
  ["constructWitnesses", "project", "compact/project"],
  ["encodeCompact", "project", "probe/encode-write/compact"],
  ["restoreFull", "restore-contributions", "probe/read-validate-decode/contributions"],
  ["restoreCompact", "restore-compact", "probe/read-validate-decode/compact"],
]) {
  const selected = projections.filter(row => row.state === "base" && row.action === action && !row.javascript);
  assert.equal(selected.length, 3);
  projection[name] = distribution(selected.map(row => metric(row, phase)));
}
const sets = load(join(study, "project/sets.json"));
const unique = new Map(sets.base.map(entry => [entry.object.sha256, entry.object.bytes]));
projection.compactBytes = [...unique.values()].reduce((sum, bytes) => sum + bytes, 0);
projection.fullBytes = oldCosts.categories.contributions.uniqueBytes;
projection.restorePairs = [1, 2, 3].map(pair => {
  const before = projections.find(row => row.pair === pair && row.action === "restore-contributions");
  const after = projections.find(row => row.pair === pair && row.action === "restore-compact");
  const full = metric(before, "probe/read-validate-decode/contributions"), compact = metric(after, "probe/read-validate-decode/compact");
  return { pair, full, compact, seconds: compact - full, percent: 100 * (compact / full - 1) };
});

const storage = {};
for (const model of ["full", "compact"]) {
  storage[model] = {};
  for (const state of ["base", "leaf", "shared", "ffi"]) {
    const result = {};
    for (const layout of ["raw", "packed64"]) {
      const publishes = rows.filter(row => row.model === model && row.state === state && row.layout === layout && row.action === "publish");
      const restores = rows.filter(row => row.model === model && row.state === state && row.layout === layout && row.action === "restore");
      assert.equal(publishes.length, 3); assert.equal(restores.length, 3);
      for (const row of restores) {
        const published = publishes.find(value => value.pair === row.pair);
        const manifest = load(join(study, "packs/manifests", published.label + ".json"));
        const members = manifest.groups.flatMap(group => group.members);
        const kinds = [...new Set(members.map(member => member.kind))];
        row.extractedBytes = members.reduce((sum, member) => sum + member.object.bytes, 0)
          + kinds.reduce((sum, kind) => sum + statSync(join(study, "packs/restored", published.label, kind, "index.json")).size, 0);
        row.cacheAndScratchBytes = published.physicalBytes + row.extractedBytes;
      }
      result[layout] = { publication: distribution(publishes.map(row => row.process)), publicationInternal: distribution(publishes.map(row => row.seconds)),
        extraction: distribution(restores.map(row => row.seconds)), extractionInternal: distribution(restores.map(row => row.internal)),
        bytes: distribution(publishes.map(row => row.bytes)), physicalBytes: distribution(publishes.map(row => row.physicalBytes)),
        extractedBytes: distribution(restores.map(row => row.extractedBytes)), cacheAndScratchBytes: distribution(restores.map(row => row.cacheAndScratchBytes)), groups: publishes.map(row => row.groups) };
    }
    result.pairs = [1, 2, 3].map(pair => {
      const publications = rows.filter(row => row.model === model && row.state === state && row.pair === pair && row.action === "publish");
      const before = publications.find(row => row.layout === "raw").process, after = publications.find(row => row.layout === "packed64").process;
      return { pair, seconds: after - before, percent: 100 * (after / before - 1) };
    });
    storage[model][state] = result;
  }
}

const common = ["prepared", "metadata", "callers"], models = {};
for (const model of ["full", "compact"]) {
  const decode = common.reduce((sum, kind) => sum + oldCosts.categories[kind].restore.median, 0)
    + projection[model === "full" ? "restoreFull" : "restoreCompact"].median;
  const extraction = storage[model].base.packed64.extraction.median;
  const coldEncoding = common.reduce((sum, kind) => sum + oldCosts.categories[kind].encodeWrite.median, 0)
    + (model === "full" ? oldCosts.categories.contributions.encodeWrite.median : projection.constructWitnesses.median + projection.encodeCompact.median);
  const cold = {};
  for (const scenario of ["cold-app", "cold-complete"]) {
    const reference = baseline.scenarios[scenario].reference;
    const credit = median(baseline.samples.filter(row => row.scenario === scenario && row.compiler === "reference")
      .map(row => row.phases["prepared module keys"] + row.phases["global metadata key"]));
    const publication = storage[model].base.packed64.publication.median;
    cold[scenario] = { encoding: coldEncoding, publication, existingKeyCredit: credit,
      extraWithFullCredit: coldEncoding + publication - credit, extraWithoutCredit: coldEncoding + publication,
      regressionBudget: reference.process.median * 0.05,
      note: "Credit assumes existing key serialization can be retained; file-write, GC and integration effects remain unmeasured." };
  }
  const edits = {};
  for (const scenario of ["leaf", "shared", "ffi"]) {
    let changed = common.reduce((sum, kind) => sum + oldCosts.categories[kind].edits[scenario].encodeWrite.median, 0);
    if (model === "full") changed += oldCosts.categories.contributions.edits[scenario].encodeWrite.median;
    else {
      const selected = projections.filter(row => row.state === scenario + "-changed");
      assert.equal(selected.length, 3);
      changed += median(selected.map(row => metric(row, "compact/project") + metric(row, "probe/encode-write/compact")));
    }
    const publication = storage[model][scenario].packed64.publication.median;
    const knownCosts = decode + extraction + changed + publication;
    const samples = baseline.samples.filter(row => row.scenario === scenario && row.compiler === "reference").map(row => {
      const net = row.avoidableCeiling - knownCosts;
      return { pair: row.pair, baseline: row.process, gross: row.avoidableCeiling, net, percent: 100 * net / row.process,
        residual: row.process - net, remainingForTenPercent: net - row.process * 0.10 };
    });
    edits[scenario] = { decode, extraction, changedEncoding: changed, publication, knownCosts, samples,
      withinKnownPreparationBudget: knownCosts <= 11,
      netBeforeSelectiveProofs: distribution(samples.map(row => row.net)), residualBeforeSelectiveProofs: distribution(samples.map(row => row.residual)),
      percentBeforeSelectiveProofs: distribution(samples.map(row => row.percent)), remainingForTenPercent: distribution(samples.map(row => row.remainingForTenPercent)),
      proofImplemented: "global-input fence; TAST edits miss, so this accounting is not a demonstrated selective build gain" };
  }
  models[model] = { decode, extraction, totalRestore: decode + extraction,
    withinRestoreBudget: decode + extraction <= 6.5, currentBytes: storage[model].base.packed64.bytes.median,
    maximumObservedPhysicalBytes: Math.max(...Object.values(storage[model]).map(value => value.packed64.physicalBytes.max)), cold, edits };
  models[model].maximumCacheAndScratchBytes = Math.max(...Object.values(storage[model]).map(value => value.packed64.cacheAndScratchBytes.max));
  models[model].withinCurrentStorageBudget = models[model].currentBytes <= 320 * 1048576;
  models[model].withinKnownStagingBudget = models[model].maximumCacheAndScratchBytes <= 640 * 1048576;
}

const retained = ["provenance.json", "project/harness.mjs", "project/commands.json", "project/results.json", "project/sets.json", "project/summary.json",
  "compact/build-logs/summary.json", "packs/harness.mjs", "packs/commands.json", "packs/results.json", "packs/proofs.json", "packs/summary.json",
  ...["contract.go", "store.go", "session.go", "compiler.go", "modules.go", "packages.go", "main.go"].map(name => "packs/store-source/" + name)];
const evidence = { study, parent, parentInputs: provenance.inputs, projection: projectSummary.artifacts, packHelper: packsSummary.helper,
  sources: Object.fromEntries(retained.map(name => { const bytes = readFileSync(join(study, name)); return [name, { bytes: bytes.length, sha256: hash(bytes) }]; })),
  parentAnalyses: Object.fromEntries(["analysis/timings.json", "analysis/costs.json"].map(name => [name, hash(readFileSync(join(parent, name)))])),
  commandEvidence };
const report = { schema: 1, scope: "C01 bounded study, no integrated backend performance acceptance", candidate: provenance.candidate, evidence,
  projection, storage, models, commands: commands.length, rejectionChecks: commands.filter(row => row.expectedFailure).length,
  proofTransitions: load(join(study, "packs/proofs.json")), optimizationAccepted: false,
  disposition: "C01 rejected: restoration and known preparation time, cold overhead and staging bytes exceed the predeclared budgets; selective proofs are absent.",
  validationScope: { projectionHosts: ["JavaScript", "Go with eight codec workers"], newProjectionCrossHostState: "base",
    backendOracles: "Retained P00 four-host build/hit/forced comparisons; no C01-integrated backend exists",
    newRustProjectionVerified: false, newSequentialGoProjectionVerified: false },
  peakRssBytes: { projection: Math.max(...commands.filter(row => row.kind === "project").map(row => row.peakRssBytes)),
    packHelper: Math.max(...commands.filter(row => row.kind === "packs" && row.executable.endsWith("/pack-probe")).map(row => row.peakRssBytes)) },
  caveats: ["All object families are restored in this model. Lazy/selective restoration is a different unmeasured path.",
    "The complete-input fence authorizes unchanged TAST only. Compact audit digests do not prove dependency validity after edits.",
    "Microbenchmark medians are combined with P00 phases, not substituted for an integrated paired benchmark.",
    "Extraction writes temporary Wire files which are read and checked again by the typed decoder; these costs are counted.",
    "Cache-and-scratch bytes include base/changed CAS objects and all extracted Wire files, not a sampled peak during every filesystem write."] };
mkdirSync(destination);
writeFileSync(join(destination, "harness.mjs"), readFileSync(fileURLToPath(import.meta.url)));
writeFileSync(join(destination, "results.json"), JSON.stringify(report, null, 2) + "\n");
const columns = ["model", "pair", "layout", "state", "action", "label", "seconds", "process", "internal", "bytes", "physicalBytes", "extractedBytes", "cacheAndScratchBytes", "groups", "peakRssBytes"];
const csv = [columns.join(","), ...rows.map(row => columns.map(key => JSON.stringify(row[key] ?? "")).join(","))].join("\n") + "\n";
writeFileSync(join(destination, "timings.csv"), csv);
const outputs = { "docs/build-preparation-compact-c01.json": JSON.stringify(report, null, 2) + "\n", "docs/build-preparation-compact-c01.csv": csv };
for (const name of Object.keys(outputs)) assert(!existsSync(join(repository, name)), "Preserve published study evidence");
for (const [name, data] of Object.entries(outputs)) writeFileSync(join(repository, name), data);
writeFileSync(join(destination, "exports.json"), JSON.stringify(Object.fromEntries(Object.entries(outputs).map(([name, data]) => [name, hash(data)])), null, 2) + "\n");
for (const [name, value] of Object.entries(models)) console.log(`${name}: ${(value.currentBytes / 1048576).toFixed(2)} MiB, restore ${value.totalRestore.toFixed(3)} s, credited cold extra ${value.cold["cold-app"].extraWithFullCredit.toFixed(3)} s`);
