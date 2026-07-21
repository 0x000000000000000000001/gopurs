package Data_Number_Approximate

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var Tolerance = gopurs_runtime.Value{}
var Fraction = gopurs_runtime.Value{}
var eqRelative = gopurs_runtime.Value{}
var eqApproximate = gopurs_runtime.Apply(eqRelative, gopurs_runtime.Value{})
var neqApproximate = gopurs_runtime.Value{}
var eqAbsolute = gopurs_runtime.Value{}
