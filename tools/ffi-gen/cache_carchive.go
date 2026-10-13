//go:build carchive && gopurscache

package main

/*
#include <stdlib.h>
*/
import "C"

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/0x000000000000000000001/gopurs/tools/ffi-gen/gopurs_build_cache"
	"unsafe"
)

//export GopursCacheDigest
func GopursCacheDigest(content *C.uchar, length C.size_t) *C.char {
	sum := sha256.Sum256(unsafe.Slice((*byte)(unsafe.Pointer(content)), int(length)))
	return C.CString(hex.EncodeToString(sum[:]))
}

// The bootstrap copies the shared cache package next to this archive. Keep one
// Go runtime for the Rust host's parser, cache and compiler snapshot operations.
//
//export GopursBuildCache
func GopursBuildCache(content *C.uchar, length C.size_t, session C.int) *C.char {
	request := string(unsafe.Slice((*byte)(unsafe.Pointer(content)), int(length)))
	if session != 0 {
		return C.CString(gopurs_build_cache.SessionExchange(request))
	}
	return C.CString(gopurs_build_cache.Exchange(request))
}

//export GopursFreezeCompiler
func GopursFreezeCompiler() *C.char {
	if err := gopurs_build_cache.FreezeCompiler("rust"); err != nil {
		return C.CString(err.Error())
	}
	return nil
}

//export GopursReleaseCompiler
func GopursReleaseCompiler() {
	gopurs_build_cache.ReleaseCompilerSnapshot()
}
