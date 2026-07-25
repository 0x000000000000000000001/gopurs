func MergeImpl(l map[string]any) func(map[string]any) map[string]any {
	return func(r map[string]any) map[string]any {
		newMap := make(map[string]any)
		for k, v := range r {
			newMap[k] = v
		}
		for k, v := range l {
			newMap[k] = v
		}
		return newMap
	}
}
