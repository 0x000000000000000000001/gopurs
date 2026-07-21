package Control_Monad_Gen_Common

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var _ = fmt.Println
var _ = gopurs_runtime.TypeInt

var max = gopurs_runtime.Value{}
var genTuple = gopurs_runtime.Value{}
var genNonEmpty = gopurs_runtime.Value{}
var genMaybe' = gopurs_runtime.Value{}
var genMaybe = gopurs_runtime.Value{}
var genIdentity = gopurs_runtime.Value{}
var genEither' = gopurs_runtime.Value{}
var genEither = gopurs_runtime.Value{}