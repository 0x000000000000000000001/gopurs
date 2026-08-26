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

var cache_Main_Id gopurs_runtime.Value
var once_Main_Id sync.Once

func Get_Main_Id() gopurs_runtime.Value {
	once_Main_Id.Do(func() {
		cache_Main_Id = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return value0
		})
	})
	return cache_Main_Id
}

var cache_Main_test gopurs_runtime.Value
var once_Main_test sync.Once

func Get_Main_test() gopurs_runtime.Value {
	once_Main_test.Do(func() {
		cache_Main_test = gopurs_runtime.Func(func(m_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_test(m_0_box)
		})
	})
	return cache_Main_test
}

var cache_Main_test__1633551430 gopurs_runtime.Value
var once_Main_test__1633551430 sync.Once

func Get_Main_test__1633551430() gopurs_runtime.Value {
	once_Main_test__1633551430.Do(func() {
		cache_Main_test__1633551430 = gopurs_runtime.Func(func(m_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_test__1633551430(m_0_box)
		})
	})
	return cache_Main_test__1633551430
}

var cache_Main_maybe gopurs_runtime.Value
var once_Main_maybe sync.Once

func Get_Main_maybe() gopurs_runtime.Value {
	once_Main_maybe.Do(func() {
		cache_Main_maybe = gopurs_runtime.RecordDict2("bind", "return", gopurs_runtime.Func(func(ma_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t0 *Constructor_Main_Just[gopurs_runtime.Value]
				{
					if ma_0.Type == 9 && ma_0.IntVal == 3271839782 && ma_0.UnsafePtr == nil {
						__t0 = (*Constructor_Main_Just[gopurs_runtime.Value])(nil)
						goto end_branch_0
					} else {

					}
				}
				{
					if ma_0.Type == 9 && ma_0.IntVal == 3271839782 && ma_0.UnsafePtr != nil {
						__t0 = gopurs_runtime.CoerceToStruct[Constructor_Main_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(f_1, (*Constructor_Main_Just[gopurs_runtime.Value])(ma_0.UnsafePtr).V0))
						goto end_branch_0
					} else {

					}
				}
				{
					__t0 = gopurs_runtime.CoerceToStruct[Constructor_Main_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value { panic("Failed pattern match") }())
				}
			end_branch_0:
				return gopurs_runtime.Value{Type: 9, IntVal: 3271839782, UnsafePtr: unsafe.Pointer(__t0)}
			})
		}), Get_Main_Just())
	})
	return cache_Main_maybe
}

var cache_Main_maybe__2826186755 gopurs_runtime.Value
var once_Main_maybe__2826186755 sync.Once

func Get_Main_maybe__2826186755() gopurs_runtime.Value {
	once_Main_maybe__2826186755.Do(func() {
		cache_Main_maybe__2826186755 = gopurs_runtime.RecordDict2("bind", "return", gopurs_runtime.Func(func(ma_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t0 *Constructor_Main_Just[gopurs_runtime.Value]
				{
					if ma_0.Type == 9 && ma_0.IntVal == 3271839782 && ma_0.UnsafePtr == nil {
						__t0 = (*Constructor_Main_Just[gopurs_runtime.Value])(nil)
						goto end_branch_0
					} else {

					}
				}
				{
					if ma_0.Type == 9 && ma_0.IntVal == 3271839782 && ma_0.UnsafePtr != nil {
						__t0 = gopurs_runtime.CoerceToStruct[Constructor_Main_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(f_1, (*Constructor_Main_Just[gopurs_runtime.Value])(ma_0.UnsafePtr).V0))
						goto end_branch_0
					} else {

					}
				}
				{
					__t0 = gopurs_runtime.CoerceToStruct[Constructor_Main_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value { panic("Failed pattern match") }())
				}
			end_branch_0:
				return gopurs_runtime.Value{Type: 9, IntVal: 3271839782, UnsafePtr: unsafe.Pointer(__t0)}
			})
		}), Get_Main_Just())
	})
	return cache_Main_maybe__2826186755
}

var cache_Main_test2 gopurs_runtime.Value
var once_Main_test2 sync.Once

func Get_Main_test2() gopurs_runtime.Value {
	once_Main_test2.Do(func() {
		cache_Main_test2 = gopurs_runtime.Value{Type: 9, IntVal: 3271839782, UnsafePtr: unsafe.Pointer((&Constructor_Main_Just[gopurs_runtime.Value]{1, gopurs_runtime.Float(1.0)}))}
	})
	return cache_Main_test2
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

var cache_Main_id gopurs_runtime.Value
var once_Main_id sync.Once

func Get_Main_id() gopurs_runtime.Value {
	once_Main_id.Do(func() {
		cache_Main_id = gopurs_runtime.RecordDict2("bind", "return", gopurs_runtime.Func(func(ma_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply(f_1, ma_0)
			})
		}), Get_Main_Id())
	})
	return cache_Main_id
}

var cache_Main_id__2356412419 gopurs_runtime.Value
var once_Main_id__2356412419 sync.Once

func Get_Main_id__2356412419() gopurs_runtime.Value {
	once_Main_id__2356412419.Do(func() {
		cache_Main_id__2356412419 = gopurs_runtime.RecordDict2("bind", "return", gopurs_runtime.Func(func(ma_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply(f_1, ma_0)
			})
		}), Get_Main_Id())
	})
	return cache_Main_id__2356412419
}

var cache_Main_test1 gopurs_runtime.Value
var once_Main_test1 sync.Once

func Get_Main_test1() gopurs_runtime.Value {
	once_Main_test1.Do(func() {
		cache_Main_test1 = gopurs_runtime.Float(1.0)
	})
	return cache_Main_test1
}

type Constructor_Main_Nothing[T_a any] struct {
	Rc uint32
}

type Constructor_Main_Just[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

type Constructor_Main_Id[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

func Call_Main_test(m_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var m_0 gopurs_runtime.Value = m_0_loop
	_ = m_0
	return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(m_0, "bind"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(m_0, "return"), gopurs_runtime.Float(1.0)), gopurs_runtime.Func(func(n1_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(m_0, "bind"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(m_0, "return"), gopurs_runtime.Str("Test")), gopurs_runtime.Func(func(n2_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(gopurs_runtime.RecordGet(m_0, "return"), gopurs_runtime.Float(n1_1.FloatVal()))
		}))
	}))
}

func Call_Main_test__1633551430(m_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var m_0 gopurs_runtime.Value = m_0_loop
	_ = m_0
	return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(m_0, "bind"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(m_0, "return"), gopurs_runtime.Float(1.0)), gopurs_runtime.Func(func(n1_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(m_0, "bind"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(m_0, "return"), gopurs_runtime.Str("Test")), gopurs_runtime.Func(func(n2_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(gopurs_runtime.RecordGet(m_0, "return"), gopurs_runtime.Float(n1_1.FloatVal()))
		}))
	}))
}
