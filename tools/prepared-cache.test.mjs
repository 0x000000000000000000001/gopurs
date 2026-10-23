import assert from "node:assert/strict";
import { spawnSync } from "node:child_process";
import { mkdtempSync, mkdirSync, rmSync, writeFileSync } from "node:fs";
import { tmpdir } from "node:os";
import { join } from "node:path";
import { fileURLToPath } from "node:url";
import test from "node:test";
import { Left } from "../output/Data.Either/index.js";
import { Cons } from "../output/Data.List.Types/index.js";
import * as Map from "../output/Data.Map/index.js";
import { Nothing } from "../output/Data.Maybe/index.js";
import * as Aff from "../output/Effect.Aff/index.js";
import * as Prepare from "../output/Gopurs.Driver.Prepare/index.js";
import * as Wire from "../output/Gopurs.Cache.Wire/index.js";
import { findTypedCompiler } from "./native-workspace.mjs";

const root = fileURLToPath(new URL("../", import.meta.url));
const runAff = action => new Promise((resolve, reject) => {
  Aff.runAff(result => () => result instanceof Left ? reject(result.value0) : resolve(result.value0))(action)();
});
const modules = cache => {
  const result = {};
  for (let tail = cache.prepared.modules; tail instanceof Cons; tail = tail.value1) {
    result[tail.value0.name] = Wire.write(Wire.codecModule(Wire.codecAnn))(tail.value0)();
  }
  return result;
};

test("resident preparation retracts and adds specializations after a caller body edit", async t => {
  const workspace = mkdtempSync(join(tmpdir(), "gopurs-prepared-proof-"));
  t.after(() => rmSync(workspace, { recursive: true, force: true }));
  const source = join(workspace, "src"), output = join(workspace, "output");
  mkdirSync(source);
  writeFileSync(join(source, "Generic.purs"), "module Generic (ignore) where\nignore :: forall a. a -> Int\nignore _ = 42\n");
  const compiler = findTypedCompiler(root, process.env.GOPURS_PURS);
  function compile(argument) {
    writeFileSync(join(source, "Caller.purs"), `module Caller (answer) where\nimport Generic (ignore)\nanswer :: Int\nanswer = ignore ${argument}\n`);
    const result = spawnSync(compiler, ["compile", "--codegen", "corefn", "--output", output,
      join(source, "Generic.purs"), join(source, "Caller.purs")], { cwd: workspace, encoding: "utf8", timeout: 30000 });
    assert.ifError(result.error);
    assert.equal(result.status, 0, result.stdout + result.stderr);
  }
  const full = () => runAff(Prepare.prepareModulesWithCache(output)(2)(1)(Nothing.value)(Map.empty));
  const delta = previous => runAff(Prepare.prepareModulesDelta(output)(2)(1)(Nothing.value)(Map.empty)(previous));
  compile("1");
  const initial = await full();
  const untouched = await delta(initial);
  assert.equal(untouched.prepared.modules, initial.prepared.modules, "identical TAST may reuse prepared values");
  compile('"hello"');
  const changed = await delta(initial), expected = await full();
  assert.notEqual(modules(changed).Generic, modules(initial).Generic, "unchanged provider needs a new specialization");
  assert.deepEqual(modules(changed), modules(expected), "resident fallback must equal full preparation");
  compile("1");
  assert.deepEqual(modules(await delta(changed)), modules(await full()), "removed demands must not survive the next edit");
});
