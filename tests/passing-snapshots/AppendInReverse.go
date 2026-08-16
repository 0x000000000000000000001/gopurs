package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_Balanced_dollar_Dict gopurs_runtime.Value
var once_Main_Balanced_dollar_Dict sync.Once

func Get_Main_Balanced_dollar_Dict() gopurs_runtime.Value {
	once_Main_Balanced_dollar_Dict.Do(func() {
		cache_Main_Balanced_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_Balanced_dollar_Dict(x_0_box)
		})
	})
	return cache_Main_Balanced_dollar_Dict
}

var cache_Main_balanced2 gopurs_runtime.Value
var once_Main_balanced2 sync.Once

func Get_Main_balanced2() gopurs_runtime.Value {
	once_Main_balanced2.Do(func() {
		cache_Main_balanced2 = gopurs_runtime.Func3(func(_dollar___unused_0_box gopurs_runtime.Value, _dollar___unused_1_box gopurs_runtime.Value, _dollar___unused_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_balanced2(_dollar___unused_0_box, _dollar___unused_1_box, _dollar___unused_2_box)
		})
	})
	return cache_Main_balanced2
}

var cache_Main_balanced1 gopurs_runtime.Value
var once_Main_balanced1 sync.Once

func Get_Main_balanced1() gopurs_runtime.Value {
	once_Main_balanced1.Do(func() {
		cache_Main_balanced1 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.RecordDict0().IntVal)), UnsafePtr: nil}
	})
	return cache_Main_balanced1
}

var cache_Main_balanced gopurs_runtime.Value
var once_Main_balanced sync.Once

func Get_Main_balanced() gopurs_runtime.Value {
	once_Main_balanced.Do(func() {
		cache_Main_balanced = gopurs_runtime.Func2(func(_dollar___unused_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_balanced(uint32(_dollar___unused_0_box.IntVal), uint32(v_1_box.IntVal)))
		})
	})
	return cache_Main_balanced
}

var cache_Main_b3 gopurs_runtime.Value
var once_Main_b3 sync.Once

func Get_Main_b3() gopurs_runtime.Value {
	once_Main_b3.Do(func() {
		cache_Main_b3 = gopurs_runtime.Str("ok")
	})
	return cache_Main_b3
}

var cache_Main_b2 gopurs_runtime.Value
var once_Main_b2 sync.Once

func Get_Main_b2() gopurs_runtime.Value {
	once_Main_b2.Do(func() {
		cache_Main_b2 = gopurs_runtime.Str("ok")
	})
	return cache_Main_b2
}

var cache_Main_b1 gopurs_runtime.Value
var once_Main_b1 sync.Once

func Get_Main_b1() gopurs_runtime.Value {
	once_Main_b1.Do(func() {
		cache_Main_b1 = gopurs_runtime.Str("ok")
	})
	return cache_Main_b1
}

var cache_Main_b0 gopurs_runtime.Value
var once_Main_b0 sync.Once

func Get_Main_b0() gopurs_runtime.Value {
	once_Main_b0.Do(func() {
		cache_Main_b0 = gopurs_runtime.Str("ok")
	})
	return cache_Main_b0
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 -> gopurs_runtime.Value
			__local_var_0_0 := gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("ok"))
			_ = __local_var_0_0
			_dollar___unused_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = _dollar___unused_1_1
			_dollar___unused_2_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("ok")), gopurs_runtime.Value{})
			_ = _dollar___unused_2_2
			_dollar___unused_3_3 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("ok")), gopurs_runtime.Value{})
			_ = _dollar___unused_3_3
			_dollar___unused_4_4 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("ok")), gopurs_runtime.Value{})
			_ = _dollar___unused_4_4
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

var cache_Main_balanced__4249458811 gopurs_runtime.Value
var once_Main_balanced__4249458811 sync.Once

func Get_Main_balanced__4249458811() gopurs_runtime.Value {
	once_Main_balanced__4249458811.Do(func() {
		cache_Main_balanced__4249458811 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_balanced__4249458811(uint32(v_0_box.IntVal)))
		})
	})
	return cache_Main_balanced__4249458811
}

var cache_Main_balanced__2760328189 gopurs_runtime.Value
var once_Main_balanced__2760328189 sync.Once

func Get_Main_balanced__2760328189() gopurs_runtime.Value {
	once_Main_balanced__2760328189.Do(func() {
		cache_Main_balanced__2760328189 = gopurs_runtime.Func2(func(_dollar___unused_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_balanced__2760328189(uint32(_dollar___unused_0_box.IntVal), uint32(v_1_box.IntVal)))
		})
	})
	return cache_Main_balanced__2760328189
}

type Constructor_Main_Balanced struct {
	Rc uint32
}

func init() {
	gopurs_runtime.StructGetters[2947706876] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_Balanced)(ptr)
		_ = c
		switch key {

		default:
			panic("Key not found in dictionary Constructor_Main_Balanced: " + key)
		}
	}
}

func Call_Main_Balanced_dollar_Dict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_balanced2(_dollar___unused_0_loop gopurs_runtime.Value, _dollar___unused_1_loop gopurs_runtime.Value, _dollar___unused_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var _dollar___unused_1 gopurs_runtime.Value = _dollar___unused_1_loop
	_ = _dollar___unused_1
	var _dollar___unused_2 gopurs_runtime.Value = _dollar___unused_2_loop
	_ = _dollar___unused_2
	return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.RecordDict0().IntVal)), UnsafePtr: nil}
}

func Call_Main_balanced(_dollar___unused_0_loop uint32, v_1_loop uint32) string {
	var _dollar___unused_0 uint32 = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var v_1 uint32 = v_1_loop
	_ = v_1
	return "ok"
}

func Call_Main_balanced__4249458811(v_0_loop uint32) string {
	var v_0 uint32 = v_0_loop
	_ = v_0
	return "ok"
}

func Call_Main_balanced__2760328189(_dollar___unused_0_loop uint32, v_1_loop uint32) string {
	var _dollar___unused_0 uint32 = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var v_1 uint32 = v_1_loop
	_ = v_1
	return "ok"
}
