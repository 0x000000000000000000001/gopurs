package Control_Monad_State

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var withState = withStateT
var runState = gopurs_runtime.Value{}
var mapState = gopurs_runtime.Value{}
var execState = gopurs_runtime.Value{}
var evalState = gopurs_runtime.Value{}
