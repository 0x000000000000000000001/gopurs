const fs = require('fs');
let content = fs.readFileSync('/Users/0x1/Documents/htdocs/gopurs/src/Gopurs/CodeGen.purs', 'utf8');

let newCode = `    tcoBindingsExpanded = map
      (\\group -> group { bindings = Array.concatMap expandBind group.bindings })
      tcoBindings

    expandBind :: Tuple Ident TcoExpr -> Array (Tuple Ident TcoExpr)
    expandBind (Tuple id@(Ident name) val) =
      let qual = modNameStr <> "." <> name
      in case Map.lookup qual instantiations of
           Just concretes ->
             let genericType = getExprType val
                 concreteArr = Set.toUnfoldable concretes :: Array ExprType
             in Tuple id val : map (\\concrete ->
                  let mangledName = name <> "__" <> mangleType concrete
                  in Tuple (Ident mangledName) (setTcoExprType concrete val)
                ) concreteArr
           Nothing -> [ Tuple id val ]

    extractFuncType :: TcoExpr -> Maybe { fArgs :: Array ExprType, fRet :: ExprType }`;

content = content.replace("    extractFuncType :: TcoExpr -> Maybe { fArgs :: Array ExprType, fRet :: ExprType }", newCode);

content = content.replace("tcoBindings\n\n    Tuple decls helpers = unsafePerformEffect do", "tcoBindingsExpanded\n\n    Tuple decls helpers = unsafePerformEffect do");

content = content.replace("            )\n            tcoBindings\n      h <- Ref.read helpersRef", "            )\n            tcoBindingsExpanded\n      h <- Ref.read helpersRef");

fs.writeFileSync('/Users/0x1/Documents/htdocs/gopurs/src/Gopurs/CodeGen.purs', content);
