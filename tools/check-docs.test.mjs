import assert from "node:assert/strict";
import { mkdirSync, mkdtempSync, rmSync, writeFileSync } from "node:fs";
import { tmpdir } from "node:os";
import { dirname, join } from "node:path";
import { fileURLToPath } from "node:url";
import { spawnSync } from "node:child_process";
import { test } from "node:test";

const tool = fileURLToPath(new URL("./check-docs.mjs", import.meta.url));

const fixture = t => {
  const root = mkdtempSync(join(tmpdir(), "gopurs-check-docs-"));
  t.after(() => rmSync(root, { recursive: true, force: true }));
  const write = (path, content) => {
    const absolute = join(root, path);
    mkdirSync(dirname(absolute), { recursive: true });
    writeFileSync(absolute, content);
  };
  const run = (...args) => spawnSync(process.execPath, [tool, "--root", root, ...args], { encoding: "utf8" });
  return { root, write, run };
};

const moduleRow = (name, path = `src/${name}.purs`, phase = "Entry") =>
  `| \`${name}\`<br>\`${path}\` | ${phase} | Test module. | — |\n`;

test("accepts relative links and GitHub anchors", t => {
  const { write, run } = fixture(t);
  write("README.md", "# Root\n\nSee [the section](docs/page.md#second-title).\n");
  write("docs/page.md", "# First title\n\n## Second title\n");
  const result = run("--links-only");
  assert.equal(result.status, 0, result.stderr);
  assert.match(result.stdout, /2 document\(s\) checked/);
});

test("reports a missing relative target", t => {
  const { write, run } = fixture(t);
  write("README.md", "# Root\n\nSee [missing](docs/nope.md).\n");
  const result = run("--links-only", "README.md");
  assert.equal(result.status, 1);
  assert.match(result.stderr, /missing target docs\/nope\.md/);
});

test("checks build-cache and testing guides by default", t => {
  const { write, run } = fixture(t);
  write("docs/build-cache.md", "# Cache\n\n[old report](build-baseline.md)\n");
  write("docs/testing.md", "# Testing\n\n[old log](history/tests.md)\n");
  const result = run("--links-only");
  assert.equal(result.status, 1);
  assert.match(result.stderr, /docs\/build-cache\.md: missing target/);
  assert.match(result.stderr, /docs\/testing\.md: missing target/);
});

test("preserves consecutive spaces in anchors after punctuation removal", t => {
  const { write, run } = fixture(t);
  write("README.md", "# Root\n\n[results](docs/page.md#local-clean-command-checks--1-november-2026)\n");
  write("docs/page.md", "## Local clean-command checks — 1 November 2026\n");
  const result = run("--links-only", "README.md");
  assert.equal(result.status, 0, result.stderr);
});

test("reports a missing anchor and accepts duplicate-heading suffixes", t => {
  const { write, run } = fixture(t);
  write("README.md", "# Root\n\n[ok](docs/page.md#title-1) [bad](docs/page.md#title-2)\n");
  write("docs/page.md", "## Title\n\n## Title\n");
  const result = run("--links-only");
  assert.equal(result.status, 1);
  assert.match(result.stderr, /missing anchor .*#title-2/);
  assert.doesNotMatch(result.stderr, /#title-1/);
});

test("ignores external links and examples in code blocks", t => {
  const { write, run } = fixture(t);
  write("README.md", "# Root\n\n[web](https://example.com/a#b)\n\n```md\n[x](missing.md)\n```\n");
  const result = run("--links-only", "README.md");
  assert.equal(result.status, 0, result.stderr);
});

test("checks every package.json files entry", t => {
  const { write, run } = fixture(t);
  write("package.json", JSON.stringify({ files: ["README.md", "missing.md"] }));
  write("README.md", "# Root\n");
  write("src/Test.purs", "module Test where\n");
  write("docs/module-map.md", "# Map\n\n| Module | Phase | Responsibility | Depends on |\n| --- | --- | --- | --- |\n" + moduleRow("Test"));
  const result = run();
  assert.equal(result.status, 1);
  assert.match(result.stderr, /files entry does not exist: missing\.md/);
});

test("source-only accepts an unbuilt checkout while the default requires the bundle", t => {
  const { write, run } = fixture(t);
  write("package.json", JSON.stringify({ files: ["README.md", "bin/gopurs.js"] }));
  write("README.md", "# Root\n");
  write("src/Test.purs", "module Test where\n");
  write("docs/module-map.md", "# Map\n\n" + moduleRow("Test"));

  const unbuilt = run();
  assert.equal(unbuilt.status, 1);
  assert.match(unbuilt.stderr, /files entry does not exist: bin\/gopurs\.js/);

  const source = run("--source-only");
  assert.equal(source.status, 0, source.stderr);
  assert.match(source.stdout, /1 packaged file\(s\) present/);
  assert.match(source.stdout, /bin\/gopurs\.js package check deferred until after build/);

  write("bin/gopurs.js", "// Built compiler bundle\n");
  const built = run();
  assert.equal(built.status, 0, built.stderr);
  assert.match(built.stdout, /2 packaged file\(s\) present/);
  assert.doesNotMatch(built.stdout, /deferred/);
});

test("source-only still requires packaged sources and checked-in generated files", t => {
  const { write, run } = fixture(t);
  write("package.json", JSON.stringify({ files: ["README.md", "bin/gopurs.js", "bin/gopurs", "tools/ffi_gen.wasm"] }));
  write("README.md", "# Root\n");
  write("src/Test.purs", "module Test where\n");
  write("docs/module-map.md", "# Map\n\n" + moduleRow("Test"));
  const result = run("--source-only");
  assert.equal(result.status, 1);
  assert.match(result.stderr, /files entry does not exist: bin\/gopurs\n/);
  assert.match(result.stderr, /files entry does not exist: tools\/ffi_gen\.wasm/);
  assert.doesNotMatch(result.stderr, /files entry does not exist: bin\/gopurs\.js/);
});

test("source-only still checks links and complete module-map coverage", t => {
  const { write, run } = fixture(t);
  write("package.json", JSON.stringify({ files: ["README.md", "bin/gopurs.js"] }));
  write("README.md", "# Root\n\n[missing](docs/missing.md)\n");
  write("src/A.purs", "module A where\n");
  write("src/B.purs", "module B where\n");
  write("docs/module-map.md", "# Map\n\n" + moduleRow("A"));
  const result = run("--source-only");
  assert.equal(result.status, 1);
  assert.match(result.stderr, /missing target docs\/missing\.md/);
  assert.match(result.stderr, /missing row for B \(src\/B\.purs\)/);
});

test("requires one module-map row per source module", t => {
  const { write, run } = fixture(t);
  write("package.json", JSON.stringify({ files: ["README.md"] }));
  write("README.md", "# Root\n");
  write("src/A.purs", "module A where\n");
  write("src/B.purs", "module B where\n");
  write("docs/module-map.md", "# Map\n\n| Module | Phase | Responsibility | Depends on |\n| --- | --- | --- | --- |\n" + moduleRow("A"));
  const result = run();
  assert.equal(result.status, 1);
  assert.match(result.stderr, /missing row for B \(src\/B\.purs\)/);
});

test("rejects a module-map path that does not match the source file", t => {
  const { write, run } = fixture(t);
  write("package.json", JSON.stringify({ files: ["README.md"] }));
  write("README.md", "# Root\n");
  write("src/A.purs", "module A where\n");
  write("docs/module-map.md", "# Map\n\n| Module | Phase | Responsibility | Depends on |\n| --- | --- | --- | --- |\n" + moduleRow("A", "src/Wrong.purs"));
  const result = run();
  assert.equal(result.status, 1);
  assert.match(result.stderr, /declares src\/Wrong\.purs, actual file is src\/A\.purs/);
});
