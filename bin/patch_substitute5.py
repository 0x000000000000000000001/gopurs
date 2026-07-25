import re

with open('/Users/0x1/Documents/htdocs/gopurs/src/Gopurs/Monomorphize/Substitute.purs', 'r') as f:
    content = f.read()

content = content.replace("import PureScript.Backend.Optimizer.Codegen.Tco (BackendEffect(..), BackendOperator1(..), BackendOperator2(..), Literal(..))\n", "")
content = content.replace("import PureScript.Backend.Optimizer.Syntax (BackendSyntax(..))", "import PureScript.Backend.Optimizer.Syntax (BackendSyntax(..), BackendEffect(..), BackendOperator1(..), BackendOperator2(..))\nimport PureScript.Backend.Optimizer.CoreFn (Literal(..), ExprType(..), Ident(..), Qualified(..))")

with open('/Users/0x1/Documents/htdocs/gopurs/src/Gopurs/Monomorphize/Substitute.purs', 'w') as f:
    f.write(content)
