import * as Data_Array from "../Data.Array/index.js";
import * as Data_Array_NonEmpty_Internal from "../Data.Array.NonEmpty.Internal/index.js";
import * as Data_Foldable from "../Data.Foldable/index.js";
import * as Data_Functor from "../Data.Functor/index.js";
import * as Data_Map_Internal from "../Data.Map.Internal/index.js";
import * as Data_Maybe from "../Data.Maybe/index.js";
import * as Data_Monoid from "../Data.Monoid/index.js";
import * as Data_Newtype from "../Data.Newtype/index.js";
import * as Data_Ord from "../Data.Ord/index.js";
import * as Data_Tuple from "../Data.Tuple/index.js";
import * as PureScript_Backend_Optimizer_Codegen_Tco from "../PureScript.Backend.Optimizer.Codegen.Tco/index.js";
import * as PureScript_Backend_Optimizer_CoreFn from "../PureScript.Backend.Optimizer.CoreFn/index.js";
import * as PureScript_Backend_Optimizer_Syntax from "../PureScript.Backend.Optimizer.Syntax/index.js";
var mempty = /* #__PURE__ */ Data_Monoid.mempty(PureScript_Backend_Optimizer_Codegen_Tco.monoidTcoAnalysis);
var unify = function (v) {
    return function (v1) {
        return function (v2) {
            if (v instanceof PureScript_Backend_Optimizer_CoreFn.TypeVar) {
                return Data_Map_Internal.insert(Data_Ord.ordString)(v.value0)(v1)(v2);
            };
            if (v instanceof PureScript_Backend_Optimizer_CoreFn.Func && v1 instanceof PureScript_Backend_Optimizer_CoreFn.Func) {
                var len2 = Data_Array.length(v1.value0);
                var len1 = Data_Array.length(v.value0);
                var args2$prime = (function () {
                    var $28 = len2 > len1;
                    if ($28) {
                        return Data_Array.drop(len2 - len1 | 0)(v1.value0);
                    };
                    return v1.value0;
                })();
                var args1$prime = (function () {
                    var $29 = len1 > len2;
                    if ($29) {
                        return Data_Array.drop(len1 - len2 | 0)(v.value0);
                    };
                    return v.value0;
                })();
                var acc$prime = Data_Foldable.foldl(Data_Foldable.foldableArray)(function (acc$prime$prime) {
                    return function (v3) {
                        return unify(v3.value0)(v3.value1)(acc$prime$prime);
                    };
                })(v2)(Data_Array.zip(args1$prime)(args2$prime));
                return unify(v.value1)(v1.value1)(acc$prime);
            };
            if (v instanceof PureScript_Backend_Optimizer_CoreFn["Array"] && v1 instanceof PureScript_Backend_Optimizer_CoreFn["Array"]) {
                return unify(v.value0)(v1.value0)(v2);
            };
            if (v instanceof PureScript_Backend_Optimizer_CoreFn.Record && v1 instanceof PureScript_Backend_Optimizer_CoreFn.Record) {
                return unify(v.value0)(v1.value0)(v2);
            };
            if (v instanceof PureScript_Backend_Optimizer_CoreFn.Row && v1 instanceof PureScript_Backend_Optimizer_CoreFn.Row) {
                return Data_Foldable.foldl(Data_Foldable.foldableArray)(function (acc$prime) {
                    return function (v3) {
                        return unify(v3.value0.value1)(v3.value1.value1)(acc$prime);
                    };
                })(v2)(Data_Array.zip(v.value0)(v1.value0));
            };
            if (v instanceof PureScript_Backend_Optimizer_CoreFn.TypeApp && v1 instanceof PureScript_Backend_Optimizer_CoreFn.TypeApp) {
                return Data_Foldable.foldl(Data_Foldable.foldableArray)(function (acc$prime) {
                    return function (v3) {
                        return unify(v3.value0)(v3.value1)(acc$prime);
                    };
                })(unify(v.value0)(v1.value0)(v2))(Data_Array.zip(v.value1)(v1.value1));
            };
            if (v instanceof PureScript_Backend_Optimizer_CoreFn.ForAll && v1 instanceof PureScript_Backend_Optimizer_CoreFn.ForAll) {
                return unify(v.value1)(v1.value1)(v2);
            };
            if (v instanceof PureScript_Backend_Optimizer_CoreFn.ConstrainedType && v1 instanceof PureScript_Backend_Optimizer_CoreFn.ConstrainedType) {
                return unify(v.value1)(v1.value1)(v2);
            };
            return v2;
        };
    };
};
var substituteExprType = function (subst) {
    return function (t) {
        if (t instanceof PureScript_Backend_Optimizer_CoreFn.TypeVar) {
            var v = Data_Map_Internal.lookup(Data_Ord.ordString)(t.value0)(subst);
            if (v instanceof Data_Maybe.Just) {
                return v.value0;
            };
            if (v instanceof Data_Maybe.Nothing) {
                return t;
            };
            throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Substitute (line 42, column 16 - line 44, column 17): " + [ v.constructor.name ]);
        };
        if (t instanceof PureScript_Backend_Optimizer_CoreFn.Func) {
            return new PureScript_Backend_Optimizer_CoreFn.Func(Data_Functor.map(Data_Functor.functorArray)(substituteExprType(subst))(t.value0), substituteExprType(subst)(t.value1));
        };
        if (t instanceof PureScript_Backend_Optimizer_CoreFn["Array"]) {
            return new PureScript_Backend_Optimizer_CoreFn["Array"](substituteExprType(subst)(t.value0));
        };
        if (t instanceof PureScript_Backend_Optimizer_CoreFn.Record) {
            return new PureScript_Backend_Optimizer_CoreFn.Record(substituteExprType(subst)(t.value0));
        };
        if (t instanceof PureScript_Backend_Optimizer_CoreFn.Row) {
            return new PureScript_Backend_Optimizer_CoreFn.Row(Data_Functor.map(Data_Functor.functorArray)(function (v) {
                return new Data_Tuple.Tuple(v.value0, substituteExprType(subst)(v.value1));
            })(t.value0), Data_Functor.map(Data_Maybe.functorMaybe)(substituteExprType(subst))(t.value1));
        };
        if (t instanceof PureScript_Backend_Optimizer_CoreFn.TypeApp) {
            return new PureScript_Backend_Optimizer_CoreFn.TypeApp(substituteExprType(subst)(t.value0), Data_Functor.map(Data_Functor.functorArray)(substituteExprType(subst))(t.value1));
        };
        if (t instanceof PureScript_Backend_Optimizer_CoreFn.ForAll) {
            return new PureScript_Backend_Optimizer_CoreFn.ForAll(t.value0, substituteExprType(subst)(t.value1));
        };
        if (t instanceof PureScript_Backend_Optimizer_CoreFn.ConstrainedType) {
            return new PureScript_Backend_Optimizer_CoreFn.ConstrainedType(Data_Functor.map(Data_Functor.functorArray)(function (v) {
                return new Data_Tuple.Tuple(v.value0, Data_Functor.map(Data_Functor.functorArray)(substituteExprType(subst))(v.value1));
            })(t.value0), substituteExprType(subst)(t.value1));
        };
        return t;
    };
};
var substituteAst = function (insts) {
    return function (mangle) {
        var go = function (mbTy) {
            return function (v) {
                if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.Typed) {
                    return new PureScript_Backend_Optimizer_Codegen_Tco.TcoExpr(v.value0, new PureScript_Backend_Optimizer_Syntax.Typed(v.value1.value0, go(new Data_Maybe.Just(v.value1.value0))(v.value1.value1)));
                };
                if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.Var) {
                    var fullName = (function () {
                        if (v.value1.value0.value0 instanceof Data_Maybe.Just) {
                            return Data_Newtype.unwrap()(v.value1.value0.value0.value0) + ("." + v.value1.value0.value1);
                        };
                        if (v.value1.value0.value0 instanceof Data_Maybe.Nothing) {
                            return v.value1.value0.value1;
                        };
                        throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Substitute (line 71, column 20 - line 73, column 26): " + [ v.value1.value0.value0.constructor.name ]);
                    })();
                    var v1 = Data_Map_Internal.lookup(Data_Ord.ordString)(fullName)(insts);
                    if (v1 instanceof Data_Maybe.Just) {
                        var varTy = (function () {
                            if (mbTy instanceof Data_Maybe.Just) {
                                return mbTy.value0;
                            };
                            if (mbTy instanceof Data_Maybe.Nothing) {
                                return new PureScript_Backend_Optimizer_CoreFn.TypeVar("gopurs_runtime.Value");
                            };
                            throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Substitute (line 77, column 21 - line 79, column 64): " + [ mbTy.constructor.name ]);
                        })();
                        var $99 = Data_Array.elem(PureScript_Backend_Optimizer_CoreFn.eqExprType)(varTy)(v1.value0);
                        if ($99) {
                            var mangled = v.value1.value0.value1 + ("__" + mangle(varTy));
                            return new PureScript_Backend_Optimizer_Codegen_Tco.TcoExpr(v.value0, new PureScript_Backend_Optimizer_Syntax.Var(new PureScript_Backend_Optimizer_CoreFn.Qualified(v.value1.value0.value0, mangled)));
                        };
                        return new PureScript_Backend_Optimizer_Codegen_Tco.TcoExpr(v.value0, new PureScript_Backend_Optimizer_Syntax.Var(new PureScript_Backend_Optimizer_CoreFn.Qualified(v.value1.value0.value0, v.value1.value0.value1)));
                    };
                    if (v1 instanceof Data_Maybe.Nothing) {
                        return new PureScript_Backend_Optimizer_Codegen_Tco.TcoExpr(v.value0, new PureScript_Backend_Optimizer_Syntax.Var(new PureScript_Backend_Optimizer_CoreFn.Qualified(v.value1.value0.value0, v.value1.value0.value1)));
                    };
                    throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Substitute (line 74, column 10 - line 84, column 65): " + [ v1.constructor.name ]);
                };
                if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.App) {
                    return new PureScript_Backend_Optimizer_Codegen_Tco.TcoExpr(v.value0, new PureScript_Backend_Optimizer_Syntax.App(go(Data_Maybe.Nothing.value)(v.value1.value0), Data_Functor.map(Data_Array_NonEmpty_Internal.functorNonEmptyArray)(go(Data_Maybe.Nothing.value))(v.value1.value1)));
                };
                if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.Abs) {
                    return new PureScript_Backend_Optimizer_Codegen_Tco.TcoExpr(v.value0, new PureScript_Backend_Optimizer_Syntax.Abs(v.value1.value0, go(Data_Maybe.Nothing.value)(v.value1.value1)));
                };
                if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.UncurriedApp) {
                    return new PureScript_Backend_Optimizer_Codegen_Tco.TcoExpr(v.value0, new PureScript_Backend_Optimizer_Syntax.UncurriedApp(go(Data_Maybe.Nothing.value)(v.value1.value0), Data_Functor.map(Data_Functor.functorArray)(go(Data_Maybe.Nothing.value))(v.value1.value1)));
                };
                if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.UncurriedAbs) {
                    return new PureScript_Backend_Optimizer_Codegen_Tco.TcoExpr(v.value0, new PureScript_Backend_Optimizer_Syntax.UncurriedAbs(v.value1.value0, go(Data_Maybe.Nothing.value)(v.value1.value1)));
                };
                if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.UncurriedEffectApp) {
                    return new PureScript_Backend_Optimizer_Codegen_Tco.TcoExpr(v.value0, new PureScript_Backend_Optimizer_Syntax.UncurriedEffectApp(go(Data_Maybe.Nothing.value)(v.value1.value0), Data_Functor.map(Data_Functor.functorArray)(go(Data_Maybe.Nothing.value))(v.value1.value1)));
                };
                if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.UncurriedEffectAbs) {
                    return new PureScript_Backend_Optimizer_Codegen_Tco.TcoExpr(v.value0, new PureScript_Backend_Optimizer_Syntax.UncurriedEffectAbs(v.value1.value0, go(Data_Maybe.Nothing.value)(v.value1.value1)));
                };
                if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.Accessor) {
                    return new PureScript_Backend_Optimizer_Codegen_Tco.TcoExpr(v.value0, new PureScript_Backend_Optimizer_Syntax.Accessor(go(Data_Maybe.Nothing.value)(v.value1.value0), v.value1.value1));
                };
                if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.Update) {
                    return new PureScript_Backend_Optimizer_Codegen_Tco.TcoExpr(v.value0, new PureScript_Backend_Optimizer_Syntax.Update(go(Data_Maybe.Nothing.value)(v.value1.value0), Data_Functor.map(Data_Functor.functorArray)(function (v1) {
                        return new PureScript_Backend_Optimizer_CoreFn.Prop(v1.value0, go(Data_Maybe.Nothing.value)(v1.value1));
                    })(v.value1.value1)));
                };
                if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.CtorSaturated) {
                    return new PureScript_Backend_Optimizer_Codegen_Tco.TcoExpr(v.value0, new PureScript_Backend_Optimizer_Syntax.CtorSaturated(v.value1.value0, v.value1.value1, v.value1.value2, v.value1.value3, Data_Functor.map(Data_Functor.functorArray)(function (v1) {
                        return new Data_Tuple.Tuple(v1.value0, go(Data_Maybe.Nothing.value)(v1.value1));
                    })(v.value1.value4)));
                };
                if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.CtorDef) {
                    return new PureScript_Backend_Optimizer_Codegen_Tco.TcoExpr(v.value0, v.value1);
                };
                if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.LetRec) {
                    return new PureScript_Backend_Optimizer_Codegen_Tco.TcoExpr(v.value0, new PureScript_Backend_Optimizer_Syntax.LetRec(v.value1.value0, Data_Functor.map(Data_Array_NonEmpty_Internal.functorNonEmptyArray)(function (v1) {
                        return new Data_Tuple.Tuple(v1.value0, go(Data_Maybe.Nothing.value)(v1.value1));
                    })(v.value1.value1), go(Data_Maybe.Nothing.value)(v.value1.value2)));
                };
                if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.Let) {
                    return new PureScript_Backend_Optimizer_Codegen_Tco.TcoExpr(v.value0, new PureScript_Backend_Optimizer_Syntax.Let(v.value1.value0, v.value1.value1, go(Data_Maybe.Nothing.value)(v.value1.value2), go(Data_Maybe.Nothing.value)(v.value1.value3)));
                };
                if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.EffectBind) {
                    return new PureScript_Backend_Optimizer_Codegen_Tco.TcoExpr(v.value0, new PureScript_Backend_Optimizer_Syntax.EffectBind(v.value1.value0, v.value1.value1, go(Data_Maybe.Nothing.value)(v.value1.value2), go(Data_Maybe.Nothing.value)(v.value1.value3)));
                };
                if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.EffectPure) {
                    return new PureScript_Backend_Optimizer_Codegen_Tco.TcoExpr(v.value0, new PureScript_Backend_Optimizer_Syntax.EffectPure(go(Data_Maybe.Nothing.value)(v.value1.value0)));
                };
                if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.EffectDefer) {
                    return new PureScript_Backend_Optimizer_Codegen_Tco.TcoExpr(v.value0, new PureScript_Backend_Optimizer_Syntax.EffectDefer(go(Data_Maybe.Nothing.value)(v.value1.value0)));
                };
                if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.Branch) {
                    return new PureScript_Backend_Optimizer_Codegen_Tco.TcoExpr(v.value0, new PureScript_Backend_Optimizer_Syntax.Branch(Data_Functor.map(Data_Array_NonEmpty_Internal.functorNonEmptyArray)(function (v1) {
                        return new PureScript_Backend_Optimizer_Syntax.Pair(go(Data_Maybe.Nothing.value)(v1.value0), go(Data_Maybe.Nothing.value)(v1.value1));
                    })(v.value1.value0), go(Data_Maybe.Nothing.value)(v.value1.value1)));
                };
                if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.Lit) {
                    return new PureScript_Backend_Optimizer_Codegen_Tco.TcoExpr(v.value0, new PureScript_Backend_Optimizer_Syntax.Lit((function () {
                        if (v.value1.value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitArray) {
                            return new PureScript_Backend_Optimizer_CoreFn.LitArray(Data_Functor.map(Data_Functor.functorArray)(go(Data_Maybe.Nothing.value))(v.value1.value0.value0));
                        };
                        if (v.value1.value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitRecord) {
                            return new PureScript_Backend_Optimizer_CoreFn.LitRecord(Data_Functor.map(Data_Functor.functorArray)(function (v1) {
                                return new PureScript_Backend_Optimizer_CoreFn.Prop(v1.value0, go(Data_Maybe.Nothing.value)(v1.value1));
                            })(v.value1.value0.value0));
                        };
                        return v.value1.value0;
                    })()));
                };
                if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.PrimEffect) {
                    return new PureScript_Backend_Optimizer_Codegen_Tco.TcoExpr(v.value0, new PureScript_Backend_Optimizer_Syntax.PrimEffect((function () {
                        if (v.value1.value0 instanceof PureScript_Backend_Optimizer_Syntax.EffectRefNew) {
                            return new PureScript_Backend_Optimizer_Syntax.EffectRefNew(go(Data_Maybe.Nothing.value)(v.value1.value0.value0));
                        };
                        if (v.value1.value0 instanceof PureScript_Backend_Optimizer_Syntax.EffectRefRead) {
                            return new PureScript_Backend_Optimizer_Syntax.EffectRefRead(go(Data_Maybe.Nothing.value)(v.value1.value0.value0));
                        };
                        if (v.value1.value0 instanceof PureScript_Backend_Optimizer_Syntax.EffectRefWrite) {
                            return new PureScript_Backend_Optimizer_Syntax.EffectRefWrite(go(Data_Maybe.Nothing.value)(v.value1.value0.value0), go(Data_Maybe.Nothing.value)(v.value1.value0.value1));
                        };
                        return v.value1.value0;
                    })()));
                };
                if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.PrimOp) {
                    return new PureScript_Backend_Optimizer_Codegen_Tco.TcoExpr(v.value0, new PureScript_Backend_Optimizer_Syntax.PrimOp((function () {
                        if (v.value1.value0 instanceof PureScript_Backend_Optimizer_Syntax.Op1) {
                            return new PureScript_Backend_Optimizer_Syntax.Op1(v.value1.value0.value0, go(Data_Maybe.Nothing.value)(v.value1.value0.value1));
                        };
                        if (v.value1.value0 instanceof PureScript_Backend_Optimizer_Syntax.Op2) {
                            return new PureScript_Backend_Optimizer_Syntax.Op2(v.value1.value0.value0, go(Data_Maybe.Nothing.value)(v.value1.value0.value1), go(Data_Maybe.Nothing.value)(v.value1.value0.value2));
                        };
                        throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Substitute (line 110, column 37 - line 112, column 61): " + [ v.value1.value0.constructor.name ]);
                    })()));
                };
                if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.Local) {
                    return new PureScript_Backend_Optimizer_Codegen_Tco.TcoExpr(v.value0, v.value1);
                };
                if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.Fail) {
                    return new PureScript_Backend_Optimizer_Codegen_Tco.TcoExpr(v.value0, v.value1);
                };
                if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.PrimUndefined) {
                    return new PureScript_Backend_Optimizer_Codegen_Tco.TcoExpr(v.value0, v.value1);
                };
                throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Substitute (line 67, column 29 - line 115, column 35): " + [ v.value1.constructor.name ]);
            };
        };
        return go(Data_Maybe.Nothing.value);
    };
};
var setBackendSyntaxType = function (ty) {
    return function (syn) {
        if (syn instanceof PureScript_Backend_Optimizer_Syntax.Typed) {
            return new PureScript_Backend_Optimizer_Syntax.Typed(ty, syn.value1);
        };
        return new PureScript_Backend_Optimizer_Syntax.Typed(ty, new PureScript_Backend_Optimizer_Codegen_Tco.TcoExpr(mempty, syn));
    };
};
var setTcoExprType = function (ty) {
    return function (v) {
        return new PureScript_Backend_Optimizer_Codegen_Tco.TcoExpr(v.value0, setBackendSyntaxType(ty)(v.value1));
    };
};
var mapTcoExprTypes = function (f) {
    var go = function (v) {
        if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.Typed) {
            return new PureScript_Backend_Optimizer_Codegen_Tco.TcoExpr(v.value0, new PureScript_Backend_Optimizer_Syntax.Typed(f(v.value1.value0), go(v.value1.value1)));
        };
        if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.Var) {
            return new PureScript_Backend_Optimizer_Codegen_Tco.TcoExpr(v.value0, new PureScript_Backend_Optimizer_Syntax.Var(v.value1.value0));
        };
        if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.App) {
            return new PureScript_Backend_Optimizer_Codegen_Tco.TcoExpr(v.value0, new PureScript_Backend_Optimizer_Syntax.App(go(v.value1.value0), Data_Functor.map(Data_Array_NonEmpty_Internal.functorNonEmptyArray)(go)(v.value1.value1)));
        };
        if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.Abs) {
            return new PureScript_Backend_Optimizer_Codegen_Tco.TcoExpr(v.value0, new PureScript_Backend_Optimizer_Syntax.Abs(v.value1.value0, go(v.value1.value1)));
        };
        if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.UncurriedApp) {
            return new PureScript_Backend_Optimizer_Codegen_Tco.TcoExpr(v.value0, new PureScript_Backend_Optimizer_Syntax.UncurriedApp(go(v.value1.value0), Data_Functor.map(Data_Functor.functorArray)(go)(v.value1.value1)));
        };
        if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.UncurriedAbs) {
            return new PureScript_Backend_Optimizer_Codegen_Tco.TcoExpr(v.value0, new PureScript_Backend_Optimizer_Syntax.UncurriedAbs(v.value1.value0, go(v.value1.value1)));
        };
        if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.UncurriedEffectApp) {
            return new PureScript_Backend_Optimizer_Codegen_Tco.TcoExpr(v.value0, new PureScript_Backend_Optimizer_Syntax.UncurriedEffectApp(go(v.value1.value0), Data_Functor.map(Data_Functor.functorArray)(go)(v.value1.value1)));
        };
        if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.UncurriedEffectAbs) {
            return new PureScript_Backend_Optimizer_Codegen_Tco.TcoExpr(v.value0, new PureScript_Backend_Optimizer_Syntax.UncurriedEffectAbs(v.value1.value0, go(v.value1.value1)));
        };
        if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.Accessor) {
            return new PureScript_Backend_Optimizer_Codegen_Tco.TcoExpr(v.value0, new PureScript_Backend_Optimizer_Syntax.Accessor(go(v.value1.value0), v.value1.value1));
        };
        if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.Update) {
            return new PureScript_Backend_Optimizer_Codegen_Tco.TcoExpr(v.value0, new PureScript_Backend_Optimizer_Syntax.Update(go(v.value1.value0), Data_Functor.map(Data_Functor.functorArray)(function (v1) {
                return new PureScript_Backend_Optimizer_CoreFn.Prop(v1.value0, go(v1.value1));
            })(v.value1.value1)));
        };
        if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.CtorSaturated) {
            return new PureScript_Backend_Optimizer_Codegen_Tco.TcoExpr(v.value0, new PureScript_Backend_Optimizer_Syntax.CtorSaturated(v.value1.value0, v.value1.value1, v.value1.value2, v.value1.value3, Data_Functor.map(Data_Functor.functorArray)(function (v1) {
                return new Data_Tuple.Tuple(v1.value0, go(v1.value1));
            })(v.value1.value4)));
        };
        if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.CtorDef) {
            return new PureScript_Backend_Optimizer_Codegen_Tco.TcoExpr(v.value0, new PureScript_Backend_Optimizer_Syntax.CtorDef(v.value1.value0, v.value1.value1, v.value1.value2, v.value1.value3));
        };
        if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.LetRec) {
            return new PureScript_Backend_Optimizer_Codegen_Tco.TcoExpr(v.value0, new PureScript_Backend_Optimizer_Syntax.LetRec(v.value1.value0, Data_Functor.map(Data_Array_NonEmpty_Internal.functorNonEmptyArray)(function (v1) {
                return new Data_Tuple.Tuple(v1.value0, go(v1.value1));
            })(v.value1.value1), go(v.value1.value2)));
        };
        if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.Let) {
            return new PureScript_Backend_Optimizer_Codegen_Tco.TcoExpr(v.value0, new PureScript_Backend_Optimizer_Syntax.Let(v.value1.value0, v.value1.value1, go(v.value1.value2), go(v.value1.value3)));
        };
        if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.EffectBind) {
            return new PureScript_Backend_Optimizer_Codegen_Tco.TcoExpr(v.value0, new PureScript_Backend_Optimizer_Syntax.EffectBind(v.value1.value0, v.value1.value1, go(v.value1.value2), go(v.value1.value3)));
        };
        if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.EffectPure) {
            return new PureScript_Backend_Optimizer_Codegen_Tco.TcoExpr(v.value0, new PureScript_Backend_Optimizer_Syntax.EffectPure(go(v.value1.value0)));
        };
        if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.EffectDefer) {
            return new PureScript_Backend_Optimizer_Codegen_Tco.TcoExpr(v.value0, new PureScript_Backend_Optimizer_Syntax.EffectDefer(go(v.value1.value0)));
        };
        if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.Branch) {
            return new PureScript_Backend_Optimizer_Codegen_Tco.TcoExpr(v.value0, new PureScript_Backend_Optimizer_Syntax.Branch(Data_Functor.map(Data_Array_NonEmpty_Internal.functorNonEmptyArray)(function (v1) {
                return new PureScript_Backend_Optimizer_Syntax.Pair(go(v1.value0), go(v1.value1));
            })(v.value1.value0), go(v.value1.value1)));
        };
        if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.Lit) {
            return new PureScript_Backend_Optimizer_Codegen_Tco.TcoExpr(v.value0, new PureScript_Backend_Optimizer_Syntax.Lit((function () {
                if (v.value1.value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitArray) {
                    return new PureScript_Backend_Optimizer_CoreFn.LitArray(Data_Functor.map(Data_Functor.functorArray)(go)(v.value1.value0.value0));
                };
                if (v.value1.value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitRecord) {
                    return new PureScript_Backend_Optimizer_CoreFn.LitRecord(Data_Functor.map(Data_Functor.functorArray)(function (v1) {
                        return new PureScript_Backend_Optimizer_CoreFn.Prop(v1.value0, go(v1.value1));
                    })(v.value1.value0.value0));
                };
                return v.value1.value0;
            })()));
        };
        if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.PrimEffect) {
            return new PureScript_Backend_Optimizer_Codegen_Tco.TcoExpr(v.value0, new PureScript_Backend_Optimizer_Syntax.PrimEffect((function () {
                if (v.value1.value0 instanceof PureScript_Backend_Optimizer_Syntax.EffectRefNew) {
                    return new PureScript_Backend_Optimizer_Syntax.EffectRefNew(go(v.value1.value0.value0));
                };
                if (v.value1.value0 instanceof PureScript_Backend_Optimizer_Syntax.EffectRefRead) {
                    return new PureScript_Backend_Optimizer_Syntax.EffectRefRead(go(v.value1.value0.value0));
                };
                if (v.value1.value0 instanceof PureScript_Backend_Optimizer_Syntax.EffectRefWrite) {
                    return new PureScript_Backend_Optimizer_Syntax.EffectRefWrite(go(v.value1.value0.value0), go(v.value1.value0.value1));
                };
                return v.value1.value0;
            })()));
        };
        if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.PrimOp) {
            return new PureScript_Backend_Optimizer_Codegen_Tco.TcoExpr(v.value0, new PureScript_Backend_Optimizer_Syntax.PrimOp((function () {
                if (v.value1.value0 instanceof PureScript_Backend_Optimizer_Syntax.Op1) {
                    return new PureScript_Backend_Optimizer_Syntax.Op1(v.value1.value0.value0, go(v.value1.value0.value1));
                };
                if (v.value1.value0 instanceof PureScript_Backend_Optimizer_Syntax.Op2) {
                    return new PureScript_Backend_Optimizer_Syntax.Op2(v.value1.value0.value0, go(v.value1.value0.value1), go(v.value1.value0.value2));
                };
                throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Substitute (line 148, column 37 - line 150, column 45): " + [ v.value1.value0.constructor.name ]);
            })()));
        };
        if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.Local) {
            return new PureScript_Backend_Optimizer_Codegen_Tco.TcoExpr(v.value0, new PureScript_Backend_Optimizer_Syntax.Local(v.value1.value0, v.value1.value1));
        };
        if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.Fail) {
            return new PureScript_Backend_Optimizer_Codegen_Tco.TcoExpr(v.value0, new PureScript_Backend_Optimizer_Syntax.Fail(v.value1.value0));
        };
        if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.PrimUndefined) {
            return new PureScript_Backend_Optimizer_Codegen_Tco.TcoExpr(v.value0, PureScript_Backend_Optimizer_Syntax.PrimUndefined.value);
        };
        throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Substitute (line 120, column 24 - line 153, column 45): " + [ v.value1.constructor.name ]);
    };
    return go;
};
export {
    unify,
    substituteExprType,
    setTcoExprType,
    setBackendSyntaxType,
    substituteAst,
    mapTcoExprTypes
};
