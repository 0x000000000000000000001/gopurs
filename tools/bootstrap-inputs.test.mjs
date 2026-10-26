import assert from "node:assert/strict";
import { execFileSync, spawnSync } from "node:child_process";
import { existsSync, mkdirSync, mkdtempSync, readFileSync, realpathSync, rmSync, writeFileSync } from "node:fs";
import { tmpdir } from "node:os";
import { join } from "node:path";
import test from "node:test";
import { findTypedCompiler, nativeWorkspaceConfig } from "./native-workspace.mjs";

function fixture(t) {
  const directory = mkdtempSync(join(tmpdir(), "gopurs-bootstrap-inputs-"));
  t.after(() => rmSync(directory, { recursive: true, force: true }));
  const root = join(directory, "compiler");
  const upstream = join(directory, "upstream");
  const destination = join(directory, "libraries");
  for (const path of [root, upstream, destination]) mkdirSync(path);
  const git = (args, cwd = upstream) => execFileSync("git", args, { cwd, encoding: "utf8", stdio: ["ignore", "pipe", "pipe"] }).trim();
  git(["init", "-q"]);
  git(["config", "user.name", "Bootstrap test"]);
  git(["config", "user.email", "bootstrap@example.invalid"]);
  writeFileSync(join(upstream, "spago.yaml"), "package:\n  name: console\n");
  git(["add", "."]);
  git(["commit", "-qm", "Pinned source"]);
  const pinned = git(["rev-parse", "HEAD"]);
  writeFileSync(join(upstream, "later.txt"), "Must not appear in the pinned checkout\n");
  git(["add", "."]);
  git(["commit", "-qm", "Later source"]);
  writeFileSync(join(root, "bootstrap-sources.json"), JSON.stringify({
    schema: 1, libraries: {
      console: { url: upstream, rev: pinned },
      effect: { url: upstream, rev: pinned },
    },
  }));
  const run = (locked, names = ["console"]) => spawnSync(process.execPath,
    [new URL("./setup-libraries.mjs", import.meta.url).pathname, root, destination, locked ? "1" : "0", ...names],
    { encoding: "utf8", timeout: 30000 });
  return { root, directory, destination, git, pinned, run };
}

test("setup installs the pinned commit, not the repository's current branch", t => {
  const f = fixture(t);
  const result = f.run(true);
  assert.equal(result.status, 0, result.stderr);
  const checkout = join(f.destination, "gopurs-console");
  assert.equal(f.git(["rev-parse", "HEAD"], checkout), f.pinned);
  assert.equal(existsSync(join(checkout, "later.txt")), false);
  assert.equal(f.run(true).status, 0, "locked setup is idempotent");
});

test("locked setup rejects local changes before cloning another library", t => {
  const f = fixture(t);
  assert.equal(f.run(true).status, 0);
  const source = join(f.destination, "gopurs-console/spago.yaml");
  writeFileSync(source, "local work\n");
  const result = f.run(true, ["console", "effect"]);
  assert.equal(result.status, 1);
  assert.match(result.stderr, /has local changes/);
  assert.equal(existsSync(join(f.destination, "gopurs-effect")), false);
  assert.equal(readFileSync(source, "utf8"), "local work\n");
  assert.equal(f.run(false).status, 0, "development setup preserves the checkout");
});

test("locked setup rejects a different source revision without resetting it", t => {
  const f = fixture(t);
  assert.equal(f.run(true).status, 0);
  const checkout = join(f.destination, "gopurs-console");
  f.git(["checkout", "--detach", "origin/HEAD"], checkout);
  const before = f.git(["rev-parse", "HEAD"], checkout);
  const result = f.run(true);
  assert.equal(result.status, 1);
  assert.match(result.stderr, /release requires/);
  assert.equal(f.git(["rev-parse", "HEAD"], checkout), before);
});

test("native workspace selects the lockfile optimizer and versioned codec input", t => {
  const f = fixture(t);
  const previous = { pbo: process.env.GOPURS_PBO_DIR, libraries: process.env.GOPURS_LIBRARY_ROOT };
  delete process.env.GOPURS_PBO_DIR;
  process.env.GOPURS_LIBRARY_ROOT = f.destination;
  t.after(() => {
    for (const [key, value] of [["GOPURS_PBO_DIR", previous.pbo], ["GOPURS_LIBRARY_ROOT", previous.libraries]]) {
      if (value === undefined) delete process.env[key]; else process.env[key] = value;
    }
  });
  writeFileSync(join(f.root, "spago.yaml"), "package:\n  name: gopurs\n  dependencies:\n    - backend-optimizer\nworkspace:\n");
  writeFileSync(join(f.root, "spago.lock"), JSON.stringify({ packages: { "backend-optimizer": { type: "git", rev: f.pinned } } }));
  const optimizer = join(f.root, ".spago/p/backend-optimizer", f.pinned);
  const codecs = join(f.root, "support/argonaut-codecs");
  for (const path of [optimizer, codecs]) {
    mkdirSync(path, { recursive: true });
    writeFileSync(join(path, "spago.yaml"), "package:\n");
  }
  assert.equal(f.run(true).status, 0);
  const config = nativeWorkspaceConfig(f.root);
  assert.ok(config.includes(JSON.stringify(optimizer)));
  assert.ok(config.includes(JSON.stringify(codecs)));
  assert.ok(config.includes(JSON.stringify(join(f.destination, "gopurs-console"))));
  process.env.GOPURS_PBO_DIR = codecs;
  assert.ok(nativeWorkspaceConfig(f.root).includes(`    backend-optimizer:\n      path: ${JSON.stringify(codecs)}`));
});

test("typed compiler discovery honors GOPURS_PURS without a sibling compiler checkout", t => {
  const f = fixture(t);
  const configured = join(f.directory, "nix-purs");
  const explicit = join(f.directory, "explicit-purs");
  for (const file of [configured, explicit]) writeFileSync(file, "#!/bin/sh\nexit 0\n", { mode: 0o755 });
  const previous = process.env.GOPURS_PURS;
  process.env.GOPURS_PURS = configured;
  t.after(() => {
    if (previous === undefined) delete process.env.GOPURS_PURS;
    else process.env.GOPURS_PURS = previous;
  });
  assert.equal(findTypedCompiler(f.root), realpathSync(configured));
  assert.equal(findTypedCompiler(f.root, explicit), realpathSync(explicit), "an explicit argument still overrides the environment");
  delete process.env.GOPURS_PURS;
  assert.throws(() => findTypedCompiler(f.root), /No local TAST compiler found/);
});
