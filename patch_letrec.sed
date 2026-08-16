/LetRec lvl bindings body ->/,/combinedRecVars =/c\
      LetRec lvl bindings body ->\
        let\
          allocRes = foldl\
            ( \acc (Tuple (Ident ident) val) ->\
                let\
                  oldName = localId (Just (Ident ident)) lvl\
                  gId = unsafePerformEffect do\
                    curr <- Ref.read helpersRef\
                    Ref.modify_ (\\r -> r { globalId = r.globalId + 1 }) helpersRef\
                    pure curr.globalId\
                  newName = oldName <> "_" <> show acc.nextId <> "_" <> show gId\
                  expectedGoTypeFromAst = exprTypeToGoType (unsafePerformEffect (Ref.read helpersRef)).pointerAdtPaths (unsafePerformEffect (Ref.read helpersRef)).enumAdts (unsafePerformEffect (Ref.read helpersRef)).elidedCtors modNameStr (getExprType val)\
                  -- Note: For LetRec, we MUST use expectedGoTypeFromAst because we cannot translate the body yet (mutual recursion).\
                in\
                  { newBound: Map.insert oldName { name: newName, goType: expectedGoTypeFromAst } acc.newBound, newNames: Array.snoc acc.newNames { oldName, newName }, exprType: TypeValue, nextId: acc.nextId + 1 }\
            )\
            { newBound: bound, newNames: [], exprType: TypeValue, nextId }\
            (toArray bindings)\
\
          combinedRecVars =
