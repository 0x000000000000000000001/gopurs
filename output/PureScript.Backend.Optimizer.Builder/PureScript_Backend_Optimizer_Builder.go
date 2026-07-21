package PureScript_Backend_Optimizer_Builder

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var insert = gopurs_runtime.Apply(insert, gopurs_runtime.Apply(ordQualified, ordString))
var buildModules = gopurs_runtime.Value{}
