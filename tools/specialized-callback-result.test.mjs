import assert from "node:assert/strict";
import { spawnSync } from "node:child_process";
import { mkdirSync, rmSync, symlinkSync, writeFileSync } from "node:fs";
import { delimiter, join } from "node:path";
import { fileURLToPath } from "node:url";
import test from "node:test";
import { corePackages, createWorkspace, prepareFixture } from "./test-workspace.mjs";
import { findTypedCompiler } from "./build/native-workspace.mjs";

const root = fileURLToPath(new URL("../", import.meta.url));
const modes = [{ name: "js", GOPURS_JS: "1" }, { name: "native", GOPURS_JS: "0" }]
  .filter(mode => !process.env.GOPURS_TEST_MODES || process.env.GOPURS_TEST_MODES.split(",").includes(mode.name));

test("specialized higher-order callbacks preserve Maybe record results", async t => {
  assert.ok(modes.length > 0);
  const workspace = createWorkspace();
  let passed = false;
  t.after(() => {
    if (passed && process.env.GOPURS_TEST_KEEP_WORKSPACE !== "1") rmSync(workspace, { recursive: true, force: true });
    else console.log(`Specialized callback workspace: ${workspace}`);
  });
  const typedBin = join(workspace, "typed-bin");
  mkdirSync(typedBin);
  symlinkSync(findTypedCompiler(root, process.env.GOPURS_PURS), join(typedBin, "purs"));
  const env = {
    ...process.env,
    PATH: [typedBin, join(root, "node_modules/.bin"), process.env.PATH].join(delimiter),
    GOWORK: "off", GOPURS_DAEMON: "0", GOPURS_RESIDENT: "0", GOPURS_RUST: "0",
    GOPURS_FORCE_BUILD: "1", GOPURS_GO_PACKAGES: "0", GOPURS_GO_PACKAGE_OVERLAY: "0", GOPURS_BUILD_MANIFEST: "1",
  };
  for (const key of ["GOPURS_ALLOC_PROFILE", "GOPURS_COMPILER_SNAPSHOT", "PPROF"]) delete env[key];
  const { directory } = prepareFixture(root, workspace,
    join(root, "tools/fixtures/specialized-callback-result.purs"), 0, corePackages(root));
  let sequence = 0;
  const run = (label, command, args, cwd = directory, extra = {}) => {
    const result = spawnSync(command, args, {
      cwd, env: { ...env, ...extra }, encoding: "utf8", timeout: 120_000, maxBuffer: 16 * 1024 * 1024,
    });
    writeFileSync(join(directory, "logs", `${sequence++}-${label}.log`), (result.stdout ?? "") + (result.stderr ?? ""));
    assert.ifError(result.error);
    assert.equal(result.status, 0, `${label}\n${result.stdout}\n${result.stderr}`);
    return result.stdout;
  };
  run("frontend", "spago", ["build", "-q", "--offline"]);
  const reference = run("reference", "node", ["--input-type=module", "-e", "import { main } from './output/Main/index.js'; main();"]);
  assert.equal(reference, "0\n3\n1\n4\n7\n");
  for (const mode of modes) await t.test(mode.name, () => {
    const compiler = process.env.GOPURS_TEST_COMPILER ?? join(root, "bin/gopurs");
    const command = /\.m?js$/.test(compiler) ? process.execPath : compiler;
    const prefix = command === compiler ? [] : ["--stack-size=65536", compiler];
    run(mode.name, command, [...prefix, "--main", "Main"], directory, { GOPURS_JS: mode.GOPURS_JS });
    assert.equal(run(`${mode.name}-go`, "go", ["run", "./main"], join(directory, "output")), reference);
  });
  passed = true;
});
