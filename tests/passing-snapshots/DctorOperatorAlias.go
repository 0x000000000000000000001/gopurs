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
			return Call_Main_get3(v_0_box, gopurs_runtime.CoerceToStruct[Constructor_List_Cons[gopurs_runtime.Value]](v1_1_box))
		})
	})
	return cache_Main_get3
}

var cache_Main_get3__1718674842 gopurs_runtime.Value
var once_Main_get3__1718674842 sync.Once

func Get_Main_get3__1718674842() gopurs_runtime.Value {
	once_Main_get3__1718674842.Do(func() {
		cache_Main_get3__1718674842 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Float(Call_Main_get3__1718674842(v_0_box.FloatVal(), Rebox_Main_2855836466_3211731498(gopurs_runtime.CoerceToStruct[Constructor_List_Cons[gopurs_runtime.Value]](v1_1_box))))
		})
	})
	return cache_Main_get3__1718674842
}

var cache_Main_get2 gopurs_runtime.Value
var once_Main_get2 sync.Once

func Get_Main_get2() gopurs_runtime.Value {
	once_Main_get2.Do(func() {
		cache_Main_get2 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_get2(v_0_box, gopurs_runtime.CoerceToStruct[Constructor_List_Cons[gopurs_runtime.Value]](v1_1_box))
		})
	})
	return cache_Main_get2
}

var cache_Main_get2__2795389194 gopurs_runtime.Value
var once_Main_get2__2795389194 sync.Once

func Get_Main_get2__2795389194() gopurs_runtime.Value {
	once_Main_get2__2795389194.Do(func() {
		cache_Main_get2__2795389194 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_get2__2795389194(v_0_box.IntVal, Rebox_Main_2855836466_3483143753(gopurs_runtime.CoerceToStruct[Constructor_List_Cons[gopurs_runtime.Value]](v1_1_box))))
		})
	})
	return cache_Main_get2__2795389194
}

var cache_Main_get1 gopurs_runtime.Value
var once_Main_get1 sync.Once

func Get_Main_get1() gopurs_runtime.Value {
	once_Main_get1.Do(func() {
		cache_Main_get1 = gopurs_runtime.Func2(func(y_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_get1(y_0_box, gopurs_runtime.CoerceToStruct[Constructor_List_Cons[gopurs_runtime.Value]](xs_1_box))
		})
	})
	return cache_Main_get1
}

var cache_Main_get1__2795389194 gopurs_runtime.Value
var once_Main_get1__2795389194 sync.Once

func Get_Main_get1__2795389194() gopurs_runtime.Value {
	once_Main_get1__2795389194.Do(func() {
		cache_Main_get1__2795389194 = gopurs_runtime.Func2(func(y_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_get1__2795389194(y_0_box.IntVal, Rebox_Main_2855836466_3483143753(gopurs_runtime.CoerceToStruct[Constructor_List_Cons[gopurs_runtime.Value]](xs_1_box))))
		})
	})
	return cache_Main_get1__2795389194
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("Incorrect result!"), gopurs_runtime.Bool(true)), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("Incorrect result!"), gopurs_runtime.Bool(true)), gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("Incorrect result!"), gopurs_runtime.Bool(true)), gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
				}))
			}))
		}))
	})
	return cache_Main_main
}

func Call_Main_get3(v_0_loop gopurs_runtime.Value, v1_1_loop *Constructor_List_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	var v1_1 *Constructor_List_Cons[gopurs_runtime.Value] = v1_1_loop
	_ = v1_1
	var __t2 gopurs_runtime.Value
	{
		var __t_and_1 bool = false
		if v1_1 != nil {

			var __t_tag_0 *Constructor_List_Cons[gopurs_runtime.Value] = (v1_1).V1
			_ = __t_tag_0
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

func Call_Main_get3__1718674842(v_0_loop float64, v1_1_loop *Constructor_List_Cons[float64]) float64 {
get3__1718674842:
	for {
		if false {
			continue get3__1718674842
		}
		var v_0 float64 = v_0_loop
		_ = v_0
		var v1_1 *Constructor_List_Cons[float64] = v1_1_loop
		_ = v1_1
		var __t2 float64
		{
			var __t_and_1 bool = false
			if v1_1 != nil {

				var __t_tag_0 *Constructor_List_Cons[float64] = (v1_1).V1
				_ = __t_tag_0
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
}

func Call_Main_get2(v_0_loop gopurs_runtime.Value, v1_1_loop *Constructor_List_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	var v1_1 *Constructor_List_Cons[gopurs_runtime.Value] = v1_1_loop
	_ = v1_1
	var __t2 gopurs_runtime.Value
	{
		var __t_and_1 bool = false
		if v1_1 != nil {

			var __t_tag_0 *Constructor_List_Cons[gopurs_runtime.Value] = (v1_1).V1
			_ = __t_tag_0
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

func Call_Main_get2__2795389194(v_0_loop int64, v1_1_loop *Constructor_List_Cons[int64]) int64 {
get2__2795389194:
	for {
		if false {
			continue get2__2795389194
		}
		var v_0 int64 = v_0_loop
		_ = v_0
		var v1_1 *Constructor_List_Cons[int64] = v1_1_loop
		_ = v1_1
		var __t2 int64
		{
			var __t_and_1 bool = false
			if v1_1 != nil {

				var __t_tag_0 *Constructor_List_Cons[int64] = (v1_1).V1
				_ = __t_tag_0
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
}

func Call_Main_get1(y_0_loop gopurs_runtime.Value, xs_1_loop *Constructor_List_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
	var y_0 gopurs_runtime.Value = y_0_loop
	_ = y_0
	var xs_1 *Constructor_List_Cons[gopurs_runtime.Value] = xs_1_loop
	_ = xs_1
	var __t2 gopurs_runtime.Value
	{
		var __t_and_1 bool = false
		if xs_1 != nil {

			var __t_tag_0 *Constructor_List_Cons[gopurs_runtime.Value] = (xs_1).V1
			_ = __t_tag_0
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

func Call_Main_get1__2795389194(y_0_loop int64, xs_1_loop *Constructor_List_Cons[int64]) int64 {
get1__2795389194:
	for {
		if false {
			continue get1__2795389194
		}
		var y_0 int64 = y_0_loop
		_ = y_0
		var xs_1 *Constructor_List_Cons[int64] = xs_1_loop
		_ = xs_1
		var __t2 int64
		{
			var __t_and_1 bool = false
			if xs_1 != nil {

				var __t_tag_0 *Constructor_List_Cons[int64] = (xs_1).V1
				_ = __t_tag_0
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
}

func Rebox_Main_2855836466_3211731498(in *Constructor_List_Cons[gopurs_runtime.Value]) *Constructor_List_Cons[float64] {
	if in == nil {
		return nil
	}
	out := &Constructor_List_Cons[float64]{}
	out.V0 = in.V0.FloatVal()
	out.V1 = Rebox_Main_2855836466_3211731498(in.V1)
	return out
}

func Rebox_Main_2855836466_3483143753(in *Constructor_List_Cons[gopurs_runtime.Value]) *Constructor_List_Cons[int64] {
	if in == nil {
		return nil
	}
	out := &Constructor_List_Cons[int64]{}
	out.V0 = in.V0.IntVal
	out.V1 = Rebox_Main_2855836466_3483143753(in.V1)
	return out
}
