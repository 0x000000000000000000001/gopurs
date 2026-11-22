package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Map_SemigroupMap gopurs_runtime.Value
var once_Data_Map_SemigroupMap sync.Once

func Get_Data_Map_SemigroupMap() gopurs_runtime.Value {
	once_Data_Map_SemigroupMap.Do(func() {
		cache_Data_Map_SemigroupMap = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Map_SemigroupMap(x_0_box)
		})
	})
	return cache_Data_Map_SemigroupMap
}

var cache_Data_Map_semigroupSemigroupMap gopurs_runtime.Value
var once_Data_Map_semigroupSemigroupMap sync.Once

func Get_Data_Map_semigroupSemigroupMap() gopurs_runtime.Value {
	once_Data_Map_semigroupSemigroupMap.Do(func() {
		cache_Data_Map_semigroupSemigroupMap = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Map_semigroupSemigroupMap(dictOrd_0_box)
		})
	})
	return cache_Data_Map_semigroupSemigroupMap
}

var cache_Data_Map_newtypeSemigroupMap gopurs_runtime.Value
var once_Data_Map_newtypeSemigroupMap sync.Once

func Get_Data_Map_newtypeSemigroupMap() gopurs_runtime.Value {
	once_Data_Map_newtypeSemigroupMap.Do(func() {
		cache_Data_Map_newtypeSemigroupMap = gopurs_runtime.Value{Type: 9, IntVal: 3322196858, UnsafePtr: unsafe.Pointer((&Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{}
		})}))}
	})
	return cache_Data_Map_newtypeSemigroupMap
}

var cache_Data_Map_monoidSemigroupMap gopurs_runtime.Value
var once_Data_Map_monoidSemigroupMap sync.Once

func Get_Data_Map_monoidSemigroupMap() gopurs_runtime.Value {
	once_Data_Map_monoidSemigroupMap.Do(func() {
		cache_Data_Map_monoidSemigroupMap = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Map_monoidSemigroupMap(dictOrd_0_box)
		})
	})
	return cache_Data_Map_monoidSemigroupMap
}

var cache_Data_Map_keys gopurs_runtime.Value
var once_Data_Map_keys sync.Once

func Get_Data_Map_keys() gopurs_runtime.Value {
	once_Data_Map_keys.Do(func() {
		cache_Data_Map_keys = gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Set_fromMap(), Call_Data_Functor_void(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Map_Internal_functorMap())))
	})
	return cache_Data_Map_keys
}

var cache_Data_Map_eqSemigroupMap gopurs_runtime.Value
var once_Data_Map_eqSemigroupMap sync.Once

func Get_Data_Map_eqSemigroupMap() gopurs_runtime.Value {
	once_Data_Map_eqSemigroupMap.Do(func() {
		cache_Data_Map_eqSemigroupMap = gopurs_runtime.Func2(func(dictEq_0_box gopurs_runtime.Value, dictEq1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Map_eqSemigroupMap(dictEq_0_box, dictEq1_1_box)
		})
	})
	return cache_Data_Map_eqSemigroupMap
}

func Call_Data_Map_SemigroupMap(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Data_Map_semigroupSemigroupMap(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
	_ = dictOrd_0
	// TAST (Let): unionWith_1_0 shape=App(Var) bindingType=(Func [(Func [(TypeVar v$scope3), (TypeVar v$scope3)] (TypeVar v$scope3)), (ADT ["Data","Map","Internal","Map"] [(TypeVar k$scope2), (TypeVar v$scope3)]), (ADT ["Data","Map","Internal","Map"] [(TypeVar k$scope2), (TypeVar v$scope3)])] (ADT ["Data","Map","Internal","Map"] [(TypeVar k$scope2), (TypeVar v$scope3)]))
	unionWith_1_0 := gopurs_runtime.Apply(Get_Data_Map_Internal_unionWithImpl(), gopurs_runtime.Apply(Get_Data_Map_Internal_compareInt(), dictOrd_0))
	_ = unionWith_1_0
	return gopurs_runtime.Func(func(dictSemigroup_2 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): append_3_1 shape=App(Var) bindingType=(Func [(TypeVar v$scope3), (TypeVar v$scope3)] (TypeVar v$scope3))
		append_3_1 := Call_Data_Semigroup_go__append(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](dictSemigroup_2))
		_ = append_3_1
		return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer((&Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply3(unionWith_1_0, append_3_1, v_4, v1_5)
		})}))}
	})
}

func Call_Data_Map_monoidSemigroupMap(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
	_ = dictOrd_0
	// TAST (Let): semigroupSemigroupMap1_1_0 shape=App(Var) bindingType=Any
	semigroupSemigroupMap1_1_0 := Call_Data_Map_semigroupSemigroupMap(dictOrd_0)
	_ = semigroupSemigroupMap1_1_0
	return gopurs_runtime.Func(func(dictSemigroup_2 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): semigroupSemigroupMap2_3_1 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(ADT ["Data","Map","Internal","Map"] [(TypeVar k$scope16), (TypeVar v$scope17)])])
		semigroupSemigroupMap2_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(semigroupSemigroupMap1_1_0, dictSemigroup_2))
		_ = semigroupSemigroupMap2_3_1
		return gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer((&Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(semigroupSemigroupMap2_3_1)}
		}), Get_Data_Map_Internal_empty()}))}
	})
}

func Call_Data_Map_eqSemigroupMap(dictEq_0_loop gopurs_runtime.Value, dictEq1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
	_ = dictEq_0
	var dictEq1_1 gopurs_runtime.Value = dictEq1_1_loop
	_ = dictEq1_1
	return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_Map_Internal_eqMap(dictEq_0, dictEq1_1)))}
}
