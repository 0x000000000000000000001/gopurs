import assert from 'node:assert/strict';
import { spawnSync } from 'node:child_process';
import { mkdirSync, mkdtempSync, rmSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join } from 'node:path';
import test from 'node:test';
import { Tuple } from '../output/Data.Tuple/index.js';
import * as Go from '../output/Gopurs.GoAst/index.js';
import { iterationBindings, loopParameters } from '../output/Gopurs.GoFunctions/index.js';
import { loop } from '../output/Gopurs.Int32Loops/index.js';
import { printGoDecl, printGoExpr } from '../output/Gopurs.Printer/index.js';
import { runtimeGoCode } from '../output/Gopurs.Runtime/index.js';
import { intAdd } from '../../gopurs-prelude/src/Data/Semiring.js';
import { intSub } from '../../gopurs-prelude/src/Data/Ring.js';
import { intDiv } from '../../gopurs-prelude/src/Data/EuclideanRing.js';
import { zshr } from '../../gopurs-integers/src/Data/Int/Bits.js';

const i64 = Go.TypeInt64.value;
const variable = name => new Go.GoVar(name);
const cast = (name, value) => new Go.GoCall(variable(name), [value]);
const integer = n => cast('int64', new Go.GoInt(n));
const call = (name, ...args) => new Go.GoCall(variable(name), args);
const runtime = (name, ...args) => new Go.GoCall(new Go.GoSelector(variable('gopurs_runtime'), name), args);
const binary = (op, left, right) => new Go.GoBinOp(op, left, right);
const n = variable('remaining'), seed = variable('seed'), step = variable('step');
const params = ['remaining', 'seed', 'step'].map(name => new Tuple(name, i64));
const assignSeed = expr => new Go.GoMutate('seed_loop', expr);
const again = new Go.GoContinue('again');
const body = update => [
    new Go.GoIfElse(binary('==', n, integer(0)), [new Go.GoReturn(seed)], []),
    new Go.GoMutate('remaining_loop', runtime('IntSub', n, integer(1))),
    ...update, again,
];
const original = statements => new Go.GoFor('again', [...iterationBindings(params), ...statements]);
const optimized = statements => loop('again')(params)(i64)(statements);

test('integer loop admission rejects wide assignments, opaque references and nested scopes', () => {
    const add = body([assignSeed(runtime('IntAdd', seed, step))]);
    const code = printGoExpr(optimized(add));
    assert.match(code, /remaining_loop := int32\(remaining_loop\)/);
    assert.match(code, /seed_loop := int32\(seed_loop\)/);
    assert.doesNotMatch(code, /step_loop := int32/);
    assert.match(code, /int64\(int32\(seed_loop\)\)/, 'guard the initial accumulator');
    assert.match(code, /return int64\(seed\)/, 'widen at the return boundary');

    for (const expr of [runtime('IntZshr', seed, step), runtime('IntDiv', seed, step), call('external', seed), integer(4294967295)]) {
        const mixed = printGoExpr(optimized(body([assignSeed(expr)])));
        assert.match(mixed, /remaining_loop := int32/);
        assert.doesNotMatch(mixed, /seed_loop := int32/);
    }
    const branches = body([new Go.GoIfElse(binary('<', seed, integer(0)),
        [assignSeed(runtime('IntAdd', seed, step))], [assignSeed(runtime('IntZshr', seed, step))])]);
    assert.doesNotMatch(printGoExpr(optimized(branches)), /seed_loop := int32/,
        'one wide branch invalidates the slot proof');

    for (const unsafe of [
        Go.rawGo('observe(remaining, seed)'),
        new Go.GoAssign('remaining', integer(1)),
        new Go.GoAssign('seed', integer(1)),
        new Go.GoAssign('remaining_loop', integer(1)),
        new Go.GoAssign('seed_loop', integer(1)),
        new Go.GoMutate('remaining', integer(1)),
        new Go.GoMutate('seed', integer(1)),
        call('observe', variable('remaining_loop'), variable('seed_loop')),
        call('save', new Go.GoFuncLit([], [], seed, i64)),
        new Go.GoFor('nested', [new Go.GoContinue('nested')]),
        new Go.GoContinue('outer'),
        new Go.GoRecordDict(new Go.TypeRecord([]), []),
    ]) {
        // Shadowing just one parameter can leave the other admissible. Require
        // that the affected parameter is never narrowed, or that the whole loop
        // is rejected when the AST form/scope is unknown.
        const statements = [unsafe, ...add];
        const rendered = printGoExpr(optimized(statements));
        if (unsafe instanceof Go.GoAssign || unsafe instanceof Go.GoMutate) {
            const name = unsafe.value0.replace(/_loop$/, '');
            assert.ok(!rendered.includes(`${name}_loop := int32`), rendered);
        } else assert.equal(rendered, printGoExpr(original(statements)));
    }
    const noBackedge = [new Go.GoReturn(seed)];
    assert.equal(printGoExpr(optimized(noBackedge)), printGoExpr(original(noBackedge)));
    assert.equal(printGoExpr(loop('again')(params)(Go.TypeValue.value)(add)), printGoExpr(original(add)));
});

test('native loops preserve signed overflow, wide entry/exit values and operand effects against JS', t => {
    const scenarios = [
        { name: 'Add', update: [assignSeed(runtime('IntAdd', seed, step))], next: (s, d) => intAdd(s)(d) },
        { name: 'Subtract', update: [assignSeed(runtime('IntSub', seed, step))], next: (s, d) => intSub(s)(d) },
        { name: 'Branch', update: [new Go.GoIfElse(binary('<', seed, integer(0)),
            [assignSeed(runtime('IntSub', seed, step))], [assignSeed(runtime('IntAdd', seed, step))])],
          next: (s, d) => s < 0 ? intSub(s)(d) : intAdd(s)(d) },
        { name: 'Unsigned', update: [assignSeed(runtime('IntZshr', seed, step))], next: (s, d) => zshr(s)(d) },
        { name: 'Division', update: [assignSeed(runtime('IntDiv', seed, step))], next: (s, d) => intDiv(s)(d) },
        { name: 'WideBranch', update: [new Go.GoIfElse(binary('<', seed, integer(0)),
            [assignSeed(runtime('IntAdd', seed, step))], [assignSeed(runtime('IntZshr', seed, step))])],
          next: (s, d) => s < 0 ? intAdd(s)(d) : zshr(s)(d) },
        { name: 'ConstantOverflow', update: [assignSeed(runtime('IntAdd', integer(2147483647), integer(1)))],
          next: () => intAdd(2147483647)(1) },
        { name: 'UnsignedLiteral', update: [assignSeed(integer(4294967295))], next: () => 4294967295 },
        { name: 'AddUnsignedLiteral', update: [assignSeed(runtime('IntAdd', seed, integer(4294967295)))],
          next: s => intAdd(s)(4294967295) },
        { name: 'CompareUnsignedLiteral', update: [new Go.GoIfElse(binary('<', seed, integer(4294967295)),
            [assignSeed(runtime('IntAdd', seed, integer(1)))], [assignSeed(runtime('IntSub', seed, integer(1)))])],
          next: s => s < 4294967295 ? intAdd(s)(1) : intSub(s)(1) },
        { name: 'Effects', update: [assignSeed(runtime('IntAdd', call('left', seed), call('right', step)))],
          next: (s, d) => intAdd(s)(d), effects: true },
        { name: 'Simultaneous', update: [assignSeed(runtime('IntAdd', seed, n))], next: (s, d, count) => intAdd(s)(count) },
        { name: 'WideComparison', update: [new Go.GoIfElse(binary('<', seed, step),
            [assignSeed(runtime('IntAdd', seed, integer(1)))], [assignSeed(runtime('IntSub', seed, integer(1)))])],
          next: (s, d) => s < d ? intAdd(s)(1) : intSub(s)(1) },
    ];
    const definitions = [];
    const cases = [];
    const seeds = [-2147483648, -2147483647, -1, 0, 1, 2147483647, 2147483648, 4294967295, 4294967296];
    const steps = [-2147483648, -3, -1, 0, 1, 2147483647, 2147483648, 4294967295];
    for (const scenario of scenarios) {
        const statements = body(scenario.update);
        for (const [prefix, value] of [['Before', original(statements)], ['After', optimized(statements)]]) {
            definitions.push(printGoDecl(new Go.GoFunctionDecl({
                name: prefix + scenario.name, params: loopParameters(params), result: i64, body: value,
            })));
        }
        for (const count of [0, 1, 2, 7, 4294967297, -4294967295]) {
            for (const initial of seeds) for (const delta of steps) {
                let remaining = count, value = initial, visits = 0;
                const seen = [];
                while (remaining !== 0) {
                    assert.ok(visits++ < 20, 'bounded oracle inputs');
                    if (scenario.effects) seen.push(value, delta);
                    value = scenario.next(value, delta, remaining);
                    remaining = intSub(remaining)(1);
                }
                cases.push(`{"${scenario.name}", ${count}, ${initial}, ${delta}, ${value}, `
                    + `"${scenario.effects ? 'LR'.repeat(visits) : ''}", []int64{${seen.join(',')}}}`);
            }
        }
    }
    const directory = process.env.GOPURS_INT32_TEST_OUTPUT
        ?? mkdtempSync(join(tmpdir(), 'gopurs-int32-loops-'));
    if (!process.env.GOPURS_INT32_TEST_OUTPUT) t.after(() => rmSync(directory, { recursive: true, force: true }));
    mkdirSync(join(directory, 'gopurs_runtime'), { recursive: true });
    writeFileSync(join(directory, 'go.mod'), 'module gopurs/output\n\ngo 1.22\n');
    writeFileSync(join(directory, 'gopurs_runtime/runtime.go'), runtimeGoCode);
    writeFileSync(join(directory, 'loops.go'), `package loops
import "gopurs/output/gopurs_runtime"
var order string
var seen []int64
func left(v int64) int64 { order += "L"; seen = append(seen, v); return v }
func right(v int64) int64 { order += "R"; seen = append(seen, v); return v }
${definitions.join('\n')}
`);
    writeFileSync(join(directory, 'loops_test.go'), `package loops
import ("testing"; "slices")
func TestLoopsAgainstJS(t *testing.T) {
    functions := map[string][2]func(int64,int64,int64)int64{
        ${scenarios.map(s => `"${s.name}": {Before${s.name}, After${s.name}},`).join('\n')}
    }
    for _, c := range []struct {
        name string; n, seed, step, want int64; order string; seen []int64
    }{${cases.join(',\n')}} {
        for path, f := range functions[c.name] {
            order = ""; seen = nil
            got := f(c.n, c.seed, c.step)
            if got != c.want || order != c.order || !slices.Equal(seen,c.seen) {
                t.Fatalf("%s path=%d n=%d seed=%d step=%d: got %d want %d; effects %q %v want %q %v",
                    c.name,path,c.n,c.seed,c.step,got,c.want,order,seen,c.order,c.seen)
            }
        }
    }
    t.Log("${cases.length} JS oracle cases, before and after")
}
`);
    const result = spawnSync('go', ['test', '-race', '-count=1', '-v', '.'], {
        cwd: directory, encoding: 'utf8', timeout: 120_000,
        env: { ...process.env, GOWORK: 'off' },
    });
    assert.ifError(result.error);
    assert.equal(result.status, 0, result.stdout + result.stderr);
    assert.match(result.stdout, /--- PASS: TestLoopsAgainstJS/);
    if (process.env.GOPURS_INT32_TEST_OUTPUT) writeFileSync(join(directory, 'go-test.log'), result.stdout + result.stderr);
});
