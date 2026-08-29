module Main where

import Prelude

import Effect (Effect)
import Effect.Console as Console
import Effect.Class (liftEffect)
import Effect.Aff (Aff, launchAff_, attempt)
import Node.FS.Aff as FS
import Node.FS.Stats as Stats
import Node.Encoding (Encoding(..))
import Node.Process as Process
import Gopurs.CodeGen as CodeGen
import Data.Argonaut.Parser (jsonParser)
import Data.Either (Either(..))
import Data.Bifunctor (lmap)
import Data.Argonaut.Decode.Error (printJsonDecodeError)
import Data.Array as Array
import Partial.Unsafe (unsafePartial)
import Data.Tuple (Tuple(..), fst, snd)
import Data.List as List
import Data.List (List)
import Data.Traversable (traverse)
import Data.Maybe (Maybe(..), isJust, fromMaybe)
import Data.Map as Map
import Data.Foldable (foldl)

import Data.Set as Set
import Data.String.Pattern (Pattern(..), Replacement(..))
import Data.String as String
import Data.Newtype (unwrap)
import PureScript.Backend.Optimizer.Builder (buildModules)
import PureScript.Backend.Optimizer.Monomorphize (collectInstantiations, InstantiationMap, collectAllTypes, monomorphize, transitiveCollect, getExprAnn)
import PureScript.Backend.Optimizer.Semantics.Foreign (coreForeignSemantics)
import PureScript.Backend.Optimizer.CoreFn (Module(..), Ann(..), importName, Bind(..), Binding(..), ExprType(..), Expr(..), Ident(..))
import Gopurs.CodeGen (translate, getStructName)
import Gopurs.Runtime (runtimeGoCode)
import PureScript.Backend.Optimizer.FfiSupport (findFfiFile)
import Gopurs.GoAst (sanitizeName)
import Gopurs.FfiSupport (extractFfiAst)
import Gopurs.FfiTypes (FfiDecl)
import PureScript.Backend.Optimizer.App (coreFnModulesFromOutput, parseCLIArgs, writeCache, loadDirectives)
import Data.Argonaut.Decode (decodeJson)
import Effect.Unsafe (unsafePerformEffect)
import PureScript.Backend.Optimizer.Semantics (InlineDirectiveMap)

type PreparedData =
  { directives :: InlineDirectiveMap
  , elidedCtors :: Set.Set String
  , ctorTypes :: Map.Map String { vars :: Array String, fields :: Array ExprType }
  , globalTypes :: Map.Map String ExprType
  , classDeclsMap :: Map.Map String String
  , classDeclsFields :: Map.Map String { vars :: Array String, fields :: Array { name :: String, "type" :: ExprType } }
  , instantiations :: InstantiationMap
  , monomorphizedModules :: List (Module Ann)
  , adtTypes :: Set.Set ExprType
  , pointerAdtPaths :: Map.Map String { ctorName :: String, arity :: Int }
  , pointerAdtNodes :: Set.Set String
  , pointerAdtLeaves :: Map.Map String { nodeBaseStruct :: String, nodeCtor :: String }
  , enumAdts :: Set.Set String
  , enumCtors :: Set.Set String
  , targetMainModules :: Array String
  }

loadAndPrepareModules :: { mbMainModule :: Maybe String } -> Aff PreparedData
loadAndPrepareModules args = do
  finalModules <- coreFnModulesFromOutput "output"

  let elidedCtors = Array.foldl (\acc (Module mod) ->
        Array.foldl (\acc' decl ->
          if Array.length decl.constructors == 1 then
            case Array.head decl.constructors of
              Just ctor ->
                if Array.length ctor.fields == 1 then
                  let structName = getStructName (unwrap mod.name) Nothing ctor.name
                      constrName = "Constructor_" <> String.drop 5 structName
                  in Set.insert constrName acc'
                else acc'
              Nothing -> acc'
          else acc'
        ) acc mod.dataDecls
      ) Set.empty (Array.fromFoldable finalModules)

  directives <- loadDirectives

  let
    ctorTypes = foldl (\acc (Module m) ->
      let modStr = String.replaceAll (Pattern ".") (Replacement "_") (unwrap m.name)
      in foldl (\acc' decl ->
        foldl (\acc'' ctor ->
          Map.insert (modStr <> "." <> ctor.name) { vars: decl.vars, fields: ctor.fields } acc''
        ) acc' decl.constructors
      ) acc m.dataDecls
    ) Map.empty finalModules

  let globalTypes = buildGlobalTypes (Array.fromFoldable finalModules)
  let

    classDeclsMap = foldl (\acc (Module m) ->
      foldl (\acc' c ->
        let superclassFields = Array.mapWithIndex (\i super ->
                  fromMaybe "" (Array.last (fst super)) <> show i
                ) c.superclasses
            methodFields = map fst c.methods
            allFields = superclassFields <> methodFields
            key = String.joinWith "," allFields
        in Map.insert key (unwrap m.name <> "." <> c.name) acc'
      ) acc m.classDecls
    ) Map.empty finalModules

    classDeclsFields = foldl (\acc (Module m) ->
      foldl (\acc' c ->
        let superclassFields = Array.mapWithIndex (\i super ->
                  let superName = fromMaybe "" (Array.last (fst super))
                  in Tuple (superName <> show i) Any
                ) c.superclasses
            methodFields = c.methods
            allFields = Array.sortBy (comparing fst) (superclassFields <> methodFields)
            fieldsWithTypes = map (\(Tuple name ty) -> { name, "type": ty }) allFields
            vars = c.vars
        in Map.insert (unwrap m.name <> "." <> c.name) { vars, fields: fieldsWithTypes } acc'
      ) acc m.classDecls
    ) Map.empty finalModules

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
  
  let rawInstantiations = foldl collectInstantiations Map.empty finalModulesWithClassDecls
      
  let globalAstMap = foldl (\acc (Module m) ->
        foldl (\acc' b -> case b of
          NonRec (Binding ann id e) -> Map.insert (unwrap m.name <> "." <> unwrap id) (Binding ann id e) acc'
          Rec binds -> foldl (\a (Binding ann id e) -> Map.insert (unwrap m.name <> "." <> unwrap id) (Binding ann id e) a) acc' binds
        ) acc m.decls
      ) Map.empty finalModulesWithClassDecls

  let transitiveInstantiations = transitiveCollect globalAstMap rawInstantiations
  let instantiations = Map.filterKeys (\k -> case Map.lookup k globalTypes of
                                            Just t -> hasTypeVariables t
                                            Nothing -> false) transitiveInstantiations

  let monomorphizedModules = map (monomorphize globalAstMap instantiations) finalModulesWithClassDecls

  let allTypes = foldl (\acc mod -> Set.union acc (collectAllTypes mod)) Set.empty finalModulesWithClassDecls
  let adtTypes = Set.filter (\t -> case t of
        ADT _ _ _ -> true
        _ -> false) allTypes
  let
    -- pointerAdts computation
    pointerAdtPathsRaw = foldl (\acc (Module m) ->
      foldl (\acc' d ->
        let
          ctorsWithFields = Array.filter (\c -> Array.length c.fields > 0) d.constructors
          ctorsWithoutFields = Array.filter (\c -> Array.length c.fields == 0) d.constructors
        in
          if Array.length ctorsWithFields == 1 then
            let
              adtPath = unwrap m.name <> "." <> d.name
              nodeCtor = (unsafePartial (Array.unsafeIndex ctorsWithFields 0)).name
              leafCtor = if Array.length ctorsWithoutFields == 1 then (unsafePartial (Array.unsafeIndex ctorsWithoutFields 0)).name else ""
              pkgNameStr = String.replaceAll (Pattern ".") (Replacement "_") (unwrap m.name)
              nodeBaseStruct = "Data_" <> pkgNameStr <> "_" <> sanitizeName nodeCtor
              leafBaseStruct = if leafCtor /= "" then "Data_" <> pkgNameStr <> "_" <> sanitizeName leafCtor else ""
              arity = Array.length d.vars
            in
              Array.snoc acc' { adtPath, nodeCtor, leafCtor, nodeBaseStruct, leafBaseStruct, arity }
          else acc'
      ) acc m.dataDecls
     ) [] finalModulesWithClassDecls

    pointerAdtPaths = Map.fromFoldable (map (\info -> Tuple info.adtPath { ctorName: info.nodeCtor, arity: info.arity }) pointerAdtPathsRaw)
    pointerAdtNodes = Set.fromFoldable (map _.nodeBaseStruct pointerAdtPathsRaw)
    pointerAdtLeaves = Map.fromFoldable (Array.mapMaybe (\info -> if info.leafBaseStruct /= "" then Just (Tuple info.leafBaseStruct { nodeBaseStruct: info.nodeBaseStruct, nodeCtor: info.nodeCtor }) else Nothing) pointerAdtPathsRaw)

    enumAdtsRaw = foldl (\acc (Module m) ->
      foldl (\acc' d ->
        let
          ctorsWithFields = Array.filter (\c -> Array.length c.fields > 0) d.constructors
        in
          if Array.length ctorsWithFields == 0 && Array.length d.constructors > 0 then
            let
              adtPath = unwrap m.name <> "." <> d.name
              pkgNameStr = String.replaceAll (Pattern ".") (Replacement "_") (unwrap m.name)
              ctorBaseStructs = map (\c -> "Data_" <> pkgNameStr <> "_" <> sanitizeName c.name) d.constructors
            in Array.snoc acc' { adtPath, ctors: ctorBaseStructs }
          else acc'
      ) acc m.dataDecls
     ) [] finalModulesWithClassDecls

    enumAdts = Set.fromFoldable (map _.adtPath enumAdtsRaw)
    enumCtors = Set.fromFoldable (Array.concatMap _.ctors enumAdtsRaw)

    targetMainModules = case args.mbMainModule of
      Just mainMod -> [ mainMod ]
      Nothing -> Array.mapMaybe (\(Module m) -> if isJust (Array.elemIndex (Ident "main") m.exports) then Just (unwrap m.name) else Nothing) (Array.fromFoldable finalModules)

  pure { directives
       , elidedCtors
       , ctorTypes
       , globalTypes
       , classDeclsMap
       , classDeclsFields
       , instantiations
       , monomorphizedModules
       , adtTypes
       , pointerAdtPaths
       , pointerAdtNodes
       , pointerAdtLeaves
       , enumAdts
       , enumCtors
       , targetMainModules
       }

buildGlobalTypes :: Array (Module Ann) -> Map.Map String ExprType
buildGlobalTypes modules = Array.foldl (\acc (Module m) ->
  let modName = unwrap m.name
      processBind acc' (NonRec (Binding (Ann ann) (Ident name) expr)) =
        let ty = case ann.type of
                   Just t -> Just t
                   Nothing -> case let (Ann exprAnn) = getExprAnn expr in exprAnn.type of
                                Just t -> Just t
                                Nothing -> inferExprType expr
        in case ty of
            Just t -> Map.insert (modName <> "." <> name) t acc'
            Nothing -> acc'
      processBind acc' (Rec bindings) = Array.foldl (\a b -> processBind a (NonRec b)) acc' bindings
      withDecls = Array.foldl processBind acc m.decls
  in foldl (\acc' (Tuple (Ident name) mbTy) -> 
        case mbTy of
          Just ty -> Map.insert (modName <> "." <> name) ty acc'
          Nothing -> acc'
      ) withDecls (Map.toUnfoldable m.foreign :: Array (Tuple Ident (Maybe ExprType)))
  ) Map.empty modules

inferExprType :: Expr Ann -> Maybe ExprType
inferExprType (ExprApp _ fn _) = case getExprAnn fn of
  Ann { type: Just ty } -> getReturnType ty
  _ -> case inferExprType fn of
         Just ty -> getReturnType ty
         Nothing -> Nothing
inferExprType (ExprTypeApp _ fn _) = inferExprType fn
inferExprType _ = Nothing

getReturnType :: ExprType -> Maybe ExprType
getReturnType (ForAll _ t) = getReturnType t
getReturnType (ConstrainedType _ t) = getReturnType t
getReturnType (Func _ ret) = Just ret
getReturnType _ = Nothing

hasTypeVariables :: ExprType -> Boolean
hasTypeVariables (TypeVar v) = String.take 1 v == String.toLower (String.take 1 v) && v /= "gopurs_runtime.Value"

hasTypeVariables (Func args ret) = Array.any hasTypeVariables args || hasTypeVariables ret
hasTypeVariables (Array t) = hasTypeVariables t
hasTypeVariables (Record row) = hasTypeVariables row
hasTypeVariables (Row props tail) = 
  let tailHas = case tail of
        Nothing -> false
        Just t -> hasTypeVariables t
  in Array.any (\(Tuple _ v) -> hasTypeVariables v) props || tailHas
hasTypeVariables (TypeApp c args) = hasTypeVariables c || Array.any hasTypeVariables args
hasTypeVariables (ForAll _ body) = hasTypeVariables body
hasTypeVariables (ConstrainedType constraints body) = Array.any (\(Tuple _ a) -> Array.any hasTypeVariables a) constraints || hasTypeVariables body
hasTypeVariables Int = false
hasTypeVariables String = false
hasTypeVariables Char = false
hasTypeVariables Number = false
hasTypeVariables Boolean = false
hasTypeVariables Unit = false
hasTypeVariables (TypeLevelString _) = false
hasTypeVariables (ADT _ _ args) = Array.any hasTypeVariables args
hasTypeVariables Any = false

cacheVersion :: String
cacheVersion = "1.0.0"

main :: Effect Unit
main = launchAff_ do
  argsRaw <- liftEffect Process.argv
  let args = parseCLIArgs argsRaw

  prepared <- loadAndPrepareModules { mbMainModule: args.mbMainModule }

  _ <- attempt (FS.mkdir "output/gopurs_runtime")
  FS.writeTextFile UTF8 "output/gopurs_runtime/runtime.go" runtimeGoCode

  _ <- attempt (FS.mkdir "output/purescript")

  FS.writeTextFile UTF8 "output/go.mod" "module gopurs/output\n\ngo 1.22\n"

  let
    directives = prepared.directives
    enumAdts = prepared.enumAdts
    enumCtors = prepared.enumCtors
    pointerAdtPaths = prepared.pointerAdtPaths
    pointerAdtNodes = prepared.pointerAdtNodes
    pointerAdtLeaves = prepared.pointerAdtLeaves
    adtTypes = prepared.adtTypes
    elidedCtors = prepared.elidedCtors
    ctorTypes = prepared.ctorTypes
    globalTypes = prepared.globalTypes
    instantiations = prepared.instantiations
    classDeclsMap = prepared.classDeclsMap
    classDeclsFields = prepared.classDeclsFields
    monomorphizedModules = prepared.monomorphizedModules
    targetMainModules = prepared.targetMainModules

  buildModules
    { directives: directives
    , analyzeCustom: \_ _ -> Nothing
    , foreignSemantics: coreForeignSemantics
    , traceIdents: Set.empty
    , rewriteLimit: fromMaybe 10_000 args.mbRewriteLimit
    , onPrepareModule: \_ (Module m) -> pure (Module m)
    , onSkipModule: \_ (Module coreFnMod) -> do
        let modNameStr = unwrap coreFnMod.name
        let safeModName = String.replaceAll (Pattern ".") (Replacement "_") modNameStr
        let cachePath = "output/purescript/" <> safeModName <> ".gopurs-cache.json"
        res <- pure Nothing
        case res of
          Just _ -> do
            let ffiPath = "output/purescript/" <> safeModName <> "_ffi.go"
            ffiStatRes <- attempt (FS.stat ffiPath)
            cacheStatRes <- attempt (FS.stat cachePath)
            case ffiStatRes, cacheStatRes of
              Right ffiStat, Right cacheStat ->
                if Stats.modifiedTimeMs ffiStat > Stats.modifiedTimeMs cacheStat then do
                  pure Nothing
                else pure res
              _, _ -> pure res
          Nothing -> pure Nothing
    , onCodegenModule: \_ (Module coreFnMod) backendMod _ -> do
        let modNameStr = unwrap backendMod.name
        let safeModName = String.replaceAll (Pattern ".") (Replacement "_") modNameStr
        let importsArray = map (\i -> String.split (Pattern ".") (unwrap (importName i))) coreFnMod.imports

        let goFile = translate enumAdts enumCtors pointerAdtPaths pointerAdtNodes pointerAdtLeaves adtTypes elidedCtors ctorTypes globalTypes instantiations classDeclsMap classDeclsFields importsArray backendMod
        FS.writeTextFile UTF8 ("output/purescript/" <> safeModName <> ".go") goFile

        when (Array.length (Array.fromFoldable backendMod.foreign) > 0) do
          ffiPathMb <- liftEffect $ findFfiFile ".go" [] args.mbFfiDir modNameStr (Just coreFnMod.path)
          case ffiPathMb of
            Just ffiPath -> do
              content <- FS.readTextFile UTF8 ffiPath
              jsonStr <- liftEffect $ extractFfiAst modNameStr content
              let parsed = (jsonParser jsonStr >>= (decodeJson >>> lmap printJsonDecodeError)) :: Either String (Array FfiDecl)
              let ffiDecls = case parsed of
                               Right d -> d
                               Left err -> unsafePerformEffect (Console.log ("JSON Parse error for " <> modNameStr <> ": " <> err) *> pure [])

              let lines = String.split (Pattern "\n") (String.replaceAll (Pattern "\r") (Replacement "") content)
              let otherLines = Array.filter (\l -> not (String.contains (Pattern "package ") l)) lines
              let finalPkgLine = "package purescript"
              let hasImport = String.contains (Pattern "\"gopurs/output/gopurs_runtime\"") content
              let importLine = if hasImport then "" else "import \"gopurs/output/gopurs_runtime\"\n"
              
              let prefixedFfiDecls = map (\d -> d { name = safeModName <> "_" <> d.name }) ffiDecls
              let renamedContentLines = map (\l -> Array.foldl (\acc decl ->
                                                if decl.isVar 
                                                then String.replaceAll (Pattern ("var " <> decl.name)) (Replacement ("var " <> safeModName <> "_" <> decl.name)) acc
                                                else String.replaceAll (Pattern ("func " <> decl.name)) (Replacement ("func " <> safeModName <> "_" <> decl.name)) acc
                                             ) l ffiDecls) otherLines
              
              let newContent = finalPkgLine <> "\n\n" <> importLine <> "\n" <> String.joinWith "\n" renamedContentLines <> "\n\n// --- Auto-generated FFI wrappers ---\n" <> CodeGen.generateFfiBridge safeModName backendMod.dataDecls prefixedFfiDecls (Map.toUnfoldable backendMod.foreign)
              FS.writeTextFile UTF8 ("output/purescript/" <> safeModName <> "_ffi.go") newContent
            Nothing -> do

              let dummyContent = "package purescript\n\nimport \"gopurs/output/gopurs_runtime\"\n\n" <> CodeGen.generateFfiBridge safeModName backendMod.dataDecls [] (Map.toUnfoldable backendMod.foreign)
              FS.writeTextFile UTF8 ("output/purescript/" <> safeModName <> "_ffi.go") dummyContent
        writeCache cacheVersion ("output/purescript/" <> safeModName <> ".gopurs-cache.json") backendMod
    }
    (List.fromFoldable monomorphizedModules)

  _ <- traverse
    ( \mainMod -> do
        let pkgName = String.replaceAll (Pattern ".") (Replacement "_") mainMod
        let mainEntryPoint = "package main\n\nimport (\n\t\"os\"\n\t\"runtime/pprof\"\n\t\"gopurs/output/purescript\"\n\t\"gopurs/output/gopurs_runtime\"\n)\n\nfunc main() {\n\tif os.Getenv(\"PPROF\") == \"1\" {\n\t\tf, err := os.Create(\"cpu.prof\")\n\t\tif err != nil { panic(err) }\n\t\tpprof.StartCPUProfile(f)\n\t\tdefer pprof.StopCPUProfile()\n\t}\n\n\tgopurs_runtime.Apply(purescript.Get_" <> pkgName <> "_main(), gopurs_runtime.Value{})\n\n\tgopurs_runtime.EventLoopWait()\n\n\tif os.Getenv(\"PPROF\") == \"1\" {\n\t\tmf, err := os.Create(\"mem.prof\")\n\t\tif err != nil { panic(err) }\n\t\tpprof.WriteHeapProfile(mf)\n\t\tmf.Close()\n\t}\n}\n"
        
        _ <- attempt (FS.mkdir "output/main")
        FS.writeTextFile UTF8 ("output/main/main.go") mainEntryPoint

        _ <- attempt (FS.mkdir ("output/" <> mainMod))
        _ <- attempt (FS.mkdir ("output/" <> mainMod <> "/main"))
        FS.writeTextFile UTF8 ("output/" <> mainMod <> "/main/main.go") mainEntryPoint
    )
    targetMainModules

  pure unit

