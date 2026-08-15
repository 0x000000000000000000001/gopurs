const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');

code = code.replace(/StmtLeaf \(GoAssign name \(unsafePerformEffect[^]*?\*> pure resBinding\.expr\)\)/, 
`StmtLeaf (GoAssign name (const resBinding.expr (unsafePerformEffect (if String.contains (Pattern "__local_var") name then Effect.Console.log ("MakeLet " <> name <> " (" <> modNameStr <> ") -> " <> synName <> ", expectedGoType: " <> goTypeToStr expectedGoType) else pure unit))))`);

code = code.replace(/StmtLeaf \(GoRaw \("var " <> name <> " " <> goTypeToStr expectedGoType <> " = " <> printGoExpr \(unboxGoExpr \(unsafePerformEffect[^]*?\*> pure resBinding\.expr\) resBinding\.exprType expectedGoType\)\)\)/,
`StmtLeaf (GoRaw ("var " <> name <> " " <> goTypeToStr expectedGoType <> " = " <> printGoExpr (unboxGoExpr (const resBinding.expr (unsafePerformEffect (if String.contains (Pattern "__local_var") name then Effect.Console.log ("MakeLet " <> name <> " (" <> modNameStr <> ") -> " <> synName <> ", expectedGoType: " <> goTypeToStr expectedGoType) else pure unit))) resBinding.exprType expectedGoType)))`);

fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
