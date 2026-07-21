package Control_Monad_RWS

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var _ = fmt.Println
var _ = gopurs_runtime.TypeInt

var withRWS = withRWST
var rws = gopurs_runtime.Value{}
var runRWS = gopurs_runtime.Value{}
var mapRWS = gopurs_runtime.Value{}
var execRWS = gopurs_runtime.Value{}
var evalRWS = gopurs_runtime.Value{}