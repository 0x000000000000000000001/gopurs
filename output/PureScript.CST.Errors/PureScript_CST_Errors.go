package PureScript_CST_Errors

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var _ = fmt.Println
var _ = gopurs_runtime.TypeInt

var UnexpectedEof = gopurs_runtime.Value{}
var ExpectedEof = gopurs_runtime.Value{}
var UnexpectedToken = gopurs_runtime.Value{}
var ExpectedToken = gopurs_runtime.Value{}
var ExpectedClass = gopurs_runtime.Value{}
var LexExpected = gopurs_runtime.Value{}
var LexInvalidCharEscape = gopurs_runtime.Value{}
var LexCharEscapeOutOfRange = gopurs_runtime.Value{}
var LexHexOutOfRange = gopurs_runtime.Value{}
var LexIntOutOfRange = gopurs_runtime.Value{}
var LexNumberOutOfRange = gopurs_runtime.Value{}
var RecoveredError = gopurs_runtime.Value{}
var printTokenError = gopurs_runtime.Value{}
var printParseError = gopurs_runtime.Value{}