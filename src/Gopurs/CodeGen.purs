module Gopurs.CodeGen where

import Prelude
import Control.Alternative (guard)
import PureScript.Backend.Optimizer.Syntax (BackendSyntax(Var, Local, Lit, App, Abs, UncurriedApp, UncurriedAbs, UncurriedEffectApp, UncurriedEffectAbs, Accessor, Update, CtorSaturated, CtorDef, LetRec, Let, EffectBind, EffectPure, EffectDefer, Branch, PrimOp, PrimEffect, Fail, Typed), BackendAccessor(..), Pair(..), Level, BackendOperator(..), BackendOperator1(..), BackendOperator2(..), BackendEffect(..))
import PureScript.Backend.Optimizer.Syntax as Syn
import PureScript.Backend.Optimizer.Convert (BackendModule)
import Data.String as String
import Data.Array as Array
import Data.Maybe (Maybe(..), fromMaybe)
import Data.Newtype (unwrap)
import PureScript.Backend.Optimizer.CoreFn (DataDecl, ExprType(..), Ident(..), Literal(..), ModuleName(..), Prop(..), Qualified(..))
import Data.Tuple (Tuple(..), fst, snd)
import Data.Array.NonEmpty as NonEmptyArray
import Data.Array.NonEmpty (NonEmptyArray, fromArray, toArray)
import Effect (Effect)
import Effect.Unsafe (unsafePerformEffect)
import Effect.Ref (Ref)
import Effect.Ref as Ref

import Data.String.Pattern (Pattern(..), Replacement(..))
import Data.Map (Map)
import Debug as Debug
import Data.Map as Map
import Data.Set (Set)
import Data.Set as Set
import Data.Foldable (foldl, foldMap, any)
import Data.List as List
import Data.Traversable (traverse)

import Gopurs.GoAst (GoDecl, GoExpr(..), GoType(..), goTypeToStr, sanitizeName)
import Gopurs.GoAst as GoAst

import Gopurs.Printer (printGoFile, printGoExpr, printGoDeclVar)
import PureScript.Backend.Optimizer.Codegen.Tco as Tco
import PureScript.Backend.Optimizer.Codegen.Tco (TcoExpr(..))
import PureScript.Backend.Optimizer.FreeVars (freeVars, localId)
import PureScript.Backend.Optimizer.FfiSupport (hashString)
import Gopurs.FfiTypes (TypeNode, FfiDecl)
import Gopurs.FfiBridge as FfiBridge
import Gopurs.ThunkFusion (optimizeThunkProducers)
import Gopurs.GoTypes as GoTypes
import Gopurs.CodegenState as State
import Gopurs.GoConversions as GoConversions
import Gopurs.PrimitiveExprs as PrimitiveExprs

type CodegenMetadataRow :: Row Type
type CodegenMetadataRow = State.CodegenMetadataRow

type CodegenMetadata = State.CodegenMetadata

type CodegenState = State.CodegenState

type LocalBinding =
  { name :: String
  , goType :: GoType
  }

-- Keys are original localId values; each binding's name is the emitted Go name,
-- which may have been renamed.
type LocalEnv = Map String LocalBinding

type FunctionInfo =
  { fullName :: String
  , fArgs :: Array GoType
  , fRet :: GoType
  , arity :: Int
  }

type ModuleFunctions = Map String FunctionInfo

type LoopTarget =
  { ident :: String
  , params :: Array String
  , loopParams :: Array String
  , goTypes :: Array GoType
  , fRet :: GoType
  }

type LoopContext = Array LoopTarget

type ExprOptions =
  { isTail :: Boolean
  , inEffectBlock :: Boolean
  }

type ExprResult =
  { stmts :: StmtTree
  , expr :: GoExpr
  , exprType :: GoType
  , nextId :: Int
  }

foreign import memoizedFreeVarsImpl :: (TcoExpr -> Set String) -> TcoExpr -> Set String

memoizedFreeVars :: TcoExpr -> Set String
memoizedFreeVars = memoizedFreeVarsImpl freeVars

type UnboxedADT = GoConversions.UnboxedADT

unboxableADTs :: Map String UnboxedADT
unboxableADTs = GoConversions.unboxableADTs

getUnboxedADT :: ExprType -> Maybe (Tuple String UnboxedADT)
getUnboxedADT = GoConversions.getUnboxedADT

coerceGoExpr :: Ref CodegenState -> String -> GoExpr -> GoType -> GoType -> GoExpr
coerceGoExpr = GoConversions.coerceGoExpr

boxGoExpr :: Ref CodegenState -> String -> GoExpr -> GoType -> GoExpr
boxGoExpr = GoConversions.boxGoExpr

boxGoExprImpl :: Ref CodegenState -> String -> GoExpr -> GoType -> GoExpr
boxGoExprImpl = GoConversions.boxGoExprImpl

isClosedRowTail :: Maybe ExprType -> Boolean
isClosedRowTail = GoTypes.isClosedRowTail

exprTypeToGoType :: Map.Map String { ctorName :: String, arity :: Int } -> Set.Set String -> Set.Set String -> String -> ExprType -> GoType
exprTypeToGoType = GoTypes.exprTypeToGoType

exprTypeToGenericGoType :: Map.Map String { ctorName :: String, arity :: Int } -> Set.Set String -> Set.Set String -> Array String -> String -> ExprType -> GoType
exprTypeToGenericGoType = GoTypes.exprTypeToGenericGoType

structFieldGoType :: Map.Map String { ctorName :: String, arity :: Int } -> Set.Set String -> Set.Set String -> Array String -> String -> ExprType -> GoType
structFieldGoType = GoTypes.structFieldGoType

instantiateGenericGoType :: Map.Map String GoType -> GoType -> GoType
instantiateGenericGoType = GoTypes.instantiateGenericGoType

unboxGoExpr :: Ref CodegenState -> String -> GoExpr -> GoType -> GoType -> GoExpr
unboxGoExpr = GoConversions.unboxGoExpr

capitalize :: String -> String
capitalize = GoAst.capitalize

data StmtTree = StmtEmpty | StmtLeaf GoExpr | StmtAppend StmtTree StmtTree

instance Semigroup StmtTree where
  append StmtEmpty a = a
  append a StmtEmpty = a
  append a b = StmtAppend a b

instance Monoid StmtTree where
  mempty = StmtEmpty

flattenStmts :: StmtTree -> Array GoExpr
flattenStmts tree = Array.fromFoldable (go List.Nil tree)
  where
  go acc StmtEmpty = acc
  go acc (StmtLeaf s) = List.Cons s acc
  go acc (StmtAppend a b) =
    let
      acc' = go acc b
    in
      go acc' a

wrapInStmts :: Array String -> StmtTree -> GoType -> GoExpr -> GoExpr
wrapInStmts _ stmts retType expr =
  let
    stmtsArr = flattenStmts stmts
  in
    if Array.length stmtsArr == 0 then expr
    else GoCall (GoFuncLit [] stmtsArr expr retType) []

-- Reuse an unchanged constructor without mutating it. Restrict this to direct
-- projections of one typed local, with exactly one constant field replaced.
-- Comparing Number fields would be unsound for observable signed zero.
constructorReuse :: LocalEnv -> GoType -> Array Boolean -> GoExpr -> Maybe { source :: GoExpr, condition :: GoExpr }
constructorReuse bound resultType constants constructor = case resultType, constructor of
  TypeStructPointer _ _ _ _, GoConstructor _ ctor typeArgs fields ->
    case Array.catMaybes (Array.mapWithIndex (\index constant -> if constant then Just index else Nothing) constants) of
      [ changedIndex ] -> do
        guard (Array.length constants == Array.length fields)
        sourceName <- Array.head (Array.mapMaybe
          (\(Tuple index field) -> case field of
            GoConstructorAccess (GoVar name) sourceCtor sourceTypeArgs sourceIndex true
              | index /= changedIndex && sourceCtor == ctor
                  && sourceTypeArgs == typeArgs && sourceIndex == index -> Just name
            _ -> Nothing)
          (Array.mapWithIndex Tuple fields))
        sourceBinding <- Array.find (\binding -> binding.name == sourceName)
          (Array.fromFoldable (Map.values bound))
        guard (sourceBinding.goType == resultType)
        let
          source = GoVar sourceName
          projection index = GoConstructorAccess source ctor typeArgs index true
        guard (Array.all identity (Array.mapWithIndex
          (\index field -> index == changedIndex || field == projection index) fields))
        replacement <- Array.index fields changedIndex
        pure
          { source
          , condition: GoBinOp "&&" (GoBinOp "!=" source (GoRaw "nil"))
              (GoBinOp "==" (projection changedIndex) replacement)
          }
      _ -> Nothing
  _, _ -> Nothing

type CurriedAbs =
  { args :: NonEmptyArray (Tuple (Maybe Ident) Level)
  , body :: TcoExpr
  }

-- Collect adjacent curried lambdas without crossing a computation.
-- Typed is transparent only when looking for another Abs; the terminal body
-- keeps its original annotations for translation.
collectCurriedAbs :: NonEmptyArray (Tuple (Maybe Ident) Level) -> TcoExpr -> CurriedAbs
collectCurriedAbs args body =
  case lookAheadAbs body of
    Just next -> collectCurriedAbs (args <> next.args) next.body
    Nothing -> { args, body }
  where
  lookAheadAbs :: TcoExpr -> Maybe CurriedAbs
  lookAheadAbs (TcoExpr _ syntax) = case syntax of
    Typed _ inner -> lookAheadAbs inner
    Abs nextArgs nextBody -> Just { args: nextArgs, body: nextBody }
    _ -> Nothing

extractUncurriedAbs :: TcoExpr -> Maybe { args :: Array String, body :: TcoExpr }
extractUncurriedAbs tcoExpr@(TcoExpr _ syntax) = case syntax of
  UncurriedAbs args body ->
    let
      thisArgs = map (\(Tuple mbI lvl) -> localId mbI lvl) args
    in
      case extractUncurriedAbs body of
        Just inner -> Just { args: thisArgs <> inner.args, body: inner.body }
        Nothing -> Just { args: thisArgs, body }
  Abs args body ->
    let
      thisArgs = map (\(Tuple mbI lvl) -> localId mbI lvl) (toArray args)
    in
      case extractUncurriedAbs body of
        Just inner -> Just { args: thisArgs <> inner.args, body: inner.body }
        Nothing -> Just { args: thisArgs, body }
  Typed _ inner -> extractUncurriedAbs inner
  _ -> Nothing

data GoSpineArg = GoSpineApp (Array TcoExpr) | GoSpineTypeApp ExprType

getGoSpineArgs :: Array GoSpineArg -> Array TcoExpr
getGoSpineArgs = Array.concatMap extractApp
  where
  extractApp (GoSpineApp a) = a
  extractApp _ = []

unwrapForSpine :: TcoExpr -> BackendSyntax TcoExpr
unwrapForSpine (TcoExpr _ syn) = case syn of
  Typed _ inner -> unwrapForSpine inner
  _ -> syn

collectGoSpine :: TcoExpr -> Tuple TcoExpr (Array GoSpineArg)
collectGoSpine e =
  case unwrapForSpine e of
    App f args ->
      let
        Tuple f' args' = collectGoSpine f
      in
        Tuple f' (args' <> [GoSpineApp (toArray args)])
    UncurriedApp f args ->
      let
        Tuple f' args' = collectGoSpine f
      in
        Tuple f' (args' <> [GoSpineApp args])
    Syn.TypeApp f ty ->
      let
        Tuple f' args' = collectGoSpine f
      in
        Tuple f' (args' <> [GoSpineTypeApp ty])
    _ -> Tuple e []

getBaseStructName :: String -> Maybe ModuleName -> String -> String
getBaseStructName modNameStr mbMod ctorName =
  let
    modNamePart = case mbMod of
      Just mn -> sanitizeName (String.replaceAll (Pattern ".") (Replacement "_") (unwrap mn))
      Nothing -> modNameStr
  in
    "Data_" <> modNamePart <> "_" <> sanitizeName ctorName

getStructName :: String -> Maybe ModuleName -> String -> String
getStructName modNameStr mbMod ctorName =
  getBaseStructName modNameStr mbMod ctorName

type ReboxFields = GoConversions.ReboxFields

registerReboxPair :: Ref CodegenState -> GoType -> GoType -> Effect Unit
registerReboxPair = GoConversions.registerReboxPair

findReboxFields :: CodegenState -> String -> Maybe ReboxFields
findReboxFields = GoConversions.findReboxFields

renderReboxFunction :: Ref CodegenState -> CodegenState -> String -> Map String String -> Tuple GoType GoType -> Maybe (Tuple String String)
renderReboxFunction = GoConversions.renderReboxFunction

generateReboxFunctions :: Ref CodegenState -> String -> Effect (Array String)
generateReboxFunctions = GoConversions.generateReboxFunctions

translate :: CodegenMetadata -> BackendModule -> String
translate { enumAdts, enumCtors, pointerAdtPaths, pointerAdtNodes, pointerAdtLeaves, elidedCtors, ctorTypes, globalTypes, classDeclsFields } inputMod =

  let
    mod = optimizeThunkProducers inputMod
    modNameStrOrig = unwrap mod.name
    modNameStr = String.replaceAll (Pattern ".") (Replacement "_") modNameStrOrig

    codegenStateRef :: Ref CodegenState
    codegenStateRef = unsafePerformEffect do
      let
        structDecls = Array.concatMap
          ( \decl ->
              Array.concatMap
                ( \ctor ->
                    let
                      fieldTypes = ctor.fields
                      goFieldTypes = map (structFieldGoType pointerAdtPaths enumAdts elidedCtors decl.vars modNameStr) fieldTypes
                      structName = "Constructor_" <> modNameStr <> "_" <> sanitizeName ctor.name

                      typeParams =
                        if Array.length decl.vars > 0 then
                          "[" <> String.joinWith ", " (map (\v -> "T_" <> sanitizeName v <> " any") decl.vars) <> "]"
                        else ""

                      fieldsStr = Array.cons "Rc uint32" (Array.mapWithIndex (\i ty -> "V" <> show i <> " " <> goTypeToStr ty) goFieldTypes)
                      structDecl = "type " <> structName <> typeParams <> " struct {\n\t" <> String.joinWith "\n\t" fieldsStr <> "\n}\n"

                      fullName = unwrap mod.name <> "." <> ctor.name
                      getterDecl = case Map.lookup fullName classDeclsFields of
                        Just info ->
                          let
                            typeParamsGetter =
                              if Array.length decl.vars > 0 then
                                "[" <> String.joinWith ", " (map (const "any") decl.vars) <> "]"
                              else ""
                            cases = Array.mapWithIndex (\i f -> "\t\tcase \"" <> f.name <> "\": return gopurs_runtime.Box(c.V" <> show i <> ")") info.fields
                            pkgNameStr = String.replaceAll (Pattern ".") (Replacement "_") (unwrap mod.name)
                            baseStructName = "Data_" <> pkgNameStr <> "_" <> sanitizeName ctor.name
                            hashStr = hashString baseStructName
                          in
                            "func init() {\n\tgopurs_runtime.StructGetters[" <> hashStr <> "] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {\n\t\tc := (*" <> structName <> typeParamsGetter <> ")(ptr)\n\t\t_ = c\n\t\tswitch key {\n" <> String.joinWith "\n" cases <> "\n\t\tdefault: panic(\"Key not found in dictionary " <> structName <> ": \" + key)\n\t\t}\n\t}\n}\n"
                        Nothing -> ""

                    in
                      if getterDecl == "" then [ structDecl ] else [ structDecl, getterDecl ]
                )
                decl.constructors
          )
          mod.dataDecls

      -- Rebox requests belong to this invocation of translate.
      Ref.new { decls: [], rawDecls: structDecls, elidedCtors, ctorTypes, pointerAdtPaths, pointerAdtNodes, pointerAdtLeaves, enumAdts, enumCtors, globalTypes, classDeclsFields, globalId: 0, reboxPairs: Set.empty }

    Tuple _ tcoBindings = foldl
      ( \(Tuple env acc) group ->
          let
            neBindings = fromArray group.bindings

            env' = case neBindings of
              Just ne | group.recursive -> Tco.topLevelTcoEnvGroup mod.name ne <> env
              _ -> env
            tcoBinds = map
              ( \(Tuple id val) -> Tuple id (Tco.analyze env' val) )
              group.bindings
          in
            Tuple env' (Array.snoc acc { recursive: group.recursive, bindings: tcoBinds })
      )
      (Tuple [] [])
      mod.bindings

    getExprType :: TcoExpr -> ExprType
    getExprType (TcoExpr _ (Typed ty _)) = ty
    getExprType _ = Any -- fallback

    tcoBindingsExpanded = tcoBindings

    unwrapFunc :: Array (Tuple Ident TcoExpr) -> Array (Tuple String FunctionInfo)
    unwrapFunc binds =
      Array.concatMap
        ( \(Tuple (Ident name) val) ->
            case extractUncurriedAbs val of
              Just { args, body } ->
                let
                  typeSig = extractFuncType val
                  fArgsGo = case typeSig of
                    Just { fArgs } -> map (exprTypeToGoType pointerAdtPaths enumAdts elidedCtors modNameStr) (Array.take (Array.length args) fArgs)
                    Nothing -> Array.replicate (Array.length args) TypeValue
                  fRetGo = case typeSig of
                    Just { fArgs, fRet } ->
                      let
                        ret = if Array.length args < Array.length fArgs then TypeValue else exprTypeToGoType pointerAdtPaths enumAdts elidedCtors modNameStr fRet
                      in
                        case getUnboxedADT fRet of
                          Just (Tuple adtName adt) -> TypeStructValue adtName adt.signature
                          Nothing -> ret
                    Nothing -> TypeValue
                  fullName = "Call_" <> modNameStr <> "_" <> sanitizeName name
                in
                  [ Tuple (sanitizeName name) { fullName, fArgs: fArgsGo, fRet: fRetGo, arity: Array.length args } ]
              Nothing ->
                let
                  fullName = "Call_" <> modNameStr <> "_" <> sanitizeName name
                in
                  [ Tuple (sanitizeName name) { fullName, fArgs: [], fRet: TypeValue, arity: 0 } ]
        )
        binds

    moduleFunctions :: ModuleFunctions
    moduleFunctions = Map.fromFoldable $ Array.concatMap
      ( \group ->
          if group.recursive then
            let
              mutRecBinds = traverse (\(Tuple _ val) -> extractUncurriedAbs val) group.bindings
            in
              case mutRecBinds of
                Just _ -> unwrapFunc group.bindings
                Nothing -> map (\(Tuple (Ident name) _) -> Tuple (sanitizeName name) { fullName: "Call_" <> modNameStr <> "_" <> sanitizeName name, fArgs: [], fRet: TypeValue, arity: 0 }) group.bindings
          else
            unwrapFunc group.bindings
      )
      tcoBindingsExpanded

    Tuple decls helpers = unsafePerformEffect do
      let
        d = Array.concatMap
          ( \group ->
              let
                recVars = if group.recursive then map (\(Tuple (Ident name) _) -> sanitizeName name) group.bindings else []

                processBindingGroup :: Array (Tuple Ident TcoExpr) -> Boolean -> Array GoDecl
                processBindingGroup binds _ =
                  let
                    mutRecBinds = traverse (\(Tuple (Ident name) val) -> map (\abs -> { ident: sanitizeName name, args: abs.args, body: abs.body, fvs: memoizedFreeVars val, val: val }) (extractUncurriedAbs val)) binds
                  in
                    case mutRecBinds of
                      Just fns ->
                        let
                          fnWrapperStmts = map
                            ( \fn ->
                                let
                                  paramsWithTypes = case extractExprFuncType (getExprType fn.val) of
                                    Just { fArgs } -> Array.zipWith (\p goType -> Tuple p goType) fn.args (map (exprTypeToGoType (unsafePerformEffect (Ref.read codegenStateRef)).pointerAdtPaths (unsafePerformEffect (Ref.read codegenStateRef)).enumAdts (unsafePerformEffect (Ref.read codegenStateRef)).elidedCtors modNameStr) fArgs <> Array.replicate (Array.length fn.args - Array.length fArgs) TypeValue)
                                    Nothing -> map (\p -> Tuple p TypeValue) fn.args

                                  newBound = foldl (\acc (Tuple idStr goType) -> Map.insert idStr { name: idStr, goType } acc) Map.empty paramsWithTypes

                                  isSelfRecursiveLoop = group.recursive && Array.length group.bindings == 1
                                  fRet = case extractExprFuncType (getExprType fn.val) of
                                    Just { fRet: rt } -> exprTypeToGoType (unsafePerformEffect (Ref.read codegenStateRef)).pointerAdtPaths (unsafePerformEffect (Ref.read codegenStateRef)).enumAdts (unsafePerformEffect (Ref.read codegenStateRef)).elidedCtors modNameStr rt
                                    Nothing -> TypeValue
                                  mbExpectedRet = case extractExprFuncType (getExprType fn.val) of
                                    Just { fRet: rt } -> Just rt
                                    Nothing -> Nothing
                                  currentLoopCtx :: LoopContext
                                  currentLoopCtx = if isSelfRecursiveLoop then [ { ident: fn.ident, params: map fst paramsWithTypes, loopParams: map (\p -> fst p <> "_loop") paramsWithTypes, goTypes: map snd paramsWithTypes, fRet } ] else []
                                  resBodyMut = translateExprWithExpectedType codegenStateRef 0 modNameStr recVars moduleFunctions newBound (Just fn.ident) currentLoopCtx { isTail: isSelfRecursiveLoop, inEffectBlock: false } mbExpectedRet 0 fn.body

                                  goName = fn.ident
                                  initVars = Array.concatMap (\(Tuple p goT) -> [ GoRaw ("var " <> p <> " " <> goTypeToStr goT <> " = " <> p <> "_loop"), GoRaw ("_ = " <> p) ]) paramsWithTypes

                                  arity = Array.length fn.args

                                  expectedRetType = case Map.lookup goName moduleFunctions of
                                    Just { fRet, fArgs } | arity < Array.length fArgs -> TypeValue
                                    Just { fRet } -> fRet
                                    Nothing -> TypeValue

                                  goParams = String.joinWith ", " (map (\(Tuple p goT) -> p <> "_loop " <> goTypeToStr goT) paramsWithTypes)

                                  funcExpr =
                                    if arity >= 1 && arity <= 10 then
                                      let
                                        coercedExpr = coerceGoExpr codegenStateRef modNameStr resBodyMut.expr resBodyMut.exprType expectedRetType
                                        bodyStmts = initVars <> flattenStmts resBodyMut.stmts <> [ GoReturn coercedExpr ]
                                        funcBody = if isSelfRecursiveLoop then GoFor goName bodyStmts else GoBlock bodyStmts
                                      in
                                        unsafePerformEffect do
                                          let callFuncDecl = "func Call_" <> modNameStr <> "_" <> goName <> "(" <> goParams <> ") " <> goTypeToStr expectedRetType <> " {\n" <> printGoExpr funcBody <> "\n}"
                                          Ref.modify_ (\r -> r { rawDecls = Array.snoc r.rawDecls callFuncDecl }) codegenStateRef
                                          let wrapperParams = map (\(Tuple p _) -> p <> "_box") paramsWithTypes
                                          let callExpr = GoCall (GoVar ("Call_" <> modNameStr <> "_" <> goName)) (map (\(Tuple p goT) -> unboxGoExpr codegenStateRef modNameStr (GoVar (p <> "_box")) TypeValue goT) paramsWithTypes)
                                          let boxedRes = boxGoExpr codegenStateRef modNameStr callExpr expectedRetType
                                          let wrapperFunc = GoRaw ("func(" <> String.joinWith ", " (map (\p -> p <> " gopurs_runtime.Value") wrapperParams) <> ") gopurs_runtime.Value {\nreturn " <> printGoExpr boxedRes <> "\n}")
                                          let funcWrapperName = if arity == 1 then "gopurs_runtime.Func" else "gopurs_runtime.Func" <> show arity
                                          pure $ GoRaw (funcWrapperName <> "(" <> printGoExpr wrapperFunc <> ")")
                                    else
                                      let
                                        bodyStmts = initVars <> flattenStmts resBodyMut.stmts <> [ GoReturn (boxGoExpr codegenStateRef modNameStr resBodyMut.expr resBodyMut.exprType) ]
                                        funcBody = if isSelfRecursiveLoop then GoFor goName bodyStmts else GoBlock bodyStmts
                                        iife = GoRaw ("func() gopurs_runtime.Value {\n" <> printGoExpr funcBody <> "\n}()")
                                      in
                                        Array.foldr (\(Tuple p goT) acc -> GoCall (GoSelector (GoVar "gopurs_runtime") "Func") [ GoRaw ("func(" <> p <> "_box gopurs_runtime.Value) gopurs_runtime.Value {\nvar " <> p <> "_loop " <> goTypeToStr goT <> " = " <> printGoExpr (unboxGoExpr codegenStateRef modNameStr (GoVar (p <> "_box")) TypeValue goT) <> "\nreturn " <> printGoExpr acc <> "\n}") ]) iife paramsWithTypes
                                in
                                  { identifier: modNameStr <> "_" <> goName, expression: funcExpr, goType: TypeValue }
                            )
                            fns
                        in
                          fnWrapperStmts
                      Nothing ->
                        Array.concatMap
                          ( \(Tuple (Ident name) expr) ->
                              let
                                res = translateExprWithExpectedType codegenStateRef 0 modNameStr recVars moduleFunctions Map.empty (Just (sanitizeName name)) [] { isTail: false, inEffectBlock: false } (Just (getExprType expr)) 0 expr
                              in
                                [ { identifier: modNameStr <> "_" <> sanitizeName name, expression: wrapInStmts [] res.stmts TypeValue (boxGoExpr codegenStateRef modNameStr res.expr res.exprType), goType: TypeValue } ]
                          )
                          binds
              in
                if group.recursive then
                  processBindingGroup group.bindings true
                else
                  Array.concatMap (\b -> processBindingGroup [ b ] false) group.bindings
          )
          tcoBindingsExpanded
      h <- Ref.read codegenStateRef
      pure (Tuple d h)

    allDeclsAst = decls <> helpers.decls
    declsStr = String.joinWith "\\n" (map printGoDeclVar allDeclsAst) <> "\\n" <> String.joinWith "\\n" helpers.rawDecls

    parts = [ declsStr ]
    usedPkgNames =
      Set.toUnfoldable $ Set.fromFoldable $ Array.mapMaybe
        ( \part ->
            let
              subParts = String.split (Pattern ".") part
            in
              Array.head subParts
        )
        (fromMaybe [] (Array.tail parts)) :: Array String

    goImports = Set.toUnfoldable $ Set.fromFoldable $
      (if Array.length allDeclsAst > 0 || Array.length (Array.fromFoldable mod.foreign) > 0 then [ "gopurs/output/gopurs_runtime" ] else [])
        <> (if Array.length allDeclsAst > 0 then [ "sync" ] else [])
        <> (if String.contains (Pattern "math.") declsStr then [ "math" ] else [])
        <> Array.mapMaybe
          ( \pkg ->
              if pkg /= modNameStr && pkg /= "Prim" && not (String.indexOf (Pattern "Prim_") pkg == Just 0) then Just ("gopurs/output/" <> String.replaceAll (Pattern "_") (Replacement ".") pkg)
              else Nothing
          )
          usedPkgNames

    goFile =
      { packageName: "purescript"
      , imports: goImports
      , decls: allDeclsAst
      , rawDecls: helpers.rawDecls <> unsafePerformEffect (generateReboxFunctions codegenStateRef modNameStr)
      , foreigns: map (\(Tuple (Ident name) type_) -> { pursName: modNameStr <> "_" <> sanitizeName name, goName: "_Gopurs_" <> modNameStr <> "_" <> capitalize (sanitizeName name), exprType: type_ }) (Map.toUnfoldable mod.foreign)
      }
  in
    printGoFile goFile

isEffectNode :: TcoExpr -> Boolean
isEffectNode expr = case unwrapTcoExpr expr of
  EffectBind _ _ _ _ -> true
  EffectPure _ -> true
  EffectDefer _ -> false
  PrimEffect _ -> true
  UncurriedEffectApp _ _ -> true
  Let _ _ _ body -> isEffectNode body
  LetRec _ _ body -> isEffectNode body
  _ -> false

getArityFromType :: ExprType -> Int
getArityFromType = go 0
  where
  go acc (ForAll _ t) = go acc t
  go acc (ConstrainedType _ t) = go acc t
  go acc (Func args ret) = go (acc + Array.length args) ret
  go acc _ = acc

isClosureNode :: forall r. Ref { globalTypes :: Map.Map String ExprType | r } -> TcoExpr -> Boolean
isClosureNode codegenStateRef expr = case unwrapTcoExpr expr of
  Abs _ _ -> true
  UncurriedAbs _ _ -> true
  App _ _ ->
    let
      Tuple flatFn flatArgsSpine = collectGoSpine expr
      flatArgs = getGoSpineArgs flatArgsSpine
      expectedArity = getArityFromType (getExprType flatFn)
      actualArity = Array.length flatArgs
    in
      case unwrapTcoExpr flatFn of
        Var (Qualified mbMn (Ident i)) ->
          let
            h = unsafePerformEffect (Ref.read codegenStateRef)
            vType = case mbMn of
              Just mn -> Map.lookup (unwrap mn <> "." <> i) h.globalTypes
              Nothing -> Nothing

            expectedArity2 = case vType of
              Just t -> getArityFromType t
              Nothing -> 0
          in
            actualArity < expectedArity || actualArity < expectedArity2 || i == "foldrArray" || i == "foldlArray" || i == "traverse_" || i == "for_" || i == "traverseArrayImpl"
        _ -> actualArity < expectedArity
  UncurriedApp _ _ ->
    let
      Tuple flatFn flatArgsSpine = collectGoSpine expr
      flatArgs = getGoSpineArgs flatArgsSpine
      expectedArity = getArityFromType (getExprType flatFn)
      actualArity = Array.length flatArgs
    in
      case unwrapTcoExpr flatFn of
        Var (Qualified mbMn (Ident i)) ->
          let
            h = unsafePerformEffect (Ref.read codegenStateRef)
            vType = case mbMn of
              Just mn -> Map.lookup (unwrap mn <> "." <> i) h.globalTypes
              Nothing -> Nothing

            expectedArity2 = case vType of
              Just t -> getArityFromType t
              Nothing -> 0
          in
            actualArity < expectedArity || actualArity < expectedArity2 || i == "foldrArray" || i == "foldlArray" || i == "traverse_" || i == "for_" || i == "traverseArrayImpl"
        _ -> actualArity < expectedArity
  Let _ _ _ body -> isClosureNode codegenStateRef body
  LetRec _ _ body -> isClosureNode codegenStateRef body
  Typed _ inner -> isClosureNode codegenStateRef inner
  _ -> false

unwrapTcoExpr :: TcoExpr -> BackendSyntax TcoExpr
unwrapTcoExpr (TcoExpr _ syn) = case syn of
  Typed _ inner -> unwrapTcoExpr inner
  Syn.TypeApp inner _ -> unwrapTcoExpr inner
  _ -> syn

printTcoExprShape :: TcoExpr -> String
printTcoExprShape e = case unwrapTcoExpr e of
  Let ident lvl val body -> "Let(" <> printTcoExprShape body <> ")"
  Abs ident body -> "Abs(" <> printTcoExprShape body <> ")"
  App fn arg -> "App(" <> printTcoExprShape fn <> ")"
  Branch branches def -> "Branch(" <> String.joinWith ", " (map (\(Pair _ expr) -> printTcoExprShape expr) (toArray branches)) <> ", def=" <> printTcoExprShape def <> ")"
  Var _ -> "Var"
  LetRec _ _ body -> "LetRec(" <> printTcoExprShape body <> ")"
  Lit (LitInt _) -> "LitInt"
  Lit (LitNumber _) -> "LitNumber"
  Lit (LitString _) -> "LitString"
  Lit (LitChar _) -> "LitChar"
  Lit (LitBoolean _) -> "LitBoolean"
  Lit (LitArray _) -> "LitArray"
  Lit (LitRecord _) -> "LitRecord"
  UncurriedAbs _ body -> "UncurriedAbs(" <> printTcoExprShape body <> ")"
  UncurriedEffectAbs _ body -> "UncurriedEffectAbs(" <> printTcoExprShape body <> ")"
  UncurriedApp fn args -> "UncurriedApp(" <> printTcoExprShape fn <> ")"
  UncurriedEffectApp fn args -> "UncurriedEffectApp(" <> printTcoExprShape fn <> ")"
  EffectBind _ _ _ body -> "EffectBind(" <> printTcoExprShape body <> ")"
  EffectPure _ -> "EffectPure"
  Typed tp inner -> "Typed(" <> printExprType tp <> ", " <> printTcoExprShape inner <> ")"
  _ -> "Other"

extractExprFuncType :: ExprType -> Maybe { fArgs :: Array ExprType, fRet :: ExprType }
extractExprFuncType ty =
  let
    flattenFuncType acc (Func args ret) = flattenFuncType (acc <> args) ret
    flattenFuncType acc ret = { fArgs: acc, fRet: ret }

    getFunc (Func a r) = Just (flattenFuncType a r)
    getFunc (ConstrainedType constraints innerTy) =
      case getFunc innerTy of
        Just i ->
          let
            constraintTypes = map
              ( \(Tuple qual args) ->
                  let
                    qualStr = String.joinWith "." qual
                  in
                    ADT qualStr qual args
              )
              constraints
          in
            Just (i { fArgs = constraintTypes <> i.fArgs })
        Nothing -> Nothing
    getFunc (ForAll _ innerTy) = getFunc innerTy
    getFunc _ = Nothing
  in
    getFunc ty

extractFuncType :: TcoExpr -> Maybe { fArgs :: Array ExprType, fRet :: ExprType }
extractFuncType (TcoExpr _ (Typed ty inner)) =
  case extractExprFuncType ty of
    Just r -> Just r
    Nothing -> extractFuncType inner
extractFuncType _ = Nothing

-- Require the resolved intrinsic and concrete scalar callback. Outer Typed
-- annotations on the seed can describe the enclosing application, so only a
-- literal establishes its Int type here; opaque/dynamic seeds keep the copies.
isIntArrayFold :: TcoExpr -> Array TcoExpr -> Boolean
isIntArrayFold fn args = case unwrapTcoExpr fn, args of
  Var (Qualified (Just (ModuleName "Data.Foldable")) (Ident "foldlArray")), [ callback, seed, _ ] ->
    case extractFuncType callback, unwrapTcoExpr seed of
      Just { fArgs: [ Int, Int ], fRet: Int }, Lit (LitInt _) -> true
      _, _ -> false
  _, _ -> false

normalizeFreshIntArrayRoundtrip :: String -> GoExpr -> GoExpr
normalizeFreshIntArrayRoundtrip suffix expr = case expr of
  GoBoxIntArray (GoUnboxIntArray (GoCall (GoSelector (GoVar "gopurs_runtime") "Array") [ GoFreshFilterArray filtered ])) ->
    let
      sourceName = "source_int_array_" <> suffix
      itemsName = "items_int_array_" <> suffix
      indexName = "i_int_array_" <> suffix
      valueName = "v_int_array_" <> suffix
      source = GoCall (GoSelector (GoVar "gopurs_runtime") "Array") [ filtered ]
      -- The source runs once, before these IIFE-local names enter scope. The
      -- filter's make/append owns this buffer; do not search through variables,
      -- calls or storage for a marker. Normalize every Value before the fold
      -- to preserve IntVal/tag/pointer semantics of the two original copies.
      body = GoBlock
        [ GoAssign itemsName (GoCall (GoRaw "(*[]gopurs_runtime.Value)") [ GoSelector (GoVar sourceName) "UnsafePtr" ])
        , GoForRange (indexName <> ", " <> valueName <> " := range *" <> itemsName)
            [ GoMutate ("(*" <> itemsName <> ")[" <> indexName <> "]")
                (GoCall (GoSelector (GoVar "gopurs_runtime") "Int") [ GoSelector (GoVar valueName) "IntVal" ])
            ]
        , GoReturn (GoVar sourceName)
        ]
    in
      GoIIFE sourceName source body
  _ -> expr

getExprType :: TcoExpr -> ExprType
getExprType (TcoExpr _ syn) = case syn of
  Typed t _ -> t
  PrimOp op -> case op of
    Op1 OpIntNegate _ -> Int
    Op1 OpIntBitNot _ -> Int
    Op1 OpNumberNegate _ -> Number
    Op1 OpBooleanNot _ -> Boolean
    Op1 (OpIsTag _) _ -> Boolean
    Op1 OpArrayLength _ -> Int
    Op2 (OpIntNum _) _ _ -> Int
    Op2 (OpIntOrd _) _ _ -> Boolean
    Op2 OpIntBitZeroFillShiftRight _ _ -> Int
    Op2 (OpNumberNum _) _ _ -> Number
    Op2 (OpNumberOrd _) _ _ -> Boolean
    Op2 OpStringAppend _ _ -> String
    Op2 (OpStringOrd _) _ _ -> Boolean
    Op2 (OpCharOrd _) _ _ -> Boolean
    Op2 (OpBooleanOrd _) _ _ -> Boolean
    Op2 OpBooleanAnd _ _ -> Boolean
    Op2 OpBooleanOr _ _ -> Boolean
    _ -> Any
  _ -> Any

executeIfOpaque :: TcoExpr -> GoExpr -> GoExpr

executeIfOpaque expr goExpr =
  if isEffectNode expr then goExpr
  else GoCall (GoSelector (GoVar "gopurs_runtime") "Apply") [ goExpr, GoRaw "gopurs_runtime.Value{}" ]

translateExpr :: Ref CodegenState -> Int -> String -> Array String -> ModuleFunctions -> LocalEnv -> Maybe String -> LoopContext -> ExprOptions -> Int -> TcoExpr -> ExprResult
translateExpr codegenStateRef depth modNameStr recVars moduleFunctions bound tcoIdent loopCtx options nextId tcoExpr =
  translateExprWithExpectedType codegenStateRef depth modNameStr recVars moduleFunctions bound tcoIdent loopCtx options Nothing nextId tcoExpr

translateExprWithExpectedType :: Ref CodegenState -> Int -> String -> Array String -> ModuleFunctions -> LocalEnv -> Maybe String -> LoopContext -> ExprOptions -> Maybe ExprType -> Int -> TcoExpr -> ExprResult
translateExprWithExpectedType codegenStateRef depth modNameStr recVars moduleFunctions bound tcoIdent loopCtx options@{ isTail, inEffectBlock } mbExpectedExprType nextId tcoExpr@(TcoExpr tcoAnalysis expr) =
  let
    elidedCtors = (unsafePerformEffect (Ref.read codegenStateRef)).elidedCtors
    isEff = isEffectNode tcoExpr
  in
    if isEff && not inEffectBlock then
      let
        res = translateExpr codegenStateRef depth modNameStr recVars moduleFunctions bound tcoIdent loopCtx { isTail: false, inEffectBlock: true } nextId tcoExpr
        funcExpr = GoRaw ("gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {\n" <> printGoExpr (GoBlock (flattenStmts res.stmts <> [ GoReturn (boxGoExpr codegenStateRef modNameStr res.expr res.exprType) ])) <> "\n})")
      in
        { stmts: StmtEmpty, expr: funcExpr, exprType: TypeValue, nextId: res.nextId }
    else
      let
        liftIfNeeded mkNodeThunk = mkNodeThunk unit
      in
        case expr of
          Typed type_ a ->
            let
              expectedGoType = exprTypeToGoType (unsafePerformEffect (Ref.read codegenStateRef)).pointerAdtPaths (unsafePerformEffect (Ref.read codegenStateRef)).enumAdts (unsafePerformEffect (Ref.read codegenStateRef)).elidedCtors modNameStr type_
              h = unsafePerformEffect (Ref.read codegenStateRef)

            in
              case unwrapTcoExpr a, expectedGoType of
                Lit (LitRecord props), TypeStructPointer baseStructName fullName fullPath tArgs ->
                  case Map.lookup fullName h.classDeclsFields of
                    Just classInfo ->
                      let
                        classFields = classInfo.fields
                        typeArgs = case type_ of
                          ADT fullName _ tArgs ->
                            let
                              mapped = map (exprTypeToGoType h.pointerAdtPaths h.enumAdts h.elidedCtors modNameStr) tArgs
                              arity = case Map.lookup fullName h.pointerAdtPaths of
                                Just info -> info.arity
                                Nothing -> Array.length classInfo.vars
                            in
                              Array.take arity mapped
                          _ -> map (const TypeValue) classInfo.vars
                        instMap = Map.fromFoldable (Array.zip classInfo.vars typeArgs)
                        propMap = Map.fromFoldable (map (\(Prop k v) -> Tuple k v) props)
                        sortedVals = Array.mapMaybe (\f -> map (\v -> { field: f, val: v }) (Map.lookup f.name propMap)) classFields
                        accProps = foldl
                          ( \acc item ->
                              let
                                genericGoType = exprTypeToGenericGoType h.pointerAdtPaths h.enumAdts h.elidedCtors classInfo.vars modNameStr item.field."type"
                                expectedType = instantiateGenericGoType instMap genericGoType
                                expectedExprType = item.field."type"

                                newBound = case unwrapTcoExpr item.val, extractExprFuncType expectedExprType of
                                  Abs args _, Just { fArgs } ->
                                    let
                                      paramsWithTypes = Array.zipWith (\(Tuple mbI lvl) fArgTy -> Tuple (localId mbI lvl) (exprTypeToGoType h.pointerAdtPaths h.enumAdts h.elidedCtors modNameStr fArgTy)) (toArray args) (fArgs <> Array.replicate (Array.length (toArray args) - Array.length fArgs) Any)
                                    in
                                      foldl (\b (Tuple idStr goType) -> Map.insert idStr { name: idStr, goType } b) bound paramsWithTypes
                                  UncurriedAbs args _, Just { fArgs } ->
                                    let
                                      paramsWithTypes = Array.zipWith (\(Tuple mbI lvl) fArgTy -> Tuple (localId mbI lvl) (exprTypeToGoType h.pointerAdtPaths h.enumAdts h.elidedCtors modNameStr fArgTy)) args (fArgs <> Array.replicate (Array.length args - Array.length fArgs) Any)
                                    in
                                      foldl (\b (Tuple idStr goType) -> Map.insert idStr { name: idStr, goType } b) bound paramsWithTypes
                                  _, _ -> bound

                                resVal = translateExpr codegenStateRef (depth + 1) modNameStr recVars moduleFunctions newBound Nothing [] { isTail: false, inEffectBlock: false } acc.nextId item.val
                                coercedExpr = coerceGoExpr codegenStateRef modNameStr resVal.expr resVal.exprType expectedType
                              in
                                { stmts: acc.stmts <> resVal.stmts, exprs: Array.snoc acc.exprs coercedExpr, exprType: TypeValue, nextId: resVal.nextId }
                          )
                          { stmts: StmtEmpty, exprs: [], exprType: TypeValue, nextId }
                          sortedVals
                      in
                        let
                          typeArgsForDict = case expectedGoType of
                            TypeStructPointer _ _ _ args -> args
                            _ -> []
                          monoStructName = case String.indexOf (Pattern "[") fullPath of
                            Just idx -> String.take idx fullPath
                            Nothing -> fullPath
                        in
                          { stmts: accProps.stmts, expr: GoConstructor (hashString baseStructName) monoStructName typeArgsForDict accProps.exprs, exprType: expectedGoType, nextId: accProps.nextId }
                    Nothing ->
                      let
                        res = translateExprWithExpectedType codegenStateRef depth modNameStr recVars moduleFunctions bound tcoIdent loopCtx options (Just type_) nextId a
                      in
                        case res.exprType of
                          TypeStructPointer _ _ _ _ -> res
                          _ ->
                            if expectedGoType == res.exprType then res
                            else if isClosureNode codegenStateRef a then res
                            else
                              { stmts: res.stmts, expr: coerceGoExpr codegenStateRef modNameStr res.expr res.exprType expectedGoType, exprType: expectedGoType, nextId: res.nextId }
                Let ident lvl val body, _ ->
                  let
                    newBody = case body of
                      TcoExpr bAnn _ -> TcoExpr bAnn (Typed type_ body)
                    newLetShape = Let ident lvl val newBody
                    newA = case a of
                      TcoExpr ann _ -> TcoExpr ann newLetShape
                  in
                    translateExpr codegenStateRef depth modNameStr recVars moduleFunctions bound tcoIdent loopCtx options nextId newA
                LetRec lvl bindings body, _ ->
                  let
                    newBody = case body of
                      TcoExpr bAnn _ -> TcoExpr bAnn (Typed type_ body)
                    newLetShape = LetRec lvl bindings newBody
                    newA = case a of
                      TcoExpr ann _ -> TcoExpr ann newLetShape
                  in
                    translateExpr codegenStateRef depth modNameStr recVars moduleFunctions bound tcoIdent loopCtx options nextId newA
                _, _ ->
                  let
                    res = translateExprWithExpectedType codegenStateRef depth modNameStr recVars moduleFunctions bound tcoIdent loopCtx options (Just type_) nextId a
                  in
                    case res.exprType of
                      TypeStructPointer _ _ _ _ -> res
                      _ ->
                        if expectedGoType == res.exprType then res
                        else if isClosureNode codegenStateRef a then res
                        else
                          { stmts: res.stmts, expr: coerceGoExpr codegenStateRef modNameStr res.expr res.exprType expectedGoType, exprType: expectedGoType, nextId: res.nextId }
          Var (Qualified mbMn (Ident i)) ->
            let
              safeName = sanitizeName i
              h = unsafePerformEffect (Ref.read codegenStateRef)
              vType = case mbMn of
                Just mn ->
                  let
                    modStr = unwrap mn
                  in
                    case Map.lookup (modStr <> "." <> i) h.globalTypes of
                      Just ty -> exprTypeToGoType h.pointerAdtPaths h.enumAdts h.elidedCtors modNameStr ty
                      Nothing -> TypeValue
                Nothing -> TypeValue
            in
              case mbMn of
                Just mn ->
                  let
                    modStr = unwrap mn
                    modPkg = String.replaceAll (Pattern ".") (Replacement "_") modStr
                    rawCall = GoCall (GoVar ("Get_" <> modPkg <> "_" <> safeName)) []
                  in
                    { stmts: StmtEmpty, expr: unboxGoExpr codegenStateRef modNameStr rawCall TypeValue vType, exprType: vType, nextId }
                Nothing ->
                  let
                    rawCall = Debug.trace ("mbMn is Nothing for safeName: " <> safeName) (\_ -> GoCall (GoVar ("Get_" <> modNameStr <> "_" <> safeName)) [])
                  in
                    { stmts: StmtEmpty, expr: unboxGoExpr codegenStateRef modNameStr rawCall TypeValue vType, exprType: vType, nextId }

          Local mbIdent lvl ->
            let
              v = fromMaybe { name: localId mbIdent lvl, goType: TypeValue } (Map.lookup (localId mbIdent lvl) bound)
            in
              { stmts: StmtEmpty, expr: GoVar v.name, exprType: v.goType, nextId }

          Lit value
            | Just result <- PrimitiveExprs.literal value ->
                { stmts: StmtEmpty, expr: result.expr, exprType: result.exprType, nextId }

          Lit (LitArray xs) ->
            let
              accXs = foldl
                ( \acc val ->
                    let
                      resVal = translateExpr codegenStateRef (depth + 1) modNameStr recVars moduleFunctions bound Nothing [] { isTail: false, inEffectBlock: false } acc.nextId val
                    in
                      { stmts: acc.stmts <> resVal.stmts, exprs: Array.snoc acc.exprs resVal.expr, exprTypes: Array.snoc acc.exprTypes resVal.exprType, nextId: resVal.nextId }
                )
                { stmts: StmtEmpty, exprs: [], exprTypes: [], nextId }
                xs

              mbElemType = Array.head accXs.exprTypes
              isAllSame = Array.all (\t -> Just t == mbElemType) accXs.exprTypes
              
              expectedElemType = case mbExpectedExprType of
                Just exTy ->
                  let
                    h = unsafePerformEffect (Ref.read codegenStateRef)
                  in
                    case exprTypeToGenericGoType h.pointerAdtPaths h.enumAdts h.elidedCtors [] modNameStr exTy of
                      TypeNativeArray et -> Just et
                      _ -> Nothing
                Nothing -> Nothing
                
              finalElemType = case expectedElemType of
                Just t -> Just t
                Nothing -> if isAllSame then mbElemType else Nothing
            in
              case finalElemType of
                Just elemType | (isAllSame || Array.length xs == 0) && elemType /= TypeValue ->
                  let
                    goTypeArr = TypeNativeArray elemType
                  in
                    { stmts: accXs.stmts, expr: GoRaw (goTypeToStr goTypeArr <> "{" <> String.joinWith ", " (map printGoExpr accXs.exprs) <> "}"), exprType: goTypeArr, nextId: accXs.nextId }
                _ ->
                  let
                    boxedExprs = Array.zipWith (\expr ty -> boxGoExpr codegenStateRef modNameStr expr ty) accXs.exprs accXs.exprTypes
                  in
                    { stmts: accXs.stmts, expr: GoCall (GoSelector (GoVar "gopurs_runtime") "Array") [ GoRaw ("[]gopurs_runtime.Value{" <> String.joinWith ", " (map printGoExpr boxedExprs) <> "}") ], exprType: TypeValue, nextId: accXs.nextId }

          Lit (LitRecord props) ->
            let
              sortedProps = Array.sortBy (comparing \(Prop k _) -> k) props
              baseExprType = getExprType tcoExpr
              exprType = case baseExprType of
                Record _ -> baseExprType
                _ -> fromMaybe baseExprType mbExpectedExprType

              mbRecordType = case exprType of
                Record (Row fields _) -> Just fields
                _ -> Nothing

              goRecordType = exprTypeToGoType (unsafePerformEffect (Ref.read codegenStateRef)).pointerAdtPaths (unsafePerformEffect (Ref.read codegenStateRef)).enumAdts (unsafePerformEffect (Ref.read codegenStateRef)).elidedCtors modNameStr exprType
              
              recordFields = case mbRecordType of
                Just fields -> Map.fromFoldable fields
                Nothing -> Map.empty

              accProps = foldl
                ( \acc (Prop key val) ->
                    let
                      expectedExprType = fromMaybe Any (Map.lookup key recordFields)
                      newBound = case unwrapTcoExpr val, extractExprFuncType expectedExprType of
                        Abs args _, Just { fArgs } ->
                          let
                            paramsWithTypes = Array.zipWith (\(Tuple mbI lvl) fArgTy -> Tuple (localId mbI lvl) (exprTypeToGoType (unsafePerformEffect (Ref.read codegenStateRef)).pointerAdtPaths (unsafePerformEffect (Ref.read codegenStateRef)).enumAdts (unsafePerformEffect (Ref.read codegenStateRef)).elidedCtors modNameStr fArgTy)) (toArray args) (fArgs <> Array.replicate (Array.length (toArray args) - Array.length fArgs) Any)
                          in
                            foldl (\b (Tuple idStr goType) -> Map.insert idStr { name: idStr, goType } b) bound paramsWithTypes
                        UncurriedAbs args _, Just { fArgs } ->
                          let
                            paramsWithTypes = Array.zipWith (\(Tuple mbI lvl) fArgTy -> Tuple (localId mbI lvl) (exprTypeToGoType (unsafePerformEffect (Ref.read codegenStateRef)).pointerAdtPaths (unsafePerformEffect (Ref.read codegenStateRef)).enumAdts (unsafePerformEffect (Ref.read codegenStateRef)).elidedCtors modNameStr fArgTy)) args (fArgs <> Array.replicate (Array.length args - Array.length fArgs) Any)
                          in
                            foldl (\b (Tuple idStr goType) -> Map.insert idStr { name: idStr, goType } b) bound paramsWithTypes
                        _, _ -> bound

                      resVal = translateExpr codegenStateRef (depth + 1) modNameStr recVars moduleFunctions newBound Nothing [] { isTail: false, inEffectBlock: false } acc.nextId val

                      expectedGoType = exprTypeToGoType (unsafePerformEffect (Ref.read codegenStateRef)).pointerAdtPaths (unsafePerformEffect (Ref.read codegenStateRef)).enumAdts (unsafePerformEffect (Ref.read codegenStateRef)).elidedCtors modNameStr expectedExprType
                      coercedVal = coerceGoExpr codegenStateRef modNameStr resVal.expr resVal.exprType expectedGoType
                    in
                      { stmts: acc.stmts <> resVal.stmts, exprs: Array.snoc acc.exprs (Tuple key coercedVal), exprType: TypeValue, nextId: resVal.nextId }
                )
                { stmts: StmtEmpty, exprs: [], exprType: TypeValue, nextId }
                sortedProps
            in
              { stmts: accProps.stmts, expr: GoRecordDict goRecordType accProps.exprs, exprType: goRecordType, nextId: accProps.nextId }

          expr_ | (case expr_ of
                     App _ _ -> true
                     Syn.TypeApp _ _ -> true
                     _ -> false) ->
            let
              Tuple flatFn flatArgsSpine = collectGoSpine tcoExpr
              flatArgs = getGoSpineArgs flatArgsSpine

              isTailCallTo =
                if isTail then case unwrapTcoExpr flatFn of
                  Local mbIdent lvl ->
                    let
                      v = fromMaybe { name: localId mbIdent lvl, goType: TypeValue } (Map.lookup (localId mbIdent lvl) bound)
                    in
                      Array.findIndex (\ctx -> ctx.ident == v.name) loopCtx
                  Var (Qualified mbMod (Ident name)) ->
                    let
                      fullName = sanitizeName name
                    in
                      Array.findIndex (\ctx -> ctx.ident == fullName) loopCtx
                  _ -> Nothing
                else Nothing

            in
              case isTailCallTo of
                Just index ->
                  let
                    accFinal = foldl
                      ( \acc arg ->
                          let
                            argRes = translateExpr codegenStateRef (depth + 1) modNameStr recVars moduleFunctions bound Nothing [] { isTail: false, inEffectBlock: false } acc.nextId arg
                          in
                            { stmts: acc.stmts <> argRes.stmts, exprs: Array.snoc acc.exprs argRes.expr, exprTypes: Array.snoc acc.exprTypes argRes.exprType, nextId: argRes.nextId }
                      )
                      { stmts: StmtEmpty, exprs: [], exprTypes: [], nextId }
                      flatArgs
                    targetCtx = fromMaybe { ident: "", params: [], loopParams: [], goTypes: [], fRet: TypeValue } (Array.index loopCtx index)
                    assigns = Array.mapWithIndex
                      ( \i paramName ->
                          let
                            argExpr = fromMaybe (GoRaw "nil") (Array.index accFinal.exprs i)
                            argType = fromMaybe TypeValue (Array.index accFinal.exprTypes i)
                            expectedType = fromMaybe TypeValue (Array.index targetCtx.goTypes i)
                          in
                            GoMutate paramName (coerceGoExpr codegenStateRef modNameStr argExpr argType expectedType)
                      )
                      targetCtx.loopParams
                  in
                    let
                      expectedGoType = exprTypeToGoType (unsafePerformEffect (Ref.read codegenStateRef)).pointerAdtPaths (unsafePerformEffect (Ref.read codegenStateRef)).enumAdts (unsafePerformEffect (Ref.read codegenStateRef)).elidedCtors modNameStr (case getExprType tcoExpr of
                                Any -> fromMaybe Any mbExpectedExprType
                                ty -> ty)
                      expectedGoTypeStr = goTypeToStr expectedGoType
                    in
                    { stmts: accFinal.stmts <> foldMap StmtLeaf assigns <> StmtLeaf (GoContinue targetCtx.ident), expr: GoRaw ("func() " <> expectedGoTypeStr <> " { panic(\"unreachable\") }()"), exprType: expectedGoType, nextId: accFinal.nextId }

                Nothing ->
                  let
                    getVar :: BackendSyntax TcoExpr -> Maybe { mbMod :: Maybe ModuleName, name :: String }
                    getVar (Typed _ inner) = getVar (unwrapTcoExpr inner)
                    getVar (Var (Qualified mbMod (Ident name))) = Just { mbMod, name }
                    getVar (Local mbIdent lvl) =
                      let
                        resolvedName = (fromMaybe { name: localId mbIdent lvl, goType: TypeValue } (Map.lookup (localId mbIdent lvl) bound)).name
                      in
                        Just { mbMod: Nothing, name: resolvedName }
                    getVar (Lit _) = Just { mbMod: Nothing, name: "Lit" }
                    getVar (App _ _) = Just { mbMod: Nothing, name: "App" }
                    getVar (Abs _ _) = Just { mbMod: Nothing, name: "Abs" }
                    getVar (UncurriedApp _ _) = Just { mbMod: Nothing, name: "UncurriedApp" }
                    getVar (UncurriedAbs _ _) = Just { mbMod: Nothing, name: "UncurriedAbs" }
                    getVar (UncurriedEffectApp _ _) = Just { mbMod: Nothing, name: "UncurriedEffectApp" }
                    getVar (UncurriedEffectAbs _ _) = Just { mbMod: Nothing, name: "UncurriedEffectAbs" }
                    getVar (Accessor _ _) = Just { mbMod: Nothing, name: "Accessor" }
                    getVar (Update _ _) = Just { mbMod: Nothing, name: "Update" }
                    getVar (CtorSaturated _ _ _ _ _) = Just { mbMod: Nothing, name: "CtorSaturated" }
                    getVar (CtorDef _ _ _ _) = Just { mbMod: Nothing, name: "CtorDef" }
                    getVar (LetRec _ _ _) = Just { mbMod: Nothing, name: "LetRec" }
                    getVar (Let _ _ _ _) = Just { mbMod: Nothing, name: "Let" }
                    getVar (EffectBind _ _ _ _) = Just { mbMod: Nothing, name: "EffectBind" }
                    getVar (EffectPure _) = Just { mbMod: Nothing, name: "EffectPure" }
                    getVar (EffectDefer _) = Just { mbMod: Nothing, name: "EffectDefer" }
                    getVar _ = Just { mbMod: Nothing, name: "Unknown" }

                    mbIntrinsic = case getVar (unwrapTcoExpr flatFn) of
                      Just { name: "arrayMap" } ->
                        if Array.length flatArgs >= 2 then Just "arrayMap" else Nothing
                      Just { name: "foldlArray" } ->
                        if Array.length flatArgs >= 3 then Just "foldlArray" else Nothing
                      Just { mbMod, name: "filter" } | mbMod == Just (ModuleName "Data.Array") || (mbMod == Nothing && modNameStr == "Data.Array") ->
                        if Array.length flatArgs >= 2 then Just "filter" else Nothing
                      _ -> Nothing

                    mbDirectCall = case getVar (unwrapTcoExpr flatFn) of
                      Just { mbMod, name } ->
                        let
                          isLocal = map (String.replaceAll (Pattern ".") (Replacement "_") <<< unwrap) mbMod == Just modNameStr || mbMod == Nothing
                          modPrefix = case mbMod of
                            Just mn -> String.replaceAll (Pattern ".") (Replacement "_") (unwrap mn)
                            Nothing -> modNameStr
                          fromModuleFunctions = if isLocal then Map.lookup name moduleFunctions else Nothing
                          fromTypeSig = case extractFuncType flatFn of
                            Just { fArgs, fRet } ->
                              Just { fullName: "Call_" <> modPrefix <> "_" <> sanitizeName name, fArgs: map (exprTypeToGoType (unsafePerformEffect (Ref.read codegenStateRef)).pointerAdtPaths (unsafePerformEffect (Ref.read codegenStateRef)).enumAdts (unsafePerformEffect (Ref.read codegenStateRef)).elidedCtors modNameStr) fArgs, fRet: exprTypeToGoType (unsafePerformEffect (Ref.read codegenStateRef)).pointerAdtPaths (unsafePerformEffect (Ref.read codegenStateRef)).enumAdts (unsafePerformEffect (Ref.read codegenStateRef)).elidedCtors modNameStr fRet, arity: Array.length fArgs }
                            Nothing ->
                              Nothing

                          entry = case fromTypeSig of
                            Just e | not isLocal -> Just e
                            _ -> fromModuleFunctions
                        in
                          case entry of
                            Just e ->
                              if Array.length flatArgs >= e.arity && e.arity >= 1 then Just e else Nothing
                            Nothing -> Nothing
                      Nothing -> Nothing
                  in
                    case mbIntrinsic of
                      Just intrinsicName ->
                        let
                          accArgs = foldl
                            ( \acc arg ->
                                let
                                  argRes = translateExpr codegenStateRef (depth + 1) modNameStr recVars moduleFunctions bound Nothing [] { isTail: false, inEffectBlock: false } acc.nextId arg
                                in
                                  { stmts: acc.stmts <> argRes.stmts, exprs: Array.snoc acc.exprs (boxGoExpr codegenStateRef modNameStr argRes.expr argRes.exprType), exprTypes: Array.snoc acc.exprTypes argRes.exprType, nextId: argRes.nextId }
                            )
                            { stmts: StmtEmpty, exprs: [], exprTypes: [], nextId }
                            flatArgs

                          iifeName = intrinsicName <> show depth
                          arrValName = "arr_val_" <> iifeName
                          arrGoName = "arr_go_" <> iifeName
                          resGoName = "res_go_" <> iifeName
                          iName = "i_" <> iifeName
                          vName = "v_" <> iifeName

                          iifeExpr = case intrinsicName of
                            "arrayMap" ->
                              let
                                fExpr = fromMaybe (GoRaw "nil") (Array.index accArgs.exprs 0)
                                arrExpr = fromMaybe (GoRaw "nil") (Array.index accArgs.exprs 1)
                                loopBody = GoMutate (resGoName <> "[" <> iName <> "]") (GoCall (GoSelector (GoVar "gopurs_runtime") "Apply") [ fExpr, GoVar vName ])
                                iifeBody = GoBlock
                                  [ GoAssign arrGoName (GoCall (GoRaw "(*[]gopurs_runtime.Value)") [ GoSelector (GoVar arrValName) "UnsafePtr" ])
                                  , GoAssign resGoName (GoCall (GoVar "make") [ GoRaw "[]gopurs_runtime.Value", GoCall (GoVar "len") [ GoRaw ("*" <> arrGoName) ] ])
                                  , GoForRange (iName <> ", " <> vName <> " := range *" <> arrGoName) [ loopBody ]
                                  , GoReturn (GoCall (GoSelector (GoVar "gopurs_runtime") "Array") [ GoVar resGoName ])
                                  ]
                              in
                                GoIIFE arrValName arrExpr iifeBody

                            "foldlArray" ->
                              let
                                fExpr = fromMaybe (GoRaw "nil") (Array.index accArgs.exprs 0)
                                initExpr = fromMaybe (GoRaw "nil") (Array.index accArgs.exprs 1)
                                boxedArrExpr = fromMaybe (GoRaw "nil") (Array.index accArgs.exprs 2)
                                arrExpr = if isIntArrayFold flatFn flatArgs then
                                  normalizeFreshIntArrayRoundtrip (iifeName <> "_" <> show accArgs.nextId) boxedArrExpr
                                else boxedArrExpr
                                loopBody = GoMutate resGoName (GoCall (GoSelector (GoVar "gopurs_runtime") "Apply2") [ fExpr, GoVar resGoName, GoVar vName ])
                                iifeBody = GoBlock
                                  [ GoAssign resGoName initExpr
                                  , GoAssign arrGoName (GoCall (GoRaw "(*[]gopurs_runtime.Value)") [ GoSelector (GoVar arrValName) "UnsafePtr" ])
                                  , GoForRange ("_, " <> vName <> " := range *" <> arrGoName) [ loopBody ]
                                  , GoReturn (GoVar resGoName)
                                  ]
                              in
                                GoIIFE arrValName arrExpr iifeBody

                            "filter" ->
                              let
                                fExpr = fromMaybe (GoRaw "nil") (Array.index accArgs.exprs 0)
                                arrExpr = fromMaybe (GoRaw "nil") (Array.index accArgs.exprs 1)
                                condExpr = GoCall (GoSelector (GoVar "gopurs_runtime") "Apply") [ fExpr, GoVar vName ]
                                isTrueExpr = GoCall (GoSelector condExpr "BoolVal") []
                                loopBody = GoIfElse isTrueExpr [ GoMutate resGoName (GoCall (GoVar "append") [ GoVar resGoName, GoVar vName ]) ] []
                                iifeBody = GoBlock
                                  [ GoAssign arrGoName (GoCall (GoRaw "(*[]gopurs_runtime.Value)") [ GoSelector (GoVar arrValName) "UnsafePtr" ])
                                  , GoAssign resGoName (GoCall (GoVar "make") [ GoRaw "[]gopurs_runtime.Value", GoRaw "0" ])
                                  , GoForRange ("_, " <> vName <> " := range *" <> arrGoName) [ loopBody ]
                                  , GoReturn (GoCall (GoSelector (GoVar "gopurs_runtime") "Array") [ GoVar resGoName ])
                                  ]
                              in
                                GoIIFE arrValName arrExpr iifeBody

                            _ -> GoRaw "nil"

                          arity = if intrinsicName == "foldlArray" then 3 else 2
                          accArgsRemaining = Array.drop arity accArgs.exprs
                          accArgsRemainingTypes = Array.drop arity accArgs.exprTypes
                          accArgsRemainingBoxed = Array.zipWith (\arg t -> boxGoExpr codegenStateRef modNameStr arg t) accArgsRemaining accArgsRemainingTypes

                          buildApp :: GoExpr -> Array GoExpr -> GoExpr
                          buildApp fExpr argExprs =
                            let
                              len = Array.length argExprs
                            in
                              if len == 0 then fExpr
                              else if len == 1 then GoCall (GoSelector (GoVar "gopurs_runtime") "Apply") [ fExpr, fromMaybe (GoRaw "nil") (Array.index argExprs 0) ]
                              else if len >= 2 && len <= 10 then
                                GoCall (GoSelector (GoVar "gopurs_runtime") ("Apply" <> show len)) (Array.cons fExpr argExprs)
                              else
                                let
                                  chunk = Array.take 10 argExprs
                                  rest = Array.drop 10 argExprs
                                in
                                  buildApp (buildApp fExpr chunk) rest

                          finalExpr = buildApp iifeExpr accArgsRemainingBoxed
                        in
                          { stmts: accArgs.stmts, expr: finalExpr, exprType: TypeValue, nextId: accArgs.nextId }

                      Nothing ->
                        let
                          buildApp :: GoExpr -> Array GoExpr -> GoExpr
                          buildApp fExpr argExprs =
                            let
                              len = Array.length argExprs
                            in
                              if len == 0 then fExpr
                              else if len == 1 then GoCall (GoSelector (GoVar "gopurs_runtime") "Apply") [ fExpr, fromMaybe (GoRaw "nil") (Array.index argExprs 0) ]
                              else if len >= 2 && len <= 10 then
                                GoCall (GoSelector (GoVar "gopurs_runtime") ("Apply" <> show len)) (Array.cons fExpr argExprs)
                              else
                                let
                                  chunk = Array.take 10 argExprs
                                  rest = Array.drop 10 argExprs
                                in
                                  buildApp (buildApp fExpr chunk) rest
                        in
                          case mbDirectCall of
                            Just { fullName, fArgs, fRet, arity } ->
                              let
                                accArgs = foldl
                                  ( \acc arg ->
                                      let
                                        argRes = translateExpr codegenStateRef (depth + 1) modNameStr recVars moduleFunctions bound Nothing [] { isTail: false, inEffectBlock: false } acc.nextId arg
                                      in
                                        { stmts: acc.stmts <> argRes.stmts, exprs: Array.snoc acc.exprs argRes.expr, exprTypes: Array.snoc acc.exprTypes argRes.exprType, nextId: argRes.nextId }
                                  )
                                  { stmts: StmtEmpty, exprs: [], exprTypes: [], nextId }
                                  flatArgs

                                accArgsArity = Array.take arity accArgs.exprs
                                accArgsRemaining = Array.drop arity accArgs.exprs
                                accArgsRemainingTypes = Array.drop arity accArgs.exprTypes
                                accArgsRemainingBoxed = Array.zipWith (\arg t -> boxGoExpr codegenStateRef modNameStr arg t) accArgsRemaining accArgsRemainingTypes

                                callArgs = Array.mapWithIndex
                                  ( \i argExprValue ->
                                      let
                                        expectedType = fromMaybe TypeValue (Array.index fArgs i)
                                        actualType = fromMaybe TypeValue (Array.index accArgs.exprTypes i)
                                      in
                                        coerceGoExpr codegenStateRef modNameStr argExprValue actualType expectedType
                                  )
                                  accArgsArity

                                callExpr = GoCall (GoVar fullName) callArgs
                                finalExpr = if Array.length accArgsRemainingBoxed == 0 then callExpr else buildApp (boxGoExpr codegenStateRef modNameStr callExpr fRet) accArgsRemainingBoxed
                                finalExprType = if Array.length accArgsRemainingBoxed == 0 then fRet else TypeValue
                              in
                                { stmts: accArgs.stmts, expr: finalExpr, exprType: finalExprType, nextId: accArgs.nextId }

                            Nothing ->
                              let
                                resFn = translateExpr codegenStateRef (depth + 1) modNameStr recVars moduleFunctions bound Nothing [] { isTail: false, inEffectBlock: false } nextId flatFn
                                accArgs = foldl
                                  ( \acc arg ->
                                      let
                                        argRes = translateExpr codegenStateRef (depth + 1) modNameStr recVars moduleFunctions bound Nothing [] { isTail: false, inEffectBlock: false } acc.nextId arg
                                      in
                                        { stmts: acc.stmts <> argRes.stmts, exprs: Array.snoc acc.exprs argRes.expr, exprTypes: Array.snoc acc.exprTypes argRes.exprType, nextId: argRes.nextId }
                                  )
                                  { stmts: resFn.stmts, exprs: [], exprTypes: [], nextId: resFn.nextId }
                                  flatArgs

                                finalExprType = case resFn.exprType of
                                  TypeFunc fArgs fRet | Array.length fArgs == Array.length flatArgs -> fRet
                                  _ -> TypeValue

                                finalExpr = case resFn.exprType of
                                  TypeFunc fArgs fRet | Array.length fArgs == Array.length flatArgs ->
                                    let
                                      callArgs = Array.mapWithIndex
                                        ( \i expected ->
                                            let
                                              arg = fromMaybe (GoRaw "nil") (Array.index accArgs.exprs i)
                                              actual = fromMaybe TypeValue (Array.index accArgs.exprTypes i)
                                            in
                                              coerceGoExpr codegenStateRef modNameStr arg actual expected
                                        )
                                        fArgs
                                    in
                                      GoCall resFn.expr callArgs
                                  TypeFunc fArgs fRet | Array.length flatArgs > Array.length fArgs ->
                                    let
                                      arity = Array.length fArgs
                                      accArgsArity = Array.take arity accArgs.exprs
                                      accArgsRemaining = Array.drop arity accArgs.exprs
                                      accArgsRemainingTypes = Array.drop arity accArgs.exprTypes
                                      accArgsRemainingBoxed = Array.zipWith (\arg t -> boxGoExpr codegenStateRef modNameStr arg t) accArgsRemaining accArgsRemainingTypes

                                      callArgs = Array.mapWithIndex
                                        ( \i argExprValue ->
                                            let
                                              expectedType = fromMaybe TypeValue (Array.index fArgs i)
                                              actualType = fromMaybe TypeValue (Array.index accArgs.exprTypes i)
                                            in
                                              coerceGoExpr codegenStateRef modNameStr argExprValue actualType expectedType
                                        )
                                        accArgsArity

                                      callExpr = GoCall resFn.expr callArgs
                                    in
                                      buildApp (boxGoExpr codegenStateRef modNameStr callExpr fRet) accArgsRemainingBoxed
                                  _ ->
                                    let
                                      boxedArgs = Array.zipWith (\arg actual -> boxGoExpr codegenStateRef modNameStr arg actual) accArgs.exprs accArgs.exprTypes
                                    in
                                      buildApp (boxGoExpr codegenStateRef modNameStr resFn.expr resFn.exprType) boxedArgs
                              in
                                { stmts: accArgs.stmts, expr: finalExpr, exprType: finalExprType, nextId: accArgs.nextId }

          Abs args body ->
            let
              grouped = collectCurriedAbs args body
              consumed = NonEmptyArray.length grouped.args
              mbFuncTy = case extractFuncType tcoExpr of
                Just r -> Just r
                Nothing -> case mbExpectedExprType of
                  Just ty -> extractExprFuncType ty
                  Nothing -> Nothing

              -- A computation may still separate us from further lambdas.
              -- Consume only the parameters actually collected, not the full
              -- flattened function type's argument list.
              mbBodyType = case mbFuncTy of
                Just { fArgs, fRet } | consumed <= Array.length fArgs ->
                  Just case Array.drop consumed fArgs of
                    [] -> fRet
                    remaining -> Func remaining fRet
                _ -> Nothing

              paramsWithTypes = map (\(Tuple mbI lvl) -> Tuple (localId mbI lvl) TypeValue) (toArray grouped.args)

              newBound = foldl (\acc (Tuple idStr goType) -> Map.insert idStr { name: idStr, goType } acc) bound paramsWithTypes
              params = map fst paramsWithTypes
              resBody = translateExprWithExpectedType codegenStateRef (depth + 1) modNameStr recVars moduleFunctions newBound Nothing [] { isTail, inEffectBlock: false } mbBodyType nextId grouped.body

              buildFunc :: Array String -> GoExpr -> GoExpr
              buildFunc ps innerExpr =
                let
                  len = Array.length ps
                  bodyStr = case innerExpr of
                    GoBlock _ -> printGoExpr innerExpr
                    _ -> "return " <> printGoExpr innerExpr
                in
                  if len == 1 then
                    GoCall (GoSelector (GoVar "gopurs_runtime") "Func") [ GoRaw ("func(" <> fromMaybe "" (Array.index ps 0) <> " gopurs_runtime.Value) gopurs_runtime.Value {\n" <> bodyStr <> "\n}") ]
                  else if len >= 2 && len <= 5 then
                    let
                      goParams = String.joinWith ", " (map (\p -> p <> " gopurs_runtime.Value") ps)
                    in
                      GoCall (GoSelector (GoVar "gopurs_runtime") ("Func" <> show len)) [ GoRaw ("func(" <> goParams <> ") gopurs_runtime.Value {\n" <> bodyStr <> "\n}") ]
                  else
                    let
                      chunk = Array.take 5 ps
                      rest = Array.drop 5 ps
                    in
                      buildFunc chunk (buildFunc rest innerExpr)

              funcExpr = buildFunc params (GoBlock (flattenStmts resBody.stmts <> [ GoReturn (boxGoExpr codegenStateRef modNameStr resBody.expr resBody.exprType) ]))
            in
              { stmts: StmtEmpty, expr: funcExpr, exprType: TypeValue, nextId: resBody.nextId }

          UncurriedApp fn args ->
            let
              getVar :: BackendSyntax TcoExpr -> Maybe { mbMod :: Maybe ModuleName, name :: String }
              getVar (Typed _ inner) = getVar (unwrapTcoExpr inner)
              getVar (Var (Qualified mbMod (Ident name))) = Just { mbMod, name }
              getVar _ = Nothing

              mbIntrinsic = case getVar (unwrapTcoExpr fn) of
                Just { name: "arrayMap" } ->
                  if Array.length args >= 2 then Just "arrayMap" else Nothing
                Just { name: "foldlArray" } ->
                  if Array.length args >= 3 then Just "foldlArray" else Nothing
                Just { mbMod, name: "filterImpl" } | mbMod == Just (ModuleName "Data.Array") || (mbMod == Nothing && modNameStr == "Data.Array") ->
                  if Array.length args >= 2 then Just "filterImpl" else Nothing
                _ -> Nothing
            in
              case mbIntrinsic of
                Just intrinsicName ->
                  let
                    accArgs = foldl
                      ( \acc arg ->
                          let
                            argRes = translateExpr codegenStateRef (depth + 1) modNameStr recVars moduleFunctions bound Nothing [] { isTail: false, inEffectBlock: false } acc.nextId arg
                          in
                            { stmts: acc.stmts <> argRes.stmts, exprs: Array.snoc acc.exprs argRes.expr, exprTypes: Array.snoc acc.exprTypes argRes.exprType, nextId: argRes.nextId }
                      )
                      { stmts: StmtEmpty, exprs: [], exprTypes: [], nextId }
                      args

                    iifeName = intrinsicName <> show depth
                    arrValName = "arr_val_" <> iifeName
                    arrGoName = "arr_go_" <> iifeName
                    resGoName = "res_go_" <> iifeName
                    iName = "i_" <> iifeName
                    vName = "v_" <> iifeName

                    iifeExpr = case intrinsicName of
                      "arrayMap" ->
                        let
                          fExprRaw = fromMaybe (GoRaw "nil") (Array.index accArgs.exprs 0)
                          fExprType = fromMaybe TypeValue (Array.index accArgs.exprTypes 0)
                          arrExprRaw = fromMaybe (GoRaw "nil") (Array.index accArgs.exprs 1)
                          arrExprType = fromMaybe TypeValue (Array.index accArgs.exprTypes 1)

                          mbFnVar = case Array.index args 0 of
                            Just fArg -> getVar (unwrapTcoExpr fArg)
                            Nothing -> Nothing

                          fnFullName = case mbFnVar of
                            Just { mbMod: Just (ModuleName mn), name } -> String.replaceAll (Pattern ".") (Replacement "_") mn <> "." <> name
                            Just { mbMod: Nothing, name } -> modNameStr <> "." <> name
                            Nothing -> ""

                          mbFnArityInfo = Map.lookup fnFullName moduleFunctions

                          elemType = case arrExprType of
                            TypeNativeArray inner -> inner
                            _ -> TypeValue

                          retType = case fExprType of
                            TypeFunc _ ret -> ret
                            _ -> case mbFnArityInfo of
                              Just info -> info.fRet
                              _ -> TypeValue

                          finalRetType = TypeNativeArray retType

                          arrGoAssignment = case arrExprType of
                            TypeNativeArray _ -> GoAssign arrGoName (GoVar arrValName)
                            _ -> GoAssign arrGoName (GoCall (GoRaw "(*[]gopurs_runtime.Value)") [ GoSelector (GoVar arrValName) "UnsafePtr" ])

                          arrGoRangeTarget = case arrExprType of
                            TypeNativeArray _ -> arrGoName
                            _ -> "*" <> arrGoName

                          loopBody = case mbFnArityInfo of
                            Just info | info.arity == 1 ->
                              let
                                expectedArgType = fromMaybe TypeValue (Array.index info.fArgs 0)
                              in
                                GoMutate (resGoName <> "[" <> iName <> "]") (unboxGoExpr codegenStateRef modNameStr (GoCall (GoVar ("Call_" <> String.replaceAll (Pattern ".") (Replacement "_") fnFullName)) [ unboxGoExpr codegenStateRef modNameStr (GoVar vName) elemType expectedArgType ]) info.fRet retType)
                            _ ->
                              GoMutate (resGoName <> "[" <> iName <> "]") (unboxGoExpr codegenStateRef modNameStr (GoCall (GoSelector (GoVar "gopurs_runtime") "Apply") [ boxGoExpr codegenStateRef modNameStr fExprRaw fExprType, boxGoExpr codegenStateRef modNameStr (GoVar vName) elemType ]) TypeValue retType)

                          iifeBodyStmts =
                            [ arrGoAssignment
                            , GoAssign resGoName (GoCall (GoVar "make") [ GoRaw ("[]" <> goTypeToStr retType), GoCall (GoVar "len") [ GoRaw arrGoRangeTarget ] ])
                            , GoForRange (iName <> ", " <> vName <> " := range " <> arrGoRangeTarget) [ loopBody ]
                            ]
                        in
                          { stmts: accArgs.stmts, expr: GoCall (GoFuncLit [] (Array.cons (GoAssign arrValName arrExprRaw) (Array.cons (GoMutate "_" (GoVar arrValName)) iifeBodyStmts)) (GoVar resGoName) finalRetType) [], exprType: finalRetType, nextId: accArgs.nextId }

                      "foldlArray" ->
                        let
                          fExpr = fromMaybe (GoRaw "nil") (Array.index accArgs.exprs 0)
                          fExprType = fromMaybe TypeValue (Array.index accArgs.exprTypes 0)
                          initExpr = fromMaybe (GoRaw "nil") (Array.index accArgs.exprs 1)
                          initExprType = fromMaybe TypeValue (Array.index accArgs.exprTypes 1)
                          arrExpr = fromMaybe (GoRaw "nil") (Array.index accArgs.exprs 2)
                          arrExprType = fromMaybe TypeValue (Array.index accArgs.exprTypes 2)

                          mbFnVar = case Array.index args 0 of
                            Just fArg -> getVar (unwrapTcoExpr fArg)
                            Nothing -> Nothing

                          fnFullName = case mbFnVar of
                            Just { mbMod: Just (ModuleName mn), name } -> String.replaceAll (Pattern ".") (Replacement "_") mn <> "." <> name
                            Just { mbMod: Nothing, name } -> modNameStr <> "." <> name
                            Nothing -> ""

                          mbFnArityInfo = Map.lookup fnFullName moduleFunctions

                          elemType = case arrExprType of
                            TypeNativeArray inner -> inner
                            _ -> TypeValue

                          loopBody = case mbFnArityInfo of
                            Just info | info.arity == 2 ->
                              let
                                expectedArg0 = fromMaybe TypeValue (Array.index info.fArgs 0)
                                expectedArg1 = fromMaybe TypeValue (Array.index info.fArgs 1)
                              in
                                GoMutate resGoName (unboxGoExpr codegenStateRef modNameStr (GoCall (GoVar ("Call_" <> String.replaceAll (Pattern ".") (Replacement "_") fnFullName)) [ unboxGoExpr codegenStateRef modNameStr (GoVar resGoName) initExprType expectedArg0, unboxGoExpr codegenStateRef modNameStr (GoVar vName) elemType expectedArg1 ]) info.fRet initExprType)
                            _ ->
                              GoMutate resGoName (unboxGoExpr codegenStateRef modNameStr (GoCall (GoSelector (GoVar "gopurs_runtime") "Apply2") [ boxGoExpr codegenStateRef modNameStr fExpr fExprType, boxGoExpr codegenStateRef modNameStr (GoVar resGoName) initExprType, boxGoExpr codegenStateRef modNameStr (GoVar vName) elemType ]) TypeValue initExprType)

                          arrGoAssignment = case arrExprType of
                            TypeNativeArray _ -> GoAssign arrGoName (GoVar arrValName)
                            _ -> GoAssign arrGoName (GoCall (GoRaw "(*[]gopurs_runtime.Value)") [ GoSelector (GoVar arrValName) "UnsafePtr" ])

                          arrGoRangeTarget = case arrExprType of
                            TypeNativeArray _ -> arrGoName
                            _ -> "*" <> arrGoName

                          iifeBody = GoBlock
                            [ GoAssign resGoName initExpr
                            , arrGoAssignment
                            , GoForRange ("_, " <> vName <> " := range " <> arrGoRangeTarget) [ loopBody ]
                            , GoReturn (GoVar resGoName)
                            ]
                        in
                          { stmts: accArgs.stmts, expr: GoIIFE arrValName arrExpr iifeBody, exprType: initExprType, nextId: accArgs.nextId }

                      "filterImpl" ->
                        let
                          fExprRaw = fromMaybe (GoRaw "nil") (Array.index accArgs.exprs 0)
                          fExprType = fromMaybe TypeValue (Array.index accArgs.exprTypes 0)
                          arrExprRaw = fromMaybe (GoRaw "nil") (Array.index accArgs.exprs 1)
                          arrExprType = fromMaybe TypeValue (Array.index accArgs.exprTypes 1)

                          mbFnVar = case Array.index args 0 of
                            Just fArg -> getVar (unwrapTcoExpr fArg)
                            Nothing -> Nothing

                          fnFullName = case mbFnVar of
                            Just { mbMod: Just (ModuleName mn), name } -> String.replaceAll (Pattern ".") (Replacement "_") mn <> "." <> name
                            Just { mbMod: Nothing, name } -> modNameStr <> "." <> name
                            Nothing -> ""

                          mbFnArityInfo = Map.lookup fnFullName moduleFunctions

                          elemType = case arrExprType of
                            TypeNativeArray inner -> inner
                            _ -> TypeValue

                          isTrueExpr = case mbFnArityInfo of
                            Just info | info.arity == 1 ->
                              let
                                expectedArgType = fromMaybe TypeValue (Array.index info.fArgs 0)
                              in
                                GoCall (GoVar ("Call_" <> String.replaceAll (Pattern ".") (Replacement "_") fnFullName)) [ unboxGoExpr codegenStateRef modNameStr (GoVar vName) elemType expectedArgType ]
                            _ ->
                              let
                                condExpr = GoCall (GoSelector (GoVar "gopurs_runtime") "Apply") [ boxGoExpr codegenStateRef modNameStr fExprRaw fExprType, boxGoExpr codegenStateRef modNameStr (GoVar vName) elemType ]
                              in
                                GoCall (GoSelector condExpr "BoolVal") []

                          loopBody = GoIfElse isTrueExpr [ GoMutate resGoName (GoCall (GoVar "append") [ GoVar resGoName, GoVar vName ]) ] []

                          arrGoAssignment = case arrExprType of
                            TypeNativeArray _ -> GoAssign arrGoName (GoVar arrValName)
                            _ -> GoAssign arrGoName (GoCall (GoRaw "(*[]gopurs_runtime.Value)") [ GoSelector (GoVar arrValName) "UnsafePtr" ])

                          arrGoRangeTarget = case arrExprType of
                            TypeNativeArray _ -> arrGoName
                            _ -> "*" <> arrGoName

                          iifeBodyStmts =
                            [ arrGoAssignment
                            , GoAssign resGoName (GoCall (GoVar "make") [ GoRaw ("[]" <> goTypeToStr elemType), GoRaw "0" ])
                            , GoForRange ("_, " <> vName <> " := range " <> arrGoRangeTarget) [ loopBody ]
                            ]
                          filterExpr = GoCall (GoFuncLit [] (Array.cons (GoAssign arrValName arrExprRaw) (Array.cons (GoMutate "_" (GoVar arrValName)) iifeBodyStmts)) (GoVar resGoName) (TypeNativeArray elemType)) []
                          freshFilterExpr = if elemType == TypeValue && Array.length args == 2 then GoFreshFilterArray filterExpr else filterExpr
                        in
                          { stmts: accArgs.stmts, expr: freshFilterExpr, exprType: TypeNativeArray elemType, nextId: accArgs.nextId }

                      _ -> { stmts: accArgs.stmts, expr: GoRaw "nil", exprType: TypeValue, nextId: accArgs.nextId }
                  in
                    iifeExpr
                Nothing ->
                  let
                    Tuple flatFn flatArgsSpine = collectGoSpine tcoExpr
                    flatArgs = getGoSpineArgs flatArgsSpine
                    isTailCallTo =
                      if isTail then case unwrapTcoExpr flatFn of
                        Local mbIdent lvl ->
                          let
                            v = fromMaybe { name: localId mbIdent lvl, goType: TypeValue } (Map.lookup (localId mbIdent lvl) bound)
                          in
                            Array.findIndex (\ctx -> ctx.ident == v.name) loopCtx
                        Var (Qualified mbMod (Ident name)) ->
                          let
                            fullName = sanitizeName name
                          in
                            Array.findIndex (\ctx -> ctx.ident == fullName) loopCtx
                        _ -> Nothing
                      else Nothing
                  in
                    case isTailCallTo of
                      Just index ->
                        let
                          accFinal = foldl
                            ( \acc arg ->
                                let
                                  argRes = translateExpr codegenStateRef (depth + 1) modNameStr recVars moduleFunctions bound Nothing [] { isTail: false, inEffectBlock: false } acc.nextId arg
                                in
                                  { stmts: acc.stmts <> argRes.stmts, exprs: Array.snoc acc.exprs argRes.expr, exprTypes: Array.snoc acc.exprTypes argRes.exprType, nextId: argRes.nextId }
                            )
                            { stmts: StmtEmpty, exprs: [], exprTypes: [], nextId }
                            flatArgs
                          targetCtx = fromMaybe { ident: "", params: [], loopParams: [], goTypes: [], fRet: TypeValue } (Array.index loopCtx index)
                          assigns = Array.mapWithIndex
                            ( \i paramName ->
                                let
                                  argExpr = fromMaybe (GoRaw "nil") (Array.index accFinal.exprs i)
                                  argType = fromMaybe TypeValue (Array.index accFinal.exprTypes i)
                                  expectedType = fromMaybe TypeValue (Array.index targetCtx.goTypes i)
                                in
                                  GoMutate paramName (coerceGoExpr codegenStateRef modNameStr argExpr argType expectedType)
                            )
                            targetCtx.loopParams
                        in
                          let
                            expectedGoType = exprTypeToGoType (unsafePerformEffect (Ref.read codegenStateRef)).pointerAdtPaths (unsafePerformEffect (Ref.read codegenStateRef)).enumAdts (unsafePerformEffect (Ref.read codegenStateRef)).elidedCtors modNameStr (case getExprType tcoExpr of
                                Any -> fromMaybe Any mbExpectedExprType
                                ty -> ty)
                            expectedGoTypeStr = goTypeToStr expectedGoType
                          in
                          { stmts: accFinal.stmts <> foldMap StmtLeaf assigns <> StmtLeaf (GoContinue targetCtx.ident), expr: GoRaw ("func() " <> expectedGoTypeStr <> " { panic(\"unreachable\") }()"), exprType: expectedGoType, nextId: accFinal.nextId }
                      Nothing ->
                        let
                          resFn = translateExpr codegenStateRef (depth + 1) modNameStr recVars moduleFunctions bound Nothing [] { isTail: false, inEffectBlock: false } nextId fn
                          accArgs = foldl
                            ( \acc arg ->
                                let
                                  argRes = translateExpr codegenStateRef (depth + 1) modNameStr recVars moduleFunctions bound Nothing [] { isTail: false, inEffectBlock: false } acc.nextId arg
                                in
                                  { stmts: acc.stmts <> argRes.stmts, exprs: Array.snoc acc.exprs argRes.expr, exprTypes: Array.snoc acc.exprTypes argRes.exprType, nextId: argRes.nextId }
                            )
                            { stmts: resFn.stmts, exprs: [], exprTypes: [], nextId: resFn.nextId }
                            args
                          len = Array.length args
                          goFuncName = if len >= 2 && len <= 10 then "UncurriedApp" <> show len else "UncurriedApp"
                        in
                          case resFn.exprType of
                            TypeFunc fArgs fRet | Array.length fArgs == len ->
                              let
                                callArgs = Array.mapWithIndex
                                  ( \i expected ->
                                      let
                                        arg = fromMaybe (GoRaw "nil") (Array.index accArgs.exprs i)
                                        actual = fromMaybe TypeValue (Array.index accArgs.exprTypes i)
                                      in
                                        coerceGoExpr codegenStateRef modNameStr arg actual expected
                                  )
                                  fArgs
                              in
                                { stmts: accArgs.stmts, expr: boxGoExpr codegenStateRef modNameStr (GoCall resFn.expr callArgs) fRet, exprType: TypeValue, nextId: accArgs.nextId }
                            _ ->
                              let
                                boxedArgs = Array.zipWith (\arg actual -> boxGoExpr codegenStateRef modNameStr arg actual) accArgs.exprs accArgs.exprTypes
                              in
                                { stmts: accArgs.stmts, expr: GoCall (GoSelector (GoVar "gopurs_runtime") goFuncName) (Array.cons (boxGoExpr codegenStateRef modNameStr resFn.expr resFn.exprType) boxedArgs), exprType: TypeValue, nextId: accArgs.nextId }

          UncurriedAbs args body -> liftIfNeeded \_ ->
            let
              mbFuncTy = case extractExprFuncType (getExprType tcoExpr) of
                Just r -> Just r
                Nothing -> case mbExpectedExprType of
                  Just ty -> extractExprFuncType ty
                  Nothing -> Nothing

              paramsWithTypes = map (\(Tuple mbI lvl) -> Tuple (localId mbI lvl) TypeValue) args

              newBound = foldl (\acc (Tuple idStr goType) -> Map.insert idStr { name: idStr, goType } acc) bound paramsWithTypes

              goParams = String.joinWith ", " (map (\(Tuple p goT) -> p <> " " <> goTypeToStr goT) paramsWithTypes)
              resBody = translateExprWithExpectedType codegenStateRef (depth + 1) modNameStr recVars moduleFunctions newBound Nothing [] { isTail, inEffectBlock: false } (case mbFuncTy of
                Just { fRet } -> Just fRet
                Nothing -> Nothing) nextId body
              arity = Array.length args
            in
              if arity >= 2 && arity <= 10 then
                case tcoIdent of
                  Just topName ->
                    let
                      callFuncDecl = "func Call_" <> modNameStr <> "_" <> topName <> "(" <> goParams <> ") gopurs_runtime.Value {\n" <> printGoExpr (GoBlock (flattenStmts resBody.stmts <> [ GoReturn (boxGoExpr codegenStateRef modNameStr resBody.expr resBody.exprType) ])) <> "\n}"
                      funcExpr = unsafePerformEffect do
                        Ref.modify_ (\r -> r { rawDecls = Array.snoc r.rawDecls callFuncDecl }) codegenStateRef
                        pure $ GoRaw ("gopurs_runtime.Func" <> show arity <> "(Call_" <> modNameStr <> "_" <> topName <> ")")
                    in
                      { stmts: StmtEmpty, expr: funcExpr, exprType: TypeValue, nextId: resBody.nextId }
                  Nothing ->
                    let
                      funcExpr = GoRaw ("gopurs_runtime.Func" <> show arity <> "(func(" <> goParams <> ") gopurs_runtime.Value {\n" <> printGoExpr (GoBlock (flattenStmts resBody.stmts <> [ GoReturn (boxGoExpr codegenStateRef modNameStr resBody.expr resBody.exprType) ])) <> "\n})")
                    in
                      { stmts: StmtEmpty, expr: funcExpr, exprType: TypeValue, nextId: resBody.nextId }
              else
                let
                  params = map fst paramsWithTypes
                  makeCurried [] = resBody.expr
                  makeCurried [ p ] = GoFunc p TypeValue TypeValue (GoBlock (flattenStmts resBody.stmts <> [ GoReturn (boxGoExpr codegenStateRef modNameStr resBody.expr resBody.exprType) ]))
                  makeCurried ps = case Array.uncons ps of
                    Just { head: p, tail: rest } -> GoFunc p TypeValue TypeValue (makeCurried rest)
                    Nothing -> resBody.expr
                in
                  { stmts: StmtEmpty, expr: makeCurried params, exprType: TypeValue, nextId: resBody.nextId }

          UncurriedEffectApp fn args ->
            let
              resFn = translateExpr codegenStateRef (depth + 1) modNameStr recVars moduleFunctions bound Nothing [] { isTail: false, inEffectBlock: false } nextId fn
              accArgs = foldl
                ( \acc arg ->
                    let
                      argRes = translateExpr codegenStateRef (depth + 1) modNameStr recVars moduleFunctions bound Nothing [] { isTail: false, inEffectBlock: false } acc.nextId arg
                    in
                      { stmts: acc.stmts <> argRes.stmts, exprs: Array.snoc acc.exprs argRes.expr, exprTypes: Array.snoc acc.exprTypes argRes.exprType, nextId: argRes.nextId }
                )
                { stmts: resFn.stmts, exprs: [], exprTypes: [], nextId: resFn.nextId }
                args
            in
              let
                len = Array.length args
                goFuncName = if len >= 2 && len <= 5 then "UncurriedApp" <> show len else "UncurriedApp"
              in
                case resFn.exprType of
                  TypeFunc fArgs fRet | Array.length fArgs == len ->
                    let
                      callArgs = Array.mapWithIndex
                        ( \i expected ->
                            let
                              arg = fromMaybe (GoRaw "nil") (Array.index accArgs.exprs i)
                              actual = fromMaybe TypeValue (Array.index accArgs.exprTypes i)
                            in
                              coerceGoExpr codegenStateRef modNameStr arg actual expected
                        )
                        fArgs
                    in
                      { stmts: accArgs.stmts, expr: boxGoExpr codegenStateRef modNameStr (GoCall resFn.expr callArgs) fRet, exprType: TypeValue, nextId: accArgs.nextId }
                  _ ->
                    let
                      boxedArgs = Array.zipWith (\arg actual -> boxGoExpr codegenStateRef modNameStr arg actual) accArgs.exprs accArgs.exprTypes
                    in
                      { stmts: accArgs.stmts, expr: GoCall (GoSelector (GoVar "gopurs_runtime") goFuncName) (Array.cons (boxGoExpr codegenStateRef modNameStr resFn.expr resFn.exprType) boxedArgs), exprType: TypeValue, nextId: accArgs.nextId }

          UncurriedEffectAbs args body -> liftIfNeeded \_ ->
            let
              mbFuncTy = case extractExprFuncType (getExprType tcoExpr) of
                Just r -> Just r
                Nothing -> case mbExpectedExprType of
                  Just ty -> extractExprFuncType ty
                  Nothing -> Nothing

              paramsWithTypes = map (\(Tuple mbI lvl) -> Tuple (localId mbI lvl) TypeValue) args
              newBound = foldl (\acc (Tuple idStr goType) -> Map.insert idStr { name: idStr, goType } acc) bound paramsWithTypes
              goParams = String.joinWith ", " (map (\(Tuple p goT) -> p <> " " <> goTypeToStr goT) paramsWithTypes)
              resBody = translateExprWithExpectedType codegenStateRef (depth + 1) modNameStr recVars moduleFunctions newBound Nothing [] { isTail, inEffectBlock: false } (case mbFuncTy of
                Just { fRet } -> Just fRet
                Nothing -> Nothing) nextId body
              arity = Array.length args
            in
              if arity >= 2 && arity <= 5 then
                let
                  funcExpr = GoRaw ("gopurs_runtime.Func" <> show arity <> "(func(" <> goParams <> ") gopurs_runtime.Value {\n" <> printGoExpr (GoBlock (flattenStmts resBody.stmts <> [ GoReturn (GoRaw ("gopurs_runtime.Apply(" <> printGoExpr (boxGoExpr codegenStateRef modNameStr resBody.expr resBody.exprType) <> ", gopurs_runtime.Value{})")) ])) <> "\n})")
                in
                  { stmts: StmtEmpty, expr: funcExpr, exprType: TypeValue, nextId: resBody.nextId }
              else
                let
                  params = map fst paramsWithTypes
                  makeCurried [] = GoRaw ("gopurs_runtime.Apply(" <> printGoExpr (boxGoExpr codegenStateRef modNameStr resBody.expr resBody.exprType) <> ", gopurs_runtime.Value{})")
                  makeCurried [ p ] = GoFunc p TypeValue TypeValue (GoBlock (flattenStmts resBody.stmts <> [ GoReturn (GoRaw ("gopurs_runtime.Apply(" <> printGoExpr (boxGoExpr codegenStateRef modNameStr resBody.expr resBody.exprType) <> ", gopurs_runtime.Value{})")) ]))
                  makeCurried ps = case Array.uncons ps of
                    Just { head: p, tail: rest } -> GoFunc p TypeValue TypeValue (makeCurried rest)
                    Nothing -> resBody.expr
                in
                  { stmts: StmtEmpty, expr: makeCurried params, exprType: TypeValue, nextId: resBody.nextId }

          EffectBind mbIdent lvl binding body ->
            let
              stripEffectDefer (TcoExpr a syn) = case unwrapTcoExpr (TcoExpr a syn) of
                EffectDefer inner -> let Tuple _ i = stripEffectDefer inner in Tuple true i
                Abs _ inner -> let Tuple _ i = stripEffectDefer inner in Tuple true i
                Let ident lvl_ val body_ -> let Tuple stripped i = stripEffectDefer body_ in Tuple stripped (TcoExpr a (Let ident lvl_ val i))
                LetRec lvl_ bindings_ body_ -> let Tuple stripped i = stripEffectDefer body_ in Tuple stripped (TcoExpr a (LetRec lvl_ bindings_ i))
                _ -> Tuple false (TcoExpr a syn)
              Tuple wasStripped realBinding = stripEffectDefer binding
              originalName = localId mbIdent lvl
              name = originalName <> "_" <> show nextId
              newBound = Map.insert originalName { name, goType: TypeValue } bound
              resBinding = translateExpr codegenStateRef (depth + 1) modNameStr recVars moduleFunctions bound Nothing [] { isTail: false, inEffectBlock: true } (nextId + 1) realBinding
              resBody = translateExprWithExpectedType codegenStateRef (depth + 1) modNameStr recVars moduleFunctions newBound Nothing loopCtx { isTail, inEffectBlock: true } mbExpectedExprType resBinding.nextId body
              bindingExpr = if wasStripped then boxGoExpr codegenStateRef modNameStr resBinding.expr resBinding.exprType else executeIfOpaque realBinding (boxGoExpr codegenStateRef modNameStr resBinding.expr resBinding.exprType)
              bodyExpr = executeIfOpaque body resBody.expr
            in
              { stmts: resBinding.stmts <> StmtLeaf (GoAssign name bindingExpr) <> resBody.stmts, expr: bodyExpr, exprType: resBody.exprType, nextId: resBody.nextId }

          EffectPure binding ->
            translateExpr codegenStateRef (depth + 1) modNameStr recVars moduleFunctions bound Nothing [] { isTail: false, inEffectBlock: false } nextId binding

          EffectDefer binding ->
            let
              resBinding = translateExpr codegenStateRef (depth + 1) modNameStr recVars moduleFunctions bound Nothing [] { isTail: false, inEffectBlock: true } nextId binding
              funcExpr = GoRaw ("gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {\n" <> printGoExpr (GoBlock (flattenStmts resBinding.stmts <> [ GoReturn (boxGoExpr codegenStateRef modNameStr resBinding.expr resBinding.exprType) ])) <> "\n})")
            in
              { stmts: StmtEmpty, expr: funcExpr, exprType: TypeValue, nextId: resBinding.nextId }

          Let mbIdent lvl binding body ->
            let
              originalName = localId mbIdent lvl
              name = originalName <> "_" <> show nextId
              expectedGoTypeFromAst = exprTypeToGoType (unsafePerformEffect (Ref.read codegenStateRef)).pointerAdtPaths (unsafePerformEffect (Ref.read codegenStateRef)).enumAdts (unsafePerformEffect (Ref.read codegenStateRef)).elidedCtors modNameStr (getExprType binding)

              mbFunc = extractUncurriedAbs binding
            in
              case mbFunc of
                Just abs | Array.length abs.args > 0 ->
                  let
                    fArgsAst = case extractExprFuncType (getExprType binding) of
                      Just { fArgs } -> fArgs
                      Nothing -> []
                    paramsWithTypes = Array.zipWith
                      ( \idStr ty ->
                          Tuple idStr (exprTypeToGoType (unsafePerformEffect (Ref.read codegenStateRef)).pointerAdtPaths (unsafePerformEffect (Ref.read codegenStateRef)).enumAdts (unsafePerformEffect (Ref.read codegenStateRef)).elidedCtors modNameStr ty)
                      )
                      abs.args
                      (fArgsAst <> Array.replicate (max 0 (Array.length abs.args - Array.length fArgsAst)) Any)
                    goTypes = map snd paramsWithTypes

                    localModuleFunctions = Map.insert name { fullName: "Call_local_" <> modNameStr <> "_" <> name, fArgs: goTypes, fRet: TypeValue, arity: Array.length abs.args } moduleFunctions
                    declStmts = [ GoRaw ("var Call_local_" <> modNameStr <> "_" <> name <> " func(" <> String.joinWith ", " (map goTypeToStr goTypes) <> ") gopurs_runtime.Value"), GoRaw ("_ = Call_local_" <> modNameStr <> "_" <> name), GoRaw ("var " <> name <> " gopurs_runtime.Value"), GoRaw ("_ = " <> name) ]

                    loopBound = foldl (\acc (Tuple idStr goT) -> Map.insert idStr { name: idStr, goType: goT } acc) bound paramsWithTypes
                    resBodyMut = translateExpr codegenStateRef (depth + 1) modNameStr recVars localModuleFunctions loopBound (Just name) [] { isTail: true, inEffectBlock: false } (nextId + 1) abs.body

                    goParamsNative = String.joinWith ", " (map (\(Tuple p goT) -> p <> "_loop " <> goTypeToStr goT) paramsWithTypes)
                    initVars = Array.concatMap (\(Tuple p goT) -> [ GoRaw ("var " <> p <> " " <> goTypeToStr goT <> " = " <> p <> "_loop"), GoRaw ("_ = " <> p) ]) paramsWithTypes
                    funcBody = GoBlock (initVars <> flattenStmts resBodyMut.stmts <> [ GoReturn (boxGoExpr codegenStateRef modNameStr resBodyMut.expr resBodyMut.exprType) ])
                    nativeAssignment = GoMutate ("Call_local_" <> modNameStr <> "_" <> name) (GoRaw ("func(" <> goParamsNative <> ") gopurs_runtime.Value {\n" <> printGoExpr funcBody <> "\n}"))

                    nativeCallExpr = GoCall (GoVar ("Call_local_" <> modNameStr <> "_" <> name)) (map (\(Tuple p goT) -> unboxGoExpr codegenStateRef modNameStr (GoVar (p <> "_loop_val")) TypeValue goT) paramsWithTypes)
                    funcExpr = Array.foldr (\(Tuple p goT) acc -> GoCall (GoSelector (GoVar "gopurs_runtime") "Func") [ GoRaw ("func(" <> p <> "_loop_val gopurs_runtime.Value) gopurs_runtime.Value {\nreturn " <> printGoExpr acc <> "\n}") ]) nativeCallExpr paramsWithTypes

                    newBound = Map.insert originalName { name, goType: TypeValue } bound
                    resBodyOuter = translateExpr codegenStateRef (depth + 1) modNameStr recVars localModuleFunctions newBound Nothing loopCtx options resBodyMut.nextId body
                  in
                    { stmts: foldMap StmtLeaf declStmts <> StmtLeaf nativeAssignment <> StmtLeaf (GoMutate name funcExpr) <> resBodyOuter.stmts, expr: resBodyOuter.expr, exprType: resBodyOuter.exprType, nextId: resBodyOuter.nextId }

                _ ->
                  let
                    resBinding = translateExpr codegenStateRef (depth + 1) modNameStr recVars moduleFunctions bound Nothing [] { isTail: false, inEffectBlock: false } (nextId + 1) binding
                    actualGoType = if expectedGoTypeFromAst == TypeValue then resBinding.exprType else expectedGoTypeFromAst
                    newBound = Map.insert originalName { name, goType: actualGoType } bound
                    resBody = translateExprWithExpectedType codegenStateRef (depth + 1) modNameStr recVars moduleFunctions newBound Nothing loopCtx options mbExpectedExprType resBinding.nextId body
                    letStmt =
                      if actualGoType == resBinding.exprType then
                        StmtLeaf (GoAssign name resBinding.expr)
                      else
                        StmtLeaf (GoRaw ("var " <> name <> " " <> goTypeToStr actualGoType <> " = " <> printGoExpr (unboxGoExpr codegenStateRef modNameStr resBinding.expr resBinding.exprType actualGoType)))
                  in
                    { stmts: resBinding.stmts <> StmtLeaf (GoRaw ("// TAST (Let): " <> name <> " shape=" <> printTcoExprShape binding <> " bindingType=" <> printExprType (getExprType binding))) <> letStmt <> resBody.stmts, expr: resBody.expr, exprType: resBody.exprType, nextId: resBody.nextId }

          LetRec lvl bindings body ->
            let
              allocRes = foldl
                ( \acc (Tuple (Ident ident) val) ->
                    let
                      oldName = localId (Just (Ident ident)) lvl
                      gId = unsafePerformEffect do
                        curr <- Ref.read codegenStateRef
                        Ref.modify_ (\r -> r { globalId = r.globalId + 1 }) codegenStateRef
                        pure curr.globalId
                      newName = oldName <> "_" <> show acc.nextId <> "_" <> show gId
                      expectedGoTypeFromAst = exprTypeToGoType (unsafePerformEffect (Ref.read codegenStateRef)).pointerAdtPaths (unsafePerformEffect (Ref.read codegenStateRef)).enumAdts (unsafePerformEffect (Ref.read codegenStateRef)).elidedCtors modNameStr (getExprType val)
                    in
                      { newBound: Map.insert oldName { name: newName, goType: expectedGoTypeFromAst } acc.newBound, newNames: Array.snoc acc.newNames { oldName, newName }, exprType: TypeValue, nextId: acc.nextId + 1 }
                )
                { newBound: bound, newNames: [], exprType: TypeValue, nextId }
                (toArray bindings)

              combinedRecVars = recVars <> map (\(Tuple (Ident i) _) -> sanitizeName i) (toArray bindings)

              isLoop = (unwrap tcoAnalysis).role.isLoop
              mutRecBinds =
                if isLoop then
                  traverse (\(Tuple (Ident name) val) -> map (\abs -> { ident: sanitizeName name, args: abs.args, body: abs.body, fvs: memoizedFreeVars val, val }) (extractUncurriedAbs val)) (toArray bindings)
                else Nothing
            in
              case mutRecBinds of
                Just fns ->
                  let
                    loopCtxs = map
                      ( \fn ->
                          let
                            oldName = localId (Just (Ident fn.ident)) lvl
                            boundInfo = fromMaybe { name: oldName, goType: TypeValue } (Map.lookup oldName allocRes.newBound)
                            newName = boundInfo.name
                            goType = boundInfo.goType
                            fRet = case goType of
                              TypeFunc _ r -> r
                              _ -> TypeValue
                            fArgs = case extractExprFuncType (getExprType fn.val) of
                              Just { fArgs: a } -> map (exprTypeToGoType (unsafePerformEffect (Ref.read codegenStateRef)).pointerAdtPaths (unsafePerformEffect (Ref.read codegenStateRef)).enumAdts (unsafePerformEffect (Ref.read codegenStateRef)).elidedCtors modNameStr) a
                              Nothing -> []
                            paramsWithTypes = Array.zipWith (\idStr goT -> Tuple idStr goT) fn.args (fArgs <> Array.replicate (max 0 (Array.length fn.args - Array.length fArgs)) TypeValue)
                          in
                            { ident: newName, params: fn.args, loopParams: map (\p -> p <> "_loop") fn.args, goTypes: map snd paramsWithTypes, fRet }
                      )
                      fns

                    combinedLoopCtx = loopCtxs <> loopCtx
                    
                    prepopulatedFunctions = foldl (\accCtx fn ->
                        let
                          oldName = localId (Just (Ident fn.ident)) lvl
                          boundInfo = fromMaybe { name: oldName, goType: TypeValue } (Map.lookup oldName allocRes.newBound)
                          newName = boundInfo.name
                          fArgs = case extractExprFuncType (getExprType fn.val) of
                            Just { fArgs: a } -> map (exprTypeToGoType (unsafePerformEffect (Ref.read codegenStateRef)).pointerAdtPaths (unsafePerformEffect (Ref.read codegenStateRef)).enumAdts (unsafePerformEffect (Ref.read codegenStateRef)).elidedCtors modNameStr) a
                            Nothing -> []
                          fRet = case extractExprFuncType (getExprType fn.val) of
                            Just { fRet: r } -> exprTypeToGoType (unsafePerformEffect (Ref.read codegenStateRef)).pointerAdtPaths (unsafePerformEffect (Ref.read codegenStateRef)).enumAdts (unsafePerformEffect (Ref.read codegenStateRef)).elidedCtors modNameStr r
                            Nothing -> TypeValue
                          paramsWithTypes = Array.zipWith (\idStr goT -> Tuple idStr goT) fn.args (fArgs <> Array.replicate (max 0 (Array.length fn.args - Array.length fArgs)) TypeValue)
                        in
                          Map.insert newName { fullName: "Call_local_" <> modNameStr <> "_" <> newName, fArgs: map snd paramsWithTypes, fRet: fRet, arity: Array.length fn.args } accCtx
                      ) moduleFunctions fns

                    prepopulatedBound = foldl (\accCtx fn -> 
                        let
                          oldName = localId (Just (Ident fn.ident)) lvl
                          boundInfo = fromMaybe { name: oldName, goType: TypeValue } (Map.lookup oldName allocRes.newBound)
                          newName = boundInfo.name
                          fArgs = case extractExprFuncType (getExprType fn.val) of
                            Just { fArgs: a } -> map (exprTypeToGoType (unsafePerformEffect (Ref.read codegenStateRef)).pointerAdtPaths (unsafePerformEffect (Ref.read codegenStateRef)).enumAdts (unsafePerformEffect (Ref.read codegenStateRef)).elidedCtors modNameStr) a
                            Nothing -> []
                          fRet = case extractExprFuncType (getExprType fn.val) of
                            Just { fRet: r } -> exprTypeToGoType (unsafePerformEffect (Ref.read codegenStateRef)).pointerAdtPaths (unsafePerformEffect (Ref.read codegenStateRef)).enumAdts (unsafePerformEffect (Ref.read codegenStateRef)).elidedCtors modNameStr r
                            Nothing -> TypeValue
                          paramsWithTypes = Array.zipWith (\idStr goT -> Tuple idStr goT) fn.args (fArgs <> Array.replicate (max 0 (Array.length fn.args - Array.length fArgs)) TypeValue)
                        in
                          Map.insert oldName { name: newName, goType: TypeFunc (map snd paramsWithTypes) fRet } accCtx
                      ) allocRes.newBound fns

                    resData = foldl
                      ( \acc fn ->
                          let
                            oldName = localId (Just (Ident fn.ident)) lvl
                            boundInfo = fromMaybe { name: oldName, goType: TypeValue } (Map.lookup oldName prepopulatedBound)
                            newName = boundInfo.name
                            fArgs = case extractExprFuncType (getExprType fn.val) of
                              Just { fArgs: a } -> map (exprTypeToGoType (unsafePerformEffect (Ref.read codegenStateRef)).pointerAdtPaths (unsafePerformEffect (Ref.read codegenStateRef)).enumAdts (unsafePerformEffect (Ref.read codegenStateRef)).elidedCtors modNameStr) a
                              Nothing -> []
                            paramsWithTypes = Array.zipWith (\idStr goT -> Tuple idStr goT) fn.args (fArgs <> Array.replicate (max 0 (Array.length fn.args - Array.length fArgs)) TypeValue)
                            currentLoopCtx = [ { ident: newName, params: fn.args, loopParams: map (\p -> p <> "_loop") fn.args, goTypes: map snd paramsWithTypes, fRet: TypeValue } ]
                            loopBound = foldl (\acc2 (Tuple idStr goT) -> Map.insert idStr { name: idStr, goType: goT } acc2) prepopulatedBound paramsWithTypes
                            mbExpectedRet = case extractExprFuncType (getExprType fn.val) of
                              Just { fRet: r } -> Just r
                              Nothing -> Nothing
                            resBodyMut = translateExprWithExpectedType codegenStateRef (depth + 1) modNameStr combinedRecVars acc.moduleFunctions loopBound (Just newName) currentLoopCtx { isTail: true, inEffectBlock: false } mbExpectedRet acc.nextId fn.body
                            trueFRet = resBodyMut.exprType

                            initVars = Array.concatMap (\(Tuple p goT) -> [ GoRaw ("var " <> p <> " " <> goTypeToStr goT <> " = " <> p <> "_loop"), GoRaw ("_ = " <> p) ]) paramsWithTypes

                            funcBody = GoFor newName (initVars <> flattenStmts resBodyMut.stmts <> [ GoReturn resBodyMut.expr ])

                            goParamsNative = String.joinWith ", " (map (\(Tuple p goT) -> p <> "_loop " <> goTypeToStr goT) paramsWithTypes)
                            nativeAssignment = GoMutate ("Call_local_" <> modNameStr <> "_" <> newName) (GoRaw ("func(" <> goParamsNative <> ") " <> goTypeToStr trueFRet <> " {\n" <> printGoExpr funcBody <> "\n}"))

                            nativeCallExpr = GoCall (GoVar ("Call_local_" <> modNameStr <> "_" <> newName)) (map (\(Tuple p goT) -> unboxGoExpr codegenStateRef modNameStr (GoVar (p <> "_loop_val")) TypeValue goT) paramsWithTypes)
                            funcExpr = Array.foldr (\(Tuple p goT) accExpr -> GoCall (GoSelector (GoVar "gopurs_runtime") "Func") [ GoRaw ("func(" <> p <> "_loop_val gopurs_runtime.Value) gopurs_runtime.Value {\nreturn " <> printGoExpr accExpr <> "\n}") ]) (boxGoExpr codegenStateRef modNameStr nativeCallExpr trueFRet) paramsWithTypes
                            
                            newFunctions = Map.insert newName { fullName: "Call_local_" <> modNameStr <> "_" <> newName, fArgs: map snd paramsWithTypes, fRet: trueFRet, arity: Array.length fn.args } acc.moduleFunctions
                            newBound2 = Map.insert oldName { name: newName, goType: TypeFunc (map snd paramsWithTypes) trueFRet } acc.newBound
                            declStmtsLocal = [ GoRaw ("var Call_local_" <> modNameStr <> "_" <> newName <> " func(" <> String.joinWith ", " (map goTypeToStr (map snd paramsWithTypes)) <> ") " <> goTypeToStr trueFRet), GoRaw ("_ = Call_local_" <> modNameStr <> "_" <> newName), GoRaw ("var " <> newName <> " gopurs_runtime.Value"), GoRaw ("_ = " <> newName) ]
                          in
                            { stmts: acc.stmts <> declStmtsLocal <> [ nativeAssignment, GoMutate newName funcExpr ], nextId: resBodyMut.nextId, moduleFunctions: newFunctions, newBound: newBound2 }
                      )
                      { stmts: [], nextId: allocRes.nextId, moduleFunctions: prepopulatedFunctions, newBound: prepopulatedBound }
                      fns

                    resBodyOuter = translateExpr codegenStateRef (depth + 1) modNameStr combinedRecVars resData.moduleFunctions resData.newBound Nothing loopCtx options resData.nextId body
                  in
                    { stmts: foldMap StmtLeaf resData.stmts <> resBodyOuter.stmts, expr: resBodyOuter.expr, exprType: resBodyOuter.exprType, nextId: resBodyOuter.nextId }

                Nothing ->
                  let
                    accBindings = foldl
                      ( \acc (Tuple (Tuple (Ident ident) val) alloc) ->
                          let
                            res = translateExpr codegenStateRef (depth + 1) modNameStr combinedRecVars moduleFunctions allocRes.newBound Nothing [] { isTail: false, inEffectBlock: false } acc.nextId val
                            expectedGoType = (fromMaybe { name: alloc.newName, goType: TypeValue } (Map.lookup alloc.oldName allocRes.newBound)).goType
                            assignedVal = if expectedGoType == res.exprType then res.expr else unboxGoExpr codegenStateRef modNameStr res.expr res.exprType expectedGoType
                          in
                            { stmts: acc.stmts <> res.stmts, exprs: Array.snoc acc.exprs { key: alloc.newName, value: assignedVal, goType: expectedGoType }, exprType: TypeValue, nextId: res.nextId }
                      )
                      { stmts: StmtEmpty, exprs: [], exprType: TypeValue, nextId: allocRes.nextId }
                      (Array.zip (toArray bindings) allocRes.newNames)

                    declStmts = map (\b -> GoRaw ("var " <> b.key <> " " <> goTypeToStr b.goType <> "\n_ = " <> b.key <> "\n// FALLBACK TCO: isLoop=" <> show isLoop <> " len=" <> show (Array.length (toArray bindings)))) accBindings.exprs
                    assignStmts = map (\b -> GoMutate b.key b.value) accBindings.exprs

                    resBody = translateExprWithExpectedType codegenStateRef (depth + 1) modNameStr combinedRecVars moduleFunctions allocRes.newBound Nothing loopCtx options mbExpectedExprType accBindings.nextId body
                  in
                    { stmts: foldMap StmtLeaf declStmts <> accBindings.stmts <> foldMap StmtLeaf assignStmts <> resBody.stmts, expr: resBody.expr, exprType: resBody.exprType, nextId: resBody.nextId }

          Accessor obj accessor ->
            let
              resObj = translateExpr codegenStateRef (depth + 1) modNameStr recVars moduleFunctions bound Nothing [] { isTail: false, inEffectBlock: false } nextId obj
            in
              case accessor of
                GetProp prop ->
                  case resObj.exprType of
                    TypeRecord fields ->
                      let
                        fieldGoType = fromMaybe TypeValue (Map.lookup prop (Map.fromFoldable fields))
                      in
                        { stmts: resObj.stmts, expr: GoStructAccess resObj.expr (sanitizeName prop), exprType: fieldGoType, nextId: resObj.nextId }
                    TypeStructPointer _ fullName _ _ ->
                      let
                        h = unsafePerformEffect (Ref.read codegenStateRef)

                      in
                        case Map.lookup fullName h.classDeclsFields of
                          Just info ->
                            case Array.findIndex (\f -> f.name == prop) info.fields of
                              Just idx ->
                                let
                                  unboxedObj = unboxGoExpr codegenStateRef modNameStr resObj.expr resObj.exprType resObj.exprType
                                  fieldExpr = GoStructAccess unboxedObj ("V" <> show idx)
                                  boxedFieldExpr = GoCall (GoSelector (GoVar "gopurs_runtime") "Box") [ fieldExpr ]
                                in
                                  { stmts: resObj.stmts, expr: boxedFieldExpr, exprType: TypeValue, nextId: resObj.nextId }
                              Nothing ->
                                let
                                  _ = unit
                                in
                                  { stmts: resObj.stmts, expr: GoRecordAccess (boxGoExpr codegenStateRef modNameStr resObj.expr resObj.exprType) prop, exprType: TypeValue, nextId: resObj.nextId }
                          Nothing -> { stmts: resObj.stmts, expr: GoRecordAccess (boxGoExpr codegenStateRef modNameStr resObj.expr resObj.exprType) prop, exprType: TypeValue, nextId: resObj.nextId }
                    TypeValue -> { stmts: resObj.stmts, expr: GoRecordAccess (boxGoExpr codegenStateRef modNameStr resObj.expr resObj.exprType) prop, exprType: TypeValue, nextId: resObj.nextId }
                    _ -> { stmts: resObj.stmts, expr: GoRecordAccess (boxGoExpr codegenStateRef modNameStr resObj.expr resObj.exprType) prop, exprType: TypeValue, nextId: resObj.nextId }
                GetIndex idx -> { stmts: resObj.stmts, expr: GoCall (GoSelector (GoVar "gopurs_runtime") "ArrayAccess") [ (boxGoExpr codegenStateRef modNameStr resObj.expr resObj.exprType), GoInt idx ], exprType: TypeValue, nextId: resObj.nextId }
                GetCtorField (Qualified mbMod _) _ _ (Ident ctorName) _ idx ->
                  let
                    defMod = case mbMod of
                      Just (ModuleName mod) -> String.replaceAll (Pattern ".") (Replacement "_") mod
                      Nothing -> modNameStr
                    structName = "Constructor_" <> defMod <> "_" <> sanitizeName ctorName
                    key = defMod <> "." <> ctorName
                    helpers = unsafePerformEffect (Ref.read codegenStateRef)
                  in
                    if Set.member structName elidedCtors then
                      { stmts: resObj.stmts, expr: coerceGoExpr codegenStateRef modNameStr resObj.expr resObj.exprType TypeValue, exprType: TypeValue, nextId: resObj.nextId }
                    else
                      let
                        fields = fromMaybe [] (map _.fields (Map.lookup key helpers.ctorTypes))
                        monoStructName = structName

                        expectedType = case Array.index fields idx of
                          Just ty -> exprTypeToGoType (unsafePerformEffect (Ref.read codegenStateRef)).pointerAdtPaths (unsafePerformEffect (Ref.read codegenStateRef)).enumAdts (unsafePerformEffect (Ref.read codegenStateRef)).elidedCtors modNameStr ty
                          Nothing -> TypeValue

                        typeArgs = case resObj.exprType of
                          TypeStructPointer _ _ _ tArgs -> tArgs
                          _ -> case getExprType obj of
                            ADT fullName _ tArgs ->
                              let
                                mapped = map (exprTypeToGoType (unsafePerformEffect (Ref.read codegenStateRef)).pointerAdtPaths (unsafePerformEffect (Ref.read codegenStateRef)).enumAdts (unsafePerformEffect (Ref.read codegenStateRef)).elidedCtors modNameStr) tArgs
                                arity = case Map.lookup fullName (unsafePerformEffect (Ref.read codegenStateRef)).pointerAdtPaths of
                                  Just info -> info.arity
                                  Nothing -> Array.length mapped
                              in
                                Array.take arity mapped
                            _ -> case Map.lookup key helpers.ctorTypes of
                              Just ctorInfo -> map (const TypeValue) ctorInfo.vars
                              Nothing -> []

                        isNative = case resObj.exprType of
                          TypeStructPointer _ _ _ _ -> true
                          _ -> false

                        actualFieldType = case resObj.exprType of
                          TypeStructPointer _ _ _ tArgs ->
                            case Map.lookup key helpers.ctorTypes of
                              Just ctorInfo ->
                                let
                                  env = Map.fromFoldable (Array.zip ctorInfo.vars tArgs)
                                  genericTy = structFieldGoType (unsafePerformEffect (Ref.read codegenStateRef)).pointerAdtPaths (unsafePerformEffect (Ref.read codegenStateRef)).enumAdts (unsafePerformEffect (Ref.read codegenStateRef)).elidedCtors ctorInfo.vars modNameStr (fromMaybe (TypeVar "") (Array.index fields idx))
                                in
                                  instantiateGenericGoType env genericTy
                              Nothing -> expectedType
                          _ -> expectedType

                        exprAccess =
                          if isNative then
                            GoConstructorAccess resObj.expr monoStructName typeArgs idx true
                          else
                            GoConstructorAccess (boxGoExpr codegenStateRef modNameStr resObj.expr resObj.exprType) monoStructName typeArgs idx false
                      in
                        { stmts: resObj.stmts, expr: exprAccess, exprType: actualFieldType, nextId: resObj.nextId }

          Update obj props ->
            let
              resObj = translateExpr codegenStateRef (depth + 1) modNameStr recVars moduleFunctions bound Nothing [] { isTail: false, inEffectBlock: false } nextId obj
              accProps = foldl
                ( \acc (Prop key val) ->
                    let
                      resVal = translateExpr codegenStateRef (depth + 1) modNameStr recVars moduleFunctions bound Nothing [] { isTail: false, inEffectBlock: false } acc.nextId val
                    in
                      { stmts: acc.stmts <> resVal.stmts, exprs: Array.snoc acc.exprs { key, expr: resVal.expr, goType: resVal.exprType }, exprType: TypeValue, nextId: resVal.nextId }
                )
                { stmts: StmtEmpty, exprs: [], exprType: TypeValue, nextId: resObj.nextId }
                props
            in
              case resObj.exprType of
                TypeRecord fields ->
                  let
                    coercedUpdates = map
                      ( \p ->
                          let
                            expectedGoType = fromMaybe TypeValue (Map.lookup p.key (Map.fromFoldable fields))
                            coercedVal = coerceGoExpr codegenStateRef modNameStr p.expr p.goType expectedGoType
                          in
                            Tuple p.key coercedVal
                      )
                      accProps.exprs
                  in
                    { stmts: resObj.stmts <> accProps.stmts, expr: GoRecordUpdateNative resObj.exprType resObj.expr coercedUpdates, exprType: resObj.exprType, nextId: accProps.nextId }
                _ ->
                  let
                    boxedExprs = map (\p -> Tuple p.key (boxGoExpr codegenStateRef modNameStr p.expr p.goType)) accProps.exprs
                  in
                    { stmts: resObj.stmts <> accProps.stmts, expr: GoRecordUpdateDict (boxGoExpr codegenStateRef modNameStr resObj.expr resObj.exprType) boxedExprs, exprType: TypeValue, nextId: accProps.nextId }

          CtorDef _ _ (Ident name) fields ->
            let
              helpers = unsafePerformEffect (Ref.read codegenStateRef)
              ctorType = getExprType tcoExpr
              expectedGoType = exprTypeToGoType (unsafePerformEffect (Ref.read codegenStateRef)).pointerAdtPaths (unsafePerformEffect (Ref.read codegenStateRef)).enumAdts (unsafePerformEffect (Ref.read codegenStateRef)).elidedCtors modNameStr ctorType

              trueModPart = case expectedGoType of
                TypeStructPointer _ fn _ _ ->
                  let
                    parts = String.split (Pattern ".") fn
                  in
                    String.joinWith "." (Array.slice 0 (Array.length parts - 1) parts)
                _ -> modNameStr

              trueModPartUnderscores = String.replaceAll (Pattern ".") (Replacement "_") trueModPart

              structName = "Constructor_" <> trueModPartUnderscores <> "_" <> sanitizeName name
              baseStructName = "Data_" <> trueModPartUnderscores <> "_" <> sanitizeName name
              key = trueModPartUnderscores <> "." <> name

              fullName = case expectedGoType of
                TypeStructPointer _ fn _ _ -> fn
                _ -> if key == "Test_RBTree.E" then "Test.RBTree.Tree" else key

              ctorInfo = Map.lookup key helpers.ctorTypes
              classInfo = Map.lookup fullName helpers.classDeclsFields

              fields' = case ctorInfo of
                Just info -> info.fields
                Nothing -> case classInfo of
                  Just info -> map _."type" info.fields
                  Nothing -> []

              vars' = case ctorInfo of
                Just info -> info.vars
                Nothing -> case classInfo of
                  Just info -> info.vars
                  Nothing -> []

              typeArgs = case getExprType tcoExpr of
                ADT fullName _ tArgs ->
                  let
                    mapped = map (exprTypeToGoType (unsafePerformEffect (Ref.read codegenStateRef)).pointerAdtPaths (unsafePerformEffect (Ref.read codegenStateRef)).enumAdts (unsafePerformEffect (Ref.read codegenStateRef)).elidedCtors modNameStr) tArgs
                    arity = case Map.lookup fullName (unsafePerformEffect (Ref.read codegenStateRef)).pointerAdtPaths of
                      Just info -> info.arity
                      Nothing -> Array.length vars'
                  in
                    Array.take arity mapped
                TypeApp fn arg ->
                  let
                    unwrapTypeApp (TypeApp f a) acc = unwrapTypeApp f (a <> acc)
                    unwrapTypeApp other acc = Tuple other acc
                  in
                    case unwrapTypeApp (getExprType tcoExpr) [] of
                      Tuple (ADT fnName _ tArgs) allArgs ->
                        let
                          mapped = map (exprTypeToGoType (unsafePerformEffect (Ref.read codegenStateRef)).pointerAdtPaths (unsafePerformEffect (Ref.read codegenStateRef)).enumAdts (unsafePerformEffect (Ref.read codegenStateRef)).elidedCtors modNameStr) (tArgs <> allArgs)
                        in
                          Array.take (Array.length vars') mapped
                      _ -> map (const TypeValue) vars'
                _ -> map (const TypeValue) vars'
              typeArgsStr = if Array.length typeArgs > 0 then "[" <> String.joinWith ", " (map goTypeToStr typeArgs) <> "]" else ""
              instMap = Map.fromFoldable (Array.zip vars' typeArgs)
              coercedFields = Array.mapWithIndex
                ( \i f ->
                    let
                      expectedType = case Array.index fields' i of
                        Just ty ->
                          let
                            genericGoType = exprTypeToGenericGoType helpers.pointerAdtPaths helpers.enumAdts helpers.elidedCtors vars' modNameStr ty
                          in
                            instantiateGenericGoType instMap genericGoType
                        Nothing -> TypeValue
                    in
                      coerceGoExpr codegenStateRef modNameStr (GoVar (sanitizeName f)) TypeValue expectedType
                )
                fields
              isElided = Set.member structName helpers.elidedCtors
              isPointerAdtLeaf = Map.member baseStructName helpers.pointerAdtLeaves
              isEnum = Set.member baseStructName helpers.enumCtors
              boxedCtor = boxGoExpr codegenStateRef modNameStr (GoConstructor (hashString baseStructName) structName typeArgs coercedFields) (TypeStructPointer baseStructName fullName (structName <> typeArgsStr) typeArgs)

              finalExprType =
                if isEnum then TypeUint32
                else if isPointerAdtLeaf then
                  let
                    nodeInfo = Map.lookup baseStructName helpers.pointerAdtLeaves
                    nodeBaseStruct = case nodeInfo of
                      Just info -> info.nodeBaseStruct
                      Nothing -> ""
                    nodeCtorName = case nodeInfo of
                      Just info -> info.nodeCtor
                      Nothing -> ""
                    nodeStruct = "Constructor_" <> modNameStr <> "_" <> sanitizeName nodeCtorName
                    nodeFullPath = nodeStruct <> typeArgsStr
                  in
                    TypeStructPointer nodeBaseStruct fullName nodeFullPath typeArgs
                else if Array.length fields == 0 then
                  TypeStructPointer baseStructName fullName (structName <> typeArgsStr) typeArgs
                else TypeValue

              funcExpr =
                if isElided then
                  case Array.head fields of
                    Just f -> GoCall (GoSelector (GoVar "gopurs_runtime") "Func") [ GoRaw ("func(" <> sanitizeName f <> " gopurs_runtime.Value) gopurs_runtime.Value {\nreturn " <> printGoExpr (coerceGoExpr codegenStateRef modNameStr (GoVar (sanitizeName f)) TypeValue TypeValue) <> "\n}") ]
                    Nothing -> Array.foldr (\f inner -> GoCall (GoSelector (GoVar "gopurs_runtime") "Func") [ GoRaw ("func(" <> sanitizeName f <> " gopurs_runtime.Value) gopurs_runtime.Value {\nreturn " <> printGoExpr inner <> "\n}") ]) boxedCtor fields
                else if isPointerAdtLeaf then
                  let
                    nodeCtorName = case Map.lookup baseStructName helpers.pointerAdtLeaves of
                      Just info -> info.nodeCtor
                      Nothing -> ""
                    nodeStruct = "Constructor_" <> modNameStr <> "_" <> sanitizeName nodeCtorName <> typeArgsStr
                  in
                    GoRaw ("(*" <> nodeStruct <> ")(nil)")
                else if isEnum then GoRaw (hashString baseStructName)
                else if Array.length fields == 0 then
                  GoConstructor (hashString baseStructName) structName typeArgs coercedFields
                else
                  Array.foldr (\f inner -> GoCall (GoSelector (GoVar "gopurs_runtime") "Func") [ GoRaw ("func(" <> sanitizeName f <> " gopurs_runtime.Value) gopurs_runtime.Value {\nreturn " <> printGoExpr inner <> "\n}") ]) boxedCtor fields
            in
              { stmts: StmtEmpty, expr: funcExpr, exprType: finalExprType, nextId }

          CtorSaturated (Qualified mbMod _) _ _ (Ident name) props ->
            let
              helpers = unsafePerformEffect (Ref.read codegenStateRef)

              ctorType = case getExprType tcoExpr of
                Any -> fromMaybe Any mbExpectedExprType
                ty -> 
                  if hasTypeVars ty then 
                    case mbExpectedExprType of
                      Just expectedTy | not (hasTypeVars expectedTy) -> expectedTy
                      _ -> ty
                  else ty
                  
              expectedGoType = exprTypeToGoType (unsafePerformEffect (Ref.read codegenStateRef)).pointerAdtPaths (unsafePerformEffect (Ref.read codegenStateRef)).enumAdts (unsafePerformEffect (Ref.read codegenStateRef)).elidedCtors modNameStr ctorType

              modPart = case mbMod of
                Just (ModuleName mn) -> mn
                Nothing -> modNameStr

              trueModPart = case expectedGoType of
                TypeStructPointer _ fn _ _ ->
                  let
                    parts = String.split (Pattern ".") fn
                  in
                    String.joinWith "." (Array.slice 0 (Array.length parts - 1) parts)
                _ -> modPart

              trueModPartUnderscores = String.replaceAll (Pattern ".") (Replacement "_") trueModPart

              baseStructName = "Data_" <> trueModPartUnderscores <> "_" <> sanitizeName name
              adtFullName = case ctorType of
                ADT fn _ _ -> Just fn
                _ -> Nothing
                
              structName = "Constructor_" <> trueModPartUnderscores <> "_" <> sanitizeName name
              key = trueModPartUnderscores <> "." <> name

              fullName = case expectedGoType of
                TypeStructPointer _ fn _ _ -> fn
                _ -> if key == "Test_RBTree.E" then "Test.RBTree.Tree" else key

              ctorInfo = Map.lookup key helpers.ctorTypes
              classInfo = Map.lookup fullName helpers.classDeclsFields

              fields = case ctorInfo of
                Just info -> info.fields
                Nothing -> case classInfo of
                  Just info -> map _."type" info.fields
                  Nothing -> []

              vars = case ctorInfo of
                Just info -> info.vars
                Nothing -> case classInfo of
                  Just info -> info.vars
                  Nothing -> []

              typeArgs = case ctorType of
                ADT fullName _ tArgs ->
                  let
                    mapped = map (exprTypeToGoType (unsafePerformEffect (Ref.read codegenStateRef)).pointerAdtPaths (unsafePerformEffect (Ref.read codegenStateRef)).enumAdts (unsafePerformEffect (Ref.read codegenStateRef)).elidedCtors modNameStr) tArgs
                    arity = case Map.lookup fullName (unsafePerformEffect (Ref.read codegenStateRef)).pointerAdtPaths of
                      Just info -> info.arity
                      Nothing -> Array.length vars
                  in
                    Array.take arity mapped
                _ -> map (const TypeValue) vars
              instMap = Map.fromFoldable (Array.zip vars typeArgs)

              accProps = foldl
                ( \acc (Tuple _ val) ->
                    let
                      expectedExprType = case Array.index fields acc.fieldIdx of
                        Just ty -> ty
                        Nothing -> Any
                      expectedType = case adtFullName >>= \fn -> Map.lookup fn unboxableADTs of
                        Just adt -> fromMaybe TypeValue (Array.index adt.signature acc.fieldIdx)
                        Nothing -> case Array.index fields acc.fieldIdx of
                          Just ty ->
                            let
                              genericGoType = exprTypeToGenericGoType helpers.pointerAdtPaths helpers.enumAdts helpers.elidedCtors vars modNameStr ty
                            in
                              instantiateGenericGoType instMap genericGoType
                          Nothing -> TypeValue

                      newBound = case unwrapTcoExpr val, extractExprFuncType expectedExprType of
                        Abs args _, Just { fArgs } ->
                          let
                            paramsWithTypes = Array.zipWith (\(Tuple mbI lvl) fArgTy -> Tuple (localId mbI lvl) (exprTypeToGoType helpers.pointerAdtPaths helpers.enumAdts helpers.elidedCtors modNameStr fArgTy)) (toArray args) (fArgs <> Array.replicate (Array.length (toArray args) - Array.length fArgs) Any)
                          in
                            foldl (\b (Tuple idStr goType) -> Map.insert idStr { name: idStr, goType } b) bound paramsWithTypes
                        UncurriedAbs args _, Just { fArgs } ->
                          let
                            paramsWithTypes = Array.zipWith (\(Tuple mbI lvl) fArgTy -> Tuple (localId mbI lvl) (exprTypeToGoType helpers.pointerAdtPaths helpers.enumAdts helpers.elidedCtors modNameStr fArgTy)) args (fArgs <> Array.replicate (Array.length args - Array.length fArgs) Any)
                          in
                            foldl (\b (Tuple idStr goType) -> Map.insert idStr { name: idStr, goType } b) bound paramsWithTypes
                        _, _ -> bound

                      resVal = translateExpr codegenStateRef (depth + 1) modNameStr recVars moduleFunctions newBound Nothing [] { isTail: false, inEffectBlock: false } acc.nextId val
                      coercedExpr = coerceGoExpr codegenStateRef modNameStr resVal.expr resVal.exprType expectedType
                      isConstant = expectedType == resVal.exprType && case expectedType, unwrapTcoExpr val of
                        TypeInt64, Lit (LitInt _) -> true
                        TypeBool, Lit (LitBoolean _) -> true
                        TypeUint32, CtorSaturated _ _ _ _ ctorFields -> Array.null ctorFields
                        TypeUint32, CtorDef _ _ _ ctorFields -> Array.null ctorFields
                        _, _ -> false
                    in
                      { stmts: acc.stmts <> resVal.stmts, exprs: Array.snoc acc.exprs coercedExpr, exprType: TypeValue, nextId: resVal.nextId, fieldIdx: acc.fieldIdx + 1, constants: Array.snoc acc.constants isConstant }
                )
                { stmts: StmtEmpty, exprs: [], exprType: TypeValue, nextId, fieldIdx: 0, constants: [] }
                props

              isElided = Set.member structName helpers.elidedCtors
              isPointerAdtLeaf = Map.member baseStructName helpers.pointerAdtLeaves

              modPart' = trueModPartUnderscores
              monoStructName = "Constructor_" <> modPart' <> "_" <> sanitizeName name

              typeArgsCtor = case ctorType of
                ADT fullName _ tArgs ->
                  let
                    mapped = map (exprTypeToGoType (unsafePerformEffect (Ref.read codegenStateRef)).pointerAdtPaths (unsafePerformEffect (Ref.read codegenStateRef)).enumAdts (unsafePerformEffect (Ref.read codegenStateRef)).elidedCtors modNameStr) tArgs
                    arity = case Map.lookup fullName (unsafePerformEffect (Ref.read codegenStateRef)).pointerAdtPaths of
                      Just info -> info.arity
                      Nothing -> Array.length mapped
                  in
                    Array.take arity mapped
                _ -> case Map.lookup key helpers.ctorTypes of
                  Just ctorInfo -> map (const TypeValue) ctorInfo.vars
                  Nothing -> []
              typeArgsStr = if Array.length typeArgsCtor > 0 then "[" <> String.joinWith ", " (map goTypeToStr typeArgsCtor) <> "]" else ""
              fullPath = monoStructName <> typeArgsStr
              isEnum = Set.member baseStructName helpers.enumCtors
              res =
                if isElided then
                  case Array.head accProps.exprs of
                    Just expr -> { expr: boxGoExpr codegenStateRef modNameStr expr (fromMaybe TypeValue (map (exprTypeToGoType (unsafePerformEffect (Ref.read codegenStateRef)).pointerAdtPaths (unsafePerformEffect (Ref.read codegenStateRef)).enumAdts (unsafePerformEffect (Ref.read codegenStateRef)).elidedCtors modNameStr) (Array.index fields 0))), exprType: TypeValue }
                    Nothing -> { expr: GoConstructor (hashString baseStructName) monoStructName typeArgsCtor accProps.exprs, exprType: TypeStructPointer baseStructName fullName fullPath typeArgsCtor }
                else case adtFullName >>= \fn -> Map.lookup fn unboxableADTs >>= \adt -> Just (Tuple fn adt) of
                  Just (Tuple fn adt) ->
                    { expr: GoStructValue fn adt.signature (adt.mapConstructor name accProps.exprs)
                    , exprType: TypeStructValue fn adt.signature
                    }
                  Nothing ->
                    if isPointerAdtLeaf then
                      let
                        nodeInfo = Map.lookup baseStructName helpers.pointerAdtLeaves
                        nodeCtorName = case nodeInfo of
                          Just info -> info.nodeCtor
                          Nothing -> ""
                        nodeBaseStruct = case nodeInfo of
                          Just info -> info.nodeBaseStruct
                          Nothing -> ""
                        nodeStruct = "Constructor_" <> modPart' <> "_" <> sanitizeName nodeCtorName
                        nodeFullPath = nodeStruct <> typeArgsStr
                      in
                        { expr: GoRaw ("(*" <> nodeFullPath <> ")(nil)"), exprType: TypeStructPointer nodeBaseStruct fullName nodeFullPath typeArgsCtor }
                    else if isEnum then { expr: GoRaw (hashString baseStructName), exprType: TypeUint32 }
                    else { expr: GoConstructor (hashString baseStructName) monoStructName typeArgsCtor accProps.exprs, exprType: TypeStructPointer baseStructName fullName fullPath typeArgsCtor }
              reuse = case accProps.stmts of
                StmtEmpty -> constructorReuse bound res.exprType accProps.constants res.expr
                _ -> Nothing
            in
              case reuse of
                Just { source, condition } ->
                  let
                    resultName = "__reuse_" <> show accProps.nextId
                    declare = GoRaw ("var " <> resultName <> " " <> goTypeToStr res.exprType)
                    choose = GoIfElse condition [ GoMutate resultName source ] [ GoMutate resultName res.expr ]
                  in
                    { stmts: StmtLeaf declare <> StmtLeaf choose, expr: GoVar resultName, exprType: res.exprType, nextId: accProps.nextId + 1 }
                Nothing ->
                  { stmts: accProps.stmts, expr: res.expr, exprType: res.exprType, nextId: accProps.nextId }

          Fail msg ->
            let
              expectedGoType = exprTypeToGoType (unsafePerformEffect (Ref.read codegenStateRef)).pointerAdtPaths (unsafePerformEffect (Ref.read codegenStateRef)).enumAdts (unsafePerformEffect (Ref.read codegenStateRef)).elidedCtors modNameStr (case getExprType tcoExpr of
                                Any -> fromMaybe Any mbExpectedExprType
                                ty -> ty)
              expectedGoTypeStr = goTypeToStr expectedGoType
            in
            { stmts: StmtEmpty, expr: GoRaw ("func() " <> expectedGoTypeStr <> " { panic(" <> printGoExpr (GoString msg) <> ") }()"), exprType: expectedGoType, nextId }

          Branch branches def ->
            let
              resDef = translateExprWithExpectedType codegenStateRef (depth + 1) modNameStr recVars moduleFunctions bound Nothing loopCtx { isTail, inEffectBlock: false } mbExpectedExprType nextId def

              computedBranches = foldl
                ( \acc (Pair condExpr bodyExpr) ->
                    let
                      resCond = translateExpr codegenStateRef (depth + 1) modNameStr recVars moduleFunctions bound Nothing [] { isTail: false, inEffectBlock: false } acc.nextId condExpr
                      resBody = translateExprWithExpectedType codegenStateRef (depth + 1) modNameStr recVars moduleFunctions bound Nothing loopCtx { isTail, inEffectBlock: false } mbExpectedExprType resCond.nextId bodyExpr
                    in
                      { nextId: resBody.nextId, results: acc.results <> [ { cond: resCond, body: resBody } ] }
                )
                { nextId: resDef.nextId, results: [] }
                (toArray branches)

              isFailNode = case unwrapTcoExpr def of
                Fail _ -> true
                _ -> false

              allTypes = (if isFailNode then [] else [ resDef.exprType ]) <> map (\r -> r.body.exprType) computedBranches.results
              
              hasTypeValue = Array.any
                ( \t -> case t of
                    TypeValue -> true
                    _ -> false
                )
                allTypes

              expectedGoType =
                if hasTypeValue then TypeValue
                else
                  let
                    nubbed = Array.nub (map goTypeToStr allTypes)
                  in
                    if Array.length nubbed == 1 then fromMaybe TypeValue (Array.head allTypes)
                    else TypeValue

              tmpVar = "__t" <> show computedBranches.nextId
              declTmp = StmtLeaf (GoRaw ("var " <> tmpVar <> " " <> goTypeToStr expectedGoType))
              labelName = "end_branch_" <> show computedBranches.nextId

              buildIfs = foldl
                ( \acc r ->
                    let
                      goIf = GoIfElse (unboxGoExpr codegenStateRef modNameStr r.cond.expr r.cond.exprType TypeBool) (flattenStmts r.body.stmts <> [ GoMutate tmpVar (coerceGoExpr codegenStateRef modNameStr r.body.expr r.body.exprType expectedGoType), GoRaw ("goto " <> labelName) ]) []
                    in
                      acc <> StmtLeaf (GoRaw "{") <> r.cond.stmts <> StmtLeaf goIf <> StmtLeaf (GoRaw "}")
                )
                StmtEmpty
                computedBranches.results
            in
              { stmts: declTmp <> buildIfs <> StmtLeaf (GoRaw "{") <> resDef.stmts <> StmtLeaf (GoMutate tmpVar (coerceGoExpr codegenStateRef modNameStr resDef.expr resDef.exprType expectedGoType)) <> StmtLeaf (GoRaw "}") <> StmtLeaf (GoRaw (labelName <> ":")), expr: GoVar tmpVar, exprType: expectedGoType, nextId: computedBranches.nextId + 1 }

          PrimOp (Op1 op1 e)
            | Just emit <- PrimitiveExprs.unary codegenStateRef modNameStr op1 ->
                let
                  resE = translateExpr codegenStateRef (depth + 1) modNameStr recVars moduleFunctions bound Nothing [] { isTail: false, inEffectBlock: false } nextId e
                  result = emit { expr: resE.expr, exprType: resE.exprType }
                in
                  { stmts: resE.stmts, expr: result.expr, exprType: result.exprType, nextId: resE.nextId }

          PrimOp (Op1 (OpIsTag (Qualified mbMod (Ident tag))) e) ->
            let
              resE = translateExpr codegenStateRef (depth + 1) modNameStr recVars moduleFunctions bound Nothing [] { isTail: false, inEffectBlock: false } nextId e
            in
              let
                baseStructName = getBaseStructName modNameStr mbMod tag
                hashStr = hashString baseStructName
                helpers = unsafePerformEffect (Ref.read codegenStateRef)

                isNativePointer = case resE.exprType of
                  TypeStructPointer typedBaseStructName _ _ _ ->
                    typedBaseStructName == baseStructName ||
                      ( case Map.lookup baseStructName helpers.pointerAdtLeaves of
                          Just nodeInfo -> typedBaseStructName == nodeInfo.nodeBaseStruct
                          Nothing -> false
                      )
                  _ -> false
              in
                case resE.expr of
                  GoVar _ ->
                    let
                      exprStr =
                        if isNativePointer then
                          case Map.lookup baseStructName helpers.pointerAdtLeaves of
                            Just _ -> "(" <> printGoExpr resE.expr <> " == nil)"
                            Nothing -> "(" <> printGoExpr resE.expr <> " != nil)"
                        else case Map.lookup baseStructName helpers.pointerAdtLeaves of
                          Just nodeInfo -> "(" <> printGoExpr (boxGoExpr codegenStateRef modNameStr resE.expr resE.exprType) <> ".Type == 9 && " <> printGoExpr (boxGoExpr codegenStateRef modNameStr resE.expr resE.exprType) <> ".IntVal == " <> hashString nodeInfo.nodeBaseStruct <> " && " <> printGoExpr (boxGoExpr codegenStateRef modNameStr resE.expr resE.exprType) <> ".UnsafePtr == nil)"
                          Nothing ->
                            if Set.member baseStructName helpers.pointerAdtNodes then
                              "(" <> printGoExpr (boxGoExpr codegenStateRef modNameStr resE.expr resE.exprType) <> ".Type == 9 && " <> printGoExpr (boxGoExpr codegenStateRef modNameStr resE.expr resE.exprType) <> ".IntVal == " <> hashStr <> " && " <> printGoExpr (boxGoExpr codegenStateRef modNameStr resE.expr resE.exprType) <> ".UnsafePtr != nil)"
                            else if Set.member baseStructName helpers.enumCtors then
                              "(" <> printGoExpr (unboxGoExpr codegenStateRef modNameStr resE.expr resE.exprType TypeUint32) <> " == " <> hashStr <> ")"
                            else
                              "(" <> printGoExpr (boxGoExpr codegenStateRef modNameStr resE.expr resE.exprType) <> ".Type == 9 && " <> printGoExpr (boxGoExpr codegenStateRef modNameStr resE.expr resE.exprType) <> ".IntVal == " <> hashStr <> ")"
                    in
                      { stmts: resE.stmts, expr: GoRaw exprStr, exprType: TypeBool, nextId: resE.nextId }
                  _ ->
                    let
                      tmpVar = "__t_tag_" <> show resE.nextId
                      declTmp =
                        if isNativePointer || resE.exprType /= TypeValue then
                          StmtLeaf (GoRaw ("var " <> tmpVar <> " " <> goTypeToStr resE.exprType <> " = " <> printGoExpr resE.expr))
                        else
                          StmtLeaf (GoRaw ("var " <> tmpVar <> " gopurs_runtime.Value = " <> printGoExpr (boxGoExpr codegenStateRef modNameStr resE.expr resE.exprType)))

                      exprStr =
                        if isNativePointer then
                          case Map.lookup baseStructName helpers.pointerAdtLeaves of
                            Just _ -> "(" <> tmpVar <> " == nil)"
                            Nothing -> "(" <> tmpVar <> " != nil)"
                        else if resE.exprType /= TypeValue then
                          "(uint32(" <> tmpVar <> ") == " <> hashStr <> ")"
                        else case Map.lookup baseStructName helpers.pointerAdtLeaves of
                          Just nodeInfo -> "(" <> tmpVar <> ".Type == 9 && " <> tmpVar <> ".IntVal == " <> hashString nodeInfo.nodeBaseStruct <> " && " <> tmpVar <> ".UnsafePtr == nil)"
                          Nothing ->
                            if Set.member baseStructName helpers.pointerAdtNodes then
                              "(" <> tmpVar <> ".Type == 9 && " <> tmpVar <> ".IntVal == " <> hashStr <> " && " <> tmpVar <> ".UnsafePtr != nil)"
                            else if Set.member baseStructName helpers.enumCtors then
                              "(" <> printGoExpr (unboxGoExpr codegenStateRef modNameStr (GoVar tmpVar) TypeValue TypeUint32) <> " == " <> hashStr <> ")"
                            else
                              "(" <> tmpVar <> ".Type == 9 && " <> tmpVar <> ".IntVal == " <> hashStr <> ")"
                    in
                      { stmts: resE.stmts <> declTmp, expr: GoRaw exprStr, exprType: TypeBool, nextId: resE.nextId + 1 }

          PrimOp (Op1 OpArrayLength e) ->
            let
              resE = translateExpr codegenStateRef (depth + 1) modNameStr recVars moduleFunctions bound Nothing [] { isTail: false, inEffectBlock: false } nextId e
            in
              case resE.exprType of
                TypeNativeArray _ -> { stmts: resE.stmts, expr: GoCall (GoSelector (GoVar "gopurs_runtime") "Int") [ GoCall (GoVar "int64") [ GoCall (GoVar "len") [ resE.expr ] ] ], exprType: TypeValue, nextId: resE.nextId }
                _ -> { stmts: resE.stmts, expr: GoCall (GoSelector (GoVar "gopurs_runtime") "Int") [ GoCall (GoVar "int64") [ GoCall (GoSelector (GoVar "gopurs_runtime") "ArrayLength") [ boxGoExpr codegenStateRef modNameStr resE.expr resE.exprType ] ] ], exprType: TypeValue, nextId: resE.nextId }

          PrimOp (Op2 OpBooleanAnd e1 e2) ->
            let
              res1 = translateExpr codegenStateRef (depth + 1) modNameStr recVars moduleFunctions bound Nothing [] { isTail: false, inEffectBlock: false } nextId e1
              res2 = translateExpr codegenStateRef (depth + 1) modNameStr recVars moduleFunctions bound Nothing [] { isTail: false, inEffectBlock: false } res1.nextId e2
              isEmptyStmts StmtEmpty = true
              isEmptyStmts (StmtAppend s1 s2) = isEmptyStmts s1 && isEmptyStmts s2
              isEmptyStmts _ = false
            in
              if isEmptyStmts res2.stmts then
                { expr: GoBinOp "&&" (unboxGoExpr codegenStateRef modNameStr res1.expr res1.exprType TypeBool) (unboxGoExpr codegenStateRef modNameStr res2.expr res2.exprType TypeBool), exprType: TypeBool, stmts: res1.stmts <> res2.stmts, nextId: res2.nextId }
              else
                let
                  tmpVar = "__t_and_" <> show res2.nextId
                  declTmp = StmtLeaf (GoRaw ("var " <> tmpVar <> " bool = false\nif " <> printGoExpr (unboxGoExpr codegenStateRef modNameStr res1.expr res1.exprType TypeBool) <> " {\n"))
                  assignTmp = StmtLeaf (GoRaw (tmpVar <> " = " <> printGoExpr (unboxGoExpr codegenStateRef modNameStr res2.expr res2.exprType TypeBool) <> "\n}"))
                in
                  { expr: GoRaw tmpVar, exprType: TypeBool, stmts: res1.stmts <> declTmp <> res2.stmts <> assignTmp, nextId: res2.nextId + 1 }

          PrimOp (Op2 OpBooleanOr e1 e2) ->
            let
              res1 = translateExpr codegenStateRef (depth + 1) modNameStr recVars moduleFunctions bound Nothing [] { isTail: false, inEffectBlock: false } nextId e1
              res2 = translateExpr codegenStateRef (depth + 1) modNameStr recVars moduleFunctions bound Nothing [] { isTail: false, inEffectBlock: false } res1.nextId e2
              isEmptyStmts StmtEmpty = true
              isEmptyStmts (StmtAppend s1 s2) = isEmptyStmts s1 && isEmptyStmts s2
              isEmptyStmts _ = false
            in
              if isEmptyStmts res2.stmts then
                { expr: GoBinOp "||" (unboxGoExpr codegenStateRef modNameStr res1.expr res1.exprType TypeBool) (unboxGoExpr codegenStateRef modNameStr res2.expr res2.exprType TypeBool), exprType: TypeBool, stmts: res1.stmts <> res2.stmts, nextId: res2.nextId }
              else
                let
                  tmpVar = "__t_or_" <> show res2.nextId
                  declTmp = StmtLeaf (GoRaw ("var " <> tmpVar <> " bool = true\nif !(" <> printGoExpr (unboxGoExpr codegenStateRef modNameStr res1.expr res1.exprType TypeBool) <> ") {\n"))
                  assignTmp = StmtLeaf (GoRaw (tmpVar <> " = " <> printGoExpr (unboxGoExpr codegenStateRef modNameStr res2.expr res2.exprType TypeBool) <> "\n}"))
                in
                  { expr: GoRaw tmpVar, exprType: TypeBool, stmts: res1.stmts <> declTmp <> res2.stmts <> assignTmp, nextId: res2.nextId + 1 }

          PrimOp (Op2 OpArrayIndex e1 e2) ->
            let
              res1 = translateExpr codegenStateRef (depth + 1) modNameStr recVars moduleFunctions bound Nothing [] { isTail: false, inEffectBlock: false } nextId e1
              res2 = translateExpr codegenStateRef (depth + 1) modNameStr recVars moduleFunctions bound Nothing [] { isTail: false, inEffectBlock: false } res1.nextId e2
              result = case res1.exprType of
                TypeNativeArray innerType -> { expr: boxGoExpr codegenStateRef modNameStr (GoRaw (printGoExpr res1.expr <> "[" <> printGoExpr (unboxGoExpr codegenStateRef modNameStr res2.expr res2.exprType TypeInt64) <> "]")) innerType, exprType: TypeValue }
                _ -> { expr: GoCall (GoSelector (GoVar "gopurs_runtime") "ArrayAccess") [ boxGoExpr codegenStateRef modNameStr res1.expr res1.exprType, GoCall (GoVar "int") [ unboxGoExpr codegenStateRef modNameStr res2.expr res2.exprType TypeInt64 ] ], exprType: TypeValue }
            in
              { stmts: res1.stmts <> res2.stmts, expr: result.expr, exprType: result.exprType, nextId: res2.nextId }

          PrimOp (Op2 op2 e1 e2)
            | Just emit <- PrimitiveExprs.binary codegenStateRef modNameStr op2 ->
                let
                  res1 = translateExpr codegenStateRef (depth + 1) modNameStr recVars moduleFunctions bound Nothing [] { isTail: false, inEffectBlock: false } nextId e1
                  res2 = translateExpr codegenStateRef (depth + 1) modNameStr recVars moduleFunctions bound Nothing [] { isTail: false, inEffectBlock: false } res1.nextId e2
                  result = emit { expr: res1.expr, exprType: res1.exprType } { expr: res2.expr, exprType: res2.exprType }
                in
                  { stmts: res1.stmts <> res2.stmts, expr: result.expr, exprType: result.exprType, nextId: res2.nextId }

          PrimEffect eff -> case eff of
            EffectRefNew a ->
              let
                resA = translateExpr codegenStateRef (depth + 1) modNameStr recVars moduleFunctions bound Nothing [] { isTail: false, inEffectBlock: false } nextId a
                refIdent = "__local_ref_" <> show resA.nextId
                declStmt = GoAssign refIdent (boxGoExpr codegenStateRef modNameStr resA.expr resA.exprType)
                ifaceIdent = "__local_iface_" <> show resA.nextId
                ifaceStmt = GoRaw ("var " <> ifaceIdent <> " interface{} = " <> refIdent)
              in
                { stmts: resA.stmts <> StmtLeaf declStmt <> StmtLeaf ifaceStmt
                , expr: GoRaw ("gopurs_runtime.Any(&" <> ifaceIdent <> ")")
                , exprType: TypeValue
                , nextId: resA.nextId + 1
                }
            EffectRefRead a ->
              let
                resA = translateExpr codegenStateRef (depth + 1) modNameStr recVars moduleFunctions bound Nothing [] { isTail: false, inEffectBlock: false } nextId a
              in
                { stmts: resA.stmts
                , expr: GoRaw ("(*(" <> printGoExpr resA.expr <> ".PtrVal().(*interface{}))).(gopurs_runtime.Value)")
                , exprType: TypeValue
                , nextId: resA.nextId
                }
            EffectRefWrite ref val ->
              let
                resRef = translateExpr codegenStateRef (depth + 1) modNameStr recVars moduleFunctions bound Nothing [] { isTail: false, inEffectBlock: false } nextId ref
                resVal = translateExpr codegenStateRef (depth + 1) modNameStr recVars moduleFunctions bound Nothing [] { isTail: false, inEffectBlock: false } resRef.nextId val
                writeStmt = GoRaw ("*(" <> printGoExpr resRef.expr <> ".PtrVal().(*interface{})) = " <> printGoExpr (boxGoExpr codegenStateRef modNameStr resVal.expr resVal.exprType))
              in
                { stmts: resRef.stmts <> resVal.stmts <> StmtLeaf writeStmt
                , expr: boxGoExpr codegenStateRef modNameStr resVal.expr resVal.exprType
                , exprType: TypeValue
                , nextId: resVal.nextId
                }

          _ -> { stmts: StmtEmpty, expr: GoVar "gopurs_runtime.Value{}", exprType: TypeValue, nextId }

printTypeNode :: TypeNode -> String
printTypeNode = FfiBridge.printTypeNode

exprTypeToDummyTypeNode :: ExprType -> TypeNode
exprTypeToDummyTypeNode = FfiBridge.exprTypeToDummyTypeNode

getTastReturnType :: ExprType -> Maybe ExprType
getTastReturnType = FfiBridge.getTastReturnType

resolveNewtype :: Array DataDecl -> ExprType -> ExprType
resolveNewtype = FfiBridge.resolveNewtype

flattenFuncArgs :: ExprType -> Array ExprType
flattenFuncArgs = FfiBridge.flattenFuncArgs

unwrapValueToFunc :: Array DataDecl -> TypeNode -> Maybe ExprType -> String -> Int -> Int -> String
unwrapValueToFunc = FfiBridge.unwrapValueToFunc

boxFfiValue :: TypeNode -> String -> String
boxFfiValue = FfiBridge.boxFfiValue

wrapReturn :: Array DataDecl -> TypeNode -> Maybe ExprType -> String -> String
wrapReturn = FfiBridge.wrapReturn

printExprType :: ExprType -> String
printExprType = GoTypes.printExprType

getTastArgType :: ExprType -> Int -> Maybe ExprType
getTastArgType = FfiBridge.getTastArgType

isStandardPursFunc :: TypeNode -> Boolean
isStandardPursFunc = FfiBridge.isStandardPursFunc

generateWrapperFunc :: Array DataDecl -> FfiDecl -> Maybe ExprType -> String
generateWrapperFunc = FfiBridge.generateWrapperFunc

generateFfiBridge :: String -> Array DataDecl -> Array FfiDecl -> Array (Tuple Ident (Maybe ExprType)) -> String
generateFfiBridge = FfiBridge.generateFfiBridge

hasTypeVars :: ExprType -> Boolean
hasTypeVars = case _ of
  TypeVar _ -> true
  Array t -> hasTypeVars t
  ADT _ _ ts -> any hasTypeVars ts
  TypeApp t ts -> hasTypeVars t || any hasTypeVars ts
  Func ts t -> any hasTypeVars ts || hasTypeVars t
  Record t -> hasTypeVars t
  Row ts tail ->
    any (\(Tuple _ t) -> hasTypeVars t) ts ||
      case tail of
        Just t -> hasTypeVars t
        Nothing -> false
  ForAll _ t -> hasTypeVars t
  ConstrainedType cs t -> any (\(Tuple _ ts) -> any hasTypeVars ts) cs || hasTypeVars t
  _ -> false
