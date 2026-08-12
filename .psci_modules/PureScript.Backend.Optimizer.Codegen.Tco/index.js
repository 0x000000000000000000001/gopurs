import * as Control_Alternative from "../Control.Alternative/index.js";
import * as Control_Applicative from "../Control.Applicative/index.js";
import * as Control_Bind from "../Control.Bind/index.js";
import * as Control_Category from "../Control.Category/index.js";
import * as Data_Array from "../Data.Array/index.js";
import * as Data_Array_NonEmpty from "../Data.Array.NonEmpty/index.js";
import * as Data_Array_NonEmpty_Internal from "../Data.Array.NonEmpty.Internal/index.js";
import * as Data_Boolean from "../Data.Boolean/index.js";
import * as Data_Eq from "../Data.Eq/index.js";
import * as Data_Foldable from "../Data.Foldable/index.js";
import * as Data_Function from "../Data.Function/index.js";
import * as Data_Functor from "../Data.Functor/index.js";
import * as Data_List from "../Data.List/index.js";
import * as Data_List_Types from "../Data.List.Types/index.js";
import * as Data_Map from "../Data.Map/index.js";
import * as Data_Map_Internal from "../Data.Map.Internal/index.js";
import * as Data_Maybe from "../Data.Maybe/index.js";
import * as Data_Monoid from "../Data.Monoid/index.js";
import * as Data_Ord from "../Data.Ord/index.js";
import * as Data_Ordering from "../Data.Ordering/index.js";
import * as Data_Semigroup from "../Data.Semigroup/index.js";
import * as Data_Semiring from "../Data.Semiring/index.js";
import * as Data_Set from "../Data.Set/index.js";
import * as Data_Traversable from "../Data.Traversable/index.js";
import * as Data_Tuple from "../Data.Tuple/index.js";
import * as Data_Unfoldable from "../Data.Unfoldable/index.js";
import * as PureScript_Backend_Optimizer_CoreFn from "../PureScript.Backend.Optimizer.CoreFn/index.js";
import * as PureScript_Backend_Optimizer_Syntax from "../PureScript.Backend.Optimizer.Syntax/index.js";
var eqQualified = /* #__PURE__ */ PureScript_Backend_Optimizer_CoreFn.eqQualified(PureScript_Backend_Optimizer_CoreFn.eqIdent);
var eqMaybe = /* #__PURE__ */ Data_Maybe.eqMaybe(PureScript_Backend_Optimizer_CoreFn.eqIdent);
var ordQualified = /* #__PURE__ */ PureScript_Backend_Optimizer_CoreFn.ordQualified(PureScript_Backend_Optimizer_CoreFn.ordIdent);
var ordMaybe = /* #__PURE__ */ Data_Maybe.ordMaybe(PureScript_Backend_Optimizer_CoreFn.ordIdent);
var fold = /* #__PURE__ */ Data_Foldable.fold(Data_Foldable.foldableMaybe)(Data_Monoid.monoidArray);
var add = /* #__PURE__ */ Data_Semiring.add(Data_Semiring.semiringInt);
var identity = /* #__PURE__ */ Control_Category.identity(Control_Category.categoryFn);
var guard = /* #__PURE__ */ Control_Alternative.guard(Data_Maybe.alternativeMaybe);
var pure = /* #__PURE__ */ Control_Applicative.pure(Data_Array_NonEmpty_Internal.applicativeNonEmptyArray);
var TcoUsage = function (x) {
    return x;
};
var TcoTopLevel = /* #__PURE__ */ (function () {
    function TcoTopLevel(value0) {
        this.value0 = value0;
    };
    TcoTopLevel.create = function (value0) {
        return new TcoTopLevel(value0);
    };
    return TcoTopLevel;
})();
var TcoLocal = /* #__PURE__ */ (function () {
    function TcoLocal(value0, value1) {
        this.value0 = value0;
        this.value1 = value1;
    };
    TcoLocal.create = function (value0) {
        return function (value1) {
            return new TcoLocal(value0, value1);
        };
    };
    return TcoLocal;
})();
var TcoAnalysis = function (x) {
    return x;
};
var TcoExpr = /* #__PURE__ */ (function () {
    function TcoExpr(value0, value1) {
        this.value0 = value0;
        this.value1 = value1;
    };
    TcoExpr.create = function (value0) {
        return function (value1) {
            return new TcoExpr(value0, value1);
        };
    };
    return TcoExpr;
})();
var semigroupTcoUsage = {
    append: function (v) {
        return function (v1) {
            return {
                total: v.total + v1.total | 0,
                arities: Data_Set.union(Data_Ord.ordInt)(v.arities)(v1.arities),
                call: v.call + v1.call | 0,
                readWrite: v.readWrite + v1.readWrite | 0
            };
        };
    }
};
var append = /* #__PURE__ */ Data_Semigroup.append(semigroupTcoUsage);
var newtypeTcoAnalysis_ = {
    Coercible0: function () {
        return undefined;
    }
};
var monoidTcoUsage = {
    mempty: {
        total: 0,
        arities: Data_Set.empty,
        call: 0,
        readWrite: 0
    },
    Semigroup0: function () {
        return semigroupTcoUsage;
    }
};
var eqTcoRef = {
    eq: function (x) {
        return function (y) {
            if (x instanceof TcoTopLevel && y instanceof TcoTopLevel) {
                return Data_Eq.eq(eqQualified)(x.value0)(y.value0);
            };
            if (x instanceof TcoLocal && y instanceof TcoLocal) {
                return Data_Eq.eq(eqMaybe)(x.value0)(y.value0) && Data_Eq.eq(PureScript_Backend_Optimizer_Syntax.eqLevel)(x.value1)(y.value1);
            };
            return false;
        };
    }
};
var ordTcoRef = {
    compare: function (x) {
        return function (y) {
            if (x instanceof TcoTopLevel && y instanceof TcoTopLevel) {
                return Data_Ord.compare(ordQualified)(x.value0)(y.value0);
            };
            if (x instanceof TcoTopLevel) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof TcoTopLevel) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof TcoLocal && y instanceof TcoLocal) {
                var v = Data_Ord.compare(ordMaybe)(x.value0)(y.value0);
                if (v instanceof Data_Ordering.LT) {
                    return Data_Ordering.LT.value;
                };
                if (v instanceof Data_Ordering.GT) {
                    return Data_Ordering.GT.value;
                };
                return Data_Ord.compare(PureScript_Backend_Optimizer_Syntax.ordLevel)(x.value1)(y.value1);
            };
            throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Codegen.Tco (line 0, column 0 - line 0, column 0): " + [ x.constructor.name, y.constructor.name ]);
        };
    },
    Eq0: function () {
        return eqTcoRef;
    }
};
var nub = /* #__PURE__ */ Data_Array.nub(ordTcoRef);
var withTcoRole = function (role) {
    return function (v) {
        return {
            usages: v.usages,
            tailCalls: v.tailCalls,
            role: role
        };
    };
};
var usedTopLevel = function (v) {
    return Data_Set.mapMaybe(ordQualified)(function (v1) {
        if (v1 instanceof TcoTopLevel) {
            return new Data_Maybe.Just(v1.value0);
        };
        if (v1 instanceof TcoLocal) {
            return Data_Maybe.Nothing.value;
        };
        throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Codegen.Tco (line 133, column 18 - line 135, column 30): " + [ v1.constructor.name ]);
    })(Data_Map.keys(v.usages));
};
var unwrapRefHead = function ($copy_v) {
    var $tco_done = false;
    var $tco_result;
    function $tco_loop(v) {
        if (v instanceof PureScript_Backend_Optimizer_Syntax.Local) {
            $tco_done = true;
            return new Data_Maybe.Just(new TcoLocal(v.value0, v.value1));
        };
        if (v instanceof PureScript_Backend_Optimizer_Syntax.Typed) {
            $copy_v = v.value1;
            return;
        };
        $tco_done = true;
        return Data_Maybe.Nothing.value;
    };
    while (!$tco_done) {
        $tco_result = $tco_loop($copy_v);
    };
    return $tco_result;
};
var unwrapAppHead = function ($copy_v) {
    return function ($copy_arity) {
        var $tco_var_v = $copy_v;
        var $tco_done = false;
        var $tco_result;
        function $tco_loop(v, arity) {
            if (v instanceof PureScript_Backend_Optimizer_Syntax.Local) {
                $tco_done = true;
                return new Data_Maybe.Just(new Data_Tuple.Tuple(new TcoLocal(v.value0, v.value1), arity));
            };
            if (v instanceof PureScript_Backend_Optimizer_Syntax.Var) {
                $tco_done = true;
                return new Data_Maybe.Just(new Data_Tuple.Tuple(new TcoTopLevel(v.value0), arity));
            };
            if (v instanceof PureScript_Backend_Optimizer_Syntax.Typed) {
                $tco_var_v = v.value1;
                $copy_arity = arity;
                return;
            };
            if (v instanceof PureScript_Backend_Optimizer_Syntax.App) {
                $tco_var_v = v.value0;
                $copy_arity = arity + Data_Array_NonEmpty.length(v.value1) | 0;
                return;
            };
            if (v instanceof PureScript_Backend_Optimizer_Syntax.UncurriedApp) {
                $tco_var_v = v.value0;
                $copy_arity = arity + Data_Array.length(v.value1) | 0;
                return;
            };
            if (v instanceof PureScript_Backend_Optimizer_Syntax.UncurriedEffectApp) {
                $tco_var_v = v.value0;
                $copy_arity = arity + Data_Array.length(v.value1) | 0;
                return;
            };
            $tco_done = true;
            return Data_Maybe.Nothing.value;
        };
        while (!$tco_done) {
            $tco_result = $tco_loop($tco_var_v, $copy_arity);
        };
        return $tco_result;
    };
};
var unwindTcoScope = /* #__PURE__ */ (function () {
    var go = function ($copy_pop) {
        return function ($copy_v) {
            var $tco_var_pop = $copy_pop;
            var $tco_done = false;
            var $tco_result;
            function $tco_loop(pop, v) {
                if (v instanceof Data_List_Types.Cons) {
                    if (Data_List["null"](v.value1)) {
                        $tco_done = true;
                        return new Data_Maybe.Just(new Data_Tuple.Tuple(v.value0.value0, pop));
                    };
                    if (Data_Boolean.otherwise) {
                        $tco_var_pop = new Data_List_Types.Cons(v.value0.value0, pop);
                        $copy_v = v.value1;
                        return;
                    };
                };
                if (v instanceof Data_List_Types.Nil) {
                    $tco_done = true;
                    return Data_Maybe.Nothing.value;
                };
                throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Codegen.Tco (line 69, column 12 - line 76, column 14): " + [ v.constructor.name ]);
            };
            while (!$tco_done) {
                $tco_result = $tco_loop($tco_var_pop, $copy_v);
            };
            return $tco_result;
        };
    };
    return go(Data_List_Types.Nil.value);
})();
var unTcoExpr = function (v) {
    return v.value1;
};
var tcoRefEffect = function (ident) {
    return function (v) {
        return {
            tailCalls: v.tailCalls,
            role: v.role,
            usages: Data_Map_Internal.insertWith(ordTcoRef)(append)(ident)({
                total: 1,
                call: 0,
                arities: Data_Set.empty,
                readWrite: 1
            })(v.usages)
        };
    };
};
var tcoNoTailCalls = function (v) {
    return {
        usages: v.usages,
        tailCalls: Data_Map_Internal.empty,
        role: {
            isLoop: v.role.isLoop,
            joins: [  ]
        }
    };
};
var tcoCall = function (ident) {
    return function (arity) {
        return function (v) {
            return {
                role: v.role,
                usages: Data_Map_Internal.insertWith(ordTcoRef)(append)(ident)({
                    total: 1,
                    call: 1,
                    arities: Data_Set.singleton(arity),
                    readWrite: 0
                })(v.usages),
                tailCalls: Data_Map_Internal.insert(ordTcoRef)(ident)(1)(v.tailCalls)
            };
        };
    };
};
var tcoAnalysisOf = function (v) {
    return v.value0;
};
var tcoRefBinding = function (ref) {
    return function (v) {
        if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.Abs) {
            var v1 = tcoRefBinding(ref)(v.value1.value1);
            if (v1 instanceof Data_Maybe.Just) {
                return new Data_Maybe.Just({
                    ref: ref,
                    analysis: v1.value0.analysis,
                    arity: Data_Array_NonEmpty.length(v.value1.value0) + v1.value0.arity | 0
                });
            };
            if (v1 instanceof Data_Maybe.Nothing) {
                return new Data_Maybe.Just({
                    ref: ref,
                    analysis: tcoAnalysisOf(v.value1.value1),
                    arity: Data_Array_NonEmpty.length(v.value1.value0)
                });
            };
            throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Codegen.Tco (line 193, column 5 - line 195, column 95): " + [ v1.constructor.name ]);
        };
        if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.UncurriedAbs) {
            var v1 = tcoRefBinding(ref)(v.value1.value1);
            if (v1 instanceof Data_Maybe.Just) {
                return new Data_Maybe.Just({
                    ref: ref,
                    analysis: v1.value0.analysis,
                    arity: Data_Array.length(v.value1.value0) + v1.value0.arity | 0
                });
            };
            if (v1 instanceof Data_Maybe.Nothing) {
                return new Data_Maybe.Just({
                    ref: ref,
                    analysis: tcoAnalysisOf(v.value1.value1),
                    arity: Data_Array.length(v.value1.value0)
                });
            };
            throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Codegen.Tco (line 197, column 5 - line 199, column 87): " + [ v1.constructor.name ]);
        };
        if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.Typed) {
            return tcoRefBinding(ref)(v.value1.value1);
        };
        return Data_Maybe.Nothing.value;
    };
};
var tcoRefBindings = function (toTcoRef) {
    return Data_Traversable.traverse(Data_Array_NonEmpty_Internal.traversableNonEmptyArray)(Data_Maybe.applicativeMaybe)(function (v) {
        return tcoRefBinding(toTcoRef(v.value0))(v.value1);
    });
};
var topLevelTcoRefBindings = function (mod) {
    return tcoRefBindings((function () {
        var $270 = PureScript_Backend_Optimizer_CoreFn.Qualified.create(new Data_Maybe.Just(mod));
        return function ($271) {
            return TcoTopLevel.create($270($271));
        };
    })());
};
var syntacticArity = function (v) {
    if (v instanceof PureScript_Backend_Optimizer_Syntax.Abs) {
        var v1 = syntacticArity(v.value1);
        if (v1 instanceof Data_Maybe.Just) {
            return new Data_Maybe.Just(Data_Array_NonEmpty.length(v.value0) + v1.value0 | 0);
        };
        if (v1 instanceof Data_Maybe.Nothing) {
            return new Data_Maybe.Just(Data_Array_NonEmpty.length(v.value0));
        };
        throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Codegen.Tco (line 241, column 5 - line 243, column 50): " + [ v1.constructor.name ]);
    };
    if (v instanceof PureScript_Backend_Optimizer_Syntax.UncurriedAbs) {
        var v1 = syntacticArity(v.value1);
        if (v1 instanceof Data_Maybe.Just) {
            return new Data_Maybe.Just(Data_Array.length(v.value0) + v1.value0 | 0);
        };
        if (v1 instanceof Data_Maybe.Nothing) {
            return new Data_Maybe.Just(Data_Array.length(v.value0));
        };
        throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Codegen.Tco (line 245, column 5 - line 247, column 42): " + [ v1.constructor.name ]);
    };
    if (v instanceof PureScript_Backend_Optimizer_Syntax.Typed) {
        return syntacticArity(v.value1);
    };
    return Data_Maybe.Nothing.value;
};
var tcoEnvGroup = function (toTcoRef) {
    var go = function (v) {
        return Data_Functor.map(Data_Maybe.functorMaybe)(Data_Tuple.Tuple.create(toTcoRef(v.value0)))(syntacticArity(v.value1));
    };
    var $272 = Data_Traversable.traverse(Data_Traversable.traversableArray)(Data_Maybe.applicativeMaybe)(go);
    return function ($273) {
        return fold($272(Data_Array_NonEmpty.toArray($273)));
    };
};
var topLevelTcoEnvGroup = function (mod) {
    return tcoEnvGroup((function () {
        var $274 = PureScript_Backend_Optimizer_CoreFn.Qualified.create(new Data_Maybe.Just(mod));
        return function ($275) {
            return TcoTopLevel.create($274($275));
        };
    })());
};
var popTcoScope = function (ref) {
    var go = function ($copy_stack) {
        return function ($copy_v) {
            var $tco_var_stack = $copy_stack;
            var $tco_done = false;
            var $tco_result;
            function $tco_loop(stack, v) {
                var v1 = function (v2) {
                    if (v instanceof Data_List_Types.Cons && Data_Boolean.otherwise) {
                        $tco_var_stack = new Data_List_Types.Cons(v.value0.value0, stack);
                        $copy_v = v.value1;
                        return;
                    };
                    $tco_done = true;
                    return Data_Maybe.Nothing.value;
                };
                if (v instanceof Data_List_Types.Cons) {
                    var $192 = Data_Array_NonEmpty.findIndex(Data_Eq.eq(eqTcoRef)(ref))(v.value0.value1);
                    if ($192 instanceof Data_Maybe.Just) {
                        $tco_done = true;
                        return new Data_Maybe.Just({
                            ident: v.value0.value0,
                            group: v.value0.value1,
                            index: $192.value0,
                            stack: stack
                        });
                    };
                    return v1(true);
                };
                return v1(true);
            };
            while (!$tco_done) {
                $tco_result = $tco_loop($tco_var_stack, $copy_v);
            };
            return $tco_result;
        };
    };
    return go(Data_List_Types.Nil.value);
};
var overTcoAnalysis = function (f) {
    return function (v) {
        return new TcoExpr(f(v.value0), v.value1);
    };
};
var noTcoRole = {
    joins: [  ],
    isLoop: false
};
var semigroupTcoAnalysis = {
    append: function (v) {
        return function (v1) {
            return {
                usages: Data_Map_Internal.unionWith(ordTcoRef)(append)(v.usages)(v1.usages),
                tailCalls: Data_Map_Internal.unionWith(ordTcoRef)(add)(v.tailCalls)(v1.tailCalls),
                role: noTcoRole
            };
        };
    }
};
var monoidTcoAnalysis = {
    mempty: {
        usages: Data_Map_Internal.empty,
        tailCalls: Data_Map_Internal.empty,
        role: noTcoRole
    },
    Semigroup0: function () {
        return semigroupTcoAnalysis;
    }
};
var mempty = /* #__PURE__ */ Data_Monoid.mempty(monoidTcoAnalysis);
var localTcoRefBindings = function (level) {
    return tcoRefBindings(function (ident) {
        return new TcoLocal(new Data_Maybe.Just(ident), level);
    });
};
var localTcoEnvGroup = function (level) {
    return tcoEnvGroup(function (ident) {
        return new TcoLocal(new Data_Maybe.Just(ident), level);
    });
};
var isUniformTailCall = function (v) {
    return function (ref) {
        return function (arity) {
            return Control_Bind.bind(Data_Maybe.bindMaybe)(Data_Map_Internal.lookup(ordTcoRef)(ref)(v.tailCalls))(function (numTailCalls) {
                return Control_Bind.bind(Data_Maybe.bindMaybe)(Data_Map_Internal.lookup(ordTcoRef)(ref)(v.usages))(function (v1) {
                    var v2 = Data_Set.toUnfoldable(Data_Unfoldable.unfoldableArray)(v1.arities);
                    if (v2.length === 1) {
                        return new Data_Maybe.Just(v2[0] === arity && v1.total === numTailCalls);
                    };
                    return Data_Maybe.Nothing.value;
                });
            });
        };
    };
};
var isTailCalledIn = function (analysis) {
    return function (group) {
        var tailCalled = Data_Array_NonEmpty.mapMaybe(function (b) {
            return isUniformTailCall(analysis)(b.ref)(b.arity);
        })(group);
        return !Data_Array["null"](tailCalled) && Data_Array.all(identity)(tailCalled);
    };
};
var tcoRoleIsLoop = function (group) {
    return Data_Array_NonEmpty.all((function () {
        var $276 = Data_Function.flip(isTailCalledIn)(group);
        return function ($277) {
            return $276((function (v) {
                return v.analysis;
            })($277));
        };
    })())(group);
};
var tcoRoleJoins = function (env) {
    return function (analysis) {
        return function (group) {
            return Control_Bind.discard(Control_Bind.discardUnit)(Control_Bind.bindArray)(Control_Alternative.guard(Control_Alternative.alternativeArray)(isTailCalledIn(analysis)(group)))(function () {
                return nub(Data_Foldable.foldMap(Data_Array_NonEmpty_Internal.foldableNonEmptyArray)(Data_Monoid.monoidArray)(function (b) {
                    return Data_Array.mapMaybe(function (v) {
                        return Data_Functor.voidRight(Data_Maybe.functorMaybe)(v.value0)(Control_Bind.bindFlipped(Data_Maybe.bindMaybe)(guard)(isUniformTailCall(b.analysis)(v.value0)(v.value1)));
                    })(env);
                })(group));
            });
        };
    };
};
var inTcoScope = function (ref) {
    var go = function ($copy_v) {
        var $tco_done = false;
        var $tco_result;
        function $tco_loop(v) {
            if (v instanceof Data_List_Types.Cons) {
                if (Data_Array_NonEmpty.elem(eqTcoRef)(ref)(v.value0.value1)) {
                    $tco_done = true;
                    return true;
                };
                if (Data_Boolean.otherwise) {
                    $copy_v = v.value1;
                    return;
                };
            };
            $tco_done = true;
            return false;
        };
        while (!$tco_done) {
            $tco_result = $tco_loop($copy_v);
        };
        return $tco_result;
    };
    return go;
};
var hasTcoRole = function (v) {
    return v.isLoop || !Data_Array["null"](v.joins);
};
var analyze = function (env) {
    return function (v) {
        if (v instanceof PureScript_Backend_Optimizer_Syntax.Var) {
            return new TcoExpr(tcoCall(new TcoTopLevel(v.value0))(0)(mempty), new PureScript_Backend_Optimizer_Syntax.Var(v.value0));
        };
        if (v instanceof PureScript_Backend_Optimizer_Syntax.Local) {
            return new TcoExpr(tcoCall(new TcoLocal(v.value0, v.value1))(0)(mempty), new PureScript_Backend_Optimizer_Syntax.Local(v.value0, v.value1));
        };
        var v1 = function (v2) {
            var v3 = function (v4) {
                var v5 = function (v6) {
                    var v7 = function (v8) {
                        if (v instanceof PureScript_Backend_Optimizer_Syntax.Typed) {
                            var inner$prime = analyze(env)(v.value1);
                            return new TcoExpr(tcoAnalysisOf(inner$prime), new PureScript_Backend_Optimizer_Syntax.Typed(v.value0, inner$prime));
                        };
                        if (v instanceof PureScript_Backend_Optimizer_Syntax.Branch) {
                            var branches$prime = Data_Functor.map(Data_Array_NonEmpty_Internal.functorNonEmptyArray)(function (v9) {
                                return new PureScript_Backend_Optimizer_Syntax.Pair(overTcoAnalysis(tcoNoTailCalls)(analyze(env)(v9.value0)), analyze(env)(v9.value1));
                            })(v.value0);
                            var def$prime = analyze(env)(v.value1);
                            var analysis = Data_Semigroup.append(semigroupTcoAnalysis)(Data_Foldable.foldMap(Data_Array_NonEmpty_Internal.foldableNonEmptyArray)(monoidTcoAnalysis)(Data_Foldable.foldMap(PureScript_Backend_Optimizer_Syntax.foldablePair)(monoidTcoAnalysis)(tcoAnalysisOf))(branches$prime))(tcoAnalysisOf(def$prime));
                            return new TcoExpr(analysis, new PureScript_Backend_Optimizer_Syntax.Branch(branches$prime, def$prime));
                        };
                        if (v instanceof PureScript_Backend_Optimizer_Syntax.Let) {
                            var binding$prime = analyze(env)(v.value2);
                            var body$prime = analyze(env)(v.value3);
                            var expr$prime = new PureScript_Backend_Optimizer_Syntax.Let(v.value0, v.value1, binding$prime, body$prime);
                            var refBinding = tcoRefBinding(new TcoLocal(v.value0, v.value1))(binding$prime);
                            var role = {
                                isLoop: false,
                                joins: Data_Foldable.foldMap(Data_Foldable.foldableMaybe)(Data_Monoid.monoidArray)((function () {
                                    var $278 = tcoRoleJoins(env)(tcoAnalysisOf(body$prime));
                                    return function ($279) {
                                        return $278(pure($279));
                                    };
                                })())(refBinding)
                            };
                            if (refBinding instanceof Data_Maybe.Just && hasTcoRole(role)) {
                                var analysis = Data_Semigroup.append(semigroupTcoAnalysis)(refBinding.value0.analysis)(tcoAnalysisOf(body$prime));
                                return new TcoExpr(withTcoRole(role)(analysis), expr$prime);
                            };
                            var analysis = Data_Semigroup.append(semigroupTcoAnalysis)(tcoNoTailCalls(tcoAnalysisOf(binding$prime)))(tcoAnalysisOf(body$prime));
                            return new TcoExpr(analysis, expr$prime);
                        };
                        if (v instanceof PureScript_Backend_Optimizer_Syntax.LetRec) {
                            var group = localTcoEnvGroup(v.value0)(v.value1);
                            var env$prime = Data_Semigroup.append(Data_Semigroup.semigroupArray)(group)(env);
                            var bindings$prime = Data_Functor.map(Data_Array_NonEmpty_Internal.functorNonEmptyArray)(Data_Functor.map(Data_Tuple.functorTuple)(analyze(env$prime)))(v.value1);
                            var body$prime = analyze(env$prime)(v.value2);
                            var expr$prime = new PureScript_Backend_Optimizer_Syntax.LetRec(v.value0, bindings$prime, body$prime);
                            var refBindings = localTcoRefBindings(v.value0)(bindings$prime);
                            var isLoop = Data_Maybe.maybe(false)(tcoRoleIsLoop)(refBindings);
                            var role = {
                                isLoop: isLoop,
                                joins: Data_Foldable.foldMap(Data_Foldable.foldableMaybe)(Data_Monoid.monoidArray)(tcoRoleJoins(env)(tcoAnalysisOf(body$prime)))(refBindings)
                            };
                            var $241 = hasTcoRole(role);
                            if ($241) {
                                var analysis = Data_Semigroup.append(semigroupTcoAnalysis)(Data_Foldable.foldMap(Data_Foldable.foldableMaybe)(monoidTcoAnalysis)(Data_Foldable.foldMap(Data_Array_NonEmpty_Internal.foldableNonEmptyArray)(monoidTcoAnalysis)(function (v9) {
                                    return v9.analysis;
                                }))(refBindings))(tcoAnalysisOf(body$prime));
                                return new TcoExpr(withTcoRole(role)(analysis), expr$prime);
                            };
                            var analysis = Data_Semigroup.append(semigroupTcoAnalysis)(tcoNoTailCalls(Data_Foldable.foldMap(Data_Array_NonEmpty_Internal.foldableNonEmptyArray)(monoidTcoAnalysis)(Data_Foldable.foldMap(Data_Foldable.foldableTuple)(monoidTcoAnalysis)(tcoAnalysisOf))(bindings$prime)))(tcoAnalysisOf(body$prime));
                            return new TcoExpr(analysis, expr$prime);
                        };
                        var expr$prime = Data_Functor.map(PureScript_Backend_Optimizer_Syntax.functorBackendSyntax)(analyze(env))(v);
                        return new TcoExpr(tcoNoTailCalls(Data_Foldable.foldMap(PureScript_Backend_Optimizer_Syntax.foldableBackendSyntax)(monoidTcoAnalysis)(tcoAnalysisOf)(expr$prime)), expr$prime);
                    };
                    if (v instanceof PureScript_Backend_Optimizer_Syntax.PrimEffect && v.value0 instanceof PureScript_Backend_Optimizer_Syntax.EffectRefWrite) {
                        var $246 = unwrapRefHead(v.value0.value0);
                        if ($246 instanceof Data_Maybe.Just) {
                            var ref$prime = analyze(env)(v.value0.value0);
                            var val$prime = analyze(env)(v.value0.value1);
                            var analysis = tcoRefEffect($246.value0)(tcoAnalysisOf(val$prime));
                            return new TcoExpr(analysis, new PureScript_Backend_Optimizer_Syntax.PrimEffect(new PureScript_Backend_Optimizer_Syntax.EffectRefWrite(ref$prime, val$prime)));
                        };
                        return v7(true);
                    };
                    return v7(true);
                };
                if (v instanceof PureScript_Backend_Optimizer_Syntax.PrimEffect && v.value0 instanceof PureScript_Backend_Optimizer_Syntax.EffectRefRead) {
                    var $252 = unwrapRefHead(v.value0.value0);
                    if ($252 instanceof Data_Maybe.Just) {
                        var ref$prime = analyze(env)(v.value0.value0);
                        var analysis = tcoRefEffect($252.value0)(mempty);
                        return new TcoExpr(analysis, new PureScript_Backend_Optimizer_Syntax.PrimEffect(new PureScript_Backend_Optimizer_Syntax.EffectRefRead(ref$prime)));
                    };
                    return v5(true);
                };
                return v5(true);
            };
            if (v instanceof PureScript_Backend_Optimizer_Syntax.UncurriedApp) {
                var $257 = unwrapAppHead(v.value0)(Data_Array.length(v.value1));
                if ($257 instanceof Data_Maybe.Just) {
                    var hd$prime = analyze(env)(v.value0);
                    var tl$prime = Data_Functor.map(Data_Functor.functorArray)((function () {
                        var $280 = overTcoAnalysis(tcoNoTailCalls);
                        var $281 = analyze(env);
                        return function ($282) {
                            return $280($281($282));
                        };
                    })())(v.value1);
                    var analysis2 = tcoCall($257.value0.value0)($257.value0.value1)(Data_Foldable.foldMap(Data_Foldable.foldableArray)(monoidTcoAnalysis)(tcoAnalysisOf)(tl$prime));
                    return new TcoExpr(analysis2, new PureScript_Backend_Optimizer_Syntax.UncurriedApp(hd$prime, tl$prime));
                };
                return v3(true);
            };
            return v3(true);
        };
        if (v instanceof PureScript_Backend_Optimizer_Syntax.App) {
            var $264 = unwrapAppHead(v.value0)(Data_Array_NonEmpty.length(v.value1));
            if ($264 instanceof Data_Maybe.Just) {
                var hd$prime = analyze(env)(v.value0);
                var tl$prime = Data_Functor.map(Data_Array_NonEmpty_Internal.functorNonEmptyArray)((function () {
                    var $283 = overTcoAnalysis(tcoNoTailCalls);
                    var $284 = analyze(env);
                    return function ($285) {
                        return $283($284($285));
                    };
                })())(v.value1);
                var analysis2 = tcoCall($264.value0.value0)($264.value0.value1)(Data_Foldable.foldMap(Data_Array_NonEmpty_Internal.foldableNonEmptyArray)(monoidTcoAnalysis)(tcoAnalysisOf)(tl$prime));
                return new TcoExpr(analysis2, new PureScript_Backend_Optimizer_Syntax.App(hd$prime, tl$prime));
            };
            return v1(true);
        };
        return v1(true);
    };
};
export {
    TcoTopLevel,
    TcoLocal,
    popTcoScope,
    inTcoScope,
    unwindTcoScope,
    TcoUsage,
    TcoAnalysis,
    usedTopLevel,
    noTcoRole,
    hasTcoRole,
    TcoExpr,
    unTcoExpr,
    tcoAnalysisOf,
    overTcoAnalysis,
    tcoCall,
    tcoRefEffect,
    tcoNoTailCalls,
    withTcoRole,
    isUniformTailCall,
    tcoRefBinding,
    tcoRefBindings,
    tcoEnvGroup,
    localTcoRefBindings,
    localTcoEnvGroup,
    topLevelTcoRefBindings,
    topLevelTcoEnvGroup,
    isTailCalledIn,
    tcoRoleIsLoop,
    tcoRoleJoins,
    syntacticArity,
    analyze,
    unwrapAppHead,
    unwrapRefHead,
    eqTcoRef,
    ordTcoRef,
    semigroupTcoUsage,
    monoidTcoUsage,
    newtypeTcoAnalysis_,
    semigroupTcoAnalysis,
    monoidTcoAnalysis
};
