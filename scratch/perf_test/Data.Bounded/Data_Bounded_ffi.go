package Data_Bounded

import (
	"gopurs/output/gopurs_runtime"
	"math"
)

var TopInt = gopurs_runtime.Int(2147483647)
var BottomInt = gopurs_runtime.Int(-2147483648)

var TopChar = gopurs_runtime.Str(string(rune(65535)))
var BottomChar = gopurs_runtime.Str(string(rune(0)))

var TopNumber = gopurs_runtime.Float(math.Inf(1))
var BottomNumber = gopurs_runtime.Float(math.Inf(-1))
