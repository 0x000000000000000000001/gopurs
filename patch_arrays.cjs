const fs = require('fs');

// Data/Array.go
let arrayGo = fs.readFileSync('../gopurs-arrays/src/Data/Array.go', 'utf8');
arrayGo = arrayGo.replace(
	`return gopurs_runtime.Record(map[string]gopurs_runtime.Value{
			"yes": gopurs_runtime.Array(yes),
			"no":  gopurs_runtime.Array(no),
		})`,
	`return gopurs_runtime.RecordDict([]string{"no", "yes"}, []gopurs_runtime.Value{gopurs_runtime.Array(no), gopurs_runtime.Array(yes)})`);
fs.writeFileSync('../gopurs-arrays/src/Data/Array.go', arrayGo);

// Data/Array/ST.go
let arraySt = fs.readFileSync('../gopurs-arrays/src/Data/Array/ST.go', 'utf8');
arraySt = arraySt.replace(
	`		dict := map[string]gopurs_runtime.Value{
			"value": v,
			"index": gopurs_runtime.Int(int64(i)),
		}
		res[i] = gopurs_runtime.Record(dict)`,
	`		res[i] = gopurs_runtime.RecordDict([]string{"index", "value"}, []gopurs_runtime.Value{gopurs_runtime.Int(int64(i)), v})`);
fs.writeFileSync('../gopurs-arrays/src/Data/Array/ST.go', arraySt);
