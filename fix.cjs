const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');

const liftIfNeededCode = `
  let
    liftIfNeeded mkNode =
      if depth > 100 then unsafePerformEffect do
        let allVarsSet = freeVars tcoExpr
        let availableVars = Set.union (Set.fromFoldable (Map.keys bound)) (Set.fromFoldable (Map.keys namedBound))
        let fvsSet = Set.intersection allVarsSet availableVars
        let fvs = Array.fromFoldable fvsSet
        helperId <- Ref.modify' (\\id -> { state: id + 1, value: id }) globalHelperId
        let helperName = "__helper_" <> show helperId
        let newNextId = nextId + 1
        let res = translateExprImpl helpersRef 0 modNameStr recVars namedBound bound Nothing [] false newNextId tcoExpr
        Ref.modify_ (\\arr -> Array.snoc arr (GoHelper helperName fvs (wrapInStmts [] res.stmts res.expr))) helpersRef
        pure { stmts: StmtEmpty, expr: GoHelperCall helperName (map GoVar fvs), nextId: res.nextId }
      else mkNode
`;

code = code.replace(
  'translateExprImpl modNameStr recVars namedBound bound tcoIdent loopCtx isTail nextId tcoExpr@(TcoExpr _ expr) = case expr of',
  'translateExprImpl helpersRef depth modNameStr recVars namedBound bound tcoIdent loopCtx isTail nextId tcoExpr@(TcoExpr _ expr) = case expr of\n' + liftIfNeededCode
);

code = code.replace(
  'translateExprImpl :: String -> Array String -> Map String String -> Map String String -> Maybe String -> Array { ident :: String, params :: Array String } -> Boolean -> Int -> TcoExpr -> { stmts :: StmtTree, expr :: GoExpr, nextId :: Int }',
  'translateExprImpl :: Ref (Array GoDecl) -> Int -> String -> Array String -> Map String String -> Map String String -> Maybe String -> Array { ident :: String, params :: Array String } -> Boolean -> Int -> TcoExpr -> { stmts :: StmtTree, expr :: GoExpr, nextId :: Int }'
);

code = code.replace(
  'Abs args body ->',
  'Abs args body -> liftIfNeeded $'
);

code = code.replace(
  'UncurriedAbs args body ->',
  'UncurriedAbs args body -> liftIfNeeded $'
);

code = code.replace(
  'UncurriedEffectAbs args body ->',
  'UncurriedEffectAbs args body -> liftIfNeeded $'
);

code = code.replace(
  'LetRec lvl bindings body ->',
  'LetRec lvl bindings body -> liftIfNeeded $'
);

code = code.replace(
  'let res = translateExprImpl modNameStr',
  'let res = translateExprImpl helpersRef 0 modNameStr'
);

fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
