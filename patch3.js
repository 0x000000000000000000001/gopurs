import fs from 'fs';
let code = fs.readFileSync('bin/gopurs.js', 'utf8');
code = code.replace('var analyzeBind = (v) => {', `
var analyzeBind = (v) => {
  if (v.tag === "TcoBinding") {
    console.log("Analyzing bind:", v._2);
  }
`);
fs.writeFileSync('bin/gopurs.js', code);
