package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_replicateM_ gopurs_runtime.Value
var once_Main_replicateM_ sync.Once

func Get_Main_replicateM_() gopurs_runtime.Value {
	once_Main_replicateM_.Do(func() {
		cache_Main_replicateM_ = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_replicateM_(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](dictMonad_0_box))
		})
	})
	return cache_Main_replicateM_
}

var cache_Main_replicateM___1903616512 gopurs_runtime.Value
var once_Main_replicateM___1903616512 sync.Once

func Get_Main_replicateM___1903616512() gopurs_runtime.Value {
	once_Main_replicateM___1903616512.Do(func() {
		cache_Main_replicateM___1903616512 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_replicateM___1903616512(__eta_norm_1_0_box.FloatVal(), __eta_norm_0_1_box)
		})
	})
	return cache_Main_replicateM___1903616512
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Main_replicateM___1903616512(10.0, gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Hello World!"))), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
		}))
	})
	return cache_Main_main
}

func Call_Main_replicateM_(dictMonad_0_loop *Constructor_Control_Monad_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
replicateM_:
	for {
		if false {
			continue replicateM_
		}
		var dictMonad_0 *Constructor_Control_Monad_Monad[gopurs_runtime.Value] = dictMonad_0_loop
		_ = dictMonad_0
		// TAST (Let): Applicative0_1_0 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m$scope2)])
		Applicative0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonad_0.V0, gopurs_runtime.Value{}))
		_ = Applicative0_1_0
		// TAST (Let): Bind1_2_1 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m$scope2)])
		Bind1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonad_0.V1, gopurs_runtime.Value{}))
		_ = Bind1_2_1
		return gopurs_runtime.Func2(func(v_3 gopurs_runtime.Value, v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t2 gopurs_runtime.Value
			{
				if (v_3.FloatVal()) == (0.0) {
					__t2 = gopurs_runtime.Apply(Applicative0_1_0.V1, Get_Data_Unit_unit())
					goto end_branch_2
				} else {

				}
			}
			{
				__t2 = gopurs_runtime.Apply2(Bind1_2_1.V1, v1_4, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply2(Call_Main_replicateM_(dictMonad_0), gopurs_runtime.Float((v_3.FloatVal())-(1.0)), v1_4)
				}))
			}
		end_branch_2:
			return __t2
		})
	}
}

func Call_Main_replicateM___1903616512(__eta_norm_1_0_loop float64, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
replicateM___1903616512:
	for {
		if false {
			continue replicateM___1903616512
		}
		var __eta_norm_1_0 float64 = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		var __t1 gopurs_runtime.Value
		{
			if (__eta_norm_1_0) == (0.0) {
				__t1 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
					return Get_Data_Unit_unit()
				})
				goto end_branch_1
			} else {

			}
		}
		{
			__t1 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
				_dollar___unused_2_0 := gopurs_runtime.Apply(__eta_norm_0_1, gopurs_runtime.Value{})
				_ = _dollar___unused_2_0
				return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Main_replicateM_(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](Get_Effect_monadEffect())), gopurs_runtime.Float((__eta_norm_1_0)-(1.0)), __eta_norm_0_1), gopurs_runtime.Value{})
			})
		}
	end_branch_1:
		return __t1
	}
}
