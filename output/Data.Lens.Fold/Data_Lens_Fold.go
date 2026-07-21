package Data_Lens_Fold

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var fromFoldable = gopurs_runtime.Apply(runFn2, fromFoldableImpl)
var unfolded = gopurs_runtime.Value{}
var replicated = gopurs_runtime.Value{}
var ifoldMapOf = gopurs_runtime.Value{}
var ifoldlOf = gopurs_runtime.Value{}
var ifoldrOf = gopurs_runtime.Value{}
var itoListOf = gopurs_runtime.Value{}
var itraverseOf_ = gopurs_runtime.Value{}
var iforOf_ = gopurs_runtime.Value{}
var ifindOf = gopurs_runtime.Value{}
var ianyOf = gopurs_runtime.Value{}
var iallOf = gopurs_runtime.Value{}
var folded = gopurs_runtime.Value{}
var foldMapOf = unsafeCoerce
var foldOf = gopurs_runtime.Value{}
var foldlOf = gopurs_runtime.Value{}
var foldrOf = gopurs_runtime.Value{}
var maximumOf = gopurs_runtime.Value{}
var minimumOf = gopurs_runtime.Value{}
var toListOf = gopurs_runtime.Value{}
var toArrayOf = gopurs_runtime.Value{}
var toArrayOfOn = gopurs_runtime.Value{}
var toListOfOn = gopurs_runtime.Value{}
var traverseOf_ = gopurs_runtime.Value{}
var has = gopurs_runtime.Value{}
var hasn't = gopurs_runtime.Value{}
var lastOf = gopurs_runtime.Value{}
var lengthOf = gopurs_runtime.Value{}
var preview = gopurs_runtime.Value{}
var previewOn = gopurs_runtime.Value{}
var productOf = gopurs_runtime.Value{}
var sequenceOf_ = gopurs_runtime.Value{}
var sumOf = gopurs_runtime.Value{}
var firstOf = gopurs_runtime.Value{}
var findOf = gopurs_runtime.Value{}
var filtered = gopurs_runtime.Value{}
var anyOf = gopurs_runtime.Value{}
var elemOf = gopurs_runtime.Value{}
var orOf = gopurs_runtime.Value{}
var allOf = gopurs_runtime.Value{}
var andOf = gopurs_runtime.Value{}
var notElemOf = gopurs_runtime.Value{}
