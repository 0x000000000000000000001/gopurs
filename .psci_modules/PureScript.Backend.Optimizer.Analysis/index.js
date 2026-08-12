import * as Data_Array from "../Data.Array/index.js";
import * as Data_Array_NonEmpty from "../Data.Array.NonEmpty/index.js";
import * as Data_Array_NonEmpty_Internal from "../Data.Array.NonEmpty.Internal/index.js";
import * as Data_Boolean from "../Data.Boolean/index.js";
import * as Data_Foldable from "../Data.Foldable/index.js";
import * as Data_Functor from "../Data.Functor/index.js";
import * as Data_Map_Internal from "../Data.Map.Internal/index.js";
import * as Data_Maybe from "../Data.Maybe/index.js";
import * as Data_Monoid from "../Data.Monoid/index.js";
import * as Data_Newtype from "../Data.Newtype/index.js";
import * as Data_Ord from "../Data.Ord/index.js";
import * as Data_Ordering from "../Data.Ordering/index.js";
import * as Data_Semigroup from "../Data.Semigroup/index.js";
import * as Data_Set from "../Data.Set/index.js";
import * as Data_String_CodeUnits from "../Data.String.CodeUnits/index.js";
import * as Data_Tuple from "../Data.Tuple/index.js";
import * as PureScript_Backend_Optimizer_CoreFn from "../PureScript.Backend.Optimizer.CoreFn/index.js";
import * as PureScript_Backend_Optimizer_Syntax from "../PureScript.Backend.Optimizer.Syntax/index.js";
var ordQualified = /* #__PURE__ */ PureScript_Backend_Optimizer_CoreFn.ordQualified(PureScript_Backend_Optimizer_CoreFn.ordIdent);
var unwrap = /* #__PURE__ */ Data_Newtype.unwrap();
var KnownNeutral = /* #__PURE__ */ (function () {
    function KnownNeutral() {

    };
    KnownNeutral.value = new KnownNeutral();
    return KnownNeutral;
})();
var Unknown = /* #__PURE__ */ (function () {
    function Unknown() {

    };
    Unknown.value = new Unknown();
    return Unknown;
})();
var Trivial = /* #__PURE__ */ (function () {
    function Trivial() {

    };
    Trivial.value = new Trivial();
    return Trivial;
})();
var Deref = /* #__PURE__ */ (function () {
    function Deref() {

    };
    Deref.value = new Deref();
    return Deref;
})();
var KnownSize = /* #__PURE__ */ (function () {
    function KnownSize() {

    };
    KnownSize.value = new KnownSize();
    return KnownSize;
})();
var NonTrivial = /* #__PURE__ */ (function () {
    function NonTrivial() {

    };
    NonTrivial.value = new NonTrivial();
    return NonTrivial;
})();
var CaptureNone = /* #__PURE__ */ (function () {
    function CaptureNone() {

    };
    CaptureNone.value = new CaptureNone();
    return CaptureNone;
})();
var CaptureBranch = /* #__PURE__ */ (function () {
    function CaptureBranch() {

    };
    CaptureBranch.value = new CaptureBranch();
    return CaptureBranch;
})();
var CaptureClosure = /* #__PURE__ */ (function () {
    function CaptureClosure() {

    };
    CaptureClosure.value = new CaptureClosure();
    return CaptureClosure;
})();
var Usage = function (x) {
    return x;
};
var BackendAnalysis = function (x) {
    return x;
};
var semigroupResultTerm = {
    append: function (v) {
        return function (v1) {
            if (v instanceof Unknown) {
                return Unknown.value;
            };
            if (v1 instanceof Unknown) {
                return Unknown.value;
            };
            return KnownNeutral.value;
        };
    }
};
var newtypeUsage_ = {
    Coercible0: function () {
        return undefined;
    }
};
var newtypeBackendAnalysis_ = {
    Coercible0: function () {
        return undefined;
    }
};
var monoidResultTerm = /* #__PURE__ */ (function () {
    return {
        mempty: KnownNeutral.value,
        Semigroup0: function () {
            return semigroupResultTerm;
        }
    };
})();
var eqResultTerm = {
    eq: function (x) {
        return function (y) {
            if (x instanceof KnownNeutral && y instanceof KnownNeutral) {
                return true;
            };
            if (x instanceof Unknown && y instanceof Unknown) {
                return true;
            };
            return false;
        };
    }
};
var eqComplexity = {
    eq: function (x) {
        return function (y) {
            if (x instanceof Trivial && y instanceof Trivial) {
                return true;
            };
            if (x instanceof Deref && y instanceof Deref) {
                return true;
            };
            if (x instanceof KnownSize && y instanceof KnownSize) {
                return true;
            };
            if (x instanceof NonTrivial && y instanceof NonTrivial) {
                return true;
            };
            return false;
        };
    }
};
var ordComplexity = {
    compare: function (x) {
        return function (y) {
            if (x instanceof Trivial && y instanceof Trivial) {
                return Data_Ordering.EQ.value;
            };
            if (x instanceof Trivial) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof Trivial) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof Deref && y instanceof Deref) {
                return Data_Ordering.EQ.value;
            };
            if (x instanceof Deref) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof Deref) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof KnownSize && y instanceof KnownSize) {
                return Data_Ordering.EQ.value;
            };
            if (x instanceof KnownSize) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof KnownSize) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof NonTrivial && y instanceof NonTrivial) {
                return Data_Ordering.EQ.value;
            };
            throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Analysis (line 0, column 0 - line 0, column 0): " + [ x.constructor.name, y.constructor.name ]);
        };
    },
    Eq0: function () {
        return eqComplexity;
    }
};
var semigroupComplexity = {
    append: /* #__PURE__ */ Data_Ord.max(ordComplexity)
};
var monoidComplexity = /* #__PURE__ */ (function () {
    return {
        mempty: Trivial.value,
        Semigroup0: function () {
            return semigroupComplexity;
        }
    };
})();
var eqCapture = {
    eq: function (x) {
        return function (y) {
            if (x instanceof CaptureNone && y instanceof CaptureNone) {
                return true;
            };
            if (x instanceof CaptureBranch && y instanceof CaptureBranch) {
                return true;
            };
            if (x instanceof CaptureClosure && y instanceof CaptureClosure) {
                return true;
            };
            return false;
        };
    }
};
var ordCapture = {
    compare: function (x) {
        return function (y) {
            if (x instanceof CaptureNone && y instanceof CaptureNone) {
                return Data_Ordering.EQ.value;
            };
            if (x instanceof CaptureNone) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof CaptureNone) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof CaptureBranch && y instanceof CaptureBranch) {
                return Data_Ordering.EQ.value;
            };
            if (x instanceof CaptureBranch) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof CaptureBranch) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof CaptureClosure && y instanceof CaptureClosure) {
                return Data_Ordering.EQ.value;
            };
            throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Analysis (line 0, column 0 - line 0, column 0): " + [ x.constructor.name, y.constructor.name ]);
        };
    },
    Eq0: function () {
        return eqCapture;
    }
};
var semigroupCapture = {
    append: /* #__PURE__ */ Data_Ord.max(ordCapture)
};
var monoidCapture = /* #__PURE__ */ (function () {
    return {
        mempty: CaptureNone.value,
        Semigroup0: function () {
            return semigroupCapture;
        }
    };
})();
var mempty = /* #__PURE__ */ Data_Monoid.mempty(monoidCapture);
var semigroupUsage = {
    append: function (v) {
        return function (v1) {
            return {
                total: v.total + v1.total | 0,
                captured: Data_Semigroup.append(semigroupCapture)(v.captured)(v1.captured),
                arities: Data_Set.union(Data_Ord.ordInt)(v.arities)(v1.arities),
                call: v.call + v1.call | 0,
                access: v.access + v1.access | 0,
                "case": v["case"] + v1["case"] | 0,
                update: v.update + v1.update | 0
            };
        };
    }
};
var append = /* #__PURE__ */ Data_Semigroup.append(semigroupUsage);
var monoidUsage = {
    mempty: {
        total: 0,
        captured: mempty,
        arities: Data_Set.empty,
        call: 0,
        access: 0,
        "case": 0,
        update: 0
    },
    Semigroup0: function () {
        return semigroupUsage;
    }
};
var semigroupBackendAnalysis = {
    append: function (v) {
        return function (v1) {
            return {
                usages: Data_Map_Internal.unionWith(PureScript_Backend_Optimizer_Syntax.ordLevel)(append)(v.usages)(v1.usages),
                size: v.size + v1.size | 0,
                complexity: Data_Semigroup.append(semigroupComplexity)(v.complexity)(v1.complexity),
                args: [  ],
                rewrite: v.rewrite || v1.rewrite,
                deps: Data_Set.union(ordQualified)(v.deps)(v1.deps),
                result: Data_Semigroup.append(semigroupResultTerm)(v.result)(v1.result),
                externs: v.externs || v1.externs
            };
        };
    }
};
var monoidBackendAnalysis = /* #__PURE__ */ (function () {
    return {
        mempty: {
            usages: Data_Map_Internal.empty,
            size: 0,
            complexity: Trivial.value,
            args: [  ],
            rewrite: false,
            deps: Data_Set.empty,
            result: KnownNeutral.value,
            externs: false
        },
        Semigroup0: function () {
            return semigroupBackendAnalysis;
        }
    };
})();
var mempty1 = /* #__PURE__ */ Data_Monoid.mempty(monoidBackendAnalysis);
var withRewrite = function (v) {
    return {
        usages: v.usages,
        size: v.size,
        complexity: v.complexity,
        args: v.args,
        deps: v.deps,
        result: v.result,
        externs: v.externs,
        rewrite: true
    };
};
var withResult = function (r) {
    return function (v) {
        return {
            usages: v.usages,
            size: v.size,
            complexity: v.complexity,
            args: v.args,
            rewrite: v.rewrite,
            deps: v.deps,
            externs: v.externs,
            result: r
        };
    };
};
var withArgs = function (args) {
    return function (v) {
        return {
            usages: v.usages,
            size: v.size,
            complexity: v.complexity,
            rewrite: v.rewrite,
            deps: v.deps,
            result: v.result,
            externs: v.externs,
            args: args
        };
    };
};
var usedDep = function (dep) {
    return function (v) {
        return {
            usages: v.usages,
            size: v.size,
            complexity: v.complexity,
            args: v.args,
            rewrite: v.rewrite,
            result: v.result,
            externs: v.externs,
            deps: Data_Set.insert(ordQualified)(dep)(v.deps)
        };
    };
};
var used = function (level) {
    var v = Data_Monoid.mempty(monoidBackendAnalysis);
    return {
        size: v.size,
        complexity: v.complexity,
        args: v.args,
        rewrite: v.rewrite,
        deps: v.deps,
        result: v.result,
        externs: v.externs,
        usages: Data_Map_Internal.singleton(level)({
            total: 1,
            captured: mempty,
            arities: Data_Set.empty,
            call: 0,
            access: 0,
            "case": 0,
            update: 0
        })
    };
};
var updated = function (level) {
    return function (v) {
        return {
            size: v.size,
            complexity: v.complexity,
            args: v.args,
            rewrite: v.rewrite,
            deps: v.deps,
            result: v.result,
            externs: v.externs,
            usages: Data_Map_Internal.update(PureScript_Backend_Optimizer_Syntax.ordLevel)((function () {
                var $287 = Data_Newtype.over()()(Usage)(function (us) {
                    return {
                        access: us.access,
                        arities: us.arities,
                        call: us.call,
                        captured: us.captured,
                        "case": us["case"],
                        total: us.total,
                        update: us.update + 1 | 0
                    };
                });
                return function ($288) {
                    return Data_Maybe.Just.create($287($288));
                };
            })())(level)(v.usages)
        };
    };
};
var externs = function (v) {
    return {
        usages: v.usages,
        size: v.size,
        complexity: v.complexity,
        args: v.args,
        rewrite: v.rewrite,
        deps: v.deps,
        result: v.result,
        externs: true
    };
};
var complex = function (complexity) {
    return function (v) {
        return {
            usages: v.usages,
            size: v.size,
            args: v.args,
            rewrite: v.rewrite,
            deps: v.deps,
            result: v.result,
            externs: v.externs,
            complexity: Data_Semigroup.append(semigroupComplexity)(v.complexity)(complexity)
        };
    };
};
var cased = function (level) {
    return function (v) {
        return {
            size: v.size,
            complexity: v.complexity,
            args: v.args,
            rewrite: v.rewrite,
            deps: v.deps,
            result: v.result,
            externs: v.externs,
            usages: Data_Map_Internal.update(PureScript_Backend_Optimizer_Syntax.ordLevel)((function () {
                var $289 = Data_Newtype.over()()(Usage)(function (us) {
                    return {
                        access: us.access,
                        arities: us.arities,
                        call: us.call,
                        captured: us.captured,
                        total: us.total,
                        update: us.update,
                        "case": us["case"] + 1 | 0
                    };
                });
                return function ($290) {
                    return Data_Maybe.Just.create($289($290));
                };
            })())(level)(v.usages)
        };
    };
};
var capture = function (cap) {
    return function (v) {
        return {
            size: v.size,
            complexity: v.complexity,
            args: v.args,
            rewrite: v.rewrite,
            deps: v.deps,
            result: v.result,
            externs: v.externs,
            usages: Data_Functor.map(Data_Map_Internal.functorMap)(Data_Newtype.over()()(Usage)(function (v2) {
                return {
                    access: v2.access,
                    arities: v2.arities,
                    call: v2.call,
                    "case": v2["case"],
                    total: v2.total,
                    update: v2.update,
                    captured: cap
                };
            }))(v.usages)
        };
    };
};
var callArity = function (lvl) {
    return function (arity) {
        return function (v) {
            return {
                size: v.size,
                complexity: v.complexity,
                args: v.args,
                rewrite: v.rewrite,
                deps: v.deps,
                result: v.result,
                externs: v.externs,
                usages: Data_Map_Internal.update(PureScript_Backend_Optimizer_Syntax.ordLevel)((function () {
                    var $291 = Data_Newtype.over()()(Usage)(function (us) {
                        return {
                            access: us.access,
                            captured: us.captured,
                            "case": us["case"],
                            total: us.total,
                            update: us.update,
                            arities: Data_Set.insert(Data_Ord.ordInt)(arity)(us.arities),
                            call: us.call + 1 | 0
                        };
                    });
                    return function ($292) {
                        return Data_Maybe.Just.create($291($292));
                    };
                })())(lvl)(v.usages)
            };
        };
    };
};
var bump = function (v) {
    return {
        usages: v.usages,
        complexity: v.complexity,
        args: v.args,
        rewrite: v.rewrite,
        deps: v.deps,
        result: v.result,
        externs: v.externs,
        size: v.size + 1 | 0
    };
};
var boundArg = function (level) {
    return function (v) {
        var v1 = Data_Map_Internal.pop(PureScript_Backend_Optimizer_Syntax.ordLevel)(level)(v.usages);
        if (v1 instanceof Data_Maybe.Nothing) {
            return {
                usages: v.usages,
                size: v.size,
                complexity: v.complexity,
                rewrite: v.rewrite,
                deps: v.deps,
                result: v.result,
                externs: v.externs,
                args: Data_Array.cons(Data_Monoid.mempty(monoidUsage))(v.args)
            };
        };
        if (v1 instanceof Data_Maybe.Just) {
            return {
                size: v.size,
                complexity: v.complexity,
                rewrite: v.rewrite,
                deps: v.deps,
                result: v.result,
                externs: v.externs,
                usages: v1.value0.value1,
                args: Data_Array.cons(v1.value0.value0)(v.args)
            };
        };
        throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Analysis (line 129, column 38 - line 133, column 66): " + [ v1.constructor.name ]);
    };
};
var bound = function (level) {
    return function (v) {
        return {
            size: v.size,
            complexity: v.complexity,
            args: v.args,
            rewrite: v.rewrite,
            deps: v.deps,
            result: v.result,
            externs: v.externs,
            usages: Data_Map_Internal["delete"](PureScript_Backend_Optimizer_Syntax.ordLevel)(level)(v.usages)
        };
    };
};
var analysisOf = function (dict) {
    return dict.analysisOf;
};
var analyzeDefault = function (dictHasAnalysis) {
    var $293 = Data_Foldable.foldMap(PureScript_Backend_Optimizer_Syntax.foldableBackendSyntax)(monoidBackendAnalysis)(analysisOf(dictHasAnalysis));
    return function ($294) {
        return bump($293($294));
    };
};
var resultOf = function (dictHasAnalysis) {
    var $295 = analysisOf(dictHasAnalysis);
    return function ($296) {
        return (function (v) {
            return v.result;
        })(unwrap($295($296)));
    };
};
var accessed = function (level) {
    return function (v) {
        return {
            size: v.size,
            complexity: v.complexity,
            args: v.args,
            rewrite: v.rewrite,
            deps: v.deps,
            result: v.result,
            externs: v.externs,
            usages: Data_Map_Internal.update(PureScript_Backend_Optimizer_Syntax.ordLevel)((function () {
                var $297 = Data_Newtype.over()()(Usage)(function (us) {
                    return {
                        arities: us.arities,
                        call: us.call,
                        captured: us.captured,
                        "case": us["case"],
                        total: us.total,
                        update: us.update,
                        access: us.access + 1 | 0
                    };
                });
                return function ($298) {
                    return Data_Maybe.Just.create($297($298));
                };
            })())(level)(v.usages)
        };
    };
};
var analyze = function (dictHasAnalysis) {
    var analysisOf1 = analysisOf(dictHasAnalysis);
    var resultOf1 = resultOf(dictHasAnalysis);
    return function (dictHasSyntax) {
        return function (externAnalysis) {
            return function (expr) {
                if (expr instanceof PureScript_Backend_Optimizer_Syntax.Var) {
                    var analysis = bump(externs(usedDep(expr.value0)(mempty1)));
                    var v = externAnalysis(expr.value0)(Data_Maybe.Nothing.value);
                    if (v instanceof Data_Maybe.Just) {
                        return withArgs(v.value0.args)(analysis);
                    };
                    if (v instanceof Data_Maybe.Nothing) {
                        return analysis;
                    };
                    throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Analysis (line 206, column 5 - line 210, column 17): " + [ v.constructor.name ]);
                };
                if (expr instanceof PureScript_Backend_Optimizer_Syntax.Local) {
                    return bump(used(expr.value1));
                };
                if (expr instanceof PureScript_Backend_Optimizer_Syntax.Let) {
                    return withResult(resultOf(dictHasAnalysis)(expr.value3))(bump(complex(NonTrivial.value)(Data_Semigroup.append(semigroupBackendAnalysis)(analysisOf(dictHasAnalysis)(expr.value2))(bound(expr.value1)(analysisOf(dictHasAnalysis)(expr.value3))))));
                };
                if (expr instanceof PureScript_Backend_Optimizer_Syntax.LetRec) {
                    return withResult(resultOf(dictHasAnalysis)(expr.value2))(complex(NonTrivial.value)(bound(expr.value0)(bump(Data_Semigroup.append(semigroupBackendAnalysis)(Data_Foldable.foldMap(Data_Array_NonEmpty_Internal.foldableNonEmptyArray)(monoidBackendAnalysis)(function ($299) {
                        return analysisOf1(Data_Tuple.snd($299));
                    })(expr.value1))(analysisOf(dictHasAnalysis)(expr.value2))))));
                };
                if (expr instanceof PureScript_Backend_Optimizer_Syntax.EffectBind) {
                    return withResult(Unknown.value)(complex(NonTrivial.value)(capture(CaptureClosure.value)(bump(Data_Semigroup.append(semigroupBackendAnalysis)(analysisOf(dictHasAnalysis)(expr.value2))(bound(expr.value1)(analysisOf(dictHasAnalysis)(expr.value3)))))));
                };
                if (expr instanceof PureScript_Backend_Optimizer_Syntax.EffectPure) {
                    return withResult(Unknown.value)(capture(CaptureClosure.value)(bump(analysisOf(dictHasAnalysis)(expr.value0))));
                };
                if (expr instanceof PureScript_Backend_Optimizer_Syntax.EffectDefer) {
                    return withResult(Unknown.value)(capture(CaptureClosure.value)(bump(analysisOf(dictHasAnalysis)(expr.value0))));
                };
                if (expr instanceof PureScript_Backend_Optimizer_Syntax.Abs) {
                    return withResult(KnownNeutral.value)(complex(KnownSize.value)(capture(CaptureClosure.value)(Data_Foldable.foldr(Data_Array_NonEmpty_Internal.foldableNonEmptyArray)(function ($300) {
                        return boundArg(Data_Tuple.snd($300));
                    })(analyzeDefault(dictHasAnalysis)(expr))(expr.value0))));
                };
                if (expr instanceof PureScript_Backend_Optimizer_Syntax.UncurriedAbs) {
                    return withResult(KnownNeutral.value)(complex(KnownSize.value)(capture(CaptureClosure.value)(Data_Foldable.foldr(Data_Foldable.foldableArray)(function ($301) {
                        return boundArg(Data_Tuple.snd($301));
                    })(analyzeDefault(dictHasAnalysis)(expr))(expr.value0))));
                };
                if (expr instanceof PureScript_Backend_Optimizer_Syntax.UncurriedApp) {
                    var analysis = withResult(Unknown.value)(complex(NonTrivial.value)(analyzeDefault(dictHasAnalysis)(expr)));
                    var v = PureScript_Backend_Optimizer_Syntax.syntaxOf(dictHasSyntax)(expr.value0);
                    if (v instanceof Data_Maybe.Just && v.value0 instanceof PureScript_Backend_Optimizer_Syntax.Local) {
                        return callArity(v.value0.value1)(Data_Array.length(expr.value1))(analysis);
                    };
                    return analysis;
                };
                if (expr instanceof PureScript_Backend_Optimizer_Syntax.UncurriedEffectAbs) {
                    return withResult(KnownNeutral.value)(complex(KnownSize.value)(capture(CaptureClosure.value)(Data_Foldable.foldr(Data_Foldable.foldableArray)(function ($302) {
                        return boundArg(Data_Tuple.snd($302));
                    })(analyzeDefault(dictHasAnalysis)(expr))(expr.value0))));
                };
                if (expr instanceof PureScript_Backend_Optimizer_Syntax.UncurriedEffectApp) {
                    var analysis = withResult(Unknown.value)(complex(NonTrivial.value)(capture(CaptureClosure.value)(analyzeDefault(dictHasAnalysis)(expr))));
                    var v = PureScript_Backend_Optimizer_Syntax.syntaxOf(dictHasSyntax)(expr.value0);
                    if (v instanceof Data_Maybe.Just && v.value0 instanceof PureScript_Backend_Optimizer_Syntax.Local) {
                        return callArity(v.value0.value1)(Data_Array.length(expr.value1))(analysis);
                    };
                    return analysis;
                };
                var v = function (v1) {
                    if (expr instanceof PureScript_Backend_Optimizer_Syntax.Update) {
                        var analysis = withResult(Unknown.value)(complex(NonTrivial.value)(analyzeDefault(dictHasAnalysis)(expr)));
                        var v2 = PureScript_Backend_Optimizer_Syntax.syntaxOf(dictHasSyntax)(expr.value0);
                        if (v2 instanceof Data_Maybe.Just && v2.value0 instanceof PureScript_Backend_Optimizer_Syntax.Local) {
                            return updated(v2.value0.value1)(analysis);
                        };
                        return analysis;
                    };
                    if (expr instanceof PureScript_Backend_Optimizer_Syntax.CtorSaturated) {
                        return withResult(KnownNeutral.value)(bump(usedDep(expr.value0)(Data_Foldable.foldMap(Data_Foldable.foldableArray)(monoidBackendAnalysis)(Data_Foldable.foldMap(Data_Foldable.foldableTuple)(monoidBackendAnalysis)(analysisOf1))(expr.value4))));
                    };
                    if (expr instanceof PureScript_Backend_Optimizer_Syntax.CtorDef) {
                        return complex(NonTrivial.value)(analyzeDefault(dictHasAnalysis)(expr));
                    };
                    if (expr instanceof PureScript_Backend_Optimizer_Syntax.Branch) {
                        var v2 = Data_Array_NonEmpty.head(expr.value0);
                        var result = Data_Foldable.foldMap(Data_Array_NonEmpty_Internal.foldableNonEmptyArray)(monoidResultTerm)(function ($303) {
                            return resultOf1(PureScript_Backend_Optimizer_Syntax.sndPair($303));
                        })(expr.value0);
                        return withResult(result)(complex(NonTrivial.value)(Data_Semigroup.append(semigroupBackendAnalysis)(analysisOf(dictHasAnalysis)(v2.value0))(Data_Semigroup.append(semigroupBackendAnalysis)(capture(CaptureBranch.value)(analysisOf(dictHasAnalysis)(v2.value1)))(Data_Semigroup.append(semigroupBackendAnalysis)(capture(CaptureBranch.value)(Data_Foldable.foldMap(Data_Foldable.foldableArray)(monoidBackendAnalysis)(Data_Foldable.foldMap(PureScript_Backend_Optimizer_Syntax.foldablePair)(monoidBackendAnalysis)(analysisOf1))(Data_Array_NonEmpty.tail(expr.value0))))(capture(CaptureBranch.value)(analysisOf(dictHasAnalysis)(expr.value1)))))));
                    };
                    if (expr instanceof PureScript_Backend_Optimizer_Syntax.Fail) {
                        return complex(NonTrivial.value)(analyzeDefault(dictHasAnalysis)(expr));
                    };
                    if (expr instanceof PureScript_Backend_Optimizer_Syntax.PrimOp) {
                        var analysis = withResult(Unknown.value)(complex(NonTrivial.value)(analyzeDefault(dictHasAnalysis)(expr)));
                        var v2 = function (v3) {
                            if (expr.value0 instanceof PureScript_Backend_Optimizer_Syntax.Op1 && (expr.value0.value0 instanceof PureScript_Backend_Optimizer_Syntax.OpIsTag && Data_Boolean.otherwise)) {
                                return usedDep(expr.value0.value0.value0)(analysis);
                            };
                            return analysis;
                        };
                        if (expr.value0 instanceof PureScript_Backend_Optimizer_Syntax.Op1 && expr.value0.value0 instanceof PureScript_Backend_Optimizer_Syntax.OpIsTag) {
                            var $216 = PureScript_Backend_Optimizer_Syntax.syntaxOf(dictHasSyntax)(expr.value0.value1);
                            if ($216 instanceof Data_Maybe.Just && $216.value0 instanceof PureScript_Backend_Optimizer_Syntax.Local) {
                                return cased($216.value0.value1)(usedDep(expr.value0.value0.value0)(analysis));
                            };
                            return v2(true);
                        };
                        return v2(true);
                    };
                    if (expr instanceof PureScript_Backend_Optimizer_Syntax.PrimEffect) {
                        return withResult(Unknown.value)(complex(NonTrivial.value)(capture(CaptureClosure.value)(analyzeDefault(dictHasAnalysis)(expr))));
                    };
                    if (expr instanceof PureScript_Backend_Optimizer_Syntax.PrimUndefined) {
                        return analyzeDefault(dictHasAnalysis)(expr);
                    };
                    if (expr instanceof PureScript_Backend_Optimizer_Syntax.Accessor) {
                        var analysis = (function () {
                            if (expr.value1 instanceof PureScript_Backend_Optimizer_Syntax.GetCtorField) {
                                return withResult(Unknown.value)(usedDep(expr.value1.value0)(analyzeDefault(dictHasAnalysis)(expr)));
                            };
                            return withResult(Unknown.value)(analyzeDefault(dictHasAnalysis)(expr));
                        })();
                        var v2 = PureScript_Backend_Optimizer_Syntax.syntaxOf(dictHasSyntax)(expr.value0);
                        if (v2 instanceof Data_Maybe.Just && v2.value0 instanceof PureScript_Backend_Optimizer_Syntax.Accessor) {
                            return analysis;
                        };
                        if (v2 instanceof Data_Maybe.Just && v2.value0 instanceof PureScript_Backend_Optimizer_Syntax.Local) {
                            return accessed(v2.value0.value1)(complex(Deref.value)(analysis));
                        };
                        if (v2 instanceof Data_Maybe.Just && v2.value0 instanceof PureScript_Backend_Optimizer_Syntax.Var) {
                            var v3 = function (v4) {
                                return complex(Trivial.value)(analysis);
                            };
                            if (expr.value1 instanceof PureScript_Backend_Optimizer_Syntax.GetProp) {
                                var $240 = externAnalysis(v2.value0.value0)(new Data_Maybe.Just(expr.value1.value0));
                                if ($240 instanceof Data_Maybe.Just) {
                                    return withArgs($240.value0.args)(complex(Trivial.value)(analysis));
                                };
                                return v3(true);
                            };
                            return v3(true);
                        };
                        return complex(Deref.value)(analysis);
                    };
                    if (expr instanceof PureScript_Backend_Optimizer_Syntax.Lit) {
                        var analysis = withResult(KnownNeutral.value)(analyzeDefault(dictHasAnalysis)(expr));
                        if (expr.value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitArray && Data_Array.length(expr.value0.value0) > 0) {
                            return complex(KnownSize.value)(analysis);
                        };
                        if (expr.value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitRecord) {
                            if (Data_Array.length(expr.value0.value0) > 0) {
                                return complex(KnownSize.value)(analysis);
                            };
                            if (Data_Boolean.otherwise) {
                                return analysis;
                            };
                        };
                        if (expr.value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitString && Data_String_CodeUnits.length(expr.value0.value0) > 128) {
                            return complex(KnownSize.value)(analysis);
                        };
                        return analysis;
                    };
                    if (expr instanceof PureScript_Backend_Optimizer_Syntax.Typed) {
                        return analysisOf(dictHasAnalysis)(expr.value1);
                    };
                    throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Analysis (line 203, column 1 - line 203, column 150): " + [ expr.constructor.name ]);
                };
                if (expr instanceof PureScript_Backend_Optimizer_Syntax.App) {
                    var $256 = analysisOf(dictHasAnalysis)(expr.value0);
                    var remainingArgs = Data_Array.drop(Data_Array_NonEmpty.length(expr.value1))($256.args);
                    var analysis = (function () {
                        if (Data_Array["null"](remainingArgs)) {
                            return complex(NonTrivial.value)(analyzeDefault(dictHasAnalysis)(expr));
                        };
                        if (Data_Boolean.otherwise) {
                            return analyzeDefault(dictHasAnalysis)(expr);
                        };
                        throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Analysis (line 298, column 5 - line 302, column 30): " + [  ]);
                    })();
                    return withArgs(remainingArgs)((function () {
                        var v1 = PureScript_Backend_Optimizer_Syntax.syntaxOf(dictHasSyntax)(expr.value0);
                        if (v1 instanceof Data_Maybe.Just && v1.value0 instanceof PureScript_Backend_Optimizer_Syntax.Local) {
                            return withResult(Unknown.value)(callArity(v1.value0.value1)(Data_Array_NonEmpty.length(expr.value1))(bump(analysis)));
                        };
                        return withResult(Unknown.value)(bump(analysis));
                    })());
                };
                return v(true);
            };
        };
    };
};
var analyzeEffectBlock = function (dictHasAnalysis) {
    var analysisOf1 = analysisOf(dictHasAnalysis);
    return function (dictHasSyntax) {
        return function (externAnalysis) {
            return function (expr) {
                if (expr instanceof PureScript_Backend_Optimizer_Syntax.Let) {
                    return withResult(resultOf(dictHasAnalysis)(expr.value3))(complex(NonTrivial.value)(bump(Data_Semigroup.append(semigroupBackendAnalysis)(analysisOf(dictHasAnalysis)(expr.value2))(bound(expr.value1)(analysisOf(dictHasAnalysis)(expr.value3))))));
                };
                if (expr instanceof PureScript_Backend_Optimizer_Syntax.LetRec) {
                    return withResult(resultOf(dictHasAnalysis)(expr.value2))(complex(NonTrivial.value)(bound(expr.value0)(bump(Data_Semigroup.append(semigroupBackendAnalysis)(Data_Foldable.foldMap(Data_Array_NonEmpty_Internal.foldableNonEmptyArray)(monoidBackendAnalysis)(function ($304) {
                        return analysisOf1(Data_Tuple.snd($304));
                    })(expr.value1))(analysisOf(dictHasAnalysis)(expr.value2))))));
                };
                if (expr instanceof PureScript_Backend_Optimizer_Syntax.EffectBind) {
                    return withResult(Unknown.value)(complex(NonTrivial.value)(bump(Data_Semigroup.append(semigroupBackendAnalysis)(analysisOf(dictHasAnalysis)(expr.value2))(bound(expr.value1)(analysisOf(dictHasAnalysis)(expr.value3))))));
                };
                if (expr instanceof PureScript_Backend_Optimizer_Syntax.EffectPure) {
                    return withResult(Unknown.value)(bump(analysisOf(dictHasAnalysis)(expr.value0)));
                };
                if (expr instanceof PureScript_Backend_Optimizer_Syntax.EffectDefer) {
                    return withResult(Unknown.value)(bump(analysisOf(dictHasAnalysis)(expr.value0)));
                };
                if (expr instanceof PureScript_Backend_Optimizer_Syntax.UncurriedEffectApp) {
                    var analysis = withResult(Unknown.value)(complex(NonTrivial.value)(analyzeDefault(dictHasAnalysis)(expr)));
                    var v = PureScript_Backend_Optimizer_Syntax.syntaxOf(dictHasSyntax)(expr.value0);
                    if (v instanceof Data_Maybe.Just && v.value0 instanceof PureScript_Backend_Optimizer_Syntax.Local) {
                        return callArity(v.value0.value1)(Data_Array.length(expr.value1))(analysis);
                    };
                    return analysis;
                };
                if (expr instanceof PureScript_Backend_Optimizer_Syntax.PrimEffect) {
                    return withResult(Unknown.value)(complex(NonTrivial.value)(analyzeDefault(dictHasAnalysis)(expr)));
                };
                if (expr instanceof PureScript_Backend_Optimizer_Syntax.Typed) {
                    return analysisOf(dictHasAnalysis)(expr.value1);
                };
                return analyze(dictHasAnalysis)(dictHasSyntax)(externAnalysis)(expr);
            };
        };
    };
};
export {
    analysisOf,
    CaptureNone,
    CaptureBranch,
    CaptureClosure,
    Usage,
    Trivial,
    Deref,
    KnownSize,
    NonTrivial,
    KnownNeutral,
    Unknown,
    BackendAnalysis,
    bound,
    boundArg,
    withArgs,
    withRewrite,
    used,
    accessed,
    cased,
    updated,
    usedDep,
    bump,
    complex,
    capture,
    callArity,
    withResult,
    externs,
    resultOf,
    analyze,
    analyzeEffectBlock,
    analyzeDefault,
    eqCapture,
    ordCapture,
    semigroupCapture,
    monoidCapture,
    newtypeUsage_,
    semigroupUsage,
    monoidUsage,
    eqComplexity,
    ordComplexity,
    semigroupComplexity,
    monoidComplexity,
    eqResultTerm,
    semigroupResultTerm,
    monoidResultTerm,
    newtypeBackendAnalysis_,
    semigroupBackendAnalysis,
    monoidBackendAnalysis
};
