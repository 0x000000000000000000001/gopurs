package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Decidable_identity gopurs_runtime.Value
var once_Data_Decidable_identity sync.Once

func Get_Data_Decidable_identity() gopurs_runtime.Value {
	once_Data_Decidable_identity.Do(func() {
		cache_Data_Decidable_identity = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Decidable_identity(x_0_box)
		})
	})
	return cache_Data_Decidable_identity
}

var cache_Data_Decidable_Decidable_dollar_Dict gopurs_runtime.Value
var once_Data_Decidable_Decidable_dollar_Dict sync.Once

func Get_Data_Decidable_Decidable_dollar_Dict() gopurs_runtime.Value {
	once_Data_Decidable_Decidable_dollar_Dict.Do(func() {
		cache_Data_Decidable_Decidable_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 336732346, UnsafePtr: unsafe.Pointer(Call_Data_Decidable_Decidable_dollar_Dict(func() struct {
				Decide0    gopurs_runtime.Value
				Divisible1 gopurs_runtime.Value
				lose       gopurs_runtime.Value
				/*gopurs:record-v1:3:0044006500630069006400650030,0044006900760069007300690062006c00650031,006c006f00730065*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Decide0    gopurs_runtime.Value
					Divisible1 gopurs_runtime.Value
					lose       gopurs_runtime.Value
					/*gopurs:record-v1:3:0044006500630069006400650030,0044006900760069007300690062006c00650031,006c006f00730065*/
				}{}
				clone.Decide0 = gopurs_runtime.RecordGet(orig, "Decide0")
				clone.Divisible1 = gopurs_runtime.RecordGet(orig, "Divisible1")
				clone.lose = gopurs_runtime.RecordGet(orig, "lose")
				return clone
			}()))}
		})
	})
	return cache_Data_Decidable_Decidable_dollar_Dict
}

var cache_Data_Decidable_lose gopurs_runtime.Value
var once_Data_Decidable_lose sync.Once

func Get_Data_Decidable_lose() gopurs_runtime.Value {
	once_Data_Decidable_lose.Do(func() {
		cache_Data_Decidable_lose = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Decidable_lose(gopurs_runtime.CoerceToStruct[Constructor_Data_Decidable_Decidable[gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Data_Decidable_lose
}

var cache_Data_Decidable_lost gopurs_runtime.Value
var once_Data_Decidable_lost sync.Once

func Get_Data_Decidable_lost() gopurs_runtime.Value {
	once_Data_Decidable_lost.Do(func() {
		cache_Data_Decidable_lost = gopurs_runtime.Func(func(dictDecidable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Decidable_lost(dictDecidable_0_box)
		})
	})
	return cache_Data_Decidable_lost
}

var cache_Data_Decidable_decidablePredicate gopurs_runtime.Value
var once_Data_Decidable_decidablePredicate sync.Once

func Get_Data_Decidable_decidablePredicate() gopurs_runtime.Value {
	once_Data_Decidable_decidablePredicate.Do(func() {
		cache_Data_Decidable_decidablePredicate = gopurs_runtime.Value{Type: 9, IntVal: 336732346, UnsafePtr: unsafe.Pointer((&Constructor_Data_Decidable_Decidable[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1618621146, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Decide_Decide[gopurs_runtime.Value]](Get_Data_Decide_choosePredicate()))}
		}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2265116602, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Divisible_Divisible[gopurs_runtime.Value]](Get_Data_Divisible_divisiblePredicate()))}
		}), gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, a_1 gopurs_runtime.Value) gopurs_runtime.Value {
			var Call_local_Data_Decidable_spin__2182866177_2_0_0 func(gopurs_runtime.Value) bool
			_ = Call_local_Data_Decidable_spin__2182866177_2_0_0
			var spin__2182866177_2_0_0 gopurs_runtime.Value
			_ = spin__2182866177_2_0_0
			var Call_local_Data_Decidable_spin_2_1_1 func(gopurs_runtime.Value) bool
			_ = Call_local_Data_Decidable_spin_2_1_1
			var spin_2_1_1 gopurs_runtime.Value
			_ = spin_2_1_1
			Call_local_Data_Decidable_spin__2182866177_2_0_0 = func(v_3_loop gopurs_runtime.Value) bool {
			spin__2182866177_2_0_0:
				for {
					if false {
						continue spin__2182866177_2_0_0
					}
					var v_3 gopurs_runtime.Value = v_3_loop
					_ = v_3
					v_3_loop = v_3
					continue spin__2182866177_2_0_0
					return func() bool { panic("unreachable") }()
				}
			}
			spin__2182866177_2_0_0 = gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Bool(Call_local_Data_Decidable_spin__2182866177_2_0_0(v_3_loop_val))
			})
			Call_local_Data_Decidable_spin_2_1_1 = func(v_3_loop gopurs_runtime.Value) bool {
			spin_2_1_1:
				for {
					if false {
						continue spin_2_1_1
					}
					var v_3 gopurs_runtime.Value = v_3_loop
					_ = v_3
					return Call_local_Data_Decidable_spin__2182866177_2_0_0(v_3)
				}
			}
			spin_2_1_1 = gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Bool(Call_local_Data_Decidable_spin_2_1_1(v_3_loop_val))
			})
			return gopurs_runtime.Bool(Call_local_Data_Decidable_spin__2182866177_2_0_0(gopurs_runtime.Apply(f_0, a_1)))
		})}))}
	})
	return cache_Data_Decidable_decidablePredicate
}

var cache_Data_Decidable_decidableOp gopurs_runtime.Value
var once_Data_Decidable_decidableOp sync.Once

func Get_Data_Decidable_decidableOp() gopurs_runtime.Value {
	once_Data_Decidable_decidableOp.Do(func() {
		cache_Data_Decidable_decidableOp = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Decidable_decidableOp(dictMonoid_0_box)
		})
	})
	return cache_Data_Decidable_decidableOp
}

var cache_Data_Decidable_decidableEquivalence gopurs_runtime.Value
var once_Data_Decidable_decidableEquivalence sync.Once

func Get_Data_Decidable_decidableEquivalence() gopurs_runtime.Value {
	once_Data_Decidable_decidableEquivalence.Do(func() {
		cache_Data_Decidable_decidableEquivalence = gopurs_runtime.Value{Type: 9, IntVal: 336732346, UnsafePtr: unsafe.Pointer((&Constructor_Data_Decidable_Decidable[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1618621146, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Decide_Decide[gopurs_runtime.Value]](Get_Data_Decide_chooseEquivalence()))}
		}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2265116602, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Divisible_Divisible[gopurs_runtime.Value]](Get_Data_Divisible_divisibleEquivalence()))}
		}), gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, a_1 gopurs_runtime.Value) gopurs_runtime.Value {
			var Call_local_Data_Decidable_spin__1769020947_2_0_4 func(gopurs_runtime.Value) gopurs_runtime.Value
			_ = Call_local_Data_Decidable_spin__1769020947_2_0_4
			var spin__1769020947_2_0_4 gopurs_runtime.Value
			_ = spin__1769020947_2_0_4
			var Call_local_Data_Decidable_spin_2_1_5 func(gopurs_runtime.Value) gopurs_runtime.Value
			_ = Call_local_Data_Decidable_spin_2_1_5
			var spin_2_1_5 gopurs_runtime.Value
			_ = spin_2_1_5
			Call_local_Data_Decidable_spin__1769020947_2_0_4 = func(v_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
			spin__1769020947_2_0_4:
				for {
					if false {
						continue spin__1769020947_2_0_4
					}
					var v_3 gopurs_runtime.Value = v_3_loop
					_ = v_3
					v_3_loop = v_3
					continue spin__1769020947_2_0_4
					return func() gopurs_runtime.Value { panic("unreachable") }()
				}
			}
			spin__1769020947_2_0_4 = gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return Call_local_Data_Decidable_spin__1769020947_2_0_4(v_3_loop_val)
			})
			Call_local_Data_Decidable_spin_2_1_5 = func(v_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
			spin_2_1_5:
				for {
					if false {
						continue spin_2_1_5
					}
					var v_3 gopurs_runtime.Value = v_3_loop
					_ = v_3
					return Call_local_Data_Decidable_spin__1769020947_2_0_4(v_3)
				}
			}
			spin_2_1_5 = gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return Call_local_Data_Decidable_spin_2_1_5(v_3_loop_val)
			})
			return Call_local_Data_Decidable_spin__1769020947_2_0_4(gopurs_runtime.Apply(f_0, a_1))
		})}))}
	})
	return cache_Data_Decidable_decidableEquivalence
}

var cache_Data_Decidable_decidableComparison gopurs_runtime.Value
var once_Data_Decidable_decidableComparison sync.Once

func Get_Data_Decidable_decidableComparison() gopurs_runtime.Value {
	once_Data_Decidable_decidableComparison.Do(func() {
		cache_Data_Decidable_decidableComparison = gopurs_runtime.Value{Type: 9, IntVal: 336732346, UnsafePtr: unsafe.Pointer((&Constructor_Data_Decidable_Decidable[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1618621146, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Decide_Decide[gopurs_runtime.Value]](Get_Data_Decide_chooseComparison()))}
		}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2265116602, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Divisible_Divisible[gopurs_runtime.Value]](Get_Data_Divisible_divisibleComparison()))}
		}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, a_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
			var Call_local_Data_Decidable_spin__1133204955_3_0_6 func(gopurs_runtime.Value) uint32
			_ = Call_local_Data_Decidable_spin__1133204955_3_0_6
			var spin__1133204955_3_0_6 gopurs_runtime.Value
			_ = spin__1133204955_3_0_6
			var Call_local_Data_Decidable_spin_3_1_7 func(gopurs_runtime.Value) uint32
			_ = Call_local_Data_Decidable_spin_3_1_7
			var spin_3_1_7 gopurs_runtime.Value
			_ = spin_3_1_7
			Call_local_Data_Decidable_spin__1133204955_3_0_6 = func(v_4_loop gopurs_runtime.Value) uint32 {
			spin__1133204955_3_0_6:
				for {
					if false {
						continue spin__1133204955_3_0_6
					}
					var v_4 gopurs_runtime.Value = v_4_loop
					_ = v_4
					v_4_loop = v_4
					continue spin__1133204955_3_0_6
					return func() uint32 { panic("unreachable") }()
				}
			}
			spin__1133204955_3_0_6 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_local_Data_Decidable_spin__1133204955_3_0_6(v_4_loop_val)), UnsafePtr: nil}
			})
			Call_local_Data_Decidable_spin_3_1_7 = func(v_4_loop gopurs_runtime.Value) uint32 {
			spin_3_1_7:
				for {
					if false {
						continue spin_3_1_7
					}
					var v_4 gopurs_runtime.Value = v_4_loop
					_ = v_4
					return Call_local_Data_Decidable_spin__1133204955_3_0_6(v_4)
				}
			}
			spin_3_1_7 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_local_Data_Decidable_spin_3_1_7(v_4_loop_val)), UnsafePtr: nil}
			})
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_local_Data_Decidable_spin__1133204955_3_0_6(gopurs_runtime.Apply(f_0, a_1))), UnsafePtr: nil}
		})}))}
	})
	return cache_Data_Decidable_decidableComparison
}

type Constructor_Data_Decidable_Decidable[T_f any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
}

//gopurs:init-v1:00670065006e006500720061007400650064002f0044006100740061005f0044006500630069006400610062006c0065:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f0044006100740061005f0044006500630069006400610062006c0065005f0044006500630069006400610062006c0065
func init() {
	gopurs_runtime.StructGetters[336732346] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Decidable_Decidable[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "Decide0":
			return gopurs_runtime.Box(c.V0)
		case "Divisible1":
			return gopurs_runtime.Box(c.V1)
		case "lose":
			return gopurs_runtime.Box(c.V2)
		default:
			panic("Key not found in dictionary Constructor_Data_Decidable_Decidable: " + key)
		}
	}
}

func Call_Data_Decidable_identity(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Data_Decidable_Decidable_dollar_Dict(x_0_loop struct {
	Decide0    gopurs_runtime.Value
	Divisible1 gopurs_runtime.Value
	lose       gopurs_runtime.Value
	/*gopurs:record-v1:3:0044006500630069006400650030,0044006900760069007300690062006c00650031,006c006f00730065*/
}) *Constructor_Data_Decidable_Decidable[gopurs_runtime.Value] {
	var x_0 struct {
		Decide0    gopurs_runtime.Value
		Divisible1 gopurs_runtime.Value
		lose       gopurs_runtime.Value
		/*gopurs:record-v1:3:0044006500630069006400650030,0044006900760069007300690062006c00650031,006c006f00730065*/
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Data_Decidable_Decidable[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict3("Decide0", "Divisible1", "lose", orig.Decide0, orig.Divisible1, orig.lose)
	}())
}

func Call_Data_Decidable_lose(dict_0_loop *Constructor_Data_Decidable_Decidable[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Data_Decidable_Decidable[gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V2
}

func Call_Data_Decidable_lost(dictDecidable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictDecidable_0 gopurs_runtime.Value = dictDecidable_0_loop
	_ = dictDecidable_0
	return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictDecidable_0, "lose"), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return x_1
	}))
}

func Call_Data_Decidable_decidableOp(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
	_ = dictMonoid_0
	// TAST (Let): chooseOp_1_0 shape=App(Var) bindingType=(ADT ["Data","Decide","Decide"] [(Func [(TypeVar b)] (TypeVar r$scope10))])
	chooseOp_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Decide_Decide[gopurs_runtime.Value]](Call_Data_Decide_chooseOp(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})))
	_ = chooseOp_1_0
	// TAST (Let): divisibleOp_2_1 shape=App(Var) bindingType=(ADT ["Data","Divisible","Divisible"] [(Func [(TypeVar b)] (TypeVar r$scope10))])
	divisibleOp_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Divisible_Divisible[gopurs_runtime.Value]](Call_Data_Divisible_divisibleOp(dictMonoid_0))
	_ = divisibleOp_2_1
	return gopurs_runtime.Value{Type: 9, IntVal: 336732346, UnsafePtr: unsafe.Pointer((&Constructor_Data_Decidable_Decidable[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 1618621146, UnsafePtr: unsafe.Pointer(chooseOp_1_0)}
	}), gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 2265116602, UnsafePtr: unsafe.Pointer(divisibleOp_2_1)}
	}), gopurs_runtime.Func2(func(f_3 gopurs_runtime.Value, a_4 gopurs_runtime.Value) gopurs_runtime.Value {
		var Call_local_Data_Decidable_spin__1769020947_5_2_2 func(gopurs_runtime.Value) gopurs_runtime.Value
		_ = Call_local_Data_Decidable_spin__1769020947_5_2_2
		var spin__1769020947_5_2_2 gopurs_runtime.Value
		_ = spin__1769020947_5_2_2
		var Call_local_Data_Decidable_spin_5_3_3 func(gopurs_runtime.Value) gopurs_runtime.Value
		_ = Call_local_Data_Decidable_spin_5_3_3
		var spin_5_3_3 gopurs_runtime.Value
		_ = spin_5_3_3
		Call_local_Data_Decidable_spin__1769020947_5_2_2 = func(v_6_loop gopurs_runtime.Value) gopurs_runtime.Value {
		spin__1769020947_5_2_2:
			for {
				if false {
					continue spin__1769020947_5_2_2
				}
				var v_6 gopurs_runtime.Value = v_6_loop
				_ = v_6
				v_6_loop = v_6
				continue spin__1769020947_5_2_2
				return func() gopurs_runtime.Value { panic("unreachable") }()
			}
		}
		spin__1769020947_5_2_2 = gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_local_Data_Decidable_spin__1769020947_5_2_2(v_6_loop_val)
		})
		Call_local_Data_Decidable_spin_5_3_3 = func(v_6_loop gopurs_runtime.Value) gopurs_runtime.Value {
		spin_5_3_3:
			for {
				if false {
					continue spin_5_3_3
				}
				var v_6 gopurs_runtime.Value = v_6_loop
				_ = v_6
				return Call_local_Data_Decidable_spin__1769020947_5_2_2(v_6)
			}
		}
		spin_5_3_3 = gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_local_Data_Decidable_spin_5_3_3(v_6_loop_val)
		})
		return Call_local_Data_Decidable_spin__1769020947_5_2_2(gopurs_runtime.Apply(f_3, a_4))
	})}))}
}
