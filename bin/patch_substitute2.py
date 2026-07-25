import re

with open('/Users/0x1/Documents/htdocs/gopurs/src/Gopurs/Monomorphize/Substitute.purs', 'r') as f:
    content = f.read()

content = content.replace("case mbMn of { Just mn -> unwrap mn <> \".\" <> name; Nothing -> name }", "(case mbMn of\\n        Just mn -> unwrap mn <> \".\" <> name\\n        Nothing -> name)")

# Let's fix missing imports
if 'import Data.Newtype (unwrap)' not in content:
    content = content.replace('import PureScript.Backend.Optimizer.CoreFn (Ident(..), Qualified(..))', 'import PureScript.Backend.Optimizer.CoreFn (Ident(..), Qualified(..))\nimport Data.Newtype (unwrap)\nimport PureScript.Backend.Optimizer.Codegen.Tco (BackendEffect(..), BackendOperator1(..), BackendOperator2(..), Literal(..))\n')

with open('/Users/0x1/Documents/htdocs/gopurs/src/Gopurs/Monomorphize/Substitute.purs', 'w') as f:
    f.write(content)
