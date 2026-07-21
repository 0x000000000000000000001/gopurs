package Effect_Exception

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var catchException = gopurs_runtime.Value{}
var error = gopurs_runtime.Value{}
var errorWithCause = gopurs_runtime.Value{}
var errorWithName = gopurs_runtime.Value{}
var message = gopurs_runtime.Value{}
var name = gopurs_runtime.Value{}
var showErrorImpl = gopurs_runtime.Value{}
var stackImpl = gopurs_runtime.Value{}
var throwException = gopurs_runtime.Value{}
var try = gopurs_runtime.Value{}
var throw = gopurs_runtime.Value{}
var stack = gopurs_runtime.Apply(stackImpl, Just)
var showError = gopurs_runtime.Value{}
