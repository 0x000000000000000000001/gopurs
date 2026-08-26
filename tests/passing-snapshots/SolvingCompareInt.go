package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_Proxy gopurs_runtime.Value
var once_Main_Proxy sync.Once

func Get_Main_Proxy() gopurs_runtime.Value {
	once_Main_Proxy.Do(func() {
		cache_Main_Proxy = gopurs_runtime.Value{Type: 9, IntVal: int64(227768594), UnsafePtr: nil}
	})
	return cache_Main_Proxy
}

var cache_Main_AssertIsGT_dollar_Dict gopurs_runtime.Value
var once_Main_AssertIsGT_dollar_Dict sync.Once

func Get_Main_AssertIsGT_dollar_Dict() gopurs_runtime.Value {
	once_Main_AssertIsGT_dollar_Dict.Do(func() {
		cache_Main_AssertIsGT_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_AssertIsGT_dollar_Dict(x_0_box)
		})
	})
	return cache_Main_AssertIsGT_dollar_Dict
}

var cache_Main_assertIsGTGT gopurs_runtime.Value
var once_Main_assertIsGTGT sync.Once

func Get_Main_assertIsGTGT() gopurs_runtime.Value {
	once_Main_assertIsGTGT.Do(func() {
		cache_Main_assertIsGTGT = gopurs_runtime.Value{Type: 9, IntVal: 3071150869, UnsafePtr: unsafe.Pointer((&Constructor_Main_AssertIsGT[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(true)
		})}))}
	})
	return cache_Main_assertIsGTGT
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

var cache_Main_assertLesser gopurs_runtime.Value
var once_Main_assertLesser sync.Once

func Get_Main_assertLesser() gopurs_runtime.Value {
	once_Main_assertLesser.Do(func() {
		cache_Main_assertLesser = gopurs_runtime.Func(func(_dollar___unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_assertLesser(_dollar___unused_0_box)
		})
	})
	return cache_Main_assertLesser
}

var cache_Main_assertLesser1 gopurs_runtime.Value
var once_Main_assertLesser1 sync.Once

func Get_Main_assertLesser1() gopurs_runtime.Value {
	once_Main_assertLesser1.Do(func() {
		cache_Main_assertLesser1 = gopurs_runtime.Value{Type: 9, IntVal: int64(227768594), UnsafePtr: nil}
	})
	return cache_Main_assertLesser1
}

var cache_Main_assertLesser2 gopurs_runtime.Value
var once_Main_assertLesser2 sync.Once

func Get_Main_assertLesser2() gopurs_runtime.Value {
	once_Main_assertLesser2.Do(func() {
		cache_Main_assertLesser2 = gopurs_runtime.Value{Type: 9, IntVal: int64(227768594), UnsafePtr: nil}
	})
	return cache_Main_assertLesser2
}

var cache_Main_assertLesser3 gopurs_runtime.Value
var once_Main_assertLesser3 sync.Once

func Get_Main_assertLesser3() gopurs_runtime.Value {
	once_Main_assertLesser3.Do(func() {
		cache_Main_assertLesser3 = gopurs_runtime.Value{Type: 9, IntVal: int64(227768594), UnsafePtr: nil}
	})
	return cache_Main_assertLesser3
}

var cache_Main_assertLesser4 gopurs_runtime.Value
var once_Main_assertLesser4 sync.Once

func Get_Main_assertLesser4() gopurs_runtime.Value {
	once_Main_assertLesser4.Do(func() {
		cache_Main_assertLesser4 = gopurs_runtime.Value{Type: 9, IntVal: int64(227768594), UnsafePtr: nil}
	})
	return cache_Main_assertLesser4
}

var cache_Main_litLt gopurs_runtime.Value
var once_Main_litLt sync.Once

func Get_Main_litLt() gopurs_runtime.Value {
	once_Main_litLt.Do(func() {
		cache_Main_litLt = gopurs_runtime.Value{Type: 9, IntVal: int64(227768594), UnsafePtr: nil}
	})
	return cache_Main_litLt
}

var cache_Main_litTransLT gopurs_runtime.Value
var once_Main_litTransLT sync.Once

func Get_Main_litTransLT() gopurs_runtime.Value {
	once_Main_litTransLT.Do(func() {
		cache_Main_litTransLT = gopurs_runtime.Func(func(_dollar___unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_litTransLT(_dollar___unused_0_box)
		})
	})
	return cache_Main_litTransLT
}

var cache_Main_litTransRange gopurs_runtime.Value
var once_Main_litTransRange sync.Once

func Get_Main_litTransRange() gopurs_runtime.Value {
	once_Main_litTransRange.Do(func() {
		cache_Main_litTransRange = gopurs_runtime.Func2(func(_dollar___unused_0_box gopurs_runtime.Value, _dollar___unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_litTransRange(_dollar___unused_0_box, _dollar___unused_1_box)
		})
	})
	return cache_Main_litTransRange
}

var cache_Main_symmLt gopurs_runtime.Value
var once_Main_symmLt sync.Once

func Get_Main_symmLt() gopurs_runtime.Value {
	once_Main_symmLt.Do(func() {
		cache_Main_symmLt = gopurs_runtime.Func(func(_dollar___unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_symmLt(_dollar___unused_0_box)
		})
	})
	return cache_Main_symmLt
}

var cache_Main_transEqLt gopurs_runtime.Value
var once_Main_transEqLt sync.Once

func Get_Main_transEqLt() gopurs_runtime.Value {
	once_Main_transEqLt.Do(func() {
		cache_Main_transEqLt = gopurs_runtime.Func3(func(_dollar___unused_0_box gopurs_runtime.Value, _dollar___unused_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_transEqLt(_dollar___unused_0_box, _dollar___unused_1_box, uint32(v_2_box.IntVal))), UnsafePtr: nil}
		})
	})
	return cache_Main_transEqLt
}

var cache_Main_transLt gopurs_runtime.Value
var once_Main_transLt sync.Once

func Get_Main_transLt() gopurs_runtime.Value {
	once_Main_transLt.Do(func() {
		cache_Main_transLt = gopurs_runtime.Func3(func(_dollar___unused_0_box gopurs_runtime.Value, _dollar___unused_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_transLt(_dollar___unused_0_box, _dollar___unused_1_box, uint32(v_2_box.IntVal))), UnsafePtr: nil}
		})
	})
	return cache_Main_transLt
}

var cache_Main_transLtEq gopurs_runtime.Value
var once_Main_transLtEq sync.Once

func Get_Main_transLtEq() gopurs_runtime.Value {
	once_Main_transLtEq.Do(func() {
		cache_Main_transLtEq = gopurs_runtime.Func3(func(_dollar___unused_0_box gopurs_runtime.Value, _dollar___unused_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_transLtEq(_dollar___unused_0_box, _dollar___unused_1_box, uint32(v_2_box.IntVal))), UnsafePtr: nil}
		})
	})
	return cache_Main_transLtEq
}

var cache_Main_transSymmEqLt gopurs_runtime.Value
var once_Main_transSymmEqLt sync.Once

func Get_Main_transSymmEqLt() gopurs_runtime.Value {
	once_Main_transSymmEqLt.Do(func() {
		cache_Main_transSymmEqLt = gopurs_runtime.Func3(func(_dollar___unused_0_box gopurs_runtime.Value, _dollar___unused_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_transSymmEqLt(_dollar___unused_0_box, _dollar___unused_1_box, uint32(v_2_box.IntVal))), UnsafePtr: nil}
		})
	})
	return cache_Main_transSymmEqLt
}

var cache_Main_transSymmLt gopurs_runtime.Value
var once_Main_transSymmLt sync.Once

func Get_Main_transSymmLt() gopurs_runtime.Value {
	once_Main_transSymmLt.Do(func() {
		cache_Main_transSymmLt = gopurs_runtime.Func3(func(_dollar___unused_0_box gopurs_runtime.Value, _dollar___unused_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_transSymmLt(_dollar___unused_0_box, _dollar___unused_1_box, uint32(v_2_box.IntVal))), UnsafePtr: nil}
		})
	})
	return cache_Main_transSymmLt
}

var cache_Main_transSymmLtEq gopurs_runtime.Value
var once_Main_transSymmLtEq sync.Once

func Get_Main_transSymmLtEq() gopurs_runtime.Value {
	once_Main_transSymmLtEq.Do(func() {
		cache_Main_transSymmLtEq = gopurs_runtime.Func3(func(_dollar___unused_0_box gopurs_runtime.Value, _dollar___unused_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_transSymmLtEq(_dollar___unused_0_box, _dollar___unused_1_box, uint32(v_2_box.IntVal))), UnsafePtr: nil}
		})
	})
	return cache_Main_transSymmLtEq
}

var cache_Main_withFacts gopurs_runtime.Value
var once_Main_withFacts sync.Once

func Get_Main_withFacts() gopurs_runtime.Value {
	once_Main_withFacts.Do(func() {
		cache_Main_withFacts = gopurs_runtime.Func2(func(_dollar___unused_0_box gopurs_runtime.Value, _dollar___unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_withFacts(_dollar___unused_0_box, _dollar___unused_1_box)
		})
	})
	return cache_Main_withFacts
}

var cache_Main_assertIsGT gopurs_runtime.Value
var once_Main_assertIsGT sync.Once

func Get_Main_assertIsGT() gopurs_runtime.Value {
	once_Main_assertIsGT.Do(func() {
		cache_Main_assertIsGT = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_assertIsGT(gopurs_runtime.CoerceToStruct[Constructor_Main_AssertIsGT[gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Main_assertIsGT
}

var cache_Main_assertIsGT__865065017 gopurs_runtime.Value
var once_Main_assertIsGT__865065017 sync.Once

func Get_Main_assertIsGT__865065017() gopurs_runtime.Value {
	once_Main_assertIsGT__865065017.Do(func() {
		cache_Main_assertIsGT__865065017 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_assertIsGT__865065017(gopurs_runtime.CoerceToStruct[Constructor_Main_AssertIsGT[gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Main_assertIsGT__865065017
}

var cache_Main_assertIsGT__1442557951 gopurs_runtime.Value
var once_Main_assertIsGT__1442557951 sync.Once

func Get_Main_assertIsGT__1442557951() gopurs_runtime.Value {
	once_Main_assertIsGT__1442557951.Do(func() {
		cache_Main_assertIsGT__1442557951 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_assertIsGT__1442557951(gopurs_runtime.CoerceToStruct[Constructor_Main_AssertIsGT[gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Main_assertIsGT__1442557951
}

var cache_Main_infer gopurs_runtime.Value
var once_Main_infer sync.Once

func Get_Main_infer() gopurs_runtime.Value {
	once_Main_infer.Do(func() {
		cache_Main_infer = gopurs_runtime.Func4(func(_dollar___unused_0_box gopurs_runtime.Value, dictAssertIsGT_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value, v1_3_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(Call_Main_infer(_dollar___unused_0_box, gopurs_runtime.CoerceToStruct[Constructor_Main_AssertIsGT[gopurs_runtime.Value]](dictAssertIsGT_1_box), uint32(v_2_box.IntVal), uint32(v1_3_box.IntVal)))
		})
	})
	return cache_Main_infer
}

var cache_Main_infer__668843825 gopurs_runtime.Value
var once_Main_infer__668843825 sync.Once

func Get_Main_infer__668843825() gopurs_runtime.Value {
	once_Main_infer__668843825.Do(func() {
		cache_Main_infer__668843825 = gopurs_runtime.Func4(func(_dollar___unused_0_box gopurs_runtime.Value, dictAssertIsGT_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value, v1_3_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(Call_Main_infer__668843825(_dollar___unused_0_box, gopurs_runtime.CoerceToStruct[Constructor_Main_AssertIsGT[gopurs_runtime.Value]](dictAssertIsGT_1_box), uint32(v_2_box.IntVal), uint32(v1_3_box.IntVal)))
		})
	})
	return cache_Main_infer__668843825
}

var cache_Main_inferSolved gopurs_runtime.Value
var once_Main_inferSolved sync.Once

func Get_Main_inferSolved() gopurs_runtime.Value {
	once_Main_inferSolved.Do(func() {
		cache_Main_inferSolved = gopurs_runtime.Func5(func(_dollar___unused_0_box gopurs_runtime.Value, _dollar___unused_1_box gopurs_runtime.Value, m_2_box gopurs_runtime.Value, v_3_box gopurs_runtime.Value, p_4_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(Call_Main_inferSolved(_dollar___unused_0_box, _dollar___unused_1_box, uint32(m_2_box.IntVal), uint32(v_3_box.IntVal), uint32(p_4_box.IntVal)))
		})
	})
	return cache_Main_inferSolved
}

var cache_Main_assertGreater gopurs_runtime.Value
var once_Main_assertGreater sync.Once

func Get_Main_assertGreater() gopurs_runtime.Value {
	once_Main_assertGreater.Do(func() {
		cache_Main_assertGreater = gopurs_runtime.Func(func(_dollar___unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_assertGreater(_dollar___unused_0_box)
		})
	})
	return cache_Main_assertGreater
}

var cache_Main_assertGreater1 gopurs_runtime.Value
var once_Main_assertGreater1 sync.Once

func Get_Main_assertGreater1() gopurs_runtime.Value {
	once_Main_assertGreater1.Do(func() {
		cache_Main_assertGreater1 = gopurs_runtime.Value{Type: 9, IntVal: int64(227768594), UnsafePtr: nil}
	})
	return cache_Main_assertGreater1
}

var cache_Main_assertGreater2 gopurs_runtime.Value
var once_Main_assertGreater2 sync.Once

func Get_Main_assertGreater2() gopurs_runtime.Value {
	once_Main_assertGreater2.Do(func() {
		cache_Main_assertGreater2 = gopurs_runtime.Value{Type: 9, IntVal: int64(227768594), UnsafePtr: nil}
	})
	return cache_Main_assertGreater2
}

var cache_Main_assertGreater3 gopurs_runtime.Value
var once_Main_assertGreater3 sync.Once

func Get_Main_assertGreater3() gopurs_runtime.Value {
	once_Main_assertGreater3.Do(func() {
		cache_Main_assertGreater3 = gopurs_runtime.Value{Type: 9, IntVal: int64(227768594), UnsafePtr: nil}
	})
	return cache_Main_assertGreater3
}

var cache_Main_litGt gopurs_runtime.Value
var once_Main_litGt sync.Once

func Get_Main_litGt() gopurs_runtime.Value {
	once_Main_litGt.Do(func() {
		cache_Main_litGt = gopurs_runtime.Value{Type: 9, IntVal: int64(227768594), UnsafePtr: nil}
	})
	return cache_Main_litGt
}

var cache_Main_litTransGT gopurs_runtime.Value
var once_Main_litTransGT sync.Once

func Get_Main_litTransGT() gopurs_runtime.Value {
	once_Main_litTransGT.Do(func() {
		cache_Main_litTransGT = gopurs_runtime.Func(func(_dollar___unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_litTransGT(_dollar___unused_0_box)
		})
	})
	return cache_Main_litTransGT
}

var cache_Main_symmGt gopurs_runtime.Value
var once_Main_symmGt sync.Once

func Get_Main_symmGt() gopurs_runtime.Value {
	once_Main_symmGt.Do(func() {
		cache_Main_symmGt = gopurs_runtime.Func(func(_dollar___unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_symmGt(_dollar___unused_0_box)
		})
	})
	return cache_Main_symmGt
}

var cache_Main_transEqGt gopurs_runtime.Value
var once_Main_transEqGt sync.Once

func Get_Main_transEqGt() gopurs_runtime.Value {
	once_Main_transEqGt.Do(func() {
		cache_Main_transEqGt = gopurs_runtime.Func3(func(_dollar___unused_0_box gopurs_runtime.Value, _dollar___unused_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_transEqGt(_dollar___unused_0_box, _dollar___unused_1_box, uint32(v_2_box.IntVal))), UnsafePtr: nil}
		})
	})
	return cache_Main_transEqGt
}

var cache_Main_transGt gopurs_runtime.Value
var once_Main_transGt sync.Once

func Get_Main_transGt() gopurs_runtime.Value {
	once_Main_transGt.Do(func() {
		cache_Main_transGt = gopurs_runtime.Func3(func(_dollar___unused_0_box gopurs_runtime.Value, _dollar___unused_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_transGt(_dollar___unused_0_box, _dollar___unused_1_box, uint32(v_2_box.IntVal))), UnsafePtr: nil}
		})
	})
	return cache_Main_transGt
}

var cache_Main_transGtEq gopurs_runtime.Value
var once_Main_transGtEq sync.Once

func Get_Main_transGtEq() gopurs_runtime.Value {
	once_Main_transGtEq.Do(func() {
		cache_Main_transGtEq = gopurs_runtime.Func3(func(_dollar___unused_0_box gopurs_runtime.Value, _dollar___unused_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_transGtEq(_dollar___unused_0_box, _dollar___unused_1_box, uint32(v_2_box.IntVal))), UnsafePtr: nil}
		})
	})
	return cache_Main_transGtEq
}

var cache_Main_transSymmEqGt gopurs_runtime.Value
var once_Main_transSymmEqGt sync.Once

func Get_Main_transSymmEqGt() gopurs_runtime.Value {
	once_Main_transSymmEqGt.Do(func() {
		cache_Main_transSymmEqGt = gopurs_runtime.Func3(func(_dollar___unused_0_box gopurs_runtime.Value, _dollar___unused_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_transSymmEqGt(_dollar___unused_0_box, _dollar___unused_1_box, uint32(v_2_box.IntVal))), UnsafePtr: nil}
		})
	})
	return cache_Main_transSymmEqGt
}

var cache_Main_transSymmGt gopurs_runtime.Value
var once_Main_transSymmGt sync.Once

func Get_Main_transSymmGt() gopurs_runtime.Value {
	once_Main_transSymmGt.Do(func() {
		cache_Main_transSymmGt = gopurs_runtime.Func3(func(_dollar___unused_0_box gopurs_runtime.Value, _dollar___unused_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_transSymmGt(_dollar___unused_0_box, _dollar___unused_1_box, uint32(v_2_box.IntVal))), UnsafePtr: nil}
		})
	})
	return cache_Main_transSymmGt
}

var cache_Main_transSymmGtEq gopurs_runtime.Value
var once_Main_transSymmGtEq sync.Once

func Get_Main_transSymmGtEq() gopurs_runtime.Value {
	once_Main_transSymmGtEq.Do(func() {
		cache_Main_transSymmGtEq = gopurs_runtime.Func3(func(_dollar___unused_0_box gopurs_runtime.Value, _dollar___unused_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_transSymmGtEq(_dollar___unused_0_box, _dollar___unused_1_box, uint32(v_2_box.IntVal))), UnsafePtr: nil}
		})
	})
	return cache_Main_transSymmGtEq
}

var cache_Main_assertEqual gopurs_runtime.Value
var once_Main_assertEqual sync.Once

func Get_Main_assertEqual() gopurs_runtime.Value {
	once_Main_assertEqual.Do(func() {
		cache_Main_assertEqual = gopurs_runtime.Func(func(_dollar___unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_assertEqual(_dollar___unused_0_box)
		})
	})
	return cache_Main_assertEqual
}

var cache_Main_assertEqual1 gopurs_runtime.Value
var once_Main_assertEqual1 sync.Once

func Get_Main_assertEqual1() gopurs_runtime.Value {
	once_Main_assertEqual1.Do(func() {
		cache_Main_assertEqual1 = gopurs_runtime.Value{Type: 9, IntVal: int64(227768594), UnsafePtr: nil}
	})
	return cache_Main_assertEqual1
}

var cache_Main_assertEqual2 gopurs_runtime.Value
var once_Main_assertEqual2 sync.Once

func Get_Main_assertEqual2() gopurs_runtime.Value {
	once_Main_assertEqual2.Do(func() {
		cache_Main_assertEqual2 = gopurs_runtime.Value{Type: 9, IntVal: int64(227768594), UnsafePtr: nil}
	})
	return cache_Main_assertEqual2
}

var cache_Main_litEq gopurs_runtime.Value
var once_Main_litEq sync.Once

func Get_Main_litEq() gopurs_runtime.Value {
	once_Main_litEq.Do(func() {
		cache_Main_litEq = gopurs_runtime.Value{Type: 9, IntVal: int64(227768594), UnsafePtr: nil}
	})
	return cache_Main_litEq
}

var cache_Main_reflEq gopurs_runtime.Value
var once_Main_reflEq sync.Once

func Get_Main_reflEq() gopurs_runtime.Value {
	once_Main_reflEq.Do(func() {
		cache_Main_reflEq = gopurs_runtime.Value{Type: 9, IntVal: int64(227768594), UnsafePtr: nil}
	})
	return cache_Main_reflEq
}

var cache_Main_symmEq gopurs_runtime.Value
var once_Main_symmEq sync.Once

func Get_Main_symmEq() gopurs_runtime.Value {
	once_Main_symmEq.Do(func() {
		cache_Main_symmEq = gopurs_runtime.Func(func(_dollar___unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_symmEq(_dollar___unused_0_box)
		})
	})
	return cache_Main_symmEq
}

var cache_Main_transEq gopurs_runtime.Value
var once_Main_transEq sync.Once

func Get_Main_transEq() gopurs_runtime.Value {
	once_Main_transEq.Do(func() {
		cache_Main_transEq = gopurs_runtime.Func3(func(_dollar___unused_0_box gopurs_runtime.Value, _dollar___unused_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_transEq(_dollar___unused_0_box, _dollar___unused_1_box, uint32(v_2_box.IntVal))), UnsafePtr: nil}
		})
	})
	return cache_Main_transEq
}

var cache_Main_transSymmEq gopurs_runtime.Value
var once_Main_transSymmEq sync.Once

func Get_Main_transSymmEq() gopurs_runtime.Value {
	once_Main_transSymmEq.Do(func() {
		cache_Main_transSymmEq = gopurs_runtime.Func3(func(_dollar___unused_0_box gopurs_runtime.Value, _dollar___unused_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_transSymmEq(_dollar___unused_0_box, _dollar___unused_1_box, uint32(v_2_box.IntVal))), UnsafePtr: nil}
		})
	})
	return cache_Main_transSymmEq
}

type Constructor_Main_Proxy[T_n any] struct {
	Rc uint32
}

type Constructor_Main_AssertIsGT[T_o any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

func init() {
	gopurs_runtime.StructGetters[3071150869] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_AssertIsGT[any])(ptr)
		_ = c
		switch key {
		case "assertIsGT":
			return gopurs_runtime.Box(c.V0)
		default:
			panic("Key not found in dictionary Constructor_Main_AssertIsGT: " + key)
		}
	}
}

func Call_Main_AssertIsGT_dollar_Dict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_assertLesser(_dollar___unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	return gopurs_runtime.Value{Type: 9, IntVal: int64(227768594), UnsafePtr: nil}
}

func Call_Main_litTransLT(_dollar___unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	return gopurs_runtime.Value{Type: 9, IntVal: int64(227768594), UnsafePtr: nil}
}

func Call_Main_litTransRange(_dollar___unused_0_loop gopurs_runtime.Value, _dollar___unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var _dollar___unused_1 gopurs_runtime.Value = _dollar___unused_1_loop
	_ = _dollar___unused_1
	return gopurs_runtime.Value{Type: 9, IntVal: int64(227768594), UnsafePtr: nil}
}

func Call_Main_symmLt(_dollar___unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	return gopurs_runtime.Value{Type: 9, IntVal: int64(227768594), UnsafePtr: nil}
}

func Call_Main_transEqLt(_dollar___unused_0_loop gopurs_runtime.Value, _dollar___unused_1_loop gopurs_runtime.Value, v_2_loop uint32) uint32 {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var _dollar___unused_1 gopurs_runtime.Value = _dollar___unused_1_loop
	_ = _dollar___unused_1
	var v_2 uint32 = v_2_loop
	_ = v_2
	return 227768594
}

func Call_Main_transLt(_dollar___unused_0_loop gopurs_runtime.Value, _dollar___unused_1_loop gopurs_runtime.Value, v_2_loop uint32) uint32 {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var _dollar___unused_1 gopurs_runtime.Value = _dollar___unused_1_loop
	_ = _dollar___unused_1
	var v_2 uint32 = v_2_loop
	_ = v_2
	return 227768594
}

func Call_Main_transLtEq(_dollar___unused_0_loop gopurs_runtime.Value, _dollar___unused_1_loop gopurs_runtime.Value, v_2_loop uint32) uint32 {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var _dollar___unused_1 gopurs_runtime.Value = _dollar___unused_1_loop
	_ = _dollar___unused_1
	var v_2 uint32 = v_2_loop
	_ = v_2
	return 227768594
}

func Call_Main_transSymmEqLt(_dollar___unused_0_loop gopurs_runtime.Value, _dollar___unused_1_loop gopurs_runtime.Value, v_2_loop uint32) uint32 {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var _dollar___unused_1 gopurs_runtime.Value = _dollar___unused_1_loop
	_ = _dollar___unused_1
	var v_2 uint32 = v_2_loop
	_ = v_2
	return 227768594
}

func Call_Main_transSymmLt(_dollar___unused_0_loop gopurs_runtime.Value, _dollar___unused_1_loop gopurs_runtime.Value, v_2_loop uint32) uint32 {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var _dollar___unused_1 gopurs_runtime.Value = _dollar___unused_1_loop
	_ = _dollar___unused_1
	var v_2 uint32 = v_2_loop
	_ = v_2
	return 227768594
}

func Call_Main_transSymmLtEq(_dollar___unused_0_loop gopurs_runtime.Value, _dollar___unused_1_loop gopurs_runtime.Value, v_2_loop uint32) uint32 {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var _dollar___unused_1 gopurs_runtime.Value = _dollar___unused_1_loop
	_ = _dollar___unused_1
	var v_2 uint32 = v_2_loop
	_ = v_2
	return 227768594
}

func Call_Main_withFacts(_dollar___unused_0_loop gopurs_runtime.Value, _dollar___unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var _dollar___unused_1 gopurs_runtime.Value = _dollar___unused_1_loop
	_ = _dollar___unused_1
	return gopurs_runtime.Value{Type: 9, IntVal: int64(227768594), UnsafePtr: nil}
}

func Call_Main_assertIsGT(dict_0_loop *Constructor_Main_AssertIsGT[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_AssertIsGT[gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return gopurs_runtime.Box(dict_0.V0)
}

func Call_Main_assertIsGT__865065017(dict_0_loop *Constructor_Main_AssertIsGT[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_AssertIsGT[gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return gopurs_runtime.Box(dict_0.V0)
}

func Call_Main_assertIsGT__1442557951(dict_0_loop *Constructor_Main_AssertIsGT[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_AssertIsGT[gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return gopurs_runtime.Box(dict_0.V0)
}

func Call_Main_infer(_dollar___unused_0_loop gopurs_runtime.Value, dictAssertIsGT_1_loop *Constructor_Main_AssertIsGT[gopurs_runtime.Value], v_2_loop uint32, v1_3_loop uint32) bool {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var dictAssertIsGT_1 *Constructor_Main_AssertIsGT[gopurs_runtime.Value] = dictAssertIsGT_1_loop
	_ = dictAssertIsGT_1
	var v_2 uint32 = v_2_loop
	_ = v_2
	var v1_3 uint32 = v1_3_loop
	_ = v1_3
	return (gopurs_runtime.Apply(gopurs_runtime.Box(dictAssertIsGT_1.V0), gopurs_runtime.Value{Type: 9, IntVal: int64(227768594), UnsafePtr: nil}).IntVal) != (0)
}

func Call_Main_infer__668843825(_dollar___unused_0_loop gopurs_runtime.Value, dictAssertIsGT_1_loop *Constructor_Main_AssertIsGT[gopurs_runtime.Value], v_2_loop uint32, v1_3_loop uint32) bool {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var dictAssertIsGT_1 *Constructor_Main_AssertIsGT[gopurs_runtime.Value] = dictAssertIsGT_1_loop
	_ = dictAssertIsGT_1
	var v_2 uint32 = v_2_loop
	_ = v_2
	var v1_3 uint32 = v1_3_loop
	_ = v1_3
	return (gopurs_runtime.Apply(gopurs_runtime.Box(dictAssertIsGT_1.V0), gopurs_runtime.Value{Type: 9, IntVal: int64(227768594), UnsafePtr: nil}).IntVal) != (0)
}

func Call_Main_inferSolved(_dollar___unused_0_loop gopurs_runtime.Value, _dollar___unused_1_loop gopurs_runtime.Value, m_2_loop uint32, v_3_loop uint32, p_4_loop uint32) bool {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var _dollar___unused_1 gopurs_runtime.Value = _dollar___unused_1_loop
	_ = _dollar___unused_1
	var m_2 uint32 = m_2_loop
	_ = m_2
	var v_3 uint32 = v_3_loop
	_ = v_3
	var p_4 uint32 = p_4_loop
	_ = p_4
	return true
}

func Call_Main_assertGreater(_dollar___unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	return gopurs_runtime.Value{Type: 9, IntVal: int64(227768594), UnsafePtr: nil}
}

func Call_Main_litTransGT(_dollar___unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	return gopurs_runtime.Value{Type: 9, IntVal: int64(227768594), UnsafePtr: nil}
}

func Call_Main_symmGt(_dollar___unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	return gopurs_runtime.Value{Type: 9, IntVal: int64(227768594), UnsafePtr: nil}
}

func Call_Main_transEqGt(_dollar___unused_0_loop gopurs_runtime.Value, _dollar___unused_1_loop gopurs_runtime.Value, v_2_loop uint32) uint32 {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var _dollar___unused_1 gopurs_runtime.Value = _dollar___unused_1_loop
	_ = _dollar___unused_1
	var v_2 uint32 = v_2_loop
	_ = v_2
	return 227768594
}

func Call_Main_transGt(_dollar___unused_0_loop gopurs_runtime.Value, _dollar___unused_1_loop gopurs_runtime.Value, v_2_loop uint32) uint32 {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var _dollar___unused_1 gopurs_runtime.Value = _dollar___unused_1_loop
	_ = _dollar___unused_1
	var v_2 uint32 = v_2_loop
	_ = v_2
	return 227768594
}

func Call_Main_transGtEq(_dollar___unused_0_loop gopurs_runtime.Value, _dollar___unused_1_loop gopurs_runtime.Value, v_2_loop uint32) uint32 {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var _dollar___unused_1 gopurs_runtime.Value = _dollar___unused_1_loop
	_ = _dollar___unused_1
	var v_2 uint32 = v_2_loop
	_ = v_2
	return 227768594
}

func Call_Main_transSymmEqGt(_dollar___unused_0_loop gopurs_runtime.Value, _dollar___unused_1_loop gopurs_runtime.Value, v_2_loop uint32) uint32 {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var _dollar___unused_1 gopurs_runtime.Value = _dollar___unused_1_loop
	_ = _dollar___unused_1
	var v_2 uint32 = v_2_loop
	_ = v_2
	return 227768594
}

func Call_Main_transSymmGt(_dollar___unused_0_loop gopurs_runtime.Value, _dollar___unused_1_loop gopurs_runtime.Value, v_2_loop uint32) uint32 {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var _dollar___unused_1 gopurs_runtime.Value = _dollar___unused_1_loop
	_ = _dollar___unused_1
	var v_2 uint32 = v_2_loop
	_ = v_2
	return 227768594
}

func Call_Main_transSymmGtEq(_dollar___unused_0_loop gopurs_runtime.Value, _dollar___unused_1_loop gopurs_runtime.Value, v_2_loop uint32) uint32 {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var _dollar___unused_1 gopurs_runtime.Value = _dollar___unused_1_loop
	_ = _dollar___unused_1
	var v_2 uint32 = v_2_loop
	_ = v_2
	return 227768594
}

func Call_Main_assertEqual(_dollar___unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	return gopurs_runtime.Value{Type: 9, IntVal: int64(227768594), UnsafePtr: nil}
}

func Call_Main_symmEq(_dollar___unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	return gopurs_runtime.Value{Type: 9, IntVal: int64(227768594), UnsafePtr: nil}
}

func Call_Main_transEq(_dollar___unused_0_loop gopurs_runtime.Value, _dollar___unused_1_loop gopurs_runtime.Value, v_2_loop uint32) uint32 {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var _dollar___unused_1 gopurs_runtime.Value = _dollar___unused_1_loop
	_ = _dollar___unused_1
	var v_2 uint32 = v_2_loop
	_ = v_2
	return 227768594
}

func Call_Main_transSymmEq(_dollar___unused_0_loop gopurs_runtime.Value, _dollar___unused_1_loop gopurs_runtime.Value, v_2_loop uint32) uint32 {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var _dollar___unused_1 gopurs_runtime.Value = _dollar___unused_1_loop
	_ = _dollar___unused_1
	var v_2 uint32 = v_2_loop
	_ = v_2
	return 227768594
}
