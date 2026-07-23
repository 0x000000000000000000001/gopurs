const fs = require('fs');

let strGo = fs.readFileSync('../gopurs-strings/src/Data/String/CodeUnits.go', 'utf8');
strGo = strGo.replace(
	`			return gopurs_runtime.Record(map[string]gopurs_runtime.Value{
				"before": gopurs_runtime.Str(string(runes[:i])),
				"after":  gopurs_runtime.Str(string(runes[i:])),
			})`,
	`			return gopurs_runtime.RecordDict([]string{"after", "before"}, []gopurs_runtime.Value{gopurs_runtime.Str(string(runes[i:])), gopurs_runtime.Str(string(runes[:i]))})`);
fs.writeFileSync('../gopurs-strings/src/Data/String/CodeUnits.go', strGo);
