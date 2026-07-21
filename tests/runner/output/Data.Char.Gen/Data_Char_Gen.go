package Data_Char_Gen

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var toEnumWithDefaults = gopurs_runtime.Apply(toEnumWithDefaults, boundedEnumChar)
var foldable1NonEmpty = gopurs_runtime.Apply(foldable1NonEmpty, foldableArray)
var genUnicodeChar = gopurs_runtime.Value{}
var genDigitChar = gopurs_runtime.Value{}
var genAsciiChar' = gopurs_runtime.Value{}
var genAsciiChar = gopurs_runtime.Value{}
var genAlphaUppercase = gopurs_runtime.Value{}
var genAlphaLowercase = gopurs_runtime.Value{}
var genAlpha = gopurs_runtime.Value{}
