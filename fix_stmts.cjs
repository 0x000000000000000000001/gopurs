const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');

code = code.replace(
  /in GoCall \(GoRaw \("func\(\) gopurs_runtime\.Value \{\\n" <> printGoExpr checkBlock <> "\\nreturn " <> printGoExpr allocBlock <> "\\n}\(\)"\)\) \[\]/,
  `in
                     let
                       tmpVar = "__tmut_" <> show accProps.nextId
                       declTmp = StmtLeaf (GoRaw ("var " <> tmpVar <> " gopurs_runtime.Value"))
                       assignMut = StmtLeaf (GoRaw (tmpVar <> " = " <> printGoExpr (GoBlock [ GoRaw ("gopurs_runtime.Value{Type: 9, IntVal: " <> hashString baseStructName <> ", UnsafePtr: unsafe.Pointer(" <> ptrExpr <> ")}") ])))
                       assignAlloc = StmtLeaf (GoRaw (tmpVar <> " = " <> printGoExpr allocBlock))
                       
                       checkBlockStmts = StmtBlock [ GoIfElse condType (flattenStmts [GoBlock mutateStmts, GoRaw (tmpVar <> " = gopurs_runtime.Value{Type: 9, IntVal: " <> hashString baseStructName <> ", UnsafePtr: unsafe.Pointer(" <> ptrExpr <> ")}")]) [GoRaw (tmpVar <> " = " <> printGoExpr allocBlock)] ]
                       
                     in { stmts: accProps.stmts <> StmtList [declTmp] <> checkBlockStmts, expr: GoRaw tmpVar, exprType: expectedGoType, nextId: accProps.nextId + 1 }`
);

fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
