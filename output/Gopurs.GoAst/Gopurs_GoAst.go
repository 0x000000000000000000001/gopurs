package Gopurs_GoAst

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var eq2 = gopurs_runtime.Apply(eqArrayImpl, eqStringImpl)
var GoVar = gopurs_runtime.Value{}
var GoString = gopurs_runtime.Value{}
var GoInt = gopurs_runtime.Value{}
var GoCall = gopurs_runtime.Value{}
var GoSelector = gopurs_runtime.Value{}
var GoFunc = gopurs_runtime.Value{}
var GoBlock = gopurs_runtime.Value{}
var GoReturn = gopurs_runtime.Value{}
var GoAssign = gopurs_runtime.Value{}
var GoRaw = gopurs_runtime.Value{}
var eqGoExpr = gopurs_runtime.Value{}
