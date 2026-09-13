package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_T gopurs_runtime.Value
var once_Main_T sync.Once

func Get_Main_T() gopurs_runtime.Value {
	once_Main_T.Do(func() {
		cache_Main_T = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_T(x_0_box)
		})
	})
	return cache_Main_T
}

var cache_Main_T__3985564026 gopurs_runtime.Value
var once_Main_T__3985564026 sync.Once

func Get_Main_T__3985564026() gopurs_runtime.Value {
	once_Main_T__3985564026.Do(func() {
		cache_Main_T__3985564026 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_T__3985564026(x_0_box)
		})
	})
	return cache_Main_T__3985564026
}

var cache_Main_runT gopurs_runtime.Value
var once_Main_runT sync.Once

func Get_Main_runT() gopurs_runtime.Value {
	once_Main_runT.Do(func() {
		cache_Main_runT = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_runT(v_0_box)
		})
	})
	return cache_Main_runT
}

var cache_Main_runT__3985564026 gopurs_runtime.Value
var once_Main_runT__3985564026 sync.Once

func Get_Main_runT__3985564026() gopurs_runtime.Value {
	once_Main_runT__3985564026.Do(func() {
		cache_Main_runT__3985564026 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_runT__3985564026(v_0_box)
		})
	})
	return cache_Main_runT__3985564026
}

var cache_Main_functorT gopurs_runtime.Value
var once_Main_functorT sync.Once

func Get_Main_functorT() gopurs_runtime.Value {
	once_Main_functorT.Do(func() {
		cache_Main_functorT = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, f_0, v_1)
		})}))}
	})
	return cache_Main_functorT
}

var cache_Main_applyT gopurs_runtime.Value
var once_Main_applyT sync.Once

func Get_Main_applyT() gopurs_runtime.Value {
	once_Main_applyT.Do(func() {
		cache_Main_applyT = gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Main_functorT()))}
		}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Effect_applyEffect()).V1, v_0, v1_1)
		})}))}
	})
	return cache_Main_applyT
}

var cache_Main_bindT gopurs_runtime.Value
var once_Main_bindT sync.Once

func Get_Main_bindT() gopurs_runtime.Value {
	once_Main_bindT.Do(func() {
		cache_Main_bindT = gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Main_applyT()))}
		}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
				x_2_0 := gopurs_runtime.Apply(v_0, gopurs_runtime.Value{})
				_ = x_2_0
				return gopurs_runtime.Apply(gopurs_runtime.Apply(f_1, x_2_0), gopurs_runtime.Value{})
			})
		})}))}
	})
	return cache_Main_bindT
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Main_bindT())), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Main_bindT())), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
			}))
		}))
	})
	return cache_Main_main
}

var cache_Main_applicativeT gopurs_runtime.Value
var once_Main_applicativeT sync.Once

func Get_Main_applicativeT() gopurs_runtime.Value {
	once_Main_applicativeT.Do(func() {
		cache_Main_applicativeT = gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Main_applyT()))}
		}), gopurs_runtime.Func(func(t_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
				return t_0
			})
		})}))}
	})
	return cache_Main_applicativeT
}

var cache_Main_monadT gopurs_runtime.Value
var once_Main_monadT sync.Once

func Get_Main_monadT() gopurs_runtime.Value {
	once_Main_monadT.Do(func() {
		cache_Main_monadT = gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Main_applicativeT()))}
		}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Main_bindT()))}
		})}))}
	})
	return cache_Main_monadT
}

func Call_Main_T(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_T__3985564026(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
T__3985564026:
	for {
		if false {
			continue T__3985564026
		}
		var x_0 gopurs_runtime.Value = x_0_loop
		_ = x_0
		return x_0
	}
}

func Call_Main_runT(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	return v_0
}

func Call_Main_runT__3985564026(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
runT__3985564026:
	for {
		if false {
			continue runT__3985564026
		}
		var v_0 gopurs_runtime.Value = v_0_loop
		_ = v_0
		return v_0
	}
}
