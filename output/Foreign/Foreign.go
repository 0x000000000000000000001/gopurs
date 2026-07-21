package Foreign

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var isArray = gopurs_runtime.Value{}
var isNull = gopurs_runtime.Value{}
var isUndefined = gopurs_runtime.Value{}
var tagOf = gopurs_runtime.Value{}
var typeOf = gopurs_runtime.Value{}
var ForeignError = gopurs_runtime.Value{}
var TypeMismatch = gopurs_runtime.Value{}
var ErrorAtIndex = gopurs_runtime.Value{}
var ErrorAtProperty = gopurs_runtime.Value{}
var unsafeToForeign = unsafeCoerce
var unsafeFromForeign = unsafeCoerce
var showForeignError = gopurs_runtime.Value{}
var renderForeignError = gopurs_runtime.Value{}
var readUndefined = gopurs_runtime.Value{}
var readNullOrUndefined = gopurs_runtime.Value{}
var readNull = gopurs_runtime.Value{}
var fail = gopurs_runtime.Value{}
var readArray = gopurs_runtime.Value{}
var unsafeReadTagged = gopurs_runtime.Value{}
var readBoolean = gopurs_runtime.Value{}
var readNumber = gopurs_runtime.Value{}
var readInt = gopurs_runtime.Value{}
var readString = gopurs_runtime.Value{}
var readChar = gopurs_runtime.Value{}
var eqForeignError = gopurs_runtime.Value{}
var ordForeignError = gopurs_runtime.Value{}
