package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_Show_dollar_Dict gopurs_runtime.Value
var once_Main_Show_dollar_Dict sync.Once

func Get_Main_Show_dollar_Dict() gopurs_runtime.Value {
	once_Main_Show_dollar_Dict.Do(func() {
		cache_Main_Show_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3143145725, UnsafePtr: unsafe.Pointer(Call_Main_Show_dollar_Dict(func() struct {
				Show0 gopurs_runtime.Value
				id    gopurs_runtime.Value
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Show0 gopurs_runtime.Value
					id    gopurs_runtime.Value
				}{}
				clone.Show0 = gopurs_runtime.RecordGet(orig, "Show0")
				clone.id = gopurs_runtime.RecordGet(orig, "id")
				return clone
			}()))}
		})
	})
	return cache_Main_Show_dollar_Dict
}

var cache_Main_Show_dollar_Dict__1998906268 gopurs_runtime.Value
var once_Main_Show_dollar_Dict__1998906268 sync.Once

func Get_Main_Show_dollar_Dict__1998906268() gopurs_runtime.Value {
	once_Main_Show_dollar_Dict__1998906268.Do(func() {
		cache_Main_Show_dollar_Dict__1998906268 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3143145725, UnsafePtr: unsafe.Pointer(Rebox_Main_2177390102_590014985(Call_Main_Show_dollar_Dict__1998906268(func() struct {
				Show0 gopurs_runtime.Value
				id    gopurs_runtime.Value
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Show0 gopurs_runtime.Value
					id    gopurs_runtime.Value
				}{}
				clone.Show0 = gopurs_runtime.RecordGet(orig, "Show0")
				clone.id = gopurs_runtime.RecordGet(orig, "id")
				return clone
			}())))}
		})
	})
	return cache_Main_Show_dollar_Dict__1998906268
}

var cache_Main_showString gopurs_runtime.Value
var once_Main_showString sync.Once

func Get_Main_showString() gopurs_runtime.Value {
	once_Main_showString.Do(func() {
		cache_Main_showString = gopurs_runtime.Value{Type: 9, IntVal: 3143145725, UnsafePtr: unsafe.Pointer(Rebox_Main_2177390102_590014985((&Constructor_Main_Show[string]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString()))))}
		}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(x_0.StrVal())
		})})))}
	})
	return cache_Main_showString
}

var cache_Main_id gopurs_runtime.Value
var once_Main_id sync.Once

func Get_Main_id() gopurs_runtime.Value {
	once_Main_id.Do(func() {
		cache_Main_id = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_id(gopurs_runtime.CoerceToStruct[Constructor_Main_Show[gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Main_id
}

var cache_Main_id__3658769210 gopurs_runtime.Value
var once_Main_id__3658769210 sync.Once

func Get_Main_id__3658769210() gopurs_runtime.Value {
	once_Main_id__3658769210.Do(func() {
		cache_Main_id__3658769210 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_id__3658769210(__eta_norm_0_0_box.StrVal()))
		})
	})
	return cache_Main_id__3658769210
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

type Constructor_Main_Show[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}

func init() {
	gopurs_runtime.StructGetters[3143145725] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_Show[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "Show0":
			return gopurs_runtime.Box(c.V0)
		case "id":
			return gopurs_runtime.Box(c.V1)
		default:
			panic("Key not found in dictionary Constructor_Main_Show: " + key)
		}
	}
}

func Call_Main_Show_dollar_Dict(x_0_loop struct {
	Show0 gopurs_runtime.Value
	id    gopurs_runtime.Value
}) *Constructor_Main_Show[gopurs_runtime.Value] {
	var x_0 struct {
		Show0 gopurs_runtime.Value
		id    gopurs_runtime.Value
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Main_Show[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict2("Show0", "id", orig.Show0, orig.id)
	}())
}

func Call_Main_Show_dollar_Dict__1998906268(x_0_loop struct {
	Show0 gopurs_runtime.Value
	id    gopurs_runtime.Value
}) *Constructor_Main_Show[string] {
Show_dollar_Dict__1998906268:
	for {
		if false {
			continue Show_dollar_Dict__1998906268
		}
		var x_0 struct {
			Show0 gopurs_runtime.Value
			id    gopurs_runtime.Value
		} = x_0_loop
		_ = x_0
		return gopurs_runtime.CoerceToStruct[Constructor_Main_Show[string]](func() gopurs_runtime.Value {
			orig := x_0
			_ = orig
			return gopurs_runtime.RecordDict2("Show0", "id", orig.Show0, orig.id)
		}())
	}
}

func Call_Main_id(dict_0_loop *Constructor_Main_Show[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_Show[gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V1
}

func Call_Main_id__3658769210(__eta_norm_0_0_loop string) string {
id__3658769210:
	for {
		if false {
			continue id__3658769210
		}
		var __eta_norm_0_0 string = __eta_norm_0_0_loop
		_ = __eta_norm_0_0
		return __eta_norm_0_0
	}
}

func Rebox_Main_1386611502_1514099793(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[string] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[string]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_1514099793_1386611502(in *Constructor_Data_Show_Show[string]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_2177390102_590014985(in *Constructor_Main_Show[string]) *Constructor_Main_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_Show[gopurs_runtime.Value]{}
	out.V0 = in.V0
	out.V1 = in.V1
	return out
}
