package Data_JSDate

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var _ = fmt.Println
var _ = gopurs_runtime.TypeInt

var toUTCString = gopurs_runtime.Value{}
var toTimeString = gopurs_runtime.Value{}
var toString = gopurs_runtime.Value{}
var toInstant = gopurs_runtime.Value{}
var toISOString = gopurs_runtime.Value{}
var toDateTime = gopurs_runtime.Value{}
var toDateString = gopurs_runtime.Value{}
var toDate = gopurs_runtime.Value{}
var readDate = gopurs_runtime.Apply(unsafeReadTagged, monadIdentity)
var getUTCSeconds = gopurs_runtime.Value{}
var getUTCMonth = gopurs_runtime.Value{}
var getUTCMinutes = gopurs_runtime.Value{}
var getUTCMilliseconds = gopurs_runtime.Value{}
var getUTCHours = gopurs_runtime.Value{}
var getUTCFullYear = gopurs_runtime.Value{}
var getUTCDay = gopurs_runtime.Value{}
var getUTCDate = gopurs_runtime.Value{}
var getTimezoneOffset = gopurs_runtime.Value{}
var getTime = gopurs_runtime.Value{}
var showJSDate = gopurs_runtime.Value{}
var getSeconds = gopurs_runtime.Value{}
var getMonth = gopurs_runtime.Value{}
var getMinutes = gopurs_runtime.Value{}
var getMilliseconds = gopurs_runtime.Value{}
var getHours = gopurs_runtime.Value{}
var getFullYear = gopurs_runtime.Value{}
var getDay = gopurs_runtime.Value{}
var getDate = gopurs_runtime.Value{}
var fromDateTime = gopurs_runtime.Value{}
var eqJSDate = gopurs_runtime.Value{}
var ordJSDate = gopurs_runtime.Value{}