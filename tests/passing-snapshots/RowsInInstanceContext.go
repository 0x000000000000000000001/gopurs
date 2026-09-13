package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_identity gopurs_runtime.Value
var once_Main_identity sync.Once

func Get_Main_identity() gopurs_runtime.Value {
	once_Main_identity.Do(func() {
		cache_Main_identity = Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))})
	})
	return cache_Main_identity
}

var cache_Main_TypeEquals_dollar_Dict gopurs_runtime.Value
var once_Main_TypeEquals_dollar_Dict sync.Once

func Get_Main_TypeEquals_dollar_Dict() gopurs_runtime.Value {
	once_Main_TypeEquals_dollar_Dict.Do(func() {
		cache_Main_TypeEquals_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2698051417, UnsafePtr: unsafe.Pointer(Call_Main_TypeEquals_dollar_Dict(func() struct {
				coerce     gopurs_runtime.Value
				coerceBack gopurs_runtime.Value
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					coerce     gopurs_runtime.Value
					coerceBack gopurs_runtime.Value
				}{}
				clone.coerce = gopurs_runtime.RecordGet(orig, "coerce")
				clone.coerceBack = gopurs_runtime.RecordGet(orig, "coerceBack")
				return clone
			}()))}
		})
	})
	return cache_Main_TypeEquals_dollar_Dict
}

var cache_Main_RecordNewtype gopurs_runtime.Value
var once_Main_RecordNewtype sync.Once

func Get_Main_RecordNewtype() gopurs_runtime.Value {
	once_Main_RecordNewtype.Do(func() {
		cache_Main_RecordNewtype = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				orig := Call_Main_RecordNewtype(func() struct {
					x string
				} {
					orig := x_0_box
					_ = orig
					clone := struct {
						x string
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

var cache_Main_OldStyleNewtype_dollar_Dict gopurs_runtime.Value
var once_Main_OldStyleNewtype_dollar_Dict sync.Once

func Get_Main_OldStyleNewtype_dollar_Dict() gopurs_runtime.Value {
	once_Main_OldStyleNewtype_dollar_Dict.Do(func() {
		cache_Main_OldStyleNewtype_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 847837002, UnsafePtr: unsafe.Pointer(Call_Main_OldStyleNewtype_dollar_Dict(func() struct {
				unwrap gopurs_runtime.Value
				wrap   gopurs_runtime.Value
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					unwrap gopurs_runtime.Value
					wrap   gopurs_runtime.Value
				}{}
				clone.unwrap = gopurs_runtime.RecordGet(orig, "unwrap")
				clone.wrap = gopurs_runtime.RecordGet(orig, "wrap")
				return clone
			}()))}
		})
	})
	return cache_Main_OldStyleNewtype_dollar_Dict
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
		cache_Main_refl = gopurs_runtime.Value{Type: 9, IntVal: 2698051417, UnsafePtr: unsafe.Pointer((&Constructor_Main_TypeEquals[gopurs_runtime.Value, gopurs_runtime.Value]{1, Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))}), Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))})}))}
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

func Call_Main_TypeEquals_dollar_Dict(x_0_loop struct {
	coerce     gopurs_runtime.Value
	coerceBack gopurs_runtime.Value
}) *Constructor_Main_TypeEquals[gopurs_runtime.Value, gopurs_runtime.Value] {
	var x_0 struct {
		coerce     gopurs_runtime.Value
		coerceBack gopurs_runtime.Value
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Main_TypeEquals[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict2("coerce", "coerceBack", orig.coerce, orig.coerceBack)
	}())
}

func Call_Main_RecordNewtype(x_0_loop struct {
	x string
}) struct {
	x string
} {
	var x_0 struct {
		x string
	} = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_OldStyleNewtype_dollar_Dict(x_0_loop struct {
	unwrap gopurs_runtime.Value
	wrap   gopurs_runtime.Value
}) *Constructor_Main_OldStyleNewtype[gopurs_runtime.Value, gopurs_runtime.Value] {
	var x_0 struct {
		unwrap gopurs_runtime.Value
		wrap   gopurs_runtime.Value
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Main_OldStyleNewtype[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict2("unwrap", "wrap", orig.unwrap, orig.wrap)
	}())
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
	}, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTypeEquals_0, "coerceBack"), func() gopurs_runtime.Value {
			orig := func() struct {
				x string
			} {
				orig := v_1
				_ = orig
				clone := struct {
					x string
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
}, gopurs_runtime.Value]) *Constructor_Main_OldStyleNewtype[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_OldStyleNewtype[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = in.V0
	out.V1 = in.V1
	return out
}
