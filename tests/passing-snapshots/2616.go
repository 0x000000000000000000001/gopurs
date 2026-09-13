package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_F gopurs_runtime.Value
var once_Main_F sync.Once

func Get_Main_F() gopurs_runtime.Value {
	once_Main_F.Do(func() {
		cache_Main_F = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_F(x_0_box)
		})
	})
	return cache_Main_F
}

var cache_Main_F__153731162 gopurs_runtime.Value
var once_Main_F__153731162 sync.Once

func Get_Main_F__153731162() gopurs_runtime.Value {
	once_Main_F__153731162.Do(func() {
		cache_Main_F__153731162 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_F__153731162(x_0_box)
		})
	})
	return cache_Main_F__153731162
}

var cache_Main_unF gopurs_runtime.Value
var once_Main_unF sync.Once

func Get_Main_unF() gopurs_runtime.Value {
	once_Main_unF.Do(func() {
		cache_Main_unF = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_unF(v_0_box)
		})
	})
	return cache_Main_unF
}

var cache_Main_unF__153731162 gopurs_runtime.Value
var once_Main_unF__153731162 sync.Once

func Get_Main_unF__153731162() gopurs_runtime.Value {
	once_Main_unF__153731162.Do(func() {
		cache_Main_unF__153731162 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_unF__153731162(v_0_box)
		})
	})
	return cache_Main_unF__153731162
}

var cache_Main_functorF gopurs_runtime.Value
var once_Main_functorF sync.Once

func Get_Main_functorF() gopurs_runtime.Value {
	once_Main_functorF.Do(func() {
		cache_Main_functorF = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, m_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.RecordUpdate1(m_1, "x", gopurs_runtime.Apply(f_0, gopurs_runtime.RecordGet(m_1, "x")))
		})}))}
	})
	return cache_Main_functorF
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Str(gopurs_runtime.Apply(Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))}), gopurs_runtime.Str("Done")).StrVal()).StrVal()))
	})
	return cache_Main_main
}

func Call_Main_F(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_F__153731162(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
F__153731162:
	for {
		if false {
			continue F__153731162
		}
		var x_0 gopurs_runtime.Value = x_0_loop
		_ = x_0
		return x_0
	}
}

func Call_Main_unF(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	return v_0
}

func Call_Main_unF__153731162(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
unF__153731162:
	for {
		if false {
			continue unF__153731162
		}
		var v_0 gopurs_runtime.Value = v_0_loop
		_ = v_0
		return v_0
	}
}
