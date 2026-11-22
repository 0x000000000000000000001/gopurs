package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Decide_compose gopurs_runtime.Value
var once_Data_Decide_compose sync.Once

func Get_Data_Decide_compose() gopurs_runtime.Value {
	once_Data_Decide_compose.Do(func() {
		cache_Data_Decide_compose = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Decide_compose
}

var cache_Data_Decide_compose1 gopurs_runtime.Value
var once_Data_Decide_compose1 sync.Once

func Get_Data_Decide_compose1() gopurs_runtime.Value {
	once_Data_Decide_compose1.Do(func() {
		cache_Data_Decide_compose1 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Decide_compose1
}

var cache_Data_Decide_identity gopurs_runtime.Value
var once_Data_Decide_identity sync.Once

func Get_Data_Decide_identity() gopurs_runtime.Value {
	once_Data_Decide_identity.Do(func() {
		cache_Data_Decide_identity = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Decide_identity(x_0_box)
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})
	})
	return cache_Data_Decide_identity
}

var cache_Data_Decide_Decide_dollar_Dict gopurs_runtime.Value
var once_Data_Decide_Decide_dollar_Dict sync.Once

func Get_Data_Decide_Decide_dollar_Dict() gopurs_runtime.Value {
	once_Data_Decide_Decide_dollar_Dict.Do(func() {
		cache_Data_Decide_Decide_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1618621146, UnsafePtr: unsafe.Pointer(Call_Data_Decide_Decide_dollar_Dict(func() struct {
				Divide0 gopurs_runtime.Value
				choose  gopurs_runtime.Value
				/*gopurs:record-v1:2:0044006900760069006400650030,00630068006f006f00730065*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Divide0 gopurs_runtime.Value
					choose  gopurs_runtime.Value
					/*gopurs:record-v1:2:0044006900760069006400650030,00630068006f006f00730065*/
				}{}
				clone.Divide0 = gopurs_runtime.RecordGet(orig, "Divide0")
				clone.choose = gopurs_runtime.RecordGet(orig, "choose")
				return clone
			}()))}
		})
	})
	return cache_Data_Decide_Decide_dollar_Dict
}

var cache_Data_Decide_choosePredicate gopurs_runtime.Value
var once_Data_Decide_choosePredicate sync.Once

func Get_Data_Decide_choosePredicate() gopurs_runtime.Value {
	once_Data_Decide_choosePredicate.Do(func() {
		cache_Data_Decide_choosePredicate = gopurs_runtime.Value{Type: 9, IntVal: 1618621146, UnsafePtr: unsafe.Pointer((&Constructor_Data_Decide_Decide[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2642321722, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Divide_Divide[gopurs_runtime.Value]](Get_Data_Divide_dividePredicate()))}
		}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Func(func(v2_3 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t0 gopurs_runtime.Value
				{
					if v2_3.Type == 9 && v2_3.IntVal == 3711209382 {
						__t0 = gopurs_runtime.Apply(v_1, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_3.UnsafePtr).V0)
						goto end_branch_0
					} else {

					}
				}
				{
					if v2_3.Type == 9 && v2_3.IntVal == 2465973597 {
						__t0 = gopurs_runtime.Apply(v1_2, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_3.UnsafePtr).V0)
						goto end_branch_0
					} else {

					}
				}
				{
					__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
				}
			end_branch_0:
				return __t0
			}), f_0)
		})}))}
	})
	return cache_Data_Decide_choosePredicate
}

var cache_Data_Decide_chooseOp gopurs_runtime.Value
var once_Data_Decide_chooseOp sync.Once

func Get_Data_Decide_chooseOp() gopurs_runtime.Value {
	once_Data_Decide_chooseOp.Do(func() {
		cache_Data_Decide_chooseOp = gopurs_runtime.Func(func(dictSemigroup_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Decide_chooseOp(dictSemigroup_0_box)
		})
	})
	return cache_Data_Decide_chooseOp
}

var cache_Data_Decide_chooseEquivalence gopurs_runtime.Value
var once_Data_Decide_chooseEquivalence sync.Once

func Get_Data_Decide_chooseEquivalence() gopurs_runtime.Value {
	once_Data_Decide_chooseEquivalence.Do(func() {
		cache_Data_Decide_chooseEquivalence = gopurs_runtime.Value{Type: 9, IntVal: 1618621146, UnsafePtr: unsafe.Pointer((&Constructor_Data_Decide_Decide[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2642321722, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Divide_Divide[gopurs_runtime.Value]](Get_Data_Divide_divideEquivalence()))}
		}), gopurs_runtime.Func5(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value, a_3 gopurs_runtime.Value, b_4 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): v2_5_0 shape=App(Other) bindingType=(ADT ["Data","Either","Either"] [(TypeVar b$scope26), (TypeVar c$scope27)])
			v2_5_0 := gopurs_runtime.Apply(f_0, a_3)
			_ = v2_5_0
			var __t5 gopurs_runtime.Value
			{
				if v2_5_0.Type == 9 && v2_5_0.IntVal == 3711209382 {
					// TAST (Let): v3_6_1 shape=App(Other) bindingType=(ADT ["Data","Either","Either"] [(TypeVar b$scope26), (TypeVar c$scope27)])
					v3_6_1 := gopurs_runtime.Apply(f_0, b_4)
					_ = v3_6_1
					var __t2 gopurs_runtime.Value
					{
						if v3_6_1.Type == 9 && v3_6_1.IntVal == 3711209382 {
							__t2 = gopurs_runtime.Apply2(v_1, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_5_0.UnsafePtr).V0, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v3_6_1.UnsafePtr).V0)
							goto end_branch_2
						} else {

						}
					}
					{
						if v3_6_1.Type == 9 && v3_6_1.IntVal == 2465973597 {
							__t2 = gopurs_runtime.Bool(false)
							goto end_branch_2
						} else {

						}
					}
					{
						__t2 = gopurs_runtime.Bool(func() bool { panic("Failed pattern match") }())
					}
				end_branch_2:
					__t5 = __t2
					goto end_branch_5
				} else {

				}
			}
			{
				if v2_5_0.Type == 9 && v2_5_0.IntVal == 2465973597 {
					// TAST (Let): v3_6_3 shape=App(Other) bindingType=(ADT ["Data","Either","Either"] [(TypeVar b$scope26), (TypeVar c$scope27)])
					v3_6_3 := gopurs_runtime.Apply(f_0, b_4)
					_ = v3_6_3
					var __t4 gopurs_runtime.Value
					{
						if v3_6_3.Type == 9 && v3_6_3.IntVal == 3711209382 {
							__t4 = gopurs_runtime.Bool(false)
							goto end_branch_4
						} else {

						}
					}
					{
						if v3_6_3.Type == 9 && v3_6_3.IntVal == 2465973597 {
							__t4 = gopurs_runtime.Apply2(v1_2, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_5_0.UnsafePtr).V0, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v3_6_3.UnsafePtr).V0)
							goto end_branch_4
						} else {

						}
					}
					{
						__t4 = gopurs_runtime.Bool(func() bool { panic("Failed pattern match") }())
					}
				end_branch_4:
					__t5 = __t4
					goto end_branch_5
				} else {

				}
			}
			{
				__t5 = gopurs_runtime.Bool(func() bool { panic("Failed pattern match") }())
			}
		end_branch_5:
			return __t5
		})}))}
	})
	return cache_Data_Decide_chooseEquivalence
}

var cache_Data_Decide_chooseComparison gopurs_runtime.Value
var once_Data_Decide_chooseComparison sync.Once

func Get_Data_Decide_chooseComparison() gopurs_runtime.Value {
	once_Data_Decide_chooseComparison.Do(func() {
		cache_Data_Decide_chooseComparison = gopurs_runtime.Value{Type: 9, IntVal: 1618621146, UnsafePtr: unsafe.Pointer((&Constructor_Data_Decide_Decide[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2642321722, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Divide_Divide[gopurs_runtime.Value]](Get_Data_Divide_divideComparison()))}
		}), gopurs_runtime.Func5(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value, a_3 gopurs_runtime.Value, b_4 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): v2_5_0 shape=App(Other) bindingType=(ADT ["Data","Either","Either"] [(TypeVar b$scope34), (TypeVar c$scope35)])
			v2_5_0 := gopurs_runtime.Apply(f_0, a_3)
			_ = v2_5_0
			var __t5 gopurs_runtime.Value
			{
				if v2_5_0.Type == 9 && v2_5_0.IntVal == 3711209382 {
					// TAST (Let): v3_6_1 shape=App(Other) bindingType=(ADT ["Data","Either","Either"] [(TypeVar b$scope34), (TypeVar c$scope35)])
					v3_6_1 := gopurs_runtime.Apply(f_0, b_4)
					_ = v3_6_1
					var __t2 gopurs_runtime.Value
					{
						if v3_6_1.Type == 9 && v3_6_1.IntVal == 3711209382 {
							__t2 = gopurs_runtime.Apply2(v_1, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_5_0.UnsafePtr).V0, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v3_6_1.UnsafePtr).V0)
							goto end_branch_2
						} else {

						}
					}
					{
						if v3_6_1.Type == 9 && v3_6_1.IntVal == 2465973597 {
							__t2 = gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}
							goto end_branch_2
						} else {

						}
					}
					{
						__t2 = gopurs_runtime.Value{Type: 9, IntVal: int64(func() uint32 { panic("Failed pattern match") }()), UnsafePtr: nil}
					}
				end_branch_2:
					__t5 = __t2
					goto end_branch_5
				} else {

				}
			}
			{
				if v2_5_0.Type == 9 && v2_5_0.IntVal == 2465973597 {
					// TAST (Let): v3_6_3 shape=App(Other) bindingType=(ADT ["Data","Either","Either"] [(TypeVar b$scope34), (TypeVar c$scope35)])
					v3_6_3 := gopurs_runtime.Apply(f_0, b_4)
					_ = v3_6_3
					var __t4 gopurs_runtime.Value
					{
						if v3_6_3.Type == 9 && v3_6_3.IntVal == 3711209382 {
							__t4 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
							goto end_branch_4
						} else {

						}
					}
					{
						if v3_6_3.Type == 9 && v3_6_3.IntVal == 2465973597 {
							__t4 = gopurs_runtime.Apply2(v1_2, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_5_0.UnsafePtr).V0, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v3_6_3.UnsafePtr).V0)
							goto end_branch_4
						} else {

						}
					}
					{
						__t4 = gopurs_runtime.Value{Type: 9, IntVal: int64(func() uint32 { panic("Failed pattern match") }()), UnsafePtr: nil}
					}
				end_branch_4:
					__t5 = __t4
					goto end_branch_5
				} else {

				}
			}
			{
				__t5 = gopurs_runtime.Value{Type: 9, IntVal: int64(func() uint32 { panic("Failed pattern match") }()), UnsafePtr: nil}
			}
		end_branch_5:
			return __t5
		})}))}
	})
	return cache_Data_Decide_chooseComparison
}

var cache_Data_Decide_choose gopurs_runtime.Value
var once_Data_Decide_choose sync.Once

func Get_Data_Decide_choose() gopurs_runtime.Value {
	once_Data_Decide_choose.Do(func() {
		cache_Data_Decide_choose = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Decide_choose(gopurs_runtime.CoerceToStruct[Constructor_Data_Decide_Decide[gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Data_Decide_choose
}

var cache_Data_Decide_chosen gopurs_runtime.Value
var once_Data_Decide_chosen sync.Once

func Get_Data_Decide_chosen() gopurs_runtime.Value {
	once_Data_Decide_chosen.Do(func() {
		cache_Data_Decide_chosen = gopurs_runtime.Func(func(dictDecide_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Decide_chosen(gopurs_runtime.CoerceToStruct[Constructor_Data_Decide_Decide[gopurs_runtime.Value]](dictDecide_0_box))
		})
	})
	return cache_Data_Decide_chosen
}

type Constructor_Data_Decide_Decide[T_f any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}

//gopurs:init-v1:00670065006e006500720061007400650064002f0044006100740061005f004400650063006900640065:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f0044006100740061005f004400650063006900640065005f004400650063006900640065
func init() {
	gopurs_runtime.StructGetters[1618621146] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Decide_Decide[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "Divide0":
			return gopurs_runtime.Box(c.V0)
		case "choose":
			return gopurs_runtime.Box(c.V1)
		default:
			panic("Key not found in dictionary Constructor_Data_Decide_Decide: " + key)
		}
	}
}

func Call_Data_Decide_identity(x_0_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 bool
} {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return func() struct {
		V0 gopurs_runtime.Value
		V1 gopurs_runtime.Value
		V2 bool
	} { _v := x_0; if _v.Type == 9 && _v.IntVal == 2465973597 && _v.UnsafePtr != nil {
		return struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
			V2 bool
		}{V0: gopurs_runtime.Value{}, V1: (*(*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V0, V2: true}
	}; return struct {
		V0 gopurs_runtime.Value
		V1 gopurs_runtime.Value
		V2 bool
	}{V0: (*(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: gopurs_runtime.Value{}, V2: false} }()
}

func Call_Data_Decide_Decide_dollar_Dict(x_0_loop struct {
	Divide0 gopurs_runtime.Value
	choose  gopurs_runtime.Value
	/*gopurs:record-v1:2:0044006900760069006400650030,00630068006f006f00730065*/
}) *Constructor_Data_Decide_Decide[gopurs_runtime.Value] {
	var x_0 struct {
		Divide0 gopurs_runtime.Value
		choose  gopurs_runtime.Value
		/*gopurs:record-v1:2:0044006900760069006400650030,00630068006f006f00730065*/
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Data_Decide_Decide[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict2("Divide0", "choose", orig.Divide0, orig.choose)
	}())
}

func Call_Data_Decide_chooseOp(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
	_ = dictSemigroup_0
	// TAST (Let): divideOp_1_0 shape=App(Var) bindingType=(ADT ["Data","Divide","Divide"] [(Func [(TypeVar b)] (TypeVar r$scope10))])
	divideOp_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Divide_Divide[gopurs_runtime.Value]](Call_Data_Divide_divideOp(dictSemigroup_0))
	_ = divideOp_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1618621146, UnsafePtr: unsafe.Pointer((&Constructor_Data_Decide_Decide[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 2642321722, UnsafePtr: unsafe.Pointer(divideOp_1_0)}
	}), gopurs_runtime.Func3(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value, v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Func(func(v2_5 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t1 gopurs_runtime.Value
			{
				if v2_5.Type == 9 && v2_5.IntVal == 3711209382 {
					__t1 = gopurs_runtime.Apply(v_3, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_5.UnsafePtr).V0)
					goto end_branch_1
				} else {

				}
			}
			{
				if v2_5.Type == 9 && v2_5.IntVal == 2465973597 {
					__t1 = gopurs_runtime.Apply(v1_4, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_5.UnsafePtr).V0)
					goto end_branch_1
				} else {

				}
			}
			{
				__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_1:
			return __t1
		}), f_2)
	})}))}
}

func Call_Data_Decide_choose(dict_0_loop *Constructor_Data_Decide_Decide[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Data_Decide_Decide[gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V1
}

func Call_Data_Decide_chosen(dictDecide_0_loop *Constructor_Data_Decide_Decide[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictDecide_0 *Constructor_Data_Decide_Decide[gopurs_runtime.Value] = dictDecide_0_loop
	_ = dictDecide_0
	return gopurs_runtime.Apply(dictDecide_0.V1, gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return x_1
	}))
}
