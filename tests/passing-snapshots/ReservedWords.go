package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_o gopurs_runtime.Value
var once_Main_o sync.Once

func Get_Main_o() gopurs_runtime.Value {
	once_Main_o.Do(func() {
		cache_Main_o = gopurs_runtime.RecordDict1("type", gopurs_runtime.Str("o"))
	})
	return cache_Main_o
}

var cache_Main_p gopurs_runtime.Value
var once_Main_p sync.Once

func Get_Main_p() gopurs_runtime.Value {
	once_Main_p.Do(func() {
		cache_Main_p = func() gopurs_runtime.Value {
			origVal := Get_Main_o()
			if origVal.Type != gopurs_runtime.TypeRecord1 {
				return gopurs_runtime.RecordUpdateDict(origVal, []string{"type"}, []gopurs_runtime.Value{gopurs_runtime.Str("p")})
			}
			clone := *((*gopurs_runtime.RecordData1)(origVal.UnsafePtr))
			clone.V0 = gopurs_runtime.Str("p")
			return gopurs_runtime.Value{Type: gopurs_runtime.TypeRecord1, UnsafePtr: unsafe.Pointer(&clone)}
		}()
	})
	return cache_Main_p
}

var cache_Main_f gopurs_runtime.Value
var once_Main_f sync.Once

func Get_Main_f() gopurs_runtime.Value {
	once_Main_f.Do(func() {
		cache_Main_f = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_f(v_0_box))
		})
	})
	return cache_Main_f
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = func() gopurs_runtime.Value {
			var __t0 string
			{
				if (gopurs_runtime.RecordGet(Get_Main_p(), "type").StrVal()) == ("p") {
					__t0 = "Done"
					goto end_branch_0
				} else {

				}
			}
			{
				__t0 = "Fail"
			}
		end_branch_0:
			return gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(__t0))
		}()
	})
	return cache_Main_main
}

var cache_Main_f__3025627738 gopurs_runtime.Value
var once_Main_f__3025627738 sync.Once

func Get_Main_f__3025627738() gopurs_runtime.Value {
	once_Main_f__3025627738.Do(func() {
		cache_Main_f__3025627738 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_f__3025627738(v_0_box))
		})
	})
	return cache_Main_f__3025627738
}

func Call_Main_f(v_0_loop gopurs_runtime.Value) string {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	var __t0 string
	{
		if (gopurs_runtime.RecordGet(v_0, "type").StrVal()) == ("p") {
			__t0 = "Done"
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = "Fail"
	}
end_branch_0:
	return __t0
}

func Call_Main_f__3025627738(v_0_loop gopurs_runtime.Value) string {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	var __t0 string
	{
		if (gopurs_runtime.RecordGet(v_0, "type").StrVal()) == ("p") {
			__t0 = "Done"
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = "Fail"
	}
end_branch_0:
	return __t0
}
