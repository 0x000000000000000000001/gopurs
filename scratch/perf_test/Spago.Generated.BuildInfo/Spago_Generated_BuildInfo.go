package Spago_Generated_BuildInfo

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var spagoVersion gopurs_runtime.Value
var once_spagoVersion sync.Once
func Get_spagoVersion() gopurs_runtime.Value {
	once_spagoVersion.Do(func() {
		spagoVersion = gopurs_runtime.Str("1.0.3")
	})
	return spagoVersion
}

var pursVersion gopurs_runtime.Value
var once_pursVersion sync.Once
func Get_pursVersion() gopurs_runtime.Value {
	once_pursVersion.Do(func() {
		pursVersion = gopurs_runtime.Str("0.15.15")
	})
	return pursVersion
}

var packages gopurs_runtime.Value
var once_packages sync.Once
func Get_packages() gopurs_runtime.Value {
	once_packages.Do(func() {
		packages = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"runner": gopurs_runtime.Str("0.0.0")})
	})
	return packages
}


