package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Monad_Reader_Class_MonadAsk_dollar_Dict gopurs_runtime.Value
var once_Control_Monad_Reader_Class_MonadAsk_dollar_Dict sync.Once

func Get_Control_Monad_Reader_Class_MonadAsk_dollar_Dict() gopurs_runtime.Value {
	once_Control_Monad_Reader_Class_MonadAsk_dollar_Dict.Do(func() {
		cache_Control_Monad_Reader_Class_MonadAsk_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1229730751, UnsafePtr: unsafe.Pointer(Call_Control_Monad_Reader_Class_MonadAsk_dollar_Dict(func() struct {
				Monad0 gopurs_runtime.Value
				ask    gopurs_runtime.Value
				/*gopurs:record-v1:2:004d006f006e006100640030,00610073006b*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Monad0 gopurs_runtime.Value
					ask    gopurs_runtime.Value
					/*gopurs:record-v1:2:004d006f006e006100640030,00610073006b*/
				}{}
				clone.Monad0 = gopurs_runtime.RecordGet(orig, "Monad0")
				clone.ask = gopurs_runtime.RecordGet(orig, "ask")
				return clone
			}()))}
		})
	})
	return cache_Control_Monad_Reader_Class_MonadAsk_dollar_Dict
}

var cache_Control_Monad_Reader_Class_MonadReader_dollar_Dict gopurs_runtime.Value
var once_Control_Monad_Reader_Class_MonadReader_dollar_Dict sync.Once

func Get_Control_Monad_Reader_Class_MonadReader_dollar_Dict() gopurs_runtime.Value {
	once_Control_Monad_Reader_Class_MonadReader_dollar_Dict.Do(func() {
		cache_Control_Monad_Reader_Class_MonadReader_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2457234979, UnsafePtr: unsafe.Pointer(Call_Control_Monad_Reader_Class_MonadReader_dollar_Dict(func() struct {
				MonadAsk0 gopurs_runtime.Value
				local     gopurs_runtime.Value
				/*gopurs:record-v1:2:004d006f006e0061006400410073006b0030,006c006f00630061006c*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					MonadAsk0 gopurs_runtime.Value
					local     gopurs_runtime.Value
					/*gopurs:record-v1:2:004d006f006e0061006400410073006b0030,006c006f00630061006c*/
				}{}
				clone.MonadAsk0 = gopurs_runtime.RecordGet(orig, "MonadAsk0")
				clone.local = gopurs_runtime.RecordGet(orig, "local")
				return clone
			}()))}
		})
	})
	return cache_Control_Monad_Reader_Class_MonadReader_dollar_Dict
}

var cache_Control_Monad_Reader_Class_monadAskFun gopurs_runtime.Value
var once_Control_Monad_Reader_Class_monadAskFun sync.Once

func Get_Control_Monad_Reader_Class_monadAskFun() gopurs_runtime.Value {
	once_Control_Monad_Reader_Class_monadAskFun.Do(func() {
		cache_Control_Monad_Reader_Class_monadAskFun = gopurs_runtime.Value{Type: 9, IntVal: 1229730751, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Reader_Class_MonadAsk[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Get_Control_Monad_monadFn()
		}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return x_0
		})}))}
	})
	return cache_Control_Monad_Reader_Class_monadAskFun
}

var cache_Control_Monad_Reader_Class_monadReaderFun gopurs_runtime.Value
var once_Control_Monad_Reader_Class_monadReaderFun sync.Once

func Get_Control_Monad_Reader_Class_monadReaderFun() gopurs_runtime.Value {
	once_Control_Monad_Reader_Class_monadReaderFun.Do(func() {
		cache_Control_Monad_Reader_Class_monadReaderFun = gopurs_runtime.Value{Type: 9, IntVal: 2457234979, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Reader_Class_MonadReader[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Get_Control_Monad_Reader_Class_monadAskFun()
		}), Call_Control_Semigroupoid_composeFlipped(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))}))}
	})
	return cache_Control_Monad_Reader_Class_monadReaderFun
}

var cache_Control_Monad_Reader_Class_local gopurs_runtime.Value
var once_Control_Monad_Reader_Class_local sync.Once

func Get_Control_Monad_Reader_Class_local() gopurs_runtime.Value {
	once_Control_Monad_Reader_Class_local.Do(func() {
		cache_Control_Monad_Reader_Class_local = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Reader_Class_local(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Reader_Class_MonadReader[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Control_Monad_Reader_Class_local
}

var cache_Control_Monad_Reader_Class_ask gopurs_runtime.Value
var once_Control_Monad_Reader_Class_ask sync.Once

func Get_Control_Monad_Reader_Class_ask() gopurs_runtime.Value {
	once_Control_Monad_Reader_Class_ask.Do(func() {
		cache_Control_Monad_Reader_Class_ask = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Reader_Class_ask(dict_0_box)
		})
	})
	return cache_Control_Monad_Reader_Class_ask
}

var cache_Control_Monad_Reader_Class_asks gopurs_runtime.Value
var once_Control_Monad_Reader_Class_asks sync.Once

func Get_Control_Monad_Reader_Class_asks() gopurs_runtime.Value {
	once_Control_Monad_Reader_Class_asks.Do(func() {
		cache_Control_Monad_Reader_Class_asks = gopurs_runtime.Func(func(dictMonadAsk_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Reader_Class_asks(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Reader_Class_MonadAsk[gopurs_runtime.Value, gopurs_runtime.Value]](dictMonadAsk_0_box))
		})
	})
	return cache_Control_Monad_Reader_Class_asks
}

type Constructor_Control_Monad_Reader_Class_MonadAsk[T_r any, T_m any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}

//gopurs:init-v1:00670065006e006500720061007400650064002f0043006f006e00740072006f006c005f004d006f006e00610064005f005200650061006400650072005f0043006c006100730073:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f0043006f006e00740072006f006c005f004d006f006e00610064005f005200650061006400650072005f0043006c006100730073005f004d006f006e0061006400410073006b
func init() {
	gopurs_runtime.StructGetters[1229730751] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Control_Monad_Reader_Class_MonadAsk[gopurs_runtime.Value, gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "Monad0":
			return gopurs_runtime.Box(c.V0)
		case "ask":
			return gopurs_runtime.Box(c.V1)
		default:
			panic("Key not found in dictionary Constructor_Control_Monad_Reader_Class_MonadAsk: " + key)
		}
	}
}

type Constructor_Control_Monad_Reader_Class_MonadReader[T_r any, T_m any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}

//gopurs:init-v1:00670065006e006500720061007400650064002f0043006f006e00740072006f006c005f004d006f006e00610064005f005200650061006400650072005f0043006c006100730073:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f0043006f006e00740072006f006c005f004d006f006e00610064005f005200650061006400650072005f0043006c006100730073005f004d006f006e00610064005200650061006400650072
func init() {
	gopurs_runtime.StructGetters[2457234979] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Control_Monad_Reader_Class_MonadReader[gopurs_runtime.Value, gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "MonadAsk0":
			return gopurs_runtime.Box(c.V0)
		case "local":
			return gopurs_runtime.Box(c.V1)
		default:
			panic("Key not found in dictionary Constructor_Control_Monad_Reader_Class_MonadReader: " + key)
		}
	}
}

func Call_Control_Monad_Reader_Class_MonadAsk_dollar_Dict(x_0_loop struct {
	Monad0 gopurs_runtime.Value
	ask    gopurs_runtime.Value
	/*gopurs:record-v1:2:004d006f006e006100640030,00610073006b*/
}) *Constructor_Control_Monad_Reader_Class_MonadAsk[gopurs_runtime.Value, gopurs_runtime.Value] {
	var x_0 struct {
		Monad0 gopurs_runtime.Value
		ask    gopurs_runtime.Value
		/*gopurs:record-v1:2:004d006f006e006100640030,00610073006b*/
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Reader_Class_MonadAsk[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict2("Monad0", "ask", orig.Monad0, orig.ask)
	}())
}

func Call_Control_Monad_Reader_Class_MonadReader_dollar_Dict(x_0_loop struct {
	MonadAsk0 gopurs_runtime.Value
	local     gopurs_runtime.Value
	/*gopurs:record-v1:2:004d006f006e0061006400410073006b0030,006c006f00630061006c*/
}) *Constructor_Control_Monad_Reader_Class_MonadReader[gopurs_runtime.Value, gopurs_runtime.Value] {
	var x_0 struct {
		MonadAsk0 gopurs_runtime.Value
		local     gopurs_runtime.Value
		/*gopurs:record-v1:2:004d006f006e0061006400410073006b0030,006c006f00630061006c*/
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Reader_Class_MonadReader[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict2("MonadAsk0", "local", orig.MonadAsk0, orig.local)
	}())
}

func Call_Control_Monad_Reader_Class_local(dict_0_loop *Constructor_Control_Monad_Reader_Class_MonadReader[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Control_Monad_Reader_Class_MonadReader[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V1
}

func Call_Control_Monad_Reader_Class_ask(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dict_0 gopurs_runtime.Value = dict_0_loop
	_ = dict_0
	return gopurs_runtime.RecordGet(dict_0, "ask")
}

func Call_Control_Monad_Reader_Class_asks(dictMonadAsk_0_loop *Constructor_Control_Monad_Reader_Class_MonadAsk[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictMonadAsk_0 *Constructor_Control_Monad_Reader_Class_MonadAsk[gopurs_runtime.Value, gopurs_runtime.Value] = dictMonadAsk_0_loop
	_ = dictMonadAsk_0
	// TAST (Let): ask1_1_0 shape=App(Var) bindingType=(TypeApp (TypeVar m$scope21) [(TypeVar r$scope20)])
	ask1_1_0 := Call_Control_Monad_Reader_Class_ask(gopurs_runtime.Value{Type: 9, IntVal: 1229730751, UnsafePtr: unsafe.Pointer(dictMonadAsk_0)})
	_ = ask1_1_0
	return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictMonadAsk_0.V0, gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), f_2, ask1_1_0)
	})
}
