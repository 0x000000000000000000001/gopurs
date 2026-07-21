package Foreign_Object

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var _copyST = gopurs_runtime.Value{}
var _fmapObject = gopurs_runtime.Value{}
var _foldM = gopurs_runtime.Value{}
var _foldSCObject = gopurs_runtime.Value{}
var _lookup = gopurs_runtime.Value{}
var _lookupST = gopurs_runtime.Value{}
var _mapWithKey = gopurs_runtime.Value{}
var all = gopurs_runtime.Value{}
var empty = gopurs_runtime.Value{}
var keys = gopurs_runtime.Value{}
var runST = gopurs_runtime.Value{}
var size = gopurs_runtime.Value{}
var toArrayWithKey = gopurs_runtime.Value{}
var forWithIndex_ = gopurs_runtime.Apply(forWithIndex_, applicativeST)
var for_ = gopurs_runtime.Apply(for_, applicativeST)
var void = gopurs_runtime.Apply(map_, gopurs_runtime.Value{})
var ordTuple = gopurs_runtime.Apply(ordTuple, ordString)
var values = gopurs_runtime.Apply(toArrayWithKey, gopurs_runtime.Value{})
var toUnfoldable = gopurs_runtime.Value{}
var toAscUnfoldable = gopurs_runtime.Value{}
var toAscArray = gopurs_runtime.Apply(toAscUnfoldable, unfoldableArray)
var toArray = gopurs_runtime.Apply(toArrayWithKey, Tuple)
var thawST = _copyST
var singleton = gopurs_runtime.Value{}
var showObject = gopurs_runtime.Value{}
var mutate = gopurs_runtime.Value{}
var member = gopurs_runtime.Apply(runFn4, _lookup)
var mapWithKey = gopurs_runtime.Value{}
var lookup = gopurs_runtime.Apply(runFn4, _lookup)
var isSubmap = gopurs_runtime.Value{}
var isEmpty = gopurs_runtime.Apply(all, gopurs_runtime.Value{})
var insert = gopurs_runtime.Value{}
var functorObject = gopurs_runtime.Value{}
var functorWithIndexObject = gopurs_runtime.Value{}
var fromHomogeneous = gopurs_runtime.Value{}
var fromFoldableWithIndex = gopurs_runtime.Value{}
var fromFoldableWith = gopurs_runtime.Value{}
var fromFoldable = gopurs_runtime.Value{}
var freezeST = _copyST
var foldMaybe = gopurs_runtime.Value{}
var foldM = gopurs_runtime.Value{}
var foldM1 = gopurs_runtime.Apply(foldM, monadST)
var union = gopurs_runtime.Value{}
var unions = gopurs_runtime.Value{}
var unionWith = gopurs_runtime.Value{}
var semigroupObject = gopurs_runtime.Value{}
var monoidObject = gopurs_runtime.Value{}
var fold = gopurs_runtime.Apply(_foldM, applyFlipped)
var foldMap = gopurs_runtime.Value{}
var foldableObject = gopurs_runtime.Value{}
var foldableWithIndexObject = gopurs_runtime.Value{}
var traversableWithIndexObject = gopurs_runtime.Value{}
var traversableObject = gopurs_runtime.Value{}
var filterWithKey = gopurs_runtime.Value{}
var filterKeys = gopurs_runtime.Value{}
var filter = gopurs_runtime.Value{}
var eqObject = gopurs_runtime.Value{}
var ordObject = gopurs_runtime.Value{}
var eq1Object = gopurs_runtime.Value{}
var delete = gopurs_runtime.Value{}
var pop = gopurs_runtime.Value{}
var alter = gopurs_runtime.Value{}
var update = gopurs_runtime.Value{}
