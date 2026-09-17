package Gopurs_GoAst

import "sync"

// Each partial application owns its cache, as in the JavaScript FFI. Empty
// strings are cached too. The callback is pure; synchronize initialization so
// native compiler workers can share the memoized sanitizer.
func MemoizeName(sanitize func(string) string) func(string) string {
	names := make(map[string]string)
	var mu sync.Mutex
	return func(name string) string {
		mu.Lock()
		defer mu.Unlock()
		if result, found := names[name]; found {
			return result
		}
		result := sanitize(name)
		names[name] = result
		return result
	}
}
