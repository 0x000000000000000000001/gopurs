import assert from 'node:assert/strict';
import { spawnSync } from 'node:child_process';
import { mkdirSync, mkdtempSync, rmSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join } from 'node:path';
import test from 'node:test';
import { empty as emptyMap, insert } from '../output/Data.Map/index.js';
import { Just, Nothing } from '../output/Data.Maybe/index.js';
import { ordString } from '../output/Data.Ord/index.js';
import { empty as emptySet } from '../output/Data.Set/index.js';
import { Tuple } from '../output/Data.Tuple/index.js';
import * as CodeGen from '../output/Gopurs.CodeGen/index.js';
import { optimizeFunctionProducers } from '../output/Gopurs.FunctionFusion/index.js';
import { sanitizeName } from '../output/Gopurs.GoAst/index.js';
import { runtimeGoCode } from '../output/Gopurs.Runtime/index.js';
import * as C from '../output/PureScript.Backend.Optimizer.CoreFn/index.js';
import { NeutralExpr } from '../output/PureScript.Backend.Optimizer.Semantics/index.js';
import * as S from '../output/PureScript.Backend.Optimizer.Syntax/index.js';

const moduleName = 'CountingFixture';
const intType = C.Int.value;
const callbackType = new C.Func([intType], intType);
const iteratorType = new C.Func([callbackType, intType], intType);
const producerType = new C.Func([intType, callbackType, intType], intType);
const expr = syntax => new NeutralExpr(syntax);
const typed = (type, value) => expr(new S.Typed(type, value));
const local = (type, level) => typed(type, expr(new S.Local(Nothing.value, level)));
const integer = value => typed(intType, expr(new S.Lit(new C.LitInt(value))));
const variable = (name, module = moduleName) =>
    expr(new S.Var(new C.Qualified(module === null ? Nothing.value : new Just(module), name)));
const apply = (type, fn, ...args) => typed(type, expr(new S.App(fn, args)));
const lambda = (type, levels, body) => typed(type,
    expr(new S.Abs(levels.map(level => new Tuple(Nothing.value, level)), body)));
const binary = (type, operator, left, right) => typed(type,
    expr(new S.PrimOp(new S.Op2(operator, left, right))));
const identity = () => lambda(iteratorType, [10, 11], local(intType, 11));
const bindings = source => source.bindings.flatMap(group => group.bindings);
const binding = (source, name) => bindings(source).find(pair => pair.value0 === name).value1;
const strip = value => value instanceof S.Typed ? strip(value.value1)
    : value instanceof S.TypeApp ? strip(value.value0) : value;
const metadata = {
    elidedCtors: emptySet, ctorTypes: emptyMap, pointerAdtPaths: emptyMap,
    pointerAdtNodes: emptySet, pointerAdtLeaves: emptyMap, enumAdts: emptySet,
    enumCtors: emptySet, globalTypes: emptyMap, globalFunctions: emptyMap,
    classDeclsFields: emptyMap,
};

function fixture(options = {}) {
    const { name = 'repeat', base = variable('empty'), step = 1, reversed = false,
        recursive = true, self = variable(name), wrap = value => value,
        outerType = producerType, bodyWrap = value => value, annotation = true } = options;
    const n = local(intType, 0), f = local(callbackType, 2), x = local(intType, 3);
    const previous = options.previous ?? apply(intType, local(iteratorType, 1), f, x);
    const body = options.body ?? apply(intType, f, previous);
    const returned = options.returned ?? lambda(iteratorType, [2], lambda(callbackType, [3], body));
    const built = apply(iteratorType, typed(producerType, self),
        binary(intType, new S.OpIntNum(S.OpSubtract.value), n, integer(step)));
    const condition = binary(C.Boolean.value, new S.OpIntOrd(S.OpEq.value),
        ...(reversed ? [integer(0), n] : [n, integer(0)]));
    const bodyExpr = bodyWrap(typed(iteratorType, expr(new S.Branch([
        new S.Pair(condition, typed(iteratorType, base)),
    ], typed(iteratorType, expr(new S.Let(Nothing.value, 1, built, returned)))))));
    const abstraction = expr(new S.Abs([new Tuple(Nothing.value, 0)], bodyExpr));
    const producer = wrap(annotation ? typed(outerType, abstraction) : abstraction);
    return {
        name: moduleName,
        bindings: [
            { recursive: false, bindings: [new Tuple('empty', options.identity ?? identity())] },
            { recursive, bindings: [new Tuple(name, producer)] },
        ],
        comments: [], imports: emptySet, exports: emptySet, reExports: emptySet,
        dataTypes: emptyMap, dataDecls: [], classDecls: [], foreign: emptyMap,
        implementations: emptyMap, directives: emptyMap,
    };
}

function optimized(source) {
    const before = JSON.stringify(source);
    const result = optimizeFunctionProducers(source);
    assert.equal(JSON.stringify(source), before, 'the input IR must remain immutable');
    return result;
}

function workerNames(source, result) {
    const originalNames = new Set(bindings(source).map(pair => pair.value0));
    return bindings(result).map(pair => pair.value0).filter(name => !originalNames.has(name));
}

for (const [label, options] of [
    ['an unrelated producer name and reversed zero guard', { name: 'assemble', reversed: true }],
    ['grouped arrows and a TypeApp wrapper', {
        outerType: new C.Func([intType], iteratorType),
        wrap: value => expr(new S.TypeApp(value, intType)),
    }],
    ['an unqualified recursive name', { self: variable('repeat', null) }],
    ['an inline identity and flattened callback arguments', {
        base: identity(),
        returned: lambda(iteratorType, [2, 3], apply(intType, local(callbackType, 2),
            apply(intType, local(iteratorType, 1), local(callbackType, 2), local(intType, 3)))),
    }],
    ['a polymorphic identity', (() => {
        const typeVariable = new C.TypeVar('element');
        const polyType = new C.Func([new C.Func([typeVariable], typeVariable), typeVariable], typeVariable);
        return { identity: typed(new C.ForAll(['element'], polyType),
            lambda(polyType, [10, 11], local(typeVariable, 11))) };
    })()],
]) {
    test(`counted function fusion accepts ${label}`, () => {
        const source = fixture(options);
        const result = optimized(source);
        assert.equal(workerNames(source, result).length, 1, 'a single loop worker must be introduced');
        assert.notDeepEqual(binding(result, options.name ?? 'repeat'), binding(source, options.name ?? 'repeat'));
        assert.deepEqual(binding(result, 'empty'), binding(source, 'empty'));
    });
}

test('the negative counter path retains the original producer body', () => {
    const source = fixture();
    const result = optimized(source);
    const originalBody = strip(binding(source, 'repeat')).value1;
    const rewrittenBody = strip(strip(binding(result, 'repeat')).value1);
    assert.ok(rewrittenBody instanceof S.Branch);
    assert.equal(rewrittenBody.value0.length, 1);
    assert.deepEqual(rewrittenBody.value1, originalBody,
        'the divergent negative path must not be replaced with an identity iterator');
    const condition = strip(rewrittenBody.value0[0].value0);
    assert.ok(condition instanceof S.PrimOp);
    assert.ok(condition.value0 instanceof S.Op2);
    assert.deepEqual(condition.value0.value0, new S.OpIntOrd(S.OpGte.value));
    assert.deepEqual(condition.value0.value1, local(intType, 0));
    assert.deepEqual(condition.value0.value2, integer(0));
});

for (const [label, options, mutate = source => source] of [
    ['a non-recursive group', { recursive: false }],
    ['an absent producer type', { annotation: false }],
    ['a Number counter', { outerType: new C.Func([C.Number.value, callbackType, intType], intType) }],
    ['a decrement by two', { step: 2 }],
    ['an increment', { step: -1 }],
    ['a call to another function', { self: variable('different') }],
    ['a recursive name in another module', { self: variable('repeat', 'OtherModule') }],
    ['an identity name in another module', { base: variable('empty', 'OtherModule') }],
    ['an unknown identity name', { base: variable('unknown') }],
    ['a constant base result', { identity: lambda(iteratorType, [10, 11], integer(7)) }],
    ['a base that invokes the callback', { base: lambda(iteratorType, [10, 11],
        apply(intType, local(callbackType, 10), local(intType, 11))) }],
    ['a missing recursive application', { body: apply(intType, local(callbackType, 2), local(intType, 3)) }],
    ['a counter captured as the seed', { previous: apply(intType,
        local(iteratorType, 1), local(callbackType, 2), local(intType, 0)) }],
    ['a modified callback', { previous: apply(intType,
        local(iteratorType, 1), variable('otherCallback'), local(intType, 3)) }],
    ['a reversed callback composition', { body: apply(intType, local(iteratorType, 1),
        local(callbackType, 2), apply(intType, local(callbackType, 2), local(intType, 3))) }],
    ['an incompatible nested annotation', { previous: typed(C.Number.value,
        apply(intType, local(iteratorType, 1), local(callbackType, 2), local(intType, 3))) }],
    ['extra work while constructing the iterator', { bodyWrap: body => expr(new S.Let(Nothing.value, 9,
        apply(intType, variable('opaque'), integer(0)), body)) }],
    ['a group with multiple recursive bindings', {}, source => ({ ...source,
        bindings: [source.bindings[0], { recursive: true,
            bindings: [...source.bindings[1].bindings, new Tuple('mutual', integer(0))] }],
    })],
]) {
    test(`counted function fusion rejects ${label}`, () => {
        const source = mutate(fixture(options));
        assert.deepEqual(optimized(source), source,
            'an unsupported producer must retain its complete original implementation');
    });
}

for (const collision of ['binding', 'sanitized binding', 'foreign']) {
    test(`counted function worker avoids a ${collision} collision`, () => {
        const source = fixture(collision === 'sanitized binding' ? { name: 'repeat$' } : {});
        const [firstWorker] = workerNames(source, optimized(source));
        assert.ok(firstWorker);
        // Go escapes punctuation; this distinct source name deliberately
        // occupies the same generated identifier.
        const reserved = collision === 'sanitized binding'
            ? firstWorker.replaceAll('$', '_dollar_') : firstWorker;
        if (collision === 'sanitized binding') assert.notEqual(reserved, firstWorker);
        assert.equal(sanitizeName(reserved), sanitizeName(firstWorker));
        const userValue = integer(42);
        if (collision === 'foreign') {
            source.foreign = insert(ordString)(reserved)(intType)(emptyMap);
        } else {
            source.bindings.push({ recursive: false, bindings: [new Tuple(reserved, userValue)] });
        }
        const result = optimized(source);
        const names = workerNames(source, result);
        assert.equal(names.length, 1);
        assert.notEqual(sanitizeName(names[0]), sanitizeName(reserved));
        if (collision === 'foreign') assert.deepEqual(result.foreign, source.foreign);
        else assert.deepEqual(binding(result, reserved), userValue);
        const emitted = bindings(result).map(pair => sanitizeName(pair.value0));
        assert.equal(new Set(emitted).size, emitted.length);
    });
}

function runGenerated(t, body) {
    const source = fixture();
    const names = workerNames(source, optimized(source));
    assert.equal(names.length, 1,
        'the runtime checks must exercise the optimized producer');
    const code = CodeGen.translate(metadata)(source);
    assert.ok(code.includes(`func Call_${moduleName}_${sanitizeName(names[0])}(`),
        'the actual code generator must include the recognized loop worker');
    const directory = mkdtempSync(join(tmpdir(), 'gopurs-counted-functions-'));
    t.after(() => rmSync(directory, { recursive: true, force: true }));
    mkdirSync(join(directory, 'purescript'));
    mkdirSync(join(directory, 'gopurs_runtime'));
    writeFileSync(join(directory, 'go.mod'), 'module gopurs/output\n\ngo 1.22\n');
    writeFileSync(join(directory, 'gopurs_runtime/runtime.go'), runtimeGoCode);
    writeFileSync(join(directory, `purescript/${moduleName}.go`), code);
    writeFileSync(join(directory, 'main.go'), `package main
import (
    "fmt"
    "reflect"
    "gopurs/output/gopurs_runtime"
    "gopurs/output/purescript"
)
func check(got, want int64) {
    if got != want { panic(fmt.Sprintf("got %d; want %d", got, want)) }
}
func main() {
    _ = reflect.DeepEqual
${body}
    fmt.Println("counted iterator verified")
}
`);
    const result = spawnSync('go', ['run', '.'], {
        cwd: directory, encoding: 'utf8', timeout: 30_000,
        env: { ...process.env, GOWORK: 'off',
            GOCACHE: process.env.GOCACHE ?? join(tmpdir(), 'gopurs-counted-tests-go-cache') },
    });
    assert.ifError(result.error);
    assert.equal(result.status, 0, result.stdout + result.stderr);
    assert.equal(result.stdout, 'counted iterator verified\n');
}

test('generated iterators preserve zero, saved values, partial applications and callback order', t => {
    runGenerated(t, `
    for n := int64(0); n <= 12; n++ {
        saved := purescript.Call_CountingFixture_repeat(n)
        var observed []int64
        callback := gopurs_runtime.Func(func(value gopurs_runtime.Value) gopurs_runtime.Value {
            observed = append(observed, value.IntVal)
            return gopurs_runtime.Int(value.IntVal * 2 + 1)
        })
        partial := gopurs_runtime.Apply(saved, callback)
        if len(observed) != 0 { panic("a partial application evaluated its callback") }
        for _, seed := range []int64{2, -3, 7} {
            observed = nil
            var expected []int64
            want := seed
            for count := int64(0); count < n; count++ {
                expected = append(expected, want)
                want = want * 2 + 1
            }
            check(gopurs_runtime.Apply(partial, gopurs_runtime.Int(seed)).IntVal, want)
            if !reflect.DeepEqual(observed, expected) { panic("partial callback order changed") }
            observed = nil
            check(gopurs_runtime.Apply2(saved, callback, gopurs_runtime.Int(seed)).IntVal, want)
            if !reflect.DeepEqual(observed, expected) { panic("saved callback order changed") }
        }
        other := gopurs_runtime.Func(func(value gopurs_runtime.Value) gopurs_runtime.Value {
            return gopurs_runtime.Int(value.IntVal - 3)
        })
        check(gopurs_runtime.Apply2(saved, other, gopurs_runtime.Int(20)).IntVal, 20 - 3*n)
        // Compositions retain their higher-order meaning when both iterators
        // are reusable values, with the callback partially applied to one.
        inner := purescript.Call_CountingFixture_repeat(3)
        composed := gopurs_runtime.Apply(inner, other)
        check(gopurs_runtime.Apply2(saved, composed, gopurs_runtime.Int(20)).IntVal, 20 - 9*n)
    }
    never := gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
        panic("zero must not call the callback")
    })
    zero := purescript.Call_CountingFixture_repeat(0)
    check(gopurs_runtime.Apply2(zero, never, gopurs_runtime.Int(37)).IntVal, 37)
`);
});

test('generated iterators defer callback failure until the same invocation', t => {
    runGenerated(t, `
    calls := int64(0)
    callback := gopurs_runtime.Func(func(value gopurs_runtime.Value) gopurs_runtime.Value {
        calls++
        if calls == 3 { panic("third callback") }
        return gopurs_runtime.Int(value.IntVal + 4)
    })
    saved := purescript.Call_CountingFixture_repeat(5)
    partial := gopurs_runtime.Apply(saved, callback)
    check(calls, 0)
    func() {
        defer func() {
            if got := recover(); got != "third callback" { panic(fmt.Sprintf("unexpected panic: %v", got)) }
        }()
        gopurs_runtime.Apply(partial, gopurs_runtime.Int(10))
        panic("the callback failure disappeared")
    }()
    check(calls, 3)
    check(gopurs_runtime.Apply(partial, gopurs_runtime.Int(10)).IntVal, 30)
    check(calls, 8)
`);
});
