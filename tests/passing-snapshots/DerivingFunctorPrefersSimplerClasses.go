package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_identity gopurs_runtime.Value
var once_Main_identity sync.Once

func Get_Main_identity() gopurs_runtime.Value {
	once_Main_identity.Do(func() {
		cache_Main_identity = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_identity(x_0_box)
		})
	})
	return cache_Main_identity
}

var cache_Main_MonoAndPro gopurs_runtime.Value
var once_Main_MonoAndPro sync.Once

func Get_Main_MonoAndPro() gopurs_runtime.Value {
	once_Main_MonoAndPro.Do(func() {
		cache_Main_MonoAndPro = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_MonoAndPro(x_0_box)
		})
	})
	return cache_Main_MonoAndPro
}

var cache_Main_Test3 gopurs_runtime.Value
var once_Main_Test3 sync.Once

func Get_Main_Test3() gopurs_runtime.Value {
	once_Main_Test3.Do(func() {
		cache_Main_Test3 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_Test3(x_0_box)
		})
	})
	return cache_Main_Test3
}

var cache_Main_MonoAndBi gopurs_runtime.Value
var once_Main_MonoAndBi sync.Once

func Get_Main_MonoAndBi() gopurs_runtime.Value {
	once_Main_MonoAndBi.Do(func() {
		cache_Main_MonoAndBi = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_MonoAndBi(x_0_box)
		})
	})
	return cache_Main_MonoAndBi
}

var cache_Main_Test1 gopurs_runtime.Value
var once_Main_Test1 sync.Once

func Get_Main_Test1() gopurs_runtime.Value {
	once_Main_Test1.Do(func() {
		cache_Main_Test1 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_Test1(x_0_box)
		})
	})
	return cache_Main_Test1
}

var cache_Main_Test4 gopurs_runtime.Value
var once_Main_Test4 sync.Once

func Get_Main_Test4() gopurs_runtime.Value {
	once_Main_Test4.Do(func() {
		cache_Main_Test4 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_Test4(x_0_box)
		})
	})
	return cache_Main_Test4
}

var cache_Main_Test2 gopurs_runtime.Value
var once_Main_Test2 sync.Once

func Get_Main_Test2() gopurs_runtime.Value {
	once_Main_Test2.Do(func() {
		cache_Main_Test2 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_Test2(x_0_box)
		})
	})
	return cache_Main_Test2
}

var cache_Main_profunctorMonoAndPro gopurs_runtime.Value
var once_Main_profunctorMonoAndPro sync.Once

func Get_Main_profunctorMonoAndPro() gopurs_runtime.Value {
	once_Main_profunctorMonoAndPro.Do(func() {
		cache_Main_profunctorMonoAndPro = gopurs_runtime.Value{Type: 9, IntVal: 2367018778, UnsafePtr: unsafe.Pointer((&Constructor_Data_Profunctor_Profunctor{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("Profunctor instance was used but the Functor instance was expected"), gopurs_runtime.Bool(false))
				})
			})
		})}))}
	})
	return cache_Main_profunctorMonoAndPro
}

var cache_Main_profunctorExclusivelyPro gopurs_runtime.Value
var once_Main_profunctorExclusivelyPro sync.Once

func Get_Main_profunctorExclusivelyPro() gopurs_runtime.Value {
	once_Main_profunctorExclusivelyPro.Do(func() {
		cache_Main_profunctorExclusivelyPro = gopurs_runtime.Value{Type: 9, IntVal: 2367018778, UnsafePtr: unsafe.Pointer((&Constructor_Data_Profunctor_Profunctor{1, gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(m_2 gopurs_runtime.Value) gopurs_runtime.Value {
					return func() gopurs_runtime.Value { panic("Failed pattern match") }()
				})
			})
		})}))}
	})
	return cache_Main_profunctorExclusivelyPro
}

var cache_Main_functorTest4 gopurs_runtime.Value
var once_Main_functorTest4 sync.Once

func Get_Main_functorTest4() gopurs_runtime.Value {
	once_Main_functorTest4.Do(func() {
		cache_Main_functorTest4 = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor{1, gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return func() gopurs_runtime.Value { panic("Failed pattern match") }()
			})
		})}))}
	})
	return cache_Main_functorTest4
}

var cache_Main_functorMonoAndPro gopurs_runtime.Value
var once_Main_functorMonoAndPro sync.Once

func Get_Main_functorMonoAndPro() gopurs_runtime.Value {
	once_Main_functorMonoAndPro.Do(func() {
		cache_Main_functorMonoAndPro = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor{1, gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return m_1
			})
		})}))}
	})
	return cache_Main_functorMonoAndPro
}

var cache_Main_functorTest3 gopurs_runtime.Value
var once_Main_functorTest3 sync.Once

func Get_Main_functorTest3() gopurs_runtime.Value {
	once_Main_functorTest3.Do(func() {
		cache_Main_functorTest3 = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor{1, gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return m_1
			})
		})}))}
	})
	return cache_Main_functorTest3
}

var cache_Main_functorMonoAndBi gopurs_runtime.Value
var once_Main_functorMonoAndBi sync.Once

func Get_Main_functorMonoAndBi() gopurs_runtime.Value {
	once_Main_functorMonoAndBi.Do(func() {
		cache_Main_functorMonoAndBi = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor{1, gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return m_1
			})
		})}))}
	})
	return cache_Main_functorMonoAndBi
}

var cache_Main_functorTest1 gopurs_runtime.Value
var once_Main_functorTest1 sync.Once

func Get_Main_functorTest1() gopurs_runtime.Value {
	once_Main_functorTest1.Do(func() {
		cache_Main_functorTest1 = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor{1, gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return m_1
			})
		})}))}
	})
	return cache_Main_functorTest1
}

var cache_Main_bifunctorMonoAndBi gopurs_runtime.Value
var once_Main_bifunctorMonoAndBi sync.Once

func Get_Main_bifunctorMonoAndBi() gopurs_runtime.Value {
	once_Main_bifunctorMonoAndBi.Do(func() {
		cache_Main_bifunctorMonoAndBi = gopurs_runtime.Value{Type: 9, IntVal: 4141114362, UnsafePtr: unsafe.Pointer((&Constructor_Data_Bifunctor_Bifunctor{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("Bifunctor instance was used but the Functor instance was expected"), gopurs_runtime.Bool(false))
				})
			})
		})}))}
	})
	return cache_Main_bifunctorMonoAndBi
}

var cache_Main_bifunctorExclusivelyBi gopurs_runtime.Value
var once_Main_bifunctorExclusivelyBi sync.Once

func Get_Main_bifunctorExclusivelyBi() gopurs_runtime.Value {
	once_Main_bifunctorExclusivelyBi.Do(func() {
		cache_Main_bifunctorExclusivelyBi = gopurs_runtime.Value{Type: 9, IntVal: 4141114362, UnsafePtr: unsafe.Pointer((&Constructor_Data_Bifunctor_Bifunctor{1, gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(m_2 gopurs_runtime.Value) gopurs_runtime.Value {
					return func() gopurs_runtime.Value { panic("Failed pattern match") }()
				})
			})
		})}))}
	})
	return cache_Main_bifunctorExclusivelyBi
}

var cache_Main_functorTest2 gopurs_runtime.Value
var once_Main_functorTest2 sync.Once

func Get_Main_functorTest2() gopurs_runtime.Value {
	once_Main_functorTest2.Do(func() {
		cache_Main_functorTest2 = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor{1, gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return func() gopurs_runtime.Value { panic("Failed pattern match") }()
			})
		})}))}
	})
	return cache_Main_functorTest2
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			_dollar___unused_0_0 := Get_Data_Unit_unit()
			_ = _dollar___unused_0_0
			_dollar___unused_1_1 := Get_Data_Unit_unit()
			_ = _dollar___unused_1_1
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

var cache_Main_functorMonoAndBi__2782926784 gopurs_runtime.Value
var once_Main_functorMonoAndBi__2782926784 sync.Once

func Get_Main_functorMonoAndBi__2782926784() gopurs_runtime.Value {
	once_Main_functorMonoAndBi__2782926784.Do(func() {
		cache_Main_functorMonoAndBi__2782926784 = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor{1, gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return m_1
			})
		})}))}
	})
	return cache_Main_functorMonoAndBi__2782926784
}

var cache_Main_functorMonoAndBi__2068759430 gopurs_runtime.Value
var once_Main_functorMonoAndBi__2068759430 sync.Once

func Get_Main_functorMonoAndBi__2068759430() gopurs_runtime.Value {
	once_Main_functorMonoAndBi__2068759430.Do(func() {
		cache_Main_functorMonoAndBi__2068759430 = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor{1, gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return m_1
			})
		})}))}
	})
	return cache_Main_functorMonoAndBi__2068759430
}

var cache_Main_functorMonoAndPro__2782926784 gopurs_runtime.Value
var once_Main_functorMonoAndPro__2782926784 sync.Once

func Get_Main_functorMonoAndPro__2782926784() gopurs_runtime.Value {
	once_Main_functorMonoAndPro__2782926784.Do(func() {
		cache_Main_functorMonoAndPro__2782926784 = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor{1, gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return m_1
			})
		})}))}
	})
	return cache_Main_functorMonoAndPro__2782926784
}

var cache_Main_functorMonoAndPro__2068759430 gopurs_runtime.Value
var once_Main_functorMonoAndPro__2068759430 sync.Once

func Get_Main_functorMonoAndPro__2068759430() gopurs_runtime.Value {
	once_Main_functorMonoAndPro__2068759430.Do(func() {
		cache_Main_functorMonoAndPro__2068759430 = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor{1, gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return m_1
			})
		})}))}
	})
	return cache_Main_functorMonoAndPro__2068759430
}

var cache_Main_functorTest1__4279421935 gopurs_runtime.Value
var once_Main_functorTest1__4279421935 sync.Once

func Get_Main_functorTest1__4279421935() gopurs_runtime.Value {
	once_Main_functorTest1__4279421935.Do(func() {
		cache_Main_functorTest1__4279421935 = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor{1, gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return m_1
			})
		})}))}
	})
	return cache_Main_functorTest1__4279421935
}

var cache_Main_functorTest3__4279421935 gopurs_runtime.Value
var once_Main_functorTest3__4279421935 sync.Once

func Get_Main_functorTest3__4279421935() gopurs_runtime.Value {
	once_Main_functorTest3__4279421935.Do(func() {
		cache_Main_functorTest3__4279421935 = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor{1, gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return m_1
			})
		})}))}
	})
	return cache_Main_functorTest3__4279421935
}

func Call_Main_identity(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_MonoAndPro(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_Test3(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_MonoAndBi(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_Test1(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_Test4(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_Test2(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}
