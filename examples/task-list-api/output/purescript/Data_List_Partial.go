package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_List_Partial_tail gopurs_runtime.Value
var once_Data_List_Partial_tail sync.Once

func Get_Data_List_Partial_tail() gopurs_runtime.Value {
	once_Data_List_Partial_tail.Do(func() {
		cache_Data_List_Partial_tail = gopurs_runtime.Func2(func(_dollar___unused_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_Partial_tail(_dollar___unused_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_1_box)))}
		})
	})
	return cache_Data_List_Partial_tail
}

var cache_Data_List_Partial_last gopurs_runtime.Value
var once_Data_List_Partial_last sync.Once

func Get_Data_List_Partial_last() gopurs_runtime.Value {
	once_Data_List_Partial_last.Do(func() {
		cache_Data_List_Partial_last = gopurs_runtime.Func2(func(_dollar___unused_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_Partial_last(_dollar___unused_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_1_box))
		})
	})
	return cache_Data_List_Partial_last
}

var cache_Data_List_Partial_go__init gopurs_runtime.Value
var once_Data_List_Partial_go__init sync.Once

func Get_Data_List_Partial_go__init() gopurs_runtime.Value {
	once_Data_List_Partial_go__init.Do(func() {
		cache_Data_List_Partial_go__init = gopurs_runtime.Func2(func(_dollar___unused_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_Partial_go__init(_dollar___unused_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_1_box)))}
		})
	})
	return cache_Data_List_Partial_go__init
}

var cache_Data_List_Partial_head gopurs_runtime.Value
var once_Data_List_Partial_head sync.Once

func Get_Data_List_Partial_head() gopurs_runtime.Value {
	once_Data_List_Partial_head.Do(func() {
		cache_Data_List_Partial_head = gopurs_runtime.Func2(func(_dollar___unused_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_Partial_head(_dollar___unused_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_1_box))
		})
	})
	return cache_Data_List_Partial_head
}

func Call_Data_List_Partial_tail(_dollar___unused_0_loop gopurs_runtime.Value, v_1_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var v_1 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_1_loop
	_ = v_1
	var __t0 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
	{
		if v_1 != nil {
			__t0 = (v_1).V1
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
	}
end_branch_0:
	return __t0
}

func Call_Data_List_Partial_last(_dollar___unused_0_loop gopurs_runtime.Value, v_1_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
last:
	for {
		if false {
			continue last
		}
		var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
		_ = _dollar___unused_0
		var v_1 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_1_loop
		_ = v_1
		var __t2 gopurs_runtime.Value
		{
			if v_1 != nil {
				var __t1 gopurs_runtime.Value
				{
					var __t_tag_0 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = (v_1).V1
					_ = __t_tag_0
					if __t_tag_0 == nil {
						__t1 = (v_1).V0
						goto end_branch_1
					} else {

					}
				}
				{
					_dollar___unused_0_loop = gopurs_runtime.Value{}
					v_1_loop = (v_1).V1
					continue last
					__t1 = func() gopurs_runtime.Value { panic("unreachable") }()
				}
			end_branch_1:
				__t2 = __t1
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
}

func Call_Data_List_Partial_go__init(_dollar___unused_0_loop gopurs_runtime.Value, v_1_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
go__init:
	for {
		if false {
			continue go__init
		}
		var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
		_ = _dollar___unused_0
		var v_1 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_1_loop
		_ = v_1
		var __t2 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
		{
			if v_1 != nil {
				var __t1 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
				{
					var __t_tag_0 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = (v_1).V1
					_ = __t_tag_0
					if __t_tag_0 == nil {
						__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))})
						goto end_branch_1
					} else {

					}
				}
				{
					__t1 = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v_1).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(Get_Data_List_Partial_go__init(), gopurs_runtime.Value{}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((v_1).V1)}))})
				}
			end_branch_1:
				__t2 = __t1
				goto end_branch_2
			} else {

			}
		}
		{
			__t2 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
		}
	end_branch_2:
		return __t2
	}
}

func Call_Data_List_Partial_head(_dollar___unused_0_loop gopurs_runtime.Value, v_1_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var v_1 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_1_loop
	_ = v_1
	var __t0 gopurs_runtime.Value
	{
		if v_1 != nil {
			__t0 = (v_1).V0
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
