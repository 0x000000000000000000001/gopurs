import fs from 'fs';
const file = '/Users/0x1/Documents/htdocs/gopurs/gopurs/tests/runner/output/purescript/Main.go';
let content = fs.readFileSync(file, 'utf8');

const oldCode = `			arr_prime__1_2 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = arr_prime__1_2
			// TAST (Let): __local_var_2_4 -> gopurs_runtime.Value
			__local_var_2_4 := gopurs_runtime.Apply2(Get_Data_Array_mapMaybe(), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {`;

const newCode = `			arr_prime__1_2 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = arr_prime__1_2
            fmt.Printf("arr_prime__1_2 TYPE: %d, UnsafePtr: %v\\n", arr_prime__1_2.Type, arr_prime__1_2.UnsafePtr)
			// TAST (Let): __local_var_2_4 -> gopurs_runtime.Value
			__local_var_2_4 := gopurs_runtime.Apply2(Get_Data_Array_mapMaybe(), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {`;

if (content.includes(oldCode)) {
  content = content.replace(oldCode, newCode);
  fs.writeFileSync(file, content);
  console.log("Patched Main.go successfully!");
} else {
  console.error("Could not find old code block.");
}
