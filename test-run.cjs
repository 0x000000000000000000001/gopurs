const fs = require('fs');
const content = fs.readFileSync('tests/runner/output/purescript/Main.go', 'utf8');
const start = content.indexOf('func Get_Main_functorFun2()');
const end = content.indexOf('return cache_Main_functorFun2', start) + 32;
console.log(content.slice(start, end));
