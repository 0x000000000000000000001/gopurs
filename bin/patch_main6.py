import re

with open('/Users/0x1/Documents/htdocs/gopurs/src/Main.purs', 'r') as f:
    content = f.read()

lines = content.split('\n')
new_lines = []
for line in lines:
    if 'hasTypeVariables (App t1 t2)' in line: continue
    if 'hasTypeVariables (TypeLevelString _)' in line: continue
    if 'hasTypeVariables (TypeLevelInt _)' in line: continue
    new_lines.append(line)

# Also Any should return false
new_lines = '\n'.join(new_lines)
new_lines = new_lines.replace('hasTypeVariables (ADT _) = false', 'hasTypeVariables (ADT _) = false\nhasTypeVariables Any = false')

with open('/Users/0x1/Documents/htdocs/gopurs/src/Main.purs', 'w') as f:
    f.write(new_lines)
