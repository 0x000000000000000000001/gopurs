module Gopurs.GoTypes
  ( isClosedRowTail
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
import Gopurs.GoAst (GoType(..), goTypeToStr, sanitizeName)
import PureScript.Backend.Optimizer.CoreFn (ExprType(..))

-- Nothing and Just Any count as closed row tails in this mapping. Closed
-- records become native structs with fields sorted by label; other tails use Value.
isClosedRowTail :: Maybe ExprType -> Boolean
isClosedRowTail Nothing = true
isClosedRowTail (Just Any) = true
isClosedRowTail _ = false

-- TAST annotations and ADT metadata select the internal Go representation.
-- FFI wrappers separately reconcile it with the declared foreign Go signature.
exprTypeToGoType :: Map.Map String { ctorName :: String, arity :: Int } -> Set.Set String -> Set.Set String -> String -> ExprType -> GoType
exprTypeToGoType _ _ _ _ Int = TypeInt64
exprTypeToGoType _ _ _ _ Number = TypeFloat64
exprTypeToGoType _ _ _ _ String = TypeString
exprTypeToGoType _ _ _ _ Char = TypeString
exprTypeToGoType _ _ _ _ Boolean = TypeBool
exprTypeToGoType ptrPaths enumAdts elided modNameStr (Array ty) = TypeNativeArray (exprTypeToGoType ptrPaths enumAdts elided modNameStr ty)
exprTypeToGoType ptrPaths enumAdts elided modNameStr (Record (Row fields tail)) | isClosedRowTail tail = TypeRecord (map (\(Tuple k v) -> Tuple k (exprTypeToGoType ptrPaths enumAdts elided modNameStr v)) (Array.sortBy (comparing \(Tuple k _) -> k) fields))
exprTypeToGoType ptrPaths enumAdts elided modNameStr (Record _) = TypeValue
exprTypeToGoType ptrPaths enumAdts elided modNameStr (ADT fullName path args) =
  let
    ctorName = fromMaybe "" (Array.last path)
    pkgNameStr = String.replaceAll (Pattern ".") (Replacement "_") (String.joinWith "." (Array.slice 0 (Array.length path - 1) path))
    monoStructName = "Constructor_" <> pkgNameStr <> "_" <> sanitizeName ctorName
  in
    if Set.member monoStructName elided then TypeValue
    else if Set.member fullName enumAdts then TypeUint32
    else
      case
        ( case Map.lookup fullName ptrPaths of
            Just i -> Just i
            Nothing -> Map.lookup (fullName <> "$Dict") ptrPaths
        )
        of
        Just info ->
          let
            baseStructName = "Data_" <> pkgNameStr <> "_" <> sanitizeName info.ctorName
            monoStructName' = "Constructor_" <> pkgNameStr <> "_" <> sanitizeName info.ctorName
            typeArgsMapped = map (exprTypeToGoType ptrPaths enumAdts elided modNameStr) args
            typeArgsMappedTruncated = Array.take info.arity typeArgsMapped
            paddedTypeArgs = typeArgsMappedTruncated <> Array.replicate (info.arity - Array.length typeArgsMappedTruncated) TypeValue
            typeArgsStr = if Array.length paddedTypeArgs > 0 then "[" <> String.joinWith ", " (map goTypeToStr paddedTypeArgs) <> "]" else ""
          in
            TypeStructPointer baseStructName fullName (monoStructName' <> typeArgsStr) paddedTypeArgs
        Nothing -> TypeValue
exprTypeToGoType ptrPaths enumAdts elided modNameStr (TypeApp fn arg) =
  let
    unwrapTypeApp :: ExprType -> Array ExprType -> Tuple ExprType (Array ExprType)
    unwrapTypeApp (TypeApp f a) acc = unwrapTypeApp f (a <> acc)
    unwrapTypeApp other acc = Tuple other acc
  in
    case unwrapTypeApp (TypeApp fn arg) [] of
      Tuple (ADT fullName path args) allArgs -> exprTypeToGoType ptrPaths enumAdts elided modNameStr (ADT fullName path (args <> allArgs))
      _ -> TypeValue
exprTypeToGoType _ _ _ _ (TypeVar v) = TypeValue
exprTypeToGoType _ _ _ _ _ = TypeValue

-- Ordinary type variables fall back to Value. In a generic declaration,
-- variables listed in typeVars can instead become Go type parameters.
exprTypeToGenericGoType :: Map.Map String { ctorName :: String, arity :: Int } -> Set.Set String -> Set.Set String -> Array String -> String -> ExprType -> GoType
exprTypeToGenericGoType ptrPaths enumAdts elidedCtors typeVars modNameStr (Record (Row fields tail)) | isClosedRowTail tail = TypeRecord (map (\(Tuple k v) -> Tuple k (exprTypeToGenericGoType ptrPaths enumAdts elidedCtors typeVars modNameStr v)) (Array.sortBy (comparing \(Tuple k _) -> k) fields))
exprTypeToGenericGoType _ _ _ _ _ (Record _) = TypeValue
exprTypeToGenericGoType ptrPaths enumAdts elidedCtors typeVars modNameStr (TypeApp fn arg) =
  let
    unwrapTypeApp :: ExprType -> Array ExprType -> Tuple ExprType (Array ExprType)
    unwrapTypeApp (TypeApp f a) acc = unwrapTypeApp f (a <> acc)
    unwrapTypeApp other acc = Tuple other acc
  in
    case unwrapTypeApp (TypeApp fn arg) [] of
      Tuple (ADT fullName path args) allArgs -> exprTypeToGenericGoType ptrPaths enumAdts elidedCtors typeVars modNameStr (ADT fullName path (args <> allArgs))
      _ -> TypeValue
exprTypeToGenericGoType _ _ _ typeVars _ (TypeVar v) | Array.elem v typeVars = TypeGenericParam v
exprTypeToGenericGoType ptrPaths enumAdts elided typeVars modNameStr (ADT fullName path args) =
  if Set.member fullName enumAdts then TypeUint32
  else
    case
      ( case Map.lookup fullName ptrPaths of
          Just i -> Just i
          Nothing -> Map.lookup (fullName <> "$Dict") ptrPaths
      )
      of
        Just info ->
          let
            pkgNameStr = String.replaceAll (Pattern ".") (Replacement "_") (String.joinWith "." (Array.slice 0 (Array.length path - 1) path))
            monoStructName = "Constructor_" <> pkgNameStr <> "_" <> sanitizeName info.ctorName
            baseStructName = "Data_" <> pkgNameStr <> "_" <> sanitizeName info.ctorName
          in
          if Set.member monoStructName elided then TypeValue
          else if info.arity == 0 then TypeStructPointer baseStructName fullName monoStructName []
          else
            let
              finalArgs =
                if Array.length args == info.arity then
                  map (exprTypeToGenericGoType ptrPaths enumAdts elided typeVars modNameStr) args
                else if Array.length typeVars == info.arity then
                  map TypeGenericParam typeVars
                else
                  Array.replicate info.arity TypeValue
              typeArgsStr = if Array.length finalArgs > 0 then "[" <> String.joinWith ", " (map goTypeToStr finalArgs) <> "]" else ""
            in
              TypeStructPointer baseStructName fullName (monoStructName <> typeArgsStr) finalArgs
        Nothing -> TypeValue
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
instantiateGenericGoType env (TypeStructPointer base key full typeArgs) =
  let
    newTypeArgs = map (instantiateGenericGoType env) typeArgs
    typeArgsStr = if Array.length newTypeArgs > 0 then "[" <> String.joinWith ", " (map goTypeToStr newTypeArgs) <> "]" else ""
    monoStructName = case String.indexOf (Pattern "[") full of
      Just i -> String.take i full
      Nothing -> full
  in
    TypeStructPointer base key (monoStructName <> typeArgsStr) newTypeArgs
instantiateGenericGoType env (TypeFunc args ret) = TypeFunc (map (instantiateGenericGoType env) args) (instantiateGenericGoType env ret)
instantiateGenericGoType env t = t
