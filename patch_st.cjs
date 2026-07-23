const fs = require('fs');

let stGo = fs.readFileSync('../gopurs-st/src/Control/Monad/ST/Internal.go', 'utf8');
stGo = stGo.replace(
	`			return gopurs_runtime.Record(map[string]gopurs_runtime.Value{
				"value": val,
				"state": gopurs_runtime.Any(nil),
			})`,
	`			return gopurs_runtime.RecordDict([]string{"state", "value"}, []gopurs_runtime.Value{gopurs_runtime.Any(nil), val})`);
fs.writeFileSync('../gopurs-st/src/Control/Monad/ST/Internal.go', stGo);
