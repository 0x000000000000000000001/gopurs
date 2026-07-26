const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');

const target = `
          case resObj.exprType of
            TypeRecord fields ->
              let
                staticUpdates = Array.catMaybes (map (\\(Tuple key val) ->
                  case Array.findIndex (\\(Tuple k _) -> k == key) fields of
                    Just idx -> Just (Tuple idx val)
                    Nothing -> Nothing
                ) accProps.exprs)
              in
                if Array.length staticUpdates == Array.length accProps.exprs then
                  { stmts: resObj.stmts <> accProps.stmts, expr: GoRecordUpdateStatic (boxGoExpr resObj.expr resObj.exprType) (Array.length fields) staticUpdates, exprType: TypeValue, nextId: accProps.nextId }
                else
                  { stmts: resObj.stmts <> accProps.stmts, expr: GoRecordUpdateDict (boxGoExpr resObj.expr resObj.exprType) accProps.exprs, exprType: TypeValue, nextId: accProps.nextId }
            _ ->
              { stmts: resObj.stmts <> accProps.stmts, expr: GoRecordUpdateDict (boxGoExpr resObj.expr resObj.exprType) accProps.exprs, exprType: TypeValue, nextId: accProps.nextId }
`;

const replacement = `
          case resObj.exprType of
            TypeRecord fields ->
              let nativeUpdates = map (\\(Tuple key val) -> Tuple (sanitizeName key) val) accProps.exprs
              in { stmts: resObj.stmts <> accProps.stmts, expr: GoRecordUpdateNative (boxGoExpr resObj.expr resObj.exprType) (goTypeToStr resObj.exprType) nativeUpdates, exprType: resObj.exprType, nextId: accProps.nextId }
            _ ->
              { stmts: resObj.stmts <> accProps.stmts, expr: GoRecordUpdateDict (boxGoExpr resObj.expr resObj.exprType) accProps.exprs, exprType: TypeValue, nextId: accProps.nextId }
`;

// we need to be careful with exact whitespace
// so let's just do a manual replacement using regex if needed
code = code.replace(/case resObj\.exprType of[\s\S]*?nextId: accProps\.nextId }/g, replacement.trim());
fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
