package Data_Traversable

import "gopurs/output/gopurs_runtime"

var TraverseArrayImpl = gopurs_runtime.Func(func(apply gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(mapFn gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(pure gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(f gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(arrayVal gopurs_runtime.Value) gopurs_runtime.Value {
					
					array1 := gopurs_runtime.Func(func(a gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Array([]gopurs_runtime.Value{a})
					})
					
					array2 := gopurs_runtime.Func(func(a gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(b gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Array([]gopurs_runtime.Value{a, b})
						})
					})
					
					array3 := gopurs_runtime.Func(func(a gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(b gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Func(func(c gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Array([]gopurs_runtime.Value{a, b, c})
							})
						})
					})
					
					concat2 := gopurs_runtime.Func(func(xsVal gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(ysVal gopurs_runtime.Value) gopurs_runtime.Value {
							xs := xsVal.PtrVal.([]gopurs_runtime.Value)
							ys := ysVal.PtrVal.([]gopurs_runtime.Value)
							res := make([]gopurs_runtime.Value, 0, len(xs)+len(ys))
							res = append(res, xs...)
							res = append(res, ys...)
							return gopurs_runtime.Array(res)
						})
					})
					
					arr := arrayVal.PtrVal.([]gopurs_runtime.Value)
					
					var goFn func(int, int) gopurs_runtime.Value
					goFn = func(bot, top int) gopurs_runtime.Value {
						switch top - bot {
						case 0:
							return gopurs_runtime.Apply(pure, gopurs_runtime.Array([]gopurs_runtime.Value{}))
						case 1:
							return gopurs_runtime.Apply(gopurs_runtime.Apply(mapFn, array1), gopurs_runtime.Apply(f, arr[bot]))
						case 2:
							return gopurs_runtime.Apply(gopurs_runtime.Apply(apply, gopurs_runtime.Apply(gopurs_runtime.Apply(mapFn, array2), gopurs_runtime.Apply(f, arr[bot]))), gopurs_runtime.Apply(f, arr[bot+1]))
						case 3:
							return gopurs_runtime.Apply(gopurs_runtime.Apply(apply, gopurs_runtime.Apply(gopurs_runtime.Apply(apply, gopurs_runtime.Apply(gopurs_runtime.Apply(mapFn, array3), gopurs_runtime.Apply(f, arr[bot]))), gopurs_runtime.Apply(f, arr[bot+1]))), gopurs_runtime.Apply(f, arr[bot+2]))
						default:
							pivot := bot + ((top-bot)/4)*2
							return gopurs_runtime.Apply(gopurs_runtime.Apply(apply, gopurs_runtime.Apply(gopurs_runtime.Apply(mapFn, concat2), goFn(bot, pivot))), goFn(pivot, top))
						}
					}
					
					return goFn(0, len(arr))
				})
			})
		})
	})
})
