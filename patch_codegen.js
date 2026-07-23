const fs = require('fs');

let content = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');

const appOld = `            Nothing ->
              let
                resFn = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars namedBound bound Nothing [] false false nextId fn
                accArgs = foldl
                  ( \\acc arg ->
                      let
                        argRes = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars namedBound bound Nothing [] false false acc.nextId arg
                      in
                        { stmts: acc.stmts <> argRes.stmts, exprs: Array.snoc acc.exprs argRes.expr, nextId: argRes.nextId }
                  )
                  { stmts: resFn.stmts, exprs: [], nextId: resFn.nextId }
                  argsArr

                finalExpr = Array.foldl (\\acc argExpr -> GoCall (GoSelector (GoVar "gopurs_runtime") "Apply") [ acc, argExpr ]) resFn.expr accArgs.exprs
              in
                { stmts: accArgs.stmts, expr: finalExpr, nextId: accArgs.nextId }`;

const appNew = `            Nothing ->
              let
                resFn = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars namedBound bound Nothing [] false false nextId flatFn
                accArgs = foldl
                  ( \\acc arg ->
                      let
                        argRes = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars namedBound bound Nothing [] false false acc.nextId arg
                      in
                        { stmts: acc.stmts <> argRes.stmts, exprs: Array.snoc acc.exprs argRes.expr, nextId: argRes.nextId }
                  )
                  { stmts: resFn.stmts, exprs: [], nextId: resFn.nextId }
                  flatArgs

                buildApp :: GoExpr -> Array GoExpr -> GoExpr
                buildApp fExpr argExprs =
                  let len = Array.length argExprs
                  in
                    if len == 0 then fExpr
                    else if len == 1 then GoCall (GoSelector (GoVar "gopurs_runtime") "Apply") [ fExpr, fromMaybe (GoRaw "nil") (Array.index argExprs 0) ]
                    else if len >= 2 && len <= 5 then
                      GoCall (GoSelector (GoVar "gopurs_runtime") ("Apply" <> show len)) (Array.cons fExpr argExprs)
                    else
                      let chunk = Array.take 5 argExprs
                          rest = Array.drop 5 argExprs
                      in buildApp (buildApp fExpr chunk) rest

                finalExpr = buildApp resFn.expr accArgs.exprs
              in
                { stmts: accArgs.stmts, expr: finalExpr, nextId: accArgs.nextId }`;

const absOld = `      Abs args body ->
        let
          params = map (\\(Tuple mbI lvl) -> localId mbI lvl) (toArray args)
          resBody = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars namedBound bound Nothing loopCtx isTail false nextId body
          funcExpr = Array.foldr
            ( \\p acc ->
                let
                  bodyStr = case acc of
                    GoBlock _ -> printGoExpr acc
                    _ -> "return " <> printGoExpr acc
                in
                  GoCall (GoSelector (GoVar "gopurs_runtime") "Func") [ GoRaw ("func(" <> p <> " gopurs_runtime.Value) gopurs_runtime.Value {\\n" <> bodyStr <> "\\n}") ]
            )
            (GoBlock (flattenStmts resBody.stmts <> [ GoReturn resBody.expr ]))
            params
        in
          { stmts: StmtEmpty, expr: funcExpr, nextId: resBody.nextId }`;

const absNew = `      Abs args body ->
        let
          params = map (\\(Tuple mbI lvl) -> localId mbI lvl) (toArray args)
          resBody = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars namedBound bound Nothing loopCtx isTail false nextId body
          
          buildFunc :: Array String -> GoExpr -> GoExpr
          buildFunc ps innerExpr =
            let
              len = Array.length ps
              bodyStr = case innerExpr of
                GoBlock _ -> printGoExpr innerExpr
                _ -> "return " <> printGoExpr innerExpr
            in
              if len == 1 then
                GoCall (GoSelector (GoVar "gopurs_runtime") "Func") [ GoRaw ("func(" <> fromMaybe "" (Array.index ps 0) <> " gopurs_runtime.Value) gopurs_runtime.Value {\\n" <> bodyStr <> "\\n}") ]
              else if len >= 2 && len <= 5 then
                let goParams = String.joinWith ", " (map (\\p -> p <> " gopurs_runtime.Value") ps)
                in GoCall (GoSelector (GoVar "gopurs_runtime") ("Func" <> show len)) [ GoRaw ("func(" <> goParams <> ") gopurs_runtime.Value {\\n" <> bodyStr <> "\\n}") ]
              else
                let chunk = Array.take 5 ps
                    rest = Array.drop 5 ps
                in buildFunc chunk (buildFunc rest innerExpr)
                
          funcExpr = buildFunc params (GoBlock (flattenStmts resBody.stmts <> [ GoReturn resBody.expr ]))
        in
          { stmts: StmtEmpty, expr: funcExpr, nextId: resBody.nextId }`;

if (content.includes(appOld) && content.includes(absOld)) {
  content = content.replace(appOld, appNew);
  content = content.replace(absOld, absNew);
  fs.writeFileSync('src/Gopurs/CodeGen.purs', content, 'utf8');
  console.log("Patched CodeGen.purs successfully");
} else {
  console.log("Could not find the target code in CodeGen.purs");
}
