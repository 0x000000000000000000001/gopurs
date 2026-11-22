package purescript

import "gopurs/output/gopurs_runtime"

import (
	"time"
)

func createDateTime(y, m, d, h, mi, s, ms int) time.Time {
	return time.Date(y, time.Month(m+1), d, h, mi, s, ms*1000000, time.UTC)
}

func Data_DateTime_Instant_FromDateTimeImpl(y int, mo int, d int, h int, mi int, s int, ms int) float64 {
	dt := createDateTime(y, mo-1, d, h, mi, s, ms)
	return float64(dt.UnixMilli())
}

func Data_DateTime_Instant_ToDateTimeImpl(ctor func(int) func(int) func(int) func(int) func(int) func(int) func(int) interface{}, instant float64) interface{} {
	dt := time.UnixMilli(int64(instant)).UTC()

	return ctor(dt.Year())(int(dt.Month()))(dt.Day())(dt.Hour())(dt.Minute())(dt.Second())(dt.Nanosecond() / 1000000)
}

// --- Auto-generated FFI wrappers ---
var _Gopurs_Data_DateTime_Instant_FromDateTimeImpl = // TAST: (ADT ["Data","Function","Uncurried","Fn7"] [Int, Int, Int, Int, Int, Int, Int, Number])
gopurs_runtime.Func7(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value, arg5 gopurs_runtime.Value, arg6 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int](arg0)
	go_arg1 := gopurs_runtime.Unbox[int](arg1)
	go_arg2 := gopurs_runtime.Unbox[int](arg2)
	go_arg3 := gopurs_runtime.Unbox[int](arg3)
	go_arg4 := gopurs_runtime.Unbox[int](arg4)
	go_arg5 := gopurs_runtime.Unbox[int](arg5)
	go_arg6 := gopurs_runtime.Unbox[int](arg6)
	go_res := Data_DateTime_Instant_FromDateTimeImpl(go_arg0, go_arg1, go_arg2, go_arg3, go_arg4, go_arg5, go_arg6)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_DateTime_Instant_ToDateTimeImpl = // TAST: (Func [(Func [Int, Int, Int, Int, Int, Int, Int] (ADT ["Data","DateTime","DateTime"] [])), Number] (ADT ["Data","DateTime","DateTime"] []))
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 int) func(int) func(int) func(int) func(int) func(int) func(int) any {
		inner_res0 := gopurs_runtime.Apply(arg0, gopurs_runtime.Int(int64(p0_0)))
		return func(p1_0 int) func(int) func(int) func(int) func(int) func(int) any {
			inner_res1 := gopurs_runtime.Apply(inner_res0, gopurs_runtime.Int(int64(p1_0)))
			return func(p2_0 int) func(int) func(int) func(int) func(int) any {
				inner_res2 := gopurs_runtime.Apply(inner_res1, gopurs_runtime.Int(int64(p2_0)))
				return func(p3_0 int) func(int) func(int) func(int) any {
					inner_res3 := gopurs_runtime.Apply(inner_res2, gopurs_runtime.Int(int64(p3_0)))
					return func(p4_0 int) func(int) func(int) any {
						inner_res4 := gopurs_runtime.Apply(inner_res3, gopurs_runtime.Int(int64(p4_0)))
						return func(p5_0 int) func(int) any {
							inner_res5 := gopurs_runtime.Apply(inner_res4, gopurs_runtime.Int(int64(p5_0)))
							return func(p6_0 int) any {
								return gopurs_runtime.Apply(inner_res5, gopurs_runtime.Int(int64(p6_0)))
							}
						}
					}
				}
			}
		}
	}
	go_arg1 := gopurs_runtime.Unbox[float64](arg1)
	go_res := Data_DateTime_Instant_ToDateTimeImpl(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})

// --- Boxed-result FFI workers ---
func Data_DateTime_Instant_ToDateTimeImpl_nativeWorker(arg0 gopurs_runtime.Value, arg1 float64) gopurs_runtime.Value {
	go_arg0 := func(p0_0 int) func(int) func(int) func(int) func(int) func(int) func(int) any {
		inner_res0 := gopurs_runtime.Apply(arg0, gopurs_runtime.Int(int64(p0_0)))
		return func(p1_0 int) func(int) func(int) func(int) func(int) func(int) any {
			inner_res1 := gopurs_runtime.Apply(inner_res0, gopurs_runtime.Int(int64(p1_0)))
			return func(p2_0 int) func(int) func(int) func(int) func(int) any {
				inner_res2 := gopurs_runtime.Apply(inner_res1, gopurs_runtime.Int(int64(p2_0)))
				return func(p3_0 int) func(int) func(int) func(int) any {
					inner_res3 := gopurs_runtime.Apply(inner_res2, gopurs_runtime.Int(int64(p3_0)))
					return func(p4_0 int) func(int) func(int) any {
						inner_res4 := gopurs_runtime.Apply(inner_res3, gopurs_runtime.Int(int64(p4_0)))
						return func(p5_0 int) func(int) any {
							inner_res5 := gopurs_runtime.Apply(inner_res4, gopurs_runtime.Int(int64(p5_0)))
							return func(p6_0 int) any {
								return gopurs_runtime.Apply(inner_res5, gopurs_runtime.Int(int64(p6_0)))
							}
						}
					}
				}
			}
		}
	}
	go_arg1 := arg1
	go_res := Data_DateTime_Instant_ToDateTimeImpl(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
}
