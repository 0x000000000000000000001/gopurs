package Control_Comonad_Env_Class

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var _ = fmt.Println
var _ = gopurs_runtime.TypeInt

var local = gopurs_runtime.Value{}
var comonadAskTuple = gopurs_runtime.Value{}
var comonadEnvTuple = gopurs_runtime.Value{}
var comonadAskEnvT = gopurs_runtime.Value{}
var comonadEnvEnvT = gopurs_runtime.Value{}
var ask = gopurs_runtime.Value{}
var asks = gopurs_runtime.Value{}
var comonadAskStoreT = gopurs_runtime.Value{}
var comonadEnvStoreT = gopurs_runtime.Value{}
var comonadAskTracedT = gopurs_runtime.Value{}
var comonadEnvTracedT = gopurs_runtime.Value{}