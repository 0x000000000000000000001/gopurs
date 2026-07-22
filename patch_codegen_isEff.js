import fs from 'fs';
let content = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');

const isEffectNodeSrc = `
isEffectNode :: forall a. BackendSyntax a -> Boolean
isEffectNode = case _ of
  EffectBind _ _ _ _ -> true
  EffectPure _ -> true
  EffectDefer _ -> true
  PrimEffect _ -> true
  UncurriedEffectApp _ _ -> true
  _ -> false

executeIfOpaque :: forall a. BackendSyntax a -> GoExpr -> GoExpr
executeIfOpaque expr goExpr =
  if isEffectNode expr then goExpr
  else GoCall (GoSelector (GoVar "gopurs_runtime") "Apply") [ goExpr, GoRaw "gopurs_runtime.Value{}" ]
`;

// Insert right before translateExprImpl ::
content = content.replace(
  'translateExprImpl :: Ref (Array GoDecl) -> Int -> String -> Array String -> Map String String -> Map String String -> Maybe String -> Array { ident :: String, params :: Array String, loopParams :: Array String } -> Boolean -> Int -> TcoExpr -> { stmts :: StmtTree, expr :: GoExpr, nextId :: Int }',
  isEffectNodeSrc + '\ntranslateExprImpl :: Ref (Array GoDecl) -> Int -> String -> Array String -> Map String String -> Map String String -> Maybe String -> Array { ident :: String, params :: Array String, loopParams :: Array String } -> Boolean -> Int -> TcoExpr -> { stmts :: StmtTree, expr :: GoExpr, nextId :: Int }'
);

fs.writeFileSync('src/Gopurs/CodeGen.purs', content);
