const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');

code = code.replace(/_ = unsafePerformEffect([^]*?)newBound = Map\.insert originalName \{ name, goType: expectedGoType \} bound/, 
`newBound = Map.insert originalName { name, goType: expectedGoType } bound`);

code = code.replace(/letStmt = if expectedGoType == resBinding\.exprType then\n\s*StmtLeaf \(GoAssign name resBinding\.expr\)\n\s*else\n\s*StmtLeaf \(GoRaw \("var " <> name <> " " <> goTypeToStr expectedGoType <> " = " <> printGoExpr \(unboxGoExpr resBinding\.expr resBinding\.exprType expectedGoType\)\)\)/,
`letStmt = if expectedGoType == resBinding.exprType then
                      StmtLeaf (GoAssign name (unsafePerformEffect (if String.contains (Pattern "__local_var") name then Effect.Console.log ("MakeLet " <> name <> " (" <> modNameStr <> ") -> " <> synName <> ", expectedGoType: " <> goTypeToStr expectedGoType) else pure unit) *> pure resBinding.expr))
                    else
                      StmtLeaf (GoRaw ("var " <> name <> " " <> goTypeToStr expectedGoType <> " = " <> printGoExpr (unboxGoExpr (unsafePerformEffect (if String.contains (Pattern "__local_var") name then Effect.Console.log ("MakeLet " <> name <> " (" <> modNameStr <> ") -> " <> synName <> ", expectedGoType: " <> goTypeToStr expectedGoType) else pure unit) *> pure resBinding.expr) resBinding.exprType expectedGoType)))`);

fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
