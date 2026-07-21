package Effect_Class_Console

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var _ = fmt.Println
var _ = gopurs_runtime.TypeInt

var warnShow = gopurs_runtime.Value{}
var warn = gopurs_runtime.Value{}
var timeLog = gopurs_runtime.Value{}
var timeEnd = gopurs_runtime.Value{}
var time = gopurs_runtime.Value{}
var logShow = gopurs_runtime.Value{}
var log = gopurs_runtime.Value{}
var infoShow = gopurs_runtime.Value{}
var info = gopurs_runtime.Value{}
var groupEnd = gopurs_runtime.Value{}
var groupCollapsed = gopurs_runtime.Value{}
var group = gopurs_runtime.Value{}
var grouped = gopurs_runtime.Value{}
var errorShow = gopurs_runtime.Value{}
var error = gopurs_runtime.Value{}
var debugShow = gopurs_runtime.Value{}
var debug = gopurs_runtime.Value{}
var clear = gopurs_runtime.Value{}