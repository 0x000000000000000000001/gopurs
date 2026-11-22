package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Monoid_semigroupRecord gopurs_runtime.Value
var once_Data_Monoid_semigroupRecord sync.Once

func Get_Data_Monoid_semigroupRecord() gopurs_runtime.Value {
	once_Data_Monoid_semigroupRecord.Do(func() {
		cache_Data_Monoid_semigroupRecord = gopurs_runtime.Apply(Get_Data_Semigroup_semigroupRecord(), gopurs_runtime.Value{})
	})
	return cache_Data_Monoid_semigroupRecord
}

var cache_Data_Monoid_lessThanOrEq gopurs_runtime.Value
var once_Data_Monoid_lessThanOrEq sync.Once

func Get_Data_Monoid_lessThanOrEq() gopurs_runtime.Value {
	once_Data_Monoid_lessThanOrEq.Do(func() {
		cache_Data_Monoid_lessThanOrEq = Get_Data_Ord_lessThanOrEq__3348776441()
	})
	return cache_Data_Monoid_lessThanOrEq
}

var cache_Data_Monoid_eq gopurs_runtime.Value
var once_Data_Monoid_eq sync.Once

func Get_Data_Monoid_eq() gopurs_runtime.Value {
	once_Data_Monoid_eq.Do(func() {
		cache_Data_Monoid_eq = Get_Data_Eq_eq__3020190905()
	})
	return cache_Data_Monoid_eq
}

var cache_Data_Monoid_mod gopurs_runtime.Value
var once_Data_Monoid_mod sync.Once

func Get_Data_Monoid_mod() gopurs_runtime.Value {
	once_Data_Monoid_mod.Do(func() {
		cache_Data_Monoid_mod = Get_Data_EuclideanRing_mod__478053998()
	})
	return cache_Data_Monoid_mod
}

var cache_Data_Monoid_div gopurs_runtime.Value
var once_Data_Monoid_div sync.Once

func Get_Data_Monoid_div() gopurs_runtime.Value {
	once_Data_Monoid_div.Do(func() {
		cache_Data_Monoid_div = Get_Data_EuclideanRing_div__478053998()
	})
	return cache_Data_Monoid_div
}

var cache_Data_Monoid_MonoidRecord_dollar_Dict gopurs_runtime.Value
var once_Data_Monoid_MonoidRecord_dollar_Dict sync.Once

func Get_Data_Monoid_MonoidRecord_dollar_Dict() gopurs_runtime.Value {
	once_Data_Monoid_MonoidRecord_dollar_Dict.Do(func() {
		cache_Data_Monoid_MonoidRecord_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2415148183, UnsafePtr: unsafe.Pointer(Call_Data_Monoid_MonoidRecord_dollar_Dict(func() struct {
				SemigroupRecord0 gopurs_runtime.Value
				memptyRecord     gopurs_runtime.Value
				/*gopurs:record-v1:2:00530065006d006900670072006f00750070005200650063006f007200640030,006d0065006d007000740079005200650063006f00720064*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					SemigroupRecord0 gopurs_runtime.Value
					memptyRecord     gopurs_runtime.Value
					/*gopurs:record-v1:2:00530065006d006900670072006f00750070005200650063006f007200640030,006d0065006d007000740079005200650063006f00720064*/
				}{}
				clone.SemigroupRecord0 = gopurs_runtime.RecordGet(orig, "SemigroupRecord0")
				clone.memptyRecord = gopurs_runtime.RecordGet(orig, "memptyRecord")
				return clone
			}()))}
		})
	})
	return cache_Data_Monoid_MonoidRecord_dollar_Dict
}

var cache_Data_Monoid_Monoid_dollar_Dict gopurs_runtime.Value
var once_Data_Monoid_Monoid_dollar_Dict sync.Once

func Get_Data_Monoid_Monoid_dollar_Dict() gopurs_runtime.Value {
	once_Data_Monoid_Monoid_dollar_Dict.Do(func() {
		cache_Data_Monoid_Monoid_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(Call_Data_Monoid_Monoid_dollar_Dict(func() struct {
				Semigroup0 gopurs_runtime.Value
				mempty     gopurs_runtime.Value
				/*gopurs:record-v1:2:00530065006d006900670072006f007500700030,006d0065006d007000740079*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Semigroup0 gopurs_runtime.Value
					mempty     gopurs_runtime.Value
					/*gopurs:record-v1:2:00530065006d006900670072006f007500700030,006d0065006d007000740079*/
				}{}
				clone.Semigroup0 = gopurs_runtime.RecordGet(orig, "Semigroup0")
				clone.mempty = gopurs_runtime.RecordGet(orig, "mempty")
				return clone
			}()))}
		})
	})
	return cache_Data_Monoid_Monoid_dollar_Dict
}

var cache_Data_Monoid_Monoid_dollar_Dict__3455996042 gopurs_runtime.Value
var once_Data_Monoid_Monoid_dollar_Dict__3455996042 sync.Once

func Get_Data_Monoid_Monoid_dollar_Dict__3455996042() gopurs_runtime.Value {
	once_Data_Monoid_Monoid_dollar_Dict__3455996042.Do(func() {
		cache_Data_Monoid_Monoid_dollar_Dict__3455996042 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(Call_Data_Monoid_Monoid_dollar_Dict__3455996042(func() struct {
				Semigroup0 gopurs_runtime.Value
				mempty     gopurs_runtime.Value
				/*gopurs:record-v1:2:00530065006d006900670072006f007500700030,006d0065006d007000740079*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Semigroup0 gopurs_runtime.Value
					mempty     gopurs_runtime.Value
					/*gopurs:record-v1:2:00530065006d006900670072006f007500700030,006d0065006d007000740079*/
				}{}
				clone.Semigroup0 = gopurs_runtime.RecordGet(orig, "Semigroup0")
				clone.mempty = gopurs_runtime.RecordGet(orig, "mempty")
				return clone
			}()))}
		})
	})
	return cache_Data_Monoid_Monoid_dollar_Dict__3455996042
}

var cache_Data_Monoid_Monoid_dollar_Dict__2980887040 gopurs_runtime.Value
var once_Data_Monoid_Monoid_dollar_Dict__2980887040 sync.Once

func Get_Data_Monoid_Monoid_dollar_Dict__2980887040() gopurs_runtime.Value {
	once_Data_Monoid_Monoid_dollar_Dict__2980887040.Do(func() {
		cache_Data_Monoid_Monoid_dollar_Dict__2980887040 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(Rebox_Data_Monoid_2443930307_1201789390(Call_Data_Monoid_Monoid_dollar_Dict__2980887040(func() struct {
				Semigroup0 gopurs_runtime.Value
				mempty     uint32
				/*gopurs:record-v1:2:00530065006d006900670072006f007500700030,006d0065006d007000740079*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Semigroup0 gopurs_runtime.Value
					mempty     uint32
					/*gopurs:record-v1:2:00530065006d006900670072006f007500700030,006d0065006d007000740079*/
				}{}
				clone.Semigroup0 = gopurs_runtime.RecordGet(orig, "Semigroup0")
				clone.mempty = uint32(gopurs_runtime.RecordGet(orig, "mempty").IntVal)
				return clone
			}())))}
		})
	})
	return cache_Data_Monoid_Monoid_dollar_Dict__2980887040
}

var cache_Data_Monoid_Monoid_dollar_Dict__3050967606 gopurs_runtime.Value
var once_Data_Monoid_Monoid_dollar_Dict__3050967606 sync.Once

func Get_Data_Monoid_Monoid_dollar_Dict__3050967606() gopurs_runtime.Value {
	once_Data_Monoid_Monoid_dollar_Dict__3050967606.Do(func() {
		cache_Data_Monoid_Monoid_dollar_Dict__3050967606 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(Call_Data_Monoid_Monoid_dollar_Dict__3050967606(func() struct {
				Semigroup0 gopurs_runtime.Value
				mempty     gopurs_runtime.Value
				/*gopurs:record-v1:2:00530065006d006900670072006f007500700030,006d0065006d007000740079*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Semigroup0 gopurs_runtime.Value
					mempty     gopurs_runtime.Value
					/*gopurs:record-v1:2:00530065006d006900670072006f007500700030,006d0065006d007000740079*/
				}{}
				clone.Semigroup0 = gopurs_runtime.RecordGet(orig, "Semigroup0")
				clone.mempty = gopurs_runtime.RecordGet(orig, "mempty")
				return clone
			}()))}
		})
	})
	return cache_Data_Monoid_Monoid_dollar_Dict__3050967606
}

var cache_Data_Monoid_Monoid_dollar_Dict__676023645 gopurs_runtime.Value
var once_Data_Monoid_Monoid_dollar_Dict__676023645 sync.Once

func Get_Data_Monoid_Monoid_dollar_Dict__676023645() gopurs_runtime.Value {
	once_Data_Monoid_Monoid_dollar_Dict__676023645.Do(func() {
		cache_Data_Monoid_Monoid_dollar_Dict__676023645 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(Rebox_Data_Monoid_321927638_1201789390(Call_Data_Monoid_Monoid_dollar_Dict__676023645(func() struct {
				Semigroup0 gopurs_runtime.Value
				mempty     float64
				/*gopurs:record-v1:2:00530065006d006900670072006f007500700030,006d0065006d007000740079*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Semigroup0 gopurs_runtime.Value
					mempty     float64
					/*gopurs:record-v1:2:00530065006d006900670072006f007500700030,006d0065006d007000740079*/
				}{}
				clone.Semigroup0 = gopurs_runtime.RecordGet(orig, "Semigroup0")
				clone.mempty = gopurs_runtime.RecordGet(orig, "mempty").FloatVal()
				return clone
			}())))}
		})
	})
	return cache_Data_Monoid_Monoid_dollar_Dict__676023645
}

var cache_Data_Monoid_Monoid_dollar_Dict__2308908914 gopurs_runtime.Value
var once_Data_Monoid_Monoid_dollar_Dict__2308908914 sync.Once

func Get_Data_Monoid_Monoid_dollar_Dict__2308908914() gopurs_runtime.Value {
	once_Data_Monoid_Monoid_dollar_Dict__2308908914.Do(func() {
		cache_Data_Monoid_Monoid_dollar_Dict__2308908914 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(Rebox_Data_Monoid_2860751713_1201789390(Call_Data_Monoid_Monoid_dollar_Dict__2308908914(func() struct {
				Semigroup0 gopurs_runtime.Value
				mempty     struct {
					dotAll     bool
					global     bool
					ignoreCase bool
					multiline  bool
					sticky     bool
					unicode    bool
					/*gopurs:record-v1:6:0064006f00740041006c006c,0067006c006f00620061006c,00690067006e006f007200650043006100730065,006d0075006c00740069006c0069006e0065,0073007400690063006b0079,0075006e00690063006f00640065*/
				}
				/*gopurs:record-v1:2:00530065006d006900670072006f007500700030,006d0065006d007000740079*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Semigroup0 gopurs_runtime.Value
					mempty     struct {
						dotAll     bool
						global     bool
						ignoreCase bool
						multiline  bool
						sticky     bool
						unicode    bool
						/*gopurs:record-v1:6:0064006f00740041006c006c,0067006c006f00620061006c,00690067006e006f007200650043006100730065,006d0075006c00740069006c0069006e0065,0073007400690063006b0079,0075006e00690063006f00640065*/
					}
					/*gopurs:record-v1:2:00530065006d006900670072006f007500700030,006d0065006d007000740079*/
				}{}
				clone.Semigroup0 = gopurs_runtime.RecordGet(orig, "Semigroup0")
				clone.mempty = func() struct {
					dotAll     bool
					global     bool
					ignoreCase bool
					multiline  bool
					sticky     bool
					unicode    bool
					/*gopurs:record-v1:6:0064006f00740041006c006c,0067006c006f00620061006c,00690067006e006f007200650043006100730065,006d0075006c00740069006c0069006e0065,0073007400690063006b0079,0075006e00690063006f00640065*/
				} {
					orig := gopurs_runtime.RecordGet(orig, "mempty")
					_ = orig
					clone := struct {
						dotAll     bool
						global     bool
						ignoreCase bool
						multiline  bool
						sticky     bool
						unicode    bool
						/*gopurs:record-v1:6:0064006f00740041006c006c,0067006c006f00620061006c,00690067006e006f007200650043006100730065,006d0075006c00740069006c0069006e0065,0073007400690063006b0079,0075006e00690063006f00640065*/
					}{}
					clone.dotAll = (gopurs_runtime.RecordGet(orig, "dotAll").IntVal) != (0)
					clone.global = (gopurs_runtime.RecordGet(orig, "global").IntVal) != (0)
					clone.ignoreCase = (gopurs_runtime.RecordGet(orig, "ignoreCase").IntVal) != (0)
					clone.multiline = (gopurs_runtime.RecordGet(orig, "multiline").IntVal) != (0)
					clone.sticky = (gopurs_runtime.RecordGet(orig, "sticky").IntVal) != (0)
					clone.unicode = (gopurs_runtime.RecordGet(orig, "unicode").IntVal) != (0)
					return clone
				}()
				return clone
			}())))}
		})
	})
	return cache_Data_Monoid_Monoid_dollar_Dict__2308908914
}

var cache_Data_Monoid_Monoid_dollar_Dict__3607938411 gopurs_runtime.Value
var once_Data_Monoid_Monoid_dollar_Dict__3607938411 sync.Once

func Get_Data_Monoid_Monoid_dollar_Dict__3607938411() gopurs_runtime.Value {
	once_Data_Monoid_Monoid_dollar_Dict__3607938411.Do(func() {
		cache_Data_Monoid_Monoid_dollar_Dict__3607938411 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(Rebox_Data_Monoid_1950344881_1201789390(Call_Data_Monoid_Monoid_dollar_Dict__3607938411(func() struct {
				Semigroup0 gopurs_runtime.Value
				mempty     string
				/*gopurs:record-v1:2:00530065006d006900670072006f007500700030,006d0065006d007000740079*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Semigroup0 gopurs_runtime.Value
					mempty     string
					/*gopurs:record-v1:2:00530065006d006900670072006f007500700030,006d0065006d007000740079*/
				}{}
				clone.Semigroup0 = gopurs_runtime.RecordGet(orig, "Semigroup0")
				clone.mempty = gopurs_runtime.RecordGet(orig, "mempty").StrVal()
				return clone
			}())))}
		})
	})
	return cache_Data_Monoid_Monoid_dollar_Dict__3607938411
}

var cache_Data_Monoid_Monoid_dollar_Dict__2878286104 gopurs_runtime.Value
var once_Data_Monoid_Monoid_dollar_Dict__2878286104 sync.Once

func Get_Data_Monoid_Monoid_dollar_Dict__2878286104() gopurs_runtime.Value {
	once_Data_Monoid_Monoid_dollar_Dict__2878286104.Do(func() {
		cache_Data_Monoid_Monoid_dollar_Dict__2878286104 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(Call_Data_Monoid_Monoid_dollar_Dict__2878286104(func() struct {
				Semigroup0 gopurs_runtime.Value
				mempty     gopurs_runtime.Value
				/*gopurs:record-v1:2:00530065006d006900670072006f007500700030,006d0065006d007000740079*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Semigroup0 gopurs_runtime.Value
					mempty     gopurs_runtime.Value
					/*gopurs:record-v1:2:00530065006d006900670072006f007500700030,006d0065006d007000740079*/
				}{}
				clone.Semigroup0 = gopurs_runtime.RecordGet(orig, "Semigroup0")
				clone.mempty = gopurs_runtime.RecordGet(orig, "mempty")
				return clone
			}()))}
		})
	})
	return cache_Data_Monoid_Monoid_dollar_Dict__2878286104
}

var cache_Data_Monoid_monoidUnit gopurs_runtime.Value
var once_Data_Monoid_monoidUnit sync.Once

func Get_Data_Monoid_monoidUnit() gopurs_runtime.Value {
	once_Data_Monoid_monoidUnit.Do(func() {
		cache_Data_Monoid_monoidUnit = gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer((&Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](Get_Data_Semigroup_semigroupUnit()))}
		}), Get_Data_Unit_unit()}))}
	})
	return cache_Data_Monoid_monoidUnit
}

var cache_Data_Monoid_monoidString gopurs_runtime.Value
var once_Data_Monoid_monoidString sync.Once

func Get_Data_Monoid_monoidString() gopurs_runtime.Value {
	once_Data_Monoid_monoidString.Do(func() {
		cache_Data_Monoid_monoidString = gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(Rebox_Data_Monoid_1950344881_1201789390((&Constructor_Data_Monoid_Monoid[string]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(Rebox_Data_Monoid_443971153_4179793454(Rebox_Data_Monoid_4179793454_443971153(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](Get_Data_Semigroup_semigroupString()))))}
		}), ""})))}
	})
	return cache_Data_Monoid_monoidString
}

var cache_Data_Monoid_monoidRecordNil gopurs_runtime.Value
var once_Data_Monoid_monoidRecordNil sync.Once

func Get_Data_Monoid_monoidRecordNil() gopurs_runtime.Value {
	once_Data_Monoid_monoidRecordNil.Do(func() {
		cache_Data_Monoid_monoidRecordNil = gopurs_runtime.Value{Type: 9, IntVal: 2415148183, UnsafePtr: unsafe.Pointer((&Constructor_Data_Monoid_MonoidRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Get_Data_Semigroup_semigroupRecordNil()
		}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.RecordDict0()
		})}))}
	})
	return cache_Data_Monoid_monoidRecordNil
}

var cache_Data_Monoid_monoidOrdering gopurs_runtime.Value
var once_Data_Monoid_monoidOrdering sync.Once

func Get_Data_Monoid_monoidOrdering() gopurs_runtime.Value {
	once_Data_Monoid_monoidOrdering.Do(func() {
		cache_Data_Monoid_monoidOrdering = gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(Rebox_Data_Monoid_2443930307_1201789390((&Constructor_Data_Monoid_Monoid[uint32]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(Rebox_Data_Monoid_1625289059_4179793454(Rebox_Data_Monoid_4179793454_1625289059(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](Get_Data_Ordering_semigroupOrdering()))))}
		}), 902936544})))}
	})
	return cache_Data_Monoid_monoidOrdering
}

var cache_Data_Monoid_monoidArray gopurs_runtime.Value
var once_Data_Monoid_monoidArray sync.Once

func Get_Data_Monoid_monoidArray() gopurs_runtime.Value {
	once_Data_Monoid_monoidArray.Do(func() {
		cache_Data_Monoid_monoidArray = gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(Rebox_Data_Monoid_67307368_1201789390((&Constructor_Data_Monoid_Monoid[[]gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Get_Data_Semigroup_semigroupArray()
		}), *((*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{}).UnsafePtr))})))}
	})
	return cache_Data_Monoid_monoidArray
}

var cache_Data_Monoid_memptyRecord gopurs_runtime.Value
var once_Data_Monoid_memptyRecord sync.Once

func Get_Data_Monoid_memptyRecord() gopurs_runtime.Value {
	once_Data_Monoid_memptyRecord.Do(func() {
		cache_Data_Monoid_memptyRecord = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Monoid_memptyRecord(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_MonoidRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Data_Monoid_memptyRecord
}

var cache_Data_Monoid_monoidRecord gopurs_runtime.Value
var once_Data_Monoid_monoidRecord sync.Once

func Get_Data_Monoid_monoidRecord() gopurs_runtime.Value {
	once_Data_Monoid_monoidRecord.Do(func() {
		cache_Data_Monoid_monoidRecord = gopurs_runtime.Func2(func(_dollar___unused_0_box gopurs_runtime.Value, dictMonoidRecord_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Monoid_monoidRecord(_dollar___unused_0_box, dictMonoidRecord_1_box)
		})
	})
	return cache_Data_Monoid_monoidRecord
}

var cache_Data_Monoid_mempty gopurs_runtime.Value
var once_Data_Monoid_mempty sync.Once

func Get_Data_Monoid_mempty() gopurs_runtime.Value {
	once_Data_Monoid_mempty.Do(func() {
		cache_Data_Monoid_mempty = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Monoid_mempty(dict_0_box)
		})
	})
	return cache_Data_Monoid_mempty
}

var cache_Data_Monoid_mempty__1856050949 gopurs_runtime.Value
var once_Data_Monoid_mempty__1856050949 sync.Once

func Get_Data_Monoid_mempty__1856050949() gopurs_runtime.Value {
	once_Data_Monoid_mempty__1856050949.Do(func() {
		cache_Data_Monoid_mempty__1856050949 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Monoid_mempty__1856050949(__eta0_0_box)
		})
	})
	return cache_Data_Monoid_mempty__1856050949
}

var cache_Data_Monoid_monoidFn gopurs_runtime.Value
var once_Data_Monoid_monoidFn sync.Once

func Get_Data_Monoid_monoidFn() gopurs_runtime.Value {
	once_Data_Monoid_monoidFn.Do(func() {
		cache_Data_Monoid_monoidFn = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Monoid_monoidFn(dictMonoid_0_box)
		})
	})
	return cache_Data_Monoid_monoidFn
}

var cache_Data_Monoid_monoidRecordCons gopurs_runtime.Value
var once_Data_Monoid_monoidRecordCons sync.Once

func Get_Data_Monoid_monoidRecordCons() gopurs_runtime.Value {
	once_Data_Monoid_monoidRecordCons.Do(func() {
		cache_Data_Monoid_monoidRecordCons = gopurs_runtime.Func2(func(dictIsSymbol_0_box gopurs_runtime.Value, dictMonoid_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Monoid_monoidRecordCons(dictIsSymbol_0_box, dictMonoid_1_box)
		})
	})
	return cache_Data_Monoid_monoidRecordCons
}

var cache_Data_Monoid_power gopurs_runtime.Value
var once_Data_Monoid_power sync.Once

func Get_Data_Monoid_power() gopurs_runtime.Value {
	once_Data_Monoid_power.Do(func() {
		cache_Data_Monoid_power = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Monoid_power(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_0_box))
		})
	})
	return cache_Data_Monoid_power
}

var cache_Data_Monoid_guard gopurs_runtime.Value
var once_Data_Monoid_guard sync.Once

func Get_Data_Monoid_guard() gopurs_runtime.Value {
	once_Data_Monoid_guard.Do(func() {
		cache_Data_Monoid_guard = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Monoid_guard(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_0_box))
		})
	})
	return cache_Data_Monoid_guard
}

type Constructor_Data_Monoid_MonoidRecord[T_rowlist any, T_row any, T_subrow any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}

//gopurs:init-v1:00670065006e006500720061007400650064002f0044006100740061005f004d006f006e006f00690064:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f0044006100740061005f004d006f006e006f00690064005f004d006f006e006f00690064005200650063006f00720064
func init() {
	gopurs_runtime.StructGetters[2415148183] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Monoid_MonoidRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "SemigroupRecord0":
			return gopurs_runtime.Box(c.V0)
		case "memptyRecord":
			return gopurs_runtime.Box(c.V1)
		default:
			panic("Key not found in dictionary Constructor_Data_Monoid_MonoidRecord: " + key)
		}
	}
}

type Constructor_Data_Monoid_Monoid[T_m any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 T_m
}

//gopurs:init-v1:00670065006e006500720061007400650064002f0044006100740061005f004d006f006e006f00690064:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f0044006100740061005f004d006f006e006f00690064005f004d006f006e006f00690064
func init() {
	gopurs_runtime.StructGetters[1722653594] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Monoid_Monoid[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "Semigroup0":
			return gopurs_runtime.Box(c.V0)
		case "mempty":
			return gopurs_runtime.Box(c.V1)
		default:
			panic("Key not found in dictionary Constructor_Data_Monoid_Monoid: " + key)
		}
	}
}

func Call_Data_Monoid_MonoidRecord_dollar_Dict(x_0_loop struct {
	SemigroupRecord0 gopurs_runtime.Value
	memptyRecord     gopurs_runtime.Value
	/*gopurs:record-v1:2:00530065006d006900670072006f00750070005200650063006f007200640030,006d0065006d007000740079005200650063006f00720064*/
}) *Constructor_Data_Monoid_MonoidRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value] {
	var x_0 struct {
		SemigroupRecord0 gopurs_runtime.Value
		memptyRecord     gopurs_runtime.Value
		/*gopurs:record-v1:2:00530065006d006900670072006f00750070005200650063006f007200640030,006d0065006d007000740079005200650063006f00720064*/
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_MonoidRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict2("SemigroupRecord0", "memptyRecord", orig.SemigroupRecord0, orig.memptyRecord)
	}())
}

func Call_Data_Monoid_Monoid_dollar_Dict(x_0_loop struct {
	Semigroup0 gopurs_runtime.Value
	mempty     gopurs_runtime.Value
	/*gopurs:record-v1:2:00530065006d006900670072006f007500700030,006d0065006d007000740079*/
}) *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value] {
	var x_0 struct {
		Semigroup0 gopurs_runtime.Value
		mempty     gopurs_runtime.Value
		/*gopurs:record-v1:2:00530065006d006900670072006f007500700030,006d0065006d007000740079*/
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict2("Semigroup0", "mempty", orig.Semigroup0, orig.mempty)
	}())
}

func Call_Data_Monoid_Monoid_dollar_Dict__3455996042(x_0_loop struct {
	Semigroup0 gopurs_runtime.Value
	mempty     gopurs_runtime.Value
	/*gopurs:record-v1:2:00530065006d006900670072006f007500700030,006d0065006d007000740079*/
}) *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value] {
Monoid_dollar_Dict__3455996042:
	for {
		if false {
			continue Monoid_dollar_Dict__3455996042
		}
		var x_0 struct {
			Semigroup0 gopurs_runtime.Value
			mempty     gopurs_runtime.Value
			/*gopurs:record-v1:2:00530065006d006900670072006f007500700030,006d0065006d007000740079*/
		} = x_0_loop
		_ = x_0
		return gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			orig := x_0
			_ = orig
			return gopurs_runtime.RecordDict2("Semigroup0", "mempty", orig.Semigroup0, orig.mempty)
		}())
	}
}

func Call_Data_Monoid_Monoid_dollar_Dict__2980887040(x_0_loop struct {
	Semigroup0 gopurs_runtime.Value
	mempty     uint32
	/*gopurs:record-v1:2:00530065006d006900670072006f007500700030,006d0065006d007000740079*/
}) *Constructor_Data_Monoid_Monoid[uint32] {
Monoid_dollar_Dict__2980887040:
	for {
		if false {
			continue Monoid_dollar_Dict__2980887040
		}
		var x_0 struct {
			Semigroup0 gopurs_runtime.Value
			mempty     uint32
			/*gopurs:record-v1:2:00530065006d006900670072006f007500700030,006d0065006d007000740079*/
		} = x_0_loop
		_ = x_0
		return gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[uint32]](func() gopurs_runtime.Value {
			orig := x_0
			_ = orig
			return gopurs_runtime.RecordDict2("Semigroup0", "mempty", orig.Semigroup0, gopurs_runtime.Value{Type: 9, IntVal: int64(orig.mempty), UnsafePtr: nil})
		}())
	}
}

func Call_Data_Monoid_Monoid_dollar_Dict__3050967606(x_0_loop struct {
	Semigroup0 gopurs_runtime.Value
	mempty     gopurs_runtime.Value
	/*gopurs:record-v1:2:00530065006d006900670072006f007500700030,006d0065006d007000740079*/
}) *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value] {
Monoid_dollar_Dict__3050967606:
	for {
		if false {
			continue Monoid_dollar_Dict__3050967606
		}
		var x_0 struct {
			Semigroup0 gopurs_runtime.Value
			mempty     gopurs_runtime.Value
			/*gopurs:record-v1:2:00530065006d006900670072006f007500700030,006d0065006d007000740079*/
		} = x_0_loop
		_ = x_0
		return gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			orig := x_0
			_ = orig
			return gopurs_runtime.RecordDict2("Semigroup0", "mempty", orig.Semigroup0, orig.mempty)
		}())
	}
}

func Call_Data_Monoid_Monoid_dollar_Dict__676023645(x_0_loop struct {
	Semigroup0 gopurs_runtime.Value
	mempty     float64
	/*gopurs:record-v1:2:00530065006d006900670072006f007500700030,006d0065006d007000740079*/
}) *Constructor_Data_Monoid_Monoid[float64] {
Monoid_dollar_Dict__676023645:
	for {
		if false {
			continue Monoid_dollar_Dict__676023645
		}
		var x_0 struct {
			Semigroup0 gopurs_runtime.Value
			mempty     float64
			/*gopurs:record-v1:2:00530065006d006900670072006f007500700030,006d0065006d007000740079*/
		} = x_0_loop
		_ = x_0
		return gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[float64]](func() gopurs_runtime.Value {
			orig := x_0
			_ = orig
			return gopurs_runtime.RecordDict2("Semigroup0", "mempty", orig.Semigroup0, gopurs_runtime.Float(orig.mempty))
		}())
	}
}

func Call_Data_Monoid_Monoid_dollar_Dict__2308908914(x_0_loop struct {
	Semigroup0 gopurs_runtime.Value
	mempty     struct {
		dotAll     bool
		global     bool
		ignoreCase bool
		multiline  bool
		sticky     bool
		unicode    bool
		/*gopurs:record-v1:6:0064006f00740041006c006c,0067006c006f00620061006c,00690067006e006f007200650043006100730065,006d0075006c00740069006c0069006e0065,0073007400690063006b0079,0075006e00690063006f00640065*/
	}
	/*gopurs:record-v1:2:00530065006d006900670072006f007500700030,006d0065006d007000740079*/
}) *Constructor_Data_Monoid_Monoid[struct {
	dotAll     bool
	global     bool
	ignoreCase bool
	multiline  bool
	sticky     bool
	unicode    bool
	/*gopurs:record-v1:6:0064006f00740041006c006c,0067006c006f00620061006c,00690067006e006f007200650043006100730065,006d0075006c00740069006c0069006e0065,0073007400690063006b0079,0075006e00690063006f00640065*/
}] {
Monoid_dollar_Dict__2308908914:
	for {
		if false {
			continue Monoid_dollar_Dict__2308908914
		}
		var x_0 struct {
			Semigroup0 gopurs_runtime.Value
			mempty     struct {
				dotAll     bool
				global     bool
				ignoreCase bool
				multiline  bool
				sticky     bool
				unicode    bool
				/*gopurs:record-v1:6:0064006f00740041006c006c,0067006c006f00620061006c,00690067006e006f007200650043006100730065,006d0075006c00740069006c0069006e0065,0073007400690063006b0079,0075006e00690063006f00640065*/
			}
			/*gopurs:record-v1:2:00530065006d006900670072006f007500700030,006d0065006d007000740079*/
		} = x_0_loop
		_ = x_0
		return gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[struct {
			dotAll     bool
			global     bool
			ignoreCase bool
			multiline  bool
			sticky     bool
			unicode    bool
			/*gopurs:record-v1:6:0064006f00740041006c006c,0067006c006f00620061006c,00690067006e006f007200650043006100730065,006d0075006c00740069006c0069006e0065,0073007400690063006b0079,0075006e00690063006f00640065*/
		}]](func() gopurs_runtime.Value {
			orig := x_0
			_ = orig
			return gopurs_runtime.RecordDict2("Semigroup0", "mempty", orig.Semigroup0, func() gopurs_runtime.Value {
				orig := orig.mempty
				_ = orig
				return gopurs_runtime.RecordDict([]string{"dotAll", "global", "ignoreCase", "multiline", "sticky", "unicode"}, []gopurs_runtime.Value{gopurs_runtime.Bool(orig.dotAll), gopurs_runtime.Bool(orig.global), gopurs_runtime.Bool(orig.ignoreCase), gopurs_runtime.Bool(orig.multiline), gopurs_runtime.Bool(orig.sticky), gopurs_runtime.Bool(orig.unicode)})
			}())
		}())
	}
}

func Call_Data_Monoid_Monoid_dollar_Dict__3607938411(x_0_loop struct {
	Semigroup0 gopurs_runtime.Value
	mempty     string
	/*gopurs:record-v1:2:00530065006d006900670072006f007500700030,006d0065006d007000740079*/
}) *Constructor_Data_Monoid_Monoid[string] {
Monoid_dollar_Dict__3607938411:
	for {
		if false {
			continue Monoid_dollar_Dict__3607938411
		}
		var x_0 struct {
			Semigroup0 gopurs_runtime.Value
			mempty     string
			/*gopurs:record-v1:2:00530065006d006900670072006f007500700030,006d0065006d007000740079*/
		} = x_0_loop
		_ = x_0
		return gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[string]](func() gopurs_runtime.Value {
			orig := x_0
			_ = orig
			return gopurs_runtime.RecordDict2("Semigroup0", "mempty", orig.Semigroup0, gopurs_runtime.Str(orig.mempty))
		}())
	}
}

func Call_Data_Monoid_Monoid_dollar_Dict__2878286104(x_0_loop struct {
	Semigroup0 gopurs_runtime.Value
	mempty     gopurs_runtime.Value
	/*gopurs:record-v1:2:00530065006d006900670072006f007500700030,006d0065006d007000740079*/
}) *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value] {
Monoid_dollar_Dict__2878286104:
	for {
		if false {
			continue Monoid_dollar_Dict__2878286104
		}
		var x_0 struct {
			Semigroup0 gopurs_runtime.Value
			mempty     gopurs_runtime.Value
			/*gopurs:record-v1:2:00530065006d006900670072006f007500700030,006d0065006d007000740079*/
		} = x_0_loop
		_ = x_0
		return gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			orig := x_0
			_ = orig
			return gopurs_runtime.RecordDict2("Semigroup0", "mempty", orig.Semigroup0, orig.mempty)
		}())
	}
}

func Call_Data_Monoid_memptyRecord(dict_0_loop *Constructor_Data_Monoid_MonoidRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Data_Monoid_MonoidRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V1
}

func Call_Data_Monoid_monoidRecord(_dollar___unused_0_loop gopurs_runtime.Value, dictMonoidRecord_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var dictMonoidRecord_1 gopurs_runtime.Value = dictMonoidRecord_1_loop
	_ = dictMonoidRecord_1
	// TAST (Let): semigroupRecord1_2_0 shape=App(Var) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(Record (Row [] (TypeVar row$scope17)))])
	semigroupRecord1_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](Call_Data_Semigroup_semigroupRecord(gopurs_runtime.Value{}, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoidRecord_1, "SemigroupRecord0"), gopurs_runtime.Value{})))
	_ = semigroupRecord1_2_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer((&Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(semigroupRecord1_2_0)}
	}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoidRecord_1, "memptyRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil})}))}
}

func Call_Data_Monoid_mempty(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dict_0 gopurs_runtime.Value = dict_0_loop
	_ = dict_0
	return gopurs_runtime.RecordGet(dict_0, "mempty")
}

func Call_Data_Monoid_mempty__1856050949(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
mempty__1856050949:
	for {
		if false {
			continue mempty__1856050949
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		return gopurs_runtime.Apply(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](Get_Effect_Aff_monoidCanceler()).V1, __eta0_0)
	}
}

func Call_Data_Monoid_monoidFn(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
	_ = dictMonoid_0
	// TAST (Let): mempty1_1_0 shape=Other bindingType=(TypeVar b$scope25)
	mempty1_1_0 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
	_ = mempty1_1_0
	// TAST (Let): semigroupFn_2_1 shape=App(Var) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(Func [(TypeVar a$scope26)] (TypeVar b$scope25))])
	semigroupFn_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](Call_Data_Semigroup_semigroupFn(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})))
	_ = semigroupFn_2_1
	return gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer((&Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(semigroupFn_2_1)}
	}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return mempty1_1_0
	})}))}
}

func Call_Data_Monoid_monoidRecordCons(dictIsSymbol_0_loop gopurs_runtime.Value, dictMonoid_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictIsSymbol_0 gopurs_runtime.Value = dictIsSymbol_0_loop
	_ = dictIsSymbol_0
	var dictMonoid_1 gopurs_runtime.Value = dictMonoid_1_loop
	_ = dictMonoid_1
	// TAST (Let): mempty1_2_0 shape=App(Var) bindingType=(TypeVar focus$scope35)
	mempty1_2_0 := Call_Data_Monoid_mempty(dictMonoid_1)
	_ = mempty1_2_0
	// TAST (Let): Semigroup0_3_1 shape=App(Other) bindingType=Any
	Semigroup0_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_1, "Semigroup0"), gopurs_runtime.Value{})
	_ = Semigroup0_3_1
	return gopurs_runtime.Func2(func(_dollar___unused_4 gopurs_runtime.Value, dictMonoidRecord_5 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): semigroupRecordCons1_6_2 shape=App(Var) bindingType=(TypeApp (ADT ["Data","Semigroup","SemigroupRecord"] []) [(TypeApp (ADT ["Prim","RowList","Cons"] []) [(TypeVar key$scope34), (TypeVar focus$scope35), (TypeVar rowlistTail$scope38)]), (TypeVar row$scope39), (TypeVar subrow$scope37)])
		semigroupRecordCons1_6_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_SemigroupRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]](Call_Data_Semigroup_semigroupRecordCons(dictIsSymbol_0, gopurs_runtime.Value{}, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoidRecord_5, "SemigroupRecord0"), gopurs_runtime.Value{}), Semigroup0_3_1))
		_ = semigroupRecordCons1_6_2
		return gopurs_runtime.Value{Type: 9, IntVal: 2415148183, UnsafePtr: unsafe.Pointer((&Constructor_Data_Monoid_MonoidRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3847494007, UnsafePtr: unsafe.Pointer(semigroupRecordCons1_6_2)}
		}), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply3(Get_Record_Unsafe_unsafeSet(), gopurs_runtime.NormalizeStringValue(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIsSymbol_0, "reflectSymbol"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil})), mempty1_2_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoidRecord_5, "memptyRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}))
		})}))}
	})
}

func Call_Data_Monoid_power(dictMonoid_0_loop *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictMonoid_0 *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value] = dictMonoid_0_loop
	_ = dictMonoid_0
	// TAST (Let): mempty1_1_0 shape=App(Var) bindingType=(TypeVar m$scope52)
	mempty1_1_0 := Call_Data_Monoid_mempty(gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(dictMonoid_0)})
	_ = mempty1_1_0
	// TAST (Let): __local_var_2_1 shape=App(Other) bindingType=Any
	__local_var_2_1 := gopurs_runtime.Apply(dictMonoid_0.V0, gopurs_runtime.Value{})
	_ = __local_var_2_1
	return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
		var go__go_4_2_0 gopurs_runtime.Value
		_ = go__go_4_2_0
		var go__go_4_2_0_cell *gopurs_runtime.Value
		_ = go__go_4_2_0_cell
		// FALLBACK TCO: isLoop=false len=1
		go__go_4_2_0 = gopurs_runtime.Func(func(p_5 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t5 gopurs_runtime.Value
			{
				if (p_5.IntVal) <= (int64(0)) {
					__t5 = mempty1_1_0
					goto end_branch_5
				} else {

				}
			}
			{
				if (p_5.IntVal) == (int64(1)) {
					__t5 = x_3
					goto end_branch_5
				} else {

				}
			}
			{
				if (gopurs_runtime.IntMod(p_5.IntVal, int64(2))) == (int64(0)) {
					// TAST (Let): x_prime__6_4 shape=App(Other) bindingType=(TypeVar m$scope52)
					x_prime__6_4 := gopurs_runtime.Apply((*go__go_4_2_0_cell), gopurs_runtime.Int(gopurs_runtime.IntDiv(p_5.IntVal, int64(2))))
					_ = x_prime__6_4
					__t5 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "append"), x_prime__6_4, x_prime__6_4)
					goto end_branch_5
				} else {

				}
			}
			{
				// TAST (Let): x_prime__6_3 shape=App(Other) bindingType=(TypeVar m$scope52)
				x_prime__6_3 := gopurs_runtime.Apply((*go__go_4_2_0_cell), gopurs_runtime.Int(gopurs_runtime.IntDiv(p_5.IntVal, int64(2))))
				_ = x_prime__6_3
				__t5 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "append"), x_prime__6_3, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "append"), x_prime__6_3, x_3))
			}
		end_branch_5:
			return __t5
		})
		go__go_4_2_0_cell = &go__go_4_2_0
		return go__go_4_2_0
	})
}

func Call_Data_Monoid_guard(dictMonoid_0_loop *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictMonoid_0 *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value] = dictMonoid_0_loop
	_ = dictMonoid_0
	// TAST (Let): mempty1_1_0 shape=App(Var) bindingType=(TypeVar m$scope54)
	mempty1_1_0 := Call_Data_Monoid_mempty(gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(dictMonoid_0)})
	_ = mempty1_1_0
	return gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
		var __t1 gopurs_runtime.Value
		{
			if (v_2.IntVal) != (0) {
				__t1 = v1_3
				goto end_branch_1
			} else {

			}
		}
		{
			__t1 = mempty1_1_0
		}
	end_branch_1:
		return __t1
	})
}

func Rebox_Data_Monoid_1625289059_4179793454(in *Constructor_Data_Semigroup_Semigroup[uint32]) *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Monoid_1950344881_1201789390(in *Constructor_Data_Monoid_Monoid[string]) *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]{}
	out.V0 = in.V0
	out.V1 = gopurs_runtime.Str(in.V1)
	return out
}

func Rebox_Data_Monoid_2443930307_1201789390(in *Constructor_Data_Monoid_Monoid[uint32]) *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]{}
	out.V0 = in.V0
	out.V1 = gopurs_runtime.Value{Type: 9, IntVal: int64(in.V1), UnsafePtr: nil}
	return out
}

func Rebox_Data_Monoid_2860751713_1201789390(in *Constructor_Data_Monoid_Monoid[struct {
	dotAll     bool
	global     bool
	ignoreCase bool
	multiline  bool
	sticky     bool
	unicode    bool
	/*gopurs:record-v1:6:0064006f00740041006c006c,0067006c006f00620061006c,00690067006e006f007200650043006100730065,006d0075006c00740069006c0069006e0065,0073007400690063006b0079,0075006e00690063006f00640065*/
}]) *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]{}
	out.V0 = in.V0
	out.V1 = func() gopurs_runtime.Value {
		orig := in.V1
		_ = orig
		return gopurs_runtime.RecordDict([]string{"dotAll", "global", "ignoreCase", "multiline", "sticky", "unicode"}, []gopurs_runtime.Value{gopurs_runtime.Bool(orig.dotAll), gopurs_runtime.Bool(orig.global), gopurs_runtime.Bool(orig.ignoreCase), gopurs_runtime.Bool(orig.multiline), gopurs_runtime.Bool(orig.sticky), gopurs_runtime.Bool(orig.unicode)})
	}()
	return out
}

func Rebox_Data_Monoid_321927638_1201789390(in *Constructor_Data_Monoid_Monoid[float64]) *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]{}
	out.V0 = in.V0
	out.V1 = gopurs_runtime.Float(in.V1)
	return out
}

func Rebox_Data_Monoid_4179793454_1625289059(in *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]) *Constructor_Data_Semigroup_Semigroup[uint32] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Semigroup_Semigroup[uint32])(unsafe.Pointer(in))
}

func Rebox_Data_Monoid_4179793454_443971153(in *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]) *Constructor_Data_Semigroup_Semigroup[string] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Semigroup_Semigroup[string])(unsafe.Pointer(in))
}

func Rebox_Data_Monoid_443971153_4179793454(in *Constructor_Data_Semigroup_Semigroup[string]) *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Monoid_67307368_1201789390(in *Constructor_Data_Monoid_Monoid[[]gopurs_runtime.Value]) *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]{}
	out.V0 = in.V0
	out.V1 = gopurs_runtime.Array(in.V1)
	return out
}
