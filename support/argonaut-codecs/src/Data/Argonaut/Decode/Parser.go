package Parser

import (
	"strconv"
	"strings"
	"unicode/utf8"

	"gopurs/output/gopurs_runtime"
)

var TextDecoderABI1 = 1

type argonautTextWorker func(*recordDecodePlan, *fieldKind, argonautTextCursor) argonautSchemaResult

// The ordinary worker and its dictionary checks are retained. The stricter
// text guard requires every method to match a proven operation; opaque callback
// trees use the complete ordinary composition, preserving Json argument aliases.
func argonautCompileTextSchema(original gopurs_runtime.Value, worker argonautTextWorker, accepts func(*fieldKind) bool) gopurs_runtime.Value {
	tag, ok := gopurs_runtime.FunctionData[*typedKind](original)
	if !ok || tag == nil || tag.worker == nil || !accepts(tag.kind) {
		return original
	}
	compiled := *tag
	compiled.textWorker = worker
	return gopurs_runtime.WithFunctionData(original, &compiled)
}

func DecodeJsonStringImpl(_ int64, fallback, decoder gopurs_runtime.Value, text string) gopurs_runtime.Value {
	tag, ok := gopurs_runtime.FunctionData[*typedKind](decoder)
	if !ok || tag == nil {
		return gopurs_runtime.Apply(fallback, gopurs_runtime.Str(text))
	}
	worker, ok := tag.textWorker.(argonautTextWorker)
	if !ok || worker == nil {
		return gopurs_runtime.Apply(fallback, gopurs_runtime.Str(text))
	}
	input, ok := argonautTextIndex(text)
	if !ok {
		return gopurs_runtime.Apply(fallback, gopurs_runtime.Str(text))
	}
	result := worker(tag.errors, tag.kind, input)
	if !result.ok {
		return gopurs_runtime.Apply(tag.errors.left, result.err)
	}
	return gopurs_runtime.Apply(tag.errors.right, result.value)
}

// Validate all syntax, including ignored fields, before invoking any decoder.
// This call-local index contains offsets only. No cursor is published as Json;
// all strings entering a final value own their storage. The scanner follows
// the validated single-pass TAST implementation's grammar and numeric rules.
type argonautTextToken struct{ start, end, next uint32 }
type argonautTextDocument struct {
	text   string
	tokens []argonautTextToken
}
type argonautTextCursor struct {
	document *argonautTextDocument
	index    uint32
}
type argonautTextParser struct {
	document argonautTextDocument
	pos      int
}

const argonautTextNormalize uint32 = 1 << 31

func argonautTextIndex(text string) (argonautTextCursor, bool) {
	if uint64(len(text)) >= uint64(argonautTextNormalize) {
		return argonautTextCursor{}, false
	}
	p := argonautTextParser{document: argonautTextDocument{text: text, tokens: make([]argonautTextToken, 0, len(text)/6+8)}}
	if !p.value(0) {
		return argonautTextCursor{}, false
	}
	p.space()
	if p.pos != len(text) {
		return argonautTextCursor{}, false
	}
	return argonautTextCursor{&p.document, 0}, true
}

func (p *argonautTextParser) space() {
	for p.pos < len(p.document.text) {
		switch p.document.text[p.pos] {
		case ' ', '\t', '\n', '\r':
			p.pos++
		default:
			return
		}
	}
}

func (p *argonautTextParser) take(c byte) bool {
	if p.pos < len(p.document.text) && p.document.text[p.pos] == c {
		p.pos++
		return true
	}
	return false
}

func (p *argonautTextParser) value(depth int) bool {
	p.space()
	text := p.document.text
	if p.pos == len(text) {
		return false
	}
	index, start := len(p.document.tokens), p.pos
	switch text[start] {
	case '{', '[':
		if depth >= 10000 {
			return false
		}
		object := text[start] == '{'
		end := byte(']')
		if object {
			end = '}'
		}
		p.pos++
		p.document.tokens = append(p.document.tokens, argonautTextToken{start: uint32(start)})
		p.space()
		var count, lastKey uint32
		if !p.take(end) {
			for {
				if object {
					key := uint32(len(p.document.tokens))
					if !p.string() {
						return false
					}
					p.document.tokens[key].next = lastKey
					lastKey = key
					p.space()
					if !p.take(':') {
						return false
					}
				}
				if !p.value(depth + 1) {
					return false
				}
				count++
				p.space()
				if p.take(end) {
					break
				}
				if !p.take(',') {
					return false
				}
				p.space()
			}
		}
		if object {
			count = lastKey
		}
		p.document.tokens[index].end = count
		p.document.tokens[index].next = uint32(len(p.document.tokens))
		return true
	case '"':
		return p.string()
	case 't':
		if !strings.HasPrefix(text[start:], "true") {
			return false
		}
		p.pos += 4
	case 'f':
		if !strings.HasPrefix(text[start:], "false") {
			return false
		}
		p.pos += 5
	case 'n':
		if !strings.HasPrefix(text[start:], "null") {
			return false
		}
		p.pos += 4
	default:
		if !p.number() {
			return false
		}
	}
	p.document.tokens = append(p.document.tokens, argonautTextToken{uint32(start), uint32(p.pos), uint32(index + 1)})
	return true
}

func (p *argonautTextParser) string() bool {
	text, start := p.document.text, p.pos
	if !p.take('"') {
		return false
	}
	escaped, nonASCII := false, false
	for p.pos < len(text) {
		c := text[p.pos]
		p.pos++
		if c == '"' {
			end := uint32(p.pos)
			if escaped || (nonASCII && !utf8.ValidString(text[start+1:p.pos-1])) {
				end |= argonautTextNormalize
			}
			p.document.tokens = append(p.document.tokens, argonautTextToken{uint32(start), end, uint32(len(p.document.tokens) + 1)})
			return true
		}
		if c < 0x20 {
			return false
		}
		if c >= utf8.RuneSelf {
			nonASCII = true
		}
		if c != '\\' {
			continue
		}
		escaped = true
		if p.pos == len(text) {
			return false
		}
		escape := text[p.pos]
		p.pos++
		switch escape {
		case '"', '\\', '/', 'b', 'f', 'n', 'r', 't':
		case 'u':
			if _, ok := argonautJSONHex4(text, p.pos); !ok {
				return false
			}
			p.pos += 4
		default:
			return false
		}
	}
	return false
}

func (p *argonautTextParser) number() bool {
	text, start := p.document.text, p.pos
	p.take('-')
	if p.pos == len(text) {
		return false
	}
	if !p.take('0') {
		if text[p.pos] < '1' || text[p.pos] > '9' {
			return false
		}
		for p.pos < len(text) && text[p.pos] >= '0' && text[p.pos] <= '9' {
			p.pos++
		}
	}
	integer := true
	if p.take('.') {
		integer = false
		digits := p.pos
		for p.pos < len(text) && text[p.pos] >= '0' && text[p.pos] <= '9' {
			p.pos++
		}
		if digits == p.pos {
			return false
		}
	}
	if p.pos < len(text) && (text[p.pos] == 'e' || text[p.pos] == 'E') {
		integer = false
		p.pos++
		if !p.take('+') {
			p.take('-')
		}
		digits := p.pos
		for p.pos < len(text) && text[p.pos] >= '0' && text[p.pos] <= '9' {
			p.pos++
		}
		if digits == p.pos {
			return false
		}
	}
	if !integer || p.pos-start > 15 {
		if _, err := strconv.ParseFloat(text[start:p.pos], 64); err != nil {
			return false
		}
	}
	return true
}

func (c argonautTextCursor) kind() byte {
	if c.document == nil {
		return 0
	}
	return c.document.text[c.document.tokens[c.index].start]
}

func (c argonautTextCursor) string(owned bool) string {
	token := c.document.tokens[c.index]
	text := c.document.text[token.start+1 : (token.end&^argonautTextNormalize)-1]
	if token.end&argonautTextNormalize != 0 {
		value, ok := argonautUnquoteJSON(text)
		if !ok {
			panic("validated JSON string")
		}
		return value
	}
	if owned {
		return strings.Clone(text)
	}
	return text
}

func (c argonautTextCursor) Lookup(key string) (argonautTextCursor, bool) {
	if c.kind() == '{' {
		for at := c.document.tokens[c.index].end; at != 0; at = c.document.tokens[at].next {
			if (argonautTextCursor{c.document, at}).string(false) == key {
				return argonautTextCursor{c.document, at + 1}, true
			}
		}
	}
	return argonautTextCursor{}, false
}

// Escaping Json values use the ordinary owned parser's representation,
// including object insertion order and duplicate-key handling.
func (c argonautTextCursor) materialize() any {
	p := argonautJSONParser{text: c.document.text, pos: int(c.document.tokens[c.index].start)}
	value, ok := p.value(0)
	if !ok {
		panic("validated JSON subtree")
	}
	return value
}

func argonautTextObject(c argonautTextCursor) (argonautTextCursor, bool) { return c, c.kind() == '{' }
func argonautTextNull(c argonautTextCursor) bool                         { return c.kind() == 'n' }
func argonautTextBool(c argonautTextCursor) (bool, bool) {
	return c.kind() == 't', c.kind() == 't' || c.kind() == 'f'
}
func argonautTextString(c argonautTextCursor) (string, bool) {
	if c.kind() != '"' {
		return "", false
	}
	return c.string(true), true
}

func argonautTextNumber(c argonautTextCursor) (float64, bool) {
	k := c.kind()
	if k != '-' && (k < '0' || k > '9') {
		return 0, false
	}
	token := c.document.tokens[c.index]
	text := c.document.text[token.start:token.end]
	// Exactly representable short integers need no floating-point conversion
	// parser. Keep the sign of zero and use ParseFloat for all other tokens.
	if len(text) <= 15 {
		at, negative := 0, text[0] == '-'
		if negative {
			at++
		}
		var integer int64
		for at < len(text) && text[at] >= '0' && text[at] <= '9' {
			integer = integer*10 + int64(text[at]-'0')
			at++
		}
		if at == len(text) {
			value := float64(integer)
			if negative {
				value = -value
			}
			return value, true
		}
	}
	value, err := strconv.ParseFloat(text, 64)
	return value, err == nil
}

type argonautTextArray struct{ count int }

func (a argonautTextArray) length() int { return a.count }
func argonautTextArrayValues(c argonautTextCursor) (argonautTextArray, bool) {
	if c.kind() != '[' {
		return argonautTextArray{}, false
	}
	return argonautTextArray{int(c.document.tokens[c.index].end)}, true
}
