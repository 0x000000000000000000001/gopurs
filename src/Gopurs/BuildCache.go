// Handwritten FFI for Gopurs.BuildCache; keep in sync with the .purs foreign imports.
package Gopurs_BuildCache

import (
	"gopurs/output/gopurs_build_cache"
	"gopurs/output/gopurs_runtime"
)

func Exchange(request string, _ any) string {
	return gopurs_build_cache.Exchange(request)
}

func SessionExchange(request string, _ any) string {
	return gopurs_build_cache.SessionExchange(request)
}

func SessionExchangeOnMiss(request string, onMiss gopurs_runtime.Value, _ any) string {
	return gopurs_build_cache.SessionExchangeOnMiss(request, func(seed string) {
		gopurs_runtime.Apply(gopurs_runtime.Apply(onMiss, gopurs_runtime.Str(seed)), gopurs_runtime.Value{})
	})
}
