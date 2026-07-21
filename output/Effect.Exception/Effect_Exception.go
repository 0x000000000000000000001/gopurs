package Effect_Exception

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var _ = fmt.Println
var _ = gopurs_runtime.TypeInt

var try = gopurs_runtime.Value{}
var throw = gopurs_runtime.Value{}
var stack = gopurs_runtime.Apply(stackImpl, Just)
var showError = gopurs_runtime.Value{}