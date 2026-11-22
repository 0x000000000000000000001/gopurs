package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Divisible_Divisible_dollar_Dict gopurs_runtime.Value
var once_Data_Divisible_Divisible_dollar_Dict sync.Once

func Get_Data_Divisible_Divisible_dollar_Dict() gopurs_runtime.Value {
	once_Data_Divisible_Divisible_dollar_Dict.Do(func() {
		cache_Data_Divisible_Divisible_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2265116602, UnsafePtr: unsafe.Pointer(Call_Data_Divisible_Divisible_dollar_Dict(func() struct {
				Divide0 gopurs_runtime.Value
				conquer gopurs_runtime.Value
				/*gopurs:record-v1:2:0044006900760069006400650030,0063006f006e0071007500650072*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Divide0 gopurs_runtime.Value
					conquer gopurs_runtime.Value
					/*gopurs:record-v1:2:0044006900760069006400650030,0063006f006e0071007500650072*/
				}{}
				clone.Divide0 = gopurs_runtime.RecordGet(orig, "Divide0")
				clone.conquer = gopurs_runtime.RecordGet(orig, "conquer")
				return clone
			}()))}
		})
	})
	return cache_Data_Divisible_Divisible_dollar_Dict
}

var cache_Data_Divisible_divisiblePredicate gopurs_runtime.Value
var once_Data_Divisible_divisiblePredicate sync.Once

func Get_Data_Divisible_divisiblePredicate() gopurs_runtime.Value {
	once_Data_Divisible_divisiblePredicate.Do(func() {
		cache_Data_Divisible_divisiblePredicate = gopurs_runtime.Value{Type: 9, IntVal: 2265116602, UnsafePtr: unsafe.Pointer((&Constructor_Data_Divisible_Divisible[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2642321722, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Divide_Divide[gopurs_runtime.Value]](Get_Data_Divide_dividePredicate()))}
		}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(true)
		})}))}
	})
	return cache_Data_Divisible_divisiblePredicate
}

var cache_Data_Divisible_divisibleOp gopurs_runtime.Value
var once_Data_Divisible_divisibleOp sync.Once

func Get_Data_Divisible_divisibleOp() gopurs_runtime.Value {
	once_Data_Divisible_divisibleOp.Do(func() {
		cache_Data_Divisible_divisibleOp = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Divisible_divisibleOp(dictMonoid_0_box)
		})
	})
	return cache_Data_Divisible_divisibleOp
}

var cache_Data_Divisible_divisibleEquivalence gopurs_runtime.Value
var once_Data_Divisible_divisibleEquivalence sync.Once

func Get_Data_Divisible_divisibleEquivalence() gopurs_runtime.Value {
	once_Data_Divisible_divisibleEquivalence.Do(func() {
		cache_Data_Divisible_divisibleEquivalence = gopurs_runtime.Value{Type: 9, IntVal: 2265116602, UnsafePtr: unsafe.Pointer((&Constructor_Data_Divisible_Divisible[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2642321722, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Divide_Divide[gopurs_runtime.Value]](Get_Data_Divide_divideEquivalence()))}
		}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(true)
		})}))}
	})
	return cache_Data_Divisible_divisibleEquivalence
}

var cache_Data_Divisible_divisibleComparison gopurs_runtime.Value
var once_Data_Divisible_divisibleComparison sync.Once

func Get_Data_Divisible_divisibleComparison() gopurs_runtime.Value {
	once_Data_Divisible_divisibleComparison.Do(func() {
		cache_Data_Divisible_divisibleComparison = gopurs_runtime.Value{Type: 9, IntVal: 2265116602, UnsafePtr: unsafe.Pointer((&Constructor_Data_Divisible_Divisible[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2642321722, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Divide_Divide[gopurs_runtime.Value]](Get_Data_Divide_divideComparison()))}
		}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
		})}))}
	})
	return cache_Data_Divisible_divisibleComparison
}

var cache_Data_Divisible_conquer gopurs_runtime.Value
var once_Data_Divisible_conquer sync.Once

func Get_Data_Divisible_conquer() gopurs_runtime.Value {
	once_Data_Divisible_conquer.Do(func() {
		cache_Data_Divisible_conquer = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Divisible_conquer(dict_0_box)
		})
	})
	return cache_Data_Divisible_conquer
}

type Constructor_Data_Divisible_Divisible[T_f any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}

//gopurs:init-v1:00670065006e006500720061007400650064002f0044006100740061005f0044006900760069007300690062006c0065:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f0044006100740061005f0044006900760069007300690062006c0065005f0044006900760069007300690062006c0065
func init() {
	gopurs_runtime.StructGetters[2265116602] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Divisible_Divisible[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "Divide0":
			return gopurs_runtime.Box(c.V0)
		case "conquer":
			return gopurs_runtime.Box(c.V1)
		default:
			panic("Key not found in dictionary Constructor_Data_Divisible_Divisible: " + key)
		}
	}
}

func Call_Data_Divisible_Divisible_dollar_Dict(x_0_loop struct {
	Divide0 gopurs_runtime.Value
	conquer gopurs_runtime.Value
	/*gopurs:record-v1:2:0044006900760069006400650030,0063006f006e0071007500650072*/
}) *Constructor_Data_Divisible_Divisible[gopurs_runtime.Value] {
	var x_0 struct {
		Divide0 gopurs_runtime.Value
		conquer gopurs_runtime.Value
		/*gopurs:record-v1:2:0044006900760069006400650030,0063006f006e0071007500650072*/
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Data_Divisible_Divisible[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict2("Divide0", "conquer", orig.Divide0, orig.conquer)
	}())
}

func Call_Data_Divisible_divisibleOp(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
	_ = dictMonoid_0
	// TAST (Let): divideOp_1_0 shape=App(Var) bindingType=(ADT ["Data","Divide","Divide"] [(Func [(TypeVar b)] (TypeVar r$scope5))])
	divideOp_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Divide_Divide[gopurs_runtime.Value]](Call_Data_Divide_divideOp(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})))
	_ = divideOp_1_0
	// TAST (Let): __local_var_2_1 shape=Other bindingType=Any
	__local_var_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
	_ = __local_var_2_1
	return gopurs_runtime.Value{Type: 9, IntVal: 2265116602, UnsafePtr: unsafe.Pointer((&Constructor_Data_Divisible_Divisible[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 2642321722, UnsafePtr: unsafe.Pointer(divideOp_1_0)}
	}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return __local_var_2_1
	})}))}
}

func Call_Data_Divisible_conquer(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dict_0 gopurs_runtime.Value = dict_0_loop
	_ = dict_0
	return gopurs_runtime.RecordGet(dict_0, "conquer")
}
