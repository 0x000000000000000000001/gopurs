import assert from 'node:assert/strict';
import test from 'node:test';
import { withReboxFields } from './codegen-metadata.mjs';
import { empty as emptyMap, insert, lookup } from '../output/Data.Map/index.js';
import { Just, Nothing } from '../output/Data.Maybe/index.js';
import { ordString } from '../output/Data.Ord/index.js';
import { empty as emptySet, singleton } from '../output/Data.Set/index.js';
import { Tuple } from '../output/Data.Tuple/index.js';
import * as Ref from '../output/Effect.Ref/index.js';
import * as C from '../output/PureScript.Backend.Optimizer.CoreFn/index.js';
import * as S from '../output/PureScript.Backend.Optimizer.Syntax/index.js';
import { TcoExpr } from '../output/PureScript.Backend.Optimizer.Codegen.Tco/index.js';
import * as Go from '../output/Gopurs.GoAst/index.js';
import { generated, SymbolReference } from '../output/Gopurs.GoSymbol/index.js';
import { StmtEmpty, StmtLeaf, flattenStmts } from '../output/Gopurs.ExprContext/index.js';
import { printGoExpr } from '../output/Gopurs.Printer/index.js';
import { application } from '../output/Gopurs.CallExprs/index.js';
import { worker, workerArguments } from '../output/Gopurs.ListFilter/index.js';

const map = entries => entries.reduce((result, [key, value]) => insert(ordString)(key)(value)(result), emptyMap);
const expr = syntax => new TcoExpr(null, syntax);
const typed = (ty, value) => expr(new S.Typed(ty, value));
const local = (name, level, ty) => typed(ty, expr(new S.Local(new Just(name), level)));
const apply = (ty, fn, ...args) => typed(ty, expr(new S.App(fn, args)));
const lambda = (ty, args, body) => typed(ty, expr(new S.Abs(args.map(([name, level]) => new Tuple(new Just(name), level)), body)));
const qualified = (module, name) => new C.Qualified(new Just(module), name);

function fixture(options = {}) {
    const module = options.module ?? 'BuilderProbe';
    const node = options.node ?? 'Push', empty = options.empty ?? 'Stop';
    const names = Go.constructorNames(module)(node), leaf = Go.constructorNames(module)(empty);
    const element = options.element ?? C.Int.value;
    const list = new C.ADT(`${module}.Chain`, [module, 'Chain'], []);
    const native = Go.structPointer({ ...names, fullName: `${module}.Chain` })([]);
    const predType = new C.Func([element], C.Boolean.value);
    const loopType = new C.Func([list, list], list);
    const functionType = new C.Func([predType, list], list);
    const ctor = (name, ...values) => typed(list, expr(new S.CtorSaturated(qualified(module, name), C.SumType.value,
        'Chain', name, values.map((value, index) => new Tuple(`value${index}`, value)))));
    const field = (value, index, owner = module) => typed(index === 0 ? element : list, expr(new S.Accessor(value,
        new S.GetCtorField(qualified(owner, node), C.SumType.value, 'Chain', node, `value${index}`, index))));
    const tag = (name, value) => typed(C.Boolean.value, expr(new S.PrimOp(new S.Op1(new S.OpIsTag(qualified(module, name)), value))));
    const cases = (input, nil, cons) => typed(list, expr(new S.Branch([
        new S.Pair(tag(empty, input), nil), new S.Pair(tag(node, input), cons),
    ], expr(new S.Fail('impossible')))));
    const pred = local('predicate', 0, predType), source = local('source', 1, list);
    const input = local('input', 3, list), acc = local('acc', 4, list);
    const current = local('current', 6, list), reversed = local('reversed', 7, list);
    const reverseSelf = local('reverse', 5, loopType);
    const reverseHead = field(current, options.reverseField ?? 0, options.reverseOwner ?? module);
    const reverseStep = apply(list, reverseSelf, field(current, 1), ctor(node, reverseHead, reversed));
    const reverseLoop = lambda(options.reverseType?.(loopType) ?? loopType, [['current', 6], ['reversed', 7]],
        cases(current, reversed, reverseStep));
    const reverse = typed(list, expr(new S.LetRec(5, [new Tuple('reverse', reverseLoop)],
        apply(list, reverseSelf, acc, options.reverseSeed ?? ctor(empty)))));
    const self = local('walk', 2, loopType);
    const condition = options.predicateCall?.(pred, field(input, 0)) ?? apply(C.Boolean.value, pred, field(input, 0));
    const yes = apply(list, self, field(input, 1), ctor(node, field(input, 0), acc));
    const no = options.skip?.(self, input, acc) ?? apply(list, self, field(input, 1), acc);
    const choice = typed(list, expr(new S.Branch([new S.Pair(condition, yes)], no)));
    const loop = lambda(options.loopType?.(loopType) ?? loopType, [['input', 3], ['acc', 4]],
        cases(input, options.noReverse ? acc : reverse, choice));
    const bindings = [new Tuple('walk', loop)];
    if (options.initializer) bindings.push(new Tuple('extra', options.initializer));
    const body = typed(list, expr(new S.LetRec(2, bindings,
        apply(list, self, source, options.initialSeed ?? ctor(empty)))));
    const value = lambda(options.functionType?.(functionType) ?? functionType, [['predicate', 0], ['source', 1]], body);
    const metadata = withReboxFields({
        elidedCtors: emptySet,
        ctorTypes: map([[`${module}.${node}`, { vars: [], fields: options.fields ?? [element, list] }],
            [`${module}.${empty}`, { vars: [], fields: [] }]]),
        pointerAdtPaths: map([[`${module}.Chain`, { ctorName: node, arity: 0, sharedPairs: [] }]]),
        pointerAdtNodes: singleton(names.baseStructName),
        pointerAdtLeaves: options.noLeaf ? emptyMap : map([[leaf.baseStructName,
            { nodeBaseStruct: names.baseStructName, nodeCtor: node }]]),
        enumAdts: emptySet, enumCtors: emptySet, globalTypes: emptyMap,
        globalFunctions: emptyMap, classDeclsFields: emptyMap,
    });
    return { metadata, value, result: options.result ?? native, params: options.params ?? [
        new Tuple('predicate_0', Go.TypeValue.value), new Tuple('source_1', native),
    ] };
}

const lower = ({ metadata, params, result, value }) => worker(metadata)(params)(result)(value);
const signature = ({ metadata, params, result, value }) => workerArguments(metadata)(params)(result)(value);
const nativePredicate = new Go.TypeFunc([Go.TypeInt64.value], Go.TypeBool.value);
function visit(node, callback) {
    if (!node || typeof node !== 'object') return;
    callback(node);
    for (const value of Object.values(node)) visit(value, callback);
}

test('native filter/reverse lowers structurally with private output links and the boxed predicate ABI', () => {
    for (const options of [{}, { module: 'Other', node: 'Cell', empty: 'Done' }]) {
        const input = fixture(options), before = JSON.stringify(input);
        const result = lower(input);
        assert.ok(result instanceof Just);
        assert.equal(JSON.stringify(input), before);
        const mutations = [], constructors = [], callbacks = [];
        visit(result.value0.expr, node => {
            if (node instanceof Go.GoMutate) mutations.push(node.value0);
            if (node instanceof Go.GoConstructor) constructors.push(node);
            if (node instanceof Go.GoSelector && node.value1 === 'Apply') callbacks.push(node);
        });
        assert.deepEqual([...new Set(mutations.filter(name => name.includes('.')))], ['tail.V1'],
            'only the private output tail may have a field updated');
        assert.equal(constructors.length, 1, 'one constructor site for retained prefix cells');
        assert.equal(callbacks.length, 2, 'one boxed callback site in each disjoint scanning phase');
        assert.equal(result.value0.expr.value0.constructor, Go.GoFuncBlock, 'private builder scope');
        assert.deepEqual(result.value0.expr.value1.map(node => node.value0), ['predicate_0', 'source_1']);
    }
});

test('suffix sharing mutates only fresh nodes and never calls the predicate during prefix copying', () => {
    for (const native of [false, true]) {
        const input = fixture({ module: 'AliasProbe', node: 'Cell', empty: 'Done' });
        if (native) input.params[0] = new Tuple(input.params[0].value0, nativePredicate);
        const result = lower(input);
        assert.ok(result instanceof Just);
        const copies = [], scans = [], tailUpdates = [], allocations = [];
        visit(result.value0.expr, node => {
            if (node instanceof Go.GoForRange) {
                if (node.value0 === 'run != cursor') copies.push(node);
                else scans.push(node);
            }
            if (node instanceof Go.GoMutate && node.value0 === 'tail') tailUpdates.push(node.value1);
            if (node instanceof Go.GoAssign && node.value0 === 'node') allocations.push(node.value1);
        });
        assert.equal(copies.length, 1);
        assert.equal(copies[0].value0, 'run != cursor', 'copy only the retained run preceding a rejection');
        assert.equal(scans.length, 2);
        for (const scan of scans) {
            assert.equal(scan.value0, 'cursor != nil');
            assert.ok(scan.value1[0] instanceof Go.GoIfElse);
            assert.equal(printGoExpr(scan.value1[0].value1[0]), 'break');
            assert.deepEqual(scan.value1[1], new Go.GoMutate('cursor', new Go.GoStructAccess(new Go.GoVar('cursor'), 'V1')));
        }
        assert.ok(scans[1].value1[0].value0 instanceof Go.GoPrefixOp);
        assert.equal(scans[1].value1[0].value0.value0, '!');
        assert.deepEqual(scans[1].value1[0].value0.value1, scans[0].value1[0].value0,
            'the two phases stop on opposite outcomes of the same predicate');
        assert.equal(allocations.length, 1);
        assert.ok(allocations[0] instanceof Go.GoConstructor);
        assert.deepEqual(tailUpdates, [new Go.GoVar('node')], 'tail can only acquire a freshly allocated cell');
        const copyCalls = [], copiedFields = [];
        visit(copies[0], node => {
            if (node instanceof Go.GoCall) copyCalls.push(node);
            if (node instanceof Go.GoStructAccess) copiedFields.push([node.value0.value0, node.value1]);
        });
        assert.deepEqual(copyCalls, [], 'prefix copying does not re-evaluate either callback ABI');
        assert.deepEqual(copiedFields, [['run', 'V0'], ['run', 'V1']]);
        const body = result.value0.expr.value0.value1;
        const traversal = body.find(node => node instanceof Go.GoFor);
        const steps = traversal.value1;
        const secondScan = steps.indexOf(scans[1]), copy = steps.indexOf(copies[0]);
        const exhausted = steps[steps.indexOf(scans[0]) + 1];
        assert.ok(exhausted instanceof Go.GoIfElse);
        assert.deepEqual(exhausted.value1, [new Go.GoReturn(new Go.GoVar('output'))],
            'an exhausted rejected suffix returns the already built prefix');
        assert.deepEqual(steps[secondScan - 2], new Go.GoAssign('run', new Go.GoVar('cursor')));
        assert.deepEqual(steps[secondScan - 1], new Go.GoMutate('cursor', new Go.GoStructAccess(new Go.GoVar('cursor'), 'V1')),
            'consume the first accepted element before the retained scan');
        assert.deepEqual(steps[copy + 1], new Go.GoMutate('cursor', new Go.GoStructAccess(new Go.GoVar('cursor'), 'V1')),
            'consume the first rejected element before repeating the rejected scan');
        const end = steps[secondScan + 1];
        assert.ok(end instanceof Go.GoIfElse);
        assert.ok(end.value0 instanceof Go.GoBinOp);
        assert.equal(end.value0.value0, '==');
        assert.deepEqual(end.value0.value1, new Go.GoVar('cursor'));
        assert.match(printGoExpr(end.value0.value2), /\(nil\)$/);
        const finish = printGoExpr(new Go.GoBlock(end.value1));
        assert.match(finish, /return run/);
        assert.match(finish, /tail\.V1 = run/);
        assert.match(finish, /return output/);
        assert.deepEqual(result.value0.expr.value1.map(node => node.value0), ['predicate_0', 'source_1'],
            'the private builder still receives both already evaluated arguments');
    }
});

test('the admitted signature publishes a native predicate and the builder consumes it directly', () => {
    const input = fixture();
    const published = signature(input);
    assert.ok(published instanceof Just);
    assert.deepEqual(published.value0, [nativePredicate, input.result]);
    const nativeInput = { ...input, params: input.params.map((param, index) =>
        new Tuple(param.value0, published.value0[index])) };
    const lowered = lower(nativeInput);
    assert.ok(lowered instanceof Just);
    const text = printGoExpr(lowered.value0.expr);
    assert.match(text, /func\(predicate func\(int64\) bool/);
    assert.match(text, /predicate\(cursor.V0\)/);
    assert.doesNotMatch(text, /Apply|gopurs_runtime.Int\(/);
});

for (const [label, options] of [
    ['a quantified function', { functionType: ty => new C.ForAll(['a'], ty) }],
    ['a quantified filter loop', { loopType: ty => new C.ForAll(['a'], ty) }],
    ['a quantified reverse loop', { reverseType: ty => new C.ForAll(['a'], ty) }],
    ['non-Int elements', { element: C.String.value }],
    ['missing nil layout', { noLeaf: true }],
    ['incorrect constructor fields', { fields: [C.Int.value, C.Int.value] }],
    ['an erased result ABI', { result: Go.TypeValue.value }],
    ['incorrect parameter names', { params: [new Tuple('x', Go.TypeValue.value), new Tuple('y', Go.TypeValue.value)] }],
    ['a missing reverse', { noReverse: true }],
    ['a reverse projecting the wrong field', { reverseField: 1 }],
    ['another module constructor', { reverseOwner: 'Foreign' }],
    ['a nonempty initial seed', { initialSeed: local('shared', 99, C.Any.value) }],
    ['a nonempty reverse seed', { reverseSeed: local('shared', 99, C.Any.value) }],
    ['an evaluated recursive initializer', { initializer: apply(C.Any.value, expr(new S.Var(qualified('Foreign', 'run'))), local('x', 99, C.Any.value)) }],
    ['opaque predicate substitutions', { predicateCall: (_, head) => apply(C.Boolean.value, expr(new S.Var(qualified('Foreign', 'test'))), head) }],
    ['extra predicate evaluations', { predicateCall: (pred, head) => typed(C.Boolean.value, expr(new S.PrimOp(new S.Op2(S.OpBooleanAnd.value,
        apply(C.Boolean.value, pred, head), apply(C.Boolean.value, pred, head))))) }],
    ['a predicate applied to a captured value', { predicateCall: (pred) => apply(C.Boolean.value, pred, local('captured', 99, C.Int.value)) }],
    ['a non-progressing skip branch', { skip: (self, input, acc) => apply(C.Any.value, self, input, acc) }],
]) {
    test(`native filter builder and signature reject ${label}`, () => {
        const input = fixture(options);
        assert.equal(lower(input), Nothing.value);
        assert.equal(signature(input), Nothing.value);
    });
}

function callFixture({ imported = false, opaque = false, partial = false, admitted = true, boxedResult = false } = {}) {
    const input = fixture(), owner = imported ? 'OtherModule' : 'BuilderProbe';
    const functionType = new C.Func([new C.Func([C.Int.value], C.Boolean.value), C.Any.value], C.Any.value);
    const fn = typed(functionType, expr(new S.Var(qualified(owner, 'select'))));
    const body = typed(C.Boolean.value, expr(new S.Var(qualified('Probe', 'observe'))));
    const callback = opaque ? typed(new C.Func([C.Int.value], C.Boolean.value),
        expr(new S.Var(qualified('Probe', 'predicateFactory'))))
        : lambda(new C.Func([C.Int.value], C.Boolean.value), [['item', 2]], body);
    const source = typed(C.Any.value, expr(new S.Var(qualified('Probe', 'sourceFactory'))));
    const reference = new SymbolReference(generated(owner)(`Call_${owner}_select`));
    const info = { reference, arity: 2, fArgs: [admitted ? nativePredicate : Go.TypeValue.value, input.result], fRet: input.result };
    const context = {
        metadata: { ...input.metadata, globalFunctions: imported ? map([[`${owner}.select`, info]]) : emptyMap },
        moduleFunctions: imported ? emptyMap : map([['select', info]]),
        codegenStateRef: Ref.new({ declarations: [], globalId: 0, reboxPairs: emptySet })(),
        depth: 0, modNameStr: 'BuilderProbe', recVars: [], bound: emptyMap,
        tcoIdent: Nothing.value, loopCtx: [], options: { isTail: false, inEffectBlock: false },
        mbExpectedExprType: Nothing.value, mbExpectedGoType: Nothing.value,
    };
    const translated = [];
    const translate = ctx => nextId => value => {
        translated.push(value);
        if (value === body) {
            const binding = lookup(ordString)('item_2')(ctx.bound);
            assert.ok(binding instanceof Just);
            assert.equal(binding.value0.goType, Go.TypeInt64.value);
            return { stmts: StmtEmpty.value, expr: new Go.GoCall(new Go.GoVar('observe'), [new Go.GoVar('item_2')]),
                exprType: boxedResult ? Go.TypeValue.value : Go.TypeBool.value, nextId };
        }
        if (value === source) return { stmts: new StmtLeaf(new Go.GoCall(new Go.GoVar('sourceSetup'), [])),
            expr: new Go.GoCall(new Go.GoVar('sourceFactory'), []), exprType: input.result, nextId };
        if (value === callback) return { stmts: StmtEmpty.value,
            expr: new Go.GoCall(new Go.GoVar('predicateFactory'), []), exprType: Go.TypeValue.value, nextId };
        if (value === fn) return { stmts: StmtEmpty.value,
            expr: new Go.GoVar('boxedSelect'), exprType: Go.TypeValue.value, nextId };
        assert.fail('unexpected translation');
    };
    const value = apply(C.Any.value, fn, ...partial ? [callback] : [callback, source]);
    const result = application(translate)(context)(0)(value);
    const text = [...flattenStmts(result.stmts), new Go.GoReturn(result.expr)].map(printGoExpr).join('\n');
    return { result, text, translated, callback, body, source, reference };
}

for (const imported of [false, true]) {
    test(`native predicate calls retain ${imported ? 'imported' : 'local'} worker symbols and capture order`, () => {
        const result = callFixture({ imported });
        assert.deepEqual(result.translated, [result.body, result.source]);
        assert.match(result.text, /func\(item_2 int64\) bool/);
        assert.doesNotMatch(result.text, /Apply|predicateFactory/);
        assert.ok(result.text.indexOf('nativePredicate_arg_0 :=') < result.text.indexOf('sourceSetup()'));
        const references = [];
        visit(result.result.expr, node => { if (node instanceof Go.GoRef) references.push(node.value0); });
        assert.deepEqual(references, [result.reference]);
    });
}

test('a native predicate with an opaque Boolean result converts only its return value', () => {
    const result = callFixture({ boxedResult: true });
    assert.deepEqual(result.translated, [result.body, result.source]);
    assert.match(result.text, /observe\(item_2\).IntVal/);
    assert.doesNotMatch(result.text, /predicateFactory|Apply/);
});

test('an opaque predicate producer is captured once before source statements and adapted once', () => {
    const result = callFixture({ opaque: true });
    assert.deepEqual(result.translated, [result.callback, result.source]);
    assert.equal(result.text.match(/predicateFactory\(\)/g)?.length, 1);
    assert.match(result.text, /func\(boxedPredicate gopurs_runtime.Value\) func\(int64\) bool/);
    assert.match(result.text, /Apply\(boxedPredicate, gopurs_runtime.Int\(predicateInput\)\)/);
    assert.ok(result.text.indexOf('predicateFactory()') < result.text.indexOf('sourceSetup()'));
});

test('a partially applied filter keeps its boxed getter and callback', () => {
    const result = callFixture({ partial: true });
    assert.ok(!result.translated.includes(result.body));
    assert.match(result.text, /Apply\(boxedSelect, predicateFactory\(\)\)/);
    assert.doesNotMatch(result.text, /nativePredicate|func\(int64\) bool|Call_BuilderProbe_select/);
});

test('an ordinary callback slot does not enable native predicate lowering', () => {
    const result = callFixture({ admitted: false });
    assert.deepEqual(result.translated, [result.callback, result.source]);
    assert.doesNotMatch(result.text, /nativePredicate|func\(int64\) bool|boxedPredicate/);
});
