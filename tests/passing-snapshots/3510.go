package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_Just gopurs_runtime.Value
var once_Main_Just sync.Once

func Get_Main_Just() gopurs_runtime.Value {
	once_Main_Just.Do(func() {
		cache_Main_Just = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3271839782, UnsafePtr: unsafe.Pointer((&Constructor_Main_Just[gopurs_runtime.Value]{1, value0}))}
		})
	})
	return cache_Main_Just
}

var cache_Main_Nothing gopurs_runtime.Value
var once_Main_Nothing sync.Once

func Get_Main_Nothing() gopurs_runtime.Value {
	once_Main_Nothing.Do(func() {
		cache_Main_Nothing = gopurs_runtime.Value{Type: 9, IntVal: 3271839782, UnsafePtr: unsafe.Pointer((*Constructor_Main_Just[gopurs_runtime.Value])(nil))}
	})
	return cache_Main_Nothing
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

var cache_Main_eqT gopurs_runtime.Value
var once_Main_eqT sync.Once

func Get_Main_eqT() gopurs_runtime.Value {
	once_Main_eqT.Do(func() {
		cache_Main_eqT = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_603498402_3790796878((&Constructor_Data_Eq_Eq[*Constructor_Main_Just[int64]]{1, gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t5 bool
			{
				var __t_tag_3 *Constructor_Main_Just[int64] = Rebox_Main_4188394610_2656528265(gopurs_runtime.CoerceToStruct[Constructor_Main_Just[gopurs_runtime.Value]](x_0))
				_ = __t_tag_3
				if __t_tag_3 != nil {
					var __t_tag_4 *Constructor_Main_Just[int64] = Rebox_Main_4188394610_2656528265(gopurs_runtime.CoerceToStruct[Constructor_Main_Just[gopurs_runtime.Value]](y_1))
					_ = __t_tag_4
					__t5 = (__t_tag_4 != nil) && (((*Constructor_Main_Just[gopurs_runtime.Value])(x_0.UnsafePtr).V0.IntVal) == ((*Constructor_Main_Just[gopurs_runtime.Value])(y_1.UnsafePtr).V0.IntVal))
					goto end_branch_5
				} else {

				}
			}
			{
				var __t_tag_0 *Constructor_Main_Just[int64] = Rebox_Main_4188394610_2656528265(gopurs_runtime.CoerceToStruct[Constructor_Main_Just[gopurs_runtime.Value]](x_0))
				_ = __t_tag_0
				var __t_and_2 bool = false
				if __t_tag_0 == nil {

					var __t_tag_1 *Constructor_Main_Just[int64] = Rebox_Main_4188394610_2656528265(gopurs_runtime.CoerceToStruct[Constructor_Main_Just[gopurs_runtime.Value]](y_1))
					_ = __t_tag_1
					__t_and_2 = (__t_tag_1 == nil)
				}
				__t5 = __t_and_2
			}
		end_branch_5:
			return gopurs_runtime.Bool(__t5)
		})})))}
	})
	return cache_Main_eqT
}

type Constructor_Main_Just[T_a any] struct {
	Rc uint32
	V0 T_a
}

type Constructor_Main_Nothing[T_a any] struct {
	Rc uint32
}

func Rebox_Main_4188394610_2656528265(in *Constructor_Main_Just[gopurs_runtime.Value]) *Constructor_Main_Just[int64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_Just[int64]{}
	out.V0 = in.V0.IntVal
	return out
}

func Rebox_Main_603498402_3790796878(in *Constructor_Data_Eq_Eq[*Constructor_Main_Just[int64]]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}
