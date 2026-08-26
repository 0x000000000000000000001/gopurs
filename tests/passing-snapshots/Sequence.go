package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_Sequence_dollar_Dict gopurs_runtime.Value
var once_Main_Sequence_dollar_Dict sync.Once

func Get_Main_Sequence_dollar_Dict() gopurs_runtime.Value {
	once_Main_Sequence_dollar_Dict.Do(func() {
		cache_Main_Sequence_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_Sequence_dollar_Dict(x_0_box)
		})
	})
	return cache_Main_Sequence_dollar_Dict
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

var cache_Main_sequence gopurs_runtime.Value
var once_Main_sequence sync.Once

func Get_Main_sequence() gopurs_runtime.Value {
	once_Main_sequence.Do(func() {
		cache_Main_sequence = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_sequence(gopurs_runtime.CoerceToStruct[Constructor_Main_Sequence[gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Main_sequence
}

var cache_Main_sequence__4029670178 gopurs_runtime.Value
var once_Main_sequence__4029670178 sync.Once

func Get_Main_sequence__4029670178() gopurs_runtime.Value {
	once_Main_sequence__4029670178.Do(func() {
		cache_Main_sequence__4029670178 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_sequence__4029670178(gopurs_runtime.CoerceToStruct[Constructor_Main_Sequence[gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Main_sequence__4029670178
}

var cache_Main_sequenceList gopurs_runtime.Value
var once_Main_sequenceList sync.Once

func Get_Main_sequenceList() gopurs_runtime.Value {
	once_Main_sequenceList.Do(func() {
		cache_Main_sequenceList = gopurs_runtime.Value{Type: 9, IntVal: 3858394017, UnsafePtr: unsafe.Pointer((&Constructor_Main_Sequence[*Constructor_Main_Cons[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): Applicative0_1_0 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m)])
			Applicative0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
			_ = Applicative0_1_0
			// TAST (Let): Bind1_2_1 shape=App(Other) bindingType=Any
			Bind1_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{})
			_ = Bind1_2_1
			// TAST (Let): Apply0_3_2 shape=App(Other) bindingType=(ADT ["Control","Apply","Apply"] [(TypeVar m)])
			Apply0_3_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Bind1_2_1, "Apply0"), gopurs_runtime.Value{}))
			_ = Apply0_3_2
			// TAST (Let): Functor0_4_3 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
			Functor0_4_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Bind1_2_1, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
			_ = Functor0_4_3
			return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t4 gopurs_runtime.Value
				{
					if v_5.Type == 9 && v_5.IntVal == 322902991 && v_5.UnsafePtr == nil {
						__t4 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_1_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 322902991, UnsafePtr: unsafe.Pointer((*Constructor_Main_Cons[gopurs_runtime.Value])(nil))})
						goto end_branch_4
					} else {

					}
				}
				{
					if v_5.Type == 9 && v_5.IntVal == 322902991 && v_5.UnsafePtr != nil {
						__t4 = gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_3_2.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_4_3.V0), Get_Main_Cons(), (*Constructor_Main_Cons[gopurs_runtime.Value])(v_5.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Main_Sequence[*Constructor_Main_Cons[gopurs_runtime.Value]]](Get_Main_sequenceList()).V0), gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](dictMonad_0))}, gopurs_runtime.Value{Type: 9, IntVal: 322902991, UnsafePtr: unsafe.Pointer((*Constructor_Main_Cons[gopurs_runtime.Value])(v_5.UnsafePtr).V1)}))
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
		})}))}
	})
	return cache_Main_sequenceList
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=Other bindingType=(ADT ["Main","List"] [(ADT ["Effect","Effect"] [Unit])])
			__local_var_0_0 := (&Constructor_Main_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), (*Constructor_Main_Cons[gopurs_runtime.Value])(nil)})
			_ = __local_var_0_0
			// TAST (Let): __local_var_1_1 shape=Other bindingType=(TypeApp (TypeVar m) [(TypeVar a)])
			__local_var_1_1 := (__local_var_0_0).V0
			_ = __local_var_1_1
			// TAST (Let): __local_var_2_2 shape=App(Other) bindingType=Any
			__local_var_2_2 := gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Main_Sequence[*Constructor_Main_Cons[gopurs_runtime.Value]]](Get_Main_sequenceList()).V0), gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](Get_Effect_monadEffect()))}, gopurs_runtime.Value{Type: 9, IntVal: 322902991, UnsafePtr: unsafe.Pointer((__local_var_0_0).V1)})
			_ = __local_var_2_2
			a_prime__3_4 := gopurs_runtime.Apply(__local_var_1_1, gopurs_runtime.Value{})
			_ = a_prime__3_4
			f_prime__3_3 := gopurs_runtime.Apply(Get_Main_Cons(), a_prime__3_4)
			_ = f_prime__3_3
			a_prime__4_5 := gopurs_runtime.Apply(__local_var_2_2, gopurs_runtime.Value{})
			_ = a_prime__4_5
			return gopurs_runtime.Apply(f_prime__3_3, a_prime__4_5)
		})
	})
	return cache_Main_main
}

type Constructor_Main_Cons[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 *Constructor_Main_Cons[gopurs_runtime.Value]
}

type Constructor_Main_Nil[T_a any] struct {
	Rc uint32
}

type Constructor_Main_Sequence[T_t any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

func init() {
	gopurs_runtime.StructGetters[3858394017] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_Sequence[any])(ptr)
		_ = c
		switch key {
		case "sequence":
			return gopurs_runtime.Box(c.V0)
		default:
			panic("Key not found in dictionary Constructor_Main_Sequence: " + key)
		}
	}
}

func Call_Main_Sequence_dollar_Dict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_sequence(dict_0_loop *Constructor_Main_Sequence[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_Sequence[gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return gopurs_runtime.Box(dict_0.V0)
}

func Call_Main_sequence__4029670178(dict_0_loop *Constructor_Main_Sequence[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_Sequence[gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return gopurs_runtime.Box(dict_0.V0)
}
