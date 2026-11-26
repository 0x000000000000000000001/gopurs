// Independent source-level checks for guarded native record projection.
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
const cases = [[0, 19], [1, -7], [2, 2147483647], [7, -2147483648], [31, 2147483630], [1000, -19]];
function updated(values, count) {
  const result = [...values];
  for (let step = 0; step < count; step++) result[result.length - 1] = (result[result.length - 1] + 1) | 0;
  return result;
}
const expected = [];
for (const [count, seed] of cases) {
  const wide = Array.from({ length: 16 }, (_, index) => (seed + index) | 0), eight = wide.slice(0, 8), pair = wide.slice(0, 2);
  expected.push([wide, updated(wide, count), updated(wide, count), wide]);
  expected.push([eight, updated(eight, count), eight]);
  expected.push([pair, updated(pair, count), pair]);
  expected.push("changed", [pair[0], pair[1], pair[1]]);
}
expected.push([17, 20, 29, 30, 77, 77], [[4, 7, 11], [4, 7, 11]], [true, true], "keep λ", "keep λ");
const observed = Array.from({ length: 16 }, (_, index) => index + 5), changed = [...observed];
changed[0]--; changed[15]++;
expected.push([observed, changed, observed], [0, 1, 221]);
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

test("record projection preserves complete records, aliases, Int32 results and evaluation order", async t => {
  assert.ok(modes.length > 0);
  const workspace = createWorkspace();
  let completed = 0;
  console.log(`Record projection fixture workspace: ${workspace}`);
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
      const { directory } = prepareFixture(root, workspace, join(root, "tools/fixtures/record-projection-guard.purs"), index, packages);
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
      assert.deepEqual(oracle, expected, "original JavaScript records agree with the independent imperative reference");
      if (process.env.GOPURS_TEST_FRONTEND_ONLY !== "1") {
        for (const profile of ["release", "dev"]) {
          run(profile, process.env.GOPURS_TEST_COMPILER ?? join(root, "bin/gopurs"), [`--${profile}`, "--main", "Main"]);
          const workers = Object.fromEntries(["inspectWide", "inspectEight"].map(name => [name,
            emittedWorker(join(directory, "output"), `Get_Main_${name}`) +
            emittedWorker(join(directory, "output"), `Call_Main_${name}`)]));
          for (const [name, worker] of Object.entries(workers)) {
            assert.match(worker, /recordProjection_data/, `${profile}: ${name} must exercise guarded boxed-to-native projection`);
          }
          writeFileSync(join(directory, "logs", `${profile}-workers.json`), JSON.stringify(workers, null, 2) + "\n");
          const actual = parse(run(`${profile}-go`, "go", ["run", "./main"], join(directory, "output")));
          assert.deepEqual(actual, oracle, `${profile}: complete records, overflow, immutable sources, shared references and evaluation order agree with JavaScript`);
        }
      }
      completed++;
    });
  }
});
