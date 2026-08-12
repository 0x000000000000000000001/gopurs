// | Wraps Javascript's `RegExp` object that enables matching strings with
// | patterns defined by regular expressions.
// | For details of the underlying implementation, see [RegExp Reference at MDN](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/RegExp).
import * as $foreign from "./foreign.js";
import * as Data_Either from "../Data.Either/index.js";
import * as Data_Maybe from "../Data.Maybe/index.js";
import * as Data_String_CodeUnits from "../Data.String.CodeUnits/index.js";
import * as Data_String_Regex_Flags from "../Data.String.Regex.Flags/index.js";
var showRegex = {
    show: $foreign.showRegexImpl
};

// | Returns `Just` the index of the first match of the `Regex` in the string,
// | or `Nothing` if there is no match.
var search = /* #__PURE__ */ (function () {
    return $foreign["_search"](Data_Maybe.Just.create)(Data_Maybe.Nothing.value);
})();

// | Transforms occurrences of the `Regex` using a function of the matched
// | substring and a list of captured substrings of type `Maybe String`,
// | where `Nothing` represents an unmatched optional capturing group.
// | See the [reference](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/String/replace#Specifying_a_function_as_a_parameter).
var replace$prime = /* #__PURE__ */ (function () {
    return $foreign["_replaceBy"](Data_Maybe.Just.create)(Data_Maybe.Nothing.value);
})();

// | Returns the string representation of the given `RegexFlags`.
var renderFlags = function (v) {
    return (function () {
        if (v.global) {
            return "g";
        };
        return "";
    })() + ((function () {
        if (v.ignoreCase) {
            return "i";
        };
        return "";
    })() + ((function () {
        if (v.multiline) {
            return "m";
        };
        return "";
    })() + ((function () {
        if (v.dotAll) {
            return "s";
        };
        return "";
    })() + ((function () {
        if (v.sticky) {
            return "y";
        };
        return "";
    })() + (function () {
        if (v.unicode) {
            return "u";
        };
        return "";
    })()))));
};

// | Constructs a `Regex` from a pattern string and flags. Fails with
// | `Left error` if the pattern contains a syntax error.
var regex = function (s) {
    return function (f) {
        return $foreign.regexImpl(Data_Either.Left.create)(Data_Either.Right.create)(s)(renderFlags(f));
    };
};

// | Parses the string representation of `RegexFlags`.
var parseFlags = function (s) {
    return {
        global: Data_String_CodeUnits.contains("g")(s),
        ignoreCase: Data_String_CodeUnits.contains("i")(s),
        multiline: Data_String_CodeUnits.contains("m")(s),
        dotAll: Data_String_CodeUnits.contains("s")(s),
        sticky: Data_String_CodeUnits.contains("y")(s),
        unicode: Data_String_CodeUnits.contains("u")(s)
    };
};

// | Matches the string against the `Regex` and returns an array of matches
// | if there were any. Each match has type `Maybe String`, where `Nothing`
// | represents an unmatched optional capturing group.
// | See [reference](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/String/match).
var match = /* #__PURE__ */ (function () {
    return $foreign["_match"](Data_Maybe.Just.create)(Data_Maybe.Nothing.value);
})();

// | Returns the `RegexFlags` used to construct the given `Regex`.
var flags = function ($8) {
    return Data_String_Regex_Flags.RegexFlags($foreign.flagsImpl($8));
};
export {
    source,
    test,
    replace,
    split
} from "./foreign.js";
export {
    regex,
    flags,
    renderFlags,
    parseFlags,
    match,
    replace$prime,
    search,
    showRegex
};
