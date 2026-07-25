import re

with open('/Users/0x1/Documents/htdocs/gopurs/src/Gopurs/Monomorphize/Substitute.purs', 'r') as f:
    content = f.read()

content = content.replace("import Gopurs.Monomorphize (InstantiationMap)\n", "import Data.Set as Set\n")
content = content.replace("substituteAst :: InstantiationMap", "substituteAst :: Map String (Array ExprType)")

with open('/Users/0x1/Documents/htdocs/gopurs/src/Gopurs/Monomorphize/Substitute.purs', 'w') as f:
    f.write(content)
