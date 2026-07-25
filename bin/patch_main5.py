import re

with open('/Users/0x1/Documents/htdocs/gopurs/src/Main.purs', 'r') as f:
    content = f.read()

content = content.replace('import Data.Maybe (Maybe(..))', 'import Data.Maybe (Maybe(..))\nimport Data.Tuple (Tuple(..))')

with open('/Users/0x1/Documents/htdocs/gopurs/src/Main.purs', 'w') as f:
    f.write(content)
