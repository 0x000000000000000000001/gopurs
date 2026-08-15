package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_Con_Structor gopurs_runtime.Value
var once_Main_Con_Structor sync.Once

func Get_Main_Con_Structor() gopurs_runtime.Value {
	once_Main_Con_Structor.Do(func() {
		cache_Main_Con_Structor = gopurs_runtime.Value{Type: 9, IntVal: 762951793, UnsafePtr: unsafe.Pointer((*Constructor_Main_Con_2)(nil))}
	})
	return cache_Main_Con_Structor
}

var cache_Main_Con_2 gopurs_runtime.Value
var once_Main_Con_2 sync.Once

func Get_Main_Con_2() gopurs_runtime.Value {
	once_Main_Con_2.Do(func() {
		cache_Main_Con_2 = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 762951793, UnsafePtr: unsafe.Pointer((&Constructor_Main_Con_2{1, value0.StrVal()}))}
		})
	})
	return cache_Main_Con_2
}

var cache_Main_done gopurs_runtime.Value
var once_Main_done sync.Once

func Get_Main_done() gopurs_runtime.Value {
	once_Main_done.Do(func() {
		cache_Main_done = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_done(gopurs_runtime.CoerceToStruct[Constructor_Main_Con_2](v_0_box)))
		})
	})
	return cache_Main_done
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

type Constructor_Main_Con_Structor struct {
	Rc uint32
}

type Constructor_Main_Con_2 struct {
	Rc uint32
	V0 string
}

func Call_Main_done(v_0_loop *Constructor_Main_Con_2) string {
	var v_0 *Constructor_Main_Con_2 = v_0_loop
	_ = v_0
	var __t0 string
	{
		if v_0 != nil {
			__t0 = (v_0).V0
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = "Failed"
	}
end_branch_0:
	return __t0
}
