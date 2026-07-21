package Foreign_Object_Gen

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var fromFoldable = gopurs_runtime.Apply(fromFoldable, foldableList)
var genForeignObject = gopurs_runtime.Value{}
