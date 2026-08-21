import fs from 'fs';
let code = fs.readFileSync('bin/gopurs.js', 'utf8');
code = code.replace('const v1 = optimize((() => {', `
  console.log("Optimizing:", v._2);
  const v1 = optimize((() => {
`);
fs.writeFileSync('bin/gopurs.js', code);
