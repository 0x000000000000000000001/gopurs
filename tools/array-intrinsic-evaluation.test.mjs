import assert from 'node:assert/strict';
import { spawnSync } from 'node:child_process';
import { mkdirSync, mkdtempSync, rmSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join } from 'node:path';
import { fileURLToPath, pathToFileURL } from 'node:url';
import test from 'node:test';

const project = fileURLToPath(new URL('../', import.meta.url));
const load = module => import(pathToFileURL(join(project, 'output', module, 'index.js')).href);
const [Maps, Maybe, Ord, Sets, Ref, Go, Conversions, Printer, Runtime, Tco, Core, Syntax, Statements, Calls] = await Promise.all(
    ['Data.Map', 'Data.Maybe', 'Data.Ord', 'Data.Set', 'Effect.Ref', 'Gopurs.GoAst',
        'Gopurs.GoConversions', 'Gopurs.Printer', 'Gopurs.Runtime',
        'PureScript.Backend.Optimizer.Codegen.Tco', 'PureScript.Backend.Optimizer.CoreFn',
        'PureScript.Backend.Optimizer.Syntax', 'Gopurs.ExprContext', 'Gopurs.CallExprs'].map(load));
const { withReboxFields } = await import(pathToFileURL(join(project, 'tools/codegen-metadata.mjs')).href);
const metadata = withReboxFields({
    pointerAdtPaths: Maps.empty, pointerAdtNodes: Sets.empty, pointerAdtLeaves: Maps.empty,
    enumAdts: Sets.empty, enumCtors: Sets.empty, elidedCtors: Sets.empty, ctorTypes: Maps.empty,
    classDeclsFields: Maps.empty, globalTypes: Maps.empty, globalFunctions: Maps.empty,
});
const context = () => ({
    metadata, codegenStateRef: Ref.new({ declarations: [], globalId: 0, reboxPairs: Sets.empty })(),
    depth: 2, modNameStr: 'Probe', recVars: [], moduleFunctions: Maps.empty, bound: Maps.empty,
    tcoIdent: Maybe.Nothing.value, loopCtx: [], options: { isTail: false, inEffectBlock: false },
    mbExpectedExprType: Maybe.Nothing.value,
});
const variable = (name, module = 'Probe') => new Tco.TcoExpr(null,
    new Syntax.Var(new Core.Qualified(new Maybe.Just(module), name)));

function lower(kind, mode, extra = false) {
    const ctx = context();
    const curried = mode.startsWith('curried');
    const native = mode.includes('Native') || mode.startsWith('uncurriedWorker');
    const worker = mode.startsWith('uncurriedWorker');
    const nativeSeed = mode.endsWith('Seed');
    const callbackArity = kind === 'fold' ? 2 : 1;
    const callbackResult = kind === 'filter' ? Go.TypeBool.value : Go.TypeInt64.value;
    if (worker) ctx.moduleFunctions = Maps.insert(Ord.ordString)(`Probe.${kind}`)(
        { arity: callbackArity, fArgs: Array(callbackArity).fill(Go.TypeInt64.value), fRet: callbackResult })(Maps.empty);
    const callback = variable(worker ? kind : `callback_${kind}`);
    const seed = variable('seed');
    const array = variable('array');
    const surplus = variable('extra');
    const args = [callback, ...(kind === 'fold' ? [seed] : []), array, ...(extra ? [surplus] : [])];
    const name = kind === 'fold' ? 'foldlArray' : kind === 'map' ? 'arrayMap' : curried ? 'filter' : 'filterImpl';
    const fn = variable(name, kind === 'fold' ? 'Data.Foldable' : 'Data.Array');
    const expressions = new Map([
        // Ordinary translated function values use the boxed ABI. Native FFI callbacks
        // are adapted at the boundary; native worker metadata remains separate.
        [callback, { label: 'callback', expression: extra ? 'make_extra_callback()'
            : `make_${native ? 'bridged' : 'boxed'}_${kind}()`, type: Go.TypeValue.value }],
        [seed, { label: 'seed', expression: extra ? 'make_extra_seed()'
            : nativeSeed ? 'markValue("seed:expression", int64(1))' : 'markValue("seed:expression", gopurs_runtime.Int(1))',
            type: nativeSeed ? Go.TypeInt64.value : Go.TypeValue.value }],
        [array, { label: 'array', expression: `markValue("array:expression", ${native ? 'input' : 'boxedInput(input)'})`,
            type: native ? new Go.TypeNativeArray(Go.TypeInt64.value) : Go.TypeValue.value }],
        [surplus, { label: 'extra', expression: 'markValue("extra:expression", gopurs_runtime.Int(7))', type: Go.TypeValue.value }],
    ]);
    const translated = [];
    const translate = _child => nextId => expression => {
        const item = expressions.get(expression);
        assert.ok(item, 'only intrinsic arguments should reach the fixture translator');
        translated.push(item.label);
        return { stmts: new Statements.StmtLeaf(Go.rawGo(`mark("${item.label}:statement")`)),
            expr: Go.rawGo(item.expression), exprType: item.type, nextId };
    };
    const expression = new Tco.TcoExpr(null, curried ? new Syntax.App(fn, args) : new Syntax.UncurriedApp(fn, args));
    const result = curried ? Calls.application(translate)(ctx)(0)(expression)
        : Calls.uncurriedApplication(translate)(ctx)(0)(expression)(fn)(args);
    assert.deepEqual(translated, ['callback', ...(kind === 'fold' ? ['seed'] : []), 'array', ...(extra ? ['extra'] : [])]);
    const boxed = Conversions.boxGoExpr(ctx.codegenStateRef)('Probe')(result.expr)(result.exprType);
    const body = [...Statements.flattenStmts(result.stmts), new Go.GoReturn(boxed)]
        .map(Printer.printGoExpr).join('\n');
    return `func probe_${mode}_${kind}${extra ? '_extra' : ''}(input []int64) gopurs_runtime.Value {\n${body}\n}`;
}

test('intrinsic callback factories run once in argument order, including empty inputs and native workers', t => {
    const modes = ['curriedBoxed', 'curriedNative', 'uncurriedBoxed', 'uncurriedNative', 'uncurriedWorker'];
    const kinds = ['map', 'filter', 'fold'];
    const declarations = modes.flatMap(mode => kinds.map(kind => lower(kind, mode)));
    // Native accumulators must retain their scalar return type.
    const nativeSeedModes = ['uncurriedNativeSeed', 'uncurriedWorkerSeed'];
    declarations.push(...nativeSeedModes.map(mode => lower('fold', mode)));
    declarations.push(lower('fold', 'curriedBoxed', true));
    const directory = mkdtempSync(join(tmpdir(), 'gopurs-intrinsic-evaluation-'));
    t.after(() => rmSync(directory, { recursive: true, force: true }));
    mkdirSync(join(directory, 'gopurs_runtime'));
    writeFileSync(join(directory, 'go.mod'), 'module gopurs/output\n\ngo 1.22\n');
    writeFileSync(join(directory, 'gopurs_runtime/runtime.go'), Runtime.runtimeGoCode);
    writeFileSync(join(directory, 'main.go'), `package main
import ("fmt"; "slices"; "unsafe"; "gopurs/output/gopurs_runtime")
var trace []string
func mark(value string) { trace = append(trace, value) }
func markValue[T any](label string, value T) T { mark(label); return value }
func boxedInput(input []int64) gopurs_runtime.Value {
    values := make([]gopurs_runtime.Value, len(input))
    for i, value := range input { values[i] = gopurs_runtime.Int(value) }
    return gopurs_runtime.Array(values)
}
func Call_Probe_map(value int64) int64 { mark(fmt.Sprintf("call:%d", value)); return value * 2 }
func Call_Probe_filter(value int64) bool { mark(fmt.Sprintf("call:%d", value)); return value % 2 == 0 }
func Call_Probe_fold(acc, value int64) int64 { mark(fmt.Sprintf("call:%d", value)); return acc * 10 + value }
func make_native_map() func(int64) int64 { mark("callback:factory"); return Call_Probe_map }
func make_native_filter() func(int64) bool { mark("callback:factory"); return Call_Probe_filter }
func make_native_fold() func(int64) func(int64) int64 {
    mark("callback:factory")
    return func(acc int64) func(int64) int64 {
        return func(value int64) int64 { return Call_Probe_fold(acc, value) }
    }
}
// Model FfiBridge.Values.toValue: evaluate the native result once, then adapt
// its scalar arguments and result to the ordinary boxed function ABI.
func make_bridged_map() gopurs_runtime.Value {
    callback := make_native_map()
    return gopurs_runtime.Func(func(value gopurs_runtime.Value) gopurs_runtime.Value {
        result := callback(gopurs_runtime.Unbox[int64](value))
        return gopurs_runtime.Int(result)
    })
}
func make_bridged_filter() gopurs_runtime.Value {
    callback := make_native_filter()
    return gopurs_runtime.Func(func(value gopurs_runtime.Value) gopurs_runtime.Value {
        result := callback(gopurs_runtime.Unbox[int64](value))
        return gopurs_runtime.Box(result)
    })
}
func make_bridged_fold() gopurs_runtime.Value {
    callback := make_native_fold()
    return gopurs_runtime.Func(func(acc gopurs_runtime.Value) gopurs_runtime.Value {
        inner := callback(gopurs_runtime.Unbox[int64](acc))
        return gopurs_runtime.Func(func(value gopurs_runtime.Value) gopurs_runtime.Value {
            result := inner(gopurs_runtime.Unbox[int64](value))
            return gopurs_runtime.Int(result)
        })
    })
}
func make_boxed_map() gopurs_runtime.Value {
    mark("callback:factory"); return gopurs_runtime.Func(func(value gopurs_runtime.Value) gopurs_runtime.Value {
        return gopurs_runtime.Int(Call_Probe_map(value.IntVal))
    })
}
func make_boxed_filter() gopurs_runtime.Value {
    mark("callback:factory"); return gopurs_runtime.Func(func(value gopurs_runtime.Value) gopurs_runtime.Value {
        return gopurs_runtime.Bool(Call_Probe_filter(value.IntVal))
    })
}
func make_boxed_fold() gopurs_runtime.Value {
    mark("callback:factory"); return gopurs_runtime.Func2(func(acc, value gopurs_runtime.Value) gopurs_runtime.Value {
        return gopurs_runtime.Int(Call_Probe_fold(acc.IntVal, value.IntVal))
    })
}
func make_extra_seed() gopurs_runtime.Value {
    mark("seed:expression")
    return gopurs_runtime.Func(func(extra gopurs_runtime.Value) gopurs_runtime.Value {
        mark("result:apply"); return gopurs_runtime.Int(10 + extra.IntVal)
    })
}
func make_extra_callback() gopurs_runtime.Value {
    mark("callback:factory")
    return gopurs_runtime.Func2(func(acc, value gopurs_runtime.Value) gopurs_runtime.Value {
        mark(fmt.Sprintf("call:%d", value.IntVal))
        return gopurs_runtime.Func(func(extra gopurs_runtime.Value) gopurs_runtime.Value {
            return gopurs_runtime.Int(gopurs_runtime.Apply(acc, extra).IntVal + value.IntVal)
        })
    })
}
${declarations.join('\n')}
func main() {
    _ = unsafe.Pointer(nil)
    probes := []struct { name, kind string; run func([]int64) gopurs_runtime.Value }{
        ${[
            ...modes.flatMap(mode => kinds.map(kind => `{"${mode}", "${kind}", probe_${mode}_${kind}}`)),
            ...nativeSeedModes.map(mode => `{"${mode}", "fold", probe_${mode}_fold}`),
        ].join(',\n')},
    }
    for _, input := range [][]int64{nil, {2, 3}} {
        original := slices.Clone(input)
        for _, probe := range probes {
            trace = nil
            result := probe.run(input)
            expected := []string{"callback:statement", "callback:factory"}
            if probe.kind == "fold" { expected = append(expected, "seed:statement", "seed:expression") }
            expected = append(expected, "array:statement", "array:expression")
            for _, value := range input { expected = append(expected, fmt.Sprintf("call:%d", value)) }
            if !slices.Equal(trace, expected) { panic(fmt.Sprintf("%s/%s input=%v trace=%v expected=%v", probe.name, probe.kind, input, trace, expected)) }
            if !slices.Equal(input, original) { panic("input mutated") }
            if probe.kind == "fold" {
                value := int64(1); for _, item := range input { value = value*10 + item }
                if result.IntVal != value { panic("incorrect left fold") }
            } else {
                got := *(*[]gopurs_runtime.Value)(result.UnsafePtr)
                expected := []int64{}
                for _, value := range input { if probe.kind == "map" { expected = append(expected, value*2) } else if value%2 == 0 { expected = append(expected, value) } }
                if len(got) != len(expected) { panic("incorrect output length") }
                for i, value := range expected { if got[i].IntVal != value { panic("incorrect output element") } }
            }
        }
        trace = nil
        result := probe_curriedBoxed_fold_extra(input)
        expected := []string{"callback:statement", "callback:factory", "seed:statement", "seed:expression", "array:statement", "array:expression"}
        want := int64(17)
        for _, value := range input { expected = append(expected, fmt.Sprintf("call:%d", value)); want += value }
        expected = append(expected, "extra:statement", "extra:expression", "result:apply")
        if result.IntVal != want || !slices.Equal(trace, expected) { panic(fmt.Sprintf("surplus input=%v value=%v trace=%v expected=%v", input, result.IntVal, trace, expected)) }
    }
    fmt.Println("factory once, strict argument order, empty/nonempty, native worker and surplus stages: ok")
}
`);
    const result = spawnSync('go', ['run', '.'], { cwd: directory, encoding: 'utf8', timeout: 60_000,
        env: { ...process.env, GOWORK: 'off' } });
    assert.ifError(result.error);
    assert.equal(result.status, 0, result.stdout + result.stderr);
    assert.match(result.stdout, /factory once, strict argument order/);
});
