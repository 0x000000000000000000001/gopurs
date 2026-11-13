// Handwritten FFI for Gopurs.Host.Snapshot; keep in sync with the .purs foreign imports.
package Gopurs_Host_Snapshot

import (
	"fmt"
	"gopurs/output/gopurs_build_cache"
	"os"
)

func WithSnapshot(run func(), _ any) {
	if err := gopurs_build_cache.FreezeNativeCompiler(); err != nil {
		fmt.Fprintln(os.Stderr, "[gopurs] compiler snapshot:", err)
		os.Exit(1)
	}
	run()
}

func Release() {
	gopurs_build_cache.ReleaseCompilerSnapshot()
}
