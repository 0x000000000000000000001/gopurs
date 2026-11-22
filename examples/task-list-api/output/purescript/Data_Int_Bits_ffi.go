package purescript

import "gopurs/output/gopurs_runtime"

func Data_Int_Bits_And(n1 int, n2 int) int {
	return int(gopurs_runtime.IntBitAnd(int64(n1), int64(n2)))
}
func Data_Int_Bits_Or(n1 int, n2 int) int { return int(gopurs_runtime.IntBitOr(int64(n1), int64(n2))) }
func Data_Int_Bits_Xor(n1 int, n2 int) int {
	return int(gopurs_runtime.IntBitXor(int64(n1), int64(n2)))
}
func Data_Int_Bits_Shl(n1 int, n2 int) int  { return int(gopurs_runtime.IntShl(int64(n1), int64(n2))) }
func Data_Int_Bits_Shr(n1 int, n2 int) int  { return int(gopurs_runtime.IntShr(int64(n1), int64(n2))) }
func Data_Int_Bits_Zshr(n1 int, n2 int) int { return int(gopurs_runtime.IntZshr(int64(n1), int64(n2))) }
func Data_Int_Bits_Complement(n int) int    { return int(gopurs_runtime.IntBitNot(int64(n))) }

// --- Auto-generated FFI wrappers ---
var _Gopurs_Data_Int_Bits_And = // TAST: (Func [Int, Int] Int)
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int](arg0)
	go_arg1 := gopurs_runtime.Unbox[int](arg1)
	go_res := Data_Int_Bits_And(go_arg0, go_arg1)
	return gopurs_runtime.Int(int64(go_res))
})
var _Gopurs_Data_Int_Bits_Complement = // TAST: (Func [Int] Int)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int](arg0)
	go_res := Data_Int_Bits_Complement(go_arg0)
	return gopurs_runtime.Int(int64(go_res))
})
var _Gopurs_Data_Int_Bits_Or = // TAST: (Func [Int, Int] Int)
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int](arg0)
	go_arg1 := gopurs_runtime.Unbox[int](arg1)
	go_res := Data_Int_Bits_Or(go_arg0, go_arg1)
	return gopurs_runtime.Int(int64(go_res))
})
var _Gopurs_Data_Int_Bits_Shl = // TAST: (Func [Int, Int] Int)
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int](arg0)
	go_arg1 := gopurs_runtime.Unbox[int](arg1)
	go_res := Data_Int_Bits_Shl(go_arg0, go_arg1)
	return gopurs_runtime.Int(int64(go_res))
})
var _Gopurs_Data_Int_Bits_Shr = // TAST: (Func [Int, Int] Int)
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int](arg0)
	go_arg1 := gopurs_runtime.Unbox[int](arg1)
	go_res := Data_Int_Bits_Shr(go_arg0, go_arg1)
	return gopurs_runtime.Int(int64(go_res))
})
var _Gopurs_Data_Int_Bits_Xor = // TAST: (Func [Int, Int] Int)
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int](arg0)
	go_arg1 := gopurs_runtime.Unbox[int](arg1)
	go_res := Data_Int_Bits_Xor(go_arg0, go_arg1)
	return gopurs_runtime.Int(int64(go_res))
})
var _Gopurs_Data_Int_Bits_Zshr = // TAST: (Func [Int, Int] Int)
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int](arg0)
	go_arg1 := gopurs_runtime.Unbox[int](arg1)
	go_res := Data_Int_Bits_Zshr(go_arg0, go_arg1)
	return gopurs_runtime.Int(int64(go_res))
})
