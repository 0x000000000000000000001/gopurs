package Control_Monad_ST_Internal

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var _ = fmt.Println
var _ = gopurs_runtime.TypeInt

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