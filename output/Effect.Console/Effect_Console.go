package Effect_Console

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var _ = fmt.Println
var _ = gopurs_runtime.TypeInt

var warnShow = gopurs_runtime.Value{}
var logShow = gopurs_runtime.Value{}
var infoShow = gopurs_runtime.Value{}
var grouped = gopurs_runtime.Value{}
var errorShow = gopurs_runtime.Value{}
var debugShow = gopurs_runtime.Value{}