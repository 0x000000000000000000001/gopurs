var Fnil = []any{}

func Fcons(hd any) func([]any) []any {
	return func(tl []any) []any {
		newArr := make([]any, 0, len(tl)+1)
		newArr = append(newArr, hd)
		newArr = append(newArr, tl...)
		return newArr
	}
}

func FappendImpl(left []any) func([]any) []any {
	return func(right []any) []any {
		newArr := make([]any, 0, len(left)+len(right))
		newArr = append(newArr, left...)
		newArr = append(newArr, right...)
		return newArr
	}
}

func FflattenImpl(v [][]any) []any {
	acc := make([]any, 0)
	for _, inner := range v {
		acc = append(acc, inner...)
	}
	return acc
}

func FtoArray(vect any) any {
	return vect
}
