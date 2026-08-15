package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_ReservedWord_dollar_Dict gopurs_runtime.Value
var once_Main_ReservedWord_dollar_Dict sync.Once

func Get_Main_ReservedWord_dollar_Dict() gopurs_runtime.Value {
	once_Main_ReservedWord_dollar_Dict.Do(func() {
		cache_Main_ReservedWord_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_ReservedWord_dollar_Dict(x_0_box)
		})
	})
	return cache_Main_ReservedWord_dollar_Dict
}

var cache_Main_OverlappingStillCompiles_dollar_Dict gopurs_runtime.Value
var once_Main_OverlappingStillCompiles_dollar_Dict sync.Once

func Get_Main_OverlappingStillCompiles_dollar_Dict() gopurs_runtime.Value {
	once_Main_OverlappingStillCompiles_dollar_Dict.Do(func() {
		cache_Main_OverlappingStillCompiles_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_OverlappingStillCompiles_dollar_Dict(x_0_box)
		})
	})
	return cache_Main_OverlappingStillCompiles_dollar_Dict
}

var cache_Main_OneTypeParamChain_dollar_Dict gopurs_runtime.Value
var once_Main_OneTypeParamChain_dollar_Dict sync.Once

func Get_Main_OneTypeParamChain_dollar_Dict() gopurs_runtime.Value {
	once_Main_OneTypeParamChain_dollar_Dict.Do(func() {
		cache_Main_OneTypeParamChain_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_OneTypeParamChain_dollar_Dict(x_0_box)
		})
	})
	return cache_Main_OneTypeParamChain_dollar_Dict
}

var cache_Main_OneTypeParam_dollar_Dict gopurs_runtime.Value
var once_Main_OneTypeParam_dollar_Dict sync.Once

func Get_Main_OneTypeParam_dollar_Dict() gopurs_runtime.Value {
	once_Main_OneTypeParam_dollar_Dict.Do(func() {
		cache_Main_OneTypeParam_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_OneTypeParam_dollar_Dict(x_0_box)
		})
	})
	return cache_Main_OneTypeParam_dollar_Dict
}

var cache_Main_NoTypeParams_dollar_Dict gopurs_runtime.Value
var once_Main_NoTypeParams_dollar_Dict sync.Once

func Get_Main_NoTypeParams_dollar_Dict() gopurs_runtime.Value {
	once_Main_NoTypeParams_dollar_Dict.Do(func() {
		cache_Main_NoTypeParams_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_NoTypeParams_dollar_Dict(x_0_box)
		})
	})
	return cache_Main_NoTypeParams_dollar_Dict
}

var cache_Main_MultipleTypeParamsChain_dollar_Dict gopurs_runtime.Value
var once_Main_MultipleTypeParamsChain_dollar_Dict sync.Once

func Get_Main_MultipleTypeParamsChain_dollar_Dict() gopurs_runtime.Value {
	once_Main_MultipleTypeParamsChain_dollar_Dict.Do(func() {
		cache_Main_MultipleTypeParamsChain_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_MultipleTypeParamsChain_dollar_Dict(x_0_box)
		})
	})
	return cache_Main_MultipleTypeParamsChain_dollar_Dict
}

var cache_Main_MultipleTypeParams_dollar_Dict gopurs_runtime.Value
var once_Main_MultipleTypeParams_dollar_Dict sync.Once

func Get_Main_MultipleTypeParams_dollar_Dict() gopurs_runtime.Value {
	once_Main_MultipleTypeParams_dollar_Dict.Do(func() {
		cache_Main_MultipleTypeParams_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_MultipleTypeParams_dollar_Dict(x_0_box)
		})
	})
	return cache_Main_MultipleTypeParams_dollar_Dict
}

var cache_Main_MultipleKindParamsChain_dollar_Dict gopurs_runtime.Value
var once_Main_MultipleKindParamsChain_dollar_Dict sync.Once

func Get_Main_MultipleKindParamsChain_dollar_Dict() gopurs_runtime.Value {
	once_Main_MultipleKindParamsChain_dollar_Dict.Do(func() {
		cache_Main_MultipleKindParamsChain_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_MultipleKindParamsChain_dollar_Dict(x_0_box)
		})
	})
	return cache_Main_MultipleKindParamsChain_dollar_Dict
}

var cache_Main_MultipleKindParams_dollar_Dict gopurs_runtime.Value
var once_Main_MultipleKindParams_dollar_Dict sync.Once

func Get_Main_MultipleKindParams_dollar_Dict() gopurs_runtime.Value {
	once_Main_MultipleKindParams_dollar_Dict.Do(func() {
		cache_Main_MultipleKindParams_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_MultipleKindParams_dollar_Dict(x_0_box)
		})
	})
	return cache_Main_MultipleKindParams_dollar_Dict
}

var cache_Main_HigherKindedTypeParamsChain_dollar_Dict gopurs_runtime.Value
var once_Main_HigherKindedTypeParamsChain_dollar_Dict sync.Once

func Get_Main_HigherKindedTypeParamsChain_dollar_Dict() gopurs_runtime.Value {
	once_Main_HigherKindedTypeParamsChain_dollar_Dict.Do(func() {
		cache_Main_HigherKindedTypeParamsChain_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_HigherKindedTypeParamsChain_dollar_Dict(x_0_box)
		})
	})
	return cache_Main_HigherKindedTypeParamsChain_dollar_Dict
}

var cache_Main_HigherKindedTypeParams_dollar_Dict gopurs_runtime.Value
var once_Main_HigherKindedTypeParams_dollar_Dict sync.Once

func Get_Main_HigherKindedTypeParams_dollar_Dict() gopurs_runtime.Value {
	once_Main_HigherKindedTypeParams_dollar_Dict.Do(func() {
		cache_Main_HigherKindedTypeParams_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_HigherKindedTypeParams_dollar_Dict(x_0_box)
		})
	})
	return cache_Main_HigherKindedTypeParams_dollar_Dict
}

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
			return gopurs_runtime.Value{Type: 9, IntVal: 913637797, UnsafePtr: unsafe.Pointer((&Constructor_Main_Left{1, value0}))}
		})
	})
	return cache_Main_Left
}

var cache_Main_Right gopurs_runtime.Value
var once_Main_Right sync.Once

func Get_Main_Right() gopurs_runtime.Value {
	once_Main_Right.Do(func() {
		cache_Main_Right = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2535318782, UnsafePtr: unsafe.Pointer((&Constructor_Main_Right{1, value0}))}
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
				return gopurs_runtime.Value{Type: 9, IntVal: 2763139640, UnsafePtr: unsafe.Pointer((&Constructor_Main_Foo{1, value0, value1}))}
			})
		})
	})
	return cache_Main_Foo
}

var cache_Main_reservedWordFunction gopurs_runtime.Value
var once_Main_reservedWordFunction sync.Once

func Get_Main_reservedWordFunction() gopurs_runtime.Value {
	once_Main_reservedWordFunction.Do(func() {
		cache_Main_reservedWordFunction = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.RecordDict0().IntVal)), UnsafePtr: nil}
	})
	return cache_Main_reservedWordFunction
}

var cache_Main_reservedWordArrow gopurs_runtime.Value
var once_Main_reservedWordArrow sync.Once

func Get_Main_reservedWordArrow() gopurs_runtime.Value {
	once_Main_reservedWordArrow.Do(func() {
		cache_Main_reservedWordArrow = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.RecordDict0().IntVal)), UnsafePtr: nil}
	})
	return cache_Main_reservedWordArrow
}

var cache_Main_overlappingStillCompiles gopurs_runtime.Value
var once_Main_overlappingStillCompiles sync.Once

func Get_Main_overlappingStillCompiles() gopurs_runtime.Value {
	once_Main_overlappingStillCompiles.Do(func() {
		cache_Main_overlappingStillCompiles = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.RecordDict0().IntVal)), UnsafePtr: nil}
	})
	return cache_Main_overlappingStillCompiles
}

var cache_Main_overlappingStillCompiles1 gopurs_runtime.Value
var once_Main_overlappingStillCompiles1 sync.Once

func Get_Main_overlappingStillCompiles1() gopurs_runtime.Value {
	once_Main_overlappingStillCompiles1.Do(func() {
		cache_Main_overlappingStillCompiles1 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.RecordDict0().IntVal)), UnsafePtr: nil}
	})
	return cache_Main_overlappingStillCompiles1
}

var cache_Main_oneTypeParamChainString gopurs_runtime.Value
var once_Main_oneTypeParamChainString sync.Once

func Get_Main_oneTypeParamChainString() gopurs_runtime.Value {
	once_Main_oneTypeParamChainString.Do(func() {
		cache_Main_oneTypeParamChainString = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.RecordDict0().IntVal)), UnsafePtr: nil}
	})
	return cache_Main_oneTypeParamChainString
}

var cache_Main_oneTypeParamChainBoolean gopurs_runtime.Value
var once_Main_oneTypeParamChainBoolean sync.Once

func Get_Main_oneTypeParamChainBoolean() gopurs_runtime.Value {
	once_Main_oneTypeParamChainBoolean.Do(func() {
		cache_Main_oneTypeParamChainBoolean = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.RecordDict0().IntVal)), UnsafePtr: nil}
	})
	return cache_Main_oneTypeParamChainBoolean
}

var cache_Main_oneTypeParamBoolean gopurs_runtime.Value
var once_Main_oneTypeParamBoolean sync.Once

func Get_Main_oneTypeParamBoolean() gopurs_runtime.Value {
	once_Main_oneTypeParamBoolean.Do(func() {
		cache_Main_oneTypeParamBoolean = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.RecordDict0().IntVal)), UnsafePtr: nil}
	})
	return cache_Main_oneTypeParamBoolean
}

var cache_Main_noTypeParams gopurs_runtime.Value
var once_Main_noTypeParams sync.Once

func Get_Main_noTypeParams() gopurs_runtime.Value {
	once_Main_noTypeParams.Do(func() {
		cache_Main_noTypeParams = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.RecordDict0().IntVal)), UnsafePtr: nil}
	})
	return cache_Main_noTypeParams
}

var cache_Main_multipleTypeParamsChainBo gopurs_runtime.Value
var once_Main_multipleTypeParamsChainBo sync.Once

func Get_Main_multipleTypeParamsChainBo() gopurs_runtime.Value {
	once_Main_multipleTypeParamsChainBo.Do(func() {
		cache_Main_multipleTypeParamsChainBo = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.RecordDict0().IntVal)), UnsafePtr: nil}
	})
	return cache_Main_multipleTypeParamsChainBo
}

var cache_Main_multipleTypeParamsChainBo1 gopurs_runtime.Value
var once_Main_multipleTypeParamsChainBo1 sync.Once

func Get_Main_multipleTypeParamsChainBo1() gopurs_runtime.Value {
	once_Main_multipleTypeParamsChainBo1.Do(func() {
		cache_Main_multipleTypeParamsChainBo1 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.RecordDict0().IntVal)), UnsafePtr: nil}
	})
	return cache_Main_multipleTypeParamsChainBo1
}

var cache_Main_multipleTypeParamsChainBo2 gopurs_runtime.Value
var once_Main_multipleTypeParamsChainBo2 sync.Once

func Get_Main_multipleTypeParamsChainBo2() gopurs_runtime.Value {
	once_Main_multipleTypeParamsChainBo2.Do(func() {
		cache_Main_multipleTypeParamsChainBo2 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.RecordDict0().IntVal)), UnsafePtr: nil}
	})
	return cache_Main_multipleTypeParamsChainBo2
}

var cache_Main_multipleTypeParamsChainBo3 gopurs_runtime.Value
var once_Main_multipleTypeParamsChainBo3 sync.Once

func Get_Main_multipleTypeParamsChainBo3() gopurs_runtime.Value {
	once_Main_multipleTypeParamsChainBo3.Do(func() {
		cache_Main_multipleTypeParamsChainBo3 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.RecordDict0().IntVal)), UnsafePtr: nil}
	})
	return cache_Main_multipleTypeParamsChainBo3
}

var cache_Main_multipleTypeParamsChainBo4 gopurs_runtime.Value
var once_Main_multipleTypeParamsChainBo4 sync.Once

func Get_Main_multipleTypeParamsChainBo4() gopurs_runtime.Value {
	once_Main_multipleTypeParamsChainBo4.Do(func() {
		cache_Main_multipleTypeParamsChainBo4 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.RecordDict0().IntVal)), UnsafePtr: nil}
	})
	return cache_Main_multipleTypeParamsChainBo4
}

var cache_Main_multipleTypeParamsBoolean gopurs_runtime.Value
var once_Main_multipleTypeParamsBoolean sync.Once

func Get_Main_multipleTypeParamsBoolean() gopurs_runtime.Value {
	once_Main_multipleTypeParamsBoolean.Do(func() {
		cache_Main_multipleTypeParamsBoolean = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.RecordDict0().IntVal)), UnsafePtr: nil}
	})
	return cache_Main_multipleTypeParamsBoolean
}

var cache_Main_multipleKindParamsConstru gopurs_runtime.Value
var once_Main_multipleKindParamsConstru sync.Once

func Get_Main_multipleKindParamsConstru() gopurs_runtime.Value {
	once_Main_multipleKindParamsConstru.Do(func() {
		cache_Main_multipleKindParamsConstru = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.RecordDict0().IntVal)), UnsafePtr: nil}
	})
	return cache_Main_multipleKindParamsConstru
}

var cache_Main_multipleKindParamsChainCo gopurs_runtime.Value
var once_Main_multipleKindParamsChainCo sync.Once

func Get_Main_multipleKindParamsChainCo() gopurs_runtime.Value {
	once_Main_multipleKindParamsChainCo.Do(func() {
		cache_Main_multipleKindParamsChainCo = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.RecordDict0().IntVal)), UnsafePtr: nil}
	})
	return cache_Main_multipleKindParamsChainCo
}

var cache_Main_multipleKindParamsChainCo1 gopurs_runtime.Value
var once_Main_multipleKindParamsChainCo1 sync.Once

func Get_Main_multipleKindParamsChainCo1() gopurs_runtime.Value {
	once_Main_multipleKindParamsChainCo1.Do(func() {
		cache_Main_multipleKindParamsChainCo1 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.RecordDict0().IntVal)), UnsafePtr: nil}
	})
	return cache_Main_multipleKindParamsChainCo1
}

var cache_Main_multipleKindParamsChainCo2 gopurs_runtime.Value
var once_Main_multipleKindParamsChainCo2 sync.Once

func Get_Main_multipleKindParamsChainCo2() gopurs_runtime.Value {
	once_Main_multipleKindParamsChainCo2.Do(func() {
		cache_Main_multipleKindParamsChainCo2 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.RecordDict0().IntVal)), UnsafePtr: nil}
	})
	return cache_Main_multipleKindParamsChainCo2
}

var cache_Main_higherKindedTypeParamsCha gopurs_runtime.Value
var once_Main_higherKindedTypeParamsCha sync.Once

func Get_Main_higherKindedTypeParamsCha() gopurs_runtime.Value {
	once_Main_higherKindedTypeParamsCha.Do(func() {
		cache_Main_higherKindedTypeParamsCha = gopurs_runtime.Value{Type: 9, IntVal: 1868017351, UnsafePtr: unsafe.Pointer((&Constructor_Main_HigherKindedTypeParamsChain{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Int(0)
			})
		})}))}
	})
	return cache_Main_higherKindedTypeParamsCha
}

var cache_Main_higherKindedTypeParamsCha1 gopurs_runtime.Value
var once_Main_higherKindedTypeParamsCha1 sync.Once

func Get_Main_higherKindedTypeParamsCha1() gopurs_runtime.Value {
	once_Main_higherKindedTypeParamsCha1.Do(func() {
		cache_Main_higherKindedTypeParamsCha1 = gopurs_runtime.Value{Type: 9, IntVal: 1868017351, UnsafePtr: unsafe.Pointer((&Constructor_Main_HigherKindedTypeParamsChain{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Int(0)
			})
		})}))}
	})
	return cache_Main_higherKindedTypeParamsCha1
}

var cache_Main_higherKindedTypeParamsArr gopurs_runtime.Value
var once_Main_higherKindedTypeParamsArr sync.Once

func Get_Main_higherKindedTypeParamsArr() gopurs_runtime.Value {
	once_Main_higherKindedTypeParamsArr.Do(func() {
		cache_Main_higherKindedTypeParamsArr = gopurs_runtime.Value{Type: 9, IntVal: 3032179466, UnsafePtr: unsafe.Pointer((&Constructor_Main_HigherKindedTypeParams{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Int(0)
			})
		})}))}
	})
	return cache_Main_higherKindedTypeParamsArr
}

var cache_Main_genericGenericFoo_ gopurs_runtime.Value
var once_Main_genericGenericFoo_ sync.Once

func Get_Main_genericGenericFoo_() gopurs_runtime.Value {
	once_Main_genericGenericFoo_.Do(func() {
		cache_Main_genericGenericFoo_ = gopurs_runtime.Value{Type: 9, IntVal: 1921946594, UnsafePtr: unsafe.Pointer((&Constructor_Data_Generic_Rep_Generic{1, gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(1454898258), UnsafePtr: nil}
		}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(682342953), UnsafePtr: nil}
		})}))}
	})
	return cache_Main_genericGenericFoo_
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
		})
	})
	return cache_Main_main
}

var cache_Main_hktpChain gopurs_runtime.Value
var once_Main_hktpChain sync.Once

func Get_Main_hktpChain() gopurs_runtime.Value {
	once_Main_hktpChain.Do(func() {
		cache_Main_hktpChain = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_hktpChain(gopurs_runtime.CoerceToStruct[Constructor_Main_HigherKindedTypeParamsChain](dict_0_box))
		})
	})
	return cache_Main_hktpChain
}

var cache_Main_hktp gopurs_runtime.Value
var once_Main_hktp sync.Once

func Get_Main_hktp() gopurs_runtime.Value {
	once_Main_hktp.Do(func() {
		cache_Main_hktp = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_hktp(gopurs_runtime.CoerceToStruct[Constructor_Main_HigherKindedTypeParams](dict_0_box))
		})
	})
	return cache_Main_hktp
}

type Constructor_Main_GenericFoo struct {
	Rc uint32
}

type Constructor_Main_Left struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

type Constructor_Main_Right struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

type Constructor_Main_Foo struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}

type Constructor_Main_ReservedWord struct {
	Rc uint32
}

func init() {
	gopurs_runtime.StructGetters[3965187028] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_ReservedWord)(ptr)
		_ = c
		switch key {

		default:
			panic("Key not found in dictionary Constructor_Main_ReservedWord: " + key)
		}
	}
}

type Constructor_Main_OverlappingStillCompiles struct {
	Rc uint32
}

func init() {
	gopurs_runtime.StructGetters[1985592625] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_OverlappingStillCompiles)(ptr)
		_ = c
		switch key {

		default:
			panic("Key not found in dictionary Constructor_Main_OverlappingStillCompiles: " + key)
		}
	}
}

type Constructor_Main_OneTypeParamChain struct {
	Rc uint32
}

func init() {
	gopurs_runtime.StructGetters[4048695264] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_OneTypeParamChain)(ptr)
		_ = c
		switch key {

		default:
			panic("Key not found in dictionary Constructor_Main_OneTypeParamChain: " + key)
		}
	}
}

type Constructor_Main_OneTypeParam struct {
	Rc uint32
}

func init() {
	gopurs_runtime.StructGetters[1734428973] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_OneTypeParam)(ptr)
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

type Constructor_Main_MultipleTypeParamsChain struct {
	Rc uint32
}

func init() {
	gopurs_runtime.StructGetters[3346744871] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_MultipleTypeParamsChain)(ptr)
		_ = c
		switch key {

		default:
			panic("Key not found in dictionary Constructor_Main_MultipleTypeParamsChain: " + key)
		}
	}
}

type Constructor_Main_MultipleTypeParams struct {
	Rc uint32
}

func init() {
	gopurs_runtime.StructGetters[1755820650] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_MultipleTypeParams)(ptr)
		_ = c
		switch key {

		default:
			panic("Key not found in dictionary Constructor_Main_MultipleTypeParams: " + key)
		}
	}
}

type Constructor_Main_HigherKindedTypeParamsChain struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

func init() {
	gopurs_runtime.StructGetters[1868017351] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_HigherKindedTypeParamsChain)(ptr)
		_ = c
		switch key {
		case "hktpChain":
			return gopurs_runtime.Box(c.V0)
		default:
			panic("Key not found in dictionary Constructor_Main_HigherKindedTypeParamsChain: " + key)
		}
	}
}

type Constructor_Main_HigherKindedTypeParams struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

func init() {
	gopurs_runtime.StructGetters[3032179466] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_HigherKindedTypeParams)(ptr)
		_ = c
		switch key {
		case "hktp":
			return gopurs_runtime.Box(c.V0)
		default:
			panic("Key not found in dictionary Constructor_Main_HigherKindedTypeParams: " + key)
		}
	}
}

type Constructor_Main_MultipleKindParams struct {
	Rc uint32
}

func init() {
	gopurs_runtime.StructGetters[3989967386] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_MultipleKindParams)(ptr)
		_ = c
		switch key {

		default:
			panic("Key not found in dictionary Constructor_Main_MultipleKindParams: " + key)
		}
	}
}

type Constructor_Main_MultipleKindParamsChain struct {
	Rc uint32
}

func init() {
	gopurs_runtime.StructGetters[467154327] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_MultipleKindParamsChain)(ptr)
		_ = c
		switch key {

		default:
			panic("Key not found in dictionary Constructor_Main_MultipleKindParamsChain: " + key)
		}
	}
}

func Call_Main_ReservedWord_dollar_Dict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_OverlappingStillCompiles_dollar_Dict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_OneTypeParamChain_dollar_Dict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_OneTypeParam_dollar_Dict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_NoTypeParams_dollar_Dict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_MultipleTypeParamsChain_dollar_Dict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_MultipleTypeParams_dollar_Dict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_MultipleKindParamsChain_dollar_Dict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_MultipleKindParams_dollar_Dict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_HigherKindedTypeParamsChain_dollar_Dict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_HigherKindedTypeParams_dollar_Dict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_hktpChain(dict_0_loop *Constructor_Main_HigherKindedTypeParamsChain) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_HigherKindedTypeParamsChain = dict_0_loop
	_ = dict_0
	return gopurs_runtime.Box(dict_0.V0)
}

func Call_Main_hktp(dict_0_loop *Constructor_Main_HigherKindedTypeParams) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_HigherKindedTypeParams = dict_0_loop
	_ = dict_0
	return gopurs_runtime.Box(dict_0.V0)
}
