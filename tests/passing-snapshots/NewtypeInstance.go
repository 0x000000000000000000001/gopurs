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
		cache_Main_show = Call_Data_Show_show(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))
	})
	return cache_Main_show
}

var cache_Main_Y gopurs_runtime.Value
var once_Main_Y sync.Once

func Get_Main_Y() gopurs_runtime.Value {
	once_Main_Y.Do(func() {
		cache_Main_Y = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Array(Call_Main_Y((*(*[]gopurs_runtime.Value)((x_0_box).UnsafePtr))))
		})
	})
	return cache_Main_Y
}

var cache_Main_Singleton_dollar_Dict gopurs_runtime.Value
var once_Main_Singleton_dollar_Dict sync.Once

func Get_Main_Singleton_dollar_Dict() gopurs_runtime.Value {
	once_Main_Singleton_dollar_Dict.Do(func() {
		cache_Main_Singleton_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3773262993, UnsafePtr: unsafe.Pointer(Call_Main_Singleton_dollar_Dict(func() struct {
				singleton gopurs_runtime.Value
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					singleton gopurs_runtime.Value
				}{}
				clone.singleton = gopurs_runtime.RecordGet(orig, "singleton")
				return clone
			}()))}
		})
	})
	return cache_Main_Singleton_dollar_Dict
}

var cache_Main_ProxyArray gopurs_runtime.Value
var once_Main_ProxyArray sync.Once

func Get_Main_ProxyArray() gopurs_runtime.Value {
	once_Main_ProxyArray.Do(func() {
		cache_Main_ProxyArray = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_ProxyArray((*(*[]gopurs_runtime.Value)((x_0_box).UnsafePtr)))
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
			return func() gopurs_runtime.Value {
				_v := Call_Main_MyWriter(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](x_0_box))
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()
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
			return gopurs_runtime.Array(Call_Main_MyArray((*(*[]gopurs_runtime.Value)((x_0_box).UnsafePtr))))
		})
	})
	return cache_Main_MyArray
}

var cache_Main_MyArray__4020493786 gopurs_runtime.Value
var once_Main_MyArray__4020493786 sync.Once

func Get_Main_MyArray__4020493786() gopurs_runtime.Value {
	once_Main_MyArray__4020493786.Do(func() {
		cache_Main_MyArray__4020493786 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				arr := Call_Main_MyArray__4020493786(func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(x_0_box.UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr {
						unboxed[i] = v.IntVal
					}
					return unboxed
				}())
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Int(v)
				}
				return gopurs_runtime.Array(boxed)
			}()
		})
	})
	return cache_Main_MyArray__4020493786
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
			return gopurs_runtime.Value{Type: 9, IntVal: 2544837208, UnsafePtr: unsafe.Pointer(Call_Main_MonadWriter_dollar_Dict(func() struct {
				Monad0  gopurs_runtime.Value
				Monoid1 gopurs_runtime.Value
				tell    gopurs_runtime.Value
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Monad0  gopurs_runtime.Value
					Monoid1 gopurs_runtime.Value
					tell    gopurs_runtime.Value
				}{}
				clone.Monad0 = gopurs_runtime.RecordGet(orig, "Monad0")
				clone.Monoid1 = gopurs_runtime.RecordGet(orig, "Monoid1")
				clone.tell = gopurs_runtime.RecordGet(orig, "tell")
				return clone
			}()))}
		})
	})
	return cache_Main_MonadWriter_dollar_Dict
}

var cache_Main_Foo gopurs_runtime.Value
var once_Main_Foo sync.Once

func Get_Main_Foo() gopurs_runtime.Value {
	once_Main_Foo.Do(func() {
		cache_Main_Foo = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_Foo(uint32(x_0_box.IntVal))), UnsafePtr: nil}
		})
	})
	return cache_Main_Foo
}

var cache_Main_functorProxy2 gopurs_runtime.Value
var once_Main_functorProxy2 sync.Once

func Get_Main_functorProxy2() gopurs_runtime.Value {
	once_Main_functorProxy2.Do(func() {
		cache_Main_functorProxy2 = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Main_562069347_2812149806((&Constructor_Data_Functor_Functor[uint32]{1, gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, m_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(4225449536), UnsafePtr: nil}
		})})))}
	})
	return cache_Main_functorProxy2
}

var cache_Main_functorFoo gopurs_runtime.Value
var once_Main_functorFoo sync.Once

func Get_Main_functorFoo() gopurs_runtime.Value {
	once_Main_functorFoo.Do(func() {
		cache_Main_functorFoo = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Main_562069347_2812149806(Rebox_Main_2812149806_562069347(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Main_functorProxy2()))))}
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
		cache_Main_singletonArray = gopurs_runtime.Value{Type: 9, IntVal: 3773262993, UnsafePtr: unsafe.Pointer(Rebox_Main_2655323845_2269342371((&Constructor_Main_Singleton[gopurs_runtime.Value, []gopurs_runtime.Value]{1, gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.Array([]gopurs_runtime.Value{x_0})).UnsafePtr)))
		})})))}
	})
	return cache_Main_singletonArray
}

var cache_Main_singletonY gopurs_runtime.Value
var once_Main_singletonY sync.Once

func Get_Main_singletonY() gopurs_runtime.Value {
	once_Main_singletonY.Do(func() {
		cache_Main_singletonY = gopurs_runtime.Value{Type: 9, IntVal: 3773262993, UnsafePtr: unsafe.Pointer(Rebox_Main_2655323845_2269342371(Rebox_Main_2269342371_2655323845(gopurs_runtime.CoerceToStruct[Constructor_Main_Singleton[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Main_singletonArray()))))}
	})
	return cache_Main_singletonY
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

var cache_Main_singleton__125190524 gopurs_runtime.Value
var once_Main_singleton__125190524 sync.Once

func Get_Main_singleton__125190524() gopurs_runtime.Value {
	once_Main_singleton__125190524.Do(func() {
		cache_Main_singleton__125190524 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				arr := Call_Main_singleton__125190524(__eta_norm_0_0_box.StrVal())
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Str(v)
				}
				return gopurs_runtime.Array(boxed)
			}()
		})
	})
	return cache_Main_singleton__125190524
}

var cache_Main_showY gopurs_runtime.Value
var once_Main_showY sync.Once

func Get_Main_showY() gopurs_runtime.Value {
	once_Main_showY.Do(func() {
		cache_Main_showY = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1953100407_1386611502(Rebox_Main_1386611502_1953100407(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Call_Data_Show_showArray(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString()))))})))))}
	})
	return cache_Main_showY
}

var cache_Main_showX gopurs_runtime.Value
var once_Main_showX sync.Once

func Get_Main_showX() gopurs_runtime.Value {
	once_Main_showX.Do(func() {
		cache_Main_showX = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString()))))}
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
		cache_Main_showMyArray1 = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1953100407_1386611502(Rebox_Main_1386611502_1953100407(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Call_Main_showMyArray(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString()))))})))))}
	})
	return cache_Main_showMyArray1
}

var cache_Main_ordX gopurs_runtime.Value
var once_Main_ordX sync.Once

func Get_Main_ordX() gopurs_runtime.Value {
	once_Main_ordX.Do(func() {
		cache_Main_ordX = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Main_2406510097_4177771502(Rebox_Main_4177771502_2406510097(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordString()))))}
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
		cache_Main_functorMyWriter = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Main_2363162019_2812149806(Rebox_Main_2812149806_2363162019(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Tuple_functorTuple()))))}
	})
	return cache_Main_functorMyWriter
}

var cache_Main_functorSyn gopurs_runtime.Value
var once_Main_functorSyn sync.Once

func Get_Main_functorSyn() gopurs_runtime.Value {
	once_Main_functorSyn.Do(func() {
		cache_Main_functorSyn = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Main_functorMyWriter()))}
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

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showStringImpl(), gopurs_runtime.Str("test")).StrVal())), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Data_Show_showArray(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString()))))}), "show"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Main_singletonY(), "singleton"), gopurs_runtime.Str("test"))).StrVal())), gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Main_showMyArray(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString()))))}), "show"), func() gopurs_runtime.Value {
					arr_val_arrayMap7 := func() gopurs_runtime.Value {
						arr := []int64{int64(1), int64(2), int64(3)}
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Int(v)
						}
						return gopurs_runtime.Array(boxed)
					}()
					_ = arr_val_arrayMap7
					arr_go_arrayMap7 := (*[]gopurs_runtime.Value)(arr_val_arrayMap7.UnsafePtr)
					_ = arr_go_arrayMap7
					res_go_arrayMap7 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap7))
					_ = res_go_arrayMap7
					for i_arrayMap7, v_arrayMap7 := range *arr_go_arrayMap7 {
						res_go_arrayMap7[i_arrayMap7] = gopurs_runtime.Apply(Call_Data_Show_show(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), v_arrayMap7)
					}
					return gopurs_runtime.Array(res_go_arrayMap7)
				}()).StrVal())), gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
				}))
			}))
		}))
	})
	return cache_Main_main
}

var cache_Main_eqX gopurs_runtime.Value
var once_Main_eqX sync.Once

func Get_Main_eqX() gopurs_runtime.Value {
	once_Main_eqX.Do(func() {
		cache_Main_eqX = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))}
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
		c := (*Constructor_Main_Singleton[gopurs_runtime.Value, gopurs_runtime.Value])(ptr)
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
		c := (*Constructor_Main_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value])(ptr)
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

func Call_Main_Y(x_0_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
	var x_0 []gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_Singleton_dollar_Dict(x_0_loop struct {
	singleton gopurs_runtime.Value
}) *Constructor_Main_Singleton[gopurs_runtime.Value, gopurs_runtime.Value] {
	var x_0 struct {
		singleton gopurs_runtime.Value
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Main_Singleton[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict1("singleton", orig.singleton)
	}())
}

func Call_Main_ProxyArray(x_0_loop []gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 []gopurs_runtime.Value = x_0_loop
	_ = x_0
	return gopurs_runtime.Array(x_0)
}

func Call_Main_MyWriter(x_0_loop *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
} {
	var x_0 *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] = x_0_loop
	_ = x_0
	return func() struct {
		V0 gopurs_runtime.Value
		V1 gopurs_runtime.Value
	} { _v := gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(x_0)}; _p := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr); return struct {
		V0 gopurs_runtime.Value
		V1 gopurs_runtime.Value
	}{V0: _p.V0, V1: _p.V1} }()
}

func Call_Main_X(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_MyArray(x_0_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
	var x_0 []gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_MyArray__4020493786(x_0_loop []int64) []int64 {
MyArray__4020493786:
	for {
		if false {
			continue MyArray__4020493786
		}
		var x_0 []int64 = x_0_loop
		_ = x_0
		return x_0
	}
}

func Call_Main_Syn(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_MonadWriter_dollar_Dict(x_0_loop struct {
	Monad0  gopurs_runtime.Value
	Monoid1 gopurs_runtime.Value
	tell    gopurs_runtime.Value
}) *Constructor_Main_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value] {
	var x_0 struct {
		Monad0  gopurs_runtime.Value
		Monoid1 gopurs_runtime.Value
		tell    gopurs_runtime.Value
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Main_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict3("Monad0", "Monoid1", "tell", orig.Monad0, orig.Monoid1, orig.tell)
	}())
}

func Call_Main_Foo(x_0_loop uint32) uint32 {
	var x_0 uint32 = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_tell(dict_0_loop *Constructor_Main_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V2
}

func Call_Main_singleton(dict_0_loop *Constructor_Main_Singleton[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_Singleton[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V0
}

func Call_Main_singleton__125190524(__eta_norm_0_0_loop string) []string {
singleton__125190524:
	for {
		if false {
			continue singleton__125190524
		}
		var __eta_norm_0_0 string = __eta_norm_0_0_loop
		_ = __eta_norm_0_0
		return func() []string {
			arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Main_singletonY(), "singleton"), gopurs_runtime.Str(__eta_norm_0_0)).UnsafePtr)
			unboxed := make([]string, len(arr))
			for i, v := range arr {
				unboxed[i] = v.StrVal()
			}
			return unboxed
		}()
	}
}

func Call_Main_showMyArray(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
	_ = dictShow_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1356436936_1386611502(Rebox_Main_1386611502_1356436936(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Call_Data_Show_showArray(dictShow_0)))))}
}

func Call_Main_monadWriterTuple(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
	_ = dictMonoid_0
	// TAST (Let): monadTuple_1_0 shape=App(Var) bindingType=(ADT ["Control","Monad","Monad"] [(ADT ["Data","Tuple","Tuple"] [(TypeVar w$scope24)])])
	monadTuple_1_0 := Rebox_Main_2568689657_2157788756(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](Call_Data_Tuple_monadTuple(dictMonoid_0)))
	_ = monadTuple_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 2544837208, UnsafePtr: unsafe.Pointer(Rebox_Main_2915227655_72788106((&Constructor_Main_MonadWriter[gopurs_runtime.Value, *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(Rebox_Main_2157788756_2568689657(monadTuple_1_0))}
	}), gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_0))}
	}), gopurs_runtime.Func(func(w_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
			}{w_2, Get_Data_Unit_unit()}
			return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
		}()))}
	})})))}
}

func Call_Main_monadWriterMyWriter(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
	_ = dictMonoid_0
	return gopurs_runtime.Value{Type: 9, IntVal: 2544837208, UnsafePtr: unsafe.Pointer(Rebox_Main_2915227655_72788106(Rebox_Main_72788106_2915227655(gopurs_runtime.CoerceToStruct[Constructor_Main_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value]](Call_Main_monadWriterTuple(dictMonoid_0)))))}
}

func Call_Main_monadMyWriter(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
	_ = dictMonoid_0
	return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(Rebox_Main_2157788756_2568689657(Rebox_Main_2568689657_2157788756(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](Call_Data_Tuple_monadTuple(dictMonoid_0)))))}
}

func Call_Main_bindMyWriter(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
	_ = dictSemigroup_0
	return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(Rebox_Main_2495815764_2748095225(Rebox_Main_2748095225_2495815764(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Call_Data_Tuple_bindTuple(dictSemigroup_0)))))}
}

func Call_Main_applyMyWriter(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
	_ = dictSemigroup_0
	return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(Rebox_Main_2708463828_3741347833(Rebox_Main_3741347833_2708463828(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Call_Data_Tuple_applyTuple(dictSemigroup_0)))))}
}

func Call_Main_applicativeMyWriter(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
	_ = dictMonoid_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(Rebox_Main_894579924_1439734649(Rebox_Main_1439734649_894579924(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Call_Data_Tuple_applicativeTuple(dictMonoid_0)))))}
}

func Rebox_Main_1140313009_3790796878(in *Constructor_Data_Eq_Eq[string]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_1356436936_1386611502(in *Constructor_Data_Show_Show[[]gopurs_runtime.Value]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_1386611502_1356436936(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[[]gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[[]gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_1386611502_1514099793(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[string] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[string]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_1386611502_1636311157(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[int64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[int64]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_1386611502_1953100407(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[[]string] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[[]string]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_1439734649_894579924(in *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]) *Constructor_Control_Applicative_Applicative[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	out := &Constructor_Control_Applicative_Applicative[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{}
	out.V0 = in.V0
	out.V1 = in.V1
	return out
}

func Rebox_Main_1514099793_1386611502(in *Constructor_Data_Show_Show[string]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_1636311157_1386611502(in *Constructor_Data_Show_Show[int64]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_1953100407_1386611502(in *Constructor_Data_Show_Show[[]string]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_2157788756_2568689657(in *Constructor_Control_Monad_Monad[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Control_Monad_Monad[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Control_Monad_Monad[gopurs_runtime.Value]{}
	out.V0 = in.V0
	out.V1 = in.V1
	return out
}

func Rebox_Main_2269342371_2655323845(in *Constructor_Main_Singleton[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Main_Singleton[gopurs_runtime.Value, []gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_Singleton[gopurs_runtime.Value, []gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_2363162019_2812149806(in *Constructor_Data_Functor_Functor[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Functor_Functor[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Functor_Functor[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_2406510097_4177771502(in *Constructor_Data_Ord_Ord[string]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Ord_Ord[gopurs_runtime.Value]{}
	out.V0 = in.V0
	out.V1 = in.V1
	return out
}

func Rebox_Main_2495815764_2748095225(in *Constructor_Control_Bind_Bind[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Control_Bind_Bind[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Control_Bind_Bind[gopurs_runtime.Value]{}
	out.V0 = in.V0
	out.V1 = in.V1
	return out
}

func Rebox_Main_2568689657_2157788756(in *Constructor_Control_Monad_Monad[gopurs_runtime.Value]) *Constructor_Control_Monad_Monad[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	out := &Constructor_Control_Monad_Monad[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{}
	out.V0 = in.V0
	out.V1 = in.V1
	return out
}

func Rebox_Main_2655323845_2269342371(in *Constructor_Main_Singleton[gopurs_runtime.Value, []gopurs_runtime.Value]) *Constructor_Main_Singleton[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_Singleton[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_2708463828_3741347833(in *Constructor_Control_Apply_Apply[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Control_Apply_Apply[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Control_Apply_Apply[gopurs_runtime.Value]{}
	out.V0 = in.V0
	out.V1 = in.V1
	return out
}

func Rebox_Main_2748095225_2495815764(in *Constructor_Control_Bind_Bind[gopurs_runtime.Value]) *Constructor_Control_Bind_Bind[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	out := &Constructor_Control_Bind_Bind[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{}
	out.V0 = in.V0
	out.V1 = in.V1
	return out
}

func Rebox_Main_2812149806_2363162019(in *Constructor_Data_Functor_Functor[gopurs_runtime.Value]) *Constructor_Data_Functor_Functor[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Functor_Functor[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_2812149806_562069347(in *Constructor_Data_Functor_Functor[gopurs_runtime.Value]) *Constructor_Data_Functor_Functor[uint32] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Functor_Functor[uint32]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_2915227655_72788106(in *Constructor_Main_MonadWriter[gopurs_runtime.Value, *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Main_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = in.V0
	out.V1 = in.V1
	out.V2 = in.V2
	return out
}

func Rebox_Main_3741347833_2708463828(in *Constructor_Control_Apply_Apply[gopurs_runtime.Value]) *Constructor_Control_Apply_Apply[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	out := &Constructor_Control_Apply_Apply[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{}
	out.V0 = in.V0
	out.V1 = in.V1
	return out
}

func Rebox_Main_3790796878_1140313009(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[string] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Eq[string]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_4177771502_2406510097(in *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) *Constructor_Data_Ord_Ord[string] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Ord_Ord[string]{}
	out.V0 = in.V0
	out.V1 = in.V1
	return out
}

func Rebox_Main_562069347_2812149806(in *Constructor_Data_Functor_Functor[uint32]) *Constructor_Data_Functor_Functor[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Functor_Functor[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_72788106_2915227655(in *Constructor_Main_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Main_MonadWriter[gopurs_runtime.Value, *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_MonadWriter[gopurs_runtime.Value, *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{}
	out.V0 = in.V0
	out.V1 = in.V1
	out.V2 = in.V2
	return out
}

func Rebox_Main_894579924_1439734649(in *Constructor_Control_Applicative_Applicative[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{}
	out.V0 = in.V0
	out.V1 = in.V1
	return out
}
