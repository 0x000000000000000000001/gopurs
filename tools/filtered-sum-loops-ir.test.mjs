import assert from 'node:assert/strict';
import { spawnSync } from 'node:child_process';
import { mkdirSync, mkdtempSync, rmSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join, resolve } from 'node:path';
import { fileURLToPath, pathToFileURL } from 'node:url';
import test from 'node:test';

const output = process.argv[2] ?? fileURLToPath(new URL('../output/', import.meta.url));
const load = name => import(pathToFileURL(resolve(output, name, 'index.js')));
const [Go, Filtered, Affine, C, S, T, Maybe, M, Ord, Tuples, Free, Printer] = await Promise.all([
    'Gopurs.GoAst', 'Gopurs.FilteredSumLoops', 'Gopurs.AffineLoops',
    'PureScript.Backend.Optimizer.CoreFn', 'PureScript.Backend.Optimizer.Syntax',
    'PureScript.Backend.Optimizer.Codegen.Tco', 'Data.Maybe', 'Data.Map',
    'Data.Ord', 'Data.Tuple', 'PureScript.Backend.Optimizer.FreeVars', 'Gopurs.Printer',
].map(load));
const { Just, Nothing } = Maybe, { Tuple } = Tuples;
const int = C.Int.value, bool = C.Boolean.value, i64 = Go.TypeInt64.value;
const expr = syntax => new T.TcoExpr(null, syntax);
const typed = (type, value) => expr(new S.Typed(type, value));
const local = (name, level) => typed(int, expr(new S.Local(new Just(name), level)));
const counter = local('remaining', 0), accumulator = local('seed', 1);
const counterName = Free.localId(new Just('remaining'))(0), accumulatorName = Free.localId(new Just('seed'))(1);
const params = [new Tuple(counterName, i64), new Tuple(accumulatorName, i64)];
const integer = n => typed(int, expr(new S.Lit(new C.LitInt(n))));
const variable = (name = 'sumSelected', owner = 'Probe') => expr(new S.Var(new C.Qualified(new Just(owner), name)));
const binary = (operator, left, right, type = int) => typed(type, expr(new S.PrimOp(new S.Op2(operator, left, right))));
const add = (left, right) => binary(new S.OpIntNum(S.OpAdd.value), left, right);
const subtract = (left, right) => binary(new S.OpIntNum(S.OpSubtract.value), left, right);
const modulo = (left, right) => binary(new S.OpIntNum(S.OpMod.value), left, right);
const equal = (left, right) => binary(new S.OpIntOrd(S.OpEq.value), left, right, bool);
const less = (left, right) => binary(new S.OpIntOrd(S.OpLt.value), left, right, bool);
const call = (target, args) => typed(int, expr(new S.App(target, args)));
const branch = (condition, yes, no) => typed(int, expr(new S.Branch([new S.Pair(condition, yes)], no)));
const sentinel = new Go.GoReturn(new Go.GoInt(-98));
const self = Affine.globalTarget('Probe')('sumSelected');

function fixture(options = {}) {
    const residue = integer(options.residue ?? 0);
    const remainder = options.remainder ?? modulo(counter, integer(options.modulus ?? 2));
    const predicate = options.predicate ?? (options.commute ? equal(residue, remainder) : equal(remainder, residue));
    const selected = options.selected ?? (options.commute ? add(counter, accumulator) : add(accumulator, counter));
    const update = options.update ?? branch(predicate, selected, options.skipped ?? accumulator);
    const decrement = options.decrement ?? subtract(counter, integer(1));
    const args = options.args ?? (options.reversed ? [update, decrement] : [decrement, update]);
    const target = options.target ?? variable();
    const step = options.step ?? (options.uncurried ? typed(int, expr(new S.UncurriedApp(target, args))) : call(target, args));
    return branch(options.condition ?? less(counter, integer(options.lower ?? 1)), options.seed ?? accumulator, step);
}

function optimize(body, { parameters = params, result = i64, target = self, fallback = sentinel, hook = false } = {}) {
    const before = JSON.stringify(body);
    const value = (hook ? Affine : Filtered).wrap(target)(parameters)(result)(body)(fallback);
    assert.equal(JSON.stringify(body), before, 'input IR remains immutable');
    return value;
}

test('filtered sums admit exact local/global recurrences and parameter/operator permutations', () => {
    for (const options of [
        {}, { modulus: 1 }, { lower: 7, modulus: 5, residue: 3 },
        { commute: true }, { uncurried: true }, { reversed: true },
        { lower: 2147483647, modulus: 2147483647, residue: 2147483646 },
    ]) {
        const result = optimize(fixture(options), { parameters: options.reversed ? [...params].reverse() : params, hook: true });
        assert.notEqual(result, sentinel);
        const code = Printer.printGoExpr(result);
        assert.match(code, /func\(upper int64, seed int64\) int64/);
        assert.match(code, /int64\(int32\(/);
        assert.match(code, /2147483647/);
        assert.doesNotMatch(code, /IntMul|IntMod|IntAdd/);
        assert.equal(result.value0.at(-1), sentinel, 'fallback is unchanged');
    }
    const localSelf = local('nested', 7), key = Free.localId(new Just('nested'))(7);
    const bound = M.insert(Ord.ordString)(key)({ name: 'private_worker', goType: i64 })(M.empty);
    assert.notEqual(optimize(fixture({ target: localSelf }), { target: Affine.localTarget(bound)('private_worker') }), sentinel);
    assert.equal(optimize(fixture({ target: localSelf }), { target: Affine.localTarget(bound)('other_worker') }), sentinel);
});

test('filtered sums reject changed state, calls, effects and unsupported predicates', () => {
    const options = [
        { target: variable('different') }, { target: variable('sumSelected', 'Foreign') },
        { decrement: subtract(counter, integer(2)) }, { decrement: add(counter, integer(1)) },
        { decrement: call(variable('observe'), [counter]) },
        { selected: add(accumulator, integer(1)) }, { selected: subtract(accumulator, counter) },
        { selected: add(add(accumulator, integer(0)), counter) },
        { selected: call(variable('observe'), [counter]) }, { selected: expr(new S.EffectPure(counter)) },
        { skipped: add(accumulator, integer(1)) }, { skipped: counter },
        { seed: counter }, { seed: add(accumulator, integer(0)) },
        { remainder: modulo(add(counter, integer(1)), integer(2)) },
        { remainder: modulo(counter, accumulator) },
        { predicate: binary(new S.OpIntOrd(S.OpNotEq.value), modulo(counter, integer(2)), integer(0), bool) },
        { predicate: call(variable('predicate'), [counter]) },
        { condition: less(integer(1), counter) }, { condition: equal(counter, integer(0)) },
        { args: [subtract(counter, integer(1)), accumulator, integer(1)] },
        { step: expr(new S.EffectPure(integer(1))) },
    ];
    for (const option of options) assert.equal(optimize(fixture(option)), sentinel);
    const effectBefore = expr(new S.Let(Nothing.value, 9, call(variable('observe'), [accumulator]), fixture()));
    assert.equal(optimize(effectBefore), sentinel);
    const extraBranch = expr(new S.Branch([new S.Pair(less(counter, integer(1)), accumulator),
        new S.Pair(equal(counter, integer(2)), integer(0))], fixture()));
    assert.equal(optimize(extraBranch), sentinel);
    assert.equal(optimize(fixture(), { parameters: [...params, new Tuple('extra', i64)] }), sentinel);
    assert.equal(optimize(fixture(), { parameters: [params[0], params[0]] }), sentinel);
    assert.equal(optimize(fixture(), { parameters: [new Tuple(counterName, Go.TypeFloat64.value), params[1]] }), sentinel);
    assert.equal(optimize(fixture(), { result: Go.TypeValue.value }), sentinel);
});

test('filtered sums reject unsupported literal bounds, moduli and noncanonical residues', () => {
    for (const lower of [-2147483648, -1, 0, 2147483648, 4294967295]) assert.equal(optimize(fixture({ lower })), sentinel);
    for (const modulus of [-2147483648, -2, 0, 2147483648, 4294967295]) assert.equal(optimize(fixture({ modulus })), sentinel);
    for (const residue of [-2147483648, -1, 2, 2147483648, 4294967295]) assert.equal(optimize(fixture({ residue })), sentinel);
});

test('generated filtered sums preserve Int32 sums, empty selections and wide-value fallbacks', t => {
    const directory = process.env.GOPURS_FILTERED_SUM_TEST_OUTPUT ?? mkdtempSync(join(tmpdir(), 'gopurs-filtered-sum-loops-'));
    if (!process.env.GOPURS_FILTERED_SUM_TEST_OUTPUT) t.after(() => rmSync(directory, { recursive: true, force: true }));
    mkdirSync(directory, { recursive: true });
    const definitions = [], cases = [], guards = [];
    const configurations = [
        { lower: 1, modulus: 2, residue: 0 }, { lower: 1, modulus: 1, residue: 0 },
        { lower: 7, modulus: 5, residue: 3 }, { lower: 1, modulus: 7, residue: 6 },
        { lower: 2147483639, modulus: 5, residue: 3 },
        { lower: 2147483647, modulus: 2147483647, residue: 2147483646 },
        { lower: 1, modulus: 2147483647, residue: 2147483646 },
    ];
    const seeds = [-2147483648, -1, 0, 1, 2147483647, 2147483648, 4294967295];
    const print = (name, body) => Printer.printGoDecl(new Go.GoFunctionDecl({
        name: { owner: 'generated/Probe', name },
        params: params.map(p => new Tuple(p.value0 + '_loop', p.value1)), result: i64, body,
    }));
    for (const [index, config] of configurations.entries()) {
        const { lower, modulus, residue } = config;
        const n = counterName + '_loop', a = accumulatorName + '_loop';
        const fallback = Go.rawGo(`for ${n} >= ${lower} { if ${n}%${modulus} == ${residue} { ${a}=int64(int32(${a}+${n})) }; ${n}=int64(int32(${n}-1)) }; return ${a}`);
        definitions.push(print('Run' + index, optimize(fixture(config), { fallback })));
        definitions.push(print('Guard' + index, optimize(fixture(config))));
        for (const count of [-2147483648, -1, 0, 1, 2, 6, 7, 8, 30, 900, 4294967296, 4294967297]) for (const seed of seeds) {
            let current = count, want = seed, iterations = 0;
            while (current >= lower) {
                assert.ok(iterations++ < 1000);
                if (current % modulus === residue) want = (want + current) | 0;
                current = (current - 1) | 0;
            }
            cases.push(`{Run${index}, ${count}, ${seed}, ${want}}`);
        }
        // The independent oracle sums the indices k in k*modulus+residue.
        // BigInt preserves products beyond 2^53, where IntMul would round.
        for (const count of [2147483646, 2147483647]) for (const seed of seeds.slice(0, 5)) {
            const lo = BigInt(lower), hi = BigInt(count), m = BigInt(modulus), r = BigInt(residue);
            const begin = lo > r ? (lo - r + m - 1n) / m : 0n;
            const end = hi >= r ? (hi - r) / m : -1n;
            const terms = end >= begin ? end - begin + 1n : 0n;
            const sum = m * (end * (end + 1n) - (begin - 1n) * begin) / 2n + terms * r;
            const want = terms > 0n ? BigInt.asIntN(32, BigInt(seed) + sum) : BigInt(seed);
            cases.push(`{Run${index}, ${count}, ${seed}, ${want}}`);
        }
        for (const [count, seed] of [[lower - 1, 0], [-1, 0], [-2147483648, 0],
            [2147483648, 0], [4294967297, 0], [lower, 2147483648], [lower, 4294967295]]) {
            guards.push(`{Guard${index}, ${count}, ${seed}, -98}`);
        }
    }
    writeFileSync(join(directory, 'go.mod'), 'module filteredsumprobe\n\ngo 1.22\n');
    writeFileSync(join(directory, 'filtered.go'), `package filteredsumprobe\n${definitions.join('\n')}`);
    writeFileSync(join(directory, 'filtered_test.go'), `package filteredsumprobe
import "testing"
func TestFiltered(t *testing.T) {
 for i,c := range []struct{f func(int64,int64)int64; n,a,want int64}{${[...cases, ...guards].join(',\n')}} {
  if got:=c.f(c.n,c.a); got!=c.want {t.Fatalf("case %d n=%d a=%d got=%d want=%d",i,c.n,c.a,got,c.want)}
 }
 t.Log("${cases.length} semantic cases and ${guards.length} guard refusals passed")
}
`);
    const run = spawnSync('go', ['test', '-count=1', '-v', '.'], { cwd: directory, encoding: 'utf8', timeout: 60000, env: { ...process.env, GOWORK: 'off' } });
    assert.ifError(run.error);
    assert.equal(run.status, 0, run.stdout + run.stderr);
    if (process.env.GOPURS_FILTERED_SUM_TEST_OUTPUT) writeFileSync(join(directory, 'go-test.log'), run.stdout + run.stderr);
});
