func Add3(a string) func(string) func(string) string {
	return func(b string) func(string) string {
		return func(c string) string {
			return a + b + c
		}
	}
}
