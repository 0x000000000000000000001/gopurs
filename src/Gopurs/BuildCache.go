package Gopurs_BuildCache

import "gopurs/output/gopurs_build_cache"

func Exchange(request string, _ any) string {
	return gopurs_build_cache.Exchange(request)
}
