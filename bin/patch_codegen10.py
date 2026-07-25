import re

with open('/Users/0x1/Documents/htdocs/gopurs/src/Gopurs/CodeGen.purs', 'r') as f:
    content = f.read()

content = content.replace(
    "translate elidedCtors ctorTypes instantiations importsArray mod =",
    "translate elidedCtors ctorTypes globalTypes instantiations importsArray mod ="
)

with open('/Users/0x1/Documents/htdocs/gopurs/src/Gopurs/CodeGen.purs', 'w') as f:
    f.write(content)
