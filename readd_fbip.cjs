const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');

if (!code.includes('globalReusedVars ::')) {
  code = code.replace(
    'translate ::',
    'globalReusedVars :: Ref.Ref (Set.Set String)\nglobalReusedVars = unsafePerformEffect (Ref.new Set.empty)\n\ntranslate ::'
  );
}

code = code.replace(
  'translateDecl helpers modNameStr bound decl = case decl of',
  'translateDecl helpers modNameStr bound decl =\n  let _ = unsafePerformEffect (Ref.write Set.empty globalReusedVars)\n  in case decl of'
);

code = code.replace(
  '          isPointerAdtLeaf = Map.member baseStructName helpers.pointerAdtLeaves\n          finalExpr =',
  `          isPointerAdtLeaf = Map.member baseStructName helpers.pointerAdtLeaves
          
          deadVarOptRaw = Array.head (Array.mapMaybe (\\(Tuple _ v) -> 
              let reused = unsafePerformEffect (Ref.read globalReusedVars)
              in
              if v.goType == expectedGoType 
                 && not (Set.member v.name reused)
                 && not (Set.member v.name liveOut) 
                 && not (Set.member v.name (freeVars tcoExpr)) 
              then Just v.name 
              else Nothing
            ) (Array.fromFoldable (map (\\v -> Tuple v.name v) (Map.values bound))))
            
          deadVarOpt = deadVarOptRaw
          _ = unsafePerformEffect (case deadVarOpt of
                Just n -> Ref.modify_ (Set.insert n) globalReusedVars
                Nothing -> pure unit)
                
          finalExpr =`
);

code = code.replace(
  '               else if isPointerAdtLeaf then GoRaw ("gopurs_runtime.Value{Type: 9, IntVal: " <> hashString (fromMaybe "" (Map.lookup baseStructName helpers.pointerAdtLeaves)) <> ", UnsafePtr: nil}")\n               else GoConstructor (hashString baseStructName) (pkgPrefix <> monoStructName) typeArgs accProps.exprs',
  `               else if isPointerAdtLeaf then GoRaw ("gopurs_runtime.Value{Type: 9, IntVal: " <> hashString (fromMaybe "" (Map.lookup baseStructName helpers.pointerAdtLeaves)) <> ", UnsafePtr: nil}")
               else case deadVarOpt of
                 Just deadVar ->
                   let
                     allocExpr = GoConstructor (hashString baseStructName) (pkgPrefix <> monoStructName) typeArgs accProps.exprs
                     typeArgsStr = if Array.length typeArgs > 0 then "[" <> String.joinWith ", " (map goTypeToStr typeArgs) <> "]" else ""
                     mutateStmts = Array.mapWithIndex (\\idx arg -> GoMutate ("(*" <> pkgPrefix <> monoStructName <> typeArgsStr <> ")(" <> deadVar <> ".UnsafePtr).V" <> show idx) arg) accProps.exprs
                     mutateBlock = GoBlock (mutateStmts <> [ GoReturn (GoVar deadVar) ])
                     allocBlock = GoBlock [ GoReturn allocExpr ]
                     
                     ifCond = GoBinOp "&&" (GoBinOp "!=" (GoRaw (deadVar <> ".UnsafePtr")) (GoRaw "nil")) (GoBinOp "==" (GoRaw ("(*struct{Rc uint32})(" <> deadVar <> ".UnsafePtr).Rc")) (GoInt 1))
                     ifStmt = GoIfElse ifCond [ mutateBlock ] [ allocBlock ]
                   in
                     GoRaw ("func() gopurs_runtime.Value {\\n" <> printGoExpr ifStmt <> "\\n}()")
                 Nothing -> GoConstructor (hashString baseStructName) (pkgPrefix <> monoStructName) typeArgs accProps.exprs`
);

fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
