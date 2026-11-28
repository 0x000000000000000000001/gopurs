package gopurs_runtime

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"math"
	"reflect"
	"sort"
	"strings"
	"sync"
	"sync/atomic"
	"unicode/utf8"
	"unsafe"
)

var EventLoop sync.WaitGroup

// Keep the published *any reference envelope. A private payload lets new
// writes replace Value fields without allocating another boxed interface.
// Legacy writes may replace the payload; the next new write promotes it again.
type stValueCell struct{ value Value }

func NewSTRef(value Value) Value {
	var cell any = &stValueCell{value}
	return Any(&cell)
}

func ReadSTRef(ref Value) Value {
	return AnyToValue(*ref.PtrVal().(*any))
}

func WriteSTRef(ref, value Value) Value {
	slot := ref.PtrVal().(*any)
	if cell, ok := (*slot).(*stValueCell); ok {
		cell.value = value
	} else {
		*slot = &stValueCell{value}
	}
	return value
}

func Retain() {
	EventLoop.Add(1)
}

func Release() {
	EventLoop.Done()
}

func EventLoopWait() {
	EventLoop.Wait()
}

// FFIContext lets generated wrappers name the standard context type without
// introducing an import that could collide with imports in handwritten Go.
type FFIContext = context.Context

// EffectFnAff is the compiler's adapter for ordinary Go operations. Constructing
// the value does not start work. Each installation starts a fresh operation and
// returns its canceler without waiting for a blocking call to finish.
//
// Cancellation stops delivery and cancels the operation's context. An operation
// that ignores its context may continue running. It owns any resources it does
// not deliver; the adapter never guesses a destructor for a late result. Use an
// acquisition/release bracket when ownership must survive cancellation.
//
// Only error results enter the failure callback. A typed nil error retains Go's
// non-nil interface semantics. Panics are deliberately not recovered here.
func EffectFnAff(run func(context.Context) (Value, error)) Value {
	return EffectFnAffAcquire(func(ctx context.Context) (Value, func(), error) {
		value, err := run(ctx)
		return value, nil, err
	})
}

// EffectFnAffAcquire adds an explicit release function for acquired results.
// A non-nil release is called once on error or when cancellation wins before
// delivery. Successful delivery transfers ownership to the caller and does not
// call release; an application's normal bracket/close remains responsible.
// Release runs in the worker after the operation returns, even when the caller
// has already stopped waiting. Its panics retain ordinary Go panic semantics.
func EffectFnAffAcquire(run func(context.Context) (Value, func(), error)) Value {
	return Func2(func(onError, onSuccess Value) Value {
		operation := newFFIAffOperation(onError, onSuccess)
		go operation.runAcquire(run)
		return Func3(func(reason, _ Value, onCanceled Value) Value {
			delivered := operation.stop(reason)
			if delivered == nil {
				Apply(onCanceled, Value{})
			} else {
				select {
				case <-delivered:
					Apply(onCanceled, Value{})
				default:
					// A callback may invoke its own canceler. Return from the
					// call immediately, but acknowledge only after publication.
					go func() { <-delivered; Apply(onCanceled, Value{}) }()
				}
			}
			return Value{}
		})
	})
}

// pending -> completed or canceled is the single arbitration point. Callback
// delivery belongs to the winning completion, including when cancellation is
// requested while that callback is already executing.
type ffiAffOperation struct {
	state     atomic.Uint32
	ctx       context.Context
	cancel    context.CancelCauseFunc
	delivered chan struct{}
	onError   Value
	onSuccess Value
}

func newFFIAffOperation(onError, onSuccess Value) *ffiAffOperation {
	ctx, cancel := context.WithCancelCause(context.Background())
	return &ffiAffOperation{ctx: ctx, cancel: cancel, delivered: make(chan struct{}), onError: onError, onSuccess: onSuccess}
}

func (operation *ffiAffOperation) run(run func(context.Context) (Value, error)) {
	operation.runAcquire(func(ctx context.Context) (Value, func(), error) {
		value, err := run(ctx)
		return value, nil, err
	})
}

func (operation *ffiAffOperation) runAcquire(run func(context.Context) (Value, func(), error)) {
	defer func() {
		operation.cancel(nil)
		close(operation.delivered)
	}()
	if operation.state.Load() != 0 {
		return
	}
	value, release, err := run(operation.ctx)
	if err != nil && release != nil {
		release()
		release = nil
	}
	if !operation.state.CompareAndSwap(0, 1) {
		if release != nil {
			release()
		}
		return
	}
	if err != nil {
		Apply(operation.onError, Box(err))
	} else {
		Apply(operation.onSuccess, value)
	}
}

func (operation *ffiAffOperation) stop(reason Value) <-chan struct{} {
	if operation.state.CompareAndSwap(0, 2) {
		var cause error
		if reason.Type != 0 {
			cause = Unbox[error](reason)
		}
		operation.cancel(cause)
	} else if operation.state.Load() == 1 {
		// Completion already owns the result. Acknowledge cancellation only
		// after its callback has published it, so Aff cannot abandon ownership
		// between this adapter's arbitration and its own callback boundary.
		return operation.delivered
	}
	return nil
}

var StructGetters = map[int64]func(unsafe.Pointer, string) Value{}

const (
	TypeFunc         = 1
	TypeFunc2        = 2
	TypeFunc3        = 3
	TypeFunc4        = 4
	TypeFunc5        = 5
	TypeInt          = 6
	TypeString       = 7
	TypeRecord       = 8
	TypeConstructor  = 9
	TypeFloat        = 10
	TypeBool         = 11
	TypeArray        = 12
	TypeAny          = 13
	TypeFunc6        = 14
	TypeFunc7        = 15
	TypeFunc8        = 16
	TypeFunc9        = 17
	TypeFunc10       = 18
	TypeRecord0      = 19
	TypeRecord1      = 20
	TypeRecord2      = 21
	TypeRecord3      = 22
	TypeRecord4      = 23
	TypeRecord5      = 24
	TypeRecordData   = 25
	TypeFunc11       = 26
	TypeFunctionData = 27
	TypeIntArray     = 28
	TypeInt32Array   = 29
	TypeStringArray  = 30
)

// We do not add FloatVal or BoolVal fields to keep the struct size minimal.
// Floats are packed into IntVal using math.Float64bits, and Bools are mapped to 1/0 in IntVal.
// Adding more fields would increase the struct size and reduce pass-by-value performance.
type Value struct {
	Type      int
	IntVal    int64
	UnsafePtr unsafe.Pointer
}

// IncRef increments the Rc field if the value is a pointer ADT.
// It assumes that all pointer ADTs have `Rc uint32` as their very first field.
func IncRef(v Value) Value {
	if v.Type == TypeConstructor {
		ptr := (*struct{ Rc uint32 })(v.UnsafePtr)
		if ptr != nil {
			ptr.Rc++
		}
	}
	return v
}

func (v Value) FloatVal() float64 {
	return math.Float64frombits(uint64(v.IntVal))
}

// StrValue reads a TypeString payload.
//
// Strings are packed into the Value: a nonnegative IntVal holds the byte
// length. A negative IntVal encodes optional immutable UTF-16 and byte lengths
// as -1 - ((byteLength << 31) | codeUnits). UnsafePtr holds the immutable
// string data (unsafe.StringData). The referenced bytes must never be mutated
// for the whole lifetime of the Value, as
// required by unsafe.String. A substring keeps its whole backing alive.
func StrValue(v Value) string {
	if v.UnsafePtr == nil {
		return ""
	}
	if v.IntVal >= 0 {
		return unsafe.String((*byte)(v.UnsafePtr), int(v.IntVal))
	}
	return unsafe.String((*byte)(v.UnsafePtr), int(uint64(^v.IntVal)>>31))
}

func Str(v string) Value {
	if len(v) == 0 {
		return Value{Type: TypeString}
	}
	return Value{Type: TypeString, IntVal: int64(len(v)), UnsafePtr: unsafe.Pointer(unsafe.StringData(v))}
}

func (v Value) StrVal() string {
	return StrValue(v)
}

const stringCodeUnitsMask int64 = 1<<31 - 1

// packStringCodeUnits accepts only representable lengths. Its separate length
// arguments let boundary tests cover large strings without allocating them.
func packStringCodeUnits(bytes, units int) (int64, bool) {
	if bytes < 0 || int64(bytes) > stringCodeUnitsMask || units < 0 || units > bytes {
		return 0, false
	}
	return -1 - ((int64(bytes) << 31) | int64(units)), true
}

// StrWithCodeUnits attaches an already computed UTF-16 length to an immutable
// string. Callers must supply the exact CodeUnits length of these bytes,
// including WTF-8 surrogates and one unit per malformed UTF-8 byte. It never
// computes the length itself. Unrepresentable lengths keep the ordinary string
// representation and use the decoder when length is requested.
func StrWithCodeUnits(s string, units int) Value {
	value := Str(s)
	if len(s) == 0 {
		return value
	}
	if packed, ok := packStringCodeUnits(len(s), units); ok {
		value.IntVal = packed
	}
	return value
}

// NormalizeStringValue replaces an identity String unbox/rebox without losing
// optional immutable length metadata. Other values retain the previous path.
func NormalizeStringValue(value Value) Value {
	if value.Type == TypeString {
		return value
	}
	return Str(value.StrVal())
}

// CachedStringCodeUnits returns a known UTF-16 length, if present. The ordinary
// empty-string representation also has a known zero length without metadata.
// Unknown strings are left to the caller's normal decoder.
func CachedStringCodeUnits(value Value) (int, bool) {
	if value.Type != TypeString {
		return 0, false
	}
	if value.IntVal == 0 {
		return 0, true
	}
	if value.IntVal < 0 {
		return int((-1 - value.IntVal) & stringCodeUnitsMask), true
	}
	return 0, false
}

// StringCodeUnitsLength returns the UTF-16 length, using metadata when a
// producer already knows it. Ordinary strings use the same scan as the FFI.
func StringCodeUnitsLength(value Value) int {
	if value.Type == TypeString && value.IntVal < 0 {
		return int(^value.IntVal & stringCodeUnitsMask)
	}
	return int(StringCodeUnitsLengthSlow(value))
}

// StringCodeUnitsLengthSlow handles empty, uncached and non-string values
// out-of-line so the metadata fast path inlines into hot callers.
//
//go:noinline
func StringCodeUnitsLengthSlow(value Value) int64 {
	if value.Type == TypeString && value.IntVal == 0 {
		return 0
	}
	return int64(CountStringCodeUnits(value.StrVal()))
}

// StringCodeUnitsLengthsEqual compares UTF-16 lengths without scanning identical
// immutable string views. Other views retain their optional length metadata.
func StringCodeUnitsLengthsEqual(left, right Value) bool {
	l, r := left.StrVal(), right.StrVal()
	if len(l) == len(r) && unsafe.StringData(l) == unsafe.StringData(r) {
		return true
	}
	return StringCodeUnitsLength(left) == StringCodeUnitsLength(right)
}

// Zero entries are invalid leading bytes. Surrogate encodings use width 3,
// because native strings also preserve isolated UTF-16 units as WTF-8.
var stringUnitSequenceWidth = [256]byte{
	0xc2: 2, 0xc3: 2, 0xc4: 2, 0xc5: 2, 0xc6: 2, 0xc7: 2,
	0xc8: 2, 0xc9: 2, 0xca: 2, 0xcb: 2, 0xcc: 2, 0xcd: 2, 0xce: 2, 0xcf: 2,
	0xd0: 2, 0xd1: 2, 0xd2: 2, 0xd3: 2, 0xd4: 2, 0xd5: 2, 0xd6: 2, 0xd7: 2,
	0xd8: 2, 0xd9: 2, 0xda: 2, 0xdb: 2, 0xdc: 2, 0xdd: 2, 0xde: 2, 0xdf: 2,
	0xe0: 3, 0xe1: 3, 0xe2: 3, 0xe3: 3, 0xe4: 3, 0xe5: 3, 0xe6: 3, 0xe7: 3,
	0xe8: 3, 0xe9: 3, 0xea: 3, 0xeb: 3, 0xec: 3, 0xed: 3, 0xee: 3, 0xef: 3,
	0xf0: 4, 0xf1: 4, 0xf2: 4, 0xf3: 4, 0xf4: 4,
}

func stringASCIIUnitPrefix(s string) int {
	i := 0
	// Keep short ASCII runs on the small scan. Reading a wide block before
	// proving a long run adds work to strings alternating ASCII and Unicode.
	for i < 128 && i+32 <= len(s) {
		b := []byte(s[i : i+32])
		word := binary.LittleEndian.Uint64(b) | binary.LittleEndian.Uint64(b[8:]) |
			binary.LittleEndian.Uint64(b[16:]) | binary.LittleEndian.Uint64(b[24:])
		if word&0x8080808080808080 != 0 {
			goto tail
		}
		i += 32
	}
	for i+128 <= len(s) {
		b := []byte(s[i : i+128])
		first := binary.LittleEndian.Uint64(b) | binary.LittleEndian.Uint64(b[8:]) |
			binary.LittleEndian.Uint64(b[16:]) | binary.LittleEndian.Uint64(b[24:])
		second := binary.LittleEndian.Uint64(b[32:]) | binary.LittleEndian.Uint64(b[40:]) |
			binary.LittleEndian.Uint64(b[48:]) | binary.LittleEndian.Uint64(b[56:])
		third := binary.LittleEndian.Uint64(b[64:]) | binary.LittleEndian.Uint64(b[72:]) |
			binary.LittleEndian.Uint64(b[80:]) | binary.LittleEndian.Uint64(b[88:])
		fourth := binary.LittleEndian.Uint64(b[96:]) | binary.LittleEndian.Uint64(b[104:]) |
			binary.LittleEndian.Uint64(b[112:]) | binary.LittleEndian.Uint64(b[120:])
		if (first|second|third|fourth)&0x8080808080808080 != 0 {
			break
		}
		i += 128
	}
	for i+32 <= len(s) {
		b := []byte(s[i : i+32])
		word := binary.LittleEndian.Uint64(b) | binary.LittleEndian.Uint64(b[8:]) |
			binary.LittleEndian.Uint64(b[16:]) | binary.LittleEndian.Uint64(b[24:])
		if word&0x8080808080808080 != 0 {
			break
		}
		i += 32
	}
tail:
	for i+8 <= len(s) {
		if binary.LittleEndian.Uint64([]byte(s[i:i+8]))&0x8080808080808080 != 0 {
			break
		}
		i += 8
	}
	for i < len(s) && s[i] < utf8.RuneSelf {
		i++
	}
	return i
}

// DecodeStringCodePoint reads one PureScript code point from native UTF-8/WTF-8.
// Adjacent high/low surrogate encodings represent one supplementary code point;
// isolated surrogates retain their value. Other malformed bytes consume one
// byte and yield RuneError, matching Data.String.CodePoints.toCodePointArray.
// The empty input follows utf8.DecodeRuneInString: RuneError with width zero.
func DecodeStringCodePoint(s string) (rune, int) {
	if len(s) == 0 {
		return utf8.RuneError, 0
	}
	b0 := s[0]
	if b0 < utf8.RuneSelf {
		return rune(b0), 1
	}
	if b0 < 0xC2 {
		return utf8.RuneError, 1
	}
	if b0 < 0xE0 {
		if len(s) >= 2 && s[1]&0xC0 == 0x80 {
			return rune(b0&0x1F)<<6 | rune(s[1]&0x3F), 2
		}
		return utf8.RuneError, 1
	}
	if b0 < 0xF0 {
		if len(s) >= 3 && s[1]&0xC0 == 0x80 && s[2]&0xC0 == 0x80 && (b0 != 0xE0 || s[1] >= 0xA0) {
			r := rune(b0&0x0F)<<12 | rune(s[1]&0x3F)<<6 | rune(s[2]&0x3F)
			if b0 == 0xED && s[1] >= 0xA0 && s[1] <= 0xAF &&
				len(s) >= 6 && s[3] == 0xED && s[4] >= 0xB0 && s[4] <= 0xBF && s[5]&0xC0 == 0x80 {
				low := rune(s[3]&0x0F)<<12 | rune(s[4]&0x3F)<<6 | rune(s[5]&0x3F)
				return 0x10000 + (r-0xD800)*0x400 + low - 0xDC00, 6
			}
			return r, 3
		}
		return utf8.RuneError, 1
	}
	if b0 <= 0xF4 {
		if len(s) >= 4 && s[1]&0xC0 == 0x80 && s[2]&0xC0 == 0x80 && s[3]&0xC0 == 0x80 &&
			(b0 != 0xF0 || s[1] >= 0x90) && (b0 != 0xF4 || s[1] <= 0x8F) {
			return rune(b0&0x07)<<18 | rune(s[1]&0x3F)<<12 | rune(s[2]&0x3F)<<6 | rune(s[3]&0x3F), 4
		}
	}
	return utf8.RuneError, 1
}

// CountStringCodeUnits scans a native string using PureScript UTF-16 semantics.
func CountStringCodeUnits(s string) int {
	// Count ASCII runs a word at a time. The width table avoids trying every
	// UTF-8 sequence width for each non-ASCII character. Invalid sequences
	// still consume one byte, matching the UTF-8/WTF-8 decoder.
	units := 0
	for len(s) > 0 {
		b := s[0]
		if b < utf8.RuneSelf {
			s = s[1:]
			units++
			if len(s) == 0 {
				break
			}
			b = s[0]
			if b < utf8.RuneSelf {
				n := stringASCIIUnitPrefix(s)
				units += n
				s = s[n:]
				if len(s) == 0 {
					break
				}
				b = s[0]
			}
		}
		units++
		switch stringUnitSequenceWidth[b] {
		case 2:
			if len(s) >= 2 && s[1]&0xc0 == 0x80 {
				s = s[2:]
				continue
			}
		case 3:
			if len(s) >= 3 && s[1]&0xc0 == 0x80 && s[2]&0xc0 == 0x80 && (b != 0xe0 || s[1] >= 0xa0) {
				s = s[3:]
				continue
			}
		case 4:
			if len(s) >= 4 && s[1]&0xc0 == 0x80 && s[2]&0xc0 == 0x80 && s[3]&0xc0 == 0x80 && (b != 0xf0 || s[1] >= 0x90) && (b != 0xf4 || s[1] <= 0x8f) {
				units++
				s = s[4:]
				continue
			}
		}
		s = s[1:]
	}
	return units
}

// ConcatString concatenates PureScript UTF-16 strings in their canonical WTF-8
// representation. Each operand already combines its valid surrogate pairs;
// only a trailing high surrogate and a leading low surrogate can form a new
// pair at the join. Preserve all other bytes, including isolated surrogates.
func ConcatString(left, right string) string {
	n := len(left)
	if n >= 3 && len(right) >= 3 &&
		left[n-3] == 0xed && left[n-2]&0xf0 == 0xa0 && left[n-1]&0xc0 == 0x80 &&
		right[0] == 0xed && right[1]&0xf0 == 0xb0 && right[2]&0xc0 == 0x80 {
		high := rune(left[n-2]&0x0f)<<6 | rune(left[n-1]&0x3f)
		low := rune(right[1]&0x0f)<<6 | rune(right[2]&0x3f)
		return left[:n-3] + string(0x10000+(high<<10)+low) + right[3:]
	}
	return left + right
}

// ConcatStrings implements a strict fold of the canonical string append. The
// projection is compiler-generated (identity or Value.StrVal), never a user
// callback. Holding a trailing high surrogate preserves ConcatString's boundary
// normalization, including when empty fragments separate a surrogate pair.
func ConcatStrings[T any](initial string, parts []T, text func(T) string) string {
	if len(parts) == 0 {
		return initial
	}
	total := len(initial)
	for _, part := range parts {
		total += len(text(part))
	}
	var b strings.Builder
	b.Grow(total)
	// Keep the last three bytes, including partial/malformed UTF-8. An append
	// may complete a high surrogate from bytes split across previous fragments.
	var tail [3]byte
	n := 0
	appendRaw := func(s string) {
		if n+len(s) <= len(tail) {
			copy(tail[n:], s)
			n += len(s)
			return
		}
		flush := n + len(s) - len(tail)
		if flush < n {
			b.Write(tail[:flush])
			copy(tail[:], tail[flush:n])
			copy(tail[n-flush:], s)
		} else {
			b.Write(tail[:n])
			cut := flush - n
			b.WriteString(s[:cut])
			copy(tail[:], s[cut:])
		}
		n = len(tail)
	}
	appendRaw(initial)
	for _, part := range parts {
		s := text(part)
		if n == 3 && len(s) >= 3 &&
			tail[0] == 0xed && tail[1]&0xf0 == 0xa0 && tail[2]&0xc0 == 0x80 &&
			s[0] == 0xed && s[1]&0xf0 == 0xb0 && s[2]&0xc0 == 0x80 {
			high := rune(tail[1]&0x0f)<<6 | rune(tail[2]&0x3f)
			low := rune(s[1]&0x0f)<<6 | rune(s[2]&0x3f)
			n = 0
			appendRaw(string(0x10000 + (high << 10) + low))
			s = s[3:]
		}
		appendRaw(s)
	}
	b.Write(tail[:n])
	return b.String()
}

func Int(v int64) Value {
	return Value{Type: TypeInt, IntVal: v}
}

// Int storage also carries the unsigned result of >>> and the boundary
// quotient bottom / -1. Coerce at the same operations as the JS Prelude,
// rather than narrowing every stored value.
func IntNegate(a int64) int64 { return int64(-int32(a)) }
func IntAdd(a, b int64) int64 { return int64(int32(a + b)) }
func IntSub(a, b int64) int64 { return int64(int32(a - b)) }

func IntMul(a, b int64) int64 {
	// The reference Prelude uses (a * b) | 0: Number multiplication rounds
	// before ToInt32. An exact integer product (or Math.imul) differs near
	// the bounds. Reduce before converting, since a product can exceed int64.
	return int64(int32(int64(math.Mod(float64(a)*float64(b), 4294967296))))
}

func IntBitNot(a int64) int64    { return int64(^int32(a)) }
func IntBitAnd(a, b int64) int64 { return int64(int32(a) & int32(b)) }
func IntBitOr(a, b int64) int64  { return int64(int32(a) | int32(b)) }
func IntBitXor(a, b int64) int64 { return int64(int32(a) ^ int32(b)) }
func IntShl(a, b int64) int64    { return int64(int32(a) << (uint32(b) & 31)) }
func IntShr(a, b int64) int64    { return int64(int32(a) >> (uint32(b) & 31)) }
func IntZshr(a, b int64) int64   { return int64(uint32(a) >> (uint32(b) & 31)) }

// PureScript uses Euclidean division: a non-negative remainder, and zero
// for both division and modulo by zero. Go instead truncates toward zero.
func IntDiv(a, b int64) int64 {
	if b == 0 {
		return 0
	}
	q := a / b
	if a%b < 0 {
		if b > 0 {
			q--
		} else {
			q++
		}
	}
	return q
}

func IntMod(a, b int64) int64 {
	if b == 0 {
		return 0
	}
	r := a % b
	if r < 0 {
		if b > 0 {
			r += b
		} else {
			r -= b
		}
	}
	return r
}

func Float(v float64) Value {
	return Value{Type: TypeFloat, IntVal: int64(math.Float64bits(v))}
}

func NegativeZero() float64 {
	return math.Copysign(0, -1)
}

func Bool(v bool) Value {
	var i int64 = 0
	if v {
		i = 1
	}
	return Value{Type: TypeBool, IntVal: i}
}

func FloatAdd(a, b Value) Value {
	return Float(math.Float64frombits(uint64(a.IntVal)) + math.Float64frombits(uint64(b.IntVal)))
}
func FloatSub(a, b Value) Value {
	return Float(math.Float64frombits(uint64(a.IntVal)) - math.Float64frombits(uint64(b.IntVal)))
}
func FloatMul(a, b Value) Value {
	return Float(math.Float64frombits(uint64(a.IntVal)) * math.Float64frombits(uint64(b.IntVal)))
}
func FloatDiv(a, b Value) Value {
	return Float(math.Float64frombits(uint64(a.IntVal)) / math.Float64frombits(uint64(b.IntVal)))
}
func FloatNeg(a Value) Value { return Float(-math.Float64frombits(uint64(a.IntVal))) }

func FloatEq(a, b Value) Value {
	return Bool(math.Float64frombits(uint64(a.IntVal)) == math.Float64frombits(uint64(b.IntVal)))
}
func FloatNeq(a, b Value) Value {
	return Bool(math.Float64frombits(uint64(a.IntVal)) != math.Float64frombits(uint64(b.IntVal)))
}
func FloatLt(a, b Value) Value {
	return Bool(math.Float64frombits(uint64(a.IntVal)) < math.Float64frombits(uint64(b.IntVal)))
}
func FloatLte(a, b Value) Value {
	return Bool(math.Float64frombits(uint64(a.IntVal)) <= math.Float64frombits(uint64(b.IntVal)))
}
func FloatGt(a, b Value) Value {
	return Bool(math.Float64frombits(uint64(a.IntVal)) > math.Float64frombits(uint64(b.IntVal)))
}
func FloatGte(a, b Value) Value {
	return Bool(math.Float64frombits(uint64(a.IntVal)) >= math.Float64frombits(uint64(b.IntVal)))
}

func Zshr(a Value, b Value) Value {
	return Int(IntZshr(a.IntVal, b.IntVal))
}

func Shl(a Value, b Value) Value {
	return Int(IntShl(a.IntVal, b.IntVal))
}

func Shr(a Value, b Value) Value {
	return Int(IntShr(a.IntVal, b.IntVal))
}

func BitAnd(a Value, b Value) Value {
	return Int(IntBitAnd(a.IntVal, b.IntVal))
}

func BitOr(a Value, b Value) Value {
	return Int(IntBitOr(a.IntVal, b.IntVal))
}

func BitXor(a Value, b Value) Value {
	return Int(IntBitXor(a.IntVal, b.IntVal))
}

func Array(v []Value) Value {
	ptr := new([]Value)
	*ptr = v
	return Value{Type: TypeArray, UnsafePtr: unsafe.Pointer(ptr)}
}

func Int32Array(v []int32) Value {
	ptr := new([]int32)
	*ptr = v
	return Value{Type: TypeInt32Array, UnsafePtr: unsafe.Pointer(ptr)}
}

func StringArray(v []string) Value {
	ptr := new([]string)
	*ptr = v
	return Value{Type: TypeStringArray, UnsafePtr: unsafe.Pointer(ptr)}
}

func IntArray(v []int64) Value {
	if len(v) == 0 {
		ptr := new([]int32)
		if v != nil {
			*ptr = []int32{}
		}
		return Value{Type: TypeInt32Array, UnsafePtr: unsafe.Pointer(ptr)}
	}
	compact := make([]int32, len(v))
	for i, x := range v {
		n := int32(x)
		if int64(n) != x {
			ptr := new([]int64)
			*ptr = v
			return Value{Type: TypeIntArray, UnsafePtr: unsafe.Pointer(ptr)}
		}
		compact[i] = n
	}
	ptr := new([]int32)
	*ptr = compact
	return Value{Type: TypeInt32Array, UnsafePtr: unsafe.Pointer(ptr)}
}

func EqualInt32Array(a, b []int32) bool {
	if len(a) != len(b) {
		return false
	}
	if len(a) == 0 {
		return true
	}
	return bytes.Equal(
		unsafe.Slice((*byte)(unsafe.Pointer(unsafe.SliceData(a))), len(a)*4),
		unsafe.Slice((*byte)(unsafe.Pointer(unsafe.SliceData(b))), len(b)*4),
	)
}

func EqualIntArray(a, b []int64) bool {
	if len(a) != len(b) {
		return false
	}
	if len(a) == 0 {
		return true
	}
	return bytes.Equal(
		unsafe.Slice((*byte)(unsafe.Pointer(unsafe.SliceData(a))), len(a)*8),
		unsafe.Slice((*byte)(unsafe.Pointer(unsafe.SliceData(b))), len(b)*8),
	)
}

func EqualIntArrayWithValue(native []int64, boxed Value) bool {
	switch boxed.Type {
	case TypeInt32Array:
		rhs := *(*[]int32)(boxed.UnsafePtr)
		if len(native) != len(rhs) {
			return false
		}
		for i, v := range rhs {
			if native[i] != int64(v) {
				return false
			}
		}
		return true
	case TypeIntArray:
		return EqualIntArray(native, *(*[]int64)(boxed.UnsafePtr))
	default:
		rhs := *(*[]Value)(boxed.UnsafePtr)
		if len(native) != len(rhs) {
			return false
		}
		for i, v := range rhs {
			if native[i] != v.IntVal {
				return false
			}
		}
		return true
	}
}

func EqualIntArrayValues(left, right Value) bool {
	switch left.Type {
	case TypeInt32Array:
		lhs := *(*[]int32)(left.UnsafePtr)
		switch right.Type {
		case TypeInt32Array:
			return EqualInt32Array(lhs, *(*[]int32)(right.UnsafePtr))
		case TypeIntArray:
			rhs := *(*[]int64)(right.UnsafePtr)
			if len(lhs) != len(rhs) {
				return false
			}
			for i, v := range lhs {
				if int64(v) != rhs[i] {
					return false
				}
			}
			return true
		default:
			rhs := *(*[]Value)(right.UnsafePtr)
			if len(lhs) != len(rhs) {
				return false
			}
			for i, v := range lhs {
				if int64(v) != rhs[i].IntVal {
					return false
				}
			}
			return true
		}
	case TypeIntArray:
		return EqualIntArrayWithValue(*(*[]int64)(left.UnsafePtr), right)
	default:
		lhs := *(*[]Value)(left.UnsafePtr)
		switch right.Type {
		case TypeInt32Array:
			rhs := *(*[]int32)(right.UnsafePtr)
			if len(lhs) != len(rhs) {
				return false
			}
			for i, v := range rhs {
				if lhs[i].IntVal != int64(v) {
					return false
				}
			}
			return true
		case TypeIntArray:
			rhs := *(*[]int64)(right.UnsafePtr)
			if len(lhs) != len(rhs) {
				return false
			}
			for i, v := range rhs {
				if lhs[i].IntVal != v {
					return false
				}
			}
			return true
		default:
			rhs := *(*[]Value)(right.UnsafePtr)
			if len(lhs) != len(rhs) {
				return false
			}
			for i := range lhs {
				if lhs[i].IntVal != rhs[i].IntVal {
					return false
				}
			}
			return true
		}
	}
}

func ArrayValues(arr Value) []Value {
	if arr.Type != TypeInt32Array && arr.Type != TypeIntArray && arr.Type != TypeStringArray {
		if arr.UnsafePtr == nil {
			return nil
		}
		return *(*[]Value)(arr.UnsafePtr)
	}
	return arrayValuesSlow(arr)
}

func arrayValuesSlow(arr Value) []Value {
	if arr.UnsafePtr == nil {
		return nil
	}
	if arr.Type == TypeInt32Array {
		src := *(*[]int32)(arr.UnsafePtr)
		out := make([]Value, len(src))
		for i, v := range src {
			out[i] = Int(int64(v))
		}
		return out
	}
	if arr.Type == TypeStringArray {
		src := *(*[]string)(arr.UnsafePtr)
		out := make([]Value, len(src))
		for i, v := range src {
			out[i] = Str(v)
		}
		return out
	}
	src := *(*[]int64)(arr.UnsafePtr)
	out := make([]Value, len(src))
	for i, v := range src {
		out[i] = Int(v)
	}
	return out
}

func UnboxIntArray(arr Value) []int64 {
	if arr.Type == TypeInt32Array {
		if arr.UnsafePtr == nil {
			return make([]int64, 0)
		}
		src := *(*[]int32)(arr.UnsafePtr)
		out := make([]int64, len(src))
		for i, v := range src {
			out[i] = int64(v)
		}
		return out
	}
	if arr.Type == TypeIntArray {
		if arr.UnsafePtr == nil {
			return make([]int64, 0)
		}
		src := *(*[]int64)(arr.UnsafePtr)
		out := make([]int64, len(src))
		copy(out, src)
		return out
	}
	src := *(*[]Value)(arr.UnsafePtr)
	out := make([]int64, len(src))
	for i, v := range src {
		out[i] = v.IntVal
	}
	return out
}

// BorrowIntArray returns a read-only []int64 view for pure PureScript workers
// without copying when the boxed value already holds a []int64 backing slice.
func BorrowIntArray(arr Value) []int64 {
	if arr.Type == TypeIntArray {
		if arr.UnsafePtr == nil {
			return make([]int64, 0)
		}
		return *(*[]int64)(arr.UnsafePtr)
	}
	return UnboxIntArray(arr)
}

func (v Value) PtrVal() any {
	return *(*any)(v.UnsafePtr)
}

func (v Value) BoolVal() bool {
	return v.IntVal != 0
}

func (v Value) AnyVal() any {
	switch v.Type {
	case TypeInt:
		return v.IntVal
	case TypeFloat:
		return math.Float64frombits(uint64(v.IntVal))
	case TypeString:
		return StrValue(v)
	case TypeBool:
		return v.IntVal != 0
	case TypeArray:
		return *(*[]Value)(v.UnsafePtr)
	case TypeIntArray, TypeInt32Array, TypeStringArray:
		return ArrayValues(v)
	case TypeRecord:
		return *(*map[string]Value)(v.UnsafePtr)
	case TypeAny:
		return v.PtrVal()
	case TypeConstructor:
		return v
	default:
		if v.Type >= TypeRecord0 && v.Type <= TypeRecordData {
			return RecordToMap(v)
		}
		if (v.Type >= TypeFunc && v.Type <= TypeFunc5) || (v.Type >= TypeFunc6 && v.Type <= TypeFunc11) {
			return v
		}
		if v.UnsafePtr != nil {
			return v
		}
		return nil
	}
}

// RecordFieldName is the exported field ABI for native PureScript records.
// Native Go structs keep their own field contract; callers select this mapping
// explicitly when reflecting on a PureScript record.
func RecordFieldName(label string) string {
	ascii := len(label) != 0
	for i := 0; i < len(label); i++ {
		c := label[i]
		letter := c == '_' || c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z'
		ascii = ascii && (letter || i > 0 && c >= '0' && c <= '9')
	}
	if ascii {
		return "P_" + label
	}
	const hex = "0123456789abcdef"
	var result strings.Builder
	result.WriteString("PX_")
	unit := func(value rune) {
		for shift := 12; shift >= 0; shift -= 4 {
			result.WriteByte(hex[(value>>shift)&15])
		}
	}
	for i := 0; i < len(label); {
		var char rune
		var width int
		if i+2 < len(label) && label[i] == 0xed && label[i+1] >= 0xa0 && label[i+1] <= 0xbf && label[i+2]&0xc0 == 0x80 {
			char, width = rune(label[i]&15)<<12|rune(label[i+1]&63)<<6|rune(label[i+2]&63), 3
		} else {
			char, width = utf8.DecodeRuneInString(label[i:])
		}
		i += width
		if char > 0xffff {
			char -= 0x10000
			unit(0xd800 + char>>10)
			unit(0xdc00 + char&0x3ff)
		} else {
			unit(char)
		}
	}
	return result.String()
}

// RecordFieldLabel reverses only canonical PureScript record field names.
// Preserve isolated surrogates as WTF-8, just like boxed record keys.
func RecordFieldLabel(name string) (string, bool) {
	if strings.HasPrefix(name, "P_") {
		label := name[2:]
		return label, RecordFieldName(label) == name
	}
	if !strings.HasPrefix(name, "PX_") || (len(name)-3)%4 != 0 {
		return "", false
	}
	units := make([]rune, 0, (len(name)-3)/4)
	for i := 3; i < len(name); i += 4 {
		var unit rune
		for _, digit := range name[i : i+4] {
			value := strings.IndexRune("0123456789abcdef", digit)
			if value < 0 {
				return "", false
			}
			unit = unit<<4 | rune(value)
		}
		units = append(units, unit)
	}
	var label strings.Builder
	for i := 0; i < len(units); i++ {
		unit := units[i]
		if unit >= 0xd800 && unit <= 0xdbff && i+1 < len(units) && units[i+1] >= 0xdc00 && units[i+1] <= 0xdfff {
			label.WriteRune(0x10000 + (unit-0xd800)<<10 + units[i+1] - 0xdc00)
			i++
		} else if unit >= 0xd800 && unit <= 0xdfff {
			label.WriteByte(0xe0 | byte(unit>>12))
			label.WriteByte(0x80 | byte(unit>>6&63))
			label.WriteByte(0x80 | byte(unit&63))
		} else {
			label.WriteRune(unit)
		}
	}
	text := label.String()
	return text, RecordFieldName(text) == name
}

// RecordShape is canonical immutable key-order metadata for >5-field records.
// Only compiler-generated shapes and updates preserving immutable Keys carry
// it; hand-constructed RecordDict records leave Shape nil so caller mutations
// to Keys are always caught by fallback key comparisons.
type RecordShape struct {
	Keys []string
}

var recordShapeCache sync.Map // string -> *RecordShape

func InternRecordShape(keys ...string) *RecordShape {
	var sb strings.Builder
	for _, k := range keys {
		sb.WriteString(fmt.Sprintf("%d:", len(k)))
		sb.WriteString(k)
	}
	cacheKey := sb.String()
	if existing, ok := recordShapeCache.Load(cacheKey); ok {
		return existing.(*RecordShape)
	}
	copied := make([]string, len(keys))
	copy(copied, keys)
	shape := &RecordShape{Keys: copied}
	actual, _ := recordShapeCache.LoadOrStore(cacheKey, shape)
	return actual.(*RecordShape)
}

type RecordData struct {
	Keys    []string
	Vals    []Value
	Shape   *RecordShape
	IntVals []int64
}

// JSONObject is an immutable, compact JSON object supplied by the parser.
// Accessors borrow it; APIs needing a mutable foreign map materialise their own
// container. The object and every string it contains own their storage.
type JSONObject interface {
	JSONLookup(string) (any, bool)
	JSONLength() int
	JSONEntry(int) (string, any)
}

type JSONObjectView struct {
	raw     map[string]any
	boxed   map[string]Value
	compact JSONObject
}

func ReadJSONObject(raw any) (JSONObjectView, bool) {
	switch value := raw.(type) {
	case map[string]any:
		return JSONObjectView{raw: value}, true
	case map[string]Value:
		return JSONObjectView{boxed: value}, true
	case JSONObject:
		return JSONObjectView{compact: value}, true
	case Value:
		return ReadJSONObjectValue(value)
	}
	return JSONObjectView{}, false
}

// Value-level FFI callers avoid allocating an interface box for a Value merely
// to unwrap its TypeAny payload again.
func ReadJSONObjectValue(value Value) (JSONObjectView, bool) {
	if value.Type == TypeAny && value.UnsafePtr != nil {
		return ReadJSONObject(value.PtrVal())
	}
	if value.Type == TypeRecord || (value.Type >= TypeRecord0 && value.Type <= TypeRecordData) {
		return JSONObjectView{boxed: RecordToMap(value)}, true
	}
	return JSONObjectView{}, false
}

func (object JSONObjectView) Lookup(key string) (any, bool) {
	if object.compact != nil {
		return object.compact.JSONLookup(key)
	}
	if object.boxed != nil {
		v, ok := object.boxed[key]
		return v, ok
	}
	v, ok := object.raw[key]
	return v, ok
}

func (object JSONObjectView) Length() int {
	if object.compact != nil {
		return object.compact.JSONLength()
	}
	return len(object.raw) + len(object.boxed)
}

func (object JSONObjectView) Each(visit func(string, any)) {
	if object.compact != nil {
		for i := 0; i < object.compact.JSONLength(); i++ {
			key, value := object.compact.JSONEntry(i)
			visit(key, value)
		}
		return
	}
	for key, value := range object.raw {
		visit(key, value)
	}
	for key, value := range object.boxed {
		visit(key, value)
	}
}

func (object JSONObjectView) Keys() []string {
	keys := make([]string, 0, object.Length())
	object.Each(func(key string, _ any) { keys = append(keys, key) })
	return keys
}

func RecordDict(keys []string, vals []Value) Value {
	r := &RecordData{Keys: keys, Vals: vals}
	return Value{Type: TypeRecordData, UnsafePtr: unsafe.Pointer(r)}
}

func RecordDictInterned(keys []string, vals []Value) Value {
	shape := InternRecordShape(keys...)
	var intVals []int64
	if len(vals) > 0 && len(vals) == len(shape.Keys) {
		allInts := true
		for i := range vals {
			if vals[i].Type != TypeInt {
				allInts = false
				break
			}
		}
		if allInts {
			intVals = make([]int64, len(vals))
			for i := range vals {
				intVals[i] = vals[i].IntVal
			}
		}
	}
	r := &RecordData{Keys: shape.Keys, Vals: vals, Shape: shape, IntVals: intVals}
	return Value{Type: TypeRecordData, UnsafePtr: unsafe.Pointer(r)}
}

var (
	recordKeys1Cache sync.Map
	recordKeys2Cache sync.Map
	recordKeys3Cache sync.Map
	recordKeys4Cache sync.Map
	recordKeys5Cache sync.Map
)

func recordKeysCacheKey(keys ...string) string {
	var sb strings.Builder
	for _, k := range keys {
		sb.WriteString(fmt.Sprintf("%d:", len(k)))
		sb.WriteString(k)
	}
	return sb.String()
}

func InternRecordKeys1(k0 string) *RecordKeys1 {
	key := recordKeysCacheKey(k0)
	if existing, ok := recordKeys1Cache.Load(key); ok {
		return existing.(*RecordKeys1)
	}
	actual, _ := recordKeys1Cache.LoadOrStore(key, &RecordKeys1{K0: k0})
	return actual.(*RecordKeys1)
}

func InternRecordKeys2(k0, k1 string) *RecordKeys2 {
	key := recordKeysCacheKey(k0, k1)
	if existing, ok := recordKeys2Cache.Load(key); ok {
		return existing.(*RecordKeys2)
	}
	actual, _ := recordKeys2Cache.LoadOrStore(key, &RecordKeys2{K0: k0, K1: k1})
	return actual.(*RecordKeys2)
}

func InternRecordKeys3(k0, k1, k2 string) *RecordKeys3 {
	key := recordKeysCacheKey(k0, k1, k2)
	if existing, ok := recordKeys3Cache.Load(key); ok {
		return existing.(*RecordKeys3)
	}
	actual, _ := recordKeys3Cache.LoadOrStore(key, &RecordKeys3{K0: k0, K1: k1, K2: k2})
	return actual.(*RecordKeys3)
}

func InternRecordKeys4(k0, k1, k2, k3 string) *RecordKeys4 {
	key := recordKeysCacheKey(k0, k1, k2, k3)
	if existing, ok := recordKeys4Cache.Load(key); ok {
		return existing.(*RecordKeys4)
	}
	actual, _ := recordKeys4Cache.LoadOrStore(key, &RecordKeys4{K0: k0, K1: k1, K2: k2, K3: k3})
	return actual.(*RecordKeys4)
}

func InternRecordKeys5(k0, k1, k2, k3, k4 string) *RecordKeys5 {
	key := recordKeysCacheKey(k0, k1, k2, k3, k4)
	if existing, ok := recordKeys5Cache.Load(key); ok {
		return existing.(*RecordKeys5)
	}
	actual, _ := recordKeys5Cache.LoadOrStore(key, &RecordKeys5{K0: k0, K1: k1, K2: k2, K3: k3, K4: k4})
	return actual.(*RecordKeys5)
}

type RecordData0 struct{}

func RecordDict0() Value { return Value{Type: TypeRecord0, UnsafePtr: unsafe.Pointer(&RecordData0{})} }

// RecordKeys1 is shape metadata. Callers must not mutate it after use.
type RecordKeys1 struct {
	K0 string
}

type RecordData1 struct {
	*RecordKeys1
	V0 Value
}

func RecordShape1(keys *RecordKeys1, v0 Value) Value {
	return Value{Type: TypeRecord1, UnsafePtr: unsafe.Pointer(&RecordData1{keys, v0})}
}

func RecordDict1(k0 string, v0 Value) Value {
	return RecordShape1(&RecordKeys1{k0}, v0)
}

// RecordKeys2 is shape metadata. Callers must not mutate it after use.
type RecordKeys2 struct {
	K0, K1 string
}

type RecordData2 struct {
	*RecordKeys2
	V0, V1 Value
}

// TypeRecord2 uses IntVal only as a representation discriminator:
// zero is RecordData2, one is RecordData2Int. The public record type is unchanged.
// Every payload access must go through RecordFields2; keys remain immutable.
const record2IntLayout int64 = 1

type RecordData2Int struct {
	*RecordKeys2
	V0, V1 int64
}

func RecordShape2(keys *RecordKeys2, v0, v1 Value) Value {
	if v0.Type == TypeInt && v1.Type == TypeInt {
		return Value{Type: TypeRecord2, IntVal: record2IntLayout, UnsafePtr: unsafe.Pointer(&RecordData2Int{keys, v0.IntVal, v1.IntVal})}
	}
	return Value{Type: TypeRecord2, UnsafePtr: unsafe.Pointer(&RecordData2{keys, v0, v1})}
}

// RecordFields2 reads either two-field payload layout without exposing a cast.
func RecordFields2(record Value) (*RecordKeys2, Value, Value) {
	if record.IntVal == record2IntLayout {
		r := (*RecordData2Int)(record.UnsafePtr)
		return r.RecordKeys2, Int(r.V0), Int(r.V1)
	}
	r := (*RecordData2)(record.UnsafePtr)
	return r.RecordKeys2, r.V0, r.V1
}

// RecordIntFields2 reads both IntVal payloads directly without boxing values.
func RecordIntFields2(record Value) (*RecordKeys2, int64, int64) {
	if record.IntVal == record2IntLayout {
		r := (*RecordData2Int)(record.UnsafePtr)
		return r.RecordKeys2, r.V0, r.V1
	}
	r := (*RecordData2)(record.UnsafePtr)
	return r.RecordKeys2, r.V0.IntVal, r.V1.IntVal
}

// RecordAt2 is the indexed access used when the compiler proves field order.
func RecordAt2(record Value, index int) Value {
	_, v0, v1 := RecordFields2(record)
	switch index {
	case 0:
		return v0
	case 1:
		return v1
	default:
		panic("two-field record index out of range")
	}
}

func RecordDict2(k0, k1 string, v0, v1 Value) Value {
	return RecordShape2(&RecordKeys2{k0, k1}, v0, v1)
}

// RecordKeys3 is shape metadata. Callers must not mutate it after use.
type RecordKeys3 struct {
	K0, K1, K2 string
}

type RecordData3 struct {
	*RecordKeys3
	V0, V1, V2 Value
}

func RecordShape3(keys *RecordKeys3, v0, v1, v2 Value) Value {
	return Value{Type: TypeRecord3, UnsafePtr: unsafe.Pointer(&RecordData3{keys, v0, v1, v2})}
}

func RecordDict3(k0, k1, k2 string, v0, v1, v2 Value) Value {
	return RecordShape3(&RecordKeys3{k0, k1, k2}, v0, v1, v2)
}

// RecordKeys4 is shape metadata. Callers must not mutate it after use.
type RecordKeys4 struct {
	K0, K1, K2, K3 string
}

type RecordData4 struct {
	*RecordKeys4
	V0, V1, V2, V3 Value
}

func RecordShape4(keys *RecordKeys4, v0, v1, v2, v3 Value) Value {
	return Value{Type: TypeRecord4, UnsafePtr: unsafe.Pointer(&RecordData4{keys, v0, v1, v2, v3})}
}

func RecordDict4(k0, k1, k2, k3 string, v0, v1, v2, v3 Value) Value {
	return RecordShape4(&RecordKeys4{k0, k1, k2, k3}, v0, v1, v2, v3)
}

// RecordKeys5 is shape metadata. Callers must not mutate it after use.
type RecordKeys5 struct {
	K0, K1, K2, K3, K4 string
}

type RecordData5 struct {
	*RecordKeys5
	V0, V1, V2, V3, V4 Value
}

func RecordShape5(keys *RecordKeys5, v0, v1, v2, v3, v4 Value) Value {
	return Value{Type: TypeRecord5, UnsafePtr: unsafe.Pointer(&RecordData5{keys, v0, v1, v2, v3, v4})}
}

func RecordDict5(k0, k1, k2, k3, k4 string, v0, v1, v2, v3, v4 Value) Value {
	return RecordShape5(&RecordKeys5{k0, k1, k2, k3, k4}, v0, v1, v2, v3, v4)
}

func UnboxObject(v Value) map[string]any {
	if v.Type == TypeAny {
		val := v.PtrVal()
		if m, ok := val.(map[string]any); ok {
			return m
		}
		if object, ok := val.(JSONObject); ok {
			result := make(map[string]any, object.JSONLength())
			for i := 0; i < object.JSONLength(); i++ {
				key, value := object.JSONEntry(i)
				result[key] = value
			}
			return result
		}
	}
	m := RecordToMap(v)
	res := make(map[string]any, len(m))
	for k, v2 := range m {
		res[k] = v2
	}
	return res
}

func RecordToMap(obj Value) map[string]Value {
	if obj.Type == TypeRecord {
		m := *(*map[string]Value)(obj.UnsafePtr)
		res := make(map[string]Value, len(m))
		for k, v := range m {
			res[k] = v
		}
		return res
	}
	if obj.Type == TypeAny {
		v := obj.PtrVal()
		if m, ok := v.(map[string]any); ok {
			res := make(map[string]Value, len(m))
			for k, v2 := range m {
				res[k] = Box(v2)
			}
			return res
		}
		if m, ok := v.(map[string]Value); ok {
			return m
		}
		if object, ok := v.(JSONObject); ok {
			result := make(map[string]Value, object.JSONLength())
			for i := 0; i < object.JSONLength(); i++ {
				key, value := object.JSONEntry(i)
				result[key] = Box(value)
			}
			return result
		}
	}
	res := make(map[string]Value)
	switch obj.Type {
	case TypeRecord0:
	case TypeRecord1:
		r := (*RecordData1)(obj.UnsafePtr)
		res[r.K0] = r.V0
	case TypeRecord2:
		keys, v0, v1 := RecordFields2(obj)
		res[keys.K0] = v0
		res[keys.K1] = v1
	case TypeRecord3:
		r := (*RecordData3)(obj.UnsafePtr)
		res[r.K0] = r.V0
		res[r.K1] = r.V1
		res[r.K2] = r.V2
	case TypeRecord4:
		r := (*RecordData4)(obj.UnsafePtr)
		res[r.K0] = r.V0
		res[r.K1] = r.V1
		res[r.K2] = r.V2
		res[r.K3] = r.V3
	case TypeRecord5:
		r := (*RecordData5)(obj.UnsafePtr)
		res[r.K0] = r.V0
		res[r.K1] = r.V1
		res[r.K2] = r.V2
		res[r.K3] = r.V3
		res[r.K4] = r.V4
	case TypeRecordData:
		r := (*RecordData)(obj.UnsafePtr)
		for i, k := range r.Keys {
			res[k] = r.Vals[i]
		}
	}
	return res
}

func CoerceToStruct[T any](val Value) *T {
	if val.Type == TypeConstructor {
		return (*T)(val.UnsafePtr)
	}
	res := new(T)
	// Empty and single-field records need neither a map nor key sorting.
	if val.Type == TypeRecord0 {
		return res
	}
	resVal := reflect.ValueOf(res).Elem()
	if val.Type == TypeRecord1 {
		resVal.Field(1).Set(reflect.ValueOf((*RecordData1)(val.UnsafePtr).V0))
		return res
	}
	m := RecordToMap(val)
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for i, k := range keys {
		resVal.Field(i + 1).Set(reflect.ValueOf(m[k]))
	}
	return res
}

func Record(m map[string]Value) Value {
	ptr := new(map[string]Value)
	*ptr = m
	return Value{Type: TypeRecord, UnsafePtr: unsafe.Pointer(ptr)}
}

func RecordGet(obj Value, key string) Value {
	if obj.Type == 0 {
		panic("Attempt to read property '" + key + "' on uninitialized value")
	}
	if obj.Type == TypeAny {
		v := obj.PtrVal()
		if m, ok := v.(map[string]any); ok {
			return Box(m[key])
		}
		if m, ok := v.(map[string]Value); ok {
			return m[key]
		}
		if object, ok := v.(JSONObject); ok {
			value, _ := object.JSONLookup(key)
			return Box(value)
		}
	}
	switch obj.Type {
	case TypeRecord:
		return (*(*map[string]Value)(obj.UnsafePtr))[key]
	case TypeRecord1:
		r := (*RecordData1)(obj.UnsafePtr)
		if r.K0 == key {
			return r.V0
		}
	case TypeRecord2:
		keys, v0, v1 := RecordFields2(obj)
		if keys.K0 == key {
			return v0
		}
		if keys.K1 == key {
			return v1
		}
	case TypeRecord3:
		r := (*RecordData3)(obj.UnsafePtr)
		if r.K0 == key {
			return r.V0
		}
		if r.K1 == key {
			return r.V1
		}
		if r.K2 == key {
			return r.V2
		}
	case TypeRecord4:
		r := (*RecordData4)(obj.UnsafePtr)
		if r.K0 == key {
			return r.V0
		}
		if r.K1 == key {
			return r.V1
		}
		if r.K2 == key {
			return r.V2
		}
		if r.K3 == key {
			return r.V3
		}
	case TypeRecord5:
		r := (*RecordData5)(obj.UnsafePtr)
		if r.K0 == key {
			return r.V0
		}
		if r.K1 == key {
			return r.V1
		}
		if r.K2 == key {
			return r.V2
		}
		if r.K3 == key {
			return r.V3
		}
		if r.K4 == key {
			return r.V4
		}
	case TypeRecordData:
		r := (*RecordData)(obj.UnsafePtr)
		for i, k := range r.Keys {
			if k == key {
				return r.Vals[i]
			}
		}
	case TypeConstructor:
		if getter, ok := StructGetters[obj.IntVal]; ok {
			return getter(obj.UnsafePtr, key)
		}
	}
	strVal := ""
	if obj.Type == TypeString && obj.UnsafePtr != nil {
		strVal = StrValue(obj)
	}
	panic(fmt.Sprintf("Key '%s' not found in record. Object type: %d, String value: '%s', Object: %+v\n", key, obj.Type, strVal, obj))
}

// RecordSet inserts or replaces a field without modifying the original record.
// Unlike RecordUpdateDict, it supports extending compact records with new keys.
func RecordSet(orig Value, key string, value Value) Value {
	switch orig.Type {
	case TypeRecord0:
		return RecordDict1(key, value)
	case TypeRecord1:
		r := (*RecordData1)(orig.UnsafePtr)
		if key == r.K0 {
			return RecordShape1(r.RecordKeys1, value)
		}
		return RecordDict2(r.K0, key, r.V0, value)
	case TypeRecord2:
		keys, v0, v1 := RecordFields2(orig)
		switch key {
		case keys.K0:
			return RecordShape2(keys, value, v1)
		case keys.K1:
			return RecordShape2(keys, v0, value)
		}
		return RecordDict3(keys.K0, keys.K1, key, v0, v1, value)
	case TypeRecord3:
		r := (*RecordData3)(orig.UnsafePtr)
		switch key {
		case r.K0:
			return RecordShape3(r.RecordKeys3, value, r.V1, r.V2)
		case r.K1:
			return RecordShape3(r.RecordKeys3, r.V0, value, r.V2)
		case r.K2:
			return RecordShape3(r.RecordKeys3, r.V0, r.V1, value)
		}
		return RecordDict4(r.K0, r.K1, r.K2, key, r.V0, r.V1, r.V2, value)
	case TypeRecord4:
		r := (*RecordData4)(orig.UnsafePtr)
		switch key {
		case r.K0:
			return RecordShape4(r.RecordKeys4, value, r.V1, r.V2, r.V3)
		case r.K1:
			return RecordShape4(r.RecordKeys4, r.V0, value, r.V2, r.V3)
		case r.K2:
			return RecordShape4(r.RecordKeys4, r.V0, r.V1, value, r.V3)
		case r.K3:
			return RecordShape4(r.RecordKeys4, r.V0, r.V1, r.V2, value)
		}
		return RecordDict5(r.K0, r.K1, r.K2, r.K3, key, r.V0, r.V1, r.V2, r.V3, value)
	case TypeRecord5:
		r := (*RecordData5)(orig.UnsafePtr)
		switch key {
		case r.K0:
			return RecordShape5(r.RecordKeys5, value, r.V1, r.V2, r.V3, r.V4)
		case r.K1:
			return RecordShape5(r.RecordKeys5, r.V0, value, r.V2, r.V3, r.V4)
		case r.K2:
			return RecordShape5(r.RecordKeys5, r.V0, r.V1, value, r.V3, r.V4)
		case r.K3:
			return RecordShape5(r.RecordKeys5, r.V0, r.V1, r.V2, value, r.V4)
		case r.K4:
			return RecordShape5(r.RecordKeys5, r.V0, r.V1, r.V2, r.V3, value)
		}
		return RecordDict([]string{r.K0, r.K1, r.K2, r.K3, r.K4, key}, []Value{r.V0, r.V1, r.V2, r.V3, r.V4, value})
	case TypeRecordData:
		r := (*RecordData)(orig.UnsafePtr)
		for i, field := range r.Keys {
			if field == key {
				vals := append([]Value(nil), r.Vals...)
				vals[i] = value
				return Value{Type: TypeRecordData, UnsafePtr: unsafe.Pointer(&RecordData{Keys: r.Keys, Vals: vals, Shape: r.Shape})}
			}
		}
		keys := make([]string, len(r.Keys)+1)
		copy(keys, r.Keys)
		keys[len(r.Keys)] = key
		vals := make([]Value, len(r.Vals)+1)
		copy(vals, r.Vals)
		vals[len(r.Vals)] = value
		return RecordDict(keys, vals)
	case TypeRecord:
		return recordMapSet(*(*map[string]Value)(orig.UnsafePtr), key, value)
	case TypeAny:
		switch m := orig.PtrVal().(type) {
		case map[string]Value:
			return recordMapSet(m, key, value)
		case map[string]any:
			result := make(map[string]Value, len(m)+1)
			for k, v := range m {
				result[k] = Box(v)
			}
			result[key] = value
			return Record(result)
		}
	}
	return recordMapSet(RecordToMap(orig), key, value)
}

func recordMapSet(orig map[string]Value, key string, value Value) Value {
	result := make(map[string]Value, len(orig)+1)
	for k, v := range orig {
		result[k] = v
	}
	result[key] = value
	return Record(result)
}

func RecordUpdateDict(orig Value, keys []string, vals []Value) Value {
	switch orig.Type {
	case TypeRecord1:
		clone := *(*RecordData1)(orig.UnsafePtr)
		for i, k := range keys {
			if clone.K0 == k {
				clone.V0 = vals[i]
			}
		}
		return Value{Type: TypeRecord1, UnsafePtr: unsafe.Pointer(&clone)}
	case TypeRecord2:
		shape, v0, v1 := RecordFields2(orig)
		for i, k := range keys {
			if shape.K0 == k {
				v0 = vals[i]
			} else if shape.K1 == k {
				v1 = vals[i]
			}
		}
		return RecordShape2(shape, v0, v1)
	case TypeRecord3:
		clone := *(*RecordData3)(orig.UnsafePtr)
		for i, k := range keys {
			if clone.K0 == k {
				clone.V0 = vals[i]
			} else if clone.K1 == k {
				clone.V1 = vals[i]
			} else if clone.K2 == k {
				clone.V2 = vals[i]
			}
		}
		return Value{Type: TypeRecord3, UnsafePtr: unsafe.Pointer(&clone)}
	case TypeRecord4:
		clone := *(*RecordData4)(orig.UnsafePtr)
		for i, k := range keys {
			if clone.K0 == k {
				clone.V0 = vals[i]
			} else if clone.K1 == k {
				clone.V1 = vals[i]
			} else if clone.K2 == k {
				clone.V2 = vals[i]
			} else if clone.K3 == k {
				clone.V3 = vals[i]
			}
		}
		return Value{Type: TypeRecord4, UnsafePtr: unsafe.Pointer(&clone)}
	case TypeRecord5:
		clone := *(*RecordData5)(orig.UnsafePtr)
		for i, k := range keys {
			if clone.K0 == k {
				clone.V0 = vals[i]
			} else if clone.K1 == k {
				clone.V1 = vals[i]
			} else if clone.K2 == k {
				clone.V2 = vals[i]
			} else if clone.K3 == k {
				clone.V3 = vals[i]
			} else if clone.K4 == k {
				clone.V4 = vals[i]
			}
		}
		return Value{Type: TypeRecord5, UnsafePtr: unsafe.Pointer(&clone)}
	case TypeRecordData:
		r := (*RecordData)(orig.UnsafePtr)
		newVals := make([]Value, len(r.Vals))
		copy(newVals, r.Vals)
		for i, k := range keys {
			for j, rk := range r.Keys {
				if rk == k {
					newVals[j] = vals[i]
					break
				}
			}
		}
		newR := RecordData{Keys: r.Keys, Vals: newVals, Shape: r.Shape}
		return Value{Type: TypeRecordData, UnsafePtr: unsafe.Pointer(&newR)}
	}
	m := RecordToMap(orig)
	for i, k := range keys {
		m[k] = vals[i]
	}
	return Record(m)
}

func RecordUpdate1(orig Value, k1 string, v1 Value) Value {
	return RecordUpdateDict(orig, []string{k1}, []Value{v1})
}

func RecordUpdate2(orig Value, k1 string, v1 Value, k2 string, v2 Value) Value {
	return RecordUpdateDict(orig, []string{k1, k2}, []Value{v1, v2})
}

func RecordUpdate3(orig Value, k1 string, v1 Value, k2 string, v2 Value, k3 string, v3 Value) Value {
	return RecordUpdateDict(orig, []string{k1, k2, k3}, []Value{v1, v2, v3})
}

func RecordUpdate(orig Value, updates map[string]Value) Value {
	keys := make([]string, 0, len(updates))
	vals := make([]Value, 0, len(updates))
	for k, v := range updates {
		keys = append(keys, k)
		vals = append(vals, v)
	}
	return RecordUpdateDict(orig, keys, vals)
}

type ConstructorData []Value

func Constructor(tag string, args []Value) Value {
	ptr := unsafe.Pointer(nil)
	if len(args) > 0 {
		ptr = unsafe.Pointer(&args[0])
	}
	return Value{Type: TypeConstructor, UnsafePtr: ptr}
}

type ConstructorData0 struct{}

func Constructor0(tag string) Value {
	return Value{Type: TypeConstructor, UnsafePtr: unsafe.Pointer(&ConstructorData0{})}
}

type ConstructorData1 struct{ V0 Value }

func Constructor1(tag string, v0 Value) Value {
	return Value{Type: TypeConstructor, UnsafePtr: unsafe.Pointer(&ConstructorData1{v0})}
}

type ConstructorData2 struct{ V0, V1 Value }

func Constructor2(tag string, v0, v1 Value) Value {
	return Value{Type: TypeConstructor, UnsafePtr: unsafe.Pointer(&ConstructorData2{v0, v1})}
}

type ConstructorData3 struct{ V0, V1, V2 Value }

func Constructor3(tag string, v0, v1, v2 Value) Value {
	return Value{Type: TypeConstructor, UnsafePtr: unsafe.Pointer(&ConstructorData3{v0, v1, v2})}
}

type ConstructorData4 struct{ V0, V1, V2, V3 Value }

func Constructor4(tag string, v0, v1, v2, v3 Value) Value {
	return Value{Type: TypeConstructor, UnsafePtr: unsafe.Pointer(&ConstructorData4{v0, v1, v2, v3})}
}

type ConstructorData5 struct{ V0, V1, V2, V3, V4 Value }

func Constructor5(tag string, v0, v1, v2, v3, v4 Value) Value {
	return Value{Type: TypeConstructor, UnsafePtr: unsafe.Pointer(&ConstructorData5{v0, v1, v2, v3, v4})}
}

var EscapeSink any
var escapeAlwaysFalse bool

func forceEscape(f any) {
	// As in Go's internal/abi.Escape, a variable (not a constant) keeps the
	// global store visible to escape analysis without executing it. Never set
	// escapeAlwaysFalse: closures must escape without a shared write or lock.
	if escapeAlwaysFalse {
		EscapeSink = f
	}
}

// Function with 1 arg (curried)
func Func(f func(Value) Value) Value {
	forceEscape(f)
	return Value{Type: TypeFunc, UnsafePtr: *(*unsafe.Pointer)(unsafe.Pointer(&f))}
}

// Function constructors
func Func2(f func(Value, Value) Value) Value {
	forceEscape(f)
	return Value{Type: TypeFunc2, UnsafePtr: *(*unsafe.Pointer)(unsafe.Pointer(&f))}
}
func Func3(f func(Value, Value, Value) Value) Value {
	forceEscape(f)
	return Value{Type: TypeFunc3, UnsafePtr: *(*unsafe.Pointer)(unsafe.Pointer(&f))}
}
func Func4(f func(Value, Value, Value, Value) Value) Value {
	forceEscape(f)
	return Value{Type: TypeFunc4, UnsafePtr: *(*unsafe.Pointer)(unsafe.Pointer(&f))}
}
func Func5(f func(Value, Value, Value, Value, Value) Value) Value {
	forceEscape(f)
	return Value{Type: TypeFunc5, UnsafePtr: *(*unsafe.Pointer)(unsafe.Pointer(&f))}
}
func Func6(f func(Value, Value, Value, Value, Value, Value) Value) Value {
	forceEscape(f)
	return Value{Type: TypeFunc6, UnsafePtr: *(*unsafe.Pointer)(unsafe.Pointer(&f))}
}
func Func7(f func(Value, Value, Value, Value, Value, Value, Value) Value) Value {
	forceEscape(f)
	return Value{Type: TypeFunc7, UnsafePtr: *(*unsafe.Pointer)(unsafe.Pointer(&f))}
}
func Func8(f func(Value, Value, Value, Value, Value, Value, Value, Value) Value) Value {
	forceEscape(f)
	return Value{Type: TypeFunc8, UnsafePtr: *(*unsafe.Pointer)(unsafe.Pointer(&f))}
}
func Func9(f func(Value, Value, Value, Value, Value, Value, Value, Value, Value) Value) Value {
	forceEscape(f)
	return Value{Type: TypeFunc9, UnsafePtr: *(*unsafe.Pointer)(unsafe.Pointer(&f))}
}
func Func10(f func(Value, Value, Value, Value, Value, Value, Value, Value, Value, Value) Value) Value {
	forceEscape(f)
	return Value{Type: TypeFunc10, UnsafePtr: *(*unsafe.Pointer)(unsafe.Pointer(&f))}
}
func Func11(f func(Value, Value, Value, Value, Value, Value, Value, Value, Value, Value, Value) Value) Value {
	forceEscape(f)
	return Value{Type: TypeFunc11, UnsafePtr: *(*unsafe.Pointer)(unsafe.Pointer(&f))}
}

func FuncAny(f any) Value {
	ptr := new(any)
	*ptr = f
	return Value{Type: TypeFunc, UnsafePtr: unsafe.Pointer(ptr)}
}

// Function data is immutable metadata owned by the callback, not a global
// registry. The ordinary entry remains callable through every Apply variant.
// Keeping a separate tag avoids depending on Go's closure memory layout.
type functionData struct {
	function Value
	data     any
}

func WithFunctionData(function Value, data any) Value {
	if function.Type == TypeFunctionData {
		function = (*functionData)(function.UnsafePtr).function
	}
	switch function.Type {
	case TypeFunc, TypeFunc2, TypeFunc3, TypeFunc4, TypeFunc5,
		TypeFunc6, TypeFunc7, TypeFunc8, TypeFunc9, TypeFunc10, TypeFunc11:
	default:
		panic("WithFunctionData requires a function")
	}
	return Value{Type: TypeFunctionData, UnsafePtr: unsafe.Pointer(&functionData{function, data})}
}

// Only an exact metadata type is recognized. Partial application returns the
// ordinary partial callback: metadata describing the full call no longer fits.
func FunctionData[T any](function Value) (T, bool) {
	if function.Type == TypeFunctionData && function.UnsafePtr != nil {
		data, ok := (*functionData)(function.UnsafePtr).data.(T)
		return data, ok
	}
	var zero T
	return zero, false
}

// IntActionFromInt64 exposes a zero-argument func() int64 action as func() int
// without a per-call trampoline closure on 64-bit targets, where int and int64
// share identical width and register ABI. 32-bit targets adapt cleanly.
func IntActionFromInt64(fn func() int64) func() int {
	if ^uint(0)>>63 == 1 {
		return *(*func() int)(unsafe.Pointer(&fn))
	}
	return func() int { return int(fn()) }
}

func Apply(f Value, arg Value) Value {
	if f.UnsafePtr == nil {
		panic(fmt.Sprintf("PANIC in Apply: f.Type = %v, f.UnsafePtr is nil", f.Type))
	}
	switch f.Type {
	case TypeFunctionData:
		f = (*functionData)(f.UnsafePtr).function
		if f.Type != TypeFunc || f.UnsafePtr == nil {
			return Apply(f, arg)
		}
		fallthrough
	case TypeFunc:
		return (*(*func(Value) Value)(unsafe.Pointer(&f.UnsafePtr)))(arg)
	case TypeFunc2:
		fn := *(*func(Value, Value) Value)(unsafe.Pointer(&f.UnsafePtr))
		return Func(func(a Value) Value { return fn(arg, a) })
	case TypeFunc3:
		fn := *(*func(Value, Value, Value) Value)(unsafe.Pointer(&f.UnsafePtr))
		return Func2(func(a, b Value) Value { return fn(arg, a, b) })
	case TypeFunc4:
		fn := *(*func(Value, Value, Value, Value) Value)(unsafe.Pointer(&f.UnsafePtr))
		return Func3(func(a, b, c Value) Value { return fn(arg, a, b, c) })
	case TypeFunc5:
		fn := *(*func(Value, Value, Value, Value, Value) Value)(unsafe.Pointer(&f.UnsafePtr))
		return Func4(func(a, b, c, d Value) Value { return fn(arg, a, b, c, d) })
	case TypeFunc6:
		fn := *(*func(Value, Value, Value, Value, Value, Value) Value)(unsafe.Pointer(&f.UnsafePtr))
		return Func5(func(a, b, c, d, e Value) Value { return fn(arg, a, b, c, d, e) })
	case TypeFunc7:
		fn := *(*func(Value, Value, Value, Value, Value, Value, Value) Value)(unsafe.Pointer(&f.UnsafePtr))
		return Func6(func(a, b, c, d, e, f Value) Value { return fn(arg, a, b, c, d, e, f) })
	case TypeFunc8:
		fn := *(*func(Value, Value, Value, Value, Value, Value, Value, Value) Value)(unsafe.Pointer(&f.UnsafePtr))
		return Func7(func(a, b, c, d, e, f, g Value) Value { return fn(arg, a, b, c, d, e, f, g) })
	case TypeFunc9:
		fn := *(*func(Value, Value, Value, Value, Value, Value, Value, Value, Value) Value)(unsafe.Pointer(&f.UnsafePtr))
		return Func8(func(a, b, c, d, e, f, g, h Value) Value { return fn(arg, a, b, c, d, e, f, g, h) })
	case TypeFunc10:
		fn := *(*func(Value, Value, Value, Value, Value, Value, Value, Value, Value, Value) Value)(unsafe.Pointer(&f.UnsafePtr))
		return Func9(func(a, b, c, d, e, f, g, h, i Value) Value { return fn(arg, a, b, c, d, e, f, g, h, i) })
	case TypeFunc11:
		fn := *(*func(Value, Value, Value, Value, Value, Value, Value, Value, Value, Value, Value) Value)(unsafe.Pointer(&f.UnsafePtr))
		return Func10(func(a, b, c, d, e, f, g, h, i, j Value) Value { return fn(arg, a, b, c, d, e, f, g, h, i, j) })
	default:
		panic("Attempted to apply a non-function")
	}
}

// ApplyN invokes saturated functions directly and captures partial arguments once.
// A smaller arity falls back to sequential application so curried functions and
// overapplication keep their evaluation order.
func Apply2(f Value, arg1, arg2 Value) Value {
	switch f.Type {
	case TypeFunctionData:
		return Apply2((*functionData)(f.UnsafePtr).function, arg1, arg2)
	case TypeFunc2:
		return (*(*func(Value, Value) Value)(unsafe.Pointer(&f.UnsafePtr)))(arg1, arg2)
	case TypeFunc3:
		fn := *(*func(Value, Value, Value) Value)(unsafe.Pointer(&f.UnsafePtr))
		return Func(func(a Value) Value { return fn(arg1, arg2, a) })
	case TypeFunc4:
		fn := *(*func(Value, Value, Value, Value) Value)(unsafe.Pointer(&f.UnsafePtr))
		return Func2(func(a, b Value) Value { return fn(arg1, arg2, a, b) })
	case TypeFunc5:
		fn := *(*func(Value, Value, Value, Value, Value) Value)(unsafe.Pointer(&f.UnsafePtr))
		return Func3(func(a, b, c Value) Value { return fn(arg1, arg2, a, b, c) })
	case TypeFunc6:
		fn := *(*func(Value, Value, Value, Value, Value, Value) Value)(unsafe.Pointer(&f.UnsafePtr))
		return Func4(func(a, b, c, d Value) Value { return fn(arg1, arg2, a, b, c, d) })
	case TypeFunc7:
		fn := *(*func(Value, Value, Value, Value, Value, Value, Value) Value)(unsafe.Pointer(&f.UnsafePtr))
		return Func5(func(a, b, c, d, e Value) Value { return fn(arg1, arg2, a, b, c, d, e) })
	case TypeFunc8:
		fn := *(*func(Value, Value, Value, Value, Value, Value, Value, Value) Value)(unsafe.Pointer(&f.UnsafePtr))
		return Func6(func(a, b, c, d, e, f Value) Value { return fn(arg1, arg2, a, b, c, d, e, f) })
	case TypeFunc9:
		fn := *(*func(Value, Value, Value, Value, Value, Value, Value, Value, Value) Value)(unsafe.Pointer(&f.UnsafePtr))
		return Func7(func(a, b, c, d, e, f, g Value) Value { return fn(arg1, arg2, a, b, c, d, e, f, g) })
	case TypeFunc10:
		fn := *(*func(Value, Value, Value, Value, Value, Value, Value, Value, Value, Value) Value)(unsafe.Pointer(&f.UnsafePtr))
		return Func8(func(a, b, c, d, e, f, g, h Value) Value { return fn(arg1, arg2, a, b, c, d, e, f, g, h) })
	case TypeFunc11:
		fn := *(*func(Value, Value, Value, Value, Value, Value, Value, Value, Value, Value, Value) Value)(unsafe.Pointer(&f.UnsafePtr))
		return Func9(func(a, b, c, d, e, f, g, h, i Value) Value { return fn(arg1, arg2, a, b, c, d, e, f, g, h, i) })
	}
	return Apply(Apply(f, arg1), arg2)
}

func Apply3(f Value, arg1, arg2, arg3 Value) Value {
	switch f.Type {
	case TypeFunctionData:
		return Apply3((*functionData)(f.UnsafePtr).function, arg1, arg2, arg3)
	case TypeFunc3:
		return (*(*func(Value, Value, Value) Value)(unsafe.Pointer(&f.UnsafePtr)))(arg1, arg2, arg3)
	case TypeFunc4:
		fn := *(*func(Value, Value, Value, Value) Value)(unsafe.Pointer(&f.UnsafePtr))
		return Func(func(a Value) Value { return fn(arg1, arg2, arg3, a) })
	case TypeFunc5:
		fn := *(*func(Value, Value, Value, Value, Value) Value)(unsafe.Pointer(&f.UnsafePtr))
		return Func2(func(a, b Value) Value { return fn(arg1, arg2, arg3, a, b) })
	case TypeFunc6:
		fn := *(*func(Value, Value, Value, Value, Value, Value) Value)(unsafe.Pointer(&f.UnsafePtr))
		return Func3(func(a, b, c Value) Value { return fn(arg1, arg2, arg3, a, b, c) })
	case TypeFunc7:
		fn := *(*func(Value, Value, Value, Value, Value, Value, Value) Value)(unsafe.Pointer(&f.UnsafePtr))
		return Func4(func(a, b, c, d Value) Value { return fn(arg1, arg2, arg3, a, b, c, d) })
	case TypeFunc8:
		fn := *(*func(Value, Value, Value, Value, Value, Value, Value, Value) Value)(unsafe.Pointer(&f.UnsafePtr))
		return Func5(func(a, b, c, d, e Value) Value { return fn(arg1, arg2, arg3, a, b, c, d, e) })
	case TypeFunc9:
		fn := *(*func(Value, Value, Value, Value, Value, Value, Value, Value, Value) Value)(unsafe.Pointer(&f.UnsafePtr))
		return Func6(func(a, b, c, d, e, f Value) Value { return fn(arg1, arg2, arg3, a, b, c, d, e, f) })
	case TypeFunc10:
		fn := *(*func(Value, Value, Value, Value, Value, Value, Value, Value, Value, Value) Value)(unsafe.Pointer(&f.UnsafePtr))
		return Func7(func(a, b, c, d, e, f, g Value) Value { return fn(arg1, arg2, arg3, a, b, c, d, e, f, g) })
	case TypeFunc11:
		fn := *(*func(Value, Value, Value, Value, Value, Value, Value, Value, Value, Value, Value) Value)(unsafe.Pointer(&f.UnsafePtr))
		return Func8(func(a, b, c, d, e, f, g, h Value) Value { return fn(arg1, arg2, arg3, a, b, c, d, e, f, g, h) })
	}
	return Apply(Apply2(f, arg1, arg2), arg3)
}

func Apply4(f Value, arg1, arg2, arg3, arg4 Value) Value {
	switch f.Type {
	case TypeFunctionData:
		return Apply4((*functionData)(f.UnsafePtr).function, arg1, arg2, arg3, arg4)
	case TypeFunc4:
		return (*(*func(Value, Value, Value, Value) Value)(unsafe.Pointer(&f.UnsafePtr)))(arg1, arg2, arg3, arg4)
	case TypeFunc5:
		fn := *(*func(Value, Value, Value, Value, Value) Value)(unsafe.Pointer(&f.UnsafePtr))
		return Func(func(a Value) Value { return fn(arg1, arg2, arg3, arg4, a) })
	case TypeFunc6:
		fn := *(*func(Value, Value, Value, Value, Value, Value) Value)(unsafe.Pointer(&f.UnsafePtr))
		return Func2(func(a, b Value) Value { return fn(arg1, arg2, arg3, arg4, a, b) })
	case TypeFunc7:
		fn := *(*func(Value, Value, Value, Value, Value, Value, Value) Value)(unsafe.Pointer(&f.UnsafePtr))
		return Func3(func(a, b, c Value) Value { return fn(arg1, arg2, arg3, arg4, a, b, c) })
	case TypeFunc8:
		fn := *(*func(Value, Value, Value, Value, Value, Value, Value, Value) Value)(unsafe.Pointer(&f.UnsafePtr))
		return Func4(func(a, b, c, d Value) Value { return fn(arg1, arg2, arg3, arg4, a, b, c, d) })
	case TypeFunc9:
		fn := *(*func(Value, Value, Value, Value, Value, Value, Value, Value, Value) Value)(unsafe.Pointer(&f.UnsafePtr))
		return Func5(func(a, b, c, d, e Value) Value { return fn(arg1, arg2, arg3, arg4, a, b, c, d, e) })
	case TypeFunc10:
		fn := *(*func(Value, Value, Value, Value, Value, Value, Value, Value, Value, Value) Value)(unsafe.Pointer(&f.UnsafePtr))
		return Func6(func(a, b, c, d, e, f Value) Value { return fn(arg1, arg2, arg3, arg4, a, b, c, d, e, f) })
	case TypeFunc11:
		fn := *(*func(Value, Value, Value, Value, Value, Value, Value, Value, Value, Value, Value) Value)(unsafe.Pointer(&f.UnsafePtr))
		return Func7(func(a, b, c, d, e, f, g Value) Value { return fn(arg1, arg2, arg3, arg4, a, b, c, d, e, f, g) })
	}
	return Apply(Apply3(f, arg1, arg2, arg3), arg4)
}

func Apply5(f Value, arg1, arg2, arg3, arg4, arg5 Value) Value {
	switch f.Type {
	case TypeFunctionData:
		return Apply5((*functionData)(f.UnsafePtr).function, arg1, arg2, arg3, arg4, arg5)
	case TypeFunc5:
		return (*(*func(Value, Value, Value, Value, Value) Value)(unsafe.Pointer(&f.UnsafePtr)))(arg1, arg2, arg3, arg4, arg5)
	case TypeFunc6:
		fn := *(*func(Value, Value, Value, Value, Value, Value) Value)(unsafe.Pointer(&f.UnsafePtr))
		return Func(func(a Value) Value { return fn(arg1, arg2, arg3, arg4, arg5, a) })
	case TypeFunc7:
		fn := *(*func(Value, Value, Value, Value, Value, Value, Value) Value)(unsafe.Pointer(&f.UnsafePtr))
		return Func2(func(a, b Value) Value { return fn(arg1, arg2, arg3, arg4, arg5, a, b) })
	case TypeFunc8:
		fn := *(*func(Value, Value, Value, Value, Value, Value, Value, Value) Value)(unsafe.Pointer(&f.UnsafePtr))
		return Func3(func(a, b, c Value) Value { return fn(arg1, arg2, arg3, arg4, arg5, a, b, c) })
	case TypeFunc9:
		fn := *(*func(Value, Value, Value, Value, Value, Value, Value, Value, Value) Value)(unsafe.Pointer(&f.UnsafePtr))
		return Func4(func(a, b, c, d Value) Value { return fn(arg1, arg2, arg3, arg4, arg5, a, b, c, d) })
	case TypeFunc10:
		fn := *(*func(Value, Value, Value, Value, Value, Value, Value, Value, Value, Value) Value)(unsafe.Pointer(&f.UnsafePtr))
		return Func5(func(a, b, c, d, e Value) Value { return fn(arg1, arg2, arg3, arg4, arg5, a, b, c, d, e) })
	case TypeFunc11:
		fn := *(*func(Value, Value, Value, Value, Value, Value, Value, Value, Value, Value, Value) Value)(unsafe.Pointer(&f.UnsafePtr))
		return Func6(func(a, b, c, d, e, f Value) Value { return fn(arg1, arg2, arg3, arg4, arg5, a, b, c, d, e, f) })
	}
	return Apply(Apply4(f, arg1, arg2, arg3, arg4), arg5)
}

func ArrayAccess(arr Value, index int) Value {
	if arr.Type == TypeInt32Array {
		return Int(int64((*(*[]int32)(arr.UnsafePtr))[index]))
	}
	if arr.Type == TypeIntArray {
		return Int((*(*[]int64)(arr.UnsafePtr))[index])
	}
	if arr.Type == TypeStringArray {
		return Str((*(*[]string)(arr.UnsafePtr))[index])
	}
	return (*(*[]Value)(arr.UnsafePtr))[index]
}

func ArrayLength(arr Value) int {
	if arr.Type == TypeInt32Array {
		return len(*(*[]int32)(arr.UnsafePtr))
	}
	if arr.Type == TypeIntArray {
		return len(*(*[]int64)(arr.UnsafePtr))
	}
	if arr.Type == TypeStringArray {
		return len(*(*[]string)(arr.UnsafePtr))
	}
	return len(*(*[]Value)(arr.UnsafePtr))
}

//go:noinline
func SumIntArraySlow(input Value, result int64) int64 {
	if input.Type == TypeIntArray {
		items := *(*[]int64)(input.UnsafePtr)
		if len(items) < 8 {
			for _, v := range items {
				result = IntAdd(result, v)
			}
			return result
		}
		var a0, a1, a2, a3, a4, a5, a6, a7 int64
		i := 0
		for ; i <= len(items)-8; i += 8 {
			c := items[i : i+8]
			a0 += c[0]
			a1 += c[1]
			a2 += c[2]
			a3 += c[3]
			a4 += c[4]
			a5 += c[5]
			a6 += c[6]
			a7 += c[7]
		}
		for ; i < len(items); i++ {
			a0 += items[i]
		}
		return int64(int32((((result + a0) + (a1 + a2)) + ((a3 + a4) + (a5 + a6))) + a7))
	}
	items := *(*[]Value)(input.UnsafePtr)
	if len(items) < 8 {
		for _, v := range items {
			result = IntAdd(result, v.IntVal)
		}
		return result
	}
	var a0, a1, a2, a3, a4, a5, a6, a7 int64
	i := 0
	for ; i <= len(items)-8; i += 8 {
		c := items[i : i+8]
		a0 += c[0].IntVal
		a1 += c[1].IntVal
		a2 += c[2].IntVal
		a3 += c[3].IntVal
		a4 += c[4].IntVal
		a5 += c[5].IntVal
		a6 += c[6].IntVal
		a7 += c[7].IntVal
	}
	for ; i < len(items); i++ {
		a0 += items[i].IntVal
	}
	return int64(int32((((result + a0) + (a1 + a2)) + ((a3 + a4) + (a5 + a6))) + a7))
}

func Any(v any) Value {
	switch val := v.(type) {
	case Value:
		return val
	case foreignUndefined:
		return Value{}
	}
	ptr := new(any)
	*ptr = v
	return Value{Type: TypeAny, UnsafePtr: unsafe.Pointer(ptr)}
}

func UncurriedApp2(fn Value, a, b Value) Value {
	if fn.Type == TypeFunctionData {
		return Apply2(fn, a, b)
	}
	if fn.Type == TypeFunc2 {
		return (*(*func(Value, Value) Value)(unsafe.Pointer(&fn.UnsafePtr)))(a, b)
	}
	return Apply(Apply(fn, a), b)
}

func UncurriedApp3(fn Value, a, b, c Value) Value {
	if fn.Type == TypeFunctionData {
		return Apply3(fn, a, b, c)
	}
	if fn.Type == TypeFunc3 {
		return (*(*func(Value, Value, Value) Value)(unsafe.Pointer(&fn.UnsafePtr)))(a, b, c)
	}
	return Apply(Apply(Apply(fn, a), b), c)
}

func UncurriedApp4(fn Value, a, b, c, d Value) Value {
	if fn.Type == TypeFunctionData {
		return Apply4(fn, a, b, c, d)
	}
	if fn.Type == TypeFunc4 {
		return (*(*func(Value, Value, Value, Value) Value)(unsafe.Pointer(&fn.UnsafePtr)))(a, b, c, d)
	}
	return Apply(Apply(Apply(Apply(fn, a), b), c), d)
}

func UncurriedApp5(fn Value, a, b, c, d, e Value) Value {
	if fn.Type == TypeFunctionData {
		return Apply5(fn, a, b, c, d, e)
	}
	if fn.Type == TypeFunc5 {
		return (*(*func(Value, Value, Value, Value, Value) Value)(unsafe.Pointer(&fn.UnsafePtr)))(a, b, c, d, e)
	}
	return Apply(Apply(Apply(Apply(Apply(fn, a), b), c), d), e)
}

func UncurriedApp6(fn Value, a, b, c, d, e, f Value) Value {
	return Apply6(fn, a, b, c, d, e, f)
}

func UncurriedApp7(fn Value, a, b, c, d, e, f, g Value) Value {
	return Apply7(fn, a, b, c, d, e, f, g)
}

func UncurriedApp8(fn Value, a, b, c, d, e, f, g, h Value) Value {
	return Apply8(fn, a, b, c, d, e, f, g, h)
}

func UncurriedApp9(fn Value, a, b, c, d, e, f, g, h, i Value) Value {
	return Apply9(fn, a, b, c, d, e, f, g, h, i)
}

func UncurriedApp10(fn Value, a, b, c, d, e, f, g, h, i, j Value) Value {
	return Apply10(fn, a, b, c, d, e, f, g, h, i, j)
}

func UncurriedApp(fn Value, args ...Value) Value {
	// Zero-argument closures use the same ignored-Unit wrapper as Box and Wrap0.
	if len(args) == 0 {
		return Apply(fn, Value{})
	}
	res := fn
	for _, arg := range args {
		res = Apply(res, arg)
	}
	return res
}

func Unbox[T any](v any) T {
	if res, ok := v.(T); ok {
		return res
	}

	val, isValue := v.(Value)
	if !isValue {
		var t any = *new(T)
		// If T is any, we can just return v
		switch t.(type) {
		case any:
			return any(v).(T)
		}
		panic("Cannot unbox non-Value into T")
	}

	var t any = *new(T)
	switch t.(type) {
	case int32:
		if val.Type == TypeAny && val.UnsafePtr != nil {
			if value, ok := (*(*any)(val.UnsafePtr)).(int32); ok {
				return any(value).(T)
			}
		}
		return any(int32(Unbox[int64](val))).(T)
	case int64:
		if val.Type == TypeFloat {
			return any(int64(math.Float64frombits(uint64(val.IntVal)))).(T)
		}
		if val.Type == TypeAny {
			if ptr := val.UnsafePtr; ptr != nil {
				v := *(*any)(ptr)
				if f, ok := v.(float64); ok {
					return any(int64(f)).(T)
				}
				if i, ok := v.(int); ok {
					return any(int64(i)).(T)
				}
				if i, ok := v.(int64); ok {
					return any(i).(T)
				}
			}
			return any(int64(0)).(T)
		}
		return any(val.IntVal).(T)
	case int:
		if val.Type == TypeFloat {
			return any(int(math.Float64frombits(uint64(val.IntVal)))).(T)
		}
		if val.Type == TypeAny {
			if ptr := val.UnsafePtr; ptr != nil {
				v := *(*any)(ptr)
				if f, ok := v.(float64); ok {
					return any(int(f)).(T)
				}
				if i, ok := v.(int); ok {
					return any(i).(T)
				}
				if i, ok := v.(int64); ok {
					return any(int(i)).(T)
				}
			}
			return any(0).(T)
		}
		return any(int(val.IntVal)).(T)
	case string:
		return any(StrValue(val)).(T)
	case float64:
		if val.Type == TypeInt {
			return any(float64(val.IntVal)).(T)
		}
		if val.Type == TypeAny {
			if ptr := val.UnsafePtr; ptr != nil {
				v := *(*any)(ptr)
				if f, ok := v.(float64); ok {
					return any(f).(T)
				}
				if i, ok := v.(int); ok {
					return any(float64(i)).(T)
				}
				if i, ok := v.(int64); ok {
					return any(float64(i)).(T)
				}
			}
			return any(float64(0)).(T)
		}
		return any(math.Float64frombits(uint64(val.IntVal))).(T)
	case bool:
		return any(val.IntVal == 1).(T)
	case Value:
		return any(val).(T)
	case map[string]any:
		m := RecordToMap(val)
		res := make(map[string]any, len(m))
		for k, v2 := range m {
			res[k] = v2
		}
		return any(res).(T)
	case []Value:
		return any(ArrayValues(val)).(T)
	case []int64:
		return any(UnboxIntArray(val)).(T)
	case []string:
		if val.Type == TypeStringArray {
			if val.UnsafePtr == nil {
				return any(make([]string, 0)).(T)
			}
			src := *(*[]string)(val.UnsafePtr)
			out := make([]string, len(src))
			copy(out, src)
			return any(out).(T)
		}
		src := ArrayValues(val)
		out := make([]string, len(src))
		for i, v2 := range src {
			out[i] = StrValue(v2)
		}
		return any(out).(T)
	case []any:
		if val.Type == TypeInt32Array {
			var src []int32
			if val.UnsafePtr != nil {
				src = *(*[]int32)(val.UnsafePtr)
			}
			res := make([]any, len(src))
			for i, v2 := range src {
				res[i] = Int(int64(v2))
			}
			return any(res).(T)
		}
		if val.Type == TypeIntArray {
			var src []int64
			if val.UnsafePtr != nil {
				src = *(*[]int64)(val.UnsafePtr)
			}
			res := make([]any, len(src))
			for i, v2 := range src {
				res[i] = Int(v2)
			}
			return any(res).(T)
		}
		if val.Type == TypeStringArray {
			var src []string
			if val.UnsafePtr != nil {
				src = *(*[]string)(val.UnsafePtr)
			}
			res := make([]any, len(src))
			for i, v2 := range src {
				res[i] = Str(v2)
			}
			return any(res).(T)
		}
		arr := *(*[]Value)(val.UnsafePtr)
		res := make([]any, len(arr))
		for i, v2 := range arr {
			res[i] = v2
		}
		return any(res).(T)
	case func(any) any:
		res := func(arg any) any { return Apply(val, Box(arg)) }
		return any(res).(T)
	default:
		if val.Type == TypeAny {
			return (*(*any)(val.UnsafePtr)).(T)
		}
		return *(*T)(val.UnsafePtr)
	}
}

func Box[T any](val T) Value {
	if v, ok := any(val).(Value); ok {
		return v
	}
	return boxSlow(any(val))
}

// Keep Value identity small enough to inline at generated FFI boundaries.
// All other representations retain the same conversion and closure behavior.
func boxSlow(val any) Value {
	switch v := val.(type) {
	case int32:
		return Int(int64(v))
	case int64:
		return Int(v)
	case int:
		return Int(int64(v))
	case string:
		return Str(v)
	case float64:
		return Value{Type: TypeFloat, IntVal: int64(math.Float64bits(v))}
	case bool:
		if v {
			return Value{Type: TypeBool, IntVal: 1}
		}
		return Value{Type: TypeBool, IntVal: 0}
	case func():
		return Func(func(_ Value) Value { v(); return Value{} })
	case func() bool:
		return Func(func(_ Value) Value {
			if v() {
				return Value{Type: TypeBool, IntVal: 1}
			}
			return Value{Type: TypeBool, IntVal: 0}
		})
	case func() int:
		return Func(func(_ Value) Value { return Int(int64(v())) })
	case func() int64:
		return Func(func(_ Value) Value { return Int(v()) })
	case func() string:
		return Func(func(_ Value) Value { return Str(v()) })
	case func() float64:
		return Func(func(_ Value) Value { return Value{Type: TypeFloat, IntVal: int64(math.Float64bits(v()))} })
	case func() Value:
		return Func(func(_ Value) Value { return v() })
	case func(Value) Value:
		return Func(v)
	case func(any) func(any) func(any) any:
		return Func(func(arg Value) Value { return Box(v(arg)) })
	case func(any) func(any) any:
		return Func(func(arg Value) Value { return Box(v(arg)) })
	case func(any) any:
		return Func(func(arg Value) Value { return Box(v(arg)) })
	case []any:
		arr := make([]Value, len(v))
		for i, val := range v {
			arr[i] = Box(val)
		}
		return Array(arr)
	case []Value:
		return Array(v)
	case []string:
		return StringArray(v)

	case Value:
		return v
	default:
		return Any(v)
	}
}

func Wrap0[R any](f func() R) Value {
	return Func(func(_ Value) Value {
		return Box(f())
	})
}

func Wrap1[A, R any](f func(A) R) Value {
	return Func(func(a Value) Value {
		return Box(f(Unbox[A](a)))
	})
}

func Wrap2[A, B, R any](f func(A, B) R) Value {
	return Func2(func(a, b Value) Value {
		return Box(f(Unbox[A](a), Unbox[B](b)))
	})
}

func Wrap3[A, B, C, R any](f func(A, B, C) R) Value {
	return Func3(func(a, b, c Value) Value {
		return Box(f(Unbox[A](a), Unbox[B](b), Unbox[C](c)))
	})
}

func Wrap4[A, B, C, D, R any](f func(A, B, C, D) R) Value {
	return Func4(func(a, b, c, d Value) Value {
		return Box(f(Unbox[A](a), Unbox[B](b), Unbox[C](c), Unbox[D](d)))
	})
}

func Wrap5[A, B, C, D, E, R any](f func(A, B, C, D, E) R) Value {
	return Func5(func(a, b, c, d, e Value) Value {
		return Box(f(Unbox[A](a), Unbox[B](b), Unbox[C](c), Unbox[D](d), Unbox[E](e)))
	})
}

func Apply6(f Value, arg1, arg2, arg3, arg4, arg5, arg6 Value) Value {
	switch f.Type {
	case TypeFunctionData:
		return Apply6((*functionData)(f.UnsafePtr).function, arg1, arg2, arg3, arg4, arg5, arg6)
	case TypeFunc6:
		return (*(*func(Value, Value, Value, Value, Value, Value) Value)(unsafe.Pointer(&f.UnsafePtr)))(arg1, arg2, arg3, arg4, arg5, arg6)
	case TypeFunc7:
		fn := *(*func(Value, Value, Value, Value, Value, Value, Value) Value)(unsafe.Pointer(&f.UnsafePtr))
		return Func(func(a Value) Value { return fn(arg1, arg2, arg3, arg4, arg5, arg6, a) })
	case TypeFunc8:
		fn := *(*func(Value, Value, Value, Value, Value, Value, Value, Value) Value)(unsafe.Pointer(&f.UnsafePtr))
		return Func2(func(a, b Value) Value { return fn(arg1, arg2, arg3, arg4, arg5, arg6, a, b) })
	case TypeFunc9:
		fn := *(*func(Value, Value, Value, Value, Value, Value, Value, Value, Value) Value)(unsafe.Pointer(&f.UnsafePtr))
		return Func3(func(a, b, c Value) Value { return fn(arg1, arg2, arg3, arg4, arg5, arg6, a, b, c) })
	case TypeFunc10:
		fn := *(*func(Value, Value, Value, Value, Value, Value, Value, Value, Value, Value) Value)(unsafe.Pointer(&f.UnsafePtr))
		return Func4(func(a, b, c, d Value) Value { return fn(arg1, arg2, arg3, arg4, arg5, arg6, a, b, c, d) })
	case TypeFunc11:
		fn := *(*func(Value, Value, Value, Value, Value, Value, Value, Value, Value, Value, Value) Value)(unsafe.Pointer(&f.UnsafePtr))
		return Func5(func(a, b, c, d, e Value) Value { return fn(arg1, arg2, arg3, arg4, arg5, arg6, a, b, c, d, e) })
	}
	return Apply(Apply5(f, arg1, arg2, arg3, arg4, arg5), arg6)
}

func Apply7(f Value, arg1, arg2, arg3, arg4, arg5, arg6, arg7 Value) Value {
	switch f.Type {
	case TypeFunctionData:
		return Apply7((*functionData)(f.UnsafePtr).function, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
	case TypeFunc7:
		return (*(*func(Value, Value, Value, Value, Value, Value, Value) Value)(unsafe.Pointer(&f.UnsafePtr)))(arg1, arg2, arg3, arg4, arg5, arg6, arg7)
	case TypeFunc8:
		fn := *(*func(Value, Value, Value, Value, Value, Value, Value, Value) Value)(unsafe.Pointer(&f.UnsafePtr))
		return Func(func(a Value) Value { return fn(arg1, arg2, arg3, arg4, arg5, arg6, arg7, a) })
	case TypeFunc9:
		fn := *(*func(Value, Value, Value, Value, Value, Value, Value, Value, Value) Value)(unsafe.Pointer(&f.UnsafePtr))
		return Func2(func(a, b Value) Value { return fn(arg1, arg2, arg3, arg4, arg5, arg6, arg7, a, b) })
	case TypeFunc10:
		fn := *(*func(Value, Value, Value, Value, Value, Value, Value, Value, Value, Value) Value)(unsafe.Pointer(&f.UnsafePtr))
		return Func3(func(a, b, c Value) Value { return fn(arg1, arg2, arg3, arg4, arg5, arg6, arg7, a, b, c) })
	case TypeFunc11:
		fn := *(*func(Value, Value, Value, Value, Value, Value, Value, Value, Value, Value, Value) Value)(unsafe.Pointer(&f.UnsafePtr))
		return Func4(func(a, b, c, d Value) Value { return fn(arg1, arg2, arg3, arg4, arg5, arg6, arg7, a, b, c, d) })
	}
	return Apply(Apply6(f, arg1, arg2, arg3, arg4, arg5, arg6), arg7)
}

func AnyToValue(val any) Value {
	// The old ST FFI reads the reference envelope through this bridge too.
	// Return a Value copy so subsequent writes cannot change read snapshots.
	if cell, ok := val.(*stValueCell); ok {
		return cell.value
	}
	if val == nil {
		return Value{}
	}
	return Box(val)
}

func ValueToAny(val Value) any {
	switch val.Type {
	case TypeInt:
		return val.IntVal
	case TypeFloat:
		return *(*float64)(val.UnsafePtr)
	case TypeString:
		return StrValue(val)
	case TypeBool:
		return val.IntVal != 0
	}
	return val
}

func Apply8(f Value, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8 Value) Value {
	switch f.Type {
	case TypeFunctionData:
		return Apply8((*functionData)(f.UnsafePtr).function, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
	case TypeFunc8:
		return (*(*func(Value, Value, Value, Value, Value, Value, Value, Value) Value)(unsafe.Pointer(&f.UnsafePtr)))(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
	case TypeFunc9:
		fn := *(*func(Value, Value, Value, Value, Value, Value, Value, Value, Value) Value)(unsafe.Pointer(&f.UnsafePtr))
		return Func(func(a Value) Value { return fn(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, a) })
	case TypeFunc10:
		fn := *(*func(Value, Value, Value, Value, Value, Value, Value, Value, Value, Value) Value)(unsafe.Pointer(&f.UnsafePtr))
		return Func2(func(a, b Value) Value { return fn(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, a, b) })
	case TypeFunc11:
		fn := *(*func(Value, Value, Value, Value, Value, Value, Value, Value, Value, Value, Value) Value)(unsafe.Pointer(&f.UnsafePtr))
		return Func3(func(a, b, c Value) Value { return fn(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, a, b, c) })
	}
	return Apply(Apply7(f, arg1, arg2, arg3, arg4, arg5, arg6, arg7), arg8)
}

func Apply9(f Value, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9 Value) Value {
	switch f.Type {
	case TypeFunctionData:
		return Apply9((*functionData)(f.UnsafePtr).function, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9)
	case TypeFunc9:
		return (*(*func(Value, Value, Value, Value, Value, Value, Value, Value, Value) Value)(unsafe.Pointer(&f.UnsafePtr)))(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9)
	case TypeFunc10:
		fn := *(*func(Value, Value, Value, Value, Value, Value, Value, Value, Value, Value) Value)(unsafe.Pointer(&f.UnsafePtr))
		return Func(func(a Value) Value { return fn(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, a) })
	case TypeFunc11:
		fn := *(*func(Value, Value, Value, Value, Value, Value, Value, Value, Value, Value, Value) Value)(unsafe.Pointer(&f.UnsafePtr))
		return Func2(func(a, b Value) Value { return fn(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, a, b) })
	}
	return Apply(Apply8(f, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8), arg9)
}

func Apply10(f Value, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10 Value) Value {
	switch f.Type {
	case TypeFunctionData:
		return Apply10((*functionData)(f.UnsafePtr).function, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10)
	case TypeFunc10:
		return (*(*func(Value, Value, Value, Value, Value, Value, Value, Value, Value, Value) Value)(unsafe.Pointer(&f.UnsafePtr)))(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10)
	case TypeFunc11:
		fn := *(*func(Value, Value, Value, Value, Value, Value, Value, Value, Value, Value, Value) Value)(unsafe.Pointer(&f.UnsafePtr))
		return Func(func(a Value) Value { return fn(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, a) })
	}
	return Apply(Apply9(f, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9), arg10)
}

func ExtractVariant(variant interface{}) (string, interface{}, bool) {
	if val, ok := variant.(Value); ok {
		if val.Type == TypeRecord {
			m := *(*map[string]Value)(val.UnsafePtr)
			if typVal, ok := m["type"]; ok {
				if typVal.Type == TypeString {
					str := StrValue(typVal)
					valVal, hasVal := m["value"]
					if hasVal {
						return str, valVal, true
					}
					return str, nil, true
				}
			}
		} else if val.Type == TypeRecord2 {
			keys, v0, v1 := RecordFields2(val)
			var str string
			var valVal interface{}
			hasVal := false
			if keys.K0 == "type" && v0.Type == TypeString {
				str = StrValue(v0)
			} else if keys.K1 == "type" && v1.Type == TypeString {
				str = StrValue(v1)
			} else {
				return "", nil, false
			}

			if keys.K0 == "value" {
				valVal = v0
				hasVal = true
			} else if keys.K1 == "value" {
				valVal = v1
				hasVal = true
			}

			if hasVal {
				return str, valVal, true
			}
			return str, nil, true
		} else if val.Type == TypeRecord3 {
			rec := (*RecordData3)(val.UnsafePtr)
			var str string
			var valVal interface{}
			hasVal := false

			if rec.K0 == "type" && rec.V0.Type == TypeString {
				str = StrValue(rec.V0)
			} else if rec.K1 == "type" && rec.V1.Type == TypeString {
				str = StrValue(rec.V1)
			} else if rec.K2 == "type" && rec.V2.Type == TypeString {
				str = StrValue(rec.V2)
			} else {
				return "", nil, false
			}

			if rec.K0 == "value" {
				valVal = rec.V0
				hasVal = true
			} else if rec.K1 == "value" {
				valVal = rec.V1
				hasVal = true
			} else if rec.K2 == "value" {
				valVal = rec.V2
				hasVal = true
			}

			if hasVal {
				return str, valVal, true
			}
			return str, nil, true
		} else if val.Type == TypeString {
			str := StrValue(val)
			return str, nil, true
		}
	}
	return "", nil, false
}

func ReboxStruct[Src any, Dest any](srcVal *Src) *Dest {
	if srcVal == nil {
		return nil
	}

	src := reflect.ValueOf(srcVal).Elem()
	destPtr := reflect.New(reflect.TypeOf(*new(Dest)))
	dest := destPtr.Elem()

	for i := 0; i < src.NumField(); i++ {
		srcField := src.Field(i)
		destField := dest.Field(i)
		destField.Set(copyReflectField(srcField, destField.Type()))
	}

	return destPtr.Interface().(*Dest)
}

func copyReflectField(src reflect.Value, destType reflect.Type) reflect.Value {
	if src.Type() == destType {
		return src
	}

	valueType := reflect.TypeOf(Value{})

	if src.Type() == valueType {
		val := src.Interface().(Value)

		switch destType.Kind() {
		case reflect.Int64:
			if val.Type == TypeFloat {
				return reflect.ValueOf(int64(math.Float64frombits(uint64(val.IntVal))))
			}
			return reflect.ValueOf(val.IntVal)
		case reflect.Int:
			if val.Type == TypeFloat {
				return reflect.ValueOf(int(math.Float64frombits(uint64(val.IntVal))))
			}
			return reflect.ValueOf(int(val.IntVal))
		case reflect.String:
			if val.UnsafePtr != nil {
				return reflect.ValueOf(StrValue(val))
			}
			return reflect.ValueOf("")
		case reflect.Bool:
			return reflect.ValueOf(val.IntVal != 0)
		case reflect.Float64:
			return reflect.ValueOf(math.Float64frombits(uint64(val.IntVal)))
		}
	}

	if destType == valueType {
		return reflect.ValueOf(AnyToValue(src.Interface()))
	}

	if src.Kind() == reflect.Ptr && destType.Kind() == reflect.Ptr {
		if src.IsNil() {
			return reflect.Zero(destType)
		}
		srcElem := src.Elem()
		destElemPtr := reflect.New(destType.Elem())
		destElem := destElemPtr.Elem()

		for i := 0; i < srcElem.NumField(); i++ {
			destElem.Field(i).Set(copyReflectField(srcElem.Field(i), destElem.Field(i).Type()))
		}
		return destElemPtr
	}

	if src.Type().ConvertibleTo(destType) {
		return src.Convert(destType)
	}
	return reflect.Zero(destType)
}

// Undefined has no Go scalar equivalent. This opaque native marker keeps it
// distinct from null and an absent property until an explicit serializer runs.
type foreignUndefined struct{}

// IsForeignUndefined lets specialized Foreign serializers recognize the native
// marker without exposing the boxed PureScript representation to ordinary FFI.
func IsForeignUndefined(value any) bool {
	_, undefined := value.(foreignUndefined)
	return undefined
}

// ForeignValue crosses a dynamic-data boundary, not a serialization boundary.
// Native data keeps its types and identity. Legacy native containers that still
// contain boxed Values are copied only along paths requiring conversion. The
// graph walk preserves cycles and shared children without inspecting opaque
// structs or changing application-level serialization policy.
func ForeignValue(value Value) any {
	graph := foreignDataGraph{}
	graph.scan(value, nil)
	for i := 0; i < len(graph.changed); i++ {
		for _, parent := range graph.changed[i].parents {
			graph.mark(parent)
		}
	}
	return graph.convert(value)
}

type foreignDataKey struct {
	kind    int
	pointer unsafe.Pointer
	length  int
}

type foreignDataNode struct {
	source    any
	parents   []*foreignDataNode
	changed   bool
	converted bool
	result    any
}

type foreignDataGraph struct {
	nodes   map[foreignDataKey]*foreignDataNode
	changed []*foreignDataNode
}

func foreignDataContainerKey(input any) (foreignDataKey, bool) {
	switch value := input.(type) {
	case map[string]any:
		return foreignDataKey{kind: 1, pointer: reflect.ValueOf(value).UnsafePointer()}, value != nil
	case []any:
		return foreignDataKey{kind: 2, pointer: unsafe.Pointer(unsafe.SliceData(value)), length: len(value)}, value != nil
	case map[string]Value:
		return foreignDataKey{kind: 3, pointer: reflect.ValueOf(value).UnsafePointer()}, true
	case []Value:
		return foreignDataKey{kind: 4, pointer: unsafe.Pointer(unsafe.SliceData(value)), length: len(value)}, true
	case Value:
		switch value.Type {
		case TypeArray:
			var values []Value
			if value.UnsafePtr != nil {
				values = *(*[]Value)(value.UnsafePtr)
			}
			return foreignDataContainerKey(values)
		case TypeIntArray:
			var ints []int64
			if value.UnsafePtr != nil {
				ints = *(*[]int64)(value.UnsafePtr)
			}
			return foreignDataKey{kind: 5, pointer: unsafe.Pointer(unsafe.SliceData(ints)), length: len(ints)}, true
		case TypeInt32Array:
			var ints []int32
			if value.UnsafePtr != nil {
				ints = *(*[]int32)(value.UnsafePtr)
			}
			return foreignDataKey{kind: 6, pointer: unsafe.Pointer(unsafe.SliceData(ints)), length: len(ints)}, true
		case TypeStringArray:
			var strs []string
			if value.UnsafePtr != nil {
				strs = *(*[]string)(value.UnsafePtr)
			}
			return foreignDataKey{kind: 7, pointer: unsafe.Pointer(unsafe.SliceData(strs)), length: len(strs)}, true
		case TypeRecord, TypeRecord1, TypeRecord2, TypeRecord3, TypeRecord4, TypeRecord5, TypeRecordData:
			return foreignDataKey{kind: value.Type + 100, pointer: value.UnsafePtr}, true
		}
	}
	return foreignDataKey{}, false
}

func (graph *foreignDataGraph) mark(node *foreignDataNode) {
	if node != nil && !node.changed {
		node.changed = true
		graph.changed = append(graph.changed, node)
	}
}

func (graph *foreignDataGraph) scan(input any, parent *foreignDataNode) {
	if value, boxed := input.(Value); boxed {
		graph.mark(parent)
		if value.Type == TypeAny {
			if value.UnsafePtr != nil {
				graph.scan(*(*any)(value.UnsafePtr), parent)
			}
			return
		}
	}
	key, container := foreignDataContainerKey(input)
	if !container {
		return
	}
	if node := graph.nodes[key]; node != nil {
		if parent != nil {
			node.parents = append(node.parents, parent)
		}
		return
	}
	if graph.nodes == nil {
		graph.nodes = make(map[foreignDataKey]*foreignDataNode)
	}
	node := &foreignDataNode{source: input}
	graph.nodes[key] = node
	if parent != nil {
		node.parents = append(node.parents, parent)
	}
	if value, boxed := input.(Value); boxed {
		graph.mark(node)
		if value.Type == TypeArray {
			var values []Value
			if value.UnsafePtr != nil {
				values = *(*[]Value)(value.UnsafePtr)
			}
			node.source = values
		} else if value.Type == TypeIntArray || value.Type == TypeInt32Array || value.Type == TypeStringArray {
			node.source = ArrayValues(value)
		} else {
			node.source = RecordToMap(value)
		}
	}
	switch source := node.source.(type) {
	case map[string]any:
		for _, child := range source {
			graph.scan(child, node)
		}
	case []any:
		for _, child := range source {
			graph.scan(child, node)
		}
	case map[string]Value:
		graph.mark(node)
		for _, child := range source {
			graph.scan(child, node)
		}
	case []Value:
		graph.mark(node)
		for _, child := range source {
			graph.scan(child, node)
		}
	}
}

func (graph *foreignDataGraph) convert(input any) any {
	if value, boxed := input.(Value); boxed {
		switch value.Type {
		case 0:
			return foreignUndefined{}
		case TypeInt:
			return value.IntVal
		case TypeFloat:
			return value.FloatVal()
		case TypeString:
			return StrValue(value)
		case TypeBool:
			return value.IntVal != 0
		case TypeAny:
			if value.UnsafePtr == nil {
				return nil
			}
			return graph.convert(*(*any)(value.UnsafePtr))
		case TypeRecord0:
			return map[string]any{}
		case TypeArray, TypeIntArray, TypeInt32Array, TypeStringArray, TypeRecord, TypeRecord1, TypeRecord2, TypeRecord3, TypeRecord4, TypeRecord5, TypeRecordData:
			// Containers use the graph below, including their cycles.
		default:
			panic("FFI Foreign: functions and algebraic constructors require an explicit PureScript conversion to native data")
		}
	}
	key, container := foreignDataContainerKey(input)
	if !container {
		return input
	}
	node := graph.nodes[key]
	if !node.changed {
		return input
	}
	if node.converted {
		return node.result
	}
	node.converted = true
	// Publish the destination before filling it, so a back-edge refers to the
	// converted container instead of retaining an edge into the old graph.
	switch source := node.source.(type) {
	case map[string]any:
		out := make(map[string]any, len(source))
		node.result = out
		for key, child := range source {
			out[key] = graph.convert(child)
		}
	case []any:
		out := make([]any, len(source))
		node.result = out
		for i, child := range source {
			out[i] = graph.convert(child)
		}
	case map[string]Value:
		var out map[string]any
		if source != nil {
			out = make(map[string]any, len(source))
		}
		node.result = out
		for key, child := range source {
			out[key] = graph.convert(child)
		}
	case []Value:
		var out []any
		if source != nil {
			out = make([]any, len(source))
		}
		node.result = out
		for i, child := range source {
			out[i] = graph.convert(child)
		}
	}
	return node.result
}
