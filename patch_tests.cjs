const fs = require('fs');

let poly = fs.readFileSync('tests/passing/PolyLabels.go', 'utf8');
poly = poly.replace(
	`		return o.PtrVal.(map[string]gopurs_runtime.Value)[s.StrVal]`,
	`		return gopurs_runtime.RecordGet(o, s.StrVal)`);
poly = poly.replace(
	`			oldMap := o.PtrVal.(map[string]gopurs_runtime.Value)
			newMap := make(map[string]gopurs_runtime.Value)
			for k, v := range oldMap {
				newMap[k] = v
			}
			newMap[s.StrVal] = a
			return gopurs_runtime.Record(newMap)`,
	`		    if _, ok := o.PtrVal.(map[string]gopurs_runtime.Value); ok {
                return gopurs_runtime.RecordUpdate(o, map[string]gopurs_runtime.Value{s.StrVal: a})
            }
            r := o.PtrVal.(gopurs_runtime.RecordData)
            for j, k := range r.Keys {
                if k == s.StrVal {
                    newVals := make([]gopurs_runtime.Value, len(r.Vals))
                    copy(newVals, r.Vals)
                    newVals[j] = a
                    return gopurs_runtime.RecordDict(r.Keys, newVals)
                }
            }
            newKeys := make([]string, len(r.Keys)+1)
            newVals := make([]gopurs_runtime.Value, len(r.Vals)+1)
            copy(newKeys, r.Keys)
            copy(newVals, r.Vals)
            newKeys[len(r.Keys)] = s.StrVal
            newVals[len(r.Vals)] = a
            return gopurs_runtime.RecordDict(newKeys, newVals)`);
fs.writeFileSync('tests/passing/PolyLabels.go', poly);

let row = fs.readFileSync('tests/passing/RowUnion.go', 'utf8');
row = row.replace(
	`		return r.PtrVal.(map[string]gopurs_runtime.Value)["a"]`,
	`		return gopurs_runtime.RecordGet(r, "a")`);
fs.writeFileSync('tests/passing/RowUnion.go', row);
