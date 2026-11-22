package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Monad_Trans_Class_MonadTrans_dollar_Dict gopurs_runtime.Value
var once_Control_Monad_Trans_Class_MonadTrans_dollar_Dict sync.Once

func Get_Control_Monad_Trans_Class_MonadTrans_dollar_Dict() gopurs_runtime.Value {
	once_Control_Monad_Trans_Class_MonadTrans_dollar_Dict.Do(func() {
		cache_Control_Monad_Trans_Class_MonadTrans_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2835982595, UnsafePtr: unsafe.Pointer(Call_Control_Monad_Trans_Class_MonadTrans_dollar_Dict(func() struct {
				lift gopurs_runtime.Value
				/*gopurs:record-v1:1:006c006900660074*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					lift gopurs_runtime.Value
					/*gopurs:record-v1:1:006c006900660074*/
				}{}
				clone.lift = gopurs_runtime.RecordGet(orig, "lift")
				return clone
			}()))}
		})
	})
	return cache_Control_Monad_Trans_Class_MonadTrans_dollar_Dict
}

var cache_Control_Monad_Trans_Class_lift gopurs_runtime.Value
var once_Control_Monad_Trans_Class_lift sync.Once

func Get_Control_Monad_Trans_Class_lift() gopurs_runtime.Value {
	once_Control_Monad_Trans_Class_lift.Do(func() {
		cache_Control_Monad_Trans_Class_lift = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Trans_Class_lift(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Trans_Class_MonadTrans[gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Control_Monad_Trans_Class_lift
}

type Constructor_Control_Monad_Trans_Class_MonadTrans[T_t any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

//gopurs:init-v1:00670065006e006500720061007400650064002f0043006f006e00740072006f006c005f004d006f006e00610064005f005400720061006e0073005f0043006c006100730073:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f0043006f006e00740072006f006c005f004d006f006e00610064005f005400720061006e0073005f0043006c006100730073005f004d006f006e00610064005400720061006e0073
func init() {
	gopurs_runtime.StructGetters[2835982595] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Control_Monad_Trans_Class_MonadTrans[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "lift":
			return gopurs_runtime.Box(c.V0)
		default:
			panic("Key not found in dictionary Constructor_Control_Monad_Trans_Class_MonadTrans: " + key)
		}
	}
}

func Call_Control_Monad_Trans_Class_MonadTrans_dollar_Dict(x_0_loop struct {
	lift gopurs_runtime.Value
	/*gopurs:record-v1:1:006c006900660074*/
}) *Constructor_Control_Monad_Trans_Class_MonadTrans[gopurs_runtime.Value] {
	var x_0 struct {
		lift gopurs_runtime.Value
		/*gopurs:record-v1:1:006c006900660074*/
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Trans_Class_MonadTrans[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict1("lift", orig.lift)
	}())
}

func Call_Control_Monad_Trans_Class_lift(dict_0_loop *Constructor_Control_Monad_Trans_Class_MonadTrans[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Control_Monad_Trans_Class_MonadTrans[gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V0
}
