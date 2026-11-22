package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Lazy_Lazy_dollar_Dict gopurs_runtime.Value
var once_Control_Lazy_Lazy_dollar_Dict sync.Once

func Get_Control_Lazy_Lazy_dollar_Dict() gopurs_runtime.Value {
	once_Control_Lazy_Lazy_dollar_Dict.Do(func() {
		cache_Control_Lazy_Lazy_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1860244333, UnsafePtr: unsafe.Pointer(Call_Control_Lazy_Lazy_dollar_Dict(func() struct {
				go__defer gopurs_runtime.Value
				/*gopurs:record-v1:1:00640065006600650072*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					go__defer gopurs_runtime.Value
					/*gopurs:record-v1:1:00640065006600650072*/
				}{}
				clone.go__defer = gopurs_runtime.RecordGet(orig, "defer")
				return clone
			}()))}
		})
	})
	return cache_Control_Lazy_Lazy_dollar_Dict
}

var cache_Control_Lazy_Lazy_dollar_Dict__3034036302 gopurs_runtime.Value
var once_Control_Lazy_Lazy_dollar_Dict__3034036302 sync.Once

func Get_Control_Lazy_Lazy_dollar_Dict__3034036302() gopurs_runtime.Value {
	once_Control_Lazy_Lazy_dollar_Dict__3034036302.Do(func() {
		cache_Control_Lazy_Lazy_dollar_Dict__3034036302 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1860244333, UnsafePtr: unsafe.Pointer(Call_Control_Lazy_Lazy_dollar_Dict__3034036302(func() struct {
				go__defer gopurs_runtime.Value
				/*gopurs:record-v1:1:00640065006600650072*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					go__defer gopurs_runtime.Value
					/*gopurs:record-v1:1:00640065006600650072*/
				}{}
				clone.go__defer = gopurs_runtime.RecordGet(orig, "defer")
				return clone
			}()))}
		})
	})
	return cache_Control_Lazy_Lazy_dollar_Dict__3034036302
}

var cache_Control_Lazy_lazyUnit gopurs_runtime.Value
var once_Control_Lazy_lazyUnit sync.Once

func Get_Control_Lazy_lazyUnit() gopurs_runtime.Value {
	once_Control_Lazy_lazyUnit.Do(func() {
		cache_Control_Lazy_lazyUnit = gopurs_runtime.Value{Type: 9, IntVal: 1860244333, UnsafePtr: unsafe.Pointer((&Constructor_Control_Lazy_Lazy[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Get_Data_Unit_unit()
		})}))}
	})
	return cache_Control_Lazy_lazyUnit
}

var cache_Control_Lazy_lazyFn gopurs_runtime.Value
var once_Control_Lazy_lazyFn sync.Once

func Get_Control_Lazy_lazyFn() gopurs_runtime.Value {
	once_Control_Lazy_lazyFn.Do(func() {
		cache_Control_Lazy_lazyFn = gopurs_runtime.Value{Type: 9, IntVal: 1860244333, UnsafePtr: unsafe.Pointer((&Constructor_Control_Lazy_Lazy[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, x_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(f_0, Get_Data_Unit_unit(), x_1)
		})}))}
	})
	return cache_Control_Lazy_lazyFn
}

var cache_Control_Lazy_go__defer gopurs_runtime.Value
var once_Control_Lazy_go__defer sync.Once

func Get_Control_Lazy_go__defer() gopurs_runtime.Value {
	once_Control_Lazy_go__defer.Do(func() {
		cache_Control_Lazy_go__defer = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Lazy_go__defer(gopurs_runtime.CoerceToStruct[Constructor_Control_Lazy_Lazy[gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Control_Lazy_go__defer
}

var cache_Control_Lazy_fix gopurs_runtime.Value
var once_Control_Lazy_fix sync.Once

func Get_Control_Lazy_fix() gopurs_runtime.Value {
	once_Control_Lazy_fix.Do(func() {
		cache_Control_Lazy_fix = gopurs_runtime.Func2(func(dictLazy_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Lazy_fix(gopurs_runtime.CoerceToStruct[Constructor_Control_Lazy_Lazy[gopurs_runtime.Value]](dictLazy_0_box), f_1_box)
		})
	})
	return cache_Control_Lazy_fix
}

type Constructor_Control_Lazy_Lazy[T_l any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

//gopurs:init-v1:00670065006e006500720061007400650064002f0043006f006e00740072006f006c005f004c0061007a0079:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f0043006f006e00740072006f006c005f004c0061007a0079005f004c0061007a0079
func init() {
	gopurs_runtime.StructGetters[1860244333] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Control_Lazy_Lazy[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "defer":
			return gopurs_runtime.Box(c.V0)
		default:
			panic("Key not found in dictionary Constructor_Control_Lazy_Lazy: " + key)
		}
	}
}

func Call_Control_Lazy_Lazy_dollar_Dict(x_0_loop struct {
	go__defer gopurs_runtime.Value
	/*gopurs:record-v1:1:00640065006600650072*/
}) *Constructor_Control_Lazy_Lazy[gopurs_runtime.Value] {
	var x_0 struct {
		go__defer gopurs_runtime.Value
		/*gopurs:record-v1:1:00640065006600650072*/
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Control_Lazy_Lazy[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict1("defer", orig.go__defer)
	}())
}

func Call_Control_Lazy_Lazy_dollar_Dict__3034036302(x_0_loop struct {
	go__defer gopurs_runtime.Value
	/*gopurs:record-v1:1:00640065006600650072*/
}) *Constructor_Control_Lazy_Lazy[gopurs_runtime.Value] {
Lazy_dollar_Dict__3034036302:
	for {
		if false {
			continue Lazy_dollar_Dict__3034036302
		}
		var x_0 struct {
			go__defer gopurs_runtime.Value
			/*gopurs:record-v1:1:00640065006600650072*/
		} = x_0_loop
		_ = x_0
		return gopurs_runtime.CoerceToStruct[Constructor_Control_Lazy_Lazy[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			orig := x_0
			_ = orig
			return gopurs_runtime.RecordDict1("defer", orig.go__defer)
		}())
	}
}

func Call_Control_Lazy_go__defer(dict_0_loop *Constructor_Control_Lazy_Lazy[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Control_Lazy_Lazy[gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V0
}

func Call_Control_Lazy_fix(dictLazy_0_loop *Constructor_Control_Lazy_Lazy[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictLazy_0 *Constructor_Control_Lazy_Lazy[gopurs_runtime.Value] = dictLazy_0_loop
	_ = dictLazy_0
	var f_1 gopurs_runtime.Value = f_1_loop
	_ = f_1
	var go__go_2_0_0 gopurs_runtime.Value
	_ = go__go_2_0_0
	var go__go_2_0_0_cell *gopurs_runtime.Value
	_ = go__go_2_0_0_cell
	// FALLBACK TCO: isLoop=false len=1
	go__go_2_0_0 = gopurs_runtime.Apply(dictLazy_0.V0, gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(f_1, (*go__go_2_0_0_cell))
	}))
	go__go_2_0_0_cell = &go__go_2_0_0
	return go__go_2_0_0
}
