package Foreign_Object_ST

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var _ = fmt.Println
var _ = gopurs_runtime.TypeInt

var peek = gopurs_runtime.Apply(peekImpl, Just)