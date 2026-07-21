package PureScript_CST_ModuleGraph

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var _ = fmt.Println
var _ = gopurs_runtime.TypeInt

var toUnfoldable = gopurs_runtime.Value{}
var toUnfoldable1 = gopurs_runtime.Apply(toUnfoldable, unfoldableArray)
var all = gopurs_runtime.Apply(all, foldableMap)
var fromFoldable = gopurs_runtime.Apply(foldlArray, gopurs_runtime.Value{})
var fromFoldable1 = gopurs_runtime.Apply(fromFoldable, ordString)
var lookup = gopurs_runtime.Value{}
var toUnfoldable2 = gopurs_runtime.Apply(gopurs_runtime.Value{}, gopurs_runtime.Value{})
var Sorted = gopurs_runtime.Value{}
var CycleDetected = gopurs_runtime.Value{}
var topoSort = gopurs_runtime.Value{}
var topoSort1 = gopurs_runtime.Apply(topoSort, ordString)
var moduleGraph = gopurs_runtime.Value{}
var sortModules = gopurs_runtime.Value{}