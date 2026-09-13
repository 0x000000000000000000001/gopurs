package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_Leaf gopurs_runtime.Value
var once_Main_Leaf sync.Once

func Get_Main_Leaf() gopurs_runtime.Value {
	once_Main_Leaf.Do(func() {
		cache_Main_Leaf = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3604601968, UnsafePtr: unsafe.Pointer((&Constructor_Main_Leaf[gopurs_runtime.Value]{1, value0}))}
		})
	})
	return cache_Main_Leaf
}

var cache_Main_Branch gopurs_runtime.Value
var once_Main_Branch sync.Once

func Get_Main_Branch() gopurs_runtime.Value {
	once_Main_Branch.Do(func() {
		cache_Main_Branch = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 2447690122, UnsafePtr: unsafe.Pointer((&Constructor_Main_Branch[gopurs_runtime.Value]{1, value0, value1}))}
			})
		})
	})
	return cache_Main_Branch
}

var cache_Main_ConstClass_dollar_Dict gopurs_runtime.Value
var once_Main_ConstClass_dollar_Dict sync.Once

func Get_Main_ConstClass_dollar_Dict() gopurs_runtime.Value {
	once_Main_ConstClass_dollar_Dict.Do(func() {
		cache_Main_ConstClass_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4125467925, UnsafePtr: unsafe.Pointer(Call_Main_ConstClass_dollar_Dict(func() struct {
				constClass gopurs_runtime.Value
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					constClass gopurs_runtime.Value
				}{}
				clone.constClass = gopurs_runtime.RecordGet(orig, "constClass")
				return clone
			}()))}
		})
	})
	return cache_Main_ConstClass_dollar_Dict
}

var cache_Main_constClass1 gopurs_runtime.Value
var once_Main_constClass1 sync.Once

func Get_Main_constClass1() gopurs_runtime.Value {
	once_Main_constClass1.Do(func() {
		cache_Main_constClass1 = gopurs_runtime.Value{Type: 9, IntVal: 4125467925, UnsafePtr: unsafe.Pointer((&Constructor_Main_ConstClass[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(a_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return a_0
		})}))}
	})
	return cache_Main_constClass1
}

var cache_Main_treeInt_prime_ gopurs_runtime.Value
var once_Main_treeInt_prime_ sync.Once

func Get_Main_treeInt_prime_() gopurs_runtime.Value {
	once_Main_treeInt_prime_.Do(func() {
		cache_Main_treeInt_prime_ = Get_Main_Branch()
	})
	return cache_Main_treeInt_prime_
}

var cache_Main_treeInt gopurs_runtime.Value
var once_Main_treeInt sync.Once

func Get_Main_treeInt() gopurs_runtime.Value {
	once_Main_treeInt.Do(func() {
		cache_Main_treeInt = Get_Main_Leaf()
	})
	return cache_Main_treeInt
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

var cache_Main_identityCheck gopurs_runtime.Value
var once_Main_identityCheck sync.Once

func Get_Main_identityCheck() gopurs_runtime.Value {
	once_Main_identityCheck.Do(func() {
		cache_Main_identityCheck = gopurs_runtime.Int(int64(0))
	})
	return cache_Main_identityCheck
}

var cache_Main_identityPass gopurs_runtime.Value
var once_Main_identityPass sync.Once

func Get_Main_identityPass() gopurs_runtime.Value {
	once_Main_identityPass.Do(func() {
		cache_Main_identityPass = gopurs_runtime.Int(int64(0))
	})
	return cache_Main_identityPass
}

var cache_Main_constClass gopurs_runtime.Value
var once_Main_constClass sync.Once

func Get_Main_constClass() gopurs_runtime.Value {
	once_Main_constClass.Do(func() {
		cache_Main_constClass = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_constClass(gopurs_runtime.CoerceToStruct[Constructor_Main_ConstClass[gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Main_constClass
}

var cache_Main_constClassInt gopurs_runtime.Value
var once_Main_constClassInt sync.Once

func Get_Main_constClassInt() gopurs_runtime.Value {
	once_Main_constClassInt.Do(func() {
		cache_Main_constClassInt = Call_Main_constClass(Rebox_Main_662423258_1962053153(Rebox_Main_1962053153_662423258(gopurs_runtime.CoerceToStruct[Constructor_Main_ConstClass[gopurs_runtime.Value]](Get_Main_constClass1()))))
	})
	return cache_Main_constClassInt
}

var cache_Main_constCheck gopurs_runtime.Value
var once_Main_constCheck sync.Once

func Get_Main_constCheck() gopurs_runtime.Value {
	once_Main_constCheck.Do(func() {
		cache_Main_constCheck = gopurs_runtime.Int(int64(0))
	})
	return cache_Main_constCheck
}

var cache_Main_constPass gopurs_runtime.Value
var once_Main_constPass sync.Once

func Get_Main_constPass() gopurs_runtime.Value {
	once_Main_constPass.Do(func() {
		cache_Main_constPass = gopurs_runtime.Int(int64(0))
	})
	return cache_Main_constPass
}

type Constructor_Main_Leaf[T_a any] struct {
	Rc uint32
	V0 T_a
}

type Constructor_Main_Branch[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}

type Constructor_Main_ConstClass[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

func init() {
	gopurs_runtime.StructGetters[4125467925] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_ConstClass[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "constClass":
			return gopurs_runtime.Box(c.V0)
		default:
			panic("Key not found in dictionary Constructor_Main_ConstClass: " + key)
		}
	}
}

func Call_Main_ConstClass_dollar_Dict(x_0_loop struct {
	constClass gopurs_runtime.Value
}) *Constructor_Main_ConstClass[gopurs_runtime.Value] {
	var x_0 struct {
		constClass gopurs_runtime.Value
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Main_ConstClass[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict1("constClass", orig.constClass)
	}())
}

func Call_Main_constClass(dict_0_loop *Constructor_Main_ConstClass[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_ConstClass[gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V0
}

func Rebox_Main_1962053153_662423258(in *Constructor_Main_ConstClass[gopurs_runtime.Value]) *Constructor_Main_ConstClass[int64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_ConstClass[int64]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_662423258_1962053153(in *Constructor_Main_ConstClass[int64]) *Constructor_Main_ConstClass[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_ConstClass[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}
