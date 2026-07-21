package Data_Argonaut_Decode_Error

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var _ = fmt.Println
var _ = gopurs_runtime.TypeInt

var TypeMismatch = gopurs_runtime.Value{}
var UnexpectedValue = gopurs_runtime.Value{}
var AtIndex = gopurs_runtime.Value{}
var AtKey = gopurs_runtime.Value{}
var Named = gopurs_runtime.Value{}
var MissingValue = gopurs_runtime.Value{}

var showJsonDecodeError = gopurs_runtime.Value{}

var printJsonDecodeError = gopurs_runtime.Value{}
var genericJsonDecodeError = gopurs_runtime.Value{}

var eqJsonDecodeError = gopurs_runtime.Value{}

var ordJsonDecodeError = gopurs_runtime.Value{}