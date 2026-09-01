const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf-8');

code = code.replace(
  '"\\\\t\\\\tout.V" <> show i <> " = " <> printGoExpr (coerceGoExpr modNameStr (GoStructAccess (GoVar "in") ("V" <> show i)) t1 t2)',
  `let _trace2 = unsafePerformEffect (if funcName == "Rebox_Control_Comonad_Env_Class_1649752824_2829962037" then Console.log ("i=" <> show i <> " genericTy=" <> printGoType genericTy <> " t1=" <> printGoType t1 <> " t2=" <> printGoType t2) else pure unit)\n                                    in "\\t\\tout.V" <> show i <> " = " <> printGoExpr (coerceGoExpr modNameStr (GoStructAccess (GoVar "in") ("V" <> show i)) t1 t2)`
);

fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
console.log("Patched loop with trace!");
