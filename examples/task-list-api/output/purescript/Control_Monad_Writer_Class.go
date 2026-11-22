package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Monad_Writer_Class_MonadTell_dollar_Dict gopurs_runtime.Value
var once_Control_Monad_Writer_Class_MonadTell_dollar_Dict sync.Once

func Get_Control_Monad_Writer_Class_MonadTell_dollar_Dict() gopurs_runtime.Value {
	once_Control_Monad_Writer_Class_MonadTell_dollar_Dict.Do(func() {
		cache_Control_Monad_Writer_Class_MonadTell_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 551781469, UnsafePtr: unsafe.Pointer(Call_Control_Monad_Writer_Class_MonadTell_dollar_Dict(func() struct {
				Monad1     gopurs_runtime.Value
				Semigroup0 gopurs_runtime.Value
				tell       gopurs_runtime.Value
				/*gopurs:record-v1:3:004d006f006e006100640031,00530065006d006900670072006f007500700030,00740065006c006c*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Monad1     gopurs_runtime.Value
					Semigroup0 gopurs_runtime.Value
					tell       gopurs_runtime.Value
					/*gopurs:record-v1:3:004d006f006e006100640031,00530065006d006900670072006f007500700030,00740065006c006c*/
				}{}
				clone.Monad1 = gopurs_runtime.RecordGet(orig, "Monad1")
				clone.Semigroup0 = gopurs_runtime.RecordGet(orig, "Semigroup0")
				clone.tell = gopurs_runtime.RecordGet(orig, "tell")
				return clone
			}()))}
		})
	})
	return cache_Control_Monad_Writer_Class_MonadTell_dollar_Dict
}

var cache_Control_Monad_Writer_Class_MonadWriter_dollar_Dict gopurs_runtime.Value
var once_Control_Monad_Writer_Class_MonadWriter_dollar_Dict sync.Once

func Get_Control_Monad_Writer_Class_MonadWriter_dollar_Dict() gopurs_runtime.Value {
	once_Control_Monad_Writer_Class_MonadWriter_dollar_Dict.Do(func() {
		cache_Control_Monad_Writer_Class_MonadWriter_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 784743459, UnsafePtr: unsafe.Pointer(Call_Control_Monad_Writer_Class_MonadWriter_dollar_Dict(func() struct {
				MonadTell1 gopurs_runtime.Value
				Monoid0    gopurs_runtime.Value
				listen     gopurs_runtime.Value
				pass       gopurs_runtime.Value
				/*gopurs:record-v1:4:004d006f006e0061006400540065006c006c0031,004d006f006e006f006900640030,006c0069007300740065006e,0070006100730073*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					MonadTell1 gopurs_runtime.Value
					Monoid0    gopurs_runtime.Value
					listen     gopurs_runtime.Value
					pass       gopurs_runtime.Value
					/*gopurs:record-v1:4:004d006f006e0061006400540065006c006c0031,004d006f006e006f006900640030,006c0069007300740065006e,0070006100730073*/
				}{}
				clone.MonadTell1 = gopurs_runtime.RecordGet(orig, "MonadTell1")
				clone.Monoid0 = gopurs_runtime.RecordGet(orig, "Monoid0")
				clone.listen = gopurs_runtime.RecordGet(orig, "listen")
				clone.pass = gopurs_runtime.RecordGet(orig, "pass")
				return clone
			}()))}
		})
	})
	return cache_Control_Monad_Writer_Class_MonadWriter_dollar_Dict
}

var cache_Control_Monad_Writer_Class_tell gopurs_runtime.Value
var once_Control_Monad_Writer_Class_tell sync.Once

func Get_Control_Monad_Writer_Class_tell() gopurs_runtime.Value {
	once_Control_Monad_Writer_Class_tell.Do(func() {
		cache_Control_Monad_Writer_Class_tell = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Writer_Class_tell(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Writer_Class_MonadTell[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Control_Monad_Writer_Class_tell
}

var cache_Control_Monad_Writer_Class_pass gopurs_runtime.Value
var once_Control_Monad_Writer_Class_pass sync.Once

func Get_Control_Monad_Writer_Class_pass() gopurs_runtime.Value {
	once_Control_Monad_Writer_Class_pass.Do(func() {
		cache_Control_Monad_Writer_Class_pass = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Writer_Class_pass(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Writer_Class_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Control_Monad_Writer_Class_pass
}

var cache_Control_Monad_Writer_Class_listen gopurs_runtime.Value
var once_Control_Monad_Writer_Class_listen sync.Once

func Get_Control_Monad_Writer_Class_listen() gopurs_runtime.Value {
	once_Control_Monad_Writer_Class_listen.Do(func() {
		cache_Control_Monad_Writer_Class_listen = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Writer_Class_listen(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Writer_Class_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Control_Monad_Writer_Class_listen
}

var cache_Control_Monad_Writer_Class_listens gopurs_runtime.Value
var once_Control_Monad_Writer_Class_listens sync.Once

func Get_Control_Monad_Writer_Class_listens() gopurs_runtime.Value {
	once_Control_Monad_Writer_Class_listens.Do(func() {
		cache_Control_Monad_Writer_Class_listens = gopurs_runtime.Func(func(dictMonadWriter_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Writer_Class_listens(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Writer_Class_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value]](dictMonadWriter_0_box))
		})
	})
	return cache_Control_Monad_Writer_Class_listens
}

var cache_Control_Monad_Writer_Class_censor gopurs_runtime.Value
var once_Control_Monad_Writer_Class_censor sync.Once

func Get_Control_Monad_Writer_Class_censor() gopurs_runtime.Value {
	once_Control_Monad_Writer_Class_censor.Do(func() {
		cache_Control_Monad_Writer_Class_censor = gopurs_runtime.Func(func(dictMonadWriter_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Writer_Class_censor(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Writer_Class_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value]](dictMonadWriter_0_box))
		})
	})
	return cache_Control_Monad_Writer_Class_censor
}

type Constructor_Control_Monad_Writer_Class_MonadTell[T_w any, T_m any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
}

//gopurs:init-v1:00670065006e006500720061007400650064002f0043006f006e00740072006f006c005f004d006f006e00610064005f005700720069007400650072005f0043006c006100730073:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f0043006f006e00740072006f006c005f004d006f006e00610064005f005700720069007400650072005f0043006c006100730073005f004d006f006e0061006400540065006c006c
func init() {
	gopurs_runtime.StructGetters[551781469] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Control_Monad_Writer_Class_MonadTell[gopurs_runtime.Value, gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "Monad1":
			return gopurs_runtime.Box(c.V0)
		case "Semigroup0":
			return gopurs_runtime.Box(c.V1)
		case "tell":
			return gopurs_runtime.Box(c.V2)
		default:
			panic("Key not found in dictionary Constructor_Control_Monad_Writer_Class_MonadTell: " + key)
		}
	}
}

type Constructor_Control_Monad_Writer_Class_MonadWriter[T_w any, T_m any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
	V3 gopurs_runtime.Value
}

//gopurs:init-v1:00670065006e006500720061007400650064002f0043006f006e00740072006f006c005f004d006f006e00610064005f005700720069007400650072005f0043006c006100730073:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f0043006f006e00740072006f006c005f004d006f006e00610064005f005700720069007400650072005f0043006c006100730073005f004d006f006e00610064005700720069007400650072
func init() {
	gopurs_runtime.StructGetters[784743459] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Control_Monad_Writer_Class_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "MonadTell1":
			return gopurs_runtime.Box(c.V0)
		case "Monoid0":
			return gopurs_runtime.Box(c.V1)
		case "listen":
			return gopurs_runtime.Box(c.V2)
		case "pass":
			return gopurs_runtime.Box(c.V3)
		default:
			panic("Key not found in dictionary Constructor_Control_Monad_Writer_Class_MonadWriter: " + key)
		}
	}
}

func Call_Control_Monad_Writer_Class_MonadTell_dollar_Dict(x_0_loop struct {
	Monad1     gopurs_runtime.Value
	Semigroup0 gopurs_runtime.Value
	tell       gopurs_runtime.Value
	/*gopurs:record-v1:3:004d006f006e006100640031,00530065006d006900670072006f007500700030,00740065006c006c*/
}) *Constructor_Control_Monad_Writer_Class_MonadTell[gopurs_runtime.Value, gopurs_runtime.Value] {
	var x_0 struct {
		Monad1     gopurs_runtime.Value
		Semigroup0 gopurs_runtime.Value
		tell       gopurs_runtime.Value
		/*gopurs:record-v1:3:004d006f006e006100640031,00530065006d006900670072006f007500700030,00740065006c006c*/
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Writer_Class_MonadTell[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict3("Monad1", "Semigroup0", "tell", orig.Monad1, orig.Semigroup0, orig.tell)
	}())
}

func Call_Control_Monad_Writer_Class_MonadWriter_dollar_Dict(x_0_loop struct {
	MonadTell1 gopurs_runtime.Value
	Monoid0    gopurs_runtime.Value
	listen     gopurs_runtime.Value
	pass       gopurs_runtime.Value
	/*gopurs:record-v1:4:004d006f006e0061006400540065006c006c0031,004d006f006e006f006900640030,006c0069007300740065006e,0070006100730073*/
}) *Constructor_Control_Monad_Writer_Class_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value] {
	var x_0 struct {
		MonadTell1 gopurs_runtime.Value
		Monoid0    gopurs_runtime.Value
		listen     gopurs_runtime.Value
		pass       gopurs_runtime.Value
		/*gopurs:record-v1:4:004d006f006e0061006400540065006c006c0031,004d006f006e006f006900640030,006c0069007300740065006e,0070006100730073*/
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Writer_Class_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict4("MonadTell1", "Monoid0", "listen", "pass", orig.MonadTell1, orig.Monoid0, orig.listen, orig.pass)
	}())
}

func Call_Control_Monad_Writer_Class_tell(dict_0_loop *Constructor_Control_Monad_Writer_Class_MonadTell[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Control_Monad_Writer_Class_MonadTell[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V2
}

func Call_Control_Monad_Writer_Class_pass(dict_0_loop *Constructor_Control_Monad_Writer_Class_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Control_Monad_Writer_Class_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V3
}

func Call_Control_Monad_Writer_Class_listen(dict_0_loop *Constructor_Control_Monad_Writer_Class_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Control_Monad_Writer_Class_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V2
}

func Call_Control_Monad_Writer_Class_listens(dictMonadWriter_0_loop *Constructor_Control_Monad_Writer_Class_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictMonadWriter_0 *Constructor_Control_Monad_Writer_Class_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value] = dictMonadWriter_0_loop
	_ = dictMonadWriter_0
	// TAST (Let): Monad1_1_0 shape=App(Other) bindingType=Any
	Monad1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictMonadWriter_0.V0, gopurs_runtime.Value{}), "Monad1"), gopurs_runtime.Value{})
	_ = Monad1_1_0
	// TAST (Let): pure_2_1 shape=App(Var) bindingType=(Func [(ADT ["Data","Tuple","Tuple"] [(TypeVar a$scope14), (TypeVar b$scope15)])] (TypeApp (TypeVar m$scope13) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a$scope14), (TypeVar b$scope15)])]))
	pure_2_1 := Call_Control_Applicative_pure(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Applicative0"), gopurs_runtime.Value{})))
	_ = pure_2_1
	return gopurs_runtime.Func2(func(f_3 gopurs_runtime.Value, m_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply(dictMonadWriter_0.V2, m_4), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(pure_2_1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 gopurs_runtime.Value
				}{(*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0, gopurs_runtime.Apply(f_3, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V1)}
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()))})
		}))
	})
}

func Call_Control_Monad_Writer_Class_censor(dictMonadWriter_0_loop *Constructor_Control_Monad_Writer_Class_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictMonadWriter_0 *Constructor_Control_Monad_Writer_Class_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value] = dictMonadWriter_0_loop
	_ = dictMonadWriter_0
	// TAST (Let): Monad1_1_0 shape=App(Other) bindingType=Any
	Monad1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictMonadWriter_0.V0, gopurs_runtime.Value{}), "Monad1"), gopurs_runtime.Value{})
	_ = Monad1_1_0
	// TAST (Let): pure_2_1 shape=App(Var) bindingType=(Func [(ADT ["Data","Tuple","Tuple"] [(TypeVar a$scope23), (Func [(TypeVar w$scope21)] (TypeVar w$scope21))])] (TypeApp (TypeVar m$scope22) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a$scope23), (Func [(TypeVar w$scope21)] (TypeVar w$scope21))])]))
	pure_2_1 := Call_Control_Applicative_pure(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Applicative0"), gopurs_runtime.Value{})))
	_ = pure_2_1
	return gopurs_runtime.Func2(func(f_3 gopurs_runtime.Value, m_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(dictMonadWriter_0.V3, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Bind1"), gopurs_runtime.Value{}), "bind"), m_4, gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(pure_2_1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 gopurs_runtime.Value
				}{a_5, f_3}
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()))})
		})))
	})
}
