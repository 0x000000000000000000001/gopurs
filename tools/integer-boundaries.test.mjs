import assert from 'node:assert/strict';
import { spawnSync } from 'node:child_process';
import { mkdirSync, mkdtempSync, readFileSync, rmSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join } from 'node:path';
import test from 'node:test';
import { empty as emptySet } from '../output/Data.Set/index.js';
import * as Ref from '../output/Effect.Ref/index.js';
import * as Go from '../output/Gopurs.GoAst/index.js';
import { binary, unary } from '../output/Gopurs.PrimitiveExprs/index.js';
import { printGoExpr } from '../output/Gopurs.Printer/index.js';
import { runtimeGoCode } from '../output/Gopurs.Runtime/index.js';
import * as S from '../output/PureScript.Backend.Optimizer.Syntax/index.js';
import { intAdd, intMul } from '../../gopurs-prelude/src/Data/Semiring.js';
import { intSub } from '../../gopurs-prelude/src/Data/Ring.js';
import { intDiv, intMod } from '../../gopurs-prelude/src/Data/EuclideanRing.js';
import * as Bits from '../../gopurs-integers/src/Data/Int/Bits.js';

test('integer primitives and FFI match JavaScript at 32-bit boundaries', t => {
    const state = Ref.new({ reboxPairs: emptySet })();
    const operand = (name, variable) => ({
        expr: new Go.GoCall(new Go.GoVar(name), [new Go.GoVar(variable)]),
        exprType: Go.TypeInt64.value,
    });
    const emit = (op, arity) => {
        const translate = (arity === 1 ? unary : binary)(state)('Test')(op).value0;
        const first = translate(operand('left', 'a'));
        const result = arity === 1 ? first : first(operand('right', 'b'));
        return printGoExpr(result.exprType === Go.TypeValue.value ? new Go.GoSelector(result.expr, 'IntVal') : result.expr);
    };
    const values = [-2147483648, -2147483647, -65536, -46341, -32768, -1, 0, 1, 2, 31, 32,
        46341, 65536, 2147483646, 2147483647, 2147483648, 4294967295];
    const shifts = [-2147483648, -65, -33, -32, -31, -1, 0, 1, 30, 31, 32, 33, 63, 64, 65, 2147483647];
    const pairs = values.flatMap(a => values.map(b => [a, b]));
    const shiftPairs = values.flatMap(a => shifts.map(b => [a, b]));
    let seed = 0x2136;
    const next = () => (seed = (Math.imul(seed, 1664525) + 1013904223) | 0);
    for (let i = 0; i < 256; i++) {
        pairs.push([next(), next()]);
        shiftPairs.push([next(), next() % 96]);
    }
    const operations = [
        ['negate', S.OpIntNegate.value, 1, a => intSub(0)(a), 'purescript.IntSub(0, left(a))'],
        ['complement', S.OpIntBitNot.value, 1, Bits.complement, 'int64(purescript.Complement(int(left(a))))'],
        ['add', new S.OpIntNum(S.OpAdd.value), 2, intAdd, 'purescript.IntAdd(left(a), right(b))'],
        ['subtract', new S.OpIntNum(S.OpSubtract.value), 2, intSub, 'purescript.IntSub(left(a), right(b))'],
        ['multiply', new S.OpIntNum(S.OpMultiply.value), 2, intMul, 'purescript.IntMul(left(a), right(b))'],
        ['divide', new S.OpIntNum(S.OpDivide.value), 2, intDiv, 'purescript.IntDiv(left(a), right(b))'],
        ['modulo', new S.OpIntNum(S.OpMod.value), 2, intMod, 'purescript.IntMod(left(a), right(b))'],
        ['and', S.OpIntBitAnd.value, 2, Bits.and, 'int64(purescript.And(int(left(a)), int(right(b))))'],
        ['or', S.OpIntBitOr.value, 2, Bits.or, 'int64(purescript.Or(int(left(a)), int(right(b))))'],
        ['xor', S.OpIntBitXor.value, 2, Bits.xor, 'int64(purescript.Xor(int(left(a)), int(right(b))))'],
        ['shl', S.OpIntBitShiftLeft.value, 2, Bits.shl, 'int64(purescript.Shl(int(left(a)), int(right(b))))'],
        ['shr', S.OpIntBitShiftRight.value, 2, Bits.shr, 'int64(purescript.Shr(int(left(a)), int(right(b))))'],
        ['zshr', S.OpIntBitZeroFillShiftRight.value, 2, Bits.zshr, 'int64(purescript.Zshr(int(left(a)), int(right(b))))'],
    ];
    const directory = mkdtempSync(join(tmpdir(), 'gopurs-integer-boundaries-'));
    t.after(() => rmSync(directory, { recursive: true, force: true }));
    mkdirSync(join(directory, 'gopurs_runtime'));
    mkdirSync(join(directory, 'purescript'));
    writeFileSync(join(directory, 'go.mod'), 'module gopurs/output\n\ngo 1.22\n');
    writeFileSync(join(directory, 'gopurs_runtime/runtime.go'), runtimeGoCode);
    for (const [file, source] of [
        ['semiring', '../../gopurs-prelude/src/Data/Semiring.go'],
        ['ring', '../../gopurs-prelude/src/Data/Ring.go'],
        ['euclidean', '../../gopurs-prelude/src/Data/EuclideanRing.go'],
        ['bits', '../../gopurs-integers/src/Data/Int/Bits.go'],
    ]) writeFileSync(join(directory, `purescript/${file}.go`), 'package purescript\n' + readFileSync(new URL(source, import.meta.url), 'utf8'));
    writeFileSync(join(directory, 'boundaries_test.go'), `package boundaries_test
import ("testing"; "gopurs/output/gopurs_runtime"; "gopurs/output/purescript")
var calls string
func left(a int64) int64 { calls += "L"; return a }
func right(b int64) int64 { calls += "R"; return b }
func invoke(op func(int64,int64)int64, a,b int64) (value int64, panicked bool) {
    defer func() { if recover() != nil { panicked = true } }()
    return op(a,b), false
}
var _ = gopurs_runtime.Int
${operations.map(([name, op, arity, oracle, ffi]) => {
    const cases = arity === 1 ? values.map(a => [a, 0]) : name.endsWith('shr') || name === 'shl' ? shiftPairs : pairs;
    return `func Test_${name}(t *testing.T) {
        for path, operation := range map[string]func(int64,int64)int64{
            "generated": func(a,b int64)int64 { return ${emit(op, arity)} },
            "ffi": func(a,b int64)int64 { return ${ffi} },
        } {
            t.Run(path, func(t *testing.T) {
                for _, c := range [][3]int64{${cases.map(([a, b]) => `{${a},${b},${arity === 1 ? oracle(a) : oracle(a)(b)}}`).join(',')}} {
                    calls = ""
                    got, panicked := invoke(operation, c[0], c[1])
                    if panicked || got != c[2] || calls != "${arity === 1 ? 'L' : 'LR'}" {
                        t.Fatalf("%d, %d: got %d want %d; panic=%t calls=%s", c[0], c[1], got, c[2], panicked, calls)
                    }
                }
            })
        }
    }`;
}).join('\n')}
`);
    const result = spawnSync('go', ['test', '-race', '-count=1', '-v', '.'], {
        cwd: directory, encoding: 'utf8', timeout: 120_000, env: { ...process.env, GOWORK: 'off' },
    });
    assert.ifError(result.error);
    assert.equal(result.status, 0, result.stdout + result.stderr);
    for (const [name] of operations) assert.match(result.stdout, new RegExp(`--- PASS: Test_${name} `));
});
