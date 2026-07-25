import re

with open('/Users/0x1/Documents/htdocs/gopurs/src/Main.purs', 'r') as f:
    content = f.read()

old_call = "let goFile = translate elidedCtors ctorTypes instantiations importsArray backendMod"
new_call = "let goFile = translate elidedCtors ctorTypes globalTypes instantiations importsArray backendMod"
content = content.replace(old_call, new_call)

with open('/Users/0x1/Documents/htdocs/gopurs/src/Main.purs', 'w') as f:
    f.write(content)
