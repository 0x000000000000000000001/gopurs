module Gopurs.FfiBridge.ImmediateEffects
  ( key
  , signature
  , worker
  ) where

import Prelude

import Control.Alternative (guard)
import Data.Array as Array
import Data.Maybe (Maybe(..), fromMaybe)
import Data.String as String
import Gopurs.CodegenState (FunctionInfo)
import Gopurs.FfiBridge.Effects as Effects
import Gopurs.FfiBridge.TypeSupport (tastForArity)
import Gopurs.FfiBridge.Values as Values
import Gopurs.FfiTypes (FfiDecl, TypeNode(..))
import Gopurs.GoAst (GoType(..))
import Gopurs.GoSymbol (GoReference(..), foreignSymbol)
import PureScript.Backend.Optimizer.CoreFn (DataDecl, ExprType)

-- '$' cannot occur in a PureScript value identifier. Keep these signatures in
-- the existing worker table without making them ordinary callable bindings.
key :: String -> String
key name = "$gopursRunEffect$" <> name

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

signature :: String -> FfiDecl -> ExprType -> Maybe FunctionInfo
signature owner declaration tast = do
  _ <- candidate declaration tast
  pure
    { reference: SymbolReference (foreignSymbol owner (workerName declaration))
    , fArgs: map (const TypeValue) declaration.args
    , fRet: TypeValue
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
    params = String.joinWith ", " (Array.mapWithIndex (\i _ -> "arg" <> show i <> " gopurs_runtime.Value") declaration.args)
    conversions = String.joinWith "\n" (Array.concat (Array.mapWithIndex (Values.boxedArgument dataDecls (Just callTast)) declaration.args))
    args = String.joinWith ", " (Array.mapWithIndex (\i _ -> "go_arg" <> show i) declaration.args)
    result = case c.nativeResult of
      Nothing -> "go_action()\nreturn gopurs_runtime.Value{}"
      Just native -> "go_result := go_action()\nreturn " <> Values.toValue dataDecls native (Just c.result) "go_result"
  pure ("func " <> workerName declaration <> "(" <> params <> ") gopurs_runtime.Value {\n"
    <> conversions <> "\ngo_action := " <> declaration.name <> "(" <> args <> ")\n" <> result <> "\n}\n")
