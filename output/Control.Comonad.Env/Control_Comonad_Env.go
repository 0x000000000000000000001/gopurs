package Control_Comonad_Env

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var withEnv = withEnvT
var runEnv = gopurs_runtime.Value{}
var mapEnv = gopurs_runtime.Value{}
var env = gopurs_runtime.Value{}
