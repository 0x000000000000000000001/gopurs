package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_Z gopurs_runtime.Value
var once_Main_Z sync.Once

func Get_Main_Z() gopurs_runtime.Value {
	once_Main_Z.Do(func() {
		cache_Main_Z = gopurs_runtime.Value{Type: 9, IntVal: int64(1714575428), UnsafePtr: nil}
	})
	return cache_Main_Z
}

var cache_Main_S gopurs_runtime.Value
var once_Main_S sync.Once

func Get_Main_S() gopurs_runtime.Value {
	once_Main_S.Do(func() {
		cache_Main_S = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return value0
		})
	})
	return cache_Main_S
}

var cache_Main_EQ_dollar_Dict gopurs_runtime.Value
var once_Main_EQ_dollar_Dict sync.Once

func Get_Main_EQ_dollar_Dict() gopurs_runtime.Value {
	once_Main_EQ_dollar_Dict.Do(func() {
		cache_Main_EQ_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_EQ_dollar_Dict(x_0_box)
		})
	})
	return cache_Main_EQ_dollar_Dict
}

var cache_Main_test gopurs_runtime.Value
var once_Main_test sync.Once

func Get_Main_test() gopurs_runtime.Value {
	once_Main_test.Do(func() {
		cache_Main_test = gopurs_runtime.Func3(func(_dollar___unused_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_test(uint32(_dollar___unused_0_box.IntVal), a_1_box, v_2_box)
		})
	})
	return cache_Main_test
}

var cache_Main_spin gopurs_runtime.Value
var once_Main_spin sync.Once

func Get_Main_spin() gopurs_runtime.Value {
	once_Main_spin.Do(func() {
		cache_Main_spin = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_spin(a_0_box)
		})
	})
	return cache_Main_spin
}

var cache_Main_test1 gopurs_runtime.Value
var once_Main_test1 sync.Once

func Get_Main_test1() gopurs_runtime.Value {
	once_Main_test1.Do(func() {
		cache_Main_test1 = gopurs_runtime.Func(func(dictEQ_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_test1(dictEQ_0_box)
		})
	})
	return cache_Main_test1
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

var cache_Main_eqT gopurs_runtime.Value
var once_Main_eqT sync.Once

func Get_Main_eqT() gopurs_runtime.Value {
	once_Main_eqT.Do(func() {
		cache_Main_eqT = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.RecordDict0().IntVal)), UnsafePtr: nil}
	})
	return cache_Main_eqT
}

var cache_Main_eqF gopurs_runtime.Value
var once_Main_eqF sync.Once

func Get_Main_eqF() gopurs_runtime.Value {
	once_Main_eqF.Do(func() {
		cache_Main_eqF = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.RecordDict0().IntVal)), UnsafePtr: nil}
	})
	return cache_Main_eqF
}

var cache_Main_spin__669605629 gopurs_runtime.Value
var once_Main_spin__669605629 sync.Once

func Get_Main_spin__669605629() gopurs_runtime.Value {
	once_Main_spin__669605629.Do(func() {
		cache_Main_spin__669605629 = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_spin__669605629(a_0_box)
		})
	})
	return cache_Main_spin__669605629
}

var cache_Main_test__3969149901 gopurs_runtime.Value
var once_Main_test__3969149901 sync.Once

func Get_Main_test__3969149901() gopurs_runtime.Value {
	once_Main_test__3969149901.Do(func() {
		cache_Main_test__3969149901 = gopurs_runtime.Func3(func(_dollar___unused_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_test__3969149901(uint32(_dollar___unused_0_box.IntVal), a_1_box, v_2_box)
		})
	})
	return cache_Main_test__3969149901
}

type Constructor_Main_Z struct {
	Rc uint32
}

type Constructor_Main_S struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

type Constructor_Main_EQ struct {
	Rc uint32
}

func init() {
	gopurs_runtime.StructGetters[3323825930] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_EQ)(ptr)
		_ = c
		switch key {

		default:
			panic("Key not found in dictionary Constructor_Main_EQ: " + key)
		}
	}
}

func Call_Main_EQ_dollar_Dict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_test(_dollar___unused_0_loop uint32, a_1_loop gopurs_runtime.Value, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var _dollar___unused_0 uint32 = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var a_1 gopurs_runtime.Value = a_1_loop
	_ = a_1
	var v_2 gopurs_runtime.Value = v_2_loop
	_ = v_2
	return a_1
}

func Call_Main_spin(a_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
spin:
	for {
		if false {
			continue spin
		}
		var a_0 gopurs_runtime.Value = a_0_loop
		_ = a_0
		a_0_loop = a_0
		continue spin
		return gopurs_runtime.Value{}
	}
}

func Call_Main_test1(dictEQ_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictEQ_0 gopurs_runtime.Value = dictEQ_0_loop
	_ = dictEQ_0
	return Call_Main_spin(gopurs_runtime.Int(1))
}

func Call_Main_spin__669605629(a_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var a_0 gopurs_runtime.Value = a_0_loop
	_ = a_0
	return Call_Main_spin(a_0)
}

func Call_Main_test__3969149901(_dollar___unused_0_loop uint32, a_1_loop gopurs_runtime.Value, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var _dollar___unused_0 uint32 = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var a_1 gopurs_runtime.Value = a_1_loop
	_ = a_1
	var v_2 gopurs_runtime.Value = v_2_loop
	_ = v_2
	return a_1
}
