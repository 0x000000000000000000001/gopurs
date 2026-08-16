package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_Stream gopurs_runtime.Value
var once_Main_Stream sync.Once

func Get_Main_Stream() gopurs_runtime.Value {
	once_Main_Stream.Do(func() {
		cache_Main_Stream = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 1020906690, UnsafePtr: unsafe.Pointer((&Constructor_Main_Stream{1, value0, value1}))}
			})
		})
	})
	return cache_Main_Stream
}

var cache_Main_IsStream_dollar_Dict gopurs_runtime.Value
var once_Main_IsStream_dollar_Dict sync.Once

func Get_Main_IsStream_dollar_Dict() gopurs_runtime.Value {
	once_Main_IsStream_dollar_Dict.Do(func() {
		cache_Main_IsStream_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_IsStream_dollar_Dict(x_0_box)
		})
	})
	return cache_Main_IsStream_dollar_Dict
}

var cache_Main_uncons gopurs_runtime.Value
var once_Main_uncons sync.Once

func Get_Main_uncons() gopurs_runtime.Value {
	once_Main_uncons.Do(func() {
		cache_Main_uncons = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_uncons(gopurs_runtime.CoerceToStruct[Constructor_Main_IsStream](dict_0_box))
		})
	})
	return cache_Main_uncons
}

var cache_Main_streamIsStream gopurs_runtime.Value
var once_Main_streamIsStream sync.Once

func Get_Main_streamIsStream() gopurs_runtime.Value {
	once_Main_streamIsStream.Do(func() {
		cache_Main_streamIsStream = gopurs_runtime.Value{Type: 9, IntVal: 668736856, UnsafePtr: unsafe.Pointer((&Constructor_Main_IsStream{1, gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 1020906690, UnsafePtr: unsafe.Pointer((&Constructor_Main_Stream{1, x_0, xs_1}))}
			})
		}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.RecordDict2("head", "tail", (*Constructor_Main_Stream)(v_0.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1020906690, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Main_Stream](gopurs_runtime.Apply((*Constructor_Main_Stream)(v_0.UnsafePtr).V1, Get_Data_Unit_unit())))})
		})}))}
	})
	return cache_Main_streamIsStream
}

var cache_Main_cons gopurs_runtime.Value
var once_Main_cons sync.Once

func Get_Main_cons() gopurs_runtime.Value {
	once_Main_cons.Do(func() {
		cache_Main_cons = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_cons(gopurs_runtime.CoerceToStruct[Constructor_Main_IsStream](dict_0_box))
		})
	})
	return cache_Main_cons
}

var cache_Main_test gopurs_runtime.Value
var once_Main_test sync.Once

func Get_Main_test() gopurs_runtime.Value {
	once_Main_test.Do(func() {
		cache_Main_test = gopurs_runtime.Func2(func(dictIsStream_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_test(gopurs_runtime.CoerceToStruct[Constructor_Main_IsStream](dictIsStream_0_box), s_1_box)
		})
	})
	return cache_Main_test
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = func() gopurs_runtime.Value {
			var dones_0_0_0 *Constructor_Main_Stream
			_ = dones_0_0_0
			dones_0_0_0 = gopurs_runtime.CoerceToStruct[Constructor_Main_Stream](gopurs_runtime.Value{Type: 9, IntVal: 1020906690, UnsafePtr: unsafe.Pointer((&Constructor_Main_Stream{1, gopurs_runtime.Str("Done"), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 1020906690, UnsafePtr: unsafe.Pointer(dones_0_0_0)}
			})}))})
			return gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str((dones_0_0_0).V0.StrVal()))
		}()
	})
	return cache_Main_main
}

var cache_Main_cons__3099608975 gopurs_runtime.Value
var once_Main_cons__3099608975 sync.Once

func Get_Main_cons__3099608975() gopurs_runtime.Value {
	once_Main_cons__3099608975.Do(func() {
		cache_Main_cons__3099608975 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_cons__3099608975(gopurs_runtime.CoerceToStruct[Constructor_Main_IsStream](dict_0_box))
		})
	})
	return cache_Main_cons__3099608975
}

var cache_Main_streamIsStream__1249510304 gopurs_runtime.Value
var once_Main_streamIsStream__1249510304 sync.Once

func Get_Main_streamIsStream__1249510304() gopurs_runtime.Value {
	once_Main_streamIsStream__1249510304.Do(func() {
		cache_Main_streamIsStream__1249510304 = gopurs_runtime.Value{Type: 9, IntVal: 668736856, UnsafePtr: unsafe.Pointer((&Constructor_Main_IsStream{1, gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 1020906690, UnsafePtr: unsafe.Pointer((&Constructor_Main_Stream{1, x_0, xs_1}))}
			})
		}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.RecordDict2("head", "tail", gopurs_runtime.Str((*Constructor_Main_Stream)(v_0.UnsafePtr).V0.StrVal()), gopurs_runtime.Value{Type: 9, IntVal: 1020906690, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Main_Stream](gopurs_runtime.Apply((*Constructor_Main_Stream)(v_0.UnsafePtr).V1, Get_Data_Unit_unit())))})
		})}))}
	})
	return cache_Main_streamIsStream__1249510304
}

var cache_Main_streamIsStream__3175905987 gopurs_runtime.Value
var once_Main_streamIsStream__3175905987 sync.Once

func Get_Main_streamIsStream__3175905987() gopurs_runtime.Value {
	once_Main_streamIsStream__3175905987.Do(func() {
		cache_Main_streamIsStream__3175905987 = gopurs_runtime.Value{Type: 9, IntVal: 668736856, UnsafePtr: unsafe.Pointer((&Constructor_Main_IsStream{1, gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 1020906690, UnsafePtr: unsafe.Pointer((&Constructor_Main_Stream{1, x_0, xs_1}))}
			})
		}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.RecordDict2("head", "tail", gopurs_runtime.Str((*Constructor_Main_Stream)(v_0.UnsafePtr).V0.StrVal()), gopurs_runtime.Value{Type: 9, IntVal: 1020906690, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Main_Stream](gopurs_runtime.Apply((*Constructor_Main_Stream)(v_0.UnsafePtr).V1, Get_Data_Unit_unit())))})
		})}))}
	})
	return cache_Main_streamIsStream__3175905987
}

var cache_Main_test__2939414718 gopurs_runtime.Value
var once_Main_test__2939414718 sync.Once

func Get_Main_test__2939414718() gopurs_runtime.Value {
	once_Main_test__2939414718.Do(func() {
		cache_Main_test__2939414718 = gopurs_runtime.Func2(func(dictIsStream_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_test__2939414718(gopurs_runtime.CoerceToStruct[Constructor_Main_IsStream](dictIsStream_0_box), s_1_box)
		})
	})
	return cache_Main_test__2939414718
}

var cache_Main_uncons__3027971297 gopurs_runtime.Value
var once_Main_uncons__3027971297 sync.Once

func Get_Main_uncons__3027971297() gopurs_runtime.Value {
	once_Main_uncons__3027971297.Do(func() {
		cache_Main_uncons__3027971297 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_uncons__3027971297(gopurs_runtime.CoerceToStruct[Constructor_Main_IsStream](dict_0_box))
		})
	})
	return cache_Main_uncons__3027971297
}

type Constructor_Main_Stream struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}

type Constructor_Main_IsStream struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}

func init() {
	gopurs_runtime.StructGetters[668736856] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_IsStream)(ptr)
		_ = c
		switch key {
		case "cons":
			return gopurs_runtime.Box(c.V0)
		case "uncons":
			return gopurs_runtime.Box(c.V1)
		default:
			panic("Key not found in dictionary Constructor_Main_IsStream: " + key)
		}
	}
}

func Call_Main_IsStream_dollar_Dict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_uncons(dict_0_loop *Constructor_Main_IsStream) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_IsStream = dict_0_loop
	_ = dict_0
	return gopurs_runtime.Box(dict_0.V1)
}

func Call_Main_cons(dict_0_loop *Constructor_Main_IsStream) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_IsStream = dict_0_loop
	_ = dict_0
	return gopurs_runtime.Box(dict_0.V0)
}

func Call_Main_test(dictIsStream_0_loop *Constructor_Main_IsStream, s_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictIsStream_0 *Constructor_Main_IsStream = dictIsStream_0_loop
	_ = dictIsStream_0
	var s_1 gopurs_runtime.Value = s_1_loop
	_ = s_1
	// TAST (Let): v_2_0 shape=App(Other) expectedFromAst=gopurs_runtime.Value actual=gopurs_runtime.Value bindingType=Any
	v_2_0 := gopurs_runtime.Apply(gopurs_runtime.Box(dictIsStream_0.V1), s_1)
	_ = v_2_0
	// TAST (Let): __local_var_3_1 shape=Other expectedFromAst=gopurs_runtime.Value actual=gopurs_runtime.Value bindingType=Any
	__local_var_3_1 := gopurs_runtime.RecordGet(v_2_0, "tail")
	_ = __local_var_3_1
	return gopurs_runtime.Apply2(gopurs_runtime.Box(dictIsStream_0.V0), gopurs_runtime.RecordGet(v_2_0, "head"), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return __local_var_3_1
	}))
}

func Call_Main_cons__3099608975(dict_0_loop *Constructor_Main_IsStream) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_IsStream = dict_0_loop
	_ = dict_0
	return gopurs_runtime.Box(dict_0.V0)
}

func Call_Main_test__2939414718(dictIsStream_0_loop *Constructor_Main_IsStream, s_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictIsStream_0 *Constructor_Main_IsStream = dictIsStream_0_loop
	_ = dictIsStream_0
	var s_1 gopurs_runtime.Value = s_1_loop
	_ = s_1
	// TAST (Let): v_2_0 shape=App(Other) expectedFromAst=gopurs_runtime.Value actual=gopurs_runtime.Value bindingType=Any
	v_2_0 := gopurs_runtime.Apply(gopurs_runtime.Box(dictIsStream_0.V1), s_1)
	_ = v_2_0
	// TAST (Let): __local_var_3_1 shape=Other expectedFromAst=gopurs_runtime.Value actual=gopurs_runtime.Value bindingType=Any
	__local_var_3_1 := gopurs_runtime.RecordGet(v_2_0, "tail")
	_ = __local_var_3_1
	return gopurs_runtime.Apply2(gopurs_runtime.Box(dictIsStream_0.V0), gopurs_runtime.RecordGet(v_2_0, "head"), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return __local_var_3_1
	}))
}

func Call_Main_uncons__3027971297(dict_0_loop *Constructor_Main_IsStream) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_IsStream = dict_0_loop
	_ = dict_0
	return gopurs_runtime.Box(dict_0.V1)
}
