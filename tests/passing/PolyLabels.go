func UnsafeGet(s string) func(map[string]any) any {
	return func(o map[string]any) any {
		return o[s]
	}
}

func UnsafeSet(s string, a any, o map[string]any) map[string]any {
	newMap := make(map[string]any)
	for k, v := range o {
		newMap[k] = v
	}
	newMap[s] = a
	return newMap
}
