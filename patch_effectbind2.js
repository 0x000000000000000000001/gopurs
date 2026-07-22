import fs from 'fs';
let content = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');

content = content.replace(
  `      EffectBind mbIdent lvl binding body ->
        let
          originalName = localId mbIdent lvl
          name = originalName <> "_" <> show nextId
          newBound = Map.insert originalName name bound
          resBinding = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars namedBound bound Nothing [] false true (nextId + 1) binding
          resBody = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars namedBound newBound Nothing loopCtx isTail true resBinding.nextId body
          bindingExpr = executeIfOpaque (unwrapTcoExpr binding) resBinding.expr
        in
          { stmts: resBinding.stmts <> StmtLeaf (GoAssign name bindingExpr) <> resBody.stmts, expr: resBody.expr, nextId: resBody.nextId }`,
  `      EffectBind mbIdent lvl binding body ->
        let
          originalName = localId mbIdent lvl
          name = originalName <> "_" <> show nextId
          newBound = Map.insert originalName name bound
          resBinding = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars namedBound bound Nothing [] false true (nextId + 1) binding
          resBody = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars namedBound newBound Nothing loopCtx isTail true resBinding.nextId body
          bindingExpr = executeIfOpaque (unwrapTcoExpr binding) resBinding.expr
          bodyExpr = executeIfOpaque (unwrapTcoExpr body) resBody.expr
        in
          { stmts: resBinding.stmts <> StmtLeaf (GoAssign name bindingExpr) <> resBody.stmts, expr: bodyExpr, nextId: resBody.nextId }`
);

fs.writeFileSync('src/Gopurs/CodeGen.purs', content);
