import re

with open('/Users/0x1/Documents/htdocs/gopurs/src/Gopurs/CodeGen.purs', 'r') as f:
    content = f.read()

# Remove the bad imports
content = content.replace('import PureScript.Backend.Optimizer.CoreFn (\n  getExprType,\n  setTcoExprType,', 'import PureScript.Backend.Optimizer.CoreFn (')

newCode = """    mangleType :: ExprType -> String
    mangleType t = String.replaceAll (Pattern ".") (Replacement "_") (printGoType (exprTypeToGoType t))

    getExprType :: TcoExpr -> ExprType
    getExprType (TcoExpr _ (Typed ty _)) = ty
    getExprType _ = TypeApp (TypeVar "Any") (TypeVar "Any") -- fallback

    setTcoExprType :: ExprType -> TcoExpr -> TcoExpr
    setTcoExprType ty (TcoExpr a (Typed _ inner)) = TcoExpr a (Typed ty inner)
    setTcoExprType _ expr = expr

    tcoBindingsExpanded = map
"""

content = content.replace("    mangleType :: ExprType -> String\n    mangleType t = String.replaceAll (Pattern \".\") (Replacement \"_\") (printGoType (exprTypeToGoType t))\n\n    tcoBindingsExpanded = map\n", newCode)

with open('/Users/0x1/Documents/htdocs/gopurs/src/Gopurs/CodeGen.purs', 'w') as f:
    f.write(content)
