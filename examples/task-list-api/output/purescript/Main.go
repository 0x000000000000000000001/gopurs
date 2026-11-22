package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_conj gopurs_runtime.Value
var once_Main_conj sync.Once

func Get_Main_conj() gopurs_runtime.Value {
	once_Main_conj.Do(func() {
		cache_Main_conj = Get_Data_HeytingAlgebra_conj__2421783737()
	})
	return cache_Main_conj
}

var cache_Main_notEq gopurs_runtime.Value
var once_Main_notEq sync.Once

func Get_Main_notEq() gopurs_runtime.Value {
	once_Main_notEq.Do(func() {
		cache_Main_notEq = Get_Data_Eq_notEq__3478410105()
	})
	return cache_Main_notEq
}

var cache_Main_lessThanOrEq gopurs_runtime.Value
var once_Main_lessThanOrEq sync.Once

func Get_Main_lessThanOrEq() gopurs_runtime.Value {
	once_Main_lessThanOrEq.Do(func() {
		cache_Main_lessThanOrEq = Get_Data_Ord_lessThanOrEq__3348776441()
	})
	return cache_Main_lessThanOrEq
}

var cache_Main_greaterThanOrEq gopurs_runtime.Value
var once_Main_greaterThanOrEq sync.Once

func Get_Main_greaterThanOrEq() gopurs_runtime.Value {
	once_Main_greaterThanOrEq.Do(func() {
		cache_Main_greaterThanOrEq = Get_Data_Ord_greaterThanOrEq__4238969209()
	})
	return cache_Main_greaterThanOrEq
}

var cache_Main_lessThanOrEq1 gopurs_runtime.Value
var once_Main_lessThanOrEq1 sync.Once

func Get_Main_lessThanOrEq1() gopurs_runtime.Value {
	once_Main_lessThanOrEq1.Do(func() {
		cache_Main_lessThanOrEq1 = Get_Data_Ord_lessThanOrEq__4238969209()
	})
	return cache_Main_lessThanOrEq1
}

var cache_Main_gEncodeJsonCons gopurs_runtime.Value
var once_Main_gEncodeJsonCons sync.Once

func Get_Main_gEncodeJsonCons() gopurs_runtime.Value {
	once_Main_gEncodeJsonCons.Do(func() {
		cache_Main_gEncodeJsonCons = Call_Data_Argonaut_Encode_Class_gEncodeJsonCons(gopurs_runtime.Value{Type: 9, IntVal: 3228574649, UnsafePtr: unsafe.Pointer(Rebox_Main_2086421970_2831479053(Rebox_Main_2831479053_2086421970(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Encode_Class_EncodeJson[gopurs_runtime.Value]](Get_Data_Argonaut_Encode_Class_encodeJsonJString()))))})
	})
	return cache_Main_gEncodeJsonCons
}

var cache_Main_gEncodeJsonCons1 gopurs_runtime.Value
var once_Main_gEncodeJsonCons1 sync.Once

func Get_Main_gEncodeJsonCons1() gopurs_runtime.Value {
	once_Main_gEncodeJsonCons1.Do(func() {
		cache_Main_gEncodeJsonCons1 = gopurs_runtime.Apply(Call_Data_Argonaut_Encode_Class_gEncodeJsonCons(gopurs_runtime.Value{Type: 9, IntVal: 3228574649, UnsafePtr: unsafe.Pointer(Rebox_Main_2086421970_2831479053(Rebox_Main_2831479053_2086421970(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Encode_Class_EncodeJson[gopurs_runtime.Value]](Get_Data_Argonaut_Encode_Class_encodeJsonJString()))))}), Get_Data_Argonaut_Encode_Class_gEncodeJsonNil())
	})
	return cache_Main_gEncodeJsonCons1
}

var cache_Main_bindExceptT gopurs_runtime.Value
var once_Main_bindExceptT sync.Once

func Get_Main_bindExceptT() gopurs_runtime.Value {
	once_Main_bindExceptT.Do(func() {
		cache_Main_bindExceptT = gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Call_Control_Monad_Except_Trans_bindExceptT(gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](Get_Data_Identity_monadIdentity()))})))}
	})
	return cache_Main_bindExceptT
}

var cache_Main_bind gopurs_runtime.Value
var once_Main_bind sync.Once

func Get_Main_bind() gopurs_runtime.Value {
	once_Main_bind.Do(func() {
		cache_Main_bind = gopurs_runtime.RecordGet(Call_Control_Monad_Except_Trans_bindExceptT(gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](Get_Data_Identity_monadIdentity()))}), "bind")
	})
	return cache_Main_bind
}

var cache_Main_readString gopurs_runtime.Value
var once_Main_readString sync.Once

func Get_Main_readString() gopurs_runtime.Value {
	once_Main_readString.Do(func() {
		cache_Main_readString = Call_Foreign_readString(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](Get_Data_Identity_monadIdentity()))
	})
	return cache_Main_readString
}

var cache_Main_bind1 gopurs_runtime.Value
var once_Main_bind1 sync.Once

func Get_Main_bind1() gopurs_runtime.Value {
	once_Main_bind1.Do(func() {
		cache_Main_bind1 = gopurs_runtime.RecordGet(Call_Control_Monad_Except_Trans_bindExceptT(gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](Get_Data_Identity_monadIdentity()))}), "bind")
	})
	return cache_Main_bind1
}

var cache_Main_go__map gopurs_runtime.Value
var once_Main_go__map sync.Once

func Get_Main_go__map() gopurs_runtime.Value {
	once_Main_go__map.Do(func() {
		cache_Main_go__map = gopurs_runtime.RecordGet(Call_Control_Monad_Except_Trans_functorExceptT(gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Identity_functorIdentity()))}), "map")
	})
	return cache_Main_go__map
}

var cache_Main_notEq1 gopurs_runtime.Value
var once_Main_notEq1 sync.Once

func Get_Main_notEq1() gopurs_runtime.Value {
	once_Main_notEq1.Do(func() {
		cache_Main_notEq1 = Get_Data_Eq_notEq__3020190905()
	})
	return cache_Main_notEq1
}

var cache_Main_readInt gopurs_runtime.Value
var once_Main_readInt sync.Once

func Get_Main_readInt() gopurs_runtime.Value {
	once_Main_readInt.Do(func() {
		cache_Main_readInt = Call_Foreign_readInt(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](Get_Data_Identity_monadIdentity()))
	})
	return cache_Main_readInt
}

var cache_Main_pure gopurs_runtime.Value
var once_Main_pure sync.Once

func Get_Main_pure() gopurs_runtime.Value {
	once_Main_pure.Do(func() {
		cache_Main_pure = gopurs_runtime.RecordGet(Call_Control_Monad_Except_Trans_applicativeExceptT(gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](Get_Data_Identity_monadIdentity()))}), "pure")
	})
	return cache_Main_pure
}

var cache_Main_throwError gopurs_runtime.Value
var once_Main_throwError sync.Once

func Get_Main_throwError() gopurs_runtime.Value {
	once_Main_throwError.Do(func() {
		cache_Main_throwError = Get_Effect_Aff__throwError()
	})
	return cache_Main_throwError
}

var cache_Main_go__append gopurs_runtime.Value
var once_Main_go__append sync.Once

func Get_Main_go__append() gopurs_runtime.Value {
	once_Main_go__append.Do(func() {
		cache_Main_go__append = Get_Data_Semigroup_append__1694605992()
	})
	return cache_Main_go__append
}

var cache_Main_show gopurs_runtime.Value
var once_Main_show sync.Once

func Get_Main_show() gopurs_runtime.Value {
	once_Main_show.Do(func() {
		cache_Main_show = Get_Data_Show_show__3342561716()
	})
	return cache_Main_show
}

var cache_Main_pure1 gopurs_runtime.Value
var once_Main_pure1 sync.Once

func Get_Main_pure1() gopurs_runtime.Value {
	once_Main_pure1.Do(func() {
		cache_Main_pure1 = Get_Control_Applicative_pure__3763014981()
	})
	return cache_Main_pure1
}

var cache_Main_bind2 gopurs_runtime.Value
var once_Main_bind2 sync.Once

func Get_Main_bind2() gopurs_runtime.Value {
	once_Main_bind2.Do(func() {
		cache_Main_bind2 = Get_Control_Bind_bind__2624874445()
	})
	return cache_Main_bind2
}

var cache_Main_pure2 gopurs_runtime.Value
var once_Main_pure2 sync.Once

func Get_Main_pure2() gopurs_runtime.Value {
	once_Main_pure2.Do(func() {
		cache_Main_pure2 = Get_Effect_Aff__pure()
	})
	return cache_Main_pure2
}

var cache_Main_map1 gopurs_runtime.Value
var once_Main_map1 sync.Once

func Get_Main_map1() gopurs_runtime.Value {
	once_Main_map1.Do(func() {
		cache_Main_map1 = Get_Data_Functor_map__2345995242()
	})
	return cache_Main_map1
}

var cache_Main_titleIsSymbol gopurs_runtime.Value
var once_Main_titleIsSymbol sync.Once

func Get_Main_titleIsSymbol() gopurs_runtime.Value {
	once_Main_titleIsSymbol.Do(func() {
		cache_Main_titleIsSymbol = gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.StrWithCodeUnits("title", 5)
		}))
	})
	return cache_Main_titleIsSymbol
}

var cache_Main_completedIsSymbol gopurs_runtime.Value
var once_Main_completedIsSymbol sync.Once

func Get_Main_completedIsSymbol() gopurs_runtime.Value {
	once_Main_completedIsSymbol.Do(func() {
		cache_Main_completedIsSymbol = gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.StrWithCodeUnits("completed", 9)
		}))
	})
	return cache_Main_completedIsSymbol
}

var cache_Main_encodeRecord gopurs_runtime.Value
var once_Main_encodeRecord sync.Once

func Get_Main_encodeRecord() gopurs_runtime.Value {
	once_Main_encodeRecord.Do(func() {
		cache_Main_encodeRecord = Call_Data_Argonaut_Encode_Class_encodeRecord(gopurs_runtime.Apply3(Call_Data_Argonaut_Encode_Class_gEncodeJsonCons(gopurs_runtime.Value{Type: 9, IntVal: 3228574649, UnsafePtr: unsafe.Pointer(Rebox_Main_946498665_2831479053(Rebox_Main_2831479053_946498665(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Encode_Class_EncodeJson[gopurs_runtime.Value]](Get_Data_Argonaut_Encode_Class_encodeJsonJBoolean()))))}), gopurs_runtime.Apply3(Call_Data_Argonaut_Encode_Class_gEncodeJsonCons(gopurs_runtime.Value{Type: 9, IntVal: 3228574649, UnsafePtr: unsafe.Pointer(Rebox_Main_2086421970_2831479053(Rebox_Main_2831479053_2086421970(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Encode_Class_EncodeJson[gopurs_runtime.Value]](Get_Data_Argonaut_Encode_Class_encodeJsonJString()))))}), gopurs_runtime.Apply3(Call_Data_Argonaut_Encode_Class_gEncodeJsonCons(gopurs_runtime.Value{Type: 9, IntVal: 3228574649, UnsafePtr: unsafe.Pointer(Rebox_Main_2086421970_2831479053(Rebox_Main_2831479053_2086421970(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Encode_Class_EncodeJson[gopurs_runtime.Value]](Get_Data_Argonaut_Encode_Class_encodeJsonJString()))))}), Get_Data_Argonaut_Encode_Class_gEncodeJsonNil(), Get_Main_titleIsSymbol(), gopurs_runtime.Value{}), gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.StrWithCodeUnits("id", 2)
		})), gopurs_runtime.Value{}), Get_Main_completedIsSymbol(), gopurs_runtime.Value{}))
	})
	return cache_Main_encodeRecord
}

var cache_Main_bind3 gopurs_runtime.Value
var once_Main_bind3 sync.Once

func Get_Main_bind3() gopurs_runtime.Value {
	once_Main_bind3.Do(func() {
		cache_Main_bind3 = Get_Control_Bind_bind__1318103565()
	})
	return cache_Main_bind3
}

var cache_Main_traverse gopurs_runtime.Value
var once_Main_traverse sync.Once

func Get_Main_traverse() gopurs_runtime.Value {
	once_Main_traverse.Do(func() {
		cache_Main_traverse = Get_Data_Traversable_traverse__2498349992()
	})
	return cache_Main_traverse
}

var cache_Main_map2 gopurs_runtime.Value
var once_Main_map2 sync.Once

func Get_Main_map2() gopurs_runtime.Value {
	once_Main_map2.Do(func() {
		cache_Main_map2 = Get_Data_Functor_map__1317168042()
	})
	return cache_Main_map2
}

var cache_Main_fromJsonString gopurs_runtime.Value
var once_Main_fromJsonString sync.Once

func Get_Main_fromJsonString() gopurs_runtime.Value {
	once_Main_fromJsonString.Do(func() {
		cache_Main_fromJsonString = Get_Data_Argonaut_Decode_fromJsonString__2215306652()
	})
	return cache_Main_fromJsonString
}

var cache_Main_disj gopurs_runtime.Value
var once_Main_disj sync.Once

func Get_Main_disj() gopurs_runtime.Value {
	once_Main_disj.Do(func() {
		cache_Main_disj = Get_Data_HeytingAlgebra_disj__2421783737()
	})
	return cache_Main_disj
}

var cache_Main_eq gopurs_runtime.Value
var once_Main_eq sync.Once

func Get_Main_eq() gopurs_runtime.Value {
	once_Main_eq.Do(func() {
		cache_Main_eq = Get_Data_Eq_eq__3478410105()
	})
	return cache_Main_eq
}

var cache_Main_greaterThan gopurs_runtime.Value
var once_Main_greaterThan sync.Once

func Get_Main_greaterThan() gopurs_runtime.Value {
	once_Main_greaterThan.Do(func() {
		cache_Main_greaterThan = Get_Data_Ord_greaterThan__3348776441()
	})
	return cache_Main_greaterThan
}

var cache_Main_bind4 gopurs_runtime.Value
var once_Main_bind4 sync.Once

func Get_Main_bind4() gopurs_runtime.Value {
	once_Main_bind4.Do(func() {
		cache_Main_bind4 = Get_Control_Bind_bind__156704557()
	})
	return cache_Main_bind4
}

var cache_Main_fromJsonString1 gopurs_runtime.Value
var once_Main_fromJsonString1 sync.Once

func Get_Main_fromJsonString1() gopurs_runtime.Value {
	once_Main_fromJsonString1.Do(func() {
		cache_Main_fromJsonString1 = Get_Data_Argonaut_Decode_fromJsonString__2316388992()
	})
	return cache_Main_fromJsonString1
}

var cache_Main_eq1 gopurs_runtime.Value
var once_Main_eq1 sync.Once

func Get_Main_eq1() gopurs_runtime.Value {
	once_Main_eq1.Do(func() {
		cache_Main_eq1 = Get_Data_Eq_eq__3020190905()
	})
	return cache_Main_eq1
}

var cache_Main_discard gopurs_runtime.Value
var once_Main_discard sync.Once

func Get_Main_discard() gopurs_runtime.Value {
	once_Main_discard.Do(func() {
		cache_Main_discard = Get_Control_Bind_discard__143037513()
	})
	return cache_Main_discard
}

var cache_Main_applySecond gopurs_runtime.Value
var once_Main_applySecond sync.Once

func Get_Main_applySecond() gopurs_runtime.Value {
	once_Main_applySecond.Do(func() {
		cache_Main_applySecond = Get_Control_Apply_applySecond__3021785433()
	})
	return cache_Main_applySecond
}

var cache_Main_pure3 gopurs_runtime.Value
var once_Main_pure3 sync.Once

func Get_Main_pure3() gopurs_runtime.Value {
	once_Main_pure3.Do(func() {
		cache_Main_pure3 = Get_Control_Applicative_pure__2088919292()
	})
	return cache_Main_pure3
}

var cache_Main_bind5 gopurs_runtime.Value
var once_Main_bind5 sync.Once

func Get_Main_bind5() gopurs_runtime.Value {
	once_Main_bind5.Do(func() {
		cache_Main_bind5 = Get_Control_Bind_bind__4145807597()
	})
	return cache_Main_bind5
}

var cache_Main_liftEffect gopurs_runtime.Value
var once_Main_liftEffect sync.Once

func Get_Main_liftEffect() gopurs_runtime.Value {
	once_Main_liftEffect.Do(func() {
		cache_Main_liftEffect = Get_Effect_Aff__liftEffect()
	})
	return cache_Main_liftEffect
}

var cache_Main_map3 gopurs_runtime.Value
var once_Main_map3 sync.Once

func Get_Main_map3() gopurs_runtime.Value {
	once_Main_map3.Do(func() {
		cache_Main_map3 = Get_Data_Functor_map__3654485947()
	})
	return cache_Main_map3
}

var cache_Main_discard1 gopurs_runtime.Value
var once_Main_discard1 sync.Once

func Get_Main_discard1() gopurs_runtime.Value {
	once_Main_discard1.Do(func() {
		cache_Main_discard1 = Get_Control_Bind_discard__4273756336()
	})
	return cache_Main_discard1
}

var cache_Main_void gopurs_runtime.Value
var once_Main_void sync.Once

func Get_Main_void() gopurs_runtime.Value {
	once_Main_void.Do(func() {
		cache_Main_void = Call_Data_Functor_void(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_Aff_functorAff()))
	})
	return cache_Main_void
}

var cache_Main_liftEffect1 gopurs_runtime.Value
var once_Main_liftEffect1 sync.Once

func Get_Main_liftEffect1() gopurs_runtime.Value {
	once_Main_liftEffect1.Do(func() {
		cache_Main_liftEffect1 = Get_Effect_Aff__liftEffect()
	})
	return cache_Main_liftEffect1
}

var cache_Main_validId gopurs_runtime.Value
var once_Main_validId sync.Once

func Get_Main_validId() gopurs_runtime.Value {
	once_Main_validId.Do(func() {
		cache_Main_validId = gopurs_runtime.Func(func(value_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(Call_Main_validId(value_0_box.StrVal()))
		})
	})
	return cache_Main_validId
}

var cache_Main_selectTasks gopurs_runtime.Value
var once_Main_selectTasks sync.Once

func Get_Main_selectTasks() gopurs_runtime.Value {
	once_Main_selectTasks.Do(func() {
		cache_Main_selectTasks = gopurs_runtime.StrWithCodeUnits("SELECT CAST(id AS TEXT), title, completed FROM tasks", 52)
	})
	return cache_Main_selectTasks
}

var cache_Main_json gopurs_runtime.Value
var once_Main_json sync.Once

func Get_Main_json() gopurs_runtime.Value {
	once_Main_json.Do(func() {
		cache_Main_json = gopurs_runtime.Func(func(dictEncodeJson_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_json(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Encode_Class_EncodeJson[gopurs_runtime.Value]](dictEncodeJson_0_box))
		})
	})
	return cache_Main_json
}

var cache_Main_json__4250750939 gopurs_runtime.Value
var once_Main_json__4250750939 sync.Once

func Get_Main_json__4250750939() gopurs_runtime.Value {
	once_Main_json__4250750939.Do(func() {
		cache_Main_json__4250750939 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				orig := Call_Main_json__4250750939(__eta0_0_box.IntVal, func() []struct {
					completed bool
					id        string
					title     string
					/*gopurs:record-v1:3:0063006f006d0070006c0065007400650064,00690064,007400690074006c0065*/
				} {
					arr := *(*[]gopurs_runtime.Value)(__eta1_1_box.UnsafePtr)
					unboxed := make([]struct {
						completed bool
						id        string
						title     string
						/*gopurs:record-v1:3:0063006f006d0070006c0065007400650064,00690064,007400690074006c0065*/
					}, len(arr))
					for i, v := range arr {
						unboxed[i] = func() struct {
							completed bool
							id        string
							title     string
							/*gopurs:record-v1:3:0063006f006d0070006c0065007400650064,00690064,007400690074006c0065*/
						} {
							orig := v
							_ = orig
							clone := struct {
								completed bool
								id        string
								title     string
								/*gopurs:record-v1:3:0063006f006d0070006c0065007400650064,00690064,007400690074006c0065*/
							}{}
							clone.completed = (gopurs_runtime.RecordGet(orig, "completed").IntVal) != (0)
							clone.id = gopurs_runtime.RecordGet(orig, "id").StrVal()
							clone.title = gopurs_runtime.RecordGet(orig, "title").StrVal()
							return clone
						}()
					}
					return unboxed
				}())
				_ = orig
				return gopurs_runtime.RecordDict2("body", "status", gopurs_runtime.Str(orig.body), gopurs_runtime.Int(orig.status))
			}()
		})
	})
	return cache_Main_json__4250750939
}

var cache_Main_json__1164587298 gopurs_runtime.Value
var once_Main_json__1164587298 sync.Once

func Get_Main_json__1164587298() gopurs_runtime.Value {
	once_Main_json__1164587298.Do(func() {
		cache_Main_json__1164587298 = gopurs_runtime.Func(func(dictEncodeJson_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_json__1164587298(Rebox_Main_2831479053_614949298(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Encode_Class_EncodeJson[gopurs_runtime.Value]](dictEncodeJson_0_box)))
		})
	})
	return cache_Main_json__1164587298
}

var cache_Main_json__4251671259 gopurs_runtime.Value
var once_Main_json__4251671259 sync.Once

func Get_Main_json__4251671259() gopurs_runtime.Value {
	once_Main_json__4251671259.Do(func() {
		cache_Main_json__4251671259 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				orig := Call_Main_json__4251671259(__eta0_0_box.IntVal, func() struct {
					completed bool
					id        string
					title     string
					/*gopurs:record-v1:3:0063006f006d0070006c0065007400650064,00690064,007400690074006c0065*/
				} {
					orig := __eta1_1_box
					_ = orig
					clone := struct {
						completed bool
						id        string
						title     string
						/*gopurs:record-v1:3:0063006f006d0070006c0065007400650064,00690064,007400690074006c0065*/
					}{}
					clone.completed = (gopurs_runtime.RecordGet(orig, "completed").IntVal) != (0)
					clone.id = gopurs_runtime.RecordGet(orig, "id").StrVal()
					clone.title = gopurs_runtime.RecordGet(orig, "title").StrVal()
					return clone
				}())
				_ = orig
				return gopurs_runtime.RecordDict2("body", "status", gopurs_runtime.Str(orig.body), gopurs_runtime.Int(orig.status))
			}()
		})
	})
	return cache_Main_json__4251671259
}

var cache_Main_json__3478233858 gopurs_runtime.Value
var once_Main_json__3478233858 sync.Once

func Get_Main_json__3478233858() gopurs_runtime.Value {
	once_Main_json__3478233858.Do(func() {
		cache_Main_json__3478233858 = gopurs_runtime.Func(func(dictEncodeJson_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_json__3478233858(Rebox_Main_2831479053_1331040862(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Encode_Class_EncodeJson[gopurs_runtime.Value]](dictEncodeJson_0_box)))
		})
	})
	return cache_Main_json__3478233858
}

var cache_Main_json1 gopurs_runtime.Value
var once_Main_json1 sync.Once

func Get_Main_json1() gopurs_runtime.Value {
	once_Main_json1.Do(func() {
		cache_Main_json1 = func() gopurs_runtime.Value {
			// TAST (Let): toJsonString__913691249_0_0 shape=App(Var) bindingType=(Func [(Record (Row [error: String] Empty))] String)
			toJsonString__913691249_0_0 := gopurs_runtime.Apply2(Call_Control_Semigroupoid_composeFlipped(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Call_Data_Argonaut_Encode_Class_encodeJson(Rebox_Main_614949298_2831479053(Rebox_Main_2831479053_614949298(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Encode_Class_EncodeJson[gopurs_runtime.Value]](Get_Main___cached_dict_0())))), Get_Data_Argonaut_Core_stringify())
			_ = toJsonString__913691249_0_0
			return gopurs_runtime.Func2(func(status_1 gopurs_runtime.Value, value_2 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.RecordDict2("body", "status", gopurs_runtime.Apply(toJsonString__913691249_0_0, value_2), status_1)
			})
		}()
	})
	return cache_Main_json1
}

var cache_Main_json2 gopurs_runtime.Value
var once_Main_json2 sync.Once

func Get_Main_json2() gopurs_runtime.Value {
	once_Main_json2.Do(func() {
		cache_Main_json2 = Get_Main_json__4251671259()
	})
	return cache_Main_json2
}

var cache_Main_json3 gopurs_runtime.Value
var once_Main_json3 sync.Once

func Get_Main_json3() gopurs_runtime.Value {
	once_Main_json3.Do(func() {
		cache_Main_json3 = func() gopurs_runtime.Value {
			// TAST (Let): toJsonString__2805964221_0_0 shape=App(Var) bindingType=(Func [(Record (Row [status: String] Empty))] String)
			toJsonString__2805964221_0_0 := gopurs_runtime.Apply2(Call_Control_Semigroupoid_composeFlipped(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Call_Data_Argonaut_Encode_Class_encodeJson(Rebox_Main_1331040862_2831479053(Rebox_Main_2831479053_1331040862(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Encode_Class_EncodeJson[gopurs_runtime.Value]](Get_Main___cached_dict_1())))), Get_Data_Argonaut_Core_stringify())
			_ = toJsonString__2805964221_0_0
			return gopurs_runtime.Func2(func(status_1 gopurs_runtime.Value, value_2 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.RecordDict2("body", "status", gopurs_runtime.Apply(toJsonString__2805964221_0_0, value_2), status_1)
			})
		}()
	})
	return cache_Main_json3
}

var cache_Main_json4 gopurs_runtime.Value
var once_Main_json4 sync.Once

func Get_Main_json4() gopurs_runtime.Value {
	once_Main_json4.Do(func() {
		cache_Main_json4 = Get_Main_json__4250750939()
	})
	return cache_Main_json4
}

var cache_Main_failure gopurs_runtime.Value
var once_Main_failure sync.Once

func Get_Main_failure() gopurs_runtime.Value {
	once_Main_failure.Do(func() {
		cache_Main_failure = gopurs_runtime.Func2(func(status_0_box gopurs_runtime.Value, text_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				orig := Call_Main_failure(status_0_box.IntVal, text_1_box.StrVal())
				_ = orig
				return gopurs_runtime.RecordDict2("body", "status", gopurs_runtime.Str(orig.body), gopurs_runtime.Int(orig.status))
			}()
		})
	})
	return cache_Main_failure
}

var cache_Main_decodeTask gopurs_runtime.Value
var once_Main_decodeTask sync.Once

func Get_Main_decodeTask() gopurs_runtime.Value {
	once_Main_decodeTask.Do(func() {
		cache_Main_decodeTask = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_decodeTask(*((*[]gopurs_runtime.Value)(v_0_box.UnsafePtr)))
		})
	})
	return cache_Main_decodeTask
}

var cache_Main_getTask gopurs_runtime.Value
var once_Main_getTask sync.Once

func Get_Main_getTask() gopurs_runtime.Value {
	once_Main_getTask.Do(func() {
		cache_Main_getTask = gopurs_runtime.Func2(func(database_0_box gopurs_runtime.Value, id_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_getTask(database_0_box, id_1_box.StrVal())
		})
	})
	return cache_Main_getTask
}

var cache_Main_tasks gopurs_runtime.Value
var once_Main_tasks sync.Once

func Get_Main_tasks() gopurs_runtime.Value {
	once_Main_tasks.Do(func() {
		cache_Main_tasks = gopurs_runtime.Func(func(database_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_tasks(database_0_box)
		})
	})
	return cache_Main_tasks
}

var cache_Main_route gopurs_runtime.Value
var once_Main_route sync.Once

func Get_Main_route() gopurs_runtime.Value {
	once_Main_route.Do(func() {
		cache_Main_route = gopurs_runtime.Func2(func(database_0_box gopurs_runtime.Value, request_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_route(database_0_box, func() struct {
				body   string
				method string
				path   string
				/*gopurs:record-v1:3:0062006f00640079,006d006500740068006f0064,0070006100740068*/
			} {
				orig := request_1_box
				_ = orig
				clone := struct {
					body   string
					method string
					path   string
					/*gopurs:record-v1:3:0062006f00640079,006d006500740068006f0064,0070006100740068*/
				}{}
				clone.body = gopurs_runtime.RecordGet(orig, "body").StrVal()
				clone.method = gopurs_runtime.RecordGet(orig, "method").StrVal()
				clone.path = gopurs_runtime.RecordGet(orig, "path").StrVal()
				return clone
			}())
		})
	})
	return cache_Main_route
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Call_Data_Functor_void(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect())), Call_Effect_Aff_runAff__2939864241(gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t0 gopurs_runtime.Value
			{
				if v_0.Type == 9 && v_0.IntVal == 3711209382 {
					__t0 = Call_Control_Apply_applySecond__3021785433(gopurs_runtime.Apply(Get_Effect_Console_error(), gopurs_runtime.NormalizeStringValue(gopurs_runtime.Apply(Get_Effect_Exception_message(), (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V0))), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.UncurriedApp(Get_Node_Process_exitImpl(), gopurs_runtime.Int(int64(1)))
					}))
					goto end_branch_0
				} else {

				}
			}
			{
				if v_0.Type == 9 && v_0.IntVal == 2465973597 {
					__t0 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
						return Get_Data_Unit_unit()
					})
					goto end_branch_0
				} else {

				}
			}
			{
				__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_0:
			return __t0
		}), gopurs_runtime.Apply2(Get_Effect_Aff__bind(), gopurs_runtime.Apply(Get_Effect_Aff__liftEffect(), Call_Control_Applicative_liftA1(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Effect_applicativeEffect()), gopurs_runtime.Func(func(v2_0 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t1 gopurs_runtime.Value
			{
				if v2_0.Type == 9 && v2_0.IntVal == 930809136 && v2_0.UnsafePtr == nil {
					__t1 = gopurs_runtime.StrWithCodeUnits("127.0.0.1:8080", 14)
					goto end_branch_1
				} else {

				}
			}
			{
				if v2_0.Type == 9 && v2_0.IntVal == 930809136 && v2_0.UnsafePtr != nil {
					__t1 = (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v2_0.UnsafePtr).V0
					goto end_branch_1
				} else {

				}
			}
			{
				__t1 = gopurs_runtime.Str(func() string { panic("Failed pattern match") }())
			}
		end_branch_1:
			return __t1
		}), Call_Control_Applicative_liftA1(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Effect_applicativeEffect()), gopurs_runtime.Apply(Get_Foreign_Object_lookup__2830175038(), gopurs_runtime.StrWithCodeUnits("TASK_LIST_ADDRESS", 17)), Get_Node_Process_unsafeGetEnv()))), gopurs_runtime.Func(func(address_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(Get_Effect_Aff__bind(), gopurs_runtime.Apply(Get_Effect_Aff__liftEffect(), Call_Control_Applicative_liftA1(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Effect_applicativeEffect()), gopurs_runtime.Func(func(v2_1 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t2 gopurs_runtime.Value
				{
					if v2_1.Type == 9 && v2_1.IntVal == 930809136 && v2_1.UnsafePtr == nil {
						__t2 = gopurs_runtime.StrWithCodeUnits("tasks.db", 8)
						goto end_branch_2
					} else {

					}
				}
				{
					if v2_1.Type == 9 && v2_1.IntVal == 930809136 && v2_1.UnsafePtr != nil {
						__t2 = (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v2_1.UnsafePtr).V0
						goto end_branch_2
					} else {

					}
				}
				{
					__t2 = gopurs_runtime.Str(func() string { panic("Failed pattern match") }())
				}
			end_branch_2:
				return __t2
			}), Call_Control_Applicative_liftA1(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Effect_applicativeEffect()), gopurs_runtime.Apply(Get_Foreign_Object_lookup__2830175038(), gopurs_runtime.StrWithCodeUnits("TASK_LIST_DATABASE", 18)), Get_Node_Process_unsafeGetEnv()))), gopurs_runtime.Func(func(path_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(Get_Effect_Aff__bind(), Call_Effect_Aff_bracket__939099628(gopurs_runtime.Apply(Get_Database_SQLite3_open(), gopurs_runtime.NormalizeStringValue(path_1)), Get_Database_SQLite3_go__close(), gopurs_runtime.Func(func(database_2 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply2(Get_Effect_Aff__bind(), gopurs_runtime.Apply(Call_Data_Functor_void(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_Aff_functorAff())), Call_Effect_Aff_Compat_fromEffectFnAff__560502078(gopurs_runtime.Apply3(Get_Database_SQLite3_executeImpl(), database_2, gopurs_runtime.StrWithCodeUnits("CREATE TABLE IF NOT EXISTS tasks (id INTEGER PRIMARY KEY AUTOINCREMENT, title TEXT NOT NULL, completed INTEGER NOT NULL DEFAULT 0 CHECK (completed IN (0, 1)))", 158), gopurs_runtime.Array([]gopurs_runtime.Value{})))), gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
						return Call_Example_HTTP_serve(address_0.StrVal(), gopurs_runtime.Func2(func(request_4 gopurs_runtime.Value, reply_5 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Apply(Call_Data_Functor_void(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect())), Call_Effect_Aff_runAff__2603932529(gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
								var __t5 gopurs_runtime.Value
								{
									if v1_6.Type == 9 && v1_6.IntVal == 2465973597 {
										__t5 = gopurs_runtime.Apply(reply_5, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v1_6.UnsafePtr).V0)
										goto end_branch_5
									} else {

									}
								}
								{
									if v1_6.Type == 9 && v1_6.IntVal == 3711209382 {
										__t5 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
											// TAST (Let): __local_var_7_3 shape=App(Var) bindingType=Any
											__local_var_7_3 := gopurs_runtime.Apply(Get_Effect_Console_error(), gopurs_runtime.NormalizeStringValue(gopurs_runtime.Apply(Get_Effect_Exception_message(), (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v1_6.UnsafePtr).V0)))
											_ = __local_var_7_3
											__local_var_8_4 := gopurs_runtime.Apply(__local_var_7_3, gopurs_runtime.Value{})
											_ = __local_var_8_4
											return gopurs_runtime.Apply2(reply_5, func() gopurs_runtime.Value {
												orig := Call_Main_failure(int64(500), "Internal server error")
												_ = orig
												return gopurs_runtime.RecordDict2("body", "status", gopurs_runtime.Str(orig.body), gopurs_runtime.Int(orig.status))
											}(), gopurs_runtime.Value{})
										})
										goto end_branch_5
									} else {

									}
								}
								{
									__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
								}
							end_branch_5:
								return __t5
							}), Call_Main_route(database_2, func() struct {
								body   string
								method string
								path   string
								/*gopurs:record-v1:3:0062006f00640079,006d006500740068006f0064,0070006100740068*/
							} {
								orig := request_4
								_ = orig
								clone := struct {
									body   string
									method string
									path   string
									/*gopurs:record-v1:3:0062006f00640079,006d006500740068006f0064,0070006100740068*/
								}{}
								clone.body = gopurs_runtime.RecordGet(orig, "body").StrVal()
								clone.method = gopurs_runtime.RecordGet(orig, "method").StrVal()
								clone.path = gopurs_runtime.RecordGet(orig, "path").StrVal()
								return clone
							}())))
						}))
					}))
				})), gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply(Get_Effect_Aff__liftEffect(), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.StrWithCodeUnits("STOPPED", 7)))
				}))
			}))
		}))))
	})
	return cache_Main_main
}

var cache_Main___cached_dict_0 gopurs_runtime.Value
var once_Main___cached_dict_0 sync.Once

func Get_Main___cached_dict_0() gopurs_runtime.Value {
	once_Main___cached_dict_0.Do(func() {
		cache_Main___cached_dict_0 = gopurs_runtime.Value{Type: 9, IntVal: 3228574649, UnsafePtr: unsafe.Pointer(Rebox_Main_614949298_2831479053(Rebox_Main_2831479053_614949298(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Encode_Class_EncodeJson[gopurs_runtime.Value]](gopurs_runtime.Apply(Call_Data_Argonaut_Encode_Class_encodeRecord(gopurs_runtime.Apply3(Call_Data_Argonaut_Encode_Class_gEncodeJsonCons(gopurs_runtime.Value{Type: 9, IntVal: 3228574649, UnsafePtr: unsafe.Pointer(Rebox_Main_2086421970_2831479053(Rebox_Main_2831479053_2086421970(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Encode_Class_EncodeJson[gopurs_runtime.Value]](Get_Data_Argonaut_Encode_Class_encodeJsonJString()))))}), Get_Data_Argonaut_Encode_Class_gEncodeJsonNil(), gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.StrWithCodeUnits("error", 5)
		})), gopurs_runtime.Value{})), gopurs_runtime.Value{})))))}
	})
	return cache_Main___cached_dict_0
}

var cache_Main___cached_dict_1 gopurs_runtime.Value
var once_Main___cached_dict_1 sync.Once

func Get_Main___cached_dict_1() gopurs_runtime.Value {
	once_Main___cached_dict_1.Do(func() {
		cache_Main___cached_dict_1 = gopurs_runtime.Value{Type: 9, IntVal: 3228574649, UnsafePtr: unsafe.Pointer(Rebox_Main_1331040862_2831479053(Rebox_Main_2831479053_1331040862(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Encode_Class_EncodeJson[gopurs_runtime.Value]](gopurs_runtime.Apply(Call_Data_Argonaut_Encode_Class_encodeRecord(gopurs_runtime.Apply3(Call_Data_Argonaut_Encode_Class_gEncodeJsonCons(gopurs_runtime.Value{Type: 9, IntVal: 3228574649, UnsafePtr: unsafe.Pointer(Rebox_Main_2086421970_2831479053(Rebox_Main_2831479053_2086421970(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Encode_Class_EncodeJson[gopurs_runtime.Value]](Get_Data_Argonaut_Encode_Class_encodeJsonJString()))))}), Get_Data_Argonaut_Encode_Class_gEncodeJsonNil(), gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.StrWithCodeUnits("status", 6)
		})), gopurs_runtime.Value{})), gopurs_runtime.Value{})))))}
	})
	return cache_Main___cached_dict_1
}

var cache_Main___cached_dict_2 gopurs_runtime.Value
var once_Main___cached_dict_2 sync.Once

func Get_Main___cached_dict_2() gopurs_runtime.Value {
	once_Main___cached_dict_2.Do(func() {
		cache_Main___cached_dict_2 = gopurs_runtime.Value{Type: 9, IntVal: 3228574649, UnsafePtr: unsafe.Pointer(Rebox_Main_614949298_2831479053(Rebox_Main_2831479053_614949298(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Encode_Class_EncodeJson[gopurs_runtime.Value]](gopurs_runtime.Apply(Call_Data_Argonaut_Encode_Class_encodeRecord(gopurs_runtime.Apply3(Call_Data_Argonaut_Encode_Class_gEncodeJsonCons(gopurs_runtime.Value{Type: 9, IntVal: 3228574649, UnsafePtr: unsafe.Pointer(Rebox_Main_2086421970_2831479053(Rebox_Main_2831479053_2086421970(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Encode_Class_EncodeJson[gopurs_runtime.Value]](Get_Data_Argonaut_Encode_Class_encodeJsonJString()))))}), Get_Data_Argonaut_Encode_Class_gEncodeJsonNil(), gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.StrWithCodeUnits("error", 5)
		})), gopurs_runtime.Value{})), gopurs_runtime.Value{})))))}
	})
	return cache_Main___cached_dict_2
}

var cache_Main___cached_dict_3 gopurs_runtime.Value
var once_Main___cached_dict_3 sync.Once

func Get_Main___cached_dict_3() gopurs_runtime.Value {
	once_Main___cached_dict_3.Do(func() {
		cache_Main___cached_dict_3 = gopurs_runtime.Value{Type: 9, IntVal: 3228574649, UnsafePtr: unsafe.Pointer(Rebox_Main_1331040862_2831479053(Rebox_Main_2831479053_1331040862(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Encode_Class_EncodeJson[gopurs_runtime.Value]](gopurs_runtime.Apply(Call_Data_Argonaut_Encode_Class_encodeRecord(gopurs_runtime.Apply3(Call_Data_Argonaut_Encode_Class_gEncodeJsonCons(gopurs_runtime.Value{Type: 9, IntVal: 3228574649, UnsafePtr: unsafe.Pointer(Rebox_Main_2086421970_2831479053(Rebox_Main_2831479053_2086421970(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Encode_Class_EncodeJson[gopurs_runtime.Value]](Get_Data_Argonaut_Encode_Class_encodeJsonJString()))))}), Get_Data_Argonaut_Encode_Class_gEncodeJsonNil(), gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.StrWithCodeUnits("status", 6)
		})), gopurs_runtime.Value{})), gopurs_runtime.Value{})))))}
	})
	return cache_Main___cached_dict_3
}

func Call_Main_validId(value_0_loop string) bool {
	var value_0 string = value_0_loop
	_ = value_0
	return (((value_0) == ("")) != (true)) && (((gopurs_runtime.Apply(Get_Data_String_CodeUnits_length(), gopurs_runtime.Str(value_0)).IntVal) <= (int64(19))) && ((gopurs_runtime.UncurriedApp2(Get_Data_Array_allImpl(), gopurs_runtime.Func(func(char_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Bool(((char_1.StrVal()) >= ("0")) && ((char_1.StrVal()) <= ("9")))
	}), func() gopurs_runtime.Value {
		arr := Data_String_CodeUnits_ToCharArray(value_0)
		boxed := make([]gopurs_runtime.Value, len(arr))
		for i, v := range arr {
			boxed[i] = gopurs_runtime.Str(v)
		}
		return gopurs_runtime.Array(boxed)
	}()).IntVal) != (0)))
}

func Call_Main_json(dictEncodeJson_0_loop *Constructor_Data_Argonaut_Encode_Class_EncodeJson[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictEncodeJson_0 *Constructor_Data_Argonaut_Encode_Class_EncodeJson[gopurs_runtime.Value] = dictEncodeJson_0_loop
	_ = dictEncodeJson_0
	// TAST (Let): toJsonString_1_0 shape=App(Var) bindingType=(Func [(TypeVar a$scope1)] String)
	toJsonString_1_0 := gopurs_runtime.Apply2(Call_Control_Semigroupoid_composeFlipped(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Call_Data_Argonaut_Encode_Class_encodeJson(dictEncodeJson_0), Get_Data_Argonaut_Core_stringify())
	_ = toJsonString_1_0
	return gopurs_runtime.Func2(func(status_2 gopurs_runtime.Value, value_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.RecordDict2("body", "status", gopurs_runtime.Apply(toJsonString_1_0, value_3), status_2)
	})
}

func Call_Main_json__4250750939(__eta0_0_loop int64, __eta1_1_loop []struct {
	completed bool
	id        string
	title     string
	/*gopurs:record-v1:3:0063006f006d0070006c0065007400650064,00690064,007400690074006c0065*/
}) struct {
	body   string
	status int64
	/*gopurs:record-v1:2:0062006f00640079,007300740061007400750073*/
} {
json__4250750939:
	for {
		if false {
			continue json__4250750939
		}
		var __eta0_0 int64 = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 []struct {
			completed bool
			id        string
			title     string
			/*gopurs:record-v1:3:0063006f006d0070006c0065007400650064,00690064,007400690074006c0065*/
		} = __eta1_1_loop
		_ = __eta1_1
		return struct {
			body   string
			status int64
			/*gopurs:record-v1:2:0062006f00640079,007300740061007400750073*/
		}{Call_Data_Argonaut_Encode_toJsonString__2078790066(__eta1_1), __eta0_0}
	}
}

func Call_Main_json__1164587298(dictEncodeJson_0_loop *Constructor_Data_Argonaut_Encode_Class_EncodeJson[struct {
	error string
	/*gopurs:record-v1:1:006500720072006f0072*/
}]) gopurs_runtime.Value {
json__1164587298:
	for {
		if false {
			continue json__1164587298
		}
		var dictEncodeJson_0 *Constructor_Data_Argonaut_Encode_Class_EncodeJson[struct {
			error string
			/*gopurs:record-v1:1:006500720072006f0072*/
		}] = dictEncodeJson_0_loop
		_ = dictEncodeJson_0
		// TAST (Let): toJsonString__913691249_1_0 shape=App(Var) bindingType=(Func [(Record (Row [error: String] Empty))] String)
		toJsonString__913691249_1_0 := gopurs_runtime.Apply2(Call_Control_Semigroupoid_composeFlipped(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Call_Data_Argonaut_Encode_Class_encodeJson(Rebox_Main_614949298_2831479053(dictEncodeJson_0)), Get_Data_Argonaut_Core_stringify())
		_ = toJsonString__913691249_1_0
		return gopurs_runtime.Func2(func(status_2 gopurs_runtime.Value, value_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.RecordDict2("body", "status", gopurs_runtime.Apply(toJsonString__913691249_1_0, value_3), status_2)
		})
	}
}

func Call_Main_json__4251671259(__eta0_0_loop int64, __eta1_1_loop struct {
	completed bool
	id        string
	title     string
	/*gopurs:record-v1:3:0063006f006d0070006c0065007400650064,00690064,007400690074006c0065*/
}) struct {
	body   string
	status int64
	/*gopurs:record-v1:2:0062006f00640079,007300740061007400750073*/
} {
json__4251671259:
	for {
		if false {
			continue json__4251671259
		}
		var __eta0_0 int64 = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 struct {
			completed bool
			id        string
			title     string
			/*gopurs:record-v1:3:0063006f006d0070006c0065007400650064,00690064,007400690074006c0065*/
		} = __eta1_1_loop
		_ = __eta1_1
		return struct {
			body   string
			status int64
			/*gopurs:record-v1:2:0062006f00640079,007300740061007400750073*/
		}{gopurs_runtime.Apply3(Call_Control_Semigroupoid_composeFlipped(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Call_Data_Argonaut_Encode_Class_encodeJson(Rebox_Main_2037062001_2831479053(Rebox_Main_2831479053_2037062001(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Encode_Class_EncodeJson[gopurs_runtime.Value]](gopurs_runtime.Apply(Call_Data_Argonaut_Encode_Class_encodeRecord(gopurs_runtime.Apply3(Call_Data_Argonaut_Encode_Class_gEncodeJsonCons(gopurs_runtime.Value{Type: 9, IntVal: 3228574649, UnsafePtr: unsafe.Pointer(Rebox_Main_946498665_2831479053(Rebox_Main_2831479053_946498665(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Encode_Class_EncodeJson[gopurs_runtime.Value]](Get_Data_Argonaut_Encode_Class_encodeJsonJBoolean()))))}), gopurs_runtime.Apply3(Call_Data_Argonaut_Encode_Class_gEncodeJsonCons(gopurs_runtime.Value{Type: 9, IntVal: 3228574649, UnsafePtr: unsafe.Pointer(Rebox_Main_2086421970_2831479053(Rebox_Main_2831479053_2086421970(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Encode_Class_EncodeJson[gopurs_runtime.Value]](Get_Data_Argonaut_Encode_Class_encodeJsonJString()))))}), gopurs_runtime.Apply3(Call_Data_Argonaut_Encode_Class_gEncodeJsonCons(gopurs_runtime.Value{Type: 9, IntVal: 3228574649, UnsafePtr: unsafe.Pointer(Rebox_Main_2086421970_2831479053(Rebox_Main_2831479053_2086421970(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Encode_Class_EncodeJson[gopurs_runtime.Value]](Get_Data_Argonaut_Encode_Class_encodeJsonJString()))))}), Get_Data_Argonaut_Encode_Class_gEncodeJsonNil(), Get_Main_titleIsSymbol(), gopurs_runtime.Value{}), gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.StrWithCodeUnits("id", 2)
		})), gopurs_runtime.Value{}), Get_Main_completedIsSymbol(), gopurs_runtime.Value{})), gopurs_runtime.Value{}))))), Get_Data_Argonaut_Core_stringify(), func() gopurs_runtime.Value {
			orig := __eta1_1
			_ = orig
			return gopurs_runtime.RecordDict3("completed", "id", "title", gopurs_runtime.Bool(orig.completed), gopurs_runtime.Str(orig.id), gopurs_runtime.Str(orig.title))
		}()).StrVal(), __eta0_0}
	}
}

func Call_Main_json__3478233858(dictEncodeJson_0_loop *Constructor_Data_Argonaut_Encode_Class_EncodeJson[struct {
	status string
	/*gopurs:record-v1:1:007300740061007400750073*/
}]) gopurs_runtime.Value {
json__3478233858:
	for {
		if false {
			continue json__3478233858
		}
		var dictEncodeJson_0 *Constructor_Data_Argonaut_Encode_Class_EncodeJson[struct {
			status string
			/*gopurs:record-v1:1:007300740061007400750073*/
		}] = dictEncodeJson_0_loop
		_ = dictEncodeJson_0
		// TAST (Let): toJsonString__2805964221_1_0 shape=App(Var) bindingType=(Func [(Record (Row [status: String] Empty))] String)
		toJsonString__2805964221_1_0 := gopurs_runtime.Apply2(Call_Control_Semigroupoid_composeFlipped(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Call_Data_Argonaut_Encode_Class_encodeJson(Rebox_Main_1331040862_2831479053(dictEncodeJson_0)), Get_Data_Argonaut_Core_stringify())
		_ = toJsonString__2805964221_1_0
		return gopurs_runtime.Func2(func(status_2 gopurs_runtime.Value, value_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.RecordDict2("body", "status", gopurs_runtime.Apply(toJsonString__2805964221_1_0, value_3), status_2)
		})
	}
}

func Call_Main_failure(status_0_loop int64, text_1_loop string) struct {
	body   string
	status int64
	/*gopurs:record-v1:2:0062006f00640079,007300740061007400750073*/
} {
	var status_0 int64 = status_0_loop
	_ = status_0
	var text_1 string = text_1_loop
	_ = text_1
	return struct {
		body   string
		status int64
		/*gopurs:record-v1:2:0062006f00640079,007300740061007400750073*/
	}{gopurs_runtime.Apply3(Call_Control_Semigroupoid_composeFlipped(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Call_Data_Argonaut_Encode_Class_encodeJson(Rebox_Main_614949298_2831479053(Rebox_Main_2831479053_614949298(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Encode_Class_EncodeJson[gopurs_runtime.Value]](Get_Main___cached_dict_2())))), Get_Data_Argonaut_Core_stringify(), func() gopurs_runtime.Value {
		orig := struct {
			error string
			/*gopurs:record-v1:1:006500720072006f0072*/
		}{text_1}
		_ = orig
		return gopurs_runtime.RecordDict1("error", gopurs_runtime.Str(orig.error))
	}()).StrVal(), status_0}
}

func Call_Main_decodeTask(v_0_loop []gopurs_runtime.Value) gopurs_runtime.Value {
	var v_0 []gopurs_runtime.Value = v_0_loop
	_ = v_0
	var __t4 gopurs_runtime.Value
	{
		if (gopurs_runtime.Int(int64(len(v_0))).IntVal) == (int64(3)) {
			// TAST (Let): __local_var_1_0 shape=Other bindingType=Any
			__local_var_1_0 := gopurs_runtime.ArrayAccess(gopurs_runtime.Array(v_0), 2)
			_ = __local_var_1_0
			// TAST (Let): __local_var_2_1 shape=Other bindingType=Any
			__local_var_2_1 := gopurs_runtime.ArrayAccess(gopurs_runtime.Array(v_0), 1)
			_ = __local_var_2_1
			// TAST (Let): v1_3_2 shape=App(Var) bindingType=(ADT ["Data","Either","Either"] [(ADT ["Data","NonEmpty","NonEmpty"] [(ADT ["Data","List","Types","List"] []), (ADT ["Foreign","ForeignError"] [])]), (Record (Row [id: String, title: String, completed: Boolean] Empty))])
			v1_3_2 := gopurs_runtime.Apply3(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Call_Data_Newtype_unwrap(gopurs_runtime.CoerceToStruct[Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{})), Get_Control_Monad_Except_Trans_runExceptT(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Control_Monad_Except_Trans_bindExceptT(gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](Get_Data_Identity_monadIdentity()))}), "bind"), gopurs_runtime.Apply(Call_Foreign_readString(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](Get_Data_Identity_monadIdentity())), gopurs_runtime.ArrayAccess(gopurs_runtime.Array(v_0), 0)), gopurs_runtime.Func(func(id_3 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Control_Monad_Except_Trans_bindExceptT(gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](Get_Data_Identity_monadIdentity()))}), "bind"), gopurs_runtime.Apply(Call_Foreign_readString(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](Get_Data_Identity_monadIdentity())), __local_var_2_1), gopurs_runtime.Func(func(title_4 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Control_Monad_Except_Trans_bindExceptT(gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](Get_Data_Identity_monadIdentity()))}), "bind"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Control_Monad_Except_Trans_functorExceptT(gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Identity_functorIdentity()))}), "map"), gopurs_runtime.Func(func(v2_5 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Bool(((v2_5.IntVal) == (int64(0))) != (true))
					}), gopurs_runtime.Apply(Call_Foreign_readInt(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](Get_Data_Identity_monadIdentity())), __local_var_1_0)), gopurs_runtime.Func(func(completed_5 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Control_Monad_Except_Trans_applicativeExceptT(gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](Get_Data_Identity_monadIdentity()))}), "pure"), func() gopurs_runtime.Value {
							orig := struct {
								completed bool
								id        string
								title     string
								/*gopurs:record-v1:3:0063006f006d0070006c0065007400650064,00690064,007400690074006c0065*/
							}{(completed_5.IntVal) != (0), id_3.StrVal(), title_4.StrVal()}
							_ = orig
							return gopurs_runtime.RecordDict3("completed", "id", "title", gopurs_runtime.Bool(orig.completed), gopurs_runtime.Str(orig.id), gopurs_runtime.Str(orig.title))
						}())
					}))
				}))
			})))
			_ = v1_3_2
			var __t3 gopurs_runtime.Value
			{
				if v1_3_2.Type == 9 && v1_3_2.IntVal == 3711209382 {
					__t3 = gopurs_runtime.Apply(Get_Effect_Aff__throwError(), gopurs_runtime.Apply(Get_Effect_Exception_error(), gopurs_runtime.Str(gopurs_runtime.ConcatString("Unexpected SQLite row: ", gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Data_List_Types_showNonEmptyList(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Foreign_showForeignError()))}), "show"), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Rebox_Main_3123684004_1293498952(Rebox_Main_1293498952_3123684004(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]((*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3_2.UnsafePtr).V0))))}).StrVal()))))
					goto end_branch_3
				} else {

				}
			}
			{
				if v1_3_2.Type == 9 && v1_3_2.IntVal == 2465973597 {
					__t3 = gopurs_runtime.Apply(Get_Effect_Aff__pure(), (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3_2.UnsafePtr).V0)
					goto end_branch_3
				} else {

				}
			}
			{
				__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_3:
			__t4 = __t3
			goto end_branch_4
		} else {

		}
	}
	{
		__t4 = gopurs_runtime.Apply(Get_Effect_Aff__throwError(), gopurs_runtime.Apply(Get_Effect_Exception_error(), gopurs_runtime.StrWithCodeUnits("Unexpected SQLite column count", 30)))
	}
end_branch_4:
	return __t4
}

func Call_Main_getTask(database_0_loop gopurs_runtime.Value, id_1_loop string) gopurs_runtime.Value {
	var database_0 gopurs_runtime.Value = database_0_loop
	_ = database_0
	var id_1 string = id_1_loop
	_ = id_1
	return gopurs_runtime.Apply2(Get_Effect_Aff__bind(), Call_Effect_Aff_Compat_fromEffectFnAff__2546930558(gopurs_runtime.Apply3(Get_Database_SQLite3_queryImpl(), database_0, gopurs_runtime.Str(gopurs_runtime.ConcatString("SELECT CAST(id AS TEXT), title, completed FROM tasks", " WHERE id = ?")), gopurs_runtime.Array([]gopurs_runtime.Value{gopurs_runtime.Str(id_1)}))), gopurs_runtime.Func(func(result_2 gopurs_runtime.Value) gopurs_runtime.Value {
		arrayIndex_arg_1 := Get_Data_Maybe_Just()
		_ = arrayIndex_arg_1
		arrayIndex_arg_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{gopurs_runtime.Value{}, false}
			if _v.V1 {
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
		}())
		_ = arrayIndex_arg_2
		arrayIndex_arg_3 := gopurs_runtime.RecordGet(result_2, "rows")
		_ = arrayIndex_arg_3
		arrayIndex_arg_4 := int64(0)
		_ = arrayIndex_arg_4
		// TAST (Let): v_3_0 shape=UncurriedApp(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(Array (ADT ["Foreign","Foreign"] []))])
		v_3_0 := Rebox_Main_3094389156_2807397954(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			arrayIndex_value := arrayIndex_arg_3
			_ = arrayIndex_value
			arrayIndex_source_5 := (*[]gopurs_runtime.Value)(arrayIndex_value.UnsafePtr)
			_ = arrayIndex_source_5
			if ((arrayIndex_arg_4) >= (0)) && ((arrayIndex_arg_4) < (int64(len(*arrayIndex_source_5)))) {
				return gopurs_runtime.Apply(arrayIndex_arg_1, gopurs_runtime.Array(*((*[]gopurs_runtime.Value)((*arrayIndex_source_5)[arrayIndex_arg_4].UnsafePtr))))
			} else {

			}
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(arrayIndex_arg_2)}
		}()))
		_ = v_3_0
		var __t5 gopurs_runtime.Value
		{
			if v_3_0 == nil {
				__t5 = gopurs_runtime.Apply(Get_Effect_Aff__pure(), func() gopurs_runtime.Value {
					orig := Call_Main_failure(int64(404), "Task not found")
					_ = orig
					return gopurs_runtime.RecordDict2("body", "status", gopurs_runtime.Str(orig.body), gopurs_runtime.Int(orig.status))
				}())
				goto end_branch_5
			} else {

			}
		}
		{
			if v_3_0 != nil {
				__t5 = gopurs_runtime.Apply2(Get_Effect_Aff__map(), gopurs_runtime.Apply(Get_Main_json__4251671259(), gopurs_runtime.Int(int64(200))), Call_Main_decodeTask((v_3_0).V0))
				goto end_branch_5
			} else {

			}
		}
		{
			__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
		}
	end_branch_5:
		return __t5
	}))
}

func Call_Main_tasks(database_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var database_0 gopurs_runtime.Value = database_0_loop
	_ = database_0
	return gopurs_runtime.Apply2(Get_Effect_Aff__bind(), Call_Effect_Aff_Compat_fromEffectFnAff__2546930558(gopurs_runtime.Apply3(Get_Database_SQLite3_queryImpl(), database_0, gopurs_runtime.Str(gopurs_runtime.ConcatString("SELECT CAST(id AS TEXT), title, completed FROM tasks", " ORDER BY id")), gopurs_runtime.Array([]gopurs_runtime.Value{}))), gopurs_runtime.Func(func(result_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply3(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableArray()).V3, gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Effect_Aff_applicativeAff()))}, Get_Main_decodeTask(), gopurs_runtime.RecordGet(result_1, "rows"))
	}))
}

func Call_Main_route(database_0_loop gopurs_runtime.Value, request_1_loop struct {
	body   string
	method string
	path   string
	/*gopurs:record-v1:3:0062006f00640079,006d006500740068006f0064,0070006100740068*/
}) gopurs_runtime.Value {
	var database_0 gopurs_runtime.Value = database_0_loop
	_ = database_0
	var request_1 struct {
		body   string
		method string
		path   string
		/*gopurs:record-v1:3:0062006f00640079,006d006500740068006f0064,0070006100740068*/
	} = request_1_loop
	_ = request_1
	// TAST (Let): v_2_0 shape=App(Var) bindingType=(Array String)
	v_2_0 := Data_String_Common_Split("/", request_1.path)
	_ = v_2_0
	var __t14 gopurs_runtime.Value
	{
		if (gopurs_runtime.Int(int64(len(v_2_0))).IntVal) == (int64(2)) {
			var __t7 gopurs_runtime.Value
			{
				if (gopurs_runtime.ArrayAccess(func() gopurs_runtime.Value {
					arr := v_2_0
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Str(v)
					}
					return gopurs_runtime.Array(boxed)
				}(), 0).StrVal()) == ("") {
					var __t6 gopurs_runtime.Value
					{
						if (request_1.method) == ("GET") {
							var __t1 gopurs_runtime.Value
							{
								if (gopurs_runtime.ArrayAccess(func() gopurs_runtime.Value {
									arr := v_2_0
									boxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										boxed[i] = gopurs_runtime.Str(v)
									}
									return gopurs_runtime.Array(boxed)
								}(), 1).StrVal()) == ("health") {
									__t1 = gopurs_runtime.Apply(Get_Effect_Aff__pure(), func() gopurs_runtime.Value {
										orig := struct {
											body   string
											status int64
											/*gopurs:record-v1:2:0062006f00640079,007300740061007400750073*/
										}{gopurs_runtime.Apply3(Call_Control_Semigroupoid_composeFlipped(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Call_Data_Argonaut_Encode_Class_encodeJson(Rebox_Main_1331040862_2831479053(Rebox_Main_2831479053_1331040862(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Encode_Class_EncodeJson[gopurs_runtime.Value]](Get_Main___cached_dict_3())))), Get_Data_Argonaut_Core_stringify(), func() gopurs_runtime.Value {
											orig := struct {
												status string
												/*gopurs:record-v1:1:007300740061007400750073*/
											}{"ok"}
											_ = orig
											return gopurs_runtime.RecordDict1("status", gopurs_runtime.Str(orig.status))
										}()).StrVal(), int64(200)}
										_ = orig
										return gopurs_runtime.RecordDict2("body", "status", gopurs_runtime.Str(orig.body), gopurs_runtime.Int(orig.status))
									}())
									goto end_branch_1
								} else {

								}
							}
							{
								if (gopurs_runtime.ArrayAccess(func() gopurs_runtime.Value {
									arr := v_2_0
									boxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										boxed[i] = gopurs_runtime.Str(v)
									}
									return gopurs_runtime.Array(boxed)
								}(), 1).StrVal()) == ("tasks") {
									__t1 = gopurs_runtime.Apply2(Get_Effect_Aff__map(), gopurs_runtime.Apply(Get_Main_json__4250750939(), gopurs_runtime.Int(int64(200))), Call_Main_tasks(database_0))
									goto end_branch_1
								} else {

								}
							}
							{
								__t1 = gopurs_runtime.Apply(Get_Effect_Aff__pure(), func() gopurs_runtime.Value {
									orig := Call_Main_failure(int64(404), "Route not found")
									_ = orig
									return gopurs_runtime.RecordDict2("body", "status", gopurs_runtime.Str(orig.body), gopurs_runtime.Int(orig.status))
								}())
							}
						end_branch_1:
							__t6 = __t1
							goto end_branch_6
						} else {

						}
					}
					{
						if ((request_1.method) == ("POST")) && ((gopurs_runtime.ArrayAccess(func() gopurs_runtime.Value {
							arr := v_2_0
							boxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								boxed[i] = gopurs_runtime.Str(v)
							}
							return gopurs_runtime.Array(boxed)
						}(), 1).StrVal()) == ("tasks")) {
							// TAST (Let): v2_3_2 shape=App(Var) bindingType=(ADT ["Data","Either","Either"] [(ADT ["Data","Argonaut","Decode","Error","JsonDecodeError"] []), (Record (Row [title: String] Empty))])
							v2_3_2 := Call_Data_Argonaut_Decode_fromJsonString__2215306652(request_1.body)
							_ = v2_3_2
							var __t5 gopurs_runtime.Value
							{
								if !v2_3_2.V2 {
									__t5 = gopurs_runtime.Apply(Get_Effect_Aff__pure(), func() gopurs_runtime.Value {
										orig := Call_Main_failure(int64(400), "Expected a JSON object with a string title")
										_ = orig
										return gopurs_runtime.RecordDict2("body", "status", gopurs_runtime.Str(orig.body), gopurs_runtime.Int(orig.status))
									}())
									goto end_branch_5
								} else {

								}
							}
							{
								if v2_3_2.V2 {
									// TAST (Let): title_4_3 shape=App(Var) bindingType=String
									title_4_3 := Data_String_Common_Trim(v2_3_2.V1.title)
									_ = title_4_3
									var __t4 gopurs_runtime.Value
									{
										if ((title_4_3) == ("")) || ((gopurs_runtime.Apply(Get_Data_String_CodeUnits_length(), gopurs_runtime.Str(title_4_3)).IntVal) > (int64(200))) {
											__t4 = gopurs_runtime.Apply(Get_Effect_Aff__pure(), func() gopurs_runtime.Value {
												orig := Call_Main_failure(int64(400), "Title must contain 1 to 200 UTF-16 code units")
												_ = orig
												return gopurs_runtime.RecordDict2("body", "status", gopurs_runtime.Str(orig.body), gopurs_runtime.Int(orig.status))
											}())
											goto end_branch_4
										} else {

										}
									}
									{
										__t4 = gopurs_runtime.Apply2(Get_Effect_Aff__bind(), Call_Effect_Aff_Compat_fromEffectFnAff__560502078(gopurs_runtime.Apply3(Get_Database_SQLite3_executeImpl(), database_0, gopurs_runtime.StrWithCodeUnits("INSERT INTO tasks (title) VALUES (?)", 36), gopurs_runtime.Array([]gopurs_runtime.Value{gopurs_runtime.Str(title_4_3)}))), gopurs_runtime.Func(func(result_5 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Apply(Get_Effect_Aff__pure(), func() gopurs_runtime.Value {
												orig := Call_Main_json__4251671259(int64(201), struct {
													completed bool
													id        string
													title     string
													/*gopurs:record-v1:3:0063006f006d0070006c0065007400650064,00690064,007400690074006c0065*/
												}{false, gopurs_runtime.RecordGet(result_5, "lastInsertRowId").StrVal(), title_4_3})
												_ = orig
												return gopurs_runtime.RecordDict2("body", "status", gopurs_runtime.Str(orig.body), gopurs_runtime.Int(orig.status))
											}())
										}))
									}
								end_branch_4:
									__t5 = __t4
									goto end_branch_5
								} else {

								}
							}
							{
								__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
							}
						end_branch_5:
							__t6 = __t5
							goto end_branch_6
						} else {

						}
					}
					{
						__t6 = gopurs_runtime.Apply(Get_Effect_Aff__pure(), func() gopurs_runtime.Value {
							orig := Call_Main_failure(int64(404), "Route not found")
							_ = orig
							return gopurs_runtime.RecordDict2("body", "status", gopurs_runtime.Str(orig.body), gopurs_runtime.Int(orig.status))
						}())
					}
				end_branch_6:
					__t7 = __t6
					goto end_branch_7
				} else {

				}
			}
			{
				__t7 = gopurs_runtime.Apply(Get_Effect_Aff__pure(), func() gopurs_runtime.Value {
					orig := Call_Main_failure(int64(404), "Route not found")
					_ = orig
					return gopurs_runtime.RecordDict2("body", "status", gopurs_runtime.Str(orig.body), gopurs_runtime.Int(orig.status))
				}())
			}
		end_branch_7:
			__t14 = __t7
			goto end_branch_14
		} else {

		}
	}
	{
		if ((gopurs_runtime.Int(int64(len(v_2_0))).IntVal) == (int64(3))) && (((gopurs_runtime.ArrayAccess(func() gopurs_runtime.Value {
			arr := v_2_0
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = gopurs_runtime.Str(v)
			}
			return gopurs_runtime.Array(boxed)
		}(), 0).StrVal()) == ("")) && (((gopurs_runtime.ArrayAccess(func() gopurs_runtime.Value {
			arr := v_2_0
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = gopurs_runtime.Str(v)
			}
			return gopurs_runtime.Array(boxed)
		}(), 1).StrVal()) == ("tasks")) && (Call_Main_validId(gopurs_runtime.ArrayAccess(func() gopurs_runtime.Value {
			arr := v_2_0
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = gopurs_runtime.Str(v)
			}
			return gopurs_runtime.Array(boxed)
		}(), 2).StrVal())))) {
			var __t13 gopurs_runtime.Value
			{
				if (request_1.method) == ("GET") {
					__t13 = Call_Main_getTask(database_0, gopurs_runtime.ArrayAccess(func() gopurs_runtime.Value {
						arr := v_2_0
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Str(v)
						}
						return gopurs_runtime.Array(boxed)
					}(), 2).StrVal())
					goto end_branch_13
				} else {

				}
			}
			{
				if (request_1.method) == ("PATCH") {
					// TAST (Let): v2_3_8 shape=App(Var) bindingType=(ADT ["Data","Either","Either"] [(ADT ["Data","Argonaut","Decode","Error","JsonDecodeError"] []), (Record (Row [completed: Boolean] Empty))])
					v2_3_8 := Call_Data_Argonaut_Decode_fromJsonString__2316388992(request_1.body)
					_ = v2_3_8
					var __t11 gopurs_runtime.Value
					{
						if !v2_3_8.V2 {
							__t11 = gopurs_runtime.Apply(Get_Effect_Aff__pure(), func() gopurs_runtime.Value {
								orig := Call_Main_failure(int64(400), "Expected a JSON object with a boolean completed")
								_ = orig
								return gopurs_runtime.RecordDict2("body", "status", gopurs_runtime.Str(orig.body), gopurs_runtime.Int(orig.status))
							}())
							goto end_branch_11
						} else {

						}
					}
					{
						if v2_3_8.V2 {
							var __t9 int64
							{
								if v2_3_8.V1.completed {
									__t9 = int64(1)
									goto end_branch_9
								} else {

								}
							}
							{
								__t9 = int64(0)
							}
						end_branch_9:
							__t11 = gopurs_runtime.Apply2(Get_Effect_Aff__bind(), Call_Effect_Aff_Compat_fromEffectFnAff__560502078(gopurs_runtime.Apply3(Get_Database_SQLite3_executeImpl(), database_0, gopurs_runtime.StrWithCodeUnits("UPDATE tasks SET completed = ? WHERE id = ?", 43), gopurs_runtime.Array([]gopurs_runtime.Value{gopurs_runtime.Int(__t9), gopurs_runtime.ArrayAccess(func() gopurs_runtime.Value {
								arr := v_2_0
								boxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									boxed[i] = gopurs_runtime.Str(v)
								}
								return gopurs_runtime.Array(boxed)
							}(), 2)}))), gopurs_runtime.Func(func(result_4 gopurs_runtime.Value) gopurs_runtime.Value {
								var __t10 gopurs_runtime.Value
								{
									if (gopurs_runtime.RecordGet(result_4, "changes").IntVal) == (int64(0)) {
										__t10 = gopurs_runtime.Apply(Get_Effect_Aff__pure(), func() gopurs_runtime.Value {
											orig := Call_Main_failure(int64(404), "Task not found")
											_ = orig
											return gopurs_runtime.RecordDict2("body", "status", gopurs_runtime.Str(orig.body), gopurs_runtime.Int(orig.status))
										}())
										goto end_branch_10
									} else {

									}
								}
								{
									__t10 = Call_Main_getTask(database_0, gopurs_runtime.ArrayAccess(func() gopurs_runtime.Value {
										arr := v_2_0
										boxed := make([]gopurs_runtime.Value, len(arr))
										for i, v := range arr {
											boxed[i] = gopurs_runtime.Str(v)
										}
										return gopurs_runtime.Array(boxed)
									}(), 2).StrVal())
								}
							end_branch_10:
								return __t10
							}))
							goto end_branch_11
						} else {

						}
					}
					{
						__t11 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
					}
				end_branch_11:
					__t13 = __t11
					goto end_branch_13
				} else {

				}
			}
			{
				if (request_1.method) == ("DELETE") {
					__t13 = gopurs_runtime.Apply2(Get_Effect_Aff__bind(), Call_Effect_Aff_Compat_fromEffectFnAff__560502078(gopurs_runtime.Apply3(Get_Database_SQLite3_executeImpl(), database_0, gopurs_runtime.StrWithCodeUnits("DELETE FROM tasks WHERE id = ?", 30), gopurs_runtime.Array([]gopurs_runtime.Value{gopurs_runtime.ArrayAccess(func() gopurs_runtime.Value {
						arr := v_2_0
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Str(v)
						}
						return gopurs_runtime.Array(boxed)
					}(), 2)}))), gopurs_runtime.Func(func(result_3 gopurs_runtime.Value) gopurs_runtime.Value {
						var __t12 struct {
							body   string
							status int64
							/*gopurs:record-v1:2:0062006f00640079,007300740061007400750073*/
						}
						{
							if (gopurs_runtime.RecordGet(result_3, "changes").IntVal) == (int64(0)) {
								__t12 = Call_Main_failure(int64(404), "Task not found")
								goto end_branch_12
							} else {

							}
						}
						{
							__t12 = struct {
								body   string
								status int64
								/*gopurs:record-v1:2:0062006f00640079,007300740061007400750073*/
							}{"", int64(204)}
						}
					end_branch_12:
						return gopurs_runtime.Apply(Get_Effect_Aff__pure(), func() gopurs_runtime.Value {
							orig := __t12
							_ = orig
							return gopurs_runtime.RecordDict2("body", "status", gopurs_runtime.Str(orig.body), gopurs_runtime.Int(orig.status))
						}())
					}))
					goto end_branch_13
				} else {

				}
			}
			{
				__t13 = gopurs_runtime.Apply(Get_Effect_Aff__pure(), func() gopurs_runtime.Value {
					orig := Call_Main_failure(int64(405), "Method not allowed")
					_ = orig
					return gopurs_runtime.RecordDict2("body", "status", gopurs_runtime.Str(orig.body), gopurs_runtime.Int(orig.status))
				}())
			}
		end_branch_13:
			__t14 = __t13
			goto end_branch_14
		} else {

		}
	}
	{
		__t14 = gopurs_runtime.Apply(Get_Effect_Aff__pure(), func() gopurs_runtime.Value {
			orig := Call_Main_failure(int64(404), "Route not found")
			_ = orig
			return gopurs_runtime.RecordDict2("body", "status", gopurs_runtime.Str(orig.body), gopurs_runtime.Int(orig.status))
		}())
	}
end_branch_14:
	return __t14
}

func Rebox_Main_1293498952_3123684004(in *Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_1331040862_2831479053(in *Constructor_Data_Argonaut_Encode_Class_EncodeJson[struct {
	status string
	/*gopurs:record-v1:1:007300740061007400750073*/
}]) *Constructor_Data_Argonaut_Encode_Class_EncodeJson[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Argonaut_Encode_Class_EncodeJson[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_2037062001_2831479053(in *Constructor_Data_Argonaut_Encode_Class_EncodeJson[struct {
	completed bool
	id        string
	title     string
	/*gopurs:record-v1:3:0063006f006d0070006c0065007400650064,00690064,007400690074006c0065*/
}]) *Constructor_Data_Argonaut_Encode_Class_EncodeJson[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Argonaut_Encode_Class_EncodeJson[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_2086421970_2831479053(in *Constructor_Data_Argonaut_Encode_Class_EncodeJson[string]) *Constructor_Data_Argonaut_Encode_Class_EncodeJson[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Argonaut_Encode_Class_EncodeJson[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_2831479053_1331040862(in *Constructor_Data_Argonaut_Encode_Class_EncodeJson[gopurs_runtime.Value]) *Constructor_Data_Argonaut_Encode_Class_EncodeJson[struct {
	status string
	/*gopurs:record-v1:1:007300740061007400750073*/
}] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Argonaut_Encode_Class_EncodeJson[struct {
		status string
		/*gopurs:record-v1:1:007300740061007400750073*/
	}])(unsafe.Pointer(in))
}

func Rebox_Main_2831479053_2037062001(in *Constructor_Data_Argonaut_Encode_Class_EncodeJson[gopurs_runtime.Value]) *Constructor_Data_Argonaut_Encode_Class_EncodeJson[struct {
	completed bool
	id        string
	title     string
	/*gopurs:record-v1:3:0063006f006d0070006c0065007400650064,00690064,007400690074006c0065*/
}] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Argonaut_Encode_Class_EncodeJson[struct {
		completed bool
		id        string
		title     string
		/*gopurs:record-v1:3:0063006f006d0070006c0065007400650064,00690064,007400690074006c0065*/
	}])(unsafe.Pointer(in))
}

func Rebox_Main_2831479053_2086421970(in *Constructor_Data_Argonaut_Encode_Class_EncodeJson[gopurs_runtime.Value]) *Constructor_Data_Argonaut_Encode_Class_EncodeJson[string] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Argonaut_Encode_Class_EncodeJson[string])(unsafe.Pointer(in))
}

func Rebox_Main_2831479053_614949298(in *Constructor_Data_Argonaut_Encode_Class_EncodeJson[gopurs_runtime.Value]) *Constructor_Data_Argonaut_Encode_Class_EncodeJson[struct {
	error string
	/*gopurs:record-v1:1:006500720072006f0072*/
}] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Argonaut_Encode_Class_EncodeJson[struct {
		error string
		/*gopurs:record-v1:1:006500720072006f0072*/
	}])(unsafe.Pointer(in))
}

func Rebox_Main_2831479053_946498665(in *Constructor_Data_Argonaut_Encode_Class_EncodeJson[gopurs_runtime.Value]) *Constructor_Data_Argonaut_Encode_Class_EncodeJson[bool] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Argonaut_Encode_Class_EncodeJson[bool])(unsafe.Pointer(in))
}

func Rebox_Main_3094389156_2807397954(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[[]gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[[]gopurs_runtime.Value]{}
	out.V0 = *((*[]gopurs_runtime.Value)(in.V0.UnsafePtr))
	return out
}

func Rebox_Main_3123684004_1293498952(in *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) *Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_614949298_2831479053(in *Constructor_Data_Argonaut_Encode_Class_EncodeJson[struct {
	error string
	/*gopurs:record-v1:1:006500720072006f0072*/
}]) *Constructor_Data_Argonaut_Encode_Class_EncodeJson[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Argonaut_Encode_Class_EncodeJson[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_946498665_2831479053(in *Constructor_Data_Argonaut_Encode_Class_EncodeJson[bool]) *Constructor_Data_Argonaut_Encode_Class_EncodeJson[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Argonaut_Encode_Class_EncodeJson[gopurs_runtime.Value])(unsafe.Pointer(in))
}
