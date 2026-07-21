package PureScript_Backend_Optimizer_Tracer_Printer

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var fold = gopurs_runtime.Apply(gopurs_runtime.Value{}, monoidDoc)
var lines = gopurs_runtime.Apply(foldrArray, appendBreak)
var words = gopurs_runtime.Apply(foldrArray, appendSpace)
var power = gopurs_runtime.Apply(power, monoidString)
var PrecBlock = gopurs_runtime.Value{}
var PrecApply = gopurs_runtime.Value{}
var PrecAccess = gopurs_runtime.Value{}
var PrecAtom = gopurs_runtime.Value{}
var wrapPrec = gopurs_runtime.Value{}
var printUncurriedApp = gopurs_runtime.Value{}
var printUncurriedAbs = gopurs_runtime.Value{}
var printRewrite = gopurs_runtime.Value{}
var printUnpackOpCase = gopurs_runtime.Value{}
var printRecord = gopurs_runtime.Value{}
var printQualified = gopurs_runtime.Value{}
var printLevel = gopurs_runtime.Value{}
var printLet' = gopurs_runtime.Value{}
var printIdent = gopurs_runtime.Value{}
var printLocal = gopurs_runtime.Value{}
var printDistOpCase = gopurs_runtime.Value{}
var printCurriedApp = gopurs_runtime.Value{}
var printCurriedAbs = gopurs_runtime.Value{}
var printBranch = gopurs_runtime.Value{}
var printBackendRewriteCase = gopurs_runtime.Value{}
var printBackendAccessor = gopurs_runtime.Value{}
var printArray = gopurs_runtime.Value{}
var printLiteral = gopurs_runtime.Value{}
var primOp = gopurs_runtime.Value{}
var printBackendEffect = gopurs_runtime.Value{}
var printBackendOperator1 = gopurs_runtime.Value{}
var printBackendOperatorNum = gopurs_runtime.Value{}
var printBackendOperatorOrd = gopurs_runtime.Value{}
var printBackendOperator2 = gopurs_runtime.Value{}
var printBackendOperator = gopurs_runtime.Value{}
var printBackendSyntax = gopurs_runtime.Value{}
var printBackendExpr = gopurs_runtime.Apply(foldBackendExpr, printBackendSyntax)
var heading = gopurs_runtime.Value{}
var printSteps = gopurs_runtime.Value{}
var printModuleSteps = gopurs_runtime.Value{}
