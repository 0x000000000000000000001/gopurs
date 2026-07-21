package Control_Monad_Error_Class

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var _ = fmt.Println
var _ = gopurs_runtime.TypeInt

var throwError = gopurs_runtime.Value{}
var monadThrowMaybe = gopurs_runtime.Value{}
var monadThrowEither = gopurs_runtime.Value{}
var monadThrowEffect = gopurs_runtime.Value{}
var monadErrorMaybe = gopurs_runtime.Value{}
var monadErrorEither = gopurs_runtime.Value{}
var monadErrorEffect = gopurs_runtime.Value{}
var liftMaybe = gopurs_runtime.Value{}
var liftEither = gopurs_runtime.Value{}
var catchError = gopurs_runtime.Value{}
var catchJust = gopurs_runtime.Value{}
var try = gopurs_runtime.Value{}
var withResource = gopurs_runtime.Value{}