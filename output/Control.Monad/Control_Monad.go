package Control_Monad

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var _ = fmt.Println
var _ = gopurs_runtime.TypeInt

var whenM = gopurs_runtime.Value{}
var unlessM = gopurs_runtime.Value{}
var monadProxy = gopurs_runtime.Value{}
var monadFn = gopurs_runtime.Value{}
var monadArray = gopurs_runtime.Value{}
var liftM1 = gopurs_runtime.Value{}
var ap = gopurs_runtime.Value{}