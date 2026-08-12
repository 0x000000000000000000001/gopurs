// @inline Data.Argonaut.Core.caseJson always
import * as Control_Alternative from "../Control.Alternative/index.js";
import * as Control_Applicative from "../Control.Applicative/index.js";
import * as Control_Bind from "../Control.Bind/index.js";
import * as Control_Monad_Error_Class from "../Control.Monad.Error.Class/index.js";
import * as Data_Argonaut_Core from "../Data.Argonaut.Core/index.js";
import * as Data_Argonaut_Decode_Class from "../Data.Argonaut.Decode.Class/index.js";
import * as Data_Argonaut_Decode_Error from "../Data.Argonaut.Decode.Error/index.js";
import * as Data_Array from "../Data.Array/index.js";
import * as Data_Array_ST from "../Data.Array.ST/index.js";
import * as Data_Boolean from "../Data.Boolean/index.js";
import * as Data_Bounded from "../Data.Bounded/index.js";
import * as Data_Either from "../Data.Either/index.js";
import * as Data_Enum from "../Data.Enum/index.js";
import * as Data_Foldable from "../Data.Foldable/index.js";
import * as Data_Functor from "../Data.Functor/index.js";
import * as Data_Int from "../Data.Int/index.js";
import * as Data_Map_Internal from "../Data.Map.Internal/index.js";
import * as Data_Maybe from "../Data.Maybe/index.js";
import * as Data_Monoid from "../Data.Monoid/index.js";
import * as Data_String_CodePoints from "../Data.String.CodePoints/index.js";
import * as Data_String_CodeUnits from "../Data.String.CodeUnits/index.js";
import * as Data_Traversable from "../Data.Traversable/index.js";
import * as Data_Tuple from "../Data.Tuple/index.js";
import * as Data_Unit from "../Data.Unit/index.js";
import * as Foreign_Object from "../Foreign.Object/index.js";
import * as PureScript_Backend_Optimizer_CoreFn from "../PureScript.Backend.Optimizer.CoreFn/index.js";
import * as Safe_Coerce from "../Safe.Coerce/index.js";
var throwError = /* #__PURE__ */ Control_Monad_Error_Class.throwError(Control_Monad_Error_Class.monadThrowEither);
var pure = /* #__PURE__ */ Control_Applicative.pure(Data_Either.applicativeEither);
var pure1 = /* #__PURE__ */ Control_Applicative.pure(Data_Either.applicativeEither);
var throwError1 = /* #__PURE__ */ Control_Monad_Error_Class.throwError(Control_Monad_Error_Class.monadThrowEither);
var pure2 = /* #__PURE__ */ Control_Applicative.pure(Data_Either.applicativeEither);
var pure3 = /* #__PURE__ */ Control_Applicative.pure(Data_Either.applicativeEither);
var decodeJsonTuple = /* #__PURE__ */ Data_Argonaut_Decode_Class.decodeJsonTuple(Data_Argonaut_Decode_Class.decodeJsonInt)(Data_Argonaut_Decode_Class.decodeJsonInt);
var throwError2 = /* #__PURE__ */ Control_Monad_Error_Class.throwError(Control_Monad_Error_Class.monadThrowEither);
var pure4 = /* #__PURE__ */ Control_Applicative.pure(Data_Either.applicativeEither);
var pure5 = /* #__PURE__ */ Control_Applicative.pure(Data_Either.applicativeEither);
var throwError3 = /* #__PURE__ */ Control_Monad_Error_Class.throwError(Control_Monad_Error_Class.monadThrowEither);
var pure6 = /* #__PURE__ */ Control_Applicative.pure(Data_Either.applicativeEither);
var throwError4 = /* #__PURE__ */ Control_Monad_Error_Class.throwError(Control_Monad_Error_Class.monadThrowEither);
var pure7 = /* #__PURE__ */ Control_Applicative.pure(Data_Either.applicativeEither);
var pure8 = /* #__PURE__ */ Control_Applicative.pure(Data_Either.applicativeEither);
var throwError5 = /* #__PURE__ */ Control_Monad_Error_Class.throwError(Control_Monad_Error_Class.monadThrowEither);
var pure9 = /* #__PURE__ */ Control_Applicative.pure(Data_Either.applicativeEither);
var pure10 = /* #__PURE__ */ Control_Applicative.pure(Data_Either.applicativeEither);
var throwError6 = /* #__PURE__ */ Control_Monad_Error_Class.throwError(Control_Monad_Error_Class.monadThrowEither);
var pure11 = /* #__PURE__ */ Control_Applicative.pure(Data_Either.applicativeEither);
var getFieldOptional$prime = function (decode) {
    return function (obj) {
        return function (prop) {
            var v = Foreign_Object.lookup(prop)(obj);
            if (v instanceof Data_Maybe.Nothing) {
                return new Data_Either.Right(Data_Maybe.Nothing.value);
            };
            if (v instanceof Data_Maybe.Just) {
                if (Data_Argonaut_Core.isNull(v.value0)) {
                    return new Data_Either.Right(Data_Maybe.Nothing.value);
                };
                if (Data_Boolean.otherwise) {
                    return Data_Functor.map(Data_Either.functorEither)(Data_Maybe.Just.create)(decode(v.value0));
                };
            };
            throw new Error("Failed pattern match at PureScript.Backend.Optimizer.CoreFn.Json (line 490, column 3 - line 497, column 31): " + [ v.constructor.name ]);
        };
    };
};
var getField = function (decode) {
    return function (obj) {
        return function (prop) {
            var v = Foreign_Object.lookup(prop)(obj);
            if (v instanceof Data_Maybe.Nothing) {
                return new Data_Either.Left(new Data_Argonaut_Decode_Error.AtKey(prop, Data_Argonaut_Decode_Error.MissingValue.value));
            };
            if (v instanceof Data_Maybe.Just) {
                return decode(v.value0);
            };
            throw new Error("Failed pattern match at PureScript.Backend.Optimizer.CoreFn.Json (line 482, column 3 - line 486, column 18): " + [ v.constructor.name ]);
        };
    };
};
var decodeString = /* #__PURE__ */ (function () {
    var fail = function (v) {
        return new Data_Either.Left(new Data_Argonaut_Decode_Error.TypeMismatch("String"));
    };
    return Data_Argonaut_Core.caseJson(fail)(fail)(fail)(Data_Either.Right.create)(fail)(fail);
})();
var decodeProperName = /* #__PURE__ */ Safe_Coerce.coerce()(decodeString);
var decodeNumber = /* #__PURE__ */ (function () {
    var fail = function (v) {
        return new Data_Either.Left(new Data_Argonaut_Decode_Error.TypeMismatch("Number"));
    };
    return Data_Argonaut_Core.caseJson(fail)(fail)(Data_Either.Right.create)(fail)(fail)(fail);
})();
var decodeJObject = /* #__PURE__ */ (function () {
    var fail = function (v) {
        return new Data_Either.Left(new Data_Argonaut_Decode_Error.TypeMismatch("Object"));
    };
    return Data_Argonaut_Core.caseJson(fail)(fail)(fail)(fail)(fail)(Data_Either.Right.create);
})();
var decodeJArray = /* #__PURE__ */ (function () {
    var fail = function (v) {
        return new Data_Either.Left(new Data_Argonaut_Decode_Error.TypeMismatch("Array"));
    };
    return Data_Argonaut_Core.caseJson(fail)(fail)(fail)(fail)(Data_Either.Right.create)(fail);
})();
var decodeIdent = /* #__PURE__ */ Safe_Coerce.coerce()(decodeString);
var decodeBoolean = /* #__PURE__ */ (function () {
    var fail = function (v) {
        return new Data_Either.Left(new Data_Argonaut_Decode_Error.TypeMismatch("Boolean"));
    };
    return Data_Argonaut_Core.caseJson(fail)(Data_Either.Right.create)(fail)(fail)(fail)(fail);
})();
var decodeArray = function (decoder) {
    return function (json) {
        var v = decodeJArray(json);
        if (v instanceof Data_Either.Left) {
            return new Data_Either.Left(v.value0);
        };
        if (v instanceof Data_Either.Right) {
            return (function __do() {
                var out = Data_Array_ST["new"]();
                var ix = 0;
                var con = true;
                var res = Data_Unit.unit;
                var len = Data_Array.length(v.value0);
                (function () {
                    while (con) {
                        (function __do() {
                            var ix$prime = ix;
                            var $60 = ix$prime === len;
                            if ($60) {
                                var out$prime = Data_Array_ST.unsafeFreeze(out)();
                                con = false;
                                res = new Data_Either.Right(out$prime);
                                return Data_Unit.unit;
                            };
                            var v1 = decoder(v["value0"][ix$prime]);
                            if (v1 instanceof Data_Either.Left) {
                                con = false;
                                res = new Data_Either.Left(new Data_Argonaut_Decode_Error.AtIndex(ix$prime, v1.value0));
                                return Data_Unit.unit;
                            };
                            if (v1 instanceof Data_Either.Right) {
                                Data_Array_ST.push(v1.value0)(out)();
                                ix = ix$prime + 1 | 0;
                                return Data_Unit.unit;
                            };
                            throw new Error("Failed pattern match at PureScript.Backend.Optimizer.CoreFn.Json (line 469, column 9 - line 477, column 22): " + [ v1.constructor.name ]);
                        })();
                    };
                    return {};
                })();
                return res;
            })();
        };
        throw new Error("Failed pattern match at PureScript.Backend.Optimizer.CoreFn.Json (line 452, column 28 - line 478, column 19): " + [ v.constructor.name ]);
    };
};
var decodeModuleName = /* #__PURE__ */ (function () {
    var $110 = Data_Functor.map(Data_Either.functorEither)((function () {
        var $113 = Data_Foldable.intercalate(Data_Foldable.foldableArray)(Data_Monoid.monoidString)(".");
        return function ($114) {
            return PureScript_Backend_Optimizer_CoreFn.ModuleName($113($114));
        };
    })());
    var $111 = decodeArray(decodeString);
    return function ($112) {
        return $110($111($112));
    };
})();

// Either's bind implementation is not ideal from an optimization
// standpoint and generates awkward code.
var bind = function (a) {
    return function (k) {
        if (a instanceof Data_Either.Left) {
            return new Data_Either.Left(a.value0);
        };
        if (a instanceof Data_Either.Right) {
            return k(a.value0);
        };
        throw new Error("Failed pattern match at PureScript.Backend.Optimizer.CoreFn.Json (line 47, column 12 - line 51, column 9): " + [ a.constructor.name ]);
    };
};
var decodeConstructorType = function (json) {
    return bind(decodeString(json))(function (str) {
        if (str === "ProductType") {
            return Control_Applicative.pure(Data_Either.applicativeEither)(PureScript_Backend_Optimizer_CoreFn.ProductType.value);
        };
        if (str === "SumType") {
            return Control_Applicative.pure(Data_Either.applicativeEither)(PureScript_Backend_Optimizer_CoreFn.SumType.value);
        };
        return throwError(new Data_Argonaut_Decode_Error.TypeMismatch("ConstructorType"));
    });
};
var decodeImport = function (decodeAnn$prime) {
    return function (json) {
        return bind(decodeJObject(json))(function (obj) {
            return bind(getField(decodeAnn$prime)(obj)("annotation"))(function (ann) {
                return bind(getField(decodeModuleName)(obj)("moduleName"))(function (mod) {
                    return pure(new PureScript_Backend_Optimizer_CoreFn.Import(ann, mod));
                });
            });
        });
    };
};
var decodeInt = function (json) {
    return bind(decodeNumber(json))(function (num) {
        var v = Data_Int.fromNumber(num);
        if (v instanceof Data_Maybe.Nothing) {
            var $70 = num === 2.147483648e9;
            if ($70) {
                return new Data_Either.Right(Data_Bounded.bottom(Data_Bounded.boundedInt));
            };
            return new Data_Either.Left(new Data_Argonaut_Decode_Error.TypeMismatch("Int"));
        };
        if (v instanceof Data_Maybe.Just) {
            return new Data_Either.Right(v.value0);
        };
        throw new Error("Failed pattern match at PureScript.Backend.Optimizer.CoreFn.Json (line 542, column 3 - line 549, column 16): " + [ v.constructor.name ]);
    });
};
var decodeCodePoint = /* #__PURE__ */ Control_Bind.composeKleisliFlipped(Data_Either.bindEither)(/* #__PURE__ */ (function () {
    var $115 = Data_Either.note(new Data_Argonaut_Decode_Error.TypeMismatch("CodePoint"));
    var $116 = Data_Enum.toEnum(Data_String_CodePoints.boundedEnumCodePoint);
    return function ($117) {
        return $115($116($117));
    };
})())(decodeInt);
var decodeCodePointArray = /* #__PURE__ */ decodeArray(decodeCodePoint);
var decodeMeta = function (json) {
    return bind(decodeJObject(json))(function (obj) {
        return bind(getField(decodeString)(obj)("metaType"))(function (typ) {
            if (typ === "IsConstructor") {
                return bind(getField(decodeConstructorType)(obj)("constructorType"))(function (ct) {
                    return bind(getField(decodeArray(decodeIdent))(obj)("identifiers"))(function (is) {
                        return pure1(new PureScript_Backend_Optimizer_CoreFn.IsConstructor(ct, is));
                    });
                });
            };
            if (typ === "IsNewtype") {
                return Control_Applicative.pure(Data_Either.applicativeEither)(PureScript_Backend_Optimizer_CoreFn.IsNewtype.value);
            };
            if (typ === "IsTypeClassConstructor") {
                return Control_Applicative.pure(Data_Either.applicativeEither)(PureScript_Backend_Optimizer_CoreFn.IsTypeClassConstructor.value);
            };
            if (typ === "IsForeign") {
                return Control_Applicative.pure(Data_Either.applicativeEither)(PureScript_Backend_Optimizer_CoreFn.IsForeign.value);
            };
            if (typ === "IsWhere") {
                return Control_Applicative.pure(Data_Either.applicativeEither)(PureScript_Backend_Optimizer_CoreFn.IsWhere.value);
            };
            if (typ === "IsSyntheticApp") {
                return Control_Applicative.pure(Data_Either.applicativeEither)(PureScript_Backend_Optimizer_CoreFn.IsSyntheticApp.value);
            };
            return throwError1(new Data_Argonaut_Decode_Error.TypeMismatch("Meta"));
        });
    });
};
var decodeQualified = function (k) {
    return function (json) {
        return bind(decodeJObject(json))(function (obj) {
            return bind(getFieldOptional$prime(decodeModuleName)(obj)("moduleName"))(function (moduleName) {
                return bind(getField(k)(obj)("identifier"))(function (identifier) {
                    return pure2(new PureScript_Backend_Optimizer_CoreFn.Qualified(moduleName, identifier));
                });
            });
        });
    };
};
var decodeReExports = function (json) {
    return bind(decodeJObject(json))(function (obj) {
        return bind(Data_Traversable.traverse(Data_Traversable.traversableArray)(Data_Either.applicativeEither)(Data_Traversable.traverse(Data_Traversable.traversableTuple)(Data_Either.applicativeEither)(decodeArray(decodeIdent)))(Foreign_Object.toArrayWithKey(Data_Tuple.Tuple.create)(obj)))(function (all) {
            return pure3(Control_Bind.bind(Control_Bind.bindArray)(all)(function (v) {
                return Data_Functor.map(Data_Functor.functorArray)(PureScript_Backend_Optimizer_CoreFn.ReExport.create(v.value0))(v.value1);
            }));
        });
    });
};
var decodeSourcePos = function (json) {
    return bind(Data_Argonaut_Decode_Class.decodeJson(decodeJsonTuple)(json))(function (v) {
        return Control_Applicative.pure(Data_Either.applicativeEither)({
            line: v.value0,
            column: v.value1
        });
    });
};
var decodeSourceSpan = function (path) {
    return function (json) {
        return bind(decodeJObject(json))(function (obj) {
            return bind(getField(decodeSourcePos)(obj)("start"))(function (start) {
                return bind(getField(decodeSourcePos)(obj)("end"))(function (end) {
                    return Control_Applicative.pure(Data_Either.applicativeEither)({
                        path: path,
                        start: start,
                        end: end
                    });
                });
            });
        });
    };
};
var alt = function (a) {
    return function (k) {
        if (a instanceof Data_Either.Left) {
            return k(Data_Unit.unit);
        };
        if (a instanceof Data_Either.Right) {
            return a;
        };
        throw new Error("Failed pattern match at PureScript.Backend.Optimizer.CoreFn.Json (line 40, column 11 - line 42, column 15): " + [ a.constructor.name ]);
    };
};
var decodeComment = function (json) {
    return bind(decodeJObject(json))(function (obj) {
        return alt(Data_Functor.map(Data_Either.functorEither)(PureScript_Backend_Optimizer_CoreFn.LineComment.create)(getField(decodeString)(obj)("LineComment")))(function (v) {
            return Data_Functor.map(Data_Either.functorEither)(PureScript_Backend_Optimizer_CoreFn.BlockComment.create)(getField(decodeString)(obj)("BlockComment"));
        });
    });
};
var decodeField = function (j) {
    return bind(decodeJObject(j))(function (o) {
        return bind(getField(decodeString)(o)("label"))(function (l) {
            return bind(getField(decodeExprType)(o)("type"))(function (t) {
                return Control_Applicative.pure(Data_Either.applicativeEither)(new Data_Tuple.Tuple(l, t));
            });
        });
    });
};
var decodeExprType = function (json) {
    var decodeStr = bind(decodeString(json))(function (str) {
        if (str === "Int") {
            return Control_Applicative.pure(Data_Either.applicativeEither)(PureScript_Backend_Optimizer_CoreFn.Int.value);
        };
        if (str === "Number") {
            return Control_Applicative.pure(Data_Either.applicativeEither)(PureScript_Backend_Optimizer_CoreFn["Number"].value);
        };
        if (str === "String") {
            return Control_Applicative.pure(Data_Either.applicativeEither)(PureScript_Backend_Optimizer_CoreFn["String"].value);
        };
        if (str === "Char") {
            return Control_Applicative.pure(Data_Either.applicativeEither)(PureScript_Backend_Optimizer_CoreFn.Char.value);
        };
        if (str === "Boolean") {
            return Control_Applicative.pure(Data_Either.applicativeEither)(PureScript_Backend_Optimizer_CoreFn["Boolean"].value);
        };
        if (str === "Unit") {
            return Control_Applicative.pure(Data_Either.applicativeEither)(PureScript_Backend_Optimizer_CoreFn.Unit.value);
        };
        if (str === "Any") {
            return Control_Applicative.pure(Data_Either.applicativeEither)(PureScript_Backend_Optimizer_CoreFn.Any.value);
        };
        return throwError2(new Data_Argonaut_Decode_Error.TypeMismatch("ExprType"));
    });
    var decodeObj = function (v) {
        return bind(decodeJObject(json))(function (obj) {
            var v1 = Foreign_Object.lookup("Func")(obj);
            if (v1 instanceof Data_Maybe.Just) {
                return bind(decodeJObject(v1.value0))(function (funcObj) {
                    return bind(getField(decodeArray(decodeExprType))(funcObj)("args"))(function (args) {
                        return bind(getField(decodeExprType)(funcObj)("ret"))(function (ret) {
                            return Control_Applicative.pure(Data_Either.applicativeEither)(new PureScript_Backend_Optimizer_CoreFn.Func(args, ret));
                        });
                    });
                });
            };
            if (v1 instanceof Data_Maybe.Nothing) {
                var v2 = Foreign_Object.lookup("Record")(obj);
                if (v2 instanceof Data_Maybe.Just) {
                    return bind(decodeJObject(v2.value0))(function (recordObj) {
                        var entries = Foreign_Object.toArrayWithKey(Data_Tuple.Tuple.create)(recordObj);
                        return bind(Data_Traversable.traverse(Data_Traversable.traversableArray)(Data_Either.applicativeEither)(function (v3) {
                            return Data_Functor.map(Data_Either.functorEither)(Data_Tuple.Tuple.create(v3.value0))(decodeExprType(v3.value1));
                        })(entries))(function (parsedEntries) {
                            return Control_Applicative.pure(Data_Either.applicativeEither)(new PureScript_Backend_Optimizer_CoreFn.Record(new PureScript_Backend_Optimizer_CoreFn.Row(parsedEntries, Data_Maybe.Nothing.value)));
                        });
                    });
                };
                if (v2 instanceof Data_Maybe.Nothing) {
                    var v3 = Foreign_Object.lookup("Array")(obj);
                    if (v3 instanceof Data_Maybe.Just) {
                        return bind(decodeExprType(v3.value0))(function (innerType) {
                            return Control_Applicative.pure(Data_Either.applicativeEither)(new PureScript_Backend_Optimizer_CoreFn["Array"](innerType));
                        });
                    };
                    if (v3 instanceof Data_Maybe.Nothing) {
                        var v4 = Foreign_Object.lookup("ADT")(obj);
                        if (v4 instanceof Data_Maybe.Just) {
                            return bind(decodeJObject(v4.value0))(function (adtObj) {
                                return bind(getField(decodeArray(decodeString))(adtObj)("path"))(function (adtPath) {
                                    return bind(getField(decodeArray(decodeExprType))(adtObj)("args"))(function (adtArgs) {
                                        return Control_Applicative.pure(Data_Either.applicativeEither)(new PureScript_Backend_Optimizer_CoreFn.ADT(Data_Foldable.intercalate(Data_Foldable.foldableArray)(Data_Monoid.monoidString)(".")(adtPath), adtPath, adtArgs));
                                    });
                                });
                            });
                        };
                        if (v4 instanceof Data_Maybe.Nothing) {
                            var v5 = Foreign_Object.lookup("TypeVar")(obj);
                            if (v5 instanceof Data_Maybe.Just) {
                                return bind(decodeString(v5.value0))(function (tv) {
                                    return Control_Applicative.pure(Data_Either.applicativeEither)(new PureScript_Backend_Optimizer_CoreFn.TypeVar(tv));
                                });
                            };
                            if (v5 instanceof Data_Maybe.Nothing) {
                                return bind(getField(decodeString)(obj)("type"))(function (typ) {
                                    if (typ === "Int") {
                                        return Control_Applicative.pure(Data_Either.applicativeEither)(PureScript_Backend_Optimizer_CoreFn.Int.value);
                                    };
                                    if (typ === "Number") {
                                        return Control_Applicative.pure(Data_Either.applicativeEither)(PureScript_Backend_Optimizer_CoreFn["Number"].value);
                                    };
                                    if (typ === "String") {
                                        return Control_Applicative.pure(Data_Either.applicativeEither)(PureScript_Backend_Optimizer_CoreFn["String"].value);
                                    };
                                    if (typ === "Char") {
                                        return Control_Applicative.pure(Data_Either.applicativeEither)(PureScript_Backend_Optimizer_CoreFn.Char.value);
                                    };
                                    if (typ === "Boolean") {
                                        return Control_Applicative.pure(Data_Either.applicativeEither)(PureScript_Backend_Optimizer_CoreFn["Boolean"].value);
                                    };
                                    if (typ === "Unit") {
                                        return Control_Applicative.pure(Data_Either.applicativeEither)(PureScript_Backend_Optimizer_CoreFn.Unit.value);
                                    };
                                    if (typ === "Any") {
                                        return Control_Applicative.pure(Data_Either.applicativeEither)(PureScript_Backend_Optimizer_CoreFn.Any.value);
                                    };
                                    if (typ === "TypeLevelString") {
                                        return Data_Functor.map(Data_Either.functorEither)(PureScript_Backend_Optimizer_CoreFn.TypeLevelString.create)(getField(decodeString)(obj)("value"));
                                    };
                                    if (typ === "Array") {
                                        return Data_Functor.map(Data_Either.functorEither)(PureScript_Backend_Optimizer_CoreFn["Array"].create)(getField(decodeExprType)(obj)("element"));
                                    };
                                    if (typ === "TypeVar") {
                                        return Data_Functor.map(Data_Either.functorEither)(PureScript_Backend_Optimizer_CoreFn.TypeVar.create)(getField(decodeString)(obj)("name"));
                                    };
                                    if (typ === "Adt") {
                                        return bind(getField(decodeArray(decodeString))(obj)("fqn"))(function (fqn) {
                                            return bind(getField(decodeArray(decodeExprType))(obj)("args"))(function (args) {
                                                return Control_Applicative.pure(Data_Either.applicativeEither)(new PureScript_Backend_Optimizer_CoreFn.ADT(Data_Foldable.intercalate(Data_Foldable.foldableArray)(Data_Monoid.monoidString)(".")(fqn), fqn, args));
                                            });
                                        });
                                    };
                                    if (typ === "TypeApp") {
                                        return bind(getField(decodeExprType)(obj)("constructor"))(function (constructor) {
                                            return bind(getField(decodeArray(decodeExprType))(obj)("args"))(function (args) {
                                                return Control_Applicative.pure(Data_Either.applicativeEither)(new PureScript_Backend_Optimizer_CoreFn.TypeApp(constructor, args));
                                            });
                                        });
                                    };
                                    if (typ === "Func") {
                                        return bind(getField(decodeArray(decodeExprType))(obj)("args"))(function (args) {
                                            return bind(getField(decodeExprType)(obj)("ret"))(function (ret) {
                                                return Control_Applicative.pure(Data_Either.applicativeEither)(new PureScript_Backend_Optimizer_CoreFn.Func(args, ret));
                                            });
                                        });
                                    };
                                    if (typ === "Row") {
                                        return bind(getField(decodeArray(decodeField))(obj)("fields"))(function (fields) {
                                            return bind(getFieldOptional$prime(decodeExprType)(obj)("tail"))(function (tail) {
                                                return Control_Applicative.pure(Data_Either.applicativeEither)(new PureScript_Backend_Optimizer_CoreFn.Row(fields, tail));
                                            });
                                        });
                                    };
                                    if (typ === "Record") {
                                        return Data_Functor.map(Data_Either.functorEither)(PureScript_Backend_Optimizer_CoreFn.Record.create)(getField(decodeExprType)(obj)("row"));
                                    };
                                    if (typ === "ForAll") {
                                        return bind(getField(decodeArray(decodeString))(obj)("vars"))(function (vars) {
                                            return bind(getField(decodeExprType)(obj)("body"))(function (body) {
                                                return Control_Applicative.pure(Data_Either.applicativeEither)(new PureScript_Backend_Optimizer_CoreFn.ForAll(vars, body));
                                            });
                                        });
                                    };
                                    if (typ === "ConstrainedType") {
                                        return bind(getField(decodeArray(decodeConstraint))(obj)("constraints"))(function (constraints) {
                                            return bind(getField(decodeExprType)(obj)("body"))(function (body) {
                                                return Control_Applicative.pure(Data_Either.applicativeEither)(new PureScript_Backend_Optimizer_CoreFn.ConstrainedType(constraints, body));
                                            });
                                        });
                                    };
                                    return throwError2(new Data_Argonaut_Decode_Error.TypeMismatch("ExprType"));
                                });
                            };
                            throw new Error("Failed pattern match at PureScript.Backend.Optimizer.CoreFn.Json (line 150, column 24 - line 193, column 60): " + [ v5.constructor.name ]);
                        };
                        throw new Error("Failed pattern match at PureScript.Backend.Optimizer.CoreFn.Json (line 144, column 22 - line 193, column 60): " + [ v4.constructor.name ]);
                    };
                    throw new Error("Failed pattern match at PureScript.Backend.Optimizer.CoreFn.Json (line 140, column 20 - line 193, column 60): " + [ v3.constructor.name ]);
                };
                throw new Error("Failed pattern match at PureScript.Backend.Optimizer.CoreFn.Json (line 134, column 18 - line 193, column 60): " + [ v2.constructor.name ]);
            };
            throw new Error("Failed pattern match at PureScript.Backend.Optimizer.CoreFn.Json (line 128, column 5 - line 193, column 60): " + [ v1.constructor.name ]);
        });
    };
    return alt(decodeStr)(decodeObj);
};
var decodeConstraint = function (j) {
    return bind(decodeJObject(j))(function (o) {
        return bind(getField(decodeArray(decodeString))(o)("fqn"))(function (fqn) {
            return bind(getField(decodeArray(decodeExprType))(o)("args"))(function (args) {
                return Control_Applicative.pure(Data_Either.applicativeEither)(new Data_Tuple.Tuple(fqn, args));
            });
        });
    });
};
var decodeAnn = function (_path) {
    return function (json) {
        return bind(decodeJObject(json))(function (obj) {
            return bind(getFieldOptional$prime(decodeMeta)(obj)("meta"))(function (meta) {
                return bind(getFieldOptional$prime(decodeExprType)(obj)("type"))(function (type_) {
                    return pure4({
                        span: PureScript_Backend_Optimizer_CoreFn.emptySpan,
                        meta: meta,
                        type: type_
                    });
                });
            });
        });
    };
};
var decodeDataConstructor = function (json) {
    return bind(decodeJObject(json))(function (obj) {
        return bind(alt(getField(decodeString)(obj)("name"))(function (v) {
            return getField(decodeString)(obj)("constructorName");
        }))(function (name) {
            return bind(alt(getField(decodeArray(decodeExprType))(obj)("fields"))(function (v) {
                return getField(decodeArray(decodeExprType))(obj)("fieldTypes");
            }))(function (fields) {
                return Control_Applicative.pure(Data_Either.applicativeEither)({
                    name: name,
                    fields: fields
                });
            });
        });
    });
};
var decodeDataDecl = function (json) {
    return bind(decodeJObject(json))(function (obj) {
        return bind(alt(getField(decodeString)(obj)("name"))(function (v) {
            return getField(decodeString)(obj)("typeName");
        }))(function (name) {
            return bind(alt(getFieldOptional$prime(decodeArray(decodeString))(obj)("vars"))(function (v) {
                return getFieldOptional$prime(decodeArray(decodeString))(obj)("typeVars");
            }))(function (mbTypeVars) {
                var vars = Data_Maybe.fromMaybe([  ])(mbTypeVars);
                return bind(getField(decodeArray(decodeDataConstructor))(obj)("constructors"))(function (constructors) {
                    return Control_Applicative.pure(Data_Either.applicativeEither)({
                        name: name,
                        vars: vars,
                        constructors: constructors
                    });
                });
            });
        });
    });
};
var decodeMethod = function (j) {
    return bind(decodeJObject(j))(function (o) {
        return bind(getField(decodeString)(o)("name"))(function (name) {
            return bind(getField(decodeExprType)(o)("type"))(function (t) {
                return Control_Applicative.pure(Data_Either.applicativeEither)(new Data_Tuple.Tuple(name, t));
            });
        });
    });
};
var decodeClassDecl = function (json) {
    return bind(decodeJObject(json))(function (obj) {
        return bind(getField(decodeString)(obj)("name"))(function (name) {
            return bind(getFieldOptional$prime(decodeArray(decodeString))(obj)("vars"))(function (mbVars) {
                var vars = Data_Maybe.fromMaybe([  ])(mbVars);
                return bind(getField(decodeArray(decodeConstraint))(obj)("superclasses"))(function (superclasses) {
                    return bind(getField(decodeArray(decodeMethod))(obj)("methods"))(function (methods) {
                        return Control_Applicative.pure(Data_Either.applicativeEither)({
                            name: name,
                            vars: vars,
                            superclasses: superclasses,
                            methods: methods
                        });
                    });
                });
            });
        });
    });
};
var decodeStringLiteral = function (json) {
    return alt(decodeString(json))(function (v) {
        return alt(Data_Functor.map(Data_Either.functorEither)(Data_String_CodePoints.fromCodePointArray)(decodeCodePointArray(json)))(function (v1) {
            return Control_Monad_Error_Class.throwError(Control_Monad_Error_Class.monadThrowEither)(new Data_Argonaut_Decode_Error.TypeMismatch("StringLiteral"));
        });
    });
};
var decodeRecord = /* #__PURE__ */ (function () {
    var decodeProp = function (decoder) {
        return function (json) {
            return bind(decodeJArray(json))(function (arr) {
                if (arr.length === 2) {
                    return bind(decodeStringLiteral(arr[0]))(function (prop) {
                        return bind(decoder(arr[1]))(function (value) {
                            return pure5(new PureScript_Backend_Optimizer_CoreFn.Prop(prop, value));
                        });
                    });
                };
                return new Data_Either.Left(new Data_Argonaut_Decode_Error.TypeMismatch("Tuple"));
            });
        };
    };
    return function ($118) {
        return decodeArray(decodeProp($118));
    };
})();
var decodeLiteral = function (dec) {
    return function (json) {
        return bind(decodeJObject(json))(function (obj) {
            return bind(getField(decodeString)(obj)("literalType"))(function (typ) {
                if (typ === "IntLiteral") {
                    return Data_Functor.map(Data_Either.functorEither)(PureScript_Backend_Optimizer_CoreFn.LitInt.create)(getField(decodeInt)(obj)("value"));
                };
                if (typ === "NumberLiteral") {
                    return Data_Functor.map(Data_Either.functorEither)(PureScript_Backend_Optimizer_CoreFn.LitNumber.create)(getField(decodeNumber)(obj)("value"));
                };
                if (typ === "StringLiteral") {
                    return Data_Functor.map(Data_Either.functorEither)(PureScript_Backend_Optimizer_CoreFn.LitString.create)(getField(decodeStringLiteral)(obj)("value"));
                };
                if (typ === "CharLiteral") {
                    return bind(getField(decodeString)(obj)("value"))(function (str) {
                        return Data_Functor.map(Data_Either.functorEither)(PureScript_Backend_Optimizer_CoreFn.LitChar.create)(Data_Either.note(new Data_Argonaut_Decode_Error.TypeMismatch("Char"))(Control_Bind.discard(Control_Bind.discardUnit)(Data_Maybe.bindMaybe)(Control_Alternative.guard(Data_Maybe.alternativeMaybe)(Data_String_CodeUnits.length(str) === 1))(function () {
                            return Data_Array.head(Data_String_CodeUnits.toCharArray(str));
                        })));
                    });
                };
                if (typ === "BooleanLiteral") {
                    return Data_Functor.map(Data_Either.functorEither)(PureScript_Backend_Optimizer_CoreFn.LitBoolean.create)(getField(decodeBoolean)(obj)("value"));
                };
                if (typ === "ArrayLiteral") {
                    return Data_Functor.map(Data_Either.functorEither)(PureScript_Backend_Optimizer_CoreFn.LitArray.create)(getField(decodeArray(dec))(obj)("value"));
                };
                if (typ === "ObjectLiteral") {
                    return Data_Functor.map(Data_Either.functorEither)(PureScript_Backend_Optimizer_CoreFn.LitRecord.create)(getField(decodeRecord(dec))(obj)("value"));
                };
                return throwError3(new Data_Argonaut_Decode_Error.TypeMismatch("Literal"));
            });
        });
    };
};
var decodeBinder = function (decAnn) {
    return function (json) {
        return bind(decodeJObject(json))(function (obj) {
            return bind(getField(decAnn)(obj)("annotation"))(function (ann) {
                return bind(getField(decodeString)(obj)("binderType"))(function (typ) {
                    if (typ === "NullBinder") {
                        return pure6(new PureScript_Backend_Optimizer_CoreFn.BinderNull(ann));
                    };
                    if (typ === "VarBinder") {
                        return Data_Functor.map(Data_Either.functorEither)(PureScript_Backend_Optimizer_CoreFn.BinderVar.create(ann))(getField(decodeIdent)(obj)("identifier"));
                    };
                    if (typ === "LiteralBinder") {
                        return Data_Functor.map(Data_Either.functorEither)(PureScript_Backend_Optimizer_CoreFn.BinderLit.create(ann))(getField(decodeLiteral(decodeBinder(decAnn)))(obj)("literal"));
                    };
                    if (typ === "ConstructorBinder") {
                        return bind(getField(decodeQualified(decodeProperName))(obj)("typeName"))(function (tyn) {
                            return bind(alt(getField(decodeQualified(decodeIdent))(obj)("name"))(function (v) {
                                return getField(decodeQualified(decodeIdent))(obj)("constructorName");
                            }))(function (ctn) {
                                return bind(getField(decodeArray(decodeBinder(decAnn)))(obj)("binders"))(function (binders) {
                                    return pure6(new PureScript_Backend_Optimizer_CoreFn.BinderConstructor(ann, tyn, ctn, binders));
                                });
                            });
                        });
                    };
                    if (typ === "NamedBinder") {
                        return bind(getField(decodeIdent)(obj)("identifier"))(function (ident) {
                            return bind(getField(decodeBinder(decAnn))(obj)("binder"))(function (binder) {
                                return pure6(new PureScript_Backend_Optimizer_CoreFn.BinderNamed(ann, ident, binder));
                            });
                        });
                    };
                    return throwError4(new Data_Argonaut_Decode_Error.TypeMismatch("Binder"));
                });
            });
        });
    };
};
var decodeGuard = function (decAnn) {
    return function (json) {
        return bind(decodeJObject(json))(function (obj) {
            return bind(getField(decodeExpr(decAnn))(obj)("guard"))(function (guard) {
                return bind(getField(decodeExpr(decAnn))(obj)("expression"))(function (expr) {
                    return pure7(new PureScript_Backend_Optimizer_CoreFn.Guard(guard, expr));
                });
            });
        });
    };
};
var decodeExpr = function (decAnn) {
    return function (json) {
        return bind(decodeJObject(json))(function (obj) {
            return bind(getField(decAnn)(obj)("annotation"))(function (ann) {
                return bind(getField(decodeString)(obj)("type"))(function (typ) {
                    if (typ === "Var") {
                        return Data_Functor.map(Data_Either.functorEither)(PureScript_Backend_Optimizer_CoreFn.ExprVar.create(ann))(getField(decodeQualified(decodeIdent))(obj)("value"));
                    };
                    if (typ === "Literal") {
                        return Data_Functor.map(Data_Either.functorEither)(PureScript_Backend_Optimizer_CoreFn.ExprLit.create(ann))(getField(decodeLiteral(decodeExpr(decAnn)))(obj)("value"));
                    };
                    if (typ === "Constructor") {
                        return bind(getField(decodeProperName)(obj)("typeName"))(function (tyn) {
                            return bind(alt(getField(decodeIdent)(obj)("name"))(function (v) {
                                return getField(decodeIdent)(obj)("constructorName");
                            }))(function (con) {
                                return bind(alt(getField(decodeArray(decodeStringLiteral))(obj)("fields"))(function (v) {
                                    return getField(decodeArray(decodeStringLiteral))(obj)("fieldNames");
                                }))(function (is) {
                                    return pure8(new PureScript_Backend_Optimizer_CoreFn.ExprConstructor(ann, tyn, con, is));
                                });
                            });
                        });
                    };
                    if (typ === "Accessor") {
                        return bind(getField(decodeExpr(decAnn))(obj)("expression"))(function (e) {
                            return bind(getField(decodeStringLiteral)(obj)("fieldName"))(function (f) {
                                return pure8(new PureScript_Backend_Optimizer_CoreFn.ExprAccessor(ann, e, f));
                            });
                        });
                    };
                    if (typ === "ObjectUpdate") {
                        return bind(getField(decodeExpr(decAnn))(obj)("expression"))(function (e) {
                            return bind(getField(decodeRecord(decodeExpr(decAnn)))(obj)("updates"))(function (us) {
                                return pure8(new PureScript_Backend_Optimizer_CoreFn.ExprUpdate(ann, e, us));
                            });
                        });
                    };
                    if (typ === "Abs") {
                        return bind(getField(decodeIdent)(obj)("argument"))(function (idn) {
                            return bind(getField(decodeExpr(decAnn))(obj)("body"))(function (e) {
                                return pure8(new PureScript_Backend_Optimizer_CoreFn.ExprAbs(ann, idn, e));
                            });
                        });
                    };
                    if (typ === "App") {
                        return bind(getField(decodeExpr(decAnn))(obj)("abstraction"))(function (e1) {
                            return bind(getField(decodeExpr(decAnn))(obj)("argument"))(function (e2) {
                                return pure8(new PureScript_Backend_Optimizer_CoreFn.ExprApp(ann, e1, e2));
                            });
                        });
                    };
                    if (typ === "Case") {
                        return bind(getField(decodeArray(decodeExpr(decAnn)))(obj)("caseExpressions"))(function (cs) {
                            return bind(getField(decodeArray(decodeCaseAlternative(decAnn)))(obj)("caseAlternatives"))(function (cas) {
                                return pure8(new PureScript_Backend_Optimizer_CoreFn.ExprCase(ann, cs, cas));
                            });
                        });
                    };
                    if (typ === "Let") {
                        return bind(getField(decodeArray(decodeBind(decAnn)))(obj)("binds"))(function (bs) {
                            return bind(getField(decodeExpr(decAnn))(obj)("expression"))(function (e) {
                                return pure8(new PureScript_Backend_Optimizer_CoreFn.ExprLet(ann, bs, e));
                            });
                        });
                    };
                    return throwError5(new Data_Argonaut_Decode_Error.TypeMismatch("Expr"));
                });
            });
        });
    };
};
var decodeCaseAlternative = function (decAnn) {
    return function (json) {
        return bind(decodeJObject(json))(function (obj) {
            return bind(getField(decodeArray(decodeBinder(decAnn)))(obj)("binders"))(function (binders) {
                return bind(getField(decodeBoolean)(obj)("isGuarded"))(function (isGuarded) {
                    if (isGuarded) {
                        return bind(getField(decodeArray(decodeGuard(decAnn)))(obj)("expressions"))(function (es) {
                            return pure9(new PureScript_Backend_Optimizer_CoreFn.CaseAlternative(binders, new PureScript_Backend_Optimizer_CoreFn.Guarded(es)));
                        });
                    };
                    return bind(getField(decodeExpr(decAnn))(obj)("expression"))(function (e) {
                        return pure9(new PureScript_Backend_Optimizer_CoreFn.CaseAlternative(binders, new PureScript_Backend_Optimizer_CoreFn.Unconditional(e)));
                    });
                });
            });
        });
    };
};
var decodeBinding = function (decAnn) {
    return function (obj) {
        return bind(getField(decAnn)(obj)("annotation"))(function (ann) {
            return bind(getField(decodeIdent)(obj)("identifier"))(function (ident) {
                return bind(getField(decodeExpr(decAnn))(obj)("expression"))(function (expr) {
                    return pure10(new PureScript_Backend_Optimizer_CoreFn.Binding(ann, ident, expr));
                });
            });
        });
    };
};
var decodeBind = function (decAnn) {
    return function (json) {
        return bind(decodeJObject(json))(function (obj) {
            return bind(getField(decodeString)(obj)("bindType"))(function (typ) {
                if (typ === "NonRec") {
                    return Data_Functor.map(Data_Either.functorEither)(PureScript_Backend_Optimizer_CoreFn.NonRec.create)(decodeBinding(decAnn)(obj));
                };
                if (typ === "Rec") {
                    return Data_Functor.map(Data_Either.functorEither)(PureScript_Backend_Optimizer_CoreFn.Rec.create)(getField(decodeArray(Control_Bind.composeKleisli(Data_Either.bindEither)(decodeJObject)(decodeBinding(decAnn))))(obj)("binds"));
                };
                return throwError6(new Data_Argonaut_Decode_Error.TypeMismatch("Bind"));
            });
        });
    };
};
var decodeModule$prime = function (decodeAnn$prime) {
    return function (json) {
        return bind(decodeJObject(json))(function (obj) {
            return bind(getField(decodeModuleName)(obj)("moduleName"))(function (name) {
                return bind(getField(decodeString)(obj)("modulePath"))(function (path) {
                    return bind(getField(decodeSourceSpan(path))(obj)("sourceSpan"))(function (span) {
                        return bind(getField(decodeArray(decodeImport(decodeAnn$prime(path))))(obj)("imports"))(function (imports) {
                            return bind(getField(decodeArray(decodeIdent))(obj)("exports"))(function (exports) {
                                return bind(getField(decodeReExports)(obj)("reExports"))(function (reExports) {
                                    return bind(getField(decodeArray(decodeDataDecl))(obj)("dataDecls"))(function (dataDecls) {
                                        return bind(getFieldOptional$prime(decodeArray(decodeClassDecl))(obj)("classDecls"))(function (mbClassDecls) {
                                            var classDecls = Data_Maybe.fromMaybe([  ])(mbClassDecls);
                                            return bind(getField(decodeArray(decodeBind(decodeAnn$prime(path))))(obj)("decls"))(function (decls) {
                                                return bind(getField(decodeArray(decodeIdent))(obj)("foreign"))(function (foreign_arr) {
                                                    return bind(Data_Functor.map(Data_Either.functorEither)(Data_Maybe.fromMaybe(Foreign_Object.empty))(getFieldOptional$prime(decodeJObject)(obj)("foreignAnnotations")))(function (foreign_anns) {
                                                        return bind(Data_Traversable.traverse(Data_Traversable.traversableArray)(Data_Either.applicativeEither)(function (v) {
                                                            var v1 = Foreign_Object.lookup(v)(foreign_anns);
                                                            if (v1 instanceof Data_Maybe.Just) {
                                                                return bind(decodeAnn(path)(v1.value0))(function (v2) {
                                                                    return Control_Applicative.pure(Data_Either.applicativeEither)(new Data_Tuple.Tuple(v, v2.type));
                                                                });
                                                            };
                                                            if (v1 instanceof Data_Maybe.Nothing) {
                                                                return Control_Applicative.pure(Data_Either.applicativeEither)(new Data_Tuple.Tuple(v, Data_Maybe.Nothing.value));
                                                            };
                                                            throw new Error("Failed pattern match at PureScript.Backend.Optimizer.CoreFn.Json (line 277, column 5 - line 282, column 35): " + [ v1.constructor.name ]);
                                                        })(foreign_arr))(function (foreign_list) {
                                                            var foreignMap = Data_Map_Internal.fromFoldable(PureScript_Backend_Optimizer_CoreFn.ordIdent)(Data_Foldable.foldableArray)(foreign_list);
                                                            return bind(getField(decodeArray(decodeComment))(obj)("comments"))(function (comments) {
                                                                return pure11({
                                                                    name: name,
                                                                    path: path,
                                                                    span: span,
                                                                    imports: imports,
                                                                    exports: exports,
                                                                    reExports: reExports,
                                                                    dataDecls: dataDecls,
                                                                    classDecls: classDecls,
                                                                    decls: decls,
                                                                    foreign: foreignMap,
                                                                    comments: comments
                                                                });
                                                            });
                                                        });
                                                    });
                                                });
                                            });
                                        });
                                    });
                                });
                            });
                        });
                    });
                });
            });
        });
    };
};
var decodeModule = /* #__PURE__ */ decodeModule$prime(decodeAnn);
export {
    decodeModule,
    decodeModule$prime,
    decodeAnn
};
