const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');

code = code.replace("          _traceLet = if String.contains (Pattern \"__local_var_2_0\") name then Debug.trace (\"\\n==== LET \" <> name <> \" ===\\nexpectedType=\" <> goTypeToStr expectedGoType <> \"\\nactualType=\" <> goTypeToStr resBinding.exprType <> \"\\nAST=\" <> printTcoExprShape binding <> \"\\nTAST_TYPE=\" <> printExprType (getExprType binding) <> \"\\n==================\\n\") (\\_ -> unit) else unit\n", "          _traceLet = if String.contains (Pattern \"arrayMap\") (printTcoExprShape binding) then Debug.trace (\"\\n==== LET \" <> name <> \" ===\\nexpectedType=\" <> goTypeToStr expectedGoType <> \"\\nactualType=\" <> goTypeToStr resBinding.exprType <> \"\\nAST=\" <> printTcoExprShape binding <> \"\\nTAST_TYPE=\" <> printExprType (getExprType binding) <> \"\\n==================\\n\") (\\_ -> unit) else unit\n");

fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
