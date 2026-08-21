import fs from 'fs';
let code = fs.readFileSync('bin/gopurs.js', 'utf8');
code = code.replace('var evalAccessor = (env) => (lhs) => (accessor) => floatLet(lhs)((v) => {', `
var evalAccessorDepth = 0;
var evalAccessor = (env) => (lhs) => (accessor) => {
  evalAccessorDepth++;
  if (evalAccessorDepth > 2000) {
    console.log("CRASH IN evalAccessor, lhs=", JSON.stringify(lhs).substring(0,200));
    process.exit(1);
  }
  var res = floatLet(lhs)((v) => {
`);
code = code.replace(/return evalRecord\(env\)\(v\._1\)\(v\._2\);\n  }\n  return fail\(\);\n}\)\)\(accessor\);/, `return evalRecord(env)(v._1)(v._2);
  }
  return fail();
}))(accessor);
evalAccessorDepth--;
return res;
`);
fs.writeFileSync('bin/gopurs.js', code);
