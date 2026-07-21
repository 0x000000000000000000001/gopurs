package PureScript_Backend_Optimizer_Debug

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var _ = fmt.Println
var _ = gopurs_runtime.TypeInt

var traceWhen = gopurs_runtime.Value{}
var time = gopurs_runtime.Value{}
var spyWhen = gopurs_runtime.Value{}