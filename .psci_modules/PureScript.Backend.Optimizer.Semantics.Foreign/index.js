import * as Data_Array from "../Data.Array/index.js";
import * as Data_Array_NonEmpty from "../Data.Array.NonEmpty/index.js";
import * as Data_Boolean from "../Data.Boolean/index.js";
import * as Data_Either from "../Data.Either/index.js";
import * as Data_Enum from "../Data.Enum/index.js";
import * as Data_Eq from "../Data.Eq/index.js";
import * as Data_Foldable from "../Data.Foldable/index.js";
import * as Data_Functor from "../Data.Functor/index.js";
import * as Data_List from "../Data.List/index.js";
import * as Data_List_Types from "../Data.List.Types/index.js";
import * as Data_Map_Internal from "../Data.Map.Internal/index.js";
import * as Data_Maybe from "../Data.Maybe/index.js";
import * as Data_Ord from "../Data.Ord/index.js";
import * as Data_Semigroup from "../Data.Semigroup/index.js";
import * as Data_Show from "../Data.Show/index.js";
import * as Data_String_CodePoints from "../Data.String.CodePoints/index.js";
import * as Data_Tuple from "../Data.Tuple/index.js";
import * as PureScript_Backend_Optimizer_CoreFn from "../PureScript.Backend.Optimizer.CoreFn/index.js";
import * as PureScript_Backend_Optimizer_Semantics from "../PureScript.Backend.Optimizer.Semantics/index.js";
import * as PureScript_Backend_Optimizer_Syntax from "../PureScript.Backend.Optimizer.Syntax/index.js";
var fromEnum = /* #__PURE__ */ Data_Enum.fromEnum(Data_String_CodePoints.boundedEnumCodePoint);
var viewCopyRecord = function (v) {
    if (v instanceof PureScript_Backend_Optimizer_Semantics.SemRef && (v.value0 instanceof PureScript_Backend_Optimizer_Semantics.EvalExtern && (v.value0.value0.value0 instanceof Data_Maybe.Just && (v.value0.value0.value0.value0 === "Record.Builder" && (v.value0.value0.value1 === "copyRecord" && (v.value1.length === 1 && (v["value1"][0] instanceof PureScript_Backend_Optimizer_Semantics.ExternApp && v["value1"][0].value0.length === 1))))))) {
        return new Data_Maybe.Just(v["value1"][0]["value0"][0]);
    };
    return Data_Maybe.Nothing.value;
};
var qualified = function (mod) {
    return function (id) {
        return new PureScript_Backend_Optimizer_CoreFn.Qualified(new Data_Maybe.Just(mod), id);
    };
};
var record_builder_copyRecord = /* #__PURE__ */ (function () {
    var go = function (v) {
        return function (v1) {
            return function (v2) {
                if (v2.length === 1 && (v2[0] instanceof PureScript_Backend_Optimizer_Semantics.ExternApp && (v2[0].value0.length === 1 && (v2[0]["value0"][0] instanceof PureScript_Backend_Optimizer_Semantics.NeutLit && v2[0]["value0"][0].value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitRecord)))) {
                    return new Data_Maybe.Just(v2[0]["value0"][0]);
                };
                if (v2.length === 1 && (v2[0] instanceof PureScript_Backend_Optimizer_Semantics.ExternApp && (v2[0].value0.length === 1 && v2[0]["value0"][0] instanceof PureScript_Backend_Optimizer_Semantics.NeutUpdate))) {
                    return new Data_Maybe.Just(v2[0]["value0"][0]);
                };
                return Data_Maybe.Nothing.value;
            };
        };
    };
    return new Data_Tuple.Tuple(qualified("Record.Builder")("copyRecord"), go);
})();
var record_builder_unsafeDelete = /* #__PURE__ */ (function () {
    var go = function (v) {
        return function (v1) {
            return function (v2) {
                if (v2.length === 1 && (v2[0] instanceof PureScript_Backend_Optimizer_Semantics.ExternApp && (v2[0].value0.length === 2 && (v2[0]["value0"][0] instanceof PureScript_Backend_Optimizer_Semantics.NeutLit && (v2[0]["value0"][0].value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitString && (v2[0]["value0"][1] instanceof PureScript_Backend_Optimizer_Semantics.NeutLit && v2[0]["value0"][1].value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitRecord)))))) {
                    return new Data_Maybe.Just(new PureScript_Backend_Optimizer_Semantics.NeutLit(new PureScript_Backend_Optimizer_CoreFn.LitRecord(Data_Array.filter((function () {
                        var $426 = Data_Eq.eq(Data_Eq.eqString)(v2[0]["value0"][0].value0.value0);
                        return function ($427) {
                            return !$426(PureScript_Backend_Optimizer_CoreFn.propKey($427));
                        };
                    })())(v2[0]["value0"][1].value0.value0))));
                };
                return Data_Maybe.Nothing.value;
            };
        };
    };
    return new Data_Tuple.Tuple(qualified("Record.Builder")("unsafeDelete"), go);
})();
var record_builder_unsafeInsert = /* #__PURE__ */ (function () {
    var go = function (v) {
        return function (v1) {
            return function (v2) {
                if (v2.length === 1 && (v2[0] instanceof PureScript_Backend_Optimizer_Semantics.ExternApp && (v2[0].value0.length === 3 && (v2[0]["value0"][0] instanceof PureScript_Backend_Optimizer_Semantics.NeutLit && (v2[0]["value0"][0].value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitString && (v2[0]["value0"][2] instanceof PureScript_Backend_Optimizer_Semantics.NeutLit && v2[0]["value0"][2].value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitRecord)))))) {
                    return new Data_Maybe.Just(new PureScript_Backend_Optimizer_Semantics.NeutLit(new PureScript_Backend_Optimizer_CoreFn.LitRecord(Data_Array.snoc(v2[0]["value0"][2].value0.value0)(new PureScript_Backend_Optimizer_CoreFn.Prop(v2[0]["value0"][0].value0.value0, v2[0]["value0"][1])))));
                };
                if (v2.length === 1 && (v2[0] instanceof PureScript_Backend_Optimizer_Semantics.ExternApp && (v2[0].value0.length === 3 && (v2[0]["value0"][0] instanceof PureScript_Backend_Optimizer_Semantics.NeutLit && (v2[0]["value0"][0].value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitString && v2[0]["value0"][2] instanceof PureScript_Backend_Optimizer_Semantics.NeutUpdate))))) {
                    return new Data_Maybe.Just(PureScript_Backend_Optimizer_Semantics.evalUpdate(v2[0]["value0"][2])([ new PureScript_Backend_Optimizer_CoreFn.Prop(v2[0]["value0"][0].value0.value0, v2[0]["value0"][1]) ]));
                };
                var v3 = function (v4) {
                    return Data_Maybe.Nothing.value;
                };
                if (v2.length === 1 && (v2[0] instanceof PureScript_Backend_Optimizer_Semantics.ExternApp && (v2[0].value0.length === 3 && (v2[0]["value0"][0] instanceof PureScript_Backend_Optimizer_Semantics.NeutLit && v2[0]["value0"][0].value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitString)))) {
                    var $146 = viewCopyRecord(v2[0]["value0"][2]);
                    if ($146 instanceof Data_Maybe.Just) {
                        return new Data_Maybe.Just(PureScript_Backend_Optimizer_Semantics.evalUpdate($146.value0)([ new PureScript_Backend_Optimizer_CoreFn.Prop(v2[0]["value0"][0].value0.value0, v2[0]["value0"][1]) ]));
                    };
                    return v3(true);
                };
                return v3(true);
            };
        };
    };
    return new Data_Tuple.Tuple(qualified("Record.Builder")("unsafeInsert"), go);
})();
var record_builder_unsafeModify = /* #__PURE__ */ (function () {
    var go = function (env) {
        return function (v) {
            return function (v1) {
                if (v1.length === 1 && (v1[0] instanceof PureScript_Backend_Optimizer_Semantics.ExternApp && (v1[0].value0.length === 3 && (v1[0]["value0"][0] instanceof PureScript_Backend_Optimizer_Semantics.NeutLit && (v1[0]["value0"][0].value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitString && (v1[0]["value0"][2] instanceof PureScript_Backend_Optimizer_Semantics.NeutLit && v1[0]["value0"][2].value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitRecord)))))) {
                    var props$prime = Data_Functor.map(Data_Functor.functorArray)(function (v2) {
                        var $157 = v1[0]["value0"][0].value0.value0 === v2.value0;
                        if ($157) {
                            return new PureScript_Backend_Optimizer_CoreFn.Prop(v2.value0, PureScript_Backend_Optimizer_Semantics.evalApp(env)(v1[0]["value0"][1])([ v2.value1 ]));
                        };
                        return new PureScript_Backend_Optimizer_CoreFn.Prop(v2.value0, v2.value1);
                    })(v1[0]["value0"][2].value0.value0);
                    return new Data_Maybe.Just(new PureScript_Backend_Optimizer_Semantics.NeutLit(new PureScript_Backend_Optimizer_CoreFn.LitRecord(props$prime)));
                };
                if (v1.length === 1 && (v1[0] instanceof PureScript_Backend_Optimizer_Semantics.ExternApp && (v1[0].value0.length === 3 && (v1[0]["value0"][0] instanceof PureScript_Backend_Optimizer_Semantics.NeutLit && (v1[0]["value0"][0].value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitString && (v1[0]["value0"][2] instanceof PureScript_Backend_Optimizer_Semantics.NeutUpdate && v1[0]["value0"][2].value0 instanceof PureScript_Backend_Optimizer_Semantics.NeutLocal)))))) {
                    var update = new PureScript_Backend_Optimizer_CoreFn.Prop(v1[0]["value0"][0].value0.value0, PureScript_Backend_Optimizer_Semantics.evalApp(env)(v1[0]["value0"][1])([ PureScript_Backend_Optimizer_Semantics.evalAccessor(env)(v1[0]["value0"][2].value0)(new PureScript_Backend_Optimizer_Syntax.GetProp(v1[0]["value0"][0].value0.value0)) ]));
                    return new Data_Maybe.Just(PureScript_Backend_Optimizer_Semantics.evalUpdate(v1[0]["value0"][2])([ update ]));
                };
                var v2 = function (v3) {
                    return Data_Maybe.Nothing.value;
                };
                if (v1.length === 1 && (v1[0] instanceof PureScript_Backend_Optimizer_Semantics.ExternApp && (v1[0].value0.length === 3 && (v1[0]["value0"][0] instanceof PureScript_Backend_Optimizer_Semantics.NeutLit && v1[0]["value0"][0].value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitString)))) {
                    var $181 = viewCopyRecord(v1[0]["value0"][2]);
                    if ($181 instanceof Data_Maybe.Just) {
                        return new Data_Maybe.Just(PureScript_Backend_Optimizer_Semantics.makeLet(Data_Maybe.Nothing.value)($181.value0)(function (r$prime) {
                            var update = new PureScript_Backend_Optimizer_CoreFn.Prop(v1[0]["value0"][0].value0.value0, PureScript_Backend_Optimizer_Semantics.evalApp(env)(v1[0]["value0"][1])([ PureScript_Backend_Optimizer_Semantics.evalAccessor(env)(r$prime)(new PureScript_Backend_Optimizer_Syntax.GetProp(v1[0]["value0"][0].value0.value0)) ]));
                            return PureScript_Backend_Optimizer_Semantics.evalUpdate($181.value0)([ update ]);
                        }));
                    };
                    return v2(true);
                };
                return v2(true);
            };
        };
    };
    return new Data_Tuple.Tuple(qualified("Record.Builder")("unsafeModify"), go);
})();
var record_builder_unsafeRename = /* #__PURE__ */ (function () {
    var go = function (v) {
        return function (v1) {
            return function (v2) {
                if (v2.length === 1 && (v2[0] instanceof PureScript_Backend_Optimizer_Semantics.ExternApp && (v2[0].value0.length === 3 && (v2[0]["value0"][0] instanceof PureScript_Backend_Optimizer_Semantics.NeutLit && (v2[0]["value0"][0].value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitString && (v2[0]["value0"][1] instanceof PureScript_Backend_Optimizer_Semantics.NeutLit && (v2[0]["value0"][1].value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitString && (v2[0]["value0"][2] instanceof PureScript_Backend_Optimizer_Semantics.NeutLit && v2[0]["value0"][2].value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitRecord)))))))) {
                    var props$prime = Data_Functor.map(Data_Functor.functorArray)(function (v3) {
                        var $192 = v2[0]["value0"][0].value0.value0 === v3.value0;
                        if ($192) {
                            return new PureScript_Backend_Optimizer_CoreFn.Prop(v2[0]["value0"][1].value0.value0, v3.value1);
                        };
                        return new PureScript_Backend_Optimizer_CoreFn.Prop(v3.value0, v3.value1);
                    })(v2[0]["value0"][2].value0.value0);
                    return new Data_Maybe.Just(new PureScript_Backend_Optimizer_Semantics.NeutLit(new PureScript_Backend_Optimizer_CoreFn.LitRecord(props$prime)));
                };
                return Data_Maybe.Nothing.value;
            };
        };
    };
    return new Data_Tuple.Tuple(qualified("Record.Builder")("unsafeRename"), go);
})();
var record_unsafe_union_unsafeUnionFn = /* #__PURE__ */ (function () {
    var go = function (v) {
        return function (v1) {
            return function (v2) {
                if (v2.length === 1 && (v2[0] instanceof PureScript_Backend_Optimizer_Semantics.ExternUncurriedApp && (v2[0].value0.length === 2 && (v2[0]["value0"][0] instanceof PureScript_Backend_Optimizer_Semantics.NeutLit && (v2[0]["value0"][0].value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitRecord && (v2[0]["value0"][1] instanceof PureScript_Backend_Optimizer_Semantics.NeutLit && v2[0]["value0"][1].value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitRecord)))))) {
                    return new Data_Maybe.Just(new PureScript_Backend_Optimizer_Semantics.NeutLit(new PureScript_Backend_Optimizer_CoreFn.LitRecord(Data_Functor.map(Data_Functor.functorArray)(Data_Array_NonEmpty.head)(Data_Array.groupAllBy(Data_Ord.comparing(Data_Ord.ordString)(PureScript_Backend_Optimizer_CoreFn.propKey))(Data_Semigroup.append(Data_Semigroup.semigroupArray)(v2[0]["value0"][0].value0.value0)(v2[0]["value0"][1].value0.value0))))));
                };
                return Data_Maybe.Nothing.value;
            };
        };
    };
    return new Data_Tuple.Tuple(qualified("Record.Unsafe.Union")("unsafeUnionFn"), go);
})();
var record_unsafe_unsafeDelete = /* #__PURE__ */ (function () {
    var go = function (v) {
        return function (v1) {
            return function (v2) {
                if (v2.length === 1 && (v2[0] instanceof PureScript_Backend_Optimizer_Semantics.ExternApp && (v2[0].value0.length === 2 && (v2[0]["value0"][0] instanceof PureScript_Backend_Optimizer_Semantics.NeutLit && (v2[0]["value0"][0].value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitString && (v2[0]["value0"][1] instanceof PureScript_Backend_Optimizer_Semantics.NeutLit && v2[0]["value0"][1].value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitRecord)))))) {
                    return new Data_Maybe.Just(new PureScript_Backend_Optimizer_Semantics.NeutLit(new PureScript_Backend_Optimizer_CoreFn.LitRecord(Data_Array.filter((function () {
                        var $428 = Data_Eq.eq(Data_Eq.eqString)(v2[0]["value0"][0].value0.value0);
                        return function ($429) {
                            return !$428(PureScript_Backend_Optimizer_CoreFn.propKey($429));
                        };
                    })())(v2[0]["value0"][1].value0.value0))));
                };
                return Data_Maybe.Nothing.value;
            };
        };
    };
    return new Data_Tuple.Tuple(qualified("Record.Unsafe")("unsafeDelete"), go);
})();
var record_unsafe_unsafeGet = /* #__PURE__ */ (function () {
    var go = function (env) {
        return function (v) {
            return function (v1) {
                if (v1.length === 1 && (v1[0] instanceof PureScript_Backend_Optimizer_Semantics.ExternApp && (v1[0].value0.length === 1 && (v1[0]["value0"][0] instanceof PureScript_Backend_Optimizer_Semantics.NeutLit && v1[0]["value0"][0].value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitString)))) {
                    return new Data_Maybe.Just(new PureScript_Backend_Optimizer_Semantics.SemLam(Data_Maybe.Nothing.value, function (r) {
                        return PureScript_Backend_Optimizer_Semantics.evalAccessor(env)(r)(new PureScript_Backend_Optimizer_Syntax.GetProp(v1[0]["value0"][0].value0.value0));
                    }));
                };
                return Data_Maybe.Nothing.value;
            };
        };
    };
    return new Data_Tuple.Tuple(qualified("Record.Unsafe")("unsafeGet"), go);
})();
var record_unsafe_unsafeHas = /* #__PURE__ */ (function () {
    var go = function (v) {
        return function (v1) {
            return function (v2) {
                if (v2.length === 1 && (v2[0] instanceof PureScript_Backend_Optimizer_Semantics.ExternApp && (v2[0].value0.length === 2 && (v2[0]["value0"][0] instanceof PureScript_Backend_Optimizer_Semantics.NeutLit && (v2[0]["value0"][0].value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitString && (v2[0]["value0"][1] instanceof PureScript_Backend_Optimizer_Semantics.NeutLit && v2[0]["value0"][1].value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitRecord)))))) {
                    return new Data_Maybe.Just(PureScript_Backend_Optimizer_Semantics.liftBoolean(Data_Array.any((function () {
                        var $430 = Data_Eq.eq(Data_Eq.eqString)(v2[0]["value0"][0].value0.value0);
                        return function ($431) {
                            return $430(PureScript_Backend_Optimizer_CoreFn.propKey($431));
                        };
                    })())(v2[0]["value0"][1].value0.value0)));
                };
                return Data_Maybe.Nothing.value;
            };
        };
    };
    return new Data_Tuple.Tuple(qualified("Record.Unsafe")("unsafeHas"), go);
})();
var record_unsafe_unsafeSet = /* #__PURE__ */ (function () {
    var go = function (v) {
        return function (v1) {
            return function (v2) {
                if (v2.length === 1 && (v2[0] instanceof PureScript_Backend_Optimizer_Semantics.ExternApp && (v2[0].value0.length === 3 && (v2[0]["value0"][0] instanceof PureScript_Backend_Optimizer_Semantics.NeutLit && v2[0]["value0"][0].value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitString)))) {
                    return new Data_Maybe.Just(PureScript_Backend_Optimizer_Semantics.evalUpdate(v2[0]["value0"][2])([ new PureScript_Backend_Optimizer_CoreFn.Prop(v2[0]["value0"][0].value0.value0, v2[0]["value0"][1]) ]));
                };
                return Data_Maybe.Nothing.value;
            };
        };
    };
    return new Data_Tuple.Tuple(qualified("Record.Unsafe")("unsafeSet"), go);
})();
var runEffectFn = function (mod) {
    return function (name) {
        return function (n) {
            var goRunEffectFn = function (env) {
                return function (acc) {
                    return function (head) {
                        return function (v) {
                            if (v instanceof Data_List_Types.Nil) {
                                return PureScript_Backend_Optimizer_Semantics.evalUncurriedEffectApp(env)(head)(acc);
                            };
                            if (v instanceof Data_List_Types.Cons) {
                                return PureScript_Backend_Optimizer_Semantics.makeLet(Data_Maybe.Nothing.value)(v.value0)(function (nextArg) {
                                    return goRunEffectFn(env)(Data_Array.snoc(acc)(nextArg))(head)(v.value1);
                                });
                            };
                            throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Semantics.Foreign (line 294, column 32 - line 299, column 61): " + [ v.constructor.name ]);
                        };
                    };
                };
            };
            var go = function (env) {
                return function (v) {
                    return function (v1) {
                        var v2 = function (v3) {
                            return Data_Maybe.Nothing.value;
                        };
                        if (v1.length === 1 && v1[0] instanceof PureScript_Backend_Optimizer_Semantics.ExternApp) {
                            var $251 = Data_Array.uncons(v1[0].value0);
                            if ($251 instanceof Data_Maybe.Just) {
                                var $252 = Data_Array.length($251.value0.tail) === n;
                                if ($252) {
                                    return new Data_Maybe.Just(goRunEffectFn(env)([  ])($251.value0.head)(Data_List.fromFoldable(Data_Foldable.foldableArray)($251.value0.tail)));
                                };
                                return v2(true);
                            };
                            return v2(true);
                        };
                        return v2(true);
                    };
                };
            };
            return new Data_Tuple.Tuple(qualified(mod)(name + Data_Show.show(Data_Show.showInt)(n)), go);
        };
    };
};
var test_data_undefinedOr_compareUndefinedOrImpl = /* #__PURE__ */ (function () {
    var goCmp = function (env) {
        return function (lt) {
            return function (eq) {
                return function (gt) {
                    return function (compareFn) {
                        return function (a) {
                            return function (b) {
                                if (a instanceof PureScript_Backend_Optimizer_Semantics.NeutPrimUndefined && b instanceof PureScript_Backend_Optimizer_Semantics.NeutPrimUndefined) {
                                    return eq;
                                };
                                if (a instanceof PureScript_Backend_Optimizer_Semantics.NeutPrimUndefined) {
                                    return lt;
                                };
                                if (b instanceof PureScript_Backend_Optimizer_Semantics.NeutPrimUndefined) {
                                    return gt;
                                };
                                return PureScript_Backend_Optimizer_Semantics.evalApp(env)(compareFn)([ a, b ]);
                            };
                        };
                    };
                };
            };
        };
    };
    var go = function (env) {
        return function (v) {
            return function (v1) {
                if (v1.length === 1 && (v1[0] instanceof PureScript_Backend_Optimizer_Semantics.ExternApp && v1[0].value0.length === 6)) {
                    return new Data_Maybe.Just(goCmp(env)(v1[0]["value0"][0])(v1[0]["value0"][1])(v1[0]["value0"][2])(v1[0]["value0"][3])(v1[0]["value0"][4])(v1[0]["value0"][5]));
                };
                if (v1.length === 1 && (v1[0] instanceof PureScript_Backend_Optimizer_Semantics.ExternApp && v1[0].value0.length === 5)) {
                    return new Data_Maybe.Just(new PureScript_Backend_Optimizer_Semantics.SemLam(Data_Maybe.Nothing.value, function (b) {
                        return goCmp(env)(v1[0]["value0"][0])(v1[0]["value0"][1])(v1[0]["value0"][2])(v1[0]["value0"][3])(v1[0]["value0"][4])(b);
                    }));
                };
                if (v1.length === 1 && (v1[0] instanceof PureScript_Backend_Optimizer_Semantics.ExternApp && v1[0].value0.length === 4)) {
                    return new Data_Maybe.Just(new PureScript_Backend_Optimizer_Semantics.SemLam(Data_Maybe.Nothing.value, function (a) {
                        return new PureScript_Backend_Optimizer_Semantics.SemLam(Data_Maybe.Nothing.value, function (b) {
                            return goCmp(env)(v1[0]["value0"][0])(v1[0]["value0"][1])(v1[0]["value0"][2])(v1[0]["value0"][3])(a)(b);
                        });
                    }));
                };
                if (v1.length === 1 && (v1[0] instanceof PureScript_Backend_Optimizer_Semantics.ExternApp && v1[0].value0.length === 3)) {
                    return new Data_Maybe.Just(new PureScript_Backend_Optimizer_Semantics.SemLam(Data_Maybe.Nothing.value, function (compareFn) {
                        return new PureScript_Backend_Optimizer_Semantics.SemLam(Data_Maybe.Nothing.value, function (a) {
                            return new PureScript_Backend_Optimizer_Semantics.SemLam(Data_Maybe.Nothing.value, function (b) {
                                return goCmp(env)(v1[0]["value0"][0])(v1[0]["value0"][1])(v1[0]["value0"][2])(compareFn)(a)(b);
                            });
                        });
                    }));
                };
                if (v1.length === 1 && (v1[0] instanceof PureScript_Backend_Optimizer_Semantics.ExternApp && v1[0].value0.length === 2)) {
                    return new Data_Maybe.Just(new PureScript_Backend_Optimizer_Semantics.SemLam(Data_Maybe.Nothing.value, function (gt) {
                        return new PureScript_Backend_Optimizer_Semantics.SemLam(Data_Maybe.Nothing.value, function (compareFn) {
                            return new PureScript_Backend_Optimizer_Semantics.SemLam(Data_Maybe.Nothing.value, function (a) {
                                return new PureScript_Backend_Optimizer_Semantics.SemLam(Data_Maybe.Nothing.value, function (b) {
                                    return goCmp(env)(v1[0]["value0"][0])(v1[0]["value0"][1])(gt)(compareFn)(a)(b);
                                });
                            });
                        });
                    }));
                };
                if (v1.length === 1 && (v1[0] instanceof PureScript_Backend_Optimizer_Semantics.ExternApp && v1[0].value0.length === 1)) {
                    return new Data_Maybe.Just(new PureScript_Backend_Optimizer_Semantics.SemLam(Data_Maybe.Nothing.value, function (eq) {
                        return new PureScript_Backend_Optimizer_Semantics.SemLam(Data_Maybe.Nothing.value, function (gt) {
                            return new PureScript_Backend_Optimizer_Semantics.SemLam(Data_Maybe.Nothing.value, function (compareFn) {
                                return new PureScript_Backend_Optimizer_Semantics.SemLam(Data_Maybe.Nothing.value, function (a) {
                                    return new PureScript_Backend_Optimizer_Semantics.SemLam(Data_Maybe.Nothing.value, function (b) {
                                        return goCmp(env)(v1[0]["value0"][0])(eq)(gt)(compareFn)(a)(b);
                                    });
                                });
                            });
                        });
                    }));
                };
                if (v1.length === 0) {
                    return new Data_Maybe.Just(new PureScript_Backend_Optimizer_Semantics.SemLam(Data_Maybe.Nothing.value, function (lt) {
                        return new PureScript_Backend_Optimizer_Semantics.SemLam(Data_Maybe.Nothing.value, function (eq) {
                            return new PureScript_Backend_Optimizer_Semantics.SemLam(Data_Maybe.Nothing.value, function (gt) {
                                return new PureScript_Backend_Optimizer_Semantics.SemLam(Data_Maybe.Nothing.value, function (compareFn) {
                                    return new PureScript_Backend_Optimizer_Semantics.SemLam(Data_Maybe.Nothing.value, function (a) {
                                        return new PureScript_Backend_Optimizer_Semantics.SemLam(Data_Maybe.Nothing.value, function (b) {
                                            return goCmp(env)(lt)(eq)(gt)(compareFn)(a)(b);
                                        });
                                    });
                                });
                            });
                        });
                    }));
                };
                return Data_Maybe.Nothing.value;
            };
        };
    };
    return new Data_Tuple.Tuple(qualified("Test.Data.UndefinedOr")("compareUndefinedOrImpl"), go);
})();
var unsafe_coerce_unsafeCoerce = /* #__PURE__ */ (function () {
    var go = function (v) {
        return function (v1) {
            return function (v2) {
                if (v2.length === 1 && (v2[0] instanceof PureScript_Backend_Optimizer_Semantics.ExternApp && v2[0].value0.length === 1)) {
                    return new Data_Maybe.Just(v2[0]["value0"][0]);
                };
                return Data_Maybe.Nothing.value;
            };
        };
    };
    return new Data_Tuple.Tuple(qualified("Unsafe.Coerce")("unsafeCoerce"), go);
})();
var primUnaryOperator = function (op) {
    return function (env) {
        return function (v) {
            return function (v1) {
                if (v1.length === 1 && (v1[0] instanceof PureScript_Backend_Optimizer_Semantics.ExternApp && v1[0].value0.length === 1)) {
                    return new Data_Maybe.Just(PureScript_Backend_Optimizer_Semantics.evalPrimOp(env)(new PureScript_Backend_Optimizer_Syntax.Op1(op, v1[0]["value0"][0])));
                };
                return Data_Maybe.Nothing.value;
            };
        };
    };
};
var primBinaryOperator = function (op) {
    return function (env) {
        return function (v) {
            return function (v1) {
                if (v1.length === 1 && (v1[0] instanceof PureScript_Backend_Optimizer_Semantics.ExternApp && v1[0].value0.length === 1)) {
                    return new Data_Maybe.Just(PureScript_Backend_Optimizer_Semantics.makeLet(Data_Maybe.Nothing.value)(v1[0]["value0"][0])(function (a$prime) {
                        return new PureScript_Backend_Optimizer_Semantics.SemLam(Data_Maybe.Nothing.value, function (b$prime) {
                            return PureScript_Backend_Optimizer_Semantics.evalPrimOp(env)(new PureScript_Backend_Optimizer_Syntax.Op2(op, a$prime, b$prime));
                        });
                    }));
                };
                return Data_Maybe.Nothing.value;
            };
        };
    };
};
var partial_unsafe_unsafePartial = /* #__PURE__ */ (function () {
    var go = function (v) {
        return function (v1) {
            return function (v2) {
                if (v2.length === 1 && (v2[0] instanceof PureScript_Backend_Optimizer_Semantics.ExternApp && (v2[0].value0.length === 1 && v2[0]["value0"][0] instanceof PureScript_Backend_Optimizer_Semantics.SemLam))) {
                    return new Data_Maybe.Just(v2[0]["value0"][0].value1(new PureScript_Backend_Optimizer_Semantics.NeutLit(new PureScript_Backend_Optimizer_CoreFn.LitRecord([  ]))));
                };
                return Data_Maybe.Nothing.value;
            };
        };
    };
    return new Data_Tuple.Tuple(qualified("Partial.Unsafe")("_unsafePartial"), go);
})();
var mkEffectFn = function (mod) {
    return function (name) {
        return function (n) {
            var go = function (env) {
                return function (v) {
                    return function (v1) {
                        if (v1.length === 1 && (v1[0] instanceof PureScript_Backend_Optimizer_Semantics.ExternApp && v1[0].value0.length === 1)) {
                            return new Data_Maybe.Just(new PureScript_Backend_Optimizer_Semantics.SemMkEffectFn(PureScript_Backend_Optimizer_Semantics.evalMkFn(env)(n)(v1[0]["value0"][0])));
                        };
                        return Data_Maybe.Nothing.value;
                    };
                };
            };
            return new Data_Tuple.Tuple(qualified(mod)(name + Data_Show.show(Data_Show.showInt)(n)), go);
        };
    };
};
var isQualified = function (mod) {
    return function (tag) {
        return function (v) {
            if (v.value0 instanceof Data_Maybe.Just) {
                return mod === v.value0.value0 && tag === v.value1;
            };
            return false;
        };
    };
};
var primOrdImplOperator = function (op) {
    return function (env) {
        return function (v) {
            return function (v1) {
                if (v1.length === 2 && (v1[0] instanceof PureScript_Backend_Optimizer_Semantics.ExternApp && (v1[0].value0.length === 5 && (v1[1] instanceof PureScript_Backend_Optimizer_Semantics.ExternPrimOp && v1[1].value0 instanceof PureScript_Backend_Optimizer_Syntax.OpIsTag)))) {
                    if (isQualified("Data.Ordering")("LT")(v1[1].value0.value0)) {
                        return new Data_Maybe.Just(PureScript_Backend_Optimizer_Semantics.evalPrimOp(env)(new PureScript_Backend_Optimizer_Syntax.Op2(op(PureScript_Backend_Optimizer_Syntax.OpLt.value), v1[0]["value0"][3], v1[0]["value0"][4])));
                    };
                    if (isQualified("Data.Ordering")("GT")(v1[1].value0.value0)) {
                        return new Data_Maybe.Just(PureScript_Backend_Optimizer_Semantics.evalPrimOp(env)(new PureScript_Backend_Optimizer_Syntax.Op2(op(PureScript_Backend_Optimizer_Syntax.OpGt.value), v1[0]["value0"][3], v1[0]["value0"][4])));
                    };
                    if (isQualified("Data.Ordering")("EQ")(v1[1].value0.value0)) {
                        return new Data_Maybe.Just(PureScript_Backend_Optimizer_Semantics.evalPrimOp(env)(new PureScript_Backend_Optimizer_Syntax.Op2(op(PureScript_Backend_Optimizer_Syntax.OpEq.value), v1[0]["value0"][3], v1[0]["value0"][4])));
                    };
                };
                return Data_Maybe.Nothing.value;
            };
        };
    };
};
var primOrdOperator = function (op) {
    return function (env) {
        return function (v) {
            return function (v1) {
                if (v1.length === 3 && (v1[0] instanceof PureScript_Backend_Optimizer_Semantics.ExternAccessor && (v1[0].value0 instanceof PureScript_Backend_Optimizer_Syntax.GetProp && (v1[0].value0.value0 === "compare" && (v1[1] instanceof PureScript_Backend_Optimizer_Semantics.ExternApp && (v1[1].value0.length === 2 && (v1[2] instanceof PureScript_Backend_Optimizer_Semantics.ExternPrimOp && v1[2].value0 instanceof PureScript_Backend_Optimizer_Syntax.OpIsTag))))))) {
                    if (isQualified("Data.Ordering")("LT")(v1[2].value0.value0)) {
                        return new Data_Maybe.Just(PureScript_Backend_Optimizer_Semantics.evalPrimOp(env)(new PureScript_Backend_Optimizer_Syntax.Op2(op(PureScript_Backend_Optimizer_Syntax.OpLt.value), v1[1]["value0"][0], v1[1]["value0"][1])));
                    };
                    if (isQualified("Data.Ordering")("GT")(v1[2].value0.value0)) {
                        return new Data_Maybe.Just(PureScript_Backend_Optimizer_Semantics.evalPrimOp(env)(new PureScript_Backend_Optimizer_Syntax.Op2(op(PureScript_Backend_Optimizer_Syntax.OpGt.value), v1[1]["value0"][0], v1[1]["value0"][1])));
                    };
                    if (isQualified("Data.Ordering")("EQ")(v1[2].value0.value0)) {
                        return new Data_Maybe.Just(PureScript_Backend_Optimizer_Semantics.evalPrimOp(env)(new PureScript_Backend_Optimizer_Syntax.Op2(op(PureScript_Backend_Optimizer_Syntax.OpEq.value), v1[1]["value0"][0], v1[1]["value0"][1])));
                    };
                };
                return Data_Maybe.Nothing.value;
            };
        };
    };
};
var effect_uncurried_runEffectFn = /* #__PURE__ */ runEffectFn("Effect.Uncurried")("runEffectFn");
var effect_uncurried_mkEffectFn = /* #__PURE__ */ mkEffectFn("Effect.Uncurried")("mkEffectFn");
var effectUnsafePerform = function (v) {
    return function (v1) {
        return function (v2) {
            if (v2.length === 1 && (v2[0] instanceof PureScript_Backend_Optimizer_Semantics.ExternApp && (v2[0].value0.length === 1 && v2[0]["value0"][0] instanceof PureScript_Backend_Optimizer_Semantics.SemEffectPure))) {
                return new Data_Maybe.Just(v2[0]["value0"][0].value0);
            };
            return Data_Maybe.Nothing.value;
        };
    };
};
var effect_unsafe_unsafePerformEffect = /* #__PURE__ */ (function () {
    return new Data_Tuple.Tuple(qualified("Effect.Unsafe")("unsafePerformEffect"), effectUnsafePerform);
})();
var effectRefWrite = function (v) {
    return function (v1) {
        return function (v2) {
            if (v2.length === 1 && (v2[0] instanceof PureScript_Backend_Optimizer_Semantics.ExternApp && v2[0].value0.length === 2)) {
                return new Data_Maybe.Just(PureScript_Backend_Optimizer_Semantics.makeLet(Data_Maybe.Nothing.value)(v2[0]["value0"][0])(function (val$prime) {
                    return PureScript_Backend_Optimizer_Semantics.makeLet(Data_Maybe.Nothing.value)(v2[0]["value0"][1])(function (ref$prime) {
                        return new PureScript_Backend_Optimizer_Semantics.NeutPrimEffect(new PureScript_Backend_Optimizer_Syntax.EffectRefWrite(ref$prime, val$prime));
                    });
                }));
            };
            return Data_Maybe.Nothing.value;
        };
    };
};
var effect_ref_write = /* #__PURE__ */ (function () {
    return new Data_Tuple.Tuple(qualified("Effect.Ref")("write"), effectRefWrite);
})();
var effectRefRead = function (v) {
    return function (v1) {
        return function (v2) {
            if (v2.length === 1 && (v2[0] instanceof PureScript_Backend_Optimizer_Semantics.ExternApp && v2[0].value0.length === 1)) {
                return new Data_Maybe.Just(PureScript_Backend_Optimizer_Semantics.makeLet(Data_Maybe.Nothing.value)(v2[0]["value0"][0])(function (val$prime) {
                    return new PureScript_Backend_Optimizer_Semantics.NeutPrimEffect(new PureScript_Backend_Optimizer_Syntax.EffectRefRead(val$prime));
                }));
            };
            return Data_Maybe.Nothing.value;
        };
    };
};
var effect_ref_read = /* #__PURE__ */ (function () {
    return new Data_Tuple.Tuple(qualified("Effect.Ref")("read"), effectRefRead);
})();
var effectRefNew = function (v) {
    return function (v1) {
        return function (v2) {
            if (v2.length === 1 && (v2[0] instanceof PureScript_Backend_Optimizer_Semantics.ExternApp && v2[0].value0.length === 1)) {
                return new Data_Maybe.Just(PureScript_Backend_Optimizer_Semantics.makeLet(Data_Maybe.Nothing.value)(v2[0]["value0"][0])(function (val$prime) {
                    return new PureScript_Backend_Optimizer_Semantics.NeutPrimEffect(new PureScript_Backend_Optimizer_Syntax.EffectRefNew(val$prime));
                }));
            };
            return Data_Maybe.Nothing.value;
        };
    };
};
var effect_ref_new = /* #__PURE__ */ (function () {
    return new Data_Tuple.Tuple(qualified("Effect.Ref")("_new"), effectRefNew);
})();
var effectRefModify = function (env) {
    return function (v) {
        return function (v1) {
            if (v1.length === 1 && (v1[0] instanceof PureScript_Backend_Optimizer_Semantics.ExternApp && v1[0].value0.length === 2)) {
                return new Data_Maybe.Just(PureScript_Backend_Optimizer_Semantics.makeLet(Data_Maybe.Nothing.value)(v1[0]["value0"][0])(function (fn$prime) {
                    return PureScript_Backend_Optimizer_Semantics.makeLet(Data_Maybe.Nothing.value)(v1[0]["value0"][1])(function (ref$prime) {
                        return PureScript_Backend_Optimizer_Semantics.makeEffectBind(Data_Maybe.Nothing.value)(new PureScript_Backend_Optimizer_Semantics.NeutPrimEffect(new PureScript_Backend_Optimizer_Syntax.EffectRefRead(ref$prime)))(function (val) {
                            return new PureScript_Backend_Optimizer_Semantics.NeutPrimEffect(new PureScript_Backend_Optimizer_Syntax.EffectRefWrite(ref$prime, PureScript_Backend_Optimizer_Semantics.evalApp(env)(fn$prime)([ val ])));
                        });
                    });
                }));
            };
            return Data_Maybe.Nothing.value;
        };
    };
};
var effect_ref_modify = /* #__PURE__ */ (function () {
    return new Data_Tuple.Tuple(qualified("Effect.Ref")("modify"), effectRefModify);
})();
var effectPure = function (v) {
    return function (v1) {
        return function (v2) {
            if (v2.length === 1 && (v2[0] instanceof PureScript_Backend_Optimizer_Semantics.ExternApp && v2[0].value0.length === 1)) {
                return new Data_Maybe.Just(PureScript_Backend_Optimizer_Semantics.makeLet(Data_Maybe.Nothing.value)(v2[0]["value0"][0])(PureScript_Backend_Optimizer_Semantics.SemEffectPure.create));
            };
            return Data_Maybe.Nothing.value;
        };
    };
};
var effect_pureE = /* #__PURE__ */ (function () {
    return new Data_Tuple.Tuple(qualified("Effect")("pureE"), effectPure);
})();
var effectMap = function (env) {
    return function (v) {
        return function (v1) {
            if (v1.length === 1 && (v1[0] instanceof PureScript_Backend_Optimizer_Semantics.ExternApp && v1[0].value0.length === 1)) {
                return new Data_Maybe.Just(PureScript_Backend_Optimizer_Semantics.makeLet(Data_Maybe.Nothing.value)(v1[0]["value0"][0])(function (fn$prime) {
                    return new PureScript_Backend_Optimizer_Semantics.SemLam(Data_Maybe.Nothing.value, function (val) {
                        return PureScript_Backend_Optimizer_Semantics.makeEffectBind(Data_Maybe.Nothing.value)(val)(function (nextVal) {
                            return new PureScript_Backend_Optimizer_Semantics.SemEffectPure(PureScript_Backend_Optimizer_Semantics.evalApp(env)(fn$prime)([ nextVal ]));
                        });
                    });
                }));
            };
            return Data_Maybe.Nothing.value;
        };
    };
};
var effectBind = function (env) {
    return function (v) {
        return function (v1) {
            if (v1.length === 1 && (v1[0] instanceof PureScript_Backend_Optimizer_Semantics.ExternApp && (v1[0].value0.length === 2 && v1[0]["value0"][1] instanceof PureScript_Backend_Optimizer_Semantics.SemLam))) {
                return new Data_Maybe.Just(PureScript_Backend_Optimizer_Semantics.makeLet(Data_Maybe.Nothing.value)(v1[0]["value0"][0])(function (nextEff) {
                    return PureScript_Backend_Optimizer_Semantics.makeEffectBind(v1[0]["value0"][1].value0)(nextEff)(v1[0]["value0"][1].value1);
                }));
            };
            if (v1.length === 1 && (v1[0] instanceof PureScript_Backend_Optimizer_Semantics.ExternApp && v1[0].value0.length === 2)) {
                return new Data_Maybe.Just(PureScript_Backend_Optimizer_Semantics.makeLet(Data_Maybe.Nothing.value)(v1[0]["value0"][0])(function (nextEff) {
                    return PureScript_Backend_Optimizer_Semantics.makeLet(Data_Maybe.Nothing.value)(v1[0]["value0"][1])(function (nextK) {
                        return PureScript_Backend_Optimizer_Semantics.makeEffectBind(Data_Maybe.Nothing.value)(nextEff)(function (a) {
                            return PureScript_Backend_Optimizer_Semantics.evalApp(env)(nextK)([ a ]);
                        });
                    });
                }));
            };
            return Data_Maybe.Nothing.value;
        };
    };
};
var effect_bindE = /* #__PURE__ */ (function () {
    return new Data_Tuple.Tuple(qualified("Effect")("bindE"), effectBind);
})();
var data_string_codePoints_toCodePointArray = /* #__PURE__ */ (function () {
    var go = function (v) {
        return function (v1) {
            return function (v2) {
                if (v2.length === 1 && (v2[0] instanceof PureScript_Backend_Optimizer_Semantics.ExternApp && (v2[0].value0.length === 1 && (v2[0]["value0"][0] instanceof PureScript_Backend_Optimizer_Semantics.NeutLit && v2[0]["value0"][0].value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitString)))) {
                    return new Data_Maybe.Just(new PureScript_Backend_Optimizer_Semantics.NeutLit(new PureScript_Backend_Optimizer_CoreFn.LitArray(Data_Functor.map(Data_Functor.functorArray)(function ($432) {
                        return PureScript_Backend_Optimizer_Semantics.NeutLit.create(PureScript_Backend_Optimizer_CoreFn.LitInt.create(fromEnum($432)));
                    })(Data_String_CodePoints.toCodePointArray(v2[0]["value0"][0].value0.value0)))));
                };
                return Data_Maybe.Nothing.value;
            };
        };
    };
    return new Data_Tuple.Tuple(qualified("Data.String.CodePoints")("toCodePointArray"), go);
})();
var data_semiring_numMul = /* #__PURE__ */ (function () {
    return new Data_Tuple.Tuple(qualified("Data.Semiring")("numMul"), primBinaryOperator(new PureScript_Backend_Optimizer_Syntax.OpNumberNum(PureScript_Backend_Optimizer_Syntax.OpMultiply.value)));
})();
var data_semiring_numAdd = /* #__PURE__ */ (function () {
    return new Data_Tuple.Tuple(qualified("Data.Semiring")("numAdd"), primBinaryOperator(new PureScript_Backend_Optimizer_Syntax.OpNumberNum(PureScript_Backend_Optimizer_Syntax.OpAdd.value)));
})();
var data_semiring_intMul = /* #__PURE__ */ (function () {
    return new Data_Tuple.Tuple(qualified("Data.Semiring")("intMul"), primBinaryOperator(new PureScript_Backend_Optimizer_Syntax.OpIntNum(PureScript_Backend_Optimizer_Syntax.OpMultiply.value)));
})();
var data_semiring_intAdd = /* #__PURE__ */ (function () {
    return new Data_Tuple.Tuple(qualified("Data.Semiring")("intAdd"), primBinaryOperator(new PureScript_Backend_Optimizer_Syntax.OpIntNum(PureScript_Backend_Optimizer_Syntax.OpAdd.value)));
})();
var data_semigroup_concatString = /* #__PURE__ */ (function () {
    return new Data_Tuple.Tuple(qualified("Data.Semigroup")("concatString"), primBinaryOperator(PureScript_Backend_Optimizer_Syntax.OpStringAppend.value));
})();
var data_semigroup_concatArray = /* #__PURE__ */ (function () {
    var go = function (env) {
        return function (qual) {
            return function (v) {
                if (v.length === 1 && (v[0] instanceof PureScript_Backend_Optimizer_Semantics.ExternApp && (v[0].value0.length === 2 && (v[0]["value0"][0] instanceof PureScript_Backend_Optimizer_Semantics.NeutLit && (v[0]["value0"][0].value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitArray && (v[0]["value0"][1] instanceof PureScript_Backend_Optimizer_Semantics.NeutLit && v[0]["value0"][1].value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitArray)))))) {
                    return new Data_Maybe.Just(new PureScript_Backend_Optimizer_Semantics.NeutLit(new PureScript_Backend_Optimizer_CoreFn.LitArray(Data_Semigroup.append(Data_Semigroup.semigroupArray)(v[0]["value0"][0].value0.value0)(v[0]["value0"][1].value0.value0))));
                };
                if (v.length === 1 && (v[0] instanceof PureScript_Backend_Optimizer_Semantics.ExternApp && v[0].value0.length === 2)) {
                    return new Data_Maybe.Just(PureScript_Backend_Optimizer_Semantics.evalAssocOp(env)(new Data_Either.Left(qual))(v[0]["value0"][0])(v[0]["value0"][1]));
                };
                return Data_Maybe.Nothing.value;
            };
        };
    };
    return new Data_Tuple.Tuple(qualified("Data.Semigroup")("concatArray"), go);
})();
var data_ring_numSub = /* #__PURE__ */ (function () {
    return new Data_Tuple.Tuple(qualified("Data.Ring")("numSub"), primBinaryOperator(new PureScript_Backend_Optimizer_Syntax.OpNumberNum(PureScript_Backend_Optimizer_Syntax.OpSubtract.value)));
})();
var data_ring_intSub = /* #__PURE__ */ (function () {
    return new Data_Tuple.Tuple(qualified("Data.Ring")("intSub"), primBinaryOperator(new PureScript_Backend_Optimizer_Syntax.OpIntNum(PureScript_Backend_Optimizer_Syntax.OpSubtract.value)));
})();
var data_ord_ordString = /* #__PURE__ */ (function () {
    return new Data_Tuple.Tuple(qualified("Data.Ord")("ordString"), primOrdOperator(PureScript_Backend_Optimizer_Syntax.OpStringOrd.create));
})();
var data_ord_ordNumber = /* #__PURE__ */ (function () {
    return new Data_Tuple.Tuple(qualified("Data.Ord")("ordNumber"), primOrdOperator(PureScript_Backend_Optimizer_Syntax.OpNumberOrd.create));
})();
var data_ord_ordIntImpl = /* #__PURE__ */ (function () {
    return new Data_Tuple.Tuple(qualified("Data.Ord")("ordIntImpl"), primOrdImplOperator(PureScript_Backend_Optimizer_Syntax.OpIntOrd.create));
})();
var data_ord_ordInt = /* #__PURE__ */ (function () {
    return new Data_Tuple.Tuple(qualified("Data.Ord")("ordInt"), primOrdOperator(PureScript_Backend_Optimizer_Syntax.OpIntOrd.create));
})();
var data_ord_ordChar = /* #__PURE__ */ (function () {
    return new Data_Tuple.Tuple(qualified("Data.Ord")("ordChar"), primOrdOperator(PureScript_Backend_Optimizer_Syntax.OpCharOrd.create));
})();
var data_ord_ordBoolean = /* #__PURE__ */ (function () {
    return new Data_Tuple.Tuple(qualified("Data.Ord")("ordBoolean"), primOrdOperator(PureScript_Backend_Optimizer_Syntax.OpBooleanOrd.create));
})();
var data_int_bits_zshr = /* #__PURE__ */ (function () {
    return new Data_Tuple.Tuple(qualified("Data.Int.Bits")("zshr"), primBinaryOperator(PureScript_Backend_Optimizer_Syntax.OpIntBitZeroFillShiftRight.value));
})();
var data_int_bits_xor = /* #__PURE__ */ (function () {
    return new Data_Tuple.Tuple(qualified("Data.Int.Bits")("xor"), primBinaryOperator(PureScript_Backend_Optimizer_Syntax.OpIntBitXor.value));
})();
var data_int_bits_shr = /* #__PURE__ */ (function () {
    return new Data_Tuple.Tuple(qualified("Data.Int.Bits")("shr"), primBinaryOperator(PureScript_Backend_Optimizer_Syntax.OpIntBitShiftRight.value));
})();
var data_int_bits_shl = /* #__PURE__ */ (function () {
    return new Data_Tuple.Tuple(qualified("Data.Int.Bits")("shl"), primBinaryOperator(PureScript_Backend_Optimizer_Syntax.OpIntBitShiftLeft.value));
})();
var data_int_bits_or = /* #__PURE__ */ (function () {
    return new Data_Tuple.Tuple(qualified("Data.Int.Bits")("or"), primBinaryOperator(PureScript_Backend_Optimizer_Syntax.OpIntBitOr.value));
})();
var data_int_bits_complement = /* #__PURE__ */ (function () {
    return new Data_Tuple.Tuple(qualified("Data.Int.Bits")("complement"), primUnaryOperator(PureScript_Backend_Optimizer_Syntax.OpIntBitNot.value));
})();
var data_int_bits_and = /* #__PURE__ */ (function () {
    return new Data_Tuple.Tuple(qualified("Data.Int.Bits")("and"), primBinaryOperator(PureScript_Backend_Optimizer_Syntax.OpIntBitAnd.value));
})();
var data_heytingAlgebra_boolNot = /* #__PURE__ */ (function () {
    return new Data_Tuple.Tuple(qualified("Data.HeytingAlgebra")("boolNot"), primUnaryOperator(PureScript_Backend_Optimizer_Syntax.OpBooleanNot.value));
})();
var data_heytingAlgebra_boolDisj = /* #__PURE__ */ (function () {
    return new Data_Tuple.Tuple(qualified("Data.HeytingAlgebra")("boolDisj"), primBinaryOperator(PureScript_Backend_Optimizer_Syntax.OpBooleanOr.value));
})();
var data_heytingAlgebra_boolConj = /* #__PURE__ */ (function () {
    return new Data_Tuple.Tuple(qualified("Data.HeytingAlgebra")("boolConj"), primBinaryOperator(PureScript_Backend_Optimizer_Syntax.OpBooleanAnd.value));
})();
var data_function_uncurried_runFn = function (n) {
    var goRunFn = function (env) {
        return function (n$prime) {
            return function (head) {
                return function (tail) {
                    if (n$prime <= 0) {
                        return PureScript_Backend_Optimizer_Semantics.evalUncurriedApp(env)(head)(tail);
                    };
                    if (Data_Boolean.otherwise) {
                        return new PureScript_Backend_Optimizer_Semantics.SemLam(Data_Maybe.Nothing.value, function (val) {
                            return goRunFn(env)(n$prime - 1 | 0)(head)(Data_Array.snoc(tail)(val));
                        });
                    };
                    throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Semantics.Foreign (line 261, column 3 - line 266, column 58): " + [ env.constructor.name, n$prime.constructor.name, head.constructor.name, tail.constructor.name ]);
                };
            };
        };
    };
    var go = function (env) {
        return function (v) {
            return function (v1) {
                var v2 = function (v3) {
                    return Data_Maybe.Nothing.value;
                };
                if (v1.length === 1 && v1[0] instanceof PureScript_Backend_Optimizer_Semantics.ExternApp) {
                    var $408 = Data_Array.uncons(v1[0].value0);
                    if ($408 instanceof Data_Maybe.Just) {
                        return new Data_Maybe.Just(goRunFn(env)(n - Data_Array.length($408.value0.tail) | 0)($408.value0.head)($408.value0.tail));
                    };
                    return v2(true);
                };
                return v2(true);
            };
        };
    };
    return new Data_Tuple.Tuple(qualified("Data.Function.Uncurried")("runFn" + Data_Show.show(Data_Show.showInt)(n)), go);
};
var data_function_uncurried_mkFn = function (n) {
    var go = function (env) {
        return function (v) {
            return function (v1) {
                if (v1.length === 1 && (v1[0] instanceof PureScript_Backend_Optimizer_Semantics.ExternApp && v1[0].value0.length === 1)) {
                    return new Data_Maybe.Just(new PureScript_Backend_Optimizer_Semantics.SemMkFn(PureScript_Backend_Optimizer_Semantics.evalMkFn(env)(n)(v1[0]["value0"][0])));
                };
                return Data_Maybe.Nothing.value;
            };
        };
    };
    return new Data_Tuple.Tuple(qualified("Data.Function.Uncurried")("mkFn" + Data_Show.show(Data_Show.showInt)(n)), go);
};
var data_euclideanRing_numDiv = /* #__PURE__ */ (function () {
    return new Data_Tuple.Tuple(qualified("Data.EuclideanRing")("numDiv"), primBinaryOperator(new PureScript_Backend_Optimizer_Syntax.OpNumberNum(PureScript_Backend_Optimizer_Syntax.OpDivide.value)));
})();
var data_euclideanRing_intDiv = /* #__PURE__ */ (function () {
    return new Data_Tuple.Tuple(qualified("Data.EuclideanRing")("intDiv"), primBinaryOperator(new PureScript_Backend_Optimizer_Syntax.OpIntNum(PureScript_Backend_Optimizer_Syntax.OpDivide.value)));
})();
var data_eq_eqStringImpl = /* #__PURE__ */ (function () {
    return new Data_Tuple.Tuple(qualified("Data.Eq")("eqStringImpl"), primBinaryOperator(new PureScript_Backend_Optimizer_Syntax.OpStringOrd(PureScript_Backend_Optimizer_Syntax.OpEq.value)));
})();
var data_eq_eqNumberImpl = /* #__PURE__ */ (function () {
    return new Data_Tuple.Tuple(qualified("Data.Eq")("eqNumberImpl"), primBinaryOperator(new PureScript_Backend_Optimizer_Syntax.OpNumberOrd(PureScript_Backend_Optimizer_Syntax.OpEq.value)));
})();
var data_eq_eqIntImpl = /* #__PURE__ */ (function () {
    return new Data_Tuple.Tuple(qualified("Data.Eq")("eqIntImpl"), primBinaryOperator(new PureScript_Backend_Optimizer_Syntax.OpIntOrd(PureScript_Backend_Optimizer_Syntax.OpEq.value)));
})();
var data_eq_eqCharImpl = /* #__PURE__ */ (function () {
    return new Data_Tuple.Tuple(qualified("Data.Eq")("eqCharImpl"), primBinaryOperator(new PureScript_Backend_Optimizer_Syntax.OpCharOrd(PureScript_Backend_Optimizer_Syntax.OpEq.value)));
})();
var data_eq_eqBooleanImpl = /* #__PURE__ */ (function () {
    return new Data_Tuple.Tuple(qualified("Data.Eq")("eqBooleanImpl"), primBinaryOperator(new PureScript_Backend_Optimizer_Syntax.OpBooleanOrd(PureScript_Backend_Optimizer_Syntax.OpEq.value)));
})();
var data_array_unsafeIndexImpl = /* #__PURE__ */ (function () {
    var go = function (env) {
        return function (v) {
            return function (v1) {
                if (v1.length === 1 && (v1[0] instanceof PureScript_Backend_Optimizer_Semantics.ExternUncurriedApp && v1[0].value0.length === 2)) {
                    return new Data_Maybe.Just(PureScript_Backend_Optimizer_Semantics.makeLet(Data_Maybe.Nothing.value)(v1[0]["value0"][0])(function (a$prime) {
                        return PureScript_Backend_Optimizer_Semantics.makeLet(Data_Maybe.Nothing.value)(v1[0]["value0"][1])(function (b$prime) {
                            return PureScript_Backend_Optimizer_Semantics.evalPrimOp(env)(new PureScript_Backend_Optimizer_Syntax.Op2(PureScript_Backend_Optimizer_Syntax.OpArrayIndex.value, a$prime, b$prime));
                        });
                    }));
                };
                if (v1.length === 1 && (v1[0] instanceof PureScript_Backend_Optimizer_Semantics.ExternApp && v1[0].value0.length === 1)) {
                    return new Data_Maybe.Just(PureScript_Backend_Optimizer_Semantics.makeLet(Data_Maybe.Nothing.value)(v1[0]["value0"][0])(function (a$prime) {
                        return new PureScript_Backend_Optimizer_Semantics.SemLam(Data_Maybe.Nothing.value, function (b$prime) {
                            return PureScript_Backend_Optimizer_Semantics.evalPrimOp(env)(new PureScript_Backend_Optimizer_Syntax.Op2(PureScript_Backend_Optimizer_Syntax.OpArrayIndex.value, a$prime, b$prime));
                        });
                    }));
                };
                return Data_Maybe.Nothing.value;
            };
        };
    };
    return new Data_Tuple.Tuple(qualified("Data.Array")("unsafeIndexImpl"), go);
})();
var data_array_length = /* #__PURE__ */ (function () {
    return new Data_Tuple.Tuple(qualified("Data.Array")("length"), primUnaryOperator(PureScript_Backend_Optimizer_Syntax.OpArrayLength.value));
})();
var control_monad_st_uncurried_runSTFn = /* #__PURE__ */ runEffectFn("Control.Monad.ST.Uncurried")("runSTFn");
var control_monad_st_uncurried_mkSTFn = /* #__PURE__ */ mkEffectFn("Control.Monad.ST.Uncurried")("mkSTFn");
var control_monad_st_internal_write = /* #__PURE__ */ (function () {
    return new Data_Tuple.Tuple(qualified("Control.Monad.ST.Internal")("write"), effectRefWrite);
})();
var control_monad_st_internal_run = /* #__PURE__ */ (function () {
    return new Data_Tuple.Tuple(qualified("Control.Monad.ST.Internal")("run"), effectUnsafePerform);
})();
var control_monad_st_internal_read = /* #__PURE__ */ (function () {
    return new Data_Tuple.Tuple(qualified("Control.Monad.ST.Internal")("read"), effectRefRead);
})();
var control_monad_st_internal_pure = /* #__PURE__ */ (function () {
    return new Data_Tuple.Tuple(qualified("Control.Monad.ST.Internal")("pure_"), effectPure);
})();
var control_monad_st_internal_new = /* #__PURE__ */ (function () {
    return new Data_Tuple.Tuple(qualified("Control.Monad.ST.Internal")("new"), effectRefNew);
})();
var control_monad_st_internal_modify = /* #__PURE__ */ (function () {
    return new Data_Tuple.Tuple(qualified("Control.Monad.ST.Internal")("modify"), effectRefModify);
})();
var control_monad_st_internal_map = /* #__PURE__ */ (function () {
    return new Data_Tuple.Tuple(qualified("Control.Monad.ST.Internal")("map_"), effectMap);
})();
var control_monad_st_internal_bind = /* #__PURE__ */ (function () {
    return new Data_Tuple.Tuple(qualified("Control.Monad.ST.Internal")("bind_"), effectBind);
})();
var coreForeignSemantics = /* #__PURE__ */ (function () {
    var oneToTen = Data_Array.range(1)(10);
    var semantics = Data_Semigroup.append(Data_Semigroup.semigroupArray)([ control_monad_st_internal_bind, control_monad_st_internal_map, control_monad_st_internal_modify, control_monad_st_internal_new, control_monad_st_internal_pure, control_monad_st_internal_read, control_monad_st_internal_run, control_monad_st_internal_write, data_array_length, data_array_unsafeIndexImpl, data_eq_eqBooleanImpl, data_eq_eqCharImpl, data_eq_eqIntImpl, data_eq_eqNumberImpl, data_eq_eqStringImpl, data_euclideanRing_intDiv, data_euclideanRing_numDiv, data_heytingAlgebra_boolConj, data_heytingAlgebra_boolDisj, data_heytingAlgebra_boolNot, data_int_bits_and, data_int_bits_complement, data_int_bits_or, data_int_bits_shl, data_int_bits_shr, data_int_bits_xor, data_int_bits_zshr, data_ord_ordBoolean, data_ord_ordChar, data_ord_ordInt, data_ord_ordIntImpl, data_ord_ordNumber, data_ord_ordString, data_ring_intSub, data_ring_numSub, data_semigroup_concatArray, data_semigroup_concatString, data_semiring_intAdd, data_semiring_intMul, data_semiring_numAdd, data_semiring_numMul, data_string_codePoints_toCodePointArray, effect_bindE, effect_pureE, effect_unsafe_unsafePerformEffect, partial_unsafe_unsafePartial, record_builder_copyRecord, record_builder_unsafeDelete, record_builder_unsafeInsert, record_builder_unsafeModify, record_builder_unsafeRename, record_unsafe_union_unsafeUnionFn, record_unsafe_unsafeDelete, record_unsafe_unsafeGet, record_unsafe_unsafeHas, record_unsafe_unsafeSet, unsafe_coerce_unsafeCoerce ])(Data_Semigroup.append(Data_Semigroup.semigroupArray)(Data_Functor.map(Data_Functor.functorArray)(data_function_uncurried_mkFn)(oneToTen))(Data_Semigroup.append(Data_Semigroup.semigroupArray)(Data_Functor.map(Data_Functor.functorArray)(data_function_uncurried_runFn)(oneToTen))(Data_Semigroup.append(Data_Semigroup.semigroupArray)(Data_Functor.map(Data_Functor.functorArray)(effect_uncurried_mkEffectFn)(oneToTen))(Data_Semigroup.append(Data_Semigroup.semigroupArray)(Data_Functor.map(Data_Functor.functorArray)(effect_uncurried_runEffectFn)(oneToTen))(Data_Semigroup.append(Data_Semigroup.semigroupArray)(Data_Functor.map(Data_Functor.functorArray)(control_monad_st_uncurried_mkSTFn)(oneToTen))(Data_Functor.map(Data_Functor.functorArray)(control_monad_st_uncurried_runSTFn)(oneToTen)))))));
    return Data_Map_Internal.fromFoldable(PureScript_Backend_Optimizer_CoreFn.ordQualified(PureScript_Backend_Optimizer_CoreFn.ordIdent))(Data_Foldable.foldableArray)(semantics);
})();
export {
    qualified,
    coreForeignSemantics,
    effect_bindE,
    effect_pureE,
    effect_ref_new,
    effect_ref_read,
    effect_ref_write,
    effect_ref_modify,
    effect_unsafe_unsafePerformEffect,
    control_monad_st_internal_bind,
    control_monad_st_internal_map,
    control_monad_st_internal_pure,
    control_monad_st_internal_new,
    control_monad_st_internal_read,
    control_monad_st_internal_run,
    control_monad_st_internal_write,
    control_monad_st_internal_modify,
    data_array_unsafeIndexImpl,
    data_array_length,
    data_eq_eqBooleanImpl,
    data_eq_eqIntImpl,
    data_eq_eqNumberImpl,
    data_eq_eqCharImpl,
    data_eq_eqStringImpl,
    data_ord_ordBoolean,
    data_ord_ordInt,
    data_ord_ordIntImpl,
    data_ord_ordNumber,
    data_ord_ordChar,
    data_ord_ordString,
    data_semiring_intAdd,
    data_semiring_intMul,
    data_semiring_numAdd,
    data_semiring_numMul,
    data_ring_intSub,
    data_ring_numSub,
    data_euclideanRing_intDiv,
    data_euclideanRing_numDiv,
    data_int_bits_and,
    data_int_bits_complement,
    data_int_bits_or,
    data_int_bits_shl,
    data_int_bits_shr,
    data_int_bits_xor,
    data_int_bits_zshr,
    data_function_uncurried_mkFn,
    data_function_uncurried_runFn,
    mkEffectFn,
    effect_uncurried_mkEffectFn,
    control_monad_st_uncurried_mkSTFn,
    runEffectFn,
    effect_uncurried_runEffectFn,
    control_monad_st_uncurried_runSTFn,
    data_heytingAlgebra_boolConj,
    data_heytingAlgebra_boolDisj,
    data_heytingAlgebra_boolNot,
    data_string_codePoints_toCodePointArray,
    unsafe_coerce_unsafeCoerce,
    primBinaryOperator,
    primUnaryOperator,
    primOrdOperator,
    primOrdImplOperator,
    effectBind,
    effectMap,
    effectPure,
    effectRefNew,
    effectRefRead,
    effectRefWrite,
    effectRefModify,
    effectUnsafePerform,
    isQualified,
    data_semigroup_concatArray,
    data_semigroup_concatString,
    partial_unsafe_unsafePartial,
    record_builder_copyRecord,
    record_builder_unsafeInsert,
    record_builder_unsafeDelete,
    record_builder_unsafeModify,
    viewCopyRecord,
    record_builder_unsafeRename,
    record_unsafe_unsafeGet,
    record_unsafe_unsafeHas,
    record_unsafe_unsafeSet,
    record_unsafe_unsafeDelete,
    record_unsafe_union_unsafeUnionFn,
    test_data_undefinedOr_compareUndefinedOrImpl
};
