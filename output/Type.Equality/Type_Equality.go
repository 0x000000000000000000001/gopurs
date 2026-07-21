package Type_Equality

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var _ = fmt.Println
var _ = gopurs_runtime.TypeInt

var refl = gopurs_runtime.Value{}
var proof = gopurs_runtime.Value{}
var to = gopurs_runtime.Value{}
var from = gopurs_runtime.Value{}