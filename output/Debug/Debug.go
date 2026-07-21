package Debug

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var _ = fmt.Println
var _ = gopurs_runtime.TypeInt

var warn = gopurs_runtime.Value{}
var traceTime = gopurs_runtime.Value{}
var trace = gopurs_runtime.Value{}
var traceM = gopurs_runtime.Value{}
var spy = gopurs_runtime.Value{}
var spyWith = gopurs_runtime.Value{}
var debugger = gopurs_runtime.Value{}