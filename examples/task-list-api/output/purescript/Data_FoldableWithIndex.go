package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_FoldableWithIndex_compose gopurs_runtime.Value
var once_Data_FoldableWithIndex_compose sync.Once

func Get_Data_FoldableWithIndex_compose() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_compose.Do(func() {
		cache_Data_FoldableWithIndex_compose = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_FoldableWithIndex_compose
}

var cache_Data_FoldableWithIndex_foldr gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldr sync.Once

func Get_Data_FoldableWithIndex_foldr() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldr.Do(func() {
		cache_Data_FoldableWithIndex_foldr = gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableMultiplicative()).V2
	})
	return cache_Data_FoldableWithIndex_foldr
}

var cache_Data_FoldableWithIndex_foldl gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldl sync.Once

func Get_Data_FoldableWithIndex_foldl() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldl.Do(func() {
		cache_Data_FoldableWithIndex_foldl = gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableMultiplicative()).V1
	})
	return cache_Data_FoldableWithIndex_foldl
}

var cache_Data_FoldableWithIndex_foldMap gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldMap sync.Once

func Get_Data_FoldableWithIndex_foldMap() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldMap.Do(func() {
		cache_Data_FoldableWithIndex_foldMap = gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableMultiplicative()).V0
	})
	return cache_Data_FoldableWithIndex_foldMap
}

var cache_Data_FoldableWithIndex_foldr1 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldr1 sync.Once

func Get_Data_FoldableWithIndex_foldr1() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldr1.Do(func() {
		cache_Data_FoldableWithIndex_foldr1 = Rebox_Data_FoldableWithIndex_1680800814_1146820559(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableMaybe())).V2
	})
	return cache_Data_FoldableWithIndex_foldr1
}

var cache_Data_FoldableWithIndex_foldl1 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldl1 sync.Once

func Get_Data_FoldableWithIndex_foldl1() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldl1.Do(func() {
		cache_Data_FoldableWithIndex_foldl1 = Rebox_Data_FoldableWithIndex_1680800814_1146820559(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableMaybe())).V1
	})
	return cache_Data_FoldableWithIndex_foldl1
}

var cache_Data_FoldableWithIndex_foldMap1 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldMap1 sync.Once

func Get_Data_FoldableWithIndex_foldMap1() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldMap1.Do(func() {
		cache_Data_FoldableWithIndex_foldMap1 = Rebox_Data_FoldableWithIndex_1680800814_1146820559(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableMaybe())).V0
	})
	return cache_Data_FoldableWithIndex_foldMap1
}

var cache_Data_FoldableWithIndex_foldr2 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldr2 sync.Once

func Get_Data_FoldableWithIndex_foldr2() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldr2.Do(func() {
		cache_Data_FoldableWithIndex_foldr2 = Rebox_Data_FoldableWithIndex_1680800814_1146820559(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableLast())).V2
	})
	return cache_Data_FoldableWithIndex_foldr2
}

var cache_Data_FoldableWithIndex_foldl2 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldl2 sync.Once

func Get_Data_FoldableWithIndex_foldl2() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldl2.Do(func() {
		cache_Data_FoldableWithIndex_foldl2 = Rebox_Data_FoldableWithIndex_1680800814_1146820559(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableLast())).V1
	})
	return cache_Data_FoldableWithIndex_foldl2
}

var cache_Data_FoldableWithIndex_foldMap2 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldMap2 sync.Once

func Get_Data_FoldableWithIndex_foldMap2() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldMap2.Do(func() {
		cache_Data_FoldableWithIndex_foldMap2 = Rebox_Data_FoldableWithIndex_1680800814_1146820559(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableLast())).V0
	})
	return cache_Data_FoldableWithIndex_foldMap2
}

var cache_Data_FoldableWithIndex_foldr3 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldr3 sync.Once

func Get_Data_FoldableWithIndex_foldr3() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldr3.Do(func() {
		cache_Data_FoldableWithIndex_foldr3 = Rebox_Data_FoldableWithIndex_1680800814_1146820559(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableFirst())).V2
	})
	return cache_Data_FoldableWithIndex_foldr3
}

var cache_Data_FoldableWithIndex_foldl3 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldl3 sync.Once

func Get_Data_FoldableWithIndex_foldl3() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldl3.Do(func() {
		cache_Data_FoldableWithIndex_foldl3 = Rebox_Data_FoldableWithIndex_1680800814_1146820559(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableFirst())).V1
	})
	return cache_Data_FoldableWithIndex_foldl3
}

var cache_Data_FoldableWithIndex_foldMap3 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldMap3 sync.Once

func Get_Data_FoldableWithIndex_foldMap3() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldMap3.Do(func() {
		cache_Data_FoldableWithIndex_foldMap3 = Rebox_Data_FoldableWithIndex_1680800814_1146820559(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableFirst())).V0
	})
	return cache_Data_FoldableWithIndex_foldMap3
}

var cache_Data_FoldableWithIndex_foldr4 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldr4 sync.Once

func Get_Data_FoldableWithIndex_foldr4() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldr4.Do(func() {
		cache_Data_FoldableWithIndex_foldr4 = gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableDual()).V2
	})
	return cache_Data_FoldableWithIndex_foldr4
}

var cache_Data_FoldableWithIndex_foldl4 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldl4 sync.Once

func Get_Data_FoldableWithIndex_foldl4() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldl4.Do(func() {
		cache_Data_FoldableWithIndex_foldl4 = gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableDual()).V1
	})
	return cache_Data_FoldableWithIndex_foldl4
}

var cache_Data_FoldableWithIndex_foldMap4 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldMap4 sync.Once

func Get_Data_FoldableWithIndex_foldMap4() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldMap4.Do(func() {
		cache_Data_FoldableWithIndex_foldMap4 = gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableDual()).V0
	})
	return cache_Data_FoldableWithIndex_foldMap4
}

var cache_Data_FoldableWithIndex_foldr5 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldr5 sync.Once

func Get_Data_FoldableWithIndex_foldr5() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldr5.Do(func() {
		cache_Data_FoldableWithIndex_foldr5 = gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableDisj()).V2
	})
	return cache_Data_FoldableWithIndex_foldr5
}

var cache_Data_FoldableWithIndex_foldl5 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldl5 sync.Once

func Get_Data_FoldableWithIndex_foldl5() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldl5.Do(func() {
		cache_Data_FoldableWithIndex_foldl5 = gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableDisj()).V1
	})
	return cache_Data_FoldableWithIndex_foldl5
}

var cache_Data_FoldableWithIndex_foldMap5 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldMap5 sync.Once

func Get_Data_FoldableWithIndex_foldMap5() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldMap5.Do(func() {
		cache_Data_FoldableWithIndex_foldMap5 = gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableDisj()).V0
	})
	return cache_Data_FoldableWithIndex_foldMap5
}

var cache_Data_FoldableWithIndex_foldr6 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldr6 sync.Once

func Get_Data_FoldableWithIndex_foldr6() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldr6.Do(func() {
		cache_Data_FoldableWithIndex_foldr6 = gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableConj()).V2
	})
	return cache_Data_FoldableWithIndex_foldr6
}

var cache_Data_FoldableWithIndex_foldl6 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldl6 sync.Once

func Get_Data_FoldableWithIndex_foldl6() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldl6.Do(func() {
		cache_Data_FoldableWithIndex_foldl6 = gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableConj()).V1
	})
	return cache_Data_FoldableWithIndex_foldl6
}

var cache_Data_FoldableWithIndex_foldMap6 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldMap6 sync.Once

func Get_Data_FoldableWithIndex_foldMap6() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldMap6.Do(func() {
		cache_Data_FoldableWithIndex_foldMap6 = gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableConj()).V0
	})
	return cache_Data_FoldableWithIndex_foldMap6
}

var cache_Data_FoldableWithIndex_foldr7 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldr7 sync.Once

func Get_Data_FoldableWithIndex_foldr7() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldr7.Do(func() {
		cache_Data_FoldableWithIndex_foldr7 = gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableAdditive()).V2
	})
	return cache_Data_FoldableWithIndex_foldr7
}

var cache_Data_FoldableWithIndex_foldl7 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldl7 sync.Once

func Get_Data_FoldableWithIndex_foldl7() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldl7.Do(func() {
		cache_Data_FoldableWithIndex_foldl7 = gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableAdditive()).V1
	})
	return cache_Data_FoldableWithIndex_foldl7
}

var cache_Data_FoldableWithIndex_foldMap7 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldMap7 sync.Once

func Get_Data_FoldableWithIndex_foldMap7() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldMap7.Do(func() {
		cache_Data_FoldableWithIndex_foldMap7 = gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableAdditive()).V0
	})
	return cache_Data_FoldableWithIndex_foldMap7
}

var cache_Data_FoldableWithIndex_compose1 gopurs_runtime.Value
var once_Data_FoldableWithIndex_compose1 sync.Once

func Get_Data_FoldableWithIndex_compose1() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_compose1.Do(func() {
		cache_Data_FoldableWithIndex_compose1 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_FoldableWithIndex_compose1
}

var cache_Data_FoldableWithIndex_compose2 gopurs_runtime.Value
var once_Data_FoldableWithIndex_compose2 sync.Once

func Get_Data_FoldableWithIndex_compose2() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_compose2.Do(func() {
		cache_Data_FoldableWithIndex_compose2 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_FoldableWithIndex_compose2
}

var cache_Data_FoldableWithIndex_compose3 gopurs_runtime.Value
var once_Data_FoldableWithIndex_compose3 sync.Once

func Get_Data_FoldableWithIndex_compose3() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_compose3.Do(func() {
		cache_Data_FoldableWithIndex_compose3 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_FoldableWithIndex_compose3
}

var cache_Data_FoldableWithIndex_compose4 gopurs_runtime.Value
var once_Data_FoldableWithIndex_compose4 sync.Once

func Get_Data_FoldableWithIndex_compose4() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_compose4.Do(func() {
		cache_Data_FoldableWithIndex_compose4 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_FoldableWithIndex_compose4
}

var cache_Data_FoldableWithIndex_compose5 gopurs_runtime.Value
var once_Data_FoldableWithIndex_compose5 sync.Once

func Get_Data_FoldableWithIndex_compose5() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_compose5.Do(func() {
		cache_Data_FoldableWithIndex_compose5 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_FoldableWithIndex_compose5
}

var cache_Data_FoldableWithIndex_compose6 gopurs_runtime.Value
var once_Data_FoldableWithIndex_compose6 sync.Once

func Get_Data_FoldableWithIndex_compose6() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_compose6.Do(func() {
		cache_Data_FoldableWithIndex_compose6 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_FoldableWithIndex_compose6
}

var cache_Data_FoldableWithIndex_compose7 gopurs_runtime.Value
var once_Data_FoldableWithIndex_compose7 sync.Once

func Get_Data_FoldableWithIndex_compose7() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_compose7.Do(func() {
		cache_Data_FoldableWithIndex_compose7 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_FoldableWithIndex_compose7
}

var cache_Data_FoldableWithIndex_compose8 gopurs_runtime.Value
var once_Data_FoldableWithIndex_compose8 sync.Once

func Get_Data_FoldableWithIndex_compose8() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_compose8.Do(func() {
		cache_Data_FoldableWithIndex_compose8 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_FoldableWithIndex_compose8
}

var cache_Data_FoldableWithIndex_compose9 gopurs_runtime.Value
var once_Data_FoldableWithIndex_compose9 sync.Once

func Get_Data_FoldableWithIndex_compose9() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_compose9.Do(func() {
		cache_Data_FoldableWithIndex_compose9 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_FoldableWithIndex_compose9
}

var cache_Data_FoldableWithIndex_compose10 gopurs_runtime.Value
var once_Data_FoldableWithIndex_compose10 sync.Once

func Get_Data_FoldableWithIndex_compose10() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_compose10.Do(func() {
		cache_Data_FoldableWithIndex_compose10 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_FoldableWithIndex_compose10
}

var cache_Data_FoldableWithIndex_compose11 gopurs_runtime.Value
var once_Data_FoldableWithIndex_compose11 sync.Once

func Get_Data_FoldableWithIndex_compose11() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_compose11.Do(func() {
		cache_Data_FoldableWithIndex_compose11 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_FoldableWithIndex_compose11
}

var cache_Data_FoldableWithIndex_compose12 gopurs_runtime.Value
var once_Data_FoldableWithIndex_compose12 sync.Once

func Get_Data_FoldableWithIndex_compose12() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_compose12.Do(func() {
		cache_Data_FoldableWithIndex_compose12 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_FoldableWithIndex_compose12
}

var cache_Data_FoldableWithIndex_compose13 gopurs_runtime.Value
var once_Data_FoldableWithIndex_compose13 sync.Once

func Get_Data_FoldableWithIndex_compose13() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_compose13.Do(func() {
		cache_Data_FoldableWithIndex_compose13 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_FoldableWithIndex_compose13
}

var cache_Data_FoldableWithIndex_compose14 gopurs_runtime.Value
var once_Data_FoldableWithIndex_compose14 sync.Once

func Get_Data_FoldableWithIndex_compose14() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_compose14.Do(func() {
		cache_Data_FoldableWithIndex_compose14 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_FoldableWithIndex_compose14
}

var cache_Data_FoldableWithIndex_unwrap gopurs_runtime.Value
var once_Data_FoldableWithIndex_unwrap sync.Once

func Get_Data_FoldableWithIndex_unwrap() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_unwrap.Do(func() {
		cache_Data_FoldableWithIndex_unwrap = Call_Safe_Coerce_coerce(gopurs_runtime.Value{})
	})
	return cache_Data_FoldableWithIndex_unwrap
}

var cache_Data_FoldableWithIndex_unwrap1 gopurs_runtime.Value
var once_Data_FoldableWithIndex_unwrap1 sync.Once

func Get_Data_FoldableWithIndex_unwrap1() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_unwrap1.Do(func() {
		cache_Data_FoldableWithIndex_unwrap1 = Call_Safe_Coerce_coerce(gopurs_runtime.Value{})
	})
	return cache_Data_FoldableWithIndex_unwrap1
}

var cache_Data_FoldableWithIndex_monoidDual gopurs_runtime.Value
var once_Data_FoldableWithIndex_monoidDual sync.Once

func Get_Data_FoldableWithIndex_monoidDual() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_monoidDual.Do(func() {
		cache_Data_FoldableWithIndex_monoidDual = gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](Call_Data_Monoid_Dual_monoidDual(Call_Data_Monoid_Endo_monoidEndo(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))}))))}
	})
	return cache_Data_FoldableWithIndex_monoidDual
}

var cache_Data_FoldableWithIndex_compose15 gopurs_runtime.Value
var once_Data_FoldableWithIndex_compose15 sync.Once

func Get_Data_FoldableWithIndex_compose15() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_compose15.Do(func() {
		cache_Data_FoldableWithIndex_compose15 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_FoldableWithIndex_compose15
}

var cache_Data_FoldableWithIndex_compose16 gopurs_runtime.Value
var once_Data_FoldableWithIndex_compose16 sync.Once

func Get_Data_FoldableWithIndex_compose16() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_compose16.Do(func() {
		cache_Data_FoldableWithIndex_compose16 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_FoldableWithIndex_compose16
}

var cache_Data_FoldableWithIndex_unwrap2 gopurs_runtime.Value
var once_Data_FoldableWithIndex_unwrap2 sync.Once

func Get_Data_FoldableWithIndex_unwrap2() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_unwrap2.Do(func() {
		cache_Data_FoldableWithIndex_unwrap2 = Call_Safe_Coerce_coerce(gopurs_runtime.Value{})
	})
	return cache_Data_FoldableWithIndex_unwrap2
}

var cache_Data_FoldableWithIndex_monoidEndo gopurs_runtime.Value
var once_Data_FoldableWithIndex_monoidEndo sync.Once

func Get_Data_FoldableWithIndex_monoidEndo() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_monoidEndo.Do(func() {
		cache_Data_FoldableWithIndex_monoidEndo = gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](Call_Data_Monoid_Endo_monoidEndo(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))})))}
	})
	return cache_Data_FoldableWithIndex_monoidEndo
}

var cache_Data_FoldableWithIndex_compose17 gopurs_runtime.Value
var once_Data_FoldableWithIndex_compose17 sync.Once

func Get_Data_FoldableWithIndex_compose17() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_compose17.Do(func() {
		cache_Data_FoldableWithIndex_compose17 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_FoldableWithIndex_compose17
}

var cache_Data_FoldableWithIndex_unwrap3 gopurs_runtime.Value
var once_Data_FoldableWithIndex_unwrap3 sync.Once

func Get_Data_FoldableWithIndex_unwrap3() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_unwrap3.Do(func() {
		cache_Data_FoldableWithIndex_unwrap3 = Call_Safe_Coerce_coerce(gopurs_runtime.Value{})
	})
	return cache_Data_FoldableWithIndex_unwrap3
}

var cache_Data_FoldableWithIndex_monoidEndo1 gopurs_runtime.Value
var once_Data_FoldableWithIndex_monoidEndo1 sync.Once

func Get_Data_FoldableWithIndex_monoidEndo1() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_monoidEndo1.Do(func() {
		cache_Data_FoldableWithIndex_monoidEndo1 = gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](Call_Data_Monoid_Endo_monoidEndo(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))})))}
	})
	return cache_Data_FoldableWithIndex_monoidEndo1
}

var cache_Data_FoldableWithIndex_compose18 gopurs_runtime.Value
var once_Data_FoldableWithIndex_compose18 sync.Once

func Get_Data_FoldableWithIndex_compose18() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_compose18.Do(func() {
		cache_Data_FoldableWithIndex_compose18 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_FoldableWithIndex_compose18
}

var cache_Data_FoldableWithIndex_unwrap4 gopurs_runtime.Value
var once_Data_FoldableWithIndex_unwrap4 sync.Once

func Get_Data_FoldableWithIndex_unwrap4() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_unwrap4.Do(func() {
		cache_Data_FoldableWithIndex_unwrap4 = Call_Data_Newtype_unwrap(gopurs_runtime.CoerceToStruct[Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{}))
	})
	return cache_Data_FoldableWithIndex_unwrap4
}

var cache_Data_FoldableWithIndex_compose19 gopurs_runtime.Value
var once_Data_FoldableWithIndex_compose19 sync.Once

func Get_Data_FoldableWithIndex_compose19() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_compose19.Do(func() {
		cache_Data_FoldableWithIndex_compose19 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_FoldableWithIndex_compose19
}

var cache_Data_FoldableWithIndex_compose20 gopurs_runtime.Value
var once_Data_FoldableWithIndex_compose20 sync.Once

func Get_Data_FoldableWithIndex_compose20() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_compose20.Do(func() {
		cache_Data_FoldableWithIndex_compose20 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_FoldableWithIndex_compose20
}

var cache_Data_FoldableWithIndex_unwrap5 gopurs_runtime.Value
var once_Data_FoldableWithIndex_unwrap5 sync.Once

func Get_Data_FoldableWithIndex_unwrap5() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_unwrap5.Do(func() {
		cache_Data_FoldableWithIndex_unwrap5 = Call_Data_Newtype_unwrap(gopurs_runtime.CoerceToStruct[Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{}))
	})
	return cache_Data_FoldableWithIndex_unwrap5
}

var cache_Data_FoldableWithIndex_compose21 gopurs_runtime.Value
var once_Data_FoldableWithIndex_compose21 sync.Once

func Get_Data_FoldableWithIndex_compose21() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_compose21.Do(func() {
		cache_Data_FoldableWithIndex_compose21 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_FoldableWithIndex_compose21
}

var cache_Data_FoldableWithIndex_FoldableWithIndex_dollar_Dict gopurs_runtime.Value
var once_Data_FoldableWithIndex_FoldableWithIndex_dollar_Dict sync.Once

func Get_Data_FoldableWithIndex_FoldableWithIndex_dollar_Dict() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_FoldableWithIndex_dollar_Dict.Do(func() {
		cache_Data_FoldableWithIndex_FoldableWithIndex_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(Call_Data_FoldableWithIndex_FoldableWithIndex_dollar_Dict(func() struct {
				Foldable0        gopurs_runtime.Value
				foldMapWithIndex gopurs_runtime.Value
				foldlWithIndex   gopurs_runtime.Value
				foldrWithIndex   gopurs_runtime.Value
				/*gopurs:record-v1:4:0046006f006c006400610062006c00650030,0066006f006c0064004d0061007000570069007400680049006e006400650078,0066006f006c0064006c00570069007400680049006e006400650078,0066006f006c0064007200570069007400680049006e006400650078*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Foldable0        gopurs_runtime.Value
					foldMapWithIndex gopurs_runtime.Value
					foldlWithIndex   gopurs_runtime.Value
					foldrWithIndex   gopurs_runtime.Value
					/*gopurs:record-v1:4:0046006f006c006400610062006c00650030,0066006f006c0064004d0061007000570069007400680049006e006400650078,0066006f006c0064006c00570069007400680049006e006400650078,0066006f006c0064007200570069007400680049006e006400650078*/
				}{}
				clone.Foldable0 = gopurs_runtime.RecordGet(orig, "Foldable0")
				clone.foldMapWithIndex = gopurs_runtime.RecordGet(orig, "foldMapWithIndex")
				clone.foldlWithIndex = gopurs_runtime.RecordGet(orig, "foldlWithIndex")
				clone.foldrWithIndex = gopurs_runtime.RecordGet(orig, "foldrWithIndex")
				return clone
			}()))}
		})
	})
	return cache_Data_FoldableWithIndex_FoldableWithIndex_dollar_Dict
}

var cache_Data_FoldableWithIndex_foldrWithIndex gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldrWithIndex sync.Once

func Get_Data_FoldableWithIndex_foldrWithIndex() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldrWithIndex.Do(func() {
		cache_Data_FoldableWithIndex_foldrWithIndex = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_FoldableWithIndex_foldrWithIndex(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Data_FoldableWithIndex_foldrWithIndex
}

var cache_Data_FoldableWithIndex_traverseWithIndex_ gopurs_runtime.Value
var once_Data_FoldableWithIndex_traverseWithIndex_ sync.Once

func Get_Data_FoldableWithIndex_traverseWithIndex_() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_traverseWithIndex_.Do(func() {
		cache_Data_FoldableWithIndex_traverseWithIndex_ = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_FoldableWithIndex_traverseWithIndex_(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0_box))
		})
	})
	return cache_Data_FoldableWithIndex_traverseWithIndex_
}

var cache_Data_FoldableWithIndex_forWithIndex_ gopurs_runtime.Value
var once_Data_FoldableWithIndex_forWithIndex_ sync.Once

func Get_Data_FoldableWithIndex_forWithIndex_() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_forWithIndex_.Do(func() {
		cache_Data_FoldableWithIndex_forWithIndex_ = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_FoldableWithIndex_forWithIndex_(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0_box))
		})
	})
	return cache_Data_FoldableWithIndex_forWithIndex_
}

var cache_Data_FoldableWithIndex_foldrDefault gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldrDefault sync.Once

func Get_Data_FoldableWithIndex_foldrDefault() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldrDefault.Do(func() {
		cache_Data_FoldableWithIndex_foldrDefault = gopurs_runtime.Func2(func(dictFoldableWithIndex_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_FoldableWithIndex_foldrDefault(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dictFoldableWithIndex_0_box), f_1_box)
		})
	})
	return cache_Data_FoldableWithIndex_foldrDefault
}

var cache_Data_FoldableWithIndex_foldlWithIndex gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldlWithIndex sync.Once

func Get_Data_FoldableWithIndex_foldlWithIndex() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldlWithIndex.Do(func() {
		cache_Data_FoldableWithIndex_foldlWithIndex = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_FoldableWithIndex_foldlWithIndex(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Data_FoldableWithIndex_foldlWithIndex
}

var cache_Data_FoldableWithIndex_foldlDefault gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldlDefault sync.Once

func Get_Data_FoldableWithIndex_foldlDefault() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldlDefault.Do(func() {
		cache_Data_FoldableWithIndex_foldlDefault = gopurs_runtime.Func2(func(dictFoldableWithIndex_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_FoldableWithIndex_foldlDefault(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dictFoldableWithIndex_0_box), f_1_box)
		})
	})
	return cache_Data_FoldableWithIndex_foldlDefault
}

var cache_Data_FoldableWithIndex_foldableWithIndexTuple gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldableWithIndexTuple sync.Once

func Get_Data_FoldableWithIndex_foldableWithIndexTuple() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldableWithIndexTuple.Do(func() {
		cache_Data_FoldableWithIndex_foldableWithIndexTuple = gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(Rebox_Data_FoldableWithIndex_3374046885_3725484264((&Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Get_Data_Foldable_foldableTuple()
		}), gopurs_runtime.Func3(func(dictMonoid_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(f_1, Get_Data_Unit_unit(), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1)
		}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply3(f_0, Get_Data_Unit_unit(), z_1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1)
		}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply3(f_0, Get_Data_Unit_unit(), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1, z_1)
		})})))}
	})
	return cache_Data_FoldableWithIndex_foldableWithIndexTuple
}

var cache_Data_FoldableWithIndex_foldableWithIndexMultiplicative gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldableWithIndexMultiplicative sync.Once

func Get_Data_FoldableWithIndex_foldableWithIndexMultiplicative() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldableWithIndexMultiplicative.Do(func() {
		cache_Data_FoldableWithIndex_foldableWithIndexMultiplicative = gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer((&Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableMultiplicative()))}
		}), gopurs_runtime.Func2(func(dictMonoid_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_2_0 shape=App(Other) bindingType=(Func [(TypeVar a$scope73)] (TypeVar m$scope74))
			__local_var_2_0 := gopurs_runtime.Apply(f_1, Get_Data_Unit_unit())
			_ = __local_var_2_0
			return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply(__local_var_2_0, v_3)
			})
		}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_1_1 shape=App(Other) bindingType=(Func [(TypeVar b$scope72), (TypeVar a$scope71)] (TypeVar b$scope72))
			__local_var_1_1 := gopurs_runtime.Apply(f_0, Get_Data_Unit_unit())
			_ = __local_var_1_1
			return gopurs_runtime.Func2(func(z_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(__local_var_1_1, z_2, v_3)
			})
		}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_1_2 shape=App(Other) bindingType=(Func [(TypeVar a$scope69), (TypeVar b$scope70)] (TypeVar b$scope70))
			__local_var_1_2 := gopurs_runtime.Apply(f_0, Get_Data_Unit_unit())
			_ = __local_var_1_2
			return gopurs_runtime.Func2(func(z_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(__local_var_1_2, v_3, z_2)
			})
		})}))}
	})
	return cache_Data_FoldableWithIndex_foldableWithIndexMultiplicative
}

var cache_Data_FoldableWithIndex_foldableWithIndexMaybe gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldableWithIndexMaybe sync.Once

func Get_Data_FoldableWithIndex_foldableWithIndexMaybe() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldableWithIndexMaybe.Do(func() {
		cache_Data_FoldableWithIndex_foldableWithIndexMaybe = gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(Rebox_Data_FoldableWithIndex_2286084809_3725484264((&Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, *Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(Rebox_Data_FoldableWithIndex_1146820559_1680800814(Rebox_Data_FoldableWithIndex_1680800814_1146820559(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableMaybe()))))}
		}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): foldMap8_1_0 shape=App(Other) bindingType=(Func [(Func [(TypeVar a$scope87)] (TypeVar m$scope88)), (ADT ["Data","Maybe","Maybe"] [(TypeVar a$scope87)])] (TypeVar m$scope88))
			foldMap8_1_0 := gopurs_runtime.Apply(Rebox_Data_FoldableWithIndex_1680800814_1146820559(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableMaybe())).V0, dictMonoid_0)
			_ = foldMap8_1_0
			return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply(foldMap8_1_0, gopurs_runtime.Apply(f_2, Get_Data_Unit_unit()))
			})
		}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_1_1 shape=App(Other) bindingType=(Func [(TypeVar b$scope86), (TypeVar a$scope85)] (TypeVar b$scope86))
			__local_var_1_1 := gopurs_runtime.Apply(f_0, Get_Data_Unit_unit())
			_ = __local_var_1_1
			return gopurs_runtime.Func2(func(v1_2 gopurs_runtime.Value, v2_3 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t2 gopurs_runtime.Value
				{
					if v2_3.Type == 9 && v2_3.IntVal == 930809136 && v2_3.UnsafePtr == nil {
						__t2 = v1_2
						goto end_branch_2
					} else {

					}
				}
				{
					if v2_3.Type == 9 && v2_3.IntVal == 930809136 && v2_3.UnsafePtr != nil {
						__t2 = gopurs_runtime.Apply2(__local_var_1_1, v1_2, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v2_3.UnsafePtr).V0)
						goto end_branch_2
					} else {

					}
				}
				{
					__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
				}
			end_branch_2:
				return __t2
			})
		}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_1_3 shape=App(Other) bindingType=(Func [(TypeVar a$scope83), (TypeVar b$scope84)] (TypeVar b$scope84))
			__local_var_1_3 := gopurs_runtime.Apply(f_0, Get_Data_Unit_unit())
			_ = __local_var_1_3
			return gopurs_runtime.Func2(func(v1_2 gopurs_runtime.Value, v2_3 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t4 gopurs_runtime.Value
				{
					if v2_3.Type == 9 && v2_3.IntVal == 930809136 && v2_3.UnsafePtr == nil {
						__t4 = v1_2
						goto end_branch_4
					} else {

					}
				}
				{
					if v2_3.Type == 9 && v2_3.IntVal == 930809136 && v2_3.UnsafePtr != nil {
						__t4 = gopurs_runtime.Apply2(__local_var_1_3, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v2_3.UnsafePtr).V0, v1_2)
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
		})})))}
	})
	return cache_Data_FoldableWithIndex_foldableWithIndexMaybe
}

var cache_Data_FoldableWithIndex_foldableWithIndexLast gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldableWithIndexLast sync.Once

func Get_Data_FoldableWithIndex_foldableWithIndexLast() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldableWithIndexLast.Do(func() {
		cache_Data_FoldableWithIndex_foldableWithIndexLast = gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(Rebox_Data_FoldableWithIndex_2286084809_3725484264((&Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, *Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(Rebox_Data_FoldableWithIndex_1146820559_1680800814(Rebox_Data_FoldableWithIndex_1680800814_1146820559(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableLast()))))}
		}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): foldMap3_1_0 shape=App(Other) bindingType=(Func [(Func [(TypeVar a$scope376)] (TypeVar m$scope377)), (ADT ["Data","Maybe","Maybe"] [(TypeVar a$scope376)])] (TypeVar m$scope377))
			foldMap3_1_0 := gopurs_runtime.Apply(Rebox_Data_FoldableWithIndex_1680800814_1146820559(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableMaybe())).V0, gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_0))})
			_ = foldMap3_1_0
			return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
				// TAST (Let): __local_var_3_1 shape=App(Other) bindingType=(Func [(TypeVar a$scope101)] (TypeVar m$scope102))
				__local_var_3_1 := gopurs_runtime.Apply(f_2, Get_Data_Unit_unit())
				_ = __local_var_3_1
				return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply2(foldMap3_1_0, __local_var_3_1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_4))})
				})
			})
		}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_1_2 shape=App(Other) bindingType=(Func [(TypeVar b$scope100), (TypeVar a$scope99)] (TypeVar b$scope100))
			__local_var_1_2 := gopurs_runtime.Apply(f_0, Get_Data_Unit_unit())
			_ = __local_var_1_2
			return gopurs_runtime.Func2(func(z_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t3 gopurs_runtime.Value
				{
					if v_3.Type == 9 && v_3.IntVal == 930809136 && v_3.UnsafePtr == nil {
						__t3 = z_2
						goto end_branch_3
					} else {

					}
				}
				{
					if v_3.Type == 9 && v_3.IntVal == 930809136 && v_3.UnsafePtr != nil {
						__t3 = gopurs_runtime.Apply2(__local_var_1_2, z_2, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v_3.UnsafePtr).V0)
						goto end_branch_3
					} else {

					}
				}
				{
					__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
				}
			end_branch_3:
				return __t3
			})
		}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_1_4 shape=App(Other) bindingType=(Func [(TypeVar a$scope97), (TypeVar b$scope98)] (TypeVar b$scope98))
			__local_var_1_4 := gopurs_runtime.Apply(f_0, Get_Data_Unit_unit())
			_ = __local_var_1_4
			return gopurs_runtime.Func2(func(z_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t5 gopurs_runtime.Value
				{
					if v_3.Type == 9 && v_3.IntVal == 930809136 && v_3.UnsafePtr == nil {
						__t5 = z_2
						goto end_branch_5
					} else {

					}
				}
				{
					if v_3.Type == 9 && v_3.IntVal == 930809136 && v_3.UnsafePtr != nil {
						__t5 = gopurs_runtime.Apply2(__local_var_1_4, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v_3.UnsafePtr).V0, z_2)
						goto end_branch_5
					} else {

					}
				}
				{
					__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
				}
			end_branch_5:
				return __t5
			})
		})})))}
	})
	return cache_Data_FoldableWithIndex_foldableWithIndexLast
}

var cache_Data_FoldableWithIndex_foldableWithIndexIdentity gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldableWithIndexIdentity sync.Once

func Get_Data_FoldableWithIndex_foldableWithIndexIdentity() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldableWithIndexIdentity.Do(func() {
		cache_Data_FoldableWithIndex_foldableWithIndexIdentity = gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer((&Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableIdentity()))}
		}), gopurs_runtime.Func3(func(dictMonoid_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(f_1, Get_Data_Unit_unit(), v_2)
		}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply3(f_0, Get_Data_Unit_unit(), z_1, v_2)
		}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply3(f_0, Get_Data_Unit_unit(), v_2, z_1)
		})}))}
	})
	return cache_Data_FoldableWithIndex_foldableWithIndexIdentity
}

var cache_Data_FoldableWithIndex_foldableWithIndexFirst gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldableWithIndexFirst sync.Once

func Get_Data_FoldableWithIndex_foldableWithIndexFirst() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldableWithIndexFirst.Do(func() {
		cache_Data_FoldableWithIndex_foldableWithIndexFirst = gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(Rebox_Data_FoldableWithIndex_2286084809_3725484264((&Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, *Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(Rebox_Data_FoldableWithIndex_1146820559_1680800814(Rebox_Data_FoldableWithIndex_1680800814_1146820559(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableFirst()))))}
		}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): foldMap3_1_0 shape=App(Other) bindingType=(Func [(Func [(TypeVar a$scope363)] (TypeVar m$scope364)), (ADT ["Data","Maybe","Maybe"] [(TypeVar a$scope363)])] (TypeVar m$scope364))
			foldMap3_1_0 := gopurs_runtime.Apply(Rebox_Data_FoldableWithIndex_1680800814_1146820559(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableMaybe())).V0, gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_0))})
			_ = foldMap3_1_0
			return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
				// TAST (Let): __local_var_3_1 shape=App(Other) bindingType=(Func [(TypeVar a$scope129)] (TypeVar m$scope130))
				__local_var_3_1 := gopurs_runtime.Apply(f_2, Get_Data_Unit_unit())
				_ = __local_var_3_1
				return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply2(foldMap3_1_0, __local_var_3_1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_4))})
				})
			})
		}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_1_2 shape=App(Other) bindingType=(Func [(TypeVar b$scope128), (TypeVar a$scope127)] (TypeVar b$scope128))
			__local_var_1_2 := gopurs_runtime.Apply(f_0, Get_Data_Unit_unit())
			_ = __local_var_1_2
			return gopurs_runtime.Func2(func(z_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t3 gopurs_runtime.Value
				{
					if v_3.Type == 9 && v_3.IntVal == 930809136 && v_3.UnsafePtr == nil {
						__t3 = z_2
						goto end_branch_3
					} else {

					}
				}
				{
					if v_3.Type == 9 && v_3.IntVal == 930809136 && v_3.UnsafePtr != nil {
						__t3 = gopurs_runtime.Apply2(__local_var_1_2, z_2, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v_3.UnsafePtr).V0)
						goto end_branch_3
					} else {

					}
				}
				{
					__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
				}
			end_branch_3:
				return __t3
			})
		}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_1_4 shape=App(Other) bindingType=(Func [(TypeVar a$scope125), (TypeVar b$scope126)] (TypeVar b$scope126))
			__local_var_1_4 := gopurs_runtime.Apply(f_0, Get_Data_Unit_unit())
			_ = __local_var_1_4
			return gopurs_runtime.Func2(func(z_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t5 gopurs_runtime.Value
				{
					if v_3.Type == 9 && v_3.IntVal == 930809136 && v_3.UnsafePtr == nil {
						__t5 = z_2
						goto end_branch_5
					} else {

					}
				}
				{
					if v_3.Type == 9 && v_3.IntVal == 930809136 && v_3.UnsafePtr != nil {
						__t5 = gopurs_runtime.Apply2(__local_var_1_4, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v_3.UnsafePtr).V0, z_2)
						goto end_branch_5
					} else {

					}
				}
				{
					__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
				}
			end_branch_5:
				return __t5
			})
		})})))}
	})
	return cache_Data_FoldableWithIndex_foldableWithIndexFirst
}

var cache_Data_FoldableWithIndex_foldableWithIndexEither gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldableWithIndexEither sync.Once

func Get_Data_FoldableWithIndex_foldableWithIndexEither() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldableWithIndexEither.Do(func() {
		cache_Data_FoldableWithIndex_foldableWithIndexEither = gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer((&Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Get_Data_Foldable_foldableEither()
		}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): mempty_1_0 shape=App(Var) bindingType=(TypeVar m$scope146)
			mempty_1_0 := Call_Data_Monoid_mempty(dictMonoid_0)
			_ = mempty_1_0
			return gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t1 gopurs_runtime.Value
				{
					if v1_3.Type == 9 && v1_3.IntVal == 3711209382 {
						__t1 = mempty_1_0
						goto end_branch_1
					} else {

					}
				}
				{
					if v1_3.Type == 9 && v1_3.IntVal == 2465973597 {
						__t1 = gopurs_runtime.Apply2(v_2, Get_Data_Unit_unit(), (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0)
						goto end_branch_1
					} else {

					}
				}
				{
					__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
				}
			end_branch_1:
				return __t1
			})
		}), gopurs_runtime.Func3(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value, v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t2 gopurs_runtime.Value
			{
				if v2_2.Type == 9 && v2_2.IntVal == 3711209382 {
					__t2 = v1_1
					goto end_branch_2
				} else {

				}
			}
			{
				if v2_2.Type == 9 && v2_2.IntVal == 2465973597 {
					__t2 = gopurs_runtime.Apply3(v_0, Get_Data_Unit_unit(), v1_1, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)
					goto end_branch_2
				} else {

				}
			}
			{
				__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_2:
			return __t2
		}), gopurs_runtime.Func3(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value, v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t3 gopurs_runtime.Value
			{
				if v2_2.Type == 9 && v2_2.IntVal == 3711209382 {
					__t3 = v1_1
					goto end_branch_3
				} else {

				}
			}
			{
				if v2_2.Type == 9 && v2_2.IntVal == 2465973597 {
					__t3 = gopurs_runtime.Apply3(v_0, Get_Data_Unit_unit(), (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0, v1_1)
					goto end_branch_3
				} else {

				}
			}
			{
				__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_3:
			return __t3
		})}))}
	})
	return cache_Data_FoldableWithIndex_foldableWithIndexEither
}

var cache_Data_FoldableWithIndex_foldableWithIndexDual gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldableWithIndexDual sync.Once

func Get_Data_FoldableWithIndex_foldableWithIndexDual() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldableWithIndexDual.Do(func() {
		cache_Data_FoldableWithIndex_foldableWithIndexDual = gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer((&Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableDual()))}
		}), gopurs_runtime.Func2(func(dictMonoid_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_2_0 shape=App(Other) bindingType=(Func [(TypeVar a$scope159)] (TypeVar m$scope160))
			__local_var_2_0 := gopurs_runtime.Apply(f_1, Get_Data_Unit_unit())
			_ = __local_var_2_0
			return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply(__local_var_2_0, v_3)
			})
		}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_1_1 shape=App(Other) bindingType=(Func [(TypeVar b$scope158), (TypeVar a$scope157)] (TypeVar b$scope158))
			__local_var_1_1 := gopurs_runtime.Apply(f_0, Get_Data_Unit_unit())
			_ = __local_var_1_1
			return gopurs_runtime.Func2(func(z_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(__local_var_1_1, z_2, v_3)
			})
		}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_1_2 shape=App(Other) bindingType=(Func [(TypeVar a$scope155), (TypeVar b$scope156)] (TypeVar b$scope156))
			__local_var_1_2 := gopurs_runtime.Apply(f_0, Get_Data_Unit_unit())
			_ = __local_var_1_2
			return gopurs_runtime.Func2(func(z_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(__local_var_1_2, v_3, z_2)
			})
		})}))}
	})
	return cache_Data_FoldableWithIndex_foldableWithIndexDual
}

var cache_Data_FoldableWithIndex_foldableWithIndexDisj gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldableWithIndexDisj sync.Once

func Get_Data_FoldableWithIndex_foldableWithIndexDisj() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldableWithIndexDisj.Do(func() {
		cache_Data_FoldableWithIndex_foldableWithIndexDisj = gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer((&Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableDisj()))}
		}), gopurs_runtime.Func2(func(dictMonoid_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_2_0 shape=App(Other) bindingType=(Func [(TypeVar a$scope173)] (TypeVar m$scope174))
			__local_var_2_0 := gopurs_runtime.Apply(f_1, Get_Data_Unit_unit())
			_ = __local_var_2_0
			return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply(__local_var_2_0, v_3)
			})
		}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_1_1 shape=App(Other) bindingType=(Func [(TypeVar b$scope172), (TypeVar a$scope171)] (TypeVar b$scope172))
			__local_var_1_1 := gopurs_runtime.Apply(f_0, Get_Data_Unit_unit())
			_ = __local_var_1_1
			return gopurs_runtime.Func2(func(z_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(__local_var_1_1, z_2, v_3)
			})
		}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_1_2 shape=App(Other) bindingType=(Func [(TypeVar a$scope169), (TypeVar b$scope170)] (TypeVar b$scope170))
			__local_var_1_2 := gopurs_runtime.Apply(f_0, Get_Data_Unit_unit())
			_ = __local_var_1_2
			return gopurs_runtime.Func2(func(z_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(__local_var_1_2, v_3, z_2)
			})
		})}))}
	})
	return cache_Data_FoldableWithIndex_foldableWithIndexDisj
}

var cache_Data_FoldableWithIndex_foldableWithIndexConst gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldableWithIndexConst sync.Once

func Get_Data_FoldableWithIndex_foldableWithIndexConst() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldableWithIndexConst.Do(func() {
		cache_Data_FoldableWithIndex_foldableWithIndexConst = gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer((&Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Get_Data_Foldable_foldableConst()
		}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): mempty_1_0 shape=App(Var) bindingType=(TypeVar m$scope190)
			mempty_1_0 := Call_Data_Monoid_mempty(dictMonoid_0)
			_ = mempty_1_0
			return gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
				return mempty_1_0
			})
		}), gopurs_runtime.Func3(func(v_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return z_1
		}), gopurs_runtime.Func3(func(v_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return z_1
		})}))}
	})
	return cache_Data_FoldableWithIndex_foldableWithIndexConst
}

var cache_Data_FoldableWithIndex_foldableWithIndexConj gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldableWithIndexConj sync.Once

func Get_Data_FoldableWithIndex_foldableWithIndexConj() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldableWithIndexConj.Do(func() {
		cache_Data_FoldableWithIndex_foldableWithIndexConj = gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer((&Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableConj()))}
		}), gopurs_runtime.Func2(func(dictMonoid_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_2_0 shape=App(Other) bindingType=(Func [(TypeVar a$scope203)] (TypeVar m$scope204))
			__local_var_2_0 := gopurs_runtime.Apply(f_1, Get_Data_Unit_unit())
			_ = __local_var_2_0
			return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply(__local_var_2_0, v_3)
			})
		}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_1_1 shape=App(Other) bindingType=(Func [(TypeVar b$scope202), (TypeVar a$scope201)] (TypeVar b$scope202))
			__local_var_1_1 := gopurs_runtime.Apply(f_0, Get_Data_Unit_unit())
			_ = __local_var_1_1
			return gopurs_runtime.Func2(func(z_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(__local_var_1_1, z_2, v_3)
			})
		}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_1_2 shape=App(Other) bindingType=(Func [(TypeVar a$scope199), (TypeVar b$scope200)] (TypeVar b$scope200))
			__local_var_1_2 := gopurs_runtime.Apply(f_0, Get_Data_Unit_unit())
			_ = __local_var_1_2
			return gopurs_runtime.Func2(func(z_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(__local_var_1_2, v_3, z_2)
			})
		})}))}
	})
	return cache_Data_FoldableWithIndex_foldableWithIndexConj
}

var cache_Data_FoldableWithIndex_foldableWithIndexAdditive gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldableWithIndexAdditive sync.Once

func Get_Data_FoldableWithIndex_foldableWithIndexAdditive() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldableWithIndexAdditive.Do(func() {
		cache_Data_FoldableWithIndex_foldableWithIndexAdditive = gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer((&Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableAdditive()))}
		}), gopurs_runtime.Func2(func(dictMonoid_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_2_0 shape=App(Other) bindingType=(Func [(TypeVar a$scope217)] (TypeVar m$scope218))
			__local_var_2_0 := gopurs_runtime.Apply(f_1, Get_Data_Unit_unit())
			_ = __local_var_2_0
			return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply(__local_var_2_0, v_3)
			})
		}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_1_1 shape=App(Other) bindingType=(Func [(TypeVar b$scope216), (TypeVar a$scope215)] (TypeVar b$scope216))
			__local_var_1_1 := gopurs_runtime.Apply(f_0, Get_Data_Unit_unit())
			_ = __local_var_1_1
			return gopurs_runtime.Func2(func(z_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(__local_var_1_1, z_2, v_3)
			})
		}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_1_2 shape=App(Other) bindingType=(Func [(TypeVar a$scope213), (TypeVar b$scope214)] (TypeVar b$scope214))
			__local_var_1_2 := gopurs_runtime.Apply(f_0, Get_Data_Unit_unit())
			_ = __local_var_1_2
			return gopurs_runtime.Func2(func(z_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(__local_var_1_2, v_3, z_2)
			})
		})}))}
	})
	return cache_Data_FoldableWithIndex_foldableWithIndexAdditive
}

var cache_Data_FoldableWithIndex_foldWithIndexM gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldWithIndexM sync.Once

func Get_Data_FoldableWithIndex_foldWithIndexM() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldWithIndexM.Do(func() {
		cache_Data_FoldableWithIndex_foldWithIndexM = gopurs_runtime.Func4(func(dictFoldableWithIndex_0_box gopurs_runtime.Value, dictMonad_1_box gopurs_runtime.Value, f_2_box gopurs_runtime.Value, a0_3_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_FoldableWithIndex_foldWithIndexM(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dictFoldableWithIndex_0_box), gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](dictMonad_1_box), f_2_box, a0_3_box)
		})
	})
	return cache_Data_FoldableWithIndex_foldWithIndexM
}

var cache_Data_FoldableWithIndex_foldMapWithIndexDefaultR gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldMapWithIndexDefaultR sync.Once

func Get_Data_FoldableWithIndex_foldMapWithIndexDefaultR() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldMapWithIndexDefaultR.Do(func() {
		cache_Data_FoldableWithIndex_foldMapWithIndexDefaultR = gopurs_runtime.Func2(func(dictFoldableWithIndex_0_box gopurs_runtime.Value, dictMonoid_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_FoldableWithIndex_foldMapWithIndexDefaultR(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dictFoldableWithIndex_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_1_box))
		})
	})
	return cache_Data_FoldableWithIndex_foldMapWithIndexDefaultR
}

var cache_Data_FoldableWithIndex_foldableWithIndexArray gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldableWithIndexArray sync.Once

func Get_Data_FoldableWithIndex_foldableWithIndexArray() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldableWithIndexArray.Do(func() {
		cache_Data_FoldableWithIndex_foldableWithIndexArray = gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(Rebox_Data_FoldableWithIndex_2491554675_3725484264((&Constructor_Data_FoldableWithIndex_FoldableWithIndex[int64, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))}
		}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_FoldableWithIndex_foldMapWithIndexDefaultR(Rebox_Data_FoldableWithIndex_2491554675_3725484264(Rebox_Data_FoldableWithIndex_3725484264_2491554675(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_FoldableWithIndex_foldableWithIndexArray()))), gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_0))
		}), Get_Data_FoldableWithIndex_foldlWithIndexImpl(), Get_Data_FoldableWithIndex_foldrWithIndexImpl()})))}
	})
	return cache_Data_FoldableWithIndex_foldableWithIndexArray
}

var cache_Data_FoldableWithIndex_foldMapWithIndexDefaultL gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldMapWithIndexDefaultL sync.Once

func Get_Data_FoldableWithIndex_foldMapWithIndexDefaultL() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldMapWithIndexDefaultL.Do(func() {
		cache_Data_FoldableWithIndex_foldMapWithIndexDefaultL = gopurs_runtime.Func2(func(dictFoldableWithIndex_0_box gopurs_runtime.Value, dictMonoid_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_FoldableWithIndex_foldMapWithIndexDefaultL(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dictFoldableWithIndex_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_1_box))
		})
	})
	return cache_Data_FoldableWithIndex_foldMapWithIndexDefaultL
}

var cache_Data_FoldableWithIndex_foldMapWithIndex gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldMapWithIndex sync.Once

func Get_Data_FoldableWithIndex_foldMapWithIndex() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldMapWithIndex.Do(func() {
		cache_Data_FoldableWithIndex_foldMapWithIndex = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_FoldableWithIndex_foldMapWithIndex(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Data_FoldableWithIndex_foldMapWithIndex
}

var cache_Data_FoldableWithIndex_foldableWithIndexApp gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldableWithIndexApp sync.Once

func Get_Data_FoldableWithIndex_foldableWithIndexApp() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldableWithIndexApp.Do(func() {
		cache_Data_FoldableWithIndex_foldableWithIndexApp = gopurs_runtime.Func(func(dictFoldableWithIndex_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_FoldableWithIndex_foldableWithIndexApp(dictFoldableWithIndex_0_box)
		})
	})
	return cache_Data_FoldableWithIndex_foldableWithIndexApp
}

var cache_Data_FoldableWithIndex_foldableWithIndexCompose gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldableWithIndexCompose sync.Once

func Get_Data_FoldableWithIndex_foldableWithIndexCompose() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldableWithIndexCompose.Do(func() {
		cache_Data_FoldableWithIndex_foldableWithIndexCompose = gopurs_runtime.Func(func(dictFoldableWithIndex_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_FoldableWithIndex_foldableWithIndexCompose(dictFoldableWithIndex_0_box)
		})
	})
	return cache_Data_FoldableWithIndex_foldableWithIndexCompose
}

var cache_Data_FoldableWithIndex_foldableWithIndexCoproduct gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldableWithIndexCoproduct sync.Once

func Get_Data_FoldableWithIndex_foldableWithIndexCoproduct() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldableWithIndexCoproduct.Do(func() {
		cache_Data_FoldableWithIndex_foldableWithIndexCoproduct = gopurs_runtime.Func(func(dictFoldableWithIndex_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_FoldableWithIndex_foldableWithIndexCoproduct(dictFoldableWithIndex_0_box)
		})
	})
	return cache_Data_FoldableWithIndex_foldableWithIndexCoproduct
}

var cache_Data_FoldableWithIndex_foldableWithIndexProduct gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldableWithIndexProduct sync.Once

func Get_Data_FoldableWithIndex_foldableWithIndexProduct() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldableWithIndexProduct.Do(func() {
		cache_Data_FoldableWithIndex_foldableWithIndexProduct = gopurs_runtime.Func(func(dictFoldableWithIndex_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_FoldableWithIndex_foldableWithIndexProduct(dictFoldableWithIndex_0_box)
		})
	})
	return cache_Data_FoldableWithIndex_foldableWithIndexProduct
}

var cache_Data_FoldableWithIndex_foldlWithIndexDefault gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldlWithIndexDefault sync.Once

func Get_Data_FoldableWithIndex_foldlWithIndexDefault() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldlWithIndexDefault.Do(func() {
		cache_Data_FoldableWithIndex_foldlWithIndexDefault = gopurs_runtime.Func(func(dictFoldableWithIndex_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_FoldableWithIndex_foldlWithIndexDefault(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dictFoldableWithIndex_0_box))
		})
	})
	return cache_Data_FoldableWithIndex_foldlWithIndexDefault
}

var cache_Data_FoldableWithIndex_foldrWithIndexDefault gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldrWithIndexDefault sync.Once

func Get_Data_FoldableWithIndex_foldrWithIndexDefault() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldrWithIndexDefault.Do(func() {
		cache_Data_FoldableWithIndex_foldrWithIndexDefault = gopurs_runtime.Func(func(dictFoldableWithIndex_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_FoldableWithIndex_foldrWithIndexDefault(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dictFoldableWithIndex_0_box))
		})
	})
	return cache_Data_FoldableWithIndex_foldrWithIndexDefault
}

var cache_Data_FoldableWithIndex_surroundMapWithIndex gopurs_runtime.Value
var once_Data_FoldableWithIndex_surroundMapWithIndex sync.Once

func Get_Data_FoldableWithIndex_surroundMapWithIndex() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_surroundMapWithIndex.Do(func() {
		cache_Data_FoldableWithIndex_surroundMapWithIndex = gopurs_runtime.Func(func(dictFoldableWithIndex_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_FoldableWithIndex_surroundMapWithIndex(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dictFoldableWithIndex_0_box))
		})
	})
	return cache_Data_FoldableWithIndex_surroundMapWithIndex
}

var cache_Data_FoldableWithIndex_foldMapDefault gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldMapDefault sync.Once

func Get_Data_FoldableWithIndex_foldMapDefault() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldMapDefault.Do(func() {
		cache_Data_FoldableWithIndex_foldMapDefault = gopurs_runtime.Func2(func(dictFoldableWithIndex_0_box gopurs_runtime.Value, dictMonoid_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_FoldableWithIndex_foldMapDefault(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dictFoldableWithIndex_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_1_box))
		})
	})
	return cache_Data_FoldableWithIndex_foldMapDefault
}

var cache_Data_FoldableWithIndex_findWithIndex gopurs_runtime.Value
var once_Data_FoldableWithIndex_findWithIndex sync.Once

func Get_Data_FoldableWithIndex_findWithIndex() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_findWithIndex.Do(func() {
		cache_Data_FoldableWithIndex_findWithIndex = gopurs_runtime.Func2(func(dictFoldableWithIndex_0_box gopurs_runtime.Value, p_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_FoldableWithIndex_findWithIndex(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dictFoldableWithIndex_0_box), p_1_box)
		})
	})
	return cache_Data_FoldableWithIndex_findWithIndex
}

var cache_Data_FoldableWithIndex_findMapWithIndex gopurs_runtime.Value
var once_Data_FoldableWithIndex_findMapWithIndex sync.Once

func Get_Data_FoldableWithIndex_findMapWithIndex() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_findMapWithIndex.Do(func() {
		cache_Data_FoldableWithIndex_findMapWithIndex = gopurs_runtime.Func2(func(dictFoldableWithIndex_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_FoldableWithIndex_findMapWithIndex(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dictFoldableWithIndex_0_box), f_1_box)
		})
	})
	return cache_Data_FoldableWithIndex_findMapWithIndex
}

var cache_Data_FoldableWithIndex_anyWithIndex gopurs_runtime.Value
var once_Data_FoldableWithIndex_anyWithIndex sync.Once

func Get_Data_FoldableWithIndex_anyWithIndex() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_anyWithIndex.Do(func() {
		cache_Data_FoldableWithIndex_anyWithIndex = gopurs_runtime.Func2(func(dictFoldableWithIndex_0_box gopurs_runtime.Value, dictHeytingAlgebra_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_FoldableWithIndex_anyWithIndex(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dictFoldableWithIndex_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_HeytingAlgebra_HeytingAlgebra[gopurs_runtime.Value]](dictHeytingAlgebra_1_box))
		})
	})
	return cache_Data_FoldableWithIndex_anyWithIndex
}

var cache_Data_FoldableWithIndex_allWithIndex gopurs_runtime.Value
var once_Data_FoldableWithIndex_allWithIndex sync.Once

func Get_Data_FoldableWithIndex_allWithIndex() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_allWithIndex.Do(func() {
		cache_Data_FoldableWithIndex_allWithIndex = gopurs_runtime.Func2(func(dictFoldableWithIndex_0_box gopurs_runtime.Value, dictHeytingAlgebra_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_FoldableWithIndex_allWithIndex(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dictFoldableWithIndex_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_HeytingAlgebra_HeytingAlgebra[gopurs_runtime.Value]](dictHeytingAlgebra_1_box))
		})
	})
	return cache_Data_FoldableWithIndex_allWithIndex
}

var cache_Data_FoldableWithIndex___cached_dict_0 gopurs_runtime.Value
var once_Data_FoldableWithIndex___cached_dict_0 sync.Once

func Get_Data_FoldableWithIndex___cached_dict_0() gopurs_runtime.Value {
	once_Data_FoldableWithIndex___cached_dict_0.Do(func() {
		cache_Data_FoldableWithIndex___cached_dict_0 = gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](Call_Data_Monoid_Dual_monoidDual(Call_Data_Monoid_Endo_monoidEndo(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))}))))}
	})
	return cache_Data_FoldableWithIndex___cached_dict_0
}

var cache_Data_FoldableWithIndex___cached_dict_1 gopurs_runtime.Value
var once_Data_FoldableWithIndex___cached_dict_1 sync.Once

func Get_Data_FoldableWithIndex___cached_dict_1() gopurs_runtime.Value {
	once_Data_FoldableWithIndex___cached_dict_1.Do(func() {
		cache_Data_FoldableWithIndex___cached_dict_1 = gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](Call_Data_Monoid_Endo_monoidEndo(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))})))}
	})
	return cache_Data_FoldableWithIndex___cached_dict_1
}

var cache_Data_FoldableWithIndex___cached_dict_2 gopurs_runtime.Value
var once_Data_FoldableWithIndex___cached_dict_2 sync.Once

func Get_Data_FoldableWithIndex___cached_dict_2() gopurs_runtime.Value {
	once_Data_FoldableWithIndex___cached_dict_2.Do(func() {
		cache_Data_FoldableWithIndex___cached_dict_2 = gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](Call_Data_Monoid_Endo_monoidEndo(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))})))}
	})
	return cache_Data_FoldableWithIndex___cached_dict_2
}

type Constructor_Data_FoldableWithIndex_FoldableWithIndex[T_i any, T_f any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
	V3 gopurs_runtime.Value
}

//gopurs:init-v1:00670065006e006500720061007400650064002f0044006100740061005f0046006f006c006400610062006c006500570069007400680049006e006400650078:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f0044006100740061005f0046006f006c006400610062006c006500570069007400680049006e006400650078005f0046006f006c006400610062006c006500570069007400680049006e006400650078
func init() {
	gopurs_runtime.StructGetters[74250362] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "Foldable0":
			return gopurs_runtime.Box(c.V0)
		case "foldMapWithIndex":
			return gopurs_runtime.Box(c.V1)
		case "foldlWithIndex":
			return gopurs_runtime.Box(c.V2)
		case "foldrWithIndex":
			return gopurs_runtime.Box(c.V3)
		default:
			panic("Key not found in dictionary Constructor_Data_FoldableWithIndex_FoldableWithIndex: " + key)
		}
	}
}

func Call_Data_FoldableWithIndex_FoldableWithIndex_dollar_Dict(x_0_loop struct {
	Foldable0        gopurs_runtime.Value
	foldMapWithIndex gopurs_runtime.Value
	foldlWithIndex   gopurs_runtime.Value
	foldrWithIndex   gopurs_runtime.Value
	/*gopurs:record-v1:4:0046006f006c006400610062006c00650030,0066006f006c0064004d0061007000570069007400680049006e006400650078,0066006f006c0064006c00570069007400680049006e006400650078,0066006f006c0064007200570069007400680049006e006400650078*/
}) *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] {
	var x_0 struct {
		Foldable0        gopurs_runtime.Value
		foldMapWithIndex gopurs_runtime.Value
		foldlWithIndex   gopurs_runtime.Value
		foldrWithIndex   gopurs_runtime.Value
		/*gopurs:record-v1:4:0046006f006c006400610062006c00650030,0066006f006c0064004d0061007000570069007400680049006e006400650078,0066006f006c0064006c00570069007400680049006e006400650078,0066006f006c0064007200570069007400680049006e006400650078*/
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", orig.Foldable0, orig.foldMapWithIndex, orig.foldlWithIndex, orig.foldrWithIndex)
	}())
}

func Call_Data_FoldableWithIndex_foldrWithIndex(dict_0_loop *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V3
}

func Call_Data_FoldableWithIndex_traverseWithIndex_(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictApplicative_0 *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
	_ = dictApplicative_0
	// TAST (Let): applySecond_1_0 shape=App(Var) bindingType=(Func [(TypeApp (TypeVar m$scope13) [(TypeVar b$scope11)]), (TypeApp (TypeVar m$scope13) [Unit])] (TypeApp (TypeVar m$scope13) [Unit]))
	applySecond_1_0 := gopurs_runtime.Apply(Get_Control_Apply_applySecond(), gopurs_runtime.Apply(dictApplicative_0.V0, gopurs_runtime.Value{}))
	_ = applySecond_1_0
	return gopurs_runtime.Func2(func(dictFoldableWithIndex_2 gopurs_runtime.Value, f_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldableWithIndex_2, "foldrWithIndex"), gopurs_runtime.Func(func(i_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), applySecond_1_0, gopurs_runtime.Apply(f_3, i_4))
		}), gopurs_runtime.Apply(dictApplicative_0.V1, Get_Data_Unit_unit()))
	})
}

func Call_Data_FoldableWithIndex_forWithIndex_(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictApplicative_0 *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
	_ = dictApplicative_0
	// TAST (Let): traverseWithIndex_1_1_0 shape=App(Var) bindingType=Any
	traverseWithIndex_1_1_0 := Call_Data_FoldableWithIndex_traverseWithIndex_(dictApplicative_0)
	_ = traverseWithIndex_1_1_0
	return gopurs_runtime.Func(func(dictFoldableWithIndex_2 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_3_1 shape=App(Other) bindingType=Any
		__local_var_3_1 := gopurs_runtime.Apply(traverseWithIndex_1_1_0, dictFoldableWithIndex_2)
		_ = __local_var_3_1
		return gopurs_runtime.Func2(func(b_4 gopurs_runtime.Value, a_5 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(__local_var_3_1, a_5, b_4)
		})
	})
}

func Call_Data_FoldableWithIndex_foldrDefault(dictFoldableWithIndex_0_loop *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value], f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictFoldableWithIndex_0 *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dictFoldableWithIndex_0_loop
	_ = dictFoldableWithIndex_0
	var f_1 gopurs_runtime.Value = f_1_loop
	_ = f_1
	return gopurs_runtime.Apply(dictFoldableWithIndex_0.V3, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return f_1
	}))
}

func Call_Data_FoldableWithIndex_foldlWithIndex(dict_0_loop *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V2
}

func Call_Data_FoldableWithIndex_foldlDefault(dictFoldableWithIndex_0_loop *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value], f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictFoldableWithIndex_0 *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dictFoldableWithIndex_0_loop
	_ = dictFoldableWithIndex_0
	var f_1 gopurs_runtime.Value = f_1_loop
	_ = f_1
	return gopurs_runtime.Apply(dictFoldableWithIndex_0.V2, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return f_1
	}))
}

func Call_Data_FoldableWithIndex_foldWithIndexM(dictFoldableWithIndex_0_loop *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value], dictMonad_1_loop *Constructor_Control_Monad_Monad[gopurs_runtime.Value], f_2_loop gopurs_runtime.Value, a0_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictFoldableWithIndex_0 *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dictFoldableWithIndex_0_loop
	_ = dictFoldableWithIndex_0
	var dictMonad_1 *Constructor_Control_Monad_Monad[gopurs_runtime.Value] = dictMonad_1_loop
	_ = dictMonad_1
	var f_2 gopurs_runtime.Value = f_2_loop
	_ = f_2
	var a0_3 gopurs_runtime.Value = a0_3_loop
	_ = a0_3
	return gopurs_runtime.Apply2(dictFoldableWithIndex_0.V2, gopurs_runtime.Func3(func(i_4 gopurs_runtime.Value, ma_5 gopurs_runtime.Value, b_6 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_7_0 shape=App(Other) bindingType=(Func [(TypeVar a$scope227), (TypeVar b$scope228)] (TypeApp (TypeVar m$scope226) [(TypeVar a$scope227)]))
		__local_var_7_0 := gopurs_runtime.Apply(f_2, i_4)
		_ = __local_var_7_0
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictMonad_1.V1, gopurs_runtime.Value{}), "bind"), ma_5, gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(__local_var_7_0, a_8, b_6)
		}))
	}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictMonad_1.V0, gopurs_runtime.Value{}), "pure"), a0_3))
}

func Call_Data_FoldableWithIndex_foldMapWithIndexDefaultR(dictFoldableWithIndex_0_loop *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value], dictMonoid_1_loop *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictFoldableWithIndex_0 *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dictFoldableWithIndex_0_loop
	_ = dictFoldableWithIndex_0
	var dictMonoid_1 *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value] = dictMonoid_1_loop
	_ = dictMonoid_1
	// TAST (Let): mempty_2_0 shape=App(Var) bindingType=(TypeVar m$scope236)
	mempty_2_0 := Call_Data_Monoid_mempty(gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(dictMonoid_1)})
	_ = mempty_2_0
	return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(dictFoldableWithIndex_0.V3, gopurs_runtime.Func3(func(i_4 gopurs_runtime.Value, x_5 gopurs_runtime.Value, acc_6 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictMonoid_1.V0, gopurs_runtime.Value{}), "append"), gopurs_runtime.Apply2(f_3, i_4, x_5), acc_6)
		}), mempty_2_0)
	})
}

func Call_Data_FoldableWithIndex_foldMapWithIndexDefaultL(dictFoldableWithIndex_0_loop *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value], dictMonoid_1_loop *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictFoldableWithIndex_0 *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dictFoldableWithIndex_0_loop
	_ = dictFoldableWithIndex_0
	var dictMonoid_1 *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value] = dictMonoid_1_loop
	_ = dictMonoid_1
	// TAST (Let): mempty_2_0 shape=App(Var) bindingType=(TypeVar m$scope262)
	mempty_2_0 := Call_Data_Monoid_mempty(gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(dictMonoid_1)})
	_ = mempty_2_0
	return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(dictFoldableWithIndex_0.V2, gopurs_runtime.Func3(func(i_4 gopurs_runtime.Value, acc_5 gopurs_runtime.Value, x_6 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictMonoid_1.V0, gopurs_runtime.Value{}), "append"), acc_5, gopurs_runtime.Apply2(f_3, i_4, x_6))
		}), mempty_2_0)
	})
}

func Call_Data_FoldableWithIndex_foldMapWithIndex(dict_0_loop *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V1
}

func Call_Data_FoldableWithIndex_foldableWithIndexApp(dictFoldableWithIndex_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictFoldableWithIndex_0 gopurs_runtime.Value = dictFoldableWithIndex_0_loop
	_ = dictFoldableWithIndex_0
	// TAST (Let): foldableApp_1_0 shape=App(Var) bindingType=(ADT ["Data","Foldable","Foldable"] [(TypeApp (TypeApp (TypeVar f) [(TypeVar a)]) [(TypeVar f$scope270)])])
	foldableApp_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Call_Data_Foldable_foldableApp(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "Foldable0"), gopurs_runtime.Value{})))
	_ = foldableApp_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer((&Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(foldableApp_1_0)}
	}), gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): foldMapWithIndex2_3_1 shape=App(Other) bindingType=(Func [(Func [(TypeVar a$scope269), (TypeVar a$scope283)] (TypeVar m$scope284)), (TypeApp (TypeVar f$scope270) [(TypeVar a$scope283)])] (TypeVar m$scope284))
		foldMapWithIndex2_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "foldMapWithIndex"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_2))})
		_ = foldMapWithIndex2_3_1
		return gopurs_runtime.Func2(func(f_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(foldMapWithIndex2_3_1, f_4, v_5)
		})
	}), gopurs_runtime.Func3(func(f_2 gopurs_runtime.Value, z_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "foldlWithIndex"), f_2, z_3, v_4)
	}), gopurs_runtime.Func3(func(f_2 gopurs_runtime.Value, z_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "foldrWithIndex"), f_2, z_3, v_4)
	})}))}
}

func Call_Data_FoldableWithIndex_foldableWithIndexCompose(dictFoldableWithIndex_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictFoldableWithIndex_0 gopurs_runtime.Value = dictFoldableWithIndex_0_loop
	_ = dictFoldableWithIndex_0
	// TAST (Let): foldableCompose_1_0 shape=App(Var) bindingType=Any
	foldableCompose_1_0 := gopurs_runtime.Apply(Get_Data_Foldable_foldableCompose(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "Foldable0"), gopurs_runtime.Value{}))
	_ = foldableCompose_1_0
	return gopurs_runtime.Func(func(dictFoldableWithIndex1_2 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): foldlWithIndex2_3_1 shape=App(Var) bindingType=(Func [(Func [(TypeVar b$scope291), (TypeVar b$scope304), (TypeVar a$scope303)] (TypeVar b$scope304)), (TypeVar b$scope304), (TypeApp (TypeVar g$scope292) [(TypeVar a$scope303)])] (TypeVar b$scope304))
		foldlWithIndex2_3_1 := Call_Data_FoldableWithIndex_foldlWithIndex(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dictFoldableWithIndex1_2))
		_ = foldlWithIndex2_3_1
		// TAST (Let): foldMapWithIndex2_4_2 shape=App(Var) bindingType=Any
		foldMapWithIndex2_4_2 := Call_Data_FoldableWithIndex_foldMapWithIndex(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dictFoldableWithIndex1_2))
		_ = foldMapWithIndex2_4_2
		// TAST (Let): foldableCompose1_5_3 shape=App(Other) bindingType=(ADT ["Data","Foldable","Foldable"] [(TypeApp (TypeApp (TypeVar f) [(TypeApp (TypeVar g) [(TypeVar a)])]) [(TypeVar f$scope290), (TypeVar g$scope292)])])
		foldableCompose1_5_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](gopurs_runtime.Apply(foldableCompose_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldableWithIndex1_2, "Foldable0"), gopurs_runtime.Value{})))
		_ = foldableCompose1_5_3
		return gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(Rebox_Data_FoldableWithIndex_139552293_3725484264((&Constructor_Data_FoldableWithIndex_FoldableWithIndex[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value], gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(foldableCompose1_5_3)}
		}), gopurs_runtime.Func(func(dictMonoid_6 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): foldMapWithIndex3_7_4 shape=App(Other) bindingType=(Func [(Func [(TypeVar a$scope289), (TypeApp (TypeVar g$scope292) [(TypeVar a$scope306)])] (TypeVar m$scope307)), (TypeApp (TypeVar f$scope290) [(TypeApp (TypeVar g$scope292) [(TypeVar a$scope306)])])] (TypeVar m$scope307))
			foldMapWithIndex3_7_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "foldMapWithIndex"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_6))})
			_ = foldMapWithIndex3_7_4
			// TAST (Let): foldMapWithIndex4_8_5 shape=App(Other) bindingType=(Func [(Func [(TypeVar b$scope291), (TypeVar a$scope306)] (TypeVar m$scope307)), (TypeApp (TypeVar g$scope292) [(TypeVar a$scope306)])] (TypeVar m$scope307))
			foldMapWithIndex4_8_5 := gopurs_runtime.Apply(foldMapWithIndex2_4_2, dictMonoid_6)
			_ = foldMapWithIndex4_8_5
			return gopurs_runtime.Func2(func(f_9 gopurs_runtime.Value, v_10 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(foldMapWithIndex3_7_4, gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), foldMapWithIndex4_8_5, gopurs_runtime.Func2(func(a_11 gopurs_runtime.Value, b_12 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply(f_9, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
						_v := struct {
							V0 gopurs_runtime.Value
							V1 gopurs_runtime.Value
						}{a_11, b_12}
						return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
					}()))})
				})), v_10)
			})
		}), gopurs_runtime.Func3(func(f_6 gopurs_runtime.Value, i_7 gopurs_runtime.Value, v_8 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "foldlWithIndex"), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), foldlWithIndex2_3_1, gopurs_runtime.Func2(func(a_9 gopurs_runtime.Value, b_10 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply(f_6, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
					_v := struct {
						V0 gopurs_runtime.Value
						V1 gopurs_runtime.Value
					}{a_9, b_10}
					return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
				}()))})
			})), i_7, v_8)
		}), gopurs_runtime.Func3(func(f_6 gopurs_runtime.Value, i_7 gopurs_runtime.Value, v_8 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "foldrWithIndex"), gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
				// TAST (Let): __local_var_10_6 shape=App(Other) bindingType=(Func [(TypeVar b), (TypeApp (TypeVar g$scope292) [(TypeVar a)])] (TypeVar b))
				__local_var_10_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldableWithIndex1_2, "foldrWithIndex"), gopurs_runtime.Func(func(b_10 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply(f_6, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
						_v := struct {
							V0 gopurs_runtime.Value
							V1 gopurs_runtime.Value
						}{a_9, b_10}
						return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
					}()))})
				}))
				_ = __local_var_10_6
				return gopurs_runtime.Func2(func(b_11 gopurs_runtime.Value, a_12 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply2(__local_var_10_6, a_12, b_11)
				})
			}), i_7, v_8)
		})})))}
	})
}

func Call_Data_FoldableWithIndex_foldableWithIndexCoproduct(dictFoldableWithIndex_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictFoldableWithIndex_0 gopurs_runtime.Value = dictFoldableWithIndex_0_loop
	_ = dictFoldableWithIndex_0
	// TAST (Let): foldableCoproduct_1_0 shape=App(Var) bindingType=Any
	foldableCoproduct_1_0 := gopurs_runtime.Apply(Get_Data_Foldable_foldableCoproduct(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "Foldable0"), gopurs_runtime.Value{}))
	_ = foldableCoproduct_1_0
	return gopurs_runtime.Func(func(dictFoldableWithIndex1_2 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): foldableCoproduct1_3_1 shape=App(Other) bindingType=(ADT ["Data","Foldable","Foldable"] [(TypeApp (ADT ["Data","Either","Either"] [(TypeApp (TypeVar f) [(TypeVar a)]), (TypeApp (TypeVar g) [(TypeVar a)])]) [(TypeVar f$scope314), (TypeVar g$scope316)])])
		foldableCoproduct1_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](gopurs_runtime.Apply(foldableCoproduct_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldableWithIndex1_2, "Foldable0"), gopurs_runtime.Value{})))
		_ = foldableCoproduct1_3_1
		return gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer((&Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(foldableCoproduct1_3_1)}
		}), gopurs_runtime.Func(func(dictMonoid_4 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): foldMapWithIndex3_5_2 shape=App(Other) bindingType=(Func [(Func [(TypeVar a$scope313), (TypeVar a$scope343)] (TypeVar m$scope344)), (TypeApp (TypeVar f$scope314) [(TypeVar a$scope343)])] (TypeVar m$scope344))
			foldMapWithIndex3_5_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "foldMapWithIndex"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_4))})
			_ = foldMapWithIndex3_5_2
			// TAST (Let): foldMapWithIndex4_6_3 shape=App(Other) bindingType=(Func [(Func [(TypeVar b$scope315), (TypeVar a$scope343)] (TypeVar m$scope344)), (TypeApp (TypeVar g$scope316) [(TypeVar a$scope343)])] (TypeVar m$scope344))
			foldMapWithIndex4_6_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldableWithIndex1_2, "foldMapWithIndex"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_4))})
			_ = foldMapWithIndex4_6_3
			return gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(Get_Data_Functor_Coproduct_coproduct(), gopurs_runtime.Apply(foldMapWithIndex3_5_2, gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), f_7, Get_Data_Either_Left())), gopurs_runtime.Apply(foldMapWithIndex4_6_3, gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), f_7, Get_Data_Either_Right())))
			})
		}), gopurs_runtime.Func2(func(f_4 gopurs_runtime.Value, z_5 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(Get_Data_Functor_Coproduct_coproduct(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "foldlWithIndex"), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), f_4, Get_Data_Either_Left()), z_5), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldableWithIndex1_2, "foldlWithIndex"), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), f_4, Get_Data_Either_Right()), z_5))
		}), gopurs_runtime.Func2(func(f_4 gopurs_runtime.Value, z_5 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(Get_Data_Functor_Coproduct_coproduct(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "foldrWithIndex"), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), f_4, Get_Data_Either_Left()), z_5), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldableWithIndex1_2, "foldrWithIndex"), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), f_4, Get_Data_Either_Right()), z_5))
		})}))}
	})
}

func Call_Data_FoldableWithIndex_foldableWithIndexProduct(dictFoldableWithIndex_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictFoldableWithIndex_0 gopurs_runtime.Value = dictFoldableWithIndex_0_loop
	_ = dictFoldableWithIndex_0
	// TAST (Let): foldableProduct_1_0 shape=App(Var) bindingType=Any
	foldableProduct_1_0 := gopurs_runtime.Apply(Get_Data_Foldable_foldableProduct(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "Foldable0"), gopurs_runtime.Value{}))
	_ = foldableProduct_1_0
	return gopurs_runtime.Func(func(dictFoldableWithIndex1_2 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): foldableProduct1_3_1 shape=App(Other) bindingType=(ADT ["Data","Foldable","Foldable"] [(TypeApp (ADT ["Data","Tuple","Tuple"] [(TypeApp (TypeVar f) [(TypeVar a)]), (TypeApp (TypeVar g) [(TypeVar a)])]) [(TypeVar f$scope357), (TypeVar g$scope359)])])
		foldableProduct1_3_1 := Rebox_Data_FoldableWithIndex_1680800814_4173511203(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](gopurs_runtime.Apply(foldableProduct_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldableWithIndex1_2, "Foldable0"), gopurs_runtime.Value{}))))
		_ = foldableProduct1_3_1
		return gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(Rebox_Data_FoldableWithIndex_3374046885_3725484264((&Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(Rebox_Data_FoldableWithIndex_4173511203_1680800814(foldableProduct1_3_1))}
		}), gopurs_runtime.Func(func(dictMonoid_4 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): foldMapWithIndex3_5_2 shape=App(Other) bindingType=(Func [(Func [(TypeVar a$scope356), (TypeVar a$scope384)] (TypeVar m$scope385)), (TypeApp (TypeVar f$scope357) [(TypeVar a$scope384)])] (TypeVar m$scope385))
			foldMapWithIndex3_5_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "foldMapWithIndex"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_4))})
			_ = foldMapWithIndex3_5_2
			// TAST (Let): foldMapWithIndex4_6_3 shape=App(Other) bindingType=(Func [(Func [(TypeVar b$scope358), (TypeVar a$scope384)] (TypeVar m$scope385)), (TypeApp (TypeVar g$scope359) [(TypeVar a$scope384)])] (TypeVar m$scope385))
			foldMapWithIndex4_6_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldableWithIndex1_2, "foldMapWithIndex"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_4))})
			_ = foldMapWithIndex4_6_3
			return gopurs_runtime.Func2(func(f_7 gopurs_runtime.Value, v_8 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_4, "Semigroup0"), gopurs_runtime.Value{}), "append"), gopurs_runtime.Apply2(foldMapWithIndex3_5_2, gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), f_7, Get_Data_Either_Left()), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V0), gopurs_runtime.Apply2(foldMapWithIndex4_6_3, gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), f_7, Get_Data_Either_Right()), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V1))
			})
		}), gopurs_runtime.Func3(func(f_4 gopurs_runtime.Value, z_5 gopurs_runtime.Value, v_6 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldableWithIndex1_2, "foldlWithIndex"), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), f_4, Get_Data_Either_Right()), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "foldlWithIndex"), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), f_4, Get_Data_Either_Left()), z_5, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V1)
		}), gopurs_runtime.Func3(func(f_4 gopurs_runtime.Value, z_5 gopurs_runtime.Value, v_6 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "foldrWithIndex"), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), f_4, Get_Data_Either_Left()), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldableWithIndex1_2, "foldrWithIndex"), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), f_4, Get_Data_Either_Right()), z_5, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V1), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V0)
		})})))}
	})
}

func Call_Data_FoldableWithIndex_foldlWithIndexDefault(dictFoldableWithIndex_0_loop *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictFoldableWithIndex_0 *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dictFoldableWithIndex_0_loop
	_ = dictFoldableWithIndex_0
	// TAST (Let): foldMapWithIndex1_1_0 shape=App(Other) bindingType=(Func [(Func [(TypeVar i$scope396), (TypeVar a$scope398)] (TypeApp (TypeApp (TypeVar c) [(TypeVar a), (TypeVar a)]) [(ADT ["Prim","Function"] []), (TypeVar b$scope399)])), (TypeApp (TypeVar f$scope397) [(TypeVar a$scope398)])] (TypeApp (TypeApp (TypeVar c) [(TypeVar a), (TypeVar a)]) [(ADT ["Prim","Function"] []), (TypeVar b$scope399)]))
	foldMapWithIndex1_1_0 := gopurs_runtime.Apply(dictFoldableWithIndex_0.V1, gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](Get_Data_FoldableWithIndex___cached_dict_0()))})
	_ = foldMapWithIndex1_1_0
	return gopurs_runtime.Func3(func(c_2 gopurs_runtime.Value, u_3 gopurs_runtime.Value, xs_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(Call_Safe_Coerce_coerce(gopurs_runtime.Value{}), gopurs_runtime.Apply(Call_Safe_Coerce_coerce(gopurs_runtime.Value{}), gopurs_runtime.Apply2(foldMapWithIndex1_1_0, gopurs_runtime.Func(func(i_5 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_6_1 shape=App(Other) bindingType=(Func [(TypeVar b$scope399), (TypeVar a$scope398)] (TypeVar b$scope399))
			__local_var_6_1 := gopurs_runtime.Apply(c_2, i_5)
			_ = __local_var_6_1
			return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Monoid_Dual_Dual(), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Monoid_Endo_Endo(), gopurs_runtime.Func2(func(b_7 gopurs_runtime.Value, a_8 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(__local_var_6_1, a_8, b_7)
			})))
		}), xs_4)), u_3)
	})
}

func Call_Data_FoldableWithIndex_foldrWithIndexDefault(dictFoldableWithIndex_0_loop *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictFoldableWithIndex_0 *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dictFoldableWithIndex_0_loop
	_ = dictFoldableWithIndex_0
	// TAST (Let): foldMapWithIndex1_1_0 shape=App(Other) bindingType=(Func [(Func [(TypeVar i$scope410), (TypeVar a$scope412)] (TypeApp (TypeApp (TypeVar c) [(TypeVar a), (TypeVar a)]) [(ADT ["Prim","Function"] []), (TypeVar b$scope413)])), (TypeApp (TypeVar f$scope411) [(TypeVar a$scope412)])] (TypeApp (TypeApp (TypeVar c) [(TypeVar a), (TypeVar a)]) [(ADT ["Prim","Function"] []), (TypeVar b$scope413)]))
	foldMapWithIndex1_1_0 := gopurs_runtime.Apply(dictFoldableWithIndex_0.V1, gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](Get_Data_FoldableWithIndex___cached_dict_1()))})
	_ = foldMapWithIndex1_1_0
	return gopurs_runtime.Func3(func(c_2 gopurs_runtime.Value, u_3 gopurs_runtime.Value, xs_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(Call_Safe_Coerce_coerce(gopurs_runtime.Value{}), gopurs_runtime.Apply2(foldMapWithIndex1_1_0, gopurs_runtime.Func(func(i_5 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Monoid_Endo_Endo(), gopurs_runtime.Apply(c_2, i_5))
		}), xs_4), u_3)
	})
}

func Call_Data_FoldableWithIndex_surroundMapWithIndex(dictFoldableWithIndex_0_loop *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictFoldableWithIndex_0 *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dictFoldableWithIndex_0_loop
	_ = dictFoldableWithIndex_0
	// TAST (Let): foldMapWithIndex1_1_0 shape=App(Other) bindingType=(Func [(Func [(TypeVar i$scope422), (TypeVar a$scope424)] (TypeApp (TypeApp (TypeVar c) [(TypeVar a), (TypeVar a)]) [(ADT ["Prim","Function"] []), (TypeVar m$scope425)])), (TypeApp (TypeVar f$scope423) [(TypeVar a$scope424)])] (TypeApp (TypeApp (TypeVar c) [(TypeVar a), (TypeVar a)]) [(ADT ["Prim","Function"] []), (TypeVar m$scope425)]))
	foldMapWithIndex1_1_0 := gopurs_runtime.Apply(dictFoldableWithIndex_0.V1, gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](Get_Data_FoldableWithIndex___cached_dict_2()))})
	_ = foldMapWithIndex1_1_0
	return gopurs_runtime.Func4(func(dictSemigroup_2 gopurs_runtime.Value, d_3 gopurs_runtime.Value, t_4 gopurs_runtime.Value, f_5 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(Call_Safe_Coerce_coerce(gopurs_runtime.Value{}), gopurs_runtime.Apply2(foldMapWithIndex1_1_0, gopurs_runtime.Func3(func(i_6 gopurs_runtime.Value, a_7 gopurs_runtime.Value, m_8 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_2, "append"), d_3, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_2, "append"), gopurs_runtime.Apply2(t_4, i_6, a_7), m_8))
		}), f_5), d_3)
	})
}

func Call_Data_FoldableWithIndex_foldMapDefault(dictFoldableWithIndex_0_loop *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value], dictMonoid_1_loop *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictFoldableWithIndex_0 *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dictFoldableWithIndex_0_loop
	_ = dictFoldableWithIndex_0
	var dictMonoid_1 *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value] = dictMonoid_1_loop
	_ = dictMonoid_1
	// TAST (Let): foldMapWithIndex2_2_0 shape=App(Other) bindingType=(Func [(Func [Any, (TypeVar a$scope435)] (TypeVar m$scope436)), (TypeApp (TypeVar f$scope434) [(TypeVar a$scope435)])] (TypeVar m$scope436))
	foldMapWithIndex2_2_0 := gopurs_runtime.Apply(dictFoldableWithIndex_0.V1, gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(dictMonoid_1)})
	_ = foldMapWithIndex2_2_0
	return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(foldMapWithIndex2_2_0, gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return f_3
		}))
	})
}

func Call_Data_FoldableWithIndex_findWithIndex(dictFoldableWithIndex_0_loop *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value], p_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictFoldableWithIndex_0 *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dictFoldableWithIndex_0_loop
	_ = dictFoldableWithIndex_0
	var p_1 gopurs_runtime.Value = p_1_loop
	_ = p_1
	return gopurs_runtime.Apply2(dictFoldableWithIndex_0.V2, gopurs_runtime.Func3(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value, v2_4 gopurs_runtime.Value) gopurs_runtime.Value {
		var __t0 gopurs_runtime.Value
		{
			if (v1_3.Type == 9 && v1_3.IntVal == 930809136 && v1_3.UnsafePtr == nil) && ((gopurs_runtime.Apply2(p_1, v_2, v2_4).IntVal) != (0)) {
				__t0 = func() gopurs_runtime.Value {
					_v := struct {
						V0 gopurs_runtime.Value
						V1 bool
					}{gopurs_runtime.RecordDict2("index", "value", v_2, v2_4), true}
					if _v.V1 {
						return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
					}
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
				}()
				goto end_branch_0
			} else {

			}
		}
		{
			__t0 = v1_3
		}
	end_branch_0:
		return __t0
	}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_FoldableWithIndex_516634048_3094389156(Rebox_Data_FoldableWithIndex_3094389156_516634048(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		_v := struct {
			V0 struct {
				index gopurs_runtime.Value
				value gopurs_runtime.Value
				/*gopurs:record-v1:2:0069006e006400650078,00760061006c00750065*/
			}
			V1 bool
		}{struct {
			index gopurs_runtime.Value
			value gopurs_runtime.Value
			/*gopurs:record-v1:2:0069006e006400650078,00760061006c00750065*/
		}{}, false}
		if _v.V1 {
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: func() gopurs_runtime.Value {
				orig := _v.V0
				_ = orig
				return gopurs_runtime.RecordDict2("index", "value", orig.index, orig.value)
			}()})}
		}
		return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
	}()))))})
}

func Call_Data_FoldableWithIndex_findMapWithIndex(dictFoldableWithIndex_0_loop *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value], f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictFoldableWithIndex_0 *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dictFoldableWithIndex_0_loop
	_ = dictFoldableWithIndex_0
	var f_1 gopurs_runtime.Value = f_1_loop
	_ = f_1
	return gopurs_runtime.Apply2(dictFoldableWithIndex_0.V2, gopurs_runtime.Func3(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value, v2_4 gopurs_runtime.Value) gopurs_runtime.Value {
		var __t0 gopurs_runtime.Value
		{
			if v1_3.Type == 9 && v1_3.IntVal == 930809136 && v1_3.UnsafePtr == nil {
				__t0 = gopurs_runtime.Apply2(f_1, v_2, v2_4)
				goto end_branch_0
			} else {

			}
		}
		{
			__t0 = v1_3
		}
	end_branch_0:
		return __t0
	}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		_v := struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{gopurs_runtime.Value{}, false}
		if _v.V1 {
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
		}
		return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
	}()))})
}

func Call_Data_FoldableWithIndex_anyWithIndex(dictFoldableWithIndex_0_loop *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value], dictHeytingAlgebra_1_loop *Constructor_Data_HeytingAlgebra_HeytingAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictFoldableWithIndex_0 *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dictFoldableWithIndex_0_loop
	_ = dictFoldableWithIndex_0
	var dictHeytingAlgebra_1 *Constructor_Data_HeytingAlgebra_HeytingAlgebra[gopurs_runtime.Value] = dictHeytingAlgebra_1_loop
	_ = dictHeytingAlgebra_1
	// TAST (Let): foldMapWithIndex2_2_0 shape=App(Other) bindingType=(Func [(Func [(TypeVar i$scope458), (TypeVar a$scope459)] (TypeVar b$scope460)), (TypeApp (TypeVar f$scope461) [(TypeVar a$scope459)])] (TypeVar b$scope460))
	foldMapWithIndex2_2_0 := gopurs_runtime.Apply(dictFoldableWithIndex_0.V1, gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](Call_Data_Monoid_Disj_monoidDisj(gopurs_runtime.Value{Type: 9, IntVal: 926771738, UnsafePtr: unsafe.Pointer(dictHeytingAlgebra_1)})))})
	_ = foldMapWithIndex2_2_0
	return gopurs_runtime.Func(func(t_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Call_Data_Newtype_unwrap(gopurs_runtime.CoerceToStruct[Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{})), gopurs_runtime.Apply(foldMapWithIndex2_2_0, gopurs_runtime.Func(func(i_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Monoid_Disj_Disj(), gopurs_runtime.Apply(t_3, i_4))
		})))
	})
}

func Call_Data_FoldableWithIndex_allWithIndex(dictFoldableWithIndex_0_loop *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value], dictHeytingAlgebra_1_loop *Constructor_Data_HeytingAlgebra_HeytingAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictFoldableWithIndex_0 *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dictFoldableWithIndex_0_loop
	_ = dictFoldableWithIndex_0
	var dictHeytingAlgebra_1 *Constructor_Data_HeytingAlgebra_HeytingAlgebra[gopurs_runtime.Value] = dictHeytingAlgebra_1_loop
	_ = dictHeytingAlgebra_1
	// TAST (Let): foldMapWithIndex2_2_0 shape=App(Other) bindingType=(Func [(Func [(TypeVar i$scope469), (TypeVar a$scope470)] (TypeVar b$scope471)), (TypeApp (TypeVar f$scope472) [(TypeVar a$scope470)])] (TypeVar b$scope471))
	foldMapWithIndex2_2_0 := gopurs_runtime.Apply(dictFoldableWithIndex_0.V1, gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](Call_Data_Monoid_Conj_monoidConj(gopurs_runtime.Value{Type: 9, IntVal: 926771738, UnsafePtr: unsafe.Pointer(dictHeytingAlgebra_1)})))})
	_ = foldMapWithIndex2_2_0
	return gopurs_runtime.Func(func(t_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Call_Data_Newtype_unwrap(gopurs_runtime.CoerceToStruct[Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{})), gopurs_runtime.Apply(foldMapWithIndex2_2_0, gopurs_runtime.Func(func(i_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Monoid_Conj_Conj(), gopurs_runtime.Apply(t_3, i_4))
		})))
	})
}

func Rebox_Data_FoldableWithIndex_1146820559_1680800814(in *Constructor_Data_Foldable_Foldable[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Foldable_Foldable[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_FoldableWithIndex_139552293_3725484264(in *Constructor_Data_FoldableWithIndex_FoldableWithIndex[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value], gopurs_runtime.Value]) *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_FoldableWithIndex_1680800814_1146820559(in *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]) *Constructor_Data_Foldable_Foldable[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Foldable_Foldable[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Data_FoldableWithIndex_1680800814_4173511203(in *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]) *Constructor_Data_Foldable_Foldable[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Foldable_Foldable[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Data_FoldableWithIndex_2286084809_3725484264(in *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, *Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_FoldableWithIndex_2491554675_3725484264(in *Constructor_Data_FoldableWithIndex_FoldableWithIndex[int64, gopurs_runtime.Value]) *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_FoldableWithIndex_3094389156_516634048(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[struct {
	index gopurs_runtime.Value
	value gopurs_runtime.Value
	/*gopurs:record-v1:2:0069006e006400650078,00760061006c00750065*/
}] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[struct {
		index gopurs_runtime.Value
		value gopurs_runtime.Value
		/*gopurs:record-v1:2:0069006e006400650078,00760061006c00750065*/
	}]{}
	out.V0 = func() struct {
		index gopurs_runtime.Value
		value gopurs_runtime.Value
		/*gopurs:record-v1:2:0069006e006400650078,00760061006c00750065*/
	} {
		orig := in.V0
		_ = orig
		clone := struct {
			index gopurs_runtime.Value
			value gopurs_runtime.Value
			/*gopurs:record-v1:2:0069006e006400650078,00760061006c00750065*/
		}{}
		clone.index = gopurs_runtime.RecordGet(orig, "index")
		clone.value = gopurs_runtime.RecordGet(orig, "value")
		return clone
	}()
	return out
}

func Rebox_Data_FoldableWithIndex_3374046885_3725484264(in *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_FoldableWithIndex_3725484264_2491554675(in *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_FoldableWithIndex_FoldableWithIndex[int64, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_FoldableWithIndex_FoldableWithIndex[int64, gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_FoldableWithIndex_4173511203_1680800814(in *Constructor_Data_Foldable_Foldable[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Foldable_Foldable[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_FoldableWithIndex_516634048_3094389156(in *Constructor_Data_Maybe_Just[struct {
	index gopurs_runtime.Value
	value gopurs_runtime.Value
	/*gopurs:record-v1:2:0069006e006400650078,00760061006c00750065*/
}]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
	out.V0 = func() gopurs_runtime.Value {
		orig := in.V0
		_ = orig
		return gopurs_runtime.RecordDict2("index", "value", orig.index, orig.value)
	}()
	return out
}

func Get_Data_FoldableWithIndex_foldlWithIndexImpl() gopurs_runtime.Value {
	return _Gopurs_Data_FoldableWithIndex_FoldlWithIndexImpl
}

func Get_Data_FoldableWithIndex_foldrWithIndexImpl() gopurs_runtime.Value {
	return _Gopurs_Data_FoldableWithIndex_FoldrWithIndexImpl
}
