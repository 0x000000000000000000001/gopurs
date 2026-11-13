#!/usr/bin/env node
import assert from "node:assert/strict";
import { execFileSync } from "node:child_process";
import { existsSync, lstatSync, mkdirSync, mkdtempSync, readFileSync, readdirSync, realpathSync } from "node:fs";
import { dirname, isAbsolute, join, resolve } from "node:path";
import { fileURLToPath } from "node:url";
import {
  downloadAndExtract,
  fileSha256,
  flattenPages,
  githubAPI,
  githubContext,
  readVerifiedTag,
  selectLatestJobs,
  selectQualifiedInputs,
  validateFinalizationContext,
  validateReportProducers,
  readRecoveryBase,
  selectRecoveryInputs,
} from "./release-inputs.mjs";
import {
  isRecovery,
  readRecoverySource,
  recoveryBase,
  recoveryTag,
  validateRecoveryIdentity,
} from "./release-recovery.mjs";

const targets = ["linux-amd64", "linux-arm64", "darwin-arm64"];
const limit = 1024 ** 3;
const timestamp = (value) => {
  assert.equal(typeof value, "string");
  assert.match(value, /^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(?:\.\d+)?Z$/);
  const parsed = Date.parse(value);
  assert.ok(Number.isFinite(parsed));
  return parsed;
};

// Selection always uses the real current run. The embedded signing attempt is
// adjusted only after binding it to the selected successful finalizer job.
export function selectFinalizedInputs({ run, jobs, artifacts, tag, context, recovery }) {
  assert.ok(context.tag !== recoveryTag || recovery, "Final release requires the bounded RC20 recovery");
  const latest = selectLatestJobs(jobs, context);
  const names = targets.map((target) => `qualify and attest (${target})`);
  assert.deepEqual(
    [...latest.keys()].filter((name) => name.startsWith("qualify and attest (")).sort(),
    [...names].sort(),
  );
  const selected = targets.map((target, index) => {
    const current = { run, jobs, artifacts, tag, context: { ...context, target } };
    const qualification = recovery ? selectRecoveryInputs({ ...recovery, current }) : selectQualifiedInputs(current);
    const job = latest.get(names[index]);
    assert.equal(job.status, "completed");
    assert.equal(job.conclusion, "success", "The latest finalization job did not pass");
    assert.ok(
      qualification.jobs.every(
        (producer) =>
          (recovery && producer.runId === recoveryBase.runId) || Number(producer.attempt) <= job.run_attempt,
      ),
      "Finalization predates the current qualification",
    );
    assert.ok(job.run_attempt <= Number(context.runAttempt));
    const started = timestamp(job.started_at),
      completed = timestamp(job.completed_at);
    assert.ok(started <= completed);
    for (const qualified of qualification.jobs) {
      const producers =
        recovery && qualified.runId === recoveryBase.runId
          ? selectLatestJobs(recovery.base.jobs, recovery.base.context)
          : latest;
      assert.ok(
        timestamp(producers.get(qualified.name).completed_at) <= started,
        "Finalization started before the selected qualification completed",
      );
    }
    if (recovery) {
      assert.ok(
        timestamp(latest.get(qualification.packageJob.name).completed_at) <= started,
        "Finalization started before the final package smoke completed",
      );
    }
    const name = `release-${target}-attempt-${job.run_attempt}`;
    const matches = artifacts.filter((artifact) => artifact.name === name);
    assert.equal(matches.length, 1, `Expected one artifact named ${name}`);
    const artifact = matches[0];
    assert.ok(
      ![...qualification.artifacts, qualification.environmentArtifact].some((input) => input.id === artifact.id),
      "Release artifact aliases a qualification input",
    );
    assert.ok(Number.isSafeInteger(artifact.id) && artifact.id > 0);
    assert.ok(
      Number.isSafeInteger(artifact.size_in_bytes) && artifact.size_in_bytes > 0 && artifact.size_in_bytes <= limit,
    );
    assert.match(artifact.digest, /^sha256:[a-f0-9]{64}$/);
    assert.equal(artifact.expired, false);
    assert.equal(String(artifact.workflow_run?.id), String(context.runId));
    assert.equal(artifact.workflow_run.head_sha, context.commit);
    const created = timestamp(artifact.created_at);
    assert.ok(
      started <= created && created <= completed,
      "Release artifact was not created during the selected finalizer",
    );
    return {
      target,
      job: { id: job.id, name: job.name, attempt: String(job.run_attempt) },
      artifact: { id: artifact.id, name, size: artifact.size_in_bytes, digest: artifact.digest },
      finalization: {
        ...qualification,
        github: { ...qualification.github, finalizationAttempt: String(job.run_attempt) },
      },
    };
  });
  assert.equal(new Set(selected.map((item) => item.artifact.id)).size, targets.length);
  for (const item of selected) {
    assert.deepEqual(item.finalization.jobs, selected[0].finalization.jobs);
  }
  return selected;
}

export function validateFinalizedQuality(quality, selected, env) {
  assert.equal(quality.schema, 1);
  assert.equal(quality.requestedFull, true);
  assert.equal(quality.full, true);
  assert.equal(quality.phase, "all");
  assert.equal(quality.shard, null);
  assert.equal(quality.status, "passed");
  const expected = { ...selected.finalization, qualificationIdentity: quality.identity };
  assert.deepEqual(quality.finalization, expected, "Embedded finalization differs from the selected API evidence");
  validateFinalizationContext(expected, quality.identity, {
    ...env,
    GOPURS_RELEASE_TARGET: selected.target,
    GITHUB_RUN_ATTEMPT: selected.job.attempt,
  });
  assert.ok(Array.isArray(quality.phases) && quality.phases.length === 5, "Missing phase producer identities");
  const cores = quality.phases.filter((phase) => phase.phase === "core");
  const fixtures = quality.phases.filter((phase) => phase.phase === "fixtures");
  assert.equal(cores.length, 1);
  assert.equal(fixtures.length, 4);
  assert.equal(cores[0].shard, null);
  for (const phase of quality.phases) assert.ok(Number.isSafeInteger(phase.gates) && phase.gates > 0);
  for (const phase of fixtures) assert.equal(phase.shard?.count, 4);
  assert.deepEqual(cores[0].identity, quality.identity);
  assert.deepEqual(fixtures.map((phase) => phase.shard?.index).sort(), [0, 1, 2, 3]);
  const comparable = (identity) => ({ ...identity, github: { ...identity.github, runAttempt: null } });
  for (const phase of fixtures) {
    if (expected.schema === 3) validateRecoveryIdentity(expected, phase.identity, { shard: phase.shard.index });
    else assert.deepEqual(comparable(phase.identity), comparable(quality.identity));
  }
  if (expected.schema === 3) {
    assert.deepEqual(quality.publicationSource, expected.source);
    assert.equal(cores[0].gates, 222);
    assert.equal(quality.gates.length, 384);
    assert.deepEqual(quality.fixtures, { selectedPerProfile: 401, passed: { release: 401, dev: 401 } });
  }
  validateReportProducers(expected, cores[0], fixtures);
  return quality;
}

export function verifyFinalizedDirectory(directory, selected, env) {
  assert.ok(lstatSync(directory).isDirectory() && !lstatSync(directory).isSymbolicLink());
  const version = selected.finalization.source.tag.slice(1),
    prefix = `gopurs-${version}-${selected.target}`;
  const names = readdirSync(directory).sort(),
    files = new Set(names);
  for (const name of names) {
    assert.match(name, /^[A-Za-z0-9][A-Za-z0-9._-]*$/);
    assert.ok(
      name === `gopurs-${version}.tgz` || name === prefix || name.startsWith(prefix + "-"),
      "Unexpected release file",
    );
    const stat = lstatSync(join(directory, name));
    assert.ok(stat.isFile() && !stat.isSymbolicLink());
  }
  const checksumName = `${prefix}-SHA256SUMS`;
  const unchecked = new Set([checksumName, `${prefix}-provenance.sigstore.json`, `${prefix}-sbom.sigstore.json`]);
  for (const name of [
    ...unchecked,
    `${prefix}-quality.json`,
    `${prefix}-build.json`,
    `${prefix}-bootstrap-checksums.txt`,
    `${prefix}-release-notes.md`,
    prefix,
    `gopurs-${version}.tgz`,
  ])
    assert.ok(files.has(name), `Missing release file: ${name}`);
  const rows = readFileSync(join(directory, checksumName), "utf8").trimEnd().split("\n"),
    checked = new Set();
  for (const row of rows) {
    const match = /^([a-f0-9]{64})  ([A-Za-z0-9][A-Za-z0-9._-]*)$/.exec(row);
    assert.ok(match, "Invalid checksum row");
    const [, hash, name] = match;
    assert.ok(files.has(name) && !unchecked.has(name) && !checked.has(name), "Invalid checksum inventory");
    assert.equal(fileSha256(join(directory, name)), hash, `Release checksum differs: ${name}`);
    checked.add(name);
  }
  assert.deepEqual(
    [...checked].sort(),
    names.filter((name) => !unchecked.has(name)),
    "Incomplete checksum inventory",
  );
  const json = (name) => JSON.parse(readFileSync(join(directory, name), "utf8"));
  const quality = validateFinalizedQuality(json(`${prefix}-quality.json`), selected, env);
  const build = json(`${prefix}-build.json`);
  assert.equal(build.schema, 1);
  assert.equal(build.version, version);
  assert.equal(build.target, selected.target);
  assert.deepEqual(build.source, {
    commit: selected.finalization.source.commit,
    tag: selected.finalization.source.tag,
    dirty: false,
  });
  assert.equal(build.preview, false);
  if (selected.finalization.schema === 3) {
    for (const name of [`${prefix}-smoke.json`, `${prefix}-recovery-inputs.tar.gz`])
      assert.ok(checked.has(name), `Missing recovery evidence: ${name}`);
    assert.deepEqual(build.compilerOrigin, {
      base: recoveryBase,
      identity: quality.identity,
      artifact: selected.finalization.artifacts[0],
    });
  }
  assert.equal(
    fileSha256(join(directory, prefix)),
    quality.identity.compiler.nativeSha256,
    "Qualified native compiler differs",
  );
  assert.ok(Array.isArray(build.artifacts));
  const buildNames = new Set();
  for (const artifact of build.artifacts) {
    assert.ok(files.has(artifact.name) && checked.has(artifact.name) && !buildNames.has(artifact.name));
    assert.match(artifact.sha256, /^[a-f0-9]{64}$/);
    assert.equal(fileSha256(join(directory, artifact.name)), artifact.sha256);
    buildNames.add(artifact.name);
  }
  assert.ok(buildNames.has(prefix) && buildNames.has(`${prefix}-bootstrap-checksums.txt`));
  const bootstrap = readFileSync(join(directory, `${prefix}-bootstrap-checksums.txt`), "utf8")
    .split("\n")
    .filter((line) => line.endsWith("  bin/gopurs.js"));
  assert.deepEqual(bootstrap, [`${quality.identity.compiler.javascriptSha256}  bin/gopurs.js`]);
  return {
    target: selected.target,
    artifactId: selected.artifact.id,
    producerAttempts: selected.finalization.jobs.map((job) => ({ name: job.name, attempt: job.attempt })),
    finalizationAttempt: selected.job.attempt,
    files: names.length,
  };
}

async function main(args, env = process.env) {
  assert.ok(args.length === 2 && args[0] === "--out" && args[1], "Expected --out DIRECTORY");
  const root = resolve(fileURLToPath(new URL("../../", import.meta.url))),
    output = resolve(args[1]);
  assert.equal(realpathSync(dirname(output)), realpathSync(root), "Output must be a fresh directory in the checkout");
  assert.ok(!existsSync(output));
  assert.ok(isAbsolute(env.RUNNER_TEMP));
  const context = githubContext({ ...env, GOPURS_RELEASE_TARGET: targets[0] });
  const git = (...values) => execFileSync("git", values, { cwd: root, encoding: "utf8" }).trim();
  assert.equal(git("rev-parse", "HEAD"), context.commit);
  assert.equal(git("status", "--porcelain", "--untracked-files=all"), "");
  const prefix = `repos/${context.repository}/actions/runs/${context.runId}`;
  const run = githubAPI(prefix, { env });
  const jobs = flattenPages(githubAPI(`${prefix}/jobs?filter=all&per_page=100`, { env, paginate: true }), "jobs");
  const artifacts = flattenPages(githubAPI(`${prefix}/artifacts?per_page=100`, { env, paginate: true }), "artifacts");
  const tag = readVerifiedTag(context, { env, root });
  let recovery;
  if (isRecovery(env)) {
    const { qualityInventory } = await import("../release/release-quality-report.mjs");
    recovery = { ...readRecoverySource(root), base: readRecoveryBase(env), inventory: qualityInventory(root) };
  }
  const selected = selectFinalizedInputs({ run, jobs, artifacts, tag, context, recovery });
  const archiveDir = mkdtempSync(join(realpathSync(env.RUNNER_TEMP), "gopurs-finalized-"));
  mkdirSync(output, { mode: 0o700 });
  const results = [];
  let remaining = limit;
  for (const item of selected) {
    const directory = join(output, `release-${item.target}`);
    remaining -= downloadAndExtract(item.artifact, directory, {
      repository: context.repository,
      env,
      maximum: remaining,
      archiveDir,
    }).bytes;
    results.push(verifyFinalizedDirectory(directory, item, env));
  }
  console.log(
    JSON.stringify(
      {
        status: "finalized-artifacts-verified",
        runId: context.runId,
        draftAttempt: context.runAttempt,
        targets: results,
      },
      null,
      2,
    ),
  );
}

if (process.argv[1] && resolve(process.argv[1]) === fileURLToPath(import.meta.url)) await main(process.argv.slice(2));
