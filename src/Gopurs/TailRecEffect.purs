module Gopurs.TailRecEffect (emit) where

import Prelude

import Control.Alternative (guard)
import Data.Array as Array
import Data.Maybe (Maybe(..), fromMaybe)
import Data.Tuple (Tuple(..))
import Gopurs.AdtExprs as AdtExprs
import Gopurs.CallAnalysis (qualifiedTarget)
import Gopurs.CallArguments as CallArguments
import Gopurs.ExprAnalysis (unwrapTcoExpr)
import Gopurs.ExprContext (ExprContext, ExprResult, StmtTree(..), TranslateExpr, childContext)
import Gopurs.GoAst (rawGo, GoExpr(..), GoType(..))
import Gopurs.GoConversions (boxGoExpr, coerceGoExpr)
import Gopurs.TailRecEffectStep as Step
import PureScript.Backend.Optimizer.Codegen.Tco (TcoExpr)
import PureScript.Backend.Optimizer.CoreFn (ExprType(..), ModuleName(..))
import PureScript.Backend.Optimizer.Syntax (BackendAccessor(GetProp), BackendSyntax(Accessor))

-- These two standard instances run synchronous effects sequentially. Their
-- implementation's private Step reference is not observable by the callback,
-- so it can be replaced by a loop-local value. User references keep their
-- ordinary representation and synchronization.
emit :: TranslateExpr -> ExprContext -> Int -> TcoExpr -> Array TcoExpr -> Maybe ExprResult
emit translate context nextId fn args = do
  supplied <- recognize fn args
  pure case supplied of
    [ dictionary, callback, seed ] | Just plan <- Step.recognize context callback ->
      emitInline translate context nextId dictionary seed plan
    _ -> emitOpaque translate context nextId supplied

emitInline :: TranslateExpr -> ExprContext -> Int -> TcoExpr -> TcoExpr -> Step.Plan -> ExprResult
emitInline translate context@{ codegenStateRef, modNameStr } nextId dictionary seed plan =
  let
    captured = CallArguments.capture { prefix: "tailRecEffect_arg_", markUsed: true }
      (\index currentId arg ->
        let
          result = translate (childContext context Nothing) currentId arg
          target = if index == 0 then TypeValue else plan.stateType
        in result { expr = coerceGoExpr codegenStateRef modNameStr result.expr result.exprType target, exprType = target })
      nextId [ dictionary, seed ]
    stateName = "tailRecEffect_state_" <> show captured.nextId
    loopName = "tailRecEffect_loop_" <> show captured.nextId
    body = Step.emit translate context captured.nextId stateName loopName plan
    initial = fromMaybe (rawGo "nil") (Array.index captured.exprs 1)
  in
    { stmts: captured.stmts
    , expr: GoCall (GoSelector (GoVar "gopurs_runtime") "Func")
        [ GoFuncBlock [ Tuple "_" TypeValue ]
            [ GoAssign stateName initial, GoFor loopName body.statements ] TypeValue ]
    , exprType: TypeValue
    , nextId: body.nextId
    }

emitOpaque :: TranslateExpr -> ExprContext -> Int -> Array TcoExpr -> ExprResult
emitOpaque translate context@{ metadata, codegenStateRef, modNameStr } nextId supplied =
  let
    captured = CallArguments.capture { prefix: "tailRecEffect_arg_", markUsed: true }
      (\_ currentId arg ->
        let result = translate (childContext context Nothing) currentId arg
        in result { expr = boxGoExpr codegenStateRef modNameStr result.expr result.exprType, exprType = TypeValue })
      nextId supplied
    argument index = fromMaybe (rawGo "nil") (Array.index captured.exprs index)
    suffix = show captured.nextId
    stateName = "tailRecEffect_state_" <> suffix
    stepName = "tailRecEffect_step_" <> suffix
    loopName = "tailRecEffect_loop_" <> suffix
    stepModule = Just (ModuleName "Control.Monad.Rec.Class")
    stepType = ADT "Control.Monad.Rec.Class.Step" [ "Control", "Monad", "Rec", "Class", "Step" ] [ Any, Any ]
    step = { expr: GoVar stepName, exprType: TypeValue, stmts: StmtEmpty, nextId: captured.nextId }
    tag name = (AdtExprs.isTag metadata codegenStateRef modNameStr stepModule name step).expr
    payload name =
      let value = AdtExprs.getField metadata codegenStateRef modNameStr
            { moduleName: stepModule, ctorName: name, index: 0 }
            { expr: GoVar stepName, exprType: TypeValue, sourceType: stepType }
      in boxGoExpr codegenStateRef modNameStr value.expr value.exprType
    apply f a = GoCall (GoSelector (GoVar "gopurs_runtime") "Apply") [ f, a ]
    body =
      [ GoAssign stateName (argument 2)
      , GoFor loopName
          [ GoAssign stepName (apply (apply (argument 1) (GoVar stateName)) (rawGo "gopurs_runtime.Value{}"))
          , GoIfElse (tag "Loop") [ GoMutate stateName (payload "Loop"), GoContinue loopName ] []
          , GoIfElse (tag "Done") [ GoReturn (payload "Done") ] []
          , GoCall (GoVar "panic") [ GoString "Failed pattern match" ]
          ]
      ]
  in
    { stmts: captured.stmts
    , expr: GoCall (GoSelector (GoVar "gopurs_runtime") "Func") [ GoFuncBlock [ Tuple "_" TypeValue ] body TypeValue ]
    , exprType: TypeValue
    , nextId: captured.nextId
    }

-- Only saturated applications with a proven standard instance qualify. A
-- user-defined MonadRec, partial application or extra argument uses the usual
-- call path, preserving its own evaluation and effect semantics.
recognize :: TcoExpr -> Array TcoExpr -> Maybe (Array TcoExpr)
recognize fn args = case unwrapTcoExpr fn of
  Accessor dictionary (GetProp "tailRecM") -> do
    guard (knownInstance dictionary && Array.length args == 2)
    pure (Array.cons dictionary args)
  _ -> do
    guard (named "Control.Monad.Rec.Class" "tailRecM" fn && Array.length args == 3)
    dictionary <- Array.head args
    guard (knownInstance dictionary)
    pure args

knownInstance :: TcoExpr -> Boolean
knownInstance expr = named "Control.Monad.Rec.Class" "monadRecEffect" expr
  || named "Control.Monad.ST.Internal" "monadRecST" expr

named :: String -> String -> TcoExpr -> Boolean
named moduleName name expr = case qualifiedTarget expr of
  Just target -> target.mbMod == Just (ModuleName moduleName) && target.name == name
  Nothing -> false
