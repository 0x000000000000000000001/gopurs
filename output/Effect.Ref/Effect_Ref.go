package Effect_Ref

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var _ = fmt.Println
var _ = gopurs_runtime.TypeInt

var new = _new
var modify' = modifyImpl
var modify = gopurs_runtime.Value{}
var modify_ = gopurs_runtime.Value{}