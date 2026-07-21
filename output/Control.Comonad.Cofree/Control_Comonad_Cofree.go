package Control_Comonad_Cofree

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var _ = fmt.Println
var _ = gopurs_runtime.TypeInt

var state = gopurs_runtime.Value{}
var monadRecStateT = gopurs_runtime.Apply(monadRecStateT, monadRecIdentity)
var tail = gopurs_runtime.Value{}
var mkCofree = gopurs_runtime.Value{}
var lazyCofree = gopurs_runtime.Value{}

var hoistCofree = gopurs_runtime.Value{}

var head = gopurs_runtime.Value{}
var functorCofree = gopurs_runtime.Value{}
var functorWithIndexCofree = gopurs_runtime.Value{}
var foldableCofree = gopurs_runtime.Value{}
var foldableWithIndexCofree = gopurs_runtime.Value{}

var traversableCofree = gopurs_runtime.Value{}

var traversableWithIndexCofree = gopurs_runtime.Value{}
var extendCofree = gopurs_runtime.Value{}

var eqCofree = gopurs_runtime.Value{}

var ordCofree = gopurs_runtime.Value{}

var eq1Cofree = gopurs_runtime.Value{}
var ord1Cofree = gopurs_runtime.Value{}
var deferCofree = gopurs_runtime.Value{}

var semigroupCofree = gopurs_runtime.Value{}

var monoidCofree = gopurs_runtime.Value{}

var comonadCofree = gopurs_runtime.Value{}
var explore = gopurs_runtime.Value{}
var exploreM = gopurs_runtime.Value{}

var buildCofree = gopurs_runtime.Value{}

var monadCofree = gopurs_runtime.Value{}
var bindCofree = gopurs_runtime.Value{}
var applyCofree = gopurs_runtime.Value{}
var applicativeCofree = gopurs_runtime.Value{}