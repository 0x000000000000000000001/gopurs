package Control_Monad_State_Class

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var _ = fmt.Println
var _ = gopurs_runtime.TypeInt

var state = gopurs_runtime.Value{}
var put = gopurs_runtime.Value{}
var modify_ = gopurs_runtime.Value{}
var modify = gopurs_runtime.Value{}
var gets = gopurs_runtime.Value{}
var get = gopurs_runtime.Value{}