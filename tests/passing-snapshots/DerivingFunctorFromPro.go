package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_eqArray gopurs_runtime.Value
var once_Main_eqArray sync.Once

func Get_Main_eqArray() gopurs_runtime.Value {
	once_Main_eqArray.Do(func() {
		cache_Main_eqArray = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_4186496832_3790796878(Rebox_Main_3790796878_4186496832(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_Eq_eqArray(Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("value")
		})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})))))))}
	})
	return cache_Main_eqArray
}

var cache_Main_Test1 gopurs_runtime.Value
var once_Main_Test1 sync.Once

func Get_Main_Test1() gopurs_runtime.Value {
	once_Main_Test1.Do(func() {
		cache_Main_Test1 = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3720114489, UnsafePtr: unsafe.Pointer((&Constructor_Main_Test1[gopurs_runtime.Value]{1, value0}))}
		})
	})
	return cache_Main_Test1
}

var cache_Main_Test1__525027919 gopurs_runtime.Value
var once_Main_Test1__525027919 sync.Once

func Get_Main_Test1__525027919() gopurs_runtime.Value {
	once_Main_Test1__525027919.Do(func() {
		cache_Main_Test1__525027919 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_Test1__525027919(__eta_norm_0_0_box)
		})
	})
	return cache_Main_Test1__525027919
}

var cache_Main_Test2 gopurs_runtime.Value
var once_Main_Test2 sync.Once

func Get_Main_Test2() gopurs_runtime.Value {
	once_Main_Test2.Do(func() {
		cache_Main_Test2 = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2375191994, UnsafePtr: unsafe.Pointer((&Constructor_Main_Test2[gopurs_runtime.Value]{1, func() struct {
				f gopurs_runtime.Value
			} {
				orig := value0
				_ = orig
				clone := struct {
					f gopurs_runtime.Value
				}{}
				clone.f = gopurs_runtime.RecordGet(orig, "f")
				return clone
			}()}))}
		})
	})
	return cache_Main_Test2
}

var cache_Main_Test2__3611333742 gopurs_runtime.Value
var once_Main_Test2__3611333742 sync.Once

func Get_Main_Test2__3611333742() gopurs_runtime.Value {
	once_Main_Test2__3611333742.Do(func() {
		cache_Main_Test2__3611333742 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_Test2__3611333742(func() struct {
				f gopurs_runtime.Value
			} {
				orig := __eta_norm_0_0_box
				_ = orig
				clone := struct {
					f gopurs_runtime.Value
				}{}
				clone.f = gopurs_runtime.RecordGet(orig, "f")
				return clone
			}())
		})
	})
	return cache_Main_Test2__3611333742
}

var cache_Main_functorTest gopurs_runtime.Value
var once_Main_functorTest sync.Once

func Get_Main_functorTest() gopurs_runtime.Value {
	once_Main_functorTest.Do(func() {
		cache_Main_functorTest = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, m_1 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t0 gopurs_runtime.Value
			{
				if m_1.Type == 9 && m_1.IntVal == 3720114489 {
					__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3720114489, UnsafePtr: unsafe.Pointer((&Constructor_Main_Test1[gopurs_runtime.Value]{1, Call_Control_Semigroupoid_composeFlipped(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()), gopurs_runtime.Func(func(b2c_2 gopurs_runtime.Value) gopurs_runtime.Value {
						return Call_Control_Semigroupoid_composeFlipped(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()), gopurs_runtime.Apply(Get_Data_Functor_arrayMap(), f_0), Call_Control_Semigroupoid_composeFlipped(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()), b2c_2, gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
							return x_3
						})))
					}), Call_Control_Semigroupoid_composeFlipped(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()), (*Constructor_Main_Test1[gopurs_runtime.Value])(m_1.UnsafePtr).V0, gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
						return x_2
					})))}))}
					goto end_branch_0
				} else {

				}
			}
			{
				if m_1.Type == 9 && m_1.IntVal == 2375191994 {
					__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2375191994, UnsafePtr: unsafe.Pointer((&Constructor_Main_Test2[gopurs_runtime.Value]{1, func() struct {
						f gopurs_runtime.Value
					} {
						clone := (*Constructor_Main_Test2[gopurs_runtime.Value])(m_1.UnsafePtr).V0
						clone.f = Call_Control_Semigroupoid_composeFlipped(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()), gopurs_runtime.Func(func(b2c_2 gopurs_runtime.Value) gopurs_runtime.Value {
							return Call_Control_Semigroupoid_composeFlipped(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()), gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
								return func() gopurs_runtime.Value {
									orig := func() struct {
										a gopurs_runtime.Value
									} {
										orig := gopurs_runtime.RecordUpdate1(v1_3, "a", gopurs_runtime.Apply(f_0, gopurs_runtime.RecordGet(v1_3, "a")))
										_ = orig
										clone := struct {
											a gopurs_runtime.Value
										}{}
										clone.a = gopurs_runtime.RecordGet(orig, "a")
										return clone
									}()
									_ = orig
									return gopurs_runtime.RecordDict1("a", orig.a)
								}()
							}), Call_Control_Semigroupoid_composeFlipped(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()), b2c_2, gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
								return x_3
							})))
						}), Call_Control_Semigroupoid_composeFlipped(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()), (*Constructor_Main_Test2[gopurs_runtime.Value])(m_1.UnsafePtr).V0.f, gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
							return x_2
						})))
						return clone
					}()}))}
					goto end_branch_0
				} else {

				}
			}
			{
				__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_0:
			return __t0
		})}))}
	})
	return cache_Main_functorTest
}

var cache_Main_adapt gopurs_runtime.Value
var once_Main_adapt sync.Once

func Get_Main_adapt() gopurs_runtime.Value {
	once_Main_adapt.Do(func() {
		cache_Main_adapt = gopurs_runtime.Apply(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Main_functorTest()).V0, gopurs_runtime.Func(func(n_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				orig := struct {
					value int64
				}{gopurs_runtime.IntAdd(n_0.IntVal, int64(1))}
				_ = orig
				return gopurs_runtime.RecordDict1("value", gopurs_runtime.Int(orig.value))
			}()
		}))
	})
	return cache_Main_adapt
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=Any
			__local_var_0_0 := gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Int(int64(7)))
			_ = __local_var_0_0
			__local_var_1_2 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = __local_var_1_2
			__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_1_2), gopurs_runtime.Value{})
			_ = __local_var_1_1
			// TAST (Let): v_2_3 shape=App(Other) bindingType=(ADT ["Main","Test"] [(Record (Row [value: Int] Empty))])
			v_2_3 := gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Main_functorTest()).V0, gopurs_runtime.Func(func(n_2 gopurs_runtime.Value) gopurs_runtime.Value {
				return func() gopurs_runtime.Value {
					orig := struct {
						value int64
					}{gopurs_runtime.IntAdd(n_2.IntVal, int64(1))}
					_ = orig
					return gopurs_runtime.RecordDict1("value", gopurs_runtime.Int(orig.value))
				}()
			}), gopurs_runtime.Value{Type: 9, IntVal: 3720114489, UnsafePtr: unsafe.Pointer(Rebox_Main_2574500310_2598952845((&Constructor_Main_Test1[int64]{1, gopurs_runtime.Func(func(consume_2 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Int(gopurs_runtime.Apply(consume_2, func() gopurs_runtime.Value {
					arr := []int64{__local_var_1_1.IntVal, gopurs_runtime.IntAdd(__local_var_1_1.IntVal, int64(1))}
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Int(v)
					}
					return gopurs_runtime.Array(boxed)
				}()).IntVal)
			})})))})
			_ = v_2_3
			__local_var_3_4 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("map - nested function arguments"), gopurs_runtime.Bool((v_2_3.Type == 9 && v_2_3.IntVal == 3720114489) && ((gopurs_runtime.Apply((*Constructor_Main_Test1[gopurs_runtime.Value])(v_2_3.UnsafePtr).V0, gopurs_runtime.Func(func(xs_3 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t5 int64
				{
					if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Data_Eq_eqArray(Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Str("value")
					})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}))), "eq"), xs_3, func() gopurs_runtime.Value {
						arr := []struct {
							value int64
						}{struct {
							value int64
						}{int64(8)}, struct {
							value int64
						}{int64(9)}}
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = func() gopurs_runtime.Value {
								orig := v
								_ = orig
								return gopurs_runtime.RecordDict1("value", gopurs_runtime.Int(orig.value))
							}()
						}
						return gopurs_runtime.Array(boxed)
					}()).IntVal) != (0) {
						__t5 = int64(42)
						goto end_branch_5
					} else {

					}
				}
				{
					__t5 = int64(0)
				}
			end_branch_5:
				return gopurs_runtime.Int(__t5)
			})).IntVal) == (int64(42))))), gopurs_runtime.Value{})
			_ = __local_var_3_4
			// TAST (Let): v_4_6 shape=App(Other) bindingType=(ADT ["Main","Test"] [(Record (Row [value: Int] Empty))])
			v_4_6 := gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Main_functorTest()).V0, gopurs_runtime.Func(func(n_4 gopurs_runtime.Value) gopurs_runtime.Value {
				return func() gopurs_runtime.Value {
					orig := struct {
						value int64
					}{gopurs_runtime.IntAdd(n_4.IntVal, int64(1))}
					_ = orig
					return gopurs_runtime.RecordDict1("value", gopurs_runtime.Int(orig.value))
				}()
			}), gopurs_runtime.Value{Type: 9, IntVal: 2375191994, UnsafePtr: unsafe.Pointer((&Constructor_Main_Test2[gopurs_runtime.Value]{1, struct {
				f gopurs_runtime.Value
			}{gopurs_runtime.Func(func(consume_4 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Int(gopurs_runtime.Apply(consume_4, func() gopurs_runtime.Value {
					orig := struct {
						a int64
					}{__local_var_1_1.IntVal}
					_ = orig
					return gopurs_runtime.RecordDict1("a", gopurs_runtime.Int(orig.a))
				}()).IntVal)
			})}}))})
			_ = v_4_6
			__local_var_5_7 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("map - nested record function arguments"), gopurs_runtime.Bool((v_4_6.Type == 9 && v_4_6.IntVal == 2375191994) && ((gopurs_runtime.Apply((*Constructor_Main_Test2[gopurs_runtime.Value])(v_4_6.UnsafePtr).V0.f, gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Int(gopurs_runtime.IntMul(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(x_5, "a"), "value").IntVal, int64(2)))
			})).IntVal) == (int64(16))))), gopurs_runtime.Value{})
			_ = __local_var_5_7
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

type Constructor_Main_Test1[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

type Constructor_Main_Test2[T_a any] struct {
	Rc uint32
	V0 struct {
		f gopurs_runtime.Value
	}
}

func Call_Main_Test1__525027919(__eta_norm_0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
Test1__525027919:
	for {
		if false {
			continue Test1__525027919
		}
		var __eta_norm_0_0 gopurs_runtime.Value = __eta_norm_0_0_loop
		_ = __eta_norm_0_0
		return gopurs_runtime.Value{Type: 9, IntVal: 3720114489, UnsafePtr: unsafe.Pointer(Rebox_Main_2574500310_2598952845((&Constructor_Main_Test1[int64]{1, __eta_norm_0_0})))}
	}
}

func Call_Main_Test2__3611333742(__eta_norm_0_0_loop struct {
	f gopurs_runtime.Value
}) gopurs_runtime.Value {
Test2__3611333742:
	for {
		if false {
			continue Test2__3611333742
		}
		var __eta_norm_0_0 struct {
			f gopurs_runtime.Value
		} = __eta_norm_0_0_loop
		_ = __eta_norm_0_0
		return gopurs_runtime.Value{Type: 9, IntVal: 2375191994, UnsafePtr: unsafe.Pointer(Rebox_Main_3188207925_1106546350((&Constructor_Main_Test2[int64]{1, __eta_norm_0_0})))}
	}
}

func Rebox_Main_1053099733_3790796878(in *Constructor_Data_Eq_Eq[int64]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_2574500310_2598952845(in *Constructor_Main_Test1[int64]) *Constructor_Main_Test1[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Main_Test1[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_3188207925_1106546350(in *Constructor_Main_Test2[int64]) *Constructor_Main_Test2[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Main_Test2[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_3790796878_1053099733(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[int64] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[int64])(unsafe.Pointer(in))
}

func Rebox_Main_3790796878_4186496832(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[[]struct {
	value int64
}] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[[]struct {
		value int64
	}])(unsafe.Pointer(in))
}

func Rebox_Main_4186496832_3790796878(in *Constructor_Data_Eq_Eq[[]struct {
	value int64
}]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[gopurs_runtime.Value])(unsafe.Pointer(in))
}
