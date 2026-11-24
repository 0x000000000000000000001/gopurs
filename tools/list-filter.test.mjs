// Check observable semantics and list-node allocation growth independently of
// the suffix-sharing pass's structural IR tests and generated helper names.
import assert from "node:assert/strict";
import { spawnSync } from "node:child_process";
import { mkdirSync, rmSync, symlinkSync, writeFileSync } from "node:fs";
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
  [4, 3, 2], [0, 0, 0], [-4, 4, 2], [-4, 4, -3], [-4, 4, 0],
  [-4, 4, 1], [-4, 4, -1], [1, 9, 2], [1, 9, 20], [1, 12, 3], [8, 8, 3],
];
const expected = cases.flatMap(([lower, upper, divisor]) => {
  const source = Array.from({ length: Math.max(0, upper - lower + 1) }, (_, index) => lower + index);
  // PureScript's IntMod returns zero for divisor zero, including negative inputs.
  const keep = (value) => divisor === 0 || value % divisor === 0;
  const selected = source.filter(keep);
  const checksum = (value) => value * value + 3;
  const fingerprint = source.reduce((result, value) => (Math.imul(result, 31) + value) | 0, 0);
  const visits = source.map((value) => [value, fingerprint]);
  const factoryVisit = [-1001, divisor];
  return [
    [selected, source.filter((value) => !keep(value)), selected.filter((value) => value > 0),
      source, selected, selected, selected.map(checksum), source.map(checksum)],
    selected,
    visits,
    selected,
    visits,
    selected,
    [factoryVisit, [-1002, fingerprint], ...visits],
    selected,
    visits,
    [factoryVisit],
    selected,
    selected,
    [factoryVisit, ...visits, ...visits],
  ];
});
const parseResults = (stdout) => stdout.trim().split("\n").map((line) => JSON.parse(line));

const allocationProbe = `package listfilterprobe

import (
  "runtime"
  "testing"
  p "gopurs/output/purescript/p_root"
  _ "gopurs/output/purescript/init_all"
)

func TestReleaseFilterCopiesOnlyRetainedValuesBeforeTheLastRejection(t *testing.T) {
  for _, divisor := range []int64{0, 1, 2, 3, -3} {
    var previousAllocations float64
    var previousCopied int64
    for _, size := range []int64{0, 1, 64, 128, 256} {
      // Input construction and all observations stay outside the measured call.
      source := p.Call_Main_makeValues(0, size - 1)
      originalFingerprint := p.Call_Main_fingerprintInts(source)
      retained, fingerprint, lastRejected := int64(0), int64(0), int64(-1)
      for value := int64(0); value < size; value++ {
        if divisor == 0 || value % divisor == 0 {
          retained++
          fingerprint = int64(int32(fingerprint * 31 + value))
        } else {
          lastRejected = value
        }
      }
      copied := int64(0)
      for value := int64(0); value < lastRejected; value++ {
        if divisor == 0 || value % divisor == 0 { copied++ }
      }
      result := source
      allocations := testing.AllocsPerRun(10, func() { result = p.Call_Main_keepDivisible(divisor, source) })
      if got := p.Call_Main_countInts(result); got != retained { t.Fatalf("size=%d divisor=%d: retained %d, want %d", size, divisor, got, retained) }
      if got := p.Call_Main_fingerprintInts(result); got != fingerprint { t.Fatalf("size=%d divisor=%d: reordered result", size, divisor) }
      if p.Call_Main_countInts(source) != size || p.Call_Main_fingerprintInts(source) != originalFingerprint {
        t.Fatalf("size=%d divisor=%d: mutated shared input", size, divisor)
      }
      // Only retained values before the last rejection need new nodes. The
      // immutable suffix can be returned or linked directly. Permit a small
      // fixed callback/worker setup cost, but no copies of that suffix.
      overhead := allocations - float64(copied)
      if overhead < 0 || overhead > 8 { t.Fatalf("size=%d divisor=%d: %g allocations for %d copied values", size, divisor, allocations, copied) }
      if size >= 128 && allocations - previousAllocations != float64(copied - previousCopied) {
        t.Fatalf("size=%d divisor=%d: allocation growth is not one node per copied prefix value", size, divisor)
      }
      previousAllocations, previousCopied = allocations, copied
      runtime.KeepAlive(result)
      runtime.KeepAlive(source)
    }
  }
}

type node = p.Constructor_FilterLibrary_Cell[int64]

func TestEverySelectionMaskSharesTheUnchangedSuffix(t *testing.T) {
  checked := 0
  for size := 0; size <= 8; size++ {
    source := p.Call_Main_makeValues(0, int64(size - 1))
    nodes, links := make([]*node, 0, size), make([]*node, 0, size)
    originalNodes := make(map[*node]bool, size)
    for cursor := source; cursor != nil; cursor = cursor.V1 {
      nodes, links = append(nodes, cursor), append(links, cursor.V1)
      originalNodes[cursor] = true
    }
    if len(nodes) != size { t.Fatalf("size=%d: malformed source", size) }
    assertSource := func() {
      cursor := source
      for index, original := range nodes {
        if cursor != original || original.V0 != int64(index) || original.V1 != links[index] {
          t.Fatalf("size=%d: source node %d was changed", size, index)
        }
        cursor = cursor.V1
      }
      if cursor != nil { t.Fatal("source grew a tail") }
    }
    for mask := 0; mask < 1 << uint(size); mask++ {
      expected, lastRejected := []int64{}, -1
      for index := 0; index < size; index++ {
        if mask & (1 << uint(index)) != 0 { expected = append(expected, int64(index)) } else { lastRejected = index }
      }
      copied := 0
      for _, value := range expected { if value < int64(lastRejected) { copied++ } }
      var suffix *node
      if lastRejected + 1 < size { suffix = nodes[lastRejected + 1] }
      visits := 0
      predicate := func(value int64) bool {
        if visits >= size || value != int64(visits) { t.Fatalf("size=%d mask=%d: callback visit %d got %d", size, mask, visits, value) }
        assertSource()
        selected := mask & (1 << uint(visits)) != 0
        visits++
        return selected
      }
      result := source
      allocations := testing.AllocsPerRun(10, func() {
        visits = 0
        result = p.Call_FilterLibrary_selectInts(predicate, source)
        if visits != size { t.Fatalf("size=%d mask=%d: got %d callback visits", size, mask, visits) }
      })
      if allocations != float64(copied) {
        t.Fatalf("size=%d mask=%d: %g allocations, expected %d copied prefix nodes", size, mask, allocations, copied)
      }
      cursor := result
      for index := 0; index < copied; index++ {
        if cursor == nil || originalNodes[cursor] || cursor.V0 != expected[index] {
          t.Fatalf("size=%d mask=%d: prefix node %d is not a fresh selected node", size, mask, index)
        }
        cursor = cursor.V1
      }
      if cursor != suffix { t.Fatalf("size=%d mask=%d: unchanged suffix was not shared", size, mask) }
      cursor = result
      for _, value := range expected {
        if cursor == nil || cursor.V0 != value { t.Fatalf("size=%d mask=%d: result reordered or truncated", size, mask) }
        cursor = cursor.V1
      }
      if cursor != nil { t.Fatalf("size=%d mask=%d: extra result nodes", size, mask) }
      assertSource()
      runtime.KeepAlive(result)
      runtime.KeepAlive(source)
      checked++
    }
  }
  if checked != 511 { t.Fatalf("checked %d masks, want 511", checked) }
  t.Log("511 masks: exact allocation count, source identity, callback order and unchanged suffix passed")
}

func TestCallbackPanicLeavesEverySourceNodeUntouched(t *testing.T) {
  for _, stop := range []int{0, 3, 7} {
    source := p.Call_Main_makeValues(0, 7)
    nodes, links := []*node{}, []*node{}
    for cursor := source; cursor != nil; cursor = cursor.V1 {
      nodes, links = append(nodes, cursor), append(links, cursor.V1)
    }
    assertSource := func() {
      cursor := source
      for index, original := range nodes {
        if cursor != original || original.V0 != int64(index) || original.V1 != links[index] { t.Fatalf("panic at %d changed input node %d", stop, index) }
        cursor = cursor.V1
      }
      if cursor != nil { t.Fatal("panic path grew an input tail") }
    }
    visits, token := 0, new(int)
    caught := func() (caught any) {
      defer func() { caught = recover() }()
      p.Call_FilterLibrary_selectInts(func(value int64) bool {
        if value != int64(visits) { t.Fatalf("panic at %d: out-of-order callback %d", stop, visits) }
        assertSource()
        visits++
        if value == int64(stop) { panic(token) }
        return value % 2 == 0
      }, source)
      return nil
    }()
    if caught != token || visits != stop + 1 { t.Fatalf("panic at %d: caught %v after %d visits", stop, caught, visits) }
    assertSource()
    runtime.KeepAlive(source)
  }
}
`;

test("stable list filtering preserves callbacks and shares only an unchanged immutable suffix", async (t) => {
  assert.ok(modes.length > 0);
  const workspace = createWorkspace();
  let completed = 0;
  console.log(`List filter workspace: ${workspace}`);
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
    GOWORK: "off", GOPURS_DAEMON: "0", GOPURS_RESIDENT: "0", GOPURS_RUST: "0",
    GOPURS_GO_PACKAGES: "1", GOPURS_GO_PACKAGE_OVERLAY: "1", GOPURS_BUILD_MANIFEST: "1",
    GOPURS_PREPARE_JOBS: "2", GOPURS_PBO_JOBS: "2", GOPURS_EMIT_JOBS: "2", GOPURS_PIPELINE: "1",
    GOPURS_FORCE_BUILD: "1",
  };
  for (const key of ["GOPURS_ALLOC_PROFILE", "GOPURS_COMPILER_SNAPSHOT", "PPROF"]) delete env[key];
  const packages = corePackages(root);
  for (const [index, mode] of modes.entries()) {
    await t.test(mode.name, () => {
      const { directory } = prepareFixture(root, workspace, join(root, "tools/fixtures/list-filter.purs"), index, packages);
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
      assert.deepEqual(parseResults(reference), expected, "reference factories run once, then callbacks run in input order");
      for (const profile of ["release", "dev"]) {
        run(profile, process.env.GOPURS_TEST_COMPILER ?? join(root, "bin/gopurs"), [`--${profile}`, "--main", "Main"]);
        assert.deepEqual(parseResults(run(`${profile}-go`, "go", ["run", "./main"], join(directory, "output"))),
          expected, `${profile}: ${cases.length} inputs preserve aliases, factory order, native and opaque callbacks, imported workers and reusable partials`);
        if (profile === "release") {
          const output = join(directory, "output");
          mkdirSync(join(output, "list-filter-probe"));
          writeFileSync(join(output, "list-filter-probe", "filter_test.go"), allocationProbe);
          run("release-allocations", "go", ["test", "-count=1", "-v", "./list-filter-probe"], output);
        }
      }
      completed++;
    });
  }
});
