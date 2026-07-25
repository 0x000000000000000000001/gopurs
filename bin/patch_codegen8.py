import re

with open('/Users/0x1/Documents/htdocs/gopurs/src/Gopurs/CodeGen.purs', 'r') as f:
    content = f.read()

# Replace expandBind again
old_expand = """    expandBind :: Tuple Ident TcoExpr -> Array (Tuple Ident TcoExpr)
    expandBind (Tuple id@(Ident name) val) =
      let qual = modNameStr <> "." <> name
          instsMap = map Set.toUnfoldable instantiations
          baseVal = substituteAst instsMap mangleType getExprType val
      in case Map.lookup qual instantiations of
           Just concretes ->
             let genericType = getExprType val
                 concreteArr = Set.toUnfoldable concretes :: Array ExprType
             in Tuple id baseVal `Array.cons` map (\\concrete ->
                  let mangledName = name <> "__" <> mangleType concrete
                      subst = unify genericType concrete Map.empty
                      mangledVal = mapTcoExprTypes (substituteExprType subst) val
                      mangledVal' = substituteAst instsMap mangleType getExprType mangledVal
                  in Tuple (Ident mangledName) (setTcoExprType concrete mangledVal')
                ) concreteArr
           Nothing -> [ Tuple id baseVal ]"""

new_expand = """    expandBind :: Tuple Ident TcoExpr -> Array (Tuple Ident TcoExpr)
    expandBind (Tuple id@(Ident name) val) =
      let qual = modNameStr <> "." <> name
          instsMap = map Set.toUnfoldable instantiations
          baseVal = substituteAst instsMap mangleType val
      in case Map.lookup qual instantiations of
           Just concretes ->
             let genericType = getExprType val
                 concreteArr = Set.toUnfoldable concretes :: Array ExprType
             in Tuple id baseVal `Array.cons` Array.mapMaybe (\\concrete ->
                  let subst = unify genericType concrete Map.empty
                  in if Map.isEmpty subst then Nothing else Just $
                       let mangledName = name <> "__" <> mangleType concrete
                           mangledVal = mapTcoExprTypes (substituteExprType subst) val
                           mangledVal' = substituteAst instsMap mangleType mangledVal
                       in Tuple (Ident mangledName) (setTcoExprType concrete mangledVal')
                ) concreteArr
           Nothing -> [ Tuple id baseVal ]"""

content = content.replace(old_expand, new_expand)

# Also fix the Substitute import because we changed the signature of substituteAst to remove getExprType
content = content.replace("substituteAst instsMap mangleType getExprType val", "substituteAst instsMap mangleType val")
content = content.replace("substituteAst instsMap mangleType getExprType mangledVal", "substituteAst instsMap mangleType mangledVal")

with open('/Users/0x1/Documents/htdocs/gopurs/src/Gopurs/CodeGen.purs', 'w') as f:
    f.write(content)
