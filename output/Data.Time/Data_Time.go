package Data_Time

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var fromJust = gopurs_runtime.Value{}
var negateDuration = gopurs_runtime.Apply(negateDuration, durationMilliseconds)
var Time = gopurs_runtime.Value{}
var showTime = gopurs_runtime.Value{}
var setSecond = gopurs_runtime.Value{}
var setMinute = gopurs_runtime.Value{}
var setMillisecond = gopurs_runtime.Value{}
var setHour = gopurs_runtime.Value{}
var second = gopurs_runtime.Value{}
var minute = gopurs_runtime.Value{}
var millisecond = gopurs_runtime.Value{}
var millisToTime = gopurs_runtime.Value{}
var hour = gopurs_runtime.Value{}
var timeToMillis = gopurs_runtime.Value{}
var eqTime = gopurs_runtime.Value{}
var ordTime = gopurs_runtime.Value{}
var diff = gopurs_runtime.Value{}
var boundedTime = gopurs_runtime.Value{}
var maxTime = gopurs_runtime.Apply(timeToMillis, gopurs_runtime.Value{})
var minTime = gopurs_runtime.Apply(timeToMillis, gopurs_runtime.Value{})
var adjust = gopurs_runtime.Value{}
