package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Effect_monadEffect gopurs_runtime.Value
var once_Effect_monadEffect sync.Once

func Get_Effect_monadEffect() gopurs_runtime.Value {
	once_Effect_monadEffect.Do(func() {
		cache_Effect_monadEffect = gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Effect_applicativeEffect()))}
		}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect()))}
		})}))}
	})
	return cache_Effect_monadEffect
}

var cache_Effect_bindEffect gopurs_runtime.Value
var once_Effect_bindEffect sync.Once

func Get_Effect_bindEffect() gopurs_runtime.Value {
	once_Effect_bindEffect.Do(func() {
		cache_Effect_bindEffect = gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Effect_applyEffect()))}
		}), Get_Effect_bindE()}))}
	})
	return cache_Effect_bindEffect
}

var cache_Effect_applyEffect gopurs_runtime.Value
var once_Effect_applyEffect sync.Once

func Get_Effect_applyEffect() gopurs_runtime.Value {
	once_Effect_applyEffect.Do(func() {
		cache_Effect_applyEffect = gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()))}
		}), Call_Control_Monad_ap(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](Get_Effect_monadEffect()))}))}
	})
	return cache_Effect_applyEffect
}

var cache_Effect_applicativeEffect gopurs_runtime.Value
var once_Effect_applicativeEffect sync.Once

func Get_Effect_applicativeEffect() gopurs_runtime.Value {
	once_Effect_applicativeEffect.Do(func() {
		cache_Effect_applicativeEffect = gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Effect_applyEffect()))}
		}), Get_Effect_pureE()}))}
	})
	return cache_Effect_applicativeEffect
}

var cache_Effect_functorEffect gopurs_runtime.Value
var once_Effect_functorEffect sync.Once

func Get_Effect_functorEffect() gopurs_runtime.Value {
	once_Effect_functorEffect.Do(func() {
		cache_Effect_functorEffect = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Control_Applicative_liftA1(), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Effect_applicativeEffect()))})}))}
	})
	return cache_Effect_functorEffect
}

var cache_Effect_lift2 gopurs_runtime.Value
var once_Effect_lift2 sync.Once

func Get_Effect_lift2() gopurs_runtime.Value {
	once_Effect_lift2.Do(func() {
		cache_Effect_lift2 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value, b_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Effect_lift2(f_0_box, a_1_box, b_2_box)
		})
	})
	return cache_Effect_lift2
}

var cache_Effect_semigroupEffect gopurs_runtime.Value
var once_Effect_semigroupEffect sync.Once

func Get_Effect_semigroupEffect() gopurs_runtime.Value {
	once_Effect_semigroupEffect.Do(func() {
		cache_Effect_semigroupEffect = gopurs_runtime.Func(func(dictSemigroup_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Effect_semigroupEffect(dictSemigroup_0_box)
		})
	})
	return cache_Effect_semigroupEffect
}

var cache_Effect_monoidEffect gopurs_runtime.Value
var once_Effect_monoidEffect sync.Once

func Get_Effect_monoidEffect() gopurs_runtime.Value {
	once_Effect_monoidEffect.Do(func() {
		cache_Effect_monoidEffect = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Effect_monoidEffect(dictMonoid_0_box)
		})
	})
	return cache_Effect_monoidEffect
}

func Call_Effect_lift2(f_0_loop gopurs_runtime.Value, a_1_loop gopurs_runtime.Value, b_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	var a_1 gopurs_runtime.Value = a_1_loop
	_ = a_1
	var b_2 gopurs_runtime.Value = b_2_loop
	_ = b_2
	return gopurs_runtime.Apply2(Call_Control_Monad_ap(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](Get_Effect_monadEffect())), Call_Control_Applicative_liftA1(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Effect_applicativeEffect()), f_0, a_1), b_2)
}

func Call_Effect_semigroupEffect(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
	_ = dictSemigroup_0
	// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=(Func [(TypeVar a$scope1), (TypeVar a$scope1)] (TypeVar a$scope1))
	__local_var_1_0 := Call_Data_Semigroup_go__append(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](dictSemigroup_0))
	_ = __local_var_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer((&Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(a_2 gopurs_runtime.Value, b_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(Call_Control_Monad_ap(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](Get_Effect_monadEffect())), Call_Control_Applicative_liftA1(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Effect_applicativeEffect()), __local_var_1_0, a_2), b_3)
	})}))}
}

func Call_Effect_monoidEffect(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
	_ = dictMonoid_0
	// TAST (Let): semigroupEffect1_1_0 shape=App(Var) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(ADT ["Effect","Effect"] [(TypeVar a$scope4)])])
	semigroupEffect1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](Call_Effect_semigroupEffect(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})))
	_ = semigroupEffect1_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer((&Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(semigroupEffect1_1_0)}
	}), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_2_1 shape=Other bindingType=Any
		__local_var_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
		_ = __local_var_2_1
		return __local_var_2_1
	})}))}
}

func Get_Effect_bindE() gopurs_runtime.Value {
	return _Gopurs_Effect_BindE
}

func Get_Effect_forE() gopurs_runtime.Value {
	return _Gopurs_Effect_ForE
}

func Get_Effect_foreachE() gopurs_runtime.Value {
	return _Gopurs_Effect_ForeachE
}

func Get_Effect_pureE() gopurs_runtime.Value {
	return _Gopurs_Effect_PureE
}

func Get_Effect_untilE() gopurs_runtime.Value {
	return _Gopurs_Effect_UntilE
}

func Get_Effect_whileE() gopurs_runtime.Value {
	return _Gopurs_Effect_WhileE
}
