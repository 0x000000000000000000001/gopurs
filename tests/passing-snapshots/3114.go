package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_barIsSymbol gopurs_runtime.Value
var once_Main_barIsSymbol sync.Once

func Get_Main_barIsSymbol() gopurs_runtime.Value {
	once_Main_barIsSymbol.Do(func() {
		cache_Main_barIsSymbol = gopurs_runtime.Value{Type: 9, IntVal: 2134024384, UnsafePtr: unsafe.Pointer(&Constructor_Data_Symbol_IsSymbol{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("bar")
		})})}
	})
	return cache_Main_barIsSymbol
}

var cache_Main_showTuple gopurs_runtime.Value
var once_Main_showTuple sync.Once

func Get_Main_showTuple() gopurs_runtime.Value {
	once_Main_showTuple.Do(func() {
		cache_Main_showTuple = gopurs_runtime.Func(func(dictShow1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_showTuple(dictShow1_0_box)
		})
	})
	return cache_Main_showTuple
}

var cache_Main_showTuple1 gopurs_runtime.Value
var once_Main_showTuple1 sync.Once

func Get_Main_showTuple1() gopurs_runtime.Value {
	once_Main_showTuple1.Do(func() {
		cache_Main_showTuple1 = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(&Constructor_Data_Show_Show{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str((((("(Tuple ") + (gopurs_runtime.Apply(Get_Data_Show_showStringImpl(), (*Constructor_Data_Tuple_Tuple)(v_0.UnsafePtr).V0).StrVal())) + (" ")) + (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), (*Constructor_Data_Tuple_Tuple)(v_0.UnsafePtr).V1).StrVal())) + (")"))
		})})}
	})
	return cache_Main_showTuple1
}

var cache_Main_fooIsSymbol gopurs_runtime.Value
var once_Main_fooIsSymbol sync.Once

func Get_Main_fooIsSymbol() gopurs_runtime.Value {
	once_Main_fooIsSymbol.Do(func() {
		cache_Main_fooIsSymbol = gopurs_runtime.Value{Type: 9, IntVal: 2134024384, UnsafePtr: unsafe.Pointer(&Constructor_Data_Symbol_IsSymbol{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("foo")
		})})}
	})
	return cache_Main_fooIsSymbol
}

var cache_Main_showMaybe gopurs_runtime.Value
var once_Main_showMaybe sync.Once

func Get_Main_showMaybe() gopurs_runtime.Value {
	once_Main_showMaybe.Do(func() {
		cache_Main_showMaybe = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(&Constructor_Data_Show_Show{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t0 string
			{
				if v_0.Type == 9 && v_0.IntVal == 930809136 && v_0.UnsafePtr != nil {
					__t0 = (("(Just ") + (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), (*Constructor_Data_Maybe_Just)(v_0.UnsafePtr).V0).StrVal())) + (")")
					goto end_branch_0
				} else {

				}
			}
			{
				if v_0.Type == 9 && v_0.IntVal == 930809136 && v_0.UnsafePtr == nil {
					__t0 = "Nothing"
					goto end_branch_0
				} else {

				}
			}
			{
				__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }().StrVal()
			}
		end_branch_0:
			return gopurs_runtime.Str(__t0)
		})})}
	})
	return cache_Main_showMaybe
}

var cache_Main_showTuple2 gopurs_runtime.Value
var once_Main_showTuple2 sync.Once

func Get_Main_showTuple2() gopurs_runtime.Value {
	once_Main_showTuple2.Do(func() {
		cache_Main_showTuple2 = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(&Constructor_Data_Show_Show{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str((((("(Tuple ") + (gopurs_runtime.Apply(Get_Data_Show_showStringImpl(), (*Constructor_Data_Tuple_Tuple)(v_0.UnsafePtr).V0).StrVal())) + (" ")) + (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), (*Constructor_Data_Tuple_Tuple)(v_0.UnsafePtr).V1).StrVal())) + (")"))
		})})}
	})
	return cache_Main_showTuple2
}

var cache_Main_showMaybe1 gopurs_runtime.Value
var once_Main_showMaybe1 sync.Once

func Get_Main_showMaybe1() gopurs_runtime.Value {
	once_Main_showMaybe1.Do(func() {
		cache_Main_showMaybe1 = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(&Constructor_Data_Show_Show{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t0 string
			{
				if v_0.Type == 9 && v_0.IntVal == 930809136 && v_0.UnsafePtr != nil {
					__t0 = (("(Just ") + (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), (*Constructor_Data_Maybe_Just)(v_0.UnsafePtr).V0).StrVal())) + (")")
					goto end_branch_0
				} else {

				}
			}
			{
				if v_0.Type == 9 && v_0.IntVal == 930809136 && v_0.UnsafePtr == nil {
					__t0 = "Nothing"
					goto end_branch_0
				} else {

				}
			}
			{
				__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }().StrVal()
			}
		end_branch_0:
			return gopurs_runtime.Str(__t0)
		})})}
	})
	return cache_Main_showMaybe1
}

var cache_Main__foo gopurs_runtime.Value
var once_Main__foo sync.Once

func Get_Main__foo() gopurs_runtime.Value {
	once_Main__foo.Do(func() {
		cache_Main__foo = gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}
	})
	return cache_Main__foo
}

var cache_Main__bar gopurs_runtime.Value
var once_Main__bar sync.Once

func Get_Main__bar() gopurs_runtime.Value {
	once_Main__bar.Do(func() {
		cache_Main__bar = gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}
	})
	return cache_Main__bar
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

func Call_Main_showTuple(dictShow1_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictShow1_0 gopurs_runtime.Value = dictShow1_0_loop
	_ = dictShow1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(&Constructor_Data_Show_Show{1, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Str((((("(Tuple ") + (gopurs_runtime.Apply(Get_Data_Show_showStringImpl(), (*Constructor_Data_Tuple_Tuple)(v_1.UnsafePtr).V0).StrVal())) + (" ")) + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow1_0, "show"), (*Constructor_Data_Tuple_Tuple)(v_1.UnsafePtr).V1).StrVal())) + (")"))
	})})}
}
