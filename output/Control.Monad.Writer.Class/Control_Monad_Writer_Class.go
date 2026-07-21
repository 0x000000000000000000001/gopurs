package Control_Monad_Writer_Class

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var _ = fmt.Println
var _ = gopurs_runtime.TypeInt

var tell = gopurs_runtime.Value{}
var pass = gopurs_runtime.Value{}
var listen = gopurs_runtime.Value{}
var listens = gopurs_runtime.Value{}
var censor = gopurs_runtime.Value{}