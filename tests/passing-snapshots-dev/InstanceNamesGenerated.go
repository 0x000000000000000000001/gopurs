package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_GenericFoo gopurs_runtime.Value
var once_Main_GenericFoo sync.Once

func Get_Main_GenericFoo() gopurs_runtime.Value {
	once_Main_GenericFoo.Do(func() {
		cache_Main_GenericFoo = gopurs_runtime.Value{Type: 9, IntVal: int64(682342953), UnsafePtr: nil}
	})
	return cache_Main_GenericFoo
}

var cache_Main_Left gopurs_runtime.Value
var once_Main_Left sync.Once

func Get_Main_Left() gopurs_runtime.Value {
	once_Main_Left.Do(func() {
		cache_Main_Left = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 913637797, UnsafePtr: unsafe.Pointer((&Constructor_Main_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, value0}))}
		})
	})
	return cache_Main_Left
}

var cache_Main_Right gopurs_runtime.Value
var once_Main_Right sync.Once

func Get_Main_Right() gopurs_runtime.Value {
	once_Main_Right.Do(func() {
		cache_Main_Right = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2535318782, UnsafePtr: unsafe.Pointer((&Constructor_Main_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, value0}))}
		})
	})
	return cache_Main_Right
}

var cache_Main_Foo gopurs_runtime.Value
var once_Main_Foo sync.Once

func Get_Main_Foo() gopurs_runtime.Value {
	once_Main_Foo.Do(func() {
		cache_Main_Foo = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 2763139640, UnsafePtr: unsafe.Pointer((&Constructor_Main_Foo[gopurs_runtime.Value, gopurs_runtime.Value]{1, value0, value1}))}
			})
		})
	})
	return cache_Main_Foo
}

var cache_Main_reservedWordFunction gopurs_runtime.Value
var once_Main_reservedWordFunction sync.Once

func Get_Main_reservedWordFunction() gopurs_runtime.Value {
	once_Main_reservedWordFunction.Do(func() {
		cache_Main_reservedWordFunction = func() gopurs_runtime.Value {
			orig := struct {

				/*gopurs:record-v1:0:*/
			}{}
			_ = orig
			return gopurs_runtime.RecordDict0()
		}()
	})
	return cache_Main_reservedWordFunction
}

var cache_Main_reservedWordArrow gopurs_runtime.Value
var once_Main_reservedWordArrow sync.Once

func Get_Main_reservedWordArrow() gopurs_runtime.Value {
	once_Main_reservedWordArrow.Do(func() {
		cache_Main_reservedWordArrow = func() gopurs_runtime.Value {
			orig := struct {

				/*gopurs:record-v1:0:*/
			}{}
			_ = orig
			return gopurs_runtime.RecordDict0()
		}()
	})
	return cache_Main_reservedWordArrow
}

var cache_Main_overlappingStillCompiles gopurs_runtime.Value
var once_Main_overlappingStillCompiles sync.Once

func Get_Main_overlappingStillCompiles() gopurs_runtime.Value {
	once_Main_overlappingStillCompiles.Do(func() {
		cache_Main_overlappingStillCompiles = func() gopurs_runtime.Value {
			orig := struct {

				/*gopurs:record-v1:0:*/
			}{}
			_ = orig
			return gopurs_runtime.RecordDict0()
		}()
	})
	return cache_Main_overlappingStillCompiles
}

var cache_Main_overlappingStillCompiles1 gopurs_runtime.Value
var once_Main_overlappingStillCompiles1 sync.Once

func Get_Main_overlappingStillCompiles1() gopurs_runtime.Value {
	once_Main_overlappingStillCompiles1.Do(func() {
		cache_Main_overlappingStillCompiles1 = func() gopurs_runtime.Value {
			orig := struct {

				/*gopurs:record-v1:0:*/
			}{}
			_ = orig
			return gopurs_runtime.RecordDict0()
		}()
	})
	return cache_Main_overlappingStillCompiles1
}

var cache_Main_oneTypeParamChainString gopurs_runtime.Value
var once_Main_oneTypeParamChainString sync.Once

func Get_Main_oneTypeParamChainString() gopurs_runtime.Value {
	once_Main_oneTypeParamChainString.Do(func() {
		cache_Main_oneTypeParamChainString = func() gopurs_runtime.Value {
			orig := struct {

				/*gopurs:record-v1:0:*/
			}{}
			_ = orig
			return gopurs_runtime.RecordDict0()
		}()
	})
	return cache_Main_oneTypeParamChainString
}

var cache_Main_oneTypeParamChainBoolean gopurs_runtime.Value
var once_Main_oneTypeParamChainBoolean sync.Once

func Get_Main_oneTypeParamChainBoolean() gopurs_runtime.Value {
	once_Main_oneTypeParamChainBoolean.Do(func() {
		cache_Main_oneTypeParamChainBoolean = func() gopurs_runtime.Value {
			orig := struct {

				/*gopurs:record-v1:0:*/
			}{}
			_ = orig
			return gopurs_runtime.RecordDict0()
		}()
	})
	return cache_Main_oneTypeParamChainBoolean
}

var cache_Main_oneTypeParamBoolean gopurs_runtime.Value
var once_Main_oneTypeParamBoolean sync.Once

func Get_Main_oneTypeParamBoolean() gopurs_runtime.Value {
	once_Main_oneTypeParamBoolean.Do(func() {
		cache_Main_oneTypeParamBoolean = func() gopurs_runtime.Value {
			orig := struct {

				/*gopurs:record-v1:0:*/
			}{}
			_ = orig
			return gopurs_runtime.RecordDict0()
		}()
	})
	return cache_Main_oneTypeParamBoolean
}

var cache_Main_noTypeParams gopurs_runtime.Value
var once_Main_noTypeParams sync.Once

func Get_Main_noTypeParams() gopurs_runtime.Value {
	once_Main_noTypeParams.Do(func() {
		cache_Main_noTypeParams = func() gopurs_runtime.Value {
			orig := struct {

				/*gopurs:record-v1:0:*/
			}{}
			_ = orig
			return gopurs_runtime.RecordDict0()
		}()
	})
	return cache_Main_noTypeParams
}

var cache_Main_multipleTypeParamsChainBo gopurs_runtime.Value
var once_Main_multipleTypeParamsChainBo sync.Once

func Get_Main_multipleTypeParamsChainBo() gopurs_runtime.Value {
	once_Main_multipleTypeParamsChainBo.Do(func() {
		cache_Main_multipleTypeParamsChainBo = func() gopurs_runtime.Value {
			orig := struct {

				/*gopurs:record-v1:0:*/
			}{}
			_ = orig
			return gopurs_runtime.RecordDict0()
		}()
	})
	return cache_Main_multipleTypeParamsChainBo
}

var cache_Main_multipleTypeParamsChainBo1 gopurs_runtime.Value
var once_Main_multipleTypeParamsChainBo1 sync.Once

func Get_Main_multipleTypeParamsChainBo1() gopurs_runtime.Value {
	once_Main_multipleTypeParamsChainBo1.Do(func() {
		cache_Main_multipleTypeParamsChainBo1 = func() gopurs_runtime.Value {
			orig := struct {

				/*gopurs:record-v1:0:*/
			}{}
			_ = orig
			return gopurs_runtime.RecordDict0()
		}()
	})
	return cache_Main_multipleTypeParamsChainBo1
}

var cache_Main_multipleTypeParamsChainBo2 gopurs_runtime.Value
var once_Main_multipleTypeParamsChainBo2 sync.Once

func Get_Main_multipleTypeParamsChainBo2() gopurs_runtime.Value {
	once_Main_multipleTypeParamsChainBo2.Do(func() {
		cache_Main_multipleTypeParamsChainBo2 = func() gopurs_runtime.Value {
			orig := struct {

				/*gopurs:record-v1:0:*/
			}{}
			_ = orig
			return gopurs_runtime.RecordDict0()
		}()
	})
	return cache_Main_multipleTypeParamsChainBo2
}

var cache_Main_multipleTypeParamsChainBo3 gopurs_runtime.Value
var once_Main_multipleTypeParamsChainBo3 sync.Once

func Get_Main_multipleTypeParamsChainBo3() gopurs_runtime.Value {
	once_Main_multipleTypeParamsChainBo3.Do(func() {
		cache_Main_multipleTypeParamsChainBo3 = func() gopurs_runtime.Value {
			orig := struct {

				/*gopurs:record-v1:0:*/
			}{}
			_ = orig
			return gopurs_runtime.RecordDict0()
		}()
	})
	return cache_Main_multipleTypeParamsChainBo3
}

var cache_Main_multipleTypeParamsChainBo4 gopurs_runtime.Value
var once_Main_multipleTypeParamsChainBo4 sync.Once

func Get_Main_multipleTypeParamsChainBo4() gopurs_runtime.Value {
	once_Main_multipleTypeParamsChainBo4.Do(func() {
		cache_Main_multipleTypeParamsChainBo4 = func() gopurs_runtime.Value {
			orig := struct {

				/*gopurs:record-v1:0:*/
			}{}
			_ = orig
			return gopurs_runtime.RecordDict0()
		}()
	})
	return cache_Main_multipleTypeParamsChainBo4
}

var cache_Main_multipleTypeParamsBoolean gopurs_runtime.Value
var once_Main_multipleTypeParamsBoolean sync.Once

func Get_Main_multipleTypeParamsBoolean() gopurs_runtime.Value {
	once_Main_multipleTypeParamsBoolean.Do(func() {
		cache_Main_multipleTypeParamsBoolean = func() gopurs_runtime.Value {
			orig := struct {

				/*gopurs:record-v1:0:*/
			}{}
			_ = orig
			return gopurs_runtime.RecordDict0()
		}()
	})
	return cache_Main_multipleTypeParamsBoolean
}

var cache_Main_multipleKindParamsConstru gopurs_runtime.Value
var once_Main_multipleKindParamsConstru sync.Once

func Get_Main_multipleKindParamsConstru() gopurs_runtime.Value {
	once_Main_multipleKindParamsConstru.Do(func() {
		cache_Main_multipleKindParamsConstru = func() gopurs_runtime.Value {
			orig := struct {

				/*gopurs:record-v1:0:*/
			}{}
			_ = orig
			return gopurs_runtime.RecordDict0()
		}()
	})
	return cache_Main_multipleKindParamsConstru
}

var cache_Main_multipleKindParamsChainCo gopurs_runtime.Value
var once_Main_multipleKindParamsChainCo sync.Once

func Get_Main_multipleKindParamsChainCo() gopurs_runtime.Value {
	once_Main_multipleKindParamsChainCo.Do(func() {
		cache_Main_multipleKindParamsChainCo = func() gopurs_runtime.Value {
			orig := struct {

				/*gopurs:record-v1:0:*/
			}{}
			_ = orig
			return gopurs_runtime.RecordDict0()
		}()
	})
	return cache_Main_multipleKindParamsChainCo
}

var cache_Main_multipleKindParamsChainCo1 gopurs_runtime.Value
var once_Main_multipleKindParamsChainCo1 sync.Once

func Get_Main_multipleKindParamsChainCo1() gopurs_runtime.Value {
	once_Main_multipleKindParamsChainCo1.Do(func() {
		cache_Main_multipleKindParamsChainCo1 = func() gopurs_runtime.Value {
			orig := struct {

				/*gopurs:record-v1:0:*/
			}{}
			_ = orig
			return gopurs_runtime.RecordDict0()
		}()
	})
	return cache_Main_multipleKindParamsChainCo1
}

var cache_Main_multipleKindParamsChainCo2 gopurs_runtime.Value
var once_Main_multipleKindParamsChainCo2 sync.Once

func Get_Main_multipleKindParamsChainCo2() gopurs_runtime.Value {
	once_Main_multipleKindParamsChainCo2.Do(func() {
		cache_Main_multipleKindParamsChainCo2 = func() gopurs_runtime.Value {
			orig := struct {

				/*gopurs:record-v1:0:*/
			}{}
			_ = orig
			return gopurs_runtime.RecordDict0()
		}()
	})
	return cache_Main_multipleKindParamsChainCo2
}

var cache_Main_higherKindedTypeParamsCha gopurs_runtime.Value
var once_Main_higherKindedTypeParamsCha sync.Once

func Get_Main_higherKindedTypeParamsCha() gopurs_runtime.Value {
	once_Main_higherKindedTypeParamsCha.Do(func() {
		cache_Main_higherKindedTypeParamsCha = gopurs_runtime.Value{Type: 9, IntVal: 1868017351, UnsafePtr: unsafe.Pointer((&Constructor_Main_HigherKindedTypeParamsChain[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(int64(0))
		})}))}
	})
	return cache_Main_higherKindedTypeParamsCha
}

var cache_Main_higherKindedTypeParamsCha1 gopurs_runtime.Value
var once_Main_higherKindedTypeParamsCha1 sync.Once

func Get_Main_higherKindedTypeParamsCha1() gopurs_runtime.Value {
	once_Main_higherKindedTypeParamsCha1.Do(func() {
		cache_Main_higherKindedTypeParamsCha1 = gopurs_runtime.Value{Type: 9, IntVal: 1868017351, UnsafePtr: unsafe.Pointer((&Constructor_Main_HigherKindedTypeParamsChain[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(int64(0))
		})}))}
	})
	return cache_Main_higherKindedTypeParamsCha1
}

var cache_Main_higherKindedTypeParamsArr gopurs_runtime.Value
var once_Main_higherKindedTypeParamsArr sync.Once

func Get_Main_higherKindedTypeParamsArr() gopurs_runtime.Value {
	once_Main_higherKindedTypeParamsArr.Do(func() {
		cache_Main_higherKindedTypeParamsArr = gopurs_runtime.Value{Type: 9, IntVal: 3032179466, UnsafePtr: unsafe.Pointer((&Constructor_Main_HigherKindedTypeParams[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(int64(0))
		})}))}
	})
	return cache_Main_higherKindedTypeParamsArr
}

var cache_Main_genericGenericFoo_ gopurs_runtime.Value
var once_Main_genericGenericFoo_ sync.Once

func Get_Main_genericGenericFoo_() gopurs_runtime.Value {
	once_Main_genericGenericFoo_.Do(func() {
		cache_Main_genericGenericFoo_ = gopurs_runtime.Value{Type: 9, IntVal: 1921946594, UnsafePtr: unsafe.Pointer(Rebox_Main_855013872_2818661616((&Constructor_Data_Generic_Rep_Generic[uint32, uint32]{1, gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(1454898258), UnsafePtr: nil}
		}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(682342953), UnsafePtr: nil}
		})})))}
	})
	return cache_Main_genericGenericFoo_
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=Any
			__local_var_0_0 := Call_Lib_namedExportStillWorksUnit(int64(0))
			_ = __local_var_0_0
			_dollar___unused_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = _dollar___unused_1_1
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

var cache_Main_hktpChain gopurs_runtime.Value
var once_Main_hktpChain sync.Once

func Get_Main_hktpChain() gopurs_runtime.Value {
	once_Main_hktpChain.Do(func() {
		cache_Main_hktpChain = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_hktpChain(gopurs_runtime.CoerceToStruct[Constructor_Main_HigherKindedTypeParamsChain[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Main_hktpChain
}

var cache_Main_hktp gopurs_runtime.Value
var once_Main_hktp sync.Once

func Get_Main_hktp() gopurs_runtime.Value {
	once_Main_hktp.Do(func() {
		cache_Main_hktp = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_hktp(gopurs_runtime.CoerceToStruct[Constructor_Main_HigherKindedTypeParams[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Main_hktp
}

type Constructor_Main_GenericFoo struct {
	Rc uint32
}

type Constructor_Main_Left[T_l any, T_r any] struct {
	Rc uint32
	V0 T_l
}

type Constructor_Main_Right[T_l any, T_r any] struct {
	Rc uint32
	V0 T_r
}

type Constructor_Main_Foo[T_a any, T_b any] struct {
	Rc uint32
	V0 T_a
	V1 T_b
}

type Constructor_Main_ReservedWord[T_a any] struct {
	Rc uint32
}

//gopurs:init-v1:00670065006e006500720061007400650064002f004d00610069006e:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f004d00610069006e005f005200650073006500720076006500640057006f00720064
func init() {
	gopurs_runtime.StructGetters[3965187028] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_ReservedWord[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {

		default:
			panic("Key not found in dictionary Constructor_Main_ReservedWord: " + key)
		}
	}
}

type Constructor_Main_OverlappingStillCompiles[T_a any] struct {
	Rc uint32
}

//gopurs:init-v1:00670065006e006500720061007400650064002f004d00610069006e:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f004d00610069006e005f004f007600650072006c0061007000700069006e0067005300740069006c006c0043006f006d00700069006c00650073
func init() {
	gopurs_runtime.StructGetters[1985592625] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_OverlappingStillCompiles[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {

		default:
			panic("Key not found in dictionary Constructor_Main_OverlappingStillCompiles: " + key)
		}
	}
}

type Constructor_Main_OneTypeParamChain[T_a any] struct {
	Rc uint32
}

//gopurs:init-v1:00670065006e006500720061007400650064002f004d00610069006e:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f004d00610069006e005f004f006e006500540079007000650050006100720061006d0043006800610069006e
func init() {
	gopurs_runtime.StructGetters[4048695264] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_OneTypeParamChain[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {

		default:
			panic("Key not found in dictionary Constructor_Main_OneTypeParamChain: " + key)
		}
	}
}

type Constructor_Main_OneTypeParam[T_a any] struct {
	Rc uint32
}

//gopurs:init-v1:00670065006e006500720061007400650064002f004d00610069006e:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f004d00610069006e005f004f006e006500540079007000650050006100720061006d
func init() {
	gopurs_runtime.StructGetters[1734428973] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_OneTypeParam[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {

		default:
			panic("Key not found in dictionary Constructor_Main_OneTypeParam: " + key)
		}
	}
}

type Constructor_Main_NoTypeParams struct {
	Rc uint32
}

//gopurs:init-v1:00670065006e006500720061007400650064002f004d00610069006e:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f004d00610069006e005f004e006f00540079007000650050006100720061006d0073
func init() {
	gopurs_runtime.StructGetters[224514203] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_NoTypeParams)(ptr)
		_ = c
		switch key {

		default:
			panic("Key not found in dictionary Constructor_Main_NoTypeParams: " + key)
		}
	}
}

type Constructor_Main_MultipleTypeParamsChain[T_a any, T_b any, T_c any, T_d any, T_e any] struct {
	Rc uint32
}

//gopurs:init-v1:00670065006e006500720061007400650064002f004d00610069006e:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f004d00610069006e005f004d0075006c007400690070006c006500540079007000650050006100720061006d00730043006800610069006e
func init() {
	gopurs_runtime.StructGetters[3346744871] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_MultipleTypeParamsChain[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(ptr)
		_ = c
		switch key {

		default:
			panic("Key not found in dictionary Constructor_Main_MultipleTypeParamsChain: " + key)
		}
	}
}

type Constructor_Main_MultipleTypeParams[T_a any, T_b any, T_c any, T_d any, T_e any] struct {
	Rc uint32
}

//gopurs:init-v1:00670065006e006500720061007400650064002f004d00610069006e:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f004d00610069006e005f004d0075006c007400690070006c006500540079007000650050006100720061006d0073
func init() {
	gopurs_runtime.StructGetters[1755820650] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_MultipleTypeParams[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(ptr)
		_ = c
		switch key {

		default:
			panic("Key not found in dictionary Constructor_Main_MultipleTypeParams: " + key)
		}
	}
}

type Constructor_Main_HigherKindedTypeParamsChain[T_f any, T_g any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

//gopurs:init-v1:00670065006e006500720061007400650064002f004d00610069006e:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f004d00610069006e005f004800690067006800650072004b0069006e00640065006400540079007000650050006100720061006d00730043006800610069006e
func init() {
	gopurs_runtime.StructGetters[1868017351] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_HigherKindedTypeParamsChain[gopurs_runtime.Value, gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "hktpChain":
			return gopurs_runtime.Box(c.V0)
		default:
			panic("Key not found in dictionary Constructor_Main_HigherKindedTypeParamsChain: " + key)
		}
	}
}

type Constructor_Main_HigherKindedTypeParams[T_f any, T_g any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

//gopurs:init-v1:00670065006e006500720061007400650064002f004d00610069006e:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f004d00610069006e005f004800690067006800650072004b0069006e00640065006400540079007000650050006100720061006d0073
func init() {
	gopurs_runtime.StructGetters[3032179466] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_HigherKindedTypeParams[gopurs_runtime.Value, gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "hktp":
			return gopurs_runtime.Box(c.V0)
		default:
			panic("Key not found in dictionary Constructor_Main_HigherKindedTypeParams: " + key)
		}
	}
}

type Constructor_Main_MultipleKindParams[T_customKind any] struct {
	Rc uint32
}

//gopurs:init-v1:00670065006e006500720061007400650064002f004d00610069006e:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f004d00610069006e005f004d0075006c007400690070006c0065004b0069006e00640050006100720061006d0073
func init() {
	gopurs_runtime.StructGetters[3989967386] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_MultipleKindParams[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {

		default:
			panic("Key not found in dictionary Constructor_Main_MultipleKindParams: " + key)
		}
	}
}

type Constructor_Main_MultipleKindParamsChain[T_customKind any] struct {
	Rc uint32
}

//gopurs:init-v1:00670065006e006500720061007400650064002f004d00610069006e:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f004d00610069006e005f004d0075006c007400690070006c0065004b0069006e00640050006100720061006d00730043006800610069006e
func init() {
	gopurs_runtime.StructGetters[467154327] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_MultipleKindParamsChain[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {

		default:
			panic("Key not found in dictionary Constructor_Main_MultipleKindParamsChain: " + key)
		}
	}
}

func Call_Main_hktpChain(dict_0_loop *Constructor_Main_HigherKindedTypeParamsChain[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_HigherKindedTypeParamsChain[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V0
}

func Call_Main_hktp(dict_0_loop *Constructor_Main_HigherKindedTypeParams[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_HigherKindedTypeParams[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V0
}

func Rebox_Main_855013872_2818661616(in *Constructor_Data_Generic_Rep_Generic[uint32, uint32]) *Constructor_Data_Generic_Rep_Generic[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Generic_Rep_Generic[gopurs_runtime.Value, gopurs_runtime.Value])(unsafe.Pointer(in))
}
