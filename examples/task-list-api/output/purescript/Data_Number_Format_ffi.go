package purescript

import "gopurs/output/gopurs_runtime"

import (
	"fmt"
	"strings"
)

func trimExp(s string) string {
	s = strings.Replace(s, "e+0", "e+", 1)
	s = strings.Replace(s, "e-0", "e-", 1)
	return s
}

func Data_Number_Format_ToExponentialNative(fractionDigits int, n float64) string {
	return trimExp(fmt.Sprintf("%.*e", fractionDigits, n))
}
func Data_Number_Format_ToFixedNative(fractionDigits int, n float64) string {
	return fmt.Sprintf("%.*f", fractionDigits, n)
}
func Data_Number_Format_ToPrecisionNative(precision int, n float64) string {
	return trimExp(fmt.Sprintf("%.*g", precision, n))
}
func Data_Number_Format_ToString(n float64) string {
	return fmt.Sprintf("%v", n)
}

// --- Auto-generated FFI wrappers ---
var _Gopurs_Data_Number_Format_ToExponentialNative = // TAST: (Func [Int, Number] String)
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int](arg0)
	go_arg1 := gopurs_runtime.Unbox[float64](arg1)
	go_res := Data_Number_Format_ToExponentialNative(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_Number_Format_ToFixedNative = // TAST: (Func [Int, Number] String)
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int](arg0)
	go_arg1 := gopurs_runtime.Unbox[float64](arg1)
	go_res := Data_Number_Format_ToFixedNative(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_Number_Format_ToPrecisionNative = // TAST: (Func [Int, Number] String)
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int](arg0)
	go_arg1 := gopurs_runtime.Unbox[float64](arg1)
	go_res := Data_Number_Format_ToPrecisionNative(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_Number_Format_ToString = // TAST: (Func [Number] String)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[float64](arg0)
	go_res := Data_Number_Format_ToString(go_arg0)
	return gopurs_runtime.Box(go_res)
})
