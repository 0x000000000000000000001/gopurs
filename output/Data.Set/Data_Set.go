package Data_Set

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var _ = fmt.Println
var _ = gopurs_runtime.TypeInt

var Set = gopurs_runtime.Value{}
var union = gopurs_runtime.Value{}
var toggle = gopurs_runtime.Value{}
var toMap = gopurs_runtime.Value{}
var toUnfoldable = gopurs_runtime.Value{}
var toUnfoldable1 = gopurs_runtime.Apply(toUnfoldable, unfoldableArray)
var size = gopurs_runtime.Apply(unsafeCoerce, size)
var singleton = gopurs_runtime.Value{}
var showSet = gopurs_runtime.Value{}
var semigroupSet = gopurs_runtime.Value{}
var member = gopurs_runtime.Value{}
var isEmpty = gopurs_runtime.Apply(unsafeCoerce, isEmpty)
var intersection = gopurs_runtime.Value{}
var insert = gopurs_runtime.Value{}
var fromMap = Set
var foldableSet = gopurs_runtime.Value{}
var findMin = gopurs_runtime.Value{}
var findMax = gopurs_runtime.Value{}
var filter = gopurs_runtime.Value{}
var eqSet = gopurs_runtime.Value{}
var ordSet = gopurs_runtime.Value{}
var eq1Set = gopurs_runtime.Value{}
var ord1Set = gopurs_runtime.Value{}
var empty = gopurs_runtime.Value{}
var fromFoldable = gopurs_runtime.Value{}
var map = gopurs_runtime.Value{}
var mapMaybe = gopurs_runtime.Value{}
var monoidSet = gopurs_runtime.Value{}
var unions = gopurs_runtime.Value{}
var difference = gopurs_runtime.Value{}
var subset = gopurs_runtime.Value{}
var properSubset = gopurs_runtime.Value{}
var delete = gopurs_runtime.Value{}
var checkValid = gopurs_runtime.Value{}
var catMaybes = gopurs_runtime.Value{}