import assert from 'node:assert/strict';
import { spawnSync } from 'node:child_process';
import { mkdirSync, mkdtempSync, rmSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join, resolve } from 'node:path';
import { fileURLToPath, pathToFileURL } from 'node:url';
import test from 'node:test';

const output = process.argv[2] ?? fileURLToPath(new URL('../output/', import.meta.url));
const load = name => import(pathToFileURL(resolve(output, name, 'index.js')));
const [Go, A, C, S, T, Maybe, M, Ord, Tuples, Free, Printer] = await Promise.all([
    'Gopurs.GoAst', 'Gopurs.AffineLoops', 'PureScript.Backend.Optimizer.CoreFn',
    'PureScript.Backend.Optimizer.Syntax', 'PureScript.Backend.Optimizer.Codegen.Tco',
    'Data.Maybe', 'Data.Map', 'Data.Ord', 'Data.Tuple',
    'PureScript.Backend.Optimizer.FreeVars', 'Gopurs.Printer',
].map(load));
const { Just, Nothing } = Maybe, { Tuple } = Tuples;
const int = C.Int.value, i64 = Go.TypeInt64.value;
const expr = syntax => new T.TcoExpr(null, syntax);
const typed = (type, value) => expr(new S.Typed(type, value));
const local = (name, level) => typed(int, expr(new S.Local(new Just(name), level)));
const counter = local('remaining', 0), accumulator = local('seed', 1);
const counterName = Free.localId(new Just('remaining'))(0), accumulatorName = Free.localId(new Just('seed'))(1);
const params = [new Tuple(counterName, i64), new Tuple(accumulatorName, i64)];
const integer = n => typed(int, expr(new S.Lit(new C.LitInt(n))));
const variable = (name = 'repeat', owner = 'Probe') => expr(new S.Var(new C.Qualified(new Just(owner), name)));
const binary = (operator, left, right, type = int) => typed(type, expr(new S.PrimOp(new S.Op2(operator, left, right))));
const add = (left, right) => binary(new S.OpIntNum(S.OpAdd.value), left, right);
const subtract = (left, right) => binary(new S.OpIntNum(S.OpSubtract.value), left, right);
const condition = binary(new S.OpIntOrd(S.OpEq.value), counter, integer(0), C.Boolean.value);
const call = (target, args) => typed(int, expr(new S.App(target, args)));
const sentinel = new Go.GoReturn(new Go.GoInt(-98));
const self = A.globalTarget('Probe')('repeat');
function fixture(options = {}) {
    const conditionExpr = options.condition ?? condition;
    const decrement = options.decrement ?? subtract(counter, integer(1));
    const update = options.update ?? add(accumulator, integer(options.increment ?? 1));
    const step = options.step ?? call(options.target ?? variable(), [decrement, update]);
    return typed(int, expr(new S.Branch([new S.Pair(conditionExpr, options.seed ?? accumulator)], step)));
}
function optimize(body, parameters = params, result = i64, target = self, fallback = sentinel) {
    const before = JSON.stringify(body);
    const value = A.wrap(target)(parameters)(result)(body)(fallback);
    assert.equal(JSON.stringify(body), before, 'input IR remains immutable');
    return value;
}

test('affine admission proves the entire recurrence and the exact self target', () => {
    for (const increment of [-2147483648, -7, 0, 1, 7, 2147483647]) {
        const body = fixture({ increment });
        const result = optimize(body);
        assert.notEqual(result, sentinel);
        const code = Printer.printGoExpr(result);
        assert.match(code, /int64\(int32\(/);
        assert.match(code, /\*.*int64\(/);
        assert.doesNotMatch(code, /IntMul/);
        assert.match(code, /2147483647/);
        assert.equal(result.value0.at(-1), sentinel, 'fallback is unchanged');
    }
    const swapped = fixture({ condition: binary(new S.OpIntOrd(S.OpEq.value), integer(0), counter, C.Boolean.value),
        update: add(integer(-3), accumulator) });
    assert.notEqual(optimize(swapped), sentinel);
    const localSelf = local('nested', 7), localKey = Free.localId(new Just('nested'))(7);
    const bound = M.insert(Ord.ordString)(localKey)({ name: 'allocated_worker', goType: i64 })(M.empty);
    assert.notEqual(optimize(fixture({ target: localSelf }), params, i64, A.localTarget(bound)('allocated_worker')), sentinel);
    assert.equal(optimize(fixture({ target: localSelf }), params, i64, A.localTarget(bound)('another_worker')), sentinel);
});

test('affine admission rejects calls, effects, changing increments and altered recurrences', () => {
    const options = [
        { target: variable('different') }, { target: variable('repeat', 'Foreign') },
        { decrement: subtract(counter, integer(2)) }, { decrement: add(counter, integer(1)) },
        { update: add(accumulator, counter) }, { update: subtract(accumulator, integer(1)) },
        { update: call(variable('observe'), [accumulator]) },
        { update: add(accumulator, integer(4294967295)) },
        { seed: counter }, { seed: add(accumulator, integer(1)) },
        { condition: binary(new S.OpIntOrd(S.OpLt.value), counter, integer(0), C.Boolean.value) },
        { step: call(variable(), [subtract(counter, integer(1)), accumulator, integer(1)]) },
        { step: expr(new S.EffectPure(integer(1))) },
    ];
    for (const option of options) assert.equal(optimize(fixture(option)), sentinel);
    const effectBefore = expr(new S.Let(Nothing.value, 9, call(variable('observe'), [accumulator]), fixture()));
    assert.equal(optimize(effectBefore), sentinel);
    assert.equal(optimize(fixture(), [...params, new Tuple('extra', i64)]), sentinel);
    assert.equal(optimize(fixture(), [new Tuple(counterName, Go.TypeFloat64.value), params[1]]), sentinel);
    assert.equal(optimize(fixture(), params, Go.TypeValue.value), sentinel);
});

test('generated affine fast paths preserve Int addition and leave unsupported inputs to the loop', t => {
    const directory = process.env.GOPURS_AFFINE_TEST_OUTPUT ?? mkdtempSync(join(tmpdir(), 'gopurs-affine-loops-'));
    if (!process.env.GOPURS_AFFINE_TEST_OUTPUT) t.after(() => rmSync(directory, { recursive: true, force: true }));
    mkdirSync(directory, { recursive: true });
    const definitions = [], cases = [], guards = [];
    const increments = [-2147483648, -2147483647, -7, 0, 1, 7, 2147483647];
    const seeds = [-2147483648, -1, 0, 1, 2147483647, 2147483648, 4294967295];
    const print = (name, body) => Printer.printGoDecl(new Go.GoFunctionDecl({
        name: { owner: 'generated/Probe', name },
        params: params.map(p => new Tuple(p.value0 + '_loop', p.value1)), result: i64, body,
    }));
    for (const [index, increment] of increments.entries()) {
        const n = counterName + '_loop', a = accumulatorName + '_loop';
        const fallback = Go.rawGo(`for ${n} != 0 { ${n} = int64(int32(${n} - 1)); ${a} = int64(int32(${a} + int64(${increment}))) }; return ${a}`);
        definitions.push(print('Run' + index, optimize(fixture({ increment }), params, i64, self, fallback)));
        definitions.push(print('Guard' + index, optimize(fixture({ increment }))));
        for (const count of [0, 1, 2, 7, 1000, 4294967297, -4294967295]) for (const seed of seeds) {
            let n = count, value = seed, iterations = 0;
            while (n !== 0) { assert.ok(iterations++ < 1002); n = (n - 1) | 0; value = (value + increment) | 0; }
            cases.push(`{Run${index}, ${count}, ${seed}, ${value}}`);
        }
        // Large counts execute only the proven path. BigInt supplies an exact
        // integer oracle independent of the Go formula and JS Number IntMul.
        for (const count of [2147483646, 2147483647]) for (const seed of seeds.slice(0, 5)) {
            const want = BigInt.asIntN(32, BigInt(seed) + BigInt(count) * BigInt(increment));
            cases.push(`{Run${index}, ${count}, ${seed}, ${want}}`);
        }
        for (const [count, seed] of [[0, 0], [-1, 0], [-2147483648, 0], [2147483648, 0], [4294967297, 0], [1, 2147483648]]) {
            guards.push(`{Guard${index}, ${count}, ${seed}, -98}`);
        }
    }
    writeFileSync(join(directory, 'go.mod'), 'module affineprobe\n\ngo 1.22\n');
    writeFileSync(join(directory, 'affine.go'), `package affineprobe\n${definitions.join('\n')}`);
    writeFileSync(join(directory, 'affine_test.go'), `package affineprobe
import "testing"
func TestAffine(t *testing.T) {
 for i,c := range []struct{f func(int64,int64)int64; n,a,want int64}{${[...cases, ...guards].join(',\n')}} {
  if got:=c.f(c.n,c.a); got!=c.want {t.Fatalf("case %d n=%d a=%d got=%d want=%d",i,c.n,c.a,got,c.want)}
 }
 t.Log("${cases.length} semantic cases and ${guards.length} guard refusals passed")
}
`);
    const run = spawnSync('go', ['test', '-count=1', '-v', '.'], { cwd: directory, encoding: 'utf8', timeout: 60000, env: { ...process.env, GOWORK: 'off' } });
    assert.ifError(run.error);
    assert.equal(run.status, 0, run.stdout + run.stderr);
    if (process.env.GOPURS_AFFINE_TEST_OUTPUT) writeFileSync(join(directory, 'go-test.log'), run.stdout + run.stderr);
});
