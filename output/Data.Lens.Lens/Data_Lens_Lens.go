package Data_Lens_Lens

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var _ = fmt.Println
var _ = gopurs_runtime.TypeInt

var withLens = gopurs_runtime.Value{}
var withIndexedLens = gopurs_runtime.Value{}
var lensStore = gopurs_runtime.Value{}
var lens' = gopurs_runtime.Value{}
var lens = gopurs_runtime.Value{}
var ilens' = gopurs_runtime.Value{}
var ilens = gopurs_runtime.Value{}
var cloneLens = gopurs_runtime.Value{}
var cloneIndexedLens = gopurs_runtime.Value{}