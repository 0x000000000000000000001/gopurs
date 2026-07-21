package Data_Functor

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var _ = fmt.Println
var _ = gopurs_runtime.TypeInt

var map = gopurs_runtime.Value{}
var mapFlipped = gopurs_runtime.Value{}
var void = gopurs_runtime.Value{}
var voidLeft = gopurs_runtime.Value{}
var voidRight = gopurs_runtime.Value{}
var functorProxy = gopurs_runtime.Value{}
var functorFn = gopurs_runtime.Value{}
var functorArray = gopurs_runtime.Value{}
var flap = gopurs_runtime.Value{}