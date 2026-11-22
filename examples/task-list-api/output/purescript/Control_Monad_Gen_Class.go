package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Monad_Gen_Class_MonadGen_dollar_Dict gopurs_runtime.Value
var once_Control_Monad_Gen_Class_MonadGen_dollar_Dict sync.Once

func Get_Control_Monad_Gen_Class_MonadGen_dollar_Dict() gopurs_runtime.Value {
	once_Control_Monad_Gen_Class_MonadGen_dollar_Dict.Do(func() {
		cache_Control_Monad_Gen_Class_MonadGen_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2254593219, UnsafePtr: unsafe.Pointer(Call_Control_Monad_Gen_Class_MonadGen_dollar_Dict(func() struct {
				Monad0      gopurs_runtime.Value
				chooseBool  gopurs_runtime.Value
				chooseFloat gopurs_runtime.Value
				chooseInt   gopurs_runtime.Value
				resize      gopurs_runtime.Value
				sized       gopurs_runtime.Value
				/*gopurs:record-v1:6:004d006f006e006100640030,00630068006f006f007300650042006f006f006c,00630068006f006f007300650046006c006f00610074,00630068006f006f007300650049006e0074,0072006500730069007a0065,00730069007a00650064*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Monad0      gopurs_runtime.Value
					chooseBool  gopurs_runtime.Value
					chooseFloat gopurs_runtime.Value
					chooseInt   gopurs_runtime.Value
					resize      gopurs_runtime.Value
					sized       gopurs_runtime.Value
					/*gopurs:record-v1:6:004d006f006e006100640030,00630068006f006f007300650042006f006f006c,00630068006f006f007300650046006c006f00610074,00630068006f006f007300650049006e0074,0072006500730069007a0065,00730069007a00650064*/
				}{}
				clone.Monad0 = gopurs_runtime.RecordGet(orig, "Monad0")
				clone.chooseBool = gopurs_runtime.RecordGet(orig, "chooseBool")
				clone.chooseFloat = gopurs_runtime.RecordGet(orig, "chooseFloat")
				clone.chooseInt = gopurs_runtime.RecordGet(orig, "chooseInt")
				clone.resize = gopurs_runtime.RecordGet(orig, "resize")
				clone.sized = gopurs_runtime.RecordGet(orig, "sized")
				return clone
			}()))}
		})
	})
	return cache_Control_Monad_Gen_Class_MonadGen_dollar_Dict
}

var cache_Control_Monad_Gen_Class_sized gopurs_runtime.Value
var once_Control_Monad_Gen_Class_sized sync.Once

func Get_Control_Monad_Gen_Class_sized() gopurs_runtime.Value {
	once_Control_Monad_Gen_Class_sized.Do(func() {
		cache_Control_Monad_Gen_Class_sized = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Gen_Class_sized(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Control_Monad_Gen_Class_sized
}

var cache_Control_Monad_Gen_Class_resize gopurs_runtime.Value
var once_Control_Monad_Gen_Class_resize sync.Once

func Get_Control_Monad_Gen_Class_resize() gopurs_runtime.Value {
	once_Control_Monad_Gen_Class_resize.Do(func() {
		cache_Control_Monad_Gen_Class_resize = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Gen_Class_resize(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Control_Monad_Gen_Class_resize
}

var cache_Control_Monad_Gen_Class_chooseInt gopurs_runtime.Value
var once_Control_Monad_Gen_Class_chooseInt sync.Once

func Get_Control_Monad_Gen_Class_chooseInt() gopurs_runtime.Value {
	once_Control_Monad_Gen_Class_chooseInt.Do(func() {
		cache_Control_Monad_Gen_Class_chooseInt = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Gen_Class_chooseInt(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Control_Monad_Gen_Class_chooseInt
}

var cache_Control_Monad_Gen_Class_chooseInt__3759223575 gopurs_runtime.Value
var once_Control_Monad_Gen_Class_chooseInt__3759223575 sync.Once

func Get_Control_Monad_Gen_Class_chooseInt__3759223575() gopurs_runtime.Value {
	once_Control_Monad_Gen_Class_chooseInt__3759223575.Do(func() {
		cache_Control_Monad_Gen_Class_chooseInt__3759223575 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Gen_Class_chooseInt__3759223575(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Control_Monad_Gen_Class_chooseInt__3759223575
}

var cache_Control_Monad_Gen_Class_chooseFloat gopurs_runtime.Value
var once_Control_Monad_Gen_Class_chooseFloat sync.Once

func Get_Control_Monad_Gen_Class_chooseFloat() gopurs_runtime.Value {
	once_Control_Monad_Gen_Class_chooseFloat.Do(func() {
		cache_Control_Monad_Gen_Class_chooseFloat = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Gen_Class_chooseFloat(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Control_Monad_Gen_Class_chooseFloat
}

var cache_Control_Monad_Gen_Class_chooseBool gopurs_runtime.Value
var once_Control_Monad_Gen_Class_chooseBool sync.Once

func Get_Control_Monad_Gen_Class_chooseBool() gopurs_runtime.Value {
	once_Control_Monad_Gen_Class_chooseBool.Do(func() {
		cache_Control_Monad_Gen_Class_chooseBool = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Gen_Class_chooseBool(dict_0_box)
		})
	})
	return cache_Control_Monad_Gen_Class_chooseBool
}

type Constructor_Control_Monad_Gen_Class_MonadGen[T_m any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
	V3 gopurs_runtime.Value
	V4 gopurs_runtime.Value
	V5 gopurs_runtime.Value
}

//gopurs:init-v1:00670065006e006500720061007400650064002f0043006f006e00740072006f006c005f004d006f006e00610064005f00470065006e005f0043006c006100730073:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f0043006f006e00740072006f006c005f004d006f006e00610064005f00470065006e005f0043006c006100730073005f004d006f006e0061006400470065006e
func init() {
	gopurs_runtime.StructGetters[2254593219] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "Monad0":
			return gopurs_runtime.Box(c.V0)
		case "chooseBool":
			return gopurs_runtime.Box(c.V1)
		case "chooseFloat":
			return gopurs_runtime.Box(c.V2)
		case "chooseInt":
			return gopurs_runtime.Box(c.V3)
		case "resize":
			return gopurs_runtime.Box(c.V4)
		case "sized":
			return gopurs_runtime.Box(c.V5)
		default:
			panic("Key not found in dictionary Constructor_Control_Monad_Gen_Class_MonadGen: " + key)
		}
	}
}

func Call_Control_Monad_Gen_Class_MonadGen_dollar_Dict(x_0_loop struct {
	Monad0      gopurs_runtime.Value
	chooseBool  gopurs_runtime.Value
	chooseFloat gopurs_runtime.Value
	chooseInt   gopurs_runtime.Value
	resize      gopurs_runtime.Value
	sized       gopurs_runtime.Value
	/*gopurs:record-v1:6:004d006f006e006100640030,00630068006f006f007300650042006f006f006c,00630068006f006f007300650046006c006f00610074,00630068006f006f007300650049006e0074,0072006500730069007a0065,00730069007a00650064*/
}) *Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value] {
	var x_0 struct {
		Monad0      gopurs_runtime.Value
		chooseBool  gopurs_runtime.Value
		chooseFloat gopurs_runtime.Value
		chooseInt   gopurs_runtime.Value
		resize      gopurs_runtime.Value
		sized       gopurs_runtime.Value
		/*gopurs:record-v1:6:004d006f006e006100640030,00630068006f006f007300650042006f006f006c,00630068006f006f007300650046006c006f00610074,00630068006f006f007300650049006e0074,0072006500730069007a0065,00730069007a00650064*/
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict([]string{"Monad0", "chooseBool", "chooseFloat", "chooseInt", "resize", "sized"}, []gopurs_runtime.Value{orig.Monad0, orig.chooseBool, orig.chooseFloat, orig.chooseInt, orig.resize, orig.sized})
	}())
}

func Call_Control_Monad_Gen_Class_sized(dict_0_loop *Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V5
}

func Call_Control_Monad_Gen_Class_resize(dict_0_loop *Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V4
}

func Call_Control_Monad_Gen_Class_chooseInt(dict_0_loop *Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V3
}

func Call_Control_Monad_Gen_Class_chooseInt__3759223575(dict_0_loop *Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]) gopurs_runtime.Value {
chooseInt__3759223575:
	for {
		if false {
			continue chooseInt__3759223575
		}
		var dict_0 *Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value] = dict_0_loop
		_ = dict_0
		return dict_0.V3
	}
}

func Call_Control_Monad_Gen_Class_chooseFloat(dict_0_loop *Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V2
}

func Call_Control_Monad_Gen_Class_chooseBool(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dict_0 gopurs_runtime.Value = dict_0_loop
	_ = dict_0
	return gopurs_runtime.RecordGet(dict_0, "chooseBool")
}
