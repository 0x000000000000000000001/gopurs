import * as Data_Argonaut_Core from "../Data.Argonaut.Core/index.js";
import * as Data_Array from "../Data.Array/index.js";
import * as Data_Array_NonEmpty from "../Data.Array.NonEmpty/index.js";
import * as Data_Either from "../Data.Either/index.js";
import * as Data_Function from "../Data.Function/index.js";
import * as Data_Functor from "../Data.Functor/index.js";
import * as Data_Int from "../Data.Int/index.js";
import * as Data_List from "../Data.List/index.js";
import * as Data_List_NonEmpty from "../Data.List.NonEmpty/index.js";
import * as Data_List_Types from "../Data.List.Types/index.js";
import * as Data_Map_Internal from "../Data.Map.Internal/index.js";
import * as Data_Maybe from "../Data.Maybe/index.js";
import * as Data_Set from "../Data.Set/index.js";
import * as Data_String_CodePoints from "../Data.String.CodePoints/index.js";
import * as Data_String_CodeUnits from "../Data.String.CodeUnits/index.js";
import * as Data_String_NonEmpty_Internal from "../Data.String.NonEmpty.Internal/index.js";
import * as Data_Tuple from "../Data.Tuple/index.js";
import * as Data_Unfoldable from "../Data.Unfoldable/index.js";
import * as Data_Void from "../Data.Void/index.js";
import * as Foreign_Object from "../Foreign.Object/index.js";
var toUnfoldable = /* #__PURE__ */ Data_List.toUnfoldable(Data_Unfoldable.unfoldableArray);
var toUnfoldable1 = /* #__PURE__ */ Data_Map_Internal.toUnfoldable(Data_List_Types.unfoldableList);
var toUnfoldable2 = /* #__PURE__ */ Data_Set.toUnfoldable(Data_List_Types.unfoldableList);
var fromFoldable = /* #__PURE__ */ Foreign_Object.fromFoldable(Data_List_Types.foldableList);
var extend = function (encoder) {
    return function (v) {
        var $36 = Data_Argonaut_Core.caseJsonObject(Data_Argonaut_Core.jsonSingletonObject(v.value0)(v.value1))((function () {
            var $38 = Foreign_Object.insert(v.value0)(v.value1);
            return function ($39) {
                return Data_Argonaut_Core.fromObject($38($39));
            };
        })());
        return function ($37) {
            return $36(encoder($37));
        };
    };
};

// | The named Encoders of the `(~>?)` operator.
var extendOptional = function (encoder) {
    return function (v) {
        if (v instanceof Data_Maybe.Just) {
            return extend(encoder)(v.value0);
        };
        if (v instanceof Data_Maybe.Nothing) {
            return encoder;
        };
        throw new Error("Failed pattern match at Data.Argonaut.Encode.Encoders (line 121, column 26 - line 123, column 21): " + [ v.constructor.name ]);
    };
};
var encodeVoid = Data_Void.absurd;
var encodeUnit = /* #__PURE__ */ Data_Function["const"](Data_Argonaut_Core.jsonNull);
var encodeTuple = function (encoderA) {
    return function (encoderB) {
        return function (v) {
            return Data_Argonaut_Core.fromArray([ encoderA(v.value0), encoderB(v.value1) ]);
        };
    };
};
var encodeString = Data_Argonaut_Core.fromString;
var encodeNumber = Data_Argonaut_Core.fromNumber;
var encodeNonEmptyString = function ($40) {
    return Data_Argonaut_Core.fromString(Data_String_NonEmpty_Internal.toString($40));
};
var encodeMaybe = function (encoder) {
    return function (v) {
        if (v instanceof Data_Maybe.Nothing) {
            return Data_Argonaut_Core.jsonNull;
        };
        if (v instanceof Data_Maybe.Just) {
            return encoder(v.value0);
        };
        throw new Error("Failed pattern match at Data.Argonaut.Encode.Encoders (line 31, column 23 - line 33, column 22): " + [ v.constructor.name ]);
    };
};
var encodeList = function (encoder) {
    var $41 = Data_Functor.map(Data_Functor.functorArray)(encoder);
    return function ($42) {
        return Data_Argonaut_Core.fromArray($41(toUnfoldable($42)));
    };
};
var encodeMap = function (dictOrd) {
    return function (encoderA) {
        return function (encoderB) {
            var $43 = encodeList(encodeTuple(encoderA)(encoderB));
            return function ($44) {
                return $43(toUnfoldable1($44));
            };
        };
    };
};
var encodeNonEmptyList = function (encoder) {
    var $45 = encodeList(encoder);
    return function ($46) {
        return $45(Data_List_NonEmpty.toList($46));
    };
};
var encodeNonEmpty_List = function (encoder) {
    return function (v) {
        return encodeList(encoder)(new Data_List_Types.Cons(v.value0, v.value1));
    };
};
var encodeSet = function (dictOrd) {
    return function (encoder) {
        var $47 = encodeList(encoder);
        return function ($48) {
            return $47(toUnfoldable2($48));
        };
    };
};
var encodeInt = function ($49) {
    return Data_Argonaut_Core.fromNumber(Data_Int.toNumber($49));
};
var encodeIdentity = function (encoder) {
    return function (v) {
        return encoder(v);
    };
};
var encodeForeignObject = function (encoder) {
    var $50 = Data_Functor.map(Foreign_Object.functorObject)(encoder);
    return function ($51) {
        return Data_Argonaut_Core.fromObject($50($51));
    };
};
var encodeEither = function (encoderA) {
    return function (encoderB) {
        var obj = function (encoder) {
            return function (tag) {
                return function (x) {
                    return Data_Argonaut_Core.fromObject(fromFoldable(new Data_List_Types.Cons(new Data_Tuple.Tuple("tag", Data_Argonaut_Core.fromString(tag)), new Data_List_Types.Cons(new Data_Tuple.Tuple("value", encoder(x)), Data_List_Types.Nil.value))));
                };
            };
        };
        return Data_Either.either(obj(encoderA)("Left"))(obj(encoderB)("Right"));
    };
};
var encodeCodePoint = function ($52) {
    return encodeString(Data_String_CodePoints.singleton($52));
};
var encodeChar = function ($53) {
    return encodeString(Data_String_CodeUnits.singleton($53));
};
var encodeBoolean = Data_Argonaut_Core.fromBoolean;
var encodeArray = function (encoder) {
    var $54 = Data_Functor.map(Data_Functor.functorArray)(encoder);
    return function ($55) {
        return Data_Argonaut_Core.fromArray($54($55));
    };
};
var encodeNonEmptyArray = function (encoder) {
    var $56 = encodeArray(encoder);
    return function ($57) {
        return $56(Data_Array_NonEmpty.toArray($57));
    };
};
var encodeNonEmpty_Array = function (encoder) {
    return function (v) {
        return encodeArray(encoder)(Data_Array.cons(v.value0)(v.value1));
    };
};
var assocOptional = function (encoder) {
    return function (k) {
        return Data_Functor.map(Data_Maybe.functorMaybe)((function () {
            var $58 = Data_Tuple.Tuple.create(k);
            return function ($59) {
                return $58(encoder($59));
            };
        })());
    };
};
var assoc = function (encoder) {
    return function (k) {
        var $60 = Data_Tuple.Tuple.create(k);
        return function ($61) {
            return $60(encoder($61));
        };
    };
};
export {
    encodeIdentity,
    encodeMaybe,
    encodeTuple,
    encodeEither,
    encodeUnit,
    encodeBoolean,
    encodeNumber,
    encodeInt,
    encodeString,
    encodeCodePoint,
    encodeNonEmptyString,
    encodeNonEmpty_Array,
    encodeNonEmptyArray,
    encodeNonEmpty_List,
    encodeNonEmptyList,
    encodeChar,
    encodeArray,
    encodeList,
    encodeForeignObject,
    encodeSet,
    encodeMap,
    encodeVoid,
    assoc,
    assocOptional,
    extend,
    extendOptional
};
