import assert from 'node:assert/strict';
import test from 'node:test';
import { Just, Nothing } from '../output/Data.Maybe/index.js';
import { Tuple } from '../output/Data.Tuple/index.js';
import { TcoExpr } from '../output/PureScript.Backend.Optimizer.Codegen.Tco/index.js';
import * as C from '../output/PureScript.Backend.Optimizer.CoreFn/index.js';
import * as S from '../output/PureScript.Backend.Optimizer.Syntax/index.js';
import { emit, totalIntegerMaybe } from '../output/Gopurs.ArrayTraverseTotal/index.js';

const e = syntax => new TcoExpr(null, syntax);
const typed = (type, value) => e(new S.Typed(type, value));
const int = value => e(new S.Lit(new C.LitInt(value)));
const bool = value => e(new S.Lit(new C.LitBoolean(value)));
const local = (level, type) => {
    const value = e(new S.Local(new Just('value'), level));
    return type ? typed(type, value) : value;
};
const qualified = (owner, name) => new C.Qualified(new Just(owner), name);
const ctor = (name, fields = [], owner = 'Data.Maybe') => e(new S.CtorSaturated(
    qualified(owner, name), C.SumType.value, 'Maybe', name, fields.map(value => new Tuple('value0', value))));
const just = value => ctor('Just', [value]);
const nothing = ctor('Nothing');
const unary = (op, value) => e(new S.PrimOp(new S.Op1(op, value)));
const binary = (op, left, right) => e(new S.PrimOp(new S.Op2(op, left, right)));
const choose = (condition, yes, no) => e(new S.Branch([new S.Pair(condition, yes)], no));
const add = (a, b) => binary(new S.OpIntNum(S.OpAdd.value), a, b);
const eq = (a, b) => binary(new S.OpIntOrd(S.OpEq.value), a, b);
const accepted = body => assert.equal(totalIntegerMaybe([2])(body), true);
const rejected = body => assert.equal(totalIntegerMaybe([2])(body), false);

test('total Maybe proof admits constructors, total primitive arithmetic and evaluated captures', () => {
    accepted(just(int(7))); accepted(nothing); accepted(just(local(2)));
    accepted(just(add(local(2), local(3, C.Int.value))));
    for (const operation of [S.OpAdd, S.OpSubtract, S.OpMultiply]) {
        accepted(just(binary(new S.OpIntNum(operation.value), local(2), int(-2147483648))));
    }
    for (const operation of [S.OpIntNegate, S.OpIntBitNot]) accepted(just(unary(operation.value, local(2))));
    const selected = choose(eq(local(2), local(3, C.Int.value)), nothing, just(add(local(2), int(1))));
    accepted(selected);
    accepted(choose(binary(S.OpBooleanAnd.value, bool(true), unary(S.OpBooleanNot.value, local(4, C.Boolean.value))), selected, just(int(0))));
    accepted(choose(binary(new S.OpBooleanOrd(S.OpEq.value), bool(false), bool(true)), nothing, just(int(1))));
});

test('total Maybe proof rejects partial or unknown work even under an unreachable branch', () => {
    const unknown = e(new S.App(e(new S.Var(qualified('Foreign', 'run'))), [local(2)]));
    const divide = binary(new S.OpIntNum(S.OpDivide.value), int(1), int(0));
    const mod = binary(new S.OpIntNum(S.OpMod.value), int(1), int(0));
    const arrayIndex = binary(S.OpArrayIndex.value, local(3), int(0));
    for (const partial of [unknown, divide, mod, arrayIndex, local(99), local(3, C.String.value)]) {
        rejected(just(partial));
        rejected(choose(bool(true), nothing, just(partial)));
        rejected(choose(eq(partial, int(0)), nothing, just(int(1))));
    }
    rejected(ctor('Just', [int(1)], 'Other.Maybe'));
    rejected(ctor('Nothing', [], 'Other.Maybe'));
    rejected(ctor('Just', [])); rejected(ctor('Nothing', [int(1)]));
    rejected(choose(local(3), nothing, just(int(1))));
    rejected(e(new S.Let(new Just('temporary'), 5, unknown, just(local(5, C.Int.value)))));
});

test('total path requires full application, exact TAST callback type and a nonrecursive context', () => {
    const maybeInt = new C.ADT('Data.Maybe.Maybe', ['Data', 'Maybe', 'Maybe'], [C.Int.value]);
    const callback = typed(new C.Func([C.Int.value], maybeInt), e(new S.Abs([new Tuple(new Just('value'), 2)], just(local(2)))));
    const array = local(3, new C.Array(C.Int.value));
    const neverTranslate = () => { throw new Error('rejected candidate must not translate arguments'); };
    const check = (context, supplied) => assert.deepEqual(emit(neverTranslate)(context)(0)(false)(0)(supplied), Nothing.value);
    check({ recVars: [] }, [callback]);
    check({ recVars: ['recursiveValue'] }, [callback, array]);
    check({ recVars: [] }, [typed(new C.Func([C.Any.value], maybeInt), callback.value1.value1), array]);
    check({ recVars: [] }, [typed(new C.Func([C.Int.value], C.Int.value), callback.value1.value1), array]);
    check({ recVars: [] }, [typed(new C.Func([C.Int.value], maybeInt), local(4)), array]);
});
