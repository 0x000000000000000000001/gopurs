import assert from 'node:assert/strict';
import { execFileSync } from 'node:child_process';
import { copyFileSync, mkdirSync, mkdtempSync, readFileSync, rmSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join } from 'node:path';
import { fileURLToPath } from 'node:url';
import test from 'node:test';
import { Just } from '../output/Data.Maybe/index.js';
import { Tuple } from '../output/Data.Tuple/index.js';
import { generateFfiBridge } from '../output/Gopurs.FfiBridge/index.js';
import { prepareFfi } from '../output/Gopurs.FfiSupport/index.js';
import * as C from '../output/PureScript.Backend.Optimizer.CoreFn/index.js';

const root = fileURLToPath(new URL('../', import.meta.url));
const fn = (args, result) => new C.Func(args, result);
const opaque = name => new C.ADT(name, name.split('.'), []);
const effect = result => new C.ADT('Effect.Effect', ['Effect', 'Effect'], [result]);
const strings = new C.Array(C.String.value);
const pool = opaque('Gopurs.Cache.Pool.Pool');
const unitEffect = effect(C.Unit.value);
const modules = [
  ['Cache/Pool', [
    ['create', effect(pool)],
    ['intern', fn([pool, C.String.value], effect(C.Int.value))],
    ['internRef', fn([pool, C.String.value], effect(C.String.value))],
    ['finish', fn([pool], effect(C.String.value))],
  ]],
  ['GoCode', [['referencedImportsImpl', fn([fn([C.String.value], strings), C.String.value], strings)]]],
  ['GoImports', [['concatStringArrays', fn([new C.Array(strings)], strings)]]],
  ['BuildCache', [
    ['exchange', fn([C.String.value], effect(C.String.value))],
    ['sessionExchange', fn([C.String.value], effect(C.String.value))],
    ['sessionExchangeOnMiss', fn([C.String.value, fn([C.String.value], unitEffect)], effect(C.String.value))],
  ]],
  ['Host/Snapshot', [['withSnapshot', fn([unitEffect], unitEffect)], ['release', unitEffect]]],
  ['Host/Metrics', [['setMemProfileRate', fn([C.Int.value], unitEffect)], ['now', effect(C.Number.value)]]],
];

test('compiler FFI generated bridges preserve native handles, nested arrays and effect timing', t => {
  const directory = mkdtempSync(join(tmpdir(), 'gopurs-compiler-ffi-'));
  t.after(() => rmSync(directory, { recursive: true, force: true }));
  for (const name of ['gopurs_runtime', 'gopurs_build_cache']) mkdirSync(join(directory, name));
  writeFileSync(join(directory, 'go.mod'), 'module gopurs/output\n\ngo 1.22\n');
  copyFileSync(join(root, 'runtime/runtime.go'), join(directory, 'gopurs_runtime/runtime.go'));
  for (const [name, entries] of modules) {
    const moduleName = 'Gopurs.' + name.replaceAll('/', '.');
    const prefix = moduleName.replaceAll('.', '_');
    const path = join(root, 'src/Gopurs', name + '.go');
    const prepared = prepareFfi({ moduleName, path })(prefix + '_')(readFileSync(path, 'utf8'))();
    const foreigns = entries.map(([name, type]) => new Tuple(name, new Just(type)));
    const bridge = generateFfiBridge(prefix)([])(prepared.decls)(foreigns);
    writeFileSync(join(directory, prefix + '.go'),
      'package main\nimport "gopurs/output/gopurs_runtime"\n' + prepared.content + '\n' + bridge);
  }
  // These host operations record calls so the test can observe the generated
  // effect boundary without starting compiler snapshots or a persistent cache.
  writeFileSync(join(directory, 'gopurs_build_cache/cache.go'), `package gopurs_build_cache
var Events []string
func Exchange(request string) string { Events = append(Events, "exchange:" + request); return request }
func SessionExchange(request string) string { Events = append(Events, "session:" + request); return request }
func SessionExchangeOnMiss(request string, onMiss func(string)) string {
    Events = append(Events, "miss:" + request)
    onMiss("seed")
    Events = append(Events, "returned")
    return request
}
func FreezeNativeCompiler() error { Events = append(Events, "freeze"); return nil }
func ReleaseCompilerSnapshot() { Events = append(Events, "release") }
`);
  copyFileSync(join(root, 'tools/native-ffi/compiler_bridge_test.go'), join(directory, 'bridge_test.go'));
  const output = execFileSync('go', ['test', '-race', '-count=1', '-timeout=45s', '-v', './...'], {
    cwd: directory, encoding: 'utf8', timeout: 90_000,
    env: { ...process.env, GOWORK: 'off' }, stdio: ['ignore', 'pipe', 'pipe'],
  });
  for (const name of ['CompilerHandles', 'CompilerImportArrays', 'CompilerHostEffects']) {
    assert.match(output, new RegExp(`--- PASS: Test${name}`));
  }
});
