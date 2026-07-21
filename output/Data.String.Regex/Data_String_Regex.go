package Data_String_Regex

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var _ = fmt.Println
var _ = gopurs_runtime.TypeInt

var showRegex = gopurs_runtime.Value{}
var search = gopurs_runtime.Apply(_search, Just)
var replace' = gopurs_runtime.Apply(_replaceBy, Just)
var renderFlags = gopurs_runtime.Value{}
var regex = gopurs_runtime.Value{}
var parseFlags = gopurs_runtime.Value{}
var match = gopurs_runtime.Apply(_match, Just)
var flags = gopurs_runtime.Value{}