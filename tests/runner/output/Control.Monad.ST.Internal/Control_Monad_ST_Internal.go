package Control_Monad_ST_Internal

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var bind_ = gopurs_runtime.Value{}
var for = gopurs_runtime.Value{}
var foreach = gopurs_runtime.Value{}
var map_ = gopurs_runtime.Value{}
var modifyImpl = gopurs_runtime.Value{}
var new = gopurs_runtime.Value{}
var pure_ = gopurs_runtime.Value{}
var read = gopurs_runtime.Value{}
var run = gopurs_runtime.Value{}
var while = gopurs_runtime.Value{}
var write = gopurs_runtime.Value{}
var modify' = modifyImpl
var modify = gopurs_runtime.Value{}
var functorST = gopurs_runtime.Value{}
var void = gopurs_runtime.Apply(map_, gopurs_runtime.Value{})
var monadST = gopurs_runtime.Value{}
var bindST = gopurs_runtime.Value{}
var applyST = gopurs_runtime.Value{}
var applicativeST = gopurs_runtime.Value{}
var lift2 = gopurs_runtime.Value{}
var semigroupST = gopurs_runtime.Value{}
var monadRecST = gopurs_runtime.Value{}
var monoidST = gopurs_runtime.Value{}
