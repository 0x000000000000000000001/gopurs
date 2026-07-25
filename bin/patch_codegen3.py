import re

with open('/Users/0x1/Documents/htdocs/gopurs/src/Gopurs/CodeGen.purs', 'r') as f:
    content = f.read()

# Add imports
if 'mangleType' not in content[:1000]:
    content = content.replace('import Gopurs.GoAst', 'import Gopurs.GoAst (mangleType)\nimport Gopurs.GoAst')
if 'getExprType' not in content[:1000]:
    content = content.replace('import PureScript.Backend.Optimizer.CoreFn (', 'import PureScript.Backend.Optimizer.CoreFn (\n  getExprType,\n  setTcoExprType,')

with open('/Users/0x1/Documents/htdocs/gopurs/src/Gopurs/CodeGen.purs', 'w') as f:
    f.write(content)
