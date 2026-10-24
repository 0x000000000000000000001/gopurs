package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_RecordNewtype gopurs_runtime.Value
var once_Main_RecordNewtype sync.Once

func Get_Main_RecordNewtype() gopurs_runtime.Value {
	once_Main_RecordNewtype.Do(func() {
		cache_Main_RecordNewtype = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				orig := Call_Main_RecordNewtype(func() struct {
					x string
					/*gopurs:record-v1:1:0078*/
				} {
					orig := x_0_box
					_ = orig
					clone := struct {
						x string
						/*gopurs:record-v1:1:0078*/
					}{}
					clone.x = gopurs_runtime.RecordGet(orig, "x").StrVal()
					return clone
				}())
				_ = orig
				return gopurs_runtime.RecordDict1("x", gopurs_runtime.Str(orig.x))
			}()
		})
	})
	return cache_Main_RecordNewtype
}

var cache_Main_wrap gopurs_runtime.Value
var once_Main_wrap sync.Once

func Get_Main_wrap() gopurs_runtime.Value {
	once_Main_wrap.Do(func() {
		cache_Main_wrap = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_wrap(gopurs_runtime.CoerceToStruct[Constructor_Main_OldStyleNewtype[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Main_wrap
}

var cache_Main_unwrap gopurs_runtime.Value
var once_Main_unwrap sync.Once

func Get_Main_unwrap() gopurs_runtime.Value {
	once_Main_unwrap.Do(func() {
		cache_Main_unwrap = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_unwrap(gopurs_runtime.CoerceToStruct[Constructor_Main_OldStyleNewtype[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Main_unwrap
}

var cache_Main_refl gopurs_runtime.Value
var once_Main_refl sync.Once

func Get_Main_refl() gopurs_runtime.Value {
	once_Main_refl.Do(func() {
		cache_Main_refl = gopurs_runtime.Value{Type: 9, IntVal: 2698051417, UnsafePtr: unsafe.Pointer((&Constructor_Main_TypeEquals[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return x_0
		}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return x_0
		})}))}
	})
	return cache_Main_refl
}

var cache_Main_coerceBack gopurs_runtime.Value
var once_Main_coerceBack sync.Once

func Get_Main_coerceBack() gopurs_runtime.Value {
	once_Main_coerceBack.Do(func() {
		cache_Main_coerceBack = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_coerceBack(gopurs_runtime.CoerceToStruct[Constructor_Main_TypeEquals[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Main_coerceBack
}

var cache_Main_coerce gopurs_runtime.Value
var once_Main_coerce sync.Once

func Get_Main_coerce() gopurs_runtime.Value {
	once_Main_coerce.Do(func() {
		cache_Main_coerce = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_coerce(gopurs_runtime.CoerceToStruct[Constructor_Main_TypeEquals[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Main_coerce
}

var cache_Main_newtypeRecordNewtype gopurs_runtime.Value
var once_Main_newtypeRecordNewtype sync.Once

func Get_Main_newtypeRecordNewtype() gopurs_runtime.Value {
	once_Main_newtypeRecordNewtype.Do(func() {
		cache_Main_newtypeRecordNewtype = gopurs_runtime.Func(func(dictTypeEquals_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_newtypeRecordNewtype(dictTypeEquals_0_box)
		})
	})
	return cache_Main_newtypeRecordNewtype
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Main_newtypeRecordNewtype(Get_Main_refl()), "unwrap"), func() gopurs_runtime.Value {
			orig := struct {
				x string
				/*gopurs:record-v1:1:0078*/
			}{"Done"}
			_ = orig
			return gopurs_runtime.RecordDict1("x", gopurs_runtime.Str(orig.x))
		}()), "x").StrVal()))
	})
	return cache_Main_main
}

type Constructor_Main_TypeEquals[T_a any, T_b any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}

//gopurs:init-v1:00670065006e006500720061007400650064002f004d00610069006e:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f004d00610069006e005f00540079007000650045007100750061006c0073
func init() {
	gopurs_runtime.StructGetters[2698051417] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_TypeEquals[gopurs_runtime.Value, gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "coerce":
			return gopurs_runtime.Box(c.V0)
		case "coerceBack":
			return gopurs_runtime.Box(c.V1)
		default:
			panic("Key not found in dictionary Constructor_Main_TypeEquals: " + key)
		}
	}
}

type Constructor_Main_OldStyleNewtype[T_t any, T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}

//gopurs:init-v1:00670065006e006500720061007400650064002f004d00610069006e:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f004d00610069006e005f004f006c0064005300740079006c0065004e006500770074007900700065
func init() {
	gopurs_runtime.StructGetters[847837002] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_OldStyleNewtype[gopurs_runtime.Value, gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "unwrap":
			return gopurs_runtime.Box(c.V0)
		case "wrap":
			return gopurs_runtime.Box(c.V1)
		default:
			panic("Key not found in dictionary Constructor_Main_OldStyleNewtype: " + key)
		}
	}
}

func Call_Main_RecordNewtype(x_0_loop struct {
	x string
	/*gopurs:record-v1:1:0078*/
}) struct {
	x string
	/*gopurs:record-v1:1:0078*/
} {
	var x_0 struct {
		x string
		/*gopurs:record-v1:1:0078*/
	} = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_wrap(dict_0_loop *Constructor_Main_OldStyleNewtype[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_OldStyleNewtype[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V1
}

func Call_Main_unwrap(dict_0_loop *Constructor_Main_OldStyleNewtype[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_OldStyleNewtype[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V0
}

func Call_Main_coerceBack(dict_0_loop *Constructor_Main_TypeEquals[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_TypeEquals[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V1
}

func Call_Main_coerce(dict_0_loop *Constructor_Main_TypeEquals[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_TypeEquals[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V0
}

func Call_Main_newtypeRecordNewtype(dictTypeEquals_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictTypeEquals_0 gopurs_runtime.Value = dictTypeEquals_0_loop
	_ = dictTypeEquals_0
	return gopurs_runtime.Value{Type: 9, IntVal: 847837002, UnsafePtr: unsafe.Pointer(Rebox_Main_2718996007_2559642264((&Constructor_Main_OldStyleNewtype[struct {
		x string
		/*gopurs:record-v1:1:0078*/
	}, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTypeEquals_0, "coerceBack"), func() gopurs_runtime.Value {
			orig := func() struct {
				x string
				/*gopurs:record-v1:1:0078*/
			} {
				orig := v_1
				_ = orig
				clone := struct {
					x string
					/*gopurs:record-v1:1:0078*/
				}{}
				clone.x = gopurs_runtime.RecordGet(orig, "x").StrVal()
				return clone
			}()
			_ = orig
			return gopurs_runtime.RecordDict1("x", gopurs_runtime.Str(orig.x))
		}())
	}), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return x_1
	}), Call_Main_coerce(gopurs_runtime.CoerceToStruct[Constructor_Main_TypeEquals[gopurs_runtime.Value, gopurs_runtime.Value]](dictTypeEquals_0)))})))}
}

func Rebox_Main_2718996007_2559642264(in *Constructor_Main_OldStyleNewtype[struct {
	x string
	/*gopurs:record-v1:1:0078*/
}, gopurs_runtime.Value]) *Constructor_Main_OldStyleNewtype[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Main_OldStyleNewtype[gopurs_runtime.Value, gopurs_runtime.Value])(unsafe.Pointer(in))
}
