package PureScript_CST_Print

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var power = gopurs_runtime.Apply(power, monoidString)
var foldMap = gopurs_runtime.Apply(gopurs_runtime.Value{}, monoidString)
var ShowLayout = gopurs_runtime.Value{}
var HideLayout = gopurs_runtime.Value{}
var printQualified = gopurs_runtime.Value{}
var printTokenWithOption = gopurs_runtime.Value{}
var printToken = gopurs_runtime.Apply(printTokenWithOption, gopurs_runtime.Value{})
var printLineFeed = gopurs_runtime.Value{}
var printComment = gopurs_runtime.Value{}
var printSourceTokenWithOption = gopurs_runtime.Value{}
var printSourceToken = gopurs_runtime.Apply(printSourceTokenWithOption, gopurs_runtime.Value{})
