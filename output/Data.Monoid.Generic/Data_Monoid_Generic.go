package Data_Monoid_Generic

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var _ = fmt.Println
var _ = gopurs_runtime.TypeInt

var genericMonoidNoArguments = gopurs_runtime.Value{}
var genericMonoidArgument = gopurs_runtime.Value{}
var genericMempty' = gopurs_runtime.Value{}
var genericMonoidConstructor = gopurs_runtime.Value{}
var genericMonoidProduct = gopurs_runtime.Value{}
var genericMempty = gopurs_runtime.Value{}