package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Enum_Gen_foldable1NonEmpty gopurs_runtime.Value
var once_Data_Enum_Gen_foldable1NonEmpty sync.Once

func Get_Data_Enum_Gen_foldable1NonEmpty() gopurs_runtime.Value {
	once_Data_Enum_Gen_foldable1NonEmpty.Do(func() {
		cache_Data_Enum_Gen_foldable1NonEmpty = gopurs_runtime.Value{Type: 9, IntVal: 2465059545, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_Gen_4217626592_4151366573(Rebox_Data_Enum_Gen_4151366573_4217626592(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]](Call_Data_NonEmpty_foldable1NonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))})))))}
	})
	return cache_Data_Enum_Gen_foldable1NonEmpty
}

var cache_Data_Enum_Gen_genBoundedEnum gopurs_runtime.Value
var once_Data_Enum_Gen_genBoundedEnum sync.Once

func Get_Data_Enum_Gen_genBoundedEnum() gopurs_runtime.Value {
	once_Data_Enum_Gen_genBoundedEnum.Do(func() {
		cache_Data_Enum_Gen_genBoundedEnum = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Enum_Gen_genBoundedEnum(dictMonadGen_0_box)
		})
	})
	return cache_Data_Enum_Gen_genBoundedEnum
}

var cache_Data_Enum_Gen___cached_dict_0 gopurs_runtime.Value
var once_Data_Enum_Gen___cached_dict_0 sync.Once

func Get_Data_Enum_Gen___cached_dict_0() gopurs_runtime.Value {
	once_Data_Enum_Gen___cached_dict_0.Do(func() {
		cache_Data_Enum_Gen___cached_dict_0 = gopurs_runtime.Value{Type: 9, IntVal: 2465059545, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_Gen_4217626592_4151366573(Rebox_Data_Enum_Gen_4151366573_4217626592(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]](Call_Data_NonEmpty_foldable1NonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))})))))}
	})
	return cache_Data_Enum_Gen___cached_dict_0
}

func Call_Data_Enum_Gen_genBoundedEnum(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
	_ = dictMonadGen_0
	// TAST (Let): elements_1_0 shape=App(Var) bindingType=(Func [(ADT ["Data","NonEmpty","NonEmpty"] [(ADT ["Prim","Array"] []), (TypeVar a$scope3)])] (TypeApp (TypeVar m$scope2) [(TypeVar a$scope3)]))
	elements_1_0 := gopurs_runtime.Apply(Call_Control_Monad_Gen_elements(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]](dictMonadGen_0)), gopurs_runtime.Value{Type: 9, IntVal: 2465059545, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_Gen_4217626592_4151366573(Rebox_Data_Enum_Gen_4151366573_4217626592(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]](Get_Data_Enum_Gen___cached_dict_0()))))})
	_ = elements_1_0
	return gopurs_runtime.Func(func(dictBoundedEnum_2 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): Enum1_3_1 shape=App(Other) bindingType=(ADT ["Data","Enum","Enum"] [(TypeVar a$scope3)])
		Enum1_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum_2, "Enum1"), gopurs_runtime.Value{}))
		_ = Enum1_3_1
		// TAST (Let): Bounded0_4_2 shape=App(Other) bindingType=(ADT ["Data","Bounded","Bounded"] [(TypeVar a$scope3)])
		Bounded0_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum_2, "Bounded0"), gopurs_runtime.Value{}))
		_ = Bounded0_4_2
		// TAST (Let): v_5_3 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar a$scope3)])
		v_5_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Enum1_3_1.V2, Bounded0_4_2.V1))
		_ = v_5_3
		var __t4 gopurs_runtime.Value
		{
			if v_5_3 != nil {
				__t4 = gopurs_runtime.Apply(elements_1_0, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer((&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, Bounded0_4_2.V1, gopurs_runtime.Apply3(Call_Data_Enum_enumFromTo(Enum1_3_1), gopurs_runtime.Value{Type: 9, IntVal: 3553002490, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value]](Get_Data_Unfoldable1_unfoldable1Array()))}, (v_5_3).V0, Bounded0_4_2.V2)}))})
				goto end_branch_4
			} else {

			}
		}
		{
			if v_5_3 == nil {
				__t4 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{}), "Applicative0"), gopurs_runtime.Value{}), "pure"), Bounded0_4_2.V1)
				goto end_branch_4
			} else {

			}
		}
		{
			__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
		}
	end_branch_4:
		return __t4
	})
}

func Rebox_Data_Enum_Gen_4151366573_4217626592(in *Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]) *Constructor_Data_Semigroup_Foldable_Foldable1[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Semigroup_Foldable_Foldable1[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Data_Enum_Gen_4217626592_4151366573(in *Constructor_Data_Semigroup_Foldable_Foldable1[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value])(unsafe.Pointer(in))
}
