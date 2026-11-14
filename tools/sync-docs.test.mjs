import assert from "node:assert/strict";
import { spawnSync } from "node:child_process";
import { mkdirSync, mkdtempSync, readFileSync, rmSync, writeFileSync } from "node:fs";
import { tmpdir } from "node:os";
import { dirname, join } from "node:path";
import { fileURLToPath } from "node:url";
import test from "node:test";

const tool = fileURLToPath(new URL("./sync-docs.mjs", import.meta.url));
const names = {
  "README.md": ["source-version", "install"],
  "docs/quickstart.md": ["source-version", "toolchain", "frontend", "install", "console"],
  "docs/migration.md": ["source-version", "toolchain", "libraries"],
  "docs/publication-scope.md": ["identity"],
};
const history = "## History\nRC14 used `gopurs@0.1.0-rc.14`; its original evidence is unchanged.\n";

function fixture(t) {
  const root = mkdtempSync(join(tmpdir(), "gopurs-sync-docs-"));
  t.after(() => rmSync(root, { recursive: true, force: true }));
  const write = (path, content) => {
    mkdirSync(dirname(join(root, path)), { recursive: true });
    writeFileSync(join(root, path), content);
  };
  const read = (path) => readFileSync(join(root, path), "utf8");
  const json = (path, value) => write(path, JSON.stringify(value));
  const pin = (name, letter) => ({ url: `https://github.com/example/${name}.git`, rev: letter.repeat(40) });
  const sources = {
    compiler: pin("purescript", "a"),
    libraries: { prelude: pin("gopurs-prelude", "b"), console: pin("gopurs-console", "c"), arrays: pin("gopurs-arrays", "d"), strings: pin("gopurs-strings", "e") },
  };
  json("package.json", { version: "0.1.0-rc.14" });
  json("bootstrap-sources.json", sources);
  for (const [file, blocks] of Object.entries(names)) {
    write(file, "# Guide\n\n" + blocks.map((name) => `<!-- gopurs:${name}:start -->\nstale\n<!-- gopurs:${name}:end -->\n`).join("\n") + "\n" + history);
  }
  return { root, read, write, json, sources,
    snapshot: () => Object.fromEntries(Object.keys(names).map((file) => [file, read(file)])),
    run: (...args) => spawnSync(process.execPath, [tool, "--root", root, ...args], { encoding: "utf8" }),
  };
}

test("default check is read-only and writing is idempotent", (t) => {
  const f = fixture(t), before = f.snapshot();
  const stale = f.run();
  assert.equal(stale.status, 1);
  assert.match(stale.stderr, /out of date: README\.md/);
  assert.deepEqual(f.snapshot(), before);
  const updated = f.run("--write");
  assert.equal(updated.status, 0, updated.stderr);
  assert.equal(f.run("--check").status, 0);
  const after = f.snapshot(), repeated = f.run("--write");
  assert.equal(repeated.status, 0, repeated.stderr);
  assert.match(repeated.stdout, /0 changed/);
  assert.deepEqual(f.snapshot(), after);
});

test("a version change updates npm, Nix and source identity without rewriting history", (t) => {
  const f = fixture(t);
  assert.equal(f.run("--write").status, 0);
  f.json("package.json", { version: "0.2.0-rc.1" });
  assert.equal(f.run("--check").status, 1);
  assert.equal(f.run("--write").status, 0);
  for (const file of ["README.md", "docs/quickstart.md", "docs/migration.md"]) {
    assert.match(f.read(file), /npm install --save-dev --ignore-scripts gopurs@0\.2\.0-rc\.1/);
    assert.match(f.read(file), /\/releases\/latest/);
  }
  for (const file of ["docs/quickstart.md", "docs/migration.md"]) {
    assert.match(f.read(file), /nix develop github:0x000000000000000000001\/gopurs\/v0\.2\.0-rc\.1/);
  }
  assert.match(f.read("docs/publication-scope.md"), /Documentation target \| `0\.2\.0-rc\.1`/);
  for (const file of Object.keys(names)) assert(f.read(file).endsWith(history), file);
  assert.equal(f.run().status, 0);
});

test("compiler and library pin changes update their exact install references", (t) => {
  const f = fixture(t);
  assert.equal(f.run("--write").status, 0);
  f.sources.compiler.rev = "f".repeat(40);
  f.sources.libraries.console.rev = "1".repeat(40);
  f.sources.libraries.prelude.rev = "2".repeat(40);
  f.json("bootstrap-sources.json", f.sources);
  const stale = f.run();
  assert.equal(stale.status, 1);
  assert.match(stale.stderr, /out of date: docs\/quickstart\.md/);
  assert.match(stale.stderr, /out of date: docs\/migration\.md/);
  assert.equal(f.run("--write").status, 0);
  assert(f.read("docs/quickstart.md").includes(`https://github.com/example/purescript/tree/${"f".repeat(40)}`));
  assert(f.read("docs/quickstart.md").includes(`git -C gopurs-console checkout --detach ${"1".repeat(40)}`));
  assert(f.read("docs/migration.md").includes(`ref: ${"2".repeat(40)}`));
  assert(f.read("docs/migration.md").includes(`ref: ${"1".repeat(40)}`));
  assert.equal(f.run().status, 0);
});

for (const [name, change] of [
  ["missing", (source) => source.replace("<!-- gopurs:identity:end -->", "")],
  ["duplicate", (source) => source + "<!-- gopurs:identity:start -->\n"],
  ["unknown", (source) => source + "<!-- gopurs:other:start -->\n<!-- gopurs:other:end -->\n"],
  ["reversed", (source) => source.replace("identity:start", "identity:swap").replace("identity:end", "identity:start").replace("identity:swap", "identity:end")],
]) test(`${name} markers fail before any document is written`, (t) => {
  const f = fixture(t);
  f.write("docs/publication-scope.md", change(f.read("docs/publication-scope.md")));
  const before = f.snapshot(), result = f.run("--write");
  assert.equal(result.status, 1);
  assert.match(result.stderr, /docs\/publication-scope\.md: .*markers/);
  assert.deepEqual(f.snapshot(), before);
});

test("nested markers fail before writing", (t) => {
  const f = fixture(t);
  f.write("README.md", ["<!-- gopurs:source-version:start -->", "<!-- gopurs:install:start -->", "<!-- gopurs:install:end -->", "<!-- gopurs:source-version:end -->"].join("\n"));
  const before = f.snapshot(), result = f.run("--write");
  assert.equal(result.status, 1);
  assert.match(result.stderr, /nested generated documentation markers/);
  assert.deepEqual(f.snapshot(), before);
});

test("invalid metadata and conflicting modes cannot modify documentation", (t) => {
  const f = fixture(t), before = f.snapshot();
  assert.equal(f.run("--check", "--write").status, 1);
  f.json("package.json", { version: "latest" });
  assert.match(f.run("--write").stderr, /unsupported documentation version/);
  f.json("package.json", { version: "0.1.1" });
  delete f.sources.libraries.console;
  f.json("bootstrap-sources.json", f.sources);
  assert.match(f.run("--write").stderr, /invalid console source pin/);
  assert.deepEqual(f.snapshot(), before);
});
