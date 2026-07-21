package Data_Traversable

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var _ = fmt.Println
var _ = gopurs_runtime.TypeInt

var traverse = gopurs_runtime.Value{}
var traversableTuple = gopurs_runtime.Value{}
var traversableMultiplicative = gopurs_runtime.Value{}
var traversableMaybe = gopurs_runtime.Value{}
var traversableIdentity = gopurs_runtime.Value{}
var traversableEither = gopurs_runtime.Value{}
var traversableDual = gopurs_runtime.Value{}
var traversableDisj = gopurs_runtime.Value{}
var traversableConst = gopurs_runtime.Value{}
var traversableConj = gopurs_runtime.Value{}

var traversableCompose = gopurs_runtime.Value{}

var traversableAdditive = gopurs_runtime.Value{}
var sequenceDefault = gopurs_runtime.Value{}

var traversableArray = gopurs_runtime.Value{}

var sequence = gopurs_runtime.Value{}
var traversableApp = gopurs_runtime.Value{}
var traversableCoproduct = gopurs_runtime.Value{}
var traversableFirst = gopurs_runtime.Value{}
var traversableLast = gopurs_runtime.Value{}
var traversableProduct = gopurs_runtime.Value{}
var traverseDefault = gopurs_runtime.Value{}
var mapAccumR = gopurs_runtime.Value{}
var scanr = gopurs_runtime.Value{}
var mapAccumL = gopurs_runtime.Value{}
var scanl = gopurs_runtime.Value{}
var for = gopurs_runtime.Value{}