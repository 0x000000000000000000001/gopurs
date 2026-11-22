package purescript

import "gopurs/output/gopurs_runtime"

import "math"

var Data_Bounded_TopInt = 2147483647
var Data_Bounded_BottomInt = -2147483648
var Data_Bounded_TopChar = string(rune(65535))
var Data_Bounded_BottomChar = string(rune(0))
var Data_Bounded_TopNumber = math.Inf(1)
var Data_Bounded_BottomNumber = math.Inf(-1)

// --- Auto-generated FFI wrappers ---
var _Gopurs_Data_Bounded_BottomChar = // TAST: Char
gopurs_runtime.Box(Data_Bounded_BottomChar)
var _Gopurs_Data_Bounded_BottomInt = // TAST: Int
gopurs_runtime.Box(Data_Bounded_BottomInt)
var _Gopurs_Data_Bounded_BottomNumber = // TAST: Number
gopurs_runtime.Box(Data_Bounded_BottomNumber)
var _Gopurs_Data_Bounded_TopChar = // TAST: Char
gopurs_runtime.Box(Data_Bounded_TopChar)
var _Gopurs_Data_Bounded_TopInt = // TAST: Int
gopurs_runtime.Box(Data_Bounded_TopInt)
var _Gopurs_Data_Bounded_TopNumber = // TAST: Number
gopurs_runtime.Box(Data_Bounded_TopNumber)
