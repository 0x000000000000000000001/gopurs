module Gopurs.FfiBridge.ImmediateEffects
  ( key
  , factoryKey
  , signature
  , factorySignature
  , worker
  ) where

import Prelude

import Control.Alternative (guard)
import Data.Array as Array
import Data.Maybe (Maybe(..), fromMaybe)
import Data.String as String
import Data.Traversable (traverse)
import Data.Tuple (Tuple(..))
import Gopurs.CodegenState (FunctionInfo)
import Gopurs.FfiBridge.Effects as Effects
import Gopurs.FfiBridge.TypeSupport (getTastArgType, tastForArity)
import Gopurs.FfiBridge.Values as Values
import Gopurs.FfiTypes (FfiDecl, TypeNode(..))
import Gopurs.GoAst (GoType(..), goTypeToStr)
import Gopurs.GoSymbol (GoReference(..), foreignSymbol)
import PureScript.Backend.Optimizer.CoreFn (DataDecl, ExprType(..))

-- '$' cannot occur in a PureScript value identifier. Keep these signatures in
-- the existing worker table without making them ordinary callable bindings.
key :: String -> String
key name = "$gopursRunEffect$" <> name

factoryKey :: String -> String
factoryKey name = "$gopursEffectFactory$" <> name

type Candidate = { result :: ExprType, nativeResult :: Maybe TypeNode }

candidate :: FfiDecl -> ExprType -> Maybe Candidate
candidate declaration tast = do
  contract <- Effects.effectContract tast
  guard (contract.kind == "effect" && not declaration.isVar && Array.null declaration.typeParams)
  guard (Array.length declaration.args > 0 && Array.length declaration.args == Array.length contract.args)
  guard (Effects.problem declaration tast == Nothing)
  case declaration.ret, declaration.results of
    Just (TFunc [] result), [ _ ] -> Just { result: contract.result, nativeResult: result }
    _, _ -> Nothing

workerName :: FfiDecl -> String
workerName declaration = declaration.name <> "_runEffectWorker"

-- A polymorphic argument already crosses this native interface as Value.
-- Let the caller retain that conversion; Foreign and structured arguments
-- still need their ordinary bridge adaptation inside the worker.
argumentType :: ExprType -> Int -> TypeNode -> GoType
argumentType tast index native = case native, getTastArgType tast index of
  TNamed name, Just (TypeVar _) | name == "any" || name == "interface{}" -> TypeInterface "interface{}"
  _, _ -> TypeValue

directArgumentType :: ExprType -> Int -> TypeNode -> Maybe GoType
directArgumentType tast index native = case native, getTastArgType tast index of
  TNamed name, Just (TypeVar _) | name == "any" || name == "interface{}" -> Just (TypeInterface "interface{}")
  TNamed "gopurs_runtime.Value", _ -> Just TypeValue
  TNamed "int64", Just Int -> Just TypeInt64
  TNamed "float64", Just Number -> Just TypeFloat64
  TNamed "bool", Just Boolean -> Just TypeBool
  TNamed "string", Just String -> Just TypeString
  _, _ -> Nothing

-- Direct invocation must preserve Values.toValue without module dataDecls
-- (so never admit ADTs/newtypes/records/functions through an `any` boundary).
directReturnType :: Maybe TypeNode -> ExprType -> Maybe GoType
directReturnType native result = case native, result of
  Just (TNamed "gopurs_runtime.Value"), _ -> Just TypeValue
  Just (TNamed name), _ | (name == "any" || name == "interface{}") && simpleBoxable result -> Just (TypeInterface "interface{}")
  Just (TNamed "int64"), Int -> Just TypeInt64
  Just (TNamed "float64"), Number -> Just TypeFloat64
  Just (TNamed "bool"), Boolean -> Just TypeBool
  Just (TNamed "string"), String -> Just TypeString
  _, _ -> Nothing
  where
  simpleBoxable = case _ of
    TypeVar _ -> true
    Int -> true
    Number -> true
    String -> true
    Boolean -> true
    Char -> true
    Unit -> true
    _ -> false

signature :: String -> FfiDecl -> ExprType -> Maybe FunctionInfo
signature owner declaration tast = do
  _ <- candidate declaration tast
  let callTast = fromMaybe tast (tastForArity (Array.length declaration.args) tast)
  pure
    { arrayFold: Nothing, codePointFold: Nothing, reference: SymbolReference (foreignSymbol owner (workerName declaration))
    , fArgs: Array.mapWithIndex (argumentType callTast) declaration.args
    , fRet: TypeValue
    , arity: Array.length declaration.args
    }

factorySignature :: String -> FfiDecl -> ExprType -> Maybe FunctionInfo
factorySignature owner declaration tast = do
  c <- candidate declaration tast
  let callTast = fromMaybe tast (tastForArity (Array.length declaration.args) tast)
  fArgs <- traverse (\(Tuple index native) -> directArgumentType callTast index native) (Array.mapWithIndex Tuple declaration.args)
  ret <- directReturnType c.nativeResult c.result
  pure
    { arrayFold: Nothing, codePointFold: Nothing, reference: SymbolReference (foreignSymbol owner declaration.name)
    , fArgs
    , fRet: TypeFunc [] ret
    , arity: Array.length declaration.args
    }

-- Only an immediately executed action uses this worker. Its original boxed
-- entry remains authoritative for saved actions and first-class functions.
-- The factory is still called before its thunk, and dynamic native results
-- retain Values.toValue's ordinary Box fallback (they need not be Value).
worker :: Array DataDecl -> FfiDecl -> ExprType -> Maybe String
worker dataDecls declaration tast = do
  c <- candidate declaration tast
  let
    callTast = fromMaybe tast (tastForArity (Array.length declaration.args) tast)
    params = String.joinWith ", " (Array.mapWithIndex (\i native -> "arg" <> show i <> " " <> goTypeToStr (argumentType callTast i native)) declaration.args)
    convert i native = case argumentType callTast i native of
      TypeInterface _ -> [ "go_arg" <> show i <> " := arg" <> show i ]
      _ -> Values.boxedArgument dataDecls (Just callTast) i native
    conversions = String.joinWith "\n" (Array.concat (Array.mapWithIndex convert declaration.args))
    args = String.joinWith ", " (Array.mapWithIndex (\i _ -> "go_arg" <> show i) declaration.args)
    result = case c.nativeResult of
      Nothing -> "go_action()\nreturn gopurs_runtime.Value{}"
      Just native -> "go_result := go_action()\nreturn " <> Values.toValue dataDecls native (Just c.result) "go_result"
  pure ("func " <> workerName declaration <> "(" <> params <> ") gopurs_runtime.Value {\n"
    <> conversions <> "\ngo_action := " <> declaration.name <> "(" <> args <> ")\n" <> result <> "\n}\n")
