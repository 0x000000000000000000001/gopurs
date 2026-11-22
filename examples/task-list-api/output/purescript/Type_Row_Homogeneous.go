package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Type_Row_Homogeneous_HomogeneousRowList_dollar_Dict gopurs_runtime.Value
var once_Type_Row_Homogeneous_HomogeneousRowList_dollar_Dict sync.Once

func Get_Type_Row_Homogeneous_HomogeneousRowList_dollar_Dict() gopurs_runtime.Value {
	once_Type_Row_Homogeneous_HomogeneousRowList_dollar_Dict.Do(func() {
		cache_Type_Row_Homogeneous_HomogeneousRowList_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Type_Row_Homogeneous_HomogeneousRowList_dollar_Dict(func() struct {

				/*gopurs:record-v1:0:*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {

					/*gopurs:record-v1:0:*/
				}{}

				return clone
			}())
		})
	})
	return cache_Type_Row_Homogeneous_HomogeneousRowList_dollar_Dict
}

var cache_Type_Row_Homogeneous_HomogeneousRowList_dollar_Dict1 gopurs_runtime.Value
var once_Type_Row_Homogeneous_HomogeneousRowList_dollar_Dict1 sync.Once

func Get_Type_Row_Homogeneous_HomogeneousRowList_dollar_Dict1() gopurs_runtime.Value {
	once_Type_Row_Homogeneous_HomogeneousRowList_dollar_Dict1.Do(func() {
		cache_Type_Row_Homogeneous_HomogeneousRowList_dollar_Dict1 = func() gopurs_runtime.Value {
			orig := struct {

				/*gopurs:record-v1:0:*/
			}{}
			_ = orig
			return gopurs_runtime.RecordDict0()
		}()
	})
	return cache_Type_Row_Homogeneous_HomogeneousRowList_dollar_Dict1
}

var cache_Type_Row_Homogeneous_Homogeneous_dollar_Dict gopurs_runtime.Value
var once_Type_Row_Homogeneous_Homogeneous_dollar_Dict sync.Once

func Get_Type_Row_Homogeneous_Homogeneous_dollar_Dict() gopurs_runtime.Value {
	once_Type_Row_Homogeneous_Homogeneous_dollar_Dict.Do(func() {
		cache_Type_Row_Homogeneous_Homogeneous_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Type_Row_Homogeneous_Homogeneous_dollar_Dict(func() struct {

				/*gopurs:record-v1:0:*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {

					/*gopurs:record-v1:0:*/
				}{}

				return clone
			}())
		})
	})
	return cache_Type_Row_Homogeneous_Homogeneous_dollar_Dict
}

var cache_Type_Row_Homogeneous_Homogeneous_dollar_Dict1 gopurs_runtime.Value
var once_Type_Row_Homogeneous_Homogeneous_dollar_Dict1 sync.Once

func Get_Type_Row_Homogeneous_Homogeneous_dollar_Dict1() gopurs_runtime.Value {
	once_Type_Row_Homogeneous_Homogeneous_dollar_Dict1.Do(func() {
		cache_Type_Row_Homogeneous_Homogeneous_dollar_Dict1 = func() gopurs_runtime.Value {
			orig := struct {

				/*gopurs:record-v1:0:*/
			}{}
			_ = orig
			return gopurs_runtime.RecordDict0()
		}()
	})
	return cache_Type_Row_Homogeneous_Homogeneous_dollar_Dict1
}

var cache_Type_Row_Homogeneous_homogeneousRowListNil gopurs_runtime.Value
var once_Type_Row_Homogeneous_homogeneousRowListNil sync.Once

func Get_Type_Row_Homogeneous_homogeneousRowListNil() gopurs_runtime.Value {
	once_Type_Row_Homogeneous_homogeneousRowListNil.Do(func() {
		cache_Type_Row_Homogeneous_homogeneousRowListNil = func() gopurs_runtime.Value {
			orig := struct {

				/*gopurs:record-v1:0:*/
			}{}
			_ = orig
			return gopurs_runtime.RecordDict0()
		}()
	})
	return cache_Type_Row_Homogeneous_homogeneousRowListNil
}

var cache_Type_Row_Homogeneous_homogeneousRowListCons gopurs_runtime.Value
var once_Type_Row_Homogeneous_homogeneousRowListCons sync.Once

func Get_Type_Row_Homogeneous_homogeneousRowListCons() gopurs_runtime.Value {
	once_Type_Row_Homogeneous_homogeneousRowListCons.Do(func() {
		cache_Type_Row_Homogeneous_homogeneousRowListCons = gopurs_runtime.Func2(func(_dollar___unused_0_box gopurs_runtime.Value, dictTypeEquals_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Type_Row_Homogeneous_homogeneousRowListCons(_dollar___unused_0_box, dictTypeEquals_1_box)
		})
	})
	return cache_Type_Row_Homogeneous_homogeneousRowListCons
}

var cache_Type_Row_Homogeneous_homogeneous gopurs_runtime.Value
var once_Type_Row_Homogeneous_homogeneous sync.Once

func Get_Type_Row_Homogeneous_homogeneous() gopurs_runtime.Value {
	once_Type_Row_Homogeneous_homogeneous.Do(func() {
		cache_Type_Row_Homogeneous_homogeneous = gopurs_runtime.Func2(func(_dollar___unused_0_box gopurs_runtime.Value, _dollar___unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Type_Row_Homogeneous_homogeneous(_dollar___unused_0_box, _dollar___unused_1_box)
		})
	})
	return cache_Type_Row_Homogeneous_homogeneous
}

type Constructor_Type_Row_Homogeneous_HomogeneousRowList[T_rowList any, T_fieldType any] struct {
	Rc uint32
}

//gopurs:init-v1:00670065006e006500720061007400650064002f0054007900700065005f0052006f0077005f0048006f006d006f00670065006e0065006f00750073:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f0054007900700065005f0052006f0077005f0048006f006d006f00670065006e0065006f00750073005f0048006f006d006f00670065006e0065006f007500730052006f0077004c006900730074
func init() {
	gopurs_runtime.StructGetters[2816475247] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Type_Row_Homogeneous_HomogeneousRowList[gopurs_runtime.Value, gopurs_runtime.Value])(ptr)
		_ = c
		switch key {

		default:
			panic("Key not found in dictionary Constructor_Type_Row_Homogeneous_HomogeneousRowList: " + key)
		}
	}
}

type Constructor_Type_Row_Homogeneous_Homogeneous[T_row any, T_fieldType any] struct {
	Rc uint32
}

//gopurs:init-v1:00670065006e006500720061007400650064002f0054007900700065005f0052006f0077005f0048006f006d006f00670065006e0065006f00750073:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f0054007900700065005f0052006f0077005f0048006f006d006f00670065006e0065006f00750073005f0048006f006d006f00670065006e0065006f00750073
func init() {
	gopurs_runtime.StructGetters[4189377959] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Type_Row_Homogeneous_Homogeneous[gopurs_runtime.Value, gopurs_runtime.Value])(ptr)
		_ = c
		switch key {

		default:
			panic("Key not found in dictionary Constructor_Type_Row_Homogeneous_Homogeneous: " + key)
		}
	}
}

func Call_Type_Row_Homogeneous_HomogeneousRowList_dollar_Dict(x_0_loop struct {

	/*gopurs:record-v1:0:*/
}) gopurs_runtime.Value {
	var x_0 struct {

		/*gopurs:record-v1:0:*/
	} = x_0_loop
	_ = x_0
	return func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict0()
	}()
}

func Call_Type_Row_Homogeneous_Homogeneous_dollar_Dict(x_0_loop struct {

	/*gopurs:record-v1:0:*/
}) gopurs_runtime.Value {
	var x_0 struct {

		/*gopurs:record-v1:0:*/
	} = x_0_loop
	_ = x_0
	return func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict0()
	}()
}

func Call_Type_Row_Homogeneous_homogeneousRowListCons(_dollar___unused_0_loop gopurs_runtime.Value, dictTypeEquals_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var dictTypeEquals_1 gopurs_runtime.Value = dictTypeEquals_1_loop
	_ = dictTypeEquals_1
	return Get_Type_Row_Homogeneous_HomogeneousRowList_dollar_Dict1()
}

func Call_Type_Row_Homogeneous_homogeneous(_dollar___unused_0_loop gopurs_runtime.Value, _dollar___unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var _dollar___unused_1 gopurs_runtime.Value = _dollar___unused_1_loop
	_ = _dollar___unused_1
	return Get_Type_Row_Homogeneous_Homogeneous_dollar_Dict1()
}
