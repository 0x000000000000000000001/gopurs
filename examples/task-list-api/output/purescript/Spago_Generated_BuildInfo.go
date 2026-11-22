package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Spago_Generated_BuildInfo_spagoVersion gopurs_runtime.Value
var once_Spago_Generated_BuildInfo_spagoVersion sync.Once

func Get_Spago_Generated_BuildInfo_spagoVersion() gopurs_runtime.Value {
	once_Spago_Generated_BuildInfo_spagoVersion.Do(func() {
		cache_Spago_Generated_BuildInfo_spagoVersion = gopurs_runtime.StrWithCodeUnits("1.0.3", 5)
	})
	return cache_Spago_Generated_BuildInfo_spagoVersion
}

var cache_Spago_Generated_BuildInfo_pursVersion gopurs_runtime.Value
var once_Spago_Generated_BuildInfo_pursVersion sync.Once

func Get_Spago_Generated_BuildInfo_pursVersion() gopurs_runtime.Value {
	once_Spago_Generated_BuildInfo_pursVersion.Do(func() {
		cache_Spago_Generated_BuildInfo_pursVersion = gopurs_runtime.StrWithCodeUnits("0.15.17", 7)
	})
	return cache_Spago_Generated_BuildInfo_pursVersion
}

var cache_Spago_Generated_BuildInfo_packages gopurs_runtime.Value
var once_Spago_Generated_BuildInfo_packages sync.Once

func Get_Spago_Generated_BuildInfo_packages() gopurs_runtime.Value {
	once_Spago_Generated_BuildInfo_packages.Do(func() {
		cache_Spago_Generated_BuildInfo_packages = func() gopurs_runtime.Value {
			orig := struct {
				gopurs_field_116_97_115_107_45_108_105_115_116_45_97_112_105 string
				/*gopurs:record-v1:1:007400610073006b002d006c006900730074002d006100700069*/
			}{"0.0.0"}
			_ = orig
			return gopurs_runtime.RecordDict1("task-list-api", gopurs_runtime.Str(orig.gopurs_field_116_97_115_107_45_108_105_115_116_45_97_112_105))
		}()
	})
	return cache_Spago_Generated_BuildInfo_packages
}
