package Control_Monad_Free

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var _ = fmt.Println
var _ = gopurs_runtime.TypeInt

var Free = gopurs_runtime.Value{}
var Return = gopurs_runtime.Value{}
var Bind = gopurs_runtime.Value{}

var toView = gopurs_runtime.Value{}

var runFreeM = gopurs_runtime.Value{}
var runFree = gopurs_runtime.Value{}
var resume' = gopurs_runtime.Value{}
var resume = gopurs_runtime.Value{}
var wrap = gopurs_runtime.Value{}
var suspendF = gopurs_runtime.Value{}

var freeMonad = gopurs_runtime.Value{}
var freeFunctor = gopurs_runtime.Value{}
var freeBind = gopurs_runtime.Value{}
var freeApply = gopurs_runtime.Value{}
var freeApplicative = gopurs_runtime.Value{}

var lift2 = gopurs_runtime.Value{}
var semigroupFree = gopurs_runtime.Value{}

var freeMonadRec = gopurs_runtime.Value{}

var liftF = gopurs_runtime.Value{}
var freeMonadTrans = gopurs_runtime.Value{}
var monoidFree = gopurs_runtime.Value{}
var substFree = gopurs_runtime.Value{}
var hoistFree = gopurs_runtime.Value{}
var foldableFree = gopurs_runtime.Value{}

var traversableFree = gopurs_runtime.Value{}

var foldFree = gopurs_runtime.Value{}

var eqFree = gopurs_runtime.Value{}

var ordFree = gopurs_runtime.Value{}

var eq1Free = gopurs_runtime.Value{}
var ord1Free = gopurs_runtime.Value{}