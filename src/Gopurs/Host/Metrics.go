// Handwritten FFI for Gopurs.Host.Metrics; keep in sync with the .purs foreign imports.
package Gopurs_Host_Metrics

import (
	"runtime"
	"time"
)

func SetMemProfileRate(rate int64, _ any) {
	runtime.MemProfileRate = int(rate)
}

var metricsEpoch = time.Now()

func Now() float64 {
	return float64(time.Since(metricsEpoch).Nanoseconds()) / 1e6
}
