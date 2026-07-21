package PureScript_CST_TokenStream

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var _ = fmt.Println
var _ = gopurs_runtime.TypeInt

var TokenEOF = gopurs_runtime.Value{}
var TokenError = gopurs_runtime.Value{}
var TokenCons = gopurs_runtime.Value{}
var TokenStream = gopurs_runtime.Value{}
var newtypeTokenStream = gopurs_runtime.Value{}
var step = gopurs_runtime.Value{}
var unwindLayout = gopurs_runtime.Value{}
var layoutStack = gopurs_runtime.Value{}
var consTokens = gopurs_runtime.Value{}