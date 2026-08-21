import fs from 'fs';
let code = fs.readFileSync('bin/gopurs.js', 'utf8');
code = code.replace('var monomorphizeExpr = (globalAstMap) =>', `
var maxDepth = 0;
var getDepth = (expr) => {
  if (!expr) return 0;
  if (expr.tag === "ExprApp") return 1 + Math.max(getDepth(expr._1), getDepth(expr._2));
  if (expr.tag === "ExprLet") return 1 + getDepth(expr._2);
  if (expr.tag === "ExprAbs") return 1 + getDepth(expr._2);
  return 1;
};
var monomorphizeExpr = (globalAstMap) =>`);
code = code.replace('var f$p = monomorphizeExpr(globalAstMap)(modName)(instMap)(localDicts)(f_spine);', `
var f$p = monomorphizeExpr(globalAstMap)(modName)(instMap)(localDicts)(f_spine);
var d = getDepth(f$p);
if (d > maxDepth) { maxDepth = d; console.log("New max depth:", maxDepth); }
`);
fs.writeFileSync('bin/gopurs.js', code);
