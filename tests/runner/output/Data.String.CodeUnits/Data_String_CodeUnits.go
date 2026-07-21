package Data_String_CodeUnits

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var _charAt = gopurs_runtime.Value{}
var _indexOf = gopurs_runtime.Value{}
var _indexOfStartingAt = gopurs_runtime.Value{}
var _lastIndexOf = gopurs_runtime.Value{}
var _lastIndexOfStartingAt = gopurs_runtime.Value{}
var _toChar = gopurs_runtime.Value{}
var countPrefix = gopurs_runtime.Value{}
var drop = gopurs_runtime.Value{}
var fromCharArray = gopurs_runtime.Value{}
var length = gopurs_runtime.Value{}
var singleton = gopurs_runtime.Value{}
var slice = gopurs_runtime.Value{}
var splitAt = gopurs_runtime.Value{}
var take = gopurs_runtime.Value{}
var toCharArray = gopurs_runtime.Value{}
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
