package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
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
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=App(Var) expectedFromAst=gopurs_runtime.Value actual=gopurs_runtime.Value bindingType=Any
			__local_var_0_0 := gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Foo"))
			_ = __local_var_0_0
			_dollar___unused_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = _dollar___unused_1_1
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), gopurs_runtime.Value{})
		})
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
	return func() gopurs_runtime.Value {
		origVal := o_0
		if origVal.Type != gopurs_runtime.TypeRecord1 {
			return gopurs_runtime.RecordUpdateDict(origVal, []string{"foo"}, []gopurs_runtime.Value{gopurs_runtime.Str("Foo")})
		}
		clone := *((*gopurs_runtime.RecordData1)(origVal.UnsafePtr))
		clone.V0 = gopurs_runtime.Str("Foo")
		return gopurs_runtime.Value{Type: gopurs_runtime.TypeRecord1, UnsafePtr: unsafe.Pointer(&clone)}
	}()
}

func Call_Main_update1(o_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var o_0 gopurs_runtime.Value = o_0_loop
	_ = o_0
	return func() gopurs_runtime.Value {
		origVal := o_0
		if origVal.Type != gopurs_runtime.TypeRecord1 {
			return gopurs_runtime.RecordUpdateDict(origVal, []string{"foo"}, []gopurs_runtime.Value{gopurs_runtime.Str("Foo")})
		}
		clone := *((*gopurs_runtime.RecordData1)(origVal.UnsafePtr))
		clone.V0 = gopurs_runtime.Str("Foo")
		return gopurs_runtime.Value{Type: gopurs_runtime.TypeRecord1, UnsafePtr: unsafe.Pointer(&clone)}
	}()
}

func Call_Main_replace(o_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var o_0 gopurs_runtime.Value = o_0_loop
	_ = o_0
	var __t0 gopurs_runtime.Value
	{
		if (gopurs_runtime.RecordGet(o_0, "foo").StrVal()) == ("Foo") {
			__t0 = func() gopurs_runtime.Value {
				origVal := o_0
				if origVal.Type != gopurs_runtime.TypeRecord2 {
					return gopurs_runtime.RecordUpdateDict(origVal, []string{"foo"}, []gopurs_runtime.Value{gopurs_runtime.Str("Bar")})
				}
				clone := *((*gopurs_runtime.RecordData2)(origVal.UnsafePtr))
				clone.V1 = gopurs_runtime.Str("Bar")
				return gopurs_runtime.Value{Type: gopurs_runtime.TypeRecord2, UnsafePtr: unsafe.Pointer(&clone)}
			}()
			goto end_branch_0
		} else {

		}
	}
	{
		if (gopurs_runtime.RecordGet(o_0, "foo").StrVal()) == ("Bar") {
			__t0 = func() gopurs_runtime.Value {
				origVal := o_0
				if origVal.Type != gopurs_runtime.TypeRecord2 {
					return gopurs_runtime.RecordUpdateDict(origVal, []string{"bar"}, []gopurs_runtime.Value{gopurs_runtime.Str("Baz")})
				}
				clone := *((*gopurs_runtime.RecordData2)(origVal.UnsafePtr))
				clone.V0 = gopurs_runtime.Str("Baz")
				return gopurs_runtime.Value{Type: gopurs_runtime.TypeRecord2, UnsafePtr: unsafe.Pointer(&clone)}
			}()
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
	return func() gopurs_runtime.Value {
		origVal := o_0
		if origVal.Type != gopurs_runtime.TypeRecord1 {
			return gopurs_runtime.RecordUpdateDict(origVal, []string{"foo"}, []gopurs_runtime.Value{gopurs_runtime.Str("Foo")})
		}
		clone := *((*gopurs_runtime.RecordData1)(origVal.UnsafePtr))
		clone.V0 = gopurs_runtime.Str("Foo")
		return gopurs_runtime.Value{Type: gopurs_runtime.TypeRecord1, UnsafePtr: unsafe.Pointer(&clone)}
	}()
}

func Call_Main_inferPolyUpdate(o_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var o_0 gopurs_runtime.Value = o_0_loop
	_ = o_0
	return func() gopurs_runtime.Value {
		origVal := o_0
		if origVal.Type != gopurs_runtime.TypeRecord1 {
			return gopurs_runtime.RecordUpdateDict(origVal, []string{"foo"}, []gopurs_runtime.Value{gopurs_runtime.Str("Foo")})
		}
		clone := *((*gopurs_runtime.RecordData1)(origVal.UnsafePtr))
		clone.V0 = gopurs_runtime.Str("Foo")
		return gopurs_runtime.Value{Type: gopurs_runtime.TypeRecord1, UnsafePtr: unsafe.Pointer(&clone)}
	}()
}
