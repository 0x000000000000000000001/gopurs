package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_Nothing gopurs_runtime.Value
var once_Main_Nothing sync.Once

func Get_Main_Nothing() gopurs_runtime.Value {
	once_Main_Nothing.Do(func() {
		cache_Main_Nothing = gopurs_runtime.Value{Type: 9, IntVal: 3271839782, UnsafePtr: unsafe.Pointer((*Constructor_Main_Just[gopurs_runtime.Value])(nil))}
	})
	return cache_Main_Nothing
}

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

var cache_Main_Just__1283418637 gopurs_runtime.Value
var once_Main_Just__1283418637 sync.Once

func Get_Main_Just__1283418637() gopurs_runtime.Value {
	once_Main_Just__1283418637.Do(func() {
		cache_Main_Just__1283418637 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3271839782, UnsafePtr: unsafe.Pointer(Rebox_Main_2656528265_4188394610(Call_Main_Just__1283418637(__eta_norm_0_0_box.IntVal)))}
		})
	})
	return cache_Main_Just__1283418637
}

var cache_Main_weirdsum gopurs_runtime.Value
var once_Main_weirdsum sync.Once

func Get_Main_weirdsum() gopurs_runtime.Value {
	once_Main_weirdsum.Do(func() {
		cache_Main_weirdsum = gopurs_runtime.Func3(func(accum_0_box gopurs_runtime.Value, f1_1_box gopurs_runtime.Value, n_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_weirdsum(accum_0_box.IntVal, f1_1_box, n_2_box.IntVal))
		})
	})
	return cache_Main_weirdsum
}

var cache_Main_tricksyinners gopurs_runtime.Value
var once_Main_tricksyinners sync.Once

func Get_Main_tricksyinners() gopurs_runtime.Value {
	once_Main_tricksyinners.Do(func() {
		cache_Main_tricksyinners = gopurs_runtime.Func2(func(accum_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_tricksyinners(accum_0_box.IntVal, x_1_box.IntVal))
		})
	})
	return cache_Main_tricksyinners
}

var cache_Main_g gopurs_runtime.Value
var once_Main_g sync.Once

func Get_Main_g() gopurs_runtime.Value {
	once_Main_g.Do(func() {
		cache_Main_g = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_g(x_0_box.IntVal))
		})
	})
	return cache_Main_g
}

var cache_Main_f gopurs_runtime.Value
var once_Main_f sync.Once

func Get_Main_f() gopurs_runtime.Value {
	once_Main_f.Do(func() {
		cache_Main_f = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_f(x_0_box.IntVal))
		})
	})
	return cache_Main_f
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___627669702("", struct {
			actual   int64
			expected int64
		}{Call_Main_f(int64(99999)), int64(0)}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___627669702("", struct {
				actual   int64
				expected int64
			}{Call_Main_g(int64(100000)), int64(0)}), gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___627669702("", struct {
					actual   int64
					expected int64
				}{Call_Main_weirdsum(int64(0), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
					var __t0 *Constructor_Main_Just[int64]
					{
						if (x_2.IntVal) < (int64(5)) {
							__t0 = (&Constructor_Main_Just[int64]{1, (int64(2)) * (x_2.IntVal)})
							goto end_branch_0
						} else {

						}
					}
					{
						__t0 = (*Constructor_Main_Just[int64])(nil)
					}
				end_branch_0:
					return gopurs_runtime.Value{Type: 9, IntVal: 3271839782, UnsafePtr: unsafe.Pointer(Rebox_Main_2656528265_4188394610(__t0))}
				}), int64(100000)), int64(20)}), gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___627669702("", struct {
						actual   int64
						expected int64
					}{Call_Main_tricksyinners(int64(0), int64(100000)), int64(200009)}), gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
					}))
				}))
			}))
		}))
	})
	return cache_Main_main
}

type Constructor_Main_Nothing[T_a any] struct {
	Rc uint32
}

type Constructor_Main_Just[T_a any] struct {
	Rc uint32
	V0 T_a
}

func Call_Main_Just__1283418637(__eta_norm_0_0_loop int64) *Constructor_Main_Just[int64] {
Just__1283418637:
	for {
		if false {
			continue Just__1283418637
		}
		var __eta_norm_0_0 int64 = __eta_norm_0_0_loop
		_ = __eta_norm_0_0
		return (&Constructor_Main_Just[int64]{1, __eta_norm_0_0})
	}
}

func Call_Main_weirdsum(accum_0_loop int64, f1_1_loop gopurs_runtime.Value, n_2_loop int64) int64 {
weirdsum:
	for {
		if false {
			continue weirdsum
		}
		var accum_0 int64 = accum_0_loop
		_ = accum_0
		var f1_1 gopurs_runtime.Value = f1_1_loop
		_ = f1_1
		var n_2 int64 = n_2_loop
		_ = n_2
		var __t2 int64
		{
			if (n_2) == (int64(0)) {
				__t2 = accum_0
				goto end_branch_2
			} else {

			}
		}
		{
			// TAST (Let): __local_var_3_0 shape=App(Other) bindingType=(ADT ["Main","Maybe"] [Int])
			__local_var_3_0 := Rebox_Main_4188394610_2656528265(gopurs_runtime.CoerceToStruct[Constructor_Main_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(f1_1, gopurs_runtime.Int(n_2))))
			_ = __local_var_3_0
			var __t1 int64
			{
				if __local_var_3_0 != nil {
					accum_0_loop = (accum_0) + ((__local_var_3_0).V0)
					f1_1_loop = f1_1
					n_2_loop = (n_2) - (int64(1))
					continue weirdsum
					__t1 = func() int64 { panic("unreachable") }()
					goto end_branch_1
				} else {

				}
			}
			{
				accum_0_loop = accum_0
				f1_1_loop = f1_1
				n_2_loop = (n_2) - (int64(1))
				continue weirdsum
				__t1 = func() int64 { panic("unreachable") }()
			}
		end_branch_1:
			__t2 = __t1
		}
	end_branch_2:
		return __t2
	}
}

func Call_Main_tricksyinners(accum_0_loop int64, x_1_loop int64) int64 {
tricksyinners:
	for {
		if false {
			continue tricksyinners
		}
		var accum_0 int64 = accum_0_loop
		_ = accum_0
		var x_1 int64 = x_1_loop
		_ = x_1
		var __t0 int64
		{
			if (x_1) == (int64(0)) {
				__t0 = (accum_0) + (((x_1) + (int64(3))) * ((x_1) + (int64(3))))
				goto end_branch_0
			} else {

			}
		}
		{
			accum_0_loop = (accum_0) + (int64(2))
			x_1_loop = (x_1) - (int64(1))
			continue tricksyinners
			__t0 = func() int64 { panic("unreachable") }()
		}
	end_branch_0:
		return __t0
	}
}

func Call_Main_g(x_0_loop int64) int64 {
g:
	for {
		if false {
			continue g
		}
		var x_0 int64 = x_0_loop
		_ = x_0
		var __t0 int64
		{
			if (x_0) == (int64(0)) {
				__t0 = int64(0)
				goto end_branch_0
			} else {

			}
		}
		{
			if (x_0) == (x_0) {
				x_0_loop = (x_0) - (int64(1))
				continue g
				__t0 = func() int64 { panic("unreachable") }()
				goto end_branch_0
			} else {

			}
		}
		{
			x_0_loop = (x_0) - (int64(2))
			continue g
			__t0 = func() int64 { panic("unreachable") }()
		}
	end_branch_0:
		return __t0
	}
}

func Call_Main_f(x_0_loop int64) int64 {
f:
	for {
		if false {
			continue f
		}
		var x_0 int64 = x_0_loop
		_ = x_0
		var __t0 int64
		{
			if (x_0) == (int64(0)) {
				__t0 = int64(0)
				goto end_branch_0
			} else {

			}
		}
		{
			x_0_loop = (x_0) - (int64(1))
			continue f
			__t0 = func() int64 { panic("unreachable") }()
		}
	end_branch_0:
		return __t0
	}
}

func Rebox_Main_2656528265_4188394610(in *Constructor_Main_Just[int64]) *Constructor_Main_Just[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_Just[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Int(in.V0)
	return out
}

func Rebox_Main_4188394610_2656528265(in *Constructor_Main_Just[gopurs_runtime.Value]) *Constructor_Main_Just[int64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_Just[int64]{}
	out.V0 = in.V0.IntVal
	return out
}
