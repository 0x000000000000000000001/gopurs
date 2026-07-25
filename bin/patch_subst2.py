import re

with open('/Users/0x1/Documents/htdocs/gopurs/src/Gopurs/Monomorphize/Substitute.purs', 'r') as f:
    content = f.read()

old_code = """        Just concretes ->
          let
            varTy = case mbTy of
                      Just ty -> ty
                      Nothing -> TypeVar "gopurs_runtime.Value"
            mangled = name <> "__" <> mangle varTy
          in TcoExpr a (Var (Qualified mbMn (Ident mangled)))"""

new_code = """        Just concretes ->
          let
            varTy = case mbTy of
                      Just ty -> ty
                      Nothing -> TypeVar "gopurs_runtime.Value"
          in if Array.elem varTy concretes then
               let mangled = name <> "__" <> mangle varTy
               in TcoExpr a (Var (Qualified mbMn (Ident mangled)))
             else TcoExpr a (Var (Qualified mbMn (Ident name)))"""

content = content.replace(old_code, new_code)

if 'import Data.Array as Array' not in content:
    content = content.replace('import Data.Map (Map)', 'import Data.Array as Array\nimport Data.Map (Map)')

with open('/Users/0x1/Documents/htdocs/gopurs/src/Gopurs/Monomorphize/Substitute.purs', 'w') as f:
    f.write(content)
