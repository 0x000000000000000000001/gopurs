import fs from 'fs';
let code = fs.readFileSync('bin/gopurs.js', 'utf8');
code = code.replace('var evalApp = (env) => (hd) => (spine) => {', `
var evalApp = (env) => (hd) => (spine) => {
  if (spine.length > 50) {
    console.log("HUGE SPINE:", spine.length, "hd =", JSON.stringify(hd).substring(0,200));
  }
`);
fs.writeFileSync('bin/gopurs.js', code);
