package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Array_ST_bind gopurs_runtime.Value
var once_Data_Array_ST_bind sync.Once

func Get_Data_Array_ST_bind() gopurs_runtime.Value {
	once_Data_Array_ST_bind.Do(func() {
		cache_Data_Array_ST_bind = Get_Control_Monad_ST_Internal_bind_()
	})
	return cache_Data_Array_ST_bind
}

var cache_Data_Array_ST_bind1 gopurs_runtime.Value
var once_Data_Array_ST_bind1 sync.Once

func Get_Data_Array_ST_bind1() gopurs_runtime.Value {
	once_Data_Array_ST_bind1.Do(func() {
		cache_Data_Array_ST_bind1 = Get_Control_Monad_ST_Internal_bind_()
	})
	return cache_Data_Array_ST_bind1
}

var cache_Data_Array_ST_negate gopurs_runtime.Value
var once_Data_Array_ST_negate sync.Once

func Get_Data_Array_ST_negate() gopurs_runtime.Value {
	once_Data_Array_ST_negate.Do(func() {
		cache_Data_Array_ST_negate = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Data_Array_ST_negate(a_0_box.IntVal))
		})
	})
	return cache_Data_Array_ST_negate
}

var cache_Data_Array_ST_bind2 gopurs_runtime.Value
var once_Data_Array_ST_bind2 sync.Once

func Get_Data_Array_ST_bind2() gopurs_runtime.Value {
	once_Data_Array_ST_bind2.Do(func() {
		cache_Data_Array_ST_bind2 = Get_Control_Monad_ST_Internal_bind_()
	})
	return cache_Data_Array_ST_bind2
}

var cache_Data_Array_ST_bind3 gopurs_runtime.Value
var once_Data_Array_ST_bind3 sync.Once

func Get_Data_Array_ST_bind3() gopurs_runtime.Value {
	once_Data_Array_ST_bind3.Do(func() {
		cache_Data_Array_ST_bind3 = Get_Control_Monad_ST_Internal_bind_()
	})
	return cache_Data_Array_ST_bind3
}

var cache_Data_Array_ST_pure gopurs_runtime.Value
var once_Data_Array_ST_pure sync.Once

func Get_Data_Array_ST_pure() gopurs_runtime.Value {
	once_Data_Array_ST_pure.Do(func() {
		cache_Data_Array_ST_pure = Get_Control_Monad_ST_Internal_pure_()
	})
	return cache_Data_Array_ST_pure
}

var cache_Data_Array_ST_unshiftAll gopurs_runtime.Value
var once_Data_Array_ST_unshiftAll sync.Once

func Get_Data_Array_ST_unshiftAll() gopurs_runtime.Value {
	once_Data_Array_ST_unshiftAll.Do(func() {
		cache_Data_Array_ST_unshiftAll = gopurs_runtime.Apply(Get_Control_Monad_ST_Uncurried_runSTFn2(), Get_Data_Array_ST_unshiftAllImpl())
	})
	return cache_Data_Array_ST_unshiftAll
}

var cache_Data_Array_ST_unshift gopurs_runtime.Value
var once_Data_Array_ST_unshift sync.Once

func Get_Data_Array_ST_unshift() gopurs_runtime.Value {
	once_Data_Array_ST_unshift.Do(func() {
		cache_Data_Array_ST_unshift = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_ST_unshift(a_0_box)
		})
	})
	return cache_Data_Array_ST_unshift
}

var cache_Data_Array_ST_unsafeThaw gopurs_runtime.Value
var once_Data_Array_ST_unsafeThaw sync.Once

func Get_Data_Array_ST_unsafeThaw() gopurs_runtime.Value {
	once_Data_Array_ST_unsafeThaw.Do(func() {
		cache_Data_Array_ST_unsafeThaw = gopurs_runtime.Apply(Get_Control_Monad_ST_Uncurried_runSTFn1(), Get_Data_Array_ST_unsafeThawImpl())
	})
	return cache_Data_Array_ST_unsafeThaw
}

var cache_Data_Array_ST_unsafeFreeze gopurs_runtime.Value
var once_Data_Array_ST_unsafeFreeze sync.Once

func Get_Data_Array_ST_unsafeFreeze() gopurs_runtime.Value {
	once_Data_Array_ST_unsafeFreeze.Do(func() {
		cache_Data_Array_ST_unsafeFreeze = gopurs_runtime.Apply(Get_Control_Monad_ST_Uncurried_runSTFn1(), Get_Data_Array_ST_unsafeFreezeImpl())
	})
	return cache_Data_Array_ST_unsafeFreeze
}

var cache_Data_Array_ST_toAssocArray gopurs_runtime.Value
var once_Data_Array_ST_toAssocArray sync.Once

func Get_Data_Array_ST_toAssocArray() gopurs_runtime.Value {
	once_Data_Array_ST_toAssocArray.Do(func() {
		cache_Data_Array_ST_toAssocArray = gopurs_runtime.Apply(Get_Control_Monad_ST_Uncurried_runSTFn1(), Get_Data_Array_ST_toAssocArrayImpl())
	})
	return cache_Data_Array_ST_toAssocArray
}

var cache_Data_Array_ST_thaw gopurs_runtime.Value
var once_Data_Array_ST_thaw sync.Once

func Get_Data_Array_ST_thaw() gopurs_runtime.Value {
	once_Data_Array_ST_thaw.Do(func() {
		cache_Data_Array_ST_thaw = gopurs_runtime.Apply(Get_Control_Monad_ST_Uncurried_runSTFn1(), Get_Data_Array_ST_thawImpl())
	})
	return cache_Data_Array_ST_thaw
}

var cache_Data_Array_ST_withArray gopurs_runtime.Value
var once_Data_Array_ST_withArray sync.Once

func Get_Data_Array_ST_withArray() gopurs_runtime.Value {
	once_Data_Array_ST_withArray.Do(func() {
		cache_Data_Array_ST_withArray = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_ST_withArray(f_0_box, *((*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)))
		})
	})
	return cache_Data_Array_ST_withArray
}

var cache_Data_Array_ST_splice gopurs_runtime.Value
var once_Data_Array_ST_splice sync.Once

func Get_Data_Array_ST_splice() gopurs_runtime.Value {
	once_Data_Array_ST_splice.Do(func() {
		cache_Data_Array_ST_splice = gopurs_runtime.Apply(Get_Control_Monad_ST_Uncurried_runSTFn4(), Get_Data_Array_ST_spliceImpl())
	})
	return cache_Data_Array_ST_splice
}

var cache_Data_Array_ST_sortBy gopurs_runtime.Value
var once_Data_Array_ST_sortBy sync.Once

func Get_Data_Array_ST_sortBy() gopurs_runtime.Value {
	once_Data_Array_ST_sortBy.Do(func() {
		cache_Data_Array_ST_sortBy = gopurs_runtime.Func(func(comp_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_ST_sortBy(comp_0_box)
		})
	})
	return cache_Data_Array_ST_sortBy
}

var cache_Data_Array_ST_sortWith gopurs_runtime.Value
var once_Data_Array_ST_sortWith sync.Once

func Get_Data_Array_ST_sortWith() gopurs_runtime.Value {
	once_Data_Array_ST_sortWith.Do(func() {
		cache_Data_Array_ST_sortWith = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_ST_sortWith(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box), f_1_box)
		})
	})
	return cache_Data_Array_ST_sortWith
}

var cache_Data_Array_ST_sort gopurs_runtime.Value
var once_Data_Array_ST_sort sync.Once

func Get_Data_Array_ST_sort() gopurs_runtime.Value {
	once_Data_Array_ST_sort.Do(func() {
		cache_Data_Array_ST_sort = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_ST_sort(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
		})
	})
	return cache_Data_Array_ST_sort
}

var cache_Data_Array_ST_shift gopurs_runtime.Value
var once_Data_Array_ST_shift sync.Once

func Get_Data_Array_ST_shift() gopurs_runtime.Value {
	once_Data_Array_ST_shift.Do(func() {
		cache_Data_Array_ST_shift = gopurs_runtime.Apply3(Get_Control_Monad_ST_Uncurried_runSTFn3(), Get_Data_Array_ST_shiftImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{gopurs_runtime.Value{}, false}
			if _v.V1 {
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
		}()))})
	})
	return cache_Data_Array_ST_shift
}

var cache_Data_Array_ST_run gopurs_runtime.Value
var once_Data_Array_ST_run sync.Once

func Get_Data_Array_ST_run() gopurs_runtime.Value {
	once_Data_Array_ST_run.Do(func() {
		cache_Data_Array_ST_run = gopurs_runtime.Func(func(st_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Array(Call_Data_Array_ST_run(st_0_box))
		})
	})
	return cache_Data_Array_ST_run
}

var cache_Data_Array_ST_pushAll gopurs_runtime.Value
var once_Data_Array_ST_pushAll sync.Once

func Get_Data_Array_ST_pushAll() gopurs_runtime.Value {
	once_Data_Array_ST_pushAll.Do(func() {
		cache_Data_Array_ST_pushAll = gopurs_runtime.Apply(Get_Control_Monad_ST_Uncurried_runSTFn2(), Get_Data_Array_ST_pushAllImpl())
	})
	return cache_Data_Array_ST_pushAll
}

var cache_Data_Array_ST_push gopurs_runtime.Value
var once_Data_Array_ST_push sync.Once

func Get_Data_Array_ST_push() gopurs_runtime.Value {
	once_Data_Array_ST_push.Do(func() {
		cache_Data_Array_ST_push = gopurs_runtime.Apply(Get_Control_Monad_ST_Uncurried_runSTFn2(), Get_Data_Array_ST_pushImpl())
	})
	return cache_Data_Array_ST_push
}

var cache_Data_Array_ST_pop gopurs_runtime.Value
var once_Data_Array_ST_pop sync.Once

func Get_Data_Array_ST_pop() gopurs_runtime.Value {
	once_Data_Array_ST_pop.Do(func() {
		cache_Data_Array_ST_pop = gopurs_runtime.Apply3(Get_Control_Monad_ST_Uncurried_runSTFn3(), Get_Data_Array_ST_popImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{gopurs_runtime.Value{}, false}
			if _v.V1 {
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
		}()))})
	})
	return cache_Data_Array_ST_pop
}

var cache_Data_Array_ST_poke gopurs_runtime.Value
var once_Data_Array_ST_poke sync.Once

func Get_Data_Array_ST_poke() gopurs_runtime.Value {
	once_Data_Array_ST_poke.Do(func() {
		cache_Data_Array_ST_poke = gopurs_runtime.Apply(Get_Control_Monad_ST_Uncurried_runSTFn3(), Get_Data_Array_ST_pokeImpl())
	})
	return cache_Data_Array_ST_poke
}

var cache_Data_Array_ST_peek gopurs_runtime.Value
var once_Data_Array_ST_peek sync.Once

func Get_Data_Array_ST_peek() gopurs_runtime.Value {
	once_Data_Array_ST_peek.Do(func() {
		cache_Data_Array_ST_peek = gopurs_runtime.Apply3(Get_Control_Monad_ST_Uncurried_runSTFn4(), Get_Data_Array_ST_peekImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{gopurs_runtime.Value{}, false}
			if _v.V1 {
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
		}()))})
	})
	return cache_Data_Array_ST_peek
}

var cache_Data_Array_ST_go__new gopurs_runtime.Value
var once_Data_Array_ST_go__new sync.Once

func Get_Data_Array_ST_go__new() gopurs_runtime.Value {
	once_Data_Array_ST_go__new.Do(func() {
		cache_Data_Array_ST_go__new = Get_Data_Array_ST_newImpl()
	})
	return cache_Data_Array_ST_go__new
}

var cache_Data_Array_ST_modify gopurs_runtime.Value
var once_Data_Array_ST_modify sync.Once

func Get_Data_Array_ST_modify() gopurs_runtime.Value {
	once_Data_Array_ST_modify.Do(func() {
		cache_Data_Array_ST_modify = gopurs_runtime.Func3(func(i_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, xs_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_ST_modify(i_0_box.IntVal, f_1_box, xs_2_box)
		})
	})
	return cache_Data_Array_ST_modify
}

var cache_Data_Array_ST_length gopurs_runtime.Value
var once_Data_Array_ST_length sync.Once

func Get_Data_Array_ST_length() gopurs_runtime.Value {
	once_Data_Array_ST_length.Do(func() {
		cache_Data_Array_ST_length = gopurs_runtime.Apply(Get_Control_Monad_ST_Uncurried_runSTFn1(), Get_Data_Array_ST_lengthImpl())
	})
	return cache_Data_Array_ST_length
}

var cache_Data_Array_ST_freeze gopurs_runtime.Value
var once_Data_Array_ST_freeze sync.Once

func Get_Data_Array_ST_freeze() gopurs_runtime.Value {
	once_Data_Array_ST_freeze.Do(func() {
		cache_Data_Array_ST_freeze = gopurs_runtime.Apply(Get_Control_Monad_ST_Uncurried_runSTFn1(), Get_Data_Array_ST_freezeImpl())
	})
	return cache_Data_Array_ST_freeze
}

var cache_Data_Array_ST_clone gopurs_runtime.Value
var once_Data_Array_ST_clone sync.Once

func Get_Data_Array_ST_clone() gopurs_runtime.Value {
	once_Data_Array_ST_clone.Do(func() {
		cache_Data_Array_ST_clone = gopurs_runtime.Apply(Get_Control_Monad_ST_Uncurried_runSTFn1(), Get_Data_Array_ST_cloneImpl())
	})
	return cache_Data_Array_ST_clone
}

func Call_Data_Array_ST_negate(a_0_loop int64) int64 {
	var a_0 int64 = a_0_loop
	_ = a_0
	return gopurs_runtime.IntNegate(a_0)
}

func Call_Data_Array_ST_unshift(a_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var a_0 gopurs_runtime.Value = a_0_loop
	_ = a_0
	return gopurs_runtime.Apply2(Get_Control_Monad_ST_Uncurried_runSTFn2(), Get_Data_Array_ST_unshiftAllImpl(), gopurs_runtime.Array([]gopurs_runtime.Value{a_0}))
}

func Call_Data_Array_ST_withArray(f_0_loop gopurs_runtime.Value, xs_1_loop []gopurs_runtime.Value) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	var xs_1 []gopurs_runtime.Value = xs_1_loop
	_ = xs_1
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		result_2_0 := gopurs_runtime.UncurriedApp(Get_Data_Array_ST_thawImpl(), gopurs_runtime.Array(xs_1))
		_ = result_2_0
		_dollar___unused_3_1 := gopurs_runtime.Apply2(f_0, result_2_0, gopurs_runtime.Value{})
		_ = _dollar___unused_3_1
		return gopurs_runtime.UncurriedApp(Get_Data_Array_ST_unsafeFreezeImpl(), result_2_0)
	})
}

func Call_Data_Array_ST_sortBy(comp_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var comp_0 gopurs_runtime.Value = comp_0_loop
	_ = comp_0
	return gopurs_runtime.Apply3(Get_Control_Monad_ST_Uncurried_runSTFn3(), Get_Data_Array_ST_sortByImpl(), comp_0, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
		var __t3 gopurs_runtime.Value
		{
			var __t_tag_0 uint32 = uint32(v_1.IntVal)
			_ = __t_tag_0
			if uint32(__t_tag_0) == 380165415 {
				__t3 = gopurs_runtime.Int(int64(1))
				goto end_branch_3
			} else {

			}
		}
		{
			var __t_tag_1 uint32 = uint32(v_1.IntVal)
			_ = __t_tag_1
			if uint32(__t_tag_1) == 902936544 {
				__t3 = gopurs_runtime.Int(int64(0))
				goto end_branch_3
			} else {

			}
		}
		{
			var __t_tag_2 uint32 = uint32(v_1.IntVal)
			_ = __t_tag_2
			if uint32(__t_tag_2) == 1527465420 {
				__t3 = gopurs_runtime.Int(int64(-1))
				goto end_branch_3
			} else {

			}
		}
		{
			__t3 = gopurs_runtime.Int(func() int64 { panic("Failed pattern match") }())
		}
	end_branch_3:
		return __t3
	}))
}

func Call_Data_Array_ST_sortWith(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
	_ = dictOrd_0
	var f_1 gopurs_runtime.Value = f_1_loop
	_ = f_1
	return Call_Data_Array_ST_sortBy(gopurs_runtime.Func2(func(x_2 gopurs_runtime.Value, y_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(dictOrd_0.V1, gopurs_runtime.Apply(f_1, x_2), gopurs_runtime.Apply(f_1, y_3))
	}))
}

func Call_Data_Array_ST_sort(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
	_ = dictOrd_0
	return Call_Data_Array_ST_sortBy(Call_Data_Ord_compare(dictOrd_0))
}

func Call_Data_Array_ST_run(st_0_loop gopurs_runtime.Value) []gopurs_runtime.Value {
	var st_0 gopurs_runtime.Value = st_0_loop
	_ = st_0
	return *((*[]gopurs_runtime.Value)(gopurs_runtime.Apply(Get_Control_Monad_ST_Internal_run(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		__local_var_1_0 := gopurs_runtime.Apply(st_0, gopurs_runtime.Value{})
		_ = __local_var_1_0
		return gopurs_runtime.Apply2(Get_Data_Array_ST_unsafeFreeze(), __local_var_1_0, gopurs_runtime.Value{})
	})).UnsafePtr))
}

func Call_Data_Array_ST_modify(i_0_loop int64, f_1_loop gopurs_runtime.Value, xs_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var i_0 int64 = i_0_loop
	_ = i_0
	var f_1 gopurs_runtime.Value = f_1_loop
	_ = f_1
	var xs_2 gopurs_runtime.Value = xs_2_loop
	_ = xs_2
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		entry_3_0 := gopurs_runtime.UncurriedApp4(Get_Data_Array_ST_peekImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}, gopurs_runtime.Int(i_0), xs_2)
		_ = entry_3_0
		var __t1 gopurs_runtime.Value
		{
			if entry_3_0.Type == 9 && entry_3_0.IntVal == 930809136 && entry_3_0.UnsafePtr != nil {
				__t1 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.UncurriedApp3(Get_Data_Array_ST_pokeImpl(), gopurs_runtime.Int(i_0), gopurs_runtime.Apply(f_1, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(entry_3_0.UnsafePtr).V0), xs_2)
				})
				goto end_branch_1
			} else {

			}
		}
		{
			if entry_3_0.Type == 9 && entry_3_0.IntVal == 930809136 && entry_3_0.UnsafePtr == nil {
				__t1 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Bool(false)
				})
				goto end_branch_1
			} else {

			}
		}
		{
			__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
		}
	end_branch_1:
		return gopurs_runtime.Apply(__t1, gopurs_runtime.Value{})
	})
}

func Get_Data_Array_ST_cloneImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Array_ST_CloneImpl
}

func Get_Data_Array_ST_freezeImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Array_ST_FreezeImpl
}

func Get_Data_Array_ST_lengthImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Array_ST_LengthImpl
}

func Get_Data_Array_ST_newImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Array_ST_NewImpl
}

func Get_Data_Array_ST_peekImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Array_ST_PeekImpl
}

func Get_Data_Array_ST_pokeImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Array_ST_PokeImpl
}

func Get_Data_Array_ST_popImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Array_ST_PopImpl
}

func Get_Data_Array_ST_pushAllImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Array_ST_PushAllImpl
}

func Get_Data_Array_ST_pushImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Array_ST_PushImpl
}

func Get_Data_Array_ST_shiftImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Array_ST_ShiftImpl
}

func Get_Data_Array_ST_sortByImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Array_ST_SortByImpl
}

func Get_Data_Array_ST_spliceImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Array_ST_SpliceImpl
}

func Get_Data_Array_ST_thawImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Array_ST_ThawImpl
}

func Get_Data_Array_ST_toAssocArrayImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Array_ST_ToAssocArrayImpl
}

func Get_Data_Array_ST_unsafeFreezeImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Array_ST_UnsafeFreezeImpl
}

func Get_Data_Array_ST_unsafeThawImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Array_ST_UnsafeThawImpl
}

func Get_Data_Array_ST_unshiftAllImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Array_ST_UnshiftAllImpl
}
