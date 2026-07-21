package Data_Set_NonEmpty

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var unionSet = gopurs_runtime.Value{}
var toUnfoldable1 = gopurs_runtime.Value{}
var toUnfoldable11 = gopurs_runtime.Apply(toUnfoldable1, unfoldable1Array)
var toUnfoldable12 = gopurs_runtime.Apply(toUnfoldable1, unfoldable1NonEmptyList)
var toUnfoldable = gopurs_runtime.Value{}
var toSet = gopurs_runtime.Value{}
var subset = gopurs_runtime.Value{}
var size = gopurs_runtime.Apply(unsafeCoerce, size)
var singleton = gopurs_runtime.Apply(unsafeCoerce, singleton)
var showNonEmptySet = gopurs_runtime.Value{}
var semigroupNonEmptySet = gopurs_runtime.Value{}
var properSubset = gopurs_runtime.Value{}
var ordNonEmptySet = gopurs_runtime.Value{}
var ord1NonEmptySet = ord1Set
var min = gopurs_runtime.Value{}
var member = gopurs_runtime.Value{}
var max = gopurs_runtime.Value{}
var mapMaybe = gopurs_runtime.Value{}
var map = gopurs_runtime.Value{}
var insert = gopurs_runtime.Value{}
var fromSet = gopurs_runtime.Value{}
var intersection = gopurs_runtime.Value{}
var fromFoldable1 = gopurs_runtime.Value{}
var fromFoldable = gopurs_runtime.Value{}
var foldableNonEmptySet = foldableSet
var foldable1NonEmptySet = gopurs_runtime.Value{}
var filter = gopurs_runtime.Value{}
var eqNonEmptySet = gopurs_runtime.Value{}
var eq1NonEmptySet = eq1Set
var difference = gopurs_runtime.Value{}
var delete = gopurs_runtime.Value{}
var cons = gopurs_runtime.Value{}
