const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');

const targetBranch = `
  Branch branches def ->
    let resDef = translateExprImpl helpersRef (depth + 1) modNameStr recVars namedBound bound Nothing loopCtx isTail nextId def
        tmpVar = "__t" <> show resDef.nextId
        declTmp = StmtLeaf (GoRaw ("var " <> tmpVar <> " gopurs_runtime.Value"))
        
        buildIfs = Array.foldr
          ( \\(Pair condExpr bodyExpr) accNext ->
              let resCond = translateExprImpl helpersRef (depth + 1) modNameStr recVars namedBound bound Nothing [] false accNext.nextId condExpr
                  resBody = translateExprImpl helpersRef (depth + 1) modNameStr recVars namedBound bound Nothing loopCtx isTail resCond.nextId bodyExpr
                  goIf = GoIfElse resCond.expr (flattenStmts resBody.stmts <> [GoMutate tmpVar resBody.expr]) (flattenStmts accNext.stmts)
              in { stmts: resCond.stmts <> StmtLeaf goIf, nextId: resBody.nextId }
          )
          { stmts: resDef.stmts <> StmtLeaf (GoMutate tmpVar resDef.expr), nextId: resDef.nextId + 1 }
          (toArray branches)
    in { stmts: declTmp <> buildIfs.stmts, expr: GoVar tmpVar, nextId: buildIfs.nextId }
`;

const replacementBranch = `
  Branch branches def ->
    let resDef = translateExprImpl helpersRef (depth + 1) modNameStr recVars namedBound bound Nothing loopCtx isTail nextId def
        tmpVar = "__t" <> show resDef.nextId
        declTmp = StmtLeaf (GoRaw ("var " <> tmpVar <> " gopurs_runtime.Value"))
        labelName = "end_branch_" <> show resDef.nextId
        
        buildIfs = foldl
          ( \\acc (Pair condExpr bodyExpr) ->
              let resCond = translateExprImpl helpersRef (depth + 1) modNameStr recVars namedBound bound Nothing [] false acc.nextId condExpr
                  resBody = translateExprImpl helpersRef (depth + 1) modNameStr recVars namedBound bound Nothing loopCtx isTail resCond.nextId bodyExpr
                  goIf = GoIfElse resCond.expr (flattenStmts resBody.stmts <> [GoMutate tmpVar resBody.expr, GoRaw ("goto " <> labelName)]) []
              in { stmts: acc.stmts <> StmtLeaf (GoRaw "{") <> resCond.stmts <> StmtLeaf goIf <> StmtLeaf (GoRaw "}"), nextId: resBody.nextId }
          )
          { stmts: StmtEmpty, nextId: resDef.nextId + 1 }
          (toArray branches)
    in { stmts: declTmp <> buildIfs.stmts <> StmtLeaf (GoRaw "{") <> resDef.stmts <> StmtLeaf (GoMutate tmpVar resDef.expr) <> StmtLeaf (GoRaw "}") <> StmtLeaf (GoRaw (labelName <> ":")), expr: GoVar tmpVar, nextId: buildIfs.nextId }
`;

code = code.replace(/  Branch branches def ->[\s\S]*?(?=  PrimOp op -> case op of)/m, replacementBranch.trim() + '\n\n');
fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
