package Control_Monad_Trampoline

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var runTrampoline = gopurs_runtime.Apply(runFree, functorFn)
var done = gopurs_runtime.Value{}
var delay = liftF
