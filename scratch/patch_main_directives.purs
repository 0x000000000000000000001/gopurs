import PureScript.Backend.Optimizer.Semantics (InlineDirectiveMap, EvalRef(..))

  let extendedDirectives = Map.foldrWithIndex (\origName hashInfos accDir ->
        let
          parts = String.split (Pattern ".") origName
          mbIdent = Array.last parts
          mbMod = if Array.length parts > 1 then Just (ModuleName (String.joinWith "." (Array.dropEnd 1 parts))) else Nothing
        in case mbIdent, mbMod of
          Just identStr, Just modName ->
            let
              origRef = ExternRef (Qualified (Just modName) (Ident identStr))
            in case Map.lookup origRef accDir of
              Just dirValue ->
                Map.foldrWithIndex (\hash _ accDir' ->
                  Map.insert (LocalRef (Ident (identStr <> "__" <> hash))) dirValue accDir'
                ) accDir hashInfos
              Nothing -> accDir
          _, _ -> accDir
      ) directives instantiations
