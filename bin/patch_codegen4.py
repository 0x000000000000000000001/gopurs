import re

with open('/Users/0x1/Documents/htdocs/gopurs/src/Gopurs/CodeGen.purs', 'r') as f:
    content = f.read()

# Remove the bad import
content = content.replace('import Gopurs.GoAst (mangleType)\n', '')

newCode = """    mangleType :: ExprType -> String
    mangleType t = String.replaceAll (Pattern ".") (Replacement "_") (printGoType (exprTypeToGoType t))

    tcoBindingsExpanded = map
"""

content = content.replace("    tcoBindingsExpanded = map\n", newCode)

with open('/Users/0x1/Documents/htdocs/gopurs/src/Gopurs/CodeGen.purs', 'w') as f:
    f.write(content)
