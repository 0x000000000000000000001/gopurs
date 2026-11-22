package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Ord_Down_go__append gopurs_runtime.Value
var once_Data_Ord_Down_go__append sync.Once

func Get_Data_Ord_Down_go__append() gopurs_runtime.Value {
	once_Data_Ord_Down_go__append.Do(func() {
		cache_Data_Ord_Down_go__append = Get_Data_Semigroup_append__1694605992()
	})
	return cache_Data_Ord_Down_go__append
}

var cache_Data_Ord_Down_Down gopurs_runtime.Value
var once_Data_Ord_Down_Down sync.Once

func Get_Data_Ord_Down_Down() gopurs_runtime.Value {
	once_Data_Ord_Down_Down.Do(func() {
		cache_Data_Ord_Down_Down = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Ord_Down_Down(x_0_box)
		})
	})
	return cache_Data_Ord_Down_Down
}

var cache_Data_Ord_Down_showDown gopurs_runtime.Value
var once_Data_Ord_Down_showDown sync.Once

func Get_Data_Ord_Down_showDown() gopurs_runtime.Value {
	once_Data_Ord_Down_showDown.Do(func() {
		cache_Data_Ord_Down_showDown = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Ord_Down_showDown(dictShow_0_box)
		})
	})
	return cache_Data_Ord_Down_showDown
}

var cache_Data_Ord_Down_newtypeDown gopurs_runtime.Value
var once_Data_Ord_Down_newtypeDown sync.Once

func Get_Data_Ord_Down_newtypeDown() gopurs_runtime.Value {
	once_Data_Ord_Down_newtypeDown.Do(func() {
		cache_Data_Ord_Down_newtypeDown = gopurs_runtime.Value{Type: 9, IntVal: 3322196858, UnsafePtr: unsafe.Pointer((&Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{}
		})}))}
	})
	return cache_Data_Ord_Down_newtypeDown
}

var cache_Data_Ord_Down_eqDown gopurs_runtime.Value
var once_Data_Ord_Down_eqDown sync.Once

func Get_Data_Ord_Down_eqDown() gopurs_runtime.Value {
	once_Data_Ord_Down_eqDown.Do(func() {
		cache_Data_Ord_Down_eqDown = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Ord_Down_eqDown(dictEq_0_box)
		})
	})
	return cache_Data_Ord_Down_eqDown
}

var cache_Data_Ord_Down_ordDown gopurs_runtime.Value
var once_Data_Ord_Down_ordDown sync.Once

func Get_Data_Ord_Down_ordDown() gopurs_runtime.Value {
	once_Data_Ord_Down_ordDown.Do(func() {
		cache_Data_Ord_Down_ordDown = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Ord_Down_ordDown(dictOrd_0_box)
		})
	})
	return cache_Data_Ord_Down_ordDown
}

var cache_Data_Ord_Down_boundedDown gopurs_runtime.Value
var once_Data_Ord_Down_boundedDown sync.Once

func Get_Data_Ord_Down_boundedDown() gopurs_runtime.Value {
	once_Data_Ord_Down_boundedDown.Do(func() {
		cache_Data_Ord_Down_boundedDown = gopurs_runtime.Func(func(dictBounded_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Ord_Down_boundedDown(dictBounded_0_box)
		})
	})
	return cache_Data_Ord_Down_boundedDown
}

func Call_Data_Ord_Down_Down(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Data_Ord_Down_showDown(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
	_ = dictShow_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer((&Constructor_Data_Show_Show[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Str(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString("(Down ", gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), v_1).StrVal()), ")"))
	})}))}
}

func Call_Data_Ord_Down_eqDown(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
	_ = dictEq_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0))}
}

func Call_Data_Ord_Down_ordDown(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
	_ = dictOrd_0
	// TAST (Let): eqDown1_1_0 shape=App(Other) bindingType=(ADT ["Data","Eq","Eq"] [(TypeVar a$scope4)])
	eqDown1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{}))
	_ = eqDown1_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ord_Ord[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(eqDown1_1_0)}
	}), gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_4_1 shape=App(Other) bindingType=(ADT ["Data","Ordering","Ordering"] [])
		__local_var_4_1 := uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), v_2, v1_3).IntVal)
		_ = __local_var_4_1
		var __t2 gopurs_runtime.Value
		{
			if __local_var_4_1 == 380165415 {
				__t2 = gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}
				goto end_branch_2
			} else {

			}
		}
		{
			if __local_var_4_1 == 902936544 {
				__t2 = gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
				goto end_branch_2
			} else {

			}
		}
		{
			if __local_var_4_1 == 1527465420 {
				__t2 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
				goto end_branch_2
			} else {

			}
		}
		{
			__t2 = gopurs_runtime.Value{Type: 9, IntVal: int64(func() uint32 { panic("Failed pattern match") }()), UnsafePtr: nil}
		}
	end_branch_2:
		return __t2
	})}))}
}

func Call_Data_Ord_Down_boundedDown(dictBounded_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictBounded_0 gopurs_runtime.Value = dictBounded_0_loop
	_ = dictBounded_0
	// TAST (Let): ordDown1_1_0 shape=App(Var) bindingType=(ADT ["Data","Ord","Ord"] [(TypeVar a$scope14)])
	ordDown1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Call_Data_Ord_Down_ordDown(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBounded_0, "Ord0"), gopurs_runtime.Value{})))
	_ = ordDown1_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer((&Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(ordDown1_1_0)}
	}), gopurs_runtime.RecordGet(dictBounded_0, "top"), gopurs_runtime.RecordGet(dictBounded_0, "bottom")}))}
}
