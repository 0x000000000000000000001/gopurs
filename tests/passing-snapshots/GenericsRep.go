package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_genericEqSum gopurs_runtime.Value
var once_Main_genericEqSum sync.Once

func Get_Main_genericEqSum() gopurs_runtime.Value {
	once_Main_genericEqSum.Do(func() {
		cache_Main_genericEqSum = gopurs_runtime.Apply(Get_Data_Eq_Generic_genericEqSum(), Call_Data_Eq_Generic_genericEqConstructor(gopurs_runtime.Value{Type: 9, IntVal: 106035173, UnsafePtr: unsafe.Pointer(Rebox_Main_4167548092_1498596945(Rebox_Main_1498596945_4167548092(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Generic_GenericEq[gopurs_runtime.Value]](Get_Data_Eq_Generic_genericEqNoArguments()))))}))
	})
	return cache_Main_genericEqSum
}

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

var cache_Main_Z__643203383 gopurs_runtime.Value
var once_Main_Z__643203383 sync.Once

func Get_Main_Z__643203383() gopurs_runtime.Value {
	once_Main_Z__643203383.Do(func() {
		cache_Main_Z__643203383 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1714575428, UnsafePtr: unsafe.Pointer(Rebox_Main_857262859_4105062992(Call_Main_Z__643203383(__eta_norm_1_0_box.IntVal, Rebox_Main_4105062992_857262859(gopurs_runtime.CoerceToStruct[Constructor_Main_Z[gopurs_runtime.Value]](__eta_norm_0_unused_1_box)))))}
		})
	})
	return cache_Main_Z__643203383
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

var cache_Main_X__3655787559 gopurs_runtime.Value
var once_Main_X__3655787559 sync.Once

func Get_Main_X__3655787559() gopurs_runtime.Value {
	once_Main_X__3655787559.Do(func() {
		cache_Main_X__3655787559 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_X__3655787559(__eta_norm_0_0_box.IntVal)
		})
	})
	return cache_Main_X__3655787559
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
				} {
					orig := x_0_box
					_ = orig
					clone := struct {
						x int64
						y gopurs_runtime.Value
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
			var __t2 gopurs_runtime.Value
			{
				var __t_tag_0 *Constructor_Main_Z[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Main_Z[gopurs_runtime.Value]](x_0)
				_ = __t_tag_0
				if __t_tag_0 == nil {
					__t2 = gopurs_runtime.Value{Type: 9, IntVal: 3478632216, UnsafePtr: unsafe.Pointer(Rebox_Main_4264667741_1323331594((&Constructor_Data_Generic_Rep_Inl[uint32, *Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, *Constructor_Main_Z[gopurs_runtime.Value]]]{1, 1454898258})))}
					goto end_branch_2
				} else {

				}
			}
			{
				var __t_tag_1 *Constructor_Main_Z[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Main_Z[gopurs_runtime.Value]](x_0)
				_ = __t_tag_1
				if __t_tag_1 != nil {
					__t2 = gopurs_runtime.Value{Type: 9, IntVal: 492034566, UnsafePtr: unsafe.Pointer(Rebox_Main_1784434947_2687169876((&Constructor_Data_Generic_Rep_Inr[uint32, *Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, *Constructor_Main_Z[gopurs_runtime.Value]]]{1, Rebox_Main_372708202_966244383((&Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Main_Z[gopurs_runtime.Value])(x_0.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1714575428, UnsafePtr: unsafe.Pointer((*Constructor_Main_Z[gopurs_runtime.Value])(x_0.UnsafePtr).V1)}}))})))}
					goto end_branch_2
				} else {

				}
			}
			{
				__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_2:
			return __t2
		}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t3 *Constructor_Main_Z[gopurs_runtime.Value]
			{
				if x_0.Type == 9 && x_0.IntVal == 3478632216 {
					__t3 = gopurs_runtime.CoerceToStruct[Constructor_Main_Z[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1714575428, UnsafePtr: unsafe.Pointer((*Constructor_Main_Z[gopurs_runtime.Value])(nil))})
					goto end_branch_3
				} else {

				}
			}
			{
				if x_0.Type == 9 && x_0.IntVal == 492034566 {
					__t3 = (&Constructor_Main_Z[gopurs_runtime.Value]{1, (*Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Generic_Rep_Inr[gopurs_runtime.Value, gopurs_runtime.Value])(x_0.UnsafePtr).V0.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Main_Z[gopurs_runtime.Value]]((*Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Generic_Rep_Inr[gopurs_runtime.Value, gopurs_runtime.Value])(x_0.UnsafePtr).V0.UnsafePtr).V1)})
					goto end_branch_3
				} else {

				}
			}
			{
				__t3 = func() *Constructor_Main_Z[gopurs_runtime.Value] { panic("Failed pattern match") }()
			}
		end_branch_3:
			return gopurs_runtime.Value{Type: 9, IntVal: 1714575428, UnsafePtr: unsafe.Pointer(__t3)}
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
		}, struct {
			x int64
			y string
		}]{1, gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				orig := func() struct {
					x int64
					y string
				} {
					orig := func() gopurs_runtime.Value {
						orig := func() struct {
							x int64
							y gopurs_runtime.Value
						} {
							orig := x_0
							_ = orig
							clone := struct {
								x int64
								y gopurs_runtime.Value
							}{}
							clone.x = gopurs_runtime.RecordGet(orig, "x").IntVal
							clone.y = gopurs_runtime.RecordGet(orig, "y")
							return clone
						}()
						_ = orig
						return gopurs_runtime.RecordDict2("x", "y", gopurs_runtime.Int(orig.x), orig.y)
					}()
					_ = orig
					clone := struct {
						x int64
						y string
					}{}
					clone.x = gopurs_runtime.RecordGet(orig, "x").IntVal
					clone.y = gopurs_runtime.RecordGet(orig, "y").StrVal()
					return clone
				}()
				_ = orig
				return gopurs_runtime.RecordDict2("x", "y", gopurs_runtime.Int(orig.x), gopurs_runtime.Str(orig.y))
			}()
		}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				orig := func() struct {
					x int64
					y gopurs_runtime.Value
				} {
					orig := func() gopurs_runtime.Value {
						orig := func() struct {
							x int64
							y string
						} {
							orig := x_0
							_ = orig
							clone := struct {
								x int64
								y string
							}{}
							clone.x = gopurs_runtime.RecordGet(orig, "x").IntVal
							clone.y = gopurs_runtime.RecordGet(orig, "y").StrVal()
							return clone
						}()
						_ = orig
						return gopurs_runtime.RecordDict2("x", "y", gopurs_runtime.Int(orig.x), gopurs_runtime.Str(orig.y))
					}()
					_ = orig
					clone := struct {
						x int64
						y gopurs_runtime.Value
					}{}
					clone.x = gopurs_runtime.RecordGet(orig, "x").IntVal
					clone.y = gopurs_runtime.RecordGet(orig, "y")
					return clone
				}()
				_ = orig
				return gopurs_runtime.RecordDict2("x", "y", gopurs_runtime.Int(orig.x), orig.y)
			}()
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

var cache_Main_eqY1 gopurs_runtime.Value
var once_Main_eqY1 sync.Once

func Get_Main_eqY1() gopurs_runtime.Value {
	once_Main_eqY1.Do(func() {
		cache_Main_eqY1 = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1121961536_3790796878(Rebox_Main_3790796878_1121961536(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Main_eqY(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})))))}
	})
	return cache_Main_eqY1
}

var cache_Main_eqY2 gopurs_runtime.Value
var once_Main_eqY2 sync.Once

func Get_Main_eqY2() gopurs_runtime.Value {
	once_Main_eqY2.Do(func() {
		cache_Main_eqY2 = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_3989972411_3790796878(Rebox_Main_3790796878_3989972411(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Main_eqY(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main_eqZ()))})))))}
	})
	return cache_Main_eqY2
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

var cache_Main_eqX1 gopurs_runtime.Value
var once_Main_eqX1 sync.Once

func Get_Main_eqX1() gopurs_runtime.Value {
	once_Main_eqX1.Do(func() {
		cache_Main_eqX1 = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Main_eqX(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})))}
	})
	return cache_Main_eqX1
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = func() gopurs_runtime.Value {
			var __t0 string
			{
				if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Main_eqX(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), "eq"), gopurs_runtime.Int(int64(0)), gopurs_runtime.Int(int64(1))).IntVal) != (0) {
					__t0 = "true"
					goto end_branch_0
				} else {

				}
			}
			{
				__t0 = "false"
			}
		end_branch_0:
			return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(__t0)), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t1 string
				{
					if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Main_eqX(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), "eq"), gopurs_runtime.Int(int64(1)), gopurs_runtime.Int(int64(1))).IntVal) != (0) {
						__t1 = "true"
						goto end_branch_1
					} else {

					}
				}
				{
					__t1 = "false"
				}
			end_branch_1:
				return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(__t1)), gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
					var __t2 string
					{
						if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Main_eqY(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), "eq"), gopurs_runtime.Value{Type: 9, IntVal: 1714575428, UnsafePtr: unsafe.Pointer(Rebox_Main_857262859_4105062992((&Constructor_Main_Z[int64]{1, int64(1), (*Constructor_Main_Z[int64])(nil)})))}, gopurs_runtime.Value{Type: 9, IntVal: 1714575428, UnsafePtr: unsafe.Pointer(Rebox_Main_857262859_4105062992((&Constructor_Main_Z[int64]{1, int64(1), (*Constructor_Main_Z[int64])(nil)})))}).IntVal) != (0) {
							__t2 = "true"
							goto end_branch_2
						} else {

						}
					}
					{
						__t2 = "false"
					}
				end_branch_2:
					return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(__t2)), gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
						var __t3 string
						{
							if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Main_eqY(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), "eq"), gopurs_runtime.Value{Type: 9, IntVal: 1714575428, UnsafePtr: unsafe.Pointer(Rebox_Main_857262859_4105062992((&Constructor_Main_Z[int64]{1, int64(1), (*Constructor_Main_Z[int64])(nil)})))}, gopurs_runtime.Value{Type: 9, IntVal: 1714575428, UnsafePtr: unsafe.Pointer(Rebox_Main_857262859_4105062992((*Constructor_Main_Z[int64])(nil)))}).IntVal) != (0) {
								__t3 = "true"
								goto end_branch_3
							} else {

							}
						}
						{
							__t3 = "false"
						}
					end_branch_3:
						return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(__t3)), gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
							var __t4 string
							{
								if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Main_eqY(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main_eqZ()))}), "eq"), gopurs_runtime.Value{Type: 9, IntVal: 1714575428, UnsafePtr: unsafe.Pointer((*Constructor_Main_Z[gopurs_runtime.Value])(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1714575428, UnsafePtr: unsafe.Pointer((*Constructor_Main_Z[gopurs_runtime.Value])(nil))}).IntVal) != (0) {
									__t4 = "true"
									goto end_branch_4
								} else {

								}
							}
							{
								__t4 = "false"
							}
						end_branch_4:
							return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(__t4)), gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
							}))
						}))
					}))
				}))
			}))
		}()
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

func Call_Main_Z__643203383(__eta_norm_1_0_loop int64, __eta_norm_0_unused_1_loop *Constructor_Main_Z[int64]) *Constructor_Main_Z[int64] {
Z__643203383:
	for {
		if false {
			continue Z__643203383
		}
		var __eta_norm_1_0 int64 = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_unused_1 *Constructor_Main_Z[int64] = __eta_norm_0_unused_1_loop
		_ = __eta_norm_0_unused_1
		return (&Constructor_Main_Z[int64]{1, __eta_norm_1_0, (*Constructor_Main_Z[int64])(nil)})
	}
}

func Call_Main_X__3655787559(__eta_norm_0_0_loop int64) gopurs_runtime.Value {
X__3655787559:
	for {
		if false {
			continue X__3655787559
		}
		var __eta_norm_0_0 int64 = __eta_norm_0_0_loop
		_ = __eta_norm_0_0
		return gopurs_runtime.Int(__eta_norm_0_0)
	}
}

func Call_Main_W(x_0_loop struct {
	x int64
	y gopurs_runtime.Value
}) struct {
	x int64
	y gopurs_runtime.Value
} {
	var x_0 struct {
		x int64
		y gopurs_runtime.Value
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
			var __t3 gopurs_runtime.Value
			{
				var __t_tag_1 *Constructor_Main_Z[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Main_Z[gopurs_runtime.Value]](xs_2)
				_ = __t_tag_1
				if __t_tag_1 == nil {
					__t3 = gopurs_runtime.Value{Type: 9, IntVal: 3478632216, UnsafePtr: unsafe.Pointer(Rebox_Main_4264667741_1323331594((&Constructor_Data_Generic_Rep_Inl[uint32, *Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, *Constructor_Main_Z[gopurs_runtime.Value]]]{1, 1454898258})))}
					goto end_branch_3
				} else {

				}
			}
			{
				var __t_tag_2 *Constructor_Main_Z[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Main_Z[gopurs_runtime.Value]](xs_2)
				_ = __t_tag_2
				if __t_tag_2 != nil {
					__t3 = gopurs_runtime.Value{Type: 9, IntVal: 492034566, UnsafePtr: unsafe.Pointer(Rebox_Main_1784434947_2687169876((&Constructor_Data_Generic_Rep_Inr[uint32, *Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, *Constructor_Main_Z[gopurs_runtime.Value]]]{1, Rebox_Main_372708202_966244383((&Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Main_Z[gopurs_runtime.Value])(xs_2.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1714575428, UnsafePtr: unsafe.Pointer((*Constructor_Main_Z[gopurs_runtime.Value])(xs_2.UnsafePtr).V1)}}))})))}
					goto end_branch_3
				} else {

				}
			}
			{
				__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_3:
			var __t6 gopurs_runtime.Value
			{
				var __t_tag_4 *Constructor_Main_Z[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Main_Z[gopurs_runtime.Value]](ys_3)
				_ = __t_tag_4
				if __t_tag_4 == nil {
					__t6 = gopurs_runtime.Value{Type: 9, IntVal: 3478632216, UnsafePtr: unsafe.Pointer(Rebox_Main_4264667741_1323331594((&Constructor_Data_Generic_Rep_Inl[uint32, *Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, *Constructor_Main_Z[gopurs_runtime.Value]]]{1, 1454898258})))}
					goto end_branch_6
				} else {

				}
			}
			{
				var __t_tag_5 *Constructor_Main_Z[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Main_Z[gopurs_runtime.Value]](ys_3)
				_ = __t_tag_5
				if __t_tag_5 != nil {
					__t6 = gopurs_runtime.Value{Type: 9, IntVal: 492034566, UnsafePtr: unsafe.Pointer(Rebox_Main_1784434947_2687169876((&Constructor_Data_Generic_Rep_Inr[uint32, *Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, *Constructor_Main_Z[gopurs_runtime.Value]]]{1, Rebox_Main_372708202_966244383((&Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Main_Z[gopurs_runtime.Value])(ys_3.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1714575428, UnsafePtr: unsafe.Pointer((*Constructor_Main_Z[gopurs_runtime.Value])(ys_3.UnsafePtr).V1)}}))})))}
					goto end_branch_6
				} else {

				}
			}
			{
				__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_6:
			return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Data_Eq_Generic_genericEqSum(Call_Data_Eq_Generic_genericEqConstructor(gopurs_runtime.Value{Type: 9, IntVal: 106035173, UnsafePtr: unsafe.Pointer(Rebox_Main_4167548092_1498596945(Rebox_Main_1498596945_4167548092(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Generic_GenericEq[gopurs_runtime.Value]](Get_Data_Eq_Generic_genericEqNoArguments()))))}), Call_Data_Eq_Generic_genericEqConstructor(gopurs_runtime.Apply(genericEqProduct_1_0, Call_Data_Eq_Generic_genericEqArgument(Call_Main_eqY(dictEq_0))))), "genericEq'"), __t3, __t6).IntVal) != (0))
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
		return gopurs_runtime.Bool((gopurs_runtime.Apply2(genericEqConstructor_1_0.V0, xs_2, ys_3).IntVal) != (0))
	})}))}
}

func Rebox_Main_1053099733_3790796878(in *Constructor_Data_Eq_Eq[int64]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_1121961536_3790796878(in *Constructor_Data_Eq_Eq[*Constructor_Main_Z[int64]]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_1498596945_4167548092(in *Constructor_Data_Eq_Generic_GenericEq[gopurs_runtime.Value]) *Constructor_Data_Eq_Generic_GenericEq[uint32] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Generic_GenericEq[uint32]{}
	out.V0 = in.V0
	return out
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
	out := &Constructor_Data_Generic_Rep_Generic[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = in.V0
	out.V1 = in.V1
	return out
}

func Rebox_Main_372708202_966244383(in *Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, *Constructor_Main_Z[gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, *Constructor_Main_Z[gopurs_runtime.Value]]{}
	out.V0 = in.V0
	out.V1 = gopurs_runtime.CoerceToStruct[Constructor_Main_Z[gopurs_runtime.Value]](in.V1)
	return out
}

func Rebox_Main_3790796878_1053099733(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[int64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Eq[int64]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_3790796878_1121961536(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[*Constructor_Main_Z[int64]] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Eq[*Constructor_Main_Z[int64]]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_3790796878_3989972411(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[*Constructor_Main_Z[gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Eq[*Constructor_Main_Z[gopurs_runtime.Value]]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_3989972411_3790796878(in *Constructor_Data_Eq_Eq[*Constructor_Main_Z[gopurs_runtime.Value]]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_4105062992_857262859(in *Constructor_Main_Z[gopurs_runtime.Value]) *Constructor_Main_Z[int64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_Z[int64]{}
	out.V0 = in.V0.IntVal
	out.V1 = Rebox_Main_4105062992_857262859(in.V1)
	return out
}

func Rebox_Main_4167548092_1498596945(in *Constructor_Data_Eq_Generic_GenericEq[uint32]) *Constructor_Data_Eq_Generic_GenericEq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Generic_GenericEq[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
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
}, struct {
	x int64
	y string
}]) *Constructor_Data_Generic_Rep_Generic[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Generic_Rep_Generic[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = in.V0
	out.V1 = in.V1
	return out
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
