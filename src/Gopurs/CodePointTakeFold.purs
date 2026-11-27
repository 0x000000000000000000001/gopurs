module Gopurs.CodePointTakeFold (emit) where

import Prelude

import Control.Alternative (guard)
import Control.Monad.State (evalState)
import Data.Array as Array
import Data.Maybe (Maybe(..))
import Data.Tuple (Tuple(..))
import Gopurs.CallAnalysis (collectGoSpine, getGoSpineArgs, qualifiedTarget)
import Gopurs.CodePointFold as CodePointFold
import Gopurs.CodegenState (CodePointFoldDescriptor)
import Gopurs.ExprContext (ExprContext, ExprResult, TranslateExpr)
import Gopurs.ImmediateApplications.Scope (freshen, maximumLevel)
import PureScript.Backend.Optimizer.Codegen.Tco (TcoExpr(..))
import PureScript.Backend.Optimizer.Codegen.Tco as Tco
import PureScript.Backend.Optimizer.CoreFn (ExprType(Int, String, Func), Literal(..), ModuleName(..))
import PureScript.Backend.Optimizer.Semantics (NeutralExpr(..))
import PureScript.Backend.Optimizer.Syntax (BackendSyntax(..))

-- Remove only the immediately consumed canonical take. The count and string
-- are captured in their original order, even for empty/negative prefixes.
-- Neither a stored substring nor an unknown/captured consumer is rewritten.
emit :: TranslateExpr -> ExprContext -> Int -> CodePointFoldDescriptor -> TcoExpr -> Maybe ExprResult
emit translate context nextId descriptor producer = do
  guard (Array.null context.recVars && CodePointFold.validDescriptor descriptor && onlyCurried producer)
  guard (CodePointFold.compatibleT String producer)
  let Tuple function spine = collectGoSpine producer
  target <- qualifiedTarget function
  guard (target.mbMod == Just (ModuleName "Data.String.CodePoints") && target.name == "take")
  guard (CodePointFold.compatibleT (Func [ Int, String ] String) function)
  Tuple count input <- case getGoSpineArgs spine of
    [ count, input ] -> Just (Tuple count input)
    _ -> Nothing
  let
    start = max (maximumLevel (neutral producer)) (maximumLevel descriptor.callback) + nextId + 1
    callback = Tco.analyze [] (evalState (freshen descriptor.callback) start)
    seed = Tco.analyze [] (NeutralExpr (Typed Int (NeutralExpr (Lit (LitInt descriptor.seed)))))
  CodePointFold.emitBounded translate context nextId callback seed count input

onlyCurried :: TcoExpr -> Boolean
onlyCurried (TcoExpr _ syntax) = case syntax of
  App function _ -> onlyCurried function
  Typed _ value -> onlyCurried value
  TypeApp value _ -> onlyCurried value
  UncurriedApp _ _ -> false
  UncurriedEffectApp _ _ -> false
  _ -> true

neutral :: TcoExpr -> NeutralExpr
neutral (TcoExpr _ syntax) = NeutralExpr (map neutral syntax)
