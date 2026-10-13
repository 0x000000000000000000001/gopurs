package Gopurs_Cache_Pool

import (
	"gopurs/output/gopurs_runtime"
	"strings"
)

type pool struct {
	index map[string]int64
	nodes []string
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
