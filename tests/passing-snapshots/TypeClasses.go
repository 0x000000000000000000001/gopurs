package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_show gopurs_runtime.Value
var once_Main_show sync.Once

func Get_Main_show() gopurs_runtime.Value {
	once_Main_show.Do(func() {
		cache_Main_show = Get_Data_Show_showStringImpl()
	})
	return cache_Main_show
}

var cache_Main_pure gopurs_runtime.Value
var once_Main_pure sync.Once

func Get_Main_pure() gopurs_runtime.Value {
	once_Main_pure.Do(func() {
		cache_Main_pure = gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeFn()).V1)
	})
	return cache_Main_pure
}

var cache_Main_Nothing gopurs_runtime.Value
var once_Main_Nothing sync.Once

func Get_Main_Nothing() gopurs_runtime.Value {
	once_Main_Nothing.Do(func() {
		cache_Main_Nothing = gopurs_runtime.Value{Type: 9, IntVal: 3271839782, UnsafePtr: unsafe.Pointer((*Constructor_Main_Just)(nil))}
	})
	return cache_Main_Nothing
}

var cache_Main_Just gopurs_runtime.Value
var once_Main_Just sync.Once

func Get_Main_Just() gopurs_runtime.Value {
	once_Main_Just.Do(func() {
		cache_Main_Just = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3271839782, UnsafePtr: unsafe.Pointer((&Constructor_Main_Just{1, value0}))}
		})
	})
	return cache_Main_Just
}

var cache_Main_Data gopurs_runtime.Value
var once_Main_Data sync.Once

func Get_Main_Data() gopurs_runtime.Value {
	once_Main_Data.Do(func() {
		cache_Main_Data = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return value0
		})
	})
	return cache_Main_Data
}

var cache_Main_test8 gopurs_runtime.Value
var once_Main_test8 sync.Once

func Get_Main_test8() gopurs_runtime.Value {
	once_Main_test8.Do(func() {
		cache_Main_test8 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_test8(v_0_box))
		})
	})
	return cache_Main_test8
}

var cache_Main_test7 gopurs_runtime.Value
var once_Main_test7 sync.Once

func Get_Main_test7() gopurs_runtime.Value {
	once_Main_test7.Do(func() {
		cache_Main_test7 = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_test7(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show](dictShow_0_box))
		})
	})
	return cache_Main_test7
}

var cache_Main_test4 gopurs_runtime.Value
var once_Main_test4 sync.Once

func Get_Main_test4() gopurs_runtime.Value {
	once_Main_test4.Do(func() {
		cache_Main_test4 = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_test4(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](dictMonad_0_box))
		})
	})
	return cache_Main_test4
}

var cache_Main_test1 gopurs_runtime.Value
var once_Main_test1 sync.Once

func Get_Main_test1() gopurs_runtime.Value {
	once_Main_test1.Do(func() {
		cache_Main_test1 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_test1(v_0_box))
		})
	})
	return cache_Main_test1
}

var cache_Main_showData gopurs_runtime.Value
var once_Main_showData sync.Once

func Get_Main_showData() gopurs_runtime.Value {
	once_Main_showData.Do(func() {
		cache_Main_showData = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_showData(dictShow_0_box)
		})
	})
	return cache_Main_showData
}

var cache_Main_showData1 gopurs_runtime.Value
var once_Main_showData1 sync.Once

func Get_Main_showData1() gopurs_runtime.Value {
	once_Main_showData1.Do(func() {
		cache_Main_showData1 = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer((&Constructor_Data_Show_Show{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str((("Data (") + (gopurs_runtime.Apply(Get_Data_Show_showStringImpl(), v_0).StrVal())) + (")"))
		})}))}
	})
	return cache_Main_showData1
}

var cache_Main_test3 gopurs_runtime.Value
var once_Main_test3 sync.Once

func Get_Main_test3() gopurs_runtime.Value {
	once_Main_test3.Do(func() {
		cache_Main_test3 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_test3(v_0_box))
		})
	})
	return cache_Main_test3
}

var cache_Main_runReader gopurs_runtime.Value
var once_Main_runReader sync.Once

func Get_Main_runReader() gopurs_runtime.Value {
	once_Main_runReader.Do(func() {
		cache_Main_runReader = gopurs_runtime.Func2(func(r_0_box gopurs_runtime.Value, f1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_runReader(r_0_box, f1_1_box)
		})
	})
	return cache_Main_runReader
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=App(Var) expectedFromAst=gopurs_runtime.Value actual=gopurs_runtime.Value bindingType=Any
			__local_var_0_0 := gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showStringImpl(), gopurs_runtime.Str("Hello")).StrVal()))
			_ = __local_var_0_0
			_dollar___unused_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = _dollar___unused_1_1
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

var cache_Main_f gopurs_runtime.Value
var once_Main_f sync.Once

func Get_Main_f() gopurs_runtime.Value {
	once_Main_f.Do(func() {
		cache_Main_f = gopurs_runtime.Func2(func(dictShow_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_f(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show](dictShow_0_box), x_1_box))
		})
	})
	return cache_Main_f
}

var cache_Main_test2 gopurs_runtime.Value
var once_Main_test2 sync.Once

func Get_Main_test2() gopurs_runtime.Value {
	once_Main_test2.Do(func() {
		cache_Main_test2 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_test2(v_0_box))
		})
	})
	return cache_Main_test2
}

var cache_Main_ask gopurs_runtime.Value
var once_Main_ask sync.Once

func Get_Main_ask() gopurs_runtime.Value {
	once_Main_ask.Do(func() {
		cache_Main_ask = gopurs_runtime.Func(func(r_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_ask(r_0_box)
		})
	})
	return cache_Main_ask
}

var cache_Main_test9 gopurs_runtime.Value
var once_Main_test9 sync.Once

func Get_Main_test9() gopurs_runtime.Value {
	once_Main_test9.Do(func() {
		cache_Main_test9 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Float(Call_Main_test9(v_0_box))
		})
	})
	return cache_Main_test9
}

var cache_Main_monadMaybe gopurs_runtime.Value
var once_Main_monadMaybe sync.Once

func Get_Main_monadMaybe() gopurs_runtime.Value {
	once_Main_monadMaybe.Do(func() {
		cache_Main_monadMaybe = gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Monad{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Main_applicativeMaybe()))}
		}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](Get_Main_bindMaybe()))}
		})}))}
	})
	return cache_Main_monadMaybe
}

var cache_Main_functorMaybe gopurs_runtime.Value
var once_Main_functorMaybe sync.Once

func Get_Main_functorMaybe() gopurs_runtime.Value {
	once_Main_functorMaybe.Do(func() {
		cache_Main_functorMaybe = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor{1, gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](Get_Main_bindMaybe()).V1), a_1, gopurs_runtime.Func(func(a_prime__2 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Main_applicativeMaybe()).V1), gopurs_runtime.Apply(f_0, a_prime__2))
				}))
			})
		})}))}
	})
	return cache_Main_functorMaybe
}

var cache_Main_bindMaybe gopurs_runtime.Value
var once_Main_bindMaybe sync.Once

func Get_Main_bindMaybe() gopurs_runtime.Value {
	once_Main_bindMaybe.Do(func() {
		cache_Main_bindMaybe = gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Get_Main_applyMaybe()))}
		}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t0 *Constructor_Main_Just
				{
					if v_0.Type == 9 && v_0.IntVal == 3271839782 && v_0.UnsafePtr == nil {
						__t0 = (*Constructor_Main_Just)(nil)
						goto end_branch_0
					} else {

					}
				}
				{
					if v_0.Type == 9 && v_0.IntVal == 3271839782 && v_0.UnsafePtr != nil {
						__t0 = gopurs_runtime.CoerceToStruct[Constructor_Main_Just](gopurs_runtime.Apply(v1_1, (*Constructor_Main_Just)(v_0.UnsafePtr).V0))
						goto end_branch_0
					} else {

					}
				}
				{
					__t0 = gopurs_runtime.CoerceToStruct[Constructor_Main_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
				}
			end_branch_0:
				return gopurs_runtime.Value{Type: 9, IntVal: 3271839782, UnsafePtr: unsafe.Pointer(__t0)}
			})
		})}))}
	})
	return cache_Main_bindMaybe
}

var cache_Main_applyMaybe gopurs_runtime.Value
var once_Main_applyMaybe sync.Once

func Get_Main_applyMaybe() gopurs_runtime.Value {
	once_Main_applyMaybe.Do(func() {
		cache_Main_applyMaybe = gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Main_functorMaybe()))}
		}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t1 *Constructor_Main_Just
				{
					if f_0.Type == 9 && f_0.IntVal == 3271839782 && f_0.UnsafePtr == nil {
						__t1 = (*Constructor_Main_Just)(nil)
						goto end_branch_1
					} else {

					}
				}
				{
					if f_0.Type == 9 && f_0.IntVal == 3271839782 && f_0.UnsafePtr != nil {
						var __t0 *Constructor_Main_Just
						{
							if a_1.Type == 9 && a_1.IntVal == 3271839782 && a_1.UnsafePtr == nil {
								__t0 = (*Constructor_Main_Just)(nil)
								goto end_branch_0
							} else {

							}
						}
						{
							if a_1.Type == 9 && a_1.IntVal == 3271839782 && a_1.UnsafePtr != nil {
								__t0 = gopurs_runtime.CoerceToStruct[Constructor_Main_Just](gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Main_applicativeMaybe()).V1), gopurs_runtime.Apply((*Constructor_Main_Just)(f_0.UnsafePtr).V0, (*Constructor_Main_Just)(a_1.UnsafePtr).V0)))
								goto end_branch_0
							} else {

							}
						}
						{
							__t0 = gopurs_runtime.CoerceToStruct[Constructor_Main_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
						}
					end_branch_0:
						__t1 = __t0
						goto end_branch_1
					} else {

					}
				}
				{
					__t1 = gopurs_runtime.CoerceToStruct[Constructor_Main_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
				}
			end_branch_1:
				return gopurs_runtime.Value{Type: 9, IntVal: 3271839782, UnsafePtr: unsafe.Pointer(__t1)}
			})
		})}))}
	})
	return cache_Main_applyMaybe
}

var cache_Main_applicativeMaybe gopurs_runtime.Value
var once_Main_applicativeMaybe sync.Once

func Get_Main_applicativeMaybe() gopurs_runtime.Value {
	once_Main_applicativeMaybe.Do(func() {
		cache_Main_applicativeMaybe = gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Get_Main_applyMaybe()))}
		}), Get_Main_Just()}))}
	})
	return cache_Main_applicativeMaybe
}

var cache_Main_test5 gopurs_runtime.Value
var once_Main_test5 sync.Once

func Get_Main_test5() gopurs_runtime.Value {
	once_Main_test5.Do(func() {
		cache_Main_test5 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3271839782, UnsafePtr: unsafe.Pointer(Call_Main_test5(v_0_box))}
		})
	})
	return cache_Main_test5
}

var cache_Main_monadData gopurs_runtime.Value
var once_Main_monadData sync.Once

func Get_Main_monadData() gopurs_runtime.Value {
	once_Main_monadData.Do(func() {
		cache_Main_monadData = gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Monad{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Main_applicativeData()))}
		}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](Get_Main_bindData()))}
		})}))}
	})
	return cache_Main_monadData
}

var cache_Main_functorData gopurs_runtime.Value
var once_Main_functorData sync.Once

func Get_Main_functorData() gopurs_runtime.Value {
	once_Main_functorData.Do(func() {
		cache_Main_functorData = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor{1, gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](Get_Main_bindData()).V1), a_1, gopurs_runtime.Func(func(a_prime__2 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Main_applicativeData()).V1), gopurs_runtime.Apply(f_0, a_prime__2))
				}))
			})
		})}))}
	})
	return cache_Main_functorData
}

var cache_Main_bindData gopurs_runtime.Value
var once_Main_bindData sync.Once

func Get_Main_bindData() gopurs_runtime.Value {
	once_Main_bindData.Do(func() {
		cache_Main_bindData = gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Get_Main_applyData()))}
		}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(f1_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply(f1_1, v_0)
			})
		})}))}
	})
	return cache_Main_bindData
}

var cache_Main_applyData gopurs_runtime.Value
var once_Main_applyData sync.Once

func Get_Main_applyData() gopurs_runtime.Value {
	once_Main_applyData.Do(func() {
		cache_Main_applyData = gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Main_functorData()))}
		}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Main_applicativeData()).V1), gopurs_runtime.Apply(f_0, a_1))
			})
		})}))}
	})
	return cache_Main_applyData
}

var cache_Main_applicativeData gopurs_runtime.Value
var once_Main_applicativeData sync.Once

func Get_Main_applicativeData() gopurs_runtime.Value {
	once_Main_applicativeData.Do(func() {
		cache_Main_applicativeData = gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Get_Main_applyData()))}
		}), Get_Main_Data()}))}
	})
	return cache_Main_applicativeData
}

var cache_Main_f__83006887 gopurs_runtime.Value
var once_Main_f__83006887 sync.Once

func Get_Main_f__83006887() gopurs_runtime.Value {
	once_Main_f__83006887.Do(func() {
		cache_Main_f__83006887 = gopurs_runtime.Func2(func(dictShow_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_f__83006887(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show](dictShow_0_box), x_1_box))
		})
	})
	return cache_Main_f__83006887
}

var cache_Main_f__2688021959 gopurs_runtime.Value
var once_Main_f__2688021959 sync.Once

func Get_Main_f__2688021959() gopurs_runtime.Value {
	once_Main_f__2688021959.Do(func() {
		cache_Main_f__2688021959 = gopurs_runtime.Func2(func(dictShow_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_f__2688021959(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show](dictShow_0_box), x_1_box.FloatVal()))
		})
	})
	return cache_Main_f__2688021959
}

var cache_Main_f__2742601362 gopurs_runtime.Value
var once_Main_f__2742601362 sync.Once

func Get_Main_f__2742601362() gopurs_runtime.Value {
	once_Main_f__2742601362.Do(func() {
		cache_Main_f__2742601362 = gopurs_runtime.Func2(func(dictShow_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_f__2742601362(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show](dictShow_0_box), x_1_box))
		})
	})
	return cache_Main_f__2742601362
}

var cache_Main_test7__2742601362 gopurs_runtime.Value
var once_Main_test7__2742601362 sync.Once

func Get_Main_test7__2742601362() gopurs_runtime.Value {
	once_Main_test7__2742601362.Do(func() {
		cache_Main_test7__2742601362 = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_test7__2742601362(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show](dictShow_0_box))
		})
	})
	return cache_Main_test7__2742601362
}

type Constructor_Main_Nothing struct {
	Rc uint32
}

type Constructor_Main_Just struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

type Constructor_Main_Data struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

func Call_Main_test8(v_0_loop gopurs_runtime.Value) string {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	return gopurs_runtime.Apply(Get_Data_Show_showStringImpl(), gopurs_runtime.Str("testing")).StrVal()
}

func Call_Main_test7(dictShow_0_loop *Constructor_Data_Show_Show) gopurs_runtime.Value {
	var dictShow_0 *Constructor_Data_Show_Show = dictShow_0_loop
	_ = dictShow_0
	return gopurs_runtime.Box(dictShow_0.V0)
}

func Call_Main_test4(dictMonad_0_loop *Constructor_Control_Monad_Monad) gopurs_runtime.Value {
	var dictMonad_0 *Constructor_Control_Monad_Monad = dictMonad_0_loop
	_ = dictMonad_0
	// TAST (Let): Applicative0_1_0 shape=App(Other) expectedFromAst=*Constructor_Control_Applicative_Applicative actual=*Constructor_Control_Applicative_Applicative bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m)])
	Applicative0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V0), gopurs_runtime.Value{}))
	_ = Applicative0_1_0
	return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_1_0.V1), gopurs_runtime.Float(1.0))
	})
}

func Call_Main_test1(v_0_loop gopurs_runtime.Value) string {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	return gopurs_runtime.Apply(Get_Data_Show_showStringImpl(), gopurs_runtime.Str("testing")).StrVal()
}

func Call_Main_showData(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
	_ = dictShow_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer((&Constructor_Data_Show_Show{1, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Str((("Data (") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), v_1).StrVal())) + (")"))
	})}))}
}

func Call_Main_test3(v_0_loop gopurs_runtime.Value) string {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	return (("Data (") + (gopurs_runtime.Apply(Get_Data_Show_showStringImpl(), gopurs_runtime.Str("testing")).StrVal())) + (")")
}

func Call_Main_runReader(r_0_loop gopurs_runtime.Value, f1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var r_0 gopurs_runtime.Value = r_0_loop
	_ = r_0
	var f1_1 gopurs_runtime.Value = f1_1_loop
	_ = f1_1
	return gopurs_runtime.Apply(f1_1, r_0)
}

func Call_Main_f(dictShow_0_loop *Constructor_Data_Show_Show, x_1_loop gopurs_runtime.Value) string {
	var dictShow_0 *Constructor_Data_Show_Show = dictShow_0_loop
	_ = dictShow_0
	var x_1 gopurs_runtime.Value = x_1_loop
	_ = x_1
	return gopurs_runtime.Apply(gopurs_runtime.Box(dictShow_0.V0), x_1).StrVal()
}

func Call_Main_test2(v_0_loop gopurs_runtime.Value) string {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	return gopurs_runtime.Apply(Get_Data_Show_showStringImpl(), gopurs_runtime.Str("testing")).StrVal()
}

func Call_Main_ask(r_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var r_0 gopurs_runtime.Value = r_0_loop
	_ = r_0
	return r_0
}

func Call_Main_test9(v_0_loop gopurs_runtime.Value) float64 {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	return gopurs_runtime.Float(1.0).FloatVal()
}

func Call_Main_test5(v_0_loop gopurs_runtime.Value) *Constructor_Main_Just {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	return gopurs_runtime.CoerceToStruct[Constructor_Main_Just](gopurs_runtime.Value{Type: 9, IntVal: 3271839782, UnsafePtr: unsafe.Pointer((&Constructor_Main_Just{1, gopurs_runtime.Float(2.0)}))})
}

func Call_Main_f__83006887(dictShow_0_loop *Constructor_Data_Show_Show, x_1_loop gopurs_runtime.Value) string {
	var dictShow_0 *Constructor_Data_Show_Show = dictShow_0_loop
	_ = dictShow_0
	var x_1 gopurs_runtime.Value = x_1_loop
	_ = x_1
	return gopurs_runtime.Apply(gopurs_runtime.Box(dictShow_0.V0), x_1).StrVal()
}

func Call_Main_f__2688021959(dictShow_0_loop *Constructor_Data_Show_Show, x_1_loop float64) string {
	var dictShow_0 *Constructor_Data_Show_Show = dictShow_0_loop
	_ = dictShow_0
	var x_1 float64 = x_1_loop
	_ = x_1
	return gopurs_runtime.Apply(gopurs_runtime.Box(dictShow_0.V0), gopurs_runtime.Float(x_1)).StrVal()
}

func Call_Main_f__2742601362(dictShow_0_loop *Constructor_Data_Show_Show, x_1_loop gopurs_runtime.Value) string {
	var dictShow_0 *Constructor_Data_Show_Show = dictShow_0_loop
	_ = dictShow_0
	var x_1 gopurs_runtime.Value = x_1_loop
	_ = x_1
	return gopurs_runtime.Apply(gopurs_runtime.Box(dictShow_0.V0), x_1).StrVal()
}

func Call_Main_test7__2742601362(dictShow_0_loop *Constructor_Data_Show_Show) gopurs_runtime.Value {
	var dictShow_0 *Constructor_Data_Show_Show = dictShow_0_loop
	_ = dictShow_0
	return gopurs_runtime.Box(dictShow_0.V0)
}
