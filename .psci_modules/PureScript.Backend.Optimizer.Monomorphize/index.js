import * as Data_Array from "../Data.Array/index.js";
import * as Data_Foldable from "../Data.Foldable/index.js";
import * as Data_Function from "../Data.Function/index.js";
import * as Data_Functor from "../Data.Functor/index.js";
import * as Data_Map_Internal from "../Data.Map.Internal/index.js";
import * as Data_Maybe from "../Data.Maybe/index.js";
import * as Data_Newtype from "../Data.Newtype/index.js";
import * as Data_Ord from "../Data.Ord/index.js";
import * as Data_Semigroup from "../Data.Semigroup/index.js";
import * as Data_Set from "../Data.Set/index.js";
import * as Data_String_Common from "../Data.String.Common/index.js";
import * as Data_Tuple from "../Data.Tuple/index.js";
import * as Data_Unfoldable from "../Data.Unfoldable/index.js";
import * as PureScript_Backend_Optimizer_CoreFn from "../PureScript.Backend.Optimizer.CoreFn/index.js";
import * as PureScript_Backend_Optimizer_Substitute from "../PureScript.Backend.Optimizer.Substitute/index.js";
var union = /* #__PURE__ */ Data_Map_Internal.union(PureScript_Backend_Optimizer_CoreFn.ordExprType);
var rebuildApp = function (finalAnn) {
    return function (f) {
        return function (args) {
            var v = Data_Array.uncons(args);
            if (v instanceof Data_Maybe.Nothing) {
                return f;
            };
            if (v instanceof Data_Maybe.Just) {
                var firstApp = new PureScript_Backend_Optimizer_CoreFn.ExprApp(finalAnn, f, v.value0.head);
                return Data_Array.foldl(function (acc) {
                    return function (arg) {
                        return new PureScript_Backend_Optimizer_CoreFn.ExprApp(finalAnn, acc, arg);
                    };
                })(firstApp)(v.value0.tail);
            };
            throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Monomorphize (line 366, column 3 - line 370, column 74): " + [ v.constructor.name ]);
        };
    };
};
var partitionArgs = function ($copy_v) {
    return function ($copy_v1) {
        var $tco_var_v = $copy_v;
        var $tco_done = false;
        var $tco_result;
        function $tco_loop(v, v1) {
            if (v instanceof PureScript_Backend_Optimizer_CoreFn.ConstrainedType) {
                var numDicts = Data_Array.length(v.value0);
                var normalArgs = Data_Array.drop(numDicts)(v1);
                var dictArgs = Data_Array.take(numDicts)(v1);
                $tco_done = true;
                return {
                    dictArgs: dictArgs,
                    normalArgs: normalArgs
                };
            };
            if (v instanceof PureScript_Backend_Optimizer_CoreFn.ForAll) {
                $tco_var_v = v.value1;
                $copy_v1 = v1;
                return;
            };
            $tco_done = true;
            return {
                dictArgs: [  ],
                normalArgs: v1
            };
        };
        while (!$tco_done) {
            $tco_result = $tco_loop($tco_var_v, $copy_v1);
        };
        return $tco_result;
    };
};
var mapAnn = function (f) {
    return function (v) {
        return {
            span: v.span,
            meta: v.meta,
            type: Data_Functor.map(Data_Maybe.functorMaybe)(f)(v.type)
        };
    };
};
var rewriteExpr = function (f) {
    var goProp = function (v) {
        return new PureScript_Backend_Optimizer_CoreFn.Prop(v.value0, go(v.value1));
    };
    var goGuard = function (v) {
        return new PureScript_Backend_Optimizer_CoreFn.Guard(go(v.value0), go(v.value1));
    };
    var goCaseGuard = function (v) {
        if (v instanceof PureScript_Backend_Optimizer_CoreFn.Unconditional) {
            return new PureScript_Backend_Optimizer_CoreFn.Unconditional(go(v.value0));
        };
        if (v instanceof PureScript_Backend_Optimizer_CoreFn.Guarded) {
            return new PureScript_Backend_Optimizer_CoreFn.Guarded(Data_Functor.map(Data_Functor.functorArray)(goGuard)(v.value0));
        };
        throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Monomorphize (line 236, column 3 - line 236, column 55): " + [ v.constructor.name ]);
    };
    var goBinding = function (v) {
        return new PureScript_Backend_Optimizer_CoreFn.Binding(mapAnn(f)(v.value0), v.value1, go(v.value2));
    };
    var goBind = function (v) {
        if (v instanceof PureScript_Backend_Optimizer_CoreFn.NonRec) {
            return new PureScript_Backend_Optimizer_CoreFn.NonRec(goBinding(v.value0));
        };
        if (v instanceof PureScript_Backend_Optimizer_CoreFn.Rec) {
            return new PureScript_Backend_Optimizer_CoreFn.Rec(Data_Functor.map(Data_Functor.functorArray)(goBinding)(v.value0));
        };
        throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Monomorphize (line 229, column 3 - line 229, column 43): " + [ v.constructor.name ]);
    };
    var goAlt = function (v) {
        return new PureScript_Backend_Optimizer_CoreFn.CaseAlternative(v.value0, goCaseGuard(v.value1));
    };
    var go = function (expr) {
        if (expr instanceof PureScript_Backend_Optimizer_CoreFn.ExprVar) {
            return new PureScript_Backend_Optimizer_CoreFn.ExprVar(mapAnn(f)(expr.value0), expr.value1);
        };
        if (expr instanceof PureScript_Backend_Optimizer_CoreFn.ExprLit) {
            return new PureScript_Backend_Optimizer_CoreFn.ExprLit(mapAnn(f)(expr.value0), Data_Functor.map(PureScript_Backend_Optimizer_CoreFn.functorLiteral)(go)(expr.value1));
        };
        if (expr instanceof PureScript_Backend_Optimizer_CoreFn.ExprApp) {
            return new PureScript_Backend_Optimizer_CoreFn.ExprApp(mapAnn(f)(expr.value0), go(expr.value1), go(expr.value2));
        };
        if (expr instanceof PureScript_Backend_Optimizer_CoreFn.ExprAbs) {
            return new PureScript_Backend_Optimizer_CoreFn.ExprAbs(mapAnn(f)(expr.value0), expr.value1, go(expr.value2));
        };
        if (expr instanceof PureScript_Backend_Optimizer_CoreFn.ExprLet) {
            return new PureScript_Backend_Optimizer_CoreFn.ExprLet(mapAnn(f)(expr.value0), Data_Functor.map(Data_Functor.functorArray)(goBind)(expr.value1), go(expr.value2));
        };
        if (expr instanceof PureScript_Backend_Optimizer_CoreFn.ExprCase) {
            return new PureScript_Backend_Optimizer_CoreFn.ExprCase(mapAnn(f)(expr.value0), Data_Functor.map(Data_Functor.functorArray)(go)(expr.value1), Data_Functor.map(Data_Functor.functorArray)(goAlt)(expr.value2));
        };
        if (expr instanceof PureScript_Backend_Optimizer_CoreFn.ExprConstructor) {
            return new PureScript_Backend_Optimizer_CoreFn.ExprConstructor(mapAnn(f)(expr.value0), expr.value1, expr.value2, expr.value3);
        };
        if (expr instanceof PureScript_Backend_Optimizer_CoreFn.ExprAccessor) {
            return new PureScript_Backend_Optimizer_CoreFn.ExprAccessor(mapAnn(f)(expr.value0), go(expr.value1), expr.value2);
        };
        if (expr instanceof PureScript_Backend_Optimizer_CoreFn.ExprUpdate) {
            return new PureScript_Backend_Optimizer_CoreFn.ExprUpdate(mapAnn(f)(expr.value0), go(expr.value1), Data_Functor.map(Data_Functor.functorArray)(goProp)(expr.value2));
        };
        throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Monomorphize (line 218, column 13 - line 227, column 82): " + [ expr.constructor.name ]);
    };
    return go;
};
var mangleType = function (v) {
    if (v instanceof PureScript_Backend_Optimizer_CoreFn.Int) {
        return "Int";
    };
    if (v instanceof PureScript_Backend_Optimizer_CoreFn["Number"]) {
        return "Number";
    };
    if (v instanceof PureScript_Backend_Optimizer_CoreFn["String"]) {
        return "String";
    };
    if (v instanceof PureScript_Backend_Optimizer_CoreFn.Char) {
        return "Char";
    };
    if (v instanceof PureScript_Backend_Optimizer_CoreFn["Boolean"]) {
        return "Boolean";
    };
    if (v instanceof PureScript_Backend_Optimizer_CoreFn.Unit) {
        return "Unit";
    };
    if (v instanceof PureScript_Backend_Optimizer_CoreFn.Any) {
        return "Any";
    };
    if (v instanceof PureScript_Backend_Optimizer_CoreFn.TypeLevelString) {
        return "TypeLevelString_" + v.value0;
    };
    if (v instanceof PureScript_Backend_Optimizer_CoreFn["Array"]) {
        return "Array_" + mangleType(v.value0);
    };
    if (v instanceof PureScript_Backend_Optimizer_CoreFn.Func) {
        return "Func_" + (Data_String_Common.joinWith("_")(Data_Functor.map(Data_Functor.functorArray)(mangleType)(v.value0)) + ("_" + mangleType(v.value1)));
    };
    if (v instanceof PureScript_Backend_Optimizer_CoreFn.Record) {
        return "Record_" + mangleType(v.value0);
    };
    if (v instanceof PureScript_Backend_Optimizer_CoreFn.Row) {
        return "Row_" + (Data_String_Common.joinWith("_")(Data_Functor.map(Data_Functor.functorArray)(function (v1) {
            return v1.value0 + ("_" + mangleType(v1.value1));
        })(v.value0)) + ("_" + Data_Maybe.maybe("Empty")(mangleType)(v.value1)));
    };
    if (v instanceof PureScript_Backend_Optimizer_CoreFn.TypeApp) {
        return "TypeApp_" + (mangleType(v.value0) + ("_" + Data_String_Common.joinWith("_")(Data_Functor.map(Data_Functor.functorArray)(mangleType)(v.value1))));
    };
    if (v instanceof PureScript_Backend_Optimizer_CoreFn.ForAll) {
        return "ForAll_" + (Data_String_Common.joinWith("_")(v.value0) + ("_" + mangleType(v.value1)));
    };
    if (v instanceof PureScript_Backend_Optimizer_CoreFn.ConstrainedType) {
        return "ConstrainedType_" + (Data_String_Common.joinWith("_")(Data_Functor.map(Data_Functor.functorArray)(function (v1) {
            return Data_String_Common.joinWith("_")(v1.value0) + ("_" + Data_String_Common.joinWith("_")(Data_Functor.map(Data_Functor.functorArray)(mangleType)(v1.value1)));
        })(v.value0)) + ("_" + mangleType(v.value1)));
    };
    if (v instanceof PureScript_Backend_Optimizer_CoreFn.ADT) {
        return "ADT_" + (Data_String_Common.joinWith("_")(v.value1) + (function () {
            var $145 = Data_Array.length(v.value2) === 0;
            if ($145) {
                return "";
            };
            return "_" + Data_String_Common.joinWith("_")(Data_Functor.map(Data_Functor.functorArray)(mangleType)(v.value2));
        })());
    };
    if (v instanceof PureScript_Backend_Optimizer_CoreFn.TypeVar) {
        return "Var_" + v.value0;
    };
    if (v instanceof PureScript_Backend_Optimizer_CoreFn.Any) {
        return "Any";
    };
    throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Monomorphize (line 52, column 1 - line 52, column 33): " + [ v.constructor.name ]);
};
var isStatic = function (v) {
    if (v instanceof PureScript_Backend_Optimizer_CoreFn.ExprVar && v.value1.value0 instanceof Data_Maybe.Just) {
        return true;
    };
    if (v instanceof PureScript_Backend_Optimizer_CoreFn.ExprVar && v.value1.value0 instanceof Data_Maybe.Nothing) {
        return false;
    };
    if (v instanceof PureScript_Backend_Optimizer_CoreFn.ExprApp) {
        return isStatic(v.value1) && isStatic(v.value2);
    };
    if (v instanceof PureScript_Backend_Optimizer_CoreFn.ExprAccessor) {
        return isStatic(v.value1);
    };
    return false;
};
var getExprAnn = function (v) {
    if (v instanceof PureScript_Backend_Optimizer_CoreFn.ExprVar) {
        return v.value0;
    };
    if (v instanceof PureScript_Backend_Optimizer_CoreFn.ExprLit) {
        return v.value0;
    };
    if (v instanceof PureScript_Backend_Optimizer_CoreFn.ExprApp) {
        return v.value0;
    };
    if (v instanceof PureScript_Backend_Optimizer_CoreFn.ExprAbs) {
        return v.value0;
    };
    if (v instanceof PureScript_Backend_Optimizer_CoreFn.ExprLet) {
        return v.value0;
    };
    if (v instanceof PureScript_Backend_Optimizer_CoreFn.ExprCase) {
        return v.value0;
    };
    if (v instanceof PureScript_Backend_Optimizer_CoreFn.ExprConstructor) {
        return v.value0;
    };
    if (v instanceof PureScript_Backend_Optimizer_CoreFn.ExprAccessor) {
        return v.value0;
    };
    if (v instanceof PureScript_Backend_Optimizer_CoreFn.ExprUpdate) {
        return v.value0;
    };
    throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Monomorphize (line 85, column 14 - line 94, column 28): " + [ v.constructor.name ]);
};
var inferExprType = function (expr) {
    if (expr instanceof PureScript_Backend_Optimizer_CoreFn.ExprApp) {
        var fTy = inferExprType(expr.value1);
        if (fTy instanceof Data_Maybe.Just && fTy.value0 instanceof PureScript_Backend_Optimizer_CoreFn.Func) {
            var $195 = Data_Array.length(fTy.value0.value0) > 1;
            if ($195) {
                return new Data_Maybe.Just(new PureScript_Backend_Optimizer_CoreFn.Func(Data_Maybe.fromMaybe([  ])(Data_Array.tail(fTy.value0.value0)), fTy.value0.value1));
            };
            return new Data_Maybe.Just(fTy.value0.value1);
        };
        return Data_Maybe.Nothing.value;
    };
    var v = getExprAnn(expr);
    return v.type;
};
var defaultToAny = function (v) {
    if (v instanceof PureScript_Backend_Optimizer_CoreFn.TypeVar) {
        return PureScript_Backend_Optimizer_CoreFn.Any.value;
    };
    if (v instanceof PureScript_Backend_Optimizer_CoreFn["Array"]) {
        return new PureScript_Backend_Optimizer_CoreFn["Array"](defaultToAny(v.value0));
    };
    if (v instanceof PureScript_Backend_Optimizer_CoreFn.Func) {
        return new PureScript_Backend_Optimizer_CoreFn.Func(Data_Functor.map(Data_Functor.functorArray)(defaultToAny)(v.value0), defaultToAny(v.value1));
    };
    if (v instanceof PureScript_Backend_Optimizer_CoreFn.Record) {
        return new PureScript_Backend_Optimizer_CoreFn.Record(defaultToAny(v.value0));
    };
    if (v instanceof PureScript_Backend_Optimizer_CoreFn.Row) {
        return new PureScript_Backend_Optimizer_CoreFn.Row(Data_Functor.map(Data_Functor.functorArray)(function (v1) {
            return new Data_Tuple.Tuple(v1.value0, defaultToAny(v1.value1));
        })(v.value0), Data_Functor.map(Data_Maybe.functorMaybe)(defaultToAny)(v.value1));
    };
    if (v instanceof PureScript_Backend_Optimizer_CoreFn.TypeApp) {
        return new PureScript_Backend_Optimizer_CoreFn.TypeApp(defaultToAny(v.value0), Data_Functor.map(Data_Functor.functorArray)(defaultToAny)(v.value1));
    };
    if (v instanceof PureScript_Backend_Optimizer_CoreFn.ForAll) {
        return new PureScript_Backend_Optimizer_CoreFn.ForAll(v.value0, defaultToAny(v.value1));
    };
    if (v instanceof PureScript_Backend_Optimizer_CoreFn.ConstrainedType) {
        return new PureScript_Backend_Optimizer_CoreFn.ConstrainedType(Data_Functor.map(Data_Functor.functorArray)(function (v1) {
            return new Data_Tuple.Tuple(v1.value0, Data_Functor.map(Data_Functor.functorArray)(defaultToAny)(v1.value1));
        })(v.value0), defaultToAny(v.value1));
    };
    if (v instanceof PureScript_Backend_Optimizer_CoreFn.ADT) {
        return new PureScript_Backend_Optimizer_CoreFn.ADT(v.value0, v.value1, Data_Functor.map(Data_Functor.functorArray)(defaultToAny)(v.value2));
    };
    return v;
};
var collectTypesFromExpr = function (expr) {
    return function (acc) {
        if (expr instanceof PureScript_Backend_Optimizer_CoreFn.ExprVar) {
            return Data_Maybe.maybe(acc)(function (t) {
                return Data_Set.insert(PureScript_Backend_Optimizer_CoreFn.ordExprType)(t)(acc);
            })(expr.value0.type);
        };
        if (expr instanceof PureScript_Backend_Optimizer_CoreFn.ExprLit) {
            return Data_Foldable.foldl(PureScript_Backend_Optimizer_CoreFn.foldableLiteral)(function (a) {
                return function (e) {
                    return collectTypesFromExpr(e)(a);
                };
            })(Data_Maybe.maybe(acc)(function (t) {
                return Data_Set.insert(PureScript_Backend_Optimizer_CoreFn.ordExprType)(t)(acc);
            })(expr.value0.type))(expr.value1);
        };
        if (expr instanceof PureScript_Backend_Optimizer_CoreFn.ExprApp) {
            return collectTypesFromExpr(expr.value2)(collectTypesFromExpr(expr.value1)(Data_Maybe.maybe(acc)(function (t) {
                return Data_Set.insert(PureScript_Backend_Optimizer_CoreFn.ordExprType)(t)(acc);
            })(expr.value0.type)));
        };
        if (expr instanceof PureScript_Backend_Optimizer_CoreFn.ExprAbs) {
            return collectTypesFromExpr(expr.value2)(Data_Maybe.maybe(acc)(function (t) {
                return Data_Set.insert(PureScript_Backend_Optimizer_CoreFn.ordExprType)(t)(acc);
            })(expr.value0.type));
        };
        if (expr instanceof PureScript_Backend_Optimizer_CoreFn.ExprLet) {
            return Data_Foldable.foldl(Data_Foldable.foldableArray)(function (a) {
                return function (b) {
                    return collectTypesFromBind(b)(a);
                };
            })(collectTypesFromExpr(expr.value2)(Data_Maybe.maybe(acc)(function (t) {
                return Data_Set.insert(PureScript_Backend_Optimizer_CoreFn.ordExprType)(t)(acc);
            })(expr.value0.type)))(expr.value1);
        };
        if (expr instanceof PureScript_Backend_Optimizer_CoreFn.ExprCase) {
            return Data_Foldable.foldl(Data_Foldable.foldableArray)(function (a) {
                return function (alt) {
                    return collectTypesFromAlt(alt)(a);
                };
            })(Data_Foldable.foldl(Data_Foldable.foldableArray)(Data_Function.flip(collectTypesFromExpr))(Data_Maybe.maybe(acc)(function (t) {
                return Data_Set.insert(PureScript_Backend_Optimizer_CoreFn.ordExprType)(t)(acc);
            })(expr.value0.type))(expr.value1))(expr.value2);
        };
        if (expr instanceof PureScript_Backend_Optimizer_CoreFn.ExprConstructor) {
            return Data_Maybe.maybe(acc)(function (t) {
                return Data_Set.insert(PureScript_Backend_Optimizer_CoreFn.ordExprType)(t)(acc);
            })(expr.value0.type);
        };
        if (expr instanceof PureScript_Backend_Optimizer_CoreFn.ExprAccessor) {
            return collectTypesFromExpr(expr.value1)(Data_Maybe.maybe(acc)(function (t) {
                return Data_Set.insert(PureScript_Backend_Optimizer_CoreFn.ordExprType)(t)(acc);
            })(expr.value0.type));
        };
        if (expr instanceof PureScript_Backend_Optimizer_CoreFn.ExprUpdate) {
            return Data_Foldable.foldl(Data_Foldable.foldableArray)(function (a) {
                return function (v) {
                    return collectTypesFromExpr(v.value1)(a);
                };
            })(collectTypesFromExpr(expr.value1)(Data_Maybe.maybe(acc)(function (t) {
                return Data_Set.insert(PureScript_Backend_Optimizer_CoreFn.ordExprType)(t)(acc);
            })(expr.value0.type)))(expr.value2);
        };
        throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Monomorphize (line 193, column 33 - line 202, column 161): " + [ expr.constructor.name ]);
    };
};
var collectTypesFromBind = function (v) {
    return function (v1) {
        if (v instanceof PureScript_Backend_Optimizer_CoreFn.NonRec) {
            return collectTypesFromExpr(v.value0.value2)(v1);
        };
        if (v instanceof PureScript_Backend_Optimizer_CoreFn.Rec) {
            return Data_Foldable.foldl(Data_Foldable.foldableArray)(function (a) {
                return function (v2) {
                    return collectTypesFromExpr(v2.value2)(a);
                };
            })(v1)(v.value0);
        };
        throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Monomorphize (line 204, column 1 - line 204, column 65): " + [ v.constructor.name, v1.constructor.name ]);
    };
};
var collectTypesFromAlt = function (v) {
    return function (v1) {
        if (v.value1 instanceof PureScript_Backend_Optimizer_CoreFn.Unconditional) {
            return collectTypesFromExpr(v.value1.value0)(v1);
        };
        if (v.value1 instanceof PureScript_Backend_Optimizer_CoreFn.Guarded) {
            return Data_Foldable.foldl(Data_Foldable.foldableArray)(function (a) {
                return function (v2) {
                    return collectTypesFromExpr(v2.value1)(collectTypesFromExpr(v2.value0)(a));
                };
            })(v1)(v.value1.value0);
        };
        throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Monomorphize (line 208, column 1 - line 208, column 75): " + [ v.constructor.name, v1.constructor.name ]);
    };
};
var collectAppSpine = /* #__PURE__ */ (function () {
    var go = function ($copy_v) {
        return function ($copy_v1) {
            var $tco_var_v = $copy_v;
            var $tco_done = false;
            var $tco_result;
            function $tco_loop(v, v1) {
                if (v1 instanceof PureScript_Backend_Optimizer_CoreFn.ExprApp) {
                    $tco_var_v = Data_Array.cons(v1.value2)(v);
                    $copy_v1 = v1.value1;
                    return;
                };
                $tco_done = true;
                return {
                    f: v1,
                    args: v
                };
            };
            while (!$tco_done) {
                $tco_result = $tco_loop($tco_var_v, $copy_v1);
            };
            return $tco_result;
        };
    };
    return go([  ]);
})();
var monomorphizeProp = function (modName) {
    return function (instMap) {
        return function (v) {
            return new PureScript_Backend_Optimizer_CoreFn.Prop(v.value0, monomorphizeExpr(modName)(instMap)(v.value1));
        };
    };
};
var monomorphizeGuard = function (modName) {
    return function (instMap) {
        return function (v) {
            return new PureScript_Backend_Optimizer_CoreFn.Guard(monomorphizeExpr(modName)(instMap)(v.value0), monomorphizeExpr(modName)(instMap)(v.value1));
        };
    };
};
var monomorphizeExpr = function (modName) {
    return function (instMap) {
        return function (expr) {
            if (expr instanceof PureScript_Backend_Optimizer_CoreFn.ExprApp) {
                var v = collectAppSpine(expr);
                var f$prime = monomorphizeExpr(modName)(instMap)(v.f);
                var args$prime = Data_Functor.map(Data_Functor.functorArray)(monomorphizeExpr(modName)(instMap))(v.args);
                if (f$prime instanceof PureScript_Backend_Optimizer_CoreFn.ExprVar) {
                    if (f$prime.value0.type instanceof Data_Maybe.Just && f$prime.value0.type.value0 instanceof PureScript_Backend_Optimizer_CoreFn.Func) {
                        var qualName = (function () {
                            if (f$prime.value1.value0 instanceof Data_Maybe.Just) {
                                return Data_Newtype.unwrap()(f$prime.value1.value0.value0) + ("." + f$prime.value1.value1);
                            };
                            if (f$prime.value1.value0 instanceof Data_Maybe.Nothing) {
                                return modName + ("." + f$prime.value1.value1);
                            };
                            throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Monomorphize (line 316, column 26 - line 318, column 50): " + [ f$prime.value1.value0.constructor.name ]);
                        })();
                        var argTypes = Data_Array.mapMaybe(inferExprType)(args$prime);
                        var substArgs = Data_Array.foldl(function (substAcc) {
                            return function (v1) {
                                return PureScript_Backend_Optimizer_Substitute.unify(v1.value0)(v1.value1)(substAcc);
                            };
                        })(Data_Map_Internal.empty)(Data_Array.zip(f$prime.value0.type.value0.value0)(argTypes));
                        var subst = (function () {
                            if (expr.value0.type instanceof Data_Maybe.Just) {
                                var remainingType = (function () {
                                    var $303 = Data_Array.length(args$prime) < Data_Array.length(f$prime.value0.type.value0.value0);
                                    if ($303) {
                                        return new PureScript_Backend_Optimizer_CoreFn.Func(Data_Array.drop(Data_Array.length(args$prime))(f$prime.value0.type.value0.value0), f$prime.value0.type.value0.value1);
                                    };
                                    return f$prime.value0.type.value0.value1;
                                })();
                                return PureScript_Backend_Optimizer_Substitute.unify(remainingType)(expr.value0.type.value0)(substArgs);
                            };
                            if (expr.value0.type instanceof Data_Maybe.Nothing) {
                                return substArgs;
                            };
                            throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Monomorphize (line 307, column 23 - line 313, column 37): " + [ expr.value0.type.constructor.name ]);
                        })();
                        var instType = PureScript_Backend_Optimizer_Substitute.substituteExprType(subst)(f$prime.value0.type.value0);
                        var v1 = partitionArgs(f$prime.value0.type.value0)(args$prime);
                        var filteredArgs = Data_Semigroup.append(Data_Semigroup.semigroupArray)(Data_Array.filter(function (d) {
                            return !isStatic(d);
                        })(v1.dictArgs))(v1.normalArgs);
                        var $306 = Data_Map_Internal.isEmpty(subst);
                        if ($306) {
                            return rebuildApp(expr.value0)(f$prime)(args$prime);
                        };
                        var v2 = Data_Map_Internal.lookup(Data_Ord.ordString)(qualName)(instMap);
                        if (v2 instanceof Data_Maybe.Just) {
                            var specializedName = f$prime.value1.value1 + ("_" + mangleType(defaultToAny(instType)));
                            var specializedVar = new PureScript_Backend_Optimizer_CoreFn.ExprVar(f$prime.value0, new PureScript_Backend_Optimizer_CoreFn.Qualified(f$prime.value1.value0, specializedName));
                            return rebuildApp(expr.value0)(specializedVar)(filteredArgs);
                        };
                        if (v2 instanceof Data_Maybe.Nothing) {
                            return rebuildApp(expr.value0)(f$prime)(args$prime);
                        };
                        throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Monomorphize (line 325, column 20 - line 331, column 48): " + [ v2.constructor.name ]);
                    };
                    return rebuildApp(expr.value0)(f$prime)(args$prime);
                };
                return rebuildApp(expr.value0)(f$prime)(args$prime);
            };
            if (expr instanceof PureScript_Backend_Optimizer_CoreFn.ExprVar) {
                return new PureScript_Backend_Optimizer_CoreFn.ExprVar(expr.value0, expr.value1);
            };
            if (expr instanceof PureScript_Backend_Optimizer_CoreFn.ExprLit) {
                return new PureScript_Backend_Optimizer_CoreFn.ExprLit(expr.value0, Data_Functor.map(PureScript_Backend_Optimizer_CoreFn.functorLiteral)(monomorphizeExpr(modName)(instMap))(expr.value1));
            };
            if (expr instanceof PureScript_Backend_Optimizer_CoreFn.ExprAbs) {
                return new PureScript_Backend_Optimizer_CoreFn.ExprAbs(expr.value0, expr.value1, monomorphizeExpr(modName)(instMap)(expr.value2));
            };
            if (expr instanceof PureScript_Backend_Optimizer_CoreFn.ExprLet) {
                return new PureScript_Backend_Optimizer_CoreFn.ExprLet(expr.value0, Data_Functor.map(Data_Functor.functorArray)(monomorphizeBindLocal(modName)(instMap))(expr.value1), monomorphizeExpr(modName)(instMap)(expr.value2));
            };
            if (expr instanceof PureScript_Backend_Optimizer_CoreFn.ExprCase) {
                return new PureScript_Backend_Optimizer_CoreFn.ExprCase(expr.value0, Data_Functor.map(Data_Functor.functorArray)(monomorphizeExpr(modName)(instMap))(expr.value1), Data_Functor.map(Data_Functor.functorArray)(monomorphizeAlt(modName)(instMap))(expr.value2));
            };
            if (expr instanceof PureScript_Backend_Optimizer_CoreFn.ExprConstructor) {
                return new PureScript_Backend_Optimizer_CoreFn.ExprConstructor(expr.value0, expr.value1, expr.value2, expr.value3);
            };
            if (expr instanceof PureScript_Backend_Optimizer_CoreFn.ExprAccessor) {
                return new PureScript_Backend_Optimizer_CoreFn.ExprAccessor(expr.value0, monomorphizeExpr(modName)(instMap)(expr.value1), expr.value2);
            };
            if (expr instanceof PureScript_Backend_Optimizer_CoreFn.ExprUpdate) {
                return new PureScript_Backend_Optimizer_CoreFn.ExprUpdate(expr.value0, monomorphizeExpr(modName)(instMap)(expr.value1), Data_Functor.map(Data_Functor.functorArray)(monomorphizeProp(modName)(instMap))(expr.value2));
            };
            throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Monomorphize (line 292, column 41 - line 342, column 127): " + [ expr.constructor.name ]);
        };
    };
};
var monomorphizeCaseGuard = function (v) {
    return function (v1) {
        return function (v2) {
            if (v2 instanceof PureScript_Backend_Optimizer_CoreFn.Unconditional) {
                return new PureScript_Backend_Optimizer_CoreFn.Unconditional(monomorphizeExpr(v)(v1)(v2.value0));
            };
            if (v2 instanceof PureScript_Backend_Optimizer_CoreFn.Guarded) {
                return new PureScript_Backend_Optimizer_CoreFn.Guarded(Data_Functor.map(Data_Functor.functorArray)(monomorphizeGuard(v)(v1))(v2.value0));
            };
            throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Monomorphize (line 354, column 1 - line 354, column 86): " + [ v.constructor.name, v1.constructor.name, v2.constructor.name ]);
        };
    };
};
var monomorphizeBindingLocal = function (modName) {
    return function (instMap) {
        return function (v) {
            return new PureScript_Backend_Optimizer_CoreFn.Binding(v.value0, v.value1, monomorphizeExpr(modName)(instMap)(v.value2));
        };
    };
};
var monomorphizeBindLocal = function (v) {
    return function (v1) {
        return function (v2) {
            if (v2 instanceof PureScript_Backend_Optimizer_CoreFn.NonRec) {
                return new PureScript_Backend_Optimizer_CoreFn.NonRec(monomorphizeBindingLocal(v)(v1)(v2.value0));
            };
            if (v2 instanceof PureScript_Backend_Optimizer_CoreFn.Rec) {
                return new PureScript_Backend_Optimizer_CoreFn.Rec(Data_Functor.map(Data_Functor.functorArray)(monomorphizeBindingLocal(v)(v1))(v2.value0));
            };
            throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Monomorphize (line 344, column 1 - line 344, column 76): " + [ v.constructor.name, v1.constructor.name, v2.constructor.name ]);
        };
    };
};
var monomorphizeAlt = function (modName) {
    return function (instMap) {
        return function (v) {
            return new PureScript_Backend_Optimizer_CoreFn.CaseAlternative(v.value0, monomorphizeCaseGuard(modName)(instMap)(v.value1));
        };
    };
};
var collectProp = function (modName) {
    return function (acc) {
        return function (v) {
            return collectExpr(modName)(acc)(v.value1);
        };
    };
};
var collectGuard = function (modName) {
    return function (acc) {
        return function (v) {
            return collectExpr(modName)(collectExpr(modName)(acc)(v.value0))(v.value1);
        };
    };
};
var collectExpr = function (modName) {
    return function (acc) {
        return function (expr) {
            if (expr instanceof PureScript_Backend_Optimizer_CoreFn.ExprVar) {
                if (expr.value0.type instanceof Data_Maybe.Just) {
                    var qualName = (function () {
                        if (expr.value1.value0 instanceof Data_Maybe.Just) {
                            return Data_Newtype.unwrap()(expr.value1.value0.value0) + ("." + expr.value1.value1);
                        };
                        if (expr.value1.value0 instanceof Data_Maybe.Nothing) {
                            return modName + ("." + expr.value1.value1);
                        };
                        throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Monomorphize (line 131, column 24 - line 133, column 48): " + [ expr.value1.value0.constructor.name ]);
                    })();
                    return Data_Map_Internal.insertWith(Data_Ord.ordString)(union)(qualName)(Data_Map_Internal.singleton(defaultToAny(expr.value0.type.value0))({
                        dictArgs: [  ]
                    }))(acc);
                };
                if (expr.value0.type instanceof Data_Maybe.Nothing) {
                    return acc;
                };
                throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Monomorphize (line 129, column 5 - line 135, column 21): " + [ expr.value0.type.constructor.name ]);
            };
            if (expr instanceof PureScript_Backend_Optimizer_CoreFn.ExprApp) {
                var v = collectAppSpine(expr);
                var acc1 = collectExpr(modName)(acc)(v.f);
                var acc2 = Data_Foldable.foldl(Data_Foldable.foldableArray)(collectExpr(modName))(acc1)(v.args);
                if (v.f instanceof PureScript_Backend_Optimizer_CoreFn.ExprVar) {
                    if (v.f.value0.type instanceof Data_Maybe.Just && v.f.value0.type.value0 instanceof PureScript_Backend_Optimizer_CoreFn.Func) {
                        var qualName = (function () {
                            if (v.f.value1.value0 instanceof Data_Maybe.Just) {
                                return Data_Newtype.unwrap()(v.f.value1.value0.value0) + ("." + v.f.value1.value1);
                            };
                            if (v.f.value1.value0 instanceof Data_Maybe.Nothing) {
                                return modName + ("." + v.f.value1.value1);
                            };
                            throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Monomorphize (line 159, column 26 - line 161, column 50): " + [ v.f.value1.value0.constructor.name ]);
                        })();
                        var argTypes = Data_Array.mapMaybe(inferExprType)(v.args);
                        var substArgs = Data_Array.foldl(function (substAcc) {
                            return function (v1) {
                                return PureScript_Backend_Optimizer_Substitute.unify(v1.value0)(v1.value1)(substAcc);
                            };
                        })(Data_Map_Internal.empty)(Data_Array.zip(v.f.value0.type.value0.value0)(argTypes));
                        var appType = (function () {
                            var v1 = getExprAnn(expr);
                            return v1.type;
                        })();
                        var subst = (function () {
                            if (appType instanceof Data_Maybe.Just) {
                                var remainingType = (function () {
                                    var $396 = Data_Array.length(v.args) < Data_Array.length(v.f.value0.type.value0.value0);
                                    if ($396) {
                                        return new PureScript_Backend_Optimizer_CoreFn.Func(Data_Array.drop(Data_Array.length(v.args))(v.f.value0.type.value0.value0), v.f.value0.type.value0.value1);
                                    };
                                    return v.f.value0.type.value0.value1;
                                })();
                                return PureScript_Backend_Optimizer_Substitute.unify(remainingType)(appType.value0)(substArgs);
                            };
                            if (appType instanceof Data_Maybe.Nothing) {
                                return substArgs;
                            };
                            throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Monomorphize (line 150, column 23 - line 156, column 37): " + [ appType.constructor.name ]);
                        })();
                        var instType = PureScript_Backend_Optimizer_Substitute.substituteExprType(subst)(v.f.value0.type.value0);
                        var v1 = partitionArgs(v.f.value0.type.value0)(v.args);
                        var $399 = Data_Map_Internal.isEmpty(subst);
                        if ($399) {
                            return acc2;
                        };
                        return Data_Map_Internal.insertWith(Data_Ord.ordString)(union)(qualName)(Data_Map_Internal.singleton(defaultToAny(instType))({
                            dictArgs: v1.dictArgs
                        }))(acc2);
                    };
                    return acc2;
                };
                return acc2;
            };
            if (expr instanceof PureScript_Backend_Optimizer_CoreFn.ExprLit) {
                return Data_Foldable.foldl(PureScript_Backend_Optimizer_CoreFn.foldableLiteral)(collectExpr(modName))(acc)(expr.value1);
            };
            if (expr instanceof PureScript_Backend_Optimizer_CoreFn.ExprConstructor) {
                return acc;
            };
            if (expr instanceof PureScript_Backend_Optimizer_CoreFn.ExprAccessor) {
                return collectExpr(modName)(acc)(expr.value1);
            };
            if (expr instanceof PureScript_Backend_Optimizer_CoreFn.ExprUpdate) {
                return Data_Foldable.foldl(Data_Foldable.foldableArray)(collectProp(modName))(collectExpr(modName)(acc)(expr.value1))(expr.value2);
            };
            if (expr instanceof PureScript_Backend_Optimizer_CoreFn.ExprAbs) {
                return collectExpr(modName)(acc)(expr.value2);
            };
            if (expr instanceof PureScript_Backend_Optimizer_CoreFn.ExprCase) {
                return Data_Foldable.foldl(Data_Foldable.foldableArray)(collectAlt(modName))(Data_Foldable.foldl(Data_Foldable.foldableArray)(collectExpr(modName))(acc)(expr.value1))(expr.value2);
            };
            if (expr instanceof PureScript_Backend_Optimizer_CoreFn.ExprLet) {
                return Data_Foldable.foldl(Data_Foldable.foldableArray)(collectBind(modName))(collectExpr(modName)(acc)(expr.value2))(expr.value1);
            };
            throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Monomorphize (line 127, column 32 - line 176, column 85): " + [ expr.constructor.name ]);
        };
    };
};
var collectBinding = function (modName) {
    return function (acc) {
        return function (v) {
            return collectExpr(modName)(acc)(v.value2);
        };
    };
};
var collectBind = function (v) {
    return function (v1) {
        return function (v2) {
            if (v2 instanceof PureScript_Backend_Optimizer_CoreFn.NonRec) {
                return collectBinding(v)(v1)(v2.value0);
            };
            if (v2 instanceof PureScript_Backend_Optimizer_CoreFn.Rec) {
                return Data_Foldable.foldl(Data_Foldable.foldableArray)(collectBinding(v))(v1)(v2.value0);
            };
            throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Monomorphize (line 77, column 1 - line 77, column 74): " + [ v.constructor.name, v1.constructor.name, v2.constructor.name ]);
        };
    };
};
var collectAlt = function (modName) {
    return function (acc) {
        return function (v) {
            if (v.value1 instanceof PureScript_Backend_Optimizer_CoreFn.Unconditional) {
                return collectExpr(modName)(acc)(v.value1.value0);
            };
            if (v.value1 instanceof PureScript_Backend_Optimizer_CoreFn.Guarded) {
                return Data_Foldable.foldl(Data_Foldable.foldableArray)(collectGuard(modName))(acc)(v.value1.value0);
            };
            throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Monomorphize (line 182, column 49 - line 184, column 60): " + [ v.value1.constructor.name ]);
        };
    };
};
var collectInstantiations = function (acc) {
    return function (v) {
        var modNameStr = Data_Newtype.unwrap()(v.name);
        return Data_Foldable.foldl(Data_Foldable.foldableArray)(collectBind(modNameStr))(acc)(v.decls);
    };
};
var collectAllTypes = function (v) {
    return Data_Foldable.foldl(Data_Foldable.foldableArray)(function (a) {
        return function (b) {
            return collectTypesFromBind(b)(a);
        };
    })(Data_Set.empty)(v.decls);
};
var applyDicts = function (args) {
    return function (body) {
        var go = function (dicts) {
            return function (e) {
                var v = Data_Array.uncons(dicts);
                if (v instanceof Data_Maybe.Nothing) {
                    return e;
                };
                if (v instanceof Data_Maybe.Just) {
                    var $457 = isStatic(v.value0.head);
                    if ($457) {
                        if (e instanceof PureScript_Backend_Optimizer_CoreFn.ExprAbs) {
                            var body$prime = go(v.value0.tail)(e.value2);
                            return new PureScript_Backend_Optimizer_CoreFn.ExprLet(getExprAnn(body$prime), [ new PureScript_Backend_Optimizer_CoreFn.NonRec(new PureScript_Backend_Optimizer_CoreFn.Binding(getExprAnn(v.value0.head), e.value1, v.value0.head)) ], body$prime);
                        };
                        return e;
                    };
                    if (e instanceof PureScript_Backend_Optimizer_CoreFn.ExprAbs) {
                        return new PureScript_Backend_Optimizer_CoreFn.ExprAbs(e.value0, e.value1, go(v.value0.tail)(e.value2));
                    };
                    return e;
                };
                throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Monomorphize (line 246, column 16 - line 258, column 17): " + [ v.constructor.name ]);
            };
        };
        return go(args)(body);
    };
};
var monomorphizeBinding = function (modName) {
    return function (instMap) {
        return function (v) {
            var qualName = modName + ("." + v.value1);
            var v1 = Data_Map_Internal.lookup(Data_Ord.ordString)(qualName)(instMap);
            if (v1 instanceof Data_Maybe.Just) {
                return Data_Semigroup.append(Data_Semigroup.semigroupArray)(Data_Functor.map(Data_Functor.functorArray)(function (v2) {
                    var newName = v.value1 + ("_" + mangleType(v2.value0));
                    var genericType = (function () {
                        var v3 = getExprAnn(v.value2);
                        return Data_Maybe.fromMaybe(PureScript_Backend_Optimizer_CoreFn.Any.value)(v3.type);
                    })();
                    var subst = PureScript_Backend_Optimizer_Substitute.unify(genericType)(v2.value0)(Data_Map_Internal.empty);
                    var substFn = function (t) {
                        return PureScript_Backend_Optimizer_Substitute.substituteExprType(subst)(t);
                    };
                    var exprWithDicts = applyDicts(v2.value1.dictArgs)(v.value2);
                    var specializedExpr = rewriteExpr(substFn)(monomorphizeExpr(modName)(instMap)(exprWithDicts));
                    return new PureScript_Backend_Optimizer_CoreFn.Binding(mapAnn(substFn)(v.value0), newName, specializedExpr);
                })(Data_Map_Internal.toUnfoldable(Data_Unfoldable.unfoldableArray)(v1.value0)))([ new PureScript_Backend_Optimizer_CoreFn.Binding(v.value0, v.value1, monomorphizeExpr(modName)(instMap)(v.value2)) ]);
            };
            if (v1 instanceof Data_Maybe.Nothing) {
                return [ new PureScript_Backend_Optimizer_CoreFn.Binding(v.value0, v.value1, monomorphizeExpr(modName)(instMap)(v.value2)) ];
            };
            throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Monomorphize (line 274, column 6 - line 289, column 75): " + [ v1.constructor.name ]);
        };
    };
};
var monomorphizeBind = function (v) {
    return function (v1) {
        return function (v2) {
            if (v2 instanceof PureScript_Backend_Optimizer_CoreFn.NonRec) {
                return Data_Functor.map(Data_Functor.functorArray)(PureScript_Backend_Optimizer_CoreFn.NonRec.create)(monomorphizeBinding(v)(v1)(v2.value0));
            };
            if (v2 instanceof PureScript_Backend_Optimizer_CoreFn.Rec) {
                return [ new PureScript_Backend_Optimizer_CoreFn.Rec(Data_Array.concatMap(monomorphizeBinding(v)(v1))(v2.value0)) ];
            };
            throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Monomorphize (line 265, column 1 - line 265, column 79): " + [ v.constructor.name, v1.constructor.name, v2.constructor.name ]);
        };
    };
};
var monomorphize = function (instMap) {
    return function (v) {
        var modNameStr = Data_Newtype.unwrap()(v.name);
        return {
            name: v.name,
            path: v.path,
            span: v.span,
            imports: v.imports,
            exports: v.exports,
            reExports: v.reExports,
            dataDecls: v.dataDecls,
            classDecls: v.classDecls,
            foreign: v.foreign,
            comments: v.comments,
            decls: Data_Array.concatMap(monomorphizeBind(modNameStr)(instMap))(v.decls)
        };
    };
};
export {
    collectInstantiations,
    collectAllTypes,
    mangleType,
    defaultToAny,
    collectAppSpine,
    getExprAnn,
    inferExprType,
    monomorphize
};
