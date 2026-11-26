// Source-level checks independent of the compiler's read-only callback proof.
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
const cases = [
  [[], 19, 7, true], [[1], -7, 1, false], [[1, 2, 3, 4], 0, -3, true],
  [[-3, 0, -3, 5], -19, -7, false], [[2147483647, -2147483648, -1, 0], 2147483647, 1, true],
  [[-2147483648, 2147483647], -2147483648, -1, false], [[-7, 19, 8], 17, 2147483647, true],
  [[0, 2, 4, 6, 8, 10, 12], -19, -2147483648, true],
];
function fold(values, seed, callback) {
  let total = seed;
  for (let index = 0; index < values.length; index++) total = callback(index, total, values[index]) | 0;
  return total;
}
const weighted = (index, total, value) => (total + Math.imul(1 + index % 2, value)) | 0;
const ordered = (index, total, value) => ((Math.imul(total, 31) + Math.imul(index, 7)) | 0) + value;
const expected = cases.map(([values, seed, delta, enabled]) => {
  const sum = fold(values, seed, (_index, total, value) => total + value);
  return [values, [
    fold(values, seed, weighted), fold(values, seed, ordered),
    fold(values, seed, (index, total, value) => enabled ? (((total + value) | 0) + delta | 0) + index : total - value),
    fold(values, seed, (index, total) => total + index), values.length ? 7 : seed,
    fold(values, seed, ordered), fold(values, seed, weighted), seed,
    fold(values, seed, (index, total, value) => ((total + index) | 0) + value),
    fold(values, 23, (index, total, value) => ((Math.imul(total, 7) + value) | 0) + index), 23, seed,
  ], [sum, ...values.map((_value, index) => index)], values];
});
expected.push("seed:", "seed:0a12bc");
const arrays = [[1, 2], [], [3, -1, 4]], stagedTrace = [[1], [2, 5]];
for (const values of arrays) {
  stagedTrace.push([3, values.length]);
  expected.push(fold(values, 5, (index, total, value) => {
    stagedTrace.push([4, index], [5, total], [6, value]);
    return Math.imul(total, 10) + index + value;
  }));
}
expected.push(stagedTrace);
const inlineTrace = [];
for (const values of arrays) expected.push(fold(values, 7, (index, total, value) => {
  inlineTrace.push([7, index, total, value]);
  return total + index + value;
}));
expected.push(inlineTrace, 20, [[9, 0], [9, 1], [9, 2]], true, [[8, 0, 9, 1], [8, 1, 10, 2], [8, 2, 12, 3]]);
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

test("indexed array fold preserves argument stages, runtime seeds, indices, Int32 results and strict callbacks", async t => {
  assert.ok(modes.length > 0);
  const workspace = createWorkspace();
  let completed = 0;
  console.log(`Indexed fold fixture workspace: ${workspace}`);
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
      const { directory } = prepareFixture(root, workspace, join(root, "tools/fixtures/indexed-array-fold.purs"), index, packages);
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
      assert.deepEqual(oracle, expected, "original JavaScript fold agrees with independent imperative reference and strict callback order");
      if (process.env.GOPURS_TEST_FRONTEND_ONLY !== "1") {
        for (const profile of ["release", "dev"]) {
          run(profile, process.env.GOPURS_TEST_COMPILER ?? join(root, "bin/gopurs"), [`--${profile}`, "--main", "Main"]);
          const workers = Object.fromEntries(["weighted", "ordered", "indices", "divisionZero"].map(name =>
            [name, emittedWorker(join(directory, "output"), `Call_Main_${name}`)]));
          for (const [name, worker] of Object.entries(workers)) {
            assert.match(worker, /indexedFold_values_/, `${profile}: ${name} must exercise the indexed fold path`);
          }
          // This partial dictionary alias and callbacks whose unused arguments
          // lack an exact TAST type may conservatively retain their FFI path.
          workers.fixedPartial = emittedWorker(join(directory, "output"), "Get_Main_fixedPartial");
          workers.constant = emittedWorker(join(directory, "output"), "Call_Main_constant");
          workers.captured = emittedWorker(join(directory, "output"), "Call_Main_captured");
          workers.runUnknown = emittedWorker(join(directory, "output"), "Call_Main_runUnknown");
          assert.ok(workers.runUnknown, `${profile}: unknown callback worker must exist`);
          assert.doesNotMatch(workers.runUnknown, /indexedFold_values_/, `${profile}: unknown callback must retain the copying FFI path`);
          writeFileSync(join(directory, "logs", `${profile}-workers.json`), JSON.stringify(workers, null, 2) + "\n");
          const actual = parse(run(`${profile}-go`, "go", ["run", "./main"], join(directory, "output")));
          assert.deepEqual(actual, oracle, `${profile}: runtime seeds, indices, overflow, immutable inputs, evaluation stages and late exception agree with JavaScript`);
        }
      }
      completed++;
    });
  }
});
