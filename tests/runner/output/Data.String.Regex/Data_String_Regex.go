package Data_String_Regex

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var _match = gopurs_runtime.Value{}
var _replaceBy = gopurs_runtime.Value{}
var _search = gopurs_runtime.Value{}
var flagsImpl = gopurs_runtime.Value{}
var regexImpl = gopurs_runtime.Value{}
var replace = gopurs_runtime.Value{}
var showRegexImpl = gopurs_runtime.Value{}
var source = gopurs_runtime.Value{}
var split = gopurs_runtime.Value{}
var test = gopurs_runtime.Value{}
var showRegex = gopurs_runtime.Value{}
var search = gopurs_runtime.Apply(_search, Just)
var replace' = gopurs_runtime.Apply(_replaceBy, Just)
var renderFlags = gopurs_runtime.Value{}
var regex = gopurs_runtime.Value{}
var parseFlags = gopurs_runtime.Value{}
var match = gopurs_runtime.Apply(_match, Just)
var flags = gopurs_runtime.Value{}
