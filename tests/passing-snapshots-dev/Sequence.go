package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

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

var cache_Main_sequenceList gopurs_runtime.Value
var once_Main_sequenceList sync.Once

func Get_Main_sequenceList() gopurs_runtime.Value {
	once_Main_sequenceList.Do(func() {
		cache_Main_sequenceList = gopurs_runtime.Value{Type: 9, IntVal: 3858394017, UnsafePtr: unsafe.Pointer(Rebox_Main_2264881643_916613845((&Constructor_Main_Sequence[*Constructor_Main_Cons[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): Applicative0_1_0 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m$scope6)])
			Applicative0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
			_ = Applicative0_1_0
			// TAST (Let): Bind1_2_1 shape=App(Other) bindingType=Any
			Bind1_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{})
			_ = Bind1_2_1
			// TAST (Let): Apply0_3_2 shape=App(Other) bindingType=(ADT ["Control","Apply","Apply"] [(TypeVar m$scope6)])
			Apply0_3_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Bind1_2_1, "Apply0"), gopurs_runtime.Value{}))
			_ = Apply0_3_2
			// TAST (Let): Functor0_4_3 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m$scope6)])
			Functor0_4_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Bind1_2_1, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
			_ = Functor0_4_3
			return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t4 gopurs_runtime.Value
				{
					if v_5.Type == 9 && v_5.IntVal == 322902991 && v_5.UnsafePtr == nil {
						__t4 = gopurs_runtime.Apply(Call_Control_Applicative_pure(Applicative0_1_0), gopurs_runtime.Value{Type: 9, IntVal: 322902991, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Main_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 322902991, UnsafePtr: unsafe.Pointer((*Constructor_Main_Cons[gopurs_runtime.Value])(nil))}))})
						goto end_branch_4
					} else {

					}
				}
				{
					if v_5.Type == 9 && v_5.IntVal == 322902991 && v_5.UnsafePtr != nil {
						__t4 = gopurs_runtime.Apply2(Call_Control_Apply_apply(Apply0_3_2), gopurs_runtime.Apply2(Call_Data_Functor_go__map(Functor0_4_3), Get_Main_Cons(), (*Constructor_Main_Cons[gopurs_runtime.Value])(v_5.UnsafePtr).V0), gopurs_runtime.Apply2(Rebox_Main_916613845_2264881643(gopurs_runtime.CoerceToStruct[Constructor_Main_Sequence[gopurs_runtime.Value]](Get_Main_sequenceList())).V0, gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](dictMonad_0))}, gopurs_runtime.Value{Type: 9, IntVal: 322902991, UnsafePtr: unsafe.Pointer((*Constructor_Main_Cons[gopurs_runtime.Value])(v_5.UnsafePtr).V1)}))
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
		})})))}
	})
	return cache_Main_sequenceList
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = func() gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=Other bindingType=(ADT ["Main","List"] [(ADT ["Effect","Effect"] [Unit])])
			__local_var_0_0 := (&Constructor_Main_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), (*Constructor_Main_Cons[gopurs_runtime.Value])(nil)})
			_ = __local_var_0_0
			return gopurs_runtime.Apply2(Call_Control_Monad_ap(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](Get_Effect_monadEffect())), gopurs_runtime.Apply2(Call_Control_Applicative_liftA1(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Effect_applicativeEffect())), Get_Main_Cons(), (__local_var_0_0).V0), gopurs_runtime.Apply2(Rebox_Main_916613845_2264881643(gopurs_runtime.CoerceToStruct[Constructor_Main_Sequence[gopurs_runtime.Value]](Get_Main_sequenceList())).V0, gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](Get_Effect_monadEffect()))}, gopurs_runtime.Value{Type: 9, IntVal: 322902991, UnsafePtr: unsafe.Pointer((__local_var_0_0).V1)}))
		}()
	})
	return cache_Main_main
}

type Constructor_Main_Cons[T_a any] struct {
	Rc uint32
	V0 T_a
	V1 *Constructor_Main_Cons[T_a]
}

type Constructor_Main_Nil[T_a any] struct {
	Rc uint32
}

type Constructor_Main_Sequence[T_t any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

//gopurs:init-v1:00670065006e006500720061007400650064002f004d00610069006e:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f004d00610069006e005f00530065007100750065006e00630065
func init() {
	gopurs_runtime.StructGetters[3858394017] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_Sequence[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "sequence":
			return gopurs_runtime.Box(c.V0)
		default:
			panic("Key not found in dictionary Constructor_Main_Sequence: " + key)
		}
	}
}

func Call_Main_sequence(dict_0_loop *Constructor_Main_Sequence[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_Sequence[gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V0
}

func Rebox_Main_2264881643_916613845(in *Constructor_Main_Sequence[*Constructor_Main_Cons[gopurs_runtime.Value]]) *Constructor_Main_Sequence[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Main_Sequence[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_916613845_2264881643(in *Constructor_Main_Sequence[gopurs_runtime.Value]) *Constructor_Main_Sequence[*Constructor_Main_Cons[gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Main_Sequence[*Constructor_Main_Cons[gopurs_runtime.Value]])(unsafe.Pointer(in))
}
