// Handwritten FFI for Gopurs.Diagnostics; keep in sync with the .purs foreign imports.
package Gopurs_Diagnostics

import "sync/atomic"

var observed atomic.Bool

func Mark(_ any) any      { observed.Store(true); return nil }
func Recorded(_ any) bool { return observed.Load() }
