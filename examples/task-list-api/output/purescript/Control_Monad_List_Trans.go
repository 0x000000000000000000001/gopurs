package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Monad_List_Trans_compose gopurs_runtime.Value
var once_Control_Monad_List_Trans_compose sync.Once

func Get_Control_Monad_List_Trans_compose() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_compose.Do(func() {
		cache_Control_Monad_List_Trans_compose = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Control_Monad_List_Trans_compose
}

var cache_Control_Monad_List_Trans_compose1 gopurs_runtime.Value
var once_Control_Monad_List_Trans_compose1 sync.Once

func Get_Control_Monad_List_Trans_compose1() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_compose1.Do(func() {
		cache_Control_Monad_List_Trans_compose1 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Control_Monad_List_Trans_compose1
}

var cache_Control_Monad_List_Trans_go__map gopurs_runtime.Value
var once_Control_Monad_List_Trans_go__map sync.Once

func Get_Control_Monad_List_Trans_go__map() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_go__map.Do(func() {
		cache_Control_Monad_List_Trans_go__map = Rebox_Control_Monad_List_Trans_2812149806_3689823567(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Maybe_functorMaybe())).V0
	})
	return cache_Control_Monad_List_Trans_go__map
}

var cache_Control_Monad_List_Trans_map1 gopurs_runtime.Value
var once_Control_Monad_List_Trans_map1 sync.Once

func Get_Control_Monad_List_Trans_map1() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_map1.Do(func() {
		cache_Control_Monad_List_Trans_map1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Lazy_functorLazy()).V0
	})
	return cache_Control_Monad_List_Trans_map1
}

var cache_Control_Monad_List_Trans_map2 gopurs_runtime.Value
var once_Control_Monad_List_Trans_map2 sync.Once

func Get_Control_Monad_List_Trans_map2() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_map2.Do(func() {
		cache_Control_Monad_List_Trans_map2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Lazy_functorLazy()).V0
	})
	return cache_Control_Monad_List_Trans_map2
}

var cache_Control_Monad_List_Trans_sub gopurs_runtime.Value
var once_Control_Monad_List_Trans_sub sync.Once

func Get_Control_Monad_List_Trans_sub() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_sub.Do(func() {
		cache_Control_Monad_List_Trans_sub = Get_Data_Ring_sub__4157804846()
	})
	return cache_Control_Monad_List_Trans_sub
}

var cache_Control_Monad_List_Trans_map3 gopurs_runtime.Value
var once_Control_Monad_List_Trans_map3 sync.Once

func Get_Control_Monad_List_Trans_map3() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_map3.Do(func() {
		cache_Control_Monad_List_Trans_map3 = Rebox_Control_Monad_List_Trans_2812149806_3689823567(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Maybe_functorMaybe())).V0
	})
	return cache_Control_Monad_List_Trans_map3
}

var cache_Control_Monad_List_Trans_map4 gopurs_runtime.Value
var once_Control_Monad_List_Trans_map4 sync.Once

func Get_Control_Monad_List_Trans_map4() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_map4.Do(func() {
		cache_Control_Monad_List_Trans_map4 = gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Lazy_functorLazy()).V0
	})
	return cache_Control_Monad_List_Trans_map4
}

var cache_Control_Monad_List_Trans_identity gopurs_runtime.Value
var once_Control_Monad_List_Trans_identity sync.Once

func Get_Control_Monad_List_Trans_identity() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_identity.Do(func() {
		cache_Control_Monad_List_Trans_identity = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_List_Trans_identity(x_0_box)
		})
	})
	return cache_Control_Monad_List_Trans_identity
}

var cache_Control_Monad_List_Trans_map5 gopurs_runtime.Value
var once_Control_Monad_List_Trans_map5 sync.Once

func Get_Control_Monad_List_Trans_map5() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_map5.Do(func() {
		cache_Control_Monad_List_Trans_map5 = Rebox_Control_Monad_List_Trans_2812149806_3689823567(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Maybe_functorMaybe())).V0
	})
	return cache_Control_Monad_List_Trans_map5
}

var cache_Control_Monad_List_Trans_map6 gopurs_runtime.Value
var once_Control_Monad_List_Trans_map6 sync.Once

func Get_Control_Monad_List_Trans_map6() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_map6.Do(func() {
		cache_Control_Monad_List_Trans_map6 = gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Lazy_functorLazy()).V0
	})
	return cache_Control_Monad_List_Trans_map6
}

var cache_Control_Monad_List_Trans_map7 gopurs_runtime.Value
var once_Control_Monad_List_Trans_map7 sync.Once

func Get_Control_Monad_List_Trans_map7() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_map7.Do(func() {
		cache_Control_Monad_List_Trans_map7 = gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Lazy_functorLazy()).V0
	})
	return cache_Control_Monad_List_Trans_map7
}

var cache_Control_Monad_List_Trans_map8 gopurs_runtime.Value
var once_Control_Monad_List_Trans_map8 sync.Once

func Get_Control_Monad_List_Trans_map8() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_map8.Do(func() {
		cache_Control_Monad_List_Trans_map8 = gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Lazy_functorLazy()).V0
	})
	return cache_Control_Monad_List_Trans_map8
}

var cache_Control_Monad_List_Trans_map9 gopurs_runtime.Value
var once_Control_Monad_List_Trans_map9 sync.Once

func Get_Control_Monad_List_Trans_map9() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_map9.Do(func() {
		cache_Control_Monad_List_Trans_map9 = gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Lazy_functorLazy()).V0
	})
	return cache_Control_Monad_List_Trans_map9
}

var cache_Control_Monad_List_Trans_pure gopurs_runtime.Value
var once_Control_Monad_List_Trans_pure sync.Once

func Get_Control_Monad_List_Trans_pure() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_pure.Do(func() {
		cache_Control_Monad_List_Trans_pure = gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Data_Lazy_applicativeLazy()).V1
	})
	return cache_Control_Monad_List_Trans_pure
}

var cache_Control_Monad_List_Trans_pure1 gopurs_runtime.Value
var once_Control_Monad_List_Trans_pure1 sync.Once

func Get_Control_Monad_List_Trans_pure1() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_pure1.Do(func() {
		cache_Control_Monad_List_Trans_pure1 = gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Data_Lazy_applicativeLazy()).V1
	})
	return cache_Control_Monad_List_Trans_pure1
}

var cache_Control_Monad_List_Trans_map10 gopurs_runtime.Value
var once_Control_Monad_List_Trans_map10 sync.Once

func Get_Control_Monad_List_Trans_map10() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_map10.Do(func() {
		cache_Control_Monad_List_Trans_map10 = gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Lazy_functorLazy()).V0
	})
	return cache_Control_Monad_List_Trans_map10
}

var cache_Control_Monad_List_Trans_identity1 gopurs_runtime.Value
var once_Control_Monad_List_Trans_identity1 sync.Once

func Get_Control_Monad_List_Trans_identity1() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_identity1.Do(func() {
		cache_Control_Monad_List_Trans_identity1 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Control_Monad_List_Trans_identity1(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](x_0_box))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Control_Monad_List_Trans_identity1
}

var cache_Control_Monad_List_Trans_map11 gopurs_runtime.Value
var once_Control_Monad_List_Trans_map11 sync.Once

func Get_Control_Monad_List_Trans_map11() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_map11.Do(func() {
		cache_Control_Monad_List_Trans_map11 = gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Lazy_functorLazy()).V0
	})
	return cache_Control_Monad_List_Trans_map11
}

var cache_Control_Monad_List_Trans_compose2 gopurs_runtime.Value
var once_Control_Monad_List_Trans_compose2 sync.Once

func Get_Control_Monad_List_Trans_compose2() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_compose2.Do(func() {
		cache_Control_Monad_List_Trans_compose2 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Control_Monad_List_Trans_compose2
}

var cache_Control_Monad_List_Trans_compose3 gopurs_runtime.Value
var once_Control_Monad_List_Trans_compose3 sync.Once

func Get_Control_Monad_List_Trans_compose3() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_compose3.Do(func() {
		cache_Control_Monad_List_Trans_compose3 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Control_Monad_List_Trans_compose3
}

var cache_Control_Monad_List_Trans_Yield gopurs_runtime.Value
var once_Control_Monad_List_Trans_Yield sync.Once

func Get_Control_Monad_List_Trans_Yield() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_Yield.Do(func() {
		cache_Control_Monad_List_Trans_Yield = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, value0, value1}))}
			})
		})
	})
	return cache_Control_Monad_List_Trans_Yield
}

var cache_Control_Monad_List_Trans_Skip gopurs_runtime.Value
var once_Control_Monad_List_Trans_Skip sync.Once

func Get_Control_Monad_List_Trans_Skip() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_Skip.Do(func() {
		cache_Control_Monad_List_Trans_Skip = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, value0}))}
		})
	})
	return cache_Control_Monad_List_Trans_Skip
}

var cache_Control_Monad_List_Trans_Done gopurs_runtime.Value
var once_Control_Monad_List_Trans_Done sync.Once

func Get_Control_Monad_List_Trans_Done() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_Done.Do(func() {
		cache_Control_Monad_List_Trans_Done = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}
	})
	return cache_Control_Monad_List_Trans_Done
}

var cache_Control_Monad_List_Trans_ListT gopurs_runtime.Value
var once_Control_Monad_List_Trans_ListT sync.Once

func Get_Control_Monad_List_Trans_ListT() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_ListT.Do(func() {
		cache_Control_Monad_List_Trans_ListT = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_List_Trans_ListT(x_0_box)
		})
	})
	return cache_Control_Monad_List_Trans_ListT
}

var cache_Control_Monad_List_Trans_wrapLazy gopurs_runtime.Value
var once_Control_Monad_List_Trans_wrapLazy sync.Once

func Get_Control_Monad_List_Trans_wrapLazy() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_wrapLazy.Do(func() {
		cache_Control_Monad_List_Trans_wrapLazy = gopurs_runtime.Func2(func(dictApplicative_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_List_Trans_wrapLazy(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0_box), v_1_box)
		})
	})
	return cache_Control_Monad_List_Trans_wrapLazy
}

var cache_Control_Monad_List_Trans_wrapEffect gopurs_runtime.Value
var once_Control_Monad_List_Trans_wrapEffect sync.Once

func Get_Control_Monad_List_Trans_wrapEffect() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_wrapEffect.Do(func() {
		cache_Control_Monad_List_Trans_wrapEffect = gopurs_runtime.Func2(func(dictFunctor_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_List_Trans_wrapEffect(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](dictFunctor_0_box), v_1_box)
		})
	})
	return cache_Control_Monad_List_Trans_wrapEffect
}

var cache_Control_Monad_List_Trans_unfold gopurs_runtime.Value
var once_Control_Monad_List_Trans_unfold sync.Once

func Get_Control_Monad_List_Trans_unfold() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_unfold.Do(func() {
		cache_Control_Monad_List_Trans_unfold = gopurs_runtime.Func3(func(dictMonad_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, z_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_List_Trans_unfold(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](dictMonad_0_box), f_1_box, z_2_box)
		})
	})
	return cache_Control_Monad_List_Trans_unfold
}

var cache_Control_Monad_List_Trans_uncons gopurs_runtime.Value
var once_Control_Monad_List_Trans_uncons sync.Once

func Get_Control_Monad_List_Trans_uncons() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_uncons.Do(func() {
		cache_Control_Monad_List_Trans_uncons = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_List_Trans_uncons(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](dictMonad_0_box))
		})
	})
	return cache_Control_Monad_List_Trans_uncons
}

var cache_Control_Monad_List_Trans_tail gopurs_runtime.Value
var once_Control_Monad_List_Trans_tail sync.Once

func Get_Control_Monad_List_Trans_tail() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_tail.Do(func() {
		cache_Control_Monad_List_Trans_tail = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_List_Trans_tail(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](dictMonad_0_box))
		})
	})
	return cache_Control_Monad_List_Trans_tail
}

var cache_Control_Monad_List_Trans_stepMap gopurs_runtime.Value
var once_Control_Monad_List_Trans_stepMap sync.Once

func Get_Control_Monad_List_Trans_stepMap() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_stepMap.Do(func() {
		cache_Control_Monad_List_Trans_stepMap = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_List_Trans_stepMap(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](dictFunctor_0_box), f_1_box, v_2_box)
		})
	})
	return cache_Control_Monad_List_Trans_stepMap
}

var cache_Control_Monad_List_Trans_takeWhile gopurs_runtime.Value
var once_Control_Monad_List_Trans_takeWhile sync.Once

func Get_Control_Monad_List_Trans_takeWhile() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_takeWhile.Do(func() {
		cache_Control_Monad_List_Trans_takeWhile = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_List_Trans_takeWhile(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0_box))
		})
	})
	return cache_Control_Monad_List_Trans_takeWhile
}

var cache_Control_Monad_List_Trans_scanl gopurs_runtime.Value
var once_Control_Monad_List_Trans_scanl sync.Once

func Get_Control_Monad_List_Trans_scanl() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_scanl.Do(func() {
		cache_Control_Monad_List_Trans_scanl = gopurs_runtime.Func4(func(dictMonad_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, b_2_box gopurs_runtime.Value, l_3_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_List_Trans_scanl(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](dictMonad_0_box), f_1_box, b_2_box, l_3_box)
		})
	})
	return cache_Control_Monad_List_Trans_scanl
}

var cache_Control_Monad_List_Trans_prepend_prime_ gopurs_runtime.Value
var once_Control_Monad_List_Trans_prepend_prime_ sync.Once

func Get_Control_Monad_List_Trans_prepend_prime_() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_prepend_prime_.Do(func() {
		cache_Control_Monad_List_Trans_prepend_prime_ = gopurs_runtime.Func3(func(dictApplicative_0_box gopurs_runtime.Value, h_1_box gopurs_runtime.Value, t_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_List_Trans_prepend_prime_(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0_box), h_1_box, t_2_box)
		})
	})
	return cache_Control_Monad_List_Trans_prepend_prime_
}

var cache_Control_Monad_List_Trans_prepend gopurs_runtime.Value
var once_Control_Monad_List_Trans_prepend sync.Once

func Get_Control_Monad_List_Trans_prepend() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_prepend.Do(func() {
		cache_Control_Monad_List_Trans_prepend = gopurs_runtime.Func3(func(dictApplicative_0_box gopurs_runtime.Value, h_1_box gopurs_runtime.Value, t_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_List_Trans_prepend(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0_box), h_1_box, t_2_box)
		})
	})
	return cache_Control_Monad_List_Trans_prepend
}

var cache_Control_Monad_List_Trans_nil gopurs_runtime.Value
var once_Control_Monad_List_Trans_nil sync.Once

func Get_Control_Monad_List_Trans_nil() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_nil.Do(func() {
		cache_Control_Monad_List_Trans_nil = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_List_Trans_nil(dictApplicative_0_box)
		})
	})
	return cache_Control_Monad_List_Trans_nil
}

var cache_Control_Monad_List_Trans_singleton gopurs_runtime.Value
var once_Control_Monad_List_Trans_singleton sync.Once

func Get_Control_Monad_List_Trans_singleton() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_singleton.Do(func() {
		cache_Control_Monad_List_Trans_singleton = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_List_Trans_singleton(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0_box))
		})
	})
	return cache_Control_Monad_List_Trans_singleton
}

var cache_Control_Monad_List_Trans_take gopurs_runtime.Value
var once_Control_Monad_List_Trans_take sync.Once

func Get_Control_Monad_List_Trans_take() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_take.Do(func() {
		cache_Control_Monad_List_Trans_take = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_List_Trans_take(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0_box))
		})
	})
	return cache_Control_Monad_List_Trans_take
}

var cache_Control_Monad_List_Trans_zipWith_prime_ gopurs_runtime.Value
var once_Control_Monad_List_Trans_zipWith_prime_ sync.Once

func Get_Control_Monad_List_Trans_zipWith_prime_() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_zipWith_prime_.Do(func() {
		cache_Control_Monad_List_Trans_zipWith_prime_ = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_List_Trans_zipWith_prime_(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](dictMonad_0_box))
		})
	})
	return cache_Control_Monad_List_Trans_zipWith_prime_
}

var cache_Control_Monad_List_Trans_zipWith gopurs_runtime.Value
var once_Control_Monad_List_Trans_zipWith sync.Once

func Get_Control_Monad_List_Trans_zipWith() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_zipWith.Do(func() {
		cache_Control_Monad_List_Trans_zipWith = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_List_Trans_zipWith(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](dictMonad_0_box))
		})
	})
	return cache_Control_Monad_List_Trans_zipWith
}

var cache_Control_Monad_List_Trans_newtypeListT gopurs_runtime.Value
var once_Control_Monad_List_Trans_newtypeListT sync.Once

func Get_Control_Monad_List_Trans_newtypeListT() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_newtypeListT.Do(func() {
		cache_Control_Monad_List_Trans_newtypeListT = gopurs_runtime.Value{Type: 9, IntVal: 3322196858, UnsafePtr: unsafe.Pointer((&Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{}
		})}))}
	})
	return cache_Control_Monad_List_Trans_newtypeListT
}

var cache_Control_Monad_List_Trans_mapMaybe gopurs_runtime.Value
var once_Control_Monad_List_Trans_mapMaybe sync.Once

func Get_Control_Monad_List_Trans_mapMaybe() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_mapMaybe.Do(func() {
		cache_Control_Monad_List_Trans_mapMaybe = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_List_Trans_mapMaybe(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](dictFunctor_0_box), f_1_box, v_2_box)
		})
	})
	return cache_Control_Monad_List_Trans_mapMaybe
}

var cache_Control_Monad_List_Trans_iterate gopurs_runtime.Value
var once_Control_Monad_List_Trans_iterate sync.Once

func Get_Control_Monad_List_Trans_iterate() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_iterate.Do(func() {
		cache_Control_Monad_List_Trans_iterate = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_List_Trans_iterate(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](dictMonad_0_box))
		})
	})
	return cache_Control_Monad_List_Trans_iterate
}

var cache_Control_Monad_List_Trans_repeat gopurs_runtime.Value
var once_Control_Monad_List_Trans_repeat sync.Once

func Get_Control_Monad_List_Trans_repeat() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_repeat.Do(func() {
		cache_Control_Monad_List_Trans_repeat = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_List_Trans_repeat(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](dictMonad_0_box))
		})
	})
	return cache_Control_Monad_List_Trans_repeat
}

var cache_Control_Monad_List_Trans_head gopurs_runtime.Value
var once_Control_Monad_List_Trans_head sync.Once

func Get_Control_Monad_List_Trans_head() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_head.Do(func() {
		cache_Control_Monad_List_Trans_head = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_List_Trans_head(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](dictMonad_0_box))
		})
	})
	return cache_Control_Monad_List_Trans_head
}

var cache_Control_Monad_List_Trans_functorListT gopurs_runtime.Value
var once_Control_Monad_List_Trans_functorListT sync.Once

func Get_Control_Monad_List_Trans_functorListT() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_functorListT.Do(func() {
		cache_Control_Monad_List_Trans_functorListT = gopurs_runtime.Func(func(dictFunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_List_Trans_functorListT(dictFunctor_0_box)
		})
	})
	return cache_Control_Monad_List_Trans_functorListT
}

var cache_Control_Monad_List_Trans_fromEffect gopurs_runtime.Value
var once_Control_Monad_List_Trans_fromEffect sync.Once

func Get_Control_Monad_List_Trans_fromEffect() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_fromEffect.Do(func() {
		cache_Control_Monad_List_Trans_fromEffect = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_List_Trans_fromEffect(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0_box))
		})
	})
	return cache_Control_Monad_List_Trans_fromEffect
}

var cache_Control_Monad_List_Trans_monadTransListT gopurs_runtime.Value
var once_Control_Monad_List_Trans_monadTransListT sync.Once

func Get_Control_Monad_List_Trans_monadTransListT() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_monadTransListT.Do(func() {
		cache_Control_Monad_List_Trans_monadTransListT = gopurs_runtime.Value{Type: 9, IntVal: 2835982595, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Trans_Class_MonadTrans[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_List_Trans_fromEffect(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{})))
		})}))}
	})
	return cache_Control_Monad_List_Trans_monadTransListT
}

var cache_Control_Monad_List_Trans_lift gopurs_runtime.Value
var once_Control_Monad_List_Trans_lift sync.Once

func Get_Control_Monad_List_Trans_lift() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_lift.Do(func() {
		cache_Control_Monad_List_Trans_lift = gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Trans_Class_MonadTrans[gopurs_runtime.Value]](Get_Control_Monad_List_Trans_monadTransListT()).V0
	})
	return cache_Control_Monad_List_Trans_lift
}

var cache_Control_Monad_List_Trans_foldlRec_prime_ gopurs_runtime.Value
var once_Control_Monad_List_Trans_foldlRec_prime_ sync.Once

func Get_Control_Monad_List_Trans_foldlRec_prime_() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_foldlRec_prime_.Do(func() {
		cache_Control_Monad_List_Trans_foldlRec_prime_ = gopurs_runtime.Func(func(dictMonadRec_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_List_Trans_foldlRec_prime_(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]](dictMonadRec_0_box))
		})
	})
	return cache_Control_Monad_List_Trans_foldlRec_prime_
}

var cache_Control_Monad_List_Trans_runListTRec gopurs_runtime.Value
var once_Control_Monad_List_Trans_runListTRec sync.Once

func Get_Control_Monad_List_Trans_runListTRec() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_runListTRec.Do(func() {
		cache_Control_Monad_List_Trans_runListTRec = gopurs_runtime.Func(func(dictMonadRec_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_List_Trans_runListTRec(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]](dictMonadRec_0_box))
		})
	})
	return cache_Control_Monad_List_Trans_runListTRec
}

var cache_Control_Monad_List_Trans_foldlRec gopurs_runtime.Value
var once_Control_Monad_List_Trans_foldlRec sync.Once

func Get_Control_Monad_List_Trans_foldlRec() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_foldlRec.Do(func() {
		cache_Control_Monad_List_Trans_foldlRec = gopurs_runtime.Func(func(dictMonadRec_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_List_Trans_foldlRec(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]](dictMonadRec_0_box))
		})
	})
	return cache_Control_Monad_List_Trans_foldlRec
}

var cache_Control_Monad_List_Trans_foldl_prime_ gopurs_runtime.Value
var once_Control_Monad_List_Trans_foldl_prime_ sync.Once

func Get_Control_Monad_List_Trans_foldl_prime_() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_foldl_prime_.Do(func() {
		cache_Control_Monad_List_Trans_foldl_prime_ = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_List_Trans_foldl_prime_(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](dictMonad_0_box))
		})
	})
	return cache_Control_Monad_List_Trans_foldl_prime_
}

var cache_Control_Monad_List_Trans_runListT gopurs_runtime.Value
var once_Control_Monad_List_Trans_runListT sync.Once

func Get_Control_Monad_List_Trans_runListT() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_runListT.Do(func() {
		cache_Control_Monad_List_Trans_runListT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_List_Trans_runListT(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](dictMonad_0_box))
		})
	})
	return cache_Control_Monad_List_Trans_runListT
}

var cache_Control_Monad_List_Trans_foldl gopurs_runtime.Value
var once_Control_Monad_List_Trans_foldl sync.Once

func Get_Control_Monad_List_Trans_foldl() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_foldl.Do(func() {
		cache_Control_Monad_List_Trans_foldl = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_List_Trans_foldl(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](dictMonad_0_box))
		})
	})
	return cache_Control_Monad_List_Trans_foldl
}

var cache_Control_Monad_List_Trans_filter gopurs_runtime.Value
var once_Control_Monad_List_Trans_filter sync.Once

func Get_Control_Monad_List_Trans_filter() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_filter.Do(func() {
		cache_Control_Monad_List_Trans_filter = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_List_Trans_filter(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](dictFunctor_0_box), f_1_box, v_2_box)
		})
	})
	return cache_Control_Monad_List_Trans_filter
}

var cache_Control_Monad_List_Trans_dropWhile gopurs_runtime.Value
var once_Control_Monad_List_Trans_dropWhile sync.Once

func Get_Control_Monad_List_Trans_dropWhile() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_dropWhile.Do(func() {
		cache_Control_Monad_List_Trans_dropWhile = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_List_Trans_dropWhile(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0_box))
		})
	})
	return cache_Control_Monad_List_Trans_dropWhile
}

var cache_Control_Monad_List_Trans_drop gopurs_runtime.Value
var once_Control_Monad_List_Trans_drop sync.Once

func Get_Control_Monad_List_Trans_drop() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_drop.Do(func() {
		cache_Control_Monad_List_Trans_drop = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_List_Trans_drop(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0_box))
		})
	})
	return cache_Control_Monad_List_Trans_drop
}

var cache_Control_Monad_List_Trans_cons gopurs_runtime.Value
var once_Control_Monad_List_Trans_cons sync.Once

func Get_Control_Monad_List_Trans_cons() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_cons.Do(func() {
		cache_Control_Monad_List_Trans_cons = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_List_Trans_cons(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0_box))
		})
	})
	return cache_Control_Monad_List_Trans_cons
}

var cache_Control_Monad_List_Trans_unfoldable1ListT gopurs_runtime.Value
var once_Control_Monad_List_Trans_unfoldable1ListT sync.Once

func Get_Control_Monad_List_Trans_unfoldable1ListT() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_unfoldable1ListT.Do(func() {
		cache_Control_Monad_List_Trans_unfoldable1ListT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_List_Trans_unfoldable1ListT(dictMonad_0_box)
		})
	})
	return cache_Control_Monad_List_Trans_unfoldable1ListT
}

var cache_Control_Monad_List_Trans_unfoldableListT gopurs_runtime.Value
var once_Control_Monad_List_Trans_unfoldableListT sync.Once

func Get_Control_Monad_List_Trans_unfoldableListT() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_unfoldableListT.Do(func() {
		cache_Control_Monad_List_Trans_unfoldableListT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_List_Trans_unfoldableListT(dictMonad_0_box)
		})
	})
	return cache_Control_Monad_List_Trans_unfoldableListT
}

var cache_Control_Monad_List_Trans_semigroupListT gopurs_runtime.Value
var once_Control_Monad_List_Trans_semigroupListT sync.Once

func Get_Control_Monad_List_Trans_semigroupListT() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_semigroupListT.Do(func() {
		cache_Control_Monad_List_Trans_semigroupListT = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_List_Trans_semigroupListT(dictApplicative_0_box)
		})
	})
	return cache_Control_Monad_List_Trans_semigroupListT
}

var cache_Control_Monad_List_Trans_concat gopurs_runtime.Value
var once_Control_Monad_List_Trans_concat sync.Once

func Get_Control_Monad_List_Trans_concat() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_concat.Do(func() {
		cache_Control_Monad_List_Trans_concat = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_List_Trans_concat(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0_box))
		})
	})
	return cache_Control_Monad_List_Trans_concat
}

var cache_Control_Monad_List_Trans_monoidListT gopurs_runtime.Value
var once_Control_Monad_List_Trans_monoidListT sync.Once

func Get_Control_Monad_List_Trans_monoidListT() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_monoidListT.Do(func() {
		cache_Control_Monad_List_Trans_monoidListT = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_List_Trans_monoidListT(dictApplicative_0_box)
		})
	})
	return cache_Control_Monad_List_Trans_monoidListT
}

var cache_Control_Monad_List_Trans_catMaybes gopurs_runtime.Value
var once_Control_Monad_List_Trans_catMaybes sync.Once

func Get_Control_Monad_List_Trans_catMaybes() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_catMaybes.Do(func() {
		cache_Control_Monad_List_Trans_catMaybes = gopurs_runtime.Func(func(dictFunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_List_Trans_catMaybes(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](dictFunctor_0_box))
		})
	})
	return cache_Control_Monad_List_Trans_catMaybes
}

var cache_Control_Monad_List_Trans_monadListT gopurs_runtime.Value
var once_Control_Monad_List_Trans_monadListT sync.Once

func Get_Control_Monad_List_Trans_monadListT() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_monadListT.Do(func() {
		cache_Control_Monad_List_Trans_monadListT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_List_Trans_monadListT(dictMonad_0_box)
		})
	})
	return cache_Control_Monad_List_Trans_monadListT
}

var cache_Control_Monad_List_Trans_bindListT gopurs_runtime.Value
var once_Control_Monad_List_Trans_bindListT sync.Once

func Get_Control_Monad_List_Trans_bindListT() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_bindListT.Do(func() {
		cache_Control_Monad_List_Trans_bindListT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_List_Trans_bindListT(dictMonad_0_box)
		})
	})
	return cache_Control_Monad_List_Trans_bindListT
}

var cache_Control_Monad_List_Trans_applyListT gopurs_runtime.Value
var once_Control_Monad_List_Trans_applyListT sync.Once

func Get_Control_Monad_List_Trans_applyListT() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_applyListT.Do(func() {
		cache_Control_Monad_List_Trans_applyListT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_List_Trans_applyListT(dictMonad_0_box)
		})
	})
	return cache_Control_Monad_List_Trans_applyListT
}

var cache_Control_Monad_List_Trans_applicativeListT gopurs_runtime.Value
var once_Control_Monad_List_Trans_applicativeListT sync.Once

func Get_Control_Monad_List_Trans_applicativeListT() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_applicativeListT.Do(func() {
		cache_Control_Monad_List_Trans_applicativeListT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_List_Trans_applicativeListT(dictMonad_0_box)
		})
	})
	return cache_Control_Monad_List_Trans_applicativeListT
}

var cache_Control_Monad_List_Trans_monadEffectListT gopurs_runtime.Value
var once_Control_Monad_List_Trans_monadEffectListT sync.Once

func Get_Control_Monad_List_Trans_monadEffectListT() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_monadEffectListT.Do(func() {
		cache_Control_Monad_List_Trans_monadEffectListT = gopurs_runtime.Func(func(dictMonadEffect_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_List_Trans_monadEffectListT(dictMonadEffect_0_box)
		})
	})
	return cache_Control_Monad_List_Trans_monadEffectListT
}

var cache_Control_Monad_List_Trans_monadSTListT gopurs_runtime.Value
var once_Control_Monad_List_Trans_monadSTListT sync.Once

func Get_Control_Monad_List_Trans_monadSTListT() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_monadSTListT.Do(func() {
		cache_Control_Monad_List_Trans_monadSTListT = gopurs_runtime.Func(func(dictMonadST_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_List_Trans_monadSTListT(dictMonadST_0_box)
		})
	})
	return cache_Control_Monad_List_Trans_monadSTListT
}

var cache_Control_Monad_List_Trans_altListT gopurs_runtime.Value
var once_Control_Monad_List_Trans_altListT sync.Once

func Get_Control_Monad_List_Trans_altListT() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_altListT.Do(func() {
		cache_Control_Monad_List_Trans_altListT = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_List_Trans_altListT(dictApplicative_0_box)
		})
	})
	return cache_Control_Monad_List_Trans_altListT
}

var cache_Control_Monad_List_Trans_plusListT gopurs_runtime.Value
var once_Control_Monad_List_Trans_plusListT sync.Once

func Get_Control_Monad_List_Trans_plusListT() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_plusListT.Do(func() {
		cache_Control_Monad_List_Trans_plusListT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_List_Trans_plusListT(dictMonad_0_box)
		})
	})
	return cache_Control_Monad_List_Trans_plusListT
}

var cache_Control_Monad_List_Trans_alternativeListT gopurs_runtime.Value
var once_Control_Monad_List_Trans_alternativeListT sync.Once

func Get_Control_Monad_List_Trans_alternativeListT() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_alternativeListT.Do(func() {
		cache_Control_Monad_List_Trans_alternativeListT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_List_Trans_alternativeListT(dictMonad_0_box)
		})
	})
	return cache_Control_Monad_List_Trans_alternativeListT
}

var cache_Control_Monad_List_Trans_monadPlusListT gopurs_runtime.Value
var once_Control_Monad_List_Trans_monadPlusListT sync.Once

func Get_Control_Monad_List_Trans_monadPlusListT() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_monadPlusListT.Do(func() {
		cache_Control_Monad_List_Trans_monadPlusListT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_List_Trans_monadPlusListT(dictMonad_0_box)
		})
	})
	return cache_Control_Monad_List_Trans_monadPlusListT
}

type Constructor_Control_Monad_List_Trans_Yield[T_a any, T_s any] struct {
	Rc uint32
	V0 T_a
	V1 gopurs_runtime.Value
}

type Constructor_Control_Monad_List_Trans_Skip[T_a any, T_s any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

type Constructor_Control_Monad_List_Trans_Done[T_a any, T_s any] struct {
	Rc uint32
}

func Call_Control_Monad_List_Trans_identity(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Control_Monad_List_Trans_identity1(x_0_loop *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
	var x_0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = x_0_loop
	_ = x_0
	return func() struct {
		V0 gopurs_runtime.Value
		V1 bool
	} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(x_0)}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
		return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
	}; return struct {
		V0 gopurs_runtime.Value
		V1 bool
	}{V0: gopurs_runtime.Value{}, V1: false} }()
}

func Call_Control_Monad_List_Trans_ListT(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Control_Monad_List_Trans_wrapLazy(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value], v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictApplicative_0 *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
	_ = dictApplicative_0
	var v_1 gopurs_runtime.Value = v_1_loop
	_ = v_1
	return gopurs_runtime.Apply(dictApplicative_0.V1, gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, v_1}))})
}

func Call_Control_Monad_List_Trans_wrapEffect(dictFunctor_0_loop *Constructor_Data_Functor_Functor[gopurs_runtime.Value], v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictFunctor_0 *Constructor_Data_Functor_Functor[gopurs_runtime.Value] = dictFunctor_0_loop
	_ = dictFunctor_0
	var v_1 gopurs_runtime.Value = v_1_loop
	_ = v_1
	return gopurs_runtime.Apply2(dictFunctor_0.V0, gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Control_Monad_List_Trans_Skip(), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Lazy_go__defer(), Get_Data_Function_go__const())), v_1)
}

func Call_Control_Monad_List_Trans_unfold(dictMonad_0_loop *Constructor_Control_Monad_Monad[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, z_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
unfold:
	for {
		if false {
			continue unfold
		}
		var dictMonad_0 *Constructor_Control_Monad_Monad[gopurs_runtime.Value] = dictMonad_0_loop
		_ = dictMonad_0
		var f_1 gopurs_runtime.Value = f_1_loop
		_ = f_1
		var z_2 gopurs_runtime.Value = z_2_loop
		_ = z_2
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictMonad_0.V1, gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t1 gopurs_runtime.Value
			{
				if v_3.Type == 9 && v_3.IntVal == 930809136 && v_3.UnsafePtr != nil {
					// TAST (Let): __local_var_4_0 shape=Other bindingType=Any
					__local_var_4_0 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v_3.UnsafePtr).V0.UnsafePtr).V0
					_ = __local_var_4_0
					__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v_3.UnsafePtr).V0.UnsafePtr).V1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
						return Call_Control_Monad_List_Trans_unfold(dictMonad_0, f_1, __local_var_4_0)
					}))}))}
					goto end_branch_1
				} else {

				}
			}
			{
				if v_3.Type == 9 && v_3.IntVal == 930809136 && v_3.UnsafePtr == nil {
					__t1 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}
					goto end_branch_1
				} else {

				}
			}
			{
				__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_1:
			return __t1
		}), gopurs_runtime.Apply(f_1, z_2))
	}
}

func Call_Control_Monad_List_Trans_uncons(dictMonad_0_loop *Constructor_Control_Monad_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
uncons:
	for {
		if false {
			continue uncons
		}
		var dictMonad_0 *Constructor_Control_Monad_Monad[gopurs_runtime.Value] = dictMonad_0_loop
		_ = dictMonad_0
		// TAST (Let): pure2_1_0 shape=App(Var) bindingType=(Func [(ADT ["Data","Maybe","Maybe"] [(ADT ["Data","Tuple","Tuple"] [(TypeVar a$scope47), (TypeApp (TypeVar f$scope46) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a$scope47), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f$scope46), (TypeVar a$scope47)])])])])])] (TypeApp (TypeVar f$scope46) [(ADT ["Data","Maybe","Maybe"] [(ADT ["Data","Tuple","Tuple"] [(TypeVar a$scope47), (TypeApp (TypeVar f$scope46) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a$scope47), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f$scope46), (TypeVar a$scope47)])])])])])]))
		pure2_1_0 := Call_Control_Applicative_pure(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonad_0.V0, gopurs_runtime.Value{})))
		_ = pure2_1_0
		return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictMonad_0.V1, gopurs_runtime.Value{}), "bind"), v_2, gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t1 gopurs_runtime.Value
				{
					if v1_3.Type == 9 && v1_3.IntVal == 1320412129 {
						__t1 = gopurs_runtime.Apply(pure2_1_0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Control_Monad_List_Trans_4010058633_3094389156(Rebox_Control_Monad_List_Trans_3094389156_4010058633(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
							_v := struct {
								V0 gopurs_runtime.Value
								V1 bool
							}{gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer((&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0, gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V1)}))}, true}
							if _v.V1 {
								return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
							}
							return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
						}()))))})
						goto end_branch_1
					} else {

					}
				}
				{
					if v1_3.Type == 9 && v1_3.IntVal == 813447293 {
						__t1 = gopurs_runtime.Apply(Call_Control_Monad_List_Trans_uncons(dictMonad_0), gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0))
						goto end_branch_1
					} else {

					}
				}
				{
					if v1_3.Type == 9 && v1_3.IntVal == 489128924 {
						__t1 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictMonad_0.V0, gopurs_runtime.Value{}), "pure"), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Control_Monad_List_Trans_4010058633_3094389156(Rebox_Control_Monad_List_Trans_3094389156_4010058633(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))))})
						goto end_branch_1
					} else {

					}
				}
				{
					__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
				}
			end_branch_1:
				return __t1
			}))
		})
	}
}

func Call_Control_Monad_List_Trans_tail(dictMonad_0_loop *Constructor_Control_Monad_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictMonad_0 *Constructor_Control_Monad_Monad[gopurs_runtime.Value] = dictMonad_0_loop
	_ = dictMonad_0
	// TAST (Let): uncons1_1_0 shape=App(Var) bindingType=(Func [(TypeApp (TypeVar f$scope58) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a$scope59), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f$scope58), (TypeVar a$scope59)])])])] (TypeApp (TypeVar f$scope58) [(ADT ["Data","Maybe","Maybe"] [(ADT ["Data","Tuple","Tuple"] [(TypeVar a$scope59), (TypeApp (TypeVar f$scope58) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a$scope59), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f$scope58), (TypeVar a$scope59)])])])])])]))
	uncons1_1_0 := Call_Control_Monad_List_Trans_uncons(dictMonad_0)
	_ = uncons1_1_0
	return gopurs_runtime.Func(func(l_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictMonad_0.V1, gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t1 struct {
				V0 gopurs_runtime.Value
				V1 bool
			}
			{
				if v1_3.Type == 9 && v1_3.IntVal == 930809136 && v1_3.UnsafePtr != nil {
					__t1 = struct {
						V0 gopurs_runtime.Value
						V1 bool
					}{Call_Data_Tuple_snd(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v1_3.UnsafePtr).V0)), true}
					goto end_branch_1
				} else {

				}
			}
			{
				__t1 = struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Value{}, false}
			}
		end_branch_1:
			return func() gopurs_runtime.Value {
				_v := __t1
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		}), gopurs_runtime.Apply(uncons1_1_0, l_2))
	})
}

func Call_Control_Monad_List_Trans_stepMap(dictFunctor_0_loop *Constructor_Data_Functor_Functor[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictFunctor_0 *Constructor_Data_Functor_Functor[gopurs_runtime.Value] = dictFunctor_0_loop
	_ = dictFunctor_0
	var f_1 gopurs_runtime.Value = f_1_loop
	_ = f_1
	var v_2 gopurs_runtime.Value = v_2_loop
	_ = v_2
	return gopurs_runtime.Apply2(dictFunctor_0.V0, f_1, v_2)
}

func Call_Control_Monad_List_Trans_takeWhile(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
takeWhile:
	for {
		if false {
			continue takeWhile
		}
		var dictApplicative_0 *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
		_ = dictApplicative_0
		// TAST (Let): __local_var_1_0 shape=App(Other) bindingType=Any
		__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictApplicative_0.V0, gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
		_ = __local_var_1_0
		return gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "map"), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t5 gopurs_runtime.Value
				{
					if v_4.Type == 9 && v_4.IntVal == 1320412129 {
						var __t2 gopurs_runtime.Value
						{
							if (gopurs_runtime.Apply(f_2, (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0).IntVal) != (0) {
								// TAST (Let): __local_var_5_1 shape=App(Var) bindingType=Any
								__local_var_5_1 := gopurs_runtime.Apply(Call_Control_Monad_List_Trans_takeWhile(dictApplicative_0), f_2)
								_ = __local_var_5_1
								__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Apply(__local_var_5_1, gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1))
								}))}))}
								goto end_branch_2
							} else {

							}
						}
						{
							__t2 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}
						}
					end_branch_2:
						__t5 = __t2
						goto end_branch_5
					} else {

					}
				}
				{
					if v_4.Type == 9 && v_4.IntVal == 813447293 {
						// TAST (Let): __local_var_5_3 shape=Other bindingType=Any
						__local_var_5_3 := (*Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0
						_ = __local_var_5_3
						// TAST (Let): __local_var_6_4 shape=App(Var) bindingType=Any
						__local_var_6_4 := gopurs_runtime.Apply(Call_Control_Monad_List_Trans_takeWhile(dictApplicative_0), f_2)
						_ = __local_var_6_4
						__t5 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Apply(__local_var_6_4, gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_5_3))
						}))}))}
						goto end_branch_5
					} else {

					}
				}
				{
					if v_4.Type == 9 && v_4.IntVal == 489128924 {
						__t5 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}
						goto end_branch_5
					} else {

					}
				}
				{
					__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
				}
			end_branch_5:
				return __t5
			}), v_3)
		})
	}
}

func Call_Control_Monad_List_Trans_scanl(dictMonad_0_loop *Constructor_Control_Monad_Monad[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, b_2_loop gopurs_runtime.Value, l_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonad_0 *Constructor_Control_Monad_Monad[gopurs_runtime.Value] = dictMonad_0_loop
	_ = dictMonad_0
	var f_1 gopurs_runtime.Value = f_1_loop
	_ = f_1
	var b_2 gopurs_runtime.Value = b_2_loop
	_ = b_2
	var l_3 gopurs_runtime.Value = l_3_loop
	_ = l_3
	return Call_Control_Monad_List_Trans_unfold(dictMonad_0, gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_5_0 shape=Other bindingType=Any
		__local_var_5_0 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0
		_ = __local_var_5_0
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictMonad_0.V1, gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t1 gopurs_runtime.Value
			{
				if v1_6.Type == 9 && v1_6.IntVal == 1320412129 {
					__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer((&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer((&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(f_1, __local_var_5_0, (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v1_6.UnsafePtr).V0), gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v1_6.UnsafePtr).V1)}))}, __local_var_5_0}))}}))}
					goto end_branch_1
				} else {

				}
			}
			{
				if v1_6.Type == 9 && v1_6.IntVal == 813447293 {
					__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer((&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer((&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, __local_var_5_0, gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v1_6.UnsafePtr).V0)}))}, __local_var_5_0}))}}))}
					goto end_branch_1
				} else {

				}
			}
			{
				if v1_6.Type == 9 && v1_6.IntVal == 489128924 {
					__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}
					goto end_branch_1
				} else {

				}
			}
			{
				__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Control_Monad_List_Trans_2549197956_3094389156(func() *Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value], gopurs_runtime.Value]] {
					panic("Failed pattern match")
				}()))}
			}
		end_branch_1:
			return __t1
		}), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1)
	}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
		_v := struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
		}{b_2, l_3}
		return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
	}()))})
}

func Call_Control_Monad_List_Trans_prepend_prime_(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value], h_1_loop gopurs_runtime.Value, t_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictApplicative_0 *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
	_ = dictApplicative_0
	var h_1 gopurs_runtime.Value = h_1_loop
	_ = h_1
	var t_2 gopurs_runtime.Value = t_2_loop
	_ = t_2
	return gopurs_runtime.Apply(dictApplicative_0.V1, gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, h_1, t_2}))})
}

func Call_Control_Monad_List_Trans_prepend(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value], h_1_loop gopurs_runtime.Value, t_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictApplicative_0 *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
	_ = dictApplicative_0
	var h_1 gopurs_runtime.Value = h_1_loop
	_ = h_1
	var t_2 gopurs_runtime.Value = t_2_loop
	_ = t_2
	return gopurs_runtime.Apply(dictApplicative_0.V1, gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, h_1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return t_2
	}))}))})
}

func Call_Control_Monad_List_Trans_nil(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
	_ = dictApplicative_0
	return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)})
}

func Call_Control_Monad_List_Trans_singleton(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictApplicative_0 *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
	_ = dictApplicative_0
	// TAST (Let): nil1_1_0 shape=App(Var) bindingType=(TypeApp (TypeVar f$scope133) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a$scope134), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f$scope133), (TypeVar a$scope134)])])])
	nil1_1_0 := Call_Control_Monad_List_Trans_nil(gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(dictApplicative_0)})
	_ = nil1_1_0
	return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(dictApplicative_0.V1, gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_2, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return nil1_1_0
		}))}))})
	})
}

func Call_Control_Monad_List_Trans_take(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
take:
	for {
		if false {
			continue take
		}
		var dictApplicative_0 *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
		_ = dictApplicative_0
		// TAST (Let): nil1_1_0 shape=App(Var) bindingType=(TypeApp (TypeVar f$scope137) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a$scope138), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f$scope137), (TypeVar a$scope138)])])])
		nil1_1_0 := Call_Control_Monad_List_Trans_nil(gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(dictApplicative_0)})
		_ = nil1_1_0
		// TAST (Let): __local_var_2_1 shape=App(Other) bindingType=Any
		__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictApplicative_0.V0, gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
		_ = __local_var_2_1
		return gopurs_runtime.Func2(func(v_3 gopurs_runtime.Value, v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t7 gopurs_runtime.Value
			{
				if (v_3.IntVal) == (int64(0)) {
					__t7 = nil1_1_0
					goto end_branch_7
				} else {

				}
			}
			{
				__t7 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "map"), gopurs_runtime.Func(func(v2_5 gopurs_runtime.Value) gopurs_runtime.Value {
					var __t6 gopurs_runtime.Value
					{
						if v2_5.Type == 9 && v2_5.IntVal == 1320412129 {
							// TAST (Let): __local_var_6_2 shape=Other bindingType=Any
							__local_var_6_2 := (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v2_5.UnsafePtr).V1
							_ = __local_var_6_2
							// TAST (Let): __local_var_7_3 shape=App(Var) bindingType=Any
							__local_var_7_3 := gopurs_runtime.Apply(Call_Control_Monad_List_Trans_take(dictApplicative_0), gopurs_runtime.Int(gopurs_runtime.IntSub(v_3.IntVal, int64(1))))
							_ = __local_var_7_3
							__t6 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v2_5.UnsafePtr).V0, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Apply(__local_var_7_3, gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_6_2))
							}))}))}
							goto end_branch_6
						} else {

						}
					}
					{
						if v2_5.Type == 9 && v2_5.IntVal == 813447293 {
							// TAST (Let): __local_var_6_4 shape=Other bindingType=Any
							__local_var_6_4 := (*Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v2_5.UnsafePtr).V0
							_ = __local_var_6_4
							// TAST (Let): __local_var_7_5 shape=App(Var) bindingType=Any
							__local_var_7_5 := gopurs_runtime.Apply(Call_Control_Monad_List_Trans_take(dictApplicative_0), gopurs_runtime.Int(v_3.IntVal))
							_ = __local_var_7_5
							__t6 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Apply(__local_var_7_5, gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_6_4))
							}))}))}
							goto end_branch_6
						} else {

						}
					}
					{
						if v2_5.Type == 9 && v2_5.IntVal == 489128924 {
							__t6 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}
							goto end_branch_6
						} else {

						}
					}
					{
						__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
					}
				end_branch_6:
					return __t6
				}), v1_4)
			}
		end_branch_7:
			return __t7
		})
	}
}

func Call_Control_Monad_List_Trans_zipWith_prime_(dictMonad_0_loop *Constructor_Control_Monad_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
zipWith_prime_:
	for {
		if false {
			continue zipWith_prime_
		}
		var dictMonad_0 *Constructor_Control_Monad_Monad[gopurs_runtime.Value] = dictMonad_0_loop
		_ = dictMonad_0
		// TAST (Let): Applicative0_1_0 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar f$scope153)])
		Applicative0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonad_0.V0, gopurs_runtime.Value{}))
		_ = Applicative0_1_0
		// TAST (Let): nil1_2_1 shape=App(Other) bindingType=(TypeApp (TypeVar f$scope153) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar c$scope156), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f$scope153), (TypeVar c$scope156)])])])
		nil1_2_1 := gopurs_runtime.Apply(Applicative0_1_0.V1, gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)})
		_ = nil1_2_1
		// TAST (Let): Functor0_3_2 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar f$scope153)])
		Functor0_3_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictMonad_0.V1, gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
		_ = Functor0_3_2
		// TAST (Let): prepend_prime_1_4_3 shape=App(Var) bindingType=(Func [(TypeVar c$scope156), (ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f$scope153) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar c$scope156), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f$scope153), (TypeVar c$scope156)])])])])] (TypeApp (TypeVar f$scope153) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar c$scope156), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f$scope153), (TypeVar c$scope156)])])]))
		prepend_prime_1_4_3 := gopurs_runtime.Apply(Get_Control_Monad_List_Trans_prepend_prime_(), gopurs_runtime.Apply(dictMonad_0.V0, gopurs_runtime.Value{}))
		_ = prepend_prime_1_4_3
		// TAST (Let): Bind1_5_4 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar f$scope153)])
		Bind1_5_4 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonad_0.V1, gopurs_runtime.Value{}))
		_ = Bind1_5_4
		// TAST (Let): uncons1_6_5 shape=App(Var) bindingType=(Func [(TypeApp (TypeVar f$scope153) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a$scope154), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f$scope153), (TypeVar a$scope154)])])])] (TypeApp (TypeVar f$scope153) [(ADT ["Data","Maybe","Maybe"] [(ADT ["Data","Tuple","Tuple"] [(TypeVar a$scope154), (TypeApp (TypeVar f$scope153) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a$scope154), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f$scope153), (TypeVar a$scope154)])])])])])]))
		uncons1_6_5 := Call_Control_Monad_List_Trans_uncons(dictMonad_0)
		_ = uncons1_6_5
		// TAST (Let): uncons2_7_6 shape=App(Var) bindingType=(Func [(TypeApp (TypeVar f$scope153) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar b$scope155), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f$scope153), (TypeVar b$scope155)])])])] (TypeApp (TypeVar f$scope153) [(ADT ["Data","Maybe","Maybe"] [(ADT ["Data","Tuple","Tuple"] [(TypeVar b$scope155), (TypeApp (TypeVar f$scope153) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar b$scope155), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f$scope153), (TypeVar b$scope155)])])])])])]))
		uncons2_7_6 := Call_Control_Monad_List_Trans_uncons(dictMonad_0)
		_ = uncons2_7_6
		return gopurs_runtime.Func3(func(f_8 gopurs_runtime.Value, fa_9 gopurs_runtime.Value, fb_10 gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_List_Trans_wrapEffect(Functor0_3_2, gopurs_runtime.Apply2(Bind1_5_4.V1, gopurs_runtime.Apply(uncons1_6_5, fa_9), gopurs_runtime.Func(func(ua_11 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(Bind1_5_4.V1, gopurs_runtime.Apply(uncons2_7_6, fb_10), gopurs_runtime.Func(func(ub_12 gopurs_runtime.Value) gopurs_runtime.Value {
					var __t10 gopurs_runtime.Value
					{
						if ub_12.Type == 9 && ub_12.IntVal == 930809136 && ub_12.UnsafePtr == nil {
							__t10 = gopurs_runtime.Apply(Applicative0_1_0.V1, nil1_2_1)
							goto end_branch_10
						} else {

						}
					}
					{
						if ua_11.Type == 9 && ua_11.IntVal == 930809136 && ua_11.UnsafePtr == nil {
							__t10 = gopurs_runtime.Apply(Applicative0_1_0.V1, nil1_2_1)
							goto end_branch_10
						} else {

						}
					}
					{
						if (ua_11.Type == 9 && ua_11.IntVal == 930809136 && ua_11.UnsafePtr != nil) && (ub_12.Type == 9 && ub_12.IntVal == 930809136 && ub_12.UnsafePtr != nil) {
							// TAST (Let): __local_var_13_7 shape=Other bindingType=Any
							__local_var_13_7 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(ua_11.UnsafePtr).V0.UnsafePtr).V1
							_ = __local_var_13_7
							// TAST (Let): __local_var_14_8 shape=Other bindingType=Any
							__local_var_14_8 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(ub_12.UnsafePtr).V0.UnsafePtr).V1
							_ = __local_var_14_8
							// TAST (Let): __local_var_15_9 shape=App(Var) bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f$scope153) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar c$scope156), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f$scope153), (TypeVar c$scope156)])])])])
							__local_var_15_9 := gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v2_15 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Apply4(Get_Control_Monad_List_Trans_zipWith_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(dictMonad_0)}, f_8, __local_var_13_7, __local_var_14_8)
							}))
							_ = __local_var_15_9
							__t10 = gopurs_runtime.Apply2(Functor0_3_2.V0, gopurs_runtime.Func(func(a_16 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Apply2(prepend_prime_1_4_3, a_16, __local_var_15_9)
							}), gopurs_runtime.Apply2(f_8, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(ua_11.UnsafePtr).V0.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(ub_12.UnsafePtr).V0.UnsafePtr).V0))
							goto end_branch_10
						} else {

						}
					}
					{
						__t10 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
					}
				end_branch_10:
					return __t10
				}))
			})))
		})
	}
}

func Call_Control_Monad_List_Trans_zipWith(dictMonad_0_loop *Constructor_Control_Monad_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictMonad_0 *Constructor_Control_Monad_Monad[gopurs_runtime.Value] = dictMonad_0_loop
	_ = dictMonad_0
	// TAST (Let): pure2_1_0 shape=App(Var) bindingType=(Func [(TypeVar c$scope169)] (TypeApp (TypeVar f$scope166) [(TypeVar c$scope169)]))
	pure2_1_0 := Call_Control_Applicative_pure(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonad_0.V0, gopurs_runtime.Value{})))
	_ = pure2_1_0
	// TAST (Let): zipWith_prime_1_2_1 shape=App(Var) bindingType=(Func [(Func [(TypeVar a$scope167), (TypeVar b$scope168)] (TypeApp (TypeVar f$scope166) [(TypeVar c$scope169)])), (TypeApp (TypeVar f$scope166) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a$scope167), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f$scope166), (TypeVar a$scope167)])])]), (TypeApp (TypeVar f$scope166) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar b$scope168), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f$scope166), (TypeVar b$scope168)])])])] (TypeApp (TypeVar f$scope166) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar c$scope169), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f$scope166), (TypeVar c$scope169)])])]))
	zipWith_prime_1_2_1 := gopurs_runtime.Apply(Get_Control_Monad_List_Trans_zipWith_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(dictMonad_0)})
	_ = zipWith_prime_1_2_1
	return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(zipWith_prime_1_2_1, gopurs_runtime.Func2(func(a_4 gopurs_runtime.Value, b_5 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(pure2_1_0, gopurs_runtime.Apply2(f_3, a_4, b_5))
		}))
	})
}

func Call_Control_Monad_List_Trans_mapMaybe(dictFunctor_0_loop *Constructor_Data_Functor_Functor[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
mapMaybe:
	for {
		if false {
			continue mapMaybe
		}
		var dictFunctor_0 *Constructor_Data_Functor_Functor[gopurs_runtime.Value] = dictFunctor_0_loop
		_ = dictFunctor_0
		var f_1 gopurs_runtime.Value = f_1_loop
		_ = f_1
		var v_2 gopurs_runtime.Value = v_2_loop
		_ = v_2
		return gopurs_runtime.Apply2(dictFunctor_0.V0, gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t6 gopurs_runtime.Value
			{
				if v_3.Type == 9 && v_3.IntVal == 1320412129 {
					// TAST (Let): __local_var_4_0 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar b$scope199)])
					__local_var_4_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(f_1, (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0))
					_ = __local_var_4_0
					var __t3 gopurs_runtime.Value
					{
						if __local_var_4_0 != nil {
							// TAST (Let): __local_var_5_2 shape=App(Var) bindingType=Any
							__local_var_5_2 := gopurs_runtime.Apply2(Get_Control_Monad_List_Trans_mapMaybe(), gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(dictFunctor_0)}, f_1)
							_ = __local_var_5_2
							__t3 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, (__local_var_4_0).V0, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Apply(__local_var_5_2, gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1))
							}))}))}
							goto end_branch_3
						} else {

						}
					}
					{
						// TAST (Let): __local_var_5_1 shape=App(Var) bindingType=Any
						__local_var_5_1 := gopurs_runtime.Apply2(Get_Control_Monad_List_Trans_mapMaybe(), gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(dictFunctor_0)}, f_1)
						_ = __local_var_5_1
						__t3 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Apply(__local_var_5_1, gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1))
						}))}))}
					}
				end_branch_3:
					__t6 = __t3
					goto end_branch_6
				} else {

				}
			}
			{
				if v_3.Type == 9 && v_3.IntVal == 813447293 {
					// TAST (Let): __local_var_4_4 shape=Other bindingType=Any
					__local_var_4_4 := (*Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0
					_ = __local_var_4_4
					// TAST (Let): __local_var_5_5 shape=App(Var) bindingType=Any
					__local_var_5_5 := gopurs_runtime.Apply2(Get_Control_Monad_List_Trans_mapMaybe(), gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(dictFunctor_0)}, f_1)
					_ = __local_var_5_5
					__t6 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Apply(__local_var_5_5, gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_4_4))
					}))}))}
					goto end_branch_6
				} else {

				}
			}
			{
				if v_3.Type == 9 && v_3.IntVal == 489128924 {
					__t6 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}
					goto end_branch_6
				} else {

				}
			}
			{
				__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_6:
			return __t6
		}), v_2)
	}
}

func Call_Control_Monad_List_Trans_iterate(dictMonad_0_loop *Constructor_Control_Monad_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictMonad_0 *Constructor_Control_Monad_Monad[gopurs_runtime.Value] = dictMonad_0_loop
	_ = dictMonad_0
	// TAST (Let): pure2_1_0 shape=App(Var) bindingType=(Func [(ADT ["Data","Maybe","Maybe"] [(ADT ["Data","Tuple","Tuple"] [(TypeVar a$scope217), (TypeVar a$scope217)])])] (TypeApp (TypeVar f$scope216) [(ADT ["Data","Maybe","Maybe"] [(ADT ["Data","Tuple","Tuple"] [(TypeVar a$scope217), (TypeVar a$scope217)])])]))
	pure2_1_0 := Call_Control_Applicative_pure(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonad_0.V0, gopurs_runtime.Value{})))
	_ = pure2_1_0
	return gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, a_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return Call_Control_Monad_List_Trans_unfold(dictMonad_0, gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(pure2_1_0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Control_Monad_List_Trans_4010058633_3094389156(Rebox_Control_Monad_List_Trans_3094389156_4010058633(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer((&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_2, x_4), x_4}))}, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))))})
		}), a_3)
	})
}

func Call_Control_Monad_List_Trans_repeat(dictMonad_0_loop *Constructor_Control_Monad_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictMonad_0 *Constructor_Control_Monad_Monad[gopurs_runtime.Value] = dictMonad_0_loop
	_ = dictMonad_0
	return gopurs_runtime.Apply(Call_Control_Monad_List_Trans_iterate(dictMonad_0), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return x_1
	}))
}

func Call_Control_Monad_List_Trans_head(dictMonad_0_loop *Constructor_Control_Monad_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictMonad_0 *Constructor_Control_Monad_Monad[gopurs_runtime.Value] = dictMonad_0_loop
	_ = dictMonad_0
	// TAST (Let): uncons1_1_0 shape=App(Var) bindingType=(Func [(TypeApp (TypeVar f$scope228) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a$scope229), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f$scope228), (TypeVar a$scope229)])])])] (TypeApp (TypeVar f$scope228) [(ADT ["Data","Maybe","Maybe"] [(ADT ["Data","Tuple","Tuple"] [(TypeVar a$scope229), (TypeApp (TypeVar f$scope228) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a$scope229), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f$scope228), (TypeVar a$scope229)])])])])])]))
	uncons1_1_0 := Call_Control_Monad_List_Trans_uncons(dictMonad_0)
	_ = uncons1_1_0
	return gopurs_runtime.Func(func(l_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictMonad_0.V1, gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t1 struct {
				V0 gopurs_runtime.Value
				V1 bool
			}
			{
				if v1_3.Type == 9 && v1_3.IntVal == 930809136 && v1_3.UnsafePtr != nil {
					__t1 = struct {
						V0 gopurs_runtime.Value
						V1 bool
					}{Call_Data_Tuple_fst(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v1_3.UnsafePtr).V0)), true}
					goto end_branch_1
				} else {

				}
			}
			{
				__t1 = struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Value{}, false}
			}
		end_branch_1:
			return func() gopurs_runtime.Value {
				_v := __t1
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		}), gopurs_runtime.Apply(uncons1_1_0, l_2))
	})
}

func Call_Control_Monad_List_Trans_functorListT(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
functorListT:
	for {
		if false {
			continue functorListT
		}
		var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
		_ = dictFunctor_0
		return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t4 gopurs_runtime.Value
				{
					if v_3.Type == 9 && v_3.IntVal == 1320412129 {
						// TAST (Let): __local_var_4_0 shape=Other bindingType=Any
						__local_var_4_0 := (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1
						_ = __local_var_4_0
						// TAST (Let): __local_var_5_1 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f$scope231) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a$scope235), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f$scope231), (TypeVar a$scope235)])])])] (TypeApp (TypeVar f$scope231) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar b$scope236), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f$scope231), (TypeVar b$scope236)])])]))
						__local_var_5_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_functorListT(dictFunctor_0), "map"), f_1)
						_ = __local_var_5_1
						__t4 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_1, (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Apply(__local_var_5_1, gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_4_0))
						}))}))}
						goto end_branch_4
					} else {

					}
				}
				{
					if v_3.Type == 9 && v_3.IntVal == 813447293 {
						// TAST (Let): __local_var_4_2 shape=Other bindingType=Any
						__local_var_4_2 := (*Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0
						_ = __local_var_4_2
						// TAST (Let): __local_var_5_3 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f$scope231) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a$scope235), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f$scope231), (TypeVar a$scope235)])])])] (TypeApp (TypeVar f$scope231) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar b$scope236), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f$scope231), (TypeVar b$scope236)])])]))
						__local_var_5_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_functorListT(dictFunctor_0), "map"), f_1)
						_ = __local_var_5_3
						__t4 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Apply(__local_var_5_3, gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_4_2))
						}))}))}
						goto end_branch_4
					} else {

					}
				}
				{
					if v_3.Type == 9 && v_3.IntVal == 489128924 {
						__t4 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}
						goto end_branch_4
					} else {

					}
				}
				{
					__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
				}
			end_branch_4:
				return __t4
			}), v_2)
		})}))}
	}
}

func Call_Control_Monad_List_Trans_fromEffect(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictApplicative_0 *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
	_ = dictApplicative_0
	// TAST (Let): nil1_1_0 shape=App(Other) bindingType=(TypeApp (TypeVar f$scope245) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a$scope246), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f$scope245), (TypeVar a$scope246)])])])
	nil1_1_0 := gopurs_runtime.Apply(dictApplicative_0.V1, gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)})
	_ = nil1_1_0
	return gopurs_runtime.Func(func(fa_2 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_3_1 shape=App(Var) bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f$scope245) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a$scope246), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f$scope245), (TypeVar a$scope246)])])])])
		__local_var_3_1 := gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return nil1_1_0
		}))
		_ = __local_var_3_1
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictApplicative_0.V0, gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_4, __local_var_3_1}))}
		}), fa_2)
	})
}

func Call_Control_Monad_List_Trans_foldlRec_prime_(dictMonadRec_0_loop *Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictMonadRec_0 *Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value] = dictMonadRec_0_loop
	_ = dictMonadRec_0
	// TAST (Let): Monad0_1_0 shape=App(Other) bindingType=Any
	Monad0_1_0 := gopurs_runtime.Apply(dictMonadRec_0.V0, gopurs_runtime.Value{})
	_ = Monad0_1_0
	// TAST (Let): __local_var_2_1 shape=App(Other) bindingType=Any
	__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{})
	_ = __local_var_2_1
	// TAST (Let): Bind1_3_2 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar f$scope260)])
	Bind1_3_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
	_ = Bind1_3_2
	// TAST (Let): uncons1_4_3 shape=App(Var) bindingType=(Func [(TypeApp (TypeVar f$scope260) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a$scope261), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f$scope260), (TypeVar a$scope261)])])])] (TypeApp (TypeVar f$scope260) [(ADT ["Data","Maybe","Maybe"] [(ADT ["Data","Tuple","Tuple"] [(TypeVar a$scope261), (TypeApp (TypeVar f$scope260) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a$scope261), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f$scope260), (TypeVar a$scope261)])])])])])]))
	uncons1_4_3 := Call_Control_Monad_List_Trans_uncons(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonadRec_0.V0, gopurs_runtime.Value{})))
	_ = uncons1_4_3
	return gopurs_runtime.Func3(func(f_5 gopurs_runtime.Value, a_6 gopurs_runtime.Value, b_7 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(dictMonadRec_0.V1, gopurs_runtime.Func(func(o_8 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_9_4 shape=Other bindingType=(TypeVar a$scope8)
			__local_var_9_4 := gopurs_runtime.RecordGet(o_8, "a")
			_ = __local_var_9_4
			return gopurs_runtime.Apply2(Bind1_3_2.V1, gopurs_runtime.Apply(uncons1_4_3, gopurs_runtime.RecordGet(o_8, "b")), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t6 gopurs_runtime.Value
				{
					if v_10.Type == 9 && v_10.IntVal == 930809136 && v_10.UnsafePtr == nil {
						__t6 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(Rebox_Control_Monad_List_Trans_1239976062_3603546092((&Constructor_Control_Monad_Rec_Class_Done[struct {
							a gopurs_runtime.Value
							b gopurs_runtime.Value
							/*gopurs:record-v1:2:0061,0062*/
						}, gopurs_runtime.Value]{1, __local_var_9_4})))})
						goto end_branch_6
					} else {

					}
				}
				{
					if v_10.Type == 9 && v_10.IntVal == 930809136 && v_10.UnsafePtr != nil {
						// TAST (Let): __local_var_11_5 shape=Other bindingType=Any
						__local_var_11_5 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v_10.UnsafePtr).V0.UnsafePtr).V1
						_ = __local_var_11_5
						__t6 = gopurs_runtime.Apply2(Bind1_3_2.V1, gopurs_runtime.Apply2(f_5, __local_var_9_4, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v_10.UnsafePtr).V0.UnsafePtr).V0), gopurs_runtime.Func(func(b_prime__12 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(Rebox_Control_Monad_List_Trans_784458114_4008603408((&Constructor_Control_Monad_Rec_Class_Loop[struct {
								a gopurs_runtime.Value
								b gopurs_runtime.Value
								/*gopurs:record-v1:2:0061,0062*/
							}, gopurs_runtime.Value]{1, struct {
								a gopurs_runtime.Value
								b gopurs_runtime.Value
								/*gopurs:record-v1:2:0061,0062*/
							}{b_prime__12, __local_var_11_5}})))})
						}))
						goto end_branch_6
					} else {

					}
				}
				{
					__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
				}
			end_branch_6:
				return __t6
			}))
		}), func() gopurs_runtime.Value {
			orig := struct {
				a gopurs_runtime.Value
				b gopurs_runtime.Value
				/*gopurs:record-v1:2:0061,0062*/
			}{a_6, b_7}
			_ = orig
			return gopurs_runtime.RecordDict2("a", "b", orig.a, orig.b)
		}())
	})
}

func Call_Control_Monad_List_Trans_runListTRec(dictMonadRec_0_loop *Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictMonadRec_0 *Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value] = dictMonadRec_0_loop
	_ = dictMonadRec_0
	return gopurs_runtime.Apply3(Get_Control_Monad_List_Trans_foldlRec_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 3709389635, UnsafePtr: unsafe.Pointer(dictMonadRec_0)}, gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictMonadRec_0.V0, gopurs_runtime.Value{}), "Applicative0"), gopurs_runtime.Value{}), "pure"), Get_Data_Unit_unit())
	}), Get_Data_Unit_unit())
}

func Call_Control_Monad_List_Trans_foldlRec(dictMonadRec_0_loop *Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictMonadRec_0 *Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value] = dictMonadRec_0_loop
	_ = dictMonadRec_0
	// TAST (Let): Monad0_1_0 shape=App(Other) bindingType=Any
	Monad0_1_0 := gopurs_runtime.Apply(dictMonadRec_0.V0, gopurs_runtime.Value{})
	_ = Monad0_1_0
	// TAST (Let): __local_var_2_1 shape=App(Other) bindingType=Any
	__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{})
	_ = __local_var_2_1
	// TAST (Let): uncons1_3_2 shape=App(Var) bindingType=(Func [(TypeApp (TypeVar f$scope274) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a$scope275), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f$scope274), (TypeVar a$scope275)])])])] (TypeApp (TypeVar f$scope274) [(ADT ["Data","Maybe","Maybe"] [(ADT ["Data","Tuple","Tuple"] [(TypeVar a$scope275), (TypeApp (TypeVar f$scope274) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a$scope275), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f$scope274), (TypeVar a$scope275)])])])])])]))
	uncons1_3_2 := Call_Control_Monad_List_Trans_uncons(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonadRec_0.V0, gopurs_runtime.Value{})))
	_ = uncons1_3_2
	return gopurs_runtime.Func3(func(f_4 gopurs_runtime.Value, a_5 gopurs_runtime.Value, b_6 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(dictMonadRec_0.V1, gopurs_runtime.Func(func(o_7 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_8_3 shape=Other bindingType=(TypeVar a$scope8)
			__local_var_8_3 := gopurs_runtime.RecordGet(o_7, "a")
			_ = __local_var_8_3
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply(uncons1_3_2, gopurs_runtime.RecordGet(o_7, "b")), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t4 gopurs_runtime.Value
				{
					if v_9.Type == 9 && v_9.IntVal == 930809136 && v_9.UnsafePtr == nil {
						__t4 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(Rebox_Control_Monad_List_Trans_1239976062_3603546092((&Constructor_Control_Monad_Rec_Class_Done[struct {
							a gopurs_runtime.Value
							b gopurs_runtime.Value
							/*gopurs:record-v1:2:0061,0062*/
						}, gopurs_runtime.Value]{1, __local_var_8_3})))})
						goto end_branch_4
					} else {

					}
				}
				{
					if v_9.Type == 9 && v_9.IntVal == 930809136 && v_9.UnsafePtr != nil {
						__t4 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(Rebox_Control_Monad_List_Trans_784458114_4008603408((&Constructor_Control_Monad_Rec_Class_Loop[struct {
							a gopurs_runtime.Value
							b gopurs_runtime.Value
							/*gopurs:record-v1:2:0061,0062*/
						}, gopurs_runtime.Value]{1, struct {
							a gopurs_runtime.Value
							b gopurs_runtime.Value
							/*gopurs:record-v1:2:0061,0062*/
						}{gopurs_runtime.Apply2(f_4, __local_var_8_3, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v_9.UnsafePtr).V0.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v_9.UnsafePtr).V0.UnsafePtr).V1}})))})
						goto end_branch_4
					} else {

					}
				}
				{
					__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
				}
			end_branch_4:
				return __t4
			}))
		}), func() gopurs_runtime.Value {
			orig := struct {
				a gopurs_runtime.Value
				b gopurs_runtime.Value
				/*gopurs:record-v1:2:0061,0062*/
			}{a_5, b_6}
			_ = orig
			return gopurs_runtime.RecordDict2("a", "b", orig.a, orig.b)
		}())
	})
}

func Call_Control_Monad_List_Trans_foldl_prime_(dictMonad_0_loop *Constructor_Control_Monad_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictMonad_0 *Constructor_Control_Monad_Monad[gopurs_runtime.Value] = dictMonad_0_loop
	_ = dictMonad_0
	// TAST (Let): Bind1_1_0 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar f$scope284)])
	Bind1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonad_0.V1, gopurs_runtime.Value{}))
	_ = Bind1_1_0
	// TAST (Let): uncons1_2_1 shape=App(Var) bindingType=(Func [(TypeApp (TypeVar f$scope284) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a$scope285), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f$scope284), (TypeVar a$scope285)])])])] (TypeApp (TypeVar f$scope284) [(ADT ["Data","Maybe","Maybe"] [(ADT ["Data","Tuple","Tuple"] [(TypeVar a$scope285), (TypeApp (TypeVar f$scope284) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a$scope285), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f$scope284), (TypeVar a$scope285)])])])])])]))
	uncons1_2_1 := Call_Control_Monad_List_Trans_uncons(dictMonad_0)
	_ = uncons1_2_1
	return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
		var loop_4_2_0 gopurs_runtime.Value
		_ = loop_4_2_0
		var loop_4_2_0_cell *gopurs_runtime.Value
		_ = loop_4_2_0_cell
		// FALLBACK TCO: isLoop=false len=1
		loop_4_2_0 = gopurs_runtime.Func2(func(b_5 gopurs_runtime.Value, l_6 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(Bind1_1_0.V1, gopurs_runtime.Apply(uncons1_2_1, l_6), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t3 gopurs_runtime.Value
				{
					if v_7.Type == 9 && v_7.IntVal == 930809136 && v_7.UnsafePtr == nil {
						__t3 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictMonad_0.V0, gopurs_runtime.Value{}), "pure"), b_5)
						goto end_branch_3
					} else {

					}
				}
				{
					if v_7.Type == 9 && v_7.IntVal == 930809136 && v_7.UnsafePtr != nil {
						__t3 = gopurs_runtime.Apply2(Bind1_1_0.V1, gopurs_runtime.Apply2(f_3, b_5, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v_7.UnsafePtr).V0.UnsafePtr).V0), gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Apply2((*loop_4_2_0_cell), a_8, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v_7.UnsafePtr).V0.UnsafePtr).V1)
						}))
						goto end_branch_3
					} else {

					}
				}
				{
					__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
				}
			end_branch_3:
				return __t3
			}))
		})
		loop_4_2_0_cell = &loop_4_2_0
		return loop_4_2_0
	})
}

func Call_Control_Monad_List_Trans_runListT(dictMonad_0_loop *Constructor_Control_Monad_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictMonad_0 *Constructor_Control_Monad_Monad[gopurs_runtime.Value] = dictMonad_0_loop
	_ = dictMonad_0
	return gopurs_runtime.Apply3(Get_Control_Monad_List_Trans_foldl_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(dictMonad_0)}, gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictMonad_0.V0, gopurs_runtime.Value{}), "pure"), Get_Data_Unit_unit())
	}), Get_Data_Unit_unit())
}

func Call_Control_Monad_List_Trans_foldl(dictMonad_0_loop *Constructor_Control_Monad_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictMonad_0 *Constructor_Control_Monad_Monad[gopurs_runtime.Value] = dictMonad_0_loop
	_ = dictMonad_0
	// TAST (Let): uncons1_1_0 shape=App(Var) bindingType=(Func [(TypeApp (TypeVar f$scope294) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a$scope295), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f$scope294), (TypeVar a$scope295)])])])] (TypeApp (TypeVar f$scope294) [(ADT ["Data","Maybe","Maybe"] [(ADT ["Data","Tuple","Tuple"] [(TypeVar a$scope295), (TypeApp (TypeVar f$scope294) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a$scope295), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f$scope294), (TypeVar a$scope295)])])])])])]))
	uncons1_1_0 := Call_Control_Monad_List_Trans_uncons(dictMonad_0)
	_ = uncons1_1_0
	return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
		var loop_3_1_1 gopurs_runtime.Value
		_ = loop_3_1_1
		var loop_3_1_1_cell *gopurs_runtime.Value
		_ = loop_3_1_1_cell
		// FALLBACK TCO: isLoop=false len=1
		loop_3_1_1 = gopurs_runtime.Func2(func(b_4 gopurs_runtime.Value, l_5 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictMonad_0.V1, gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply(uncons1_1_0, l_5), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t2 gopurs_runtime.Value
				{
					if v_6.Type == 9 && v_6.IntVal == 930809136 && v_6.UnsafePtr == nil {
						__t2 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictMonad_0.V0, gopurs_runtime.Value{}), "pure"), b_4)
						goto end_branch_2
					} else {

					}
				}
				{
					if v_6.Type == 9 && v_6.IntVal == 930809136 && v_6.UnsafePtr != nil {
						__t2 = gopurs_runtime.Apply2((*loop_3_1_1_cell), gopurs_runtime.Apply2(f_2, b_4, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v_6.UnsafePtr).V0.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v_6.UnsafePtr).V0.UnsafePtr).V1)
						goto end_branch_2
					} else {

					}
				}
				{
					__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
				}
			end_branch_2:
				return __t2
			}))
		})
		loop_3_1_1_cell = &loop_3_1_1
		return loop_3_1_1
	})
}

func Call_Control_Monad_List_Trans_filter(dictFunctor_0_loop *Constructor_Data_Functor_Functor[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
filter:
	for {
		if false {
			continue filter
		}
		var dictFunctor_0 *Constructor_Data_Functor_Functor[gopurs_runtime.Value] = dictFunctor_0_loop
		_ = dictFunctor_0
		var f_1 gopurs_runtime.Value = f_1_loop
		_ = f_1
		var v_2 gopurs_runtime.Value = v_2_loop
		_ = v_2
		return gopurs_runtime.Apply2(dictFunctor_0.V0, gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t6 gopurs_runtime.Value
			{
				if v_3.Type == 9 && v_3.IntVal == 1320412129 {
					// TAST (Let): __local_var_4_0 shape=Other bindingType=Any
					__local_var_4_0 := (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1
					_ = __local_var_4_0
					// TAST (Let): __local_var_5_1 shape=App(Var) bindingType=Any
					__local_var_5_1 := gopurs_runtime.Apply2(Get_Control_Monad_List_Trans_filter(), gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(dictFunctor_0)}, f_1)
					_ = __local_var_5_1
					// TAST (Let): s_prime__6_2 shape=App(Var) bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f$scope299) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a$scope300), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f$scope299), (TypeVar a$scope300)])])])])
					s_prime__6_2 := gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Apply(__local_var_5_1, gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_4_0))
					}))
					_ = s_prime__6_2
					var __t3 gopurs_runtime.Value
					{
						if (gopurs_runtime.Apply(f_1, (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0).IntVal) != (0) {
							__t3 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0, s_prime__6_2}))}
							goto end_branch_3
						} else {

						}
					}
					{
						__t3 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, s_prime__6_2}))}
					}
				end_branch_3:
					__t6 = __t3
					goto end_branch_6
				} else {

				}
			}
			{
				if v_3.Type == 9 && v_3.IntVal == 813447293 {
					// TAST (Let): __local_var_4_4 shape=Other bindingType=Any
					__local_var_4_4 := (*Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0
					_ = __local_var_4_4
					// TAST (Let): __local_var_5_5 shape=App(Var) bindingType=Any
					__local_var_5_5 := gopurs_runtime.Apply2(Get_Control_Monad_List_Trans_filter(), gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(dictFunctor_0)}, f_1)
					_ = __local_var_5_5
					__t6 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Apply(__local_var_5_5, gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_4_4))
					}))}))}
					goto end_branch_6
				} else {

				}
			}
			{
				if v_3.Type == 9 && v_3.IntVal == 489128924 {
					__t6 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}
					goto end_branch_6
				} else {

				}
			}
			{
				__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_6:
			return __t6
		}), v_2)
	}
}

func Call_Control_Monad_List_Trans_dropWhile(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
dropWhile:
	for {
		if false {
			continue dropWhile
		}
		var dictApplicative_0 *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
		_ = dictApplicative_0
		// TAST (Let): __local_var_1_0 shape=App(Other) bindingType=Any
		__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictApplicative_0.V0, gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
		_ = __local_var_1_0
		return gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "map"), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t5 gopurs_runtime.Value
				{
					if v_4.Type == 9 && v_4.IntVal == 1320412129 {
						var __t2 gopurs_runtime.Value
						{
							if (gopurs_runtime.Apply(f_2, (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0).IntVal) != (0) {
								// TAST (Let): __local_var_5_1 shape=App(Var) bindingType=Any
								__local_var_5_1 := gopurs_runtime.Apply(Call_Control_Monad_List_Trans_dropWhile(dictApplicative_0), f_2)
								_ = __local_var_5_1
								__t2 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Apply(__local_var_5_1, gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1))
								}))}))}
								goto end_branch_2
							} else {

							}
						}
						{
							__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0, (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1}))}
						}
					end_branch_2:
						__t5 = __t2
						goto end_branch_5
					} else {

					}
				}
				{
					if v_4.Type == 9 && v_4.IntVal == 813447293 {
						// TAST (Let): __local_var_5_3 shape=Other bindingType=Any
						__local_var_5_3 := (*Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0
						_ = __local_var_5_3
						// TAST (Let): __local_var_6_4 shape=App(Var) bindingType=Any
						__local_var_6_4 := gopurs_runtime.Apply(Call_Control_Monad_List_Trans_dropWhile(dictApplicative_0), f_2)
						_ = __local_var_6_4
						__t5 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Apply(__local_var_6_4, gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_5_3))
						}))}))}
						goto end_branch_5
					} else {

					}
				}
				{
					if v_4.Type == 9 && v_4.IntVal == 489128924 {
						__t5 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}
						goto end_branch_5
					} else {

					}
				}
				{
					__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
				}
			end_branch_5:
				return __t5
			}), v_3)
		})
	}
}

func Call_Control_Monad_List_Trans_drop(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
drop:
	for {
		if false {
			continue drop
		}
		var dictApplicative_0 *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
		_ = dictApplicative_0
		// TAST (Let): __local_var_1_0 shape=App(Other) bindingType=Any
		__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictApplicative_0.V0, gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
		_ = __local_var_1_0
		return gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t6 gopurs_runtime.Value
			{
				if (v_2.IntVal) == (int64(0)) {
					__t6 = v1_3
					goto end_branch_6
				} else {

				}
			}
			{
				__t6 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "map"), gopurs_runtime.Func(func(v2_4 gopurs_runtime.Value) gopurs_runtime.Value {
					var __t5 gopurs_runtime.Value
					{
						if v2_4.Type == 9 && v2_4.IntVal == 1320412129 {
							// TAST (Let): __local_var_5_1 shape=Other bindingType=Any
							__local_var_5_1 := (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v2_4.UnsafePtr).V1
							_ = __local_var_5_1
							// TAST (Let): __local_var_6_2 shape=App(Var) bindingType=Any
							__local_var_6_2 := gopurs_runtime.Apply(Call_Control_Monad_List_Trans_drop(dictApplicative_0), gopurs_runtime.Int(gopurs_runtime.IntSub(v_2.IntVal, int64(1))))
							_ = __local_var_6_2
							__t5 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Apply(__local_var_6_2, gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_5_1))
							}))}))}
							goto end_branch_5
						} else {

						}
					}
					{
						if v2_4.Type == 9 && v2_4.IntVal == 813447293 {
							// TAST (Let): __local_var_5_3 shape=Other bindingType=Any
							__local_var_5_3 := (*Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v2_4.UnsafePtr).V0
							_ = __local_var_5_3
							// TAST (Let): __local_var_6_4 shape=App(Var) bindingType=Any
							__local_var_6_4 := gopurs_runtime.Apply(Call_Control_Monad_List_Trans_drop(dictApplicative_0), gopurs_runtime.Int(v_2.IntVal))
							_ = __local_var_6_4
							__t5 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Apply(__local_var_6_4, gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_5_3))
							}))}))}
							goto end_branch_5
						} else {

						}
					}
					{
						if v2_4.Type == 9 && v2_4.IntVal == 489128924 {
							__t5 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}
							goto end_branch_5
						} else {

						}
					}
					{
						__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
					}
				end_branch_5:
					return __t5
				}), v1_3)
			}
		end_branch_6:
			return __t6
		})
	}
}

func Call_Control_Monad_List_Trans_cons(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictApplicative_0 *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
	_ = dictApplicative_0
	// TAST (Let): pure2_1_0 shape=App(Var) bindingType=(Func [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a$scope346), (TypeApp (TypeVar f$scope345) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a$scope346), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f$scope345), (TypeVar a$scope346)])])])])] (TypeApp (TypeVar f$scope345) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a$scope346), (TypeApp (TypeVar f$scope345) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a$scope346), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f$scope345), (TypeVar a$scope346)])])])])]))
	pure2_1_0 := Call_Control_Applicative_pure(dictApplicative_0)
	_ = pure2_1_0
	return gopurs_runtime.Func2(func(lh_2 gopurs_runtime.Value, t_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(pure2_1_0, gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_force(), lh_2), t_3}))})
	})
}

func Call_Control_Monad_List_Trans_unfoldable1ListT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
	_ = dictMonad_0
	// TAST (Let): Applicative0_1_0 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar f$scope353)])
	Applicative0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
	_ = Applicative0_1_0
	// TAST (Let): singleton1_2_1 shape=App(Var) bindingType=(Func [(TypeVar a$scope357)] (TypeApp (TypeVar f$scope353) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a$scope357), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f$scope353), (TypeVar a$scope357)])])]))
	singleton1_2_1 := Call_Control_Monad_List_Trans_singleton(Applicative0_1_0)
	_ = singleton1_2_1
	// TAST (Let): cons1_3_2 shape=App(Var) bindingType=(Func [(ADT ["Data","Lazy","Lazy"] [(TypeVar a$scope357)]), (ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f$scope353) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a$scope357), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f$scope353), (TypeVar a$scope357)])])])])] (TypeApp (TypeVar f$scope353) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a$scope357), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f$scope353), (TypeVar a$scope357)])])]))
	cons1_3_2 := Call_Control_Monad_List_Trans_cons(Applicative0_1_0)
	_ = cons1_3_2
	return gopurs_runtime.Value{Type: 9, IntVal: 3553002490, UnsafePtr: unsafe.Pointer((&Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_4 gopurs_runtime.Value, b_5 gopurs_runtime.Value) gopurs_runtime.Value {
		var go__go_6_3_2 gopurs_runtime.Value
		_ = go__go_6_3_2
		var go__go_6_3_2_cell *gopurs_runtime.Value
		_ = go__go_6_3_2_cell
		// FALLBACK TCO: isLoop=false len=1
		go__go_6_3_2 = gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t8 gopurs_runtime.Value
			{
				var __t_tag_4 gopurs_runtime.Value = (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V1
				_ = __t_tag_4
				if __t_tag_4.Type == 9 && __t_tag_4.IntVal == 930809136 && __t_tag_4.UnsafePtr == nil {
					__t8 = gopurs_runtime.Apply(singleton1_2_1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V0)
					goto end_branch_8
				} else {

				}
			}
			{
				var __t_tag_5 gopurs_runtime.Value = (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V1
				_ = __t_tag_5
				if __t_tag_5.Type == 9 && __t_tag_5.IntVal == 930809136 && __t_tag_5.UnsafePtr != nil {
					// TAST (Let): __local_var_8_6 shape=Other bindingType=Any
					__local_var_8_6 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V0
					_ = __local_var_8_6
					// TAST (Let): __local_var_9_7 shape=Other bindingType=Any
					__local_var_9_7 := (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V1.UnsafePtr).V0
					_ = __local_var_9_7
					__t8 = gopurs_runtime.Apply2(cons1_3_2, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
						return __local_var_8_6
					})), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v1_10 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Apply((*go__go_6_3_2_cell), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Control_Monad_List_Trans_3804580809_138441832(Rebox_Control_Monad_List_Trans_138441832_3804580809(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(f_4, __local_var_9_7)))))})
					})))
					goto end_branch_8
				} else {

				}
			}
			{
				__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_8:
			return __t8
		})
		go__go_6_3_2_cell = &go__go_6_3_2
		return gopurs_runtime.Apply(go__go_6_3_2, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Control_Monad_List_Trans_3804580809_138441832(Rebox_Control_Monad_List_Trans_138441832_3804580809(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(f_4, b_5)))))})
	})}))}
}

func Call_Control_Monad_List_Trans_unfoldableListT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
	_ = dictMonad_0
	// TAST (Let): Applicative0_1_0 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar f$scope361)])
	Applicative0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
	_ = Applicative0_1_0
	// TAST (Let): nil1_2_1 shape=App(Other) bindingType=(TypeApp (TypeVar f$scope361) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a$scope365), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f$scope361), (TypeVar a$scope365)])])])
	nil1_2_1 := gopurs_runtime.Apply(Applicative0_1_0.V1, gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)})
	_ = nil1_2_1
	// TAST (Let): cons1_3_2 shape=App(Var) bindingType=(Func [(ADT ["Data","Lazy","Lazy"] [(TypeVar a$scope365)]), (ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f$scope361) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a$scope365), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f$scope361), (TypeVar a$scope365)])])])])] (TypeApp (TypeVar f$scope361) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a$scope365), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f$scope361), (TypeVar a$scope365)])])]))
	cons1_3_2 := Call_Control_Monad_List_Trans_cons(Applicative0_1_0)
	_ = cons1_3_2
	// TAST (Let): unfoldable1ListT1_4_3 shape=App(Var) bindingType=(ADT ["Data","Unfoldable1","Unfoldable1"] [(TypeApp (TypeVar f$scope361) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f$scope361), (TypeVar a)])])])])
	unfoldable1ListT1_4_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value]](Call_Control_Monad_List_Trans_unfoldable1ListT(dictMonad_0))
	_ = unfoldable1ListT1_4_3
	return gopurs_runtime.Value{Type: 9, IntVal: 2670894170, UnsafePtr: unsafe.Pointer((&Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 3553002490, UnsafePtr: unsafe.Pointer(unfoldable1ListT1_4_3)}
	}), gopurs_runtime.Func2(func(f_5 gopurs_runtime.Value, b_6 gopurs_runtime.Value) gopurs_runtime.Value {
		var go__go_7_4_3 gopurs_runtime.Value
		_ = go__go_7_4_3
		var go__go_7_4_3_cell *gopurs_runtime.Value
		_ = go__go_7_4_3_cell
		// FALLBACK TCO: isLoop=false len=1
		go__go_7_4_3 = gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t7 gopurs_runtime.Value
			{
				if v_8.Type == 9 && v_8.IntVal == 930809136 && v_8.UnsafePtr == nil {
					__t7 = nil1_2_1
					goto end_branch_7
				} else {

				}
			}
			{
				if v_8.Type == 9 && v_8.IntVal == 930809136 && v_8.UnsafePtr != nil {
					// TAST (Let): __local_var_9_5 shape=Other bindingType=Any
					__local_var_9_5 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v_8.UnsafePtr).V0.UnsafePtr).V0
					_ = __local_var_9_5
					// TAST (Let): __local_var_10_6 shape=Other bindingType=Any
					__local_var_10_6 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v_8.UnsafePtr).V0.UnsafePtr).V1
					_ = __local_var_10_6
					__t7 = gopurs_runtime.Apply2(cons1_3_2, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
						return __local_var_9_5
					})), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v1_11 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Apply((*go__go_7_4_3_cell), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Control_Monad_List_Trans_4010058633_3094389156(Rebox_Control_Monad_List_Trans_3094389156_4010058633(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(f_5, __local_var_10_6)))))})
					})))
					goto end_branch_7
				} else {

				}
			}
			{
				__t7 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_7:
			return __t7
		})
		go__go_7_4_3_cell = &go__go_7_4_3
		return gopurs_runtime.Apply(go__go_7_4_3, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Control_Monad_List_Trans_4010058633_3094389156(Rebox_Control_Monad_List_Trans_3094389156_4010058633(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(f_5, b_6)))))})
	})}))}
}

func Call_Control_Monad_List_Trans_semigroupListT(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
	_ = dictApplicative_0
	return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer((&Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{1, Call_Control_Monad_List_Trans_concat(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0))}))}
}

func Call_Control_Monad_List_Trans_concat(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictApplicative_0 *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
	_ = dictApplicative_0
	// TAST (Let): __local_var_1_0 shape=App(Other) bindingType=Any
	__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictApplicative_0.V0, gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
	_ = __local_var_1_0
	return gopurs_runtime.Func2(func(x_2 gopurs_runtime.Value, y_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "map"), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t3 gopurs_runtime.Value
			{
				if v_4.Type == 9 && v_4.IntVal == 1320412129 {
					// TAST (Let): __local_var_5_1 shape=Other bindingType=Any
					__local_var_5_1 := (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1
					_ = __local_var_5_1
					__t3 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_semigroupListT(gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(dictApplicative_0)}), "append"), gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_5_1), y_3)
					}))}))}
					goto end_branch_3
				} else {

				}
			}
			{
				if v_4.Type == 9 && v_4.IntVal == 813447293 {
					// TAST (Let): __local_var_5_2 shape=Other bindingType=Any
					__local_var_5_2 := (*Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0
					_ = __local_var_5_2
					__t3 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_semigroupListT(gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(dictApplicative_0)}), "append"), gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_5_2), y_3)
					}))}))}
					goto end_branch_3
				} else {

				}
			}
			{
				if v_4.Type == 9 && v_4.IntVal == 489128924 {
					__t3 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
						return y_3
					}))}))}
					goto end_branch_3
				} else {

				}
			}
			{
				__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_3:
			return __t3
		}), x_2)
	})
}

func Call_Control_Monad_List_Trans_monoidListT(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
	_ = dictApplicative_0
	// TAST (Let): semigroupListT1_1_0 shape=App(Var) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeApp (TypeVar f$scope179) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a$scope180), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f$scope179), (TypeVar a$scope180)])])])])
	semigroupListT1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](Call_Control_Monad_List_Trans_semigroupListT(dictApplicative_0))
	_ = semigroupListT1_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer((&Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(semigroupListT1_1_0)}
	}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)})}))}
}

func Call_Control_Monad_List_Trans_catMaybes(dictFunctor_0_loop *Constructor_Data_Functor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictFunctor_0 *Constructor_Data_Functor_Functor[gopurs_runtime.Value] = dictFunctor_0_loop
	_ = dictFunctor_0
	return gopurs_runtime.Apply2(Get_Control_Monad_List_Trans_mapMaybe(), gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(dictFunctor_0)}, gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return x_1
	}))
}

func Call_Control_Monad_List_Trans_monadListT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
	_ = dictMonad_0
	return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return Call_Control_Monad_List_Trans_applicativeListT(dictMonad_0)
	}), gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return Call_Control_Monad_List_Trans_bindListT(dictMonad_0)
	})}))}
}

func Call_Control_Monad_List_Trans_bindListT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
	_ = dictMonad_0
	// TAST (Let): __local_var_1_0 shape=App(Other) bindingType=Any
	__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
	_ = __local_var_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return Call_Control_Monad_List_Trans_applyListT(dictMonad_0)
	}), gopurs_runtime.Func2(func(fa_2 gopurs_runtime.Value, f_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "map"), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t4 gopurs_runtime.Value
			{
				if v_4.Type == 9 && v_4.IntVal == 1320412129 {
					// TAST (Let): __local_var_5_1 shape=Other bindingType=Any
					__local_var_5_1 := (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0
					_ = __local_var_5_1
					// TAST (Let): __local_var_6_2 shape=Other bindingType=Any
					__local_var_6_2 := (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1
					_ = __local_var_6_2
					__t4 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_semigroupListT(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{})), "append"), gopurs_runtime.Apply(f_3, __local_var_5_1), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_bindListT(dictMonad_0), "bind"), gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_6_2), f_3))
					}))}))}
					goto end_branch_4
				} else {

				}
			}
			{
				if v_4.Type == 9 && v_4.IntVal == 813447293 {
					// TAST (Let): __local_var_5_3 shape=Other bindingType=Any
					__local_var_5_3 := (*Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0
					_ = __local_var_5_3
					__t4 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_bindListT(dictMonad_0), "bind"), gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_5_3), f_3)
					}))}))}
					goto end_branch_4
				} else {

				}
			}
			{
				if v_4.Type == 9 && v_4.IntVal == 489128924 {
					__t4 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}
					goto end_branch_4
				} else {

				}
			}
			{
				__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_4:
			return __t4
		}), fa_2)
	})}))}
}

func Call_Control_Monad_List_Trans_applyListT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
	_ = dictMonad_0
	// TAST (Let): functorListT1_1_0 shape=App(Var) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar f$scope403) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f$scope403), (TypeVar a)])])])])
	functorListT1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Call_Control_Monad_List_Trans_functorListT(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})))
	_ = functorListT1_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorListT1_1_0)}
	}), Call_Control_Monad_ap(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](Call_Control_Monad_List_Trans_monadListT(dictMonad_0)))}))}
}

func Call_Control_Monad_List_Trans_applicativeListT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
	_ = dictMonad_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return Call_Control_Monad_List_Trans_applyListT(dictMonad_0)
	}), Call_Control_Monad_List_Trans_singleton(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{})))}))}
}

func Call_Control_Monad_List_Trans_monadEffectListT(dictMonadEffect_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
	_ = dictMonadEffect_0
	// TAST (Let): Monad0_1_0 shape=App(Other) bindingType=Any
	Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "Monad0"), gopurs_runtime.Value{})
	_ = Monad0_1_0
	// TAST (Let): monadListT1_2_1 shape=App(Var) bindingType=(ADT ["Control","Monad","Monad"] [(TypeApp (TypeVar m$scope189) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar m$scope189), (TypeVar a)])])])])
	monadListT1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](Call_Control_Monad_List_Trans_monadListT(Monad0_1_0))
	_ = monadListT1_2_1
	return gopurs_runtime.Value{Type: 9, IntVal: 2217729261, UnsafePtr: unsafe.Pointer((&Constructor_Effect_Class_MonadEffect[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadListT1_2_1)}
	}), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Call_Control_Monad_List_Trans_fromEffect(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}))), Call_Effect_Class_liftEffect(gopurs_runtime.CoerceToStruct[Constructor_Effect_Class_MonadEffect[gopurs_runtime.Value]](dictMonadEffect_0)))}))}
}

func Call_Control_Monad_List_Trans_monadSTListT(dictMonadST_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonadST_0 gopurs_runtime.Value = dictMonadST_0_loop
	_ = dictMonadST_0
	// TAST (Let): Monad0_1_0 shape=App(Other) bindingType=Any
	Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadST_0, "Monad0"), gopurs_runtime.Value{})
	_ = Monad0_1_0
	// TAST (Let): monadListT1_2_1 shape=App(Var) bindingType=(ADT ["Control","Monad","Monad"] [(TypeApp (TypeVar m$scope3) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar m$scope3), (TypeVar a)])])])])
	monadListT1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](Call_Control_Monad_List_Trans_monadListT(Monad0_1_0))
	_ = monadListT1_2_1
	return gopurs_runtime.Value{Type: 9, IntVal: 2155655715, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_ST_Class_MonadST[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadListT1_2_1)}
	}), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Call_Control_Monad_List_Trans_fromEffect(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}))), Call_Control_Monad_ST_Class_liftST(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_ST_Class_MonadST[gopurs_runtime.Value, gopurs_runtime.Value]](dictMonadST_0)))}))}
}

func Call_Control_Monad_List_Trans_altListT(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
	_ = dictApplicative_0
	// TAST (Let): functorListT1_1_0 shape=App(Var) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar f$scope418) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f$scope418), (TypeVar a)])])])])
	functorListT1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Call_Control_Monad_List_Trans_functorListT(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})))
	_ = functorListT1_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer((&Constructor_Control_Alt_Alt[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorListT1_1_0)}
	}), Call_Control_Monad_List_Trans_concat(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0))}))}
}

func Call_Control_Monad_List_Trans_plusListT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
	_ = dictMonad_0
	// TAST (Let): Applicative0_1_0 shape=App(Other) bindingType=Any
	Applicative0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{})
	_ = Applicative0_1_0
	// TAST (Let): altListT1_2_1 shape=App(Var) bindingType=(ADT ["Control","Alt","Alt"] [(TypeApp (TypeVar f$scope127) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f$scope127), (TypeVar a)])])])])
	altListT1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Alt_Alt[gopurs_runtime.Value]](Call_Control_Monad_List_Trans_altListT(Applicative0_1_0))
	_ = altListT1_2_1
	return gopurs_runtime.Value{Type: 9, IntVal: 3709470893, UnsafePtr: unsafe.Pointer((&Constructor_Control_Plus_Plus[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer(altListT1_2_1)}
	}), Call_Control_Monad_List_Trans_nil(Applicative0_1_0)}))}
}

func Call_Control_Monad_List_Trans_alternativeListT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
	_ = dictMonad_0
	// TAST (Let): applicativeListT1_1_0 shape=App(Var) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeApp (TypeVar f$scope415) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f$scope415), (TypeVar a)])])])])
	applicativeListT1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Call_Control_Monad_List_Trans_applicativeListT(dictMonad_0))
	_ = applicativeListT1_1_0
	// TAST (Let): plusListT1_2_1 shape=App(Var) bindingType=(ADT ["Control","Plus","Plus"] [(TypeApp (TypeVar f$scope415) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f$scope415), (TypeVar a)])])])])
	plusListT1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Plus_Plus[gopurs_runtime.Value]](Call_Control_Monad_List_Trans_plusListT(dictMonad_0))
	_ = plusListT1_2_1
	return gopurs_runtime.Value{Type: 9, IntVal: 397869517, UnsafePtr: unsafe.Pointer((&Constructor_Control_Alternative_Alternative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(applicativeListT1_1_0)}
	}), gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 3709470893, UnsafePtr: unsafe.Pointer(plusListT1_2_1)}
	})}))}
}

func Call_Control_Monad_List_Trans_monadPlusListT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
	_ = dictMonad_0
	// TAST (Let): monadListT1_1_0 shape=App(Var) bindingType=(ADT ["Control","Monad","Monad"] [(TypeApp (TypeVar f$scope183) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f$scope183), (TypeVar a)])])])])
	monadListT1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](Call_Control_Monad_List_Trans_monadListT(dictMonad_0))
	_ = monadListT1_1_0
	// TAST (Let): alternativeListT1_2_1 shape=App(Var) bindingType=(ADT ["Control","Alternative","Alternative"] [(TypeApp (TypeVar f$scope183) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f$scope183), (TypeVar a)])])])])
	alternativeListT1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Alternative_Alternative[gopurs_runtime.Value]](Call_Control_Monad_List_Trans_alternativeListT(dictMonad_0))
	_ = alternativeListT1_2_1
	return gopurs_runtime.Value{Type: 9, IntVal: 3236234573, UnsafePtr: unsafe.Pointer((&Constructor_Control_MonadPlus_MonadPlus[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 397869517, UnsafePtr: unsafe.Pointer(alternativeListT1_2_1)}
	}), gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadListT1_1_0)}
	})}))}
}

func Rebox_Control_Monad_List_Trans_1239976062_3603546092(in *Constructor_Control_Monad_Rec_Class_Done[struct {
	a gopurs_runtime.Value
	b gopurs_runtime.Value
	/*gopurs:record-v1:2:0061,0062*/
}, gopurs_runtime.Value]) *Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Control_Monad_List_Trans_138441832_3804580809(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, *Constructor_Data_Maybe_Just[gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, *Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{}
	out.V0 = in.V0
	out.V1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](in.V1)
	return out
}

func Rebox_Control_Monad_List_Trans_1785332133_138441832(in *Constructor_Data_Tuple_Tuple[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value], gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(in.V0)}
	out.V1 = in.V1
	return out
}

func Rebox_Control_Monad_List_Trans_2549197956_3094389156(in *Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value], gopurs_runtime.Value]]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Control_Monad_List_Trans_1785332133_138441832(in.V0))}
	return out
}

func Rebox_Control_Monad_List_Trans_2812149806_3689823567(in *Constructor_Data_Functor_Functor[gopurs_runtime.Value]) *Constructor_Data_Functor_Functor[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Functor_Functor[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Control_Monad_List_Trans_3094389156_4010058633(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{}
	out.V0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](in.V0)
	return out
}

func Rebox_Control_Monad_List_Trans_3804580809_138441832(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, *Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = in.V0
	out.V1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(in.V1)}
	return out
}

func Rebox_Control_Monad_List_Trans_4010058633_3094389156(in *Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(in.V0)}
	return out
}

func Rebox_Control_Monad_List_Trans_784458114_4008603408(in *Constructor_Control_Monad_Rec_Class_Loop[struct {
	a gopurs_runtime.Value
	b gopurs_runtime.Value
	/*gopurs:record-v1:2:0061,0062*/
}, gopurs_runtime.Value]) *Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = func() gopurs_runtime.Value {
		orig := in.V0
		_ = orig
		return gopurs_runtime.RecordDict2("a", "b", orig.a, orig.b)
	}()
	return out
}
