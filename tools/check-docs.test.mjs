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

const moduleRow = (name, path = `src/${name}.purs`, phase = "Entrée") =>
  `| \`${name}\`<br>\`${path}\` | \`${name}\` | ${phase} | Test module. | — |\n`;

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
  write("docs/module-map.md", "# Map\n\n| Module | Cible | Phase | Responsabilité | Dépend de |\n| --- | --- | --- | --- | --- |\n" + moduleRow("Test"));
  const result = run();
  assert.equal(result.status, 1);
  assert.match(result.stderr, /files entry does not exist: missing\.md/);
});

test("requires one module-map row per source module", t => {
  const { write, run } = fixture(t);
  write("package.json", JSON.stringify({ files: ["README.md"] }));
  write("README.md", "# Root\n");
  write("src/A.purs", "module A where\n");
  write("src/B.purs", "module B where\n");
  write("docs/module-map.md", "# Map\n\n| Module | Cible | Phase | Responsabilité | Dépend de |\n| --- | --- | --- | --- | --- |\n" + moduleRow("A"));
  const result = run();
  assert.equal(result.status, 1);
  assert.match(result.stderr, /missing row for B \(src\/B\.purs\)/);
});

test("rejects a module-map path that does not match the source file", t => {
  const { write, run } = fixture(t);
  write("package.json", JSON.stringify({ files: ["README.md"] }));
  write("README.md", "# Root\n");
  write("src/A.purs", "module A where\n");
  write("docs/module-map.md", "# Map\n\n| Module | Cible | Phase | Responsabilité | Dépend de |\n| --- | --- | --- | --- | --- |\n" + moduleRow("A", "src/Wrong.purs"));
  const result = run();
  assert.equal(result.status, 1);
  assert.match(result.stderr, /declares src\/Wrong\.purs, actual file is src\/A\.purs/);
});
