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
		cache_Main_genericEqSum = gopurs_runtime.Func(func(dictGenericEq1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_genericEqSum(dictGenericEq1_0_box)
		})
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
			return Call_Main_W(x_0_box)
		})
	})
	return cache_Main_W
}

var cache_Main_genericZ gopurs_runtime.Value
var once_Main_genericZ sync.Once

func Get_Main_genericZ() gopurs_runtime.Value {
	once_Main_genericZ.Do(func() {
		cache_Main_genericZ = gopurs_runtime.Value{Type: 9, IntVal: 1921946594, UnsafePtr: unsafe.Pointer((&Constructor_Data_Generic_Rep_Generic[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Generic_Rep_Generic[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Main_genericZ()).V0), x_0)
		}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Generic_Rep_Generic[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Main_genericZ()).V1), x_0)
		})}))}
	})
	return cache_Main_genericZ
}

var cache_Main_genericY gopurs_runtime.Value
var once_Main_genericY sync.Once

func Get_Main_genericY() gopurs_runtime.Value {
	once_Main_genericY.Do(func() {
		cache_Main_genericY = gopurs_runtime.Value{Type: 9, IntVal: 1921946594, UnsafePtr: unsafe.Pointer((&Constructor_Data_Generic_Rep_Generic[*Constructor_Main_Z[gopurs_runtime.Value], gopurs_runtime.Value]{1, gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t0 gopurs_runtime.Value
			{
				if x_0.Type == 9 && x_0.IntVal == 1714575428 && x_0.UnsafePtr == nil {
					__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3478632216, UnsafePtr: unsafe.Pointer((&Constructor_Data_Generic_Rep_Inl[uint32, *Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, *Constructor_Main_Z[gopurs_runtime.Value]]]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(1454898258), UnsafePtr: nil}}))}
					goto end_branch_0
				} else {

				}
			}
			{
				if x_0.Type == 9 && x_0.IntVal == 1714575428 && x_0.UnsafePtr != nil {
					__t0 = gopurs_runtime.Value{Type: 9, IntVal: 492034566, UnsafePtr: unsafe.Pointer((&Constructor_Data_Generic_Rep_Inr[uint32, *Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, *Constructor_Main_Z[gopurs_runtime.Value]]]{1, gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer((&Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Main_Z[gopurs_runtime.Value])(x_0.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1714575428, UnsafePtr: unsafe.Pointer((*Constructor_Main_Z[gopurs_runtime.Value])(x_0.UnsafePtr).V1)}}))}}))}
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
			var __t1 *Constructor_Main_Z[gopurs_runtime.Value]
			{
				if x_0.Type == 9 && x_0.IntVal == 3478632216 {
					__t1 = (*Constructor_Main_Z[gopurs_runtime.Value])(nil)
					goto end_branch_1
				} else {

				}
			}
			{
				if x_0.Type == 9 && x_0.IntVal == 492034566 {
					__t1 = (&Constructor_Main_Z[gopurs_runtime.Value]{1, (*Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Generic_Rep_Inr[gopurs_runtime.Value, gopurs_runtime.Value])(x_0.UnsafePtr).V0.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Main_Z[gopurs_runtime.Value]]((*Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Generic_Rep_Inr[gopurs_runtime.Value, gopurs_runtime.Value])(x_0.UnsafePtr).V0.UnsafePtr).V1)})
					goto end_branch_1
				} else {

				}
			}
			{
				__t1 = gopurs_runtime.CoerceToStruct[Constructor_Main_Z[gopurs_runtime.Value]](func() gopurs_runtime.Value { panic("Failed pattern match") }())
			}
		end_branch_1:
			return gopurs_runtime.Value{Type: 9, IntVal: 1714575428, UnsafePtr: unsafe.Pointer(__t1)}
		})}))}
	})
	return cache_Main_genericY
}

var cache_Main_genericY__2055741087 gopurs_runtime.Value
var once_Main_genericY__2055741087 sync.Once

func Get_Main_genericY__2055741087() gopurs_runtime.Value {
	once_Main_genericY__2055741087.Do(func() {
		cache_Main_genericY__2055741087 = gopurs_runtime.Value{Type: 9, IntVal: 1921946594, UnsafePtr: unsafe.Pointer((&Constructor_Data_Generic_Rep_Generic[*Constructor_Main_Z[gopurs_runtime.Value], gopurs_runtime.Value]{1, gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t0 gopurs_runtime.Value
			{
				if x_0.Type == 9 && x_0.IntVal == 1714575428 && x_0.UnsafePtr == nil {
					__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3478632216, UnsafePtr: unsafe.Pointer((&Constructor_Data_Generic_Rep_Inl[uint32, *Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, *Constructor_Main_Z[gopurs_runtime.Value]]]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(1454898258), UnsafePtr: nil}}))}
					goto end_branch_0
				} else {

				}
			}
			{
				if x_0.Type == 9 && x_0.IntVal == 1714575428 && x_0.UnsafePtr != nil {
					__t0 = gopurs_runtime.Value{Type: 9, IntVal: 492034566, UnsafePtr: unsafe.Pointer((&Constructor_Data_Generic_Rep_Inr[uint32, *Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, *Constructor_Main_Z[gopurs_runtime.Value]]]{1, gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer((&Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Main_Z[gopurs_runtime.Value])(x_0.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1714575428, UnsafePtr: unsafe.Pointer((*Constructor_Main_Z[gopurs_runtime.Value])(x_0.UnsafePtr).V1)}}))}}))}
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
			var __t1 *Constructor_Main_Z[gopurs_runtime.Value]
			{
				if x_0.Type == 9 && x_0.IntVal == 3478632216 {
					__t1 = (*Constructor_Main_Z[gopurs_runtime.Value])(nil)
					goto end_branch_1
				} else {

				}
			}
			{
				if x_0.Type == 9 && x_0.IntVal == 492034566 {
					__t1 = (&Constructor_Main_Z[gopurs_runtime.Value]{1, (*Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Generic_Rep_Inr[gopurs_runtime.Value, gopurs_runtime.Value])(x_0.UnsafePtr).V0.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Main_Z[gopurs_runtime.Value]]((*Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Generic_Rep_Inr[gopurs_runtime.Value, gopurs_runtime.Value])(x_0.UnsafePtr).V0.UnsafePtr).V1)})
					goto end_branch_1
				} else {

				}
			}
			{
				__t1 = gopurs_runtime.CoerceToStruct[Constructor_Main_Z[gopurs_runtime.Value]](func() gopurs_runtime.Value { panic("Failed pattern match") }())
			}
		end_branch_1:
			return gopurs_runtime.Value{Type: 9, IntVal: 1714575428, UnsafePtr: unsafe.Pointer(__t1)}
		})}))}
	})
	return cache_Main_genericY__2055741087
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

var cache_Main_genericX__2110195358 gopurs_runtime.Value
var once_Main_genericX__2110195358 sync.Once

func Get_Main_genericX__2110195358() gopurs_runtime.Value {
	once_Main_genericX__2110195358.Do(func() {
		cache_Main_genericX__2110195358 = gopurs_runtime.Value{Type: 9, IntVal: 1921946594, UnsafePtr: unsafe.Pointer((&Constructor_Data_Generic_Rep_Generic[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return x_0
		}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return x_0
		})}))}
	})
	return cache_Main_genericX__2110195358
}

var cache_Main_genericW gopurs_runtime.Value
var once_Main_genericW sync.Once

func Get_Main_genericW() gopurs_runtime.Value {
	once_Main_genericW.Do(func() {
		cache_Main_genericW = gopurs_runtime.Value{Type: 9, IntVal: 1921946594, UnsafePtr: unsafe.Pointer((&Constructor_Data_Generic_Rep_Generic[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return x_0
		}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return x_0
		})}))}
	})
	return cache_Main_genericW
}

var cache_Main_eqZ gopurs_runtime.Value
var once_Main_eqZ sync.Once

func Get_Main_eqZ() gopurs_runtime.Value {
	once_Main_eqZ.Do(func() {
		cache_Main_eqZ = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Bool(true)
			})
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
		cache_Main_eqY1 = func() gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=LitRecord bindingType=(Record (Row [genericEq': (Func [(TypeVar a), (TypeVar a)] Boolean)] Any))
			__local_var_0_0 := gopurs_runtime.RecordDict1("genericEq'", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Bool((v_0.IntVal) == (v1_1.IntVal))
				})
			}))
			_ = __local_var_0_0
			return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq[*Constructor_Main_Z[int64]]{1, gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(ys_2 gopurs_runtime.Value) gopurs_runtime.Value {
					// TAST (Let): __local_var_3_1 shape=App(Var) bindingType=Any
					__local_var_3_1 := Call_Main_eqY(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[int64]](Get_Data_Eq_eqInt()))})
					_ = __local_var_3_1
					// TAST (Let): __local_var_4_2 shape=LitRecord bindingType=(Record (Row [genericEq': (Func [(TypeVar a), (TypeVar a)] Boolean)] Any))
					__local_var_4_2 := gopurs_runtime.RecordDict1("genericEq'", gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_1, "eq"), v_4, v1_5).IntVal) != (0))
						})
					}))
					_ = __local_var_4_2
					// TAST (Let): __local_var_5_3 shape=LitRecord bindingType=(Record (Row [genericEq': (Func [(ADT ["Data","Generic","Rep","Product"] [(TypeVar a), (TypeVar b)]), (ADT ["Data","Generic","Rep","Product"] [(TypeVar a), (TypeVar b)])] Boolean)] Any))
					__local_var_5_3 := gopurs_runtime.RecordDict1("genericEq'", gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_0_0, "genericEq'"), (*Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0, (*Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v1_6.UnsafePtr).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_2, "genericEq'"), (*Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V1, (*Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v1_6.UnsafePtr).V1).IntVal) != (0)))
						})
					}))
					_ = __local_var_5_3
					// TAST (Let): __local_var_6_4 shape=LitRecord bindingType=(Record (Row [genericEq': (Func [(TypeVar a), (TypeVar a)] Boolean)] Any))
					__local_var_6_4 := gopurs_runtime.RecordDict1("genericEq'", gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_5_3, "genericEq'"), v_6, v1_7).IntVal) != (0))
						})
					}))
					_ = __local_var_6_4
					var __t15 bool
					{
						var __t_tag_5 *Constructor_Main_Z[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Main_Z[gopurs_runtime.Value]](xs_1)
						if __t_tag_5 == nil {
							var __t8 bool
							{
								var __t_tag_6 *Constructor_Main_Z[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Main_Z[gopurs_runtime.Value]](ys_2)
								if __t_tag_6 == nil {
									__t8 = true
									goto end_branch_8
								} else {

								}
							}
							{
								var __t_tag_7 *Constructor_Main_Z[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Main_Z[gopurs_runtime.Value]](ys_2)
								if __t_tag_7 != nil {
									__t8 = false
									goto end_branch_8
								} else {

								}
							}
							{
								__t8 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
							}
						end_branch_8:
							__t15 = __t8
							goto end_branch_15
						} else {

						}
					}
					{
						var __t_tag_9 *Constructor_Main_Z[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Main_Z[gopurs_runtime.Value]](xs_1)
						if __t_tag_9 != nil {
							var __t12 bool
							{
								var __t_tag_10 *Constructor_Main_Z[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Main_Z[gopurs_runtime.Value]](ys_2)
								if __t_tag_10 == nil {
									__t12 = false
									goto end_branch_12
								} else {

								}
							}
							{
								var __t_tag_11 *Constructor_Main_Z[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Main_Z[gopurs_runtime.Value]](ys_2)
								if __t_tag_11 != nil {
									__t12 = (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_6_4, "genericEq'"), gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer((&Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Main_Z[gopurs_runtime.Value])(xs_1.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1714575428, UnsafePtr: unsafe.Pointer((*Constructor_Main_Z[gopurs_runtime.Value])(xs_1.UnsafePtr).V1)}}))}, gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer((&Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Main_Z[gopurs_runtime.Value])(ys_2.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1714575428, UnsafePtr: unsafe.Pointer((*Constructor_Main_Z[gopurs_runtime.Value])(ys_2.UnsafePtr).V1)}}))}).IntVal) != (0)
									goto end_branch_12
								} else {

								}
							}
							{
								__t12 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
							}
						end_branch_12:
							__t15 = __t12
							goto end_branch_15
						} else {

						}
					}
					{
						var __t_tag_13 *Constructor_Main_Z[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Main_Z[gopurs_runtime.Value]](ys_2)
						if __t_tag_13 == nil {
							__t15 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
							goto end_branch_15
						} else {

						}
					}
					{
						var __t_tag_14 *Constructor_Main_Z[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Main_Z[gopurs_runtime.Value]](ys_2)
						if __t_tag_14 != nil {
							__t15 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
							goto end_branch_15
						} else {

						}
					}
					{
						__t15 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
					}
				end_branch_15:
					return gopurs_runtime.Bool(__t15)
				})
			})}))}
		}()
	})
	return cache_Main_eqY1
}

var cache_Main_eqY2 gopurs_runtime.Value
var once_Main_eqY2 sync.Once

func Get_Main_eqY2() gopurs_runtime.Value {
	once_Main_eqY2.Do(func() {
		cache_Main_eqY2 = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq[*Constructor_Main_Z[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(xs_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(ys_1 gopurs_runtime.Value) gopurs_runtime.Value {
				// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=Any
				__local_var_2_0 := Call_Main_eqY(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main_eqZ()))})
				_ = __local_var_2_0
				// TAST (Let): __local_var_3_1 shape=LitRecord bindingType=(Record (Row [genericEq': (Func [(TypeVar a), (TypeVar a)] Boolean)] Any))
				__local_var_3_1 := gopurs_runtime.RecordDict1("genericEq'", gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_0, "eq"), v_3, v1_4).IntVal) != (0))
					})
				}))
				_ = __local_var_3_1
				// TAST (Let): __local_var_4_2 shape=LitRecord bindingType=(Record (Row [genericEq': (Func [(ADT ["Data","Generic","Rep","Product"] [(TypeVar a), (TypeVar b)]), (ADT ["Data","Generic","Rep","Product"] [(TypeVar a), (TypeVar b)])] Boolean)] Any))
				__local_var_4_2 := gopurs_runtime.RecordDict1("genericEq'", gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_1, "genericEq'"), (*Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1, (*Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V1).IntVal) != (0))
					})
				}))
				_ = __local_var_4_2
				// TAST (Let): __local_var_5_3 shape=LitRecord bindingType=(Record (Row [genericEq': (Func [(TypeVar a), (TypeVar a)] Boolean)] Any))
				__local_var_5_3 := gopurs_runtime.RecordDict1("genericEq'", gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_2, "genericEq'"), v_5, v1_6).IntVal) != (0))
					})
				}))
				_ = __local_var_5_3
				var __t14 bool
				{
					var __t_tag_4 *Constructor_Main_Z[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Main_Z[gopurs_runtime.Value]](xs_0)
					if __t_tag_4 == nil {
						var __t7 bool
						{
							var __t_tag_5 *Constructor_Main_Z[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Main_Z[gopurs_runtime.Value]](ys_1)
							if __t_tag_5 == nil {
								__t7 = true
								goto end_branch_7
							} else {

							}
						}
						{
							var __t_tag_6 *Constructor_Main_Z[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Main_Z[gopurs_runtime.Value]](ys_1)
							if __t_tag_6 != nil {
								__t7 = false
								goto end_branch_7
							} else {

							}
						}
						{
							__t7 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
						}
					end_branch_7:
						__t14 = __t7
						goto end_branch_14
					} else {

					}
				}
				{
					var __t_tag_8 *Constructor_Main_Z[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Main_Z[gopurs_runtime.Value]](xs_0)
					if __t_tag_8 != nil {
						var __t11 bool
						{
							var __t_tag_9 *Constructor_Main_Z[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Main_Z[gopurs_runtime.Value]](ys_1)
							if __t_tag_9 == nil {
								__t11 = false
								goto end_branch_11
							} else {

							}
						}
						{
							var __t_tag_10 *Constructor_Main_Z[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Main_Z[gopurs_runtime.Value]](ys_1)
							if __t_tag_10 != nil {
								__t11 = (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_5_3, "genericEq'"), gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer((&Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Main_Z[gopurs_runtime.Value])(xs_0.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1714575428, UnsafePtr: unsafe.Pointer((*Constructor_Main_Z[gopurs_runtime.Value])(xs_0.UnsafePtr).V1)}}))}, gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer((&Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Main_Z[gopurs_runtime.Value])(ys_1.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1714575428, UnsafePtr: unsafe.Pointer((*Constructor_Main_Z[gopurs_runtime.Value])(ys_1.UnsafePtr).V1)}}))}).IntVal) != (0)
								goto end_branch_11
							} else {

							}
						}
						{
							__t11 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
						}
					end_branch_11:
						__t14 = __t11
						goto end_branch_14
					} else {

					}
				}
				{
					var __t_tag_12 *Constructor_Main_Z[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Main_Z[gopurs_runtime.Value]](ys_1)
					if __t_tag_12 == nil {
						__t14 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
						goto end_branch_14
					} else {

					}
				}
				{
					var __t_tag_13 *Constructor_Main_Z[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Main_Z[gopurs_runtime.Value]](ys_1)
					if __t_tag_13 != nil {
						__t14 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
						goto end_branch_14
					} else {

					}
				}
				{
					__t14 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
				}
			end_branch_14:
				return gopurs_runtime.Bool(__t14)
			})
		})}))}
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
		cache_Main_eqX1 = func() gopurs_runtime.Value {
			// TAST (Let): __local_var_0_1 shape=LitRecord bindingType=(Record (Row [genericEq': (Func [(TypeVar a), (TypeVar a)] Boolean)] Any))
			__local_var_0_1 := gopurs_runtime.RecordDict1("genericEq'", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Bool((v_0.IntVal) == (v1_1.IntVal))
				})
			}))
			_ = __local_var_0_1
			// TAST (Let): genericEqConstructor_0_0 shape=Let(LitRecord) bindingType=(ADT ["Data","Eq","Generic","GenericEq"] [Any])
			genericEqConstructor_0_0 := (&Constructor_Data_Eq_Generic_GenericEq[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_0_1, "genericEq'"), v_1, v1_2).IntVal) != (0))
				})
			})})
			_ = genericEqConstructor_0_0
			return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(ys_2 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.Box(genericEqConstructor_0_0.V0), xs_1, ys_2).IntVal) != (0))
				})
			})}))}
		}()
	})
	return cache_Main_eqX1
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=(ADT ["Effect","Effect"] [Unit])
			__local_var_0_0 := gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("false"))
			_ = __local_var_0_0
			_dollar___unused_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = _dollar___unused_1_1
			_dollar___unused_2_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("true")), gopurs_runtime.Value{})
			_ = _dollar___unused_2_2
			_dollar___unused_3_3 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("true")), gopurs_runtime.Value{})
			_ = _dollar___unused_3_3
			_dollar___unused_4_4 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("false")), gopurs_runtime.Value{})
			_ = _dollar___unused_4_4
			var __t6 string
			{
				if (gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[*Constructor_Main_Z[gopurs_runtime.Value]]](Get_Main_eqY2()).V0), gopurs_runtime.Value{Type: 9, IntVal: 1714575428, UnsafePtr: unsafe.Pointer((*Constructor_Main_Z[gopurs_runtime.Value])(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1714575428, UnsafePtr: unsafe.Pointer((*Constructor_Main_Z[gopurs_runtime.Value])(nil))}).IntVal) != (0)).IntVal) != (0) {
					__t6 = "true"
					goto end_branch_6
				} else {

				}
			}
			{
				__t6 = "false"
			}
		end_branch_6:
			_dollar___unused_5_5 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(__t6)), gopurs_runtime.Value{})
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
	V0 gopurs_runtime.Value
	V1 *Constructor_Main_Z[gopurs_runtime.Value]
}

type Constructor_Main_X[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

func Call_Main_genericEqSum(dictGenericEq1_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictGenericEq1_0 gopurs_runtime.Value = dictGenericEq1_0_loop
	_ = dictGenericEq1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 106035173, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Generic_GenericEq[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t0 bool
			{
				if v_1.Type == 9 && v_1.IntVal == 3478632216 {
					__t0 = (v1_2.Type == 9 && v1_2.IntVal == 3478632216)
					goto end_branch_0
				} else {

				}
			}
			{
				__t0 = (v_1.Type == 9 && v_1.IntVal == 492034566) && ((v1_2.Type == 9 && v1_2.IntVal == 492034566) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericEq1_0, "genericEq'"), (*Constructor_Data_Generic_Rep_Inr[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V0, (*Constructor_Data_Generic_Rep_Inr[gopurs_runtime.Value, gopurs_runtime.Value])(v1_2.UnsafePtr).V0).IntVal) != (0)))
			}
		end_branch_0:
			return gopurs_runtime.Bool(__t0)
		})
	})}))}
}

func Call_Main_W(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
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
		// TAST (Let): __local_var_1_0 shape=LitRecord bindingType=(Record (Row [genericEq': (Func [(TypeVar a), (TypeVar a)] Boolean)] Any))
		__local_var_1_0 := gopurs_runtime.RecordDict1("genericEq'", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), v_1, v1_2).IntVal) != (0))
			})
		}))
		_ = __local_var_1_0
		return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq[*Constructor_Main_Z[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(ys_3 gopurs_runtime.Value) gopurs_runtime.Value {
				// TAST (Let): __local_var_4_1 shape=App(Var) bindingType=Any
				__local_var_4_1 := Call_Main_eqY(dictEq_0)
				_ = __local_var_4_1
				// TAST (Let): __local_var_5_2 shape=LitRecord bindingType=(Record (Row [genericEq': (Func [(TypeVar a), (TypeVar a)] Boolean)] Any))
				__local_var_5_2 := gopurs_runtime.RecordDict1("genericEq'", gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_1, "eq"), v_5, v1_6).IntVal) != (0))
					})
				}))
				_ = __local_var_5_2
				// TAST (Let): __local_var_6_3 shape=LitRecord bindingType=(Record (Row [genericEq': (Func [(ADT ["Data","Generic","Rep","Product"] [(TypeVar a), (TypeVar b)]), (ADT ["Data","Generic","Rep","Product"] [(TypeVar a), (TypeVar b)])] Boolean)] Any))
				__local_var_6_3 := gopurs_runtime.RecordDict1("genericEq'", gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "genericEq'"), (*Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V0, (*Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v1_7.UnsafePtr).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_5_2, "genericEq'"), (*Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V1, (*Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v1_7.UnsafePtr).V1).IntVal) != (0)))
					})
				}))
				_ = __local_var_6_3
				// TAST (Let): __local_var_7_4 shape=LitRecord bindingType=(Record (Row [genericEq': (Func [(TypeVar a), (TypeVar a)] Boolean)] Any))
				__local_var_7_4 := gopurs_runtime.RecordDict1("genericEq'", gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_6_3, "genericEq'"), v_7, v1_8).IntVal) != (0))
					})
				}))
				_ = __local_var_7_4
				var __t15 bool
				{
					var __t_tag_5 *Constructor_Main_Z[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Main_Z[gopurs_runtime.Value]](xs_2)
					if __t_tag_5 == nil {
						var __t8 bool
						{
							var __t_tag_6 *Constructor_Main_Z[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Main_Z[gopurs_runtime.Value]](ys_3)
							if __t_tag_6 == nil {
								__t8 = true
								goto end_branch_8
							} else {

							}
						}
						{
							var __t_tag_7 *Constructor_Main_Z[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Main_Z[gopurs_runtime.Value]](ys_3)
							if __t_tag_7 != nil {
								__t8 = false
								goto end_branch_8
							} else {

							}
						}
						{
							__t8 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
						}
					end_branch_8:
						__t15 = __t8
						goto end_branch_15
					} else {

					}
				}
				{
					var __t_tag_9 *Constructor_Main_Z[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Main_Z[gopurs_runtime.Value]](xs_2)
					if __t_tag_9 != nil {
						var __t12 bool
						{
							var __t_tag_10 *Constructor_Main_Z[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Main_Z[gopurs_runtime.Value]](ys_3)
							if __t_tag_10 == nil {
								__t12 = false
								goto end_branch_12
							} else {

							}
						}
						{
							var __t_tag_11 *Constructor_Main_Z[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Main_Z[gopurs_runtime.Value]](ys_3)
							if __t_tag_11 != nil {
								__t12 = (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_7_4, "genericEq'"), gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer((&Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Main_Z[gopurs_runtime.Value])(xs_2.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1714575428, UnsafePtr: unsafe.Pointer((*Constructor_Main_Z[gopurs_runtime.Value])(xs_2.UnsafePtr).V1)}}))}, gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer((&Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Main_Z[gopurs_runtime.Value])(ys_3.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1714575428, UnsafePtr: unsafe.Pointer((*Constructor_Main_Z[gopurs_runtime.Value])(ys_3.UnsafePtr).V1)}}))}).IntVal) != (0)
								goto end_branch_12
							} else {

							}
						}
						{
							__t12 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
						}
					end_branch_12:
						__t15 = __t12
						goto end_branch_15
					} else {

					}
				}
				{
					var __t_tag_13 *Constructor_Main_Z[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Main_Z[gopurs_runtime.Value]](ys_3)
					if __t_tag_13 == nil {
						__t15 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
						goto end_branch_15
					} else {

					}
				}
				{
					var __t_tag_14 *Constructor_Main_Z[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Main_Z[gopurs_runtime.Value]](ys_3)
					if __t_tag_14 != nil {
						__t15 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
						goto end_branch_15
					} else {

					}
				}
				{
					__t15 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
				}
			end_branch_15:
				return gopurs_runtime.Bool(__t15)
			})
		})}))}
	}
}

func Call_Main_eqX(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
	_ = dictEq_0
	// TAST (Let): __local_var_1_1 shape=LitRecord bindingType=(Record (Row [genericEq': (Func [(TypeVar a), (TypeVar a)] Boolean)] Any))
	__local_var_1_1 := gopurs_runtime.RecordDict1("genericEq'", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), v_1, v1_2).IntVal) != (0))
		})
	}))
	_ = __local_var_1_1
	// TAST (Let): genericEqConstructor_1_0 shape=Let(LitRecord) bindingType=(ADT ["Data","Eq","Generic","GenericEq"] [Any])
	genericEqConstructor_1_0 := (&Constructor_Data_Eq_Generic_GenericEq[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "genericEq'"), v_2, v1_3).IntVal) != (0))
		})
	})})
	_ = genericEqConstructor_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(ys_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.Box(genericEqConstructor_1_0.V0), xs_2, ys_3).IntVal) != (0))
		})
	})}))}
}
