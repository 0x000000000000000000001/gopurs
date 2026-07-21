package PureScript_CST_Parser_Monad

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var _ = fmt.Println
var _ = gopurs_runtime.TypeInt

var toUnfoldable = gopurs_runtime.Apply(gopurs_runtime.Value{}, gopurs_runtime.Value{})
var More = gopurs_runtime.Value{}
var Done = gopurs_runtime.Value{}
var ParseFail = gopurs_runtime.Value{}
var ParseSucc = gopurs_runtime.Value{}
var Parser = gopurs_runtime.Value{}
var lazyParser = gopurs_runtime.Value{}
var functorParser = gopurs_runtime.Value{}
var applyParser = gopurs_runtime.Value{}
var bindParser = gopurs_runtime.Value{}
var applicativeParser = gopurs_runtime.Value{}
var monadParser = gopurs_runtime.Value{}
var altParser = gopurs_runtime.Value{}
var try = gopurs_runtime.Value{}
var take = gopurs_runtime.Value{}
var runParser' = gopurs_runtime.Value{}
var recover = gopurs_runtime.Value{}
var optional = gopurs_runtime.Value{}
var many = gopurs_runtime.Value{}
var lookAhead = gopurs_runtime.Value{}
var initialParserState = gopurs_runtime.Value{}
var fromParserResult = gopurs_runtime.Value{}
var runParser = gopurs_runtime.Value{}
var fail = gopurs_runtime.Value{}
var eof = gopurs_runtime.Apply(mkFn4, gopurs_runtime.Value{})