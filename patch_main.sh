sed -i '' -e 's/accWithBinds = Array.foldl processBind acc m.decls/accWithBinds = Array.foldl processBind acc m.decls\n      accWithForeigns = Array.foldl (\\a (Tuple (Ident name) typeMb) -> case typeMb of\n        Just t -> Map.insert (modName <> "." <> name) t a\n        Nothing -> a) accWithBinds (Map.toUnfoldable m.foreign :: Array (Tuple Ident (Maybe ExprType)))\n  in\n      accWithForeigns/' src/Main.purs

sed -i '' -e 's/accWithBinds$/accWithForeigns/' src/Main.purs
