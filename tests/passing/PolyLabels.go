func UnsafeGet(s string) func(map[string]any) any {
	return func(o map[string]any) any {
		return o[s]
	}
}

func UnsafeSet(s string) func(any) func(map[string]any) map[string]any {
	return func(a any) func(map[string]any) map[string]any {
		return func(o map[string]any) map[string]any {
			newMap := make(map[string]any)
			for k, v := range o {
				newMap[k] = v
			}
			newMap[s] = a
			return newMap
		}
	}
}
