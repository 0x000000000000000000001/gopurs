package Gopurs_Metrics

import (
	"runtime"
	"time"
)

func SetMemProfileRate(rate int64, _ gopurs_runtime.Value) gopurs_runtime.Value {
	runtime.MemProfileRate = int(rate)
	return gopurs_runtime.Value{}
}

var metricsEpoch = time.Now()

func Now() float64 {
	return float64(time.Since(metricsEpoch).Nanoseconds()) / 1e6
}
