package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Predicate_compose gopurs_runtime.Value
var once_Data_Predicate_compose sync.Once

func Get_Data_Predicate_compose() gopurs_runtime.Value {
	once_Data_Predicate_compose.Do(func() {
		cache_Data_Predicate_compose = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Predicate_compose
}

var cache_Data_Predicate_Predicate gopurs_runtime.Value
var once_Data_Predicate_Predicate sync.Once

func Get_Data_Predicate_Predicate() gopurs_runtime.Value {
	once_Data_Predicate_Predicate.Do(func() {
		cache_Data_Predicate_Predicate = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Predicate_Predicate(x_0_box)
		})
	})
	return cache_Data_Predicate_Predicate
}

var cache_Data_Predicate_newtypePredicate gopurs_runtime.Value
var once_Data_Predicate_newtypePredicate sync.Once

func Get_Data_Predicate_newtypePredicate() gopurs_runtime.Value {
	once_Data_Predicate_newtypePredicate.Do(func() {
		cache_Data_Predicate_newtypePredicate = gopurs_runtime.Value{Type: 9, IntVal: 3322196858, UnsafePtr: unsafe.Pointer((&Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{}
		})}))}
	})
	return cache_Data_Predicate_newtypePredicate
}

var cache_Data_Predicate_heytingAlgebraPredicate gopurs_runtime.Value
var once_Data_Predicate_heytingAlgebraPredicate sync.Once

func Get_Data_Predicate_heytingAlgebraPredicate() gopurs_runtime.Value {
	once_Data_Predicate_heytingAlgebraPredicate.Do(func() {
		cache_Data_Predicate_heytingAlgebraPredicate = gopurs_runtime.Value{Type: 9, IntVal: 926771738, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_HeytingAlgebra_HeytingAlgebra[gopurs_runtime.Value]](Call_Data_HeytingAlgebra_heytingAlgebraFunction(gopurs_runtime.Value{Type: 9, IntVal: 926771738, UnsafePtr: unsafe.Pointer(Rebox_Data_Predicate_3591112874_2663347022(Rebox_Data_Predicate_2663347022_3591112874(gopurs_runtime.CoerceToStruct[Constructor_Data_HeytingAlgebra_HeytingAlgebra[gopurs_runtime.Value]](Get_Data_HeytingAlgebra_heytingAlgebraBoolean()))))})))}
	})
	return cache_Data_Predicate_heytingAlgebraPredicate
}

var cache_Data_Predicate_contravariantPredicate gopurs_runtime.Value
var once_Data_Predicate_contravariantPredicate sync.Once

func Get_Data_Predicate_contravariantPredicate() gopurs_runtime.Value {
	once_Data_Predicate_contravariantPredicate.Do(func() {
		cache_Data_Predicate_contravariantPredicate = gopurs_runtime.Value{Type: 9, IntVal: 85171506, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Contravariant_Contravariant[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), v_1, f_0)
		})}))}
	})
	return cache_Data_Predicate_contravariantPredicate
}

var cache_Data_Predicate_booleanAlgebraPredicate gopurs_runtime.Value
var once_Data_Predicate_booleanAlgebraPredicate sync.Once

func Get_Data_Predicate_booleanAlgebraPredicate() gopurs_runtime.Value {
	once_Data_Predicate_booleanAlgebraPredicate.Do(func() {
		cache_Data_Predicate_booleanAlgebraPredicate = gopurs_runtime.Value{Type: 9, IntVal: 3257204378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_BooleanAlgebra_BooleanAlgebra[gopurs_runtime.Value]](Call_Data_BooleanAlgebra_booleanAlgebraFn(gopurs_runtime.Value{Type: 9, IntVal: 3257204378, UnsafePtr: unsafe.Pointer(Rebox_Data_Predicate_1931817002_72969294(Rebox_Data_Predicate_72969294_1931817002(gopurs_runtime.CoerceToStruct[Constructor_Data_BooleanAlgebra_BooleanAlgebra[gopurs_runtime.Value]](Get_Data_BooleanAlgebra_booleanAlgebraBoolean()))))})))}
	})
	return cache_Data_Predicate_booleanAlgebraPredicate
}

func Call_Data_Predicate_Predicate(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Rebox_Data_Predicate_1931817002_72969294(in *Constructor_Data_BooleanAlgebra_BooleanAlgebra[bool]) *Constructor_Data_BooleanAlgebra_BooleanAlgebra[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_BooleanAlgebra_BooleanAlgebra[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Predicate_2663347022_3591112874(in *Constructor_Data_HeytingAlgebra_HeytingAlgebra[gopurs_runtime.Value]) *Constructor_Data_HeytingAlgebra_HeytingAlgebra[bool] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_HeytingAlgebra_HeytingAlgebra[bool]{}
	out.V0 = in.V0
	out.V1 = in.V1
	out.V2 = (in.V2.IntVal) != (0)
	out.V3 = in.V3
	out.V4 = in.V4
	out.V5 = (in.V5.IntVal) != (0)
	return out
}

func Rebox_Data_Predicate_3591112874_2663347022(in *Constructor_Data_HeytingAlgebra_HeytingAlgebra[bool]) *Constructor_Data_HeytingAlgebra_HeytingAlgebra[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_HeytingAlgebra_HeytingAlgebra[gopurs_runtime.Value]{}
	out.V0 = in.V0
	out.V1 = in.V1
	out.V2 = gopurs_runtime.Bool(in.V2)
	out.V3 = in.V3
	out.V4 = in.V4
	out.V5 = gopurs_runtime.Bool(in.V5)
	return out
}

func Rebox_Data_Predicate_72969294_1931817002(in *Constructor_Data_BooleanAlgebra_BooleanAlgebra[gopurs_runtime.Value]) *Constructor_Data_BooleanAlgebra_BooleanAlgebra[bool] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_BooleanAlgebra_BooleanAlgebra[bool])(unsafe.Pointer(in))
}
