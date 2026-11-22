package purescript

import "gopurs/output/gopurs_runtime"

import (
	"time"
)

func createUTC(y, mo, d, h, m, s, ms int) time.Time {
	return time.Date(y, time.Month(mo+1), d, h, m, s, ms*1000000, time.UTC)
}

func getInt(m map[string]interface{}, key string) int {
	v := m[key]
	if val, ok := v.(int64); ok {
		return int(val)
	}
	if val, ok := v.(int); ok {
		return val
	}
	if val, ok := v.(float64); ok {
		return int(val)
	}
	// Fallback for gopurs_runtime.Value
	return int(v.(gopurs_runtime.Value).IntVal)
}

func Data_DateTime_CalcDiff(rec1 map[string]interface{}, rec2 map[string]interface{}) float64 {
	msUTC1 := createUTC(getInt(rec1, "year"), getInt(rec1, "month")-1, getInt(rec1, "day"), getInt(rec1, "hour"), getInt(rec1, "minute"), getInt(rec1, "second"), getInt(rec1, "millisecond")).UnixMilli()
	msUTC2 := createUTC(getInt(rec2, "year"), getInt(rec2, "month")-1, getInt(rec2, "day"), getInt(rec2, "hour"), getInt(rec2, "minute"), getInt(rec2, "second"), getInt(rec2, "millisecond")).UnixMilli()
	return float64(msUTC1 - msUTC2)
}

func Data_DateTime_AdjustImpl(just func(interface{}) interface{}, nothing interface{}, offset float64, rec map[string]interface{}) interface{} {
	t := createUTC(getInt(rec, "year"), getInt(rec, "month")-1, getInt(rec, "day"), getInt(rec, "hour"), getInt(rec, "minute"), getInt(rec, "second"), getInt(rec, "millisecond"))
	ms := t.UnixMilli() + int64(offset)
	dt := time.UnixMilli(ms).UTC()

	resMap := make(map[string]interface{})
	resMap["year"] = dt.Year()
	resMap["month"] = int(dt.Month())
	resMap["day"] = dt.Day()
	resMap["hour"] = dt.Hour()
	resMap["minute"] = dt.Minute()
	resMap["second"] = dt.Second()
	resMap["millisecond"] = dt.Nanosecond() / 1000000

	return just(resMap)
}

// --- Auto-generated FFI wrappers ---
var _Gopurs_Data_DateTime_AdjustImpl = // TAST: (Func [(ForAll [a] (Func [(TypeVar a)] (ADT ["Data","Maybe","Maybe"] [(TypeVar a)]))), (ForAll [a] (ADT ["Data","Maybe","Maybe"] [(TypeVar a)])), Number, (Record (Row [year: Int, month: Int, day: Int, hour: Int, minute: Int, second: Int, millisecond: Int] Empty))] (ADT ["Data","Maybe","Maybe"] [(Record (Row [year: Int, month: Int, day: Int, hour: Int, minute: Int, second: Int, millisecond: Int] Empty))]))
gopurs_runtime.Func4(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any) any {
		return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
	}
	go_arg1 := arg1
	go_arg2 := gopurs_runtime.Unbox[float64](arg2)
	arg3_map := gopurs_runtime.RecordToMap(arg3)
	go_arg3 := make(map[string]any)
	for k, v := range arg3_map {
		go_arg3[k] = v
	}
	go_res := Data_DateTime_AdjustImpl(go_arg0, go_arg1, go_arg2, go_arg3)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_DateTime_CalcDiff = // TAST: (ADT ["Data","Function","Uncurried","Fn2"] [(Record (Row [year: Int, month: Int, day: Int, hour: Int, minute: Int, second: Int, millisecond: Int] Empty)), (Record (Row [year: Int, month: Int, day: Int, hour: Int, minute: Int, second: Int, millisecond: Int] Empty)), Number])
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	arg0_map := gopurs_runtime.RecordToMap(arg0)
	go_arg0 := make(map[string]any)
	for k, v := range arg0_map {
		go_arg0[k] = v
	}
	arg1_map := gopurs_runtime.RecordToMap(arg1)
	go_arg1 := make(map[string]any)
	for k, v := range arg1_map {
		go_arg1[k] = v
	}
	go_res := Data_DateTime_CalcDiff(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})

// --- Boxed-result FFI workers ---
func Data_DateTime_AdjustImpl_nativeWorker(arg0 gopurs_runtime.Value, arg1 any, arg2 float64, arg3 map[string]any) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any) any {
		return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
	}
	go_arg1 := arg1
	go_arg2 := arg2
	go_arg3 := arg3
	go_res := Data_DateTime_AdjustImpl(go_arg0, go_arg1, go_arg2, go_arg3)
	return gopurs_runtime.Box(go_res)
}
