package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_BIsSymbol gopurs_runtime.Value
var once_Main_BIsSymbol sync.Once

func Get_Main_BIsSymbol() gopurs_runtime.Value {
	once_Main_BIsSymbol.Do(func() {
		cache_Main_BIsSymbol = gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("B")
		}))
	})
	return cache_Main_BIsSymbol
}

var cache_Main_genericShowConstructor gopurs_runtime.Value
var once_Main_genericShowConstructor sync.Once

func Get_Main_genericShowConstructor() gopurs_runtime.Value {
	once_Main_genericShowConstructor.Do(func() {
		cache_Main_genericShowConstructor = gopurs_runtime.Func(func(dictIsSymbol_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_genericShowConstructor(dictIsSymbol_0_box)
		})
	})
	return cache_Main_genericShowConstructor
}

var cache_Main_genericShowConstructor1 gopurs_runtime.Value
var once_Main_genericShowConstructor1 sync.Once

func Get_Main_genericShowConstructor1() gopurs_runtime.Value {
	once_Main_genericShowConstructor1.Do(func() {
		cache_Main_genericShowConstructor1 = gopurs_runtime.RecordDict1("genericShow'", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("Z")
		}))
	})
	return cache_Main_genericShowConstructor1
}

var cache_Main_B2IsSymbol gopurs_runtime.Value
var once_Main_B2IsSymbol sync.Once

func Get_Main_B2IsSymbol() gopurs_runtime.Value {
	once_Main_B2IsSymbol.Do(func() {
		cache_Main_B2IsSymbol = gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("B2")
		}))
	})
	return cache_Main_B2IsSymbol
}

var cache_Main_genericShowConstructor2 gopurs_runtime.Value
var once_Main_genericShowConstructor2 sync.Once

func Get_Main_genericShowConstructor2() gopurs_runtime.Value {
	once_Main_genericShowConstructor2.Do(func() {
		cache_Main_genericShowConstructor2 = gopurs_runtime.RecordDict1("genericShow'", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("Z2")
		}))
	})
	return cache_Main_genericShowConstructor2
}

var cache_Main_A2 gopurs_runtime.Value
var once_Main_A2 sync.Once

func Get_Main_A2() gopurs_runtime.Value {
	once_Main_A2.Do(func() {
		cache_Main_A2 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_A2(x_0_box)
		})
	})
	return cache_Main_A2
}

var cache_Main_B2 gopurs_runtime.Value
var once_Main_B2 sync.Once

func Get_Main_B2() gopurs_runtime.Value {
	once_Main_B2.Do(func() {
		cache_Main_B2 = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4102534158, UnsafePtr: unsafe.Pointer((&Constructor_Main_B2{1, value0}))}
		})
	})
	return cache_Main_B2
}

var cache_Main_Z2 gopurs_runtime.Value
var once_Main_Z2 sync.Once

func Get_Main_Z2() gopurs_runtime.Value {
	once_Main_Z2.Do(func() {
		cache_Main_Z2 = gopurs_runtime.Value{Type: 9, IntVal: 4102534158, UnsafePtr: unsafe.Pointer((*Constructor_Main_B2)(nil))}
	})
	return cache_Main_Z2
}

var cache_Main_C2 gopurs_runtime.Value
var once_Main_C2 sync.Once

func Get_Main_C2() gopurs_runtime.Value {
	once_Main_C2.Do(func() {
		cache_Main_C2 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_C2(x_0_box)
		})
	})
	return cache_Main_C2
}

var cache_Main_A gopurs_runtime.Value
var once_Main_A sync.Once

func Get_Main_A() gopurs_runtime.Value {
	once_Main_A.Do(func() {
		cache_Main_A = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_A(x_0_box)
		})
	})
	return cache_Main_A
}

var cache_Main_B gopurs_runtime.Value
var once_Main_B sync.Once

func Get_Main_B() gopurs_runtime.Value {
	once_Main_B.Do(func() {
		cache_Main_B = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4250879068, UnsafePtr: unsafe.Pointer((&Constructor_Main_B{1, gopurs_runtime.CoerceToStruct[Constructor_Main_B](value0)}))}
		})
	})
	return cache_Main_B
}

var cache_Main_Z gopurs_runtime.Value
var once_Main_Z sync.Once

func Get_Main_Z() gopurs_runtime.Value {
	once_Main_Z.Do(func() {
		cache_Main_Z = gopurs_runtime.Value{Type: 9, IntVal: 4250879068, UnsafePtr: unsafe.Pointer((*Constructor_Main_B)(nil))}
	})
	return cache_Main_Z
}

var cache_Main_C gopurs_runtime.Value
var once_Main_C sync.Once

func Get_Main_C() gopurs_runtime.Value {
	once_Main_C.Do(func() {
		cache_Main_C = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_C(x_0_box)
		})
	})
	return cache_Main_C
}

var cache_Main_genericC_ gopurs_runtime.Value
var once_Main_genericC_ sync.Once

func Get_Main_genericC_() gopurs_runtime.Value {
	once_Main_genericC_.Do(func() {
		cache_Main_genericC_ = gopurs_runtime.Value{Type: 9, IntVal: 1921946594, UnsafePtr: unsafe.Pointer((&Constructor_Data_Generic_Rep_Generic{1, gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4250879068, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Main_B](x_0))}
		}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4250879068, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Main_B](x_0))}
		})}))}
	})
	return cache_Main_genericC_
}

var cache_Main_genericC2_ gopurs_runtime.Value
var once_Main_genericC2_ sync.Once

func Get_Main_genericC2_() gopurs_runtime.Value {
	once_Main_genericC2_.Do(func() {
		cache_Main_genericC2_ = gopurs_runtime.Value{Type: 9, IntVal: 1921946594, UnsafePtr: unsafe.Pointer((&Constructor_Data_Generic_Rep_Generic{1, gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return x_0
		}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return x_0
		})}))}
	})
	return cache_Main_genericC2_
}

var cache_Main_genericB_ gopurs_runtime.Value
var once_Main_genericB_ sync.Once

func Get_Main_genericB_() gopurs_runtime.Value {
	once_Main_genericB_.Do(func() {
		cache_Main_genericB_ = gopurs_runtime.Value{Type: 9, IntVal: 1921946594, UnsafePtr: unsafe.Pointer((&Constructor_Data_Generic_Rep_Generic{1, gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t0 gopurs_runtime.Value
			{
				if x_0.Type == 9 && x_0.IntVal == 4250879068 && x_0.UnsafePtr != nil {
					__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3478632216, UnsafePtr: unsafe.Pointer((&Constructor_Data_Generic_Rep_Inl{1, gopurs_runtime.Value{Type: 9, IntVal: 4250879068, UnsafePtr: unsafe.Pointer((*Constructor_Main_B)(x_0.UnsafePtr).V0)}}))}
					goto end_branch_0
				} else {

				}
			}
			{
				if x_0.Type == 9 && x_0.IntVal == 4250879068 && x_0.UnsafePtr == nil {
					__t0 = gopurs_runtime.Value{Type: 9, IntVal: 492034566, UnsafePtr: unsafe.Pointer((&Constructor_Data_Generic_Rep_Inr{1, gopurs_runtime.Value{Type: 9, IntVal: int64(1454898258), UnsafePtr: nil}}))}
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
			var __t1 *Constructor_Main_B
			{
				if x_0.Type == 9 && x_0.IntVal == 3478632216 {
					__t1 = (&Constructor_Main_B{1, gopurs_runtime.CoerceToStruct[Constructor_Main_B]((*Constructor_Data_Generic_Rep_Inl)(x_0.UnsafePtr).V0)})
					goto end_branch_1
				} else {

				}
			}
			{
				if x_0.Type == 9 && x_0.IntVal == 492034566 {
					__t1 = (*Constructor_Main_B)(nil)
					goto end_branch_1
				} else {

				}
			}
			{
				__t1 = gopurs_runtime.CoerceToStruct[Constructor_Main_B](func() gopurs_runtime.Value { panic("Failed pattern match") }())
			}
		end_branch_1:
			return gopurs_runtime.Value{Type: 9, IntVal: 4250879068, UnsafePtr: unsafe.Pointer(__t1)}
		})}))}
	})
	return cache_Main_genericB_
}

var cache_Main_showB gopurs_runtime.Value
var once_Main_showB sync.Once

func Get_Main_showB() gopurs_runtime.Value {
	once_Main_showB.Do(func() {
		cache_Main_showB = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer((&Constructor_Data_Show_Show{1, gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
			__local_var_1_0 := gopurs_runtime.RecordDict1("genericShowArgs", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return func() gopurs_runtime.Value {
					arr := []string{gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show](Get_Main_showC()).V0), v_1).StrVal()}
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Str(v)
					}
					return gopurs_runtime.Array(boxed)
				}()
			}))
			_ = __local_var_1_0
			// TAST (Let): __local_var_2_1 -> gopurs_runtime.Value
			__local_var_2_1 := gopurs_runtime.RecordDict1("genericShow'", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
				// TAST (Let): v1_3_2 -> gopurs_runtime.Value
				v1_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "genericShowArgs"), v_2)
				_ = v1_3_2
				var __t3 string
				{
					if (gopurs_runtime.Int(int64(gopurs_runtime.ArrayLength(v1_3_2))).IntVal) == (0) {
						__t3 = "B"
						goto end_branch_3
					} else {

					}
				}
				{
					__t3 = (("(") + (gopurs_runtime.Apply2(Get_Data_Show_Generic_intercalate(), gopurs_runtime.Str(" "), func() gopurs_runtime.Value {
						arr := func() []string {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(Get_Data_Semigroup_concatArray(), func() gopurs_runtime.Value {
									arr := []string{"B"}
									boxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										boxed[i] = gopurs_runtime.Str(v)
									}
									return gopurs_runtime.Array(boxed)
								}(), v1_3_2).UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}()).UnsafePtr)
							unboxed := make([]string, len(arr))
							for i, v := range arr {
								unboxed[i] = v.StrVal()
							}
							return unboxed
						}()
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Str(v)
						}
						return gopurs_runtime.Array(boxed)
					}()).StrVal())) + (")")
				}
			end_branch_3:
				return gopurs_runtime.Str(__t3)
			}))
			_ = __local_var_2_1
			var __t6 string
			{
				var __t_tag_4 *Constructor_Main_B = gopurs_runtime.CoerceToStruct[Constructor_Main_B](x_0)
				if __t_tag_4 != nil {
					__t6 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "genericShow'"), gopurs_runtime.Value{Type: 9, IntVal: 4250879068, UnsafePtr: unsafe.Pointer((*Constructor_Main_B)(x_0.UnsafePtr).V0)}).StrVal()
					goto end_branch_6
				} else {

				}
			}
			{
				var __t_tag_5 *Constructor_Main_B = gopurs_runtime.CoerceToStruct[Constructor_Main_B](x_0)
				if __t_tag_5 == nil {
					__t6 = "Z"
					goto end_branch_6
				} else {

				}
			}
			{
				__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }().StrVal()
			}
		end_branch_6:
			return gopurs_runtime.Str(__t6)
		})}))}
	})
	return cache_Main_showB
}

var cache_Main_showA gopurs_runtime.Value
var once_Main_showA sync.Once

func Get_Main_showA() gopurs_runtime.Value {
	once_Main_showA.Do(func() {
		cache_Main_showA = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show](Get_Main_showB()))}
	})
	return cache_Main_showA
}

var cache_Main_showC gopurs_runtime.Value
var once_Main_showC sync.Once

func Get_Main_showC() gopurs_runtime.Value {
	once_Main_showC.Do(func() {
		cache_Main_showC = func() gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 -> gopurs_runtime.Value
			__local_var_0_0 := gopurs_runtime.RecordDict1("genericShowArgs", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
				// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
				__local_var_1_1 := gopurs_runtime.RecordDict1("genericShowArgs", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
					return func() gopurs_runtime.Value {
						arr := []string{gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show](Get_Main_showC()).V0), v_1).StrVal()}
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Str(v)
						}
						return gopurs_runtime.Array(boxed)
					}()
				}))
				_ = __local_var_1_1
				// TAST (Let): __local_var_2_2 -> gopurs_runtime.Value
				__local_var_2_2 := gopurs_runtime.RecordDict1("genericShow'", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
					// TAST (Let): v1_3_3 -> gopurs_runtime.Value
					v1_3_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "genericShowArgs"), v_2)
					_ = v1_3_3
					var __t4 string
					{
						if (gopurs_runtime.Int(int64(gopurs_runtime.ArrayLength(v1_3_3))).IntVal) == (0) {
							__t4 = "B"
							goto end_branch_4
						} else {

						}
					}
					{
						__t4 = (("(") + (gopurs_runtime.Apply2(Get_Data_Show_Generic_intercalate(), gopurs_runtime.Str(" "), func() gopurs_runtime.Value {
							arr := func() []string {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(Get_Data_Semigroup_concatArray(), func() gopurs_runtime.Value {
										arr := []string{"B"}
										boxed := make([]gopurs_runtime.Value, len(arr))
										for i, v := range arr {
											boxed[i] = gopurs_runtime.Str(v)
										}
										return gopurs_runtime.Array(boxed)
									}(), v1_3_3).UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}()).UnsafePtr)
								unboxed := make([]string, len(arr))
								for i, v := range arr {
									unboxed[i] = v.StrVal()
								}
								return unboxed
							}()
							boxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								boxed[i] = gopurs_runtime.Str(v)
							}
							return gopurs_runtime.Array(boxed)
						}()).StrVal())) + (")")
					}
				end_branch_4:
					return gopurs_runtime.Str(__t4)
				}))
				_ = __local_var_2_2
				var __t5 string
				{
					if v_0.Type == 9 && v_0.IntVal == 4250879068 && v_0.UnsafePtr != nil {
						__t5 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "genericShow'"), gopurs_runtime.Value{Type: 9, IntVal: 4250879068, UnsafePtr: unsafe.Pointer((*Constructor_Main_B)(v_0.UnsafePtr).V0)}).StrVal()
						goto end_branch_5
					} else {

					}
				}
				{
					if v_0.Type == 9 && v_0.IntVal == 4250879068 && v_0.UnsafePtr == nil {
						__t5 = "Z"
						goto end_branch_5
					} else {

					}
				}
				{
					__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }().StrVal()
				}
			end_branch_5:
				return func() gopurs_runtime.Value {
					arr := []string{__t5}
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Str(v)
					}
					return gopurs_runtime.Array(boxed)
				}()
			}))
			_ = __local_var_0_0
			// TAST (Let): __local_var_1_6 -> gopurs_runtime.Value
			__local_var_1_6 := gopurs_runtime.RecordDict1("genericShow'", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
				// TAST (Let): v1_2_7 -> gopurs_runtime.Value
				v1_2_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_0_0, "genericShowArgs"), v_1)
				_ = v1_2_7
				var __t8 string
				{
					if (gopurs_runtime.Int(int64(gopurs_runtime.ArrayLength(v1_2_7))).IntVal) == (0) {
						__t8 = "C"
						goto end_branch_8
					} else {

					}
				}
				{
					__t8 = (("(") + (gopurs_runtime.Apply2(Get_Data_Show_Generic_intercalate(), gopurs_runtime.Str(" "), func() gopurs_runtime.Value {
						arr := func() []string {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(Get_Data_Semigroup_concatArray(), func() gopurs_runtime.Value {
									arr := []string{"C"}
									boxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										boxed[i] = gopurs_runtime.Str(v)
									}
									return gopurs_runtime.Array(boxed)
								}(), v1_2_7).UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}()).UnsafePtr)
							unboxed := make([]string, len(arr))
							for i, v := range arr {
								unboxed[i] = v.StrVal()
							}
							return unboxed
						}()
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Str(v)
						}
						return gopurs_runtime.Array(boxed)
					}()).StrVal())) + (")")
				}
			end_branch_8:
				return gopurs_runtime.Str(__t8)
			}))
			_ = __local_var_1_6
			return gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer((&Constructor_Data_Show_Show{1, gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_6, "genericShow'"), gopurs_runtime.Value{Type: 9, IntVal: 4250879068, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Main_B](x_2))}).StrVal())
			})}))}
		}()
	})
	return cache_Main_showC
}

var cache_Main_genericB2_ gopurs_runtime.Value
var once_Main_genericB2_ sync.Once

func Get_Main_genericB2_() gopurs_runtime.Value {
	once_Main_genericB2_.Do(func() {
		cache_Main_genericB2_ = gopurs_runtime.Value{Type: 9, IntVal: 1921946594, UnsafePtr: unsafe.Pointer((&Constructor_Data_Generic_Rep_Generic{1, gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t0 gopurs_runtime.Value
			{
				if x_0.Type == 9 && x_0.IntVal == 4102534158 && x_0.UnsafePtr != nil {
					__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3478632216, UnsafePtr: unsafe.Pointer((&Constructor_Data_Generic_Rep_Inl{1, (*Constructor_Main_B2)(x_0.UnsafePtr).V0}))}
					goto end_branch_0
				} else {

				}
			}
			{
				if x_0.Type == 9 && x_0.IntVal == 4102534158 && x_0.UnsafePtr == nil {
					__t0 = gopurs_runtime.Value{Type: 9, IntVal: 492034566, UnsafePtr: unsafe.Pointer((&Constructor_Data_Generic_Rep_Inr{1, gopurs_runtime.Value{Type: 9, IntVal: int64(1454898258), UnsafePtr: nil}}))}
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
			var __t1 *Constructor_Main_B2
			{
				if x_0.Type == 9 && x_0.IntVal == 3478632216 {
					__t1 = (&Constructor_Main_B2{1, (*Constructor_Data_Generic_Rep_Inl)(x_0.UnsafePtr).V0})
					goto end_branch_1
				} else {

				}
			}
			{
				if x_0.Type == 9 && x_0.IntVal == 492034566 {
					__t1 = (*Constructor_Main_B2)(nil)
					goto end_branch_1
				} else {

				}
			}
			{
				__t1 = gopurs_runtime.CoerceToStruct[Constructor_Main_B2](func() gopurs_runtime.Value { panic("Failed pattern match") }())
			}
		end_branch_1:
			return gopurs_runtime.Value{Type: 9, IntVal: 4102534158, UnsafePtr: unsafe.Pointer(__t1)}
		})}))}
	})
	return cache_Main_genericB2_
}

var cache_Main_showB2 gopurs_runtime.Value
var once_Main_showB2 sync.Once

func Get_Main_showB2() gopurs_runtime.Value {
	once_Main_showB2.Do(func() {
		cache_Main_showB2 = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer((&Constructor_Data_Show_Show{1, gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
			__local_var_1_0 := gopurs_runtime.RecordDict1("genericShowArgs", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return func() gopurs_runtime.Value {
					arr := []string{gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show](Get_Main_showC2()).V0), v_1).StrVal()}
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Str(v)
					}
					return gopurs_runtime.Array(boxed)
				}()
			}))
			_ = __local_var_1_0
			// TAST (Let): __local_var_2_1 -> gopurs_runtime.Value
			__local_var_2_1 := gopurs_runtime.RecordDict1("genericShow'", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
				// TAST (Let): v1_3_2 -> gopurs_runtime.Value
				v1_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "genericShowArgs"), v_2)
				_ = v1_3_2
				var __t3 string
				{
					if (gopurs_runtime.Int(int64(gopurs_runtime.ArrayLength(v1_3_2))).IntVal) == (0) {
						__t3 = "B2"
						goto end_branch_3
					} else {

					}
				}
				{
					__t3 = (("(") + (gopurs_runtime.Apply2(Get_Data_Show_Generic_intercalate(), gopurs_runtime.Str(" "), func() gopurs_runtime.Value {
						arr := func() []string {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(Get_Data_Semigroup_concatArray(), func() gopurs_runtime.Value {
									arr := []string{"B2"}
									boxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										boxed[i] = gopurs_runtime.Str(v)
									}
									return gopurs_runtime.Array(boxed)
								}(), v1_3_2).UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}()).UnsafePtr)
							unboxed := make([]string, len(arr))
							for i, v := range arr {
								unboxed[i] = v.StrVal()
							}
							return unboxed
						}()
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Str(v)
						}
						return gopurs_runtime.Array(boxed)
					}()).StrVal())) + (")")
				}
			end_branch_3:
				return gopurs_runtime.Str(__t3)
			}))
			_ = __local_var_2_1
			var __t6 string
			{
				var __t_tag_4 *Constructor_Main_B2 = gopurs_runtime.CoerceToStruct[Constructor_Main_B2](x_0)
				if __t_tag_4 != nil {
					__t6 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "genericShow'"), (*Constructor_Main_B2)(x_0.UnsafePtr).V0).StrVal()
					goto end_branch_6
				} else {

				}
			}
			{
				var __t_tag_5 *Constructor_Main_B2 = gopurs_runtime.CoerceToStruct[Constructor_Main_B2](x_0)
				if __t_tag_5 == nil {
					__t6 = "Z2"
					goto end_branch_6
				} else {

				}
			}
			{
				__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }().StrVal()
			}
		end_branch_6:
			return gopurs_runtime.Str(__t6)
		})}))}
	})
	return cache_Main_showB2
}

var cache_Main_showA2 gopurs_runtime.Value
var once_Main_showA2 sync.Once

func Get_Main_showA2() gopurs_runtime.Value {
	once_Main_showA2.Do(func() {
		cache_Main_showA2 = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer((&Constructor_Data_Show_Show{1, gopurs_runtime.Func(func(record_0 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
			__local_var_1_0 := gopurs_runtime.RecordDict1("genericShowArgs", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return func() gopurs_runtime.Value {
					arr := []string{gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show](Get_Main_showC2()).V0), v_1).StrVal()}
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Str(v)
					}
					return gopurs_runtime.Array(boxed)
				}()
			}))
			_ = __local_var_1_0
			// TAST (Let): __local_var_2_1 -> gopurs_runtime.Value
			__local_var_2_1 := gopurs_runtime.RecordDict1("genericShow'", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
				// TAST (Let): v1_3_2 -> gopurs_runtime.Value
				v1_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "genericShowArgs"), v_2)
				_ = v1_3_2
				var __t3 string
				{
					if (gopurs_runtime.Int(int64(gopurs_runtime.ArrayLength(v1_3_2))).IntVal) == (0) {
						__t3 = "B2"
						goto end_branch_3
					} else {

					}
				}
				{
					__t3 = (("(") + (gopurs_runtime.Apply2(Get_Data_Show_Generic_intercalate(), gopurs_runtime.Str(" "), func() gopurs_runtime.Value {
						arr := func() []string {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(Get_Data_Semigroup_concatArray(), func() gopurs_runtime.Value {
									arr := []string{"B2"}
									boxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										boxed[i] = gopurs_runtime.Str(v)
									}
									return gopurs_runtime.Array(boxed)
								}(), v1_3_2).UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}()).UnsafePtr)
							unboxed := make([]string, len(arr))
							for i, v := range arr {
								unboxed[i] = v.StrVal()
							}
							return unboxed
						}()
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Str(v)
						}
						return gopurs_runtime.Array(boxed)
					}()).StrVal())) + (")")
				}
			end_branch_3:
				return gopurs_runtime.Str(__t3)
			}))
			_ = __local_var_2_1
			var __t6 string
			{
				var __t_tag_4 gopurs_runtime.Value = gopurs_runtime.RecordGet(record_0, "x")
				if __t_tag_4.Type == 9 && __t_tag_4.IntVal == 4102534158 && __t_tag_4.UnsafePtr != nil {
					__t6 = (("{ x: ") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "genericShow'"), (*Constructor_Main_B2)(gopurs_runtime.RecordGet(record_0, "x").UnsafePtr).V0).StrVal())) + (" }")
					goto end_branch_6
				} else {

				}
			}
			{
				var __t_tag_5 gopurs_runtime.Value = gopurs_runtime.RecordGet(record_0, "x")
				if __t_tag_5.Type == 9 && __t_tag_5.IntVal == 4102534158 && __t_tag_5.UnsafePtr == nil {
					__t6 = "{ x: Z2 }"
					goto end_branch_6
				} else {

				}
			}
			{
				__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }().StrVal()
			}
		end_branch_6:
			return gopurs_runtime.Str(__t6)
		})}))}
	})
	return cache_Main_showA2
}

var cache_Main_showC2 gopurs_runtime.Value
var once_Main_showC2 sync.Once

func Get_Main_showC2() gopurs_runtime.Value {
	once_Main_showC2.Do(func() {
		cache_Main_showC2 = func() gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 -> gopurs_runtime.Value
			__local_var_0_0 := gopurs_runtime.RecordDict1("genericShowArgs", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
				return func() gopurs_runtime.Value {
					arr := []string{gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show](Get_Main_showA2()).V0), v_0).StrVal()}
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Str(v)
					}
					return gopurs_runtime.Array(boxed)
				}()
			}))
			_ = __local_var_0_0
			// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
			__local_var_1_1 := gopurs_runtime.RecordDict1("genericShow'", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
				// TAST (Let): v1_2_2 -> gopurs_runtime.Value
				v1_2_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_0_0, "genericShowArgs"), v_1)
				_ = v1_2_2
				var __t3 string
				{
					if (gopurs_runtime.Int(int64(gopurs_runtime.ArrayLength(v1_2_2))).IntVal) == (0) {
						__t3 = "C2"
						goto end_branch_3
					} else {

					}
				}
				{
					__t3 = (("(") + (gopurs_runtime.Apply2(Get_Data_Show_Generic_intercalate(), gopurs_runtime.Str(" "), func() gopurs_runtime.Value {
						arr := func() []string {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(Get_Data_Semigroup_concatArray(), func() gopurs_runtime.Value {
									arr := []string{"C2"}
									boxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										boxed[i] = gopurs_runtime.Str(v)
									}
									return gopurs_runtime.Array(boxed)
								}(), v1_2_2).UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}()).UnsafePtr)
							unboxed := make([]string, len(arr))
							for i, v := range arr {
								unboxed[i] = v.StrVal()
							}
							return unboxed
						}()
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Str(v)
						}
						return gopurs_runtime.Array(boxed)
					}()).StrVal())) + (")")
				}
			end_branch_3:
				return gopurs_runtime.Str(__t3)
			}))
			_ = __local_var_1_1
			return gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer((&Constructor_Data_Show_Show{1, gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "genericShow'"), x_2).StrVal())
			})}))}
		}()
	})
	return cache_Main_showC2
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

type Constructor_Main_B2 struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

type Constructor_Main_Z2 struct {
	Rc uint32
}

type Constructor_Main_B struct {
	Rc uint32
	V0 *Constructor_Main_B
}

type Constructor_Main_Z struct {
	Rc uint32
}

func Call_Main_genericShowConstructor(dictIsSymbol_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictIsSymbol_0 gopurs_runtime.Value = dictIsSymbol_0_loop
	_ = dictIsSymbol_0
	return gopurs_runtime.Value{Type: 9, IntVal: 2730968613, UnsafePtr: unsafe.Pointer((&Constructor_Data_Show_Generic_GenericShow{1, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIsSymbol_0, "reflectSymbol"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}).StrVal())
	})}))}
}

func Call_Main_A2(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_C2(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_A(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_C(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}
