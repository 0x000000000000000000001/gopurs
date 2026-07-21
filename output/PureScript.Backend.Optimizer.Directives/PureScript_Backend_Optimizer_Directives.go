package PureScript_Backend_Optimizer_Directives

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var expectMap = gopurs_runtime.Value{}
var keyword = gopurs_runtime.Value{}
var label = gopurs_runtime.Apply(expectMap, gopurs_runtime.Value{})
var natural = gopurs_runtime.Apply(expectMap, gopurs_runtime.Value{})
var qualified = gopurs_runtime.Apply(expectMap, gopurs_runtime.Value{})
var unqualified = gopurs_runtime.Apply(expectMap, gopurs_runtime.Value{})
var equals = gopurs_runtime.Apply(expectMap, gopurs_runtime.Value{})
var parseInlineDirective = gopurs_runtime.Apply(gopurs_runtime.Value{}, gopurs_runtime.Apply(gopurs_runtime.Value{}, gopurs_runtime.Value{}))
var dotDot = gopurs_runtime.Apply(expectMap, gopurs_runtime.Value{})
var dot = gopurs_runtime.Apply(expectMap, gopurs_runtime.Value{})
var parseInlineAccessor = gopurs_runtime.Apply(gopurs_runtime.Value{}, gopurs_runtime.Apply(gopurs_runtime.Value{}, InlineProp))
var parseDirective = gopurs_runtime.Apply(applyFirst, applyParser)
var parseDirectiveMaybe = gopurs_runtime.Apply(gopurs_runtime.Value{}, gopurs_runtime.Apply(gopurs_runtime.Value{}, Just))
var parseDirectiveLine = gopurs_runtime.Value{}
var parseDirectiveFile = gopurs_runtime.Value{}
var parseDirectiveExport = gopurs_runtime.Value{}
var parseDirectiveHeader = gopurs_runtime.Value{}
