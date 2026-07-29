import fs from 'fs';
let code = fs.readFileSync('src/Gopurs/FfiSupport.js', 'utf8');

const parseStart = code.indexOf('let parseFuncType = function(t) {');
const parseEnd = code.indexOf('let goFuncArgsBoxed = parsedArgs', parseStart);
const blockToMove = code.substring(parseStart, parseEnd);

code = code.substring(0, parseStart) + code.substring(parseEnd);

const insertPoint = code.indexOf('// 1. Generate Native Call_X proxy');
code = code.substring(0, insertPoint) + blockToMove + '\n                ' + code.substring(insertPoint);

fs.writeFileSync('src/Gopurs/FfiSupport.js', code);
