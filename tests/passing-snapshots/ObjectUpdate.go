package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_update2 gopurs_runtime.Value
var once_Main_update2 sync.Once

func Get_Main_update2() gopurs_runtime.Value {
	once_Main_update2.Do(func() {
		cache_Main_update2 = gopurs_runtime.Func(func(o_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_update2(o_0_box)
		})
	})
	return cache_Main_update2
}

var cache_Main_update1 gopurs_runtime.Value
var once_Main_update1 sync.Once

func Get_Main_update1() gopurs_runtime.Value {
	once_Main_update1.Do(func() {
		cache_Main_update1 = gopurs_runtime.Func(func(o_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_update1(o_0_box)
		})
	})
	return cache_Main_update1
}

var cache_Main_replace gopurs_runtime.Value
var once_Main_replace sync.Once

func Get_Main_replace() gopurs_runtime.Value {
	once_Main_replace.Do(func() {
		cache_Main_replace = gopurs_runtime.Func(func(o_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_replace(o_0_box)
		})
	})
	return cache_Main_replace
}

var cache_Main_polyUpdate gopurs_runtime.Value
var once_Main_polyUpdate sync.Once

func Get_Main_polyUpdate() gopurs_runtime.Value {
	once_Main_polyUpdate.Do(func() {
		cache_Main_polyUpdate = gopurs_runtime.Func(func(o_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_polyUpdate(o_0_box)
		})
	})
	return cache_Main_polyUpdate
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.RecordGet(Call_Main_update1(func() gopurs_runtime.Value {
			orig := struct {
				foo string
			}{""}
			_ = orig
			return gopurs_runtime.RecordDict1("foo", gopurs_runtime.Str(orig.foo))
		}()), "foo").StrVal())), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
		}))
	})
	return cache_Main_main
}

var cache_Main_inferPolyUpdate gopurs_runtime.Value
var once_Main_inferPolyUpdate sync.Once

func Get_Main_inferPolyUpdate() gopurs_runtime.Value {
	once_Main_inferPolyUpdate.Do(func() {
		cache_Main_inferPolyUpdate = gopurs_runtime.Func(func(o_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_inferPolyUpdate(o_0_box)
		})
	})
	return cache_Main_inferPolyUpdate
}

func Call_Main_update2(o_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var o_0 gopurs_runtime.Value = o_0_loop
	_ = o_0
	return gopurs_runtime.RecordUpdate1(o_0, "foo", gopurs_runtime.Str("Foo"))
}

func Call_Main_update1(o_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var o_0 gopurs_runtime.Value = o_0_loop
	_ = o_0
	return gopurs_runtime.RecordUpdate1(o_0, "foo", gopurs_runtime.Str("Foo"))
}

func Call_Main_replace(o_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var o_0 gopurs_runtime.Value = o_0_loop
	_ = o_0
	var __t0 gopurs_runtime.Value
	{
		if (gopurs_runtime.RecordGet(o_0, "foo").StrVal()) == ("Foo") {
			__t0 = gopurs_runtime.RecordUpdate1(o_0, "foo", gopurs_runtime.Str("Bar"))
			goto end_branch_0
		} else {

		}
	}
	{
		if (gopurs_runtime.RecordGet(o_0, "foo").StrVal()) == ("Bar") {
			__t0 = gopurs_runtime.RecordUpdate1(o_0, "bar", gopurs_runtime.Str("Baz"))
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = o_0
	}
end_branch_0:
	return __t0
}

func Call_Main_polyUpdate(o_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var o_0 gopurs_runtime.Value = o_0_loop
	_ = o_0
	return gopurs_runtime.RecordUpdate1(o_0, "foo", gopurs_runtime.Str("Foo"))
}

func Call_Main_inferPolyUpdate(o_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var o_0 gopurs_runtime.Value = o_0_loop
	_ = o_0
	return gopurs_runtime.RecordUpdate1(o_0, "foo", gopurs_runtime.Str("Foo"))
}
