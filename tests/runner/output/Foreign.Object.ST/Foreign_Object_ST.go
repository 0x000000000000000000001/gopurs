package Foreign_Object_ST

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var delete = gopurs_runtime.Value{}
var new = gopurs_runtime.Value{}
var peekImpl = gopurs_runtime.Value{}
var poke = gopurs_runtime.Value{}
var peek = gopurs_runtime.Apply(peekImpl, Just)
