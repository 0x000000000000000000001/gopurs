package Dodo

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var max = gopurs_runtime.Value{}
var max1 = gopurs_runtime.Value{}
var min = gopurs_runtime.Value{}
var power = gopurs_runtime.Apply(power, monoidString)
var Printer = gopurs_runtime.Value{}
var Doc = gopurs_runtime.Value{}
var Dedent = gopurs_runtime.Value{}
var LeaveAnnotation = gopurs_runtime.Value{}
var LeaveFlexGroup = gopurs_runtime.Value{}
var LeaveLocal = gopurs_runtime.Value{}
var NoFlexGroup = gopurs_runtime.Value{}
var FlexGroupPending = gopurs_runtime.Value{}
var FlexGroupReset = gopurs_runtime.Value{}
var withPosition = WithPosition
var withLocalOptions = Local
var twoSpaces = gopurs_runtime.Value{}
var text = gopurs_runtime.Value{}
var tabs = gopurs_runtime.Value{}
var space = gopurs_runtime.Value{}
var plainText = gopurs_runtime.Value{}
var locally = gopurs_runtime.Value{}
var indent = gopurs_runtime.Value{}
var fourSpaces = gopurs_runtime.Value{}
var foldWith = gopurs_runtime.Value{}
var foldWithSeparator = gopurs_runtime.Value{}
var flexSelect = gopurs_runtime.Value{}
var flexGroup = gopurs_runtime.Value{}
var flexAlt = FlexAlt
var encloseWithSeparator = gopurs_runtime.Value{}
var encloseEmptyAlt = gopurs_runtime.Value{}
var enclose = gopurs_runtime.Value{}
var calcRibbonWidth = gopurs_runtime.Value{}
var storeOptions = gopurs_runtime.Value{}
var print = gopurs_runtime.Value{}
var break = gopurs_runtime.Value{}
var softBreak = gopurs_runtime.Value{}
var spaceBreak = gopurs_runtime.Value{}
var appendSpaceBreak = gopurs_runtime.Value{}
var paragraph = gopurs_runtime.Value{}
var textParagraph = gopurs_runtime.Value{}
var appendSpace = gopurs_runtime.Value{}
var words = gopurs_runtime.Value{}
var appendBreak = gopurs_runtime.Value{}
var lines = gopurs_runtime.Value{}
var annotate = gopurs_runtime.Value{}
var align = gopurs_runtime.Value{}
var alignCurrentColumn = gopurs_runtime.Value{}
