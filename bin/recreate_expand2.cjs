const fs = require('fs');
let content = fs.readFileSync('/Users/0x1/Documents/htdocs/gopurs/src/Gopurs/CodeGen.purs', 'utf8');

let tcoBindingsTarget = `                  if group.recursive then
                    processBindingGroup group.bindings true
                  else
                    Array.concatMap (\\b -> processBindingGroup [b] false) group.bindings
            )
            tcoBindings`;

let tcoBindingsReplacement = `                  if group.recursive then
                    processBindingGroup group.bindings true
                  else
                    Array.concatMap (\\b -> processBindingGroup [b] false) group.bindings
            )
            tcoBindingsExpanded
            
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

content = content.replace(tcoBindingsTarget, tcoBindingsReplacement);

fs.writeFileSync('/Users/0x1/Documents/htdocs/gopurs/src/Gopurs/CodeGen.purs', content);
