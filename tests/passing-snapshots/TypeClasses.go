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
		cache_Main_show = Call_Data_Show_show(Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString()))))
	})
	return cache_Main_show
}

var cache_Main_pure gopurs_runtime.Value
var once_Main_pure sync.Once

func Get_Main_pure() gopurs_runtime.Value {
	once_Main_pure.Do(func() {
		cache_Main_pure = Call_Control_Applicative_pure(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Control_Applicative_applicativeFn()))
	})
	return cache_Main_pure
}

var cache_Main_Nothing gopurs_runtime.Value
var once_Main_Nothing sync.Once

func Get_Main_Nothing() gopurs_runtime.Value {
	once_Main_Nothing.Do(func() {
		cache_Main_Nothing = gopurs_runtime.Value{Type: 9, IntVal: 3271839782, UnsafePtr: unsafe.Pointer((*Constructor_Main_Just[gopurs_runtime.Value])(nil))}
	})
	return cache_Main_Nothing
}

var cache_Main_Just gopurs_runtime.Value
var once_Main_Just sync.Once

func Get_Main_Just() gopurs_runtime.Value {
	once_Main_Just.Do(func() {
		cache_Main_Just = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3271839782, UnsafePtr: unsafe.Pointer((&Constructor_Main_Just[gopurs_runtime.Value]{1, value0}))}
		})
	})
	return cache_Main_Just
}

var cache_Main_Just__309914989 gopurs_runtime.Value
var once_Main_Just__309914989 sync.Once

func Get_Main_Just__309914989() gopurs_runtime.Value {
	once_Main_Just__309914989.Do(func() {
		cache_Main_Just__309914989 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3271839782, UnsafePtr: unsafe.Pointer(Rebox_Main_2858737514_4188394610(Call_Main_Just__309914989(__eta_norm_0_0_box.FloatVal())))}
		})
	})
	return cache_Main_Just__309914989
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

var cache_Main_Data__4294158511 gopurs_runtime.Value
var once_Main_Data__4294158511 sync.Once

func Get_Main_Data__4294158511() gopurs_runtime.Value {
	once_Main_Data__4294158511.Do(func() {
		cache_Main_Data__4294158511 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_Data__4294158511(__eta_norm_0_0_box.StrVal())
		})
	})
	return cache_Main_Data__4294158511
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
			return Call_Main_test7(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](dictShow_0_box))
		})
	})
	return cache_Main_test7
}

var cache_Main_test7__2030676012 gopurs_runtime.Value
var once_Main_test7__2030676012 sync.Once

func Get_Main_test7__2030676012() gopurs_runtime.Value {
	once_Main_test7__2030676012.Do(func() {
		cache_Main_test7__2030676012 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_test7__2030676012(__eta_norm_0_0_box.StrVal()))
		})
	})
	return cache_Main_test7__2030676012
}

var cache_Main_test4 gopurs_runtime.Value
var once_Main_test4 sync.Once

func Get_Main_test4() gopurs_runtime.Value {
	once_Main_test4.Do(func() {
		cache_Main_test4 = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_test4(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](dictMonad_0_box))
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
		cache_Main_showData1 = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Call_Main_showData(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString()))))})))}
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
		cache_Main_main = gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Call_Data_Show_show(Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString())))), gopurs_runtime.Str("Hello")).StrVal())), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
		}))
	})
	return cache_Main_main
}

var cache_Main_f gopurs_runtime.Value
var once_Main_f sync.Once

func Get_Main_f() gopurs_runtime.Value {
	once_Main_f.Do(func() {
		cache_Main_f = gopurs_runtime.Func2(func(dictShow_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_f(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](dictShow_0_box), x_1_box))
		})
	})
	return cache_Main_f
}

var cache_Main_f__2030676012 gopurs_runtime.Value
var once_Main_f__2030676012 sync.Once

func Get_Main_f__2030676012() gopurs_runtime.Value {
	once_Main_f__2030676012.Do(func() {
		cache_Main_f__2030676012 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_f__2030676012(x_0_box.StrVal()))
		})
	})
	return cache_Main_f__2030676012
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
		cache_Main_monadMaybe = gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(Rebox_Main_107968622_2568689657((&Constructor_Control_Monad_Monad[*Constructor_Main_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(Rebox_Main_3923698542_1439734649(Rebox_Main_1439734649_3923698542(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Main_applicativeMaybe()))))}
		}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(Rebox_Main_2534018798_2748095225(Rebox_Main_2748095225_2534018798(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Main_bindMaybe()))))}
		})})))}
	})
	return cache_Main_monadMaybe
}

var cache_Main_functorMaybe gopurs_runtime.Value
var once_Main_functorMaybe sync.Once

func Get_Main_functorMaybe() gopurs_runtime.Value {
	once_Main_functorMaybe.Do(func() {
		cache_Main_functorMaybe = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Main_4035997145_2812149806((&Constructor_Data_Functor_Functor[*Constructor_Main_Just[gopurs_runtime.Value]]{1, Call_Control_Monad_liftM1(Rebox_Main_107968622_2568689657(Rebox_Main_2568689657_107968622(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](Get_Main_monadMaybe()))))})))}
	})
	return cache_Main_functorMaybe
}

var cache_Main_bindMaybe gopurs_runtime.Value
var once_Main_bindMaybe sync.Once

func Get_Main_bindMaybe() gopurs_runtime.Value {
	once_Main_bindMaybe.Do(func() {
		cache_Main_bindMaybe = gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(Rebox_Main_2534018798_2748095225((&Constructor_Control_Bind_Bind[*Constructor_Main_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(Rebox_Main_2381757806_3741347833(Rebox_Main_3741347833_2381757806(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Main_applyMaybe()))))}
		}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t0 *Constructor_Main_Just[gopurs_runtime.Value]
			{
				if v_0.Type == 9 && v_0.IntVal == 3271839782 && v_0.UnsafePtr == nil {
					__t0 = (*Constructor_Main_Just[gopurs_runtime.Value])(nil)
					goto end_branch_0
				} else {

				}
			}
			{
				if v_0.Type == 9 && v_0.IntVal == 3271839782 && v_0.UnsafePtr != nil {
					__t0 = gopurs_runtime.CoerceToStruct[Constructor_Main_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(v1_1, (*Constructor_Main_Just[gopurs_runtime.Value])(v_0.UnsafePtr).V0))
					goto end_branch_0
				} else {

				}
			}
			{
				__t0 = func() *Constructor_Main_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
			}
		end_branch_0:
			return gopurs_runtime.Value{Type: 9, IntVal: 3271839782, UnsafePtr: unsafe.Pointer(__t0)}
		})})))}
	})
	return cache_Main_bindMaybe
}

var cache_Main_applyMaybe gopurs_runtime.Value
var once_Main_applyMaybe sync.Once

func Get_Main_applyMaybe() gopurs_runtime.Value {
	once_Main_applyMaybe.Do(func() {
		cache_Main_applyMaybe = gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(Rebox_Main_2381757806_3741347833((&Constructor_Control_Apply_Apply[*Constructor_Main_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Main_4035997145_2812149806(Rebox_Main_2812149806_4035997145(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Main_functorMaybe()))))}
		}), Call_Control_Monad_ap(Rebox_Main_107968622_2568689657(Rebox_Main_2568689657_107968622(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](Get_Main_monadMaybe()))))})))}
	})
	return cache_Main_applyMaybe
}

var cache_Main_applicativeMaybe gopurs_runtime.Value
var once_Main_applicativeMaybe sync.Once

func Get_Main_applicativeMaybe() gopurs_runtime.Value {
	once_Main_applicativeMaybe.Do(func() {
		cache_Main_applicativeMaybe = gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(Rebox_Main_3923698542_1439734649((&Constructor_Control_Applicative_Applicative[*Constructor_Main_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(Rebox_Main_2381757806_3741347833(Rebox_Main_3741347833_2381757806(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Main_applyMaybe()))))}
		}), Get_Main_Just()})))}
	})
	return cache_Main_applicativeMaybe
}

var cache_Main_test5 gopurs_runtime.Value
var once_Main_test5 sync.Once

func Get_Main_test5() gopurs_runtime.Value {
	once_Main_test5.Do(func() {
		cache_Main_test5 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3271839782, UnsafePtr: unsafe.Pointer(Rebox_Main_2858737514_4188394610(Call_Main_test5(v_0_box)))}
		})
	})
	return cache_Main_test5
}

var cache_Main_monadData gopurs_runtime.Value
var once_Main_monadData sync.Once

func Get_Main_monadData() gopurs_runtime.Value {
	once_Main_monadData.Do(func() {
		cache_Main_monadData = gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Main_applicativeData()))}
		}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Main_bindData()))}
		})}))}
	})
	return cache_Main_monadData
}

var cache_Main_functorData gopurs_runtime.Value
var once_Main_functorData sync.Once

func Get_Main_functorData() gopurs_runtime.Value {
	once_Main_functorData.Do(func() {
		cache_Main_functorData = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, Call_Control_Monad_liftM1(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](Get_Main_monadData()))}))}
	})
	return cache_Main_functorData
}

var cache_Main_bindData gopurs_runtime.Value
var once_Main_bindData sync.Once

func Get_Main_bindData() gopurs_runtime.Value {
	once_Main_bindData.Do(func() {
		cache_Main_bindData = gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Main_applyData()))}
		}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, f1_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(f1_1, v_0)
		})}))}
	})
	return cache_Main_bindData
}

var cache_Main_applyData gopurs_runtime.Value
var once_Main_applyData sync.Once

func Get_Main_applyData() gopurs_runtime.Value {
	once_Main_applyData.Do(func() {
		cache_Main_applyData = gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Main_functorData()))}
		}), Call_Control_Monad_ap(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](Get_Main_monadData()))}))}
	})
	return cache_Main_applyData
}

var cache_Main_applicativeData gopurs_runtime.Value
var once_Main_applicativeData sync.Once

func Get_Main_applicativeData() gopurs_runtime.Value {
	once_Main_applicativeData.Do(func() {
		cache_Main_applicativeData = gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Main_applyData()))}
		}), Get_Main_Data()}))}
	})
	return cache_Main_applicativeData
}

type Constructor_Main_Nothing[T_a any] struct {
	Rc uint32
}

type Constructor_Main_Just[T_a any] struct {
	Rc uint32
	V0 T_a
}

type Constructor_Main_Data[T_a any] struct {
	Rc uint32
	V0 T_a
}

func Call_Main_Just__309914989(__eta_norm_0_0_loop float64) *Constructor_Main_Just[float64] {
Just__309914989:
	for {
		if false {
			continue Just__309914989
		}
		var __eta_norm_0_0 float64 = __eta_norm_0_0_loop
		_ = __eta_norm_0_0
		return (&Constructor_Main_Just[float64]{1, __eta_norm_0_0})
	}
}

func Call_Main_Data__4294158511(__eta_norm_0_0_loop string) gopurs_runtime.Value {
Data__4294158511:
	for {
		if false {
			continue Data__4294158511
		}
		var __eta_norm_0_0 string = __eta_norm_0_0_loop
		_ = __eta_norm_0_0
		return gopurs_runtime.Str(__eta_norm_0_0)
	}
}

func Call_Main_test8(v_0_loop gopurs_runtime.Value) string {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	return gopurs_runtime.Apply(Call_Data_Show_show(Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString())))), gopurs_runtime.Str("testing")).StrVal()
}

func Call_Main_test7(dictShow_0_loop *Constructor_Data_Show_Show[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictShow_0 *Constructor_Data_Show_Show[gopurs_runtime.Value] = dictShow_0_loop
	_ = dictShow_0
	return Call_Data_Show_show(dictShow_0)
}

func Call_Main_test7__2030676012(__eta_norm_0_0_loop string) string {
test7__2030676012:
	for {
		if false {
			continue test7__2030676012
		}
		var __eta_norm_0_0 string = __eta_norm_0_0_loop
		_ = __eta_norm_0_0
		return gopurs_runtime.Apply(Call_Data_Show_show(Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString())))), gopurs_runtime.Str(__eta_norm_0_0)).StrVal()
	}
}

func Call_Main_test4(dictMonad_0_loop *Constructor_Control_Monad_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictMonad_0 *Constructor_Control_Monad_Monad[gopurs_runtime.Value] = dictMonad_0_loop
	_ = dictMonad_0
	// TAST (Let): Applicative0_1_0 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m$scope6)])
	Applicative0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonad_0.V0, gopurs_runtime.Value{}))
	_ = Applicative0_1_0
	return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(Applicative0_1_0.V1, gopurs_runtime.Float(1.0))
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
	return gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer((&Constructor_Data_Show_Show[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Str((("Data (") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), v_1).StrVal())) + (")"))
	})}))}
}

func Call_Main_test3(v_0_loop gopurs_runtime.Value) string {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Main_showData(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString()))))}), "show"), gopurs_runtime.Str("testing")).StrVal()
}

func Call_Main_runReader(r_0_loop gopurs_runtime.Value, f1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var r_0 gopurs_runtime.Value = r_0_loop
	_ = r_0
	var f1_1 gopurs_runtime.Value = f1_1_loop
	_ = f1_1
	return gopurs_runtime.Apply(f1_1, r_0)
}

func Call_Main_f(dictShow_0_loop *Constructor_Data_Show_Show[gopurs_runtime.Value], x_1_loop gopurs_runtime.Value) string {
	var dictShow_0 *Constructor_Data_Show_Show[gopurs_runtime.Value] = dictShow_0_loop
	_ = dictShow_0
	var x_1 gopurs_runtime.Value = x_1_loop
	_ = x_1
	return gopurs_runtime.Apply(dictShow_0.V0, x_1).StrVal()
}

func Call_Main_f__2030676012(x_0_loop string) string {
f__2030676012:
	for {
		if false {
			continue f__2030676012
		}
		var x_0 string = x_0_loop
		_ = x_0
		return gopurs_runtime.Apply(Get_Data_Show_showStringImpl(), gopurs_runtime.Str(x_0)).StrVal()
	}
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
	return Call_Main_runReader(gopurs_runtime.Float(0.0), gopurs_runtime.Apply2(Get_Control_Bind_bind__3164072432(), Get_Main_ask(), gopurs_runtime.Func(func(n_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(Call_Control_Applicative_pure(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Control_Applicative_applicativeFn())), gopurs_runtime.Float((n_1.FloatVal())+(1.0)))
	}))).FloatVal()
}

func Call_Main_test5(v_0_loop gopurs_runtime.Value) *Constructor_Main_Just[float64] {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	return Rebox_Main_4188394610_2858737514((&Constructor_Main_Just[gopurs_runtime.Value]{1, gopurs_runtime.Float(2.0)}))
}

func Rebox_Main_107968622_2568689657(in *Constructor_Control_Monad_Monad[*Constructor_Main_Just[gopurs_runtime.Value]]) *Constructor_Control_Monad_Monad[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Control_Monad_Monad[gopurs_runtime.Value]{}
	out.V0 = in.V0
	out.V1 = in.V1
	return out
}

func Rebox_Main_1386611502_1514099793(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[string] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[string]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_1439734649_3923698542(in *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]) *Constructor_Control_Applicative_Applicative[*Constructor_Main_Just[gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	out := &Constructor_Control_Applicative_Applicative[*Constructor_Main_Just[gopurs_runtime.Value]]{}
	out.V0 = in.V0
	out.V1 = in.V1
	return out
}

func Rebox_Main_1514099793_1386611502(in *Constructor_Data_Show_Show[string]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_2381757806_3741347833(in *Constructor_Control_Apply_Apply[*Constructor_Main_Just[gopurs_runtime.Value]]) *Constructor_Control_Apply_Apply[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Control_Apply_Apply[gopurs_runtime.Value]{}
	out.V0 = in.V0
	out.V1 = in.V1
	return out
}

func Rebox_Main_2534018798_2748095225(in *Constructor_Control_Bind_Bind[*Constructor_Main_Just[gopurs_runtime.Value]]) *Constructor_Control_Bind_Bind[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Control_Bind_Bind[gopurs_runtime.Value]{}
	out.V0 = in.V0
	out.V1 = in.V1
	return out
}

func Rebox_Main_2568689657_107968622(in *Constructor_Control_Monad_Monad[gopurs_runtime.Value]) *Constructor_Control_Monad_Monad[*Constructor_Main_Just[gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	out := &Constructor_Control_Monad_Monad[*Constructor_Main_Just[gopurs_runtime.Value]]{}
	out.V0 = in.V0
	out.V1 = in.V1
	return out
}

func Rebox_Main_2748095225_2534018798(in *Constructor_Control_Bind_Bind[gopurs_runtime.Value]) *Constructor_Control_Bind_Bind[*Constructor_Main_Just[gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	out := &Constructor_Control_Bind_Bind[*Constructor_Main_Just[gopurs_runtime.Value]]{}
	out.V0 = in.V0
	out.V1 = in.V1
	return out
}

func Rebox_Main_2812149806_4035997145(in *Constructor_Data_Functor_Functor[gopurs_runtime.Value]) *Constructor_Data_Functor_Functor[*Constructor_Main_Just[gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Functor_Functor[*Constructor_Main_Just[gopurs_runtime.Value]]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_2858737514_4188394610(in *Constructor_Main_Just[float64]) *Constructor_Main_Just[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_Just[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Float(in.V0)
	return out
}

func Rebox_Main_3741347833_2381757806(in *Constructor_Control_Apply_Apply[gopurs_runtime.Value]) *Constructor_Control_Apply_Apply[*Constructor_Main_Just[gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	out := &Constructor_Control_Apply_Apply[*Constructor_Main_Just[gopurs_runtime.Value]]{}
	out.V0 = in.V0
	out.V1 = in.V1
	return out
}

func Rebox_Main_3923698542_1439734649(in *Constructor_Control_Applicative_Applicative[*Constructor_Main_Just[gopurs_runtime.Value]]) *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{}
	out.V0 = in.V0
	out.V1 = in.V1
	return out
}

func Rebox_Main_4035997145_2812149806(in *Constructor_Data_Functor_Functor[*Constructor_Main_Just[gopurs_runtime.Value]]) *Constructor_Data_Functor_Functor[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Functor_Functor[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_4188394610_2858737514(in *Constructor_Main_Just[gopurs_runtime.Value]) *Constructor_Main_Just[float64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_Just[float64]{}
	out.V0 = in.V0.FloatVal()
	return out
}
