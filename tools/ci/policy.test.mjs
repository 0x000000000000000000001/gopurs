import assert from "node:assert/strict";
import test from "node:test";
import { readFileSync } from "node:fs";
import { buildMatrix, docsOnly, parseChanges, validateTagVersion, fullFixtureMatrix } from "./policy.mjs";
import { recoveryBase, recoveryHash, recoveryShards, validateRecoverySource, validateRecoveryManifest } from "./release-recovery.mjs";
import { qualityInventory } from "../release/release-quality-report.mjs";
import { fileURLToPath } from "node:url";
import { spawnSync } from "node:child_process";
import { tmpdir } from "node:os";
import { join } from "node:path";

const change = (path, oldMode = "100644", newMode = "100644") => ({ path, oldMode, newMode });
const docs = [change("README.md"), change("docs/reference/ffi.md")];
const code = [change("src/Main.purs")];

function recoverySource() {
  const before = Buffer.from("old golden"), after = Buffer.from("correct golden");
  const path = "tests/passing-snapshots/1110.go";
  const manifest = { schema: 1, enabled: true, base: recoveryBase, targetTag: "v0.1.0",
    goldens: [{ path, fixture: "tests/passing/1110.purs", beforeSha256: recoveryHash(before), afterSha256: recoveryHash(after) }] };
  const changes = [{ ...change(path), before, after }, ...["package.json", "package-lock.json"].map(path => ({
    ...change(path), before: JSON.stringify({ version: "0.1.0-rc.20", ...(path.includes("lock") ? { packages: { "": { version: "0.1.0-rc.20" } } } : {}) }),
    after: JSON.stringify({ version: "0.1.0", ...(path.includes("lock") ? { packages: { "": { version: "0.1.0" } } } : {}) }),
  }))];
  return { changes, manifest };
}

test("RC20 recovery accepts only exact goldens and three release version fields", () => {
  const { changes, manifest } = recoverySource();
  const result = validateRecoverySource(changes, manifest);
  assert.equal(result.changes.length, 3);
  assert.equal(result.manifestSha256, recoveryHash(JSON.stringify(manifest)));
  for (const path of ["src/Gopurs/Render.purs", "bootstrap-sources.json", "flake.lock", "spago.lock", "bin/test", "tools/c1c/main.go", "tools/ffi_gen.wasm", "tests/passing/1110.purs", "tests/passing-snapshots/4105.go"])
    assert.throws(() => validateRecoverySource([...changes, { ...change(path), before: "old", after: "new" }], manifest), /input changed/);
  for (const mutate of [
    value => { value.changes[0].after = "different golden"; },
    value => { value.changes[0].newMode = "120000"; },
    value => { value.changes.pop(); },
    value => { value.manifest.enabled = false; },
    value => { value.manifest.base = { ...recoveryBase, commit: "f".repeat(40) }; },
    value => { value.manifest.targetTag = "v0.1.1"; },
    value => { value.changes[2].after = JSON.stringify({ version: "0.1.0", packages: { "": { version: "0.1.0" }, altered: { version: "2" } } }); },
    value => { value.changes[1].after = JSON.stringify({ version: "0.1.0", scripts: { build: "something else" } }); },
  ]) {
    const value = recoverySource(); mutate(value);
    assert.throws(() => validateRecoverySource(value.changes, value.manifest), mutate.toString());
  }
});

test("RC20 recovery maps only the reviewed golden profiles to existing shards", () => {
  const { manifest } = recoverySource();
  const inventory = qualityInventory(fileURLToPath(new URL("../../", import.meta.url)));
  for (const [path, expected] of [["tests/passing-snapshots/1110.go", 0], ["tests/passing-snapshots/4105.go", 1], ["tests/passing-snapshots-dev/4105.go", 2]]) {
    manifest.goldens[0].path = path;
    manifest.goldens[0].fixture = `tests/passing/${path.split("/").at(-1).slice(0, -3)}.purs`;
    assert.deepEqual(recoveryShards(inventory, manifest), [expected]);
  }
  manifest.goldens[0].path = "tests/passing-snapshots/not-a-fixture.go";
  manifest.goldens[0].fixture = "tests/passing/not-a-fixture.purs";
  assert.throws(() => recoveryShards(inventory, manifest), /unique fixture/);
  manifest.goldens.push(structuredClone(manifest.goldens[0]));
  assert.throws(() => validateRecoveryManifest(manifest), /Duplicate golden/);
});

test("release preparation actually activates its explicit recovery context option", () => {
  const missing = join(tmpdir(), `gopurs-required-recovery-context-${process.pid}.json`);
  const result = spawnSync(process.execPath, [fileURLToPath(new URL("../release/release-prepare.mjs", import.meta.url)),
    "--out", join(tmpdir(), "unused-recovery-output"), "--native", join(tmpdir(), "unused-native"),
    "--recovery-context", missing, "--original-directory", tmpdir()], { encoding: "utf8" });
  assert.notEqual(result.status, 0);
  assert.ok(result.stderr.includes(missing) && result.stderr.includes("ENOENT"), result.stderr);
});
const oldSha = "1".repeat(40), newSha = "2".repeat(40), zeroSha = "0".repeat(40);
const raw = (path, status = "M", oldMode = "100644", newMode = "100644") =>
  `:${oldMode} ${newMode} ${status === "A" ? zeroSha : oldSha} ${status === "D" ? zeroSha : newSha} ${status}\0${path}\0`;

function policy(input) {
  const result = buildMatrix(input);
  assert.equal(typeof result.build, "boolean");
  assert.equal(typeof result.reason, "string");
  assert(result.reason.trim().length > 0);
  assert(Array.isArray(result.matrix.include));
  for (const row of result.matrix.include) {
    assert.equal(typeof row.target, "string");
    assert.equal(typeof row.os, "string");
    assert(row.os.length > 0);
  }
  return result;
}

function expectTargets(input, targets) {
  const result = policy(input);
  assert.equal(result.build, true);
  assert.deepEqual(result.matrix.include.map(row => row.target).sort(), [...targets].sort());
  return result;
}

test("docs-only accepts the explicit root whitelist and nested Markdown documents", () => {
  const paths = ["README.md", "CHANGELOG.md", "CODE_OF_CONDUCT.md", "CONTRIBUTING.md", "SECURITY.md",
    "THIRD_PARTY_NOTICES.md", "todo.md", "docs/cli.md", "docs/reference/compiler/types.md"];
  for (const path of paths) assert.equal(docsOnly([change(path)]), true, path);
  assert.equal(docsOnly(paths.map(path => change(path))), true);
  assert.equal(docsOnly([]), false);
});

test("docs-only includes regular-file additions and deletions, never mode or type changes", () => {
  assert.equal(docsOnly([change("docs/added.md", "000000", "100644")]), true);
  assert.equal(docsOnly([change("docs/deleted.md", "100644", "000000")]), true);
  for (const [oldMode, newMode] of [
    ["100755", "100755"], ["100644", "100755"], ["100755", "100644"],
    ["120000", "120000"], ["100644", "120000"], ["120000", "100644"],
    ["000000", "120000"], ["120000", "000000"], ["000000", "100755"],
    ["160000", "160000"], ["040000", "040000"], ["999999", "100644"],
    ["100644", undefined], [undefined, "100644"], ["10064", "100644"],
  ]) assert.equal(docsOnly([{ path: "README.md", oldMode, newMode }]), false, `${oldMode} -> ${newMode}`);
});

test("unlisted paths, deceptive document names and mixed source changes require a build", () => {
  for (const path of [
    "src/README.md", "LICENSE", "package.json", "docs/config.json", ".github/workflows/ci.yml",
    "readme.md", "README.md.backup", "README.md/src/Main.purs", "documentation/cli.md",
    "docs.md", "docs/guide.MD", "docs/new guide.md", "docs/../src/Main.md", "docs//guide.md", "./README.md",
    "/README.md", "docs\\guide.md", "README.md\nsrc/Main.purs", "docs/guide.md\npackage.json",
  ]) assert.equal(docsOnly([change(path)]), false, JSON.stringify(path));
  assert.equal(docsOnly([...docs, ...code]), false);
});

test("tag and explicit full runs always build all three supported targets", () => {
  const supported = [
    { target: "linux-amd64", os: "ubuntu-24.04" },
    { target: "linux-arm64", os: "ubuntu-24.04-arm" },
    { target: "darwin-arm64", os: "macos-15" },
  ];
  const all = supported.map(row => row.target);
  for (const changes of [docs, code, []]) {
    for (const event of [
      { eventName: "push", ref: "refs/tags/v0.1.0-rc.3", full: false },
      { eventName: "workflow_dispatch", ref: "refs/heads/main", full: true },
      { eventName: "pull_request", ref: "refs/pull/4/merge", full: true },
    ]) {
      const result = expectTargets({ ...event, changes }, all);
      assert.deepEqual(result.matrix.include, supported, "build runner matrix");
    }
  }
  const workflow = readFileSync(new URL("../../.github/workflows/ci.yml", import.meta.url), "utf8");
  assert.deepEqual(fullFixtureMatrix.include, supported.flatMap(platform => [0, 1, 2, 3].map(shard => ({ ...platform, shard }))), "complete fixture matrix keeps twelve jobs on the supported runners");
  assert.match(workflow, /matrix: \$\{\{ fromJSON\(needs\.plan\.outputs\.fixture_matrix\) \}\}/);
});

test("ordinary PRs and branch pushes skip only known document-only changes", () => {
  for (const [eventName, ref] of [["pull_request", "refs/pull/4/merge"], ["push", "refs/heads/main"], ["push", "refs/heads/edge"]]) {
    assert.equal(policy({ eventName, ref, full: false, changes: docs }).build, false);
    for (const changes of [code, [], [...docs, ...code], [change("README.md", "100644", "120000")]]) {
      expectTargets({ eventName, ref, full: false, changes }, ["linux-amd64"]);
    }
  }
});

test("manual non-full runs still build Linux for docs and full intent is strictly boolean", () => {
  for (const full of [false, undefined, "true", "false"]) {
    expectTargets({ eventName: "workflow_dispatch", ref: "refs/heads/main", full, changes: docs }, ["linux-amd64"]);
  }
  expectTargets({ eventName: "unknown", ref: "refs/heads/main", full: false, changes: docs }, ["linux-amd64"]);
});

test("raw NUL diff parsing preserves modes and the exact path for A, D and M records", () => {
  const records = raw("README.md") + raw("docs/new guide.md", "A", "000000", "100644")
    + raw("docs/removed.md", "D", "100644", "000000");
  assert.deepEqual(parseChanges(records), [change("README.md"), change("docs/new guide.md", "000000", "100644"),
    change("docs/removed.md", "100644", "000000")]);
  assert.deepEqual(parseChanges(""), []);
});

test("--no-renames represents a document rename as delete plus add without hiding source deletion", () => {
  const renamed = raw("docs/old-name.md", "D", "100644", "000000")
    + raw("docs/new-name.md", "A", "000000", "100644");
  assert.deepEqual(parseChanges(renamed), [change("docs/old-name.md", "100644", "000000"),
    change("docs/new-name.md", "000000", "100644")]);
  assert.equal(docsOnly(parseChanges(renamed)), true);
  const movedSource = raw("src/Main.purs", "D", "100644", "000000")
    + raw("docs/Main.md", "A", "000000", "100644");
  assert.equal(docsOnly(parseChanges(movedSource)), false);
});

test("raw parsing does not split a hostile newline path into trusted document records", () => {
  const path = "README.md\n:100644 100644 " + oldSha + " " + newSha + " M\nsrc/Main.purs";
  assert.deepEqual(parseChanges(raw(path)), [change(path)]);
  assert.equal(docsOnly(parseChanges(raw(path))), false);
  const unknown = "strange\tname\nwith spaces.purs";
  assert.deepEqual(parseChanges(raw(unknown)), [change(unknown)]);
  assert.equal(docsOnly(parseChanges(raw(unknown))), false);
});

test("raw parsing rejects rename/copy/type-change/unmerged and unknown statuses", () => {
  for (const status of ["R100", "C100", "T", "U", "X", "B", "M100", "?"]) {
    assert.throws(() => parseChanges(raw("README.md", status)), status);
  }
  assert.throws(() => parseChanges(raw("old.md", "R100") + "new.md\0"));
});

test("raw parsing rejects truncated or malformed records instead of treating them as docs-only", () => {
  const header = `:100644 100644 ${oldSha} ${newSha} M`;
  for (const input of [header, header + "\0", header + "\0README.md", header + "\0\0", "README.md\0",
    `100644 100644 ${oldSha} ${newSha} M\0README.md\0`,
    `:10064 100644 ${oldSha} ${newSha} M\0README.md\0`,
    raw("README.md") + "unparsed trailing data",
  ]) assert.throws(() => parseChanges(input), JSON.stringify(input));
});

test("early tag validation accepts matching RC and stable release metadata", () => {
  for (const version of ["0.1.0-rc.3", "0.1.0"]) {
    validateTagVersion(`refs/tags/v${version}`, { version }, { version, packages: { "": { version } } },
      `# Changelog\n\n## [Unreleased]\n\n## [${version}] - 2026-10-28\nRelease notes.\n`);
  }
});

test("early tag validation rejects version, tag, lockfile and changelog mismatches", () => {
  const version = "0.1.0-rc.3";
  const valid = () => [`refs/tags/v${version}`, { version }, { version, packages: { "": { version } } }, `## [${version}]\nNotes\n`];
  for (const mutate of [
    a => { a[0] = "refs/tags/v0.1.0-rc.2"; }, a => { a[0] = "refs/heads/main"; },
    a => { a[1].version = "0.1.0-beta.1"; }, a => { a[1].version = "0.1"; },
    a => { a[1].version = "0.1.0+build"; }, a => { delete a[1].version; },
    a => { a[2].version = "0.1.0-rc.2"; }, a => { delete a[2].version; },
    a => { a[2].packages[""].version = "0.1.0-rc.2"; }, a => { a[2].packages = {}; },
    a => { a[3] = "## [Unreleased]\n"; }, a => { a[3] = `Mention ## [${version}] in prose\n`; },
    a => { a[3] = `### [${version}]\n`; },
  ]) {
    const args = valid(); mutate(args); assert.throws(() => validateTagVersion(...args));
  }
});
