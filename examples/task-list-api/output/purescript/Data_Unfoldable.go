package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Unfoldable_go__map gopurs_runtime.Value
var once_Data_Unfoldable_go__map sync.Once

func Get_Data_Unfoldable_go__map() gopurs_runtime.Value {
	once_Data_Unfoldable_go__map.Do(func() {
		cache_Data_Unfoldable_go__map = Rebox_Data_Unfoldable_2812149806_3689823567(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Maybe_functorMaybe())).V0
	})
	return cache_Data_Unfoldable_go__map
}

var cache_Data_Unfoldable_fromJust gopurs_runtime.Value
var once_Data_Unfoldable_fromJust sync.Once

func Get_Data_Unfoldable_fromJust() gopurs_runtime.Value {
	once_Data_Unfoldable_fromJust.Do(func() {
		cache_Data_Unfoldable_fromJust = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Unfoldable_fromJust(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_0_box))
		})
	})
	return cache_Data_Unfoldable_fromJust
}

var cache_Data_Unfoldable_lessThanOrEq gopurs_runtime.Value
var once_Data_Unfoldable_lessThanOrEq sync.Once

func Get_Data_Unfoldable_lessThanOrEq() gopurs_runtime.Value {
	once_Data_Unfoldable_lessThanOrEq.Do(func() {
		cache_Data_Unfoldable_lessThanOrEq = Get_Data_Ord_lessThanOrEq__3348776441()
	})
	return cache_Data_Unfoldable_lessThanOrEq
}

var cache_Data_Unfoldable_sub gopurs_runtime.Value
var once_Data_Unfoldable_sub sync.Once

func Get_Data_Unfoldable_sub() gopurs_runtime.Value {
	once_Data_Unfoldable_sub.Do(func() {
		cache_Data_Unfoldable_sub = Get_Data_Ring_sub__4157804846()
	})
	return cache_Data_Unfoldable_sub
}

var cache_Data_Unfoldable_map1 gopurs_runtime.Value
var once_Data_Unfoldable_map1 sync.Once

func Get_Data_Unfoldable_map1() gopurs_runtime.Value {
	once_Data_Unfoldable_map1.Do(func() {
		cache_Data_Unfoldable_map1 = Rebox_Data_Unfoldable_2812149806_3689823567(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Maybe_functorMaybe())).V0
	})
	return cache_Data_Unfoldable_map1
}

var cache_Data_Unfoldable_Unfoldable_dollar_Dict gopurs_runtime.Value
var once_Data_Unfoldable_Unfoldable_dollar_Dict sync.Once

func Get_Data_Unfoldable_Unfoldable_dollar_Dict() gopurs_runtime.Value {
	once_Data_Unfoldable_Unfoldable_dollar_Dict.Do(func() {
		cache_Data_Unfoldable_Unfoldable_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2670894170, UnsafePtr: unsafe.Pointer(Call_Data_Unfoldable_Unfoldable_dollar_Dict(func() struct {
				Unfoldable10 gopurs_runtime.Value
				unfoldr      gopurs_runtime.Value
				/*gopurs:record-v1:2:0055006e0066006f006c006400610062006c006500310030,0075006e0066006f006c00640072*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Unfoldable10 gopurs_runtime.Value
					unfoldr      gopurs_runtime.Value
					/*gopurs:record-v1:2:0055006e0066006f006c006400610062006c006500310030,0075006e0066006f006c00640072*/
				}{}
				clone.Unfoldable10 = gopurs_runtime.RecordGet(orig, "Unfoldable10")
				clone.unfoldr = gopurs_runtime.RecordGet(orig, "unfoldr")
				return clone
			}()))}
		})
	})
	return cache_Data_Unfoldable_Unfoldable_dollar_Dict
}

var cache_Data_Unfoldable_unfoldr gopurs_runtime.Value
var once_Data_Unfoldable_unfoldr sync.Once

func Get_Data_Unfoldable_unfoldr() gopurs_runtime.Value {
	once_Data_Unfoldable_unfoldr.Do(func() {
		cache_Data_Unfoldable_unfoldr = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Unfoldable_unfoldr(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Data_Unfoldable_unfoldr
}

var cache_Data_Unfoldable_unfoldr__4022650341 gopurs_runtime.Value
var once_Data_Unfoldable_unfoldr__4022650341 sync.Once

func Get_Data_Unfoldable_unfoldr__4022650341() gopurs_runtime.Value {
	once_Data_Unfoldable_unfoldr__4022650341.Do(func() {
		cache_Data_Unfoldable_unfoldr__4022650341 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Data_Unfoldable_2442833393_849153993(Call_Data_Unfoldable_unfoldr__4022650341(__eta0_0_box, __eta1_1_box.IntVal)))}
		})
	})
	return cache_Data_Unfoldable_unfoldr__4022650341
}

var cache_Data_Unfoldable_unfoldableArray gopurs_runtime.Value
var once_Data_Unfoldable_unfoldableArray sync.Once

func Get_Data_Unfoldable_unfoldableArray() gopurs_runtime.Value {
	once_Data_Unfoldable_unfoldableArray.Do(func() {
		cache_Data_Unfoldable_unfoldableArray = gopurs_runtime.Value{Type: 9, IntVal: 2670894170, UnsafePtr: unsafe.Pointer((&Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3553002490, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value]](Get_Data_Unfoldable1_unfoldable1Array()))}
		}), gopurs_runtime.Apply4(Get_Data_Unfoldable_unfoldrArrayImpl(), Get_Data_Maybe_isNothing(), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
	return cache_Data_Unfoldable_unfoldableArray
}

var cache_Data_Unfoldable_unfoldr__2862945647 gopurs_runtime.Value
var once_Data_Unfoldable_unfoldr__2862945647 sync.Once

func Get_Data_Unfoldable_unfoldr__2862945647() gopurs_runtime.Value {
	once_Data_Unfoldable_unfoldr__2862945647.Do(func() {
		cache_Data_Unfoldable_unfoldr__2862945647 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Unfoldable_unfoldr__2862945647(__eta0_0_box, __eta1_1_box.StrVal())
		})
	})
	return cache_Data_Unfoldable_unfoldr__2862945647
}

var cache_Data_Unfoldable_unfoldableMaybe gopurs_runtime.Value
var once_Data_Unfoldable_unfoldableMaybe sync.Once

func Get_Data_Unfoldable_unfoldableMaybe() gopurs_runtime.Value {
	once_Data_Unfoldable_unfoldableMaybe.Do(func() {
		cache_Data_Unfoldable_unfoldableMaybe = gopurs_runtime.Value{Type: 9, IntVal: 2670894170, UnsafePtr: unsafe.Pointer(Rebox_Data_Unfoldable_1401495535_2738507278((&Constructor_Data_Unfoldable_Unfoldable[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3553002490, UnsafePtr: unsafe.Pointer(Rebox_Data_Unfoldable_4212083983_2187088110(Rebox_Data_Unfoldable_2187088110_4212083983(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value]](Get_Data_Unfoldable1_unfoldable1Maybe()))))}
		}), gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, b_1 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_2_0 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(ADT ["Data","Tuple","Tuple"] [(TypeVar a$scope6), (TypeVar b$scope7)])])
			__local_var_2_0 := Rebox_Data_Unfoldable_3094389156_4010058633(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(f_0, b_1)))
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
					}{Call_Data_Tuple_fst((__local_var_2_0).V0), true}
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
		})})))}
	})
	return cache_Data_Unfoldable_unfoldableMaybe
}

var cache_Data_Unfoldable_replicate gopurs_runtime.Value
var once_Data_Unfoldable_replicate sync.Once

func Get_Data_Unfoldable_replicate() gopurs_runtime.Value {
	once_Data_Unfoldable_replicate.Do(func() {
		cache_Data_Unfoldable_replicate = gopurs_runtime.Func3(func(dictUnfoldable_0_box gopurs_runtime.Value, n_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Unfoldable_replicate(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]](dictUnfoldable_0_box), n_1_box.IntVal, v_2_box)
		})
	})
	return cache_Data_Unfoldable_replicate
}

var cache_Data_Unfoldable_replicateA gopurs_runtime.Value
var once_Data_Unfoldable_replicateA sync.Once

func Get_Data_Unfoldable_replicateA() gopurs_runtime.Value {
	once_Data_Unfoldable_replicateA.Do(func() {
		cache_Data_Unfoldable_replicateA = gopurs_runtime.Func3(func(dictApplicative_0_box gopurs_runtime.Value, dictUnfoldable_1_box gopurs_runtime.Value, dictTraversable_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Unfoldable_replicateA(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]](dictUnfoldable_1_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](dictTraversable_2_box))
		})
	})
	return cache_Data_Unfoldable_replicateA
}

var cache_Data_Unfoldable_none gopurs_runtime.Value
var once_Data_Unfoldable_none sync.Once

func Get_Data_Unfoldable_none() gopurs_runtime.Value {
	once_Data_Unfoldable_none.Do(func() {
		cache_Data_Unfoldable_none = gopurs_runtime.Func(func(dictUnfoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Unfoldable_none(dictUnfoldable_0_box)
		})
	})
	return cache_Data_Unfoldable_none
}

var cache_Data_Unfoldable_fromMaybe gopurs_runtime.Value
var once_Data_Unfoldable_fromMaybe sync.Once

func Get_Data_Unfoldable_fromMaybe() gopurs_runtime.Value {
	once_Data_Unfoldable_fromMaybe.Do(func() {
		cache_Data_Unfoldable_fromMaybe = gopurs_runtime.Func(func(dictUnfoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Unfoldable_fromMaybe(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]](dictUnfoldable_0_box))
		})
	})
	return cache_Data_Unfoldable_fromMaybe
}

type Constructor_Data_Unfoldable_Unfoldable[T_t any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}

//gopurs:init-v1:00670065006e006500720061007400650064002f0044006100740061005f0055006e0066006f006c006400610062006c0065:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f0044006100740061005f0055006e0066006f006c006400610062006c0065005f0055006e0066006f006c006400610062006c0065
func init() {
	gopurs_runtime.StructGetters[2670894170] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "Unfoldable10":
			return gopurs_runtime.Box(c.V0)
		case "unfoldr":
			return gopurs_runtime.Box(c.V1)
		default:
			panic("Key not found in dictionary Constructor_Data_Unfoldable_Unfoldable: " + key)
		}
	}
}

func Call_Data_Unfoldable_fromJust(v_0_loop *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
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

func Call_Data_Unfoldable_Unfoldable_dollar_Dict(x_0_loop struct {
	Unfoldable10 gopurs_runtime.Value
	unfoldr      gopurs_runtime.Value
	/*gopurs:record-v1:2:0055006e0066006f006c006400610062006c006500310030,0075006e0066006f006c00640072*/
}) *Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value] {
	var x_0 struct {
		Unfoldable10 gopurs_runtime.Value
		unfoldr      gopurs_runtime.Value
		/*gopurs:record-v1:2:0055006e0066006f006c006400610062006c006500310030,0075006e0066006f006c00640072*/
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict2("Unfoldable10", "unfoldr", orig.Unfoldable10, orig.unfoldr)
	}())
}

func Call_Data_Unfoldable_unfoldr(dict_0_loop *Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V1
}

func Call_Data_Unfoldable_unfoldr__4022650341(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop int64) *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]] {
unfoldr__4022650341:
	for {
		if false {
			continue unfoldr__4022650341
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 int64 = __eta1_1_loop
		_ = __eta1_1
		return Rebox_Data_Unfoldable_849153993_2442833393(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(Rebox_Data_Unfoldable_2738507278_132451362(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]](Get_Data_List_Types_unfoldableList())).V1, __eta0_0, gopurs_runtime.Int(__eta1_1))))
	}
}

func Call_Data_Unfoldable_unfoldr__2862945647(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop string) gopurs_runtime.Value {
unfoldr__2862945647:
	for {
		if false {
			continue unfoldr__2862945647
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 string = __eta1_1_loop
		_ = __eta1_1
		return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]](Get_Data_Unfoldable_unfoldableArray()).V1, __eta0_0, gopurs_runtime.Str(__eta1_1))
	}
}

func Call_Data_Unfoldable_replicate(dictUnfoldable_0_loop *Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value], n_1_loop int64, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictUnfoldable_0 *Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value] = dictUnfoldable_0_loop
	_ = dictUnfoldable_0
	var n_1 int64 = n_1_loop
	_ = n_1
	var v_2 gopurs_runtime.Value = v_2_loop
	_ = v_2
	return gopurs_runtime.Apply2(dictUnfoldable_0.V1, gopurs_runtime.Func(func(i_3 gopurs_runtime.Value) gopurs_runtime.Value {
		var __t0 struct {
			V0 gopurs_runtime.Value
			V1 bool
		}
		{
			if (i_3.IntVal) <= (int64(0)) {
				__t0 = struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Value{}, false}
				goto end_branch_0
			} else {

			}
		}
		{
			__t0 = struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer((&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, v_2, gopurs_runtime.Int(gopurs_runtime.IntSub(i_3.IntVal, int64(1)))}))}, true}
		}
	end_branch_0:
		return func() gopurs_runtime.Value {
			_v := __t0
			if _v.V1 {
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
		}()
	}), gopurs_runtime.Int(n_1))
}

func Call_Data_Unfoldable_replicateA(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value], dictUnfoldable_1_loop *Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value], dictTraversable_2_loop *Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictApplicative_0 *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
	_ = dictApplicative_0
	var dictUnfoldable_1 *Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value] = dictUnfoldable_1_loop
	_ = dictUnfoldable_1
	var dictTraversable_2 *Constructor_Data_Traversable_Traversable[gopurs_runtime.Value] = dictTraversable_2_loop
	_ = dictTraversable_2
	// TAST (Let): sequence_3_0 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f$scope39) [(TypeApp (TypeVar m$scope38) [(TypeVar a$scope40)])])] (TypeApp (TypeVar m$scope38) [(TypeApp (TypeVar f$scope39) [(TypeVar a$scope40)])]))
	sequence_3_0 := gopurs_runtime.Apply(dictTraversable_2.V2, gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(dictApplicative_0)})
	_ = sequence_3_0
	return gopurs_runtime.Func2(func(n_4 gopurs_runtime.Value, m_5 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(sequence_3_0, gopurs_runtime.Apply2(dictUnfoldable_1.V1, gopurs_runtime.Func(func(i_6 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t1 struct {
				V0 gopurs_runtime.Value
				V1 bool
			}
			{
				if (i_6.IntVal) <= (int64(0)) {
					__t1 = struct {
						V0 gopurs_runtime.Value
						V1 bool
					}{gopurs_runtime.Value{}, false}
					goto end_branch_1
				} else {

				}
			}
			{
				__t1 = struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer((&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, m_5, gopurs_runtime.Int(gopurs_runtime.IntSub(i_6.IntVal, int64(1)))}))}, true}
			}
		end_branch_1:
			return func() gopurs_runtime.Value {
				_v := __t1
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		}), gopurs_runtime.Int(n_4.IntVal)))
	})
}

func Call_Data_Unfoldable_none(dictUnfoldable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictUnfoldable_0 gopurs_runtime.Value = dictUnfoldable_0_loop
	_ = dictUnfoldable_0
	return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictUnfoldable_0, "unfoldr"), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}
	}), Get_Data_Unit_unit())
}

func Call_Data_Unfoldable_fromMaybe(dictUnfoldable_0_loop *Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictUnfoldable_0 *Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value] = dictUnfoldable_0_loop
	_ = dictUnfoldable_0
	return gopurs_runtime.Apply(dictUnfoldable_0.V1, gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
		var __t0 struct {
			V0 gopurs_runtime.Value
			V1 bool
		}
		{
			if b_1.Type == 9 && b_1.IntVal == 930809136 && b_1.UnsafePtr != nil {
				__t0 = struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer((&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(b_1.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}}))}, true}
				goto end_branch_0
			} else {

			}
		}
		{
			__t0 = struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{gopurs_runtime.Value{}, false}
		}
	end_branch_0:
		return func() gopurs_runtime.Value {
			_v := __t0
			if _v.V1 {
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
		}()
	}))
}

func Rebox_Data_Unfoldable_138441832_3132786365(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[uint32, float64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[uint32, float64]{}
	out.V0 = uint32(in.V0.IntVal)
	out.V1 = in.V1.FloatVal()
	return out
}

func Rebox_Data_Unfoldable_1401495535_2738507278(in *Constructor_Data_Unfoldable_Unfoldable[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Unfoldable_2187088110_4212083983(in *Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value]) *Constructor_Data_Unfoldable1_Unfoldable1[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Unfoldable1_Unfoldable1[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Data_Unfoldable_2442833393_849153993(in *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_Unfoldable_3132786365_138441832(in.V0))}
	out.V1 = Rebox_Data_Unfoldable_2442833393_849153993(in.V1)
	return out
}

func Rebox_Data_Unfoldable_2738507278_132451362(in *Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]) *Constructor_Data_Unfoldable_Unfoldable[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Unfoldable_Unfoldable[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Data_Unfoldable_2812149806_3689823567(in *Constructor_Data_Functor_Functor[gopurs_runtime.Value]) *Constructor_Data_Functor_Functor[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Functor_Functor[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Data_Unfoldable_3094389156_4010058633(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{}
	out.V0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](in.V0)
	return out
}

func Rebox_Data_Unfoldable_3132786365_138441832(in *Constructor_Data_Tuple_Tuple[uint32, float64]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Value{Type: 9, IntVal: int64(in.V0), UnsafePtr: nil}
	out.V1 = gopurs_runtime.Float(in.V1)
	return out
}

func Rebox_Data_Unfoldable_4212083983_2187088110(in *Constructor_Data_Unfoldable1_Unfoldable1[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Unfoldable_849153993_2442833393(in *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]]{}
	out.V0 = Rebox_Data_Unfoldable_138441832_3132786365(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](in.V0))
	out.V1 = Rebox_Data_Unfoldable_849153993_2442833393(in.V1)
	return out
}

func Get_Data_Unfoldable_unfoldrArrayImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Unfoldable_UnfoldrArrayImpl
}
