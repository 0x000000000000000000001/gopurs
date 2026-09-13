package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_Test_dollar_Dict gopurs_runtime.Value
var once_Main_Test_dollar_Dict sync.Once

func Get_Main_Test_dollar_Dict() gopurs_runtime.Value {
	once_Main_Test_dollar_Dict.Do(func() {
		cache_Main_Test_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3423238824, UnsafePtr: unsafe.Pointer(Call_Main_Test_dollar_Dict(func() struct {
				fn  gopurs_runtime.Value
				val gopurs_runtime.Value
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					fn  gopurs_runtime.Value
					val gopurs_runtime.Value
				}{}
				clone.fn = gopurs_runtime.RecordGet(orig, "fn")
				clone.val = gopurs_runtime.RecordGet(orig, "val")
				return clone
			}()))}
		})
	})
	return cache_Main_Test_dollar_Dict
}

var cache_Main_Test_dollar_Dict__1170844877 gopurs_runtime.Value
var once_Main_Test_dollar_Dict__1170844877 sync.Once

func Get_Main_Test_dollar_Dict__1170844877() gopurs_runtime.Value {
	once_Main_Test_dollar_Dict__1170844877.Do(func() {
		cache_Main_Test_dollar_Dict__1170844877 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3423238824, UnsafePtr: unsafe.Pointer(Rebox_Main_1531651672_2549914428(Call_Main_Test_dollar_Dict__1170844877(func() struct {
				fn  gopurs_runtime.Value
				val bool
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					fn  gopurs_runtime.Value
					val bool
				}{}
				clone.fn = gopurs_runtime.RecordGet(orig, "fn")
				clone.val = (gopurs_runtime.RecordGet(orig, "val").IntVal) != (0)
				return clone
			}())))}
		})
	})
	return cache_Main_Test_dollar_Dict__1170844877
}

var cache_Main_val gopurs_runtime.Value
var once_Main_val sync.Once

func Get_Main_val() gopurs_runtime.Value {
	once_Main_val.Do(func() {
		cache_Main_val = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_val(dict_0_box)
		})
	})
	return cache_Main_val
}

var cache_Main_testBoolean gopurs_runtime.Value
var once_Main_testBoolean sync.Once

func Get_Main_testBoolean() gopurs_runtime.Value {
	once_Main_testBoolean.Do(func() {
		cache_Main_testBoolean = gopurs_runtime.Value{Type: 9, IntVal: 3423238824, UnsafePtr: unsafe.Pointer(Rebox_Main_1531651672_2549914428((&Constructor_Main_Test[bool]{1, gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool((y_1.IntVal) != (0))
		}), true})))}
	})
	return cache_Main_testBoolean
}

var cache_Main_fn gopurs_runtime.Value
var once_Main_fn sync.Once

func Get_Main_fn() gopurs_runtime.Value {
	once_Main_fn.Do(func() {
		cache_Main_fn = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_fn(gopurs_runtime.CoerceToStruct[Constructor_Main_Test[gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Main_fn
}

var cache_Main_fn__425925499 gopurs_runtime.Value
var once_Main_fn__425925499 sync.Once

func Get_Main_fn__425925499() gopurs_runtime.Value {
	once_Main_fn__425925499.Do(func() {
		cache_Main_fn__425925499 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(Call_Main_fn__425925499((__eta_norm_1_0_box.IntVal) != (0), (__eta_norm_0_unused_1_box.IntVal) != (0)))
		})
	})
	return cache_Main_fn__425925499
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = func() gopurs_runtime.Value {
			var __t0 string
			{
				if (Call_Main_val(gopurs_runtime.Value{Type: 9, IntVal: 3423238824, UnsafePtr: unsafe.Pointer(Rebox_Main_1531651672_2549914428(Rebox_Main_2549914428_1531651672(gopurs_runtime.CoerceToStruct[Constructor_Main_Test[gopurs_runtime.Value]](Get_Main_testBoolean()))))}).IntVal) != (0) {
					__t0 = "true"
					goto end_branch_0
				} else {

				}
			}
			{
				__t0 = "false"
			}
		end_branch_0:
			return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(__t0)), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
			}))
		}()
	})
	return cache_Main_main
}

type Constructor_Main_Test[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 T_a
}

func init() {
	gopurs_runtime.StructGetters[3423238824] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_Test[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "fn":
			return gopurs_runtime.Box(c.V0)
		case "val":
			return gopurs_runtime.Box(c.V1)
		default:
			panic("Key not found in dictionary Constructor_Main_Test: " + key)
		}
	}
}

func Call_Main_Test_dollar_Dict(x_0_loop struct {
	fn  gopurs_runtime.Value
	val gopurs_runtime.Value
}) *Constructor_Main_Test[gopurs_runtime.Value] {
	var x_0 struct {
		fn  gopurs_runtime.Value
		val gopurs_runtime.Value
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Main_Test[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict2("fn", "val", orig.fn, orig.val)
	}())
}

func Call_Main_Test_dollar_Dict__1170844877(x_0_loop struct {
	fn  gopurs_runtime.Value
	val bool
}) *Constructor_Main_Test[bool] {
Test_dollar_Dict__1170844877:
	for {
		if false {
			continue Test_dollar_Dict__1170844877
		}
		var x_0 struct {
			fn  gopurs_runtime.Value
			val bool
		} = x_0_loop
		_ = x_0
		return gopurs_runtime.CoerceToStruct[Constructor_Main_Test[bool]](func() gopurs_runtime.Value {
			orig := x_0
			_ = orig
			return gopurs_runtime.RecordDict2("fn", "val", orig.fn, gopurs_runtime.Bool(orig.val))
		}())
	}
}

func Call_Main_val(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dict_0 gopurs_runtime.Value = dict_0_loop
	_ = dict_0
	return gopurs_runtime.RecordGet(dict_0, "val")
}

func Call_Main_fn(dict_0_loop *Constructor_Main_Test[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_Test[gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V0
}

func Call_Main_fn__425925499(__eta_norm_1_0_loop bool, __eta_norm_0_unused_1_loop bool) bool {
fn__425925499:
	for {
		if false {
			continue fn__425925499
		}
		var __eta_norm_1_0 bool = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_unused_1 bool = __eta_norm_0_unused_1_loop
		_ = __eta_norm_0_unused_1
		return (Call_Main_val(gopurs_runtime.Value{Type: 9, IntVal: 3423238824, UnsafePtr: unsafe.Pointer(Rebox_Main_1531651672_2549914428(Rebox_Main_2549914428_1531651672(gopurs_runtime.CoerceToStruct[Constructor_Main_Test[gopurs_runtime.Value]](Get_Main_testBoolean()))))}).IntVal) != (0)
	}
}

func Rebox_Main_1531651672_2549914428(in *Constructor_Main_Test[bool]) *Constructor_Main_Test[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_Test[gopurs_runtime.Value]{}
	out.V0 = in.V0
	out.V1 = gopurs_runtime.Bool(in.V1)
	return out
}

func Rebox_Main_2549914428_1531651672(in *Constructor_Main_Test[gopurs_runtime.Value]) *Constructor_Main_Test[bool] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_Test[bool]{}
	out.V0 = in.V0
	out.V1 = (in.V1.IntVal) != (0)
	return out
}
