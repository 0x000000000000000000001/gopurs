package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_get3 gopurs_runtime.Value
var once_Main_get3 sync.Once

func Get_Main_get3() gopurs_runtime.Value {
	once_Main_get3.Do(func() {
		cache_Main_get3 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_get3(v_0_box, gopurs_runtime.CoerceToStruct[Constructor_List_Cons](v1_1_box))
		})
	})
	return cache_Main_get3
}

var cache_Main_get2 gopurs_runtime.Value
var once_Main_get2 sync.Once

func Get_Main_get2() gopurs_runtime.Value {
	once_Main_get2.Do(func() {
		cache_Main_get2 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_get2(v_0_box, gopurs_runtime.CoerceToStruct[Constructor_List_Cons](v1_1_box))
		})
	})
	return cache_Main_get2
}

var cache_Main_get1 gopurs_runtime.Value
var once_Main_get1 sync.Once

func Get_Main_get1() gopurs_runtime.Value {
	once_Main_get1.Do(func() {
		cache_Main_get1 = gopurs_runtime.Func2(func(y_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_get1(y_0_box, gopurs_runtime.CoerceToStruct[Constructor_List_Cons](xs_1_box))
		})
	})
	return cache_Main_get1
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=App(Var) expectedFromAst=gopurs_runtime.Value actual=gopurs_runtime.Value bindingType=Any
			__local_var_0_0 := gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("Incorrect result!"), gopurs_runtime.Bool(true))
			_ = __local_var_0_0
			_dollar___unused_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = _dollar___unused_1_1
			_dollar___unused_2_2 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("Incorrect result!"), gopurs_runtime.Bool(true)), gopurs_runtime.Value{})
			_ = _dollar___unused_2_2
			_dollar___unused_3_3 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("Incorrect result!"), gopurs_runtime.Bool(true)), gopurs_runtime.Value{})
			_ = _dollar___unused_3_3
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

var cache_Main_get1__3420859751 gopurs_runtime.Value
var once_Main_get1__3420859751 sync.Once

func Get_Main_get1__3420859751() gopurs_runtime.Value {
	once_Main_get1__3420859751.Do(func() {
		cache_Main_get1__3420859751 = gopurs_runtime.Func2(func(y_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_get1__3420859751(y_0_box, gopurs_runtime.CoerceToStruct[Constructor_List_Cons](xs_1_box))
		})
	})
	return cache_Main_get1__3420859751
}

var cache_Main_get2__3420859751 gopurs_runtime.Value
var once_Main_get2__3420859751 sync.Once

func Get_Main_get2__3420859751() gopurs_runtime.Value {
	once_Main_get2__3420859751.Do(func() {
		cache_Main_get2__3420859751 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_get2__3420859751(v_0_box, gopurs_runtime.CoerceToStruct[Constructor_List_Cons](v1_1_box))
		})
	})
	return cache_Main_get2__3420859751
}

var cache_Main_get3__3420859751 gopurs_runtime.Value
var once_Main_get3__3420859751 sync.Once

func Get_Main_get3__3420859751() gopurs_runtime.Value {
	once_Main_get3__3420859751.Do(func() {
		cache_Main_get3__3420859751 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_get3__3420859751(v_0_box, gopurs_runtime.CoerceToStruct[Constructor_List_Cons](v1_1_box))
		})
	})
	return cache_Main_get3__3420859751
}

func Call_Main_get3(v_0_loop gopurs_runtime.Value, v1_1_loop *Constructor_List_Cons) gopurs_runtime.Value {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	var v1_1 *Constructor_List_Cons = v1_1_loop
	_ = v1_1
	var __t2 gopurs_runtime.Value
	{
		var __t_and_1 bool = false
		if v1_1 != nil {

			var __t_tag_0 *Constructor_List_Cons = (v1_1).V1
			__t_and_1 = (__t_tag_0 != nil)
		}
		if __t_and_1 {
			__t2 = ((v1_1).V1).V0
			goto end_branch_2
		} else {

		}
	}
	{
		__t2 = v_0
	}
end_branch_2:
	return __t2
}

func Call_Main_get2(v_0_loop gopurs_runtime.Value, v1_1_loop *Constructor_List_Cons) gopurs_runtime.Value {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	var v1_1 *Constructor_List_Cons = v1_1_loop
	_ = v1_1
	var __t2 gopurs_runtime.Value
	{
		var __t_and_1 bool = false
		if v1_1 != nil {

			var __t_tag_0 *Constructor_List_Cons = (v1_1).V1
			__t_and_1 = (__t_tag_0 != nil)
		}
		if __t_and_1 {
			__t2 = ((v1_1).V1).V0
			goto end_branch_2
		} else {

		}
	}
	{
		__t2 = v_0
	}
end_branch_2:
	return __t2
}

func Call_Main_get1(y_0_loop gopurs_runtime.Value, xs_1_loop *Constructor_List_Cons) gopurs_runtime.Value {
	var y_0 gopurs_runtime.Value = y_0_loop
	_ = y_0
	var xs_1 *Constructor_List_Cons = xs_1_loop
	_ = xs_1
	var __t2 gopurs_runtime.Value
	{
		var __t_and_1 bool = false
		if xs_1 != nil {

			var __t_tag_0 *Constructor_List_Cons = (xs_1).V1
			__t_and_1 = (__t_tag_0 != nil)
		}
		if __t_and_1 {
			__t2 = ((xs_1).V1).V0
			goto end_branch_2
		} else {

		}
	}
	{
		__t2 = y_0
	}
end_branch_2:
	return __t2
}

func Call_Main_get1__3420859751(y_0_loop gopurs_runtime.Value, xs_1_loop *Constructor_List_Cons) gopurs_runtime.Value {
	var y_0 gopurs_runtime.Value = y_0_loop
	_ = y_0
	var xs_1 *Constructor_List_Cons = xs_1_loop
	_ = xs_1
	var __t2 gopurs_runtime.Value
	{
		var __t_and_1 bool = false
		if xs_1 != nil {

			var __t_tag_0 *Constructor_List_Cons = (xs_1).V1
			__t_and_1 = (__t_tag_0 != nil)
		}
		if __t_and_1 {
			__t2 = ((xs_1).V1).V0
			goto end_branch_2
		} else {

		}
	}
	{
		__t2 = y_0
	}
end_branch_2:
	return __t2
}

func Call_Main_get2__3420859751(v_0_loop gopurs_runtime.Value, v1_1_loop *Constructor_List_Cons) gopurs_runtime.Value {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	var v1_1 *Constructor_List_Cons = v1_1_loop
	_ = v1_1
	var __t2 gopurs_runtime.Value
	{
		var __t_and_1 bool = false
		if v1_1 != nil {

			var __t_tag_0 *Constructor_List_Cons = (v1_1).V1
			__t_and_1 = (__t_tag_0 != nil)
		}
		if __t_and_1 {
			__t2 = ((v1_1).V1).V0
			goto end_branch_2
		} else {

		}
	}
	{
		__t2 = v_0
	}
end_branch_2:
	return __t2
}

func Call_Main_get3__3420859751(v_0_loop gopurs_runtime.Value, v1_1_loop *Constructor_List_Cons) gopurs_runtime.Value {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	var v1_1 *Constructor_List_Cons = v1_1_loop
	_ = v1_1
	var __t2 gopurs_runtime.Value
	{
		var __t_and_1 bool = false
		if v1_1 != nil {

			var __t_tag_0 *Constructor_List_Cons = (v1_1).V1
			__t_and_1 = (__t_tag_0 != nil)
		}
		if __t_and_1 {
			__t2 = ((v1_1).V1).V0
			goto end_branch_2
		} else {

		}
	}
	{
		__t2 = v_0
	}
end_branch_2:
	return __t2
}
