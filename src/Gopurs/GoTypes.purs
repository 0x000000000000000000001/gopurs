module Gopurs.GoTypes
  ( isClosedRowTail
  , visibleRecordFields
  , printExprType
  , AppliedAdt
  , appliedAdt
  , exprTypeToGoType
  , exprTypeToGenericGoType
  , structFieldGoType
  , instantiateGenericGoType
  ) where

import Prelude

import Data.Array as Array
import Data.Map as Map
import Data.Maybe (Maybe(..), fromMaybe)
import Data.Set as Set
import Data.String as String
import Data.String.Pattern (Pattern(..), Replacement(..))
import Data.Tuple (Tuple(..))
import Gopurs.AdtMetadata (PointerAdtPaths)
import Gopurs.GoAst (GoType(..), constructorNames, structPointer)
import PureScript.Backend.Optimizer.CoreFn (ExprType(..))

-- Diagnostic rendering of TAST annotations, shared by codegen and FFI comments.
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
    let
      tailStr = case tail of
        Nothing -> "Empty"
        Just t -> printExprType t
    in
      "(Row [" <> String.joinWith ", " (map (\(Tuple k v) -> k <> ": " <> printExprType v) props) <> "] " <> tailStr <> ")"
  TypeApp c args -> "(TypeApp " <> printExprType c <> " [" <> String.joinWith ", " (map printExprType args) <> "])"
  ForAll vars body -> "(ForAll [" <> String.joinWith ", " vars <> "] " <> printExprType body <> ")"
  ConstrainedType _ body -> "(ConstrainedType " <> printExprType body <> ")"
  ADT _ path args -> "(ADT " <> show path <> " [" <> String.joinWith ", " (map printExprType args) <> "])"
  TypeVar v -> "(TypeVar " <> v <> ")"
  Any -> "Any"

-- Only an absent tail proves that a row is closed. An unknown tail (Any)
-- may contain additional fields, which a native struct would discard.
isClosedRowTail :: Maybe ExprType -> Boolean
isClosedRowTail Nothing = true
isClosedRowTail _ = false

-- Row.Union can retain duplicate labels in the TAST. Only the first
-- occurrence is visible in the runtime record, including its field type.
visibleRecordFields :: forall a. Array (Tuple String a) -> Array (Tuple String a)
visibleRecordFields = Array.nubBy (comparing \(Tuple label _) -> label)

type AppliedAdt = { fullName :: String, path :: Array String, args :: Array ExprType }

-- Read nested applications from the inside out, after the arguments already
-- carried by the ADT. This does not unwrap ForAll or constrained annotations.
appliedAdt :: ExprType -> Maybe AppliedAdt
appliedAdt ty = go ty []
  where
  go (TypeApp fn args) applied = go fn (args <> applied)
  go (ADT fullName path args) applied = Just { fullName, path, args: args <> applied }
  go _ _ = Nothing

-- TAST annotations and ADT metadata select the internal Go representation.
-- FFI wrappers separately reconcile it with the declared foreign Go signature.
exprTypeToGoType :: PointerAdtPaths -> Set.Set String -> Set.Set String -> String -> ExprType -> GoType
exprTypeToGoType _ _ _ _ Int = TypeInt64
exprTypeToGoType _ _ _ _ Number = TypeFloat64
exprTypeToGoType _ _ _ _ String = TypeString
exprTypeToGoType _ _ _ _ Char = TypeString
exprTypeToGoType _ _ _ _ Boolean = TypeBool
exprTypeToGoType ptrPaths enumAdts elided modNameStr (Array ty) = TypeNativeArray (exprTypeToGoType ptrPaths enumAdts elided modNameStr ty)
exprTypeToGoType ptrPaths enumAdts elided modNameStr (Record (Row fields tail)) | isClosedRowTail tail =
  recordGoType (exprTypeToGoType ptrPaths enumAdts elided modNameStr) fields
exprTypeToGoType ptrPaths enumAdts elided modNameStr ty = case appliedAdt ty of
  Just { fullName, path, args } ->
    let
      -- Value annotations test elision using the ADT path, before looking up
      -- its payload constructor. Generic fields below use that constructor.
      annotationNames = adtConstructorNames path (fromMaybe "" (Array.last path))
      toGoType = exprTypeToGoType ptrPaths enumAdts elided modNameStr
    in
      if Set.member annotationNames.structName elided then TypeValue
      else if Set.member fullName enumAdts then TypeUint32
      else case lookupPointerAdt fullName ptrPaths of
        Just info ->
          let names = adtConstructorNames path info.ctorName
          in structPointer { baseStructName: names.baseStructName, fullName, structName: names.structName }
            (valueTypeArguments info.arity (map toGoType args))
        Nothing -> TypeValue
  Nothing -> TypeValue

-- Ordinary type variables fall back to Value. In a generic declaration,
-- variables listed in typeVars can instead become Go type parameters.
exprTypeToGenericGoType :: PointerAdtPaths -> Set.Set String -> Set.Set String -> Array String -> String -> ExprType -> GoType
exprTypeToGenericGoType ptrPaths enumAdts elided typeVars modNameStr (Record (Row fields tail)) | isClosedRowTail tail =
  recordGoType (exprTypeToGenericGoType ptrPaths enumAdts elided typeVars modNameStr) fields
exprTypeToGenericGoType _ _ _ typeVars _ (TypeVar v) | Array.elem v typeVars = TypeGenericParam v
exprTypeToGenericGoType ptrPaths enumAdts elided typeVars modNameStr ty = case appliedAdt ty of
  Just { fullName, path, args } ->
    if Set.member fullName enumAdts then TypeUint32
    else case lookupPointerAdt fullName ptrPaths of
      Just info ->
        let
          names = adtConstructorNames path info.ctorName
          toGoType = exprTypeToGenericGoType ptrPaths enumAdts elided typeVars modNameStr
        in
          if Set.member names.structName elided then TypeValue
          else structPointer { baseStructName: names.baseStructName, fullName, structName: names.structName }
            (genericTypeArguments toGoType typeVars info.arity args)
      Nothing -> TypeValue
  -- Arrays and other annotations retain the ordinary value representation.
  Nothing -> exprTypeToGoType ptrPaths enumAdts elided modNameStr ty

recordGoType :: (ExprType -> GoType) -> Array (Tuple String ExprType) -> GoType
recordGoType toGoType fields =
  TypeRecord (map (\(Tuple name ty) -> Tuple name (toGoType ty))
    (Array.sortBy (comparing \(Tuple name _) -> name) (visibleRecordFields fields)))

lookupPointerAdt :: String -> PointerAdtPaths -> Maybe { ctorName :: String, arity :: Int }
lookupPointerAdt fullName paths = case Map.lookup fullName paths of
  Just info -> Just info
  Nothing -> Map.lookup (fullName <> "$Dict") paths

adtConstructorNames :: Array String -> String -> { baseStructName :: String, structName :: String }
adtConstructorNames path =
  let
    moduleName = String.joinWith "." (Array.take (Array.length path - 1) path)
    modulePrefix = String.replaceAll (Pattern ".") (Replacement "_") moduleName
  in
    constructorNames modulePrefix

-- A value keeps the supplied prefix, pads missing slots with Value, and drops
-- excess arguments. Its pointer always has the metadata's declared arity.
valueTypeArguments :: Int -> Array GoType -> Array GoType
valueTypeArguments arity args =
  let supplied = Array.take arity args
  in supplied <> Array.replicate (arity - Array.length supplied) TypeValue

-- A generic field needs a complete instantiation. Otherwise it inherits all
-- declaration parameters when their count fits, or erases every slot to Value.
genericTypeArguments :: (ExprType -> GoType) -> Array String -> Int -> Array ExprType -> Array GoType
genericTypeArguments toGoType vars arity args
  | arity == 0 = []
  | Array.length args == arity = map toGoType args
  | Array.length vars == arity = map TypeGenericParam vars
  | otherwise = Array.replicate arity TypeValue

structFieldGoType :: PointerAdtPaths -> Set.Set String -> Set.Set String -> Array String -> String -> ExprType -> GoType
structFieldGoType ptrPaths enumAdts elidedCtors typeVars modStr ty =
  case exprTypeToGenericGoType ptrPaths enumAdts elidedCtors typeVars modStr ty of
    TypeInterface _ -> TypeValue
    other -> other

instantiateGenericGoType :: Map.Map String GoType -> GoType -> GoType
instantiateGenericGoType env (TypeGenericParam v) = fromMaybe TypeValue (Map.lookup v env)
instantiateGenericGoType env (TypeRecord fields) = TypeRecord (map (\(Tuple k v) -> Tuple k (instantiateGenericGoType env v)) fields)
instantiateGenericGoType env (TypeNativeArray ty) = TypeNativeArray (instantiateGenericGoType env ty)
instantiateGenericGoType env (TypeStructPointer pointer) =
  structPointer pointer (map (instantiateGenericGoType env) pointer.typeArgs)
instantiateGenericGoType env (TypeFunc args ret) = TypeFunc (map (instantiateGenericGoType env) args) (instantiateGenericGoType env ret)
instantiateGenericGoType _ t = t
