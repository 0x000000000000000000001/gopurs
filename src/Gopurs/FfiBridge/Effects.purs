module Gopurs.FfiBridge.Effects
  ( effectContract
  , isLegacy
  , wrapper
  , problem
  ) where

import Prelude

import Data.Array as Array
import Data.Maybe (Maybe(..), fromMaybe)
import Data.String as String
import Gopurs.FfiBridge.Structured as Structured
import Gopurs.FfiBridge.TypeSupport (getTastArgType, getTastReturnType)
import Gopurs.FfiBridge.Values as Values
import Gopurs.FfiTypes (FfiDecl, TypeNode(..))
import PureScript.Backend.Optimizer.CoreFn (DataDecl, ExprType(..))

type Contract = { kind :: String, args :: Array ExprType, result :: ExprType }

-- Select by the source type's identity, never by the shape of its callback ABI.
effectContract :: ExprType -> Maybe Contract
effectContract (ForAll _ body) = effectContract body
effectContract (Func args result) = map (\c -> c { args = args <> c.args }) (effectContract result)
effectContract (ADT name _ [ result ])
  | name == "Effect.Aff.Compat.EffectFnAff" = Just { kind: "aff", args: [], result }
  | name == "Effect.Effect" = Just { kind: "effect", args: [], result }
effectContract _ = Nothing

isLegacy :: FfiDecl -> Boolean
isLegacy d = case d.ret of
  Just (TFunc _ _) -> true
  Just (TNamed "gopurs_runtime.Value") -> true
  Just (TNamed "any") -> true
  Just (TNamed "interface{}") -> true
  _ -> d.isVar

legacyContract :: Contract -> FfiDecl -> Boolean
legacyContract c d = isLegacy d ||
  (c.kind == "effect" && Array.length d.args == Array.length c.args + 1 && fromMaybe false (unitArgument <$> Array.last d.args))
  where
  unitArgument (TNamed "any") = true
  unitArgument (TNamed "interface{}") = true
  unitArgument (TNamed "gopurs_runtime.Value") = true
  unitArgument _ = false

isContext :: TypeNode -> Boolean
isContext (TIdentified _ "context.Context" _) = true
isContext _ = false

isError :: TypeNode -> Boolean
isError (TIdentified _ "builtin.error" _) = true
isError _ = false

isCleanup :: TypeNode -> Boolean
isCleanup (TFunc [] Nothing) = true
isCleanup _ = false

acquisition :: FfiDecl -> Boolean
acquisition d = case d.results of
  [ _, cleanup, err ] -> isCleanup cleanup && isError err
  _ -> false

unitAcquisition :: FfiDecl -> ExprType -> Boolean
unitAcquisition d Unit = case d.results of
  [ cleanup, err ] -> isCleanup cleanup && isError err
  _ -> false
unitAcquisition _ _ = false

valueResults :: FfiDecl -> Array TypeNode
valueResults d = if acquisition d then Array.take 1 d.results else Array.filter (not <<< isError) d.results


problem :: FfiDecl -> ExprType -> Maybe String
problem d tast | not (isLegacy d) && constrainedContract tast = Just "Ordinary Go FFI does not infer type-class dictionaries for a constrained import. Pass explicit typed callbacks or use the existing explicit runtime ABI."
problem d tast | not (isLegacy d) && directAff tast = Just "Ordinary Go implementations of Aff are not supported; use EffectFnAff and fromEffectFnAff."
problem d tast | not (isLegacy d) && (Array.length d.results /= 1 || fromMaybe false (ordinaryResult <$> d.ret)) && expandedAffCallbacks tast =
  Just "This PureScript result requires an EffectFn2/EffectFn3 callback, but the Go implementation returns an ordinary result. If the source declares EffectFnAff, select the matching typed purs frontend on PATH (including node_modules/.bin) and rebuild from clean frontend output, preserving go.mod/go.sum: the nominal Effect.Aff.Compat.EffectFnAff type must be retained. The purs version number alone does not identify its TAST contract. Explicit callback imports require a compatible callback or legacy runtime implementation."
problem d _ | Array.any unsupported (d.args <> d.results) = Just "This Go signature contains an unsupported variadic, fixed array, multiple-result callback or unresolved type. Use slices and single-result callbacks at this boundary."
problem d tast | Just reason <- structuredProblem d tast = Just reason
problem d tast = case effectContract tast of
  Just c | c.kind == "aff" && ambiguousResult d -> Just "An any result is ambiguous between a native dynamic value and a legacy EffectFnAff callback. Return a concrete Go type for native data, or gopurs_runtime.Value for an explicit legacy action."
  Just c | not (legacyContract c d) ->
    let
      contexts = Array.filter isContext d.args
      args = Array.filter (not <<< isContext) d.args
      results = if unitAcquisition d c.result then [] else valueResults d
      errors = Array.filter isError d.results
    in
      if Array.length args /= Array.length c.args then Just "Go parameters do not match the PureScript arguments (only a leading context.Context is implicit)."
      else if not (Array.null contexts) && (c.kind /= "aff" || Array.length contexts /= 1 || not (fromMaybe false (isContext <$> Array.head d.args))) then Just "Only EffectFnAff accepts one implicit context.Context, as its first Go parameter."
      else if Array.length results > 1 || Array.length errors > 1 || (not (Array.null errors) && not (fromMaybe false (isError <$> Array.last d.results))) then Just "Expected T, (T, error), error, or no Go result."
      else if c.kind == "effect" && not (Array.null errors) then Just "A Go error result requires EffectFnAff; Effect does not select an error policy."
      else if (acquisition d || unitAcquisition d c.result) && c.kind /= "aff" then Just "An acquisition cleanup result requires EffectFnAff."
      else if Array.null results && c.result /= Unit then Just "A Go function without a value result requires a PureScript Unit result."
      else case Array.head (Array.mapMaybe identity (Array.zipWith scalarProblem args c.args <> case Array.head results of
        Just node -> [scalarProblem node c.result]
        Nothing -> [])) of
        Just reason -> Just reason
        Nothing -> Nothing
  _ | Array.length d.results > 1 -> Just "Multiple Go results require a supported EffectFnAff result contract."
  _ -> Nothing

wrapper :: Array DataDecl -> FfiDecl -> ExprType -> Maybe String
wrapper dataDecls d tast = case effectContract tast of
  Just c | not (legacyContract c d) -> Just (render c)
  _ -> Nothing
  where
  render c =
    let
      context = fromMaybe false (isContext <$> Array.head d.args)
      args = if context then Array.drop 1 d.args else d.args
      argTast = Just (Func c.args c.result)
      conversions = String.joinWith "\n" (Array.concat (Array.mapWithIndex (Values.boxedArgument dataDecls argTast) args))
      callArgs = (if context then ["ctx"] else []) <> Array.mapWithIndex (\i _ -> "go_arg" <> show i) args
      call = d.name <> "(" <> String.joinWith ", " callArgs <> ")"
      hasError = fromMaybe false (isError <$> Array.last d.results)
      cleanup = acquisition d || unitAcquisition d c.result
      result = if unitAcquisition d c.result then Nothing else Array.head (valueResults d)
      value = case result of
        Nothing -> "gopurs_runtime.Value{}"
        Just r -> Values.toValue dataDecls r (Just c.result) "go_res"
      invocation = if cleanup then
        (if unitAcquisition d c.result then "go_cleanup, go_err := " else "go_res, go_cleanup, go_err := ") <> call <> "\nif go_err != nil { return gopurs_runtime.Value{}, go_cleanup, go_err }\nreturn " <> value <> ", go_cleanup, nil"
        else case result, hasError of
          Nothing, false -> call <> "\nreturn " <> value <> (if c.kind == "aff" then ", nil" else "")
          Nothing, true -> "go_err := " <> call <> "\nreturn " <> value <> ", go_err"
          Just _, false -> "go_res := " <> call <> "\nreturn " <> value <> (if c.kind == "aff" then ", nil" else "")
          Just _, true -> "go_res, go_err := " <> call <> "\nif go_err != nil { return gopurs_runtime.Value{}, go_err }\nreturn " <> value <> ", nil"
      action = if c.kind == "aff" then
        "gopurs_runtime." <> (if cleanup then "EffectFnAffAcquire" else "EffectFnAff") <> "(func(ctx gopurs_runtime.FFIContext) (gopurs_runtime.Value, " <> (if cleanup then "func(), " else "") <> "error) {\n" <> conversions <> "\n" <> invocation <> "\n})"
        else "gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {\n" <> conversions <> "\n" <> invocation <> "\n})"
      arity = Array.length args
      params = String.joinWith ", " (Array.mapWithIndex (\i _ -> "arg" <> show i <> " gopurs_runtime.Value") args)
    in
      if arity == 0 then action
      else "gopurs_runtime.Func" <> (if arity == 1 then "" else show arity) <> "(func(" <> params <> ") gopurs_runtime.Value {\nreturn " <> action <> "\n})"

unsupported :: TypeNode -> Boolean
unsupported = case _ of
  TUnknown _ -> true
  TMultiFunc _ _ -> true
  TVariadic _ -> true
  TFixedArray _ _ -> true
  TFunc args ret -> Array.any unsupported args || fromMaybe false (unsupported <$> ret)
  TArray t -> unsupported t
  TMap k v -> unsupported k || unsupported v
  _ -> false

structuredProblem :: FfiDecl -> ExprType -> Maybe String
structuredProblem d tast =
  let
    contract = effectContract tast
    argNodes = case contract of
      Just _ -> Array.filter (not <<< isContext) d.args
      _ -> d.args
    argType i = case contract of
      Just c -> Array.index c.args i
      _ -> getTastArgType tast i
    resultType = case contract of
      Just c -> Just c.result
      _ -> getTastReturnType tast
    arguments = Array.mapWithIndex (\i node -> argType i >>= Structured.problem node) argNodes
    result = do
      node <- Array.head (valueResults d)
      ps <- resultType
      Structured.problem node ps
  in Array.head (Array.mapMaybe identity (arguments <> [result]))

-- Check native scalar contracts rather than relying on permissive runtime
-- unboxing (which could otherwise turn a mismatched string into an empty one).
scalarProblem :: TypeNode -> ExprType -> Maybe String
scalarProblem (TIdentified _ _ (Just t)) ps = scalarProblem t ps
scalarProblem (TArray go) (Array ps) = scalarProblem go ps
scalarProblem go ps
  | Array.elem ps [String, Boolean, Int, Number, Unit] =
      if scalarCompatible go ps then Nothing else Just "The native Go scalar does not match the PureScript value type."
  | otherwise = Nothing

scalarCompatible :: TypeNode -> ExprType -> Boolean
scalarCompatible (TNamed "gopurs_runtime.Value") _ = true
scalarCompatible (TNamed "any") _ = true
scalarCompatible (TNamed "interface{}") _ = true
scalarCompatible (TNamed "string") String = true
scalarCompatible (TNamed "bool") Boolean = true
scalarCompatible (TNamed "int") Int = true
scalarCompatible (TNamed "int32") Int = true
scalarCompatible (TNamed "int64") Int = true
scalarCompatible (TNamed "float64") Number = true
scalarCompatible _ _ = false

ambiguousResult :: FfiDecl -> Boolean
ambiguousResult d = case d.ret of
  Just (TNamed "any") -> true
  Just (TNamed "interface{}") -> true
  _ -> false

constrainedContract :: ExprType -> Boolean
constrainedContract (ConstrainedType _ _) = true
constrainedContract (ForAll _ body) = constrainedContract body
constrainedContract (Func args result) = Array.any constrainedContract args || constrainedContract result
constrainedContract _ = false

directAff :: ExprType -> Boolean
directAff (ForAll _ body) = directAff body
directAff (Func _ result) = directAff result
directAff (ADT "Effect.Aff.Aff" _ _) = true
directAff _ = false

-- This shape is only a diagnostic clue. It never selects an asynchronous ABI.
-- Explicit callback implementations keep their ABI. Opaque named Go types
-- cannot be classified here because their underlying type might be a function.
expandedAffCallbacks :: ExprType -> Boolean
expandedAffCallbacks (ForAll _ body) = expandedAffCallbacks body
expandedAffCallbacks (Func _ result) = expandedAffCallbacks result
expandedAffCallbacks (ADT "Effect.Uncurried.EffectFn2" _ [ _, _, ADT "Effect.Uncurried.EffectFn3" _ [ _, _, _, Unit ] ]) = true
expandedAffCallbacks _ = false

ordinaryResult :: TypeNode -> Boolean
ordinaryResult (TIdentified _ _ (Just underlying)) = ordinaryResult underlying
ordinaryResult (TNamed name) = Array.elem name
  [ "bool", "string", "int", "int8", "int16", "int32", "int64"
  , "uint", "uint8", "uint16", "uint32", "uint64", "uintptr"
  , "float32", "float64", "complex64", "complex128", "byte", "rune"
  ]
ordinaryResult (TArray _) = true
ordinaryResult (TMap _ _) = true
ordinaryResult (TPointer _) = true
ordinaryResult (TStruct _ _) = true
ordinaryResult _ = false
