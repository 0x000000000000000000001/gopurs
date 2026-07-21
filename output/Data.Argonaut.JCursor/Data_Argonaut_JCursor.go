package Data_Argonaut_JCursor

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var _ = fmt.Println
var _ = gopurs_runtime.TypeInt

var toUnfoldable = gopurs_runtime.Value{}
var fromFoldable = gopurs_runtime.Apply(gopurs_runtime.Value{}, Cons)
var fromFoldable1 = gopurs_runtime.Apply(foldrArray, Cons)
var decodeJson = gopurs_runtime.Apply(decodeArray, Right)
var JsonPrim = gopurs_runtime.Value{}
var JCursorTop = gopurs_runtime.Value{}
var JField = gopurs_runtime.Value{}
var JIndex = gopurs_runtime.Value{}

var showJCursor = gopurs_runtime.Value{}



var semigroupJCursor = gopurs_runtime.Value{}

var runJsonPrim = gopurs_runtime.Value{}
var showJsonPrim = gopurs_runtime.Value{}
var print = gopurs_runtime.Value{}
var primToJson = gopurs_runtime.Value{}
var primStr = gopurs_runtime.Value{}
var primNum = gopurs_runtime.Value{}
var primNull = gopurs_runtime.Value{}
var primBool = gopurs_runtime.Value{}

var toPrims = gopurs_runtime.Apply(caseJson, gopurs_runtime.Value{})

var monoidJCursor = gopurs_runtime.Value{}
var inferEmpty = gopurs_runtime.Value{}

var eqJCursor = gopurs_runtime.Value{}

var ordJCursor = gopurs_runtime.Value{}

var encodeJsonJCursor = gopurs_runtime.Value{}
var downIndex = gopurs_runtime.Value{}
var downField = gopurs_runtime.Value{}

var insideOut = gopurs_runtime.Value{}

var decodeJsonJCursor = gopurs_runtime.Value{}

var cursorSet = gopurs_runtime.Value{}

var fromPrims = gopurs_runtime.Value{}

var cursorGet = gopurs_runtime.Value{}