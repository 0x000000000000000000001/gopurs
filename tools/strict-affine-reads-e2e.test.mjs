// Source oracle across both compiler hosts, including an uncalled FFI function.
import assert from "node:assert/strict";
import { spawnSync } from "node:child_process";
import { mkdirSync, readFileSync, readdirSync, rmSync, symlinkSync, writeFileSync } from "node:fs";
import { delimiter, join } from "node:path";
import { fileURLToPath } from "node:url";
import test from "node:test";
import { corePackages, createWorkspace, prepareFixture } from "./test-workspace.mjs";
import { findTypedCompiler } from "./build/native-workspace.mjs";

const root = fileURLToPath(new URL("../", import.meta.url));
const modes = [
  { name: "js", GOPURS_JS: "1" },
  { name: "native-parallel", GOPURS_JS: "0" },
].filter(mode => !process.env.GOPURS_TEST_MODES || process.env.GOPURS_TEST_MODES.split(",").includes(mode.name));
const expected = [
  [2147483647, 7, 140, -5, -2147483635],
  [223, -14, -2147483627],
  [2147483647, 102, -2147483639],
  [2147483647, 63, -16, -2147483586],
  [24, -17],
];
const parse = stdout => stdout.trim().split("\n").map(line => JSON.parse(line));

function assertFastPaths(directory) {
  const generated = join(directory, "output/purescript");
  const sources = readdirSync(generated, { recursive: true }).filter(name => name.endsWith(".go"))
    .map(name => readFileSync(join(generated, name), "utf8"));
  for (const [name, increment, getter] of [
    ["rounds", 7, "Get_Data_Unit_unit"],
    ["localRounds", 11, "Get_Data_Unit_unit"],
    ["effectRounds", 5, "Get_Effect_Console_log"],
  ]) {
    const declaration = `func Call_Main_${name}(`;
    const source = sources.find(value => value.includes(declaration));
    assert.ok(source, `${name}: generated worker exists`);
    const start = source.indexOf("{", source.indexOf(declaration));
    let end = start + 1, depth = 1;
    for (; depth > 0 && end < source.length; end++) {
      if (source[end] === "{") depth++;
      if (source[end] === "}") depth--;
    }
    const body = source.slice(start, end);
    const guard = body.indexOf("> (int64(0))"), read = body.indexOf(`${getter}()`);
    const formula = body.indexOf(`* (int64(${increment}))`);
    assert.ok(guard >= 0 && read > guard && formula > read,
      `${name}: a positive guarded formula retains its strict global demand`);
    assert.ok(body.includes("int64(2147483647)") && body.includes("for {"), `${name}: bounded path and original loop`);
  }
  for (const [name, divisor] of [["moduloRounds",7],["localModuloRounds",3]]) {
    const declaration=`func Call_Main_${name}(`;
    const source=sources.find(value=>value.includes(declaration));
    assert.ok(source,`${name}: generated worker exists`);
    const start=source.indexOf("{",source.indexOf(declaration));let end=start+1,depth=1;
    for(;depth>0&&end<source.length;end++){if(source[end]==="{")depth++;if(source[end]==="}")depth--;}
    const body=source.slice(start,end);
    assert.ok(body.includes(`% (int64(${divisor}))`)&&body.includes(`/ (int64(${divisor}))`),`${name}: complete remainder periods`);
    assert.ok(body.includes("int64(2147483647)")&&body.includes("for {"),`${name}: guarded fast path and original loop`);
  }
}

test("strict affine reads preserve global and local recurrences, zero, overflow and uncalled effects", async t => {
  assert.ok(modes.length > 0);
  const workspace = createWorkspace();
  let completed = 0;
  console.log(`Strict affine reads fixture workspace: ${workspace}`);
  t.after(() => {
    if (completed === modes.length && process.env.GOPURS_TEST_KEEP_WORKSPACE !== "1") {
      rmSync(workspace, { recursive: true, force: true });
    }
  });
  const typedBin = join(workspace, "typed-bin");
  mkdirSync(typedBin);
  symlinkSync(findTypedCompiler(root, process.env.GOPURS_PURS), join(typedBin, "purs"));
  const env = {
    ...process.env, PATH: [typedBin, join(root, "node_modules/.bin"), process.env.PATH].join(delimiter),
    GOWORK: "off", GOPURS_DAEMON: "0", GOPURS_RESIDENT: "0", GOPURS_RUST: "0",
    GOPURS_GO_PACKAGES: "1", GOPURS_GO_PACKAGE_OVERLAY: "1", GOPURS_BUILD_MANIFEST: "1",
    GOPURS_PREPARE_JOBS: "2", GOPURS_PBO_JOBS: "2", GOPURS_EMIT_JOBS: "2",
    GOPURS_PIPELINE: "1", GOPURS_FORCE_BUILD: "1",
  };
  for (const key of ["GOPURS_ALLOC_PROFILE", "GOPURS_COMPILER_SNAPSHOT", "PPROF"]) delete env[key];
  const packages = corePackages(root);
  for (const [index, mode] of modes.entries()) {
    await t.test(mode.name, () => {
      const { directory } = prepareFixture(root, workspace, join(root, "tools/fixtures/strict-affine-reads.purs"), index, packages);
      let sequence = 0;
      const run = (label, command, args, cwd = directory) => {
        console.log(`${mode.name}: ${label}`);
        const result = spawnSync(command, args, {
          cwd, env: { ...env, GOPURS_JS: mode.GOPURS_JS }, encoding: "utf8", timeout: 120000,
          maxBuffer: 16 * 1024 * 1024,
        });
        writeFileSync(join(directory, "logs", `${sequence++}-${label}.log`), (result.stdout ?? "") + (result.stderr ?? ""));
        assert.ifError(result.error);
        assert.equal(result.status, 0, `${label}\n${result.stdout}\n${result.stderr}`);
        return result.stdout;
      };
      run("frontend", "spago", ["build", "-q", "--offline"]);
      assert.deepEqual(parse(run("javascript", "node", ["--input-type=module", "-e", 'import { main } from "./output/Main/index.js"; main();'])), expected);
      for (const profile of ["release", "dev"]) {
        run(profile, process.env.GOPURS_TEST_COMPILER ?? join(root, "bin/gopurs"), [`--${profile}`, "--main", "Main"]);
        assertFastPaths(directory);
        const actual = parse(run(`${profile}-go`, "go", ["run", "./main"], join(directory, "output")));
        assert.deepEqual(actual, expected, `${profile}: exact integers and no invocation of the retained FFI function`);
      }
      completed++;
    });
  }
});
