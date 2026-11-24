// Test the public PureScript pipeline, including evaluation order and fallback
// semantics, without depending on the optimizer's generated helper names.
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
  [1, 9, 0, 2], [9, 1, 17, 2], [-4, 4, 0, 3], [4, -4, -17, -3],
  [3, 3, 2147483647, 2], [4, 4, 2147483647, 2], [0, 0, -7, 0],
  [-2147483648, -2147483646, -1, 2], [-2147483646, -2147483648, -2147483648, -3],
  [2147483645, 2147483647, 2147483647, 2], [2147483647, 2147483645, -2147483648, 3],
  [1, 7, 13, 20], [1, 8, -2147483648, 0], [-8, -1, 2147483647, 2], [2, 2, 0, -2147483648],
];
const int32 = (value) => BigInt.asIntN(32, value);
const range = (first, last) => Array.from({ length: Math.abs(last - first) + 1 },
  (_, index) => first + index * (first <= last ? 1 : -1));
const selectedValues = (first, last, divisor) => range(first, last)
  .filter((value) => divisor === 0 || value % divisor === 0);
const selectedResidueValues = (first, last) => range(first, last)
  .filter((value) => ((BigInt(value) % 5n) + 5n) % 5n === 3n);
const fold = (values, seed, multiplier = 1n) => Number(values
  .reduce((acc, value) => int32(acc * multiplier + BigInt(value)), BigInt(seed)));
const expected = cases.flatMap(([first, last, seed, divisor]) => {
  const source = range(first, last);
  const selected = selectedValues(first, last, divisor);
  const sum = fold(selected, seed);
  const ordered = fold(selected, seed, 31n);
  const residueSum = fold(selectedResidueValues(first, last), seed);
  const trace = [[0, seed], [1, divisor], [2, first], [3, last], ...source.map((value) => [4, value])];
  return [
    [source, selected, selected.filter((value) => value > 0), source, selected, [sum, sum, ordered]],
    sum, trace,
    ordered, [...trace, ...selected.map((value) => [5, value])],
    [fold(selectedValues(first, last, 2), seed), residueSum],
    residueSum, [[0, seed], [2, first], [3, last]],
  ];
});
const parseResults = (stdout) => stdout.trim().split("\n").map((line) => JSON.parse(line));
const allocationCases = [...cases, [1, 900, 0, 2], [900, 1, -17, 3]]
  .map(([first, last, seed, divisor]) =>
    `{${first}, ${last}, ${seed}, ${divisor}, ${fold(selectedValues(first, last, divisor), seed)}}`).join(",\n");
const congruenceAllocationCases = [...cases, [1, 900, 0, 2], [900, 1, -17, 3]]
  .map(([first, last, seed]) => `{${first}, ${last}, ${seed}, ` +
    `${fold(selectedValues(first, last, 2), seed)}, ${fold(selectedResidueValues(first, last), seed)}}`).join(",\n");

// Sum the indices k of k*modulus+residue using exact BigInt triangular sums.
// This oracle does not materialize enormous arrays or repeat the emitted
// first/count/last arithmetic in Number precision.
const largeSum = (first, last, seed, modulus, residue) => {
  const lower = BigInt(Math.min(first, last)), upper = BigInt(Math.max(first, last));
  const m = BigInt(modulus), r = BigInt(residue);
  const begin = lower > r ? (lower - r + m - 1n) / m : 0n;
  const end = upper >= r ? (upper - r) / m : -1n;
  if (end < begin) return BigInt(seed);
  const count = end - begin + 1n;
  const indices = (end * (end + 1n) - (begin - 1n) * begin) / 2n;
  return int32(BigInt(seed) + m * indices + r * count);
};
const largeCases = [
  [1, 2147483647], [2147483647, 1], [1, 65536], [65536, 1],
  [1073741823, 2147483647], [2147483647, 2147483645], [2147483647, 2147483647],
].flatMap(([first, last]) => [0, -2147483648, 2147483647].map((seed) =>
  `{${first}, ${last}, ${seed}, ${largeSum(first, last, seed, 2, 0)}, ${largeSum(first, last, seed, 5, 3)}}`)).join(",\n");
const allocationProbe = `package arrayrangeprobe

import (
  "testing"
  "slices"
  p "gopurs/output/purescript/p_root"
  _ "gopurs/output/purescript/init_all"
)

var sumSink int64

func TestReleaseRangeFilterSumEliminatesArrays(t *testing.T) {
  for _, c := range []struct{first,last,seed,divisor,want int64}{${allocationCases}} {
    allocations := testing.AllocsPerRun(10, func() { sumSink = p.Call_Main_sumRange(c.first, c.last, c.seed, c.divisor) })
    if sumSink != c.want { t.Fatalf("%+v: got %d", c, sumSink) }
    if allocations != 0 { t.Fatalf("%+v: %g allocations; range and filtered buffers must be eliminated", c, allocations) }
  }
}

func TestConstantCongruencesKeepZeroAllocations(t *testing.T) {
  for _, c := range []struct{first,last,seed,even,residue int64}{${congruenceAllocationCases}} {
    for _, calculation := range []struct{run func(int64,int64,int64)int64; want int64}{
      {p.Call_Main_sumEven,c.even}, {p.Call_Main_sumResidue,c.residue},
    } {
      allocations := testing.AllocsPerRun(10, func() { sumSink = calculation.run(c.first,c.last,c.seed) })
      if sumSink != calculation.want { t.Fatalf("%+v: got %d want %d",c,sumSink,calculation.want) }
      if allocations != 0 { t.Fatalf("%+v: %g allocations",c,allocations) }
    }
  }
}

func TestConstantCongruencesHandleLargePositiveRanges(t *testing.T) {
  for _, c := range []struct{first,last,seed,even,residue int64}{${largeCases}} {
    if got := p.Call_Main_sumEven(c.first,c.last,c.seed); got != c.even { t.Fatalf("even %+v: got %d",c,got) }
    if got := p.Call_Main_sumResidue(c.first,c.last,c.seed); got != c.residue { t.Fatalf("residue %+v: got %d",c,got) }
  }
}

func TestNoncanonicalInt64UsesTheMaterializedSemantics(t *testing.T) {
  for _, bounds := range [][2]int64{
    {1,3}, {3,1}, {1,1},
    {2147483646,2147483648}, {2147483648,2147483650},
    {-2147483650,-2147483648}, {4294967295,4294967297},
    {2147483650,2147483648}, {-2147483648,-2147483650},
    {4294967297,4294967295},
  } {
    for _, seed := range []int64{0,4294967295,-2147483649} {
      for _, divisor := range []int64{0,2,3} {
        source := p.Call_Main_makeRange(bounds[0], bounds[1])
        original := slices.Clone(source)
        want := p.Call_Main_sumBorrowed(seed, divisor, source)
        got := p.Call_Main_sumRange(bounds[0], bounds[1], seed, divisor)
        if got != want { t.Fatalf("bounds=%v seed=%d divisor=%d: got %d want %d", bounds, seed, divisor, got, want) }
        if !slices.Equal(source, original) { t.Fatal("fallback changed the shared range") }
      }
    }
  }
}

func TestConstantCongruencesPreserveWideBoundsAndSeeds(t *testing.T) {
  for _, bounds := range [][2]int64{
    {1,3}, {3,1}, {1,1}, {1,2}, {2,1}, {4,4},
    {-3,3}, {3,-3},
    {2147483646,2147483648}, {2147483648,2147483650},
    {-2147483650,-2147483648}, {4294967295,4294967297},
    {2147483650,2147483648}, {-2147483648,-2147483650},
    {4294967297,4294967295},
  } {
    for _, seed := range []int64{0,4294967295,-2147483649,9223372036854775807,-9223372036854775808} {
      source := p.Call_Main_makeRange(bounds[0],bounds[1])
      original := slices.Clone(source)
      even := p.Call_Main_sumBorrowed(seed,2,source)
      residue := p.Call_Main_sumBorrowedResidue(seed,source)
      if got := p.Call_Main_sumEven(bounds[0],bounds[1],seed); got != even { t.Fatalf("even bounds=%v seed=%d got=%d want=%d",bounds,seed,got,even) }
      if got := p.Call_Main_sumResidue(bounds[0],bounds[1],seed); got != residue { t.Fatalf("residue bounds=%v seed=%d got=%d want=%d",bounds,seed,got,residue) }
      if !slices.Equal(source,original) { t.Fatal("fallback changed shared range") }
    }
  }
}
`;

test("array range/filter/sum fusion preserves range direction, effects and fallback semantics", async (t) => {
  assert.ok(modes.length > 0);
  const workspace = createWorkspace();
  let completed = 0;
  console.log(`Array range fold workspace: ${workspace}`);
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
    GOPURS_PREPARE_JOBS: "2", GOPURS_PBO_JOBS: "2", GOPURS_EMIT_JOBS: "2", GOPURS_PIPELINE: "1",
    GOPURS_FORCE_BUILD: "1",
  };
  for (const key of ["GOPURS_ALLOC_PROFILE", "GOPURS_COMPILER_SNAPSHOT", "PPROF"]) delete env[key];
  const packages = corePackages(root);
  for (const [index, mode] of modes.entries()) {
    await t.test(mode.name, () => {
      const { directory } = prepareFixture(root, workspace, join(root, "tools/fixtures/array-range-fold.purs"), index, packages);
      let sequence = 0;
      const run = (label, command, args, cwd = directory) => {
        console.log(`${mode.name}: ${label}`);
        const result = spawnSync(command, args, {
          cwd, env: { ...env, GOPURS_JS: mode.GOPURS_JS }, encoding: "utf8",
          timeout: 120000, maxBuffer: 16 * 1024 * 1024,
        });
        writeFileSync(join(directory, "logs", `${sequence++}-${label}.log`), (result.stdout ?? "") + (result.stderr ?? ""));
        assert.ifError(result.error);
        assert.equal(result.status, 0, `${label}\n${result.stdout}\n${result.stderr}`);
        return result.stdout;
      };
      run("frontend", "spago", ["build", "-q", "--offline"]);
      const reference = run("javascript", "node", ["--input-type=module", "-e",
        'import { main } from "./output/Main/index.js"; main();']);
      assert.deepEqual(parseResults(reference), expected, "reference evaluation order and inclusive ranges");
      for (const profile of ["release", "dev"]) {
        run(profile, process.env.GOPURS_TEST_COMPILER ?? join(root, "bin/gopurs"), [`--${profile}`, "--main", "Main"]);
        if (profile === "release") {
          const source = readFileSync(join(directory, "output/purescript/p_root/0_gopurs_package.go"), "utf8");
          const worker = (name) => {
            const start = source.indexOf(`func Call_Main_${name}(`);
            assert.ok(start >= 0, `${name}: generated worker exists`);
            const end = source.indexOf("\nfunc ", start + 1);
            return source.slice(start, end < 0 ? source.length : end);
          };
          const admitted = ["sumEven", "sumResidue", "observedPureSum"].map((name) => {
            const body = worker(name);
            for (const pattern of [/first :=/, /count :=/, /last :=/, /int64\(int32\(/, /range_filter_fold/]) {
              assert.match(body, pattern, `${name}: exact arithmetic sum with original loop fallback`);
            }
            return body;
          });
          for (const name of ["sumRange", "observedSum"]) {
            assert.doesNotMatch(worker(name), /count :=/, `${name}: dynamic or effectful predicate keeps its invocations`);
          }
          writeFileSync(join(directory, "logs/release-congruence.go"), admitted.join("\n"));
        }
        assert.deepEqual(parseResults(run(`${profile}-go`, "go", ["run", "./main"], join(directory, "output"))),
          expected, `${profile}: values, shared arrays and callback phases`);
        if (profile === "release") {
          const output = join(directory, "output");
          mkdirSync(join(output, "array-range-probe"));
          writeFileSync(join(output, "array-range-probe", "range_test.go"), allocationProbe);
          run("release-allocations-and-fallback", "go", ["test", "-count=1", "-timeout=30s", "-v", "./array-range-probe"], output);
        }
      }
      completed++;
    });
  }
});
