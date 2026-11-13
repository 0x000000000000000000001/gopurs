import assert from 'node:assert/strict';
import { spawnSync } from 'node:child_process';
import { copyFileSync, mkdirSync, mkdtempSync, readFileSync, rmSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join } from 'node:path';
import { fileURLToPath } from 'node:url';
import test from 'node:test';
import { withReboxFields } from './codegen-metadata.mjs';
import * as Map from '../output/Data.Map/index.js';
import { Just } from '../output/Data.Maybe/index.js';
import { empty as emptySet } from '../output/Data.Set/index.js';
import { ordString } from '../output/Data.Ord/index.js';
import { Left } from '../output/Data.Either/index.js';
import * as Aff from '../output/Effect.Aff/index.js';
import { emitModule } from '../output/Gopurs.Driver.Output/index.js';
import * as C from '../output/PureScript.Backend.Optimizer.CoreFn/index.js';

const root = fileURLToPath(new URL('../', import.meta.url));
const runAff = action => new Promise((resolve, reject) => {
  Aff.runAff(result => () => result instanceof Left ? reject(result.value0) : resolve(result.value0))(action)();
});
const map = entries => entries.reduce((result, [key, value]) => Map.insert(ordString)(key)(value)(result), Map.empty);
const fn = (args, result) => new C.Func(args, result);
const strings = new C.Array(C.String.value);
const metadata = withReboxFields({
  elidedCtors: emptySet, ctorTypes: Map.empty, pointerAdtPaths: Map.empty,
  pointerAdtNodes: emptySet, pointerAdtLeaves: Map.empty, enumAdts: emptySet,
  enumCtors: emptySet, globalTypes: Map.empty, globalFunctions: Map.empty,
  classDeclsFields: Map.empty,
});

test('the output writer derives runtime imports from Go declarations, including the compiler scanner', async t => {
  const directory = mkdtempSync(join(tmpdir(), 'gopurs-ffi-output-'));
  t.after(() => rmSync(directory, { recursive: true, force: true }));
  mkdirSync(join(directory, 'purescript'));
  mkdirSync(join(directory, 'gopurs_runtime'));
  copyFileSync(join(root, 'runtime/runtime.go'), join(directory, 'gopurs_runtime/runtime.go'));
  writeFileSync(join(directory, 'go.mod'), 'module gopurs/output\n\ngo 1.22\n');
  const cases = [
    ['Literal', `// import "gopurs/output/gopurs_runtime"
func Echo(value string) string { _ = "gopurs/output/gopurs_runtime"; return value }
`, 'echo', fn([C.String.value], C.String.value)],
    ['Grouped', `import (
    "gopurs/output/gopurs_runtime"
)
func Echo(value string) string { return gopurs_runtime.Str(value).StrVal() }
`, 'echo', fn([C.String.value], C.String.value)],
    ['Aliased', `import rt "gopurs/output/gopurs_runtime"
func Echo(value string) string { return rt.Str(value).StrVal() }
`, 'echo', fn([C.String.value], C.String.value)],
    ['RawImport', 'import `gopurs/output/gopurs_runtime`\nfunc Echo(value string) string { return gopurs_runtime.Str(value).StrVal() }\n',
      'echo', fn([C.String.value], C.String.value)],
    ['Gopurs.GoCode', readFileSync(join(root, 'src/Gopurs/GoCode.go'), 'utf8'), 'referencedImportsImpl',
      fn([fn([C.String.value], strings), C.String.value], strings)],
  ];
  for (const [name, content, binding, type] of cases) {
    const path = join(directory, name + '.go');
    writeFileSync(path, content);
    const backend = {
      name, bindings: [], comments: [], imports: emptySet, exports: emptySet, reExports: emptySet,
      dataTypes: Map.empty, dataDecls: [], classDecls: [], foreign: map([[binding, new Just(type)]]),
      implementations: Map.empty, directives: Map.empty,
    };
    const frozen = map([[name, { name, path, file: path, found: true, digest: '' }]]);
    await runAff(emitModule(directory)(frozen)(metadata)(new C.Module({ path: path.replace(/\.go$/, '.purs') }))(backend));
  }
  writeFileSync(join(directory, 'purescript/imports_test.go'), `package purescript
import (
    "testing"
    rt "gopurs/output/gopurs_runtime"
)
func TestEmittedForeignFunctions(t *testing.T) {
    for _, fn := range []rt.Value{_Gopurs_Literal_Echo, _Gopurs_Grouped_Echo, _Gopurs_Aliased_Echo, _Gopurs_RawImport_Echo} {
        if got := rt.Apply(fn, rt.Str("kept")).StrVal(); got != "kept" { t.Fatalf("echo: %q", got) }
    }
    fallback := rt.Func(func(rt.Value) rt.Value { panic("native scanner used fallback") })
    values := rt.Unbox[[]rt.Value](rt.Apply2(_Gopurs_Gopurs_GoCode_ReferencedImportsImpl, fallback, rt.Str("math.Abs(1)")))
    if len(values) != 1 || values[0].StrVal() != "math" { t.Fatal("native scanner imports") }
}
`);
  const result = spawnSync('go', ['test', '-count=1', '-v', './purescript'], {
    cwd: directory, encoding: 'utf8', timeout: 60_000, env: { ...process.env, GOWORK: 'off' },
  });
  assert.ifError(result.error);
  assert.equal(result.status, 0, result.stdout + result.stderr);
  assert.match(result.stdout, /--- PASS: TestEmittedForeignFunctions/);
});
