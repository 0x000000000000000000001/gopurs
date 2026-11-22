package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Profunctor_Choice_identity gopurs_runtime.Value
var once_Data_Profunctor_Choice_identity sync.Once

func Get_Data_Profunctor_Choice_identity() gopurs_runtime.Value {
	once_Data_Profunctor_Choice_identity.Do(func() {
		cache_Data_Profunctor_Choice_identity = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Profunctor_Choice_identity(x_0_box)
		})
	})
	return cache_Data_Profunctor_Choice_identity
}

var cache_Data_Profunctor_Choice_Choice_dollar_Dict gopurs_runtime.Value
var once_Data_Profunctor_Choice_Choice_dollar_Dict sync.Once

func Get_Data_Profunctor_Choice_Choice_dollar_Dict() gopurs_runtime.Value {
	once_Data_Profunctor_Choice_Choice_dollar_Dict.Do(func() {
		cache_Data_Profunctor_Choice_Choice_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3666633887, UnsafePtr: unsafe.Pointer(Call_Data_Profunctor_Choice_Choice_dollar_Dict(func() struct {
				Profunctor0 gopurs_runtime.Value
				left        gopurs_runtime.Value
				right       gopurs_runtime.Value
				/*gopurs:record-v1:3:00500072006f00660075006e00630074006f00720030,006c006500660074,00720069006700680074*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Profunctor0 gopurs_runtime.Value
					left        gopurs_runtime.Value
					right       gopurs_runtime.Value
					/*gopurs:record-v1:3:00500072006f00660075006e00630074006f00720030,006c006500660074,00720069006700680074*/
				}{}
				clone.Profunctor0 = gopurs_runtime.RecordGet(orig, "Profunctor0")
				clone.left = gopurs_runtime.RecordGet(orig, "left")
				clone.right = gopurs_runtime.RecordGet(orig, "right")
				return clone
			}()))}
		})
	})
	return cache_Data_Profunctor_Choice_Choice_dollar_Dict
}

var cache_Data_Profunctor_Choice_right gopurs_runtime.Value
var once_Data_Profunctor_Choice_right sync.Once

func Get_Data_Profunctor_Choice_right() gopurs_runtime.Value {
	once_Data_Profunctor_Choice_right.Do(func() {
		cache_Data_Profunctor_Choice_right = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Profunctor_Choice_right(gopurs_runtime.CoerceToStruct[Constructor_Data_Profunctor_Choice_Choice[gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Data_Profunctor_Choice_right
}

var cache_Data_Profunctor_Choice_left gopurs_runtime.Value
var once_Data_Profunctor_Choice_left sync.Once

func Get_Data_Profunctor_Choice_left() gopurs_runtime.Value {
	once_Data_Profunctor_Choice_left.Do(func() {
		cache_Data_Profunctor_Choice_left = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Profunctor_Choice_left(gopurs_runtime.CoerceToStruct[Constructor_Data_Profunctor_Choice_Choice[gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Data_Profunctor_Choice_left
}

var cache_Data_Profunctor_Choice_splitChoice gopurs_runtime.Value
var once_Data_Profunctor_Choice_splitChoice sync.Once

func Get_Data_Profunctor_Choice_splitChoice() gopurs_runtime.Value {
	once_Data_Profunctor_Choice_splitChoice.Do(func() {
		cache_Data_Profunctor_Choice_splitChoice = gopurs_runtime.Func(func(dictSemigroupoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Profunctor_Choice_splitChoice(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](dictSemigroupoid_0_box))
		})
	})
	return cache_Data_Profunctor_Choice_splitChoice
}

var cache_Data_Profunctor_Choice_fanin gopurs_runtime.Value
var once_Data_Profunctor_Choice_fanin sync.Once

func Get_Data_Profunctor_Choice_fanin() gopurs_runtime.Value {
	once_Data_Profunctor_Choice_fanin.Do(func() {
		cache_Data_Profunctor_Choice_fanin = gopurs_runtime.Func(func(dictSemigroupoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Profunctor_Choice_fanin(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](dictSemigroupoid_0_box))
		})
	})
	return cache_Data_Profunctor_Choice_fanin
}

var cache_Data_Profunctor_Choice_choiceFn gopurs_runtime.Value
var once_Data_Profunctor_Choice_choiceFn sync.Once

func Get_Data_Profunctor_Choice_choiceFn() gopurs_runtime.Value {
	once_Data_Profunctor_Choice_choiceFn.Do(func() {
		cache_Data_Profunctor_Choice_choiceFn = gopurs_runtime.Value{Type: 9, IntVal: 3666633887, UnsafePtr: unsafe.Pointer((&Constructor_Data_Profunctor_Choice_Choice[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2367018778, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Profunctor_Profunctor[gopurs_runtime.Value]](Get_Data_Profunctor_profunctorFn()))}
		}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t0 gopurs_runtime.Value
			{
				if v1_1.Type == 9 && v1_1.IntVal == 3711209382 {
					__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer((&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(v_0, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V0)}))}
					goto end_branch_0
				} else {

				}
			}
			{
				if v1_1.Type == 9 && v1_1.IntVal == 2465973597 {
					__t0 = func() gopurs_runtime.Value {
						_v := struct {
							V0 gopurs_runtime.Value
							V1 gopurs_runtime.Value
							V2 bool
						}{gopurs_runtime.Value{}, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V0, true}
						if _v.V2 {
							return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
						}
						return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
					}()
					goto end_branch_0
				} else {

				}
			}
			{
				__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_0:
			return __t0
		}), gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map")}))}
	})
	return cache_Data_Profunctor_Choice_choiceFn
}

type Constructor_Data_Profunctor_Choice_Choice[T_p any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
}

//gopurs:init-v1:00670065006e006500720061007400650064002f0044006100740061005f00500072006f00660075006e00630074006f0072005f00430068006f006900630065:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f0044006100740061005f00500072006f00660075006e00630074006f0072005f00430068006f006900630065005f00430068006f006900630065
func init() {
	gopurs_runtime.StructGetters[3666633887] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Profunctor_Choice_Choice[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "Profunctor0":
			return gopurs_runtime.Box(c.V0)
		case "left":
			return gopurs_runtime.Box(c.V1)
		case "right":
			return gopurs_runtime.Box(c.V2)
		default:
			panic("Key not found in dictionary Constructor_Data_Profunctor_Choice_Choice: " + key)
		}
	}
}

func Call_Data_Profunctor_Choice_identity(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Data_Profunctor_Choice_Choice_dollar_Dict(x_0_loop struct {
	Profunctor0 gopurs_runtime.Value
	left        gopurs_runtime.Value
	right       gopurs_runtime.Value
	/*gopurs:record-v1:3:00500072006f00660075006e00630074006f00720030,006c006500660074,00720069006700680074*/
}) *Constructor_Data_Profunctor_Choice_Choice[gopurs_runtime.Value] {
	var x_0 struct {
		Profunctor0 gopurs_runtime.Value
		left        gopurs_runtime.Value
		right       gopurs_runtime.Value
		/*gopurs:record-v1:3:00500072006f00660075006e00630074006f00720030,006c006500660074,00720069006700680074*/
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Data_Profunctor_Choice_Choice[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict3("Profunctor0", "left", "right", orig.Profunctor0, orig.left, orig.right)
	}())
}

func Call_Data_Profunctor_Choice_right(dict_0_loop *Constructor_Data_Profunctor_Choice_Choice[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Data_Profunctor_Choice_Choice[gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V2
}

func Call_Data_Profunctor_Choice_left(dict_0_loop *Constructor_Data_Profunctor_Choice_Choice[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Data_Profunctor_Choice_Choice[gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V1
}

func Call_Data_Profunctor_Choice_splitChoice(dictSemigroupoid_0_loop *Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictSemigroupoid_0 *Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value] = dictSemigroupoid_0_loop
	_ = dictSemigroupoid_0
	// TAST (Let): composeFlipped_1_0 shape=App(Var) bindingType=(Func [(TypeApp (TypeVar p$scope13) [(ADT ["Data","Either","Either"] [(TypeVar a$scope14), (TypeVar c$scope16)]), (ADT ["Data","Either","Either"] [(TypeVar b$scope15), (TypeVar c$scope16)])]), (TypeApp (TypeVar p$scope13) [(ADT ["Data","Either","Either"] [(TypeVar b$scope15), (TypeVar c$scope16)]), (ADT ["Data","Either","Either"] [(TypeVar b$scope15), (TypeVar d$scope17)])])] (TypeApp (TypeVar p$scope13) [(ADT ["Data","Either","Either"] [(TypeVar a$scope14), (TypeVar c$scope16)]), (ADT ["Data","Either","Either"] [(TypeVar b$scope15), (TypeVar d$scope17)])]))
	composeFlipped_1_0 := Call_Control_Semigroupoid_composeFlipped(dictSemigroupoid_0)
	_ = composeFlipped_1_0
	return gopurs_runtime.Func3(func(dictChoice_2 gopurs_runtime.Value, l_3 gopurs_runtime.Value, r_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(composeFlipped_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictChoice_2, "left"), l_3), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictChoice_2, "right"), r_4))
	})
}

func Call_Data_Profunctor_Choice_fanin(dictSemigroupoid_0_loop *Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictSemigroupoid_0 *Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value] = dictSemigroupoid_0_loop
	_ = dictSemigroupoid_0
	// TAST (Let): splitChoice1__614038976_1_0 shape=App(Var) bindingType=(ConstrainedType (Func [(TypeApp (TypeVar p$scope23) [(TypeVar a$scope24), (TypeVar c$scope26)]), (TypeApp (TypeVar p$scope23) [(TypeVar b$scope25), (TypeVar c$scope26)])] (TypeApp (TypeVar p$scope23) [(ADT ["Data","Either","Either"] [(TypeVar a$scope24), (TypeVar b$scope25)]), (ADT ["Data","Either","Either"] [(TypeVar c$scope26), (TypeVar c$scope26)])])))
	splitChoice1__614038976_1_0 := Call_Data_Profunctor_Choice_splitChoice(dictSemigroupoid_0)
	_ = splitChoice1__614038976_1_0
	return gopurs_runtime.Func(func(dictChoice_2 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_3_1 shape=App(Other) bindingType=Any
		__local_var_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictChoice_2, "Profunctor0"), gopurs_runtime.Value{})
		_ = __local_var_3_1
		// TAST (Let): splitChoice2_4_2 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar p$scope23) [(TypeVar a$scope24), (TypeVar c$scope26)]), (TypeApp (TypeVar p$scope23) [(TypeVar b$scope25), (TypeVar c$scope26)])] (TypeApp (TypeVar p$scope23) [(ADT ["Data","Either","Either"] [(TypeVar a$scope24), (TypeVar b$scope25)]), (ADT ["Data","Either","Either"] [(TypeVar c$scope26), (TypeVar c$scope26)])]))
		splitChoice2_4_2 := gopurs_runtime.Apply(splitChoice1__614038976_1_0, gopurs_runtime.Value{Type: 9, IntVal: 3666633887, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Profunctor_Choice_Choice[gopurs_runtime.Value]](dictChoice_2))})
		_ = splitChoice2_4_2
		return gopurs_runtime.Func2(func(l_5 gopurs_runtime.Value, r_6 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_3_1, "dimap"), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
				return x_7
			}), gopurs_runtime.Func(func(v2_7 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t3 gopurs_runtime.Value
				{
					if v2_7.Type == 9 && v2_7.IntVal == 3711209382 {
						__t3 = (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_7.UnsafePtr).V0
						goto end_branch_3
					} else {

					}
				}
				{
					if v2_7.Type == 9 && v2_7.IntVal == 2465973597 {
						__t3 = (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_7.UnsafePtr).V0
						goto end_branch_3
					} else {

					}
				}
				{
					__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
				}
			end_branch_3:
				return __t3
			}), gopurs_runtime.Apply2(splitChoice2_4_2, l_5, r_6))
		})
	})
}
