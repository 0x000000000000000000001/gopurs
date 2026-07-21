package Data_Array_Partial

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var unsafeIndex = gopurs_runtime.Apply(runFn2, unsafeIndexImpl)
var tail = gopurs_runtime.Value{}
var last = gopurs_runtime.Value{}
var init = gopurs_runtime.Value{}
var head = gopurs_runtime.Value{}
