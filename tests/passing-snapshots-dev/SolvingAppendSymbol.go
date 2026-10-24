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
		cache_Main_egBA = gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Type_Data_Symbol_go__append(gopurs_runtime.Value{}, 513803634, 513803634)), UnsafePtr: nil}
	})
	return cache_Main_egBA
}

var cache_Main_egAB gopurs_runtime.Value
var once_Main_egAB sync.Once

func Get_Main_egAB() gopurs_runtime.Value {
	once_Main_egAB.Do(func() {
		cache_Main_egAB = gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Type_Data_Symbol_go__append(gopurs_runtime.Value{}, 513803634, 513803634)), UnsafePtr: nil}
	})
	return cache_Main_egAB
}

var cache_Main_egA_prime_ gopurs_runtime.Value
var once_Main_egA_prime_ sync.Once

func Get_Main_egA_prime_() gopurs_runtime.Value {
	once_Main_egA_prime_.Do(func() {
		cache_Main_egA_prime_ = gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Type_Data_Symbol_go__append(gopurs_runtime.Value{}, 513803634, Call_Type_Data_Symbol_go__append(gopurs_runtime.Value{}, 513803634, 513803634))), UnsafePtr: nil}
	})
	return cache_Main_egA_prime_
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): gotBA_0_0 shape=Other bindingType=Boolean
			gotBA_0_0 := (gopurs_runtime.Apply(Call_Data_Symbol_reflectSymbol((&Constructor_Data_Symbol_IsSymbol[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Str("BA")
			})})), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(Get_Main_egBA().IntVal)), UnsafePtr: nil}).StrVal()) == ("BA")
			_ = gotBA_0_0
			// TAST (Let): gotAB_1_1 shape=Other bindingType=Boolean
			gotAB_1_1 := (gopurs_runtime.Apply(Call_Data_Symbol_reflectSymbol((&Constructor_Data_Symbol_IsSymbol[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Str("AB")
			})})), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(Get_Main_egAB().IntVal)), UnsafePtr: nil}).StrVal()) == ("AB")
			_ = gotAB_1_1
			// TAST (Let): gotA_prime__2_2 shape=Other bindingType=Boolean
			gotA_prime__2_2 := (gopurs_runtime.Apply(Call_Data_Symbol_reflectSymbol((&Constructor_Data_Symbol_IsSymbol[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Str("A")
			})})), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(Get_Main_egA_prime_().IntVal)), UnsafePtr: nil}).StrVal()) == ("A")
			_ = gotA_prime__2_2
			// TAST (Let): __local_var_3_4 shape=App(Var) bindingType=(ADT ["Effect","Effect"] [Unit])
			__local_var_3_4 := gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Did not get AB"))
			_ = __local_var_3_4
			var __t5 gopurs_runtime.Value
			{
				if (gotAB_1_1) != (true) {
					__t5 = __local_var_3_4
					goto end_branch_5
				} else {

				}
			}
			{
				__t5 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
					return Get_Data_Unit_unit()
				})
			}
		end_branch_5:
			// TAST (Let): __local_var_3_3 shape=Let(Branch(Other, def=EffectPure)) bindingType=(ADT ["Effect","Effect"] [Unit])
			__local_var_3_3 := __t5
			_ = __local_var_3_3
			_dollar___unused_4_6 := gopurs_runtime.Apply(__local_var_3_3, gopurs_runtime.Value{})
			_ = _dollar___unused_4_6
			// TAST (Let): __local_var_5_8 shape=App(Var) bindingType=(ADT ["Effect","Effect"] [Unit])
			__local_var_5_8 := gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Did not get BA"))
			_ = __local_var_5_8
			var __t9 gopurs_runtime.Value
			{
				if (gotBA_0_0) != (true) {
					__t9 = __local_var_5_8
					goto end_branch_9
				} else {

				}
			}
			{
				__t9 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
					return Get_Data_Unit_unit()
				})
			}
		end_branch_9:
			_dollar___unused_5_7 := gopurs_runtime.Apply(__t9, gopurs_runtime.Value{})
			_ = _dollar___unused_5_7
			// TAST (Let): __local_var_6_11 shape=App(Var) bindingType=(ADT ["Effect","Effect"] [Unit])
			__local_var_6_11 := gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Did not get A"))
			_ = __local_var_6_11
			var __t12 gopurs_runtime.Value
			{
				if (gotA_prime__2_2) != (true) {
					__t12 = __local_var_6_11
					goto end_branch_12
				} else {

				}
			}
			{
				__t12 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
					return Get_Data_Unit_unit()
				})
			}
		end_branch_12:
			_dollar___unused_6_10 := gopurs_runtime.Apply(__t12, gopurs_runtime.Value{})
			_ = _dollar___unused_6_10
			// TAST (Let): __local_var_7_13 shape=App(Var) bindingType=(ADT ["Effect","Effect"] [Unit])
			__local_var_7_13 := gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
			_ = __local_var_7_13
			var __t14 gopurs_runtime.Value
			{
				if (gotAB_1_1) && ((gotBA_0_0) && (gotA_prime__2_2)) {
					__t14 = __local_var_7_13
					goto end_branch_14
				} else {

				}
			}
			{
				__t14 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
					return Get_Data_Unit_unit()
				})
			}
		end_branch_14:
			return gopurs_runtime.Apply(__t14, gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}
