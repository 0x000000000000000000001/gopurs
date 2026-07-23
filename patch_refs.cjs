const fs = require('fs');

let refGo = fs.readFileSync('../gopurs-refs/src/Effect/Ref.go', 'utf8');
refGo = refGo.replace(
	`			return gopurs_runtime.Record(map[string]gopurs_runtime.Value{
				"value": val,
				"state": gopurs_runtime.Any(nil),
			})`,
	`			return gopurs_runtime.RecordDict([]string{"state", "value"}, []gopurs_runtime.Value{gopurs_runtime.Any(nil), val})`);
fs.writeFileSync('../gopurs-refs/src/Effect/Ref.go', refGo);
