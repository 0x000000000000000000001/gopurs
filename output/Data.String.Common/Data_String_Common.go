package Data_String_Common

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var _ = fmt.Println
var _ = gopurs_runtime.TypeInt

var null = gopurs_runtime.Value{}
var localeCompare = gopurs_runtime.Apply(_localeCompare, gopurs_runtime.Value{})