package Main
import "gopurs/output/gopurs_runtime"

func getMap(val gopurs_runtime.Value) map[string]gopurs_runtime.Value {
	return gopurs_runtime.RecordToMap(val)
}

var MergeImpl = gopurs_runtime.Func(func(l gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(r gopurs_runtime.Value) gopurs_runtime.Value {
		oldMap1 := getMap(r)
		oldMap2 := getMap(l)
		newMap := make(map[string]gopurs_runtime.Value)
		for k, v := range oldMap1 {
			newMap[k] = v
		}
		for k, v := range oldMap2 {
			newMap[k] = v
		}
		return gopurs_runtime.Record(newMap)
	})
})
