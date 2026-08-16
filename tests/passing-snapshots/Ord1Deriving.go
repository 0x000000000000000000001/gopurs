package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_Product gopurs_runtime.Value
var once_Main_Product sync.Once

func Get_Main_Product() gopurs_runtime.Value {
	once_Main_Product.Do(func() {
		cache_Main_Product = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 2670017141, UnsafePtr: unsafe.Pointer((&Constructor_Main_Product{1, value0, value1}))}
			})
		})
	})
	return cache_Main_Product
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

var cache_Main_eqMu gopurs_runtime.Value
var once_Main_eqMu sync.Once

func Get_Main_eqMu() gopurs_runtime.Value {
	once_Main_eqMu.Do(func() {
		cache_Main_eqMu = gopurs_runtime.Func2(func(dictEq_0_box gopurs_runtime.Value, dictEq1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_eqMu(dictEq_0_box, dictEq1_1_box)
		})
	})
	return cache_Main_eqMu
}

var cache_Main_ordMu gopurs_runtime.Value
var once_Main_ordMu sync.Once

func Get_Main_ordMu() gopurs_runtime.Value {
	once_Main_ordMu.Do(func() {
		cache_Main_ordMu = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_ordMu(dictOrd_0_box)
		})
	})
	return cache_Main_ordMu
}

var cache_Main_eq1Mu gopurs_runtime.Value
var once_Main_eq1Mu sync.Once

func Get_Main_eq1Mu() gopurs_runtime.Value {
	once_Main_eq1Mu.Do(func() {
		cache_Main_eq1Mu = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_eq1Mu(dictEq_0_box)
		})
	})
	return cache_Main_eq1Mu
}

var cache_Main_ord1Mu gopurs_runtime.Value
var once_Main_ord1Mu sync.Once

func Get_Main_ord1Mu() gopurs_runtime.Value {
	once_Main_ord1Mu.Do(func() {
		cache_Main_ord1Mu = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_ord1Mu(dictOrd_0_box)
		})
	})
	return cache_Main_ord1Mu
}

type Constructor_Main_Product struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}

func Call_Main_eqMu(dictEq_0_loop gopurs_runtime.Value, dictEq1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
	_ = dictEq_0
	var dictEq1_1 gopurs_runtime.Value = dictEq1_1_loop
	_ = dictEq1_1
	return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq{1, gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Main_Product)(x_2.UnsafePtr).V0, (*Constructor_Main_Product)(y_3.UnsafePtr).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq1_1, "eq"), (*Constructor_Main_Product)(x_2.UnsafePtr).V1, (*Constructor_Main_Product)(y_3.UnsafePtr).V1).IntVal) != (0)))
		})
	})}))}
}

func Call_Main_ordMu(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
	_ = dictOrd_0
	// TAST (Let): __local_var_1_0 shape=App(Other) expectedFromAst=gopurs_runtime.Value actual=gopurs_runtime.Value bindingType=Any
	__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})
	_ = __local_var_1_0
	return gopurs_runtime.Func(func(dictOrd1_2 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_3_2 shape=App(Other) expectedFromAst=gopurs_runtime.Value actual=gopurs_runtime.Value bindingType=Any
		__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_2, "Eq0"), gopurs_runtime.Value{})
		_ = __local_var_3_2
		// TAST (Let): eqMu2_3_1 shape=Let(LitRecord) expectedFromAst=*Constructor_Data_Eq_Eq actual=*Constructor_Data_Eq_Eq bindingType=(ADT ["Data","Eq","Eq"] [(ADT ["Main","Product"] [(TypeVar a), (TypeVar b)])])
		eqMu2_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(y_5 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "eq"), (*Constructor_Main_Product)(x_4.UnsafePtr).V0, (*Constructor_Main_Product)(y_5.UnsafePtr).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_2, "eq"), (*Constructor_Main_Product)(x_4.UnsafePtr).V1, (*Constructor_Main_Product)(y_5.UnsafePtr).V1).IntVal) != (0)))
			})
		})))
		_ = eqMu2_3_1
		return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ord_Ord{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(eqMu2_3_1)}
		}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(y_5 gopurs_runtime.Value) gopurs_runtime.Value {
				// TAST (Let): v_6_3 shape=App(Other) expectedFromAst=uint32 actual=uint32 bindingType=(ADT ["Data","Ordering","Ordering"] [])
				v_6_3 := uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (*Constructor_Main_Product)(x_4.UnsafePtr).V0, (*Constructor_Main_Product)(y_5.UnsafePtr).V0).IntVal)
				_ = v_6_3
				var __t4 uint32
				{
					if v_6_3 == 1527465420 {
						__t4 = 1527465420
						goto end_branch_4
					} else {

					}
				}
				{
					if v_6_3 == 380165415 {
						__t4 = 380165415
						goto end_branch_4
					} else {

					}
				}
				{
					__t4 = uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd1_2, "compare"), (*Constructor_Main_Product)(x_4.UnsafePtr).V1, (*Constructor_Main_Product)(y_5.UnsafePtr).V1).IntVal)
				}
			end_branch_4:
				return gopurs_runtime.Value{Type: 9, IntVal: int64(__t4), UnsafePtr: nil}
			})
		})}))}
	})
}

func Call_Main_eq1Mu(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
	_ = dictEq_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq1{1, gopurs_runtime.Func(func(dictEq1_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Main_Product)(x_2.UnsafePtr).V0, (*Constructor_Main_Product)(y_3.UnsafePtr).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq1_1, "eq"), (*Constructor_Main_Product)(x_2.UnsafePtr).V1, (*Constructor_Main_Product)(y_3.UnsafePtr).V1).IntVal) != (0)))
			})
		})
	})}))}
}

func Call_Main_ord1Mu(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
	_ = dictOrd_0
	// TAST (Let): __local_var_1_1 shape=App(Other) expectedFromAst=gopurs_runtime.Value actual=gopurs_runtime.Value bindingType=Any
	__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})
	_ = __local_var_1_1
	// TAST (Let): eq1Mu1_1_0 shape=Let(LitRecord) expectedFromAst=*Constructor_Data_Eq_Eq1 actual=*Constructor_Data_Eq_Eq1 bindingType=(ADT ["Data","Eq","Eq1"] [(ADT ["Main","Product"] [(TypeVar a)])])
	eq1Mu1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1](gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq1_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(y_4 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "eq"), (*Constructor_Main_Product)(x_3.UnsafePtr).V0, (*Constructor_Main_Product)(y_4.UnsafePtr).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq1_2, "eq"), (*Constructor_Main_Product)(x_3.UnsafePtr).V1, (*Constructor_Main_Product)(y_4.UnsafePtr).V1).IntVal) != (0)))
			})
		})
	})))
	_ = eq1Mu1_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1632188299, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ord_Ord1{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(eq1Mu1_1_0)}
	}), gopurs_runtime.Func(func(dictOrd1_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(y_4 gopurs_runtime.Value) gopurs_runtime.Value {
				// TAST (Let): v_5_2 shape=App(Other) expectedFromAst=uint32 actual=uint32 bindingType=(ADT ["Data","Ordering","Ordering"] [])
				v_5_2 := uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (*Constructor_Main_Product)(x_3.UnsafePtr).V0, (*Constructor_Main_Product)(y_4.UnsafePtr).V0).IntVal)
				_ = v_5_2
				var __t3 uint32
				{
					if v_5_2 == 1527465420 {
						__t3 = 1527465420
						goto end_branch_3
					} else {

					}
				}
				{
					if v_5_2 == 380165415 {
						__t3 = 380165415
						goto end_branch_3
					} else {

					}
				}
				{
					__t3 = uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd1_2, "compare"), (*Constructor_Main_Product)(x_3.UnsafePtr).V1, (*Constructor_Main_Product)(y_4.UnsafePtr).V1).IntVal)
				}
			end_branch_3:
				return gopurs_runtime.Value{Type: 9, IntVal: int64(__t3), UnsafePtr: nil}
			})
		})
	})}))}
}
