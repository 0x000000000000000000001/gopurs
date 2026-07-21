package Effect_Exception_Unsafe

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var _ = fmt.Println
var _ = gopurs_runtime.TypeInt

var unsafeThrowException = gopurs_runtime.Value{}
var unsafeThrow = gopurs_runtime.Value{}