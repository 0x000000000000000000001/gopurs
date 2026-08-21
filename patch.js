import fs from 'fs';
let code = fs.readFileSync('bin/gopurs.js', 'utf8');
code = code.replace('var evalApp = (env) => (hd) => (spine) => {', `
var stackDepth = 0;
var evalApp = (env) => (hd) => (spine) => {
  if (stackDepth > 2000) {
    console.log("CRASH AT hd =", JSON.stringify(hd));
    process.exit(1);
  }
  stackDepth++;
  try {
`);
code = code.replace(/return go\$r;\n  };\n  return go\(traceSteps \? \$List\("Cons", originalExpr, Nil\) : Nil\)\(initN\)\(originalExpr\);\n};/, `
    return go$r;
  };
  var res = go(traceSteps ? $List("Cons", originalExpr, Nil) : Nil)(initN)(originalExpr);
  stackDepth--;
  return res;
} catch (e) {
  stackDepth--;
  throw e;
}
};
`);
fs.writeFileSync('bin/gopurs.js', code);
