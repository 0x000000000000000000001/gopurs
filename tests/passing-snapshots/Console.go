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
			return Call_Main_replicateM_(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](dictMonad_0_box))
		})
	})
	return cache_Main_replicateM_
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=App(Var) expectedFromAst=gopurs_runtime.Value actual=gopurs_runtime.Value bindingType=(ADT ["Effect","Effect"] [Unit])
			__local_var_0_0 := gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Hello World!"))
			_ = __local_var_0_0
			_dollar___unused_1_2 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = _dollar___unused_1_2
			_dollar___unused_1_1 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Main_replicateM_(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](Get_Effect_monadEffect())), gopurs_runtime.Float(9.0), __local_var_0_0), gopurs_runtime.Value{})
			_ = _dollar___unused_1_1
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

var cache_Main_replicateM___3797747736 gopurs_runtime.Value
var once_Main_replicateM___3797747736 sync.Once

func Get_Main_replicateM___3797747736() gopurs_runtime.Value {
	once_Main_replicateM___3797747736.Do(func() {
		cache_Main_replicateM___3797747736 = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_replicateM___3797747736(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](dictMonad_0_box))
		})
	})
	return cache_Main_replicateM___3797747736
}

func Call_Main_replicateM_(dictMonad_0_loop *Constructor_Control_Monad_Monad) gopurs_runtime.Value {
replicateM_:
	for {
		if false {
			continue replicateM_
		}
		var dictMonad_0 *Constructor_Control_Monad_Monad = dictMonad_0_loop
		_ = dictMonad_0
		// TAST (Let): Applicative0_1_0 shape=App(Other) expectedFromAst=*Constructor_Control_Applicative_Applicative actual=*Constructor_Control_Applicative_Applicative bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m)])
		Applicative0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V0), gopurs_runtime.Value{}))
		_ = Applicative0_1_0
		// TAST (Let): Bind1_2_1 shape=App(Other) expectedFromAst=*Constructor_Control_Bind_Bind actual=*Constructor_Control_Bind_Bind bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
		Bind1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V1), gopurs_runtime.Value{}))
		_ = Bind1_2_1
		return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t2 gopurs_runtime.Value
				{
					if (v_3.FloatVal()) == (0.0) {
						__t2 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_1_0.V1), Get_Data_Unit_unit())
						goto end_branch_2
					} else {

					}
				}
				{
					__t2 = gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_1.V1), v1_4, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Apply2(Call_Main_replicateM_(dictMonad_0), gopurs_runtime.Float((v_3.FloatVal())-(1.0)), v1_4)
					}))
				}
			end_branch_2:
				return __t2
			})
		})
	}
}

func Call_Main_replicateM___3797747736(dictMonad_0_loop *Constructor_Control_Monad_Monad) gopurs_runtime.Value {
	var dictMonad_0 *Constructor_Control_Monad_Monad = dictMonad_0_loop
	_ = dictMonad_0
	// TAST (Let): Applicative0_1_0 shape=App(Other) expectedFromAst=*Constructor_Control_Applicative_Applicative actual=*Constructor_Control_Applicative_Applicative bindingType=(ADT ["Control","Applicative","Applicative"] [Any])
	Applicative0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V0), gopurs_runtime.Value{}))
	_ = Applicative0_1_0
	// TAST (Let): Bind1_2_1 shape=App(Other) expectedFromAst=*Constructor_Control_Bind_Bind actual=*Constructor_Control_Bind_Bind bindingType=(ADT ["Control","Bind","Bind"] [Any])
	Bind1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V1), gopurs_runtime.Value{}))
	_ = Bind1_2_1
	return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t6 gopurs_runtime.Value
			{
				if (v_3.FloatVal()) == (0.0) {
					__t6 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_1_0.V1), Get_Data_Unit_unit())
					goto end_branch_6
				} else {

				}
			}
			{
				__t6 = gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_1.V1), v1_4, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
					// TAST (Let): Applicative0_6_2 shape=App(Other) expectedFromAst=*Constructor_Control_Applicative_Applicative actual=*Constructor_Control_Applicative_Applicative bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m)])
					Applicative0_6_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V0), gopurs_runtime.Value{}))
					_ = Applicative0_6_2
					// TAST (Let): Bind1_7_3 shape=App(Other) expectedFromAst=*Constructor_Control_Bind_Bind actual=*Constructor_Control_Bind_Bind bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
					Bind1_7_3 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V1), gopurs_runtime.Value{}))
					_ = Bind1_7_3
					// TAST (Let): __local_var_8_4 shape=Other expectedFromAst=float64 actual=float64 bindingType=Number
					__local_var_8_4 := (v_3.FloatVal()) - (1.0)
					_ = __local_var_8_4
					var __t5 gopurs_runtime.Value
					{
						if (__local_var_8_4) == (0.0) {
							__t5 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_6_2.V1), Get_Data_Unit_unit())
							goto end_branch_5
						} else {

						}
					}
					{
						__t5 = gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_3.V1), v1_4, gopurs_runtime.Func(func(_dollar___unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Apply2(Call_Main_replicateM_(dictMonad_0), gopurs_runtime.Float((__local_var_8_4)-(1.0)), v1_4)
						}))
					}
				end_branch_5:
					return __t5
				}))
			}
		end_branch_6:
			return __t6
		})
	})
}
