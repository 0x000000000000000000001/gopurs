package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_WithArgEmpty_dollar_Dict gopurs_runtime.Value
var once_Main_WithArgEmpty_dollar_Dict sync.Once

func Get_Main_WithArgEmpty_dollar_Dict() gopurs_runtime.Value {
	once_Main_WithArgEmpty_dollar_Dict.Do(func() {
		cache_Main_WithArgEmpty_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_WithArgEmpty_dollar_Dict(func() struct {
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
				}{}

				return clone
			}())
		})
	})
	return cache_Main_WithArgEmpty_dollar_Dict
}

var cache_Main_WithArgEmpty_dollar_Dict__3150569997 gopurs_runtime.Value
var once_Main_WithArgEmpty_dollar_Dict__3150569997 sync.Once

func Get_Main_WithArgEmpty_dollar_Dict__3150569997() gopurs_runtime.Value {
	once_Main_WithArgEmpty_dollar_Dict__3150569997.Do(func() {
		cache_Main_WithArgEmpty_dollar_Dict__3150569997 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_WithArgEmpty_dollar_Dict__3150569997(func() struct {
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
				}{}

				return clone
			}())
		})
	})
	return cache_Main_WithArgEmpty_dollar_Dict__3150569997
}

var cache_Main_WithArgHasEmptySuper_dollar_Dict gopurs_runtime.Value
var once_Main_WithArgHasEmptySuper_dollar_Dict sync.Once

func Get_Main_WithArgHasEmptySuper_dollar_Dict() gopurs_runtime.Value {
	once_Main_WithArgHasEmptySuper_dollar_Dict.Do(func() {
		cache_Main_WithArgHasEmptySuper_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 785089286, UnsafePtr: unsafe.Pointer(Call_Main_WithArgHasEmptySuper_dollar_Dict(func() struct {
				WithArgEmpty0        gopurs_runtime.Value
				withArgHasEmptySuper gopurs_runtime.Value
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					WithArgEmpty0        gopurs_runtime.Value
					withArgHasEmptySuper gopurs_runtime.Value
				}{}
				clone.WithArgEmpty0 = gopurs_runtime.RecordGet(orig, "WithArgEmpty0")
				clone.withArgHasEmptySuper = gopurs_runtime.RecordGet(orig, "withArgHasEmptySuper")
				return clone
			}()))}
		})
	})
	return cache_Main_WithArgHasEmptySuper_dollar_Dict
}

var cache_Main_WithArgHasEmptySuper_dollar_Dict__2442956343 gopurs_runtime.Value
var once_Main_WithArgHasEmptySuper_dollar_Dict__2442956343 sync.Once

func Get_Main_WithArgHasEmptySuper_dollar_Dict__2442956343() gopurs_runtime.Value {
	once_Main_WithArgHasEmptySuper_dollar_Dict__2442956343.Do(func() {
		cache_Main_WithArgHasEmptySuper_dollar_Dict__2442956343 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 785089286, UnsafePtr: unsafe.Pointer(Rebox_Main_761192255_1292520722(Call_Main_WithArgHasEmptySuper_dollar_Dict__2442956343(func() struct {
				WithArgEmpty0        gopurs_runtime.Value
				withArgHasEmptySuper uint32
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					WithArgEmpty0        gopurs_runtime.Value
					withArgHasEmptySuper uint32
				}{}
				clone.WithArgEmpty0 = gopurs_runtime.RecordGet(orig, "WithArgEmpty0")
				clone.withArgHasEmptySuper = uint32(gopurs_runtime.RecordGet(orig, "withArgHasEmptySuper").IntVal)
				return clone
			}())))}
		})
	})
	return cache_Main_WithArgHasEmptySuper_dollar_Dict__2442956343
}

var cache_Main_EmptyClass_dollar_Dict gopurs_runtime.Value
var once_Main_EmptyClass_dollar_Dict sync.Once

func Get_Main_EmptyClass_dollar_Dict() gopurs_runtime.Value {
	once_Main_EmptyClass_dollar_Dict.Do(func() {
		cache_Main_EmptyClass_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_EmptyClass_dollar_Dict(func() struct {
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
				}{}

				return clone
			}())
		})
	})
	return cache_Main_EmptyClass_dollar_Dict
}

var cache_Main_Check gopurs_runtime.Value
var once_Main_Check sync.Once

func Get_Main_Check() gopurs_runtime.Value {
	once_Main_Check.Do(func() {
		cache_Main_Check = gopurs_runtime.Value{Type: 9, IntVal: int64(60647608), UnsafePtr: nil}
	})
	return cache_Main_Check
}

var cache_Main_HasEmptySuper_dollar_Dict gopurs_runtime.Value
var once_Main_HasEmptySuper_dollar_Dict sync.Once

func Get_Main_HasEmptySuper_dollar_Dict() gopurs_runtime.Value {
	once_Main_HasEmptySuper_dollar_Dict.Do(func() {
		cache_Main_HasEmptySuper_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2786347248, UnsafePtr: unsafe.Pointer(Call_Main_HasEmptySuper_dollar_Dict(func() struct {
				EmptyClass0   gopurs_runtime.Value
				hasEmptySuper uint32
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					EmptyClass0   gopurs_runtime.Value
					hasEmptySuper uint32
				}{}
				clone.EmptyClass0 = gopurs_runtime.RecordGet(orig, "EmptyClass0")
				clone.hasEmptySuper = uint32(gopurs_runtime.RecordGet(orig, "hasEmptySuper").IntVal)
				return clone
			}()))}
		})
	})
	return cache_Main_HasEmptySuper_dollar_Dict
}

var cache_Main_HasNonEmptySuper_dollar_Dict gopurs_runtime.Value
var once_Main_HasNonEmptySuper_dollar_Dict sync.Once

func Get_Main_HasNonEmptySuper_dollar_Dict() gopurs_runtime.Value {
	once_Main_HasNonEmptySuper_dollar_Dict.Do(func() {
		cache_Main_HasNonEmptySuper_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3506164383, UnsafePtr: unsafe.Pointer(Call_Main_HasNonEmptySuper_dollar_Dict(func() struct {
				HasEmptySuper0 gopurs_runtime.Value
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					HasEmptySuper0 gopurs_runtime.Value
				}{}
				clone.HasEmptySuper0 = gopurs_runtime.RecordGet(orig, "HasEmptySuper0")
				return clone
			}()))}
		})
	})
	return cache_Main_HasNonEmptySuper_dollar_Dict
}

var cache_Main_AliasEmptyClass_dollar_Dict gopurs_runtime.Value
var once_Main_AliasEmptyClass_dollar_Dict sync.Once

func Get_Main_AliasEmptyClass_dollar_Dict() gopurs_runtime.Value {
	once_Main_AliasEmptyClass_dollar_Dict.Do(func() {
		cache_Main_AliasEmptyClass_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4080323731, UnsafePtr: unsafe.Pointer(Call_Main_AliasEmptyClass_dollar_Dict(func() struct {
				EmptyClass0 gopurs_runtime.Value
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					EmptyClass0 gopurs_runtime.Value
				}{}
				clone.EmptyClass0 = gopurs_runtime.RecordGet(orig, "EmptyClass0")
				return clone
			}()))}
		})
	})
	return cache_Main_AliasEmptyClass_dollar_Dict
}

var cache_Main_withArgHasEmptySuper gopurs_runtime.Value
var once_Main_withArgHasEmptySuper sync.Once

func Get_Main_withArgHasEmptySuper() gopurs_runtime.Value {
	once_Main_withArgHasEmptySuper.Do(func() {
		cache_Main_withArgHasEmptySuper = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_withArgHasEmptySuper(dict_0_box)
		})
	})
	return cache_Main_withArgHasEmptySuper
}

var cache_Main_withArgEmptyCheck gopurs_runtime.Value
var once_Main_withArgEmptyCheck sync.Once

func Get_Main_withArgEmptyCheck() gopurs_runtime.Value {
	once_Main_withArgEmptyCheck.Do(func() {
		cache_Main_withArgEmptyCheck = func() gopurs_runtime.Value {
			orig := struct {
			}{}
			_ = orig
			return gopurs_runtime.RecordDict0()
		}()
	})
	return cache_Main_withArgEmptyCheck
}

var cache_Main_withArgHasEmptySuperCheck gopurs_runtime.Value
var once_Main_withArgHasEmptySuperCheck sync.Once

func Get_Main_withArgHasEmptySuperCheck() gopurs_runtime.Value {
	once_Main_withArgHasEmptySuperCheck.Do(func() {
		cache_Main_withArgHasEmptySuperCheck = gopurs_runtime.Value{Type: 9, IntVal: 785089286, UnsafePtr: unsafe.Pointer(Rebox_Main_761192255_1292520722((&Constructor_Main_WithArgHasEmptySuper[uint32]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{}
		}), 60647608})))}
	})
	return cache_Main_withArgHasEmptySuperCheck
}

var cache_Main_whenAccessingSuperDict gopurs_runtime.Value
var once_Main_whenAccessingSuperDict sync.Once

func Get_Main_whenAccessingSuperDict() gopurs_runtime.Value {
	once_Main_whenAccessingSuperDict.Do(func() {
		cache_Main_whenAccessingSuperDict = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Value{Type: 9, IntVal: int64(60647608), UnsafePtr: nil}.IntVal)), UnsafePtr: nil}
	})
	return cache_Main_whenAccessingSuperDict
}

var cache_Main_hasNonEmptySuperInst gopurs_runtime.Value
var once_Main_hasNonEmptySuperInst sync.Once

func Get_Main_hasNonEmptySuperInst() gopurs_runtime.Value {
	once_Main_hasNonEmptySuperInst.Do(func() {
		cache_Main_hasNonEmptySuperInst = gopurs_runtime.Func(func(dictHasEmptySuper_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_hasNonEmptySuperInst(dictHasEmptySuper_0_box)
		})
	})
	return cache_Main_hasNonEmptySuperInst
}

var cache_Main_hasEmptySuper gopurs_runtime.Value
var once_Main_hasEmptySuper sync.Once

func Get_Main_hasEmptySuper() gopurs_runtime.Value {
	once_Main_hasEmptySuper.Do(func() {
		cache_Main_hasEmptySuper = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_hasEmptySuper(dict_0_box)
		})
	})
	return cache_Main_hasEmptySuper
}

var cache_Main_eqCheck gopurs_runtime.Value
var once_Main_eqCheck sync.Once

func Get_Main_eqCheck() gopurs_runtime.Value {
	once_Main_eqCheck.Do(func() {
		cache_Main_eqCheck = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_3768443459_3790796878((&Constructor_Data_Eq_Eq[uint32]{1, gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(true)
		})})))}
	})
	return cache_Main_eqCheck
}

var cache_Main_emptyDictInst gopurs_runtime.Value
var once_Main_emptyDictInst sync.Once

func Get_Main_emptyDictInst() gopurs_runtime.Value {
	once_Main_emptyDictInst.Do(func() {
		cache_Main_emptyDictInst = func() gopurs_runtime.Value {
			orig := struct {
			}{}
			_ = orig
			return gopurs_runtime.RecordDict0()
		}()
	})
	return cache_Main_emptyDictInst
}

var cache_Main_hasEmptySuperInst gopurs_runtime.Value
var once_Main_hasEmptySuperInst sync.Once

func Get_Main_hasEmptySuperInst() gopurs_runtime.Value {
	once_Main_hasEmptySuperInst.Do(func() {
		cache_Main_hasEmptySuperInst = gopurs_runtime.Value{Type: 9, IntVal: 2786347248, UnsafePtr: unsafe.Pointer((&Constructor_Main_HasEmptySuper{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{}
		}), 60647608}))}
	})
	return cache_Main_hasEmptySuperInst
}

var cache_Main_whenHasEmptySuper gopurs_runtime.Value
var once_Main_whenHasEmptySuper sync.Once

func Get_Main_whenHasEmptySuper() gopurs_runtime.Value {
	once_Main_whenHasEmptySuper.Do(func() {
		cache_Main_whenHasEmptySuper = gopurs_runtime.Value{Type: 9, IntVal: int64(60647608), UnsafePtr: nil}
	})
	return cache_Main_whenHasEmptySuper
}

var cache_Main_whenHasNonEmptySuper gopurs_runtime.Value
var once_Main_whenHasNonEmptySuper sync.Once

func Get_Main_whenHasNonEmptySuper() gopurs_runtime.Value {
	once_Main_whenHasNonEmptySuper.Do(func() {
		cache_Main_whenHasNonEmptySuper = gopurs_runtime.Value{Type: 9, IntVal: int64(60647608), UnsafePtr: nil}
	})
	return cache_Main_whenHasNonEmptySuper
}

var cache_Main_whenEmpty gopurs_runtime.Value
var once_Main_whenEmpty sync.Once

func Get_Main_whenEmpty() gopurs_runtime.Value {
	once_Main_whenEmpty.Do(func() {
		cache_Main_whenEmpty = gopurs_runtime.Value{Type: 9, IntVal: int64(60647608), UnsafePtr: nil}
	})
	return cache_Main_whenEmpty
}

var cache_Main_aliasEmptyClassInst gopurs_runtime.Value
var once_Main_aliasEmptyClassInst sync.Once

func Get_Main_aliasEmptyClassInst() gopurs_runtime.Value {
	once_Main_aliasEmptyClassInst.Do(func() {
		cache_Main_aliasEmptyClassInst = gopurs_runtime.Value{Type: 9, IntVal: 4080323731, UnsafePtr: unsafe.Pointer((&Constructor_Main_AliasEmptyClass{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{}
		})}))}
	})
	return cache_Main_aliasEmptyClassInst
}

var cache_Main_whenAliasEmptyClass gopurs_runtime.Value
var once_Main_whenAliasEmptyClass sync.Once

func Get_Main_whenAliasEmptyClass() gopurs_runtime.Value {
	once_Main_whenAliasEmptyClass.Do(func() {
		cache_Main_whenAliasEmptyClass = gopurs_runtime.Value{Type: 9, IntVal: int64(60647608), UnsafePtr: nil}
	})
	return cache_Main_whenAliasEmptyClass
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

type Constructor_Main_Check struct {
	Rc uint32
}

type Constructor_Main_WithArgEmpty[T_t any] struct {
	Rc uint32
}

func init() {
	gopurs_runtime.StructGetters[207373949] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_WithArgEmpty[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {

		default:
			panic("Key not found in dictionary Constructor_Main_WithArgEmpty: " + key)
		}
	}
}

type Constructor_Main_WithArgHasEmptySuper[T_t any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 T_t
}

func init() {
	gopurs_runtime.StructGetters[785089286] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_WithArgHasEmptySuper[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "WithArgEmpty0":
			return gopurs_runtime.Box(c.V0)
		case "withArgHasEmptySuper":
			return gopurs_runtime.Box(c.V1)
		default:
			panic("Key not found in dictionary Constructor_Main_WithArgHasEmptySuper: " + key)
		}
	}
}

type Constructor_Main_EmptyClass struct {
	Rc uint32
}

func init() {
	gopurs_runtime.StructGetters[394228773] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_EmptyClass)(ptr)
		_ = c
		switch key {

		default:
			panic("Key not found in dictionary Constructor_Main_EmptyClass: " + key)
		}
	}
}

type Constructor_Main_HasEmptySuper struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 uint32
}

func init() {
	gopurs_runtime.StructGetters[2786347248] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_HasEmptySuper)(ptr)
		_ = c
		switch key {
		case "EmptyClass0":
			return gopurs_runtime.Box(c.V0)
		case "hasEmptySuper":
			return gopurs_runtime.Value{Type: 9, IntVal: int64(c.V1)}
		default:
			panic("Key not found in dictionary Constructor_Main_HasEmptySuper: " + key)
		}
	}
}

type Constructor_Main_HasNonEmptySuper struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

func init() {
	gopurs_runtime.StructGetters[3506164383] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_HasNonEmptySuper)(ptr)
		_ = c
		switch key {
		case "HasEmptySuper0":
			return gopurs_runtime.Box(c.V0)
		default:
			panic("Key not found in dictionary Constructor_Main_HasNonEmptySuper: " + key)
		}
	}
}

type Constructor_Main_AliasEmptyClass struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

func init() {
	gopurs_runtime.StructGetters[4080323731] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_AliasEmptyClass)(ptr)
		_ = c
		switch key {
		case "EmptyClass0":
			return gopurs_runtime.Box(c.V0)
		default:
			panic("Key not found in dictionary Constructor_Main_AliasEmptyClass: " + key)
		}
	}
}

func Call_Main_WithArgEmpty_dollar_Dict(x_0_loop struct {
}) gopurs_runtime.Value {
	var x_0 struct {
	} = x_0_loop
	_ = x_0
	return func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict0()
	}()
}

func Call_Main_WithArgEmpty_dollar_Dict__3150569997(x_0_loop struct {
}) gopurs_runtime.Value {
WithArgEmpty_dollar_Dict__3150569997:
	for {
		if false {
			continue WithArgEmpty_dollar_Dict__3150569997
		}
		var x_0 struct {
		} = x_0_loop
		_ = x_0
		return func() gopurs_runtime.Value {
			orig := x_0
			_ = orig
			return gopurs_runtime.RecordDict0()
		}()
	}
}

func Call_Main_WithArgHasEmptySuper_dollar_Dict(x_0_loop struct {
	WithArgEmpty0        gopurs_runtime.Value
	withArgHasEmptySuper gopurs_runtime.Value
}) *Constructor_Main_WithArgHasEmptySuper[gopurs_runtime.Value] {
	var x_0 struct {
		WithArgEmpty0        gopurs_runtime.Value
		withArgHasEmptySuper gopurs_runtime.Value
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Main_WithArgHasEmptySuper[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict2("WithArgEmpty0", "withArgHasEmptySuper", orig.WithArgEmpty0, orig.withArgHasEmptySuper)
	}())
}

func Call_Main_WithArgHasEmptySuper_dollar_Dict__2442956343(x_0_loop struct {
	WithArgEmpty0        gopurs_runtime.Value
	withArgHasEmptySuper uint32
}) *Constructor_Main_WithArgHasEmptySuper[uint32] {
WithArgHasEmptySuper_dollar_Dict__2442956343:
	for {
		if false {
			continue WithArgHasEmptySuper_dollar_Dict__2442956343
		}
		var x_0 struct {
			WithArgEmpty0        gopurs_runtime.Value
			withArgHasEmptySuper uint32
		} = x_0_loop
		_ = x_0
		return gopurs_runtime.CoerceToStruct[Constructor_Main_WithArgHasEmptySuper[uint32]](func() gopurs_runtime.Value {
			orig := x_0
			_ = orig
			return gopurs_runtime.RecordDict2("WithArgEmpty0", "withArgHasEmptySuper", orig.WithArgEmpty0, gopurs_runtime.Value{Type: 9, IntVal: int64(orig.withArgHasEmptySuper), UnsafePtr: nil})
		}())
	}
}

func Call_Main_EmptyClass_dollar_Dict(x_0_loop struct {
}) gopurs_runtime.Value {
	var x_0 struct {
	} = x_0_loop
	_ = x_0
	return func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict0()
	}()
}

func Call_Main_HasEmptySuper_dollar_Dict(x_0_loop struct {
	EmptyClass0   gopurs_runtime.Value
	hasEmptySuper uint32
}) *Constructor_Main_HasEmptySuper {
	var x_0 struct {
		EmptyClass0   gopurs_runtime.Value
		hasEmptySuper uint32
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Main_HasEmptySuper](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict2("EmptyClass0", "hasEmptySuper", orig.EmptyClass0, gopurs_runtime.Value{Type: 9, IntVal: int64(orig.hasEmptySuper), UnsafePtr: nil})
	}())
}

func Call_Main_HasNonEmptySuper_dollar_Dict(x_0_loop struct {
	HasEmptySuper0 gopurs_runtime.Value
}) *Constructor_Main_HasNonEmptySuper {
	var x_0 struct {
		HasEmptySuper0 gopurs_runtime.Value
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Main_HasNonEmptySuper](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict1("HasEmptySuper0", orig.HasEmptySuper0)
	}())
}

func Call_Main_AliasEmptyClass_dollar_Dict(x_0_loop struct {
	EmptyClass0 gopurs_runtime.Value
}) *Constructor_Main_AliasEmptyClass {
	var x_0 struct {
		EmptyClass0 gopurs_runtime.Value
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Main_AliasEmptyClass](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict1("EmptyClass0", orig.EmptyClass0)
	}())
}

func Call_Main_withArgHasEmptySuper(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dict_0 gopurs_runtime.Value = dict_0_loop
	_ = dict_0
	return gopurs_runtime.RecordGet(dict_0, "withArgHasEmptySuper")
}

func Call_Main_hasNonEmptySuperInst(dictHasEmptySuper_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictHasEmptySuper_0 gopurs_runtime.Value = dictHasEmptySuper_0_loop
	_ = dictHasEmptySuper_0
	return gopurs_runtime.Value{Type: 9, IntVal: 3506164383, UnsafePtr: unsafe.Pointer((&Constructor_Main_HasNonEmptySuper{1, gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 2786347248, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Main_HasEmptySuper](dictHasEmptySuper_0))}
	})}))}
}

func Call_Main_hasEmptySuper(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dict_0 gopurs_runtime.Value = dict_0_loop
	_ = dict_0
	return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.RecordGet(dict_0, "hasEmptySuper").IntVal)), UnsafePtr: nil}
}

func Rebox_Main_3768443459_3790796878(in *Constructor_Data_Eq_Eq[uint32]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_761192255_1292520722(in *Constructor_Main_WithArgHasEmptySuper[uint32]) *Constructor_Main_WithArgHasEmptySuper[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_WithArgHasEmptySuper[gopurs_runtime.Value]{}
	out.V0 = in.V0
	out.V1 = gopurs_runtime.Value{Type: 9, IntVal: int64(in.V1), UnsafePtr: nil}
	return out
}
