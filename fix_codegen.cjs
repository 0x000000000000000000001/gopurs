const fs = require('fs');
let content = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');

const targetStructDecls = `        structDecls = Array.concatMap (\\adtType ->
          case adtType of
            ADT names args ->
              if Array.length names == 0 then [] else
                let
                  modStr = String.joinWith "_" (fromMaybe [] (Array.init names))
                  declName = fromMaybe "" (Array.last names)
                in
                  if modStr == modNameStr then
                    case Array.find (\\d -> d.typeName == declName) mod.dataDecls of
                      Just decl ->
                        let
                          interfaceName = goTypeToStr (TypeInterface ("ADT_" <> String.joinWith "_" names <> (if Array.length args == 0 then "" else "_" <> String.joinWith "_" (map mangleType args))))
                          interfaceDecl = "type " <> interfaceName <> " interface {\\n\\tis" <> interfaceName <> "()\\n}\\n"
                          ctorDecls = map (\\ctor ->
                            let
                              -- substitute type vars in ctor field types using the instantiation args!
                              subst = Array.foldl (\\acc (Tuple v t) -> Map.insert v t acc) Map.empty (Array.zip decl.typeVars args)
                              monomorphizedFields = map (\\t -> substituteExprType subst t) ctor.fieldTypes
                              fieldsStr = Array.mapWithIndex (\\i ty -> "V" <> show i <> " " <> goTypeToStr (exprTypeToGoType ty)) monomorphizedFields
                              structName = "Constructor_" <> sanitizeName ctor.constructorName <> (if Array.length args == 0 then "" else "_" <> String.joinWith "_" (map mangleType args))
                            in
                              "type " <> structName <> " struct {\\n\\t" <> String.joinWith "\\n\\t" fieldsStr <> "\\n}\\n" <>
                              "func (v *" <> structName <> ") is" <> interfaceName <> "() {}\\n"
                          ) decl.constructors
                        in Array.cons interfaceDecl ctorDecls
                      Nothing -> []
                  else []
            _ -> []
        ) (Array.fromFoldable adtTypes)`;

const replacementStructDecls = `        -- group by constructor struct name
        ctorInstances = foldl (\\acc adtType ->
            case adtType of
              ADT names args ->
                if Array.length names == 0 then acc else
                  let
                    modStr = String.joinWith "_" (fromMaybe [] (Array.init names))
                    declName = fromMaybe "" (Array.last names)
                  in
                    if modStr == modNameStr then
                      case Array.find (\\d -> d.typeName == declName) mod.dataDecls of
                        Just decl ->
                          let
                            interfaceName = goTypeToStr (TypeInterface ("ADT_" <> String.joinWith "_" names <> (if Array.length args == 0 then "" else "_" <> String.joinWith "_" (map mangleType args))))
                          in foldl (\\acc' ctor -> 
                            let
                              subst = Array.foldl (\\substAcc (Tuple v t) -> Map.insert v t substAcc) Map.empty (Array.zip decl.typeVars args)
                              monomorphizedFields = map (\\t -> substituteExprType subst t) ctor.fieldTypes
                              goFieldTypes = map exprTypeToGoType monomorphizedFields
                              structName = "Constructor_" <> sanitizeName ctor.constructorName <> (if Array.length goFieldTypes == 0 then "" else "_" <> String.replaceAll (Pattern "*") (Replacement "ptr") (String.replaceAll (Pattern "[]") (Replacement "arr") (String.replaceAll (Pattern " ") (Replacement "_") (String.joinWith "_" (map goTypeToStr goFieldTypes)))))
                            in
                              Map.insertWith (\\old new -> { interfaceNames: Array.concat [old.interfaceNames, new.interfaceNames], goFieldTypes: old.goFieldTypes }) structName { interfaceNames: [interfaceName], goFieldTypes } acc'
                          ) acc decl.constructors
                        Nothing -> acc
                    else acc
              _ -> acc
          ) Map.empty (Array.fromFoldable adtTypes)

        interfaceDecls = Array.concatMap (\\adtType ->
            case adtType of
              ADT names args ->
                if Array.length names == 0 then [] else
                  let
                    modStr = String.joinWith "_" (fromMaybe [] (Array.init names))
                    declName = fromMaybe "" (Array.last names)
                  in
                    if modStr == modNameStr then
                      let
                        interfaceName = goTypeToStr (TypeInterface ("ADT_" <> String.joinWith "_" names <> (if Array.length args == 0 then "" else "_" <> String.joinWith "_" (map mangleType args))))
                      in [ "type " <> interfaceName <> " interface {\\n\\tis" <> interfaceName <> "()\\n}\\n" ]
                    else []
              _ -> []
          ) (Array.fromFoldable adtTypes)

        structDecls = interfaceDecls <> ((Map.toUnfoldable ctorInstances :: Array (Tuple String { interfaceNames :: Array String, goFieldTypes :: Array GoType })) # Array.concatMap (\\(Tuple structName { interfaceNames, goFieldTypes }) -> 
            let
              fieldsStr = Array.mapWithIndex (\\i ty -> "V" <> show i <> " " <> goTypeToStr ty) goFieldTypes
              structDecl = "type " <> structName <> " struct {\\n\\t" <> String.joinWith "\\n\\t" fieldsStr <> "\\n}\\n"
              methods = Array.nub interfaceNames # map (\\interfaceName -> 
                  "func (v *" <> structName <> ") is" <> interfaceName <> "() {}\\n"
                )
            in
              Array.cons structDecl methods
          ))`;

const targetCtorSaturated = `      CtorSaturated _ name _ _ _ ->
        let
          accProps = foldl
            ( \\acc expr ->
                let
                  res = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars moduleArities bound (Just currentType) loopCtx isTail false acc.nextId expr
                in
                  { stmts: Array.snoc acc.stmts res.stmts
                  , exprs: Array.snoc acc.exprs res.expr
                  , nextId: res.nextId
                  }
            )
            { stmts: [], exprs: [], nextId }
            args

          baseStructName = String.replaceAll (Pattern ".") (Replacement "_") (unwrap name)
          
          structName = "Constructor_" <> sanitizeName name <> (if Array.length typeArgs == 0 then "" else "_" <> String.joinWith "_" (map mangleType typeArgs))
        in
          { stmts: accProps.stmts, expr: GoConstructor (hashString baseStructName) structName accProps.exprs, exprType: TypeValue, nextId: accProps.nextId }`;

const replacementCtorSaturated = `      CtorSaturated _ name _ _ _ ->
        let
          accProps = foldl
            ( \\acc expr ->
                let
                  res = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars moduleArities bound (Just currentType) loopCtx isTail false acc.nextId expr
                in
                  { stmts: Array.snoc acc.stmts res.stmts
                  , exprs: Array.snoc acc.exprs res.expr
                  , goFieldTypes: Array.snoc acc.goFieldTypes res.exprType
                  , nextId: res.nextId
                  }
            )
            { stmts: [], exprs: [], goFieldTypes: [], nextId }
            args

          baseStructName = String.replaceAll (Pattern ".") (Replacement "_") (unwrap name)
          
          structName = "Constructor_" <> sanitizeName name <> (if Array.length accProps.goFieldTypes == 0 then "" else "_" <> String.replaceAll (Pattern "*") (Replacement "ptr") (String.replaceAll (Pattern "[]") (Replacement "arr") (String.replaceAll (Pattern " ") (Replacement "_") (String.joinWith "_" (map goTypeToStr accProps.goFieldTypes)))))
        in
          { stmts: accProps.stmts, expr: GoConstructor (hashString baseStructName) structName accProps.exprs, exprType: TypeValue, nextId: accProps.nextId }`;

content = content.replace(targetStructDecls, replacementStructDecls);
content = content.replace(targetCtorSaturated, replacementCtorSaturated);

fs.writeFileSync('src/Gopurs/CodeGen.purs', content, 'utf8');
