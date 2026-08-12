import * as Control_Applicative from "../Control.Applicative/index.js";
import * as Data_Array from "../Data.Array/index.js";
import * as Data_Array_NonEmpty from "../Data.Array.NonEmpty/index.js";
import * as Data_Array_NonEmpty_Internal from "../Data.Array.NonEmpty.Internal/index.js";
import * as Data_Boolean from "../Data.Boolean/index.js";
import * as Data_Foldable from "../Data.Foldable/index.js";
import * as Data_Functor from "../Data.Functor/index.js";
import * as Data_Maybe from "../Data.Maybe/index.js";
import * as Data_Monoid from "../Data.Monoid/index.js";
import * as Data_Ord from "../Data.Ord/index.js";
import * as Data_Ordering from "../Data.Ordering/index.js";
import * as Data_Semigroup from "../Data.Semigroup/index.js";
import * as Data_Show from "../Data.Show/index.js";
import * as Data_Tuple from "../Data.Tuple/index.js";
import * as Dodo from "../Dodo/index.js";
import * as Dodo_Common from "../Dodo.Common/index.js";
import * as Dodo_Internal from "../Dodo.Internal/index.js";
import * as PureScript_Backend_Optimizer_CoreFn from "../PureScript.Backend.Optimizer.CoreFn/index.js";
import * as PureScript_Backend_Optimizer_Semantics from "../PureScript.Backend.Optimizer.Semantics/index.js";
import * as PureScript_Backend_Optimizer_Syntax from "../PureScript.Backend.Optimizer.Syntax/index.js";
var lines = /* #__PURE__ */ Dodo.lines(Data_Foldable.foldableArray);
var pure = /* #__PURE__ */ Control_Applicative.pure(Control_Applicative.applicativeArray);
var show = /* #__PURE__ */ Data_Show.show(Data_Show.showInt);
var mempty = /* #__PURE__ */ Data_Monoid.mempty(Dodo_Internal.monoidDoc);
var PrecBlock = /* #__PURE__ */ (function () {
    function PrecBlock() {

    };
    PrecBlock.value = new PrecBlock();
    return PrecBlock;
})();
var PrecApply = /* #__PURE__ */ (function () {
    function PrecApply() {

    };
    PrecApply.value = new PrecApply();
    return PrecApply;
})();
var PrecAccess = /* #__PURE__ */ (function () {
    function PrecAccess() {

    };
    PrecAccess.value = new PrecAccess();
    return PrecAccess;
})();
var PrecAtom = /* #__PURE__ */ (function () {
    function PrecAtom() {

    };
    PrecAtom.value = new PrecAtom();
    return PrecAtom;
})();
var eqPrec = {
    eq: function (x) {
        return function (y) {
            if (x instanceof PrecBlock && y instanceof PrecBlock) {
                return true;
            };
            if (x instanceof PrecApply && y instanceof PrecApply) {
                return true;
            };
            if (x instanceof PrecAccess && y instanceof PrecAccess) {
                return true;
            };
            if (x instanceof PrecAtom && y instanceof PrecAtom) {
                return true;
            };
            return false;
        };
    }
};
var ordPrec = {
    compare: function (x) {
        return function (y) {
            if (x instanceof PrecBlock && y instanceof PrecBlock) {
                return Data_Ordering.EQ.value;
            };
            if (x instanceof PrecBlock) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof PrecBlock) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof PrecApply && y instanceof PrecApply) {
                return Data_Ordering.EQ.value;
            };
            if (x instanceof PrecApply) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof PrecApply) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof PrecAccess && y instanceof PrecAccess) {
                return Data_Ordering.EQ.value;
            };
            if (x instanceof PrecAccess) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof PrecAccess) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof PrecAtom && y instanceof PrecAtom) {
                return Data_Ordering.EQ.value;
            };
            throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Tracer.Printer (line 0, column 0 - line 0, column 0): " + [ x.constructor.name, y.constructor.name ]);
        };
    },
    Eq0: function () {
        return eqPrec;
    }
};
var wrapPrec = function (prec1) {
    return function (v) {
        if (Data_Ord.lessThan(ordPrec)(v.value0)(prec1)) {
            return Dodo_Common.pursParens(v.value1);
        };
        if (Data_Boolean.otherwise) {
            return v.value1;
        };
        throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Tracer.Printer (line 94, column 1 - line 94, column 40): " + [ prec1.constructor.name, v.constructor.name ]);
    };
};
var printUncurriedApp = function (isEffectful) {
    return function (fn) {
        return function (args) {
            return new Data_Tuple.Tuple(PrecApply.value, Dodo.flexGroup(Data_Array.fold(Dodo_Internal.monoidDoc)([ wrapPrec(PrecAccess.value)(fn), Dodo.softBreak, Dodo.indent(Data_Array.fold(Dodo_Internal.monoidDoc)([ Dodo_Common.pursParens(Dodo.foldWithSeparator(Data_Foldable.foldableArray)(Dodo_Common.leadingComma)(Data_Functor.map(Data_Functor.functorArray)(Data_Tuple.snd)(args))), Data_Monoid.guard(Dodo_Internal.monoidDoc)(isEffectful)(Dodo.text("!")) ])) ])));
        };
    };
};
var printUncurriedAbs = function (isEffectful) {
    return function (args) {
        return function (body) {
            return new Data_Tuple.Tuple(PrecBlock.value, Dodo.flexGroup(Data_Array.fold(Dodo_Internal.monoidDoc)([ Dodo.text("\\"), Dodo.flexGroup(Dodo.alignCurrentColumn(Dodo_Common.pursParens(Dodo.foldWithSeparator(Data_Foldable.foldableArray)(Dodo_Common.leadingComma)(args)))), Data_Monoid.guard(Dodo_Internal.monoidDoc)(isEffectful)(Dodo.text("!")), Dodo.space, Dodo.text("->"), Dodo.spaceBreak, Dodo.indent(Data_Tuple.snd(body)) ])));
        };
    };
};
var printRewrite = function (str) {
    return Dodo.text("{- " + (str + " -}"));
};
var printUnpackOpCase = function (v) {
    if (v instanceof PureScript_Backend_Optimizer_Semantics.UnpackRecord) {
        return printRewrite("UnpackRecord");
    };
    if (v instanceof PureScript_Backend_Optimizer_Semantics.UnpackUpdate) {
        return printRewrite("UnpackUpdate");
    };
    if (v instanceof PureScript_Backend_Optimizer_Semantics.UnpackArray) {
        return printRewrite("UnpackArray");
    };
    if (v instanceof PureScript_Backend_Optimizer_Semantics.UnpackData) {
        return printRewrite("UnpackData");
    };
    throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Tracer.Printer (line 307, column 21 - line 315, column 30): " + [ v.constructor.name ]);
};
var printRecord = function (sep) {
    var printProp = function (v) {
        return Dodo.flexGroup(Data_Array.fold(Dodo_Internal.monoidDoc)([ Dodo.text(Data_Show.show(Data_Show.showString)(v.value0)), sep, Dodo.spaceBreak, Dodo.align(2)(Dodo.indent(v.value1)) ]));
    };
    var $209 = Dodo.foldWithSeparator(Data_Foldable.foldableArray)(Dodo_Common.leadingComma);
    var $210 = Data_Functor.map(Data_Functor.functorArray)(printProp);
    return function ($211) {
        return Dodo_Common.pursCurlies($209($210($211)));
    };
};
var printQualified = function (printA) {
    return function (v) {
        return Data_Maybe.maybe(printA(v.value1))(function (v1) {
            return Data_Semigroup.append(Dodo_Internal.semigroupDoc)(Dodo.text(v1))(Data_Semigroup.append(Dodo_Internal.semigroupDoc)(Dodo.text("."))(printA(v.value1)));
        })(v.value0);
    };
};
var printLevel = function (v) {
    return Dodo.text(Data_Show.show(Data_Show.showInt)(v));
};
var printLet$prime = function (letKeyword) {
    return function (bindings) {
        return Dodo.flexGroup(Data_Array.fold(Dodo_Internal.monoidDoc)([ Dodo.text(letKeyword), Dodo.spaceBreak, Dodo.indent(lines(Data_Functor.map(Data_Functor.functorArray)(function (v) {
            return Dodo.flexGroup(Data_Array.fold(Dodo_Internal.monoidDoc)([ v.value0, Dodo.space, Dodo.text("="), Dodo.spaceBreak, Dodo.indent(v.value1) ]));
        })(bindings))), Dodo.space, Dodo.text("in") ]));
    };
};
var printLet = function (letKeyword) {
    return Data_Tuple.curry((function () {
        var $212 = printLet$prime(letKeyword);
        return function ($213) {
            return $212(pure($213));
        };
    })());
};
var printIdent = function (v) {
    return Dodo.text(v);
};
var printLocal = function (mbIdent) {
    return function (lvl) {
        if (mbIdent instanceof Data_Maybe.Just && mbIdent.value0 === "$__unused") {
            return Dodo.text("_");
        };
        return Data_Semigroup.append(Dodo_Internal.semigroupDoc)(Data_Foldable.foldMap(Data_Foldable.foldableMaybe)(Dodo_Internal.monoidDoc)(printIdent)(mbIdent))(Data_Semigroup.append(Dodo_Internal.semigroupDoc)(Dodo.text("@"))(printLevel(lvl)));
    };
};
var printQualifiedIdent = /* #__PURE__ */ printQualified(printIdent);
var printDistOpCase = function (v) {
    if (v instanceof PureScript_Backend_Optimizer_Semantics.DistApp) {
        return printRewrite("DistApp");
    };
    if (v instanceof PureScript_Backend_Optimizer_Semantics.DistUncurriedApp) {
        return printRewrite("DistUncurriedApp");
    };
    if (v instanceof PureScript_Backend_Optimizer_Semantics.DistAccessor) {
        return printRewrite("DistAccessor");
    };
    if (v instanceof PureScript_Backend_Optimizer_Semantics.DistPrimOp1) {
        return printRewrite("DistPrimOp1");
    };
    if (v instanceof PureScript_Backend_Optimizer_Semantics.DistPrimOp2L) {
        return printRewrite("DistPrimOp2L");
    };
    if (v instanceof PureScript_Backend_Optimizer_Semantics.DistPrimOp2R) {
        return printRewrite("DistPrimOp2R");
    };
    throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Tracer.Printer (line 318, column 19 - line 330, column 32): " + [ v.constructor.name ]);
};
var printCurriedApp = function (fn) {
    return function (args) {
        return new Data_Tuple.Tuple(PrecApply.value, Dodo.flexGroup(Dodo.foldWithSeparator(Data_Array_NonEmpty_Internal.foldableNonEmptyArray)(Dodo.spaceBreak)(Data_Array_NonEmpty.cons(wrapPrec(PrecApply.value)(fn))(Data_Functor.map(Data_Array_NonEmpty_Internal.functorNonEmptyArray)((function () {
            var $214 = wrapPrec(PrecAccess.value);
            return function ($215) {
                return Dodo.indent($214($215));
            };
        })())(args)))));
    };
};
var printCurriedAbs = function (args) {
    return function (body) {
        return new Data_Tuple.Tuple(PrecBlock.value, Dodo.flexGroup(Data_Array.fold(Dodo_Internal.monoidDoc)([ Dodo.text("\\"), Dodo.flexGroup(Dodo.alignCurrentColumn(Dodo.foldWithSeparator(Data_Array_NonEmpty_Internal.foldableNonEmptyArray)(Dodo.spaceBreak)(args))), Dodo.space, Dodo.text("->"), Dodo.spaceBreak, Dodo.indent(Data_Tuple.snd(body)) ])));
    };
};
var printBranch = function (conds) {
    return function (fallback) {
        return Data_Foldable.foldr(Data_Array_NonEmpty_Internal.foldableNonEmptyArray)(function (v) {
            return function (other) {
                return Data_Array.fold(Dodo_Internal.monoidDoc)([ Dodo.flexGroup(Data_Array.fold(Dodo_Internal.monoidDoc)([ Dodo.text("if"), Dodo.spaceBreak, Dodo.indent(v.value0), Dodo.space, Dodo.text("then") ])), Dodo.indent(v.value1), Dodo["break"], Dodo.text("else"), Dodo.space, other ]);
            };
        })(Data_Semigroup.append(Dodo_Internal.semigroupDoc)(Dodo["break"])(Dodo.indent(fallback)))(conds);
    };
};
var printBackendRewriteCase = function (v) {
    if (v instanceof PureScript_Backend_Optimizer_Semantics.RewriteInline) {
        return printRewrite("Inline");
    };
    if (v instanceof PureScript_Backend_Optimizer_Semantics.RewriteUncurry) {
        return printRewrite("Uncurry");
    };
    if (v instanceof PureScript_Backend_Optimizer_Semantics.RewriteStop) {
        return printRewrite("Stop");
    };
    if (v instanceof PureScript_Backend_Optimizer_Semantics.RewriteUnpackOp) {
        return printUnpackOpCase(v.value2);
    };
    if (v instanceof PureScript_Backend_Optimizer_Semantics.RewriteDistBranchesLet) {
        return printRewrite("DistLet");
    };
    if (v instanceof PureScript_Backend_Optimizer_Semantics.RewriteDistBranchesOp) {
        return printDistOpCase(v.value2);
    };
    throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Tracer.Printer (line 292, column 27 - line 304, column 27): " + [ v.constructor.name ]);
};
var printBackendAccessor = function (v) {
    if (v instanceof PureScript_Backend_Optimizer_Syntax.GetProp) {
        return Dodo.text("." + Data_Show.show(Data_Show.showString)(v.value0));
    };
    if (v instanceof PureScript_Backend_Optimizer_Syntax.GetIndex) {
        return Dodo.text("[" + (Data_Show.show(Data_Show.showInt)(v.value0) + "]"));
    };
    if (v instanceof PureScript_Backend_Optimizer_Syntax.GetCtorField) {
        return Dodo.text("#" + Data_Show.show(Data_Show.showInt)(v.value5));
    };
    throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Tracer.Printer (line 404, column 24 - line 410, column 28): " + [ v.constructor.name ]);
};
var printArray = /* #__PURE__ */ (function () {
    var $216 = Dodo.foldWithSeparator(Data_Foldable.foldableArray)(Dodo_Common.leadingComma);
    return function ($217) {
        return Dodo_Common.pursSquares($216($217));
    };
})();
var printLiteral = /* #__PURE__ */ (function () {
    var $218 = Data_Tuple.Tuple.create(PrecAtom.value);
    return function ($219) {
        return $218((function (v) {
            if (v instanceof PureScript_Backend_Optimizer_CoreFn.LitInt) {
                return Dodo.text(Data_Show.show(Data_Show.showInt)(v.value0));
            };
            if (v instanceof PureScript_Backend_Optimizer_CoreFn.LitNumber) {
                return Dodo.text(Data_Show.show(Data_Show.showNumber)(v.value0));
            };
            if (v instanceof PureScript_Backend_Optimizer_CoreFn.LitString) {
                return Dodo.text(Data_Show.show(Data_Show.showString)(v.value0));
            };
            if (v instanceof PureScript_Backend_Optimizer_CoreFn.LitChar) {
                return Dodo.text(Data_Show.show(Data_Show.showChar)(v.value0));
            };
            if (v instanceof PureScript_Backend_Optimizer_CoreFn.LitBoolean) {
                return Dodo.text(Data_Show.show(Data_Show.showBoolean)(v.value0));
            };
            if (v instanceof PureScript_Backend_Optimizer_CoreFn.LitArray) {
                return printArray(Data_Functor.map(Data_Functor.functorArray)(Data_Tuple.snd)(v.value0));
            };
            if (v instanceof PureScript_Backend_Optimizer_CoreFn.LitRecord) {
                return printRecord(Dodo.text(":"))(Data_Functor.map(Data_Functor.functorArray)(Data_Functor.map(PureScript_Backend_Optimizer_CoreFn.functorProp)(Data_Tuple.snd))(v.value0));
            };
            throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Tracer.Printer (line 413, column 35 - line 427, column 53): " + [ v.constructor.name ]);
        })($219));
    };
})();
var primOp = function (ns) {
    return function (op) {
        if (ns === "") {
            return new Data_Tuple.Tuple(PrecAtom.value, Dodo.text("#[prim." + (op + "]")));
        };
        if (Data_Boolean.otherwise) {
            return new Data_Tuple.Tuple(PrecAtom.value, Dodo.text("#[prim." + (ns + ("." + (op + "]")))));
        };
        throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Tracer.Printer (line 380, column 1 - line 380, column 38): " + [ ns.constructor.name, op.constructor.name ]);
    };
};
var printBackendEffect = function (v) {
    if (v instanceof PureScript_Backend_Optimizer_Syntax.EffectRefNew) {
        return printUncurriedApp(true)(primOp("ref")("new"))([ v.value0 ]);
    };
    if (v instanceof PureScript_Backend_Optimizer_Syntax.EffectRefRead) {
        return printUncurriedApp(true)(primOp("ref")("read"))([ v.value0 ]);
    };
    if (v instanceof PureScript_Backend_Optimizer_Syntax.EffectRefWrite) {
        return printUncurriedApp(true)(primOp("ref")("write"))([ v.value0, v.value1 ]);
    };
    throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Tracer.Printer (line 336, column 22 - line 342, column 59): " + [ v.constructor.name ]);
};
var printBackendOperator1 = function (v) {
    if (v instanceof PureScript_Backend_Optimizer_Syntax.OpBooleanNot) {
        return primOp("boolean")("not");
    };
    if (v instanceof PureScript_Backend_Optimizer_Syntax.OpIntBitNot) {
        return primOp("int")("bitnot");
    };
    if (v instanceof PureScript_Backend_Optimizer_Syntax.OpIntNegate) {
        return primOp("int")("negate");
    };
    if (v instanceof PureScript_Backend_Optimizer_Syntax.OpNumberNegate) {
        return primOp("number")("negate");
    };
    if (v instanceof PureScript_Backend_Optimizer_Syntax.OpArrayLength) {
        return primOp("array")("length");
    };
    if (v instanceof PureScript_Backend_Optimizer_Syntax.OpIsTag) {
        return new Data_Tuple.Tuple(PrecAtom.value, Data_Semigroup.append(Dodo_Internal.semigroupDoc)(Dodo.text("#[prim.istag "))(Data_Semigroup.append(Dodo_Internal.semigroupDoc)(printQualifiedIdent(v.value0))(Dodo.text("]"))));
    };
    throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Tracer.Printer (line 352, column 25 - line 358, column 96): " + [ v.constructor.name ]);
};
var printBackendOperatorNum = function (ns) {
    var $220 = primOp(ns);
    return function ($221) {
        return $220((function (v) {
            if (v instanceof PureScript_Backend_Optimizer_Syntax.OpAdd) {
                return "add";
            };
            if (v instanceof PureScript_Backend_Optimizer_Syntax.OpDivide) {
                return "div";
            };
            if (v instanceof PureScript_Backend_Optimizer_Syntax.OpMultiply) {
                return "mul";
            };
            if (v instanceof PureScript_Backend_Optimizer_Syntax.OpSubtract) {
                return "sub";
            };
            throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Tracer.Printer (line 388, column 44 - line 392, column 22): " + [ v.constructor.name ]);
        })($221));
    };
};
var printBackendOperatorOrd = function (ns) {
    var $222 = primOp(ns);
    return function ($223) {
        return $222((function (v) {
            if (v instanceof PureScript_Backend_Optimizer_Syntax.OpEq) {
                return "eq";
            };
            if (v instanceof PureScript_Backend_Optimizer_Syntax.OpNotEq) {
                return "neq";
            };
            if (v instanceof PureScript_Backend_Optimizer_Syntax.OpGt) {
                return "gt";
            };
            if (v instanceof PureScript_Backend_Optimizer_Syntax.OpGte) {
                return "gte";
            };
            if (v instanceof PureScript_Backend_Optimizer_Syntax.OpLt) {
                return "lt";
            };
            if (v instanceof PureScript_Backend_Optimizer_Syntax.OpLte) {
                return "lte";
            };
            throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Tracer.Printer (line 395, column 44 - line 401, column 17): " + [ v.constructor.name ]);
        })($223));
    };
};
var printBackendOperator2 = function (v) {
    if (v instanceof PureScript_Backend_Optimizer_Syntax.OpArrayIndex) {
        return primOp("array")("index");
    };
    if (v instanceof PureScript_Backend_Optimizer_Syntax.OpBooleanAnd) {
        return primOp("boolean")("and");
    };
    if (v instanceof PureScript_Backend_Optimizer_Syntax.OpBooleanOr) {
        return primOp("boolean")("or");
    };
    if (v instanceof PureScript_Backend_Optimizer_Syntax.OpBooleanOrd) {
        return printBackendOperatorOrd("boolean")(v.value0);
    };
    if (v instanceof PureScript_Backend_Optimizer_Syntax.OpCharOrd) {
        return printBackendOperatorOrd("char")(v.value0);
    };
    if (v instanceof PureScript_Backend_Optimizer_Syntax.OpIntBitAnd) {
        return primOp("int")("bitand");
    };
    if (v instanceof PureScript_Backend_Optimizer_Syntax.OpIntBitOr) {
        return primOp("int")("bitor");
    };
    if (v instanceof PureScript_Backend_Optimizer_Syntax.OpIntBitShiftLeft) {
        return primOp("int")("bitshl");
    };
    if (v instanceof PureScript_Backend_Optimizer_Syntax.OpIntBitShiftRight) {
        return primOp("int")("bitshr");
    };
    if (v instanceof PureScript_Backend_Optimizer_Syntax.OpIntBitXor) {
        return primOp("int")("bitxor");
    };
    if (v instanceof PureScript_Backend_Optimizer_Syntax.OpIntBitZeroFillShiftRight) {
        return primOp("int")("bitzshr");
    };
    if (v instanceof PureScript_Backend_Optimizer_Syntax.OpIntNum) {
        return printBackendOperatorNum("int")(v.value0);
    };
    if (v instanceof PureScript_Backend_Optimizer_Syntax.OpIntOrd) {
        return printBackendOperatorOrd("int")(v.value0);
    };
    if (v instanceof PureScript_Backend_Optimizer_Syntax.OpNumberNum) {
        return printBackendOperatorNum("number")(v.value0);
    };
    if (v instanceof PureScript_Backend_Optimizer_Syntax.OpNumberOrd) {
        return printBackendOperatorOrd("number")(v.value0);
    };
    if (v instanceof PureScript_Backend_Optimizer_Syntax.OpStringAppend) {
        return primOp("string")("append");
    };
    if (v instanceof PureScript_Backend_Optimizer_Syntax.OpStringOrd) {
        return printBackendOperatorOrd("string")(v.value0);
    };
    throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Tracer.Printer (line 361, column 25 - line 378, column 58): " + [ v.constructor.name ]);
};
var printBackendOperator = function (v) {
    if (v instanceof PureScript_Backend_Optimizer_Syntax.Op1) {
        return printUncurriedApp(false)(printBackendOperator1(v.value0))([ v.value1 ]);
    };
    if (v instanceof PureScript_Backend_Optimizer_Syntax.Op2) {
        return printUncurriedApp(false)(printBackendOperator2(v.value0))([ v.value1, v.value2 ]);
    };
    throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Tracer.Printer (line 345, column 24 - line 349, column 65): " + [ v.constructor.name ]);
};
var printBackendSyntax = function (v) {
    if (v instanceof PureScript_Backend_Optimizer_Syntax.Var) {
        return new Data_Tuple.Tuple(PrecAtom.value, printQualifiedIdent(v.value0));
    };
    if (v instanceof PureScript_Backend_Optimizer_Syntax.Local) {
        return new Data_Tuple.Tuple(PrecAtom.value, printLocal(v.value0)(v.value1));
    };
    if (v instanceof PureScript_Backend_Optimizer_Syntax.Lit) {
        return printLiteral(v.value0);
    };
    if (v instanceof PureScript_Backend_Optimizer_Syntax.App) {
        return printCurriedApp(v.value0)(v.value1);
    };
    if (v instanceof PureScript_Backend_Optimizer_Syntax.Abs) {
        return printCurriedAbs(Data_Functor.map(Data_Array_NonEmpty_Internal.functorNonEmptyArray)(Data_Tuple.uncurry(printLocal))(v.value0))(v.value1);
    };
    if (v instanceof PureScript_Backend_Optimizer_Syntax.UncurriedApp) {
        return printUncurriedApp(false)(v.value0)(v.value1);
    };
    if (v instanceof PureScript_Backend_Optimizer_Syntax.UncurriedAbs) {
        return printUncurriedAbs(false)(Data_Functor.map(Data_Functor.functorArray)(Data_Tuple.uncurry(printLocal))(v.value0))(v.value1);
    };
    if (v instanceof PureScript_Backend_Optimizer_Syntax.UncurriedEffectApp) {
        return printUncurriedApp(true)(v.value0)(v.value1);
    };
    if (v instanceof PureScript_Backend_Optimizer_Syntax.UncurriedEffectAbs) {
        return printUncurriedAbs(true)(Data_Functor.map(Data_Functor.functorArray)(Data_Tuple.uncurry(printLocal))(v.value0))(v.value1);
    };
    if (v instanceof PureScript_Backend_Optimizer_Syntax.Accessor) {
        return new Data_Tuple.Tuple(PrecAccess.value, Data_Array.fold(Dodo_Internal.monoidDoc)([ wrapPrec(PrecAccess.value)(v.value0), printBackendAccessor(v.value1) ]));
    };
    if (v instanceof PureScript_Backend_Optimizer_Syntax.Update) {
        return new Data_Tuple.Tuple(PrecApply.value, Dodo.flexGroup(Data_Array.fold(Dodo_Internal.monoidDoc)([ wrapPrec(PrecAccess.value)(v.value0), Dodo.spaceBreak, Dodo.indent(printRecord(Dodo.text(" ="))(Data_Functor.map(Data_Functor.functorArray)(Data_Functor.map(PureScript_Backend_Optimizer_CoreFn.functorProp)(Data_Tuple.snd))(v.value1))) ])));
    };
    if (v instanceof PureScript_Backend_Optimizer_Syntax.CtorSaturated) {
        return printUncurriedApp(false)(new Data_Tuple.Tuple(PrecAtom.value, printQualifiedIdent(v.value0)))(Data_Functor.map(Data_Functor.functorArray)(Data_Tuple.snd)(v.value4));
    };
    if (v instanceof PureScript_Backend_Optimizer_Syntax.CtorDef) {
        return new Data_Tuple.Tuple(PrecBlock.value, Dodo.words(Data_Foldable.foldableArray)([ Dodo.text("constructor"), Data_Semigroup.append(Dodo_Internal.semigroupDoc)(printIdent(v.value2))(Dodo.encloseWithSeparator(Data_Foldable.foldableArray)(Dodo.text("("))(Dodo.text(")"))(Dodo.text(", "))(Data_Array.replicate(Data_Array.length(v.value3))(Dodo.text("_")))), Dodo.text("of"), Dodo.text(v.value1) ]));
    };
    if (v instanceof PureScript_Backend_Optimizer_Syntax.LetRec) {
        return new Data_Tuple.Tuple(PrecBlock.value, Dodo.lines(Data_Foldable.foldableArray)([ printLet$prime("letrec")(Data_Array_NonEmpty.toArray(Data_Functor.map(Data_Array_NonEmpty_Internal.functorNonEmptyArray)(function (v1) {
            return new Data_Tuple.Tuple(printLocal(new Data_Maybe.Just(v1.value0))(v.value0), Data_Tuple.snd(v1.value1));
        })(v.value1))), Data_Tuple.snd(v.value2) ]));
    };
    if (v instanceof PureScript_Backend_Optimizer_Syntax.Let) {
        return new Data_Tuple.Tuple(PrecBlock.value, Dodo.lines(Data_Foldable.foldableArray)([ printLet("let")(printLocal(v.value0)(v.value1))(Data_Tuple.snd(v.value2)), Data_Tuple.snd(v.value3) ]));
    };
    if (v instanceof PureScript_Backend_Optimizer_Syntax.EffectBind) {
        return new Data_Tuple.Tuple(PrecBlock.value, Dodo.lines(Data_Foldable.foldableArray)([ printLet("let!")(printLocal(v.value0)(v.value1))(Data_Tuple.snd(v.value2)), Data_Tuple.snd(v.value3) ]));
    };
    if (v instanceof PureScript_Backend_Optimizer_Syntax.EffectPure) {
        return printUncurriedApp(true)(primOp("effect")("pure"))([ v.value0 ]);
    };
    if (v instanceof PureScript_Backend_Optimizer_Syntax.EffectDefer) {
        return printUncurriedApp(true)(primOp("effect")("defer"))([ v.value0 ]);
    };
    if (v instanceof PureScript_Backend_Optimizer_Syntax.Branch) {
        return new Data_Tuple.Tuple(PrecBlock.value, printBranch(Data_Functor.map(Data_Array_NonEmpty_Internal.functorNonEmptyArray)(Data_Functor.map(PureScript_Backend_Optimizer_Syntax.functorPair)(Data_Tuple.snd))(v.value0))(Data_Tuple.snd(v.value1)));
    };
    if (v instanceof PureScript_Backend_Optimizer_Syntax.PrimOp) {
        return printBackendOperator(v.value0);
    };
    if (v instanceof PureScript_Backend_Optimizer_Syntax.PrimEffect) {
        return printBackendEffect(v.value0);
    };
    if (v instanceof PureScript_Backend_Optimizer_Syntax.PrimUndefined) {
        return primOp("")("undefined");
    };
    if (v instanceof PureScript_Backend_Optimizer_Syntax.Fail) {
        return printUncurriedApp(false)(primOp("")("fail"))([  ]);
    };
    if (v instanceof PureScript_Backend_Optimizer_Syntax.Typed) {
        return v.value1;
    };
    throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Tracer.Printer (line 211, column 22 - line 289, column 6): " + [ v.constructor.name ]);
};
var printBackendExpr = /* #__PURE__ */ PureScript_Backend_Optimizer_Semantics.foldBackendExpr(printBackendSyntax)(function (rewrite) {
    return function (expr) {
        return new Data_Tuple.Tuple(Data_Tuple.fst(expr), Dodo.lines(Data_Foldable.foldableArray)([ printBackendRewriteCase(rewrite), Data_Tuple.snd(expr) ]));
    };
});
var heading = function (repeat) {
    return function (hd) {
        return Dodo.words(Data_Foldable.foldableArray)([ Dodo.text(Data_Monoid.power(Data_Monoid.monoidString)(repeat)(2)), hd, Dodo.withPosition(function (v) {
            return Dodo.text(Data_Monoid.power(Data_Monoid.monoidString)(repeat)(v.pageWidth - v.column | 0));
        }) ]);
    };
};
var printSteps = function (qual) {
    return function (steps) {
        var ident = printQualifiedIdent(qual);
        var renderStep = function (idx) {
            return function (step) {
                return Dodo.lines(Data_Foldable.foldableArray)([ heading((function () {
                    var $202 = idx === 0;
                    if ($202) {
                        return "+";
                    };
                    return "-";
                })())(Dodo.words(Data_Foldable.foldableArray)([ ident, Dodo.text("Step"), Dodo.text(show(idx + 1 | 0)), (function () {
                    var v = idx === (Data_Array_NonEmpty.length(steps) - 1 | 0);
                    var v1 = idx === 0;
                    if (v1 && v) {
                        return Dodo.text("(Original/Final)");
                    };
                    if (v1) {
                        return Dodo.text("(Original)");
                    };
                    if (v) {
                        return Dodo.text("(Final)");
                    };
                    return mempty;
                })() ])), Dodo.indent(Data_Tuple.snd(printBackendExpr(step))) ]);
            };
        };
        return Dodo.foldWithSeparator(Data_Foldable.foldableArray)(Data_Semigroup.append(Dodo_Internal.semigroupDoc)(Dodo["break"])(Dodo["break"]))(Data_Array.mapWithIndex(renderStep)(Data_Array_NonEmpty.toArray(steps)));
    };
};
var printModuleSteps = function (v) {
    return function (steps) {
        return Data_Array.fold(Dodo_Internal.monoidDoc)([ heading("=")(Dodo.text(v)), Dodo["break"], Dodo.withPosition(function (v1) {
            return Dodo.text(Data_Monoid.power(Data_Monoid.monoidString)("=")(v1.pageWidth));
        }), Dodo["break"], Dodo["break"], Dodo.foldWithSeparator(Data_Foldable.foldableArray)(Data_Semigroup.append(Dodo_Internal.semigroupDoc)(Dodo["break"])(Dodo["break"]))(Data_Functor.map(Data_Functor.functorArray)(Data_Tuple.uncurry(printSteps))(steps)) ]);
    };
};
export {
    printModuleSteps,
    printSteps
};
