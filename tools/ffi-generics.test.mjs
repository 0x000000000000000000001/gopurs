import assert from 'node:assert/strict';
import { spawnSync } from 'node:child_process';
import { mkdirSync, mkdtempSync, readFileSync, rmSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join } from 'node:path';
import test from 'node:test';
import { fileURLToPath } from 'node:url';
import { Nothing } from '../output/Data.Maybe/index.js';
import { Tuple } from '../output/Data.Tuple/index.js';
import { generateFfiBridge } from '../output/Gopurs.FfiBridge/index.js';
import { extractFfiDecls, prepareFfi } from '../output/Gopurs.FfiSupport/index.js';

const root = fileURLToPath(new URL('../', import.meta.url));
const ordPath = fileURLToPath(new URL('../../gopurs-prelude/src/Data/Ord.go', import.meta.url));

function prepare(moduleName, path, content, names) {
    const source = { moduleName, path };
    const prefix = moduleName.replaceAll('.', '_');
    const prepared = prepareFfi(source)(prefix + '_')(content)();
    const foreigns = names.map(name => new Tuple(name, Nothing.value));
    return {
        declarations: extractFfiDecls(source)(content)(),
        code: prepared.content + '\n' + generateFfiBridge(prefix)([])(prepared.decls)(foreigns),
    };
}

test('generic scalar FFI preserves native ordering and avoids boxing Ordering into interfaces', t => {
    // Read the production FFI rather than duplicating the optimized signatures.
    const ord = prepare('Data.Ord', ordPath, readFileSync(ordPath, 'utf8'), ['ordCharImpl', 'ordStringImpl']);
    for (const name of ['OrdCharImpl', 'OrdStringImpl']) {
        const declaration = ord.declarations.find(decl => decl.name === name);
        assert.deepEqual(declaration.typeParams, ['T']);
        assert.match(ord.code, new RegExp(`Data_Ord_${name}\\[gopurs_runtime\\.Value\\]\\(`));
    }

    const fixture = prepare('Fixture', 'Fixture.go', `
func ReturnSecond[A, B any](first A, second B) B { return second }
func Concrete(value int64) int64 { return value + 1 }
`, ['returnSecond', 'concrete']);
    assert.deepEqual(fixture.declarations.find(decl => decl.name === 'ReturnSecond').typeParams, ['A', 'B']);
    assert.match(fixture.code, /Fixture_ReturnSecond\[gopurs_runtime\.Value, gopurs_runtime\.Value\]\(/);
    assert.match(fixture.code, /gopurs_runtime\.Unbox\[int64\]\(arg0\)/);
    assert.doesNotMatch(ord.code + fixture.code, /gopurs_runtime\.Unbox\[(?:T|A|B)\]/);

    const directory = mkdtempSync(join(tmpdir(), 'gopurs-ffi-generics-'));
    t.after(() => rmSync(directory, { recursive: true, force: true }));
    mkdirSync(join(directory, 'gopurs_runtime'));
    writeFileSync(join(directory, 'go.mod'), 'module gopurs/output\n\ngo 1.22\n');
    writeFileSync(join(directory, 'gopurs_runtime/runtime.go'), readFileSync(join(root, 'runtime/runtime.go')));
    const header = 'package main\nimport "gopurs/output/gopurs_runtime"\n';
    writeFileSync(join(directory, 'ord.go'), header + ord.code);
    writeFileSync(join(directory, 'fixture.go'), header + fixture.code);
    writeFileSync(join(directory, 'ffi_test.go'), `package main

import (
    "strings"
    "testing"
    "gopurs/output/gopurs_runtime"
)

var lt = gopurs_runtime.Value{Type: gopurs_runtime.TypeConstructor, IntVal: 1527465420}
var eq = gopurs_runtime.Value{Type: gopurs_runtime.TypeConstructor, IntVal: 902936544}
var gt = gopurs_runtime.Value{Type: gopurs_runtime.TypeConstructor, IntVal: 380165415}
var sink gopurs_runtime.Value

type pair struct { x, y string }
var charPairs = []pair{
    {"a", "z"}, {"a", "a"}, {"z", "a"},
    {"é", "ê"}, {"é", "é"}, {"ê", "é"},
    {"😀", "😁"}, {"😀", "😀"}, {"😁", "😀"},
    {"\\x00", "a"}, {"\\ue000", "😀"}, {"😀", "\\ue000"},
}
var stringPairs = []pair{
    {"alpha", "beta"}, {"same", "same"}, {"beta", "alpha"},
    {"", "x"}, {"", ""}, {"x", ""},
    {"abc", "abcd"}, {"abcd", "abc"},
    {"éclair", "étude"}, {"étude", "éclair"}, {"日本語", "日本語"},
    {"\\ue000", "😀"}, {"😀", "\\ue000"},
    {strings.Repeat("prefix", 32) + "a", strings.Repeat("prefix", 32) + "b"},
    {strings.Repeat("prefix", 32) + "b", strings.Repeat("prefix", 32) + "a"},
    {strings.Repeat("same", 32), strings.Repeat("same", 32)},
}
type suite struct {
    name string
    fn gopurs_runtime.Value
    pairs []pair
}
func suites() []suite {
    return []suite{
        {"Char", _Gopurs_Data_Ord_OrdCharImpl, charPairs},
        {"String", _Gopurs_Data_Ord_OrdStringImpl, stringPairs},
    }
}

func TestNativeOrderingAndPartialApplication(t *testing.T) {
    for _, s := range suites() {
        t.Run(s.name, func(t *testing.T) {
            partial := gopurs_runtime.Apply3(s.fn, lt, eq, gt)
            for _, p := range s.pairs {
                // Preserve the existing Go ordering, including Unicode cases;
                // this does not redefine ordering for other backends.
                want := eq
                if cmp := strings.Compare(p.x, p.y); cmp < 0 { want = lt } else if cmp > 0 { want = gt }
                x, y := gopurs_runtime.Str(p.x), gopurs_runtime.Str(p.y)
                direct := gopurs_runtime.Apply5(s.fn, lt, eq, gt, x, y)
                curried := gopurs_runtime.Apply2(partial, x, y)
                if direct != want || curried != want {
                    t.Fatalf("(%q,%q): direct=%+v partial=%+v want=%+v", p.x, p.y, direct, curried, want)
                }
            }
        })
    }
}

func TestOrderingAllocations(t *testing.T) {
    for _, s := range suites() {
        t.Run(s.name, func(t *testing.T) {
            type boxedPair struct { x, y gopurs_runtime.Value }
            inputs := make([]boxedPair, len(s.pairs))
            for i, p := range s.pairs { inputs[i] = boxedPair{gopurs_runtime.Str(p.x), gopurs_runtime.Str(p.y)} }
            fn, index := s.fn, 0
            preboxed := testing.AllocsPerRun(1000, func() {
                p := inputs[index % len(inputs)]
                index++
                sink = gopurs_runtime.Apply5(fn, lt, eq, gt, p.x, p.y)
            })
            if preboxed != 0 { t.Fatalf("preboxed: got %g allocations, want 0", preboxed) }
            boxed := testing.AllocsPerRun(1000, func() {
                p := s.pairs[index % len(s.pairs)]
                index++
                sink = gopurs_runtime.Apply5(fn, lt, eq, gt, gopurs_runtime.Str(p.x), gopurs_runtime.Str(p.y))
            })
            if boxed != 0 { t.Fatalf("Str per call: got %g allocations, want 0 (packed strings)", boxed) }
        })
    }
}

func TestMultipleTypeParametersAndConcreteArgument(t *testing.T) {
    first, second := gopurs_runtime.Str("first"), gopurs_runtime.Int(42)
    if got := gopurs_runtime.Apply2(_Gopurs_Fixture_ReturnSecond, first, second); got != second {
        t.Fatalf("second generic argument: got %+v, want %+v", got, second)
    }
    if got := gopurs_runtime.Apply(_Gopurs_Fixture_Concrete, gopurs_runtime.Int(41)); got != second {
        t.Fatalf("concrete int64 argument: got %+v, want %+v", got, second)
    }
}
`);
    const result = spawnSync('go', ['test', '-count=1', '-v', './...'], {
        cwd: directory, encoding: 'utf8', timeout: 60_000,
        env: { ...process.env, GOWORK: 'off' }, maxBuffer: 1024 * 1024,
    });
    assert.ifError(result.error);
    assert.equal(result.signal, null);
    assert.equal(result.status, 0, result.stdout + result.stderr);
    assert.match(result.stdout, /--- PASS: TestOrderingAllocations/);
    assert.match(result.stdout, /--- PASS: TestNativeOrderingAndPartialApplication/);
    assert.match(result.stdout, /--- PASS: TestMultipleTypeParametersAndConcreteArgument/);
});

test('generic array callbacks preserve short circuiting, values and input arrays', t => {
    const arrayPath = fileURLToPath(new URL('../../gopurs-arrays/src/Data/Array.go', import.meta.url));
    const arrays = prepare('Data.Array', arrayPath, readFileSync(arrayPath, 'utf8'), ['anyImpl']);
    const anyDeclaration = arrays.declarations.find(decl => decl.name === 'AnyImpl');
    assert.equal(anyDeclaration.typeParams.length, 1);
    assert.match(arrays.code, /Data_Array_AnyImpl\[gopurs_runtime\.Value\]\(/);

    const fixture = prepare('ArrayFixture', 'ArrayFixture.go', `
func Transform[A, B any](f func(A) B, xs []A) []B {
    result := make([]B, len(xs))
    for i, x := range xs { result[i] = f(x) }
    return result
}
func ReturnCallback[A, B any](f func(A) B) func(A) B { return f }
`, ['transform', 'returnCallback']);
    assert.doesNotMatch(fixture.code, /gopurs_runtime\.Unbox\[(?:A|B)\]/);

    const directory = mkdtempSync(join(tmpdir(), 'gopurs-ffi-array-generics-'));
    t.after(() => rmSync(directory, { recursive: true, force: true }));
    mkdirSync(join(directory, 'gopurs_runtime'));
    writeFileSync(join(directory, 'go.mod'), 'module gopurs/output\n\ngo 1.22\n');
    writeFileSync(join(directory, 'gopurs_runtime/runtime.go'), readFileSync(join(root, 'runtime/runtime.go')));
    const header = 'package main\nimport "gopurs/output/gopurs_runtime"\n';
    writeFileSync(join(directory, 'array.go'), header + arrays.code);
    writeFileSync(join(directory, 'fixture.go'), header + fixture.code);
    writeFileSync(join(directory, 'ffi_test.go'), `package main

import (
    "reflect"
    "testing"
    rt "gopurs/output/gopurs_runtime"
)

var arraySink rt.Value

func TestAnyOrderShortCircuitAndImmutability(t *testing.T) {
    for _, tc := range []struct { name string; values []int64; match int64; want bool; visited []int64 }{
        {"empty", []int64{}, 1, false, []int64{}},
        {"middle", []int64{1, 2, 3}, 2, true, []int64{1, 2}},
        {"first", []int64{1, 2, 3}, 1, true, []int64{1}},
        {"absent", []int64{1, 2, 3}, 4, false, []int64{1, 2, 3}},
    } {
        t.Run(tc.name, func(t *testing.T) {
            values := make([]rt.Value, len(tc.values))
            for i, x := range tc.values { values[i] = rt.Int(x) }
            original := append([]rt.Value{}, values...)
            visited := []int64{}
            predicate := rt.Func(func(x rt.Value) rt.Value {
                visited = append(visited, x.IntVal)
                return rt.Bool(x.IntVal == tc.match)
            })
            got := rt.Apply2(_Gopurs_Data_Array_AnyImpl, predicate, rt.Array(values))
            if (got.IntVal != 0) != tc.want { t.Fatalf("any = %v, want %v", got, tc.want) }
            if !reflect.DeepEqual(visited, tc.visited) { t.Fatalf("visits = %v, want %v", visited, tc.visited) }
            if !reflect.DeepEqual(values, original) { t.Fatalf("input changed: %v -> %v", original, values) }
        })
    }
}

func TestAnyAllocationsDoNotGrowWithArrayLength(t *testing.T) {
    // Prebox the inputs and predicate outside measurement. A match at the first
    // element must not allocate a conversion of all remaining array elements.
    predicate := rt.Func(func(rt.Value) rt.Value { return rt.Bool(true) })
    count := func(size int) float64 {
        values := make([]rt.Value, size)
        for i := range values { values[i] = rt.Int(int64(i)) }
        boxed := rt.Array(values)
        return testing.AllocsPerRun(100, func() {
            arraySink = rt.Apply2(_Gopurs_Data_Array_AnyImpl, predicate, boxed)
        })
    }
    small, large := count(1), count(1024)
    if large > small { t.Fatalf("allocations grow with array length: 1=%g, 1024=%g", small, large) }
}

func TestGenericCallbackAndArrayResults(t *testing.T) {
    input := []rt.Value{rt.Str("alpha"), rt.Str("beta")}
    original := append([]rt.Value{}, input...)
    first := rt.Value{Type: rt.TypeConstructor, IntVal: 12345}
    second := rt.Value{Type: rt.TypeConstructor, IntVal: 67890}
    visited := []string{}
    callback := rt.Func(func(value rt.Value) rt.Value {
        text := rt.Unbox[string](value)
        visited = append(visited, text)
        if text == "alpha" { return first }
        return second
    })
    result := rt.Apply2(_Gopurs_ArrayFixture_Transform, callback, rt.Array(input))
    if got := rt.Unbox[[]rt.Value](result); !reflect.DeepEqual(got, []rt.Value{first, second}) {
        t.Fatalf("generic callback results changed: %v", got)
    }
    if !reflect.DeepEqual(visited, []string{"alpha", "beta"}) { t.Fatalf("callback order: %v", visited) }
    if !reflect.DeepEqual(input, original) { t.Fatalf("input changed: %v -> %v", original, input) }
    returned := rt.Apply(_Gopurs_ArrayFixture_ReturnCallback, callback)
    if got := rt.Apply(returned, input[1]); got != second { t.Fatalf("returned generic callback: %v", got) }
}
`);
    const result = spawnSync('go', ['test', '-count=1', '-v', './...'], {
        cwd: directory, encoding: 'utf8', timeout: 60_000,
        env: { ...process.env, GOWORK: 'off' }, maxBuffer: 1024 * 1024,
    });
    assert.ifError(result.error);
    assert.equal(result.signal, null);
    assert.equal(result.status, 0, result.stdout + result.stderr);
    assert.match(result.stdout, /--- PASS: TestAnyOrderShortCircuitAndImmutability/);
    assert.match(result.stdout, /--- PASS: TestAnyAllocationsDoNotGrowWithArrayLength/);
    assert.match(result.stdout, /--- PASS: TestGenericCallbackAndArrayResults/);
});

test('Value FFI callbacks and traversal preserve Applicative values and evaluation order', t => {
    const traversePath = fileURLToPath(new URL('../../gopurs-foldable-traversable/src/Data/Traversable.go', import.meta.url));
    const traversal = prepare('Data.Traversable', traversePath, readFileSync(traversePath, 'utf8'), ['traverseArrayImpl']);
    const fixture = prepare('ValueFixture', 'ValueFixture.go', `
import "gopurs/output/gopurs_runtime"
var Calls int64
func HigherOrder(f func(func(gopurs_runtime.Value) gopurs_runtime.Value, gopurs_runtime.Value) gopurs_runtime.Value, value gopurs_runtime.Value) gopurs_runtime.Value {
    return f(func(x gopurs_runtime.Value) gopurs_runtime.Value {
        Calls++
        return gopurs_runtime.Int(x.IntVal + 1)
    }, value)
}
func Deferred() func(gopurs_runtime.Value) gopurs_runtime.Value {
    return func(x gopurs_runtime.Value) gopurs_runtime.Value {
        Calls++
        return x
    }
}
`, ['higherOrder', 'deferred']);
    assert.match(fixture.code, /gopurs_runtime\.Func\(p0_0\)/);
    assert.match(fixture.code, /return gopurs_runtime\.Func\(go_res\)/);
    assert.doesNotMatch(fixture.code + traversal.code, /gopurs_runtime\.Box\(|make\(\[\](?:any|interface\{\})/);

    const directory = mkdtempSync(join(tmpdir(), 'gopurs-ffi-value-traverse-'));
    t.after(() => rmSync(directory, { recursive: true, force: true }));
    mkdirSync(join(directory, 'gopurs_runtime'));
    writeFileSync(join(directory, 'go.mod'), 'module gopurs/output\n\ngo 1.22\n');
    writeFileSync(join(directory, 'gopurs_runtime/runtime.go'), readFileSync(join(root, 'runtime/runtime.go')));
    // These production/fixture sources supply their own runtime import.
    writeFileSync(join(directory, 'traversal.go'), 'package main\n' + traversal.code);
    writeFileSync(join(directory, 'fixture.go'), 'package main\n' + fixture.code);
    writeFileSync(join(directory, 'ffi_test.go'), `package main
import (
    "reflect"
    "testing"
    rt "gopurs/output/gopurs_runtime"
)

func TestHigherOrderValueCallbacks(t *testing.T) {
    ValueFixture_Calls = 0
    deferred := rt.Apply(_Gopurs_ValueFixture_Deferred, rt.Value{})
    if ValueFixture_Calls != 0 { t.Fatal("returned callback ran during conversion") }
    value := rt.Str("same boxed value")
    if got := rt.Apply(deferred, value); got != value || ValueFixture_Calls != 1 {
        t.Fatalf("callback lost its value or effects: %+v, calls=%d", got, ValueFixture_Calls)
    }
    higher := rt.Func2(func(f, x rt.Value) rt.Value {
        return rt.Apply(f, rt.Apply(f, x))
    })
    if got := rt.Apply2(_Gopurs_ValueFixture_HigherOrder, higher, rt.Int(40)); got.IntVal != 42 || ValueFixture_Calls != 3 {
        t.Fatalf("higher-order callback: %+v, calls=%d", got, ValueFixture_Calls)
    }
}

func values(n int) []rt.Value {
    out := make([]rt.Value, n)
    for i := range out { out[i] = rt.Int(int64(i)) }
    return out
}
func array(v rt.Value) []rt.Value { return rt.Unbox[[]rt.Value](v) }
func traverse(apply, mapFn, pure, f rt.Value, input []rt.Value) rt.Value {
    concat := rt.Func2(func(a, b rt.Value) rt.Value {
        out := append([]rt.Value{}, array(a)...)
        return rt.Array(append(out, array(b)...))
    })
    return rt.Apply6(_Gopurs_Data_Traversable_TraverseArrayImpl, apply, mapFn, pure, concat, f, rt.Array(input))
}

type either struct { failed bool; value rt.Value }
func TestTraversalEitherAndStrictCallbackOrder(t *testing.T) {
    pure := rt.Func(func(x rt.Value) rt.Value { return rt.Any(either{value: x}) })
    mapFn := rt.Func2(func(f, x rt.Value) rt.Value {
        e := x.PtrVal().(either)
        if e.failed { return x }
        return rt.Any(either{value: rt.Apply(f, e.value)})
    })
    apply := rt.Func2(func(f, x rt.Value) rt.Value {
        ef, ex := f.PtrVal().(either), x.PtrVal().(either)
        if ef.failed { return f }; if ex.failed { return x }
        return rt.Any(either{value: rt.Apply(ef.value, ex.value)})
    })
    for _, n := range []int{0, 1, 2, 3, 4, 7, 8, 9} {
        input := values(n)
        for fail := -1; fail < n; fail++ {
            visited := []int64{}
            f := rt.Func(func(x rt.Value) rt.Value {
                visited = append(visited, x.IntVal)
                // Multiple failures must still retain the first error.
                return rt.Any(either{failed: fail >= 0 && x.IntVal >= int64(fail), value: x})
            })
            got := traverse(apply, mapFn, pure, f, input).PtrVal().(either)
            wantVisits := make([]int64, n)
            for i := range wantVisits { wantVisits[i] = int64(i) }
            if !reflect.DeepEqual(visited, wantVisits) { t.Fatalf("n=%d fail=%d visits=%v", n, fail, visited) }
            if fail >= 0 {
                if !got.failed || got.value.IntVal != int64(fail) { t.Fatalf("first error changed: %+v", got) }
            } else if got.failed || !reflect.DeepEqual(array(got.value), input) { t.Fatalf("success changed: %+v", got) }
            if !reflect.DeepEqual(input, values(n)) { t.Fatal("input mutated") }
        }
    }
}

func TestTraversalArrayCartesianOrder(t *testing.T) {
    pure := rt.Func(func(x rt.Value) rt.Value { return rt.Array([]rt.Value{x}) })
    mapFn := rt.Func2(func(f, x rt.Value) rt.Value {
        out := make([]rt.Value, len(array(x)))
        for i, v := range array(x) { out[i] = rt.Apply(f, v) }
        return rt.Array(out)
    })
    apply := rt.Func2(func(f, x rt.Value) rt.Value {
        out := []rt.Value{}
        for _, fn := range array(f) { for _, v := range array(x) { out = append(out, rt.Apply(fn, v)) } }
        return rt.Array(out)
    })
    f := rt.Func(func(x rt.Value) rt.Value { return rt.Array([]rt.Value{x, rt.Int(x.IntVal + 10)}) })
    for _, n := range []int{0, 1, 2, 3, 4, 7} {
        got := array(traverse(apply, mapFn, pure, f, values(n)))
        if len(got) != 1 << n { t.Fatalf("n=%d alternatives=%d", n, len(got)) }
        for i, result := range got {
            row := array(result)
            if len(row) != n { t.Fatalf("n=%d row length=%d", n, len(row)) }
            for j, x := range row {
                want := int64(j + 10 * ((i >> (n - 1 - j)) & 1))
                if x.IntVal != want { t.Fatalf("n=%d row=%d column=%d: %d != %d", n, i, j, x.IntVal, want) }
            }
        }
    }
}

type stateResult struct { value rt.Value; state int64 }
func TestTraversalStateEffectOrder(t *testing.T) {
    pure := rt.Func(func(x rt.Value) rt.Value { return rt.Func(func(s rt.Value) rt.Value {
        return rt.Any(stateResult{x, s.IntVal})
    }) })
    mapFn := rt.Func2(func(f, x rt.Value) rt.Value { return rt.Func(func(s rt.Value) rt.Value {
        result := rt.Apply(x, s).PtrVal().(stateResult)
        return rt.Any(stateResult{rt.Apply(f, result.value), result.state})
    }) })
    apply := rt.Func2(func(f, x rt.Value) rt.Value { return rt.Func(func(s rt.Value) rt.Value {
        left := rt.Apply(f, s).PtrVal().(stateResult)
        right := rt.Apply(x, rt.Int(left.state)).PtrVal().(stateResult)
        return rt.Any(stateResult{rt.Apply(left.value, right.value), right.state})
    }) })
    for _, n := range []int{0, 1, 2, 3, 4, 7, 8, 9} {
        effects := []int64{}
        f := rt.Func(func(x rt.Value) rt.Value { return rt.Func(func(s rt.Value) rt.Value {
            effects = append(effects, x.IntVal)
            return rt.Any(stateResult{rt.Int(x.IntVal + s.IntVal), s.IntVal + 1})
        }) })
        computation := traverse(apply, mapFn, pure, f, values(n))
        if len(effects) != 0 { t.Fatal("State effects ran during traversal construction") }
        result := rt.Apply(computation, rt.Int(10)).PtrVal().(stateResult)
        if result.state != int64(10+n) || len(effects) != n { t.Fatalf("State result: %+v effects=%v", result, effects) }
        for i, v := range array(result.value) {
            if v.IntVal != int64(10+2*i) || effects[i] != int64(i) { t.Fatalf("State order: %+v effects=%v", result, effects) }
        }
    }
}
`);
    const result = spawnSync('go', ['test', '-count=1', '-v', './...'], {
        cwd: directory, encoding: 'utf8', timeout: 60_000,
        env: { ...process.env, GOWORK: 'off' }, maxBuffer: 1024 * 1024,
    });
    assert.ifError(result.error);
    assert.equal(result.signal, null);
    assert.equal(result.status, 0, result.stdout + result.stderr);
    for (const name of ['HigherOrderValueCallbacks', 'TraversalEitherAndStrictCallbackOrder', 'TraversalArrayCartesianOrder', 'TraversalStateEffectOrder']) {
        assert.match(result.stdout, new RegExp(`--- PASS: Test${name}`));
    }
});
