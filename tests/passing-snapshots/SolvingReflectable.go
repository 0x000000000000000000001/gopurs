package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_refString gopurs_runtime.Value
var once_Main_refString sync.Once

func Get_Main_refString() gopurs_runtime.Value {
	once_Main_refString.Do(func() {
		cache_Main_refString = gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}
	})
	return cache_Main_refString
}

var cache_Main_refStringPass gopurs_runtime.Value
var once_Main_refStringPass sync.Once

func Get_Main_refStringPass() gopurs_runtime.Value {
	once_Main_refStringPass.Do(func() {
		cache_Main_refStringPass = gopurs_runtime.Bool(true)
	})
	return cache_Main_refStringPass
}

var cache_Main_refOrderingLT gopurs_runtime.Value
var once_Main_refOrderingLT sync.Once

func Get_Main_refOrderingLT() gopurs_runtime.Value {
	once_Main_refOrderingLT.Do(func() {
		cache_Main_refOrderingLT = gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}
	})
	return cache_Main_refOrderingLT
}

var cache_Main_refOrderingGT gopurs_runtime.Value
var once_Main_refOrderingGT sync.Once

func Get_Main_refOrderingGT() gopurs_runtime.Value {
	once_Main_refOrderingGT.Do(func() {
		cache_Main_refOrderingGT = gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}
	})
	return cache_Main_refOrderingGT
}

var cache_Main_refOrderingEQ gopurs_runtime.Value
var once_Main_refOrderingEQ sync.Once

func Get_Main_refOrderingEQ() gopurs_runtime.Value {
	once_Main_refOrderingEQ.Do(func() {
		cache_Main_refOrderingEQ = gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}
	})
	return cache_Main_refOrderingEQ
}

var cache_Main_refOrderingPass gopurs_runtime.Value
var once_Main_refOrderingPass sync.Once

func Get_Main_refOrderingPass() gopurs_runtime.Value {
	once_Main_refOrderingPass.Do(func() {
		cache_Main_refOrderingPass = gopurs_runtime.Bool(true)
	})
	return cache_Main_refOrderingPass
}

var cache_Main_refInt gopurs_runtime.Value
var once_Main_refInt sync.Once

func Get_Main_refInt() gopurs_runtime.Value {
	once_Main_refInt.Do(func() {
		cache_Main_refInt = gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}
	})
	return cache_Main_refInt
}

var cache_Main_refIntPass gopurs_runtime.Value
var once_Main_refIntPass sync.Once

func Get_Main_refIntPass() gopurs_runtime.Value {
	once_Main_refIntPass.Do(func() {
		cache_Main_refIntPass = gopurs_runtime.Bool(true)
	})
	return cache_Main_refIntPass
}

var cache_Main_refBooleanT gopurs_runtime.Value
var once_Main_refBooleanT sync.Once

func Get_Main_refBooleanT() gopurs_runtime.Value {
	once_Main_refBooleanT.Do(func() {
		cache_Main_refBooleanT = gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}
	})
	return cache_Main_refBooleanT
}

var cache_Main_refBooleanF gopurs_runtime.Value
var once_Main_refBooleanF sync.Once

func Get_Main_refBooleanF() gopurs_runtime.Value {
	once_Main_refBooleanF.Do(func() {
		cache_Main_refBooleanF = gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}
	})
	return cache_Main_refBooleanF
}

var cache_Main_refBooleanPass gopurs_runtime.Value
var once_Main_refBooleanPass sync.Once

func Get_Main_refBooleanPass() gopurs_runtime.Value {
	once_Main_refBooleanPass.Do(func() {
		cache_Main_refBooleanPass = gopurs_runtime.Bool(true)
	})
	return cache_Main_refBooleanPass
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}
