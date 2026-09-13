import assert from 'node:assert/strict';
import { spawnSync } from 'node:child_process';
import { mkdirSync, mkdtempSync, rmSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join } from 'node:path';
import test from 'node:test';
import { Just } from '../output/Data.Maybe/index.js';
import { empty as emptySet } from '../output/Data.Set/index.js';
import * as Ref from '../output/Effect.Ref/index.js';
import * as Go from '../output/Gopurs.GoAst/index.js';
import { binary } from '../output/Gopurs.PrimitiveExprs/index.js';
import { printGoExpr } from '../output/Gopurs.Printer/index.js';
import { runtimeGoCode } from '../output/Gopurs.Runtime/index.js';
import * as S from '../output/PureScript.Backend.Optimizer.Syntax/index.js';

test('integer division and modulo obey Euclidean laws and evaluate operands once', t => {
    const state = Ref.new({ reboxPairs: emptySet })();
    const operand = (name, value) => ({
        expr: new Go.GoCall(new Go.GoVar(name), [new Go.GoVar(value)]),
        exprType: Go.TypeInt64.value,
    });
    const emit = operator => {
        const emitter = binary(state)('Test')(new S.OpIntNum(operator));
        assert.ok(emitter instanceof Just);
        const result = emitter.value0(operand('left', 'a'))(operand('right', 'b'));
        assert.deepEqual(result.exprType, Go.TypeInt64.value);
        return printGoExpr(result.expr);
    };
    const division = emit(S.OpDivide.value);
    const modulo = emit(S.OpMod.value);
    const directory = mkdtempSync(join(tmpdir(), 'gopurs-integer-division-'));
    t.after(() => rmSync(directory, { recursive: true, force: true }));
    mkdirSync(join(directory, 'gopurs_runtime'));
    writeFileSync(join(directory, 'go.mod'), 'module gopurs/output\n\ngo 1.22\n');
    writeFileSync(join(directory, 'gopurs_runtime/runtime.go'), runtimeGoCode);
    writeFileSync(join(directory, 'main.go'), `package main
import (
    "fmt"
    "os"
    "gopurs/output/gopurs_runtime"
)
var leftCalls, rightCalls int
func left(value int64) int64 { leftCalls++; return value }
func right(value int64) int64 { rightCalls++; return value }
func divide(a, b int64) int64 { return ${division} }
func modulo(a, b int64) int64 { return ${modulo} }
func evaluate(operation func(int64, int64) int64, a, b int64) (result int64, panicked bool) {
    defer func() { if recover() != nil { panicked = true } }()
    return operation(a, b), false
}
func main() {
    _ = gopurs_runtime.Int
    cases := [][4]int64{
        {7, 3, 2, 1}, {-7, 3, -3, 2}, {7, -3, -2, 1}, {-7, -3, 3, 2},
        {9, 3, 3, 0}, {-9, 3, -3, 0}, {9, -3, -3, 0}, {-9, -3, 3, 0},
        {0, 3, 0, 0}, {0, -3, 0, 0}, {0, 0, 0, 0}, {7, 0, 0, 0}, {-7, 0, 0, 0},
        {-2147483648, -1, 2147483648, 0}, {-2147483648, 1, -2147483648, 0},
        {2147483647, -1, -2147483647, 0}, {2147483647, 1, 2147483647, 0},
        {-2147483648, 3, -715827883, 1}, {-2147483648, -3, 715827883, 1},
        {2147483647, 3, 715827882, 1}, {2147483647, -3, -715827882, 1},
        {-2147483648, 2147483647, -2, 2147483646},
        {2147483647, -2147483648, 0, 2147483647},
        {-2147483648, -2147483648, 1, 0},
    }
    failures := 0
    for _, c := range cases {
        a, b, expectedQuotient, expectedRemainder := c[0], c[1], c[2], c[3]
        values := [2]int64{}
        for index, operation := range []func(int64, int64) int64{divide, modulo} {
            leftCalls, rightCalls = 0, 0
            value, panicked := evaluate(operation, a, b)
            values[index] = value
            if panicked || leftCalls != 1 || rightCalls != 1 {
                fmt.Printf("operation %d (%d,%d): panic=%t operand calls=%d,%d\\n",
                    index, a, b, panicked, leftCalls, rightCalls)
                failures++
            }
        }
        q, r := values[0], values[1]
        if q != expectedQuotient || r != expectedRemainder {
            fmt.Printf("(%d,%d): quotient=%d remainder=%d; expected %d,%d\\n",
                a, b, q, r, expectedQuotient, expectedRemainder)
            failures++
        }
        if b != 0 {
            magnitude := b
            if magnitude < 0 { magnitude = -magnitude }
            if a != b*q+r || r < 0 || r >= magnitude {
                fmt.Printf("(%d,%d): Euclidean law failed for quotient=%d remainder=%d\\n", a, b, q, r)
                failures++
            }
        }
    }
    if failures != 0 { os.Exit(1) }
    fmt.Printf("%d cases passed", len(cases))
}
`);
    const result = spawnSync('go', ['run', '.'], {
        cwd: directory, encoding: 'utf8', timeout: 30_000,
        env: { ...process.env, GOWORK: 'off' },
    });
    assert.ifError(result.error);
    assert.equal(result.status, 0, result.stdout + result.stderr);
    assert.equal(result.stdout, '24 cases passed');
});
