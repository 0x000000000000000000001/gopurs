package Data_EuclideanRing

import (
	"gopurs/output/gopurs_runtime"
	"math"
)

var IntDegree = gopurs_runtime.Func(func(x gopurs_runtime.Value) gopurs_runtime.Value {
	xv := x.IntVal
	if xv < 0 {
		xv = -xv
	}
	if xv > 2147483647 {
		xv = 2147483647
	}
	return gopurs_runtime.Int(int(xv))
})

var IntDiv = gopurs_runtime.Func(func(x gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(y gopurs_runtime.Value) gopurs_runtime.Value {
		if y.IntVal == 0 {
			return gopurs_runtime.Int(0)
		}
		if y.IntVal > 0 {
			return gopurs_runtime.Int(int(math.Floor(float64(x.IntVal) / float64(y.IntVal))))
		}
		return gopurs_runtime.Int(int(-math.Floor(float64(x.IntVal) / float64(-y.IntVal))))
	})
})

var IntMod = gopurs_runtime.Func(func(x gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(y gopurs_runtime.Value) gopurs_runtime.Value {
		if y.IntVal == 0 {
			return gopurs_runtime.Int(0)
		}
		yy := y.IntVal
		if yy < 0 {
			yy = -yy
		}
		return gopurs_runtime.Int(int(((x.IntVal % yy) + yy) % yy))
	})
})

var NumDiv = gopurs_runtime.Func(func(n1 gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(n2 gopurs_runtime.Value) gopurs_runtime.Value {
		f1 := math.Float64frombits(uint64(n1.IntVal))
		f2 := math.Float64frombits(uint64(n2.IntVal))
		return gopurs_runtime.Float(f1 / f2)
	})
})
