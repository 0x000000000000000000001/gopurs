package PureScript_CST_Lexer

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var _ = fmt.Println
var _ = gopurs_runtime.TypeInt

var foldMap = gopurs_runtime.Apply(gopurs_runtime.Value{}, monoidArray)
var fold1 = gopurs_runtime.Apply(gopurs_runtime.Value{}, gopurs_runtime.Value{})
var consTokens = gopurs_runtime.Apply(consTokens, foldableArray)
var LexFail = gopurs_runtime.Value{}
var LexSucc = gopurs_runtime.Value{}
var isCharCodePoint = gopurs_runtime.Value{}
var isCharChar = gopurs_runtime.Value{}
var toModuleName = gopurs_runtime.Value{}
var qualLength = gopurs_runtime.Value{}
var optional = gopurs_runtime.Value{}
var mkUnexpected = gopurs_runtime.Value{}
var regex = gopurs_runtime.Value{}
var shebangComment = gopurs_runtime.Apply(regex, gopurs_runtime.Apply(LexExpected, gopurs_runtime.Str("shebang")))
var satisfy = gopurs_runtime.Value{}
var string = gopurs_runtime.Value{}
var many = gopurs_runtime.Value{}
var functorLex = gopurs_runtime.Value{}
var spaceComment = gopurs_runtime.Value{}
var char' = gopurs_runtime.Value{}
var char = gopurs_runtime.Value{}
var bumpText = gopurs_runtime.Value{}
var bumpToken = gopurs_runtime.Value{}
var bumpComment = gopurs_runtime.Value{}
var applyLex = gopurs_runtime.Value{}
var altLex = gopurs_runtime.Value{}
var comment = gopurs_runtime.Apply(gopurs_runtime.Value{}, gopurs_runtime.Apply(regex, gopurs_runtime.Apply(LexExpected, gopurs_runtime.Str("block comment"))))
var lineComment = gopurs_runtime.Apply(gopurs_runtime.Value{}, gopurs_runtime.Value{})
var leadingComments = gopurs_runtime.Apply(many, gopurs_runtime.Apply(gopurs_runtime.Value{}, gopurs_runtime.Value{}))
var oneLineComment = gopurs_runtime.Value{}
var leadingShebangs = gopurs_runtime.Value{}
var leadingModuleComments = gopurs_runtime.Value{}
var token = gopurs_runtime.Value{}
var lexToken = gopurs_runtime.Value{}
var trailingComments = gopurs_runtime.Apply(many, gopurs_runtime.Apply(gopurs_runtime.Value{}, gopurs_runtime.Value{}))
var lexWithState' = gopurs_runtime.Value{}
var lexModule = gopurs_runtime.Apply(lexWithState', leadingModuleComments)
var lexWithState = gopurs_runtime.Apply(lexWithState', leadingComments)
var lex = gopurs_runtime.Apply(lexWithState, gopurs_runtime.Value{})