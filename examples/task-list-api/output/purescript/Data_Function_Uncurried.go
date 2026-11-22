package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Data_Function_Uncurried_runFn1 gopurs_runtime.Value
var once_Data_Function_Uncurried_runFn1 sync.Once

func Get_Data_Function_Uncurried_runFn1() gopurs_runtime.Value {
	once_Data_Function_Uncurried_runFn1.Do(func() {
		cache_Data_Function_Uncurried_runFn1 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Function_Uncurried_runFn1(f_0_box)
		})
	})
	return cache_Data_Function_Uncurried_runFn1
}

var cache_Data_Function_Uncurried_runFn1__3939456157 gopurs_runtime.Value
var once_Data_Function_Uncurried_runFn1__3939456157 sync.Once

func Get_Data_Function_Uncurried_runFn1__3939456157() gopurs_runtime.Value {
	once_Data_Function_Uncurried_runFn1__3939456157.Do(func() {
		cache_Data_Function_Uncurried_runFn1__3939456157 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Data_Function_Uncurried_runFn1__3939456157(f_unused_0_box, __eta_norm_0_1_box))
		})
	})
	return cache_Data_Function_Uncurried_runFn1__3939456157
}

var cache_Data_Function_Uncurried_mkFn1 gopurs_runtime.Value
var once_Data_Function_Uncurried_mkFn1 sync.Once

func Get_Data_Function_Uncurried_mkFn1() gopurs_runtime.Value {
	once_Data_Function_Uncurried_mkFn1.Do(func() {
		cache_Data_Function_Uncurried_mkFn1 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Function_Uncurried_mkFn1(f_0_box)
		})
	})
	return cache_Data_Function_Uncurried_mkFn1
}

func Call_Data_Function_Uncurried_runFn1(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	return f_0
}

func Call_Data_Function_Uncurried_runFn1__3939456157(f_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) string {
runFn1__3939456157:
	for {
		if false {
			continue runFn1__3939456157
		}
		var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
		_ = f_unused_0
		var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		return gopurs_runtime.Apply(Get_Node_Symbol_showSymbolImpl(), __eta_norm_0_1).StrVal()
	}
}

func Call_Data_Function_Uncurried_mkFn1(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	return f_0
}

func Get_Data_Function_Uncurried_mkFn0() gopurs_runtime.Value {
	return _Gopurs_Data_Function_Uncurried_MkFn0
}

func Get_Data_Function_Uncurried_mkFn10() gopurs_runtime.Value {
	return _Gopurs_Data_Function_Uncurried_MkFn10
}

func Get_Data_Function_Uncurried_mkFn2() gopurs_runtime.Value {
	return _Gopurs_Data_Function_Uncurried_MkFn2
}

func Get_Data_Function_Uncurried_mkFn3() gopurs_runtime.Value {
	return _Gopurs_Data_Function_Uncurried_MkFn3
}

func Get_Data_Function_Uncurried_mkFn4() gopurs_runtime.Value {
	return _Gopurs_Data_Function_Uncurried_MkFn4
}

func Get_Data_Function_Uncurried_mkFn5() gopurs_runtime.Value {
	return _Gopurs_Data_Function_Uncurried_MkFn5
}

func Get_Data_Function_Uncurried_mkFn6() gopurs_runtime.Value {
	return _Gopurs_Data_Function_Uncurried_MkFn6
}

func Get_Data_Function_Uncurried_mkFn7() gopurs_runtime.Value {
	return _Gopurs_Data_Function_Uncurried_MkFn7
}

func Get_Data_Function_Uncurried_mkFn8() gopurs_runtime.Value {
	return _Gopurs_Data_Function_Uncurried_MkFn8
}

func Get_Data_Function_Uncurried_mkFn9() gopurs_runtime.Value {
	return _Gopurs_Data_Function_Uncurried_MkFn9
}

func Get_Data_Function_Uncurried_runFn0() gopurs_runtime.Value {
	return _Gopurs_Data_Function_Uncurried_RunFn0
}

func Get_Data_Function_Uncurried_runFn10() gopurs_runtime.Value {
	return _Gopurs_Data_Function_Uncurried_RunFn10
}

func Get_Data_Function_Uncurried_runFn2() gopurs_runtime.Value {
	return _Gopurs_Data_Function_Uncurried_RunFn2
}

func Get_Data_Function_Uncurried_runFn3() gopurs_runtime.Value {
	return _Gopurs_Data_Function_Uncurried_RunFn3
}

func Get_Data_Function_Uncurried_runFn4() gopurs_runtime.Value {
	return _Gopurs_Data_Function_Uncurried_RunFn4
}

func Get_Data_Function_Uncurried_runFn5() gopurs_runtime.Value {
	return _Gopurs_Data_Function_Uncurried_RunFn5
}

func Get_Data_Function_Uncurried_runFn6() gopurs_runtime.Value {
	return _Gopurs_Data_Function_Uncurried_RunFn6
}

func Get_Data_Function_Uncurried_runFn7() gopurs_runtime.Value {
	return _Gopurs_Data_Function_Uncurried_RunFn7
}

func Get_Data_Function_Uncurried_runFn8() gopurs_runtime.Value {
	return _Gopurs_Data_Function_Uncurried_RunFn8
}

func Get_Data_Function_Uncurried_runFn9() gopurs_runtime.Value {
	return _Gopurs_Data_Function_Uncurried_RunFn9
}
