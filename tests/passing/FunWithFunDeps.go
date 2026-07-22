package Main
import "gopurs/output/gopurs_runtime"
var Fnil = gopurs_runtime.Array([]gopurs_runtime.Value{})
var Fcons = gopurs_runtime.Func(func(hd gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(tl gopurs_runtime.Value) gopurs_runtime.Value {
		tlArr := tl.PtrVal.([]gopurs_runtime.Value)
		newArr := make([]gopurs_runtime.Value, 0, len(tlArr)+1)
		newArr = append(newArr, hd)
		newArr = append(newArr, tlArr...)
		return gopurs_runtime.Array(newArr)
	})
})
var FappendImpl = gopurs_runtime.Func(func(left gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(right gopurs_runtime.Value) gopurs_runtime.Value {
		leftArr := left.PtrVal.([]gopurs_runtime.Value)
		rightArr := right.PtrVal.([]gopurs_runtime.Value)
		newArr := make([]gopurs_runtime.Value, 0, len(leftArr)+len(rightArr))
		newArr = append(newArr, leftArr...)
		newArr = append(newArr, rightArr...)
		return gopurs_runtime.Array(newArr)
	})
})
var FflattenImpl = gopurs_runtime.Func(func(v gopurs_runtime.Value) gopurs_runtime.Value {
	vArr := v.PtrVal.([]gopurs_runtime.Value)
	acc := make([]gopurs_runtime.Value, 0)
	for _, inner := range vArr {
		acc = append(acc, inner.PtrVal.([]gopurs_runtime.Value)...)
	}
	return gopurs_runtime.Array(acc)
})
var FtoArray = gopurs_runtime.Func(func(vect gopurs_runtime.Value) gopurs_runtime.Value {
	return vect
})
