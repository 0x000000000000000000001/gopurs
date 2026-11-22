package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Ordering_LT gopurs_runtime.Value
var once_Data_Ordering_LT sync.Once

func Get_Data_Ordering_LT() gopurs_runtime.Value {
	once_Data_Ordering_LT.Do(func() {
		cache_Data_Ordering_LT = gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}
	})
	return cache_Data_Ordering_LT
}

var cache_Data_Ordering_GT gopurs_runtime.Value
var once_Data_Ordering_GT sync.Once

func Get_Data_Ordering_GT() gopurs_runtime.Value {
	once_Data_Ordering_GT.Do(func() {
		cache_Data_Ordering_GT = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
	})
	return cache_Data_Ordering_GT
}

var cache_Data_Ordering_EQ gopurs_runtime.Value
var once_Data_Ordering_EQ sync.Once

func Get_Data_Ordering_EQ() gopurs_runtime.Value {
	once_Data_Ordering_EQ.Do(func() {
		cache_Data_Ordering_EQ = gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
	})
	return cache_Data_Ordering_EQ
}

var cache_Data_Ordering_showOrdering gopurs_runtime.Value
var once_Data_Ordering_showOrdering sync.Once

func Get_Data_Ordering_showOrdering() gopurs_runtime.Value {
	once_Data_Ordering_showOrdering.Do(func() {
		cache_Data_Ordering_showOrdering = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Data_Ordering_1209612131_1386611502((&Constructor_Data_Show_Show[uint32]{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t3 gopurs_runtime.Value
			{
				var __t_tag_0 uint32 = uint32(v_0.IntVal)
				_ = __t_tag_0
				if uint32(__t_tag_0) == 1527465420 {
					__t3 = gopurs_runtime.StrWithCodeUnits("LT", 2)
					goto end_branch_3
				} else {

				}
			}
			{
				var __t_tag_1 uint32 = uint32(v_0.IntVal)
				_ = __t_tag_1
				if uint32(__t_tag_1) == 380165415 {
					__t3 = gopurs_runtime.StrWithCodeUnits("GT", 2)
					goto end_branch_3
				} else {

				}
			}
			{
				var __t_tag_2 uint32 = uint32(v_0.IntVal)
				_ = __t_tag_2
				if uint32(__t_tag_2) == 902936544 {
					__t3 = gopurs_runtime.StrWithCodeUnits("EQ", 2)
					goto end_branch_3
				} else {

				}
			}
			{
				__t3 = gopurs_runtime.Str(func() string { panic("Failed pattern match") }())
			}
		end_branch_3:
			return __t3
		})})))}
	})
	return cache_Data_Ordering_showOrdering
}

var cache_Data_Ordering_semigroupOrdering gopurs_runtime.Value
var once_Data_Ordering_semigroupOrdering sync.Once

func Get_Data_Ordering_semigroupOrdering() gopurs_runtime.Value {
	once_Data_Ordering_semigroupOrdering.Do(func() {
		cache_Data_Ordering_semigroupOrdering = gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(Rebox_Data_Ordering_1625289059_4179793454((&Constructor_Data_Semigroup_Semigroup[uint32]{1, gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t3 gopurs_runtime.Value
			{
				var __t_tag_0 uint32 = uint32(v_0.IntVal)
				_ = __t_tag_0
				if uint32(__t_tag_0) == 1527465420 {
					__t3 = gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}
					goto end_branch_3
				} else {

				}
			}
			{
				var __t_tag_1 uint32 = uint32(v_0.IntVal)
				_ = __t_tag_1
				if uint32(__t_tag_1) == 380165415 {
					__t3 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
					goto end_branch_3
				} else {

				}
			}
			{
				var __t_tag_2 uint32 = uint32(v_0.IntVal)
				_ = __t_tag_2
				if uint32(__t_tag_2) == 902936544 {
					__t3 = v1_1
					goto end_branch_3
				} else {

				}
			}
			{
				__t3 = gopurs_runtime.Value{Type: 9, IntVal: int64(func() uint32 { panic("Failed pattern match") }()), UnsafePtr: nil}
			}
		end_branch_3:
			return __t3
		})})))}
	})
	return cache_Data_Ordering_semigroupOrdering
}

var cache_Data_Ordering_invert gopurs_runtime.Value
var once_Data_Ordering_invert sync.Once

func Get_Data_Ordering_invert() gopurs_runtime.Value {
	once_Data_Ordering_invert.Do(func() {
		cache_Data_Ordering_invert = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Data_Ordering_invert(uint32(v_0_box.IntVal))), UnsafePtr: nil}
		})
	})
	return cache_Data_Ordering_invert
}

var cache_Data_Ordering_eqOrdering gopurs_runtime.Value
var once_Data_Ordering_eqOrdering sync.Once

func Get_Data_Ordering_eqOrdering() gopurs_runtime.Value {
	once_Data_Ordering_eqOrdering.Do(func() {
		cache_Data_Ordering_eqOrdering = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Ordering_3768443459_3790796878((&Constructor_Data_Eq_Eq[uint32]{1, gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t7 bool
			{
				var __t_tag_3 uint32 = uint32(v_0.IntVal)
				_ = __t_tag_3
				if uint32(__t_tag_3) == 1527465420 {
					var __t_tag_4 uint32 = uint32(v1_1.IntVal)
					_ = __t_tag_4
					__t7 = (uint32(__t_tag_4) == 1527465420)
					goto end_branch_7
				} else {

				}
			}
			{
				var __t_tag_5 uint32 = uint32(v_0.IntVal)
				_ = __t_tag_5
				if uint32(__t_tag_5) == 380165415 {
					var __t_tag_6 uint32 = uint32(v1_1.IntVal)
					_ = __t_tag_6
					__t7 = (uint32(__t_tag_6) == 380165415)
					goto end_branch_7
				} else {

				}
			}
			{
				var __t_tag_0 uint32 = uint32(v_0.IntVal)
				_ = __t_tag_0
				var __t_and_2 bool = false
				if uint32(__t_tag_0) == 902936544 {

					var __t_tag_1 uint32 = uint32(v1_1.IntVal)
					_ = __t_tag_1
					__t_and_2 = (uint32(__t_tag_1) == 902936544)
				}
				__t7 = __t_and_2
			}
		end_branch_7:
			return gopurs_runtime.Bool(__t7)
		})})))}
	})
	return cache_Data_Ordering_eqOrdering
}

type Constructor_Data_Ordering_LT struct {
	Rc uint32
}

type Constructor_Data_Ordering_GT struct {
	Rc uint32
}

type Constructor_Data_Ordering_EQ struct {
	Rc uint32
}

func Call_Data_Ordering_invert(v_0_loop uint32) uint32 {
	var v_0 uint32 = v_0_loop
	_ = v_0
	var __t0 uint32
	{
		if v_0 == 380165415 {
			__t0 = 1527465420
			goto end_branch_0
		} else {

		}
	}
	{
		if v_0 == 902936544 {
			__t0 = 902936544
			goto end_branch_0
		} else {

		}
	}
	{
		if v_0 == 1527465420 {
			__t0 = 380165415
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = func() uint32 { panic("Failed pattern match") }()
	}
end_branch_0:
	return __t0
}

func Rebox_Data_Ordering_1209612131_1386611502(in *Constructor_Data_Show_Show[uint32]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Ordering_1625289059_4179793454(in *Constructor_Data_Semigroup_Semigroup[uint32]) *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Ordering_3768443459_3790796878(in *Constructor_Data_Eq_Eq[uint32]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[gopurs_runtime.Value])(unsafe.Pointer(in))
}
