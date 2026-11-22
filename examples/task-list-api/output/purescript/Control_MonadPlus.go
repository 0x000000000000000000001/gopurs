package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_MonadPlus_MonadPlus_dollar_Dict gopurs_runtime.Value
var once_Control_MonadPlus_MonadPlus_dollar_Dict sync.Once

func Get_Control_MonadPlus_MonadPlus_dollar_Dict() gopurs_runtime.Value {
	once_Control_MonadPlus_MonadPlus_dollar_Dict.Do(func() {
		cache_Control_MonadPlus_MonadPlus_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3236234573, UnsafePtr: unsafe.Pointer(Call_Control_MonadPlus_MonadPlus_dollar_Dict(func() struct {
				Alternative1 gopurs_runtime.Value
				Monad0       gopurs_runtime.Value
				/*gopurs:record-v1:2:0041006c007400650072006e006100740069007600650031,004d006f006e006100640030*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Alternative1 gopurs_runtime.Value
					Monad0       gopurs_runtime.Value
					/*gopurs:record-v1:2:0041006c007400650072006e006100740069007600650031,004d006f006e006100640030*/
				}{}
				clone.Alternative1 = gopurs_runtime.RecordGet(orig, "Alternative1")
				clone.Monad0 = gopurs_runtime.RecordGet(orig, "Monad0")
				return clone
			}()))}
		})
	})
	return cache_Control_MonadPlus_MonadPlus_dollar_Dict
}

var cache_Control_MonadPlus_MonadPlus_dollar_Dict__1423976851 gopurs_runtime.Value
var once_Control_MonadPlus_MonadPlus_dollar_Dict__1423976851 sync.Once

func Get_Control_MonadPlus_MonadPlus_dollar_Dict__1423976851() gopurs_runtime.Value {
	once_Control_MonadPlus_MonadPlus_dollar_Dict__1423976851.Do(func() {
		cache_Control_MonadPlus_MonadPlus_dollar_Dict__1423976851 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3236234573, UnsafePtr: unsafe.Pointer(Rebox_Control_MonadPlus_2656832021_2394491833(Call_Control_MonadPlus_MonadPlus_dollar_Dict__1423976851(func() struct {
				Alternative1 gopurs_runtime.Value
				Monad0       gopurs_runtime.Value
				/*gopurs:record-v1:2:0041006c007400650072006e006100740069007600650031,004d006f006e006100640030*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Alternative1 gopurs_runtime.Value
					Monad0       gopurs_runtime.Value
					/*gopurs:record-v1:2:0041006c007400650072006e006100740069007600650031,004d006f006e006100640030*/
				}{}
				clone.Alternative1 = gopurs_runtime.RecordGet(orig, "Alternative1")
				clone.Monad0 = gopurs_runtime.RecordGet(orig, "Monad0")
				return clone
			}())))}
		})
	})
	return cache_Control_MonadPlus_MonadPlus_dollar_Dict__1423976851
}

var cache_Control_MonadPlus_MonadPlus_dollar_Dict__2939863607 gopurs_runtime.Value
var once_Control_MonadPlus_MonadPlus_dollar_Dict__2939863607 sync.Once

func Get_Control_MonadPlus_MonadPlus_dollar_Dict__2939863607() gopurs_runtime.Value {
	once_Control_MonadPlus_MonadPlus_dollar_Dict__2939863607.Do(func() {
		cache_Control_MonadPlus_MonadPlus_dollar_Dict__2939863607 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3236234573, UnsafePtr: unsafe.Pointer(Call_Control_MonadPlus_MonadPlus_dollar_Dict__2939863607(func() struct {
				Alternative1 gopurs_runtime.Value
				Monad0       gopurs_runtime.Value
				/*gopurs:record-v1:2:0041006c007400650072006e006100740069007600650031,004d006f006e006100640030*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Alternative1 gopurs_runtime.Value
					Monad0       gopurs_runtime.Value
					/*gopurs:record-v1:2:0041006c007400650072006e006100740069007600650031,004d006f006e006100640030*/
				}{}
				clone.Alternative1 = gopurs_runtime.RecordGet(orig, "Alternative1")
				clone.Monad0 = gopurs_runtime.RecordGet(orig, "Monad0")
				return clone
			}()))}
		})
	})
	return cache_Control_MonadPlus_MonadPlus_dollar_Dict__2939863607
}

var cache_Control_MonadPlus_monadPlusArray gopurs_runtime.Value
var once_Control_MonadPlus_monadPlusArray sync.Once

func Get_Control_MonadPlus_monadPlusArray() gopurs_runtime.Value {
	once_Control_MonadPlus_monadPlusArray.Do(func() {
		cache_Control_MonadPlus_monadPlusArray = gopurs_runtime.Value{Type: 9, IntVal: 3236234573, UnsafePtr: unsafe.Pointer((&Constructor_Control_MonadPlus_MonadPlus[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 397869517, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Alternative_Alternative[gopurs_runtime.Value]](Get_Control_Alternative_alternativeArray()))}
		}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](Get_Control_Monad_monadArray()))}
		})}))}
	})
	return cache_Control_MonadPlus_monadPlusArray
}

type Constructor_Control_MonadPlus_MonadPlus[T_m any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}

//gopurs:init-v1:00670065006e006500720061007400650064002f0043006f006e00740072006f006c005f004d006f006e006100640050006c00750073:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f0043006f006e00740072006f006c005f004d006f006e006100640050006c00750073005f004d006f006e006100640050006c00750073
func init() {
	gopurs_runtime.StructGetters[3236234573] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Control_MonadPlus_MonadPlus[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "Alternative1":
			return gopurs_runtime.Box(c.V0)
		case "Monad0":
			return gopurs_runtime.Box(c.V1)
		default:
			panic("Key not found in dictionary Constructor_Control_MonadPlus_MonadPlus: " + key)
		}
	}
}

func Call_Control_MonadPlus_MonadPlus_dollar_Dict(x_0_loop struct {
	Alternative1 gopurs_runtime.Value
	Monad0       gopurs_runtime.Value
	/*gopurs:record-v1:2:0041006c007400650072006e006100740069007600650031,004d006f006e006100640030*/
}) *Constructor_Control_MonadPlus_MonadPlus[gopurs_runtime.Value] {
	var x_0 struct {
		Alternative1 gopurs_runtime.Value
		Monad0       gopurs_runtime.Value
		/*gopurs:record-v1:2:0041006c007400650072006e006100740069007600650031,004d006f006e006100640030*/
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Control_MonadPlus_MonadPlus[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict2("Alternative1", "Monad0", orig.Alternative1, orig.Monad0)
	}())
}

func Call_Control_MonadPlus_MonadPlus_dollar_Dict__1423976851(x_0_loop struct {
	Alternative1 gopurs_runtime.Value
	Monad0       gopurs_runtime.Value
	/*gopurs:record-v1:2:0041006c007400650072006e006100740069007600650031,004d006f006e006100640030*/
}) *Constructor_Control_MonadPlus_MonadPlus[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] {
MonadPlus_dollar_Dict__1423976851:
	for {
		if false {
			continue MonadPlus_dollar_Dict__1423976851
		}
		var x_0 struct {
			Alternative1 gopurs_runtime.Value
			Monad0       gopurs_runtime.Value
			/*gopurs:record-v1:2:0041006c007400650072006e006100740069007600650031,004d006f006e006100640030*/
		} = x_0_loop
		_ = x_0
		return gopurs_runtime.CoerceToStruct[Constructor_Control_MonadPlus_MonadPlus[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](func() gopurs_runtime.Value {
			orig := x_0
			_ = orig
			return gopurs_runtime.RecordDict2("Alternative1", "Monad0", orig.Alternative1, orig.Monad0)
		}())
	}
}

func Call_Control_MonadPlus_MonadPlus_dollar_Dict__2939863607(x_0_loop struct {
	Alternative1 gopurs_runtime.Value
	Monad0       gopurs_runtime.Value
	/*gopurs:record-v1:2:0041006c007400650072006e006100740069007600650031,004d006f006e006100640030*/
}) *Constructor_Control_MonadPlus_MonadPlus[gopurs_runtime.Value] {
MonadPlus_dollar_Dict__2939863607:
	for {
		if false {
			continue MonadPlus_dollar_Dict__2939863607
		}
		var x_0 struct {
			Alternative1 gopurs_runtime.Value
			Monad0       gopurs_runtime.Value
			/*gopurs:record-v1:2:0041006c007400650072006e006100740069007600650031,004d006f006e006100640030*/
		} = x_0_loop
		_ = x_0
		return gopurs_runtime.CoerceToStruct[Constructor_Control_MonadPlus_MonadPlus[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			orig := x_0
			_ = orig
			return gopurs_runtime.RecordDict2("Alternative1", "Monad0", orig.Alternative1, orig.Monad0)
		}())
	}
}

func Rebox_Control_MonadPlus_2656832021_2394491833(in *Constructor_Control_MonadPlus_MonadPlus[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]) *Constructor_Control_MonadPlus_MonadPlus[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Control_MonadPlus_MonadPlus[gopurs_runtime.Value])(unsafe.Pointer(in))
}
