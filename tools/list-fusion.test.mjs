// End-to-end semantics complement the pass's structural admission tests.
import assert from "node:assert/strict";
import { spawnSync } from "node:child_process";
import { mkdirSync, readFileSync, rmSync, symlinkSync, writeFileSync } from "node:fs";
import { delimiter, join } from "node:path";
import { fileURLToPath } from "node:url";
import test from "node:test";
import { corePackages, createWorkspace, prepareFixture } from "./test-workspace.mjs";
import { findTypedCompiler } from "./build/native-workspace.mjs";

const root = fileURLToPath(new URL("../", import.meta.url));
const modes = [
  { name: "js", GOPURS_JS: "1" },
  { name: "native-parallel", GOPURS_JS: "0" },
].filter((mode) => !process.env.GOPURS_TEST_MODES || process.env.GOPURS_TEST_MODES.split(",").includes(mode.name));
const cases = [
  [4, 3, 123],
  [2, 2, 0],
  [3, 3, 7],
  [0, 0, -7],
  [-7, -1, 0],
  [-8, 8, 17],
  [1, 9, 2147483647],
  [1, 10, -2147483648],
  [1, 900, 0],
  [-7, -1, -2147483648],
  [1, 900, -17],
];
const nearMaxCases = [
  [2147483630, 17],
  [2147483631, -2147483648],
  [2147483633, 2147483647],
  [2147483640, -2147483648],
  [2147483646, 0],
  [2147483647, 17],
  [2147483647, -2147483648],
];

// Array operations and exact BigInt arithmetic provide an oracle independent
// of both the fixture's recursive lists and the compiler's fusion algorithm.
const int32 = (value) => BigInt.asIntN(32, value);
const sum = (values, seed) => values.reduce((acc, value) => int32(acc + value), BigInt(seed));
const fingerprint = (values, seed) => values.reduce((acc, value) => int32(acc * 31n + value), BigInt(seed));
const range = (lower, upper) =>
  Array.from({ length: Math.max(0, upper - lower + 1) }, (_, index) => BigInt(lower + index));
const expected = cases.map(([lower, upper, seed]) => {
  const source = range(lower, upper);
  const even = source.filter((value) => value % 2n === 0n).reverse();
  const fromOne = range(1, upper).filter((value) => value % 2n === 0n).reverse();
  const odd = range(-8, upper).filter((value) => value % 2n !== 0n).reverse();
  const third = range(-5, upper).filter((value) => value % 3n === 0n).reverse();
  const residue = range(1, upper).filter((value) => value % 5n === 3n).reverse();
  const mapped = even.map((value) => int32(int32(value * value) - int32(3n * value) + 2n));
  return [
    sum(fromOne, seed),
    sum(odd, seed),
    sum(third, seed),
    sum(mapped, seed),
    fingerprint(fromOne, seed),
    sum(even, seed),
    fingerprint(source, 0),
    fingerprint(even, 0),
    sum(source, 0),
    int32(sum(fromOne, seed) + sum(odd, -17)),
    sum(residue, seed),
    sum(residue, seed),
  ].map(Number);
}).concat(nearMaxCases.map(([upper, seed]) => [
  Number(sum(range(2147483631, upper).filter((value) => value % 5n === 3n), seed)),
]));
const parseResults = (stdout) => stdout.trim().split("\n").map((line) => JSON.parse(line));

const allocationProbe = `package listfusionprobe

import (
  "testing"
  p "gopurs/output/purescript/p_root"
  _ "gopurs/output/purescript/init_all"
)

var resultSink int64

func TestReleaseFusionEliminatesListAllocations(t *testing.T) {
  cases := []struct {
    name string
    run func(int64, int64) int64
    lower, modulus, residue int64
    odd bool
  }{
    {"even-from-one", p.Call_Main_score, 1, 2, 0, false},
    {"odd-from-negative", p.Call_Main_oddScore, -8, 2, 0, true},
    {"multiples-of-three", p.Call_Main_thirdScore, -5, 3, 0, false},
    {"nonzero-residue", p.Call_Main_residueScore, 1, 5, 3, false},
    {"scalar-recurrence", p.Call_Main_scalarResidueScore, 1, 5, 3, false},
  }
  for _, c := range cases {
    t.Run(c.name, func(t *testing.T) {
      for _, upper := range []int64{-10, -1, 0, 1, 2, 9, 900} {
        want := int64(17)
        for value := c.lower; value <= upper; value++ {
          selected := value % c.modulus == c.residue
          if c.odd { selected = !selected }
          if selected { want += value }
        }
        allocations := testing.AllocsPerRun(10, func() { resultSink = c.run(upper, 17) })
        if resultSink != want { t.Fatalf("upper=%d: got %d, want %d", upper, resultSink, want) }
        if allocations != 0 {
          t.Fatalf("upper=%d: fusion allocated %g objects; producer/filter list cells must be eliminated", upper, allocations)
        }
      }
    })
  }
}

func TestWideSeedsUseTheOriginalInt32Recurrence(t *testing.T) {
  for _, run := range []func(int64, int64) int64{p.Call_Main_residueScore, p.Call_Main_scalarResidueScore} {
    for _, upper := range []int64{-3, 0, 1, 2, 3, 4, 8} {
      for _, seed := range []int64{-9223372036854775808, -2147483649, 2147483648, 4294967295, 9223372036854775807} {
        want := seed
        for value := upper; value >= 1; value-- {
          if value % 5 == 3 { want = int64(int32(want + value)) }
        }
        if got := run(upper, seed); got != want {
          t.Fatalf("upper=%d seed=%d: got %d, want %d", upper, seed, got, want)
        }
      }
    }
  }
}
`;

test("list producer/filter/fold fusion preserves order, Int32 sums and shared values", async (t) => {
  assert.ok(modes.length > 0);
  const workspace = createWorkspace();
  let completed = 0;
  console.log(`List fusion workspace: ${workspace}`);
  t.after(() => {
    if (completed === modes.length && process.env.GOPURS_TEST_KEEP_WORKSPACE !== "1") {
      rmSync(workspace, { recursive: true, force: true });
    }
  });
  const typedBin = join(workspace, "typed-bin");
  mkdirSync(typedBin);
  symlinkSync(findTypedCompiler(root, process.env.GOPURS_PURS), join(typedBin, "purs"));
  const env = {
    ...process.env,
    PATH: [typedBin, join(root, "node_modules/.bin"), process.env.PATH].join(delimiter),
    GOWORK: "off",
    GOPURS_DAEMON: "0",
    GOPURS_RESIDENT: "0",
    GOPURS_RUST: "0",
    GOPURS_GO_PACKAGES: "1",
    GOPURS_GO_PACKAGE_OVERLAY: "1",
    GOPURS_BUILD_MANIFEST: "1",
    GOPURS_PREPARE_JOBS: "2",
    GOPURS_PBO_JOBS: "2",
    GOPURS_EMIT_JOBS: "2",
    GOPURS_PIPELINE: "1",
    GOPURS_FORCE_BUILD: "1",
  };
  for (const key of ["GOPURS_ALLOC_PROFILE", "GOPURS_COMPILER_SNAPSHOT", "PPROF"]) delete env[key];
  const packages = corePackages(root);
  for (const [index, mode] of modes.entries()) {
    await t.test(mode.name, () => {
      const { directory } = prepareFixture(root, workspace, join(root, "tools/fixtures/list-fusion.purs"), index, packages);
      let sequence = 0;
      const run = (label, command, args, cwd = directory) => {
        console.log(`${mode.name}: ${label}`);
        const result = spawnSync(command, args, {
          cwd,
          env: { ...env, GOPURS_JS: mode.GOPURS_JS },
          encoding: "utf8",
          timeout: 120000,
          maxBuffer: 16 * 1024 * 1024,
        });
        writeFileSync(join(directory, "logs", `${sequence++}-${label}.log`), (result.stdout ?? "") + (result.stderr ?? ""));
        assert.ifError(result.error);
        assert.equal(result.status, 0, `${label}\n${result.stdout}\n${result.stderr}`);
        return result.stdout;
      };
      run("frontend", "spago", ["build", "-q", "--offline"]);
      const reference = run("javascript", "node", [
        "--input-type=module",
        "-e",
        'import { main } from "./output/Main/index.js"; main();',
      ]);
      assert.deepEqual(parseResults(reference), expected, "PureScript semantics match the independent oracle");
      for (const profile of ["release", "dev"]) {
        run(profile, process.env.GOPURS_TEST_COMPILER ?? join(root, "bin/gopurs"), [`--${profile}`, "--main", "Main"]);
        if (profile === "release") {
          const source = readFileSync(join(directory, "output/purescript/p_root/0_gopurs_package.go"), "utf8");
          const admitted = ["residueScore", "scalarResidueScore", "nearMaxScore"].map((name) => {
            const start = source.indexOf(`func Call_Main_${name}(`);
            assert.ok(start >= 0, `${name}: generated worker exists`);
            const end = source.indexOf("\nfunc ", start + 1);
            const worker = source.slice(start, end < 0 ? source.length : end);
            for (const pattern of [/first :=/, /count :=/, /last :=/, /int64\(2147483647\)/, /int64\(int32\(/]) {
              assert.match(worker, pattern, `${name}: guarded arithmetic sum replaces repeated additions`);
            }
            return worker;
          });
          writeFileSync(join(directory, "logs/release-congruence.go"), admitted.join("\n"));
        }
        const generated = run(`${profile}-go`, "go", ["run", "./main"], join(directory, "output"));
        assert.deepEqual(parseResults(generated), expected,
          `${profile}: ${cases.length} ranges with twelve observations and ${nearMaxCases.length} ranges near MaxInt32`);
        if (profile === "release") {
          const output = join(directory, "output");
          mkdirSync(join(output, "list-fusion-probe"));
          writeFileSync(join(output, "list-fusion-probe", "fusion_test.go"), allocationProbe);
          run("release-allocations", "go", ["test", "-count=1", "-v", "./list-fusion-probe"], output);
        }
      }
      completed++;
    });
  }
});
