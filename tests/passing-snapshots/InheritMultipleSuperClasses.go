package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_identity gopurs_runtime.Value
var once_Main_identity sync.Once

func Get_Main_identity() gopurs_runtime.Value {
	once_Main_identity.Do(func() {
		cache_Main_identity = Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))})
	})
	return cache_Main_identity
}

var cache_Main_identity1 gopurs_runtime.Value
var once_Main_identity1 sync.Once

func Get_Main_identity1() gopurs_runtime.Value {
	once_Main_identity1.Do(func() {
		cache_Main_identity1 = Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))})
	})
	return cache_Main_identity1
}

var cache_Main_identity2 gopurs_runtime.Value
var once_Main_identity2 sync.Once

func Get_Main_identity2() gopurs_runtime.Value {
	once_Main_identity2.Do(func() {
		cache_Main_identity2 = Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))})
	})
	return cache_Main_identity2
}

var cache_Main_identity3 gopurs_runtime.Value
var once_Main_identity3 sync.Once

func Get_Main_identity3() gopurs_runtime.Value {
	once_Main_identity3.Do(func() {
		cache_Main_identity3 = Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))})
	})
	return cache_Main_identity3
}

var cache_Main_Eg2_dollar_Dict gopurs_runtime.Value
var once_Main_Eg2_dollar_Dict sync.Once

func Get_Main_Eg2_dollar_Dict() gopurs_runtime.Value {
	once_Main_Eg2_dollar_Dict.Do(func() {
		cache_Main_Eg2_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2207147950, UnsafePtr: unsafe.Pointer(Call_Main_Eg2_dollar_Dict(func() struct {
				Functor0 gopurs_runtime.Value
				Functor1 gopurs_runtime.Value
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Functor0 gopurs_runtime.Value
					Functor1 gopurs_runtime.Value
				}{}
				clone.Functor0 = gopurs_runtime.RecordGet(orig, "Functor0")
				clone.Functor1 = gopurs_runtime.RecordGet(orig, "Functor1")
				return clone
			}()))}
		})
	})
	return cache_Main_Eg2_dollar_Dict
}

var cache_Main_Eg1_dollar_Dict gopurs_runtime.Value
var once_Main_Eg1_dollar_Dict sync.Once

func Get_Main_Eg1_dollar_Dict() gopurs_runtime.Value {
	once_Main_Eg1_dollar_Dict.Do(func() {
		cache_Main_Eg1_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1828296749, UnsafePtr: unsafe.Pointer(Call_Main_Eg1_dollar_Dict(func() struct {
				Functor0 gopurs_runtime.Value
				Functor1 gopurs_runtime.Value
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Functor0 gopurs_runtime.Value
					Functor1 gopurs_runtime.Value
				}{}
				clone.Functor0 = gopurs_runtime.RecordGet(orig, "Functor0")
				clone.Functor1 = gopurs_runtime.RecordGet(orig, "Functor1")
				return clone
			}()))}
		})
	})
	return cache_Main_Eg1_dollar_Dict
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

var cache_Main_g2 gopurs_runtime.Value
var once_Main_g2 sync.Once

func Get_Main_g2() gopurs_runtime.Value {
	once_Main_g2.Do(func() {
		cache_Main_g2 = gopurs_runtime.Func(func(dictEg2_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_g2(gopurs_runtime.CoerceToStruct[Constructor_Main_Eg2[gopurs_runtime.Value, gopurs_runtime.Value]](dictEg2_0_box))
		})
	})
	return cache_Main_g2
}

var cache_Main_g1 gopurs_runtime.Value
var once_Main_g1 sync.Once

func Get_Main_g1() gopurs_runtime.Value {
	once_Main_g1.Do(func() {
		cache_Main_g1 = gopurs_runtime.Func(func(dictEg1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_g1(gopurs_runtime.CoerceToStruct[Constructor_Main_Eg1[gopurs_runtime.Value, gopurs_runtime.Value]](dictEg1_0_box))
		})
	})
	return cache_Main_g1
}

var cache_Main_f2 gopurs_runtime.Value
var once_Main_f2 sync.Once

func Get_Main_f2() gopurs_runtime.Value {
	once_Main_f2.Do(func() {
		cache_Main_f2 = gopurs_runtime.Func(func(dictEg2_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_f2(gopurs_runtime.CoerceToStruct[Constructor_Main_Eg2[gopurs_runtime.Value, gopurs_runtime.Value]](dictEg2_0_box))
		})
	})
	return cache_Main_f2
}

var cache_Main_f1 gopurs_runtime.Value
var once_Main_f1 sync.Once

func Get_Main_f1() gopurs_runtime.Value {
	once_Main_f1.Do(func() {
		cache_Main_f1 = gopurs_runtime.Func(func(dictEg1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_f1(gopurs_runtime.CoerceToStruct[Constructor_Main_Eg1[gopurs_runtime.Value, gopurs_runtime.Value]](dictEg1_0_box))
		})
	})
	return cache_Main_f1
}

type Constructor_Main_Eg2[T_f any, T_g any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}

func init() {
	gopurs_runtime.StructGetters[2207147950] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_Eg2[gopurs_runtime.Value, gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "Functor0":
			return gopurs_runtime.Box(c.V0)
		case "Functor1":
			return gopurs_runtime.Box(c.V1)
		default:
			panic("Key not found in dictionary Constructor_Main_Eg2: " + key)
		}
	}
}

type Constructor_Main_Eg1[T_f any, T_g any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}

func init() {
	gopurs_runtime.StructGetters[1828296749] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_Eg1[gopurs_runtime.Value, gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "Functor0":
			return gopurs_runtime.Box(c.V0)
		case "Functor1":
			return gopurs_runtime.Box(c.V1)
		default:
			panic("Key not found in dictionary Constructor_Main_Eg1: " + key)
		}
	}
}

func Call_Main_Eg2_dollar_Dict(x_0_loop struct {
	Functor0 gopurs_runtime.Value
	Functor1 gopurs_runtime.Value
}) *Constructor_Main_Eg2[gopurs_runtime.Value, gopurs_runtime.Value] {
	var x_0 struct {
		Functor0 gopurs_runtime.Value
		Functor1 gopurs_runtime.Value
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Main_Eg2[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict2("Functor0", "Functor1", orig.Functor0, orig.Functor1)
	}())
}

func Call_Main_Eg1_dollar_Dict(x_0_loop struct {
	Functor0 gopurs_runtime.Value
	Functor1 gopurs_runtime.Value
}) *Constructor_Main_Eg1[gopurs_runtime.Value, gopurs_runtime.Value] {
	var x_0 struct {
		Functor0 gopurs_runtime.Value
		Functor1 gopurs_runtime.Value
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Main_Eg1[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict2("Functor0", "Functor1", orig.Functor0, orig.Functor1)
	}())
}

func Call_Main_g2(dictEg2_0_loop *Constructor_Main_Eg2[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictEg2_0 *Constructor_Main_Eg2[gopurs_runtime.Value, gopurs_runtime.Value] = dictEg2_0_loop
	_ = dictEg2_0
	return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictEg2_0.V0, gopurs_runtime.Value{}), "map"), Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))}))
}

func Call_Main_g1(dictEg1_0_loop *Constructor_Main_Eg1[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictEg1_0 *Constructor_Main_Eg1[gopurs_runtime.Value, gopurs_runtime.Value] = dictEg1_0_loop
	_ = dictEg1_0
	return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictEg1_0.V1, gopurs_runtime.Value{}), "map"), Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))}))
}

func Call_Main_f2(dictEg2_0_loop *Constructor_Main_Eg2[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictEg2_0 *Constructor_Main_Eg2[gopurs_runtime.Value, gopurs_runtime.Value] = dictEg2_0_loop
	_ = dictEg2_0
	return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictEg2_0.V1, gopurs_runtime.Value{}), "map"), Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))}))
}

func Call_Main_f1(dictEg1_0_loop *Constructor_Main_Eg1[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictEg1_0 *Constructor_Main_Eg1[gopurs_runtime.Value, gopurs_runtime.Value] = dictEg1_0_loop
	_ = dictEg1_0
	return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictEg1_0.V0, gopurs_runtime.Value{}), "map"), Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))}))
}
