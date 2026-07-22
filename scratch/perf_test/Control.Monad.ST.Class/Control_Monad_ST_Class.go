package Control_Monad_ST_Class

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Category "gopurs/output/Control.Category"
	pkg_Control_Monad_ST_Internal "gopurs/output/Control.Monad.ST.Internal"
	pkg_Effect "gopurs/output/Effect"
	pkg_Unsafe_Coerce "gopurs/output/Unsafe.Coerce"
)

var monadSTST gopurs_runtime.Value
var once_monadSTST sync.Once
func Get_monadSTST() gopurs_runtime.Value {
	once_monadSTST.Do(func() {
		monadSTST = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"liftST": pkg_Control_Category.Get_categoryFn().PtrVal.(map[string]gopurs_runtime.Value)["identity"], "Monad0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Monad_ST_Internal.Get_monadST()
})})
	})
	return monadSTST
}

var monadSTEffect gopurs_runtime.Value
var once_monadSTEffect sync.Once
func Get_monadSTEffect() gopurs_runtime.Value {
	once_monadSTEffect.Do(func() {
		monadSTEffect = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"liftST": pkg_Unsafe_Coerce.Get_unsafeCoerce(), "Monad0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Effect.Get_monadEffect()
})})
	})
	return monadSTEffect
}

var liftST gopurs_runtime.Value
var once_liftST sync.Once
func Get_liftST() gopurs_runtime.Value {
	once_liftST.Do(func() {
		liftST = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["liftST"]
})
	})
	return liftST
}


