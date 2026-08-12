import * as Control_Alternative from "../Control.Alternative/index.js";
import * as Control_Applicative from "../Control.Applicative/index.js";
import * as Control_Bind from "../Control.Bind/index.js";
import * as Control_Category from "../Control.Category/index.js";
import * as Data_Array from "../Data.Array/index.js";
import * as Data_Array_NonEmpty from "../Data.Array.NonEmpty/index.js";
import * as Data_Array_NonEmpty_Internal from "../Data.Array.NonEmpty.Internal/index.js";
import * as Data_Boolean from "../Data.Boolean/index.js";
import * as Data_Bounded from "../Data.Bounded/index.js";
import * as Data_Either from "../Data.Either/index.js";
import * as Data_Eq from "../Data.Eq/index.js";
import * as Data_EuclideanRing from "../Data.EuclideanRing/index.js";
import * as Data_Foldable from "../Data.Foldable/index.js";
import * as Data_Function from "../Data.Function/index.js";
import * as Data_Functor from "../Data.Functor/index.js";
import * as Data_HeytingAlgebra from "../Data.HeytingAlgebra/index.js";
import * as Data_Lazy from "../Data.Lazy/index.js";
import * as Data_List from "../Data.List/index.js";
import * as Data_List_Types from "../Data.List.Types/index.js";
import * as Data_Map_Internal from "../Data.Map.Internal/index.js";
import * as Data_Maybe from "../Data.Maybe/index.js";
import * as Data_Monoid from "../Data.Monoid/index.js";
import * as Data_Newtype from "../Data.Newtype/index.js";
import * as Data_Ord from "../Data.Ord/index.js";
import * as Data_Ordering from "../Data.Ordering/index.js";
import * as Data_Semigroup from "../Data.Semigroup/index.js";
import * as Data_Set from "../Data.Set/index.js";
import * as Data_Show from "../Data.Show/index.js";
import * as Data_String_CodePoints from "../Data.String.CodePoints/index.js";
import * as Data_Tuple from "../Data.Tuple/index.js";
import * as Data_Unfoldable from "../Data.Unfoldable/index.js";
import * as Partial_Unsafe from "../Partial.Unsafe/index.js";
import * as PureScript_Backend_Optimizer_Analysis from "../PureScript.Backend.Optimizer.Analysis/index.js";
import * as PureScript_Backend_Optimizer_CoreFn from "../PureScript.Backend.Optimizer.CoreFn/index.js";
import * as PureScript_Backend_Optimizer_Syntax from "../PureScript.Backend.Optimizer.Syntax/index.js";
import * as PureScript_Backend_Optimizer_Utils from "../PureScript.Backend.Optimizer.Utils/index.js";
var $runtime_lazy = function (name, moduleName, init) {
    var state = 0;
    var val;
    return function (lineNumber) {
        if (state === 2) return val;
        if (state === 1) throw new ReferenceError(name + " was needed before it finished initializing (module " + moduleName + ", line " + lineNumber + ")", moduleName, lineNumber);
        state = 1;
        val = init();
        state = 2;
        return val;
    };
};
var eqQualified = /* #__PURE__ */ PureScript_Backend_Optimizer_CoreFn.eqQualified(PureScript_Backend_Optimizer_CoreFn.eqIdent);
var eqTuple = /* #__PURE__ */ Data_Tuple.eqTuple(Data_Eq.eqString);
var eqMaybe = /* #__PURE__ */ Data_Maybe.eqMaybe(PureScript_Backend_Optimizer_CoreFn.eqIdent);
var ordQualified = /* #__PURE__ */ PureScript_Backend_Optimizer_CoreFn.ordQualified(PureScript_Backend_Optimizer_CoreFn.ordIdent);
var ordMaybe = /* #__PURE__ */ Data_Maybe.ordMaybe(PureScript_Backend_Optimizer_CoreFn.ordIdent);
var eqNonEmptyArray = /* #__PURE__ */ Data_Array_NonEmpty_Internal.eqNonEmptyArray(/* #__PURE__ */ Data_Tuple.eqTuple(/* #__PURE__ */ Data_Maybe.eqMaybe(PureScript_Backend_Optimizer_CoreFn.eqIdent))(PureScript_Backend_Optimizer_Syntax.eqLevel));
var append = /* #__PURE__ */ Data_Semigroup.append(PureScript_Backend_Optimizer_Analysis.semigroupBackendAnalysis);
var and = /* #__PURE__ */ Data_Foldable.and(Data_Array_NonEmpty_Internal.foldableNonEmptyArray)(Data_HeytingAlgebra.heytingAlgebraBoolean);
var unwrap = /* #__PURE__ */ Data_Newtype.unwrap();
var eqEither = /* #__PURE__ */ Data_Either.eqEither(/* #__PURE__ */ PureScript_Backend_Optimizer_CoreFn.eqQualified(PureScript_Backend_Optimizer_CoreFn.eqIdent))(PureScript_Backend_Optimizer_Syntax.eqBackendOperator2);
var identity = /* #__PURE__ */ Control_Category.identity(Control_Category.categoryFn);
var lookup = /* #__PURE__ */ Data_Foldable.lookup(Data_Array_NonEmpty_Internal.foldableNonEmptyArray)(PureScript_Backend_Optimizer_CoreFn.eqIdent);
var toUnfoldable = /* #__PURE__ */ Data_Array.toUnfoldable(Data_List_Types.unfoldableList);
var unwrap1 = /* #__PURE__ */ Data_Newtype.unwrap();
var mempty = /* #__PURE__ */ Data_Monoid.mempty(PureScript_Backend_Optimizer_Analysis.monoidBackendAnalysis);
var top = /* #__PURE__ */ Data_Bounded.top(Data_Bounded.boundedInt);
var UnpackRecord = /* #__PURE__ */ (function () {
    function UnpackRecord(value0) {
        this.value0 = value0;
    };
    UnpackRecord.create = function (value0) {
        return new UnpackRecord(value0);
    };
    return UnpackRecord;
})();
var UnpackUpdate = /* #__PURE__ */ (function () {
    function UnpackUpdate(value0, value1) {
        this.value0 = value0;
        this.value1 = value1;
    };
    UnpackUpdate.create = function (value0) {
        return function (value1) {
            return new UnpackUpdate(value0, value1);
        };
    };
    return UnpackUpdate;
})();
var UnpackArray = /* #__PURE__ */ (function () {
    function UnpackArray(value0) {
        this.value0 = value0;
    };
    UnpackArray.create = function (value0) {
        return new UnpackArray(value0);
    };
    return UnpackArray;
})();
var UnpackData = /* #__PURE__ */ (function () {
    function UnpackData(value0, value1, value2, value3, value4) {
        this.value0 = value0;
        this.value1 = value1;
        this.value2 = value2;
        this.value3 = value3;
        this.value4 = value4;
    };
    UnpackData.create = function (value0) {
        return function (value1) {
            return function (value2) {
                return function (value3) {
                    return function (value4) {
                        return new UnpackData(value0, value1, value2, value3, value4);
                    };
                };
            };
        };
    };
    return UnpackData;
})();
var SemConditional = /* #__PURE__ */ (function () {
    function SemConditional(value0, value1) {
        this.value0 = value0;
        this.value1 = value1;
    };
    SemConditional.create = function (value0) {
        return function (value1) {
            return new SemConditional(value0, value1);
        };
    };
    return SemConditional;
})();
var NeutralExpr = function (x) {
    return x;
};
var MkFnApplied = /* #__PURE__ */ (function () {
    function MkFnApplied(value0) {
        this.value0 = value0;
    };
    MkFnApplied.create = function (value0) {
        return new MkFnApplied(value0);
    };
    return MkFnApplied;
})();
var MkFnNext = /* #__PURE__ */ (function () {
    function MkFnNext(value0, value1) {
        this.value0 = value0;
        this.value1 = value1;
    };
    MkFnNext.create = function (value0) {
        return function (value1) {
            return new MkFnNext(value0, value1);
        };
    };
    return MkFnNext;
})();
var One = /* #__PURE__ */ (function () {
    function One(value0) {
        this.value0 = value0;
    };
    One.create = function (value0) {
        return new One(value0);
    };
    return One;
})();
var Group = /* #__PURE__ */ (function () {
    function Group(value0) {
        this.value0 = value0;
    };
    Group.create = function (value0) {
        return new Group(value0);
    };
    return Group;
})();
var InlineDefault = /* #__PURE__ */ (function () {
    function InlineDefault() {

    };
    InlineDefault.value = new InlineDefault();
    return InlineDefault;
})();
var InlineNever = /* #__PURE__ */ (function () {
    function InlineNever() {

    };
    InlineNever.value = new InlineNever();
    return InlineNever;
})();
var InlineAlways = /* #__PURE__ */ (function () {
    function InlineAlways() {

    };
    InlineAlways.value = new InlineAlways();
    return InlineAlways;
})();
var InlineArity = /* #__PURE__ */ (function () {
    function InlineArity(value0) {
        this.value0 = value0;
    };
    InlineArity.create = function (value0) {
        return new InlineArity(value0);
    };
    return InlineArity;
})();
var InlineProp = /* #__PURE__ */ (function () {
    function InlineProp(value0) {
        this.value0 = value0;
    };
    InlineProp.create = function (value0) {
        return new InlineProp(value0);
    };
    return InlineProp;
})();
var InlineSpineProp = /* #__PURE__ */ (function () {
    function InlineSpineProp(value0) {
        this.value0 = value0;
    };
    InlineSpineProp.create = function (value0) {
        return new InlineSpineProp(value0);
    };
    return InlineSpineProp;
})();
var InlineRef = /* #__PURE__ */ (function () {
    function InlineRef() {

    };
    InlineRef.value = new InlineRef();
    return InlineRef;
})();
var EvalExtern = /* #__PURE__ */ (function () {
    function EvalExtern(value0) {
        this.value0 = value0;
    };
    EvalExtern.create = function (value0) {
        return new EvalExtern(value0);
    };
    return EvalExtern;
})();
var EvalLocal = /* #__PURE__ */ (function () {
    function EvalLocal(value0, value1) {
        this.value0 = value0;
        this.value1 = value1;
    };
    EvalLocal.create = function (value0) {
        return function (value1) {
            return new EvalLocal(value0, value1);
        };
    };
    return EvalLocal;
})();
var DistApp = /* #__PURE__ */ (function () {
    function DistApp(value0) {
        this.value0 = value0;
    };
    DistApp.create = function (value0) {
        return new DistApp(value0);
    };
    return DistApp;
})();
var DistUncurriedApp = /* #__PURE__ */ (function () {
    function DistUncurriedApp(value0) {
        this.value0 = value0;
    };
    DistUncurriedApp.create = function (value0) {
        return new DistUncurriedApp(value0);
    };
    return DistUncurriedApp;
})();
var DistAccessor = /* #__PURE__ */ (function () {
    function DistAccessor(value0) {
        this.value0 = value0;
    };
    DistAccessor.create = function (value0) {
        return new DistAccessor(value0);
    };
    return DistAccessor;
})();
var DistPrimOp1 = /* #__PURE__ */ (function () {
    function DistPrimOp1(value0) {
        this.value0 = value0;
    };
    DistPrimOp1.create = function (value0) {
        return new DistPrimOp1(value0);
    };
    return DistPrimOp1;
})();
var DistPrimOp2L = /* #__PURE__ */ (function () {
    function DistPrimOp2L(value0, value1) {
        this.value0 = value0;
        this.value1 = value1;
    };
    DistPrimOp2L.create = function (value0) {
        return function (value1) {
            return new DistPrimOp2L(value0, value1);
        };
    };
    return DistPrimOp2L;
})();
var DistPrimOp2R = /* #__PURE__ */ (function () {
    function DistPrimOp2R(value0, value1) {
        this.value0 = value0;
        this.value1 = value1;
    };
    DistPrimOp2R.create = function (value0) {
        return function (value1) {
            return new DistPrimOp2R(value0, value1);
        };
    };
    return DistPrimOp2R;
})();
var ExternExpr = /* #__PURE__ */ (function () {
    function ExternExpr(value0, value1) {
        this.value0 = value0;
        this.value1 = value1;
    };
    ExternExpr.create = function (value0) {
        return function (value1) {
            return new ExternExpr(value0, value1);
        };
    };
    return ExternExpr;
})();
var ExternDict = /* #__PURE__ */ (function () {
    function ExternDict(value0, value1) {
        this.value0 = value0;
        this.value1 = value1;
    };
    ExternDict.create = function (value0) {
        return function (value1) {
            return new ExternDict(value0, value1);
        };
    };
    return ExternDict;
})();
var ExternCtor = /* #__PURE__ */ (function () {
    function ExternCtor(value0, value1, value2, value3, value4) {
        this.value0 = value0;
        this.value1 = value1;
        this.value2 = value2;
        this.value3 = value3;
        this.value4 = value4;
    };
    ExternCtor.create = function (value0) {
        return function (value1) {
            return function (value2) {
                return function (value3) {
                    return function (value4) {
                        return new ExternCtor(value0, value1, value2, value3, value4);
                    };
                };
            };
        };
    };
    return ExternCtor;
})();
var SemTyped = /* #__PURE__ */ (function () {
    function SemTyped(value0, value1) {
        this.value0 = value0;
        this.value1 = value1;
    };
    SemTyped.create = function (value0) {
        return function (value1) {
            return new SemTyped(value0, value1);
        };
    };
    return SemTyped;
})();
var SemRef = /* #__PURE__ */ (function () {
    function SemRef(value0, value1, value2) {
        this.value0 = value0;
        this.value1 = value1;
        this.value2 = value2;
    };
    SemRef.create = function (value0) {
        return function (value1) {
            return function (value2) {
                return new SemRef(value0, value1, value2);
            };
        };
    };
    return SemRef;
})();
var SemLam = /* #__PURE__ */ (function () {
    function SemLam(value0, value1) {
        this.value0 = value0;
        this.value1 = value1;
    };
    SemLam.create = function (value0) {
        return function (value1) {
            return new SemLam(value0, value1);
        };
    };
    return SemLam;
})();
var SemMkFn = /* #__PURE__ */ (function () {
    function SemMkFn(value0) {
        this.value0 = value0;
    };
    SemMkFn.create = function (value0) {
        return new SemMkFn(value0);
    };
    return SemMkFn;
})();
var SemMkEffectFn = /* #__PURE__ */ (function () {
    function SemMkEffectFn(value0) {
        this.value0 = value0;
    };
    SemMkEffectFn.create = function (value0) {
        return new SemMkEffectFn(value0);
    };
    return SemMkEffectFn;
})();
var SemLet = /* #__PURE__ */ (function () {
    function SemLet(value0, value1, value2) {
        this.value0 = value0;
        this.value1 = value1;
        this.value2 = value2;
    };
    SemLet.create = function (value0) {
        return function (value1) {
            return function (value2) {
                return new SemLet(value0, value1, value2);
            };
        };
    };
    return SemLet;
})();
var SemLetRec = /* #__PURE__ */ (function () {
    function SemLetRec(value0, value1) {
        this.value0 = value0;
        this.value1 = value1;
    };
    SemLetRec.create = function (value0) {
        return function (value1) {
            return new SemLetRec(value0, value1);
        };
    };
    return SemLetRec;
})();
var SemEffectBind = /* #__PURE__ */ (function () {
    function SemEffectBind(value0, value1, value2) {
        this.value0 = value0;
        this.value1 = value1;
        this.value2 = value2;
    };
    SemEffectBind.create = function (value0) {
        return function (value1) {
            return function (value2) {
                return new SemEffectBind(value0, value1, value2);
            };
        };
    };
    return SemEffectBind;
})();
var SemEffectPure = /* #__PURE__ */ (function () {
    function SemEffectPure(value0) {
        this.value0 = value0;
    };
    SemEffectPure.create = function (value0) {
        return new SemEffectPure(value0);
    };
    return SemEffectPure;
})();
var SemEffectDefer = /* #__PURE__ */ (function () {
    function SemEffectDefer(value0) {
        this.value0 = value0;
    };
    SemEffectDefer.create = function (value0) {
        return new SemEffectDefer(value0);
    };
    return SemEffectDefer;
})();
var SemBranch = /* #__PURE__ */ (function () {
    function SemBranch(value0, value1) {
        this.value0 = value0;
        this.value1 = value1;
    };
    SemBranch.create = function (value0) {
        return function (value1) {
            return new SemBranch(value0, value1);
        };
    };
    return SemBranch;
})();
var SemAssocOp = /* #__PURE__ */ (function () {
    function SemAssocOp(value0, value1) {
        this.value0 = value0;
        this.value1 = value1;
    };
    SemAssocOp.create = function (value0) {
        return function (value1) {
            return new SemAssocOp(value0, value1);
        };
    };
    return SemAssocOp;
})();
var NeutLocal = /* #__PURE__ */ (function () {
    function NeutLocal(value0, value1) {
        this.value0 = value0;
        this.value1 = value1;
    };
    NeutLocal.create = function (value0) {
        return function (value1) {
            return new NeutLocal(value0, value1);
        };
    };
    return NeutLocal;
})();
var NeutVar = /* #__PURE__ */ (function () {
    function NeutVar(value0) {
        this.value0 = value0;
    };
    NeutVar.create = function (value0) {
        return new NeutVar(value0);
    };
    return NeutVar;
})();
var NeutStop = /* #__PURE__ */ (function () {
    function NeutStop(value0) {
        this.value0 = value0;
    };
    NeutStop.create = function (value0) {
        return new NeutStop(value0);
    };
    return NeutStop;
})();
var NeutData = /* #__PURE__ */ (function () {
    function NeutData(value0, value1, value2, value3, value4) {
        this.value0 = value0;
        this.value1 = value1;
        this.value2 = value2;
        this.value3 = value3;
        this.value4 = value4;
    };
    NeutData.create = function (value0) {
        return function (value1) {
            return function (value2) {
                return function (value3) {
                    return function (value4) {
                        return new NeutData(value0, value1, value2, value3, value4);
                    };
                };
            };
        };
    };
    return NeutData;
})();
var NeutCtorDef = /* #__PURE__ */ (function () {
    function NeutCtorDef(value0, value1, value2, value3, value4) {
        this.value0 = value0;
        this.value1 = value1;
        this.value2 = value2;
        this.value3 = value3;
        this.value4 = value4;
    };
    NeutCtorDef.create = function (value0) {
        return function (value1) {
            return function (value2) {
                return function (value3) {
                    return function (value4) {
                        return new NeutCtorDef(value0, value1, value2, value3, value4);
                    };
                };
            };
        };
    };
    return NeutCtorDef;
})();
var NeutApp = /* #__PURE__ */ (function () {
    function NeutApp(value0, value1) {
        this.value0 = value0;
        this.value1 = value1;
    };
    NeutApp.create = function (value0) {
        return function (value1) {
            return new NeutApp(value0, value1);
        };
    };
    return NeutApp;
})();
var NeutAccessor = /* #__PURE__ */ (function () {
    function NeutAccessor(value0, value1) {
        this.value0 = value0;
        this.value1 = value1;
    };
    NeutAccessor.create = function (value0) {
        return function (value1) {
            return new NeutAccessor(value0, value1);
        };
    };
    return NeutAccessor;
})();
var NeutUpdate = /* #__PURE__ */ (function () {
    function NeutUpdate(value0, value1) {
        this.value0 = value0;
        this.value1 = value1;
    };
    NeutUpdate.create = function (value0) {
        return function (value1) {
            return new NeutUpdate(value0, value1);
        };
    };
    return NeutUpdate;
})();
var NeutLit = /* #__PURE__ */ (function () {
    function NeutLit(value0) {
        this.value0 = value0;
    };
    NeutLit.create = function (value0) {
        return new NeutLit(value0);
    };
    return NeutLit;
})();
var NeutFail = /* #__PURE__ */ (function () {
    function NeutFail(value0) {
        this.value0 = value0;
    };
    NeutFail.create = function (value0) {
        return new NeutFail(value0);
    };
    return NeutFail;
})();
var NeutUncurriedApp = /* #__PURE__ */ (function () {
    function NeutUncurriedApp(value0, value1) {
        this.value0 = value0;
        this.value1 = value1;
    };
    NeutUncurriedApp.create = function (value0) {
        return function (value1) {
            return new NeutUncurriedApp(value0, value1);
        };
    };
    return NeutUncurriedApp;
})();
var NeutUncurriedEffectApp = /* #__PURE__ */ (function () {
    function NeutUncurriedEffectApp(value0, value1) {
        this.value0 = value0;
        this.value1 = value1;
    };
    NeutUncurriedEffectApp.create = function (value0) {
        return function (value1) {
            return new NeutUncurriedEffectApp(value0, value1);
        };
    };
    return NeutUncurriedEffectApp;
})();
var NeutPrimOp = /* #__PURE__ */ (function () {
    function NeutPrimOp(value0) {
        this.value0 = value0;
    };
    NeutPrimOp.create = function (value0) {
        return new NeutPrimOp(value0);
    };
    return NeutPrimOp;
})();
var NeutPrimEffect = /* #__PURE__ */ (function () {
    function NeutPrimEffect(value0) {
        this.value0 = value0;
    };
    NeutPrimEffect.create = function (value0) {
        return new NeutPrimEffect(value0);
    };
    return NeutPrimEffect;
})();
var NeutPrimUndefined = /* #__PURE__ */ (function () {
    function NeutPrimUndefined() {

    };
    NeutPrimUndefined.value = new NeutPrimUndefined();
    return NeutPrimUndefined;
})();
var ExternApp = /* #__PURE__ */ (function () {
    function ExternApp(value0) {
        this.value0 = value0;
    };
    ExternApp.create = function (value0) {
        return new ExternApp(value0);
    };
    return ExternApp;
})();
var ExternUncurriedApp = /* #__PURE__ */ (function () {
    function ExternUncurriedApp(value0) {
        this.value0 = value0;
    };
    ExternUncurriedApp.create = function (value0) {
        return new ExternUncurriedApp(value0);
    };
    return ExternUncurriedApp;
})();
var ExternAccessor = /* #__PURE__ */ (function () {
    function ExternAccessor(value0) {
        this.value0 = value0;
    };
    ExternAccessor.create = function (value0) {
        return new ExternAccessor(value0);
    };
    return ExternAccessor;
})();
var ExternPrimOp = /* #__PURE__ */ (function () {
    function ExternPrimOp(value0) {
        this.value0 = value0;
    };
    ExternPrimOp.create = function (value0) {
        return new ExternPrimOp(value0);
    };
    return ExternPrimOp;
})();
var Env = function (x) {
    return x;
};
var RewriteInline = /* #__PURE__ */ (function () {
    function RewriteInline(value0, value1, value2, value3) {
        this.value0 = value0;
        this.value1 = value1;
        this.value2 = value2;
        this.value3 = value3;
    };
    RewriteInline.create = function (value0) {
        return function (value1) {
            return function (value2) {
                return function (value3) {
                    return new RewriteInline(value0, value1, value2, value3);
                };
            };
        };
    };
    return RewriteInline;
})();
var RewriteUncurry = /* #__PURE__ */ (function () {
    function RewriteUncurry(value0, value1, value2, value3, value4) {
        this.value0 = value0;
        this.value1 = value1;
        this.value2 = value2;
        this.value3 = value3;
        this.value4 = value4;
    };
    RewriteUncurry.create = function (value0) {
        return function (value1) {
            return function (value2) {
                return function (value3) {
                    return function (value4) {
                        return new RewriteUncurry(value0, value1, value2, value3, value4);
                    };
                };
            };
        };
    };
    return RewriteUncurry;
})();
var RewriteStop = /* #__PURE__ */ (function () {
    function RewriteStop(value0) {
        this.value0 = value0;
    };
    RewriteStop.create = function (value0) {
        return new RewriteStop(value0);
    };
    return RewriteStop;
})();
var RewriteUnpackOp = /* #__PURE__ */ (function () {
    function RewriteUnpackOp(value0, value1, value2, value3) {
        this.value0 = value0;
        this.value1 = value1;
        this.value2 = value2;
        this.value3 = value3;
    };
    RewriteUnpackOp.create = function (value0) {
        return function (value1) {
            return function (value2) {
                return function (value3) {
                    return new RewriteUnpackOp(value0, value1, value2, value3);
                };
            };
        };
    };
    return RewriteUnpackOp;
})();
var RewriteDistBranchesLet = /* #__PURE__ */ (function () {
    function RewriteDistBranchesLet(value0, value1, value2, value3, value4) {
        this.value0 = value0;
        this.value1 = value1;
        this.value2 = value2;
        this.value3 = value3;
        this.value4 = value4;
    };
    RewriteDistBranchesLet.create = function (value0) {
        return function (value1) {
            return function (value2) {
                return function (value3) {
                    return function (value4) {
                        return new RewriteDistBranchesLet(value0, value1, value2, value3, value4);
                    };
                };
            };
        };
    };
    return RewriteDistBranchesLet;
})();
var RewriteDistBranchesOp = /* #__PURE__ */ (function () {
    function RewriteDistBranchesOp(value0, value1, value2) {
        this.value0 = value0;
        this.value1 = value1;
        this.value2 = value2;
    };
    RewriteDistBranchesOp.create = function (value0) {
        return function (value1) {
            return function (value2) {
                return new RewriteDistBranchesOp(value0, value1, value2);
            };
        };
    };
    return RewriteDistBranchesOp;
})();
var ExprSyntax = /* #__PURE__ */ (function () {
    function ExprSyntax(value0, value1) {
        this.value0 = value0;
        this.value1 = value1;
    };
    ExprSyntax.create = function (value0) {
        return function (value1) {
            return new ExprSyntax(value0, value1);
        };
    };
    return ExprSyntax;
})();
var ExprRewrite = /* #__PURE__ */ (function () {
    function ExprRewrite(value0, value1) {
        this.value0 = value0;
        this.value1 = value1;
    };
    ExprRewrite.create = function (value0) {
        return function (value1) {
            return new ExprRewrite(value0, value1);
        };
    };
    return ExprRewrite;
})();
var Ctx = function (x) {
    return x;
};
var newtypeNeutralExpr_ = {
    Coercible0: function () {
        return undefined;
    }
};
var newtypeEnv_ = {
    Coercible0: function () {
        return undefined;
    }
};
var hasSyntaxBackendExpr = {
    syntaxOf: /* #__PURE__ */ (function () {
        var go = function ($copy_v) {
            var $tco_done = false;
            var $tco_result;
            function $tco_loop(v) {
                if (v instanceof ExprSyntax && v.value1 instanceof PureScript_Backend_Optimizer_Syntax.Typed) {
                    $copy_v = v.value1.value1;
                    return;
                };
                if (v instanceof ExprSyntax) {
                    $tco_done = true;
                    return new Data_Maybe.Just(v.value1);
                };
                $tco_done = true;
                return Data_Maybe.Nothing.value;
            };
            while (!$tco_done) {
                $tco_result = $tco_loop($copy_v);
            };
            return $tco_result;
        };
        return go;
    })()
};
var syntaxOf = /* #__PURE__ */ PureScript_Backend_Optimizer_Syntax.syntaxOf(hasSyntaxBackendExpr);
var hasAnalysisBackendExpr = {
    analysisOf: function (v) {
        if (v instanceof ExprSyntax) {
            return v.value0;
        };
        if (v instanceof ExprRewrite) {
            return v.value0;
        };
        throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Semantics (line 139, column 16 - line 141, column 25): " + [ v.constructor.name ]);
    }
};
var analysisOf = /* #__PURE__ */ PureScript_Backend_Optimizer_Analysis.analysisOf(hasAnalysisBackendExpr);
var eqUnpackOp = function (dictEq) {
    var eqArray = Data_Eq.eqArray(PureScript_Backend_Optimizer_CoreFn.eqProp(dictEq));
    var eqArray1 = Data_Eq.eqArray(dictEq);
    var eqArray2 = Data_Eq.eqArray(eqTuple(dictEq));
    return {
        eq: function (x) {
            return function (y) {
                if (x instanceof UnpackRecord && y instanceof UnpackRecord) {
                    return Data_Eq.eq(eqArray)(x.value0)(y.value0);
                };
                if (x instanceof UnpackUpdate && y instanceof UnpackUpdate) {
                    return Data_Eq.eq(dictEq)(x.value0)(y.value0) && Data_Eq.eq(eqArray)(x.value1)(y.value1);
                };
                if (x instanceof UnpackArray && y instanceof UnpackArray) {
                    return Data_Eq.eq(eqArray1)(x.value0)(y.value0);
                };
                if (x instanceof UnpackData && y instanceof UnpackData) {
                    return Data_Eq.eq(eqQualified)(x.value0)(y.value0) && Data_Eq.eq(PureScript_Backend_Optimizer_CoreFn.eqConstructorType)(x.value1)(y.value1) && Data_Eq.eq(PureScript_Backend_Optimizer_CoreFn.eqProperName)(x.value2)(y.value2) && Data_Eq.eq(PureScript_Backend_Optimizer_CoreFn.eqIdent)(x.value3)(y.value3) && Data_Eq.eq(eqArray2)(x.value4)(y.value4);
                };
                return false;
            };
        }
    };
};
var eqInlineAccessor = {
    eq: function (x) {
        return function (y) {
            if (x instanceof InlineProp && y instanceof InlineProp) {
                return x.value0 === y.value0;
            };
            if (x instanceof InlineSpineProp && y instanceof InlineSpineProp) {
                return x.value0 === y.value0;
            };
            if (x instanceof InlineRef && y instanceof InlineRef) {
                return true;
            };
            return false;
        };
    }
};
var ordInlineAccessor = {
    compare: function (x) {
        return function (y) {
            if (x instanceof InlineProp && y instanceof InlineProp) {
                return Data_Ord.compare(Data_Ord.ordString)(x.value0)(y.value0);
            };
            if (x instanceof InlineProp) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof InlineProp) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof InlineSpineProp && y instanceof InlineSpineProp) {
                return Data_Ord.compare(Data_Ord.ordString)(x.value0)(y.value0);
            };
            if (x instanceof InlineSpineProp) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof InlineSpineProp) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof InlineRef && y instanceof InlineRef) {
                return Data_Ordering.EQ.value;
            };
            throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Semantics (line 0, column 0 - line 0, column 0): " + [ x.constructor.name, y.constructor.name ]);
        };
    },
    Eq0: function () {
        return eqInlineAccessor;
    }
};
var eqEvalRef = {
    eq: function (x) {
        return function (y) {
            if (x instanceof EvalExtern && y instanceof EvalExtern) {
                return Data_Eq.eq(eqQualified)(x.value0)(y.value0);
            };
            if (x instanceof EvalLocal && y instanceof EvalLocal) {
                return Data_Eq.eq(eqMaybe)(x.value0)(y.value0) && Data_Eq.eq(PureScript_Backend_Optimizer_Syntax.eqLevel)(x.value1)(y.value1);
            };
            return false;
        };
    }
};
var ordEvalRef = {
    compare: function (x) {
        return function (y) {
            if (x instanceof EvalExtern && y instanceof EvalExtern) {
                return Data_Ord.compare(ordQualified)(x.value0)(y.value0);
            };
            if (x instanceof EvalExtern) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof EvalExtern) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof EvalLocal && y instanceof EvalLocal) {
                var v = Data_Ord.compare(ordMaybe)(x.value0)(y.value0);
                if (v instanceof Data_Ordering.LT) {
                    return Data_Ordering.LT.value;
                };
                if (v instanceof Data_Ordering.GT) {
                    return Data_Ordering.GT.value;
                };
                return Data_Ord.compare(PureScript_Backend_Optimizer_Syntax.ordLevel)(x.value1)(y.value1);
            };
            throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Semantics (line 0, column 0 - line 0, column 0): " + [ x.constructor.name, y.constructor.name ]);
        };
    },
    Eq0: function () {
        return eqEvalRef;
    }
};
var eqDistOp = function (dictEq) {
    var eqNonEmptyArray1 = Data_Array_NonEmpty_Internal.eqNonEmptyArray(dictEq);
    var eqArray = Data_Eq.eqArray(dictEq);
    return {
        eq: function (x) {
            return function (y) {
                if (x instanceof DistApp && y instanceof DistApp) {
                    return Data_Eq.eq(eqNonEmptyArray1)(x.value0)(y.value0);
                };
                if (x instanceof DistUncurriedApp && y instanceof DistUncurriedApp) {
                    return Data_Eq.eq(eqArray)(x.value0)(y.value0);
                };
                if (x instanceof DistAccessor && y instanceof DistAccessor) {
                    return Data_Eq.eq(PureScript_Backend_Optimizer_Syntax.eqBackendAccessor)(x.value0)(y.value0);
                };
                if (x instanceof DistPrimOp1 && y instanceof DistPrimOp1) {
                    return Data_Eq.eq(PureScript_Backend_Optimizer_Syntax.eqBackendOperator1)(x.value0)(y.value0);
                };
                if (x instanceof DistPrimOp2L && y instanceof DistPrimOp2L) {
                    return Data_Eq.eq(PureScript_Backend_Optimizer_Syntax.eqBackendOperator2)(x.value0)(y.value0) && Data_Eq.eq(dictEq)(x.value1)(y.value1);
                };
                if (x instanceof DistPrimOp2R && y instanceof DistPrimOp2R) {
                    return Data_Eq.eq(dictEq)(x.value0)(y.value0) && Data_Eq.eq(PureScript_Backend_Optimizer_Syntax.eqBackendOperator2)(x.value1)(y.value1);
                };
                return false;
            };
        }
    };
};
var eqBackendRewrite = function (dictEq) {
    var eqUnpackOp1 = eqUnpackOp(dictEq);
    var eqNonEmptyArray1 = Data_Array_NonEmpty_Internal.eqNonEmptyArray(PureScript_Backend_Optimizer_Syntax.eqPair(dictEq));
    var eqDistOp1 = eqDistOp(dictEq);
    return {
        eq: function (x) {
            return function (y) {
                if (x instanceof RewriteInline && y instanceof RewriteInline) {
                    return Data_Eq.eq(eqMaybe)(x.value0)(y.value0) && Data_Eq.eq(PureScript_Backend_Optimizer_Syntax.eqLevel)(x.value1)(y.value1) && Data_Eq.eq(dictEq)(x.value2)(y.value2) && Data_Eq.eq(dictEq)(x.value3)(y.value3);
                };
                if (x instanceof RewriteUncurry && y instanceof RewriteUncurry) {
                    return Data_Eq.eq(eqMaybe)(x.value0)(y.value0) && Data_Eq.eq(PureScript_Backend_Optimizer_Syntax.eqLevel)(x.value1)(y.value1) && Data_Eq.eq(eqNonEmptyArray)(x.value2)(y.value2) && Data_Eq.eq(dictEq)(x.value3)(y.value3) && Data_Eq.eq(dictEq)(x.value4)(y.value4);
                };
                if (x instanceof RewriteStop && y instanceof RewriteStop) {
                    return Data_Eq.eq(eqQualified)(x.value0)(y.value0);
                };
                if (x instanceof RewriteUnpackOp && y instanceof RewriteUnpackOp) {
                    return Data_Eq.eq(eqMaybe)(x.value0)(y.value0) && Data_Eq.eq(PureScript_Backend_Optimizer_Syntax.eqLevel)(x.value1)(y.value1) && Data_Eq.eq(eqUnpackOp1)(x.value2)(y.value2) && Data_Eq.eq(dictEq)(x.value3)(y.value3);
                };
                if (x instanceof RewriteDistBranchesLet && y instanceof RewriteDistBranchesLet) {
                    return Data_Eq.eq(eqMaybe)(x.value0)(y.value0) && Data_Eq.eq(PureScript_Backend_Optimizer_Syntax.eqLevel)(x.value1)(y.value1) && Data_Eq.eq(eqNonEmptyArray1)(x.value2)(y.value2) && Data_Eq.eq(dictEq)(x.value3)(y.value3) && Data_Eq.eq(dictEq)(x.value4)(y.value4);
                };
                if (x instanceof RewriteDistBranchesOp && y instanceof RewriteDistBranchesOp) {
                    return Data_Eq.eq(eqNonEmptyArray1)(x.value0)(y.value0) && Data_Eq.eq(dictEq)(x.value1)(y.value1) && Data_Eq.eq(eqDistOp1)(x.value2)(y.value2);
                };
                return false;
            };
        }
    };
};
var eqBackendExpr = {
    eq: function (v) {
        return function (v1) {
            if (v instanceof ExprSyntax && v1 instanceof ExprSyntax) {
                return v.value0.size === v1.value0.size && Data_Eq.eq(PureScript_Backend_Optimizer_Syntax.eqBackendSyntax(eqBackendExpr))(v.value1)(v1.value1);
            };
            if (v instanceof ExprRewrite && v1 instanceof ExprRewrite) {
                return v.value0.size === v1.value0.size && Data_Eq.eq(eqBackendRewrite(eqBackendExpr))(v.value1)(v1.value1);
            };
            return false;
        };
    }
};
var snocApp = function (prev) {
    return function (next) {
        var v = Data_Array.last(prev);
        if (v instanceof Data_Maybe.Just && v.value0 instanceof ExternApp) {
            return Data_Array.snoc(Data_Array.dropEnd(1)(prev))(new ExternApp(Data_Array.snoc(v.value0.value0)(next)));
        };
        return Data_Array.snoc(prev)(new ExternApp([ next ]));
    };
};
var snocSpine = function (spine) {
    return function (v) {
        if (v instanceof ExternApp) {
            return Data_Foldable.foldl(Data_Foldable.foldableArray)(snocApp)(spine)(v.value0);
        };
        return Data_Array.snoc(spine)(v);
    };
};
var simplifyCondIsTag = function (v) {
    return function (v1) {
        return function (v2) {
            if (v1.value0 instanceof ExprSyntax && (v1.value0.value1 instanceof PureScript_Backend_Optimizer_Syntax.PrimOp && (v1.value0.value1.value0 instanceof PureScript_Backend_Optimizer_Syntax.Op1 && (v1.value0.value1.value0.value0 instanceof PureScript_Backend_Optimizer_Syntax.OpIsTag && (v1.value1 instanceof ExprSyntax && (v1.value1.value1 instanceof PureScript_Backend_Optimizer_Syntax.Lit && (v1.value1.value1.value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitBoolean && (!v1.value1.value1.value0.value0 && (v2 instanceof ExprSyntax && (v2.value1 instanceof PureScript_Backend_Optimizer_Syntax.PrimOp && (v2.value1.value0 instanceof PureScript_Backend_Optimizer_Syntax.Op1 && (v2.value1.value0.value0 instanceof PureScript_Backend_Optimizer_Syntax.OpIsTag && Data_Eq.eq(eqBackendExpr)(v1.value0.value1.value0.value1)(v2.value1.value0.value1))))))))))))) {
                return new Data_Maybe.Just(v2);
            };
            return Data_Maybe.Nothing.value;
        };
    };
};
var shouldUnpackUpdate = function (ident) {
    return function (level) {
        return function (binding) {
            return function (body) {
                var v = PureScript_Backend_Optimizer_Analysis.analysisOf(hasAnalysisBackendExpr)(body);
                var v1 = function (v2) {
                    return Data_Maybe.Nothing.value;
                };
                if (binding instanceof ExprSyntax && binding.value1 instanceof PureScript_Backend_Optimizer_Syntax.Update) {
                    var $656 = Data_Map_Internal.lookup(PureScript_Backend_Optimizer_Syntax.ordLevel)(level)(v.usages);
                    if ($656 instanceof Data_Maybe.Just) {
                        var $657 = $656.value0.total === ($656.value0.access + $656.value0.update | 0);
                        if ($657) {
                            var analysis = PureScript_Backend_Optimizer_Analysis.updated(level)(Data_Semigroup.append(PureScript_Backend_Optimizer_Analysis.semigroupBackendAnalysis)(PureScript_Backend_Optimizer_Analysis.analysisOf(hasAnalysisBackendExpr)(binding.value1.value0))(Data_Foldable.foldr(Data_Foldable.foldableArray)(function ($2086) {
                                return append(analysisOf(PureScript_Backend_Optimizer_CoreFn.propValue($2086)));
                            })(PureScript_Backend_Optimizer_Analysis.complex(PureScript_Backend_Optimizer_Analysis.NonTrivial.value)(PureScript_Backend_Optimizer_Analysis.bound(level)(v)))(binding.value1.value1)));
                            return new Data_Maybe.Just(new ExprRewrite(PureScript_Backend_Optimizer_Analysis.withRewrite(analysis), new RewriteUnpackOp(ident, level, new UnpackUpdate(binding.value1.value0, binding.value1.value1), body)));
                        };
                        return v1(true);
                    };
                    return v1(true);
                };
                return v1(true);
            };
        };
    };
};
var shouldUnpackRecord = function (ident) {
    return function (level) {
        return function (binding) {
            return function (body) {
                var v = PureScript_Backend_Optimizer_Analysis.analysisOf(hasAnalysisBackendExpr)(body);
                var v1 = function (v2) {
                    return Data_Maybe.Nothing.value;
                };
                if (binding instanceof ExprSyntax && (binding.value1 instanceof PureScript_Backend_Optimizer_Syntax.Lit && binding.value1.value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitRecord)) {
                    var $665 = Data_Map_Internal.lookup(PureScript_Backend_Optimizer_Syntax.ordLevel)(level)(v.usages);
                    if ($665 instanceof Data_Maybe.Just) {
                        var $666 = $665.value0.total === ($665.value0.access + $665.value0.update | 0);
                        if ($666) {
                            var analysis = Data_Foldable.foldr(Data_Foldable.foldableArray)(function ($2087) {
                                return append(analysisOf(PureScript_Backend_Optimizer_CoreFn.propValue($2087)));
                            })(PureScript_Backend_Optimizer_Analysis.complex(PureScript_Backend_Optimizer_Analysis.NonTrivial.value)(PureScript_Backend_Optimizer_Analysis.bound(level)(v)))(binding.value1.value0.value0);
                            return new Data_Maybe.Just(new ExprRewrite(PureScript_Backend_Optimizer_Analysis.withRewrite(analysis), new RewriteUnpackOp(ident, level, new UnpackRecord(binding.value1.value0.value0), body)));
                        };
                        return v1(true);
                    };
                    return v1(true);
                };
                return v1(true);
            };
        };
    };
};
var shouldUnpackCtor = function (ident) {
    return function (level) {
        return function (a) {
            return function (body) {
                var v = PureScript_Backend_Optimizer_Analysis.analysisOf(hasAnalysisBackendExpr)(body);
                var v1 = function (v2) {
                    return Data_Maybe.Nothing.value;
                };
                if (a instanceof ExprSyntax && a.value1 instanceof PureScript_Backend_Optimizer_Syntax.CtorSaturated) {
                    var $674 = Data_Map_Internal.lookup(PureScript_Backend_Optimizer_Syntax.ordLevel)(level)(v.usages);
                    if ($674 instanceof Data_Maybe.Just) {
                        var $675 = $674.value0.total === ($674.value0.access + $674["value0"]["case"] | 0);
                        if ($675) {
                            var analysis = Data_Foldable.foldr(Data_Foldable.foldableArray)(function ($2088) {
                                return append(analysisOf(Data_Tuple.snd($2088)));
                            })(PureScript_Backend_Optimizer_Analysis.complex(PureScript_Backend_Optimizer_Analysis.NonTrivial.value)(PureScript_Backend_Optimizer_Analysis.bound(level)(v)))(a.value1.value4);
                            return new Data_Maybe.Just(new ExprRewrite(PureScript_Backend_Optimizer_Analysis.withRewrite(analysis), new RewriteUnpackOp(ident, level, new UnpackData(a.value1.value0, a.value1.value1, a.value1.value2, a.value1.value3, a.value1.value4), body)));
                        };
                        return v1(true);
                    };
                    return v1(true);
                };
                return v1(true);
            };
        };
    };
};
var shouldUnpackArray = function (ident) {
    return function (level) {
        return function (binding) {
            return function (body) {
                var v = PureScript_Backend_Optimizer_Analysis.analysisOf(hasAnalysisBackendExpr)(body);
                var v1 = function (v2) {
                    return Data_Maybe.Nothing.value;
                };
                if (binding instanceof ExprSyntax && (binding.value1 instanceof PureScript_Backend_Optimizer_Syntax.Lit && binding.value1.value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitArray)) {
                    var $686 = Data_Map_Internal.lookup(PureScript_Backend_Optimizer_Syntax.ordLevel)(level)(v.usages);
                    if ($686 instanceof Data_Maybe.Just) {
                        var $687 = $686.value0.total === $686.value0.access;
                        if ($687) {
                            var analysis = Data_Foldable.foldr(Data_Foldable.foldableArray)(function ($2089) {
                                return append(analysisOf($2089));
                            })(PureScript_Backend_Optimizer_Analysis.complex(PureScript_Backend_Optimizer_Analysis.NonTrivial.value)(PureScript_Backend_Optimizer_Analysis.bound(level)(v)))(binding.value1.value0.value0);
                            return new Data_Maybe.Just(new ExprRewrite(PureScript_Backend_Optimizer_Analysis.withRewrite(analysis), new RewriteUnpackOp(ident, level, new UnpackArray(binding.value1.value0.value0), body)));
                        };
                        return v1(true);
                    };
                    return v1(true);
                };
                return v1(true);
            };
        };
    };
};
var shouldUncurryAbs = function (ident) {
    return function (level) {
        return function (a) {
            return function (b) {
                var v = PureScript_Backend_Optimizer_Analysis.analysisOf(hasAnalysisBackendExpr)(b);
                var v1 = function (v2) {
                    return Data_Maybe.Nothing.value;
                };
                if (a instanceof ExprSyntax && a.value1 instanceof PureScript_Backend_Optimizer_Syntax.Abs) {
                    var $695 = Data_Map_Internal.lookup(PureScript_Backend_Optimizer_Syntax.ordLevel)(level)(v.usages);
                    if ($695 instanceof Data_Maybe.Just) {
                        var $696 = Data_Set.toUnfoldable(Data_Unfoldable.unfoldableArray)($695.value0.arities);
                        if ($696.length === 1) {
                            var $697 = $696[0] === Data_Array_NonEmpty.length(a.value1.value0);
                            if ($697) {
                                var analysis = PureScript_Backend_Optimizer_Analysis.withResult(PureScript_Backend_Optimizer_Analysis.resultOf(hasAnalysisBackendExpr)(b))(PureScript_Backend_Optimizer_Analysis.bump(PureScript_Backend_Optimizer_Analysis.complex(PureScript_Backend_Optimizer_Analysis.NonTrivial.value)(Data_Semigroup.append(PureScript_Backend_Optimizer_Analysis.semigroupBackendAnalysis)(PureScript_Backend_Optimizer_Analysis.analysisOf(hasAnalysisBackendExpr)(a))(PureScript_Backend_Optimizer_Analysis.bound(level)(PureScript_Backend_Optimizer_Analysis.analysisOf(hasAnalysisBackendExpr)(b))))));
                                return new Data_Maybe.Just(new ExprRewrite(PureScript_Backend_Optimizer_Analysis.withRewrite(analysis), new RewriteUncurry(ident, level, a.value1.value0, a.value1.value1, b)));
                            };
                            return v1(true);
                        };
                        return v1(true);
                    };
                    return v1(true);
                };
                return v1(true);
            };
        };
    };
};
var shouldInlineExternReference = function (v) {
    return function (v1) {
        return function (v2) {
            return Data_Ord.lessThanOrEq(PureScript_Backend_Optimizer_Analysis.ordComplexity)(v1.complexity)(PureScript_Backend_Optimizer_Analysis.Deref.value) && v1.size < 16;
        };
    };
};
var shouldInlineExternLiteral = function (v) {
    if (v instanceof PureScript_Backend_Optimizer_CoreFn.LitInt) {
        return true;
    };
    if (v instanceof PureScript_Backend_Optimizer_CoreFn.LitNumber) {
        return true;
    };
    if (v instanceof PureScript_Backend_Optimizer_CoreFn.LitString) {
        return Data_String_CodePoints.length(v.value0) <= 32;
    };
    if (v instanceof PureScript_Backend_Optimizer_CoreFn.LitChar) {
        return true;
    };
    if (v instanceof PureScript_Backend_Optimizer_CoreFn.LitBoolean) {
        return true;
    };
    if (v instanceof PureScript_Backend_Optimizer_CoreFn.LitArray) {
        return Data_Array["null"](v.value0);
    };
    if (v instanceof PureScript_Backend_Optimizer_CoreFn.LitRecord) {
        return Data_Array["null"](v.value0);
    };
    throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Semantics (line 1673, column 29 - line 1680, column 30): " + [ v.constructor.name ]);
};
var shouldInlineExternAppArg = function (v) {
    return function (v1) {
        if (v1 instanceof SemLam) {
            return Data_Ord.lessThanOrEq(PureScript_Backend_Optimizer_Analysis.ordCapture)(v.captured)(PureScript_Backend_Optimizer_Analysis.CaptureBranch.value) && (v.total > 0 && v.call === v.total);
        };
        return false;
    };
};
var shouldInlineExternApp = function (v) {
    return function (v1) {
        return function (v2) {
            return function (args) {
                var delayed = Data_Array.length(v1.args) > 0;
                return Data_Ord.lessThanOrEq(PureScript_Backend_Optimizer_Analysis.ordComplexity)(v1.complexity)(PureScript_Backend_Optimizer_Analysis.Deref.value) && v1.size < 16 || (Data_Map_Internal.isEmpty(v1.usages) && (!v1.externs && v1.size < 64) || (delayed && (Data_Array.length(v1.args) <= Data_Array.length(args) && v1.size < 16) || delayed && (Data_Foldable.or(Data_Foldable.foldableArray)(Data_HeytingAlgebra.heytingAlgebraBoolean)(Data_Array.zipWith(shouldInlineExternAppArg)(v1.args)(args)) && v1.size < 16)));
            };
        };
    };
};
var shouldInlineExternAccessor = function (v) {
    return function (v1) {
        return function (v2) {
            return function (v3) {
                return Data_Ord.lessThanOrEq(PureScript_Backend_Optimizer_Analysis.ordComplexity)(v1.complexity)(PureScript_Backend_Optimizer_Analysis.Deref.value) && v1.size < 16;
            };
        };
    };
};
var shouldEtaReduce = function (level1) {
    return function (binding) {
        var isSameArg = function (v) {
            return function (v1) {
                if (v1 instanceof ExprSyntax && v1.value1 instanceof PureScript_Backend_Optimizer_Syntax.Local) {
                    return Data_Eq.eq(PureScript_Backend_Optimizer_Syntax.eqLevel)(v.value1)(v1.value1.value1);
                };
                return false;
            };
        };
        return function (v) {
            var v1 = function (v2) {
                return Data_Maybe.Nothing.value;
            };
            if (v instanceof ExprSyntax && (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.Abs && (v.value1.value1 instanceof ExprSyntax && (v.value1.value1.value1 instanceof PureScript_Backend_Optimizer_Syntax.App && (v.value1.value1.value1.value0 instanceof ExprSyntax && v.value1.value1.value1.value0.value1 instanceof PureScript_Backend_Optimizer_Syntax.Local))))) {
                var $736 = Data_Eq.eq(PureScript_Backend_Optimizer_Syntax.eqLevel)(level1)(v.value1.value1.value1.value0.value1.value1);
                if ($736) {
                    var $737 = Data_Array_NonEmpty.length(v.value1.value0) === Data_Array_NonEmpty.length(v.value1.value1.value1.value1);
                    if ($737) {
                        var $738 = and(Data_Array_NonEmpty.zipWith(isSameArg)(v.value1.value0)(v.value1.value1.value1.value1));
                        if ($738) {
                            return new Data_Maybe.Just(binding);
                        };
                        return v1(true);
                    };
                    return v1(true);
                };
                return v1(true);
            };
            return v1(true);
        };
    };
};
var shouldDistributeBranches = function (ident) {
    return function (level) {
        return function (a) {
            return function (body) {
                var v = PureScript_Backend_Optimizer_Analysis.analysisOf(hasAnalysisBackendExpr)(body);
                var v1 = function (v2) {
                    return Data_Maybe.Nothing.value;
                };
                if (a instanceof ExprSyntax && a.value1 instanceof PureScript_Backend_Optimizer_Syntax.Branch) {
                    var $753 = v.size <= 128;
                    if ($753) {
                        var $754 = Data_Eq.eq(PureScript_Backend_Optimizer_Analysis.eqResultTerm)(a.value0.result)(PureScript_Backend_Optimizer_Analysis.KnownNeutral.value);
                        if ($754) {
                            var $755 = Data_Map_Internal.lookup(PureScript_Backend_Optimizer_Syntax.ordLevel)(level)(v.usages);
                            if ($755 instanceof Data_Maybe.Just) {
                                var $756 = $755.value0.total === ($755.value0.access + $755["value0"]["case"] | 0);
                                if ($756) {
                                    var analysis = Data_Semigroup.append(PureScript_Backend_Optimizer_Analysis.semigroupBackendAnalysis)(PureScript_Backend_Optimizer_Analysis.analysisOf(hasAnalysisBackendExpr)(a))(PureScript_Backend_Optimizer_Analysis.bound(level)(PureScript_Backend_Optimizer_Analysis.analysisOf(hasAnalysisBackendExpr)(body)));
                                    return new Data_Maybe.Just(new ExprRewrite(PureScript_Backend_Optimizer_Analysis.withRewrite(analysis), new RewriteDistBranchesLet(ident, level, a.value1.value0, a.value1.value1, body)));
                                };
                                return v1(true);
                            };
                            return v1(true);
                        };
                        return v1(true);
                    };
                    return v1(true);
                };
                return v1(true);
            };
        };
    };
};
var shouldDistributeBranchUncurriedApps = function (analysis1) {
    return function (branches) {
        return function (def) {
            return function (spine) {
                var $762 = Data_Array.all(function ($2090) {
                    return (function (v) {
                        return Data_Ord.lessThanOrEq(PureScript_Backend_Optimizer_Analysis.ordComplexity)(v)(PureScript_Backend_Optimizer_Analysis.Deref.value);
                    })((function (v) {
                        return v.complexity;
                    })(unwrap(analysisOf($2090))));
                })(spine);
                if ($762) {
                    var analysis = Data_Semigroup.append(PureScript_Backend_Optimizer_Analysis.semigroupBackendAnalysis)(analysis1)(Data_Foldable.foldMap(Data_Foldable.foldableArray)(PureScript_Backend_Optimizer_Analysis.monoidBackendAnalysis)(analysisOf)(spine));
                    return new Data_Maybe.Just(new ExprRewrite(PureScript_Backend_Optimizer_Analysis.withRewrite(analysis), new RewriteDistBranchesOp(branches, def, new DistUncurriedApp(spine))));
                };
                return Data_Maybe.Nothing.value;
            };
        };
    };
};
var shouldDistributeBranchPrimOp2R = function (analysis1) {
    return function (branches) {
        return function (def) {
            return function (lhs) {
                return function (op2) {
                    var $763 = Data_Ord.lessThanOrEq(PureScript_Backend_Optimizer_Analysis.ordComplexity)((Data_Newtype.unwrap()(PureScript_Backend_Optimizer_Analysis.analysisOf(hasAnalysisBackendExpr)(lhs))).complexity)(PureScript_Backend_Optimizer_Analysis.Deref.value);
                    if ($763) {
                        var analysis = PureScript_Backend_Optimizer_Analysis.bump(Data_Semigroup.append(PureScript_Backend_Optimizer_Analysis.semigroupBackendAnalysis)(analysis1)(PureScript_Backend_Optimizer_Analysis.analysisOf(hasAnalysisBackendExpr)(lhs)));
                        return new Data_Maybe.Just(new ExprRewrite(PureScript_Backend_Optimizer_Analysis.withRewrite(analysis), new RewriteDistBranchesOp(branches, def, new DistPrimOp2R(lhs, op2))));
                    };
                    return Data_Maybe.Nothing.value;
                };
            };
        };
    };
};
var shouldDistributeBranchPrimOp2L = function (analysis1) {
    return function (branches) {
        return function (def) {
            return function (op2) {
                return function (rhs) {
                    var $764 = Data_Ord.lessThanOrEq(PureScript_Backend_Optimizer_Analysis.ordComplexity)((Data_Newtype.unwrap()(PureScript_Backend_Optimizer_Analysis.analysisOf(hasAnalysisBackendExpr)(rhs))).complexity)(PureScript_Backend_Optimizer_Analysis.Deref.value);
                    if ($764) {
                        var analysis = PureScript_Backend_Optimizer_Analysis.bump(Data_Semigroup.append(PureScript_Backend_Optimizer_Analysis.semigroupBackendAnalysis)(analysis1)(PureScript_Backend_Optimizer_Analysis.analysisOf(hasAnalysisBackendExpr)(rhs)));
                        return new Data_Maybe.Just(new ExprRewrite(PureScript_Backend_Optimizer_Analysis.withRewrite(analysis), new RewriteDistBranchesOp(branches, def, new DistPrimOp2L(op2, rhs))));
                    };
                    return Data_Maybe.Nothing.value;
                };
            };
        };
    };
};
var shouldDistributeBranchPrimOp1 = function (analysis1) {
    return function (branches) {
        return function (def) {
            return function (op) {
                var analysis = PureScript_Backend_Optimizer_Analysis.bump(analysis1);
                return new Data_Maybe.Just(new ExprRewrite(PureScript_Backend_Optimizer_Analysis.withRewrite(analysis), new RewriteDistBranchesOp(branches, def, new DistPrimOp1(op))));
            };
        };
    };
};
var shouldDistributeBranchApps = function (analysis1) {
    return function (branches) {
        return function (def) {
            return function (spine) {
                var $765 = Data_Array_NonEmpty.all(function ($2091) {
                    return (function (v) {
                        return Data_Ord.lessThanOrEq(PureScript_Backend_Optimizer_Analysis.ordComplexity)(v)(PureScript_Backend_Optimizer_Analysis.Deref.value);
                    })((function (v) {
                        return v.complexity;
                    })(unwrap(analysisOf($2091))));
                })(spine);
                if ($765) {
                    var analysis = Data_Semigroup.append(PureScript_Backend_Optimizer_Analysis.semigroupBackendAnalysis)(analysis1)(Data_Foldable.foldMap(Data_Array_NonEmpty_Internal.foldableNonEmptyArray)(PureScript_Backend_Optimizer_Analysis.monoidBackendAnalysis)(analysisOf)(spine));
                    return new Data_Maybe.Just(new ExprRewrite(PureScript_Backend_Optimizer_Analysis.withRewrite(analysis), new RewriteDistBranchesOp(branches, def, new DistApp(spine))));
                };
                return Data_Maybe.Nothing.value;
            };
        };
    };
};
var shouldDistributeBranchAccessor = function (analysis1) {
    return function (branches) {
        return function (def) {
            return function (acc) {
                var analysis = PureScript_Backend_Optimizer_Analysis.bump(analysis1);
                return new Data_Maybe.Just(new ExprRewrite(PureScript_Backend_Optimizer_Analysis.withRewrite(analysis), new RewriteDistBranchesOp(branches, def, new DistAccessor(acc))));
            };
        };
    };
};
var rewriteInline = function (ident) {
    return function (level) {
        return function (binding) {
            return function (body) {
                var s2 = PureScript_Backend_Optimizer_Analysis.analysisOf(hasAnalysisBackendExpr)(body);
                var powAnalysis = (function () {
                    var v = Data_Map_Internal.lookup(PureScript_Backend_Optimizer_Syntax.ordLevel)(level)((Data_Newtype.unwrap()(s2)).usages);
                    if (v instanceof Data_Maybe.Just) {
                        return Data_Semigroup.append(PureScript_Backend_Optimizer_Analysis.semigroupBackendAnalysis)(s2)(Data_Monoid.power(PureScript_Backend_Optimizer_Analysis.monoidBackendAnalysis)(PureScript_Backend_Optimizer_Analysis.analysisOf(hasAnalysisBackendExpr)(binding))(v.value0.total));
                    };
                    if (v instanceof Data_Maybe.Nothing) {
                        return s2;
                    };
                    throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Semantics (line 1449, column 19 - line 1454, column 11): " + [ v.constructor.name ]);
                })();
                return new ExprRewrite(PureScript_Backend_Optimizer_Analysis.withRewrite(PureScript_Backend_Optimizer_Analysis.bound(level)(powAnalysis)), new RewriteInline(ident, level, binding, body));
            };
        };
    };
};
var rewriteBranches = function (k) {
    var go = function (v) {
        if (v instanceof SemLet) {
            return new SemLet(v.value0, v.value1, Data_Functor.map(Data_Functor.functorFn)(go)(v.value2));
        };
        if (v instanceof SemLetRec) {
            return new SemLetRec(v.value0, Data_Functor.map(Data_Functor.functorFn)(go)(v.value1));
        };
        if (v instanceof SemBranch) {
            return new SemBranch(Data_Functor.map(Data_Array_NonEmpty_Internal.functorNonEmptyArray)(function (v1) {
                return new SemConditional(v1.value0, Data_Functor.map(Data_Lazy.functorLazy)(go)(v1.value1));
            })(v.value0), Data_Functor.map(Data_Lazy.functorLazy)(go)(v.value1));
        };
        return k(v);
    };
    return go;
};
var purely = function (v) {
    if (v.effect) {
        return {
            currentLevel: v.currentLevel,
            lookupExtern: v.lookupExtern,
            analyze: v.analyze,
            effect: false
        };
    };
    if (Data_Boolean.otherwise) {
        return v;
    };
    throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Semantics (line 1195, column 1 - line 1195, column 21): " + [ v.constructor.name ]);
};
var primOpOrdNot = function (v) {
    if (v instanceof PureScript_Backend_Optimizer_Syntax.OpEq) {
        return PureScript_Backend_Optimizer_Syntax.OpNotEq.value;
    };
    if (v instanceof PureScript_Backend_Optimizer_Syntax.OpNotEq) {
        return PureScript_Backend_Optimizer_Syntax.OpEq.value;
    };
    if (v instanceof PureScript_Backend_Optimizer_Syntax.OpLt) {
        return PureScript_Backend_Optimizer_Syntax.OpGte.value;
    };
    if (v instanceof PureScript_Backend_Optimizer_Syntax.OpLte) {
        return PureScript_Backend_Optimizer_Syntax.OpGt.value;
    };
    if (v instanceof PureScript_Backend_Optimizer_Syntax.OpGt) {
        return PureScript_Backend_Optimizer_Syntax.OpLte.value;
    };
    if (v instanceof PureScript_Backend_Optimizer_Syntax.OpGte) {
        return PureScript_Backend_Optimizer_Syntax.OpLt.value;
    };
    throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Semantics (line 871, column 16 - line 877, column 16): " + [ v.constructor.name ]);
};
var nextLevel = function (v) {
    return new Data_Tuple.Tuple(v.currentLevel, {
        lookupExtern: v.lookupExtern,
        analyze: v.analyze,
        effect: v.effect,
        currentLevel: v.currentLevel + 1 | 0
    });
};
var neutralSpine = /* #__PURE__ */ (function () {
    var go = function (hd) {
        return function (v) {
            if (v instanceof ExternApp) {
                return new NeutApp(hd, v.value0);
            };
            if (v instanceof ExternUncurriedApp) {
                return new NeutUncurriedApp(hd, v.value0);
            };
            if (v instanceof ExternAccessor) {
                return new NeutAccessor(hd, v.value0);
            };
            if (v instanceof ExternPrimOp) {
                return new NeutPrimOp(new PureScript_Backend_Optimizer_Syntax.Op1(v.value0, hd));
            };
            throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Semantics (line 488, column 11 - line 496, column 30): " + [ v.constructor.name ]);
        };
    };
    return Data_Foldable.foldl(Data_Foldable.foldableArray)(go);
})();
var neutralApp = function (hd) {
    return function (spine) {
        if (Data_Array["null"](spine)) {
            return hd;
        };
        if (Data_Boolean.otherwise) {
            if (hd instanceof NeutApp) {
                return new NeutApp(hd.value0, Data_Semigroup.append(Data_Semigroup.semigroupArray)(hd.value1)(spine));
            };
            return new NeutApp(hd, spine);
        };
        throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Semantics (line 498, column 1 - line 498, column 77): " + [ hd.constructor.name, spine.constructor.name ]);
    };
};
var lookupLocal = function (v) {
    return function (v1) {
        return Data_Array.index(v.locals)(v1);
    };
};
var liftString = function ($2092) {
    return NeutLit.create(PureScript_Backend_Optimizer_CoreFn.LitString.create($2092));
};
var liftOp2 = function (op) {
    return function (a) {
        return function (b) {
            return new NeutPrimOp(new PureScript_Backend_Optimizer_Syntax.Op2(op, a, b));
        };
    };
};
var liftOp1 = function (op) {
    return function (a) {
        return new NeutPrimOp(new PureScript_Backend_Optimizer_Syntax.Op1(op, a));
    };
};
var liftNumber = function ($2093) {
    return NeutLit.create(PureScript_Backend_Optimizer_CoreFn.LitNumber.create($2093));
};
var liftInt = function ($2094) {
    return NeutLit.create(PureScript_Backend_Optimizer_CoreFn.LitInt.create($2094));
};
var liftBoolean = function ($2095) {
    return NeutLit.create(PureScript_Backend_Optimizer_CoreFn.LitBoolean.create($2095));
};
var isSimplePredicate = function (v) {
    if (v instanceof ExprSyntax) {
        if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.Lit) {
            return true;
        };
        if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.Var) {
            return true;
        };
        if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.Local) {
            return true;
        };
        if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.PrimOp) {
            return true;
        };
        return false;
    };
    return false;
};
var isReference = function (v) {
    if (v instanceof PureScript_Backend_Optimizer_Syntax.Var) {
        return true;
    };
    if (v instanceof PureScript_Backend_Optimizer_Syntax.Local) {
        return true;
    };
    return false;
};
var isRefExpr = function (v) {
    if (v instanceof PureScript_Backend_Optimizer_Syntax.Var) {
        return true;
    };
    if (v instanceof PureScript_Backend_Optimizer_Syntax.Lit) {
        return true;
    };
    if (v instanceof PureScript_Backend_Optimizer_Syntax.CtorSaturated) {
        return true;
    };
    if (v instanceof PureScript_Backend_Optimizer_Syntax.Accessor) {
        return true;
    };
    if (v instanceof PureScript_Backend_Optimizer_Syntax.Update) {
        return true;
    };
    if (v instanceof PureScript_Backend_Optimizer_Syntax.PrimOp) {
        return true;
    };
    return false;
};
var isPartialAssocOp = function ($copy_v) {
    var $tco_done = false;
    var $tco_result;
    function $tco_loop(v) {
        if (v instanceof SemTyped) {
            $copy_v = v.value1;
            return;
        };
        if (v instanceof SemAssocOp) {
            $tco_done = true;
            return true;
        };
        $tco_done = true;
        return false;
    };
    while (!$tco_done) {
        $tco_result = $tco_loop($copy_v);
    };
    return $tco_result;
};
var isKnownEffect = function ($2096) {
    return (function (v) {
        if (v instanceof Data_Maybe.Just && v.value0 instanceof PureScript_Backend_Optimizer_Syntax.PrimEffect) {
            return true;
        };
        if (v instanceof Data_Maybe.Just && v.value0 instanceof PureScript_Backend_Optimizer_Syntax.UncurriedEffectApp) {
            return true;
        };
        if (v instanceof Data_Maybe.Just && v.value0 instanceof PureScript_Backend_Optimizer_Syntax.EffectBind) {
            return true;
        };
        if (v instanceof Data_Maybe.Just && v.value0 instanceof PureScript_Backend_Optimizer_Syntax.EffectDefer) {
            return true;
        };
        return false;
    })(syntaxOf($2096));
};
var isEffectSemantics = function ($copy_v) {
    var $tco_done = false;
    var $tco_result;
    function $tco_loop(v) {
        if (v instanceof SemTyped) {
            $copy_v = v.value1;
            return;
        };
        if (v instanceof SemMkEffectFn) {
            $tco_done = true;
            return true;
        };
        if (v instanceof SemEffectBind) {
            $tco_done = true;
            return true;
        };
        if (v instanceof SemEffectPure) {
            $tco_done = true;
            return true;
        };
        if (v instanceof SemEffectDefer) {
            $tco_done = true;
            return true;
        };
        $tco_done = true;
        return false;
    };
    while (!$tco_done) {
        $tco_result = $tco_loop($copy_v);
    };
    return $tco_result;
};
var isAssocPrimOp = function (v) {
    if (v instanceof PureScript_Backend_Optimizer_Syntax.OpIntNum && v.value0 instanceof PureScript_Backend_Optimizer_Syntax.OpAdd) {
        return true;
    };
    if (v instanceof PureScript_Backend_Optimizer_Syntax.OpIntNum && v.value0 instanceof PureScript_Backend_Optimizer_Syntax.OpMultiply) {
        return true;
    };
    if (v instanceof PureScript_Backend_Optimizer_Syntax.OpNumberNum && v.value0 instanceof PureScript_Backend_Optimizer_Syntax.OpAdd) {
        return true;
    };
    if (v instanceof PureScript_Backend_Optimizer_Syntax.OpNumberNum && v.value0 instanceof PureScript_Backend_Optimizer_Syntax.OpMultiply) {
        return true;
    };
    if (v instanceof PureScript_Backend_Optimizer_Syntax.OpStringAppend) {
        return true;
    };
    return false;
};
var isAbs = function ($2097) {
    return (function (v) {
        if (v instanceof Data_Maybe.Just && v.value0 instanceof PureScript_Backend_Optimizer_Syntax.Abs) {
            return true;
        };
        if (v instanceof Data_Maybe.Just && v.value0 instanceof PureScript_Backend_Optimizer_Syntax.UncurriedAbs) {
            return true;
        };
        if (v instanceof Data_Maybe.Just && v.value0 instanceof PureScript_Backend_Optimizer_Syntax.UncurriedEffectAbs) {
            return true;
        };
        if (v instanceof Data_Maybe.Just && v.value0 instanceof PureScript_Backend_Optimizer_Syntax.EffectDefer) {
            return true;
        };
        return false;
    })(syntaxOf($2097));
};
var shouldInlineLet = function (level) {
    return function (a) {
        return function (b) {
            var v = PureScript_Backend_Optimizer_Analysis.analysisOf(hasAnalysisBackendExpr)(a);
            var v1 = PureScript_Backend_Optimizer_Analysis.analysisOf(hasAnalysisBackendExpr)(b);
            var v2 = Data_Map_Internal.lookup(PureScript_Backend_Optimizer_Syntax.ordLevel)(level)(v1.usages);
            if (v2 instanceof Data_Maybe.Nothing) {
                return true;
            };
            if (v2 instanceof Data_Maybe.Just) {
                return Data_Eq.eq(PureScript_Backend_Optimizer_Analysis.eqComplexity)(v.complexity)(PureScript_Backend_Optimizer_Analysis.Trivial.value) || (Data_Eq.eq(PureScript_Backend_Optimizer_Analysis.eqCapture)(v2.value0.captured)(PureScript_Backend_Optimizer_Analysis.CaptureNone.value) && v2.value0.total === 1 || (Data_Ord.lessThanOrEq(PureScript_Backend_Optimizer_Analysis.ordCapture)(v2.value0.captured)(PureScript_Backend_Optimizer_Analysis.CaptureBranch.value) && (Data_Ord.lessThanOrEq(PureScript_Backend_Optimizer_Analysis.ordComplexity)(v.complexity)(PureScript_Backend_Optimizer_Analysis.Deref.value) && v.size < 5) || (Data_Eq.eq(PureScript_Backend_Optimizer_Analysis.eqComplexity)(v.complexity)(PureScript_Backend_Optimizer_Analysis.Deref.value) && v2.value0.call === v2.value0.total || (Data_Eq.eq(PureScript_Backend_Optimizer_Analysis.eqComplexity)(v.complexity)(PureScript_Backend_Optimizer_Analysis.KnownSize.value) && v2.value0.total === 1 || (isAbs(a) && (v2.value0.total === 1 || (Data_Map_Internal.isEmpty(v.usages) || v.size < 16)) || isKnownEffect(a) && v2.value0.total === 1)))));
            };
            throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Semantics (line 1614, column 3 - line 1624, column 43): " + [ v2.constructor.name ]);
        };
    };
};
var insertDirective = function (ref) {
    return function (acc) {
        return function (dir) {
            return Data_Map_Internal.alter(ordEvalRef)(function (v) {
                if (v instanceof Data_Maybe.Just) {
                    return new Data_Maybe.Just(Data_Map_Internal.insert(ordInlineAccessor)(acc)(dir)(v.value0));
                };
                if (v instanceof Data_Maybe.Nothing) {
                    return new Data_Maybe.Just(Data_Map_Internal.singleton(acc)(dir));
                };
                throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Semantics (line 200, column 3 - line 204, column 35): " + [ v.constructor.name ]);
            })(ref);
        };
    };
};
var guardFailOver = function (dictFoldable) {
    return function (f) {
        return function (as) {
            return function (k) {
                var toFail = function (expr) {
                    if (expr instanceof NeutFail) {
                        return new Data_Maybe.Just(expr);
                    };
                    return Data_Maybe.Nothing.value;
                };
                var v = Data_Foldable.findMap(dictFoldable)(function ($2098) {
                    return toFail(f($2098));
                })(as);
                if (v instanceof Data_Maybe.Just) {
                    return v.value0;
                };
                if (v instanceof Data_Maybe.Nothing) {
                    return k(as);
                };
                throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Semantics (line 1807, column 3 - line 1809, column 20): " + [ v.constructor.name ]);
            };
        };
    };
};
var guardFail = function (sem) {
    return function (k) {
        if (sem instanceof NeutFail) {
            return new NeutFail(sem.value0);
        };
        return k(sem);
    };
};
var foldBackendExpr = function (foldSyntax) {
    return function (foldRewrite) {
        var go = function (v) {
            if (v instanceof ExprSyntax) {
                return foldSyntax(Data_Functor.map(PureScript_Backend_Optimizer_Syntax.functorBackendSyntax)(go)(v.value1));
            };
            if (v instanceof ExprRewrite) {
                return foldRewrite(v.value1)((function () {
                    if (v.value1 instanceof RewriteInline) {
                        return foldSyntax(new PureScript_Backend_Optimizer_Syntax.Let(v.value1.value0, v.value1.value1, go(v.value1.value2), go(v.value1.value3)));
                    };
                    if (v.value1 instanceof RewriteUncurry) {
                        return foldSyntax(new PureScript_Backend_Optimizer_Syntax.Let(v.value1.value0, v.value1.value1, foldSyntax(new PureScript_Backend_Optimizer_Syntax.Abs(v.value1.value2, go(v.value1.value3))), go(v.value1.value4)));
                    };
                    if (v.value1 instanceof RewriteStop) {
                        return foldSyntax(new PureScript_Backend_Optimizer_Syntax.Var(v.value1.value0));
                    };
                    if (v.value1 instanceof RewriteUnpackOp) {
                        if (v.value1.value2 instanceof UnpackRecord) {
                            return foldSyntax(new PureScript_Backend_Optimizer_Syntax.Let(v.value1.value0, v.value1.value1, foldSyntax(new PureScript_Backend_Optimizer_Syntax.Lit(new PureScript_Backend_Optimizer_CoreFn.LitRecord(Data_Functor.map(Data_Functor.functorArray)(Data_Functor.map(PureScript_Backend_Optimizer_CoreFn.functorProp)(go))(v.value1.value2.value0)))), go(v.value1.value3)));
                        };
                        if (v.value1.value2 instanceof UnpackUpdate) {
                            return foldSyntax(new PureScript_Backend_Optimizer_Syntax.Let(v.value1.value0, v.value1.value1, foldSyntax(new PureScript_Backend_Optimizer_Syntax.Update(go(v.value1.value2.value0), Data_Functor.map(Data_Functor.functorArray)(Data_Functor.map(PureScript_Backend_Optimizer_CoreFn.functorProp)(go))(v.value1.value2.value1))), go(v.value1.value3)));
                        };
                        if (v.value1.value2 instanceof UnpackArray) {
                            return foldSyntax(new PureScript_Backend_Optimizer_Syntax.Let(v.value1.value0, v.value1.value1, foldSyntax(new PureScript_Backend_Optimizer_Syntax.Lit(new PureScript_Backend_Optimizer_CoreFn.LitArray(Data_Functor.map(Data_Functor.functorArray)(go)(v.value1.value2.value0)))), go(v.value1.value3)));
                        };
                        if (v.value1.value2 instanceof UnpackData) {
                            return foldSyntax(new PureScript_Backend_Optimizer_Syntax.Let(v.value1.value0, v.value1.value1, foldSyntax(new PureScript_Backend_Optimizer_Syntax.CtorSaturated(v.value1.value2.value0, v.value1.value2.value1, v.value1.value2.value2, v.value1.value2.value3, Data_Functor.map(Data_Functor.functorArray)(Data_Functor.map(Data_Tuple.functorTuple)(go))(v.value1.value2.value4))), go(v.value1.value3)));
                        };
                        throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Semantics (line 1742, column 11 - line 1750, column 117): " + [ v.value1.value2.constructor.name ]);
                    };
                    if (v.value1 instanceof RewriteDistBranchesLet) {
                        return foldSyntax(new PureScript_Backend_Optimizer_Syntax.Let(v.value1.value0, v.value1.value1, foldSyntax(new PureScript_Backend_Optimizer_Syntax.Branch(Data_Functor.map(Data_Array_NonEmpty_Internal.functorNonEmptyArray)(Data_Functor.map(PureScript_Backend_Optimizer_Syntax.functorPair)(go))(v.value1.value2), go(v.value1.value3))), go(v.value1.value4)));
                    };
                    if (v.value1 instanceof RewriteDistBranchesOp) {
                        var branches$prime = foldSyntax(new PureScript_Backend_Optimizer_Syntax.Branch(Data_Functor.map(Data_Array_NonEmpty_Internal.functorNonEmptyArray)(Data_Functor.map(PureScript_Backend_Optimizer_Syntax.functorPair)(go))(v.value1.value0), go(v.value1.value1)));
                        if (v.value1.value2 instanceof DistApp) {
                            return foldSyntax(new PureScript_Backend_Optimizer_Syntax.App(branches$prime, Data_Functor.map(Data_Array_NonEmpty_Internal.functorNonEmptyArray)(go)(v.value1.value2.value0)));
                        };
                        if (v.value1.value2 instanceof DistUncurriedApp) {
                            return foldSyntax(new PureScript_Backend_Optimizer_Syntax.UncurriedApp(branches$prime, Data_Functor.map(Data_Functor.functorArray)(go)(v.value1.value2.value0)));
                        };
                        if (v.value1.value2 instanceof DistAccessor) {
                            return foldSyntax(new PureScript_Backend_Optimizer_Syntax.Accessor(branches$prime, v.value1.value2.value0));
                        };
                        if (v.value1.value2 instanceof DistPrimOp1) {
                            return foldSyntax(new PureScript_Backend_Optimizer_Syntax.PrimOp(new PureScript_Backend_Optimizer_Syntax.Op1(v.value1.value2.value0, branches$prime)));
                        };
                        if (v.value1.value2 instanceof DistPrimOp2L) {
                            return foldSyntax(new PureScript_Backend_Optimizer_Syntax.PrimOp(new PureScript_Backend_Optimizer_Syntax.Op2(v.value1.value2.value0, branches$prime, go(v.value1.value2.value1))));
                        };
                        if (v.value1.value2 instanceof DistPrimOp2R) {
                            return foldSyntax(new PureScript_Backend_Optimizer_Syntax.PrimOp(new PureScript_Backend_Optimizer_Syntax.Op2(v.value1.value2.value1, go(v.value1.value2.value0), branches$prime)));
                        };
                        throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Semantics (line 1755, column 11 - line 1767, column 63): " + [ v.value1.value2.constructor.name ]);
                    };
                    throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Semantics (line 1734, column 7 - line 1767, column 63): " + [ v.value1.constructor.name ]);
                })());
            };
            throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Semantics (line 1730, column 8 - line 1767, column 63): " + [ v.constructor.name ]);
        };
        return go;
    };
};
var freeze = function (init) {
    return new Data_Tuple.Tuple(PureScript_Backend_Optimizer_Analysis.analysisOf(hasAnalysisBackendExpr)(init), foldBackendExpr(NeutralExpr)(function (v) {
        return function (neutExpr) {
            return neutExpr;
        };
    })(init));
};
var $lazy_floatLetWith = /* #__PURE__ */ $runtime_lazy("floatLetWith", "PureScript.Backend.Optimizer.Semantics", function () {
    var go = function ($copy_f) {
        return function ($copy_ident1) {
            return function ($copy_binding1) {
                return function ($copy_k1) {
                    var $tco_var_f = $copy_f;
                    var $tco_var_ident1 = $copy_ident1;
                    var $tco_var_binding1 = $copy_binding1;
                    var $tco_done = false;
                    var $tco_result;
                    function $tco_loop(f, ident1, binding1, k1) {
                        if (binding1 instanceof SemLet) {
                            $tco_var_f = $lazy_makeLet(631);
                            $tco_var_ident1 = binding1.value0;
                            $tco_var_binding1 = binding1.value1;
                            $copy_k1 = function (nextBinding2) {
                                return f(ident1)(binding1.value2(nextBinding2))(k1);
                            };
                            return;
                        };
                        if (binding1 instanceof SemLetRec) {
                            $tco_done = true;
                            return new SemLetRec(binding1.value0, function (nextBindings) {
                                return $lazy_makeLet(635)(ident1)(binding1.value1(nextBindings))(k1);
                            });
                        };
                        if (binding1 instanceof NeutFail) {
                            $tco_done = true;
                            return binding1;
                        };
                        $tco_done = true;
                        return f(ident1)(binding1)(k1);
                    };
                    while (!$tco_done) {
                        $tco_result = $tco_loop($tco_var_f, $tco_var_ident1, $tco_var_binding1, $copy_k1);
                    };
                    return $tco_result;
                };
            };
        };
    };
    return go;
});
var $lazy_makeLet = /* #__PURE__ */ $runtime_lazy("makeLet", "PureScript.Backend.Optimizer.Semantics", function () {
    var go = function (ident) {
        return function (binding) {
            return function (k) {
                if (binding instanceof SemRef && binding.value1.length === 0) {
                    return k(binding);
                };
                if (binding instanceof NeutLocal) {
                    return k(binding);
                };
                if (binding instanceof NeutStop) {
                    return k(binding);
                };
                if (binding instanceof NeutVar) {
                    return k(binding);
                };
                return new SemLet(ident, binding, k);
            };
        };
    };
    return $lazy_floatLetWith(604)(go);
});
var floatLetWith = /* #__PURE__ */ $lazy_floatLetWith(621);
var makeLet = /* #__PURE__ */ $lazy_makeLet(603);
var floatLet = /* #__PURE__ */ (function () {
    return floatLetWith(Data_Function["const"](Data_Function.applyFlipped))(Data_Maybe.Nothing.value);
})();
var $lazy_makeEffectBind = /* #__PURE__ */ $runtime_lazy("makeEffectBind", "PureScript.Backend.Optimizer.Semantics", function () {
    var go = function ($copy_ident1) {
        return function ($copy_binding1) {
            return function ($copy_k1) {
                var $tco_var_ident1 = $copy_ident1;
                var $tco_var_binding1 = $copy_binding1;
                var $tco_done = false;
                var $tco_result;
                function $tco_loop(ident1, binding1, k1) {
                    if (binding1 instanceof SemLet) {
                        $tco_done = true;
                        return makeLet(binding1.value0)(binding1.value1)(function (nextBinding2) {
                            return $lazy_makeEffectBind(592)(ident1)(binding1.value2(nextBinding2))(k1);
                        });
                    };
                    if (binding1 instanceof SemEffectBind) {
                        $tco_var_ident1 = binding1.value0;
                        $tco_var_binding1 = binding1.value1;
                        $copy_k1 = function (nextBinding2) {
                            return $lazy_makeEffectBind(595)(ident1)(binding1.value2(nextBinding2))(k1);
                        };
                        return;
                    };
                    if (binding1 instanceof SemEffectDefer) {
                        $tco_done = true;
                        return new SemEffectDefer(floatLet(binding1.value0)(function (nextBinding2) {
                            return $lazy_makeEffectBind(598)(ident1)(nextBinding2)(k1);
                        }));
                    };
                    $tco_done = true;
                    return floatLet(binding1)(function (nextBinding2) {
                        return new SemEffectBind(ident1, nextBinding2, k1);
                    });
                };
                while (!$tco_done) {
                    $tco_result = $tco_loop($tco_var_ident1, $tco_var_binding1, $copy_k1);
                };
                return $tco_result;
            };
        };
    };
    return go;
});
var makeEffectBind = /* #__PURE__ */ $lazy_makeEffectBind(586);
var evalUpdate = function (lhs) {
    return function (props) {
        return floatLet(lhs)(function (v) {
            if (v instanceof SemTyped) {
                return evalUpdate(v.value1)(props);
            };
            if (v instanceof NeutLit && v.value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitRecord) {
                return new NeutLit(new PureScript_Backend_Optimizer_CoreFn.LitRecord(Data_Functor.map(Data_Functor.functorArray)(Data_Array_NonEmpty.head)(Data_Array.groupAllBy(Data_Ord.comparing(Data_Ord.ordString)(PureScript_Backend_Optimizer_CoreFn.propKey))(Data_Semigroup.append(Data_Semigroup.semigroupArray)(props)(v.value0.value0)))));
            };
            if (v instanceof NeutUpdate) {
                return new NeutUpdate(v.value0, Data_Functor.map(Data_Functor.functorArray)(Data_Array_NonEmpty.head)(Data_Array.groupAllBy(Data_Ord.comparing(Data_Ord.ordString)(PureScript_Backend_Optimizer_CoreFn.propKey))(Data_Semigroup.append(Data_Semigroup.semigroupArray)(props)(v.value1))));
            };
            return new NeutUpdate(v, props);
        });
    };
};
var evalUncurriedBeta = function (fn) {
    return function (mk) {
        return function (spine) {
            var go = function (v) {
                return function (v1) {
                    if (v instanceof MkFnNext && (v1 instanceof Data_List_Types.Cons && v1.value0 instanceof NeutFail)) {
                        return new NeutFail(v1.value0.value0);
                    };
                    if (v instanceof MkFnNext && v1 instanceof Data_List_Types.Cons) {
                        return makeLet(Data_Maybe.Nothing.value)(v1.value0)(function (nextArg) {
                            return go(v.value1(nextArg))(v1.value1);
                        });
                    };
                    if (v instanceof MkFnNext) {
                        return Partial_Unsafe.unsafeCrashWith("Uncurried function applied to too few arguments");
                    };
                    if (v instanceof MkFnApplied && v1 instanceof Data_List_Types.Nil) {
                        return v.value0;
                    };
                    if (v instanceof MkFnApplied) {
                        return fn(v.value0)(Data_List.toUnfoldable(Data_Unfoldable.unfoldableArray)(v1));
                    };
                    throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Semantics (line 459, column 8 - line 470, column 36): " + [ v.constructor.name, v1.constructor.name ]);
                };
            };
            return go(mk)(Data_List.fromFoldable(Data_Foldable.foldableArray)(spine));
        };
    };
};

// Duplicate because inlined operators behave differently with NaN.
// This just ensures we get the same behavior for const eval.
var evalPrimOpOrdNumber = function (op) {
    return function (x) {
        return function (y) {
            if (op instanceof PureScript_Backend_Optimizer_Syntax.OpEq) {
                return x === y;
            };
            if (op instanceof PureScript_Backend_Optimizer_Syntax.OpNotEq) {
                return x !== y;
            };
            if (op instanceof PureScript_Backend_Optimizer_Syntax.OpGt) {
                return x > y;
            };
            if (op instanceof PureScript_Backend_Optimizer_Syntax.OpGte) {
                return x >= y;
            };
            if (op instanceof PureScript_Backend_Optimizer_Syntax.OpLt) {
                return x < y;
            };
            if (op instanceof PureScript_Backend_Optimizer_Syntax.OpLte) {
                return x <= y;
            };
            throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Semantics (line 805, column 30 - line 811, column 18): " + [ op.constructor.name ]);
        };
    };
};
var evalPrimOpOrd = function (dictOrd) {
    var Eq0 = dictOrd.Eq0();
    return function (op) {
        return function (x) {
            return function (y) {
                if (op instanceof PureScript_Backend_Optimizer_Syntax.OpEq) {
                    return Data_Eq.eq(Eq0)(x)(y);
                };
                if (op instanceof PureScript_Backend_Optimizer_Syntax.OpNotEq) {
                    return Data_Eq.notEq(Eq0)(x)(y);
                };
                if (op instanceof PureScript_Backend_Optimizer_Syntax.OpGt) {
                    return Data_Ord.greaterThan(dictOrd)(x)(y);
                };
                if (op instanceof PureScript_Backend_Optimizer_Syntax.OpGte) {
                    return Data_Ord.greaterThanOrEq(dictOrd)(x)(y);
                };
                if (op instanceof PureScript_Backend_Optimizer_Syntax.OpLt) {
                    return Data_Ord.lessThan(dictOrd)(x)(y);
                };
                if (op instanceof PureScript_Backend_Optimizer_Syntax.OpLte) {
                    return Data_Ord.lessThanOrEq(dictOrd)(x)(y);
                };
                throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Semantics (line 794, column 24 - line 800, column 18): " + [ op.constructor.name ]);
            };
        };
    };
};
var evalPrimOpNot = function (v) {
    if (v instanceof PureScript_Backend_Optimizer_Syntax.Op1) {
        if (v.value0 instanceof PureScript_Backend_Optimizer_Syntax.OpBooleanNot) {
            return v.value1;
        };
        return liftOp1(PureScript_Backend_Optimizer_Syntax.OpBooleanNot.value)(liftOp1(v.value0)(v.value1));
    };
    if (v instanceof PureScript_Backend_Optimizer_Syntax.Op2) {
        if (v.value0 instanceof PureScript_Backend_Optimizer_Syntax.OpIntOrd) {
            return liftOp2(new PureScript_Backend_Optimizer_Syntax.OpIntOrd(primOpOrdNot(v.value0.value0)))(v.value1)(v.value2);
        };
        if (v.value0 instanceof PureScript_Backend_Optimizer_Syntax.OpNumberOrd) {
            return liftOp2(new PureScript_Backend_Optimizer_Syntax.OpNumberOrd(primOpOrdNot(v.value0.value0)))(v.value1)(v.value2);
        };
        if (v.value0 instanceof PureScript_Backend_Optimizer_Syntax.OpStringOrd) {
            return liftOp2(new PureScript_Backend_Optimizer_Syntax.OpStringOrd(primOpOrdNot(v.value0.value0)))(v.value1)(v.value2);
        };
        if (v.value0 instanceof PureScript_Backend_Optimizer_Syntax.OpCharOrd) {
            return liftOp2(new PureScript_Backend_Optimizer_Syntax.OpCharOrd(primOpOrdNot(v.value0.value0)))(v.value1)(v.value2);
        };
        if (v.value0 instanceof PureScript_Backend_Optimizer_Syntax.OpBooleanOrd) {
            return liftOp2(new PureScript_Backend_Optimizer_Syntax.OpBooleanOrd(primOpOrdNot(v.value0.value0)))(v.value1)(v.value2);
        };
        return liftOp1(PureScript_Backend_Optimizer_Syntax.OpBooleanNot.value)(liftOp2(v.value0)(v.value1)(v.value2));
    };
    throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Semantics (line 848, column 17 - line 868, column 46): " + [ v.constructor.name ]);
};
var evalEvalRef = function (v) {
    if (v instanceof EvalExtern) {
        return new NeutVar(v.value0);
    };
    if (v instanceof EvalLocal) {
        return new NeutLocal(v.value0, v.value1);
    };
    throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Semantics (line 956, column 15 - line 960, column 24): " + [ v.constructor.name ]);
};
var $$eval = function (dict) {
    return dict["eval"];
};
var evalPair = function (dictEval) {
    return function (env) {
        return function (v) {
            return new SemConditional(Data_Lazy.defer(function (v1) {
                return $$eval(dictEval)(env)(v.value0);
            }), Data_Lazy.defer(function (v1) {
                return $$eval(dictEval)(env)(v.value1);
            }));
        };
    };
};
var effectfully = function (v) {
    if (v.effect) {
        return v;
    };
    if (Data_Boolean.otherwise) {
        return {
            currentLevel: v.currentLevel,
            lookupExtern: v.lookupExtern,
            analyze: v.analyze,
            effect: true
        };
    };
    throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Semantics (line 1190, column 1 - line 1190, column 26): " + [ v.constructor.name ]);
};
var deref = function ($copy_v) {
    var $tco_done = false;
    var $tco_result;
    function $tco_loop(v) {
        if (v instanceof SemTyped) {
            $copy_v = v.value1;
            return;
        };
        if (v instanceof SemRef) {
            $tco_done = true;
            return Data_Lazy.force(v.value2);
        };
        $tco_done = true;
        return v;
    };
    while (!$tco_done) {
        $tco_result = $tco_loop($copy_v);
    };
    return $tco_result;
};
var evalBranches = function (v) {
    return function (initConds) {
        return function (initDef) {
            var go = function ($copy_acc) {
                return function ($copy_conds) {
                    return function ($copy_def) {
                        var $tco_var_acc = $copy_acc;
                        var $tco_var_conds = $copy_conds;
                        var $tco_done = false;
                        var $tco_result;
                        function $tco_loop(acc, conds, def) {
                            var v1 = Data_Array.uncons(conds);
                            if (v1 instanceof Data_Maybe.Just) {
                                var v2 = deref(Data_Lazy.force(v1.value0.head.value0));
                                if (v2 instanceof NeutLit && v2.value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitBoolean) {
                                    if (v2.value0.value0) {
                                        $tco_var_acc = acc;
                                        $tco_var_conds = [  ];
                                        $copy_def = v1.value0.head.value1;
                                        return;
                                    };
                                    if (Data_Boolean.otherwise) {
                                        $tco_var_acc = acc;
                                        $tco_var_conds = v1.value0.tail;
                                        $copy_def = def;
                                        return;
                                    };
                                };
                                if (v2 instanceof NeutFail) {
                                    $tco_var_acc = acc;
                                    $tco_var_conds = [  ];
                                    $copy_def = Data_Lazy.defer(function (v3) {
                                        return new NeutFail(v2.value0);
                                    });
                                    return;
                                };
                                $tco_var_acc = Data_Array.snoc(acc)(v1.value0.head);
                                $tco_var_conds = v1.value0.tail;
                                $copy_def = def;
                                return;
                            };
                            if (v1 instanceof Data_Maybe.Nothing) {
                                var v2 = Data_Array_NonEmpty.fromArray(acc);
                                if (v2 instanceof Data_Maybe.Just) {
                                    $tco_done = true;
                                    return new SemBranch(v2.value0, def);
                                };
                                if (v2 instanceof Data_Maybe.Nothing) {
                                    $tco_done = true;
                                    return Data_Lazy.force(def);
                                };
                                throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Semantics (line 564, column 7 - line 568, column 20): " + [ v2.constructor.name ]);
                            };
                            throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Semantics (line 552, column 22 - line 568, column 20): " + [ v1.constructor.name ]);
                        };
                        while (!$tco_done) {
                            $tco_result = $tco_loop($tco_var_acc, $tco_var_conds, $copy_def);
                        };
                        return $tco_result;
                    };
                };
            };
            return go([  ])(Data_Array_NonEmpty.toArray(initConds))(initDef);
        };
    };
};
var evalPrimOpNumInt = function (op) {
    return function (x) {
        return function (y) {
            var v = function (v1) {
                if (Data_Boolean.otherwise) {
                    return Data_Maybe.Nothing.value;
                };
                throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Semantics (line 825, column 1 - line 825, column 105): " + [ op.constructor.name, x.constructor.name, y.constructor.name ]);
            };
            var $1023 = deref(x);
            if ($1023 instanceof NeutLit && $1023.value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitInt) {
                var $1024 = deref(y);
                if ($1024 instanceof NeutLit && $1024.value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitInt) {
                    if (op instanceof PureScript_Backend_Optimizer_Syntax.OpAdd) {
                        var res = $1023.value0.value0 + $1024.value0.value0 | 0;
                        var $1026 = $1024.value0.value0 > 0 && res < $1023.value0.value0 || $1024.value0.value0 < 0 && res > $1023.value0.value0;
                        if ($1026) {
                            return Data_Maybe.Nothing.value;
                        };
                        return new Data_Maybe.Just(liftInt(res));
                    };
                    if (op instanceof PureScript_Backend_Optimizer_Syntax.OpMultiply) {
                        var res = $1023.value0.value0 * $1024.value0.value0 | 0;
                        var $1027 = $1023.value0.value0 !== Data_EuclideanRing.div(Data_EuclideanRing.euclideanRingInt)(res)($1024.value0.value0);
                        if ($1027) {
                            return Data_Maybe.Nothing.value;
                        };
                        return new Data_Maybe.Just(liftInt(res));
                    };
                    if (op instanceof PureScript_Backend_Optimizer_Syntax.OpSubtract) {
                        var res = $1023.value0.value0 - $1024.value0.value0 | 0;
                        var $1028 = $1024.value0.value0 > 0 && res > $1023.value0.value0 || $1024.value0.value0 < 0 && res < $1023.value0.value0;
                        if ($1028) {
                            return Data_Maybe.Nothing.value;
                        };
                        return new Data_Maybe.Just(liftInt(res));
                    };
                    if (op instanceof PureScript_Backend_Optimizer_Syntax.OpDivide) {
                        return new Data_Maybe.Just(liftInt(Data_EuclideanRing.div(Data_EuclideanRing.euclideanRingInt)($1023.value0.value0)($1024.value0.value0)));
                    };
                    throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Semantics (line 829, column 7 - line 843, column 33): " + [ op.constructor.name ]);
                };
                return v(true);
            };
            return v(true);
        };
    };
};
var evalPrimOpNumNumber = function (op) {
    return function (x) {
        return function (y) {
            var v = function (v1) {
                if (Data_Boolean.otherwise) {
                    return Data_Maybe.Nothing.value;
                };
                throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Semantics (line 813, column 1 - line 813, column 108): " + [ op.constructor.name, x.constructor.name, y.constructor.name ]);
            };
            var $1039 = deref(x);
            if ($1039 instanceof NeutLit && $1039.value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitNumber) {
                var $1040 = deref(y);
                if ($1040 instanceof NeutLit && $1040.value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitNumber) {
                    return new Data_Maybe.Just(liftNumber((function () {
                        if (op instanceof PureScript_Backend_Optimizer_Syntax.OpAdd) {
                            return $1039.value0.value0 + $1040.value0.value0;
                        };
                        if (op instanceof PureScript_Backend_Optimizer_Syntax.OpMultiply) {
                            return $1039.value0.value0 * $1040.value0.value0;
                        };
                        if (op instanceof PureScript_Backend_Optimizer_Syntax.OpSubtract) {
                            return $1039.value0.value0 - $1040.value0.value0;
                        };
                        if (op instanceof PureScript_Backend_Optimizer_Syntax.OpDivide) {
                            return $1039.value0.value0 / $1040.value0.value0;
                        };
                        throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Semantics (line 817, column 25 - line 821, column 26): " + [ op.constructor.name ]);
                    })()));
                };
                return v(true);
            };
            return v(true);
        };
    };
};
var evalRefSpine = function (env) {
    return function (ref) {
        return function (spine) {
            return function (sem) {
                return function (v) {
                    if (v instanceof ExternApp) {
                        return neutralSpine(evalEvalRef(ref))(spine);
                    };
                    if (v instanceof ExternUncurriedApp) {
                        return neutralSpine(evalEvalRef(ref))(spine);
                    };
                    if (v instanceof ExternAccessor) {
                        return evalAccessor(env)(Data_Lazy.force(sem))(v.value0);
                    };
                    if (v instanceof ExternPrimOp) {
                        return evalPrimOp(env)(new PureScript_Backend_Optimizer_Syntax.Op1(v.value0, Data_Lazy.force(sem)));
                    };
                    throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Semantics (line 945, column 34 - line 953, column 40): " + [ v.constructor.name ]);
                };
            };
        };
    };
};
var evalRef = function (v) {
    return function (ref) {
        return function (spine) {
            return function (last) {
                return function (sem) {
                    var spine$prime = snocSpine(spine)(last);
                    var v1 = function (v2) {
                        return new SemRef(ref, spine$prime, Data_Lazy.defer(function (v3) {
                            return deref(evalRefSpine(v)(ref)(spine$prime)(sem)(last));
                        }));
                    };
                    if (ref instanceof EvalExtern) {
                        var $1057 = v.evalExternSpine(v)(ref.value0)(spine$prime);
                        if ($1057 instanceof Data_Maybe.Just) {
                            return $1057.value0;
                        };
                        return v1(true);
                    };
                    return v1(true);
                };
            };
        };
    };
};
var evalPrimOp = function ($copy_env) {
    return function ($copy_v) {
        var $tco_var_env = $copy_env;
        var $tco_done = false;
        var $tco_result;
        function $tco_loop(env, v) {
            if (v instanceof PureScript_Backend_Optimizer_Syntax.Op1) {
                var v1 = function (v2) {
                    var v3 = function (v4) {
                        var v5 = function (v6) {
                            var v7 = function (v8) {
                                var v9 = function (v10) {
                                    var v11 = function (v12) {
                                        var v13 = function (v14) {
                                            if (v.value1 instanceof SemRef) {
                                                $tco_done = true;
                                                return evalRef(env)(v.value1.value0)(v.value1.value1)(new ExternPrimOp(v.value0))(v.value1.value2);
                                            };
                                            if (v.value1 instanceof NeutFail) {
                                                $tco_done = true;
                                                return new NeutFail(v.value1.value0);
                                            };
                                            $tco_done = true;
                                            return floatLet(v.value1)((function () {
                                                var $2099 = PureScript_Backend_Optimizer_Syntax.Op1.create(v.value0);
                                                return function ($2100) {
                                                    return NeutPrimOp.create($2099($2100));
                                                };
                                            })());
                                        };
                                        if (v.value0 instanceof PureScript_Backend_Optimizer_Syntax.OpNumberNegate) {
                                            var $1069 = deref(v.value1);
                                            if ($1069 instanceof NeutLit && $1069.value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitNumber) {
                                                $tco_done = true;
                                                return liftNumber(-$1069.value0.value0);
                                            };
                                            return v13(true);
                                        };
                                        return v13(true);
                                    };
                                    if (v.value0 instanceof PureScript_Backend_Optimizer_Syntax.OpIntNegate) {
                                        var $1074 = deref(v.value1);
                                        if ($1074 instanceof NeutLit && $1074.value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitInt) {
                                            $tco_done = true;
                                            return liftInt(-$1074.value0.value0 | 0);
                                        };
                                        return v11(true);
                                    };
                                    return v11(true);
                                };
                                if (v.value0 instanceof PureScript_Backend_Optimizer_Syntax.OpArrayLength) {
                                    var $1079 = deref(v.value1);
                                    if ($1079 instanceof NeutLit && $1079.value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitArray) {
                                        $tco_done = true;
                                        return liftInt(Data_Array.length($1079.value0.value0));
                                    };
                                    return v9(true);
                                };
                                return v9(true);
                            };
                            if (v.value0 instanceof PureScript_Backend_Optimizer_Syntax.OpIsTag) {
                                var $1084 = deref(v.value1);
                                if ($1084 instanceof NeutData) {
                                    $tco_done = true;
                                    return liftBoolean(Data_Eq.eq(eqQualified)(v.value0.value0)($1084.value0));
                                };
                                return v7(true);
                            };
                            return v7(true);
                        };
                        if (v.value0 instanceof PureScript_Backend_Optimizer_Syntax.OpIntBitNot) {
                            var $1093 = deref(v.value1);
                            if ($1093 instanceof NeutLit && $1093.value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitInt) {
                                $tco_done = true;
                                return liftInt(~$1093.value0.value0);
                            };
                            return v5(true);
                        };
                        return v5(true);
                    };
                    if (v.value0 instanceof PureScript_Backend_Optimizer_Syntax.OpBooleanNot) {
                        if (v.value1 instanceof NeutPrimOp) {
                            $tco_done = true;
                            return evalPrimOpNot(v.value1.value0);
                        };
                        return v3(true);
                    };
                    return v3(true);
                };
                if (v.value0 instanceof PureScript_Backend_Optimizer_Syntax.OpBooleanNot) {
                    var $1102 = deref(v.value1);
                    if ($1102 instanceof NeutLit && $1102.value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitBoolean) {
                        $tco_done = true;
                        return liftBoolean(!$1102.value0.value0);
                    };
                    return v1(true);
                };
                return v1(true);
            };
            if (v instanceof PureScript_Backend_Optimizer_Syntax.Op2) {
                var v1 = function (v2) {
                    var v3 = function (v4) {
                        var v5 = function (v6) {
                            var v7 = function (v8) {
                                var v9 = function (v10) {
                                    var v11 = function (v12) {
                                        var v13 = function (v14) {
                                            var v15 = function (v16) {
                                                var v17 = function (v18) {
                                                    var v19 = function (v20) {
                                                        var v21 = function (v22) {
                                                            var v23 = function (v24) {
                                                                var v25 = function (v26) {
                                                                    var v27 = function (v28) {
                                                                        var v29 = function (v30) {
                                                                            var v31 = function (v32) {
                                                                                var v33 = function (v34) {
                                                                                    var v35 = function (v36) {
                                                                                        var v37 = function (v38) {
                                                                                            var v39 = function (v40) {
                                                                                                var v41 = function (v42) {
                                                                                                    var v43 = function (v44) {
                                                                                                        var v45 = function (v46) {
                                                                                                            var v47 = function (v48) {
                                                                                                                var v49 = function (v50) {
                                                                                                                    var v51 = function (v52) {
                                                                                                                        var v53 = function (v54) {
                                                                                                                            if (v.value0 instanceof PureScript_Backend_Optimizer_Syntax.OpBooleanAnd) {
                                                                                                                                if (v.value1 instanceof NeutFail) {
                                                                                                                                    return new NeutFail(v.value1.value0);
                                                                                                                                };
                                                                                                                                if (v.value2 instanceof NeutFail) {
                                                                                                                                    return new NeutFail(v.value2.value0);
                                                                                                                                };
                                                                                                                                return new NeutPrimOp(new PureScript_Backend_Optimizer_Syntax.Op2(v.value0, v.value1, v.value2));
                                                                                                                            };
                                                                                                                            if (v.value0 instanceof PureScript_Backend_Optimizer_Syntax.OpBooleanOr) {
                                                                                                                                if (v.value1 instanceof NeutFail) {
                                                                                                                                    return new NeutFail(v.value1.value0);
                                                                                                                                };
                                                                                                                                if (v.value2 instanceof NeutFail) {
                                                                                                                                    return new NeutFail(v.value2.value0);
                                                                                                                                };
                                                                                                                                return new NeutPrimOp(new PureScript_Backend_Optimizer_Syntax.Op2(v.value0, v.value1, v.value2));
                                                                                                                            };
                                                                                                                            if (v.value1 instanceof NeutFail) {
                                                                                                                                return new NeutFail(v.value1.value0);
                                                                                                                            };
                                                                                                                            if (v.value2 instanceof NeutFail) {
                                                                                                                                return new NeutFail(v.value2.value0);
                                                                                                                            };
                                                                                                                            return floatLet(v.value1)(function (x$prime) {
                                                                                                                                return floatLet(v.value2)(function (y$prime) {
                                                                                                                                    var $1120 = isAssocPrimOp(v.value0);
                                                                                                                                    if ($1120) {
                                                                                                                                        return evalAssocOp(env)(new Data_Either.Right(v.value0))(x$prime)(y$prime);
                                                                                                                                    };
                                                                                                                                    return new NeutPrimOp(new PureScript_Backend_Optimizer_Syntax.Op2(v.value0, x$prime, y$prime));
                                                                                                                                });
                                                                                                                            });
                                                                                                                        };
                                                                                                                        if (v.value0 instanceof PureScript_Backend_Optimizer_Syntax.OpArrayIndex) {
                                                                                                                            if (v.value2 instanceof NeutLit && v.value2.value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitInt) {
                                                                                                                                return evalAccessor(env)(v.value1)(new PureScript_Backend_Optimizer_Syntax.GetIndex(v.value2.value0.value0));
                                                                                                                            };
                                                                                                                            return v53(true);
                                                                                                                        };
                                                                                                                        return v53(true);
                                                                                                                    };
                                                                                                                    if (v.value0 instanceof PureScript_Backend_Optimizer_Syntax.OpStringAppend) {
                                                                                                                        if (v.value1 instanceof NeutLit && v.value1.value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitString) {
                                                                                                                            if (v.value2 instanceof NeutLit && v.value2.value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitString) {
                                                                                                                                return liftString(v.value1.value0.value0 + v.value2.value0.value0);
                                                                                                                            };
                                                                                                                            return v51(true);
                                                                                                                        };
                                                                                                                        return v51(true);
                                                                                                                    };
                                                                                                                    return v51(true);
                                                                                                                };
                                                                                                                if (v.value0 instanceof PureScript_Backend_Optimizer_Syntax.OpStringOrd) {
                                                                                                                    var $1133 = deref(v.value1);
                                                                                                                    if ($1133 instanceof NeutLit && $1133.value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitString) {
                                                                                                                        var $1134 = deref(v.value2);
                                                                                                                        if ($1134 instanceof NeutLit && $1134.value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitString) {
                                                                                                                            return liftBoolean(evalPrimOpOrd(Data_Ord.ordString)(v.value0.value0)($1133.value0.value0)($1134.value0.value0));
                                                                                                                        };
                                                                                                                        return v49(true);
                                                                                                                    };
                                                                                                                    return v49(true);
                                                                                                                };
                                                                                                                return v49(true);
                                                                                                            };
                                                                                                            if (v.value0 instanceof PureScript_Backend_Optimizer_Syntax.OpNumberOrd) {
                                                                                                                var $1141 = deref(v.value1);
                                                                                                                if ($1141 instanceof NeutLit && $1141.value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitNumber) {
                                                                                                                    var $1142 = deref(v.value2);
                                                                                                                    if ($1142 instanceof NeutLit && $1142.value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitNumber) {
                                                                                                                        return liftBoolean(evalPrimOpOrdNumber(v.value0.value0)($1141.value0.value0)($1142.value0.value0));
                                                                                                                    };
                                                                                                                    return v47(true);
                                                                                                                };
                                                                                                                return v47(true);
                                                                                                            };
                                                                                                            return v47(true);
                                                                                                        };
                                                                                                        if (v.value0 instanceof PureScript_Backend_Optimizer_Syntax.OpNumberNum) {
                                                                                                            var $1149 = evalPrimOpNumNumber(v.value0.value0)(v.value1)(v.value2);
                                                                                                            if ($1149 instanceof Data_Maybe.Just) {
                                                                                                                return $1149.value0;
                                                                                                            };
                                                                                                            return v45(true);
                                                                                                        };
                                                                                                        return v45(true);
                                                                                                    };
                                                                                                    if (v.value0 instanceof PureScript_Backend_Optimizer_Syntax.OpNumberNum && v.value0.value0 instanceof PureScript_Backend_Optimizer_Syntax.OpSubtract) {
                                                                                                        var $1153 = deref(v.value1);
                                                                                                        if ($1153 instanceof NeutLit && ($1153.value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitNumber && $1153.value0.value0 === 0.0)) {
                                                                                                            $tco_var_env = env;
                                                                                                            $copy_v = new PureScript_Backend_Optimizer_Syntax.Op1(PureScript_Backend_Optimizer_Syntax.OpNumberNegate.value, v.value2);
                                                                                                            return;
                                                                                                        };
                                                                                                        $tco_done = true;
                                                                                                        return v43(true);
                                                                                                    };
                                                                                                    $tco_done = true;
                                                                                                    return v43(true);
                                                                                                };
                                                                                                if (v.value0 instanceof PureScript_Backend_Optimizer_Syntax.OpIntOrd) {
                                                                                                    var $1158 = deref(v.value1);
                                                                                                    if ($1158 instanceof NeutLit && $1158.value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitInt) {
                                                                                                        var $1159 = deref(v.value2);
                                                                                                        if ($1159 instanceof NeutLit && $1159.value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitInt) {
                                                                                                            $tco_done = true;
                                                                                                            return liftBoolean(evalPrimOpOrd(Data_Ord.ordInt)(v.value0.value0)($1158.value0.value0)($1159.value0.value0));
                                                                                                        };
                                                                                                        return v41(true);
                                                                                                    };
                                                                                                    return v41(true);
                                                                                                };
                                                                                                return v41(true);
                                                                                            };
                                                                                            if (v.value0 instanceof PureScript_Backend_Optimizer_Syntax.OpIntNum) {
                                                                                                var $1166 = evalPrimOpNumInt(v.value0.value0)(v.value1)(v.value2);
                                                                                                if ($1166 instanceof Data_Maybe.Just) {
                                                                                                    $tco_done = true;
                                                                                                    return $1166.value0;
                                                                                                };
                                                                                                return v39(true);
                                                                                            };
                                                                                            return v39(true);
                                                                                        };
                                                                                        if (v.value0 instanceof PureScript_Backend_Optimizer_Syntax.OpIntNum && v.value0.value0 instanceof PureScript_Backend_Optimizer_Syntax.OpSubtract) {
                                                                                            var $1170 = deref(v.value1);
                                                                                            if ($1170 instanceof NeutLit && ($1170.value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitInt && $1170.value0.value0 === 0)) {
                                                                                                $tco_var_env = env;
                                                                                                $copy_v = new PureScript_Backend_Optimizer_Syntax.Op1(PureScript_Backend_Optimizer_Syntax.OpIntNegate.value, v.value2);
                                                                                                return;
                                                                                            };
                                                                                            return v37(true);
                                                                                        };
                                                                                        return v37(true);
                                                                                    };
                                                                                    if (v.value0 instanceof PureScript_Backend_Optimizer_Syntax.OpIntBitZeroFillShiftRight) {
                                                                                        var $1175 = deref(v.value1);
                                                                                        if ($1175 instanceof NeutLit && $1175.value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitInt) {
                                                                                            var $1176 = deref(v.value2);
                                                                                            if ($1176 instanceof NeutLit && $1176.value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitInt) {
                                                                                                $tco_done = true;
                                                                                                return liftInt($1175.value0.value0 >>> $1176.value0.value0);
                                                                                            };
                                                                                            return v35(true);
                                                                                        };
                                                                                        return v35(true);
                                                                                    };
                                                                                    return v35(true);
                                                                                };
                                                                                if (v.value0 instanceof PureScript_Backend_Optimizer_Syntax.OpIntBitXor) {
                                                                                    var $1182 = deref(v.value1);
                                                                                    if ($1182 instanceof NeutLit && $1182.value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitInt) {
                                                                                        var $1183 = deref(v.value2);
                                                                                        if ($1183 instanceof NeutLit && $1183.value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitInt) {
                                                                                            $tco_done = true;
                                                                                            return liftInt($1182.value0.value0 ^ $1183.value0.value0);
                                                                                        };
                                                                                        return v33(true);
                                                                                    };
                                                                                    return v33(true);
                                                                                };
                                                                                return v33(true);
                                                                            };
                                                                            if (v.value0 instanceof PureScript_Backend_Optimizer_Syntax.OpIntBitShiftRight) {
                                                                                var $1189 = deref(v.value1);
                                                                                if ($1189 instanceof NeutLit && $1189.value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitInt) {
                                                                                    var $1190 = deref(v.value2);
                                                                                    if ($1190 instanceof NeutLit && $1190.value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitInt) {
                                                                                        $tco_done = true;
                                                                                        return liftInt($1189.value0.value0 >> $1190.value0.value0);
                                                                                    };
                                                                                    return v31(true);
                                                                                };
                                                                                return v31(true);
                                                                            };
                                                                            return v31(true);
                                                                        };
                                                                        if (v.value0 instanceof PureScript_Backend_Optimizer_Syntax.OpIntBitShiftLeft) {
                                                                            var $1196 = deref(v.value1);
                                                                            if ($1196 instanceof NeutLit && $1196.value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitInt) {
                                                                                var $1197 = deref(v.value2);
                                                                                if ($1197 instanceof NeutLit && $1197.value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitInt) {
                                                                                    $tco_done = true;
                                                                                    return liftInt($1196.value0.value0 << $1197.value0.value0);
                                                                                };
                                                                                return v29(true);
                                                                            };
                                                                            return v29(true);
                                                                        };
                                                                        return v29(true);
                                                                    };
                                                                    if (v.value0 instanceof PureScript_Backend_Optimizer_Syntax.OpIntBitOr) {
                                                                        var $1203 = deref(v.value1);
                                                                        if ($1203 instanceof NeutLit && $1203.value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitInt) {
                                                                            var $1204 = deref(v.value2);
                                                                            if ($1204 instanceof NeutLit && $1204.value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitInt) {
                                                                                $tco_done = true;
                                                                                return liftInt($1203.value0.value0 | $1204.value0.value0);
                                                                            };
                                                                            return v27(true);
                                                                        };
                                                                        return v27(true);
                                                                    };
                                                                    return v27(true);
                                                                };
                                                                if (v.value0 instanceof PureScript_Backend_Optimizer_Syntax.OpIntBitAnd) {
                                                                    var $1210 = deref(v.value1);
                                                                    if ($1210 instanceof NeutLit && $1210.value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitInt) {
                                                                        var $1211 = deref(v.value2);
                                                                        if ($1211 instanceof NeutLit && $1211.value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitInt) {
                                                                            $tco_done = true;
                                                                            return liftInt($1210.value0.value0 & $1211.value0.value0);
                                                                        };
                                                                        return v25(true);
                                                                    };
                                                                    return v25(true);
                                                                };
                                                                return v25(true);
                                                            };
                                                            if (v.value0 instanceof PureScript_Backend_Optimizer_Syntax.OpCharOrd) {
                                                                var $1217 = deref(v.value1);
                                                                if ($1217 instanceof NeutLit && $1217.value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitChar) {
                                                                    var $1218 = deref(v.value2);
                                                                    if ($1218 instanceof NeutLit && $1218.value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitChar) {
                                                                        $tco_done = true;
                                                                        return liftBoolean(evalPrimOpOrd(Data_Ord.ordChar)(v.value0.value0)($1217.value0.value0)($1218.value0.value0));
                                                                    };
                                                                    return v23(true);
                                                                };
                                                                return v23(true);
                                                            };
                                                            return v23(true);
                                                        };
                                                        if (v.value0 instanceof PureScript_Backend_Optimizer_Syntax.OpBooleanOrd) {
                                                            var $1225 = deref(v.value1);
                                                            if ($1225 instanceof NeutLit && $1225.value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitBoolean) {
                                                                var $1226 = deref(v.value2);
                                                                if ($1226 instanceof NeutLit && $1226.value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitBoolean) {
                                                                    $tco_done = true;
                                                                    return liftBoolean(evalPrimOpOrd(Data_Ord.ordBoolean)(v.value0.value0)($1225.value0.value0)($1226.value0.value0));
                                                                };
                                                                return v21(true);
                                                            };
                                                            return v21(true);
                                                        };
                                                        return v21(true);
                                                    };
                                                    if (v.value0 instanceof PureScript_Backend_Optimizer_Syntax.OpBooleanOrd && v.value0.value0 instanceof PureScript_Backend_Optimizer_Syntax.OpEq) {
                                                        var $1233 = deref(v.value2);
                                                        if ($1233 instanceof NeutLit && $1233.value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitBoolean) {
                                                            if ($1233.value0.value0) {
                                                                $tco_done = true;
                                                                return v.value1;
                                                            };
                                                            $tco_var_env = env;
                                                            $copy_v = new PureScript_Backend_Optimizer_Syntax.Op1(PureScript_Backend_Optimizer_Syntax.OpBooleanNot.value, v.value1);
                                                            return;
                                                        };
                                                        return v19(true);
                                                    };
                                                    return v19(true);
                                                };
                                                if (v.value0 instanceof PureScript_Backend_Optimizer_Syntax.OpBooleanOrd && v.value0.value0 instanceof PureScript_Backend_Optimizer_Syntax.OpEq) {
                                                    var $1239 = deref(v.value1);
                                                    if ($1239 instanceof NeutLit && $1239.value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitBoolean) {
                                                        if ($1239.value0.value0) {
                                                            $tco_done = true;
                                                            return v.value2;
                                                        };
                                                        $tco_var_env = env;
                                                        $copy_v = new PureScript_Backend_Optimizer_Syntax.Op1(PureScript_Backend_Optimizer_Syntax.OpBooleanNot.value, v.value2);
                                                        return;
                                                    };
                                                    return v17(true);
                                                };
                                                return v17(true);
                                            };
                                            if (v.value0 instanceof PureScript_Backend_Optimizer_Syntax.OpBooleanOr) {
                                                var $1245 = deref(v.value2);
                                                if ($1245 instanceof NeutLit && ($1245.value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitBoolean && $1245.value0.value0)) {
                                                    $tco_done = true;
                                                    return v.value2;
                                                };
                                                return v15(true);
                                            };
                                            return v15(true);
                                        };
                                        if (v.value0 instanceof PureScript_Backend_Optimizer_Syntax.OpBooleanOr) {
                                            var $1249 = deref(v.value1);
                                            if ($1249 instanceof NeutLit && ($1249.value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitBoolean && $1249.value0.value0)) {
                                                $tco_done = true;
                                                return v.value1;
                                            };
                                            return v13(true);
                                        };
                                        return v13(true);
                                    };
                                    if (v.value0 instanceof PureScript_Backend_Optimizer_Syntax.OpBooleanOr) {
                                        var $1253 = deref(v.value2);
                                        if ($1253 instanceof NeutLit && ($1253.value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitBoolean && !$1253.value0.value0)) {
                                            $tco_done = true;
                                            return v.value1;
                                        };
                                        return v11(true);
                                    };
                                    return v11(true);
                                };
                                if (v.value0 instanceof PureScript_Backend_Optimizer_Syntax.OpBooleanOr) {
                                    var $1257 = deref(v.value1);
                                    if ($1257 instanceof NeutLit && ($1257.value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitBoolean && !$1257.value0.value0)) {
                                        $tco_done = true;
                                        return v.value2;
                                    };
                                    return v9(true);
                                };
                                return v9(true);
                            };
                            if (v.value0 instanceof PureScript_Backend_Optimizer_Syntax.OpBooleanAnd) {
                                var $1261 = deref(v.value2);
                                if ($1261 instanceof NeutLit && ($1261.value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitBoolean && $1261.value0.value0)) {
                                    $tco_done = true;
                                    return v.value1;
                                };
                                return v7(true);
                            };
                            return v7(true);
                        };
                        if (v.value0 instanceof PureScript_Backend_Optimizer_Syntax.OpBooleanAnd) {
                            var $1265 = deref(v.value1);
                            if ($1265 instanceof NeutLit && ($1265.value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitBoolean && $1265.value0.value0)) {
                                $tco_done = true;
                                return v.value2;
                            };
                            return v5(true);
                        };
                        return v5(true);
                    };
                    if (v.value0 instanceof PureScript_Backend_Optimizer_Syntax.OpBooleanAnd) {
                        var $1269 = deref(v.value2);
                        if ($1269 instanceof NeutLit && ($1269.value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitBoolean && !$1269.value0.value0)) {
                            $tco_done = true;
                            return v.value2;
                        };
                        return v3(true);
                    };
                    return v3(true);
                };
                if (v.value0 instanceof PureScript_Backend_Optimizer_Syntax.OpBooleanAnd) {
                    var $1273 = deref(v.value1);
                    if ($1273 instanceof NeutLit && ($1273.value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitBoolean && !$1273.value0.value0)) {
                        $tco_done = true;
                        return v.value1;
                    };
                    return v1(true);
                };
                return v1(true);
            };
            throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Semantics (line 651, column 18 - line 791, column 45): " + [ v.constructor.name ]);
        };
        while (!$tco_done) {
            $tco_result = $tco_loop($tco_var_env, $copy_v);
        };
        return $tco_result;
    };
};
var evalAssocOp$prime = function (v) {
    return function (op) {
        return function (a) {
            return function (b) {
                if (op instanceof Data_Either.Left) {
                    var v1 = v.evalExternSpine(v)(op.value0)([ new ExternApp([ a, b ]) ]);
                    if (v1 instanceof Data_Maybe.Just) {
                        return v1.value0;
                    };
                    if (v1 instanceof Data_Maybe.Nothing) {
                        return new SemAssocOp(op, Data_Array_NonEmpty["cons$prime"](a)([ b ]));
                    };
                    throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Semantics (line 925, column 5 - line 929, column 52): " + [ v1.constructor.name ]);
                };
                if (op instanceof Data_Either.Right) {
                    return evalPrimOp(v)(new PureScript_Backend_Optimizer_Syntax.Op2(op.value0, a, b));
                };
                throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Semantics (line 923, column 35 - line 931, column 36): " + [ op.constructor.name ]);
            };
        };
    };
};
var evalAssocOp = function ($copy_env) {
    return function ($copy_op1) {
        return function ($copy_v) {
            return function ($copy_v1) {
                var $tco_var_env = $copy_env;
                var $tco_var_op1 = $copy_op1;
                var $tco_var_v = $copy_v;
                var $tco_done = false;
                var $tco_result;
                function $tco_loop(env, op1, v, v1) {
                    if (v instanceof SemTyped) {
                        $tco_var_env = env;
                        $tco_var_op1 = op1;
                        $tco_var_v = v.value1;
                        $copy_v1 = v1;
                        return;
                    };
                    if (v1 instanceof SemTyped) {
                        $tco_var_env = env;
                        $tco_var_op1 = op1;
                        $tco_var_v = v;
                        $copy_v1 = v1.value1;
                        return;
                    };
                    var v2 = function (v3) {
                        if (v1 instanceof SemAssocOp && Data_Eq.eq(eqEither)(op1)(v1.value0)) {
                            var v4 = evalAssocOp$prime(env)(op1)(v)(Data_Array_NonEmpty.head(v1.value1));
                            if (v4 instanceof SemAssocOp && Data_Eq.eq(eqEither)(v1.value0)(v4.value0)) {
                                return new SemAssocOp(op1, Data_Array_NonEmpty.appendArray(v4.value1)(Data_Array_NonEmpty.tail(v1.value1)));
                            };
                            return new SemAssocOp(op1, Data_Array_NonEmpty["cons$prime"](v4)(Data_Array_NonEmpty.tail(v1.value1)));
                        };
                        if (v instanceof SemAssocOp && Data_Eq.eq(eqEither)(op1)(v.value0)) {
                            var v4 = evalAssocOp$prime(env)(op1)(Data_Array_NonEmpty.last(v.value1))(v1);
                            if (v4 instanceof SemAssocOp && Data_Eq.eq(eqEither)(v.value0)(v4.value0)) {
                                return new SemAssocOp(op1, Data_Array_NonEmpty.prependArray(Data_Array_NonEmpty.init(v.value1))(v4.value1));
                            };
                            return new SemAssocOp(op1, Data_Array_NonEmpty["snoc$prime"](Data_Array_NonEmpty.init(v.value1))(v4));
                        };
                        return new SemAssocOp(op1, Data_Array_NonEmpty["cons$prime"](v)([ v1 ]));
                    };
                    if (v instanceof SemAssocOp && v1 instanceof SemAssocOp) {
                        var $1308 = Data_Eq.eq(eqEither)(op1)(v.value0);
                        if ($1308) {
                            var $1309 = Data_Eq.eq(eqEither)(v.value0)(v1.value0);
                            if ($1309) {
                                var v3 = evalAssocOp$prime(env)(op1)(Data_Array_NonEmpty.last(v.value1))(Data_Array_NonEmpty.head(v1.value1));
                                if (v3 instanceof SemAssocOp && Data_Eq.eq(eqEither)(v1.value0)(v3.value0)) {
                                    $tco_done = true;
                                    return new SemAssocOp(op1, Data_Array_NonEmpty.prependArray(Data_Array_NonEmpty.init(v.value1))(Data_Array_NonEmpty.appendArray(v3.value1)(Data_Array_NonEmpty.tail(v1.value1))));
                                };
                                $tco_done = true;
                                return new SemAssocOp(op1, Data_Array_NonEmpty.prependArray(Data_Array_NonEmpty.init(v.value1))(Data_Array_NonEmpty["cons$prime"](v3)(Data_Array_NonEmpty.tail(v1.value1))));
                            };
                            $tco_done = true;
                            return v2(true);
                        };
                        $tco_done = true;
                        return v2(true);
                    };
                    $tco_done = true;
                    return v2(true);
                };
                while (!$tco_done) {
                    $tco_result = $tco_loop($tco_var_env, $tco_var_op1, $tco_var_v, $copy_v1);
                };
                return $tco_result;
            };
        };
    };
};
var evalAccessor = function (env) {
    return function (lhs) {
        return function (accessor) {
            return floatLet(lhs)(function (v) {
                if (v instanceof SemTyped) {
                    return evalAccessor(env)(v.value1)(accessor);
                };
                if (v instanceof SemRef) {
                    return evalRef(env)(v.value0)(v.value1)(new ExternAccessor(accessor))(v.value2);
                };
                var v1 = function (v2) {
                    var v3 = function (v4) {
                        var v5 = function (v6) {
                            var v7 = function (v8) {
                                if (v instanceof NeutFail) {
                                    return new NeutFail(v.value0);
                                };
                                return new NeutAccessor(v, accessor);
                            };
                            if (v instanceof NeutData) {
                                if (accessor instanceof PureScript_Backend_Optimizer_Syntax.GetCtorField) {
                                    var $1327 = Data_Array.index(v.value4)(accessor.value5);
                                    if ($1327 instanceof Data_Maybe.Just) {
                                        return $1327.value0.value1;
                                    };
                                    return v7(true);
                                };
                                return v7(true);
                            };
                            return v7(true);
                        };
                        if (v instanceof NeutLit && v.value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitArray) {
                            if (accessor instanceof PureScript_Backend_Optimizer_Syntax.GetIndex) {
                                var $1344 = Data_Array.index(v.value0.value0)(accessor.value0);
                                if ($1344 instanceof Data_Maybe.Just) {
                                    return $1344.value0;
                                };
                                return v5(true);
                            };
                            return v5(true);
                        };
                        return v5(true);
                    };
                    if (v instanceof NeutUpdate) {
                        if (accessor instanceof PureScript_Backend_Optimizer_Syntax.GetProp) {
                            var v4 = Data_Array.findMap(function (v5) {
                                return Data_Functor.voidLeft(Data_Maybe.functorMaybe)(Control_Alternative.guard(Data_Maybe.alternativeMaybe)(v5.value0 === accessor.value0))(v5.value1);
                            })(v.value1);
                            if (v4 instanceof Data_Maybe.Just) {
                                return v4.value0;
                            };
                            if (v4 instanceof Data_Maybe.Nothing) {
                                return evalAccessor(env)(v.value0)(accessor);
                            };
                            throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Semantics (line 520, column 9 - line 524, column 42): " + [ v4.constructor.name ]);
                        };
                        return v3(true);
                    };
                    return v3(true);
                };
                if (v instanceof NeutLit && v.value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitRecord) {
                    if (accessor instanceof PureScript_Backend_Optimizer_Syntax.GetProp) {
                        var $1364 = Data_Array.findMap(function (v2) {
                            return Data_Functor.voidLeft(Data_Maybe.functorMaybe)(Control_Alternative.guard(Data_Maybe.alternativeMaybe)(v2.value0 === accessor.value0))(v2.value1);
                        })(v.value0.value0);
                        if ($1364 instanceof Data_Maybe.Just) {
                            return $1364.value0;
                        };
                        return v1(true);
                    };
                    return v1(true);
                };
                return v1(true);
            });
        };
    };
};
var caseString = function (v) {
    if (v instanceof NeutLit && v.value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitString) {
        return new Data_Maybe.Just(v.value0.value0);
    };
    return Data_Maybe.Nothing.value;
};
var caseNumber = function (v) {
    if (v instanceof NeutLit && v.value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitNumber) {
        return new Data_Maybe.Just(v.value0.value0);
    };
    return Data_Maybe.Nothing.value;
};
var caseInt = function (v) {
    if (v instanceof NeutLit && v.value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitInt) {
        return new Data_Maybe.Just(v.value0.value0);
    };
    return Data_Maybe.Nothing.value;
};
var buildStop = function (v) {
    return function (stop) {
        return new ExprRewrite(v.analyze(v)(new PureScript_Backend_Optimizer_Syntax.Var(stop)), new RewriteStop(stop));
    };
};
var buildDefault = function (v) {
    return function (expr) {
        return new ExprSyntax(v.analyze(v)(expr), expr);
    };
};
var build = function (ctx) {
    return function (v) {
        if (v instanceof PureScript_Backend_Optimizer_Syntax.App && (v.value0 instanceof ExprSyntax && v.value0.value1 instanceof PureScript_Backend_Optimizer_Syntax.App)) {
            return build(ctx)(new PureScript_Backend_Optimizer_Syntax.App(v.value0.value1.value0, Data_Semigroup.append(Data_Array_NonEmpty_Internal.semigroupNonEmptyArray)(v.value0.value1.value1)(v.value1)));
        };
        if (v instanceof PureScript_Backend_Optimizer_Syntax.Abs && (v.value1 instanceof ExprSyntax && v.value1.value1 instanceof PureScript_Backend_Optimizer_Syntax.Abs)) {
            return build(ctx)(new PureScript_Backend_Optimizer_Syntax.Abs(Data_Semigroup.append(Data_Array_NonEmpty_Internal.semigroupNonEmptyArray)(v.value0)(v.value1.value1.value0), v.value1.value1.value1));
        };
        if (v instanceof PureScript_Backend_Optimizer_Syntax.Let && shouldInlineLet(v.value1)(v.value2)(v.value3)) {
            return rewriteInline(v.value0)(v.value1)(v.value2)(v.value3);
        };
        var v1 = function (v2) {
            var v3 = function (v4) {
                var v5 = function (v6) {
                    var v7 = function (v8) {
                        var v9 = function (v10) {
                            var v11 = function (v12) {
                                var v13 = function (v14) {
                                    var v15 = function (v16) {
                                        var v17 = function (v18) {
                                            var v19 = function (v20) {
                                                var v21 = function (v22) {
                                                    var v23 = function (v24) {
                                                        var v25 = function (v26) {
                                                            if (v instanceof PureScript_Backend_Optimizer_Syntax.EffectBind && (v.value2 instanceof ExprSyntax && v.value2.value1 instanceof PureScript_Backend_Optimizer_Syntax.EffectPure)) {
                                                                return build(ctx)(new PureScript_Backend_Optimizer_Syntax.EffectDefer(build(ctx)(new PureScript_Backend_Optimizer_Syntax.Let(v.value0, v.value1, v.value2.value1.value0, v.value3))));
                                                            };
                                                            if (v instanceof PureScript_Backend_Optimizer_Syntax.EffectBind && (v.value2 instanceof ExprSyntax && v.value2.value1 instanceof PureScript_Backend_Optimizer_Syntax.EffectDefer)) {
                                                                return build(ctx)(new PureScript_Backend_Optimizer_Syntax.EffectBind(v.value0, v.value1, v.value2.value1.value0, v.value3));
                                                            };
                                                            if (v instanceof PureScript_Backend_Optimizer_Syntax.EffectBind && (v.value3 instanceof ExprSyntax && v.value3.value1 instanceof PureScript_Backend_Optimizer_Syntax.EffectDefer)) {
                                                                return build(ctx)(new PureScript_Backend_Optimizer_Syntax.EffectBind(v.value0, v.value1, v.value2, v.value3.value1.value0));
                                                            };
                                                            if (v instanceof PureScript_Backend_Optimizer_Syntax.EffectBind && (v.value3 instanceof ExprSyntax && (v.value3.value1 instanceof PureScript_Backend_Optimizer_Syntax.EffectPure && (v.value3.value1.value0 instanceof ExprSyntax && (v.value3.value1.value0.value1 instanceof PureScript_Backend_Optimizer_Syntax.Local && Data_Eq.eq(PureScript_Backend_Optimizer_Syntax.eqLevel)(v.value1)(v.value3.value1.value0.value1.value1)))))) {
                                                                return v.value2;
                                                            };
                                                            if (v instanceof PureScript_Backend_Optimizer_Syntax.EffectDefer && (v.value0 instanceof ExprSyntax && v.value0.value1 instanceof PureScript_Backend_Optimizer_Syntax.EffectDefer)) {
                                                                return v.value0;
                                                            };
                                                            if (v instanceof PureScript_Backend_Optimizer_Syntax.PrimOp && (v.value0 instanceof PureScript_Backend_Optimizer_Syntax.Op1 && (v.value0.value0 instanceof PureScript_Backend_Optimizer_Syntax.OpBooleanNot && (v.value0.value1 instanceof ExprSyntax && (v.value0.value1.value1 instanceof PureScript_Backend_Optimizer_Syntax.PrimOp && (v.value0.value1.value1.value0 instanceof PureScript_Backend_Optimizer_Syntax.Op1 && v.value0.value1.value1.value0.value0 instanceof PureScript_Backend_Optimizer_Syntax.OpBooleanNot)))))) {
                                                                return v.value0.value1.value1.value0.value1;
                                                            };
                                                            return buildDefault(ctx)(v);
                                                        };
                                                        if (v instanceof PureScript_Backend_Optimizer_Syntax.PrimOp && (v.value0 instanceof PureScript_Backend_Optimizer_Syntax.Op2 && (v.value0.value2 instanceof ExprSyntax && v.value0.value2.value1 instanceof PureScript_Backend_Optimizer_Syntax.Branch))) {
                                                            var $1447 = shouldDistributeBranchPrimOp2R(v.value0.value2.value0)(v.value0.value2.value1.value0)(v.value0.value2.value1.value1)(v.value0.value1)(v.value0.value0);
                                                            if ($1447 instanceof Data_Maybe.Just) {
                                                                return $1447.value0;
                                                            };
                                                            return v25(true);
                                                        };
                                                        return v25(true);
                                                    };
                                                    if (v instanceof PureScript_Backend_Optimizer_Syntax.PrimOp && (v.value0 instanceof PureScript_Backend_Optimizer_Syntax.Op2 && (v.value0.value1 instanceof ExprSyntax && v.value0.value1.value1 instanceof PureScript_Backend_Optimizer_Syntax.Branch))) {
                                                        var $1458 = shouldDistributeBranchPrimOp2L(v.value0.value1.value0)(v.value0.value1.value1.value0)(v.value0.value1.value1.value1)(v.value0.value0)(v.value0.value2);
                                                        if ($1458 instanceof Data_Maybe.Just) {
                                                            return $1458.value0;
                                                        };
                                                        return v23(true);
                                                    };
                                                    return v23(true);
                                                };
                                                if (v instanceof PureScript_Backend_Optimizer_Syntax.PrimOp && (v.value0 instanceof PureScript_Backend_Optimizer_Syntax.Op1 && (v.value0.value1 instanceof ExprSyntax && v.value0.value1.value1 instanceof PureScript_Backend_Optimizer_Syntax.Branch))) {
                                                    var $1469 = shouldDistributeBranchPrimOp1(v.value0.value1.value0)(v.value0.value1.value1.value0)(v.value0.value1.value1.value1)(v.value0.value0);
                                                    if ($1469 instanceof Data_Maybe.Just) {
                                                        return $1469.value0;
                                                    };
                                                    return v21(true);
                                                };
                                                return v21(true);
                                            };
                                            if (v instanceof PureScript_Backend_Optimizer_Syntax.Accessor && (v.value0 instanceof ExprSyntax && v.value0.value1 instanceof PureScript_Backend_Optimizer_Syntax.Branch)) {
                                                var $1479 = shouldDistributeBranchAccessor(v.value0.value0)(v.value0.value1.value0)(v.value0.value1.value1)(v.value1);
                                                if ($1479 instanceof Data_Maybe.Just) {
                                                    return $1479.value0;
                                                };
                                                return v19(true);
                                            };
                                            return v19(true);
                                        };
                                        if (v instanceof PureScript_Backend_Optimizer_Syntax.UncurriedApp && (v.value0 instanceof ExprSyntax && v.value0.value1 instanceof PureScript_Backend_Optimizer_Syntax.Branch)) {
                                            var $1488 = shouldDistributeBranchUncurriedApps(v.value0.value0)(v.value0.value1.value0)(v.value0.value1.value1)(v.value1);
                                            if ($1488 instanceof Data_Maybe.Just) {
                                                return $1488.value0;
                                            };
                                            return v17(true);
                                        };
                                        return v17(true);
                                    };
                                    if (v instanceof PureScript_Backend_Optimizer_Syntax.App && (v.value0 instanceof ExprSyntax && v.value0.value1 instanceof PureScript_Backend_Optimizer_Syntax.Branch)) {
                                        var $1497 = shouldDistributeBranchApps(v.value0.value0)(v.value0.value1.value0)(v.value0.value1.value1)(v.value1);
                                        if ($1497 instanceof Data_Maybe.Just) {
                                            return $1497.value0;
                                        };
                                        return v15(true);
                                    };
                                    return v15(true);
                                };
                                if (v instanceof PureScript_Backend_Optimizer_Syntax.Let) {
                                    var $1506 = shouldEtaReduce(v.value1)(v.value2)(v.value3);
                                    if ($1506 instanceof Data_Maybe.Just) {
                                        return $1506.value0;
                                    };
                                    return v13(true);
                                };
                                return v13(true);
                            };
                            if (v instanceof PureScript_Backend_Optimizer_Syntax.Let) {
                                var $1513 = shouldDistributeBranches(v.value0)(v.value1)(v.value2)(v.value3);
                                if ($1513 instanceof Data_Maybe.Just) {
                                    return $1513.value0;
                                };
                                return v11(true);
                            };
                            return v11(true);
                        };
                        if (v instanceof PureScript_Backend_Optimizer_Syntax.Let) {
                            var $1520 = shouldUnpackArray(v.value0)(v.value1)(v.value2)(v.value3);
                            if ($1520 instanceof Data_Maybe.Just) {
                                return $1520.value0;
                            };
                            return v9(true);
                        };
                        return v9(true);
                    };
                    if (v instanceof PureScript_Backend_Optimizer_Syntax.Let) {
                        var $1527 = shouldUnpackCtor(v.value0)(v.value1)(v.value2)(v.value3);
                        if ($1527 instanceof Data_Maybe.Just) {
                            return $1527.value0;
                        };
                        return v7(true);
                    };
                    return v7(true);
                };
                if (v instanceof PureScript_Backend_Optimizer_Syntax.Let) {
                    var $1534 = shouldUnpackUpdate(v.value0)(v.value1)(v.value2)(v.value3);
                    if ($1534 instanceof Data_Maybe.Just) {
                        return $1534.value0;
                    };
                    return v5(true);
                };
                return v5(true);
            };
            if (v instanceof PureScript_Backend_Optimizer_Syntax.Let) {
                var $1541 = shouldUnpackRecord(v.value0)(v.value1)(v.value2)(v.value3);
                if ($1541 instanceof Data_Maybe.Just) {
                    return $1541.value0;
                };
                return v3(true);
            };
            return v3(true);
        };
        if (v instanceof PureScript_Backend_Optimizer_Syntax.Let) {
            var $1548 = shouldUncurryAbs(v.value0)(v.value1)(v.value2)(v.value3);
            if ($1548 instanceof Data_Maybe.Just) {
                return $1548.value0;
            };
            return v1(true);
        };
        return v1(true);
    };
};
var simplifyCondBoolean = function (ctx) {
    return function (v) {
        return function (v1) {
            if (v.value1 instanceof ExprSyntax && (v.value1.value1 instanceof PureScript_Backend_Optimizer_Syntax.Lit && (v.value1.value1.value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitBoolean && (v1 instanceof ExprSyntax && (v1.value1 instanceof PureScript_Backend_Optimizer_Syntax.Lit && v1.value1.value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitBoolean))))) {
                if (v.value1.value1.value0.value0 === v1.value1.value0.value0) {
                    return new Data_Maybe.Just(v.value1);
                };
                if (v.value1.value1.value0.value0 && !v1.value1.value0.value0) {
                    return new Data_Maybe.Just(v.value0);
                };
                if (!v.value1.value1.value0.value0 && v1.value1.value0.value0) {
                    return new Data_Maybe.Just(build(ctx)(new PureScript_Backend_Optimizer_Syntax.PrimOp(new PureScript_Backend_Optimizer_Syntax.Op1(PureScript_Backend_Optimizer_Syntax.OpBooleanNot.value, v.value0))));
                };
            };
            if (v.value1 instanceof ExprSyntax && (v.value1.value1 instanceof PureScript_Backend_Optimizer_Syntax.Lit && (v.value1.value1.value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitBoolean && (v.value1.value1.value0.value0 && isSimplePredicate(v1))))) {
                return new Data_Maybe.Just(build(ctx)(new PureScript_Backend_Optimizer_Syntax.PrimOp(new PureScript_Backend_Optimizer_Syntax.Op2(PureScript_Backend_Optimizer_Syntax.OpBooleanOr.value, v.value0, v1))));
            };
            if (v1 instanceof ExprSyntax && (v1.value1 instanceof PureScript_Backend_Optimizer_Syntax.Lit && (v1.value1.value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitBoolean && !v1.value1.value0.value0))) {
                return new Data_Maybe.Just(build(ctx)(new PureScript_Backend_Optimizer_Syntax.PrimOp(new PureScript_Backend_Optimizer_Syntax.Op2(PureScript_Backend_Optimizer_Syntax.OpBooleanAnd.value, v.value0, v.value1))));
            };
            return Data_Maybe.Nothing.value;
        };
    };
};
var simplifyCondRedundantElse = function (ctx) {
    return function (v) {
        return function (v1) {
            var v2 = function (v3) {
                return Data_Maybe.Nothing.value;
            };
            if (v1 instanceof ExprSyntax && v1.value1 instanceof PureScript_Backend_Optimizer_Syntax.Branch) {
                var $1580 = Data_Array_NonEmpty.head(v1.value1.value0);
                if ($1580.value0 instanceof ExprSyntax && ($1580.value0.value1 instanceof PureScript_Backend_Optimizer_Syntax.PrimOp && ($1580.value0.value1.value0 instanceof PureScript_Backend_Optimizer_Syntax.Op1 && $1580.value0.value1.value0.value0 instanceof PureScript_Backend_Optimizer_Syntax.OpBooleanNot))) {
                    var $1581 = Data_Eq.eq(eqBackendExpr)(v.value0)($1580.value0.value1.value0.value1);
                    if ($1581) {
                        return new Data_Maybe.Just(buildBranchCond(ctx)(new PureScript_Backend_Optimizer_Syntax.Pair(v.value0, v.value1))($1580.value1));
                    };
                    return v2(true);
                };
                return v2(true);
            };
            return v2(true);
        };
    };
};
var simplifyCondLiftAnd = function (ctx) {
    return function (pair) {
        return function (def1) {
            var v = function (v1) {
                return Data_Maybe.Nothing.value;
            };
            if (pair.value1 instanceof ExprSyntax && pair.value1.value1 instanceof PureScript_Backend_Optimizer_Syntax.Branch) {
                var $1596 = Data_Array_NonEmpty.toArray(pair.value1.value1.value0);
                if ($1596.length === 1) {
                    var $1597 = Data_Eq.eq(eqBackendExpr)(def1)(pair.value1.value1.value1);
                    if ($1597) {
                        return new Data_Maybe.Just(buildBranchCond(ctx)(new PureScript_Backend_Optimizer_Syntax.Pair(build(ctx)(new PureScript_Backend_Optimizer_Syntax.PrimOp(new PureScript_Backend_Optimizer_Syntax.Op2(PureScript_Backend_Optimizer_Syntax.OpBooleanAnd.value, pair.value0, $1596[0].value0))), $1596[0].value1))(def1));
                    };
                    return v(true);
                };
                return v(true);
            };
            return v(true);
        };
    };
};
var buildBranchCond = function (ctx) {
    return function (pair) {
        return function (def) {
            var v = function (v1) {
                var v2 = function (v3) {
                    var v4 = function (v5) {
                        var v6 = function (v7) {
                            var v8 = function (v9) {
                                if (Data_Boolean.otherwise) {
                                    return build(ctx)(new PureScript_Backend_Optimizer_Syntax.Branch(Data_Array_NonEmpty.singleton(pair), def));
                                };
                                throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Semantics (line 1421, column 1 - line 1421, column 89): " + [ ctx.constructor.name, pair.constructor.name, def.constructor.name ]);
                            };
                            if (def instanceof ExprSyntax && def.value1 instanceof PureScript_Backend_Optimizer_Syntax.Branch) {
                                return build(ctx)(new PureScript_Backend_Optimizer_Syntax.Branch(Data_Array_NonEmpty.cons(pair)(def.value1.value0), def.value1.value1));
                            };
                            return v8(true);
                        };
                        var $1621 = simplifyCondRedundantElse(ctx)(pair)(def);
                        if ($1621 instanceof Data_Maybe.Just) {
                            return $1621.value0;
                        };
                        return v6(true);
                    };
                    var $1626 = simplifyCondLiftAnd(ctx)(pair)(def);
                    if ($1626 instanceof Data_Maybe.Just) {
                        return $1626.value0;
                    };
                    return v4(true);
                };
                var $1631 = simplifyCondBoolean(ctx)(pair)(def);
                if ($1631 instanceof Data_Maybe.Just) {
                    return $1631.value0;
                };
                return v2(true);
            };
            var $1636 = simplifyCondIsTag(ctx)(pair)(def);
            if ($1636 instanceof Data_Maybe.Just) {
                return $1636.value0;
            };
            return v(true);
        };
    };
};
var $lazy_quote = /* #__PURE__ */ $runtime_lazy("quote", "PureScript.Backend.Optimizer.Semantics", function () {
    var go = function (ctx) {
        return function (v) {
            if (v instanceof SemTyped) {
                return build(ctx)(new PureScript_Backend_Optimizer_Syntax.Typed(v.value0, go(ctx)(v.value1)));
            };
            if (v instanceof SemLet) {
                var v1 = nextLevel(ctx);
                return build(ctx)(new PureScript_Backend_Optimizer_Syntax.Let(v.value0, v1.value0, $lazy_quote(1209)(purely(ctx))(v.value1), $lazy_quote(1209)(v1.value1)(v.value2(new SemRef(new EvalLocal(v.value0, v1.value0), [  ], Data_Lazy.defer(function (v2) {
                    return deref(v.value1);
                }))))));
            };
            if (v instanceof SemLetRec) {
                var v1 = nextLevel(ctx);
                var neutBindings = Data_Functor.map(Data_Array_NonEmpty_Internal.functorNonEmptyArray)(function (v2) {
                    return new Data_Tuple.Tuple(v2.value0, Data_Lazy.defer(function (v3) {
                        return new NeutLocal(new Data_Maybe.Just(v2.value0), v1.value0);
                    }));
                })(v.value0);
                return build(ctx)(new PureScript_Backend_Optimizer_Syntax.LetRec(v1.value0, Data_Functor.map(Data_Array_NonEmpty_Internal.functorNonEmptyArray)(Data_Functor.map(Data_Tuple.functorTuple)(function (b) {
                    return $lazy_quote(1218)(purely(v1.value1))(b(neutBindings));
                }))(v.value0), $lazy_quote(1219)(v1.value1)(v.value1(neutBindings))));
            };
            if (v instanceof SemEffectBind) {
                var ctx$prime = effectfully(ctx);
                var v1 = nextLevel(ctx$prime);
                return build(ctx)(new PureScript_Backend_Optimizer_Syntax.EffectBind(v.value0, v1.value0, $lazy_quote(1223)(ctx$prime)(v.value1), $lazy_quote(1223)(v1.value1)(v.value2(new NeutLocal(v.value0, v1.value0)))));
            };
            if (v instanceof SemEffectPure) {
                return build(ctx)(new PureScript_Backend_Optimizer_Syntax.EffectPure($lazy_quote(1225)(purely(ctx))(v.value0)));
            };
            if (v instanceof SemEffectDefer) {
                return build(ctx)(new PureScript_Backend_Optimizer_Syntax.EffectDefer($lazy_quote(1227)(effectfully(ctx))(v.value0)));
            };
            if (v instanceof SemBranch) {
                var ctx$prime = purely(ctx);
                var quoteCond = function (v1) {
                    return new PureScript_Backend_Optimizer_Syntax.Pair($lazy_quote(1230)(ctx$prime)(Data_Lazy.force(v1.value0)), $lazy_quote(1230)(ctx)(Data_Lazy.force(v1.value1)));
                };
                var branches$prime = Data_Functor.map(Data_Array_NonEmpty_Internal.functorNonEmptyArray)(quoteCond)(v.value0);
                return Data_Foldable.foldr(Data_Array_NonEmpty_Internal.foldableNonEmptyArray)(buildBranchCond(ctx))($lazy_quote(1232)(ctx)(Data_Lazy.force(v.value1)))(branches$prime);
            };
            if (v instanceof SemRef) {
                if (v.value0 instanceof EvalExtern) {
                    return go(ctx)(neutralSpine(new NeutVar(v.value0.value0))(v.value1));
                };
                if (v.value0 instanceof EvalLocal) {
                    return go(ctx)(neutralSpine(new NeutLocal(v.value0.value0, v.value0.value1))(v.value1));
                };
                throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Semantics (line 1236, column 7 - line 1240, column 57): " + [ v.value0.constructor.name ]);
            };
            if (v instanceof SemLam) {
                var v1 = nextLevel(ctx);
                return build(ctx)(new PureScript_Backend_Optimizer_Syntax.Abs(Data_Array_NonEmpty.singleton(new Data_Tuple.Tuple(v.value0, v1.value0)), $lazy_quote(1243)(purely(v1.value1))(v.value1(new NeutLocal(v.value0, v1.value0)))));
            };
            if (v instanceof SemMkFn) {
                var loop = function ($copy_ctx$prime) {
                    return function ($copy_idents) {
                        return function ($copy_v1) {
                            var $tco_var_ctx$prime = $copy_ctx$prime;
                            var $tco_var_idents = $copy_idents;
                            var $tco_done = false;
                            var $tco_result;
                            function $tco_loop(ctx$prime, idents, v1) {
                                if (v1 instanceof MkFnNext) {
                                    var v2 = nextLevel(ctx$prime);
                                    $tco_var_ctx$prime = v2.value1;
                                    $tco_var_idents = Data_Array.snoc(idents)(new Data_Tuple.Tuple(v1.value0, v2.value0));
                                    $copy_v1 = v1.value1(new NeutLocal(v1.value0, v2.value0));
                                    return;
                                };
                                if (v1 instanceof MkFnApplied) {
                                    $tco_done = true;
                                    return build(ctx$prime)(new PureScript_Backend_Optimizer_Syntax.UncurriedAbs(idents, $lazy_quote(1251)(purely(ctx$prime))(v1.value0)));
                                };
                                throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Semantics (line 1246, column 28 - line 1251, column 72): " + [ v1.constructor.name ]);
                            };
                            while (!$tco_done) {
                                $tco_result = $tco_loop($tco_var_ctx$prime, $tco_var_idents, $copy_v1);
                            };
                            return $tco_result;
                        };
                    };
                };
                return loop(ctx)([  ])(v.value0);
            };
            if (v instanceof SemMkEffectFn) {
                var loop = function ($copy_ctx$prime) {
                    return function ($copy_idents) {
                        return function ($copy_v1) {
                            var $tco_var_ctx$prime = $copy_ctx$prime;
                            var $tco_var_idents = $copy_idents;
                            var $tco_done1 = false;
                            var $tco_result;
                            function $tco_loop(ctx$prime, idents, v1) {
                                if (v1 instanceof MkFnNext) {
                                    var v2 = nextLevel(ctx$prime);
                                    $tco_var_ctx$prime = v2.value1;
                                    $tco_var_idents = Data_Array.snoc(idents)(new Data_Tuple.Tuple(v1.value0, v2.value0));
                                    $copy_v1 = v1.value1(new NeutLocal(v1.value0, v2.value0));
                                    return;
                                };
                                if (v1 instanceof MkFnApplied) {
                                    $tco_done1 = true;
                                    return build(ctx$prime)(new PureScript_Backend_Optimizer_Syntax.UncurriedEffectAbs(idents, $lazy_quote(1260)(purely(ctx$prime))(v1.value0)));
                                };
                                throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Semantics (line 1255, column 28 - line 1260, column 78): " + [ v1.constructor.name ]);
                            };
                            while (!$tco_done1) {
                                $tco_result = $tco_loop($tco_var_ctx$prime, $tco_var_idents, $copy_v1);
                            };
                            return $tco_result;
                        };
                    };
                };
                return loop(ctx)([  ])(v.value0);
            };
            if (v instanceof SemAssocOp) {
                return PureScript_Backend_Optimizer_Utils.foldl1Array(function (a) {
                    return function (b) {
                        if (v.value0 instanceof Data_Either.Left) {
                            return build(ctx)(new PureScript_Backend_Optimizer_Syntax.App(build(ctx)(new PureScript_Backend_Optimizer_Syntax.Var(v.value0.value0)), Data_Array_NonEmpty["cons$prime"](a)([ $lazy_quote(1266)(ctx)(b) ])));
                        };
                        if (v.value0 instanceof Data_Either.Right) {
                            return build(ctx)(new PureScript_Backend_Optimizer_Syntax.PrimOp(new PureScript_Backend_Optimizer_Syntax.Op2(v.value0.value0, a, $lazy_quote(1268)(ctx)(b))));
                        };
                        throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Semantics (line 1264, column 19 - line 1268, column 62): " + [ v.value0.constructor.name ]);
                    };
                })($lazy_quote(1270)(ctx))(v.value1);
            };
            if (v instanceof NeutLocal) {
                return build(ctx)(new PureScript_Backend_Optimizer_Syntax.Local(v.value0, v.value1));
            };
            if (v instanceof NeutVar) {
                return build(ctx)(new PureScript_Backend_Optimizer_Syntax.Var(v.value0));
            };
            if (v instanceof NeutStop) {
                return buildStop(ctx)(v.value0);
            };
            if (v instanceof NeutData) {
                return build(ctx)(new PureScript_Backend_Optimizer_Syntax.CtorSaturated(v.value0, v.value1, v.value2, v.value3, Data_Functor.map(Data_Functor.functorArray)(Data_Functor.map(Data_Tuple.functorTuple)($lazy_quote(1279)(ctx)))(v.value4)));
            };
            if (v instanceof NeutCtorDef) {
                return build(ctx)(new PureScript_Backend_Optimizer_Syntax.CtorDef(v.value1, v.value2, v.value3, v.value4));
            };
            if (v instanceof NeutUncurriedApp) {
                var ctx$prime = purely(ctx);
                var hd$prime = $lazy_quote(1284)(ctx$prime)(v.value0);
                return build(ctx)(new PureScript_Backend_Optimizer_Syntax.UncurriedApp(hd$prime, Data_Functor.map(Data_Functor.functorArray)($lazy_quote(1285)(ctx$prime))(v.value1)));
            };
            if (v instanceof NeutUncurriedEffectApp) {
                var ctx$prime = purely(ctx);
                var hd$prime = $lazy_quote(1288)(ctx$prime)(v.value0);
                return build(ctx)(new PureScript_Backend_Optimizer_Syntax.UncurriedEffectApp(hd$prime, Data_Functor.map(Data_Functor.functorArray)($lazy_quote(1289)(ctx$prime))(v.value1)));
            };
            if (v instanceof NeutApp) {
                var ctx$prime = purely(ctx);
                var hd$prime = $lazy_quote(1292)(ctx$prime)(v.value0);
                var v1 = Data_Array_NonEmpty.fromArray(Data_Functor.map(Data_Functor.functorArray)($lazy_quote(1293)(ctx$prime))(v.value1));
                if (v1 instanceof Data_Maybe.Nothing) {
                    return hd$prime;
                };
                if (v1 instanceof Data_Maybe.Just) {
                    return build(ctx)(new PureScript_Backend_Optimizer_Syntax.App(hd$prime, v1.value0));
                };
                throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Semantics (line 1293, column 7 - line 1297, column 35): " + [ v1.constructor.name ]);
            };
            if (v instanceof NeutAccessor) {
                return build(ctx)(new PureScript_Backend_Optimizer_Syntax.Accessor($lazy_quote(1299)(ctx)(v.value0), v.value1));
            };
            if (v instanceof NeutUpdate) {
                return build(ctx)(new PureScript_Backend_Optimizer_Syntax.Update($lazy_quote(1301)(ctx)(v.value0), Data_Functor.map(Data_Functor.functorArray)(Data_Functor.map(PureScript_Backend_Optimizer_CoreFn.functorProp)($lazy_quote(1301)(ctx)))(v.value1)));
            };
            if (v instanceof NeutLit) {
                return build(ctx)(new PureScript_Backend_Optimizer_Syntax.Lit(Data_Functor.map(PureScript_Backend_Optimizer_CoreFn.functorLiteral)($lazy_quote(1303)(ctx))(v.value0)));
            };
            if (v instanceof NeutPrimOp) {
                return build(ctx)(new PureScript_Backend_Optimizer_Syntax.PrimOp(Data_Functor.map(PureScript_Backend_Optimizer_Syntax.functorBackendOperator)($lazy_quote(1305)(ctx))(v.value0)));
            };
            if (v instanceof NeutPrimEffect) {
                return build(ctx)(new PureScript_Backend_Optimizer_Syntax.PrimEffect(Data_Functor.map(PureScript_Backend_Optimizer_Syntax.functorBackendEffect)($lazy_quote(1307)(purely(ctx)))(v.value0)));
            };
            if (v instanceof NeutPrimUndefined) {
                return build(ctx)(PureScript_Backend_Optimizer_Syntax.PrimUndefined.value);
            };
            if (v instanceof NeutFail) {
                return build(ctx)(new PureScript_Backend_Optimizer_Syntax.Fail(v.value0));
            };
            throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Semantics (line 1203, column 12 - line 1311, column 27): " + [ v.constructor.name ]);
        };
    };
    return go;
});
var quote = /* #__PURE__ */ $lazy_quote(1200);
var bindLocal = function (v) {
    return function (sem) {
        return {
            currentModule: v.currentModule,
            evalExternRef: v.evalExternRef,
            evalExternSpine: v.evalExternSpine,
            directives: v.directives,
            locals: Data_Array.snoc(v.locals)(sem)
        };
    };
};
var evalApp = function (env) {
    return function (hd) {
        return function (spine) {
            var go = function (env$prime) {
                return function (v) {
                    return function (v1) {
                        if (v instanceof SemTyped) {
                            return go(env$prime)(v.value1)(v1);
                        };
                        if (v1 instanceof Data_List_Types.Cons && v1.value0 instanceof SemTyped) {
                            return go(env$prime)(v)(new Data_List_Types.Cons(v1.value0.value1, v1.value1));
                        };
                        if (v1 instanceof Data_List_Types.Cons && v1.value0 instanceof NeutFail) {
                            return new NeutFail(v1.value0.value0);
                        };
                        if (v instanceof NeutFail) {
                            return new NeutFail(v.value0);
                        };
                        if (v instanceof SemLam && v1 instanceof Data_List_Types.Cons) {
                            return makeLet(Data_Maybe.Nothing.value)(v1.value0)(function (nextArg) {
                                return go(env$prime)(v.value1(nextArg))(v1.value1);
                            });
                        };
                        if (v instanceof SemRef && v1 instanceof Data_List_Types.Cons) {
                            return go(env$prime)(evalRef(env$prime)(v.value0)(v.value1)(new ExternApp([ v1.value0 ]))(v.value2))(v1.value1);
                        };
                        if (v instanceof SemLet) {
                            return new SemLet(v.value0, v.value1, function (nextVal) {
                                return makeLet(Data_Maybe.Nothing.value)(v.value2(nextVal))(function (nextFn) {
                                    return go(bindLocal(bindLocal(env$prime)(new One(nextVal)))(new One(nextFn)))(nextFn)(v1);
                                });
                            });
                        };
                        if (v instanceof SemLetRec) {
                            return new SemLetRec(v.value0, function (nextVals) {
                                return makeLet(Data_Maybe.Nothing.value)(v.value1(nextVals))(function (nextFn) {
                                    return go(bindLocal(bindLocal(env$prime)(new Group(nextVals)))(new One(nextFn)))(nextFn)(v1);
                                });
                            });
                        };
                        if (v instanceof NeutCtorDef && Data_Array.length(v.value4) === Data_List.length(v1)) {
                            return Partial_Unsafe.unsafeCrashWith("CRASH CtorDef");
                        };
                        if (v1 instanceof Data_List_Types.Nil) {
                            return v;
                        };
                        return new NeutApp(v, Data_List.toUnfoldable(Data_Unfoldable.unfoldableArray)(v1));
                    };
                };
            };
            return go(env)(hd)(Data_List.fromFoldable(Data_Foldable.foldableArray)(spine));
        };
    };
};
var evalMkFn = function (env) {
    return function (n) {
        return function (sem) {
            if (n === 0) {
                return new MkFnApplied(sem);
            };
            if (Data_Boolean.otherwise) {
                if (sem instanceof SemLam) {
                    return new MkFnNext(sem.value0, (function () {
                        var $2101 = evalMkFn(env)(n - 1 | 0);
                        return function ($2102) {
                            return $2101(sem.value1($2102));
                        };
                    })());
                };
                return new MkFnNext(Data_Maybe.Nothing.value, function (nextArg) {
                    var env$prime = bindLocal(env)(new One(nextArg));
                    return evalMkFn(env$prime)(n - 1 | 0)(evalApp(env$prime)(sem)([ nextArg ]));
                });
            };
            throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Semantics (line 1769, column 1 - line 1769, column 68): " + [ env.constructor.name, n.constructor.name, sem.constructor.name ]);
        };
    };
};
var evalUncurriedApp = function (env) {
    return function (hd) {
        return function (spine) {
            if (hd instanceof SemTyped) {
                return evalUncurriedApp(env)(hd.value1)(spine);
            };
            if (hd instanceof SemMkFn) {
                return evalUncurriedBeta(NeutUncurriedApp.create)(hd.value0)(spine);
            };
            if (hd instanceof SemRef) {
                return guardFailOver(Data_Foldable.foldableArray)(identity)(spine)(function (spine$prime) {
                    return evalRef(env)(hd.value0)(hd.value1)(new ExternUncurriedApp(spine$prime))(hd.value2);
                });
            };
            if (hd instanceof SemLet) {
                return new SemLet(hd.value0, hd.value1, function (nextVal) {
                    return makeLet(Data_Maybe.Nothing.value)(hd.value2(nextVal))(function (nextFn) {
                        return evalUncurriedApp(bindLocal(bindLocal(env)(new One(nextVal)))(new One(nextFn)))(nextFn)(spine);
                    });
                });
            };
            if (hd instanceof NeutFail) {
                return new NeutFail(hd.value0);
            };
            return guardFailOver(Data_Foldable.foldableArray)(identity)(spine)(NeutUncurriedApp.create(hd));
        };
    };
};
var evalSpine = function (env) {
    var go = function (hd) {
        return function (v) {
            if (v instanceof ExternApp) {
                return evalApp(env)(hd)(v.value0);
            };
            if (v instanceof ExternUncurriedApp) {
                return evalUncurriedApp(env)(hd)(v.value0);
            };
            if (v instanceof ExternAccessor) {
                return evalAccessor(env)(hd)(v.value0);
            };
            if (v instanceof ExternPrimOp) {
                return evalPrimOp(env)(new PureScript_Backend_Optimizer_Syntax.Op1(v.value0, hd));
            };
            throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Semantics (line 475, column 11 - line 483, column 34): " + [ v.constructor.name ]);
        };
    };
    return Data_Foldable.foldl(Data_Foldable.foldableArray)(go);
};
var mkUncurriedAppRewrite = function (env) {
    return function (hd) {
        var go = function (acc) {
            return function (n) {
                if (n === 0) {
                    return evalUncurriedApp(env)(hd)(acc);
                };
                if (Data_Boolean.otherwise) {
                    return new SemLam(Data_Maybe.Nothing.value, function (arg) {
                        return go(Data_Array.snoc(acc)(arg))(n - 1 | 0);
                    });
                };
                throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Semantics (line 1784, column 3 - line 1788, column 42): " + [ acc.constructor.name, n.constructor.name ]);
            };
        };
        return go([  ]);
    };
};
var evalUncurriedEffectApp = function (env) {
    return function (hd) {
        return function (spine) {
            if (hd instanceof SemTyped) {
                return evalUncurriedEffectApp(env)(hd.value1)(spine);
            };
            if (hd instanceof SemMkEffectFn) {
                return evalUncurriedBeta(NeutUncurriedEffectApp.create)(hd.value0)(spine);
            };
            if (hd instanceof SemLet) {
                return new SemLet(hd.value0, hd.value1, function (nextVal) {
                    return makeLet(Data_Maybe.Nothing.value)(hd.value2(nextVal))(function (nextFn) {
                        return evalUncurriedEffectApp(bindLocal(bindLocal(env)(new One(nextVal)))(new One(nextFn)))(nextFn)(spine);
                    });
                });
            };
            if (hd instanceof NeutFail) {
                return new NeutFail(hd.value0);
            };
            return guardFailOver(Data_Foldable.foldableArray)(identity)(spine)(NeutUncurriedEffectApp.create(hd));
        };
    };
};
var mkFnFromArgs = function (dictEval) {
    var eval1 = $$eval(dictEval);
    return function (env) {
        return function (args) {
            return function (body) {
                return new SemMkFn(Data_Foldable.foldr(Data_Foldable.foldableArray)(function (v) {
                    return function (next) {
                        return function (env$prime) {
                            return new MkFnNext(v.value0, (function () {
                                var $2103 = bindLocal(env$prime);
                                return function ($2104) {
                                    return next($2103(One.create($2104)));
                                };
                            })());
                        };
                    };
                })((function () {
                    var $2105 = Data_Function.flip(eval1)(body);
                    return function ($2106) {
                        return MkFnApplied.create($2105($2106));
                    };
                })())(args)(env));
            };
        };
    };
};
var evalBackendSyntax = function (dictEval) {
    var eval1 = $$eval(dictEval);
    return {
        "eval": function (v) {
            return function (v1) {
                if (v1 instanceof PureScript_Backend_Optimizer_Syntax.Var) {
                    var v2 = v.evalExternSpine(v)(v1.value0)([  ]);
                    if (v2 instanceof Data_Maybe.Just) {
                        return v2.value0;
                    };
                    if (v2 instanceof Data_Maybe.Nothing) {
                        return new SemRef(new EvalExtern(v1.value0), [  ], Data_Lazy.defer(function (v3) {
                            var v4 = v.evalExternRef(v)(v1.value0);
                            if (v4 instanceof Data_Maybe.Just) {
                                return deref(v4.value0);
                            };
                            if (v4 instanceof Data_Maybe.Nothing) {
                                return new NeutVar(v1.value0);
                            };
                            throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Semantics (line 230, column 13 - line 234, column 29): " + [ v4.constructor.name ]);
                        }));
                    };
                    throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Semantics (line 225, column 7 - line 234, column 29): " + [ v2.constructor.name ]);
                };
                if (v1 instanceof PureScript_Backend_Optimizer_Syntax.Local) {
                    var v2 = lookupLocal(v)(v1.value1);
                    if (v2 instanceof Data_Maybe.Just && v2.value0 instanceof One) {
                        return v2.value0.value0;
                    };
                    var v3 = function (v4) {
                        return Partial_Unsafe.unsafeCrashWith("Unbound local at level " + Data_Show.show(Data_Show.showInt)(Data_Newtype.unwrap()(v1.value1)));
                    };
                    if (v2 instanceof Data_Maybe.Just && v2.value0 instanceof Group) {
                        var $1810 = Control_Bind.bindFlipped(Data_Maybe.bindMaybe)(Data_Function.flip(lookup)(v2.value0.value0))(v1.value0);
                        if ($1810 instanceof Data_Maybe.Just) {
                            return Data_Lazy.force($1810.value0);
                        };
                        return v3(true);
                    };
                    return v3(true);
                };
                if (v1 instanceof PureScript_Backend_Optimizer_Syntax.App) {
                    return evalApp(v)($$eval(dictEval)(v)(v1.value0))(Data_Array_NonEmpty.toArray(Data_Functor.map(Data_Array_NonEmpty_Internal.functorNonEmptyArray)($$eval(dictEval)(v))(v1.value1)));
                };
                if (v1 instanceof PureScript_Backend_Optimizer_Syntax.UncurriedApp) {
                    return evalUncurriedApp(v)($$eval(dictEval)(v)(v1.value0))(Data_Functor.map(Data_Functor.functorArray)($$eval(dictEval)(v))(v1.value1));
                };
                if (v1 instanceof PureScript_Backend_Optimizer_Syntax.UncurriedAbs) {
                    var loop = function (env$prime) {
                        return function (v2) {
                            if (v2 instanceof Data_List_Types.Nil) {
                                return new MkFnApplied($$eval(dictEval)(env$prime)(v1.value1));
                            };
                            if (v2 instanceof Data_List_Types.Cons) {
                                return new MkFnNext(v2.value0, function (nextArg) {
                                    return loop(bindLocal(env$prime)(new One(nextArg)))(v2.value1);
                                });
                            };
                            throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Semantics (line 248, column 21 - line 253, column 53): " + [ v2.constructor.name ]);
                        };
                    };
                    return new SemMkFn(loop(v)(toUnfoldable(Data_Functor.map(Data_Functor.functorArray)(Data_Tuple.fst)(v1.value0))));
                };
                if (v1 instanceof PureScript_Backend_Optimizer_Syntax.UncurriedEffectApp) {
                    return evalUncurriedEffectApp(v)($$eval(dictEval)(v)(v1.value0))(Data_Functor.map(Data_Functor.functorArray)($$eval(dictEval)(v))(v1.value1));
                };
                if (v1 instanceof PureScript_Backend_Optimizer_Syntax.UncurriedEffectAbs) {
                    var loop = function (env$prime) {
                        return function (v2) {
                            if (v2 instanceof Data_List_Types.Nil) {
                                return new MkFnApplied($$eval(dictEval)(env$prime)(v1.value1));
                            };
                            if (v2 instanceof Data_List_Types.Cons) {
                                return new MkFnNext(v2.value0, function (nextArg) {
                                    return loop(bindLocal(env$prime)(new One(nextArg)))(v2.value1);
                                });
                            };
                            throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Semantics (line 259, column 21 - line 264, column 53): " + [ v2.constructor.name ]);
                        };
                    };
                    return new SemMkEffectFn(loop(v)(toUnfoldable(Data_Functor.map(Data_Functor.functorArray)(Data_Tuple.fst)(v1.value0))));
                };
                if (v1 instanceof PureScript_Backend_Optimizer_Syntax.Abs) {
                    return PureScript_Backend_Optimizer_Utils.foldr1Array(function (v2) {
                        return function (next) {
                            return function (env$prime) {
                                return new SemLam(v2.value0, (function () {
                                    var $2107 = bindLocal(env$prime);
                                    return function ($2108) {
                                        return next($2107(One.create($2108)));
                                    };
                                })());
                            };
                        };
                    })(function (v2) {
                        return function (env$prime) {
                            return new SemLam(v2.value0, (function () {
                                var $2109 = Data_Function.flip(eval1)(v1.value1);
                                var $2110 = bindLocal(env$prime);
                                return function ($2111) {
                                    return $2109($2110(One.create($2111)));
                                };
                            })());
                        };
                    })(v1.value0)(v);
                };
                if (v1 instanceof PureScript_Backend_Optimizer_Syntax.Let) {
                    return makeLet(v1.value0)($$eval(dictEval)(v)(v1.value2))((function () {
                        var $2112 = Data_Function.flip(eval1)(v1.value3);
                        var $2113 = bindLocal(v);
                        return function ($2114) {
                            return $2112($2113(One.create($2114)));
                        };
                    })());
                };
                if (v1 instanceof PureScript_Backend_Optimizer_Syntax.LetRec) {
                    var bindGroup = function (sem) {
                        var $2115 = Data_Function.flip(eval1)(sem);
                        var $2116 = bindLocal(v);
                        return function ($2117) {
                            return $2115($2116(Group.create($2117)));
                        };
                    };
                    return new SemLetRec(Data_Functor.map(Data_Array_NonEmpty_Internal.functorNonEmptyArray)(Data_Functor.map(Data_Tuple.functorTuple)(bindGroup))(v1.value1), bindGroup(v1.value2));
                };
                if (v1 instanceof PureScript_Backend_Optimizer_Syntax.EffectBind) {
                    return makeEffectBind(v1.value0)($$eval(dictEval)(v)(v1.value2))((function () {
                        var $2118 = Data_Function.flip(eval1)(v1.value3);
                        var $2119 = bindLocal(v);
                        return function ($2120) {
                            return $2118($2119(One.create($2120)));
                        };
                    })());
                };
                if (v1 instanceof PureScript_Backend_Optimizer_Syntax.EffectPure) {
                    return guardFail($$eval(dictEval)(v)(v1.value0))(SemEffectPure.create);
                };
                if (v1 instanceof PureScript_Backend_Optimizer_Syntax.EffectDefer) {
                    return guardFail($$eval(dictEval)(v)(v1.value0))(SemEffectDefer.create);
                };
                if (v1 instanceof PureScript_Backend_Optimizer_Syntax.Accessor) {
                    return evalAccessor(v)($$eval(dictEval)(v)(v1.value0))(v1.value1);
                };
                if (v1 instanceof PureScript_Backend_Optimizer_Syntax.Update) {
                    return evalUpdate($$eval(dictEval)(v)(v1.value0))(Data_Functor.map(Data_Functor.functorArray)(Data_Functor.map(PureScript_Backend_Optimizer_CoreFn.functorProp)($$eval(dictEval)(v)))(v1.value1));
                };
                if (v1 instanceof PureScript_Backend_Optimizer_Syntax.Branch) {
                    return evalBranches(v)(Data_Functor.map(Data_Array_NonEmpty_Internal.functorNonEmptyArray)(evalPair(dictEval)(v))(v1.value0))(Data_Lazy.defer(function (v2) {
                        return $$eval(dictEval)(v)(v1.value1);
                    }));
                };
                if (v1 instanceof PureScript_Backend_Optimizer_Syntax.PrimOp) {
                    return evalPrimOp(v)(Data_Functor.map(PureScript_Backend_Optimizer_Syntax.functorBackendOperator)($$eval(dictEval)(v))(v1.value0));
                };
                if (v1 instanceof PureScript_Backend_Optimizer_Syntax.PrimEffect) {
                    return guardFailOver(PureScript_Backend_Optimizer_Syntax.foldableBackendEffect)(identity)(Data_Functor.map(PureScript_Backend_Optimizer_Syntax.functorBackendEffect)($$eval(dictEval)(v))(v1.value0))(NeutPrimEffect.create);
                };
                if (v1 instanceof PureScript_Backend_Optimizer_Syntax.PrimUndefined) {
                    return NeutPrimUndefined.value;
                };
                if (v1 instanceof PureScript_Backend_Optimizer_Syntax.Lit) {
                    return guardFailOver(PureScript_Backend_Optimizer_CoreFn.foldableLiteral)(identity)(Data_Functor.map(PureScript_Backend_Optimizer_CoreFn.functorLiteral)($$eval(dictEval)(v))(v1.value0))(NeutLit.create);
                };
                if (v1 instanceof PureScript_Backend_Optimizer_Syntax.Fail) {
                    return new NeutFail(v1.value0);
                };
                if (v1 instanceof PureScript_Backend_Optimizer_Syntax.CtorDef) {
                    return new NeutCtorDef(new PureScript_Backend_Optimizer_CoreFn.Qualified(new Data_Maybe.Just((Data_Newtype.unwrap()(v)).currentModule), v1.value2), v1.value0, v1.value1, v1.value2, v1.value3);
                };
                if (v1 instanceof PureScript_Backend_Optimizer_Syntax.CtorSaturated) {
                    return guardFailOver(Data_Foldable.foldableArray)(Data_Tuple.snd)(Data_Functor.map(Data_Functor.functorArray)(Data_Functor.map(Data_Tuple.functorTuple)($$eval(dictEval)(v)))(v1.value4))(NeutData.create(v1.value0)(v1.value1)(v1.value2)(v1.value3));
                };
                if (v1 instanceof PureScript_Backend_Optimizer_Syntax.Typed) {
                    return new SemTyped(v1.value0, $$eval(dictEval)(v)(v1.value1));
                };
                throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Semantics (line 223, column 22 - line 304, column 30): " + [ v1.constructor.name ]);
            };
        }
    };
};
var $lazy_evalBackendExpr = /* #__PURE__ */ $runtime_lazy("evalBackendExpr", "PureScript.Backend.Optimizer.Semantics", function () {
    return {
        "eval": (function () {
            var go = function ($copy_env) {
                return function ($copy_v) {
                    var $tco_var_env = $copy_env;
                    var $tco_done = false;
                    var $tco_result;
                    function $tco_loop(env, v) {
                        if (v instanceof ExprRewrite) {
                            if (v.value1 instanceof RewriteInline) {
                                $tco_var_env = bindLocal(env)(new One($$eval($lazy_evalBackendExpr(0))(env)(v.value1.value2)));
                                $copy_v = v.value1.value3;
                                return;
                            };
                            if (v.value1 instanceof RewriteUncurry) {
                                $tco_done = true;
                                return new SemLet(v.value1.value0, mkFnFromArgs($lazy_evalBackendExpr(0))(env)(Data_Array_NonEmpty.toArray(v.value1.value2))(v.value1.value3), function (newFn) {
                                    return $$eval($lazy_evalBackendExpr(0))(bindLocal(env)(new One(mkUncurriedAppRewrite(env)(newFn)(Data_Array_NonEmpty.length(v.value1.value2)))))(v.value1.value4);
                                });
                            };
                            if (v.value1 instanceof RewriteStop) {
                                $tco_done = true;
                                return new NeutStop(v.value1.value0);
                            };
                            if (v.value1 instanceof RewriteUnpackOp) {
                                if (v.value1.value2 instanceof UnpackRecord) {
                                    $tco_done = true;
                                    return Data_Foldable.foldr(Data_Foldable.foldableArray)(function (v1) {
                                        return function (next) {
                                            return function (props$prime) {
                                                return makeLet(Data_Maybe.Nothing.value)($$eval($lazy_evalBackendExpr(0))(env)(v1.value1))(function (val) {
                                                    return next(Data_Array.snoc(props$prime)(new PureScript_Backend_Optimizer_CoreFn.Prop(v1.value0, val)));
                                                });
                                            };
                                        };
                                    })((function () {
                                        var $2121 = Data_Function.flip($$eval($lazy_evalBackendExpr(0)))(v.value1.value3);
                                        var $2122 = bindLocal(env);
                                        return function ($2123) {
                                            return $2121($2122(One.create(NeutLit.create(PureScript_Backend_Optimizer_CoreFn.LitRecord.create($2123)))));
                                        };
                                    })())(v.value1.value2.value0)([  ]);
                                };
                                if (v.value1.value2 instanceof UnpackUpdate) {
                                    $tco_done = true;
                                    return makeLet(Data_Maybe.Nothing.value)($$eval($lazy_evalBackendExpr(0))(env)(v.value1.value2.value0))(function (hd$prime) {
                                        return Data_Foldable.foldr(Data_Foldable.foldableArray)(function (v1) {
                                            return function (next) {
                                                return function (props$prime) {
                                                    return makeLet(Data_Maybe.Nothing.value)($$eval($lazy_evalBackendExpr(0))(env)(v1.value1))(function (val) {
                                                        return next(Data_Array.snoc(props$prime)(new PureScript_Backend_Optimizer_CoreFn.Prop(v1.value0, val)));
                                                    });
                                                };
                                            };
                                        })((function () {
                                            var $2124 = Data_Function.flip($$eval($lazy_evalBackendExpr(0)))(v.value1.value3);
                                            var $2125 = bindLocal(env);
                                            var $2126 = NeutUpdate.create(hd$prime);
                                            return function ($2127) {
                                                return $2124($2125(One.create($2126($2127))));
                                            };
                                        })())(v.value1.value2.value1)([  ]);
                                    });
                                };
                                if (v.value1.value2 instanceof UnpackArray) {
                                    $tco_done = true;
                                    return Data_Foldable.foldr(Data_Foldable.foldableArray)(function (expr) {
                                        return function (next) {
                                            return function (exprs$prime) {
                                                return makeLet(Data_Maybe.Nothing.value)($$eval($lazy_evalBackendExpr(0))(env)(expr))(function (val) {
                                                    return next(Data_Array.snoc(exprs$prime)(val));
                                                });
                                            };
                                        };
                                    })((function () {
                                        var $2128 = Data_Function.flip($$eval($lazy_evalBackendExpr(0)))(v.value1.value3);
                                        var $2129 = bindLocal(env);
                                        return function ($2130) {
                                            return $2128($2129(One.create(NeutLit.create(PureScript_Backend_Optimizer_CoreFn.LitArray.create($2130)))));
                                        };
                                    })())(v.value1.value2.value0)([  ]);
                                };
                                if (v.value1.value2 instanceof UnpackData) {
                                    $tco_done = true;
                                    return Data_Foldable.foldr(Data_Foldable.foldableArray)(function (v1) {
                                        return function (next) {
                                            return function (props$prime) {
                                                return makeLet(Data_Maybe.Nothing.value)($$eval($lazy_evalBackendExpr(0))(env)(v1.value1))(function (val) {
                                                    return next(Data_Array.snoc(props$prime)(new Data_Tuple.Tuple(v1.value0, val)));
                                                });
                                            };
                                        };
                                    })((function () {
                                        var $2131 = Data_Function.flip($$eval($lazy_evalBackendExpr(0)))(v.value1.value3);
                                        var $2132 = bindLocal(env);
                                        var $2133 = NeutData.create(v.value1.value2.value0)(v.value1.value2.value1)(v.value1.value2.value2)(v.value1.value2.value3);
                                        return function ($2134) {
                                            return $2131($2132(One.create($2133($2134))));
                                        };
                                    })())(v.value1.value2.value4)([  ]);
                                };
                                throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Semantics (line 320, column 13 - line 357, column 21): " + [ v.value1.value2.constructor.name ]);
                            };
                            if (v.value1 instanceof RewriteDistBranchesLet) {
                                $tco_done = true;
                                return rewriteBranches((function () {
                                    var $2135 = Data_Function.flip($$eval($lazy_evalBackendExpr(0)))(v.value1.value4);
                                    var $2136 = bindLocal(env);
                                    return function ($2137) {
                                        return $2135($2136(One.create($2137)));
                                    };
                                })())(evalBranches(env)(Data_Functor.map(Data_Array_NonEmpty_Internal.functorNonEmptyArray)(evalPair($lazy_evalBackendExpr(0))(env))(v.value1.value2))(Data_Lazy.defer(function (v1) {
                                    return $$eval($lazy_evalBackendExpr(0))(env)(v.value1.value3);
                                })));
                            };
                            if (v.value1 instanceof RewriteDistBranchesOp) {
                                var dist = (function () {
                                    if (v.value1.value2 instanceof DistApp) {
                                        return Data_Function.flip(evalApp(env))(Data_Array_NonEmpty.toArray(Data_Functor.map(Data_Array_NonEmpty_Internal.functorNonEmptyArray)($$eval($lazy_evalBackendExpr(0))(env))(v.value1.value2.value0)));
                                    };
                                    if (v.value1.value2 instanceof DistUncurriedApp) {
                                        return Data_Function.flip(evalUncurriedApp(env))(Data_Functor.map(Data_Functor.functorArray)($$eval($lazy_evalBackendExpr(0))(env))(v.value1.value2.value0));
                                    };
                                    if (v.value1.value2 instanceof DistAccessor) {
                                        return Data_Function.flip(evalAccessor(env))(v.value1.value2.value0);
                                    };
                                    if (v.value1.value2 instanceof DistPrimOp1) {
                                        var $2138 = evalPrimOp(env);
                                        var $2139 = PureScript_Backend_Optimizer_Syntax.Op1.create(v.value1.value2.value0);
                                        return function ($2140) {
                                            return $2138($2139($2140));
                                        };
                                    };
                                    if (v.value1.value2 instanceof DistPrimOp2L) {
                                        var $2141 = evalPrimOp(env);
                                        var $2142 = Data_Function.flip(PureScript_Backend_Optimizer_Syntax.Op2.create(v.value1.value2.value0))($$eval($lazy_evalBackendExpr(0))(env)(v.value1.value2.value1));
                                        return function ($2143) {
                                            return $2141($2142($2143));
                                        };
                                    };
                                    if (v.value1.value2 instanceof DistPrimOp2R) {
                                        var $2144 = evalPrimOp(env);
                                        var $2145 = PureScript_Backend_Optimizer_Syntax.Op2.create(v.value1.value2.value1)($$eval($lazy_evalBackendExpr(0))(env)(v.value1.value2.value0));
                                        return function ($2146) {
                                            return $2144($2145($2146));
                                        };
                                    };
                                    throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Semantics (line 364, column 20 - line 376, column 58): " + [ v.value1.value2.constructor.name ]);
                                })();
                                $tco_done = true;
                                return rewriteBranches(dist)(evalBranches(env)(Data_Functor.map(Data_Array_NonEmpty_Internal.functorNonEmptyArray)(evalPair($lazy_evalBackendExpr(0))(env))(v.value1.value0))(Data_Lazy.defer(function (v1) {
                                    return $$eval($lazy_evalBackendExpr(0))(env)(v.value1.value1);
                                })));
                            };
                            throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Semantics (line 311, column 9 - line 376, column 58): " + [ v.value1.constructor.name ]);
                        };
                        if (v instanceof ExprSyntax) {
                            $tco_done = true;
                            return $$eval(evalBackendSyntax($lazy_evalBackendExpr(0)))(env)(v.value1);
                        };
                        throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Semantics (line 309, column 14 - line 378, column 22): " + [ v.constructor.name ]);
                    };
                    while (!$tco_done) {
                        $tco_result = $tco_loop($tco_var_env, $copy_v);
                    };
                    return $tco_result;
                };
            };
            return go;
        })()
    };
});
var evalBackendExpr = /* #__PURE__ */ $lazy_evalBackendExpr(306);
var optimize = function (traceSteps) {
    return function (ctx) {
        return function (env) {
            return function (v) {
                return function (initN) {
                    return function (originalExpr) {
                        var goStep = function (n) {
                            return function (expr1) {
                                if (n === 0) {
                                    var name = Data_Foldable.foldMap(Data_Foldable.foldableMaybe)(Data_Monoid.monoidString)(function ($2147) {
                                        return (function (v1) {
                                            return v1 + ".";
                                        })(unwrap1($2147));
                                    })(v.value0) + v.value1;
                                    return Partial_Unsafe.unsafeCrashWith(name + ": Possible infinite optimization loop.");
                                };
                                if (Data_Boolean.otherwise) {
                                    var expr2 = quote(ctx)($$eval(evalBackendExpr)(env)(expr1));
                                    var v1 = PureScript_Backend_Optimizer_Analysis.analysisOf(hasAnalysisBackendExpr)(expr2);
                                    return new Data_Tuple.Tuple(v1.rewrite, expr2);
                                };
                                throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Semantics (line 1714, column 3 - line 1714, column 60): " + [ n.constructor.name, expr1.constructor.name ]);
                            };
                        };
                        var go = function ($copy_steps) {
                            return function ($copy_n) {
                                return function ($copy_expr1) {
                                    var $tco_var_steps = $copy_steps;
                                    var $tco_var_n = $copy_n;
                                    var $tco_done = false;
                                    var $tco_result;
                                    function $tco_loop(steps, n, expr1) {
                                        var v1 = goStep(n)(expr1);
                                        var newSteps = (function () {
                                            if (traceSteps) {
                                                return new Data_List_Types.Cons(v1.value1, steps);
                                            };
                                            return steps;
                                        })();
                                        if (v1.value0) {
                                            $tco_var_steps = newSteps;
                                            $tco_var_n = n - 1 | 0;
                                            $copy_expr1 = v1.value1;
                                            return;
                                        };
                                        $tco_done = true;
                                        return new Data_Tuple.Tuple(Data_Array.reverse(Data_List.toUnfoldable(Data_Unfoldable.unfoldableArray)(newSteps)), v1.value1);
                                    };
                                    while (!$tco_done) {
                                        $tco_result = $tco_loop($tco_var_steps, $tco_var_n, $copy_expr1);
                                    };
                                    return $tco_result;
                                };
                            };
                        };
                        return go((function () {
                            if (traceSteps) {
                                return Control_Applicative.pure(Data_List_Types.applicativeList)(originalExpr);
                            };
                            return Data_List_Types.Nil.value;
                        })())(initN)(originalExpr);
                    };
                };
            };
        };
    };
};
var evalNeutralExpr = {
    "eval": function (env) {
        return function (v) {
            return $$eval(evalBackendSyntax(evalNeutralExpr))(env)(v);
        };
    }
};
var evalBackendSyntax1 = /* #__PURE__ */ evalBackendSyntax(evalNeutralExpr);
var analysisFromDirective = function (v) {
    return function (v1) {
        if (v1 instanceof InlineAlways) {
            return mempty;
        };
        if (v1 instanceof InlineNever) {
            return {
                usages: v.usages,
                args: v.args,
                rewrite: v.rewrite,
                deps: v.deps,
                result: v.result,
                externs: v.externs,
                complexity: PureScript_Backend_Optimizer_Analysis.NonTrivial.value,
                size: top
            };
        };
        if (v1 instanceof InlineArity) {
            return {
                usages: v.usages,
                size: v.size,
                complexity: v.complexity,
                rewrite: v.rewrite,
                deps: v.deps,
                result: v.result,
                externs: v.externs,
                args: Data_Array.take(v1.value0)(v.args)
            };
        };
        if (v1 instanceof InlineDefault) {
            return v;
        };
        throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Semantics (line 1137, column 52 - line 1145, column 29): " + [ v1.constructor.name ]);
    };
};
var addStop = function (v) {
    return function (ref) {
        return function (acc) {
            return {
                currentModule: v.currentModule,
                evalExternRef: v.evalExternRef,
                evalExternSpine: v.evalExternSpine,
                locals: v.locals,
                directives: Data_Map_Internal.alter(ordEvalRef)(function (v2) {
                    if (v2 instanceof Data_Maybe.Just) {
                        return new Data_Maybe.Just(Data_Map_Internal.insert(ordInlineAccessor)(acc)(InlineNever.value)(v2.value0));
                    };
                    return new Data_Maybe.Just(Data_Map_Internal.singleton(acc)(InlineNever.value));
                })(ref)(v.directives)
            };
        };
    };
};
var envForGroup = function (env) {
    return function (ref) {
        return function (acc) {
            return function (group) {
                if (Data_Array["null"](group)) {
                    return env;
                };
                if (Data_Boolean.otherwise) {
                    return addStop(env)(ref)(acc);
                };
                throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Semantics (line 969, column 1 - line 969, column 82): " + [ env.constructor.name, ref.constructor.name, acc.constructor.name, group.constructor.name ]);
            };
        };
    };
};
var evalExternFromImpl = function (v) {
    return function (qual) {
        return function (v1) {
            return function (spine) {
                if (spine.length === 0) {
                    if (v1.value1 instanceof ExternExpr) {
                        var ref = new EvalExtern(qual);
                        var v2 = Control_Bind.bind(Data_Maybe.bindMaybe)(Data_Map_Internal.lookup(ordEvalRef)(ref)(v.directives))(Data_Map_Internal.lookup(ordInlineAccessor)(InlineRef.value));
                        if (v2 instanceof Data_Maybe.Just && v2.value0 instanceof InlineNever) {
                            return new Data_Maybe.Just(new NeutStop(qual));
                        };
                        if (v2 instanceof Data_Maybe.Just && v2.value0 instanceof InlineAlways) {
                            return new Data_Maybe.Just($$eval(evalNeutralExpr)(envForGroup(v)(ref)(InlineRef.value)(v1.value1.value0))(v1.value1.value1));
                        };
                        if (v2 instanceof Data_Maybe.Just && v2.value0 instanceof InlineArity) {
                            return Data_Maybe.Nothing.value;
                        };
                        if (v1.value1.value1 instanceof PureScript_Backend_Optimizer_Syntax.Lit && shouldInlineExternLiteral(v1.value1.value1.value0)) {
                            return new Data_Maybe.Just($$eval(evalNeutralExpr)(envForGroup(v)(ref)(InlineRef.value)(v1.value1.value0))(v1.value1.value1));
                        };
                        if (shouldInlineExternReference(qual)(v1.value0)(v1.value1.value1)) {
                            return new Data_Maybe.Just($$eval(evalNeutralExpr)(envForGroup(v)(ref)(InlineRef.value)(v1.value1.value0))(v1.value1.value1));
                        };
                        return Data_Maybe.Nothing.value;
                    };
                    if (v1.value1 instanceof ExternCtor && v1.value1.value4.length === 0) {
                        return new Data_Maybe.Just(new NeutData(qual, v1.value1.value1, v1.value1.value2, v1.value1.value3, [  ]));
                    };
                    return Data_Maybe.Nothing.value;
                };
                if (spine.length === 1 && (spine[0] instanceof ExternAccessor && spine[0].value0 instanceof PureScript_Backend_Optimizer_Syntax.GetProp)) {
                    if (v1.value1 instanceof ExternExpr) {
                        var ref = new EvalExtern(qual);
                        var v2 = Control_Bind.bind(Data_Maybe.bindMaybe)(Data_Map_Internal.lookup(ordEvalRef)(ref)(v.directives))(Data_Map_Internal.lookup(ordInlineAccessor)(new InlineProp(spine[0].value0.value0)));
                        if (v2 instanceof Data_Maybe.Just && v2.value0 instanceof InlineNever) {
                            return new Data_Maybe.Just(neutralSpine(new NeutStop(qual))(spine));
                        };
                        if (v2 instanceof Data_Maybe.Just && v2.value0 instanceof InlineAlways) {
                            return new Data_Maybe.Just(evalSpine(v)($$eval(evalNeutralExpr)(envForGroup(v)(ref)(new InlineProp(spine[0].value0.value0))(v1.value1.value0))(v1.value1.value1))(spine));
                        };
                        return Data_Maybe.Nothing.value;
                    };
                    var v2 = function (v3) {
                        return Data_Maybe.Nothing.value;
                    };
                    if (v1.value1 instanceof ExternDict) {
                        var $1989 = PureScript_Backend_Optimizer_CoreFn.findProp(spine[0].value0.value0)(v1.value1.value1);
                        if ($1989 instanceof Data_Maybe.Just) {
                            var ref = new EvalExtern(qual);
                            var v3 = Control_Bind.bind(Data_Maybe.bindMaybe)(Data_Map_Internal.lookup(ordEvalRef)(ref)(v.directives))(Data_Map_Internal.lookup(ordInlineAccessor)(new InlineProp(spine[0].value0.value0)));
                            if (v3 instanceof Data_Maybe.Just && v3.value0 instanceof InlineNever) {
                                return new Data_Maybe.Just(neutralSpine(new NeutStop(qual))(spine));
                            };
                            if (v3 instanceof Data_Maybe.Just && v3.value0 instanceof InlineAlways) {
                                return new Data_Maybe.Just($$eval(evalNeutralExpr)(envForGroup(v)(ref)(new InlineProp(spine[0].value0.value0))(v1.value1.value0))($1989.value0.value1));
                            };
                            if (v3 instanceof Data_Maybe.Just && v3.value0 instanceof InlineArity) {
                                return Data_Maybe.Nothing.value;
                            };
                            if (shouldInlineExternAccessor(qual)($1989.value0.value0)($1989.value0.value1)(spine[0].value0)) {
                                return new Data_Maybe.Just($$eval(evalNeutralExpr)(envForGroup(v)(ref)(new InlineProp(spine[0].value0.value0))(v1.value1.value0))($1989.value0.value1));
                            };
                            return Data_Maybe.Nothing.value;
                        };
                        return v2(true);
                    };
                    return v2(true);
                };
                if (spine.length === 2 && (spine[0] instanceof ExternAccessor && (spine[0].value0 instanceof PureScript_Backend_Optimizer_Syntax.GetProp && spine[1] instanceof ExternApp))) {
                    if (v1.value1 instanceof ExternExpr) {
                        var ref = new EvalExtern(qual);
                        var v2 = Control_Bind.bind(Data_Maybe.bindMaybe)(Data_Map_Internal.lookup(ordEvalRef)(ref)(v.directives))(Data_Map_Internal.lookup(ordInlineAccessor)(new InlineProp(spine[0].value0.value0)));
                        if (v2 instanceof Data_Maybe.Just && v2.value0 instanceof InlineNever) {
                            return new Data_Maybe.Just(neutralSpine(new NeutStop(qual))(spine));
                        };
                        if (v2 instanceof Data_Maybe.Just && v2.value0 instanceof InlineAlways) {
                            return new Data_Maybe.Just(evalSpine(v)($$eval(evalNeutralExpr)(envForGroup(v)(ref)(new InlineProp(spine[0].value0.value0))(v1.value1.value0))(v1.value1.value1))(spine));
                        };
                        if (v2 instanceof Data_Maybe.Just && v2.value0 instanceof InlineArity) {
                            if (Data_Array.length(spine[1].value0) >= v2.value0.value0) {
                                return new Data_Maybe.Just(evalSpine(v)($$eval(evalNeutralExpr)(envForGroup(v)(ref)(new InlineProp(spine[0].value0.value0))(v1.value1.value0))(v1.value1.value1))(spine));
                            };
                            if (Data_Boolean.otherwise) {
                                return Data_Maybe.Nothing.value;
                            };
                        };
                        return Data_Maybe.Nothing.value;
                    };
                    var v2 = function (v3) {
                        return Data_Maybe.Nothing.value;
                    };
                    if (v1.value1 instanceof ExternDict) {
                        var $2012 = PureScript_Backend_Optimizer_CoreFn.findProp(spine[0].value0.value0)(v1.value1.value1);
                        if ($2012 instanceof Data_Maybe.Just) {
                            var ref = new EvalExtern(qual);
                            var v3 = Control_Bind.bind(Data_Maybe.bindMaybe)(Data_Map_Internal.lookup(ordEvalRef)(ref)(v.directives))(Data_Map_Internal.lookup(ordInlineAccessor)(new InlineProp(spine[0].value0.value0)));
                            if (v3 instanceof Data_Maybe.Just && v3.value0 instanceof InlineNever) {
                                return new Data_Maybe.Just(neutralSpine(new NeutStop(qual))(spine));
                            };
                            if (v3 instanceof Data_Maybe.Just && v3.value0 instanceof InlineAlways) {
                                return new Data_Maybe.Just(evalApp(v)($$eval(evalNeutralExpr)(envForGroup(v)(ref)(new InlineProp(spine[0].value0.value0))(v1.value1.value0))($2012.value0.value1))(spine[1].value0));
                            };
                            if (v3 instanceof Data_Maybe.Just && v3.value0 instanceof InlineArity) {
                                if (Data_Array.length(spine[1].value0) >= v3.value0.value0) {
                                    return new Data_Maybe.Just(evalApp(v)($$eval(evalNeutralExpr)(envForGroup(v)(ref)(new InlineProp(spine[0].value0.value0))(v1.value1.value0))($2012.value0.value1))(spine[1].value0));
                                };
                                if (Data_Boolean.otherwise) {
                                    return Data_Maybe.Nothing.value;
                                };
                            };
                            if (shouldInlineExternApp(qual)($2012.value0.value0)($2012.value0.value1)(spine[1].value0)) {
                                return new Data_Maybe.Just(evalApp(v)($$eval(evalNeutralExpr)(envForGroup(v)(ref)(new InlineProp(spine[0].value0.value0))(v1.value1.value0))($2012.value0.value1))(spine[1].value0));
                            };
                            return Data_Maybe.Nothing.value;
                        };
                        return v2(true);
                    };
                    return v2(true);
                };
                if (spine.length === 1 && spine[0] instanceof ExternApp) {
                    if (v1.value1 instanceof ExternExpr) {
                        var ref = new EvalExtern(qual);
                        var v2 = Control_Bind.bind(Data_Maybe.bindMaybe)(Data_Map_Internal.lookup(ordEvalRef)(ref)(v.directives))(Data_Map_Internal.lookup(ordInlineAccessor)(InlineRef.value));
                        if (v2 instanceof Data_Maybe.Just && v2.value0 instanceof InlineNever) {
                            return new Data_Maybe.Just(neutralSpine(new NeutStop(qual))(spine));
                        };
                        if (v2 instanceof Data_Maybe.Just && v2.value0 instanceof InlineAlways) {
                            return new Data_Maybe.Just(evalApp(v)($$eval(evalNeutralExpr)(envForGroup(v)(ref)(InlineRef.value)(v1.value1.value0))(v1.value1.value1))(spine[0].value0));
                        };
                        if (v2 instanceof Data_Maybe.Just && v2.value0 instanceof InlineArity) {
                            if (Data_Array.length(spine[0].value0) >= v2.value0.value0) {
                                return new Data_Maybe.Just(evalApp(v)($$eval(evalNeutralExpr)(envForGroup(v)(ref)(InlineRef.value)(v1.value1.value0))(v1.value1.value1))(spine[0].value0));
                            };
                            if (Data_Boolean.otherwise) {
                                return Data_Maybe.Nothing.value;
                            };
                        };
                        if (shouldInlineExternApp(qual)(v1.value0)(v1.value1.value1)(spine[0].value0)) {
                            return new Data_Maybe.Just(evalApp(v)($$eval(evalNeutralExpr)(envForGroup(v)(ref)(InlineRef.value)(v1.value1.value0))(v1.value1.value1))(spine[0].value0));
                        };
                        return Data_Maybe.Nothing.value;
                    };
                    if (v1.value1 instanceof ExternCtor && Data_Array.length(v1.value1.value4) === Data_Array.length(spine[0].value0)) {
                        return new Data_Maybe.Just(new NeutData(qual, v1.value1.value1, v1.value1.value2, v1.value1.value3, Data_Array.zip(v1.value1.value4)(spine[0].value0)));
                    };
                    return Data_Maybe.Nothing.value;
                };
                if (spine.length === 2 && (spine[0] instanceof ExternApp && (spine[1] instanceof ExternAccessor && spine[1].value0 instanceof PureScript_Backend_Optimizer_Syntax.GetProp))) {
                    if (v1.value1 instanceof ExternExpr) {
                        var ref = new EvalExtern(qual);
                        var v2 = Control_Bind.bind(Data_Maybe.bindMaybe)(Data_Map_Internal.lookup(ordEvalRef)(ref)(v.directives))(Data_Map_Internal.lookup(ordInlineAccessor)(new InlineSpineProp(spine[1].value0.value0)));
                        if (v2 instanceof Data_Maybe.Just && v2.value0 instanceof InlineNever) {
                            return new Data_Maybe.Just(neutralSpine(new NeutStop(qual))(spine));
                        };
                        if (v2 instanceof Data_Maybe.Just && v2.value0 instanceof InlineAlways) {
                            return new Data_Maybe.Just(evalSpine(v)($$eval(evalNeutralExpr)(envForGroup(v)(ref)(new InlineSpineProp(spine[1].value0.value0))(v1.value1.value0))(v1.value1.value1))(spine));
                        };
                        return Data_Maybe.Nothing.value;
                    };
                    return Data_Maybe.Nothing.value;
                };
                if (spine.length === 3 && (spine[0] instanceof ExternApp && (spine[1] instanceof ExternAccessor && (spine[1].value0 instanceof PureScript_Backend_Optimizer_Syntax.GetProp && spine[2] instanceof ExternApp)))) {
                    if (v1.value1 instanceof ExternExpr) {
                        var ref = new EvalExtern(qual);
                        var v2 = Control_Bind.bind(Data_Maybe.bindMaybe)(Data_Map_Internal.lookup(ordEvalRef)(ref)(v.directives))(Data_Map_Internal.lookup(ordInlineAccessor)(new InlineSpineProp(spine[1].value0.value0)));
                        if (v2 instanceof Data_Maybe.Just && v2.value0 instanceof InlineNever) {
                            return new Data_Maybe.Just(neutralSpine(new NeutStop(qual))(spine));
                        };
                        if (v2 instanceof Data_Maybe.Just && v2.value0 instanceof InlineAlways) {
                            return new Data_Maybe.Just(evalSpine(v)($$eval(evalNeutralExpr)(envForGroup(v)(ref)(new InlineSpineProp(spine[1].value0.value0))(v1.value1.value0))(v1.value1.value1))(spine));
                        };
                        if (v2 instanceof Data_Maybe.Just && (v2.value0 instanceof InlineArity && Data_Array.length(spine[2].value0) >= v2.value0.value0)) {
                            return new Data_Maybe.Just(evalSpine(v)($$eval(evalNeutralExpr)(envForGroup(v)(ref)(new InlineSpineProp(spine[1].value0.value0))(v1.value1.value0))(v1.value1.value1))(spine));
                        };
                        return Data_Maybe.Nothing.value;
                    };
                    return Data_Maybe.Nothing.value;
                };
                return Data_Maybe.Nothing.value;
            };
        };
    };
};
var evalExternRefFromImpl = function (env) {
    return function (qual) {
        return function (v) {
            if (v.value1 instanceof ExternExpr && isRefExpr(v.value1.value1)) {
                return $$eval(evalBackendSyntax1)(envForGroup(env)(new EvalExtern(qual))(InlineRef.value)(v.value1.value0))(v.value1.value1);
            };
            if (v.value1 instanceof ExternDict) {
                return new NeutLit(new PureScript_Backend_Optimizer_CoreFn.LitRecord(Data_Functor.map(Data_Functor.functorArray)(function (v1) {
                    return new PureScript_Backend_Optimizer_CoreFn.Prop(v1.value0, $$eval(evalBackendSyntax1)(envForGroup(env)(new EvalExtern(qual))(new InlineProp(v1.value0))(v.value1.value0))(v1.value1.value1));
                })(v.value1.value1)));
            };
            return new NeutVar(qual);
        };
    };
};
export {
    $$eval as eval,
    MkFnApplied,
    MkFnNext,
    SemTyped,
    SemRef,
    SemLam,
    SemMkFn,
    SemMkEffectFn,
    SemLet,
    SemLetRec,
    SemEffectBind,
    SemEffectPure,
    SemEffectDefer,
    SemBranch,
    SemAssocOp,
    NeutLocal,
    NeutVar,
    NeutStop,
    NeutData,
    NeutCtorDef,
    NeutApp,
    NeutAccessor,
    NeutUpdate,
    NeutLit,
    NeutFail,
    NeutUncurriedApp,
    NeutUncurriedEffectApp,
    NeutPrimOp,
    NeutPrimEffect,
    NeutPrimUndefined,
    SemConditional,
    ExprSyntax,
    ExprRewrite,
    RewriteInline,
    RewriteUncurry,
    RewriteStop,
    RewriteUnpackOp,
    RewriteDistBranchesLet,
    RewriteDistBranchesOp,
    UnpackRecord,
    UnpackUpdate,
    UnpackArray,
    UnpackData,
    DistApp,
    DistUncurriedApp,
    DistAccessor,
    DistPrimOp1,
    DistPrimOp2L,
    DistPrimOp2R,
    ExternExpr,
    ExternDict,
    ExternCtor,
    One,
    Group,
    ExternApp,
    ExternUncurriedApp,
    ExternAccessor,
    ExternPrimOp,
    EvalExtern,
    EvalLocal,
    InlineProp,
    InlineSpineProp,
    InlineRef,
    InlineDefault,
    InlineNever,
    InlineAlways,
    InlineArity,
    Env,
    lookupLocal,
    bindLocal,
    insertDirective,
    addStop,
    snocApp,
    evalApp,
    evalUncurriedApp,
    evalUncurriedEffectApp,
    evalUncurriedBeta,
    evalSpine,
    neutralSpine,
    neutralApp,
    evalAccessor,
    evalUpdate,
    evalBranches,
    rewriteBranches,
    evalPair,
    makeEffectBind,
    makeLet,
    floatLet,
    floatLetWith,
    deref,
    evalPrimOp,
    evalPrimOpOrd,
    evalPrimOpOrdNumber,
    evalPrimOpNumNumber,
    evalPrimOpNumInt,
    evalPrimOpNot,
    primOpOrdNot,
    isAssocPrimOp,
    evalAssocOp,
    evalAssocOp$prime,
    evalRef,
    evalRefSpine,
    evalEvalRef,
    snocSpine,
    envForGroup,
    evalExternFromImpl,
    evalExternRefFromImpl,
    isRefExpr,
    analysisFromDirective,
    liftBoolean,
    liftInt,
    liftNumber,
    liftString,
    liftOp1,
    liftOp2,
    caseString,
    caseInt,
    caseNumber,
    Ctx,
    nextLevel,
    effectfully,
    purely,
    quote,
    build,
    buildBranchCond,
    simplifyCondIsTag,
    simplifyCondBoolean,
    isSimplePredicate,
    simplifyCondRedundantElse,
    simplifyCondLiftAnd,
    buildStop,
    buildDefault,
    rewriteInline,
    isReference,
    shouldEtaReduce,
    shouldUnpackCtor,
    shouldUnpackRecord,
    shouldUnpackUpdate,
    shouldUnpackArray,
    shouldDistributeBranches,
    shouldDistributeBranchApps,
    shouldDistributeBranchUncurriedApps,
    shouldDistributeBranchAccessor,
    shouldDistributeBranchPrimOp1,
    shouldDistributeBranchPrimOp2L,
    shouldDistributeBranchPrimOp2R,
    shouldUncurryAbs,
    shouldInlineLet,
    shouldInlineExternReference,
    shouldInlineExternApp,
    shouldInlineExternAppArg,
    isEffectSemantics,
    isPartialAssocOp,
    shouldInlineExternAccessor,
    shouldInlineExternLiteral,
    isAbs,
    isKnownEffect,
    NeutralExpr,
    optimize,
    freeze,
    foldBackendExpr,
    evalMkFn,
    mkUncurriedAppRewrite,
    mkFnFromArgs,
    guardFail,
    guardFailOver,
    eqBackendExpr,
    eqBackendRewrite,
    eqUnpackOp,
    eqDistOp,
    hasAnalysisBackendExpr,
    hasSyntaxBackendExpr,
    eqEvalRef,
    ordEvalRef,
    eqInlineAccessor,
    ordInlineAccessor,
    newtypeEnv_,
    evalBackendSyntax,
    evalBackendExpr,
    evalNeutralExpr,
    newtypeNeutralExpr_
};
