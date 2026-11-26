// Independent source-level checks for guarded groups of immutable record reads.
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
const cases = [[0, 19, 1], [2147483647, -2147483648, 1], [-7, 9, -3],
  [-2147483648, 2147483647, -1], [2147483639, -11, 31], [-19, 3, -2147483648]];
const sum = values => { let total = 0; for (const value of values) total = (total + value) | 0; return total; };
const expected = [];
for (const [left, right, delta] of cases) {
  const l = Array.from({ length: 8 }, (_, i) => (left + i) | 0), r = Array.from({ length: 8 }, (_, i) => (right + i) | 0);
  let repeated = (l[7] + l[0]) | 0;
  repeated = (repeated + l[7]) | 0; repeated = (repeated - l[0]) | 0;
  repeated = (repeated + l[1]) | 0; repeated = (repeated - l[1]) | 0;
  const updated = [...l]; updated[7] = (updated[7] + delta) | 0;
  expected.push([sum(l), repeated, (l[0] - l[7]) | 0, sum([...l, ...r]),
    (Math.imul(sum(l), 3) + delta) | 0, sum([...updated, ...l]), sum(l)]);
  expected.push(l[0] < l[7] && l[1] === l[1], [l, r]);
}
expected.push(73, 68, [1, 2, 3, 4, 5, 6, 7, 8, 9, 10], 68, [], true, [5]);
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

test("grouped record reads preserve values, immutable sources, call order and selected exceptions", async t => {
  assert.ok(modes.length > 0);
  const workspace = createWorkspace();
  let completed = 0;
  console.log(`Record read groups fixture workspace: ${workspace}`);
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
      const { directory } = prepareFixture(root, workspace, join(root, "tools/fixtures/record-read-groups.purs"), index, packages);
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
          const workers = Object.fromEntries(["readSum", "readUpdated", "readArithmetic"].map(name => [name,
            emittedWorker(join(directory, "output"), `Get_Main_${name}`) +
            emittedWorker(join(directory, "output"), `Call_Main_${name}`)]));
          for (const [name, worker] of Object.entries(workers)) {
            assert.match(worker, /recordReadGroup_data/, `${profile}: ${name} must exercise grouped record reads`);
          }
          for (const name of ["withUnknown", "withFactories"]) {
            workers[name] = emittedWorker(join(directory, "output"), `Get_Main_${name}`) +
              emittedWorker(join(directory, "output"), `Call_Main_${name}`);
            assert.ok(workers[name]);
            assert.doesNotMatch(workers[name], /recordReadGroup_data/, `${profile}: calls must retain their original evaluation boundary`);
          }
          writeFileSync(join(directory, "logs", `${profile}-workers.json`), JSON.stringify(workers, null, 2) + "\n");
          const actual = parse(run(`${profile}-go`, "go", ["run", "./main"], join(directory, "output")));
          assert.deepEqual(actual, oracle, `${profile}: repeated/partial reads, multiple records, Int32 arithmetic, immutable sources, call order and branch exceptions agree with JavaScript`);
        }
      }
      completed++;
    });
  }
});
