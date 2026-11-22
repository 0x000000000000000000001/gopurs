package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Foreign_Index_disj gopurs_runtime.Value
var once_Foreign_Index_disj sync.Once

func Get_Foreign_Index_disj() gopurs_runtime.Value {
	once_Foreign_Index_disj.Do(func() {
		cache_Foreign_Index_disj = Get_Data_HeytingAlgebra_disj__2421783737()
	})
	return cache_Foreign_Index_disj
}

var cache_Foreign_Index_eq gopurs_runtime.Value
var once_Foreign_Index_eq sync.Once

func Get_Foreign_Index_eq() gopurs_runtime.Value {
	once_Foreign_Index_eq.Do(func() {
		cache_Foreign_Index_eq = Get_Data_Eq_eq__3478410105()
	})
	return cache_Foreign_Index_eq
}

var cache_Foreign_Index_Index_dollar_Dict gopurs_runtime.Value
var once_Foreign_Index_Index_dollar_Dict sync.Once

func Get_Foreign_Index_Index_dollar_Dict() gopurs_runtime.Value {
	once_Foreign_Index_Index_dollar_Dict.Do(func() {
		cache_Foreign_Index_Index_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3155911316, UnsafePtr: unsafe.Pointer(Call_Foreign_Index_Index_dollar_Dict(func() struct {
				errorAt        gopurs_runtime.Value
				hasOwnProperty gopurs_runtime.Value
				hasProperty    gopurs_runtime.Value
				index          gopurs_runtime.Value
				/*gopurs:record-v1:4:006500720072006f007200410074,006800610073004f0077006e00500072006f00700065007200740079,00680061007300500072006f00700065007200740079,0069006e006400650078*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					errorAt        gopurs_runtime.Value
					hasOwnProperty gopurs_runtime.Value
					hasProperty    gopurs_runtime.Value
					index          gopurs_runtime.Value
					/*gopurs:record-v1:4:006500720072006f007200410074,006800610073004f0077006e00500072006f00700065007200740079,00680061007300500072006f00700065007200740079,0069006e006400650078*/
				}{}
				clone.errorAt = gopurs_runtime.RecordGet(orig, "errorAt")
				clone.hasOwnProperty = gopurs_runtime.RecordGet(orig, "hasOwnProperty")
				clone.hasProperty = gopurs_runtime.RecordGet(orig, "hasProperty")
				clone.index = gopurs_runtime.RecordGet(orig, "index")
				return clone
			}()))}
		})
	})
	return cache_Foreign_Index_Index_dollar_Dict
}

var cache_Foreign_Index_Indexable_dollar_Dict gopurs_runtime.Value
var once_Foreign_Index_Indexable_dollar_Dict sync.Once

func Get_Foreign_Index_Indexable_dollar_Dict() gopurs_runtime.Value {
	once_Foreign_Index_Indexable_dollar_Dict.Do(func() {
		cache_Foreign_Index_Indexable_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1103386302, UnsafePtr: unsafe.Pointer(Call_Foreign_Index_Indexable_dollar_Dict(func() struct {
				ix gopurs_runtime.Value
				/*gopurs:record-v1:1:00690078*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					ix gopurs_runtime.Value
					/*gopurs:record-v1:1:00690078*/
				}{}
				clone.ix = gopurs_runtime.RecordGet(orig, "ix")
				return clone
			}()))}
		})
	})
	return cache_Foreign_Index_Indexable_dollar_Dict
}

var cache_Foreign_Index_unsafeReadProp gopurs_runtime.Value
var once_Foreign_Index_unsafeReadProp sync.Once

func Get_Foreign_Index_unsafeReadProp() gopurs_runtime.Value {
	once_Foreign_Index_unsafeReadProp.Do(func() {
		cache_Foreign_Index_unsafeReadProp = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Foreign_Index_unsafeReadProp(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](dictMonad_0_box))
		})
	})
	return cache_Foreign_Index_unsafeReadProp
}

var cache_Foreign_Index_readProp gopurs_runtime.Value
var once_Foreign_Index_readProp sync.Once

func Get_Foreign_Index_readProp() gopurs_runtime.Value {
	once_Foreign_Index_readProp.Do(func() {
		cache_Foreign_Index_readProp = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Foreign_Index_readProp(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](dictMonad_0_box))
		})
	})
	return cache_Foreign_Index_readProp
}

var cache_Foreign_Index_readIndex gopurs_runtime.Value
var once_Foreign_Index_readIndex sync.Once

func Get_Foreign_Index_readIndex() gopurs_runtime.Value {
	once_Foreign_Index_readIndex.Do(func() {
		cache_Foreign_Index_readIndex = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Foreign_Index_readIndex(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](dictMonad_0_box))
		})
	})
	return cache_Foreign_Index_readIndex
}

var cache_Foreign_Index_ix gopurs_runtime.Value
var once_Foreign_Index_ix sync.Once

func Get_Foreign_Index_ix() gopurs_runtime.Value {
	once_Foreign_Index_ix.Do(func() {
		cache_Foreign_Index_ix = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Foreign_Index_ix(gopurs_runtime.CoerceToStruct[Constructor_Foreign_Index_Indexable[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Foreign_Index_ix
}

var cache_Foreign_Index_index gopurs_runtime.Value
var once_Foreign_Index_index sync.Once

func Get_Foreign_Index_index() gopurs_runtime.Value {
	once_Foreign_Index_index.Do(func() {
		cache_Foreign_Index_index = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Foreign_Index_index(gopurs_runtime.CoerceToStruct[Constructor_Foreign_Index_Index[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Foreign_Index_index
}

var cache_Foreign_Index_Indexable_dollar_Dict1 gopurs_runtime.Value
var once_Foreign_Index_Indexable_dollar_Dict1 sync.Once

func Get_Foreign_Index_Indexable_dollar_Dict1() gopurs_runtime.Value {
	once_Foreign_Index_Indexable_dollar_Dict1.Do(func() {
		cache_Foreign_Index_Indexable_dollar_Dict1 = gopurs_runtime.Value{Type: 9, IntVal: 1103386302, UnsafePtr: unsafe.Pointer((&Constructor_Foreign_Index_Indexable[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(dictIndex_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Foreign_Index_index(gopurs_runtime.CoerceToStruct[Constructor_Foreign_Index_Index[gopurs_runtime.Value, gopurs_runtime.Value]](dictIndex_0))
		})}))}
	})
	return cache_Foreign_Index_Indexable_dollar_Dict1
}

var cache_Foreign_Index_indexableExceptT gopurs_runtime.Value
var once_Foreign_Index_indexableExceptT sync.Once

func Get_Foreign_Index_indexableExceptT() gopurs_runtime.Value {
	once_Foreign_Index_indexableExceptT.Do(func() {
		cache_Foreign_Index_indexableExceptT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Foreign_Index_indexableExceptT(dictMonad_0_box)
		})
	})
	return cache_Foreign_Index_indexableExceptT
}

var cache_Foreign_Index_indexableForeign gopurs_runtime.Value
var once_Foreign_Index_indexableForeign sync.Once

func Get_Foreign_Index_indexableForeign() gopurs_runtime.Value {
	once_Foreign_Index_indexableForeign.Do(func() {
		cache_Foreign_Index_indexableForeign = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Foreign_Index_indexableForeign(dictMonad_0_box)
		})
	})
	return cache_Foreign_Index_indexableForeign
}

var cache_Foreign_Index_hasPropertyImpl gopurs_runtime.Value
var once_Foreign_Index_hasPropertyImpl sync.Once

func Get_Foreign_Index_hasPropertyImpl() gopurs_runtime.Value {
	once_Foreign_Index_hasPropertyImpl.Do(func() {
		cache_Foreign_Index_hasPropertyImpl = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(Call_Foreign_Index_hasPropertyImpl(v_0_box, v1_1_box))
		})
	})
	return cache_Foreign_Index_hasPropertyImpl
}

var cache_Foreign_Index_hasProperty gopurs_runtime.Value
var once_Foreign_Index_hasProperty sync.Once

func Get_Foreign_Index_hasProperty() gopurs_runtime.Value {
	once_Foreign_Index_hasProperty.Do(func() {
		cache_Foreign_Index_hasProperty = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Foreign_Index_hasProperty(gopurs_runtime.CoerceToStruct[Constructor_Foreign_Index_Index[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Foreign_Index_hasProperty
}

var cache_Foreign_Index_hasOwnPropertyImpl gopurs_runtime.Value
var once_Foreign_Index_hasOwnPropertyImpl sync.Once

func Get_Foreign_Index_hasOwnPropertyImpl() gopurs_runtime.Value {
	once_Foreign_Index_hasOwnPropertyImpl.Do(func() {
		cache_Foreign_Index_hasOwnPropertyImpl = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(Call_Foreign_Index_hasOwnPropertyImpl(v_0_box, v1_1_box))
		})
	})
	return cache_Foreign_Index_hasOwnPropertyImpl
}

var cache_Foreign_Index_indexInt gopurs_runtime.Value
var once_Foreign_Index_indexInt sync.Once

func Get_Foreign_Index_indexInt() gopurs_runtime.Value {
	once_Foreign_Index_indexInt.Do(func() {
		cache_Foreign_Index_indexInt = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Foreign_Index_indexInt(dictMonad_0_box)
		})
	})
	return cache_Foreign_Index_indexInt
}

var cache_Foreign_Index_indexString gopurs_runtime.Value
var once_Foreign_Index_indexString sync.Once

func Get_Foreign_Index_indexString() gopurs_runtime.Value {
	once_Foreign_Index_indexString.Do(func() {
		cache_Foreign_Index_indexString = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Foreign_Index_indexString(dictMonad_0_box)
		})
	})
	return cache_Foreign_Index_indexString
}

var cache_Foreign_Index_hasOwnProperty gopurs_runtime.Value
var once_Foreign_Index_hasOwnProperty sync.Once

func Get_Foreign_Index_hasOwnProperty() gopurs_runtime.Value {
	once_Foreign_Index_hasOwnProperty.Do(func() {
		cache_Foreign_Index_hasOwnProperty = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Foreign_Index_hasOwnProperty(gopurs_runtime.CoerceToStruct[Constructor_Foreign_Index_Index[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Foreign_Index_hasOwnProperty
}

var cache_Foreign_Index_errorAt gopurs_runtime.Value
var once_Foreign_Index_errorAt sync.Once

func Get_Foreign_Index_errorAt() gopurs_runtime.Value {
	once_Foreign_Index_errorAt.Do(func() {
		cache_Foreign_Index_errorAt = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Foreign_Index_errorAt(gopurs_runtime.CoerceToStruct[Constructor_Foreign_Index_Index[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Foreign_Index_errorAt
}

type Constructor_Foreign_Index_Index[T_i any, T_m any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
	V3 gopurs_runtime.Value
}

//gopurs:init-v1:00670065006e006500720061007400650064002f0046006f0072006500690067006e005f0049006e006400650078:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f0046006f0072006500690067006e005f0049006e006400650078005f0049006e006400650078
func init() {
	gopurs_runtime.StructGetters[3155911316] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Foreign_Index_Index[gopurs_runtime.Value, gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "errorAt":
			return gopurs_runtime.Box(c.V0)
		case "hasOwnProperty":
			return gopurs_runtime.Box(c.V1)
		case "hasProperty":
			return gopurs_runtime.Box(c.V2)
		case "index":
			return gopurs_runtime.Box(c.V3)
		default:
			panic("Key not found in dictionary Constructor_Foreign_Index_Index: " + key)
		}
	}
}

type Constructor_Foreign_Index_Indexable[T_a any, T_m any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

//gopurs:init-v1:00670065006e006500720061007400650064002f0046006f0072006500690067006e005f0049006e006400650078:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f0046006f0072006500690067006e005f0049006e006400650078005f0049006e00640065007800610062006c0065
func init() {
	gopurs_runtime.StructGetters[1103386302] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Foreign_Index_Indexable[gopurs_runtime.Value, gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "ix":
			return gopurs_runtime.Box(c.V0)
		default:
			panic("Key not found in dictionary Constructor_Foreign_Index_Indexable: " + key)
		}
	}
}

func Call_Foreign_Index_Index_dollar_Dict(x_0_loop struct {
	errorAt        gopurs_runtime.Value
	hasOwnProperty gopurs_runtime.Value
	hasProperty    gopurs_runtime.Value
	index          gopurs_runtime.Value
	/*gopurs:record-v1:4:006500720072006f007200410074,006800610073004f0077006e00500072006f00700065007200740079,00680061007300500072006f00700065007200740079,0069006e006400650078*/
}) *Constructor_Foreign_Index_Index[gopurs_runtime.Value, gopurs_runtime.Value] {
	var x_0 struct {
		errorAt        gopurs_runtime.Value
		hasOwnProperty gopurs_runtime.Value
		hasProperty    gopurs_runtime.Value
		index          gopurs_runtime.Value
		/*gopurs:record-v1:4:006500720072006f007200410074,006800610073004f0077006e00500072006f00700065007200740079,00680061007300500072006f00700065007200740079,0069006e006400650078*/
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Foreign_Index_Index[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict4("errorAt", "hasOwnProperty", "hasProperty", "index", orig.errorAt, orig.hasOwnProperty, orig.hasProperty, orig.index)
	}())
}

func Call_Foreign_Index_Indexable_dollar_Dict(x_0_loop struct {
	ix gopurs_runtime.Value
	/*gopurs:record-v1:1:00690078*/
}) *Constructor_Foreign_Index_Indexable[gopurs_runtime.Value, gopurs_runtime.Value] {
	var x_0 struct {
		ix gopurs_runtime.Value
		/*gopurs:record-v1:1:00690078*/
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Foreign_Index_Indexable[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict1("ix", orig.ix)
	}())
}

func Call_Foreign_Index_unsafeReadProp(dictMonad_0_loop *Constructor_Control_Monad_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictMonad_0 *Constructor_Control_Monad_Monad[gopurs_runtime.Value] = dictMonad_0_loop
	_ = dictMonad_0
	// TAST (Let): fail__2179776687_1_0 shape=App(Var) bindingType=(Func [(ADT ["Foreign","ForeignError"] [])] (TypeApp (TypeVar m$scope3) [(ADT ["Data","Either","Either"] [(ADT ["Data","NonEmpty","NonEmpty"] [(ADT ["Data","List","Types","List"] []), (ADT ["Foreign","ForeignError"] [])]), (ADT ["Foreign","Foreign"] [])])]))
	fail__2179776687_1_0 := gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Call_Control_Monad_Error_Class_throwError(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Error_Class_MonadThrow[gopurs_runtime.Value, gopurs_runtime.Value]](Call_Control_Monad_Except_Trans_monadThrowExceptT(gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(dictMonad_0)}))), Get_Data_List_NonEmpty_singleton())
	_ = fail__2179776687_1_0
	// TAST (Let): pure_2_1 shape=App(Var) bindingType=(Func [(ADT ["Foreign","Foreign"] [])] (TypeApp (TypeVar m$scope3) [(ADT ["Data","Either","Either"] [(ADT ["Data","NonEmpty","NonEmpty"] [(ADT ["Data","List","Types","List"] []), (ADT ["Foreign","ForeignError"] [])]), (ADT ["Foreign","Foreign"] [])])]))
	pure_2_1 := Call_Control_Applicative_pure(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Call_Control_Monad_Except_Trans_applicativeExceptT(gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(dictMonad_0)})))
	_ = pure_2_1
	return gopurs_runtime.Func2(func(k_3 gopurs_runtime.Value, value_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.UncurriedApp4(Get_Foreign_Index_unsafeReadPropImpl(), gopurs_runtime.Apply(fail__2179776687_1_0, gopurs_runtime.Value{Type: 9, IntVal: 2536121783, UnsafePtr: unsafe.Pointer((&Constructor_Foreign_TypeMismatch{1, "object", gopurs_runtime.Apply(Get_Foreign_typeOf(), value_4).StrVal()}))}), pure_2_1, k_3, value_4)
	})
}

func Call_Foreign_Index_readProp(dictMonad_0_loop *Constructor_Control_Monad_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictMonad_0 *Constructor_Control_Monad_Monad[gopurs_runtime.Value] = dictMonad_0_loop
	_ = dictMonad_0
	return Call_Foreign_Index_unsafeReadProp(dictMonad_0)
}

func Call_Foreign_Index_readIndex(dictMonad_0_loop *Constructor_Control_Monad_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictMonad_0 *Constructor_Control_Monad_Monad[gopurs_runtime.Value] = dictMonad_0_loop
	_ = dictMonad_0
	return Call_Foreign_Index_unsafeReadProp(dictMonad_0)
}

func Call_Foreign_Index_ix(dict_0_loop *Constructor_Foreign_Index_Indexable[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Foreign_Index_Indexable[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V0
}

func Call_Foreign_Index_index(dict_0_loop *Constructor_Foreign_Index_Index[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Foreign_Index_Index[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V3
}

func Call_Foreign_Index_indexableExceptT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
	_ = dictMonad_0
	// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=Any
	__local_var_1_0 := Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Call_Control_Monad_Except_Trans_bindExceptT(dictMonad_0)))
	_ = __local_var_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1103386302, UnsafePtr: unsafe.Pointer((&Constructor_Foreign_Index_Indexable[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(dictIndex_2 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): index1_3_1 shape=App(Var) bindingType=(Func [(ADT ["Foreign","Foreign"] []), (TypeVar i$scope25)] (TypeApp (TypeVar m$scope21) [(ADT ["Data","Either","Either"] [(ADT ["Data","NonEmpty","NonEmpty"] [(ADT ["Data","List","Types","List"] []), (ADT ["Foreign","ForeignError"] [])]), (ADT ["Foreign","Foreign"] [])])]))
		index1_3_1 := Call_Foreign_Index_index(gopurs_runtime.CoerceToStruct[Constructor_Foreign_Index_Index[gopurs_runtime.Value, gopurs_runtime.Value]](dictIndex_2))
		_ = index1_3_1
		return gopurs_runtime.Func2(func(f_4 gopurs_runtime.Value, i_5 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(__local_var_1_0, f_4, gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(index1_3_1, a_6, i_5)
			}))
		})
	})}))}
}

func Call_Foreign_Index_indexableForeign(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
	_ = dictMonad_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1103386302, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Foreign_Index_Indexable[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Foreign_Index_Indexable_dollar_Dict1()))}
}

func Call_Foreign_Index_hasPropertyImpl(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value) bool {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	var v1_1 gopurs_runtime.Value = v1_1_loop
	_ = v1_1
	var __t0 bool
	{
		if (gopurs_runtime.Apply(Get_Foreign_isNull(), v1_1).IntVal) != (0) {
			__t0 = false
			goto end_branch_0
		} else {

		}
	}
	{
		if (gopurs_runtime.Apply(Get_Foreign_isUndefined(), v1_1).IntVal) != (0) {
			__t0 = false
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = (((gopurs_runtime.Apply(Get_Foreign_typeOf(), v1_1).StrVal()) == ("object")) || ((gopurs_runtime.Apply(Get_Foreign_typeOf(), v1_1).StrVal()) == ("function"))) && ((gopurs_runtime.UncurriedApp2(Get_Foreign_Index_unsafeHasProperty(), v_0, v1_1).IntVal) != (0))
	}
end_branch_0:
	return __t0
}

func Call_Foreign_Index_hasProperty(dict_0_loop *Constructor_Foreign_Index_Index[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Foreign_Index_Index[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V2
}

func Call_Foreign_Index_hasOwnPropertyImpl(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value) bool {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	var v1_1 gopurs_runtime.Value = v1_1_loop
	_ = v1_1
	var __t0 bool
	{
		if (gopurs_runtime.Apply(Get_Foreign_isNull(), v1_1).IntVal) != (0) {
			__t0 = false
			goto end_branch_0
		} else {

		}
	}
	{
		if (gopurs_runtime.Apply(Get_Foreign_isUndefined(), v1_1).IntVal) != (0) {
			__t0 = false
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = (((gopurs_runtime.Apply(Get_Foreign_typeOf(), v1_1).StrVal()) == ("object")) || ((gopurs_runtime.Apply(Get_Foreign_typeOf(), v1_1).StrVal()) == ("function"))) && ((gopurs_runtime.UncurriedApp2(Get_Foreign_Index_unsafeHasOwnProperty(), v_0, v1_1).IntVal) != (0))
	}
end_branch_0:
	return __t0
}

func Call_Foreign_Index_indexInt(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
	_ = dictMonad_0
	// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=Any
	__local_var_1_0 := Call_Foreign_Index_readIndex(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](dictMonad_0))
	_ = __local_var_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 3155911316, UnsafePtr: unsafe.Pointer(Rebox_Foreign_Index_1827029757_375302278((&Constructor_Foreign_Index_Index[int64, gopurs_runtime.Value]{1, Get_Foreign_ErrorAtIndex(), Get_Foreign_Index_hasOwnPropertyImpl(), Get_Foreign_Index_hasPropertyImpl(), gopurs_runtime.Func2(func(b_2 gopurs_runtime.Value, a_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(__local_var_1_0, gopurs_runtime.Int(a_3.IntVal), b_2)
	})})))}
}

func Call_Foreign_Index_indexString(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
	_ = dictMonad_0
	// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=Any
	__local_var_1_0 := Call_Foreign_Index_readProp(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](dictMonad_0))
	_ = __local_var_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 3155911316, UnsafePtr: unsafe.Pointer(Rebox_Foreign_Index_3744877529_375302278((&Constructor_Foreign_Index_Index[string, gopurs_runtime.Value]{1, Get_Foreign_ErrorAtProperty(), Get_Foreign_Index_hasOwnPropertyImpl(), Get_Foreign_Index_hasPropertyImpl(), gopurs_runtime.Func2(func(b_2 gopurs_runtime.Value, a_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(__local_var_1_0, gopurs_runtime.NormalizeStringValue(a_3), b_2)
	})})))}
}

func Call_Foreign_Index_hasOwnProperty(dict_0_loop *Constructor_Foreign_Index_Index[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Foreign_Index_Index[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V1
}

func Call_Foreign_Index_errorAt(dict_0_loop *Constructor_Foreign_Index_Index[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Foreign_Index_Index[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V0
}

func Rebox_Foreign_Index_1827029757_375302278(in *Constructor_Foreign_Index_Index[int64, gopurs_runtime.Value]) *Constructor_Foreign_Index_Index[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Foreign_Index_Index[gopurs_runtime.Value, gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Foreign_Index_3744877529_375302278(in *Constructor_Foreign_Index_Index[string, gopurs_runtime.Value]) *Constructor_Foreign_Index_Index[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Foreign_Index_Index[gopurs_runtime.Value, gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Get_Foreign_Index_unsafeHasOwnProperty() gopurs_runtime.Value {
	return _Gopurs_Foreign_Index_UnsafeHasOwnProperty
}

func Get_Foreign_Index_unsafeHasProperty() gopurs_runtime.Value {
	return _Gopurs_Foreign_Index_UnsafeHasProperty
}

func Get_Foreign_Index_unsafeReadPropImpl() gopurs_runtime.Value {
	return _Gopurs_Foreign_Index_UnsafeReadPropImpl
}
