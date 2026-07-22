import re

with open('src/Gopurs/CodeGen.purs', 'r') as f:
    lines = f.readlines()

new_lines = []
for i, line in enumerate(lines):
    if line.startswith('translateExprImpl ::'):
        new_lines.append(line.replace('Boolean -> Int -> TcoExpr', 'Boolean -> Boolean -> Int -> TcoExpr'))
    elif line.startswith('translateExprImpl helpersRef depth'):
        new_lines.append(line.replace('isTail nextId tcoExpr', 'isTail inEffect nextId tcoExpr'))
    else:
        # replace calls:
        # translateExprImpl arg1 arg2 arg3 arg4 arg5 arg6 arg7 arg8 arg9 nextId tcoExpr
        # We know arg9 is the 9th argument after translateExprImpl.
        # But some calls span multiple lines.
        # Actually, in CodeGen.purs, almost all translateExprImpl calls are on a single line!
        
        # Let's use a regex to find all translateExprImpl
        # translateExprImpl helpersRef (depth + 1) modNameStr recVars namedBound bound Nothing [] false nextId binding
        
        # Actually, we can just find 'translateExprImpl' and manually add 'false' before the nextId parameter if we just write a good regex.
        # Wait, if we just define `translateExpr` that takes `inEffect`, and `translateExprImpl` that DOES NOT.
        pass
        
