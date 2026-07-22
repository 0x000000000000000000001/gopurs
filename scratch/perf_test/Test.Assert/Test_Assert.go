package Test_Assert

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Eq "gopurs/output/Data.Eq"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Show "gopurs/output/Data.Show"
	pkg_Effect "gopurs/output/Effect"
	pkg_Effect_Console "gopurs/output/Effect.Console"
	pkg_Data_Unit "gopurs/output/Data.Unit"
)

var assert_prime gopurs_runtime.Value
var once_assert_prime sync.Once
func Get_assert_prime() gopurs_runtime.Value {
	once_assert_prime.Do(func() {
		assert_prime = Get_assertImpl()
	})
	return assert_prime
}

var assertEqual_prime gopurs_runtime.Value
var once_assertEqual_prime sync.Once
func Get_assertEqual_prime() gopurs_runtime.Value {
	once_assertEqual_prime.Do(func() {
		assertEqual_prime = gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictShow_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(userMessage_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
result_4_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(dictEq_0.PtrVal.(map[string]gopurs_runtime.Value)["eq"], v_3.PtrVal.(map[string]gopurs_runtime.Value)["actual"]), v_3.PtrVal.(map[string]gopurs_runtime.Value)["expected"])
var __t2 gopurs_runtime.Value
if (gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Eq.Get_eqStringImpl(), userMessage_2), gopurs_runtime.Str(""))).IntVal != 0 {
__t2 = gopurs_runtime.Str("")
} else {
__t2 = gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), userMessage_2), gopurs_runtime.Str("\n"))
}
message_5_1 := gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), __t2), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Str("Expected: ")), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Apply(dictShow_1.PtrVal.(map[string]gopurs_runtime.Value)["show"], v_3.PtrVal.(map[string]gopurs_runtime.Value)["expected"])), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Str("\nActual:   ")), gopurs_runtime.Apply(dictShow_1.PtrVal.(map[string]gopurs_runtime.Value)["show"], v_3.PtrVal.(map[string]gopurs_runtime.Value)["actual"])))))
__local_var_6_3 := gopurs_runtime.Apply(pkg_Effect_Console.Get_error(), message_5_1)
var __t4 gopurs_runtime.Value
if (gopurs_runtime.Bool(result_4_0.IntVal == 0)).IntVal != 0 {
__t4 = __local_var_6_3
} else {
if (result_4_0).IntVal != 0 {
__t4 = gopurs_runtime.Apply(pkg_Effect.Get_pureE(), pkg_Data_Unit.Get_unit())
} else {
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Effect.Get_bindE(), __t4), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_assertImpl(), message_5_1), result_4_0)
}))
})
})
})
})
	})
	return assertEqual_prime
}

var assertEqual gopurs_runtime.Value
var once_assertEqual sync.Once
func Get_assertEqual() gopurs_runtime.Value {
	once_assertEqual.Do(func() {
		assertEqual = gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictShow_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_assertEqual_prime(), dictEq_0), dictShow_1), gopurs_runtime.Str(""))
})
})
	})
	return assertEqual
}

var assertFalse gopurs_runtime.Value
var once_assertFalse sync.Once
func Get_assertFalse() gopurs_runtime.Value {
	once_assertFalse.Do(func() {
		assertFalse = gopurs_runtime.Func(func(actual_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_assertEqual_prime(), pkg_Data_Eq.Get_eqBoolean()), pkg_Data_Show.Get_showBoolean()), gopurs_runtime.Str("")), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"actual": actual_0, "expected": gopurs_runtime.Bool(false)}))
})
	})
	return assertFalse
}

var assertTrue gopurs_runtime.Value
var once_assertTrue sync.Once
func Get_assertTrue() gopurs_runtime.Value {
	once_assertTrue.Do(func() {
		assertTrue = gopurs_runtime.Func(func(actual_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_assertEqual_prime(), pkg_Data_Eq.Get_eqBoolean()), pkg_Data_Show.Get_showBoolean()), gopurs_runtime.Str("")), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"actual": actual_0, "expected": gopurs_runtime.Bool(true)}))
})
	})
	return assertTrue
}

var assertFalse_prime gopurs_runtime.Value
var once_assertFalse_prime sync.Once
func Get_assertFalse_prime() gopurs_runtime.Value {
	once_assertFalse_prime.Do(func() {
		assertFalse_prime = gopurs_runtime.Func(func(message_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(actual_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_assertEqual_prime(), pkg_Data_Eq.Get_eqBoolean()), pkg_Data_Show.Get_showBoolean()), message_0), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"actual": actual_1, "expected": gopurs_runtime.Bool(false)}))
})
})
	})
	return assertFalse_prime
}

var assertTrue_prime gopurs_runtime.Value
var once_assertTrue_prime sync.Once
func Get_assertTrue_prime() gopurs_runtime.Value {
	once_assertTrue_prime.Do(func() {
		assertTrue_prime = gopurs_runtime.Func(func(message_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(actual_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_assertEqual_prime(), pkg_Data_Eq.Get_eqBoolean()), pkg_Data_Show.Get_showBoolean()), message_0), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"actual": actual_1, "expected": gopurs_runtime.Bool(true)}))
})
})
	})
	return assertTrue_prime
}

var assertThrows_prime gopurs_runtime.Value
var once_assertThrows_prime sync.Once
func Get_assertThrows_prime() gopurs_runtime.Value {
	once_assertThrows_prime.Do(func() {
		assertThrows_prime = gopurs_runtime.Func(func(msg_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(fn_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Effect.Get_bindE(), gopurs_runtime.Apply(Get_checkThrows(), fn_1)), gopurs_runtime.Apply(Get_assertImpl(), msg_0))
})
})
	})
	return assertThrows_prime
}

var assertThrows gopurs_runtime.Value
var once_assertThrows sync.Once
func Get_assertThrows() gopurs_runtime.Value {
	once_assertThrows.Do(func() {
		assertThrows = gopurs_runtime.Apply(Get_assertThrows_prime(), gopurs_runtime.Str("Assertion failed: An error should have been thrown"))
	})
	return assertThrows
}

var assert gopurs_runtime.Value
var once_assert sync.Once
func Get_assert() gopurs_runtime.Value {
	once_assert.Do(func() {
		assert = gopurs_runtime.Apply(Get_assertImpl(), gopurs_runtime.Str("Assertion failed"))
	})
	return assert
}

func Get_assertImpl() gopurs_runtime.Value {
	return AssertImpl
}

func Get_checkThrows() gopurs_runtime.Value {
	return CheckThrows
}
