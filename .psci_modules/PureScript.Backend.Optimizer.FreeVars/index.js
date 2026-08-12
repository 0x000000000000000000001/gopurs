import * as Data_Array from "../Data.Array/index.js";
import * as Data_Array_NonEmpty from "../Data.Array.NonEmpty/index.js";
import * as Data_Map_Internal from "../Data.Map.Internal/index.js";
import * as Data_Maybe from "../Data.Maybe/index.js";
import * as Data_Ord from "../Data.Ord/index.js";
import * as Data_Set from "../Data.Set/index.js";
import * as Data_Show from "../Data.Show/index.js";
import * as Data_String_Common from "../Data.String.Common/index.js";
import * as PureScript_Backend_Optimizer_CoreFn from "../PureScript.Backend.Optimizer.CoreFn/index.js";
import * as PureScript_Backend_Optimizer_Syntax from "../PureScript.Backend.Optimizer.Syntax/index.js";
var sanitizeName = function (name) {
    var s1 = Data_String_Common.replaceAll("\"")("_quote_")(Data_String_Common.replaceAll(".")("_dot_")(Data_String_Common.replaceAll("'")("_prime")(Data_String_Common.replaceAll("$")("_dollar")(name))));
    var $16 = s1 === "break" || (s1 === "default" || (s1 === "func" || (s1 === "interface" || (s1 === "select" || (s1 === "case" || (s1 === "defer" || (s1 === "go" || (s1 === "map" || (s1 === "struct" || (s1 === "chan" || (s1 === "else" || (s1 === "goto" || (s1 === "package" || (s1 === "switch" || (s1 === "const" || (s1 === "fallthrough" || (s1 === "if" || (s1 === "range" || (s1 === "type" || (s1 === "continue" || (s1 === "for" || (s1 === "import" || (s1 === "return" || s1 === "var")))))))))))))))))))))));
    if ($16) {
        return "go__" + s1;
    };
    return s1;
};
var paramTypes = function (v) {
    if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.Typed && (v.value1.value1.value1 instanceof PureScript_Backend_Optimizer_Syntax.Local && v.value1.value1.value1.value0 instanceof Data_Maybe.Just)) {
        return Data_Map_Internal.singleton(v.value1.value1.value1.value0.value0 + ("_" + Data_Show.show(Data_Show.showInt)(v.value1.value1.value1.value1)))(v.value1.value0);
    };
    if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.Typed) {
        return paramTypes(v.value1.value1);
    };
    if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.Var) {
        return Data_Map_Internal.empty;
    };
    if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.Local) {
        return Data_Map_Internal.empty;
    };
    if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.Lit) {
        if (v.value1.value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitArray) {
            return Data_Array.foldl(function (acc) {
                return function (e) {
                    return Data_Map_Internal.union(Data_Ord.ordString)(acc)(paramTypes(e));
                };
            })(Data_Map_Internal.empty)(v.value1.value0.value0);
        };
        if (v.value1.value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitRecord) {
            return Data_Array.foldl(function (acc) {
                return function (prop) {
                    return Data_Map_Internal.union(Data_Ord.ordString)(acc)(paramTypes(PureScript_Backend_Optimizer_CoreFn.propValue(prop)));
                };
            })(Data_Map_Internal.empty)(v.value1.value0.value0);
        };
        return Data_Map_Internal.empty;
    };
    if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.App) {
        return Data_Array.foldl(function (acc) {
            return function (e) {
                return Data_Map_Internal.union(Data_Ord.ordString)(acc)(paramTypes(e));
            };
        })(paramTypes(v.value1.value0))(Data_Array_NonEmpty.toArray(v.value1.value1));
    };
    if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.Abs) {
        return paramTypes(v.value1.value1);
    };
    if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.UncurriedApp) {
        return Data_Array.foldl(function (acc) {
            return function (e) {
                return Data_Map_Internal.union(Data_Ord.ordString)(acc)(paramTypes(e));
            };
        })(paramTypes(v.value1.value0))(v.value1.value1);
    };
    if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.UncurriedAbs) {
        return paramTypes(v.value1.value1);
    };
    if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.UncurriedEffectApp) {
        return Data_Array.foldl(function (acc) {
            return function (e) {
                return Data_Map_Internal.union(Data_Ord.ordString)(acc)(paramTypes(e));
            };
        })(paramTypes(v.value1.value0))(v.value1.value1);
    };
    if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.UncurriedEffectAbs) {
        return paramTypes(v.value1.value1);
    };
    if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.Accessor) {
        return paramTypes(v.value1.value0);
    };
    if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.Update) {
        return Data_Array.foldl(function (acc) {
            return function (prop) {
                return Data_Map_Internal.union(Data_Ord.ordString)(acc)(paramTypes(PureScript_Backend_Optimizer_CoreFn.propValue(prop)));
            };
        })(paramTypes(v.value1.value0))(v.value1.value1);
    };
    if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.CtorSaturated) {
        return Data_Array.foldl(function (acc) {
            return function (v1) {
                return Data_Map_Internal.union(Data_Ord.ordString)(acc)(paramTypes(v1.value1));
            };
        })(Data_Map_Internal.empty)(v.value1.value4);
    };
    if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.CtorDef) {
        return Data_Map_Internal.empty;
    };
    if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.LetRec) {
        return Data_Array.foldl(function (acc) {
            return function (v1) {
                return Data_Map_Internal.union(Data_Ord.ordString)(acc)(paramTypes(v1.value1));
            };
        })(paramTypes(v.value1.value2))(Data_Array_NonEmpty.toArray(v.value1.value1));
    };
    if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.Let) {
        return Data_Map_Internal.union(Data_Ord.ordString)(paramTypes(v.value1.value2))(paramTypes(v.value1.value3));
    };
    if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.EffectBind) {
        return Data_Map_Internal.union(Data_Ord.ordString)(paramTypes(v.value1.value2))(paramTypes(v.value1.value3));
    };
    if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.EffectPure) {
        return paramTypes(v.value1.value0);
    };
    if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.EffectDefer) {
        return paramTypes(v.value1.value0);
    };
    if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.Branch) {
        return Data_Array.foldl(function (acc) {
            return function (v1) {
                return Data_Map_Internal.union(Data_Ord.ordString)(acc)(Data_Map_Internal.union(Data_Ord.ordString)(paramTypes(v1.value0))(paramTypes(v1.value1)));
            };
        })(paramTypes(v.value1.value1))(Data_Array_NonEmpty.toArray(v.value1.value0));
    };
    if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.PrimOp) {
        if (v.value1.value0 instanceof PureScript_Backend_Optimizer_Syntax.Op1) {
            return paramTypes(v.value1.value0.value1);
        };
        if (v.value1.value0 instanceof PureScript_Backend_Optimizer_Syntax.Op2) {
            return Data_Map_Internal.union(Data_Ord.ordString)(paramTypes(v.value1.value0.value1))(paramTypes(v.value1.value0.value2));
        };
        throw new Error("Failed pattern match at PureScript.Backend.Optimizer.FreeVars (line 117, column 16 - line 119, column 61): " + [ v.value1.value0.constructor.name ]);
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
    throw new Error("Failed pattern match at PureScript.Backend.Optimizer.FreeVars (line 91, column 31 - line 122, column 22): " + [ v.value1.constructor.name ]);
};
var localId = function (v) {
    return function (v1) {
        if (v instanceof Data_Maybe.Just) {
            return sanitizeName(v.value0) + ("_" + Data_Show.show(Data_Show.showInt)(v1));
        };
        if (v instanceof Data_Maybe.Nothing) {
            return "__local_var_" + Data_Show.show(Data_Show.showInt)(v1);
        };
        throw new Error("Failed pattern match at PureScript.Backend.Optimizer.FreeVars (line 25, column 1 - line 25, column 42): " + [ v.constructor.name, v1.constructor.name ]);
    };
};

// | Computes the set of free variables (by unique localId) in a given `TcoExpr`.
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
        var argsSet = Data_Array.foldl(function (acc) {
            return function (v1) {
                return Data_Set.insert(Data_Ord.ordString)(localId(v1.value0)(v1.value1))(acc);
            };
        })(Data_Set.empty)(Data_Array_NonEmpty.toArray(v.value1.value0));
        return Data_Set.difference(Data_Ord.ordString)(bodyVars)(argsSet);
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
        var argsSet = Data_Array.foldl(function (acc) {
            return function (v1) {
                return Data_Set.insert(Data_Ord.ordString)(localId(v1.value0)(v1.value1))(acc);
            };
        })(Data_Set.empty)(v.value1.value0);
        return Data_Set.difference(Data_Ord.ordString)(bodyVars)(argsSet);
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
        var argsSet = Data_Array.foldl(function (acc) {
            return function (v1) {
                return Data_Set.insert(Data_Ord.ordString)(localId(v1.value0)(v1.value1))(acc);
            };
        })(Data_Set.empty)(v.value1.value0);
        return Data_Set.difference(Data_Ord.ordString)(bodyVars)(argsSet);
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
        var bindsSet = Data_Array.foldl(function (acc) {
            return function (v1) {
                return Data_Set.insert(Data_Ord.ordString)(localId(new Data_Maybe.Just(v1.value0))(v.value1.value0))(acc);
            };
        })(Data_Set.empty)(Data_Array_NonEmpty.toArray(v.value1.value1));
        return Data_Set.difference(Data_Ord.ordString)(Data_Set.union(Data_Ord.ordString)(bodyVars)(bindsVars))(bindsSet);
    };
    if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.Let) {
        return Data_Set.union(Data_Ord.ordString)(freeVars(v.value1.value2))(Data_Set.difference(Data_Ord.ordString)(freeVars(v.value1.value3))(Data_Set.singleton(localId(v.value1.value0)(v.value1.value1))));
    };
    if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.EffectBind) {
        return Data_Set.union(Data_Ord.ordString)(freeVars(v.value1.value2))(Data_Set.difference(Data_Ord.ordString)(freeVars(v.value1.value3))(Data_Set.singleton(localId(v.value1.value0)(v.value1.value1))));
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
        var defVars = freeVars(v.value1.value1);
        return Data_Set.union(Data_Ord.ordString)(defVars)(pairsVars);
    };
    if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.PrimOp) {
        if (v.value1.value0 instanceof PureScript_Backend_Optimizer_Syntax.Op1) {
            return freeVars(v.value1.value0.value1);
        };
        if (v.value1.value0 instanceof PureScript_Backend_Optimizer_Syntax.Op2) {
            return Data_Set.union(Data_Ord.ordString)(freeVars(v.value1.value0.value1))(freeVars(v.value1.value0.value2));
        };
        throw new Error("Failed pattern match at PureScript.Backend.Optimizer.FreeVars (line 82, column 16 - line 84, column 57): " + [ v.value1.value0.constructor.name ]);
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
    throw new Error("Failed pattern match at PureScript.Backend.Optimizer.FreeVars (line 31, column 31 - line 88, column 26): " + [ v.value1.constructor.name ]);
};
export {
    sanitizeName,
    localId,
    freeVars,
    paramTypes
};
