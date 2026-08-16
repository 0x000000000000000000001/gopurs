package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_2_colon_30IsSymbol gopurs_runtime.Value
var once_Main_2_colon_30IsSymbol sync.Once

func Get_Main_2_colon_30IsSymbol() gopurs_runtime.Value {
	once_Main_2_colon_30IsSymbol.Do(func() {
		cache_Main_2_colon_30IsSymbol = gopurs_runtime.Value{Type: 9, IntVal: 2134024384, UnsafePtr: unsafe.Pointer((&Constructor_Data_Symbol_IsSymbol{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("2:30")
		})}))}
	})
	return cache_Main_2_colon_30IsSymbol
}

var cache_Main_2IsSymbol gopurs_runtime.Value
var once_Main_2IsSymbol sync.Once

func Get_Main_2IsSymbol() gopurs_runtime.Value {
	once_Main_2IsSymbol.Do(func() {
		cache_Main_2IsSymbol = gopurs_runtime.Value{Type: 9, IntVal: 2134024384, UnsafePtr: unsafe.Pointer((&Constructor_Data_Symbol_IsSymbol{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("2")
		})}))}
	})
	return cache_Main_2IsSymbol
}

var cache_Main_twoThirty gopurs_runtime.Value
var once_Main_twoThirty sync.Once

func Get_Main_twoThirty() gopurs_runtime.Value {
	once_Main_twoThirty.Do(func() {
		cache_Main_twoThirty = gopurs_runtime.Str("2:30")
	})
	return cache_Main_twoThirty
}

var cache_Main_two gopurs_runtime.Value
var once_Main_two sync.Once

func Get_Main_two() gopurs_runtime.Value {
	once_Main_two.Do(func() {
		cache_Main_two = gopurs_runtime.Str("2")
	})
	return cache_Main_two
}

var cache_Main_reflectSymbol_prime_ gopurs_runtime.Value
var once_Main_reflectSymbol_prime_ sync.Once

func Get_Main_reflectSymbol_prime_() gopurs_runtime.Value {
	once_Main_reflectSymbol_prime_.Do(func() {
		cache_Main_reflectSymbol_prime_ = gopurs_runtime.Func(func(dictIsSymbol_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_reflectSymbol_prime_(gopurs_runtime.CoerceToStruct[Constructor_Data_Symbol_IsSymbol](dictIsSymbol_0_box))
		})
	})
	return cache_Main_reflectSymbol_prime_
}

var cache_Main_two2 gopurs_runtime.Value
var once_Main_two2 sync.Once

func Get_Main_two2() gopurs_runtime.Value {
	once_Main_two2.Do(func() {
		cache_Main_two2 = gopurs_runtime.Str("2")
	})
	return cache_Main_two2
}

var cache_Main_twoThirty2 gopurs_runtime.Value
var once_Main_twoThirty2 sync.Once

func Get_Main_twoThirty2() gopurs_runtime.Value {
	once_Main_twoThirty2.Do(func() {
		cache_Main_twoThirty2 = gopurs_runtime.Str("2:30")
	})
	return cache_Main_twoThirty2
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

var cache_Main_reflectSymbol_prime___178557939 gopurs_runtime.Value
var once_Main_reflectSymbol_prime___178557939 sync.Once

func Get_Main_reflectSymbol_prime___178557939() gopurs_runtime.Value {
	once_Main_reflectSymbol_prime___178557939.Do(func() {
		cache_Main_reflectSymbol_prime___178557939 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_reflectSymbol_prime___178557939(__eta0_0_box)
		})
	})
	return cache_Main_reflectSymbol_prime___178557939
}

var cache_Main_reflectSymbol_prime___3961507637 gopurs_runtime.Value
var once_Main_reflectSymbol_prime___3961507637 sync.Once

func Get_Main_reflectSymbol_prime___3961507637() gopurs_runtime.Value {
	once_Main_reflectSymbol_prime___3961507637.Do(func() {
		cache_Main_reflectSymbol_prime___3961507637 = gopurs_runtime.Func(func(dictIsSymbol_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_reflectSymbol_prime___3961507637(gopurs_runtime.CoerceToStruct[Constructor_Data_Symbol_IsSymbol](dictIsSymbol_0_box))
		})
	})
	return cache_Main_reflectSymbol_prime___3961507637
}

func Call_Main_reflectSymbol_prime_(dictIsSymbol_0_loop *Constructor_Data_Symbol_IsSymbol) gopurs_runtime.Value {
	var dictIsSymbol_0 *Constructor_Data_Symbol_IsSymbol = dictIsSymbol_0_loop
	_ = dictIsSymbol_0
	return gopurs_runtime.Box(dictIsSymbol_0.V0)
}

func Call_Main_reflectSymbol_prime___178557939(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
	_ = __eta0_0
	return gopurs_runtime.Str("2")
}

func Call_Main_reflectSymbol_prime___3961507637(dictIsSymbol_0_loop *Constructor_Data_Symbol_IsSymbol) gopurs_runtime.Value {
	var dictIsSymbol_0 *Constructor_Data_Symbol_IsSymbol = dictIsSymbol_0_loop
	_ = dictIsSymbol_0
	return gopurs_runtime.Box(dictIsSymbol_0.V0)
}
