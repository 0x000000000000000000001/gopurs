#!/usr/bin/env node
import { execFileSync } from "node:child_process";
import { createHash } from "node:crypto";
import { mkdirSync, readFileSync, readdirSync, writeFileSync } from "node:fs";
import { arch, platform } from "node:os";
import { dirname, join, relative, resolve } from "node:path";
import { fileURLToPath } from "node:url";
import { isDeepStrictEqual } from "node:util";
import { fixtureBatches } from "./release-quality-resources.mjs";
import { parseOptions, selectFixtures } from "./test-selection.mjs";
import { validateFinalizationContext } from "./ci/release-inputs.mjs";

export const fixtureShardCount = 4;
export const emissionGateName = "Native emission cancellation and concurrency (race)";
export const packageGateName = "npm package surface";
const requireValue = (condition, message) => { if (!condition) throw new Error(message); };
const sha256 = file => createHash("sha256").update(readFileSync(file)).digest("hex");

export function qualityOptions(args) {
  const options = { full: false, phase: "all", shard: null, report: undefined };
  let shardIndex, shardCount;
  for (let index = 0; index < args.length; index++) {
    const arg = args[index];
    if (arg === "--full") options.full = true;
    else if (["--phase", "--shard", "--shard-count", "--report"].includes(arg) && args[index + 1]) {
      const value = args[++index];
      if (arg === "--phase") options.phase = value;
      else if (arg === "--report") options.report = resolve(value);
      else {
        requireValue(/^\d+$/.test(value), `${arg} must be an integer`);
        if (arg === "--shard") shardIndex = Number(value);
        else shardCount = Number(value);
      }
    } else throw new Error(`Unknown or incomplete option: ${arg}`);
  }
  requireValue(["all", "core", "fixtures"].includes(options.phase), "Unknown quality phase");
  requireValue(options.full || options.phase === "all", "Partial phases require --full");
  if (options.phase === "fixtures") {
    requireValue(Number.isSafeInteger(shardIndex) && Number.isSafeInteger(shardCount)
      && shardCount > 0 && shardIndex >= 0 && shardIndex < shardCount, "Fixtures require a valid --shard and --shard-count");
    options.shard = { index: shardIndex, count: shardCount };
  } else requireValue(shardIndex === undefined && shardCount === undefined, "Shard options require --phase fixtures");
  return options;
}

export function qualityIdentity(root, env = process.env) {
  const git = args => execFileSync("git", args, { cwd: root, encoding: "utf8" }).trim();
  const target = `${platform()}-${{ x64: "amd64", arm64: "arm64" }[arch()] ?? arch()}`;
  requireValue(!env.GOPURS_RELEASE_TARGET || env.GOPURS_RELEASE_TARGET === target, "Release target differs from the executing host");
  const identity = {
    source: { commit: git(["rev-parse", "HEAD"]), dirty: Boolean(git(["status", "--porcelain", "--untracked-files=all"])) },
    compiler: { javascriptSha256: sha256(join(root, "bin/gopurs.js")), nativeSha256: sha256(join(root, "bin/gopurs-native")) },
    target,
    github: { repository: env.GITHUB_REPOSITORY ?? null, runId: env.GITHUB_RUN_ID ?? null, runAttempt: env.GITHUB_RUN_ATTEMPT ?? null },
  };
  validateIdentity(identity);
  return identity;
}

function validateIdentity(identity) {
  requireValue(identity && /^[a-f0-9]{40}$/.test(identity.source?.commit) && typeof identity.source?.dirty === "boolean", "Invalid source identity");
  requireValue(["javascriptSha256", "nativeSha256"].every(key => /^[a-f0-9]{64}$/.test(identity.compiler?.[key])), "Invalid compiler identity");
  requireValue(["linux-amd64", "linux-arm64", "darwin-arm64"].includes(identity.target), "Invalid target identity");
  const github = identity.github;
  requireValue(github && (github.repository === null && github.runId === null && github.runAttempt === null
    || /^[^/\s]+\/[^/\s]+$/.test(github.repository ?? "") && /^[1-9]\d*$/.test(github.runId ?? "")
      && /^[1-9]\d*$/.test(github.runAttempt ?? "")), "Incomplete GitHub run identity");
}

export function qualityInventory(root) {
  const nodeTests = readdirSync(join(root, "tools")).filter(name => name.endsWith(".test.mjs")).sort();
  const { selected, skipped } = selectFixtures(root, parseOptions(["--all"], { env: { UPDATE_SNAPSHOTS: "0" } }));
  requireValue(!skipped.length, "Release quality requires every fixture, without exclusions");
  const targets = selected.map(target => relative(root, target));
  return { nodeTests, targets, batches: fixtureBatches(targets).map((batch, batchIndex) => ({ ...batch, batchIndex,
    name: `${batch.profile} fixtures ${batch.start + 1}-${batch.start + batch.targets.length}` })) };
}

export function coreGateNames(inventory) {
  return [...inventory.nodeTests.map(name => `Node ${name}`), ...["build-cache", "c1c", "ffi-gen"].map(name => `Go ${name} (race)`),
    emissionGateName, packageGateName];
}

export function shardBatches(batches, shard) {
  // Rotate each block so interleaved Release/Dev batches reach every shard.
  // A plain modulo assigns only Release to even shards and only Dev to odd ones.
  return batches.filter((_, index) => (index + Math.floor(index / shard.count)) % shard.count === shard.index);
}

export function validateQualityBuild(identity, build) {
  requireValue(build?.schema === 1 && build.source?.dirty === false && build.source.commit === identity.source.commit,
    "Original build source differs from quality reports");
  requireValue(build.target === identity.target, "Original build target differs from quality reports");
  const name = `gopurs-${build.version}-${build.target}`;
  const artifacts = build.artifacts?.filter(artifact => artifact.name === name);
  requireValue(artifacts?.length === 1 && artifacts[0].sha256 === identity.compiler.nativeSha256,
    "Original build native hash differs from quality reports");
}

function validateGates(report, names) {
  requireValue(Array.isArray(report.gates) && report.gates.length === names.length, `Incomplete ${report.phase} gate inventory`);
  report.gates.forEach((gate, index) => {
    requireValue(gate.index === index + 1 && gate.name === names[index], `Unexpected ${report.phase} gate ${index + 1}`);
    requireValue(gate.passed === true && gate.exitCode === 0 && gate.signal == null && !gate.error,
      `Unsuccessful gate: ${gate.name}`);
  });
}

export function mergeQualityReports(core, fixtures, inventory, build) {
  validateIdentity(core?.identity);
  requireValue(core.identity.source.dirty === false, "Dirty source cannot qualify a full release");
  const reports = [core, ...fixtures];
  for (const report of reports) {
    requireValue(report?.schema === 1 && report.requestedFull === true && report.full === false && report.status === "passed",
      "Only successful partial full-campaign reports can be merged");
    requireValue(isDeepStrictEqual(report.identity, core.identity), "Source/compiler/target/GitHub run identities differ");
  }
  requireValue(core.phase === "core" && core.shard === null, "Expected a core report");
  validateGates(core, coreGateNames(inventory));
  requireValue(fixtures.length === fixtureShardCount, `Exactly ${fixtureShardCount} fixture shards are required`);
  const seen = new Set(), fixtureGates = [];
  for (const report of fixtures) {
    requireValue(report.phase === "fixtures" && report.shard?.count === fixtureShardCount
      && Number.isInteger(report.shard.index) && report.shard.index >= 0 && report.shard.index < fixtureShardCount,
    "Invalid fixture shard identity");
    requireValue(!seen.has(report.shard.index), "Duplicate fixture shard");
    seen.add(report.shard.index);
    const expected = shardBatches(inventory.batches, report.shard);
    validateGates(report, expected.map(batch => batch.name));
    report.gates.forEach((gate, index) => {
      const batch = expected[index], actual = gate.fixtures;
      requireValue(actual?.batchIndex === batch.batchIndex && actual.profile === batch.profile
        && actual.passed === batch.targets.length && isDeepStrictEqual(actual.targets, batch.targets),
      `Fixture coverage differs for ${batch.name}`);
      fixtureGates.push(gate);
    });
    for (const profile of ["release", "dev"]) {
      const count = expected.filter(batch => batch.profile === profile).reduce((total, batch) => total + batch.targets.length, 0);
      requireValue(report.fixtures?.selectedPerProfile === inventory.targets.length && report.fixtures.passed?.[profile] === count,
        "Fixture shard summary differs from its verified batches");
    }
  }
  if (build) validateQualityBuild(core.identity, build);
  fixtureGates.sort((a, b) => a.fixtures.batchIndex - b.fixtures.batchIndex);
  const gates = [...core.gates.slice(0, -1), ...fixtureGates, core.gates.at(-1)].map((gate, index) => ({ ...gate, index: index + 1 }));
  return { schema: 1, requestedFull: true, full: true, phase: "all", shard: null, status: "passed", identity: core.identity,
    fixtures: { selectedPerProfile: inventory.targets.length, passed: { release: inventory.targets.length, dev: inventory.targets.length } },
    phases: reports.map(report => ({ phase: report.phase, shard: report.shard, gates: report.gates.length })), gates };
}

function main(args) {
  const options = { fixtures: [] };
  for (let index = 0; index < args.length; index++) {
    const arg = args[index];
    if (arg === "--help") {
      console.log("Usage: node tools/release-quality-report.mjs --core FILE --fixtures FILE (four times) --build FILE --out FILE [--finalization-context FILE]\nMerge only complete, identical-build core and fixture phases. Keep fixture-reports/ alongside each quality.json. A checked finalization context binds a later signing attempt to the original qualification.");
      return;
    }
    requireValue(["--core", "--fixtures", "--build", "--out", "--finalization-context"].includes(arg) && args[index + 1], `Unknown or incomplete option: ${arg}`);
    const value = resolve(args[++index]);
    if (arg === "--fixtures") options.fixtures.push(value);
    else {
      requireValue(!options[arg.slice(2)], `Repeated option: ${arg}`);
      options[arg.slice(2)] = value;
    }
  }
  requireValue(options.core && options.out && options.build, "--core, --build and --out are required");
  const json = file => JSON.parse(readFileSync(file, "utf8"));
  const root = fileURLToPath(new URL("../", import.meta.url));
  const core = json(options.core), fixtures = options.fixtures.map(json), build = json(options.build);
  const merged = mergeQualityReports(core, fixtures, qualityInventory(root), build);
  const commit = execFileSync("git", ["rev-parse", "HEAD"], { cwd: root, encoding: "utf8" }).trim();
  requireValue(merged.identity.source.commit === commit, "Quality source differs from the validating checkout");
  if (process.env.GITHUB_SHA) requireValue(merged.identity.source.commit === process.env.GITHUB_SHA,
    "Quality source differs from GITHUB_SHA");
  if (options["finalization-context"]) {
    const context = json(options["finalization-context"]);
    validateFinalizationContext(context, merged.identity, process.env);
    merged.finalization = context;
  } else if (process.env.GITHUB_RUN_ID) requireValue(isDeepStrictEqual(merged.identity.github,
    { repository: process.env.GITHUB_REPOSITORY, runId: process.env.GITHUB_RUN_ID, runAttempt: process.env.GITHUB_RUN_ATTEMPT }),
  "Quality reports belong to another GitHub run or attempt");
  for (let index = 0; index < fixtures.length; index++) {
    for (const gate of fixtures[index].gates) {
      const file = gate.fixtures.file;
      requireValue(/^fixture-reports\/[a-zA-Z0-9-]+\.json$/.test(file), "Invalid stable fixture report path");
      const evidence = json(join(dirname(options.fixtures[index]), file));
      requireValue(evidence.version === 1 && evidence.kind === "fixtures" && evidence.buildProfile === gate.fixtures.profile
        && evidence.updateSnapshots === false && Array.isArray(evidence.skipped) && !evidence.skipped.length && evidence.finishedAt
        && Array.isArray(evidence.targets) && evidence.targets.length === gate.fixtures.targets.length
        && evidence.targets.every((row, i) => row.status === "passed" && relative(evidence.root, row.target) === gate.fixtures.targets[i]),
      `Stable fixture evidence differs: ${file}`);
    }
  }
  const bootstrapName = `gopurs-${build.version}-${build.target}-bootstrap-checksums.txt`;
  const bootstrapFile = join(dirname(options.build), bootstrapName);
  const manifest = build.artifacts.filter(artifact => artifact.name === bootstrapName);
  requireValue(manifest.length === 1 && sha256(bootstrapFile) === manifest[0].sha256, "Original bootstrap checksums differ from build metadata");
  const jsRows = readFileSync(bootstrapFile, "utf8").split(/\r?\n/).filter(line => line.endsWith("  bin/gopurs.js"));
  requireValue(jsRows.length === 1 && jsRows[0] === `${merged.identity.compiler.javascriptSha256}  bin/gopurs.js`,
    "Original build JavaScript hash differs from quality reports");
  merged.reports = { core: options.core, fixtures: options.fixtures, build: options.build };
  mkdirSync(dirname(options.out), { recursive: true });
  writeFileSync(options.out, JSON.stringify(merged, null, 2) + "\n", { flag: "wx" });
  console.log(`Release quality: ${merged.gates.length} full gates verified; ${merged.fixtures.passed.release} Release and ${merged.fixtures.passed.dev} Dev fixtures.`);
}

if (process.argv[1] && resolve(process.argv[1]) === fileURLToPath(import.meta.url)) main(process.argv.slice(2));
