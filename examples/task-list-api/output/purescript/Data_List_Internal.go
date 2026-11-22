package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_List_Internal_Leaf gopurs_runtime.Value
var once_Data_List_Internal_Leaf sync.Once

func Get_Data_List_Internal_Leaf() gopurs_runtime.Value {
	once_Data_List_Internal_Leaf.Do(func() {
		cache_Data_List_Internal_Leaf = gopurs_runtime.Value{Type: 9, IntVal: 2764020654, UnsafePtr: unsafe.Pointer(nil)}
	})
	return cache_Data_List_Internal_Leaf
}

var cache_Data_List_Internal_Two gopurs_runtime.Value
var once_Data_List_Internal_Two sync.Once

func Get_Data_List_Internal_Two() gopurs_runtime.Value {
	once_Data_List_Internal_Two.Do(func() {
		cache_Data_List_Internal_Two = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(value2 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Value{Type: 9, IntVal: 1177901036, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Internal_Two[gopurs_runtime.Value]{1, value0, value1, value2}))}
				})
			})
		})
	})
	return cache_Data_List_Internal_Two
}

var cache_Data_List_Internal_Three gopurs_runtime.Value
var once_Data_List_Internal_Three sync.Once

func Get_Data_List_Internal_Three() gopurs_runtime.Value {
	once_Data_List_Internal_Three.Do(func() {
		cache_Data_List_Internal_Three = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(value2 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(value3 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(value4 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Value{Type: 9, IntVal: 1064476974, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Internal_Three[gopurs_runtime.Value]{1, value0, value1, value2, value3, value4}))}
						})
					})
				})
			})
		})
	})
	return cache_Data_List_Internal_Three
}

var cache_Data_List_Internal_TwoLeft gopurs_runtime.Value
var once_Data_List_Internal_TwoLeft sync.Once

func Get_Data_List_Internal_TwoLeft() gopurs_runtime.Value {
	once_Data_List_Internal_TwoLeft.Do(func() {
		cache_Data_List_Internal_TwoLeft = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 1304506903, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Internal_TwoLeft[gopurs_runtime.Value]{1, value0, value1}))}
			})
		})
	})
	return cache_Data_List_Internal_TwoLeft
}

var cache_Data_List_Internal_TwoRight gopurs_runtime.Value
var once_Data_List_Internal_TwoRight sync.Once

func Get_Data_List_Internal_TwoRight() gopurs_runtime.Value {
	once_Data_List_Internal_TwoRight.Do(func() {
		cache_Data_List_Internal_TwoRight = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 2884341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Internal_TwoRight[gopurs_runtime.Value]{1, value0, value1}))}
			})
		})
	})
	return cache_Data_List_Internal_TwoRight
}

var cache_Data_List_Internal_ThreeLeft gopurs_runtime.Value
var once_Data_List_Internal_ThreeLeft sync.Once

func Get_Data_List_Internal_ThreeLeft() gopurs_runtime.Value {
	once_Data_List_Internal_ThreeLeft.Do(func() {
		cache_Data_List_Internal_ThreeLeft = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(value2 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(value3 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Value{Type: 9, IntVal: 2195694037, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Internal_ThreeLeft[gopurs_runtime.Value]{1, value0, value1, value2, value3}))}
					})
				})
			})
		})
	})
	return cache_Data_List_Internal_ThreeLeft
}

var cache_Data_List_Internal_ThreeMiddle gopurs_runtime.Value
var once_Data_List_Internal_ThreeMiddle sync.Once

func Get_Data_List_Internal_ThreeMiddle() gopurs_runtime.Value {
	once_Data_List_Internal_ThreeMiddle.Do(func() {
		cache_Data_List_Internal_ThreeMiddle = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(value2 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(value3 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Value{Type: 9, IntVal: 1584522659, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Internal_ThreeMiddle[gopurs_runtime.Value]{1, value0, value1, value2, value3}))}
					})
				})
			})
		})
	})
	return cache_Data_List_Internal_ThreeMiddle
}

var cache_Data_List_Internal_ThreeRight gopurs_runtime.Value
var once_Data_List_Internal_ThreeRight sync.Once

func Get_Data_List_Internal_ThreeRight() gopurs_runtime.Value {
	once_Data_List_Internal_ThreeRight.Do(func() {
		cache_Data_List_Internal_ThreeRight = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(value2 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(value3 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Value{Type: 9, IntVal: 3952671150, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Internal_ThreeRight[gopurs_runtime.Value]{1, value0, value1, value2, value3}))}
					})
				})
			})
		})
	})
	return cache_Data_List_Internal_ThreeRight
}

var cache_Data_List_Internal_KickUp gopurs_runtime.Value
var once_Data_List_Internal_KickUp sync.Once

func Get_Data_List_Internal_KickUp() gopurs_runtime.Value {
	once_Data_List_Internal_KickUp.Do(func() {
		cache_Data_List_Internal_KickUp = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(value2 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Value{Type: 9, IntVal: 2023586927, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Internal_KickUp[gopurs_runtime.Value]{1, value0, value1, value2}))}
				})
			})
		})
	})
	return cache_Data_List_Internal_KickUp
}

var cache_Data_List_Internal_fromZipper gopurs_runtime.Value
var once_Data_List_Internal_fromZipper sync.Once

func Get_Data_List_Internal_fromZipper() gopurs_runtime.Value {
	once_Data_List_Internal_fromZipper.Do(func() {
		cache_Data_List_Internal_fromZipper = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_Internal_fromZipper(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_0_box), v1_1_box)
		})
	})
	return cache_Data_List_Internal_fromZipper
}

var cache_Data_List_Internal_insertAndLookupBy gopurs_runtime.Value
var once_Data_List_Internal_insertAndLookupBy sync.Once

func Get_Data_List_Internal_insertAndLookupBy() gopurs_runtime.Value {
	once_Data_List_Internal_insertAndLookupBy.Do(func() {
		cache_Data_List_Internal_insertAndLookupBy = gopurs_runtime.Func3(func(comp_0_box gopurs_runtime.Value, k_1_box gopurs_runtime.Value, orig_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				orig := Call_Data_List_Internal_insertAndLookupBy(comp_0_box, k_1_box, orig_2_box)
				_ = orig
				return gopurs_runtime.RecordDict2("found", "result", gopurs_runtime.Bool(orig.found), orig.result)
			}()
		})
	})
	return cache_Data_List_Internal_insertAndLookupBy
}

var cache_Data_List_Internal_emptySet gopurs_runtime.Value
var once_Data_List_Internal_emptySet sync.Once

func Get_Data_List_Internal_emptySet() gopurs_runtime.Value {
	once_Data_List_Internal_emptySet.Do(func() {
		cache_Data_List_Internal_emptySet = gopurs_runtime.Value{Type: 9, IntVal: 2764020654, UnsafePtr: unsafe.Pointer(nil)}
	})
	return cache_Data_List_Internal_emptySet
}

type Constructor_Data_List_Internal_Leaf[T_k any] struct {
	Rc uint32
}

type Constructor_Data_List_Internal_Two[T_k any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 T_k
	V2 gopurs_runtime.Value
}

type Constructor_Data_List_Internal_Three[T_k any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 T_k
	V2 gopurs_runtime.Value
	V3 T_k
	V4 gopurs_runtime.Value
}

type Constructor_Data_List_Internal_TwoLeft[T_k any] struct {
	Rc uint32
	V0 T_k
	V1 gopurs_runtime.Value
}

type Constructor_Data_List_Internal_TwoRight[T_k any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 T_k
}

type Constructor_Data_List_Internal_ThreeLeft[T_k any] struct {
	Rc uint32
	V0 T_k
	V1 gopurs_runtime.Value
	V2 T_k
	V3 gopurs_runtime.Value
}

type Constructor_Data_List_Internal_ThreeMiddle[T_k any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 T_k
	V2 T_k
	V3 gopurs_runtime.Value
}

type Constructor_Data_List_Internal_ThreeRight[T_k any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 T_k
	V2 gopurs_runtime.Value
	V3 T_k
}

type Constructor_Data_List_Internal_KickUp[T_k any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 T_k
	V2 gopurs_runtime.Value
}

func Call_Data_List_Internal_fromZipper(v_0_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
fromZipper:
	for {
		if false {
			continue fromZipper
		}
		var v_0 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_0_loop
		_ = v_0
		var v1_1 gopurs_runtime.Value = v1_1_loop
		_ = v1_1
		var __t6 gopurs_runtime.Value
		{
			if v_0 == nil {
				__t6 = v1_1
				goto end_branch_6
			} else {

			}
		}
		{
			if v_0 != nil {
				var __t5 gopurs_runtime.Value
				{
					var __t_tag_0 gopurs_runtime.Value = (v_0).V0
					_ = __t_tag_0
					if __t_tag_0.Type == 9 && __t_tag_0.IntVal == 1304506903 {
						v_0_loop = (v_0).V1
						v1_1_loop = gopurs_runtime.Value{Type: 9, IntVal: 1177901036, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Internal_Two[gopurs_runtime.Value]{1, v1_1, (*Constructor_Data_List_Internal_TwoLeft[gopurs_runtime.Value])((v_0).V0.UnsafePtr).V0, (*Constructor_Data_List_Internal_TwoLeft[gopurs_runtime.Value])((v_0).V0.UnsafePtr).V1}))}
						continue fromZipper
						__t5 = func() gopurs_runtime.Value { panic("unreachable") }()
						goto end_branch_5
					} else {

					}
				}
				{
					var __t_tag_1 gopurs_runtime.Value = (v_0).V0
					_ = __t_tag_1
					if __t_tag_1.Type == 9 && __t_tag_1.IntVal == 2884341868 {
						v_0_loop = (v_0).V1
						v1_1_loop = gopurs_runtime.Value{Type: 9, IntVal: 1177901036, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Internal_Two[gopurs_runtime.Value]{1, (*Constructor_Data_List_Internal_TwoRight[gopurs_runtime.Value])((v_0).V0.UnsafePtr).V0, (*Constructor_Data_List_Internal_TwoRight[gopurs_runtime.Value])((v_0).V0.UnsafePtr).V1, v1_1}))}
						continue fromZipper
						__t5 = func() gopurs_runtime.Value { panic("unreachable") }()
						goto end_branch_5
					} else {

					}
				}
				{
					var __t_tag_2 gopurs_runtime.Value = (v_0).V0
					_ = __t_tag_2
					if __t_tag_2.Type == 9 && __t_tag_2.IntVal == 2195694037 {
						v_0_loop = (v_0).V1
						v1_1_loop = gopurs_runtime.Value{Type: 9, IntVal: 1064476974, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Internal_Three[gopurs_runtime.Value]{1, v1_1, (*Constructor_Data_List_Internal_ThreeLeft[gopurs_runtime.Value])((v_0).V0.UnsafePtr).V0, (*Constructor_Data_List_Internal_ThreeLeft[gopurs_runtime.Value])((v_0).V0.UnsafePtr).V1, (*Constructor_Data_List_Internal_ThreeLeft[gopurs_runtime.Value])((v_0).V0.UnsafePtr).V2, (*Constructor_Data_List_Internal_ThreeLeft[gopurs_runtime.Value])((v_0).V0.UnsafePtr).V3}))}
						continue fromZipper
						__t5 = func() gopurs_runtime.Value { panic("unreachable") }()
						goto end_branch_5
					} else {

					}
				}
				{
					var __t_tag_3 gopurs_runtime.Value = (v_0).V0
					_ = __t_tag_3
					if __t_tag_3.Type == 9 && __t_tag_3.IntVal == 1584522659 {
						v_0_loop = (v_0).V1
						v1_1_loop = gopurs_runtime.Value{Type: 9, IntVal: 1064476974, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Internal_Three[gopurs_runtime.Value]{1, (*Constructor_Data_List_Internal_ThreeMiddle[gopurs_runtime.Value])((v_0).V0.UnsafePtr).V0, (*Constructor_Data_List_Internal_ThreeMiddle[gopurs_runtime.Value])((v_0).V0.UnsafePtr).V1, v1_1, (*Constructor_Data_List_Internal_ThreeMiddle[gopurs_runtime.Value])((v_0).V0.UnsafePtr).V2, (*Constructor_Data_List_Internal_ThreeMiddle[gopurs_runtime.Value])((v_0).V0.UnsafePtr).V3}))}
						continue fromZipper
						__t5 = func() gopurs_runtime.Value { panic("unreachable") }()
						goto end_branch_5
					} else {

					}
				}
				{
					var __t_tag_4 gopurs_runtime.Value = (v_0).V0
					_ = __t_tag_4
					if __t_tag_4.Type == 9 && __t_tag_4.IntVal == 3952671150 {
						v_0_loop = (v_0).V1
						v1_1_loop = gopurs_runtime.Value{Type: 9, IntVal: 1064476974, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Internal_Three[gopurs_runtime.Value]{1, (*Constructor_Data_List_Internal_ThreeRight[gopurs_runtime.Value])((v_0).V0.UnsafePtr).V0, (*Constructor_Data_List_Internal_ThreeRight[gopurs_runtime.Value])((v_0).V0.UnsafePtr).V1, (*Constructor_Data_List_Internal_ThreeRight[gopurs_runtime.Value])((v_0).V0.UnsafePtr).V2, (*Constructor_Data_List_Internal_ThreeRight[gopurs_runtime.Value])((v_0).V0.UnsafePtr).V3, v1_1}))}
						continue fromZipper
						__t5 = func() gopurs_runtime.Value { panic("unreachable") }()
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
			__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
		}
	end_branch_6:
		return __t6
	}
}

func Call_Data_List_Internal_insertAndLookupBy(comp_0_loop gopurs_runtime.Value, k_1_loop gopurs_runtime.Value, orig_2_loop gopurs_runtime.Value) struct {
	found  bool
	result gopurs_runtime.Value
	/*gopurs:record-v1:2:0066006f0075006e0064,0072006500730075006c0074*/
} {
	var comp_0 gopurs_runtime.Value = comp_0_loop
	_ = comp_0
	var k_1 gopurs_runtime.Value = k_1_loop
	_ = k_1
	var orig_2 gopurs_runtime.Value = orig_2_loop
	_ = orig_2
	var Call_local_Data_List_Internal_up_3_0_0 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Internal_KickUp[gopurs_runtime.Value]) gopurs_runtime.Value
	_ = Call_local_Data_List_Internal_up_3_0_0
	var up_3_0_0 gopurs_runtime.Value
	_ = up_3_0_0
	Call_local_Data_List_Internal_up_3_0_0 = func(v_4_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_5_loop *Constructor_Data_List_Internal_KickUp[gopurs_runtime.Value]) gopurs_runtime.Value {
	up_3_0_0:
		for {
			if false {
				continue up_3_0_0
			}
			var v_4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_4_loop
			_ = v_4
			var v1_5 *Constructor_Data_List_Internal_KickUp[gopurs_runtime.Value] = v1_5_loop
			_ = v1_5
			var __t7 gopurs_runtime.Value
			{
				if v_4 == nil {
					__t7 = gopurs_runtime.Value{Type: 9, IntVal: 1177901036, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Internal_Two[gopurs_runtime.Value]{1, (v1_5).V0, (v1_5).V1, (v1_5).V2}))}
					goto end_branch_7
				} else {

				}
			}
			{
				if v_4 != nil {
					var __t6 gopurs_runtime.Value
					{
						var __t_tag_1 gopurs_runtime.Value = (v_4).V0
						_ = __t_tag_1
						if __t_tag_1.Type == 9 && __t_tag_1.IntVal == 1304506903 {
							__t6 = Call_Data_List_Internal_fromZipper((v_4).V1, gopurs_runtime.Value{Type: 9, IntVal: 1064476974, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Internal_Three[gopurs_runtime.Value]{1, (v1_5).V0, (v1_5).V1, (v1_5).V2, (*Constructor_Data_List_Internal_TwoLeft[gopurs_runtime.Value])((v_4).V0.UnsafePtr).V0, (*Constructor_Data_List_Internal_TwoLeft[gopurs_runtime.Value])((v_4).V0.UnsafePtr).V1}))})
							goto end_branch_6
						} else {

						}
					}
					{
						var __t_tag_2 gopurs_runtime.Value = (v_4).V0
						_ = __t_tag_2
						if __t_tag_2.Type == 9 && __t_tag_2.IntVal == 2884341868 {
							__t6 = Call_Data_List_Internal_fromZipper((v_4).V1, gopurs_runtime.Value{Type: 9, IntVal: 1064476974, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Internal_Three[gopurs_runtime.Value]{1, (*Constructor_Data_List_Internal_TwoRight[gopurs_runtime.Value])((v_4).V0.UnsafePtr).V0, (*Constructor_Data_List_Internal_TwoRight[gopurs_runtime.Value])((v_4).V0.UnsafePtr).V1, (v1_5).V0, (v1_5).V1, (v1_5).V2}))})
							goto end_branch_6
						} else {

						}
					}
					{
						var __t_tag_3 gopurs_runtime.Value = (v_4).V0
						_ = __t_tag_3
						if __t_tag_3.Type == 9 && __t_tag_3.IntVal == 2195694037 {
							v_4_loop = (v_4).V1
							v1_5_loop = (&Constructor_Data_List_Internal_KickUp[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1177901036, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Internal_Two[gopurs_runtime.Value]{1, (v1_5).V0, (v1_5).V1, (v1_5).V2}))}, (*Constructor_Data_List_Internal_ThreeLeft[gopurs_runtime.Value])((v_4).V0.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1177901036, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Internal_Two[gopurs_runtime.Value]{1, (*Constructor_Data_List_Internal_ThreeLeft[gopurs_runtime.Value])((v_4).V0.UnsafePtr).V1, (*Constructor_Data_List_Internal_ThreeLeft[gopurs_runtime.Value])((v_4).V0.UnsafePtr).V2, (*Constructor_Data_List_Internal_ThreeLeft[gopurs_runtime.Value])((v_4).V0.UnsafePtr).V3}))}})
							continue up_3_0_0
							__t6 = func() gopurs_runtime.Value { panic("unreachable") }()
							goto end_branch_6
						} else {

						}
					}
					{
						var __t_tag_4 gopurs_runtime.Value = (v_4).V0
						_ = __t_tag_4
						if __t_tag_4.Type == 9 && __t_tag_4.IntVal == 1584522659 {
							v_4_loop = (v_4).V1
							v1_5_loop = (&Constructor_Data_List_Internal_KickUp[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1177901036, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Internal_Two[gopurs_runtime.Value]{1, (*Constructor_Data_List_Internal_ThreeMiddle[gopurs_runtime.Value])((v_4).V0.UnsafePtr).V0, (*Constructor_Data_List_Internal_ThreeMiddle[gopurs_runtime.Value])((v_4).V0.UnsafePtr).V1, (v1_5).V0}))}, (v1_5).V1, gopurs_runtime.Value{Type: 9, IntVal: 1177901036, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Internal_Two[gopurs_runtime.Value]{1, (v1_5).V2, (*Constructor_Data_List_Internal_ThreeMiddle[gopurs_runtime.Value])((v_4).V0.UnsafePtr).V2, (*Constructor_Data_List_Internal_ThreeMiddle[gopurs_runtime.Value])((v_4).V0.UnsafePtr).V3}))}})
							continue up_3_0_0
							__t6 = func() gopurs_runtime.Value { panic("unreachable") }()
							goto end_branch_6
						} else {

						}
					}
					{
						var __t_tag_5 gopurs_runtime.Value = (v_4).V0
						_ = __t_tag_5
						if __t_tag_5.Type == 9 && __t_tag_5.IntVal == 3952671150 {
							v_4_loop = (v_4).V1
							v1_5_loop = (&Constructor_Data_List_Internal_KickUp[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1177901036, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Internal_Two[gopurs_runtime.Value]{1, (*Constructor_Data_List_Internal_ThreeRight[gopurs_runtime.Value])((v_4).V0.UnsafePtr).V0, (*Constructor_Data_List_Internal_ThreeRight[gopurs_runtime.Value])((v_4).V0.UnsafePtr).V1, (*Constructor_Data_List_Internal_ThreeRight[gopurs_runtime.Value])((v_4).V0.UnsafePtr).V2}))}, (*Constructor_Data_List_Internal_ThreeRight[gopurs_runtime.Value])((v_4).V0.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 1177901036, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Internal_Two[gopurs_runtime.Value]{1, (v1_5).V0, (v1_5).V1, (v1_5).V2}))}})
							continue up_3_0_0
							__t6 = func() gopurs_runtime.Value { panic("unreachable") }()
							goto end_branch_6
						} else {

						}
					}
					{
						__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
					}
				end_branch_6:
					__t7 = __t6
					goto end_branch_7
				} else {

				}
			}
			{
				__t7 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_7:
			return __t7
		}
	}
	up_3_0_0 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_local_Data_List_Internal_up_3_0_0(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_4_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Internal_KickUp[gopurs_runtime.Value]](v1_5_loop_val))
		})
	})
	var Call_local_Data_List_Internal_down_4_8_1 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value) struct {
		found  bool
		result gopurs_runtime.Value
		/*gopurs:record-v1:2:0066006f0075006e0064,0072006500730075006c0074*/
	}
	_ = Call_local_Data_List_Internal_down_4_8_1
	var down_4_8_1 gopurs_runtime.Value
	_ = down_4_8_1
	Call_local_Data_List_Internal_down_4_8_1 = func(v_5_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_6_loop gopurs_runtime.Value) struct {
		found  bool
		result gopurs_runtime.Value
		/*gopurs:record-v1:2:0066006f0075006e0064,0072006500730075006c0074*/
	} {
	down_4_8_1:
		for {
			if false {
				continue down_4_8_1
			}
			var v_5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_5_loop
			_ = v_5
			var v1_6 gopurs_runtime.Value = v1_6_loop
			_ = v1_6
			var __t15 struct {
				found  bool
				result gopurs_runtime.Value
				/*gopurs:record-v1:2:0066006f0075006e0064,0072006500730075006c0074*/
			}
			{
				if v1_6.Type == 9 && v1_6.IntVal == 2764020654 {
					__t15 = struct {
						found  bool
						result gopurs_runtime.Value
						/*gopurs:record-v1:2:0066006f0075006e0064,0072006500730075006c0074*/
					}{false, Call_local_Data_List_Internal_up_3_0_0(v_5, (&Constructor_Data_List_Internal_KickUp[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2764020654, UnsafePtr: unsafe.Pointer(nil)}, k_1, gopurs_runtime.Value{Type: 9, IntVal: 2764020654, UnsafePtr: unsafe.Pointer(nil)}}))}
					goto end_branch_15
				} else {

				}
			}
			{
				if v1_6.Type == 9 && v1_6.IntVal == 1177901036 {
					// TAST (Let): v2_7_9 shape=App(Other) bindingType=(ADT ["Data","Ordering","Ordering"] [])
					v2_7_9 := uint32(gopurs_runtime.Apply2(comp_0, k_1, (*Constructor_Data_List_Internal_Two[gopurs_runtime.Value])(v1_6.UnsafePtr).V1).IntVal)
					_ = v2_7_9
					var __t10 struct {
						found  bool
						result gopurs_runtime.Value
						/*gopurs:record-v1:2:0066006f0075006e0064,0072006500730075006c0074*/
					}
					{
						if v2_7_9 == 902936544 {
							__t10 = struct {
								found  bool
								result gopurs_runtime.Value
								/*gopurs:record-v1:2:0066006f0075006e0064,0072006500730075006c0074*/
							}{true, orig_2}
							goto end_branch_10
						} else {

						}
					}
					{
						if v2_7_9 == 1527465420 {
							v_5_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1304506903, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Internal_TwoLeft[gopurs_runtime.Value]{1, (*Constructor_Data_List_Internal_Two[gopurs_runtime.Value])(v1_6.UnsafePtr).V1, (*Constructor_Data_List_Internal_Two[gopurs_runtime.Value])(v1_6.UnsafePtr).V2}))}, v_5})
							v1_6_loop = (*Constructor_Data_List_Internal_Two[gopurs_runtime.Value])(v1_6.UnsafePtr).V0
							continue down_4_8_1
							__t10 = func() struct {
								found  bool
								result gopurs_runtime.Value
								/*gopurs:record-v1:2:0066006f0075006e0064,0072006500730075006c0074*/
							} {
								panic("unreachable")
							}()
							goto end_branch_10
						} else {

						}
					}
					{
						v_5_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2884341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Internal_TwoRight[gopurs_runtime.Value]{1, (*Constructor_Data_List_Internal_Two[gopurs_runtime.Value])(v1_6.UnsafePtr).V0, (*Constructor_Data_List_Internal_Two[gopurs_runtime.Value])(v1_6.UnsafePtr).V1}))}, v_5})
						v1_6_loop = (*Constructor_Data_List_Internal_Two[gopurs_runtime.Value])(v1_6.UnsafePtr).V2
						continue down_4_8_1
						__t10 = func() struct {
							found  bool
							result gopurs_runtime.Value
							/*gopurs:record-v1:2:0066006f0075006e0064,0072006500730075006c0074*/
						} {
							panic("unreachable")
						}()
					}
				end_branch_10:
					__t15 = __t10
					goto end_branch_15
				} else {

				}
			}
			{
				if v1_6.Type == 9 && v1_6.IntVal == 1064476974 {
					// TAST (Let): v2_7_11 shape=App(Other) bindingType=(ADT ["Data","Ordering","Ordering"] [])
					v2_7_11 := uint32(gopurs_runtime.Apply2(comp_0, k_1, (*Constructor_Data_List_Internal_Three[gopurs_runtime.Value])(v1_6.UnsafePtr).V1).IntVal)
					_ = v2_7_11
					var __t14 struct {
						found  bool
						result gopurs_runtime.Value
						/*gopurs:record-v1:2:0066006f0075006e0064,0072006500730075006c0074*/
					}
					{
						if v2_7_11 == 902936544 {
							__t14 = struct {
								found  bool
								result gopurs_runtime.Value
								/*gopurs:record-v1:2:0066006f0075006e0064,0072006500730075006c0074*/
							}{true, orig_2}
							goto end_branch_14
						} else {

						}
					}
					{
						// TAST (Let): v3_8_12 shape=App(Other) bindingType=(ADT ["Data","Ordering","Ordering"] [])
						v3_8_12 := uint32(gopurs_runtime.Apply2(comp_0, k_1, (*Constructor_Data_List_Internal_Three[gopurs_runtime.Value])(v1_6.UnsafePtr).V3).IntVal)
						_ = v3_8_12
						var __t13 struct {
							found  bool
							result gopurs_runtime.Value
							/*gopurs:record-v1:2:0066006f0075006e0064,0072006500730075006c0074*/
						}
						{
							if v3_8_12 == 902936544 {
								__t13 = struct {
									found  bool
									result gopurs_runtime.Value
									/*gopurs:record-v1:2:0066006f0075006e0064,0072006500730075006c0074*/
								}{true, orig_2}
								goto end_branch_13
							} else {

							}
						}
						{
							if v2_7_11 == 1527465420 {
								v_5_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2195694037, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Internal_ThreeLeft[gopurs_runtime.Value]{1, (*Constructor_Data_List_Internal_Three[gopurs_runtime.Value])(v1_6.UnsafePtr).V1, (*Constructor_Data_List_Internal_Three[gopurs_runtime.Value])(v1_6.UnsafePtr).V2, (*Constructor_Data_List_Internal_Three[gopurs_runtime.Value])(v1_6.UnsafePtr).V3, (*Constructor_Data_List_Internal_Three[gopurs_runtime.Value])(v1_6.UnsafePtr).V4}))}, v_5})
								v1_6_loop = (*Constructor_Data_List_Internal_Three[gopurs_runtime.Value])(v1_6.UnsafePtr).V0
								continue down_4_8_1
								__t13 = func() struct {
									found  bool
									result gopurs_runtime.Value
									/*gopurs:record-v1:2:0066006f0075006e0064,0072006500730075006c0074*/
								} {
									panic("unreachable")
								}()
								goto end_branch_13
							} else {

							}
						}
						{
							if (v2_7_11 == 380165415) && (v3_8_12 == 1527465420) {
								v_5_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1584522659, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Internal_ThreeMiddle[gopurs_runtime.Value]{1, (*Constructor_Data_List_Internal_Three[gopurs_runtime.Value])(v1_6.UnsafePtr).V0, (*Constructor_Data_List_Internal_Three[gopurs_runtime.Value])(v1_6.UnsafePtr).V1, (*Constructor_Data_List_Internal_Three[gopurs_runtime.Value])(v1_6.UnsafePtr).V3, (*Constructor_Data_List_Internal_Three[gopurs_runtime.Value])(v1_6.UnsafePtr).V4}))}, v_5})
								v1_6_loop = (*Constructor_Data_List_Internal_Three[gopurs_runtime.Value])(v1_6.UnsafePtr).V2
								continue down_4_8_1
								__t13 = func() struct {
									found  bool
									result gopurs_runtime.Value
									/*gopurs:record-v1:2:0066006f0075006e0064,0072006500730075006c0074*/
								} {
									panic("unreachable")
								}()
								goto end_branch_13
							} else {

							}
						}
						{
							v_5_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 3952671150, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Internal_ThreeRight[gopurs_runtime.Value]{1, (*Constructor_Data_List_Internal_Three[gopurs_runtime.Value])(v1_6.UnsafePtr).V0, (*Constructor_Data_List_Internal_Three[gopurs_runtime.Value])(v1_6.UnsafePtr).V1, (*Constructor_Data_List_Internal_Three[gopurs_runtime.Value])(v1_6.UnsafePtr).V2, (*Constructor_Data_List_Internal_Three[gopurs_runtime.Value])(v1_6.UnsafePtr).V3}))}, v_5})
							v1_6_loop = (*Constructor_Data_List_Internal_Three[gopurs_runtime.Value])(v1_6.UnsafePtr).V4
							continue down_4_8_1
							__t13 = func() struct {
								found  bool
								result gopurs_runtime.Value
								/*gopurs:record-v1:2:0066006f0075006e0064,0072006500730075006c0074*/
							} {
								panic("unreachable")
							}()
						}
					end_branch_13:
						__t14 = __t13
					}
				end_branch_14:
					__t15 = __t14
					goto end_branch_15
				} else {

				}
			}
			{
				__t15 = func() struct {
					found  bool
					result gopurs_runtime.Value
					/*gopurs:record-v1:2:0066006f0075006e0064,0072006500730075006c0074*/
				} {
					panic("Failed pattern match")
				}()
			}
		end_branch_15:
			return __t15
		}
	}
	down_4_8_1 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				orig := Call_local_Data_List_Internal_down_4_8_1(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_5_loop_val), v1_6_loop_val)
				_ = orig
				return gopurs_runtime.RecordDict2("found", "result", gopurs_runtime.Bool(orig.found), orig.result)
			}()
		})
	})
	return Call_local_Data_List_Internal_down_4_8_1((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil), orig_2)
}
