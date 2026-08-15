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
		cache_Main_void = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_void(a_0_box)
		})
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
				return gopurs_runtime.Value{Type: 9, IntVal: 322902991, UnsafePtr: unsafe.Pointer((&Constructor_Main_Cons{1, value0, gopurs_runtime.CoerceToStruct[Constructor_Main_Cons](value1)}))}
			})
		})
	})
	return cache_Main_Cons
}

var cache_Main_Nil gopurs_runtime.Value
var once_Main_Nil sync.Once

func Get_Main_Nil() gopurs_runtime.Value {
	once_Main_Nil.Do(func() {
		cache_Main_Nil = gopurs_runtime.Value{Type: 9, IntVal: 322902991, UnsafePtr: unsafe.Pointer((*Constructor_Main_Cons)(nil))}
	})
	return cache_Main_Nil
}

var cache_Main_sequenceListSeq gopurs_runtime.Value
var once_Main_sequenceListSeq sync.Once

func Get_Main_sequenceListSeq() gopurs_runtime.Value {
	once_Main_sequenceListSeq.Do(func() {
		cache_Main_sequenceListSeq = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_sequenceListSeq(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](dictMonad_0_box))
		})
	})
	return cache_Main_sequenceListSeq
}

var cache_Main_sequenceList_prime__prime_ gopurs_runtime.Value
var once_Main_sequenceList_prime__prime_ sync.Once

func Get_Main_sequenceList_prime__prime_() gopurs_runtime.Value {
	once_Main_sequenceList_prime__prime_.Do(func() {
		cache_Main_sequenceList_prime__prime_ = gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): Applicative0_1_0 -> *Constructor_Control_Applicative_Applicative
			Applicative0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
			_ = Applicative0_1_0
			// TAST (Let): Bind1_2_1 -> gopurs_runtime.Value
			Bind1_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{})
			_ = Bind1_2_1
			// TAST (Let): Apply0_3_2 -> *Constructor_Control_Apply_Apply
			Apply0_3_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Bind1_2_1, "Apply0"), gopurs_runtime.Value{}))
			_ = Apply0_3_2
			// TAST (Let): Functor0_4_3 -> *Constructor_Data_Functor_Functor
			Functor0_4_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Bind1_2_1, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
			_ = Functor0_4_3
			return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t4 gopurs_runtime.Value
				{
					if v_5.Type == 9 && v_5.IntVal == 322902991 && v_5.UnsafePtr == nil {
						__t4 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_1_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 322902991, UnsafePtr: unsafe.Pointer((*Constructor_Main_Cons)(nil))})
						goto end_branch_4
					} else {

					}
				}
				{
					if v_5.Type == 9 && v_5.IntVal == 322902991 && v_5.UnsafePtr != nil {
						__t4 = gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_3_2.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_4_3.V0), Get_Main_Cons(), (*Constructor_Main_Cons)(v_5.UnsafePtr).V0), gopurs_runtime.Apply(Call_Main_sequenceListSeq(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](dictMonad_0)), gopurs_runtime.Value{Type: 9, IntVal: 322902991, UnsafePtr: unsafe.Pointer((*Constructor_Main_Cons)(v_5.UnsafePtr).V1)}))
						goto end_branch_4
					} else {

					}
				}
				{
					__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
				}
			end_branch_4:
				return __t4
			})
		})
	})
	return cache_Main_sequenceList_prime__prime_
}

var cache_Main_sequenceList gopurs_runtime.Value
var once_Main_sequenceList sync.Once

func Get_Main_sequenceList() gopurs_runtime.Value {
	once_Main_sequenceList.Do(func() {
		cache_Main_sequenceList = gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): Applicative0_1_0 -> *Constructor_Control_Applicative_Applicative
			Applicative0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
			_ = Applicative0_1_0
			// TAST (Let): Bind1_2_1 -> gopurs_runtime.Value
			Bind1_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{})
			_ = Bind1_2_1
			// TAST (Let): Apply0_3_2 -> *Constructor_Control_Apply_Apply
			Apply0_3_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Bind1_2_1, "Apply0"), gopurs_runtime.Value{}))
			_ = Apply0_3_2
			// TAST (Let): Functor0_4_3 -> *Constructor_Data_Functor_Functor
			Functor0_4_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Bind1_2_1, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
			_ = Functor0_4_3
			return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t4 gopurs_runtime.Value
				{
					if v_5.Type == 9 && v_5.IntVal == 322902991 && v_5.UnsafePtr == nil {
						__t4 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_1_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 322902991, UnsafePtr: unsafe.Pointer((*Constructor_Main_Cons)(nil))})
						goto end_branch_4
					} else {

					}
				}
				{
					if v_5.Type == 9 && v_5.IntVal == 322902991 && v_5.UnsafePtr != nil {
						__t4 = gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_3_2.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_4_3.V0), Get_Main_Cons(), (*Constructor_Main_Cons)(v_5.UnsafePtr).V0), gopurs_runtime.Apply(Call_Main_sequenceListSeq(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](dictMonad_0)), gopurs_runtime.Value{Type: 9, IntVal: 322902991, UnsafePtr: unsafe.Pointer((*Constructor_Main_Cons)(v_5.UnsafePtr).V1)}))
						goto end_branch_4
					} else {

					}
				}
				{
					__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
				}
			end_branch_4:
				return __t4
			})
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
			// TAST (Let): Applicative0_1_0 -> *Constructor_Control_Applicative_Applicative
			Applicative0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
			_ = Applicative0_1_0
			// TAST (Let): Bind1_2_1 -> gopurs_runtime.Value
			Bind1_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{})
			_ = Bind1_2_1
			// TAST (Let): Apply0_3_2 -> *Constructor_Control_Apply_Apply
			Apply0_3_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Bind1_2_1, "Apply0"), gopurs_runtime.Value{}))
			_ = Apply0_3_2
			// TAST (Let): Functor0_4_3 -> *Constructor_Data_Functor_Functor
			Functor0_4_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Bind1_2_1, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
			_ = Functor0_4_3
			return gopurs_runtime.Func(func(val_5 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t4 gopurs_runtime.Value
				{
					if val_5.Type == 9 && val_5.IntVal == 322902991 && val_5.UnsafePtr == nil {
						__t4 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_1_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 322902991, UnsafePtr: unsafe.Pointer((*Constructor_Main_Cons)(nil))})
						goto end_branch_4
					} else {

					}
				}
				{
					if val_5.Type == 9 && val_5.IntVal == 322902991 && val_5.UnsafePtr != nil {
						__t4 = gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_3_2.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_4_3.V0), Get_Main_Cons(), (*Constructor_Main_Cons)(val_5.UnsafePtr).V0), gopurs_runtime.Apply2(Get_Main_sequenceList_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](dictMonad_0))}, gopurs_runtime.Value{Type: 9, IntVal: 322902991, UnsafePtr: unsafe.Pointer((*Constructor_Main_Cons)(val_5.UnsafePtr).V1)}))
						goto end_branch_4
					} else {

					}
				}
				{
					__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
				}
			end_branch_4:
				return __t4
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
			// TAST (Let): Applicative0_1_0 -> *Constructor_Control_Applicative_Applicative
			Applicative0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
			_ = Applicative0_1_0
			// TAST (Let): Bind1_2_1 -> gopurs_runtime.Value
			Bind1_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{})
			_ = Bind1_2_1
			// TAST (Let): Apply0_3_2 -> *Constructor_Control_Apply_Apply
			Apply0_3_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Bind1_2_1, "Apply0"), gopurs_runtime.Value{}))
			_ = Apply0_3_2
			// TAST (Let): Functor0_4_3 -> *Constructor_Data_Functor_Functor
			Functor0_4_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Bind1_2_1, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
			_ = Functor0_4_3
			return gopurs_runtime.Func(func(val_5 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t4 gopurs_runtime.Value
				{
					if val_5.Type == 9 && val_5.IntVal == 322902991 && val_5.UnsafePtr == nil {
						__t4 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_1_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 322902991, UnsafePtr: unsafe.Pointer((*Constructor_Main_Cons)(nil))})
						goto end_branch_4
					} else {

					}
				}
				{
					if val_5.Type == 9 && val_5.IntVal == 322902991 && val_5.UnsafePtr != nil {
						__t4 = gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_3_2.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_4_3.V0), Get_Main_Cons(), (*Constructor_Main_Cons)(val_5.UnsafePtr).V0), gopurs_runtime.Apply2(Get_Main_sequenceList_prime__prime__prime_(), gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](dictMonad_0))}, gopurs_runtime.Value{Type: 9, IntVal: 322902991, UnsafePtr: unsafe.Pointer((*Constructor_Main_Cons)(val_5.UnsafePtr).V1)}))
						goto end_branch_4
					} else {

					}
				}
				{
					__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
				}
			end_branch_4:
				return __t4
			})
		})
	})
	return cache_Main_sequenceList_prime__prime__prime_
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 -> gopurs_runtime.Value
			__local_var_0_0 := gopurs_runtime.Apply2(Get_Main_sequenceList(), gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](Get_Effect_monadEffect()))}, gopurs_runtime.Value{Type: 9, IntVal: 322902991, UnsafePtr: unsafe.Pointer((&Constructor_Main_Cons{1, gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), (*Constructor_Main_Cons)(nil)}))})
			_ = __local_var_0_0
			a_prime__1_2 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = a_prime__1_2
			_dollar___unused_1_1 := Get_Data_Unit_unit()
			_ = _dollar___unused_1_1
			a_prime__2_4 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Main_sequenceList_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](Get_Effect_monadEffect()))}, gopurs_runtime.Value{Type: 9, IntVal: 322902991, UnsafePtr: unsafe.Pointer((&Constructor_Main_Cons{1, gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), (*Constructor_Main_Cons)(nil)}))}), gopurs_runtime.Value{})
			_ = a_prime__2_4
			_dollar___unused_2_3 := Get_Data_Unit_unit()
			_ = _dollar___unused_2_3
			a_prime__3_6 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Main_sequenceList_prime__prime_(), gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](Get_Effect_monadEffect()))}, gopurs_runtime.Value{Type: 9, IntVal: 322902991, UnsafePtr: unsafe.Pointer((&Constructor_Main_Cons{1, gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), (*Constructor_Main_Cons)(nil)}))}), gopurs_runtime.Value{})
			_ = a_prime__3_6
			_dollar___unused_3_5 := Get_Data_Unit_unit()
			_ = _dollar___unused_3_5
			a_prime__4_7 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Main_sequenceList_prime__prime__prime_(), gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](Get_Effect_monadEffect()))}, gopurs_runtime.Value{Type: 9, IntVal: 322902991, UnsafePtr: unsafe.Pointer((&Constructor_Main_Cons{1, gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), (*Constructor_Main_Cons)(nil)}))}), gopurs_runtime.Value{})
			_ = a_prime__4_7
			return Get_Data_Unit_unit()
		})
	})
	return cache_Main_main
}

var cache_Main_sequence__4065837736 gopurs_runtime.Value
var once_Main_sequence__4065837736 sync.Once

func Get_Main_sequence__4065837736() gopurs_runtime.Value {
	once_Main_sequence__4065837736.Do(func() {
		cache_Main_sequence__4065837736 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, dictMonad_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_sequence__4065837736(v_0_box, dictMonad_1_box)
		})
	})
	return cache_Main_sequence__4065837736
}

var cache_Main_sequenceListSeq__2621562580 gopurs_runtime.Value
var once_Main_sequenceListSeq__2621562580 sync.Once

func Get_Main_sequenceListSeq__2621562580() gopurs_runtime.Value {
	once_Main_sequenceListSeq__2621562580.Do(func() {
		cache_Main_sequenceListSeq__2621562580 = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_sequenceListSeq__2621562580(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](dictMonad_0_box))
		})
	})
	return cache_Main_sequenceListSeq__2621562580
}

type Constructor_Main_Sequence struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

type Constructor_Main_Cons struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 *Constructor_Main_Cons
}

type Constructor_Main_Nil struct {
	Rc uint32
}

func Call_Main_void(a_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var a_0 gopurs_runtime.Value = a_0_loop
	_ = a_0
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		a_prime__1_0 := gopurs_runtime.Apply(a_0, gopurs_runtime.Value{})
		_ = a_prime__1_0
		return Get_Data_Unit_unit()
	})
}

func Call_Main_sequenceListSeq(dictMonad_0_loop *Constructor_Control_Monad_Monad) gopurs_runtime.Value {
sequenceListSeq:
	for {
		if false {
			continue sequenceListSeq
		}
		var dictMonad_0 *Constructor_Control_Monad_Monad = dictMonad_0_loop
		_ = dictMonad_0
		// TAST (Let): Applicative0_1_0 -> *Constructor_Control_Applicative_Applicative
		Applicative0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V0), gopurs_runtime.Value{}))
		_ = Applicative0_1_0
		// TAST (Let): Bind1_2_1 -> gopurs_runtime.Value
		Bind1_2_1 := gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V1), gopurs_runtime.Value{})
		_ = Bind1_2_1
		// TAST (Let): Apply0_3_2 -> *Constructor_Control_Apply_Apply
		Apply0_3_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Bind1_2_1, "Apply0"), gopurs_runtime.Value{}))
		_ = Apply0_3_2
		// TAST (Let): Functor0_4_3 -> *Constructor_Data_Functor_Functor
		Functor0_4_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Bind1_2_1, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
		_ = Functor0_4_3
		return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t4 gopurs_runtime.Value
			{
				if v_5.Type == 9 && v_5.IntVal == 322902991 && v_5.UnsafePtr == nil {
					__t4 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_1_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 322902991, UnsafePtr: unsafe.Pointer((*Constructor_Main_Cons)(nil))})
					goto end_branch_4
				} else {

				}
			}
			{
				if v_5.Type == 9 && v_5.IntVal == 322902991 && v_5.UnsafePtr != nil {
					__t4 = gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_3_2.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_4_3.V0), Get_Main_Cons(), (*Constructor_Main_Cons)(v_5.UnsafePtr).V0), gopurs_runtime.Apply(Call_Main_sequenceListSeq(dictMonad_0), gopurs_runtime.Value{Type: 9, IntVal: 322902991, UnsafePtr: unsafe.Pointer((*Constructor_Main_Cons)(v_5.UnsafePtr).V1)}))
					goto end_branch_4
				} else {

				}
			}
			{
				__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_4:
			return __t4
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

func Call_Main_sequence__4065837736(v_0_loop gopurs_runtime.Value, dictMonad_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	var dictMonad_1 gopurs_runtime.Value = dictMonad_1_loop
	_ = dictMonad_1
	return gopurs_runtime.Apply(v_0, dictMonad_1)
}

func Call_Main_sequenceListSeq__2621562580(dictMonad_0_loop *Constructor_Control_Monad_Monad) gopurs_runtime.Value {
	var dictMonad_0 *Constructor_Control_Monad_Monad = dictMonad_0_loop
	_ = dictMonad_0
	// TAST (Let): Applicative0_1_0 -> *Constructor_Control_Applicative_Applicative
	Applicative0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V0), gopurs_runtime.Value{}))
	_ = Applicative0_1_0
	// TAST (Let): Bind1_2_1 -> gopurs_runtime.Value
	Bind1_2_1 := gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V1), gopurs_runtime.Value{})
	_ = Bind1_2_1
	// TAST (Let): Apply0_3_2 -> *Constructor_Control_Apply_Apply
	Apply0_3_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Bind1_2_1, "Apply0"), gopurs_runtime.Value{}))
	_ = Apply0_3_2
	// TAST (Let): Functor0_4_3 -> *Constructor_Data_Functor_Functor
	Functor0_4_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Bind1_2_1, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
	_ = Functor0_4_3
	return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
		var __t11 gopurs_runtime.Value
		{
			if v_5.Type == 9 && v_5.IntVal == 322902991 && v_5.UnsafePtr == nil {
				__t11 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_1_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 322902991, UnsafePtr: unsafe.Pointer((*Constructor_Main_Cons)(nil))})
				goto end_branch_11
			} else {

			}
		}
		{
			if v_5.Type == 9 && v_5.IntVal == 322902991 && v_5.UnsafePtr != nil {
				// TAST (Let): Applicative0_6_4 -> *Constructor_Control_Applicative_Applicative
				Applicative0_6_4 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V0), gopurs_runtime.Value{}))
				_ = Applicative0_6_4
				// TAST (Let): Bind1_7_5 -> gopurs_runtime.Value
				Bind1_7_5 := gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V1), gopurs_runtime.Value{})
				_ = Bind1_7_5
				// TAST (Let): Apply0_8_6 -> *Constructor_Control_Apply_Apply
				Apply0_8_6 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Bind1_7_5, "Apply0"), gopurs_runtime.Value{}))
				_ = Apply0_8_6
				// TAST (Let): Functor0_9_7 -> *Constructor_Data_Functor_Functor
				Functor0_9_7 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Bind1_7_5, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
				_ = Functor0_9_7
				var __t10 gopurs_runtime.Value
				{
					var __t_tag_8 *Constructor_Main_Cons = (*Constructor_Main_Cons)(v_5.UnsafePtr).V1
					if __t_tag_8 == nil {
						__t10 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_6_4.V1), gopurs_runtime.Value{Type: 9, IntVal: 322902991, UnsafePtr: unsafe.Pointer((*Constructor_Main_Cons)(nil))})
						goto end_branch_10
					} else {

					}
				}
				{
					var __t_tag_9 *Constructor_Main_Cons = (*Constructor_Main_Cons)(v_5.UnsafePtr).V1
					if __t_tag_9 != nil {
						__t10 = gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_8_6.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_9_7.V0), Get_Main_Cons(), ((*Constructor_Main_Cons)(v_5.UnsafePtr).V1).V0), gopurs_runtime.Apply(Call_Main_sequenceListSeq(dictMonad_0), gopurs_runtime.Value{Type: 9, IntVal: 322902991, UnsafePtr: unsafe.Pointer(((*Constructor_Main_Cons)(v_5.UnsafePtr).V1).V1)}))
						goto end_branch_10
					} else {

					}
				}
				{
					__t10 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
				}
			end_branch_10:
				__t11 = gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_3_2.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_4_3.V0), Get_Main_Cons(), (*Constructor_Main_Cons)(v_5.UnsafePtr).V0), __t10)
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
}
