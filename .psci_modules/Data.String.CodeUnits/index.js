import * as $foreign from "./foreign.js";
import * as Data_Maybe from "../Data.Maybe/index.js";
import * as Data_String_Unsafe from "../Data.String.Unsafe/index.js";

// | Returns the first character and the rest of the string,
// | if the string is not empty.
// |
// | ```purescript
// | uncons "" == Nothing
// | uncons "Hello World" == Just { head: 'H', tail: "ello World" }
// | ```
// |
var uncons = function (v) {
    if (v === "") {
        return Data_Maybe.Nothing.value;
    };
    return new Data_Maybe.Just({
        head: Data_String_Unsafe.charAt(0)(v),
        tail: $foreign.drop(1)(v)
    });
};

// | Converts the string to a character, if the length of the string is
// | exactly `1`.
// |
// | ```purescript
// | toChar "l" == Just 'l'
// | toChar "Hi" == Nothing -- since length is not 1
// | ```
var toChar = /* #__PURE__ */ (function () {
    return $foreign["_toChar"](Data_Maybe.Just.create)(Data_Maybe.Nothing.value);
})();

// | Returns the longest prefix (possibly empty) of characters that satisfy
// | the predicate.
// |
// | ```purescript
// | takeWhile (_ /= ':') "http://purescript.org" == "http"
// | ```
// |
var takeWhile = function (p) {
    return function (s) {
        return $foreign.take($foreign.countPrefix(p)(s))(s);
    };
};

// | Returns the last `n` characters of the string.
// |
// | ```purescript
// | takeRight 5 "Hello World" == "World"
// | ```
// |
var takeRight = function (i) {
    return function (s) {
        return $foreign.drop($foreign.length(s) - i | 0)(s);
    };
};

// | If the string ends with the given suffix, return the portion of the
// | string left after removing it, as a `Just` value. Otherwise, return
// | `Nothing`.
// |
// | ```purescript
// | stripSuffix (Pattern ".exe") "psc.exe" == Just "psc"
// | stripSuffix (Pattern ".exe") "psc" == Nothing
// | ```
var stripSuffix = function (v) {
    return function (str) {
        var v1 = $foreign.splitAt($foreign.length(str) - $foreign.length(v) | 0)(str);
        var $11 = v1.after === v;
        if ($11) {
            return new Data_Maybe.Just(v1.before);
        };
        return Data_Maybe.Nothing.value;
    };
};

//-----------------------------------------------------------------------------
// `stripPrefix`, `stripSuffix`, and `contains` are CodeUnit/CodePoint agnostic
// as they are based on patterns rather than lengths/indices, but they need to
// be defined in here to avoid a circular module dependency
//-----------------------------------------------------------------------------
// | If the string starts with the given prefix, return the portion of the
// | string left after removing it, as a `Just` value. Otherwise, return `Nothing`.
// |
// | ```purescript
// | stripPrefix (Pattern "http:") "http://purescript.org" == Just "//purescript.org"
// | stripPrefix (Pattern "http:") "https://purescript.org" == Nothing
// | ```
var stripPrefix = function (v) {
    return function (str) {
        var v1 = $foreign.splitAt($foreign.length(v))(str);
        var $17 = v1.before === v;
        if ($17) {
            return new Data_Maybe.Just(v1.after);
        };
        return Data_Maybe.Nothing.value;
    };
};

// | Returns the index of the last occurrence of the pattern in the
// | given string, starting at the specified index and searching
// | backwards towards the beginning of the string.
// |
// | Starting at a negative index is equivalent to starting at 0 and
// | starting at an index greater than the string length is equivalent
// | to searching in the whole string.
// |
// | Returns `Nothing` if there is no match.
// |
// | ```purescript
// | lastIndexOf' (Pattern "a") (-1) "ababa" == Just 0
// | lastIndexOf' (Pattern "a") 1 "ababa" == Just 0
// | lastIndexOf' (Pattern "a") 3 "ababa" == Just 2
// | lastIndexOf' (Pattern "a") 4 "ababa" == Just 4
// | lastIndexOf' (Pattern "a") 5 "ababa" == Just 4
// | ```
// |
var lastIndexOf$prime = /* #__PURE__ */ (function () {
    return $foreign["_lastIndexOfStartingAt"](Data_Maybe.Just.create)(Data_Maybe.Nothing.value);
})();

// | Returns the index of the last occurrence of the pattern in the
// | given string. Returns `Nothing` if there is no match.
// |
// | ```purescript
// | lastIndexOf (Pattern "c") "abcdc" == Just 4
// | lastIndexOf (Pattern "c") "aaa" == Nothing
// | ```
// |
var lastIndexOf = /* #__PURE__ */ (function () {
    return $foreign["_lastIndexOf"](Data_Maybe.Just.create)(Data_Maybe.Nothing.value);
})();

// | Returns the index of the first occurrence of the pattern in the
// | given string, starting at the specified index. Returns `Nothing` if there is
// | no match.
// |
// | ```purescript
// | indexOf' (Pattern "a") 2 "ababa" == Just 2
// | indexOf' (Pattern "a") 3 "ababa" == Just 4
// | ```
// |
var indexOf$prime = /* #__PURE__ */ (function () {
    return $foreign["_indexOfStartingAt"](Data_Maybe.Just.create)(Data_Maybe.Nothing.value);
})();

// | Returns the index of the first occurrence of the pattern in the
// | given string. Returns `Nothing` if there is no match.
// |
// | ```purescript
// | indexOf (Pattern "c") "abcdc" == Just 2
// | indexOf (Pattern "c") "aaa" == Nothing
// | ```
// |
var indexOf = /* #__PURE__ */ (function () {
    return $foreign["_indexOf"](Data_Maybe.Just.create)(Data_Maybe.Nothing.value);
})();

// | Returns the suffix remaining after `takeWhile`.
// |
// | ```purescript
// | dropWhile (_ /= '.') "Test.purs" == ".purs"
// | ```
// |
var dropWhile = function (p) {
    return function (s) {
        return $foreign.drop($foreign.countPrefix(p)(s))(s);
    };
};

// | Returns the string without the last `n` characters.
// |
// | ```purescript
// | dropRight 6 "Hello World" == "Hello"
// | ```
// |
var dropRight = function (i) {
    return function (s) {
        return $foreign.take($foreign.length(s) - i | 0)(s);
    };
};

// | Checks whether the pattern appears in the given string.
// |
// | ```purescript
// | contains (Pattern "needle") "haystack with needle" == true
// | contains (Pattern "needle") "haystack" == false
// | ```
var contains = function (pat) {
    var $20 = indexOf(pat);
    return function ($21) {
        return Data_Maybe.isJust($20($21));
    };
};

// | Returns the character at the given index, if the index is within bounds.
// |
// | ```purescript
// | charAt 2 "Hello" == Just 'l'
// | charAt 10 "Hello" == Nothing
// | ```
// |
var charAt = /* #__PURE__ */ (function () {
    return $foreign["_charAt"](Data_Maybe.Just.create)(Data_Maybe.Nothing.value);
})();
export {
    singleton,
    fromCharArray,
    toCharArray,
    length,
    countPrefix,
    take,
    drop,
    slice,
    splitAt
} from "./foreign.js";
export {
    stripPrefix,
    stripSuffix,
    contains,
    charAt,
    toChar,
    uncons,
    indexOf,
    indexOf$prime,
    lastIndexOf,
    lastIndexOf$prime,
    takeRight,
    takeWhile,
    dropRight,
    dropWhile
};
