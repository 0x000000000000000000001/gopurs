package Data_Ord

import (
	"gopurs/output/gopurs_runtime"
	"math"
)

var OrdBooleanImpl = gopurs_runtime.Func(func(lt gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(eq gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(gt gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(x gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(y gopurs_runtime.Value) gopurs_runtime.Value {
					if x.IntVal < y.IntVal {
						return lt
					} else if x.IntVal == y.IntVal {
						return eq
					} else {
						return gt
					}
				})
			})
		})
	})
})

var OrdIntImpl = OrdBooleanImpl

var OrdCharImpl = gopurs_runtime.Func(func(lt gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(eq gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(gt gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(x gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(y gopurs_runtime.Value) gopurs_runtime.Value {
					if x.StrVal < y.StrVal {
						return lt
					} else if x.StrVal == y.StrVal {
						return eq
					} else {
						return gt
					}
				})
			})
		})
	})
})

var OrdStringImpl = OrdCharImpl

var OrdNumberImpl = gopurs_runtime.Func(func(lt gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(eq gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(gt gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(x gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(y gopurs_runtime.Value) gopurs_runtime.Value {
					xf := math.Float64frombits(uint64(x.IntVal))
					yf := math.Float64frombits(uint64(y.IntVal))
					if xf < yf {
						return lt
					} else if xf == yf {
						return eq
					} else {
						return gt
					}
				})
			})
		})
	})
})

var OrdArrayImpl = gopurs_runtime.Func(func(f gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(xs gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(ys gopurs_runtime.Value) gopurs_runtime.Value {
			xsArr := xs.PtrVal.([]gopurs_runtime.Value)
			ysArr := ys.PtrVal.([]gopurs_runtime.Value)
			i := 0
			xlen := len(xsArr)
			ylen := len(ysArr)
			for i < xlen && i < ylen {
				o := gopurs_runtime.Apply(gopurs_runtime.Apply(f, xsArr[i]), ysArr[i]).IntVal
				if o != 0 {
					return gopurs_runtime.Int(int(o))
				}
				i++
			}
			if xlen == ylen {
				return gopurs_runtime.Int(0)
			} else if xlen > ylen {
				return gopurs_runtime.Int(-1)
			} else {
				return gopurs_runtime.Int(1)
			}
		})
	})
})
