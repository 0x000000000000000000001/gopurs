import re

with open('/Users/0x1/Documents/htdocs/gopurs/src/Gopurs/CodeGen.purs', 'r') as f:
    content = f.read()

# Add InstantiationMap to imports
if 'type InstantiationMap' not in content:
    content = content.replace('import Gopurs.GoAst', 'import Gopurs.Monomorphize (InstantiationMap)\nimport Gopurs.GoAst')

# Fix translate signature
content = content.replace(
    'translate :: Set.Set String -> Map.Map String (Array ExprType) -> Array (Array String) -> BackendModule -> String\ntranslate elidedCtors ctorTypes importsArray mod =',
    'translate :: Set.Set String -> Map.Map String (Array ExprType) -> InstantiationMap -> Array (Array String) -> BackendModule -> String\ntranslate elidedCtors ctorTypes instantiations importsArray mod ='
)

with open('/Users/0x1/Documents/htdocs/gopurs/src/Gopurs/CodeGen.purs', 'w') as f:
    f.write(content)
