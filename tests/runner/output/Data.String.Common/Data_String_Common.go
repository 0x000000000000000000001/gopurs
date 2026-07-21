package Data_String_Common

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var _localeCompare = gopurs_runtime.Value{}
var joinWith = gopurs_runtime.Value{}
var replace = gopurs_runtime.Value{}
var replaceAll = gopurs_runtime.Value{}
var split = gopurs_runtime.Value{}
var toLower = gopurs_runtime.Value{}
var toUpper = gopurs_runtime.Value{}
var trim = gopurs_runtime.Value{}
var null = gopurs_runtime.Value{}
var localeCompare = gopurs_runtime.Apply(_localeCompare, gopurs_runtime.Value{})
