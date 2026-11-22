package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Unfoldable1_fromJust gopurs_runtime.Value
var once_Data_Unfoldable1_fromJust sync.Once

func Get_Data_Unfoldable1_fromJust() gopurs_runtime.Value {
	once_Data_Unfoldable1_fromJust.Do(func() {
		cache_Data_Unfoldable1_fromJust = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Unfoldable1_fromJust(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_0_box))
		})
	})
	return cache_Data_Unfoldable1_fromJust
}

var cache_Data_Unfoldable1_lessThanOrEq gopurs_runtime.Value
var once_Data_Unfoldable1_lessThanOrEq sync.Once

func Get_Data_Unfoldable1_lessThanOrEq() gopurs_runtime.Value {
	once_Data_Unfoldable1_lessThanOrEq.Do(func() {
		cache_Data_Unfoldable1_lessThanOrEq = Get_Data_Ord_lessThanOrEq__3348776441()
	})
	return cache_Data_Unfoldable1_lessThanOrEq
}

var cache_Data_Unfoldable1_sub gopurs_runtime.Value
var once_Data_Unfoldable1_sub sync.Once

func Get_Data_Unfoldable1_sub() gopurs_runtime.Value {
	once_Data_Unfoldable1_sub.Do(func() {
		cache_Data_Unfoldable1_sub = Get_Data_Ring_sub__4157804846()
	})
	return cache_Data_Unfoldable1_sub
}

var cache_Data_Unfoldable1_add gopurs_runtime.Value
var once_Data_Unfoldable1_add sync.Once

func Get_Data_Unfoldable1_add() gopurs_runtime.Value {
	once_Data_Unfoldable1_add.Do(func() {
		cache_Data_Unfoldable1_add = Get_Data_Semiring_add__4086903470()
	})
	return cache_Data_Unfoldable1_add
}

var cache_Data_Unfoldable1_eq gopurs_runtime.Value
var once_Data_Unfoldable1_eq sync.Once

func Get_Data_Unfoldable1_eq() gopurs_runtime.Value {
	once_Data_Unfoldable1_eq.Do(func() {
		cache_Data_Unfoldable1_eq = Get_Data_Eq_eq__3020190905()
	})
	return cache_Data_Unfoldable1_eq
}

var cache_Data_Unfoldable1_greaterThanOrEq gopurs_runtime.Value
var once_Data_Unfoldable1_greaterThanOrEq sync.Once

func Get_Data_Unfoldable1_greaterThanOrEq() gopurs_runtime.Value {
	once_Data_Unfoldable1_greaterThanOrEq.Do(func() {
		cache_Data_Unfoldable1_greaterThanOrEq = Get_Data_Ord_greaterThanOrEq__3348776441()
	})
	return cache_Data_Unfoldable1_greaterThanOrEq
}

var cache_Data_Unfoldable1_negate gopurs_runtime.Value
var once_Data_Unfoldable1_negate sync.Once

func Get_Data_Unfoldable1_negate() gopurs_runtime.Value {
	once_Data_Unfoldable1_negate.Do(func() {
		cache_Data_Unfoldable1_negate = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Data_Unfoldable1_negate(a_0_box.IntVal))
		})
	})
	return cache_Data_Unfoldable1_negate
}

var cache_Data_Unfoldable1_greaterThan gopurs_runtime.Value
var once_Data_Unfoldable1_greaterThan sync.Once

func Get_Data_Unfoldable1_greaterThan() gopurs_runtime.Value {
	once_Data_Unfoldable1_greaterThan.Do(func() {
		cache_Data_Unfoldable1_greaterThan = Get_Data_Ord_greaterThan__3348776441()
	})
	return cache_Data_Unfoldable1_greaterThan
}

var cache_Data_Unfoldable1_Unfoldable1_dollar_Dict gopurs_runtime.Value
var once_Data_Unfoldable1_Unfoldable1_dollar_Dict sync.Once

func Get_Data_Unfoldable1_Unfoldable1_dollar_Dict() gopurs_runtime.Value {
	once_Data_Unfoldable1_Unfoldable1_dollar_Dict.Do(func() {
		cache_Data_Unfoldable1_Unfoldable1_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3553002490, UnsafePtr: unsafe.Pointer(Call_Data_Unfoldable1_Unfoldable1_dollar_Dict(func() struct {
				unfoldr1 gopurs_runtime.Value
				/*gopurs:record-v1:1:0075006e0066006f006c006400720031*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					unfoldr1 gopurs_runtime.Value
					/*gopurs:record-v1:1:0075006e0066006f006c006400720031*/
				}{}
				clone.unfoldr1 = gopurs_runtime.RecordGet(orig, "unfoldr1")
				return clone
			}()))}
		})
	})
	return cache_Data_Unfoldable1_Unfoldable1_dollar_Dict
}

var cache_Data_Unfoldable1_unfoldr1 gopurs_runtime.Value
var once_Data_Unfoldable1_unfoldr1 sync.Once

func Get_Data_Unfoldable1_unfoldr1() gopurs_runtime.Value {
	once_Data_Unfoldable1_unfoldr1.Do(func() {
		cache_Data_Unfoldable1_unfoldr1 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Unfoldable1_unfoldr1(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Data_Unfoldable1_unfoldr1
}

var cache_Data_Unfoldable1_unfoldable1Maybe gopurs_runtime.Value
var once_Data_Unfoldable1_unfoldable1Maybe sync.Once

func Get_Data_Unfoldable1_unfoldable1Maybe() gopurs_runtime.Value {
	once_Data_Unfoldable1_unfoldable1Maybe.Do(func() {
		cache_Data_Unfoldable1_unfoldable1Maybe = gopurs_runtime.Value{Type: 9, IntVal: 3553002490, UnsafePtr: unsafe.Pointer(Rebox_Data_Unfoldable1_4212083983_2187088110((&Constructor_Data_Unfoldable1_Unfoldable1[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, b_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{(*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply(f_0, b_1).UnsafePtr).V0, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})})))}
	})
	return cache_Data_Unfoldable1_unfoldable1Maybe
}

var cache_Data_Unfoldable1_unfoldable1Array gopurs_runtime.Value
var once_Data_Unfoldable1_unfoldable1Array sync.Once

func Get_Data_Unfoldable1_unfoldable1Array() gopurs_runtime.Value {
	once_Data_Unfoldable1_unfoldable1Array.Do(func() {
		cache_Data_Unfoldable1_unfoldable1Array = gopurs_runtime.Value{Type: 9, IntVal: 3553002490, UnsafePtr: unsafe.Pointer((&Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value]{1, gopurs_runtime.Apply4(Get_Data_Unfoldable1_unfoldr1ArrayImpl(), Get_Data_Maybe_isNothing(), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t0 gopurs_runtime.Value
			{
				if v_0.Type == 9 && v_0.IntVal == 930809136 && v_0.UnsafePtr != nil {
					__t0 = (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v_0.UnsafePtr).V0
					goto end_branch_0
				} else {

				}
			}
			{
				__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_0:
			return __t0
		}), Get_Data_Tuple_fst(), Get_Data_Tuple_snd())}))}
	})
	return cache_Data_Unfoldable1_unfoldable1Array
}

var cache_Data_Unfoldable1_replicate1 gopurs_runtime.Value
var once_Data_Unfoldable1_replicate1 sync.Once

func Get_Data_Unfoldable1_replicate1() gopurs_runtime.Value {
	once_Data_Unfoldable1_replicate1.Do(func() {
		cache_Data_Unfoldable1_replicate1 = gopurs_runtime.Func3(func(dictUnfoldable1_0_box gopurs_runtime.Value, n_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Unfoldable1_replicate1(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value]](dictUnfoldable1_0_box), n_1_box.IntVal, v_2_box)
		})
	})
	return cache_Data_Unfoldable1_replicate1
}

var cache_Data_Unfoldable1_replicate1A gopurs_runtime.Value
var once_Data_Unfoldable1_replicate1A sync.Once

func Get_Data_Unfoldable1_replicate1A() gopurs_runtime.Value {
	once_Data_Unfoldable1_replicate1A.Do(func() {
		cache_Data_Unfoldable1_replicate1A = gopurs_runtime.Func3(func(dictApply_0_box gopurs_runtime.Value, dictUnfoldable1_1_box gopurs_runtime.Value, dictTraversable1_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Unfoldable1_replicate1A(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](dictApply_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value]](dictUnfoldable1_1_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Traversable_Traversable1[gopurs_runtime.Value]](dictTraversable1_2_box))
		})
	})
	return cache_Data_Unfoldable1_replicate1A
}

var cache_Data_Unfoldable1_singleton gopurs_runtime.Value
var once_Data_Unfoldable1_singleton sync.Once

func Get_Data_Unfoldable1_singleton() gopurs_runtime.Value {
	once_Data_Unfoldable1_singleton.Do(func() {
		cache_Data_Unfoldable1_singleton = gopurs_runtime.Func2(func(dictUnfoldable1_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Unfoldable1_singleton(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value]](dictUnfoldable1_0_box), v_1_box)
		})
	})
	return cache_Data_Unfoldable1_singleton
}

var cache_Data_Unfoldable1_go__range gopurs_runtime.Value
var once_Data_Unfoldable1_go__range sync.Once

func Get_Data_Unfoldable1_go__range() gopurs_runtime.Value {
	once_Data_Unfoldable1_go__range.Do(func() {
		cache_Data_Unfoldable1_go__range = gopurs_runtime.Func3(func(dictUnfoldable1_0_box gopurs_runtime.Value, start_1_box gopurs_runtime.Value, end_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Unfoldable1_go__range(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value]](dictUnfoldable1_0_box), start_1_box.IntVal, end_2_box.IntVal)
		})
	})
	return cache_Data_Unfoldable1_go__range
}

var cache_Data_Unfoldable1_iterateN gopurs_runtime.Value
var once_Data_Unfoldable1_iterateN sync.Once

func Get_Data_Unfoldable1_iterateN() gopurs_runtime.Value {
	once_Data_Unfoldable1_iterateN.Do(func() {
		cache_Data_Unfoldable1_iterateN = gopurs_runtime.Func4(func(dictUnfoldable1_0_box gopurs_runtime.Value, n_1_box gopurs_runtime.Value, f_2_box gopurs_runtime.Value, s_3_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Unfoldable1_iterateN(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value]](dictUnfoldable1_0_box), n_1_box.IntVal, f_2_box, s_3_box)
		})
	})
	return cache_Data_Unfoldable1_iterateN
}

type Constructor_Data_Unfoldable1_Unfoldable1[T_t any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

//gopurs:init-v1:00670065006e006500720061007400650064002f0044006100740061005f0055006e0066006f006c006400610062006c00650031:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f0044006100740061005f0055006e0066006f006c006400610062006c00650031005f0055006e0066006f006c006400610062006c00650031
func init() {
	gopurs_runtime.StructGetters[3553002490] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "unfoldr1":
			return gopurs_runtime.Box(c.V0)
		default:
			panic("Key not found in dictionary Constructor_Data_Unfoldable1_Unfoldable1: " + key)
		}
	}
}

func Call_Data_Unfoldable1_fromJust(v_0_loop *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
	var v_0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = v_0_loop
	_ = v_0
	var __t0 gopurs_runtime.Value
	{
		if v_0 != nil {
			__t0 = (v_0).V0
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
	}
end_branch_0:
	return __t0
}

func Call_Data_Unfoldable1_negate(a_0_loop int64) int64 {
	var a_0 int64 = a_0_loop
	_ = a_0
	return gopurs_runtime.IntNegate(a_0)
}

func Call_Data_Unfoldable1_Unfoldable1_dollar_Dict(x_0_loop struct {
	unfoldr1 gopurs_runtime.Value
	/*gopurs:record-v1:1:0075006e0066006f006c006400720031*/
}) *Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value] {
	var x_0 struct {
		unfoldr1 gopurs_runtime.Value
		/*gopurs:record-v1:1:0075006e0066006f006c006400720031*/
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict1("unfoldr1", orig.unfoldr1)
	}())
}

func Call_Data_Unfoldable1_unfoldr1(dict_0_loop *Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V0
}

func Call_Data_Unfoldable1_replicate1(dictUnfoldable1_0_loop *Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value], n_1_loop int64, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictUnfoldable1_0 *Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value] = dictUnfoldable1_0_loop
	_ = dictUnfoldable1_0
	var n_1 int64 = n_1_loop
	_ = n_1
	var v_2 gopurs_runtime.Value = v_2_loop
	_ = v_2
	return gopurs_runtime.Apply2(dictUnfoldable1_0.V0, gopurs_runtime.Func(func(i_3 gopurs_runtime.Value) gopurs_runtime.Value {
		var __t0 struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
		}
		{
			if (i_3.IntVal) <= (int64(0)) {
				__t0 = struct {
					V0 gopurs_runtime.Value
					V1 gopurs_runtime.Value
				}{v_2, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}}
				goto end_branch_0
			} else {

			}
		}
		{
			__t0 = struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
			}{v_2, func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Int(gopurs_runtime.IntSub(i_3.IntVal, int64(1))), true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()}
		}
	end_branch_0:
		return func() gopurs_runtime.Value {
			_v := __t0
			return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
		}()
	}), gopurs_runtime.Int(gopurs_runtime.IntSub(n_1, int64(1))))
}

func Call_Data_Unfoldable1_replicate1A(dictApply_0_loop *Constructor_Control_Apply_Apply[gopurs_runtime.Value], dictUnfoldable1_1_loop *Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value], dictTraversable1_2_loop *Constructor_Data_Semigroup_Traversable_Traversable1[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictApply_0 *Constructor_Control_Apply_Apply[gopurs_runtime.Value] = dictApply_0_loop
	_ = dictApply_0
	var dictUnfoldable1_1 *Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value] = dictUnfoldable1_1_loop
	_ = dictUnfoldable1_1
	var dictTraversable1_2 *Constructor_Data_Semigroup_Traversable_Traversable1[gopurs_runtime.Value] = dictTraversable1_2_loop
	_ = dictTraversable1_2
	// TAST (Let): sequence1_3_0 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f$scope42) [(TypeApp (TypeVar m$scope41) [(TypeVar a$scope43)])])] (TypeApp (TypeVar m$scope41) [(TypeApp (TypeVar f$scope42) [(TypeVar a$scope43)])]))
	sequence1_3_0 := gopurs_runtime.Apply(dictTraversable1_2.V2, gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(dictApply_0)})
	_ = sequence1_3_0
	return gopurs_runtime.Func2(func(n_4 gopurs_runtime.Value, m_5 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(sequence1_3_0, gopurs_runtime.Apply2(dictUnfoldable1_1.V0, gopurs_runtime.Func(func(i_6 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t1 struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
			}
			{
				if (i_6.IntVal) <= (int64(0)) {
					__t1 = struct {
						V0 gopurs_runtime.Value
						V1 gopurs_runtime.Value
					}{m_5, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}}
					goto end_branch_1
				} else {

				}
			}
			{
				__t1 = struct {
					V0 gopurs_runtime.Value
					V1 gopurs_runtime.Value
				}{m_5, func() gopurs_runtime.Value {
					_v := struct {
						V0 gopurs_runtime.Value
						V1 bool
					}{gopurs_runtime.Int(gopurs_runtime.IntSub(i_6.IntVal, int64(1))), true}
					if _v.V1 {
						return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
					}
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
				}()}
			}
		end_branch_1:
			return func() gopurs_runtime.Value {
				_v := __t1
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()
		}), gopurs_runtime.Int(gopurs_runtime.IntSub(n_4.IntVal, int64(1)))))
	})
}

func Call_Data_Unfoldable1_singleton(dictUnfoldable1_0_loop *Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value], v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictUnfoldable1_0 *Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value] = dictUnfoldable1_0_loop
	_ = dictUnfoldable1_0
	var v_1 gopurs_runtime.Value = v_1_loop
	_ = v_1
	return gopurs_runtime.Apply2(dictUnfoldable1_0.V0, gopurs_runtime.Func(func(i_2 gopurs_runtime.Value) gopurs_runtime.Value {
		var __t0 struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
		}
		{
			if (i_2.IntVal) <= (int64(0)) {
				__t0 = struct {
					V0 gopurs_runtime.Value
					V1 gopurs_runtime.Value
				}{v_1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}}
				goto end_branch_0
			} else {

			}
		}
		{
			__t0 = struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
			}{v_1, func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Int(gopurs_runtime.IntSub(i_2.IntVal, int64(1))), true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()}
		}
	end_branch_0:
		return func() gopurs_runtime.Value {
			_v := __t0
			return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
		}()
	}), gopurs_runtime.Int(int64(0)))
}

func Call_Data_Unfoldable1_go__range(dictUnfoldable1_0_loop *Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value], start_1_loop int64, end_2_loop int64) gopurs_runtime.Value {
	var dictUnfoldable1_0 *Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value] = dictUnfoldable1_0_loop
	_ = dictUnfoldable1_0
	var start_1 int64 = start_1_loop
	_ = start_1
	var end_2 int64 = end_2_loop
	_ = end_2
	var __t1 int64
	{
		if (end_2) >= (start_1) {
			__t1 = int64(1)
			goto end_branch_1
		} else {

		}
	}
	{
		__t1 = int64(-1)
	}
end_branch_1:
	// TAST (Let): __local_var_3_0 shape=Branch(LitInt, def=LitInt) bindingType=Int
	__local_var_3_0 := __t1
	_ = __local_var_3_0
	return gopurs_runtime.Apply2(dictUnfoldable1_0.V0, gopurs_runtime.Func(func(i_4 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): i_prime__5_2 shape=Other bindingType=Int
		i_prime__5_2 := gopurs_runtime.IntAdd(i_4.IntVal, __local_var_3_0)
		_ = i_prime__5_2
		var __t3 gopurs_runtime.Value
		{
			if (i_4.IntVal) == (end_2) {
				__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}
				goto end_branch_3
			} else {

			}
		}
		{
			__t3 = func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Int(i_prime__5_2), true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		}
	end_branch_3:
		return func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
			}{i_4, __t3}
			return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
		}()
	}), gopurs_runtime.Int(start_1))
}

func Call_Data_Unfoldable1_iterateN(dictUnfoldable1_0_loop *Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value], n_1_loop int64, f_2_loop gopurs_runtime.Value, s_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictUnfoldable1_0 *Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value] = dictUnfoldable1_0_loop
	_ = dictUnfoldable1_0
	var n_1 int64 = n_1_loop
	_ = n_1
	var f_2 gopurs_runtime.Value = f_2_loop
	_ = f_2
	var s_3 gopurs_runtime.Value = s_3_loop
	_ = s_3
	return gopurs_runtime.Apply2(dictUnfoldable1_0.V0, gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
		var __t0 gopurs_runtime.Value
		{
			if ((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1.IntVal) > (int64(0)) {
				__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer((&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_2, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0), gopurs_runtime.Int(gopurs_runtime.IntSub((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1.IntVal, int64(1)))}))}}))}
				goto end_branch_0
			} else {

			}
		}
		{
			__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}
		}
	end_branch_0:
		return func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
			}{(*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0, __t0}
			return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
		}()
	}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_Unfoldable1_3415943795_138441832(Rebox_Data_Unfoldable1_138441832_3415943795(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
		_v := struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
		}{s_3, gopurs_runtime.Int(gopurs_runtime.IntSub(n_1, int64(1)))}
		return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
	}()))))})
}

func Rebox_Data_Unfoldable1_138441832_3415943795(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, int64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, int64]{}
	out.V0 = in.V0
	out.V1 = in.V1.IntVal
	return out
}

func Rebox_Data_Unfoldable1_3415943795_138441832(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, int64]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = in.V0
	out.V1 = gopurs_runtime.Int(in.V1)
	return out
}

func Rebox_Data_Unfoldable1_4212083983_2187088110(in *Constructor_Data_Unfoldable1_Unfoldable1[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Get_Data_Unfoldable1_unfoldr1ArrayImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Unfoldable1_Unfoldr1ArrayImpl
}
