package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Type_Proxy_Proxy gopurs_runtime.Value
var once_Type_Proxy_Proxy sync.Once

func Get_Type_Proxy_Proxy() gopurs_runtime.Value {
	once_Type_Proxy_Proxy.Do(func() {
		cache_Type_Proxy_Proxy = gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}
	})
	return cache_Type_Proxy_Proxy
}

type Constructor_Type_Proxy_Proxy[T_a any] struct {
	Rc uint32
}
