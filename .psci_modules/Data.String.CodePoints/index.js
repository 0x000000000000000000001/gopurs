// | These functions allow PureScript strings to be treated as if they were
// | sequences of Unicode code points instead of their true underlying
// | implementation (sequences of UTF-16 code units). For nearly all uses of
// | strings, these functions should be preferred over the ones in
// | `Data.String.CodeUnits`.
import * as $foreign from "./foreign.js";
import * as Data_Array from "../Data.Array/index.js";
import * as Data_Boolean from "../Data.Boolean/index.js";
import * as Data_Bounded from "../Data.Bounded/index.js";
import * as Data_Enum from "../Data.Enum/index.js";
import * as Data_EuclideanRing from "../Data.EuclideanRing/index.js";
import * as Data_Functor from "../Data.Functor/index.js";
import * as Data_Int from "../Data.Int/index.js";
import * as Data_Maybe from "../Data.Maybe/index.js";
import * as Data_Ord from "../Data.Ord/index.js";
import * as Data_String_CodeUnits from "../Data.String.CodeUnits/index.js";
import * as Data_String_Common from "../Data.String.Common/index.js";
import * as Data_String_Unsafe from "../Data.String.Unsafe/index.js";
import * as Data_Tuple from "../Data.Tuple/index.js";
import * as Data_Unfoldable from "../Data.Unfoldable/index.js";
var $runtime_lazy = function (name, moduleName, init) {
    var state = 0;
    var val;
    return function (lineNumber) {
        if (state === 2) return val;
        if (state === 1) throw new ReferenceError(name + " was needed before it finished initializing (module " + moduleName + ", line " + lineNumber + ")", moduleName, lineNumber);
        state = 1;
        val = init();
        state = 2;
        return val;
    };
};

// | CodePoint is an `Int` bounded between `0` and `0x10FFFF`, corresponding to
// | Unicode code points.
var CodePoint = function (x) {
    return x;
};
var unsurrogate = function (lead) {
    return function (trail) {
        return (((lead - 55296 | 0) * 1024 | 0) + (trail - 56320 | 0) | 0) + 65536 | 0;
    };
};
var showCodePoint = {
    show: function (v) {
        return "(CodePoint 0x" + (Data_String_Common.toUpper(Data_Int.toStringAs(Data_Int.hexadecimal)(v)) + ")");
    }
};
var isTrail = function (cu) {
    return 56320 <= cu && cu <= 57343;
};
var isLead = function (cu) {
    return 55296 <= cu && cu <= 56319;
};

// | Returns a record with the first code point and the remaining code points
// | of the string. Returns `Nothing` if the string is empty. Operates in
// | constant space and time.
// |
// | ```purescript
// | >>> uncons "𝐀𝐀 c 𝐀"
// | Just { head: CodePoint 0x1D400, tail: "𝐀 c 𝐀" }
// | >>> uncons ""
// | Nothing
// | ```
// |
var uncons = function (s) {
    var v = Data_String_CodeUnits.length(s);
    if (v === 0) {
        return Data_Maybe.Nothing.value;
    };
    if (v === 1) {
        return new Data_Maybe.Just({
            head: Data_Enum.fromEnum(Data_Enum.boundedEnumChar)(Data_String_Unsafe.charAt(0)(s)),
            tail: ""
        });
    };
    var cu1 = Data_Enum.fromEnum(Data_Enum.boundedEnumChar)(Data_String_Unsafe.charAt(1)(s));
    var cu0 = Data_Enum.fromEnum(Data_Enum.boundedEnumChar)(Data_String_Unsafe.charAt(0)(s));
    var $23 = isLead(cu0) && isTrail(cu1);
    if ($23) {
        return new Data_Maybe.Just({
            head: unsurrogate(cu0)(cu1),
            tail: Data_String_CodeUnits.drop(2)(s)
        });
    };
    return new Data_Maybe.Just({
        head: cu0,
        tail: Data_String_CodeUnits.drop(1)(s)
    });
};
var unconsButWithTuple = function (s) {
    return Data_Functor.map(Data_Maybe.functorMaybe)(function (v) {
        return new Data_Tuple.Tuple(v.head, v.tail);
    })(uncons(s));
};
var toCodePointArrayFallback = function (s) {
    return Data_Unfoldable.unfoldr(Data_Unfoldable.unfoldableArray)(unconsButWithTuple)(s);
};
var unsafeCodePointAt0Fallback = function (s) {
    var cu0 = Data_Enum.fromEnum(Data_Enum.boundedEnumChar)(Data_String_Unsafe.charAt(0)(s));
    var $27 = isLead(cu0) && Data_String_CodeUnits.length(s) > 1;
    if ($27) {
        var cu1 = Data_Enum.fromEnum(Data_Enum.boundedEnumChar)(Data_String_Unsafe.charAt(1)(s));
        var $28 = isTrail(cu1);
        if ($28) {
            return unsurrogate(cu0)(cu1);
        };
        return cu0;
    };
    return cu0;
};

// WARN: this function expects the String parameter to be non-empty
var unsafeCodePointAt0 = /* #__PURE__ */ $foreign["_unsafeCodePointAt0"](unsafeCodePointAt0Fallback);

// | Creates an array of code points from a string. Operates in space and time
// | linear to the length of the string.
// |
// | ```purescript
// | >>> codePointArray = toCodePointArray "b 𝐀𝐀"
// | >>> codePointArray
// | [CodePoint 0x62, CodePoint 0x20, CodePoint 0x1D400, CodePoint 0x1D400]
// | >>> map singleton codePointArray
// | ["b", " ", "𝐀", "𝐀"]
// | ```
// |
var toCodePointArray = /* #__PURE__ */ $foreign["_toCodePointArray"](toCodePointArrayFallback)(unsafeCodePointAt0);

// | Returns the number of code points in the string. Operates in constant
// | space and in time linear to the length of the string.
// |
// | ```purescript
// | >>> length "b 𝐀𝐀 c 𝐀"
// | 8
// | -- compare to Data.String:
// | >>> length "b 𝐀𝐀 c 𝐀"
// | 11
// | ```
// |
var length = function ($54) {
    return Data_Array.length(toCodePointArray($54));
};

// | Returns the number of code points preceding the last match of the given
// | pattern in the string. Returns `Nothing` when no matches are found.
// |
// | ```purescript
// | >>> lastIndexOf (Pattern "𝐀") "b 𝐀𝐀 c 𝐀"
// | Just 7
// | >>> lastIndexOf (Pattern "o") "b 𝐀𝐀 c 𝐀"
// | Nothing
// | ```
// |
var lastIndexOf = function (p) {
    return function (s) {
        return Data_Functor.map(Data_Maybe.functorMaybe)(function (i) {
            return length(Data_String_CodeUnits.take(i)(s));
        })(Data_String_CodeUnits.lastIndexOf(p)(s));
    };
};

// | Returns the number of code points preceding the first match of the given
// | pattern in the string. Returns `Nothing` when no matches are found.
// |
// | ```purescript
// | >>> indexOf (Pattern "𝐀") "b 𝐀𝐀 c 𝐀"
// | Just 2
// | >>> indexOf (Pattern "o") "b 𝐀𝐀 c 𝐀"
// | Nothing
// | ```
// |
var indexOf = function (p) {
    return function (s) {
        return Data_Functor.map(Data_Maybe.functorMaybe)(function (i) {
            return length(Data_String_CodeUnits.take(i)(s));
        })(Data_String_CodeUnits.indexOf(p)(s));
    };
};
var fromCharCode = /* #__PURE__ */ (function () {
    var $55 = Data_Enum.toEnumWithDefaults(Data_Enum.boundedEnumChar)(Data_Bounded.bottom(Data_Bounded.boundedChar))(Data_Bounded.top(Data_Bounded.boundedChar));
    return function ($56) {
        return Data_String_CodeUnits.singleton($55($56));
    };
})();
var singletonFallback = function (v) {
    if (v <= 65535) {
        return fromCharCode(v);
    };
    var lead = Data_EuclideanRing.div(Data_EuclideanRing.euclideanRingInt)(v - 65536 | 0)(1024) + 55296 | 0;
    var trail = Data_EuclideanRing.mod(Data_EuclideanRing.euclideanRingInt)(v - 65536 | 0)(1024) + 56320 | 0;
    return fromCharCode(lead) + fromCharCode(trail);
};

// | Creates a string from an array of code points. Operates in space and time
// | linear to the length of the array.
// |
// | ```purescript
// | >>> codePointArray = toCodePointArray "c 𝐀"
// | >>> codePointArray
// | [CodePoint 0x63, CodePoint 0x20, CodePoint 0x1D400]
// | >>> fromCodePointArray codePointArray
// | "c 𝐀"
// | ```
// |
var fromCodePointArray = /* #__PURE__ */ $foreign["_fromCodePointArray"](singletonFallback);

// | Creates a string containing just the given code point. Operates in
// | constant space and time.
// |
// | ```purescript
// | >>> map singleton (toEnum 0x1D400)
// | Just "𝐀"
// | ```
// |
var singleton = /* #__PURE__ */ $foreign["_singleton"](singletonFallback);
var takeFallback = function (v) {
    return function (v1) {
        if (v < 1) {
            return "";
        };
        var v2 = uncons(v1);
        if (v2 instanceof Data_Maybe.Just) {
            return singleton(v2.value0.head) + takeFallback(v - 1 | 0)(v2.value0.tail);
        };
        return v1;
    };
};

// | Returns a string containing the given number of code points from the
// | beginning of the given string. If the string does not have that many code
// | points, returns the empty string. Operates in constant space and in time
// | linear to the given number.
// |
// | ```purescript
// | >>> take 3 "b 𝐀𝐀 c 𝐀"
// | "b 𝐀"
// | -- compare to Data.String:
// | >>> take 3 "b 𝐀𝐀 c 𝐀"
// | "b �"
// | ```
// |
var take = /* #__PURE__ */ $foreign["_take"](takeFallback);

// | Returns the number of code points preceding the first match of the given
// | pattern in the string. Pattern matches following the given index will be
// | ignored.
// |
// | Giving a negative index is equivalent to giving 0 and giving an index
// | greater than the number of code points in the string is equivalent to
// | searching in the whole string.
// |
// | Returns `Nothing` when no matches are found.
// |
// | ```purescript
// | >>> lastIndexOf' (Pattern "𝐀") (-1) "b 𝐀𝐀 c 𝐀"
// | Nothing
// | >>> lastIndexOf' (Pattern "𝐀") 0 "b 𝐀𝐀 c 𝐀"
// | Nothing
// | >>> lastIndexOf' (Pattern "𝐀") 5 "b 𝐀𝐀 c 𝐀"
// | Just 3
// | >>> lastIndexOf' (Pattern "𝐀") 8 "b 𝐀𝐀 c 𝐀"
// | Just 7
// | >>> lastIndexOf' (Pattern "o") 5 "b 𝐀𝐀 c 𝐀"
// | Nothing
// | ```
// |
var lastIndexOf$prime = function (p) {
    return function (i) {
        return function (s) {
            var i$prime = Data_String_CodeUnits.length(take(i)(s));
            return Data_Functor.map(Data_Maybe.functorMaybe)(function (k) {
                return length(Data_String_CodeUnits.take(k)(s));
            })(Data_String_CodeUnits["lastIndexOf$prime"](p)(i$prime)(s));
        };
    };
};

// | Splits a string into two substrings, where `before` contains the code
// | points up to (but not including) the given index, and `after` contains the
// | rest of the string, from that index on.
// |
// | ```purescript
// | >>> splitAt 3 "b 𝐀𝐀 c 𝐀"
// | { before: "b 𝐀", after: "𝐀 c 𝐀" }
// | ```
// |
// | Thus the length of `(splitAt i s).before` will equal either `i` or
// | `length s`, if that is shorter. (Or if `i` is negative the length will be
// | 0.)
// |
// | In code:
// | ```purescript
// | length (splitAt i s).before == min (max i 0) (length s)
// | (splitAt i s).before <> (splitAt i s).after == s
// | splitAt i s == {before: take i s, after: drop i s}
// | ```
var splitAt = function (i) {
    return function (s) {
        var before = take(i)(s);
        return {
            before: before,
            after: Data_String_CodeUnits.drop(Data_String_CodeUnits.length(before))(s)
        };
    };
};
var eqCodePoint = {
    eq: function (x) {
        return function (y) {
            return x === y;
        };
    }
};
var ordCodePoint = {
    compare: function (x) {
        return function (y) {
            return Data_Ord.compare(Data_Ord.ordInt)(x)(y);
        };
    },
    Eq0: function () {
        return eqCodePoint;
    }
};

// | Drops the given number of code points from the beginning of the string. If
// | the string does not have that many code points, returns the empty string.
// | Operates in constant space and in time linear to the given number.
// |
// | ```purescript
// | >>> drop 5 "𝐀𝐀 b c"
// | "c"
// | -- compared to Data.String:
// | >>> drop 5 "𝐀𝐀 b c"
// | "b c" -- because "𝐀" occupies 2 code units
// | ```
// |
var drop = function (n) {
    return function (s) {
        return Data_String_CodeUnits.drop(Data_String_CodeUnits.length(take(n)(s)))(s);
    };
};

// | Returns the number of code points preceding the first match of the given
// | pattern in the string. Pattern matches preceding the given index will be
// | ignored. Returns `Nothing` when no matches are found.
// |
// | ```purescript
// | >>> indexOf' (Pattern "𝐀") 4 "b 𝐀𝐀 c 𝐀"
// | Just 7
// | >>> indexOf' (Pattern "o") 4 "b 𝐀𝐀 c 𝐀"
// | Nothing
// | ```
// |
var indexOf$prime = function (p) {
    return function (i) {
        return function (s) {
            var s$prime = drop(i)(s);
            return Data_Functor.map(Data_Maybe.functorMaybe)(function (k) {
                return i + length(Data_String_CodeUnits.take(k)(s$prime)) | 0;
            })(Data_String_CodeUnits.indexOf(p)(s$prime));
        };
    };
};
var countTail = function ($copy_p) {
    return function ($copy_s) {
        return function ($copy_accum) {
            var $tco_var_p = $copy_p;
            var $tco_var_s = $copy_s;
            var $tco_done = false;
            var $tco_result;
            function $tco_loop(p, s, accum) {
                var v = uncons(s);
                if (v instanceof Data_Maybe.Just) {
                    var $41 = p(v.value0.head);
                    if ($41) {
                        $tco_var_p = p;
                        $tco_var_s = v.value0.tail;
                        $copy_accum = accum + 1 | 0;
                        return;
                    };
                    $tco_done = true;
                    return accum;
                };
                $tco_done = true;
                return accum;
            };
            while (!$tco_done) {
                $tco_result = $tco_loop($tco_var_p, $tco_var_s, $copy_accum);
            };
            return $tco_result;
        };
    };
};
var countFallback = function (p) {
    return function (s) {
        return countTail(p)(s)(0);
    };
};

// | Returns the number of code points in the leading sequence of code points
// | which all match the given predicate. Operates in constant space and in
// | time linear to the length of the string.
// |
// | ```purescript
// | >>> countPrefix (\c -> fromEnum c == 0x1D400) "𝐀𝐀 b c 𝐀"
// | 2
// | ```
// |
var countPrefix = /* #__PURE__ */ $foreign["_countPrefix"](countFallback)(unsafeCodePointAt0);

// | Drops the leading sequence of code points which all match the given
// | predicate from the string. Operates in constant space and in time linear
// | to the length of the string.
// |
// | ```purescript
// | >>> dropWhile (\c -> fromEnum c == 0x1D400) "𝐀𝐀 b c 𝐀"
// | " b c 𝐀"
// | ```
// |
var dropWhile = function (p) {
    return function (s) {
        return drop(countPrefix(p)(s))(s);
    };
};

// | Returns a string containing the leading sequence of code points which all
// | match the given predicate from the string. Operates in constant space and
// | in time linear to the length of the string.
// |
// | ```purescript
// | >>> takeWhile (\c -> fromEnum c == 0x1D400) "𝐀𝐀 b c 𝐀"
// | "𝐀𝐀"
// | ```
// |
var takeWhile = function (p) {
    return function (s) {
        return take(countPrefix(p)(s))(s);
    };
};

// | Creates a `CodePoint` from a given `Char`.
// |
// | ```purescript
// | >>> codePointFromChar 'B'
// | CodePoint 0x42 -- represents 'B'
// | ```
// |
var codePointFromChar = /* #__PURE__ */ (function () {
    var $57 = Data_Enum.fromEnum(Data_Enum.boundedEnumChar);
    return function ($58) {
        return CodePoint($57($58));
    };
})();
var codePointAtFallback = function ($copy_n) {
    return function ($copy_s) {
        var $tco_var_n = $copy_n;
        var $tco_done = false;
        var $tco_result;
        function $tco_loop(n, s) {
            var v = uncons(s);
            if (v instanceof Data_Maybe.Just) {
                var $46 = n === 0;
                if ($46) {
                    $tco_done = true;
                    return new Data_Maybe.Just(v.value0.head);
                };
                $tco_var_n = n - 1 | 0;
                $copy_s = v.value0.tail;
                return;
            };
            $tco_done = true;
            return Data_Maybe.Nothing.value;
        };
        while (!$tco_done) {
            $tco_result = $tco_loop($tco_var_n, $copy_s);
        };
        return $tco_result;
    };
};

// | Returns the first code point of the string after dropping the given number
// | of code points from the beginning, if there is such a code point. Operates
// | in constant space and in time linear to the given index.
// |
// | ```purescript
// | >>> codePointAt 1 "𝐀𝐀𝐀𝐀"
// | Just (CodePoint 0x1D400) -- represents "𝐀"
// | -- compare to Data.String:
// | >>> charAt 1 "𝐀𝐀𝐀𝐀"
// | Just '�'
// | ```
// |
var codePointAt = function (v) {
    return function (v1) {
        if (v < 0) {
            return Data_Maybe.Nothing.value;
        };
        if (v === 0 && v1 === "") {
            return Data_Maybe.Nothing.value;
        };
        if (v === 0) {
            return new Data_Maybe.Just(unsafeCodePointAt0(v1));
        };
        return $foreign["_codePointAt"](codePointAtFallback)(Data_Maybe.Just.create)(Data_Maybe.Nothing.value)(unsafeCodePointAt0)(v)(v1);
    };
};
var boundedCodePoint = {
    bottom: 0,
    top: 1114111,
    Ord0: function () {
        return ordCodePoint;
    }
};
var boundedEnumCodePoint = /* #__PURE__ */ (function () {
    return {
        cardinality: 1114111 + 1 | 0,
        fromEnum: function (v) {
            return v;
        },
        toEnum: function (n) {
            if (n >= 0 && n <= 1114111) {
                return new Data_Maybe.Just(n);
            };
            if (Data_Boolean.otherwise) {
                return Data_Maybe.Nothing.value;
            };
            throw new Error("Failed pattern match at Data.String.CodePoints (line 63, column 1 - line 68, column 26): " + [ n.constructor.name ]);
        },
        Bounded0: function () {
            return boundedCodePoint;
        },
        Enum1: function () {
            return $lazy_enumCodePoint(0);
        }
    };
})();
var $lazy_enumCodePoint = /* #__PURE__ */ $runtime_lazy("enumCodePoint", "Data.String.CodePoints", function () {
    return {
        succ: Data_Enum.defaultSucc(Data_Enum.toEnum(boundedEnumCodePoint))(Data_Enum.fromEnum(boundedEnumCodePoint)),
        pred: Data_Enum.defaultPred(Data_Enum.toEnum(boundedEnumCodePoint))(Data_Enum.fromEnum(boundedEnumCodePoint)),
        Ord0: function () {
            return ordCodePoint;
        }
    };
});
var enumCodePoint = /* #__PURE__ */ $lazy_enumCodePoint(59);
export {
    codePointFromChar,
    singleton,
    fromCodePointArray,
    toCodePointArray,
    codePointAt,
    uncons,
    length,
    countPrefix,
    indexOf,
    indexOf$prime,
    lastIndexOf,
    lastIndexOf$prime,
    take,
    takeWhile,
    drop,
    dropWhile,
    splitAt,
    eqCodePoint,
    ordCodePoint,
    showCodePoint,
    boundedCodePoint,
    enumCodePoint,
    boundedEnumCodePoint
};
export {
    contains,
    stripPrefix,
    stripSuffix
} from "../Data.String.CodeUnits/index.js";
