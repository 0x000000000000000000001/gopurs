const fs = require('fs');

let unsafe = fs.readFileSync('../gopurs-prelude/src/Record/Unsafe.go', 'utf8');

// UnsafeHas
unsafe = unsafe.replace(
	`		_, ok := rec.PtrVal.(map[string]gopurs_runtime.Value)[label.StrVal]
		return gopurs_runtime.Bool(ok)`,
	`		if m, ok := rec.PtrVal.(map[string]gopurs_runtime.Value); ok {
			_, ok2 := m[label.StrVal]
			return gopurs_runtime.Bool(ok2)
		}
		r := rec.PtrVal.(gopurs_runtime.RecordData)
		for _, k := range r.Keys {
			if k == label.StrVal {
				return gopurs_runtime.Bool(true)
			}
		}
		return gopurs_runtime.Bool(false)`);

// UnsafeGet
unsafe = unsafe.replace(
	`		return rec.PtrVal.(map[string]gopurs_runtime.Value)[label.StrVal]`,
	`		return gopurs_runtime.RecordGet(rec, label.StrVal)`);

// UnsafeSet
unsafe = unsafe.replace(
	`			old := rec.PtrVal.(map[string]gopurs_runtime.Value)
			newMap := make(map[string]gopurs_runtime.Value, len(old)+1)
			for k, v := range old {
				newMap[k] = v
			}
			newMap[label.StrVal] = value
			return gopurs_runtime.Record(newMap)`,
	`		    if _, ok := rec.PtrVal.(map[string]gopurs_runtime.Value); ok {
                return gopurs_runtime.RecordUpdate(rec, map[string]gopurs_runtime.Value{label.StrVal: value})
            }
            // For RecordData, if the key exists, RecordUpdateDict will overwrite it.
            // If the key does not exist (UnsafeSet adding a new key), RecordUpdateDict won't add it.
            // But UnsafeSet in PureScript is used to create NEW records dynamically.
            r := rec.PtrVal.(gopurs_runtime.RecordData)
            for j, k := range r.Keys {
                if k == label.StrVal {
                    newVals := make([]gopurs_runtime.Value, len(r.Vals))
                    copy(newVals, r.Vals)
                    newVals[j] = value
                    return gopurs_runtime.RecordDict(r.Keys, newVals)
                }
            }
            newKeys := make([]string, len(r.Keys)+1)
            newVals := make([]gopurs_runtime.Value, len(r.Vals)+1)
            copy(newKeys, r.Keys)
            copy(newVals, r.Vals)
            newKeys[len(r.Keys)] = label.StrVal
            newVals[len(r.Vals)] = value
            return gopurs_runtime.RecordDict(newKeys, newVals)`);

// UnsafeDelete
unsafe = unsafe.replace(
	`		old := rec.PtrVal.(map[string]gopurs_runtime.Value)
		newMap := make(map[string]gopurs_runtime.Value)
		for k, v := range old {
			if k != label.StrVal {
				newMap[k] = v
			}
		}
		return gopurs_runtime.Record(newMap)`,
	`		if m, ok := rec.PtrVal.(map[string]gopurs_runtime.Value); ok {
			newMap := make(map[string]gopurs_runtime.Value)
			for k, v := range m {
				if k != label.StrVal {
					newMap[k] = v
				}
			}
			return gopurs_runtime.Record(newMap)
		}
		r := rec.PtrVal.(gopurs_runtime.RecordData)
		newKeys := make([]string, 0, len(r.Keys))
		newVals := make([]gopurs_runtime.Value, 0, len(r.Vals))
		for i, k := range r.Keys {
			if k != label.StrVal {
				newKeys = append(newKeys, k)
				newVals = append(newVals, r.Vals[i])
			}
		}
		return gopurs_runtime.RecordDict(newKeys, newVals)`);

fs.writeFileSync('../gopurs-prelude/src/Record/Unsafe.go', unsafe);
