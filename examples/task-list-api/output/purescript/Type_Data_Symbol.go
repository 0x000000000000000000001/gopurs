package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Type_Data_Symbol_Equals_dollar_Dict gopurs_runtime.Value
var once_Type_Data_Symbol_Equals_dollar_Dict sync.Once

func Get_Type_Data_Symbol_Equals_dollar_Dict() gopurs_runtime.Value {
	once_Type_Data_Symbol_Equals_dollar_Dict.Do(func() {
		cache_Type_Data_Symbol_Equals_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Type_Data_Symbol_Equals_dollar_Dict(func() struct {

				/*gopurs:record-v1:0:*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {

					/*gopurs:record-v1:0:*/
				}{}

				return clone
			}())
		})
	})
	return cache_Type_Data_Symbol_Equals_dollar_Dict
}

var cache_Type_Data_Symbol_Equals_dollar_Dict1 gopurs_runtime.Value
var once_Type_Data_Symbol_Equals_dollar_Dict1 sync.Once

func Get_Type_Data_Symbol_Equals_dollar_Dict1() gopurs_runtime.Value {
	once_Type_Data_Symbol_Equals_dollar_Dict1.Do(func() {
		cache_Type_Data_Symbol_Equals_dollar_Dict1 = func() gopurs_runtime.Value {
			orig := struct {

				/*gopurs:record-v1:0:*/
			}{}
			_ = orig
			return gopurs_runtime.RecordDict0()
		}()
	})
	return cache_Type_Data_Symbol_Equals_dollar_Dict1
}

var cache_Type_Data_Symbol_uncons gopurs_runtime.Value
var once_Type_Data_Symbol_uncons sync.Once

func Get_Type_Data_Symbol_uncons() gopurs_runtime.Value {
	once_Type_Data_Symbol_uncons.Do(func() {
		cache_Type_Data_Symbol_uncons = gopurs_runtime.Func2(func(_dollar___unused_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				orig := Call_Type_Data_Symbol_uncons(_dollar___unused_0_box, uint32(v_1_box.IntVal))
				_ = orig
				return gopurs_runtime.RecordDict2("head", "tail", gopurs_runtime.Value{Type: 9, IntVal: int64(orig.head), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(orig.tail), UnsafePtr: nil})
			}()
		})
	})
	return cache_Type_Data_Symbol_uncons
}

var cache_Type_Data_Symbol_equalsSymbol gopurs_runtime.Value
var once_Type_Data_Symbol_equalsSymbol sync.Once

func Get_Type_Data_Symbol_equalsSymbol() gopurs_runtime.Value {
	once_Type_Data_Symbol_equalsSymbol.Do(func() {
		cache_Type_Data_Symbol_equalsSymbol = gopurs_runtime.Func2(func(_dollar___unused_0_box gopurs_runtime.Value, _dollar___unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Type_Data_Symbol_equalsSymbol(_dollar___unused_0_box, _dollar___unused_1_box)
		})
	})
	return cache_Type_Data_Symbol_equalsSymbol
}

var cache_Type_Data_Symbol_equals gopurs_runtime.Value
var once_Type_Data_Symbol_equals sync.Once

func Get_Type_Data_Symbol_equals() gopurs_runtime.Value {
	once_Type_Data_Symbol_equals.Do(func() {
		cache_Type_Data_Symbol_equals = gopurs_runtime.Func3(func(_dollar___unused_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, v1_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Type_Data_Symbol_equals(_dollar___unused_0_box, uint32(v_1_box.IntVal), uint32(v1_2_box.IntVal))), UnsafePtr: nil}
		})
	})
	return cache_Type_Data_Symbol_equals
}

var cache_Type_Data_Symbol_compare gopurs_runtime.Value
var once_Type_Data_Symbol_compare sync.Once

func Get_Type_Data_Symbol_compare() gopurs_runtime.Value {
	once_Type_Data_Symbol_compare.Do(func() {
		cache_Type_Data_Symbol_compare = gopurs_runtime.Func3(func(_dollar___unused_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, v1_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Type_Data_Symbol_compare(_dollar___unused_0_box, uint32(v_1_box.IntVal), uint32(v1_2_box.IntVal))), UnsafePtr: nil}
		})
	})
	return cache_Type_Data_Symbol_compare
}

var cache_Type_Data_Symbol_go__append gopurs_runtime.Value
var once_Type_Data_Symbol_go__append sync.Once

func Get_Type_Data_Symbol_go__append() gopurs_runtime.Value {
	once_Type_Data_Symbol_go__append.Do(func() {
		cache_Type_Data_Symbol_go__append = gopurs_runtime.Func3(func(_dollar___unused_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, v1_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Type_Data_Symbol_go__append(_dollar___unused_0_box, uint32(v_1_box.IntVal), uint32(v1_2_box.IntVal))), UnsafePtr: nil}
		})
	})
	return cache_Type_Data_Symbol_go__append
}

type Constructor_Type_Data_Symbol_Equals[T_lhs any, T_rhs any, T_out any] struct {
	Rc uint32
}

//gopurs:init-v1:00670065006e006500720061007400650064002f0054007900700065005f0044006100740061005f00530079006d0062006f006c:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f0054007900700065005f0044006100740061005f00530079006d0062006f006c005f0045007100750061006c0073
func init() {
	gopurs_runtime.StructGetters[3155987556] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Type_Data_Symbol_Equals[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(ptr)
		_ = c
		switch key {

		default:
			panic("Key not found in dictionary Constructor_Type_Data_Symbol_Equals: " + key)
		}
	}
}

func Call_Type_Data_Symbol_Equals_dollar_Dict(x_0_loop struct {

	/*gopurs:record-v1:0:*/
}) gopurs_runtime.Value {
	var x_0 struct {

		/*gopurs:record-v1:0:*/
	} = x_0_loop
	_ = x_0
	return func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict0()
	}()
}

func Call_Type_Data_Symbol_uncons(_dollar___unused_0_loop gopurs_runtime.Value, v_1_loop uint32) struct {
	head uint32
	tail uint32
	/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
} {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var v_1 uint32 = v_1_loop
	_ = v_1
	return struct {
		head uint32
		tail uint32
		/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
	}{513803634, 513803634}
}

func Call_Type_Data_Symbol_equalsSymbol(_dollar___unused_0_loop gopurs_runtime.Value, _dollar___unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var _dollar___unused_1 gopurs_runtime.Value = _dollar___unused_1_loop
	_ = _dollar___unused_1
	return Get_Type_Data_Symbol_Equals_dollar_Dict1()
}

func Call_Type_Data_Symbol_equals(_dollar___unused_0_loop gopurs_runtime.Value, v_1_loop uint32, v1_2_loop uint32) uint32 {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var v_1 uint32 = v_1_loop
	_ = v_1
	var v1_2 uint32 = v1_2_loop
	_ = v1_2
	return 513803634
}

func Call_Type_Data_Symbol_compare(_dollar___unused_0_loop gopurs_runtime.Value, v_1_loop uint32, v1_2_loop uint32) uint32 {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var v_1 uint32 = v_1_loop
	_ = v_1
	var v1_2 uint32 = v1_2_loop
	_ = v1_2
	return 513803634
}

func Call_Type_Data_Symbol_go__append(_dollar___unused_0_loop gopurs_runtime.Value, v_1_loop uint32, v1_2_loop uint32) uint32 {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var v_1 uint32 = v_1_loop
	_ = v_1
	var v1_2 uint32 = v1_2_loop
	_ = v1_2
	return 513803634
}
