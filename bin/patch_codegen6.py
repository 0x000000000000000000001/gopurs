import re

with open('/Users/0x1/Documents/htdocs/gopurs/src/Gopurs/CodeGen.purs', 'r') as f:
    content = f.read()

newCode = """    getExprType :: TcoExpr -> ExprType
    getExprType (TcoExpr _ (Typed ty _)) = ty
    getExprType _ = Any -- fallback
"""

content = content.replace("    getExprType :: TcoExpr -> ExprType\n    getExprType (TcoExpr _ (Typed ty _)) = ty\n    getExprType _ = TypeApp (TypeVar \"Any\") (TypeVar \"Any\") -- fallback\n", newCode)

with open('/Users/0x1/Documents/htdocs/gopurs/src/Gopurs/CodeGen.purs', 'w') as f:
    f.write(content)
