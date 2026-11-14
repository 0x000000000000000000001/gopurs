// One release transition, not a general cross-commit qualification mechanism.
import assert from "node:assert/strict";
import { createHash } from "node:crypto";
import { execFileSync } from "node:child_process";
import { readFileSync } from "node:fs";
import { join } from "node:path";
import { fileURLToPath } from "node:url";

export const recoveryBase = Object.freeze({
  repository: "0x000000000000000000001/gopurs",
  commit: "2202368d4b19849fdc365afacaffa4a617b0ea43",
  tag: "v0.1.0-rc.20",
  tagObject: "9a80c154eff82589275d6790e158e0636066197c",
  runId: "43412666316",
});
export const recoveryTag = "v0.1.1";
export const recoveryManifestPath = "tools/ci/rc20-recovery.json";
export const recoveryPlatforms = Object.freeze([
  { target: "linux-amd64", os: "ubuntu-24.04" },
  { target: "linux-arm64", os: "ubuntu-24.04-arm" },
  { target: "darwin-arm64", os: "macos-15" },
]);
// These release controls are freshly tested by the plan job. No compiler,
// fixture runner, library source, lock dependency or runtime file is exempt.
export const recoveryControls = Object.freeze([
  ".github/workflows/ci.yml",
  recoveryManifestPath,
  "tools/ci/release-recovery.mjs",
  "tools/ci/policy.mjs",
  "tools/ci/policy.test.mjs",
  "tools/ci/release-inputs.mjs",
  "tools/ci/release-inputs.test.mjs",
  "tools/ci/fixture-environment.mjs",
  "tools/ci/fixture-environment.test.mjs",
  "tools/ci/release-finalized.mjs",
  "tools/ci/release-finalized.test.mjs",
  "tools/release/release-quality-report.mjs",
  "tools/release-quality-resources.test.mjs",
  "tools/release/release-prepare.mjs",
]);
export const recoveryDocs = Object.freeze([
  "README.md",
  "CHANGELOG.md",
  "docs/quickstart.md",
  "docs/migration.md",
  "docs/bootstrap.md",
  "docs/ffi-idiomatic.md",
  "docs/publication-scope.md",
  "docs/library-source-audit.md",
  "docs/release-checklist.md",
  "docs/quality-0.1.0.md",
  "docs/release-benchmark.md",
  "docs/limitations.md",
  "docs/release.md",
  "docs/benchmark-results/2026-11-13-rc20.json",
  "docs/benchmark-results/2026-11-13-rc20.csv",
]);
const recoveryDataHashes = Object.freeze({
  "docs/benchmark-results/2026-11-13-rc20.json": "d980cc08ae0298a6637759995d2920820f0e6cbf26a4d724285d80ca175ad62e",
  "docs/benchmark-results/2026-11-13-rc20.csv": "f97b52098fdc14f3b3cabaabb61286a042a398f8466647e60975a97d53f0cae2",
});
// These two existing fixture companions only migrate their FFI calling shape.
// PureScript fixtures and every other executable fixture input remain sealed.
const recoveryFixtureSources = Object.freeze({
  "tests/passing/PolyLabels.go": "tests/passing/PolyLabels.purs",
  "tests/passing/RowUnion.go": "tests/passing/RowUnion.purs",
});
export const recoveryHash = (bytes) => createHash("sha256").update(bytes).digest("hex");
export const isRecovery = (env = process.env) => env.GITHUB_REF === `refs/tags/${recoveryTag}`;
const hash = (value) => assert.match(value, /^[a-f0-9]{64}$/);

export function validateRecoveryManifest(manifest) {
  assert.deepEqual(Object.keys(manifest).sort(), [
    "base",
    "enabled",
    "fixtureSources",
    "goldens",
    "schema",
    "targetTag",
  ]);
  assert.equal(manifest.schema, 1);
  assert.equal(manifest.enabled, true, "RC20 recovery has not been sealed");
  assert.deepEqual(manifest.base, recoveryBase);
  assert.equal(manifest.targetTag, recoveryTag);
  assert.ok(Array.isArray(manifest.goldens) && manifest.goldens.length > 0);
  const paths = new Set();
  for (const golden of manifest.goldens) {
    assert.deepEqual(Object.keys(golden).sort(), ["afterSha256", "beforeSha256", "fixture", "path"]);
    assert.match(golden.path, /^tests\/passing-snapshots(?:-dev)?\/[A-Za-z0-9_-]+\.go$/);
    assert.match(golden.fixture, /^tests\/passing\/[A-Za-z0-9_-]+\.purs$/);
    const name = golden.fixture.split("/").at(-1).slice(0, -5);
    assert.ok(
      [`${name}.go`, `${name}_ffi.go`].includes(golden.path.split("/").at(-1)),
      "Snapshot is not owned by the named fixture",
    );
    assert.ok(!paths.has(golden.path), "Duplicate golden");
    paths.add(golden.path);
    hash(golden.beforeSha256);
    hash(golden.afterSha256);
    assert.notEqual(golden.beforeSha256, golden.afterSha256);
  }
  assert.ok(Array.isArray(manifest.fixtureSources));
  for (const source of manifest.fixtureSources) {
    assert.deepEqual(Object.keys(source).sort(), ["afterSha256", "beforeSha256", "fixture", "path"]);
    assert.ok(Object.hasOwn(recoveryFixtureSources, source.path), "Unreviewed fixture source path");
    assert.equal(source.fixture, recoveryFixtureSources[source.path], "Fixture source owner changed");
    assert.ok(!paths.has(source.path), "Duplicate fixture source");
    paths.add(source.path);
    hash(source.beforeSha256);
    hash(source.afterSha256);
    assert.notEqual(source.beforeSha256, source.afterSha256);
  }
  return manifest;
}

// changes contains all Git changes, with full file bytes and modes, not an
// allowlisted subset supplied by a caller. The Git reader below produces it.
export function validateRecoverySource(changes, manifest) {
  validateRecoveryManifest(manifest);
  const fixtureChanges = new Map([...manifest.goldens, ...manifest.fixtureSources].map((entry) => [entry.path, entry]));
  const seen = new Set();
  const rows = changes.map(({ path, oldMode, newMode, before, after }) => {
    assert.ok(!seen.has(path), "Duplicate source change");
    seen.add(path);
    assert.equal(newMode, "100644", `Nonregular recovery input: ${path}`);
    assert.ok(
      oldMode === "100644" ||
        (oldMode === "000000" && (recoveryControls.includes(path) || recoveryDocs.includes(path))),
      `Removed or nonregular input: ${path}`,
    );
    const beforeSha256 = before === null ? null : recoveryHash(before),
      afterSha256 = recoveryHash(after);
    if (Object.hasOwn(recoveryDataHashes, path)) {
      assert.equal(before, null);
      assert.equal(afterSha256, recoveryDataHashes[path], "Reviewed benchmark export bytes changed");
    } else if (fixtureChanges.has(path)) {
      assert.deepEqual(
        { path, fixture: fixtureChanges.get(path).fixture, beforeSha256, afterSha256 },
        fixtureChanges.get(path),
        `Unreviewed fixture bytes: ${path}`,
      );
    } else if (path === "package.json" || path === "package-lock.json") {
      const original = JSON.parse(before),
        current = JSON.parse(after);
      assert.equal(original.version, "0.1.0-rc.20");
      assert.equal(current.version, "0.1.1");
      original.version = current.version;
      if (path === "package-lock.json") {
        assert.equal(original.packages[""].version, "0.1.0-rc.20");
        assert.equal(current.packages[""].version, "0.1.1");
        original.packages[""].version = current.packages[""].version;
      }
      assert.deepEqual(current, original, `Only release version fields may differ: ${path}`);
    } else
      assert.ok(
        recoveryControls.includes(path) || recoveryDocs.includes(path),
        `Executable/source input changed: ${path}`,
      );
    return { path, oldMode, newMode, beforeSha256, afterSha256 };
  });
  for (const path of [...fixtureChanges.keys(), "package.json", "package-lock.json"])
    assert.ok(seen.has(path), `Missing required release change: ${path}`);
  return {
    manifestSha256: recoveryHash(JSON.stringify(manifest)),
    changes: rows.sort((a, b) => (a.path < b.path ? -1 : a.path > b.path ? 1 : 0)),
  };
}

export function readRecoverySource(root = fileURLToPath(new URL("../../", import.meta.url))) {
  const git = (...args) => execFileSync("git", args, { cwd: root, maxBuffer: 32 * 1024 ** 2 });
  assert.equal(
    git("status", "--porcelain", "--untracked-files=all").toString().trim(),
    "",
    "Recovery requires a clean checkout",
  );
  assert.equal(git("rev-parse", `${recoveryBase.tag}^{commit}`).toString().trim(), recoveryBase.commit);
  assert.equal(git("rev-parse", recoveryBase.tag).toString().trim(), recoveryBase.tagObject);
  const manifest = JSON.parse(readFileSync(join(root, recoveryManifestPath), "utf8"));
  const fields = git("diff", "--raw", "-z", "--no-renames", "--no-abbrev", recoveryBase.commit, "HEAD")
    .toString()
    .split("\0");
  assert.equal(fields.pop(), "");
  assert.equal(fields.length % 2, 0);
  const changes = [];
  for (let index = 0; index < fields.length; index += 2) {
    const match = /^:([0-7]{6}) ([0-7]{6}) [a-f0-9]{40} [a-f0-9]{40} [AM]$/.exec(fields[index]);
    assert.ok(match, "Recovery does not permit deletion, rename or mode changes");
    const path = fields[index + 1];
    changes.push({
      path,
      oldMode: match[1],
      newMode: match[2],
      before: match[1] === "000000" ? null : git("show", `${recoveryBase.commit}:${path}`),
      after: git("show", `HEAD:${path}`),
    });
  }
  return { manifest, proof: validateRecoverySource(changes, manifest) };
}

export function recoveryShards(inventory, manifest) {
  validateRecoveryManifest(manifest);
  assert.equal(inventory.targets.length, 401);
  assert.equal(inventory.batches.length, 162);
  const affected = new Set();
  const includeFixture = (entry, profile) => {
    const matches = inventory.batches.filter(
      (batch) => batch.profile === profile && batch.targets.includes(entry.fixture),
    );
    assert.equal(matches.length, 1, `No unique fixture batch (${profile}): ${entry.path}`);
    const index = matches[0].batchIndex;
    affected.add((index + Math.floor(index / 4)) % 4);
  };
  for (const golden of manifest.goldens) includeFixture(golden, golden.path.includes("-dev/") ? "dev" : "release");
  for (const source of manifest.fixtureSources)
    for (const profile of ["release", "dev"]) includeFixture(source, profile);
  return [...affected].sort();
}

export function validateRecoveryIdentity(context, identity, { core = false, shard } = {}) {
  assert.equal(context.schema, 3);
  assert.deepEqual(context.recovery.base, recoveryBase);
  assert.equal(context.source.tag, recoveryTag);
  assert.notEqual(context.source.commit, recoveryBase.commit);
  assert.equal(context.github.repository, recoveryBase.repository);
  assert.notEqual(context.github.runId, recoveryBase.runId);
  assert.ok(context.qualificationIdentity?.environment, "Missing original compiler and prepared environment identity");
  for (const key of ["javascriptSha256", "nativeSha256"]) {
    hash(identity.compiler?.[key]);
    hash(context.qualificationIdentity.compiler?.[key]);
  }
  const name = core ? `build and test (${context.target})` : `fixtures (${context.target}, shard ${shard})`;
  const producer = context.jobs.find((job) => job.name === name);
  assert.ok(producer, `Missing producer: ${name}`);
  assert.deepEqual(identity.source, { commit: producer.commit, dirty: false });
  assert.deepEqual(identity.github, {
    repository: recoveryBase.repository,
    runId: producer.runId,
    runAttempt: producer.attempt,
  });
  assert.equal(identity.target, context.target);
  if (core) {
    assert.equal(producer.commit, recoveryBase.commit);
    assert.equal(producer.runId, recoveryBase.runId);
  } else {
    const inherited = producer.runId === recoveryBase.runId;
    assert.equal(producer.commit, inherited ? recoveryBase.commit : context.source.commit);
    assert.equal(producer.runId, inherited ? recoveryBase.runId : context.github.runId);
    assert.ok(
      !inherited ||
        !context.recovery.required.some((entry) => entry.target === context.target && entry.shard === shard),
      "Invalidated shard cannot be inherited",
    );
  }
  assert.deepEqual(identity.compiler, context.qualificationIdentity.compiler, "Compiler bytes changed");
  assert.deepEqual(identity.environment, context.qualificationIdentity.environment, "Prepared tools/closure changed");
  return identity;
}

export function validateRecoveryEnvironment(manifest, current, selected) {
  assert.equal(selected.schema, 3);
  assert.equal(selected.source.commit, current.source.commit);
  assert.deepEqual(current.github, {
    repository: selected.github.repository,
    runId: selected.github.runId,
    runAttempt: selected.github.finalizationAttempt,
  });
  assert.equal(manifest.source.commit, recoveryBase.commit);
  assert.equal(manifest.github.runId, recoveryBase.runId);
  assert.equal(manifest.target, current.target);
  assert.deepEqual(manifest.compiler, current.compiler, "Inherited compiler bytes differ");
  validateRecoveryIdentity(selected, selected.qualificationIdentity, { core: true });
  assert.deepEqual(manifest.environment, selected.qualificationIdentity.environment);
  return { ...current, source: manifest.source, github: manifest.github };
}

export function validateRecoverySmokeCommands(commands) {
  assert.ok(Array.isArray(commands) && commands.length > 0);
  const expectedFailures = new Set();
  for (const host of ["javascript", "native"]) {
    const suffix = host === "javascript" ? "/node_modules/.bin/gopurs" : "/node_modules/gopurs/bin/gopurs";
    const drivers = commands
      .map((command, index) => ({ command, index }))
      .filter(
        ({ command }) =>
          typeof command.cwd === "string" &&
          command.cwd.endsWith(`/${host}`) &&
          typeof command.command === "string" &&
          command.command.endsWith(suffix) &&
          JSON.stringify(command.args) === JSON.stringify(["--main", "Main"]),
      );
    // The unchanged smoke runner performs cold, warm, changed signature,
    // deliberately invalid signature, then recovery, for each compiler host.
    assert.deepEqual(
      drivers.map(({ command }) => command.status),
      [0, 0, 0, 1, 0],
      `Missing or unsuccessful installed ${host} cache/signature transitions`,
    );
    expectedFailures.add(drivers[3].index);
  }
  commands.forEach((command, index) =>
    assert.equal(command.status, expectedFailures.has(index) ? 1 : 0, "Unexpected failed installed-package command"),
  );
}

export function validateRecoveryPackage(
  context,
  { build, nativeSha256, javascriptSha256, manifest, sourceManifest, smoke, packageSha256 },
) {
  validateRecoveryIdentity(context, context.qualificationIdentity, { core: true });
  assert.equal(build.schema, 1);
  assert.equal(build.preview, false);
  assert.equal(build.version, "0.1.1");
  assert.equal(build.target, context.target);
  assert.deepEqual(build.source, { commit: context.source.commit, tag: recoveryTag, dirty: false });
  assert.deepEqual(build.compilerOrigin, {
    base: recoveryBase,
    identity: context.qualificationIdentity,
    artifact: context.artifacts[0],
  });
  assert.equal(
    nativeSha256,
    context.qualificationIdentity.compiler.nativeSha256,
    "Final native compiler bytes changed",
  );
  assert.equal(
    javascriptSha256,
    context.qualificationIdentity.compiler.javascriptSha256,
    "Final JavaScript compiler bytes changed",
  );
  assert.equal(manifest.name, "gopurs");
  assert.equal(manifest.version, "0.1.1");
  assert.deepEqual(manifest, sourceManifest, "Packaged metadata differs from the signed final source");
  assert.equal(smoke.version, "0.1.1");
  assert.equal(smoke.packageSHA256, packageSha256, "Smoke tested a different package archive");
  assert.deepEqual(
    smoke.tests.map((entry) => entry.host),
    ["javascript", "native"],
  );
  for (const entry of smoke.tests) {
    assert.equal(entry.passed, true);
    assert.deepEqual(entry.stages, [
      "official-javascript",
      "cold",
      "warm",
      "signature-change",
      "invalid-signature",
      "recovery",
    ]);
  }
  validateRecoverySmokeCommands(smoke.commands);
  return build;
}
