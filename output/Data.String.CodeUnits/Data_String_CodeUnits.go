package Data_String_CodeUnits

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var _ = fmt.Println
var _ = gopurs_runtime.TypeInt

var uncons = gopurs_runtime.Value{}
var toChar = gopurs_runtime.Apply(_toChar, Just)
var takeWhile = gopurs_runtime.Value{}
var takeRight = gopurs_runtime.Value{}
var stripSuffix = gopurs_runtime.Value{}
var stripPrefix = gopurs_runtime.Value{}
var lastIndexOf' = gopurs_runtime.Apply(_lastIndexOfStartingAt, Just)
var lastIndexOf = gopurs_runtime.Apply(_lastIndexOf, Just)
var indexOf' = gopurs_runtime.Apply(_indexOfStartingAt, Just)
var indexOf = gopurs_runtime.Apply(_indexOf, Just)
var dropWhile = gopurs_runtime.Value{}
var dropRight = gopurs_runtime.Value{}
var contains = gopurs_runtime.Value{}
var charAt = gopurs_runtime.Apply(_charAt, Just)