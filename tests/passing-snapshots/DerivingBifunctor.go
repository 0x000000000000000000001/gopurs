package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_bifoldl gopurs_runtime.Value
var once_Main_bifoldl sync.Once

func Get_Main_bifoldl() gopurs_runtime.Value {
	once_Main_bifoldl.Do(func() {
		cache_Main_bifoldl = gopurs_runtime.Func4(func(f_0_box gopurs_runtime.Value, g_1_box gopurs_runtime.Value, z_2_box gopurs_runtime.Value, v_3_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_bifoldl(f_0_box, g_1_box, z_2_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](v_3_box))
		})
	})
	return cache_Main_bifoldl
}

var cache_Main_bifoldl1 gopurs_runtime.Value
var once_Main_bifoldl1 sync.Once

func Get_Main_bifoldl1() gopurs_runtime.Value {
	once_Main_bifoldl1.Do(func() {
		cache_Main_bifoldl1 = gopurs_runtime.Func4(func(f_0_box gopurs_runtime.Value, g_1_box gopurs_runtime.Value, z_2_box gopurs_runtime.Value, v_3_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_bifoldl1(f_0_box, g_1_box, z_2_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](v_3_box))
		})
	})
	return cache_Main_bifoldl1
}

var cache_Main_bifoldr gopurs_runtime.Value
var once_Main_bifoldr sync.Once

func Get_Main_bifoldr() gopurs_runtime.Value {
	once_Main_bifoldr.Do(func() {
		cache_Main_bifoldr = gopurs_runtime.Func4(func(f_0_box gopurs_runtime.Value, g_1_box gopurs_runtime.Value, z_2_box gopurs_runtime.Value, v_3_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_bifoldr(f_0_box, g_1_box, z_2_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](v_3_box))
		})
	})
	return cache_Main_bifoldr
}

var cache_Main_bifoldr1 gopurs_runtime.Value
var once_Main_bifoldr1 sync.Once

func Get_Main_bifoldr1() gopurs_runtime.Value {
	once_Main_bifoldr1.Do(func() {
		cache_Main_bifoldr1 = gopurs_runtime.Func4(func(f_0_box gopurs_runtime.Value, g_1_box gopurs_runtime.Value, z_2_box gopurs_runtime.Value, v_3_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_bifoldr1(f_0_box, g_1_box, z_2_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](v_3_box))
		})
	})
	return cache_Main_bifoldr1
}

var cache_Main_bifoldMap gopurs_runtime.Value
var once_Main_bifoldMap sync.Once

func Get_Main_bifoldMap() gopurs_runtime.Value {
	once_Main_bifoldMap.Do(func() {
		cache_Main_bifoldMap = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_bifoldMap(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](dictMonoid_0_box))
		})
	})
	return cache_Main_bifoldMap
}

var cache_Main_identity gopurs_runtime.Value
var once_Main_identity sync.Once

func Get_Main_identity() gopurs_runtime.Value {
	once_Main_identity.Do(func() {
		cache_Main_identity = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_identity(x_0_box)
		})
	})
	return cache_Main_identity
}

var cache_Main_identity1 gopurs_runtime.Value
var once_Main_identity1 sync.Once

func Get_Main_identity1() gopurs_runtime.Value {
	once_Main_identity1.Do(func() {
		cache_Main_identity1 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_identity1(x_0_box)
		})
	})
	return cache_Main_identity1
}

var cache_Main_Test0 gopurs_runtime.Value
var once_Main_Test0 sync.Once

func Get_Main_Test0() gopurs_runtime.Value {
	once_Main_Test0.Do(func() {
		cache_Main_Test0 = gopurs_runtime.Value{Type: 9, IntVal: 2074462008, UnsafePtr: unsafe.Pointer(nil)}
	})
	return cache_Main_Test0
}

var cache_Main_Test1 gopurs_runtime.Value
var once_Main_Test1 sync.Once

func Get_Main_Test1() gopurs_runtime.Value {
	once_Main_Test1.Do(func() {
		cache_Main_Test1 = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 3720114489, UnsafePtr: unsafe.Pointer((&Constructor_Main_Test1{1, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(value0.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						unboxed[i] = v
					}
					return unboxed
				}(), value1}))}
			})
		})
	})
	return cache_Main_Test1
}

var cache_Main_Test2 gopurs_runtime.Value
var once_Main_Test2 sync.Once

func Get_Main_Test2() gopurs_runtime.Value {
	once_Main_Test2.Do(func() {
		cache_Main_Test2 = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 2375191994, UnsafePtr: unsafe.Pointer((&Constructor_Main_Test2{1, value0.IntVal, value1}))}
			})
		})
	})
	return cache_Main_Test2
}

var cache_Main_Test3 gopurs_runtime.Value
var once_Main_Test3 sync.Once

func Get_Main_Test3() gopurs_runtime.Value {
	once_Main_Test3.Do(func() {
		cache_Main_Test3 = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(value2 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(value3 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Value{Type: 9, IntVal: 227416251, UnsafePtr: unsafe.Pointer((&Constructor_Main_Test3{1, value0.IntVal, value1, value2, value3}))}
					})
				})
			})
		})
	})
	return cache_Main_Test3
}

var cache_Main_Test4 gopurs_runtime.Value
var once_Main_Test4 sync.Once

func Get_Main_Test4() gopurs_runtime.Value {
	once_Main_Test4.Do(func() {
		cache_Main_Test4 = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 3712677948, UnsafePtr: unsafe.Pointer((&Constructor_Main_Test4{1, func() []*Constructor_Data_Tuple_Tuple {
					arr := *(*[]gopurs_runtime.Value)(value0.UnsafePtr)
					unboxed := make([]*Constructor_Data_Tuple_Tuple, len(arr))
					for i, v := range arr {
						unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](v)
					}
					return unboxed
				}(), gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](value1)}))}
			})
		})
	})
	return cache_Main_Test4
}

var cache_Main_Test5 gopurs_runtime.Value
var once_Main_Test5 sync.Once

func Get_Main_Test5() gopurs_runtime.Value {
	once_Main_Test5.Do(func() {
		cache_Main_Test5 = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1063363133, UnsafePtr: unsafe.Pointer((&Constructor_Main_Test5{1, value0}))}
		})
	})
	return cache_Main_Test5
}

var cache_Main_FromProAndContra gopurs_runtime.Value
var once_Main_FromProAndContra sync.Once

func Get_Main_FromProAndContra() gopurs_runtime.Value {
	once_Main_FromProAndContra.Do(func() {
		cache_Main_FromProAndContra = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 2092050667, UnsafePtr: unsafe.Pointer((&Constructor_Main_FromProAndContra{1, value0, value1}))}
			})
		})
	})
	return cache_Main_FromProAndContra
}

var cache_Main_bifunctorTest gopurs_runtime.Value
var once_Main_bifunctorTest sync.Once

func Get_Main_bifunctorTest() gopurs_runtime.Value {
	once_Main_bifunctorTest.Do(func() {
		cache_Main_bifunctorTest = gopurs_runtime.Func(func(dictBifunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_bifunctorTest(dictBifunctor_0_box)
		})
	})
	return cache_Main_bifunctorTest
}

var cache_Main_bifunctorFromProAndContra gopurs_runtime.Value
var once_Main_bifunctorFromProAndContra sync.Once

func Get_Main_bifunctorFromProAndContra() gopurs_runtime.Value {
	once_Main_bifunctorFromProAndContra.Do(func() {
		cache_Main_bifunctorFromProAndContra = gopurs_runtime.Value{Type: 9, IntVal: 4141114362, UnsafePtr: unsafe.Pointer((&Constructor_Data_Bifunctor_Bifunctor{1, gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(m_2 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Value{Type: 9, IntVal: 2092050667, UnsafePtr: unsafe.Pointer((&Constructor_Main_FromProAndContra{1, gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Apply((*Constructor_Main_FromProAndContra)(m_2.UnsafePtr).V0, gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Apply(x_3, gopurs_runtime.Apply(f_0, x_4))
						}))
					}), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Apply((*Constructor_Main_FromProAndContra)(m_2.UnsafePtr).V1, gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Apply(x_3, gopurs_runtime.Apply(g_1, x_4))
						}))
					})}))}
				})
			})
		})}))}
	})
	return cache_Main_bifunctorFromProAndContra
}

var cache_Main_bifoldableTest gopurs_runtime.Value
var once_Main_bifoldableTest sync.Once

func Get_Main_bifoldableTest() gopurs_runtime.Value {
	once_Main_bifoldableTest.Do(func() {
		cache_Main_bifoldableTest = gopurs_runtime.Func(func(dictBifoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_bifoldableTest(dictBifoldable_0_box)
		})
	})
	return cache_Main_bifoldableTest
}

var cache_Main_bitraversableTest gopurs_runtime.Value
var once_Main_bitraversableTest sync.Once

func Get_Main_bitraversableTest() gopurs_runtime.Value {
	once_Main_bitraversableTest.Do(func() {
		cache_Main_bitraversableTest = gopurs_runtime.Func(func(dictBitraversable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_bitraversableTest(dictBitraversable_0_box)
		})
	})
	return cache_Main_bitraversableTest
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

type Constructor_Main_Test0 struct {
	Rc uint32
}

type Constructor_Main_Test1 struct {
	Rc uint32
	V0 []gopurs_runtime.Value
	V1 gopurs_runtime.Value
}

type Constructor_Main_Test2 struct {
	Rc uint32
	V0 int64
	V1 gopurs_runtime.Value
}

type Constructor_Main_Test3 struct {
	Rc uint32
	V0 int64
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
	V3 gopurs_runtime.Value
}

type Constructor_Main_Test4 struct {
	Rc uint32
	V0 []*Constructor_Data_Tuple_Tuple
	V1 *Constructor_Data_Tuple_Tuple
}

type Constructor_Main_Test5 struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

type Constructor_Main_FromProAndContra struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}

func Call_Main_bifoldl(f_0_loop gopurs_runtime.Value, g_1_loop gopurs_runtime.Value, z_2_loop gopurs_runtime.Value, v_3_loop *Constructor_Data_Tuple_Tuple) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	var g_1 gopurs_runtime.Value = g_1_loop
	_ = g_1
	var z_2 gopurs_runtime.Value = z_2_loop
	_ = z_2
	var v_3 *Constructor_Data_Tuple_Tuple = v_3_loop
	_ = v_3
	return gopurs_runtime.Apply2(g_1, gopurs_runtime.Apply2(f_0, z_2, (v_3).V0), (v_3).V1)
}

func Call_Main_bifoldl1(f_0_loop gopurs_runtime.Value, g_1_loop gopurs_runtime.Value, z_2_loop gopurs_runtime.Value, v_3_loop *Constructor_Data_Tuple_Tuple) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	var g_1 gopurs_runtime.Value = g_1_loop
	_ = g_1
	var z_2 gopurs_runtime.Value = z_2_loop
	_ = z_2
	var v_3 *Constructor_Data_Tuple_Tuple = v_3_loop
	_ = v_3
	return gopurs_runtime.Apply2(g_1, gopurs_runtime.Apply2(f_0, z_2, (v_3).V0), (v_3).V1)
}

func Call_Main_bifoldr(f_0_loop gopurs_runtime.Value, g_1_loop gopurs_runtime.Value, z_2_loop gopurs_runtime.Value, v_3_loop *Constructor_Data_Tuple_Tuple) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	var g_1 gopurs_runtime.Value = g_1_loop
	_ = g_1
	var z_2 gopurs_runtime.Value = z_2_loop
	_ = z_2
	var v_3 *Constructor_Data_Tuple_Tuple = v_3_loop
	_ = v_3
	return gopurs_runtime.Apply2(f_0, (v_3).V0, gopurs_runtime.Apply2(g_1, (v_3).V1, z_2))
}

func Call_Main_bifoldr1(f_0_loop gopurs_runtime.Value, g_1_loop gopurs_runtime.Value, z_2_loop gopurs_runtime.Value, v_3_loop *Constructor_Data_Tuple_Tuple) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	var g_1 gopurs_runtime.Value = g_1_loop
	_ = g_1
	var z_2 gopurs_runtime.Value = z_2_loop
	_ = z_2
	var v_3 *Constructor_Data_Tuple_Tuple = v_3_loop
	_ = v_3
	return gopurs_runtime.Apply2(f_0, (v_3).V0, gopurs_runtime.Apply2(g_1, (v_3).V1, z_2))
}

func Call_Main_bifoldMap(dictMonoid_0_loop *Constructor_Data_Monoid_Monoid) gopurs_runtime.Value {
	var dictMonoid_0 *Constructor_Data_Monoid_Monoid = dictMonoid_0_loop
	_ = dictMonoid_0
	// TAST (Let): Semigroup0_1_0 shape=App(Other) expectedFromAst=*Constructor_Data_Semigroup_Semigroup actual=*Constructor_Data_Semigroup_Semigroup bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar m)])
	Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonoid_0.V0), gopurs_runtime.Value{}))
	_ = Semigroup0_1_0
	return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(g_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_0.V0), gopurs_runtime.Apply(f_2, (*Constructor_Data_Tuple_Tuple)(v_4.UnsafePtr).V0), gopurs_runtime.Apply(g_3, (*Constructor_Data_Tuple_Tuple)(v_4.UnsafePtr).V1))
			})
		})
	})
}

func Call_Main_identity(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_identity1(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_bifunctorTest(dictBifunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictBifunctor_0 gopurs_runtime.Value = dictBifunctor_0_loop
	_ = dictBifunctor_0
	return gopurs_runtime.Value{Type: 9, IntVal: 4141114362, UnsafePtr: unsafe.Pointer((&Constructor_Data_Bifunctor_Bifunctor{1, gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(g_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t0 gopurs_runtime.Value
				{
					if m_3.Type == 9 && m_3.IntVal == 2074462008 {
						__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2074462008, UnsafePtr: unsafe.Pointer(nil)}
						goto end_branch_0
					} else {

					}
				}
				{
					if m_3.Type == 9 && m_3.IntVal == 3720114489 {
						__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3720114489, UnsafePtr: unsafe.Pointer((&Constructor_Main_Test1{1, func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
								arr_val_arrayMap6 := gopurs_runtime.Array((*Constructor_Main_Test1)(m_3.UnsafePtr).V0)
								_ = arr_val_arrayMap6
								arr_go_arrayMap6 := (*[]gopurs_runtime.Value)(arr_val_arrayMap6.UnsafePtr)
								_ = arr_go_arrayMap6
								res_go_arrayMap6 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap6))
								_ = res_go_arrayMap6
								for i_arrayMap6, v_arrayMap6 := range *arr_go_arrayMap6 {
									res_go_arrayMap6[i_arrayMap6] = gopurs_runtime.Apply(f_1, v_arrayMap6)
								}
								return gopurs_runtime.Array(res_go_arrayMap6)
							}().UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}(), gopurs_runtime.Apply(g_2, (*Constructor_Main_Test1)(m_3.UnsafePtr).V1)}))}
						goto end_branch_0
					} else {

					}
				}
				{
					if m_3.Type == 9 && m_3.IntVal == 2375191994 {
						__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2375191994, UnsafePtr: unsafe.Pointer((&Constructor_Main_Test2{1, (*Constructor_Main_Test2)(m_3.UnsafePtr).V0, (*Constructor_Main_Test2)(m_3.UnsafePtr).V1}))}
						goto end_branch_0
					} else {

					}
				}
				{
					if m_3.Type == 9 && m_3.IntVal == 227416251 {
						__t0 = gopurs_runtime.Value{Type: 9, IntVal: 227416251, UnsafePtr: unsafe.Pointer((&Constructor_Main_Test3{1, (*Constructor_Main_Test3)(m_3.UnsafePtr).V0, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictBifunctor_0, "bimap"), f_1, g_2, (*Constructor_Main_Test3)(m_3.UnsafePtr).V1), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictBifunctor_0, "bimap"), f_1, gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
							return x_4
						}), (*Constructor_Main_Test3)(m_3.UnsafePtr).V2), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictBifunctor_0, "bimap"), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
							return x_4
						}), g_2, (*Constructor_Main_Test3)(m_3.UnsafePtr).V3)}))}
						goto end_branch_0
					} else {

					}
				}
				{
					if m_3.Type == 9 && m_3.IntVal == 3712677948 {
						__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3712677948, UnsafePtr: unsafe.Pointer((&Constructor_Main_Test4{1, func() []*Constructor_Data_Tuple_Tuple {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
									arr_val_arrayMap6 := func() gopurs_runtime.Value {
										arr := (*Constructor_Main_Test4)(m_3.UnsafePtr).V0
										boxed := make([]gopurs_runtime.Value, len(arr))
										for i, v := range arr {
											boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)}
										}
										return gopurs_runtime.Array(boxed)
									}()
									_ = arr_val_arrayMap6
									arr_go_arrayMap6 := (*[]gopurs_runtime.Value)(arr_val_arrayMap6.UnsafePtr)
									_ = arr_go_arrayMap6
									res_go_arrayMap6 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap6))
									_ = res_go_arrayMap6
									for i_arrayMap6, v_arrayMap6 := range *arr_go_arrayMap6 {
										res_go_arrayMap6[i_arrayMap6] = gopurs_runtime.Apply(gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer((&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_1, (*Constructor_Data_Tuple_Tuple)(v_4.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v_4.UnsafePtr).V1}))}
										}), v_arrayMap6)
									}
									return gopurs_runtime.Array(res_go_arrayMap6)
								}().UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}()).UnsafePtr)
							unboxed := make([]*Constructor_Data_Tuple_Tuple, len(arr))
							for i, v := range arr {
								unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](v)
							}
							return unboxed
						}(), gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer((&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(g_2, ((*Constructor_Main_Test4)(m_3.UnsafePtr).V1).V0), ((*Constructor_Main_Test4)(m_3.UnsafePtr).V1).V1}))})}))}
						goto end_branch_0
					} else {

					}
				}
				{
					if m_3.Type == 9 && m_3.IntVal == 1063363133 {
						__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1063363133, UnsafePtr: unsafe.Pointer((&Constructor_Main_Test5{1, func() gopurs_runtime.Value {
							origVal := (*Constructor_Main_Test5)(m_3.UnsafePtr).V0
							if origVal.Type != gopurs_runtime.TypeRecord1 {
								return gopurs_runtime.RecordUpdateDict(origVal, []string{"nested"}, []gopurs_runtime.Value{func() gopurs_runtime.Value {
									arr := func() []gopurs_runtime.Value {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
											arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
												arr_val_arrayMap7 := func() gopurs_runtime.Value {
													arr := func() []gopurs_runtime.Value {
														arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_Test5)(m_3.UnsafePtr).V0, "nested").UnsafePtr)
														unboxed := make([]gopurs_runtime.Value, len(arr))
														for i, v := range arr {
															unboxed[i] = v
														}
														return unboxed
													}()
													boxed := make([]gopurs_runtime.Value, len(arr))
													for i, v := range arr {
														boxed[i] = v
													}
													return gopurs_runtime.Array(boxed)
												}()
												_ = arr_val_arrayMap7
												arr_go_arrayMap7 := (*[]gopurs_runtime.Value)(arr_val_arrayMap7.UnsafePtr)
												_ = arr_go_arrayMap7
												res_go_arrayMap7 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap7))
												_ = res_go_arrayMap7
												for i_arrayMap7, v_arrayMap7 := range *arr_go_arrayMap7 {
													res_go_arrayMap7[i_arrayMap7] = gopurs_runtime.Apply(gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
														return gopurs_runtime.RecordUpdate1(v1_4, "x", gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictBifunctor_0, "bimap"), gopurs_runtime.Func(func(v2_5 gopurs_runtime.Value) gopurs_runtime.Value {
															return gopurs_runtime.RecordUpdate1(v2_5, "a", gopurs_runtime.Apply(f_1, gopurs_runtime.RecordGet(v2_5, "a")))
														}), gopurs_runtime.Func(func(v2_5 gopurs_runtime.Value) gopurs_runtime.Value {
															return gopurs_runtime.RecordUpdate1(v2_5, "b", gopurs_runtime.Apply(g_2, gopurs_runtime.RecordGet(v2_5, "b")))
														}), gopurs_runtime.RecordGet(v1_4, "x")))
													}), v_arrayMap7)
												}
												return gopurs_runtime.Array(res_go_arrayMap7)
											}().UnsafePtr)
											unboxed := make([]gopurs_runtime.Value, len(arr))
											for i, v := range arr {
												unboxed[i] = v
											}
											return unboxed
										}()).UnsafePtr)
										unboxed := make([]gopurs_runtime.Value, len(arr))
										for i, v := range arr {
											unboxed[i] = v
										}
										return unboxed
									}()
									boxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										boxed[i] = v
									}
									return gopurs_runtime.Array(boxed)
								}()})
							}
							clone := *((*gopurs_runtime.RecordData1)(origVal.UnsafePtr))
							clone.V0 = func() gopurs_runtime.Value {
								arr := func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
										arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
											arr_val_arrayMap7 := func() gopurs_runtime.Value {
												arr := func() []gopurs_runtime.Value {
													arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_Test5)(m_3.UnsafePtr).V0, "nested").UnsafePtr)
													unboxed := make([]gopurs_runtime.Value, len(arr))
													for i, v := range arr {
														unboxed[i] = v
													}
													return unboxed
												}()
												boxed := make([]gopurs_runtime.Value, len(arr))
												for i, v := range arr {
													boxed[i] = v
												}
												return gopurs_runtime.Array(boxed)
											}()
											_ = arr_val_arrayMap7
											arr_go_arrayMap7 := (*[]gopurs_runtime.Value)(arr_val_arrayMap7.UnsafePtr)
											_ = arr_go_arrayMap7
											res_go_arrayMap7 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap7))
											_ = res_go_arrayMap7
											for i_arrayMap7, v_arrayMap7 := range *arr_go_arrayMap7 {
												res_go_arrayMap7[i_arrayMap7] = gopurs_runtime.Apply(gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
													return gopurs_runtime.RecordUpdate1(v1_4, "x", gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictBifunctor_0, "bimap"), gopurs_runtime.Func(func(v2_5 gopurs_runtime.Value) gopurs_runtime.Value {
														return gopurs_runtime.RecordUpdate1(v2_5, "a", gopurs_runtime.Apply(f_1, gopurs_runtime.RecordGet(v2_5, "a")))
													}), gopurs_runtime.Func(func(v2_5 gopurs_runtime.Value) gopurs_runtime.Value {
														return gopurs_runtime.RecordUpdate1(v2_5, "b", gopurs_runtime.Apply(g_2, gopurs_runtime.RecordGet(v2_5, "b")))
													}), gopurs_runtime.RecordGet(v1_4, "x")))
												}), v_arrayMap7)
											}
											return gopurs_runtime.Array(res_go_arrayMap7)
										}().UnsafePtr)
										unboxed := make([]gopurs_runtime.Value, len(arr))
										for i, v := range arr {
											unboxed[i] = v
										}
										return unboxed
									}()).UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}()
								boxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									boxed[i] = v
								}
								return gopurs_runtime.Array(boxed)
							}()
							return gopurs_runtime.Value{Type: gopurs_runtime.TypeRecord1, UnsafePtr: unsafe.Pointer(&clone)}
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
			})
		})
	})}))}
}

func Call_Main_bifoldableTest(dictBifoldable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictBifoldable_0 gopurs_runtime.Value = dictBifoldable_0_loop
	_ = dictBifoldable_0
	return gopurs_runtime.Value{Type: 9, IntVal: 4001671834, UnsafePtr: unsafe.Pointer((&Constructor_Data_Bifoldable_Bifoldable{1, gopurs_runtime.Func(func(dictMonoid_1 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): mempty_2_0 shape=Other expectedFromAst=gopurs_runtime.Value actual=gopurs_runtime.Value bindingType=(TypeVar m)
		mempty_2_0 := gopurs_runtime.RecordGet(dictMonoid_1, "mempty")
		_ = mempty_2_0
		// TAST (Let): Semigroup0_3_1 shape=App(Other) expectedFromAst=*Constructor_Data_Semigroup_Semigroup actual=*Constructor_Data_Semigroup_Semigroup bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar m)])
		Semigroup0_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_1, "Semigroup0"), gopurs_runtime.Value{}))
		_ = Semigroup0_3_1
		// TAST (Let): bifoldMap2_4_2 shape=App(Other) expectedFromAst=gopurs_runtime.Value actual=gopurs_runtime.Value bindingType=(Func [(Func [(TypeVar a)] (TypeVar m)), (Func [Int] (TypeVar m)), (TypeApp (TypeVar f) [(TypeVar a), Int])] (TypeVar m))
		bifoldMap2_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBifoldable_0, "bifoldMap"), dictMonoid_1)
		_ = bifoldMap2_4_2
		// TAST (Let): Semigroup0_5_4 shape=App(Other) expectedFromAst=*Constructor_Data_Semigroup_Semigroup actual=*Constructor_Data_Semigroup_Semigroup bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar m)])
		Semigroup0_5_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_1, "Semigroup0"), gopurs_runtime.Value{}))
		_ = Semigroup0_5_4
		// TAST (Let): bifoldMap3_5_3 shape=Let(Abs(Abs(Abs(App(Other))))) expectedFromAst=gopurs_runtime.Value actual=gopurs_runtime.Value bindingType=(Func [(Func [(TypeVar a)] (TypeVar m)), (Func [Int] (TypeVar m)), (ADT ["Data","Tuple","Tuple"] [(TypeVar a), Int])] (TypeVar m))
		bifoldMap3_5_3 := gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(g_7 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_5_4.V0), gopurs_runtime.Apply(f_6, (*Constructor_Data_Tuple_Tuple)(v_8.UnsafePtr).V0), gopurs_runtime.Apply(g_7, (*Constructor_Data_Tuple_Tuple)(v_8.UnsafePtr).V1))
				})
			})
		})
		_ = bifoldMap3_5_3
		// TAST (Let): Semigroup0_6_6 shape=App(Other) expectedFromAst=*Constructor_Data_Semigroup_Semigroup actual=*Constructor_Data_Semigroup_Semigroup bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar m)])
		Semigroup0_6_6 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_1, "Semigroup0"), gopurs_runtime.Value{}))
		_ = Semigroup0_6_6
		// TAST (Let): bifoldMap4_6_5 shape=Let(Abs(Abs(Abs(App(Other))))) expectedFromAst=gopurs_runtime.Value actual=gopurs_runtime.Value bindingType=(Func [(Func [(TypeVar b)] (TypeVar m)), (Func [Int] (TypeVar m)), (ADT ["Data","Tuple","Tuple"] [(TypeVar b), Int])] (TypeVar m))
		bifoldMap4_6_5 := gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(g_8 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_6_6.V0), gopurs_runtime.Apply(f_7, (*Constructor_Data_Tuple_Tuple)(v_9.UnsafePtr).V0), gopurs_runtime.Apply(g_8, (*Constructor_Data_Tuple_Tuple)(v_9.UnsafePtr).V1))
				})
			})
		})
		_ = bifoldMap4_6_5
		return gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(g_8 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(m_9 gopurs_runtime.Value) gopurs_runtime.Value {
					var __t11 gopurs_runtime.Value
					{
						if m_9.Type == 9 && m_9.IntVal == 2074462008 {
							__t11 = mempty_2_0
							goto end_branch_11
						} else {

						}
					}
					{
						if m_9.Type == 9 && m_9.IntVal == 3720114489 {
							// TAST (Let): Semigroup0_10_7 shape=App(Other) expectedFromAst=*Constructor_Data_Semigroup_Semigroup actual=*Constructor_Data_Semigroup_Semigroup bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar m)])
							Semigroup0_10_7 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_1, "Semigroup0"), gopurs_runtime.Value{}))
							_ = Semigroup0_10_7
							__t11 = gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_3_1.V0), gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(x_11 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Func(func(acc_12 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_10_7.V0), gopurs_runtime.Apply(f_7, x_11), acc_12)
								})
							}), gopurs_runtime.RecordGet(dictMonoid_1, "mempty"), gopurs_runtime.Array((*Constructor_Main_Test1)(m_9.UnsafePtr).V0)), gopurs_runtime.Apply(g_8, (*Constructor_Main_Test1)(m_9.UnsafePtr).V1))
							goto end_branch_11
						} else {

						}
					}
					{
						if m_9.Type == 9 && m_9.IntVal == 2375191994 {
							__t11 = mempty_2_0
							goto end_branch_11
						} else {

						}
					}
					{
						if m_9.Type == 9 && m_9.IntVal == 227416251 {
							__t11 = gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_3_1.V0), gopurs_runtime.Apply4(gopurs_runtime.RecordGet(dictBifoldable_0, "bifoldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](dictMonoid_1))}, f_7, g_8, (*Constructor_Main_Test3)(m_9.UnsafePtr).V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_3_1.V0), gopurs_runtime.Apply3(bifoldMap2_4_2, f_7, gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.RecordGet(dictMonoid_1, "mempty")
							}), (*Constructor_Main_Test3)(m_9.UnsafePtr).V2), gopurs_runtime.Apply4(gopurs_runtime.RecordGet(dictBifoldable_0, "bifoldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](dictMonoid_1))}, gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.RecordGet(dictMonoid_1, "mempty")
							}), g_8, (*Constructor_Main_Test3)(m_9.UnsafePtr).V3)))
							goto end_branch_11
						} else {

						}
					}
					{
						if m_9.Type == 9 && m_9.IntVal == 3712677948 {
							// TAST (Let): Semigroup0_10_8 shape=App(Other) expectedFromAst=*Constructor_Data_Semigroup_Semigroup actual=*Constructor_Data_Semigroup_Semigroup bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar m)])
							Semigroup0_10_8 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_1, "Semigroup0"), gopurs_runtime.Value{}))
							_ = Semigroup0_10_8
							// TAST (Let): __local_var_11_9 shape=App(Other) expectedFromAst=gopurs_runtime.Value actual=gopurs_runtime.Value bindingType=(Func [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), Int])] (TypeVar m))
							__local_var_11_9 := gopurs_runtime.Apply2(bifoldMap3_5_3, f_7, gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.RecordGet(dictMonoid_1, "mempty")
							}))
							_ = __local_var_11_9
							__t11 = gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_3_1.V0), gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(x_12 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Func(func(acc_13 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_10_8.V0), gopurs_runtime.Apply(__local_var_11_9, x_12), acc_13)
								})
							}), gopurs_runtime.RecordGet(dictMonoid_1, "mempty"), func() gopurs_runtime.Value {
								arr := (*Constructor_Main_Test4)(m_9.UnsafePtr).V0
								boxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)}
								}
								return gopurs_runtime.Array(boxed)
							}()), gopurs_runtime.Apply3(bifoldMap4_6_5, g_8, gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.RecordGet(dictMonoid_1, "mempty")
							}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer((*Constructor_Main_Test4)(m_9.UnsafePtr).V1)}))
							goto end_branch_11
						} else {

						}
					}
					{
						if m_9.Type == 9 && m_9.IntVal == 1063363133 {
							// TAST (Let): Semigroup0_10_10 shape=App(Other) expectedFromAst=*Constructor_Data_Semigroup_Semigroup actual=*Constructor_Data_Semigroup_Semigroup bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar m)])
							Semigroup0_10_10 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_1, "Semigroup0"), gopurs_runtime.Value{}))
							_ = Semigroup0_10_10
							__t11 = gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(x_11 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Func(func(acc_12 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_10_10.V0), gopurs_runtime.Apply4(gopurs_runtime.RecordGet(dictBifoldable_0, "bifoldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](dictMonoid_1))}, gopurs_runtime.Func(func(v2_13 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Apply(f_7, gopurs_runtime.RecordGet(v2_13, "a"))
									}), gopurs_runtime.Func(func(v2_13 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Apply(g_8, gopurs_runtime.RecordGet(v2_13, "b"))
									}), gopurs_runtime.RecordGet(x_11, "x")), acc_12)
								})
							}), gopurs_runtime.RecordGet(dictMonoid_1, "mempty"), func() gopurs_runtime.Value {
								arr := func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_Test5)(m_9.UnsafePtr).V0, "nested").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}()
								boxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									boxed[i] = v
								}
								return gopurs_runtime.Array(boxed)
							}())
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
			})
		})
	}), gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(g_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(z_3 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
					var __t12 gopurs_runtime.Value
					{
						if m_4.Type == 9 && m_4.IntVal == 2074462008 {
							__t12 = z_3
							goto end_branch_12
						} else {

						}
					}
					{
						if m_4.Type == 9 && m_4.IntVal == 3720114489 {
							__t12 = gopurs_runtime.Apply2(g_2, func() gopurs_runtime.Value {
								arr_val_foldlArray7 := gopurs_runtime.Array((*Constructor_Main_Test1)(m_4.UnsafePtr).V0)
								_ = arr_val_foldlArray7
								res_go_foldlArray7 := z_3
								_ = res_go_foldlArray7
								arr_go_foldlArray7 := (*[]gopurs_runtime.Value)(arr_val_foldlArray7.UnsafePtr)
								_ = arr_go_foldlArray7
								for _, v_foldlArray7 := range *arr_go_foldlArray7 {
									res_go_foldlArray7 = gopurs_runtime.Apply2(f_1, res_go_foldlArray7, v_foldlArray7)
								}
								return res_go_foldlArray7
							}(), (*Constructor_Main_Test1)(m_4.UnsafePtr).V1)
							goto end_branch_12
						} else {

						}
					}
					{
						if m_4.Type == 9 && m_4.IntVal == 2375191994 {
							__t12 = z_3
							goto end_branch_12
						} else {

						}
					}
					{
						if m_4.Type == 9 && m_4.IntVal == 227416251 {
							__t12 = gopurs_runtime.Apply4(gopurs_runtime.RecordGet(dictBifoldable_0, "bifoldl"), Get_Data_Function_go__const(), g_2, gopurs_runtime.Apply4(gopurs_runtime.RecordGet(dictBifoldable_0, "bifoldl"), f_1, Get_Data_Function_go__const(), gopurs_runtime.Apply4(gopurs_runtime.RecordGet(dictBifoldable_0, "bifoldl"), f_1, g_2, z_3, (*Constructor_Main_Test3)(m_4.UnsafePtr).V1), (*Constructor_Main_Test3)(m_4.UnsafePtr).V2), (*Constructor_Main_Test3)(m_4.UnsafePtr).V3)
							goto end_branch_12
						} else {

						}
					}
					{
						if m_4.Type == 9 && m_4.IntVal == 3712677948 {
							__t12 = gopurs_runtime.Apply2(g_2, func() gopurs_runtime.Value {
								arr_val_foldlArray7 := func() gopurs_runtime.Value {
									arr := (*Constructor_Main_Test4)(m_4.UnsafePtr).V0
									boxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)}
									}
									return gopurs_runtime.Array(boxed)
								}()
								_ = arr_val_foldlArray7
								res_go_foldlArray7 := z_3
								_ = res_go_foldlArray7
								arr_go_foldlArray7 := (*[]gopurs_runtime.Value)(arr_val_foldlArray7.UnsafePtr)
								_ = arr_go_foldlArray7
								for _, v_foldlArray7 := range *arr_go_foldlArray7 {
									res_go_foldlArray7 = gopurs_runtime.Apply2(gopurs_runtime.Func(func(z_5 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Apply2(f_1, z_5, (*Constructor_Data_Tuple_Tuple)(v_6.UnsafePtr).V0)
										})
									}), res_go_foldlArray7, v_foldlArray7)
								}
								return res_go_foldlArray7
							}(), ((*Constructor_Main_Test4)(m_4.UnsafePtr).V1).V0)
							goto end_branch_12
						} else {

						}
					}
					{
						if m_4.Type == 9 && m_4.IntVal == 1063363133 {
							__t12 = func() gopurs_runtime.Value {
								arr_val_foldlArray6 := func() gopurs_runtime.Value {
									arr := func() []gopurs_runtime.Value {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_Test5)(m_4.UnsafePtr).V0, "nested").UnsafePtr)
										unboxed := make([]gopurs_runtime.Value, len(arr))
										for i, v := range arr {
											unboxed[i] = v
										}
										return unboxed
									}()
									boxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										boxed[i] = v
									}
									return gopurs_runtime.Array(boxed)
								}()
								_ = arr_val_foldlArray6
								res_go_foldlArray6 := z_3
								_ = res_go_foldlArray6
								arr_go_foldlArray6 := (*[]gopurs_runtime.Value)(arr_val_foldlArray6.UnsafePtr)
								_ = arr_go_foldlArray6
								for _, v_foldlArray6 := range *arr_go_foldlArray6 {
									res_go_foldlArray6 = gopurs_runtime.Apply2(gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Func(func(v2_6 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Apply4(gopurs_runtime.RecordGet(dictBifoldable_0, "bifoldl"), gopurs_runtime.Func(func(v3_7 gopurs_runtime.Value) gopurs_runtime.Value {
												return gopurs_runtime.Func(func(v4_8 gopurs_runtime.Value) gopurs_runtime.Value {
													return gopurs_runtime.Apply2(f_1, v3_7, gopurs_runtime.RecordGet(v4_8, "a"))
												})
											}), gopurs_runtime.Func(func(v3_7 gopurs_runtime.Value) gopurs_runtime.Value {
												return gopurs_runtime.Func(func(v4_8 gopurs_runtime.Value) gopurs_runtime.Value {
													return gopurs_runtime.Apply2(g_2, v3_7, gopurs_runtime.RecordGet(v4_8, "b"))
												})
											}), v1_5, gopurs_runtime.RecordGet(v2_6, "x"))
										})
									}), res_go_foldlArray6, v_foldlArray6)
								}
								return res_go_foldlArray6
							}()
							goto end_branch_12
						} else {

						}
					}
					{
						__t12 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
					}
				end_branch_12:
					return __t12
				})
			})
		})
	}), gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(g_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(z_3 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
					var __t13 gopurs_runtime.Value
					{
						if m_4.Type == 9 && m_4.IntVal == 2074462008 {
							__t13 = z_3
							goto end_branch_13
						} else {

						}
					}
					{
						if m_4.Type == 9 && m_4.IntVal == 3720114489 {
							__t13 = gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), f_1, gopurs_runtime.Apply2(g_2, (*Constructor_Main_Test1)(m_4.UnsafePtr).V1, z_3), gopurs_runtime.Array((*Constructor_Main_Test1)(m_4.UnsafePtr).V0))
							goto end_branch_13
						} else {

						}
					}
					{
						if m_4.Type == 9 && m_4.IntVal == 2375191994 {
							__t13 = z_3
							goto end_branch_13
						} else {

						}
					}
					{
						if m_4.Type == 9 && m_4.IntVal == 227416251 {
							__t13 = gopurs_runtime.Apply4(gopurs_runtime.RecordGet(dictBifoldable_0, "bifoldr"), f_1, g_2, gopurs_runtime.Apply4(gopurs_runtime.RecordGet(dictBifoldable_0, "bifoldr"), f_1, gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
									return x_6
								})
							}), gopurs_runtime.Apply4(gopurs_runtime.RecordGet(dictBifoldable_0, "bifoldr"), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
									return x_6
								})
							}), g_2, z_3, (*Constructor_Main_Test3)(m_4.UnsafePtr).V3), (*Constructor_Main_Test3)(m_4.UnsafePtr).V2), (*Constructor_Main_Test3)(m_4.UnsafePtr).V1)
							goto end_branch_13
						} else {

						}
					}
					{
						if m_4.Type == 9 && m_4.IntVal == 3712677948 {
							__t13 = gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(b_5 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Apply2(f_1, (*Constructor_Data_Tuple_Tuple)(b_5.UnsafePtr).V0, a_6)
								})
							}), gopurs_runtime.Apply2(g_2, ((*Constructor_Main_Test4)(m_4.UnsafePtr).V1).V0, z_3), func() gopurs_runtime.Value {
								arr := (*Constructor_Main_Test4)(m_4.UnsafePtr).V0
								boxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)}
								}
								return gopurs_runtime.Array(boxed)
							}())
							goto end_branch_13
						} else {

						}
					}
					{
						if m_4.Type == 9 && m_4.IntVal == 1063363133 {
							__t13 = gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Func(func(v2_6 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Apply4(gopurs_runtime.RecordGet(dictBifoldable_0, "bifoldr"), gopurs_runtime.Func(func(v3_7 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Func(func(v4_8 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Apply2(f_1, gopurs_runtime.RecordGet(v3_7, "a"), v4_8)
										})
									}), gopurs_runtime.Func(func(v3_7 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Func(func(v4_8 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Apply2(g_2, gopurs_runtime.RecordGet(v3_7, "b"), v4_8)
										})
									}), v2_6, gopurs_runtime.RecordGet(v1_5, "x"))
								})
							}), z_3, func() gopurs_runtime.Value {
								arr := func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_Test5)(m_4.UnsafePtr).V0, "nested").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}()
								boxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									boxed[i] = v
								}
								return gopurs_runtime.Array(boxed)
							}())
							goto end_branch_13
						} else {

						}
					}
					{
						__t13 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
					}
				end_branch_13:
					return __t13
				})
			})
		})
	})}))}
}

func Call_Main_bitraversableTest(dictBitraversable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
bitraversableTest:
	for {
		if false {
			continue bitraversableTest
		}
		var dictBitraversable_0 gopurs_runtime.Value = dictBitraversable_0_loop
		_ = dictBitraversable_0
		// TAST (Let): __local_var_1_1 shape=App(Other) expectedFromAst=gopurs_runtime.Value actual=gopurs_runtime.Value bindingType=Any
		__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBitraversable_0, "Bifunctor0"), gopurs_runtime.Value{})
		_ = __local_var_1_1
		// TAST (Let): bifunctorTest1_1_0 shape=Let(LitRecord) expectedFromAst=*Constructor_Data_Bifunctor_Bifunctor actual=*Constructor_Data_Bifunctor_Bifunctor bindingType=(ADT ["Data","Bifunctor","Bifunctor"] [(ADT ["Main","Test"] [(TypeVar f)])])
		bifunctorTest1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Bifunctor_Bifunctor](gopurs_runtime.RecordDict1("bimap", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(g_3 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
					var __t2 gopurs_runtime.Value
					{
						if m_4.Type == 9 && m_4.IntVal == 2074462008 {
							__t2 = gopurs_runtime.Value{Type: 9, IntVal: 2074462008, UnsafePtr: unsafe.Pointer(nil)}
							goto end_branch_2
						} else {

						}
					}
					{
						if m_4.Type == 9 && m_4.IntVal == 3720114489 {
							__t2 = gopurs_runtime.Value{Type: 9, IntVal: 3720114489, UnsafePtr: unsafe.Pointer((&Constructor_Main_Test1{1, func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
									arr_val_arrayMap8 := gopurs_runtime.Array((*Constructor_Main_Test1)(m_4.UnsafePtr).V0)
									_ = arr_val_arrayMap8
									arr_go_arrayMap8 := (*[]gopurs_runtime.Value)(arr_val_arrayMap8.UnsafePtr)
									_ = arr_go_arrayMap8
									res_go_arrayMap8 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap8))
									_ = res_go_arrayMap8
									for i_arrayMap8, v_arrayMap8 := range *arr_go_arrayMap8 {
										res_go_arrayMap8[i_arrayMap8] = gopurs_runtime.Apply(f_2, v_arrayMap8)
									}
									return gopurs_runtime.Array(res_go_arrayMap8)
								}().UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}(), gopurs_runtime.Apply(g_3, (*Constructor_Main_Test1)(m_4.UnsafePtr).V1)}))}
							goto end_branch_2
						} else {

						}
					}
					{
						if m_4.Type == 9 && m_4.IntVal == 2375191994 {
							__t2 = gopurs_runtime.Value{Type: 9, IntVal: 2375191994, UnsafePtr: unsafe.Pointer((&Constructor_Main_Test2{1, (*Constructor_Main_Test2)(m_4.UnsafePtr).V0, (*Constructor_Main_Test2)(m_4.UnsafePtr).V1}))}
							goto end_branch_2
						} else {

						}
					}
					{
						if m_4.Type == 9 && m_4.IntVal == 227416251 {
							__t2 = gopurs_runtime.Value{Type: 9, IntVal: 227416251, UnsafePtr: unsafe.Pointer((&Constructor_Main_Test3{1, (*Constructor_Main_Test3)(m_4.UnsafePtr).V0, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_1_1, "bimap"), f_2, g_3, (*Constructor_Main_Test3)(m_4.UnsafePtr).V1), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_1_1, "bimap"), f_2, gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
								return x_5
							}), (*Constructor_Main_Test3)(m_4.UnsafePtr).V2), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_1_1, "bimap"), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
								return x_5
							}), g_3, (*Constructor_Main_Test3)(m_4.UnsafePtr).V3)}))}
							goto end_branch_2
						} else {

						}
					}
					{
						if m_4.Type == 9 && m_4.IntVal == 3712677948 {
							__t2 = gopurs_runtime.Value{Type: 9, IntVal: 3712677948, UnsafePtr: unsafe.Pointer((&Constructor_Main_Test4{1, func() []*Constructor_Data_Tuple_Tuple {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
										arr_val_arrayMap8 := func() gopurs_runtime.Value {
											arr := (*Constructor_Main_Test4)(m_4.UnsafePtr).V0
											boxed := make([]gopurs_runtime.Value, len(arr))
											for i, v := range arr {
												boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)}
											}
											return gopurs_runtime.Array(boxed)
										}()
										_ = arr_val_arrayMap8
										arr_go_arrayMap8 := (*[]gopurs_runtime.Value)(arr_val_arrayMap8.UnsafePtr)
										_ = arr_go_arrayMap8
										res_go_arrayMap8 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap8))
										_ = res_go_arrayMap8
										for i_arrayMap8, v_arrayMap8 := range *arr_go_arrayMap8 {
											res_go_arrayMap8[i_arrayMap8] = gopurs_runtime.Apply(gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
												return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer((&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_2, (*Constructor_Data_Tuple_Tuple)(v_5.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v_5.UnsafePtr).V1}))}
											}), v_arrayMap8)
										}
										return gopurs_runtime.Array(res_go_arrayMap8)
									}().UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}()).UnsafePtr)
								unboxed := make([]*Constructor_Data_Tuple_Tuple, len(arr))
								for i, v := range arr {
									unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](v)
								}
								return unboxed
							}(), gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer((&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(g_3, ((*Constructor_Main_Test4)(m_4.UnsafePtr).V1).V0), ((*Constructor_Main_Test4)(m_4.UnsafePtr).V1).V1}))})}))}
							goto end_branch_2
						} else {

						}
					}
					{
						if m_4.Type == 9 && m_4.IntVal == 1063363133 {
							__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1063363133, UnsafePtr: unsafe.Pointer((&Constructor_Main_Test5{1, func() gopurs_runtime.Value {
								origVal := (*Constructor_Main_Test5)(m_4.UnsafePtr).V0
								if origVal.Type != gopurs_runtime.TypeRecord1 {
									return gopurs_runtime.RecordUpdateDict(origVal, []string{"nested"}, []gopurs_runtime.Value{func() gopurs_runtime.Value {
										arr := func() []gopurs_runtime.Value {
											arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
												arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
													arr_val_arrayMap9 := func() gopurs_runtime.Value {
														arr := func() []gopurs_runtime.Value {
															arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_Test5)(m_4.UnsafePtr).V0, "nested").UnsafePtr)
															unboxed := make([]gopurs_runtime.Value, len(arr))
															for i, v := range arr {
																unboxed[i] = v
															}
															return unboxed
														}()
														boxed := make([]gopurs_runtime.Value, len(arr))
														for i, v := range arr {
															boxed[i] = v
														}
														return gopurs_runtime.Array(boxed)
													}()
													_ = arr_val_arrayMap9
													arr_go_arrayMap9 := (*[]gopurs_runtime.Value)(arr_val_arrayMap9.UnsafePtr)
													_ = arr_go_arrayMap9
													res_go_arrayMap9 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap9))
													_ = res_go_arrayMap9
													for i_arrayMap9, v_arrayMap9 := range *arr_go_arrayMap9 {
														res_go_arrayMap9[i_arrayMap9] = gopurs_runtime.Apply(gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
															return gopurs_runtime.RecordUpdate1(v1_5, "x", gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_1_1, "bimap"), gopurs_runtime.Func(func(v2_6 gopurs_runtime.Value) gopurs_runtime.Value {
																return gopurs_runtime.RecordUpdate1(v2_6, "a", gopurs_runtime.Apply(f_2, gopurs_runtime.RecordGet(v2_6, "a")))
															}), gopurs_runtime.Func(func(v2_6 gopurs_runtime.Value) gopurs_runtime.Value {
																return gopurs_runtime.RecordUpdate1(v2_6, "b", gopurs_runtime.Apply(g_3, gopurs_runtime.RecordGet(v2_6, "b")))
															}), gopurs_runtime.RecordGet(v1_5, "x")))
														}), v_arrayMap9)
													}
													return gopurs_runtime.Array(res_go_arrayMap9)
												}().UnsafePtr)
												unboxed := make([]gopurs_runtime.Value, len(arr))
												for i, v := range arr {
													unboxed[i] = v
												}
												return unboxed
											}()).UnsafePtr)
											unboxed := make([]gopurs_runtime.Value, len(arr))
											for i, v := range arr {
												unboxed[i] = v
											}
											return unboxed
										}()
										boxed := make([]gopurs_runtime.Value, len(arr))
										for i, v := range arr {
											boxed[i] = v
										}
										return gopurs_runtime.Array(boxed)
									}()})
								}
								clone := *((*gopurs_runtime.RecordData1)(origVal.UnsafePtr))
								clone.V0 = func() gopurs_runtime.Value {
									arr := func() []gopurs_runtime.Value {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
											arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
												arr_val_arrayMap9 := func() gopurs_runtime.Value {
													arr := func() []gopurs_runtime.Value {
														arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_Test5)(m_4.UnsafePtr).V0, "nested").UnsafePtr)
														unboxed := make([]gopurs_runtime.Value, len(arr))
														for i, v := range arr {
															unboxed[i] = v
														}
														return unboxed
													}()
													boxed := make([]gopurs_runtime.Value, len(arr))
													for i, v := range arr {
														boxed[i] = v
													}
													return gopurs_runtime.Array(boxed)
												}()
												_ = arr_val_arrayMap9
												arr_go_arrayMap9 := (*[]gopurs_runtime.Value)(arr_val_arrayMap9.UnsafePtr)
												_ = arr_go_arrayMap9
												res_go_arrayMap9 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap9))
												_ = res_go_arrayMap9
												for i_arrayMap9, v_arrayMap9 := range *arr_go_arrayMap9 {
													res_go_arrayMap9[i_arrayMap9] = gopurs_runtime.Apply(gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
														return gopurs_runtime.RecordUpdate1(v1_5, "x", gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_1_1, "bimap"), gopurs_runtime.Func(func(v2_6 gopurs_runtime.Value) gopurs_runtime.Value {
															return gopurs_runtime.RecordUpdate1(v2_6, "a", gopurs_runtime.Apply(f_2, gopurs_runtime.RecordGet(v2_6, "a")))
														}), gopurs_runtime.Func(func(v2_6 gopurs_runtime.Value) gopurs_runtime.Value {
															return gopurs_runtime.RecordUpdate1(v2_6, "b", gopurs_runtime.Apply(g_3, gopurs_runtime.RecordGet(v2_6, "b")))
														}), gopurs_runtime.RecordGet(v1_5, "x")))
													}), v_arrayMap9)
												}
												return gopurs_runtime.Array(res_go_arrayMap9)
											}().UnsafePtr)
											unboxed := make([]gopurs_runtime.Value, len(arr))
											for i, v := range arr {
												unboxed[i] = v
											}
											return unboxed
										}()).UnsafePtr)
										unboxed := make([]gopurs_runtime.Value, len(arr))
										for i, v := range arr {
											unboxed[i] = v
										}
										return unboxed
									}()
									boxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										boxed[i] = v
									}
									return gopurs_runtime.Array(boxed)
								}()
								return gopurs_runtime.Value{Type: gopurs_runtime.TypeRecord1, UnsafePtr: unsafe.Pointer(&clone)}
							}()}))}
							goto end_branch_2
						} else {

						}
					}
					{
						__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
					}
				end_branch_2:
					return __t2
				})
			})
		})))
		_ = bifunctorTest1_1_0
		// TAST (Let): __local_var_2_4 shape=App(Other) expectedFromAst=gopurs_runtime.Value actual=gopurs_runtime.Value bindingType=Any
		__local_var_2_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBitraversable_0, "Bifoldable1"), gopurs_runtime.Value{})
		_ = __local_var_2_4
		// TAST (Let): bifoldableTest1_2_3 shape=Let(LitRecord) expectedFromAst=*Constructor_Data_Bifoldable_Bifoldable actual=*Constructor_Data_Bifoldable_Bifoldable bindingType=(ADT ["Data","Bifoldable","Bifoldable"] [(ADT ["Main","Test"] [(TypeVar f)])])
		bifoldableTest1_2_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Bifoldable_Bifoldable](gopurs_runtime.RecordDict3("bifoldMap", "bifoldl", "bifoldr", gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): mempty_4_5 shape=Other expectedFromAst=gopurs_runtime.Value actual=gopurs_runtime.Value bindingType=(TypeVar m)
			mempty_4_5 := gopurs_runtime.RecordGet(dictMonoid_3, "mempty")
			_ = mempty_4_5
			// TAST (Let): Semigroup0_5_6 shape=App(Other) expectedFromAst=*Constructor_Data_Semigroup_Semigroup actual=*Constructor_Data_Semigroup_Semigroup bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar m)])
			Semigroup0_5_6 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_3, "Semigroup0"), gopurs_runtime.Value{}))
			_ = Semigroup0_5_6
			// TAST (Let): bifoldMap2_6_7 shape=App(Other) expectedFromAst=gopurs_runtime.Value actual=gopurs_runtime.Value bindingType=(Func [(Func [(TypeVar a)] (TypeVar m)), (Func [Int] (TypeVar m)), (TypeApp (TypeVar f) [(TypeVar a), Int])] (TypeVar m))
			bifoldMap2_6_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_4, "bifoldMap"), dictMonoid_3)
			_ = bifoldMap2_6_7
			// TAST (Let): Semigroup0_7_9 shape=App(Other) expectedFromAst=*Constructor_Data_Semigroup_Semigroup actual=*Constructor_Data_Semigroup_Semigroup bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar m)])
			Semigroup0_7_9 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_3, "Semigroup0"), gopurs_runtime.Value{}))
			_ = Semigroup0_7_9
			// TAST (Let): bifoldMap3_7_8 shape=Let(Abs(Abs(Abs(App(Other))))) expectedFromAst=gopurs_runtime.Value actual=gopurs_runtime.Value bindingType=(Func [(Func [(TypeVar a)] (TypeVar m)), (Func [Int] (TypeVar m)), (ADT ["Data","Tuple","Tuple"] [(TypeVar a), Int])] (TypeVar m))
			bifoldMap3_7_8 := gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(g_9 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_7_9.V0), gopurs_runtime.Apply(f_8, (*Constructor_Data_Tuple_Tuple)(v_10.UnsafePtr).V0), gopurs_runtime.Apply(g_9, (*Constructor_Data_Tuple_Tuple)(v_10.UnsafePtr).V1))
					})
				})
			})
			_ = bifoldMap3_7_8
			// TAST (Let): Semigroup0_8_11 shape=App(Other) expectedFromAst=*Constructor_Data_Semigroup_Semigroup actual=*Constructor_Data_Semigroup_Semigroup bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar m)])
			Semigroup0_8_11 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_3, "Semigroup0"), gopurs_runtime.Value{}))
			_ = Semigroup0_8_11
			// TAST (Let): bifoldMap4_8_10 shape=Let(Abs(Abs(Abs(App(Other))))) expectedFromAst=gopurs_runtime.Value actual=gopurs_runtime.Value bindingType=(Func [(Func [(TypeVar b)] (TypeVar m)), (Func [Int] (TypeVar m)), (ADT ["Data","Tuple","Tuple"] [(TypeVar b), Int])] (TypeVar m))
			bifoldMap4_8_10 := gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(g_10 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_8_11.V0), gopurs_runtime.Apply(f_9, (*Constructor_Data_Tuple_Tuple)(v_11.UnsafePtr).V0), gopurs_runtime.Apply(g_10, (*Constructor_Data_Tuple_Tuple)(v_11.UnsafePtr).V1))
					})
				})
			})
			_ = bifoldMap4_8_10
			return gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(g_10 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(m_11 gopurs_runtime.Value) gopurs_runtime.Value {
						var __t16 gopurs_runtime.Value
						{
							if m_11.Type == 9 && m_11.IntVal == 2074462008 {
								__t16 = mempty_4_5
								goto end_branch_16
							} else {

							}
						}
						{
							if m_11.Type == 9 && m_11.IntVal == 3720114489 {
								// TAST (Let): Semigroup0_12_12 shape=App(Other) expectedFromAst=*Constructor_Data_Semigroup_Semigroup actual=*Constructor_Data_Semigroup_Semigroup bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar m)])
								Semigroup0_12_12 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_3, "Semigroup0"), gopurs_runtime.Value{}))
								_ = Semigroup0_12_12
								__t16 = gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_5_6.V0), gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(x_13 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Func(func(acc_14 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_12_12.V0), gopurs_runtime.Apply(f_9, x_13), acc_14)
									})
								}), gopurs_runtime.RecordGet(dictMonoid_3, "mempty"), gopurs_runtime.Array((*Constructor_Main_Test1)(m_11.UnsafePtr).V0)), gopurs_runtime.Apply(g_10, (*Constructor_Main_Test1)(m_11.UnsafePtr).V1))
								goto end_branch_16
							} else {

							}
						}
						{
							if m_11.Type == 9 && m_11.IntVal == 2375191994 {
								__t16 = mempty_4_5
								goto end_branch_16
							} else {

							}
						}
						{
							if m_11.Type == 9 && m_11.IntVal == 227416251 {
								__t16 = gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_5_6.V0), gopurs_runtime.Apply4(gopurs_runtime.RecordGet(__local_var_2_4, "bifoldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](dictMonoid_3))}, f_9, g_10, (*Constructor_Main_Test3)(m_11.UnsafePtr).V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_5_6.V0), gopurs_runtime.Apply3(bifoldMap2_6_7, f_9, gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.RecordGet(dictMonoid_3, "mempty")
								}), (*Constructor_Main_Test3)(m_11.UnsafePtr).V2), gopurs_runtime.Apply4(gopurs_runtime.RecordGet(__local_var_2_4, "bifoldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](dictMonoid_3))}, gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.RecordGet(dictMonoid_3, "mempty")
								}), g_10, (*Constructor_Main_Test3)(m_11.UnsafePtr).V3)))
								goto end_branch_16
							} else {

							}
						}
						{
							if m_11.Type == 9 && m_11.IntVal == 3712677948 {
								// TAST (Let): Semigroup0_12_13 shape=App(Other) expectedFromAst=*Constructor_Data_Semigroup_Semigroup actual=*Constructor_Data_Semigroup_Semigroup bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar m)])
								Semigroup0_12_13 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_3, "Semigroup0"), gopurs_runtime.Value{}))
								_ = Semigroup0_12_13
								// TAST (Let): __local_var_13_14 shape=App(Other) expectedFromAst=gopurs_runtime.Value actual=gopurs_runtime.Value bindingType=(Func [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), Int])] (TypeVar m))
								__local_var_13_14 := gopurs_runtime.Apply2(bifoldMap3_7_8, f_9, gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.RecordGet(dictMonoid_3, "mempty")
								}))
								_ = __local_var_13_14
								__t16 = gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_5_6.V0), gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(x_14 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Func(func(acc_15 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_12_13.V0), gopurs_runtime.Apply(__local_var_13_14, x_14), acc_15)
									})
								}), gopurs_runtime.RecordGet(dictMonoid_3, "mempty"), func() gopurs_runtime.Value {
									arr := (*Constructor_Main_Test4)(m_11.UnsafePtr).V0
									boxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)}
									}
									return gopurs_runtime.Array(boxed)
								}()), gopurs_runtime.Apply3(bifoldMap4_8_10, g_10, gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.RecordGet(dictMonoid_3, "mempty")
								}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer((*Constructor_Main_Test4)(m_11.UnsafePtr).V1)}))
								goto end_branch_16
							} else {

							}
						}
						{
							if m_11.Type == 9 && m_11.IntVal == 1063363133 {
								// TAST (Let): Semigroup0_12_15 shape=App(Other) expectedFromAst=*Constructor_Data_Semigroup_Semigroup actual=*Constructor_Data_Semigroup_Semigroup bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar m)])
								Semigroup0_12_15 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_3, "Semigroup0"), gopurs_runtime.Value{}))
								_ = Semigroup0_12_15
								__t16 = gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(x_13 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Func(func(acc_14 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_12_15.V0), gopurs_runtime.Apply4(gopurs_runtime.RecordGet(__local_var_2_4, "bifoldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](dictMonoid_3))}, gopurs_runtime.Func(func(v2_15 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Apply(f_9, gopurs_runtime.RecordGet(v2_15, "a"))
										}), gopurs_runtime.Func(func(v2_15 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Apply(g_10, gopurs_runtime.RecordGet(v2_15, "b"))
										}), gopurs_runtime.RecordGet(x_13, "x")), acc_14)
									})
								}), gopurs_runtime.RecordGet(dictMonoid_3, "mempty"), func() gopurs_runtime.Value {
									arr := func() []gopurs_runtime.Value {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_Test5)(m_11.UnsafePtr).V0, "nested").UnsafePtr)
										unboxed := make([]gopurs_runtime.Value, len(arr))
										for i, v := range arr {
											unboxed[i] = v
										}
										return unboxed
									}()
									boxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										boxed[i] = v
									}
									return gopurs_runtime.Array(boxed)
								}())
								goto end_branch_16
							} else {

							}
						}
						{
							__t16 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
						}
					end_branch_16:
						return __t16
					})
				})
			})
		}), gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(g_4 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(z_5 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(m_6 gopurs_runtime.Value) gopurs_runtime.Value {
						var __t17 gopurs_runtime.Value
						{
							if m_6.Type == 9 && m_6.IntVal == 2074462008 {
								__t17 = z_5
								goto end_branch_17
							} else {

							}
						}
						{
							if m_6.Type == 9 && m_6.IntVal == 3720114489 {
								__t17 = gopurs_runtime.Apply2(g_4, func() gopurs_runtime.Value {
									arr_val_foldlArray10 := gopurs_runtime.Array((*Constructor_Main_Test1)(m_6.UnsafePtr).V0)
									_ = arr_val_foldlArray10
									res_go_foldlArray10 := z_5
									_ = res_go_foldlArray10
									arr_go_foldlArray10 := (*[]gopurs_runtime.Value)(arr_val_foldlArray10.UnsafePtr)
									_ = arr_go_foldlArray10
									for _, v_foldlArray10 := range *arr_go_foldlArray10 {
										res_go_foldlArray10 = gopurs_runtime.Apply2(f_3, res_go_foldlArray10, v_foldlArray10)
									}
									return res_go_foldlArray10
								}(), (*Constructor_Main_Test1)(m_6.UnsafePtr).V1)
								goto end_branch_17
							} else {

							}
						}
						{
							if m_6.Type == 9 && m_6.IntVal == 2375191994 {
								__t17 = z_5
								goto end_branch_17
							} else {

							}
						}
						{
							if m_6.Type == 9 && m_6.IntVal == 227416251 {
								__t17 = gopurs_runtime.Apply4(gopurs_runtime.RecordGet(__local_var_2_4, "bifoldl"), Get_Data_Function_go__const(), g_4, gopurs_runtime.Apply4(gopurs_runtime.RecordGet(__local_var_2_4, "bifoldl"), f_3, Get_Data_Function_go__const(), gopurs_runtime.Apply4(gopurs_runtime.RecordGet(__local_var_2_4, "bifoldl"), f_3, g_4, z_5, (*Constructor_Main_Test3)(m_6.UnsafePtr).V1), (*Constructor_Main_Test3)(m_6.UnsafePtr).V2), (*Constructor_Main_Test3)(m_6.UnsafePtr).V3)
								goto end_branch_17
							} else {

							}
						}
						{
							if m_6.Type == 9 && m_6.IntVal == 3712677948 {
								__t17 = gopurs_runtime.Apply2(g_4, func() gopurs_runtime.Value {
									arr_val_foldlArray10 := func() gopurs_runtime.Value {
										arr := (*Constructor_Main_Test4)(m_6.UnsafePtr).V0
										boxed := make([]gopurs_runtime.Value, len(arr))
										for i, v := range arr {
											boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)}
										}
										return gopurs_runtime.Array(boxed)
									}()
									_ = arr_val_foldlArray10
									res_go_foldlArray10 := z_5
									_ = res_go_foldlArray10
									arr_go_foldlArray10 := (*[]gopurs_runtime.Value)(arr_val_foldlArray10.UnsafePtr)
									_ = arr_go_foldlArray10
									for _, v_foldlArray10 := range *arr_go_foldlArray10 {
										res_go_foldlArray10 = gopurs_runtime.Apply2(gopurs_runtime.Func(func(z_7 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
												return gopurs_runtime.Apply2(f_3, z_7, (*Constructor_Data_Tuple_Tuple)(v_8.UnsafePtr).V0)
											})
										}), res_go_foldlArray10, v_foldlArray10)
									}
									return res_go_foldlArray10
								}(), ((*Constructor_Main_Test4)(m_6.UnsafePtr).V1).V0)
								goto end_branch_17
							} else {

							}
						}
						{
							if m_6.Type == 9 && m_6.IntVal == 1063363133 {
								__t17 = func() gopurs_runtime.Value {
									arr_val_foldlArray9 := func() gopurs_runtime.Value {
										arr := func() []gopurs_runtime.Value {
											arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_Test5)(m_6.UnsafePtr).V0, "nested").UnsafePtr)
											unboxed := make([]gopurs_runtime.Value, len(arr))
											for i, v := range arr {
												unboxed[i] = v
											}
											return unboxed
										}()
										boxed := make([]gopurs_runtime.Value, len(arr))
										for i, v := range arr {
											boxed[i] = v
										}
										return gopurs_runtime.Array(boxed)
									}()
									_ = arr_val_foldlArray9
									res_go_foldlArray9 := z_5
									_ = res_go_foldlArray9
									arr_go_foldlArray9 := (*[]gopurs_runtime.Value)(arr_val_foldlArray9.UnsafePtr)
									_ = arr_go_foldlArray9
									for _, v_foldlArray9 := range *arr_go_foldlArray9 {
										res_go_foldlArray9 = gopurs_runtime.Apply2(gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Func(func(v2_8 gopurs_runtime.Value) gopurs_runtime.Value {
												return gopurs_runtime.Apply4(gopurs_runtime.RecordGet(__local_var_2_4, "bifoldl"), gopurs_runtime.Func(func(v3_9 gopurs_runtime.Value) gopurs_runtime.Value {
													return gopurs_runtime.Func(func(v4_10 gopurs_runtime.Value) gopurs_runtime.Value {
														return gopurs_runtime.Apply2(f_3, v3_9, gopurs_runtime.RecordGet(v4_10, "a"))
													})
												}), gopurs_runtime.Func(func(v3_9 gopurs_runtime.Value) gopurs_runtime.Value {
													return gopurs_runtime.Func(func(v4_10 gopurs_runtime.Value) gopurs_runtime.Value {
														return gopurs_runtime.Apply2(g_4, v3_9, gopurs_runtime.RecordGet(v4_10, "b"))
													})
												}), v1_7, gopurs_runtime.RecordGet(v2_8, "x"))
											})
										}), res_go_foldlArray9, v_foldlArray9)
									}
									return res_go_foldlArray9
								}()
								goto end_branch_17
							} else {

							}
						}
						{
							__t17 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
						}
					end_branch_17:
						return __t17
					})
				})
			})
		}), gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(g_4 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(z_5 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(m_6 gopurs_runtime.Value) gopurs_runtime.Value {
						var __t18 gopurs_runtime.Value
						{
							if m_6.Type == 9 && m_6.IntVal == 2074462008 {
								__t18 = z_5
								goto end_branch_18
							} else {

							}
						}
						{
							if m_6.Type == 9 && m_6.IntVal == 3720114489 {
								__t18 = gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), f_3, gopurs_runtime.Apply2(g_4, (*Constructor_Main_Test1)(m_6.UnsafePtr).V1, z_5), gopurs_runtime.Array((*Constructor_Main_Test1)(m_6.UnsafePtr).V0))
								goto end_branch_18
							} else {

							}
						}
						{
							if m_6.Type == 9 && m_6.IntVal == 2375191994 {
								__t18 = z_5
								goto end_branch_18
							} else {

							}
						}
						{
							if m_6.Type == 9 && m_6.IntVal == 227416251 {
								__t18 = gopurs_runtime.Apply4(gopurs_runtime.RecordGet(__local_var_2_4, "bifoldr"), f_3, g_4, gopurs_runtime.Apply4(gopurs_runtime.RecordGet(__local_var_2_4, "bifoldr"), f_3, gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
										return x_8
									})
								}), gopurs_runtime.Apply4(gopurs_runtime.RecordGet(__local_var_2_4, "bifoldr"), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
										return x_8
									})
								}), g_4, z_5, (*Constructor_Main_Test3)(m_6.UnsafePtr).V3), (*Constructor_Main_Test3)(m_6.UnsafePtr).V2), (*Constructor_Main_Test3)(m_6.UnsafePtr).V1)
								goto end_branch_18
							} else {

							}
						}
						{
							if m_6.Type == 9 && m_6.IntVal == 3712677948 {
								__t18 = gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(b_7 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Apply2(f_3, (*Constructor_Data_Tuple_Tuple)(b_7.UnsafePtr).V0, a_8)
									})
								}), gopurs_runtime.Apply2(g_4, ((*Constructor_Main_Test4)(m_6.UnsafePtr).V1).V0, z_5), func() gopurs_runtime.Value {
									arr := (*Constructor_Main_Test4)(m_6.UnsafePtr).V0
									boxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)}
									}
									return gopurs_runtime.Array(boxed)
								}())
								goto end_branch_18
							} else {

							}
						}
						{
							if m_6.Type == 9 && m_6.IntVal == 1063363133 {
								__t18 = gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Func(func(v2_8 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Apply4(gopurs_runtime.RecordGet(__local_var_2_4, "bifoldr"), gopurs_runtime.Func(func(v3_9 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Func(func(v4_10 gopurs_runtime.Value) gopurs_runtime.Value {
												return gopurs_runtime.Apply2(f_3, gopurs_runtime.RecordGet(v3_9, "a"), v4_10)
											})
										}), gopurs_runtime.Func(func(v3_9 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Func(func(v4_10 gopurs_runtime.Value) gopurs_runtime.Value {
												return gopurs_runtime.Apply2(g_4, gopurs_runtime.RecordGet(v3_9, "b"), v4_10)
											})
										}), v2_8, gopurs_runtime.RecordGet(v1_7, "x"))
									})
								}), z_5, func() gopurs_runtime.Value {
									arr := func() []gopurs_runtime.Value {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_Test5)(m_6.UnsafePtr).V0, "nested").UnsafePtr)
										unboxed := make([]gopurs_runtime.Value, len(arr))
										for i, v := range arr {
											unboxed[i] = v
										}
										return unboxed
									}()
									boxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										boxed[i] = v
									}
									return gopurs_runtime.Array(boxed)
								}())
								goto end_branch_18
							} else {

							}
						}
						{
							__t18 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
						}
					end_branch_18:
						return __t18
					})
				})
			})
		})))
		_ = bifoldableTest1_2_3
		return gopurs_runtime.Value{Type: 9, IntVal: 3704227322, UnsafePtr: unsafe.Pointer((&Constructor_Data_Bitraversable_Bitraversable{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4001671834, UnsafePtr: unsafe.Pointer(bifoldableTest1_2_3)}
		}), gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4141114362, UnsafePtr: unsafe.Pointer(bifunctorTest1_1_0)}
		}), gopurs_runtime.Func(func(dictApplicative_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply4(gopurs_runtime.RecordGet(Call_Main_bitraversableTest(dictBitraversable_0), "bitraverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_3))}, gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
					return x_5
				}), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
					return x_5
				}), v_4)
			})
		}), gopurs_runtime.Func(func(dictApplicative_3 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): Apply0_4_19 shape=App(Other) expectedFromAst=*Constructor_Control_Apply_Apply actual=*Constructor_Control_Apply_Apply bindingType=(ADT ["Control","Apply","Apply"] [(TypeVar f)])
			Apply0_4_19 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}))
			_ = Apply0_4_19
			// TAST (Let): Functor0_5_20 shape=App(Other) expectedFromAst=*Constructor_Data_Functor_Functor actual=*Constructor_Data_Functor_Functor bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar f)])
			Functor0_5_20 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
			_ = Functor0_5_20
			return gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(g_7 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(m_8 gopurs_runtime.Value) gopurs_runtime.Value {
						var __t28 gopurs_runtime.Value
						{
							if m_8.Type == 9 && m_8.IntVal == 2074462008 {
								__t28 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2074462008, UnsafePtr: unsafe.Pointer(nil)})
								goto end_branch_28
							} else {

							}
						}
						{
							if m_8.Type == 9 && m_8.IntVal == 3720114489 {
								// TAST (Let): Apply0_9_21 shape=App(Other) expectedFromAst=gopurs_runtime.Value actual=gopurs_runtime.Value bindingType=Any
								Apply0_9_21 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{})
								_ = Apply0_9_21
								__t28 = gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_4_19.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_5_20.V0), gopurs_runtime.Func(func(v2_9 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Func(func(v3_10 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Value{Type: 9, IntVal: 3720114489, UnsafePtr: unsafe.Pointer((&Constructor_Main_Test1{1, func() []gopurs_runtime.Value {
											arr := *(*[]gopurs_runtime.Value)(v2_9.UnsafePtr)
											unboxed := make([]gopurs_runtime.Value, len(arr))
											for i, v := range arr {
												unboxed[i] = v
											}
											return unboxed
										}(), v3_10}))}
									})
								}), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), gopurs_runtime.RecordGet(Apply0_9_21, "apply"), gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_9_21, "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.RecordGet(dictApplicative_3, "pure"), Get_Data_Semigroup_concatArray(), f_6, gopurs_runtime.Array((*Constructor_Main_Test1)(m_8.UnsafePtr).V0))), gopurs_runtime.Apply(g_7, (*Constructor_Main_Test1)(m_8.UnsafePtr).V1))
								goto end_branch_28
							} else {

							}
						}
						{
							if m_8.Type == 9 && m_8.IntVal == 2375191994 {
								__t28 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2375191994, UnsafePtr: unsafe.Pointer((&Constructor_Main_Test2{1, (*Constructor_Main_Test2)(m_8.UnsafePtr).V0, (*Constructor_Main_Test2)(m_8.UnsafePtr).V1}))})
								goto end_branch_28
							} else {

							}
						}
						{
							if m_8.Type == 9 && m_8.IntVal == 227416251 {
								// TAST (Let): __local_var_9_22 shape=Other expectedFromAst=gopurs_runtime.Value actual=int64 bindingType=Any
								__local_var_9_22 := (*Constructor_Main_Test3)(m_8.UnsafePtr).V0
								_ = __local_var_9_22
								__t28 = gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_4_19.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_4_19.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_5_20.V0), gopurs_runtime.Func(func(v4_10 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Func(func(v5_11 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Func(func(v6_12 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Value{Type: 9, IntVal: 227416251, UnsafePtr: unsafe.Pointer((&Constructor_Main_Test3{1, __local_var_9_22, v4_10, v5_11, v6_12}))}
										})
									})
								}), gopurs_runtime.Apply4(gopurs_runtime.RecordGet(dictBitraversable_0, "bitraverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_3))}, f_6, g_7, (*Constructor_Main_Test3)(m_8.UnsafePtr).V1)), gopurs_runtime.Apply4(gopurs_runtime.RecordGet(dictBitraversable_0, "bitraverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_3))}, f_6, gopurs_runtime.RecordGet(dictApplicative_3, "pure"), (*Constructor_Main_Test3)(m_8.UnsafePtr).V2)), gopurs_runtime.Apply4(gopurs_runtime.RecordGet(dictBitraversable_0, "bitraverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_3))}, gopurs_runtime.RecordGet(dictApplicative_3, "pure"), g_7, (*Constructor_Main_Test3)(m_8.UnsafePtr).V3))
								goto end_branch_28
							} else {

							}
						}
						{
							if m_8.Type == 9 && m_8.IntVal == 3712677948 {
								// TAST (Let): Apply0_9_23 shape=App(Other) expectedFromAst=gopurs_runtime.Value actual=gopurs_runtime.Value bindingType=Any
								Apply0_9_23 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{})
								_ = Apply0_9_23
								// TAST (Let): Apply0_10_24 shape=App(Other) expectedFromAst=*Constructor_Control_Apply_Apply actual=*Constructor_Control_Apply_Apply bindingType=(ADT ["Control","Apply","Apply"] [(TypeVar f)])
								Apply0_10_24 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}))
								_ = Apply0_10_24
								// TAST (Let): Functor0_11_25 shape=App(Other) expectedFromAst=*Constructor_Data_Functor_Functor actual=*Constructor_Data_Functor_Functor bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar f)])
								Functor0_11_25 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
								_ = Functor0_11_25
								__t28 = gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_4_19.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_5_20.V0), gopurs_runtime.Func(func(v2_9 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Func(func(v3_10 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Value{Type: 9, IntVal: 3712677948, UnsafePtr: unsafe.Pointer((&Constructor_Main_Test4{1, func() []*Constructor_Data_Tuple_Tuple {
											arr := *(*[]gopurs_runtime.Value)(v2_9.UnsafePtr)
											unboxed := make([]*Constructor_Data_Tuple_Tuple, len(arr))
											for i, v := range arr {
												unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](v)
											}
											return unboxed
										}(), gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](v3_10)}))}
									})
								}), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), gopurs_runtime.RecordGet(Apply0_9_23, "apply"), gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_9_23, "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.RecordGet(dictApplicative_3, "pure"), Get_Data_Semigroup_concatArray(), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_10_24.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_11_25.V0), Get_Data_Tuple_Tuple(), gopurs_runtime.Apply(f_6, (*Constructor_Data_Tuple_Tuple)(v_12.UnsafePtr).V0)), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "pure"), (*Constructor_Data_Tuple_Tuple)(v_12.UnsafePtr).V1))
								}), func() gopurs_runtime.Value {
									arr := (*Constructor_Main_Test4)(m_8.UnsafePtr).V0
									boxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)}
									}
									return gopurs_runtime.Array(boxed)
								}())), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}), "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), Get_Data_Tuple_Tuple(), gopurs_runtime.Apply(g_7, ((*Constructor_Main_Test4)(m_8.UnsafePtr).V1).V0)), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "pure"), ((*Constructor_Main_Test4)(m_8.UnsafePtr).V1).V1)))
								goto end_branch_28
							} else {

							}
						}
						{
							if m_8.Type == 9 && m_8.IntVal == 1063363133 {
								// TAST (Let): __local_var_9_26 shape=Other expectedFromAst=gopurs_runtime.Value actual=gopurs_runtime.Value bindingType=Any
								__local_var_9_26 := (*Constructor_Main_Test5)(m_8.UnsafePtr).V0
								_ = __local_var_9_26
								// TAST (Let): Apply0_10_27 shape=App(Other) expectedFromAst=gopurs_runtime.Value actual=gopurs_runtime.Value bindingType=Any
								Apply0_10_27 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{})
								_ = Apply0_10_27
								__t28 = gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_5_20.V0), gopurs_runtime.Func(func(v1_10 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Value{Type: 9, IntVal: 1063363133, UnsafePtr: unsafe.Pointer((&Constructor_Main_Test5{1, func() gopurs_runtime.Value {
										origVal := __local_var_9_26
										if origVal.Type != gopurs_runtime.TypeRecord1 {
											return gopurs_runtime.RecordUpdateDict(origVal, []string{"nested"}, []gopurs_runtime.Value{func() gopurs_runtime.Value {
												arr := func() []gopurs_runtime.Value {
													arr := *(*[]gopurs_runtime.Value)(v1_10.UnsafePtr)
													unboxed := make([]gopurs_runtime.Value, len(arr))
													for i, v := range arr {
														unboxed[i] = v
													}
													return unboxed
												}()
												boxed := make([]gopurs_runtime.Value, len(arr))
												for i, v := range arr {
													boxed[i] = v
												}
												return gopurs_runtime.Array(boxed)
											}()})
										}
										clone := *((*gopurs_runtime.RecordData1)(origVal.UnsafePtr))
										clone.V0 = func() gopurs_runtime.Value {
											arr := func() []gopurs_runtime.Value {
												arr := *(*[]gopurs_runtime.Value)(v1_10.UnsafePtr)
												unboxed := make([]gopurs_runtime.Value, len(arr))
												for i, v := range arr {
													unboxed[i] = v
												}
												return unboxed
											}()
											boxed := make([]gopurs_runtime.Value, len(arr))
											for i, v := range arr {
												boxed[i] = v
											}
											return gopurs_runtime.Array(boxed)
										}()
										return gopurs_runtime.Value{Type: gopurs_runtime.TypeRecord1, UnsafePtr: unsafe.Pointer(&clone)}
									}()}))}
								}), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), gopurs_runtime.RecordGet(Apply0_10_27, "apply"), gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_10_27, "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.RecordGet(dictApplicative_3, "pure"), Get_Data_Semigroup_concatArray(), gopurs_runtime.Func(func(v1_11 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_5_20.V0), gopurs_runtime.Func(func(v2_12 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.RecordUpdate1(v1_11, "x", v2_12)
									}), gopurs_runtime.Apply4(gopurs_runtime.RecordGet(dictBitraversable_0, "bitraverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_3))}, gopurs_runtime.Func(func(v2_12 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_5_20.V0), gopurs_runtime.Func(func(v3_13 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.RecordUpdate1(v2_12, "a", v3_13)
										}), gopurs_runtime.Apply(f_6, gopurs_runtime.RecordGet(v2_12, "a")))
									}), gopurs_runtime.Func(func(v2_12 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_5_20.V0), gopurs_runtime.Func(func(v3_13 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.RecordUpdate1(v2_12, "b", v3_13)
										}), gopurs_runtime.Apply(g_7, gopurs_runtime.RecordGet(v2_12, "b")))
									}), gopurs_runtime.RecordGet(v1_11, "x")))
								}), func() gopurs_runtime.Value {
									arr := func() []gopurs_runtime.Value {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(__local_var_9_26, "nested").UnsafePtr)
										unboxed := make([]gopurs_runtime.Value, len(arr))
										for i, v := range arr {
											unboxed[i] = v
										}
										return unboxed
									}()
									boxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										boxed[i] = v
									}
									return gopurs_runtime.Array(boxed)
								}()))
								goto end_branch_28
							} else {

							}
						}
						{
							__t28 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
						}
					end_branch_28:
						return __t28
					})
				})
			})
		})}))}
	}
}
