import * as Control_Applicative from "../Control.Applicative/index.js";
import * as Control_Apply from "../Control.Apply/index.js";
import * as Data_Array_NonEmpty_Internal from "../Data.Array.NonEmpty.Internal/index.js";
import * as Data_Eq from "../Data.Eq/index.js";
import * as Data_Foldable from "../Data.Foldable/index.js";
import * as Data_Function from "../Data.Function/index.js";
import * as Data_Functor from "../Data.Functor/index.js";
import * as Data_Maybe from "../Data.Maybe/index.js";
import * as Data_Monoid from "../Data.Monoid/index.js";
import * as Data_Ord from "../Data.Ord/index.js";
import * as Data_Ordering from "../Data.Ordering/index.js";
import * as Data_Semigroup from "../Data.Semigroup/index.js";
import * as Data_Traversable from "../Data.Traversable/index.js";
import * as Data_Tuple from "../Data.Tuple/index.js";
import * as PureScript_Backend_Optimizer_CoreFn from "../PureScript.Backend.Optimizer.CoreFn/index.js";
var eqQualified = /* #__PURE__ */ PureScript_Backend_Optimizer_CoreFn.eqQualified(PureScript_Backend_Optimizer_CoreFn.eqIdent);
var ordQualified = /* #__PURE__ */ PureScript_Backend_Optimizer_CoreFn.ordQualified(PureScript_Backend_Optimizer_CoreFn.ordIdent);
var eqMaybe = /* #__PURE__ */ Data_Maybe.eqMaybe(PureScript_Backend_Optimizer_CoreFn.eqIdent);
var eqTuple = /* #__PURE__ */ Data_Tuple.eqTuple(Data_Eq.eqString);
var eqArray = /* #__PURE__ */ Data_Eq.eqArray(Data_Eq.eqString);
var eqTuple1 = /* #__PURE__ */ Data_Tuple.eqTuple(PureScript_Backend_Optimizer_CoreFn.eqIdent);
var Pair = /* #__PURE__ */ (function () {
    function Pair(value0, value1) {
        this.value0 = value0;
        this.value1 = value1;
    };
    Pair.create = function (value0) {
        return function (value1) {
            return new Pair(value0, value1);
        };
    };
    return Pair;
})();
var Level = function (x) {
    return x;
};
var OpEq = /* #__PURE__ */ (function () {
    function OpEq() {

    };
    OpEq.value = new OpEq();
    return OpEq;
})();
var OpNotEq = /* #__PURE__ */ (function () {
    function OpNotEq() {

    };
    OpNotEq.value = new OpNotEq();
    return OpNotEq;
})();
var OpGt = /* #__PURE__ */ (function () {
    function OpGt() {

    };
    OpGt.value = new OpGt();
    return OpGt;
})();
var OpGte = /* #__PURE__ */ (function () {
    function OpGte() {

    };
    OpGte.value = new OpGte();
    return OpGte;
})();
var OpLt = /* #__PURE__ */ (function () {
    function OpLt() {

    };
    OpLt.value = new OpLt();
    return OpLt;
})();
var OpLte = /* #__PURE__ */ (function () {
    function OpLte() {

    };
    OpLte.value = new OpLte();
    return OpLte;
})();
var OpAdd = /* #__PURE__ */ (function () {
    function OpAdd() {

    };
    OpAdd.value = new OpAdd();
    return OpAdd;
})();
var OpDivide = /* #__PURE__ */ (function () {
    function OpDivide() {

    };
    OpDivide.value = new OpDivide();
    return OpDivide;
})();
var OpMultiply = /* #__PURE__ */ (function () {
    function OpMultiply() {

    };
    OpMultiply.value = new OpMultiply();
    return OpMultiply;
})();
var OpSubtract = /* #__PURE__ */ (function () {
    function OpSubtract() {

    };
    OpSubtract.value = new OpSubtract();
    return OpSubtract;
})();
var OpArrayIndex = /* #__PURE__ */ (function () {
    function OpArrayIndex() {

    };
    OpArrayIndex.value = new OpArrayIndex();
    return OpArrayIndex;
})();
var OpBooleanAnd = /* #__PURE__ */ (function () {
    function OpBooleanAnd() {

    };
    OpBooleanAnd.value = new OpBooleanAnd();
    return OpBooleanAnd;
})();
var OpBooleanOr = /* #__PURE__ */ (function () {
    function OpBooleanOr() {

    };
    OpBooleanOr.value = new OpBooleanOr();
    return OpBooleanOr;
})();
var OpBooleanOrd = /* #__PURE__ */ (function () {
    function OpBooleanOrd(value0) {
        this.value0 = value0;
    };
    OpBooleanOrd.create = function (value0) {
        return new OpBooleanOrd(value0);
    };
    return OpBooleanOrd;
})();
var OpCharOrd = /* #__PURE__ */ (function () {
    function OpCharOrd(value0) {
        this.value0 = value0;
    };
    OpCharOrd.create = function (value0) {
        return new OpCharOrd(value0);
    };
    return OpCharOrd;
})();
var OpIntBitAnd = /* #__PURE__ */ (function () {
    function OpIntBitAnd() {

    };
    OpIntBitAnd.value = new OpIntBitAnd();
    return OpIntBitAnd;
})();
var OpIntBitOr = /* #__PURE__ */ (function () {
    function OpIntBitOr() {

    };
    OpIntBitOr.value = new OpIntBitOr();
    return OpIntBitOr;
})();
var OpIntBitShiftLeft = /* #__PURE__ */ (function () {
    function OpIntBitShiftLeft() {

    };
    OpIntBitShiftLeft.value = new OpIntBitShiftLeft();
    return OpIntBitShiftLeft;
})();
var OpIntBitShiftRight = /* #__PURE__ */ (function () {
    function OpIntBitShiftRight() {

    };
    OpIntBitShiftRight.value = new OpIntBitShiftRight();
    return OpIntBitShiftRight;
})();
var OpIntBitXor = /* #__PURE__ */ (function () {
    function OpIntBitXor() {

    };
    OpIntBitXor.value = new OpIntBitXor();
    return OpIntBitXor;
})();
var OpIntBitZeroFillShiftRight = /* #__PURE__ */ (function () {
    function OpIntBitZeroFillShiftRight() {

    };
    OpIntBitZeroFillShiftRight.value = new OpIntBitZeroFillShiftRight();
    return OpIntBitZeroFillShiftRight;
})();
var OpIntNum = /* #__PURE__ */ (function () {
    function OpIntNum(value0) {
        this.value0 = value0;
    };
    OpIntNum.create = function (value0) {
        return new OpIntNum(value0);
    };
    return OpIntNum;
})();
var OpIntOrd = /* #__PURE__ */ (function () {
    function OpIntOrd(value0) {
        this.value0 = value0;
    };
    OpIntOrd.create = function (value0) {
        return new OpIntOrd(value0);
    };
    return OpIntOrd;
})();
var OpNumberNum = /* #__PURE__ */ (function () {
    function OpNumberNum(value0) {
        this.value0 = value0;
    };
    OpNumberNum.create = function (value0) {
        return new OpNumberNum(value0);
    };
    return OpNumberNum;
})();
var OpNumberOrd = /* #__PURE__ */ (function () {
    function OpNumberOrd(value0) {
        this.value0 = value0;
    };
    OpNumberOrd.create = function (value0) {
        return new OpNumberOrd(value0);
    };
    return OpNumberOrd;
})();
var OpStringAppend = /* #__PURE__ */ (function () {
    function OpStringAppend() {

    };
    OpStringAppend.value = new OpStringAppend();
    return OpStringAppend;
})();
var OpStringOrd = /* #__PURE__ */ (function () {
    function OpStringOrd(value0) {
        this.value0 = value0;
    };
    OpStringOrd.create = function (value0) {
        return new OpStringOrd(value0);
    };
    return OpStringOrd;
})();
var OpBooleanNot = /* #__PURE__ */ (function () {
    function OpBooleanNot() {

    };
    OpBooleanNot.value = new OpBooleanNot();
    return OpBooleanNot;
})();
var OpIntBitNot = /* #__PURE__ */ (function () {
    function OpIntBitNot() {

    };
    OpIntBitNot.value = new OpIntBitNot();
    return OpIntBitNot;
})();
var OpIntNegate = /* #__PURE__ */ (function () {
    function OpIntNegate() {

    };
    OpIntNegate.value = new OpIntNegate();
    return OpIntNegate;
})();
var OpNumberNegate = /* #__PURE__ */ (function () {
    function OpNumberNegate() {

    };
    OpNumberNegate.value = new OpNumberNegate();
    return OpNumberNegate;
})();
var OpArrayLength = /* #__PURE__ */ (function () {
    function OpArrayLength() {

    };
    OpArrayLength.value = new OpArrayLength();
    return OpArrayLength;
})();
var OpIsTag = /* #__PURE__ */ (function () {
    function OpIsTag(value0) {
        this.value0 = value0;
    };
    OpIsTag.create = function (value0) {
        return new OpIsTag(value0);
    };
    return OpIsTag;
})();
var Op1 = /* #__PURE__ */ (function () {
    function Op1(value0, value1) {
        this.value0 = value0;
        this.value1 = value1;
    };
    Op1.create = function (value0) {
        return function (value1) {
            return new Op1(value0, value1);
        };
    };
    return Op1;
})();
var Op2 = /* #__PURE__ */ (function () {
    function Op2(value0, value1, value2) {
        this.value0 = value0;
        this.value1 = value1;
        this.value2 = value2;
    };
    Op2.create = function (value0) {
        return function (value1) {
            return function (value2) {
                return new Op2(value0, value1, value2);
            };
        };
    };
    return Op2;
})();
var EffectRefNew = /* #__PURE__ */ (function () {
    function EffectRefNew(value0) {
        this.value0 = value0;
    };
    EffectRefNew.create = function (value0) {
        return new EffectRefNew(value0);
    };
    return EffectRefNew;
})();
var EffectRefRead = /* #__PURE__ */ (function () {
    function EffectRefRead(value0) {
        this.value0 = value0;
    };
    EffectRefRead.create = function (value0) {
        return new EffectRefRead(value0);
    };
    return EffectRefRead;
})();
var EffectRefWrite = /* #__PURE__ */ (function () {
    function EffectRefWrite(value0, value1) {
        this.value0 = value0;
        this.value1 = value1;
    };
    EffectRefWrite.create = function (value0) {
        return function (value1) {
            return new EffectRefWrite(value0, value1);
        };
    };
    return EffectRefWrite;
})();
var GetProp = /* #__PURE__ */ (function () {
    function GetProp(value0) {
        this.value0 = value0;
    };
    GetProp.create = function (value0) {
        return new GetProp(value0);
    };
    return GetProp;
})();
var GetIndex = /* #__PURE__ */ (function () {
    function GetIndex(value0) {
        this.value0 = value0;
    };
    GetIndex.create = function (value0) {
        return new GetIndex(value0);
    };
    return GetIndex;
})();
var GetCtorField = /* #__PURE__ */ (function () {
    function GetCtorField(value0, value1, value2, value3, value4, value5) {
        this.value0 = value0;
        this.value1 = value1;
        this.value2 = value2;
        this.value3 = value3;
        this.value4 = value4;
        this.value5 = value5;
    };
    GetCtorField.create = function (value0) {
        return function (value1) {
            return function (value2) {
                return function (value3) {
                    return function (value4) {
                        return function (value5) {
                            return new GetCtorField(value0, value1, value2, value3, value4, value5);
                        };
                    };
                };
            };
        };
    };
    return GetCtorField;
})();
var Var = /* #__PURE__ */ (function () {
    function Var(value0) {
        this.value0 = value0;
    };
    Var.create = function (value0) {
        return new Var(value0);
    };
    return Var;
})();
var Local = /* #__PURE__ */ (function () {
    function Local(value0, value1) {
        this.value0 = value0;
        this.value1 = value1;
    };
    Local.create = function (value0) {
        return function (value1) {
            return new Local(value0, value1);
        };
    };
    return Local;
})();
var Lit = /* #__PURE__ */ (function () {
    function Lit(value0) {
        this.value0 = value0;
    };
    Lit.create = function (value0) {
        return new Lit(value0);
    };
    return Lit;
})();
var App = /* #__PURE__ */ (function () {
    function App(value0, value1) {
        this.value0 = value0;
        this.value1 = value1;
    };
    App.create = function (value0) {
        return function (value1) {
            return new App(value0, value1);
        };
    };
    return App;
})();
var Abs = /* #__PURE__ */ (function () {
    function Abs(value0, value1) {
        this.value0 = value0;
        this.value1 = value1;
    };
    Abs.create = function (value0) {
        return function (value1) {
            return new Abs(value0, value1);
        };
    };
    return Abs;
})();
var UncurriedApp = /* #__PURE__ */ (function () {
    function UncurriedApp(value0, value1) {
        this.value0 = value0;
        this.value1 = value1;
    };
    UncurriedApp.create = function (value0) {
        return function (value1) {
            return new UncurriedApp(value0, value1);
        };
    };
    return UncurriedApp;
})();
var UncurriedAbs = /* #__PURE__ */ (function () {
    function UncurriedAbs(value0, value1) {
        this.value0 = value0;
        this.value1 = value1;
    };
    UncurriedAbs.create = function (value0) {
        return function (value1) {
            return new UncurriedAbs(value0, value1);
        };
    };
    return UncurriedAbs;
})();
var UncurriedEffectApp = /* #__PURE__ */ (function () {
    function UncurriedEffectApp(value0, value1) {
        this.value0 = value0;
        this.value1 = value1;
    };
    UncurriedEffectApp.create = function (value0) {
        return function (value1) {
            return new UncurriedEffectApp(value0, value1);
        };
    };
    return UncurriedEffectApp;
})();
var UncurriedEffectAbs = /* #__PURE__ */ (function () {
    function UncurriedEffectAbs(value0, value1) {
        this.value0 = value0;
        this.value1 = value1;
    };
    UncurriedEffectAbs.create = function (value0) {
        return function (value1) {
            return new UncurriedEffectAbs(value0, value1);
        };
    };
    return UncurriedEffectAbs;
})();
var Accessor = /* #__PURE__ */ (function () {
    function Accessor(value0, value1) {
        this.value0 = value0;
        this.value1 = value1;
    };
    Accessor.create = function (value0) {
        return function (value1) {
            return new Accessor(value0, value1);
        };
    };
    return Accessor;
})();
var Update = /* #__PURE__ */ (function () {
    function Update(value0, value1) {
        this.value0 = value0;
        this.value1 = value1;
    };
    Update.create = function (value0) {
        return function (value1) {
            return new Update(value0, value1);
        };
    };
    return Update;
})();
var CtorSaturated = /* #__PURE__ */ (function () {
    function CtorSaturated(value0, value1, value2, value3, value4) {
        this.value0 = value0;
        this.value1 = value1;
        this.value2 = value2;
        this.value3 = value3;
        this.value4 = value4;
    };
    CtorSaturated.create = function (value0) {
        return function (value1) {
            return function (value2) {
                return function (value3) {
                    return function (value4) {
                        return new CtorSaturated(value0, value1, value2, value3, value4);
                    };
                };
            };
        };
    };
    return CtorSaturated;
})();
var CtorDef = /* #__PURE__ */ (function () {
    function CtorDef(value0, value1, value2, value3) {
        this.value0 = value0;
        this.value1 = value1;
        this.value2 = value2;
        this.value3 = value3;
    };
    CtorDef.create = function (value0) {
        return function (value1) {
            return function (value2) {
                return function (value3) {
                    return new CtorDef(value0, value1, value2, value3);
                };
            };
        };
    };
    return CtorDef;
})();
var LetRec = /* #__PURE__ */ (function () {
    function LetRec(value0, value1, value2) {
        this.value0 = value0;
        this.value1 = value1;
        this.value2 = value2;
    };
    LetRec.create = function (value0) {
        return function (value1) {
            return function (value2) {
                return new LetRec(value0, value1, value2);
            };
        };
    };
    return LetRec;
})();
var Let = /* #__PURE__ */ (function () {
    function Let(value0, value1, value2, value3) {
        this.value0 = value0;
        this.value1 = value1;
        this.value2 = value2;
        this.value3 = value3;
    };
    Let.create = function (value0) {
        return function (value1) {
            return function (value2) {
                return function (value3) {
                    return new Let(value0, value1, value2, value3);
                };
            };
        };
    };
    return Let;
})();
var EffectBind = /* #__PURE__ */ (function () {
    function EffectBind(value0, value1, value2, value3) {
        this.value0 = value0;
        this.value1 = value1;
        this.value2 = value2;
        this.value3 = value3;
    };
    EffectBind.create = function (value0) {
        return function (value1) {
            return function (value2) {
                return function (value3) {
                    return new EffectBind(value0, value1, value2, value3);
                };
            };
        };
    };
    return EffectBind;
})();
var EffectPure = /* #__PURE__ */ (function () {
    function EffectPure(value0) {
        this.value0 = value0;
    };
    EffectPure.create = function (value0) {
        return new EffectPure(value0);
    };
    return EffectPure;
})();
var EffectDefer = /* #__PURE__ */ (function () {
    function EffectDefer(value0) {
        this.value0 = value0;
    };
    EffectDefer.create = function (value0) {
        return new EffectDefer(value0);
    };
    return EffectDefer;
})();
var Branch = /* #__PURE__ */ (function () {
    function Branch(value0, value1) {
        this.value0 = value0;
        this.value1 = value1;
    };
    Branch.create = function (value0) {
        return function (value1) {
            return new Branch(value0, value1);
        };
    };
    return Branch;
})();
var PrimOp = /* #__PURE__ */ (function () {
    function PrimOp(value0) {
        this.value0 = value0;
    };
    PrimOp.create = function (value0) {
        return new PrimOp(value0);
    };
    return PrimOp;
})();
var PrimEffect = /* #__PURE__ */ (function () {
    function PrimEffect(value0) {
        this.value0 = value0;
    };
    PrimEffect.create = function (value0) {
        return new PrimEffect(value0);
    };
    return PrimEffect;
})();
var PrimUndefined = /* #__PURE__ */ (function () {
    function PrimUndefined() {

    };
    PrimUndefined.value = new PrimUndefined();
    return PrimUndefined;
})();
var Fail = /* #__PURE__ */ (function () {
    function Fail(value0) {
        this.value0 = value0;
    };
    Fail.create = function (value0) {
        return new Fail(value0);
    };
    return Fail;
})();
var Typed = /* #__PURE__ */ (function () {
    function Typed(value0, value1) {
        this.value0 = value0;
        this.value1 = value1;
    };
    Typed.create = function (value0) {
        return function (value1) {
            return new Typed(value0, value1);
        };
    };
    return Typed;
})();
var ordLevel = Data_Ord.ordInt;
var newtypeLevel_ = {
    Coercible0: function () {
        return undefined;
    }
};
var functorPair = {
    map: function (f) {
        return function (m) {
            return new Pair(f(m.value0), f(m.value1));
        };
    }
};
var functorBackendOperator = {
    map: function (f) {
        return function (m) {
            if (m instanceof Op1) {
                return new Op1(m.value0, f(m.value1));
            };
            if (m instanceof Op2) {
                return new Op2(m.value0, f(m.value1), f(m.value2));
            };
            throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Syntax (line 0, column 0 - line 0, column 0): " + [ m.constructor.name ]);
        };
    }
};
var functorBackendEffect = {
    map: function (f) {
        return function (m) {
            if (m instanceof EffectRefNew) {
                return new EffectRefNew(f(m.value0));
            };
            if (m instanceof EffectRefRead) {
                return new EffectRefRead(f(m.value0));
            };
            if (m instanceof EffectRefWrite) {
                return new EffectRefWrite(f(m.value0), f(m.value1));
            };
            throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Syntax (line 0, column 0 - line 0, column 0): " + [ m.constructor.name ]);
        };
    }
};
var functorBackendSyntax = {
    map: function (f) {
        return function (m) {
            if (m instanceof Var) {
                return new Var(m.value0);
            };
            if (m instanceof Local) {
                return new Local(m.value0, m.value1);
            };
            if (m instanceof Lit) {
                return new Lit(Data_Functor.map(PureScript_Backend_Optimizer_CoreFn.functorLiteral)(f)(m.value0));
            };
            if (m instanceof App) {
                return new App(f(m.value0), Data_Functor.map(Data_Array_NonEmpty_Internal.functorNonEmptyArray)(f)(m.value1));
            };
            if (m instanceof Abs) {
                return new Abs(m.value0, f(m.value1));
            };
            if (m instanceof UncurriedApp) {
                return new UncurriedApp(f(m.value0), Data_Functor.map(Data_Functor.functorArray)(f)(m.value1));
            };
            if (m instanceof UncurriedAbs) {
                return new UncurriedAbs(m.value0, f(m.value1));
            };
            if (m instanceof UncurriedEffectApp) {
                return new UncurriedEffectApp(f(m.value0), Data_Functor.map(Data_Functor.functorArray)(f)(m.value1));
            };
            if (m instanceof UncurriedEffectAbs) {
                return new UncurriedEffectAbs(m.value0, f(m.value1));
            };
            if (m instanceof Accessor) {
                return new Accessor(f(m.value0), m.value1);
            };
            if (m instanceof Update) {
                return new Update(f(m.value0), Data_Functor.map(Data_Functor.functorArray)(Data_Functor.map(PureScript_Backend_Optimizer_CoreFn.functorProp)(f))(m.value1));
            };
            if (m instanceof CtorSaturated) {
                return new CtorSaturated(m.value0, m.value1, m.value2, m.value3, Data_Functor.map(Data_Functor.functorArray)(Data_Functor.map(Data_Tuple.functorTuple)(f))(m.value4));
            };
            if (m instanceof CtorDef) {
                return new CtorDef(m.value0, m.value1, m.value2, m.value3);
            };
            if (m instanceof LetRec) {
                return new LetRec(m.value0, Data_Functor.map(Data_Array_NonEmpty_Internal.functorNonEmptyArray)(Data_Functor.map(Data_Tuple.functorTuple)(f))(m.value1), f(m.value2));
            };
            if (m instanceof Let) {
                return new Let(m.value0, m.value1, f(m.value2), f(m.value3));
            };
            if (m instanceof EffectBind) {
                return new EffectBind(m.value0, m.value1, f(m.value2), f(m.value3));
            };
            if (m instanceof EffectPure) {
                return new EffectPure(f(m.value0));
            };
            if (m instanceof EffectDefer) {
                return new EffectDefer(f(m.value0));
            };
            if (m instanceof Branch) {
                return new Branch(Data_Functor.map(Data_Array_NonEmpty_Internal.functorNonEmptyArray)(Data_Functor.map(functorPair)(f))(m.value0), f(m.value1));
            };
            if (m instanceof PrimOp) {
                return new PrimOp(Data_Functor.map(functorBackendOperator)(f)(m.value0));
            };
            if (m instanceof PrimEffect) {
                return new PrimEffect(Data_Functor.map(functorBackendEffect)(f)(m.value0));
            };
            if (m instanceof PrimUndefined) {
                return PrimUndefined.value;
            };
            if (m instanceof Fail) {
                return new Fail(m.value0);
            };
            if (m instanceof Typed) {
                return new Typed(m.value0, f(m.value1));
            };
            throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Syntax (line 0, column 0 - line 0, column 0): " + [ m.constructor.name ]);
        };
    }
};
var foldablePair = {
    foldl: function (f) {
        return function (acc) {
            return function (v) {
                return f(f(acc)(v.value0))(v.value1);
            };
        };
    },
    foldr: function (f) {
        return function (acc) {
            return function (v) {
                return f(v.value0)(f(v.value1)(acc));
            };
        };
    },
    foldMap: function (dictMonoid) {
        var Semigroup0 = dictMonoid.Semigroup0();
        return function (f) {
            return function (v) {
                return Data_Semigroup.append(Semigroup0)(f(v.value0))(f(v.value1));
            };
        };
    }
};
var traversablePair = {
    sequence: function (dictApplicative) {
        return function (a) {
            return Data_Traversable.sequenceDefault(traversablePair)(dictApplicative)(a);
        };
    },
    traverse: function (dictApplicative) {
        var Apply0 = dictApplicative.Apply0();
        var Functor0 = (dictApplicative.Apply0()).Functor0();
        return function (f) {
            return function (v) {
                return Control_Apply.apply(Apply0)(Data_Functor.map(Functor0)(Pair.create)(f(v.value0)))(f(v.value1));
            };
        };
    },
    Functor0: function () {
        return functorPair;
    },
    Foldable1: function () {
        return foldablePair;
    }
};
var foldableBackendOperator = {
    foldr: function (a) {
        return Data_Foldable.foldrDefault(foldableBackendOperator)(a);
    },
    foldl: function (a) {
        return Data_Foldable.foldlDefault(foldableBackendOperator)(a);
    },
    foldMap: function (dictMonoid) {
        var Semigroup0 = dictMonoid.Semigroup0();
        return function (f) {
            return function (v) {
                if (v instanceof Op1) {
                    return f(v.value1);
                };
                if (v instanceof Op2) {
                    return Data_Semigroup.append(Semigroup0)(f(v.value1))(f(v.value2));
                };
                throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Syntax (line 239, column 15 - line 241, column 28): " + [ v.constructor.name ]);
            };
        };
    }
};
var traversableBackendOperato = {
    sequence: function (dictApplicative) {
        return function (a) {
            return Data_Traversable.sequenceDefault(traversableBackendOperato)(dictApplicative)(a);
        };
    },
    traverse: function (dictApplicative) {
        var Functor0 = (dictApplicative.Apply0()).Functor0();
        var Apply0 = dictApplicative.Apply0();
        return function (f) {
            return function (v) {
                if (v instanceof Op1) {
                    return Data_Functor.map(Functor0)(Op1.create(v.value0))(f(v.value1));
                };
                if (v instanceof Op2) {
                    return Control_Apply.apply(Apply0)(Data_Functor.map(Functor0)(Op2.create(v.value0))(f(v.value1)))(f(v.value2));
                };
                throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Syntax (line 245, column 16 - line 247, column 39): " + [ v.constructor.name ]);
            };
        };
    },
    Functor0: function () {
        return functorBackendOperator;
    },
    Foldable1: function () {
        return foldableBackendOperator;
    }
};
var foldableBackendEffect = {
    foldr: function (a) {
        return Data_Foldable.foldrDefault(foldableBackendEffect)(a);
    },
    foldl: function (a) {
        return Data_Foldable.foldlDefault(foldableBackendEffect)(a);
    },
    foldMap: function (dictMonoid) {
        var Semigroup0 = dictMonoid.Semigroup0();
        return function (f) {
            return function (v) {
                if (v instanceof EffectRefNew) {
                    return f(v.value0);
                };
                if (v instanceof EffectRefRead) {
                    return f(v.value0);
                };
                if (v instanceof EffectRefWrite) {
                    return Data_Semigroup.append(Semigroup0)(f(v.value0))(f(v.value1));
                };
                throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Syntax (line 254, column 15 - line 257, column 37): " + [ v.constructor.name ]);
            };
        };
    }
};
var foldableBackendSyntax = {
    foldr: function (a) {
        return Data_Foldable.foldrDefault(foldableBackendSyntax)(a);
    },
    foldl: function (a) {
        return Data_Foldable.foldlDefault(foldableBackendSyntax)(a);
    },
    foldMap: function (dictMonoid) {
        var mempty = Data_Monoid.mempty(dictMonoid);
        var Semigroup0 = dictMonoid.Semigroup0();
        return function (f) {
            return function (v) {
                if (v instanceof Var) {
                    return mempty;
                };
                if (v instanceof Local) {
                    return mempty;
                };
                if (v instanceof Lit) {
                    if (v.value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitArray) {
                        return Data_Foldable.foldMap(Data_Foldable.foldableArray)(dictMonoid)(f)(v.value0.value0);
                    };
                    if (v.value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitRecord) {
                        return Data_Foldable.foldMap(Data_Foldable.foldableArray)(dictMonoid)(Data_Foldable.foldMap(PureScript_Backend_Optimizer_CoreFn.foldableProp)(dictMonoid)(f))(v.value0.value0);
                    };
                    return mempty;
                };
                if (v instanceof App) {
                    return Data_Semigroup.append(Semigroup0)(f(v.value0))(Data_Foldable.foldMap(Data_Array_NonEmpty_Internal.foldableNonEmptyArray)(dictMonoid)(f)(v.value1));
                };
                if (v instanceof Abs) {
                    return f(v.value1);
                };
                if (v instanceof UncurriedApp) {
                    return Data_Semigroup.append(Semigroup0)(f(v.value0))(Data_Foldable.foldMap(Data_Foldable.foldableArray)(dictMonoid)(f)(v.value1));
                };
                if (v instanceof UncurriedAbs) {
                    return f(v.value1);
                };
                if (v instanceof UncurriedEffectApp) {
                    return Data_Semigroup.append(Semigroup0)(f(v.value0))(Data_Foldable.foldMap(Data_Foldable.foldableArray)(dictMonoid)(f)(v.value1));
                };
                if (v instanceof UncurriedEffectAbs) {
                    return f(v.value1);
                };
                if (v instanceof Accessor) {
                    return f(v.value0);
                };
                if (v instanceof Update) {
                    return Data_Semigroup.append(Semigroup0)(f(v.value0))(Data_Foldable.foldMap(Data_Foldable.foldableArray)(dictMonoid)(Data_Foldable.foldMap(PureScript_Backend_Optimizer_CoreFn.foldableProp)(dictMonoid)(f))(v.value1));
                };
                if (v instanceof LetRec) {
                    return Data_Semigroup.append(Semigroup0)(Data_Foldable.foldMap(Data_Array_NonEmpty_Internal.foldableNonEmptyArray)(dictMonoid)(Data_Foldable.foldMap(Data_Foldable.foldableTuple)(dictMonoid)(f))(v.value1))(f(v.value2));
                };
                if (v instanceof Let) {
                    return Data_Semigroup.append(Semigroup0)(f(v.value2))(f(v.value3));
                };
                if (v instanceof EffectBind) {
                    return Data_Semigroup.append(Semigroup0)(f(v.value2))(f(v.value3));
                };
                if (v instanceof EffectPure) {
                    return f(v.value0);
                };
                if (v instanceof EffectDefer) {
                    return f(v.value0);
                };
                if (v instanceof Branch) {
                    return Data_Semigroup.append(Semigroup0)(Data_Foldable.foldMap(Data_Array_NonEmpty_Internal.foldableNonEmptyArray)(dictMonoid)(Data_Foldable.foldMap(foldablePair)(dictMonoid)(f))(v.value0))(f(v.value1));
                };
                if (v instanceof PrimOp) {
                    return Data_Foldable.foldMap(foldableBackendOperator)(dictMonoid)(f)(v.value0);
                };
                if (v instanceof PrimEffect) {
                    return Data_Foldable.foldMap(foldableBackendEffect)(dictMonoid)(f)(v.value0);
                };
                if (v instanceof PrimUndefined) {
                    return mempty;
                };
                if (v instanceof CtorSaturated) {
                    return Data_Foldable.foldMap(Data_Foldable.foldableArray)(dictMonoid)(Data_Foldable.foldMap(Data_Foldable.foldableTuple)(dictMonoid)(f))(v.value4);
                };
                if (v instanceof CtorDef) {
                    return mempty;
                };
                if (v instanceof Fail) {
                    return mempty;
                };
                if (v instanceof Typed) {
                    return f(v.value1);
                };
                throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Syntax (line 134, column 15 - line 162, column 21): " + [ v.constructor.name ]);
            };
        };
    }
};
var traversableBackendEffect = {
    sequence: function (dictApplicative) {
        return function (a) {
            return Data_Traversable.sequenceDefault(traversableBackendEffect)(dictApplicative)(a);
        };
    },
    traverse: function (dictApplicative) {
        var Functor0 = (dictApplicative.Apply0()).Functor0();
        var Apply0 = dictApplicative.Apply0();
        return function (f) {
            return function (v) {
                if (v instanceof EffectRefNew) {
                    return Data_Functor.map(Functor0)(EffectRefNew.create)(f(v.value0));
                };
                if (v instanceof EffectRefRead) {
                    return Data_Functor.map(Functor0)(EffectRefRead.create)(f(v.value0));
                };
                if (v instanceof EffectRefWrite) {
                    return Control_Apply.apply(Apply0)(Data_Functor.map(Functor0)(EffectRefWrite.create)(f(v.value0)))(f(v.value1));
                };
                throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Syntax (line 261, column 16 - line 264, column 57): " + [ v.constructor.name ]);
            };
        };
    },
    Functor0: function () {
        return functorBackendEffect;
    },
    Foldable1: function () {
        return foldableBackendEffect;
    }
};
var traversableBackendSyntax = {
    sequence: function (dictApplicative) {
        return function (a) {
            return Data_Traversable.sequenceDefault(traversableBackendSyntax)(dictApplicative)(a);
        };
    },
    traverse: function (dictApplicative) {
        var Functor0 = (dictApplicative.Apply0()).Functor0();
        var Apply0 = dictApplicative.Apply0();
        return function (f) {
            return function (v) {
                if (v instanceof Var) {
                    return Control_Applicative.pure(dictApplicative)(new Var(v.value0));
                };
                if (v instanceof Local) {
                    return Control_Applicative.pure(dictApplicative)(new Local(v.value0, v.value1));
                };
                if (v instanceof Lit) {
                    if (v.value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitInt) {
                        return Control_Applicative.pure(dictApplicative)(new Lit(new PureScript_Backend_Optimizer_CoreFn.LitInt(v.value0.value0)));
                    };
                    if (v.value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitNumber) {
                        return Control_Applicative.pure(dictApplicative)(new Lit(new PureScript_Backend_Optimizer_CoreFn.LitNumber(v.value0.value0)));
                    };
                    if (v.value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitString) {
                        return Control_Applicative.pure(dictApplicative)(new Lit(new PureScript_Backend_Optimizer_CoreFn.LitString(v.value0.value0)));
                    };
                    if (v.value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitChar) {
                        return Control_Applicative.pure(dictApplicative)(new Lit(new PureScript_Backend_Optimizer_CoreFn.LitChar(v.value0.value0)));
                    };
                    if (v.value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitBoolean) {
                        return Control_Applicative.pure(dictApplicative)(new Lit(new PureScript_Backend_Optimizer_CoreFn.LitBoolean(v.value0.value0)));
                    };
                    if (v.value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitArray) {
                        return Data_Functor.map(Functor0)(function ($838) {
                            return Lit.create(PureScript_Backend_Optimizer_CoreFn.LitArray.create($838));
                        })(Data_Traversable.traverse(Data_Traversable.traversableArray)(dictApplicative)(f)(v.value0.value0));
                    };
                    if (v.value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitRecord) {
                        return Data_Functor.map(Functor0)(function ($839) {
                            return Lit.create(PureScript_Backend_Optimizer_CoreFn.LitRecord.create($839));
                        })(Data_Traversable.traverse(Data_Traversable.traversableArray)(dictApplicative)(Data_Traversable.traverse(PureScript_Backend_Optimizer_CoreFn.traversableProp)(dictApplicative)(f))(v.value0.value0));
                    };
                    throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Syntax (line 172, column 7 - line 179, column 71): " + [ v.value0.constructor.name ]);
                };
                if (v instanceof App) {
                    return Control_Apply.apply(Apply0)(Data_Functor.map(Functor0)(App.create)(f(v.value0)))(Data_Traversable.traverse(Data_Array_NonEmpty_Internal.traversableNonEmptyArray)(dictApplicative)(f)(v.value1));
                };
                if (v instanceof Abs) {
                    return Data_Functor.map(Functor0)(Abs.create(v.value0))(f(v.value1));
                };
                if (v instanceof UncurriedApp) {
                    return Control_Apply.apply(Apply0)(Data_Functor.map(Functor0)(UncurriedApp.create)(f(v.value0)))(Data_Traversable.traverse(Data_Traversable.traversableArray)(dictApplicative)(f)(v.value1));
                };
                if (v instanceof UncurriedAbs) {
                    return Data_Functor.map(Functor0)(UncurriedAbs.create(v.value0))(f(v.value1));
                };
                if (v instanceof UncurriedEffectApp) {
                    return Control_Apply.apply(Apply0)(Data_Functor.map(Functor0)(UncurriedEffectApp.create)(f(v.value0)))(Data_Traversable.traverse(Data_Traversable.traversableArray)(dictApplicative)(f)(v.value1));
                };
                if (v instanceof UncurriedEffectAbs) {
                    return Data_Functor.map(Functor0)(UncurriedEffectAbs.create(v.value0))(f(v.value1));
                };
                if (v instanceof Accessor) {
                    return Data_Functor.map(Functor0)(Data_Function.flip(Accessor.create)(v.value1))(f(v.value0));
                };
                if (v instanceof Update) {
                    return Control_Apply.apply(Apply0)(Data_Functor.map(Functor0)(Update.create)(f(v.value0)))(Data_Traversable.traverse(Data_Traversable.traversableArray)(dictApplicative)(Data_Traversable.traverse(PureScript_Backend_Optimizer_CoreFn.traversableProp)(dictApplicative)(f))(v.value1));
                };
                if (v instanceof CtorDef) {
                    return Control_Applicative.pure(dictApplicative)(new CtorDef(v.value0, v.value1, v.value2, v.value3));
                };
                if (v instanceof CtorSaturated) {
                    return Data_Functor.map(Functor0)(CtorSaturated.create(v.value0)(v.value1)(v.value2)(v.value3))(Data_Traversable.traverse(Data_Traversable.traversableArray)(dictApplicative)(Data_Traversable.traverse(Data_Traversable.traversableTuple)(dictApplicative)(f))(v.value4));
                };
                if (v instanceof LetRec) {
                    return Control_Apply.apply(Apply0)(Data_Functor.map(Functor0)(LetRec.create(v.value0))(Data_Traversable.traverse(Data_Array_NonEmpty_Internal.traversableNonEmptyArray)(dictApplicative)(Data_Traversable.traverse(Data_Traversable.traversableTuple)(dictApplicative)(f))(v.value1)))(f(v.value2));
                };
                if (v instanceof Let) {
                    return Control_Apply.apply(Apply0)(Data_Functor.map(Functor0)(Let.create(v.value0)(v.value1))(f(v.value2)))(f(v.value3));
                };
                if (v instanceof EffectBind) {
                    return Control_Apply.apply(Apply0)(Data_Functor.map(Functor0)(EffectBind.create(v.value0)(v.value1))(f(v.value2)))(f(v.value3));
                };
                if (v instanceof EffectPure) {
                    return Data_Functor.map(Functor0)(EffectPure.create)(f(v.value0));
                };
                if (v instanceof EffectDefer) {
                    return Data_Functor.map(Functor0)(EffectDefer.create)(f(v.value0));
                };
                if (v instanceof Branch) {
                    return Control_Apply.apply(Apply0)(Data_Functor.map(Functor0)(Branch.create)(Data_Traversable.traverse(Data_Array_NonEmpty_Internal.traversableNonEmptyArray)(dictApplicative)(Data_Traversable.traverse(traversablePair)(dictApplicative)(f))(v.value0)))(f(v.value1));
                };
                if (v instanceof PrimOp) {
                    return Data_Functor.map(Functor0)(PrimOp.create)(Data_Traversable.traverse(traversableBackendOperato)(dictApplicative)(f)(v.value0));
                };
                if (v instanceof PrimEffect) {
                    return Data_Functor.map(Functor0)(PrimEffect.create)(Data_Traversable.traverse(traversableBackendEffect)(dictApplicative)(f)(v.value0));
                };
                if (v instanceof PrimUndefined) {
                    return Control_Applicative.pure(dictApplicative)(PrimUndefined.value);
                };
                if (v instanceof Fail) {
                    return Control_Applicative.pure(dictApplicative)(new Fail(v.value0));
                };
                if (v instanceof Typed) {
                    return Data_Functor.map(Functor0)(Typed.create(v.value0))(f(v.value1));
                };
                throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Syntax (line 166, column 16 - line 221, column 22): " + [ v.constructor.name ]);
            };
        };
    },
    Functor0: function () {
        return functorBackendSyntax;
    },
    Foldable1: function () {
        return foldableBackendSyntax;
    }
};
var eqPair = function (dictEq) {
    return {
        eq: function (x) {
            return function (y) {
                return Data_Eq.eq(dictEq)(x.value0)(y.value0) && Data_Eq.eq(dictEq)(x.value1)(y.value1);
            };
        }
    };
};
var eqLevel = Data_Eq.eqInt;
var eqTuple2 = /* #__PURE__ */ Data_Tuple.eqTuple(/* #__PURE__ */ Data_Maybe.eqMaybe(PureScript_Backend_Optimizer_CoreFn.eqIdent))(eqLevel);
var eqNonEmptyArray = /* #__PURE__ */ Data_Array_NonEmpty_Internal.eqNonEmptyArray(eqTuple2);
var eqArray1 = /* #__PURE__ */ Data_Eq.eqArray(eqTuple2);
var eqBackendOperatorOrd = {
    eq: function (x) {
        return function (y) {
            if (x instanceof OpEq && y instanceof OpEq) {
                return true;
            };
            if (x instanceof OpNotEq && y instanceof OpNotEq) {
                return true;
            };
            if (x instanceof OpGt && y instanceof OpGt) {
                return true;
            };
            if (x instanceof OpGte && y instanceof OpGte) {
                return true;
            };
            if (x instanceof OpLt && y instanceof OpLt) {
                return true;
            };
            if (x instanceof OpLte && y instanceof OpLte) {
                return true;
            };
            return false;
        };
    }
};
var ordBackendOperatorOrd = {
    compare: function (x) {
        return function (y) {
            if (x instanceof OpEq && y instanceof OpEq) {
                return Data_Ordering.EQ.value;
            };
            if (x instanceof OpEq) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof OpEq) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof OpNotEq && y instanceof OpNotEq) {
                return Data_Ordering.EQ.value;
            };
            if (x instanceof OpNotEq) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof OpNotEq) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof OpGt && y instanceof OpGt) {
                return Data_Ordering.EQ.value;
            };
            if (x instanceof OpGt) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof OpGt) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof OpGte && y instanceof OpGte) {
                return Data_Ordering.EQ.value;
            };
            if (x instanceof OpGte) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof OpGte) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof OpLt && y instanceof OpLt) {
                return Data_Ordering.EQ.value;
            };
            if (x instanceof OpLt) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof OpLt) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof OpLte && y instanceof OpLte) {
                return Data_Ordering.EQ.value;
            };
            throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Syntax (line 0, column 0 - line 0, column 0): " + [ x.constructor.name, y.constructor.name ]);
        };
    },
    Eq0: function () {
        return eqBackendOperatorOrd;
    }
};
var eqBackendOperatorNum = {
    eq: function (x) {
        return function (y) {
            if (x instanceof OpAdd && y instanceof OpAdd) {
                return true;
            };
            if (x instanceof OpDivide && y instanceof OpDivide) {
                return true;
            };
            if (x instanceof OpMultiply && y instanceof OpMultiply) {
                return true;
            };
            if (x instanceof OpSubtract && y instanceof OpSubtract) {
                return true;
            };
            return false;
        };
    }
};
var ordBackendOperatorNum = {
    compare: function (x) {
        return function (y) {
            if (x instanceof OpAdd && y instanceof OpAdd) {
                return Data_Ordering.EQ.value;
            };
            if (x instanceof OpAdd) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof OpAdd) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof OpDivide && y instanceof OpDivide) {
                return Data_Ordering.EQ.value;
            };
            if (x instanceof OpDivide) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof OpDivide) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof OpMultiply && y instanceof OpMultiply) {
                return Data_Ordering.EQ.value;
            };
            if (x instanceof OpMultiply) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof OpMultiply) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof OpSubtract && y instanceof OpSubtract) {
                return Data_Ordering.EQ.value;
            };
            throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Syntax (line 0, column 0 - line 0, column 0): " + [ x.constructor.name, y.constructor.name ]);
        };
    },
    Eq0: function () {
        return eqBackendOperatorNum;
    }
};
var eqBackendOperator2 = {
    eq: function (x) {
        return function (y) {
            if (x instanceof OpArrayIndex && y instanceof OpArrayIndex) {
                return true;
            };
            if (x instanceof OpBooleanAnd && y instanceof OpBooleanAnd) {
                return true;
            };
            if (x instanceof OpBooleanOr && y instanceof OpBooleanOr) {
                return true;
            };
            if (x instanceof OpBooleanOrd && y instanceof OpBooleanOrd) {
                return Data_Eq.eq(eqBackendOperatorOrd)(x.value0)(y.value0);
            };
            if (x instanceof OpCharOrd && y instanceof OpCharOrd) {
                return Data_Eq.eq(eqBackendOperatorOrd)(x.value0)(y.value0);
            };
            if (x instanceof OpIntBitAnd && y instanceof OpIntBitAnd) {
                return true;
            };
            if (x instanceof OpIntBitOr && y instanceof OpIntBitOr) {
                return true;
            };
            if (x instanceof OpIntBitShiftLeft && y instanceof OpIntBitShiftLeft) {
                return true;
            };
            if (x instanceof OpIntBitShiftRight && y instanceof OpIntBitShiftRight) {
                return true;
            };
            if (x instanceof OpIntBitXor && y instanceof OpIntBitXor) {
                return true;
            };
            if (x instanceof OpIntBitZeroFillShiftRight && y instanceof OpIntBitZeroFillShiftRight) {
                return true;
            };
            if (x instanceof OpIntNum && y instanceof OpIntNum) {
                return Data_Eq.eq(eqBackendOperatorNum)(x.value0)(y.value0);
            };
            if (x instanceof OpIntOrd && y instanceof OpIntOrd) {
                return Data_Eq.eq(eqBackendOperatorOrd)(x.value0)(y.value0);
            };
            if (x instanceof OpNumberNum && y instanceof OpNumberNum) {
                return Data_Eq.eq(eqBackendOperatorNum)(x.value0)(y.value0);
            };
            if (x instanceof OpNumberOrd && y instanceof OpNumberOrd) {
                return Data_Eq.eq(eqBackendOperatorOrd)(x.value0)(y.value0);
            };
            if (x instanceof OpStringAppend && y instanceof OpStringAppend) {
                return true;
            };
            if (x instanceof OpStringOrd && y instanceof OpStringOrd) {
                return Data_Eq.eq(eqBackendOperatorOrd)(x.value0)(y.value0);
            };
            return false;
        };
    }
};
var ordBackendOperator2 = {
    compare: function (x) {
        return function (y) {
            if (x instanceof OpArrayIndex && y instanceof OpArrayIndex) {
                return Data_Ordering.EQ.value;
            };
            if (x instanceof OpArrayIndex) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof OpArrayIndex) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof OpBooleanAnd && y instanceof OpBooleanAnd) {
                return Data_Ordering.EQ.value;
            };
            if (x instanceof OpBooleanAnd) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof OpBooleanAnd) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof OpBooleanOr && y instanceof OpBooleanOr) {
                return Data_Ordering.EQ.value;
            };
            if (x instanceof OpBooleanOr) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof OpBooleanOr) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof OpBooleanOrd && y instanceof OpBooleanOrd) {
                return Data_Ord.compare(ordBackendOperatorOrd)(x.value0)(y.value0);
            };
            if (x instanceof OpBooleanOrd) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof OpBooleanOrd) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof OpCharOrd && y instanceof OpCharOrd) {
                return Data_Ord.compare(ordBackendOperatorOrd)(x.value0)(y.value0);
            };
            if (x instanceof OpCharOrd) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof OpCharOrd) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof OpIntBitAnd && y instanceof OpIntBitAnd) {
                return Data_Ordering.EQ.value;
            };
            if (x instanceof OpIntBitAnd) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof OpIntBitAnd) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof OpIntBitOr && y instanceof OpIntBitOr) {
                return Data_Ordering.EQ.value;
            };
            if (x instanceof OpIntBitOr) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof OpIntBitOr) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof OpIntBitShiftLeft && y instanceof OpIntBitShiftLeft) {
                return Data_Ordering.EQ.value;
            };
            if (x instanceof OpIntBitShiftLeft) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof OpIntBitShiftLeft) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof OpIntBitShiftRight && y instanceof OpIntBitShiftRight) {
                return Data_Ordering.EQ.value;
            };
            if (x instanceof OpIntBitShiftRight) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof OpIntBitShiftRight) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof OpIntBitXor && y instanceof OpIntBitXor) {
                return Data_Ordering.EQ.value;
            };
            if (x instanceof OpIntBitXor) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof OpIntBitXor) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof OpIntBitZeroFillShiftRight && y instanceof OpIntBitZeroFillShiftRight) {
                return Data_Ordering.EQ.value;
            };
            if (x instanceof OpIntBitZeroFillShiftRight) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof OpIntBitZeroFillShiftRight) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof OpIntNum && y instanceof OpIntNum) {
                return Data_Ord.compare(ordBackendOperatorNum)(x.value0)(y.value0);
            };
            if (x instanceof OpIntNum) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof OpIntNum) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof OpIntOrd && y instanceof OpIntOrd) {
                return Data_Ord.compare(ordBackendOperatorOrd)(x.value0)(y.value0);
            };
            if (x instanceof OpIntOrd) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof OpIntOrd) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof OpNumberNum && y instanceof OpNumberNum) {
                return Data_Ord.compare(ordBackendOperatorNum)(x.value0)(y.value0);
            };
            if (x instanceof OpNumberNum) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof OpNumberNum) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof OpNumberOrd && y instanceof OpNumberOrd) {
                return Data_Ord.compare(ordBackendOperatorOrd)(x.value0)(y.value0);
            };
            if (x instanceof OpNumberOrd) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof OpNumberOrd) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof OpStringAppend && y instanceof OpStringAppend) {
                return Data_Ordering.EQ.value;
            };
            if (x instanceof OpStringAppend) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof OpStringAppend) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof OpStringOrd && y instanceof OpStringOrd) {
                return Data_Ord.compare(ordBackendOperatorOrd)(x.value0)(y.value0);
            };
            throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Syntax (line 0, column 0 - line 0, column 0): " + [ x.constructor.name, y.constructor.name ]);
        };
    },
    Eq0: function () {
        return eqBackendOperator2;
    }
};
var eqBackendOperator1 = {
    eq: function (x) {
        return function (y) {
            if (x instanceof OpBooleanNot && y instanceof OpBooleanNot) {
                return true;
            };
            if (x instanceof OpIntBitNot && y instanceof OpIntBitNot) {
                return true;
            };
            if (x instanceof OpIntNegate && y instanceof OpIntNegate) {
                return true;
            };
            if (x instanceof OpNumberNegate && y instanceof OpNumberNegate) {
                return true;
            };
            if (x instanceof OpArrayLength && y instanceof OpArrayLength) {
                return true;
            };
            if (x instanceof OpIsTag && y instanceof OpIsTag) {
                return Data_Eq.eq(eqQualified)(x.value0)(y.value0);
            };
            return false;
        };
    }
};
var ordBackendOperator1 = {
    compare: function (x) {
        return function (y) {
            if (x instanceof OpBooleanNot && y instanceof OpBooleanNot) {
                return Data_Ordering.EQ.value;
            };
            if (x instanceof OpBooleanNot) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof OpBooleanNot) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof OpIntBitNot && y instanceof OpIntBitNot) {
                return Data_Ordering.EQ.value;
            };
            if (x instanceof OpIntBitNot) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof OpIntBitNot) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof OpIntNegate && y instanceof OpIntNegate) {
                return Data_Ordering.EQ.value;
            };
            if (x instanceof OpIntNegate) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof OpIntNegate) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof OpNumberNegate && y instanceof OpNumberNegate) {
                return Data_Ordering.EQ.value;
            };
            if (x instanceof OpNumberNegate) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof OpNumberNegate) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof OpArrayLength && y instanceof OpArrayLength) {
                return Data_Ordering.EQ.value;
            };
            if (x instanceof OpArrayLength) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof OpArrayLength) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof OpIsTag && y instanceof OpIsTag) {
                return Data_Ord.compare(ordQualified)(x.value0)(y.value0);
            };
            throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Syntax (line 0, column 0 - line 0, column 0): " + [ x.constructor.name, y.constructor.name ]);
        };
    },
    Eq0: function () {
        return eqBackendOperator1;
    }
};
var eqBackendOperator = function (dictEq) {
    return {
        eq: function (x) {
            return function (y) {
                if (x instanceof Op1 && y instanceof Op1) {
                    return Data_Eq.eq(eqBackendOperator1)(x.value0)(y.value0) && Data_Eq.eq(dictEq)(x.value1)(y.value1);
                };
                if (x instanceof Op2 && y instanceof Op2) {
                    return Data_Eq.eq(eqBackendOperator2)(x.value0)(y.value0) && Data_Eq.eq(dictEq)(x.value1)(y.value1) && Data_Eq.eq(dictEq)(x.value2)(y.value2);
                };
                return false;
            };
        }
    };
};
var eqBackendEffect = function (dictEq) {
    return {
        eq: function (x) {
            return function (y) {
                if (x instanceof EffectRefNew && y instanceof EffectRefNew) {
                    return Data_Eq.eq(dictEq)(x.value0)(y.value0);
                };
                if (x instanceof EffectRefRead && y instanceof EffectRefRead) {
                    return Data_Eq.eq(dictEq)(x.value0)(y.value0);
                };
                if (x instanceof EffectRefWrite && y instanceof EffectRefWrite) {
                    return Data_Eq.eq(dictEq)(x.value0)(y.value0) && Data_Eq.eq(dictEq)(x.value1)(y.value1);
                };
                return false;
            };
        }
    };
};
var eqBackendAccessor = {
    eq: function (x) {
        return function (y) {
            if (x instanceof GetProp && y instanceof GetProp) {
                return x.value0 === y.value0;
            };
            if (x instanceof GetIndex && y instanceof GetIndex) {
                return x.value0 === y.value0;
            };
            if (x instanceof GetCtorField && y instanceof GetCtorField) {
                return Data_Eq.eq(eqQualified)(x.value0)(y.value0) && Data_Eq.eq(PureScript_Backend_Optimizer_CoreFn.eqConstructorType)(x.value1)(y.value1) && Data_Eq.eq(PureScript_Backend_Optimizer_CoreFn.eqProperName)(x.value2)(y.value2) && Data_Eq.eq(PureScript_Backend_Optimizer_CoreFn.eqIdent)(x.value3)(y.value3) && x.value4 === y.value4 && x.value5 === y.value5;
            };
            return false;
        };
    }
};
var eqBackendSyntax = function (dictEq) {
    var eqLiteral = PureScript_Backend_Optimizer_CoreFn.eqLiteral(dictEq);
    var eqNonEmptyArray1 = Data_Array_NonEmpty_Internal.eqNonEmptyArray(dictEq);
    var eqArray2 = Data_Eq.eqArray(dictEq);
    var eqArray3 = Data_Eq.eqArray(PureScript_Backend_Optimizer_CoreFn.eqProp(dictEq));
    var eqArray4 = Data_Eq.eqArray(eqTuple(dictEq));
    var eqNonEmptyArray2 = Data_Array_NonEmpty_Internal.eqNonEmptyArray(eqTuple1(dictEq));
    var eqNonEmptyArray3 = Data_Array_NonEmpty_Internal.eqNonEmptyArray(eqPair(dictEq));
    var eqBackendOperator3 = eqBackendOperator(dictEq);
    var eqBackendEffect1 = eqBackendEffect(dictEq);
    return {
        eq: function (x) {
            return function (y) {
                if (x instanceof Var && y instanceof Var) {
                    return Data_Eq.eq(eqQualified)(x.value0)(y.value0);
                };
                if (x instanceof Local && y instanceof Local) {
                    return Data_Eq.eq(eqMaybe)(x.value0)(y.value0) && Data_Eq.eq(eqLevel)(x.value1)(y.value1);
                };
                if (x instanceof Lit && y instanceof Lit) {
                    return Data_Eq.eq(eqLiteral)(x.value0)(y.value0);
                };
                if (x instanceof App && y instanceof App) {
                    return Data_Eq.eq(dictEq)(x.value0)(y.value0) && Data_Eq.eq(eqNonEmptyArray1)(x.value1)(y.value1);
                };
                if (x instanceof Abs && y instanceof Abs) {
                    return Data_Eq.eq(eqNonEmptyArray)(x.value0)(y.value0) && Data_Eq.eq(dictEq)(x.value1)(y.value1);
                };
                if (x instanceof UncurriedApp && y instanceof UncurriedApp) {
                    return Data_Eq.eq(dictEq)(x.value0)(y.value0) && Data_Eq.eq(eqArray2)(x.value1)(y.value1);
                };
                if (x instanceof UncurriedAbs && y instanceof UncurriedAbs) {
                    return Data_Eq.eq(eqArray1)(x.value0)(y.value0) && Data_Eq.eq(dictEq)(x.value1)(y.value1);
                };
                if (x instanceof UncurriedEffectApp && y instanceof UncurriedEffectApp) {
                    return Data_Eq.eq(dictEq)(x.value0)(y.value0) && Data_Eq.eq(eqArray2)(x.value1)(y.value1);
                };
                if (x instanceof UncurriedEffectAbs && y instanceof UncurriedEffectAbs) {
                    return Data_Eq.eq(eqArray1)(x.value0)(y.value0) && Data_Eq.eq(dictEq)(x.value1)(y.value1);
                };
                if (x instanceof Accessor && y instanceof Accessor) {
                    return Data_Eq.eq(dictEq)(x.value0)(y.value0) && Data_Eq.eq(eqBackendAccessor)(x.value1)(y.value1);
                };
                if (x instanceof Update && y instanceof Update) {
                    return Data_Eq.eq(dictEq)(x.value0)(y.value0) && Data_Eq.eq(eqArray3)(x.value1)(y.value1);
                };
                if (x instanceof CtorSaturated && y instanceof CtorSaturated) {
                    return Data_Eq.eq(eqQualified)(x.value0)(y.value0) && Data_Eq.eq(PureScript_Backend_Optimizer_CoreFn.eqConstructorType)(x.value1)(y.value1) && Data_Eq.eq(PureScript_Backend_Optimizer_CoreFn.eqProperName)(x.value2)(y.value2) && Data_Eq.eq(PureScript_Backend_Optimizer_CoreFn.eqIdent)(x.value3)(y.value3) && Data_Eq.eq(eqArray4)(x.value4)(y.value4);
                };
                if (x instanceof CtorDef && y instanceof CtorDef) {
                    return Data_Eq.eq(PureScript_Backend_Optimizer_CoreFn.eqConstructorType)(x.value0)(y.value0) && Data_Eq.eq(PureScript_Backend_Optimizer_CoreFn.eqProperName)(x.value1)(y.value1) && Data_Eq.eq(PureScript_Backend_Optimizer_CoreFn.eqIdent)(x.value2)(y.value2) && Data_Eq.eq(eqArray)(x.value3)(y.value3);
                };
                if (x instanceof LetRec && y instanceof LetRec) {
                    return Data_Eq.eq(eqLevel)(x.value0)(y.value0) && Data_Eq.eq(eqNonEmptyArray2)(x.value1)(y.value1) && Data_Eq.eq(dictEq)(x.value2)(y.value2);
                };
                if (x instanceof Let && y instanceof Let) {
                    return Data_Eq.eq(eqMaybe)(x.value0)(y.value0) && Data_Eq.eq(eqLevel)(x.value1)(y.value1) && Data_Eq.eq(dictEq)(x.value2)(y.value2) && Data_Eq.eq(dictEq)(x.value3)(y.value3);
                };
                if (x instanceof EffectBind && y instanceof EffectBind) {
                    return Data_Eq.eq(eqMaybe)(x.value0)(y.value0) && Data_Eq.eq(eqLevel)(x.value1)(y.value1) && Data_Eq.eq(dictEq)(x.value2)(y.value2) && Data_Eq.eq(dictEq)(x.value3)(y.value3);
                };
                if (x instanceof EffectPure && y instanceof EffectPure) {
                    return Data_Eq.eq(dictEq)(x.value0)(y.value0);
                };
                if (x instanceof EffectDefer && y instanceof EffectDefer) {
                    return Data_Eq.eq(dictEq)(x.value0)(y.value0);
                };
                if (x instanceof Branch && y instanceof Branch) {
                    return Data_Eq.eq(eqNonEmptyArray3)(x.value0)(y.value0) && Data_Eq.eq(dictEq)(x.value1)(y.value1);
                };
                if (x instanceof PrimOp && y instanceof PrimOp) {
                    return Data_Eq.eq(eqBackendOperator3)(x.value0)(y.value0);
                };
                if (x instanceof PrimEffect && y instanceof PrimEffect) {
                    return Data_Eq.eq(eqBackendEffect1)(x.value0)(y.value0);
                };
                if (x instanceof PrimUndefined && y instanceof PrimUndefined) {
                    return true;
                };
                if (x instanceof Fail && y instanceof Fail) {
                    return x.value0 === y.value0;
                };
                if (x instanceof Typed && y instanceof Typed) {
                    return Data_Eq.eq(PureScript_Backend_Optimizer_CoreFn.eqExprType)(x.value0)(y.value0) && Data_Eq.eq(dictEq)(x.value1)(y.value1);
                };
                return false;
            };
        }
    };
};
var ordBackendAccessor = {
    compare: function (x) {
        return function (y) {
            if (x instanceof GetProp && y instanceof GetProp) {
                return Data_Ord.compare(Data_Ord.ordString)(x.value0)(y.value0);
            };
            if (x instanceof GetProp) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof GetProp) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof GetIndex && y instanceof GetIndex) {
                return Data_Ord.compare(Data_Ord.ordInt)(x.value0)(y.value0);
            };
            if (x instanceof GetIndex) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof GetIndex) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof GetCtorField && y instanceof GetCtorField) {
                var v = Data_Ord.compare(ordQualified)(x.value0)(y.value0);
                if (v instanceof Data_Ordering.LT) {
                    return Data_Ordering.LT.value;
                };
                if (v instanceof Data_Ordering.GT) {
                    return Data_Ordering.GT.value;
                };
                var v1 = Data_Ord.compare(PureScript_Backend_Optimizer_CoreFn.ordConstructorType)(x.value1)(y.value1);
                if (v1 instanceof Data_Ordering.LT) {
                    return Data_Ordering.LT.value;
                };
                if (v1 instanceof Data_Ordering.GT) {
                    return Data_Ordering.GT.value;
                };
                var v2 = Data_Ord.compare(PureScript_Backend_Optimizer_CoreFn.ordProperName)(x.value2)(y.value2);
                if (v2 instanceof Data_Ordering.LT) {
                    return Data_Ordering.LT.value;
                };
                if (v2 instanceof Data_Ordering.GT) {
                    return Data_Ordering.GT.value;
                };
                var v3 = Data_Ord.compare(PureScript_Backend_Optimizer_CoreFn.ordIdent)(x.value3)(y.value3);
                if (v3 instanceof Data_Ordering.LT) {
                    return Data_Ordering.LT.value;
                };
                if (v3 instanceof Data_Ordering.GT) {
                    return Data_Ordering.GT.value;
                };
                var v4 = Data_Ord.compare(Data_Ord.ordString)(x.value4)(y.value4);
                if (v4 instanceof Data_Ordering.LT) {
                    return Data_Ordering.LT.value;
                };
                if (v4 instanceof Data_Ordering.GT) {
                    return Data_Ordering.GT.value;
                };
                return Data_Ord.compare(Data_Ord.ordInt)(x.value5)(y.value5);
            };
            throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Syntax (line 0, column 0 - line 0, column 0): " + [ x.constructor.name, y.constructor.name ]);
        };
    },
    Eq0: function () {
        return eqBackendAccessor;
    }
};
var syntaxOf = function (dict) {
    return dict.syntaxOf;
};
var sndPair = function (v) {
    return v.value1;
};
var fstPair = function (v) {
    return v.value0;
};
export {
    syntaxOf,
    Var,
    Local,
    Lit,
    App,
    Abs,
    UncurriedApp,
    UncurriedAbs,
    UncurriedEffectApp,
    UncurriedEffectAbs,
    Accessor,
    Update,
    CtorSaturated,
    CtorDef,
    LetRec,
    Let,
    EffectBind,
    EffectPure,
    EffectDefer,
    Branch,
    PrimOp,
    PrimEffect,
    PrimUndefined,
    Fail,
    Typed,
    Level,
    Pair,
    fstPair,
    sndPair,
    GetProp,
    GetIndex,
    GetCtorField,
    Op1,
    Op2,
    OpBooleanNot,
    OpIntBitNot,
    OpIntNegate,
    OpNumberNegate,
    OpArrayLength,
    OpIsTag,
    OpArrayIndex,
    OpBooleanAnd,
    OpBooleanOr,
    OpBooleanOrd,
    OpCharOrd,
    OpIntBitAnd,
    OpIntBitOr,
    OpIntBitShiftLeft,
    OpIntBitShiftRight,
    OpIntBitXor,
    OpIntBitZeroFillShiftRight,
    OpIntNum,
    OpIntOrd,
    OpNumberNum,
    OpNumberOrd,
    OpStringAppend,
    OpStringOrd,
    OpAdd,
    OpDivide,
    OpMultiply,
    OpSubtract,
    OpEq,
    OpNotEq,
    OpGt,
    OpGte,
    OpLt,
    OpLte,
    EffectRefNew,
    EffectRefRead,
    EffectRefWrite,
    eqBackendSyntax,
    eqLevel,
    ordLevel,
    newtypeLevel_,
    eqPair,
    eqBackendAccessor,
    ordBackendAccessor,
    eqBackendOperator,
    eqBackendOperator1,
    ordBackendOperator1,
    eqBackendOperator2,
    ordBackendOperator2,
    eqBackendOperatorNum,
    ordBackendOperatorNum,
    eqBackendOperatorOrd,
    ordBackendOperatorOrd,
    eqBackendEffect,
    functorBackendSyntax,
    foldableBackendSyntax,
    traversableBackendSyntax,
    functorPair,
    foldablePair,
    traversablePair,
    functorBackendOperator,
    foldableBackendOperator,
    traversableBackendOperato,
    functorBackendEffect,
    foldableBackendEffect,
    traversableBackendEffect
};
