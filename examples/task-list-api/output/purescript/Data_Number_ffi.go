package purescript

import "gopurs/output/gopurs_runtime"

import (
	"math"
	"strconv"
)

func Data_Number_IsNaN(n float64) bool {
	return math.IsNaN(n)
}

func Data_Number_IsFinite(v float64) bool {
	return !math.IsNaN(v) && !math.IsInf(v, 0)
}

func Data_Number_FromStringImpl(str string, isFinite func(float64) bool, just func(float64) interface{}, nothing interface{}) interface{} {
	val, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return nothing
	}
	if math.IsNaN(val) || math.IsInf(val, 0) {
		return nothing
	}
	return just(val)
}

func Data_Number_Abs(n float64) float64 {
	return math.Abs(n)
}

func Data_Number_Acos(n float64) float64 {
	return math.Acos(n)
}

func Data_Number_Asin(n float64) float64 {
	return math.Asin(n)
}

func Data_Number_Atan(n float64) float64 {
	return math.Atan(n)
}

func Data_Number_Atan2(y float64, x float64) float64 {
	return math.Atan2(y, x)
}

func Data_Number_Ceil(n float64) float64 {
	return math.Ceil(n)
}

func Data_Number_Cos(n float64) float64 {
	return math.Cos(n)
}

func Data_Number_Exp(n float64) float64 {
	return math.Exp(n)
}

func Data_Number_Floor(n float64) float64 {
	return math.Floor(n)
}

func Data_Number_Log(n float64) float64 {
	return math.Log(n)
}

func Data_Number_Max(n1 float64, n2 float64) float64 {
	return math.Max(n1, n2)
}

func Data_Number_Min(n1 float64, n2 float64) float64 {
	return math.Min(n1, n2)
}

func Data_Number_Pow(n float64, p float64) float64 {
	return math.Pow(n, p)
}

func Data_Number_Remainder(n float64, m float64) float64 {
	return math.Mod(n, m)
}

func Data_Number_Round(n float64) float64 {
	return math.Round(n)
}

func Data_Number_Sign(v float64) float64 {
	if math.IsNaN(v) || v == 0 {
		return v
	}
	if v < 0 {
		return -1
	}
	return 1
}

func Data_Number_Sin(n float64) float64 {
	return math.Sin(n)
}

func Data_Number_Sqrt(n float64) float64 {
	return math.Sqrt(n)
}

func Data_Number_Tan(n float64) float64 {
	return math.Tan(n)
}

func Data_Number_Trunc(n float64) float64 {
	return math.Trunc(n)
}

var Data_Number_Infinity = math.Inf(1)
var Data_Number_Nan = math.NaN()

// --- Auto-generated FFI wrappers ---
var _Gopurs_Data_Number_Abs = // TAST: (Func [Number] Number)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[float64](arg0)
	go_res := Data_Number_Abs(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_Number_Acos = // TAST: (Func [Number] Number)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[float64](arg0)
	go_res := Data_Number_Acos(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_Number_Asin = // TAST: (Func [Number] Number)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[float64](arg0)
	go_res := Data_Number_Asin(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_Number_Atan = // TAST: (Func [Number] Number)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[float64](arg0)
	go_res := Data_Number_Atan(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_Number_Atan2 = // TAST: (Func [Number, Number] Number)
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[float64](arg0)
	go_arg1 := gopurs_runtime.Unbox[float64](arg1)
	go_res := Data_Number_Atan2(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_Number_Ceil = // TAST: (Func [Number] Number)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[float64](arg0)
	go_res := Data_Number_Ceil(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_Number_Cos = // TAST: (Func [Number] Number)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[float64](arg0)
	go_res := Data_Number_Cos(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_Number_Exp = // TAST: (Func [Number] Number)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[float64](arg0)
	go_res := Data_Number_Exp(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_Number_Floor = // TAST: (Func [Number] Number)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[float64](arg0)
	go_res := Data_Number_Floor(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_Number_FromStringImpl = // TAST: (ADT ["Data","Function","Uncurried","Fn4"] [String, (Func [Number] Boolean), (ForAll [a] (Func [(TypeVar a)] (ADT ["Data","Maybe","Maybe"] [(TypeVar a)]))), (ForAll [a] (ADT ["Data","Maybe","Maybe"] [(TypeVar a)])), (ADT ["Data","Maybe","Maybe"] [Number])])
gopurs_runtime.Func4(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_arg1 := func(p0_0 float64) bool {
		inner_res0 := gopurs_runtime.Apply(arg1, gopurs_runtime.Box(p0_0))
		return gopurs_runtime.Unbox[bool](inner_res0)
	}
	go_arg2 := func(p0_0 float64) any {
		return gopurs_runtime.Apply(arg2, gopurs_runtime.Box(p0_0))
	}
	go_arg3 := arg3
	go_res := Data_Number_FromStringImpl(go_arg0, go_arg1, go_arg2, go_arg3)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_Number_Infinity = // TAST: Number
gopurs_runtime.Box(Data_Number_Infinity)
var _Gopurs_Data_Number_IsFinite = // TAST: (Func [Number] Boolean)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[float64](arg0)
	go_res := Data_Number_IsFinite(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_Number_IsNaN = // TAST: (Func [Number] Boolean)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[float64](arg0)
	go_res := Data_Number_IsNaN(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_Number_Log = // TAST: (Func [Number] Number)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[float64](arg0)
	go_res := Data_Number_Log(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_Number_Max = // TAST: (Func [Number, Number] Number)
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[float64](arg0)
	go_arg1 := gopurs_runtime.Unbox[float64](arg1)
	go_res := Data_Number_Max(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_Number_Min = // TAST: (Func [Number, Number] Number)
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[float64](arg0)
	go_arg1 := gopurs_runtime.Unbox[float64](arg1)
	go_res := Data_Number_Min(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_Number_Nan = // TAST: Number
gopurs_runtime.Box(Data_Number_Nan)
var _Gopurs_Data_Number_Pow = // TAST: (Func [Number, Number] Number)
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[float64](arg0)
	go_arg1 := gopurs_runtime.Unbox[float64](arg1)
	go_res := Data_Number_Pow(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_Number_Remainder = // TAST: (Func [Number, Number] Number)
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[float64](arg0)
	go_arg1 := gopurs_runtime.Unbox[float64](arg1)
	go_res := Data_Number_Remainder(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_Number_Round = // TAST: (Func [Number] Number)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[float64](arg0)
	go_res := Data_Number_Round(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_Number_Sign = // TAST: (Func [Number] Number)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[float64](arg0)
	go_res := Data_Number_Sign(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_Number_Sin = // TAST: (Func [Number] Number)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[float64](arg0)
	go_res := Data_Number_Sin(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_Number_Sqrt = // TAST: (Func [Number] Number)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[float64](arg0)
	go_res := Data_Number_Sqrt(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_Number_Tan = // TAST: (Func [Number] Number)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[float64](arg0)
	go_res := Data_Number_Tan(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_Number_Trunc = // TAST: (Func [Number] Number)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[float64](arg0)
	go_res := Data_Number_Trunc(go_arg0)
	return gopurs_runtime.Box(go_res)
})

// --- Boxed-result FFI workers ---
func Data_Number_FromStringImpl_nativeWorker(arg0 string, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 any) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := func(p0_0 float64) bool {
		inner_res0 := gopurs_runtime.Apply(arg1, gopurs_runtime.Box(p0_0))
		return gopurs_runtime.Unbox[bool](inner_res0)
	}
	go_arg2 := func(p0_0 float64) any {
		return gopurs_runtime.Apply(arg2, gopurs_runtime.Box(p0_0))
	}
	go_arg3 := arg3
	go_res := Data_Number_FromStringImpl(go_arg0, go_arg1, go_arg2, go_arg3)
	return gopurs_runtime.Box(go_res)
}
