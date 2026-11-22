package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Ord_Max_go__append gopurs_runtime.Value
var once_Data_Ord_Max_go__append sync.Once

func Get_Data_Ord_Max_go__append() gopurs_runtime.Value {
	once_Data_Ord_Max_go__append.Do(func() {
		cache_Data_Ord_Max_go__append = Get_Data_Semigroup_append__1694605992()
	})
	return cache_Data_Ord_Max_go__append
}

var cache_Data_Ord_Max_Max gopurs_runtime.Value
var once_Data_Ord_Max_Max sync.Once

func Get_Data_Ord_Max_Max() gopurs_runtime.Value {
	once_Data_Ord_Max_Max.Do(func() {
		cache_Data_Ord_Max_Max = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Ord_Max_Max(x_0_box)
		})
	})
	return cache_Data_Ord_Max_Max
}

var cache_Data_Ord_Max_showMax gopurs_runtime.Value
var once_Data_Ord_Max_showMax sync.Once

func Get_Data_Ord_Max_showMax() gopurs_runtime.Value {
	once_Data_Ord_Max_showMax.Do(func() {
		cache_Data_Ord_Max_showMax = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Ord_Max_showMax(dictShow_0_box)
		})
	})
	return cache_Data_Ord_Max_showMax
}

var cache_Data_Ord_Max_semigroupMax gopurs_runtime.Value
var once_Data_Ord_Max_semigroupMax sync.Once

func Get_Data_Ord_Max_semigroupMax() gopurs_runtime.Value {
	once_Data_Ord_Max_semigroupMax.Do(func() {
		cache_Data_Ord_Max_semigroupMax = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Ord_Max_semigroupMax(dictOrd_0_box)
		})
	})
	return cache_Data_Ord_Max_semigroupMax
}

var cache_Data_Ord_Max_newtypeMax gopurs_runtime.Value
var once_Data_Ord_Max_newtypeMax sync.Once

func Get_Data_Ord_Max_newtypeMax() gopurs_runtime.Value {
	once_Data_Ord_Max_newtypeMax.Do(func() {
		cache_Data_Ord_Max_newtypeMax = gopurs_runtime.Value{Type: 9, IntVal: 3322196858, UnsafePtr: unsafe.Pointer((&Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{}
		})}))}
	})
	return cache_Data_Ord_Max_newtypeMax
}

var cache_Data_Ord_Max_monoidMax gopurs_runtime.Value
var once_Data_Ord_Max_monoidMax sync.Once

func Get_Data_Ord_Max_monoidMax() gopurs_runtime.Value {
	once_Data_Ord_Max_monoidMax.Do(func() {
		cache_Data_Ord_Max_monoidMax = gopurs_runtime.Func(func(dictBounded_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Ord_Max_monoidMax(dictBounded_0_box)
		})
	})
	return cache_Data_Ord_Max_monoidMax
}

var cache_Data_Ord_Max_eqMax gopurs_runtime.Value
var once_Data_Ord_Max_eqMax sync.Once

func Get_Data_Ord_Max_eqMax() gopurs_runtime.Value {
	once_Data_Ord_Max_eqMax.Do(func() {
		cache_Data_Ord_Max_eqMax = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Ord_Max_eqMax(dictEq_0_box)
		})
	})
	return cache_Data_Ord_Max_eqMax
}

var cache_Data_Ord_Max_ordMax gopurs_runtime.Value
var once_Data_Ord_Max_ordMax sync.Once

func Get_Data_Ord_Max_ordMax() gopurs_runtime.Value {
	once_Data_Ord_Max_ordMax.Do(func() {
		cache_Data_Ord_Max_ordMax = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Ord_Max_ordMax(dictOrd_0_box)
		})
	})
	return cache_Data_Ord_Max_ordMax
}

func Call_Data_Ord_Max_Max(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Data_Ord_Max_showMax(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
	_ = dictShow_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer((&Constructor_Data_Show_Show[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Str(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString("(Max ", gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), v_1).StrVal()), ")"))
	})}))}
}

func Call_Data_Ord_Max_semigroupMax(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
	_ = dictOrd_0
	return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer((&Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): v_3_0 shape=App(Other) bindingType=(ADT ["Data","Ordering","Ordering"] [])
		v_3_0 := uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), v_1, v1_2).IntVal)
		_ = v_3_0
		var __t1 gopurs_runtime.Value
		{
			if v_3_0 == 1527465420 {
				__t1 = v1_2
				goto end_branch_1
			} else {

			}
		}
		{
			if v_3_0 == 902936544 {
				__t1 = v_1
				goto end_branch_1
			} else {

			}
		}
		{
			if v_3_0 == 380165415 {
				__t1 = v_1
				goto end_branch_1
			} else {

			}
		}
		{
			__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
		}
	end_branch_1:
		return __t1
	})}))}
}

func Call_Data_Ord_Max_monoidMax(dictBounded_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictBounded_0 gopurs_runtime.Value = dictBounded_0_loop
	_ = dictBounded_0
	// TAST (Let): semigroupMax1_1_0 shape=App(Var) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar a$scope16)])
	semigroupMax1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](Call_Data_Ord_Max_semigroupMax(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBounded_0, "Ord0"), gopurs_runtime.Value{})))
	_ = semigroupMax1_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer((&Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(semigroupMax1_1_0)}
	}), gopurs_runtime.RecordGet(dictBounded_0, "bottom")}))}
}

func Call_Data_Ord_Max_eqMax(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
	_ = dictEq_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0))}
}

func Call_Data_Ord_Max_ordMax(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
	_ = dictOrd_0
	// TAST (Let): eqMax1_1_0 shape=App(Other) bindingType=(ADT ["Data","Eq","Eq"] [(TypeVar a$scope8)])
	eqMax1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{}))
	_ = eqMax1_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ord_Ord[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(eqMax1_1_0)}
	}), gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), v_2, v1_3)
	})}))}
}
