import fs from 'fs';
const code = fs.readFileSync('/Users/0x1/Documents/htdocs/gopurs/gopurs/tests/runner/output/purescript/Data_Array.go', 'utf8');
const lines = code.split('\n');
let inFunc = false;
let braces = 0;
for (let line of lines) {
  if (line.startsWith('func Call_Data_Array_mapMaybe(f_0_loop gopurs_runtime.Value)')) {
    inFunc = true;
  }
  if (inFunc) {
    console.log(line);
    braces += (line.match(/\{/g) || []).length;
    braces -= (line.match(/\}/g) || []).length;
    if (braces === 0) break;
  }
}
