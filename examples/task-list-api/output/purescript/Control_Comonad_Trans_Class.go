package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Comonad_Trans_Class_ComonadTrans_dollar_Dict gopurs_runtime.Value
var once_Control_Comonad_Trans_Class_ComonadTrans_dollar_Dict sync.Once

func Get_Control_Comonad_Trans_Class_ComonadTrans_dollar_Dict() gopurs_runtime.Value {
	once_Control_Comonad_Trans_Class_ComonadTrans_dollar_Dict.Do(func() {
		cache_Control_Comonad_Trans_Class_ComonadTrans_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3399197123, UnsafePtr: unsafe.Pointer(Call_Control_Comonad_Trans_Class_ComonadTrans_dollar_Dict(func() struct {
				lower gopurs_runtime.Value
				/*gopurs:record-v1:1:006c006f007700650072*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					lower gopurs_runtime.Value
					/*gopurs:record-v1:1:006c006f007700650072*/
				}{}
				clone.lower = gopurs_runtime.RecordGet(orig, "lower")
				return clone
			}()))}
		})
	})
	return cache_Control_Comonad_Trans_Class_ComonadTrans_dollar_Dict
}

var cache_Control_Comonad_Trans_Class_lower gopurs_runtime.Value
var once_Control_Comonad_Trans_Class_lower sync.Once

func Get_Control_Comonad_Trans_Class_lower() gopurs_runtime.Value {
	once_Control_Comonad_Trans_Class_lower.Do(func() {
		cache_Control_Comonad_Trans_Class_lower = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Comonad_Trans_Class_lower(gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Trans_Class_ComonadTrans[gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Control_Comonad_Trans_Class_lower
}

var cache_Control_Comonad_Trans_Class_comonadTransIdentityT gopurs_runtime.Value
var once_Control_Comonad_Trans_Class_comonadTransIdentityT sync.Once

func Get_Control_Comonad_Trans_Class_comonadTransIdentityT() gopurs_runtime.Value {
	once_Control_Comonad_Trans_Class_comonadTransIdentityT.Do(func() {
		cache_Control_Comonad_Trans_Class_comonadTransIdentityT = gopurs_runtime.Value{Type: 9, IntVal: 3399197123, UnsafePtr: unsafe.Pointer((&Constructor_Control_Comonad_Trans_Class_ComonadTrans[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(dictComonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Get_Control_Monad_Identity_Trans_runIdentityT()
		})}))}
	})
	return cache_Control_Comonad_Trans_Class_comonadTransIdentityT
}

type Constructor_Control_Comonad_Trans_Class_ComonadTrans[T_f any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

//gopurs:init-v1:00670065006e006500720061007400650064002f0043006f006e00740072006f006c005f0043006f006d006f006e00610064005f005400720061006e0073005f0043006c006100730073:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f0043006f006e00740072006f006c005f0043006f006d006f006e00610064005f005400720061006e0073005f0043006c006100730073005f0043006f006d006f006e00610064005400720061006e0073
func init() {
	gopurs_runtime.StructGetters[3399197123] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Control_Comonad_Trans_Class_ComonadTrans[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "lower":
			return gopurs_runtime.Box(c.V0)
		default:
			panic("Key not found in dictionary Constructor_Control_Comonad_Trans_Class_ComonadTrans: " + key)
		}
	}
}

func Call_Control_Comonad_Trans_Class_ComonadTrans_dollar_Dict(x_0_loop struct {
	lower gopurs_runtime.Value
	/*gopurs:record-v1:1:006c006f007700650072*/
}) *Constructor_Control_Comonad_Trans_Class_ComonadTrans[gopurs_runtime.Value] {
	var x_0 struct {
		lower gopurs_runtime.Value
		/*gopurs:record-v1:1:006c006f007700650072*/
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Trans_Class_ComonadTrans[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict1("lower", orig.lower)
	}())
}

func Call_Control_Comonad_Trans_Class_lower(dict_0_loop *Constructor_Control_Comonad_Trans_Class_ComonadTrans[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Control_Comonad_Trans_Class_ComonadTrans[gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V0
}
