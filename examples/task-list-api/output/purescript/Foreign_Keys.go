package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Foreign_Keys_eq gopurs_runtime.Value
var once_Foreign_Keys_eq sync.Once

func Get_Foreign_Keys_eq() gopurs_runtime.Value {
	once_Foreign_Keys_eq.Do(func() {
		cache_Foreign_Keys_eq = Get_Data_Eq_eq__3478410105()
	})
	return cache_Foreign_Keys_eq
}

var cache_Foreign_Keys_keys gopurs_runtime.Value
var once_Foreign_Keys_keys sync.Once

func Get_Foreign_Keys_keys() gopurs_runtime.Value {
	once_Foreign_Keys_keys.Do(func() {
		cache_Foreign_Keys_keys = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Foreign_Keys_keys(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](dictMonad_0_box))
		})
	})
	return cache_Foreign_Keys_keys
}

func Call_Foreign_Keys_keys(dictMonad_0_loop *Constructor_Control_Monad_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictMonad_0 *Constructor_Control_Monad_Monad[gopurs_runtime.Value] = dictMonad_0_loop
	_ = dictMonad_0
	// TAST (Let): fail_1_0 shape=App(Var) bindingType=(Func [(ADT ["Foreign","ForeignError"] [])] (TypeApp (TypeVar m$scope1) [(ADT ["Data","Either","Either"] [(ADT ["Data","NonEmpty","NonEmpty"] [(ADT ["Data","List","Types","List"] []), (ADT ["Foreign","ForeignError"] [])]), (Array String)])]))
	fail_1_0 := Call_Foreign_fail(dictMonad_0)
	_ = fail_1_0
	// TAST (Let): pure_2_1 shape=App(Var) bindingType=(Func [(Array String)] (TypeApp (TypeVar m$scope1) [(ADT ["Data","Either","Either"] [(ADT ["Data","NonEmpty","NonEmpty"] [(ADT ["Data","List","Types","List"] []), (ADT ["Foreign","ForeignError"] [])]), (Array String)])]))
	pure_2_1 := Call_Control_Applicative_pure(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Call_Control_Monad_Except_Trans_applicativeExceptT(gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(dictMonad_0)})))
	_ = pure_2_1
	return gopurs_runtime.Func(func(value_3 gopurs_runtime.Value) gopurs_runtime.Value {
		var __t2 gopurs_runtime.Value
		{
			if (gopurs_runtime.Apply(Get_Foreign_isNull(), value_3).IntVal) != (0) {
				__t2 = gopurs_runtime.Apply(fail_1_0, gopurs_runtime.Value{Type: 9, IntVal: 2536121783, UnsafePtr: unsafe.Pointer((&Constructor_Foreign_TypeMismatch{1, "object", "null"}))})
				goto end_branch_2
			} else {

			}
		}
		{
			if (gopurs_runtime.Apply(Get_Foreign_isUndefined(), value_3).IntVal) != (0) {
				__t2 = gopurs_runtime.Apply(fail_1_0, gopurs_runtime.Value{Type: 9, IntVal: 2536121783, UnsafePtr: unsafe.Pointer((&Constructor_Foreign_TypeMismatch{1, "object", "undefined"}))})
				goto end_branch_2
			} else {

			}
		}
		{
			if (gopurs_runtime.Apply(Get_Foreign_typeOf(), value_3).StrVal()) == ("object") {
				__t2 = gopurs_runtime.Apply(pure_2_1, func() gopurs_runtime.Value {
					arr := func() []string {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(Get_Foreign_Keys_unsafeKeys(), value_3).UnsafePtr)
						unboxed := make([]string, len(arr))
						for i, v := range arr {
							unboxed[i] = v.StrVal()
						}
						return unboxed
					}()
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Str(v)
					}
					return gopurs_runtime.Array(boxed)
				}())
				goto end_branch_2
			} else {

			}
		}
		{
			__t2 = gopurs_runtime.Apply(fail_1_0, gopurs_runtime.Value{Type: 9, IntVal: 2536121783, UnsafePtr: unsafe.Pointer((&Constructor_Foreign_TypeMismatch{1, "object", gopurs_runtime.Apply(Get_Foreign_typeOf(), value_3).StrVal()}))})
		}
	end_branch_2:
		return __t2
	})
}

func Get_Foreign_Keys_unsafeKeys() gopurs_runtime.Value {
	return _Gopurs_Foreign_Keys_UnsafeKeys
}
