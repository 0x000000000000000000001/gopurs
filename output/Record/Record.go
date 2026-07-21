package Record

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var _ = fmt.Println
var _ = gopurs_runtime.TypeInt

var union = gopurs_runtime.Value{}
var set = gopurs_runtime.Value{}
var nub = gopurs_runtime.Value{}
var merge = gopurs_runtime.Value{}
var insert = gopurs_runtime.Value{}
var get = gopurs_runtime.Value{}
var modify = gopurs_runtime.Value{}
var equalFieldsNil = gopurs_runtime.Value{}
var equalFields = gopurs_runtime.Value{}
var equalFieldsCons = gopurs_runtime.Value{}
var equal = gopurs_runtime.Value{}
var disjointUnion = gopurs_runtime.Value{}
var delete = gopurs_runtime.Value{}
var rename = gopurs_runtime.Value{}