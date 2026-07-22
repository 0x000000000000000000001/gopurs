import fs from 'fs';

let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');

if (!code.includes('import Effect.Ref')) {
    code = code.replace(/import Effect\.Unsafe \(unsafePerformEffect\)/, 'import Effect.Unsafe (unsafePerformEffect)\nimport Effect.Ref (Ref)\nimport Effect.Ref as Ref\nimport Gopurs.FreeVars (freeVars)\nimport Data.Set as Set');
}

code = code.replace(
  /translate :: Array \(Array String\) -> BackendModule -> String\ntranslate importsArray mod =\n  let\n    modNameStr = String\.replaceAll \(Pattern "\."\) \(Replacement "_"\) \(unwrap mod\.name\)\n    _ = unsafePerformEffect \(Console\.log \("Translating module " <> modNameStr\)\)/g,
  `translate :: Array (Array String) -> BackendModule -> String
translate importsArray mod =
  let
    modNameStr = String.replaceAll (Pattern ".") (Replacement "_") (unwrap mod.name)
    _ = unsafePerformEffect (Console.log ("Translating module " <> modNameStr))
    helpersRef = unsafePerformEffect (Ref.new [])`
);

code = code.replace(
  /goFile = \{ packageName: modNameStr\n             , imports: goImports\n             , decls: decls\n             , foreigns:/g,
  `goFile = { packageName: modNameStr
             , imports: goImports
             , decls: decls <> unsafePerformEffect (Ref.read helpersRef)
             , foreigns:`
);

code = code.replace(
    /translateExprImpl :: String -> Array String -> Map String String -> Map String String -> Maybe String -> Array \{ ident :: String, params :: Array String \} -> Boolean -> Int -> TcoExpr -> \{ stmts :: StmtTree, expr :: GoExpr, nextId :: Int \}/g,
    'translateExprImpl :: Ref (Array GoDecl) -> Int -> String -> Array String -> Map String String -> Map String String -> Maybe String -> Array { ident :: String, params :: Array String } -> Boolean -> Int -> TcoExpr -> { stmts :: StmtTree, expr :: GoExpr, nextId :: Int }'
);

code = code.replace(
    /translateExprImpl modNameStr recVars namedBound bound tcoIdent loopCtx isTail nextId tcoExpr@\(TcoExpr _ expr\) = case expr of/g,
    `translateExprImpl helpersRef depth modNameStr recVars namedBound bound tcoIdent loopCtx isTail nextId tcoExpr@(TcoExpr _ expr) =
  let
    liftIfNeeded mkNodeThunk =
      if depth > 10 then unsafePerformEffect do
        let fvsSet = freeVars tcoExpr
        let fvs = Array.fromFoldable fvsSet
        let helperName = "__helper_" <> show nextId
        let newNextId = nextId + 1
        let res = translateExprImpl helpersRef 0 modNameStr recVars namedBound bound Nothing [] false newNextId tcoExpr
        
        let helperExpr = if Array.length fvs == 0 
                         then GoFunc "_" (wrapInStmts [] res.stmts res.expr)
                         else Array.foldr (\\fv accFunc -> GoFunc fv accFunc) (wrapInStmts [] res.stmts res.expr) fvs
        
        Ref.modify_ (\\arr -> Array.snoc arr { identifier: helperName, expression: helperExpr }) helpersRef
        
        let callExpr = if Array.length fvs == 0
                       then GoCall (GoSelector (GoVar "gopurs_runtime") "Apply") [GoVar helperName, GoRaw "gopurs_runtime.Int(0)"]
                       else Array.foldl (\\accCall fv -> GoCall (GoSelector (GoVar "gopurs_runtime") "Apply") [accCall, GoVar fv]) (GoVar helperName) fvs
        
        pure { stmts: StmtEmpty, expr: callExpr, nextId: res.nextId }
      else mkNodeThunk unit
  in case expr of`
);

code = code.replace(/translateExprImpl modNameStr/g, 'translateExprImpl helpersRef (depth + 1) modNameStr');

code = code.replace(
  /translateExprImpl helpersRef \(depth \+ 1\) modNameStr recVars Map\.empty Map\.empty Nothing loopCtxs true 0 fn\.body/g,
  'translateExprImpl helpersRef 0 modNameStr recVars Map.empty Map.empty Nothing loopCtxs true 0 fn.body'
);
code = code.replace(
  /translateExprImpl helpersRef \(depth \+ 1\) modNameStr recVars Map\.empty Map\.empty \(Just \(sanitizeName name\)\) \[\] false 0 expr/g,
  'translateExprImpl helpersRef 0 modNameStr recVars Map.empty Map.empty (Just (sanitizeName name)) [] false 0 expr'
);
code = code.replace(
  /translateExprImpl helpersRef \(depth \+ 1\) modNameStr \[\] Map\.empty Map\.empty \(Just \(sanitizeName name\)\) \[\] false 0 expr/g,
  'translateExprImpl helpersRef 0 modNameStr [] Map.empty Map.empty (Just (sanitizeName name)) [] false 0 expr'
);

let parts = code.split('translateExprImpl helpersRef depth modNameStr');
let top = parts[0];
let bottom = parts[1];

bottom = bottom.replace(/LetRec lvl binds body ->/g, 'LetRec lvl binds body -> liftIfNeeded \\_ ->');
bottom = bottom.replace(/Abs arg body ->/g, 'Abs arg body -> liftIfNeeded \\_ ->');
bottom = bottom.replace(/UncurriedAbs args body ->/g, 'UncurriedAbs args body -> liftIfNeeded \\_ ->');
bottom = bottom.replace(/UncurriedEffectAbs args body ->/g, 'UncurriedEffectAbs args body -> liftIfNeeded \\_ ->');

code = top + 'translateExprImpl helpersRef depth modNameStr' + bottom;

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
console.log('Patched CodeGen.purs successfully');
