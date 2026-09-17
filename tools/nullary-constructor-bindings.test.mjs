import assert from 'node:assert/strict';
import { spawnSync } from 'node:child_process';
import { mkdirSync, mkdtempSync, rmSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join } from 'node:path';
import test from 'node:test';
import * as Go from '../output/Gopurs.GoAst/index.js';
import { collectImports } from '../output/Gopurs.GoImports/index.js';
import { printGoExpr } from '../output/Gopurs.Printer/index.js';

test('nullary constructor values infer their type in local bindings, including generic imports', t => {
    const empty = new Go.GoConstructor('123', 'Constructor_Empty', [], []);
    const generic = new Go.GoConstructor('456', 'Constructor_Generic', [Go.TypeValue.value], []);
    const declarations = [[new Go.GoFunctionDecl({
        name: 'Probe', params: [], result: Go.TypeBool.value,
        body: new Go.GoBlock([
            new Go.GoAssign('plain', empty),
            new Go.GoAssign('generic', generic),
            new Go.GoReturn(new Go.GoBinOp('&&',
                new Go.GoBinOp('==', new Go.GoVar('plain'), new Go.GoVar('nil')),
                new Go.GoBinOp('==', new Go.GoVar('generic'), new Go.GoVar('nil')))),
        ]),
    })]];
    assert.deepEqual(collectImports(declarations), ['gopurs/output/gopurs_runtime'],
        'a nullary constructor still refers to its generic argument types');
    const syncGeneric = new Go.GoConstructor('789', 'Constructor_Generic', [new Go.TypeInterface('sync.Mutex')], []);
    const iife = new Go.GoIIFE('inside', syncGeneric, Go.rawGo('gopurs_runtime.Value{}'));
    const discardIife = new Go.GoIIFE('_', generic, Go.rawGo('gopurs_runtime.Value{}'));
    assert.deepEqual(collectImports([[new Go.GoInitDecl(new Go.GoAssign('value', iife))]]),
        ['gopurs/output/gopurs_runtime', 'sync']);
    const imports = collectImports([...declarations, [new Go.GoInitDecl(new Go.GoAssign('value', iife))]]);
    const directory = mkdtempSync(join(tmpdir(), 'gopurs-nullary-bindings-'));
    t.after(() => rmSync(directory, { recursive: true, force: true }));
    mkdirSync(join(directory, 'gopurs_runtime'));
    writeFileSync(join(directory, 'go.mod'), 'module gopurs/output\n\ngo 1.22\n');
    writeFileSync(join(directory, 'gopurs_runtime/runtime.go'), 'package gopurs_runtime\ntype Value struct{}\n');
    writeFileSync(join(directory, 'main.go'), `package main
import (
    "fmt"
    ${imports.map(name => JSON.stringify(name)).join('\n    ')}
)
type Constructor_Empty struct { Rc uint32 }
type Constructor_Generic[T any] struct { Rc uint32 }
func acceptEmpty(value *Constructor_Empty) bool { return value == nil }
func main() {
    ${printGoExpr(new Go.GoAssign('plain', empty))}
    ${printGoExpr(new Go.GoAssign('generic', generic))}
    _ = ${printGoExpr(iife)}
    _ = ${printGoExpr(discardIife)}
    fmt.Printf("%t %t %t", plain == nil, generic == nil, acceptEmpty(${printGoExpr(empty)}))
}
`);
    const result = spawnSync('go', ['run', '.'], {
        cwd: directory, encoding: 'utf8', timeout: 30_000,
        env: { ...process.env, GOWORK: 'off' },
    });
    assert.ifError(result.error);
    assert.equal(result.status, 0, result.stderr);
    assert.equal(result.stdout, 'true true true');
});

test('nullary constructors use their surrounding type without unnecessary imports', () => {
    const generic = new Go.GoConstructor('456', 'Constructor_Generic', [Go.TypeValue.value], []);
    assert.equal(printGoExpr(generic), 'nil');
    assert.equal(printGoExpr(new Go.GoReturn(generic)), 'return nil');
    const nilComparison = new Go.GoBinOp('==', new Go.GoVar('pointer'), generic);
    assert.deepEqual(collectImports([[new Go.GoFunctionDecl({
        name: 'Probe', params: [], result: Go.TypeBool.value, body: new Go.GoReturn(nilComparison),
    })]]), [], 'a nil comparison does not render the generic type argument');
});
