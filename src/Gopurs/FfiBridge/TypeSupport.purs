module Gopurs.FfiBridge.TypeSupport
  ( printTypeNode
  , instantiateDeclaration
  , getTastArgType
  , getTastReturnType
  , tastForArity
  , resolveNewtype
  , flattenFuncArgs
  , isStandardPursFunc
  , isUncurriedFunction
  , isRuntimeValueNode
  , isPointerNode
  ) where

import Prelude

import Data.Array as Array
import Data.Map as Map
import Data.Maybe (Maybe(..), fromMaybe)
import Data.String as String
import Data.Tuple (Tuple(..))
import Gopurs.FfiTypes (FfiDecl, TypeNode(..))
import PureScript.Backend.Optimizer.CoreFn (DataDecl, ExprType(..))
import PureScript.Backend.Optimizer.TypeSubstitution (substitute)

printTypeNode :: TypeNode -> String
printTypeNode (TNamed n) = n
printTypeNode (TFunc args ret) =
  let
    argsStr = String.joinWith ", " (map printTypeNode args)
    retStr = case ret of
      Nothing -> ""
      Just r -> " " <> printTypeNode r
  in
    "func(" <> argsStr <> ")" <> retStr
printTypeNode (TArray elem) = "[]" <> printTypeNode elem
printTypeNode (TMap k v) = "map[" <> printTypeNode k <> "]" <> printTypeNode v
printTypeNode (TIdentified n _ _) = n
printTypeNode (TPointer t) = "*" <> printTypeNode t
printTypeNode (TFixedArray t n) = "[" <> n <> "]" <> printTypeNode t
printTypeNode (TVariadic t) = "..." <> printTypeNode t
printTypeNode (TInterface s) = s
printTypeNode (TMultiFunc args results) = "func(" <> String.joinWith ", " (map printTypeNode args) <> ") (" <> String.joinWith ", " (map printTypeNode results) <> ")"
printTypeNode (TStruct source _) = source
printTypeNode (TUnknown s) = s

-- Pointer representations stay inside the FFI module; callers pass an opaque
-- boxed handle. Follow local aliases as well as explicit pointer nodes.
isPointerNode :: TypeNode -> Boolean
isPointerNode (TPointer _) = true
isPointerNode (TIdentified _ _ (Just underlying)) = isPointerNode underlying
isPointerNode (TUnknown name) = String.take 1 name == "*"
isPointerNode _ = false

-- Wrapper calls instantiate Go type parameters with Value. The same
-- substitution must reach callbacks and containers before adapting values.
instantiateDeclaration :: FfiDecl -> FfiDecl
instantiateDeclaration declaration = declaration
  { args = map instantiate declaration.args
  , ret = map instantiate declaration.ret
  , results = map instantiate declaration.results
  }
  where
  instantiate = case _ of
    TNamed name | Array.elem name declaration.typeParams -> TNamed "gopurs_runtime.Value"
    TFunc args ret -> TFunc (map instantiate args) (map instantiate ret)
    TMultiFunc args results -> TMultiFunc (map instantiate args) (map instantiate results)
    TPointer t -> TPointer (instantiate t)
    TFixedArray t n -> TFixedArray (instantiate t) n
    TVariadic t -> TVariadic (instantiate t)
    TStruct source fields -> TStruct source (map (\f -> f { type = instantiate f.type }) fields)
    TIdentified n i u -> TIdentified n i (map instantiate u)
    TArray item -> TArray (instantiate item)
    TMap key value -> TMap (instantiate key) (instantiate value)
    other -> other

getTastArgType :: ExprType -> Int -> Maybe ExprType
getTastArgType (Func args _) i = Array.index args i
getTastArgType (ForAll _ body) i = getTastArgType body i
getTastArgType (ADT name _ args) i | isUncurriedFunction name = Array.index (Array.dropEnd 1 args) i
getTastArgType _ _ = Nothing

getTastReturnType :: ExprType -> Maybe ExprType
getTastReturnType (Func _ ret) = Just ret
getTastReturnType (ForAll _ body) = getTastReturnType body
getTastReturnType (ADT name _ args) | name == "Effect.Effect" || isUncurriedFunction name = Array.last args
getTastReturnType _ = Nothing

-- A native declaration can consume arguments from a returned function, or
-- return a function before consuming the whole first TAST argument group.
-- Keep the residual result at the actual Go call boundary; flattening every
-- result would lose genuinely returned functions.
tastForArity :: Int -> ExprType -> Maybe ExprType
tastForArity arity tast = do
  call <- consume arity tast
  pure (Func call.args call.result)
  where
  consume count (ForAll _ body) = consume count body
  consume count (Func args result) = group count args result
  consume count (ADT name _ args) | isUncurriedFunction name = do
    result <- Array.last args
    group count (Array.dropEnd 1 args) result
  consume _ _ = Nothing

  group count args result
    | count < 0 = Nothing
    | count < Array.length args = Just
        { args: Array.take count args
        , result: Func (Array.drop count args) result
        }
    | count == Array.length args = Just { args, result }
    | otherwise = do
        rest <- consume (count - Array.length args) result
        pure { args: args <> rest.args, result: rest.result }

-- Single-field data constructors use the same elided layout as newtypes in
-- this backend. Instantiate their fields before following that representation.
resolveNewtype :: Array DataDecl -> ExprType -> ExprType
resolveNewtype declarations = go []
  where
  go seen original@(ADT fullName path args)
    | Array.elem fullName seen = original
    | otherwise =
        case Array.find (\d -> d.name == fullName || d.name == fromMaybe "" (Array.last path)) declarations of
          Just decl | Array.length decl.vars == Array.length args -> case decl.constructors of
            [ { fields: [ field ] } ] -> go (seen <> [fullName]) (substitute (Map.fromFoldable (Array.zip decl.vars args)) field)
            _ -> original
          _ -> original
  go seen (Func args ret) = Func (map (go seen) args) (go seen ret)
  go seen (Array t) = Array (go seen t)
  go seen (Record row) = Record (go seen row)
  go seen (Row fields tail) = Row (map (\(Tuple k v) -> Tuple k (go seen v)) fields) (map (go seen) tail)
  go seen (TypeApp c args) = TypeApp (go seen c) (map (go seen) args)
  go seen (ForAll vars body) = ForAll vars (go seen body)
  go seen (ConstrainedType constraints body) = ConstrainedType (map (\(Tuple c args) -> Tuple c (map (go seen) args)) constraints) (go seen body)
  go _ other = other

flattenFuncArgs :: ExprType -> Array ExprType
flattenFuncArgs (Func args ret) = args <> flattenFuncArgs ret
flattenFuncArgs _ = []

isStandardPursFunc :: TypeNode -> Boolean
isStandardPursFunc (TFunc args ret) =
  let
    argsAny = Array.all (\a -> printTypeNode a == "any" || printTypeNode a == "interface{}" || printTypeNode a == "gopurs_runtime.Value") args
    retStr = case ret of
      Nothing -> ""
      Just r -> printTypeNode r
  in
    if Array.length args == 0 then
      retStr == "" || retStr == "bool" || retStr == "int" || retStr == "int64" || retStr == "string" || retStr == "float64" || retStr == "gopurs_runtime.Value" || retStr == "any" || retStr == "interface{}"
    else if argsAny then
      case ret of
        Nothing -> true
        Just r@(TFunc _ _) -> isStandardPursFunc r
        Just _ -> retStr == "any" || retStr == "interface{}" || retStr == "gopurs_runtime.Value"
    else
      false
isStandardPursFunc _ = false

isRuntimeValueNode :: TypeNode -> Boolean
isRuntimeValueNode = case _ of
  TNamed "gopurs_runtime.Value" -> true
  _ -> false

isUncurriedFunction :: String -> Boolean
isUncurriedFunction name = Array.any (\arity -> name == "Effect.Uncurried.EffectFn" <> show arity || name == "Data.Function.Uncurried.Fn" <> show arity) (Array.range 0 10)
