// Handwritten FFI for Gopurs.Cache.Pool; keep in sync with the .purs foreign imports.
package Gopurs_Cache_Pool

import (
	"gopurs/output/gopurs_runtime"
	"strconv"
	"strings"
)

type pool struct {
	index      map[string]int64
	nodes      []string
	references []string
}

func Create(_ any) gopurs_runtime.Value {
	return gopurs_runtime.Any(&pool{index: map[string]int64{}})
}

func Intern(handle gopurs_runtime.Value, text string, _ any) int64 {
	p := handle.AnyVal().(*pool)
	if index, ok := p.index[text]; ok {
		return index
	}
	index := int64(len(p.nodes))
	p.index[text] = index
	p.nodes = append(p.nodes, text)
	return index
}

func Finish(handle gopurs_runtime.Value, _ any) string {
	return strings.Join(handle.AnyVal().(*pool).nodes, ",")
}

// The reference is owned by this portable-artifact pool, like its JSON nodes.
// Numeric Intern and node insertion order retain their original behavior.
func InternRef(handle gopurs_runtime.Value, text string, effect any) string {
	index := Intern(handle, text, effect)
	p := handle.AnyVal().(*pool)
	for int64(len(p.references)) <= index {
		p.references = append(p.references, "")
	}
	if reference := p.references[index]; reference != "" {
		return reference
	}
	var buffer [32]byte
	n := copy(buffer[:], `["int","`)
	bytes := strconv.AppendInt(buffer[:n], index, 10)
	bytes = append(bytes, '"', ']')
	reference := string(bytes)
	p.references[index] = reference
	return reference
}
