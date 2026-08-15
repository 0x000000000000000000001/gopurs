import fs from 'fs';
const code = fs.readFileSync('/Users/0x1/Documents/htdocs/gopurs/gopurs/tests/runner/output/purescript/Main.go', 'utf8');
const lines = code.split('\n');
let inFunc = false;
let braces = 0;
for (let line of lines) {
  if (line.includes('Get_Main_main()')) {
    inFunc = true;
  }
  if (inFunc) {
    console.log(line);
    braces += (line.match(/\{/g) || []).length;
    braces -= (line.match(/\}/g) || []).length;
    if (braces === 0) break;
  }
}
