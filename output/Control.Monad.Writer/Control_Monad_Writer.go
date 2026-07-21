package Control_Monad_Writer

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var _ = fmt.Println
var _ = gopurs_runtime.TypeInt

var writer = gopurs_runtime.Value{}
var runWriter = gopurs_runtime.Value{}
var mapWriter = gopurs_runtime.Value{}
var execWriter = gopurs_runtime.Value{}