package Data_Profunctor_Split

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var SplitF = gopurs_runtime.Value{}
var unSplit = gopurs_runtime.Value{}
var split = gopurs_runtime.Value{}
var profunctorSplit = gopurs_runtime.Value{}
var lowerSplit = gopurs_runtime.Value{}
var liftSplit = gopurs_runtime.Apply(split, gopurs_runtime.Value{})
var hoistSplit = gopurs_runtime.Value{}
var functorSplit = gopurs_runtime.Value{}
