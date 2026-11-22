package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Foreign_Object_Gen_fromFoldable gopurs_runtime.Value
var once_Foreign_Object_Gen_fromFoldable sync.Once

func Get_Foreign_Object_Gen_fromFoldable() gopurs_runtime.Value {
	once_Foreign_Object_Gen_fromFoldable.Do(func() {
		cache_Foreign_Object_Gen_fromFoldable = Call_Foreign_Object_fromFoldable(Rebox_Foreign_Object_Gen_1022383170_1680800814(Rebox_Foreign_Object_Gen_1680800814_1022383170(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Types_foldableList()))))
	})
	return cache_Foreign_Object_Gen_fromFoldable
}

var cache_Foreign_Object_Gen_genForeignObject gopurs_runtime.Value
var once_Foreign_Object_Gen_genForeignObject sync.Once

func Get_Foreign_Object_Gen_genForeignObject() gopurs_runtime.Value {
	once_Foreign_Object_Gen_genForeignObject.Do(func() {
		cache_Foreign_Object_Gen_genForeignObject = gopurs_runtime.Func2(func(dictMonadRec_0_box gopurs_runtime.Value, dictMonadGen_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Foreign_Object_Gen_genForeignObject(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]](dictMonadRec_0_box), gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]](dictMonadGen_1_box))
		})
	})
	return cache_Foreign_Object_Gen_genForeignObject
}

func Call_Foreign_Object_Gen_genForeignObject(dictMonadRec_0_loop *Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value], dictMonadGen_1_loop *Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]) gopurs_runtime.Value {
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
	// TAST (Let): Functor0_4_2 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m$scope2)])
	Functor0_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Bind1_3_1, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
	_ = Functor0_4_2
	// TAST (Let): unfoldable1_5_3 shape=App(Var) bindingType=(Func [(TypeApp (TypeVar m$scope2) [(ADT ["Data","Tuple","Tuple"] [String, (TypeVar a$scope3)])])] (TypeApp (TypeVar m$scope2) [(ADT ["Data","List","Types","List"] [(ADT ["Data","Tuple","Tuple"] [String, (TypeVar a$scope3)])])]))
	unfoldable1_5_3 := gopurs_runtime.Apply(Call_Control_Monad_Gen_unfoldable(dictMonadRec_0, dictMonadGen_1), gopurs_runtime.Value{Type: 9, IntVal: 2670894170, UnsafePtr: unsafe.Pointer(Rebox_Foreign_Object_Gen_132451362_2738507278(Rebox_Foreign_Object_Gen_2738507278_132451362(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]](Get_Data_List_Types_unfoldableList()))))})
	_ = unfoldable1_5_3
	return gopurs_runtime.Func2(func(genKey_6 gopurs_runtime.Value, genValue_7 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(dictMonadGen_1.V5, gopurs_runtime.Func(func(size_8 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_0, "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply2(dictMonadGen_1.V3, gopurs_runtime.Int(int64(0)), gopurs_runtime.Int(size_8.IntVal)), gopurs_runtime.Func(func(newSize_9 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(dictMonadGen_1.V4, gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
					return newSize_9
				}), gopurs_runtime.Apply2(Functor0_4_2.V0, Call_Foreign_Object_fromFoldable(Rebox_Foreign_Object_Gen_1022383170_1680800814(Rebox_Foreign_Object_Gen_1680800814_1022383170(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Types_foldableList())))), gopurs_runtime.Apply(unfoldable1_5_3, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Bind1_3_1, "Apply0"), gopurs_runtime.Value{}), "apply"), gopurs_runtime.Apply2(Functor0_4_2.V0, Get_Data_Tuple_Tuple(), genKey_6), genValue_7))))
			}))
		}))
	})
}

func Rebox_Foreign_Object_Gen_1022383170_1680800814(in *Constructor_Data_Foldable_Foldable[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]) *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Foldable_Foldable[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Foreign_Object_Gen_132451362_2738507278(in *Constructor_Data_Unfoldable_Unfoldable[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]) *Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Foreign_Object_Gen_1680800814_1022383170(in *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]) *Constructor_Data_Foldable_Foldable[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Foldable_Foldable[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Foreign_Object_Gen_2738507278_132451362(in *Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]) *Constructor_Data_Unfoldable_Unfoldable[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Unfoldable_Unfoldable[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]])(unsafe.Pointer(in))
}
