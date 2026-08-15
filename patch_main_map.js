import fs from 'fs';
const file = '/Users/0x1/Documents/htdocs/gopurs/gopurs/tests/runner/output/purescript/Main.go';
const lines = fs.readFileSync(file, 'utf8').split('\n');
let print = false;
for (let i=0; i<lines.length; i++) {
  if (lines[i].includes('Get_Data_Traversable_traverseArrayImpl()')) {
    print = true;
  }
  if (print && i < 700) {
    console.log(lines[i]);
  }
}
