// Independent producer/consumer phase, ownership and imported-worker checks.
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
  [[], [], 99, 7, 19], [[1], [], 1, 1, -7],
  [[1, 2, 3, 4], [5, 6], 1, -3, 0], [[1, 2, 3, 4], [5], 3, 7, 5],
  [[1, 2, 3, 4], [-2, -3], 4, -1, 2147483647], [[-3, 0, -3, 5], [0], 99, -2147483648, -19],
  [[2147483647, -2147483648, -1, 0], [1, 2147483647], 19, 1, 2147483647],
  [[-2147483648, 2147483647], [], -2147483648, -1, -2147483648],
  [[], [3, -1, 4], 3, 1, 7], [[4, -8, 9], [], 9, 2147483647, -7],
];
function weighted(values, seed = 0) {
  let result = seed;
  for (let index = 0; index < values.length; index++) result = (result + Math.imul(1 + index % 2, values[index])) | 0;
  return result;
}
function expectedFor(consumerSeed = 0, seedRuns = 0) {
  const result = [];
  for (const [left, right, rejected, delta, seed] of cases) {
    const mapped = left.map(value => (value + 1) | 0), appended = [...left, ...right];
    const rejectedInput = left.some(value => value === rejected);
    const mappedScore = weighted(mapped, consumerSeed);
    result.push([mappedScore, weighted(appended, consumerSeed), mappedScore,
      rejectedInput ? -1 : mappedScore, weighted(left.map(value => (value + delta) | 0), seed),
      weighted(appended, seed), mappedScore]);
    const encoded = rejectedInput ? [0] : [1, ...mapped];
    result.push([left, right, mapped, appended, encoded, encoded, mapped, left, right]);
  }
  result.push([weighted([1, 2, 3], consumerSeed), consumerSeed, weighted([4, -1], consumerSeed)], weighted([2, 0, -3], consumerSeed));
  result.push(29, [[11], [12, 7], [10], [13, 4], [1, 1], [1, 2], [1, 3], [1, 4],
    [2, 0, 7, 2], [2, 1, 9, 3], [2, 2, 15, 4], [2, 3, 19, 5]]);
  result.push([12, 7, 10], [[11], [12, 7], [2, 0, 7, 1], [2, 1, 8, 2], [2, 0, 7, 3]]);
  result.push("producer failure", [[3, 1], [3, 2], [3, 3]]);
  result.push("later traversal failure", [[5, 1], [5, 2], [5, 3]]);
  result.push("right input failure", [[14, 4], [15]], seedRuns);
  return result;
}
const expected = expectedFor();
const unknownSeed = `(unsafePerformEffect do
  Ref.modify_ (_ + 1) seedRuns
  cell <- Ref.new 17
  Ref.read cell)`;
function ineligibleConsumer(original) {
  const changed = original.replace(/(consume = foldlWithIndex[^\n]*\)) 0\n/, `$1 ${unknownSeed}\n`);
  assert.notEqual(changed, original, "consumer seed substitution must happen exactly once");
  return changed;
}
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

test("array intermediates preserve complete values, producer phases, alias boundaries and imported worker invalidation", async t => {
  assert.ok(modes.length > 0);
  const workspace = createWorkspace();
  let completed = 0;
  console.log(`Array intermediates fixture workspace: ${workspace}`);
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
      const { directory } = prepareFixture(root, workspace, join(root, "tools/fixtures/array-intermediates.purs"), index, packages);
      let sequence = 0;
      const run = (label, command, args, cwd = directory, overrides = {}) => {
        console.log(`${mode.name}: ${label}`);
        const result = spawnSync(command, args, {
          cwd, env: { ...env, GOPURS_JS: mode.GOPURS_JS, ...overrides }, encoding: "utf8", timeout: 120000,
          maxBuffer: 16 * 1024 * 1024,
        });
        writeFileSync(join(directory, "logs", `${sequence++}-${label}.log`), (result.stdout ?? "") + (result.stderr ?? ""));
        assert.ifError(result.error);
        assert.equal(result.status, 0, `${label}\n${result.stdout}\n${result.stderr}`);
        return result.stdout;
      };
      run("frontend", "spago", ["build", "-q", "--offline"]);
      const oracle = parse(run("javascript", "node", ["--input-type=module", "-e", 'import { main } from "./output/Main/index.js"; main();']));
      assert.deepEqual(oracle, expected, "original JavaScript arrays agree with the independent imperative reference and phase traces");
      if (process.env.GOPURS_TEST_FRONTEND_ONLY !== "1") {
        for (const profile of ["release", "dev"]) {
          run(profile, process.env.GOPURS_TEST_COMPILER ?? join(root, "bin/gopurs"), [`--${profile}`, "--main", "Main"]);
          const workers = Object.fromEntries(["mapScore", "appendScore", "traverseSuccess"].map(name => [name,
            emittedWorker(join(directory, "output"), `Call_Main_${name}`)]));
          workers.consumer = emittedWorker(join(directory, "output"), "Call_Consumer_consume");
          workers.consumerGetter = emittedWorker(join(directory, "output"), "Get_Consumer_consume");
          if (profile === "release") {
            assert.match(workers.consumer, /func Call_Consumer_consume\([^)]*\[\]int64/, "literal-seed consumer must publish a native array worker");
            assert.match(workers.mapScore, /Call_Consumer_consume/, "map must actually call the imported native worker");
            assert.match(workers.mapScore, /Call_Consumer_consume\(nativeArray_arg_\d+\)/, "map output must reach the native worker without boxing and unboxing");
            assert.match(workers.appendScore, /concatInt_output_/, "append must use the native concatenation buffer");
            assert.match(workers.traverseSuccess, /Constructor_Data_Maybe_Just\[\[\]int64\]/, "traverse must retain a native Maybe array payload");
            assert.match(workers.consumerGetter, /indexedFold_values_/, "the public getter must retain its direct boxed-array fold");
          }
          writeFileSync(join(directory, "logs", `${profile}-workers.json`), JSON.stringify(workers, null, 2) + "\n");
          const actual = parse(run(`${profile}-go`, "go", ["run", "./main"], join(directory, "output")));
          assert.deepEqual(actual, oracle, `${profile}: complete arrays, Int32 arithmetic, immutable inputs, producer phases and exception priority agree with JavaScript`);
          if (profile === "release" && process.env.GOPURS_TEST_SKIP_INCREMENTAL !== "1") {
            const consumerPath = join(directory, "src", "Consumer.purs");
            const originalConsumer = readFileSync(consumerPath, "utf8");
            const changedConsumer = ineligibleConsumer(originalConsumer);
            writeFileSync(consumerPath, changedConsumer);
            run("incremental-frontend", "spago", ["build", "-q", "--offline"]);
            const changedOracle = parse(run("incremental-javascript", "node", ["--input-type=module", "-e", 'import { main } from "./output/Main/index.js"; main();']));
            assert.deepEqual(changedOracle, expectedFor(17, 1), "unknown seed is evaluated once and changes every applicable result");
            run("incremental-release", process.env.GOPURS_TEST_COMPILER ?? join(root, "bin/gopurs"), ["--release", "--main", "Main"], directory, { GOPURS_FORCE_BUILD: "0" });
            const changedWorker = emittedWorker(join(directory, "output"), "Call_Consumer_consume");
            assert.doesNotMatch(changedWorker, /func Call_Consumer_consume\([^)]*\[\]int64/, "unknown seed must lose native-worker eligibility");
            const incremental = parse(run("incremental-go", "go", ["run", "./main"], join(directory, "output")));
            assert.deepEqual(incremental, changedOracle, "cached callers must observe the changed consumer ABI and seed");
            const clean = prepareFixture(root, workspace, join(root, "tools/fixtures/array-intermediates.purs"), 100 + index, packages).directory;
            writeFileSync(join(clean, "src", "Consumer.purs"), changedConsumer);
            run("clean-frontend", "spago", ["build", "-q", "--offline"], clean);
            run("clean-release", process.env.GOPURS_TEST_COMPILER ?? join(root, "bin/gopurs"), ["--release", "--main", "Main"], clean);
            const rebuilt = parse(run("clean-go", "go", ["run", "./main"], join(clean, "output")));
            assert.deepEqual(rebuilt, incremental, "incremental consumer ABI change must match a fresh build");
            writeFileSync(join(directory, "logs", "incremental-validation.json"), JSON.stringify({
              transition: "literal seed 0 / native worker -> effectful unknown seed 17 / boxed closure",
              seed_evaluations: 1, original_observations: expected.length, changed_observations: changedOracle.length,
              clean_equals_incremental: true, native_worker_removed: true,
            }, null, 2) + "\n");
            writeFileSync(consumerPath, originalConsumer);
            run("restore-frontend", "spago", ["build", "-q", "--offline"]);
          }
        }
      }
      completed++;
    });
  }
});
