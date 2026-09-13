package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_symB gopurs_runtime.Value
var once_Main_symB sync.Once

func Get_Main_symB() gopurs_runtime.Value {
	once_Main_symB.Do(func() {
		cache_Main_symB = gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}
	})
	return cache_Main_symB
}

var cache_Main_symA gopurs_runtime.Value
var once_Main_symA sync.Once

func Get_Main_symA() gopurs_runtime.Value {
	once_Main_symA.Do(func() {
		cache_Main_symA = gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}
	})
	return cache_Main_symA
}

var cache_Main_sym gopurs_runtime.Value
var once_Main_sym sync.Once

func Get_Main_sym() gopurs_runtime.Value {
	once_Main_sym.Do(func() {
		cache_Main_sym = gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}
	})
	return cache_Main_sym
}

var cache_Main_egBA gopurs_runtime.Value
var once_Main_egBA sync.Once

func Get_Main_egBA() gopurs_runtime.Value {
	once_Main_egBA.Do(func() {
		cache_Main_egBA = gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}
	})
	return cache_Main_egBA
}

var cache_Main_egAB gopurs_runtime.Value
var once_Main_egAB sync.Once

func Get_Main_egAB() gopurs_runtime.Value {
	once_Main_egAB.Do(func() {
		cache_Main_egAB = gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}
	})
	return cache_Main_egAB
}

var cache_Main_egA_prime_ gopurs_runtime.Value
var once_Main_egA_prime_ sync.Once

func Get_Main_egA_prime_() gopurs_runtime.Value {
	once_Main_egA_prime_.Do(func() {
		cache_Main_egA_prime_ = gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}
	})
	return cache_Main_egA_prime_
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			return Get_Data_Unit_unit()
		}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
				return Get_Data_Unit_unit()
			}), gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
					return Get_Data_Unit_unit()
				}), gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
				}))
			}))
		}))
	})
	return cache_Main_main
}
