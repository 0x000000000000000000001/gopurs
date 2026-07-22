package Data_Array_ST

import "gopurs/output/gopurs_runtime"
import gosort "sort"

var New_ = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
	arr := make([]gopurs_runtime.Value, 0)
	return gopurs_runtime.Value{PtrVal: &arr}
})

var PeekImpl = gopurs_runtime.Func(func(just gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(nothing gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(iVal gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(xsVal gopurs_runtime.Value) gopurs_runtime.Value {
				i := int(iVal.IntVal)
				xs := *(xsVal.PtrVal.(*[]gopurs_runtime.Value))
				if i >= 0 && i < len(xs) {
					return gopurs_runtime.Apply(just, xs[i])
				}
				return nothing
			})
		})
	})
})

var PokeImpl = gopurs_runtime.Func(func(iVal gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(a gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(xsVal gopurs_runtime.Value) gopurs_runtime.Value {
			i := int(iVal.IntVal)
			xsPtr := xsVal.PtrVal.(*[]gopurs_runtime.Value)
			xs := *xsPtr
			ret := i >= 0 && i < len(xs)
			if ret {
				xs[i] = a
			}
			if ret {
				return gopurs_runtime.Int(1) // Assuming boolean mapped to 1/0
			}
			return gopurs_runtime.Int(0)
		})
	})
})

var LengthImpl = gopurs_runtime.Func(func(xsVal gopurs_runtime.Value) gopurs_runtime.Value {
	xs := *(xsVal.PtrVal.(*[]gopurs_runtime.Value))
	return gopurs_runtime.Int(len(xs))
})

var PopImpl = gopurs_runtime.Func(func(just gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(nothing gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(xsVal gopurs_runtime.Value) gopurs_runtime.Value {
			xsPtr := xsVal.PtrVal.(*[]gopurs_runtime.Value)
			xs := *xsPtr
			if len(xs) > 0 {
				last := xs[len(xs)-1]
				*xsPtr = xs[:len(xs)-1]
				return gopurs_runtime.Apply(just, last)
			}
			return nothing
		})
	})
})

var PushAllImpl = gopurs_runtime.Func(func(asVal gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(xsVal gopurs_runtime.Value) gopurs_runtime.Value {
		as := asVal.PtrVal.([]gopurs_runtime.Value)
		xsPtr := xsVal.PtrVal.(*[]gopurs_runtime.Value)
		*xsPtr = append(*xsPtr, as...)
		return gopurs_runtime.Int(len(*xsPtr))
	})
})

var ShiftImpl = gopurs_runtime.Func(func(just gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(nothing gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(xsVal gopurs_runtime.Value) gopurs_runtime.Value {
			xsPtr := xsVal.PtrVal.(*[]gopurs_runtime.Value)
			xs := *xsPtr
			if len(xs) > 0 {
				first := xs[0]
				*xsPtr = xs[1:]
				return gopurs_runtime.Apply(just, first)
			}
			return nothing
		})
	})
})

var UnshiftAllImpl = gopurs_runtime.Func(func(asVal gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(xsVal gopurs_runtime.Value) gopurs_runtime.Value {
		as := asVal.PtrVal.([]gopurs_runtime.Value)
		xsPtr := xsVal.PtrVal.(*[]gopurs_runtime.Value)
		newArr := make([]gopurs_runtime.Value, 0, len(as)+len(*xsPtr))
		newArr = append(newArr, as...)
		newArr = append(newArr, *xsPtr...)
		*xsPtr = newArr
		return gopurs_runtime.Int(len(*xsPtr))
	})
})

var SpliceImpl = gopurs_runtime.Func(func(iVal gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(howManyVal gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(bsVal gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(xsVal gopurs_runtime.Value) gopurs_runtime.Value {
				i := int(iVal.IntVal)
				howMany := int(howManyVal.IntVal)
				bs := bsVal.PtrVal.([]gopurs_runtime.Value)
				xsPtr := xsVal.PtrVal.(*[]gopurs_runtime.Value)
				xs := *xsPtr
				
				if i < 0 {
					i = len(xs) + i
					if i < 0 {
						i = 0
					}
				} else if i > len(xs) {
					i = len(xs)
				}
				
				if howMany < 0 {
					howMany = 0
				} else if i+howMany > len(xs) {
					howMany = len(xs) - i
				}
				
				removed := make([]gopurs_runtime.Value, howMany)
				copy(removed, xs[i:i+howMany])
				
				newArr := make([]gopurs_runtime.Value, 0, len(xs)-howMany+len(bs))
				newArr = append(newArr, xs[:i]...)
				newArr = append(newArr, bs...)
				newArr = append(newArr, xs[i+howMany:]...)
				*xsPtr = newArr
				
				return gopurs_runtime.Array(removed)
			})
		})
	})
})

var UnsafeFreezeImpl = gopurs_runtime.Func(func(xsVal gopurs_runtime.Value) gopurs_runtime.Value {
	xs := *(xsVal.PtrVal.(*[]gopurs_runtime.Value))
	return gopurs_runtime.Array(xs)
})

var UnsafeThawImpl = gopurs_runtime.Func(func(xsVal gopurs_runtime.Value) gopurs_runtime.Value {
	xs := xsVal.PtrVal.([]gopurs_runtime.Value)
	return gopurs_runtime.Value{PtrVal: &xs}
})

var FreezeImpl = gopurs_runtime.Func(func(xsVal gopurs_runtime.Value) gopurs_runtime.Value {
	xs := *(xsVal.PtrVal.(*[]gopurs_runtime.Value))
	copyArr := make([]gopurs_runtime.Value, len(xs))
	copy(copyArr, xs)
	return gopurs_runtime.Array(copyArr)
})

var ThawImpl = gopurs_runtime.Func(func(xsVal gopurs_runtime.Value) gopurs_runtime.Value {
	xs := xsVal.PtrVal.([]gopurs_runtime.Value)
	copyArr := make([]gopurs_runtime.Value, len(xs))
	copy(copyArr, xs)
	return gopurs_runtime.Value{PtrVal: &copyArr}
})

var CloneImpl = gopurs_runtime.Func(func(xsVal gopurs_runtime.Value) gopurs_runtime.Value {
	xs := *(xsVal.PtrVal.(*[]gopurs_runtime.Value))
	copyArr := make([]gopurs_runtime.Value, len(xs))
	copy(copyArr, xs)
	return gopurs_runtime.Value{PtrVal: &copyArr}
})

var SortByImpl = gopurs_runtime.Func(func(compare gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(fromOrdering gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(xsVal gopurs_runtime.Value) gopurs_runtime.Value {
			xsPtr := xsVal.PtrVal.(*[]gopurs_runtime.Value)
			xs := *xsPtr
			
			gosort.SliceStable(xs, func(i, j int) bool {
				c := gopurs_runtime.Apply(fromOrdering, gopurs_runtime.Apply(gopurs_runtime.Apply(compare, xs[i]), xs[j]))
				return c.IntVal < 0
			})
			
			return xsVal
		})
	})
})

var ToAssocArrayImpl = gopurs_runtime.Func(func(xsVal gopurs_runtime.Value) gopurs_runtime.Value {
	xs := *(xsVal.PtrVal.(*[]gopurs_runtime.Value))
	as := make([]gopurs_runtime.Value, len(xs))
	for i, v := range xs {
		dict := map[string]gopurs_runtime.Value{
			"value": v,
			"index": gopurs_runtime.Int(i),
		}
		as[i] = gopurs_runtime.Value{PtrVal: dict} // Ensure it's a valid PtrVal record
	}
	return gopurs_runtime.Array(as)
})

var PushImpl = gopurs_runtime.Func(func(aVal gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(xsVal gopurs_runtime.Value) gopurs_runtime.Value {
		xsPtr := xsVal.PtrVal.(*[]gopurs_runtime.Value)
		*xsPtr = append(*xsPtr, aVal)
		return gopurs_runtime.Int(len(*xsPtr))
	})
})
