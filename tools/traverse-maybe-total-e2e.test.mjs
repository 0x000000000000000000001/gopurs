// Source-level checks independent of the compiler's totality recognizer.
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
  [[], 99, 1], [[1], 1, 1], [[1], 99, 1],
  [[1, 2, 3, 4], 1, 7], [[1, 2, 3, 4], 2, -3], [[1, 2, 3, 4], 4, 1], [[1, 2, 3, 4], 99, 1],
  [[-3, 0, -3, 5], -3, -7], [[2147483647, -2147483648, -1, 0], 19, 1],
  [[-2147483648, 2147483647], 19, -1], [[-7, 19, 8], 8, 2147483647],
  [[0, 2, 4, 6, 8, 10, 12], 6, -2147483648],
];
function reference(values, rejected, delta, indexed = false, fail = true) {
  const result = [1];
  let failed = false;
  for (let index = 0; index < values.length; index++) {
    const value = values[index];
    if (fail && (indexed ? index : value) === rejected) failed = true;
    result.push((((value + delta) | 0) + (indexed ? index : 0)) | 0);
  }
  return failed ? [0] : result;
}
function branchReference(values) {
  const result = [1];
  let failed = false;
  for (const value of values) {
    if (value < 0) result.push((-value) | 0);
    else if (value === 2) failed = true;
    else result.push((Math.imul(value, 3) - 7) | 0);
  }
  return failed ? [0] : result;
}
const expected = cases.map(([values, rejected, delta]) => {
  const selected = reference(values, rejected, delta);
  return [values, selected, selected, selected, reference(values, rejected, delta, false, false),
    reference(values, rejected, delta, true), [1, ...values.map(() => 7)], values.length ? [0] : [1],
    rejected === 1 && values.length ? [0] : reference(values, rejected, 1, false, false), branchReference(values), selected, values];
});
for (const rejected of [1, 2, 4, 99]) {
  expected.push(reference([1, 2, 3, 4], rejected, 1));
  expected.push([[1, rejected], [2, 4], [3, 1], [3, 2], [3, 3], [3, 4]]);
}
expected.push(true, [[4, 1], [4, 2], [4, 3]]);
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

test("total Maybe traversal preserves arbitrary arrays, captured values, Int32 results and strict fallback callbacks", async t => {
  assert.ok(modes.length > 0);
  const workspace = createWorkspace();
  let completed = 0;
  console.log(`Total Maybe fixture workspace: ${workspace}`);
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
      const { directory } = prepareFixture(root, workspace, join(root, "tools/fixtures/traverse-maybe-total.purs"), index, packages);
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
      assert.deepEqual(oracle, expected, "original JavaScript traversal agrees with independent imperative reference and strict callback order");
      if (process.env.GOPURS_TEST_FRONTEND_ONLY !== "1") {
        for (const profile of ["release", "dev"]) {
          run(profile, process.env.GOPURS_TEST_COMPILER ?? join(root, "bin/gopurs"), [`--${profile}`, "--main", "Main"]);
          const workers = Object.fromEntries(["constantSuccess", "constantFailure", "capturedFlag", "branched"].map(name =>
            [name, emittedWorker(join(directory, "output"), `Call_Main_${name}`)]));
          for (const [name, worker] of Object.entries(workers)) {
            assert.match(worker, /traverseTotal_values_/, `${profile}: ${name} must exercise the total path`);
          }
          writeFileSync(join(directory, "logs", `${profile}-workers.json`), JSON.stringify(workers, null, 2) + "\n");
          const actual = parse(run(`${profile}-go`, "go", ["run", "./main"], join(directory, "output")));
          assert.deepEqual(actual, oracle, `${profile}: complete arrays, overflow, immutable inputs, evaluation order and late exception agree with JavaScript`);
        }
      }
      completed++;
    });
  }
});
