import * as Control_Applicative from "../Control.Applicative/index.js";
import * as Control_Apply from "../Control.Apply/index.js";
import * as Control_Bind from "../Control.Bind/index.js";
import * as Data_Argonaut_Core from "../Data.Argonaut.Core/index.js";
import * as Data_Argonaut_Decode_Error from "../Data.Argonaut.Decode.Error/index.js";
import * as Data_Array from "../Data.Array/index.js";
import * as Data_Array_NonEmpty from "../Data.Array.NonEmpty/index.js";
import * as Data_Array_NonEmpty_Internal from "../Data.Array.NonEmpty.Internal/index.js";
import * as Data_Bifunctor from "../Data.Bifunctor/index.js";
import * as Data_Boolean from "../Data.Boolean/index.js";
import * as Data_Either from "../Data.Either/index.js";
import * as Data_Foldable from "../Data.Foldable/index.js";
import * as Data_Function from "../Data.Function/index.js";
import * as Data_Functor from "../Data.Functor/index.js";
import * as Data_Identity from "../Data.Identity/index.js";
import * as Data_Int from "../Data.Int/index.js";
import * as Data_List from "../Data.List/index.js";
import * as Data_List_NonEmpty from "../Data.List.NonEmpty/index.js";
import * as Data_List_Types from "../Data.List.Types/index.js";
import * as Data_Map_Internal from "../Data.Map.Internal/index.js";
import * as Data_Maybe from "../Data.Maybe/index.js";
import * as Data_NonEmpty from "../Data.NonEmpty/index.js";
import * as Data_Set from "../Data.Set/index.js";
import * as Data_String_CodePoints from "../Data.String.CodePoints/index.js";
import * as Data_String_NonEmpty_Internal from "../Data.String.NonEmpty.Internal/index.js";
import * as Data_Traversable from "../Data.Traversable/index.js";
import * as Data_TraversableWithIndex from "../Data.TraversableWithIndex/index.js";
import * as Data_Tuple from "../Data.Tuple/index.js";
import * as Data_Unit from "../Data.Unit/index.js";
import * as Foreign_Object from "../Foreign.Object/index.js";
var fromFoldable = /* #__PURE__ */ Data_List.fromFoldable(Data_Foldable.foldableArray);
var fromFoldable1 = /* #__PURE__ */ Data_Set.fromFoldable(Data_List_Types.foldableList);
var traversableNonEmpty = /* #__PURE__ */ Data_NonEmpty.traversableNonEmpty(Data_Traversable.traversableArray);
var traversableNonEmpty1 = /* #__PURE__ */ Data_NonEmpty.traversableNonEmpty(Data_List_Types.traversableList);
var getFieldOptional$prime = function (decoder) {
    return function (obj) {
        return function (str) {
            var decode = function (json) {
                var $12 = Data_Argonaut_Core.isNull(json);
                if ($12) {
                    return Control_Applicative.pure(Data_Either.applicativeEither)(Data_Maybe.Nothing.value);
                };
                return Data_Functor.map(Data_Either.functorEither)(Data_Maybe.Just.create)(Data_Bifunctor.lmap(Data_Bifunctor.bifunctorEither)(Data_Argonaut_Decode_Error.AtKey.create(str))(decoder(json)));
            };
            return Data_Maybe.maybe(Control_Applicative.pure(Data_Either.applicativeEither)(Data_Maybe.Nothing.value))(decode)(Foreign_Object.lookup(str)(obj));
        };
    };
};
var getFieldOptional = function (decoder) {
    return function (obj) {
        return function (str) {
            var decode = (function () {
                var $21 = Data_Bifunctor.lmap(Data_Bifunctor.bifunctorEither)(Data_Argonaut_Decode_Error.AtKey.create(str));
                return function ($22) {
                    return $21(decoder($22));
                };
            })();
            return Data_Maybe.maybe(Control_Applicative.pure(Data_Either.applicativeEither)(Data_Maybe.Nothing.value))((function () {
                var $23 = Data_Functor.map(Data_Either.functorEither)(Data_Maybe.Just.create);
                return function ($24) {
                    return $23(decode($24));
                };
            })())(Foreign_Object.lookup(str)(obj));
        };
    };
};
var getField = function (decoder) {
    return function (obj) {
        return function (str) {
            return Data_Maybe.maybe(new Data_Either.Left(new Data_Argonaut_Decode_Error.AtKey(str, Data_Argonaut_Decode_Error.MissingValue.value)))((function () {
                var $25 = Data_Bifunctor.lmap(Data_Bifunctor.bifunctorEither)(Data_Argonaut_Decode_Error.AtKey.create(str));
                return function ($26) {
                    return $25(decoder($26));
                };
            })())(Foreign_Object.lookup(str)(obj));
        };
    };
};
var decodeVoid = function (v) {
    return new Data_Either.Left(new Data_Argonaut_Decode_Error.UnexpectedValue(Data_Argonaut_Core.fromString("Value cannot be Void")));
};
var decodeString = /* #__PURE__ */ (function () {
    return Data_Argonaut_Core.caseJsonString(new Data_Either.Left(new Data_Argonaut_Decode_Error.TypeMismatch("String")))(Data_Either.Right.create);
})();
var decodeNumber = /* #__PURE__ */ (function () {
    return Data_Argonaut_Core.caseJsonNumber(new Data_Either.Left(new Data_Argonaut_Decode_Error.TypeMismatch("Number")))(Data_Either.Right.create);
})();
var decodeNull = /* #__PURE__ */ (function () {
    return Data_Argonaut_Core.caseJsonNull(new Data_Either.Left(new Data_Argonaut_Decode_Error.TypeMismatch("null")))(Data_Function["const"](new Data_Either.Right(Data_Unit.unit)));
})();
var decodeNonEmptyString = function (json) {
    return Control_Bind.bindFlipped(Data_Either.bindEither)(Data_Either.note(new Data_Argonaut_Decode_Error.Named("NonEmptyString", new Data_Argonaut_Decode_Error.UnexpectedValue(json))))(Data_Functor.map(Data_Either.functorEither)(Data_String_NonEmpty_Internal.fromString)(decodeString(json)));
};
var decodeMaybe = function (decoder) {
    return function (json) {
        if (Data_Argonaut_Core.isNull(json)) {
            return Control_Applicative.pure(Data_Either.applicativeEither)(Data_Maybe.Nothing.value);
        };
        if (Data_Boolean.otherwise) {
            return Data_Functor.map(Data_Either.functorEither)(Data_Maybe.Just.create)(decoder(json));
        };
        throw new Error("Failed pattern match at Data.Argonaut.Decode.Decoders (line 37, column 1 - line 41, column 38): " + [ decoder.constructor.name, json.constructor.name ]);
    };
};
var decodeJObject = /* #__PURE__ */ (function () {
    var $27 = Data_Either.note(new Data_Argonaut_Decode_Error.TypeMismatch("Object"));
    return function ($28) {
        return $27(Data_Argonaut_Core.toObject($28));
    };
})();
var decodeJArray = /* #__PURE__ */ (function () {
    var $29 = Data_Either.note(new Data_Argonaut_Decode_Error.TypeMismatch("Array"));
    return function ($30) {
        return $29(Data_Argonaut_Core.toArray($30));
    };
})();
var decodeList = function (decoder) {
    return Control_Bind.composeKleisliFlipped(Data_Either.bindEither)((function () {
        var $31 = Data_Bifunctor.lmap(Data_Bifunctor.bifunctorEither)(Data_Argonaut_Decode_Error.Named.create("List"));
        var $32 = Data_Traversable.traverse(Data_List_Types.traversableList)(Data_Either.applicativeEither)(decoder);
        return function ($33) {
            return $31($32($33));
        };
    })())(Data_Functor.map(Data_Functor.functorFn)(Data_Functor.map(Data_Either.functorEither)(fromFoldable))(decodeJArray));
};
var decodeSet = function (dictOrd) {
    var fromFoldable2 = fromFoldable1(dictOrd);
    return function (decoder) {
        var $34 = Data_Functor.map(Data_Either.functorEither)(fromFoldable2);
        var $35 = decodeList(decoder);
        return function ($36) {
            return $34($35($36));
        };
    };
};
var decodeNonEmptyArray = function (decoder) {
    return Control_Bind.composeKleisliFlipped(Data_Either.bindEither)((function () {
        var $37 = Data_Bifunctor.lmap(Data_Bifunctor.bifunctorEither)(Data_Argonaut_Decode_Error.Named.create("NonEmptyArray"));
        var $38 = Data_Traversable.traverse(Data_Array_NonEmpty_Internal.traversableNonEmptyArray)(Data_Either.applicativeEither)(decoder);
        return function ($39) {
            return $37($38($39));
        };
    })())(Control_Bind.composeKleisliFlipped(Data_Either.bindEither)((function () {
        var $40 = Data_Functor.map(Data_Either.functorEither)(function (x) {
            return Data_Array_NonEmpty["cons$prime"](x.head)(x.tail);
        });
        var $41 = Data_Either.note(new Data_Argonaut_Decode_Error.TypeMismatch("NonEmptyArray"));
        return function ($42) {
            return $40($41(Data_Array.uncons($42)));
        };
    })())(decodeJArray));
};
var decodeNonEmptyList = function (decoder) {
    return Control_Bind.composeKleisliFlipped(Data_Either.bindEither)((function () {
        var $43 = Data_Bifunctor.lmap(Data_Bifunctor.bifunctorEither)(Data_Argonaut_Decode_Error.Named.create("NonEmptyList"));
        var $44 = Data_Traversable.traverse(Data_List_Types.traversableNonEmptyList)(Data_Either.applicativeEither)(decoder);
        return function ($45) {
            return $43($44($45));
        };
    })())(Control_Bind.composeKleisliFlipped(Data_Either.bindEither)((function () {
        var $46 = Data_Functor.map(Data_Either.functorEither)(function (x) {
            return Data_List_NonEmpty["cons$prime"](x.head)(x.tail);
        });
        var $47 = Data_Either.note(new Data_Argonaut_Decode_Error.TypeMismatch("NonEmptyList"));
        return function ($48) {
            return $46($47(Data_List.uncons($48)));
        };
    })())(Data_Functor.map(Data_Functor.functorFn)(Data_Functor.map(Data_Either.functorEither)(fromFoldable))(decodeJArray)));
};
var decodeNonEmpty_Array = function (decoder) {
    return Control_Bind.composeKleisliFlipped(Data_Either.bindEither)((function () {
        var $49 = Data_Bifunctor.lmap(Data_Bifunctor.bifunctorEither)(Data_Argonaut_Decode_Error.Named.create("NonEmpty Array"));
        var $50 = Data_Traversable.traverse(traversableNonEmpty)(Data_Either.applicativeEither)(decoder);
        return function ($51) {
            return $49($50($51));
        };
    })())(Control_Bind.composeKleisliFlipped(Data_Either.bindEither)((function () {
        var $52 = Data_Functor.map(Data_Either.functorEither)(function (x) {
            return new Data_NonEmpty.NonEmpty(x.head, x.tail);
        });
        var $53 = Data_Either.note(new Data_Argonaut_Decode_Error.TypeMismatch("NonEmpty Array"));
        return function ($54) {
            return $52($53(Data_Array.uncons($54)));
        };
    })())(decodeJArray));
};
var decodeNonEmpty_List = function (decoder) {
    return Control_Bind.composeKleisliFlipped(Data_Either.bindEither)((function () {
        var $55 = Data_Bifunctor.lmap(Data_Bifunctor.bifunctorEither)(Data_Argonaut_Decode_Error.Named.create("NonEmpty List"));
        var $56 = Data_Traversable.traverse(traversableNonEmpty1)(Data_Either.applicativeEither)(decoder);
        return function ($57) {
            return $55($56($57));
        };
    })())(Control_Bind.composeKleisliFlipped(Data_Either.bindEither)((function () {
        var $58 = Data_Functor.map(Data_Either.functorEither)(function (x) {
            return new Data_NonEmpty.NonEmpty(x.head, x.tail);
        });
        var $59 = Data_Either.note(new Data_Argonaut_Decode_Error.TypeMismatch("NonEmpty List"));
        return function ($60) {
            return $58($59(Data_List.uncons($60)));
        };
    })())(Data_Functor.map(Data_Functor.functorFn)(Data_Functor.map(Data_Either.functorEither)(fromFoldable))(decodeJArray)));
};
var decodeInt = /* #__PURE__ */ Control_Bind.composeKleisliFlipped(Data_Either.bindEither)(/* #__PURE__ */ (function () {
    var $61 = Data_Either.note(new Data_Argonaut_Decode_Error.TypeMismatch("Integer"));
    return function ($62) {
        return $61(Data_Int.fromNumber($62));
    };
})())(decodeNumber);
var decodeIdentity = function (decoder) {
    return function (json) {
        return Data_Functor.map(Data_Either.functorEither)(Data_Identity.Identity)(decoder(json));
    };
};
var decodeForeignObject = function (decoder) {
    return Control_Bind.composeKleisliFlipped(Data_Either.bindEither)((function () {
        var $63 = Data_Bifunctor.lmap(Data_Bifunctor.bifunctorEither)(Data_Argonaut_Decode_Error.Named.create("ForeignObject"));
        var $64 = Data_Traversable.traverse(Foreign_Object.traversableObject)(Data_Either.applicativeEither)(decoder);
        return function ($65) {
            return $63($64($65));
        };
    })())(decodeJObject);
};
var decodeEither = function (decoderA) {
    return function (decoderB) {
        return function (json) {
            return Data_Bifunctor.lmap(Data_Bifunctor.bifunctorEither)(Data_Argonaut_Decode_Error.Named.create("Either"))(Control_Bind.bind(Data_Either.bindEither)(decodeJObject(json))(function (obj) {
                return Control_Bind.bind(Data_Either.bindEither)(Data_Either.note(new Data_Argonaut_Decode_Error.AtKey("tag", Data_Argonaut_Decode_Error.MissingValue.value))(Foreign_Object.lookup("tag")(obj)))(function (tag) {
                    return Control_Bind.bind(Data_Either.bindEither)(Data_Either.note(new Data_Argonaut_Decode_Error.AtKey("value", Data_Argonaut_Decode_Error.MissingValue.value))(Foreign_Object.lookup("value")(obj)))(function (val) {
                        var v = Data_Argonaut_Core.toString(tag);
                        if (v instanceof Data_Maybe.Just && v.value0 === "Right") {
                            return Data_Functor.map(Data_Either.functorEither)(Data_Either.Right.create)(decoderB(val));
                        };
                        if (v instanceof Data_Maybe.Just && v.value0 === "Left") {
                            return Data_Functor.map(Data_Either.functorEither)(Data_Either.Left.create)(decoderA(val));
                        };
                        return new Data_Either.Left(new Data_Argonaut_Decode_Error.AtKey("tag", new Data_Argonaut_Decode_Error.UnexpectedValue(tag)));
                    });
                });
            }));
        };
    };
};
var decodeCodePoint = function (json) {
    return Control_Bind.bindFlipped(Data_Either.bindEither)(Data_Either.note(new Data_Argonaut_Decode_Error.Named("CodePoint", new Data_Argonaut_Decode_Error.UnexpectedValue(json))))(Data_Functor.map(Data_Either.functorEither)(Data_String_CodePoints.codePointAt(0))(decodeString(json)));
};
var decodeBoolean = /* #__PURE__ */ (function () {
    return Data_Argonaut_Core.caseJsonBoolean(new Data_Either.Left(new Data_Argonaut_Decode_Error.TypeMismatch("Boolean")))(Data_Either.Right.create);
})();
var decodeArray = function (decoder) {
    return Control_Bind.composeKleisliFlipped(Data_Either.bindEither)((function () {
        var $66 = Data_Bifunctor.lmap(Data_Bifunctor.bifunctorEither)(Data_Argonaut_Decode_Error.Named.create("Array"));
        var $67 = Data_TraversableWithIndex.traverseWithIndex(Data_TraversableWithIndex.traversableWithIndexArray)(Data_Either.applicativeEither)(function (i) {
            var $69 = Data_Bifunctor.lmap(Data_Bifunctor.bifunctorEither)(Data_Argonaut_Decode_Error.AtIndex.create(i));
            return function ($70) {
                return $69(decoder($70));
            };
        });
        return function ($68) {
            return $66($67($68));
        };
    })())(decodeJArray);
};
var decodeTuple = function (decoderA) {
    return function (decoderB) {
        return function (json) {
            var f = function (v) {
                if (v.length === 2) {
                    return Control_Apply.apply(Data_Either.applyEither)(Data_Functor.map(Data_Either.functorEither)(Data_Tuple.Tuple.create)(decoderA(v[0])))(decoderB(v[1]));
                };
                return new Data_Either.Left(new Data_Argonaut_Decode_Error.TypeMismatch("Tuple"));
            };
            return Control_Bind.bind(Data_Either.bindEither)(decodeArray(Data_Either.Right.create)(json))(f);
        };
    };
};
var decodeMap = function (dictOrd) {
    var fromFoldable2 = Data_Map_Internal.fromFoldable(dictOrd)(Data_List_Types.foldableList);
    return function (decoderA) {
        return function (decoderB) {
            var $71 = Data_Functor.map(Data_Either.functorEither)(fromFoldable2);
            var $72 = decodeList(decodeTuple(decoderA)(decoderB));
            return function ($73) {
                return $71($72($73));
            };
        };
    };
};
export {
    decodeIdentity,
    decodeMaybe,
    decodeTuple,
    decodeEither,
    decodeNull,
    decodeBoolean,
    decodeNumber,
    decodeInt,
    decodeString,
    decodeNonEmptyString,
    decodeNonEmpty_Array,
    decodeNonEmptyArray,
    decodeNonEmpty_List,
    decodeNonEmptyList,
    decodeCodePoint,
    decodeForeignObject,
    decodeArray,
    decodeList,
    decodeSet,
    decodeMap,
    decodeVoid,
    decodeJArray,
    decodeJObject,
    getField,
    getFieldOptional,
    getFieldOptional$prime
};
