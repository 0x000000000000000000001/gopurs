package purescript

import "gopurs/output/gopurs_runtime"

import (
	"encoding/json"
	"strconv"
	"strings"
	"unicode/utf16"
	"unicode/utf8"
)

func Data_Argonaut_Parser__JsonParser(fail func(any) any, succ func(any) any, str string) any {
	result, err := argonautParseJSON(str)
	if err != nil {
		return fail(err.Error())
	}
	return succ(result)
}

type argonautJSONParser struct {
	text string
	pos  int
}

type argonautJSONEntry struct {
	key   string
	value any
}

// Small objects avoid a hash table's minimum bucket allocation. Inline fields
// also avoid a second allocation for an interface-held slice header.
// Each instantiation has exactly the required capacity, rather than retaining
// an eight-field buffer for every two-field object. Larger objects use maps.
// These methods implement gopurs_runtime.JSONObject without coupling the parser
// itself to the boxed-value runtime.
type argonautJSONFields interface {
	[0]argonautJSONEntry | [1]argonautJSONEntry | [2]argonautJSONEntry | [3]argonautJSONEntry | [4]argonautJSONEntry |
		[5]argonautJSONEntry | [6]argonautJSONEntry | [7]argonautJSONEntry | [8]argonautJSONEntry
}
type argonautJSONSmallObject[T argonautJSONFields] struct{ fields T }

func (object *argonautJSONSmallObject[T]) JSONLength() int { return len(object.fields) }
func (object *argonautJSONSmallObject[T]) JSONEntry(index int) (string, any) {
	entry := object.fields[index]
	return entry.key, entry.value
}
func (object *argonautJSONSmallObject[T]) JSONLookup(key string) (any, bool) {
	for i := 0; i < len(object.fields); i++ {
		entry := object.fields[i]
		if entry.key == key {
			return entry.value, true
		}
	}
	return nil, false
}
func (object *argonautJSONSmallObject[T]) MarshalJSON() ([]byte, error) {
	fields := make(map[string]any, len(object.fields))
	for i := 0; i < len(object.fields); i++ {
		entry := object.fields[i]
		fields[entry.key] = entry.value
	}
	return json.Marshal(fields)
}

func argonautSmallObject(fields []argonautJSONEntry) any {
	switch len(fields) {
	case 0:
		return &argonautJSONSmallObject[[0]argonautJSONEntry]{}
	case 1:
		return &argonautJSONSmallObject[[1]argonautJSONEntry]{[1]argonautJSONEntry{fields[0]}}
	case 2:
		return &argonautJSONSmallObject[[2]argonautJSONEntry]{[2]argonautJSONEntry{fields[0], fields[1]}}
	case 3:
		return &argonautJSONSmallObject[[3]argonautJSONEntry]{[3]argonautJSONEntry{fields[0], fields[1], fields[2]}}
	case 4:
		return &argonautJSONSmallObject[[4]argonautJSONEntry]{[4]argonautJSONEntry{fields[0], fields[1], fields[2], fields[3]}}
	case 5:
		return &argonautJSONSmallObject[[5]argonautJSONEntry]{[5]argonautJSONEntry{fields[0], fields[1], fields[2], fields[3], fields[4]}}
	case 6:
		return &argonautJSONSmallObject[[6]argonautJSONEntry]{[6]argonautJSONEntry{fields[0], fields[1], fields[2], fields[3], fields[4], fields[5]}}
	case 7:
		return &argonautJSONSmallObject[[7]argonautJSONEntry]{[7]argonautJSONEntry{fields[0], fields[1], fields[2], fields[3], fields[4], fields[5], fields[6]}}
	case 8:
		return &argonautJSONSmallObject[[8]argonautJSONEntry]{[8]argonautJSONEntry{fields[0], fields[1], fields[2], fields[3], fields[4], fields[5], fields[6], fields[7]}}
	default:
		panic("argonautSmallObject: more than eight fields")
	}
}

// Parse the existing Json representation in one pass over the input string.
// Keep encoding/json as the authority for malformed inputs and error messages.
// Extracted strings own their storage, so a small result does not retain the
// complete source document. Unicode replacement follows encoding/json's rules.
func argonautParseJSON(text string) (any, error) {
	p := argonautJSONParser{text: text}
	value, ok := p.value(0)
	p.space()
	if ok && p.pos == len(text) {
		return value, nil
	}
	var fallback any
	err := json.Unmarshal([]byte(text), &fallback)
	return fallback, err
}

func (p *argonautJSONParser) space() {
	for p.pos < len(p.text) {
		switch p.text[p.pos] {
		case ' ', '\n', '\r', '\t':
			p.pos++
		default:
			return
		}
	}
}

func (p *argonautJSONParser) value(depth int) (any, bool) {
	p.space()
	if p.pos >= len(p.text) || depth > 10000 {
		return nil, false
	}
	switch p.text[p.pos] {
	case '"':
		return p.string()
	case '{':
		return p.object(depth + 1)
	case '[':
		return p.array(depth + 1)
	case 't':
		if strings.HasPrefix(p.text[p.pos:], "true") {
			p.pos += 4
			return true, true
		}
	case 'f':
		if strings.HasPrefix(p.text[p.pos:], "false") {
			p.pos += 5
			return false, true
		}
	case 'n':
		if strings.HasPrefix(p.text[p.pos:], "null") {
			p.pos += 4
			return nil, true
		}
	default:
		return p.number()
	}
	return nil, false
}

func (p *argonautJSONParser) string() (string, bool) {
	if p.pos >= len(p.text) || p.text[p.pos] != '"' {
		return "", false
	}
	start := p.pos
	p.pos++
	escaped, nonASCII := false, false
	for p.pos < len(p.text) {
		c := p.text[p.pos]
		p.pos++
		if c == '"' {
			if escaped || (nonASCII && !utf8.ValidString(p.text[start+1:p.pos-1])) {
				return argonautUnquoteJSON(p.text[start+1 : p.pos-1])
			}
			return strings.Clone(p.text[start+1 : p.pos-1]), true
		}
		if c < 0x20 {
			return "", false
		}
		if c >= utf8.RuneSelf {
			nonASCII = true
		}
		if c == '\\' {
			escaped = true
			if p.pos >= len(p.text) {
				return "", false
			}
			p.pos++
		}
	}
	return "", false
}

// Decode JSON escapes without starting a complete decoder for each string.
// The builder owns the output; invalid UTF-8 and unpaired UTF-16 surrogates
// become U+FFFD, as they do in encoding/json. Invalid JSON uses the full fallback.
func argonautUnquoteJSON(text string) (string, bool) {
	var output strings.Builder
	output.Grow(len(text))
	for i := 0; i < len(text); {
		start := i
		for i < len(text) && text[i] >= 0x20 && text[i] < utf8.RuneSelf && text[i] != '\\' {
			i++
		}
		output.WriteString(text[start:i])
		if i == len(text) {
			break
		}
		c := text[i]
		if c < 0x20 {
			return "", false
		}
		if c >= utf8.RuneSelf {
			r, size := utf8.DecodeRuneInString(text[i:])
			output.WriteRune(r)
			i += size
			continue
		}
		i++ // backslash
		if i == len(text) {
			return "", false
		}
		escape := text[i]
		i++
		switch escape {
		case '"', '\\', '/':
			output.WriteByte(escape)
		case 'b':
			output.WriteByte('\b')
		case 'f':
			output.WriteByte('\f')
		case 'n':
			output.WriteByte('\n')
		case 'r':
			output.WriteByte('\r')
		case 't':
			output.WriteByte('\t')
		case 'u':
			r, ok := argonautJSONHex4(text, i)
			if !ok {
				return "", false
			}
			i += 4
			if r >= 0xd800 && r <= 0xdbff {
				if i+6 <= len(text) && text[i] == '\\' && text[i+1] == 'u' {
					low, valid := argonautJSONHex4(text, i+2)
					if valid && low >= 0xdc00 && low <= 0xdfff {
						r = utf16.DecodeRune(r, low)
						i += 6
					} else {
						r = utf8.RuneError
					}
				} else {
					r = utf8.RuneError
				}
			} else if r >= 0xdc00 && r <= 0xdfff {
				r = utf8.RuneError
			}
			output.WriteRune(r)
		default:
			return "", false
		}
	}
	return output.String(), true
}

func argonautJSONHex4(text string, pos int) (rune, bool) {
	if len(text)-pos < 4 {
		return 0, false
	}
	var value rune
	for _, c := range []byte(text[pos : pos+4]) {
		value <<= 4
		switch {
		case c >= '0' && c <= '9':
			value |= rune(c - '0')
		case c >= 'a' && c <= 'f':
			value |= rune(c-'a') + 10
		case c >= 'A' && c <= 'F':
			value |= rune(c-'A') + 10
		default:
			return 0, false
		}
	}
	return value, true
}

func (p *argonautJSONParser) number() (any, bool) {
	start := p.pos
	if p.pos < len(p.text) && p.text[p.pos] == '-' {
		p.pos++
	}
	if p.pos >= len(p.text) {
		return nil, false
	}
	if p.text[p.pos] == '0' {
		p.pos++
	} else {
		if p.text[p.pos] < '1' || p.text[p.pos] > '9' {
			return nil, false
		}
		for p.pos < len(p.text) && p.text[p.pos] >= '0' && p.text[p.pos] <= '9' {
			p.pos++
		}
	}
	if p.pos < len(p.text) && p.text[p.pos] == '.' {
		p.pos++
		digits := p.pos
		for p.pos < len(p.text) && p.text[p.pos] >= '0' && p.text[p.pos] <= '9' {
			p.pos++
		}
		if p.pos == digits {
			return nil, false
		}
	}
	if p.pos < len(p.text) && (p.text[p.pos] == 'e' || p.text[p.pos] == 'E') {
		p.pos++
		if p.pos < len(p.text) && (p.text[p.pos] == '+' || p.text[p.pos] == '-') {
			p.pos++
		}
		digits := p.pos
		for p.pos < len(p.text) && p.text[p.pos] >= '0' && p.text[p.pos] <= '9' {
			p.pos++
		}
		if p.pos == digits {
			return nil, false
		}
	}
	n, err := strconv.ParseFloat(p.text[start:p.pos], 64)
	return n, err == nil
}

func (p *argonautJSONParser) object(depth int) (any, bool) {
	if depth > 10000 {
		return nil, false
	}
	p.pos++
	p.space()
	var small [8]argonautJSONEntry
	count := 0
	var obj map[string]any
	if p.pos < len(p.text) && p.text[p.pos] == '}' {
		p.pos++
		return argonautSmallObject(nil), true
	}
	for {
		key, ok := p.string()
		if !ok {
			return nil, false
		}
		p.space()
		if p.pos >= len(p.text) || p.text[p.pos] != ':' {
			return nil, false
		}
		p.pos++
		value, ok := p.value(depth)
		if !ok {
			return nil, false
		}
		if obj != nil {
			obj[key] = value
		} else {
			slot := 0
			for slot < count && small[slot].key != key {
				slot++
			}
			if slot < count {
				small[slot].value = value
			} else if count < len(small) {
				small[count] = argonautJSONEntry{key, value}
				count++
			} else {
				obj = make(map[string]any, count+1)
				for _, entry := range small {
					obj[entry.key] = entry.value
				}
				obj[key] = value
			}
		}
		p.space()
		if p.pos >= len(p.text) {
			return nil, false
		}
		c := p.text[p.pos]
		p.pos++
		if c == '}' {
			if obj == nil {
				return argonautSmallObject(small[:count]), true
			}
			return obj, true
		}
		if c != ',' {
			return nil, false
		}
		p.space()
	}
}

func (p *argonautJSONParser) array(depth int) (any, bool) {
	if depth > 10000 {
		return nil, false
	}
	p.pos++
	p.space()
	var small [8]any
	count := 0
	var values []any
	if p.pos < len(p.text) && p.text[p.pos] == ']' {
		p.pos++
		return []any{}, true
	}
	for {
		value, ok := p.value(depth)
		if !ok {
			return nil, false
		}
		if values != nil {
			values = append(values, value)
		} else if count < len(small) {
			small[count] = value
			count++
		} else {
			values = make([]any, count, count*2)
			copy(values, small[:])
			values = append(values, value)
		}
		p.space()
		if p.pos >= len(p.text) {
			return nil, false
		}
		c := p.text[p.pos]
		p.pos++
		if c == ']' {
			if values == nil {
				values = make([]any, count)
				copy(values, small[:count])
			}
			return values, true
		}
		if c != ',' {
			return nil, false
		}
	}
}

// --- Auto-generated FFI wrappers ---
var _Gopurs_Data_Argonaut_Parser__JsonParser = // TAST: (ForAll [a] (ADT ["Data","Function","Uncurried","Fn3"] [(Func [String] (TypeVar a)), (Func [(ADT ["Data","Argonaut","Core","Json"] [])] (TypeVar a)), String, (TypeVar a)]))
gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any) any {
		return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
	}
	go_arg1 := func(p0_0 any) any {
		return gopurs_runtime.Apply(arg1, gopurs_runtime.Box(p0_0))
	}
	go_arg2 := gopurs_runtime.Unbox[string](arg2)
	go_res := Data_Argonaut_Parser__JsonParser(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})
