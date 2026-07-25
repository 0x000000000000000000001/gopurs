import re

with open('/Users/0x1/Documents/htdocs/gopurs/src/Main.purs', 'r') as f:
    content = f.read()

# Add a function to build the global types
new_code = """
buildGlobalTypes :: Array (Module Ann) -> Map.Map String ExprType
buildGlobalTypes modules = Array.foldl (\\acc (Module m) ->
  let modName = unwrap m.name
      processBind acc' (NonRec (Binding _ (Ident name) expr)) =
        case (getExprAnn expr) of
          Ann ann -> case ann.type of
            Just t -> Map.insert (modName <> "." <> name) t acc'
            Nothing -> acc'
      processBind acc' (Rec bindings) = Array.foldl (\\a (Binding _ (Ident name) expr) ->
        case (getExprAnn expr) of
          Ann ann -> case ann.type of
            Just t -> Map.insert (modName <> "." <> name) t a
            Nothing -> a
        ) acc' bindings
  in Array.foldl processBind acc m.decls
  ) Map.empty modules

hasTypeVariables :: ExprType -> Boolean
hasTypeVariables (TypeVar v) = String.take 1 v == String.toLower (String.take 1 v) && v /= "gopurs_runtime.Value"
hasTypeVariables (Func args ret) = Array.any hasTypeVariables args || hasTypeVariables ret
hasTypeVariables (Array t) = hasTypeVariables t
hasTypeVariables (Record props) = Array.any (\\(Tuple _ v) -> hasTypeVariables v) props
hasTypeVariables (App t1 t2) = hasTypeVariables t1 || hasTypeVariables t2
hasTypeVariables (TypeLevelString _) = false
hasTypeVariables (TypeLevelInt _) = false
hasTypeVariables Int = false
hasTypeVariables String = false
hasTypeVariables Char = false
hasTypeVariables Number = false
hasTypeVariables Boolean = false
hasTypeVariables (ADT _) = false
"""

if 'buildGlobalTypes' not in content:
    # insert before readCoreFnModule
    content = content.replace('readCoreFnModule :: String -> Aff (Maybe (Module Ann))', new_code + '\nreadCoreFnModule :: String -> Aff (Maybe (Module Ann))')

# Add missing imports for the new code
content = content.replace('import PureScript.Backend.Optimizer.CoreFn (Module(..), Ann, importName)', 'import PureScript.Backend.Optimizer.CoreFn (Module(..), Ann, importName, Bind(..), Binding(..), ExprType(..), Ident(..))\nimport Data.String as String\nimport Gopurs.Monomorphize (getExprAnn)')

# Filter instantiations
old_inst = 'let instantiations = foldl collectInstantiations Map.empty finalModules'
new_inst = """let globalTypes = buildGlobalTypes (Array.fromFoldable finalModules)
  let rawInstantiations = foldl collectInstantiations Map.empty finalModules
  let instantiations = Map.filterKeys (\\k -> case Map.lookup k globalTypes of
                                            Just t -> hasTypeVariables t
                                            Nothing -> false) rawInstantiations"""

content = content.replace(old_inst, new_inst)

with open('/Users/0x1/Documents/htdocs/gopurs/src/Main.purs', 'w') as f:
    f.write(content)
