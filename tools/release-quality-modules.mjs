import assert from "node:assert/strict";
import { execFileSync } from "node:child_process";
import { createHash } from "node:crypto";
import { lstatSync, mkdirSync, readFileSync, realpathSync, writeFileSync } from "node:fs";
import { basename, dirname, isAbsolute, join } from "node:path";
import { parseOptions, selectModules } from "./test-selection.mjs";

function validateInventory(inventory) {
  assert.ok(Array.isArray(inventory) && inventory.length > 0, "Empty module inventory");
  for (const entry of inventory) {
    assert.match(entry.name, /^gopurs-[a-z0-9]+(?:-[a-z0-9]+)*$/);
    assert.match(entry.rev, /^[a-f0-9]{40}$/);
  }
  const names = inventory.map(entry => entry.name);
  assert.deepEqual(names, [...new Set(names)].sort(), "Module inventory must be unique and sorted");
}

export function moduleInventory(root) {
  const manifest = JSON.parse(readFileSync(join(root, "bootstrap-sources.json"), "utf8"));
  assert.equal(manifest.schema, 1);
  assert.ok(manifest.libraries && typeof manifest.libraries === "object" && !Array.isArray(manifest.libraries));
  const inventory = Object.keys(manifest.libraries).sort().map(name => {
    assert.match(name, /^[a-z0-9]+(?:-[a-z0-9]+)*$/);
    return { name: "gopurs-" + name, rev: manifest.libraries[name].rev };
  });
  validateInventory(inventory);
  return inventory;
}

export function moduleGateName(profile, name) {
  assert.ok(["release", "dev"].includes(profile), "Invalid module profile");
  assert.match(name, /^gopurs-[a-z0-9]+(?:-[a-z0-9]+)*$/);
  return `Library module ${name} (${profile})`;
}

export function verifyModuleInputs(root, inventory) {
  assert.deepEqual(inventory, moduleInventory(root), "Module pins differ from bootstrap-sources.json");
  const selected = selectModules(root, parseOptions(["--all"], { modules: true }));
  assert.deepEqual(selected.map(directory => basename(directory)), inventory.map(entry => entry.name),
    "Installed module tests differ from the pinned inventory");
  selected.forEach((directory, index) => {
    const git = (...args) => execFileSync("git", args, { cwd: directory, encoding: "utf8" }).trim();
    assert.equal(git("rev-parse", "HEAD"), inventory[index].rev, `Wrong module revision: ${inventory[index].name}`);
    assert.equal(git("status", "--porcelain", "--untracked-files=all"), "", `Dirty module: ${inventory[index].name}`);
  });
  return inventory;
}

function completed(value) {
  assert.equal(typeof value.startedAt, "string");
  assert.equal(typeof value.finishedAt, "string");
  const start = Date.parse(value.startedAt), finish = Date.parse(value.finishedAt);
  assert.ok(Number.isFinite(start) && Number.isFinite(finish) && finish >= start, "Incomplete module timing");
  assert.equal(value.error, undefined, "Module campaign contains an error");
}

function validateReport(report, profile, inventory) {
  assert.equal(report.version, 1); assert.equal(report.kind, "modules");
  assert.equal(typeof report.root, "string"); assert.ok(isAbsolute(report.root));
  assert.equal(report.buildProfile, profile); assert.equal(report.updateSnapshots, false);
  assert.equal(report.resumedFrom, null, "A resumed subset cannot qualify all libraries");
  assert.deepEqual(report.skipped, []); completed(report);
  assert.ok(Array.isArray(report.targets));
  assert.deepEqual(report.targets.map(row => row.target), inventory.map(entry => entry.name), "Module coverage differs");
  for (const row of report.targets) {
    assert.equal(row.status, "passed", `Unsuccessful module: ${row.target}`);
    assert.equal(typeof row.directory, "string"); assert.ok(isAbsolute(row.directory));
    completed(row);
  }
}

export function validateModuleEvidence(evidence, profile, inventory) {
  validateInventory(inventory); assert.equal(inventory.length, 1, "Each module gate qualifies one library");
  moduleGateName(profile, inventory[0].name);
  assert.equal(evidence.profile, profile); assert.deepEqual(evidence.inputs, inventory, "Module input pins differ");
  validateReport(evidence.report, profile, inventory);
  assert.equal(evidence.file, `module-reports/${profile}/${inventory[0].name}.json`);
  assert.ok(Array.isArray(evidence.logs));
  assert.deepEqual(evidence.logs.map(log => log.target), inventory.map(entry => entry.name), "Module log coverage differs");
  for (const log of evidence.logs) {
    assert.equal(log.file, `module-reports/${profile}/${log.target}.log`);
    assert.match(log.sha256, /^[a-f0-9]{64}$/);
  }
  return evidence;
}

function regularBytes(file) {
  assert.ok(lstatSync(file).isFile(), `Expected a regular module evidence file: ${file}`);
  return readFileSync(file);
}

export function collectModuleEvidence(log, root, profile, inventory, reportDirectory) {
  validateInventory(inventory); assert.equal(inventory.length, 1, "Each module gate qualifies one library");
  moduleGateName(profile, inventory[0].name);
  const matches = [...regularBytes(log).toString("utf8").matchAll(/^Results: (.+)\r?$/gm)];
  assert.equal(matches.length, 1, "Expected one module campaign Results path");
  const file = matches[0][1].trimEnd();
  assert.ok(isAbsolute(file), "Expected an absolute module campaign report path");
  const bytes = regularBytes(file), report = JSON.parse(bytes.toString("utf8"));
  validateReport(report, profile, inventory);
  assert.equal(report.root, realpathSync(root), "Module campaign belongs to another checkout");
  const workspace = realpathSync(dirname(file));
  // TestCampaign gives each target this private directory; read no arbitrary
  // paths supplied by a malformed report and retain logs before cleanup.
  const logs = report.targets.map((row, index) => {
    const directory = join(workspace, `${index}-${row.target}`);
    assert.equal(realpathSync(row.directory), directory, "Module log directory differs from its campaign");
    const source = join(directory, "logs", "test.log");
    assert.equal(realpathSync(source), source, "Module log path traverses a link");
    const content = regularBytes(source);
    return { target: row.target, file: `module-reports/${profile}/${row.target}.log`,
      sha256: createHash("sha256").update(content).digest("hex"), content };
  });
  const evidence = { profile, inputs: structuredClone(inventory), report,
    file: `module-reports/${profile}/${inventory[0].name}.json`, logs: logs.map(({ content, ...entry }) => entry) };
  validateModuleEvidence(evidence, profile, inventory);
  mkdirSync(join(reportDirectory, "module-reports", profile), { recursive: true });
  for (const entry of logs) writeFileSync(join(reportDirectory, entry.file), entry.content, { flag: "wx" });
  writeFileSync(join(reportDirectory, evidence.file), bytes, { flag: "wx" });
  return evidence;
}
