import re

with open('/Users/0x1/Documents/htdocs/gopurs/src/Gopurs/Monomorphize/Substitute.purs', 'r') as f:
    content = f.read()

# Remove the bad imports from the middle
bad_imports = """import Gopurs.Monomorphize (InstantiationMap)
import PureScript.Backend.Optimizer.CoreFn (Ident(..), Qualified(..))
import Data.Newtype (unwrap)
import PureScript.Backend.Optimizer.Codegen.Tco (BackendEffect(..), BackendOperator1(..), BackendOperator2(..), Literal(..))
"""
content = content.replace(bad_imports, "")

# Add them to the top
if 'import Gopurs.Monomorphize' not in content[:500]:
    content = content.replace("import PureScript.Backend.Optimizer.Syntax (BackendSyntax(..))", "import PureScript.Backend.Optimizer.Syntax (BackendSyntax(..))\n" + bad_imports)

with open('/Users/0x1/Documents/htdocs/gopurs/src/Gopurs/Monomorphize/Substitute.purs', 'w') as f:
    f.write(content)
