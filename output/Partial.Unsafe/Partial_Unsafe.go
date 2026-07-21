package Partial_Unsafe

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var _ = fmt.Println
var _ = gopurs_runtime.TypeInt

var unsafePartial = _unsafePartial
var unsafeCrashWith = gopurs_runtime.Value{}