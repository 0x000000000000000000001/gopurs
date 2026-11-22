package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Node_Encoding_ASCII gopurs_runtime.Value
var once_Node_Encoding_ASCII sync.Once

func Get_Node_Encoding_ASCII() gopurs_runtime.Value {
	once_Node_Encoding_ASCII.Do(func() {
		cache_Node_Encoding_ASCII = gopurs_runtime.Value{Type: 9, IntVal: int64(2598402200), UnsafePtr: nil}
	})
	return cache_Node_Encoding_ASCII
}

var cache_Node_Encoding_UTF8 gopurs_runtime.Value
var once_Node_Encoding_UTF8 sync.Once

func Get_Node_Encoding_UTF8() gopurs_runtime.Value {
	once_Node_Encoding_UTF8.Do(func() {
		cache_Node_Encoding_UTF8 = gopurs_runtime.Value{Type: 9, IntVal: int64(346813718), UnsafePtr: nil}
	})
	return cache_Node_Encoding_UTF8
}

var cache_Node_Encoding_UTF16LE gopurs_runtime.Value
var once_Node_Encoding_UTF16LE sync.Once

func Get_Node_Encoding_UTF16LE() gopurs_runtime.Value {
	once_Node_Encoding_UTF16LE.Do(func() {
		cache_Node_Encoding_UTF16LE = gopurs_runtime.Value{Type: 9, IntVal: int64(2352678688), UnsafePtr: nil}
	})
	return cache_Node_Encoding_UTF16LE
}

var cache_Node_Encoding_UCS2 gopurs_runtime.Value
var once_Node_Encoding_UCS2 sync.Once

func Get_Node_Encoding_UCS2() gopurs_runtime.Value {
	once_Node_Encoding_UCS2.Do(func() {
		cache_Node_Encoding_UCS2 = gopurs_runtime.Value{Type: 9, IntVal: int64(4287456318), UnsafePtr: nil}
	})
	return cache_Node_Encoding_UCS2
}

var cache_Node_Encoding_Base64 gopurs_runtime.Value
var once_Node_Encoding_Base64 sync.Once

func Get_Node_Encoding_Base64() gopurs_runtime.Value {
	once_Node_Encoding_Base64.Do(func() {
		cache_Node_Encoding_Base64 = gopurs_runtime.Value{Type: 9, IntVal: int64(2405785054), UnsafePtr: nil}
	})
	return cache_Node_Encoding_Base64
}

var cache_Node_Encoding_Base64Url gopurs_runtime.Value
var once_Node_Encoding_Base64Url sync.Once

func Get_Node_Encoding_Base64Url() gopurs_runtime.Value {
	once_Node_Encoding_Base64Url.Do(func() {
		cache_Node_Encoding_Base64Url = gopurs_runtime.Value{Type: 9, IntVal: int64(2476550165), UnsafePtr: nil}
	})
	return cache_Node_Encoding_Base64Url
}

var cache_Node_Encoding_Latin1 gopurs_runtime.Value
var once_Node_Encoding_Latin1 sync.Once

func Get_Node_Encoding_Latin1() gopurs_runtime.Value {
	once_Node_Encoding_Latin1.Do(func() {
		cache_Node_Encoding_Latin1 = gopurs_runtime.Value{Type: 9, IntVal: int64(3576003686), UnsafePtr: nil}
	})
	return cache_Node_Encoding_Latin1
}

var cache_Node_Encoding_Binary gopurs_runtime.Value
var once_Node_Encoding_Binary sync.Once

func Get_Node_Encoding_Binary() gopurs_runtime.Value {
	once_Node_Encoding_Binary.Do(func() {
		cache_Node_Encoding_Binary = gopurs_runtime.Value{Type: 9, IntVal: int64(349879974), UnsafePtr: nil}
	})
	return cache_Node_Encoding_Binary
}

var cache_Node_Encoding_Hex gopurs_runtime.Value
var once_Node_Encoding_Hex sync.Once

func Get_Node_Encoding_Hex() gopurs_runtime.Value {
	once_Node_Encoding_Hex.Do(func() {
		cache_Node_Encoding_Hex = gopurs_runtime.Value{Type: 9, IntVal: int64(106823740), UnsafePtr: nil}
	})
	return cache_Node_Encoding_Hex
}

var cache_Node_Encoding_showEncoding gopurs_runtime.Value
var once_Node_Encoding_showEncoding sync.Once

func Get_Node_Encoding_showEncoding() gopurs_runtime.Value {
	once_Node_Encoding_showEncoding.Do(func() {
		cache_Node_Encoding_showEncoding = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Node_Encoding_1209612131_1386611502((&Constructor_Data_Show_Show[uint32]{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t9 gopurs_runtime.Value
			{
				var __t_tag_0 uint32 = uint32(v_0.IntVal)
				_ = __t_tag_0
				if uint32(__t_tag_0) == 2598402200 {
					__t9 = gopurs_runtime.StrWithCodeUnits("ASCII", 5)
					goto end_branch_9
				} else {

				}
			}
			{
				var __t_tag_1 uint32 = uint32(v_0.IntVal)
				_ = __t_tag_1
				if uint32(__t_tag_1) == 346813718 {
					__t9 = gopurs_runtime.StrWithCodeUnits("UTF8", 4)
					goto end_branch_9
				} else {

				}
			}
			{
				var __t_tag_2 uint32 = uint32(v_0.IntVal)
				_ = __t_tag_2
				if uint32(__t_tag_2) == 2352678688 {
					__t9 = gopurs_runtime.StrWithCodeUnits("UTF16LE", 7)
					goto end_branch_9
				} else {

				}
			}
			{
				var __t_tag_3 uint32 = uint32(v_0.IntVal)
				_ = __t_tag_3
				if uint32(__t_tag_3) == 4287456318 {
					__t9 = gopurs_runtime.StrWithCodeUnits("UCS2", 4)
					goto end_branch_9
				} else {

				}
			}
			{
				var __t_tag_4 uint32 = uint32(v_0.IntVal)
				_ = __t_tag_4
				if uint32(__t_tag_4) == 2405785054 {
					__t9 = gopurs_runtime.StrWithCodeUnits("Base64", 6)
					goto end_branch_9
				} else {

				}
			}
			{
				var __t_tag_5 uint32 = uint32(v_0.IntVal)
				_ = __t_tag_5
				if uint32(__t_tag_5) == 2476550165 {
					__t9 = gopurs_runtime.StrWithCodeUnits("Base64Url", 9)
					goto end_branch_9
				} else {

				}
			}
			{
				var __t_tag_6 uint32 = uint32(v_0.IntVal)
				_ = __t_tag_6
				if uint32(__t_tag_6) == 3576003686 {
					__t9 = gopurs_runtime.StrWithCodeUnits("Latin1", 6)
					goto end_branch_9
				} else {

				}
			}
			{
				var __t_tag_7 uint32 = uint32(v_0.IntVal)
				_ = __t_tag_7
				if uint32(__t_tag_7) == 349879974 {
					__t9 = gopurs_runtime.StrWithCodeUnits("Binary", 6)
					goto end_branch_9
				} else {

				}
			}
			{
				var __t_tag_8 uint32 = uint32(v_0.IntVal)
				_ = __t_tag_8
				if uint32(__t_tag_8) == 106823740 {
					__t9 = gopurs_runtime.StrWithCodeUnits("Hex", 3)
					goto end_branch_9
				} else {

				}
			}
			{
				__t9 = gopurs_runtime.Str(func() string { panic("Failed pattern match") }())
			}
		end_branch_9:
			return __t9
		})})))}
	})
	return cache_Node_Encoding_showEncoding
}

var cache_Node_Encoding_encodingToNode gopurs_runtime.Value
var once_Node_Encoding_encodingToNode sync.Once

func Get_Node_Encoding_encodingToNode() gopurs_runtime.Value {
	once_Node_Encoding_encodingToNode.Do(func() {
		cache_Node_Encoding_encodingToNode = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Node_Encoding_encodingToNode(uint32(v_0_box.IntVal)))
		})
	})
	return cache_Node_Encoding_encodingToNode
}

var cache_Node_Encoding_byteLength gopurs_runtime.Value
var once_Node_Encoding_byteLength sync.Once

func Get_Node_Encoding_byteLength() gopurs_runtime.Value {
	once_Node_Encoding_byteLength.Do(func() {
		cache_Node_Encoding_byteLength = gopurs_runtime.Func2(func(str_0_box gopurs_runtime.Value, enc_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Node_Encoding_byteLength(str_0_box.StrVal(), uint32(enc_1_box.IntVal)))
		})
	})
	return cache_Node_Encoding_byteLength
}

type Constructor_Node_Encoding_ASCII struct {
	Rc uint32
}

type Constructor_Node_Encoding_UTF8 struct {
	Rc uint32
}

type Constructor_Node_Encoding_UTF16LE struct {
	Rc uint32
}

type Constructor_Node_Encoding_UCS2 struct {
	Rc uint32
}

type Constructor_Node_Encoding_Base64 struct {
	Rc uint32
}

type Constructor_Node_Encoding_Base64Url struct {
	Rc uint32
}

type Constructor_Node_Encoding_Latin1 struct {
	Rc uint32
}

type Constructor_Node_Encoding_Binary struct {
	Rc uint32
}

type Constructor_Node_Encoding_Hex struct {
	Rc uint32
}

func Call_Node_Encoding_encodingToNode(v_0_loop uint32) string {
	var v_0 uint32 = v_0_loop
	_ = v_0
	var __t0 string
	{
		if v_0 == 2598402200 {
			__t0 = "ascii"
			goto end_branch_0
		} else {

		}
	}
	{
		if v_0 == 346813718 {
			__t0 = "utf8"
			goto end_branch_0
		} else {

		}
	}
	{
		if v_0 == 2352678688 {
			__t0 = "utf16le"
			goto end_branch_0
		} else {

		}
	}
	{
		if v_0 == 4287456318 {
			__t0 = "ucs2"
			goto end_branch_0
		} else {

		}
	}
	{
		if v_0 == 2405785054 {
			__t0 = "base64"
			goto end_branch_0
		} else {

		}
	}
	{
		if v_0 == 2476550165 {
			__t0 = "base64url"
			goto end_branch_0
		} else {

		}
	}
	{
		if v_0 == 3576003686 {
			__t0 = "latin1"
			goto end_branch_0
		} else {

		}
	}
	{
		if v_0 == 349879974 {
			__t0 = "binary"
			goto end_branch_0
		} else {

		}
	}
	{
		if v_0 == 106823740 {
			__t0 = "hex"
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = func() string { panic("Failed pattern match") }()
	}
end_branch_0:
	return __t0
}

func Call_Node_Encoding_byteLength(str_0_loop string, enc_1_loop uint32) int64 {
	var str_0 string = str_0_loop
	_ = str_0
	var enc_1 uint32 = enc_1_loop
	_ = enc_1
	var __t0 string
	{
		if enc_1 == 2598402200 {
			__t0 = "ascii"
			goto end_branch_0
		} else {

		}
	}
	{
		if enc_1 == 346813718 {
			__t0 = "utf8"
			goto end_branch_0
		} else {

		}
	}
	{
		if enc_1 == 2352678688 {
			__t0 = "utf16le"
			goto end_branch_0
		} else {

		}
	}
	{
		if enc_1 == 4287456318 {
			__t0 = "ucs2"
			goto end_branch_0
		} else {

		}
	}
	{
		if enc_1 == 2405785054 {
			__t0 = "base64"
			goto end_branch_0
		} else {

		}
	}
	{
		if enc_1 == 2476550165 {
			__t0 = "base64url"
			goto end_branch_0
		} else {

		}
	}
	{
		if enc_1 == 3576003686 {
			__t0 = "latin1"
			goto end_branch_0
		} else {

		}
	}
	{
		if enc_1 == 349879974 {
			__t0 = "binary"
			goto end_branch_0
		} else {

		}
	}
	{
		if enc_1 == 106823740 {
			__t0 = "hex"
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = func() string { panic("Failed pattern match") }()
	}
end_branch_0:
	return gopurs_runtime.UncurriedApp2(Get_Node_Encoding_byteLengthImpl(), gopurs_runtime.Str(str_0), gopurs_runtime.Str(__t0)).IntVal
}

func Rebox_Node_Encoding_1209612131_1386611502(in *Constructor_Data_Show_Show[uint32]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Get_Node_Encoding_byteLengthImpl() gopurs_runtime.Value {
	return _Gopurs_Node_Encoding_ByteLengthImpl
}
