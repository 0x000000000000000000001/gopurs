package Data_Reflectable

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var _ = fmt.Println
var _ = gopurs_runtime.TypeInt

var reifiableString = gopurs_runtime.Value{}
var reifiableOrdering = gopurs_runtime.Value{}
var reifiableInt = gopurs_runtime.Value{}
var reifiableBoolean = gopurs_runtime.Value{}
var reifyType = gopurs_runtime.Value{}
var reflectType = gopurs_runtime.Value{}