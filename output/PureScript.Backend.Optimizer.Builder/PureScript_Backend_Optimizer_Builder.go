package PureScript_Backend_Optimizer_Builder

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var _ = fmt.Println
var _ = gopurs_runtime.TypeInt

var insert = gopurs_runtime.Apply(insert, gopurs_runtime.Apply(ordQualified, ordString))
var buildModules = gopurs_runtime.Value{}