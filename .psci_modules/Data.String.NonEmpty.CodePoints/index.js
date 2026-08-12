import * as Data_Array_NonEmpty from "../Data.Array.NonEmpty/index.js";
import * as Data_Maybe from "../Data.Maybe/index.js";
import * as Data_Semigroup_Foldable from "../Data.Semigroup.Foldable/index.js";
import * as Data_String_CodePoints from "../Data.String.CodePoints/index.js";
import * as Data_String_NonEmpty_Internal from "../Data.String.NonEmpty.Internal/index.js";

// For internal use only. Do not export.
var toNonEmptyString = Data_String_NonEmpty_Internal.NonEmptyString;
var snoc = function (c) {
    return function (s) {
        return toNonEmptyString(s + Data_String_CodePoints.singleton(c));
    };
};
var singleton = function ($14) {
    return toNonEmptyString(Data_String_CodePoints.singleton($14));
};

// For internal use only. Do not export.
var liftS = function (f) {
    return function (v) {
        return f(v);
    };
};
var takeWhile = function (f) {
    var $15 = liftS(Data_String_CodePoints.takeWhile(f));
    return function ($16) {
        return Data_String_NonEmpty_Internal.fromString($15($16));
    };
};
var lastIndexOf$prime = function (pat) {
    var $17 = Data_String_CodePoints["lastIndexOf$prime"](pat);
    return function ($18) {
        return liftS($17($18));
    };
};
var lastIndexOf = function ($19) {
    return liftS(Data_String_CodePoints.lastIndexOf($19));
};
var indexOf$prime = function (pat) {
    var $20 = Data_String_CodePoints["indexOf$prime"](pat);
    return function ($21) {
        return liftS($20($21));
    };
};
var indexOf = function ($22) {
    return liftS(Data_String_CodePoints.indexOf($22));
};

// For internal use only. Do not export.
var fromNonEmptyString = function (v) {
    return v;
};
var length = function ($23) {
    return Data_String_CodePoints.length(fromNonEmptyString($23));
};
var splitAt = function (i) {
    return function (nes) {
        var v = Data_String_CodePoints.splitAt(i)(fromNonEmptyString(nes));
        return {
            before: Data_String_NonEmpty_Internal.fromString(v.before),
            after: Data_String_NonEmpty_Internal.fromString(v.after)
        };
    };
};
var take = function (i) {
    return function (nes) {
        var s = fromNonEmptyString(nes);
        var $11 = i < 1;
        if ($11) {
            return Data_Maybe.Nothing.value;
        };
        return new Data_Maybe.Just(toNonEmptyString(Data_String_CodePoints.take(i)(s)));
    };
};
var toCodePointArray = function ($24) {
    return Data_String_CodePoints.toCodePointArray(fromNonEmptyString($24));
};
var toNonEmptyCodePointArray = /* #__PURE__ */ (function () {
    var $25 = Data_Maybe.fromJust();
    return function ($26) {
        return $25(Data_Array_NonEmpty.fromArray(toCodePointArray($26)));
    };
})();
var uncons = function (nes) {
    var s = fromNonEmptyString(nes);
    return {
        head: Data_Maybe.fromJust()(Data_String_CodePoints.codePointAt(0)(s)),
        tail: Data_String_NonEmpty_Internal.fromString(Data_String_CodePoints.drop(1)(s))
    };
};
var fromFoldable1 = function (dictFoldable1) {
    return Data_Semigroup_Foldable.foldMap1(dictFoldable1)(Data_String_NonEmpty_Internal.semigroupNonEmptyString)(singleton);
};
var fromCodePointArray = function (v) {
    if (v.length === 0) {
        return Data_Maybe.Nothing.value;
    };
    return new Data_Maybe.Just(toNonEmptyString(Data_String_CodePoints.fromCodePointArray(v)));
};
var fromNonEmptyCodePointArray = /* #__PURE__ */ (function () {
    var $27 = Data_Maybe.fromJust();
    return function ($28) {
        return $27(fromCodePointArray(Data_Array_NonEmpty.toArray($28)));
    };
})();
var dropWhile = function (f) {
    var $29 = liftS(Data_String_CodePoints.dropWhile(f));
    return function ($30) {
        return Data_String_NonEmpty_Internal.fromString($29($30));
    };
};
var drop = function (i) {
    return function (nes) {
        var s = fromNonEmptyString(nes);
        var $13 = i >= Data_String_CodePoints.length(s);
        if ($13) {
            return Data_Maybe.Nothing.value;
        };
        return new Data_Maybe.Just(toNonEmptyString(Data_String_CodePoints.drop(i)(s)));
    };
};
var countPrefix = function ($31) {
    return liftS(Data_String_CodePoints.countPrefix($31));
};
var cons = function (c) {
    return function (s) {
        return toNonEmptyString(Data_String_CodePoints.singleton(c) + s);
    };
};
var codePointAt = function ($32) {
    return liftS(Data_String_CodePoints.codePointAt($32));
};
export {
    fromCodePointArray,
    fromNonEmptyCodePointArray,
    singleton,
    cons,
    snoc,
    fromFoldable1,
    toCodePointArray,
    toNonEmptyCodePointArray,
    codePointAt,
    indexOf,
    indexOf$prime,
    lastIndexOf,
    lastIndexOf$prime,
    uncons,
    length,
    take,
    takeWhile,
    drop,
    dropWhile,
    countPrefix,
    splitAt
};
