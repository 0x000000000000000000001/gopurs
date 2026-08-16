const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');

// We want to patch the `Typed type_ a` case in `translateExprImpl_`
// If `res.exprType == TypeValue` and `isPrimitive expectedGoType` and it looks like a function, don't coerce.
const replace = `
          _, _ ->
            let res = translateExprImpl_ helpersRef depth modNameStr recVars moduleArities bound tcoIdent loopCtx isTail inEffectBlock nextId a
            in case res.exprType of
              TypeStructPointer _ _ _ _ -> res
              _ -> 
                let 
                  isPrimitive = case expectedGoType of
                    TypeBool -> true
                    TypeInt64 -> true
                    TypeFloat64 -> true
                    TypeUint32 -> true
                    TypeString -> true
                    _ -> false
                  
                  isFuncLit = case res.expr of
                    GoCall (GoSelector (GoVar "gopurs_runtime") "Func") _ -> true
                    GoFuncLit _ _ _ _ -> true
                    _ -> false
                in
                  if isPrimitive && isFuncLit && res.exprType == TypeValue then
                    -- The optimizer generated an invalid Typed node coercing a function to a primitive.
                    -- Ignore the coercion.
                    { stmts: res.stmts, expr: res.expr, exprType: TypeValue, nextId: res.nextId }
                  else
                    { stmts: res.stmts, expr: coerceGoExpr res.expr res.exprType expectedGoType, exprType: expectedGoType, nextId: res.nextId }
`;

code = code.replace(/_, _ ->\n\s*let res = translateExprImpl_ [^\n]+\n\s*in case res\.exprType of\n\s*TypeStructPointer _ _ _ _ -> res\n\s*_ -> \n\s*let \n\s*isPrimitive = case expectedGoType of[\s\S]*?else\n\s*\{ stmts: res\.stmts, expr: coerceGoExpr res\.expr res\.exprType expectedGoType, exprType: expectedGoType, nextId: res\.nextId \}/, replace);

fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
