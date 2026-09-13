package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

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

var cache_Main_Test2 gopurs_runtime.Value
var once_Main_Test2 sync.Once

func Get_Main_Test2() gopurs_runtime.Value {
	once_Main_Test2.Do(func() {
		cache_Main_Test2 = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2375191994, UnsafePtr: unsafe.Pointer((&Constructor_Main_Test2[gopurs_runtime.Value]{1, func() struct {
				x gopurs_runtime.Value
			} {
				orig := value0
				_ = orig
				clone := struct {
					x gopurs_runtime.Value
				}{}
				clone.x = gopurs_runtime.RecordGet(orig, "x")
				return clone
			}()}))}
		})
	})
	return cache_Main_Test2
}

var cache_Main_functorTest gopurs_runtime.Value
var once_Main_functorTest sync.Once

func Get_Main_functorTest() gopurs_runtime.Value {
	once_Main_functorTest.Do(func() {
		cache_Main_functorTest = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, m_1 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t0 gopurs_runtime.Value
			{
				if m_1.Type == 9 && m_1.IntVal == 3720114489 {
					__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3720114489, UnsafePtr: unsafe.Pointer((&Constructor_Main_Test1[gopurs_runtime.Value]{1, gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), (*Constructor_Main_Test1[gopurs_runtime.Value])(m_1.UnsafePtr).V0, gopurs_runtime.Apply(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Contravariant_Contravariant[gopurs_runtime.Value]](Get_Data_Predicate_contravariantPredicate()).V0, f_0))}))}
					goto end_branch_0
				} else {

				}
			}
			{
				if m_1.Type == 9 && m_1.IntVal == 2375191994 {
					__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2375191994, UnsafePtr: unsafe.Pointer((&Constructor_Main_Test2[gopurs_runtime.Value]{1, func() struct {
						x gopurs_runtime.Value
					} {
						clone := (*Constructor_Main_Test2[gopurs_runtime.Value])(m_1.UnsafePtr).V0
						clone.x = gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), (*Constructor_Main_Test2[gopurs_runtime.Value])(m_1.UnsafePtr).V0.x, gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
							return func() gopurs_runtime.Value {
								orig := func() struct {
									y gopurs_runtime.Value
								} {
									orig := gopurs_runtime.RecordUpdate1(v1_2, "y", gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.RecordGet(v1_2, "y"), f_0))
									_ = orig
									clone := struct {
										y gopurs_runtime.Value
									}{}
									clone.y = gopurs_runtime.RecordGet(orig, "y")
									return clone
								}()
								_ = orig
								return gopurs_runtime.RecordDict1("y", orig.y)
							}()
						}))
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

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
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
		x gopurs_runtime.Value
	}
}
