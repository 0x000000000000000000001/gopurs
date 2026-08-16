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
import Partial.Unsafe (unsafePartial)

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

import PureScript.Backend.Optimizer.Monomorphize (InstantiationMap)
import PureScript.Backend.Optimizer.Monomorphize as Monomorphize
import PureScript.Backend.Optimizer.Substitute (unify, substituteExprType, mapTcoExprTypes, substituteAst)
import Gopurs.GoAst (GoFile, GoDecl, GoExpr(..), GoType(..), goTypeToStr)


import Gopurs.Printer (printGoFile, printGoExpr, printGoDeclVar)
import PureScript.Backend.Optimizer.Codegen.Tco as Tco
import PureScript.Backend.Optimizer.Codegen.Tco (TcoExpr(..), tcoAnalysisOf)
import PureScript.Backend.Optimizer.FreeVars (freeVars, localId, paramTypes, sanitizeName)
import Node.Path as Path
import Node.FS.Sync as FS
import Data.Tuple (Tuple(..), fst)
import PureScript.Backend.Optimizer.FfiSupport (hashString)
import Gopurs.FfiTypes (TypeNode(..), FfiDecl)
import Data.Maybe (fromMaybe)

coerceGoExpr :: GoExpr -> GoType -> GoType -> GoExpr
coerceGoExpr expr from to | from == to = expr
coerceGoExpr expr (TypeStructPointer b1 f1 s1 a1) (TypeStructPointer b2 f2 s2 a2) | b1 == b2 && s1 == s2 && a1 == a2 = expr
coerceGoExpr expr from TypeValue = boxGoExpr expr from
coerceGoExpr expr TypeValue to = unboxGoExpr expr TypeValue to
coerceGoExpr expr from to = unboxGoExpr (boxGoExpr expr from) TypeValue to

boxGoExpr :: GoExpr -> GoType -> GoExpr
boxGoExpr expr TypeValue = expr
boxGoExpr expr TypeInt64 = GoCall (GoSelector (GoVar "gopurs_runtime") "Int") [ expr ]
boxGoExpr expr TypeFloat64 = GoCall (GoSelector (GoVar "gopurs_runtime") "Float") [ expr ]
boxGoExpr expr TypeString = GoCall (GoSelector (GoVar "gopurs_runtime") "Str") [ expr ]
boxGoExpr expr TypeBool = GoCall (GoSelector (GoVar "gopurs_runtime") "Bool") [ expr ]
boxGoExpr expr (TypeStructPointer baseStructName _ _ _) = GoRaw ("gopurs_runtime.Value{Type: 9, IntVal: " <> hashString baseStructName <> ", UnsafePtr: unsafe.Pointer(" <> printGoExpr expr <> ")}")
boxGoExpr expr (TypeRecord _) = expr
boxGoExpr expr (TypeInterface _) = expr
boxGoExpr expr (TypeNativeArray TypeValue) = GoCall (GoSelector (GoVar "gopurs_runtime") "Array") [ expr ]
boxGoExpr expr (TypeNativeArray inner) = GoRaw ("func() gopurs_runtime.Value {\n\t\t\t\t\tarr := " <> printGoExpr expr <> "\n\t\t\t\t\tboxed := make([]gopurs_runtime.Value, len(arr))\n\t\t\t\t\tfor i, v := range arr { boxed[i] = " <> printGoExpr (boxGoExpr (GoVar "v") inner) <> " }\n\t\t\t\t\treturn gopurs_runtime.Array(boxed)\n\t\t\t\t}()")
boxGoExpr expr TypeUint32 = GoRaw ("gopurs_runtime.Value{Type: 9, IntVal: int64(" <> printGoExpr expr <> "), UnsafePtr: nil}")
boxGoExpr expr (TypeGenericParam _) = expr
boxGoExpr expr (TypeFunc _ _) = expr

mangleType :: Map.Map String { ctorName :: String, arity :: Int } -> Set.Set String -> Set.Set String -> String -> ExprType -> String
mangleType ptrPaths enumAdts elidedCtors modNameStr t = 
  let
    typeStr = goTypeToStr (exprTypeToGoType ptrPaths enumAdts elidedCtors modNameStr t)
    typeStrNoPkg = String.replaceAll (Pattern "pkg_") (Replacement "") typeStr
    typeStrSafe = String.replaceAll (Pattern ".") (Replacement "_") typeStrNoPkg
    typeStrSafe2 = String.replaceAll (Pattern "[]") (Replacement "arr") typeStrSafe
    typeStrSafe3 = String.replaceAll (Pattern "*") (Replacement "ptr") typeStrSafe2
    typeStrSafe4 = String.replaceAll (Pattern "[") (Replacement "_") typeStrSafe3
    typeStrSafe5 = String.replaceAll (Pattern "]") (Replacement "_") typeStrSafe4
    typeStrSafe6 = String.replaceAll (Pattern ",") (Replacement "_") typeStrSafe5
    cleanType = String.replaceAll (Pattern " ") (Replacement "_") typeStrSafe6
  in hashString (Monomorphize.mangleType t)

exprTypeToGoType :: Map.Map String { ctorName :: String, arity :: Int } -> Set.Set String -> Set.Set String -> String -> ExprType -> GoType
exprTypeToGoType _ _ _ _ Int = TypeInt64
exprTypeToGoType _ _ _ _ Number = TypeFloat64
exprTypeToGoType _ _ _ _ String = TypeString
exprTypeToGoType _ _ _ _ Char = TypeString
exprTypeToGoType _ _ _ _ Boolean = TypeBool
exprTypeToGoType ptrPaths enumAdts elided modNameStr (Array ty) = TypeNativeArray (exprTypeToGoType ptrPaths enumAdts elided modNameStr ty)
exprTypeToGoType ptrPaths enumAdts elided modNameStr (Record (Row fields _)) = TypeRecord (map (\(Tuple k v) -> Tuple k (exprTypeToGoType ptrPaths enumAdts elided modNameStr v)) (Array.sortBy (comparing \(Tuple k _) -> k) fields))
exprTypeToGoType ptrPaths enumAdts elided modNameStr (Record _) = TypeValue
exprTypeToGoType ptrPaths enumAdts elided modNameStr (ADT fullName path args) =
  let
    ctorName = fromMaybe "" (Array.last path)
    pkgNameStr = String.replaceAll (Pattern ".") (Replacement "_") (String.joinWith "." (Array.slice 0 (Array.length path - 1) path))
    monoStructName = "Constructor_" <> pkgNameStr <> "_" <> sanitizeName ctorName
  in
  if Set.member monoStructName elided then TypeValue else
  if Set.member fullName enumAdts then TypeUint32
  else case Map.lookup fullName ptrPaths of
  Just info -> 
    let 
      baseStructName = "Data_" <> pkgNameStr <> "_" <> sanitizeName info.ctorName
      monoStructName' = "Constructor_" <> pkgNameStr <> "_" <> sanitizeName info.ctorName
      typeArgsMapped = map (exprTypeToGoType ptrPaths enumAdts elided modNameStr) args
      typeArgsMappedTruncated = Array.take info.arity typeArgsMapped
      paddedTypeArgs = typeArgsMappedTruncated <> Array.replicate (info.arity - Array.length typeArgsMappedTruncated) TypeValue
      typeArgsStr = ""
    in TypeStructPointer baseStructName fullName (monoStructName' <> typeArgsStr) paddedTypeArgs
  Nothing -> TypeValue
exprTypeToGoType ptrPaths enumAdts elided modNameStr (TypeApp fn arg) =
  let
    unwrapTypeApp :: ExprType -> Array ExprType -> Tuple ExprType (Array ExprType)
    unwrapTypeApp (TypeApp f a) acc = unwrapTypeApp f (a <> acc)
    unwrapTypeApp other acc = Tuple other acc
  in case unwrapTypeApp (TypeApp fn arg) [] of
    Tuple (ADT fullName path args) allArgs -> exprTypeToGoType ptrPaths enumAdts elided modNameStr (ADT fullName path (args <> allArgs))
    _ -> TypeValue
exprTypeToGoType _ _ _ _ (TypeVar v) = TypeValue
exprTypeToGoType _ _ _ _ _ = TypeValue

exprTypeToGenericGoType :: Map.Map String { ctorName :: String, arity :: Int } -> Set.Set String -> Set.Set String -> Array String -> String -> ExprType -> GoType
exprTypeToGenericGoType ptrPaths enumAdts elidedCtors typeVars modNameStr (Record (Row fields _)) = TypeRecord (map (\(Tuple k v) -> Tuple k (exprTypeToGenericGoType ptrPaths enumAdts elidedCtors typeVars modNameStr v)) (Array.sortBy (comparing \(Tuple k _) -> k) fields))
exprTypeToGenericGoType _ _ _ _ _ (Record _) = TypeValue
exprTypeToGenericGoType ptrPaths enumAdts elidedCtors typeVars modNameStr (TypeApp fn arg) =
  let
    unwrapTypeApp :: ExprType -> Array ExprType -> Tuple ExprType (Array ExprType)
    unwrapTypeApp (TypeApp f a) acc = unwrapTypeApp f (a <> acc)
    unwrapTypeApp other acc = Tuple other acc
  in case unwrapTypeApp (TypeApp fn arg) [] of
    Tuple (ADT fullName path args) allArgs -> exprTypeToGenericGoType ptrPaths enumAdts elidedCtors typeVars modNameStr (ADT fullName path (args <> allArgs))
    _ -> TypeValue
exprTypeToGenericGoType _ _ _ typeVars _ (TypeVar v) | Array.elem v typeVars = TypeValue
exprTypeToGenericGoType ptrPaths enumAdts elidedCtors _ modNameStr ty = exprTypeToGoType ptrPaths enumAdts elidedCtors modNameStr ty

structFieldGoType :: Map.Map String { ctorName :: String, arity :: Int } -> Set.Set String -> Set.Set String -> Array String -> String -> ExprType -> GoType
structFieldGoType ptrPaths enumAdts elidedCtors typeVars modStr ty = 
  case exprTypeToGenericGoType ptrPaths enumAdts elidedCtors typeVars modStr ty of
    TypeInterface _ -> TypeValue
    other -> other

instantiateGenericGoType :: Map.Map String GoType -> GoType -> GoType
instantiateGenericGoType env (TypeGenericParam v) = fromMaybe TypeValue (Map.lookup v env)
instantiateGenericGoType env (TypeRecord fields) = TypeRecord (map (\(Tuple k v) -> Tuple k (instantiateGenericGoType env v)) fields)
instantiateGenericGoType env (TypeNativeArray ty) = TypeNativeArray (instantiateGenericGoType env ty)
instantiateGenericGoType env (TypeStructPointer base key full typeArgs) = TypeStructPointer base key full (map (instantiateGenericGoType env) typeArgs)
instantiateGenericGoType env (TypeFunc args ret) = TypeFunc (map (instantiateGenericGoType env) args) (instantiateGenericGoType env ret)
instantiateGenericGoType env t = t


unboxGoExpr :: GoExpr -> GoType -> GoType -> GoExpr
unboxGoExpr expr currentType desiredType =
  if currentType == desiredType then expr
  else if goTypeToStr currentType == goTypeToStr desiredType && String.contains (Pattern "Constructor_Test_RBTree_T") (goTypeToStr currentType) then
    let
      cArgs = case currentType of
        TypeStructPointer b1 k1 f1 args1 -> "cBase=" <> b1 <> ", cKey=" <> k1 <> ", cFull=" <> f1 <> ", cLen=" <> show (Array.length args1)
        _ -> "none"
      dArgs = case desiredType of
        TypeStructPointer b2 k2 f2 args2 -> "dBase=" <> b2 <> ", dKey=" <> k2 <> ", dFull=" <> f2 <> ", dLen=" <> show (Array.length args2)
        _ -> "none"
    in Debug.trace ("MISMATCH AGAIN: " <> cArgs <> " vs " <> dArgs <> ". Structurally equal arrays? " <> show (currentType == desiredType)) \_ ->
    unboxGoExpr (boxGoExpr expr currentType) TypeValue desiredType
  else if currentType /= TypeValue then
    unboxGoExpr (boxGoExpr expr currentType) TypeValue desiredType
  else case desiredType of
    TypeValue -> boxGoExpr expr currentType
    (TypeRecord _) -> boxGoExpr expr currentType
    TypeInt64 -> GoSelector expr "IntVal"
    TypeFloat64 -> GoCall (GoSelector expr "FloatVal") []
    TypeString -> GoCall (GoSelector expr "StrVal") []
    TypeBool -> GoBinOp "!=" (GoSelector expr "IntVal") (GoInt 0)
    TypeUint32 -> GoRaw ("uint32(" <> printGoExpr (GoSelector expr "IntVal") <> ")")
    (TypeStructPointer _ _ fullPath _) -> GoCall (GoRaw ("gopurs_runtime.CoerceToStruct[" <> fullPath <> "]")) [ expr ]
    (TypeInterface _) -> expr
    (TypeNativeArray inner) -> case currentType of
      TypeNativeArray currentInner -> 
        GoRaw ("func() " <> goTypeToStr desiredType <> " {\n\t\t\t\t\tarr := " <> printGoExpr expr <> "\n\t\t\t\t\tunboxed := make(" <> goTypeToStr desiredType <> ", len(arr))\n\t\t\t\t\tfor i, v := range arr { unboxed[i] = " <> printGoExpr (unboxGoExpr (GoVar "v") currentInner inner) <> " }\n\t\t\t\t\treturn unboxed\n\t\t\t\t}()")
      _ -> 
        GoRaw ("func() " <> goTypeToStr desiredType <> " {\n\t\t\t\t\tarr := *(*[]gopurs_runtime.Value)(" <> printGoExpr expr <> ".UnsafePtr)\n\t\t\t\t\tunboxed := make(" <> goTypeToStr desiredType <> ", len(arr))\n\t\t\t\t\tfor i, v := range arr { unboxed[i] = " <> printGoExpr (unboxGoExpr (GoVar "v") TypeValue inner) <> " }\n\t\t\t\t\treturn unboxed\n\t\t\t\t}()")
    (TypeGenericParam _) -> expr
    (TypeFunc _ _) -> expr


capitalize :: String -> String
capitalize "" = ""
capitalize s =
  let
    firstChar = String.take 1 s
  in
    if firstChar >= "a" && firstChar <= "z" then String.toUpper firstChar <> String.drop 1 s
    else if firstChar == "_" then "_" <> capitalize (String.drop 1 s)
    else s


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
    let
      thisArgs = map (\(Tuple mbI lvl) -> localId mbI lvl) args
    in case extractUncurriedAbs body of
      Just inner -> Just { args: thisArgs <> inner.args, body: inner.body, fvs: Set.union (freeVars tcoExpr) inner.fvs }
      Nothing -> Just { args: thisArgs, body, fvs: freeVars tcoExpr }
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
    UncurriedApp f args ->
      let
        Tuple f' args' = flattenApp f
      in
        Tuple f' (args' <> args)
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

globalReusedVars :: Ref.Ref (Set.Set String)
globalReusedVars = unsafePerformEffect (Ref.new Set.empty)

translate :: Set.Set String -> Set.Set String -> Map.Map String { ctorName :: String, arity :: Int } -> Set.Set String -> Map.Map String { nodeBaseStruct :: String, nodeCtor :: String } -> Set.Set ExprType -> Set.Set String -> Map.Map String { vars :: Array String, fields :: Array ExprType } -> Map.Map String ExprType -> InstantiationMap -> Map.Map String String -> Map.Map String { vars :: Array String, fields :: Array { name :: String, "type" :: ExprType } } -> Array (Array String) -> BackendModule -> String
translate enumAdts enumCtors pointerAdtPaths pointerAdtNodes pointerAdtLeaves adtTypes elidedCtors ctorTypes globalTypes rawInstantiations classDeclsMap classDeclsFields importsArray mod =

  let
    modNameStrOrig = unwrap mod.name
    modNameStr = String.replaceAll (Pattern ".") (Replacement "_") modNameStrOrig
    
    flatImportsSet = Set.fromFoldable (map (String.joinWith ".") importsArray)
    
    isSafeType :: ExprType -> Boolean
    isSafeType t =
      let
        check m = m == "" || m == modNameStrOrig || Set.member m flatImportsSet || m == "Prim" || String.indexOf (Pattern "Prim.") m == Just 0
        
        checkSafe (ADT _ parts args) = 
          let adtFullName = String.joinWith "." parts
              modPart = case SCU.lastIndexOf (Pattern ".") adtFullName of
                          Just idx -> SCU.take idx adtFullName
                          Nothing -> ""
          in (if Map.member adtFullName pointerAdtPaths then check modPart else true) && Array.all checkSafe args
        checkSafe (Array ty) = checkSafe ty
        checkSafe (Func args ret) = Array.all checkSafe args && checkSafe ret
        checkSafe (Record row) = checkSafe row
        checkSafe (Row fields tail) = 
          let tailSafe = case tail of
                Nothing -> true
                Just ty -> checkSafe ty
          in Array.all (\(Tuple _ ty) -> checkSafe ty) fields && tailSafe
        checkSafe (TypeApp c args) = checkSafe c && Array.all checkSafe args
        checkSafe (ForAll _ body) = checkSafe body
        checkSafe (ConstrainedType constraints body) = Array.all (\(Tuple _ a) -> Array.all checkSafe a) constraints && checkSafe body
        checkSafe _ = true
      in
        checkSafe t
    
    instantiations = Map.mapMaybe (\typeMap ->
        let set = Set.fromFoldable (Map.keys typeMap)
            safeSet = Set.filter isSafeType set
        in if Set.isEmpty safeSet then Nothing else Just safeSet
      ) rawInstantiations


    helpersRef = unsafePerformEffect do
      let
        structDecls = Array.concatMap (\decl ->
            Array.concatMap (\ctor ->
              let
                fieldTypes = ctor.fields
                goFieldTypes = map (structFieldGoType pointerAdtPaths enumAdts elidedCtors decl.vars modNameStr) fieldTypes
                structName = "Constructor_" <> modNameStr <> "_" <> sanitizeName ctor.name
                
                typeParams = ""
                
                fieldsStr = Array.cons "Rc uint32" (Array.mapWithIndex (\i ty -> "V" <> show i <> " " <> goTypeToStr ty) goFieldTypes)
                structDecl = "type " <> structName <> typeParams <> " struct {\n\t" <> String.joinWith "\n\t" fieldsStr <> "\n}\n"
                
                fullName = unwrap mod.name <> "." <> ctor.name
                getterDecl = case Map.lookup fullName classDeclsFields of
                  Just info ->
                    let 
                      typeParamsGetter = ""
                      cases = Array.mapWithIndex (\i f -> "\t\tcase \"" <> f.name <> "\": return gopurs_runtime.Box(c.V" <> show i <> ")") info.fields
                      pkgNameStr = String.replaceAll (Pattern ".") (Replacement "_") (unwrap mod.name)
                      baseStructName = "Data_" <> pkgNameStr <> "_" <> sanitizeName ctor.name
                      hashStr = hashString baseStructName
                    in "func init() {\n\tgopurs_runtime.StructGetters[" <> hashStr <> "] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {\n\t\tc := (*" <> structName <> typeParamsGetter <> ")(ptr)\n\t\t_ = c\n\t\tswitch key {\n" <> String.joinWith "\n" cases <> "\n\t\tdefault: panic(\"Key not found in dictionary " <> structName <> ": \" + key)\n\t\t}\n\t}\n}\n"
                  Nothing -> ""
                  
              in
                if getterDecl == "" then [ structDecl ] else [ structDecl, getterDecl ]
            ) decl.constructors
          ) mod.dataDecls

      Ref.new { decls: [], rawDecls: structDecls, elidedCtors, ctorTypes, pointerAdtPaths, pointerAdtNodes, pointerAdtLeaves, enumAdts, enumCtors, globalTypes, classDeclsFields, globalId: 0 }

    Tuple _ tcoBindings = foldl
      ( \(Tuple env acc) group ->
          let
            neBindings = fromArray group.bindings

            env' = case neBindings of
              Just ne | group.recursive -> Tco.topLevelTcoEnvGroup mod.name ne <> env
              _ -> env
            tcoBinds = map (\(Tuple id val) ->
              let nameStr = unwrap id
                  _ = if modNameStrOrig == "Data.Set" && String.indexOf (Pattern "toList") nameStr == Just 0 then unsafePerformEffect (Console.log ("CodeGen received Data.Set binding: " <> nameStr)) else unit
              in Tuple id (Tco.analyze env' val)
            ) group.bindings
          in
            Tuple env' (Array.snoc acc { recursive: group.recursive, bindings: tcoBinds })
      )
      (Tuple [] [])
      mod.bindings

    getExprType :: TcoExpr -> ExprType
    getExprType (TcoExpr _ (Typed ty _)) = ty
    getExprType _ = Any -- fallback

    setTcoExprType :: ExprType -> TcoExpr -> TcoExpr
    setTcoExprType ty (TcoExpr a (Typed _ inner)) = TcoExpr a (Typed ty inner)
    setTcoExprType _ expr = expr

    tcoBindingsExpanded = tcoBindings

    unwrapFunc :: Array (Tuple Ident TcoExpr) -> Array (Tuple String { fullName :: String, fArgs :: Array GoType, fRet :: GoType, arity :: Int })
    unwrapFunc binds =
      Array.concatMap
        ( \(Tuple (Ident name) val) ->
            case extractUncurriedAbs val of
              Just { args, body, fvs } ->
                let
                  typeSig = extractFuncType val
                  fArgsGo = case typeSig of
                    Just { fArgs } -> map (exprTypeToGoType pointerAdtPaths enumAdts elidedCtors modNameStr) (Array.take (Array.length args) fArgs)
                    Nothing -> Array.replicate (Array.length args) TypeValue
                  fRetGo = case typeSig of
                    Just { fArgs, fRet } -> 
                      let ret = if Array.length args < Array.length fArgs then TypeValue else exprTypeToGoType pointerAdtPaths enumAdts elidedCtors modNameStr fRet
                      in ret
                    Nothing -> TypeValue
                  fullName = "Call_" <> modNameStr <> "_" <> sanitizeName name
                in
                  [ Tuple (sanitizeName name) { fullName, fArgs: fArgsGo, fRet: fRetGo, arity: Array.length args } ]
              Nothing ->
                let fullName = "Call_" <> modNameStr <> "_" <> sanitizeName name
                in [ Tuple (sanitizeName name) { fullName, fArgs: [], fRet: TypeValue, arity: 0 } ]
        )
        binds


    moduleArities :: Map String { fullName :: String, fArgs :: Array GoType, fRet :: GoType, arity :: Int }
    moduleArities = Map.fromFoldable $ Array.concatMap
      ( \group -> 
          if group.recursive then
            let
              mutRecBinds = traverse (\(Tuple _ val) -> extractUncurriedAbs val) group.bindings
            in case mutRecBinds of
              Just _ -> unwrapFunc group.bindings
              Nothing -> map (\(Tuple (Ident name) _) -> Tuple (sanitizeName name) { fullName: "Call_" <> modNameStr <> "_" <> sanitizeName name, fArgs: [], fRet: TypeValue, arity: 0 }) group.bindings
          else
            unwrapFunc group.bindings
      )
      tcoBindingsExpanded

    Tuple decls helpers = unsafePerformEffect do
      let d = Array.concatMap
            ( \group ->
                let
                  recVars = if group.recursive then map (\(Tuple (Ident name) _) -> sanitizeName name) group.bindings else []
                  processBindingGroup :: Array (Tuple Ident TcoExpr) -> Boolean -> Array GoDecl
                  processBindingGroup binds isRec =
                    let
                      mutRecBinds = traverse (\(Tuple (Ident name) val) -> map (\abs -> { ident: sanitizeName name, args: abs.args, body: abs.body, fvs: abs.fvs, val: val }) (extractUncurriedAbs val)) binds
                    in
                      case mutRecBinds of
                        Just fns ->
                          let
                            loopCtxs = map (\fn -> { ident: fn.ident, params: fn.args, loopParams: map (\p -> p <> "_loop") fn.args }) fns

                            fnWrapperStmts = map
                              ( \fn ->
                                let
                                  paramsWithTypes = case extractExprFuncType (getExprType fn.val) of
                                    Just { fArgs } -> Array.zipWith (\p goType -> Tuple p goType) fn.args (map (exprTypeToGoType (unsafePerformEffect (Ref.read helpersRef)).pointerAdtPaths (unsafePerformEffect (Ref.read helpersRef)).enumAdts (unsafePerformEffect (Ref.read helpersRef)).elidedCtors modNameStr) fArgs <> Array.replicate (Array.length fn.args - Array.length fArgs) TypeValue)
                                    Nothing -> map (\p -> Tuple p TypeValue) fn.args

                                  newBound = foldl (\acc (Tuple idStr goType) -> Map.insert idStr { name: idStr, goType } acc) Map.empty paramsWithTypes
                                  
                                  isSelfRecursiveLoop = group.recursive && Array.length group.bindings == 1
                                  currentLoopCtx = if isSelfRecursiveLoop then [ { ident: fn.ident, params: map fst paramsWithTypes, loopParams: map (\p -> fst p <> "_loop") paramsWithTypes, goTypes: map snd paramsWithTypes } ] else []
                                  resBodyMut = translateExprImpl_ helpersRef 0 modNameStr recVars moduleArities newBound (Just fn.ident) currentLoopCtx isSelfRecursiveLoop false 0 fn.body
                                  goName = fn.ident
                                  loopParams = map (\(Tuple idStr _) -> idStr <> "_loop") paramsWithTypes
                                  initVars = Array.concatMap (\(Tuple p goT) -> [ GoRaw ("var " <> p <> " " <> goTypeToStr goT <> " = " <> p <> "_loop"), GoRaw ("_ = " <> p) ]) paramsWithTypes
                                  
                                  expectedRetType = case Map.lookup goName moduleArities of
                                    Just { fRet } -> fRet
                                    Nothing -> TypeValue
                                  
                                  arity = Array.length fn.args
                                  goParams = String.joinWith ", " (map (\(Tuple p goT) -> p <> "_loop " <> goTypeToStr goT) paramsWithTypes)
                                  
                                  funcExpr = if arity >= 1 && arity <= 10 then
                                    let
                                      coercedExpr = coerceGoExpr resBodyMut.expr resBodyMut.exprType expectedRetType
                                      bodyStmts = initVars <> flattenStmts resBodyMut.stmts <> [ GoReturn coercedExpr ]
                                      funcBody = if isSelfRecursiveLoop then GoFor goName bodyStmts else GoBlock bodyStmts
                                    in unsafePerformEffect do
                                      let callFuncDecl = "func Call_" <> modNameStr <> "_" <> goName <> "(" <> goParams <> ") " <> goTypeToStr expectedRetType <> " {\n" <> printGoExpr funcBody <> "\n}"
                                      Ref.modify_ (\r -> r { rawDecls = Array.snoc r.rawDecls callFuncDecl }) helpersRef
                                      let wrapperParams = map (\(Tuple p _) -> p <> "_box") paramsWithTypes
                                      let callExpr = GoCall (GoVar ("Call_" <> modNameStr <> "_" <> goName)) (map (\(Tuple p goT) -> unboxGoExpr (GoVar (p <> "_box")) TypeValue goT) paramsWithTypes)
                                      let boxedRes = boxGoExpr callExpr expectedRetType
                                      let wrapperFunc = GoRaw ("func(" <> String.joinWith ", " (map (\p -> p <> " gopurs_runtime.Value") wrapperParams) <> ") gopurs_runtime.Value {\nreturn " <> printGoExpr boxedRes <> "\n}")
                                      let funcWrapperName = if arity == 1 then "gopurs_runtime.Func" else "gopurs_runtime.Func" <> show arity
                                      pure $ GoRaw (funcWrapperName <> "(" <> printGoExpr wrapperFunc <> ")")
                                  else
                                    let
                                      bodyStmts = initVars <> flattenStmts resBodyMut.stmts <> [ GoReturn (boxGoExpr resBodyMut.expr resBodyMut.exprType) ]
                                      funcBody = if isSelfRecursiveLoop then GoFor goName bodyStmts else GoBlock bodyStmts
                                      iife = GoRaw ("func() gopurs_runtime.Value {\n" <> printGoExpr funcBody <> "\n}()")
                                    in
                                      Array.foldr (\(Tuple p goT) acc -> GoCall (GoSelector (GoVar "gopurs_runtime") "Func") [ GoRaw ("func(" <> p <> "_box gopurs_runtime.Value) gopurs_runtime.Value {\nvar " <> p <> "_loop " <> goTypeToStr goT <> " = " <> printGoExpr (unboxGoExpr (GoVar (p <> "_box")) TypeValue goT) <> "\nreturn " <> printGoExpr acc <> "\n}") ]) iife paramsWithTypes
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
                                  res = translateExprImpl_ helpersRef 0 modNameStr recVars moduleArities Map.empty (Just (sanitizeName name)) [] false false 0 expr
                                in
                                  [ { identifier: modNameStr <> "_" <> sanitizeName name, expression: wrapInStmts [] res.stmts (boxGoExpr res.expr res.exprType), goType: TypeValue } ]
                            )
                            binds
                in
                  if group.recursive then
                    processBindingGroup group.bindings true
                  else
                    Array.concatMap (\b -> processBindingGroup [b] false) group.bindings
            )
            tcoBindingsExpanded
      h <- Ref.read helpersRef
      pure (Tuple d h)

    allDeclsAst = decls <> helpers.decls
    declsStr = String.joinWith "\\n" (map printGoDeclVar allDeclsAst) <> "\\n" <> String.joinWith "\\n" helpers.rawDecls

    parts = [declsStr]
    usedPkgNames = Set.toUnfoldable $ Set.fromFoldable $ Array.mapMaybe
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
      , rawDecls: helpers.rawDecls
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

unwrapTcoExpr :: TcoExpr -> BackendSyntax TcoExpr
unwrapTcoExpr (TcoExpr _ syn) = case syn of
  Typed _ inner -> unwrapTcoExpr inner
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
          let constraintTypes = map (\(Tuple qual args) ->
                let
                   qualStr = String.joinWith "." qual
                in ADT qualStr qual args
              ) constraints
          in Just (i { fArgs = constraintTypes <> i.fArgs })
        Nothing -> Nothing
    getFunc (ForAll _ innerTy) = getFunc innerTy
    getFunc _ = Nothing
  in getFunc ty

extractFuncType :: TcoExpr -> Maybe { fArgs :: Array ExprType, fRet :: ExprType }
extractFuncType (TcoExpr _ (Typed ty inner)) =
  case extractExprFuncType ty of
    Just r -> Just r
    Nothing -> extractFuncType inner
extractFuncType _ = Nothing

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

getExprTypeArity :: ExprType -> Int
getExprTypeArity (Func args ret) = Array.length args + getExprTypeArity ret
getExprTypeArity _ = 0

executeIfOpaque :: TcoExpr -> GoExpr -> GoExpr

executeIfOpaque expr goExpr =
  if isEffectNode expr then goExpr
  else GoCall (GoSelector (GoVar "gopurs_runtime") "Apply") [ goExpr, GoRaw "gopurs_runtime.Value{}" ]


translateExprImpl :: Ref { decls :: Array GoDecl, rawDecls :: Array String, elidedCtors :: Set.Set String, ctorTypes :: Map String { vars :: Array String, fields :: Array ExprType }, pointerAdtPaths :: Map String { ctorName :: String, arity :: Int }, pointerAdtNodes :: Set String, pointerAdtLeaves :: Map String { nodeBaseStruct :: String, nodeCtor :: String }, enumAdts :: Set.Set String, enumCtors :: Set.Set String, globalTypes :: Map.Map String ExprType, classDeclsFields :: Map String { vars :: Array String, fields :: Array { name :: String, "type" :: ExprType } }, globalId :: Int } -> Int -> String -> Array String -> Map String { fullName :: String, fArgs :: Array GoType, fRet :: GoType, arity :: Int } -> Map String { name :: String, goType :: GoType } -> Maybe String -> Array { ident :: String, params :: Array String, loopParams :: Array String, goTypes :: Array GoType } -> Boolean -> Int -> TcoExpr -> { stmts :: StmtTree, expr :: GoExpr, exprType :: GoType, nextId :: Int }
translateExprImpl helpersRef depth modNameStr recVars moduleArities bound tcoIdent loopCtx isTail nextId tcoExpr =
  translateExprImpl_ helpersRef depth modNameStr recVars moduleArities bound tcoIdent loopCtx isTail false nextId tcoExpr

translateExprImpl_ :: Ref { decls :: Array GoDecl, rawDecls :: Array String, elidedCtors :: Set.Set String, ctorTypes :: Map String { vars :: Array String, fields :: Array ExprType }, pointerAdtPaths :: Map String { ctorName :: String, arity :: Int }, pointerAdtNodes :: Set String, pointerAdtLeaves :: Map String { nodeBaseStruct :: String, nodeCtor :: String }, enumAdts :: Set.Set String, enumCtors :: Set.Set String, globalTypes :: Map.Map String ExprType, classDeclsFields :: Map String { vars :: Array String, fields :: Array { name :: String, "type" :: ExprType } }, globalId :: Int } -> Int -> String -> Array String -> Map String { fullName :: String, fArgs :: Array GoType, fRet :: GoType, arity :: Int } -> Map String { name :: String, goType :: GoType } -> Maybe String -> Array { ident :: String, params :: Array String, loopParams :: Array String, goTypes :: Array GoType } -> Boolean -> Boolean -> Int -> TcoExpr -> { stmts :: StmtTree, expr :: GoExpr, exprType :: GoType, nextId :: Int }
translateExprImpl_ helpersRef depth modNameStr recVars moduleArities bound tcoIdent loopCtx isTail inEffectBlock nextId tcoExpr@(TcoExpr tcoAnalysis expr) =
  let
    _ = unsafePerformEffect (if depth == 0 then Ref.write Set.empty globalReusedVars else pure unit)
    elidedCtors = (unsafePerformEffect (Ref.read helpersRef)).elidedCtors
    isEff = isEffectNode tcoExpr
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
        let gId = unsafePerformEffect do
              curr <- Ref.read helpersRef
              Ref.modify_ (\r -> r { globalId = r.globalId + 1 }) helpersRef
              pure curr.globalId
        let helperName = modNameStr <> "__helper_" <> show gId
        let newNextId = nextId + 1
        let newBound = foldl (\acc fv -> Map.insert fv { name: fv, goType: TypeValue } acc) bound fvs
        let res = translateExprImpl_ helpersRef 0 modNameStr recVars moduleArities newBound Nothing [] false inEffectBlock newNextId tcoExpr

        let
          helperExpr =
            if Array.length fvs == 0 then GoFunc "_" TypeValue TypeValue (wrapInStmts [] res.stmts res.expr)
            else Array.foldr (\fv accFunc -> GoFunc fv TypeValue TypeValue accFunc) (wrapInStmts [] res.stmts res.expr) fvs

        Ref.modify_ (\r -> r { decls = Array.snoc r.decls { identifier: helperName, expression: helperExpr, goType: TypeValue } }) helpersRef

        let
          callExpr =
            if Array.length fvs == 0 then GoCall (GoSelector (GoVar "gopurs_runtime") "Apply") [ GoCall (GoVar ("Get_" <> helperName)) [], GoRaw "gopurs_runtime.Int(0)" ]
            else Array.foldl (\accCall fv -> GoCall (GoSelector (GoVar "gopurs_runtime") "Apply") [ accCall, boxGoExpr (GoVar (fromMaybe fv (map _.name (Map.lookup fv bound)))) (fromMaybe TypeValue (map _.goType (Map.lookup fv bound))) ]) (GoCall (GoVar ("Get_" <> helperName)) []) fvs

        pure { stmts: StmtEmpty, expr: callExpr, exprType: TypeValue, nextId: res.nextId }
      else mkNodeThunk unit
  in
    case expr of
      Typed type_ a ->
        let
          expectedGoType = exprTypeToGoType (unsafePerformEffect (Ref.read helpersRef)).pointerAdtPaths (unsafePerformEffect (Ref.read helpersRef)).enumAdts (unsafePerformEffect (Ref.read helpersRef)).elidedCtors modNameStr type_
          h = unsafePerformEffect (Ref.read helpersRef)

        in case unwrapTcoExpr a, expectedGoType of
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
                      in Array.take arity mapped
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
                              let paramsWithTypes = Array.zipWith (\(Tuple mbI lvl) fArgTy -> Tuple (localId mbI lvl) (exprTypeToGoType h.pointerAdtPaths h.enumAdts h.elidedCtors modNameStr fArgTy)) (toArray args) (fArgs <> Array.replicate (Array.length (toArray args) - Array.length fArgs) Any)
                              in foldl (\b (Tuple idStr goType) -> Map.insert idStr { name: idStr, goType } b) bound paramsWithTypes
                            UncurriedAbs args _, Just { fArgs } ->
                              let paramsWithTypes = Array.zipWith (\(Tuple mbI lvl) fArgTy -> Tuple (localId mbI lvl) (exprTypeToGoType h.pointerAdtPaths h.enumAdts h.elidedCtors modNameStr fArgTy)) args (fArgs <> Array.replicate (Array.length args - Array.length fArgs) Any)
                              in foldl (\b (Tuple idStr goType) -> Map.insert idStr { name: idStr, goType } b) bound paramsWithTypes
                            _, _ -> bound
                            
                          resVal = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars moduleArities newBound Nothing [] false false acc.nextId item.val
                          coercedExpr = coerceGoExpr resVal.expr resVal.exprType expectedType
                        in
                          { stmts: acc.stmts <> resVal.stmts, exprs: Array.snoc acc.exprs coercedExpr, exprType: TypeValue, nextId: resVal.nextId }
                    )
                    { stmts: StmtEmpty, exprs: [], exprType: TypeValue, nextId }
                    sortedVals
                in
                  { stmts: accProps.stmts, expr: GoConstructor (hashString baseStructName) fullPath [] accProps.exprs, exprType: expectedGoType, nextId: accProps.nextId }
              Nothing ->
                let res = translateExprImpl_ helpersRef depth modNameStr recVars moduleArities bound tcoIdent loopCtx isTail inEffectBlock nextId a
                in case res.exprType of
                  TypeStructPointer _ _ _ _ -> res
                  _ -> 
                    if expectedGoType == TypeValue then res
                    else { stmts: res.stmts, expr: coerceGoExpr res.expr res.exprType expectedGoType, exprType: expectedGoType, nextId: res.nextId }
          _, _ ->
            let res = translateExprImpl_ helpersRef depth modNameStr recVars moduleArities bound tcoIdent loopCtx isTail inEffectBlock nextId a
            in case res.exprType of
              TypeStructPointer _ _ _ _ -> res
              _ -> 
                if expectedGoType == TypeValue then res
                else { stmts: res.stmts, expr: coerceGoExpr res.expr res.exprType expectedGoType, exprType: expectedGoType, nextId: res.nextId }
      Var (Qualified mbMn (Ident i)) ->
        let
          safeName = sanitizeName i
          h = unsafePerformEffect (Ref.read helpersRef)
          vType = case mbMn of
            Just mn ->
              let modStr = unwrap mn
              in case Map.lookup (modStr <> "." <> i) h.globalTypes of
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
                { stmts: StmtEmpty, expr: unboxGoExpr rawCall TypeValue vType, exprType: vType, nextId }
            Nothing ->
              let
                rawCall = Debug.trace ("mbMn is Nothing for safeName: " <> safeName) (\_ -> GoCall (GoVar ("Get_" <> modNameStr <> "_" <> safeName)) [])
              in
                { stmts: StmtEmpty, expr: unboxGoExpr rawCall TypeValue vType, exprType: vType, nextId }

      Local mbIdent lvl ->
        let
          v = fromMaybe { name: localId mbIdent lvl, goType: TypeValue } (Map.lookup (localId mbIdent lvl) bound)
        in
          { stmts: StmtEmpty, expr: GoVar v.name, exprType: v.goType, nextId }

      Lit (LitString s) -> { stmts: StmtEmpty, expr: GoString s, exprType: TypeString, nextId }
      Lit (LitInt i) -> { stmts: StmtEmpty, expr: GoInt i, exprType: TypeInt64, nextId }
      Lit (LitNumber n) -> 
        let 
          expr = if n == 0.0 && 1.0 / n < 0.0 then GoCall (GoSelector (GoVar "gopurs_runtime") "NegativeZero") [] else GoRaw (show n)
        in { stmts: StmtEmpty, expr, exprType: TypeFloat64, nextId }
      Lit (LitBoolean b) -> { stmts: StmtEmpty, expr: GoRaw (if b then "true" else "false"), exprType: TypeBool, nextId }
      Lit (LitChar c) -> { stmts: StmtEmpty, expr: GoString (SCU.singleton c), exprType: TypeString, nextId }

      Lit (LitArray xs) ->
        let
          accXs = foldl
            ( \acc val ->
                let
                  resVal = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars moduleArities bound Nothing [] false false acc.nextId val
                in
                  { stmts: acc.stmts <> resVal.stmts, exprs: Array.snoc acc.exprs resVal.expr, exprTypes: Array.snoc acc.exprTypes resVal.exprType, nextId: resVal.nextId }
            )
            { stmts: StmtEmpty, exprs: [], exprTypes: [], nextId }
            xs
          
          mbElemType = Array.head accXs.exprTypes
          isAllSame = Array.all (\t -> Just t == mbElemType) accXs.exprTypes
        in
          case mbElemType of
            Just elemType | isAllSame && elemType /= TypeValue ->
              let goTypeArr = TypeNativeArray elemType
              in { stmts: accXs.stmts, expr: GoRaw (goTypeToStr goTypeArr <> "{" <> String.joinWith ", " (map printGoExpr accXs.exprs) <> "}"), exprType: goTypeArr, nextId: accXs.nextId }
            _ ->
              let boxedExprs = Array.zipWith (\expr ty -> boxGoExpr expr ty) accXs.exprs accXs.exprTypes
              in { stmts: accXs.stmts, expr: GoCall (GoSelector (GoVar "gopurs_runtime") "Array") [ GoRaw ("[]gopurs_runtime.Value{" <> String.joinWith ", " (map printGoExpr boxedExprs) <> "}") ], exprType: TypeValue, nextId: accXs.nextId }

      Lit (LitRecord props) ->
        let
          sortedProps = Array.sortBy (comparing \(Prop k _) -> k) props
          
          expectedGlobalType = case tcoIdent of
            Just ident ->
              let h = unsafePerformEffect (Ref.read helpersRef)
                  ty = Map.lookup (modNameStr <> "." <> ident) h.globalTypes

              in ty
            Nothing -> Nothing
            
          recordFields = case expectedGlobalType of
            Just (Record (Row fields _)) -> Map.fromFoldable fields
            Just ty ->
              let
                unwrapTypeApp :: ExprType -> Array ExprType -> Tuple ExprType (Array ExprType)
                unwrapTypeApp (TypeApp f a) acc = unwrapTypeApp f (a <> acc)
                unwrapTypeApp other acc = Tuple other acc
                Tuple mbAdt allArgs = unwrapTypeApp ty []
              in case mbAdt of
                ADT fullName path adtArgs ->
                  let h = unsafePerformEffect (Ref.read helpersRef)

                  in case Map.lookup fullName h.classDeclsFields of
                    Just info -> 
                      let
                        finalArgs = adtArgs <> allArgs
                        instMap = Map.fromFoldable (Array.zip info.vars finalArgs)
                        instantiatedFields = map (\f -> Tuple f.name (substituteExprType instMap f."type")) info.fields
                      in Map.fromFoldable instantiatedFields
                    Nothing -> Map.empty
                _ -> Map.empty
            _ -> Map.empty
            
          accProps = foldl
            ( \acc (Prop key val) ->
                let
                  expectedExprType = fromMaybe Any (Map.lookup key recordFields)
                  newBound = case unwrapTcoExpr val, extractExprFuncType expectedExprType of
                    Abs args _, Just { fArgs } ->
                      let paramsWithTypes = Array.zipWith (\(Tuple mbI lvl) fArgTy -> Tuple (localId mbI lvl) (exprTypeToGoType (unsafePerformEffect (Ref.read helpersRef)).pointerAdtPaths (unsafePerformEffect (Ref.read helpersRef)).enumAdts (unsafePerformEffect (Ref.read helpersRef)).elidedCtors modNameStr fArgTy)) (toArray args) (fArgs <> Array.replicate (Array.length (toArray args) - Array.length fArgs) Any)
                      in foldl (\b (Tuple idStr goType) -> Map.insert idStr { name: idStr, goType } b) bound paramsWithTypes
                    UncurriedAbs args _, Just { fArgs } ->
                      let paramsWithTypes = Array.zipWith (\(Tuple mbI lvl) fArgTy -> Tuple (localId mbI lvl) (exprTypeToGoType (unsafePerformEffect (Ref.read helpersRef)).pointerAdtPaths (unsafePerformEffect (Ref.read helpersRef)).enumAdts (unsafePerformEffect (Ref.read helpersRef)).elidedCtors modNameStr fArgTy)) args (fArgs <> Array.replicate (Array.length args - Array.length fArgs) Any)
                      in foldl (\b (Tuple idStr goType) -> Map.insert idStr { name: idStr, goType } b) bound paramsWithTypes
                    _, _ -> bound
                    
                  resVal = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars moduleArities newBound Nothing [] false false acc.nextId val
                in
                  { stmts: acc.stmts <> resVal.stmts, exprs: Array.snoc acc.exprs (Tuple key (boxGoExpr resVal.expr resVal.exprType)), exprType: TypeValue, nextId: resVal.nextId }
            )
            { stmts: StmtEmpty, exprs: [], exprType: TypeValue, nextId }
            sortedProps
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
                      fromModuleArities = if isLocal then Map.lookup name moduleArities else Nothing
                      fromTypeSig = case extractFuncType flatFn of
                        Just { fArgs, fRet } ->
                          Just { fullName: "Call_" <> modPrefix <> "_" <> sanitizeName name, fArgs: map (exprTypeToGoType (unsafePerformEffect (Ref.read helpersRef)).pointerAdtPaths (unsafePerformEffect (Ref.read helpersRef)).enumAdts (unsafePerformEffect (Ref.read helpersRef)).elidedCtors modNameStr) fArgs, fRet: exprTypeToGoType (unsafePerformEffect (Ref.read helpersRef)).pointerAdtPaths (unsafePerformEffect (Ref.read helpersRef)).enumAdts (unsafePerformEffect (Ref.read helpersRef)).elidedCtors modNameStr fRet, arity: Array.length fArgs }
                        Nothing ->
                          Nothing
                      
                      entry = case fromTypeSig of
                        Just e | not isLocal -> Just e
                        _ -> fromModuleArities
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
                              argRes = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars moduleArities bound Nothing [] false false acc.nextId arg
                            in
                              { stmts: acc.stmts <> argRes.stmts, exprs: Array.snoc acc.exprs (boxGoExpr argRes.expr argRes.exprType), exprTypes: Array.snoc acc.exprTypes argRes.exprType, nextId: argRes.nextId }
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
                            iifeBody = GoBlock [
                              GoAssign arrGoName (GoCall (GoRaw "(*[]gopurs_runtime.Value)") [ GoSelector (GoVar arrValName) "UnsafePtr" ]),
                              GoAssign resGoName (GoCall (GoVar "make") [ GoRaw "[]gopurs_runtime.Value", GoCall (GoVar "len") [ GoRaw ("*" <> arrGoName) ] ]),
                              GoForRange (iName <> ", " <> vName <> " := range *" <> arrGoName) [ loopBody ],
                              GoReturn (GoCall (GoSelector (GoVar "gopurs_runtime") "Array") [ GoVar resGoName ])
                            ]
                          in GoIIFE arrValName arrExpr iifeBody
                          
                        "foldlArray" ->
                          let
                            fExpr = fromMaybe (GoRaw "nil") (Array.index accArgs.exprs 0)
                            initExpr = fromMaybe (GoRaw "nil") (Array.index accArgs.exprs 1)
                            arrExpr = fromMaybe (GoRaw "nil") (Array.index accArgs.exprs 2)
                            loopBody = GoMutate resGoName (GoCall (GoSelector (GoVar "gopurs_runtime") "Apply2") [ fExpr, GoVar resGoName, GoVar vName ])
                            iifeBody = GoBlock [
                              GoAssign resGoName initExpr,
                              GoAssign arrGoName (GoCall (GoRaw "(*[]gopurs_runtime.Value)") [ GoSelector (GoVar arrValName) "UnsafePtr" ]),
                              GoForRange ("_, " <> vName <> " := range *" <> arrGoName) [ loopBody ],
                              GoReturn (GoVar resGoName)
                            ]
                          in GoIIFE arrValName arrExpr iifeBody
                          
                        "filterImpl" ->
                          let
                            fExpr = fromMaybe (GoRaw "nil") (Array.index accArgs.exprs 0)
                            arrExpr = fromMaybe (GoRaw "nil") (Array.index accArgs.exprs 1)
                            condExpr = GoCall (GoSelector (GoVar "gopurs_runtime") "Apply") [ fExpr, GoVar vName ]
                            isTrueExpr = GoCall (GoSelector condExpr "BoolVal") []
                            loopBody = GoIfElse isTrueExpr [ GoMutate resGoName (GoCall (GoVar "append") [ GoVar resGoName, GoVar vName ]) ] []
                            iifeBody = GoBlock [
                              GoAssign arrGoName (GoCall (GoRaw "(*[]gopurs_runtime.Value)") [ GoSelector (GoVar arrValName) "UnsafePtr" ]),
                              GoAssign resGoName (GoCall (GoVar "make") [ GoRaw "[]gopurs_runtime.Value", GoRaw "0" ]),
                              GoForRange ("_, " <> vName <> " := range *" <> arrGoName) [ loopBody ],
                              GoReturn (GoCall (GoSelector (GoVar "gopurs_runtime") "Array") [ GoVar resGoName ])
                            ]
                          in GoIIFE arrValName arrExpr iifeBody
                          
                        _ -> GoRaw "nil"
                        
                      arity = if intrinsicName == "foldlArray" then 3 else 2
                      accArgsRemaining = Array.drop arity accArgs.exprs
                      accArgsRemainingTypes = Array.drop arity accArgs.exprTypes
                      accArgsRemainingBoxed = Array.zipWith (\arg t -> boxGoExpr arg t) accArgsRemaining accArgsRemainingTypes
                      
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
                            
                      finalExpr = buildApp iifeExpr accArgsRemainingBoxed
                    in
                      { stmts: accArgs.stmts, expr: finalExpr, exprType: TypeValue, nextId: accArgs.nextId }

                  Nothing ->
                    let
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
                            accArgsRemainingTypes = Array.drop arity accArgs.exprTypes
                            accArgsRemainingBoxed = Array.zipWith (\arg t -> boxGoExpr arg t) accArgsRemaining accArgsRemainingTypes
                            
                            callArgs = Array.mapWithIndex (\i argExprValue ->
                                let
                                  expectedType = fromMaybe TypeValue (Array.index fArgs i)
                                  actualType = fromMaybe TypeValue (Array.index accArgs.exprTypes i)
                                in unboxGoExpr argExprValue actualType expectedType
                              ) accArgsArity
                              
                            callExpr = GoCall (GoVar fullName) callArgs
                            finalExpr = if Array.length accArgsRemainingBoxed == 0 then callExpr else buildApp (boxGoExpr callExpr fRet) accArgsRemainingBoxed
                            finalExprType = if Array.length accArgsRemainingBoxed == 0 then fRet else TypeValue
                          in
                            { stmts: accArgs.stmts, expr: finalExpr, exprType: finalExprType, nextId: accArgs.nextId }
  
                        Nothing ->
                          let
                            resFn = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars moduleArities bound Nothing [] false false nextId flatFn
                            accArgs = foldl
                              ( \acc arg ->
                                  let
                                    argRes = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars moduleArities bound Nothing [] false false acc.nextId arg
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
                                  callArgs = Array.mapWithIndex (\i expected ->
                                      let arg = fromMaybe (GoRaw "nil") (Array.index accArgs.exprs i)
                                          actual = fromMaybe TypeValue (Array.index accArgs.exprTypes i)
                                      in unboxGoExpr arg actual expected
                                    ) fArgs
                                in GoCall resFn.expr callArgs
                              TypeFunc fArgs fRet | Array.length flatArgs > Array.length fArgs ->
                                let
                                  arity = Array.length fArgs
                                  accArgsArity = Array.take arity accArgs.exprs
                                  accArgsRemaining = Array.drop arity accArgs.exprs
                                  accArgsRemainingTypes = Array.drop arity accArgs.exprTypes
                                  accArgsRemainingBoxed = Array.zipWith (\arg t -> boxGoExpr arg t) accArgsRemaining accArgsRemainingTypes
                                  
                                  callArgs = Array.mapWithIndex (\i argExprValue ->
                                      let
                                        expectedType = fromMaybe TypeValue (Array.index fArgs i)
                                        actualType = fromMaybe TypeValue (Array.index accArgs.exprTypes i)
                                      in unboxGoExpr argExprValue actualType expectedType
                                    ) accArgsArity
                                    
                                  callExpr = GoCall resFn.expr callArgs
                                in buildApp (boxGoExpr callExpr fRet) accArgsRemainingBoxed
                              _ ->
                                let 
                                  boxedArgs = Array.zipWith (\arg actual -> boxGoExpr arg actual) accArgs.exprs accArgs.exprTypes
                                in buildApp (boxGoExpr resFn.expr resFn.exprType) boxedArgs
                          in
                            { stmts: accArgs.stmts, expr: finalExpr, exprType: finalExprType, nextId: accArgs.nextId }

      Abs args body ->
        let
          paramsWithTypes = case extractFuncType tcoExpr of
            Just { fArgs } -> Array.zipWith (\(Tuple mbI lvl) goType -> Tuple (localId mbI lvl) goType) (toArray args) (map (exprTypeToGoType (unsafePerformEffect (Ref.read helpersRef)).pointerAdtPaths (unsafePerformEffect (Ref.read helpersRef)).enumAdts (unsafePerformEffect (Ref.read helpersRef)).elidedCtors modNameStr) fArgs <> Array.replicate (Array.length (toArray args) - Array.length fArgs) TypeValue)
            Nothing -> map (\(Tuple mbI lvl) -> Tuple (localId mbI lvl) TypeValue) (toArray args)
            
          newBound = foldl (\acc (Tuple idStr goType) -> Map.insert idStr { name: idStr, goType } acc) bound paramsWithTypes
          params = map fst paramsWithTypes
          resBody = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars moduleArities newBound Nothing [] isTail false nextId body
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
                        argRes = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars moduleArities bound Nothing [] false false acc.nextId arg
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
                      
                      mbFnArityInfo = Map.lookup fnFullName moduleArities
                      
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
                          in GoMutate (resGoName <> "[" <> iName <> "]") (unboxGoExpr (GoCall (GoVar ("Call_" <> String.replaceAll (Pattern ".") (Replacement "_") fnFullName)) [ unboxGoExpr (GoVar vName) elemType expectedArgType ]) info.fRet retType)
                        _ ->
                          GoMutate (resGoName <> "[" <> iName <> "]") (unboxGoExpr (GoCall (GoSelector (GoVar "gopurs_runtime") "Apply") [ boxGoExpr fExprRaw fExprType, boxGoExpr (GoVar vName) elemType ]) TypeValue retType)
                          
                      iifeBodyStmts = [
                        arrGoAssignment,
                        GoAssign resGoName (GoCall (GoVar "make") [ GoRaw ("[]" <> goTypeToStr retType), GoCall (GoVar "len") [ GoRaw arrGoRangeTarget ] ]),
                        GoForRange (iName <> ", " <> vName <> " := range " <> arrGoRangeTarget) [ loopBody ]
                      ]
                    in { stmts: accArgs.stmts, expr: GoCall (GoFuncLit [] (Array.cons (GoAssign arrValName arrExprRaw) (Array.cons (GoMutate "_" (GoVar arrValName)) iifeBodyStmts)) (GoVar resGoName) finalRetType) [], exprType: finalRetType, nextId: accArgs.nextId }
                    
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
                      
                      mbFnArityInfo = Map.lookup fnFullName moduleArities
                      
                      elemType = case arrExprType of
                        TypeNativeArray inner -> inner
                        _ -> TypeValue
                      
                      loopBody = case mbFnArityInfo of
                        Just info | info.arity == 2 ->
                          let
                            expectedArg0 = fromMaybe TypeValue (Array.index info.fArgs 0)
                            expectedArg1 = fromMaybe TypeValue (Array.index info.fArgs 1)
                          in GoMutate resGoName (unboxGoExpr (GoCall (GoVar ("Call_" <> String.replaceAll (Pattern ".") (Replacement "_") fnFullName)) [ unboxGoExpr (GoVar resGoName) initExprType expectedArg0, unboxGoExpr (GoVar vName) elemType expectedArg1 ]) info.fRet initExprType)
                        _ ->
                          GoMutate resGoName (unboxGoExpr (GoCall (GoSelector (GoVar "gopurs_runtime") "Apply2") [ boxGoExpr fExpr fExprType, boxGoExpr (GoVar resGoName) initExprType, boxGoExpr (GoVar vName) elemType ]) TypeValue initExprType)
                          
                      arrGoAssignment = case arrExprType of
                        TypeNativeArray _ -> GoAssign arrGoName (GoVar arrValName)
                        _ -> GoAssign arrGoName (GoCall (GoRaw "(*[]gopurs_runtime.Value)") [ GoSelector (GoVar arrValName) "UnsafePtr" ])
                      
                      arrGoRangeTarget = case arrExprType of
                        TypeNativeArray _ -> arrGoName
                        _ -> "*" <> arrGoName
                        
                      iifeBody = GoBlock [
                        GoAssign resGoName initExpr,
                        arrGoAssignment,
                        GoForRange ("_, " <> vName <> " := range " <> arrGoRangeTarget) [ loopBody ],
                        GoReturn (GoVar resGoName)
                      ]
                    in { stmts: accArgs.stmts, expr: GoIIFE arrValName arrExpr iifeBody, exprType: initExprType, nextId: accArgs.nextId }
                    
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
                      
                      mbFnArityInfo = Map.lookup fnFullName moduleArities
                      
                      elemType = case arrExprType of
                        TypeNativeArray inner -> inner
                        _ -> TypeValue
                        
                      isTrueExpr = case mbFnArityInfo of
                        Just info | info.arity == 1 ->
                          let
                            expectedArgType = fromMaybe TypeValue (Array.index info.fArgs 0)
                          in GoCall (GoVar ("Call_" <> String.replaceAll (Pattern ".") (Replacement "_") fnFullName)) [ unboxGoExpr (GoVar vName) elemType expectedArgType ]
                        _ ->
                          let condExpr = GoCall (GoSelector (GoVar "gopurs_runtime") "Apply") [ boxGoExpr fExprRaw fExprType, boxGoExpr (GoVar vName) elemType ]
                          in GoCall (GoSelector condExpr "BoolVal") []
                          
                      loopBody = GoIfElse isTrueExpr [ GoMutate resGoName (GoCall (GoVar "append") [ GoVar resGoName, GoVar vName ]) ] []
                      
                      arrGoAssignment = case arrExprType of
                        TypeNativeArray _ -> GoAssign arrGoName (GoVar arrValName)
                        _ -> GoAssign arrGoName (GoCall (GoRaw "(*[]gopurs_runtime.Value)") [ GoSelector (GoVar arrValName) "UnsafePtr" ])
                      
                      arrGoRangeTarget = case arrExprType of
                        TypeNativeArray _ -> arrGoName
                        _ -> "*" <> arrGoName
                        
                      iifeBodyStmts = [
                        arrGoAssignment,
                        GoAssign resGoName (GoCall (GoVar "make") [ GoRaw ("[]" <> goTypeToStr elemType), GoRaw "0" ]),
                        GoForRange ("_, " <> vName <> " := range " <> arrGoRangeTarget) [ loopBody ]
                      ]
                    in { stmts: accArgs.stmts, expr: GoCall (GoFuncLit [] (Array.cons (GoAssign arrValName arrExprRaw) (Array.cons (GoMutate "_" (GoVar arrValName)) iifeBodyStmts)) (GoVar resGoName) (TypeNativeArray elemType)) [], exprType: TypeNativeArray elemType, nextId: accArgs.nextId }
                    
                  _ -> { stmts: accArgs.stmts, expr: GoRaw "nil", exprType: TypeValue, nextId: accArgs.nextId }
              in
                iifeExpr
            Nothing ->
              let
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
                      resFn = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars moduleArities bound Nothing [] false false nextId fn
                      accArgs = foldl
                        ( \acc arg ->
                            let
                              argRes = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars moduleArities bound Nothing [] false false acc.nextId arg
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
                            callArgs = Array.mapWithIndex (\i expected ->
                                let arg = fromMaybe (GoRaw "nil") (Array.index accArgs.exprs i)
                                    actual = fromMaybe TypeValue (Array.index accArgs.exprTypes i)
                                in unboxGoExpr arg actual expected
                              ) fArgs
                          in
                            { stmts: accArgs.stmts, expr: boxGoExpr (GoCall resFn.expr callArgs) fRet, exprType: TypeValue, nextId: accArgs.nextId }
                        _ ->
                          let
                            boxedArgs = Array.zipWith (\arg actual -> boxGoExpr arg actual) accArgs.exprs accArgs.exprTypes
                          in
                            { stmts: accArgs.stmts, expr: GoCall (GoSelector (GoVar "gopurs_runtime") goFuncName) (Array.cons (boxGoExpr resFn.expr resFn.exprType) boxedArgs), exprType: TypeValue, nextId: accArgs.nextId }

      UncurriedAbs args body -> liftIfNeeded \_ ->
        let
          paramsWithTypes = case extractExprFuncType (getExprType tcoExpr) of
            Just { fArgs } -> Array.zipWith (\(Tuple mbI lvl) goType -> Tuple (localId mbI lvl) goType) args (map (exprTypeToGoType (unsafePerformEffect (Ref.read helpersRef)).pointerAdtPaths (unsafePerformEffect (Ref.read helpersRef)).enumAdts (unsafePerformEffect (Ref.read helpersRef)).elidedCtors modNameStr) fArgs <> Array.replicate (Array.length args - Array.length fArgs) TypeValue)
            Nothing -> map (\(Tuple mbI lvl) -> Tuple (localId mbI lvl) TypeValue) args

          newBound = foldl (\acc (Tuple idStr goType) -> Map.insert idStr { name: idStr, goType } acc) bound paramsWithTypes
          
          goParams = String.joinWith ", " (map (\(Tuple p goT) -> p <> " " <> goTypeToStr goT) paramsWithTypes)
          resBody = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars moduleArities newBound Nothing [] isTail false nextId body
          arity = Array.length args
        in if arity >= 2 && arity <= 10 then
          case tcoIdent of
            Just topName ->
              let
                callFuncDecl = "func Call_" <> modNameStr <> "_" <> topName <> "(" <> goParams <> ") gopurs_runtime.Value {\n" <> printGoExpr (GoBlock (flattenStmts resBody.stmts <> [ GoReturn (boxGoExpr resBody.expr resBody.exprType) ])) <> "\n}"
                funcExpr = unsafePerformEffect do
                  Ref.modify_ (\r -> r { rawDecls = Array.snoc r.rawDecls callFuncDecl }) helpersRef
                  pure $ GoRaw ("gopurs_runtime.Func" <> show arity <> "(Call_" <> modNameStr <> "_" <> topName <> ")")
              in { stmts: StmtEmpty, expr: funcExpr, exprType: TypeValue, nextId: resBody.nextId }
            Nothing ->
              let
                funcExpr = GoRaw ("gopurs_runtime.Func" <> show arity <> "(func(" <> goParams <> ") gopurs_runtime.Value {\n" <> printGoExpr (GoBlock (flattenStmts resBody.stmts <> [ GoReturn (boxGoExpr resBody.expr resBody.exprType) ])) <> "\n})")
              in { stmts: StmtEmpty, expr: funcExpr, exprType: TypeValue, nextId: resBody.nextId }
        else
          let
            params = map fst paramsWithTypes
            makeCurried [] = resBody.expr
            makeCurried [p] = GoFunc p TypeValue TypeValue (GoBlock (flattenStmts resBody.stmts <> [ GoReturn (boxGoExpr resBody.expr resBody.exprType) ]))
            makeCurried ps = case Array.uncons ps of
              Just { head: p, tail: rest } -> GoFunc p TypeValue TypeValue (makeCurried rest)
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
                  callArgs = Array.mapWithIndex (\i expected ->
                      let arg = fromMaybe (GoRaw "nil") (Array.index accArgs.exprs i)
                          actual = fromMaybe TypeValue (Array.index accArgs.exprTypes i)
                      in unboxGoExpr arg actual expected
                    ) fArgs
                in
                  { stmts: accArgs.stmts, expr: boxGoExpr (GoCall resFn.expr callArgs) fRet, exprType: TypeValue, nextId: accArgs.nextId }
              _ ->
                let
                  boxedArgs = Array.zipWith (\arg actual -> boxGoExpr arg actual) accArgs.exprs accArgs.exprTypes
                in
                  { stmts: accArgs.stmts, expr: GoCall (GoSelector (GoVar "gopurs_runtime") goFuncName) (Array.cons (boxGoExpr resFn.expr resFn.exprType) boxedArgs), exprType: TypeValue, nextId: accArgs.nextId }

      UncurriedEffectAbs args body -> liftIfNeeded \_ ->
        let
          paramsWithTypes = case extractExprFuncType (getExprType tcoExpr) of
            Just { fArgs } -> Array.zipWith (\(Tuple mbI lvl) goType -> Tuple (localId mbI lvl) goType) args (map (exprTypeToGoType (unsafePerformEffect (Ref.read helpersRef)).pointerAdtPaths (unsafePerformEffect (Ref.read helpersRef)).enumAdts (unsafePerformEffect (Ref.read helpersRef)).elidedCtors modNameStr) fArgs <> Array.replicate (Array.length args - Array.length fArgs) TypeValue)
            Nothing -> map (\(Tuple mbI lvl) -> Tuple (localId mbI lvl) TypeValue) args
          newBound = foldl (\acc (Tuple idStr goType) -> Map.insert idStr { name: idStr, goType } acc) bound paramsWithTypes
          goParams = String.joinWith ", " (map (\(Tuple p goT) -> p <> " " <> goTypeToStr goT) paramsWithTypes)
          resBody = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars moduleArities newBound Nothing [] isTail false nextId body
          arity = Array.length args
        in if arity >= 2 && arity <= 5 then
          let
            funcExpr = GoRaw ("gopurs_runtime.Func" <> show arity <> "(func(" <> goParams <> ") gopurs_runtime.Value {\n" <> printGoExpr (GoBlock (flattenStmts resBody.stmts <> [ GoReturn (GoRaw ("gopurs_runtime.Apply(" <> printGoExpr (boxGoExpr resBody.expr resBody.exprType) <> ", gopurs_runtime.Value{})")) ])) <> "\n})")
          in { stmts: StmtEmpty, expr: funcExpr, exprType: TypeValue, nextId: resBody.nextId }
        else
          let
            params = map fst paramsWithTypes
            makeCurried [] = GoRaw ("gopurs_runtime.Apply(" <> printGoExpr (boxGoExpr resBody.expr resBody.exprType) <> ", gopurs_runtime.Value{})")
            makeCurried [p] = GoFunc p TypeValue TypeValue (GoBlock (flattenStmts resBody.stmts <> [ GoReturn (GoRaw ("gopurs_runtime.Apply(" <> printGoExpr (boxGoExpr resBody.expr resBody.exprType) <> ", gopurs_runtime.Value{})")) ]))
            makeCurried ps = case Array.uncons ps of
              Just { head: p, tail: rest } -> GoFunc p TypeValue TypeValue (makeCurried rest)
              Nothing -> resBody.expr
          in { stmts: StmtEmpty, expr: makeCurried params, exprType: TypeValue, nextId: resBody.nextId }

      EffectBind mbIdent lvl binding body ->
        let
          stripEffectDefer (TcoExpr a syn) = case unwrapTcoExpr (TcoExpr a syn) of
              EffectDefer inner -> stripEffectDefer inner
              Abs _ inner -> stripEffectDefer inner
              Let ident lvl val body -> TcoExpr a (Let ident lvl val (stripEffectDefer body))
              LetRec lvl bindings body -> TcoExpr a (LetRec lvl bindings (stripEffectDefer body))
              _ -> TcoExpr a syn
          realBinding = stripEffectDefer binding
          originalName = localId mbIdent lvl
          name = originalName <> "_" <> show nextId
          newBound = Map.insert originalName { name, goType: TypeValue } bound
          resBinding = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars moduleArities bound Nothing [] false true (nextId + 1) realBinding
          resBody = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars moduleArities newBound Nothing loopCtx isTail true resBinding.nextId body
          bindingExpr = executeIfOpaque realBinding (boxGoExpr resBinding.expr resBinding.exprType)
          bodyExpr = executeIfOpaque body resBody.expr
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
          expectedGoType = exprTypeToGoType (unsafePerformEffect (Ref.read helpersRef)).pointerAdtPaths (unsafePerformEffect (Ref.read helpersRef)).enumAdts (unsafePerformEffect (Ref.read helpersRef)).elidedCtors modNameStr (getExprType binding)

          newBound = Map.insert originalName { name, goType: expectedGoType } bound
          resBody = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars moduleArities newBound Nothing loopCtx isTail inEffectBlock resBinding.nextId body
          
          letStmt = if expectedGoType == resBinding.exprType then
                      StmtLeaf (GoAssign name resBinding.expr)
                    else
                      StmtLeaf (GoRaw ("var " <> name <> " " <> goTypeToStr expectedGoType <> " = " <> printGoExpr (unboxGoExpr resBinding.expr resBinding.exprType expectedGoType)))
        in
          { stmts: resBinding.stmts <> StmtLeaf (GoRaw ("// TAST (Let): " <> name <> " -> " <> goTypeToStr expectedGoType)) <> letStmt <> resBody.stmts, expr: resBody.expr, exprType: resBody.exprType, nextId: resBody.nextId }

      LetRec lvl bindings body ->
        let
          allocRes = foldl
            ( \acc (Tuple (Ident ident) val) ->
                let
                  oldName = localId (Just (Ident ident)) lvl
                  gId = unsafePerformEffect do
                    curr <- Ref.read helpersRef
                    Ref.modify_ (\r -> r { globalId = r.globalId + 1 }) helpersRef
                    pure curr.globalId
                  newName = oldName <> "_" <> show acc.nextId <> "_" <> show gId
                  expectedGoType = exprTypeToGoType (unsafePerformEffect (Ref.read helpersRef)).pointerAdtPaths (unsafePerformEffect (Ref.read helpersRef)).enumAdts (unsafePerformEffect (Ref.read helpersRef)).elidedCtors modNameStr (getExprType val)
                in
                  { newBound: Map.insert oldName { name: newName, goType: expectedGoType } acc.newBound, newNames: Array.snoc acc.newNames { oldName, newName }, exprType: TypeValue, nextId: acc.nextId + 1 }
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
                          in Tuple idStr (exprTypeToGoType (unsafePerformEffect (Ref.read helpersRef)).pointerAdtPaths (unsafePerformEffect (Ref.read helpersRef)).enumAdts (unsafePerformEffect (Ref.read helpersRef)).elidedCtors modNameStr t)
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
                            in Tuple idStr (exprTypeToGoType (unsafePerformEffect (Ref.read helpersRef)).pointerAdtPaths (unsafePerformEffect (Ref.read helpersRef)).enumAdts (unsafePerformEffect (Ref.read helpersRef)).elidedCtors modNameStr t)
                          ) fn.args
                        currentLoopCtx = [ { ident: newName, params: fn.args, loopParams: map (\p -> p <> "_loop") fn.args, goTypes: map snd paramsWithTypes } ]
                        loopBound = foldl (\acc (Tuple idStr goT) -> Map.insert idStr { name: idStr, goType: goT } acc) allocRes.newBound paramsWithTypes
                        resBodyMut = translateExprImpl_ helpersRef (depth + 1) modNameStr combinedRecVars moduleArities loopBound (Just newName) currentLoopCtx true false currNextId fn.body
                        
                        loopParams = map (\(Tuple idStr _) -> idStr <> "_loop") paramsWithTypes
                        initVars = Array.concatMap (\(Tuple p goT) -> [ GoRaw ("var " <> p <> " " <> goTypeToStr goT <> " = " <> p <> "_loop"), GoRaw ("_ = " <> p) ]) paramsWithTypes
                        
                        funcBody = GoFor newName (initVars <> flattenStmts resBodyMut.stmts <> [ GoReturn (boxGoExpr resBodyMut.expr resBodyMut.exprType) ])
                        
                        -- Declare native loop vars from the boxed closure args
                        iifeInitVars = map (\(Tuple p goT) -> 
                          if goT == TypeValue then GoRaw ("var " <> p <> "_loop gopurs_runtime.Value = " <> p <> "_loop_val")
                          else GoRaw ("var " <> p <> "_loop " <> goTypeToStr goT <> " = " <> printGoExpr (coerceGoExpr (GoRaw (p <> "_loop_val")) TypeValue goT))
                        ) paramsWithTypes
                        
                        iife = GoRaw ("func() gopurs_runtime.Value {\n" <> printGoExpr (GoBlock (iifeInitVars <> [funcBody])) <> "\n}()")
                        funcExpr = Array.foldr (\(Tuple p goT) acc -> GoCall (GoSelector (GoVar "gopurs_runtime") "Func") [ GoRaw ("func(" <> p <> "_loop_val gopurs_runtime.Value) gopurs_runtime.Value {\nreturn " <> printGoExpr acc <> "\n}") ]) iife paramsWithTypes
                      in
                        Tuple (Array.snoc accStmts (GoMutate newName funcExpr)) resBodyMut.nextId
                  )
                  (Tuple [] allocRes.nextId)
                  fns
                
                resBodyOuter = translateExprImpl_ helpersRef (depth + 1) modNameStr combinedRecVars moduleArities allocRes.newBound Nothing loopCtx isTail inEffectBlock nextId' body
              in
                { stmts: foldMap StmtLeaf declStmts <> foldMap StmtLeaf fnWrapperStmts <> resBodyOuter.stmts, expr: resBodyOuter.expr, exprType: resBodyOuter.exprType, nextId: resBodyOuter.nextId }
            
            Nothing ->
              let
                accBindings = foldl
                  ( \acc (Tuple (Tuple (Ident ident) val) alloc) ->
                      let
                        res = translateExprImpl_ helpersRef (depth + 1) modNameStr combinedRecVars moduleArities allocRes.newBound Nothing [] false false acc.nextId val
                        expectedGoType = (fromMaybe { name: alloc.newName, goType: TypeValue } (Map.lookup alloc.oldName allocRes.newBound)).goType
                        assignedVal = if expectedGoType == res.exprType then res.expr else unboxGoExpr res.expr res.exprType expectedGoType
                      in
                        { stmts: acc.stmts <> res.stmts, exprs: Array.snoc acc.exprs { key: alloc.newName, value: assignedVal, goType: expectedGoType }, exprType: TypeValue, nextId: res.nextId }
                  )
                  { stmts: StmtEmpty, exprs: [], exprType: TypeValue, nextId: allocRes.nextId }
                  (Array.zip (toArray bindings) allocRes.newNames)

                declStmts = map (\b -> GoRaw ("var " <> b.key <> " " <> goTypeToStr b.goType <> "\n_ = " <> b.key)) accBindings.exprs
                assignStmts = map (\b -> GoMutate b.key b.value) accBindings.exprs

                resBody = translateExprImpl_ helpersRef (depth + 1) modNameStr combinedRecVars moduleArities allocRes.newBound Nothing loopCtx isTail inEffectBlock accBindings.nextId body
              in
                { stmts: foldMap StmtLeaf declStmts <> accBindings.stmts <> foldMap StmtLeaf assignStmts <> resBody.stmts, expr: resBody.expr, exprType: resBody.exprType, nextId: resBody.nextId }

      Accessor obj accessor ->
        let
          resObj = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars moduleArities bound Nothing [] false false nextId obj
        in
          case accessor of
            GetProp prop ->
              case resObj.exprType of
                TypeRecord fields -> { stmts: resObj.stmts, expr: GoRecordAccess (boxGoExpr resObj.expr resObj.exprType) prop, exprType: TypeValue, nextId: resObj.nextId }
                TypeStructPointer _ fullName _ _ ->
                  let
                    h = unsafePerformEffect (Ref.read helpersRef)

                  in case Map.lookup fullName h.classDeclsFields of
                    Just info ->
                      case Array.findIndex (\f -> f.name == prop) info.fields of
                        Just idx ->
                          let
                            unboxedObj = unboxGoExpr resObj.expr resObj.exprType resObj.exprType
                            fieldExpr = GoStructAccess unboxedObj ("V" <> show idx)
                            boxedFieldExpr = GoCall (GoSelector (GoVar "gopurs_runtime") "Box") [ fieldExpr ]
                          in { stmts: resObj.stmts, expr: boxedFieldExpr, exprType: TypeValue, nextId: resObj.nextId }
                        Nothing ->
                          let _ = unit
                          in { stmts: resObj.stmts, expr: GoRecordAccess (boxGoExpr resObj.expr resObj.exprType) prop, exprType: TypeValue, nextId: resObj.nextId }
                    Nothing -> { stmts: resObj.stmts, expr: GoRecordAccess (boxGoExpr resObj.expr resObj.exprType) prop, exprType: TypeValue, nextId: resObj.nextId }
                TypeValue -> { stmts: resObj.stmts, expr: GoRecordAccess (boxGoExpr resObj.expr resObj.exprType) prop, exprType: TypeValue, nextId: resObj.nextId }
                _ -> { stmts: resObj.stmts, expr: GoRecordAccess (boxGoExpr resObj.expr resObj.exprType) prop, exprType: TypeValue, nextId: resObj.nextId }
            GetIndex idx -> { stmts: resObj.stmts, expr: GoCall (GoSelector (GoVar "gopurs_runtime") "ArrayAccess") [ (boxGoExpr resObj.expr resObj.exprType), GoInt idx ], exprType: TypeValue, nextId: resObj.nextId }
            GetCtorField (Qualified mbMod _) _ _ (Ident ctorName) _ idx ->
              let
                defMod = case mbMod of
                  Just (ModuleName mod) -> String.replaceAll (Pattern ".") (Replacement "_") mod
                  Nothing -> modNameStr
                structName = "Constructor_" <> defMod <> "_" <> sanitizeName ctorName
                key = defMod <> "." <> ctorName
                helpers = unsafePerformEffect (Ref.read helpersRef)
              in
                if Set.member structName elidedCtors then
                  { stmts: resObj.stmts, expr: coerceGoExpr resObj.expr resObj.exprType TypeValue, exprType: TypeValue, nextId: resObj.nextId }
                else
                  let
                    fields = fromMaybe [] (map _.fields (Map.lookup key helpers.ctorTypes))
                    monoStructName = structName
                      

                      
                    expectedType = case Array.index fields idx of
                      Just ty -> exprTypeToGoType (unsafePerformEffect (Ref.read helpersRef)).pointerAdtPaths (unsafePerformEffect (Ref.read helpersRef)).enumAdts (unsafePerformEffect (Ref.read helpersRef)).elidedCtors modNameStr ty
                      Nothing -> TypeValue
                      
                    typeArgs = case getExprType obj of
                      ADT fullName _ tArgs -> 
                        let
                          mapped = map (exprTypeToGoType (unsafePerformEffect (Ref.read helpersRef)).pointerAdtPaths (unsafePerformEffect (Ref.read helpersRef)).enumAdts (unsafePerformEffect (Ref.read helpersRef)).elidedCtors modNameStr) tArgs
                          arity = case Map.lookup fullName (unsafePerformEffect (Ref.read helpersRef)).pointerAdtPaths of
                            Just info -> info.arity
                            Nothing -> Array.length mapped
                        in Array.take arity mapped
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
                              genericTy = structFieldGoType (unsafePerformEffect (Ref.read helpersRef)).pointerAdtPaths (unsafePerformEffect (Ref.read helpersRef)).enumAdts (unsafePerformEffect (Ref.read helpersRef)).elidedCtors ctorInfo.vars modNameStr (fromMaybe (TypeVar "") (Array.index fields idx))
                            in instantiateGenericGoType env genericTy
                          Nothing -> expectedType
                      _ -> expectedType
                    
                    exprAccess = if isNative then
                                   GoConstructorAccess resObj.expr monoStructName typeArgs idx true
                                 else
                                   GoConstructorAccess (boxGoExpr resObj.expr resObj.exprType) monoStructName typeArgs idx false
                  in
                    if isNative then
                      { stmts: resObj.stmts, expr: coerceGoExpr exprAccess actualFieldType expectedType, exprType: expectedType, nextId: resObj.nextId }
                    else
                      { stmts: resObj.stmts, expr: coerceGoExpr exprAccess expectedType expectedType, exprType: expectedType, nextId: resObj.nextId }

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
          case resObj.exprType of
            TypeRecord fields ->
              let
                staticUpdates = Array.catMaybes (map (\(Tuple key val) ->
                  case Array.findIndex (\(Tuple k _) -> k == key) fields of
                    Just idx -> Just (Tuple idx val)
                    Nothing -> Nothing
                ) accProps.exprs)
              in
                if Array.length staticUpdates == Array.length accProps.exprs then
                  { stmts: resObj.stmts <> accProps.stmts, expr: GoRecordUpdateStatic (boxGoExpr resObj.expr resObj.exprType) (Array.length fields) staticUpdates accProps.exprs, exprType: TypeValue, nextId: accProps.nextId }
                else
                  { stmts: resObj.stmts <> accProps.stmts, expr: GoRecordUpdateDict (boxGoExpr resObj.expr resObj.exprType) accProps.exprs, exprType: TypeValue, nextId: accProps.nextId }
            _ ->
              { stmts: resObj.stmts <> accProps.stmts, expr: GoRecordUpdateDict (boxGoExpr resObj.expr resObj.exprType) accProps.exprs, exprType: TypeValue, nextId: accProps.nextId }

      CtorDef _ _ (Ident name) fields ->
        let
          structName = "Constructor_" <> modNameStr <> "_" <> sanitizeName name
          baseStructName = getBaseStructName modNameStr Nothing name
          key = modNameStr <> "." <> name
          helpers = unsafePerformEffect (Ref.read helpersRef)
          ctorType = getExprType tcoExpr
          expectedGoType = exprTypeToGoType (unsafePerformEffect (Ref.read helpersRef)).pointerAdtPaths (unsafePerformEffect (Ref.read helpersRef)).enumAdts (unsafePerformEffect (Ref.read helpersRef)).elidedCtors modNameStr ctorType
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
                mapped = map (exprTypeToGoType (unsafePerformEffect (Ref.read helpersRef)).pointerAdtPaths (unsafePerformEffect (Ref.read helpersRef)).enumAdts (unsafePerformEffect (Ref.read helpersRef)).elidedCtors modNameStr) tArgs
                arity = case Map.lookup fullName (unsafePerformEffect (Ref.read helpersRef)).pointerAdtPaths of
                  Just info -> info.arity
                  Nothing -> Array.length vars'
              in Array.take arity mapped
            TypeApp fn arg ->
              let
                unwrapTypeApp (TypeApp f a) acc = unwrapTypeApp f (a <> acc)
                unwrapTypeApp other acc = Tuple other acc
              in case unwrapTypeApp (getExprType tcoExpr) [] of
                Tuple (ADT fnName _ tArgs) allArgs ->
                  let mapped = map (exprTypeToGoType (unsafePerformEffect (Ref.read helpersRef)).pointerAdtPaths (unsafePerformEffect (Ref.read helpersRef)).enumAdts (unsafePerformEffect (Ref.read helpersRef)).elidedCtors modNameStr) (tArgs <> allArgs)
                  in Array.take (Array.length vars') mapped
                _ -> map (const TypeValue) vars'
            _ -> map (const TypeValue) vars'
          typeArgsStr = ""
          instMap = Map.fromFoldable (Array.zip vars' typeArgs)
          coercedFields = Array.mapWithIndex (\i f -> 
            let
              expectedExprType = case Array.index fields' i of
                Just ty -> ty
                Nothing -> Any
              expectedType = case Array.index fields' i of
                Just ty -> 
                  let genericGoType = exprTypeToGenericGoType helpers.pointerAdtPaths helpers.enumAdts helpers.elidedCtors vars' modNameStr ty
                  in instantiateGenericGoType instMap genericGoType
                Nothing -> TypeValue
            in coerceGoExpr (GoVar (sanitizeName f)) TypeValue expectedType
          ) fields
          isElided = Set.member structName helpers.elidedCtors
          isPointerAdtLeaf = Map.member baseStructName helpers.pointerAdtLeaves
          isEnum = Set.member baseStructName helpers.enumCtors
          boxedCtor = boxGoExpr (GoConstructor (hashString baseStructName) structName typeArgs coercedFields) (TypeStructPointer baseStructName fullName (structName <> typeArgsStr) typeArgs)
          
          finalExprType = if isEnum then TypeUint32 
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
              in TypeStructPointer nodeBaseStruct fullName nodeFullPath typeArgs
            else if Array.length fields == 0 then
              TypeStructPointer baseStructName fullName (structName <> typeArgsStr) typeArgs
            else TypeValue
          
          funcExpr = if isElided then
              case Array.head fields of
                Just f -> GoCall (GoSelector (GoVar "gopurs_runtime") "Func") [ GoRaw ("func(" <> sanitizeName f <> " gopurs_runtime.Value) gopurs_runtime.Value {\nreturn " <> printGoExpr (coerceGoExpr (GoVar (sanitizeName f)) TypeValue TypeValue) <> "\n}") ]
                Nothing -> Array.foldr (\f inner -> GoCall (GoSelector (GoVar "gopurs_runtime") "Func") [ GoRaw ("func(" <> sanitizeName f <> " gopurs_runtime.Value) gopurs_runtime.Value {\nreturn " <> printGoExpr inner <> "\n}") ]) boxedCtor fields
            else if isPointerAdtLeaf then 
              let 
                nodeCtorName = case Map.lookup baseStructName helpers.pointerAdtLeaves of
                  Just info -> info.nodeCtor
                  Nothing -> ""
                nodeStruct = "Constructor_" <> modNameStr <> "_" <> sanitizeName nodeCtorName <> typeArgsStr
              in GoRaw ("(*" <> nodeStruct <> ")(nil)")
            else if isEnum then GoRaw (hashString baseStructName)
            else if Array.length fields == 0 then
              GoConstructor (hashString baseStructName) structName typeArgs coercedFields
            else
              Array.foldr (\f inner -> GoCall (GoSelector (GoVar "gopurs_runtime") "Func") [ GoRaw ("func(" <> sanitizeName f <> " gopurs_runtime.Value) gopurs_runtime.Value {\nreturn " <> printGoExpr inner <> "\n}") ]) boxedCtor fields
        in
          { stmts: StmtEmpty, expr: funcExpr, exprType: finalExprType, nextId }

      CtorSaturated (Qualified mbMod _) _ _ (Ident name) props ->
        let
          baseStructName = getBaseStructName modNameStr mbMod name
          modPart = case mbMod of
            Just (ModuleName mn) -> String.replaceAll (Pattern ".") (Replacement "_") mn
            Nothing -> modNameStr
          structName = "Constructor_" <> modPart <> "_" <> sanitizeName name
          key = modPart <> "." <> name
          helpers = unsafePerformEffect (Ref.read helpersRef)
          
          ctorType = getExprType tcoExpr
          expectedGoType = exprTypeToGoType (unsafePerformEffect (Ref.read helpersRef)).pointerAdtPaths (unsafePerformEffect (Ref.read helpersRef)).enumAdts (unsafePerformEffect (Ref.read helpersRef)).elidedCtors modNameStr ctorType
          
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
                mapped = map (exprTypeToGoType (unsafePerformEffect (Ref.read helpersRef)).pointerAdtPaths (unsafePerformEffect (Ref.read helpersRef)).enumAdts (unsafePerformEffect (Ref.read helpersRef)).elidedCtors modNameStr) tArgs
                arity = case Map.lookup fullName (unsafePerformEffect (Ref.read helpersRef)).pointerAdtPaths of
                  Just info -> info.arity
                  Nothing -> Array.length vars
              in Array.take arity mapped
            _ -> map (const TypeValue) vars
          instMap = Map.fromFoldable (Array.zip vars typeArgs)
          
          accProps = foldl
            ( \acc (Tuple _ val) ->
                let
                  expectedExprType = case Array.index fields acc.fieldIdx of
                    Just ty -> ty
                    Nothing -> Any
                  expectedType = case Array.index fields acc.fieldIdx of
                    Just ty -> 
                      let genericGoType = exprTypeToGenericGoType helpers.pointerAdtPaths helpers.enumAdts helpers.elidedCtors vars modNameStr ty
                      in instantiateGenericGoType instMap genericGoType
                    Nothing -> TypeValue
                    
                  newBound = case unwrapTcoExpr val, extractExprFuncType expectedExprType of
                    Abs args _, Just { fArgs } ->
                      let paramsWithTypes = Array.zipWith (\(Tuple mbI lvl) fArgTy -> Tuple (localId mbI lvl) (exprTypeToGoType helpers.pointerAdtPaths helpers.enumAdts helpers.elidedCtors modNameStr fArgTy)) (toArray args) (fArgs <> Array.replicate (Array.length (toArray args) - Array.length fArgs) Any)
                      in foldl (\b (Tuple idStr goType) -> Map.insert idStr { name: idStr, goType } b) bound paramsWithTypes
                    UncurriedAbs args _, Just { fArgs } ->
                      let paramsWithTypes = Array.zipWith (\(Tuple mbI lvl) fArgTy -> Tuple (localId mbI lvl) (exprTypeToGoType helpers.pointerAdtPaths helpers.enumAdts helpers.elidedCtors modNameStr fArgTy)) args (fArgs <> Array.replicate (Array.length args - Array.length fArgs) Any)
                      in foldl (\b (Tuple idStr goType) -> Map.insert idStr { name: idStr, goType } b) bound paramsWithTypes
                    _, _ -> bound

                  resVal = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars moduleArities newBound Nothing [] false false acc.nextId val
                  coercedExpr = coerceGoExpr resVal.expr resVal.exprType expectedType
                in
                  { stmts: acc.stmts <> resVal.stmts, exprs: Array.snoc acc.exprs coercedExpr, exprType: TypeValue, nextId: resVal.nextId, fieldIdx: acc.fieldIdx + 1 }
            )
            { stmts: StmtEmpty, exprs: [], exprType: TypeValue, nextId, fieldIdx: 0 }
            props
            
          isElided = Set.member structName helpers.elidedCtors
          isPointerAdtLeaf = Map.member baseStructName helpers.pointerAdtLeaves
          
          deadVarOptRaw = if Array.length props == 0 || isElided || isPointerAdtLeaf then Nothing else Array.head (Array.mapMaybe (\(Tuple _ v) -> 
              let reused = unsafePerformEffect (Ref.read globalReusedVars)
              in
              if v.goType == expectedGoType 
                 && not (Set.member v.name reused)
                 && not (Set.member v.name (freeVars tcoExpr)) 
              then Just v.name 
              else Nothing
            ) (Array.fromFoldable (map (\v -> Tuple v.name v) (Map.values bound))))
            
          deadVarOpt = deadVarOptRaw
          _ = unsafePerformEffect (case deadVarOpt of
                Just n -> Ref.modify_ (Set.insert n) globalReusedVars
                Nothing -> pure unit)
                
          modPart' = case mbMod of
            Just (ModuleName mod) -> String.replaceAll (Pattern ".") (Replacement "_") mod
            _ -> modNameStr
          monoStructName = "Constructor_" <> modPart' <> "_" <> sanitizeName name

          typeArgs = case ctorType of
            ADT fullName _ tArgs ->
              let
                mapped = map (exprTypeToGoType (unsafePerformEffect (Ref.read helpersRef)).pointerAdtPaths (unsafePerformEffect (Ref.read helpersRef)).enumAdts (unsafePerformEffect (Ref.read helpersRef)).elidedCtors modNameStr) tArgs
                arity = case Map.lookup fullName (unsafePerformEffect (Ref.read helpersRef)).pointerAdtPaths of
                  Just info -> info.arity
                  Nothing -> Array.length mapped
              in Array.take arity mapped
            _ -> case Map.lookup key helpers.ctorTypes of
              Just ctorInfo -> map (const TypeValue) ctorInfo.vars
              Nothing -> []
          typeArgsStr = ""
          fullPath = monoStructName <> typeArgsStr
          isEnum = Set.member baseStructName helpers.enumCtors
          res = if isElided then
              case Array.head accProps.exprs of
                Just expr -> { expr: boxGoExpr expr (fromMaybe TypeValue (map (exprTypeToGoType (unsafePerformEffect (Ref.read helpersRef)).pointerAdtPaths (unsafePerformEffect (Ref.read helpersRef)).enumAdts (unsafePerformEffect (Ref.read helpersRef)).elidedCtors modNameStr) (Array.index fields 0))), exprType: TypeValue }
                Nothing -> { expr: GoConstructor (hashString baseStructName) monoStructName typeArgs accProps.exprs, exprType: TypeStructPointer baseStructName fullName fullPath typeArgs }
            else if isPointerAdtLeaf then 
              let 
                nodeInfo = Map.lookup baseStructName helpers.pointerAdtLeaves
                nodeCtorName = case nodeInfo of
                  Just info -> info.nodeCtor
                  Nothing -> ""
                nodeBaseStruct = case nodeInfo of
                  Just info -> info.nodeBaseStruct
                  Nothing -> ""
                nodeStruct = "Constructor_" <> modPart <> "_" <> sanitizeName nodeCtorName
                nodeFullPath = nodeStruct <> typeArgsStr
              in { expr: GoRaw ("(*" <> nodeFullPath <> ")(nil)"), exprType: TypeStructPointer nodeBaseStruct fullName nodeFullPath typeArgs }
            else if isEnum then { expr: GoRaw (hashString baseStructName), exprType: TypeUint32 }
            else { expr: GoConstructor (hashString baseStructName) monoStructName typeArgs accProps.exprs, exprType: TypeStructPointer baseStructName fullName fullPath typeArgs }
        in
          { stmts: accProps.stmts, expr: res.expr, exprType: res.exprType, nextId: accProps.nextId }

      Fail msg ->
        { stmts: StmtEmpty, expr: GoRaw ("func() gopurs_runtime.Value { panic(" <> printGoExpr (GoString msg) <> ") }()"), exprType: TypeValue, nextId }

      Branch branches def ->
        let
          resDef = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars moduleArities bound Nothing loopCtx isTail false nextId def
          
          computedBranches = foldl
            ( \acc (Pair condExpr bodyExpr) ->
                let
                  resCond = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars moduleArities bound Nothing [] false false acc.nextId condExpr
                  resBody = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars moduleArities bound Nothing loopCtx isTail false resCond.nextId bodyExpr
                in
                  { nextId: resBody.nextId, results: acc.results <> [{ cond: resCond, body: resBody }] }
            )
            { nextId: resDef.nextId, results: [] }
            (toArray branches)
            
          allTypes = [ resDef.exprType ] <> map (\r -> r.body.exprType) computedBranches.results
          validTypes = Array.filter (\t -> case t of 
            TypeValue -> false
            _ -> true) allTypes
          allTypesStr = map goTypeToStr validTypes
          
          expectedGoType = 
            let nubbed = Array.nub allTypesStr
            in if Array.length nubbed == 1 
               then fromMaybe TypeValue (Array.head validTypes)
               else if Array.length (Array.nub (map goTypeToStr allTypes)) == 1 then resDef.exprType else TypeValue
            
          tmpVar = "__t" <> show computedBranches.nextId
          declTmp = StmtLeaf (GoRaw ("var " <> tmpVar <> " " <> goTypeToStr expectedGoType))
          labelName = "end_branch_" <> show computedBranches.nextId

          buildIfs = foldl
            ( \acc r ->
                let
                  goIf = GoIfElse (unboxGoExpr r.cond.expr r.cond.exprType TypeBool) (flattenStmts r.body.stmts <> [ GoMutate tmpVar (coerceGoExpr r.body.expr r.body.exprType expectedGoType), GoRaw ("goto " <> labelName) ]) []
                in
                  acc <> StmtLeaf (GoRaw "{") <> r.cond.stmts <> StmtLeaf goIf <> StmtLeaf (GoRaw "}")
            )
            StmtEmpty
            computedBranches.results
        in
          { stmts: declTmp <> buildIfs <> StmtLeaf (GoRaw "{") <> resDef.stmts <> StmtLeaf (GoMutate tmpVar (coerceGoExpr resDef.expr resDef.exprType expectedGoType)) <> StmtLeaf (GoRaw "}") <> StmtLeaf (GoRaw (labelName <> ":")), expr: GoVar tmpVar, exprType: expectedGoType, nextId: computedBranches.nextId + 1 }

      PrimOp op -> case op of
        Op1 op1 e ->
          let
            resE = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars moduleArities bound Nothing [] false false nextId e
            goOp = case op1 of
              OpBooleanNot -> { stmts: resE.stmts, expr: GoBinOp "!=" (unboxGoExpr resE.expr resE.exprType TypeBool) (GoRaw "true"), exprType: TypeBool, nextId: resE.nextId }
              OpIntNegate -> { stmts: resE.stmts, expr: GoPrefixOp "-" (unboxGoExpr resE.expr resE.exprType TypeInt64), exprType: TypeInt64, nextId: resE.nextId }
              OpIntBitNot -> { stmts: resE.stmts, expr: GoBinOp "^" (GoRaw "^0") (unboxGoExpr resE.expr resE.exprType TypeInt64), exprType: TypeInt64, nextId: resE.nextId }
              OpNumberNegate -> { stmts: resE.stmts, expr: GoPrefixOp "-" (unboxGoExpr resE.expr resE.exprType TypeFloat64), exprType: TypeFloat64, nextId: resE.nextId }
              OpIsTag (Qualified mbMod (Ident tag)) ->
                let
                  baseStructName = getBaseStructName modNameStr mbMod tag
                  hashStr = hashString baseStructName
                  helpers = unsafePerformEffect (Ref.read helpersRef)
                  
                  isNativePointer = case resE.exprType of
                    TypeStructPointer typedBaseStructName _ _ _ ->
                      typedBaseStructName == baseStructName ||
                        (case Map.lookup baseStructName helpers.pointerAdtLeaves of
                           Just nodeInfo -> typedBaseStructName == nodeInfo.nodeBaseStruct
                           Nothing -> false)
                    _ -> false
                in
                  case resE.expr of
                    GoVar _ ->
                      let
                        exprStr = if isNativePointer then
                          case Map.lookup baseStructName helpers.pointerAdtLeaves of
                            Just _ -> "(" <> printGoExpr resE.expr <> " == nil)"
                            Nothing -> "(" <> printGoExpr resE.expr <> " != nil)"
                        else case Map.lookup baseStructName helpers.pointerAdtLeaves of
                          Just nodeInfo -> "(" <> printGoExpr (boxGoExpr resE.expr resE.exprType) <> ".Type == 9 && " <> printGoExpr (boxGoExpr resE.expr resE.exprType) <> ".IntVal == " <> hashString nodeInfo.nodeBaseStruct <> " && " <> printGoExpr (boxGoExpr resE.expr resE.exprType) <> ".UnsafePtr == nil)"
                          Nothing -> if Set.member baseStructName helpers.pointerAdtNodes then
                            "(" <> printGoExpr (boxGoExpr resE.expr resE.exprType) <> ".Type == 9 && " <> printGoExpr (boxGoExpr resE.expr resE.exprType) <> ".IntVal == " <> hashStr <> " && " <> printGoExpr (boxGoExpr resE.expr resE.exprType) <> ".UnsafePtr != nil)"
                          else if Set.member baseStructName helpers.enumCtors then
                            "(" <> printGoExpr (unboxGoExpr resE.expr resE.exprType TypeUint32) <> " == " <> hashStr <> ")"
                          else
                            "(" <> printGoExpr (boxGoExpr resE.expr resE.exprType) <> ".Type == 9 && " <> printGoExpr (boxGoExpr resE.expr resE.exprType) <> ".IntVal == " <> hashStr <> ")"
                      in { stmts: resE.stmts, expr: GoRaw exprStr, exprType: TypeBool, nextId: resE.nextId }
                    _ ->
                      let
                        tmpVar = "__t_tag_" <> show resE.nextId
                        declTmp = if isNativePointer || resE.exprType /= TypeValue then
                           StmtLeaf (GoRaw ("var " <> tmpVar <> " " <> goTypeToStr resE.exprType <> " = " <> printGoExpr resE.expr))
                        else
                           StmtLeaf (GoRaw ("var " <> tmpVar <> " gopurs_runtime.Value = " <> printGoExpr (boxGoExpr resE.expr resE.exprType)))

                        exprStr = if isNativePointer then
                          case Map.lookup baseStructName helpers.pointerAdtLeaves of
                            Just _ -> "(" <> tmpVar <> " == nil)"
                            Nothing -> "(" <> tmpVar <> " != nil)"
                        else if resE.exprType /= TypeValue then
                          "(uint32(" <> tmpVar <> ") == " <> hashStr <> ")"
                        else case Map.lookup baseStructName helpers.pointerAdtLeaves of
                          Just nodeInfo -> "(" <> tmpVar <> ".Type == 9 && " <> tmpVar <> ".IntVal == " <> hashString nodeInfo.nodeBaseStruct <> " && " <> tmpVar <> ".UnsafePtr == nil)"
                          Nothing -> if Set.member baseStructName helpers.pointerAdtNodes then
                            "(" <> tmpVar <> ".Type == 9 && " <> tmpVar <> ".IntVal == " <> hashStr <> " && " <> tmpVar <> ".UnsafePtr != nil)"
                          else if Set.member baseStructName helpers.enumCtors then
                            "(" <> printGoExpr (unboxGoExpr (GoVar tmpVar) TypeValue TypeUint32) <> " == " <> hashStr <> ")"
                          else
                            "(" <> tmpVar <> ".Type == 9 && " <> tmpVar <> ".IntVal == " <> hashStr <> ")"
                      in { stmts: resE.stmts <> declTmp, expr: GoRaw exprStr, exprType: TypeBool, nextId: resE.nextId + 1 }
              OpArrayLength -> 
                case resE.exprType of
                  TypeNativeArray _ -> { stmts: resE.stmts, expr: GoCall (GoSelector (GoVar "gopurs_runtime") "Int") [ GoCall (GoVar "int64") [ GoCall (GoVar "len") [ resE.expr ] ] ], exprType: TypeValue, nextId: resE.nextId }
                  _ -> { stmts: resE.stmts, expr: GoCall (GoSelector (GoVar "gopurs_runtime") "Int") [ GoCall (GoVar "int64") [ GoCall (GoSelector (GoVar "gopurs_runtime") "ArrayLength") [ boxGoExpr resE.expr resE.exprType ] ] ], exprType: TypeValue, nextId: resE.nextId }
          in
            goOp
        Op2 OpBooleanAnd e1 e2 ->
          let
            res1 = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars moduleArities bound Nothing [] false false nextId e1
            res2 = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars moduleArities bound Nothing [] false false res1.nextId e2
            isEmptyStmts StmtEmpty = true
            isEmptyStmts (StmtAppend s1 s2) = isEmptyStmts s1 && isEmptyStmts s2
            isEmptyStmts _ = false
          in
            if isEmptyStmts res2.stmts then
              { expr: GoBinOp "&&" (unboxGoExpr res1.expr res1.exprType TypeBool) (unboxGoExpr res2.expr res2.exprType TypeBool), exprType: TypeBool, stmts: res1.stmts <> res2.stmts, nextId: res2.nextId }
            else
              let
                tmpVar = "__t_and_" <> show res2.nextId
                declTmp = StmtLeaf (GoRaw ("var " <> tmpVar <> " bool = false\nif " <> printGoExpr (unboxGoExpr res1.expr res1.exprType TypeBool) <> " {\n"))
                assignTmp = StmtLeaf (GoRaw (tmpVar <> " = " <> printGoExpr (unboxGoExpr res2.expr res2.exprType TypeBool) <> "\n}"))
              in
                { expr: GoRaw tmpVar, exprType: TypeBool, stmts: res1.stmts <> declTmp <> res2.stmts <> assignTmp, nextId: res2.nextId + 1 }
        Op2 OpBooleanOr e1 e2 ->
          let
            res1 = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars moduleArities bound Nothing [] false false nextId e1
            res2 = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars moduleArities bound Nothing [] false false res1.nextId e2
            isEmptyStmts StmtEmpty = true
            isEmptyStmts (StmtAppend s1 s2) = isEmptyStmts s1 && isEmptyStmts s2
            isEmptyStmts _ = false
          in
            if isEmptyStmts res2.stmts then
              { expr: GoBinOp "||" (unboxGoExpr res1.expr res1.exprType TypeBool) (unboxGoExpr res2.expr res2.exprType TypeBool), exprType: TypeBool, stmts: res1.stmts <> res2.stmts, nextId: res2.nextId }
            else
              let
                tmpVar = "__t_or_" <> show res2.nextId
                declTmp = StmtLeaf (GoRaw ("var " <> tmpVar <> " bool = true\nif !(" <> printGoExpr (unboxGoExpr res1.expr res1.exprType TypeBool) <> ") {\n"))
                assignTmp = StmtLeaf (GoRaw (tmpVar <> " = " <> printGoExpr (unboxGoExpr res2.expr res2.exprType TypeBool) <> "\n}"))
              in
                { expr: GoRaw tmpVar, exprType: TypeBool, stmts: res1.stmts <> declTmp <> res2.stmts <> assignTmp, nextId: res2.nextId + 1 }
        Op2 op2 e1 e2 ->
          let
            res1 = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars moduleArities bound Nothing [] false false nextId e1
            res2 = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars moduleArities bound Nothing [] false false res1.nextId e2
            goOp = case op2 of
              OpArrayIndex -> 
                case res1.exprType of
                  TypeNativeArray innerType -> { expr: boxGoExpr (GoRaw (printGoExpr res1.expr <> "[" <> printGoExpr (unboxGoExpr res2.expr res2.exprType TypeInt64) <> "]")) innerType, exprType: TypeValue }
                  _ -> { expr: GoCall (GoSelector (GoVar "gopurs_runtime") "ArrayAccess") [ boxGoExpr res1.expr res1.exprType, GoCall (GoVar "int") [ unboxGoExpr res2.expr res2.exprType TypeInt64 ] ], exprType: TypeValue }
              OpIntNum OpAdd -> { expr: GoBinOp "+" (unboxGoExpr res1.expr res1.exprType TypeInt64) (unboxGoExpr res2.expr res2.exprType TypeInt64), exprType: TypeInt64 }
              OpIntNum OpSubtract -> { expr: GoBinOp "-" (unboxGoExpr res1.expr res1.exprType TypeInt64) (unboxGoExpr res2.expr res2.exprType TypeInt64), exprType: TypeInt64 }
              OpIntNum OpMultiply -> { expr: GoBinOp "*" (unboxGoExpr res1.expr res1.exprType TypeInt64) (unboxGoExpr res2.expr res2.exprType TypeInt64), exprType: TypeInt64 }
              OpIntNum OpDivide -> { expr: GoBinOp "/" (unboxGoExpr res1.expr res1.exprType TypeInt64) (unboxGoExpr res2.expr res2.exprType TypeInt64), exprType: TypeInt64 }
              OpIntBitAnd -> { expr: GoBinOp "&" (unboxGoExpr res1.expr res1.exprType TypeInt64) (unboxGoExpr res2.expr res2.exprType TypeInt64), exprType: TypeInt64 }
              OpIntBitOr -> { expr: GoBinOp "|" (unboxGoExpr res1.expr res1.exprType TypeInt64) (unboxGoExpr res2.expr res2.exprType TypeInt64), exprType: TypeInt64 }
              OpIntBitXor -> { expr: GoBinOp "^" (unboxGoExpr res1.expr res1.exprType TypeInt64) (unboxGoExpr res2.expr res2.exprType TypeInt64), exprType: TypeInt64 }
              OpIntBitShiftLeft -> { expr: GoBinOp "<<" (unboxGoExpr res1.expr res1.exprType TypeInt64) (unboxGoExpr res2.expr res2.exprType TypeInt64), exprType: TypeInt64 }
              OpIntBitShiftRight -> { expr: GoBinOp ">>" (unboxGoExpr res1.expr res1.exprType TypeInt64) (unboxGoExpr res2.expr res2.exprType TypeInt64), exprType: TypeInt64 }
              OpIntBitZeroFillShiftRight -> { expr: GoCall (GoSelector (GoVar "gopurs_runtime") "Zshr") [ boxGoExpr res1.expr res1.exprType, boxGoExpr res2.expr res2.exprType ], exprType: TypeValue }
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
              OpBooleanAnd -> { expr: GoRaw "panic(\"unreachable\")", exprType: TypeValue }
              OpBooleanOr -> { expr: GoRaw "panic(\"unreachable\")", exprType: TypeValue }
          in
            { stmts: res1.stmts <> res2.stmts, expr: goOp.expr, exprType: goOp.exprType, nextId: res2.nextId }

      PrimEffect eff -> case eff of
        EffectRefNew a ->
          let
            resA = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars moduleArities bound Nothing [] false false nextId a
            refIdent = "__local_ref_" <> show resA.nextId
            declStmt = GoAssign refIdent (boxGoExpr resA.expr resA.exprType)
            ifaceIdent = "__local_iface_" <> show resA.nextId
            ifaceStmt = GoRaw ("var " <> ifaceIdent <> " interface{} = " <> refIdent)
          in
            { stmts: resA.stmts <> StmtLeaf declStmt <> StmtLeaf ifaceStmt
            , expr: GoRaw ("gopurs_runtime.Any(&" <> ifaceIdent <> ")")
            , exprType: TypeValue, nextId: resA.nextId + 1
            }
        EffectRefRead a ->
          let
            resA = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars moduleArities bound Nothing [] false false nextId a
          in
            { stmts: resA.stmts
            , expr: GoRaw ("(*(" <> printGoExpr resA.expr <> ".PtrVal().(*interface{}))).(gopurs_runtime.Value)")
            , exprType: TypeValue, nextId: resA.nextId
            }
        EffectRefWrite ref val ->
          let
            resRef = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars moduleArities bound Nothing [] false false nextId ref
            resVal = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars moduleArities bound Nothing [] false false resRef.nextId val
            writeStmt = GoRaw ("*(" <> printGoExpr resRef.expr <> ".PtrVal().(*interface{})) = " <> printGoExpr (boxGoExpr resVal.expr resVal.exprType))
          in
            { stmts: resRef.stmts <> resVal.stmts <> StmtLeaf writeStmt
            , expr: boxGoExpr resVal.expr resVal.exprType
            , exprType: TypeValue, nextId: resVal.nextId
            }

      _ -> { stmts: StmtEmpty, expr: GoVar "gopurs_runtime.Value{}", exprType: TypeValue, nextId }



printTypeNode :: TypeNode -> String
printTypeNode (TNamed n) = n
printTypeNode (TFunc args ret) = 
  let 
    argsStr = String.joinWith ", " (map printTypeNode args)
    retStr = case ret of
      Nothing -> ""
      Just r -> " " <> printTypeNode r
  in "func(" <> argsStr <> ")" <> retStr
printTypeNode (TArray elem) = "[]" <> printTypeNode elem
printTypeNode (TMap k v) = "map[" <> printTypeNode k <> "]" <> printTypeNode v
printTypeNode (TUnknown s) = s

exprTypeToDummyTypeNode :: ExprType -> TypeNode
exprTypeToDummyTypeNode (Func args ret) = TFunc (map exprTypeToDummyTypeNode args) (Just (exprTypeToDummyTypeNode ret))
exprTypeToDummyTypeNode (Array elem) = TArray (exprTypeToDummyTypeNode elem)
exprTypeToDummyTypeNode (Record _) = TMap (TNamed "string") (TNamed "any")
exprTypeToDummyTypeNode _ = TNamed "any"

getTastReturnType :: ExprType -> Maybe ExprType
getTastReturnType (Func _ ret) = Just ret
getTastReturnType _ = Nothing

resolveNewtype :: Array DataDecl -> ExprType -> ExprType
resolveNewtype dataDecls (ADT fullName path args) = 
  let typeName = fullName
      mbDecl = Array.find (\d -> d.name == typeName || (Array.length path > 0 && d.name == fromMaybe "" (Array.last path))) dataDecls
  in case mbDecl of
       Just decl ->
         if Array.length decl.constructors == 1 then
           case Array.head decl.constructors of
             Just ctor ->
               if Array.length ctor.fields == 1 then
                 case Array.head ctor.fields of
                   Just fieldT -> resolveNewtype dataDecls fieldT
                   Nothing -> ADT fullName path args
               else ADT fullName path args
             Nothing -> ADT fullName path args
         else ADT fullName path args
       Nothing -> ADT fullName path args
resolveNewtype dataDecls (Func fArgs ret) = Func (map (resolveNewtype dataDecls) fArgs) (resolveNewtype dataDecls ret)
resolveNewtype dataDecls (Array elem) = Array (resolveNewtype dataDecls elem)
resolveNewtype dataDecls (Record row) = Record (resolveNewtype dataDecls row)
resolveNewtype dataDecls (Row fields tail) = Row (map (\(Tuple k v) -> Tuple k (resolveNewtype dataDecls v)) fields) (map (resolveNewtype dataDecls) tail)
resolveNewtype dataDecls (TypeApp c args) = TypeApp (resolveNewtype dataDecls c) (map (resolveNewtype dataDecls) args)
resolveNewtype dataDecls (ForAll vars body) = ForAll vars (resolveNewtype dataDecls body)
resolveNewtype dataDecls (ConstrainedType constraints body) = ConstrainedType (map (\(Tuple c a) -> Tuple c (map (resolveNewtype dataDecls) a)) constraints) (resolveNewtype dataDecls body)
resolveNewtype _ other = other

flattenFuncArgs :: ExprType -> Array ExprType
flattenFuncArgs (Func args ret) = args <> flattenFuncArgs ret
flattenFuncArgs _ = []

unwrapValueToFunc :: Array DataDecl -> TypeNode -> Maybe ExprType -> String -> Int -> Int -> String
unwrapValueToFunc dataDecls (TFunc args ret) mbTast valName depth _ = 
  let
    paramsArr = Array.mapWithIndex (\cidx atype -> "p" <> show depth <> "_" <> show cidx <> " " <> printTypeNode atype) args
    applyArgsArr = Array.mapWithIndex (\cidx atype ->
      case atype of
        TNamed "gopurs_runtime.Value" -> "p" <> show depth <> "_" <> show cidx
        TFunc _ _ -> 
          let wrapped = wrapReturn dataDecls atype (mbTast >>= \tast -> getTastArgType tast cidx) ("p" <> show depth <> "_" <> show cidx)
          in String.replaceAll (Pattern "\n") (Replacement "\n\t\t") wrapped
        _ -> "gopurs_runtime.Box(p" <> show depth <> "_" <> show cidx <> ")"
    ) args
    params = String.joinWith ", " paramsArr
    applyArgs = String.joinWith ", " applyArgsArr
    
    applyCall =
      if Array.length args == 1 then
        "gopurs_runtime.Apply(" <> valName <> ", " <> applyArgs <> ")"
      else if Array.length args > 1 then
        "gopurs_runtime.Apply" <> show (Array.length args) <> "(" <> valName <> ", " <> applyArgs <> ")"
      else
        "gopurs_runtime.Apply(" <> valName <> ", gopurs_runtime.Value{})"
        
  in case ret of
    Nothing ->
      "func(" <> params <> ") {\n\t\t" <> applyCall <> "\n\t}"
    Just (TArray elem) | printTypeNode elem /= "gopurs_runtime.Value" ->
      let elemType = printTypeNode elem
          retStr = printTypeNode (TArray elem)
      in "func(" <> params <> ") " <> retStr <> " {\n\t\tinner_res" <> show depth <> " := " <> applyCall <> "\n\t\tres_arr" <> show depth <> " := *(*[]gopurs_runtime.Value)(inner_res" <> show depth <> ".UnsafePtr)\n\t\tres_go" <> show depth <> " := make(" <> retStr <> ", len(res_arr" <> show depth <> "))\n\t\tfor i, v := range res_arr" <> show depth <> " { res_go" <> show depth <> "[i] = gopurs_runtime.Unbox[" <> elemType <> "](v) }\n\t\treturn res_go" <> show depth <> "\n\t}"
    Just (TNamed "any") -> "func(" <> params <> ") any {\n\t\treturn " <> applyCall <> "\n\t}"
    Just (TNamed "interface{}") -> "func(" <> params <> ") interface{} {\n\t\treturn " <> applyCall <> "\n\t}"
    Just (TNamed "gopurs_runtime.Value") -> "func(" <> params <> ") gopurs_runtime.Value {\n\t\treturn " <> applyCall <> "\n\t}"
    Just f@(TFunc _ _) ->
      let innerUnwrap = unwrapValueToFunc dataDecls f (mbTast >>= getTastReturnType) ("inner_res" <> show depth) (depth + 1) 0
      in "func(" <> params <> ") " <> printTypeNode f <> " {\n\t\tinner_res" <> show depth <> " := " <> applyCall <> "\n\t\treturn " <> innerUnwrap <> "\n\t}"
    Just r ->
      "func(" <> params <> ") " <> printTypeNode r <> " {\n\t\tinner_res" <> show depth <> " := " <> applyCall <> "\n\t\treturn gopurs_runtime.Unbox[" <> printTypeNode r <> "](inner_res" <> show depth <> ")\n\t}"
unwrapValueToFunc dataDecls (TNamed anyT) mbTast valName depth cidx | anyT == "any" || anyT == "interface{}" || anyT == "gopurs_runtime.Value" =
  let resolvedTast = map (resolveNewtype dataDecls) mbTast
  in case resolvedTast of
    Just (Record (Row fields _)) ->
      let
        fieldStr = Array.mapWithIndex (\i (Tuple fK fT) ->
            "\t\t\t\tres_map[\"" <> fK <> "\"] = " <> unwrapValueToFunc dataDecls (TNamed "any") (Just fT) ("_raw[\"" <> fK <> "\"]") (depth + 1) i
          ) fields
      in
        "func() map[string]any {\n\t\t\t_raw := gopurs_runtime.RecordToMap(" <> valName <> ")\n\t\t\tres_map := make(map[string]any)\n" <> String.joinWith "\n" fieldStr <> "\n\t\t\treturn res_map\n\t\t}()"
    Just f@(Func _ _) ->
      unwrapValueToFunc dataDecls (exprTypeToDummyTypeNode f) resolvedTast valName depth cidx
    _ -> "gopurs_runtime.Unbox[" <> anyT <> "](" <> valName <> ")"
unwrapValueToFunc _ t _ valName _ _ = "gopurs_runtime.Unbox[" <> printTypeNode t <> "](" <> valName <> ")"

wrapReturn :: Array DataDecl -> TypeNode -> Maybe ExprType -> String -> String
wrapReturn dataDecls (TFunc args ret) mbTast valName = 
  let
    innerT = ret
    argT = Array.head args
    
    genInner val innerWrap = 
      "gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {\n\t\t\t" <> val <> "\n\t\t\treturn " <> innerWrap <> "\n\t\t})"
      
    genInnerArg val innerWrap argUnwrap =
      "gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {\n\t\t\t" <> val <> "(" <> argUnwrap <> ")\n\t\t\treturn " <> innerWrap <> "\n\t\t})"
  in case argT of
    Nothing ->
      case innerT of
        Nothing -> genInner (valName <> "()") "gopurs_runtime.Value{}"
        Just r -> genInner ("inner_res := " <> valName <> "()") (wrapReturn dataDecls r (mbTast >>= getTastReturnType) "inner_res")
    Just a ->
      let 
        argUnwrap = case a of
          TNamed "any" -> "arg"
          TNamed "interface{}" -> "arg"
          TNamed "gopurs_runtime.Value" -> "arg"
          f@(TFunc _ _) -> String.replaceAll (Pattern "\n") (Replacement "\n\t\t\t") (unwrapValueToFunc dataDecls f (mbTast >>= \tast -> getTastArgType tast 0) "arg" 99 0)
          _ -> "gopurs_runtime.Unbox[" <> printTypeNode a <> "](arg)"
      in case innerT of
        Nothing -> genInnerArg valName "gopurs_runtime.Value{}" argUnwrap
        Just r -> genInnerArg ("inner_res := " <> valName) (wrapReturn dataDecls r (mbTast >>= getTastReturnType) "inner_res") argUnwrap
wrapReturn _ (TArray elem) _ valName | printTypeNode elem /= "gopurs_runtime.Value" = 
  "func() gopurs_runtime.Value {\n\t\t\tres_arr := make([]gopurs_runtime.Value, len(" <> valName <> "))\n\t\t\tfor i, v := range " <> valName <> " { res_arr[i] = gopurs_runtime.Box(v) }\n\t\t\treturn gopurs_runtime.Array(res_arr)\n\t\t}()"
wrapReturn dataDecls (TMap _ _) (Just (Record (Row fields _))) valName = 
  let
    fieldStr = Array.mapWithIndex (\i (Tuple fK fT) ->
        "\t\t\t\tres_map[\"" <> fK <> "\"] = " <> wrapReturn dataDecls (TNamed "any") (Just fT) ("_raw[\"" <> fK <> "\"]")
      ) fields
  in
    "func() gopurs_runtime.Value {\n\t\t\t_raw := " <> valName <> "\n\t\t\tres_map := make(map[string]gopurs_runtime.Value)\n" <> String.joinWith "\n" fieldStr <> "\n\t\t\treturn gopurs_runtime.Record(res_map)\n\t\t}()"
wrapReturn dataDecls (TMap _ _) mbTast valName = 
  let resolvedTast = map (resolveNewtype dataDecls) mbTast
      isOpaque = case resolvedTast of
                   Just (Record _) -> false
                   _ -> true
  in if isOpaque then
       "gopurs_runtime.Any(" <> valName <> ")"
     else
       "func() gopurs_runtime.Value {\n\t\t\tres_map := make(map[string]gopurs_runtime.Value)\n\t\t\tfor k, v := range " <> valName <> " { res_map[k] = gopurs_runtime.Box(v) }\n\t\t\treturn gopurs_runtime.Record(res_map)\n\t\t}()"
wrapReturn dataDecls (TNamed anyT) mbTast valName | anyT == "any" || anyT == "interface{}" || anyT == "gopurs_runtime.Value" =
  let resolvedTast = map (resolveNewtype dataDecls) mbTast
  in case resolvedTast of
    Just (Record (Row fields _)) ->
      let
        fieldStr = Array.mapWithIndex (\i (Tuple fK fT) ->
            "\t\t\t\tres_map[\"" <> fK <> "\"] = " <> wrapReturn dataDecls (TNamed "any") (Just fT) ("_raw[\"" <> fK <> "\"]")
          ) fields
      in
        "func() gopurs_runtime.Value {\n\t\t\t_raw := gopurs_runtime.RecordToMap(gopurs_runtime.Box(" <> valName <> "))\n\t\t\tres_map := make(map[string]gopurs_runtime.Value)\n" <> String.joinWith "\n" fieldStr <> "\n\t\t\treturn gopurs_runtime.Record(res_map)\n\t\t}()"
    Just f@(Func _ _) ->
      let
        fArgs = flattenFuncArgs f
        arity = Array.length fArgs
        
        genWrap args remaining depth =
           if remaining == 0 then
             let 
               castArgs = String.joinWith ", " (Array.replicate arity "any")
               castType = "func(" <> castArgs <> ") any"
               
               invokeArgs = Array.mapWithIndex (\i argT ->
                   let pName = "p" <> show (depth - arity + i)
                   in case argT of
                        Func _ _ -> 
                           let cbArgs = flattenFuncArgs argT
                               cbParams = Array.mapWithIndex (\ci _ -> "cb_arg" <> show ci) cbArgs
                               cbParamsDecl = String.joinWith ", " (map (\n -> n <> " any") cbParams)
                               applyChain = Array.foldl (\acc a -> "gopurs_runtime.Apply(" <> acc <> ", gopurs_runtime.Box(" <> a <> "))") pName cbParams
                           in "func(" <> cbParamsDecl <> ") any { return " <> applyChain <> " }"
                        _ -> "gopurs_runtime.Unbox[any](" <> pName <> ")"
                 ) fArgs
                 
               invokeCall = "fn(" <> String.joinWith ", " invokeArgs <> ")"
               fallbackChain = Array.foldl (\acc p -> "gopurs_runtime.Apply(" <> acc <> ", " <> p <> ")") ("(" <> valName <> ".(gopurs_runtime.Value))") (Array.mapWithIndex (\i _ -> "p" <> show (depth - arity + i)) fArgs)
               
             in 
             "func() gopurs_runtime.Value {\n\t\t\t\tif fn, ok := " <> valName <> ".(" <> castType <> "); ok {\n\t\t\t\t\treturn gopurs_runtime.Box(" <> invokeCall <> ")\n\t\t\t\t}\n\t\t\t\treturn " <> fallbackChain <> "\n\t\t\t}()"
           else
             let currArg = "p" <> show depth
             in "gopurs_runtime.Func(func(" <> currArg <> " gopurs_runtime.Value) gopurs_runtime.Value {\n\t\t\treturn " <> genWrap args (remaining - 1) (depth + 1) <> "\n\t\t})"
             
      in genWrap fArgs arity 0
    _ -> "gopurs_runtime.Box(" <> valName <> ")"
wrapReturn _ _ _ valName = "gopurs_runtime.Box(" <> valName <> ")"

printExprType :: ExprType -> String
printExprType = case _ of
  Int -> "Int"
  Number -> "Number"
  String -> "String"
  Char -> "Char"
  Boolean -> "Boolean"
  Unit -> "Unit"
  TypeLevelString s -> "(TypeLevelString " <> s <> ")"
  Array e -> "(Array " <> printExprType e <> ")"
  Func args ret -> "(Func [" <> String.joinWith ", " (map printExprType args) <> "] " <> printExprType ret <> ")"
  Record row -> "(Record " <> printExprType row <> ")"
  Row props tail -> 
    let tailStr = case tail of
          Nothing -> "Empty"
          Just t -> printExprType t
    in "(Row [" <> String.joinWith ", " (map (\(Tuple k v) -> k <> ": " <> printExprType v) props) <> "] " <> tailStr <> ")"
  TypeApp c args -> "(TypeApp " <> printExprType c <> " [" <> String.joinWith ", " (map printExprType args) <> "])"
  ForAll vars body -> "(ForAll [" <> String.joinWith ", " vars <> "] " <> printExprType body <> ")"
  ConstrainedType constraints body -> "(ConstrainedType " <> printExprType body <> ")"
  ADT fullName path args -> "(ADT " <> show path <> " [" <> String.joinWith ", " (map printExprType args) <> "])"
  TypeVar v -> "(TypeVar " <> v <> ")"
  Any -> "Any"

getTastArgType :: ExprType -> Int -> Maybe ExprType
getTastArgType (Func args _) i = Array.index args i
getTastArgType _ _ = Nothing

isStandardPursFunc :: TypeNode -> Boolean
isStandardPursFunc (TFunc args ret) =
  let 
    argsAny = Array.all (\a -> printTypeNode a == "any" || printTypeNode a == "interface{}" || printTypeNode a == "gopurs_runtime.Value") args
    retStr = case ret of
      Nothing -> ""
      Just r -> printTypeNode r
  in if Array.length args == 0 then
       retStr == "" || retStr == "bool" || retStr == "int" || retStr == "int64" || retStr == "string" || retStr == "float64" || retStr == "gopurs_runtime.Value" || retStr == "any" || retStr == "interface{}"
     else if argsAny then
       case ret of
         Nothing -> true
         Just r@(TFunc _ _) -> isStandardPursFunc r
         Just r -> retStr == "any" || retStr == "interface{}" || retStr == "gopurs_runtime.Value"
     else
       false
isStandardPursFunc _ = false

generateWrapperFunc :: Array DataDecl -> FfiDecl -> Maybe ExprType -> String
generateWrapperFunc dataDecls d mbTast = 
  let tastComment = case mbTast of
        Just tast -> "// TAST: " <> printExprType tast <> "\n"
        Nothing -> "// TAST: Unknown\n"
  in tastComment <>
  if d.isVar then
    "gopurs_runtime.Box(" <> d.name <> ")"
  else
    let
      arity = Array.length d.args
      funcConstructor = if arity > 1 then "Func" <> show arity else "Func"
      
      boxedArgs = 
        if arity == 0 then
          "_ gopurs_runtime.Value"
        else if arity == 1 then
          "arg0 gopurs_runtime.Value"
        else
          String.joinWith ", " (Array.mapWithIndex (\i _ -> "arg" <> show i <> " gopurs_runtime.Value") d.args)
          
      substituteGeneric :: TypeNode -> TypeNode
      substituteGeneric t = 
        let s = printTypeNode t 
            res = Array.foldl (\acc tp -> String.replaceAll (Pattern ("\\b" <> tp <> "\\b")) (Replacement "gopurs_runtime.Value") acc) s d.typeParams
        in TUnknown res 

      newLines = []
      
      callFunc = if Array.length d.typeParams > 0 then
          d.name <> "[" <> String.joinWith ", " (map (const "gopurs_runtime.Value") d.typeParams) <> "]"
        else d.name

      processArg i t =
        let typStr = printTypeNode t 
            elemType = String.drop 2 typStr
            tastArg = mbTast >>= \tast -> getTastArgType tast i
        in case t of
          TFunc _ _ -> 
            let 
              isOpaque = case tastArg of
                           Just (ADT _ _ _) -> true
                           Just Any -> true
                           Just (TypeVar _) -> true
                           _ -> false
            in if isOpaque && not (isStandardPursFunc t) then
                 [ "\tgo_arg" <> show i <> " := (*(*any)(arg" <> show i <> ".UnsafePtr)).(" <> typStr <> ")" ]
               else
                 let unwrapStr = unwrapValueToFunc dataDecls t tastArg ("arg" <> show i) 0 0
                     indented = String.replaceAll (Pattern "\n") (Replacement "\n\t") unwrapStr
                 in [ "\tgo_arg" <> show i <> " := " <> indented ]
          TArray _ | typStr /= "[]gopurs_runtime.Value" ->
            let 
              et = if elemType == "any" then "interface{}" else elemType
            in
              [ "\targ" <> show i <> "_arr := *(*[]gopurs_runtime.Value)(arg" <> show i <> ".UnsafePtr)"
              , "\tgo_arg" <> show i <> " := make(" <> typStr <> ", len(arg" <> show i <> "_arr))"
              , if et == "interface{}" then
                  "\tfor i, v := range arg" <> show i <> "_arr { go_arg" <> show i <> "[i] = v }"
                else
                  "\tfor i, v := range arg" <> show i <> "_arr { go_arg" <> show i <> "[i] = gopurs_runtime.Unbox[" <> et <> "](v) }"
              ]
          TNamed "any" -> [ "\tgo_arg" <> show i <> " := arg" <> show i ]
          TNamed "interface{}" -> [ "\tgo_arg" <> show i <> " := arg" <> show i ]
          TNamed "gopurs_runtime.Value" -> [ "\tgo_arg" <> show i <> " := arg" <> show i ]
          TMap _ _ ->
            let 
              et = String.drop (String.indexOf (Pattern "]") typStr # fromMaybe 0 # add 1) typStr
              resolvedTast = map (resolveNewtype dataDecls) tastArg
              isOpaque = case resolvedTast of
                           Just (Record _) -> false
                           _ -> true
            in if isOpaque then
                 [ "\tgo_arg" <> show i <> " := gopurs_runtime.UnboxObject(arg" <> show i <> ")" ]
               else if et == "any" || et == "interface{}" then
                 [ "\targ" <> show i <> "_map := gopurs_runtime.RecordToMap(arg" <> show i <> ")"
                 , "\tgo_arg" <> show i <> " := make(" <> typStr <> ")"
                 , "\tfor k, v := range arg" <> show i <> "_map { go_arg" <> show i <> "[k] = v }"
                 ]
               else
                 [ "\tgo_arg" <> show i <> " := arg" <> show i <> ".PtrVal().(" <> typStr <> ")" ]
          _ -> [ "\tgo_arg" <> show i <> " := gopurs_runtime.Unbox[" <> typStr <> "](arg" <> show i <> ")" ]
          
      argsCode = Array.concat (Array.mapWithIndex processArg d.args)
      callArgs = String.joinWith ", " (Array.mapWithIndex (\i _ -> "go_arg" <> show i) d.args)
      
      retCode = case d.ret of
        Nothing -> 
          [ "\t" <> callFunc <> "(" <> callArgs <> ")"
          , "\treturn gopurs_runtime.Value{}"
          ]
        Just r ->
          let wrapCode = wrapReturn dataDecls r (mbTast >>= getTastReturnType) "go_res"
              indentedWrap = String.replaceAll (Pattern "\n") (Replacement "\n\t") wrapCode
          in
          [ "\tgo_res := " <> callFunc <> "(" <> callArgs <> ")"
          , "\treturn " <> indentedWrap
          ]
             
      fullCode = "gopurs_runtime." <> funcConstructor <> "(func(" <> boxedArgs <> ") gopurs_runtime.Value {\n" <>
                 String.joinWith "\n" argsCode <> "\n" <>
                 String.joinWith "\n" retCode <> "\n" <>
                 "})"
    in fullCode

generateFfiBridge :: String -> Array DataDecl -> Array FfiDecl -> Array (Tuple Ident (Maybe ExprType)) -> String
generateFfiBridge modNameStr dataDecls decls foreigns = 
  String.joinWith "\n" (map genBridge foreigns)
  where
  genBridge (Tuple ident mbTast) = 
    let 
      pursName = unwrap ident
      sanitized = sanitizeName pursName
      capName = capitalize sanitized
      exportName = "_Gopurs_" <> modNameStr <> "_" <> capName
      
      fallback1 = modNameStr <> "_" <> capitalize pursName
      fallback2 = fallback1 <> "_"
      
      findDecl n = Array.find (\d -> d.name == n) decls
      
      match = case findDecl fallback1 of
                Just d -> Just d
                Nothing -> findDecl fallback2
    in 
      case match of
        Nothing -> 
          "var " <> exportName <> " = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value { panic(\"FFI not implemented: " <> pursName <> "\"); return gopurs_runtime.Value{} })"
        Just d ->
          "var " <> exportName <> " = " <> generateWrapperFunc dataDecls d mbTast
