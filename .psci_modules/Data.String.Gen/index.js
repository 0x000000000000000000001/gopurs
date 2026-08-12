import * as Control_Bind from "../Control.Bind/index.js";
import * as Control_Monad_Gen from "../Control.Monad.Gen/index.js";
import * as Control_Monad_Gen_Class from "../Control.Monad.Gen.Class/index.js";
import * as Data_Char_Gen from "../Data.Char.Gen/index.js";
import * as Data_Function from "../Data.Function/index.js";
import * as Data_Functor from "../Data.Functor/index.js";
import * as Data_Ord from "../Data.Ord/index.js";
import * as Data_String_CodeUnits from "../Data.String.CodeUnits/index.js";
import * as Data_Unfoldable from "../Data.Unfoldable/index.js";

// | Generates a string using the specified character generator.
var genString = function (dictMonadRec) {
    return function (dictMonadGen) {
        var Monad0 = dictMonadGen.Monad0();
        var Bind1 = Monad0.Bind1();
        var Functor0 = ((Monad0.Bind1()).Apply0()).Functor0();
        return function (genChar) {
            return Control_Monad_Gen_Class.sized(dictMonadGen)(function (size) {
                return Control_Bind.bind(Bind1)(Control_Monad_Gen_Class.chooseInt(dictMonadGen)(1)(Data_Ord.max(Data_Ord.ordInt)(1)(size)))(function (newSize) {
                    return Control_Monad_Gen_Class.resize(dictMonadGen)(Data_Function["const"](newSize))(Data_Functor.map(Functor0)(Data_String_CodeUnits.fromCharArray)(Control_Monad_Gen.unfoldable(dictMonadRec)(dictMonadGen)(Data_Unfoldable.unfoldableArray)(genChar)));
                });
            });
        };
    };
};

// | Generates a string using characters from the Unicode basic multilingual
// | plain.
var genUnicodeString = function (dictMonadRec) {
    return function (dictMonadGen) {
        return genString(dictMonadRec)(dictMonadGen)(Data_Char_Gen.genUnicodeChar(dictMonadGen));
    };
};

// | Generates a string made up of numeric digits.
var genDigitString = function (dictMonadRec) {
    return function (dictMonadGen) {
        return genString(dictMonadRec)(dictMonadGen)(Data_Char_Gen.genDigitChar(dictMonadGen));
    };
};

// | Generates a string using the ASCII character set.
var genAsciiString$prime = function (dictMonadRec) {
    return function (dictMonadGen) {
        return genString(dictMonadRec)(dictMonadGen)(Data_Char_Gen["genAsciiChar$prime"](dictMonadGen));
    };
};

// | Generates a string using the ASCII character set, excluding control codes.
var genAsciiString = function (dictMonadRec) {
    return function (dictMonadGen) {
        return genString(dictMonadRec)(dictMonadGen)(Data_Char_Gen.genAsciiChar(dictMonadGen));
    };
};

// | Generates a string using uppercase characters from the basic Latin alphabet.
var genAlphaUppercaseString = function (dictMonadRec) {
    return function (dictMonadGen) {
        return genString(dictMonadRec)(dictMonadGen)(Data_Char_Gen.genAlphaUppercase(dictMonadGen));
    };
};

// | Generates a string using characters from the basic Latin alphabet.
var genAlphaString = function (dictMonadRec) {
    return function (dictMonadGen) {
        return genString(dictMonadRec)(dictMonadGen)(Data_Char_Gen.genAlpha(dictMonadGen));
    };
};

// | Generates a string using lowercase characters from the basic Latin alphabet.
var genAlphaLowercaseString = function (dictMonadRec) {
    return function (dictMonadGen) {
        return genString(dictMonadRec)(dictMonadGen)(Data_Char_Gen.genAlphaLowercase(dictMonadGen));
    };
};
export {
    genString,
    genUnicodeString,
    genAsciiString,
    genAsciiString$prime,
    genDigitString,
    genAlphaString,
    genAlphaLowercaseString,
    genAlphaUppercaseString
};
