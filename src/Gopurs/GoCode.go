package Gopurs_GoCode

import (
	"sort"

	"gopurs/output/gopurs_runtime"
)

// ReferencedImportsImpl scans Go code text for the runtime dependencies it
// mentions. Same rules and result order as the PureScript implementation, in a
// single byte pass: no code-point array, no intermediate import arrays, no
// per-character dispatch.
func ReferencedImportsImpl(fallback gopurs_runtime.Value, textValue gopurs_runtime.Value) gopurs_runtime.Value {
	_ = fallback
	text := textValue.StrVal()
	size := len(text)
	imports := make([]string, 0, 4)
	for index := 0; index < size; {
		char := text[index]
		switch {
		case char == '"' || char == '\'' || char == '`':
			index = skipQuoted(text, char, index+1)
		case char == '/' && index+1 < size && text[index+1] == '/':
			index = skipLine(text, index+2)
		case char == '/' && index+1 < size && text[index+1] == '*':
			index = skipComment(text, index+2)
		case isIdentifierByte(char):
			end := skipIdentifier(text, index+1)
			if end < size && text[end] == '.' {
				if path, ok := knownPackage(text[index:end]); ok && !containsImport(imports, path) {
					imports = append(imports, path)
				}
			}
			index = end
		default:
			index++
		}
	}
	sort.Strings(imports)
	out := make([]gopurs_runtime.Value, len(imports))
	for i, path := range imports {
		out[i] = gopurs_runtime.Str(path)
	}
	return gopurs_runtime.Array(out)
}

// isIdentifierByte mirrors the PureScript code-point test: ASCII letters,
// digits and underscore, plus every non-ASCII byte (UTF-8 continuations).
func isIdentifierByte(char byte) bool {
	switch {
	case char >= 97 && char <= 122:
		return true
	case char >= 65 && char <= 90:
		return true
	case char == 95:
		return true
	case char >= 48 && char <= 57:
		return true
	case char > 127:
		return true
	}
	return false
}

func skipIdentifier(text string, index int) int {
	for index < len(text) && isIdentifierByte(text[index]) {
		index++
	}
	return index
}

func skipQuoted(text string, quote byte, index int) int {
	for index < len(text) {
		char := text[index]
		if char == quote {
			return index + 1
		}
		if quote != '`' && char == '\\' {
			index += 2
			continue
		}
		index++
	}
	return len(text)
}

func skipLine(text string, index int) int {
	for index < len(text) && text[index] != '\n' {
		index++
	}
	return index
}

func skipComment(text string, index int) int {
	for index < len(text) {
		if text[index] == '*' && index+1 < len(text) && text[index+1] == '/' {
			return index + 2
		}
		index++
	}
	return len(text)
}

func knownPackage(name string) (string, bool) {
	switch name {
	case "gopurs_runtime":
		return "gopurs/output/gopurs_runtime", true
	case "math":
		return "math", true
	case "sync":
		return "sync", true
	case "unsafe":
		return "unsafe", true
	}
	return "", false
}

func containsImport(imports []string, path string) bool {
	for _, existing := range imports {
		if existing == path {
			return true
		}
	}
	return false
}
