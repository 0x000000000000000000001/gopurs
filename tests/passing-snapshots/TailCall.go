package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_C gopurs_runtime.Value
var once_Main_C sync.Once

func Get_Main_C() gopurs_runtime.Value {
	once_Main_C.Do(func() {
		cache_Main_C = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 2167983901, UnsafePtr: unsafe.Pointer((&Constructor_Main_C[gopurs_runtime.Value]{1, value0, gopurs_runtime.CoerceToStruct[Constructor_Main_C[gopurs_runtime.Value]](value1)}))}
			})
		})
	})
	return cache_Main_C
}

var cache_Main_N gopurs_runtime.Value
var once_Main_N sync.Once

func Get_Main_N() gopurs_runtime.Value {
	once_Main_N.Do(func() {
		cache_Main_N = gopurs_runtime.Value{Type: 9, IntVal: 2167983901, UnsafePtr: unsafe.Pointer((*Constructor_Main_C[gopurs_runtime.Value])(nil))}
	})
	return cache_Main_N
}

var cache_Main_test gopurs_runtime.Value
var once_Main_test sync.Once

func Get_Main_test() gopurs_runtime.Value {
	once_Main_test.Do(func() {
		cache_Main_test = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Float(Call_Main_test(v_0_box.FloatVal(), gopurs_runtime.CoerceToStruct[Constructor_Main_C[float64]](v1_1_box)))
		})
	})
	return cache_Main_test
}

var cache_Main_notATailCall gopurs_runtime.Value
var once_Main_notATailCall sync.Once

func Get_Main_notATailCall() gopurs_runtime.Value {
	once_Main_notATailCall.Do(func() {
		cache_Main_notATailCall = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_notATailCall(x_0_box)
		})
	})
	return cache_Main_notATailCall
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=(ADT ["Effect","Effect"] [Unit])
			__local_var_0_0 := gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showNumberImpl(), gopurs_runtime.Float(Call_Main_test(0.0, (&Constructor_Main_C[float64]{1, gopurs_runtime.Float(1.0), (&Constructor_Main_C[gopurs_runtime.Value]{1, gopurs_runtime.Float(2.0), (&Constructor_Main_C[gopurs_runtime.Value]{1, gopurs_runtime.Float(3.0), (*Constructor_Main_C[gopurs_runtime.Value])(nil)})})})))).StrVal()))
			_ = __local_var_0_0
			_dollar___unused_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = _dollar___unused_1_1
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

var cache_Main_loop gopurs_runtime.Value
var once_Main_loop sync.Once

func Get_Main_loop() gopurs_runtime.Value {
	once_Main_loop.Do(func() {
		cache_Main_loop = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_loop(x_0_box.FloatVal())
		})
	})
	return cache_Main_loop
}

var cache_Main_loop__3510510677 gopurs_runtime.Value
var once_Main_loop__3510510677 sync.Once

func Get_Main_loop__3510510677() gopurs_runtime.Value {
	once_Main_loop__3510510677.Do(func() {
		cache_Main_loop__3510510677 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_loop__3510510677(x_0_box.FloatVal())
		})
	})
	return cache_Main_loop__3510510677
}

type Constructor_Main_C[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 *Constructor_Main_C[gopurs_runtime.Value]
}

type Constructor_Main_N[T_a any] struct {
	Rc uint32
}

func Call_Main_test(v_0_loop float64, v1_1_loop *Constructor_Main_C[float64]) float64 {
test:
	for {
		if false {
			continue test
		}
		var v_0 float64 = v_0_loop
		_ = v_0
		var v1_1 *Constructor_Main_C[float64] = v1_1_loop
		_ = v1_1
		var __t0 float64
		{
			if v1_1 == nil {
				__t0 = v_0
				goto end_branch_0
			} else {

			}
		}
		{
			if v1_1 != nil {
				v_0_loop = (v_0) + ((v1_1).V0.FloatVal())
				v1_1_loop = gopurs_runtime.CoerceToStruct[Constructor_Main_C[float64]](gopurs_runtime.Value{Type: 9, IntVal: 2167983901, UnsafePtr: unsafe.Pointer((v1_1).V1)})
				continue test
				__t0 = gopurs_runtime.Value{}.FloatVal()
				goto end_branch_0
			} else {

			}
		}
		{
			__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }().FloatVal()
		}
	end_branch_0:
		return __t0
	}
}

func Call_Main_notATailCall(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_loop(x_0_loop float64) gopurs_runtime.Value {
loop:
	for {
		if false {
			continue loop
		}
		var x_0 float64 = x_0_loop
		_ = x_0
		x_0_loop = (x_0) + (1.0)
		continue loop
		return gopurs_runtime.Value{}
	}
}

func Call_Main_loop__3510510677(x_0_loop float64) gopurs_runtime.Value {
	var x_0 float64 = x_0_loop
	_ = x_0
	return Call_Main_loop((x_0) + (2.0))
}
