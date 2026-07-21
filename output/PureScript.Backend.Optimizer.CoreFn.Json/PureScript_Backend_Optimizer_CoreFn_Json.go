package PureScript_Backend_Optimizer_CoreFn_Json

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var _ = fmt.Println
var _ = gopurs_runtime.TypeInt

var unsafeIndex = gopurs_runtime.Apply(runFn2, unsafeIndexImpl)
var traverse = gopurs_runtime.Apply(gopurs_runtime.Value{}, applicativeEither)
var guard = gopurs_runtime.Apply(guard, alternativeMaybe)
var getFieldOptional' = gopurs_runtime.Value{}
var getField = gopurs_runtime.Value{}
var decodeString = gopurs_runtime.Value{}
var decodeProperName = gopurs_runtime.Apply(unsafeCoerce, decodeString)
var decodeNumber = gopurs_runtime.Value{}
var decodeJObject = gopurs_runtime.Value{}
var decodeJArray = gopurs_runtime.Value{}
var decodeIdent = gopurs_runtime.Apply(unsafeCoerce, decodeString)
var decodeBoolean = gopurs_runtime.Value{}
var decodeArray = gopurs_runtime.Value{}
var decodeModuleName = gopurs_runtime.Value{}
var decodeConstructorType = gopurs_runtime.Value{}
var decodeImport = gopurs_runtime.Value{}
var decodeInt = gopurs_runtime.Value{}
var decodeCodePoint = gopurs_runtime.Value{}
var decodeMeta = gopurs_runtime.Value{}
var decodeAnn = gopurs_runtime.Value{}
var decodeQualified = gopurs_runtime.Value{}
var decodeReExports = gopurs_runtime.Value{}
var decodeRecord = gopurs_runtime.Value{}
var decodeSourcePos = gopurs_runtime.Value{}
var decodeSourceSpan = gopurs_runtime.Value{}
var decodeComment = gopurs_runtime.Value{}
var decodeStringLiteral = gopurs_runtime.Value{}
var decodeLiteral = gopurs_runtime.Value{}

var decodeBinder = gopurs_runtime.Value{}

var decodeGuard = gopurs_runtime.Value{}
var decodeExpr = gopurs_runtime.Value{}
var decodeCaseAlternative = gopurs_runtime.Value{}
var decodeBinding = gopurs_runtime.Value{}
var decodeBind = gopurs_runtime.Value{}

var decodeModule' = gopurs_runtime.Value{}
var decodeModule = gopurs_runtime.Apply(decodeModule', decodeAnn)