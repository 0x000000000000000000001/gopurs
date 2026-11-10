import assert from "node:assert/strict";
import { createHash } from "node:crypto";
import { execFileSync } from "node:child_process";
import {
  closeSync,
  existsSync,
  lstatSync,
  mkdirSync,
  mkdtempSync,
  openSync,
  readFileSync,
  readSync,
  realpathSync,
  writeFileSync,
} from "node:fs";
import { dirname, isAbsolute, join, relative, resolve } from "node:path";
import { fileURLToPath } from "node:url";

const targets = ["linux-amd64", "linux-arm64", "darwin-arm64"];
const workflow = ".github/workflows/ci.yml";
const limit = 1024 ** 3;
const sha = (value) => assert.match(value, /^[a-f0-9]{40}$/);
const positive = (value) => assert.ok(Number.isSafeInteger(value) && value > 0, "Expected a positive safe integer");
const id = (value) => {
  assert.match(String(value), /^[1-9][0-9]*$/);
  positive(Number(value));
  return String(value);
};
const jobNames = targets.flatMap((target) => [
  "build and test (" + target + ")",
  ...[0, 1, 2, 3].map((shard) => "fixtures (" + target + ", shard " + shard + ")"),
]);
const artifactNames = (target) => [
  "candidate-" + target,
  ...[0, 1, 2, 3].map((shard) => "fixtures-" + target + "-" + shard),
];
const producerName = (target, index) =>
  index === 0 ? "build and test (" + target + ")" : "fixtures (" + target + ", shard " + (index - 1) + ")";
const timestamp = (value) => {
  assert.equal(typeof value, "string");
  assert.match(value, /^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(?:\.\d+)?Z$/);
  const result = Date.parse(value);
  assert.ok(Number.isFinite(result), "Invalid GitHub timestamp");
  return result;
};

function validateContext(context) {
  assert.match(context.repository, /^[A-Za-z0-9_.-]+\/[A-Za-z0-9_.-]+$/);
  sha(context.commit);
  assert.match(context.tag, /^v\d+\.\d+\.\d+(?:-rc\.\d+)?$/);
  id(context.runId);
  id(context.runAttempt);
  assert.ok(targets.includes(context.target));
}

function copiedExecution(job, previous) {
  if (!job.created_at || !job.completed_at || timestamp(job.created_at) <= timestamp(job.completed_at)) return false;
  // GitHub may copy a completed job into a rerun with a new ID/attempt and
  // creation time, while retaining the original execution and producing no
  // artifact. Keep its original identity only when the latest real execution
  // matches; never fall back across a newer failed or pending execution.
  assert.ok(previous && previous.run_attempt < job.run_attempt, "Copied job has no original execution");
  assert.ok(job.status === "completed" && job.conclusion === "success", "Copied job did not succeed");
  assert.ok(
    timestamp(previous.created_at) <= timestamp(previous.started_at) &&
      timestamp(previous.started_at) <= timestamp(previous.completed_at),
    "Copied job has invalid original timing",
  );
  assert.ok(Array.isArray(job.steps) && job.steps.length > 0, "Copied job lacks execution steps");
  assert.ok(
    Number.isSafeInteger(job.runner_id) &&
      job.runner_id > 0 &&
      typeof job.runner_name === "string" &&
      job.runner_name.length > 0,
    "Copied job lacks runner identity",
  );
  const execution = (value) =>
    Object.fromEntries(
      ["status", "conclusion", "started_at", "completed_at", "runner_id", "runner_name", "steps"].map((key) => [
        key,
        value[key],
      ]),
    );
  assert.deepEqual(execution(job), execution(previous), "Copied job differs from the latest original execution");
  return true;
}

export function selectLatestJobs(jobs, context) {
  assert.ok(Array.isArray(jobs));
  const latest = new Map(),
    seen = new Set(),
    attempts = new Set();
  for (const job of jobs) {
    positive(job.id);
    positive(job.run_attempt);
    assert.equal(String(job.run_id), id(context.runId));
    assert.equal(job.head_sha, context.commit);
    assert.ok(job.run_attempt <= Number(context.runAttempt), "Job belongs to a future attempt");
    assert.equal(typeof job.name, "string");
    assert.ok(job.name.length > 0);
    assert.ok(!seen.has(job.id), "Duplicate job ID");
    seen.add(job.id);
    const key = JSON.stringify([job.name, job.run_attempt]);
    assert.ok(!attempts.has(key), "Ambiguous jobs in the same attempt");
    attempts.add(key);
  }
  for (const job of [...jobs].sort((left, right) => left.run_attempt - right.run_attempt)) {
    const previous = latest.get(job.name);
    if (!copiedExecution(job, previous)) latest.set(job.name, job);
  }
  return latest;
}

function validateArtifact(artifact) {
  positive(artifact.id);
  positive(artifact.size);
  assert.ok(artifact.size <= limit, "Artifact exceeds 1 GiB");
  assert.match(artifact.digest, /^sha256:[a-f0-9]{64}$/);
}

function validateRun(run, context, preview = false) {
  if (preview) {
    assert.match(context.repository, /^[A-Za-z0-9_.-]+\/[A-Za-z0-9_.-]+$/);
    sha(context.commit);
    id(context.runId);
    id(context.runAttempt);
    assert.ok(targets.includes(context.target));
    assert.match(context.ref, /^refs\/heads\/[^\s]+$/);
  } else validateContext(context);
  assert.equal(String(run.id), id(context.runId));
  assert.equal(String(run.run_attempt), id(context.runAttempt));
  assert.equal(run.head_sha, context.commit);
  assert.equal(run.head_branch, preview ? context.ref.slice(11) : context.tag);
  assert.equal(run.event, preview ? "workflow_dispatch" : "push");
  assert.equal(run.path, workflow);
  assert.equal(run.repository?.full_name, context.repository);
  assert.equal(run.head_repository?.full_name, context.repository);
  // An aggregate run can be queued while its required producers have succeeded.
  // Producer completion and artifact provenance are checked separately below.
  assert.ok(["queued", "in_progress", "completed"].includes(run.status));
}

function validateTag(tag, context) {
  assert.equal(tag.ref?.ref, "refs/tags/" + context.tag);
  assert.equal(tag.ref.object?.type, "tag");
  sha(tag.ref.object.sha);
  assert.equal(tag.object?.sha, tag.ref.object.sha);
  assert.equal(tag.object.tag, context.tag);
  assert.equal(tag.object.object?.type, "commit");
  assert.equal(tag.object.object.sha, context.commit);
  assert.equal(tag.object.verification?.verified, true);
}

function successful(job) {
  assert.ok(job, "Missing producer job");
  assert.equal(job.status, "completed");
  assert.equal(job.conclusion, "success", "The latest qualification job did not pass");
  assert.ok(timestamp(job.started_at) <= timestamp(job.completed_at));
}

export function producerArtifactName(logicalName, attempt) {
  assert.match(
    logicalName,
    /^(candidate|quality|fixture-environment|fixtures)-(linux-amd64|linux-arm64|darwin-arm64)(?:-[0-3])?$/,
  );
  return logicalName + "-attempt-" + id(attempt);
}

function selectArtifact(artifacts, logicalName, producer, context) {
  const name = producerArtifactName(logicalName, producer.run_attempt);
  const matches = artifacts.filter((artifact) => artifact.name === name);
  assert.equal(matches.length, 1, "Expected one artifact named " + name);
  const artifact = matches[0];
  assert.equal(artifact.expired, false);
  assert.equal(String(artifact.workflow_run?.id), id(context.runId));
  assert.equal(artifact.workflow_run.head_sha, context.commit);
  const created = timestamp(artifact.created_at);
  assert.ok(
    timestamp(producer.started_at) <= created && created <= timestamp(producer.completed_at),
    "Artifact was not created during its qualified producer job",
  );
  const result = {
    id: artifact.id,
    name,
    logicalName,
    digest: artifact.digest,
    size: artifact.size_in_bytes,
    producerJobId: producer.id,
    attempt: String(producer.run_attempt),
  };
  validateArtifact(result);
  return result;
}

function selection({ run, jobs, artifacts, tag, context }, { preview = false, fixtures = false } = {}) {
  validateRun(run, context, preview);
  if (!preview) validateTag(tag, context);
  const latest = selectLatestJobs(jobs, context);
  const names = fixtures ? [producerName(context.target, 0)] : jobNames;
  if (!fixtures) {
    const qualification = [...latest.values()].filter((job) => /^(build and test|fixtures) \(/.test(job.name));
    assert.deepEqual(
      qualification.map((job) => job.name).sort(),
      [...jobNames].sort(),
      "Expected exactly three builds and twelve fixture jobs",
    );
  }
  names.forEach((name) => successful(latest.get(name)));
  assert.ok(Array.isArray(artifacts));
  const logicalNames = fixtures
    ? ["candidate-" + context.target, "fixture-environment-" + context.target]
    : artifactNames(context.target);
  const selected = logicalNames.map((name, index) =>
    selectArtifact(artifacts, name, latest.get(producerName(context.target, fixtures ? 0 : index)), context),
  );
  assert.equal(new Set(selected.map((artifact) => artifact.id)).size, selected.length);
  const environmentArtifact = fixtures
    ? undefined
    : selectArtifact(
        artifacts,
        "fixture-environment-" + context.target,
        latest.get(producerName(context.target, 0)),
        context,
      );
  if (environmentArtifact)
    assert.ok(
      !selected.some((artifact) => artifact.id === environmentArtifact.id),
      "Environment artifact aliases another input",
    );
  return {
    schema: 2,
    ...(environmentArtifact ? { environmentArtifact } : {}),
    source: preview
      ? { commit: context.commit, ref: context.ref }
      : { commit: context.commit, tag: context.tag, tagObject: tag.object.sha },
    github: { repository: context.repository, runId: id(context.runId), finalizationAttempt: id(context.runAttempt) },
    target: context.target,
    jobs: names.map((name) => {
      const job = latest.get(name);
      return { id: job.id, name, attempt: String(job.run_attempt) };
    }),
    artifacts: selected,
  };
}

export function selectQualifiedInputs(data) {
  return selection(data);
}
export function selectPreviewInputs(data) {
  return selection(data, { preview: true });
}
export function selectFixtureInputs(data, preview = false) {
  return selection(data, { preview, fixtures: true });
}

export function githubContext(env = process.env) {
  assert.equal(env.CI, "true");
  assert.equal(env.GITHUB_ACTIONS, "true");
  assert.equal(env.RUNNER_ENVIRONMENT, "github-hosted");
  assert.equal(env.GITHUB_SERVER_URL, "https://github.com");
  assert.equal(env.GITHUB_EVENT_NAME, "push");
  assert.equal(env.GITHUB_REF_TYPE, "tag");
  const context = {
    repository: env.GITHUB_REPOSITORY,
    commit: env.GITHUB_SHA,
    tag: env.GITHUB_REF_NAME,
    runId: env.GITHUB_RUN_ID,
    runAttempt: env.GITHUB_RUN_ATTEMPT,
    target: env.GOPURS_RELEASE_TARGET,
  };
  validateContext(context);
  assert.equal(env.GITHUB_REF, "refs/tags/" + context.tag);
  assert.equal(env.GITHUB_WORKFLOW_SHA, context.commit);
  assert.equal(env.GITHUB_WORKFLOW_REF, context.repository + "/" + workflow + "@refs/tags/" + context.tag);
  return context;
}

function previewContext(env) {
  assert.equal(env.CI, "true");
  assert.equal(env.GITHUB_ACTIONS, "true");
  assert.equal(env.RUNNER_ENVIRONMENT, "github-hosted");
  assert.equal(env.GITHUB_SERVER_URL, "https://github.com");
  assert.equal(env.GITHUB_EVENT_NAME, "workflow_dispatch");
  assert.equal(env.GITHUB_REF_TYPE, "branch");
  const context = {
    repository: env.GITHUB_REPOSITORY,
    commit: env.GITHUB_SHA,
    ref: env.GITHUB_REF,
    runId: env.GITHUB_RUN_ID,
    runAttempt: env.GITHUB_RUN_ATTEMPT,
    target: env.GOPURS_RELEASE_TARGET,
  };
  assert.match(context.ref, /^refs\/heads\/[^\s]+$/);
  assert.equal(env.GITHUB_REF_NAME, context.ref.slice(11));
  assert.equal(env.GITHUB_WORKFLOW_SHA, context.commit);
  assert.equal(env.GITHUB_WORKFLOW_REF, context.repository + "/" + workflow + "@" + context.ref);
  return context;
}

export function validateEnvironmentIdentity(environment) {
  assert.ok(environment && typeof environment === "object", "Missing exact qualification environment");
  assert.deepEqual(Object.keys(environment).sort(), ["closureSha256", "profile", "profileSha256", "shell"]);
  assert.match(environment.profile, /^\/nix\/store\/[a-z0-9]{32}-[^/\s]+-env$/);
  assert.match(environment.shell, /^\/nix\/store\/[0-9abcdfghijklmnpqrsvwxyz]{32}-[A-Za-z0-9+._?=-]+\/bin\/bash$/);
  for (const key of ["profileSha256", "closureSha256"]) assert.match(environment[key], /^[a-f0-9]{64}$/);
}

function validateQualificationIdentity(context, identity) {
  assert.deepEqual(identity.source, { commit: context.source.commit, dirty: false });
  assert.equal(identity.target, context.target);
  const build = context.jobs.find((job) => job.name === producerName(context.target, 0));
  assert.ok(build, "Missing original build identity");
  assert.deepEqual(identity.github, {
    repository: context.github.repository,
    runId: context.github.runId,
    runAttempt: build.attempt,
  });
  for (const key of ["javascriptSha256", "nativeSha256"]) assert.match(identity.compiler?.[key], /^[a-f0-9]{64}$/);
  validateEnvironmentIdentity(identity.environment);
}

export function validateFinalizationContext(context, identity, env = process.env, { preview = false } = {}) {
  const current = preview ? previewContext(env) : githubContext(env);
  assert.equal(context.schema, 2);
  assert.equal(context.source.commit, current.commit);
  if (preview) assert.deepEqual(context.source, { commit: current.commit, ref: current.ref });
  else {
    sha(context.source?.tagObject);
    assert.equal(context.source.tag, current.tag);
  }
  assert.equal(context.target, current.target);
  assert.equal(context.github?.repository, current.repository);
  assert.equal(context.github.runId, id(current.runId));
  assert.equal(context.github.finalizationAttempt, id(current.runAttempt));
  assert.deepEqual(Object.keys(context.github).sort(), ["finalizationAttempt", "repository", "runId"]);
  assert.deepEqual(
    context.jobs.map((job) => job.name),
    jobNames,
  );
  for (const job of context.jobs) {
    positive(job.id);
    assert.ok(Number(id(job.attempt)) <= Number(current.runAttempt));
  }
  assert.equal(new Set(context.jobs.map((job) => job.id)).size, jobNames.length);
  assert.deepEqual(
    context.artifacts.map((artifact) => artifact.logicalName),
    artifactNames(current.target),
  );
  assert.equal(new Set(context.artifacts.map((artifact) => artifact.id)).size, context.artifacts.length);
  context.artifacts.forEach((artifact, index) => {
    validateArtifact(artifact);
    const producer = context.jobs.find((job) => job.name === producerName(current.target, index));
    assert.equal(artifact.attempt, producer.attempt);
    assert.equal(artifact.producerJobId, producer.id);
    assert.equal(artifact.name, producerArtifactName(artifact.logicalName, producer.attempt));
  });
  const environment = context.environmentArtifact,
    build = context.jobs.find((job) => job.name === producerName(current.target, 0));
  validateArtifact(environment);
  assert.equal(environment.logicalName, "fixture-environment-" + current.target);
  assert.equal(environment.attempt, build.attempt);
  assert.equal(environment.producerJobId, build.id);
  assert.equal(environment.name, producerArtifactName(environment.logicalName, build.attempt));
  assert.ok(!context.artifacts.some((artifact) => artifact.id === environment.id));
  validateQualificationIdentity(context, identity);
  assert.deepEqual(
    context.qualificationIdentity,
    identity,
    "Qualification identity differs from the downloaded core report",
  );
  return context;
}

export function validateReportProducers(context, core, fixtures) {
  for (const report of [core, ...fixtures]) {
    const index = report.phase === "core" ? 0 : report.shard.index + 1;
    const artifact = context.artifacts[index];
    assert.equal(
      report.identity.github.runAttempt,
      artifact.attempt,
      "Report attempt differs from its actual artifact producer",
    );
    assert.equal(report.identity.github.runId, context.github.runId);
    assert.equal(report.identity.github.repository, context.github.repository);
  }
}

export function githubAPI(endpoint, { env = process.env, paginate = false } = {}) {
  assert.match(endpoint, /^repos\/[A-Za-z0-9_.-]+\/[A-Za-z0-9_.-]+\//);
  const args = ["api", "--hostname", "github.com", ...(paginate ? ["--paginate", "--slurp"] : []), endpoint];
  return JSON.parse(
    execFileSync("gh", args, {
      env: { ...env, GH_HOST: "github.com", GH_PROMPT_DISABLED: "1" },
      timeout: 60_000,
      maxBuffer: 32 * 1024 ** 2,
    }),
  );
}

export function readVerifiedTag(
  context,
  { env = process.env, root = fileURLToPath(new URL("../../", import.meta.url)) } = {},
) {
  validateContext(context);
  const git = (...args) => execFileSync("git", args, { cwd: root, encoding: "utf8" }).trim();
  assert.equal(git("rev-parse", "HEAD"), context.commit);
  assert.equal(git("cat-file", "-t", "refs/tags/" + context.tag), "tag");
  assert.equal(git("rev-parse", "refs/tags/" + context.tag + "^{commit}"), context.commit);
  const prefix = "repos/" + context.repository + "/";
  const ref = githubAPI(prefix + "git/ref/tags/" + context.tag, { env });
  assert.equal(ref.ref, "refs/tags/" + context.tag);
  assert.equal(ref.object?.type, "tag");
  assert.equal(ref.object.sha, git("rev-parse", "refs/tags/" + context.tag));
  const object = githubAPI(prefix + "git/tags/" + ref.object.sha, { env });
  assert.equal(object.sha, ref.object.sha);
  assert.equal(object.tag, context.tag);
  assert.equal(object.object?.type, "commit");
  assert.equal(object.object.sha, context.commit);
  assert.equal(object.verification?.verified, true);
  return { ref, object };
}

export function flattenPages(pages, key) {
  assert.ok(Array.isArray(pages) && pages.length > 0);
  const rows = pages.flatMap((page) => {
    assert.ok(Array.isArray(page[key]));
    assert.equal(page.total_count, pages[0].total_count);
    return page[key];
  });
  assert.equal(rows.length, pages[0].total_count, "Incomplete paginated GitHub response");
  return rows;
}

export function fileSha256(file) {
  const digest = createHash("sha256"),
    buffer = Buffer.alloc(1024 ** 2),
    fd = openSync(file, "r");
  try {
    for (let count; (count = readSync(fd, buffer, 0, buffer.length, null)) > 0;)
      digest.update(buffer.subarray(0, count));
  } finally {
    closeSync(fd);
  }
  return digest.digest("hex");
}

export function verifyArtifactFile(artifact, file) {
  validateArtifact(artifact);
  const stat = lstatSync(file);
  assert.ok(stat.isFile() && !stat.isSymbolicLink());
  assert.equal(stat.size, artifact.size);
  assert.equal("sha256:" + fileSha256(file), artifact.digest, "Artifact SHA-256 differs from GitHub");
}

export function downloadById(artifact, file, { repository, env = process.env }) {
  validateArtifact(artifact);
  assert.match(repository, /^[A-Za-z0-9_.-]+\/[A-Za-z0-9_.-]+$/);
  const fd = openSync(file, "wx", 0o600);
  try {
    execFileSync(
      "gh",
      ["api", "--hostname", "github.com", "repos/" + repository + "/actions/artifacts/" + artifact.id + "/zip"],
      {
        env: { ...env, GH_HOST: "github.com", GH_PROMPT_DISABLED: "1" },
        stdio: ["ignore", fd, "pipe"],
        timeout: 600_000,
        maxBuffer: 1024 ** 2,
      },
    );
  } finally {
    closeSync(fd);
  }
  verifyArtifactFile(artifact, file);
}

const extractPython = String.raw`
import json, os, pathlib, shutil, stat, sys, unicodedata, zipfile
archive, destination, maximum = sys.argv[1], pathlib.Path(sys.argv[2]), int(sys.argv[3])
assert not destination.exists() and not destination.is_symlink(), 'Destination already exists'
with zipfile.ZipFile(archive) as source:
    entries, names, files, total = [], set(), set(), 0
    for info in source.infolist():
        name = info.orig_filename
        assert name == info.filename and name and '\x00' not in name and '\\' not in name, 'Invalid ZIP path'
        assert not name.startswith('/') and not (len(name) > 1 and name[1] == ':'), 'Absolute ZIP path'
        parts = (name[:-1] if info.is_dir() else name).split('/')
        assert all(part not in ('', '.', '..') for part in parts), 'Unsafe ZIP path'
        canonical = unicodedata.normalize('NFC', '/'.join(parts)).casefold()
        assert canonical not in names, 'Duplicate ZIP path'
        names.add(canonical)
        mode = stat.S_IFMT(info.external_attr >> 16)
        assert mode in ((0, stat.S_IFDIR) if info.is_dir() else (0, stat.S_IFREG)), 'Nonregular ZIP member'
        assert not info.flag_bits & 1, 'Encrypted ZIP member'
        total += info.file_size
        assert 0 <= info.file_size and total <= maximum, 'ZIP exceeds extraction budget'
        if not info.is_dir(): files.add(canonical)
        entries.append((info, parts))
    for name in names:
        assert not any('/'.join(name.split('/')[:i]) in files for i in range(1, len(name.split('/')))), 'ZIP file is also a parent directory'
    destination.mkdir(mode=0o700)
    for info, parts in entries:
        target = destination.joinpath(*parts)
        if info.is_dir(): target.mkdir(mode=0o700, parents=True, exist_ok=True)
        else:
            target.parent.mkdir(mode=0o700, parents=True, exist_ok=True)
            with source.open(info) as reader, target.open('xb') as writer:
                shutil.copyfileobj(reader, writer, 1024 * 1024)
            assert target.stat().st_size == info.file_size, 'Extracted size mismatch'
print(json.dumps({'bytes': total}))
`;

export function extractArtifact(artifact, archive, destination, maximum = limit) {
  assert.ok(Number.isSafeInteger(maximum) && maximum >= 0 && maximum <= limit);
  verifyArtifactFile(artifact, archive);
  return JSON.parse(
    execFileSync("python3", ["-I", "-c", extractPython, archive, destination, String(maximum)], {
      timeout: 120_000,
      maxBuffer: 1024 ** 2,
      stdio: "pipe",
    }),
  ).bytes;
}

export function downloadAndExtract(
  artifact,
  destination,
  { repository, env = process.env, maximum = limit, archiveDir } = {},
) {
  assert.ok(isAbsolute(env.RUNNER_TEMP));
  const directory = archiveDir ?? mkdtempSync(join(realpathSync(env.RUNNER_TEMP), "gopurs-artifact-"));
  const archive = join(directory, String(artifact.id) + ".zip");
  downloadById(artifact, archive, { repository, env });
  return { archive, bytes: extractArtifact(artifact, archive, destination, maximum) };
}

function prepare(args, env = process.env) {
  const command = args.shift(),
    fixturesOnly = command === "prepare-fixtures",
    preview = env.GITHUB_REF_TYPE === "branch";
  assert.ok(["prepare", "prepare-preview", "prepare-fixtures"].includes(command));
  if (!fixturesOnly) assert.equal(command === "prepare-preview", preview);
  const current = preview ? previewContext(env) : githubContext(env),
    options = {};
  const second = fixturesOnly ? "--environment" : "--fixtures";
  for (let index = 0; index < args.length; index += 2) {
    const name = args[index];
    assert.ok(["--release", second, "--out"].includes(name) && args[index + 1] && !options[name]);
    options[name] = resolve(args[index + 1]);
  }
  assert.ok(options["--release"] && options[second] && options["--out"]);
  assert.ok(isAbsolute(env.RUNNER_TEMP));
  const temp = realpathSync(env.RUNNER_TEMP);
  const paths = Object.values(options);
  for (const path of paths) {
    const rel = relative(temp, path),
      parent = realpathSync(dirname(path));
    assert.ok(rel && !rel.startsWith("..") && !isAbsolute(rel) && !existsSync(path));
    assert.ok(parent === temp || (!relative(temp, parent).startsWith("..") && !isAbsolute(relative(temp, parent))));
  }
  assert.equal(new Set(paths).size, 3);
  for (const left of paths) for (const right of paths) if (left !== right) assert.ok(!right.startsWith(left + "/"));
  const root = fileURLToPath(new URL("../../", import.meta.url));
  const git = (...values) => execFileSync("git", values, { cwd: root, encoding: "utf8" }).trim();
  assert.equal(git("rev-parse", "HEAD"), current.commit);
  assert.equal(git("status", "--porcelain", "--untracked-files=all"), "");
  const prefix = "repos/" + current.repository + "/";
  const run = githubAPI(prefix + "actions/runs/" + current.runId, { env });
  const jobs = flattenPages(
    githubAPI(prefix + "actions/runs/" + current.runId + "/jobs?filter=all&per_page=100", { env, paginate: true }),
    "jobs",
  );
  const artifacts = flattenPages(
    githubAPI(prefix + "actions/runs/" + current.runId + "/artifacts?per_page=100", { env, paginate: true }),
    "artifacts",
  );
  const tag = preview ? undefined : readVerifiedTag(current, { env, root });
  const selected = selection({ run, jobs, artifacts, tag, context: current }, { preview, fixtures: fixturesOnly });
  const privateDir = mkdtempSync(join(temp, "gopurs-release-inputs-"));
  if (!fixturesOnly) mkdirSync(options[second], { mode: 0o700 });
  let remaining = limit;
  for (const artifact of selected.artifacts) {
    const destination = artifact.logicalName.startsWith("candidate-")
      ? options["--release"]
      : fixturesOnly
        ? options[second]
        : join(options[second], artifact.logicalName);
    remaining -= downloadAndExtract(artifact, destination, {
      repository: current.repository,
      env,
      maximum: remaining,
      archiveDir: privateDir,
    }).bytes;
  }
  const version = JSON.parse(readFileSync(join(root, "package.json"), "utf8")).version;
  const core = JSON.parse(
    readFileSync(join(options["--release"], "gopurs-" + version + "-" + current.target + "-quality.json"), "utf8"),
  );
  validateQualificationIdentity(selected, core.identity);
  selected.qualificationIdentity = structuredClone(core.identity);
  if (!fixturesOnly) validateFinalizationContext(selected, core.identity, env, { preview });
  writeFileSync(options["--out"], JSON.stringify(selected, null, 2) + "\n", { flag: "wx", mode: 0o600 });
  console.log(
    JSON.stringify({
      context: options["--out"],
      producers: selected.jobs,
      finalizationAttempt: selected.github.finalizationAttempt,
    }),
  );
}

if (process.argv[1] && resolve(process.argv[1]) === fileURLToPath(import.meta.url)) prepare(process.argv.slice(2));
