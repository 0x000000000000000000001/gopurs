package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_List_Lazy_NonEmpty_compose gopurs_runtime.Value
var once_Data_List_Lazy_NonEmpty_compose sync.Once

func Get_Data_List_Lazy_NonEmpty_compose() gopurs_runtime.Value {
	once_Data_List_Lazy_NonEmpty_compose.Do(func() {
		cache_Data_List_Lazy_NonEmpty_compose = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_List_Lazy_NonEmpty_compose
}

var cache_Data_List_Lazy_NonEmpty_go__map gopurs_runtime.Value
var once_Data_List_Lazy_NonEmpty_go__map sync.Once

func Get_Data_List_Lazy_NonEmpty_go__map() gopurs_runtime.Value {
	once_Data_List_Lazy_NonEmpty_go__map.Do(func() {
		cache_Data_List_Lazy_NonEmpty_go__map = Rebox_Data_List_Lazy_NonEmpty_2812149806_3689823567(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Maybe_functorMaybe())).V0
	})
	return cache_Data_List_Lazy_NonEmpty_go__map
}

var cache_Data_List_Lazy_NonEmpty_add gopurs_runtime.Value
var once_Data_List_Lazy_NonEmpty_add sync.Once

func Get_Data_List_Lazy_NonEmpty_add() gopurs_runtime.Value {
	once_Data_List_Lazy_NonEmpty_add.Do(func() {
		cache_Data_List_Lazy_NonEmpty_add = Get_Data_Semiring_add__4086903470()
	})
	return cache_Data_List_Lazy_NonEmpty_add
}

var cache_Data_List_Lazy_NonEmpty_compose1 gopurs_runtime.Value
var once_Data_List_Lazy_NonEmpty_compose1 sync.Once

func Get_Data_List_Lazy_NonEmpty_compose1() gopurs_runtime.Value {
	once_Data_List_Lazy_NonEmpty_compose1.Do(func() {
		cache_Data_List_Lazy_NonEmpty_compose1 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_List_Lazy_NonEmpty_compose1
}

var cache_Data_List_Lazy_NonEmpty_go__append gopurs_runtime.Value
var once_Data_List_Lazy_NonEmpty_go__append sync.Once

func Get_Data_List_Lazy_NonEmpty_go__append() gopurs_runtime.Value {
	once_Data_List_Lazy_NonEmpty_go__append.Do(func() {
		cache_Data_List_Lazy_NonEmpty_go__append = gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_semigroupList(), "append")
	})
	return cache_Data_List_Lazy_NonEmpty_go__append
}

var cache_Data_List_Lazy_NonEmpty_uncons gopurs_runtime.Value
var once_Data_List_Lazy_NonEmpty_uncons sync.Once

func Get_Data_List_Lazy_NonEmpty_uncons() gopurs_runtime.Value {
	once_Data_List_Lazy_NonEmpty_uncons.Do(func() {
		cache_Data_List_Lazy_NonEmpty_uncons = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				orig := Call_Data_List_Lazy_NonEmpty_uncons(v_0_box)
				_ = orig
				return gopurs_runtime.RecordDict2("head", "tail", orig.head, orig.tail)
			}()
		})
	})
	return cache_Data_List_Lazy_NonEmpty_uncons
}

var cache_Data_List_Lazy_NonEmpty_toList gopurs_runtime.Value
var once_Data_List_Lazy_NonEmpty_toList sync.Once

func Get_Data_List_Lazy_NonEmpty_toList() gopurs_runtime.Value {
	once_Data_List_Lazy_NonEmpty_toList.Do(func() {
		cache_Data_List_Lazy_NonEmpty_toList = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_Lazy_NonEmpty_toList(v_0_box)
		})
	})
	return cache_Data_List_Lazy_NonEmpty_toList
}

var cache_Data_List_Lazy_NonEmpty_toUnfoldable gopurs_runtime.Value
var once_Data_List_Lazy_NonEmpty_toUnfoldable sync.Once

func Get_Data_List_Lazy_NonEmpty_toUnfoldable() gopurs_runtime.Value {
	once_Data_List_Lazy_NonEmpty_toUnfoldable.Do(func() {
		cache_Data_List_Lazy_NonEmpty_toUnfoldable = gopurs_runtime.Func(func(dictUnfoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_Lazy_NonEmpty_toUnfoldable(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]](dictUnfoldable_0_box))
		})
	})
	return cache_Data_List_Lazy_NonEmpty_toUnfoldable
}

var cache_Data_List_Lazy_NonEmpty_tail gopurs_runtime.Value
var once_Data_List_Lazy_NonEmpty_tail sync.Once

func Get_Data_List_Lazy_NonEmpty_tail() gopurs_runtime.Value {
	once_Data_List_Lazy_NonEmpty_tail.Do(func() {
		cache_Data_List_Lazy_NonEmpty_tail = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_Lazy_NonEmpty_tail(v_0_box)
		})
	})
	return cache_Data_List_Lazy_NonEmpty_tail
}

var cache_Data_List_Lazy_NonEmpty_singleton gopurs_runtime.Value
var once_Data_List_Lazy_NonEmpty_singleton sync.Once

func Get_Data_List_Lazy_NonEmpty_singleton() gopurs_runtime.Value {
	once_Data_List_Lazy_NonEmpty_singleton.Do(func() {
		cache_Data_List_Lazy_NonEmpty_singleton = gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_applicativeNonEmptyList()).V1
	})
	return cache_Data_List_Lazy_NonEmpty_singleton
}

var cache_Data_List_Lazy_NonEmpty_repeat gopurs_runtime.Value
var once_Data_List_Lazy_NonEmpty_repeat sync.Once

func Get_Data_List_Lazy_NonEmpty_repeat() gopurs_runtime.Value {
	once_Data_List_Lazy_NonEmpty_repeat.Do(func() {
		cache_Data_List_Lazy_NonEmpty_repeat = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_Lazy_NonEmpty_repeat(x_0_box)
		})
	})
	return cache_Data_List_Lazy_NonEmpty_repeat
}

var cache_Data_List_Lazy_NonEmpty_length gopurs_runtime.Value
var once_Data_List_Lazy_NonEmpty_length sync.Once

func Get_Data_List_Lazy_NonEmpty_length() gopurs_runtime.Value {
	once_Data_List_Lazy_NonEmpty_length.Do(func() {
		cache_Data_List_Lazy_NonEmpty_length = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Data_List_Lazy_NonEmpty_length(v_0_box))
		})
	})
	return cache_Data_List_Lazy_NonEmpty_length
}

var cache_Data_List_Lazy_NonEmpty_last gopurs_runtime.Value
var once_Data_List_Lazy_NonEmpty_last sync.Once

func Get_Data_List_Lazy_NonEmpty_last() gopurs_runtime.Value {
	once_Data_List_Lazy_NonEmpty_last.Do(func() {
		cache_Data_List_Lazy_NonEmpty_last = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_Lazy_NonEmpty_last(v_0_box)
		})
	})
	return cache_Data_List_Lazy_NonEmpty_last
}

var cache_Data_List_Lazy_NonEmpty_iterate gopurs_runtime.Value
var once_Data_List_Lazy_NonEmpty_iterate sync.Once

func Get_Data_List_Lazy_NonEmpty_iterate() gopurs_runtime.Value {
	once_Data_List_Lazy_NonEmpty_iterate.Do(func() {
		cache_Data_List_Lazy_NonEmpty_iterate = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_Lazy_NonEmpty_iterate(f_0_box, x_1_box)
		})
	})
	return cache_Data_List_Lazy_NonEmpty_iterate
}

var cache_Data_List_Lazy_NonEmpty_go__init gopurs_runtime.Value
var once_Data_List_Lazy_NonEmpty_go__init sync.Once

func Get_Data_List_Lazy_NonEmpty_go__init() gopurs_runtime.Value {
	once_Data_List_Lazy_NonEmpty_go__init.Do(func() {
		cache_Data_List_Lazy_NonEmpty_go__init = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_Lazy_NonEmpty_go__init(v_0_box)
		})
	})
	return cache_Data_List_Lazy_NonEmpty_go__init
}

var cache_Data_List_Lazy_NonEmpty_head gopurs_runtime.Value
var once_Data_List_Lazy_NonEmpty_head sync.Once

func Get_Data_List_Lazy_NonEmpty_head() gopurs_runtime.Value {
	once_Data_List_Lazy_NonEmpty_head.Do(func() {
		cache_Data_List_Lazy_NonEmpty_head = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_Lazy_NonEmpty_head(v_0_box)
		})
	})
	return cache_Data_List_Lazy_NonEmpty_head
}

var cache_Data_List_Lazy_NonEmpty_fromList gopurs_runtime.Value
var once_Data_List_Lazy_NonEmpty_fromList sync.Once

func Get_Data_List_Lazy_NonEmpty_fromList() gopurs_runtime.Value {
	once_Data_List_Lazy_NonEmpty_fromList.Do(func() {
		cache_Data_List_Lazy_NonEmpty_fromList = gopurs_runtime.Func(func(l_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_List_Lazy_NonEmpty_fromList(l_0_box)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_List_Lazy_NonEmpty_fromList
}

var cache_Data_List_Lazy_NonEmpty_fromFoldable gopurs_runtime.Value
var once_Data_List_Lazy_NonEmpty_fromFoldable sync.Once

func Get_Data_List_Lazy_NonEmpty_fromFoldable() gopurs_runtime.Value {
	once_Data_List_Lazy_NonEmpty_fromFoldable.Do(func() {
		cache_Data_List_Lazy_NonEmpty_fromFoldable = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_Lazy_NonEmpty_fromFoldable(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_0_box))
		})
	})
	return cache_Data_List_Lazy_NonEmpty_fromFoldable
}

var cache_Data_List_Lazy_NonEmpty_cons gopurs_runtime.Value
var once_Data_List_Lazy_NonEmpty_cons sync.Once

func Get_Data_List_Lazy_NonEmpty_cons() gopurs_runtime.Value {
	once_Data_List_Lazy_NonEmpty_cons.Do(func() {
		cache_Data_List_Lazy_NonEmpty_cons = gopurs_runtime.Func2(func(y_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_Lazy_NonEmpty_cons(y_0_box, v_1_box)
		})
	})
	return cache_Data_List_Lazy_NonEmpty_cons
}

var cache_Data_List_Lazy_NonEmpty_concatMap gopurs_runtime.Value
var once_Data_List_Lazy_NonEmpty_concatMap sync.Once

func Get_Data_List_Lazy_NonEmpty_concatMap() gopurs_runtime.Value {
	once_Data_List_Lazy_NonEmpty_concatMap.Do(func() {
		cache_Data_List_Lazy_NonEmpty_concatMap = gopurs_runtime.Func2(func(b_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_Lazy_NonEmpty_concatMap(b_0_box, a_1_box)
		})
	})
	return cache_Data_List_Lazy_NonEmpty_concatMap
}

var cache_Data_List_Lazy_NonEmpty_appendFoldable gopurs_runtime.Value
var once_Data_List_Lazy_NonEmpty_appendFoldable sync.Once

func Get_Data_List_Lazy_NonEmpty_appendFoldable() gopurs_runtime.Value {
	once_Data_List_Lazy_NonEmpty_appendFoldable.Do(func() {
		cache_Data_List_Lazy_NonEmpty_appendFoldable = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_Lazy_NonEmpty_appendFoldable(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_0_box))
		})
	})
	return cache_Data_List_Lazy_NonEmpty_appendFoldable
}

func Call_Data_List_Lazy_NonEmpty_uncons(v_0_loop gopurs_runtime.Value) struct {
	head gopurs_runtime.Value
	tail gopurs_runtime.Value
	/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
} {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	// TAST (Let): v1_1_0 shape=App(Var) bindingType=(ADT ["Data","NonEmpty","NonEmpty"] [(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])]), (TypeVar a$scope1)])
	v1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), v_0))
	_ = v1_1_0
	return struct {
		head gopurs_runtime.Value
		tail gopurs_runtime.Value
		/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
	}{(v1_1_0).V0, (v1_1_0).V1}
}

func Call_Data_List_Lazy_NonEmpty_toList(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	// TAST (Let): v1_1_0 shape=App(Var) bindingType=(ADT ["Data","NonEmpty","NonEmpty"] [(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])]), (TypeVar a$scope4)])
	v1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), v_0))
	_ = v1_1_0
	// TAST (Let): __local_var_2_1 shape=Other bindingType=Any
	__local_var_2_1 := (v1_1_0).V0
	_ = __local_var_2_1
	// TAST (Let): __local_var_3_2 shape=Other bindingType=Any
	__local_var_3_2 := (v1_1_0).V1
	_ = __local_var_3_2
	return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, __local_var_2_1, __local_var_3_2}))}
	}))
}

func Call_Data_List_Lazy_NonEmpty_toUnfoldable(dictUnfoldable_0_loop *Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictUnfoldable_0 *Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value] = dictUnfoldable_0_loop
	_ = dictUnfoldable_0
	return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(dictUnfoldable_0.V1, gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(Record (Row [head: (TypeVar a$scope9), tail: (ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope9)])])] Empty))])
		__local_var_2_0 := Rebox_Data_List_Lazy_NonEmpty_3094389156_3833657837(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			_v := Call_Data_List_Lazy_uncons(xs_1)
			if _v.V1 {
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: func() gopurs_runtime.Value {
					orig := _v.V0
					_ = orig
					return gopurs_runtime.RecordDict2("head", "tail", orig.head, orig.tail)
				}()})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
		}()))
		_ = __local_var_2_0
		var __t1 struct {
			V0 gopurs_runtime.Value
			V1 bool
		}
		{
			if __local_var_2_0 != nil {
				__t1 = struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{func() gopurs_runtime.Value {
					_v := struct {
						V0 gopurs_runtime.Value
						V1 gopurs_runtime.Value
					}{(__local_var_2_0).V0.head, (__local_var_2_0).V0.tail}
					return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
				}(), true}
				goto end_branch_1
			} else {

			}
		}
		{
			__t1 = struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{gopurs_runtime.Value{}, false}
		}
	end_branch_1:
		return func() gopurs_runtime.Value {
			_v := __t1
			if _v.V1 {
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
		}()
	})), Get_Data_List_Lazy_NonEmpty_toList())
}

func Call_Data_List_Lazy_NonEmpty_tail(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	return (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply(Get_Data_Lazy_force(), v_0).UnsafePtr).V1
}

func Call_Data_List_Lazy_NonEmpty_repeat(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer((&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, x_0, Call_Data_List_Lazy_repeat(x_0)}))}
	}))
}

func Call_Data_List_Lazy_NonEmpty_length(v_0_loop gopurs_runtime.Value) int64 {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	return gopurs_runtime.IntAdd(int64(1), gopurs_runtime.Apply(Get_Data_List_Lazy_length(), (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply(Get_Data_Lazy_force(), v_0).UnsafePtr).V1).IntVal)
}

func Call_Data_List_Lazy_NonEmpty_last(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	// TAST (Let): v1_1_0 shape=App(Var) bindingType=(ADT ["Data","NonEmpty","NonEmpty"] [(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])]), (TypeVar a$scope28)])
	v1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), v_0))
	_ = v1_1_0
	// TAST (Let): __local_var_2_1 shape=App(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar a$scope28)])
	__local_var_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Lazy_last(), (v1_1_0).V1))
	_ = __local_var_2_1
	var __t2 gopurs_runtime.Value
	{
		if __local_var_2_1 == nil {
			__t2 = (v1_1_0).V0
			goto end_branch_2
		} else {

		}
	}
	{
		if __local_var_2_1 != nil {
			__t2 = (__local_var_2_1).V0
			goto end_branch_2
		} else {

		}
	}
	{
		__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
	}
end_branch_2:
	return __t2
}

func Call_Data_List_Lazy_NonEmpty_iterate(f_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	var x_1 gopurs_runtime.Value = x_1_loop
	_ = x_1
	return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer((&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, x_1, Call_Data_List_Lazy_iterate(f_0, gopurs_runtime.Apply(f_0, x_1))}))}
	}))
}

func Call_Data_List_Lazy_NonEmpty_go__init(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	// TAST (Let): v1_1_0 shape=App(Var) bindingType=(ADT ["Data","NonEmpty","NonEmpty"] [(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])]), (TypeVar a$scope37)])
	v1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), v_0))
	_ = v1_1_0
	// TAST (Let): __local_var_2_1 shape=App(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope37)])])])
	__local_var_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Lazy_go__init(), (v1_1_0).V1))
	_ = __local_var_2_1
	var __t3 gopurs_runtime.Value
	{
		if __local_var_2_1 == nil {
			__t3 = Get_Data_List_Lazy_Types_nil()
			goto end_branch_3
		} else {

		}
	}
	{
		if __local_var_2_1 != nil {
			// TAST (Let): __local_var_3_2 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope37)])])
			__local_var_3_2 := (__local_var_2_1).V0
			_ = __local_var_3_2
			__t3 = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (v1_1_0).V0, __local_var_3_2}))}
			}))
			goto end_branch_3
		} else {

		}
	}
	{
		__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
	}
end_branch_3:
	return __t3
}

func Call_Data_List_Lazy_NonEmpty_head(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	return (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply(Get_Data_Lazy_force(), v_0).UnsafePtr).V0
}

func Call_Data_List_Lazy_NonEmpty_fromList(l_0_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
	var l_0 gopurs_runtime.Value = l_0_loop
	_ = l_0
	// TAST (Let): v_1_0 shape=App(Var) bindingType=(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope43)])
	v_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), l_0))
	_ = v_1_0
	var __t3 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
	{
		if v_1_0 == nil {
			__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
			goto end_branch_3
		} else {

		}
	}
	{
		if v_1_0 != nil {
			// TAST (Let): __local_var_2_1 shape=Other bindingType=Any
			__local_var_2_1 := (v_1_0).V0
			_ = __local_var_2_1
			// TAST (Let): __local_var_3_2 shape=Other bindingType=Any
			__local_var_3_2 := (v_1_0).V1
			_ = __local_var_3_2
			__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer((&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, __local_var_2_1, __local_var_3_2}))}
				})), true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
			goto end_branch_3
		} else {

		}
	}
	{
		__t3 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
	}
end_branch_3:
	return func() struct {
		V0 gopurs_runtime.Value
		V1 bool
	} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t3)}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
		return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
	}; return struct {
		V0 gopurs_runtime.Value
		V1 bool
	}{V0: gopurs_runtime.Value{}, V1: false} }()
}

func Call_Data_List_Lazy_NonEmpty_fromFoldable(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictFoldable_0 *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
	_ = dictFoldable_0
	return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_List_Lazy_NonEmpty_fromList(), Call_Data_List_Lazy_fromFoldable(dictFoldable_0))
}

func Call_Data_List_Lazy_NonEmpty_cons(y_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var y_0 gopurs_runtime.Value = y_0_loop
	_ = y_0
	var v_1 gopurs_runtime.Value = v_1_loop
	_ = v_1
	return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): v2_3_0 shape=App(Var) bindingType=(ADT ["Data","NonEmpty","NonEmpty"] [(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])]), (TypeVar a$scope56)])
		v2_3_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), v_1))
		_ = v2_3_0
		// TAST (Let): __local_var_4_1 shape=Other bindingType=Any
		__local_var_4_1 := (v2_3_0).V0
		_ = __local_var_4_1
		// TAST (Let): __local_var_5_2 shape=Other bindingType=Any
		__local_var_5_2 := (v2_3_0).V1
		_ = __local_var_5_2
		return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer((&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, y_0, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, __local_var_4_1, __local_var_5_2}))}
		}))}))}
	}))
}

func Call_Data_List_Lazy_NonEmpty_concatMap(b_0_loop gopurs_runtime.Value, a_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var b_0 gopurs_runtime.Value = b_0_loop
	_ = b_0
	var a_1 gopurs_runtime.Value = a_1_loop
	_ = a_1
	return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_bindNonEmptyList()).V1, a_1, b_0)
}

func Call_Data_List_Lazy_NonEmpty_appendFoldable(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictFoldable_0 *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
	_ = dictFoldable_0
	// TAST (Let): fromFoldable1__3634910753_1_0 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar t$scope68) [(TypeVar a$scope69)])] (ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope69)])]))
	fromFoldable1__3634910753_1_0 := gopurs_runtime.Apply2(dictFoldable_0.V2, Get_Data_List_Lazy_Types_cons(), Get_Data_List_Lazy_Types_nil())
	_ = fromFoldable1__3634910753_1_0
	return gopurs_runtime.Func2(func(nel_2 gopurs_runtime.Value, ys_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer((&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply(Get_Data_Lazy_force(), nel_2).UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_semigroupList(), "append"), (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply(Get_Data_Lazy_force(), nel_2).UnsafePtr).V1, gopurs_runtime.Apply(fromFoldable1__3634910753_1_0, ys_3))}))}
		}))
	})
}

func Rebox_Data_List_Lazy_NonEmpty_2812149806_3689823567(in *Constructor_Data_Functor_Functor[gopurs_runtime.Value]) *Constructor_Data_Functor_Functor[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Functor_Functor[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Data_List_Lazy_NonEmpty_3094389156_3833657837(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[struct {
	head gopurs_runtime.Value
	tail gopurs_runtime.Value
	/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
}] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[struct {
		head gopurs_runtime.Value
		tail gopurs_runtime.Value
		/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
	}]{}
	out.V0 = func() struct {
		head gopurs_runtime.Value
		tail gopurs_runtime.Value
		/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
	} {
		orig := in.V0
		_ = orig
		clone := struct {
			head gopurs_runtime.Value
			tail gopurs_runtime.Value
			/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
		}{}
		clone.head = gopurs_runtime.RecordGet(orig, "head")
		clone.tail = gopurs_runtime.RecordGet(orig, "tail")
		return clone
	}()
	return out
}
