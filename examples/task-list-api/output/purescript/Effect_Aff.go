package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Effect_Aff_bind gopurs_runtime.Value
var once_Effect_Aff_bind sync.Once

func Get_Effect_Aff_bind() gopurs_runtime.Value {
	once_Effect_Aff_bind.Do(func() {
		cache_Effect_Aff_bind = Get_Effect_bindE()
	})
	return cache_Effect_Aff_bind
}

var cache_Effect_Aff_bind__2699449543 gopurs_runtime.Value
var once_Effect_Aff_bind__2699449543 sync.Once

func Get_Effect_Aff_bind__2699449543() gopurs_runtime.Value {
	once_Effect_Aff_bind__2699449543.Do(func() {
		cache_Effect_Aff_bind__2699449543 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Effect_Aff_bind__2699449543(__eta_norm_1_0_box, __eta_norm_0_1_box)
		})
	})
	return cache_Effect_Aff_bind__2699449543
}

var cache_Effect_Aff_pure gopurs_runtime.Value
var once_Effect_Aff_pure sync.Once

func Get_Effect_Aff_pure() gopurs_runtime.Value {
	once_Effect_Aff_pure.Do(func() {
		cache_Effect_Aff_pure = Get_Effect_pureE()
	})
	return cache_Effect_Aff_pure
}

var cache_Effect_Aff_pure__621669172 gopurs_runtime.Value
var once_Effect_Aff_pure__621669172 sync.Once

func Get_Effect_Aff_pure__621669172() gopurs_runtime.Value {
	once_Effect_Aff_pure__621669172.Do(func() {
		cache_Effect_Aff_pure__621669172 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Effect_Aff_pure__621669172(func() struct {
				isSuspended gopurs_runtime.Value
				join        gopurs_runtime.Value
				kill        gopurs_runtime.Value
				onComplete  gopurs_runtime.Value
				run         gopurs_runtime.Value
				/*gopurs:record-v1:5:0069007300530075007300700065006e006400650064,006a006f0069006e,006b0069006c006c,006f006e0043006f006d0070006c006500740065,00720075006e*/
			} {
				orig := __eta_norm_0_0_box
				_ = orig
				clone := struct {
					isSuspended gopurs_runtime.Value
					join        gopurs_runtime.Value
					kill        gopurs_runtime.Value
					onComplete  gopurs_runtime.Value
					run         gopurs_runtime.Value
					/*gopurs:record-v1:5:0069007300530075007300700065006e006400650064,006a006f0069006e,006b0069006c006c,006f006e0043006f006d0070006c006500740065,00720075006e*/
				}{}
				clone.isSuspended = gopurs_runtime.RecordGet(orig, "isSuspended")
				clone.join = gopurs_runtime.RecordGet(orig, "join")
				clone.kill = gopurs_runtime.RecordGet(orig, "kill")
				clone.onComplete = gopurs_runtime.RecordGet(orig, "onComplete")
				clone.run = gopurs_runtime.RecordGet(orig, "run")
				return clone
			}())
		})
	})
	return cache_Effect_Aff_pure__621669172
}

var cache_Effect_Aff_bind1 gopurs_runtime.Value
var once_Effect_Aff_bind1 sync.Once

func Get_Effect_Aff_bind1() gopurs_runtime.Value {
	once_Effect_Aff_bind1.Do(func() {
		cache_Effect_Aff_bind1 = Get_Effect_bindE()
	})
	return cache_Effect_Aff_bind1
}

var cache_Effect_Aff_bind1__3918760327 gopurs_runtime.Value
var once_Effect_Aff_bind1__3918760327 sync.Once

func Get_Effect_Aff_bind1__3918760327() gopurs_runtime.Value {
	once_Effect_Aff_bind1__3918760327.Do(func() {
		cache_Effect_Aff_bind1__3918760327 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Effect_Aff_bind1__3918760327(__eta_norm_1_0_box, __eta_norm_0_1_box)
		})
	})
	return cache_Effect_Aff_bind1__3918760327
}

var cache_Effect_Aff_discard gopurs_runtime.Value
var once_Effect_Aff_discard sync.Once

func Get_Effect_Aff_discard() gopurs_runtime.Value {
	once_Effect_Aff_discard.Do(func() {
		cache_Effect_Aff_discard = Get_Effect_bindE()
	})
	return cache_Effect_Aff_discard
}

var cache_Effect_Aff_discard__1683979399 gopurs_runtime.Value
var once_Effect_Aff_discard__1683979399 sync.Once

func Get_Effect_Aff_discard__1683979399() gopurs_runtime.Value {
	once_Effect_Aff_discard__1683979399.Do(func() {
		cache_Effect_Aff_discard__1683979399 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Effect_Aff_discard__1683979399(__eta_norm_1_0_box, __eta_norm_0_1_box)
		})
	})
	return cache_Effect_Aff_discard__1683979399
}

var cache_Effect_Aff_pure1 gopurs_runtime.Value
var once_Effect_Aff_pure1 sync.Once

func Get_Effect_Aff_pure1() gopurs_runtime.Value {
	once_Effect_Aff_pure1.Do(func() {
		cache_Effect_Aff_pure1 = Get_Effect_pureE()
	})
	return cache_Effect_Aff_pure1
}

var cache_Effect_Aff_pure1__621669172 gopurs_runtime.Value
var once_Effect_Aff_pure1__621669172 sync.Once

func Get_Effect_Aff_pure1__621669172() gopurs_runtime.Value {
	once_Effect_Aff_pure1__621669172.Do(func() {
		cache_Effect_Aff_pure1__621669172 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Effect_Aff_pure1__621669172(func() struct {
				isSuspended gopurs_runtime.Value
				join        gopurs_runtime.Value
				kill        gopurs_runtime.Value
				onComplete  gopurs_runtime.Value
				run         gopurs_runtime.Value
				/*gopurs:record-v1:5:0069007300530075007300700065006e006400650064,006a006f0069006e,006b0069006c006c,006f006e0043006f006d0070006c006500740065,00720075006e*/
			} {
				orig := __eta_norm_0_0_box
				_ = orig
				clone := struct {
					isSuspended gopurs_runtime.Value
					join        gopurs_runtime.Value
					kill        gopurs_runtime.Value
					onComplete  gopurs_runtime.Value
					run         gopurs_runtime.Value
					/*gopurs:record-v1:5:0069007300530075007300700065006e006400650064,006a006f0069006e,006b0069006c006c,006f006e0043006f006d0070006c006500740065,00720075006e*/
				}{}
				clone.isSuspended = gopurs_runtime.RecordGet(orig, "isSuspended")
				clone.join = gopurs_runtime.RecordGet(orig, "join")
				clone.kill = gopurs_runtime.RecordGet(orig, "kill")
				clone.onComplete = gopurs_runtime.RecordGet(orig, "onComplete")
				clone.run = gopurs_runtime.RecordGet(orig, "run")
				return clone
			}())
		})
	})
	return cache_Effect_Aff_pure1__621669172
}

var cache_Effect_Aff_void gopurs_runtime.Value
var once_Effect_Aff_void sync.Once

func Get_Effect_Aff_void() gopurs_runtime.Value {
	once_Effect_Aff_void.Do(func() {
		cache_Effect_Aff_void = Call_Data_Functor_void(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()))
	})
	return cache_Effect_Aff_void
}

var cache_Effect_Aff_go__map gopurs_runtime.Value
var once_Effect_Aff_go__map sync.Once

func Get_Effect_Aff_go__map() gopurs_runtime.Value {
	once_Effect_Aff_go__map.Do(func() {
		cache_Effect_Aff_go__map = Get_Data_Functor_map__2042768475()
	})
	return cache_Effect_Aff_go__map
}

var cache_Effect_Aff_void1 gopurs_runtime.Value
var once_Effect_Aff_void1 sync.Once

func Get_Effect_Aff_void1() gopurs_runtime.Value {
	once_Effect_Aff_void1.Do(func() {
		cache_Effect_Aff_void1 = Call_Data_Functor_void(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()))
	})
	return cache_Effect_Aff_void1
}

var cache_Effect_Aff_pure2 gopurs_runtime.Value
var once_Effect_Aff_pure2 sync.Once

func Get_Effect_Aff_pure2() gopurs_runtime.Value {
	once_Effect_Aff_pure2.Do(func() {
		cache_Effect_Aff_pure2 = Get_Control_Applicative_pure__2088919292()
	})
	return cache_Effect_Aff_pure2
}

var cache_Effect_Aff_pure3 gopurs_runtime.Value
var once_Effect_Aff_pure3 sync.Once

func Get_Effect_Aff_pure3() gopurs_runtime.Value {
	once_Effect_Aff_pure3.Do(func() {
		cache_Effect_Aff_pure3 = Get_Control_Applicative_pure__2387439490()
	})
	return cache_Effect_Aff_pure3
}

var cache_Effect_Aff_bind2 gopurs_runtime.Value
var once_Effect_Aff_bind2 sync.Once

func Get_Effect_Aff_bind2() gopurs_runtime.Value {
	once_Effect_Aff_bind2.Do(func() {
		cache_Effect_Aff_bind2 = Get_Effect_bindE()
	})
	return cache_Effect_Aff_bind2
}

var cache_Effect_Aff_discard1 gopurs_runtime.Value
var once_Effect_Aff_discard1 sync.Once

func Get_Effect_Aff_discard1() gopurs_runtime.Value {
	once_Effect_Aff_discard1.Do(func() {
		cache_Effect_Aff_discard1 = Get_Effect_bindE()
	})
	return cache_Effect_Aff_discard1
}

var cache_Effect_Aff_pure4 gopurs_runtime.Value
var once_Effect_Aff_pure4 sync.Once

func Get_Effect_Aff_pure4() gopurs_runtime.Value {
	once_Effect_Aff_pure4.Do(func() {
		cache_Effect_Aff_pure4 = Get_Effect_pureE()
	})
	return cache_Effect_Aff_pure4
}

var cache_Effect_Aff_compose gopurs_runtime.Value
var once_Effect_Aff_compose sync.Once

func Get_Effect_Aff_compose() gopurs_runtime.Value {
	once_Effect_Aff_compose.Do(func() {
		cache_Effect_Aff_compose = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Effect_Aff_compose
}

var cache_Effect_Aff_compose1 gopurs_runtime.Value
var once_Effect_Aff_compose1 sync.Once

func Get_Effect_Aff_compose1() gopurs_runtime.Value {
	once_Effect_Aff_compose1.Do(func() {
		cache_Effect_Aff_compose1 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Effect_Aff_compose1
}

var cache_Effect_Aff_functorAff gopurs_runtime.Value
var once_Effect_Aff_functorAff sync.Once

func Get_Effect_Aff_functorAff() gopurs_runtime.Value {
	once_Effect_Aff_functorAff.Do(func() {
		cache_Effect_Aff_functorAff = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, Get_Effect_Aff__map()}))}
	})
	return cache_Effect_Aff_functorAff
}

var cache_Effect_Aff_monadAff gopurs_runtime.Value
var once_Effect_Aff_monadAff sync.Once

func Get_Effect_Aff_monadAff() gopurs_runtime.Value {
	once_Effect_Aff_monadAff.Do(func() {
		cache_Effect_Aff_monadAff = gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Effect_Aff_applicativeAff()))}
		}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_Aff_bindAff()))}
		})}))}
	})
	return cache_Effect_Aff_monadAff
}

var cache_Effect_Aff_bindAff gopurs_runtime.Value
var once_Effect_Aff_bindAff sync.Once

func Get_Effect_Aff_bindAff() gopurs_runtime.Value {
	once_Effect_Aff_bindAff.Do(func() {
		cache_Effect_Aff_bindAff = gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Effect_Aff_applyAff()))}
		}), Get_Effect_Aff__bind()}))}
	})
	return cache_Effect_Aff_bindAff
}

var cache_Effect_Aff_applyAff gopurs_runtime.Value
var once_Effect_Aff_applyAff sync.Once

func Get_Effect_Aff_applyAff() gopurs_runtime.Value {
	once_Effect_Aff_applyAff.Do(func() {
		cache_Effect_Aff_applyAff = gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_Aff_functorAff()))}
		}), Call_Control_Monad_ap(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](Get_Effect_Aff_monadAff()))}))}
	})
	return cache_Effect_Aff_applyAff
}

var cache_Effect_Aff_applicativeAff gopurs_runtime.Value
var once_Effect_Aff_applicativeAff sync.Once

func Get_Effect_Aff_applicativeAff() gopurs_runtime.Value {
	once_Effect_Aff_applicativeAff.Do(func() {
		cache_Effect_Aff_applicativeAff = gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Effect_Aff_applyAff()))}
		}), Get_Effect_Aff__pure()}))}
	})
	return cache_Effect_Aff_applicativeAff
}

var cache_Effect_Aff_monadEffectAff gopurs_runtime.Value
var once_Effect_Aff_monadEffectAff sync.Once

func Get_Effect_Aff_monadEffectAff() gopurs_runtime.Value {
	once_Effect_Aff_monadEffectAff.Do(func() {
		cache_Effect_Aff_monadEffectAff = gopurs_runtime.Value{Type: 9, IntVal: 2217729261, UnsafePtr: unsafe.Pointer((&Constructor_Effect_Class_MonadEffect[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](Get_Effect_Aff_monadAff()))}
		}), Get_Effect_Aff__liftEffect()}))}
	})
	return cache_Effect_Aff_monadEffectAff
}

var cache_Effect_Aff_liftEffect gopurs_runtime.Value
var once_Effect_Aff_liftEffect sync.Once

func Get_Effect_Aff_liftEffect() gopurs_runtime.Value {
	once_Effect_Aff_liftEffect.Do(func() {
		cache_Effect_Aff_liftEffect = Get_Effect_Aff__liftEffect()
	})
	return cache_Effect_Aff_liftEffect
}

var cache_Effect_Aff_compose1__4105490545 gopurs_runtime.Value
var once_Effect_Aff_compose1__4105490545 sync.Once

func Get_Effect_Aff_compose1__4105490545() gopurs_runtime.Value {
	once_Effect_Aff_compose1__4105490545.Do(func() {
		cache_Effect_Aff_compose1__4105490545 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Effect_Aff_compose1__4105490545(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
		})
	})
	return cache_Effect_Aff_compose1__4105490545
}

var cache_Effect_Aff_compose1__1269036657 gopurs_runtime.Value
var once_Effect_Aff_compose1__1269036657 sync.Once

func Get_Effect_Aff_compose1__1269036657() gopurs_runtime.Value {
	once_Effect_Aff_compose1__1269036657.Do(func() {
		cache_Effect_Aff_compose1__1269036657 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Effect_Aff_compose1__1269036657(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
		})
	})
	return cache_Effect_Aff_compose1__1269036657
}

var cache_Effect_Aff_compose2 gopurs_runtime.Value
var once_Effect_Aff_compose2 sync.Once

func Get_Effect_Aff_compose2() gopurs_runtime.Value {
	once_Effect_Aff_compose2.Do(func() {
		cache_Effect_Aff_compose2 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Effect_Aff_compose2
}

var cache_Effect_Aff_pure5 gopurs_runtime.Value
var once_Effect_Aff_pure5 sync.Once

func Get_Effect_Aff_pure5() gopurs_runtime.Value {
	once_Effect_Aff_pure5.Do(func() {
		cache_Effect_Aff_pure5 = Get_Control_Applicative_pure__3156000668()
	})
	return cache_Effect_Aff_pure5
}

var cache_Effect_Aff_Fiber gopurs_runtime.Value
var once_Effect_Aff_Fiber sync.Once

func Get_Effect_Aff_Fiber() gopurs_runtime.Value {
	once_Effect_Aff_Fiber.Do(func() {
		cache_Effect_Aff_Fiber = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				orig := Call_Effect_Aff_Fiber(func() struct {
					isSuspended gopurs_runtime.Value
					join        gopurs_runtime.Value
					kill        gopurs_runtime.Value
					onComplete  gopurs_runtime.Value
					run         gopurs_runtime.Value
					/*gopurs:record-v1:5:0069007300530075007300700065006e006400650064,006a006f0069006e,006b0069006c006c,006f006e0043006f006d0070006c006500740065,00720075006e*/
				} {
					orig := x_0_box
					_ = orig
					clone := struct {
						isSuspended gopurs_runtime.Value
						join        gopurs_runtime.Value
						kill        gopurs_runtime.Value
						onComplete  gopurs_runtime.Value
						run         gopurs_runtime.Value
						/*gopurs:record-v1:5:0069007300530075007300700065006e006400650064,006a006f0069006e,006b0069006c006c,006f006e0043006f006d0070006c006500740065,00720075006e*/
					}{}
					clone.isSuspended = gopurs_runtime.RecordGet(orig, "isSuspended")
					clone.join = gopurs_runtime.RecordGet(orig, "join")
					clone.kill = gopurs_runtime.RecordGet(orig, "kill")
					clone.onComplete = gopurs_runtime.RecordGet(orig, "onComplete")
					clone.run = gopurs_runtime.RecordGet(orig, "run")
					return clone
				}())
				_ = orig
				return gopurs_runtime.RecordDict5("isSuspended", "join", "kill", "onComplete", "run", orig.isSuspended, orig.join, orig.kill, orig.onComplete, orig.run)
			}()
		})
	})
	return cache_Effect_Aff_Fiber
}

var cache_Effect_Aff_Fiber__288829658 gopurs_runtime.Value
var once_Effect_Aff_Fiber__288829658 sync.Once

func Get_Effect_Aff_Fiber__288829658() gopurs_runtime.Value {
	once_Effect_Aff_Fiber__288829658.Do(func() {
		cache_Effect_Aff_Fiber__288829658 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				orig := Call_Effect_Aff_Fiber__288829658(func() struct {
					isSuspended gopurs_runtime.Value
					join        gopurs_runtime.Value
					kill        gopurs_runtime.Value
					onComplete  gopurs_runtime.Value
					run         gopurs_runtime.Value
					/*gopurs:record-v1:5:0069007300530075007300700065006e006400650064,006a006f0069006e,006b0069006c006c,006f006e0043006f006d0070006c006500740065,00720075006e*/
				} {
					orig := x_0_box
					_ = orig
					clone := struct {
						isSuspended gopurs_runtime.Value
						join        gopurs_runtime.Value
						kill        gopurs_runtime.Value
						onComplete  gopurs_runtime.Value
						run         gopurs_runtime.Value
						/*gopurs:record-v1:5:0069007300530075007300700065006e006400650064,006a006f0069006e,006b0069006c006c,006f006e0043006f006d0070006c006500740065,00720075006e*/
					}{}
					clone.isSuspended = gopurs_runtime.RecordGet(orig, "isSuspended")
					clone.join = gopurs_runtime.RecordGet(orig, "join")
					clone.kill = gopurs_runtime.RecordGet(orig, "kill")
					clone.onComplete = gopurs_runtime.RecordGet(orig, "onComplete")
					clone.run = gopurs_runtime.RecordGet(orig, "run")
					return clone
				}())
				_ = orig
				return gopurs_runtime.RecordDict5("isSuspended", "join", "kill", "onComplete", "run", orig.isSuspended, orig.join, orig.kill, orig.onComplete, orig.run)
			}()
		})
	})
	return cache_Effect_Aff_Fiber__288829658
}

var cache_Effect_Aff_Canceler gopurs_runtime.Value
var once_Effect_Aff_Canceler sync.Once

func Get_Effect_Aff_Canceler() gopurs_runtime.Value {
	once_Effect_Aff_Canceler.Do(func() {
		cache_Effect_Aff_Canceler = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Effect_Aff_Canceler(x_0_box)
		})
	})
	return cache_Effect_Aff_Canceler
}

var cache_Effect_Aff_newtypeCanceler gopurs_runtime.Value
var once_Effect_Aff_newtypeCanceler sync.Once

func Get_Effect_Aff_newtypeCanceler() gopurs_runtime.Value {
	once_Effect_Aff_newtypeCanceler.Do(func() {
		cache_Effect_Aff_newtypeCanceler = gopurs_runtime.Value{Type: 9, IntVal: 3322196858, UnsafePtr: unsafe.Pointer((&Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{}
		})}))}
	})
	return cache_Effect_Aff_newtypeCanceler
}

var cache_Effect_Aff_makeFiber gopurs_runtime.Value
var once_Effect_Aff_makeFiber sync.Once

func Get_Effect_Aff_makeFiber() gopurs_runtime.Value {
	once_Effect_Aff_makeFiber.Do(func() {
		cache_Effect_Aff_makeFiber = gopurs_runtime.Func(func(aff_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Effect_Aff_makeFiber(aff_0_box)
		})
	})
	return cache_Effect_Aff_makeFiber
}

var cache_Effect_Aff_makeFiber__4085038308 gopurs_runtime.Value
var once_Effect_Aff_makeFiber__4085038308 sync.Once

func Get_Effect_Aff_makeFiber__4085038308() gopurs_runtime.Value {
	once_Effect_Aff_makeFiber__4085038308.Do(func() {
		cache_Effect_Aff_makeFiber__4085038308 = gopurs_runtime.Func(func(aff_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Effect_Aff_makeFiber__4085038308(aff_0_box)
		})
	})
	return cache_Effect_Aff_makeFiber__4085038308
}

var cache_Effect_Aff_makeAff gopurs_runtime.Value
var once_Effect_Aff_makeAff sync.Once

func Get_Effect_Aff_makeAff() gopurs_runtime.Value {
	once_Effect_Aff_makeAff.Do(func() {
		cache_Effect_Aff_makeAff = gopurs_runtime.Func(func(build_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Effect_Aff_makeAff(build_0_box)
		})
	})
	return cache_Effect_Aff_makeAff
}

var cache_Effect_Aff_makeAff__1714332378 gopurs_runtime.Value
var once_Effect_Aff_makeAff__1714332378 sync.Once

func Get_Effect_Aff_makeAff__1714332378() gopurs_runtime.Value {
	once_Effect_Aff_makeAff__1714332378.Do(func() {
		cache_Effect_Aff_makeAff__1714332378 = gopurs_runtime.Func(func(build_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Effect_Aff_makeAff__1714332378(build_0_box)
		})
	})
	return cache_Effect_Aff_makeAff__1714332378
}

var cache_Effect_Aff_makeAff__1501555834 gopurs_runtime.Value
var once_Effect_Aff_makeAff__1501555834 sync.Once

func Get_Effect_Aff_makeAff__1501555834() gopurs_runtime.Value {
	once_Effect_Aff_makeAff__1501555834.Do(func() {
		cache_Effect_Aff_makeAff__1501555834 = gopurs_runtime.Func(func(build_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Effect_Aff_makeAff__1501555834(build_0_box)
		})
	})
	return cache_Effect_Aff_makeAff__1501555834
}

var cache_Effect_Aff_makeAff__2246750170 gopurs_runtime.Value
var once_Effect_Aff_makeAff__2246750170 sync.Once

func Get_Effect_Aff_makeAff__2246750170() gopurs_runtime.Value {
	once_Effect_Aff_makeAff__2246750170.Do(func() {
		cache_Effect_Aff_makeAff__2246750170 = gopurs_runtime.Func(func(build_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Effect_Aff_makeAff__2246750170(build_0_box)
		})
	})
	return cache_Effect_Aff_makeAff__2246750170
}

var cache_Effect_Aff_makeAff__158558778 gopurs_runtime.Value
var once_Effect_Aff_makeAff__158558778 sync.Once

func Get_Effect_Aff_makeAff__158558778() gopurs_runtime.Value {
	once_Effect_Aff_makeAff__158558778.Do(func() {
		cache_Effect_Aff_makeAff__158558778 = gopurs_runtime.Func(func(build_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Effect_Aff_makeAff__158558778(build_0_box)
		})
	})
	return cache_Effect_Aff_makeAff__158558778
}

var cache_Effect_Aff_makeAff__925755130 gopurs_runtime.Value
var once_Effect_Aff_makeAff__925755130 sync.Once

func Get_Effect_Aff_makeAff__925755130() gopurs_runtime.Value {
	once_Effect_Aff_makeAff__925755130.Do(func() {
		cache_Effect_Aff_makeAff__925755130 = gopurs_runtime.Func(func(build_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Effect_Aff_makeAff__925755130(build_0_box)
		})
	})
	return cache_Effect_Aff_makeAff__925755130
}

var cache_Effect_Aff_launchSuspendedAff gopurs_runtime.Value
var once_Effect_Aff_launchSuspendedAff sync.Once

func Get_Effect_Aff_launchSuspendedAff() gopurs_runtime.Value {
	once_Effect_Aff_launchSuspendedAff.Do(func() {
		cache_Effect_Aff_launchSuspendedAff = gopurs_runtime.Func(func(aff_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Effect_Aff_launchSuspendedAff(aff_0_box)
		})
	})
	return cache_Effect_Aff_launchSuspendedAff
}

var cache_Effect_Aff_launchSuspendedAff__4085038308 gopurs_runtime.Value
var once_Effect_Aff_launchSuspendedAff__4085038308 sync.Once

func Get_Effect_Aff_launchSuspendedAff__4085038308() gopurs_runtime.Value {
	once_Effect_Aff_launchSuspendedAff__4085038308.Do(func() {
		cache_Effect_Aff_launchSuspendedAff__4085038308 = gopurs_runtime.Func(func(aff_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Effect_Aff_launchSuspendedAff__4085038308(aff_0_box)
		})
	})
	return cache_Effect_Aff_launchSuspendedAff__4085038308
}

var cache_Effect_Aff_launchAff gopurs_runtime.Value
var once_Effect_Aff_launchAff sync.Once

func Get_Effect_Aff_launchAff() gopurs_runtime.Value {
	once_Effect_Aff_launchAff.Do(func() {
		cache_Effect_Aff_launchAff = gopurs_runtime.Func(func(aff_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Effect_Aff_launchAff(aff_0_box)
		})
	})
	return cache_Effect_Aff_launchAff
}

var cache_Effect_Aff_launchAff__4085038308 gopurs_runtime.Value
var once_Effect_Aff_launchAff__4085038308 sync.Once

func Get_Effect_Aff_launchAff__4085038308() gopurs_runtime.Value {
	once_Effect_Aff_launchAff__4085038308.Do(func() {
		cache_Effect_Aff_launchAff__4085038308 = gopurs_runtime.Func(func(aff_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Effect_Aff_launchAff__4085038308(aff_0_box)
		})
	})
	return cache_Effect_Aff_launchAff__4085038308
}

var cache_Effect_Aff_launchAff_ gopurs_runtime.Value
var once_Effect_Aff_launchAff_ sync.Once

func Get_Effect_Aff_launchAff_() gopurs_runtime.Value {
	once_Effect_Aff_launchAff_.Do(func() {
		cache_Effect_Aff_launchAff_ = gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Call_Data_Functor_void(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect())), Get_Effect_Aff_launchAff())
	})
	return cache_Effect_Aff_launchAff_
}

var cache_Effect_Aff_functorParAff gopurs_runtime.Value
var once_Effect_Aff_functorParAff sync.Once

func Get_Effect_Aff_functorParAff() gopurs_runtime.Value {
	once_Effect_Aff_functorParAff.Do(func() {
		cache_Effect_Aff_functorParAff = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, Get_Effect_Aff__parAffMap()}))}
	})
	return cache_Effect_Aff_functorParAff
}

var cache_Effect_Aff_map1 gopurs_runtime.Value
var once_Effect_Aff_map1 sync.Once

func Get_Effect_Aff_map1() gopurs_runtime.Value {
	once_Effect_Aff_map1.Do(func() {
		cache_Effect_Aff_map1 = Get_Effect_Aff__map()
	})
	return cache_Effect_Aff_map1
}

var cache_Effect_Aff_delay gopurs_runtime.Value
var once_Effect_Aff_delay sync.Once

func Get_Effect_Aff_delay() gopurs_runtime.Value {
	once_Effect_Aff_delay.Do(func() {
		cache_Effect_Aff_delay = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Effect_Aff_delay(v_0_box.FloatVal())
		})
	})
	return cache_Effect_Aff_delay
}

var cache_Effect_Aff_bracket gopurs_runtime.Value
var once_Effect_Aff_bracket sync.Once

func Get_Effect_Aff_bracket() gopurs_runtime.Value {
	once_Effect_Aff_bracket.Do(func() {
		cache_Effect_Aff_bracket = gopurs_runtime.Func2(func(acquire_0_box gopurs_runtime.Value, completed_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Effect_Aff_bracket(acquire_0_box, completed_1_box)
		})
	})
	return cache_Effect_Aff_bracket
}

var cache_Effect_Aff_bracket__939099628 gopurs_runtime.Value
var once_Effect_Aff_bracket__939099628 sync.Once

func Get_Effect_Aff_bracket__939099628() gopurs_runtime.Value {
	once_Effect_Aff_bracket__939099628.Do(func() {
		cache_Effect_Aff_bracket__939099628 = gopurs_runtime.Func3(func(acquire_0_box gopurs_runtime.Value, completed_unused_1_box gopurs_runtime.Value, __eta_norm_0_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Effect_Aff_bracket__939099628(acquire_0_box, completed_unused_1_box, __eta_norm_0_2_box)
		})
	})
	return cache_Effect_Aff_bracket__939099628
}

var cache_Effect_Aff_bracket__3139827911 gopurs_runtime.Value
var once_Effect_Aff_bracket__3139827911 sync.Once

func Get_Effect_Aff_bracket__3139827911() gopurs_runtime.Value {
	once_Effect_Aff_bracket__3139827911.Do(func() {
		cache_Effect_Aff_bracket__3139827911 = gopurs_runtime.Func3(func(acquire_0_box gopurs_runtime.Value, completed_unused_1_box gopurs_runtime.Value, __eta_norm_0_unused_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Effect_Aff_bracket__3139827911(acquire_0_box, completed_unused_1_box, __eta_norm_0_unused_2_box)
		})
	})
	return cache_Effect_Aff_bracket__3139827911
}

var cache_Effect_Aff_applyParAff gopurs_runtime.Value
var once_Effect_Aff_applyParAff sync.Once

func Get_Effect_Aff_applyParAff() gopurs_runtime.Value {
	once_Effect_Aff_applyParAff.Do(func() {
		cache_Effect_Aff_applyParAff = gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_Aff_functorParAff()))}
		}), Get_Effect_Aff__parAffApply()}))}
	})
	return cache_Effect_Aff_applyParAff
}

var cache_Effect_Aff_lift2 gopurs_runtime.Value
var once_Effect_Aff_lift2 sync.Once

func Get_Effect_Aff_lift2() gopurs_runtime.Value {
	once_Effect_Aff_lift2.Do(func() {
		cache_Effect_Aff_lift2 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value, b_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Effect_Aff_lift2(f_0_box, a_1_box, b_2_box)
		})
	})
	return cache_Effect_Aff_lift2
}

var cache_Effect_Aff_semigroupParAff gopurs_runtime.Value
var once_Effect_Aff_semigroupParAff sync.Once

func Get_Effect_Aff_semigroupParAff() gopurs_runtime.Value {
	once_Effect_Aff_semigroupParAff.Do(func() {
		cache_Effect_Aff_semigroupParAff = gopurs_runtime.Func(func(dictSemigroup_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Effect_Aff_semigroupParAff(dictSemigroup_0_box)
		})
	})
	return cache_Effect_Aff_semigroupParAff
}

var cache_Effect_Aff_pure6 gopurs_runtime.Value
var once_Effect_Aff_pure6 sync.Once

func Get_Effect_Aff_pure6() gopurs_runtime.Value {
	once_Effect_Aff_pure6.Do(func() {
		cache_Effect_Aff_pure6 = Get_Control_Applicative_pure__2964015333()
	})
	return cache_Effect_Aff_pure6
}

var cache_Effect_Aff_pure7 gopurs_runtime.Value
var once_Effect_Aff_pure7 sync.Once

func Get_Effect_Aff_pure7() gopurs_runtime.Value {
	once_Effect_Aff_pure7.Do(func() {
		cache_Effect_Aff_pure7 = Get_Effect_Aff__pure()
	})
	return cache_Effect_Aff_pure7
}

var cache_Effect_Aff_bind3 gopurs_runtime.Value
var once_Effect_Aff_bind3 sync.Once

func Get_Effect_Aff_bind3() gopurs_runtime.Value {
	once_Effect_Aff_bind3.Do(func() {
		cache_Effect_Aff_bind3 = Get_Effect_Aff__bind()
	})
	return cache_Effect_Aff_bind3
}

var cache_Effect_Aff_lift21 gopurs_runtime.Value
var once_Effect_Aff_lift21 sync.Once

func Get_Effect_Aff_lift21() gopurs_runtime.Value {
	once_Effect_Aff_lift21.Do(func() {
		cache_Effect_Aff_lift21 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value, b_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Effect_Aff_lift21(f_0_box, a_1_box, b_2_box)
		})
	})
	return cache_Effect_Aff_lift21
}

var cache_Effect_Aff_apply gopurs_runtime.Value
var once_Effect_Aff_apply sync.Once

func Get_Effect_Aff_apply() gopurs_runtime.Value {
	once_Effect_Aff_apply.Do(func() {
		cache_Effect_Aff_apply = gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Effect_Aff_applyAff()).V1
	})
	return cache_Effect_Aff_apply
}

var cache_Effect_Aff_pure8 gopurs_runtime.Value
var once_Effect_Aff_pure8 sync.Once

func Get_Effect_Aff_pure8() gopurs_runtime.Value {
	once_Effect_Aff_pure8.Do(func() {
		cache_Effect_Aff_pure8 = Get_Effect_Aff__pure()
	})
	return cache_Effect_Aff_pure8
}

var cache_Effect_Aff_bind4 gopurs_runtime.Value
var once_Effect_Aff_bind4 sync.Once

func Get_Effect_Aff_bind4() gopurs_runtime.Value {
	once_Effect_Aff_bind4.Do(func() {
		cache_Effect_Aff_bind4 = Get_Effect_Aff__bind()
	})
	return cache_Effect_Aff_bind4
}

var cache_Effect_Aff_discard2 gopurs_runtime.Value
var once_Effect_Aff_discard2 sync.Once

func Get_Effect_Aff_discard2() gopurs_runtime.Value {
	once_Effect_Aff_discard2.Do(func() {
		cache_Effect_Aff_discard2 = Get_Effect_Aff__bind()
	})
	return cache_Effect_Aff_discard2
}

var cache_Effect_Aff_pure9 gopurs_runtime.Value
var once_Effect_Aff_pure9 sync.Once

func Get_Effect_Aff_pure9() gopurs_runtime.Value {
	once_Effect_Aff_pure9.Do(func() {
		cache_Effect_Aff_pure9 = Get_Effect_Aff__pure()
	})
	return cache_Effect_Aff_pure9
}

var cache_Effect_Aff_bind5 gopurs_runtime.Value
var once_Effect_Aff_bind5 sync.Once

func Get_Effect_Aff_bind5() gopurs_runtime.Value {
	once_Effect_Aff_bind5.Do(func() {
		cache_Effect_Aff_bind5 = Get_Control_Bind_bind__4176309933()
	})
	return cache_Effect_Aff_bind5
}

var cache_Effect_Aff_bindFlipped gopurs_runtime.Value
var once_Effect_Aff_bindFlipped sync.Once

func Get_Effect_Aff_bindFlipped() gopurs_runtime.Value {
	once_Effect_Aff_bindFlipped.Do(func() {
		cache_Effect_Aff_bindFlipped = gopurs_runtime.Func2(func(b_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Effect_Aff_bindFlipped(b_0_box, a_1_box)
		})
	})
	return cache_Effect_Aff_bindFlipped
}

var cache_Effect_Aff_bindFlipped__879378959 gopurs_runtime.Value
var once_Effect_Aff_bindFlipped__879378959 sync.Once

func Get_Effect_Aff_bindFlipped__879378959() gopurs_runtime.Value {
	once_Effect_Aff_bindFlipped__879378959.Do(func() {
		cache_Effect_Aff_bindFlipped__879378959 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Effect_Aff_bindFlipped__879378959(__eta_norm_1_0_box, __eta_norm_0_1_box)
		})
	})
	return cache_Effect_Aff_bindFlipped__879378959
}

var cache_Effect_Aff_bindFlipped__52396111 gopurs_runtime.Value
var once_Effect_Aff_bindFlipped__52396111 sync.Once

func Get_Effect_Aff_bindFlipped__52396111() gopurs_runtime.Value {
	once_Effect_Aff_bindFlipped__52396111.Do(func() {
		cache_Effect_Aff_bindFlipped__52396111 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Effect_Aff_bindFlipped__52396111(__eta_norm_1_0_box, __eta_norm_0_1_box)
		})
	})
	return cache_Effect_Aff_bindFlipped__52396111
}

var cache_Effect_Aff_bindFlipped1 gopurs_runtime.Value
var once_Effect_Aff_bindFlipped1 sync.Once

func Get_Effect_Aff_bindFlipped1() gopurs_runtime.Value {
	once_Effect_Aff_bindFlipped1.Do(func() {
		cache_Effect_Aff_bindFlipped1 = gopurs_runtime.Func2(func(b_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Effect_Aff_bindFlipped1(b_0_box, a_1_box)
		})
	})
	return cache_Effect_Aff_bindFlipped1
}

var cache_Effect_Aff_bind6 gopurs_runtime.Value
var once_Effect_Aff_bind6 sync.Once

func Get_Effect_Aff_bind6() gopurs_runtime.Value {
	once_Effect_Aff_bind6.Do(func() {
		cache_Effect_Aff_bind6 = Get_Effect_Aff__bind()
	})
	return cache_Effect_Aff_bind6
}

var cache_Effect_Aff_pure10 gopurs_runtime.Value
var once_Effect_Aff_pure10 sync.Once

func Get_Effect_Aff_pure10() gopurs_runtime.Value {
	once_Effect_Aff_pure10.Do(func() {
		cache_Effect_Aff_pure10 = Get_Effect_Aff__pure()
	})
	return cache_Effect_Aff_pure10
}

var cache_Effect_Aff_pure11 gopurs_runtime.Value
var once_Effect_Aff_pure11 sync.Once

func Get_Effect_Aff_pure11() gopurs_runtime.Value {
	once_Effect_Aff_pure11.Do(func() {
		cache_Effect_Aff_pure11 = Get_Effect_Aff__pure()
	})
	return cache_Effect_Aff_pure11
}

var cache_Effect_Aff_cancelWith gopurs_runtime.Value
var once_Effect_Aff_cancelWith sync.Once

func Get_Effect_Aff_cancelWith() gopurs_runtime.Value {
	once_Effect_Aff_cancelWith.Do(func() {
		cache_Effect_Aff_cancelWith = gopurs_runtime.Func2(func(aff_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Effect_Aff_cancelWith(aff_0_box, v_1_box)
		})
	})
	return cache_Effect_Aff_cancelWith
}

var cache_Effect_Aff_finally gopurs_runtime.Value
var once_Effect_Aff_finally sync.Once

func Get_Effect_Aff_finally() gopurs_runtime.Value {
	once_Effect_Aff_finally.Do(func() {
		cache_Effect_Aff_finally = gopurs_runtime.Func2(func(fin_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Effect_Aff_finally(fin_0_box, a_1_box)
		})
	})
	return cache_Effect_Aff_finally
}

var cache_Effect_Aff_invincible gopurs_runtime.Value
var once_Effect_Aff_invincible sync.Once

func Get_Effect_Aff_invincible() gopurs_runtime.Value {
	once_Effect_Aff_invincible.Do(func() {
		cache_Effect_Aff_invincible = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Effect_Aff_invincible(a_0_box)
		})
	})
	return cache_Effect_Aff_invincible
}

var cache_Effect_Aff_lazyAff gopurs_runtime.Value
var once_Effect_Aff_lazyAff sync.Once

func Get_Effect_Aff_lazyAff() gopurs_runtime.Value {
	once_Effect_Aff_lazyAff.Do(func() {
		cache_Effect_Aff_lazyAff = gopurs_runtime.Value{Type: 9, IntVal: 1860244333, UnsafePtr: unsafe.Pointer((&Constructor_Control_Lazy_Lazy[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(Get_Effect_Aff__bind(), gopurs_runtime.Apply(Get_Effect_Aff__pure(), Get_Data_Unit_unit()), f_0)
		})}))}
	})
	return cache_Effect_Aff_lazyAff
}

var cache_Effect_Aff_parallelAff gopurs_runtime.Value
var once_Effect_Aff_parallelAff sync.Once

func Get_Effect_Aff_parallelAff() gopurs_runtime.Value {
	once_Effect_Aff_parallelAff.Do(func() {
		cache_Effect_Aff_parallelAff = gopurs_runtime.Value{Type: 9, IntVal: 327692956, UnsafePtr: unsafe.Pointer((&Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Effect_Aff_applyAff()))}
		}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Effect_Aff_applyParAff()))}
		}), Get_Unsafe_Coerce_unsafeCoerce(), Get_Effect_Aff__sequential()}))}
	})
	return cache_Effect_Aff_parallelAff
}

var cache_Effect_Aff_applicativeParAff gopurs_runtime.Value
var once_Effect_Aff_applicativeParAff sync.Once

func Get_Effect_Aff_applicativeParAff() gopurs_runtime.Value {
	once_Effect_Aff_applicativeParAff.Do(func() {
		cache_Effect_Aff_applicativeParAff = gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Effect_Aff_applyParAff()))}
		}), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Unsafe_Coerce_unsafeCoerce(), Get_Effect_Aff__pure())}))}
	})
	return cache_Effect_Aff_applicativeParAff
}

var cache_Effect_Aff_pure12 gopurs_runtime.Value
var once_Effect_Aff_pure12 sync.Once

func Get_Effect_Aff_pure12() gopurs_runtime.Value {
	once_Effect_Aff_pure12.Do(func() {
		cache_Effect_Aff_pure12 = gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Effect_Aff_applicativeParAff()).V1
	})
	return cache_Effect_Aff_pure12
}

var cache_Effect_Aff_parSequence_ gopurs_runtime.Value
var once_Effect_Aff_parSequence_ sync.Once

func Get_Effect_Aff_parSequence_() gopurs_runtime.Value {
	once_Effect_Aff_parSequence_.Do(func() {
		cache_Effect_Aff_parSequence_ = gopurs_runtime.Apply2(Call_Control_Parallel_parSequence_(gopurs_runtime.CoerceToStruct[Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Effect_Aff_parallelAff())), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Effect_Aff_applicativeParAff()))}, gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))})
	})
	return cache_Effect_Aff_parSequence_
}

var cache_Effect_Aff_monoidParAff gopurs_runtime.Value
var once_Effect_Aff_monoidParAff sync.Once

func Get_Effect_Aff_monoidParAff() gopurs_runtime.Value {
	once_Effect_Aff_monoidParAff.Do(func() {
		cache_Effect_Aff_monoidParAff = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Effect_Aff_monoidParAff(dictMonoid_0_box)
		})
	})
	return cache_Effect_Aff_monoidParAff
}

var cache_Effect_Aff_semigroupCanceler gopurs_runtime.Value
var once_Effect_Aff_semigroupCanceler sync.Once

func Get_Effect_Aff_semigroupCanceler() gopurs_runtime.Value {
	once_Effect_Aff_semigroupCanceler.Do(func() {
		cache_Effect_Aff_semigroupCanceler = gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer((&Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value, err_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply3(Call_Control_Parallel_parSequence_(gopurs_runtime.CoerceToStruct[Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Effect_Aff_parallelAff())), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Effect_Aff_applicativeParAff()))}, gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))}, gopurs_runtime.Array([]gopurs_runtime.Value{gopurs_runtime.Apply(v_0, err_2), gopurs_runtime.Apply(v1_1, err_2)}))
		})}))}
	})
	return cache_Effect_Aff_semigroupCanceler
}

var cache_Effect_Aff_semigroupAff gopurs_runtime.Value
var once_Effect_Aff_semigroupAff sync.Once

func Get_Effect_Aff_semigroupAff() gopurs_runtime.Value {
	once_Effect_Aff_semigroupAff.Do(func() {
		cache_Effect_Aff_semigroupAff = gopurs_runtime.Func(func(dictSemigroup_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Effect_Aff_semigroupAff(dictSemigroup_0_box)
		})
	})
	return cache_Effect_Aff_semigroupAff
}

var cache_Effect_Aff_liftEffect1 gopurs_runtime.Value
var once_Effect_Aff_liftEffect1 sync.Once

func Get_Effect_Aff_liftEffect1() gopurs_runtime.Value {
	once_Effect_Aff_liftEffect1.Do(func() {
		cache_Effect_Aff_liftEffect1 = Get_Effect_Class_liftEffect__3733620925()
	})
	return cache_Effect_Aff_liftEffect1
}

var cache_Effect_Aff_liftEffect2 gopurs_runtime.Value
var once_Effect_Aff_liftEffect2 sync.Once

func Get_Effect_Aff_liftEffect2() gopurs_runtime.Value {
	once_Effect_Aff_liftEffect2.Do(func() {
		cache_Effect_Aff_liftEffect2 = Get_Effect_Class_liftEffect__3747263901()
	})
	return cache_Effect_Aff_liftEffect2
}

var cache_Effect_Aff_liftEffect3 gopurs_runtime.Value
var once_Effect_Aff_liftEffect3 sync.Once

func Get_Effect_Aff_liftEffect3() gopurs_runtime.Value {
	once_Effect_Aff_liftEffect3.Do(func() {
		cache_Effect_Aff_liftEffect3 = Get_Effect_Aff__liftEffect()
	})
	return cache_Effect_Aff_liftEffect3
}

var cache_Effect_Aff_liftEffect4 gopurs_runtime.Value
var once_Effect_Aff_liftEffect4 sync.Once

func Get_Effect_Aff_liftEffect4() gopurs_runtime.Value {
	once_Effect_Aff_liftEffect4.Do(func() {
		cache_Effect_Aff_liftEffect4 = Get_Effect_Aff__liftEffect()
	})
	return cache_Effect_Aff_liftEffect4
}

var cache_Effect_Aff_effectCanceler gopurs_runtime.Value
var once_Effect_Aff_effectCanceler sync.Once

func Get_Effect_Aff_effectCanceler() gopurs_runtime.Value {
	once_Effect_Aff_effectCanceler.Do(func() {
		cache_Effect_Aff_effectCanceler = gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return x_0
		}), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Function_go__const(), Get_Effect_Aff__liftEffect()))
	})
	return cache_Effect_Aff_effectCanceler
}

var cache_Effect_Aff_joinFiber gopurs_runtime.Value
var once_Effect_Aff_joinFiber sync.Once

func Get_Effect_Aff_joinFiber() gopurs_runtime.Value {
	once_Effect_Aff_joinFiber.Do(func() {
		cache_Effect_Aff_joinFiber = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Effect_Aff_joinFiber(func() struct {
				isSuspended gopurs_runtime.Value
				join        gopurs_runtime.Value
				kill        gopurs_runtime.Value
				onComplete  gopurs_runtime.Value
				run         gopurs_runtime.Value
				/*gopurs:record-v1:5:0069007300530075007300700065006e006400650064,006a006f0069006e,006b0069006c006c,006f006e0043006f006d0070006c006500740065,00720075006e*/
			} {
				orig := v_0_box
				_ = orig
				clone := struct {
					isSuspended gopurs_runtime.Value
					join        gopurs_runtime.Value
					kill        gopurs_runtime.Value
					onComplete  gopurs_runtime.Value
					run         gopurs_runtime.Value
					/*gopurs:record-v1:5:0069007300530075007300700065006e006400650064,006a006f0069006e,006b0069006c006c,006f006e0043006f006d0070006c006500740065,00720075006e*/
				}{}
				clone.isSuspended = gopurs_runtime.RecordGet(orig, "isSuspended")
				clone.join = gopurs_runtime.RecordGet(orig, "join")
				clone.kill = gopurs_runtime.RecordGet(orig, "kill")
				clone.onComplete = gopurs_runtime.RecordGet(orig, "onComplete")
				clone.run = gopurs_runtime.RecordGet(orig, "run")
				return clone
			}())
		})
	})
	return cache_Effect_Aff_joinFiber
}

var cache_Effect_Aff_functorFiber gopurs_runtime.Value
var once_Effect_Aff_functorFiber sync.Once

func Get_Effect_Aff_functorFiber() gopurs_runtime.Value {
	once_Effect_Aff_functorFiber.Do(func() {
		cache_Effect_Aff_functorFiber = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Effect_Aff_179938118_2812149806((&Constructor_Data_Functor_Functor[struct {
			isSuspended gopurs_runtime.Value
			join        gopurs_runtime.Value
			kill        gopurs_runtime.Value
			onComplete  gopurs_runtime.Value
			run         gopurs_runtime.Value
			/*gopurs:record-v1:5:0069007300530075007300700065006e006400650064,006a006f0069006e,006b0069006c006c,006f006e0043006f006d0070006c006500740065,00720075006e*/
		}]{1, gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, t_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(Get_Effect_Unsafe_unsafePerformEffect(), Call_Effect_Aff_makeFiber(gopurs_runtime.Apply2(Get_Effect_Aff__map(), f_0, Call_Effect_Aff_joinFiber(func() struct {
				isSuspended gopurs_runtime.Value
				join        gopurs_runtime.Value
				kill        gopurs_runtime.Value
				onComplete  gopurs_runtime.Value
				run         gopurs_runtime.Value
				/*gopurs:record-v1:5:0069007300530075007300700065006e006400650064,006a006f0069006e,006b0069006c006c,006f006e0043006f006d0070006c006500740065,00720075006e*/
			} {
				orig := t_1
				_ = orig
				clone := struct {
					isSuspended gopurs_runtime.Value
					join        gopurs_runtime.Value
					kill        gopurs_runtime.Value
					onComplete  gopurs_runtime.Value
					run         gopurs_runtime.Value
					/*gopurs:record-v1:5:0069007300530075007300700065006e006400650064,006a006f0069006e,006b0069006c006c,006f006e0043006f006d0070006c006500740065,00720075006e*/
				}{}
				clone.isSuspended = gopurs_runtime.RecordGet(orig, "isSuspended")
				clone.join = gopurs_runtime.RecordGet(orig, "join")
				clone.kill = gopurs_runtime.RecordGet(orig, "kill")
				clone.onComplete = gopurs_runtime.RecordGet(orig, "onComplete")
				clone.run = gopurs_runtime.RecordGet(orig, "run")
				return clone
			}()))))
		})})))}
	})
	return cache_Effect_Aff_functorFiber
}

var cache_Effect_Aff_applyFiber gopurs_runtime.Value
var once_Effect_Aff_applyFiber sync.Once

func Get_Effect_Aff_applyFiber() gopurs_runtime.Value {
	once_Effect_Aff_applyFiber.Do(func() {
		cache_Effect_Aff_applyFiber = gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(Rebox_Effect_Aff_1844969297_3741347833((&Constructor_Control_Apply_Apply[struct {
			isSuspended gopurs_runtime.Value
			join        gopurs_runtime.Value
			kill        gopurs_runtime.Value
			onComplete  gopurs_runtime.Value
			run         gopurs_runtime.Value
			/*gopurs:record-v1:5:0069007300530075007300700065006e006400650064,006a006f0069006e,006b0069006c006c,006f006e0043006f006d0070006c006500740065,00720075006e*/
		}]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Effect_Aff_179938118_2812149806(Rebox_Effect_Aff_2812149806_179938118(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_Aff_functorFiber()))))}
		}), gopurs_runtime.Func2(func(t1_0 gopurs_runtime.Value, t2_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(Get_Effect_Unsafe_unsafePerformEffect(), Call_Effect_Aff_makeFiber(gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Effect_Aff_applyAff()).V1, Call_Effect_Aff_joinFiber(func() struct {
				isSuspended gopurs_runtime.Value
				join        gopurs_runtime.Value
				kill        gopurs_runtime.Value
				onComplete  gopurs_runtime.Value
				run         gopurs_runtime.Value
				/*gopurs:record-v1:5:0069007300530075007300700065006e006400650064,006a006f0069006e,006b0069006c006c,006f006e0043006f006d0070006c006500740065,00720075006e*/
			} {
				orig := t1_0
				_ = orig
				clone := struct {
					isSuspended gopurs_runtime.Value
					join        gopurs_runtime.Value
					kill        gopurs_runtime.Value
					onComplete  gopurs_runtime.Value
					run         gopurs_runtime.Value
					/*gopurs:record-v1:5:0069007300530075007300700065006e006400650064,006a006f0069006e,006b0069006c006c,006f006e0043006f006d0070006c006500740065,00720075006e*/
				}{}
				clone.isSuspended = gopurs_runtime.RecordGet(orig, "isSuspended")
				clone.join = gopurs_runtime.RecordGet(orig, "join")
				clone.kill = gopurs_runtime.RecordGet(orig, "kill")
				clone.onComplete = gopurs_runtime.RecordGet(orig, "onComplete")
				clone.run = gopurs_runtime.RecordGet(orig, "run")
				return clone
			}()), Call_Effect_Aff_joinFiber(func() struct {
				isSuspended gopurs_runtime.Value
				join        gopurs_runtime.Value
				kill        gopurs_runtime.Value
				onComplete  gopurs_runtime.Value
				run         gopurs_runtime.Value
				/*gopurs:record-v1:5:0069007300530075007300700065006e006400650064,006a006f0069006e,006b0069006c006c,006f006e0043006f006d0070006c006500740065,00720075006e*/
			} {
				orig := t2_1
				_ = orig
				clone := struct {
					isSuspended gopurs_runtime.Value
					join        gopurs_runtime.Value
					kill        gopurs_runtime.Value
					onComplete  gopurs_runtime.Value
					run         gopurs_runtime.Value
					/*gopurs:record-v1:5:0069007300530075007300700065006e006400650064,006a006f0069006e,006b0069006c006c,006f006e0043006f006d0070006c006500740065,00720075006e*/
				}{}
				clone.isSuspended = gopurs_runtime.RecordGet(orig, "isSuspended")
				clone.join = gopurs_runtime.RecordGet(orig, "join")
				clone.kill = gopurs_runtime.RecordGet(orig, "kill")
				clone.onComplete = gopurs_runtime.RecordGet(orig, "onComplete")
				clone.run = gopurs_runtime.RecordGet(orig, "run")
				return clone
			}()))))
		})})))}
	})
	return cache_Effect_Aff_applyFiber
}

var cache_Effect_Aff_applicativeFiber gopurs_runtime.Value
var once_Effect_Aff_applicativeFiber sync.Once

func Get_Effect_Aff_applicativeFiber() gopurs_runtime.Value {
	once_Effect_Aff_applicativeFiber.Do(func() {
		cache_Effect_Aff_applicativeFiber = gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(Rebox_Effect_Aff_520806353_1439734649((&Constructor_Control_Applicative_Applicative[struct {
			isSuspended gopurs_runtime.Value
			join        gopurs_runtime.Value
			kill        gopurs_runtime.Value
			onComplete  gopurs_runtime.Value
			run         gopurs_runtime.Value
			/*gopurs:record-v1:5:0069007300530075007300700065006e006400650064,006a006f0069006e,006b0069006c006c,006f006e0043006f006d0070006c006500740065,00720075006e*/
		}]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(Rebox_Effect_Aff_1844969297_3741347833(Rebox_Effect_Aff_3741347833_1844969297(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Effect_Aff_applyFiber()))))}
		}), gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(Get_Effect_Unsafe_unsafePerformEffect(), Call_Effect_Aff_makeFiber(gopurs_runtime.Apply(Get_Effect_Aff__pure(), a_0)))
		})})))}
	})
	return cache_Effect_Aff_applicativeFiber
}

var cache_Effect_Aff_forkAff gopurs_runtime.Value
var once_Effect_Aff_forkAff sync.Once

func Get_Effect_Aff_forkAff() gopurs_runtime.Value {
	once_Effect_Aff_forkAff.Do(func() {
		cache_Effect_Aff_forkAff = gopurs_runtime.Func(func(aff_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Effect_Aff_forkAff(aff_0_box)
		})
	})
	return cache_Effect_Aff_forkAff
}

var cache_Effect_Aff_killFiber gopurs_runtime.Value
var once_Effect_Aff_killFiber sync.Once

func Get_Effect_Aff_killFiber() gopurs_runtime.Value {
	once_Effect_Aff_killFiber.Do(func() {
		cache_Effect_Aff_killFiber = gopurs_runtime.Func2(func(e_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Effect_Aff_killFiber(e_0_box, func() struct {
				isSuspended gopurs_runtime.Value
				join        gopurs_runtime.Value
				kill        gopurs_runtime.Value
				onComplete  gopurs_runtime.Value
				run         gopurs_runtime.Value
				/*gopurs:record-v1:5:0069007300530075007300700065006e006400650064,006a006f0069006e,006b0069006c006c,006f006e0043006f006d0070006c006500740065,00720075006e*/
			} {
				orig := v_1_box
				_ = orig
				clone := struct {
					isSuspended gopurs_runtime.Value
					join        gopurs_runtime.Value
					kill        gopurs_runtime.Value
					onComplete  gopurs_runtime.Value
					run         gopurs_runtime.Value
					/*gopurs:record-v1:5:0069007300530075007300700065006e006400650064,006a006f0069006e,006b0069006c006c,006f006e0043006f006d0070006c006500740065,00720075006e*/
				}{}
				clone.isSuspended = gopurs_runtime.RecordGet(orig, "isSuspended")
				clone.join = gopurs_runtime.RecordGet(orig, "join")
				clone.kill = gopurs_runtime.RecordGet(orig, "kill")
				clone.onComplete = gopurs_runtime.RecordGet(orig, "onComplete")
				clone.run = gopurs_runtime.RecordGet(orig, "run")
				return clone
			}())
		})
	})
	return cache_Effect_Aff_killFiber
}

var cache_Effect_Aff_fiberCanceler gopurs_runtime.Value
var once_Effect_Aff_fiberCanceler sync.Once

func Get_Effect_Aff_fiberCanceler() gopurs_runtime.Value {
	once_Effect_Aff_fiberCanceler.Do(func() {
		cache_Effect_Aff_fiberCanceler = gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return x_0
		}), gopurs_runtime.Func2(func(b_0 gopurs_runtime.Value, a_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Effect_Aff_killFiber(a_1, func() struct {
				isSuspended gopurs_runtime.Value
				join        gopurs_runtime.Value
				kill        gopurs_runtime.Value
				onComplete  gopurs_runtime.Value
				run         gopurs_runtime.Value
				/*gopurs:record-v1:5:0069007300530075007300700065006e006400650064,006a006f0069006e,006b0069006c006c,006f006e0043006f006d0070006c006500740065,00720075006e*/
			} {
				orig := b_0
				_ = orig
				clone := struct {
					isSuspended gopurs_runtime.Value
					join        gopurs_runtime.Value
					kill        gopurs_runtime.Value
					onComplete  gopurs_runtime.Value
					run         gopurs_runtime.Value
					/*gopurs:record-v1:5:0069007300530075007300700065006e006400650064,006a006f0069006e,006b0069006c006c,006f006e0043006f006d0070006c006500740065,00720075006e*/
				}{}
				clone.isSuspended = gopurs_runtime.RecordGet(orig, "isSuspended")
				clone.join = gopurs_runtime.RecordGet(orig, "join")
				clone.kill = gopurs_runtime.RecordGet(orig, "kill")
				clone.onComplete = gopurs_runtime.RecordGet(orig, "onComplete")
				clone.run = gopurs_runtime.RecordGet(orig, "run")
				return clone
			}())
		}))
	})
	return cache_Effect_Aff_fiberCanceler
}

var cache_Effect_Aff_supervise gopurs_runtime.Value
var once_Effect_Aff_supervise sync.Once

func Get_Effect_Aff_supervise() gopurs_runtime.Value {
	once_Effect_Aff_supervise.Do(func() {
		cache_Effect_Aff_supervise = gopurs_runtime.Func(func(aff_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Effect_Aff_supervise(aff_0_box)
		})
	})
	return cache_Effect_Aff_supervise
}

var cache_Effect_Aff_suspendAff gopurs_runtime.Value
var once_Effect_Aff_suspendAff sync.Once

func Get_Effect_Aff_suspendAff() gopurs_runtime.Value {
	once_Effect_Aff_suspendAff.Do(func() {
		cache_Effect_Aff_suspendAff = gopurs_runtime.Func(func(aff_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Effect_Aff_suspendAff(aff_0_box)
		})
	})
	return cache_Effect_Aff_suspendAff
}

var cache_Effect_Aff_monadSTAff gopurs_runtime.Value
var once_Effect_Aff_monadSTAff sync.Once

func Get_Effect_Aff_monadSTAff() gopurs_runtime.Value {
	once_Effect_Aff_monadSTAff.Do(func() {
		cache_Effect_Aff_monadSTAff = gopurs_runtime.Value{Type: 9, IntVal: 2155655715, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_ST_Class_MonadST[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](Get_Effect_Aff_monadAff()))}
		}), gopurs_runtime.Apply2(Call_Control_Semigroupoid_composeFlipped(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Control_Monad_ST_Global_toEffect(), Get_Effect_Aff__liftEffect())}))}
	})
	return cache_Effect_Aff_monadSTAff
}

var cache_Effect_Aff_monadThrowAff gopurs_runtime.Value
var once_Effect_Aff_monadThrowAff sync.Once

func Get_Effect_Aff_monadThrowAff() gopurs_runtime.Value {
	once_Effect_Aff_monadThrowAff.Do(func() {
		cache_Effect_Aff_monadThrowAff = gopurs_runtime.Value{Type: 9, IntVal: 23967309, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Error_Class_MonadThrow[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](Get_Effect_Aff_monadAff()))}
		}), Get_Effect_Aff__throwError()}))}
	})
	return cache_Effect_Aff_monadThrowAff
}

var cache_Effect_Aff_monadErrorAff gopurs_runtime.Value
var once_Effect_Aff_monadErrorAff sync.Once

func Get_Effect_Aff_monadErrorAff() gopurs_runtime.Value {
	once_Effect_Aff_monadErrorAff.Do(func() {
		cache_Effect_Aff_monadErrorAff = gopurs_runtime.Value{Type: 9, IntVal: 1402181699, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Error_Class_MonadError[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 23967309, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Error_Class_MonadThrow[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Effect_Aff_monadThrowAff()))}
		}), Get_Effect_Aff__catchError()}))}
	})
	return cache_Effect_Aff_monadErrorAff
}

var cache_Effect_Aff_try gopurs_runtime.Value
var once_Effect_Aff_try sync.Once

func Get_Effect_Aff_try() gopurs_runtime.Value {
	once_Effect_Aff_try.Do(func() {
		cache_Effect_Aff_try = Call_Control_Monad_Error_Class_try(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Error_Class_MonadError[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Effect_Aff_monadErrorAff()))
	})
	return cache_Effect_Aff_try
}

var cache_Effect_Aff_try__4239810482 gopurs_runtime.Value
var once_Effect_Aff_try__4239810482 sync.Once

func Get_Effect_Aff_try__4239810482() gopurs_runtime.Value {
	once_Effect_Aff_try__4239810482.Do(func() {
		cache_Effect_Aff_try__4239810482 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Effect_Aff_try__4239810482(__eta_norm_0_0_box)
		})
	})
	return cache_Effect_Aff_try__4239810482
}

var cache_Effect_Aff_try__1646464146 gopurs_runtime.Value
var once_Effect_Aff_try__1646464146 sync.Once

func Get_Effect_Aff_try__1646464146() gopurs_runtime.Value {
	once_Effect_Aff_try__1646464146.Do(func() {
		cache_Effect_Aff_try__1646464146 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Effect_Aff_try__1646464146(__eta_norm_0_0_box)
		})
	})
	return cache_Effect_Aff_try__1646464146
}

var cache_Effect_Aff_try1 gopurs_runtime.Value
var once_Effect_Aff_try1 sync.Once

func Get_Effect_Aff_try1() gopurs_runtime.Value {
	once_Effect_Aff_try1.Do(func() {
		cache_Effect_Aff_try1 = Call_Control_Monad_Error_Class_try(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Error_Class_MonadError[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Effect_Aff_monadErrorAff()))
	})
	return cache_Effect_Aff_try1
}

var cache_Effect_Aff_catchError gopurs_runtime.Value
var once_Effect_Aff_catchError sync.Once

func Get_Effect_Aff_catchError() gopurs_runtime.Value {
	once_Effect_Aff_catchError.Do(func() {
		cache_Effect_Aff_catchError = Get_Effect_Aff__catchError()
	})
	return cache_Effect_Aff_catchError
}

var cache_Effect_Aff_attempt gopurs_runtime.Value
var once_Effect_Aff_attempt sync.Once

func Get_Effect_Aff_attempt() gopurs_runtime.Value {
	once_Effect_Aff_attempt.Do(func() {
		cache_Effect_Aff_attempt = Call_Control_Monad_Error_Class_try(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Error_Class_MonadError[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Effect_Aff_monadErrorAff()))
	})
	return cache_Effect_Aff_attempt
}

var cache_Effect_Aff_runAff gopurs_runtime.Value
var once_Effect_Aff_runAff sync.Once

func Get_Effect_Aff_runAff() gopurs_runtime.Value {
	once_Effect_Aff_runAff.Do(func() {
		cache_Effect_Aff_runAff = gopurs_runtime.Func2(func(k_0_box gopurs_runtime.Value, aff_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Effect_Aff_runAff(k_0_box, aff_1_box)
		})
	})
	return cache_Effect_Aff_runAff
}

var cache_Effect_Aff_runAff__2603932529 gopurs_runtime.Value
var once_Effect_Aff_runAff__2603932529 sync.Once

func Get_Effect_Aff_runAff__2603932529() gopurs_runtime.Value {
	once_Effect_Aff_runAff__2603932529.Do(func() {
		cache_Effect_Aff_runAff__2603932529 = gopurs_runtime.Func2(func(k_0_box gopurs_runtime.Value, aff_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Effect_Aff_runAff__2603932529(k_0_box, aff_1_box)
		})
	})
	return cache_Effect_Aff_runAff__2603932529
}

var cache_Effect_Aff_runAff__2939864241 gopurs_runtime.Value
var once_Effect_Aff_runAff__2939864241 sync.Once

func Get_Effect_Aff_runAff__2939864241() gopurs_runtime.Value {
	once_Effect_Aff_runAff__2939864241.Do(func() {
		cache_Effect_Aff_runAff__2939864241 = gopurs_runtime.Func2(func(k_0_box gopurs_runtime.Value, aff_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Effect_Aff_runAff__2939864241(k_0_box, aff_1_box)
		})
	})
	return cache_Effect_Aff_runAff__2939864241
}

var cache_Effect_Aff_runAff_ gopurs_runtime.Value
var once_Effect_Aff_runAff_ sync.Once

func Get_Effect_Aff_runAff_() gopurs_runtime.Value {
	once_Effect_Aff_runAff_.Do(func() {
		cache_Effect_Aff_runAff_ = gopurs_runtime.Func2(func(k_0_box gopurs_runtime.Value, aff_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Effect_Aff_runAff_(k_0_box, aff_1_box)
		})
	})
	return cache_Effect_Aff_runAff_
}

var cache_Effect_Aff_runAff___3349181959 gopurs_runtime.Value
var once_Effect_Aff_runAff___3349181959 sync.Once

func Get_Effect_Aff_runAff___3349181959() gopurs_runtime.Value {
	once_Effect_Aff_runAff___3349181959.Do(func() {
		cache_Effect_Aff_runAff___3349181959 = gopurs_runtime.Func2(func(k_0_box gopurs_runtime.Value, aff_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Effect_Aff_runAff___3349181959(k_0_box, aff_1_box)
		})
	})
	return cache_Effect_Aff_runAff___3349181959
}

var cache_Effect_Aff_runAff___358700103 gopurs_runtime.Value
var once_Effect_Aff_runAff___358700103 sync.Once

func Get_Effect_Aff_runAff___358700103() gopurs_runtime.Value {
	once_Effect_Aff_runAff___358700103.Do(func() {
		cache_Effect_Aff_runAff___358700103 = gopurs_runtime.Func2(func(k_0_box gopurs_runtime.Value, aff_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Effect_Aff_runAff___358700103(k_0_box, aff_1_box)
		})
	})
	return cache_Effect_Aff_runAff___358700103
}

var cache_Effect_Aff_runSuspendedAff gopurs_runtime.Value
var once_Effect_Aff_runSuspendedAff sync.Once

func Get_Effect_Aff_runSuspendedAff() gopurs_runtime.Value {
	once_Effect_Aff_runSuspendedAff.Do(func() {
		cache_Effect_Aff_runSuspendedAff = gopurs_runtime.Func2(func(k_0_box gopurs_runtime.Value, aff_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Effect_Aff_runSuspendedAff(k_0_box, aff_1_box)
		})
	})
	return cache_Effect_Aff_runSuspendedAff
}

var cache_Effect_Aff_monadRecAff gopurs_runtime.Value
var once_Effect_Aff_monadRecAff sync.Once

func Get_Effect_Aff_monadRecAff() gopurs_runtime.Value {
	once_Effect_Aff_monadRecAff.Do(func() {
		cache_Effect_Aff_monadRecAff = gopurs_runtime.Value{Type: 9, IntVal: 3709389635, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](Get_Effect_Aff_monadAff()))}
		}), gopurs_runtime.Func(func(k_0 gopurs_runtime.Value) gopurs_runtime.Value {
			var go__go_1_0_0 gopurs_runtime.Value
			_ = go__go_1_0_0
			var go__go_1_0_0_cell *gopurs_runtime.Value
			_ = go__go_1_0_0_cell
			// FALLBACK TCO: isLoop=false len=1
			go__go_1_0_0 = gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(Get_Effect_Aff__bind(), gopurs_runtime.Apply(k_0, a_2), gopurs_runtime.Func(func(res_3 gopurs_runtime.Value) gopurs_runtime.Value {
					var __t1 gopurs_runtime.Value
					{
						if res_3.Type == 9 && res_3.IntVal == 60402430 {
							__t1 = gopurs_runtime.Apply(Get_Effect_Aff__pure(), (*Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value])(res_3.UnsafePtr).V0)
							goto end_branch_1
						} else {

						}
					}
					{
						if res_3.Type == 9 && res_3.IntVal == 525585346 {
							__t1 = gopurs_runtime.Apply((*go__go_1_0_0_cell), (*Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value])(res_3.UnsafePtr).V0)
							goto end_branch_1
						} else {

						}
					}
					{
						__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
					}
				end_branch_1:
					return __t1
				}))
			})
			go__go_1_0_0_cell = &go__go_1_0_0
			return go__go_1_0_0
		})}))}
	})
	return cache_Effect_Aff_monadRecAff
}

var cache_Effect_Aff_monoidAff gopurs_runtime.Value
var once_Effect_Aff_monoidAff sync.Once

func Get_Effect_Aff_monoidAff() gopurs_runtime.Value {
	once_Effect_Aff_monoidAff.Do(func() {
		cache_Effect_Aff_monoidAff = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Effect_Aff_monoidAff(dictMonoid_0_box)
		})
	})
	return cache_Effect_Aff_monoidAff
}

var cache_Effect_Aff_nonCanceler gopurs_runtime.Value
var once_Effect_Aff_nonCanceler sync.Once

func Get_Effect_Aff_nonCanceler() gopurs_runtime.Value {
	once_Effect_Aff_nonCanceler.Do(func() {
		cache_Effect_Aff_nonCanceler = gopurs_runtime.Apply(Get_Data_Function_const__4207166068(), gopurs_runtime.Apply(Get_Effect_Aff__pure(), Get_Data_Unit_unit()))
	})
	return cache_Effect_Aff_nonCanceler
}

var cache_Effect_Aff_monoidCanceler gopurs_runtime.Value
var once_Effect_Aff_monoidCanceler sync.Once

func Get_Effect_Aff_monoidCanceler() gopurs_runtime.Value {
	once_Effect_Aff_monoidCanceler.Do(func() {
		cache_Effect_Aff_monoidCanceler = gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer((&Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](Get_Effect_Aff_semigroupCanceler()))}
		}), gopurs_runtime.Apply(Get_Data_Function_const__4207166068(), gopurs_runtime.Apply(Get_Effect_Aff__pure(), Get_Data_Unit_unit()))}))}
	})
	return cache_Effect_Aff_monoidCanceler
}

var cache_Effect_Aff_mempty gopurs_runtime.Value
var once_Effect_Aff_mempty sync.Once

func Get_Effect_Aff_mempty() gopurs_runtime.Value {
	once_Effect_Aff_mempty.Do(func() {
		cache_Effect_Aff_mempty = Get_Data_Monoid_mempty__1856050949()
	})
	return cache_Effect_Aff_mempty
}

var cache_Effect_Aff_never gopurs_runtime.Value
var once_Effect_Aff_never sync.Once

func Get_Effect_Aff_never() gopurs_runtime.Value {
	once_Effect_Aff_never.Do(func() {
		cache_Effect_Aff_never = Call_Effect_Aff_makeAff(gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
				return Get_Data_Monoid_mempty__1856050949()
			})
		}))
	})
	return cache_Effect_Aff_never
}

var cache_Effect_Aff_apathize gopurs_runtime.Value
var once_Effect_Aff_apathize sync.Once

func Get_Effect_Aff_apathize() gopurs_runtime.Value {
	once_Effect_Aff_apathize.Do(func() {
		cache_Effect_Aff_apathize = gopurs_runtime.Apply2(Call_Control_Semigroupoid_composeFlipped(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Effect_Aff_attempt(), gopurs_runtime.Apply(Get_Effect_Aff__map(), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Get_Data_Unit_unit()
		})))
	})
	return cache_Effect_Aff_apathize
}

var cache_Effect_Aff_altParAff gopurs_runtime.Value
var once_Effect_Aff_altParAff sync.Once

func Get_Effect_Aff_altParAff() gopurs_runtime.Value {
	once_Effect_Aff_altParAff.Do(func() {
		cache_Effect_Aff_altParAff = gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer((&Constructor_Control_Alt_Alt[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_Aff_functorParAff()))}
		}), Get_Effect_Aff__parAffAlt()}))}
	})
	return cache_Effect_Aff_altParAff
}

var cache_Effect_Aff_altAff gopurs_runtime.Value
var once_Effect_Aff_altAff sync.Once

func Get_Effect_Aff_altAff() gopurs_runtime.Value {
	once_Effect_Aff_altAff.Do(func() {
		cache_Effect_Aff_altAff = gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer((&Constructor_Control_Alt_Alt[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_Aff_functorAff()))}
		}), gopurs_runtime.Func2(func(a1_0 gopurs_runtime.Value, a2_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(Get_Effect_Aff__catchError(), a1_0, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
				return a2_1
			}))
		})}))}
	})
	return cache_Effect_Aff_altAff
}

var cache_Effect_Aff_plusAff gopurs_runtime.Value
var once_Effect_Aff_plusAff sync.Once

func Get_Effect_Aff_plusAff() gopurs_runtime.Value {
	once_Effect_Aff_plusAff.Do(func() {
		cache_Effect_Aff_plusAff = gopurs_runtime.Value{Type: 9, IntVal: 3709470893, UnsafePtr: unsafe.Pointer((&Constructor_Control_Plus_Plus[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Alt_Alt[gopurs_runtime.Value]](Get_Effect_Aff_altAff()))}
		}), gopurs_runtime.Apply(Get_Effect_Aff__throwError(), gopurs_runtime.Apply(Get_Effect_Exception_error(), gopurs_runtime.StrWithCodeUnits("Always fails", 12)))}))}
	})
	return cache_Effect_Aff_plusAff
}

var cache_Effect_Aff_plusParAff gopurs_runtime.Value
var once_Effect_Aff_plusParAff sync.Once

func Get_Effect_Aff_plusParAff() gopurs_runtime.Value {
	once_Effect_Aff_plusParAff.Do(func() {
		cache_Effect_Aff_plusParAff = gopurs_runtime.Value{Type: 9, IntVal: 3709470893, UnsafePtr: unsafe.Pointer((&Constructor_Control_Plus_Plus[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Alt_Alt[gopurs_runtime.Value]](Get_Effect_Aff_altParAff()))}
		}), gopurs_runtime.CoerceToStruct[Constructor_Control_Plus_Plus[gopurs_runtime.Value]](Get_Effect_Aff_plusAff()).V1}))}
	})
	return cache_Effect_Aff_plusParAff
}

var cache_Effect_Aff_alternativeParAff gopurs_runtime.Value
var once_Effect_Aff_alternativeParAff sync.Once

func Get_Effect_Aff_alternativeParAff() gopurs_runtime.Value {
	once_Effect_Aff_alternativeParAff.Do(func() {
		cache_Effect_Aff_alternativeParAff = gopurs_runtime.Value{Type: 9, IntVal: 397869517, UnsafePtr: unsafe.Pointer((&Constructor_Control_Alternative_Alternative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Effect_Aff_applicativeParAff()))}
		}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3709470893, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Plus_Plus[gopurs_runtime.Value]](Get_Effect_Aff_plusParAff()))}
		})}))}
	})
	return cache_Effect_Aff_alternativeParAff
}

func Call_Effect_Aff_bind__2699449543(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__2699449543:
	for {
		if false {
			continue bind__2699449543
		}
		var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			__local_var_2_0 := gopurs_runtime.Apply(__eta_norm_1_0, gopurs_runtime.Value{})
			_ = __local_var_2_0
			return gopurs_runtime.Apply2(__eta_norm_0_1, __local_var_2_0, gopurs_runtime.Value{})
		})
	}
}

func Call_Effect_Aff_pure__621669172(__eta_norm_0_0_loop struct {
	isSuspended gopurs_runtime.Value
	join        gopurs_runtime.Value
	kill        gopurs_runtime.Value
	onComplete  gopurs_runtime.Value
	run         gopurs_runtime.Value
	/*gopurs:record-v1:5:0069007300530075007300700065006e006400650064,006a006f0069006e,006b0069006c006c,006f006e0043006f006d0070006c006500740065,00720075006e*/
}) gopurs_runtime.Value {
pure__621669172:
	for {
		if false {
			continue pure__621669172
		}
		var __eta_norm_0_0 struct {
			isSuspended gopurs_runtime.Value
			join        gopurs_runtime.Value
			kill        gopurs_runtime.Value
			onComplete  gopurs_runtime.Value
			run         gopurs_runtime.Value
			/*gopurs:record-v1:5:0069007300530075007300700065006e006400650064,006a006f0069006e,006b0069006c006c,006f006e0043006f006d0070006c006500740065,00720075006e*/
		} = __eta_norm_0_0_loop
		_ = __eta_norm_0_0
		return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				orig := __eta_norm_0_0
				_ = orig
				return gopurs_runtime.RecordDict5("isSuspended", "join", "kill", "onComplete", "run", orig.isSuspended, orig.join, orig.kill, orig.onComplete, orig.run)
			}()
		})
	}
}

func Call_Effect_Aff_bind1__3918760327(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind1__3918760327:
	for {
		if false {
			continue bind1__3918760327
		}
		var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			__local_var_2_0 := gopurs_runtime.Apply(__eta_norm_1_0, gopurs_runtime.Value{})
			_ = __local_var_2_0
			return gopurs_runtime.Apply2(__eta_norm_0_1, __local_var_2_0, gopurs_runtime.Value{})
		})
	}
}

func Call_Effect_Aff_discard__1683979399(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
discard__1683979399:
	for {
		if false {
			continue discard__1683979399
		}
		var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			__local_var_2_0 := gopurs_runtime.Apply(__eta_norm_1_0, gopurs_runtime.Value{})
			_ = __local_var_2_0
			return gopurs_runtime.Apply2(__eta_norm_0_1, __local_var_2_0, gopurs_runtime.Value{})
		})
	}
}

func Call_Effect_Aff_pure1__621669172(__eta_norm_0_0_loop struct {
	isSuspended gopurs_runtime.Value
	join        gopurs_runtime.Value
	kill        gopurs_runtime.Value
	onComplete  gopurs_runtime.Value
	run         gopurs_runtime.Value
	/*gopurs:record-v1:5:0069007300530075007300700065006e006400650064,006a006f0069006e,006b0069006c006c,006f006e0043006f006d0070006c006500740065,00720075006e*/
}) gopurs_runtime.Value {
pure1__621669172:
	for {
		if false {
			continue pure1__621669172
		}
		var __eta_norm_0_0 struct {
			isSuspended gopurs_runtime.Value
			join        gopurs_runtime.Value
			kill        gopurs_runtime.Value
			onComplete  gopurs_runtime.Value
			run         gopurs_runtime.Value
			/*gopurs:record-v1:5:0069007300530075007300700065006e006400650064,006a006f0069006e,006b0069006c006c,006f006e0043006f006d0070006c006500740065,00720075006e*/
		} = __eta_norm_0_0_loop
		_ = __eta_norm_0_0
		return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				orig := __eta_norm_0_0
				_ = orig
				return gopurs_runtime.RecordDict5("isSuspended", "join", "kill", "onComplete", "run", orig.isSuspended, orig.join, orig.kill, orig.onComplete, orig.run)
			}()
		})
	}
}

func Call_Effect_Aff_compose1__4105490545(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose1__4105490545:
	for {
		if false {
			continue compose1__4105490545
		}
		var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
		_ = __eta_norm_1_unused_0
		var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Effect_Aff__liftEffect(), __eta_norm_0_1)
	}
}

func Call_Effect_Aff_compose1__1269036657(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose1__1269036657:
	for {
		if false {
			continue compose1__1269036657
		}
		var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
		_ = __eta_norm_1_unused_0
		var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Effect_Aff__liftEffect(), __eta_norm_0_1)
	}
}

func Call_Effect_Aff_Fiber(x_0_loop struct {
	isSuspended gopurs_runtime.Value
	join        gopurs_runtime.Value
	kill        gopurs_runtime.Value
	onComplete  gopurs_runtime.Value
	run         gopurs_runtime.Value
	/*gopurs:record-v1:5:0069007300530075007300700065006e006400650064,006a006f0069006e,006b0069006c006c,006f006e0043006f006d0070006c006500740065,00720075006e*/
}) struct {
	isSuspended gopurs_runtime.Value
	join        gopurs_runtime.Value
	kill        gopurs_runtime.Value
	onComplete  gopurs_runtime.Value
	run         gopurs_runtime.Value
	/*gopurs:record-v1:5:0069007300530075007300700065006e006400650064,006a006f0069006e,006b0069006c006c,006f006e0043006f006d0070006c006500740065,00720075006e*/
} {
	var x_0 struct {
		isSuspended gopurs_runtime.Value
		join        gopurs_runtime.Value
		kill        gopurs_runtime.Value
		onComplete  gopurs_runtime.Value
		run         gopurs_runtime.Value
		/*gopurs:record-v1:5:0069007300530075007300700065006e006400650064,006a006f0069006e,006b0069006c006c,006f006e0043006f006d0070006c006500740065,00720075006e*/
	} = x_0_loop
	_ = x_0
	return x_0
}

func Call_Effect_Aff_Fiber__288829658(x_0_loop struct {
	isSuspended gopurs_runtime.Value
	join        gopurs_runtime.Value
	kill        gopurs_runtime.Value
	onComplete  gopurs_runtime.Value
	run         gopurs_runtime.Value
	/*gopurs:record-v1:5:0069007300530075007300700065006e006400650064,006a006f0069006e,006b0069006c006c,006f006e0043006f006d0070006c006500740065,00720075006e*/
}) struct {
	isSuspended gopurs_runtime.Value
	join        gopurs_runtime.Value
	kill        gopurs_runtime.Value
	onComplete  gopurs_runtime.Value
	run         gopurs_runtime.Value
	/*gopurs:record-v1:5:0069007300530075007300700065006e006400650064,006a006f0069006e,006b0069006c006c,006f006e0043006f006d0070006c006500740065,00720075006e*/
} {
Fiber__288829658:
	for {
		if false {
			continue Fiber__288829658
		}
		var x_0 struct {
			isSuspended gopurs_runtime.Value
			join        gopurs_runtime.Value
			kill        gopurs_runtime.Value
			onComplete  gopurs_runtime.Value
			run         gopurs_runtime.Value
			/*gopurs:record-v1:5:0069007300530075007300700065006e006400650064,006a006f0069006e,006b0069006c006c,006f006e0043006f006d0070006c006500740065,00720075006e*/
		} = x_0_loop
		_ = x_0
		return x_0
	}
}

func Call_Effect_Aff_Canceler(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Effect_Aff_makeFiber(aff_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var aff_0 gopurs_runtime.Value = aff_0_loop
	_ = aff_0
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=Any
		__local_var_1_0 := gopurs_runtime.Apply(Get_Effect_Aff__makeFiberNative(), aff_0)
		_ = __local_var_1_0
		nf_2_1 := gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Value{})
		_ = nf_2_1
		return gopurs_runtime.RecordDict5("isSuspended", "join", "kill", "onComplete", "run", gopurs_runtime.Apply(Get_Effect_Aff__isSuspendedFiber(), nf_2_1), gopurs_runtime.Apply(Get_Effect_Aff__joinFiber(), nf_2_1), gopurs_runtime.Apply(Get_Effect_Aff__killFiber(), nf_2_1), gopurs_runtime.Apply(Get_Effect_Aff__onCompleteFiber(), nf_2_1), gopurs_runtime.Apply(Get_Effect_Aff__runFiber(), nf_2_1))
	})
}

func Call_Effect_Aff_makeFiber__4085038308(aff_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
makeFiber__4085038308:
	for {
		if false {
			continue makeFiber__4085038308
		}
		var aff_0 gopurs_runtime.Value = aff_0_loop
		_ = aff_0
		return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=Any
			__local_var_1_0 := gopurs_runtime.Apply(Get_Effect_Aff__makeFiberNative(), aff_0)
			_ = __local_var_1_0
			__local_var_2_1 := gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Value{})
			_ = __local_var_2_1
			return gopurs_runtime.RecordDict5("isSuspended", "join", "kill", "onComplete", "run", gopurs_runtime.Apply(Get_Effect_Aff__isSuspendedFiber(), __local_var_2_1), gopurs_runtime.Apply(Get_Effect_Aff__joinFiber(), __local_var_2_1), gopurs_runtime.Apply(Get_Effect_Aff__killFiber(), __local_var_2_1), gopurs_runtime.Apply(Get_Effect_Aff__onCompleteFiber(), __local_var_2_1), gopurs_runtime.Apply(Get_Effect_Aff__runFiber(), __local_var_2_1))
		})
	}
}

func Call_Effect_Aff_makeAff(build_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var build_0 gopurs_runtime.Value = build_0_loop
	_ = build_0
	return gopurs_runtime.Apply(Get_Effect_Aff__makeAffImpl(), gopurs_runtime.Func2(func(onError_1 gopurs_runtime.Value, onSuccess_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(build_0, gopurs_runtime.Func(func(either_3 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t0 gopurs_runtime.Value
			{
				if either_3.Type == 9 && either_3.IntVal == 3711209382 {
					__t0 = gopurs_runtime.Apply(onError_1, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(either_3.UnsafePtr).V0)
					goto end_branch_0
				} else {

				}
			}
			{
				if either_3.Type == 9 && either_3.IntVal == 2465973597 {
					__t0 = gopurs_runtime.Apply(onSuccess_2, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(either_3.UnsafePtr).V0)
					goto end_branch_0
				} else {

				}
			}
			{
				__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_0:
			return __t0
		}))
	}))
}

func Call_Effect_Aff_makeAff__1714332378(build_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
makeAff__1714332378:
	for {
		if false {
			continue makeAff__1714332378
		}
		var build_0 gopurs_runtime.Value = build_0_loop
		_ = build_0
		return gopurs_runtime.Apply(Get_Effect_Aff__makeAffImpl(), gopurs_runtime.Func2(func(onError_1 gopurs_runtime.Value, onSuccess_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(build_0, gopurs_runtime.Func(func(either_3 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t0 gopurs_runtime.Value
				{
					if either_3.Type == 9 && either_3.IntVal == 3711209382 {
						__t0 = gopurs_runtime.Apply(onError_1, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(either_3.UnsafePtr).V0)
						goto end_branch_0
					} else {

					}
				}
				{
					if either_3.Type == 9 && either_3.IntVal == 2465973597 {
						__t0 = gopurs_runtime.Apply(onSuccess_2, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(either_3.UnsafePtr).V0)
						goto end_branch_0
					} else {

					}
				}
				{
					__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
				}
			end_branch_0:
				return __t0
			}))
		}))
	}
}

func Call_Effect_Aff_makeAff__1501555834(build_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
makeAff__1501555834:
	for {
		if false {
			continue makeAff__1501555834
		}
		var build_0 gopurs_runtime.Value = build_0_loop
		_ = build_0
		return gopurs_runtime.Apply(Get_Effect_Aff__makeAffImpl(), gopurs_runtime.Func2(func(onError_1 gopurs_runtime.Value, onSuccess_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(build_0, gopurs_runtime.Func(func(either_3 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t0 gopurs_runtime.Value
				{
					if either_3.Type == 9 && either_3.IntVal == 3711209382 {
						__t0 = gopurs_runtime.Apply(onError_1, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(either_3.UnsafePtr).V0)
						goto end_branch_0
					} else {

					}
				}
				{
					if either_3.Type == 9 && either_3.IntVal == 2465973597 {
						__t0 = gopurs_runtime.Apply(onSuccess_2, gopurs_runtime.Array(*((*[]gopurs_runtime.Value)((*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(either_3.UnsafePtr).V0.UnsafePtr))))
						goto end_branch_0
					} else {

					}
				}
				{
					__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
				}
			end_branch_0:
				return __t0
			}))
		}))
	}
}

func Call_Effect_Aff_makeAff__2246750170(build_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
makeAff__2246750170:
	for {
		if false {
			continue makeAff__2246750170
		}
		var build_0 gopurs_runtime.Value = build_0_loop
		_ = build_0
		return gopurs_runtime.Apply(Get_Effect_Aff__makeAffImpl(), gopurs_runtime.Func2(func(onError_1 gopurs_runtime.Value, onSuccess_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(build_0, gopurs_runtime.Func(func(either_3 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t0 gopurs_runtime.Value
				{
					if either_3.Type == 9 && either_3.IntVal == 3711209382 {
						__t0 = gopurs_runtime.Apply(onError_1, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(either_3.UnsafePtr).V0)
						goto end_branch_0
					} else {

					}
				}
				{
					if either_3.Type == 9 && either_3.IntVal == 2465973597 {
						__t0 = gopurs_runtime.Apply(onSuccess_2, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(either_3.UnsafePtr).V0)
						goto end_branch_0
					} else {

					}
				}
				{
					__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
				}
			end_branch_0:
				return __t0
			}))
		}))
	}
}

func Call_Effect_Aff_makeAff__158558778(build_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
makeAff__158558778:
	for {
		if false {
			continue makeAff__158558778
		}
		var build_0 gopurs_runtime.Value = build_0_loop
		_ = build_0
		return gopurs_runtime.Apply(Get_Effect_Aff__makeAffImpl(), gopurs_runtime.Func2(func(onError_1 gopurs_runtime.Value, onSuccess_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(build_0, gopurs_runtime.Func(func(either_3 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t0 gopurs_runtime.Value
				{
					if either_3.Type == 9 && either_3.IntVal == 3711209382 {
						__t0 = gopurs_runtime.Apply(onError_1, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(either_3.UnsafePtr).V0)
						goto end_branch_0
					} else {

					}
				}
				{
					if either_3.Type == 9 && either_3.IntVal == 2465973597 {
						__t0 = gopurs_runtime.Apply(onSuccess_2, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(either_3.UnsafePtr).V0)
						goto end_branch_0
					} else {

					}
				}
				{
					__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
				}
			end_branch_0:
				return __t0
			}))
		}))
	}
}

func Call_Effect_Aff_makeAff__925755130(build_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
makeAff__925755130:
	for {
		if false {
			continue makeAff__925755130
		}
		var build_0 gopurs_runtime.Value = build_0_loop
		_ = build_0
		return gopurs_runtime.Apply(Get_Effect_Aff__makeAffImpl(), gopurs_runtime.Func2(func(onError_1 gopurs_runtime.Value, onSuccess_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(build_0, gopurs_runtime.Func(func(either_3 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t0 gopurs_runtime.Value
				{
					if either_3.Type == 9 && either_3.IntVal == 3711209382 {
						__t0 = gopurs_runtime.Apply(onError_1, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(either_3.UnsafePtr).V0)
						goto end_branch_0
					} else {

					}
				}
				{
					if either_3.Type == 9 && either_3.IntVal == 2465973597 {
						__t0 = gopurs_runtime.Apply(onSuccess_2, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(either_3.UnsafePtr).V0)
						goto end_branch_0
					} else {

					}
				}
				{
					__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
				}
			end_branch_0:
				return __t0
			}))
		}))
	}
}

func Call_Effect_Aff_launchSuspendedAff(aff_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var aff_0 gopurs_runtime.Value = aff_0_loop
	_ = aff_0
	return Call_Effect_Aff_makeFiber(aff_0)
}

func Call_Effect_Aff_launchSuspendedAff__4085038308(aff_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
launchSuspendedAff__4085038308:
	for {
		if false {
			continue launchSuspendedAff__4085038308
		}
		var aff_0 gopurs_runtime.Value = aff_0_loop
		_ = aff_0
		return Call_Effect_Aff_makeFiber__4085038308(aff_0)
	}
}

func Call_Effect_Aff_launchAff(aff_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var aff_0 gopurs_runtime.Value = aff_0_loop
	_ = aff_0
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=Any
		__local_var_1_0 := Call_Effect_Aff_makeFiber(aff_0)
		_ = __local_var_1_0
		fiber_2_1 := gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Value{})
		_ = fiber_2_1
		_dollar___unused_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(fiber_2_1, "run"), gopurs_runtime.Value{})
		_ = _dollar___unused_3_2
		return fiber_2_1
	})
}

func Call_Effect_Aff_launchAff__4085038308(aff_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
launchAff__4085038308:
	for {
		if false {
			continue launchAff__4085038308
		}
		var aff_0 gopurs_runtime.Value = aff_0_loop
		_ = aff_0
		return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=Any
			__local_var_1_0 := Call_Effect_Aff_makeFiber__4085038308(aff_0)
			_ = __local_var_1_0
			__local_var_2_1 := gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Value{})
			_ = __local_var_2_1
			__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "run"), gopurs_runtime.Value{})
			_ = __local_var_3_2
			return __local_var_2_1
		})
	}
}

func Call_Effect_Aff_delay(v_0_loop float64) gopurs_runtime.Value {
	var v_0 float64 = v_0_loop
	_ = v_0
	return gopurs_runtime.UncurriedApp2(Get_Effect_Aff__delay(), Get_Data_Either_Right(), gopurs_runtime.Float(v_0))
}

func Call_Effect_Aff_bracket(acquire_0_loop gopurs_runtime.Value, completed_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var acquire_0 gopurs_runtime.Value = acquire_0_loop
	_ = acquire_0
	var completed_1 gopurs_runtime.Value = completed_1_loop
	_ = completed_1
	return gopurs_runtime.Apply2(Get_Effect_Aff_generalBracket(), acquire_0, func() gopurs_runtime.Value {
		orig := struct {
			completed gopurs_runtime.Value
			failed    gopurs_runtime.Value
			killed    gopurs_runtime.Value
			/*gopurs:record-v1:3:0063006f006d0070006c0065007400650064,006600610069006c00650064,006b0069006c006c00650064*/
		}{gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return completed_1
		}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return completed_1
		}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return completed_1
		})}
		_ = orig
		return gopurs_runtime.RecordDict3("completed", "failed", "killed", orig.completed, orig.failed, orig.killed)
	}())
}

func Call_Effect_Aff_bracket__939099628(acquire_0_loop gopurs_runtime.Value, completed_unused_1_loop gopurs_runtime.Value, __eta_norm_0_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
bracket__939099628:
	for {
		if false {
			continue bracket__939099628
		}
		var acquire_0 gopurs_runtime.Value = acquire_0_loop
		_ = acquire_0
		var completed_unused_1 gopurs_runtime.Value = completed_unused_1_loop
		_ = completed_unused_1
		var __eta_norm_0_2 gopurs_runtime.Value = __eta_norm_0_2_loop
		_ = __eta_norm_0_2
		return gopurs_runtime.Apply3(Get_Effect_Aff_generalBracket(), acquire_0, func() gopurs_runtime.Value {
			orig := struct {
				completed gopurs_runtime.Value
				failed    gopurs_runtime.Value
				killed    gopurs_runtime.Value
				/*gopurs:record-v1:3:0063006f006d0070006c0065007400650064,006600610069006c00650064,006b0069006c006c00650064*/
			}{gopurs_runtime.Apply(Get_Data_Function_const__855710010(), Get_Database_SQLite3_go__close()), gopurs_runtime.Apply(Get_Data_Function_const__425793770(), Get_Database_SQLite3_go__close()), gopurs_runtime.Apply(Get_Data_Function_const__425793770(), Get_Database_SQLite3_go__close())}
			_ = orig
			return gopurs_runtime.RecordDict3("completed", "failed", "killed", orig.completed, orig.failed, orig.killed)
		}(), __eta_norm_0_2)
	}
}

func Call_Effect_Aff_bracket__3139827911(acquire_0_loop gopurs_runtime.Value, completed_unused_1_loop gopurs_runtime.Value, __eta_norm_0_unused_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
bracket__3139827911:
	for {
		if false {
			continue bracket__3139827911
		}
		var acquire_0 gopurs_runtime.Value = acquire_0_loop
		_ = acquire_0
		var completed_unused_1 gopurs_runtime.Value = completed_unused_1_loop
		_ = completed_unused_1
		var __eta_norm_0_unused_2 gopurs_runtime.Value = __eta_norm_0_unused_2_loop
		_ = __eta_norm_0_unused_2
		return gopurs_runtime.Apply3(Get_Effect_Aff_generalBracket(), acquire_0, func() gopurs_runtime.Value {
			orig := struct {
				completed gopurs_runtime.Value
				failed    gopurs_runtime.Value
				killed    gopurs_runtime.Value
				/*gopurs:record-v1:3:0063006f006d0070006c0065007400650064,006600610069006c00650064,006b0069006c006c00650064*/
			}{gopurs_runtime.Apply(Get_Data_Function_const__244128784(), gopurs_runtime.Apply2(Get_Example_HTTP_compose(), Get_Effect_Aff_Compat_fromEffectFnAff(), Get_Example_HTTP_stopImpl())), gopurs_runtime.Apply(Get_Data_Function_const__1353723968(), gopurs_runtime.Apply2(Get_Example_HTTP_compose(), Get_Effect_Aff_Compat_fromEffectFnAff(), Get_Example_HTTP_stopImpl())), gopurs_runtime.Apply(Get_Data_Function_const__1353723968(), gopurs_runtime.Apply2(Get_Example_HTTP_compose(), Get_Effect_Aff_Compat_fromEffectFnAff(), Get_Example_HTTP_stopImpl()))}
			_ = orig
			return gopurs_runtime.RecordDict3("completed", "failed", "killed", orig.completed, orig.failed, orig.killed)
		}(), gopurs_runtime.Apply2(Get_Example_HTTP_compose1(), Get_Effect_Aff_Compat_fromEffectFnAff(), Get_Example_HTTP_waitImpl()))
	}
}

func Call_Effect_Aff_lift2(f_0_loop gopurs_runtime.Value, a_1_loop gopurs_runtime.Value, b_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	var a_1 gopurs_runtime.Value = a_1_loop
	_ = a_1
	var b_2 gopurs_runtime.Value = b_2_loop
	_ = b_2
	return gopurs_runtime.Apply2(Get_Effect_Aff__parAffApply(), gopurs_runtime.Apply2(Get_Effect_Aff__parAffMap(), f_0, a_1), b_2)
}

func Call_Effect_Aff_semigroupParAff(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
	_ = dictSemigroup_0
	// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=(Func [(TypeVar a$scope1), (TypeVar a$scope1)] (TypeVar a$scope1))
	__local_var_1_0 := Call_Data_Semigroup_go__append(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](dictSemigroup_0))
	_ = __local_var_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer((&Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(a_2 gopurs_runtime.Value, b_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(Get_Effect_Aff__parAffApply(), gopurs_runtime.Apply2(Get_Effect_Aff__parAffMap(), __local_var_1_0, a_2), b_3)
	})}))}
}

func Call_Effect_Aff_lift21(f_0_loop gopurs_runtime.Value, a_1_loop gopurs_runtime.Value, b_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	var a_1 gopurs_runtime.Value = a_1_loop
	_ = a_1
	var b_2 gopurs_runtime.Value = b_2_loop
	_ = b_2
	return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Effect_Aff_applyAff()).V1, gopurs_runtime.Apply2(Get_Effect_Aff__map(), f_0, a_1), b_2)
}

func Call_Effect_Aff_bindFlipped(b_0_loop gopurs_runtime.Value, a_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var b_0 gopurs_runtime.Value = b_0_loop
	_ = b_0
	var a_1 gopurs_runtime.Value = a_1_loop
	_ = a_1
	return gopurs_runtime.Apply2(Get_Effect_Aff__bind(), a_1, b_0)
}

func Call_Effect_Aff_bindFlipped__879378959(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bindFlipped__879378959:
	for {
		if false {
			continue bindFlipped__879378959
		}
		var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		return gopurs_runtime.Apply2(Get_Effect_Aff__bind(), __eta_norm_0_1, __eta_norm_1_0)
	}
}

func Call_Effect_Aff_bindFlipped__52396111(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bindFlipped__52396111:
	for {
		if false {
			continue bindFlipped__52396111
		}
		var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		return gopurs_runtime.Apply2(Get_Effect_Aff__bind(), __eta_norm_0_1, __eta_norm_1_0)
	}
}

func Call_Effect_Aff_bindFlipped1(b_0_loop gopurs_runtime.Value, a_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var b_0 gopurs_runtime.Value = b_0_loop
	_ = b_0
	var a_1 gopurs_runtime.Value = a_1_loop
	_ = a_1
	return gopurs_runtime.Apply2(Get_Effect_Aff__bind(), a_1, b_0)
}

func Call_Effect_Aff_cancelWith(aff_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var aff_0 gopurs_runtime.Value = aff_0_loop
	_ = aff_0
	var v_1 gopurs_runtime.Value = v_1_loop
	_ = v_1
	return gopurs_runtime.Apply3(Get_Effect_Aff_generalBracket(), gopurs_runtime.Apply(Get_Effect_Aff__pure(), Get_Data_Unit_unit()), func() gopurs_runtime.Value {
		orig := struct {
			completed gopurs_runtime.Value
			failed    gopurs_runtime.Value
			killed    gopurs_runtime.Value
			/*gopurs:record-v1:3:0063006f006d0070006c0065007400650064,006600610069006c00650064,006b0069006c006c00650064*/
		}{gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return Get_Control_Applicative_pure__2964015333()
		}), gopurs_runtime.Apply(Get_Data_Function_const__1638250638(), Get_Control_Applicative_pure__2964015333()), gopurs_runtime.Func2(func(e_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(v_1, e_2)
		})}
		_ = orig
		return gopurs_runtime.RecordDict3("completed", "failed", "killed", orig.completed, orig.failed, orig.killed)
	}(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return aff_0
	}))
}

func Call_Effect_Aff_finally(fin_0_loop gopurs_runtime.Value, a_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var fin_0 gopurs_runtime.Value = fin_0_loop
	_ = fin_0
	var a_1 gopurs_runtime.Value = a_1_loop
	_ = a_1
	return gopurs_runtime.Apply3(Get_Effect_Aff_generalBracket(), gopurs_runtime.Apply(Get_Effect_Aff__pure(), Get_Data_Unit_unit()), func() gopurs_runtime.Value {
		orig := struct {
			completed gopurs_runtime.Value
			failed    gopurs_runtime.Value
			killed    gopurs_runtime.Value
			/*gopurs:record-v1:3:0063006f006d0070006c0065007400650064,006600610069006c00650064,006b0069006c006c00650064*/
		}{gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return fin_0
		}), gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return fin_0
		}), gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return fin_0
		})}
		_ = orig
		return gopurs_runtime.RecordDict3("completed", "failed", "killed", orig.completed, orig.failed, orig.killed)
	}(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return a_1
	}))
}

func Call_Effect_Aff_invincible(a_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var a_0 gopurs_runtime.Value = a_0_loop
	_ = a_0
	// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=Any
	__local_var_1_0 := gopurs_runtime.Apply(Get_Effect_Aff__pure(), Get_Data_Unit_unit())
	_ = __local_var_1_0
	return gopurs_runtime.Apply3(Get_Effect_Aff_generalBracket(), a_0, func() gopurs_runtime.Value {
		orig := struct {
			completed gopurs_runtime.Value
			failed    gopurs_runtime.Value
			killed    gopurs_runtime.Value
			/*gopurs:record-v1:3:0063006f006d0070006c0065007400650064,006600610069006c00650064,006b0069006c006c00650064*/
		}{gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return __local_var_1_0
		}), gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return __local_var_1_0
		}), gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return __local_var_1_0
		})}
		_ = orig
		return gopurs_runtime.RecordDict3("completed", "failed", "killed", orig.completed, orig.failed, orig.killed)
	}(), Get_Effect_Aff__pure())
}

func Call_Effect_Aff_monoidParAff(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
	_ = dictMonoid_0
	// TAST (Let): semigroupParAff1_1_0 shape=App(Var) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(ADT ["Effect","Aff","ParAff"] [(TypeVar a$scope28)])])
	semigroupParAff1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](Call_Effect_Aff_semigroupParAff(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})))
	_ = semigroupParAff1_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer((&Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(semigroupParAff1_1_0)}
	}), gopurs_runtime.Apply(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Effect_Aff_applicativeParAff()).V1, gopurs_runtime.RecordGet(dictMonoid_0, "mempty"))}))}
}

func Call_Effect_Aff_semigroupAff(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
	_ = dictSemigroup_0
	// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=(Func [(TypeVar a$scope5), (TypeVar a$scope5)] (TypeVar a$scope5))
	__local_var_1_0 := Call_Data_Semigroup_go__append(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](dictSemigroup_0))
	_ = __local_var_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer((&Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(a_2 gopurs_runtime.Value, b_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Effect_Aff_applyAff()).V1, gopurs_runtime.Apply2(Get_Effect_Aff__map(), __local_var_1_0, a_2), b_3)
	})}))}
}

func Call_Effect_Aff_joinFiber(v_0_loop struct {
	isSuspended gopurs_runtime.Value
	join        gopurs_runtime.Value
	kill        gopurs_runtime.Value
	onComplete  gopurs_runtime.Value
	run         gopurs_runtime.Value
	/*gopurs:record-v1:5:0069007300530075007300700065006e006400650064,006a006f0069006e,006b0069006c006c,006f006e0043006f006d0070006c006500740065,00720075006e*/
}) gopurs_runtime.Value {
	var v_0 struct {
		isSuspended gopurs_runtime.Value
		join        gopurs_runtime.Value
		kill        gopurs_runtime.Value
		onComplete  gopurs_runtime.Value
		run         gopurs_runtime.Value
		/*gopurs:record-v1:5:0069007300530075007300700065006e006400650064,006a006f0069006e,006b0069006c006c,006f006e0043006f006d0070006c006500740065,00720075006e*/
	} = v_0_loop
	_ = v_0
	return Call_Effect_Aff_makeAff(gopurs_runtime.Func(func(k_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return Call_Control_Applicative_liftA1(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Effect_applicativeEffect()), Get_Effect_Aff_effectCanceler(), gopurs_runtime.Apply2(v_0.join, gopurs_runtime.Func(func(err_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(k_1, func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 gopurs_runtime.Value
					V2 bool
				}{err_2, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}())
		}), gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(k_1, func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 gopurs_runtime.Value
					V2 bool
				}{gopurs_runtime.Value{}, a_2, true}
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}())
		})))
	}))
}

func Call_Effect_Aff_forkAff(aff_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var aff_0 gopurs_runtime.Value = aff_0_loop
	_ = aff_0
	return gopurs_runtime.Apply2(Get_Effect_Aff__bind(), gopurs_runtime.Apply(Get_Effect_Aff__forkAffNative(), aff_0), gopurs_runtime.Func(func(nf_1 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): fiber_2_0 shape=LitRecord bindingType=(Record (Row [run: (ADT ["Effect","Effect"] [Unit]), kill: (Func [(ADT ["Effect","Exception","Error"] []), (Func [(ADT ["Effect","Exception","Error"] [])] (ADT ["Effect","Effect"] [Unit])), (Func [Unit] (ADT ["Effect","Effect"] [Unit]))] (ADT ["Effect","Effect"] [(ADT ["Effect","Effect"] [Unit])])), join: (Func [(Func [(ADT ["Effect","Exception","Error"] [])] (ADT ["Effect","Effect"] [Unit])), (Func [(TypeVar a$scope106)] (ADT ["Effect","Effect"] [Unit]))] (ADT ["Effect","Effect"] [(ADT ["Effect","Effect"] [Unit])])), onComplete: (Func [(Record (Row [rethrow: Boolean, handler: (Func [(Func [(ADT ["Data","Either","Either"] [(ADT ["Effect","Exception","Error"] []), (TypeVar a$scope106)])] (ADT ["Effect","Effect"] [Unit]))] (ADT ["Effect","Effect"] [Unit]))] Empty))] (ADT ["Effect","Effect"] [(ADT ["Effect","Effect"] [Unit])])), isSuspended: (ADT ["Effect","Effect"] [Boolean])] Empty))
		fiber_2_0 := struct {
			isSuspended gopurs_runtime.Value
			join        gopurs_runtime.Value
			kill        gopurs_runtime.Value
			onComplete  gopurs_runtime.Value
			run         gopurs_runtime.Value
			/*gopurs:record-v1:5:0069007300530075007300700065006e006400650064,006a006f0069006e,006b0069006c006c,006f006e0043006f006d0070006c006500740065,00720075006e*/
		}{gopurs_runtime.Apply(Get_Effect_Aff__isSuspendedFiber(), nf_1), gopurs_runtime.Apply(Get_Effect_Aff__joinFiber(), nf_1), gopurs_runtime.Apply(Get_Effect_Aff__killFiber(), nf_1), gopurs_runtime.Apply(Get_Effect_Aff__onCompleteFiber(), nf_1), gopurs_runtime.Apply(Get_Effect_Aff__runFiber(), nf_1)}
		_ = fiber_2_0
		return gopurs_runtime.Apply2(Get_Effect_Aff__bind(), gopurs_runtime.Apply(Get_Effect_Aff__liftEffect(), gopurs_runtime.Apply(Get_Effect_Aff__runFiber(), nf_1)), gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(Get_Effect_Aff__pure(), func() gopurs_runtime.Value {
				orig := fiber_2_0
				_ = orig
				return gopurs_runtime.RecordDict5("isSuspended", "join", "kill", "onComplete", "run", orig.isSuspended, orig.join, orig.kill, orig.onComplete, orig.run)
			}())
		}))
	}))
}

func Call_Effect_Aff_killFiber(e_0_loop gopurs_runtime.Value, v_1_loop struct {
	isSuspended gopurs_runtime.Value
	join        gopurs_runtime.Value
	kill        gopurs_runtime.Value
	onComplete  gopurs_runtime.Value
	run         gopurs_runtime.Value
	/*gopurs:record-v1:5:0069007300530075007300700065006e006400650064,006a006f0069006e,006b0069006c006c,006f006e0043006f006d0070006c006500740065,00720075006e*/
}) gopurs_runtime.Value {
	var e_0 gopurs_runtime.Value = e_0_loop
	_ = e_0
	var v_1 struct {
		isSuspended gopurs_runtime.Value
		join        gopurs_runtime.Value
		kill        gopurs_runtime.Value
		onComplete  gopurs_runtime.Value
		run         gopurs_runtime.Value
		/*gopurs:record-v1:5:0069007300530075007300700065006e006400650064,006a006f0069006e,006b0069006c006c,006f006e0043006f006d0070006c006500740065,00720075006e*/
	} = v_1_loop
	_ = v_1
	return gopurs_runtime.Apply2(Get_Effect_Aff__bind(), gopurs_runtime.Apply(Get_Effect_Aff__liftEffect(), v_1.isSuspended), gopurs_runtime.Func(func(suspended_2 gopurs_runtime.Value) gopurs_runtime.Value {
		var __t0 gopurs_runtime.Value
		{
			if (suspended_2.IntVal) != (0) {
				__t0 = gopurs_runtime.Apply(Get_Effect_Aff__liftEffect(), gopurs_runtime.Apply(Call_Data_Functor_void(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect())), gopurs_runtime.Apply3(v_1.kill, e_0, gopurs_runtime.Apply(Get_Data_Function_const__2941505456(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
					return Get_Data_Unit_unit()
				})), gopurs_runtime.Apply(Get_Data_Function_const__2476124256(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
					return Get_Data_Unit_unit()
				})))))
				goto end_branch_0
			} else {

			}
		}
		{
			__t0 = Call_Effect_Aff_makeAff__925755130(gopurs_runtime.Func(func(k_3 gopurs_runtime.Value) gopurs_runtime.Value {
				return Call_Control_Applicative_liftA1(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Effect_applicativeEffect()), Get_Effect_Aff_effectCanceler(), gopurs_runtime.Apply3(v_1.kill, e_0, gopurs_runtime.Func(func(err_4 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply(k_3, func() gopurs_runtime.Value {
						_v := struct {
							V0 gopurs_runtime.Value
							V1 gopurs_runtime.Value
							V2 bool
						}{err_4, gopurs_runtime.Value{}, false}
						if _v.V2 {
							return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
						}
						return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
					}())
				}), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply(k_3, func() gopurs_runtime.Value {
						_v := struct {
							V0 gopurs_runtime.Value
							V1 gopurs_runtime.Value
							V2 bool
						}{gopurs_runtime.Value{}, Get_Data_Unit_unit(), true}
						if _v.V2 {
							return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
						}
						return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
					}())
				})))
			}))
		}
	end_branch_0:
		return __t0
	}))
}

func Call_Effect_Aff_supervise(aff_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var aff_0 gopurs_runtime.Value = aff_0_loop
	_ = aff_0
	// TAST (Let): killError_1_0 shape=App(Var) bindingType=(ADT ["Effect","Exception","Error"] [])
	killError_1_0 := gopurs_runtime.Apply(Get_Effect_Exception_error(), gopurs_runtime.StrWithCodeUnits("[Aff] Child fiber outlived parent", 33))
	_ = killError_1_0
	return gopurs_runtime.Apply3(Get_Effect_Aff_generalBracket(), gopurs_runtime.Apply(Get_Effect_Aff__liftEffect(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_2_1 shape=App(Var) bindingType=Any
		__local_var_2_1 := gopurs_runtime.Apply(Get_Effect_Aff__makeSupervisedFiber(), aff_0)
		_ = __local_var_2_1
		sup_3_2 := gopurs_runtime.Apply(__local_var_2_1, gopurs_runtime.Value{})
		_ = sup_3_2
		// TAST (Let): fiber_4_3 shape=LitRecord bindingType=(Record (Row [run: (ADT ["Effect","Effect"] [Unit]), kill: (Func [(ADT ["Effect","Exception","Error"] []), (Func [(ADT ["Effect","Exception","Error"] [])] (ADT ["Effect","Effect"] [Unit])), (Func [Unit] (ADT ["Effect","Effect"] [Unit]))] (ADT ["Effect","Effect"] [(ADT ["Effect","Effect"] [Unit])])), join: (Func [(Func [(ADT ["Effect","Exception","Error"] [])] (ADT ["Effect","Effect"] [Unit])), (Func [(TypeVar a$scope139)] (ADT ["Effect","Effect"] [Unit]))] (ADT ["Effect","Effect"] [(ADT ["Effect","Effect"] [Unit])])), onComplete: (Func [(Record (Row [rethrow: Boolean, handler: (Func [(Func [(ADT ["Data","Either","Either"] [(ADT ["Effect","Exception","Error"] []), (TypeVar a$scope139)])] (ADT ["Effect","Effect"] [Unit]))] (ADT ["Effect","Effect"] [Unit]))] Empty))] (ADT ["Effect","Effect"] [(ADT ["Effect","Effect"] [Unit])])), isSuspended: (ADT ["Effect","Effect"] [Boolean])] Empty))
		fiber_4_3 := struct {
			isSuspended gopurs_runtime.Value
			join        gopurs_runtime.Value
			kill        gopurs_runtime.Value
			onComplete  gopurs_runtime.Value
			run         gopurs_runtime.Value
			/*gopurs:record-v1:5:0069007300530075007300700065006e006400650064,006a006f0069006e,006b0069006c006c,006f006e0043006f006d0070006c006500740065,00720075006e*/
		}{gopurs_runtime.Apply(Get_Effect_Aff__isSuspendedFiber(), gopurs_runtime.RecordGet(sup_3_2, "fiber")), gopurs_runtime.Apply(Get_Effect_Aff__joinFiber(), gopurs_runtime.RecordGet(sup_3_2, "fiber")), gopurs_runtime.Apply(Get_Effect_Aff__killFiber(), gopurs_runtime.RecordGet(sup_3_2, "fiber")), gopurs_runtime.Apply(Get_Effect_Aff__onCompleteFiber(), gopurs_runtime.RecordGet(sup_3_2, "fiber")), gopurs_runtime.Apply(Get_Effect_Aff__runFiber(), gopurs_runtime.RecordGet(sup_3_2, "fiber"))}
		_ = fiber_4_3
		_dollar___unused_5_4 := gopurs_runtime.Apply(fiber_4_3.run, gopurs_runtime.Value{})
		_ = _dollar___unused_5_4
		return gopurs_runtime.RecordDict2("fiber", "supervisor", func() gopurs_runtime.Value {
			orig := fiber_4_3
			_ = orig
			return gopurs_runtime.RecordDict5("isSuspended", "join", "kill", "onComplete", "run", orig.isSuspended, orig.join, orig.kill, orig.onComplete, orig.run)
		}(), gopurs_runtime.RecordGet(sup_3_2, "supervisor"))
	})), func() gopurs_runtime.Value {
		orig := struct {
			completed gopurs_runtime.Value
			failed    gopurs_runtime.Value
			killed    gopurs_runtime.Value
			/*gopurs:record-v1:3:0063006f006d0070006c0065007400650064,006600610069006c00650064,006b0069006c006c00650064*/
		}{gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, sup_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Effect_Aff_makeAff__925755130(gopurs_runtime.Func(func(k_4 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.UncurriedApp3(Get_Effect_Aff__killAll(), killError_1_0, gopurs_runtime.RecordGet(sup_3, "supervisor"), gopurs_runtime.Apply(k_4, func() gopurs_runtime.Value {
					_v := struct {
						V0 gopurs_runtime.Value
						V1 gopurs_runtime.Value
						V2 bool
					}{gopurs_runtime.Value{}, Get_Data_Unit_unit(), true}
					if _v.V2 {
						return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
					}
					return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
				}()))
			}))
		}), gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, sup_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Effect_Aff_makeAff__925755130(gopurs_runtime.Func(func(k_4 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.UncurriedApp3(Get_Effect_Aff__killAll(), killError_1_0, gopurs_runtime.RecordGet(sup_3, "supervisor"), gopurs_runtime.Apply(k_4, func() gopurs_runtime.Value {
					_v := struct {
						V0 gopurs_runtime.Value
						V1 gopurs_runtime.Value
						V2 bool
					}{gopurs_runtime.Value{}, Get_Data_Unit_unit(), true}
					if _v.V2 {
						return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
					}
					return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
				}()))
			}))
		}), gopurs_runtime.Func2(func(err_2 gopurs_runtime.Value, sup_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply3(Call_Control_Parallel_parSequence_(gopurs_runtime.CoerceToStruct[Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Effect_Aff_parallelAff())), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Effect_Aff_applicativeParAff()))}, gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))}, gopurs_runtime.Array([]gopurs_runtime.Value{Call_Effect_Aff_killFiber(err_2, func() struct {
				isSuspended gopurs_runtime.Value
				join        gopurs_runtime.Value
				kill        gopurs_runtime.Value
				onComplete  gopurs_runtime.Value
				run         gopurs_runtime.Value
				/*gopurs:record-v1:5:0069007300530075007300700065006e006400650064,006a006f0069006e,006b0069006c006c,006f006e0043006f006d0070006c006500740065,00720075006e*/
			} {
				orig := gopurs_runtime.RecordGet(sup_3, "fiber")
				_ = orig
				clone := struct {
					isSuspended gopurs_runtime.Value
					join        gopurs_runtime.Value
					kill        gopurs_runtime.Value
					onComplete  gopurs_runtime.Value
					run         gopurs_runtime.Value
					/*gopurs:record-v1:5:0069007300530075007300700065006e006400650064,006a006f0069006e,006b0069006c006c,006f006e0043006f006d0070006c006500740065,00720075006e*/
				}{}
				clone.isSuspended = gopurs_runtime.RecordGet(orig, "isSuspended")
				clone.join = gopurs_runtime.RecordGet(orig, "join")
				clone.kill = gopurs_runtime.RecordGet(orig, "kill")
				clone.onComplete = gopurs_runtime.RecordGet(orig, "onComplete")
				clone.run = gopurs_runtime.RecordGet(orig, "run")
				return clone
			}()), Call_Effect_Aff_makeAff__925755130(gopurs_runtime.Func(func(k_4 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.UncurriedApp3(Get_Effect_Aff__killAll(), err_2, gopurs_runtime.RecordGet(sup_3, "supervisor"), gopurs_runtime.Apply(k_4, func() gopurs_runtime.Value {
					_v := struct {
						V0 gopurs_runtime.Value
						V1 gopurs_runtime.Value
						V2 bool
					}{gopurs_runtime.Value{}, Get_Data_Unit_unit(), true}
					if _v.V2 {
						return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
					}
					return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
				}()))
			}))}))
		})}
		_ = orig
		return gopurs_runtime.RecordDict3("completed", "failed", "killed", orig.completed, orig.failed, orig.killed)
	}(), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Effect_Aff_joinFiber(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.RecordGet(v_2, "fiber")
	})))
}

func Call_Effect_Aff_suspendAff(aff_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var aff_0 gopurs_runtime.Value = aff_0_loop
	_ = aff_0
	return gopurs_runtime.Apply(Get_Effect_Aff__liftEffect(), Call_Effect_Aff_makeFiber(aff_0))
}

func Call_Effect_Aff_try__4239810482(__eta_norm_0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
try__4239810482:
	for {
		if false {
			continue try__4239810482
		}
		var __eta_norm_0_0 gopurs_runtime.Value = __eta_norm_0_0_loop
		_ = __eta_norm_0_0
		return Call_Control_Monad_Error_Class_try__770583209(__eta_norm_0_0)
	}
}

func Call_Effect_Aff_try__1646464146(__eta_norm_0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
try__1646464146:
	for {
		if false {
			continue try__1646464146
		}
		var __eta_norm_0_0 gopurs_runtime.Value = __eta_norm_0_0_loop
		_ = __eta_norm_0_0
		return Call_Control_Monad_Error_Class_try__2899663753(__eta_norm_0_0)
	}
}

func Call_Effect_Aff_runAff(k_0_loop gopurs_runtime.Value, aff_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var k_0 gopurs_runtime.Value = k_0_loop
	_ = k_0
	var aff_1 gopurs_runtime.Value = aff_1_loop
	_ = aff_1
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=Any
		__local_var_2_0 := Call_Effect_Aff_makeFiber__4085038308(gopurs_runtime.Apply2(Get_Effect_Aff__bind(), gopurs_runtime.Apply(Call_Control_Monad_Error_Class_try(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Error_Class_MonadError[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Effect_Aff_monadErrorAff())), aff_1), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Effect_Aff__liftEffect(), k_0)))
		_ = __local_var_2_0
		__local_var_3_1 := gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Value{})
		_ = __local_var_3_1
		__local_var_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_1, "run"), gopurs_runtime.Value{})
		_ = __local_var_4_2
		return __local_var_3_1
	})
}

func Call_Effect_Aff_runAff__2603932529(k_0_loop gopurs_runtime.Value, aff_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
runAff__2603932529:
	for {
		if false {
			continue runAff__2603932529
		}
		var k_0 gopurs_runtime.Value = k_0_loop
		_ = k_0
		var aff_1 gopurs_runtime.Value = aff_1_loop
		_ = aff_1
		return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=Any
			__local_var_2_0 := Call_Effect_Aff_makeFiber__4085038308(gopurs_runtime.Apply2(Get_Effect_Aff__bind(), Call_Control_Monad_Error_Class_try__770583209(aff_1), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Effect_Aff__liftEffect(), k_0)))
			_ = __local_var_2_0
			__local_var_3_1 := gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Value{})
			_ = __local_var_3_1
			__local_var_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_1, "run"), gopurs_runtime.Value{})
			_ = __local_var_4_2
			return __local_var_3_1
		})
	}
}

func Call_Effect_Aff_runAff__2939864241(k_0_loop gopurs_runtime.Value, aff_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
runAff__2939864241:
	for {
		if false {
			continue runAff__2939864241
		}
		var k_0 gopurs_runtime.Value = k_0_loop
		_ = k_0
		var aff_1 gopurs_runtime.Value = aff_1_loop
		_ = aff_1
		return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=Any
			__local_var_2_0 := Call_Effect_Aff_makeFiber__4085038308(gopurs_runtime.Apply2(Get_Effect_Aff__bind(), Call_Control_Monad_Error_Class_try__2899663753(aff_1), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Effect_Aff__liftEffect(), k_0)))
			_ = __local_var_2_0
			__local_var_3_1 := gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Value{})
			_ = __local_var_3_1
			__local_var_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_1, "run"), gopurs_runtime.Value{})
			_ = __local_var_4_2
			return __local_var_3_1
		})
	}
}

func Call_Effect_Aff_runAff_(k_0_loop gopurs_runtime.Value, aff_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var k_0 gopurs_runtime.Value = k_0_loop
	_ = k_0
	var aff_1 gopurs_runtime.Value = aff_1_loop
	_ = aff_1
	return gopurs_runtime.Apply(Call_Data_Functor_void(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect())), Call_Effect_Aff_runAff(k_0, aff_1))
}

func Call_Effect_Aff_runAff___3349181959(k_0_loop gopurs_runtime.Value, aff_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
runAff___3349181959:
	for {
		if false {
			continue runAff___3349181959
		}
		var k_0 gopurs_runtime.Value = k_0_loop
		_ = k_0
		var aff_1 gopurs_runtime.Value = aff_1_loop
		_ = aff_1
		return gopurs_runtime.Apply(Call_Data_Functor_void(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect())), Call_Effect_Aff_runAff__2603932529(k_0, aff_1))
	}
}

func Call_Effect_Aff_runAff___358700103(k_0_loop gopurs_runtime.Value, aff_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
runAff___358700103:
	for {
		if false {
			continue runAff___358700103
		}
		var k_0 gopurs_runtime.Value = k_0_loop
		_ = k_0
		var aff_1 gopurs_runtime.Value = aff_1_loop
		_ = aff_1
		return gopurs_runtime.Apply(Call_Data_Functor_void(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect())), Call_Effect_Aff_runAff__2939864241(k_0, aff_1))
	}
}

func Call_Effect_Aff_runSuspendedAff(k_0_loop gopurs_runtime.Value, aff_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var k_0 gopurs_runtime.Value = k_0_loop
	_ = k_0
	var aff_1 gopurs_runtime.Value = aff_1_loop
	_ = aff_1
	return Call_Effect_Aff_makeFiber__4085038308(gopurs_runtime.Apply2(Get_Effect_Aff__bind(), gopurs_runtime.Apply(Call_Control_Monad_Error_Class_try(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Error_Class_MonadError[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Effect_Aff_monadErrorAff())), aff_1), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Effect_Aff__liftEffect(), k_0)))
}

func Call_Effect_Aff_monoidAff(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
	_ = dictMonoid_0
	// TAST (Let): semigroupAff1_1_0 shape=App(Var) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(ADT ["Effect","Aff","Aff"] [(TypeVar a$scope32)])])
	semigroupAff1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](Call_Effect_Aff_semigroupAff(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})))
	_ = semigroupAff1_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer((&Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(semigroupAff1_1_0)}
	}), gopurs_runtime.Apply(Get_Effect_Aff__pure(), gopurs_runtime.RecordGet(dictMonoid_0, "mempty"))}))}
}

func Rebox_Effect_Aff_179938118_2812149806(in *Constructor_Data_Functor_Functor[struct {
	isSuspended gopurs_runtime.Value
	join        gopurs_runtime.Value
	kill        gopurs_runtime.Value
	onComplete  gopurs_runtime.Value
	run         gopurs_runtime.Value
	/*gopurs:record-v1:5:0069007300530075007300700065006e006400650064,006a006f0069006e,006b0069006c006c,006f006e0043006f006d0070006c006500740065,00720075006e*/
}]) *Constructor_Data_Functor_Functor[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Functor_Functor[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Effect_Aff_1844969297_3741347833(in *Constructor_Control_Apply_Apply[struct {
	isSuspended gopurs_runtime.Value
	join        gopurs_runtime.Value
	kill        gopurs_runtime.Value
	onComplete  gopurs_runtime.Value
	run         gopurs_runtime.Value
	/*gopurs:record-v1:5:0069007300530075007300700065006e006400650064,006a006f0069006e,006b0069006c006c,006f006e0043006f006d0070006c006500740065,00720075006e*/
}]) *Constructor_Control_Apply_Apply[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Control_Apply_Apply[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Effect_Aff_2812149806_179938118(in *Constructor_Data_Functor_Functor[gopurs_runtime.Value]) *Constructor_Data_Functor_Functor[struct {
	isSuspended gopurs_runtime.Value
	join        gopurs_runtime.Value
	kill        gopurs_runtime.Value
	onComplete  gopurs_runtime.Value
	run         gopurs_runtime.Value
	/*gopurs:record-v1:5:0069007300530075007300700065006e006400650064,006a006f0069006e,006b0069006c006c,006f006e0043006f006d0070006c006500740065,00720075006e*/
}] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Functor_Functor[struct {
		isSuspended gopurs_runtime.Value
		join        gopurs_runtime.Value
		kill        gopurs_runtime.Value
		onComplete  gopurs_runtime.Value
		run         gopurs_runtime.Value
		/*gopurs:record-v1:5:0069007300530075007300700065006e006400650064,006a006f0069006e,006b0069006c006c,006f006e0043006f006d0070006c006500740065,00720075006e*/
	}])(unsafe.Pointer(in))
}

func Rebox_Effect_Aff_3741347833_1844969297(in *Constructor_Control_Apply_Apply[gopurs_runtime.Value]) *Constructor_Control_Apply_Apply[struct {
	isSuspended gopurs_runtime.Value
	join        gopurs_runtime.Value
	kill        gopurs_runtime.Value
	onComplete  gopurs_runtime.Value
	run         gopurs_runtime.Value
	/*gopurs:record-v1:5:0069007300530075007300700065006e006400650064,006a006f0069006e,006b0069006c006c,006f006e0043006f006d0070006c006500740065,00720075006e*/
}] {
	if in == nil {
		return nil
	}
	return (*Constructor_Control_Apply_Apply[struct {
		isSuspended gopurs_runtime.Value
		join        gopurs_runtime.Value
		kill        gopurs_runtime.Value
		onComplete  gopurs_runtime.Value
		run         gopurs_runtime.Value
		/*gopurs:record-v1:5:0069007300530075007300700065006e006400650064,006a006f0069006e,006b0069006c006c,006f006e0043006f006d0070006c006500740065,00720075006e*/
	}])(unsafe.Pointer(in))
}

func Rebox_Effect_Aff_520806353_1439734649(in *Constructor_Control_Applicative_Applicative[struct {
	isSuspended gopurs_runtime.Value
	join        gopurs_runtime.Value
	kill        gopurs_runtime.Value
	onComplete  gopurs_runtime.Value
	run         gopurs_runtime.Value
	/*gopurs:record-v1:5:0069007300530075007300700065006e006400650064,006a006f0069006e,006b0069006c006c,006f006e0043006f006d0070006c006500740065,00720075006e*/
}]) *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Control_Applicative_Applicative[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Get_Effect_Aff__bind() gopurs_runtime.Value {
	return _Gopurs_Effect_Aff__Bind
}

func Get_Effect_Aff__catchError() gopurs_runtime.Value {
	return _Gopurs_Effect_Aff__CatchError
}

func Get_Effect_Aff__delay() gopurs_runtime.Value {
	return _Gopurs_Effect_Aff__Delay
}

func Get_Effect_Aff__forkAffNative() gopurs_runtime.Value {
	return _Gopurs_Effect_Aff__ForkAffNative
}

func Get_Effect_Aff__isSuspendedFiber() gopurs_runtime.Value {
	return _Gopurs_Effect_Aff__IsSuspendedFiber
}

func Get_Effect_Aff__joinFiber() gopurs_runtime.Value {
	return _Gopurs_Effect_Aff__JoinFiber
}

func Get_Effect_Aff__killAll() gopurs_runtime.Value {
	return _Gopurs_Effect_Aff__KillAll
}

func Get_Effect_Aff__killFiber() gopurs_runtime.Value {
	return _Gopurs_Effect_Aff__KillFiber
}

func Get_Effect_Aff__liftEffect() gopurs_runtime.Value {
	return _Gopurs_Effect_Aff__LiftEffect
}

func Get_Effect_Aff__makeAffImpl() gopurs_runtime.Value {
	return _Gopurs_Effect_Aff__MakeAffImpl
}

func Get_Effect_Aff__makeFiberNative() gopurs_runtime.Value {
	return _Gopurs_Effect_Aff__MakeFiberNative
}

func Get_Effect_Aff__makeSupervisedFiber() gopurs_runtime.Value {
	return _Gopurs_Effect_Aff__MakeSupervisedFiber
}

func Get_Effect_Aff__map() gopurs_runtime.Value {
	return _Gopurs_Effect_Aff__Map
}

func Get_Effect_Aff__onCompleteFiber() gopurs_runtime.Value {
	return _Gopurs_Effect_Aff__OnCompleteFiber
}

func Get_Effect_Aff__parAffAlt() gopurs_runtime.Value {
	return _Gopurs_Effect_Aff__ParAffAlt
}

func Get_Effect_Aff__parAffApply() gopurs_runtime.Value {
	return _Gopurs_Effect_Aff__ParAffApply
}

func Get_Effect_Aff__parAffMap() gopurs_runtime.Value {
	return _Gopurs_Effect_Aff__ParAffMap
}

func Get_Effect_Aff__pure() gopurs_runtime.Value {
	return _Gopurs_Effect_Aff__Pure
}

func Get_Effect_Aff__runFiber() gopurs_runtime.Value {
	return _Gopurs_Effect_Aff__RunFiber
}

func Get_Effect_Aff__sequential() gopurs_runtime.Value {
	return _Gopurs_Effect_Aff__Sequential
}

func Get_Effect_Aff__throwError() gopurs_runtime.Value {
	return _Gopurs_Effect_Aff__ThrowError
}

func Get_Effect_Aff_generalBracket() gopurs_runtime.Value {
	return _Gopurs_Effect_Aff_GeneralBracket
}
