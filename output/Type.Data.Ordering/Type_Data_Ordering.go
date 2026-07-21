package Type_Data_Ordering

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var _ = fmt.Println
var _ = gopurs_runtime.TypeInt

var reflectOrdering = gopurs_runtime.Value{}
var isOrderingLT = gopurs_runtime.Value{}
var isOrderingGT = gopurs_runtime.Value{}
var isOrderingEQ = gopurs_runtime.Value{}
var reifyOrdering = gopurs_runtime.Value{}
var invertOrderingLT = gopurs_runtime.Value{}
var invertOrderingGT = gopurs_runtime.Value{}
var invertOrderingEQ = gopurs_runtime.Value{}
var invert = gopurs_runtime.Value{}
var equalsLTLT = gopurs_runtime.Value{}
var equalsLTGT = gopurs_runtime.Value{}
var equalsLTEQ = gopurs_runtime.Value{}
var equalsGTLT = gopurs_runtime.Value{}
var equalsGTGT = gopurs_runtime.Value{}
var equalsGTEQ = gopurs_runtime.Value{}
var equalsEQLT = gopurs_runtime.Value{}
var equalsEQGT = gopurs_runtime.Value{}
var equalsEQEQ = gopurs_runtime.Value{}
var equals = gopurs_runtime.Value{}
var appendOrderingLT = gopurs_runtime.Value{}
var appendOrderingGT = gopurs_runtime.Value{}
var appendOrderingEQ = gopurs_runtime.Value{}
var append = gopurs_runtime.Value{}