package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_Proxy gopurs_runtime.Value
var once_Main_Proxy sync.Once

func Get_Main_Proxy() gopurs_runtime.Value {
	once_Main_Proxy.Do(func() {
		cache_Main_Proxy = gopurs_runtime.Value{Type: 9, IntVal: int64(227768594), UnsafePtr: nil}
	})
	return cache_Main_Proxy
}

var cache_Main_testToString gopurs_runtime.Value
var once_Main_testToString sync.Once

func Get_Main_testToString() gopurs_runtime.Value {
	once_Main_testToString.Do(func() {
		cache_Main_testToString = gopurs_runtime.Func2(func(_dollar___unused_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_testToString(_dollar___unused_0_box, uint32(v_1_box.IntVal))), UnsafePtr: nil}
		})
	})
	return cache_Main_testToString
}

var cache_Main_testToString__762614295 gopurs_runtime.Value
var once_Main_testToString__762614295 sync.Once

func Get_Main_testToString__762614295() gopurs_runtime.Value {
	once_Main_testToString__762614295.Do(func() {
		cache_Main_testToString__762614295 = gopurs_runtime.Func2(func(_dollar___unused_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_testToString__762614295(_dollar___unused_0_box, uint32(v_1_box.IntVal))), UnsafePtr: nil}
		})
	})
	return cache_Main_testToString__762614295
}

var cache_Main_zeroToString gopurs_runtime.Value
var once_Main_zeroToString sync.Once

func Get_Main_zeroToString() gopurs_runtime.Value {
	once_Main_zeroToString.Do(func() {
		cache_Main_zeroToString = gopurs_runtime.Value{Type: 9, IntVal: int64(227768594), UnsafePtr: nil}
	})
	return cache_Main_zeroToString
}

var cache_Main_zeroToStringTA gopurs_runtime.Value
var once_Main_zeroToStringTA sync.Once

func Get_Main_zeroToStringTA() gopurs_runtime.Value {
	once_Main_zeroToStringTA.Do(func() {
		cache_Main_zeroToStringTA = gopurs_runtime.Value{Type: 9, IntVal: int64(227768594), UnsafePtr: nil}
	})
	return cache_Main_zeroToStringTA
}

var cache_Main_posToStringTA gopurs_runtime.Value
var once_Main_posToStringTA sync.Once

func Get_Main_posToStringTA() gopurs_runtime.Value {
	once_Main_posToStringTA.Do(func() {
		cache_Main_posToStringTA = gopurs_runtime.Value{Type: 9, IntVal: int64(227768594), UnsafePtr: nil}
	})
	return cache_Main_posToStringTA
}

var cache_Main_posToString gopurs_runtime.Value
var once_Main_posToString sync.Once

func Get_Main_posToString() gopurs_runtime.Value {
	once_Main_posToString.Do(func() {
		cache_Main_posToString = gopurs_runtime.Value{Type: 9, IntVal: int64(227768594), UnsafePtr: nil}
	})
	return cache_Main_posToString
}

var cache_Main_negToStringTA gopurs_runtime.Value
var once_Main_negToStringTA sync.Once

func Get_Main_negToStringTA() gopurs_runtime.Value {
	once_Main_negToStringTA.Do(func() {
		cache_Main_negToStringTA = gopurs_runtime.Value{Type: 9, IntVal: int64(227768594), UnsafePtr: nil}
	})
	return cache_Main_negToStringTA
}

var cache_Main_negToString gopurs_runtime.Value
var once_Main_negToString sync.Once

func Get_Main_negToString() gopurs_runtime.Value {
	once_Main_negToString.Do(func() {
		cache_Main_negToString = gopurs_runtime.Value{Type: 9, IntVal: int64(227768594), UnsafePtr: nil}
	})
	return cache_Main_negToString
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

var cache_Main_intMul gopurs_runtime.Value
var once_Main_intMul sync.Once

func Get_Main_intMul() gopurs_runtime.Value {
	once_Main_intMul.Do(func() {
		cache_Main_intMul = gopurs_runtime.Func3(func(_dollar___unused_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, v1_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_intMul(_dollar___unused_0_box, uint32(v_1_box.IntVal), uint32(v1_2_box.IntVal))), UnsafePtr: nil}
		})
	})
	return cache_Main_intMul
}

var cache_Main_intMul__3629321662 gopurs_runtime.Value
var once_Main_intMul__3629321662 sync.Once

func Get_Main_intMul__3629321662() gopurs_runtime.Value {
	once_Main_intMul__3629321662.Do(func() {
		cache_Main_intMul__3629321662 = gopurs_runtime.Func3(func(_dollar___unused_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, v1_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_intMul__3629321662(_dollar___unused_0_box, uint32(v_1_box.IntVal), uint32(v1_2_box.IntVal))), UnsafePtr: nil}
		})
	})
	return cache_Main_intMul__3629321662
}

var cache_Main_testMul gopurs_runtime.Value
var once_Main_testMul sync.Once

func Get_Main_testMul() gopurs_runtime.Value {
	once_Main_testMul.Do(func() {
		cache_Main_testMul = gopurs_runtime.Value{Type: 9, IntVal: int64(227768594), UnsafePtr: nil}
	})
	return cache_Main_testMul
}

var cache_Main_intAdd gopurs_runtime.Value
var once_Main_intAdd sync.Once

func Get_Main_intAdd() gopurs_runtime.Value {
	once_Main_intAdd.Do(func() {
		cache_Main_intAdd = gopurs_runtime.Func3(func(_dollar___unused_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, v1_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_intAdd(_dollar___unused_0_box, uint32(v_1_box.IntVal), uint32(v1_2_box.IntVal))), UnsafePtr: nil}
		})
	})
	return cache_Main_intAdd
}

var cache_Main_intAdd__2918969323 gopurs_runtime.Value
var once_Main_intAdd__2918969323 sync.Once

func Get_Main_intAdd__2918969323() gopurs_runtime.Value {
	once_Main_intAdd__2918969323.Do(func() {
		cache_Main_intAdd__2918969323 = gopurs_runtime.Func3(func(_dollar___unused_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, v1_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_intAdd__2918969323(_dollar___unused_0_box, uint32(v_1_box.IntVal), uint32(v1_2_box.IntVal))), UnsafePtr: nil}
		})
	})
	return cache_Main_intAdd__2918969323
}

var cache_Main_testAdd gopurs_runtime.Value
var once_Main_testAdd sync.Once

func Get_Main_testAdd() gopurs_runtime.Value {
	once_Main_testAdd.Do(func() {
		cache_Main_testAdd = gopurs_runtime.Value{Type: 9, IntVal: int64(227768594), UnsafePtr: nil}
	})
	return cache_Main_testAdd
}

var cache_Main_testAddMul gopurs_runtime.Value
var once_Main_testAddMul sync.Once

func Get_Main_testAddMul() gopurs_runtime.Value {
	once_Main_testAddMul.Do(func() {
		cache_Main_testAddMul = gopurs_runtime.Value{Type: 9, IntVal: int64(227768594), UnsafePtr: nil}
	})
	return cache_Main_testAddMul
}

var cache_Main_testMulAdd gopurs_runtime.Value
var once_Main_testMulAdd sync.Once

func Get_Main_testMulAdd() gopurs_runtime.Value {
	once_Main_testMulAdd.Do(func() {
		cache_Main_testMulAdd = gopurs_runtime.Value{Type: 9, IntVal: int64(227768594), UnsafePtr: nil}
	})
	return cache_Main_testMulAdd
}

var cache_Main__maxInt gopurs_runtime.Value
var once_Main__maxInt sync.Once

func Get_Main__maxInt() gopurs_runtime.Value {
	once_Main__maxInt.Do(func() {
		cache_Main__maxInt = gopurs_runtime.Value{Type: 9, IntVal: int64(227768594), UnsafePtr: nil}
	})
	return cache_Main__maxInt
}

var cache_Main_testBeyondMax gopurs_runtime.Value
var once_Main_testBeyondMax sync.Once

func Get_Main_testBeyondMax() gopurs_runtime.Value {
	once_Main_testBeyondMax.Do(func() {
		cache_Main_testBeyondMax = gopurs_runtime.Value{Type: 9, IntVal: int64(227768594), UnsafePtr: nil}
	})
	return cache_Main_testBeyondMax
}

var cache_Main_testMax gopurs_runtime.Value
var once_Main_testMax sync.Once

func Get_Main_testMax() gopurs_runtime.Value {
	once_Main_testMax.Do(func() {
		cache_Main_testMax = gopurs_runtime.Value{Type: 9, IntVal: int64(227768594), UnsafePtr: nil}
	})
	return cache_Main_testMax
}

type Constructor_Main_Proxy[T_a any] struct {
	Rc uint32
}

func Call_Main_testToString(_dollar___unused_0_loop gopurs_runtime.Value, v_1_loop uint32) uint32 {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var v_1 uint32 = v_1_loop
	_ = v_1
	return 227768594
}

func Call_Main_testToString__762614295(_dollar___unused_0_loop gopurs_runtime.Value, v_1_loop uint32) uint32 {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var v_1 uint32 = v_1_loop
	_ = v_1
	return 227768594
}

func Call_Main_intMul(_dollar___unused_0_loop gopurs_runtime.Value, v_1_loop uint32, v1_2_loop uint32) uint32 {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var v_1 uint32 = v_1_loop
	_ = v_1
	var v1_2 uint32 = v1_2_loop
	_ = v1_2
	return 227768594
}

func Call_Main_intMul__3629321662(_dollar___unused_0_loop gopurs_runtime.Value, v_1_loop uint32, v1_2_loop uint32) uint32 {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var v_1 uint32 = v_1_loop
	_ = v_1
	var v1_2 uint32 = v1_2_loop
	_ = v1_2
	return 227768594
}

func Call_Main_intAdd(_dollar___unused_0_loop gopurs_runtime.Value, v_1_loop uint32, v1_2_loop uint32) uint32 {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var v_1 uint32 = v_1_loop
	_ = v_1
	var v1_2 uint32 = v1_2_loop
	_ = v1_2
	return 227768594
}

func Call_Main_intAdd__2918969323(_dollar___unused_0_loop gopurs_runtime.Value, v_1_loop uint32, v1_2_loop uint32) uint32 {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var v_1 uint32 = v_1_loop
	_ = v_1
	var v1_2 uint32 = v1_2_loop
	_ = v1_2
	return 227768594
}
