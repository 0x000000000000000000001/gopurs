import assert from 'node:assert/strict';
import test from 'node:test';
import { empty as emptyMap, insert, lookup } from '../output/Data.Map/index.js';
import { Just, Nothing } from '../output/Data.Maybe/index.js';
import { ordString } from '../output/Data.Ord/index.js';
import { empty as emptySet } from '../output/Data.Set/index.js';
import { Tuple } from '../output/Data.Tuple/index.js';
import * as Ref from '../output/Effect.Ref/index.js';
import * as CodeGen from '../output/Gopurs.CodeGen/index.js';
import * as TypedExprs from '../output/Gopurs.TypedExprs/index.js';
import * as ModuleBindings from '../output/Gopurs.ModuleBindings/index.js';
import { StmtEmpty } from '../output/Gopurs.ExprContext/index.js';
import * as Go from '../output/Gopurs.GoAst/index.js';
import { printGoExpr } from '../output/Gopurs.Printer/index.js';
import { runtimeGoCode } from '../output/Gopurs.Runtime.Source/index.js';
import { TcoExpr } from '../output/PureScript.Backend.Optimizer.Codegen.Tco/index.js';
import * as C from '../output/PureScript.Backend.Optimizer.CoreFn/index.js';
import { NeutralExpr } from '../output/PureScript.Backend.Optimizer.Semantics/index.js';
import * as S from '../output/PureScript.Backend.Optimizer.Syntax/index.js';
import { withReboxFields } from './codegen-metadata.mjs';
import { goModule, runGo, writeGoFile } from './lib/go.mjs';

const metadata = withReboxFields({
    elidedCtors: emptySet, ctorTypes: emptyMap,
    pointerAdtPaths: insert(ordString)('Data.Tuple.Tuple')({ ctorName: 'Tuple', arity: 2 })(emptyMap),
    pointerAdtNodes: emptySet, pointerAdtLeaves: emptyMap, enumAdts: emptySet,
    enumCtors: emptySet, globalTypes: emptyMap, globalFunctions: emptyMap, classDeclsFields: emptyMap,
});
const position = new C.Record(new C.Row([
    new Tuple('column', C.Int.value), new Tuple('line', C.Int.value),
], Nothing.value));
const issue = new C.Record(new C.Row([
    new Tuple('error', C.Any.value), new Tuple('position', position),
], Nothing.value));
const state = new C.Record(new C.Row([
    new Tuple('consumed', C.Boolean.value), new Tuple('errors', new C.Array(issue)),
    new Tuple('stream', C.Any.value),
], Nothing.value));
const pair = new C.ADT('Data.Tuple.Tuple', ['Data', 'Tuple', 'Tuple'], [position, new C.Array(issue)]);
const expr = syntax => new NeutralExpr(syntax);
const typed = (type, value) => expr(new S.Typed(type, value));
const local = (name, level) => expr(new S.Local(new Just(name), level));
const lambda = (params, body) => expr(new S.Abs(params.map(([name, level]) => new Tuple(new Just(name), level)), body));
const letIn = (name, level, value, body) => expr(new S.Let(new Just(name), level, value, body));
const app = (fn, ...args) => expr(new S.App(fn, args));
const module = bindings => ({
    name: 'CompositeFlow', bindings: bindings.map(([name, value]) => ({ recursive: false, bindings: [new Tuple(name, value)] })),
    comments: [], imports: emptySet, exports: emptySet, reExports: emptySet,
    dataTypes: emptyMap, dataDecls: [], classDecls: [], foreign: emptyMap,
    implementations: emptyMap, directives: emptyMap,
});

test('composite annotations preserve producers; explicit consumer representations still convert', () => {
    const context = {
        metadata, codegenStateRef: Ref.new({ declarations: [], globalId: 0, reboxPairs: emptySet })(),
        depth: 0, modNameStr: 'CompositeFlow', recVars: [], moduleFunctions: emptyMap, bound: emptyMap,
        tcoIdent: Nothing.value, loopCtx: [], options: { isTail: false, inEffectBlock: false },
        mbExpectedExprType: Nothing.value, mbExpectedGoType: Nothing.value,
    };
    const source = new Go.GoCall(new Go.GoVar('produce'), []);
    const result = { stmts: StmtEmpty.value, expr: source, exprType: Go.TypeValue.value, nextId: 8 };
    const syntax = new TcoExpr(null, new S.Local(new Just('input'), 0));
    const translate = () => () => () => result;
    for (const type of [state, new C.Array(issue), pair]) {
        const preserved = TypedExprs.annotated(translate)(context)(0)(type)(syntax);
        assert.equal(preserved, result, 'annotations must not demand a composite allocation');
    }
    const target = new Go.TypeRecord([new Tuple('column', Go.TypeInt64.value), new Tuple('line', Go.TypeInt64.value)]);
    const concrete = TypedExprs.annotated(translate)({ ...context, mbExpectedGoType: new Just(target) })(0)(position)(syntax);
    assert.deepEqual(concrete.exprType, target);
    assert.match(printGoExpr(concrete.expr), /RecordGet/);
    assert.equal((printGoExpr(concrete.expr).match(/produce\(\)/g) ?? []).length, 1);
    const scalar = TypedExprs.annotated(translate)(context)(0)(C.Int.value)(syntax);
    assert.equal(printGoExpr(scalar.expr), 'produce().IntVal');
    const open = new C.Record(new C.Row([new Tuple('column', C.Int.value)], new Just(new C.TypeVar('row'))));
    assert.equal(TypedExprs.annotated(translate)(context)(0)(open)(syntax), result,
        'an open row must retain any additional fields');
});

test('Tuple worker results preserve instantiated array slots while ordinary scalar sums keep their ABI', () => {
    const scalarPair = new C.ADT('Data.Tuple.Tuple', ['Data', 'Tuple', 'Tuple'], [C.Int.value, C.Int.value]);
    const prepared = ModuleBindings.prepare(metadata)('CompositeFlow')(module([
        ['composite', typed(new C.Func([C.Any.value], pair), lambda([['input', 0]], typed(pair, local('input', 0))))],
        ['scalar', typed(new C.Func([C.Any.value], scalarPair), lambda([['input', 0]], typed(scalarPair, local('input', 0))))],
    ]));
    const composite = lookup(ordString)('composite')(prepared.functions).value0;
    assert.ok(composite.fRet instanceof Go.TypeStructPointer);
    assert.ok(composite.fRet.value0.typeArgs[1] instanceof Go.TypeNativeArray);
    const scalar = lookup(ordString)('scalar')(prepared.functions).value0;
    assert.deepEqual(scalar.fRet, new Go.TypeStructValue('Data.Tuple.Tuple', [Go.TypeValue.value, Go.TypeValue.value]));
});

test('imported Tuple workers retain native payloads and expose the same boxed getter', t => {
    const typedMetadata = withReboxFields({
        ...metadata,
        ctorTypes: insert(ordString)('Data_Tuple.Tuple')({
            vars: ['a', 'b'], fields: [new C.TypeVar('a'), new C.TypeVar('b')],
        })(emptyMap),
    });
    const arrays = new C.ADT('Data.Tuple.Tuple', ['Data', 'Tuple', 'Tuple'], [new C.Array(C.Int.value), new C.Array(C.Int.value)]);
    const fnType = new C.Func([arrays], arrays);
    const producer = CodeGen.translateWithFunctions(typedMetadata)(module([
        ['nativeIdentity', typed(fnType, lambda([['input', 0]], typed(arrays, local('input', 0))))],
    ]));
    const consumer = CodeGen.translateWithFunctions({ ...typedMetadata, globalFunctions: producer.functions })({
        ...module([
            ['nativeIdentity', typed(fnType, lambda([['input', 0]], typed(arrays, app(
                expr(new S.Var(new C.Qualified(new Just('CompositeFlow'), 'nativeIdentity'))), local('input', 0)))))],
        ]), name: 'CompositeConsumer',
    });
    const worker = consumer.file.declarationGroups.flat().find(decl => decl instanceof Go.GoFunctionDecl
        && decl.value0.name.name === 'Call_CompositeConsumer_nativeIdentity');
    assert.ok(worker);
    assert.doesNotMatch(printGoExpr(worker.value0.body), /Rebox_|gopurs_runtime\.Array|unboxed :=|boxed :=/);
    const directory = goModule(t, 'gopurs-composite-return-', { runtime: runtimeGoCode });
    writeGoFile(directory, 'purescript/CompositeFlow.go', producer.code);
    writeGoFile(directory, 'purescript/CompositeConsumer.go', consumer.code);
    writeGoFile(directory, 'purescript/Tuple.go', `package purescript
type Constructor_Data_Tuple_Tuple[A, B any] struct { Rc uint32; V0 A; V1 B }
`);
    writeGoFile(directory, 'returns_test.go', `package main
import ("testing"; "unsafe"; "slices"; r "gopurs/output/gopurs_runtime"; p "gopurs/output/purescript")
func TestNativeAndBoxed(t *testing.T) {
    if p.Call_CompositeConsumer_nativeIdentity(nil) != nil { t.Fatal("native nil changed") }
    native := &p.Constructor_Data_Tuple_Tuple[[]int64, []int64]{V0: []int64{1, 7}, V1: []int64{-9}}
    if p.Call_CompositeConsumer_nativeIdentity(native) != native { t.Fatal("direct imported return reconstructed the payload") }
    boxed := &p.Constructor_Data_Tuple_Tuple[r.Value, r.Value]{V0: r.Array([]r.Value{r.Int(1), r.Int(7)}), V1: r.Array([]r.Value{r.Int(-9)})}
    input := r.Value{Type: r.TypeConstructor, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(boxed)}
    result := r.Apply(p.Get_CompositeConsumer_nativeIdentity(), input)
    if result.Type != input.Type || result.IntVal != input.IntVal { t.Fatal("public getter tag changed") }
    tuple := (*p.Constructor_Data_Tuple_Tuple[r.Value, r.Value])(result.UnsafePtr)
    if !slices.Equal(*(*[]r.Value)(tuple.V0.UnsafePtr), *(*[]r.Value)(boxed.V0.UnsafePtr)) || !slices.Equal(*(*[]r.Value)(tuple.V1.UnsafePtr), *(*[]r.Value)(boxed.V1.UnsafePtr)) { t.Fatal("public getter values changed") }
}
`);
    runGo(directory, ['test', '-count=1', '.']);
});

test('boxed composite lets and local callback results avoid reconstruction and retain immutable aliases', t => {
    const payload = local('payload', 1);
    const helper = typed(new C.Func([C.Any.value], state), lambda([['ignored', 4]],
        typed(state, app(local('transform', 0), typed(state, local('saved', 2))))));
    const call = lambda([['transform', 0], ['payload', 1]],
        letIn('saved', 2, typed(state, app(local('transform', 0), payload)),
            letIn('forward', 3, helper, app(local('forward', 3), payload))));
    const capture = lambda([['payload', 1]], letIn('saved', 2, typed(state, payload),
        typed(new C.Func([C.Any.value], state), lambda([['ignored', 4]], typed(state, local('saved', 2))))));
    const update = lambda([['payload', 1]], letIn('saved', 2,
        typed(state, expr(new S.Update(typed(state, payload), [new C.Prop('consumed', typed(C.Boolean.value, expr(new S.Lit(new C.LitBoolean(true)))))]))),
        typed(state, local('saved', 2))));
    const code = CodeGen.translate(metadata)(module([['call', call], ['capture', capture], ['update', update]]));
    assert.match(code, /var Call_local_CompositeFlow_forward_\w+ func\(gopurs_runtime.Value\) gopurs_runtime.Value/);
    assert.doesNotMatch(code, /unboxed :=|boxed :=|clone\.(?:consumed|errors|stream)|Rebox_CompositeFlow/);
    const directory = goModule(t, 'gopurs-composite-flow-', { runtime: runtimeGoCode });
    writeGoFile(directory, 'purescript/CompositeFlow.go', code);
    writeGoFile(directory, 'flow_test.go', `package main
import ("testing"; "math"; r "gopurs/output/gopurs_runtime"; p "gopurs/output/purescript")
func TestCompositeFlow(t *testing.T) {
    stream := r.Array([]r.Value{r.Int(7)})
    position := r.RecordDict2("column", "line", r.Int(math.MinInt64), r.Int(math.MaxInt64))
    issue := r.RecordDict2("error", "position", r.Str("λ😀"), position)
    for _, errors := range [][]r.Value{nil, {}, {issue}} {
        for _, consumed := range []bool{false, true} {
            state := r.RecordDict3("consumed", "errors", "stream", r.Bool(consumed), r.Array(errors), stream)
            calls := 0
            transform := r.Func(func(value r.Value) r.Value { calls++; return value })
            got := r.Apply2(p.Get_CompositeFlow_call(), transform, state)
            if calls != 2 || got != state { t.Fatal("callback count or immutable state identity changed") }
            closure := r.Apply(p.Get_CompositeFlow_capture(), state)
            for i := 0; i < 2; i++ { if r.Apply(closure, r.Value{}) != state { t.Fatal("captured alias changed") } }
            updated := r.Apply(p.Get_CompositeFlow_update(), state)
            if r.RecordGet(updated, "consumed").IntVal != 1 || r.RecordGet(state, "consumed").IntVal != r.Bool(consumed).IntVal { t.Fatal("persistent update changed original") }
            if r.RecordGet(updated, "errors") != r.RecordGet(state, "errors") || r.RecordGet(updated, "stream") != stream { t.Fatal("unmodified fields were reconstructed") }
        }
    }
}
`);
    runGo(directory, ['test', '-count=1', '.']);
});
