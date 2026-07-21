package Data_Lens_Traversal

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var _ = fmt.Println
var _ = gopurs_runtime.TypeInt

var traversed = gopurs_runtime.Value{}
var traverseOf = unsafeCoerce
var sequenceOf = gopurs_runtime.Value{}
var itraverseOf = gopurs_runtime.Value{}
var iforOf = gopurs_runtime.Value{}
var failover = gopurs_runtime.Value{}
var elementsOf = gopurs_runtime.Value{}
var element = gopurs_runtime.Value{}
var cloneTraversal = gopurs_runtime.Value{}
var both = gopurs_runtime.Value{}