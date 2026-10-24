package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_Y gopurs_runtime.Value
var once_Main_Y sync.Once

func Get_Main_Y() gopurs_runtime.Value {
	once_Main_Y.Do(func() {
		cache_Main_Y = gopurs_runtime.Value{Type: 9, IntVal: 1714575428, UnsafePtr: unsafe.Pointer((*Constructor_Main_Z[gopurs_runtime.Value])(nil))}
	})
	return cache_Main_Y
}

var cache_Main_Z gopurs_runtime.Value
var once_Main_Z sync.Once

func Get_Main_Z() gopurs_runtime.Value {
	once_Main_Z.Do(func() {
		cache_Main_Z = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 1714575428, UnsafePtr: unsafe.Pointer((&Constructor_Main_Z[gopurs_runtime.Value]{1, value0, gopurs_runtime.CoerceToStruct[Constructor_Main_Z[gopurs_runtime.Value]](value1)}))}
			})
		})
	})
	return cache_Main_Z
}

var cache_Main_X gopurs_runtime.Value
var once_Main_X sync.Once

func Get_Main_X() gopurs_runtime.Value {
	once_Main_X.Do(func() {
		cache_Main_X = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return value0
		})
	})
	return cache_Main_X
}

var cache_Main_W gopurs_runtime.Value
var once_Main_W sync.Once

func Get_Main_W() gopurs_runtime.Value {
	once_Main_W.Do(func() {
		cache_Main_W = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				orig := Call_Main_W(func() struct {
					x int64
					y gopurs_runtime.Value
					/*gopurs:record-v1:2:0078,0079*/
				} {
					orig := x_0_box
					_ = orig
					clone := struct {
						x int64
						y gopurs_runtime.Value
						/*gopurs:record-v1:2:0078,0079*/
					}{}
					clone.x = gopurs_runtime.RecordGet(orig, "x").IntVal
					clone.y = gopurs_runtime.RecordGet(orig, "y")
					return clone
				}())
				_ = orig
				return gopurs_runtime.RecordDict2("x", "y", gopurs_runtime.Int(orig.x), orig.y)
			}()
		})
	})
	return cache_Main_W
}

var cache_Main_genericZ gopurs_runtime.Value
var once_Main_genericZ sync.Once

func Get_Main_genericZ() gopurs_runtime.Value {
	once_Main_genericZ.Do(func() {
		cache_Main_genericZ = gopurs_runtime.Value{Type: 9, IntVal: 1921946594, UnsafePtr: unsafe.Pointer((&Constructor_Data_Generic_Rep_Generic[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(gopurs_runtime.CoerceToStruct[Constructor_Data_Generic_Rep_Generic[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Main_genericZ()).V0, x_0)
		}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(gopurs_runtime.CoerceToStruct[Constructor_Data_Generic_Rep_Generic[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Main_genericZ()).V1, x_0)
		})}))}
	})
	return cache_Main_genericZ
}

var cache_Main_genericY gopurs_runtime.Value
var once_Main_genericY sync.Once

func Get_Main_genericY() gopurs_runtime.Value {
	once_Main_genericY.Do(func() {
		cache_Main_genericY = gopurs_runtime.Value{Type: 9, IntVal: 1921946594, UnsafePtr: unsafe.Pointer(Rebox_Main_2045408645_2818661616((&Constructor_Data_Generic_Rep_Generic[*Constructor_Main_Z[gopurs_runtime.Value], gopurs_runtime.Value]{1, gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t0 gopurs_runtime.Value
			{
				if x_0.Type == 9 && x_0.IntVal == 1714575428 && x_0.UnsafePtr == nil {
					__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3478632216, UnsafePtr: unsafe.Pointer((&Constructor_Data_Generic_Rep_Inl[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(1454898258), UnsafePtr: nil}}))}
					goto end_branch_0
				} else {

				}
			}
			{
				if x_0.Type == 9 && x_0.IntVal == 1714575428 && x_0.UnsafePtr != nil {
					__t0 = gopurs_runtime.Value{Type: 9, IntVal: 492034566, UnsafePtr: unsafe.Pointer((&Constructor_Data_Generic_Rep_Inr[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer((&Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Main_Z[gopurs_runtime.Value])(x_0.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1714575428, UnsafePtr: unsafe.Pointer((*Constructor_Main_Z[gopurs_runtime.Value])(x_0.UnsafePtr).V1)}}))}}))}
					goto end_branch_0
				} else {

				}
			}
			{
				__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_0:
			return __t0
		}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t1 gopurs_runtime.Value
			{
				if x_0.Type == 9 && x_0.IntVal == 3478632216 {
					__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1714575428, UnsafePtr: unsafe.Pointer((*Constructor_Main_Z[gopurs_runtime.Value])(nil))}
					goto end_branch_1
				} else {

				}
			}
			{
				if x_0.Type == 9 && x_0.IntVal == 492034566 {
					__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1714575428, UnsafePtr: unsafe.Pointer((&Constructor_Main_Z[gopurs_runtime.Value]{1, (*Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Generic_Rep_Inr[gopurs_runtime.Value, gopurs_runtime.Value])(x_0.UnsafePtr).V0.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Main_Z[gopurs_runtime.Value]]((*Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Generic_Rep_Inr[gopurs_runtime.Value, gopurs_runtime.Value])(x_0.UnsafePtr).V0.UnsafePtr).V1)}))}
					goto end_branch_1
				} else {

				}
			}
			{
				__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1714575428, UnsafePtr: unsafe.Pointer(func() *Constructor_Main_Z[gopurs_runtime.Value] { panic("Failed pattern match") }())}
			}
		end_branch_1:
			return __t1
		})})))}
	})
	return cache_Main_genericY
}

var cache_Main_genericX gopurs_runtime.Value
var once_Main_genericX sync.Once

func Get_Main_genericX() gopurs_runtime.Value {
	once_Main_genericX.Do(func() {
		cache_Main_genericX = gopurs_runtime.Value{Type: 9, IntVal: 1921946594, UnsafePtr: unsafe.Pointer((&Constructor_Data_Generic_Rep_Generic[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return x_0
		}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return x_0
		})}))}
	})
	return cache_Main_genericX
}

var cache_Main_genericW gopurs_runtime.Value
var once_Main_genericW sync.Once

func Get_Main_genericW() gopurs_runtime.Value {
	once_Main_genericW.Do(func() {
		cache_Main_genericW = gopurs_runtime.Value{Type: 9, IntVal: 1921946594, UnsafePtr: unsafe.Pointer(Rebox_Main_495020111_2818661616((&Constructor_Data_Generic_Rep_Generic[struct {
			x int64
			y gopurs_runtime.Value
			/*gopurs:record-v1:2:0078,0079*/
		}, struct {
			x int64
			y string
			/*gopurs:record-v1:2:0078,0079*/
		}]{1, gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return x_0
		}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return x_0
		})})))}
	})
	return cache_Main_genericW
}

var cache_Main_eqZ gopurs_runtime.Value
var once_Main_eqZ sync.Once

func Get_Main_eqZ() gopurs_runtime.Value {
	once_Main_eqZ.Do(func() {
		cache_Main_eqZ = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(true)
		})}))}
	})
	return cache_Main_eqZ
}

var cache_Main_eqY gopurs_runtime.Value
var once_Main_eqY sync.Once

func Get_Main_eqY() gopurs_runtime.Value {
	once_Main_eqY.Do(func() {
		cache_Main_eqY = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_eqY(dictEq_0_box)
		})
	})
	return cache_Main_eqY
}

var cache_Main_eqX gopurs_runtime.Value
var once_Main_eqX sync.Once

func Get_Main_eqX() gopurs_runtime.Value {
	once_Main_eqX.Do(func() {
		cache_Main_eqX = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_eqX(dictEq_0_box)
		})
	})
	return cache_Main_eqX
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=Any
			__local_var_0_0 := Call_Effect_Console_logShow(Rebox_Main_2735895690_1386611502(Rebox_Main_1386611502_2735895690(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showBoolean()))), gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Main_eqX(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), "eq"), gopurs_runtime.Int(int64(0)), gopurs_runtime.Int(int64(1))).IntVal) != (0)))
			_ = __local_var_0_0
			_dollar___unused_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = _dollar___unused_1_1
			_dollar___unused_2_2 := gopurs_runtime.Apply(Call_Effect_Console_logShow(Rebox_Main_2735895690_1386611502(Rebox_Main_1386611502_2735895690(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showBoolean()))), gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Main_eqX(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), "eq"), gopurs_runtime.Int(int64(1)), gopurs_runtime.Int(int64(1))).IntVal) != (0))), gopurs_runtime.Value{})
			_ = _dollar___unused_2_2
			_dollar___unused_3_3 := gopurs_runtime.Apply(Call_Effect_Console_logShow(Rebox_Main_2735895690_1386611502(Rebox_Main_1386611502_2735895690(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showBoolean()))), gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Main_eqY(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), "eq"), gopurs_runtime.Value{Type: 9, IntVal: 1714575428, UnsafePtr: unsafe.Pointer(Rebox_Main_857262859_4105062992((&Constructor_Main_Z[int64]{1, int64(1), (*Constructor_Main_Z[int64])(nil)})))}, gopurs_runtime.Value{Type: 9, IntVal: 1714575428, UnsafePtr: unsafe.Pointer(Rebox_Main_857262859_4105062992((&Constructor_Main_Z[int64]{1, int64(1), (*Constructor_Main_Z[int64])(nil)})))}).IntVal) != (0))), gopurs_runtime.Value{})
			_ = _dollar___unused_3_3
			_dollar___unused_4_4 := gopurs_runtime.Apply(Call_Effect_Console_logShow(Rebox_Main_2735895690_1386611502(Rebox_Main_1386611502_2735895690(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showBoolean()))), gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Main_eqY(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), "eq"), gopurs_runtime.Value{Type: 9, IntVal: 1714575428, UnsafePtr: unsafe.Pointer(Rebox_Main_857262859_4105062992((&Constructor_Main_Z[int64]{1, int64(1), (*Constructor_Main_Z[int64])(nil)})))}, gopurs_runtime.Value{Type: 9, IntVal: 1714575428, UnsafePtr: unsafe.Pointer(Rebox_Main_857262859_4105062992((*Constructor_Main_Z[int64])(nil)))}).IntVal) != (0))), gopurs_runtime.Value{})
			_ = _dollar___unused_4_4
			_dollar___unused_5_5 := gopurs_runtime.Apply(Call_Effect_Console_logShow(Rebox_Main_2735895690_1386611502(Rebox_Main_1386611502_2735895690(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showBoolean()))), gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Main_eqY(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main_eqZ()))}), "eq"), gopurs_runtime.Value{Type: 9, IntVal: 1714575428, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Main_Z[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1714575428, UnsafePtr: unsafe.Pointer((*Constructor_Main_Z[gopurs_runtime.Value])(nil))}))}, gopurs_runtime.Value{Type: 9, IntVal: 1714575428, UnsafePtr: unsafe.Pointer((*Constructor_Main_Z[gopurs_runtime.Value])(nil))}).IntVal) != (0))), gopurs_runtime.Value{})
			_ = _dollar___unused_5_5
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

type Constructor_Main_Y[T_a any] struct {
	Rc uint32
}

type Constructor_Main_Z[T_a any] struct {
	Rc uint32
	V0 T_a
	V1 *Constructor_Main_Z[T_a]
}

type Constructor_Main_X[T_a any] struct {
	Rc uint32
	V0 T_a
}

func Call_Main_W(x_0_loop struct {
	x int64
	y gopurs_runtime.Value
	/*gopurs:record-v1:2:0078,0079*/
}) struct {
	x int64
	y gopurs_runtime.Value
	/*gopurs:record-v1:2:0078,0079*/
} {
	var x_0 struct {
		x int64
		y gopurs_runtime.Value
		/*gopurs:record-v1:2:0078,0079*/
	} = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_eqY(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
eqY:
	for {
		if false {
			continue eqY
		}
		var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
		_ = dictEq_0
		// TAST (Let): genericEqProduct_1_0 shape=App(Var) bindingType=Any
		genericEqProduct_1_0 := gopurs_runtime.Apply(Get_Data_Eq_Generic_genericEqProduct(), Call_Data_Eq_Generic_genericEqArgument(dictEq_0))
		_ = genericEqProduct_1_0
		return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_3989972411_3790796878((&Constructor_Data_Eq_Eq[*Constructor_Main_Z[gopurs_runtime.Value]]{1, gopurs_runtime.Func2(func(xs_2 gopurs_runtime.Value, ys_3 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t1 gopurs_runtime.Value
			{
				if xs_2.Type == 9 && xs_2.IntVal == 1714575428 && xs_2.UnsafePtr == nil {
					__t1 = gopurs_runtime.Value{Type: 9, IntVal: 3478632216, UnsafePtr: unsafe.Pointer(Rebox_Main_4264667741_1323331594((&Constructor_Data_Generic_Rep_Inl[uint32, *Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, *Constructor_Main_Z[gopurs_runtime.Value]]]{1, 1454898258})))}
					goto end_branch_1
				} else {

				}
			}
			{
				if xs_2.Type == 9 && xs_2.IntVal == 1714575428 && xs_2.UnsafePtr != nil {
					__t1 = gopurs_runtime.Value{Type: 9, IntVal: 492034566, UnsafePtr: unsafe.Pointer(Rebox_Main_1784434947_2687169876((&Constructor_Data_Generic_Rep_Inr[uint32, *Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, *Constructor_Main_Z[gopurs_runtime.Value]]]{1, (&Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, *Constructor_Main_Z[gopurs_runtime.Value]]{1, (*Constructor_Main_Z[gopurs_runtime.Value])(xs_2.UnsafePtr).V0, (*Constructor_Main_Z[gopurs_runtime.Value])(xs_2.UnsafePtr).V1})})))}
					goto end_branch_1
				} else {

				}
			}
			{
				__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_1:
			var __t2 gopurs_runtime.Value
			{
				if ys_3.Type == 9 && ys_3.IntVal == 1714575428 && ys_3.UnsafePtr == nil {
					__t2 = gopurs_runtime.Value{Type: 9, IntVal: 3478632216, UnsafePtr: unsafe.Pointer(Rebox_Main_4264667741_1323331594((&Constructor_Data_Generic_Rep_Inl[uint32, *Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, *Constructor_Main_Z[gopurs_runtime.Value]]]{1, 1454898258})))}
					goto end_branch_2
				} else {

				}
			}
			{
				if ys_3.Type == 9 && ys_3.IntVal == 1714575428 && ys_3.UnsafePtr != nil {
					__t2 = gopurs_runtime.Value{Type: 9, IntVal: 492034566, UnsafePtr: unsafe.Pointer(Rebox_Main_1784434947_2687169876((&Constructor_Data_Generic_Rep_Inr[uint32, *Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, *Constructor_Main_Z[gopurs_runtime.Value]]]{1, (&Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, *Constructor_Main_Z[gopurs_runtime.Value]]{1, (*Constructor_Main_Z[gopurs_runtime.Value])(ys_3.UnsafePtr).V0, (*Constructor_Main_Z[gopurs_runtime.Value])(ys_3.UnsafePtr).V1})})))}
					goto end_branch_2
				} else {

				}
			}
			{
				__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_2:
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Data_Eq_Generic_genericEqSum(Call_Data_Eq_Generic_genericEqConstructor(gopurs_runtime.Value{Type: 9, IntVal: 106035173, UnsafePtr: unsafe.Pointer(Rebox_Main_4167548092_1498596945(Rebox_Main_1498596945_4167548092(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Generic_GenericEq[gopurs_runtime.Value]](Get_Data_Eq_Generic_genericEqNoArguments()))))}), Call_Data_Eq_Generic_genericEqConstructor(gopurs_runtime.Apply(genericEqProduct_1_0, Call_Data_Eq_Generic_genericEqArgument(Call_Main_eqY(dictEq_0))))), "genericEq'"), __t1, __t2)
		})})))}
	}
}

func Call_Main_eqX(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
	_ = dictEq_0
	// TAST (Let): genericEqConstructor_1_0 shape=App(Var) bindingType=(ADT ["Data","Eq","Generic","GenericEq"] [Any])
	genericEqConstructor_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Generic_GenericEq[gopurs_runtime.Value]](Call_Data_Eq_Generic_genericEqConstructor(Call_Data_Eq_Generic_genericEqArgument(dictEq_0)))
	_ = genericEqConstructor_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(xs_2 gopurs_runtime.Value, ys_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(genericEqConstructor_1_0.V0, xs_2, ys_3)
	})}))}
}

func Rebox_Main_1053099733_3790796878(in *Constructor_Data_Eq_Eq[int64]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_1386611502_2735895690(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[bool] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[bool])(unsafe.Pointer(in))
}

func Rebox_Main_1498596945_4167548092(in *Constructor_Data_Eq_Generic_GenericEq[gopurs_runtime.Value]) *Constructor_Data_Eq_Generic_GenericEq[uint32] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Generic_GenericEq[uint32])(unsafe.Pointer(in))
}

func Rebox_Main_1784434947_2687169876(in *Constructor_Data_Generic_Rep_Inr[uint32, *Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, *Constructor_Main_Z[gopurs_runtime.Value]]]) *Constructor_Data_Generic_Rep_Inr[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Generic_Rep_Inr[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer(Rebox_Main_966244383_372708202(in.V0))}
	return out
}

func Rebox_Main_2045408645_2818661616(in *Constructor_Data_Generic_Rep_Generic[*Constructor_Main_Z[gopurs_runtime.Value], gopurs_runtime.Value]) *Constructor_Data_Generic_Rep_Generic[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Generic_Rep_Generic[gopurs_runtime.Value, gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_2735895690_1386611502(in *Constructor_Data_Show_Show[bool]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_3790796878_1053099733(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[int64] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[int64])(unsafe.Pointer(in))
}

func Rebox_Main_3989972411_3790796878(in *Constructor_Data_Eq_Eq[*Constructor_Main_Z[gopurs_runtime.Value]]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_4167548092_1498596945(in *Constructor_Data_Eq_Generic_GenericEq[uint32]) *Constructor_Data_Eq_Generic_GenericEq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Generic_GenericEq[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_4264667741_1323331594(in *Constructor_Data_Generic_Rep_Inl[uint32, *Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, *Constructor_Main_Z[gopurs_runtime.Value]]]) *Constructor_Data_Generic_Rep_Inl[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Generic_Rep_Inl[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Value{Type: 9, IntVal: int64(in.V0), UnsafePtr: nil}
	return out
}

func Rebox_Main_495020111_2818661616(in *Constructor_Data_Generic_Rep_Generic[struct {
	x int64
	y gopurs_runtime.Value
	/*gopurs:record-v1:2:0078,0079*/
}, struct {
	x int64
	y string
	/*gopurs:record-v1:2:0078,0079*/
}]) *Constructor_Data_Generic_Rep_Generic[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Generic_Rep_Generic[gopurs_runtime.Value, gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_857262859_4105062992(in *Constructor_Main_Z[int64]) *Constructor_Main_Z[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_Z[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Int(in.V0)
	out.V1 = Rebox_Main_857262859_4105062992(in.V1)
	return out
}

func Rebox_Main_966244383_372708202(in *Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, *Constructor_Main_Z[gopurs_runtime.Value]]) *Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = in.V0
	out.V1 = gopurs_runtime.Value{Type: 9, IntVal: 1714575428, UnsafePtr: unsafe.Pointer(in.V1)}
	return out
}
