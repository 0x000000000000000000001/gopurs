import assert from 'node:assert/strict';
import { execFileSync } from 'node:child_process';
import { copyFileSync, mkdirSync, mkdtempSync, readFileSync, rmSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join } from 'node:path';
import test from 'node:test';
import { fileURLToPath } from 'node:url';
import { lookup } from '../output/Data.Map/index.js';
import { Just, Nothing } from '../output/Data.Maybe/index.js';
import { ordString } from '../output/Data.Ord/index.js';
import { Tuple } from '../output/Data.Tuple/index.js';
import { ffiFunctionInfos, ffiValueWorkers, generateFfiBridge } from '../output/Gopurs.FfiBridge/index.js';
import { prepareFfi } from '../output/Gopurs.FfiSupport/index.js';
import * as Go from '../output/Gopurs.GoAst/index.js';
import * as C from '../output/PureScript.Backend.Optimizer.CoreFn/index.js';

const root = fileURLToPath(new URL('../', import.meta.url));
const name = 'BuilderProbe';
const handle = new C.ADT('BuilderProbe.Builder', ['BuilderProbe', 'Builder'], []);
const fn = (args, result) => new C.Func(args, result);
const foreigns = [
  ['newBuilderImpl', new C.ADT('Effect.Effect', ['Effect', 'Effect'], [handle])],
  ['pushImpl', fn([handle, C.String.value], handle)],
  ['toStringImpl', fn([handle], C.String.value)],
  ['choose', fn([handle, handle, C.Boolean.value], handle)],
  ['chooseCurried', fn([handle], fn([handle, C.Boolean.value], handle))],
  ['aliasEcho', fn([handle], handle)],
  ['fresh', fn([handle], handle)],
  ['factory', fn([C.Int.value], handle)],
].map(([name, type]) => new Tuple(name, new Just(type)));

test('pointer FFI keeps direct workers, native scalar results and existing handle envelopes', t => {
  const path = join(root, 'src/Gopurs/Printer/Builder.go');
  const source = readFileSync(path, 'utf8') + `
func Choose(first, second *strings.Builder, useSecond bool) *strings.Builder {
    if useSecond { return second }; return first
}
func ChooseCurried(first, second *strings.Builder, useSecond bool) *strings.Builder {
    if useSecond { return second }; return first
}
type BuilderAlias = *strings.Builder
func AliasEcho(value BuilderAlias) BuilderAlias { return value }
func Fresh(previous *strings.Builder) *strings.Builder { return new(strings.Builder) }
func Factory[T any](seed int64) *T { return new(T) }
`;
  const prepared = prepareFfi({ moduleName: name, path })(name + '_')(source)();
  const signatures = ffiFunctionInfos(name)(foreigns)(prepared.decls);
  assert.equal(lookup(ordString)('factory')(signatures), Nothing.value);
  const workers = ffiValueWorkers(name)([])(foreigns)(prepared.decls);
  assert.doesNotMatch(workers, /BuilderProbe_Factory_nativeWorker/);
  for (const [binding, args, result] of [
    ['pushImpl', [Go.TypeValue.value, Go.TypeString.value], Go.TypeValue.value],
    ['toStringImpl', [Go.TypeValue.value], Go.TypeString.value],
    ['chooseCurried', [Go.TypeValue.value, Go.TypeValue.value, Go.TypeBool.value], Go.TypeValue.value],
    ['aliasEcho', [Go.TypeValue.value], Go.TypeValue.value],
  ]) {
    const found = lookup(ordString)(binding)(signatures);
    assert.ok(found instanceof Just, `${binding} must keep a direct native worker`);
    assert.deepEqual(found.value0.fArgs, args);
    assert.deepEqual(found.value0.fRet, result);
    assert.match(found.value0.reference.value0.name, /_nativeWorker$/);
  }
  const directory = mkdtempSync(join(tmpdir(), 'gopurs-pointer-handles-'));
  t.after(() => rmSync(directory, { recursive: true, force: true }));
  mkdirSync(join(directory, 'gopurs_runtime'));
  copyFileSync(join(root, 'runtime/runtime.go'), join(directory, 'gopurs_runtime/runtime.go'));
  writeFileSync(join(directory, 'go.mod'), 'module gopurs/output\n\ngo 1.22\n');
  writeFileSync(join(directory, 'bridge.go'), 'package main\nimport "gopurs/output/gopurs_runtime"\n'
    + prepared.content + '\n' + generateFfiBridge(name)([])(prepared.decls)(foreigns)
    + '\n' + workers);
  writeFileSync(join(directory, 'bridge_test.go'), `package main
import (
    "runtime"
    "strings"
    "testing"
    rt "gopurs/output/gopurs_runtime"
)
var sink rt.Value
func TestPointerHandles(t *testing.T) {
    generic := rt.Apply(_Gopurs_BuilderProbe_Factory, rt.Int(42))
    if got := rt.Unbox[*rt.Value](generic); got == nil || *got != (rt.Value{}) {
        t.Fatal("generic pointer factory lost its boxed instantiation")
    }
    first := rt.Apply(_Gopurs_BuilderProbe_NewBuilderImpl, rt.Value{})
    second := rt.Apply(_Gopurs_BuilderProbe_NewBuilderImpl, rt.Value{})
    rt.Unbox[*strings.Builder](first).Grow(8192)
    if got := BuilderProbe_PushImpl_nativeWorker(first, "é😀"); got != first {
        t.Fatal("direct call reboxed an unchanged handle")
    }
    if got := rt.Apply2(_Gopurs_BuilderProbe_PushImpl, first, rt.Str("!")); got != first {
        t.Fatal("first-class call reboxed an unchanged handle")
    }
    if got := BuilderProbe_Choose_nativeWorker(first, second, true); got != second {
        t.Fatal("returning the second input lost its envelope")
    }
    if got := rt.Apply3(_Gopurs_BuilderProbe_Choose, first, second, rt.Bool(false)); got != first {
        t.Fatal("boxed choice lost the first envelope")
    }
    if got := BuilderProbe_ChooseCurried_nativeWorker(first, second, true); got != second {
        t.Fatal("direct curried choice lost the second envelope at the native call arity")
    }
    if got := rt.Apply3(_Gopurs_BuilderProbe_ChooseCurried, first, second, rt.Bool(true)); got != second {
        t.Fatal("boxed curried choice lost the second envelope at the native call arity")
    }
    if got := BuilderProbe_AliasEcho_nativeWorker(first); got != first {
        t.Fatal("pointer alias direct call lost the existing envelope")
    }
    if got := rt.Apply(_Gopurs_BuilderProbe_AliasEcho, first); got != first {
        t.Fatal("pointer alias boxed call lost the existing envelope")
    }
    fresh := BuilderProbe_Fresh_nativeWorker(first)
    if rt.Unbox[*strings.Builder](fresh) == rt.Unbox[*strings.Builder](first) {
        t.Fatal("a new native pointer reused an unrelated handle")
    }
    allocations := testing.AllocsPerRun(1000, func() { sink = BuilderProbe_PushImpl_nativeWorker(first, "x") })
    if allocations != 0 { t.Fatalf("per-fragment handle allocations: %g", allocations) }
    completed := BuilderProbe_ToStringImpl_nativeWorker(first)
    if !strings.HasPrefix(completed, "é😀!") { t.Fatal("native result changed") }
    first, sink = rt.Value{}, rt.Value{}
    runtime.GC()
    if !strings.HasPrefix(completed, "é😀!") || BuilderProbe_ToStringImpl_nativeWorker(second) != "" {
        t.Fatal("completed string or independent handle lost storage")
    }
}
`);
  const output = execFileSync('go', ['test', '-race', '-count=1', '-v', './...'], {
    cwd: directory, encoding: 'utf8', timeout: 60_000,
    env: { ...process.env, GOWORK: 'off' }, stdio: ['ignore', 'pipe', 'pipe'],
  });
  assert.match(output, /--- PASS: TestPointerHandles/);
});
