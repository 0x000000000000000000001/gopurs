const fs = require('fs');
let content = fs.readFileSync('/Users/0x1/Documents/htdocs/gopurs/src/Gopurs/CodeGen.purs', 'utf8');

// 1. signature
content = content.replace(
  'translate :: Set.Set String -> Map.Map String (Array ExprType) -> Array (Array String) -> BackendModule -> String',
  'translate :: Set.Set String -> Map.Map String (Array ExprType) -> Map.Map String (Set.Set ExprType) -> Array (Array String) -> BackendModule -> String'
);

// 2. arguments
content = content.replace(
  'translate elidedCtors ctorTypes importsArray (BackendModule mod) =',
  'translate elidedCtors ctorTypes instantiations importsArray (BackendModule mod) ='
);

// 3. tcoBindingsExpanded
let tcoBindingsTarget = '      (Tuple [] [])\\n      mod.bindings\\n';
let tcoBindingsExpanded = `      (Tuple [] [])
      mod.bindings

    tcoBindingsExpanded = map
      (\\group -> group { bindings = Array.concatMap expandBind group.bindings })
      tcoBindings

    expandBind :: Tuple Ident TcoExpr -> Array (Tuple Ident TcoExpr)
    expandBind (Tuple id@(Ident name) val) =
      let qual = modNameStr <> "." <> name
      in case Map.lookup qual instantiations of
           Just concretes ->
             let genericType = getExprType val
                 concreteArr = Set.toUnfoldable concretes :: Array ExprType
             in Tuple id val \`Array.cons\` map (\\concrete ->
                  let mangledName = name <> "__" <> mangleType concrete
                  in Tuple (Ident mangledName) (setTcoExprType concrete val)
                ) concreteArr
           Nothing -> [ Tuple id val ]
`;
content = content.replace(tcoBindingsTarget, tcoBindingsExpanded);

// 4. replace usage of tcoBindings with tcoBindingsExpanded
content = content.replace(
  '    Tuple foreignDecls bindsExpanded = \\n      if Array.null tcoBindings then\\n        Tuple [] []\\n      else\\n        let\\n          foreigns = Array.filter (\\\\d -> case d of\\n            GoFuncDecl _ _ -> true\\n            _ -> false\\n          ) (processBindingGroup foreignBinds false)\\n        in Tuple foreigns (Array.concatMap (\\\\b -> processBindingGroup b.bindings b.recursive) tcoBindings)',
  '    Tuple foreignDecls bindsExpanded = \n      if Array.null tcoBindingsExpanded then\n        Tuple [] []\n      else\n        let\n          foreigns = Array.filter (\\d -> case d of\n            GoFuncDecl _ _ -> true\n            _ -> false\n          ) (processBindingGroup foreignBinds false)\n        in Tuple foreigns (Array.concatMap (\\b -> processBindingGroup b.bindings b.recursive) tcoBindingsExpanded)'
);

// 5. mbDirectCall
let mbDirectCallTarget = `                mbDirectCall = case getVar (unwrapTcoExpr flatFn) of
                  Just { mbMod, name } ->
                    let
                      isLocal = map (String.replaceAll (Pattern ".") (Replacement "_") <<< unwrap) mbMod == Just modNameStr || mbMod == Nothing
                      modPrefix = case mbMod of
                        Just (ModuleName mod) | not isLocal -> "pkg_" <> String.replaceAll (Pattern ".") (Replacement "_") mod <> "."
                        _ -> ""
                      fromModuleArities = if isLocal then Map.lookup name moduleArities else Nothing
                      fromTypeSig = case getFuncType (unwrapExpr flatFn) of
                        Just { fArgs, fRet } ->
                          Just { fullName: modPrefix <> "Call_" <> sanitizeName name, fArgs: map exprTypeToGoType fArgs, fRet: exprTypeToGoType fRet, arity: Array.length fArgs }
                        Nothing -> Nothing
                      
                      entry = case fromTypeSig of
                        Just e -> Just e
                        Nothing -> fromModuleArities
                    in
                      case entry of
                        Just e | e.arity == Array.length args ->
                          Just { fullName: e.fullName, goTypes: e.fArgs }
                        _ -> Nothing
                  Nothing -> Nothing`;

let mbDirectCallReplacement = `                mbDirectCall = case getVar (unwrapTcoExpr flatFn) of
                  Just { mbMod, name } ->
                    let
                      qualNameForLookup = case mbMod of
                        Just (ModuleName mod) -> mod <> "." <> name
                        Nothing -> modNameStr <> "." <> name
                      mangledName = case Map.lookup qualNameForLookup instantiations of
                        Just _ -> name <> "__" <> mangleType (getExprType flatFn)
                        Nothing -> name
                      isLocal = map (String.replaceAll (Pattern ".") (Replacement "_") <<< unwrap) mbMod == Just modNameStr || mbMod == Nothing
                      modPrefix = case mbMod of
                        Just (ModuleName mod) | not isLocal -> "pkg_" <> String.replaceAll (Pattern ".") (Replacement "_") mod <> "."
                        _ -> ""
                      fromModuleArities = if isLocal then Map.lookup mangledName moduleArities else Nothing
                      fromTypeSig = case getFuncType (unwrapExpr flatFn) of
                        Just { fArgs, fRet } ->
                          Just { fullName: modPrefix <> "Call_" <> sanitizeName mangledName, fArgs: map exprTypeToGoType fArgs, fRet: exprTypeToGoType fRet, arity: Array.length fArgs }
                        Nothing -> Nothing
                      
                      entry = case fromTypeSig of
                        Just e -> Just e
                        Nothing -> fromModuleArities
                    in
                      case entry of
                        Just e | e.arity == Array.length args ->
                          Just { fullName: e.fullName, goTypes: e.fArgs }
                        _ -> Nothing
                  Nothing -> Nothing`;

content = content.replace(mbDirectCallTarget, mbDirectCallReplacement);

fs.writeFileSync('/Users/0x1/Documents/htdocs/gopurs/src/Gopurs/CodeGen.purs', content);
