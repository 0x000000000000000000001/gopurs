package Effect_Aff

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var _ = fmt.Println
var _ = gopurs_runtime.TypeInt

var Canceler = gopurs_runtime.Value{}
var suspendAff = gopurs_runtime.Apply(_fork, gopurs_runtime.Value{})
var newtypeCanceler = gopurs_runtime.Value{}
var functorParAff = gopurs_runtime.Value{}
var functorAff = gopurs_runtime.Value{}
var forkAff = gopurs_runtime.Apply(_fork, gopurs_runtime.Value{})
var ffiUtil = gopurs_runtime.Value{}
var makeFiber = gopurs_runtime.Value{}
var launchAff = gopurs_runtime.Value{}
var launchAff_ = gopurs_runtime.Value{}
var launchSuspendedAff = makeFiber
var delay = gopurs_runtime.Value{}
var bracket = gopurs_runtime.Value{}
var applyParAff = gopurs_runtime.Value{}
var lift2 = gopurs_runtime.Value{}
var semigroupParAff = gopurs_runtime.Value{}

var monadAff = gopurs_runtime.Value{}
var bindAff = gopurs_runtime.Value{}
var applyAff = gopurs_runtime.Value{}
var applicativeAff = gopurs_runtime.Value{}

var lift21 = gopurs_runtime.Value{}
var cancelWith = gopurs_runtime.Value{}
var finally = gopurs_runtime.Value{}
var invincible = gopurs_runtime.Value{}
var lazyAff = gopurs_runtime.Value{}
var parallelAff = gopurs_runtime.Value{}
var applicativeParAff = gopurs_runtime.Value{}
var parSequence_ = gopurs_runtime.Apply(parTraverse_, parallelAff)
var monoidParAff = gopurs_runtime.Value{}
var semigroupCanceler = gopurs_runtime.Value{}
var semigroupAff = gopurs_runtime.Value{}
var monadEffectAff = gopurs_runtime.Value{}
var effectCanceler = gopurs_runtime.Value{}
var joinFiber = gopurs_runtime.Value{}
var functorFiber = gopurs_runtime.Value{}
var applyFiber = gopurs_runtime.Value{}
var applicativeFiber = gopurs_runtime.Value{}
var killFiber = gopurs_runtime.Value{}
var fiberCanceler = gopurs_runtime.Value{}
var supervise = gopurs_runtime.Value{}
var monadSTAff = gopurs_runtime.Value{}
var monadThrowAff = gopurs_runtime.Value{}
var monadErrorAff = gopurs_runtime.Value{}
var try = gopurs_runtime.Apply(try, monadErrorAff)
var attempt = try
var runAff = gopurs_runtime.Value{}
var runAff_ = gopurs_runtime.Value{}
var runSuspendedAff = gopurs_runtime.Value{}
var monadRecAff = gopurs_runtime.Value{}
var monoidAff = gopurs_runtime.Value{}
var nonCanceler = gopurs_runtime.Value{}
var monoidCanceler = gopurs_runtime.Value{}
var never = gopurs_runtime.Apply(makeAff, gopurs_runtime.Value{})
var apathize = gopurs_runtime.Value{}
var altParAff = gopurs_runtime.Value{}
var altAff = gopurs_runtime.Value{}
var plusAff = gopurs_runtime.Value{}
var plusParAff = gopurs_runtime.Value{}
var alternativeParAff = gopurs_runtime.Value{}