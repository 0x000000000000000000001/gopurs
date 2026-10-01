import assert from 'node:assert/strict';
import test from 'node:test';
import { empty as emptyMap, insert, lookup } from '../output/Data.Map/index.js';
import { Just, Nothing } from '../output/Data.Maybe/index.js';
import { ordString } from '../output/Data.Ord/index.js';
import { empty as emptySet, toUnfoldable } from '../output/Data.Set/index.js';
import { Tuple } from '../output/Data.Tuple/index.js';
import { unfoldableArray } from '../output/Data.Unfoldable/index.js';
import * as Ref from '../output/Effect.Ref/index.js';
import * as Constructors from '../output/Gopurs.ConstructorExprs/index.js';
import { StmtEmpty, StmtLeaf, flattenStmts } from '../output/Gopurs.ExprContext/index.js';
import * as Go from '../output/Gopurs.GoAst/index.js';
import * as Literals from '../output/Gopurs.LiteralExprs/index.js';
import { printGoExpr } from '../output/Gopurs.Printer/index.js';
import { TcoExpr } from '../output/PureScript.Backend.Optimizer.Codegen.Tco/index.js';
import * as C from '../output/PureScript.Backend.Optimizer.CoreFn/index.js';
import { localId } from '../output/PureScript.Backend.Optimizer.FreeVars/index.js';
import * as S from '../output/PureScript.Backend.Optimizer.Syntax/index.js';
import { withReboxFields } from './codegen-metadata.mjs';

const map = entries => entries.reduce((result, [key, value]) => insert(ordString)(key)(value)(result), emptyMap);
const adt = (name, args = []) => new C.ADT(`Probe.${name}`, ['Probe', name], args);
const boxInt = adt('Box', [C.Int.value]);
const metadata = withReboxFields({
    ctorTypes: map([
        ['Probe.Box', { vars: ['a'], fields: [new C.TypeVar('a')] }],
        ['Probe.Bundle', { vars: [], fields: [boxInt, boxInt] }],
        ['Probe.Cell', { vars: [], fields: [C.Int.value, C.Int.value] }],
    ]),
    pointerAdtPaths: map([
        ['Probe.Box', { ctorName: 'Box', arity: 1 }],
        ['Probe.Bundle', { ctorName: 'Bundle', arity: 0 }],
        ['Probe.Cell', { ctorName: 'Cell', arity: 0 }],
    ]),
    pointerAdtNodes: emptySet, pointerAdtLeaves: emptyMap, elidedCtors: emptySet,
    enumAdts: emptySet, enumCtors: emptySet, classDeclsFields: emptyMap,
    globalTypes: emptyMap, globalFunctions: emptyMap,
});
const pointer = (name, args = []) => Go.structPointer({
    baseStructName: `Data_Probe_${name}`, fullName: `Probe.${name}`, structName: `Constructor_Probe_${name}`,
})(args);
const boxedPointer = pointer('Box', [Go.TypeValue.value]);
const intPointer = pointer('Box', [Go.TypeInt64.value]);
const node = syntax => new TcoExpr(null, syntax);
const variable = name => node(new S.Var(new C.Qualified(Nothing.value, name)));
const recordType = fields => new C.Record(new C.Row(fields.map(([key, type]) => new Tuple(key, type)), Nothing.value));
const context = expected => ({
    metadata, codegenStateRef: Ref.new({ declarations: [], globalId: 0, reboxPairs: emptySet })(),
    depth: 4, modNameStr: 'Probe', recVars: ['recursive'], moduleFunctions: emptyMap, bound: emptyMap,
    tcoIdent: new Just('tail'), loopCtx: [{ sentinel: true }], options: { isTail: true, inEffectBlock: true },
    mbExpectedExprType: expected,
});
const pairCount = ctx => toUnfoldable(unfoldableArray)(Ref.read(ctx.codegenStateRef)().reboxPairs).length;
const statements = result => flattenStmts(result.stmts).map(printGoExpr);
const assertChild = (parent, child) => {
    assert.equal(child.depth, parent.depth + 1);
    assert.equal(child.metadata, parent.metadata);
    assert.equal(child.codegenStateRef, parent.codegenStateRef);
    assert.equal(child.moduleFunctions, parent.moduleFunctions);
    assert.equal(child.recVars, parent.recVars);
    assert.deepEqual(child.options, { isTail: false, inEffectBlock: false });
    assert.ok(child.tcoIdent instanceof Nothing);
    assert.ok(child.mbExpectedExprType instanceof Nothing);
    assert.deepEqual(child.loopCtx, []);
};
function observingTranslator(ctx, events) {
    return child => nextId => value => {
        assertChild(ctx, child);
        const name = value.value1.value0.value1;
        events.push({ name, nextId, pairs: pairCount(ctx) });
        return { stmts: new StmtLeaf(Go.rawGo(`mark_${name}()`)), expr: new Go.GoVar(name),
            exprType: boxedPointer, nextId: nextId + 1 };
    };
}

test('array elements finish translation before their common layout registers Rebox conversions', () => {
    const ctx = context(new Just(new C.Array(boxInt)));
    const events = [];
    const result = Literals.array(observingTranslator(ctx, events))(ctx)(7)([variable('first'), variable('second')]);
    assert.deepEqual(events, [{ name: 'first', nextId: 7, pairs: 0 }, { name: 'second', nextId: 8, pairs: 0 }]);
    assert.equal(pairCount(ctx), 1);
    assert.deepEqual(result.exprType, new Go.TypeNativeArray(intPointer));
    assert.deepEqual(statements(result), ['mark_first()', 'mark_second()']);
    assert.equal(result.nextId, 9);
});

for (const kind of ['record', 'constructor']) {
    test(`${kind} fields are coerced before the following field, in their own declared order`, () => {
        const expected = kind === 'record' ? recordType([['z', boxInt], ['a', boxInt]]) : adt('Bundle');
        const ctx = context(new Just(expected));
        const events = [];
        const translate = observingTranslator(ctx, events);
        const result = kind === 'record'
            ? Literals.record(translate)(ctx)(7)(C.Any.value)([new C.Prop('z', variable('second')), new C.Prop('a', variable('first'))])
            : Constructors.saturated(translate)(ctx)(7)(variable('constructor'))(new Just('Probe'))('Bundle')
                ([new Tuple('z', variable('first')), new Tuple('a', variable('second'))]);
        assert.deepEqual(events, [{ name: 'first', nextId: 7, pairs: 0 }, { name: 'second', nextId: 8, pairs: 1 }]);
        assert.equal(pairCount(ctx), 1);
        assert.deepEqual(statements(result), ['mark_first()', 'mark_second()']);
        assert.equal(result.nextId, 9);
    });
}

test('array annotations specialize empty or uniform elements, while mixed representations remain boxed', () => {
    const cases = [
        { types: [], expected: Nothing.value, result: Go.TypeValue.value },
        { types: [], expected: new Just(new C.Array(C.Int.value)), result: new Go.TypeNativeArray(Go.TypeInt64.value) },
        { types: [Go.TypeInt64.value, Go.TypeInt64.value], expected: Nothing.value, result: new Go.TypeNativeArray(Go.TypeInt64.value) },
        { types: [Go.TypeInt64.value, Go.TypeValue.value], expected: new Just(new C.Array(C.Int.value)), result: Go.TypeValue.value },
        { types: [Go.TypeInt64.value], expected: new Just(new C.Array(C.Any.value)), result: Go.TypeValue.value },
    ];
    for (const entry of cases) {
        const ctx = context(entry.expected);
        let index = 0;
        const translate = child => nextId => _value => {
            assertChild(ctx, child);
            return { stmts: StmtEmpty.value, expr: new Go.GoVar(`item${index}`), exprType: entry.types[index++], nextId: nextId + 1 };
        };
        const result = Literals.array(translate)(ctx)(3)(entry.types.map((_, index) => variable(`item${index}`)));
        assert.deepEqual(result.exprType, entry.result);
        assert.equal(index, entry.types.length);
        assert.equal(result.nextId, 3 + entry.types.length);
    }
});

test('record and constructor callbacks share parameter typing without leaking locals between fields', () => {
    const functionType = new C.Func([C.Int.value], C.Int.value);
    const x = localId(new Just('x'))(0), y = localId(new Just('y'))(1);
    for (const kind of ['record', 'constructor']) {
        for (const Lambda of [S.Abs, S.UncurriedAbs]) {
            const ctx = context(new Just(kind === 'record' ? recordType([['a', functionType], ['z', C.Any.value]]) : adt('Callbacks')));
            ctx.metadata = withReboxFields({ ...metadata,
                ctorTypes: insert(ordString)('Probe.Callbacks')({ vars: [], fields: [functionType, C.Any.value] })(metadata.ctorTypes),
            });
            const lambda = node(new Lambda([new Tuple(new Just('x'), 0), new Tuple(new Just('y'), 1)], variable('body')));
            let visited = 0;
            const translate = child => nextId => _value => {
                assertChild(ctx, child);
                if (visited++ === 0) {
                    assert.deepEqual(lookup(ordString)(x)(child.bound), new Just({ name: x, goType: Go.TypeInt64.value }));
                    assert.deepEqual(lookup(ordString)(y)(child.bound), new Just({ name: y, goType: Go.TypeValue.value }));
                } else {
                    assert.ok(lookup(ordString)(x)(child.bound) instanceof Nothing);
                    assert.ok(lookup(ordString)(y)(child.bound) instanceof Nothing);
                }
                return { stmts: StmtEmpty.value, expr: new Go.GoVar('callback'), exprType: Go.TypeValue.value, nextId };
            };
            if (kind === 'record') {
                Literals.record(translate)(ctx)(0)(C.Any.value)([new C.Prop('a', lambda), new C.Prop('z', variable('other'))]);
            } else {
                Constructors.saturated(translate)(ctx)(0)(variable('constructor'))(new Just('Probe'))('Callbacks')
                    ([new Tuple('first', lambda), new Tuple('second', variable('other'))]);
            }
            assert.equal(visited, 2);
        }
    }
});

test('constructor reuse owns one identifier and is disabled by field statements or representation changes', () => {
    for (const mode of ['reuse', 'statement', 'boxed']) {
        const ctx = context(new Just(adt('Cell')));
        ctx.bound = map([['source', { name: 'source', goType: pointer('Cell') }]]);
        const replacement = node(new S.Lit(new C.LitInt(7)));
        const field = variable('field');
        const translate = _child => nextId => value => ({
            stmts: mode === 'statement' && value === replacement ? new StmtLeaf(Go.rawGo('observe()')) : StmtEmpty.value,
            expr: value === replacement ? (mode === 'boxed'
                ? new Go.GoCall(new Go.GoSelector(new Go.GoVar('gopurs_runtime'), 'Int'), [new Go.GoInt(7)]) : new Go.GoInt(7))
                : new Go.GoConstructorAccess(new Go.GoVar('source'), 'Constructor_Probe_Cell', [], 1, true),
            exprType: mode === 'boxed' && value === replacement ? Go.TypeValue.value : Go.TypeInt64.value,
            nextId: nextId + 1,
        });
        const result = Constructors.saturated(translate)(ctx)(8)(variable('constructor'))(new Just('Probe'))('Cell')
            ([new Tuple('first', replacement), new Tuple('second', field)]);
        if (mode === 'reuse') {
            assert.equal(result.nextId, 11);
            assert.deepEqual(result.expr, new Go.GoVar('__reuse_10'));
            assert.equal(flattenStmts(result.stmts).length, 2);
            assert.match(statements(result)[0], /^var __reuse_10 \*Constructor_Probe_Cell$/);
        } else {
            assert.equal(result.nextId, 10);
            assert.ok(result.expr instanceof Go.GoConstructor);
            assert.deepEqual(statements(result), mode === 'statement' ? ['observe()'] : []);
        }
    }
});
