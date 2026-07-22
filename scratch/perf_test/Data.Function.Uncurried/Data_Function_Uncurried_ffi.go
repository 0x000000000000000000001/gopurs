package Data_Function_Uncurried

import "gopurs/output/gopurs_runtime"

func mkRunFn(arity int) gopurs_runtime.Value {
	if arity == 1 {
		return gopurs_runtime.Func(func(fn gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(a gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply(fn, a)
			})
		})
	}
	if arity == 2 {
		return gopurs_runtime.Func(func(fn gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(a gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(b gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply(gopurs_runtime.Apply(fn, a), b)
				})
			})
		})
	}
	if arity == 3 {
		return gopurs_runtime.Func(func(fn gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(a gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(b gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(c gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(fn, a), b), c)
					})
				})
			})
		})
	}
	if arity == 4 {
		return gopurs_runtime.Func(func(fn gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(a gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(b gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(c gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(d gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(fn, a), b), c), d)
						})
					})
				})
			})
		})
	}
	if arity == 5 {
		return gopurs_runtime.Func(func(fn gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(a gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(b gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(c gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(d gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Func(func(e gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(fn, a), b), c), d), e)
							})
						})
					})
				})
			})
		})
	}
	return gopurs_runtime.Value{} // fallback
}

var RunFn2 = mkRunFn(2)
var RunFn3 = mkRunFn(3)
var RunFn4 = mkRunFn(4)
var RunFn5 = mkRunFn(5)
var RunFn6 = mkRunFn(6)
var RunFn7 = mkRunFn(7)
var RunFn8 = mkRunFn(8)
var RunFn9 = mkRunFn(9)
var RunFn10 = mkRunFn(10)

func mkMkFn(arity int) gopurs_runtime.Value {
    // mkFn is essentially identity because all functions in gopurs are already curried
	return gopurs_runtime.Func(func(fn gopurs_runtime.Value) gopurs_runtime.Value {
		return fn
	})
}

var MkFn2 = mkMkFn(2)
var MkFn3 = mkMkFn(3)
var MkFn4 = mkMkFn(4)
var MkFn5 = mkMkFn(5)
var MkFn6 = mkMkFn(6)
var MkFn7 = mkMkFn(7)
var MkFn8 = mkMkFn(8)
var MkFn9 = mkMkFn(9)
var MkFn10 = mkMkFn(10)
var MkFn0 = gopurs_runtime.Func(func(fn gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(fn, gopurs_runtime.Value{})
	})
})
var RunFn0 = gopurs_runtime.Func(func(fn gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Apply(fn, gopurs_runtime.Value{})
})
