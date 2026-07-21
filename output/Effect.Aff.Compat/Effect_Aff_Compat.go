package Effect_Aff_Compat

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var _ = fmt.Println
var _ = gopurs_runtime.TypeInt

var EffectFnCanceler = gopurs_runtime.Value{}
var EffectFnAff = gopurs_runtime.Value{}
var fromEffectFnAff = gopurs_runtime.Value{}