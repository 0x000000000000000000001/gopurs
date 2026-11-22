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
import * as Ref from '../output/Effect.Ref/index.js';
import { worker } from '../output/Gopurs.ArraySnoc/index.js';
import { StmtEmpty, flattenStmts } from '../output/Gopurs.ExprContext/index.js';
import * as Go from '../output/Gopurs.GoAst/index.js';
import { LocalReference } from '../output/Gopurs.GoSymbol/index.js';
import { declaration } from '../output/Gopurs.ModuleWorkers/index.js';
import { printGoDecl, printGoExpr } from '../output/Gopurs.Printer/index.js';
import { runtimeGoCode } from '../output/Gopurs.Runtime.Source/index.js';
import { TcoExpr } from '../output/PureScript.Backend.Optimizer.Codegen.Tco/index.js';
import * as Core from '../output/PureScript.Backend.Optimizer.CoreFn/index.js';
import * as Syntax from '../output/PureScript.Backend.Optimizer.Syntax/index.js';
import { withReboxFields } from './codegen-metadata.mjs';

const expr = syntax => new TcoExpr(null, syntax);
const typed = value => expr(new Syntax.Typed(Core.Any.value, value));
const local = (name, level) => expr(new Syntax.Local(new Just(name), level));
const variable = (module, name) => expr(new Syntax.Var(new Core.Qualified(
    module === null ? Nothing.value : new Just(module), name)));
const app = (fn, args) => expr(new Syntax.App(fn, args));
const effect = (module, name, args) => expr(new Syntax.UncurriedEffectApp(variable(module, name), args));

// This is the optimized ST shape, not an approximation based on a name or type.
function canonical(options = {}) {
    const annotate = options.annotated ? typed : value => value;
    const owner = (key, fallback) => Object.hasOwn(options, key) ? options[key] : fallback;
    const xs = annotate(local('xs', 0));
    const x = annotate(local('x', 1));
    const pushName = options.pushName ?? 'push';
    const pushLevel = options.pushLevel ?? 2;
    const arrayName = options.arrayName ?? 'array';
    const arrayLevel = options.arrayLevel ?? 3;
    const unusedName = options.unusedName ?? 'unused';
    const unusedLevel = options.unusedLevel ?? 4;
    const array = annotate(local(arrayName, arrayLevel));
    const push = annotate(local(pushName, pushLevel));
    const pushValue = annotate(app(annotate(variable(owner('pushModule', 'Data.Array.ST'), 'push')),
        options.pushArguments ?? [x]));
    const thaw = annotate(effect(owner('thawModule', 'Data.Array.ST'), options.thawName ?? 'thawImpl',
        options.thawArguments ?? [xs]));
    const pushCall = annotate(app(options.pushTarget ?? push, options.applyArguments ?? [array]));
    const freeze = annotate(effect(owner('freezeModule', 'Data.Array.ST'), options.freezeName ?? 'unsafeFreezeImpl',
        options.freezeArguments ?? [array]));
    let tail = annotate(new TcoExpr(null, new Syntax.EffectBind(new Just(unusedName), unusedLevel, pushCall, freeze)));
    if (options.extraEffect) tail = expr(new Syntax.EffectBind(new Just('extra'), 5,
        effect('Probe', 'callback', [array]), tail));
    const afterPush = options.reordered
        ? expr(new Syntax.EffectBind(new Just(unusedName), unusedLevel, pushCall,
            expr(new Syntax.EffectBind(new Just(arrayName), arrayLevel, thaw, freeze))))
        : annotate(expr(new Syntax.EffectBind(new Just(arrayName), arrayLevel, thaw, tail)));
    const computation = annotate(expr(new Syntax.Let(new Just(pushName), pushLevel, pushValue, afterPush)));
    return annotate(app(annotate(variable(owner('runModule', 'Control.Monad.ST.Internal'), options.runName ?? 'run')),
        options.runArguments ?? [computation]));
}

const params = element => [new Tuple('xs_0', new Go.TypeNativeArray(element)), new Tuple('x_1', element)];
const select = (element, body = canonical(), owner = 'Data_Array', name = 'snoc') =>
    worker(owner)(name)(params(element))(new Go.TypeNativeArray(element))(body);

test('snoc requires the qualified worker, matching native ABI and complete ordered ST body', () => {
    for (const name of ['snoc', 'snoc__1895563893']) {
        for (const annotated of [false, true]) {
            assert.ok(select(Go.TypeInt64.value, canonical({ annotated }), 'Data_Array', name) instanceof Just);
        }
    }
    for (const [owner, name] of [['Other', 'snoc'], ['Data.Array', 'snoc'], ['Data_Array', 'unsnoc'],
        ['Data_Array', 'snocOther'], ['Data_Array', 'snoc_123']]) {
        assert.deepEqual(select(Go.TypeInt64.value, canonical(), owner, name), Nothing.value);
    }
    const native = new Go.TypeNativeArray(Go.TypeInt64.value);
    for (const [arguments_, result] of [
        [params(Go.TypeInt64.value).slice(0, 1), native],
        [[...params(Go.TypeInt64.value), new Tuple('extra_5', Go.TypeValue.value)], native],
        [[new Tuple('xs_0', Go.TypeValue.value), new Tuple('x_1', Go.TypeValue.value)], Go.TypeValue.value],
        [[new Tuple('xs_0', native), new Tuple('x_1', Go.TypeString.value)], native],
        [params(Go.TypeInt64.value), new Go.TypeNativeArray(Go.TypeString.value)],
        [[new Tuple('xs_0', native), new Tuple('xs_0', Go.TypeInt64.value)], native],
    ]) assert.deepEqual(worker('Data_Array')('snoc')(arguments_)(result)(canonical()), Nothing.value);

    const mutations = [
        { runModule: 'Other' }, { runModule: null }, { runName: 'differentRun' }, { runArguments: [] },
        { runArguments: [canonical(), local('extra', 9)] },
        { pushModule: 'Other' }, { pushModule: null }, { pushArguments: [local('xs', 0)] },
        { pushArguments: [local('x', 1), local('extra', 9)] },
        { thawModule: 'Other' }, { thawModule: null }, { thawName: 'unsafeThawImpl' }, { thawArguments: [local('x', 1)] },
        { applyArguments: [local('xs', 0)] }, { applyArguments: [local('array', 3), local('extra', 9)] },
        { pushTarget: variable('Other', 'push') },
        { freezeModule: 'Other' }, { freezeModule: null }, { freezeName: 'freezeImpl' }, { freezeArguments: [local('xs', 0)] },
        { extraEffect: true }, { reordered: true },
        // Rebinding a parameter, push closure, array or unused result can change semantics.
        { pushName: 'xs', pushLevel: 0 }, { pushName: 'x', pushLevel: 1 },
        { arrayName: 'push', arrayLevel: 2 }, { arrayName: 'xs', arrayLevel: 0 },
        { unusedName: 'array', unusedLevel: 3 }, { unusedName: 'push', unusedLevel: 2 },
        { unusedName: 'x', unusedLevel: 1 },
    ];
    for (const mutation of mutations) {
        assert.deepEqual(select(Go.TypeInt64.value, canonical(mutation), 'Data_Array', 'snoc__123'),
            Nothing.value, JSON.stringify(mutation));
    }
    assert.deepEqual(select(Go.TypeInt64.value, local('xs', 0), 'Data_Array', 'snoc__123'), Nothing.value,
        'a specialization-looking name and matching ABI cannot replace an arbitrary body');
});

const metadata = withReboxFields({
    pointerAdtPaths: emptyMap, pointerAdtNodes: emptySet, pointerAdtLeaves: emptyMap,
    enumAdts: emptySet, enumCtors: emptySet, elidedCtors: emptySet, ctorTypes: emptyMap,
    classDeclsFields: emptyMap, globalTypes: emptyMap, globalFunctions: emptyMap,
});

test('ModuleWorkers selects the native body before translation and retains its normal fallback', () => {
    const type = new Go.TypeNativeArray(Go.TypeInt64.value);
    const signature = { reference: new LocalReference('Call_Data_Array_snoc'),
        fArgs: [type, Go.TypeInt64.value], fRet: type, arity: 2 };
    let translated = 0;
    const fallback = _context => nextId => _body => {
        translated++;
        return { stmts: StmtEmpty.value, expr: Go.rawGo('fallback()'), exprType: type, nextId };
    };
    const context = () => ({
        metadata, codegenStateRef: Ref.new({ declarations: [], globalId: 0, reboxPairs: emptySet })(),
        depth: 0, modNameStr: 'Data_Array', recVars: [], bound: emptyMap,
        moduleFunctions: insert(ordString)('snoc')(signature)(emptyMap),
        tcoIdent: Nothing.value, loopCtx: [], options: { isTail: false, inEffectBlock: false },
        mbExpectedExprType: Nothing.value, mbExpectedGoType: Nothing.value,
    });
    for (const [body, expectedFallbacks] of [[canonical(), 0], [canonical({ extraEffect: true }), 1]]) {
        const ctx = context();
        const wrapper = declaration(fallback)(ctx)(false)({ ident: 'snoc', args: ['xs_0', 'x_1'], body,
            val: typed(body) });
        assert.equal(translated, expectedFallbacks);
        const declarations = Ref.read(ctx.codegenStateRef)().declarations.map(printGoDecl).join('\n');
        assert.match(printGoDecl(wrapper), /gopurs_runtime\.Func2/,
            'the existing boxed/partial application wrapper remains present');
        if (expectedFallbacks === 0) {
            assert.match(declarations, /snoc_output/);
            assert.doesNotMatch(declarations, /fallback\(/);
        } else assert.match(declarations, /fallback\(/);
    }
});

test('emitted snoc uses fresh storage and preserves scalars and shallow composite sharing', t => {
    const pointer = Go.structPointer({ baseStructName: 'Data_Snoc_Cell', fullName: 'Snoc.Cell',
        structName: 'Constructor_Snoc_Cell', symbol: { owner: 'generated/Snoc', name: 'Constructor_Snoc_Cell' } })([]);
    const record = new Go.TypeRecord([new Tuple('cell', pointer),
        new Tuple('values', new Go.TypeNativeArray(Go.TypeInt64.value))]);
    const types = [
        ['ints', Go.TypeInt64.value], ['numbers', Go.TypeFloat64.value], ['strings', Go.TypeString.value],
        ['booleans', Go.TypeBool.value], ['values', Go.TypeValue.value], ['pointers', pointer],
        ['records', record], ['nested', new Go.TypeNativeArray(Go.TypeInt64.value)],
        ['functions', new Go.TypeFunc([Go.TypeInt64.value], Go.TypeInt64.value)],
    ];
    const declarations = types.map(([name, type]) => {
        const result = select(type);
        assert.ok(result instanceof Just, name);
        const output = result.value0;
        assert.deepEqual(output.exprType, new Go.TypeNativeArray(type));
        const body = [...flattenStmts(output.stmts), new Go.GoReturn(output.expr)].map(printGoExpr).join('\n');
        assert.doesNotMatch(body, /Rebox_|\.Apply\(|\.Array\(|\.RecordDict|append\(/);
        return `func snoc_${name}(xs_0 []${Go.goTypeToStr(type)}, x_1 ${Go.goTypeToStr(type)}) []${Go.goTypeToStr(type)} {\n${body}\n}`;
    });
    const directory = mkdtempSync(join(tmpdir(), 'gopurs-array-snoc-'));
    t.after(() => rmSync(directory, { recursive: true, force: true }));
    mkdirSync(join(directory, 'gopurs_runtime'));
    writeFileSync(join(directory, 'go.mod'), 'module gopurs/output\n\ngo 1.22\n');
    writeFileSync(join(directory, 'gopurs_runtime/runtime.go'), runtimeGoCode);
    const cellField = Go.recordFieldName('cell');
    const valuesField = Go.recordFieldName('values');
    const recordName = Go.goTypeToStr(record);
    writeFileSync(join(directory, 'main.go'), `package main
import ("fmt"; "math"; "gopurs/output/gopurs_runtime")
type Constructor_Snoc_Cell struct { value int64 }
${declarations.join('\n')}
func check(ok bool, message string) { if !ok { panic(message) } }
func checkFresh[T comparable](call func([]T, T) []T, first, second, last T) {
    storage := make([]T, 8)
    storage[0], storage[1], storage[2] = first, second, first
    input := storage[:2]
    output := call(input, last)
    check(len(output) == 3 && output[0] == first && output[1] == second && output[2] == last, "values/order")
    check(storage[2] == first && &output[0] != &input[0], "spare capacity reused")
    output[0] = last
    check(input[0] == first, "input mutated")
    again := call(input, second)
    check(again[0] == first && again[2] == second && output[2] == last, "versions share a buffer")
    for _, empty := range [][]T{nil, make([]T, 0), make([]T, 0, 8)} {
        one := call(empty, last)
        check(one != nil && len(one) == 1 && one[0] == last && len(empty) == 0, "empty/nil")
    }
}
func main() {
    checkFresh(snoc_ints, int64(1)<<50, -(int64(1)<<55), int64(1)<<60)
    checkFresh(snoc_numbers, 1.25, -2.5, 3.75)
    checkFresh(snoc_strings, "", "λ😀", "last")
    checkFresh(snoc_booleans, true, false, false)
    negativeZero := math.Copysign(0, -1)
    numbers := snoc_numbers([]float64{negativeZero, math.NaN(), math.Inf(1)}, negativeZero)
    check(math.Signbit(numbers[0]) && math.IsNaN(numbers[1]) && math.IsInf(numbers[2], 1) && math.Signbit(numbers[3]), "float payloads")
    boxed := gopurs_runtime.RecordDict1("kept", gopurs_runtime.Int(9))
    checkFresh(snoc_values, boxed, gopurs_runtime.Int(7), gopurs_runtime.Str("last"))
    boxedResult := snoc_values([]gopurs_runtime.Value{boxed}, boxed)
    check(boxedResult[0] == boxed && boxedResult[1] == boxed, "boxed identity/tag")
    first, second, last := &Constructor_Snoc_Cell{1}, &Constructor_Snoc_Cell{2}, &Constructor_Snoc_Cell{3}
    checkFresh(snoc_pointers, first, second, last)
    pointers := snoc_pointers([]*Constructor_Snoc_Cell{first}, last)
    pointers[0].value = 41
    check(first.value == 41 && pointers[1] == last, "ADT reference sharing")
    inner := []int64{5, 7}
    tail := []int64{11}
    nestedInput := [][]int64{inner}
    nested := snoc_nested(nestedInput, tail)
    nested[0][0], nested[1][0] = 13, 17
    check(inner[0] == 13 && tail[0] == 17, "nested arrays were deep copied")
    nested[0] = []int64{19}
    check(nestedInput[0][0] == 13, "outer array shares storage")
    recordsInput := []${recordName}{{${cellField}: first, ${valuesField}: inner}}
    records := snoc_records(recordsInput, ${recordName}{${cellField}: last, ${valuesField}: tail})
    records[0].${cellField}.value = 43
    records[0].${valuesField}[0] = 47
    check(first.value == 43 && inner[0] == 47, "record reference sharing")
    records[0].${cellField} = second
    check(recordsInput[0].${cellField} == first && len(recordsInput) == 1, "record container was mutated")
    count := int64(0)
    callback := func(x int64) int64 { count++; return x+1 }
    functions := snoc_functions([]func(int64) int64{callback}, func(x int64) int64 { return x*2 })
    check(count == 0 && functions[0](10) == 11 && count == 1 && functions[1](10) == 20, "function values evaluated or changed")
    fmt.Println("snoc semantics passed")
}
`);
    const run = spawnSync('go', ['run', '.'], { cwd: directory, encoding: 'utf8', timeout: 60_000,
        env: { ...process.env, GOWORK: 'off', GOTOOLCHAIN: 'local' } });
    assert.equal(run.status, 0, run.stdout + run.stderr);
    assert.equal(run.stdout.trim(), 'snoc semantics passed');
});
