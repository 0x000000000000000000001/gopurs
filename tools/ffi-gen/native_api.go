package main

import (
	"go/token"
	"sync"
)

// The WASM runner has one process per file. A native compiler keeps this parser
// alive, so release the preceding source positions and serialize access to the
// file set used by parseExprToTypeNode and parseFFI.
var nativeParserMutex sync.Mutex

func Extract(content string) (string, error) {
	nativeParserMutex.Lock()
	defer nativeParserMutex.Unlock()
	fset = token.NewFileSet()
	return parseFFI(content)
}

func Prepare(prefix string, content string) (string, error) {
	nativeParserMutex.Lock()
	defer nativeParserMutex.Unlock()
	fset = token.NewFileSet()
	return parseFFI(content, prefix)
}
