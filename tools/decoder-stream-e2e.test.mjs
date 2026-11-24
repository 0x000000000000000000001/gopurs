import assert from "node:assert/strict";
import { spawnSync } from "node:child_process";
import { cpSync, existsSync, mkdirSync, readFileSync, readdirSync, rmSync, symlinkSync, writeFileSync } from "node:fs";
import { delimiter, join, resolve } from "node:path";
import { fileURLToPath } from "node:url";
import test from "node:test";
import { corePackages, createWorkspace, prepareFixture } from "./test-workspace.mjs";
import { findTypedCompiler } from "./build/native-workspace.mjs";

const root = fileURLToPath(new URL("../", import.meta.url));
const port = resolve(root, "../gopurs-argonaut-codecs");
const fixture = join(port, "test/typed-plans.purs");
const modes = [{ name: "js", GOPURS_JS: "1" }, { name: "native-parallel", GOPURS_JS: "0" }]
  .filter(mode => !process.env.GOPURS_TEST_MODES || process.env.GOPURS_TEST_MODES.split(",").includes(mode.name));
const streamDeclaration = /func Main___json_schema_\d+_decode_stream\b/;

test("stream decoders preserve typed results, errors, opaque fallbacks and legacy ABI", async t => {
  assert.ok(modes.length > 0);
  const workspace = createWorkspace();
  console.log(`Decoder stream fixture workspace: ${workspace}`);
  let complete = false;
  t.after(() => {
    if (complete && process.env.GOPURS_TEST_KEEP_WORKSPACE !== "1") rmSync(workspace, { recursive: true, force: true });
  });
  const typedBin = join(workspace, "typed-bin");
  mkdirSync(typedBin);
  symlinkSync(findTypedCompiler(root, process.env.GOPURS_PURS), join(typedBin, "purs"));
  const env = {
    ...process.env,
    PATH: [typedBin, join(root, "node_modules/.bin"), process.env.PATH].join(delimiter),
    GOWORK: "off", GOPURS_DAEMON: "0", GOPURS_RESIDENT: "0", GOPURS_RUST: "0",
    GOPURS_FORCE_BUILD: "1", GOPURS_GO_PACKAGES: "1", GOPURS_GO_PACKAGE_OVERLAY: "1",
    GOPURS_PREPARE_JOBS: "2", GOPURS_PBO_JOBS: "2", GOPURS_EMIT_JOBS: "2", GOPURS_PIPELINE: "1",
  };
  for (const key of ["GOPURS_ALLOC_PROFILE", "GOPURS_COMPILER_SNAPSHOT", "PPROF"]) delete env[key];
  const packages = corePackages(root);
  let sequence = 0;
  const emittedGo = output => readdirSync(output, { recursive: true })
    .filter(name => name.endsWith(".go") && !name.includes(".gopurs-cache"))
    .map(name => readFileSync(join(output, name), "utf8")).join("\n");

  async function check(mode, compiler, profiles, expectedStream, legacyPort) {
    const { directory } = prepareFixture(root, workspace, fixture, sequence++, packages);
    if (legacyPort) {
      const copy = join(workspace, ".spago", `legacy-argonaut-codecs-${sequence}`);
      mkdirSync(join(workspace, ".spago"), { recursive: true });
      cpSync(port, copy, { recursive: true, filter: path => !/(?:^|\/)(?:\.git|\.spago|output|node_modules)(?:\/|$)/.test(path) });
      for (const extension of ["purs", "js", "go"]) {
        const relative = `src/Data/Argonaut/Decode/Parser.${extension}`;
        assert.ok(existsSync(join(legacyPort, relative)), `Missing legacy parser: ${relative}`);
        cpSync(join(legacyPort, relative), join(copy, relative));
      }
      const config = join(directory, "spago.yaml");
      const original = readFileSync(config, "utf8");
      assert.ok(original.includes(JSON.stringify(port)));
      writeFileSync(config, original.replace(JSON.stringify(port), JSON.stringify(copy)));
    }
    let step = 0;
    const run = (label, command, args, cwd = directory) => {
      console.log(`${mode.name}: ${label}`);
      const result = spawnSync(command, args, { cwd, env: { ...env, GOPURS_JS: mode.GOPURS_JS }, encoding: "utf8", timeout: 180000, maxBuffer: 32 * 1024 * 1024 });
      writeFileSync(join(directory, "logs", `${step++}-${label}.log`), (result.stdout ?? "") + (result.stderr ?? ""));
      assert.ifError(result.error);
      assert.equal(result.status, 0, `${label}\n${result.stdout}\n${result.stderr}`);
      return result.stdout.trim();
    };
    run("frontend", "spago", ["build", "-q", "--offline"]);
    const oracle = run("javascript", "node", ["--input-type=module", "-e", 'import { main } from "./output/Main/index.js"; main();']);
    const output = join(directory, "output");
    for (const profile of profiles) {
      run(profile, compiler, [`--${profile}`, "--main", "Main"]);
      const emitted = emittedGo(output);
      if (expectedStream && profile === "release") assert.match(emitted, streamDeclaration);
      if (!expectedStream) assert.doesNotMatch(emitted, streamDeclaration);
      assert.equal(run(`${profile}-go`, "go", ["run", "./main"], output), oracle);
    }
  }

  for (const mode of modes) await t.test(mode.name, () => check(mode, join(root, "bin/gopurs"), ["release", "dev"], true));

  const oldCompiler = process.env.GOPURS_STREAM_OLD_COMPILER;
  const oldPort = process.env.GOPURS_STREAM_OLD_PORT;
  if (oldCompiler) await t.test("old compiler with current port", () => check({ name: "old-compiler", GOPURS_JS: "0" }, oldCompiler, ["release"], false));
  if (oldPort) await t.test("current compiler with old port", () => check({ name: "old-port", GOPURS_JS: "0" }, join(root, "bin/gopurs"), ["release"], false, oldPort));
  complete = true;
});
