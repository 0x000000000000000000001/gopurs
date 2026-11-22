package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Foreign_go__append gopurs_runtime.Value
var once_Foreign_go__append sync.Once

func Get_Foreign_go__append() gopurs_runtime.Value {
	once_Foreign_go__append.Do(func() {
		cache_Foreign_go__append = Get_Data_Semigroup_append__1694605992()
	})
	return cache_Foreign_go__append
}

var cache_Foreign_show gopurs_runtime.Value
var once_Foreign_show sync.Once

func Get_Foreign_show() gopurs_runtime.Value {
	once_Foreign_show.Do(func() {
		cache_Foreign_show = Get_Data_Show_show__3588520802()
	})
	return cache_Foreign_show
}

var cache_Foreign_show1 gopurs_runtime.Value
var once_Foreign_show1 sync.Once

func Get_Foreign_show1() gopurs_runtime.Value {
	once_Foreign_show1.Do(func() {
		cache_Foreign_show1 = Get_Data_Show_show__1986550628()
	})
	return cache_Foreign_show1
}

var cache_Foreign_disj gopurs_runtime.Value
var once_Foreign_disj sync.Once

func Get_Foreign_disj() gopurs_runtime.Value {
	once_Foreign_disj.Do(func() {
		cache_Foreign_disj = Get_Data_HeytingAlgebra_disj__2421783737()
	})
	return cache_Foreign_disj
}

var cache_Foreign_compose gopurs_runtime.Value
var once_Foreign_compose sync.Once

func Get_Foreign_compose() gopurs_runtime.Value {
	once_Foreign_compose.Do(func() {
		cache_Foreign_compose = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Foreign_compose
}

var cache_Foreign_eq gopurs_runtime.Value
var once_Foreign_eq sync.Once

func Get_Foreign_eq() gopurs_runtime.Value {
	once_Foreign_eq.Do(func() {
		cache_Foreign_eq = Get_Data_Eq_eq__3478410105()
	})
	return cache_Foreign_eq
}

var cache_Foreign_compose1 gopurs_runtime.Value
var once_Foreign_compose1 sync.Once

func Get_Foreign_compose1() gopurs_runtime.Value {
	once_Foreign_compose1.Do(func() {
		cache_Foreign_compose1 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Foreign_compose1
}

var cache_Foreign_pure gopurs_runtime.Value
var once_Foreign_pure sync.Once

func Get_Foreign_pure() gopurs_runtime.Value {
	once_Foreign_pure.Do(func() {
		cache_Foreign_pure = Get_Data_Either_Right()
	})
	return cache_Foreign_pure
}

var cache_Foreign_compose2 gopurs_runtime.Value
var once_Foreign_compose2 sync.Once

func Get_Foreign_compose2() gopurs_runtime.Value {
	once_Foreign_compose2.Do(func() {
		cache_Foreign_compose2 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Foreign_compose2
}

var cache_Foreign_pure1 gopurs_runtime.Value
var once_Foreign_pure1 sync.Once

func Get_Foreign_pure1() gopurs_runtime.Value {
	once_Foreign_pure1.Do(func() {
		cache_Foreign_pure1 = Get_Data_Either_Right()
	})
	return cache_Foreign_pure1
}

var cache_Foreign_conj gopurs_runtime.Value
var once_Foreign_conj sync.Once

func Get_Foreign_conj() gopurs_runtime.Value {
	once_Foreign_conj.Do(func() {
		cache_Foreign_conj = Get_Data_HeytingAlgebra_conj__2421783737()
	})
	return cache_Foreign_conj
}

var cache_Foreign_eq1 gopurs_runtime.Value
var once_Foreign_eq1 sync.Once

func Get_Foreign_eq1() gopurs_runtime.Value {
	once_Foreign_eq1.Do(func() {
		cache_Foreign_eq1 = Get_Data_Eq_eq__3020190905()
	})
	return cache_Foreign_eq1
}

var cache_Foreign_compare gopurs_runtime.Value
var once_Foreign_compare sync.Once

func Get_Foreign_compare() gopurs_runtime.Value {
	once_Foreign_compare.Do(func() {
		cache_Foreign_compare = Get_Data_Ord_compare__3009265731()
	})
	return cache_Foreign_compare
}

var cache_Foreign_compare1 gopurs_runtime.Value
var once_Foreign_compare1 sync.Once

func Get_Foreign_compare1() gopurs_runtime.Value {
	once_Foreign_compare1.Do(func() {
		cache_Foreign_compare1 = Get_Data_Ord_compare__1925727043()
	})
	return cache_Foreign_compare1
}

var cache_Foreign_ForeignError gopurs_runtime.Value
var once_Foreign_ForeignError sync.Once

func Get_Foreign_ForeignError() gopurs_runtime.Value {
	once_Foreign_ForeignError.Do(func() {
		cache_Foreign_ForeignError = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2355653293, UnsafePtr: unsafe.Pointer((&Constructor_Foreign_ForeignError{1, value0.StrVal()}))}
		})
	})
	return cache_Foreign_ForeignError
}

var cache_Foreign_TypeMismatch gopurs_runtime.Value
var once_Foreign_TypeMismatch sync.Once

func Get_Foreign_TypeMismatch() gopurs_runtime.Value {
	once_Foreign_TypeMismatch.Do(func() {
		cache_Foreign_TypeMismatch = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 2536121783, UnsafePtr: unsafe.Pointer((&Constructor_Foreign_TypeMismatch{1, value0.StrVal(), value1.StrVal()}))}
			})
		})
	})
	return cache_Foreign_TypeMismatch
}

var cache_Foreign_ErrorAtIndex gopurs_runtime.Value
var once_Foreign_ErrorAtIndex sync.Once

func Get_Foreign_ErrorAtIndex() gopurs_runtime.Value {
	once_Foreign_ErrorAtIndex.Do(func() {
		cache_Foreign_ErrorAtIndex = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 4093101208, UnsafePtr: unsafe.Pointer((&Constructor_Foreign_ErrorAtIndex{1, value0.IntVal, value1}))}
			})
		})
	})
	return cache_Foreign_ErrorAtIndex
}

var cache_Foreign_ErrorAtProperty gopurs_runtime.Value
var once_Foreign_ErrorAtProperty sync.Once

func Get_Foreign_ErrorAtProperty() gopurs_runtime.Value {
	once_Foreign_ErrorAtProperty.Do(func() {
		cache_Foreign_ErrorAtProperty = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 3832113409, UnsafePtr: unsafe.Pointer((&Constructor_Foreign_ErrorAtProperty{1, value0.StrVal(), value1}))}
			})
		})
	})
	return cache_Foreign_ErrorAtProperty
}

var cache_Foreign_unsafeToForeign gopurs_runtime.Value
var once_Foreign_unsafeToForeign sync.Once

func Get_Foreign_unsafeToForeign() gopurs_runtime.Value {
	once_Foreign_unsafeToForeign.Do(func() {
		cache_Foreign_unsafeToForeign = Get_Unsafe_Coerce_unsafeCoerce()
	})
	return cache_Foreign_unsafeToForeign
}

var cache_Foreign_unsafeToForeign__3031330712 gopurs_runtime.Value
var once_Foreign_unsafeToForeign__3031330712 sync.Once

func Get_Foreign_unsafeToForeign__3031330712() gopurs_runtime.Value {
	once_Foreign_unsafeToForeign__3031330712.Do(func() {
		cache_Foreign_unsafeToForeign__3031330712 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Foreign_unsafeToForeign__3031330712(__eta_norm_0_0_box.IntVal)
		})
	})
	return cache_Foreign_unsafeToForeign__3031330712
}

var cache_Foreign_unsafeToForeign__3847756350 gopurs_runtime.Value
var once_Foreign_unsafeToForeign__3847756350 sync.Once

func Get_Foreign_unsafeToForeign__3847756350() gopurs_runtime.Value {
	once_Foreign_unsafeToForeign__3847756350.Do(func() {
		cache_Foreign_unsafeToForeign__3847756350 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Foreign_unsafeToForeign__3847756350(__eta_norm_0_0_box.StrVal())
		})
	})
	return cache_Foreign_unsafeToForeign__3847756350
}

var cache_Foreign_unsafeFromForeign gopurs_runtime.Value
var once_Foreign_unsafeFromForeign sync.Once

func Get_Foreign_unsafeFromForeign() gopurs_runtime.Value {
	once_Foreign_unsafeFromForeign.Do(func() {
		cache_Foreign_unsafeFromForeign = Get_Unsafe_Coerce_unsafeCoerce()
	})
	return cache_Foreign_unsafeFromForeign
}

var cache_Foreign_showForeignError gopurs_runtime.Value
var once_Foreign_showForeignError sync.Once

func Get_Foreign_showForeignError() gopurs_runtime.Value {
	once_Foreign_showForeignError.Do(func() {
		cache_Foreign_showForeignError = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer((&Constructor_Data_Show_Show[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t0 gopurs_runtime.Value
			{
				if v_0.Type == 9 && v_0.IntVal == 2355653293 {
					__t0 = gopurs_runtime.Str(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString("(ForeignError ", Data_Show_ShowStringImpl((*Constructor_Foreign_ForeignError)(v_0.UnsafePtr).V0)), ")"))
					goto end_branch_0
				} else {

				}
			}
			{
				if v_0.Type == 9 && v_0.IntVal == 4093101208 {
					__t0 = gopurs_runtime.Str(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString("(ErrorAtIndex ", Data_Show_ShowIntImpl((*Constructor_Foreign_ErrorAtIndex)(v_0.UnsafePtr).V0)), " "), gopurs_runtime.Apply(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Foreign_showForeignError()).V0, (*Constructor_Foreign_ErrorAtIndex)(v_0.UnsafePtr).V1).StrVal()), ")"))
					goto end_branch_0
				} else {

				}
			}
			{
				if v_0.Type == 9 && v_0.IntVal == 3832113409 {
					__t0 = gopurs_runtime.Str(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString("(ErrorAtProperty ", Data_Show_ShowStringImpl((*Constructor_Foreign_ErrorAtProperty)(v_0.UnsafePtr).V0)), " "), gopurs_runtime.Apply(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Foreign_showForeignError()).V0, (*Constructor_Foreign_ErrorAtProperty)(v_0.UnsafePtr).V1).StrVal()), ")"))
					goto end_branch_0
				} else {

				}
			}
			{
				if v_0.Type == 9 && v_0.IntVal == 2536121783 {
					__t0 = gopurs_runtime.Str(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString("(TypeMismatch ", Data_Show_ShowStringImpl((*Constructor_Foreign_TypeMismatch)(v_0.UnsafePtr).V0)), " "), Data_Show_ShowStringImpl((*Constructor_Foreign_TypeMismatch)(v_0.UnsafePtr).V1)), ")"))
					goto end_branch_0
				} else {

				}
			}
			{
				__t0 = gopurs_runtime.Str(func() string { panic("Failed pattern match") }())
			}
		end_branch_0:
			return __t0
		})}))}
	})
	return cache_Foreign_showForeignError
}

var cache_Foreign_renderForeignError gopurs_runtime.Value
var once_Foreign_renderForeignError sync.Once

func Get_Foreign_renderForeignError() gopurs_runtime.Value {
	once_Foreign_renderForeignError.Do(func() {
		cache_Foreign_renderForeignError = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Foreign_renderForeignError(v_0_box))
		})
	})
	return cache_Foreign_renderForeignError
}

var cache_Foreign_readUndefined gopurs_runtime.Value
var once_Foreign_readUndefined sync.Once

func Get_Foreign_readUndefined() gopurs_runtime.Value {
	once_Foreign_readUndefined.Do(func() {
		cache_Foreign_readUndefined = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Foreign_readUndefined(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](dictMonad_0_box))
		})
	})
	return cache_Foreign_readUndefined
}

var cache_Foreign_readNullOrUndefined gopurs_runtime.Value
var once_Foreign_readNullOrUndefined sync.Once

func Get_Foreign_readNullOrUndefined() gopurs_runtime.Value {
	once_Foreign_readNullOrUndefined.Do(func() {
		cache_Foreign_readNullOrUndefined = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Foreign_readNullOrUndefined(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](dictMonad_0_box))
		})
	})
	return cache_Foreign_readNullOrUndefined
}

var cache_Foreign_readNull gopurs_runtime.Value
var once_Foreign_readNull sync.Once

func Get_Foreign_readNull() gopurs_runtime.Value {
	once_Foreign_readNull.Do(func() {
		cache_Foreign_readNull = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Foreign_readNull(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](dictMonad_0_box))
		})
	})
	return cache_Foreign_readNull
}

var cache_Foreign_fail gopurs_runtime.Value
var once_Foreign_fail sync.Once

func Get_Foreign_fail() gopurs_runtime.Value {
	once_Foreign_fail.Do(func() {
		cache_Foreign_fail = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Foreign_fail(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](dictMonad_0_box))
		})
	})
	return cache_Foreign_fail
}

var cache_Foreign_readArray gopurs_runtime.Value
var once_Foreign_readArray sync.Once

func Get_Foreign_readArray() gopurs_runtime.Value {
	once_Foreign_readArray.Do(func() {
		cache_Foreign_readArray = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Foreign_readArray(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](dictMonad_0_box))
		})
	})
	return cache_Foreign_readArray
}

var cache_Foreign_unsafeReadTagged gopurs_runtime.Value
var once_Foreign_unsafeReadTagged sync.Once

func Get_Foreign_unsafeReadTagged() gopurs_runtime.Value {
	once_Foreign_unsafeReadTagged.Do(func() {
		cache_Foreign_unsafeReadTagged = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Foreign_unsafeReadTagged(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](dictMonad_0_box))
		})
	})
	return cache_Foreign_unsafeReadTagged
}

var cache_Foreign_readBoolean gopurs_runtime.Value
var once_Foreign_readBoolean sync.Once

func Get_Foreign_readBoolean() gopurs_runtime.Value {
	once_Foreign_readBoolean.Do(func() {
		cache_Foreign_readBoolean = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Foreign_readBoolean(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](dictMonad_0_box))
		})
	})
	return cache_Foreign_readBoolean
}

var cache_Foreign_readNumber gopurs_runtime.Value
var once_Foreign_readNumber sync.Once

func Get_Foreign_readNumber() gopurs_runtime.Value {
	once_Foreign_readNumber.Do(func() {
		cache_Foreign_readNumber = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Foreign_readNumber(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](dictMonad_0_box))
		})
	})
	return cache_Foreign_readNumber
}

var cache_Foreign_readInt gopurs_runtime.Value
var once_Foreign_readInt sync.Once

func Get_Foreign_readInt() gopurs_runtime.Value {
	once_Foreign_readInt.Do(func() {
		cache_Foreign_readInt = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Foreign_readInt(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](dictMonad_0_box))
		})
	})
	return cache_Foreign_readInt
}

var cache_Foreign_readString gopurs_runtime.Value
var once_Foreign_readString sync.Once

func Get_Foreign_readString() gopurs_runtime.Value {
	once_Foreign_readString.Do(func() {
		cache_Foreign_readString = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Foreign_readString(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](dictMonad_0_box))
		})
	})
	return cache_Foreign_readString
}

var cache_Foreign_readChar gopurs_runtime.Value
var once_Foreign_readChar sync.Once

func Get_Foreign_readChar() gopurs_runtime.Value {
	once_Foreign_readChar.Do(func() {
		cache_Foreign_readChar = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Foreign_readChar(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](dictMonad_0_box))
		})
	})
	return cache_Foreign_readChar
}

var cache_Foreign_eqForeignError gopurs_runtime.Value
var once_Foreign_eqForeignError sync.Once

func Get_Foreign_eqForeignError() gopurs_runtime.Value {
	once_Foreign_eqForeignError.Do(func() {
		cache_Foreign_eqForeignError = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t0 bool
			{
				if x_0.Type == 9 && x_0.IntVal == 2355653293 {
					__t0 = (y_1.Type == 9 && y_1.IntVal == 2355653293) && (((*Constructor_Foreign_ForeignError)(x_0.UnsafePtr).V0) == ((*Constructor_Foreign_ForeignError)(y_1.UnsafePtr).V0))
					goto end_branch_0
				} else {

				}
			}
			{
				if x_0.Type == 9 && x_0.IntVal == 2536121783 {
					__t0 = (y_1.Type == 9 && y_1.IntVal == 2536121783) && ((((*Constructor_Foreign_TypeMismatch)(x_0.UnsafePtr).V0) == ((*Constructor_Foreign_TypeMismatch)(y_1.UnsafePtr).V0)) && (((*Constructor_Foreign_TypeMismatch)(x_0.UnsafePtr).V1) == ((*Constructor_Foreign_TypeMismatch)(y_1.UnsafePtr).V1)))
					goto end_branch_0
				} else {

				}
			}
			{
				if x_0.Type == 9 && x_0.IntVal == 4093101208 {
					__t0 = (y_1.Type == 9 && y_1.IntVal == 4093101208) && ((((*Constructor_Foreign_ErrorAtIndex)(x_0.UnsafePtr).V0) == ((*Constructor_Foreign_ErrorAtIndex)(y_1.UnsafePtr).V0)) && ((gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Foreign_eqForeignError()).V0, (*Constructor_Foreign_ErrorAtIndex)(x_0.UnsafePtr).V1, (*Constructor_Foreign_ErrorAtIndex)(y_1.UnsafePtr).V1).IntVal) != (0)))
					goto end_branch_0
				} else {

				}
			}
			{
				__t0 = (x_0.Type == 9 && x_0.IntVal == 3832113409) && ((y_1.Type == 9 && y_1.IntVal == 3832113409) && ((((*Constructor_Foreign_ErrorAtProperty)(x_0.UnsafePtr).V0) == ((*Constructor_Foreign_ErrorAtProperty)(y_1.UnsafePtr).V0)) && ((gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Foreign_eqForeignError()).V0, (*Constructor_Foreign_ErrorAtProperty)(x_0.UnsafePtr).V1, (*Constructor_Foreign_ErrorAtProperty)(y_1.UnsafePtr).V1).IntVal) != (0))))
			}
		end_branch_0:
			return gopurs_runtime.Bool(__t0)
		})}))}
	})
	return cache_Foreign_eqForeignError
}

var cache_Foreign_ordForeignError gopurs_runtime.Value
var once_Foreign_ordForeignError sync.Once

func Get_Foreign_ordForeignError() gopurs_runtime.Value {
	once_Foreign_ordForeignError.Do(func() {
		cache_Foreign_ordForeignError = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ord_Ord[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Foreign_eqForeignError()))}
		}), gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t9 gopurs_runtime.Value
			{
				if x_0.Type == 9 && x_0.IntVal == 2355653293 {
					var __t0 gopurs_runtime.Value
					{
						if y_1.Type == 9 && y_1.IntVal == 2355653293 {
							__t0 = gopurs_runtime.Apply2(Rebox_Foreign_4177771502_2406510097(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordString())).V1, gopurs_runtime.Str((*Constructor_Foreign_ForeignError)(x_0.UnsafePtr).V0), gopurs_runtime.Str((*Constructor_Foreign_ForeignError)(y_1.UnsafePtr).V0))
							goto end_branch_0
						} else {

						}
					}
					{
						__t0 = gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}
					}
				end_branch_0:
					__t9 = __t0
					goto end_branch_9
				} else {

				}
			}
			{
				if y_1.Type == 9 && y_1.IntVal == 2355653293 {
					__t9 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
					goto end_branch_9
				} else {

				}
			}
			{
				if x_0.Type == 9 && x_0.IntVal == 2536121783 {
					var __t3 gopurs_runtime.Value
					{
						if y_1.Type == 9 && y_1.IntVal == 2536121783 {
							// TAST (Let): v_2_1 shape=App(Other) bindingType=(ADT ["Data","Ordering","Ordering"] [])
							v_2_1 := uint32(gopurs_runtime.Apply2(Rebox_Foreign_4177771502_2406510097(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordString())).V1, gopurs_runtime.Str((*Constructor_Foreign_TypeMismatch)(x_0.UnsafePtr).V0), gopurs_runtime.Str((*Constructor_Foreign_TypeMismatch)(y_1.UnsafePtr).V0)).IntVal)
							_ = v_2_1
							var __t2 gopurs_runtime.Value
							{
								if v_2_1 == 1527465420 {
									__t2 = gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}
									goto end_branch_2
								} else {

								}
							}
							{
								if v_2_1 == 380165415 {
									__t2 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
									goto end_branch_2
								} else {

								}
							}
							{
								__t2 = gopurs_runtime.Apply2(Rebox_Foreign_4177771502_2406510097(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordString())).V1, gopurs_runtime.Str((*Constructor_Foreign_TypeMismatch)(x_0.UnsafePtr).V1), gopurs_runtime.Str((*Constructor_Foreign_TypeMismatch)(y_1.UnsafePtr).V1))
							}
						end_branch_2:
							__t3 = __t2
							goto end_branch_3
						} else {

						}
					}
					{
						__t3 = gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}
					}
				end_branch_3:
					__t9 = __t3
					goto end_branch_9
				} else {

				}
			}
			{
				if y_1.Type == 9 && y_1.IntVal == 2536121783 {
					__t9 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
					goto end_branch_9
				} else {

				}
			}
			{
				if x_0.Type == 9 && x_0.IntVal == 4093101208 {
					var __t6 gopurs_runtime.Value
					{
						if y_1.Type == 9 && y_1.IntVal == 4093101208 {
							// TAST (Let): v_2_4 shape=App(Other) bindingType=(ADT ["Data","Ordering","Ordering"] [])
							v_2_4 := uint32(gopurs_runtime.Apply2(Rebox_Foreign_4177771502_3308271157(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordInt())).V1, gopurs_runtime.Int((*Constructor_Foreign_ErrorAtIndex)(x_0.UnsafePtr).V0), gopurs_runtime.Int((*Constructor_Foreign_ErrorAtIndex)(y_1.UnsafePtr).V0)).IntVal)
							_ = v_2_4
							var __t5 gopurs_runtime.Value
							{
								if v_2_4 == 1527465420 {
									__t5 = gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}
									goto end_branch_5
								} else {

								}
							}
							{
								if v_2_4 == 380165415 {
									__t5 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
									goto end_branch_5
								} else {

								}
							}
							{
								__t5 = gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Foreign_ordForeignError()).V1, (*Constructor_Foreign_ErrorAtIndex)(x_0.UnsafePtr).V1, (*Constructor_Foreign_ErrorAtIndex)(y_1.UnsafePtr).V1)
							}
						end_branch_5:
							__t6 = __t5
							goto end_branch_6
						} else {

						}
					}
					{
						__t6 = gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}
					}
				end_branch_6:
					__t9 = __t6
					goto end_branch_9
				} else {

				}
			}
			{
				if y_1.Type == 9 && y_1.IntVal == 4093101208 {
					__t9 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
					goto end_branch_9
				} else {

				}
			}
			{
				if (x_0.Type == 9 && x_0.IntVal == 3832113409) && (y_1.Type == 9 && y_1.IntVal == 3832113409) {
					// TAST (Let): v_2_7 shape=App(Other) bindingType=(ADT ["Data","Ordering","Ordering"] [])
					v_2_7 := uint32(gopurs_runtime.Apply2(Rebox_Foreign_4177771502_2406510097(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordString())).V1, gopurs_runtime.Str((*Constructor_Foreign_ErrorAtProperty)(x_0.UnsafePtr).V0), gopurs_runtime.Str((*Constructor_Foreign_ErrorAtProperty)(y_1.UnsafePtr).V0)).IntVal)
					_ = v_2_7
					var __t8 gopurs_runtime.Value
					{
						if v_2_7 == 1527465420 {
							__t8 = gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}
							goto end_branch_8
						} else {

						}
					}
					{
						if v_2_7 == 380165415 {
							__t8 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
							goto end_branch_8
						} else {

						}
					}
					{
						__t8 = gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Foreign_ordForeignError()).V1, (*Constructor_Foreign_ErrorAtProperty)(x_0.UnsafePtr).V1, (*Constructor_Foreign_ErrorAtProperty)(y_1.UnsafePtr).V1)
					}
				end_branch_8:
					__t9 = __t8
					goto end_branch_9
				} else {

				}
			}
			{
				__t9 = gopurs_runtime.Value{Type: 9, IntVal: int64(func() uint32 { panic("Failed pattern match") }()), UnsafePtr: nil}
			}
		end_branch_9:
			return __t9
		})}))}
	})
	return cache_Foreign_ordForeignError
}

type Constructor_Foreign_ForeignError struct {
	Rc uint32
	V0 string
}

type Constructor_Foreign_TypeMismatch struct {
	Rc uint32
	V0 string
	V1 string
}

type Constructor_Foreign_ErrorAtIndex struct {
	Rc uint32
	V0 int64
	V1 gopurs_runtime.Value
}

type Constructor_Foreign_ErrorAtProperty struct {
	Rc uint32
	V0 string
	V1 gopurs_runtime.Value
}

func Call_Foreign_unsafeToForeign__3031330712(__eta_norm_0_0_loop int64) gopurs_runtime.Value {
unsafeToForeign__3031330712:
	for {
		if false {
			continue unsafeToForeign__3031330712
		}
		var __eta_norm_0_0 int64 = __eta_norm_0_0_loop
		_ = __eta_norm_0_0
		return gopurs_runtime.Int(__eta_norm_0_0)
	}
}

func Call_Foreign_unsafeToForeign__3847756350(__eta_norm_0_0_loop string) gopurs_runtime.Value {
unsafeToForeign__3847756350:
	for {
		if false {
			continue unsafeToForeign__3847756350
		}
		var __eta_norm_0_0 string = __eta_norm_0_0_loop
		_ = __eta_norm_0_0
		return gopurs_runtime.Str(__eta_norm_0_0)
	}
}

func Call_Foreign_renderForeignError(v_0_loop gopurs_runtime.Value) string {
renderForeignError:
	for {
		if false {
			continue renderForeignError
		}
		var v_0 gopurs_runtime.Value = v_0_loop
		_ = v_0
		var __t0 string
		{
			if v_0.Type == 9 && v_0.IntVal == 2355653293 {
				__t0 = (*Constructor_Foreign_ForeignError)(v_0.UnsafePtr).V0
				goto end_branch_0
			} else {

			}
		}
		{
			if v_0.Type == 9 && v_0.IntVal == 4093101208 {
				__t0 = gopurs_runtime.ConcatString(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString("Error at array index ", Data_Show_ShowIntImpl((*Constructor_Foreign_ErrorAtIndex)(v_0.UnsafePtr).V0)), ": "), Call_Foreign_renderForeignError((*Constructor_Foreign_ErrorAtIndex)(v_0.UnsafePtr).V1))
				goto end_branch_0
			} else {

			}
		}
		{
			if v_0.Type == 9 && v_0.IntVal == 3832113409 {
				__t0 = gopurs_runtime.ConcatString(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString("Error at property ", Data_Show_ShowStringImpl((*Constructor_Foreign_ErrorAtProperty)(v_0.UnsafePtr).V0)), ": "), Call_Foreign_renderForeignError((*Constructor_Foreign_ErrorAtProperty)(v_0.UnsafePtr).V1))
				goto end_branch_0
			} else {

			}
		}
		{
			if v_0.Type == 9 && v_0.IntVal == 2536121783 {
				__t0 = gopurs_runtime.ConcatString(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString("Type mismatch: expected ", (*Constructor_Foreign_TypeMismatch)(v_0.UnsafePtr).V0), ", found "), (*Constructor_Foreign_TypeMismatch)(v_0.UnsafePtr).V1)
				goto end_branch_0
			} else {

			}
		}
		{
			__t0 = func() string { panic("Failed pattern match") }()
		}
	end_branch_0:
		return __t0
	}
}

func Call_Foreign_readUndefined(dictMonad_0_loop *Constructor_Control_Monad_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictMonad_0 *Constructor_Control_Monad_Monad[gopurs_runtime.Value] = dictMonad_0_loop
	_ = dictMonad_0
	// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=Any
	__local_var_1_0 := Call_Control_Monad_Except_Trans_applicativeExceptT(gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(dictMonad_0)})
	_ = __local_var_1_0
	return gopurs_runtime.Func(func(value_2 gopurs_runtime.Value) gopurs_runtime.Value {
		var __t1 gopurs_runtime.Value
		{
			if (gopurs_runtime.Apply(Get_Foreign_isUndefined(), value_2).IntVal) != (0) {
				__t1 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))})
				goto end_branch_1
			} else {

			}
		}
		{
			__t1 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{value_2, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))})
		}
	end_branch_1:
		return __t1
	})
}

func Call_Foreign_readNullOrUndefined(dictMonad_0_loop *Constructor_Control_Monad_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictMonad_0 *Constructor_Control_Monad_Monad[gopurs_runtime.Value] = dictMonad_0_loop
	_ = dictMonad_0
	// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=Any
	__local_var_1_0 := Call_Control_Monad_Except_Trans_applicativeExceptT(gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(dictMonad_0)})
	_ = __local_var_1_0
	return gopurs_runtime.Func(func(value_2 gopurs_runtime.Value) gopurs_runtime.Value {
		var __t1 gopurs_runtime.Value
		{
			if ((gopurs_runtime.Apply(Get_Foreign_isNull(), value_2).IntVal) != (0)) || ((gopurs_runtime.Apply(Get_Foreign_isUndefined(), value_2).IntVal) != (0)) {
				__t1 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))})
				goto end_branch_1
			} else {

			}
		}
		{
			__t1 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{value_2, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))})
		}
	end_branch_1:
		return __t1
	})
}

func Call_Foreign_readNull(dictMonad_0_loop *Constructor_Control_Monad_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictMonad_0 *Constructor_Control_Monad_Monad[gopurs_runtime.Value] = dictMonad_0_loop
	_ = dictMonad_0
	// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=Any
	__local_var_1_0 := Call_Control_Monad_Except_Trans_applicativeExceptT(gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(dictMonad_0)})
	_ = __local_var_1_0
	return gopurs_runtime.Func(func(value_2 gopurs_runtime.Value) gopurs_runtime.Value {
		var __t1 gopurs_runtime.Value
		{
			if (gopurs_runtime.Apply(Get_Foreign_isNull(), value_2).IntVal) != (0) {
				__t1 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))})
				goto end_branch_1
			} else {

			}
		}
		{
			__t1 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{value_2, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))})
		}
	end_branch_1:
		return __t1
	})
}

func Call_Foreign_fail(dictMonad_0_loop *Constructor_Control_Monad_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictMonad_0 *Constructor_Control_Monad_Monad[gopurs_runtime.Value] = dictMonad_0_loop
	_ = dictMonad_0
	return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Call_Control_Monad_Error_Class_throwError(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Error_Class_MonadThrow[gopurs_runtime.Value, gopurs_runtime.Value]](Call_Control_Monad_Except_Trans_monadThrowExceptT(gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(dictMonad_0)}))), Get_Data_List_NonEmpty_singleton())
}

func Call_Foreign_readArray(dictMonad_0_loop *Constructor_Control_Monad_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictMonad_0 *Constructor_Control_Monad_Monad[gopurs_runtime.Value] = dictMonad_0_loop
	_ = dictMonad_0
	// TAST (Let): pure2_1_0 shape=App(Var) bindingType=(Func [(Array (ADT ["Foreign","Foreign"] []))] (TypeApp (TypeVar m$scope28) [(ADT ["Data","Either","Either"] [(ADT ["Data","NonEmpty","NonEmpty"] [(ADT ["Data","List","Types","List"] []), (ADT ["Foreign","ForeignError"] [])]), (Array (ADT ["Foreign","Foreign"] []))])]))
	pure2_1_0 := Call_Control_Applicative_pure(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Call_Control_Monad_Except_Trans_applicativeExceptT(gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(dictMonad_0)})))
	_ = pure2_1_0
	// TAST (Let): fail1_2_1 shape=App(Var) bindingType=(Func [(ADT ["Foreign","ForeignError"] [])] (TypeApp (TypeVar m$scope28) [(ADT ["Data","Either","Either"] [(ADT ["Data","NonEmpty","NonEmpty"] [(ADT ["Data","List","Types","List"] []), (ADT ["Foreign","ForeignError"] [])]), (Array (ADT ["Foreign","Foreign"] []))])]))
	fail1_2_1 := Call_Foreign_fail(dictMonad_0)
	_ = fail1_2_1
	return gopurs_runtime.Func(func(value_3 gopurs_runtime.Value) gopurs_runtime.Value {
		var __t2 gopurs_runtime.Value
		{
			if (gopurs_runtime.Apply(Get_Foreign_isArray(), value_3).IntVal) != (0) {
				__t2 = gopurs_runtime.Apply(pure2_1_0, gopurs_runtime.Array(*((*[]gopurs_runtime.Value)(gopurs_runtime.Apply(Get_Foreign_unsafeReadArray(), value_3).UnsafePtr))))
				goto end_branch_2
			} else {

			}
		}
		{
			__t2 = gopurs_runtime.Apply(fail1_2_1, gopurs_runtime.Value{Type: 9, IntVal: 2536121783, UnsafePtr: unsafe.Pointer((&Constructor_Foreign_TypeMismatch{1, "array", gopurs_runtime.Apply(Get_Foreign_tagOf(), value_3).StrVal()}))})
		}
	end_branch_2:
		return __t2
	})
}

func Call_Foreign_unsafeReadTagged(dictMonad_0_loop *Constructor_Control_Monad_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictMonad_0 *Constructor_Control_Monad_Monad[gopurs_runtime.Value] = dictMonad_0_loop
	_ = dictMonad_0
	// TAST (Let): fail1_1_0 shape=App(Var) bindingType=(Func [(ADT ["Foreign","ForeignError"] [])] (TypeApp (TypeVar m$scope31) [(ADT ["Data","Either","Either"] [(ADT ["Data","NonEmpty","NonEmpty"] [(ADT ["Data","List","Types","List"] []), (ADT ["Foreign","ForeignError"] [])]), (TypeVar a$scope32)])]))
	fail1_1_0 := Call_Foreign_fail(dictMonad_0)
	_ = fail1_1_0
	return gopurs_runtime.Func2(func(tag_2 gopurs_runtime.Value, value_3 gopurs_runtime.Value) gopurs_runtime.Value {
		var __t1 gopurs_runtime.Value
		{
			if (gopurs_runtime.Apply(Get_Foreign_tagOf(), value_3).StrVal()) == (tag_2.StrVal()) {
				__t1 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Control_Monad_Except_Trans_applicativeExceptT(gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(dictMonad_0)}), "pure"), value_3)
				goto end_branch_1
			} else {

			}
		}
		{
			__t1 = gopurs_runtime.Apply(fail1_1_0, gopurs_runtime.Value{Type: 9, IntVal: 2536121783, UnsafePtr: unsafe.Pointer((&Constructor_Foreign_TypeMismatch{1, tag_2.StrVal(), gopurs_runtime.Apply(Get_Foreign_tagOf(), value_3).StrVal()}))})
		}
	end_branch_1:
		return __t1
	})
}

func Call_Foreign_readBoolean(dictMonad_0_loop *Constructor_Control_Monad_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictMonad_0 *Constructor_Control_Monad_Monad[gopurs_runtime.Value] = dictMonad_0_loop
	_ = dictMonad_0
	// TAST (Let): unsafeReadTagged1__2457015837_1_0 shape=App(Var) bindingType=(Func [String, (ADT ["Foreign","Foreign"] [])] (TypeApp (TypeVar m$scope34) [(ADT ["Data","Either","Either"] [(ADT ["Data","NonEmpty","NonEmpty"] [(ADT ["Data","List","Types","List"] []), (ADT ["Foreign","ForeignError"] [])]), (ADT ["Foreign","Foreign"] [])])]))
	unsafeReadTagged1__2457015837_1_0 := Call_Foreign_unsafeReadTagged(dictMonad_0)
	_ = unsafeReadTagged1__2457015837_1_0
	return gopurs_runtime.Func(func(value_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Control_Monad_Except_Trans_functorExceptT(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictMonad_0.V1, gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})), "map"), Get_Foreign_unsafeReadBoolean(), gopurs_runtime.Apply2(unsafeReadTagged1__2457015837_1_0, gopurs_runtime.StrWithCodeUnits("Boolean", 7), value_2))
	})
}

func Call_Foreign_readNumber(dictMonad_0_loop *Constructor_Control_Monad_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictMonad_0 *Constructor_Control_Monad_Monad[gopurs_runtime.Value] = dictMonad_0_loop
	_ = dictMonad_0
	// TAST (Let): unsafeReadTagged1__2457015837_1_0 shape=App(Var) bindingType=(Func [String, (ADT ["Foreign","Foreign"] [])] (TypeApp (TypeVar m$scope36) [(ADT ["Data","Either","Either"] [(ADT ["Data","NonEmpty","NonEmpty"] [(ADT ["Data","List","Types","List"] []), (ADT ["Foreign","ForeignError"] [])]), (ADT ["Foreign","Foreign"] [])])]))
	unsafeReadTagged1__2457015837_1_0 := Call_Foreign_unsafeReadTagged(dictMonad_0)
	_ = unsafeReadTagged1__2457015837_1_0
	return gopurs_runtime.Func(func(value_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Control_Monad_Except_Trans_functorExceptT(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictMonad_0.V1, gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})), "map"), Get_Foreign_unsafeReadNumber(), gopurs_runtime.Apply2(unsafeReadTagged1__2457015837_1_0, gopurs_runtime.StrWithCodeUnits("Number", 6), value_2))
	})
}

func Call_Foreign_readInt(dictMonad_0_loop *Constructor_Control_Monad_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictMonad_0 *Constructor_Control_Monad_Monad[gopurs_runtime.Value] = dictMonad_0_loop
	_ = dictMonad_0
	// TAST (Let): go__map_1_0 shape=App(Var) bindingType=(Func [(Func [(ADT ["Data","Either","Either"] [(ADT ["Data","NonEmpty","NonEmpty"] [(ADT ["Data","List","Types","List"] []), (ADT ["Foreign","ForeignError"] [])]), Number])] (ADT ["Data","Either","Either"] [(ADT ["Data","NonEmpty","NonEmpty"] [(ADT ["Data","List","Types","List"] []), (ADT ["Foreign","ForeignError"] [])]), Int])), (TypeApp (TypeVar m$scope38) [(ADT ["Data","Either","Either"] [(ADT ["Data","NonEmpty","NonEmpty"] [(ADT ["Data","List","Types","List"] []), (ADT ["Foreign","ForeignError"] [])]), Number])])] (TypeApp (TypeVar m$scope38) [(ADT ["Data","Either","Either"] [(ADT ["Data","NonEmpty","NonEmpty"] [(ADT ["Data","List","Types","List"] []), (ADT ["Foreign","ForeignError"] [])]), Int])]))
	go__map_1_0 := Call_Data_Functor_go__map(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictMonad_0.V1, gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})))
	_ = go__map_1_0
	// TAST (Let): readNumber1__3116725317_2_1 shape=App(Var) bindingType=(Func [(ADT ["Foreign","Foreign"] [])] (TypeApp (TypeVar m$scope38) [(ADT ["Data","Either","Either"] [(ADT ["Data","NonEmpty","NonEmpty"] [(ADT ["Data","List","Types","List"] []), (ADT ["Foreign","ForeignError"] [])]), Number])]))
	readNumber1__3116725317_2_1 := Call_Foreign_readNumber(dictMonad_0)
	_ = readNumber1__3116725317_2_1
	return gopurs_runtime.Func(func(value_3 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_4_2 shape=App(Var) bindingType=Any
		__local_var_4_2 := gopurs_runtime.Apply3(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_List_Types_NonEmptyList(), Call_Data_NonEmpty_singleton(Rebox_Foreign_2173899317_3706288089(Rebox_Foreign_3706288089_2173899317(gopurs_runtime.CoerceToStruct[Constructor_Control_Plus_Plus[gopurs_runtime.Value]](Get_Data_List_Types_plusList())))), gopurs_runtime.Value{Type: 9, IntVal: 2536121783, UnsafePtr: unsafe.Pointer((&Constructor_Foreign_TypeMismatch{1, "Int", gopurs_runtime.Apply(Get_Foreign_tagOf(), value_3).StrVal()}))})
		_ = __local_var_4_2
		// TAST (Let): __local_var_5_3 shape=App(Var) bindingType=(Func [Number] (ADT ["Data","Either","Either"] [(ADT ["Data","NonEmpty","NonEmpty"] [(ADT ["Data","List","Types","List"] []), (ADT ["Foreign","ForeignError"] [])]), Int]))
		__local_var_5_3 := gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply2(Get_Data_Maybe_maybe__565733802(), func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
				V2 bool
			}{__local_var_4_2, gopurs_runtime.Value{}, false}
			if _v.V2 {
				return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
		}(), Get_Data_Either_Right()), Get_Data_Int_fromNumber())
		_ = __local_var_5_3
		return gopurs_runtime.Apply2(go__map_1_0, gopurs_runtime.Func(func(v2_6 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t4 gopurs_runtime.Value
			{
				if v2_6.Type == 9 && v2_6.IntVal == 3711209382 {
					__t4 = func() gopurs_runtime.Value {
						_v := struct {
							V0 gopurs_runtime.Value
							V1 gopurs_runtime.Value
							V2 bool
						}{__local_var_4_2, gopurs_runtime.Value{}, false}
						if _v.V2 {
							return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
						}
						return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
					}()
					goto end_branch_4
				} else {

				}
			}
			{
				if v2_6.Type == 9 && v2_6.IntVal == 2465973597 {
					__t4 = gopurs_runtime.Apply(__local_var_5_3, gopurs_runtime.Float((*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_6.UnsafePtr).V0.FloatVal()))
					goto end_branch_4
				} else {

				}
			}
			{
				__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_4:
			return __t4
		}), gopurs_runtime.Apply(readNumber1__3116725317_2_1, value_3))
	})
}

func Call_Foreign_readString(dictMonad_0_loop *Constructor_Control_Monad_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictMonad_0 *Constructor_Control_Monad_Monad[gopurs_runtime.Value] = dictMonad_0_loop
	_ = dictMonad_0
	// TAST (Let): unsafeReadTagged1__2457015837_1_0 shape=App(Var) bindingType=(Func [String, (ADT ["Foreign","Foreign"] [])] (TypeApp (TypeVar m$scope43) [(ADT ["Data","Either","Either"] [(ADT ["Data","NonEmpty","NonEmpty"] [(ADT ["Data","List","Types","List"] []), (ADT ["Foreign","ForeignError"] [])]), (ADT ["Foreign","Foreign"] [])])]))
	unsafeReadTagged1__2457015837_1_0 := Call_Foreign_unsafeReadTagged(dictMonad_0)
	_ = unsafeReadTagged1__2457015837_1_0
	return gopurs_runtime.Func(func(value_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Control_Monad_Except_Trans_functorExceptT(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictMonad_0.V1, gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})), "map"), Get_Foreign_unsafeReadString(), gopurs_runtime.Apply2(unsafeReadTagged1__2457015837_1_0, gopurs_runtime.StrWithCodeUnits("String", 6), value_2))
	})
}

func Call_Foreign_readChar(dictMonad_0_loop *Constructor_Control_Monad_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictMonad_0 *Constructor_Control_Monad_Monad[gopurs_runtime.Value] = dictMonad_0_loop
	_ = dictMonad_0
	// TAST (Let): go__map_1_0 shape=App(Var) bindingType=(Func [(Func [(ADT ["Data","Either","Either"] [(ADT ["Data","NonEmpty","NonEmpty"] [(ADT ["Data","List","Types","List"] []), (ADT ["Foreign","ForeignError"] [])]), String])] (ADT ["Data","Either","Either"] [(ADT ["Data","NonEmpty","NonEmpty"] [(ADT ["Data","List","Types","List"] []), (ADT ["Foreign","ForeignError"] [])]), Char])), (TypeApp (TypeVar m$scope45) [(ADT ["Data","Either","Either"] [(ADT ["Data","NonEmpty","NonEmpty"] [(ADT ["Data","List","Types","List"] []), (ADT ["Foreign","ForeignError"] [])]), String])])] (TypeApp (TypeVar m$scope45) [(ADT ["Data","Either","Either"] [(ADT ["Data","NonEmpty","NonEmpty"] [(ADT ["Data","List","Types","List"] []), (ADT ["Foreign","ForeignError"] [])]), Char])]))
	go__map_1_0 := Call_Data_Functor_go__map(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictMonad_0.V1, gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})))
	_ = go__map_1_0
	// TAST (Let): readString1__3695720883_2_1 shape=App(Var) bindingType=(Func [(ADT ["Foreign","Foreign"] [])] (TypeApp (TypeVar m$scope45) [(ADT ["Data","Either","Either"] [(ADT ["Data","NonEmpty","NonEmpty"] [(ADT ["Data","List","Types","List"] []), (ADT ["Foreign","ForeignError"] [])]), String])]))
	readString1__3695720883_2_1 := Call_Foreign_readString(dictMonad_0)
	_ = readString1__3695720883_2_1
	return gopurs_runtime.Func(func(value_3 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_4_2 shape=App(Var) bindingType=Any
		__local_var_4_2 := gopurs_runtime.Apply3(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_List_Types_NonEmptyList(), Call_Data_NonEmpty_singleton(Rebox_Foreign_2173899317_3706288089(Rebox_Foreign_3706288089_2173899317(gopurs_runtime.CoerceToStruct[Constructor_Control_Plus_Plus[gopurs_runtime.Value]](Get_Data_List_Types_plusList())))), gopurs_runtime.Value{Type: 9, IntVal: 2536121783, UnsafePtr: unsafe.Pointer((&Constructor_Foreign_TypeMismatch{1, "Char", gopurs_runtime.Apply(Get_Foreign_tagOf(), value_3).StrVal()}))})
		_ = __local_var_4_2
		// TAST (Let): __local_var_5_3 shape=App(Var) bindingType=(Func [String] (ADT ["Data","Either","Either"] [(ADT ["Data","NonEmpty","NonEmpty"] [(ADT ["Data","List","Types","List"] []), (ADT ["Foreign","ForeignError"] [])]), Char]))
		__local_var_5_3 := gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply2(Get_Data_Maybe_maybe__1119287216(), func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
				V2 bool
			}{__local_var_4_2, gopurs_runtime.Value{}, false}
			if _v.V2 {
				return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
		}(), Get_Data_Either_Right()), Get_Data_String_CodeUnits_toChar())
		_ = __local_var_5_3
		return gopurs_runtime.Apply2(go__map_1_0, gopurs_runtime.Func(func(v2_6 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t4 gopurs_runtime.Value
			{
				if v2_6.Type == 9 && v2_6.IntVal == 3711209382 {
					__t4 = func() gopurs_runtime.Value {
						_v := struct {
							V0 gopurs_runtime.Value
							V1 gopurs_runtime.Value
							V2 bool
						}{__local_var_4_2, gopurs_runtime.Value{}, false}
						if _v.V2 {
							return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
						}
						return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
					}()
					goto end_branch_4
				} else {

				}
			}
			{
				if v2_6.Type == 9 && v2_6.IntVal == 2465973597 {
					__t4 = gopurs_runtime.Apply(__local_var_5_3, gopurs_runtime.NormalizeStringValue((*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_6.UnsafePtr).V0))
					goto end_branch_4
				} else {

				}
			}
			{
				__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_4:
			return __t4
		}), gopurs_runtime.Apply(readString1__3695720883_2_1, value_3))
	})
}

func Rebox_Foreign_2173899317_3706288089(in *Constructor_Control_Plus_Plus[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]) *Constructor_Control_Plus_Plus[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Control_Plus_Plus[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Foreign_3706288089_2173899317(in *Constructor_Control_Plus_Plus[gopurs_runtime.Value]) *Constructor_Control_Plus_Plus[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Control_Plus_Plus[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Foreign_4177771502_2406510097(in *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) *Constructor_Data_Ord_Ord[string] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ord_Ord[string])(unsafe.Pointer(in))
}

func Rebox_Foreign_4177771502_3308271157(in *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) *Constructor_Data_Ord_Ord[int64] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ord_Ord[int64])(unsafe.Pointer(in))
}

func Get_Foreign_isArray() gopurs_runtime.Value {
	return _Gopurs_Foreign_IsArray
}

func Get_Foreign_isNull() gopurs_runtime.Value {
	return _Gopurs_Foreign_IsNull
}

func Get_Foreign_isUndefined() gopurs_runtime.Value {
	return _Gopurs_Foreign_IsUndefined
}

func Get_Foreign_tagOf() gopurs_runtime.Value {
	return _Gopurs_Foreign_TagOf
}

func Get_Foreign_typeOf() gopurs_runtime.Value {
	return _Gopurs_Foreign_TypeOf
}

func Get_Foreign_unboxForJSON() gopurs_runtime.Value {
	return _Gopurs_Foreign_UnboxForJSON
}

func Get_Foreign_unsafeReadArray() gopurs_runtime.Value {
	return _Gopurs_Foreign_UnsafeReadArray
}

func Get_Foreign_unsafeReadBoolean() gopurs_runtime.Value {
	return _Gopurs_Foreign_UnsafeReadBoolean
}

func Get_Foreign_unsafeReadNumber() gopurs_runtime.Value {
	return _Gopurs_Foreign_UnsafeReadNumber
}

func Get_Foreign_unsafeReadString() gopurs_runtime.Value {
	return _Gopurs_Foreign_UnsafeReadString
}
