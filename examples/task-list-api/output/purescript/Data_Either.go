package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Either_go__append gopurs_runtime.Value
var once_Data_Either_go__append sync.Once

func Get_Data_Either_go__append() gopurs_runtime.Value {
	once_Data_Either_go__append.Do(func() {
		cache_Data_Either_go__append = Get_Data_Semigroup_append__1694605992()
	})
	return cache_Data_Either_go__append
}

var cache_Data_Either_compose gopurs_runtime.Value
var once_Data_Either_compose sync.Once

func Get_Data_Either_compose() gopurs_runtime.Value {
	once_Data_Either_compose.Do(func() {
		cache_Data_Either_compose = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Either_compose
}

var cache_Data_Either_Left gopurs_runtime.Value
var once_Data_Either_Left sync.Once

func Get_Data_Either_Left() gopurs_runtime.Value {
	once_Data_Either_Left.Do(func() {
		cache_Data_Either_Left = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer((&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, value0}))}
		})
	})
	return cache_Data_Either_Left
}

var cache_Data_Either_Left__2610938387 gopurs_runtime.Value
var once_Data_Either_Left__2610938387 sync.Once

func Get_Data_Either_Left__2610938387() gopurs_runtime.Value {
	once_Data_Either_Left__2610938387.Do(func() {
		cache_Data_Either_Left__2610938387 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Either_Left__2610938387(__eta_norm_0_0_box)
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})
	})
	return cache_Data_Either_Left__2610938387
}

var cache_Data_Either_Left__4153321788 gopurs_runtime.Value
var once_Data_Either_Left__4153321788 sync.Once

func Get_Data_Either_Left__4153321788() gopurs_runtime.Value {
	once_Data_Either_Left__4153321788.Do(func() {
		cache_Data_Either_Left__4153321788 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Either_Left__4153321788(__eta_norm_0_0_box)
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})
	})
	return cache_Data_Either_Left__4153321788
}

var cache_Data_Either_Left__3953938069 gopurs_runtime.Value
var once_Data_Either_Left__3953938069 sync.Once

func Get_Data_Either_Left__3953938069() gopurs_runtime.Value {
	once_Data_Either_Left__3953938069.Do(func() {
		cache_Data_Either_Left__3953938069 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Either_Left__3953938069(__eta_norm_0_0_box)
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})
	})
	return cache_Data_Either_Left__3953938069
}

var cache_Data_Either_Left__1641886679 gopurs_runtime.Value
var once_Data_Either_Left__1641886679 sync.Once

func Get_Data_Either_Left__1641886679() gopurs_runtime.Value {
	once_Data_Either_Left__1641886679.Do(func() {
		cache_Data_Either_Left__1641886679 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Either_Left__1641886679(__eta_norm_0_0_box)
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})
	})
	return cache_Data_Either_Left__1641886679
}

var cache_Data_Either_Left__2912811707 gopurs_runtime.Value
var once_Data_Either_Left__2912811707 sync.Once

func Get_Data_Either_Left__2912811707() gopurs_runtime.Value {
	once_Data_Either_Left__2912811707.Do(func() {
		cache_Data_Either_Left__2912811707 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Either_Left__2912811707(__eta_norm_0_0_box)
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: func() gopurs_runtime.Value {
						orig := _v.V1
						_ = orig
						return gopurs_runtime.RecordDict2("head", "tail", orig.head, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(orig.tail)})
					}()})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})
	})
	return cache_Data_Either_Left__2912811707
}

var cache_Data_Either_Left__1871912552 gopurs_runtime.Value
var once_Data_Either_Left__1871912552 sync.Once

func Get_Data_Either_Left__1871912552() gopurs_runtime.Value {
	once_Data_Either_Left__1871912552.Do(func() {
		cache_Data_Either_Left__1871912552 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Either_Left__1871912552(__eta_norm_0_0_box)
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: func() gopurs_runtime.Value {
						orig := _v.V1
						_ = orig
						return gopurs_runtime.RecordDict2("head", "tail", orig.head, gopurs_runtime.Array(orig.tail))
					}()})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})
	})
	return cache_Data_Either_Left__1871912552
}

var cache_Data_Either_Left__3194200049 gopurs_runtime.Value
var once_Data_Either_Left__3194200049 sync.Once

func Get_Data_Either_Left__3194200049() gopurs_runtime.Value {
	once_Data_Either_Left__3194200049.Do(func() {
		cache_Data_Either_Left__3194200049 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Either_Left__3194200049(__eta_norm_0_0_box)
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})
	})
	return cache_Data_Either_Left__3194200049
}

var cache_Data_Either_Left__4151004176 gopurs_runtime.Value
var once_Data_Either_Left__4151004176 sync.Once

func Get_Data_Either_Left__4151004176() gopurs_runtime.Value {
	once_Data_Either_Left__4151004176.Do(func() {
		cache_Data_Either_Left__4151004176 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Either_Left__4151004176(Rebox_Data_Either_1293498952_3123684004(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](__eta_norm_0_0_box)))
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})
	})
	return cache_Data_Either_Left__4151004176
}

var cache_Data_Either_Left__3420637932 gopurs_runtime.Value
var once_Data_Either_Left__3420637932 sync.Once

func Get_Data_Either_Left__3420637932() gopurs_runtime.Value {
	once_Data_Either_Left__3420637932.Do(func() {
		cache_Data_Either_Left__3420637932 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Either_Left__3420637932(__eta_norm_0_0_box)
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})
	})
	return cache_Data_Either_Left__3420637932
}

var cache_Data_Either_Left__2132796443 gopurs_runtime.Value
var once_Data_Either_Left__2132796443 sync.Once

func Get_Data_Either_Left__2132796443() gopurs_runtime.Value {
	once_Data_Either_Left__2132796443.Do(func() {
		cache_Data_Either_Left__2132796443 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Either_Left__2132796443(__eta_norm_0_0_box)
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: func() gopurs_runtime.Value {
						orig := _v.V1
						_ = orig
						return gopurs_runtime.RecordDict2("buffers", "readagain", gopurs_runtime.Array(orig.buffers), gopurs_runtime.Bool(orig.readagain))
					}()})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})
	})
	return cache_Data_Either_Left__2132796443
}

var cache_Data_Either_Left__4204059810 gopurs_runtime.Value
var once_Data_Either_Left__4204059810 sync.Once

func Get_Data_Either_Left__4204059810() gopurs_runtime.Value {
	once_Data_Either_Left__4204059810.Do(func() {
		cache_Data_Either_Left__4204059810 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Either_Left__4204059810(__eta_norm_0_0_box)
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})
	})
	return cache_Data_Either_Left__4204059810
}

var cache_Data_Either_Right gopurs_runtime.Value
var once_Data_Either_Right sync.Once

func Get_Data_Either_Right() gopurs_runtime.Value {
	once_Data_Either_Right.Do(func() {
		cache_Data_Either_Right = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer((&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, value0}))}
		})
	})
	return cache_Data_Either_Right
}

var cache_Data_Either_Right__1986130775 gopurs_runtime.Value
var once_Data_Either_Right__1986130775 sync.Once

func Get_Data_Either_Right__1986130775() gopurs_runtime.Value {
	once_Data_Either_Right__1986130775.Do(func() {
		cache_Data_Either_Right__1986130775 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Either_Right__1986130775(__eta_norm_0_0_box)
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})
	})
	return cache_Data_Either_Right__1986130775
}

var cache_Data_Either_Right__1001643778 gopurs_runtime.Value
var once_Data_Either_Right__1001643778 sync.Once

func Get_Data_Either_Right__1001643778() gopurs_runtime.Value {
	once_Data_Either_Right__1001643778.Do(func() {
		cache_Data_Either_Right__1001643778 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Either_Right__1001643778(__eta_norm_0_0_box)
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})
	})
	return cache_Data_Either_Right__1001643778
}

var cache_Data_Either_Right__489929074 gopurs_runtime.Value
var once_Data_Either_Right__489929074 sync.Once

func Get_Data_Either_Right__489929074() gopurs_runtime.Value {
	once_Data_Either_Right__489929074.Do(func() {
		cache_Data_Either_Right__489929074 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Either_Right__489929074(*((*[]gopurs_runtime.Value)(__eta_norm_0_0_box.UnsafePtr)))
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})
	})
	return cache_Data_Either_Right__489929074
}

var cache_Data_Either_Right__558618578 gopurs_runtime.Value
var once_Data_Either_Right__558618578 sync.Once

func Get_Data_Either_Right__558618578() gopurs_runtime.Value {
	once_Data_Either_Right__558618578.Do(func() {
		cache_Data_Either_Right__558618578 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Either_Right__558618578(func() struct {
					buffers   []gopurs_runtime.Value
					readagain bool
					/*gopurs:record-v1:2:0062007500660066006500720073,00720065006100640061006700610069006e*/
				} {
					orig := __eta_norm_0_0_box
					_ = orig
					clone := struct {
						buffers   []gopurs_runtime.Value
						readagain bool
						/*gopurs:record-v1:2:0062007500660066006500720073,00720065006100640061006700610069006e*/
					}{}
					clone.buffers = *((*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "buffers").UnsafePtr))
					clone.readagain = (gopurs_runtime.RecordGet(orig, "readagain").IntVal) != (0)
					return clone
				}())
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: func() gopurs_runtime.Value {
						orig := _v.V1
						_ = orig
						return gopurs_runtime.RecordDict2("buffers", "readagain", gopurs_runtime.Array(orig.buffers), gopurs_runtime.Bool(orig.readagain))
					}()})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})
	})
	return cache_Data_Either_Right__558618578
}

var cache_Data_Either_Right__3870859976 gopurs_runtime.Value
var once_Data_Either_Right__3870859976 sync.Once

func Get_Data_Either_Right__3870859976() gopurs_runtime.Value {
	once_Data_Either_Right__3870859976.Do(func() {
		cache_Data_Either_Right__3870859976 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Either_Right__3870859976(__eta_norm_0_unused_0_box)
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})
	})
	return cache_Data_Either_Right__3870859976
}

var cache_Data_Either_Right__112808109 gopurs_runtime.Value
var once_Data_Either_Right__112808109 sync.Once

func Get_Data_Either_Right__112808109() gopurs_runtime.Value {
	once_Data_Either_Right__112808109.Do(func() {
		cache_Data_Either_Right__112808109 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Either_Right__112808109(__eta_norm_0_unused_0_box)
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})
	})
	return cache_Data_Either_Right__112808109
}

var cache_Data_Either_showEither gopurs_runtime.Value
var once_Data_Either_showEither sync.Once

func Get_Data_Either_showEither() gopurs_runtime.Value {
	once_Data_Either_showEither.Do(func() {
		cache_Data_Either_showEither = gopurs_runtime.Func2(func(dictShow_0_box gopurs_runtime.Value, dictShow1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Either_showEither(dictShow_0_box, dictShow1_1_box)
		})
	})
	return cache_Data_Either_showEither
}

var cache_Data_Either_note_prime_ gopurs_runtime.Value
var once_Data_Either_note_prime_ sync.Once

func Get_Data_Either_note_prime_() gopurs_runtime.Value {
	once_Data_Either_note_prime_.Do(func() {
		cache_Data_Either_note_prime_ = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Either_note_prime_(f_0_box)
		})
	})
	return cache_Data_Either_note_prime_
}

var cache_Data_Either_note gopurs_runtime.Value
var once_Data_Either_note sync.Once

func Get_Data_Either_note() gopurs_runtime.Value {
	once_Data_Either_note.Do(func() {
		cache_Data_Either_note = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Either_note(a_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v2_1_box))
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})
	})
	return cache_Data_Either_note
}

var cache_Data_Either_note__4111599802 gopurs_runtime.Value
var once_Data_Either_note__4111599802 sync.Once

func Get_Data_Either_note__4111599802() gopurs_runtime.Value {
	once_Data_Either_note__4111599802.Do(func() {
		cache_Data_Either_note__4111599802 = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Either_note__4111599802(a_0_box)
		})
	})
	return cache_Data_Either_note__4111599802
}

var cache_Data_Either_note__3211399674 gopurs_runtime.Value
var once_Data_Either_note__3211399674 sync.Once

func Get_Data_Either_note__3211399674() gopurs_runtime.Value {
	once_Data_Either_note__3211399674.Do(func() {
		cache_Data_Either_note__3211399674 = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Either_note__3211399674(a_0_box)
		})
	})
	return cache_Data_Either_note__3211399674
}

var cache_Data_Either_note__2066121466 gopurs_runtime.Value
var once_Data_Either_note__2066121466 sync.Once

func Get_Data_Either_note__2066121466() gopurs_runtime.Value {
	once_Data_Either_note__2066121466.Do(func() {
		cache_Data_Either_note__2066121466 = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Either_note__2066121466(a_0_box)
		})
	})
	return cache_Data_Either_note__2066121466
}

var cache_Data_Either_note__271183866 gopurs_runtime.Value
var once_Data_Either_note__271183866 sync.Once

func Get_Data_Either_note__271183866() gopurs_runtime.Value {
	once_Data_Either_note__271183866.Do(func() {
		cache_Data_Either_note__271183866 = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Either_note__271183866(a_0_box)
		})
	})
	return cache_Data_Either_note__271183866
}

var cache_Data_Either_note__2268946106 gopurs_runtime.Value
var once_Data_Either_note__2268946106 sync.Once

func Get_Data_Either_note__2268946106() gopurs_runtime.Value {
	once_Data_Either_note__2268946106.Do(func() {
		cache_Data_Either_note__2268946106 = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Either_note__2268946106(a_0_box)
		})
	})
	return cache_Data_Either_note__2268946106
}

var cache_Data_Either_note__67430010 gopurs_runtime.Value
var once_Data_Either_note__67430010 sync.Once

func Get_Data_Either_note__67430010() gopurs_runtime.Value {
	once_Data_Either_note__67430010.Do(func() {
		cache_Data_Either_note__67430010 = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Either_note__67430010(a_0_box)
		})
	})
	return cache_Data_Either_note__67430010
}

var cache_Data_Either_note__357786010 gopurs_runtime.Value
var once_Data_Either_note__357786010 sync.Once

func Get_Data_Either_note__357786010() gopurs_runtime.Value {
	once_Data_Either_note__357786010.Do(func() {
		cache_Data_Either_note__357786010 = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Either_note__357786010(a_0_box)
		})
	})
	return cache_Data_Either_note__357786010
}

var cache_Data_Either_genericEither gopurs_runtime.Value
var once_Data_Either_genericEither sync.Once

func Get_Data_Either_genericEither() gopurs_runtime.Value {
	once_Data_Either_genericEither.Do(func() {
		cache_Data_Either_genericEither = gopurs_runtime.Value{Type: 9, IntVal: 1921946594, UnsafePtr: unsafe.Pointer((&Constructor_Data_Generic_Rep_Generic[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t0 gopurs_runtime.Value
			{
				if x_0.Type == 9 && x_0.IntVal == 3711209382 {
					__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3478632216, UnsafePtr: unsafe.Pointer((&Constructor_Data_Generic_Rep_Inl[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(x_0.UnsafePtr).V0}))}
					goto end_branch_0
				} else {

				}
			}
			{
				if x_0.Type == 9 && x_0.IntVal == 2465973597 {
					__t0 = gopurs_runtime.Value{Type: 9, IntVal: 492034566, UnsafePtr: unsafe.Pointer((&Constructor_Data_Generic_Rep_Inr[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(x_0.UnsafePtr).V0}))}
					goto end_branch_0
				} else {

				}
			}
			{
				__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_0:
			return __t0
		}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t1 struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
				V2 bool
			}
			{
				if x_0.Type == 9 && x_0.IntVal == 3478632216 {
					__t1 = struct {
						V0 gopurs_runtime.Value
						V1 gopurs_runtime.Value
						V2 bool
					}{(*Constructor_Data_Generic_Rep_Inl[gopurs_runtime.Value, gopurs_runtime.Value])(x_0.UnsafePtr).V0, gopurs_runtime.Value{}, false}
					goto end_branch_1
				} else {

				}
			}
			{
				if x_0.Type == 9 && x_0.IntVal == 492034566 {
					__t1 = struct {
						V0 gopurs_runtime.Value
						V1 gopurs_runtime.Value
						V2 bool
					}{gopurs_runtime.Value{}, (*Constructor_Data_Generic_Rep_Inr[gopurs_runtime.Value, gopurs_runtime.Value])(x_0.UnsafePtr).V0, true}
					goto end_branch_1
				} else {

				}
			}
			{
				__t1 = func() struct {
					V0 gopurs_runtime.Value
					V1 gopurs_runtime.Value
					V2 bool
				} { _v := func() gopurs_runtime.Value { panic("Failed pattern match") }(); if _v.Type == 9 && _v.IntVal == 2465973597 && _v.UnsafePtr != nil {
					return struct {
						V0 gopurs_runtime.Value
						V1 gopurs_runtime.Value
						V2 bool
					}{V0: gopurs_runtime.Value{}, V1: (*(*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V0, V2: true}
				}; return struct {
					V0 gopurs_runtime.Value
					V1 gopurs_runtime.Value
					V2 bool
				}{V0: (*(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: gopurs_runtime.Value{}, V2: false} }()
			}
		end_branch_1:
			return func() gopurs_runtime.Value {
				_v := __t1
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})}))}
	})
	return cache_Data_Either_genericEither
}

var cache_Data_Either_functorEither gopurs_runtime.Value
var once_Data_Either_functorEither sync.Once

func Get_Data_Either_functorEither() gopurs_runtime.Value {
	once_Data_Either_functorEither.Do(func() {
		cache_Data_Either_functorEither = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, m_1 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t0 struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
				V2 bool
			}
			{
				if m_1.Type == 9 && m_1.IntVal == 3711209382 {
					__t0 = struct {
						V0 gopurs_runtime.Value
						V1 gopurs_runtime.Value
						V2 bool
					}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V0, gopurs_runtime.Value{}, false}
					goto end_branch_0
				} else {

				}
			}
			{
				if m_1.Type == 9 && m_1.IntVal == 2465973597 {
					__t0 = struct {
						V0 gopurs_runtime.Value
						V1 gopurs_runtime.Value
						V2 bool
					}{gopurs_runtime.Value{}, gopurs_runtime.Apply(f_0, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V0), true}
					goto end_branch_0
				} else {

				}
			}
			{
				__t0 = func() struct {
					V0 gopurs_runtime.Value
					V1 gopurs_runtime.Value
					V2 bool
				} { _v := func() gopurs_runtime.Value { panic("Failed pattern match") }(); if _v.Type == 9 && _v.IntVal == 2465973597 && _v.UnsafePtr != nil {
					return struct {
						V0 gopurs_runtime.Value
						V1 gopurs_runtime.Value
						V2 bool
					}{V0: gopurs_runtime.Value{}, V1: (*(*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V0, V2: true}
				}; return struct {
					V0 gopurs_runtime.Value
					V1 gopurs_runtime.Value
					V2 bool
				}{V0: (*(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: gopurs_runtime.Value{}, V2: false} }()
			}
		end_branch_0:
			return func() gopurs_runtime.Value {
				_v := __t0
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})}))}
	})
	return cache_Data_Either_functorEither
}

var cache_Data_Either_go__map gopurs_runtime.Value
var once_Data_Either_go__map sync.Once

func Get_Data_Either_go__map() gopurs_runtime.Value {
	once_Data_Either_go__map.Do(func() {
		cache_Data_Either_go__map = gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map")
	})
	return cache_Data_Either_go__map
}

var cache_Data_Either_map1 gopurs_runtime.Value
var once_Data_Either_map1 sync.Once

func Get_Data_Either_map1() gopurs_runtime.Value {
	once_Data_Either_map1.Do(func() {
		cache_Data_Either_map1 = gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map")
	})
	return cache_Data_Either_map1
}

var cache_Data_Either_invariantEither gopurs_runtime.Value
var once_Data_Either_invariantEither sync.Once

func Get_Data_Either_invariantEither() gopurs_runtime.Value {
	once_Data_Either_invariantEither.Do(func() {
		cache_Data_Either_invariantEither = gopurs_runtime.Value{Type: 9, IntVal: 2396985522, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Invariant_Invariant[gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Functor_Invariant_imapF(), Get_Data_Either_functorEither())}))}
	})
	return cache_Data_Either_invariantEither
}

var cache_Data_Either_fromRight_prime_ gopurs_runtime.Value
var once_Data_Either_fromRight_prime_ sync.Once

func Get_Data_Either_fromRight_prime_() gopurs_runtime.Value {
	once_Data_Either_fromRight_prime_.Do(func() {
		cache_Data_Either_fromRight_prime_ = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Either_fromRight_prime_(v_0_box, v1_1_box)
		})
	})
	return cache_Data_Either_fromRight_prime_
}

var cache_Data_Either_fromRight gopurs_runtime.Value
var once_Data_Either_fromRight sync.Once

func Get_Data_Either_fromRight() gopurs_runtime.Value {
	once_Data_Either_fromRight.Do(func() {
		cache_Data_Either_fromRight = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Either_fromRight(v_0_box, v1_1_box)
		})
	})
	return cache_Data_Either_fromRight
}

var cache_Data_Either_fromLeft_prime_ gopurs_runtime.Value
var once_Data_Either_fromLeft_prime_ sync.Once

func Get_Data_Either_fromLeft_prime_() gopurs_runtime.Value {
	once_Data_Either_fromLeft_prime_.Do(func() {
		cache_Data_Either_fromLeft_prime_ = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Either_fromLeft_prime_(v_0_box, v1_1_box)
		})
	})
	return cache_Data_Either_fromLeft_prime_
}

var cache_Data_Either_fromLeft gopurs_runtime.Value
var once_Data_Either_fromLeft sync.Once

func Get_Data_Either_fromLeft() gopurs_runtime.Value {
	once_Data_Either_fromLeft.Do(func() {
		cache_Data_Either_fromLeft = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Either_fromLeft(v_0_box, v1_1_box)
		})
	})
	return cache_Data_Either_fromLeft
}

var cache_Data_Either_extendEither gopurs_runtime.Value
var once_Data_Either_extendEither sync.Once

func Get_Data_Either_extendEither() gopurs_runtime.Value {
	once_Data_Either_extendEither.Do(func() {
		cache_Data_Either_extendEither = gopurs_runtime.Value{Type: 9, IntVal: 3028639021, UnsafePtr: unsafe.Pointer((&Constructor_Control_Extend_Extend[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Get_Data_Either_functorEither()
		}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t0 struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
				V2 bool
			}
			{
				if v1_1.Type == 9 && v1_1.IntVal == 3711209382 {
					__t0 = struct {
						V0 gopurs_runtime.Value
						V1 gopurs_runtime.Value
						V2 bool
					}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V0, gopurs_runtime.Value{}, false}
					goto end_branch_0
				} else {

				}
			}
			{
				__t0 = struct {
					V0 gopurs_runtime.Value
					V1 gopurs_runtime.Value
					V2 bool
				}{gopurs_runtime.Value{}, gopurs_runtime.Apply(v_0, v1_1), true}
			}
		end_branch_0:
			return func() gopurs_runtime.Value {
				_v := __t0
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})}))}
	})
	return cache_Data_Either_extendEither
}

var cache_Data_Either_eqEither gopurs_runtime.Value
var once_Data_Either_eqEither sync.Once

func Get_Data_Either_eqEither() gopurs_runtime.Value {
	once_Data_Either_eqEither.Do(func() {
		cache_Data_Either_eqEither = gopurs_runtime.Func2(func(dictEq_0_box gopurs_runtime.Value, dictEq1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Either_eqEither(dictEq_0_box, dictEq1_1_box)
		})
	})
	return cache_Data_Either_eqEither
}

var cache_Data_Either_ordEither gopurs_runtime.Value
var once_Data_Either_ordEither sync.Once

func Get_Data_Either_ordEither() gopurs_runtime.Value {
	once_Data_Either_ordEither.Do(func() {
		cache_Data_Either_ordEither = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Either_ordEither(dictOrd_0_box)
		})
	})
	return cache_Data_Either_ordEither
}

var cache_Data_Either_eq1Either gopurs_runtime.Value
var once_Data_Either_eq1Either sync.Once

func Get_Data_Either_eq1Either() gopurs_runtime.Value {
	once_Data_Either_eq1Either.Do(func() {
		cache_Data_Either_eq1Either = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Either_eq1Either(dictEq_0_box)
		})
	})
	return cache_Data_Either_eq1Either
}

var cache_Data_Either_ord1Either gopurs_runtime.Value
var once_Data_Either_ord1Either sync.Once

func Get_Data_Either_ord1Either() gopurs_runtime.Value {
	once_Data_Either_ord1Either.Do(func() {
		cache_Data_Either_ord1Either = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Either_ord1Either(dictOrd_0_box)
		})
	})
	return cache_Data_Either_ord1Either
}

var cache_Data_Either_either gopurs_runtime.Value
var once_Data_Either_either sync.Once

func Get_Data_Either_either() gopurs_runtime.Value {
	once_Data_Either_either.Do(func() {
		cache_Data_Either_either = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Either_either(v_0_box, v1_1_box, v2_2_box)
		})
	})
	return cache_Data_Either_either
}

var cache_Data_Either_either__931018797 gopurs_runtime.Value
var once_Data_Either_either__931018797 sync.Once

func Get_Data_Either_either__931018797() gopurs_runtime.Value {
	once_Data_Either_either__931018797.Do(func() {
		cache_Data_Either_either__931018797 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Either_either__931018797(v_0_box, v1_1_box, v2_2_box)
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})
	})
	return cache_Data_Either_either__931018797
}

var cache_Data_Either_either__878214054 gopurs_runtime.Value
var once_Data_Either_either__878214054 sync.Once

func Get_Data_Either_either__878214054() gopurs_runtime.Value {
	once_Data_Either_either__878214054.Do(func() {
		cache_Data_Either_either__878214054 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Either_either__878214054(v_0_box, v1_1_box, v2_2_box)
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})
	})
	return cache_Data_Either_either__878214054
}

var cache_Data_Either_either__3820936092 gopurs_runtime.Value
var once_Data_Either_either__3820936092 sync.Once

func Get_Data_Either_either__3820936092() gopurs_runtime.Value {
	once_Data_Either_either__3820936092.Do(func() {
		cache_Data_Either_either__3820936092 = gopurs_runtime.Func3(func(v_unused_0_box gopurs_runtime.Value, v1_unused_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Either_either__3820936092(v_unused_0_box, v1_unused_1_box, v2_2_box)
		})
	})
	return cache_Data_Either_either__3820936092
}

var cache_Data_Either_hush gopurs_runtime.Value
var once_Data_Either_hush sync.Once

func Get_Data_Either_hush() gopurs_runtime.Value {
	once_Data_Either_hush.Do(func() {
		cache_Data_Either_hush = gopurs_runtime.Func(func(v2_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Either_hush(v2_0_box)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_Either_hush
}

var cache_Data_Either_isLeft gopurs_runtime.Value
var once_Data_Either_isLeft sync.Once

func Get_Data_Either_isLeft() gopurs_runtime.Value {
	once_Data_Either_isLeft.Do(func() {
		cache_Data_Either_isLeft = gopurs_runtime.Func(func(v2_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(Call_Data_Either_isLeft(v2_0_box))
		})
	})
	return cache_Data_Either_isLeft
}

var cache_Data_Either_isRight gopurs_runtime.Value
var once_Data_Either_isRight sync.Once

func Get_Data_Either_isRight() gopurs_runtime.Value {
	once_Data_Either_isRight.Do(func() {
		cache_Data_Either_isRight = gopurs_runtime.Func(func(v2_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(Call_Data_Either_isRight(v2_0_box))
		})
	})
	return cache_Data_Either_isRight
}

var cache_Data_Either_choose gopurs_runtime.Value
var once_Data_Either_choose sync.Once

func Get_Data_Either_choose() gopurs_runtime.Value {
	once_Data_Either_choose.Do(func() {
		cache_Data_Either_choose = gopurs_runtime.Func(func(dictAlt_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Either_choose(gopurs_runtime.CoerceToStruct[Constructor_Control_Alt_Alt[gopurs_runtime.Value]](dictAlt_0_box))
		})
	})
	return cache_Data_Either_choose
}

var cache_Data_Either_boundedEither gopurs_runtime.Value
var once_Data_Either_boundedEither sync.Once

func Get_Data_Either_boundedEither() gopurs_runtime.Value {
	once_Data_Either_boundedEither.Do(func() {
		cache_Data_Either_boundedEither = gopurs_runtime.Func(func(dictBounded_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Either_boundedEither(dictBounded_0_box)
		})
	})
	return cache_Data_Either_boundedEither
}

var cache_Data_Either_blush gopurs_runtime.Value
var once_Data_Either_blush sync.Once

func Get_Data_Either_blush() gopurs_runtime.Value {
	once_Data_Either_blush.Do(func() {
		cache_Data_Either_blush = gopurs_runtime.Func(func(v2_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Either_blush(v2_0_box)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_Either_blush
}

var cache_Data_Either_applyEither gopurs_runtime.Value
var once_Data_Either_applyEither sync.Once

func Get_Data_Either_applyEither() gopurs_runtime.Value {
	once_Data_Either_applyEither.Do(func() {
		cache_Data_Either_applyEither = gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Get_Data_Either_functorEither()
		}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t1 struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
				V2 bool
			}
			{
				if v_0.Type == 9 && v_0.IntVal == 3711209382 {
					__t1 = struct {
						V0 gopurs_runtime.Value
						V1 gopurs_runtime.Value
						V2 bool
					}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V0, gopurs_runtime.Value{}, false}
					goto end_branch_1
				} else {

				}
			}
			{
				if v_0.Type == 9 && v_0.IntVal == 2465973597 {
					var __t0 struct {
						V0 gopurs_runtime.Value
						V1 gopurs_runtime.Value
						V2 bool
					}
					{
						if v1_1.Type == 9 && v1_1.IntVal == 3711209382 {
							__t0 = struct {
								V0 gopurs_runtime.Value
								V1 gopurs_runtime.Value
								V2 bool
							}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V0, gopurs_runtime.Value{}, false}
							goto end_branch_0
						} else {

						}
					}
					{
						if v1_1.Type == 9 && v1_1.IntVal == 2465973597 {
							__t0 = struct {
								V0 gopurs_runtime.Value
								V1 gopurs_runtime.Value
								V2 bool
							}{gopurs_runtime.Value{}, gopurs_runtime.Apply((*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V0, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V0), true}
							goto end_branch_0
						} else {

						}
					}
					{
						__t0 = func() struct {
							V0 gopurs_runtime.Value
							V1 gopurs_runtime.Value
							V2 bool
						} { _v := func() gopurs_runtime.Value { panic("Failed pattern match") }(); if _v.Type == 9 && _v.IntVal == 2465973597 && _v.UnsafePtr != nil {
							return struct {
								V0 gopurs_runtime.Value
								V1 gopurs_runtime.Value
								V2 bool
							}{V0: gopurs_runtime.Value{}, V1: (*(*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V0, V2: true}
						}; return struct {
							V0 gopurs_runtime.Value
							V1 gopurs_runtime.Value
							V2 bool
						}{V0: (*(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: gopurs_runtime.Value{}, V2: false} }()
					}
				end_branch_0:
					__t1 = __t0
					goto end_branch_1
				} else {

				}
			}
			{
				__t1 = func() struct {
					V0 gopurs_runtime.Value
					V1 gopurs_runtime.Value
					V2 bool
				} { _v := func() gopurs_runtime.Value { panic("Failed pattern match") }(); if _v.Type == 9 && _v.IntVal == 2465973597 && _v.UnsafePtr != nil {
					return struct {
						V0 gopurs_runtime.Value
						V1 gopurs_runtime.Value
						V2 bool
					}{V0: gopurs_runtime.Value{}, V1: (*(*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V0, V2: true}
				}; return struct {
					V0 gopurs_runtime.Value
					V1 gopurs_runtime.Value
					V2 bool
				}{V0: (*(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: gopurs_runtime.Value{}, V2: false} }()
			}
		end_branch_1:
			return func() gopurs_runtime.Value {
				_v := __t1
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})}))}
	})
	return cache_Data_Either_applyEither
}

var cache_Data_Either_apply gopurs_runtime.Value
var once_Data_Either_apply sync.Once

func Get_Data_Either_apply() gopurs_runtime.Value {
	once_Data_Either_apply.Do(func() {
		cache_Data_Either_apply = gopurs_runtime.RecordGet(Get_Data_Either_applyEither(), "apply")
	})
	return cache_Data_Either_apply
}

var cache_Data_Either_bindEither gopurs_runtime.Value
var once_Data_Either_bindEither sync.Once

func Get_Data_Either_bindEither() gopurs_runtime.Value {
	once_Data_Either_bindEither.Do(func() {
		cache_Data_Either_bindEither = gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Get_Data_Either_applyEither()
		}), gopurs_runtime.Func(func(v2_0 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t2 gopurs_runtime.Value
			{
				if v2_0.Type == 9 && v2_0.IntVal == 3711209382 {
					// TAST (Let): __local_var_1_0 shape=Other bindingType=(TypeVar e$scope161)
					__local_var_1_0 := (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_0.UnsafePtr).V0
					_ = __local_var_1_0
					__t2 = gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
						return func() gopurs_runtime.Value {
							_v := struct {
								V0 gopurs_runtime.Value
								V1 gopurs_runtime.Value
								V2 bool
							}{__local_var_1_0, gopurs_runtime.Value{}, false}
							if _v.V2 {
								return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
							}
							return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
						}()
					})
					goto end_branch_2
				} else {

				}
			}
			{
				if v2_0.Type == 9 && v2_0.IntVal == 2465973597 {
					// TAST (Let): __local_var_1_1 shape=Other bindingType=(TypeVar b$scope119)
					__local_var_1_1 := (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_0.UnsafePtr).V0
					_ = __local_var_1_1
					__t2 = gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Apply(f_2, __local_var_1_1)
					})
					goto end_branch_2
				} else {

				}
			}
			{
				__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_2:
			return __t2
		})}))}
	})
	return cache_Data_Either_bindEither
}

var cache_Data_Either_semigroupEither gopurs_runtime.Value
var once_Data_Either_semigroupEither sync.Once

func Get_Data_Either_semigroupEither() gopurs_runtime.Value {
	once_Data_Either_semigroupEither.Do(func() {
		cache_Data_Either_semigroupEither = gopurs_runtime.Func(func(dictSemigroup_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Either_semigroupEither(dictSemigroup_0_box)
		})
	})
	return cache_Data_Either_semigroupEither
}

var cache_Data_Either_applicativeEither gopurs_runtime.Value
var once_Data_Either_applicativeEither sync.Once

func Get_Data_Either_applicativeEither() gopurs_runtime.Value {
	once_Data_Either_applicativeEither.Do(func() {
		cache_Data_Either_applicativeEither = gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Get_Data_Either_applyEither()
		}), Get_Data_Either_Right()}))}
	})
	return cache_Data_Either_applicativeEither
}

var cache_Data_Either_monadEither gopurs_runtime.Value
var once_Data_Either_monadEither sync.Once

func Get_Data_Either_monadEither() gopurs_runtime.Value {
	once_Data_Either_monadEither.Do(func() {
		cache_Data_Either_monadEither = gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Get_Data_Either_applicativeEither()
		}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Get_Data_Either_bindEither()
		})}))}
	})
	return cache_Data_Either_monadEither
}

var cache_Data_Either_altEither gopurs_runtime.Value
var once_Data_Either_altEither sync.Once

func Get_Data_Either_altEither() gopurs_runtime.Value {
	once_Data_Either_altEither.Do(func() {
		cache_Data_Either_altEither = gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer((&Constructor_Control_Alt_Alt[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Get_Data_Either_functorEither()
		}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t0 gopurs_runtime.Value
			{
				if v_0.Type == 9 && v_0.IntVal == 3711209382 {
					__t0 = v1_1
					goto end_branch_0
				} else {

				}
			}
			{
				__t0 = v_0
			}
		end_branch_0:
			return __t0
		})}))}
	})
	return cache_Data_Either_altEither
}

type Constructor_Data_Either_Left[T_a any, T_b any] struct {
	Rc uint32
	V0 T_a
}

type Constructor_Data_Either_Right[T_a any, T_b any] struct {
	Rc uint32
	V0 T_b
}

func Call_Data_Either_Left__2610938387(__eta_norm_0_0_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 bool
} {
Left__2610938387:
	for {
		if false {
			continue Left__2610938387
		}
		var __eta_norm_0_0 gopurs_runtime.Value = __eta_norm_0_0_loop
		_ = __eta_norm_0_0
		return struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
			V2 bool
		}{__eta_norm_0_0, gopurs_runtime.Value{}, false}
	}
}

func Call_Data_Either_Left__4153321788(__eta_norm_0_0_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 bool
} {
Left__4153321788:
	for {
		if false {
			continue Left__4153321788
		}
		var __eta_norm_0_0 gopurs_runtime.Value = __eta_norm_0_0_loop
		_ = __eta_norm_0_0
		return struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
			V2 bool
		}{__eta_norm_0_0, gopurs_runtime.Value{}, false}
	}
}

func Call_Data_Either_Left__3953938069(__eta_norm_0_0_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 bool
} {
Left__3953938069:
	for {
		if false {
			continue Left__3953938069
		}
		var __eta_norm_0_0 gopurs_runtime.Value = __eta_norm_0_0_loop
		_ = __eta_norm_0_0
		return struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
			V2 bool
		}{__eta_norm_0_0, gopurs_runtime.Value{}, false}
	}
}

func Call_Data_Either_Left__1641886679(__eta_norm_0_0_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 bool
} {
Left__1641886679:
	for {
		if false {
			continue Left__1641886679
		}
		var __eta_norm_0_0 gopurs_runtime.Value = __eta_norm_0_0_loop
		_ = __eta_norm_0_0
		return struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
			V2 bool
		}{__eta_norm_0_0, gopurs_runtime.Value{}, false}
	}
}

func Call_Data_Either_Left__2912811707(__eta_norm_0_0_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 struct {
		head gopurs_runtime.Value
		tail *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
		/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
	}
	V2 bool
} {
Left__2912811707:
	for {
		if false {
			continue Left__2912811707
		}
		var __eta_norm_0_0 gopurs_runtime.Value = __eta_norm_0_0_loop
		_ = __eta_norm_0_0
		return struct {
			V0 gopurs_runtime.Value
			V1 struct {
				head gopurs_runtime.Value
				tail *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
				/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
			}
			V2 bool
		}{__eta_norm_0_0, struct {
			head gopurs_runtime.Value
			tail *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
			/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
		}{}, false}
	}
}

func Call_Data_Either_Left__1871912552(__eta_norm_0_0_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 struct {
		head gopurs_runtime.Value
		tail []gopurs_runtime.Value
		/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
	}
	V2 bool
} {
Left__1871912552:
	for {
		if false {
			continue Left__1871912552
		}
		var __eta_norm_0_0 gopurs_runtime.Value = __eta_norm_0_0_loop
		_ = __eta_norm_0_0
		return struct {
			V0 gopurs_runtime.Value
			V1 struct {
				head gopurs_runtime.Value
				tail []gopurs_runtime.Value
				/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
			}
			V2 bool
		}{__eta_norm_0_0, struct {
			head gopurs_runtime.Value
			tail []gopurs_runtime.Value
			/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
		}{}, false}
	}
}

func Call_Data_Either_Left__3194200049(__eta_norm_0_0_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 bool
} {
Left__3194200049:
	for {
		if false {
			continue Left__3194200049
		}
		var __eta_norm_0_0 gopurs_runtime.Value = __eta_norm_0_0_loop
		_ = __eta_norm_0_0
		return struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
			V2 bool
		}{__eta_norm_0_0, gopurs_runtime.Value{}, false}
	}
}

func Call_Data_Either_Left__4151004176(__eta_norm_0_0_loop *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 bool
} {
Left__4151004176:
	for {
		if false {
			continue Left__4151004176
		}
		var __eta_norm_0_0 *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = __eta_norm_0_0_loop
		_ = __eta_norm_0_0
		return struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
			V2 bool
		}{gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Rebox_Data_Either_3123684004_1293498952(__eta_norm_0_0))}, gopurs_runtime.Value{}, false}
	}
}

func Call_Data_Either_Left__3420637932(__eta_norm_0_0_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 bool
} {
Left__3420637932:
	for {
		if false {
			continue Left__3420637932
		}
		var __eta_norm_0_0 gopurs_runtime.Value = __eta_norm_0_0_loop
		_ = __eta_norm_0_0
		return struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
			V2 bool
		}{__eta_norm_0_0, gopurs_runtime.Value{}, false}
	}
}

func Call_Data_Either_Left__2132796443(__eta_norm_0_0_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 struct {
		buffers   []gopurs_runtime.Value
		readagain bool
		/*gopurs:record-v1:2:0062007500660066006500720073,00720065006100640061006700610069006e*/
	}
	V2 bool
} {
Left__2132796443:
	for {
		if false {
			continue Left__2132796443
		}
		var __eta_norm_0_0 gopurs_runtime.Value = __eta_norm_0_0_loop
		_ = __eta_norm_0_0
		return struct {
			V0 gopurs_runtime.Value
			V1 struct {
				buffers   []gopurs_runtime.Value
				readagain bool
				/*gopurs:record-v1:2:0062007500660066006500720073,00720065006100640061006700610069006e*/
			}
			V2 bool
		}{__eta_norm_0_0, struct {
			buffers   []gopurs_runtime.Value
			readagain bool
			/*gopurs:record-v1:2:0062007500660066006500720073,00720065006100640061006700610069006e*/
		}{}, false}
	}
}

func Call_Data_Either_Left__4204059810(__eta_norm_0_0_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 bool
} {
Left__4204059810:
	for {
		if false {
			continue Left__4204059810
		}
		var __eta_norm_0_0 gopurs_runtime.Value = __eta_norm_0_0_loop
		_ = __eta_norm_0_0
		return struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
			V2 bool
		}{__eta_norm_0_0, gopurs_runtime.Value{}, false}
	}
}

func Call_Data_Either_Right__1986130775(__eta_norm_0_0_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 bool
} {
Right__1986130775:
	for {
		if false {
			continue Right__1986130775
		}
		var __eta_norm_0_0 gopurs_runtime.Value = __eta_norm_0_0_loop
		_ = __eta_norm_0_0
		return struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
			V2 bool
		}{gopurs_runtime.Value{}, __eta_norm_0_0, true}
	}
}

func Call_Data_Either_Right__1001643778(__eta_norm_0_0_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 bool
} {
Right__1001643778:
	for {
		if false {
			continue Right__1001643778
		}
		var __eta_norm_0_0 gopurs_runtime.Value = __eta_norm_0_0_loop
		_ = __eta_norm_0_0
		return struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
			V2 bool
		}{gopurs_runtime.Value{}, __eta_norm_0_0, true}
	}
}

func Call_Data_Either_Right__489929074(__eta_norm_0_0_loop []gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 bool
} {
Right__489929074:
	for {
		if false {
			continue Right__489929074
		}
		var __eta_norm_0_0 []gopurs_runtime.Value = __eta_norm_0_0_loop
		_ = __eta_norm_0_0
		return struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
			V2 bool
		}{gopurs_runtime.Value{}, gopurs_runtime.Array(__eta_norm_0_0), true}
	}
}

func Call_Data_Either_Right__558618578(__eta_norm_0_0_loop struct {
	buffers   []gopurs_runtime.Value
	readagain bool
	/*gopurs:record-v1:2:0062007500660066006500720073,00720065006100640061006700610069006e*/
}) struct {
	V0 gopurs_runtime.Value
	V1 struct {
		buffers   []gopurs_runtime.Value
		readagain bool
		/*gopurs:record-v1:2:0062007500660066006500720073,00720065006100640061006700610069006e*/
	}
	V2 bool
} {
Right__558618578:
	for {
		if false {
			continue Right__558618578
		}
		var __eta_norm_0_0 struct {
			buffers   []gopurs_runtime.Value
			readagain bool
			/*gopurs:record-v1:2:0062007500660066006500720073,00720065006100640061006700610069006e*/
		} = __eta_norm_0_0_loop
		_ = __eta_norm_0_0
		return struct {
			V0 gopurs_runtime.Value
			V1 struct {
				buffers   []gopurs_runtime.Value
				readagain bool
				/*gopurs:record-v1:2:0062007500660066006500720073,00720065006100640061006700610069006e*/
			}
			V2 bool
		}{gopurs_runtime.Value{}, __eta_norm_0_0, true}
	}
}

func Call_Data_Either_Right__3870859976(__eta_norm_0_unused_0_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 bool
} {
Right__3870859976:
	for {
		if false {
			continue Right__3870859976
		}
		var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
		_ = __eta_norm_0_unused_0
		return struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
			V2 bool
		}{gopurs_runtime.Value{}, Get_Data_Unit_unit(), true}
	}
}

func Call_Data_Either_Right__112808109(__eta_norm_0_unused_0_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 bool
} {
Right__112808109:
	for {
		if false {
			continue Right__112808109
		}
		var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
		_ = __eta_norm_0_unused_0
		return struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
			V2 bool
		}{gopurs_runtime.Value{}, Get_Data_Unit_unit(), true}
	}
}

func Call_Data_Either_showEither(dictShow_0_loop gopurs_runtime.Value, dictShow1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
	_ = dictShow_0
	var dictShow1_1 gopurs_runtime.Value = dictShow1_1_loop
	_ = dictShow1_1
	return gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer((&Constructor_Data_Show_Show[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
		var __t0 gopurs_runtime.Value
		{
			if v_2.Type == 9 && v_2.IntVal == 3711209382 {
				__t0 = gopurs_runtime.Str(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString("(Left ", gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0).StrVal()), ")"))
				goto end_branch_0
			} else {

			}
		}
		{
			if v_2.Type == 9 && v_2.IntVal == 2465973597 {
				__t0 = gopurs_runtime.Str(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString("(Right ", gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow1_1, "show"), (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0).StrVal()), ")"))
				goto end_branch_0
			} else {

			}
		}
		{
			__t0 = gopurs_runtime.Str(func() string { panic("Failed pattern match") }())
		}
	end_branch_0:
		return __t0
	})}))}
}

func Call_Data_Either_note_prime_(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	return gopurs_runtime.Apply2(Get_Data_Maybe_maybe_prime_(), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Either_Left(), f_0), Get_Data_Either_Right())
}

func Call_Data_Either_note(a_0_loop gopurs_runtime.Value, v2_1_loop *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 bool
} {
	var a_0 gopurs_runtime.Value = a_0_loop
	_ = a_0
	var v2_1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = v2_1_loop
	_ = v2_1
	var __t0 gopurs_runtime.Value
	{
		if v2_1 == nil {
			__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer((&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_0}))}
			goto end_branch_0
		} else {

		}
	}
	{
		if v2_1 != nil {
			__t0 = func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 gopurs_runtime.Value
					V2 bool
				}{gopurs_runtime.Value{}, (v2_1).V0, true}
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
	}
end_branch_0:
	return func() struct {
		V0 gopurs_runtime.Value
		V1 gopurs_runtime.Value
		V2 bool
	} { _v := __t0; if _v.Type == 9 && _v.IntVal == 2465973597 && _v.UnsafePtr != nil {
		return struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
			V2 bool
		}{V0: gopurs_runtime.Value{}, V1: (*(*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V0, V2: true}
	}; return struct {
		V0 gopurs_runtime.Value
		V1 gopurs_runtime.Value
		V2 bool
	}{V0: (*(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: gopurs_runtime.Value{}, V2: false} }()
}

func Call_Data_Either_note__4111599802(a_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
note__4111599802:
	for {
		if false {
			continue note__4111599802
		}
		var a_0 gopurs_runtime.Value = a_0_loop
		_ = a_0
		return gopurs_runtime.Apply2(Get_Data_Maybe_maybe__3011962741(), func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
				V2 bool
			}{a_0, gopurs_runtime.Value{}, false}
			if _v.V2 {
				return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
		}(), Get_Data_Either_Right())
	}
}

func Call_Data_Either_note__3211399674(a_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
note__3211399674:
	for {
		if false {
			continue note__3211399674
		}
		var a_0 gopurs_runtime.Value = a_0_loop
		_ = a_0
		return gopurs_runtime.Apply2(Get_Data_Maybe_maybe__1843026042(), func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
				V2 bool
			}{a_0, gopurs_runtime.Value{}, false}
			if _v.V2 {
				return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
		}(), Get_Data_Either_Right())
	}
}

func Call_Data_Either_note__2066121466(a_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
note__2066121466:
	for {
		if false {
			continue note__2066121466
		}
		var a_0 gopurs_runtime.Value = a_0_loop
		_ = a_0
		return gopurs_runtime.Apply2(Get_Data_Maybe_maybe__2496492979(), func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
				V2 bool
			}{a_0, gopurs_runtime.Value{}, false}
			if _v.V2 {
				return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
		}(), Get_Data_Either_Right())
	}
}

func Call_Data_Either_note__271183866(a_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
note__271183866:
	for {
		if false {
			continue note__271183866
		}
		var a_0 gopurs_runtime.Value = a_0_loop
		_ = a_0
		return gopurs_runtime.Apply2(Get_Data_Maybe_maybe__3475452337(), func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
				V2 bool
			}{a_0, gopurs_runtime.Value{}, false}
			if _v.V2 {
				return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
		}(), Get_Data_Either_Right())
	}
}

func Call_Data_Either_note__2268946106(a_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
note__2268946106:
	for {
		if false {
			continue note__2268946106
		}
		var a_0 gopurs_runtime.Value = a_0_loop
		_ = a_0
		return gopurs_runtime.Apply2(Get_Data_Maybe_maybe__3900273117(), func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 struct {
					head gopurs_runtime.Value
					tail *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
					/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
				}
				V2 bool
			}{a_0, struct {
				head gopurs_runtime.Value
				tail *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
				/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
			}{}, false}
			if _v.V2 {
				return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: func() gopurs_runtime.Value {
					orig := _v.V1
					_ = orig
					return gopurs_runtime.RecordDict2("head", "tail", orig.head, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(orig.tail)})
				}()})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
		}(), Get_Data_Either_Right())
	}
}

func Call_Data_Either_note__67430010(a_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
note__67430010:
	for {
		if false {
			continue note__67430010
		}
		var a_0 gopurs_runtime.Value = a_0_loop
		_ = a_0
		return gopurs_runtime.Apply2(Get_Data_Maybe_maybe__3508964590(), func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 struct {
					head gopurs_runtime.Value
					tail []gopurs_runtime.Value
					/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
				}
				V2 bool
			}{a_0, struct {
				head gopurs_runtime.Value
				tail []gopurs_runtime.Value
				/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
			}{}, false}
			if _v.V2 {
				return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: func() gopurs_runtime.Value {
					orig := _v.V1
					_ = orig
					return gopurs_runtime.RecordDict2("head", "tail", orig.head, gopurs_runtime.Array(orig.tail))
				}()})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
		}(), Get_Data_Either_Right())
	}
}

func Call_Data_Either_note__357786010(a_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
note__357786010:
	for {
		if false {
			continue note__357786010
		}
		var a_0 gopurs_runtime.Value = a_0_loop
		_ = a_0
		return gopurs_runtime.Apply2(Get_Data_Maybe_maybe__3750041495(), func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
				V2 bool
			}{a_0, gopurs_runtime.Value{}, false}
			if _v.V2 {
				return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
		}(), Get_Data_Either_Right())
	}
}

func Call_Data_Either_fromRight_prime_(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	var v1_1 gopurs_runtime.Value = v1_1_loop
	_ = v1_1
	var __t0 gopurs_runtime.Value
	{
		if v1_1.Type == 9 && v1_1.IntVal == 2465973597 {
			__t0 = (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V0
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = gopurs_runtime.Apply(v_0, Get_Data_Unit_unit())
	}
end_branch_0:
	return __t0
}

func Call_Data_Either_fromRight(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	var v1_1 gopurs_runtime.Value = v1_1_loop
	_ = v1_1
	var __t0 gopurs_runtime.Value
	{
		if v1_1.Type == 9 && v1_1.IntVal == 2465973597 {
			__t0 = (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V0
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = v_0
	}
end_branch_0:
	return __t0
}

func Call_Data_Either_fromLeft_prime_(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	var v1_1 gopurs_runtime.Value = v1_1_loop
	_ = v1_1
	var __t0 gopurs_runtime.Value
	{
		if v1_1.Type == 9 && v1_1.IntVal == 3711209382 {
			__t0 = (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V0
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = gopurs_runtime.Apply(v_0, Get_Data_Unit_unit())
	}
end_branch_0:
	return __t0
}

func Call_Data_Either_fromLeft(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	var v1_1 gopurs_runtime.Value = v1_1_loop
	_ = v1_1
	var __t0 gopurs_runtime.Value
	{
		if v1_1.Type == 9 && v1_1.IntVal == 3711209382 {
			__t0 = (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V0
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = v_0
	}
end_branch_0:
	return __t0
}

func Call_Data_Either_eqEither(dictEq_0_loop gopurs_runtime.Value, dictEq1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
	_ = dictEq_0
	var dictEq1_1 gopurs_runtime.Value = dictEq1_1_loop
	_ = dictEq1_1
	return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(x_2 gopurs_runtime.Value, y_3 gopurs_runtime.Value) gopurs_runtime.Value {
		var __t0 bool
		{
			if x_2.Type == 9 && x_2.IntVal == 3711209382 {
				__t0 = (y_3.Type == 9 && y_3.IntVal == 3711209382) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(x_2.UnsafePtr).V0, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(y_3.UnsafePtr).V0).IntVal) != (0))
				goto end_branch_0
			} else {

			}
		}
		{
			__t0 = (x_2.Type == 9 && x_2.IntVal == 2465973597) && ((y_3.Type == 9 && y_3.IntVal == 2465973597) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq1_1, "eq"), (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(x_2.UnsafePtr).V0, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_3.UnsafePtr).V0).IntVal) != (0)))
		}
	end_branch_0:
		return gopurs_runtime.Bool(__t0)
	})}))}
}

func Call_Data_Either_ordEither(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
	_ = dictOrd_0
	// TAST (Let): eqEither1_1_0 shape=App(Var) bindingType=Any
	eqEither1_1_0 := gopurs_runtime.Apply(Get_Data_Either_eqEither(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{}))
	_ = eqEither1_1_0
	return gopurs_runtime.Func(func(dictOrd1_2 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): eqEither2_3_1 shape=App(Other) bindingType=(ADT ["Data","Eq","Eq"] [(ADT ["Data","Either","Either"] [(TypeVar a$scope12), (TypeVar b$scope13)])])
		eqEither2_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](gopurs_runtime.Apply(eqEither1_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_2, "Eq0"), gopurs_runtime.Value{})))
		_ = eqEither2_3_1
		return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ord_Ord[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(eqEither2_3_1)}
		}), gopurs_runtime.Func2(func(x_4 gopurs_runtime.Value, y_5 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t3 gopurs_runtime.Value
			{
				if x_4.Type == 9 && x_4.IntVal == 3711209382 {
					var __t2 gopurs_runtime.Value
					{
						if y_5.Type == 9 && y_5.IntVal == 3711209382 {
							__t2 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V0, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V0)
							goto end_branch_2
						} else {

						}
					}
					{
						__t2 = gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}
					}
				end_branch_2:
					__t3 = __t2
					goto end_branch_3
				} else {

				}
			}
			{
				if y_5.Type == 9 && y_5.IntVal == 3711209382 {
					__t3 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
					goto end_branch_3
				} else {

				}
			}
			{
				if (x_4.Type == 9 && x_4.IntVal == 2465973597) && (y_5.Type == 9 && y_5.IntVal == 2465973597) {
					__t3 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd1_2, "compare"), (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V0, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V0)
					goto end_branch_3
				} else {

				}
			}
			{
				__t3 = gopurs_runtime.Value{Type: 9, IntVal: int64(func() uint32 { panic("Failed pattern match") }()), UnsafePtr: nil}
			}
		end_branch_3:
			return __t3
		})}))}
	})
}

func Call_Data_Either_eq1Either(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
	_ = dictEq_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq1[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(dictEq1_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return Call_Data_Eq_eq(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_Either_eqEither(dictEq_0, dictEq1_1)))
	})}))}
}

func Call_Data_Either_ord1Either(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
	_ = dictOrd_0
	// TAST (Let): ordEither1_1_0 shape=App(Var) bindingType=Any
	ordEither1_1_0 := Call_Data_Either_ordEither(dictOrd_0)
	_ = ordEither1_1_0
	// TAST (Let): eq1Either1_2_1 shape=App(Var) bindingType=(ADT ["Data","Eq","Eq1"] [(ADT ["Data","Either","Either"] [(TypeVar a$scope16)])])
	eq1Either1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](Call_Data_Either_eq1Either(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})))
	_ = eq1Either1_2_1
	return gopurs_runtime.Value{Type: 9, IntVal: 1632188299, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ord_Ord1[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(eq1Either1_2_1)}
	}), gopurs_runtime.Func(func(dictOrd1_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return Call_Data_Ord_compare(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](gopurs_runtime.Apply(ordEither1_1_0, dictOrd1_3)))
	})}))}
}

func Call_Data_Either_either(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	var v1_1 gopurs_runtime.Value = v1_1_loop
	_ = v1_1
	var v2_2 gopurs_runtime.Value = v2_2_loop
	_ = v2_2
	var __t0 gopurs_runtime.Value
	{
		if v2_2.Type == 9 && v2_2.IntVal == 3711209382 {
			__t0 = gopurs_runtime.Apply(v_0, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)
			goto end_branch_0
		} else {

		}
	}
	{
		if v2_2.Type == 9 && v2_2.IntVal == 2465973597 {
			__t0 = gopurs_runtime.Apply(v1_1, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
	}
end_branch_0:
	return __t0
}

func Call_Data_Either_either__931018797(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 bool
} {
either__931018797:
	for {
		if false {
			continue either__931018797
		}
		var v_0 gopurs_runtime.Value = v_0_loop
		_ = v_0
		var v1_1 gopurs_runtime.Value = v1_1_loop
		_ = v1_1
		var v2_2 gopurs_runtime.Value = v2_2_loop
		_ = v2_2
		var __t0 gopurs_runtime.Value
		{
			if v2_2.Type == 9 && v2_2.IntVal == 3711209382 {
				__t0 = gopurs_runtime.Apply(v_0, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Rebox_Data_Either_3123684004_1293498952(Rebox_Data_Either_1293498952_3123684004(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]((*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0))))})
				goto end_branch_0
			} else {

			}
		}
		{
			if v2_2.Type == 9 && v2_2.IntVal == 2465973597 {
				__t0 = gopurs_runtime.Apply(v1_1, gopurs_runtime.NormalizeStringValue((*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0))
				goto end_branch_0
			} else {

			}
		}
		{
			__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
		}
	end_branch_0:
		return func() struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
			V2 bool
		} { _v := __t0; if _v.Type == 9 && _v.IntVal == 2465973597 && _v.UnsafePtr != nil {
			return struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
				V2 bool
			}{V0: gopurs_runtime.Value{}, V1: (*(*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V0, V2: true}
		}; return struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
			V2 bool
		}{V0: (*(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: gopurs_runtime.Value{}, V2: false} }()
	}
}

func Call_Data_Either_either__878214054(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 bool
} {
either__878214054:
	for {
		if false {
			continue either__878214054
		}
		var v_0 gopurs_runtime.Value = v_0_loop
		_ = v_0
		var v1_1 gopurs_runtime.Value = v1_1_loop
		_ = v1_1
		var v2_2 gopurs_runtime.Value = v2_2_loop
		_ = v2_2
		var __t0 gopurs_runtime.Value
		{
			if v2_2.Type == 9 && v2_2.IntVal == 3711209382 {
				__t0 = gopurs_runtime.Apply(v_0, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Rebox_Data_Either_3123684004_1293498952(Rebox_Data_Either_1293498952_3123684004(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]((*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0))))})
				goto end_branch_0
			} else {

			}
		}
		{
			if v2_2.Type == 9 && v2_2.IntVal == 2465973597 {
				__t0 = gopurs_runtime.Apply(v1_1, gopurs_runtime.Float((*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0.FloatVal()))
				goto end_branch_0
			} else {

			}
		}
		{
			__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
		}
	end_branch_0:
		return func() struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
			V2 bool
		} { _v := __t0; if _v.Type == 9 && _v.IntVal == 2465973597 && _v.UnsafePtr != nil {
			return struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
				V2 bool
			}{V0: gopurs_runtime.Value{}, V1: (*(*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V0, V2: true}
		}; return struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
			V2 bool
		}{V0: (*(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: gopurs_runtime.Value{}, V2: false} }()
	}
}

func Call_Data_Either_either__3820936092(v_unused_0_loop gopurs_runtime.Value, v1_unused_1_loop gopurs_runtime.Value, v2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
either__3820936092:
	for {
		if false {
			continue either__3820936092
		}
		var v_unused_0 gopurs_runtime.Value = v_unused_0_loop
		_ = v_unused_0
		var v1_unused_1 gopurs_runtime.Value = v1_unused_1_loop
		_ = v1_unused_1
		var v2_2 gopurs_runtime.Value = v2_2_loop
		_ = v2_2
		var __t0 gopurs_runtime.Value
		{
			if v2_2.Type == 9 && v2_2.IntVal == 3711209382 {
				__t0 = gopurs_runtime.Apply(Get_Partial_Unsafe_unsafeCrashWith(), gopurs_runtime.NormalizeStringValue((*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0))
				goto end_branch_0
			} else {

			}
		}
		{
			if v2_2.Type == 9 && v2_2.IntVal == 2465973597 {
				__t0 = gopurs_runtime.Apply(Get_Data_String_Regex_Unsafe_identity(), (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)
				goto end_branch_0
			} else {

			}
		}
		{
			__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
		}
	end_branch_0:
		return __t0
	}
}

func Call_Data_Either_hush(v2_0_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
	var v2_0 gopurs_runtime.Value = v2_0_loop
	_ = v2_0
	var __t0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
	{
		if v2_0.Type == 9 && v2_0.IntVal == 3711209382 {
			__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))})
			goto end_branch_0
		} else {

		}
	}
	{
		if v2_0.Type == 9 && v2_0.IntVal == 2465973597 {
			__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{(*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_0.UnsafePtr).V0, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value { panic("Failed pattern match") }())
	}
end_branch_0:
	return func() struct {
		V0 gopurs_runtime.Value
		V1 bool
	} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t0)}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
		return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
	}; return struct {
		V0 gopurs_runtime.Value
		V1 bool
	}{V0: gopurs_runtime.Value{}, V1: false} }()
}

func Call_Data_Either_isLeft(v2_0_loop gopurs_runtime.Value) bool {
	var v2_0 gopurs_runtime.Value = v2_0_loop
	_ = v2_0
	var __t0 gopurs_runtime.Value
	{
		if v2_0.Type == 9 && v2_0.IntVal == 3711209382 {
			__t0 = gopurs_runtime.Bool(true)
			goto end_branch_0
		} else {

		}
	}
	{
		if v2_0.Type == 9 && v2_0.IntVal == 2465973597 {
			__t0 = gopurs_runtime.Bool(false)
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
	}
end_branch_0:
	return (__t0.IntVal) != (0)
}

func Call_Data_Either_isRight(v2_0_loop gopurs_runtime.Value) bool {
	var v2_0 gopurs_runtime.Value = v2_0_loop
	_ = v2_0
	var __t0 gopurs_runtime.Value
	{
		if v2_0.Type == 9 && v2_0.IntVal == 3711209382 {
			__t0 = gopurs_runtime.Bool(false)
			goto end_branch_0
		} else {

		}
	}
	{
		if v2_0.Type == 9 && v2_0.IntVal == 2465973597 {
			__t0 = gopurs_runtime.Bool(true)
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
	}
end_branch_0:
	return (__t0.IntVal) != (0)
}

func Call_Data_Either_choose(dictAlt_0_loop *Constructor_Control_Alt_Alt[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictAlt_0 *Constructor_Control_Alt_Alt[gopurs_runtime.Value] = dictAlt_0_loop
	_ = dictAlt_0
	// TAST (Let): Functor0_1_0 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m$scope138)])
	Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(dictAlt_0.V0, gopurs_runtime.Value{}))
	_ = Functor0_1_0
	return gopurs_runtime.Func2(func(a_2 gopurs_runtime.Value, b_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(dictAlt_0.V1, gopurs_runtime.Apply2(Functor0_1_0.V0, Get_Data_Either_Left(), a_2), gopurs_runtime.Apply2(Functor0_1_0.V0, Get_Data_Either_Right(), b_3))
	})
}

func Call_Data_Either_boundedEither(dictBounded_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictBounded_0 gopurs_runtime.Value = dictBounded_0_loop
	_ = dictBounded_0
	// TAST (Let): bottom_1_0 shape=Other bindingType=(TypeVar a$scope147)
	bottom_1_0 := gopurs_runtime.RecordGet(dictBounded_0, "bottom")
	_ = bottom_1_0
	// TAST (Let): ordEither1_2_1 shape=App(Var) bindingType=Any
	ordEither1_2_1 := Call_Data_Either_ordEither(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBounded_0, "Ord0"), gopurs_runtime.Value{}))
	_ = ordEither1_2_1
	return gopurs_runtime.Func(func(dictBounded1_3 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): ordEither2_4_2 shape=App(Other) bindingType=(ADT ["Data","Ord","Ord"] [(ADT ["Data","Either","Either"] [(TypeVar a$scope147), (TypeVar b$scope148)])])
		ordEither2_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](gopurs_runtime.Apply(ordEither1_2_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBounded1_3, "Ord0"), gopurs_runtime.Value{})))
		_ = ordEither2_4_2
		return gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer((&Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(ordEither2_4_2)}
		}), func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
				V2 bool
			}{bottom_1_0, gopurs_runtime.Value{}, false}
			if _v.V2 {
				return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
		}(), func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
				V2 bool
			}{gopurs_runtime.Value{}, gopurs_runtime.RecordGet(dictBounded1_3, "top"), true}
			if _v.V2 {
				return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
		}()}))}
	})
}

func Call_Data_Either_blush(v2_0_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
	var v2_0 gopurs_runtime.Value = v2_0_loop
	_ = v2_0
	var __t0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
	{
		if v2_0.Type == 9 && v2_0.IntVal == 3711209382 {
			__t0 = (&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_0.UnsafePtr).V0})
			goto end_branch_0
		} else {

		}
	}
	{
		if v2_0.Type == 9 && v2_0.IntVal == 2465973597 {
			__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))})
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value { panic("Failed pattern match") }())
	}
end_branch_0:
	return func() struct {
		V0 gopurs_runtime.Value
		V1 bool
	} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t0)}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
		return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
	}; return struct {
		V0 gopurs_runtime.Value
		V1 bool
	}{V0: gopurs_runtime.Value{}, V1: false} }()
}

func Call_Data_Either_semigroupEither(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
	_ = dictSemigroup_0
	// TAST (Let): append1_1_0 shape=App(Var) bindingType=(Func [(TypeVar b$scope7), (TypeVar b$scope7)] (TypeVar b$scope7))
	append1_1_0 := Call_Data_Semigroup_go__append(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](dictSemigroup_0))
	_ = append1_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer((&Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(x_2 gopurs_runtime.Value, y_3 gopurs_runtime.Value) gopurs_runtime.Value {
		var __t2 struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
			V2 bool
		}
		{
			if x_2.Type == 9 && x_2.IntVal == 3711209382 {
				__t2 = struct {
					V0 gopurs_runtime.Value
					V1 gopurs_runtime.Value
					V2 bool
				}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(x_2.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				goto end_branch_2
			} else {

			}
		}
		{
			if x_2.Type == 9 && x_2.IntVal == 2465973597 {
				var __t1 struct {
					V0 gopurs_runtime.Value
					V1 gopurs_runtime.Value
					V2 bool
				}
				{
					if y_3.Type == 9 && y_3.IntVal == 3711209382 {
						__t1 = struct {
							V0 gopurs_runtime.Value
							V1 gopurs_runtime.Value
							V2 bool
						}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(y_3.UnsafePtr).V0, gopurs_runtime.Value{}, false}
						goto end_branch_1
					} else {

					}
				}
				{
					if y_3.Type == 9 && y_3.IntVal == 2465973597 {
						__t1 = struct {
							V0 gopurs_runtime.Value
							V1 gopurs_runtime.Value
							V2 bool
						}{gopurs_runtime.Value{}, gopurs_runtime.Apply2(append1_1_0, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(x_2.UnsafePtr).V0, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_3.UnsafePtr).V0), true}
						goto end_branch_1
					} else {

					}
				}
				{
					__t1 = func() struct {
						V0 gopurs_runtime.Value
						V1 gopurs_runtime.Value
						V2 bool
					} { _v := func() gopurs_runtime.Value { panic("Failed pattern match") }(); if _v.Type == 9 && _v.IntVal == 2465973597 && _v.UnsafePtr != nil {
						return struct {
							V0 gopurs_runtime.Value
							V1 gopurs_runtime.Value
							V2 bool
						}{V0: gopurs_runtime.Value{}, V1: (*(*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V0, V2: true}
					}; return struct {
						V0 gopurs_runtime.Value
						V1 gopurs_runtime.Value
						V2 bool
					}{V0: (*(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: gopurs_runtime.Value{}, V2: false} }()
				}
			end_branch_1:
				__t2 = __t1
				goto end_branch_2
			} else {

			}
		}
		{
			__t2 = func() struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
				V2 bool
			} { _v := func() gopurs_runtime.Value { panic("Failed pattern match") }(); if _v.Type == 9 && _v.IntVal == 2465973597 && _v.UnsafePtr != nil {
				return struct {
					V0 gopurs_runtime.Value
					V1 gopurs_runtime.Value
					V2 bool
				}{V0: gopurs_runtime.Value{}, V1: (*(*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V0, V2: true}
			}; return struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
				V2 bool
			}{V0: (*(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: gopurs_runtime.Value{}, V2: false} }()
		}
	end_branch_2:
		return func() gopurs_runtime.Value {
			_v := __t2
			if _v.V2 {
				return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
		}()
	})}))}
}

func Rebox_Data_Either_1293498952_3123684004(in *Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Either_3123684004_1293498952(in *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) *Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(unsafe.Pointer(in))
}
