package PureScript_CST

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var _ = fmt.Println
var _ = gopurs_runtime.TypeInt

var foldMap = gopurs_runtime.Apply(gopurs_runtime.Value{}, monoidString)
var ParseSucceeded = gopurs_runtime.Value{}
var ParseSucceededWithErrors = gopurs_runtime.Value{}
var ParseFailed = gopurs_runtime.Value{}
var PartialModule = gopurs_runtime.Value{}
var toRecoveredParserResult = gopurs_runtime.Value{}
var toRecovered = unsafeCoerce
var printModule = gopurs_runtime.Value{}
var parseType = gopurs_runtime.Value{}
var parsePartialModule = gopurs_runtime.Value{}
var parseModule = gopurs_runtime.Value{}
var parseImportDecl = gopurs_runtime.Value{}
var parseExpr = gopurs_runtime.Value{}
var parseDecl = gopurs_runtime.Value{}
var parseBinder = gopurs_runtime.Value{}