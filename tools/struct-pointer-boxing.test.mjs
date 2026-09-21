import assert from 'node:assert/strict';
import { spawnSync } from 'node:child_process';
import { mkdirSync, mkdtempSync, rmSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join } from 'node:path';
import test from 'node:test';
import { empty as emptySet } from '../output/Data.Set/index.js';
import * as Ref from '../output/Effect.Ref/index.js';
import * as Go from '../output/Gopurs.GoAst/index.js';
import { boxGoExpr } from '../output/Gopurs.GoConversions/index.js';
import { collectImports } from '../output/Gopurs.GoImports/index.js';
import { printGoExpr, printGoFile } from '../output/Gopurs.Printer/index.js';
import { runtimeGoCode } from '../output/Gopurs.Runtime/index.js';
import { hashString } from '../output/PureScript.Backend.Optimizer.FfiSupport/index.js';

const base = 'Constructor_Probe_Node';
const pointer = Go.structPointer({ baseStructName: base, fullName: 'Probe.Node', structName: base })([]);
const box = expression => boxGoExpr(Ref.new({ declarations: [], globalId: 0, reboxPairs: emptySet })())('Probe')(expression)(pointer);
const importsOf = expression => collectImports([[new Go.GoFunctionDecl({ name: 'probe', params: [], result: Go.TypeValue.value, body: new Go.GoReturn(expression) })]]);
const legacy = expression => `gopurs_runtime.Value{Type: 9, IntVal: ${hashString(base)}, UnsafePtr: unsafe.Pointer(${printGoExpr(expression)})}`;

for (const [name, operand, imports] of [
    ['variable', new Go.GoVar('node'), ['gopurs/output/gopurs_runtime', 'unsafe']],
    ['nested code', Go.rawGo('func() *Node { _ = "sync.Once"; /* unsafe.Pointer */ return &Node{N: int64(math.Abs(-7))} }()'), ['gopurs/output/gopurs_runtime', 'math', 'unsafe']],
    ['structured call', new Go.GoCall(new Go.GoVar('lookup'), [new Go.GoCall(new Go.GoSelector(new Go.GoVar('math'), 'Abs'), [new Go.GoInt(-7)]), new Go.GoString('sync.Once')]), ['gopurs/output/gopurs_runtime', 'math', 'unsafe']],
]) test(`pointer boxing preserves text and imports for ${name}`, () => {
    const expression = box(operand);
    assert.ok(expression instanceof Go.GoBoxStructPointer);
    assert.equal(printGoExpr(expression), legacy(operand));
    assert.deepEqual([...new Set(importsOf(expression))].sort(), imports);
    assert.deepEqual([...new Set(importsOf(expression))].sort(), [...new Set(importsOf(Go.rawGo(legacy(operand))))].sort());
});

test('structured pointer boxing preserves tag, pointer identity, nil, and single evaluation', t => {
    const expression = box(new Go.GoCall(new Go.GoVar('next'), []));
    const declarationGroups = [[new Go.GoFunctionDecl({ name: 'boxed', params: [], result: Go.TypeValue.value, body: new Go.GoReturn(expression) })]];
    const code = printGoFile({ packageName: 'main', imports: collectImports(declarationGroups), declarationGroups });
    const directory = mkdtempSync(join(tmpdir(), 'gopurs-struct-pointer-boxing-'));
    t.after(() => rmSync(directory, { recursive: true, force: true }));
    mkdirSync(join(directory, 'gopurs_runtime'));
    writeFileSync(join(directory, 'gopurs_runtime/runtime.go'), runtimeGoCode);
    writeFileSync(join(directory, 'go.mod'), 'module gopurs/output\n\ngo 1.22\n');
    writeFileSync(join(directory, 'main.go'), `${code}
type Node struct { N int64 }
var node = &Node{N: 42}
var reads int
func next() *Node { reads++; return node }
func main() {
    result := boxed()
    if reads != 1 || result.Type != 9 || result.IntVal != ${hashString(base)} || result.UnsafePtr != unsafe.Pointer(node) { panic("pointer or tag changed, or operand repeated") }
    node = nil
    result = boxed()
    if reads != 2 || result.Type != 9 || result.IntVal != ${hashString(base)} || result.UnsafePtr != nil { panic("nil pointer changed") }
}
`);
    const run = spawnSync('go', ['run', '.'], { cwd: directory, encoding: 'utf8', timeout: 30_000, env: { ...process.env, GOWORK: 'off' } });
    assert.ifError(run.error);
    assert.equal(run.status, 0, run.stdout + run.stderr);
});
