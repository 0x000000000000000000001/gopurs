// Handwritten FFI for Gopurs.BuildCache; keep in sync with the .purs foreign imports.
package Gopurs_BuildCache

import "gopurs/output/gopurs_build_cache"

func Exchange(request string, _ any) string {
	return gopurs_build_cache.Exchange(request)
}

func SessionExchange(request string, _ any) string {
	return gopurs_build_cache.SessionExchange(request)
}

func SessionExchangeOnMiss(request string, onMiss func(string) func(), _ any) string {
	return gopurs_build_cache.SessionExchangeOnMiss(request, func(seed string) {
		onMiss(seed)()
	})
}
