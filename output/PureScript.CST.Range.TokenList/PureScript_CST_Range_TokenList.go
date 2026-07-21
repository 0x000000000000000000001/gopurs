package PureScript_CST_Range_TokenList

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var unsafeIndex = gopurs_runtime.Apply(runFn2, unsafeIndexImpl)
var TokenEmpty = gopurs_runtime.Value{}
var TokenCons = gopurs_runtime.Value{}
var TokenWrap = gopurs_runtime.Value{}
var TokenAppend = gopurs_runtime.Value{}
var TokenDefer = gopurs_runtime.Value{}
var TokenArray = gopurs_runtime.Value{}
var UnconsDone = gopurs_runtime.Value{}
var UnconsMore = gopurs_runtime.Value{}
var wrap = TokenWrap
var singleton = gopurs_runtime.Value{}
var semigroupTokenList = gopurs_runtime.Value{}
var uncons2 = gopurs_runtime.Value{}
var uncons' = gopurs_runtime.Value{}
var toUnfoldable = gopurs_runtime.Value{}
var uncons = gopurs_runtime.Apply(uncons', gopurs_runtime.Value{})
var toArray = gopurs_runtime.Value{}
var monoidTokenList = gopurs_runtime.Value{}
var lazyTokenList = gopurs_runtime.Value{}
var head = gopurs_runtime.Value{}
var fromArray = gopurs_runtime.Value{}
var cons = TokenCons
