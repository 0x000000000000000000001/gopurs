module PatchMain where

    finalModulesWithClassDecls = map (\(Module m) ->
      let newDecls = map (\c ->
            let superclassFields = Array.mapWithIndex (\i super ->
                  let superName = fromMaybe "" (Array.last (fst super))
                  in Tuple (superName <> show i) Any
                ) c.superclasses
                methodFields = c.methods
                allFields = Array.sortBy (comparing fst) (superclassFields <> methodFields)
                fieldTypes = map snd allFields
            in { name: c.name, vars: c.vars, constructors: [{ name: c.name, fields: fieldTypes }] }
          ) m.classDecls
      in Module (m { dataDecls = m.dataDecls <> newDecls })
    ) finalModules
