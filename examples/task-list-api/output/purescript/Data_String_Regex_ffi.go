package purescript

import "gopurs/output/gopurs_runtime"

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode/utf16"
	"unicode/utf8"

	"github.com/dlclark/regexp2"
)

// regexp2 v1.11.5 supports Go 1.13 and preserves gopurs' Go 1.22 minimum.
// ECMAScript mode is a compatibility subset, not a complete JS RegExp engine.
// See docs/go-regex.md for the qualified cases and explicit limitations.
type GoRegex struct {
	Re     *regexp2.Regexp
	Global bool
	Flags  string
	Source string
	names  map[string]int
	cst    *regexp.Regexp
}

// Matching runs on UTF-16 units without u, and Unicode scalars with u. Keep
// isolated surrogates in canonical WTF-8, as in the native CodeUnits module.
func regexRunes(s string, unicode bool) []rune {
	out := make([]rune, 0, len(s))
	for len(s) > 0 {
		var r rune
		var size int
		if len(s) >= 3 && s[0] == 0xed && s[1] >= 0xa0 && s[1] <= 0xbf && s[2]&0xc0 == 0x80 {
			r, size = rune(s[0]&15)<<12|rune(s[1]&63)<<6|rune(s[2]&63), 3
		} else {
			r, size = utf8.DecodeRuneInString(s)
		}
		if !unicode && r > 0xffff {
			a, b := utf16.EncodeRune(r)
			out = append(out, a, b)
		} else {
			out = append(out, r)
		}
		s = s[size:]
	}
	return out
}

func regexString(rs []rune) string {
	var b strings.Builder
	for i := 0; i < len(rs); i++ {
		r := rs[i]
		if r >= 0xd800 && r <= 0xdbff && i+1 < len(rs) && rs[i+1] >= 0xdc00 && rs[i+1] <= 0xdfff {
			r = utf16.DecodeRune(r, rs[i+1])
			i++
		}
		if r >= 0xd800 && r <= 0xdfff {
			b.Write([]byte{0xed, 0x80 | byte(r>>6&63), 0x80 | byte(r&63)})
		} else {
			b.WriteRune(r)
		}
	}
	return b.String()
}

// Name every capture by its JS ordinal, including named captures. This avoids
// .NET's "unnamed groups first" numbering while retaining named references.
func regexPattern(source string, unicode bool, multiline bool, dotAll bool) (string, map[string]int, error) {
	names := make(map[string]int)
	var b strings.Builder
	inClass, ordinal := false, 0
	rs := regexRunes(source, unicode)
	for i := 0; i < len(rs); i++ {
		c := rs[i]
		if c == '\\' && i+1 < len(rs) {
			next := rs[i+1]
			if unicode && next == 'u' && i+11 < len(rs) && rs[i+6] == '\\' && rs[i+7] == 'u' {
				hi, e1 := strconv.ParseUint(string(rs[i+2:i+6]), 16, 16)
				lo, e2 := strconv.ParseUint(string(rs[i+8:i+12]), 16, 16)
				if e1 == nil && e2 == nil && hi >= 0xd800 && hi <= 0xdbff && lo >= 0xdc00 && lo <= 0xdfff {
					b.WriteRune(utf16.DecodeRune(rune(hi), rune(lo)))
					i += 11
					continue
				}
			}
			if next >= 'A' && next <= 'Z' || next >= 'a' && next <= 'z' {
				known := strings.ContainsRune("bBcdDfnrsStuvwWxk", next) || unicode && (next == 'p' || next == 'P')
				if !known {
					if unicode {
						return "", nil, fmt.Errorf("invalid Unicode identity escape \\%c", next)
					}
					b.WriteRune(next)
					i++
					continue
				}
			}
			b.WriteRune(c)
			i++
			b.WriteRune(rs[i])
			continue
		}
		if c == '[' && !inClass {
			inClass = true
		} else if c == ']' && inClass {
			inClass = false
		}
		if !inClass {
			switch c {
			case '.':
				if !dotAll {
					b.WriteString(`[^\r\n\u2028\u2029]`)
					continue
				}
			case '^':
				if multiline {
					b.WriteString(`(?:\A|(?<=[\r\n\u2028\u2029]))`)
					continue
				}
			case '$':
				if multiline {
					b.WriteString(`(?:\z|(?=[\r\n\u2028\u2029]))`)
					continue
				}
			}
		}
		if c == '(' && !inClass {
			if i+1 < len(rs) && rs[i+1] == '?' {
				if i+2 >= len(rs) {
					return "", nil, fmt.Errorf("incomplete group")
				}
				switch rs[i+2] {
				case ':', '=', '!':
				case '<':
					if i+3 < len(rs) && (rs[i+3] == '=' || rs[i+3] == '!') {
						break
					}
					end := i + 3
					for end < len(rs) && rs[end] != '>' {
						end++
					}
					if end == len(rs) {
						return "", nil, fmt.Errorf("unterminated capture name")
					}
					name := string(rs[i+3 : end])
					if name == "" {
						return "", nil, fmt.Errorf("empty capture name")
					}
					for n, r := range name {
						if !(r == '_' || r == '$' || r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' || n > 0 && r >= '0' && r <= '9') {
							return "", nil, fmt.Errorf("unsupported capture name %q", name)
						}
					}
					if _, exists := names[name]; exists {
						return "", nil, fmt.Errorf("duplicate capture name %q", name)
					}
					ordinal++
					names[name] = ordinal
					fmt.Fprintf(&b, "(?<%d>", ordinal)
					i = end
					continue
				default:
					return "", nil, fmt.Errorf("unsupported group construct")
				}
			} else {
				ordinal++
				fmt.Fprintf(&b, "(?<%d>", ordinal)
				continue
			}
		}
		if c >= 0xd800 && c <= 0xdfff {
			fmt.Fprintf(&b, `\u%04x`, c)
		} else {
			b.WriteRune(c)
		}
	}
	pattern := b.String()
	// Resolve named backreferences only outside escaped text and character classes.
	b.Reset()
	inClass = false
	for i := 0; i < len(pattern); i++ {
		c := pattern[i]
		if c == '\\' && i+1 < len(pattern) {
			if !inClass && strings.HasPrefix(pattern[i:], `\k<`) {
				end := strings.IndexByte(pattern[i+3:], '>')
				if end < 0 {
					return "", nil, fmt.Errorf("unterminated named reference")
				}
				name := pattern[i+3 : i+3+end]
				n, exists := names[name]
				if !exists {
					return "", nil, fmt.Errorf("unknown capture name %q", name)
				}
				fmt.Fprintf(&b, `\%d`, n)
				i += 3 + end
				continue
			}
			b.WriteByte(c)
			i++
			b.WriteByte(pattern[i])
			continue
		}
		if c == '[' {
			inClass = true
		} else if c == ']' {
			inClass = false
		}
		b.WriteByte(c)
	}
	return b.String(), names, nil
}

// regexp2 retains captures from earlier repetitions where JS clears them.
// Reject repeating groups containing nested captures rather than silently
// returning stale captures or using them in a later backreference.
func regexCheckRepeatingCaptures(source string) error {
	stack := []int{}
	captures := 0
	inClass := false
	for i := 0; i < len(source); i++ {
		c := source[i]
		if c == '\\' {
			i++
			continue
		}
		if c == '[' {
			inClass = true
		}
		if c == ']' {
			inClass = false
			continue
		}
		if inClass {
			continue
		}
		if c == '(' {
			capturing := i+1 == len(source) || source[i+1] != '?'
			if strings.HasPrefix(source[i:], "(?<") && i+3 < len(source) && source[i+3] != '=' && source[i+3] != '!' {
				capturing = true
			}
			if capturing {
				captures++
			}
			stack = append(stack, captures)
		}
		if c == ')' && len(stack) > 0 {
			start := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if captures > start && i+1 < len(source) && strings.ContainsRune("*+{", rune(source[i+1])) {
				return fmt.Errorf("repeated groups with nested captures are not supported by the native regex engine")
			}
		}
	}
	return nil
}

func Data_String_Regex_RegexImpl(left func(string) any, right func(any) any, source string, flags string) any {
	if err := regexCheckRepeatingCaptures(source); err != nil {
		return left(err.Error())
	}
	opts := regexp2.RegexOptions(regexp2.ECMAScript)
	seen := make(map[rune]bool)
	for _, f := range flags {
		if seen[f] {
			return left(fmt.Sprintf("duplicate regex flag %q", f))
		}
		seen[f] = true
		switch f {
		case 'g', 'y':
		case 'i':
			opts |= regexp2.IgnoreCase
		case 'm':
			opts |= regexp2.Multiline
		case 's':
			opts |= regexp2.Singleline
		case 'u':
			opts |= regexp2.Unicode
		default:
			return left(fmt.Sprintf("unsupported regex flag %q", f))
		}
	}
	pattern, names, err := regexPattern(source, seen['u'], seen['m'], seen['s'])
	if err != nil {
		return left(err.Error())
	}
	re, err := regexp2.Compile(pattern, opts)
	if err != nil {
		return left(err.Error())
	}
	// Pure Regex operations have no Either error channel. Surface timeout as an
	// exception, never as false/Nothing/a partial result. The clock is approximate.
	re.MatchTimeout = time.Second
	canonicalFlags := ""
	for _, f := range "gimsuy" {
		if seen[f] {
			canonicalFlags += string(f)
		}
	}
	return right(&GoRegex{Re: re, Global: seen['g'], Flags: canonicalFlags, Source: source, names: names, cst: regexCST(source, canonicalFlags)})
}

// Only these exact, Unicode-only lexer expressions use Go's linear engine.
// Their syntax and captures have a direct translation; the two reductions are
// checked separately. General RE2 compile success is not a compatibility proof.
// This avoids converting the entire remaining source to runes at every token.
func regexCST(source, flags string) *regexp.Regexp {
	if flags != "u" {
		return nil
	}
	pattern := source
	switch source {
	case `^(?:--[^\r\n]*)`, `^(?:#![^\r\n]*)`, `^(?: +)`, "^(?:\n+)", "^(?:(?:\r\n)+)",
		`^(?:(?:(?:\p{Lu}[\p{L}0-9_']*)\.)*)`, `^(?:\p{Lu}[\p{L}0-9_']*)`, `^(?:[\p{Ll}_][\p{L}0-9_']*)`,
		`^(?:[a-fA-F0-9]{1,6})`, `^(?:\\[ \r\n]+\\)`, `^(?:[^"\\]+)`, `^(?:""""{0,2}([^"]+"{1,2})*[^"]*""")`,
		`^(?:(0|[1-9][0-9_]*))`, `^(?:[0-9_]+)`, `^(?:[a-fA-F0-9]+)`:
	case `^(?:(?:[:!#$%&*+./<=>?@\\^|~-]|(?!\p{P})\p{S})+)`:
		// P and S are disjoint; the negative lookahead is redundant.
		pattern = `^(?:(?:[:!#$%&*+./<=>?@\\^|~-]|\p{S})+)`
	case `^(?:\{-(-(?!\})|[^-]+)*(-\}|$))`:
		// Prefer the closing delimiter before repeating. Preserve both captures.
		pattern = `^(?:\{-(-|[^-]+)*?(-\}|$))`
	default:
		return nil
	}
	re, err := regexp.Compile(pattern)
	if err != nil {
		panic(fmt.Errorf("invalid qualified CST expression: %w", err))
	}
	return re
}

func regexFind(r *GoRegex, input []rune, start int, sticky bool) *regexp2.Match {
	m, err := r.Re.FindRunesMatchStartingAt(input, start)
	if err != nil {
		panic(fmt.Errorf("Data.String.Regex: %w", err))
	}
	if sticky && m != nil && m.Index != start {
		return nil
	}
	return m
}

func regexMatches(r *GoRegex, input []rune, all bool) []*regexp2.Match {
	var matches []*regexp2.Match
	sticky := strings.Contains(r.Flags, "y")
	for pos := 0; pos <= len(input); {
		m := regexFind(r, input, pos, sticky)
		if m == nil {
			break
		}
		matches = append(matches, m)
		if !all {
			break
		}
		pos = m.Index + m.Length
		if m.Length == 0 {
			pos++
		}
	}
	return matches
}

func regexGroups(m *regexp2.Match, just func(any) any, nothing any) []any {
	groups := m.Groups()
	out := make([]any, len(groups))
	for i, g := range groups {
		if len(g.Captures) == 0 {
			out[i] = nothing
		} else {
			out[i] = just(regexString(g.Runes()))
		}
	}
	return out
}

func Data_String_Regex__Match(just func(any) any, nothing any, r *GoRegex, s string) any {
	if r.cst != nil {
		locs := r.cst.FindStringSubmatchIndex(s)
		if locs == nil {
			return nothing
		}
		out := make([]any, len(locs)/2)
		for i := range out {
			if locs[2*i] < 0 {
				out[i] = nothing
			} else {
				out[i] = just(s[locs[2*i]:locs[2*i+1]])
			}
		}
		return just(out)
	}
	input := regexRunes(s, strings.Contains(r.Flags, "u"))
	matches := regexMatches(r, input, r.Global)
	if len(matches) == 0 {
		return nothing
	}
	if !r.Global {
		return just(regexGroups(matches[0], just, nothing))
	}
	out := make([]any, len(matches))
	for i, m := range matches {
		out[i] = just(regexString(m.Runes()))
	}
	return just(out)
}

func Data_String_Regex__Search(just func(any) any, nothing any, r *GoRegex, s string) any {
	input := regexRunes(s, strings.Contains(r.Flags, "u"))
	m := regexFind(r, input, 0, strings.Contains(r.Flags, "y"))
	if m == nil {
		return nothing
	}
	index := m.Index
	if strings.Contains(r.Flags, "u") {
		for _, c := range input[:m.Index] {
			if c > 0xffff {
				index++
			}
		}
	}
	return just(index)
}

func regexReplace(r *GoRegex, s string, replacement func(*regexp2.Match, []rune) string) string {
	input := regexRunes(s, strings.Contains(r.Flags, "u"))
	matches := regexMatches(r, input, r.Global)
	var b strings.Builder
	end := 0
	for _, m := range matches {
		b.WriteString(regexString(input[end:m.Index]))
		b.WriteString(replacement(m, input))
		end = m.Index + m.Length
	}
	b.WriteString(regexString(input[end:]))
	return regexString(regexRunes(b.String(), false))
}

func Data_String_Regex__ReplaceBy(just func(any) any, nothing any, r *GoRegex, f func(string) func([]any) string, s string) string {
	return regexReplace(r, s, func(m *regexp2.Match, _ []rune) string {
		return f(regexString(m.Runes()))(regexGroups(m, just, nothing)[1:])
	})
}

func Data_String_Regex_Replace(r *GoRegex, replacement string, s string) string {
	if !strings.Contains(replacement, "$") {
		return regexReplace(r, s, func(_ *regexp2.Match, _ []rune) string {
			return replacement
		})
	}
	return regexReplace(r, s, func(m *regexp2.Match, input []rune) string {
		var b strings.Builder
		groups := m.Groups()
		for i := 0; i < len(replacement); i++ {
			if replacement[i] != '$' || i+1 == len(replacement) {
				b.WriteByte(replacement[i])
				continue
			}
			next := replacement[i+1]
			switch next {
			case '$':
				b.WriteByte('$')
				i++
			case '&':
				b.WriteString(regexString(m.Runes()))
				i++
			case '`':
				b.WriteString(regexString(input[:m.Index]))
				i++
			case '\'':
				b.WriteString(regexString(input[m.Index+m.Length:]))
				i++
			case '<':
				end := strings.IndexByte(replacement[i+2:], '>')
				if len(r.names) == 0 || end < 0 {
					b.WriteByte('$')
					continue
				}
				name := replacement[i+2 : i+2+end]
				if n, ok := r.names[name]; ok {
					b.WriteString(regexString(groups[n].Runes()))
				}
				i += end + 2
			default:
				if next < '0' || next > '9' {
					b.WriteByte('$')
					continue
				}
				n, width := int(next-'0'), 1
				if i+2 < len(replacement) && replacement[i+2] >= '0' && replacement[i+2] <= '9' {
					n2 := n*10 + int(replacement[i+2]-'0')
					if n2 > 0 && n2 < len(groups) {
						n, width = n2, 2
					}
				}
				if n == 0 || n >= len(groups) {
					b.WriteByte('$')
					continue
				}
				if len(groups[n].Captures) > 0 {
					b.WriteString(regexString(groups[n].Runes()))
				}
				i += width
			}
		}
		return b.String()
	})
}

func Data_String_Regex_Split(r *GoRegex, s string) []string {
	input := regexRunes(s, strings.Contains(r.Flags, "u"))
	if len(input) == 0 {
		if regexFind(r, input, 0, true) != nil {
			return []string{}
		}
		return []string{""}
	}
	out := []string{}
	last, pos := 0, 0
	for pos < len(input) {
		m := regexFind(r, input, pos, true)
		if m == nil {
			pos++
			continue
		}
		end := m.Index + m.Length
		if end == last {
			pos++
			continue
		}
		out = append(out, regexString(input[last:m.Index]))
		for _, g := range m.Groups()[1:] {
			if len(g.Captures) == 0 {
				panic("Data.String.Regex.split: unmatched captures cannot inhabit Array String")
			}
			out = append(out, regexString(g.Runes()))
		}
		last, pos = end, end
	}
	return append(out, regexString(input[last:]))
}

func Data_String_Regex_Test(r *GoRegex, s string) bool {
	return regexFind(r, regexRunes(s, strings.Contains(r.Flags, "u")), 0, strings.Contains(r.Flags, "y")) != nil
}

func Data_String_Regex_FlagsImpl(r *GoRegex) map[string]any {
	return map[string]any{"global": r.Global, "ignoreCase": strings.Contains(r.Flags, "i"), "multiline": strings.Contains(r.Flags, "m"), "dotAll": strings.Contains(r.Flags, "s"), "sticky": strings.Contains(r.Flags, "y"), "unicode": strings.Contains(r.Flags, "u")}
}

func Data_String_Regex_Source(r *GoRegex) string {
	if r.Source == "" {
		return "(?:)"
	}
	var b strings.Builder
	escaped := false
	for _, c := range regexRunes(r.Source, true) {
		switch c {
		case '/':
			if !escaped {
				b.WriteByte('\\')
			}
			b.WriteRune(c)
		case '\n':
			b.WriteString(`\n`)
		case '\r':
			b.WriteString(`\r`)
		case 0x2028, 0x2029:
			b.WriteString(`\u` + strconv.FormatInt(int64(c), 16))
		default:
			b.WriteString(regexString([]rune{c}))
		}
		if c == '\\' {
			escaped = !escaped
		} else {
			escaped = false
		}
	}
	return b.String()
}

func Data_String_Regex_ShowRegexImpl(r *GoRegex) string {
	return "/" + Data_String_Regex_Source(r) + "/" + r.Flags
}

// --- Auto-generated FFI wrappers ---
var _Gopurs_Data_String_Regex__Match = // TAST: (Func [(ForAll [r] (Func [(TypeVar r)] (ADT ["Data","Maybe","Maybe"] [(TypeVar r)]))), (ForAll [r] (ADT ["Data","Maybe","Maybe"] [(TypeVar r)])), (ADT ["Data","String","Regex","Regex"] []), String] (ADT ["Data","Maybe","Maybe"] [(Array (ADT ["Data","Maybe","Maybe"] [String]))]))
gopurs_runtime.Func4(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any) any {
		return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
	}
	go_arg1 := arg1
	var go_arg2 *GoRegex
	if arg2.Type == gopurs_runtime.TypeAny {
		go_arg2 = arg2.AnyVal().(*GoRegex)
	} else {
		go_arg2 = *(**GoRegex)(arg2.UnsafePtr)
	}
	go_arg3 := gopurs_runtime.Unbox[string](arg3)
	go_res := Data_String_Regex__Match(go_arg0, go_arg1, go_arg2, go_arg3)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_String_Regex__ReplaceBy = // TAST: (Func [(ForAll [r] (Func [(TypeVar r)] (ADT ["Data","Maybe","Maybe"] [(TypeVar r)]))), (ForAll [r] (ADT ["Data","Maybe","Maybe"] [(TypeVar r)])), (ADT ["Data","String","Regex","Regex"] []), (Func [String, (Array (ADT ["Data","Maybe","Maybe"] [String]))] String), String] String)
gopurs_runtime.Func5(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any) any {
		return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
	}
	go_arg1 := arg1
	var go_arg2 *GoRegex
	if arg2.Type == gopurs_runtime.TypeAny {
		go_arg2 = arg2.AnyVal().(*GoRegex)
	} else {
		go_arg2 = *(**GoRegex)(arg2.UnsafePtr)
	}
	go_arg3 := func(p0_0 string) func([]any) string {
		inner_res0 := gopurs_runtime.Apply(arg3, gopurs_runtime.Box(p0_0))
		return func(p1_0 []any) string {
			inner_res1 := gopurs_runtime.Apply(inner_res0, func() gopurs_runtime.Value {
				res_arr := make([]gopurs_runtime.Value, len(p1_0))
				for i, v := range p1_0 {
					res_arr[i] = gopurs_runtime.Box(v)
				}
				return gopurs_runtime.Array(res_arr)
			}())
			return gopurs_runtime.Unbox[string](inner_res1)
		}
	}
	go_arg4 := gopurs_runtime.Unbox[string](arg4)
	go_res := Data_String_Regex__ReplaceBy(go_arg0, go_arg1, go_arg2, go_arg3, go_arg4)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_String_Regex__Search = // TAST: (Func [(ForAll [r] (Func [(TypeVar r)] (ADT ["Data","Maybe","Maybe"] [(TypeVar r)]))), (ForAll [r] (ADT ["Data","Maybe","Maybe"] [(TypeVar r)])), (ADT ["Data","String","Regex","Regex"] []), String] (ADT ["Data","Maybe","Maybe"] [Int]))
gopurs_runtime.Func4(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any) any {
		return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
	}
	go_arg1 := arg1
	var go_arg2 *GoRegex
	if arg2.Type == gopurs_runtime.TypeAny {
		go_arg2 = arg2.AnyVal().(*GoRegex)
	} else {
		go_arg2 = *(**GoRegex)(arg2.UnsafePtr)
	}
	go_arg3 := gopurs_runtime.Unbox[string](arg3)
	go_res := Data_String_Regex__Search(go_arg0, go_arg1, go_arg2, go_arg3)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_String_Regex_FlagsImpl = // TAST: (Func [(ADT ["Data","String","Regex","Regex"] [])] (Record (Row [global: Boolean, ignoreCase: Boolean, multiline: Boolean, dotAll: Boolean, sticky: Boolean, unicode: Boolean] Empty)))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	var go_arg0 *GoRegex
	if arg0.Type == gopurs_runtime.TypeAny {
		go_arg0 = arg0.AnyVal().(*GoRegex)
	} else {
		go_arg0 = *(**GoRegex)(arg0.UnsafePtr)
	}
	go_res := Data_String_Regex_FlagsImpl(go_arg0)
	return func() gopurs_runtime.Value {
		_raw := go_res
		res_map := make(map[string]gopurs_runtime.Value)
		res_map["global"] = gopurs_runtime.Box(_raw["global"])
		res_map["ignoreCase"] = gopurs_runtime.Box(_raw["ignoreCase"])
		res_map["multiline"] = gopurs_runtime.Box(_raw["multiline"])
		res_map["dotAll"] = gopurs_runtime.Box(_raw["dotAll"])
		res_map["sticky"] = gopurs_runtime.Box(_raw["sticky"])
		res_map["unicode"] = gopurs_runtime.Box(_raw["unicode"])
		return gopurs_runtime.Record(res_map)
	}()
})
var _Gopurs_Data_String_Regex_RegexImpl = // TAST: (Func [(Func [String] (ADT ["Data","Either","Either"] [String, (ADT ["Data","String","Regex","Regex"] [])])), (Func [(ADT ["Data","String","Regex","Regex"] [])] (ADT ["Data","Either","Either"] [String, (ADT ["Data","String","Regex","Regex"] [])])), String, String] (ADT ["Data","Either","Either"] [String, (ADT ["Data","String","Regex","Regex"] [])]))
gopurs_runtime.Func4(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 string) any {
		return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
	}
	go_arg1 := func(p0_0 any) any {
		return gopurs_runtime.Apply(arg1, gopurs_runtime.Box(p0_0))
	}
	go_arg2 := gopurs_runtime.Unbox[string](arg2)
	go_arg3 := gopurs_runtime.Unbox[string](arg3)
	go_res := Data_String_Regex_RegexImpl(go_arg0, go_arg1, go_arg2, go_arg3)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_String_Regex_Replace = // TAST: (Func [(ADT ["Data","String","Regex","Regex"] []), String, String] String)
gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	var go_arg0 *GoRegex
	if arg0.Type == gopurs_runtime.TypeAny {
		go_arg0 = arg0.AnyVal().(*GoRegex)
	} else {
		go_arg0 = *(**GoRegex)(arg0.UnsafePtr)
	}
	go_arg1 := gopurs_runtime.Unbox[string](arg1)
	go_arg2 := gopurs_runtime.Unbox[string](arg2)
	go_res := Data_String_Regex_Replace(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_String_Regex_ShowRegexImpl = // TAST: (Func [(ADT ["Data","String","Regex","Regex"] [])] String)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	var go_arg0 *GoRegex
	if arg0.Type == gopurs_runtime.TypeAny {
		go_arg0 = arg0.AnyVal().(*GoRegex)
	} else {
		go_arg0 = *(**GoRegex)(arg0.UnsafePtr)
	}
	go_res := Data_String_Regex_ShowRegexImpl(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_String_Regex_Source = // TAST: (Func [(ADT ["Data","String","Regex","Regex"] [])] String)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	var go_arg0 *GoRegex
	if arg0.Type == gopurs_runtime.TypeAny {
		go_arg0 = arg0.AnyVal().(*GoRegex)
	} else {
		go_arg0 = *(**GoRegex)(arg0.UnsafePtr)
	}
	go_res := Data_String_Regex_Source(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_String_Regex_Split = // TAST: (Func [(ADT ["Data","String","Regex","Regex"] []), String] (Array String))
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	var go_arg0 *GoRegex
	if arg0.Type == gopurs_runtime.TypeAny {
		go_arg0 = arg0.AnyVal().(*GoRegex)
	} else {
		go_arg0 = *(**GoRegex)(arg0.UnsafePtr)
	}
	go_arg1 := gopurs_runtime.Unbox[string](arg1)
	go_res := Data_String_Regex_Split(go_arg0, go_arg1)
	return func() gopurs_runtime.Value {
		res_arr := make([]gopurs_runtime.Value, len(go_res))
		for i, v := range go_res {
			res_arr[i] = gopurs_runtime.Box(v)
		}
		return gopurs_runtime.Array(res_arr)
	}()
})
var _Gopurs_Data_String_Regex_Test = // TAST: (Func [(ADT ["Data","String","Regex","Regex"] []), String] Boolean)
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	var go_arg0 *GoRegex
	if arg0.Type == gopurs_runtime.TypeAny {
		go_arg0 = arg0.AnyVal().(*GoRegex)
	} else {
		go_arg0 = *(**GoRegex)(arg0.UnsafePtr)
	}
	go_arg1 := gopurs_runtime.Unbox[string](arg1)
	go_res := Data_String_Regex_Test(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})

// --- Boxed-result FFI workers ---
func Data_String_Regex_RegexImpl_nativeWorker(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 string, arg3 string) gopurs_runtime.Value {
	go_arg0 := func(p0_0 string) any {
		return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
	}
	go_arg1 := func(p0_0 any) any {
		return gopurs_runtime.Apply(arg1, gopurs_runtime.Box(p0_0))
	}
	go_arg2 := arg2
	go_arg3 := arg3
	go_res := Data_String_Regex_RegexImpl(go_arg0, go_arg1, go_arg2, go_arg3)
	return gopurs_runtime.Box(go_res)
}
