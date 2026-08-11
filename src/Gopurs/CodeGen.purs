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
coerceGoExpr expr from TypeValue = boxGoExpr expr from
coerceGoExpr expr TypeValue to = unboxGoExpr expr TypeValue to
coerceGoExpr expr from to = unboxGoExpr (boxGoExpr expr from) TypeValue to

boxGoExpr :: GoExpr -> GoType -> GoExpr
boxGoExpr expr TypeValue = expr
boxGoExpr expr TypeInt64 = GoCall (GoSelector (GoVar "gopurs_runtime") "Int") [ expr ]
boxGoExpr expr TypeFloat64 = GoCall (GoSelector (GoVar "gopurs_runtime") "Float") [ expr ]
boxGoExpr expr TypeString = GoCall (GoSelector (GoVar "gopurs_runtime") "Str") [ expr ]
boxGoExpr expr TypeBool = GoCall (GoSelector (GoVar "gopurs_runtime") "Bool") [ expr ]
boxGoExpr expr (TypeStructPointer baseStructName _) = GoRaw ("gopurs_runtime.Value{Type: 9, IntVal: " <> hashString baseStructName <> ", UnsafePtr: unsafe.Pointer(" <> printGoExpr expr <> ")}")
boxGoExpr expr (TypeRecord _) = expr
boxGoExpr expr (TypeInterface _) = expr
boxGoExpr expr (TypeNativeArray TypeValue) = GoCall (GoSelector (GoVar "gopurs_runtime") "Array") [ expr ]
boxGoExpr expr (TypeNativeArray inner) = GoRaw ("func() gopurs_runtime.Value {\n\t\t\t\t\tarr := " <> printGoExpr expr <> "\n\t\t\t\t\tboxed := make([]gopurs_runtime.Value, len(arr))\n\t\t\t\t\tfor i, v := range arr { boxed[i] = " <> printGoExpr (boxGoExpr (GoVar "v") inner) <> " }\n\t\t\t\t\treturn gopurs_runtime.Array(boxed)\n\t\t\t\t}()")
boxGoExpr expr (TypeGenericParam _) = expr
boxGoExpr expr (TypeFunc _ _) = expr

mangleType :: Map.Map String { ctorName :: String, arity :: Int } -> String -> ExprType -> String
mangleType ptrPaths modNameStr t = 
  let
    typeStr = goTypeToStr (exprTypeToGoType ptrPaths modNameStr t)
    typeStrNoPkg = String.replaceAll (Pattern "pkg_") (Replacement "") typeStr
    typeStrSafe = String.replaceAll (Pattern ".") (Replacement "_") typeStrNoPkg
    typeStrSafe2 = String.replaceAll (Pattern "[]") (Replacement "arr") typeStrSafe
    typeStrSafe3 = String.replaceAll (Pattern "*") (Replacement "ptr") typeStrSafe2
    typeStrSafe4 = String.replaceAll (Pattern "[") (Replacement "_") typeStrSafe3
    typeStrSafe5 = String.replaceAll (Pattern "]") (Replacement "_") typeStrSafe4
    typeStrSafe6 = String.replaceAll (Pattern ",") (Replacement "_") typeStrSafe5
    cleanType = String.replaceAll (Pattern " ") (Replacement "_") typeStrSafe6
  in cleanType <> "_" <> hashString (Monomorphize.mangleType t)

exprTypeToGoType :: Map.Map String { ctorName :: String, arity :: Int } -> String -> ExprType -> GoType
exprTypeToGoType _ _ Int = TypeInt64
exprTypeToGoType _ _ Number = TypeFloat64
exprTypeToGoType _ _ String = TypeString
exprTypeToGoType _ _ Char = TypeString
exprTypeToGoType _ _ Boolean = TypeBool
exprTypeToGoType ptrPaths modNameStr (Array ty) = TypeNativeArray (exprTypeToGoType ptrPaths modNameStr ty)
exprTypeToGoType ptrPaths modNameStr (Record (Row fields _)) = TypeRecord (map (\(Tuple k v) -> Tuple k (exprTypeToGoType ptrPaths modNameStr v)) (Array.sortBy (comparing \(Tuple k _) -> k) fields))
exprTypeToGoType ptrPaths modNameStr (Record _) = TypeValue
exprTypeToGoType ptrPaths modNameStr (ADT fullName path args) = case Map.lookup fullName ptrPaths of
  Just info -> 
    let 
      ctorName = info.ctorName
      pkgNameStr = String.replaceAll (Pattern ".") (Replacement "_") (String.joinWith "." (Array.slice 0 (Array.length path - 1) path))
      pkgPrefix = if pkgNameStr /= modNameStr then "pkg_" <> pkgNameStr <> "." else ""
      baseStructName = "Data_" <> pkgNameStr <> "_" <> sanitizeName ctorName
      monoStructName = "Constructor_" <> sanitizeName ctorName
      typeArgsMapped = map (exprTypeToGoType ptrPaths modNameStr) args
      paddedTypeArgs = typeArgsMapped <> Array.replicate (info.arity - Array.length typeArgsMapped) TypeValue
      typeArgsStr = if Array.length paddedTypeArgs > 0 then "[" <> String.joinWith ", " (map goTypeToStr paddedTypeArgs) <> "]" else ""
    in TypeStructPointer baseStructName (pkgPrefix <> monoStructName <> typeArgsStr)
  Nothing -> TypeValue
exprTypeToGoType _ _ (TypeVar v) = TypeValue
exprTypeToGoType _ _ _ = TypeValue

exprTypeToGenericGoType :: Map.Map String { ctorName :: String, arity :: Int } -> Array String -> String -> ExprType -> GoType
exprTypeToGenericGoType ptrPaths typeVars modNameStr (Record (Row fields _)) = TypeRecord (map (\(Tuple k v) -> Tuple k (exprTypeToGenericGoType ptrPaths typeVars modNameStr v)) (Array.sortBy (comparing \(Tuple k _) -> k) fields))
exprTypeToGenericGoType _ _ _ (Record _) = TypeValue
exprTypeToGenericGoType _ typeVars _ (TypeVar v) | Array.elem v typeVars = TypeGenericParam v
exprTypeToGenericGoType ptrPaths _ modNameStr ty = exprTypeToGoType ptrPaths modNameStr ty

structFieldGoType :: Map.Map String { ctorName :: String, arity :: Int } -> Array String -> String -> ExprType -> GoType
structFieldGoType ptrPaths typeVars modStr ty = case exprTypeToGenericGoType ptrPaths typeVars modStr ty of
  TypeInterface _ -> TypeValue
  other -> other


unboxGoExpr :: GoExpr -> GoType -> GoType -> GoExpr
unboxGoExpr expr currentType desiredType =
  if currentType == desiredType then expr
  else if currentType /= TypeValue then
    unboxGoExpr (boxGoExpr expr currentType) TypeValue desiredType
  else case desiredType of
    TypeValue -> boxGoExpr expr currentType
    (TypeRecord _) -> boxGoExpr expr currentType
    TypeInt64 -> GoSelector expr "IntVal"
    TypeFloat64 -> GoCall (GoSelector expr "FloatVal") []
    TypeString -> GoCall (GoSelector expr "StrVal") []
    TypeBool -> GoBinOp "!=" (GoSelector expr "IntVal") (GoInt 0)
    (TypeStructPointer _ fullPath) -> GoCall (GoRaw ("(*" <> fullPath <> ")")) [ GoSelector expr "UnsafePtr" ]
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

globalReusedVars :: Ref.Ref (Set.Set String)
globalReusedVars = unsafePerformEffect (Ref.new Set.empty)

translate :: Map.Map String { ctorName :: String, arity :: Int } -> Set.Set String -> Map.Map String String -> Set.Set ExprType -> Set.Set String -> Map.Map String { typeVars :: Array String, fieldTypes :: Array ExprType } -> Map.Map String ExprType -> InstantiationMap -> Array (Array String) -> BackendModule -> String
translate pointerAdtPaths pointerAdtNodes pointerAdtLeaves adtTypes elidedCtors ctorTypes globalTypes rawInstantiations importsArray mod =
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
    
    instantiations = Map.mapMaybe (\set ->
        let safeSet = Set.filter isSafeType set
        in if Set.isEmpty safeSet then Nothing else Just safeSet
      ) rawInstantiations

    _ = unsafePerformEffect (Console.log ("Translating module " <> modNameStr))
    helpersRef = unsafePerformEffect do
      let
        structDecls = Array.concatMap (\decl ->
            Array.concatMap (\ctor ->
              let
                fieldTypes = ctor.fieldTypes
                goFieldTypes = map (structFieldGoType pointerAdtPaths decl.typeVars modNameStr) fieldTypes
                structName = "Constructor_" <> sanitizeName ctor.constructorName
                
                typeParams = if Array.length decl.typeVars > 0 then "[" <> String.joinWith ", " (map (\v -> "T_" <> sanitizeName v <> " any") decl.typeVars) <> "]" else ""
                
                fieldsStr = Array.cons "Rc uint32" (Array.mapWithIndex (\i ty -> "V" <> show i <> " " <> goTypeToStr ty) goFieldTypes)
                structDecl = "type " <> structName <> typeParams <> " struct {\n\t" <> String.joinWith "\n\t" fieldsStr <> "\n}\n"
              in
                [ structDecl ]
            ) decl.constructors
          ) mod.dataDecls

      Ref.new { decls: [], rawDecls: structDecls, elidedCtors, ctorTypes, pointerAdtPaths, pointerAdtNodes, pointerAdtLeaves, globalTypes, globalId: 0 }

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

    getExprType :: TcoExpr -> ExprType
    getExprType (TcoExpr _ (Typed ty _)) = ty
    getExprType _ = Any -- fallback

    setTcoExprType :: ExprType -> TcoExpr -> TcoExpr
    setTcoExprType ty (TcoExpr a (Typed _ inner)) = TcoExpr a (Typed ty inner)
    setTcoExprType _ expr = expr

    tcoBindingsExpanded = map
      (\group -> group { bindings = Array.concatMap expandBind group.bindings })
      tcoBindings

    expandBind :: Tuple Ident TcoExpr -> Array (Tuple Ident TcoExpr)
    expandBind (Tuple id@(Ident name) val) =
      let qual = modNameStrOrig <> "." <> name
          instsMap = map Set.toUnfoldable instantiations
          baseVal = substituteAst instsMap (mangleType pointerAdtPaths modNameStr) val
      in case Map.lookup qual instantiations of
           Just concretes ->
             let genericType = fromMaybe (getExprType val) (Map.lookup qual globalTypes)
                 concreteArr = Set.toUnfoldable concretes :: Array ExprType
                 res = Tuple id baseVal `Array.cons` Array.mapMaybe (\concrete ->
                  let subst = unify genericType concrete Map.empty
                  in Just $
                       let mangledName = name <> "__" <> mangleType pointerAdtPaths modNameStr concrete
                           mangledVal = mapTcoExprTypes (substituteExprType subst) val
                           mangledVal' = substituteAst instsMap (mangleType pointerAdtPaths modNameStr) mangledVal
                       in Tuple (Ident mangledName) (setTcoExprType concrete mangledVal')
                 ) concreteArr
             in res
           Nothing -> [ Tuple id baseVal ]

    extractFuncType :: TcoExpr -> Maybe { fArgs :: Array ExprType, fRet :: ExprType }
    extractFuncType (TcoExpr _ (Typed ty inner)) =
      let
        flattenFuncType acc (Func args ret) = flattenFuncType (acc <> args) ret
        flattenFuncType acc ret = { fArgs: acc, fRet: ret }
        
        getFunc (Func a r) = Just (flattenFuncType a r)
        getFunc _ = extractFuncType inner
      in getFunc ty
    extractFuncType _ = Nothing

    unwrapFunc :: Array (Tuple Ident TcoExpr) -> Array (Tuple String { fullName :: String, fArgs :: Array GoType, fRet :: GoType, arity :: Int })
    unwrapFunc binds =
      Array.concatMap
        ( \(Tuple (Ident name) val) ->
            case extractUncurriedAbs val of
              Just { args, body, fvs } ->
                let
                  typeSig = extractFuncType val
                  fArgsGo = case typeSig of
                    Just { fArgs } -> map (exprTypeToGoType pointerAdtPaths modNameStr) (Array.take (Array.length args) fArgs)
                    Nothing -> Array.replicate (Array.length args) TypeValue
                  fRetGo = case typeSig of
                    Just { fArgs, fRet } -> 
                      if Array.length args < Array.length fArgs then
                        TypeValue
                      else
                        exprTypeToGoType pointerAdtPaths modNameStr fRet
                    Nothing -> TypeValue
                  fullName = "Call_" <> sanitizeName name
                in
                  [ Tuple (sanitizeName name) { fullName, fArgs: fArgsGo, fRet: fRetGo, arity: Array.length args } ]
              Nothing ->
                let fullName = "Call_" <> sanitizeName name
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
              Nothing -> map (\(Tuple (Ident name) _) -> Tuple (sanitizeName name) { fullName: "Call_" <> sanitizeName name, fArgs: [], fRet: TypeValue, arity: 0 }) group.bindings
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
                      mutRecBinds = traverse (\(Tuple (Ident name) val) -> map (\abs -> { ident: sanitizeName name, args: abs.args, body: abs.body, fvs: abs.fvs }) (extractUncurriedAbs val)) binds
                    in
                      case mutRecBinds of
                        Just fns ->
                          let
                            loopCtxs = map (\fn -> { ident: fn.ident, params: fn.args, loopParams: map (\p -> p <> "_loop") fn.args }) fns

                            fnWrapperStmts = map
                              ( \fn ->
                                let
                                  paramsWithTypes = case Map.lookup fn.ident moduleArities of
                                    Just { fArgs } -> Array.zip fn.args (fArgs <> Array.replicate (Array.length fn.args - Array.length fArgs) TypeValue)
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
                                      let callFuncDecl = "func Call_" <> goName <> "(" <> goParams <> ") " <> goTypeToStr expectedRetType <> " {\n" <> printGoExpr funcBody <> "\n}"
                                      Ref.modify_ (\r -> r { rawDecls = Array.snoc r.rawDecls callFuncDecl }) helpersRef
                                      let wrapperParams = map (\(Tuple p _) -> p <> "_box") paramsWithTypes
                                      let callExpr = GoCall (GoVar ("Call_" <> goName)) (map (\(Tuple p goT) -> unboxGoExpr (GoVar (p <> "_box")) TypeValue goT) paramsWithTypes)
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
                                  { identifier: goName, expression: funcExpr, goType: TypeValue }
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
                                  [ { identifier: sanitizeName name, expression: wrapInStmts [] res.stmts (boxGoExpr res.expr res.exprType), goType: TypeValue } ]
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

    parts = String.split (Pattern "pkg_") declsStr
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
      { packageName: modNameStr
      , imports: goImports
      , decls: allDeclsAst
      , rawDecls: helpers.rawDecls
      , foreigns: map (\(Tuple (Ident name) type_) -> { pursName: sanitizeName name, goName: "_Gopurs_" <> capitalize (sanitizeName name), exprType: type_ }) (Map.toUnfoldable mod.foreign)
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

getExprType :: TcoExpr -> ExprType
getExprType (TcoExpr _ syn) = case syn of
  Typed t _ -> t
  _ -> Any

getExprTypeArity :: ExprType -> Int
getExprTypeArity (Func args ret) = Array.length args + getExprTypeArity ret
getExprTypeArity _ = 0

executeIfOpaque :: forall a. BackendSyntax a -> GoExpr -> GoExpr

executeIfOpaque expr goExpr =
  if isEffectNode expr then goExpr
  else GoCall (GoSelector (GoVar "gopurs_runtime") "Apply") [ goExpr, GoRaw "gopurs_runtime.Value{}" ]


translateExprImpl :: Ref { decls :: Array GoDecl, rawDecls :: Array String, elidedCtors :: Set.Set String, ctorTypes :: Map String { typeVars :: Array String, fieldTypes :: Array ExprType }, pointerAdtPaths :: Map String { ctorName :: String, arity :: Int }, pointerAdtNodes :: Set String, pointerAdtLeaves :: Map String String, globalTypes :: Map.Map String ExprType, globalId :: Int } -> Int -> String -> Array String -> Map String { fullName :: String, fArgs :: Array GoType, fRet :: GoType, arity :: Int } -> Map String { name :: String, goType :: GoType } -> Maybe String -> Array { ident :: String, params :: Array String, loopParams :: Array String, goTypes :: Array GoType } -> Boolean -> Int -> TcoExpr -> { stmts :: StmtTree, expr :: GoExpr, exprType :: GoType, nextId :: Int }
translateExprImpl helpersRef depth modNameStr recVars moduleArities bound tcoIdent loopCtx isTail nextId tcoExpr =
  translateExprImpl_ helpersRef depth modNameStr recVars moduleArities bound tcoIdent loopCtx isTail false nextId tcoExpr

translateExprImpl_ :: Ref { decls :: Array GoDecl, rawDecls :: Array String, elidedCtors :: Set.Set String, ctorTypes :: Map String { typeVars :: Array String, fieldTypes :: Array ExprType }, pointerAdtPaths :: Map String { ctorName :: String, arity :: Int }, pointerAdtNodes :: Set String, pointerAdtLeaves :: Map String String, globalTypes :: Map.Map String ExprType, globalId :: Int } -> Int -> String -> Array String -> Map String { fullName :: String, fArgs :: Array GoType, fRet :: GoType, arity :: Int } -> Map String { name :: String, goType :: GoType } -> Maybe String -> Array { ident :: String, params :: Array String, loopParams :: Array String, goTypes :: Array GoType } -> Boolean -> Boolean -> Int -> TcoExpr -> { stmts :: StmtTree, expr :: GoExpr, exprType :: GoType, nextId :: Int }
translateExprImpl_ helpersRef depth modNameStr recVars moduleArities bound tcoIdent loopCtx isTail inEffectBlock nextId tcoExpr@(TcoExpr tcoAnalysis expr) =
  let
    _ = unsafePerformEffect (if depth == 0 then Ref.write Set.empty globalReusedVars else pure unit)
    elidedCtors = (unsafePerformEffect (Ref.read helpersRef)).elidedCtors
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
        let gId = unsafePerformEffect do
              curr <- Ref.read helpersRef
              Ref.modify_ (\r -> r { globalId = r.globalId + 1 }) helpersRef
              pure curr.globalId
        let helperName = "__helper_" <> show gId
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
          res = translateExprImpl_ helpersRef depth modNameStr recVars moduleArities bound tcoIdent loopCtx isTail inEffectBlock nextId a
          expectedGoType = exprTypeToGoType (unsafePerformEffect (Ref.read helpersRef)).pointerAdtPaths modNameStr type_
        in
          { stmts: res.stmts, expr: coerceGoExpr res.expr res.exprType expectedGoType, exprType: expectedGoType, nextId: res.nextId }
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
          accProps = foldl
            ( \acc (Prop key val) ->
                let
                  resVal = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars moduleArities bound Nothing [] false false acc.nextId val
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
                        Just (ModuleName mod) | not isLocal -> "pkg_" <> String.replaceAll (Pattern ".") (Replacement "_") mod <> "."
                        _ -> ""
                      fromModuleArities = if isLocal then Map.lookup name moduleArities else Nothing
                      fromTypeSig = case getFuncType (unwrapExpr flatFn) of
                        Just { fArgs, fRet } ->
                          Just { fullName: modPrefix <> "Call_" <> sanitizeName name, fArgs: map (exprTypeToGoType (unsafePerformEffect (Ref.read helpersRef)).pointerAdtPaths modNameStr) fArgs, fRet: exprTypeToGoType (unsafePerformEffect (Ref.read helpersRef)).pointerAdtPaths modNameStr fRet, arity: Array.length fArgs }
                        Nothing -> Nothing
                      
                      entry = case fromTypeSig of
                        Just e -> Just e
                        Nothing -> fromModuleArities
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
                          
                        "filter" ->
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
                            
                      finalExpr = buildApp iifeExpr accArgsRemaining
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
                            finalExpr = buildApp (boxGoExpr callExpr fRet) accArgsRemainingBoxed
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
                                    { stmts: acc.stmts <> argRes.stmts, exprs: Array.snoc acc.exprs argRes.expr, exprTypes: Array.snoc acc.exprTypes argRes.exprType, nextId: argRes.nextId }
                              )
                              { stmts: resFn.stmts, exprs: [], exprTypes: [], nextId: resFn.nextId }
                              flatArgs
                              
                            finalExpr = case resFn.exprType of
                              TypeFunc fArgs fRet | Array.length fArgs == Array.length flatArgs ->
                                let
                                  callArgs = Array.mapWithIndex (\i expected ->
                                      let arg = fromMaybe (GoRaw "nil") (Array.index accArgs.exprs i)
                                          actual = fromMaybe TypeValue (Array.index accArgs.exprTypes i)
                                      in unboxGoExpr arg actual expected
                                    ) fArgs
                                in boxGoExpr (GoCall resFn.expr callArgs) fRet
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
          getVar :: BackendSyntax TcoExpr -> Maybe { mbMod :: Maybe ModuleName, name :: String }
          getVar (Typed _ inner) = getVar (unwrapTcoExpr inner)
          getVar (Var (Qualified mbMod (Ident name))) = Just { mbMod, name }
          getVar _ = Nothing

          mbIntrinsic = case getVar (unwrapTcoExpr fn) of
            Just { name: "arrayMap" } ->
              if Array.length args >= 2 then Just "arrayMap" else Nothing
            Just { name: "foldlArray" } ->
              if Array.length args >= 3 then Just "foldlArray" else Nothing
            Just { mbMod, name: "filter" } | mbMod == Just (ModuleName "Data.Array") || (mbMod == Nothing && modNameStr == "Data.Array") ->
              if Array.length args >= 2 then Just "filter" else Nothing
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
                        { stmts: acc.stmts <> argRes.stmts, exprs: Array.snoc acc.exprs (boxGoExpr argRes.expr argRes.exprType), exprType: TypeValue, nextId: argRes.nextId }
                  )
                  { stmts: StmtEmpty, exprs: [], exprType: TypeValue, nextId }
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
                    
                  "filter" ->
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
              in
                { stmts: accArgs.stmts, expr: iifeExpr, exprType: TypeValue, nextId: accArgs.nextId }
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
          paramsWithTypes = case getExprType tcoExpr of
            Func fArgs _ -> Array.zipWith (\(Tuple mbI lvl) goType -> Tuple (localId mbI lvl) goType) args (map (exprTypeToGoType (unsafePerformEffect (Ref.read helpersRef)).pointerAdtPaths modNameStr) fArgs <> Array.replicate (Array.length args - Array.length fArgs) TypeValue)
            _ -> map (\(Tuple mbI lvl) -> Tuple (localId mbI lvl) TypeValue) args

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
          paramsWithTypes = case getExprType tcoExpr of
            Func fArgs _ -> Array.zipWith (\(Tuple mbI lvl) goType -> Tuple (localId mbI lvl) goType) args (map (exprTypeToGoType (unsafePerformEffect (Ref.read helpersRef)).pointerAdtPaths modNameStr) fArgs <> Array.replicate (Array.length args - Array.length fArgs) TypeValue)
            _ -> map (\(Tuple mbI lvl) -> Tuple (localId mbI lvl) TypeValue) args
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
            ( \acc (Tuple (Ident ident) val) ->
                let
                  oldName = localId (Just (Ident ident)) lvl
                  gId = unsafePerformEffect do
                    curr <- Ref.read helpersRef
                    Ref.modify_ (\r -> r { globalId = r.globalId + 1 }) helpersRef
                    pure curr.globalId
                  newName = oldName <> "_" <> show acc.nextId <> "_" <> show gId
                  expectedGoType = exprTypeToGoType (unsafePerformEffect (Ref.read helpersRef)).pointerAdtPaths modNameStr (getExprType val)
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
                          in Tuple idStr (exprTypeToGoType (unsafePerformEffect (Ref.read helpersRef)).pointerAdtPaths modNameStr t)
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
                            in Tuple idStr (exprTypeToGoType (unsafePerformEffect (Ref.read helpersRef)).pointerAdtPaths modNameStr t)
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
                
                resBodyOuter = translateExprImpl_ helpersRef (depth + 1) modNameStr combinedRecVars moduleArities allocRes.newBound Nothing loopCtx isTail false nextId' body
              in
                { stmts: foldMap StmtLeaf declStmts <> foldMap StmtLeaf fnWrapperStmts <> resBodyOuter.stmts, expr: resBodyOuter.expr, exprType: resBodyOuter.exprType, nextId: resBodyOuter.nextId }
            
            Nothing ->
              let
                accBindings = foldl
                  ( \acc (Tuple (Tuple (Ident ident) val) alloc) ->
                      let
                        res = translateExprImpl_ helpersRef (depth + 1) modNameStr combinedRecVars moduleArities allocRes.newBound Nothing [] false false acc.nextId val
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
            GetProp prop ->
              case resObj.exprType of
                TypeRecord fields -> { stmts: resObj.stmts, expr: GoRecordAccess (boxGoExpr resObj.expr resObj.exprType) prop, exprType: TypeValue, nextId: resObj.nextId }
                TypeValue -> { stmts: resObj.stmts, expr: GoRecordAccess (boxGoExpr resObj.expr resObj.exprType) prop, exprType: TypeValue, nextId: resObj.nextId }
                _ -> { stmts: resObj.stmts, expr: GoRecordAccess (boxGoExpr resObj.expr resObj.exprType) prop, exprType: TypeValue, nextId: resObj.nextId }
            GetIndex idx -> { stmts: resObj.stmts, expr: GoCall (GoSelector (GoVar "gopurs_runtime") "ArrayAccess") [ (boxGoExpr resObj.expr resObj.exprType), GoInt idx ], exprType: TypeValue, nextId: resObj.nextId }
            GetCtorField (Qualified mbMod _) _ _ (Ident ctorName) _ idx ->
              let
                structName = getStructName modNameStr mbMod ctorName
                modPart = case mbMod of
                  Just (ModuleName mn) -> String.replaceAll (Pattern ".") (Replacement "_") mn
                  Nothing -> modNameStr
                key = modPart <> "." <> ctorName
                helpers = unsafePerformEffect (Ref.read helpersRef)
              in
                if Set.member structName elidedCtors then
                  { stmts: resObj.stmts, expr: coerceGoExpr resObj.expr resObj.exprType TypeValue, exprType: TypeValue, nextId: resObj.nextId }
                else
                  let
                    fieldTypes = fromMaybe [] (map _.fieldTypes (Map.lookup key helpers.ctorTypes))
                    monoStructName = "Constructor_" <> sanitizeName ctorName
                    
                    pkgPrefix = case mbMod of
                      Just (ModuleName mod) | String.replaceAll (Pattern ".") (Replacement "_") mod /= modNameStr -> "pkg_" <> String.replaceAll (Pattern ".") (Replacement "_") mod <> "."
                      _ -> ""
                      
                    expectedType = case Array.index fieldTypes idx of
                      Just ty -> exprTypeToGoType (unsafePerformEffect (Ref.read helpersRef)).pointerAdtPaths modNameStr ty
                      Nothing -> TypeValue
                      
                    typeArgs = case getExprType obj of
                      ADT _ _ tArgs -> map (exprTypeToGoType (unsafePerformEffect (Ref.read helpersRef)).pointerAdtPaths modNameStr) tArgs
                      _ -> case Map.lookup key helpers.ctorTypes of
                        Just ctorInfo -> map (const TypeValue) ctorInfo.typeVars
                        Nothing -> []
                      
                    exprAccess = GoConstructorAccess (boxGoExpr resObj.expr resObj.exprType) (pkgPrefix <> monoStructName) typeArgs idx
                  in
                    { stmts: resObj.stmts, expr: coerceGoExpr exprAccess expectedType TypeValue, exprType: TypeValue, nextId: resObj.nextId }

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
          structName = "Constructor_" <> sanitizeName name
          baseStructName = getBaseStructName modNameStr Nothing name
          key = modNameStr <> "." <> name
          helpers = unsafePerformEffect (Ref.read helpersRef)
          fieldTypes = fromMaybe [] (map _.fieldTypes (Map.lookup key helpers.ctorTypes))
          coercedFields = Array.mapWithIndex (\i f -> 
            let expectedType = case Array.index fieldTypes i of
                  Just ty -> exprTypeToGoType (unsafePerformEffect (Ref.read helpersRef)).pointerAdtPaths modNameStr ty
                  Nothing -> TypeValue
            in coerceGoExpr (GoVar (sanitizeName f)) TypeValue expectedType
          ) fields
          typeArgs = case getExprType tcoExpr of
            ADT _ _ tArgs -> map (exprTypeToGoType (unsafePerformEffect (Ref.read helpersRef)).pointerAdtPaths modNameStr) tArgs
            _ -> case Map.lookup key helpers.ctorTypes of
              Just ctorInfo -> map (const TypeValue) ctorInfo.typeVars
              Nothing -> []
          isElided = Set.member structName helpers.elidedCtors
          isPointerAdtLeaf = Map.member baseStructName helpers.pointerAdtLeaves
          funcExpr = if isElided then
              case Array.head fields of
                Just f -> GoCall (GoSelector (GoVar "gopurs_runtime") "Func") [ GoRaw ("func(" <> sanitizeName f <> " gopurs_runtime.Value) gopurs_runtime.Value {\nreturn " <> printGoExpr (coerceGoExpr (GoVar (sanitizeName f)) TypeValue TypeValue) <> "\n}") ]
                Nothing -> Array.foldr (\f inner -> GoCall (GoSelector (GoVar "gopurs_runtime") "Func") [ GoRaw ("func(" <> sanitizeName f <> " gopurs_runtime.Value) gopurs_runtime.Value {\nreturn " <> printGoExpr inner <> "\n}") ]) (GoConstructor (hashString baseStructName) structName typeArgs coercedFields) fields
            else if isPointerAdtLeaf then GoRaw ("gopurs_runtime.Value{Type: 9, IntVal: " <> hashString (fromMaybe "" (Map.lookup baseStructName helpers.pointerAdtLeaves)) <> ", UnsafePtr: nil}")
            else
              Array.foldr (\f inner -> GoCall (GoSelector (GoVar "gopurs_runtime") "Func") [ GoRaw ("func(" <> sanitizeName f <> " gopurs_runtime.Value) gopurs_runtime.Value {\nreturn " <> printGoExpr inner <> "\n}") ]) (GoConstructor (hashString baseStructName) structName typeArgs coercedFields) fields
        in
          { stmts: StmtEmpty, expr: funcExpr, exprType: TypeValue, nextId }

      CtorSaturated (Qualified mbMod _) _ _ (Ident name) props ->
        let
          structName = getStructName modNameStr mbMod name
          baseStructName = getBaseStructName modNameStr mbMod name
          modPart = case mbMod of
            Just (ModuleName mn) -> String.replaceAll (Pattern ".") (Replacement "_") mn
            Nothing -> modNameStr
          key = modPart <> "." <> name
          helpers = unsafePerformEffect (Ref.read helpersRef)
          
          ctorType = getExprType tcoExpr
          expectedGoType = exprTypeToGoType (unsafePerformEffect (Ref.read helpersRef)).pointerAdtPaths modNameStr ctorType
          
          fieldTypes = fromMaybe [] (map _.fieldTypes (Map.lookup key helpers.ctorTypes))
          
          accProps = foldl
            ( \acc (Tuple _ val) ->
                let
                  resVal = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars moduleArities bound Nothing [] false false acc.nextId val
                  expectedType = case Array.index fieldTypes acc.fieldIdx of
                    Just ty -> exprTypeToGoType (unsafePerformEffect (Ref.read helpersRef)).pointerAdtPaths modNameStr ty
                    Nothing -> TypeValue
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
                
          finalExpr =
            let
              monoStructName = "Constructor_" <> sanitizeName name
              pkgPrefix = case mbMod of
                Just (ModuleName mod) | String.replaceAll (Pattern ".") (Replacement "_") mod /= modNameStr -> "pkg_" <> String.replaceAll (Pattern ".") (Replacement "_") mod <> "."
                _ -> ""
              typeArgs = case ctorType of
                ADT _ _ tArgs -> map (exprTypeToGoType (unsafePerformEffect (Ref.read helpersRef)).pointerAdtPaths modNameStr) tArgs
                _ -> case Map.lookup key helpers.ctorTypes of
                  Just ctorInfo -> map (const TypeValue) ctorInfo.typeVars
                  Nothing -> []
            in if isElided then
                 case Array.head accProps.exprs of
                   Just expr -> boxGoExpr expr (fromMaybe TypeValue (map (exprTypeToGoType (unsafePerformEffect (Ref.read helpersRef)).pointerAdtPaths modNameStr) (Array.index fieldTypes 0)))
                   Nothing -> GoConstructor (hashString baseStructName) (pkgPrefix <> monoStructName) typeArgs accProps.exprs
               else if isPointerAdtLeaf then GoRaw ("gopurs_runtime.Value{Type: 9, IntVal: " <> hashString (fromMaybe "" (Map.lookup baseStructName helpers.pointerAdtLeaves)) <> ", UnsafePtr: nil}")
               else GoConstructor (hashString baseStructName) (pkgPrefix <> monoStructName) typeArgs accProps.exprs
        in
          { stmts: accProps.stmts, expr: finalExpr, exprType: expectedGoType, nextId: accProps.nextId }

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
              OpBooleanNot -> { stmts: resE.stmts, expr: GoBinOp "!=" (unboxGoExpr resE.expr resE.exprType TypeBool) (GoRaw "true"), exprType: TypeBool, nextId: resE.nextId }
              OpIntNegate -> { stmts: resE.stmts, expr: GoPrefixOp "-" (unboxGoExpr resE.expr resE.exprType TypeInt64), exprType: TypeInt64, nextId: resE.nextId }
              OpIntBitNot -> { stmts: resE.stmts, expr: GoBinOp "^" (GoRaw "^0") (unboxGoExpr resE.expr resE.exprType TypeInt64), exprType: TypeInt64, nextId: resE.nextId }
              OpNumberNegate -> { stmts: resE.stmts, expr: GoPrefixOp "-" (unboxGoExpr resE.expr resE.exprType TypeFloat64), exprType: TypeFloat64, nextId: resE.nextId }
              OpIsTag (Qualified mbMod (Ident tag)) ->
                let
                  baseStructName = getBaseStructName modNameStr mbMod tag
                  hashStr = hashString baseStructName
                  helpers = unsafePerformEffect (Ref.read helpersRef)
                in
                  case resE.expr of
                    GoVar _ ->
                      let
                        exprStr = case Map.lookup baseStructName helpers.pointerAdtLeaves of
                          Just nodeBaseStruct -> "(" <> printGoExpr (boxGoExpr resE.expr resE.exprType) <> ".Type == 9 && " <> printGoExpr (boxGoExpr resE.expr resE.exprType) <> ".IntVal == " <> hashString nodeBaseStruct <> " && " <> printGoExpr (boxGoExpr resE.expr resE.exprType) <> ".UnsafePtr == nil)"
                          Nothing -> if Set.member baseStructName helpers.pointerAdtNodes then
                            "(" <> printGoExpr (boxGoExpr resE.expr resE.exprType) <> ".Type == 9 && " <> printGoExpr (boxGoExpr resE.expr resE.exprType) <> ".IntVal == " <> hashStr <> " && " <> printGoExpr (boxGoExpr resE.expr resE.exprType) <> ".UnsafePtr != nil)"
                          else
                            "(" <> printGoExpr (boxGoExpr resE.expr resE.exprType) <> ".Type == 9 && " <> printGoExpr (boxGoExpr resE.expr resE.exprType) <> ".IntVal == " <> hashStr <> ")"
                      in { stmts: resE.stmts, expr: GoRaw exprStr, exprType: TypeBool, nextId: resE.nextId }
                    _ ->
                      let
                        tmpVar = "__t_tag_" <> show resE.nextId
                        declTmp = StmtLeaf (GoRaw ("var " <> tmpVar <> " gopurs_runtime.Value = " <> printGoExpr (boxGoExpr resE.expr resE.exprType)))
                        exprStr = case Map.lookup baseStructName helpers.pointerAdtLeaves of
                          Just nodeBaseStruct -> "(" <> tmpVar <> ".Type == 9 && " <> tmpVar <> ".IntVal == " <> hashString nodeBaseStruct <> " && " <> tmpVar <> ".UnsafePtr == nil)"
                          Nothing -> if Set.member baseStructName helpers.pointerAdtNodes then
                            "(" <> tmpVar <> ".Type == 9 && " <> tmpVar <> ".IntVal == " <> hashStr <> " && " <> tmpVar <> ".UnsafePtr != nil)"
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
      mbDecl = Array.find (\d -> d.typeName == typeName || (Array.length path > 0 && d.typeName == fromMaybe "" (Array.last path))) dataDecls
  in case mbDecl of
       Just decl ->
         if Array.length decl.constructors == 1 then
           case Array.head decl.constructors of
             Just ctor ->
               if Array.length ctor.fieldTypes == 1 then
                 case Array.head ctor.fieldTypes of
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

generateFfiBridge :: Array DataDecl -> Array FfiDecl -> Array (Tuple Ident (Maybe ExprType)) -> String
generateFfiBridge dataDecls decls foreigns = 
  String.joinWith "\n" (map genBridge foreigns)
  where
  genBridge (Tuple ident mbTast) = 
    let 
      pursName = unwrap ident
      sanitized = sanitizeName pursName
      capName = capitalize sanitized
      exportName = "_Gopurs_" <> capName
      
      fallback1 = capitalize pursName
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
