package Dodo_Box

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var max = gopurs_runtime.Value{}
var power = gopurs_runtime.Apply(power, monoidDoc)
var LinePad = gopurs_runtime.Value{}
var LineDoc = gopurs_runtime.Value{}
var LineAppend = gopurs_runtime.Value{}
var FullHeight = gopurs_runtime.Value{}
var FullWidth = gopurs_runtime.Value{}
var AsIs = gopurs_runtime.Value{}
var StpDone = gopurs_runtime.Value{}
var StpLine = gopurs_runtime.Value{}
var StpPad = gopurs_runtime.Value{}
var StpHorz = gopurs_runtime.Value{}
var ResumeEnter = gopurs_runtime.Value{}
var ResumeLeave = gopurs_runtime.Value{}
var ResumeHorzR = gopurs_runtime.Value{}
var ResumeHorzH = gopurs_runtime.Value{}
var ResumeNil = gopurs_runtime.Value{}
var Start = gopurs_runtime.Value{}
var Middle = gopurs_runtime.Value{}
var End = gopurs_runtime.Value{}
var DocLine = gopurs_runtime.Value{}
var DocVApp = gopurs_runtime.Value{}
var DocHApp = gopurs_runtime.Value{}
var DocAlign = gopurs_runtime.Value{}
var DocPad = gopurs_runtime.Value{}
var DocEmpty = gopurs_runtime.Value{}
var BuildEnter = gopurs_runtime.Value{}
var BuildLeave = gopurs_runtime.Value{}
var BuildVAppR = gopurs_runtime.Value{}
var BuildHAppR = gopurs_runtime.Value{}
var BuildHAppH = gopurs_runtime.Value{}
var BuildNil = gopurs_runtime.Value{}
var Horizontal = gopurs_runtime.Value{}
var Vertical = gopurs_runtime.Value{}
var newtypeVertical_ = gopurs_runtime.Value{}
var newtypeHorizontal_ = gopurs_runtime.Value{}
var functorDocBox = gopurs_runtime.Value{}
var functorHorizontal = functorDocBox
var functorVertical = functorDocBox
var eqAlign = gopurs_runtime.Value{}
var vpadding = gopurs_runtime.Value{}
var valign = gopurs_runtime.Value{}
var sizeOf = gopurs_runtime.Value{}
var vappend = gopurs_runtime.Value{}
var vertical = gopurs_runtime.Value{}
var semigroupVertical = gopurs_runtime.Value{}
var monoidVertical = gopurs_runtime.Value{}
var power1 = gopurs_runtime.Apply(power, monoidVertical)
var resume = gopurs_runtime.Value{}
var padWithAlign = gopurs_runtime.Value{}
var isEmpty = gopurs_runtime.Value{}
var hpadding = gopurs_runtime.Value{}
var happend = gopurs_runtime.Value{}
var horizontal = gopurs_runtime.Value{}
var horizontalWithAlign = gopurs_runtime.Value{}
var semigroupHorizontal = gopurs_runtime.Value{}
var monoidHorizontal = gopurs_runtime.Value{}
var halign = gopurs_runtime.Value{}
var resize = gopurs_runtime.Value{}
var verticalWithAlign = gopurs_runtime.Value{}
var formatLine = gopurs_runtime.Value{}
var fill = gopurs_runtime.Value{}
var empty = gopurs_runtime.Value{}
var docBox = gopurs_runtime.Value{}
var build = gopurs_runtime.Value{}
var toDoc = gopurs_runtime.Value{}
