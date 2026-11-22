package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Map_Gen_genMap gopurs_runtime.Value
var once_Data_Map_Gen_genMap sync.Once

func Get_Data_Map_Gen_genMap() gopurs_runtime.Value {
	once_Data_Map_Gen_genMap.Do(func() {
		cache_Data_Map_Gen_genMap = gopurs_runtime.Func2(func(dictMonadRec_0_box gopurs_runtime.Value, dictMonadGen_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Map_Gen_genMap(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]](dictMonadRec_0_box), gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]](dictMonadGen_1_box))
		})
	})
	return cache_Data_Map_Gen_genMap
}

func Call_Data_Map_Gen_genMap(dictMonadRec_0_loop *Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value], dictMonadGen_1_loop *Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictMonadRec_0 *Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value] = dictMonadRec_0_loop
	_ = dictMonadRec_0
	var dictMonadGen_1 *Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value] = dictMonadGen_1_loop
	_ = dictMonadGen_1
	// TAST (Let): Monad0_2_0 shape=App(Other) bindingType=Any
	Monad0_2_0 := gopurs_runtime.Apply(dictMonadGen_1.V0, gopurs_runtime.Value{})
	_ = Monad0_2_0
	// TAST (Let): Bind1_3_1 shape=App(Other) bindingType=Any
	Bind1_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_0, "Bind1"), gopurs_runtime.Value{})
	_ = Bind1_3_1
	// TAST (Let): Functor0_4_2 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m$scope3)])
	Functor0_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Bind1_3_1, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
	_ = Functor0_4_2
	// TAST (Let): unfoldable1_5_3 shape=App(Var) bindingType=(Func [(TypeApp (TypeVar m$scope3) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a$scope4), (TypeVar b$scope5)])])] (TypeApp (TypeVar m$scope3) [(ADT ["Data","List","Types","List"] [(ADT ["Data","Tuple","Tuple"] [(TypeVar a$scope4), (TypeVar b$scope5)])])]))
	unfoldable1_5_3 := gopurs_runtime.Apply(Call_Control_Monad_Gen_unfoldable(dictMonadRec_0, dictMonadGen_1), gopurs_runtime.Value{Type: 9, IntVal: 2670894170, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_Gen_132451362_2738507278(Rebox_Data_Map_Gen_2738507278_132451362(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]](Get_Data_List_Types_unfoldableList()))))})
	_ = unfoldable1_5_3
	return gopurs_runtime.Func(func(dictOrd_6 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): fromFoldable_7_4 shape=App(Var) bindingType=(Func [(ADT ["Data","List","Types","List"] [(ADT ["Data","Tuple","Tuple"] [(TypeVar a$scope4), (TypeVar b$scope5)])])] (ADT ["Data","Map","Internal","Map"] [(TypeVar a$scope4), (TypeVar b$scope5)]))
		fromFoldable_7_4 := gopurs_runtime.Apply(Call_Data_Map_Internal_fromFoldable(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_6)), gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_Gen_1022383170_1680800814(Rebox_Data_Map_Gen_1680800814_1022383170(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Types_foldableList()))))})
		_ = fromFoldable_7_4
		return gopurs_runtime.Func2(func(genKey_8 gopurs_runtime.Value, genValue_9 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(dictMonadGen_1.V5, gopurs_runtime.Func(func(size_10 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_0, "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply2(dictMonadGen_1.V3, gopurs_runtime.Int(int64(0)), gopurs_runtime.Int(size_10.IntVal)), gopurs_runtime.Func(func(newSize_11 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply2(dictMonadGen_1.V4, gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
						return newSize_11
					}), gopurs_runtime.Apply2(Functor0_4_2.V0, fromFoldable_7_4, gopurs_runtime.Apply(unfoldable1_5_3, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Bind1_3_1, "Apply0"), gopurs_runtime.Value{}), "apply"), gopurs_runtime.Apply2(Functor0_4_2.V0, Get_Data_Tuple_Tuple(), genKey_8), genValue_9))))
				}))
			}))
		})
	})
}

func Rebox_Data_Map_Gen_1022383170_1680800814(in *Constructor_Data_Foldable_Foldable[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]) *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Foldable_Foldable[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Map_Gen_132451362_2738507278(in *Constructor_Data_Unfoldable_Unfoldable[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]) *Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Map_Gen_1680800814_1022383170(in *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]) *Constructor_Data_Foldable_Foldable[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Foldable_Foldable[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Data_Map_Gen_2738507278_132451362(in *Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]) *Constructor_Data_Unfoldable_Unfoldable[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Unfoldable_Unfoldable[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]])(unsafe.Pointer(in))
}
