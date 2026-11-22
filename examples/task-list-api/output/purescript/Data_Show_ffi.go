package purescript

import "gopurs/output/gopurs_runtime"

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Data_Show_ShowIntImpl(n int64) string {
	return fmt.Sprintf("%v", n)
}
func Data_Show_ShowNumberImpl(n float64) string {
	if n == 0 {
		// JavaScript Number#toString renders both signed zeros as "0".
		return "0.0"
	} else if math.IsNaN(n) {
		return "NaN"
	} else if math.IsInf(n, 1) {
		return "Infinity"
	} else if math.IsInf(n, -1) {
		return "-Infinity"
	}

	absN := math.Abs(n)
	if absN >= 1e21 || absN < 1e-6 {
		// Keep the shortest round-tripping digits, with JS notation thresholds
		// and an unpadded exponent (Go's e-07 becomes JavaScript's e-7).
		str := strconv.FormatFloat(n, 'e', -1, 64)
		exponent := strings.IndexByte(str, 'e')
		return str[:exponent+2] + strings.TrimLeft(str[exponent+2:], "0")
	}

	str := strconv.FormatFloat(n, 'f', -1, 64)
	if strings.Contains(str, ".") {
		return str
	}
	return str + ".0"
}
func Data_Show_ShowCharImpl(c string) string {
	return fmt.Sprintf("'%s'", c)
}
func Data_Show_ShowStringImpl(s string) string {
	return fmt.Sprintf("%q", s)
}
func Data_Show_ShowArrayImpl(f func(interface{}) string, arr []interface{}) string {
	res := "["
	for i, v := range arr {
		if i > 0 {
			res += ","
		}
		res += f(v)
	}
	res += "]"
	return res
}

// --- Auto-generated FFI wrappers ---
var _Gopurs_Data_Show_ShowArrayImpl = // TAST: (ForAll [a] (Func [(Func [(TypeVar a)] String), (Array (TypeVar a))] String))
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any) string {
		inner_res0 := gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		return gopurs_runtime.Unbox[string](inner_res0)
	}
	go_arg1 := func() []any {
		_source := arg1
		_out := make([]any, gopurs_runtime.ArrayLength(_source))
		for _i := range _out {
			_out[_i] = gopurs_runtime.ArrayAccess(_source, _i)
		}
		return _out
	}()
	go_res := Data_Show_ShowArrayImpl(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_Show_ShowCharImpl = // TAST: (Func [Char] String)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_res := Data_Show_ShowCharImpl(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_Show_ShowIntImpl = // TAST: (Func [Int] String)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int64](arg0)
	go_res := Data_Show_ShowIntImpl(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_Show_ShowNumberImpl = // TAST: (Func [Number] String)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[float64](arg0)
	go_res := Data_Show_ShowNumberImpl(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_Show_ShowStringImpl = // TAST: (Func [String] String)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_res := Data_Show_ShowStringImpl(go_arg0)
	return gopurs_runtime.Box(go_res)
})
