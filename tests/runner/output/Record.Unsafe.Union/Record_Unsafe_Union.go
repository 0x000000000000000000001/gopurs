package Record_Unsafe_Union

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var unsafeUnionFn = gopurs_runtime.Value{}
var unsafeUnion = gopurs_runtime.Apply(runFn2, unsafeUnionFn)
