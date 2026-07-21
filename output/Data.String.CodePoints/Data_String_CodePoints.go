package Data_String_CodePoints

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var _ = fmt.Println
var _ = gopurs_runtime.TypeInt

var unsurrogate = gopurs_runtime.Value{}
var showCodePoint = gopurs_runtime.Value{}
var isTrail = gopurs_runtime.Value{}
var isLead = gopurs_runtime.Value{}
var uncons = gopurs_runtime.Value{}
var unconsButWithTuple = gopurs_runtime.Value{}
var toCodePointArrayFallback = gopurs_runtime.Value{}
var unsafeCodePointAt0Fallback = gopurs_runtime.Value{}
var unsafeCodePointAt0 = gopurs_runtime.Apply(_unsafeCodePointAt0, unsafeCodePointAt0Fallback)
var toCodePointArray = gopurs_runtime.Apply(_toCodePointArray, toCodePointArrayFallback)
var length = gopurs_runtime.Value{}
var lastIndexOf = gopurs_runtime.Value{}
var indexOf = gopurs_runtime.Value{}
var fromCharCode = gopurs_runtime.Value{}
var singletonFallback = gopurs_runtime.Value{}
var fromCodePointArray = gopurs_runtime.Apply(_fromCodePointArray, singletonFallback)
var singleton = gopurs_runtime.Apply(_singleton, singletonFallback)

var takeFallback = gopurs_runtime.Value{}

var take = gopurs_runtime.Apply(_take, takeFallback)
var lastIndexOf' = gopurs_runtime.Value{}
var splitAt = gopurs_runtime.Value{}
var eqCodePoint = gopurs_runtime.Value{}
var ordCodePoint = gopurs_runtime.Value{}
var drop = gopurs_runtime.Value{}
var indexOf' = gopurs_runtime.Value{}

var countTail = gopurs_runtime.Value{}

var countFallback = gopurs_runtime.Value{}
var countPrefix = gopurs_runtime.Apply(_countPrefix, countFallback)
var dropWhile = gopurs_runtime.Value{}
var takeWhile = gopurs_runtime.Value{}
var codePointFromChar = gopurs_runtime.Value{}

var codePointAtFallback = gopurs_runtime.Value{}

var codePointAt = gopurs_runtime.Value{}
var boundedCodePoint = gopurs_runtime.Value{}

var boundedEnumCodePoint = gopurs_runtime.Value{}
var enumCodePoint = gopurs_runtime.Value{}