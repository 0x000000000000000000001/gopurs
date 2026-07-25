import re

with open('/Users/0x1/Documents/htdocs/gopurs/src/Gopurs/CodeGen.purs', 'r') as f:
    content = f.read()

old_sig = "translate :: Set.Set String -> Map.Map String (Array ExprType) -> InstantiationMap -> Array (Array String) -> BackendModule -> String"
new_sig = "translate :: Set.Set String -> Map.Map String (Array ExprType) -> Map.Map String ExprType -> InstantiationMap -> Array (Array String) -> BackendModule -> String"
content = content.replace(old_sig, new_sig)

old_def = "translate elidedCtors ctorTypes instantiations importsArray backendModule = "
new_def = "translate elidedCtors ctorTypes globalTypes instantiations importsArray backendModule = "
content = content.replace(old_def, new_def)

old_expand = """             let genericType = getExprType val"""
new_expand = """             let genericType = fromMaybe (getExprType val) (Map.lookup qual globalTypes)"""
content = content.replace(old_expand, new_expand)

with open('/Users/0x1/Documents/htdocs/gopurs/src/Gopurs/CodeGen.purs', 'w') as f:
    f.write(content)
