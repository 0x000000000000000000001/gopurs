package Control_Parallel_Class

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var _ = fmt.Println
var _ = gopurs_runtime.TypeInt

var ParCont = gopurs_runtime.Value{}
var sequential = gopurs_runtime.Value{}
var parallel = gopurs_runtime.Value{}
var newtypeParCont = gopurs_runtime.Value{}
var monadParWriterT = gopurs_runtime.Value{}
var monadParStar = gopurs_runtime.Value{}
var monadParReaderT = gopurs_runtime.Value{}
var monadParMaybeT = gopurs_runtime.Value{}
var monadParExceptT = gopurs_runtime.Value{}
var monadParCostar = gopurs_runtime.Value{}

var monadParParCont = gopurs_runtime.Value{}
var functorParCont = gopurs_runtime.Value{}
var applyParCont = gopurs_runtime.Value{}

var applicativeParCont = gopurs_runtime.Value{}
var altParCont = gopurs_runtime.Value{}
var plusParCont = gopurs_runtime.Value{}
var alternativeParCont = gopurs_runtime.Value{}