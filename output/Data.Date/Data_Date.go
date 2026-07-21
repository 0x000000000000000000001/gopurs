package Data_Date

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var calcDiff = gopurs_runtime.Value{}
var calcWeekday = gopurs_runtime.Value{}
var canonicalDateImpl = gopurs_runtime.Value{}
var fromJust = gopurs_runtime.Value{}
var greaterThan = gopurs_runtime.Value{}
var Date = gopurs_runtime.Value{}
var year = gopurs_runtime.Value{}
var weekday = gopurs_runtime.Apply(_unsafePartial, gopurs_runtime.Value{})
var showDate = gopurs_runtime.Value{}
var month = gopurs_runtime.Value{}
var isLeapYear = gopurs_runtime.Value{}
var lastDayOfMonth = gopurs_runtime.Value{}
var eqDate = gopurs_runtime.Value{}
var ordDate = gopurs_runtime.Value{}
var enumDate = gopurs_runtime.Value{}
var diff = gopurs_runtime.Value{}
var day = gopurs_runtime.Value{}
var canonicalDate = gopurs_runtime.Value{}
var exactDate = gopurs_runtime.Value{}
var boundedDate = gopurs_runtime.Value{}
var adjust = gopurs_runtime.Value{}
