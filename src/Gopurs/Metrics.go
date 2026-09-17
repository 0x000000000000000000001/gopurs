package Gopurs_Metrics

import "time"

var metricsEpoch = time.Now()

func Now() float64 {
	return float64(time.Since(metricsEpoch).Nanoseconds()) / 1e6
}
