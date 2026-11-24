// Exercise a renamed, strict list eliminator through the public compiler.
// The independent oracle filters arrays; it does not use a marking table.
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
  [4, 3, 123], [2, 2, 0], [2, 3, 17], [2, 30, 0], [4, 10, 7], [6, 18, -17],
  [10, 30, 2147483647], [1, 10, -2147483648], [0, 8, 23], [-4, 4, 17],
  [-8, -1, -2147483648], [-2147483647, -2147483645, 0],
  [2147483644, 2147483647, 2147483647], [2147483647, 2147483647, -2147483648],
  [2, 500, 0],
];
const int32 = (value) => BigInt.asIntN(32, value);
const progression = (lower, upper) => Array.from({ length: Math.max(0, upper - lower + 1) },
  (_, index) => BigInt(lower + index));
const modulo = (value, divisor) => {
  if (divisor === 0n) return 0n;
  const positive = divisor < 0n ? -divisor : divisor;
  return (value % positive + positive) % positive;
};
const fold = (values, seed = 0, multiplier = 1n) => values.reduce(
  (result, value) => int32(result * multiplier + value), BigInt(seed));
function eliminate(values, withTrace = false) {
  let pending = [...values];
  const selected = [], trace = [];
  const fingerprint = Number(fold(values, 0, 31n));
  while (pending.length > 0) {
    const [leader, ...tail] = pending;
    selected.push(leader);
    pending = tail.filter((value) => {
      if (withTrace) trace.push([Number(leader), Number(value), fingerprint]);
      return modulo(value, leader) !== 0n;
    });
  }
  return { selected, trace };
}
assert.deepEqual(eliminate(progression(4, 10)).selected.map(Number), [4, 5, 6, 7, 9]);
assert.equal(Number(fold(eliminate(progression(4, 10)).selected)), 31);
const expected = cases.flatMap(([lower, upper, seed]) => {
  const source = progression(lower, upper);
  const { selected, trace } = eliminate(source, true);
  const sum = Number(fold(selected, seed)), zero = Number(fold(selected));
  return [
    [source.map(Number), selected.map(Number), source.map(Number)],
    [sum, zero, zero, sum, Number(fold(selected, 0, 31n))],
    sum, trace, source.map(Number), sum,
    [[0, seed], [1, lower], [2, upper]],
  ];
});
const parseResults = (stdout) => stdout.trim().split("\n").map((line) => JSON.parse(line));

const positiveCases = [
  ...cases.filter(([lower, upper]) => lower >= 2 && upper >= lower),
  [4, 1027, 17], [4, 1028, -17], [1000, 2023, 2147483647], [1000, 2024, -2147483648],
  [2147482624, 2147483647, 17], [2147482623, 2147483647, -17],
].map(([lower, upper, seed]) => [lower, upper, seed,
  Number(fold(eliminate(progression(lower, upper)).selected, seed))]);
// Every value is greater than half the maximum, so no later value can be a
// multiple of an earlier one. This gives a separate exact large-span oracle.
const cappedCases = [1024, 1025, 1048576].flatMap((size) => [0, 2147483647, -2147483648].map((seed) => {
  const upper = 2147483647, lower = upper - size + 1;
  const exact = BigInt(size) * BigInt(lower + upper) / 2n + BigInt(seed);
  return [lower, upper, seed, Number(int32(exact))];
}));
const rows = (values) => values.map((row) => `{${row.join(",")}}`).join(",\n");
const allocationProbe = `package progressionprobe

import (
  "runtime"
  "testing"
  p "gopurs/output/purescript/p_root"
  _ "gopurs/output/purescript/init_all"
)

var resultSink int64

func TestReleaseClosedPipelineHasConstantAllocationCount(t *testing.T) {
  cases := []struct{lower,upper,seed,want int64}{${rows(positiveCases)}}
  // The scratch buffer uses at most one allocation. The enclosing public
  // function also retains one recursive-fold closure for its guarded fallback.
  // Neither allocation count grows with the input or survivor count.
  for _, c := range cases {
    allocations := testing.AllocsPerRun(10, func() { resultSink = p.Call_Main_sumRange(c.lower,c.upper,c.seed) })
    if resultSink != c.want { t.Fatalf("%+v: got %d",c,resultSink) }
    if allocations > 2 { t.Fatalf("%+v: %g allocations; list spines were not eliminated",c,allocations) }
  }
  // Run large spans only after the allocation contract has established that
  // the true optimized pipeline is admitted; the original list path is huge.
  for _, c := range []struct{lower,upper,seed,want int64}{${rows(cappedCases)}} {
    allocations:=testing.AllocsPerRun(2,func(){resultSink=p.Call_Main_sumRange(c.lower,c.upper,c.seed)})
    if resultSink!=c.want {t.Fatalf("large span %+v: got %d",c,resultSink)}
    if allocations>2 {t.Fatalf("large span %+v: %g allocations; count must remain bounded",c,allocations)}
  }
  t.Log("${positiveCases.length} allocation cases and ${cappedCases.length} exact large-span cases passed")
}

type node = p.Constructor_Main_Cell[int64]

func TestBorrowedSourcesAndNoncanonicalInputsRetainOriginalSemantics(t *testing.T) {
  // Every wide case terminates quickly under the original IntSub32 producer.
  // In particular, never start a nonempty range below MinInt32.
  bounds := [][2]int64{
    {4,3},{2,10},{4,10},{0,8},{-4,4},{-2147483647,-2147483645},
    {2147483645,2147483647},{2147483646,2147483648},
    {2147483648,2147483650},{4294967295,4294967297},
    {2,4294967296},{2,4294967297},{-2147483648,-2147483649},
    {9223372036854775807,9223372036854775806},{9223372036854775807,9223372036854775807},
  }
  checked:=0
  for _, b:=range bounds {
    source:=p.Call_Main_makeTrail(b[0],b[1])
    originals,links,values:=[]*node{},[]*node{},[]int64{}
    for cursor:=source;cursor!=nil;cursor=cursor.V1 {
      originals,links,values=append(originals,cursor),append(links,cursor.V1),append(values,cursor.V0)
      if len(originals)>100 {t.Fatal("supposedly short source did not terminate")}
    }
    checkSource:=func() {
      cursor:=source
      for index,original:=range originals {
        if cursor!=original||original.V1!=links[index]||original.V0!=values[index] {t.Fatalf("bounds=%v: mutated source node %d",b,index)}
        cursor=cursor.V1
      }
      if cursor!=nil {t.Fatal("source gained a tail")}
    }
    for _,seed:=range []int64{0,17,2147483647,-2147483648,2147483648,4294967295,-2147483649,9223372036854775807,-9223372036854775808} {
      want:=p.Call_Main_sumBorrowed(seed,source)
      checkSource()
      got:=p.Call_Main_sumRange(b[0],b[1],seed)
      if got!=want {t.Fatalf("bounds=%v seed=%d: got %d want %d",b,seed,got,want)}
      checkSource()
      checked++
    }
    runtime.KeepAlive(source)
  }
  t.Logf("%d short boundary/seed comparisons and shared-source snapshots passed",checked)
}
`;

test("closed list progression elimination preserves general bounds, Int32 sums and effectful fallbacks", async (t) => {
  assert.ok(modes.length > 0);
  const workspace = createWorkspace();
  let completed = 0;
  console.log(`List progression fold workspace: ${workspace}`);
  t.after(() => {
    if (completed === modes.length && process.env.GOPURS_TEST_KEEP_WORKSPACE !== "1") rmSync(workspace, { recursive: true, force: true });
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
      const { directory } = prepareFixture(root, workspace, join(root, "tools/fixtures/list-progression-fold.purs"), index, packages);
      let sequence = 0;
      const run = (label, command, args, cwd = directory) => {
        console.log(`${mode.name}: ${label}`);
        const result = spawnSync(command, args, { cwd, env: { ...env, GOPURS_JS: mode.GOPURS_JS }, encoding: "utf8", timeout: 120000, maxBuffer: 16 * 1024 * 1024 });
        writeFileSync(join(directory, "logs", `${sequence++}-${label}.log`), (result.stdout ?? "") + (result.stderr ?? ""));
        assert.ifError(result.error);
        assert.equal(result.status, 0, `${label}\n${result.stdout}\n${result.stderr}`);
        return result.stdout;
      };
      run("frontend", "spago", ["build", "-q", "--offline"]);
      const reference = run("javascript", "node", ["--input-type=module", "-e", 'import { main } from "./output/Main/index.js"; main();']);
      assert.deepEqual(parseResults(reference), expected, "strict array-filter oracle, general lower bounds and callback phases");
      for (const profile of ["release", "dev"]) {
        run(profile, process.env.GOPURS_TEST_COMPILER ?? join(root, "bin/gopurs"), [`--${profile}`, "--main", "Main"]);
        const emitted = readFileSync(join(directory, "output/purescript/p_root/0_gopurs_package.go"), "utf8");
        const opaqueFactory = emitted.match(/func Call_Main_observedPureRange\([\s\S]*?(?=\nfunc |$)/)?.[0];
        assert.ok(opaqueFactory, `${profile}: opaque argument factory remains available`);
        assert.doesNotMatch(opaqueFactory, /gopurs_list_progression_fold/, `${profile}: opaque seed/bounds must retain their original pipeline`);
        const actual = parseResults(run(`${profile}-go`, "go", ["run", "./main"], join(directory, "output")));
        assert.deepEqual(actual, expected, `${profile}: ${cases.length} inputs, general elimination, aliasing and effects`);
        if (profile === "release") {
          const output = join(directory, "output");
          mkdirSync(join(output, "list-progression-probe"));
          writeFileSync(join(output, "list-progression-probe", "progression_test.go"), allocationProbe);
          run("release-allocations-and-fallback", "go", ["test", "-count=1", "-v", "./list-progression-probe"], output);
        }
      }
      completed++;
    });
  }
});
