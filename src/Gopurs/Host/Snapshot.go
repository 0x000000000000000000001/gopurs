// Handwritten FFI for Gopurs.Host.Snapshot; keep in sync with the .purs foreign imports.
package Gopurs_Host_Snapshot

import (
	"fmt"
	"gopurs/output/gopurs_build_cache"
	"gopurs/output/gopurs_runtime"
	"os"
)

func WithSnapshot(run gopurs_runtime.Value, _ any) any {
	if err := gopurs_build_cache.FreezeNativeCompiler(); err != nil {
		fmt.Fprintln(os.Stderr, "[gopurs] compiler snapshot:", err)
		os.Exit(1)
	}
	gopurs_runtime.Apply(run, gopurs_runtime.Value{})
	return nil
}

func Release(_ any) any {
	gopurs_build_cache.ReleaseCompilerSnapshot()
	return nil
}
