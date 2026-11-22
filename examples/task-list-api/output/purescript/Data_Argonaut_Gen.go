package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Argonaut_Gen_negate gopurs_runtime.Value
var once_Data_Argonaut_Gen_negate sync.Once

func Get_Data_Argonaut_Gen_negate() gopurs_runtime.Value {
	once_Data_Argonaut_Gen_negate.Do(func() {
		cache_Data_Argonaut_Gen_negate = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Float(Call_Data_Argonaut_Gen_negate(a_0_box.FloatVal()))
		})
	})
	return cache_Data_Argonaut_Gen_negate
}

var cache_Data_Argonaut_Gen_foldable1NonEmpty gopurs_runtime.Value
var once_Data_Argonaut_Gen_foldable1NonEmpty sync.Once

func Get_Data_Argonaut_Gen_foldable1NonEmpty() gopurs_runtime.Value {
	once_Data_Argonaut_Gen_foldable1NonEmpty.Do(func() {
		cache_Data_Argonaut_Gen_foldable1NonEmpty = Call_Data_NonEmpty_foldable1NonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))})
	})
	return cache_Data_Argonaut_Gen_foldable1NonEmpty
}

var cache_Data_Argonaut_Gen_compose gopurs_runtime.Value
var once_Data_Argonaut_Gen_compose sync.Once

func Get_Data_Argonaut_Gen_compose() gopurs_runtime.Value {
	once_Data_Argonaut_Gen_compose.Do(func() {
		cache_Data_Argonaut_Gen_compose = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Argonaut_Gen_compose
}

var cache_Data_Argonaut_Gen_greaterThan gopurs_runtime.Value
var once_Data_Argonaut_Gen_greaterThan sync.Once

func Get_Data_Argonaut_Gen_greaterThan() gopurs_runtime.Value {
	once_Data_Argonaut_Gen_greaterThan.Do(func() {
		cache_Data_Argonaut_Gen_greaterThan = Get_Data_Ord_greaterThan__3348776441()
	})
	return cache_Data_Argonaut_Gen_greaterThan
}

var cache_Data_Argonaut_Gen_sub gopurs_runtime.Value
var once_Data_Argonaut_Gen_sub sync.Once

func Get_Data_Argonaut_Gen_sub() gopurs_runtime.Value {
	once_Data_Argonaut_Gen_sub.Do(func() {
		cache_Data_Argonaut_Gen_sub = Get_Data_Ring_sub__4157804846()
	})
	return cache_Data_Argonaut_Gen_sub
}

var cache_Data_Argonaut_Gen_min gopurs_runtime.Value
var once_Data_Argonaut_Gen_min sync.Once

func Get_Data_Argonaut_Gen_min() gopurs_runtime.Value {
	once_Data_Argonaut_Gen_min.Do(func() {
		cache_Data_Argonaut_Gen_min = Get_Data_Ord_min__4009478382()
	})
	return cache_Data_Argonaut_Gen_min
}

var cache_Data_Argonaut_Gen_genJson gopurs_runtime.Value
var once_Data_Argonaut_Gen_genJson sync.Once

func Get_Data_Argonaut_Gen_genJson() gopurs_runtime.Value {
	once_Data_Argonaut_Gen_genJson.Do(func() {
		cache_Data_Argonaut_Gen_genJson = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Argonaut_Gen_genJson(dictMonadGen_0_box)
		})
	})
	return cache_Data_Argonaut_Gen_genJson
}

func Call_Data_Argonaut_Gen_negate(a_0_loop float64) float64 {
	var a_0 float64 = a_0_loop
	_ = a_0
	return -(a_0)
}

func Call_Data_Argonaut_Gen_genJson(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
genJson:
	for {
		if false {
			continue genJson
		}
		var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
		_ = dictMonadGen_0
		// TAST (Let): Monad0_1_0 shape=App(Other) bindingType=Any
		Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{})
		_ = Monad0_1_0
		// TAST (Let): Functor0_2_1 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m$scope1)])
		Functor0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
		_ = Functor0_2_1
		// TAST (Let): chooseBool_3_2 shape=App(Var) bindingType=(TypeApp (TypeVar m$scope1) [Boolean])
		chooseBool_3_2 := Call_Control_Monad_Gen_Class_chooseBool(dictMonadGen_0)
		_ = chooseBool_3_2
		// TAST (Let): oneOf_4_3 shape=App(Var) bindingType=(Func [(ADT ["Data","NonEmpty","NonEmpty"] [(ADT ["Prim","Array"] []), (TypeApp (TypeVar m$scope1) [(ADT ["Data","Argonaut","Core","Json"] [])])])] (TypeApp (TypeVar m$scope1) [(ADT ["Data","Argonaut","Core","Json"] [])]))
		oneOf_4_3 := Call_Control_Monad_Gen_oneOf(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]](dictMonadGen_0), gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]](Get_Data_Argonaut_Gen_foldable1NonEmpty()))
		_ = oneOf_4_3
		// TAST (Let): Bind1_5_4 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m$scope1)])
		Bind1_5_4 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
		_ = Bind1_5_4
		// TAST (Let): pure1_6_5 shape=App(Var) bindingType=(Func [(ADT ["Data","Argonaut","Core","Json"] [])] (TypeApp (TypeVar m$scope1) [(ADT ["Data","Argonaut","Core","Json"] [])]))
		pure1_6_5 := Call_Control_Applicative_pure(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{})))
		_ = pure1_6_5
		// TAST (Let): __local_var_7_6 shape=App(Var) bindingType=Any
		__local_var_7_6 := Call_Control_Bind_bind(Bind1_5_4)
		_ = __local_var_7_6
		// TAST (Let): foldM_8_7 shape=App(Var) bindingType=(Func [(Func [(ADT ["Data","Argonaut","Core","Json"] []), String] (TypeApp (TypeVar m$scope1) [(ADT ["Data","Argonaut","Core","Json"] [])])), (ADT ["Data","Argonaut","Core","Json"] []), (Array String)] (TypeApp (TypeVar m$scope1) [(ADT ["Data","Argonaut","Core","Json"] [])]))
		foldM_8_7 := gopurs_runtime.Apply(Get_Data_Array_foldM(), gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{})))})
		_ = foldM_8_7
		// TAST (Let): choose_9_8 shape=App(Var) bindingType=(Func [(TypeApp (TypeVar m$scope1) [(ADT ["Data","Argonaut","Core","Json"] [])]), (TypeApp (TypeVar m$scope1) [(ADT ["Data","Argonaut","Core","Json"] [])])] (TypeApp (TypeVar m$scope1) [(ADT ["Data","Argonaut","Core","Json"] [])]))
		choose_9_8 := Call_Control_Monad_Gen_choose(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]](dictMonadGen_0))
		_ = choose_9_8
		return gopurs_runtime.Func(func(dictMonadRec_10 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): genUnicodeString_11_9 shape=App(Var) bindingType=(TypeApp (TypeVar m$scope1) [String])
			genUnicodeString_11_9 := Call_Data_String_Gen_genUnicodeString(dictMonadRec_10, dictMonadGen_0)
			_ = genUnicodeString_11_9
			// TAST (Let): unfoldable_12_10 shape=App(Var) bindingType=(Func [(TypeApp (TypeVar m$scope1) [(ADT ["Data","Argonaut","Core","Json"] [])])] (TypeApp (TypeVar m$scope1) [(Array (ADT ["Data","Argonaut","Core","Json"] []))]))
			unfoldable_12_10 := gopurs_runtime.Apply(Call_Control_Monad_Gen_unfoldable(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]](dictMonadRec_10), gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]](dictMonadGen_0)), gopurs_runtime.Value{Type: 9, IntVal: 2670894170, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]](Get_Data_Unfoldable_unfoldableArray()))})
			_ = unfoldable_12_10
			// TAST (Let): unfoldable1_13_11 shape=App(Var) bindingType=(Func [(TypeApp (TypeVar m$scope1) [String])] (TypeApp (TypeVar m$scope1) [(Array String)]))
			unfoldable1_13_11 := gopurs_runtime.Apply(Call_Control_Monad_Gen_unfoldable(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]](dictMonadRec_10), gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]](dictMonadGen_0)), gopurs_runtime.Value{Type: 9, IntVal: 2670894170, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]](Get_Data_Unfoldable_unfoldableArray()))})
			_ = unfoldable1_13_11
			return gopurs_runtime.Func(func(dictLazy_14 gopurs_runtime.Value) gopurs_runtime.Value {
				// TAST (Let): genLeaf_15_12 shape=App(Other) bindingType=(TypeApp (TypeVar m$scope1) [(ADT ["Data","Argonaut","Core","Json"] [])])
				genLeaf_15_12 := gopurs_runtime.Apply(oneOf_4_3, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer((&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), Get_Data_Argonaut_Core_jsonNull()), gopurs_runtime.Array([]gopurs_runtime.Value{gopurs_runtime.Apply2(Functor0_2_1.V0, Get_Data_Argonaut_Core_fromBoolean(), chooseBool_3_2), gopurs_runtime.Apply2(Functor0_2_1.V0, Get_Data_Argonaut_Core_fromNumber(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "chooseFloat"), gopurs_runtime.Float(-1000000.0), gopurs_runtime.Float(1000000.0))), gopurs_runtime.Apply2(Functor0_2_1.V0, Get_Data_Argonaut_Core_fromString(), genUnicodeString_11_9)})}))})
				_ = genLeaf_15_12
				// TAST (Let): genJArray_16_13 shape=App(Other) bindingType=(TypeApp (TypeVar m$scope1) [(ADT ["Data","Argonaut","Core","Json"] [])])
				genJArray_16_13 := gopurs_runtime.Apply2(Functor0_2_1.V0, Get_Data_Argonaut_Core_fromArray(), gopurs_runtime.Apply(unfoldable_12_10, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictLazy_14, "defer"), gopurs_runtime.Func(func(v_16 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply2(Call_Data_Argonaut_Gen_genJson(gopurs_runtime.Value{Type: 9, IntVal: 2254593219, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]](dictMonadGen_0))}), gopurs_runtime.Value{Type: 9, IntVal: 3709389635, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]](dictMonadRec_10))}, gopurs_runtime.Value{Type: 9, IntVal: 1860244333, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Lazy_Lazy[gopurs_runtime.Value]](dictLazy_14))})
				}))))
				_ = genJArray_16_13
				// TAST (Let): genJObject_17_14 shape=App(Other) bindingType=(TypeApp (TypeVar m$scope1) [(ADT ["Data","Argonaut","Core","Json"] [])])
				genJObject_17_14 := gopurs_runtime.Apply2(__local_var_7_6, gopurs_runtime.Apply(unfoldable1_13_11, genUnicodeString_11_9), gopurs_runtime.Apply2(foldM_8_7, gopurs_runtime.Func2(func(obj_17 gopurs_runtime.Value, k_18 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply2(Bind1_5_4.V1, gopurs_runtime.Apply2(Call_Data_Argonaut_Gen_genJson(dictMonadGen_0), dictMonadRec_10, dictLazy_14), gopurs_runtime.Func(func(v_19 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Apply(pure1_6_5, Data_Argonaut_Core__CaseJsonObject(Data_Argonaut_Core_FromObject_nativeWorker(Call_Foreign_Object_singleton__2505508685(k_18.StrVal(), v_19)), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Argonaut_Core_fromObject(), Call_Foreign_Object_insert__3487692426(k_18.StrVal(), v_19)), obj_17))
					}))
				}), Get_Data_Argonaut_Core_jsonEmptyObject()))
				_ = genJObject_17_14
				return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "resize"), gopurs_runtime.Apply(Get_Data_Ord_min__4009478382(), gopurs_runtime.Int(int64(5))), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "sized"), gopurs_runtime.Func(func(size_18 gopurs_runtime.Value) gopurs_runtime.Value {
					var __t15 gopurs_runtime.Value
					{
						if (size_18.IntVal) > (int64(1)) {
							__t15 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "resize"), gopurs_runtime.Func(func(v_19 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Int(gopurs_runtime.IntSub(v_19.IntVal, int64(1)))
							}), gopurs_runtime.Apply2(choose_9_8, genJArray_16_13, genJObject_17_14))
							goto end_branch_15
						} else {

						}
					}
					{
						__t15 = genLeaf_15_12
					}
				end_branch_15:
					return __t15
				})))
			})
		})
	}
}
