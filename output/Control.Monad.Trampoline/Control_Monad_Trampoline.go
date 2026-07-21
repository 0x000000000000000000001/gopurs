package Control_Monad_Trampoline

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var _ = fmt.Println
var _ = gopurs_runtime.TypeInt

var runTrampoline = gopurs_runtime.Apply(runFree, functorFn)
var done = gopurs_runtime.Value{}
var delay = liftF