package Data_DateTime_Instant

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var fromDateTimeImpl = gopurs_runtime.Value{}
var toDateTimeImpl = gopurs_runtime.Value{}
var negateDuration = gopurs_runtime.Apply(negateDuration, durationMilliseconds)
var unInstant = gopurs_runtime.Value{}
var toDateTime = gopurs_runtime.Apply(toDateTimeImpl, gopurs_runtime.Apply(_unsafePartial, gopurs_runtime.Value{}))
var showInstant = gopurs_runtime.Value{}
var ordDateTime = ordNumber
var instant = gopurs_runtime.Value{}
var fromDateTime = gopurs_runtime.Value{}
var fromDate = gopurs_runtime.Value{}
var eqDateTime = eqNumber
var diff = gopurs_runtime.Value{}
var boundedInstant = gopurs_runtime.Value{}
