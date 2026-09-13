package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_Pair gopurs_runtime.Value
var once_Main_Pair sync.Once

func Get_Main_Pair() gopurs_runtime.Value {
	once_Main_Pair.Do(func() {
		cache_Main_Pair = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 893478516, UnsafePtr: unsafe.Pointer((&Constructor_Main_Pair[gopurs_runtime.Value, gopurs_runtime.Value]{1, value0, value1}))}
			})
		})
	})
	return cache_Main_Pair
}

var cache_Main_Pair__3606837191 gopurs_runtime.Value
var once_Main_Pair__3606837191 sync.Once

func Get_Main_Pair__3606837191() gopurs_runtime.Value {
	once_Main_Pair__3606837191.Do(func() {
		cache_Main_Pair__3606837191 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 893478516, UnsafePtr: unsafe.Pointer(Rebox_Main_3835445094_2775928742(Call_Main_Pair__3606837191(__eta_norm_1_0_box.FloatVal(), __eta_norm_0_1_box.FloatVal())))}
		})
	})
	return cache_Main_Pair__3606837191
}

var cache_Main_eqPair gopurs_runtime.Value
var once_Main_eqPair sync.Once

func Get_Main_eqPair() gopurs_runtime.Value {
	once_Main_eqPair.Do(func() {
		cache_Main_eqPair = gopurs_runtime.Func2(func(dictEq_0_box gopurs_runtime.Value, dictEq1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_eqPair(dictEq_0_box, dictEq1_1_box)
		})
	})
	return cache_Main_eqPair
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Effect_Console_logShow(Rebox_Main_2735895690_1386611502(Rebox_Main_1386611502_2735895690(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showBoolean()))), gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Main_eqPair(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_687527510_3790796878(Rebox_Main_3790796878_687527510(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqNumber()))))}, gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_687527510_3790796878(Rebox_Main_3790796878_687527510(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqNumber()))))}), "eq"), gopurs_runtime.Value{Type: 9, IntVal: 893478516, UnsafePtr: unsafe.Pointer(Rebox_Main_3835445094_2775928742((&Constructor_Main_Pair[float64, float64]{1, 1.0, 2.0})))}, gopurs_runtime.Value{Type: 9, IntVal: 893478516, UnsafePtr: unsafe.Pointer(Rebox_Main_3835445094_2775928742((&Constructor_Main_Pair[float64, float64]{1, 1.0, 2.0})))}).IntVal) != (0))), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
		}))
	})
	return cache_Main_main
}

var cache_Main_ordPair gopurs_runtime.Value
var once_Main_ordPair sync.Once

func Get_Main_ordPair() gopurs_runtime.Value {
	once_Main_ordPair.Do(func() {
		cache_Main_ordPair = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_ordPair(dictOrd_0_box)
		})
	})
	return cache_Main_ordPair
}

type Constructor_Main_Pair[T_a any, T_b any] struct {
	Rc uint32
	V0 T_a
	V1 T_b
}

func Call_Main_Pair__3606837191(__eta_norm_1_0_loop float64, __eta_norm_0_1_loop float64) *Constructor_Main_Pair[float64, float64] {
Pair__3606837191:
	for {
		if false {
			continue Pair__3606837191
		}
		var __eta_norm_1_0 float64 = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_1 float64 = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		return (&Constructor_Main_Pair[float64, float64]{1, __eta_norm_1_0, __eta_norm_0_1})
	}
}

func Call_Main_eqPair(dictEq_0_loop gopurs_runtime.Value, dictEq1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
	_ = dictEq_0
	var dictEq1_1 gopurs_runtime.Value = dictEq1_1_loop
	_ = dictEq1_1
	return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1779257133_3790796878((&Constructor_Data_Eq_Eq[*Constructor_Main_Pair[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Main_Pair[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0, (*Constructor_Main_Pair[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq1_1, "eq"), (*Constructor_Main_Pair[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1, (*Constructor_Main_Pair[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V1).IntVal) != (0)))
	})})))}
}

func Call_Main_ordPair(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
	_ = dictOrd_0
	// TAST (Let): eqPair1_1_0 shape=App(Var) bindingType=Any
	eqPair1_1_0 := gopurs_runtime.Apply(Get_Main_eqPair(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{}))
	_ = eqPair1_1_0
	return gopurs_runtime.Func(func(dictOrd1_2 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): eqPair2_3_1 shape=App(Other) bindingType=(ADT ["Data","Eq","Eq"] [(ADT ["Main","Pair"] [(TypeVar a$scope2), (TypeVar b$scope3)])])
		eqPair2_3_1 := Rebox_Main_3790796878_1779257133(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](gopurs_runtime.Apply(eqPair1_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_2, "Eq0"), gopurs_runtime.Value{}))))
		_ = eqPair2_3_1
		return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Main_3741057677_4177771502((&Constructor_Data_Ord_Ord[*Constructor_Main_Pair[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1779257133_3790796878(eqPair2_3_1))}
		}), gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): v2_6_2 shape=App(Other) bindingType=(ADT ["Data","Ordering","Ordering"] [])
			v2_6_2 := uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (*Constructor_Main_Pair[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0, (*Constructor_Main_Pair[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V0).IntVal)
			_ = v2_6_2
			var __t3 uint32
			{
				if v2_6_2 == 902936544 {
					__t3 = uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd1_2, "compare"), (*Constructor_Main_Pair[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1, (*Constructor_Main_Pair[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V1).IntVal)
					goto end_branch_3
				} else {

				}
			}
			{
				__t3 = v2_6_2
			}
		end_branch_3:
			return gopurs_runtime.Value{Type: 9, IntVal: int64(__t3), UnsafePtr: nil}
		})})))}
	})
}

func Rebox_Main_1386611502_2735895690(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[bool] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[bool]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_1779257133_3790796878(in *Constructor_Data_Eq_Eq[*Constructor_Main_Pair[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_2735895690_1386611502(in *Constructor_Data_Show_Show[bool]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_3741057677_4177771502(in *Constructor_Data_Ord_Ord[*Constructor_Main_Pair[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Ord_Ord[gopurs_runtime.Value]{}
	out.V0 = in.V0
	out.V1 = in.V1
	return out
}

func Rebox_Main_3790796878_1779257133(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[*Constructor_Main_Pair[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Eq[*Constructor_Main_Pair[gopurs_runtime.Value, gopurs_runtime.Value]]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_3790796878_687527510(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[float64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Eq[float64]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_3835445094_2775928742(in *Constructor_Main_Pair[float64, float64]) *Constructor_Main_Pair[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_Pair[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Float(in.V0)
	out.V1 = gopurs_runtime.Float(in.V1)
	return out
}

func Rebox_Main_687527510_3790796878(in *Constructor_Data_Eq_Eq[float64]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}
