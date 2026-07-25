func RuntimeImportImpl(nothing any) func(func(any) any) func(string) func(func(any) func(any) any) func(any) any {
	return func(just func(any) any) func(string) func(func(any) func(any) any) func(any) any {
		return func(moduleName string) func(func(any) func(any) any) func(any) any {
			return func(body func(any) func(any) any) func(any) any {
				return func(_ any) any {
					errStr := "dynamic import not supported in Go"
					errVal := just(errStr)
					action := body(errVal)
					return action(nil)
				}
			}
		}
	}
}
