import assert from "node:assert/strict";
import { execFileSync } from "node:child_process";
import { createHash } from "node:crypto";
import { chmodSync, existsSync, mkdirSync, mkdtempSync, readFileSync, realpathSync, rmSync, symlinkSync, writeFileSync } from "node:fs";
import { tmpdir } from "node:os";
import { join } from "node:path";
import test from "node:test";
import { collectModuleEvidence, moduleGateName, moduleInventory, validateModuleEvidence, verifyModuleInputs } from "../release-quality-modules.mjs";

function temporary(t) {
  const directory = mkdtempSync(join(tmpdir(), "gopurs-module-quality-"));
  t.after(() => rmSync(directory, { recursive: true, force: true }));
  return directory;
}

const startedAt = "2026-10-30T00:00:00.000Z", finishedAt = "2026-10-30T00:00:01.000Z";
const pinned = [{ name: "gopurs-alpha", rev: "a".repeat(40) }];

function report(root, directory, profile) {
  return { version: 1, kind: "modules", root, startedAt, finishedAt, resumedFrom: null,
    buildProfile: profile, updateSnapshots: false, skipped: [], targets: [
      { target: "gopurs-alpha", status: "passed", directory, startedAt, finishedAt },
    ] };
}

function evidence(profile = "release") {
  return { profile, inputs: structuredClone(pinned),
    report: report("/original-checkout", "/original-workspace/0-gopurs-alpha", profile),
    file: `module-reports/${profile}/gopurs-alpha.json`, logs: [
      { target: "gopurs-alpha", file: `module-reports/${profile}/gopurs-alpha.log`, sha256: "b".repeat(64) },
    ] };
}

function captured(t, profile = "release") {
  const directory = temporary(t), root = join(directory, "checkout"), workspace = join(directory, "campaign");
  mkdirSync(root); mkdirSync(join(workspace, "0-gopurs-alpha", "logs"), { recursive: true });
  const file = join(workspace, "results.json"), log = join(directory, "outer.log"), output = join(directory, "stable");
  const testLog = join(workspace, "0-gopurs-alpha", "logs", "test.log");
  const value = report(realpathSync(root), join(workspace, "0-gopurs-alpha"), profile);
  writeFileSync(file, JSON.stringify(value, null, 2) + "\n");
  writeFileSync(testLog, "Spago compiled\nGo assertions passed\n");
  writeFileSync(log, `Selected 1 modules (explicit selection).\nResults: ${file}\nSummary: 1 passed, 0 failed, 0 pending.\n`);
  return { root, workspace, file, log, output, testLog, value };
}

test("module inventory is sorted and retains the exact manifest pins", t => {
  const root = temporary(t), file = join(root, "bootstrap-sources.json");
  const manifest = { schema: 1, libraries: { beta: { rev: "b".repeat(40) }, alpha: { rev: "a".repeat(40) } } };
  writeFileSync(file, JSON.stringify(manifest));
  assert.deepEqual(moduleInventory(root), [...pinned, { name: "gopurs-beta", rev: "b".repeat(40) }]);
  for (const invalid of [{ ...manifest, schema: 2 }, { schema: 1, libraries: {} },
    { schema: 1, libraries: { alpha: { rev: "main" } } }, { schema: 1, libraries: { "../escape": { rev: "a".repeat(40) } } }]) {
    writeFileSync(file, JSON.stringify(invalid));
    assert.throws(() => moduleInventory(root));
  }
});

test("module gates identify both their library and build profile", () => {
  assert.equal(moduleGateName("release", "gopurs-alpha"), "Library module gopurs-alpha (release)");
  assert.equal(moduleGateName("dev", "gopurs-alpha"), "Library module gopurs-alpha (dev)");
  assert.throws(() => moduleGateName("quick", "gopurs-alpha"));
  assert.throws(() => moduleGateName("release", "../alpha"));
});

test("module inputs require exact selected checkouts, Git pins and clean worktrees", t => {
  const directory = temporary(t), root = join(directory, "checkout"), libraries = join(directory, "libraries");
  mkdirSync(root); mkdirSync(libraries);
  const previous = process.env.GOPURS_LIBRARY_ROOT;
  process.env.GOPURS_LIBRARY_ROOT = libraries;
  t.after(() => { if (previous === undefined) delete process.env.GOPURS_LIBRARY_ROOT; else process.env.GOPURS_LIBRARY_ROOT = previous; });
  const manifest = { schema: 1, libraries: {} };
  for (const name of ["alpha", "beta"]) {
    const directory = join(libraries, "gopurs-" + name);
    mkdirSync(join(directory, "bin"), { recursive: true });
    writeFileSync(join(directory, "spago.yaml"), `package:\n  name: ${name}\n`);
    writeFileSync(join(directory, "bin/test"), "#!/usr/bin/env bash\nexit 0\n", { mode: 0o755 });
    const git = (...args) => execFileSync("git", ["-c", "core.hooksPath=/dev/null", "-c", "commit.gpgsign=false",
      "-c", "user.name=Module quality test", "-c", "user.email=module-quality@example.invalid", ...args], { cwd: directory, encoding: "utf8" }).trim();
    git("init", "-q", "--initial-branch=main"); git("add", "."); git("commit", "-qm", "Fixture");
    manifest.libraries[name] = { rev: git("rev-parse", "HEAD") };
  }
  const file = join(root, "bootstrap-sources.json"); writeFileSync(file, JSON.stringify(manifest));
  const inventory = moduleInventory(root);
  assert.deepEqual(verifyModuleInputs(root, inventory), inventory);
  const dirty = join(libraries, "gopurs-alpha", "untracked.txt"); writeFileSync(dirty, "dirty");
  assert.throws(() => verifyModuleInputs(root, inventory), /Dirty module/); rmSync(dirty);
  const changed = structuredClone(manifest); changed.libraries.alpha.rev = "f".repeat(40);
  writeFileSync(file, JSON.stringify(changed));
  assert.throws(() => verifyModuleInputs(root, inventory), /pins differ/);
  assert.throws(() => verifyModuleInputs(root, moduleInventory(root)), /Wrong module revision/);
  writeFileSync(file, JSON.stringify(manifest));
  const extra = join(libraries, "gopurs-extra"); mkdirSync(join(extra, "bin"), { recursive: true });
  writeFileSync(join(extra, "spago.yaml"), "package:\n  name: extra\n");
  writeFileSync(join(extra, "bin/test"), "#!/bin/sh\n", { mode: 0o755 });
  assert.throws(() => verifyModuleInputs(root, inventory), /Installed module tests differ/); rmSync(extra, { recursive: true });
  chmodSync(join(libraries, "gopurs-alpha", "bin/test"), 0o644);
  assert.throws(() => verifyModuleInputs(root, inventory), /Installed module tests differ/);
});

test("module collection retains exact JSON and regular logs after workspace cleanup", t => {
  for (const profile of ["release", "dev"]) {
    const data = captured(t, profile), original = readFileSync(data.file), logBytes = readFileSync(data.testLog);
    const result = collectModuleEvidence(data.log, data.root, profile, pinned, data.output);
    assert.deepEqual(result.report, data.value);
    assert.equal(result.logs[0].sha256, createHash("sha256").update(logBytes).digest("hex"));
    assert.deepEqual(readFileSync(join(data.output, result.file)), original);
    assert.deepEqual(readFileSync(join(data.output, result.logs[0].file)), logBytes);
    rmSync(data.workspace, { recursive: true });
    assert.equal(validateModuleEvidence(result, profile, pinned), result);
  }
});

test("embedded module evidence rejects incomplete, resumed or mismatched campaigns", () => {
  const valid = evidence();
  assert.equal(validateModuleEvidence(valid, "release", pinned), valid);
  // The finalizer has another checkout and no original workspace; validation stays pure.
  const relocated = structuredClone(valid); relocated.report.root = "/another/qualified-checkout";
  validateModuleEvidence(relocated, "release", pinned);
  const mutations = [
    value => { value.profile = "dev"; }, value => { value.inputs[0].rev = "f".repeat(40); },
    value => { value.report.kind = "fixtures"; }, value => { value.report.buildProfile = "dev"; },
    value => { value.report.resumedFrom = "/previous/results.json"; }, value => { value.report.skipped = ["gopurs-beta"]; },
    value => { value.report.error = "failure"; }, value => { delete value.report.finishedAt; },
    value => { value.report.updateSnapshots = true; }, value => { value.report.targets = []; },
    value => { value.report.targets.push(structuredClone(value.report.targets[0])); },
    value => { value.report.targets[0].status = "failed"; }, value => { value.report.targets[0].status = "pending"; },
    value => { value.report.targets[0].target = "gopurs-beta"; }, value => { value.report.targets[0].error = "failed"; },
    value => { value.file = "module-reports/../../elsewhere.json"; }, value => { value.logs = []; },
    value => { value.logs[0].file = "module-reports/dev/gopurs-alpha.log"; }, value => { value.logs[0].sha256 = "invalid"; },
  ];
  mutations.forEach((mutate, index) => {
    const invalid = structuredClone(valid); mutate(invalid);
    assert.throws(() => validateModuleEvidence(invalid, "release", pinned), undefined, "module mutation " + index);
  });
  assert.throws(() => validateModuleEvidence(valid, "release", [...pinned, { name: "gopurs-beta", rev: "b".repeat(40) }]));
});

test("collection refuses ambiguous results, foreign roots and missing or linked logs", t => {
  for (const kind of ["ambiguous", "root", "missing", "linked", "directory"]) {
    const data = captured(t);
    if (kind === "ambiguous") writeFileSync(data.log, `Results: ${data.file}\nResults: ${data.file}\n`);
    if (kind === "root") { data.value.root += "-other"; writeFileSync(data.file, JSON.stringify(data.value)); }
    if (kind === "missing") rmSync(data.testLog);
    if (kind === "linked") { rmSync(data.testLog); symlinkSync(data.file, data.testLog); }
    if (kind === "directory") { data.value.targets[0].directory = data.root; writeFileSync(data.file, JSON.stringify(data.value)); }
    assert.throws(() => collectModuleEvidence(data.log, data.root, "release", pinned, data.output), undefined, kind);
    assert.equal(existsSync(data.output), false);
  }
});
