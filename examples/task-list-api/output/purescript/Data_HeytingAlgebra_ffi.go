package purescript

import "gopurs/output/gopurs_runtime"

func Data_HeytingAlgebra_BoolConj(b1 bool, b2 bool) bool {
	return b1 && b2
}
func Data_HeytingAlgebra_BoolDisj(b1 bool, b2 bool) bool {
	return b1 || b2
}
func Data_HeytingAlgebra_BoolNot(b bool) bool {
	return !b
}

// --- Auto-generated FFI wrappers ---
var _Gopurs_Data_HeytingAlgebra_BoolConj = // TAST: (Func [Boolean, Boolean] Boolean)
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[bool](arg0)
	go_arg1 := gopurs_runtime.Unbox[bool](arg1)
	go_res := Data_HeytingAlgebra_BoolConj(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_HeytingAlgebra_BoolDisj = // TAST: (Func [Boolean, Boolean] Boolean)
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[bool](arg0)
	go_arg1 := gopurs_runtime.Unbox[bool](arg1)
	go_res := Data_HeytingAlgebra_BoolDisj(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_HeytingAlgebra_BoolNot = // TAST: (Func [Boolean] Boolean)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[bool](arg0)
	go_res := Data_HeytingAlgebra_BoolNot(go_arg0)
	return gopurs_runtime.Box(go_res)
})
