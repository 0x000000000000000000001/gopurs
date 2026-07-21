package Effect_Class

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var _ = fmt.Println
var _ = gopurs_runtime.TypeInt

var monadEffectEffect = gopurs_runtime.Value{}
var liftEffect = gopurs_runtime.Value{}