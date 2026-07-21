package Data_Int

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var fromNumberImpl = gopurs_runtime.Value{}
var fromStringAsImpl = gopurs_runtime.Value{}
var pow = gopurs_runtime.Value{}
var quot = gopurs_runtime.Value{}
var rem = gopurs_runtime.Value{}
var toNumber = gopurs_runtime.Value{}
var toStringAs = gopurs_runtime.Value{}
var Even = gopurs_runtime.Value{}
var Odd = gopurs_runtime.Value{}
var showParity = gopurs_runtime.Value{}
var radix = gopurs_runtime.Value{}
var odd = gopurs_runtime.Value{}
var octal = gopurs_runtime.Value{}
var hexadecimal = gopurs_runtime.Value{}
var fromStringAs = gopurs_runtime.Apply(fromStringAsImpl, Just)
var fromString = gopurs_runtime.Apply(fromStringAs, gopurs_runtime.Value{})
var fromNumber = gopurs_runtime.Apply(fromNumberImpl, Just)
var unsafeClamp = gopurs_runtime.Value{}
var round = gopurs_runtime.Value{}
var trunc = gopurs_runtime.Value{}
var floor = gopurs_runtime.Value{}
var even = gopurs_runtime.Value{}
var parity = gopurs_runtime.Value{}
var eqParity = gopurs_runtime.Value{}
var ordParity = gopurs_runtime.Value{}
var semiringParity = gopurs_runtime.Value{}
var ringParity = gopurs_runtime.Value{}
var divisionRingParity = gopurs_runtime.Value{}
var decimal = gopurs_runtime.Value{}
var commutativeRingParity = gopurs_runtime.Value{}
var euclideanRingParity = gopurs_runtime.Value{}
var ceil = gopurs_runtime.Value{}
var boundedParity = gopurs_runtime.Value{}
var binary = gopurs_runtime.Value{}
var base36 = gopurs_runtime.Value{}
