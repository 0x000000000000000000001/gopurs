import assert from "node:assert/strict";
import { readdirSync, readFileSync } from "node:fs";
import { dirname, join, relative } from "node:path";
import { fileURLToPath } from "node:url";
import test from "node:test";

const root = dirname(fileURLToPath(import.meta.url));

function scriptsIn(directory) {
  return readdirSync(join(root, directory), { withFileTypes: true })
    .filter(entry => entry.isFile() && entry.name.endsWith(".mjs") && !entry.name.endsWith(".test.mjs"))
    .map(entry => join(directory, entry.name));
}

test("every managed non-test tool is named in tools/README.md", () => {
  const scripts = [
    ...scriptsIn("."),
    ...scriptsIn("lib"),
    ...scriptsIn("build"),
    ...scriptsIn("release"),
    ...scriptsIn("ci"),
  ].sort();
  const readme = readFileSync(join(root, "README.md"), "utf8");
  const missing = scripts.filter(script => !readme.includes(script));
  assert.deepEqual(missing, [], `Missing tools/README.md entries:\n${missing.join("\n")}`);
});
