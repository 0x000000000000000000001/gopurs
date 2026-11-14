module Gopurs.FfiBridge.Signatures
  ( NativeCandidate
  , matchDeclaration
  , nativeCandidate
  , functionInfo
  , workerReturnsNative
  ) where

import Prelude

import Data.Array as Array
import Data.Maybe (Maybe(..))
import Data.String as String
import Data.String.Pattern (Pattern(..))
import Data.Traversable (traverse)
import Data.Tuple (Tuple(..))
import Gopurs.CodegenState (FunctionInfo)
import Gopurs.FfiTypes (FfiDecl, TypeNode(..))
import Gopurs.FfiBridge.TypeSupport (isPointerNode)
import Gopurs.GoAst (GoType(..), capitalize)
import Gopurs.GoSymbol (GoReference(..), foreignSymbol)
import PureScript.Backend.Optimizer.CoreFn (ExprType(..))

-- Worker generation and direct-call publication share this candidate. A
-- generated worker can still have an unsupported direct-call signature.
type NativeCandidate =
  { declaration :: FfiDecl
  , tast :: ExprType
  , requiresWorker :: Boolean
  , callName :: String
  }

matchDeclaration :: String -> Array FfiDecl -> String -> Maybe FfiDecl
matchDeclaration modNameStr decls pursName =
  let
    name = modNameStr <> "_" <> capitalize pursName
    findDecl n = Array.find (\d -> d.name == n) decls
  in
    case findDecl name of
      Just d -> Just d
      Nothing -> findDecl (name <> "_")

nativeCandidate :: FfiDecl -> ExprType -> Maybe NativeCandidate
nativeCandidate declaration tast
  -- The generic boxed entry instantiates its type parameters explicitly.
  -- Direct pointer adapters must not emit an uninstantiated Go function.
  | not (Array.null declaration.typeParams)
      && (Array.any isPointerNode declaration.args || case declaration.ret of
        Just result -> isPointerNode result
        Nothing -> false) = Nothing
  | Array.length declaration.results <= 1 && not (Array.any containsStruct (declaration.args <> declaration.results)) && not (mentionsForeign tast) && isPureConcreteResult tast && not declaration.isVar && Array.length declaration.args >= 1 =
      let
        requiresWorker = hasFunctionParameter declaration.args || Array.any isPointerNode declaration.args || case declaration.ret of
          Nothing -> true
          Just result -> needsValueWorker result || isPointerNode result
      in
        Just
          { declaration
          , tast
          , requiresWorker
          , callName: declaration.name <> if requiresWorker then "_nativeWorker" else ""
          }
  | otherwise = Nothing

hasFunctionParameter :: Array TypeNode -> Boolean
hasFunctionParameter = Array.any case _ of
  TFunc _ _ -> true
  _ -> false

functionInfo :: String -> NativeCandidate -> Maybe FunctionInfo
functionInfo unit { declaration, tast, requiresWorker, callName } = do
  args <- traverse argumentType declaration.args
  ret <- case declaration.ret of
    Just t | requiresWorker && workerReturnsNative declaration -> typeNodeToGoType t
    Just _ | requiresWorker -> boxedResult
    Just t -> typeNodeToGoType t
    Nothing -> boxedResult
  pure { reference: SymbolReference (foreignSymbol unit callName), fArgs: args, fRet: ret, arity: Array.length declaration.args }
  where
  argumentType (TFunc _ _) | requiresWorker = Just TypeValue
  argumentType t | requiresWorker && isPointerNode t = Just TypeValue
  argumentType t = typeNodeToGoType t

  boxedResult = if isRecordResult tast then Nothing else Just TypeValue

-- Pointer adapters need boxed inputs, but scalar outputs can keep the same
-- direct representation as a handwritten scalar FFI function.
workerReturnsNative :: FfiDecl -> Boolean
workerReturnsNative declaration = Array.any isPointerNode declaration.args && case declaration.ret of
  Just t -> case typeNodeToGoType t of
    Just TypeString -> true
    Just TypeInt64 -> true
    Just TypeFloat64 -> true
    Just TypeBool -> true
    Just TypeUint32 -> true
    Just TypeValue -> true
    _ -> false
  Nothing -> false

isRecordResult :: ExprType -> Boolean
isRecordResult = case _ of
  Func _ ret -> isRecordResult ret
  Record _ -> true
  Row _ _ -> true
  _ -> false

typeNodeToGoType :: TypeNode -> Maybe GoType
typeNodeToGoType = case _ of
  TNamed "gopurs_runtime.Value" -> Just TypeValue
  TNamed "string" -> Just TypeString
  TNamed "int64" -> Just TypeInt64
  TNamed "float64" -> Just TypeFloat64
  TNamed "bool" -> Just TypeBool
  TNamed "uint32" -> Just TypeUint32
  TNamed "any" -> Just (TypeInterface "any")
  TNamed "interface{}" -> Just (TypeInterface "interface{}")
  TArray elem -> TypeNativeArray <$> typeNodeToGoType elem
  _ -> Nothing

needsValueWorker :: TypeNode -> Boolean
needsValueWorker = case _ of
  TNamed "any" -> true
  TNamed "interface{}" -> true
  TArray elem -> needsValueWorker elem
  _ -> false

-- Keep the existing admission rule: inspect the final PureScript result,
-- rejecting quantified, constrained, unknown and effectful results.
isPureConcreteResult :: ExprType -> Boolean
isPureConcreteResult = case _ of
  Func _ ret -> isPureConcreteResult ret
  ForAll _ _ -> false
  ConstrainedType _ _ -> false
  TypeVar _ -> false
  Any -> false
  other -> not (mentionsEffect other)

mentionsEffect :: ExprType -> Boolean
mentionsEffect = case _ of
  ADT fullName _ args -> String.contains (Pattern "Effect") fullName || Array.any mentionsEffect args
  TypeApp c args -> mentionsEffect c || Array.any mentionsEffect args
  Array t -> mentionsEffect t
  Func args ret -> Array.any mentionsEffect args || mentionsEffect ret
  Record row -> mentionsEffect row
  Row fields tail ->
    Array.any (\(Tuple _ v) -> mentionsEffect v) fields
      || case tail of
        Just t -> mentionsEffect t
        Nothing -> false
  ForAll _ t -> mentionsEffect t
  ConstrainedType constraints t ->
    Array.any (\(Tuple _ as) -> Array.any mentionsEffect as) constraints || mentionsEffect t
  _ -> false

-- Dynamic Foreign boundaries require unboxing dictated by the PureScript type;
-- do not advertise their boxed entry as a direct native call.
mentionsForeign :: ExprType -> Boolean
mentionsForeign = case _ of
  ADT name _ args -> name == "Foreign.Foreign" || Array.any mentionsForeign args
  Func args ret -> Array.any mentionsForeign args || mentionsForeign ret
  Array t -> mentionsForeign t
  Record t -> mentionsForeign t
  Row fields tail -> Array.any (\(Tuple _ t) -> mentionsForeign t) fields || case tail of
    Just t -> mentionsForeign t
    Nothing -> false
  TypeApp t args -> mentionsForeign t || Array.any mentionsForeign args
  ForAll _ t -> mentionsForeign t
  ConstrainedType _ t -> mentionsForeign t
  _ -> false

-- Structured fields must go through the TAST-guided adapter. In particular a
-- callback worker returning an array of structs cannot use generic Box.
containsStruct :: TypeNode -> Boolean
containsStruct = case _ of
  TStruct _ _ -> true
  TIdentified _ _ underlying -> case underlying of
    Just t -> containsStruct t
    Nothing -> false
  TPointer t -> containsStruct t
  TArray t -> containsStruct t
  TMap k v -> containsStruct k || containsStruct v
  TFunc args ret -> Array.any containsStruct args || case ret of
    Just t -> containsStruct t
    Nothing -> false
  _ -> false
