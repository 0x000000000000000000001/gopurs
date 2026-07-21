package Data_Ordering

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var _ = fmt.Println
var _ = gopurs_runtime.TypeInt

var LT = gopurs_runtime.Value{}
var GT = gopurs_runtime.Value{}
var EQ = gopurs_runtime.Value{}
var showOrdering = gopurs_runtime.Value{}
var semigroupOrdering = gopurs_runtime.Value{}
var invert = gopurs_runtime.Value{}
var eqOrdering = gopurs_runtime.Value{}