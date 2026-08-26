package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_show gopurs_runtime.Value
var once_Main_show sync.Once

func Get_Main_show() gopurs_runtime.Value {
	once_Main_show.Do(func() {
		cache_Main_show = Get_Data_Show_showIntImpl()
	})
	return cache_Main_show
}

var cache_Main_Y gopurs_runtime.Value
var once_Main_Y sync.Once

func Get_Main_Y() gopurs_runtime.Value {
	once_Main_Y.Do(func() {
		cache_Main_Y = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_Y(x_0_box)
		})
	})
	return cache_Main_Y
}

var cache_Main_Singleton_dollar_Dict gopurs_runtime.Value
var once_Main_Singleton_dollar_Dict sync.Once

func Get_Main_Singleton_dollar_Dict() gopurs_runtime.Value {
	once_Main_Singleton_dollar_Dict.Do(func() {
		cache_Main_Singleton_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_Singleton_dollar_Dict(x_0_box)
		})
	})
	return cache_Main_Singleton_dollar_Dict
}

var cache_Main_ProxyArray gopurs_runtime.Value
var once_Main_ProxyArray sync.Once

func Get_Main_ProxyArray() gopurs_runtime.Value {
	once_Main_ProxyArray.Do(func() {
		cache_Main_ProxyArray = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_ProxyArray(x_0_box)
		})
	})
	return cache_Main_ProxyArray
}

var cache_Main_Proxy2 gopurs_runtime.Value
var once_Main_Proxy2 sync.Once

func Get_Main_Proxy2() gopurs_runtime.Value {
	once_Main_Proxy2.Do(func() {
		cache_Main_Proxy2 = gopurs_runtime.Value{Type: 9, IntVal: int64(4225449536), UnsafePtr: nil}
	})
	return cache_Main_Proxy2
}

var cache_Main_MyWriter gopurs_runtime.Value
var once_Main_MyWriter sync.Once

func Get_Main_MyWriter() gopurs_runtime.Value {
	once_Main_MyWriter.Do(func() {
		cache_Main_MyWriter = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_MyWriter(x_0_box)
		})
	})
	return cache_Main_MyWriter
}

var cache_Main_X gopurs_runtime.Value
var once_Main_X sync.Once

func Get_Main_X() gopurs_runtime.Value {
	once_Main_X.Do(func() {
		cache_Main_X = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_X(x_0_box)
		})
	})
	return cache_Main_X
}

var cache_Main_MyArray gopurs_runtime.Value
var once_Main_MyArray sync.Once

func Get_Main_MyArray() gopurs_runtime.Value {
	once_Main_MyArray.Do(func() {
		cache_Main_MyArray = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_MyArray(x_0_box)
		})
	})
	return cache_Main_MyArray
}

var cache_Main_Syn gopurs_runtime.Value
var once_Main_Syn sync.Once

func Get_Main_Syn() gopurs_runtime.Value {
	once_Main_Syn.Do(func() {
		cache_Main_Syn = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_Syn(x_0_box)
		})
	})
	return cache_Main_Syn
}

var cache_Main_MonadWriter_dollar_Dict gopurs_runtime.Value
var once_Main_MonadWriter_dollar_Dict sync.Once

func Get_Main_MonadWriter_dollar_Dict() gopurs_runtime.Value {
	once_Main_MonadWriter_dollar_Dict.Do(func() {
		cache_Main_MonadWriter_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_MonadWriter_dollar_Dict(x_0_box)
		})
	})
	return cache_Main_MonadWriter_dollar_Dict
}

var cache_Main_Foo gopurs_runtime.Value
var once_Main_Foo sync.Once

func Get_Main_Foo() gopurs_runtime.Value {
	once_Main_Foo.Do(func() {
		cache_Main_Foo = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_Foo(x_0_box)
		})
	})
	return cache_Main_Foo
}

var cache_Main_functorProxy2 gopurs_runtime.Value
var once_Main_functorProxy2 sync.Once

func Get_Main_functorProxy2() gopurs_runtime.Value {
	once_Main_functorProxy2.Do(func() {
		cache_Main_functorProxy2 = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor[uint32]{1, gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: int64(4225449536), UnsafePtr: nil}
			})
		})}))}
	})
	return cache_Main_functorProxy2
}

var cache_Main_functorProxy2__2778612294 gopurs_runtime.Value
var once_Main_functorProxy2__2778612294 sync.Once

func Get_Main_functorProxy2__2778612294() gopurs_runtime.Value {
	once_Main_functorProxy2__2778612294.Do(func() {
		cache_Main_functorProxy2__2778612294 = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor[uint32]{1, gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: int64(4225449536), UnsafePtr: nil}
			})
		})}))}
	})
	return cache_Main_functorProxy2__2778612294
}

var cache_Main_functorFoo gopurs_runtime.Value
var once_Main_functorFoo sync.Once

func Get_Main_functorFoo() gopurs_runtime.Value {
	once_Main_functorFoo.Do(func() {
		cache_Main_functorFoo = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[uint32]](Get_Main_functorProxy2()))}
	})
	return cache_Main_functorFoo
}

var cache_Main_tell gopurs_runtime.Value
var once_Main_tell sync.Once

func Get_Main_tell() gopurs_runtime.Value {
	once_Main_tell.Do(func() {
		cache_Main_tell = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_tell(gopurs_runtime.CoerceToStruct[Constructor_Main_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Main_tell
}

var cache_Main_singletonArray gopurs_runtime.Value
var once_Main_singletonArray sync.Once

func Get_Main_singletonArray() gopurs_runtime.Value {
	once_Main_singletonArray.Do(func() {
		cache_Main_singletonArray = gopurs_runtime.Value{Type: 9, IntVal: 3773262993, UnsafePtr: unsafe.Pointer((&Constructor_Main_Singleton[gopurs_runtime.Value, []gopurs_runtime.Value]{1, gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Array(func() []gopurs_runtime.Value {
				arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{x_0}).UnsafePtr)
				unboxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					unboxed[i] = v
				}
				return unboxed
			}())
		})}))}
	})
	return cache_Main_singletonArray
}

var cache_Main_singletonArray__1185293142 gopurs_runtime.Value
var once_Main_singletonArray__1185293142 sync.Once

func Get_Main_singletonArray__1185293142() gopurs_runtime.Value {
	once_Main_singletonArray__1185293142.Do(func() {
		cache_Main_singletonArray__1185293142 = gopurs_runtime.Value{Type: 9, IntVal: 3773262993, UnsafePtr: unsafe.Pointer((&Constructor_Main_Singleton[gopurs_runtime.Value, []gopurs_runtime.Value]{1, gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Array(func() []gopurs_runtime.Value {
				arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{x_0}).UnsafePtr)
				unboxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					unboxed[i] = v
				}
				return unboxed
			}())
		})}))}
	})
	return cache_Main_singletonArray__1185293142
}

var cache_Main_singletonY gopurs_runtime.Value
var once_Main_singletonY sync.Once

func Get_Main_singletonY() gopurs_runtime.Value {
	once_Main_singletonY.Do(func() {
		cache_Main_singletonY = gopurs_runtime.Value{Type: 9, IntVal: 3773262993, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Main_Singleton[gopurs_runtime.Value, []gopurs_runtime.Value]](Get_Main_singletonArray()))}
	})
	return cache_Main_singletonY
}

var cache_Main_singletonY__3643848278 gopurs_runtime.Value
var once_Main_singletonY__3643848278 sync.Once

func Get_Main_singletonY__3643848278() gopurs_runtime.Value {
	once_Main_singletonY__3643848278.Do(func() {
		cache_Main_singletonY__3643848278 = gopurs_runtime.Value{Type: 9, IntVal: 3773262993, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Main_Singleton[gopurs_runtime.Value, []gopurs_runtime.Value]](Get_Main_singletonArray()))}
	})
	return cache_Main_singletonY__3643848278
}

var cache_Main_singleton gopurs_runtime.Value
var once_Main_singleton sync.Once

func Get_Main_singleton() gopurs_runtime.Value {
	once_Main_singleton.Do(func() {
		cache_Main_singleton = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_singleton(gopurs_runtime.CoerceToStruct[Constructor_Main_Singleton[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Main_singleton
}

var cache_Main_singleton__1633385358 gopurs_runtime.Value
var once_Main_singleton__1633385358 sync.Once

func Get_Main_singleton__1633385358() gopurs_runtime.Value {
	once_Main_singleton__1633385358.Do(func() {
		cache_Main_singleton__1633385358 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_singleton__1633385358(gopurs_runtime.CoerceToStruct[Constructor_Main_Singleton[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Main_singleton__1633385358
}

var cache_Main_showY gopurs_runtime.Value
var once_Main_showY sync.Once

func Get_Main_showY() gopurs_runtime.Value {
	once_Main_showY.Do(func() {
		cache_Main_showY = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer((&Constructor_Data_Show_Show[[]string]{1, gopurs_runtime.Apply(Get_Data_Show_showArrayImpl(), Get_Data_Show_showStringImpl())}))}
	})
	return cache_Main_showY
}

var cache_Main_showX gopurs_runtime.Value
var once_Main_showX sync.Once

func Get_Main_showX() gopurs_runtime.Value {
	once_Main_showX.Do(func() {
		cache_Main_showX = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[string]](Get_Data_Show_showString()))}
	})
	return cache_Main_showX
}

var cache_Main_showMyArray gopurs_runtime.Value
var once_Main_showMyArray sync.Once

func Get_Main_showMyArray() gopurs_runtime.Value {
	once_Main_showMyArray.Do(func() {
		cache_Main_showMyArray = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_showMyArray(dictShow_0_box)
		})
	})
	return cache_Main_showMyArray
}

var cache_Main_showMyArray1 gopurs_runtime.Value
var once_Main_showMyArray1 sync.Once

func Get_Main_showMyArray1() gopurs_runtime.Value {
	once_Main_showMyArray1.Do(func() {
		cache_Main_showMyArray1 = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer((&Constructor_Data_Show_Show[[]string]{1, gopurs_runtime.Apply(Get_Data_Show_showArrayImpl(), Get_Data_Show_showStringImpl())}))}
	})
	return cache_Main_showMyArray1
}

var cache_Main_ordX gopurs_runtime.Value
var once_Main_ordX sync.Once

func Get_Main_ordX() gopurs_runtime.Value {
	once_Main_ordX.Do(func() {
		cache_Main_ordX = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[string]](Get_Data_Ord_ordString()))}
	})
	return cache_Main_ordX
}

var cache_Main_monadWriterTuple gopurs_runtime.Value
var once_Main_monadWriterTuple sync.Once

func Get_Main_monadWriterTuple() gopurs_runtime.Value {
	once_Main_monadWriterTuple.Do(func() {
		cache_Main_monadWriterTuple = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_monadWriterTuple(dictMonoid_0_box)
		})
	})
	return cache_Main_monadWriterTuple
}

var cache_Main_monadWriterMyWriter gopurs_runtime.Value
var once_Main_monadWriterMyWriter sync.Once

func Get_Main_monadWriterMyWriter() gopurs_runtime.Value {
	once_Main_monadWriterMyWriter.Do(func() {
		cache_Main_monadWriterMyWriter = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_monadWriterMyWriter(dictMonoid_0_box)
		})
	})
	return cache_Main_monadWriterMyWriter
}

var cache_Main_monadMyWriter gopurs_runtime.Value
var once_Main_monadMyWriter sync.Once

func Get_Main_monadMyWriter() gopurs_runtime.Value {
	once_Main_monadMyWriter.Do(func() {
		cache_Main_monadMyWriter = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_monadMyWriter(dictMonoid_0_box)
		})
	})
	return cache_Main_monadMyWriter
}

var cache_Main_functorProxyArray gopurs_runtime.Value
var once_Main_functorProxyArray sync.Once

func Get_Main_functorProxyArray() gopurs_runtime.Value {
	once_Main_functorProxyArray.Do(func() {
		cache_Main_functorProxyArray = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Functor_functorArray()))}
	})
	return cache_Main_functorProxyArray
}

var cache_Main_functorMyWriter gopurs_runtime.Value
var once_Main_functorMyWriter sync.Once

func Get_Main_functorMyWriter() gopurs_runtime.Value {
	once_Main_functorMyWriter.Do(func() {
		cache_Main_functorMyWriter = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_Data_Tuple_functorTuple()))}
	})
	return cache_Main_functorMyWriter
}

var cache_Main_functorMyWriter__2345709879 gopurs_runtime.Value
var once_Main_functorMyWriter__2345709879 sync.Once

func Get_Main_functorMyWriter__2345709879() gopurs_runtime.Value {
	once_Main_functorMyWriter__2345709879.Do(func() {
		cache_Main_functorMyWriter__2345709879 = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_Data_Tuple_functorTuple()))}
	})
	return cache_Main_functorMyWriter__2345709879
}

var cache_Main_functorSyn gopurs_runtime.Value
var once_Main_functorSyn sync.Once

func Get_Main_functorSyn() gopurs_runtime.Value {
	once_Main_functorSyn.Do(func() {
		cache_Main_functorSyn = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_Data_Tuple_functorTuple()))}
	})
	return cache_Main_functorSyn
}

var cache_Main_functorMyArray gopurs_runtime.Value
var once_Main_functorMyArray sync.Once

func Get_Main_functorMyArray() gopurs_runtime.Value {
	once_Main_functorMyArray.Do(func() {
		cache_Main_functorMyArray = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Functor_functorArray()))}
	})
	return cache_Main_functorMyArray
}

var cache_Main_functorMyArray__2527715796 gopurs_runtime.Value
var once_Main_functorMyArray__2527715796 sync.Once

func Get_Main_functorMyArray__2527715796() gopurs_runtime.Value {
	once_Main_functorMyArray__2527715796.Do(func() {
		cache_Main_functorMyArray__2527715796 = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Functor_functorArray()))}
	})
	return cache_Main_functorMyArray__2527715796
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=(ADT ["Effect","Effect"] [Unit])
			__local_var_0_0 := gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showStringImpl(), gopurs_runtime.Str("test")).StrVal()))
			_ = __local_var_0_0
			_dollar___unused_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = _dollar___unused_1_1
			_dollar___unused_2_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[[]string]](Get_Main_showY()).V0), func() gopurs_runtime.Value {
				arr := func() []string {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{gopurs_runtime.Str("test")}).UnsafePtr)
						unboxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							unboxed[i] = v
						}
						return unboxed
					}()).UnsafePtr)
					unboxed := make([]string, len(arr))
					for i, v := range arr {
						unboxed[i] = v.StrVal()
					}
					return unboxed
				}()
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Str(v)
				}
				return gopurs_runtime.Array(boxed)
			}()).StrVal())), gopurs_runtime.Value{})
			_ = _dollar___unused_2_2
			_dollar___unused_3_3 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[[]string]](Get_Main_showMyArray1()).V0), func() gopurs_runtime.Value {
				arr := func() []string {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
						arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
							arr_val_arrayMap6 := func() gopurs_runtime.Value {
								arr := []int64{1, 2, 3}
								boxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									boxed[i] = gopurs_runtime.Int(v)
								}
								return gopurs_runtime.Array(boxed)
							}()
							_ = arr_val_arrayMap6
							arr_go_arrayMap6 := (*[]gopurs_runtime.Value)(arr_val_arrayMap6.UnsafePtr)
							_ = arr_go_arrayMap6
							res_go_arrayMap6 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap6))
							_ = res_go_arrayMap6
							for i_arrayMap6, v_arrayMap6 := range *arr_go_arrayMap6 {
								res_go_arrayMap6[i_arrayMap6] = gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), v_arrayMap6)
							}
							return gopurs_runtime.Array(res_go_arrayMap6)
						}().UnsafePtr)
						unboxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							unboxed[i] = v
						}
						return unboxed
					}()).UnsafePtr)
					unboxed := make([]string, len(arr))
					for i, v := range arr {
						unboxed[i] = v.StrVal()
					}
					return unboxed
				}()
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Str(v)
				}
				return gopurs_runtime.Array(boxed)
			}()).StrVal())), gopurs_runtime.Value{})
			_ = _dollar___unused_3_3
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

var cache_Main_eqX gopurs_runtime.Value
var once_Main_eqX sync.Once

func Get_Main_eqX() gopurs_runtime.Value {
	once_Main_eqX.Do(func() {
		cache_Main_eqX = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[string]](Get_Data_Eq_eqString()))}
	})
	return cache_Main_eqX
}

var cache_Main_bindMyWriter gopurs_runtime.Value
var once_Main_bindMyWriter sync.Once

func Get_Main_bindMyWriter() gopurs_runtime.Value {
	once_Main_bindMyWriter.Do(func() {
		cache_Main_bindMyWriter = gopurs_runtime.Func(func(dictSemigroup_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_bindMyWriter(dictSemigroup_0_box)
		})
	})
	return cache_Main_bindMyWriter
}

var cache_Main_applyMyWriter gopurs_runtime.Value
var once_Main_applyMyWriter sync.Once

func Get_Main_applyMyWriter() gopurs_runtime.Value {
	once_Main_applyMyWriter.Do(func() {
		cache_Main_applyMyWriter = gopurs_runtime.Func(func(dictSemigroup_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_applyMyWriter(dictSemigroup_0_box)
		})
	})
	return cache_Main_applyMyWriter
}

var cache_Main_applicativeMyWriter gopurs_runtime.Value
var once_Main_applicativeMyWriter sync.Once

func Get_Main_applicativeMyWriter() gopurs_runtime.Value {
	once_Main_applicativeMyWriter.Do(func() {
		cache_Main_applicativeMyWriter = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_applicativeMyWriter(dictMonoid_0_box)
		})
	})
	return cache_Main_applicativeMyWriter
}

type Constructor_Main_Proxy2[T_a any, T_b any] struct {
	Rc uint32
}

type Constructor_Main_Singleton[T_a any, T_b any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

func init() {
	gopurs_runtime.StructGetters[3773262993] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_Singleton[any, any])(ptr)
		_ = c
		switch key {
		case "singleton":
			return gopurs_runtime.Box(c.V0)
		default:
			panic("Key not found in dictionary Constructor_Main_Singleton: " + key)
		}
	}
}

type Constructor_Main_MonadWriter[T_w any, T_m any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
}

func init() {
	gopurs_runtime.StructGetters[2544837208] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_MonadWriter[any, any])(ptr)
		_ = c
		switch key {
		case "Monad0":
			return gopurs_runtime.Box(c.V0)
		case "Monoid1":
			return gopurs_runtime.Box(c.V1)
		case "tell":
			return gopurs_runtime.Box(c.V2)
		default:
			panic("Key not found in dictionary Constructor_Main_MonadWriter: " + key)
		}
	}
}

func Call_Main_Y(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_Singleton_dollar_Dict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_ProxyArray(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_MyWriter(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_X(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_MyArray(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_Syn(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_MonadWriter_dollar_Dict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_Foo(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_tell(dict_0_loop *Constructor_Main_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return gopurs_runtime.Box(dict_0.V2)
}

func Call_Main_singleton(dict_0_loop *Constructor_Main_Singleton[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_Singleton[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return gopurs_runtime.Box(dict_0.V0)
}

func Call_Main_singleton__1633385358(dict_0_loop *Constructor_Main_Singleton[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_Singleton[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return gopurs_runtime.Box(dict_0.V0)
}

func Call_Main_showMyArray(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
	_ = dictShow_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer((&Constructor_Data_Show_Show[[]gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Show_showArrayImpl(), gopurs_runtime.RecordGet(dictShow_0, "show"))}))}
}

func Call_Main_monadWriterTuple(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
	_ = dictMonoid_0
	// TAST (Let): __local_var_1_3 shape=App(Other) bindingType=Any
	__local_var_1_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
	_ = __local_var_1_3
	// TAST (Let): applyTuple1_1_2 shape=Let(LitRecord) bindingType=(ADT ["Control","Apply","Apply"] [(ADT ["Data","Tuple","Tuple"] [(TypeVar a)])])
	applyTuple1_1_2 := (&Constructor_Control_Apply_Apply[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_Data_Tuple_functorTuple()))}
	}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer((&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_3, "append"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0), gopurs_runtime.Apply((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V1)}))}
		})
	})})
	_ = applyTuple1_1_2
	// TAST (Let): applicativeTuple1_1_1 shape=Let(LitRecord) bindingType=(ADT ["Control","Applicative","Applicative"] [(ADT ["Data","Tuple","Tuple"] [(TypeVar a)])])
	applicativeTuple1_1_1 := (&Constructor_Control_Applicative_Applicative[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyTuple1_1_2)}
	}), gopurs_runtime.Apply(Get_Data_Tuple_Tuple(), gopurs_runtime.RecordGet(dictMonoid_0, "mempty"))})
	_ = applicativeTuple1_1_1
	// TAST (Let): __local_var_2_5 shape=App(Other) bindingType=Any
	__local_var_2_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
	_ = __local_var_2_5
	// TAST (Let): applyTuple1_3_6 shape=LitRecord bindingType=(ADT ["Control","Apply","Apply"] [(ADT ["Data","Tuple","Tuple"] [(TypeVar a)])])
	applyTuple1_3_6 := (&Constructor_Control_Apply_Apply[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_Data_Tuple_functorTuple()))}
	}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer((&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_5, "append"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V0), gopurs_runtime.Apply((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V1)}))}
		})
	})})
	_ = applyTuple1_3_6
	// TAST (Let): bindTuple1_2_4 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Bind","Bind"] [(ADT ["Data","Tuple","Tuple"] [(TypeVar a)])])
	bindTuple1_2_4 := (&Constructor_Control_Bind_Bind[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyTuple1_3_6)}
	}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): v1_6_7 shape=App(Other) bindingType=(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar b)])
			v1_6_7 := gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(f_5, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1))
			_ = v1_6_7
			return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer((&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_5, "append"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0, (v1_6_7).V0), (v1_6_7).V1}))}
		})
	})})
	_ = bindTuple1_2_4
	// TAST (Let): monadTuple_1_0 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Monad","Monad"] [(ADT ["Data","Tuple","Tuple"] [(TypeVar w)])])
	monadTuple_1_0 := (&Constructor_Control_Monad_Monad[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(applicativeTuple1_1_1)}
	}), gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(bindTuple1_2_4)}
	})})
	_ = monadTuple_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 2544837208, UnsafePtr: unsafe.Pointer((&Constructor_Main_MonadWriter[gopurs_runtime.Value, *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadTuple_1_0)}
	}), gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_0))}
	}), gopurs_runtime.Func(func(w_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer((&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, w_2, Get_Data_Unit_unit()}))}
	})}))}
}

func Call_Main_monadWriterMyWriter(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
	_ = dictMonoid_0
	// TAST (Let): __local_var_1_3 shape=App(Other) bindingType=Any
	__local_var_1_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
	_ = __local_var_1_3
	// TAST (Let): applyTuple1_1_2 shape=Let(LitRecord) bindingType=(ADT ["Control","Apply","Apply"] [(ADT ["Data","Tuple","Tuple"] [(TypeVar a)])])
	applyTuple1_1_2 := (&Constructor_Control_Apply_Apply[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_Data_Tuple_functorTuple()))}
	}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer((&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_3, "append"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0), gopurs_runtime.Apply((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V1)}))}
		})
	})})
	_ = applyTuple1_1_2
	// TAST (Let): applicativeTuple1_1_1 shape=Let(LitRecord) bindingType=(ADT ["Control","Applicative","Applicative"] [(ADT ["Data","Tuple","Tuple"] [(TypeVar a)])])
	applicativeTuple1_1_1 := (&Constructor_Control_Applicative_Applicative[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyTuple1_1_2)}
	}), gopurs_runtime.Apply(Get_Data_Tuple_Tuple(), gopurs_runtime.RecordGet(dictMonoid_0, "mempty"))})
	_ = applicativeTuple1_1_1
	// TAST (Let): __local_var_2_5 shape=App(Other) bindingType=Any
	__local_var_2_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
	_ = __local_var_2_5
	// TAST (Let): applyTuple1_3_6 shape=LitRecord bindingType=(ADT ["Control","Apply","Apply"] [(ADT ["Data","Tuple","Tuple"] [(TypeVar a)])])
	applyTuple1_3_6 := (&Constructor_Control_Apply_Apply[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_Data_Tuple_functorTuple()))}
	}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer((&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_5, "append"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V0), gopurs_runtime.Apply((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V1)}))}
		})
	})})
	_ = applyTuple1_3_6
	// TAST (Let): bindTuple1_2_4 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Bind","Bind"] [(ADT ["Data","Tuple","Tuple"] [(TypeVar a)])])
	bindTuple1_2_4 := (&Constructor_Control_Bind_Bind[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyTuple1_3_6)}
	}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): v1_6_7 shape=App(Other) bindingType=(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar b)])
			v1_6_7 := gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(f_5, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1))
			_ = v1_6_7
			return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer((&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_5, "append"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0, (v1_6_7).V0), (v1_6_7).V1}))}
		})
	})})
	_ = bindTuple1_2_4
	// TAST (Let): monadTuple_1_0 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Monad","Monad"] [(ADT ["Data","Tuple","Tuple"] [(TypeVar w)])])
	monadTuple_1_0 := (&Constructor_Control_Monad_Monad[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(applicativeTuple1_1_1)}
	}), gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(bindTuple1_2_4)}
	})})
	_ = monadTuple_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 2544837208, UnsafePtr: unsafe.Pointer((&Constructor_Main_MonadWriter[gopurs_runtime.Value, *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadTuple_1_0)}
	}), gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_0))}
	}), gopurs_runtime.Func(func(w_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer((&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, w_2, Get_Data_Unit_unit()}))}
	})}))}
}

func Call_Main_monadMyWriter(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
	_ = dictMonoid_0
	// TAST (Let): __local_var_1_2 shape=App(Other) bindingType=Any
	__local_var_1_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
	_ = __local_var_1_2
	// TAST (Let): applyTuple1_1_1 shape=Let(LitRecord) bindingType=(ADT ["Control","Apply","Apply"] [(ADT ["Data","Tuple","Tuple"] [(TypeVar a)])])
	applyTuple1_1_1 := (&Constructor_Control_Apply_Apply[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_Data_Tuple_functorTuple()))}
	}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer((&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_2, "append"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0), gopurs_runtime.Apply((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V1)}))}
		})
	})})
	_ = applyTuple1_1_1
	// TAST (Let): applicativeTuple1_1_0 shape=Let(LitRecord) bindingType=(ADT ["Control","Applicative","Applicative"] [(ADT ["Data","Tuple","Tuple"] [(TypeVar a)])])
	applicativeTuple1_1_0 := (&Constructor_Control_Applicative_Applicative[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyTuple1_1_1)}
	}), gopurs_runtime.Apply(Get_Data_Tuple_Tuple(), gopurs_runtime.RecordGet(dictMonoid_0, "mempty"))})
	_ = applicativeTuple1_1_0
	// TAST (Let): __local_var_2_4 shape=App(Other) bindingType=Any
	__local_var_2_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
	_ = __local_var_2_4
	// TAST (Let): applyTuple1_3_5 shape=LitRecord bindingType=(ADT ["Control","Apply","Apply"] [(ADT ["Data","Tuple","Tuple"] [(TypeVar a)])])
	applyTuple1_3_5 := (&Constructor_Control_Apply_Apply[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_Data_Tuple_functorTuple()))}
	}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer((&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_4, "append"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V0), gopurs_runtime.Apply((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V1)}))}
		})
	})})
	_ = applyTuple1_3_5
	// TAST (Let): bindTuple1_2_3 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Bind","Bind"] [(ADT ["Data","Tuple","Tuple"] [(TypeVar a)])])
	bindTuple1_2_3 := (&Constructor_Control_Bind_Bind[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyTuple1_3_5)}
	}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): v1_6_6 shape=App(Other) bindingType=(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar b)])
			v1_6_6 := gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(f_5, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1))
			_ = v1_6_6
			return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer((&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_4, "append"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0, (v1_6_6).V0), (v1_6_6).V1}))}
		})
	})})
	_ = bindTuple1_2_3
	return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Monad[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(applicativeTuple1_1_0)}
	}), gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(bindTuple1_2_3)}
	})}))}
}

func Call_Main_bindMyWriter(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
	_ = dictSemigroup_0
	// TAST (Let): applyTuple1_1_0 shape=LitRecord bindingType=(ADT ["Control","Apply","Apply"] [(ADT ["Data","Tuple","Tuple"] [(TypeVar a)])])
	applyTuple1_1_0 := (&Constructor_Control_Apply_Apply[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_Data_Tuple_functorTuple()))}
	}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer((&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0, "append"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_2.UnsafePtr).V0), gopurs_runtime.Apply((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_2.UnsafePtr).V1)}))}
		})
	})})
	_ = applyTuple1_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyTuple1_1_0)}
	}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): v1_4_1 shape=App(Other) bindingType=(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar b)])
			v1_4_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(f_3, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1))
			_ = v1_4_1
			return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer((&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0, "append"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0, (v1_4_1).V0), (v1_4_1).V1}))}
		})
	})}))}
}

func Call_Main_applyMyWriter(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
	_ = dictSemigroup_0
	return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_Data_Tuple_functorTuple()))}
	}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer((&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0, "append"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_2.UnsafePtr).V0), gopurs_runtime.Apply((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_2.UnsafePtr).V1)}))}
		})
	})}))}
}

func Call_Main_applicativeMyWriter(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
	_ = dictMonoid_0
	// TAST (Let): __local_var_1_1 shape=App(Other) bindingType=Any
	__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
	_ = __local_var_1_1
	// TAST (Let): applyTuple1_1_0 shape=Let(LitRecord) bindingType=(ADT ["Control","Apply","Apply"] [(ADT ["Data","Tuple","Tuple"] [(TypeVar a)])])
	applyTuple1_1_0 := (&Constructor_Control_Apply_Apply[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_Data_Tuple_functorTuple()))}
	}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer((&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "append"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0), gopurs_runtime.Apply((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V1)}))}
		})
	})})
	_ = applyTuple1_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyTuple1_1_0)}
	}), gopurs_runtime.Apply(Get_Data_Tuple_Tuple(), gopurs_runtime.RecordGet(dictMonoid_0, "mempty"))}))}
}
