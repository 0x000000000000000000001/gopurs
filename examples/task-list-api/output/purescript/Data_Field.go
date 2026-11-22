package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Field_Field_dollar_Dict gopurs_runtime.Value
var once_Data_Field_Field_dollar_Dict sync.Once

func Get_Data_Field_Field_dollar_Dict() gopurs_runtime.Value {
	once_Data_Field_Field_dollar_Dict.Do(func() {
		cache_Data_Field_Field_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3997783546, UnsafePtr: unsafe.Pointer(Call_Data_Field_Field_dollar_Dict(func() struct {
				DivisionRing1  gopurs_runtime.Value
				EuclideanRing0 gopurs_runtime.Value
				/*gopurs:record-v1:2:004400690076006900730069006f006e00520069006e00670031,004500750063006c0069006400650061006e00520069006e00670030*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					DivisionRing1  gopurs_runtime.Value
					EuclideanRing0 gopurs_runtime.Value
					/*gopurs:record-v1:2:004400690076006900730069006f006e00520069006e00670031,004500750063006c0069006400650061006e00520069006e00670030*/
				}{}
				clone.DivisionRing1 = gopurs_runtime.RecordGet(orig, "DivisionRing1")
				clone.EuclideanRing0 = gopurs_runtime.RecordGet(orig, "EuclideanRing0")
				return clone
			}()))}
		})
	})
	return cache_Data_Field_Field_dollar_Dict
}

var cache_Data_Field_field gopurs_runtime.Value
var once_Data_Field_field sync.Once

func Get_Data_Field_field() gopurs_runtime.Value {
	once_Data_Field_field.Do(func() {
		cache_Data_Field_field = gopurs_runtime.Func2(func(dictEuclideanRing_0_box gopurs_runtime.Value, dictDivisionRing_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Field_field(dictEuclideanRing_0_box, dictDivisionRing_1_box)
		})
	})
	return cache_Data_Field_field
}

type Constructor_Data_Field_Field[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}

//gopurs:init-v1:00670065006e006500720061007400650064002f0044006100740061005f004600690065006c0064:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f0044006100740061005f004600690065006c0064005f004600690065006c0064
func init() {
	gopurs_runtime.StructGetters[3997783546] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Field_Field[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "DivisionRing1":
			return gopurs_runtime.Box(c.V0)
		case "EuclideanRing0":
			return gopurs_runtime.Box(c.V1)
		default:
			panic("Key not found in dictionary Constructor_Data_Field_Field: " + key)
		}
	}
}

func Call_Data_Field_Field_dollar_Dict(x_0_loop struct {
	DivisionRing1  gopurs_runtime.Value
	EuclideanRing0 gopurs_runtime.Value
	/*gopurs:record-v1:2:004400690076006900730069006f006e00520069006e00670031,004500750063006c0069006400650061006e00520069006e00670030*/
}) *Constructor_Data_Field_Field[gopurs_runtime.Value] {
	var x_0 struct {
		DivisionRing1  gopurs_runtime.Value
		EuclideanRing0 gopurs_runtime.Value
		/*gopurs:record-v1:2:004400690076006900730069006f006e00520069006e00670031,004500750063006c0069006400650061006e00520069006e00670030*/
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Data_Field_Field[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict2("DivisionRing1", "EuclideanRing0", orig.DivisionRing1, orig.EuclideanRing0)
	}())
}

func Call_Data_Field_field(dictEuclideanRing_0_loop gopurs_runtime.Value, dictDivisionRing_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictEuclideanRing_0 gopurs_runtime.Value = dictEuclideanRing_0_loop
	_ = dictEuclideanRing_0
	var dictDivisionRing_1 gopurs_runtime.Value = dictDivisionRing_1_loop
	_ = dictDivisionRing_1
	return gopurs_runtime.Value{Type: 9, IntVal: 3997783546, UnsafePtr: unsafe.Pointer((&Constructor_Data_Field_Field[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return dictDivisionRing_1
	}), gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return dictEuclideanRing_0
	})}))}
}
