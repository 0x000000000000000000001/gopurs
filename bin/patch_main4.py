import re

with open('/Users/0x1/Documents/htdocs/gopurs/src/Main.purs', 'r') as f:
    content = f.read()

content = content.replace('import PureScript.Backend.Optimizer.CoreFn (Module(..), Ann, importName, Bind(..), Binding(..), ExprType(..), Ident(..))', 'import PureScript.Backend.Optimizer.CoreFn (Module(..), Ann(..), importName, Bind(..), Binding(..), ExprType(..), Ident(..))')

content = content.replace('import Data.Tuple.Nested ((/\\))', 'import Data.Tuple.Nested ((/\\))\nimport Data.Tuple (Tuple(..))')

with open('/Users/0x1/Documents/htdocs/gopurs/src/Main.purs', 'w') as f:
    f.write(content)
