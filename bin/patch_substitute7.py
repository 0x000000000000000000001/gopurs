import re

with open('/Users/0x1/Documents/htdocs/gopurs/src/Gopurs/Monomorphize/Substitute.purs', 'r') as f:
    content = f.read()

content = content.replace("import PureScript.Backend.Optimizer.Syntax (BackendSyntax(..), BackendEffect(..), BackendOperator(..), BackendOperator1(..), BackendOperator2(..))", "import PureScript.Backend.Optimizer.Syntax (BackendSyntax(..), BackendEffect(..), BackendOperator(..), BackendOperator1(..), BackendOperator2(..), Pair(..))")

with open('/Users/0x1/Documents/htdocs/gopurs/src/Gopurs/Monomorphize/Substitute.purs', 'w') as f:
    f.write(content)
