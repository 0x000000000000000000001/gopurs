func ShowImpl(showFn func(any) any) func(any) any {
	return func(val any) any {
		return showFn(val)
	}
}
