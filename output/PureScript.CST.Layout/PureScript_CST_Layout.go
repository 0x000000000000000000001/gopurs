package PureScript_CST_Layout

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var _ = fmt.Println
var _ = gopurs_runtime.TypeInt

var LytRoot = gopurs_runtime.Value{}
var LytTopDecl = gopurs_runtime.Value{}
var LytTopDeclHead = gopurs_runtime.Value{}
var LytDeclGuard = gopurs_runtime.Value{}
var LytCase = gopurs_runtime.Value{}
var LytCaseBinders = gopurs_runtime.Value{}
var LytCaseGuard = gopurs_runtime.Value{}
var LytLambdaBinders = gopurs_runtime.Value{}
var LytParen = gopurs_runtime.Value{}
var LytBrace = gopurs_runtime.Value{}
var LytSquare = gopurs_runtime.Value{}
var LytIf = gopurs_runtime.Value{}
var LytThen = gopurs_runtime.Value{}
var LytProperty = gopurs_runtime.Value{}
var LytForall = gopurs_runtime.Value{}
var LytTick = gopurs_runtime.Value{}
var LytLet = gopurs_runtime.Value{}
var LytLetStmt = gopurs_runtime.Value{}
var LytWhere = gopurs_runtime.Value{}
var LytOf = gopurs_runtime.Value{}
var LytDo = gopurs_runtime.Value{}
var LytAdo = gopurs_runtime.Value{}
var lytToken = gopurs_runtime.Value{}
var isTopDecl = gopurs_runtime.Value{}
var isIndented = gopurs_runtime.Value{}
var eqLayoutDelim = gopurs_runtime.Value{}
var insertLayout = gopurs_runtime.Value{}
var ordLayoutDelim = gopurs_runtime.Value{}
var currentIndent = gopurs_runtime.Value{}