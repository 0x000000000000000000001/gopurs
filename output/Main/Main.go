package Main

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var filterA = gopurs_runtime.Apply(filterA, applicativeAff)
var traverse = gopurs_runtime.Apply(gopurs_runtime.Value{}, applicativeAff)
var fromFoldable = gopurs_runtime.Apply(foldrArray, Cons)
var buildModules = gopurs_runtime.Apply(buildModules, monadAff)
var readCoreFnModule = gopurs_runtime.Value{}
var Main = gopurs_runtime.Apply(gopurs_runtime.Value{}, gopurs_runtime.Apply(pureE, gopurs_runtime.Value{}))
