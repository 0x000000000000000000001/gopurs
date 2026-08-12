const fs = require('fs');
let code = fs.readFileSync('src/Main.purs', 'utf8');

code = code.replace(/classDeclsFields = foldl[\\s\\S]*?in case \\(if.*?\\)[\\s\\S]*?_ -> Map\.insert.*?acc'\\n\s*\) acc m\.classDecls\\n\s*\) Map\.empty finalModules/,
`classDeclsFields = foldl (\\acc (Module m) ->
      foldl (\\acc' c ->
        let superclassFields = Array.mapWithIndex (\\i super ->
                  let superName = fromMaybe "" (Array.last (fst super))
                  in Tuple (superName <> show i) Any
                ) c.superclasses
            methodFields = map (\\method -> Tuple method.ident method.type) c.methods
            allFields = Array.sortBy (comparing fst) (superclassFields <> methodFields)
            fieldsWithTypes = map (\\(Tuple name ty) -> { name: name, "type": ty }) allFields
            vars = map (\\(Tuple v _) -> v) c.vars
        in Map.insert (unwrap m.name <> "." <> c.name) { vars, fields: fieldsWithTypes } acc'
      ) acc m.classDecls
    ) Map.empty finalModules`);

fs.writeFileSync('src/Main.purs', code);
