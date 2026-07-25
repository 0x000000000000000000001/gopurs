module Gopurs.CodeGen where

import Prelude
import PureScript.Backend.Optimizer.Syntax (BackendSyntax(..), BackendAccessor(..), Pair(..), Level(..), BackendOperator(..), BackendOperator1(..), BackendOperator2(..), BackendOperatorOrd(..), BackendOperatorNum(..), BackendEffect(..))
import PureScript.Backend.Optimizer.Semantics (NeutralExpr(..))
import PureScript.Backend.Optimizer.Convert (BackendModule, BackendBindingGroup)
import Data.String as String
import Data.Array as Array
import Data.Maybe (Maybe(..), fromMaybe, isJust)
import Data.Newtype (unwrap)
import PureScript.Backend.Optimizer.CoreFn (Ann(..), Bind(..), Binder(..), Binding(..), CaseAlternative(..), CaseGuard(..), Comment, ConstructorType(..), DataConstructor, DataDecl, Expr(..), ExprType(..), Guard(..), Ident(..), Import(..), Literal(..), Meta(..), Module(..), ModuleName(..), Prop(..), ProperName(..), Qualified(..), ReExport, exprAnn, findProp, propKey, propValue, qualifiedModuleName, unQualified)
import Data.Tuple (Tuple(..), fst, snd)
import Data.Array.NonEmpty as NonEmptyArray
import Data.Array.NonEmpty (NonEmptyArray, fromArray, toArray)
import Effect.Console as Console
import Effect.Unsafe (unsafePerformEffect)
import Effect.Ref (Ref)
import Effect.Ref as Ref
import Gopurs.FreeVars (freeVars)
import Data.Set as Set
import Data.String.CodeUnits as SCU
import Data.String.Pattern (Pattern(..), Replacement(..))
import Data.Map (Map)
import Debug as Debug
import Data.Map as Map
import Data.Set (Set)
import Data.Set as Set
import Data.Foldable (foldl, foldMap)
import Data.List as List
import Data.Traversable (traverse)

import Gopurs.GoAst (GoFile, GoDecl, GoExpr(..))


import Gopurs.Printer (printGoFile, printGoExpr, printGoDeclVar)
import PureScript.Backend.Optimizer.Codegen.Tco as Tco
import PureScript.Backend.Optimizer.Codegen.Tco (TcoExpr(..), tcoAnalysisOf)
import Gopurs.FreeVars (freeVars, localId, paramTypes)
import Node.Path as Path
import Node.FS.Sync as FS
import Node.Encoding (Encoding(..))
import Gopurs.FfiSupport (hashString)

data GoType = TypeValue | TypeInt64 | TypeFloat64 | TypeString | TypeBool
derive instance eqGoType :: Eq GoType

boxGoExpr :: GoExpr -> GoType -> GoExpr
boxGoExpr expr TypeValue = expr
boxGoExpr expr TypeInt64 = GoCall (GoSelector (GoVar "gopurs_runtime") "Int") [ expr ]
boxGoExpr expr TypeFloat64 = GoCall (GoSelector (GoVar "gopurs_runtime") "Float") [ expr ]
boxGoExpr expr TypeString = GoCall (GoSelector (GoVar "gopurs_runtime") "Str") [ expr ]
boxGoExpr expr TypeBool = GoCall (GoSelector (GoVar "gopurs_runtime") "Bool") [ expr ]

exprTypeToGoType :: ExprType -> GoType
exprTypeToGoType Int = TypeInt64
exprTypeToGoType Number = TypeFloat64
exprTypeToGoType String = TypeString
exprTypeToGoType Boolean = TypeBool
exprTypeToGoType _ = TypeValue

goTypeToStr :: GoType -> String
goTypeToStr TypeInt64 = "int64"
goTypeToStr TypeFloat64 = "float64"
goTypeToStr TypeString = "string"
goTypeToStr TypeBool = "bool"
goTypeToStr TypeValue = "gopurs_runtime.Value"

unboxGoExpr :: GoExpr -> GoType -> GoType -> GoExpr
unboxGoExpr expr currentType desiredType =
  if currentType == desiredType then expr
  else case desiredType of
    TypeValue -> boxGoExpr expr currentType
    TypeInt64 -> GoSelector expr "IntVal"
    TypeFloat64 -> GoCall (GoSelector expr "FloatVal") []
    TypeString -> GoCall (GoSelector expr "StrVal") []
    TypeBool -> GoBinOp "!=" (GoSelector expr "IntVal") (GoInt 0)


capitalize :: String -> String
capitalize "" = ""
capitalize s =
  let
    firstChar = String.take 1 s
  in
    if firstChar >= "a" && firstChar <= "z" then String.toUpper firstChar <> String.drop 1 s
    else if firstChar == "_" then "_" <> capitalize (String.drop 1 s)
    else s

sanitizeName :: String -> String
sanitizeName name =
  let
    s1 = String.replaceAll (Pattern "'") (Replacement "_prime") (String.replaceAll (Pattern "$") (Replacement "_dollar") name)
  in
    if s1 == "break" || s1 == "default" || s1 == "func" || s1 == "interface" || s1 == "select" || s1 == "case" || s1 == "defer" || s1 == "go" || s1 == "map" || s1 == "struct" || s1 == "chan" || s1 == "else" || s1 == "goto" || s1 == "package" || s1 == "switch" || s1 == "const" || s1 == "fallthrough" || s1 == "if" || s1 == "range" || s1 == "type" || s1 == "continue" || s1 == "for" || s1 == "import" || s1 == "return" || s1 == "var" || s1 == "init" || s1 == "append" || s1 == "make" || s1 == "len" || s1 == "cap" || s1 == "new" || s1 == "close" || s1 == "delete" || s1 == "complex" || s1 == "real" || s1 == "imag" || s1 == "panic" || s1 == "recover" || s1 == "print" || s1 == "println" then s1 <> "_" else s1

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

wrapInStmts :: Array String -> StmtTree -> GoExpr -> GoExpr
wrapInStmts _ stmts expr =
  let
    stmtsArr = flattenStmts stmts
  in
    if Array.length stmtsArr == 0 then expr
    else GoRaw ("func() gopurs_runtime.Value {\n" <> printGoExpr (GoBlock (stmtsArr <> [ GoReturn expr ])) <> "\n}()")

extractUncurriedAbs :: TcoExpr -> Maybe { args :: Array String, body :: TcoExpr, fvs :: Set String }
extractUncurriedAbs tcoExpr@(TcoExpr _ syntax) = case syntax of
  UncurriedAbs args body ->
    Just { args: map (\(Tuple mbI lvl) -> localId mbI lvl) args, body, fvs: freeVars tcoExpr }
  Abs args body ->
    let
      thisArgs = map (\(Tuple mbI lvl) -> localId mbI lvl) (toArray args)
    in case extractUncurriedAbs body of
      Just inner -> Just { args: thisArgs <> inner.args, body: inner.body, fvs: Set.union (freeVars tcoExpr) inner.fvs }
      Nothing -> Just { args: thisArgs, body, fvs: freeVars tcoExpr }
  Typed _ inner -> extractUncurriedAbs inner
  _ -> Nothing

unwrapExpr :: TcoExpr -> BackendSyntax TcoExpr
unwrapExpr (TcoExpr _ e) = e

flattenApp :: TcoExpr -> Tuple TcoExpr (Array TcoExpr)
flattenApp e =
  case unwrapTcoExpr e of
    App f args ->
      let
        Tuple f' args' = flattenApp f
      in
        Tuple f' (args' <> toArray args)
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
  let
    pkgPrefix = case mbMod of
      Just mn | sanitizeName (String.replaceAll (Pattern ".") (Replacement "_") (unwrap mn)) /= modNameStr -> "pkg_" <> sanitizeName (String.replaceAll (Pattern ".") (Replacement "_") (unwrap mn)) <> "."
      _ -> ""
  in
    pkgPrefix <> getBaseStructName modNameStr mbMod ctorName

translate :: Array (Array String) -> BackendModule -> String
translate importsArray mod =
  let
    modNameStr = String.replaceAll (Pattern ".") (Replacement "_") (unwrap mod.name)
    _ = unsafePerformEffect (Console.log ("Translating module " <> modNameStr))
    helpersRef = unsafePerformEffect do
      let
        structDecls = Array.concatMap (\decl ->
          map (\ctor ->
            let
              fields = Array.mapWithIndex (\i _ -> "V" <> show i <> " gopurs_runtime.Value") ctor.fieldTypes
              structName = "Data_" <> modNameStr <> "_" <> sanitizeName ctor.constructorName
              hashStr = hashString structName
              isTagHelper = "func Is_" <> structName <> "(v gopurs_runtime.Value) bool {\n\treturn v.Type == 9 && v.IntVal == " <> hashStr <> "\n}"
            in
              "type " <> structName <> " struct {\n\t" <> String.joinWith "\n\t" fields <> "\n}\n" <> isTagHelper
          ) decl.constructors
        ) mod.dataDecls
      Ref.new { decls: [], rawDecls: structDecls }

    Tuple _ tcoBindings = foldl
      ( \(Tuple env acc) group ->
          let
            neBindings = fromArray group.bindings
            _ = unsafePerformEffect (Console.log ("Translating binding group"))
            env' = case neBindings of
              Just ne | group.recursive -> Tco.topLevelTcoEnvGroup mod.name ne <> env
              _ -> env
            tcoBinds = map (\(Tuple k v) -> Tuple k (Tco.analyze env' v)) group.bindings
          in
            Tuple env' (Array.snoc acc { recursive: group.recursive, bindings: tcoBinds })
      )
      (Tuple [] [])
      mod.bindings

    extractFuncType :: BackendSyntax TcoExpr -> Maybe { fArgs :: Array ExprType, fRet :: ExprType }
    extractFuncType (Typed (Func a r) _) =
      let
        flattenFuncType acc (Func args ret) = flattenFuncType (acc <> args) ret
        flattenFuncType acc ret = { fArgs: acc, fRet: ret }
      in Just (flattenFuncType a r)
    extractFuncType (Typed _ inner) = extractFuncType (unwrapExpr inner)
    extractFuncType _ = Nothing

    moduleArities :: Map String { fullName :: String, fArgs :: Array GoType, fRet :: GoType, arity :: Int }
    moduleArities = Map.fromFoldable $ Array.concatMap
      ( \group ->
          Array.concatMap
            ( \(Tuple (Ident name) val) ->
                case extractUncurriedAbs val of
                  Just { args, body } -> 
                    let
                      fullName = "Call_" <> sanitizeName name
                      typeSig = extractFuncType (unwrapExpr val)
                      fArgsGo = case typeSig of
                        Just { fArgs } -> map exprTypeToGoType fArgs
                        Nothing -> Array.replicate (Array.length args) TypeValue
                      fRetGo = case typeSig of
                        Just { fRet } -> exprTypeToGoType fRet
                        Nothing -> TypeValue
                    in
                      [ Tuple (sanitizeName name) { fullName, fArgs: fArgsGo, fRet: fRetGo, arity: Array.length args } ]
                  Nothing ->
                    let fullName = "Call_" <> sanitizeName name
                    in [ Tuple (sanitizeName name) { fullName, fArgs: [], fRet: TypeValue, arity: 0 } ]
            )
            group.bindings
      )
      tcoBindings

    Tuple decls helpers = unsafePerformEffect do
      let d = Array.concatMap
            ( \group ->
                let
                  recVars = if group.recursive then map (\(Tuple (Ident name) _) -> sanitizeName name) group.bindings else []
                  
                  processBindingGroup :: Array (Tuple Ident TcoExpr) -> Boolean -> Array GoDecl
                  processBindingGroup binds isRec =
                    let
                      mutRecBinds = traverse (\(Tuple (Ident name) val) -> map (\abs -> { ident: sanitizeName name, args: abs.args, body: abs.body, fvs: abs.fvs }) (extractUncurriedAbs val)) binds
                    in
                      case mutRecBinds of
                        Just fns ->
                          let
                            loopCtxs = map (\fn -> { ident: fn.ident, params: fn.args, loopParams: map (\p -> p <> "_loop") fn.args }) fns

                            fnWrapperStmts = map
                              ( \fn ->
                                let
                                  pTypes = paramTypes fn.body
                                  paramsWithTypes = map (\idStr -> 
                                      let 
                                        t = fromMaybe Any (Map.lookup idStr pTypes)
                                      in Tuple idStr (exprTypeToGoType t)
                                    ) fn.args

                                  newBound = foldl (\acc (Tuple idStr goType) -> Map.insert idStr { name: idStr, goType } acc) Map.empty paramsWithTypes
                                  
                                  isSelfRecursiveLoop = group.recursive && Array.length group.bindings == 1
                                  currentLoopCtx = if isSelfRecursiveLoop then [ { ident: fn.ident, params: map fst paramsWithTypes, loopParams: map (\p -> fst p <> "_loop") paramsWithTypes, goTypes: map snd paramsWithTypes } ] else []
                                  resBodyMut = translateExprImpl_ helpersRef 0 modNameStr recVars moduleArities newBound Nothing currentLoopCtx isSelfRecursiveLoop false 0 fn.body
                                  goName = fn.ident
                                  loopParams = map (\(Tuple idStr _) -> idStr <> "_loop") paramsWithTypes
                                  initVars = Array.concatMap (\(Tuple p goT) -> [ GoRaw ("var " <> p <> " " <> goTypeToStr goT <> " = " <> p <> "_loop"), GoRaw ("_ = " <> p) ]) paramsWithTypes
                                  
                                  arity = Array.length fn.args
                                  goParams = String.joinWith ", " (map (\(Tuple p goT) -> p <> "_loop " <> goTypeToStr goT) paramsWithTypes)
                                  
                                  funcExpr = if arity >= 2 && arity <= 10 then
                                    let
                                      bodyStmts = initVars <> flattenStmts resBodyMut.stmts <> [ GoReturn resBodyMut.expr ]
                                      funcBody = if isSelfRecursiveLoop then GoFor goName bodyStmts else GoBlock bodyStmts
                                    in unsafePerformEffect do
                                      let callFuncDecl = "func Call_" <> goName <> "(" <> goParams <> ") " <> goTypeToStr resBodyMut.exprType <> " {\n" <> printGoExpr funcBody <> "\n}"
                                      Ref.modify_ (\r -> r { rawDecls = Array.snoc r.rawDecls callFuncDecl }) helpersRef
                                      let wrapperParams = map (\(Tuple p _) -> p <> "_box") paramsWithTypes
                                      let callExpr = GoCall (GoVar ("Call_" <> goName)) (map (\(Tuple p goT) -> unboxGoExpr (GoVar (p <> "_box")) TypeValue goT) paramsWithTypes)
                                      let boxedRes = boxGoExpr callExpr resBodyMut.exprType
                                      let wrapperFunc = GoRaw ("func(" <> String.joinWith ", " (map (\p -> p <> " gopurs_runtime.Value") wrapperParams) <> ") gopurs_runtime.Value {\nreturn " <> printGoExpr boxedRes <> "\n}")
                                      pure $ GoRaw ("gopurs_runtime.Func" <> show arity <> "(" <> printGoExpr wrapperFunc <> ")")
                                  else
                                    let
                                      bodyStmts = initVars <> flattenStmts resBodyMut.stmts <> [ GoReturn (boxGoExpr resBodyMut.expr resBodyMut.exprType) ]
                                      funcBody = if isSelfRecursiveLoop then GoFor goName bodyStmts else GoBlock bodyStmts
                                      iife = GoRaw ("func() gopurs_runtime.Value {\n" <> printGoExpr funcBody <> "\n}()")
                                    in
                                      Array.foldr (\(Tuple p goT) acc -> GoCall (GoSelector (GoVar "gopurs_runtime") "Func") [ GoRaw ("func(" <> p <> "_loop " <> goTypeToStr goT <> ") gopurs_runtime.Value {\nreturn " <> printGoExpr acc <> "\n}") ]) iife paramsWithTypes
                                in
                                  { identifier: goName, expression: funcExpr }
                              )
                              fns
                          in
                            fnWrapperStmts
                        Nothing ->
                          Array.concatMap
                            ( \(Tuple (Ident name) expr) ->
                                let
                                  res = translateExprImpl_ helpersRef 0 modNameStr recVars moduleArities Map.empty (Just (sanitizeName name)) [] false false 0 expr
                                in
                                  [ { identifier: sanitizeName name, expression: wrapInStmts [] res.stmts (boxGoExpr res.expr res.exprType) } ]
                            )
                            binds
                in
                  if group.recursive then
                    processBindingGroup group.bindings true
                  else
                    Array.concatMap (\b -> processBindingGroup [b] false) group.bindings
            )
            tcoBindings
      h <- Ref.read helpersRef
      pure (Tuple d h)

    allDeclsAst = decls <> helpers.decls
    declsStr = String.joinWith "\\n" (map printGoDeclVar allDeclsAst) <> "\\n" <> String.joinWith "\\n" helpers.rawDecls

    parts = String.split (Pattern "pkg_") declsStr
    usedPkgNames = Array.nub $ Array.mapMaybe
      ( \part ->
          let
            subParts = String.split (Pattern ".") part
          in
            Array.head subParts
      )
      (fromMaybe [] (Array.tail parts))

    goImports = Array.nub $
      (if Array.length allDeclsAst > 0 || Array.length (Array.fromFoldable mod.foreign) > 0 then [ "gopurs/output/gopurs_runtime" ] else [])
        <> (if Array.length allDeclsAst > 0 then [ "sync" ] else [])
        <> Array.mapMaybe
          ( \pkg ->
              if pkg /= modNameStr && pkg /= "Prim" && not (String.indexOf (Pattern "Prim_") pkg == Just 0) then Just ("gopurs/output/" <> String.replaceAll (Pattern "_") (Replacement ".") pkg)
              else Nothing
          )
          usedPkgNames

    goFile =
      { packageName: modNameStr
      , imports: goImports
      , decls: allDeclsAst
      , rawDecls: helpers.rawDecls
      , foreigns: map (\(Ident name) -> { pursName: sanitizeName name, goName: "_Gopurs_" <> capitalize (sanitizeName name) }) (Array.fromFoldable mod.foreign)
      }
  in
    printGoFile goFile


isEffectNode :: forall a. BackendSyntax a -> Boolean
isEffectNode = case _ of
  EffectBind _ _ _ _ -> true
  EffectPure _ -> true
  EffectDefer _ -> false
  PrimEffect _ -> true
  UncurriedEffectApp _ _ -> true
  _ -> false

unwrapTcoExpr :: TcoExpr -> BackendSyntax TcoExpr
unwrapTcoExpr (TcoExpr _ syn) = case syn of
  Typed _ inner -> unwrapTcoExpr inner
  _ -> syn

executeIfOpaque :: forall a. BackendSyntax a -> GoExpr -> GoExpr

executeIfOpaque expr goExpr =
  if isEffectNode expr then goExpr
  else GoCall (GoSelector (GoVar "gopurs_runtime") "Apply") [ goExpr, GoRaw "gopurs_runtime.Value{}" ]


translateExprImpl :: Ref { decls :: Array GoDecl, rawDecls :: Array String } -> Int -> String -> Array String -> Map String { fullName :: String, fArgs :: Array GoType, fRet :: GoType, arity :: Int } -> Map String { name :: String, goType :: GoType } -> Maybe String -> Array { ident :: String, params :: Array String, loopParams :: Array String, goTypes :: Array GoType } -> Boolean -> Int -> TcoExpr -> { stmts :: StmtTree, expr :: GoExpr, exprType :: GoType, nextId :: Int }
translateExprImpl helpersRef depth modNameStr recVars moduleArities bound tcoIdent loopCtx isTail nextId tcoExpr =
  translateExprImpl_ helpersRef depth modNameStr recVars moduleArities bound tcoIdent loopCtx isTail false nextId tcoExpr

translateExprImpl_ :: Ref { decls :: Array GoDecl, rawDecls :: Array String } -> Int -> String -> Array String -> Map String { fullName :: String, fArgs :: Array GoType, fRet :: GoType, arity :: Int } -> Map String { name :: String, goType :: GoType } -> Maybe String -> Array { ident :: String, params :: Array String, loopParams :: Array String, goTypes :: Array GoType } -> Boolean -> Boolean -> Int -> TcoExpr -> { stmts :: StmtTree, expr :: GoExpr, exprType :: GoType, nextId :: Int }
translateExprImpl_ helpersRef depth modNameStr recVars moduleArities bound tcoIdent loopCtx isTail inEffectBlock nextId tcoExpr@(TcoExpr tcoAnalysis expr) =
  let
    isEff = isEffectNode expr
  in
    if isEff && not inEffectBlock then
      let
        res = translateExprImpl_ helpersRef depth modNameStr recVars moduleArities bound tcoIdent loopCtx false true nextId tcoExpr
        funcExpr = GoRaw ("gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {\n" <> printGoExpr (GoBlock (flattenStmts res.stmts <> [ GoReturn (boxGoExpr res.expr res.exprType) ])) <> "\n})")
      in
        { stmts: StmtEmpty, expr: funcExpr, exprType: TypeValue, nextId: res.nextId }
    else
  let
    liftIfNeeded mkNodeThunk =
      if depth > 10 then unsafePerformEffect do
        let fvsSet = freeVars tcoExpr
        let fvs = Array.fromFoldable fvsSet
        let helperName = "__helper_" <> show nextId
        let newNextId = nextId + 1
        let res = translateExprImpl_ helpersRef 0 modNameStr recVars moduleArities bound Nothing [] false inEffectBlock newNextId tcoExpr

        let
          helperExpr =
            if Array.length fvs == 0 then GoFunc "_" (wrapInStmts [] res.stmts res.expr)
            else Array.foldr (\fv accFunc -> GoFunc fv accFunc) (wrapInStmts [] res.stmts res.expr) fvs

        Ref.modify_ (\r -> r { decls = Array.snoc r.decls { identifier: helperName, expression: helperExpr } }) helpersRef

        let
          callExpr =
            if Array.length fvs == 0 then GoCall (GoSelector (GoVar "gopurs_runtime") "Apply") [ GoCall (GoVar ("Get_" <> helperName)) [], GoRaw "gopurs_runtime.Int(0)" ]
            else Array.foldl (\accCall fv -> GoCall (GoSelector (GoVar "gopurs_runtime") "Apply") [ accCall, GoVar fv ]) (GoCall (GoVar ("Get_" <> helperName)) []) fvs

        pure { stmts: StmtEmpty, expr: callExpr, exprType: TypeValue, nextId: res.nextId }
      else mkNodeThunk unit
  in
    case expr of
      Typed type_ a ->
        let
          res = translateExprImpl_ helpersRef depth modNameStr recVars moduleArities bound tcoIdent loopCtx isTail inEffectBlock nextId a
          expectedGoType = case type_ of
            Int -> TypeInt64
            Number -> TypeFloat64
            String -> TypeString
            Char -> TypeString
            Boolean -> TypeBool
            _ -> TypeValue
        in
          if res.exprType == expectedGoType then res
          else if expectedGoType == TypeValue then res { expr = boxGoExpr res.expr res.exprType, exprType = TypeValue }
          else res { expr = unboxGoExpr res.expr res.exprType expectedGoType, exprType = expectedGoType }
      Var (Qualified mbMn (Ident i)) ->
        let
          safeName = sanitizeName i
        in
          case mbMn of
            Just mn ->
              let
                modStr = unwrap mn
                modPkg = String.replaceAll (Pattern ".") (Replacement "_") modStr
              in
                if modPkg == modNameStr then { stmts: StmtEmpty, expr: GoCall (GoVar ("Get_" <> safeName)) [], exprType: TypeValue, nextId } else { stmts: StmtEmpty, expr: GoCall (GoSelector (GoVar ("pkg_" <> modPkg)) ("Get_" <> safeName)) [], exprType: TypeValue, nextId }
            Nothing -> { stmts: StmtEmpty, expr: GoCall (GoVar ("Get_" <> safeName)) [], exprType: TypeValue, nextId }

      Local mbIdent lvl ->
        let
          v = fromMaybe { name: localId mbIdent lvl, goType: TypeValue } (Map.lookup (localId mbIdent lvl) bound)
        in
          { stmts: StmtEmpty, expr: GoVar v.name, exprType: v.goType, nextId }

      Lit (LitString s) -> { stmts: StmtEmpty, expr: GoString s, exprType: TypeString, nextId }
      Lit (LitInt i) -> { stmts: StmtEmpty, expr: GoInt i, exprType: TypeInt64, nextId }
      Lit (LitNumber n) -> { stmts: StmtEmpty, expr: GoRaw (show n), exprType: TypeFloat64, nextId }
      Lit (LitBoolean b) -> { stmts: StmtEmpty, expr: GoRaw (if b then "true" else "false"), exprType: TypeBool, nextId }
      Lit (LitChar c) -> { stmts: StmtEmpty, expr: GoString (SCU.singleton c), exprType: TypeString, nextId }

      Lit (LitArray xs) ->
        let
          accXs = foldl
            ( \acc val ->
                let
                  resVal = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars moduleArities bound Nothing [] false false acc.nextId val
                in
                  { stmts: acc.stmts <> resVal.stmts, exprs: Array.snoc acc.exprs (boxGoExpr resVal.expr resVal.exprType), exprType: TypeValue, nextId: resVal.nextId }
            )
            { stmts: StmtEmpty, exprs: [], exprType: TypeValue, nextId }
            xs
        in
          { stmts: accXs.stmts, expr: GoCall (GoSelector (GoVar "gopurs_runtime") "Array") [ GoRaw ("[]gopurs_runtime.Value{" <> String.joinWith ", " (map printGoExpr accXs.exprs) <> "}") ], exprType: TypeValue, nextId: accXs.nextId }

      Lit (LitRecord props) ->
        let
          accProps = foldl
            ( \acc (Prop key val) ->
                let
                  resVal = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars moduleArities bound Nothing [] false false acc.nextId val
                in
                  { stmts: acc.stmts <> resVal.stmts, exprs: Array.snoc acc.exprs (Tuple key (boxGoExpr resVal.expr resVal.exprType)), exprType: TypeValue, nextId: resVal.nextId }
            )
            { stmts: StmtEmpty, exprs: [], exprType: TypeValue, nextId }
            props
        in
          { stmts: accProps.stmts, expr: GoRecordDict accProps.exprs, exprType: TypeValue, nextId: accProps.nextId }

      App fn args ->
        let
          argsArr = toArray args
          Tuple flatFn flatArgs = flattenApp tcoExpr

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
                  _ = unsafePerformEffect (if name == "deepTailRec" then Debug.trace ("isTailCallTo deepTailRec: fullName=" <> fullName <> " loopCtx.idents=" <> String.joinWith "," (map _.ident loopCtx)) \_ -> pure unit else pure unit)
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
                        argRes = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars moduleArities bound Nothing [] false false acc.nextId arg
                      in
                        { stmts: acc.stmts <> argRes.stmts, exprs: Array.snoc acc.exprs argRes.expr, exprTypes: Array.snoc acc.exprTypes argRes.exprType, nextId: argRes.nextId }
                  )
                  { stmts: StmtEmpty, exprs: [], exprTypes: [], nextId }
                  flatArgs
                targetCtx = fromMaybe { ident: "", params: [], loopParams: [], goTypes: [] } (Array.index loopCtx index)
                assigns = Array.mapWithIndex (\i paramName ->
                    let
                      argExpr = fromMaybe (GoRaw "nil") (Array.index accFinal.exprs i)
                      argType = fromMaybe TypeValue (Array.index accFinal.exprTypes i)
                      expectedType = fromMaybe TypeValue (Array.index targetCtx.goTypes i)
                    in GoMutate paramName (unboxGoExpr argExpr argType expectedType)
                  ) targetCtx.loopParams
              in
                { stmts: accFinal.stmts <> foldMap StmtLeaf assigns <> StmtLeaf (GoContinue targetCtx.ident), expr: GoRaw "gopurs_runtime.Value{}", exprType: TypeValue, nextId: accFinal.nextId }

            Nothing ->
              let
                flattenFuncType :: Array ExprType -> ExprType -> { fArgs :: Array ExprType, fRet :: ExprType }
                flattenFuncType acc (Func args ret) = flattenFuncType (acc <> args) ret
                flattenFuncType acc ret = { fArgs: acc, fRet: ret }

                getFuncType :: BackendSyntax TcoExpr -> Maybe { fArgs :: Array ExprType, fRet :: ExprType }
                getFuncType (Typed (Func a r) _) = Just (flattenFuncType a r)
                getFuncType (Typed _ inner) = getFuncType (unwrapExpr inner)
                getFuncType _ = Nothing

                getVar :: BackendSyntax TcoExpr -> Maybe { mbMod :: Maybe ModuleName, name :: String }
                getVar (Typed _ inner) = getVar (unwrapTcoExpr inner)
                getVar (Var (Qualified mbMod (Ident name))) = Just { mbMod, name }
                getVar (Local (Just (Ident name)) _) = Just { mbMod: Just (ModuleName modNameStr), name }
                getVar (Local Nothing _) = Just { mbMod: Just (ModuleName modNameStr), name: "Local_Nothing" }
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

                mbDirectCall = case getVar (unwrapTcoExpr flatFn) of
                  Just { mbMod, name } ->
                    let
                      isLocal = map (String.replaceAll (Pattern ".") (Replacement "_") <<< unwrap) mbMod == Just modNameStr || mbMod == Nothing
                    in if isLocal then
                      let
                        fromModuleArities = Map.lookup name moduleArities
                        fromTypeSig = case getFuncType (unwrapExpr flatFn) of
                          Just { fArgs, fRet } ->
                            Just { fullName: "Call_" <> sanitizeName name, fArgs: map exprTypeToGoType fArgs, fRet: exprTypeToGoType fRet, arity: Array.length fArgs }
                          Nothing -> Nothing
                        
                        entry = case fromTypeSig of
                          Just e -> Just e
                          Nothing -> fromModuleArities
                      in
                        case entry of
                          Just e ->
                            if Array.length flatArgs >= e.arity && e.arity >= 2 && e.arity <= 10 then Just e else Nothing
                          Nothing -> Nothing
                    else Nothing
                  Nothing -> Nothing
              in
                case mbDirectCall of
                  Just { fullName, fArgs, fRet, arity } ->
                    let
                      accArgs = foldl
                        ( \acc arg ->
                            let
                              argRes = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars moduleArities bound Nothing [] false false acc.nextId arg
                            in
                              { stmts: acc.stmts <> argRes.stmts, exprs: Array.snoc acc.exprs argRes.expr, exprTypes: Array.snoc acc.exprTypes argRes.exprType, nextId: argRes.nextId }
                        )
                        { stmts: StmtEmpty, exprs: [], exprTypes: [], nextId }
                        flatArgs
                        
                      accArgsArity = Array.take arity accArgs.exprs
                      accArgsRemaining = Array.drop arity accArgs.exprs
                      
                      callArgs = Array.mapWithIndex (\i argExprValue ->
                          let
                            expectedType = fromMaybe TypeValue (Array.index fArgs i)
                            actualType = fromMaybe TypeValue (Array.index accArgs.exprTypes i)
                          in unboxGoExpr argExprValue actualType expectedType
                        ) accArgsArity
                        
                      callExpr = GoCall (GoVar fullName) callArgs
                      
                      finalExpr = if Array.length accArgsRemaining == 0 then
                        callExpr
                      else
                        let
                          len = Array.length accArgsRemaining
                          goFuncName = if len >= 2 && len <= 10 then "Apply" <> show len else "Apply"
                        in
                          GoCall (GoSelector (GoVar "gopurs_runtime") goFuncName) (Array.cons callExpr accArgsRemaining)
                    in
                      { stmts: accArgs.stmts, expr: finalExpr, exprType: TypeValue, nextId: accArgs.nextId }
                  Nothing ->
                    let
                      resFn = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars moduleArities bound Nothing [] false false nextId flatFn
                      accArgs = foldl
                        ( \acc arg ->
                            let
                              argRes = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars moduleArities bound Nothing [] false false acc.nextId arg
                            in
                              { stmts: acc.stmts <> argRes.stmts, exprs: Array.snoc acc.exprs (boxGoExpr argRes.expr argRes.exprType), exprType: TypeValue, nextId: argRes.nextId }
                        )
                        { stmts: resFn.stmts, exprs: [], exprType: TypeValue, nextId: resFn.nextId }
                        flatArgs

                      buildApp :: GoExpr -> Array GoExpr -> GoExpr
                      buildApp fExpr argExprs =
                        let len = Array.length argExprs
                        in
                          if len == 0 then fExpr
                          else if len == 1 then GoCall (GoSelector (GoVar "gopurs_runtime") "Apply") [ fExpr, fromMaybe (GoRaw "nil") (Array.index argExprs 0) ]
                          else if len >= 2 && len <= 10 then
                            GoCall (GoSelector (GoVar "gopurs_runtime") ("Apply" <> show len)) (Array.cons fExpr argExprs)
                          else
                            let chunk = Array.take 10 argExprs
                                rest = Array.drop 10 argExprs
                            in buildApp (buildApp fExpr chunk) rest
                      len = Array.length accArgs.exprs
                      finalExpr = buildApp (boxGoExpr resFn.expr resFn.exprType) accArgs.exprs
                    in
                      { stmts: accArgs.stmts, expr: finalExpr, exprType: TypeValue, nextId: accArgs.nextId }

      Abs args body ->
        let
          params = map (\(Tuple mbI lvl) -> localId mbI lvl) (toArray args)
          resBody = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars moduleArities bound Nothing [] isTail false nextId body
          
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
                let goParams = String.joinWith ", " (map (\p -> p <> " gopurs_runtime.Value") ps)
                in GoCall (GoSelector (GoVar "gopurs_runtime") ("Func" <> show len)) [ GoRaw ("func(" <> goParams <> ") gopurs_runtime.Value {\n" <> bodyStr <> "\n}") ]
              else
                let chunk = Array.take 5 ps
                    rest = Array.drop 5 ps
                in buildFunc chunk (buildFunc rest innerExpr)
                
          funcExpr = buildFunc params (GoBlock (flattenStmts resBody.stmts <> [ GoReturn (boxGoExpr resBody.expr resBody.exprType) ]))
        in
          { stmts: StmtEmpty, expr: funcExpr, exprType: TypeValue, nextId: resBody.nextId }

      UncurriedApp fn args ->
        let
          resFn = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars moduleArities bound Nothing [] false false nextId fn
          accArgs = foldl
            ( \acc arg ->
                let
                  argRes = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars moduleArities bound Nothing [] false false acc.nextId arg
                in
                  { stmts: acc.stmts <> argRes.stmts, exprs: Array.snoc acc.exprs (boxGoExpr argRes.expr argRes.exprType), exprType: TypeValue, nextId: argRes.nextId }
            )
            { stmts: resFn.stmts, exprs: [], exprType: TypeValue, nextId: resFn.nextId }
            args
        in
          let
            len = Array.length args
            goFuncName = if len >= 2 && len <= 10 then "UncurriedApp" <> show len else "UncurriedApp"
          in
            { stmts: accArgs.stmts, expr: GoCall (GoSelector (GoVar "gopurs_runtime") goFuncName) (Array.cons (boxGoExpr resFn.expr resFn.exprType) accArgs.exprs), exprType: TypeValue, nextId: accArgs.nextId }

      UncurriedAbs args body -> liftIfNeeded \_ ->
        let
          pTypes = paramTypes body
          paramsWithTypes = map (\(Tuple mbI lvl) -> 
              let 
                idStr = localId mbI lvl 
                t = fromMaybe Any (Map.lookup idStr pTypes)
              in Tuple idStr (exprTypeToGoType t)
            ) args

          newBound = foldl (\acc (Tuple idStr goType) -> Map.insert idStr { name: idStr, goType } acc) bound paramsWithTypes
          
          goParams = String.joinWith ", " (map (\(Tuple p goT) -> p <> " " <> goTypeToStr goT) paramsWithTypes)
          resBody = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars moduleArities newBound Nothing [] isTail false nextId body
          arity = Array.length args
        in if arity >= 2 && arity <= 10 then
          case tcoIdent of
            Just topName ->
              let
                callFuncDecl = "func Call_" <> topName <> "(" <> goParams <> ") gopurs_runtime.Value {\n" <> printGoExpr (GoBlock (flattenStmts resBody.stmts <> [ GoReturn (boxGoExpr resBody.expr resBody.exprType) ])) <> "\n}"
                funcExpr = unsafePerformEffect do
                  Ref.modify_ (\r -> r { rawDecls = Array.snoc r.rawDecls callFuncDecl }) helpersRef
                  pure $ GoRaw ("gopurs_runtime.Func" <> show arity <> "(Call_" <> topName <> ")")
              in { stmts: StmtEmpty, expr: funcExpr, exprType: TypeValue, nextId: resBody.nextId }
            Nothing ->
              let
                funcExpr = GoRaw ("gopurs_runtime.Func" <> show arity <> "(func(" <> goParams <> ") gopurs_runtime.Value {\n" <> printGoExpr (GoBlock (flattenStmts resBody.stmts <> [ GoReturn (boxGoExpr resBody.expr resBody.exprType) ])) <> "\n})")
              in { stmts: StmtEmpty, expr: funcExpr, exprType: TypeValue, nextId: resBody.nextId }
        else
          let
            params = map fst paramsWithTypes
            makeCurried [] = resBody.expr
            makeCurried [p] = GoFunc p (GoBlock (flattenStmts resBody.stmts <> [ GoReturn (boxGoExpr resBody.expr resBody.exprType) ]))
            makeCurried ps = case Array.uncons ps of
              Just { head: p, tail: rest } -> GoFunc p (makeCurried rest)
              Nothing -> resBody.expr
          in { stmts: StmtEmpty, expr: makeCurried params, exprType: TypeValue, nextId: resBody.nextId }

      UncurriedEffectApp fn args ->
        let
          resFn = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars moduleArities bound Nothing [] false false nextId fn
          accArgs = foldl
            ( \acc arg ->
                let
                  argRes = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars moduleArities bound Nothing [] false false acc.nextId arg
                in
                  { stmts: acc.stmts <> argRes.stmts, exprs: Array.snoc acc.exprs (boxGoExpr argRes.expr argRes.exprType), exprType: TypeValue, nextId: argRes.nextId }
            )
            { stmts: resFn.stmts, exprs: [], exprType: TypeValue, nextId: resFn.nextId }
            args
        in
          let
            len = Array.length args
            goFuncName = if len >= 2 && len <= 5 then "UncurriedApp" <> show len else "UncurriedApp"
          in
            { stmts: accArgs.stmts, expr: GoCall (GoSelector (GoVar "gopurs_runtime") goFuncName) (Array.cons (boxGoExpr resFn.expr resFn.exprType) accArgs.exprs), exprType: TypeValue, nextId: accArgs.nextId }

      UncurriedEffectAbs args body -> liftIfNeeded \_ ->
        let
          params = map (\(Tuple mbI lvl) -> localId mbI lvl) args
          goParams = String.joinWith ", " (map (\p -> p <> " gopurs_runtime.Value") params)
          resBody = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars moduleArities bound Nothing [] isTail false nextId body
          arity = Array.length args
        in if arity >= 2 && arity <= 5 then
          let
            funcExpr = GoRaw ("gopurs_runtime.Func" <> show arity <> "(func(" <> goParams <> ") gopurs_runtime.Value {\n" <> printGoExpr (GoBlock (flattenStmts resBody.stmts <> [ GoReturn (boxGoExpr resBody.expr resBody.exprType) ])) <> "\n})")
          in { stmts: StmtEmpty, expr: funcExpr, exprType: TypeValue, nextId: resBody.nextId }
        else
          let
            makeCurried [] = resBody.expr
            makeCurried [p] = GoFunc p (GoBlock (flattenStmts resBody.stmts <> [ GoReturn (boxGoExpr resBody.expr resBody.exprType) ]))
            makeCurried ps = case Array.uncons ps of
              Just { head: p, tail: rest } -> GoFunc p (makeCurried rest)
              Nothing -> resBody.expr
          in { stmts: StmtEmpty, expr: makeCurried params, exprType: TypeValue, nextId: resBody.nextId }

      EffectBind mbIdent lvl binding body ->
        let
          originalName = localId mbIdent lvl
          name = originalName <> "_" <> show nextId
          newBound = Map.insert originalName { name, goType: TypeValue } bound
          resBinding = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars moduleArities bound Nothing [] false true (nextId + 1) binding
          resBody = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars moduleArities newBound Nothing loopCtx isTail true resBinding.nextId body
          bindingExpr = executeIfOpaque (unwrapTcoExpr binding) (boxGoExpr resBinding.expr resBinding.exprType)
          bodyExpr = executeIfOpaque (unwrapTcoExpr body) resBody.expr
        in
          { stmts: resBinding.stmts <> StmtLeaf (GoAssign name bindingExpr) <> resBody.stmts, expr: bodyExpr, exprType: resBody.exprType, nextId: resBody.nextId }

      EffectPure binding ->
        translateExprImpl_ helpersRef (depth + 1) modNameStr recVars moduleArities bound Nothing [] false false nextId binding

      EffectDefer binding ->
        let
          resBinding = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars moduleArities bound Nothing [] false true nextId binding
          funcExpr = GoRaw ("gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {\n" <> printGoExpr (GoBlock (flattenStmts resBinding.stmts <> [ GoReturn (boxGoExpr resBinding.expr resBinding.exprType) ])) <> "\n})")
        in
          { stmts: StmtEmpty, expr: funcExpr, exprType: TypeValue, nextId: resBinding.nextId }

      Let mbIdent lvl binding body ->
        let
          originalName = localId mbIdent lvl
          name = originalName <> "_" <> show nextId
          resBinding = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars moduleArities bound Nothing [] false false (nextId + 1) binding
          newBound = Map.insert originalName { name, goType: resBinding.exprType } bound
          resBody = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars moduleArities newBound Nothing loopCtx isTail false resBinding.nextId body
        in
          { stmts: resBinding.stmts <> StmtLeaf (GoAssign name resBinding.expr) <> resBody.stmts, expr: resBody.expr, exprType: resBody.exprType, nextId: resBody.nextId }

      LetRec lvl bindings body ->
        let
          allocRes = foldl
            ( \acc (Tuple (Ident ident) _) ->
                let
                  oldName = localId (Just (Ident ident)) lvl
                  newName = oldName <> "_" <> show acc.nextId
                in
                  { newBound: Map.insert oldName { name: newName, goType: TypeValue } acc.newBound, newNames: Array.snoc acc.newNames { oldName, newName }, exprType: TypeValue, nextId: acc.nextId + 1 }
            )
            { newBound: bound, newNames: [], exprType: TypeValue, nextId }
            (toArray bindings)

          combinedRecVars = recVars <> map (\(Tuple (Ident i) _) -> sanitizeName i) (toArray bindings)
          
          isLoop = (unwrap tcoAnalysis).role.isLoop
          mutRecBinds = if isLoop && Array.length (toArray bindings) == 1 then
              traverse (\(Tuple (Ident name) val) -> map (\abs -> { ident: sanitizeName name, args: abs.args, body: abs.body, fvs: abs.fvs }) (extractUncurriedAbs val)) (toArray bindings)
            else Nothing
        in
          case mutRecBinds of
            Just fns ->
              let
                loopCtxs = map (\fn ->
                    let
                      oldName = localId (Just (Ident fn.ident)) lvl
                      newName = (fromMaybe { name: oldName, goType: TypeValue } (Map.lookup oldName allocRes.newBound)).name
                      pTypes = paramTypes fn.body
                      paramsWithTypes = map (\idStr -> 
                          let 
                            t = fromMaybe Any (Map.lookup idStr pTypes)
                          in Tuple idStr (exprTypeToGoType t)
                        ) fn.args
                    in
                      { ident: newName, params: fn.args, loopParams: map (\p -> p <> "_loop") fn.args, goTypes: map snd paramsWithTypes }
                  ) fns
                
                combinedLoopCtx = loopCtxs <> loopCtx
                
                declStmts = map (\ctx -> GoRaw ("var " <> ctx.ident <> " gopurs_runtime.Value")) loopCtxs
                
                Tuple fnWrapperStmts nextId' = foldl
                  ( \(Tuple accStmts currNextId) fn ->
                      let
                        oldName = localId (Just (Ident fn.ident)) lvl
                        newName = (fromMaybe { name: oldName, goType: TypeValue } (Map.lookup oldName allocRes.newBound)).name
                        pTypes = paramTypes fn.body
                        paramsWithTypes = map (\idStr -> 
                            let 
                              t = fromMaybe Any (Map.lookup idStr pTypes)
                            in Tuple idStr (exprTypeToGoType t)
                          ) fn.args
                        currentLoopCtx = [ { ident: newName, params: fn.args, loopParams: map (\p -> p <> "_loop") fn.args, goTypes: map snd paramsWithTypes } ]
                        resBodyMut = translateExprImpl_ helpersRef (depth + 1) modNameStr combinedRecVars moduleArities allocRes.newBound Nothing currentLoopCtx true false currNextId fn.body
                        
                        loopParams = map (\(Tuple idStr _) -> idStr <> "_loop") paramsWithTypes
                        initVars = Array.concatMap (\(Tuple p goT) -> [ GoRaw ("var " <> p <> " " <> goTypeToStr goT <> " = " <> p <> "_loop"), GoRaw ("_ = " <> p) ]) paramsWithTypes
                        
                        funcBody = GoFor newName (initVars <> flattenStmts resBodyMut.stmts <> [ GoReturn (boxGoExpr resBodyMut.expr resBodyMut.exprType) ])
                        iife = GoRaw ("func() gopurs_runtime.Value {\n" <> printGoExpr funcBody <> "\n}()")
                        funcExpr = Array.foldr (\(Tuple p goT) acc -> GoCall (GoSelector (GoVar "gopurs_runtime") "Func") [ GoRaw ("func(" <> p <> "_loop " <> goTypeToStr goT <> ") gopurs_runtime.Value {\nreturn " <> printGoExpr acc <> "\n}") ]) iife paramsWithTypes
                      in
                        Tuple (Array.snoc accStmts (GoMutate newName funcExpr)) resBodyMut.nextId
                  )
                  (Tuple [] allocRes.nextId)
                  fns
                
                resBodyOuter = translateExprImpl_ helpersRef (depth + 1) modNameStr combinedRecVars moduleArities allocRes.newBound Nothing loopCtx isTail false nextId' body
              in
                { stmts: foldMap StmtLeaf declStmts <> foldMap StmtLeaf fnWrapperStmts <> resBodyOuter.stmts, expr: resBodyOuter.expr, exprType: TypeValue, nextId: resBodyOuter.nextId }
            
            Nothing ->
              let
                accBindings = foldl
                  ( \acc (Tuple (Tuple (Ident ident) val) alloc) ->
                      let
                        res = translateExprImpl_ helpersRef (depth + 1) modNameStr combinedRecVars moduleArities allocRes.newBound (Just alloc.newName) [] false false acc.nextId val
                      in
                        { stmts: acc.stmts <> res.stmts, exprs: Array.snoc acc.exprs { key: alloc.newName, value: boxGoExpr res.expr res.exprType }, exprType: TypeValue, nextId: res.nextId }
                  )
                  { stmts: StmtEmpty, exprs: [], exprType: TypeValue, nextId: allocRes.nextId }
                  (Array.zip (toArray bindings) allocRes.newNames)

                declStmts = map (\b -> GoRaw ("var " <> b.key <> " gopurs_runtime.Value\n_ = " <> b.key)) accBindings.exprs
                assignStmts = map (\b -> GoMutate b.key b.value) accBindings.exprs

                resBody = translateExprImpl_ helpersRef (depth + 1) modNameStr combinedRecVars moduleArities allocRes.newBound Nothing loopCtx isTail false accBindings.nextId body
              in
                { stmts: foldMap StmtLeaf declStmts <> accBindings.stmts <> foldMap StmtLeaf assignStmts <> resBody.stmts, expr: resBody.expr, exprType: resBody.exprType, nextId: resBody.nextId }

      Accessor obj accessor ->
        let
          resObj = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars moduleArities bound Nothing [] false false nextId obj
        in
          case accessor of
            GetProp prop -> { stmts: resObj.stmts, expr: GoRecordAccess (boxGoExpr resObj.expr resObj.exprType) prop, exprType: TypeValue, nextId: resObj.nextId }
            GetIndex idx -> { stmts: resObj.stmts, expr: GoCall (GoSelector (GoVar "gopurs_runtime") "ArrayAccess") [ (boxGoExpr resObj.expr resObj.exprType), GoInt idx ], exprType: TypeValue, nextId: resObj.nextId }
            GetCtorField (Qualified mbMod _) _ _ (Ident ctorName) _ idx ->
              let
                structName = getStructName modNameStr mbMod ctorName
              in
                { stmts: resObj.stmts, expr: GoConstructorAccess (boxGoExpr resObj.expr resObj.exprType) structName idx, exprType: TypeValue, nextId: resObj.nextId }

      Update obj props ->
        let
          resObj = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars moduleArities bound Nothing [] false false nextId obj
          accProps = foldl
            ( \acc (Prop key val) ->
                let
                  resVal = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars moduleArities bound Nothing [] false false acc.nextId val
                in
                  { stmts: acc.stmts <> resVal.stmts, exprs: Array.snoc acc.exprs (Tuple key (boxGoExpr resVal.expr resVal.exprType)), exprType: TypeValue, nextId: resVal.nextId }
            )
            { stmts: StmtEmpty, exprs: [], exprType: TypeValue, nextId: resObj.nextId }
            props
        in
          { stmts: resObj.stmts <> accProps.stmts, expr: GoRecordUpdateDict (boxGoExpr resObj.expr resObj.exprType) accProps.exprs, exprType: TypeValue, nextId: accProps.nextId }

      CtorDef _ _ (Ident name) fields ->
        let
          structName = getStructName modNameStr Nothing name
          baseStructName = getBaseStructName modNameStr Nothing name
          funcExpr = Array.foldr (\f inner -> GoCall (GoSelector (GoVar "gopurs_runtime") "Func") [ GoRaw ("func(" <> sanitizeName f <> " gopurs_runtime.Value) gopurs_runtime.Value {\nreturn " <> printGoExpr inner <> "\n}") ]) (GoConstructor (hashString baseStructName) structName (map (\f -> GoVar (sanitizeName f)) fields)) fields
        in
          { stmts: StmtEmpty, expr: funcExpr, exprType: TypeValue, nextId }

      CtorSaturated (Qualified mbMod _) _ _ (Ident name) props ->
        let
          structName = getStructName modNameStr mbMod name
          baseStructName = getBaseStructName modNameStr mbMod name
          accProps = foldl
            ( \acc (Tuple _ val) ->
                let
                  resVal = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars moduleArities bound Nothing [] false false acc.nextId val
                in
                  { stmts: acc.stmts <> resVal.stmts, exprs: Array.snoc acc.exprs (boxGoExpr resVal.expr resVal.exprType), exprType: TypeValue, nextId: resVal.nextId }
            )
            { stmts: StmtEmpty, exprs: [], exprType: TypeValue, nextId }
            props
        in
          { stmts: accProps.stmts, expr: GoConstructor (hashString baseStructName) structName accProps.exprs, exprType: TypeValue, nextId: accProps.nextId }

      Fail msg ->
        { stmts: StmtEmpty, expr: GoRaw ("func() gopurs_runtime.Value { panic(" <> printGoExpr (GoString msg) <> ") }()"), exprType: TypeValue, nextId }

      Branch branches def ->
        let
          resDef = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars moduleArities bound Nothing loopCtx isTail false nextId def
          tmpVar = "__t" <> show resDef.nextId
          declTmp = StmtLeaf (GoRaw ("var " <> tmpVar <> " gopurs_runtime.Value"))
          labelName = "end_branch_" <> show resDef.nextId

          buildIfs = foldl
            ( \acc (Pair condExpr bodyExpr) ->
                let
                  resCond = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars moduleArities bound Nothing [] false false acc.nextId condExpr
                  resBody = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars moduleArities bound Nothing loopCtx isTail false resCond.nextId bodyExpr
                  goIf = GoIfElse (unboxGoExpr resCond.expr resCond.exprType TypeBool) (flattenStmts resBody.stmts <> [ GoMutate tmpVar (boxGoExpr resBody.expr resBody.exprType), GoRaw ("goto " <> labelName) ]) []
                in
                  { stmts: acc.stmts <> StmtLeaf (GoRaw "{") <> resCond.stmts <> StmtLeaf goIf <> StmtLeaf (GoRaw "}"), exprType: TypeValue, nextId: resBody.nextId }
            )
            { stmts: StmtEmpty, exprType: TypeValue, nextId: resDef.nextId + 1 }
            (toArray branches)
        in
          { stmts: declTmp <> buildIfs.stmts <> StmtLeaf (GoRaw "{") <> resDef.stmts <> StmtLeaf (GoMutate tmpVar (boxGoExpr resDef.expr resDef.exprType)) <> StmtLeaf (GoRaw "}") <> StmtLeaf (GoRaw (labelName <> ":")), expr: GoVar tmpVar, exprType: TypeValue, nextId: buildIfs.nextId }

      PrimOp op -> case op of
        Op1 op1 e ->
          let
            resE = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars moduleArities bound Nothing [] false false nextId e
            goOp = case op1 of
              OpBooleanNot -> { expr: GoBinOp "!=" (unboxGoExpr resE.expr resE.exprType TypeBool) (GoRaw "true"), exprType: TypeBool }
              OpIntNegate -> { expr: GoBinOp "-" (GoInt 0) (unboxGoExpr resE.expr resE.exprType TypeInt64), exprType: TypeInt64 }
              OpIntBitNot -> { expr: GoBinOp "^" (GoRaw "^0") (unboxGoExpr resE.expr resE.exprType TypeInt64), exprType: TypeInt64 }
              OpNumberNegate -> { expr: GoCall (GoSelector (GoVar "gopurs_runtime") "FloatNeg") [ resE.expr ], exprType: TypeValue }
              OpIsTag (Qualified mbMod (Ident tag)) ->
                let
                  structName = getStructName modNameStr mbMod tag
                  baseStructName = getBaseStructName modNameStr mbMod tag
                  hashStr = hashString baseStructName
                  exprStr = "(" <> printGoExpr (boxGoExpr resE.expr resE.exprType) <> ".Type == 9 && " <> printGoExpr (boxGoExpr resE.expr resE.exprType) <> ".IntVal == " <> hashStr <> ")"
                in
                  { expr: GoRaw exprStr, exprType: TypeBool }
              OpArrayLength -> { expr: GoCall (GoSelector (GoVar "gopurs_runtime") "Int") [ GoCall (GoVar "int64") [ GoCall (GoSelector (GoVar "gopurs_runtime") "ArrayLength") [ resE.expr ] ] ], exprType: TypeValue }
          in
            { stmts: resE.stmts, expr: goOp.expr, exprType: goOp.exprType, nextId: resE.nextId }
        Op2 op2 e1 e2 ->
          let
            res1 = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars moduleArities bound Nothing [] false false nextId e1
            res2 = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars moduleArities bound Nothing [] false false res1.nextId e2
            goOp = case op2 of
              OpArrayIndex -> { expr: GoCall (GoSelector (GoVar "gopurs_runtime") "ArrayAccess") [ res1.expr, GoCall (GoVar "int") [ unboxGoExpr res2.expr res2.exprType TypeInt64 ] ], exprType: TypeValue }
              OpIntNum OpAdd -> { expr: GoBinOp "+" (unboxGoExpr res1.expr res1.exprType TypeInt64) (unboxGoExpr res2.expr res2.exprType TypeInt64), exprType: TypeInt64 }
              OpIntNum OpSubtract -> { expr: GoBinOp "-" (unboxGoExpr res1.expr res1.exprType TypeInt64) (unboxGoExpr res2.expr res2.exprType TypeInt64), exprType: TypeInt64 }
              OpIntNum OpMultiply -> { expr: GoBinOp "*" (unboxGoExpr res1.expr res1.exprType TypeInt64) (unboxGoExpr res2.expr res2.exprType TypeInt64), exprType: TypeInt64 }
              OpIntNum OpDivide -> { expr: GoBinOp "/" (unboxGoExpr res1.expr res1.exprType TypeInt64) (unboxGoExpr res2.expr res2.exprType TypeInt64), exprType: TypeInt64 }
              OpIntBitAnd -> { expr: GoBinOp "&" (unboxGoExpr res1.expr res1.exprType TypeInt64) (unboxGoExpr res2.expr res2.exprType TypeInt64), exprType: TypeInt64 }
              OpIntBitOr -> { expr: GoBinOp "|" (unboxGoExpr res1.expr res1.exprType TypeInt64) (unboxGoExpr res2.expr res2.exprType TypeInt64), exprType: TypeInt64 }
              OpIntBitXor -> { expr: GoBinOp "^" (unboxGoExpr res1.expr res1.exprType TypeInt64) (unboxGoExpr res2.expr res2.exprType TypeInt64), exprType: TypeInt64 }
              OpIntBitShiftLeft -> { expr: GoBinOp "<<" (unboxGoExpr res1.expr res1.exprType TypeInt64) (unboxGoExpr res2.expr res2.exprType TypeInt64), exprType: TypeInt64 }
              OpIntBitShiftRight -> { expr: GoBinOp ">>" (unboxGoExpr res1.expr res1.exprType TypeInt64) (unboxGoExpr res2.expr res2.exprType TypeInt64), exprType: TypeInt64 }
              OpIntBitZeroFillShiftRight -> { expr: GoCall (GoSelector (GoVar "gopurs_runtime") "Zshr") [ res1.expr, res2.expr ], exprType: TypeValue }
              OpIntOrd OpEq -> { expr: GoBinOp "==" (unboxGoExpr res1.expr res1.exprType TypeInt64) (unboxGoExpr res2.expr res2.exprType TypeInt64), exprType: TypeBool }
              OpIntOrd OpNotEq -> { expr: GoBinOp "!=" (unboxGoExpr res1.expr res1.exprType TypeInt64) (unboxGoExpr res2.expr res2.exprType TypeInt64), exprType: TypeBool }
              OpIntOrd OpLt -> { expr: GoBinOp "<" (unboxGoExpr res1.expr res1.exprType TypeInt64) (unboxGoExpr res2.expr res2.exprType TypeInt64), exprType: TypeBool }
              OpIntOrd OpLte -> { expr: GoBinOp "<=" (unboxGoExpr res1.expr res1.exprType TypeInt64) (unboxGoExpr res2.expr res2.exprType TypeInt64), exprType: TypeBool }
              OpIntOrd OpGt -> { expr: GoBinOp ">" (unboxGoExpr res1.expr res1.exprType TypeInt64) (unboxGoExpr res2.expr res2.exprType TypeInt64), exprType: TypeBool }
              OpIntOrd OpGte -> { expr: GoBinOp ">=" (unboxGoExpr res1.expr res1.exprType TypeInt64) (unboxGoExpr res2.expr res2.exprType TypeInt64), exprType: TypeBool }
              OpNumberNum OpAdd -> { expr: GoBinOp "+" (unboxGoExpr res1.expr res1.exprType TypeFloat64) (unboxGoExpr res2.expr res2.exprType TypeFloat64), exprType: TypeFloat64 }
              OpNumberNum OpSubtract -> { expr: GoBinOp "-" (unboxGoExpr res1.expr res1.exprType TypeFloat64) (unboxGoExpr res2.expr res2.exprType TypeFloat64), exprType: TypeFloat64 }
              OpNumberNum OpMultiply -> { expr: GoBinOp "*" (unboxGoExpr res1.expr res1.exprType TypeFloat64) (unboxGoExpr res2.expr res2.exprType TypeFloat64), exprType: TypeFloat64 }
              OpNumberNum OpDivide -> { expr: GoBinOp "/" (unboxGoExpr res1.expr res1.exprType TypeFloat64) (unboxGoExpr res2.expr res2.exprType TypeFloat64), exprType: TypeFloat64 }
              OpNumberOrd OpEq -> { expr: GoBinOp "==" (unboxGoExpr res1.expr res1.exprType TypeFloat64) (unboxGoExpr res2.expr res2.exprType TypeFloat64), exprType: TypeBool }
              OpNumberOrd OpNotEq -> { expr: GoBinOp "!=" (unboxGoExpr res1.expr res1.exprType TypeFloat64) (unboxGoExpr res2.expr res2.exprType TypeFloat64), exprType: TypeBool }
              OpNumberOrd OpLt -> { expr: GoBinOp "<" (unboxGoExpr res1.expr res1.exprType TypeFloat64) (unboxGoExpr res2.expr res2.exprType TypeFloat64), exprType: TypeBool }
              OpNumberOrd OpLte -> { expr: GoBinOp "<=" (unboxGoExpr res1.expr res1.exprType TypeFloat64) (unboxGoExpr res2.expr res2.exprType TypeFloat64), exprType: TypeBool }
              OpNumberOrd OpGt -> { expr: GoBinOp ">" (unboxGoExpr res1.expr res1.exprType TypeFloat64) (unboxGoExpr res2.expr res2.exprType TypeFloat64), exprType: TypeBool }
              OpNumberOrd OpGte -> { expr: GoBinOp ">=" (unboxGoExpr res1.expr res1.exprType TypeFloat64) (unboxGoExpr res2.expr res2.exprType TypeFloat64), exprType: TypeBool }
              OpStringAppend -> { expr: GoBinOp "+" (unboxGoExpr res1.expr res1.exprType TypeString) (unboxGoExpr res2.expr res2.exprType TypeString), exprType: TypeString }
              OpStringOrd OpEq -> { expr: GoBinOp "==" (unboxGoExpr res1.expr res1.exprType TypeString) (unboxGoExpr res2.expr res2.exprType TypeString), exprType: TypeBool }
              OpStringOrd OpNotEq -> { expr: GoBinOp "!=" (unboxGoExpr res1.expr res1.exprType TypeString) (unboxGoExpr res2.expr res2.exprType TypeString), exprType: TypeBool }
              OpStringOrd OpLt -> { expr: GoBinOp "<" (unboxGoExpr res1.expr res1.exprType TypeString) (unboxGoExpr res2.expr res2.exprType TypeString), exprType: TypeBool }
              OpStringOrd OpLte -> { expr: GoBinOp "<=" (unboxGoExpr res1.expr res1.exprType TypeString) (unboxGoExpr res2.expr res2.exprType TypeString), exprType: TypeBool }
              OpStringOrd OpGt -> { expr: GoBinOp ">" (unboxGoExpr res1.expr res1.exprType TypeString) (unboxGoExpr res2.expr res2.exprType TypeString), exprType: TypeBool }
              OpStringOrd OpGte -> { expr: GoBinOp ">=" (unboxGoExpr res1.expr res1.exprType TypeString) (unboxGoExpr res2.expr res2.exprType TypeString), exprType: TypeBool }
              OpCharOrd OpEq -> { expr: GoBinOp "==" (unboxGoExpr res1.expr res1.exprType TypeString) (unboxGoExpr res2.expr res2.exprType TypeString), exprType: TypeBool }
              OpCharOrd OpNotEq -> { expr: GoBinOp "!=" (unboxGoExpr res1.expr res1.exprType TypeString) (unboxGoExpr res2.expr res2.exprType TypeString), exprType: TypeBool }
              OpCharOrd OpLt -> { expr: GoBinOp "<" (unboxGoExpr res1.expr res1.exprType TypeString) (unboxGoExpr res2.expr res2.exprType TypeString), exprType: TypeBool }
              OpCharOrd OpLte -> { expr: GoBinOp "<=" (unboxGoExpr res1.expr res1.exprType TypeString) (unboxGoExpr res2.expr res2.exprType TypeString), exprType: TypeBool }
              OpCharOrd OpGt -> { expr: GoBinOp ">" (unboxGoExpr res1.expr res1.exprType TypeString) (unboxGoExpr res2.expr res2.exprType TypeString), exprType: TypeBool }
              OpCharOrd OpGte -> { expr: GoBinOp ">=" (unboxGoExpr res1.expr res1.exprType TypeString) (unboxGoExpr res2.expr res2.exprType TypeString), exprType: TypeBool }
              OpBooleanOrd OpEq -> { expr: GoBinOp "==" (unboxGoExpr res1.expr res1.exprType TypeBool) (unboxGoExpr res2.expr res2.exprType TypeBool), exprType: TypeBool }
              OpBooleanOrd OpNotEq -> { expr: GoBinOp "!=" (unboxGoExpr res1.expr res1.exprType TypeBool) (unboxGoExpr res2.expr res2.exprType TypeBool), exprType: TypeBool }
              OpBooleanOrd OpLt -> { expr: GoBinOp "<" (unboxGoExpr res1.expr res1.exprType TypeBool) (unboxGoExpr res2.expr res2.exprType TypeBool), exprType: TypeBool }
              OpBooleanOrd OpLte -> { expr: GoBinOp "<=" (unboxGoExpr res1.expr res1.exprType TypeBool) (unboxGoExpr res2.expr res2.exprType TypeBool), exprType: TypeBool }
              OpBooleanOrd OpGt -> { expr: GoBinOp ">" (unboxGoExpr res1.expr res1.exprType TypeBool) (unboxGoExpr res2.expr res2.exprType TypeBool), exprType: TypeBool }
              OpBooleanOrd OpGte -> { expr: GoBinOp ">=" (unboxGoExpr res1.expr res1.exprType TypeBool) (unboxGoExpr res2.expr res2.exprType TypeBool), exprType: TypeBool }
              OpBooleanAnd -> { expr: GoBinOp "&&" (unboxGoExpr res1.expr res1.exprType TypeBool) (unboxGoExpr res2.expr res2.exprType TypeBool), exprType: TypeBool }
              OpBooleanOr -> { expr: GoBinOp "||" (unboxGoExpr res1.expr res1.exprType TypeBool) (unboxGoExpr res2.expr res2.exprType TypeBool), exprType: TypeBool }
          in
            { stmts: res1.stmts <> res2.stmts, expr: goOp.expr, exprType: goOp.exprType, nextId: res2.nextId }

      PrimEffect eff -> case eff of
        EffectRefNew a ->
          let
            resA = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars moduleArities bound Nothing [] false false nextId a
            refIdent = "__local_ref_" <> show resA.nextId
            declStmt = GoAssign refIdent (boxGoExpr resA.expr resA.exprType)
          in
            { stmts: resA.stmts <> StmtLeaf declStmt
            , expr: GoRaw ("gopurs_runtime.Any(&" <> refIdent <> ")")
            , exprType: TypeValue, nextId: resA.nextId + 1
            }
        EffectRefRead a ->
          let
            resA = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars moduleArities bound Nothing [] false false nextId a
          in
            { stmts: resA.stmts
            , expr: GoRaw ("*(" <> printGoExpr resA.expr <> ".PtrVal().(*gopurs_runtime.Value))")
            , exprType: TypeValue, nextId: resA.nextId
            }
        EffectRefWrite ref val ->
          let
            resRef = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars moduleArities bound Nothing [] false false nextId ref
            resVal = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars moduleArities bound Nothing [] false false resRef.nextId val
            writeStmt = GoRaw ("*(" <> printGoExpr resRef.expr <> ".PtrVal().(*gopurs_runtime.Value)) = " <> printGoExpr (boxGoExpr resVal.expr resVal.exprType))
          in
            { stmts: resRef.stmts <> resVal.stmts <> StmtLeaf writeStmt
            , expr: resVal.expr
            , exprType: TypeValue, nextId: resVal.nextId
            }

      _ -> { stmts: StmtEmpty, expr: GoVar "gopurs_runtime.Value{}", exprType: TypeValue, nextId }

