package Data_Semigroup

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var _ = fmt.Println
var _ = gopurs_runtime.TypeInt

var semigroupVoid = gopurs_runtime.Value{}
var semigroupUnit = gopurs_runtime.Value{}
var semigroupString = gopurs_runtime.Value{}
var semigroupRecordNil = gopurs_runtime.Value{}
var semigroupProxy = gopurs_runtime.Value{}
var semigroupArray = gopurs_runtime.Value{}
var appendRecord = gopurs_runtime.Value{}
var semigroupRecord = gopurs_runtime.Value{}
var append = gopurs_runtime.Value{}
var semigroupFn = gopurs_runtime.Value{}
var semigroupRecordCons = gopurs_runtime.Value{}