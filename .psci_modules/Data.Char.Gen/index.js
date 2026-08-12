import * as Control_Monad_Gen from "../Control.Monad.Gen/index.js";
import * as Control_Monad_Gen_Class from "../Control.Monad.Gen.Class/index.js";
import * as Data_Bounded from "../Data.Bounded/index.js";
import * as Data_Enum from "../Data.Enum/index.js";
import * as Data_Foldable from "../Data.Foldable/index.js";
import * as Data_Functor from "../Data.Functor/index.js";
import * as Data_NonEmpty from "../Data.NonEmpty/index.js";
var foldable1NonEmpty = /* #__PURE__ */ Data_NonEmpty.foldable1NonEmpty(Data_Foldable.foldableArray);

// | Generates a character of the Unicode basic multilingual plane.
var genUnicodeChar = function (dictMonadGen) {
    return Data_Functor.map((((dictMonadGen.Monad0()).Bind1()).Apply0()).Functor0())(Data_Enum.toEnumWithDefaults(Data_Enum.boundedEnumChar)(Data_Bounded.bottom(Data_Bounded.boundedChar))(Data_Bounded.top(Data_Bounded.boundedChar)))(Control_Monad_Gen_Class.chooseInt(dictMonadGen)(0)(65536));
};

// | Generates a character that is a numeric digit.
var genDigitChar = function (dictMonadGen) {
    return Data_Functor.map((((dictMonadGen.Monad0()).Bind1()).Apply0()).Functor0())(Data_Enum.toEnumWithDefaults(Data_Enum.boundedEnumChar)(Data_Bounded.bottom(Data_Bounded.boundedChar))(Data_Bounded.top(Data_Bounded.boundedChar)))(Control_Monad_Gen_Class.chooseInt(dictMonadGen)(48)(57));
};

// | Generates a character in the ASCII character set.
var genAsciiChar$prime = function (dictMonadGen) {
    return Data_Functor.map((((dictMonadGen.Monad0()).Bind1()).Apply0()).Functor0())(Data_Enum.toEnumWithDefaults(Data_Enum.boundedEnumChar)(Data_Bounded.bottom(Data_Bounded.boundedChar))(Data_Bounded.top(Data_Bounded.boundedChar)))(Control_Monad_Gen_Class.chooseInt(dictMonadGen)(0)(127));
};

// | Generates a character in the ASCII character set, excluding control codes.
var genAsciiChar = function (dictMonadGen) {
    return Data_Functor.map((((dictMonadGen.Monad0()).Bind1()).Apply0()).Functor0())(Data_Enum.toEnumWithDefaults(Data_Enum.boundedEnumChar)(Data_Bounded.bottom(Data_Bounded.boundedChar))(Data_Bounded.top(Data_Bounded.boundedChar)))(Control_Monad_Gen_Class.chooseInt(dictMonadGen)(32)(127));
};

// | Generates an uppercase character from the basic latin alphabet.
var genAlphaUppercase = function (dictMonadGen) {
    return Data_Functor.map((((dictMonadGen.Monad0()).Bind1()).Apply0()).Functor0())(Data_Enum.toEnumWithDefaults(Data_Enum.boundedEnumChar)(Data_Bounded.bottom(Data_Bounded.boundedChar))(Data_Bounded.top(Data_Bounded.boundedChar)))(Control_Monad_Gen_Class.chooseInt(dictMonadGen)(65)(90));
};

// | Generates a lowercase character from the basic latin alphabet.
var genAlphaLowercase = function (dictMonadGen) {
    return Data_Functor.map((((dictMonadGen.Monad0()).Bind1()).Apply0()).Functor0())(Data_Enum.toEnumWithDefaults(Data_Enum.boundedEnumChar)(Data_Bounded.bottom(Data_Bounded.boundedChar))(Data_Bounded.top(Data_Bounded.boundedChar)))(Control_Monad_Gen_Class.chooseInt(dictMonadGen)(97)(122));
};

// | Generates a character from the basic latin alphabet.
var genAlpha = function (dictMonadGen) {
    return Control_Monad_Gen.oneOf(dictMonadGen)(foldable1NonEmpty)(new Data_NonEmpty.NonEmpty(genAlphaLowercase(dictMonadGen), [ genAlphaUppercase(dictMonadGen) ]));
};
export {
    genUnicodeChar,
    genAsciiChar,
    genAsciiChar$prime,
    genDigitChar,
    genAlpha,
    genAlphaLowercase,
    genAlphaUppercase
};
