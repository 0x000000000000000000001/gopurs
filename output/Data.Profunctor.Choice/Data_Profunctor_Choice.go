package Data_Profunctor_Choice

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var _ = fmt.Println
var _ = gopurs_runtime.TypeInt

var right = gopurs_runtime.Value{}
var left = gopurs_runtime.Value{}
var splitChoice = gopurs_runtime.Value{}
var fanin = gopurs_runtime.Value{}
var choiceFn = gopurs_runtime.Value{}