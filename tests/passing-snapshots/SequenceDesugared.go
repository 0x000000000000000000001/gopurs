package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_void gopurs_runtime.Value
var once_Main_void sync.Once

func Get_Main_void() gopurs_runtime.Value {
	once_Main_void.Do(func() {
		cache_Main_void = Call_Data_Functor_void(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()))
	})
	return cache_Main_void
}

var cache_Main_Sequence gopurs_runtime.Value
var once_Main_Sequence sync.Once

func Get_Main_Sequence() gopurs_runtime.Value {
	once_Main_Sequence.Do(func() {
		cache_Main_Sequence = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return value0
		})
	})
	return cache_Main_Sequence
}

var cache_Main_Cons gopurs_runtime.Value
var once_Main_Cons sync.Once

func Get_Main_Cons() gopurs_runtime.Value {
	once_Main_Cons.Do(func() {
		cache_Main_Cons = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 322902991, UnsafePtr: unsafe.Pointer((&Constructor_Main_Cons[gopurs_runtime.Value]{1, value0, gopurs_runtime.CoerceToStruct[Constructor_Main_Cons[gopurs_runtime.Value]](value1)}))}
			})
		})
	})
	return cache_Main_Cons
}

var cache_Main_Nil gopurs_runtime.Value
var once_Main_Nil sync.Once

func Get_Main_Nil() gopurs_runtime.Value {
	once_Main_Nil.Do(func() {
		cache_Main_Nil = gopurs_runtime.Value{Type: 9, IntVal: 322902991, UnsafePtr: unsafe.Pointer((*Constructor_Main_Cons[gopurs_runtime.Value])(nil))}
	})
	return cache_Main_Nil
}

var cache_Main_Cons__1817392478 gopurs_runtime.Value
var once_Main_Cons__1817392478 sync.Once

func Get_Main_Cons__1817392478() gopurs_runtime.Value {
	once_Main_Cons__1817392478.Do(func() {
		cache_Main_Cons__1817392478 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 322902991, UnsafePtr: unsafe.Pointer(Call_Main_Cons__1817392478(__eta_norm_1_0_box, gopurs_runtime.CoerceToStruct[Constructor_Main_Cons[gopurs_runtime.Value]](__eta_norm_0_unused_1_box)))}
		})
	})
	return cache_Main_Cons__1817392478
}

var cache_Main_sequenceListSeq gopurs_runtime.Value
var once_Main_sequenceListSeq sync.Once

func Get_Main_sequenceListSeq() gopurs_runtime.Value {
	once_Main_sequenceListSeq.Do(func() {
		cache_Main_sequenceListSeq = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_sequenceListSeq(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](dictMonad_0_box))
		})
	})
	return cache_Main_sequenceListSeq
}

var cache_Main_sequenceList_prime__prime_ gopurs_runtime.Value
var once_Main_sequenceList_prime__prime_ sync.Once

func Get_Main_sequenceList_prime__prime_() gopurs_runtime.Value {
	once_Main_sequenceList_prime__prime_.Do(func() {
		cache_Main_sequenceList_prime__prime_ = gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_sequenceListSeq(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](dictMonad_0))
		})
	})
	return cache_Main_sequenceList_prime__prime_
}

var cache_Main_sequenceList gopurs_runtime.Value
var once_Main_sequenceList sync.Once

func Get_Main_sequenceList() gopurs_runtime.Value {
	once_Main_sequenceList.Do(func() {
		cache_Main_sequenceList = gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_sequenceListSeq(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](dictMonad_0))
		})
	})
	return cache_Main_sequenceList
}

var cache_Main_sequence gopurs_runtime.Value
var once_Main_sequence sync.Once

func Get_Main_sequence() gopurs_runtime.Value {
	once_Main_sequence.Do(func() {
		cache_Main_sequence = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, dictMonad_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_sequence(v_0_box, dictMonad_1_box)
		})
	})
	return cache_Main_sequence
}

var cache_Main_sequenceList_prime_ gopurs_runtime.Value
var once_Main_sequenceList_prime_ sync.Once

func Get_Main_sequenceList_prime_() gopurs_runtime.Value {
	once_Main_sequenceList_prime_.Do(func() {
		cache_Main_sequenceList_prime_ = gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): Applicative0_1_0 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m$scope33)])
			Applicative0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
			_ = Applicative0_1_0
			// TAST (Let): Bind1_2_1 shape=App(Other) bindingType=Any
			Bind1_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{})
			_ = Bind1_2_1
			// TAST (Let): Apply0_3_2 shape=App(Other) bindingType=(ADT ["Control","Apply","Apply"] [(TypeVar m$scope33)])
			Apply0_3_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Bind1_2_1, "Apply0"), gopurs_runtime.Value{}))
			_ = Apply0_3_2
			// TAST (Let): Functor0_4_3 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m$scope33)])
			Functor0_4_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Bind1_2_1, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
			_ = Functor0_4_3
			return gopurs_runtime.Func(func(val_5 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t6 gopurs_runtime.Value
				{
					var __t_tag_4 *Constructor_Main_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Main_Cons[gopurs_runtime.Value]](val_5)
					_ = __t_tag_4
					if __t_tag_4 == nil {
						__t6 = gopurs_runtime.Apply(Applicative0_1_0.V1, gopurs_runtime.Value{Type: 9, IntVal: 322902991, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Main_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 322902991, UnsafePtr: unsafe.Pointer((*Constructor_Main_Cons[gopurs_runtime.Value])(nil))}))})
						goto end_branch_6
					} else {

					}
				}
				{
					var __t_tag_5 *Constructor_Main_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Main_Cons[gopurs_runtime.Value]](val_5)
					_ = __t_tag_5
					if __t_tag_5 != nil {
						__t6 = gopurs_runtime.Apply2(Apply0_3_2.V1, gopurs_runtime.Apply2(Functor0_4_3.V0, Get_Main_Cons(), (*Constructor_Main_Cons[gopurs_runtime.Value])(val_5.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.Func(func(dictMonad_6 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Apply(Get_Main_sequenceList_prime_(), dictMonad_6)
						}), gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](dictMonad_0))}, gopurs_runtime.Value{Type: 9, IntVal: 322902991, UnsafePtr: unsafe.Pointer((*Constructor_Main_Cons[gopurs_runtime.Value])(val_5.UnsafePtr).V1)}))
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
		})
	})
	return cache_Main_sequenceList_prime_
}

var cache_Main_sequenceList_prime__prime__prime_ gopurs_runtime.Value
var once_Main_sequenceList_prime__prime__prime_ sync.Once

func Get_Main_sequenceList_prime__prime__prime_() gopurs_runtime.Value {
	once_Main_sequenceList_prime__prime__prime_.Do(func() {
		cache_Main_sequenceList_prime__prime__prime_ = gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): Applicative0_1_0 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m$scope44)])
			Applicative0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
			_ = Applicative0_1_0
			// TAST (Let): Bind1_2_1 shape=App(Other) bindingType=Any
			Bind1_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{})
			_ = Bind1_2_1
			// TAST (Let): Apply0_3_2 shape=App(Other) bindingType=(ADT ["Control","Apply","Apply"] [(TypeVar m$scope44)])
			Apply0_3_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Bind1_2_1, "Apply0"), gopurs_runtime.Value{}))
			_ = Apply0_3_2
			// TAST (Let): Functor0_4_3 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m$scope44)])
			Functor0_4_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Bind1_2_1, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
			_ = Functor0_4_3
			return gopurs_runtime.Func(func(val_5 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t6 gopurs_runtime.Value
				{
					var __t_tag_4 *Constructor_Main_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Main_Cons[gopurs_runtime.Value]](val_5)
					_ = __t_tag_4
					if __t_tag_4 == nil {
						__t6 = gopurs_runtime.Apply(Applicative0_1_0.V1, gopurs_runtime.Value{Type: 9, IntVal: 322902991, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Main_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 322902991, UnsafePtr: unsafe.Pointer((*Constructor_Main_Cons[gopurs_runtime.Value])(nil))}))})
						goto end_branch_6
					} else {

					}
				}
				{
					var __t_tag_5 *Constructor_Main_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Main_Cons[gopurs_runtime.Value]](val_5)
					_ = __t_tag_5
					if __t_tag_5 != nil {
						__t6 = gopurs_runtime.Apply2(Apply0_3_2.V1, gopurs_runtime.Apply2(Functor0_4_3.V0, Get_Main_Cons(), (*Constructor_Main_Cons[gopurs_runtime.Value])(val_5.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.Func(func(dictMonad_6 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Apply(Get_Main_sequenceList_prime__prime__prime_(), dictMonad_6)
						}), gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](dictMonad_0))}, gopurs_runtime.Value{Type: 9, IntVal: 322902991, UnsafePtr: unsafe.Pointer((*Constructor_Main_Cons[gopurs_runtime.Value])(val_5.UnsafePtr).V1)}))
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
		})
	})
	return cache_Main_sequenceList_prime__prime__prime_
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Call_Data_Functor_void(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect())), gopurs_runtime.Apply2(Get_Main_sequenceList(), gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](Get_Effect_monadEffect()))}, gopurs_runtime.Value{Type: 9, IntVal: 322902991, UnsafePtr: unsafe.Pointer((&Constructor_Main_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), (*Constructor_Main_Cons[gopurs_runtime.Value])(nil)}))})), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Call_Data_Functor_void(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect())), gopurs_runtime.Apply2(Get_Main_sequenceList_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](Get_Effect_monadEffect()))}, gopurs_runtime.Value{Type: 9, IntVal: 322902991, UnsafePtr: unsafe.Pointer((&Constructor_Main_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), (*Constructor_Main_Cons[gopurs_runtime.Value])(nil)}))})), gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Call_Data_Functor_void(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect())), gopurs_runtime.Apply2(Get_Main_sequenceList_prime__prime_(), gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](Get_Effect_monadEffect()))}, gopurs_runtime.Value{Type: 9, IntVal: 322902991, UnsafePtr: unsafe.Pointer((&Constructor_Main_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), (*Constructor_Main_Cons[gopurs_runtime.Value])(nil)}))})), gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply(Call_Data_Functor_void(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect())), gopurs_runtime.Apply2(Get_Main_sequenceList_prime__prime__prime_(), gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](Get_Effect_monadEffect()))}, gopurs_runtime.Value{Type: 9, IntVal: 322902991, UnsafePtr: unsafe.Pointer((&Constructor_Main_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), (*Constructor_Main_Cons[gopurs_runtime.Value])(nil)}))}))
				}))
			}))
		}))
	})
	return cache_Main_main
}

type Constructor_Main_Sequence[T_t any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

type Constructor_Main_Cons[T_a any] struct {
	Rc uint32
	V0 T_a
	V1 *Constructor_Main_Cons[T_a]
}

type Constructor_Main_Nil[T_a any] struct {
	Rc uint32
}

func Call_Main_Cons__1817392478(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop *Constructor_Main_Cons[gopurs_runtime.Value]) *Constructor_Main_Cons[gopurs_runtime.Value] {
Cons__1817392478:
	for {
		if false {
			continue Cons__1817392478
		}
		var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_unused_1 *Constructor_Main_Cons[gopurs_runtime.Value] = __eta_norm_0_unused_1_loop
		_ = __eta_norm_0_unused_1
		return (&Constructor_Main_Cons[gopurs_runtime.Value]{1, __eta_norm_1_0, (*Constructor_Main_Cons[gopurs_runtime.Value])(nil)})
	}
}

func Call_Main_sequenceListSeq(dictMonad_0_loop *Constructor_Control_Monad_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
sequenceListSeq:
	for {
		if false {
			continue sequenceListSeq
		}
		var dictMonad_0 *Constructor_Control_Monad_Monad[gopurs_runtime.Value] = dictMonad_0_loop
		_ = dictMonad_0
		// TAST (Let): Applicative0_1_0 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m$scope2)])
		Applicative0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonad_0.V0, gopurs_runtime.Value{}))
		_ = Applicative0_1_0
		// TAST (Let): Bind1_2_1 shape=App(Other) bindingType=Any
		Bind1_2_1 := gopurs_runtime.Apply(dictMonad_0.V1, gopurs_runtime.Value{})
		_ = Bind1_2_1
		// TAST (Let): Apply0_3_2 shape=App(Other) bindingType=(ADT ["Control","Apply","Apply"] [(TypeVar m$scope2)])
		Apply0_3_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Bind1_2_1, "Apply0"), gopurs_runtime.Value{}))
		_ = Apply0_3_2
		// TAST (Let): Functor0_4_3 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m$scope2)])
		Functor0_4_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Bind1_2_1, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
		_ = Functor0_4_3
		return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t6 gopurs_runtime.Value
			{
				var __t_tag_4 *Constructor_Main_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Main_Cons[gopurs_runtime.Value]](v_5)
				_ = __t_tag_4
				if __t_tag_4 == nil {
					__t6 = gopurs_runtime.Apply(Applicative0_1_0.V1, gopurs_runtime.Value{Type: 9, IntVal: 322902991, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Main_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 322902991, UnsafePtr: unsafe.Pointer((*Constructor_Main_Cons[gopurs_runtime.Value])(nil))}))})
					goto end_branch_6
				} else {

				}
			}
			{
				var __t_tag_5 *Constructor_Main_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Main_Cons[gopurs_runtime.Value]](v_5)
				_ = __t_tag_5
				if __t_tag_5 != nil {
					__t6 = gopurs_runtime.Apply2(Apply0_3_2.V1, gopurs_runtime.Apply2(Functor0_4_3.V0, Get_Main_Cons(), (*Constructor_Main_Cons[gopurs_runtime.Value])(v_5.UnsafePtr).V0), gopurs_runtime.Apply(Call_Main_sequenceListSeq(dictMonad_0), gopurs_runtime.Value{Type: 9, IntVal: 322902991, UnsafePtr: unsafe.Pointer((*Constructor_Main_Cons[gopurs_runtime.Value])(v_5.UnsafePtr).V1)}))
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
	}
}

func Call_Main_sequence(v_0_loop gopurs_runtime.Value, dictMonad_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	var dictMonad_1 gopurs_runtime.Value = dictMonad_1_loop
	_ = dictMonad_1
	return gopurs_runtime.Apply(v_0, dictMonad_1)
}
