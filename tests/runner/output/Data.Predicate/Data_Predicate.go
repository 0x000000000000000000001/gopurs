package Data_Predicate

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var Predicate = gopurs_runtime.Value{}
var newtypePredicate = gopurs_runtime.Value{}
var heytingAlgebraPredicate = gopurs_runtime.Value{}
var contravariantPredicate = gopurs_runtime.Value{}
var booleanAlgebraPredicate = gopurs_runtime.Apply(booleanAlgebraFn, booleanAlgebraBoolean)
