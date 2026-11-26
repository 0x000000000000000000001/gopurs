// Independent source-level behavior checks for private reference loop fusion.
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
  [15, 15, 15, 15, 15, 15, 15, 15], [-12, -12, -12, -12, -12, -12, -12, -12],
  [-2147483648, 1, -2147483648, 1, -2147483648, -2147483648, 1, 1],
  [2147483646, -1, 2147483646, -1, -2147483646, -2147483646, -1, -1],
  [-12, 2, -12, 2, 28, 28, 2, 2],
  [-2147483293, 2147483646, -2147483293, 2147483646, -2147483623, -2147483623, 2147483646, 2147483646],
  [30019, 10000, 30019, 10000, 10019, 10019, 10000, 10000], [19, 19], [26, 1], 23, [17, 18],
  [8, 3], [5, 234], 133, [3, 4],
];
const parse = stdout => stdout.trim().split("\n").map(line => JSON.parse(line));

function emittedWorker(directory, name) {
  for (const entry of readdirSync(directory, { withFileTypes: true })) {
    if (entry.name === ".gopurs-cache") continue;
    const path = join(directory, entry.name);
    const source = entry.isDirectory() ? emittedWorker(path, name)
      : entry.name.endsWith(".go") ? readFileSync(path, "utf8") : "";
    const start = source.indexOf(`func ${name}(`);
    if (start >= 0) {
      const end = source.indexOf("\nfunc ", start + 1);
      return source.slice(start, end < 0 ? undefined : end);
    }
  }
  return "";
}

test("private reference loops preserve runtime bounds, overflow, replay, escaping aliases and intermediate effects", async t => {
  assert.ok(modes.length > 0);
  const workspace = createWorkspace();
  let completed = 0;
  console.log(`Reference loop fixture workspace: ${workspace}`);
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
      const { directory } = prepareFixture(root, workspace, join(root, "tools/fixtures/ref-loop-regions.purs"), index, packages);
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
      const oracle = parse(run("javascript", "node", ["--input-type=module", "-e", 'import { main } from "./output/Main/index.js"; main();']));
      assert.deepEqual(oracle, expected);
      if (process.env.GOPURS_TEST_FRONTEND_ONLY !== "1") {
        for (const profile of ["release", "dev"]) {
          run(profile, process.env.GOPURS_TEST_COMPILER ?? join(root, "bin/gopurs"), [`--${profile}`, "--main", "Main"]);
          // Dev keeps additional library wrappers (for example go__new) and
          // may conservatively retain the loop. Both profiles must agree on
          // behavior; the measured Release profile must actually specialize.
          if (profile === "release") {
            for (const name of ["effectPermuted", "stPermuted"]) {
              const worker = emittedWorker(join(directory, "output"), `Call_Main_${name}`);
              assert.ok(worker, `${profile}: ${name} worker must exist`);
              assert.doesNotMatch(worker, /Call_Main_bump(?:Effect|ST)|NewSTRef|ReadSTRef|WriteSTRef|Get_Effect_Ref_|Get_Control_Monad_ST_Internal_/,
                `${profile}: renamed constant-increment private helpers must really be fused`);
            }
            for (const name of ["effectWritePermuted", "stWritePermuted"]) {
              const worker = emittedWorker(join(directory, "output"), `Call_Main_${name}`);
              assert.match(worker, /2147483647/, `${profile}: inclusive maxBound must retain its guard`);
              assert.match(worker, /Call_Main_write(?:Effect|ST)/, `${profile}: the nonterminating maxBound path must retain the original helper`);
            }
          }
          const actual = parse(run(`${profile}-go`, "go", ["run", "./main"], join(directory, "output")));
          assert.deepEqual(actual, oracle, `${profile}: same values, deferred effects and alias observations as JavaScript`);
        }
      }
      completed++;
    });
  }
});
