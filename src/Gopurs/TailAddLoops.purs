-- Turn recursion in the right operand of IntAdd into an accumulator loop.
-- The accumulator is private and IntAdd is total/associative modulo 2^32;
-- every observable operand and argument retains its original evaluation order.
module Gopurs.TailAddLoops (worker) where

import Prelude hiding (add)

import Control.Alternative (guard)
import Data.Array as Array
import Data.Array.NonEmpty as NEA
import Data.Foldable (foldMap, foldl)
import Data.Maybe (Maybe(..))
import Data.Newtype (unwrap)
import Data.String as String
import Data.String.Pattern (Pattern(..), Replacement(..))
import Data.Tuple (Tuple(..), fst)
import Gopurs.CallAnalysis (collectGoSpine, getGoSpineArgs)
import Gopurs.ExprAnalysis (unwrapTcoExpr)
import Gopurs.ExprContext (ExprContext, TranslateExpr, childContext, flattenStmts)
import Gopurs.GoAst (GoExpr(..), GoType(..))
import Gopurs.GoConversions (coerceGoExpr)
import Gopurs.GoFunctions (Parameters, iterationBindings)
import PureScript.Backend.Optimizer.Codegen.Tco (TcoExpr(..))
import PureScript.Backend.Optimizer.CoreFn (ExprType(Int), Ident(..), Qualified(..))
import PureScript.Backend.Optimizer.FreeVars (localId)
import PureScript.Backend.Optimizer.Syntax (BackendOperator(..), BackendOperator2(..), BackendOperatorNum(..), BackendSyntax(..), Pair(..))

data Plan
  = Finish TcoExpr
  | Step TcoExpr (Array TcoExpr)
  | Choice (Array (Tuple TcoExpr Plan)) Plan

worker :: TranslateExpr -> ExprContext -> String -> Parameters -> GoType -> TcoExpr -> Maybe GoExpr
worker translate context name params resultType source = do
  guard (resultType == TypeInt64 && not (Array.null params))
  let plan = recognize context.modNameStr name (Array.length params) source
  guard (hasStep plan)
  let
    prefix = freshPrefix (Array.cons name (map fst params) <> bindingNames source)
    acc = prefix <> "acc"
    used = prefix <> "used"
    label = prefix <> "loop"
    emitted = emit translate context params { prefix, acc, used, label } 0 plan
  pure $ GoBlock
    [ GoAssign acc (GoCall (GoVar "int64") [ GoInt 0 ])
    , GoAssign used (GoVar "false")
    , GoFor label (iterationBindings params <> emitted.statements)
    ]

recognize :: String -> String -> Int -> TcoExpr -> Plan
recognize owner name arity source = case unwrapTcoExpr source of
  Branch branches fallback -> Choice
    (map (\(Pair condition value) -> Tuple condition (recognize owner name arity value)) (NEA.toArray branches))
    (recognize owner name arity fallback)
  PrimOp (Op2 (OpIntNum OpAdd) left right) ->
    let Tuple callee spine = collectGoSpine right
        args = getGoSpineArgs spine
    in case unwrapTcoExpr callee of
      Var (Qualified moduleName (Ident ident))
        | ident == name && Array.length args == arity
          && (moduleName == Nothing || map (String.replaceAll (Pattern ".") (Replacement "_") <<< unwrap) moduleName == Just owner) ->
            Step left args
      _ -> Finish source
  _ -> Finish source

hasStep :: Plan -> Boolean
hasStep = case _ of
  Step _ _ -> true
  Choice branches fallback -> Array.any (hasStep <<< \(Tuple _ body) -> body) branches || hasStep fallback
  Finish _ -> false

type Names = { prefix :: String, acc :: String, used :: String, label :: String }
type Emitted = { statements :: Array GoExpr, nextId :: Int }

emit :: TranslateExpr -> ExprContext -> Parameters -> Names -> Int -> Plan -> Emitted
emit translate context@{ codegenStateRef, modNameStr } params names nextId = case _ of
  Choice branches fallback ->
    let
      choices = foldl (\current (Tuple condition body) ->
        let cond = translate (childContext context Nothing) current.nextId condition
            selected = emit translate context params names cond.nextId body
        in { statements: current.statements <> flattenStmts cond.stmts <>
              [ GoIfElse (coerceGoExpr codegenStateRef modNameStr cond.expr cond.exprType TypeBool) selected.statements [] ]
           , nextId: selected.nextId
           }) { statements: [], nextId } branches
      final = emit translate context params names choices.nextId fallback
    in { statements: choices.statements <> final.statements, nextId: final.nextId }
  Finish source ->
    let result = translate (childContext context (Just Int)) nextId source
        valueName = names.prefix <> "result_" <> show result.nextId
        value = GoVar valueName
    in
      { statements: flattenStmts result.stmts <>
          [ GoAssign valueName (native TypeInt64 (coerceGoExpr codegenStateRef modNameStr result.expr result.exprType TypeInt64))
          , GoIfElse (GoVar names.used) [ GoReturn (add (GoVar names.acc) value) ] []
          , GoReturn value
          ]
      , nextId: result.nextId + 1
      }
  Step left args ->
    let
      first = translate (childContext context (Just Int)) nextId left
      leftName = names.prefix <> "left_" <> show first.nextId
      captured = foldl (\current (Tuple (Tuple _ ty) argument) ->
        let result = translate (childContext context Nothing) current.nextId argument
            argName = names.prefix <> "arg_" <> show result.nextId
        in { statements: current.statements <> flattenStmts result.stmts <>
               [ GoAssign argName (native ty (coerceGoExpr codegenStateRef modNameStr result.expr result.exprType ty)) ]
           , values: Array.snoc current.values (GoVar argName)
           , nextId: result.nextId + 1
           })
        { statements: flattenStmts first.stmts <>
            [ GoAssign leftName (native TypeInt64 (coerceGoExpr codegenStateRef modNameStr first.expr first.exprType TypeInt64)) ]
        , values: []
        , nextId: first.nextId + 1
        } (Array.zip params args)
      updates = Array.zipWith (\(Tuple param _) value -> GoMutate (param <> "_loop") value) params captured.values
    in
      { statements: captured.statements <>
          [ GoMutate names.acc (add (GoVar names.acc) (GoVar leftName))
          , GoMutate names.used (GoVar "true")
          ] <> updates <> [ GoContinue names.label ]
      , nextId: captured.nextId
      }

add :: GoExpr -> GoExpr -> GoExpr
add left right = GoCall (GoSelector (GoVar "gopurs_runtime") "IntAdd") [ left, right ]

-- A captured untyped literal must keep the worker ABI, rather than defaulting
-- to Go's machine int in a short declaration.
native :: GoType -> GoExpr -> GoExpr
native TypeInt64 value = GoCall (GoVar "int64") [ value ]
native TypeFloat64 value = GoCall (GoVar "float64") [ value ]
native TypeUint32 value = GoCall (GoVar "uint32") [ value ]
native _ value = value

-- Reserve a complete prefix, including temporary slots and the loop label.
-- The suffix chosen by another emitter cannot collide with this family.
freshPrefix :: Array String -> String
freshPrefix used = choose 0
  where
  choose index =
    let prefix = "__gopurs_tail_add_" <> show index <> "_"
    in if Array.any (String.contains (Pattern prefix)) used then choose (index + 1) else prefix

bindingNames :: TcoExpr -> Array String
bindingNames (TcoExpr _ syntax) = own <> foldMap bindingNames syntax
  where
  refs = map (\(Tuple name level) -> localId name level)
  own = case syntax of
    Local name level -> [ localId name level ]
    Let name level _ _ -> [ localId name level ]
    LetRec level bindings _ -> map (\(Tuple name _) -> localId (Just name) level) (NEA.toArray bindings)
    Abs args _ -> refs (NEA.toArray args)
    UncurriedAbs args _ -> refs args
    UncurriedEffectAbs args _ -> refs args
    EffectBind name level _ _ -> [ localId name level ]
    _ -> []
