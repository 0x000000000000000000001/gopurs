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
		cache_Main_bifoldl = Call_Data_Bifoldable_bifoldl(Rebox_Main_2812820739_3566843086(Rebox_Main_3566843086_2812820739(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value]](Get_Data_Bifoldable_bifoldableTuple()))))
	})
	return cache_Main_bifoldl
}

var cache_Main_bifoldl1 gopurs_runtime.Value
var once_Main_bifoldl1 sync.Once

func Get_Main_bifoldl1() gopurs_runtime.Value {
	once_Main_bifoldl1.Do(func() {
		cache_Main_bifoldl1 = Call_Data_Bifoldable_bifoldl(Rebox_Main_2812820739_3566843086(Rebox_Main_3566843086_2812820739(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value]](Get_Data_Bifoldable_bifoldableTuple()))))
	})
	return cache_Main_bifoldl1
}

var cache_Main_bifoldr gopurs_runtime.Value
var once_Main_bifoldr sync.Once

func Get_Main_bifoldr() gopurs_runtime.Value {
	once_Main_bifoldr.Do(func() {
		cache_Main_bifoldr = Call_Data_Bifoldable_bifoldr(Rebox_Main_2812820739_3566843086(Rebox_Main_3566843086_2812820739(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value]](Get_Data_Bifoldable_bifoldableTuple()))))
	})
	return cache_Main_bifoldr
}

var cache_Main_bifoldr1 gopurs_runtime.Value
var once_Main_bifoldr1 sync.Once

func Get_Main_bifoldr1() gopurs_runtime.Value {
	once_Main_bifoldr1.Do(func() {
		cache_Main_bifoldr1 = Call_Data_Bifoldable_bifoldr(Rebox_Main_2812820739_3566843086(Rebox_Main_3566843086_2812820739(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value]](Get_Data_Bifoldable_bifoldableTuple()))))
	})
	return cache_Main_bifoldr1
}

var cache_Main_bifoldMap gopurs_runtime.Value
var once_Main_bifoldMap sync.Once

func Get_Main_bifoldMap() gopurs_runtime.Value {
	once_Main_bifoldMap.Do(func() {
		cache_Main_bifoldMap = Call_Data_Bifoldable_bifoldMap(Rebox_Main_2812820739_3566843086(Rebox_Main_3566843086_2812820739(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value]](Get_Data_Bifoldable_bifoldableTuple()))))
	})
	return cache_Main_bifoldMap
}

var cache_Main_identity gopurs_runtime.Value
var once_Main_identity sync.Once

func Get_Main_identity() gopurs_runtime.Value {
	once_Main_identity.Do(func() {
		cache_Main_identity = Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))})
	})
	return cache_Main_identity
}

var cache_Main_identity1 gopurs_runtime.Value
var once_Main_identity1 sync.Once

func Get_Main_identity1() gopurs_runtime.Value {
	once_Main_identity1.Do(func() {
		cache_Main_identity1 = Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))})
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
				return gopurs_runtime.Value{Type: 9, IntVal: 3720114489, UnsafePtr: unsafe.Pointer((&Constructor_Main_Test1[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, func() []gopurs_runtime.Value {
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
				return gopurs_runtime.Value{Type: 9, IntVal: 2375191994, UnsafePtr: unsafe.Pointer((&Constructor_Main_Test2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, value0.IntVal, value1}))}
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
						return gopurs_runtime.Value{Type: 9, IntVal: 227416251, UnsafePtr: unsafe.Pointer((&Constructor_Main_Test3[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, value0.IntVal, value1, value2, value3}))}
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
				return gopurs_runtime.Value{Type: 9, IntVal: 3712677948, UnsafePtr: unsafe.Pointer((&Constructor_Main_Test4[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, func() []*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, int64] {
					arr := *(*[]gopurs_runtime.Value)(value0.UnsafePtr)
					unboxed := make([]*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, int64], len(arr))
					for i, v := range arr {
						unboxed[i] = Rebox_Main_138441832_3415943795(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v))
					}
					return unboxed
				}(), Rebox_Main_138441832_3415943795(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](value1))}))}
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
			return gopurs_runtime.Value{Type: 9, IntVal: 1063363133, UnsafePtr: unsafe.Pointer((&Constructor_Main_Test5[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, func() struct {
				nested []struct {
					x gopurs_runtime.Value
				}
			} {
				orig := value0
				_ = orig
				clone := struct {
					nested []struct {
						x gopurs_runtime.Value
					}
				}{}
				clone.nested = func() []struct {
					x gopurs_runtime.Value
				} {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "nested").UnsafePtr)
					unboxed := make([]struct {
						x gopurs_runtime.Value
					}, len(arr))
					for i, v := range arr {
						unboxed[i] = func() struct {
							x gopurs_runtime.Value
						} {
							orig := v
							_ = orig
							clone := struct {
								x gopurs_runtime.Value
							}{}
							clone.x = gopurs_runtime.RecordGet(orig, "x")
							return clone
						}()
					}
					return unboxed
				}()
				return clone
			}()}))}
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
				return gopurs_runtime.Value{Type: 9, IntVal: 2092050667, UnsafePtr: unsafe.Pointer((&Constructor_Main_FromProAndContra[gopurs_runtime.Value, gopurs_runtime.Value]{1, value0, value1}))}
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
		cache_Main_bifunctorFromProAndContra = gopurs_runtime.Value{Type: 9, IntVal: 4141114362, UnsafePtr: unsafe.Pointer(Rebox_Main_1048926258_1688994542((&Constructor_Data_Bifunctor_Bifunctor[*Constructor_Main_FromProAndContra[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, g_1 gopurs_runtime.Value, m_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2092050667, UnsafePtr: unsafe.Pointer((&Constructor_Main_FromProAndContra[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), (*Constructor_Main_FromProAndContra[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Profunctor_Profunctor[gopurs_runtime.Value]](Get_Data_Profunctor_profunctorFn()).V0, f_0, Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))}))), Call_Control_Semigroupoid_composeFlipped(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()), gopurs_runtime.Apply(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Contravariant_Contravariant[gopurs_runtime.Value]](Get_Data_Predicate_contravariantPredicate()).V0, g_1), Call_Control_Semigroupoid_composeFlipped(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()), (*Constructor_Main_FromProAndContra[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V1, Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))})))}))}
		})})))}
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

type Constructor_Main_Test0[T_f any, T_a any, T_b any] struct {
	Rc uint32
}

type Constructor_Main_Test1[T_f any, T_a any, T_b any] struct {
	Rc uint32
	V0 []gopurs_runtime.Value
	V1 T_b
}

type Constructor_Main_Test2[T_f any, T_a any, T_b any] struct {
	Rc uint32
	V0 int64
	V1 gopurs_runtime.Value
}

type Constructor_Main_Test3[T_f any, T_a any, T_b any] struct {
	Rc uint32
	V0 int64
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
	V3 gopurs_runtime.Value
}

type Constructor_Main_Test4[T_f any, T_a any, T_b any] struct {
	Rc uint32
	V0 []*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, int64]
	V1 *Constructor_Data_Tuple_Tuple[T_b, int64]
}

type Constructor_Main_Test5[T_f any, T_a any, T_b any] struct {
	Rc uint32
	V0 struct {
		nested []struct {
			x gopurs_runtime.Value
		}
	}
}

type Constructor_Main_FromProAndContra[T_a any, T_b any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}

func Call_Main_bifunctorTest(dictBifunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictBifunctor_0 gopurs_runtime.Value = dictBifunctor_0_loop
	_ = dictBifunctor_0
	return gopurs_runtime.Value{Type: 9, IntVal: 4141114362, UnsafePtr: unsafe.Pointer((&Constructor_Data_Bifunctor_Bifunctor[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(f_1 gopurs_runtime.Value, g_2 gopurs_runtime.Value, m_3 gopurs_runtime.Value) gopurs_runtime.Value {
		var __t1 gopurs_runtime.Value
		{
			if m_3.Type == 9 && m_3.IntVal == 2074462008 {
				__t1 = gopurs_runtime.Value{Type: 9, IntVal: 2074462008, UnsafePtr: unsafe.Pointer(nil)}
				goto end_branch_1
			} else {

			}
		}
		{
			if m_3.Type == 9 && m_3.IntVal == 3720114489 {
				__t1 = gopurs_runtime.Value{Type: 9, IntVal: 3720114489, UnsafePtr: unsafe.Pointer((&Constructor_Main_Test1[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
						arr_val_arrayMap4 := gopurs_runtime.Array((*Constructor_Main_Test1[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V0)
						_ = arr_val_arrayMap4
						arr_go_arrayMap4 := (*[]gopurs_runtime.Value)(arr_val_arrayMap4.UnsafePtr)
						_ = arr_go_arrayMap4
						res_go_arrayMap4 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap4))
						_ = res_go_arrayMap4
						for i_arrayMap4, v_arrayMap4 := range *arr_go_arrayMap4 {
							res_go_arrayMap4[i_arrayMap4] = gopurs_runtime.Apply(f_1, v_arrayMap4)
						}
						return gopurs_runtime.Array(res_go_arrayMap4)
					}().UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						unboxed[i] = v
					}
					return unboxed
				}(), gopurs_runtime.Apply(g_2, (*Constructor_Main_Test1[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V1)}))}
				goto end_branch_1
			} else {

			}
		}
		{
			if m_3.Type == 9 && m_3.IntVal == 2375191994 {
				__t1 = gopurs_runtime.Value{Type: 9, IntVal: 2375191994, UnsafePtr: unsafe.Pointer((&Constructor_Main_Test2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Main_Test2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V0, (*Constructor_Main_Test2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V1}))}
				goto end_branch_1
			} else {

			}
		}
		{
			if m_3.Type == 9 && m_3.IntVal == 227416251 {
				__t1 = gopurs_runtime.Value{Type: 9, IntVal: 227416251, UnsafePtr: unsafe.Pointer((&Constructor_Main_Test3[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Main_Test3[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V0, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictBifunctor_0, "bimap"), f_1, g_2, (*Constructor_Main_Test3[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V1), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictBifunctor_0, "bimap"), f_1, Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))}), (*Constructor_Main_Test3[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V2), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictBifunctor_0, "bimap"), Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))}), g_2, (*Constructor_Main_Test3[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V3)}))}
				goto end_branch_1
			} else {

			}
		}
		{
			if m_3.Type == 9 && m_3.IntVal == 3712677948 {
				// TAST (Let): __local_var_4_0 shape=App(Var) bindingType=(Func [(TypeVar c$scope18)] (TypeVar c$scope18))
				__local_var_4_0 := Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))})
				_ = __local_var_4_0
				__t1 = gopurs_runtime.Value{Type: 9, IntVal: 3712677948, UnsafePtr: unsafe.Pointer((&Constructor_Main_Test4[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, func() []*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, int64] {
					arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
						arr_val_arrayMap4 := func() gopurs_runtime.Value {
							arr := (*Constructor_Main_Test4[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V0
							boxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Main_3415943795_138441832(v))}
							}
							return gopurs_runtime.Array(boxed)
						}()
						_ = arr_val_arrayMap4
						arr_go_arrayMap4 := (*[]gopurs_runtime.Value)(arr_val_arrayMap4.UnsafePtr)
						_ = arr_go_arrayMap4
						res_go_arrayMap4 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap4))
						_ = res_go_arrayMap4
						for i_arrayMap4, v_arrayMap4 := range *arr_go_arrayMap4 {
							res_go_arrayMap4[i_arrayMap4] = gopurs_runtime.Apply(gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
									_v := struct {
										V0 gopurs_runtime.Value
										V1 gopurs_runtime.Value
									}{gopurs_runtime.Apply(f_1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0), gopurs_runtime.Apply(__local_var_4_0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V1)}
									return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
								}()))}
							}), v_arrayMap4)
						}
						return gopurs_runtime.Array(res_go_arrayMap4)
					}().UnsafePtr)
					unboxed := make([]*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, int64], len(arr))
					for i, v := range arr {
						unboxed[i] = Rebox_Main_138441832_3415943795(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v))
					}
					return unboxed
				}(), Rebox_Main_138441832_3415943795(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
					_v := struct {
						V0 gopurs_runtime.Value
						V1 gopurs_runtime.Value
					}{gopurs_runtime.Apply(g_2, ((*Constructor_Main_Test4[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V1).V0), gopurs_runtime.Apply(Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))}), gopurs_runtime.Int(((*Constructor_Main_Test4[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V1).V1))}
					return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
				}()))}))}
				goto end_branch_1
			} else {

			}
		}
		{
			if m_3.Type == 9 && m_3.IntVal == 1063363133 {
				__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1063363133, UnsafePtr: unsafe.Pointer((&Constructor_Main_Test5[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, func() struct {
					nested []struct {
						x gopurs_runtime.Value
					}
				} {
					clone := (*Constructor_Main_Test5[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V0
					clone.nested = func() []struct {
						x gopurs_runtime.Value
					} {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
								arr_val_arrayMap5 := func() gopurs_runtime.Value {
									arr := (*Constructor_Main_Test5[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V0.nested
									boxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										boxed[i] = func() gopurs_runtime.Value {
											orig := v
											_ = orig
											return gopurs_runtime.RecordDict1("x", orig.x)
										}()
									}
									return gopurs_runtime.Array(boxed)
								}()
								_ = arr_val_arrayMap5
								arr_go_arrayMap5 := (*[]gopurs_runtime.Value)(arr_val_arrayMap5.UnsafePtr)
								_ = arr_go_arrayMap5
								res_go_arrayMap5 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap5))
								_ = res_go_arrayMap5
								for i_arrayMap5, v_arrayMap5 := range *arr_go_arrayMap5 {
									res_go_arrayMap5[i_arrayMap5] = gopurs_runtime.Apply(gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
										return func() gopurs_runtime.Value {
											orig := func() struct {
												x gopurs_runtime.Value
											} {
												orig := gopurs_runtime.RecordUpdate1(v1_4, "x", gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictBifunctor_0, "bimap"), gopurs_runtime.Func(func(v2_5 gopurs_runtime.Value) gopurs_runtime.Value {
													return func() gopurs_runtime.Value {
														orig := func() struct {
															a gopurs_runtime.Value
														} {
															orig := gopurs_runtime.RecordUpdate1(v2_5, "a", gopurs_runtime.Apply(f_1, gopurs_runtime.RecordGet(v2_5, "a")))
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
												}), gopurs_runtime.Func(func(v2_5 gopurs_runtime.Value) gopurs_runtime.Value {
													return func() gopurs_runtime.Value {
														orig := func() struct {
															b gopurs_runtime.Value
														} {
															orig := gopurs_runtime.RecordUpdate1(v2_5, "b", gopurs_runtime.Apply(g_2, gopurs_runtime.RecordGet(v2_5, "b")))
															_ = orig
															clone := struct {
																b gopurs_runtime.Value
															}{}
															clone.b = gopurs_runtime.RecordGet(orig, "b")
															return clone
														}()
														_ = orig
														return gopurs_runtime.RecordDict1("b", orig.b)
													}()
												}), gopurs_runtime.RecordGet(v1_4, "x")))
												_ = orig
												clone := struct {
													x gopurs_runtime.Value
												}{}
												clone.x = gopurs_runtime.RecordGet(orig, "x")
												return clone
											}()
											_ = orig
											return gopurs_runtime.RecordDict1("x", orig.x)
										}()
									}), v_arrayMap5)
								}
								return gopurs_runtime.Array(res_go_arrayMap5)
							}().UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}()).UnsafePtr)
						unboxed := make([]struct {
							x gopurs_runtime.Value
						}, len(arr))
						for i, v := range arr {
							unboxed[i] = func() struct {
								x gopurs_runtime.Value
							} {
								orig := v
								_ = orig
								clone := struct {
									x gopurs_runtime.Value
								}{}
								clone.x = gopurs_runtime.RecordGet(orig, "x")
								return clone
							}()
						}
						return unboxed
					}()
					return clone
				}()}))}
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

func Call_Main_bifoldableTest(dictBifoldable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictBifoldable_0 gopurs_runtime.Value = dictBifoldable_0_loop
	_ = dictBifoldable_0
	// TAST (Let): bifoldl2_1_0 shape=App(Var) bindingType=(Func [(Func [(TypeVar c$scope99), (TypeVar a$scope97)] (TypeVar c$scope99)), (Func [(TypeVar c$scope99), Int] (TypeVar c$scope99)), (TypeVar c$scope99), (TypeApp (TypeVar f$scope86) [(TypeVar a$scope97), Int])] (TypeVar c$scope99))
	bifoldl2_1_0 := Call_Data_Bifoldable_bifoldl(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value]](dictBifoldable_0))
	_ = bifoldl2_1_0
	// TAST (Let): bifoldr2_2_1 shape=App(Var) bindingType=(Func [(Func [(TypeVar a$scope101), (TypeVar c$scope103)] (TypeVar c$scope103)), (Func [Int, (TypeVar c$scope103)] (TypeVar c$scope103)), (TypeVar c$scope103), (TypeApp (TypeVar f$scope86) [(TypeVar a$scope101), Int])] (TypeVar c$scope103))
	bifoldr2_2_1 := Call_Data_Bifoldable_bifoldr(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value]](dictBifoldable_0))
	_ = bifoldr2_2_1
	// TAST (Let): bifoldMap1_3_2 shape=App(Var) bindingType=Any
	bifoldMap1_3_2 := Call_Data_Bifoldable_bifoldMap(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value]](dictBifoldable_0))
	_ = bifoldMap1_3_2
	return gopurs_runtime.Value{Type: 9, IntVal: 4001671834, UnsafePtr: unsafe.Pointer((&Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(dictMonoid_4 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): mempty_5_3 shape=App(Var) bindingType=(TypeVar m$scope109)
		mempty_5_3 := Call_Data_Monoid_mempty(dictMonoid_4)
		_ = mempty_5_3
		// TAST (Let): Semigroup0_6_4 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar m$scope109)])
		Semigroup0_6_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_4, "Semigroup0"), gopurs_runtime.Value{}))
		_ = Semigroup0_6_4
		// TAST (Let): bifoldMap2_7_5 shape=App(Other) bindingType=(Func [(Func [(TypeVar a$scope110)] (TypeVar m$scope109)), (Func [Int] (TypeVar m$scope109)), (TypeApp (TypeVar f$scope86) [(TypeVar a$scope110), Int])] (TypeVar m$scope109))
		bifoldMap2_7_5 := gopurs_runtime.Apply(bifoldMap1_3_2, dictMonoid_4)
		_ = bifoldMap2_7_5
		// TAST (Let): mempty1_8_6 shape=App(Var) bindingType=(Func [Int] (TypeVar m$scope109))
		mempty1_8_6 := Call_Data_Monoid_mempty(Call_Data_Monoid_monoidFn(dictMonoid_4))
		_ = mempty1_8_6
		// TAST (Let): bifoldMap3_9_7 shape=App(Var) bindingType=(Func [(Func [(TypeVar a$scope110)] (TypeVar m$scope109)), (Func [Int] (TypeVar m$scope109)), (ADT ["Data","Tuple","Tuple"] [(TypeVar a$scope110), Int])] (TypeVar m$scope109))
		bifoldMap3_9_7 := gopurs_runtime.Apply(Call_Data_Bifoldable_bifoldMap(Rebox_Main_2812820739_3566843086(Rebox_Main_3566843086_2812820739(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value]](Get_Data_Bifoldable_bifoldableTuple())))), dictMonoid_4)
		_ = bifoldMap3_9_7
		// TAST (Let): bifoldMap4_10_8 shape=App(Var) bindingType=(Func [(Func [(TypeVar b$scope111)] (TypeVar m$scope109)), (Func [Int] (TypeVar m$scope109)), (ADT ["Data","Tuple","Tuple"] [(TypeVar b$scope111), Int])] (TypeVar m$scope109))
		bifoldMap4_10_8 := gopurs_runtime.Apply(Call_Data_Bifoldable_bifoldMap(Rebox_Main_2812820739_3566843086(Rebox_Main_3566843086_2812820739(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value]](Get_Data_Bifoldable_bifoldableTuple())))), dictMonoid_4)
		_ = bifoldMap4_10_8
		return gopurs_runtime.Func3(func(f_11 gopurs_runtime.Value, g_12 gopurs_runtime.Value, m_13 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t9 gopurs_runtime.Value
			{
				if m_13.Type == 9 && m_13.IntVal == 2074462008 {
					__t9 = mempty_5_3
					goto end_branch_9
				} else {

				}
			}
			{
				if m_13.Type == 9 && m_13.IntVal == 3720114489 {
					__t9 = gopurs_runtime.Apply2(Semigroup0_6_4.V0, gopurs_runtime.Apply2(Call_Data_Foldable_foldMapDefaultR(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()), gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_4)), f_11, gopurs_runtime.Array((*Constructor_Main_Test1[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(m_13.UnsafePtr).V0)), gopurs_runtime.Apply(g_12, (*Constructor_Main_Test1[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(m_13.UnsafePtr).V1))
					goto end_branch_9
				} else {

				}
			}
			{
				if m_13.Type == 9 && m_13.IntVal == 2375191994 {
					__t9 = mempty_5_3
					goto end_branch_9
				} else {

				}
			}
			{
				if m_13.Type == 9 && m_13.IntVal == 227416251 {
					__t9 = gopurs_runtime.Apply2(Semigroup0_6_4.V0, gopurs_runtime.Apply4(gopurs_runtime.RecordGet(dictBifoldable_0, "bifoldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_4))}, f_11, g_12, (*Constructor_Main_Test3[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(m_13.UnsafePtr).V1), gopurs_runtime.Apply2(Semigroup0_6_4.V0, gopurs_runtime.Apply3(bifoldMap2_7_5, f_11, mempty1_8_6, (*Constructor_Main_Test3[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(m_13.UnsafePtr).V2), gopurs_runtime.Apply4(gopurs_runtime.RecordGet(dictBifoldable_0, "bifoldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_4))}, mempty1_8_6, g_12, (*Constructor_Main_Test3[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(m_13.UnsafePtr).V3)))
					goto end_branch_9
				} else {

				}
			}
			{
				if m_13.Type == 9 && m_13.IntVal == 3712677948 {
					__t9 = gopurs_runtime.Apply2(Semigroup0_6_4.V0, gopurs_runtime.Apply2(Call_Data_Foldable_foldMapDefaultR(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()), gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_4)), gopurs_runtime.Apply2(bifoldMap3_9_7, f_11, mempty1_8_6), func() gopurs_runtime.Value {
						arr := (*Constructor_Main_Test4[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(m_13.UnsafePtr).V0
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Main_3415943795_138441832(v))}
						}
						return gopurs_runtime.Array(boxed)
					}()), gopurs_runtime.Apply3(bifoldMap4_10_8, g_12, mempty1_8_6, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Main_3415943795_138441832((*Constructor_Main_Test4[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(m_13.UnsafePtr).V1))}))
					goto end_branch_9
				} else {

				}
			}
			{
				if m_13.Type == 9 && m_13.IntVal == 1063363133 {
					__t9 = gopurs_runtime.Apply2(Call_Data_Foldable_foldMapDefaultR(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()), gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_4)), gopurs_runtime.Func(func(v1_14 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Apply4(gopurs_runtime.RecordGet(dictBifoldable_0, "bifoldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_4))}, gopurs_runtime.Func(func(v2_15 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Apply(f_11, gopurs_runtime.RecordGet(v2_15, "a"))
						}), gopurs_runtime.Func(func(v2_15 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Apply(g_12, gopurs_runtime.RecordGet(v2_15, "b"))
						}), gopurs_runtime.RecordGet(v1_14, "x"))
					}), func() gopurs_runtime.Value {
						arr := (*Constructor_Main_Test5[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(m_13.UnsafePtr).V0.nested
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = func() gopurs_runtime.Value {
								orig := v
								_ = orig
								return gopurs_runtime.RecordDict1("x", orig.x)
							}()
						}
						return gopurs_runtime.Array(boxed)
					}())
					goto end_branch_9
				} else {

				}
			}
			{
				__t9 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_9:
			return __t9
		})
	}), gopurs_runtime.Func4(func(f_4 gopurs_runtime.Value, g_5 gopurs_runtime.Value, z_6 gopurs_runtime.Value, m_7 gopurs_runtime.Value) gopurs_runtime.Value {
		var __t10 gopurs_runtime.Value
		{
			if m_7.Type == 9 && m_7.IntVal == 2074462008 {
				__t10 = z_6
				goto end_branch_10
			} else {

			}
		}
		{
			if m_7.Type == 9 && m_7.IntVal == 3720114489 {
				__t10 = gopurs_runtime.Apply2(g_5, func() gopurs_runtime.Value {
					arr_val_foldlArray7 := gopurs_runtime.Array((*Constructor_Main_Test1[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(m_7.UnsafePtr).V0)
					_ = arr_val_foldlArray7
					res_go_foldlArray7 := z_6
					_ = res_go_foldlArray7
					arr_go_foldlArray7 := (*[]gopurs_runtime.Value)(arr_val_foldlArray7.UnsafePtr)
					_ = arr_go_foldlArray7
					for _, v_foldlArray7 := range *arr_go_foldlArray7 {
						res_go_foldlArray7 = gopurs_runtime.Apply2(f_4, res_go_foldlArray7, v_foldlArray7)
					}
					return res_go_foldlArray7
				}(), (*Constructor_Main_Test1[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(m_7.UnsafePtr).V1)
				goto end_branch_10
			} else {

			}
		}
		{
			if m_7.Type == 9 && m_7.IntVal == 2375191994 {
				__t10 = z_6
				goto end_branch_10
			} else {

			}
		}
		{
			if m_7.Type == 9 && m_7.IntVal == 227416251 {
				__t10 = gopurs_runtime.Apply4(gopurs_runtime.RecordGet(dictBifoldable_0, "bifoldl"), Get_Data_Function_go__const(), g_5, gopurs_runtime.Apply4(bifoldl2_1_0, f_4, Get_Data_Function_go__const(), gopurs_runtime.Apply4(gopurs_runtime.RecordGet(dictBifoldable_0, "bifoldl"), f_4, g_5, z_6, (*Constructor_Main_Test3[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(m_7.UnsafePtr).V1), (*Constructor_Main_Test3[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(m_7.UnsafePtr).V2), (*Constructor_Main_Test3[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(m_7.UnsafePtr).V3)
				goto end_branch_10
			} else {

			}
		}
		{
			if m_7.Type == 9 && m_7.IntVal == 3712677948 {
				__t10 = gopurs_runtime.Apply4(Call_Data_Bifoldable_bifoldl(Rebox_Main_2812820739_3566843086(Rebox_Main_3566843086_2812820739(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value]](Get_Data_Bifoldable_bifoldableTuple())))), g_5, Get_Data_Function_go__const(), func() gopurs_runtime.Value {
					arr_val_foldlArray7 := func() gopurs_runtime.Value {
						arr := (*Constructor_Main_Test4[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(m_7.UnsafePtr).V0
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Main_3415943795_138441832(v))}
						}
						return gopurs_runtime.Array(boxed)
					}()
					_ = arr_val_foldlArray7
					res_go_foldlArray7 := z_6
					_ = res_go_foldlArray7
					arr_go_foldlArray7 := (*[]gopurs_runtime.Value)(arr_val_foldlArray7.UnsafePtr)
					_ = arr_go_foldlArray7
					for _, v_foldlArray7 := range *arr_go_foldlArray7 {
						res_go_foldlArray7 = gopurs_runtime.Apply2(gopurs_runtime.Apply2(Call_Data_Bifoldable_bifoldl(Rebox_Main_2812820739_3566843086(Rebox_Main_3566843086_2812820739(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value]](Get_Data_Bifoldable_bifoldableTuple())))), f_4, Get_Data_Function_go__const()), res_go_foldlArray7, v_foldlArray7)
					}
					return res_go_foldlArray7
				}(), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Main_3415943795_138441832((*Constructor_Main_Test4[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(m_7.UnsafePtr).V1))})
				goto end_branch_10
			} else {

			}
		}
		{
			if m_7.Type == 9 && m_7.IntVal == 1063363133 {
				__t10 = func() gopurs_runtime.Value {
					arr_val_foldlArray6 := func() gopurs_runtime.Value {
						arr := (*Constructor_Main_Test5[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(m_7.UnsafePtr).V0.nested
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = func() gopurs_runtime.Value {
								orig := v
								_ = orig
								return gopurs_runtime.RecordDict1("x", orig.x)
							}()
						}
						return gopurs_runtime.Array(boxed)
					}()
					_ = arr_val_foldlArray6
					res_go_foldlArray6 := z_6
					_ = res_go_foldlArray6
					arr_go_foldlArray6 := (*[]gopurs_runtime.Value)(arr_val_foldlArray6.UnsafePtr)
					_ = arr_go_foldlArray6
					for _, v_foldlArray6 := range *arr_go_foldlArray6 {
						res_go_foldlArray6 = gopurs_runtime.Apply2(gopurs_runtime.Func2(func(v1_8 gopurs_runtime.Value, v2_9 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Apply4(gopurs_runtime.RecordGet(dictBifoldable_0, "bifoldl"), gopurs_runtime.Func2(func(v3_10 gopurs_runtime.Value, v4_11 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Apply2(f_4, v3_10, gopurs_runtime.RecordGet(v4_11, "a"))
							}), gopurs_runtime.Func2(func(v3_10 gopurs_runtime.Value, v4_11 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Apply2(g_5, v3_10, gopurs_runtime.RecordGet(v4_11, "b"))
							}), v1_8, gopurs_runtime.RecordGet(v2_9, "x"))
						}), res_go_foldlArray6, v_foldlArray6)
					}
					return res_go_foldlArray6
				}()
				goto end_branch_10
			} else {

			}
		}
		{
			__t10 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
		}
	end_branch_10:
		return __t10
	}), gopurs_runtime.Func4(func(f_4 gopurs_runtime.Value, g_5 gopurs_runtime.Value, z_6 gopurs_runtime.Value, m_7 gopurs_runtime.Value) gopurs_runtime.Value {
		var __t16 gopurs_runtime.Value
		{
			if m_7.Type == 9 && m_7.IntVal == 2074462008 {
				__t16 = z_6
				goto end_branch_16
			} else {

			}
		}
		{
			if m_7.Type == 9 && m_7.IntVal == 3720114489 {
				__t16 = gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), f_4, gopurs_runtime.Apply2(g_5, (*Constructor_Main_Test1[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(m_7.UnsafePtr).V1, z_6), gopurs_runtime.Array((*Constructor_Main_Test1[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(m_7.UnsafePtr).V0))
				goto end_branch_16
			} else {

			}
		}
		{
			if m_7.Type == 9 && m_7.IntVal == 2375191994 {
				__t16 = z_6
				goto end_branch_16
			} else {

			}
		}
		{
			if m_7.Type == 9 && m_7.IntVal == 227416251 {
				// TAST (Let): __local_var_8_11 shape=App(Var) bindingType=Any
				__local_var_8_11 := Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))})
				_ = __local_var_8_11
				// TAST (Let): __local_var_8_12 shape=App(Var) bindingType=Any
				__local_var_8_12 := Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))})
				_ = __local_var_8_12
				__t16 = gopurs_runtime.Apply4(gopurs_runtime.RecordGet(dictBifoldable_0, "bifoldr"), f_4, g_5, gopurs_runtime.Apply4(bifoldr2_2_1, f_4, gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
					return __local_var_8_11
				}), gopurs_runtime.Apply4(gopurs_runtime.RecordGet(dictBifoldable_0, "bifoldr"), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
					return __local_var_8_12
				}), g_5, z_6, (*Constructor_Main_Test3[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(m_7.UnsafePtr).V3), (*Constructor_Main_Test3[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(m_7.UnsafePtr).V2), (*Constructor_Main_Test3[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(m_7.UnsafePtr).V1)
				goto end_branch_16
			} else {

			}
		}
		{
			if m_7.Type == 9 && m_7.IntVal == 3712677948 {
				// TAST (Let): __local_var_8_14 shape=App(Var) bindingType=Any
				__local_var_8_14 := Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))})
				_ = __local_var_8_14
				// TAST (Let): __local_var_8_13 shape=App(Var) bindingType=(Func [(TypeVar c$scope103), (ADT ["Data","Tuple","Tuple"] [(TypeVar a$scope101), Int])] (TypeVar c$scope103))
				__local_var_8_13 := gopurs_runtime.Apply2(Call_Data_Bifoldable_bifoldr(Rebox_Main_2812820739_3566843086(Rebox_Main_3566843086_2812820739(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value]](Get_Data_Bifoldable_bifoldableTuple())))), f_4, gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
					return __local_var_8_14
				}))
				_ = __local_var_8_13
				// TAST (Let): __local_var_8_15 shape=App(Var) bindingType=Any
				__local_var_8_15 := Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))})
				_ = __local_var_8_15
				__t16 = gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func2(func(b_9 gopurs_runtime.Value, a_10 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply2(__local_var_8_13, a_10, b_9)
				}), gopurs_runtime.Apply4(Call_Data_Bifoldable_bifoldr(Rebox_Main_2812820739_3566843086(Rebox_Main_3566843086_2812820739(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value]](Get_Data_Bifoldable_bifoldableTuple())))), g_5, gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
					return __local_var_8_15
				}), z_6, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Main_3415943795_138441832((*Constructor_Main_Test4[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(m_7.UnsafePtr).V1))}), func() gopurs_runtime.Value {
					arr := (*Constructor_Main_Test4[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(m_7.UnsafePtr).V0
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Main_3415943795_138441832(v))}
					}
					return gopurs_runtime.Array(boxed)
				}())
				goto end_branch_16
			} else {

			}
		}
		{
			if m_7.Type == 9 && m_7.IntVal == 1063363133 {
				__t16 = gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func2(func(v1_8 gopurs_runtime.Value, v2_9 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply4(gopurs_runtime.RecordGet(dictBifoldable_0, "bifoldr"), gopurs_runtime.Func2(func(v3_10 gopurs_runtime.Value, v4_11 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Apply2(f_4, gopurs_runtime.RecordGet(v3_10, "a"), v4_11)
					}), gopurs_runtime.Func2(func(v3_10 gopurs_runtime.Value, v4_11 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Apply2(g_5, gopurs_runtime.RecordGet(v3_10, "b"), v4_11)
					}), v2_9, gopurs_runtime.RecordGet(v1_8, "x"))
				}), z_6, func() gopurs_runtime.Value {
					arr := (*Constructor_Main_Test5[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(m_7.UnsafePtr).V0.nested
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = func() gopurs_runtime.Value {
							orig := v
							_ = orig
							return gopurs_runtime.RecordDict1("x", orig.x)
						}()
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
		// TAST (Let): bifunctorTest1_1_0 shape=App(Var) bindingType=(ADT ["Data","Bifunctor","Bifunctor"] [(ADT ["Main","Test"] [(TypeVar f$scope1)])])
		bifunctorTest1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Bifunctor_Bifunctor[gopurs_runtime.Value]](Call_Main_bifunctorTest(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBitraversable_0, "Bifunctor0"), gopurs_runtime.Value{})))
		_ = bifunctorTest1_1_0
		// TAST (Let): bifoldableTest1_2_1 shape=App(Var) bindingType=(ADT ["Data","Bifoldable","Bifoldable"] [(ADT ["Main","Test"] [(TypeVar f$scope1)])])
		bifoldableTest1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value]](Call_Main_bifoldableTest(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBitraversable_0, "Bifoldable1"), gopurs_runtime.Value{})))
		_ = bifoldableTest1_2_1
		return gopurs_runtime.Value{Type: 9, IntVal: 3704227322, UnsafePtr: unsafe.Pointer((&Constructor_Data_Bitraversable_Bitraversable[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4001671834, UnsafePtr: unsafe.Pointer(bifoldableTest1_2_1)}
		}), gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4141114362, UnsafePtr: unsafe.Pointer(bifunctorTest1_1_0)}
		}), gopurs_runtime.Func2(func(dictApplicative_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply4(gopurs_runtime.RecordGet(Call_Main_bitraversableTest(dictBitraversable_0), "bitraverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_3))}, Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))}), Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))}), v_4)
		}), gopurs_runtime.Func(func(dictApplicative_3 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): Apply0_4_2 shape=App(Other) bindingType=(ADT ["Control","Apply","Apply"] [(TypeVar f$scope11)])
			Apply0_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}))
			_ = Apply0_4_2
			// TAST (Let): Functor0_5_3 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar f$scope11)])
			Functor0_5_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
			_ = Functor0_5_3
			return gopurs_runtime.Func3(func(f_6 gopurs_runtime.Value, g_7 gopurs_runtime.Value, m_8 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t6 gopurs_runtime.Value
				{
					if m_8.Type == 9 && m_8.IntVal == 2074462008 {
						__t6 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2074462008, UnsafePtr: unsafe.Pointer(nil)})
						goto end_branch_6
					} else {

					}
				}
				{
					if m_8.Type == 9 && m_8.IntVal == 3720114489 {
						__t6 = gopurs_runtime.Apply2(Apply0_4_2.V1, gopurs_runtime.Apply2(Functor0_5_3.V0, gopurs_runtime.Func2(func(v2_9 gopurs_runtime.Value, v3_10 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Value{Type: 9, IntVal: 3720114489, UnsafePtr: unsafe.Pointer((&Constructor_Main_Test1[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(v2_9.UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}(), v3_10}))}
						}), gopurs_runtime.Apply3(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableArray()).V3, gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_3))}, f_6, gopurs_runtime.Array((*Constructor_Main_Test1[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(m_8.UnsafePtr).V0))), gopurs_runtime.Apply(g_7, (*Constructor_Main_Test1[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(m_8.UnsafePtr).V1))
						goto end_branch_6
					} else {

					}
				}
				{
					if m_8.Type == 9 && m_8.IntVal == 2375191994 {
						__t6 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2375191994, UnsafePtr: unsafe.Pointer((&Constructor_Main_Test2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Main_Test2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(m_8.UnsafePtr).V0, (*Constructor_Main_Test2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(m_8.UnsafePtr).V1}))})
						goto end_branch_6
					} else {

					}
				}
				{
					if m_8.Type == 9 && m_8.IntVal == 227416251 {
						// TAST (Let): __local_var_9_4 shape=Other bindingType=Any
						__local_var_9_4 := (*Constructor_Main_Test3[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(m_8.UnsafePtr).V0
						_ = __local_var_9_4
						__t6 = gopurs_runtime.Apply2(Apply0_4_2.V1, gopurs_runtime.Apply2(Apply0_4_2.V1, gopurs_runtime.Apply2(Functor0_5_3.V0, gopurs_runtime.Func3(func(v4_10 gopurs_runtime.Value, v5_11 gopurs_runtime.Value, v6_12 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Value{Type: 9, IntVal: 227416251, UnsafePtr: unsafe.Pointer((&Constructor_Main_Test3[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, __local_var_9_4, v4_10, v5_11, v6_12}))}
						}), gopurs_runtime.Apply4(gopurs_runtime.RecordGet(dictBitraversable_0, "bitraverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_3))}, f_6, g_7, (*Constructor_Main_Test3[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(m_8.UnsafePtr).V1)), gopurs_runtime.Apply4(gopurs_runtime.RecordGet(dictBitraversable_0, "bitraverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_3))}, f_6, Call_Control_Applicative_pure(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_3)), (*Constructor_Main_Test3[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(m_8.UnsafePtr).V2)), gopurs_runtime.Apply4(gopurs_runtime.RecordGet(dictBitraversable_0, "bitraverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_3))}, Call_Control_Applicative_pure(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_3)), g_7, (*Constructor_Main_Test3[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(m_8.UnsafePtr).V3))
						goto end_branch_6
					} else {

					}
				}
				{
					if m_8.Type == 9 && m_8.IntVal == 3712677948 {
						__t6 = gopurs_runtime.Apply2(Apply0_4_2.V1, gopurs_runtime.Apply2(Functor0_5_3.V0, gopurs_runtime.Func2(func(v2_9 gopurs_runtime.Value, v3_10 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Value{Type: 9, IntVal: 3712677948, UnsafePtr: unsafe.Pointer((&Constructor_Main_Test4[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, func() []*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, int64] {
								arr := *(*[]gopurs_runtime.Value)(v2_9.UnsafePtr)
								unboxed := make([]*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, int64], len(arr))
								for i, v := range arr {
									unboxed[i] = Rebox_Main_138441832_3415943795(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v))
								}
								return unboxed
							}(), Rebox_Main_138441832_3415943795(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v3_10))}))}
						}), gopurs_runtime.Apply3(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableArray()).V3, gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_3))}, gopurs_runtime.Apply3(Rebox_Main_3561684974_2243670499(gopurs_runtime.CoerceToStruct[Constructor_Data_Bitraversable_Bitraversable[gopurs_runtime.Value]](Get_Data_Bitraversable_bitraversableTuple())).V3, gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_3))}, f_6, Call_Control_Applicative_pure(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_3))), func() gopurs_runtime.Value {
							arr := (*Constructor_Main_Test4[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(m_8.UnsafePtr).V0
							boxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Main_3415943795_138441832(v))}
							}
							return gopurs_runtime.Array(boxed)
						}())), gopurs_runtime.Apply4(Rebox_Main_3561684974_2243670499(gopurs_runtime.CoerceToStruct[Constructor_Data_Bitraversable_Bitraversable[gopurs_runtime.Value]](Get_Data_Bitraversable_bitraversableTuple())).V3, gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_3))}, g_7, Call_Control_Applicative_pure(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_3)), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Main_3415943795_138441832((*Constructor_Main_Test4[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(m_8.UnsafePtr).V1))}))
						goto end_branch_6
					} else {

					}
				}
				{
					if m_8.Type == 9 && m_8.IntVal == 1063363133 {
						// TAST (Let): __local_var_9_5 shape=Other bindingType=Any
						__local_var_9_5 := (*Constructor_Main_Test5[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(m_8.UnsafePtr).V0
						_ = __local_var_9_5
						__t6 = gopurs_runtime.Apply2(Functor0_5_3.V0, gopurs_runtime.Func(func(v1_10 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Value{Type: 9, IntVal: 1063363133, UnsafePtr: unsafe.Pointer((&Constructor_Main_Test5[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, func() struct {
								nested []struct {
									x gopurs_runtime.Value
								}
							} {
								clone := __local_var_9_5
								clone.nested = func() []struct {
									x gopurs_runtime.Value
								} {
									arr := *(*[]gopurs_runtime.Value)(v1_10.UnsafePtr)
									unboxed := make([]struct {
										x gopurs_runtime.Value
									}, len(arr))
									for i, v := range arr {
										unboxed[i] = func() struct {
											x gopurs_runtime.Value
										} {
											orig := v
											_ = orig
											clone := struct {
												x gopurs_runtime.Value
											}{}
											clone.x = gopurs_runtime.RecordGet(orig, "x")
											return clone
										}()
									}
									return unboxed
								}()
								return clone
							}()}))}
						}), gopurs_runtime.Apply3(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableArray()).V3, gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_3))}, gopurs_runtime.Func(func(v1_10 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Apply2(Functor0_5_3.V0, gopurs_runtime.Func(func(v2_11 gopurs_runtime.Value) gopurs_runtime.Value {
								return func() gopurs_runtime.Value {
									orig := func() struct {
										x gopurs_runtime.Value
									} {
										orig := gopurs_runtime.RecordUpdate1(v1_10, "x", v2_11)
										_ = orig
										clone := struct {
											x gopurs_runtime.Value
										}{}
										clone.x = gopurs_runtime.RecordGet(orig, "x")
										return clone
									}()
									_ = orig
									return gopurs_runtime.RecordDict1("x", orig.x)
								}()
							}), gopurs_runtime.Apply4(gopurs_runtime.RecordGet(dictBitraversable_0, "bitraverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_3))}, gopurs_runtime.Func(func(v2_11 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Apply2(Functor0_5_3.V0, gopurs_runtime.Func(func(v3_12 gopurs_runtime.Value) gopurs_runtime.Value {
									return func() gopurs_runtime.Value {
										orig := func() struct {
											a gopurs_runtime.Value
										} {
											orig := gopurs_runtime.RecordUpdate1(v2_11, "a", v3_12)
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
								}), gopurs_runtime.Apply(f_6, gopurs_runtime.RecordGet(v2_11, "a")))
							}), gopurs_runtime.Func(func(v2_11 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Apply2(Functor0_5_3.V0, gopurs_runtime.Func(func(v3_12 gopurs_runtime.Value) gopurs_runtime.Value {
									return func() gopurs_runtime.Value {
										orig := func() struct {
											b gopurs_runtime.Value
										} {
											orig := gopurs_runtime.RecordUpdate1(v2_11, "b", v3_12)
											_ = orig
											clone := struct {
												b gopurs_runtime.Value
											}{}
											clone.b = gopurs_runtime.RecordGet(orig, "b")
											return clone
										}()
										_ = orig
										return gopurs_runtime.RecordDict1("b", orig.b)
									}()
								}), gopurs_runtime.Apply(g_7, gopurs_runtime.RecordGet(v2_11, "b")))
							}), gopurs_runtime.RecordGet(v1_10, "x")))
						}), func() gopurs_runtime.Value {
							arr := __local_var_9_5.nested
							boxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								boxed[i] = func() gopurs_runtime.Value {
									orig := v
									_ = orig
									return gopurs_runtime.RecordDict1("x", orig.x)
								}()
							}
							return gopurs_runtime.Array(boxed)
						}()))
						goto end_branch_6
					} else {

					}
				}
				{
					__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
				}
			end_branch_6:
				return __t6
			})
		})}))}
	}
}

func Rebox_Main_1048926258_1688994542(in *Constructor_Data_Bifunctor_Bifunctor[*Constructor_Main_FromProAndContra[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Bifunctor_Bifunctor[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Bifunctor_Bifunctor[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_138441832_3415943795(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, int64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, int64]{}
	out.V0 = in.V0
	out.V1 = in.V1.IntVal
	return out
}

func Rebox_Main_2812820739_3566843086(in *Constructor_Data_Bifoldable_Bifoldable[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value]{}
	out.V0 = in.V0
	out.V1 = in.V1
	out.V2 = in.V2
	return out
}

func Rebox_Main_3415943795_138441832(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, int64]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = in.V0
	out.V1 = gopurs_runtime.Int(in.V1)
	return out
}

func Rebox_Main_3561684974_2243670499(in *Constructor_Data_Bitraversable_Bitraversable[gopurs_runtime.Value]) *Constructor_Data_Bitraversable_Bitraversable[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Bitraversable_Bitraversable[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{}
	out.V0 = in.V0
	out.V1 = in.V1
	out.V2 = in.V2
	out.V3 = in.V3
	return out
}

func Rebox_Main_3566843086_2812820739(in *Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value]) *Constructor_Data_Bifoldable_Bifoldable[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Bifoldable_Bifoldable[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{}
	out.V0 = in.V0
	out.V1 = in.V1
	out.V2 = in.V2
	return out
}
