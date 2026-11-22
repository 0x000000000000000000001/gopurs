package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Monad_ST_Internal_go__new gopurs_runtime.Value
var once_Control_Monad_ST_Internal_go__new sync.Once

func Get_Control_Monad_ST_Internal_go__new() gopurs_runtime.Value {
	once_Control_Monad_ST_Internal_go__new.Do(func() {
		cache_Control_Monad_ST_Internal_go__new = Get_Control_Monad_ST_Internal_newImpl()
	})
	return cache_Control_Monad_ST_Internal_go__new
}

var cache_Control_Monad_ST_Internal_modify_prime_ gopurs_runtime.Value
var once_Control_Monad_ST_Internal_modify_prime_ sync.Once

func Get_Control_Monad_ST_Internal_modify_prime_() gopurs_runtime.Value {
	once_Control_Monad_ST_Internal_modify_prime_.Do(func() {
		cache_Control_Monad_ST_Internal_modify_prime_ = Get_Control_Monad_ST_Internal_modifyImpl()
	})
	return cache_Control_Monad_ST_Internal_modify_prime_
}

var cache_Control_Monad_ST_Internal_modify gopurs_runtime.Value
var once_Control_Monad_ST_Internal_modify sync.Once

func Get_Control_Monad_ST_Internal_modify() gopurs_runtime.Value {
	once_Control_Monad_ST_Internal_modify.Do(func() {
		cache_Control_Monad_ST_Internal_modify = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_ST_Internal_modify(f_0_box)
		})
	})
	return cache_Control_Monad_ST_Internal_modify
}

var cache_Control_Monad_ST_Internal_functorST gopurs_runtime.Value
var once_Control_Monad_ST_Internal_functorST sync.Once

func Get_Control_Monad_ST_Internal_functorST() gopurs_runtime.Value {
	once_Control_Monad_ST_Internal_functorST.Do(func() {
		cache_Control_Monad_ST_Internal_functorST = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, Get_Control_Monad_ST_Internal_map_()}))}
	})
	return cache_Control_Monad_ST_Internal_functorST
}

var cache_Control_Monad_ST_Internal_go__map gopurs_runtime.Value
var once_Control_Monad_ST_Internal_go__map sync.Once

func Get_Control_Monad_ST_Internal_go__map() gopurs_runtime.Value {
	once_Control_Monad_ST_Internal_go__map.Do(func() {
		cache_Control_Monad_ST_Internal_go__map = Get_Control_Monad_ST_Internal_map_()
	})
	return cache_Control_Monad_ST_Internal_go__map
}

var cache_Control_Monad_ST_Internal_void gopurs_runtime.Value
var once_Control_Monad_ST_Internal_void sync.Once

func Get_Control_Monad_ST_Internal_void() gopurs_runtime.Value {
	once_Control_Monad_ST_Internal_void.Do(func() {
		cache_Control_Monad_ST_Internal_void = gopurs_runtime.Func(func(__local_var_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_ST_Internal_void(__local_var_0_box)
		})
	})
	return cache_Control_Monad_ST_Internal_void
}

var cache_Control_Monad_ST_Internal_map1 gopurs_runtime.Value
var once_Control_Monad_ST_Internal_map1 sync.Once

func Get_Control_Monad_ST_Internal_map1() gopurs_runtime.Value {
	once_Control_Monad_ST_Internal_map1.Do(func() {
		cache_Control_Monad_ST_Internal_map1 = Get_Control_Monad_ST_Internal_map_()
	})
	return cache_Control_Monad_ST_Internal_map1
}

var cache_Control_Monad_ST_Internal_go__for gopurs_runtime.Value
var once_Control_Monad_ST_Internal_go__for sync.Once

func Get_Control_Monad_ST_Internal_go__for() gopurs_runtime.Value {
	once_Control_Monad_ST_Internal_go__for.Do(func() {
		cache_Control_Monad_ST_Internal_go__for = Get_Control_Monad_ST_Internal_forImpl()
	})
	return cache_Control_Monad_ST_Internal_go__for
}

var cache_Control_Monad_ST_Internal_monadST gopurs_runtime.Value
var once_Control_Monad_ST_Internal_monadST sync.Once

func Get_Control_Monad_ST_Internal_monadST() gopurs_runtime.Value {
	once_Control_Monad_ST_Internal_monadST.Do(func() {
		cache_Control_Monad_ST_Internal_monadST = gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Get_Control_Monad_ST_Internal_applicativeST()
		}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Get_Control_Monad_ST_Internal_bindST()
		})}))}
	})
	return cache_Control_Monad_ST_Internal_monadST
}

var cache_Control_Monad_ST_Internal_bindST gopurs_runtime.Value
var once_Control_Monad_ST_Internal_bindST sync.Once

func Get_Control_Monad_ST_Internal_bindST() gopurs_runtime.Value {
	once_Control_Monad_ST_Internal_bindST.Do(func() {
		cache_Control_Monad_ST_Internal_bindST = gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Get_Control_Monad_ST_Internal_applyST()
		}), Get_Control_Monad_ST_Internal_bind_()}))}
	})
	return cache_Control_Monad_ST_Internal_bindST
}

var cache_Control_Monad_ST_Internal_applyST gopurs_runtime.Value
var once_Control_Monad_ST_Internal_applyST sync.Once

func Get_Control_Monad_ST_Internal_applyST() gopurs_runtime.Value {
	once_Control_Monad_ST_Internal_applyST.Do(func() {
		cache_Control_Monad_ST_Internal_applyST = gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Get_Control_Monad_ST_Internal_functorST()
		}), Call_Control_Monad_ap(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](Get_Control_Monad_ST_Internal_monadST()))}))}
	})
	return cache_Control_Monad_ST_Internal_applyST
}

var cache_Control_Monad_ST_Internal_applicativeST gopurs_runtime.Value
var once_Control_Monad_ST_Internal_applicativeST sync.Once

func Get_Control_Monad_ST_Internal_applicativeST() gopurs_runtime.Value {
	once_Control_Monad_ST_Internal_applicativeST.Do(func() {
		cache_Control_Monad_ST_Internal_applicativeST = gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Get_Control_Monad_ST_Internal_applyST()
		}), Get_Control_Monad_ST_Internal_pure_()}))}
	})
	return cache_Control_Monad_ST_Internal_applicativeST
}

var cache_Control_Monad_ST_Internal_lift2 gopurs_runtime.Value
var once_Control_Monad_ST_Internal_lift2 sync.Once

func Get_Control_Monad_ST_Internal_lift2() gopurs_runtime.Value {
	once_Control_Monad_ST_Internal_lift2.Do(func() {
		cache_Control_Monad_ST_Internal_lift2 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value, b_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_ST_Internal_lift2(f_0_box, a_1_box, b_2_box)
		})
	})
	return cache_Control_Monad_ST_Internal_lift2
}

var cache_Control_Monad_ST_Internal_bind gopurs_runtime.Value
var once_Control_Monad_ST_Internal_bind sync.Once

func Get_Control_Monad_ST_Internal_bind() gopurs_runtime.Value {
	once_Control_Monad_ST_Internal_bind.Do(func() {
		cache_Control_Monad_ST_Internal_bind = Get_Control_Monad_ST_Internal_bind_()
	})
	return cache_Control_Monad_ST_Internal_bind
}

var cache_Control_Monad_ST_Internal_bindFlipped gopurs_runtime.Value
var once_Control_Monad_ST_Internal_bindFlipped sync.Once

func Get_Control_Monad_ST_Internal_bindFlipped() gopurs_runtime.Value {
	once_Control_Monad_ST_Internal_bindFlipped.Do(func() {
		cache_Control_Monad_ST_Internal_bindFlipped = gopurs_runtime.Func2(func(b_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_ST_Internal_bindFlipped(b_0_box, a_1_box)
		})
	})
	return cache_Control_Monad_ST_Internal_bindFlipped
}

var cache_Control_Monad_ST_Internal_discard gopurs_runtime.Value
var once_Control_Monad_ST_Internal_discard sync.Once

func Get_Control_Monad_ST_Internal_discard() gopurs_runtime.Value {
	once_Control_Monad_ST_Internal_discard.Do(func() {
		cache_Control_Monad_ST_Internal_discard = Get_Control_Monad_ST_Internal_bind_()
	})
	return cache_Control_Monad_ST_Internal_discard
}

var cache_Control_Monad_ST_Internal_bind1 gopurs_runtime.Value
var once_Control_Monad_ST_Internal_bind1 sync.Once

func Get_Control_Monad_ST_Internal_bind1() gopurs_runtime.Value {
	once_Control_Monad_ST_Internal_bind1.Do(func() {
		cache_Control_Monad_ST_Internal_bind1 = Get_Control_Monad_ST_Internal_bind_()
	})
	return cache_Control_Monad_ST_Internal_bind1
}

var cache_Control_Monad_ST_Internal_pure gopurs_runtime.Value
var once_Control_Monad_ST_Internal_pure sync.Once

func Get_Control_Monad_ST_Internal_pure() gopurs_runtime.Value {
	once_Control_Monad_ST_Internal_pure.Do(func() {
		cache_Control_Monad_ST_Internal_pure = Get_Control_Monad_ST_Internal_pure_()
	})
	return cache_Control_Monad_ST_Internal_pure
}

var cache_Control_Monad_ST_Internal_pure1 gopurs_runtime.Value
var once_Control_Monad_ST_Internal_pure1 sync.Once

func Get_Control_Monad_ST_Internal_pure1() gopurs_runtime.Value {
	once_Control_Monad_ST_Internal_pure1.Do(func() {
		cache_Control_Monad_ST_Internal_pure1 = Get_Control_Monad_ST_Internal_pure_()
	})
	return cache_Control_Monad_ST_Internal_pure1
}

var cache_Control_Monad_ST_Internal_semigroupST gopurs_runtime.Value
var once_Control_Monad_ST_Internal_semigroupST sync.Once

func Get_Control_Monad_ST_Internal_semigroupST() gopurs_runtime.Value {
	once_Control_Monad_ST_Internal_semigroupST.Do(func() {
		cache_Control_Monad_ST_Internal_semigroupST = gopurs_runtime.Func(func(dictSemigroup_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_ST_Internal_semigroupST(dictSemigroup_0_box)
		})
	})
	return cache_Control_Monad_ST_Internal_semigroupST
}

var cache_Control_Monad_ST_Internal_monadRecST gopurs_runtime.Value
var once_Control_Monad_ST_Internal_monadRecST sync.Once

func Get_Control_Monad_ST_Internal_monadRecST() gopurs_runtime.Value {
	once_Control_Monad_ST_Internal_monadRecST.Do(func() {
		cache_Control_Monad_ST_Internal_monadRecST = gopurs_runtime.Value{Type: 9, IntVal: 3709389635, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Get_Control_Monad_ST_Internal_monadST()
		}), gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, a_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
				// TAST (Let): __local_var_2_0 shape=App(Other) bindingType=(ADT ["Control","Monad","ST","Internal","ST"] [(TypeVar r$scope20), (ADT ["Control","Monad","Rec","Class","Step"] [(TypeVar a$scope24), (TypeVar b$scope25)])])
				__local_var_2_0 := gopurs_runtime.Apply(f_0, a_1)
				_ = __local_var_2_0
				__local_var_3_2 := gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Value{})
				_ = __local_var_3_2
				r_3_1 := gopurs_runtime.NewSTRef(__local_var_3_2)
				_ = r_3_1
				_dollar___unused_4_4 := gopurs_runtime.Apply3(Get_Control_Monad_ST_Internal_while(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
					__local_var_4_5 := gopurs_runtime.ReadSTRef(r_3_1)
					_ = __local_var_4_5
					return gopurs_runtime.Bool((__local_var_4_5.Type == 9 && __local_var_4_5.IntVal == 525585346))
				}), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
					v_4_6 := gopurs_runtime.ReadSTRef(r_3_1)
					_ = v_4_6
					var __t9 gopurs_runtime.Value
					{
						if v_4_6.Type == 9 && v_4_6.IntVal == 525585346 {
							__t9 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
								e_5_7 := gopurs_runtime.Apply2(f_0, (*Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value])(v_4_6.UnsafePtr).V0, gopurs_runtime.Value{})
								_ = e_5_7
								__local_var_6_8 := gopurs_runtime.WriteSTRef(r_3_1, e_5_7)
								_ = __local_var_6_8
								return Get_Data_Unit_unit()
							})
							goto end_branch_9
						} else {

						}
					}
					{
						if v_4_6.Type == 9 && v_4_6.IntVal == 60402430 {
							__t9 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
								return Get_Data_Unit_unit()
							})
							goto end_branch_9
						} else {

						}
					}
					{
						__t9 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
					}
				end_branch_9:
					return gopurs_runtime.Apply(__t9, gopurs_runtime.Value{})
				}), gopurs_runtime.Value{})
				_ = _dollar___unused_4_4
				__local_var_5_10 := gopurs_runtime.ReadSTRef(r_3_1)
				_ = __local_var_5_10
				var __t11 gopurs_runtime.Value
				{
					if __local_var_5_10.Type == 9 && __local_var_5_10.IntVal == 60402430 {
						__t11 = (*Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_5_10.UnsafePtr).V0
						goto end_branch_11
					} else {

					}
				}
				{
					__t11 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
				}
			end_branch_11:
				return __t11
			})
		})}))}
	})
	return cache_Control_Monad_ST_Internal_monadRecST
}

var cache_Control_Monad_ST_Internal_monoidST gopurs_runtime.Value
var once_Control_Monad_ST_Internal_monoidST sync.Once

func Get_Control_Monad_ST_Internal_monoidST() gopurs_runtime.Value {
	once_Control_Monad_ST_Internal_monoidST.Do(func() {
		cache_Control_Monad_ST_Internal_monoidST = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_ST_Internal_monoidST(dictMonoid_0_box)
		})
	})
	return cache_Control_Monad_ST_Internal_monoidST
}

func Call_Control_Monad_ST_Internal_modify(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	return gopurs_runtime.Apply(Get_Control_Monad_ST_Internal_modifyImpl(), gopurs_runtime.Func(func(s_1 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): s_prime__2_0 shape=App(Other) bindingType=(TypeVar a$scope52)
		s_prime__2_0 := gopurs_runtime.Apply(f_0, s_1)
		_ = s_prime__2_0
		return gopurs_runtime.RecordDict2("state", "value", s_prime__2_0, s_prime__2_0)
	}))
}

func Call_Control_Monad_ST_Internal_void(__local_var_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
	_ = __local_var_0
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		__local_var_1_0 := gopurs_runtime.Apply(__local_var_0, gopurs_runtime.Value{})
		_ = __local_var_1_0
		return Get_Data_Unit_unit()
	})
}

func Call_Control_Monad_ST_Internal_lift2(f_0_loop gopurs_runtime.Value, a_1_loop gopurs_runtime.Value, b_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	var a_1 gopurs_runtime.Value = a_1_loop
	_ = a_1
	var b_2 gopurs_runtime.Value = b_2_loop
	_ = b_2
	return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_applyST(), "apply"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		__local_var_3_0 := gopurs_runtime.Apply(a_1, gopurs_runtime.Value{})
		_ = __local_var_3_0
		return gopurs_runtime.Apply(f_0, __local_var_3_0)
	}), b_2)
}

func Call_Control_Monad_ST_Internal_bindFlipped(b_0_loop gopurs_runtime.Value, a_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var b_0 gopurs_runtime.Value = b_0_loop
	_ = b_0
	var a_1 gopurs_runtime.Value = a_1_loop
	_ = a_1
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		__local_var_2_0 := gopurs_runtime.Apply(a_1, gopurs_runtime.Value{})
		_ = __local_var_2_0
		return gopurs_runtime.Apply2(b_0, __local_var_2_0, gopurs_runtime.Value{})
	})
}

func Call_Control_Monad_ST_Internal_semigroupST(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
	_ = dictSemigroup_0
	// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=(Func [(TypeVar a$scope2), (TypeVar a$scope2)] (TypeVar a$scope2))
	__local_var_1_0 := Call_Data_Semigroup_go__append(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](dictSemigroup_0))
	_ = __local_var_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer((&Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(a_2 gopurs_runtime.Value, b_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_applyST(), "apply"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			__local_var_4_1 := gopurs_runtime.Apply(a_2, gopurs_runtime.Value{})
			_ = __local_var_4_1
			return gopurs_runtime.Apply(__local_var_1_0, __local_var_4_1)
		}), b_3)
	})}))}
}

func Call_Control_Monad_ST_Internal_monoidST(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
	_ = dictMonoid_0
	// TAST (Let): semigroupST1_1_0 shape=App(Var) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(ADT ["Control","Monad","ST","Internal","ST"] [(TypeVar r$scope14), (TypeVar a$scope13)])])
	semigroupST1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](Call_Control_Monad_ST_Internal_semigroupST(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})))
	_ = semigroupST1_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer((&Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(semigroupST1_1_0)}
	}), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_2_1 shape=Other bindingType=Any
		__local_var_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
		_ = __local_var_2_1
		return __local_var_2_1
	})}))}
}

func Get_Control_Monad_ST_Internal_bind_() gopurs_runtime.Value {
	return _Gopurs_Control_Monad_ST_Internal_Bind_
}

func Get_Control_Monad_ST_Internal_forImpl() gopurs_runtime.Value {
	return _Gopurs_Control_Monad_ST_Internal_ForImpl
}

func Get_Control_Monad_ST_Internal_foreach() gopurs_runtime.Value {
	return _Gopurs_Control_Monad_ST_Internal_Foreach
}

func Get_Control_Monad_ST_Internal_map_() gopurs_runtime.Value {
	return _Gopurs_Control_Monad_ST_Internal_Map_
}

func Get_Control_Monad_ST_Internal_modifyImpl() gopurs_runtime.Value {
	return _Gopurs_Control_Monad_ST_Internal_ModifyImpl
}

func Get_Control_Monad_ST_Internal_newImpl() gopurs_runtime.Value {
	return _Gopurs_Control_Monad_ST_Internal_NewImpl
}

func Get_Control_Monad_ST_Internal_pure_() gopurs_runtime.Value {
	return _Gopurs_Control_Monad_ST_Internal_Pure_
}

func Get_Control_Monad_ST_Internal_read() gopurs_runtime.Value {
	return _Gopurs_Control_Monad_ST_Internal_Read
}

func Get_Control_Monad_ST_Internal_run() gopurs_runtime.Value {
	return _Gopurs_Control_Monad_ST_Internal_Run
}

func Get_Control_Monad_ST_Internal_while() gopurs_runtime.Value {
	return _Gopurs_Control_Monad_ST_Internal_While
}

func Get_Control_Monad_ST_Internal_write() gopurs_runtime.Value {
	return _Gopurs_Control_Monad_ST_Internal_Write
}
