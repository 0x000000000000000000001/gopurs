package Data_Array

import "gopurs/output/gopurs_runtime"

var RangeImpl = gopurs_runtime.Func(func(startVal gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(endVal gopurs_runtime.Value) gopurs_runtime.Value {
		start := int(startVal.IntVal)
		end := int(endVal.IntVal)
		step := 1
		if start > end {
			step = -1
		}
		
		size := (end - start) * step + 1
		result := make([]gopurs_runtime.Value, size)
		
		i := start
		n := 0
		for i != end {
			result[n] = gopurs_runtime.Int(i)
			n++
			i += step
		}
		result[n] = gopurs_runtime.Int(i)
		return gopurs_runtime.Array(result)
	})
})

var ReplicateImpl = gopurs_runtime.Func(func(countVal gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(value gopurs_runtime.Value) gopurs_runtime.Value {
		count := int(countVal.IntVal)
		if count < 1 {
			return gopurs_runtime.Array(make([]gopurs_runtime.Value, 0))
		}
		result := make([]gopurs_runtime.Value, count)
		for i := 0; i < count; i++ {
			result[i] = value
		}
		return gopurs_runtime.Array(result)
	})
})

var Length = gopurs_runtime.Func(func(xsVal gopurs_runtime.Value) gopurs_runtime.Value {
	xs := xsVal.PtrVal.([]gopurs_runtime.Value)
	return gopurs_runtime.Int(len(xs))
})

var UnconsImpl = gopurs_runtime.Func(func(empty gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(next gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(xsVal gopurs_runtime.Value) gopurs_runtime.Value {
			xs := xsVal.PtrVal.([]gopurs_runtime.Value)
			if len(xs) == 0 {
				return gopurs_runtime.Apply(empty, gopurs_runtime.Value{})
			}
			head := xs[0]
			tail := make([]gopurs_runtime.Value, len(xs)-1)
			copy(tail, xs[1:])
			return gopurs_runtime.Apply(gopurs_runtime.Apply(next, head), gopurs_runtime.Array(tail))
		})
	})
})

var IndexImpl = gopurs_runtime.Func(func(just gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(nothing gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(xsVal gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(iVal gopurs_runtime.Value) gopurs_runtime.Value {
				xs := xsVal.PtrVal.([]gopurs_runtime.Value)
				i := int(iVal.IntVal)
				if i < 0 || i >= len(xs) {
					return nothing
				}
				return gopurs_runtime.Apply(just, xs[i])
			})
		})
	})
})

var X_updateAt = gopurs_runtime.Func(func(just gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(nothing gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(iVal gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(aVal gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(xsVal gopurs_runtime.Value) gopurs_runtime.Value {
					xs := xsVal.PtrVal.([]gopurs_runtime.Value)
					i := int(iVal.IntVal)
					if i < 0 || i >= len(xs) {
						return nothing
					}
					l1 := make([]gopurs_runtime.Value, len(xs))
					copy(l1, xs)
					l1[i] = aVal
					return gopurs_runtime.Apply(just, gopurs_runtime.Array(l1))
				})
			})
		})
	})
})

var X_insertAt = gopurs_runtime.Func(func(just gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(nothing gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(iVal gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(aVal gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(xsVal gopurs_runtime.Value) gopurs_runtime.Value {
					xs := xsVal.PtrVal.([]gopurs_runtime.Value)
					i := int(iVal.IntVal)
					if i < 0 || i > len(xs) {
						return nothing
					}
					l1 := make([]gopurs_runtime.Value, 0, len(xs)+1)
					l1 = append(l1, xs[:i]...)
					l1 = append(l1, aVal)
					l1 = append(l1, xs[i:]...)
					return gopurs_runtime.Apply(just, gopurs_runtime.Array(l1))
				})
			})
		})
	})
})

var X_deleteAt = gopurs_runtime.Func(func(just gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(nothing gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(iVal gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(xsVal gopurs_runtime.Value) gopurs_runtime.Value {
				xs := xsVal.PtrVal.([]gopurs_runtime.Value)
				i := int(iVal.IntVal)
				if i < 0 || i >= len(xs) {
					return nothing
				}
				l1 := make([]gopurs_runtime.Value, 0, len(xs)-1)
				l1 = append(l1, xs[:i]...)
				l1 = append(l1, xs[i+1:]...)
				return gopurs_runtime.Apply(just, gopurs_runtime.Array(l1))
			})
		})
	})
})

var Reverse = gopurs_runtime.Func(func(xsVal gopurs_runtime.Value) gopurs_runtime.Value {
	xs := xsVal.PtrVal.([]gopurs_runtime.Value)
	l := len(xs)
	l1 := make([]gopurs_runtime.Value, l)
	for i := 0; i < l; i++ {
		l1[i] = xs[l-1-i]
	}
	return gopurs_runtime.Array(l1)
})

var Concat = gopurs_runtime.Func(func(xssVal gopurs_runtime.Value) gopurs_runtime.Value {
	xss := xssVal.PtrVal.([]gopurs_runtime.Value)
	var result []gopurs_runtime.Value
	for _, xsVal := range xss {
		xs := xsVal.PtrVal.([]gopurs_runtime.Value)
		result = append(result, xs...)
	}
	return gopurs_runtime.Array(result)
})

var FilterImpl = gopurs_runtime.Func(func(f gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(xsVal gopurs_runtime.Value) gopurs_runtime.Value {
		xs := xsVal.PtrVal.([]gopurs_runtime.Value)
		var result []gopurs_runtime.Value
		for _, x := range xs {
			if gopurs_runtime.Apply(f, x).IntVal != 0 { // Bool mapping
				result = append(result, x)
			}
		}
		return gopurs_runtime.Array(result)
	})
})

var SliceImpl = gopurs_runtime.Func(func(sVal gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(eVal gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(lVal gopurs_runtime.Value) gopurs_runtime.Value {
			s := int(sVal.IntVal)
			e := int(eVal.IntVal)
			l := lVal.PtrVal.([]gopurs_runtime.Value)
			if s < 0 {
				s = len(l) + s
			}
			if e < 0 {
				e = len(l) + e
			}
			if s < 0 { s = 0 }
			if e > len(l) { e = len(l) }
			if s > e { s = e }
			
			res := make([]gopurs_runtime.Value, e-s)
			copy(res, l[s:e])
			return gopurs_runtime.Array(res)
		})
	})
})

var ZipWithImpl = gopurs_runtime.Func(func(f gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(xsVal gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(ysVal gopurs_runtime.Value) gopurs_runtime.Value {
			xs := xsVal.PtrVal.([]gopurs_runtime.Value)
			ys := ysVal.PtrVal.([]gopurs_runtime.Value)
			length := len(xs)
			if len(ys) < length {
				length = len(ys)
			}
			result := make([]gopurs_runtime.Value, length)
			for i := 0; i < length; i++ {
				result[i] = gopurs_runtime.Apply(gopurs_runtime.Apply(f, xs[i]), ys[i])
			}
			return gopurs_runtime.Array(result)
		})
	})
})

var UnsafeIndexImpl = gopurs_runtime.Func(func(xsVal gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(nVal gopurs_runtime.Value) gopurs_runtime.Value {
		xs := xsVal.PtrVal.([]gopurs_runtime.Value)
		n := int(nVal.IntVal)
		return xs[n]
	})
})

var SortByImpl = gopurs_runtime.Func(func(compare gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(fromOrdering gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(xsVal gopurs_runtime.Value) gopurs_runtime.Value {
			xs := xsVal.PtrVal.([]gopurs_runtime.Value)
			if len(xs) < 2 {
				return xsVal
			}
			out := make([]gopurs_runtime.Value, len(xs))
			copy(out, xs)
			// Sort implementation using merge sort or just simple sort
			// Wait, simple sort:
			for i := 0; i < len(out); i++ {
				for j := i + 1; j < len(out); j++ {
					c := gopurs_runtime.Apply(fromOrdering, gopurs_runtime.Apply(gopurs_runtime.Apply(compare, out[i]), out[j]))
					if c.IntVal > 0 { // GT
						out[i], out[j] = out[j], out[i]
					}
				}
			}
			return gopurs_runtime.Array(out)
		})
	})
})

var ScanrImpl = gopurs_runtime.Func(func(f gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(b gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(xsVal gopurs_runtime.Value) gopurs_runtime.Value {
			xs := xsVal.PtrVal.([]gopurs_runtime.Value)
			out := make([]gopurs_runtime.Value, len(xs))
			acc := b
			for i := len(xs) - 1; i >= 0; i-- {
				acc = gopurs_runtime.Apply(gopurs_runtime.Apply(f, xs[i]), acc)
				out[i] = acc
			}
			return gopurs_runtime.Array(out)
		})
	})
})

var ScanlImpl = gopurs_runtime.Func(func(f gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(b gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(xsVal gopurs_runtime.Value) gopurs_runtime.Value {
			xs := xsVal.PtrVal.([]gopurs_runtime.Value)
			out := make([]gopurs_runtime.Value, len(xs))
			acc := b
			for i := 0; i < len(xs); i++ {
				acc = gopurs_runtime.Apply(gopurs_runtime.Apply(f, acc), xs[i])
				out[i] = acc
			}
			return gopurs_runtime.Array(out)
		})
	})
})

var PartitionImpl = gopurs_runtime.Func(func(f gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(xsVal gopurs_runtime.Value) gopurs_runtime.Value {
		xs := xsVal.PtrVal.([]gopurs_runtime.Value)
		var yes []gopurs_runtime.Value
		var no []gopurs_runtime.Value
		for _, x := range xs {
			if gopurs_runtime.Apply(f, x).IntVal != 0 {
				yes = append(yes, x)
			} else {
				no = append(no, x)
			}
		}
		return gopurs_runtime.Record(map[string]gopurs_runtime.Value{
			"yes": gopurs_runtime.Array(yes),
			"no":  gopurs_runtime.Array(no),
		})
	})
})

var FromFoldableImpl = gopurs_runtime.Func(func(foldr gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(xsVal gopurs_runtime.Value) gopurs_runtime.Value {
		// A simple list-to-array since foldr passes elements.
		// wait, fromFoldableImpl uses a custom Cons.
		// Actually, we can just do what fromFoldableImpl does in JS:
		// foldr(\a acc -> Cons(a, acc))(empty)(xs)
		// Or since we just want array, maybe we can implement it without internal list?
		// fromFoldableImpl foldr xs = listToArray(foldr(curryCons)(emptyList)(xs))
		// We'll just define emptyList = [] and Cons = append
		// BUT foldr builds from right to left! 
		// Actually JS implementation uses a linked list for performance.
		
		var emptyList gopurs_runtime.Value // nil map or struct
		emptyList = gopurs_runtime.Value{PtrVal: nil}
		
		curryCons := gopurs_runtime.Func(func(head gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(tail gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{PtrVal: map[string]gopurs_runtime.Value{
					"head": head,
					"tail": tail,
				}}
			})
		})
		
		list := gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(foldr, curryCons), emptyList), xsVal)
		
		var result []gopurs_runtime.Value
		curr := list
		for curr.PtrVal != nil {
			m := curr.PtrVal.(map[string]gopurs_runtime.Value)
			result = append(result, m["head"])
			curr = m["tail"]
		}
		return gopurs_runtime.Array(result)
	})
})

var FindMapImpl = gopurs_runtime.Func(func(nothing gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(isJust gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(f gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(xsVal gopurs_runtime.Value) gopurs_runtime.Value {
				xs := xsVal.PtrVal.([]gopurs_runtime.Value)
				for _, x := range xs {
					res := gopurs_runtime.Apply(f, x)
					if gopurs_runtime.Apply(isJust, res).IntVal != 0 {
						return res
					}
				}
				return nothing
			})
		})
	})
})

var FindLastIndexImpl = gopurs_runtime.Func(func(just gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(nothing gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(f gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(xsVal gopurs_runtime.Value) gopurs_runtime.Value {
				xs := xsVal.PtrVal.([]gopurs_runtime.Value)
				for i := len(xs) - 1; i >= 0; i-- {
					if gopurs_runtime.Apply(f, xs[i]).IntVal != 0 {
						return gopurs_runtime.Apply(just, gopurs_runtime.Int(i))
					}
				}
				return nothing
			})
		})
	})
})

var FindIndexImpl = gopurs_runtime.Func(func(just gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(nothing gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(f gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(xsVal gopurs_runtime.Value) gopurs_runtime.Value {
				xs := xsVal.PtrVal.([]gopurs_runtime.Value)
				for i := 0; i < len(xs); i++ {
					if gopurs_runtime.Apply(f, xs[i]).IntVal != 0 {
						return gopurs_runtime.Apply(just, gopurs_runtime.Int(i))
					}
				}
				return nothing
			})
		})
	})
})

var AnyImpl = gopurs_runtime.Func(func(p gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(xsVal gopurs_runtime.Value) gopurs_runtime.Value {
		xs := xsVal.PtrVal.([]gopurs_runtime.Value)
		for _, x := range xs {
			if gopurs_runtime.Apply(p, x).IntVal != 0 {
				return gopurs_runtime.Int(1)
			}
		}
		return gopurs_runtime.Int(0)
	})
})

var AllImpl = gopurs_runtime.Func(func(p gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(xsVal gopurs_runtime.Value) gopurs_runtime.Value {
		xs := xsVal.PtrVal.([]gopurs_runtime.Value)
		for _, x := range xs {
			if gopurs_runtime.Apply(p, x).IntVal == 0 {
				return gopurs_runtime.Int(0)
			}
		}
		return gopurs_runtime.Int(1)
	})
})

var X_InsertAt = gopurs_runtime.Func(func(just gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(nothing gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(iVal gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(a gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(lVal gopurs_runtime.Value) gopurs_runtime.Value {
					i := int(iVal.IntVal)
					l := lVal.PtrVal.([]gopurs_runtime.Value)
					if i < 0 || i > len(l) {
						return nothing
					}
					l1 := make([]gopurs_runtime.Value, 0, len(l)+1)
					l1 = append(l1, l[:i]...)
					l1 = append(l1, a)
					l1 = append(l1, l[i:]...)
					return gopurs_runtime.Apply(just, gopurs_runtime.Array(l1))
				})
			})
		})
	})
})

var X_DeleteAt = gopurs_runtime.Func(func(just gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(nothing gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(iVal gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(lVal gopurs_runtime.Value) gopurs_runtime.Value {
				i := int(iVal.IntVal)
				l := lVal.PtrVal.([]gopurs_runtime.Value)
				if i < 0 || i >= len(l) {
					return nothing
				}
				l1 := make([]gopurs_runtime.Value, 0, len(l)-1)
				l1 = append(l1, l[:i]...)
				l1 = append(l1, l[i+1:]...)
				return gopurs_runtime.Apply(just, gopurs_runtime.Array(l1))
			})
		})
	})
})

var X_UpdateAt = gopurs_runtime.Func(func(just gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(nothing gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(iVal gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(a gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(lVal gopurs_runtime.Value) gopurs_runtime.Value {
					i := int(iVal.IntVal)
					l := lVal.PtrVal.([]gopurs_runtime.Value)
					if i < 0 || i >= len(l) {
						return nothing
					}
					l1 := make([]gopurs_runtime.Value, len(l))
					copy(l1, l)
					l1[i] = a
					return gopurs_runtime.Apply(just, gopurs_runtime.Array(l1))
				})
			})
		})
	})
})
