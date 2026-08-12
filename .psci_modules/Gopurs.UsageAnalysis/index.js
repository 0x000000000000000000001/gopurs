import * as Data_Array from "../Data.Array/index.js";
import * as Data_Array_NonEmpty from "../Data.Array.NonEmpty/index.js";
import * as Data_Function from "../Data.Function/index.js";
import * as Data_Functor from "../Data.Functor/index.js";
import * as Data_Map_Internal from "../Data.Map.Internal/index.js";
import * as Data_Maybe from "../Data.Maybe/index.js";
import * as Data_Ord from "../Data.Ord/index.js";
import * as Data_Semiring from "../Data.Semiring/index.js";
import * as Data_Set from "../Data.Set/index.js";
import * as Data_Show from "../Data.Show/index.js";
import * as Data_String_Common from "../Data.String.Common/index.js";
import * as PureScript_Backend_Optimizer_CoreFn from "../PureScript.Backend.Optimizer.CoreFn/index.js";
import * as PureScript_Backend_Optimizer_Syntax from "../PureScript.Backend.Optimizer.Syntax/index.js";
var $$delete = /* #__PURE__ */ Data_Set["delete"](Data_Ord.ordString);
var delete1 = /* #__PURE__ */ Data_Map_Internal["delete"](Data_Ord.ordString);
var sanitizeName = function (name) {
    var s1 = Data_String_Common.replaceAll("\"")("_quote_")(Data_String_Common.replaceAll(".")("_dot_")(Data_String_Common.replaceAll("'")("_prime")(Data_String_Common.replaceAll("$")("_dollar")(name))));
    var $27 = s1 === "break" || (s1 === "default" || (s1 === "func" || (s1 === "interface" || (s1 === "select" || (s1 === "case" || (s1 === "defer" || (s1 === "go" || (s1 === "map" || (s1 === "struct" || (s1 === "chan" || (s1 === "else" || (s1 === "goto" || (s1 === "package" || (s1 === "switch" || (s1 === "const" || (s1 === "fallthrough" || (s1 === "if" || (s1 === "range" || (s1 === "type" || (s1 === "continue" || (s1 === "for" || (s1 === "import" || (s1 === "return" || s1 === "var")))))))))))))))))))))));
    if ($27) {
        return "go__" + s1;
    };
    return s1;
};
var maxUsages = /* #__PURE__ */ Data_Map_Internal.unionWith(Data_Ord.ordString)(/* #__PURE__ */ Data_Ord.max(Data_Ord.ordInt));
var localId = function (v) {
    return function (v1) {
        if (v instanceof Data_Maybe.Just) {
            return sanitizeName(v.value0) + ("_" + Data_Show.show(Data_Show.showInt)(v1));
        };
        if (v instanceof Data_Maybe.Nothing) {
            return "__local_var_" + Data_Show.show(Data_Show.showInt)(v1);
        };
        throw new Error("Failed pattern match at Gopurs.UsageAnalysis (line 24, column 1 - line 24, column 42): " + [ v.constructor.name, v1.constructor.name ]);
    };
};
var freeVars = function (v) {
    if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.Var) {
        return Data_Set.empty;
    };
    if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.Local) {
        return Data_Set.singleton(localId(v.value1.value0)(v.value1.value1));
    };
    if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.Lit) {
        if (v.value1.value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitArray) {
            return Data_Array.foldl(function (acc) {
                return function (e) {
                    return Data_Set.union(Data_Ord.ordString)(acc)(freeVars(e));
                };
            })(Data_Set.empty)(v.value1.value0.value0);
        };
        if (v.value1.value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitRecord) {
            return Data_Array.foldl(function (acc) {
                return function (v1) {
                    return Data_Set.union(Data_Ord.ordString)(acc)(freeVars(v1.value1));
                };
            })(Data_Set.empty)(v.value1.value0.value0);
        };
        return Data_Set.empty;
    };
    if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.App) {
        return Data_Array.foldl(function (acc) {
            return function (e) {
                return Data_Set.union(Data_Ord.ordString)(acc)(freeVars(e));
            };
        })(freeVars(v.value1.value0))(Data_Array_NonEmpty.toArray(v.value1.value1));
    };
    if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.Abs) {
        var bodyVars = freeVars(v.value1.value1);
        var argsList = Data_Functor.map(Data_Functor.functorArray)(function (v1) {
            return localId(v1.value0)(v1.value1);
        })(Data_Array_NonEmpty.toArray(v.value1.value0));
        return Data_Array.foldl(Data_Function.flip($$delete))(bodyVars)(argsList);
    };
    if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.UncurriedApp) {
        return Data_Array.foldl(function (acc) {
            return function (e) {
                return Data_Set.union(Data_Ord.ordString)(acc)(freeVars(e));
            };
        })(freeVars(v.value1.value0))(v.value1.value1);
    };
    if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.UncurriedAbs) {
        var bodyVars = freeVars(v.value1.value1);
        var argsList = Data_Functor.map(Data_Functor.functorArray)(function (v1) {
            return localId(v1.value0)(v1.value1);
        })(v.value1.value0);
        return Data_Array.foldl(Data_Function.flip($$delete))(bodyVars)(argsList);
    };
    if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.UncurriedEffectApp) {
        return Data_Array.foldl(function (acc) {
            return function (e) {
                return Data_Set.union(Data_Ord.ordString)(acc)(freeVars(e));
            };
        })(freeVars(v.value1.value0))(v.value1.value1);
    };
    if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.UncurriedEffectAbs) {
        var bodyVars = freeVars(v.value1.value1);
        var argsList = Data_Functor.map(Data_Functor.functorArray)(function (v1) {
            return localId(v1.value0)(v1.value1);
        })(v.value1.value0);
        return Data_Array.foldl(Data_Function.flip($$delete))(bodyVars)(argsList);
    };
    if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.Accessor) {
        return freeVars(v.value1.value0);
    };
    if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.Update) {
        return Data_Array.foldl(function (acc) {
            return function (v1) {
                return Data_Set.union(Data_Ord.ordString)(acc)(freeVars(v1.value1));
            };
        })(freeVars(v.value1.value0))(v.value1.value1);
    };
    if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.CtorSaturated) {
        return Data_Array.foldl(function (acc) {
            return function (v1) {
                return Data_Set.union(Data_Ord.ordString)(acc)(freeVars(v1.value1));
            };
        })(Data_Set.empty)(v.value1.value4);
    };
    if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.CtorDef) {
        return Data_Set.empty;
    };
    if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.LetRec) {
        var bodyVars = freeVars(v.value1.value2);
        var bindsVars = Data_Array.foldl(function (acc) {
            return function (v1) {
                return Data_Set.union(Data_Ord.ordString)(acc)(freeVars(v1.value1));
            };
        })(Data_Set.empty)(Data_Array_NonEmpty.toArray(v.value1.value1));
        var bindsList = Data_Functor.map(Data_Functor.functorArray)(function (v1) {
            return localId(new Data_Maybe.Just(v1.value0))(v.value1.value0);
        })(Data_Array_NonEmpty.toArray(v.value1.value1));
        return Data_Array.foldl(Data_Function.flip($$delete))(Data_Set.union(Data_Ord.ordString)(bodyVars)(bindsVars))(bindsList);
    };
    if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.Let) {
        return Data_Set.union(Data_Ord.ordString)(freeVars(v.value1.value2))(Data_Set["delete"](Data_Ord.ordString)(localId(v.value1.value0)(v.value1.value1))(freeVars(v.value1.value3)));
    };
    if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.EffectBind) {
        return Data_Set.union(Data_Ord.ordString)(freeVars(v.value1.value2))(Data_Set["delete"](Data_Ord.ordString)(localId(v.value1.value0)(v.value1.value1))(freeVars(v.value1.value3)));
    };
    if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.EffectPure) {
        return freeVars(v.value1.value0);
    };
    if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.EffectDefer) {
        return freeVars(v.value1.value0);
    };
    if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.Branch) {
        var pairsVars = Data_Array.foldl(function (acc) {
            return function (v1) {
                return Data_Set.union(Data_Ord.ordString)(acc)(Data_Set.union(Data_Ord.ordString)(freeVars(v1.value0))(freeVars(v1.value1)));
            };
        })(Data_Set.empty)(Data_Array_NonEmpty.toArray(v.value1.value0));
        return Data_Set.union(Data_Ord.ordString)(pairsVars)(freeVars(v.value1.value1));
    };
    if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.PrimOp) {
        if (v.value1.value0 instanceof PureScript_Backend_Optimizer_Syntax.Op1) {
            return freeVars(v.value1.value0.value1);
        };
        if (v.value1.value0 instanceof PureScript_Backend_Optimizer_Syntax.Op2) {
            return Data_Set.union(Data_Ord.ordString)(freeVars(v.value1.value0.value1))(freeVars(v.value1.value0.value2));
        };
        throw new Error("Failed pattern match at Gopurs.UsageAnalysis (line 156, column 16 - line 158, column 57): " + [ v.value1.value0.constructor.name ]);
    };
    if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.PrimEffect) {
        return Data_Set.empty;
    };
    if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.PrimUndefined) {
        return Data_Set.empty;
    };
    if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.Fail) {
        return Data_Set.empty;
    };
    if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.Typed) {
        return freeVars(v.value1.value1);
    };
    throw new Error("Failed pattern match at Gopurs.UsageAnalysis (line 106, column 31 - line 162, column 26): " + [ v.value1.constructor.name ]);
};
var addUsages = /* #__PURE__ */ Data_Map_Internal.unionWith(Data_Ord.ordString)(/* #__PURE__ */ Data_Semiring.add(Data_Semiring.semiringInt));

// | Computes the usage count of each free variable in a given `TcoExpr`.
var usageCount = function (v) {
    if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.Var) {
        return Data_Map_Internal.empty;
    };
    if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.Local) {
        return Data_Map_Internal.singleton(localId(v.value1.value0)(v.value1.value1))(1);
    };
    if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.Lit) {
        if (v.value1.value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitArray) {
            return Data_Array.foldl(function (acc) {
                return function (e) {
                    return addUsages(acc)(usageCount(e));
                };
            })(Data_Map_Internal.empty)(v.value1.value0.value0);
        };
        if (v.value1.value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitRecord) {
            return Data_Array.foldl(function (acc) {
                return function (v1) {
                    return addUsages(acc)(usageCount(v1.value1));
                };
            })(Data_Map_Internal.empty)(v.value1.value0.value0);
        };
        return Data_Map_Internal.empty;
    };
    if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.App) {
        return Data_Array.foldl(function (acc) {
            return function (e) {
                return addUsages(acc)(usageCount(e));
            };
        })(usageCount(v.value1.value0))(Data_Array_NonEmpty.toArray(v.value1.value1));
    };
    if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.Abs) {
        var bodyVars = usageCount(v.value1.value1);
        var argsList = Data_Functor.map(Data_Functor.functorArray)(function (v1) {
            return localId(v1.value0)(v1.value1);
        })(Data_Array_NonEmpty.toArray(v.value1.value0));
        return Data_Array.foldl(Data_Function.flip(delete1))(bodyVars)(argsList);
    };
    if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.UncurriedApp) {
        return Data_Array.foldl(function (acc) {
            return function (e) {
                return addUsages(acc)(usageCount(e));
            };
        })(usageCount(v.value1.value0))(v.value1.value1);
    };
    if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.UncurriedAbs) {
        var bodyVars = usageCount(v.value1.value1);
        var argsList = Data_Functor.map(Data_Functor.functorArray)(function (v1) {
            return localId(v1.value0)(v1.value1);
        })(v.value1.value0);
        return Data_Array.foldl(Data_Function.flip(delete1))(bodyVars)(argsList);
    };
    if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.UncurriedEffectApp) {
        return Data_Array.foldl(function (acc) {
            return function (e) {
                return addUsages(acc)(usageCount(e));
            };
        })(usageCount(v.value1.value0))(v.value1.value1);
    };
    if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.UncurriedEffectAbs) {
        var bodyVars = usageCount(v.value1.value1);
        var argsList = Data_Functor.map(Data_Functor.functorArray)(function (v1) {
            return localId(v1.value0)(v1.value1);
        })(v.value1.value0);
        return Data_Array.foldl(Data_Function.flip(delete1))(bodyVars)(argsList);
    };
    if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.Accessor) {
        return usageCount(v.value1.value0);
    };
    if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.Update) {
        return Data_Array.foldl(function (acc) {
            return function (v1) {
                return addUsages(acc)(usageCount(v1.value1));
            };
        })(usageCount(v.value1.value0))(v.value1.value1);
    };
    if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.CtorSaturated) {
        return Data_Array.foldl(function (acc) {
            return function (v1) {
                return addUsages(acc)(usageCount(v1.value1));
            };
        })(Data_Map_Internal.empty)(v.value1.value4);
    };
    if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.CtorDef) {
        return Data_Map_Internal.empty;
    };
    if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.LetRec) {
        var bodyVars = usageCount(v.value1.value2);
        var bindsVars = Data_Array.foldl(function (acc) {
            return function (v1) {
                return addUsages(acc)(usageCount(v1.value1));
            };
        })(Data_Map_Internal.empty)(Data_Array_NonEmpty.toArray(v.value1.value1));
        var bindsList = Data_Functor.map(Data_Functor.functorArray)(function (v1) {
            return localId(new Data_Maybe.Just(v1.value0))(v.value1.value0);
        })(Data_Array_NonEmpty.toArray(v.value1.value1));
        return Data_Array.foldl(Data_Function.flip(delete1))(addUsages(bodyVars)(bindsVars))(bindsList);
    };
    if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.Let) {
        return addUsages(usageCount(v.value1.value2))(Data_Map_Internal["delete"](Data_Ord.ordString)(localId(v.value1.value0)(v.value1.value1))(usageCount(v.value1.value3)));
    };
    if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.EffectBind) {
        return addUsages(usageCount(v.value1.value2))(Data_Map_Internal["delete"](Data_Ord.ordString)(localId(v.value1.value0)(v.value1.value1))(usageCount(v.value1.value3)));
    };
    if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.EffectPure) {
        return usageCount(v.value1.value0);
    };
    if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.EffectDefer) {
        return usageCount(v.value1.value0);
    };
    if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.Branch) {
        var pairsVars = Data_Array.foldl(function (acc) {
            return function (v1) {
                return addUsages(acc)(maxUsages(usageCount(v1.value0))(usageCount(v1.value1)));
            };
        })(Data_Map_Internal.empty)(Data_Array_NonEmpty.toArray(v.value1.value0));
        var defVars = usageCount(v.value1.value1);
        return foldBranchPairs(Data_Array_NonEmpty.toArray(v.value1.value0))(usageCount(v.value1.value1));
    };
    if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.PrimOp) {
        if (v.value1.value0 instanceof PureScript_Backend_Optimizer_Syntax.Op1) {
            return usageCount(v.value1.value0.value1);
        };
        if (v.value1.value0 instanceof PureScript_Backend_Optimizer_Syntax.Op2) {
            return addUsages(usageCount(v.value1.value0.value1))(usageCount(v.value1.value0.value2));
        };
        throw new Error("Failed pattern match at Gopurs.UsageAnalysis (line 90, column 16 - line 92, column 61): " + [ v.value1.value0.constructor.name ]);
    };
    if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.PrimEffect) {
        return Data_Map_Internal.empty;
    };
    if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.PrimUndefined) {
        return Data_Map_Internal.empty;
    };
    if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.Fail) {
        return Data_Map_Internal.empty;
    };
    if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.Typed) {
        return usageCount(v.value1.value1);
    };
    throw new Error("Failed pattern match at Gopurs.UsageAnalysis (line 36, column 33 - line 96, column 28): " + [ v.value1.constructor.name ]);
};
var foldBranchPairs = function (pairs) {
    return function (def) {
        var v = Data_Array.uncons(pairs);
        if (v instanceof Data_Maybe.Nothing) {
            return def;
        };
        if (v instanceof Data_Maybe.Just) {
            var next = foldBranchPairs(v.value0.tail)(def);
            return addUsages(usageCount(v.value0.head.value0))(maxUsages(usageCount(v.value0.head.value1))(next));
        };
        throw new Error("Failed pattern match at Gopurs.UsageAnalysis (line 99, column 29 - line 103, column 70): " + [ v.constructor.name ]);
    };
};
export {
    sanitizeName,
    localId,
    addUsages,
    maxUsages,
    usageCount,
    foldBranchPairs,
    freeVars
};
