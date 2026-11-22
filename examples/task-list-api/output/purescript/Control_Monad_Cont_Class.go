package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Monad_Cont_Class_MonadCont_dollar_Dict gopurs_runtime.Value
var once_Control_Monad_Cont_Class_MonadCont_dollar_Dict sync.Once

func Get_Control_Monad_Cont_Class_MonadCont_dollar_Dict() gopurs_runtime.Value {
	once_Control_Monad_Cont_Class_MonadCont_dollar_Dict.Do(func() {
		cache_Control_Monad_Cont_Class_MonadCont_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1800060259, UnsafePtr: unsafe.Pointer(Call_Control_Monad_Cont_Class_MonadCont_dollar_Dict(func() struct {
				Monad0 gopurs_runtime.Value
				callCC gopurs_runtime.Value
				/*gopurs:record-v1:2:004d006f006e006100640030,00630061006c006c00430043*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Monad0 gopurs_runtime.Value
					callCC gopurs_runtime.Value
					/*gopurs:record-v1:2:004d006f006e006100640030,00630061006c006c00430043*/
				}{}
				clone.Monad0 = gopurs_runtime.RecordGet(orig, "Monad0")
				clone.callCC = gopurs_runtime.RecordGet(orig, "callCC")
				return clone
			}()))}
		})
	})
	return cache_Control_Monad_Cont_Class_MonadCont_dollar_Dict
}

var cache_Control_Monad_Cont_Class_callCC gopurs_runtime.Value
var once_Control_Monad_Cont_Class_callCC sync.Once

func Get_Control_Monad_Cont_Class_callCC() gopurs_runtime.Value {
	once_Control_Monad_Cont_Class_callCC.Do(func() {
		cache_Control_Monad_Cont_Class_callCC = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Cont_Class_callCC(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Cont_Class_MonadCont[gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Control_Monad_Cont_Class_callCC
}

type Constructor_Control_Monad_Cont_Class_MonadCont[T_m any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}

//gopurs:init-v1:00670065006e006500720061007400650064002f0043006f006e00740072006f006c005f004d006f006e00610064005f0043006f006e0074005f0043006c006100730073:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f0043006f006e00740072006f006c005f004d006f006e00610064005f0043006f006e0074005f0043006c006100730073005f004d006f006e006100640043006f006e0074
func init() {
	gopurs_runtime.StructGetters[1800060259] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Control_Monad_Cont_Class_MonadCont[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "Monad0":
			return gopurs_runtime.Box(c.V0)
		case "callCC":
			return gopurs_runtime.Box(c.V1)
		default:
			panic("Key not found in dictionary Constructor_Control_Monad_Cont_Class_MonadCont: " + key)
		}
	}
}

func Call_Control_Monad_Cont_Class_MonadCont_dollar_Dict(x_0_loop struct {
	Monad0 gopurs_runtime.Value
	callCC gopurs_runtime.Value
	/*gopurs:record-v1:2:004d006f006e006100640030,00630061006c006c00430043*/
}) *Constructor_Control_Monad_Cont_Class_MonadCont[gopurs_runtime.Value] {
	var x_0 struct {
		Monad0 gopurs_runtime.Value
		callCC gopurs_runtime.Value
		/*gopurs:record-v1:2:004d006f006e006100640030,00630061006c006c00430043*/
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Cont_Class_MonadCont[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict2("Monad0", "callCC", orig.Monad0, orig.callCC)
	}())
}

func Call_Control_Monad_Cont_Class_callCC(dict_0_loop *Constructor_Control_Monad_Cont_Class_MonadCont[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Control_Monad_Cont_Class_MonadCont[gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V1
}
