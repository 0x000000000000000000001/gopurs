import re

with open('/Users/0x1/Documents/htdocs/gopurs/src/Gopurs/CodeGen.purs', 'r') as f:
    content = f.read()

# Add import
if 'import Gopurs.Monomorphize.Substitute' not in content:
    content = content.replace('import Gopurs.Monomorphize (InstantiationMap)', 'import Gopurs.Monomorphize (InstantiationMap)\nimport Gopurs.Monomorphize.Substitute (unify, substituteExprType, mapTcoExprTypes, substituteAst)')

# Replace expandBind
old_expand = """    expandBind :: Tuple Ident TcoExpr -> Array (Tuple Ident TcoExpr)
    expandBind (Tuple id@(Ident name) val) =
      let qual = modNameStr <> "." <> name
      in case Map.lookup qual instantiations of
           Just concretes ->
             let genericType = getExprType val
                 concreteArr = Set.toUnfoldable concretes :: Array ExprType
             in Tuple id val `Array.cons` map (\\concrete ->
                  let mangledName = name <> "__" <> mangleType concrete
                  in Tuple (Ident mangledName) (setTcoExprType concrete val)
                ) concreteArr
           Nothing -> [ Tuple id val ]"""

new_expand = """    expandBind :: Tuple Ident TcoExpr -> Array (Tuple Ident TcoExpr)
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

content = content.replace(old_expand, new_expand)

with open('/Users/0x1/Documents/htdocs/gopurs/src/Gopurs/CodeGen.purs', 'w') as f:
    f.write(content)
