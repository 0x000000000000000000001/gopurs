package PureScript_Backend_Optimizer_CoreFn_Sort

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var lookup = gopurs_runtime.Value{}
var append = gopurs_runtime.Apply(union, ordString)
var delete = gopurs_runtime.Apply(unsafeCoerce, gopurs_runtime.Apply(delete, ordString))
var fromFoldable = gopurs_runtime.Apply(foldlArray, gopurs_runtime.Value{})
var guard = gopurs_runtime.Apply(guard, alternativeMaybe)
var member = gopurs_runtime.Value{}
var runSort = gopurs_runtime.Value{}
var sortModules = gopurs_runtime.Value{}
var resumePull = gopurs_runtime.Value{}
var pullResult = gopurs_runtime.Value{}
var emptyPull = gopurs_runtime.Value{}
