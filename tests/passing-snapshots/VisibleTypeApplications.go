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
			return gopurs_runtime.Value{Type: 9, IntVal: 3604601968, UnsafePtr: unsafe.Pointer((&Constructor_Main_Leaf{1, value0}))}
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
				return gopurs_runtime.Value{Type: 9, IntVal: 2447690122, UnsafePtr: unsafe.Pointer((&Constructor_Main_Branch{1, value0, value1}))}
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
			return Call_Main_ConstClass_dollar_Dict(x_0_box)
		})
	})
	return cache_Main_ConstClass_dollar_Dict
}

var cache_Main_constClass1 gopurs_runtime.Value
var once_Main_constClass1 sync.Once

func Get_Main_constClass1() gopurs_runtime.Value {
	once_Main_constClass1.Do(func() {
		cache_Main_constClass1 = gopurs_runtime.Value{Type: 9, IntVal: 4125467925, UnsafePtr: unsafe.Pointer((&Constructor_Main_ConstClass{1, gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return a_0
			})
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
		cache_Main_identityCheck = gopurs_runtime.Int(0)
	})
	return cache_Main_identityCheck
}

var cache_Main_identityPass gopurs_runtime.Value
var once_Main_identityPass sync.Once

func Get_Main_identityPass() gopurs_runtime.Value {
	once_Main_identityPass.Do(func() {
		cache_Main_identityPass = gopurs_runtime.Int(0)
	})
	return cache_Main_identityPass
}

var cache_Main_constClass gopurs_runtime.Value
var once_Main_constClass sync.Once

func Get_Main_constClass() gopurs_runtime.Value {
	once_Main_constClass.Do(func() {
		cache_Main_constClass = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_constClass(gopurs_runtime.CoerceToStruct[Constructor_Main_ConstClass](dict_0_box))
		})
	})
	return cache_Main_constClass
}

var cache_Main_constClassInt gopurs_runtime.Value
var once_Main_constClassInt sync.Once

func Get_Main_constClassInt() gopurs_runtime.Value {
	once_Main_constClassInt.Do(func() {
		cache_Main_constClassInt = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_constClassInt(a_0_box.IntVal, v_1_box.FloatVal()))
		})
	})
	return cache_Main_constClassInt
}

var cache_Main_constCheck gopurs_runtime.Value
var once_Main_constCheck sync.Once

func Get_Main_constCheck() gopurs_runtime.Value {
	once_Main_constCheck.Do(func() {
		cache_Main_constCheck = gopurs_runtime.Int(0)
	})
	return cache_Main_constCheck
}

var cache_Main_constPass gopurs_runtime.Value
var once_Main_constPass sync.Once

func Get_Main_constPass() gopurs_runtime.Value {
	once_Main_constPass.Do(func() {
		cache_Main_constPass = gopurs_runtime.Int(0)
	})
	return cache_Main_constPass
}

var cache_Main_constClass1__3802304696 gopurs_runtime.Value
var once_Main_constClass1__3802304696 sync.Once

func Get_Main_constClass1__3802304696() gopurs_runtime.Value {
	once_Main_constClass1__3802304696.Do(func() {
		cache_Main_constClass1__3802304696 = gopurs_runtime.Value{Type: 9, IntVal: 4125467925, UnsafePtr: unsafe.Pointer((&Constructor_Main_ConstClass{1, gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Int(a_0.IntVal)
			})
		})}))}
	})
	return cache_Main_constClass1__3802304696
}

type Constructor_Main_Leaf struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

type Constructor_Main_Branch struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}

type Constructor_Main_ConstClass struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

func init() {
	gopurs_runtime.StructGetters[4125467925] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_ConstClass)(ptr)
		_ = c
		switch key {
		case "constClass":
			return gopurs_runtime.Box(c.V0)
		default:
			panic("Key not found in dictionary Constructor_Main_ConstClass: " + key)
		}
	}
}

func Call_Main_ConstClass_dollar_Dict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_constClass(dict_0_loop *Constructor_Main_ConstClass) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_ConstClass = dict_0_loop
	_ = dict_0
	return gopurs_runtime.Box(dict_0.V0)
}

func Call_Main_constClassInt(a_0_loop int64, v_1_loop float64) int64 {
	var a_0 int64 = a_0_loop
	_ = a_0
	var v_1 float64 = v_1_loop
	_ = v_1
	return gopurs_runtime.Int(a_0).IntVal
}
