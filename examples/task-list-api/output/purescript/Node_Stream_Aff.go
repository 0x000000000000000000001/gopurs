package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Node_Stream_Aff_compose gopurs_runtime.Value
var once_Node_Stream_Aff_compose sync.Once

func Get_Node_Stream_Aff_compose() gopurs_runtime.Value {
	once_Node_Stream_Aff_compose.Do(func() {
		cache_Node_Stream_Aff_compose = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Node_Stream_Aff_compose
}

var cache_Node_Stream_Aff_bind gopurs_runtime.Value
var once_Node_Stream_Aff_bind sync.Once

func Get_Node_Stream_Aff_bind() gopurs_runtime.Value {
	once_Node_Stream_Aff_bind.Do(func() {
		cache_Node_Stream_Aff_bind = Get_Control_Bind_bind__1941644084()
	})
	return cache_Node_Stream_Aff_bind
}

var cache_Node_Stream_Aff_pure gopurs_runtime.Value
var once_Node_Stream_Aff_pure sync.Once

func Get_Node_Stream_Aff_pure() gopurs_runtime.Value {
	once_Node_Stream_Aff_pure.Do(func() {
		cache_Node_Stream_Aff_pure = Get_Control_Applicative_pure__2088919292()
	})
	return cache_Node_Stream_Aff_pure
}

var cache_Node_Stream_Aff_tailRecM gopurs_runtime.Value
var once_Node_Stream_Aff_tailRecM sync.Once

func Get_Node_Stream_Aff_tailRecM() gopurs_runtime.Value {
	once_Node_Stream_Aff_tailRecM.Do(func() {
		cache_Node_Stream_Aff_tailRecM = gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]](Get_Control_Monad_Rec_Class_monadRecEffect()).V1
	})
	return cache_Node_Stream_Aff_tailRecM
}

var cache_Node_Stream_Aff_discard gopurs_runtime.Value
var once_Node_Stream_Aff_discard sync.Once

func Get_Node_Stream_Aff_discard() gopurs_runtime.Value {
	once_Node_Stream_Aff_discard.Do(func() {
		cache_Node_Stream_Aff_discard = Get_Control_Bind_discard__932405161()
	})
	return cache_Node_Stream_Aff_discard
}

var cache_Node_Stream_Aff_pure1 gopurs_runtime.Value
var once_Node_Stream_Aff_pure1 sync.Once

func Get_Node_Stream_Aff_pure1() gopurs_runtime.Value {
	once_Node_Stream_Aff_pure1.Do(func() {
		cache_Node_Stream_Aff_pure1 = Get_Control_Applicative_pure__738246428()
	})
	return cache_Node_Stream_Aff_pure1
}

var cache_Node_Stream_Aff_bind1 gopurs_runtime.Value
var once_Node_Stream_Aff_bind1 sync.Once

func Get_Node_Stream_Aff_bind1() gopurs_runtime.Value {
	once_Node_Stream_Aff_bind1.Do(func() {
		cache_Node_Stream_Aff_bind1 = Get_Control_Bind_bind__4216591092()
	})
	return cache_Node_Stream_Aff_bind1
}

var cache_Node_Stream_Aff_add gopurs_runtime.Value
var once_Node_Stream_Aff_add sync.Once

func Get_Node_Stream_Aff_add() gopurs_runtime.Value {
	once_Node_Stream_Aff_add.Do(func() {
		cache_Node_Stream_Aff_add = Get_Data_Semiring_add__4086903470()
	})
	return cache_Node_Stream_Aff_add
}

var cache_Node_Stream_Aff_bind2 gopurs_runtime.Value
var once_Node_Stream_Aff_bind2 sync.Once

func Get_Node_Stream_Aff_bind2() gopurs_runtime.Value {
	once_Node_Stream_Aff_bind2.Do(func() {
		cache_Node_Stream_Aff_bind2 = Get_Control_Bind_bind__4149781044()
	})
	return cache_Node_Stream_Aff_bind2
}

var cache_Node_Stream_Aff_discard1 gopurs_runtime.Value
var once_Node_Stream_Aff_discard1 sync.Once

func Get_Node_Stream_Aff_discard1() gopurs_runtime.Value {
	once_Node_Stream_Aff_discard1.Do(func() {
		cache_Node_Stream_Aff_discard1 = Get_Control_Bind_discard__1192935977()
	})
	return cache_Node_Stream_Aff_discard1
}

var cache_Node_Stream_Aff_pure2 gopurs_runtime.Value
var once_Node_Stream_Aff_pure2 sync.Once

func Get_Node_Stream_Aff_pure2() gopurs_runtime.Value {
	once_Node_Stream_Aff_pure2.Do(func() {
		cache_Node_Stream_Aff_pure2 = Get_Effect_pureE()
	})
	return cache_Node_Stream_Aff_pure2
}

var cache_Node_Stream_Aff_join gopurs_runtime.Value
var once_Node_Stream_Aff_join sync.Once

func Get_Node_Stream_Aff_join() gopurs_runtime.Value {
	once_Node_Stream_Aff_join.Do(func() {
		cache_Node_Stream_Aff_join = gopurs_runtime.Apply(Get_Control_Bind_join(), gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect()))})
	})
	return cache_Node_Stream_Aff_join
}

var cache_Node_Stream_Aff_bindFlipped gopurs_runtime.Value
var once_Node_Stream_Aff_bindFlipped sync.Once

func Get_Node_Stream_Aff_bindFlipped() gopurs_runtime.Value {
	once_Node_Stream_Aff_bindFlipped.Do(func() {
		cache_Node_Stream_Aff_bindFlipped = Get_Control_Bind_bindFlipped__857303188()
	})
	return cache_Node_Stream_Aff_bindFlipped
}

var cache_Node_Stream_Aff_bind3 gopurs_runtime.Value
var once_Node_Stream_Aff_bind3 sync.Once

func Get_Node_Stream_Aff_bind3() gopurs_runtime.Value {
	once_Node_Stream_Aff_bind3.Do(func() {
		cache_Node_Stream_Aff_bind3 = Get_Effect_bindE()
	})
	return cache_Node_Stream_Aff_bind3
}

var cache_Node_Stream_Aff_liftST gopurs_runtime.Value
var once_Node_Stream_Aff_liftST sync.Once

func Get_Node_Stream_Aff_liftST() gopurs_runtime.Value {
	once_Node_Stream_Aff_liftST.Do(func() {
		cache_Node_Stream_Aff_liftST = Get_Control_Monad_ST_Global_toEffect()
	})
	return cache_Node_Stream_Aff_liftST
}

var cache_Node_Stream_Aff_mempty gopurs_runtime.Value
var once_Node_Stream_Aff_mempty sync.Once

func Get_Node_Stream_Aff_mempty() gopurs_runtime.Value {
	once_Node_Stream_Aff_mempty.Do(func() {
		cache_Node_Stream_Aff_mempty = Call_Data_Monoid_mempty(Call_Effect_monoidEffect(gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](Get_Data_Monoid_monoidUnit()))}))
	})
	return cache_Node_Stream_Aff_mempty
}

var cache_Node_Stream_Aff_bind4 gopurs_runtime.Value
var once_Node_Stream_Aff_bind4 sync.Once

func Get_Node_Stream_Aff_bind4() gopurs_runtime.Value {
	once_Node_Stream_Aff_bind4.Do(func() {
		cache_Node_Stream_Aff_bind4 = Get_Control_Bind_bind__1370440500()
	})
	return cache_Node_Stream_Aff_bind4
}

var cache_Node_Stream_Aff_discard2 gopurs_runtime.Value
var once_Node_Stream_Aff_discard2 sync.Once

func Get_Node_Stream_Aff_discard2() gopurs_runtime.Value {
	once_Node_Stream_Aff_discard2.Do(func() {
		cache_Node_Stream_Aff_discard2 = Get_Control_Bind_discard__143037513()
	})
	return cache_Node_Stream_Aff_discard2
}

var cache_Node_Stream_Aff_bind5 gopurs_runtime.Value
var once_Node_Stream_Aff_bind5 sync.Once

func Get_Node_Stream_Aff_bind5() gopurs_runtime.Value {
	once_Node_Stream_Aff_bind5.Do(func() {
		cache_Node_Stream_Aff_bind5 = Get_Control_Bind_bind__931913300()
	})
	return cache_Node_Stream_Aff_bind5
}

var cache_Node_Stream_Aff_liftST1 gopurs_runtime.Value
var once_Node_Stream_Aff_liftST1 sync.Once

func Get_Node_Stream_Aff_liftST1() gopurs_runtime.Value {
	once_Node_Stream_Aff_liftST1.Do(func() {
		cache_Node_Stream_Aff_liftST1 = Get_Control_Monad_ST_Global_toEffect()
	})
	return cache_Node_Stream_Aff_liftST1
}

var cache_Node_Stream_Aff_void gopurs_runtime.Value
var once_Node_Stream_Aff_void sync.Once

func Get_Node_Stream_Aff_void() gopurs_runtime.Value {
	once_Node_Stream_Aff_void.Do(func() {
		cache_Node_Stream_Aff_void = Call_Data_Functor_void(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()))
	})
	return cache_Node_Stream_Aff_void
}

var cache_Node_Stream_Aff_liftST2 gopurs_runtime.Value
var once_Node_Stream_Aff_liftST2 sync.Once

func Get_Node_Stream_Aff_liftST2() gopurs_runtime.Value {
	once_Node_Stream_Aff_liftST2.Do(func() {
		cache_Node_Stream_Aff_liftST2 = Get_Control_Monad_ST_Global_toEffect()
	})
	return cache_Node_Stream_Aff_liftST2
}

var cache_Node_Stream_Aff_compose1 gopurs_runtime.Value
var once_Node_Stream_Aff_compose1 sync.Once

func Get_Node_Stream_Aff_compose1() gopurs_runtime.Value {
	once_Node_Stream_Aff_compose1.Do(func() {
		cache_Node_Stream_Aff_compose1 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Node_Stream_Aff_compose1
}

var cache_Node_Stream_Aff_bind6 gopurs_runtime.Value
var once_Node_Stream_Aff_bind6 sync.Once

func Get_Node_Stream_Aff_bind6() gopurs_runtime.Value {
	once_Node_Stream_Aff_bind6.Do(func() {
		cache_Node_Stream_Aff_bind6 = Get_Control_Bind_bind__3398628340()
	})
	return cache_Node_Stream_Aff_bind6
}

var cache_Node_Stream_Aff_pure3 gopurs_runtime.Value
var once_Node_Stream_Aff_pure3 sync.Once

func Get_Node_Stream_Aff_pure3() gopurs_runtime.Value {
	once_Node_Stream_Aff_pure3.Do(func() {
		cache_Node_Stream_Aff_pure3 = Get_Control_Applicative_pure__3156000668()
	})
	return cache_Node_Stream_Aff_pure3
}

var cache_Node_Stream_Aff_bind7 gopurs_runtime.Value
var once_Node_Stream_Aff_bind7 sync.Once

func Get_Node_Stream_Aff_bind7() gopurs_runtime.Value {
	once_Node_Stream_Aff_bind7.Do(func() {
		cache_Node_Stream_Aff_bind7 = Get_Control_Bind_bind__2676662740()
	})
	return cache_Node_Stream_Aff_bind7
}

var cache_Node_Stream_Aff_pure4 gopurs_runtime.Value
var once_Node_Stream_Aff_pure4 sync.Once

func Get_Node_Stream_Aff_pure4() gopurs_runtime.Value {
	once_Node_Stream_Aff_pure4.Do(func() {
		cache_Node_Stream_Aff_pure4 = Get_Control_Applicative_pure__1379658716()
	})
	return cache_Node_Stream_Aff_pure4
}

var cache_Node_Stream_Aff_discard3 gopurs_runtime.Value
var once_Node_Stream_Aff_discard3 sync.Once

func Get_Node_Stream_Aff_discard3() gopurs_runtime.Value {
	once_Node_Stream_Aff_discard3.Do(func() {
		cache_Node_Stream_Aff_discard3 = Get_Control_Bind_discard__2928924777()
	})
	return cache_Node_Stream_Aff_discard3
}

var cache_Node_Stream_Aff_bind8 gopurs_runtime.Value
var once_Node_Stream_Aff_bind8 sync.Once

func Get_Node_Stream_Aff_bind8() gopurs_runtime.Value {
	once_Node_Stream_Aff_bind8.Do(func() {
		cache_Node_Stream_Aff_bind8 = Get_Control_Bind_bind__2639965556()
	})
	return cache_Node_Stream_Aff_bind8
}

var cache_Node_Stream_Aff_conj gopurs_runtime.Value
var once_Node_Stream_Aff_conj sync.Once

func Get_Node_Stream_Aff_conj() gopurs_runtime.Value {
	once_Node_Stream_Aff_conj.Do(func() {
		cache_Node_Stream_Aff_conj = Get_Data_HeytingAlgebra_conj__2421783737()
	})
	return cache_Node_Stream_Aff_conj
}

var cache_Node_Stream_Aff_eq gopurs_runtime.Value
var once_Node_Stream_Aff_eq sync.Once

func Get_Node_Stream_Aff_eq() gopurs_runtime.Value {
	once_Node_Stream_Aff_eq.Do(func() {
		cache_Node_Stream_Aff_eq = Get_Data_Eq_eq__3020190905()
	})
	return cache_Node_Stream_Aff_eq
}

var cache_Node_Stream_Aff_bind9 gopurs_runtime.Value
var once_Node_Stream_Aff_bind9 sync.Once

func Get_Node_Stream_Aff_bind9() gopurs_runtime.Value {
	once_Node_Stream_Aff_bind9.Do(func() {
		cache_Node_Stream_Aff_bind9 = Get_Control_Bind_bind__4020646868()
	})
	return cache_Node_Stream_Aff_bind9
}

var cache_Node_Stream_Aff_compose2 gopurs_runtime.Value
var once_Node_Stream_Aff_compose2 sync.Once

func Get_Node_Stream_Aff_compose2() gopurs_runtime.Value {
	once_Node_Stream_Aff_compose2.Do(func() {
		cache_Node_Stream_Aff_compose2 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Node_Stream_Aff_compose2
}

var cache_Node_Stream_Aff_lessThan gopurs_runtime.Value
var once_Node_Stream_Aff_lessThan sync.Once

func Get_Node_Stream_Aff_lessThan() gopurs_runtime.Value {
	once_Node_Stream_Aff_lessThan.Do(func() {
		cache_Node_Stream_Aff_lessThan = Get_Data_Ord_lessThan__3348776441()
	})
	return cache_Node_Stream_Aff_lessThan
}

var cache_Node_Stream_Aff_applySecond gopurs_runtime.Value
var once_Node_Stream_Aff_applySecond sync.Once

func Get_Node_Stream_Aff_applySecond() gopurs_runtime.Value {
	once_Node_Stream_Aff_applySecond.Do(func() {
		cache_Node_Stream_Aff_applySecond = Get_Control_Apply_applySecond__3440890937()
	})
	return cache_Node_Stream_Aff_applySecond
}

var cache_Node_Stream_Aff_bind10 gopurs_runtime.Value
var once_Node_Stream_Aff_bind10 sync.Once

func Get_Node_Stream_Aff_bind10() gopurs_runtime.Value {
	once_Node_Stream_Aff_bind10.Do(func() {
		cache_Node_Stream_Aff_bind10 = Get_Control_Bind_bind__1903947764()
	})
	return cache_Node_Stream_Aff_bind10
}

var cache_Node_Stream_Aff_bind11 gopurs_runtime.Value
var once_Node_Stream_Aff_bind11 sync.Once

func Get_Node_Stream_Aff_bind11() gopurs_runtime.Value {
	once_Node_Stream_Aff_bind11.Do(func() {
		cache_Node_Stream_Aff_bind11 = Get_Control_Bind_bind__967963220()
	})
	return cache_Node_Stream_Aff_bind11
}

var cache_Node_Stream_Aff_sub gopurs_runtime.Value
var once_Node_Stream_Aff_sub sync.Once

func Get_Node_Stream_Aff_sub() gopurs_runtime.Value {
	once_Node_Stream_Aff_sub.Do(func() {
		cache_Node_Stream_Aff_sub = Get_Data_Ring_sub__4157804846()
	})
	return cache_Node_Stream_Aff_sub
}

var cache_Node_Stream_Aff_greaterThanOrEq gopurs_runtime.Value
var once_Node_Stream_Aff_greaterThanOrEq sync.Once

func Get_Node_Stream_Aff_greaterThanOrEq() gopurs_runtime.Value {
	once_Node_Stream_Aff_greaterThanOrEq.Do(func() {
		cache_Node_Stream_Aff_greaterThanOrEq = Get_Data_Ord_greaterThanOrEq__3348776441()
	})
	return cache_Node_Stream_Aff_greaterThanOrEq
}

var cache_Node_Stream_Aff_bind12 gopurs_runtime.Value
var once_Node_Stream_Aff_bind12 sync.Once

func Get_Node_Stream_Aff_bind12() gopurs_runtime.Value {
	once_Node_Stream_Aff_bind12.Do(func() {
		cache_Node_Stream_Aff_bind12 = Get_Control_Bind_bind__2595690612()
	})
	return cache_Node_Stream_Aff_bind12
}

var cache_Node_Stream_Aff_bind13 gopurs_runtime.Value
var once_Node_Stream_Aff_bind13 sync.Once

func Get_Node_Stream_Aff_bind13() gopurs_runtime.Value {
	once_Node_Stream_Aff_bind13.Do(func() {
		cache_Node_Stream_Aff_bind13 = Get_Control_Bind_bind__1455934228()
	})
	return cache_Node_Stream_Aff_bind13
}

var cache_Node_Stream_Aff_compose3 gopurs_runtime.Value
var once_Node_Stream_Aff_compose3 sync.Once

func Get_Node_Stream_Aff_compose3() gopurs_runtime.Value {
	once_Node_Stream_Aff_compose3.Do(func() {
		cache_Node_Stream_Aff_compose3 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Node_Stream_Aff_compose3
}

var cache_Node_Stream_Aff_go__map gopurs_runtime.Value
var once_Node_Stream_Aff_go__map sync.Once

func Get_Node_Stream_Aff_go__map() gopurs_runtime.Value {
	once_Node_Stream_Aff_go__map.Do(func() {
		cache_Node_Stream_Aff_go__map = Get_Data_Functor_map__2482553499()
	})
	return cache_Node_Stream_Aff_go__map
}

var cache_Node_Stream_Aff_pure5 gopurs_runtime.Value
var once_Node_Stream_Aff_pure5 sync.Once

func Get_Node_Stream_Aff_pure5() gopurs_runtime.Value {
	once_Node_Stream_Aff_pure5.Do(func() {
		cache_Node_Stream_Aff_pure5 = gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Control_Applicative_applicativeArray()).V1
	})
	return cache_Node_Stream_Aff_pure5
}

var cache_Node_Stream_Aff_compose4 gopurs_runtime.Value
var once_Node_Stream_Aff_compose4 sync.Once

func Get_Node_Stream_Aff_compose4() gopurs_runtime.Value {
	once_Node_Stream_Aff_compose4.Do(func() {
		cache_Node_Stream_Aff_compose4 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Node_Stream_Aff_compose4
}

var cache_Node_Stream_Aff_write gopurs_runtime.Value
var once_Node_Stream_Aff_write sync.Once

func Get_Node_Stream_Aff_write() gopurs_runtime.Value {
	once_Node_Stream_Aff_write.Do(func() {
		cache_Node_Stream_Aff_write = gopurs_runtime.Func(func(dictMonadAff_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Stream_Aff_write(gopurs_runtime.CoerceToStruct[Constructor_Effect_Aff_Class_MonadAff[gopurs_runtime.Value]](dictMonadAff_0_box))
		})
	})
	return cache_Node_Stream_Aff_write
}

var cache_Node_Stream_Aff_toStringUTF8 gopurs_runtime.Value
var once_Node_Stream_Aff_toStringUTF8 sync.Once

func Get_Node_Stream_Aff_toStringUTF8() gopurs_runtime.Value {
	once_Node_Stream_Aff_toStringUTF8.Do(func() {
		cache_Node_Stream_Aff_toStringUTF8 = gopurs_runtime.Func(func(dictMonadEffect_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Stream_Aff_toStringUTF8(gopurs_runtime.CoerceToStruct[Constructor_Effect_Class_MonadEffect[gopurs_runtime.Value]](dictMonadEffect_0_box))
		})
	})
	return cache_Node_Stream_Aff_toStringUTF8
}

var cache_Node_Stream_Aff_readableToBuffers gopurs_runtime.Value
var once_Node_Stream_Aff_readableToBuffers sync.Once

func Get_Node_Stream_Aff_readableToBuffers() gopurs_runtime.Value {
	once_Node_Stream_Aff_readableToBuffers.Do(func() {
		cache_Node_Stream_Aff_readableToBuffers = gopurs_runtime.Func(func(dictMonadAff_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Stream_Aff_readableToBuffers(gopurs_runtime.CoerceToStruct[Constructor_Effect_Aff_Class_MonadAff[gopurs_runtime.Value]](dictMonadAff_0_box))
		})
	})
	return cache_Node_Stream_Aff_readableToBuffers
}

var cache_Node_Stream_Aff_readableToString gopurs_runtime.Value
var once_Node_Stream_Aff_readableToString sync.Once

func Get_Node_Stream_Aff_readableToString() gopurs_runtime.Value {
	once_Node_Stream_Aff_readableToString.Do(func() {
		cache_Node_Stream_Aff_readableToString = gopurs_runtime.Func(func(dictMonadAff_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Stream_Aff_readableToString(gopurs_runtime.CoerceToStruct[Constructor_Effect_Aff_Class_MonadAff[gopurs_runtime.Value]](dictMonadAff_0_box))
		})
	})
	return cache_Node_Stream_Aff_readableToString
}

var cache_Node_Stream_Aff_readableToStringUtf8 gopurs_runtime.Value
var once_Node_Stream_Aff_readableToStringUtf8 sync.Once

func Get_Node_Stream_Aff_readableToStringUtf8() gopurs_runtime.Value {
	once_Node_Stream_Aff_readableToStringUtf8.Do(func() {
		cache_Node_Stream_Aff_readableToStringUtf8 = gopurs_runtime.Func(func(dictMonadAff_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Stream_Aff_readableToStringUtf8(gopurs_runtime.CoerceToStruct[Constructor_Effect_Aff_Class_MonadAff[gopurs_runtime.Value]](dictMonadAff_0_box))
		})
	})
	return cache_Node_Stream_Aff_readableToStringUtf8
}

var cache_Node_Stream_Aff_readSome gopurs_runtime.Value
var once_Node_Stream_Aff_readSome sync.Once

func Get_Node_Stream_Aff_readSome() gopurs_runtime.Value {
	once_Node_Stream_Aff_readSome.Do(func() {
		cache_Node_Stream_Aff_readSome = gopurs_runtime.Func(func(dictMonadAff_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Stream_Aff_readSome(gopurs_runtime.CoerceToStruct[Constructor_Effect_Aff_Class_MonadAff[gopurs_runtime.Value]](dictMonadAff_0_box))
		})
	})
	return cache_Node_Stream_Aff_readSome
}

var cache_Node_Stream_Aff_readN gopurs_runtime.Value
var once_Node_Stream_Aff_readN sync.Once

func Get_Node_Stream_Aff_readN() gopurs_runtime.Value {
	once_Node_Stream_Aff_readN.Do(func() {
		cache_Node_Stream_Aff_readN = gopurs_runtime.Func(func(dictMonadAff_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Stream_Aff_readN(gopurs_runtime.CoerceToStruct[Constructor_Effect_Aff_Class_MonadAff[gopurs_runtime.Value]](dictMonadAff_0_box))
		})
	})
	return cache_Node_Stream_Aff_readN
}

var cache_Node_Stream_Aff_readAll gopurs_runtime.Value
var once_Node_Stream_Aff_readAll sync.Once

func Get_Node_Stream_Aff_readAll() gopurs_runtime.Value {
	once_Node_Stream_Aff_readAll.Do(func() {
		cache_Node_Stream_Aff_readAll = gopurs_runtime.Func(func(dictMonadAff_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Stream_Aff_readAll(gopurs_runtime.CoerceToStruct[Constructor_Effect_Aff_Class_MonadAff[gopurs_runtime.Value]](dictMonadAff_0_box))
		})
	})
	return cache_Node_Stream_Aff_readAll
}

var cache_Node_Stream_Aff_fromStringUTF8 gopurs_runtime.Value
var once_Node_Stream_Aff_fromStringUTF8 sync.Once

func Get_Node_Stream_Aff_fromStringUTF8() gopurs_runtime.Value {
	once_Node_Stream_Aff_fromStringUTF8.Do(func() {
		cache_Node_Stream_Aff_fromStringUTF8 = gopurs_runtime.Func(func(dictMonadEffect_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Stream_Aff_fromStringUTF8(gopurs_runtime.CoerceToStruct[Constructor_Effect_Class_MonadEffect[gopurs_runtime.Value]](dictMonadEffect_0_box))
		})
	})
	return cache_Node_Stream_Aff_fromStringUTF8
}

var cache_Node_Stream_Aff_end gopurs_runtime.Value
var once_Node_Stream_Aff_end sync.Once

func Get_Node_Stream_Aff_end() gopurs_runtime.Value {
	once_Node_Stream_Aff_end.Do(func() {
		cache_Node_Stream_Aff_end = gopurs_runtime.Func(func(dictMonadAff_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Stream_Aff_end(gopurs_runtime.CoerceToStruct[Constructor_Effect_Aff_Class_MonadAff[gopurs_runtime.Value]](dictMonadAff_0_box))
		})
	})
	return cache_Node_Stream_Aff_end
}

func Call_Node_Stream_Aff_write(dictMonadAff_0_loop *Constructor_Effect_Aff_Class_MonadAff[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictMonadAff_0 *Constructor_Effect_Aff_Class_MonadAff[gopurs_runtime.Value] = dictMonadAff_0_loop
	_ = dictMonadAff_0
	// TAST (Let): liftAff_1_0 shape=App(Var) bindingType=(Func [(ADT ["Effect","Aff","Aff"] [Unit])] (TypeApp (TypeVar m$scope2) [Unit]))
	liftAff_1_0 := Call_Effect_Aff_Class_liftAff(dictMonadAff_0)
	_ = liftAff_1_0
	return gopurs_runtime.Func2(func(w_2 gopurs_runtime.Value, bs_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply3(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), liftAff_1_0, Get_Effect_Aff_makeAff(), gopurs_runtime.Func(func(complete_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
				__local_var_5_1 := gopurs_runtime.Apply2(Get_Effect_Ref__new(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
					return Get_Data_Unit_unit()
				}), gopurs_runtime.Value{})
				_ = __local_var_5_1
				var oneWrite__38428960_6_2_0 gopurs_runtime.Value
				_ = oneWrite__38428960_6_2_0
				var oneWrite__38428960_6_2_0_cell *gopurs_runtime.Value
				_ = oneWrite__38428960_6_2_0_cell
				// FALLBACK TCO: isLoop=false len=2
				var oneWrite_6_3_1 gopurs_runtime.Value
				_ = oneWrite_6_3_1
				var oneWrite_6_3_1_cell *gopurs_runtime.Value
				_ = oneWrite_6_3_1_cell
				// FALLBACK TCO: isLoop=false len=2
				oneWrite__38428960_6_2_0 = gopurs_runtime.Func(func(i_prime__7 gopurs_runtime.Value) gopurs_runtime.Value {
					tailRecEffect_arg_4 := gopurs_runtime.Value{Type: 9, IntVal: 3709389635, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]](Get_Control_Monad_Rec_Class_monadRecEffect()))}
					_ = tailRecEffect_arg_4
					_ = tailRecEffect_arg_4
					tailRecEffect_arg_5 := i_prime__7.IntVal
					_ = tailRecEffect_arg_5
					_ = tailRecEffect_arg_5
					return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
						tailRecEffect_state_6 := tailRecEffect_arg_5
						_ = tailRecEffect_state_6
					tailRecEffect_loop_6:
						for {
							if false {
								continue tailRecEffect_loop_6
							}
							tailRecEffect_parameter_6 := tailRecEffect_state_6
							_ = tailRecEffect_parameter_6
							arrayIndex_arg_8 := Get_Data_Maybe_Just()
							_ = arrayIndex_arg_8
							arrayIndex_arg_9 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
								_v := struct {
									V0 gopurs_runtime.Value
									V1 bool
								}{gopurs_runtime.Value{}, false}
								if _v.V1 {
									return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
								}
								return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
							}())
							_ = arrayIndex_arg_9
							arrayIndex_arg_10 := bs_3
							_ = arrayIndex_arg_10
							arrayIndex_arg_11 := tailRecEffect_parameter_6
							_ = arrayIndex_arg_11
							v_9_7 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
								arrayIndex_value := arrayIndex_arg_10
								_ = arrayIndex_value
								arrayIndex_source_12 := (*[]gopurs_runtime.Value)(arrayIndex_value.UnsafePtr)
								_ = arrayIndex_source_12
								if ((arrayIndex_arg_11) >= (0)) && ((arrayIndex_arg_11) < (int64(len(*arrayIndex_source_12)))) {
									return gopurs_runtime.Apply(arrayIndex_arg_8, (*arrayIndex_source_12)[arrayIndex_arg_11])
								} else {

								}
								return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(arrayIndex_arg_9)}
							}())
							_ = v_9_7
							if v_9_7 == nil {
								__local_var_10_12 := gopurs_runtime.Apply(complete_4, func() gopurs_runtime.Value {
									_v := struct {
										V0 gopurs_runtime.Value
										V1 gopurs_runtime.Value
										V2 bool
									}{gopurs_runtime.Value{}, Get_Data_Unit_unit(), true}
									if _v.V2 {
										return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
									}
									return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
								}())
								_ = __local_var_10_12
								__local_var_11_13 := gopurs_runtime.Apply(__local_var_10_12, gopurs_runtime.Value{})
								_ = __local_var_11_13
								return Get_Data_Unit_unit()
							} else {

							}
							if v_9_7 != nil {
								__local_var_10_14 := Call_Node_Stream_write_prime_(w_2, (v_9_7).V0, gopurs_runtime.Func(func(v1_10 gopurs_runtime.Value) gopurs_runtime.Value {
									var __t15 gopurs_runtime.Value
									{
										if v1_10.Type == 9 && v1_10.IntVal == 930809136 && v1_10.UnsafePtr == nil {
											__t15 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
												return Get_Data_Unit_unit()
											})
											goto end_branch_15
										} else {

										}
									}
									{
										if v1_10.Type == 9 && v1_10.IntVal == 930809136 && v1_10.UnsafePtr != nil {
											__t15 = gopurs_runtime.Apply(complete_4, func() gopurs_runtime.Value {
												_v := struct {
													V0 gopurs_runtime.Value
													V1 gopurs_runtime.Value
													V2 bool
												}{(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v1_10.UnsafePtr).V0, gopurs_runtime.Value{}, false}
												if _v.V2 {
													return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
												}
												return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
											}())
											goto end_branch_15
										} else {

										}
									}
									{
										__t15 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
									}
								end_branch_15:
									return __t15
								}))
								_ = __local_var_10_14
								__local_var_11_16 := gopurs_runtime.Apply(__local_var_10_14, gopurs_runtime.Value{})
								_ = __local_var_11_16
								if (__local_var_11_16.IntVal) != (0) {
									tailRecEffect_state_6 = gopurs_runtime.IntAdd(tailRecEffect_parameter_6, int64(1))
									continue tailRecEffect_loop_6
								} else {

								}
								__local_var_12_17 := gopurs_runtime.Apply(Call_Node_EventEmitter_once(gopurs_runtime.CoerceToStruct[Constructor_Node_EventEmitter_EventHandle[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]](Get_Node_Stream_drainH()), gopurs_runtime.Apply((*oneWrite__38428960_6_2_0_cell), gopurs_runtime.Int(gopurs_runtime.IntAdd(tailRecEffect_parameter_6, int64(1)))), w_2), gopurs_runtime.Value{})
								_ = __local_var_12_17
								__local_var_13_18 := gopurs_runtime.Apply3(Get_Effect_Ref_write(), __local_var_12_17, __local_var_5_1, gopurs_runtime.Value{})
								_ = __local_var_13_18
								return Get_Data_Unit_unit()
							} else {

							}
							panic("Failed pattern match")
						}
					})
				})
				oneWrite__38428960_6_2_0_cell = &oneWrite__38428960_6_2_0
				oneWrite_6_3_1 = gopurs_runtime.Func(func(i_prime__7 gopurs_runtime.Value) gopurs_runtime.Value {
					tailRecEffect_arg_19 := gopurs_runtime.Value{Type: 9, IntVal: 3709389635, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]](Get_Control_Monad_Rec_Class_monadRecEffect()))}
					_ = tailRecEffect_arg_19
					_ = tailRecEffect_arg_19
					tailRecEffect_arg_20 := i_prime__7.IntVal
					_ = tailRecEffect_arg_20
					_ = tailRecEffect_arg_20
					return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
						tailRecEffect_state_21 := tailRecEffect_arg_20
						_ = tailRecEffect_state_21
					tailRecEffect_loop_21:
						for {
							if false {
								continue tailRecEffect_loop_21
							}
							tailRecEffect_parameter_21 := tailRecEffect_state_21
							_ = tailRecEffect_parameter_21
							arrayIndex_arg_23 := Get_Data_Maybe_Just()
							_ = arrayIndex_arg_23
							arrayIndex_arg_24 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
								_v := struct {
									V0 gopurs_runtime.Value
									V1 bool
								}{gopurs_runtime.Value{}, false}
								if _v.V1 {
									return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
								}
								return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
							}())
							_ = arrayIndex_arg_24
							arrayIndex_arg_25 := bs_3
							_ = arrayIndex_arg_25
							arrayIndex_arg_26 := tailRecEffect_parameter_21
							_ = arrayIndex_arg_26
							v_9_22 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
								arrayIndex_value := arrayIndex_arg_25
								_ = arrayIndex_value
								arrayIndex_source_27 := (*[]gopurs_runtime.Value)(arrayIndex_value.UnsafePtr)
								_ = arrayIndex_source_27
								if ((arrayIndex_arg_26) >= (0)) && ((arrayIndex_arg_26) < (int64(len(*arrayIndex_source_27)))) {
									return gopurs_runtime.Apply(arrayIndex_arg_23, (*arrayIndex_source_27)[arrayIndex_arg_26])
								} else {

								}
								return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(arrayIndex_arg_24)}
							}())
							_ = v_9_22
							if v_9_22 == nil {
								__local_var_10_27 := gopurs_runtime.Apply(complete_4, func() gopurs_runtime.Value {
									_v := struct {
										V0 gopurs_runtime.Value
										V1 gopurs_runtime.Value
										V2 bool
									}{gopurs_runtime.Value{}, Get_Data_Unit_unit(), true}
									if _v.V2 {
										return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
									}
									return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
								}())
								_ = __local_var_10_27
								__local_var_11_28 := gopurs_runtime.Apply(__local_var_10_27, gopurs_runtime.Value{})
								_ = __local_var_11_28
								return Get_Data_Unit_unit()
							} else {

							}
							if v_9_22 != nil {
								__local_var_10_29 := Call_Node_Stream_write_prime_(w_2, (v_9_22).V0, gopurs_runtime.Func(func(v1_10 gopurs_runtime.Value) gopurs_runtime.Value {
									var __t30 gopurs_runtime.Value
									{
										if v1_10.Type == 9 && v1_10.IntVal == 930809136 && v1_10.UnsafePtr == nil {
											__t30 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
												return Get_Data_Unit_unit()
											})
											goto end_branch_30
										} else {

										}
									}
									{
										if v1_10.Type == 9 && v1_10.IntVal == 930809136 && v1_10.UnsafePtr != nil {
											__t30 = gopurs_runtime.Apply(complete_4, func() gopurs_runtime.Value {
												_v := struct {
													V0 gopurs_runtime.Value
													V1 gopurs_runtime.Value
													V2 bool
												}{(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v1_10.UnsafePtr).V0, gopurs_runtime.Value{}, false}
												if _v.V2 {
													return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
												}
												return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
											}())
											goto end_branch_30
										} else {

										}
									}
									{
										__t30 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
									}
								end_branch_30:
									return __t30
								}))
								_ = __local_var_10_29
								__local_var_11_31 := gopurs_runtime.Apply(__local_var_10_29, gopurs_runtime.Value{})
								_ = __local_var_11_31
								if (__local_var_11_31.IntVal) != (0) {
									tailRecEffect_state_21 = gopurs_runtime.IntAdd(tailRecEffect_parameter_21, int64(1))
									continue tailRecEffect_loop_21
								} else {

								}
								__local_var_12_32 := gopurs_runtime.Apply(Call_Node_EventEmitter_once(gopurs_runtime.CoerceToStruct[Constructor_Node_EventEmitter_EventHandle[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]](Get_Node_Stream_drainH()), gopurs_runtime.Apply((*oneWrite__38428960_6_2_0_cell), gopurs_runtime.Int(gopurs_runtime.IntAdd(tailRecEffect_parameter_21, int64(1)))), w_2), gopurs_runtime.Value{})
								_ = __local_var_12_32
								__local_var_13_33 := gopurs_runtime.Apply3(Get_Effect_Ref_write(), __local_var_12_32, __local_var_5_1, gopurs_runtime.Value{})
								_ = __local_var_13_33
								return Get_Data_Unit_unit()
							} else {

							}
							panic("Failed pattern match")
						}
					})
				})
				oneWrite_6_3_1_cell = &oneWrite_6_3_1
				__local_var_7_34 := gopurs_runtime.Apply2(oneWrite__38428960_6_2_0, gopurs_runtime.Int(int64(0)), gopurs_runtime.Value{})
				_ = __local_var_7_34
				return gopurs_runtime.Apply(Get_Effect_Aff_effectCanceler(), Call_Control_Bind_join(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect()), gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_5_1)))
			})
		}))
	})
}

func Call_Node_Stream_Aff_toStringUTF8(dictMonadEffect_0_loop *Constructor_Effect_Class_MonadEffect[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictMonadEffect_0 *Constructor_Effect_Class_MonadEffect[gopurs_runtime.Value] = dictMonadEffect_0_loop
	_ = dictMonadEffect_0
	// TAST (Let): liftEffect_1_0 shape=App(Var) bindingType=(Func [(ADT ["Effect","Effect"] [String])] (TypeApp (TypeVar m$scope18) [String]))
	liftEffect_1_0 := Call_Effect_Class_liftEffect(dictMonadEffect_0)
	_ = liftEffect_1_0
	return gopurs_runtime.Func(func(bs_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(liftEffect_1_0, gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			__local_var_3_1 := gopurs_runtime.Apply2(Get_Node_Buffer_concat(), gopurs_runtime.Array(*((*[]gopurs_runtime.Value)(bs_2.UnsafePtr))), gopurs_runtime.Value{})
			_ = __local_var_3_1
			return gopurs_runtime.Apply(Call_Node_Buffer_usingFromImmutable(gopurs_runtime.Apply(Get_Node_Buffer_Immutable_toString(), gopurs_runtime.Value{Type: 9, IntVal: int64(346813718), UnsafePtr: nil}), __local_var_3_1), gopurs_runtime.Value{})
		}))
	})
}

func Call_Node_Stream_Aff_readableToBuffers(dictMonadAff_0_loop *Constructor_Effect_Aff_Class_MonadAff[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictMonadAff_0 *Constructor_Effect_Aff_Class_MonadAff[gopurs_runtime.Value] = dictMonadAff_0_loop
	_ = dictMonadAff_0
	// TAST (Let): liftAff_1_0 shape=App(Var) bindingType=(Func [(ADT ["Effect","Aff","Aff"] [(Array (ADT ["Node","Buffer","Buffer"] []))])] (TypeApp (TypeVar m$scope21) [(Array (ADT ["Node","Buffer","Buffer"] []))]))
	liftAff_1_0 := Call_Effect_Aff_Class_liftAff(dictMonadAff_0)
	_ = liftAff_1_0
	return gopurs_runtime.Func(func(r_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(liftAff_1_0, Call_Effect_Aff_makeAff__1501555834(gopurs_runtime.Func(func(complete_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
				// TAST (Let): __local_var_4_1 shape=App(Var) bindingType=Any
				__local_var_4_1 := gopurs_runtime.Apply(Get_Control_Monad_ST_Global_toEffect(), Get_Data_Array_ST_go__new())
				_ = __local_var_4_1
				bufs_5_2 := gopurs_runtime.Apply(__local_var_4_1, gopurs_runtime.Value{})
				_ = bufs_5_2
				__local_var_6_3 := gopurs_runtime.Apply2(Get_Effect_Ref__new(), Call_Data_Monoid_mempty(Call_Effect_monoidEffect(gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](Get_Data_Monoid_monoidUnit()))})), gopurs_runtime.Value{})
				_ = __local_var_6_3
				// TAST (Let): removeData_7_4 shape=App(Var) bindingType=(ADT ["Effect","Effect"] [Unit])
				removeData_7_4 := Call_Control_Bind_join(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect()), gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_6_3))
				_ = removeData_7_4
				__local_var_8_5 := gopurs_runtime.Apply(Call_Node_EventEmitter_once(gopurs_runtime.CoerceToStruct[Constructor_Node_EventEmitter_EventHandle[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]](Get_Node_Stream_errorH()), gopurs_runtime.Func(func(err_8 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
						__local_var_9_6 := gopurs_runtime.Apply(removeData_7_4, gopurs_runtime.Value{})
						_ = __local_var_9_6
						return gopurs_runtime.Apply2(complete_3, func() gopurs_runtime.Value {
							_v := struct {
								V0 gopurs_runtime.Value
								V1 gopurs_runtime.Value
								V2 bool
							}{err_8, gopurs_runtime.Value{}, false}
							if _v.V2 {
								return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
							}
							return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
						}(), gopurs_runtime.Value{})
					})
				}), r_2), gopurs_runtime.Value{})
				_ = __local_var_8_5
				__local_var_9_7 := gopurs_runtime.Apply(Call_Node_EventEmitter_once(gopurs_runtime.CoerceToStruct[Constructor_Node_EventEmitter_EventHandle[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]](Get_Node_Stream_closeH()), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
					__local_var_9_8 := gopurs_runtime.Apply(__local_var_8_5, gopurs_runtime.Value{})
					_ = __local_var_9_8
					__local_var_10_9 := gopurs_runtime.Apply(removeData_7_4, gopurs_runtime.Value{})
					_ = __local_var_10_9
					__local_var_11_10 := gopurs_runtime.Apply2(Get_Control_Monad_ST_Global_toEffect(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.UncurriedApp(Get_Data_Array_ST_unsafeFreezeImpl(), bufs_5_2)
					}), gopurs_runtime.Value{})
					_ = __local_var_11_10
					return gopurs_runtime.Apply2(complete_3, func() gopurs_runtime.Value {
						_v := struct {
							V0 gopurs_runtime.Value
							V1 gopurs_runtime.Value
							V2 bool
						}{gopurs_runtime.Value{}, __local_var_11_10, true}
						if _v.V2 {
							return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
						}
						return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
					}(), gopurs_runtime.Value{})
				}), r_2), gopurs_runtime.Value{})
				_ = __local_var_9_7
				__local_var_10_11 := gopurs_runtime.Apply(Call_Node_EventEmitter_once(gopurs_runtime.CoerceToStruct[Constructor_Node_EventEmitter_EventHandle[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]](Get_Node_Stream_endH()), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
					__local_var_10_12 := gopurs_runtime.Apply(__local_var_8_5, gopurs_runtime.Value{})
					_ = __local_var_10_12
					__local_var_11_13 := gopurs_runtime.Apply(__local_var_9_7, gopurs_runtime.Value{})
					_ = __local_var_11_13
					__local_var_12_14 := gopurs_runtime.Apply(removeData_7_4, gopurs_runtime.Value{})
					_ = __local_var_12_14
					__local_var_13_15 := gopurs_runtime.Apply2(Get_Control_Monad_ST_Global_toEffect(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.UncurriedApp(Get_Data_Array_ST_unsafeFreezeImpl(), bufs_5_2)
					}), gopurs_runtime.Value{})
					_ = __local_var_13_15
					return gopurs_runtime.Apply2(complete_3, func() gopurs_runtime.Value {
						_v := struct {
							V0 gopurs_runtime.Value
							V1 gopurs_runtime.Value
							V2 bool
						}{gopurs_runtime.Value{}, __local_var_13_15, true}
						if _v.V2 {
							return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
						}
						return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
					}(), gopurs_runtime.Value{})
				}), r_2), gopurs_runtime.Value{})
				_ = __local_var_10_11
				__local_var_11_16 := gopurs_runtime.Apply(Call_Node_EventEmitter_on(gopurs_runtime.CoerceToStruct[Constructor_Node_EventEmitter_EventHandle[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]](Get_Node_Stream_dataH()), gopurs_runtime.Func(func(buf_11 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply(Call_Data_Functor_void(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect())), gopurs_runtime.Apply(Get_Control_Monad_ST_Global_toEffect(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.UncurriedApp2(Get_Data_Array_ST_pushImpl(), buf_11, bufs_5_2)
					})))
				}), r_2), gopurs_runtime.Value{})
				_ = __local_var_11_16
				__local_var_12_17 := gopurs_runtime.Apply3(Get_Effect_Ref_write(), __local_var_11_16, __local_var_6_3, gopurs_runtime.Value{})
				_ = __local_var_12_17
				return gopurs_runtime.Apply(Get_Effect_Aff_effectCanceler(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
					__local_var_13_18 := gopurs_runtime.Apply(__local_var_8_5, gopurs_runtime.Value{})
					_ = __local_var_13_18
					__local_var_14_19 := gopurs_runtime.Apply(__local_var_9_7, gopurs_runtime.Value{})
					_ = __local_var_14_19
					__local_var_15_20 := gopurs_runtime.Apply(__local_var_10_11, gopurs_runtime.Value{})
					_ = __local_var_15_20
					return gopurs_runtime.Apply(removeData_7_4, gopurs_runtime.Value{})
				}))
			})
		})))
	})
}

func Call_Node_Stream_Aff_readableToString(dictMonadAff_0_loop *Constructor_Effect_Aff_Class_MonadAff[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictMonadAff_0 *Constructor_Effect_Aff_Class_MonadAff[gopurs_runtime.Value] = dictMonadAff_0_loop
	_ = dictMonadAff_0
	// TAST (Let): MonadEffect0_1_0 shape=App(Other) bindingType=Any
	MonadEffect0_1_0 := gopurs_runtime.Apply(dictMonadAff_0.V0, gopurs_runtime.Value{})
	_ = MonadEffect0_1_0
	// TAST (Let): readableToBuffers1__1193806748_2_1 shape=App(Var) bindingType=(Func [(ADT ["Node","Stream","Stream"] [(Row [read: (ADT ["Node","Stream","Read"] [])] (TypeVar w$scope34))])] (TypeApp (TypeVar m$scope33) [(Array (ADT ["Node","Buffer","Buffer"] []))]))
	readableToBuffers1__1193806748_2_1 := Call_Node_Stream_Aff_readableToBuffers(dictMonadAff_0)
	_ = readableToBuffers1__1193806748_2_1
	// TAST (Let): liftEffect_3_2 shape=App(Var) bindingType=(Func [(ADT ["Effect","Effect"] [String])] (TypeApp (TypeVar m$scope33) [String]))
	liftEffect_3_2 := Call_Effect_Class_liftEffect(gopurs_runtime.CoerceToStruct[Constructor_Effect_Class_MonadEffect[gopurs_runtime.Value]](MonadEffect0_1_0))
	_ = liftEffect_3_2
	return gopurs_runtime.Func2(func(r_4 gopurs_runtime.Value, enc_5 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(MonadEffect0_1_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply(readableToBuffers1__1193806748_2_1, r_4), gopurs_runtime.Func(func(bufs_6 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(liftEffect_3_2, gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
				__local_var_7_3 := gopurs_runtime.Apply2(Get_Node_Buffer_concat(), gopurs_runtime.Array(*((*[]gopurs_runtime.Value)(bufs_6.UnsafePtr))), gopurs_runtime.Value{})
				_ = __local_var_7_3
				return gopurs_runtime.Apply(Call_Node_Buffer_usingFromImmutable(gopurs_runtime.Apply(Get_Node_Buffer_Immutable_toString(), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(enc_5.IntVal)), UnsafePtr: nil}), __local_var_7_3), gopurs_runtime.Value{})
			}))
		}))
	})
}

func Call_Node_Stream_Aff_readableToStringUtf8(dictMonadAff_0_loop *Constructor_Effect_Aff_Class_MonadAff[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictMonadAff_0 *Constructor_Effect_Aff_Class_MonadAff[gopurs_runtime.Value] = dictMonadAff_0_loop
	_ = dictMonadAff_0
	// TAST (Let): readableToString1__2066326960_1_0 shape=App(Var) bindingType=(Func [(ADT ["Node","Stream","Stream"] [(Row [read: (ADT ["Node","Stream","Read"] [])] (TypeVar w$scope38))]), (ADT ["Node","Encoding","Encoding"] [])] (TypeApp (TypeVar m$scope37) [String]))
	readableToString1__2066326960_1_0 := Call_Node_Stream_Aff_readableToString(dictMonadAff_0)
	_ = readableToString1__2066326960_1_0
	return gopurs_runtime.Func(func(r_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(readableToString1__2066326960_1_0, r_2, gopurs_runtime.Value{Type: 9, IntVal: int64(346813718), UnsafePtr: nil})
	})
}

func Call_Node_Stream_Aff_readSome(dictMonadAff_0_loop *Constructor_Effect_Aff_Class_MonadAff[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictMonadAff_0 *Constructor_Effect_Aff_Class_MonadAff[gopurs_runtime.Value] = dictMonadAff_0_loop
	_ = dictMonadAff_0
	// TAST (Let): liftAff_1_0 shape=App(Var) bindingType=(Func [(ADT ["Effect","Aff","Aff"] [(Record (Row [buffers: (Array (ADT ["Node","Buffer","Buffer"] [])), readagain: Boolean] Empty))])] (TypeApp (TypeVar m$scope41) [(Record (Row [buffers: (Array (ADT ["Node","Buffer","Buffer"] [])), readagain: Boolean] Empty))]))
	liftAff_1_0 := Call_Effect_Aff_Class_liftAff(dictMonadAff_0)
	_ = liftAff_1_0
	return gopurs_runtime.Func(func(r_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply3(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), liftAff_1_0, Get_Effect_Aff_makeAff(), gopurs_runtime.Func(func(complete_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
				__local_var_4_1 := gopurs_runtime.UncurriedApp(Get_Node_Stream_readableImpl(), r_2)
				_ = __local_var_4_1
				var __t46 gopurs_runtime.Value
				{
					if (__local_var_4_1.IntVal) != (0) {
						__t46 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
							bufs_5_3 := gopurs_runtime.Apply2(Get_Control_Monad_ST_Global_toEffect(), Get_Data_Array_ST_go__new(), gopurs_runtime.Value{})
							_ = bufs_5_3
							__local_var_6_4 := gopurs_runtime.Apply(Call_Node_EventEmitter_once(gopurs_runtime.CoerceToStruct[Constructor_Node_EventEmitter_EventHandle[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]](Get_Node_Stream_errorH()), gopurs_runtime.Func(func(err_6 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Apply(complete_3, func() gopurs_runtime.Value {
									_v := struct {
										V0 gopurs_runtime.Value
										V1 struct {
											buffers   []gopurs_runtime.Value
											readagain bool
											/*gopurs:record-v1:2:0062007500660066006500720073,00720065006100640061006700610069006e*/
										}
										V2 bool
									}{err_6, struct {
										buffers   []gopurs_runtime.Value
										readagain bool
										/*gopurs:record-v1:2:0062007500660066006500720073,00720065006100640061006700610069006e*/
									}{}, false}
									if _v.V2 {
										return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: func() gopurs_runtime.Value {
											orig := _v.V1
											_ = orig
											return gopurs_runtime.RecordDict2("buffers", "readagain", gopurs_runtime.Array(orig.buffers), gopurs_runtime.Bool(orig.readagain))
										}()})}
									}
									return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
								}())
							}), r_2), gopurs_runtime.Value{})
							_ = __local_var_6_4
							__local_var_7_5 := gopurs_runtime.Apply(Call_Node_EventEmitter_once(gopurs_runtime.CoerceToStruct[Constructor_Node_EventEmitter_EventHandle[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]](Get_Node_Stream_closeH()), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
								__local_var_7_6 := gopurs_runtime.Apply(__local_var_6_4, gopurs_runtime.Value{})
								_ = __local_var_7_6
								__local_var_8_7 := gopurs_runtime.Apply2(Get_Control_Monad_ST_Global_toEffect(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.UncurriedApp(Get_Data_Array_ST_unsafeFreezeImpl(), bufs_5_3)
								}), gopurs_runtime.Value{})
								_ = __local_var_8_7
								return gopurs_runtime.Apply2(complete_3, func() gopurs_runtime.Value {
									_v := struct {
										V0 gopurs_runtime.Value
										V1 struct {
											buffers   []gopurs_runtime.Value
											readagain bool
											/*gopurs:record-v1:2:0062007500660066006500720073,00720065006100640061006700610069006e*/
										}
										V2 bool
									}{gopurs_runtime.Value{}, struct {
										buffers   []gopurs_runtime.Value
										readagain bool
										/*gopurs:record-v1:2:0062007500660066006500720073,00720065006100640061006700610069006e*/
									}{*((*[]gopurs_runtime.Value)(__local_var_8_7.UnsafePtr)), false}, true}
									if _v.V2 {
										return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: func() gopurs_runtime.Value {
											orig := _v.V1
											_ = orig
											return gopurs_runtime.RecordDict2("buffers", "readagain", gopurs_runtime.Array(orig.buffers), gopurs_runtime.Bool(orig.readagain))
										}()})}
									}
									return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
								}(), gopurs_runtime.Value{})
							}), r_2), gopurs_runtime.Value{})
							_ = __local_var_7_5
							__local_var_8_8 := gopurs_runtime.Apply(Call_Node_EventEmitter_once(gopurs_runtime.CoerceToStruct[Constructor_Node_EventEmitter_EventHandle[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]](Get_Node_Stream_endH()), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
								__local_var_8_9 := gopurs_runtime.Apply(__local_var_6_4, gopurs_runtime.Value{})
								_ = __local_var_8_9
								__local_var_9_10 := gopurs_runtime.Apply(__local_var_7_5, gopurs_runtime.Value{})
								_ = __local_var_9_10
								__local_var_10_11 := gopurs_runtime.Apply2(Get_Control_Monad_ST_Global_toEffect(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.UncurriedApp(Get_Data_Array_ST_unsafeFreezeImpl(), bufs_5_3)
								}), gopurs_runtime.Value{})
								_ = __local_var_10_11
								return gopurs_runtime.Apply2(complete_3, func() gopurs_runtime.Value {
									_v := struct {
										V0 gopurs_runtime.Value
										V1 struct {
											buffers   []gopurs_runtime.Value
											readagain bool
											/*gopurs:record-v1:2:0062007500660066006500720073,00720065006100640061006700610069006e*/
										}
										V2 bool
									}{gopurs_runtime.Value{}, struct {
										buffers   []gopurs_runtime.Value
										readagain bool
										/*gopurs:record-v1:2:0062007500660066006500720073,00720065006100640061006700610069006e*/
									}{*((*[]gopurs_runtime.Value)(__local_var_10_11.UnsafePtr)), false}, true}
									if _v.V2 {
										return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: func() gopurs_runtime.Value {
											orig := _v.V1
											_ = orig
											return gopurs_runtime.RecordDict2("buffers", "readagain", gopurs_runtime.Array(orig.buffers), gopurs_runtime.Bool(orig.readagain))
										}()})}
									}
									return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
								}(), gopurs_runtime.Value{})
							}), r_2), gopurs_runtime.Value{})
							_ = __local_var_8_8
							return gopurs_runtime.Apply3(Get_Effect_Exception_catchException(), gopurs_runtime.Func(func(err_9 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
									__local_var_10_12 := gopurs_runtime.Apply(__local_var_6_4, gopurs_runtime.Value{})
									_ = __local_var_10_12
									__local_var_11_13 := gopurs_runtime.Apply(__local_var_7_5, gopurs_runtime.Value{})
									_ = __local_var_11_13
									__local_var_12_14 := gopurs_runtime.Apply(__local_var_8_8, gopurs_runtime.Value{})
									_ = __local_var_12_14
									__local_var_13_15 := gopurs_runtime.Apply2(complete_3, func() gopurs_runtime.Value {
										_v := struct {
											V0 gopurs_runtime.Value
											V1 struct {
												buffers   []gopurs_runtime.Value
												readagain bool
												/*gopurs:record-v1:2:0062007500660066006500720073,00720065006100640061006700610069006e*/
											}
											V2 bool
										}{err_9, struct {
											buffers   []gopurs_runtime.Value
											readagain bool
											/*gopurs:record-v1:2:0062007500660066006500720073,00720065006100640061006700610069006e*/
										}{}, false}
										if _v.V2 {
											return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: func() gopurs_runtime.Value {
												orig := _v.V1
												_ = orig
												return gopurs_runtime.RecordDict2("buffers", "readagain", gopurs_runtime.Array(orig.buffers), gopurs_runtime.Bool(orig.readagain))
											}()})}
										}
										return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
									}(), gopurs_runtime.Value{})
									_ = __local_var_13_15
									return gopurs_runtime.Apply(Get_Data_Function_const__4207166068(), gopurs_runtime.Apply(Get_Effect_Aff__pure(), Get_Data_Unit_unit()))
								})
							}), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
								// TAST (Let): __local_var_9_16 shape=App(Var) bindingType=Any
								__local_var_9_16 := gopurs_runtime.Apply(Get_Effect_untilE(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
									// TAST (Let): __local_var_9_17 shape=App(Var) bindingType=Any
									__local_var_9_17 := Call_Node_Stream_read(r_2)
									_ = __local_var_9_17
									__local_var_10_18 := gopurs_runtime.Apply(__local_var_9_17, gopurs_runtime.Value{})
									_ = __local_var_10_18
									var __t21 gopurs_runtime.Value
									{
										if __local_var_10_18.Type == 9 && __local_var_10_18.IntVal == 930809136 && __local_var_10_18.UnsafePtr == nil {
											__t21 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
												return gopurs_runtime.Bool(true)
											})
											goto end_branch_21
										} else {

										}
									}
									{
										if __local_var_10_18.Type == 9 && __local_var_10_18.IntVal == 930809136 && __local_var_10_18.UnsafePtr != nil {
											__t21 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
												// TAST (Let): __local_var_11_19 shape=Other bindingType=Any
												__local_var_11_19 := (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_10_18.UnsafePtr).V0
												_ = __local_var_11_19
												__local_var_12_20 := gopurs_runtime.Apply2(Call_Data_Functor_void(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect())), gopurs_runtime.Apply(Get_Control_Monad_ST_Global_toEffect(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
													return gopurs_runtime.UncurriedApp2(Get_Data_Array_ST_pushImpl(), __local_var_11_19, bufs_5_3)
												})), gopurs_runtime.Value{})
												_ = __local_var_12_20
												return gopurs_runtime.Bool(false)
											})
											goto end_branch_21
										} else {

										}
									}
									{
										__t21 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
									}
								end_branch_21:
									return gopurs_runtime.Apply(__t21, gopurs_runtime.Value{})
								}))
								_ = __local_var_9_16
								__local_var_10_22 := gopurs_runtime.Apply(__local_var_9_16, gopurs_runtime.Value{})
								_ = __local_var_10_22
								__local_var_11_23 := gopurs_runtime.Apply2(Get_Control_Monad_ST_Global_toEffect(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.UncurriedApp(Get_Data_Array_ST_unsafeFreezeImpl(), bufs_5_3)
								}), gopurs_runtime.Value{})
								_ = __local_var_11_23
								__local_var_12_24 := gopurs_runtime.UncurriedApp(Get_Node_Stream_readableImpl(), r_2)
								_ = __local_var_12_24
								var __t45 gopurs_runtime.Value
								{
									if ((__local_var_12_24.IntVal) != (0)) && ((gopurs_runtime.Int(int64(gopurs_runtime.ArrayLength(__local_var_11_23))).IntVal) == (int64(0))) {
										__t45 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
											__local_var_13_29 := gopurs_runtime.Apply(Call_Node_EventEmitter_once(gopurs_runtime.CoerceToStruct[Constructor_Node_EventEmitter_EventHandle[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]](Get_Node_Stream_readableH()), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
												// TAST (Let): __local_var_13_30 shape=App(Var) bindingType=Any
												__local_var_13_30 := gopurs_runtime.Apply(Get_Effect_untilE(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
													// TAST (Let): __local_var_13_31 shape=App(Var) bindingType=Any
													__local_var_13_31 := Call_Node_Stream_read(r_2)
													_ = __local_var_13_31
													__local_var_14_32 := gopurs_runtime.Apply(__local_var_13_31, gopurs_runtime.Value{})
													_ = __local_var_14_32
													var __t35 gopurs_runtime.Value
													{
														if __local_var_14_32.Type == 9 && __local_var_14_32.IntVal == 930809136 && __local_var_14_32.UnsafePtr == nil {
															__t35 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
																return gopurs_runtime.Bool(true)
															})
															goto end_branch_35
														} else {

														}
													}
													{
														if __local_var_14_32.Type == 9 && __local_var_14_32.IntVal == 930809136 && __local_var_14_32.UnsafePtr != nil {
															__t35 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
																// TAST (Let): __local_var_15_33 shape=Other bindingType=Any
																__local_var_15_33 := (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_14_32.UnsafePtr).V0
																_ = __local_var_15_33
																__local_var_16_34 := gopurs_runtime.Apply2(Call_Data_Functor_void(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect())), gopurs_runtime.Apply(Get_Control_Monad_ST_Global_toEffect(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
																	return gopurs_runtime.UncurriedApp2(Get_Data_Array_ST_pushImpl(), __local_var_15_33, bufs_5_3)
																})), gopurs_runtime.Value{})
																_ = __local_var_16_34
																return gopurs_runtime.Bool(false)
															})
															goto end_branch_35
														} else {

														}
													}
													{
														__t35 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
													}
												end_branch_35:
													return gopurs_runtime.Apply(__t35, gopurs_runtime.Value{})
												}))
												_ = __local_var_13_30
												__local_var_14_36 := gopurs_runtime.Apply(__local_var_13_30, gopurs_runtime.Value{})
												_ = __local_var_14_36
												__local_var_15_37 := gopurs_runtime.Apply2(Get_Control_Monad_ST_Global_toEffect(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
													return gopurs_runtime.UncurriedApp(Get_Data_Array_ST_unsafeFreezeImpl(), bufs_5_3)
												}), gopurs_runtime.Value{})
												_ = __local_var_15_37
												__local_var_16_38 := gopurs_runtime.Apply(__local_var_6_4, gopurs_runtime.Value{})
												_ = __local_var_16_38
												__local_var_17_39 := gopurs_runtime.Apply(__local_var_7_5, gopurs_runtime.Value{})
												_ = __local_var_17_39
												__local_var_18_40 := gopurs_runtime.Apply(__local_var_8_8, gopurs_runtime.Value{})
												_ = __local_var_18_40
												__local_var_19_41 := gopurs_runtime.UncurriedApp(Get_Node_Stream_readableImpl(), r_2)
												_ = __local_var_19_41
												return gopurs_runtime.Apply2(complete_3, func() gopurs_runtime.Value {
													_v := struct {
														V0 gopurs_runtime.Value
														V1 struct {
															buffers   []gopurs_runtime.Value
															readagain bool
															/*gopurs:record-v1:2:0062007500660066006500720073,00720065006100640061006700610069006e*/
														}
														V2 bool
													}{gopurs_runtime.Value{}, struct {
														buffers   []gopurs_runtime.Value
														readagain bool
														/*gopurs:record-v1:2:0062007500660066006500720073,00720065006100640061006700610069006e*/
													}{*((*[]gopurs_runtime.Value)(__local_var_15_37.UnsafePtr)), (__local_var_19_41.IntVal) != (0)}, true}
													if _v.V2 {
														return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: func() gopurs_runtime.Value {
															orig := _v.V1
															_ = orig
															return gopurs_runtime.RecordDict2("buffers", "readagain", gopurs_runtime.Array(orig.buffers), gopurs_runtime.Bool(orig.readagain))
														}()})}
													}
													return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
												}(), gopurs_runtime.Value{})
											}), r_2), gopurs_runtime.Value{})
											_ = __local_var_13_29
											return gopurs_runtime.Apply(Get_Effect_Aff_effectCanceler(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
												__local_var_14_42 := gopurs_runtime.Apply(__local_var_6_4, gopurs_runtime.Value{})
												_ = __local_var_14_42
												__local_var_15_43 := gopurs_runtime.Apply(__local_var_7_5, gopurs_runtime.Value{})
												_ = __local_var_15_43
												__local_var_16_44 := gopurs_runtime.Apply(__local_var_8_8, gopurs_runtime.Value{})
												_ = __local_var_16_44
												return gopurs_runtime.Apply(__local_var_13_29, gopurs_runtime.Value{})
											}))
										})
										goto end_branch_45
									} else {

									}
								}
								{
									__t45 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
										__local_var_13_25 := gopurs_runtime.Apply(__local_var_6_4, gopurs_runtime.Value{})
										_ = __local_var_13_25
										__local_var_14_26 := gopurs_runtime.Apply(__local_var_7_5, gopurs_runtime.Value{})
										_ = __local_var_14_26
										__local_var_15_27 := gopurs_runtime.Apply(__local_var_8_8, gopurs_runtime.Value{})
										_ = __local_var_15_27
										__local_var_16_28 := gopurs_runtime.Apply2(complete_3, func() gopurs_runtime.Value {
											_v := struct {
												V0 gopurs_runtime.Value
												V1 struct {
													buffers   []gopurs_runtime.Value
													readagain bool
													/*gopurs:record-v1:2:0062007500660066006500720073,00720065006100640061006700610069006e*/
												}
												V2 bool
											}{gopurs_runtime.Value{}, struct {
												buffers   []gopurs_runtime.Value
												readagain bool
												/*gopurs:record-v1:2:0062007500660066006500720073,00720065006100640061006700610069006e*/
											}{*((*[]gopurs_runtime.Value)(__local_var_11_23.UnsafePtr)), (__local_var_12_24.IntVal) != (0)}, true}
											if _v.V2 {
												return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: func() gopurs_runtime.Value {
													orig := _v.V1
													_ = orig
													return gopurs_runtime.RecordDict2("buffers", "readagain", gopurs_runtime.Array(orig.buffers), gopurs_runtime.Bool(orig.readagain))
												}()})}
											}
											return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
										}(), gopurs_runtime.Value{})
										_ = __local_var_16_28
										return gopurs_runtime.Apply(Get_Data_Function_const__4207166068(), gopurs_runtime.Apply(Get_Effect_Aff__pure(), Get_Data_Unit_unit()))
									})
								}
							end_branch_45:
								return gopurs_runtime.Apply(__t45, gopurs_runtime.Value{})
							}), gopurs_runtime.Value{})
						})
						goto end_branch_46
					} else {

					}
				}
				{
					__t46 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
						__local_var_5_2 := gopurs_runtime.Apply2(complete_3, func() gopurs_runtime.Value {
							_v := struct {
								V0 gopurs_runtime.Value
								V1 struct {
									buffers   []gopurs_runtime.Value
									readagain bool
									/*gopurs:record-v1:2:0062007500660066006500720073,00720065006100640061006700610069006e*/
								}
								V2 bool
							}{gopurs_runtime.Value{}, struct {
								buffers   []gopurs_runtime.Value
								readagain bool
								/*gopurs:record-v1:2:0062007500660066006500720073,00720065006100640061006700610069006e*/
							}{*((*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{}).UnsafePtr)), false}, true}
							if _v.V2 {
								return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: func() gopurs_runtime.Value {
									orig := _v.V1
									_ = orig
									return gopurs_runtime.RecordDict2("buffers", "readagain", gopurs_runtime.Array(orig.buffers), gopurs_runtime.Bool(orig.readagain))
								}()})}
							}
							return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
						}(), gopurs_runtime.Value{})
						_ = __local_var_5_2
						return gopurs_runtime.Apply(Get_Data_Function_const__4207166068(), gopurs_runtime.Apply(Get_Effect_Aff__pure(), Get_Data_Unit_unit()))
					})
				}
			end_branch_46:
				return gopurs_runtime.Apply(__t46, gopurs_runtime.Value{})
			})
		}))
	})
}

func Call_Node_Stream_Aff_readN(dictMonadAff_0_loop *Constructor_Effect_Aff_Class_MonadAff[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictMonadAff_0 *Constructor_Effect_Aff_Class_MonadAff[gopurs_runtime.Value] = dictMonadAff_0_loop
	_ = dictMonadAff_0
	// TAST (Let): liftAff_1_0 shape=App(Var) bindingType=(Func [(ADT ["Effect","Aff","Aff"] [(Record (Row [buffers: (Array (ADT ["Node","Buffer","Buffer"] [])), readagain: Boolean] Empty))])] (TypeApp (TypeVar m$scope62) [(Record (Row [buffers: (Array (ADT ["Node","Buffer","Buffer"] [])), readagain: Boolean] Empty))]))
	liftAff_1_0 := Call_Effect_Aff_Class_liftAff(dictMonadAff_0)
	_ = liftAff_1_0
	return gopurs_runtime.Func2(func(r_2 gopurs_runtime.Value, n_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply3(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), liftAff_1_0, Get_Effect_Aff_makeAff(), gopurs_runtime.Func(func(complete_4 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t50 gopurs_runtime.Value
			{
				if (n_3.IntVal) < (int64(0)) {
					__t50 = Call_Control_Apply_applySecond__3440890937(gopurs_runtime.Apply(complete_4, func() gopurs_runtime.Value {
						_v := struct {
							V0 gopurs_runtime.Value
							V1 struct {
								buffers   []gopurs_runtime.Value
								readagain bool
								/*gopurs:record-v1:2:0062007500660066006500720073,00720065006100640061006700610069006e*/
							}
							V2 bool
						}{gopurs_runtime.Apply(Get_Effect_Exception_error(), gopurs_runtime.StrWithCodeUnits("read bytes must be > 0", 22)), struct {
							buffers   []gopurs_runtime.Value
							readagain bool
							/*gopurs:record-v1:2:0062007500660066006500720073,00720065006100640061006700610069006e*/
						}{}, false}
						if _v.V2 {
							return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: func() gopurs_runtime.Value {
								orig := _v.V1
								_ = orig
								return gopurs_runtime.RecordDict2("buffers", "readagain", gopurs_runtime.Array(orig.buffers), gopurs_runtime.Bool(orig.readagain))
							}()})}
						}
						return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
					}()), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
						// TAST (Let): __local_var_5_49 shape=App(Var) bindingType=Any
						__local_var_5_49 := gopurs_runtime.Apply(Get_Data_Function_const__4207166068(), gopurs_runtime.Apply(Get_Effect_Aff__pure(), Get_Data_Unit_unit()))
						_ = __local_var_5_49
						return __local_var_5_49
					}))
					goto end_branch_50
				} else {

				}
			}
			{
				__t50 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
					__local_var_5_1 := gopurs_runtime.UncurriedApp(Get_Node_Stream_readableImpl(), r_2)
					_ = __local_var_5_1
					var __t48 gopurs_runtime.Value
					{
						if (__local_var_5_1.IntVal) != (0) {
							__t48 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
								__local_var_6_3 := gopurs_runtime.Apply2(Get_Effect_Ref__new(), gopurs_runtime.Int(int64(0)), gopurs_runtime.Value{})
								_ = __local_var_6_3
								bufs_7_4 := gopurs_runtime.Apply2(Get_Control_Monad_ST_Global_toEffect(), Get_Data_Array_ST_go__new(), gopurs_runtime.Value{})
								_ = bufs_7_4
								__local_var_8_5 := gopurs_runtime.Apply2(Get_Effect_Ref__new(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
									return Get_Data_Unit_unit()
								}), gopurs_runtime.Value{})
								_ = __local_var_8_5
								__local_var_9_6 := gopurs_runtime.Apply(Call_Node_EventEmitter_once(gopurs_runtime.CoerceToStruct[Constructor_Node_EventEmitter_EventHandle[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]](Get_Node_Stream_errorH()), gopurs_runtime.Func(func(err_9 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
										__local_var_10_7 := gopurs_runtime.Apply(Call_Control_Bind_join(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect()), gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_8_5)), gopurs_runtime.Value{})
										_ = __local_var_10_7
										return gopurs_runtime.Apply2(complete_4, func() gopurs_runtime.Value {
											_v := struct {
												V0 gopurs_runtime.Value
												V1 struct {
													buffers   []gopurs_runtime.Value
													readagain bool
													/*gopurs:record-v1:2:0062007500660066006500720073,00720065006100640061006700610069006e*/
												}
												V2 bool
											}{err_9, struct {
												buffers   []gopurs_runtime.Value
												readagain bool
												/*gopurs:record-v1:2:0062007500660066006500720073,00720065006100640061006700610069006e*/
											}{}, false}
											if _v.V2 {
												return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: func() gopurs_runtime.Value {
													orig := _v.V1
													_ = orig
													return gopurs_runtime.RecordDict2("buffers", "readagain", gopurs_runtime.Array(orig.buffers), gopurs_runtime.Bool(orig.readagain))
												}()})}
											}
											return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
										}(), gopurs_runtime.Value{})
									})
								}), r_2), gopurs_runtime.Value{})
								_ = __local_var_9_6
								__local_var_10_8 := gopurs_runtime.Apply(Call_Node_EventEmitter_once(gopurs_runtime.CoerceToStruct[Constructor_Node_EventEmitter_EventHandle[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]](Get_Node_Stream_closeH()), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
									__local_var_10_9 := gopurs_runtime.Apply(__local_var_9_6, gopurs_runtime.Value{})
									_ = __local_var_10_9
									__local_var_11_10 := gopurs_runtime.Apply(Call_Control_Bind_join(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect()), gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_8_5)), gopurs_runtime.Value{})
									_ = __local_var_11_10
									__local_var_12_11 := gopurs_runtime.Apply2(Get_Control_Monad_ST_Global_toEffect(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.UncurriedApp(Get_Data_Array_ST_unsafeFreezeImpl(), bufs_7_4)
									}), gopurs_runtime.Value{})
									_ = __local_var_12_11
									return gopurs_runtime.Apply2(complete_4, func() gopurs_runtime.Value {
										_v := struct {
											V0 gopurs_runtime.Value
											V1 struct {
												buffers   []gopurs_runtime.Value
												readagain bool
												/*gopurs:record-v1:2:0062007500660066006500720073,00720065006100640061006700610069006e*/
											}
											V2 bool
										}{gopurs_runtime.Value{}, struct {
											buffers   []gopurs_runtime.Value
											readagain bool
											/*gopurs:record-v1:2:0062007500660066006500720073,00720065006100640061006700610069006e*/
										}{*((*[]gopurs_runtime.Value)(__local_var_12_11.UnsafePtr)), false}, true}
										if _v.V2 {
											return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: func() gopurs_runtime.Value {
												orig := _v.V1
												_ = orig
												return gopurs_runtime.RecordDict2("buffers", "readagain", gopurs_runtime.Array(orig.buffers), gopurs_runtime.Bool(orig.readagain))
											}()})}
										}
										return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
									}(), gopurs_runtime.Value{})
								}), r_2), gopurs_runtime.Value{})
								_ = __local_var_10_8
								__local_var_11_12 := gopurs_runtime.Apply(Call_Node_EventEmitter_once(gopurs_runtime.CoerceToStruct[Constructor_Node_EventEmitter_EventHandle[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]](Get_Node_Stream_endH()), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
									__local_var_11_13 := gopurs_runtime.Apply(__local_var_9_6, gopurs_runtime.Value{})
									_ = __local_var_11_13
									__local_var_12_14 := gopurs_runtime.Apply(__local_var_10_8, gopurs_runtime.Value{})
									_ = __local_var_12_14
									__local_var_13_15 := gopurs_runtime.Apply2(Get_Control_Monad_ST_Global_toEffect(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.UncurriedApp(Get_Data_Array_ST_unsafeFreezeImpl(), bufs_7_4)
									}), gopurs_runtime.Value{})
									_ = __local_var_13_15
									return gopurs_runtime.Apply2(complete_4, func() gopurs_runtime.Value {
										_v := struct {
											V0 gopurs_runtime.Value
											V1 struct {
												buffers   []gopurs_runtime.Value
												readagain bool
												/*gopurs:record-v1:2:0062007500660066006500720073,00720065006100640061006700610069006e*/
											}
											V2 bool
										}{gopurs_runtime.Value{}, struct {
											buffers   []gopurs_runtime.Value
											readagain bool
											/*gopurs:record-v1:2:0062007500660066006500720073,00720065006100640061006700610069006e*/
										}{*((*[]gopurs_runtime.Value)(__local_var_13_15.UnsafePtr)), false}, true}
										if _v.V2 {
											return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: func() gopurs_runtime.Value {
												orig := _v.V1
												_ = orig
												return gopurs_runtime.RecordDict2("buffers", "readagain", gopurs_runtime.Array(orig.buffers), gopurs_runtime.Bool(orig.readagain))
											}()})}
										}
										return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
									}(), gopurs_runtime.Value{})
								}), r_2), gopurs_runtime.Value{})
								_ = __local_var_11_12
								var Call_local_Node_Stream_Aff_tryToRead__808114339_12_16 func(gopurs_runtime.Value) gopurs_runtime.Value
								_ = Call_local_Node_Stream_Aff_tryToRead__808114339_12_16
								var tryToRead__808114339_12_16 gopurs_runtime.Value
								_ = tryToRead__808114339_12_16
								Call_local_Node_Stream_Aff_tryToRead__808114339_12_16 = func(continuation_12_loop gopurs_runtime.Value) gopurs_runtime.Value {
									var continuation_12 gopurs_runtime.Value = continuation_12_loop
									_ = continuation_12
									return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
										// TAST (Let): __local_var_13_17 shape=App(Var) bindingType=Any
										__local_var_13_17 := gopurs_runtime.Apply(Get_Effect_untilE(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
											// TAST (Let): __local_var_13_18 shape=App(Var) bindingType=Any
											__local_var_13_18 := gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_6_3)
											_ = __local_var_13_18
											__local_var_14_19 := gopurs_runtime.Apply(__local_var_13_18, gopurs_runtime.Value{})
											_ = __local_var_14_19
											__local_var_15_20 := gopurs_runtime.Apply(Call_Node_Stream_read_prime_(r_2, gopurs_runtime.IntSub(n_3.IntVal, __local_var_14_19.IntVal)), gopurs_runtime.Value{})
											_ = __local_var_15_20
											var __t27 gopurs_runtime.Value
											{
												if __local_var_15_20.Type == 9 && __local_var_15_20.IntVal == 930809136 && __local_var_15_20.UnsafePtr == nil {
													__t27 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
														return gopurs_runtime.Bool(true)
													})
													goto end_branch_27
												} else {

												}
											}
											{
												if __local_var_15_20.Type == 9 && __local_var_15_20.IntVal == 930809136 && __local_var_15_20.UnsafePtr != nil {
													__t27 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
														// TAST (Let): __local_var_16_21 shape=Other bindingType=Any
														__local_var_16_21 := (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_15_20.UnsafePtr).V0
														_ = __local_var_16_21
														__local_var_17_22 := gopurs_runtime.Apply2(Get_Control_Monad_ST_Global_toEffect(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
															return gopurs_runtime.UncurriedApp2(Get_Data_Array_ST_pushImpl(), __local_var_16_21, bufs_7_4)
														}), gopurs_runtime.Value{})
														_ = __local_var_17_22
														__local_var_18_23 := gopurs_runtime.Apply(Call_Control_Applicative_liftA1(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Effect_applicativeEffect()), Get_Node_Buffer_Immutable_size(), gopurs_runtime.Apply(Get_Node_Buffer_unsafeFreeze(), __local_var_16_21)), gopurs_runtime.Value{})
														_ = __local_var_18_23
														__local_var_19_24 := gopurs_runtime.Apply3(Get_Effect_Ref_modifyImpl(), gopurs_runtime.Func(func(s_19 gopurs_runtime.Value) gopurs_runtime.Value {
															// TAST (Let): s_prime__20_25 shape=Other bindingType=Int
															s_prime__20_25 := gopurs_runtime.IntAdd(s_19.IntVal, __local_var_18_23.IntVal)
															_ = s_prime__20_25
															return gopurs_runtime.RecordDict2("state", "value", gopurs_runtime.Int(s_prime__20_25), gopurs_runtime.Int(s_prime__20_25))
														}), __local_var_6_3, gopurs_runtime.Value{})
														_ = __local_var_19_24
														var __t26 gopurs_runtime.Value
														{
															if (__local_var_19_24.IntVal) >= (n_3.IntVal) {
																__t26 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
																	return gopurs_runtime.Bool(true)
																})
																goto end_branch_26
															} else {

															}
														}
														{
															__t26 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
																return gopurs_runtime.Bool(false)
															})
														}
													end_branch_26:
														return gopurs_runtime.Apply(__t26, gopurs_runtime.Value{})
													})
													goto end_branch_27
												} else {

												}
											}
											{
												__t27 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
											}
										end_branch_27:
											return gopurs_runtime.Apply(__t27, gopurs_runtime.Value{})
										}))
										_ = __local_var_13_17
										__local_var_14_28 := gopurs_runtime.Apply(__local_var_13_17, gopurs_runtime.Value{})
										_ = __local_var_14_28
										__local_var_15_29 := gopurs_runtime.Apply2(Get_Effect_Ref_read(), __local_var_6_3, gopurs_runtime.Value{})
										_ = __local_var_15_29
										var __t35 gopurs_runtime.Value
										{
											if (__local_var_15_29.IntVal) >= (n_3.IntVal) {
												__t35 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
													__local_var_16_30 := gopurs_runtime.Apply(__local_var_9_6, gopurs_runtime.Value{})
													_ = __local_var_16_30
													__local_var_17_31 := gopurs_runtime.Apply(__local_var_10_8, gopurs_runtime.Value{})
													_ = __local_var_17_31
													__local_var_18_32 := gopurs_runtime.Apply(__local_var_11_12, gopurs_runtime.Value{})
													_ = __local_var_18_32
													__local_var_19_33 := gopurs_runtime.Apply2(Get_Control_Monad_ST_Global_toEffect(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
														return gopurs_runtime.UncurriedApp(Get_Data_Array_ST_unsafeFreezeImpl(), bufs_7_4)
													}), gopurs_runtime.Value{})
													_ = __local_var_19_33
													__local_var_20_34 := gopurs_runtime.UncurriedApp(Get_Node_Stream_readableImpl(), r_2)
													_ = __local_var_20_34
													return gopurs_runtime.Apply2(complete_4, func() gopurs_runtime.Value {
														_v := struct {
															V0 gopurs_runtime.Value
															V1 struct {
																buffers   []gopurs_runtime.Value
																readagain bool
																/*gopurs:record-v1:2:0062007500660066006500720073,00720065006100640061006700610069006e*/
															}
															V2 bool
														}{gopurs_runtime.Value{}, struct {
															buffers   []gopurs_runtime.Value
															readagain bool
															/*gopurs:record-v1:2:0062007500660066006500720073,00720065006100640061006700610069006e*/
														}{*((*[]gopurs_runtime.Value)(__local_var_19_33.UnsafePtr)), (__local_var_20_34.IntVal) != (0)}, true}
														if _v.V2 {
															return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: func() gopurs_runtime.Value {
																orig := _v.V1
																_ = orig
																return gopurs_runtime.RecordDict2("buffers", "readagain", gopurs_runtime.Array(orig.buffers), gopurs_runtime.Bool(orig.readagain))
															}()})}
														}
														return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
													}(), gopurs_runtime.Value{})
												})
												goto end_branch_35
											} else {

											}
										}
										{
											__t35 = gopurs_runtime.Apply(continuation_12, Get_Data_Unit_unit())
										}
									end_branch_35:
										return gopurs_runtime.Apply(__t35, gopurs_runtime.Value{})
									})
								}
								tryToRead__808114339_12_16 = gopurs_runtime.Func(func(continuation_12_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
									return Call_local_Node_Stream_Aff_tryToRead__808114339_12_16(continuation_12_loop_val)
								})
								var waitToRead_13_36_2 gopurs_runtime.Value
								_ = waitToRead_13_36_2
								var waitToRead_13_36_2_cell *gopurs_runtime.Value
								_ = waitToRead_13_36_2_cell
								// FALLBACK TCO: isLoop=false len=1
								waitToRead_13_36_2 = gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
										__local_var_15_37 := gopurs_runtime.Apply(Call_Node_EventEmitter_once(gopurs_runtime.CoerceToStruct[Constructor_Node_EventEmitter_EventHandle[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]](Get_Node_Stream_readableH()), Call_local_Node_Stream_Aff_tryToRead__808114339_12_16((*waitToRead_13_36_2_cell)), r_2), gopurs_runtime.Value{})
										_ = __local_var_15_37
										return gopurs_runtime.Apply3(Get_Effect_Ref_write(), __local_var_15_37, __local_var_8_5, gopurs_runtime.Value{})
									})
								})
								waitToRead_13_36_2_cell = &waitToRead_13_36_2
								return gopurs_runtime.Apply3(Get_Effect_Exception_catchException(), gopurs_runtime.Func(func(err_14 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
										__local_var_15_38 := gopurs_runtime.Apply(__local_var_9_6, gopurs_runtime.Value{})
										_ = __local_var_15_38
										__local_var_16_39 := gopurs_runtime.Apply(__local_var_10_8, gopurs_runtime.Value{})
										_ = __local_var_16_39
										__local_var_17_40 := gopurs_runtime.Apply(__local_var_11_12, gopurs_runtime.Value{})
										_ = __local_var_17_40
										__local_var_18_41 := gopurs_runtime.Apply(Call_Control_Bind_join(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect()), gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_8_5)), gopurs_runtime.Value{})
										_ = __local_var_18_41
										__local_var_19_42 := gopurs_runtime.Apply2(complete_4, func() gopurs_runtime.Value {
											_v := struct {
												V0 gopurs_runtime.Value
												V1 struct {
													buffers   []gopurs_runtime.Value
													readagain bool
													/*gopurs:record-v1:2:0062007500660066006500720073,00720065006100640061006700610069006e*/
												}
												V2 bool
											}{err_14, struct {
												buffers   []gopurs_runtime.Value
												readagain bool
												/*gopurs:record-v1:2:0062007500660066006500720073,00720065006100640061006700610069006e*/
											}{}, false}
											if _v.V2 {
												return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: func() gopurs_runtime.Value {
													orig := _v.V1
													_ = orig
													return gopurs_runtime.RecordDict2("buffers", "readagain", gopurs_runtime.Array(orig.buffers), gopurs_runtime.Bool(orig.readagain))
												}()})}
											}
											return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
										}(), gopurs_runtime.Value{})
										_ = __local_var_19_42
										return gopurs_runtime.Apply(Get_Data_Function_const__4207166068(), gopurs_runtime.Apply(Get_Effect_Aff__pure(), Get_Data_Unit_unit()))
									})
								}), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
									// TAST (Let): __local_var_14_43 shape=App(Other) bindingType=Any
									__local_var_14_43 := Call_local_Node_Stream_Aff_tryToRead__808114339_12_16(waitToRead_13_36_2)
									_ = __local_var_14_43
									__local_var_15_44 := gopurs_runtime.Apply(__local_var_14_43, gopurs_runtime.Value{})
									_ = __local_var_15_44
									return gopurs_runtime.Apply(Get_Effect_Aff_effectCanceler(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
										__local_var_16_45 := gopurs_runtime.Apply(__local_var_9_6, gopurs_runtime.Value{})
										_ = __local_var_16_45
										__local_var_17_46 := gopurs_runtime.Apply(__local_var_10_8, gopurs_runtime.Value{})
										_ = __local_var_17_46
										__local_var_18_47 := gopurs_runtime.Apply(__local_var_11_12, gopurs_runtime.Value{})
										_ = __local_var_18_47
										return gopurs_runtime.Apply(Call_Control_Bind_join(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect()), gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_8_5)), gopurs_runtime.Value{})
									}))
								}), gopurs_runtime.Value{})
							})
							goto end_branch_48
						} else {

						}
					}
					{
						__t48 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
							__local_var_6_2 := gopurs_runtime.Apply2(complete_4, func() gopurs_runtime.Value {
								_v := struct {
									V0 gopurs_runtime.Value
									V1 struct {
										buffers   []gopurs_runtime.Value
										readagain bool
										/*gopurs:record-v1:2:0062007500660066006500720073,00720065006100640061006700610069006e*/
									}
									V2 bool
								}{gopurs_runtime.Value{}, struct {
									buffers   []gopurs_runtime.Value
									readagain bool
									/*gopurs:record-v1:2:0062007500660066006500720073,00720065006100640061006700610069006e*/
								}{*((*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{}).UnsafePtr)), false}, true}
								if _v.V2 {
									return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: func() gopurs_runtime.Value {
										orig := _v.V1
										_ = orig
										return gopurs_runtime.RecordDict2("buffers", "readagain", gopurs_runtime.Array(orig.buffers), gopurs_runtime.Bool(orig.readagain))
									}()})}
								}
								return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
							}(), gopurs_runtime.Value{})
							_ = __local_var_6_2
							return gopurs_runtime.Apply(Get_Data_Function_const__4207166068(), gopurs_runtime.Apply(Get_Effect_Aff__pure(), Get_Data_Unit_unit()))
						})
					}
				end_branch_48:
					return gopurs_runtime.Apply(__t48, gopurs_runtime.Value{})
				})
			}
		end_branch_50:
			return __t50
		}))
	})
}

func Call_Node_Stream_Aff_readAll(dictMonadAff_0_loop *Constructor_Effect_Aff_Class_MonadAff[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictMonadAff_0 *Constructor_Effect_Aff_Class_MonadAff[gopurs_runtime.Value] = dictMonadAff_0_loop
	_ = dictMonadAff_0
	// TAST (Let): liftAff_1_0 shape=App(Var) bindingType=(Func [(ADT ["Effect","Aff","Aff"] [(Array (ADT ["Node","Buffer","Buffer"] []))])] (TypeApp (TypeVar m$scope89) [(Array (ADT ["Node","Buffer","Buffer"] []))]))
	liftAff_1_0 := Call_Effect_Aff_Class_liftAff(dictMonadAff_0)
	_ = liftAff_1_0
	return gopurs_runtime.Func(func(r_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply3(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), liftAff_1_0, Get_Effect_Aff_makeAff(), gopurs_runtime.Func(func(complete_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
				__local_var_4_1 := gopurs_runtime.UncurriedApp(Get_Node_Stream_readableImpl(), r_2)
				_ = __local_var_4_1
				var __t40 gopurs_runtime.Value
				{
					if (__local_var_4_1.IntVal) != (0) {
						__t40 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
							bufs_5_3 := gopurs_runtime.Apply2(Get_Control_Monad_ST_Global_toEffect(), Get_Data_Array_ST_go__new(), gopurs_runtime.Value{})
							_ = bufs_5_3
							__local_var_6_4 := gopurs_runtime.Apply2(Get_Effect_Ref__new(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
								return Get_Data_Unit_unit()
							}), gopurs_runtime.Value{})
							_ = __local_var_6_4
							__local_var_7_5 := gopurs_runtime.Apply(Call_Node_EventEmitter_once(gopurs_runtime.CoerceToStruct[Constructor_Node_EventEmitter_EventHandle[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]](Get_Node_Stream_errorH()), gopurs_runtime.Func(func(err_7 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
									__local_var_8_6 := gopurs_runtime.Apply(Call_Control_Bind_join(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect()), gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_6_4)), gopurs_runtime.Value{})
									_ = __local_var_8_6
									return gopurs_runtime.Apply2(complete_3, func() gopurs_runtime.Value {
										_v := struct {
											V0 gopurs_runtime.Value
											V1 gopurs_runtime.Value
											V2 bool
										}{err_7, gopurs_runtime.Value{}, false}
										if _v.V2 {
											return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
										}
										return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
									}(), gopurs_runtime.Value{})
								})
							}), r_2), gopurs_runtime.Value{})
							_ = __local_var_7_5
							__local_var_8_7 := gopurs_runtime.Apply(Call_Node_EventEmitter_once(gopurs_runtime.CoerceToStruct[Constructor_Node_EventEmitter_EventHandle[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]](Get_Node_Stream_closeH()), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
								__local_var_8_8 := gopurs_runtime.Apply(__local_var_7_5, gopurs_runtime.Value{})
								_ = __local_var_8_8
								__local_var_9_9 := gopurs_runtime.Apply(Call_Control_Bind_join(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect()), gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_6_4)), gopurs_runtime.Value{})
								_ = __local_var_9_9
								__local_var_10_10 := gopurs_runtime.Apply2(Get_Control_Monad_ST_Global_toEffect(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.UncurriedApp(Get_Data_Array_ST_unsafeFreezeImpl(), bufs_5_3)
								}), gopurs_runtime.Value{})
								_ = __local_var_10_10
								return gopurs_runtime.Apply2(complete_3, func() gopurs_runtime.Value {
									_v := struct {
										V0 gopurs_runtime.Value
										V1 gopurs_runtime.Value
										V2 bool
									}{gopurs_runtime.Value{}, __local_var_10_10, true}
									if _v.V2 {
										return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
									}
									return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
								}(), gopurs_runtime.Value{})
							}), r_2), gopurs_runtime.Value{})
							_ = __local_var_8_7
							__local_var_9_11 := gopurs_runtime.Apply(Call_Node_EventEmitter_once(gopurs_runtime.CoerceToStruct[Constructor_Node_EventEmitter_EventHandle[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]](Get_Node_Stream_endH()), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
								__local_var_9_12 := gopurs_runtime.Apply(__local_var_7_5, gopurs_runtime.Value{})
								_ = __local_var_9_12
								__local_var_10_13 := gopurs_runtime.Apply(__local_var_8_7, gopurs_runtime.Value{})
								_ = __local_var_10_13
								__local_var_11_14 := gopurs_runtime.Apply2(Get_Control_Monad_ST_Global_toEffect(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.UncurriedApp(Get_Data_Array_ST_unsafeFreezeImpl(), bufs_5_3)
								}), gopurs_runtime.Value{})
								_ = __local_var_11_14
								return gopurs_runtime.Apply2(complete_3, func() gopurs_runtime.Value {
									_v := struct {
										V0 gopurs_runtime.Value
										V1 gopurs_runtime.Value
										V2 bool
									}{gopurs_runtime.Value{}, __local_var_11_14, true}
									if _v.V2 {
										return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
									}
									return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
								}(), gopurs_runtime.Value{})
							}), r_2), gopurs_runtime.Value{})
							_ = __local_var_9_11
							return gopurs_runtime.Apply3(Get_Effect_Exception_catchException(), gopurs_runtime.Func(func(err_10 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
									__local_var_11_15 := gopurs_runtime.Apply(__local_var_7_5, gopurs_runtime.Value{})
									_ = __local_var_11_15
									__local_var_12_16 := gopurs_runtime.Apply(__local_var_8_7, gopurs_runtime.Value{})
									_ = __local_var_12_16
									__local_var_13_17 := gopurs_runtime.Apply(__local_var_9_11, gopurs_runtime.Value{})
									_ = __local_var_13_17
									__local_var_14_18 := gopurs_runtime.Apply(Call_Control_Bind_join(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect()), gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_6_4)), gopurs_runtime.Value{})
									_ = __local_var_14_18
									__local_var_15_19 := gopurs_runtime.Apply2(complete_3, func() gopurs_runtime.Value {
										_v := struct {
											V0 gopurs_runtime.Value
											V1 gopurs_runtime.Value
											V2 bool
										}{err_10, gopurs_runtime.Value{}, false}
										if _v.V2 {
											return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
										}
										return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
									}(), gopurs_runtime.Value{})
									_ = __local_var_15_19
									return gopurs_runtime.Apply(Get_Data_Function_const__4207166068(), gopurs_runtime.Apply(Get_Effect_Aff__pure(), Get_Data_Unit_unit()))
								})
							}), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
								// TAST (Let): __local_var_10_20 shape=App(Var) bindingType=Any
								__local_var_10_20 := gopurs_runtime.Apply(Get_Effect_untilE(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
									// TAST (Let): __local_var_10_21 shape=App(Var) bindingType=Any
									__local_var_10_21 := Call_Node_Stream_read(r_2)
									_ = __local_var_10_21
									__local_var_11_22 := gopurs_runtime.Apply(__local_var_10_21, gopurs_runtime.Value{})
									_ = __local_var_11_22
									var __t25 gopurs_runtime.Value
									{
										if __local_var_11_22.Type == 9 && __local_var_11_22.IntVal == 930809136 && __local_var_11_22.UnsafePtr == nil {
											__t25 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
												return gopurs_runtime.Bool(true)
											})
											goto end_branch_25
										} else {

										}
									}
									{
										if __local_var_11_22.Type == 9 && __local_var_11_22.IntVal == 930809136 && __local_var_11_22.UnsafePtr != nil {
											__t25 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
												// TAST (Let): __local_var_12_23 shape=Other bindingType=Any
												__local_var_12_23 := (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_11_22.UnsafePtr).V0
												_ = __local_var_12_23
												__local_var_13_24 := gopurs_runtime.Apply2(Call_Data_Functor_void(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect())), gopurs_runtime.Apply(Get_Control_Monad_ST_Global_toEffect(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
													return gopurs_runtime.UncurriedApp2(Get_Data_Array_ST_pushImpl(), __local_var_12_23, bufs_5_3)
												})), gopurs_runtime.Value{})
												_ = __local_var_13_24
												return gopurs_runtime.Bool(false)
											})
											goto end_branch_25
										} else {

										}
									}
									{
										__t25 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
									}
								end_branch_25:
									return gopurs_runtime.Apply(__t25, gopurs_runtime.Value{})
								}))
								_ = __local_var_10_20
								__local_var_11_26 := gopurs_runtime.Apply(__local_var_10_20, gopurs_runtime.Value{})
								_ = __local_var_11_26
								var waitToRead_12_27_3 gopurs_runtime.Value
								_ = waitToRead_12_27_3
								var waitToRead_12_27_3_cell *gopurs_runtime.Value
								_ = waitToRead_12_27_3_cell
								// FALLBACK TCO: isLoop=false len=1
								waitToRead_12_27_3 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
									__local_var_13_28 := gopurs_runtime.Apply(Call_Node_EventEmitter_once(gopurs_runtime.CoerceToStruct[Constructor_Node_EventEmitter_EventHandle[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]](Get_Node_Stream_readableH()), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
										// TAST (Let): __local_var_13_29 shape=App(Var) bindingType=Any
										__local_var_13_29 := gopurs_runtime.Apply(Get_Effect_untilE(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
											// TAST (Let): __local_var_13_30 shape=App(Var) bindingType=Any
											__local_var_13_30 := Call_Node_Stream_read(r_2)
											_ = __local_var_13_30
											__local_var_14_31 := gopurs_runtime.Apply(__local_var_13_30, gopurs_runtime.Value{})
											_ = __local_var_14_31
											var __t34 gopurs_runtime.Value
											{
												if __local_var_14_31.Type == 9 && __local_var_14_31.IntVal == 930809136 && __local_var_14_31.UnsafePtr == nil {
													__t34 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
														return gopurs_runtime.Bool(true)
													})
													goto end_branch_34
												} else {

												}
											}
											{
												if __local_var_14_31.Type == 9 && __local_var_14_31.IntVal == 930809136 && __local_var_14_31.UnsafePtr != nil {
													__t34 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
														// TAST (Let): __local_var_15_32 shape=Other bindingType=Any
														__local_var_15_32 := (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_14_31.UnsafePtr).V0
														_ = __local_var_15_32
														__local_var_16_33 := gopurs_runtime.Apply2(Get_Control_Monad_ST_Global_toEffect(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
															return gopurs_runtime.UncurriedApp2(Get_Data_Array_ST_pushImpl(), __local_var_15_32, bufs_5_3)
														}), gopurs_runtime.Value{})
														_ = __local_var_16_33
														return gopurs_runtime.Bool(false)
													})
													goto end_branch_34
												} else {

												}
											}
											{
												__t34 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
											}
										end_branch_34:
											return gopurs_runtime.Apply(__t34, gopurs_runtime.Value{})
										}))
										_ = __local_var_13_29
										__local_var_14_35 := gopurs_runtime.Apply(__local_var_13_29, gopurs_runtime.Value{})
										_ = __local_var_14_35
										return gopurs_runtime.Apply((*waitToRead_12_27_3_cell), gopurs_runtime.Value{})
									}), r_2), gopurs_runtime.Value{})
									_ = __local_var_13_28
									return gopurs_runtime.Apply3(Get_Effect_Ref_write(), __local_var_13_28, __local_var_6_4, gopurs_runtime.Value{})
								})
								waitToRead_12_27_3_cell = &waitToRead_12_27_3
								__local_var_13_36 := gopurs_runtime.Apply(waitToRead_12_27_3, gopurs_runtime.Value{})
								_ = __local_var_13_36
								return gopurs_runtime.Apply(Get_Effect_Aff_effectCanceler(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
									__local_var_14_37 := gopurs_runtime.Apply(__local_var_7_5, gopurs_runtime.Value{})
									_ = __local_var_14_37
									__local_var_15_38 := gopurs_runtime.Apply(__local_var_8_7, gopurs_runtime.Value{})
									_ = __local_var_15_38
									__local_var_16_39 := gopurs_runtime.Apply(__local_var_9_11, gopurs_runtime.Value{})
									_ = __local_var_16_39
									return gopurs_runtime.Apply(Call_Control_Bind_join(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect()), gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_6_4)), gopurs_runtime.Value{})
								}))
							}), gopurs_runtime.Value{})
						})
						goto end_branch_40
					} else {

					}
				}
				{
					__t40 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
						__local_var_5_2 := gopurs_runtime.Apply2(complete_3, func() gopurs_runtime.Value {
							_v := struct {
								V0 gopurs_runtime.Value
								V1 gopurs_runtime.Value
								V2 bool
							}{gopurs_runtime.Value{}, gopurs_runtime.Array([]gopurs_runtime.Value{}), true}
							if _v.V2 {
								return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
							}
							return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
						}(), gopurs_runtime.Value{})
						_ = __local_var_5_2
						return gopurs_runtime.Apply(Get_Data_Function_const__4207166068(), gopurs_runtime.Apply(Get_Effect_Aff__pure(), Get_Data_Unit_unit()))
					})
				}
			end_branch_40:
				return gopurs_runtime.Apply(__t40, gopurs_runtime.Value{})
			})
		}))
	})
}

func Call_Node_Stream_Aff_fromStringUTF8(dictMonadEffect_0_loop *Constructor_Effect_Class_MonadEffect[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictMonadEffect_0 *Constructor_Effect_Class_MonadEffect[gopurs_runtime.Value] = dictMonadEffect_0_loop
	_ = dictMonadEffect_0
	// TAST (Let): liftEffect_1_0 shape=App(Var) bindingType=(Func [(ADT ["Effect","Effect"] [(Array (ADT ["Node","Buffer","Buffer"] []))])] (TypeApp (TypeVar m$scope109) [(Array (ADT ["Node","Buffer","Buffer"] []))]))
	liftEffect_1_0 := Call_Effect_Class_liftEffect(dictMonadEffect_0)
	_ = liftEffect_1_0
	return gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(liftEffect_1_0, Call_Control_Applicative_liftA1(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Effect_applicativeEffect()), gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Control_Applicative_applicativeArray()).V1, Call_Node_Buffer_usingToImmutable(gopurs_runtime.Apply(Get_Node_Buffer_Immutable_fromString(), gopurs_runtime.NormalizeStringValue(s_2)), gopurs_runtime.Value{Type: 9, IntVal: int64(346813718), UnsafePtr: nil})))
	})
}

func Call_Node_Stream_Aff_end(dictMonadAff_0_loop *Constructor_Effect_Aff_Class_MonadAff[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictMonadAff_0 *Constructor_Effect_Aff_Class_MonadAff[gopurs_runtime.Value] = dictMonadAff_0_loop
	_ = dictMonadAff_0
	// TAST (Let): liftAff_1_0 shape=App(Var) bindingType=(Func [(ADT ["Effect","Aff","Aff"] [Unit])] (TypeApp (TypeVar m$scope112) [Unit]))
	liftAff_1_0 := Call_Effect_Aff_Class_liftAff(dictMonadAff_0)
	_ = liftAff_1_0
	return gopurs_runtime.Func(func(w_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply3(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), liftAff_1_0, Get_Effect_Aff_makeAff(), gopurs_runtime.Func(func(complete_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
				__local_var_4_1 := gopurs_runtime.UncurriedApp2(Get_Node_Stream_endCbImpl(), w_2, gopurs_runtime.Func(func(err_4 gopurs_runtime.Value) gopurs_runtime.Value {
					// TAST (Let): __local_var_5_2 shape=UncurriedApp(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(ADT ["Effect","Exception","Error"] [])])
					__local_var_5_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp3(Get_Data_Nullable_nullable(), err_4, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))}, Get_Data_Maybe_Just()))
					_ = __local_var_5_2
					var __t3 gopurs_runtime.Value
					{
						if __local_var_5_2 == nil {
							__t3 = gopurs_runtime.Apply(complete_3, func() gopurs_runtime.Value {
								_v := struct {
									V0 gopurs_runtime.Value
									V1 gopurs_runtime.Value
									V2 bool
								}{gopurs_runtime.Value{}, Get_Data_Unit_unit(), true}
								if _v.V2 {
									return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
								}
								return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
							}())
							goto end_branch_3
						} else {

						}
					}
					{
						if __local_var_5_2 != nil {
							__t3 = gopurs_runtime.Apply(complete_3, func() gopurs_runtime.Value {
								_v := struct {
									V0 gopurs_runtime.Value
									V1 gopurs_runtime.Value
									V2 bool
								}{(__local_var_5_2).V0, gopurs_runtime.Value{}, false}
								if _v.V2 {
									return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
								}
								return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
							}())
							goto end_branch_3
						} else {

						}
					}
					{
						__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
					}
				end_branch_3:
					return gopurs_runtime.Apply(__t3, gopurs_runtime.Value{})
				}))
				_ = __local_var_4_1
				return gopurs_runtime.Apply(Get_Data_Function_const__4207166068(), gopurs_runtime.Apply(Get_Effect_Aff__pure(), Get_Data_Unit_unit()))
			})
		}))
	})
}
