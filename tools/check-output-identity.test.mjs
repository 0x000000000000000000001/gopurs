import assert from "node:assert/strict";
import { spawnSync } from "node:child_process";
import { mkdtempSync, mkdirSync, readFileSync, rmSync, writeFileSync } from "node:fs";
import { tmpdir } from "node:os";
import { dirname, join } from "node:path";
import { fileURLToPath } from "node:url";
import { test } from "node:test";

const script = fileURLToPath(new URL("./check-output-identity.mjs", import.meta.url));

const run = args => spawnSync(process.execPath, [script, ...args], { encoding: "utf8" });

const workspace = t => {
  const directory = mkdtempSync(join(tmpdir(), "gopurs-output-identity-"));
  t.after(() => rmSync(directory, { recursive: true, force: true }));
  return directory;
};

const tree = t => {
  const directory = workspace(t);
  const root = join(directory, "tree");
  mkdirSync(root);
  return { directory, root };
};

const write = (root, path, content) => {
  mkdirSync(dirname(join(root, path)), { recursive: true });
  writeFileSync(join(root, path), content);
};

test("snapshot then compare accepts an unchanged tree", t => {
  const { directory, root } = tree(t);
  write(root, "a.go", "package a\n");
  write(root, "nested/b.go", "package b\n");
  const manifest = join(directory, "manifest.json");
  assert.equal(run(["snapshot", root, manifest]).status, 0);
  const result = run(["compare", root, manifest]);
  assert.equal(result.status, 0);
  assert.match(result.stdout, /2 file\(s\), identical/);
});

test("compare reports changed, missing and extra files", t => {
  const { directory, root } = tree(t);
  write(root, "a.go", "package a\n");
  write(root, "nested/b.go", "package b\n");
  const manifest = join(directory, "manifest.json");
  assert.equal(run(["snapshot", root, manifest]).status, 0);

  write(root, "a.go", "package changed\n");
  rmSync(join(root, "nested/b.go"));
  write(root, "extra.go", "package extra\n");
  const result = run(["compare", root, manifest]);
  assert.equal(result.status, 1);
  assert.match(result.stderr, /3 difference\(s\): 1 missing, 1 extra, 1 changed/);
  assert.match(result.stderr, /missing nested\/b\.go/);
  assert.match(result.stderr, /extra\.go/);
  assert.match(result.stderr, /changed a\.go/);
});

test("ignore patterns skip matching paths in both directions", t => {
  const { directory, root } = tree(t);
  write(root, "a.go", "package a\n");
  write(root, "nested/b.go", "package b\n");
  const manifest = join(directory, "manifest.json");
  assert.equal(run(["snapshot", root, manifest]).status, 0);

  write(root, "a.go", "package changed\n");
  write(root, "nested/b.go", "package changed\n");
  const ignored = run(["compare", root, manifest, "--ignore", "a.go", "--ignore", "nested/**"]);
  assert.equal(ignored.status, 0);
});

test("diff compares two trees directly", t => {
  const baseline = tree(t);
  const current = tree(t);
  write(baseline.root, "a.go", "package a\n");
  write(current.root, "a.go", "package a\n");
  assert.equal(run(["diff", baseline.root, current.root]).status, 0);

  write(current.root, "a.go", "package c\n");
  const result = run(["diff", baseline.root, current.root]);
  assert.equal(result.status, 1);
  assert.match(result.stderr, /1 difference\(s\): 0 missing, 0 extra, 1 changed/);
});

test("snapshots are deterministic", t => {
  const { directory, root } = tree(t);
  write(root, "a.go", "package a\n");
  write(root, "nested/b.go", "package b\n");
  const first = join(directory, "first.json");
  const second = join(directory, "second.json");
  assert.equal(run(["snapshot", root, first]).status, 0);
  assert.equal(run(["snapshot", root, second]).status, 0);
  assert.equal(readFileSync(first, "utf8"), readFileSync(second, "utf8"));
});

test("usage errors exit 2", t => {
  assert.equal(run([]).status, 2);
  assert.equal(run(["compare", "only-one"]).status, 2);
  assert.equal(run(["unknown"]).status, 2);
  assert.equal(run(["snapshot", "missing", join(workspace(t), "m.json")]).status, 2);
});

test("help exits 0", () => {
  const result = run(["--help"]);
  assert.equal(result.status, 0);
  assert.match(result.stderr, /check-output-identity/);
});

test("an ignored path is excluded from both snapshot and comparison", t => {
  const { directory, root } = tree(t);
  write(root, "a.go", "package a\n");
  write(root, "ignored/b.go", "package b\n");
  const manifest = join(directory, "manifest.json");
  assert.equal(run(["snapshot", root, manifest, "--ignore", "ignored/**"]).status, 0);
  write(root, "ignored/b.go", "package changed\n");
  assert.equal(run(["compare", root, manifest, "--ignore", "ignored/**"]).status, 0);
  assert.equal(run(["compare", root, manifest]).status, 1);
});
