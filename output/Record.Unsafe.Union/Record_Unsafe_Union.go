package Record_Unsafe_Union

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var _ = fmt.Println
var _ = gopurs_runtime.TypeInt

var unsafeUnion = gopurs_runtime.Apply(runFn2, unsafeUnionFn)