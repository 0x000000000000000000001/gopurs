import assert from "node:assert/strict";
import test from "node:test";
import { buildMatrix, docsOnly, parseChanges, validateTagVersion } from "./policy.mjs";

const change = (path, oldMode = "100644", newMode = "100644") => ({ path, oldMode, newMode });
const docs = [change("README.md"), change("docs/reference/ffi.md")];
const code = [change("src/Main.purs")];
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
  const all = ["linux-amd64", "linux-arm64", "darwin-arm64"];
  for (const changes of [docs, code, []]) {
    expectTargets({ eventName: "push", ref: "refs/tags/v0.1.0-rc.3", full: false, changes }, all);
    expectTargets({ eventName: "workflow_dispatch", ref: "refs/heads/main", full: true, changes }, all);
    expectTargets({ eventName: "pull_request", ref: "refs/pull/4/merge", full: true, changes }, all);
  }
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
