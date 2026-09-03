const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');

// TcoApp 1
code = code.replace(
  /let\s+expectedGoTypeFromAst = goTypeToStr \(exprTypeToGoType \(\(unsafePerformEffect \(Ref\.read helpersRef\)\)\.pointerAdtPaths\) \(\(unsafePerformEffect \(Ref\.read helpersRef\)\)\.enumAdts\) \(\(unsafePerformEffect \(Ref\.read helpersRef\)\)\.elidedCtors\) modNameStr \(getExprType tcoExpr\)\)\s+in\s+\{ stmts: accFinal\.stmts <> foldMap StmtLeaf assigns <> StmtLeaf \(GoContinue targetCtx\.ident\), expr: GoRaw \("func\(\) " <> expectedGoTypeFromAst <> " \{ panic\(\\"unreachable\\"\) \}\(\)"\), exprType: getExprType tcoExpr, nextId: accFinal\.nextId \}/,
  `let
                      expectedGoType = exprTypeToGoType ((unsafePerformEffect (Ref.read helpersRef)).pointerAdtPaths) ((unsafePerformEffect (Ref.read helpersRef)).enumAdts) ((unsafePerformEffect (Ref.read helpersRef)).elidedCtors) modNameStr (getExprType tcoExpr)
                      expectedGoTypeStr = goTypeToStr expectedGoType
                    in
                    { stmts: accFinal.stmts <> foldMap StmtLeaf assigns <> StmtLeaf (GoContinue targetCtx.ident), expr: GoRaw ("func() " <> expectedGoTypeStr <> " { panic(\\"unreachable\\") }()"), exprType: expectedGoType, nextId: accFinal.nextId }`
);

// TcoApp 2
code = code.replace(
  /let\s+expectedGoTypeFromAst = goTypeToStr \(exprTypeToGoType \(\(unsafePerformEffect \(Ref\.read helpersRef\)\)\.pointerAdtPaths\) \(\(unsafePerformEffect \(Ref\.read helpersRef\)\)\.enumAdts\) \(\(unsafePerformEffect \(Ref\.read helpersRef\)\)\.elidedCtors\) modNameStr \(getExprType tcoExpr\)\)\s+in\s+\{ stmts: accFinal\.stmts <> foldMap StmtLeaf assigns <> StmtLeaf \(GoContinue targetCtx\.ident\), expr: GoRaw \("func\(\) " <> expectedGoTypeFromAst <> " \{ panic\(\\"unreachable\\"\) \}\(\)"\), exprType: getExprType tcoExpr, nextId: accFinal\.nextId \}/,
  `let
                          expectedGoType = exprTypeToGoType ((unsafePerformEffect (Ref.read helpersRef)).pointerAdtPaths) ((unsafePerformEffect (Ref.read helpersRef)).enumAdts) ((unsafePerformEffect (Ref.read helpersRef)).elidedCtors) modNameStr (getExprType tcoExpr)
                          expectedGoTypeStr = goTypeToStr expectedGoType
                        in
                          { stmts: accFinal.stmts <> foldMap StmtLeaf assigns <> StmtLeaf (GoContinue targetCtx.ident), expr: GoRaw ("func() " <> expectedGoTypeStr <> " { panic(\\"unreachable\\") }()"), exprType: expectedGoType, nextId: accFinal.nextId }`
);

// Fail
code = code.replace(
  /let\s+expectedGoTypeFromAst = goTypeToStr \(exprTypeToGoType \(\(unsafePerformEffect \(Ref\.read helpersRef\)\)\.pointerAdtPaths\) \(\(unsafePerformEffect \(Ref\.read helpersRef\)\)\.enumAdts\) \(\(unsafePerformEffect \(Ref\.read helpersRef\)\)\.elidedCtors\) modNameStr \(getExprType tcoExpr\)\)\s+in\s+\{ stmts: StmtEmpty, expr: GoRaw \("func\(\) " <> expectedGoTypeFromAst <> " \{ panic\(" <> printGoExpr \(GoString msg\) <> "\) \}\(\)"\), exprType: getExprType tcoExpr, nextId \}/,
  `let
              expectedGoType = exprTypeToGoType ((unsafePerformEffect (Ref.read helpersRef)).pointerAdtPaths) ((unsafePerformEffect (Ref.read helpersRef)).enumAdts) ((unsafePerformEffect (Ref.read helpersRef)).elidedCtors) modNameStr (getExprType tcoExpr)
              expectedGoTypeStr = goTypeToStr expectedGoType
            in
            { stmts: StmtEmpty, expr: GoRaw ("func() " <> expectedGoTypeStr <> " { panic(" <> printGoExpr (GoString msg) <> ") }()"), exprType: expectedGoType, nextId }`
);

fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
