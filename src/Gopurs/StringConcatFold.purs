module Gopurs.StringConcatFold (emit) where

import Prelude

import Control.Alternative (guard)
import Control.Monad.State (evalState)
import Data.Array as Array
import Data.Array.NonEmpty as NEA
import Data.Maybe (Maybe(..), fromMaybe)
import Data.Tuple (Tuple(..))
import Gopurs.ArrayIntrinsics.Source (arraySource)
import Gopurs.ArrayProducerFold as ArrayProducerFold
import Gopurs.ArrayTraverseTotal (totalInt)
import Gopurs.CallAnalysis (collectGoSpine, getGoSpineArgs, qualifiedTarget)
import Gopurs.CallArguments as CallArguments
import Gopurs.CodePointFold as CodePointFold
import Gopurs.CodegenState (CodePointFoldDescriptor)
import Gopurs.ExprContext (ExprContext, ExprResult, StmtTree(..), TranslateExpr, bindParameters, childContext, flattenStmts)
import Gopurs.GoAst (GoExpr(..), GoType(..), rawGo)
import Gopurs.GoConversions (coerceGoExpr)
import Gopurs.ImmediateApplications.Scope (freshen, maximumLevel)
import PureScript.Backend.Optimizer.Codegen.Tco (TcoExpr(..))
import PureScript.Backend.Optimizer.Codegen.Tco as Tco
import PureScript.Backend.Optimizer.CoreFn (ExprType(Int, String, Array, Func), Literal(..), ModuleName(..))
import PureScript.Backend.Optimizer.FreeVars (localId)
import PureScript.Backend.Optimizer.Semantics (NeutralExpr(..))
import PureScript.Backend.Optimizer.Syntax (BackendSyntax(..))

-- Only a closed, total code-point consumer may run before the canonical string
-- producer finishes. Captured producer arguments retain their strict order.
-- An ASCII prefix is already a sequence of complete code points. At the first
-- other byte, retain that accumulator and concatenate/decode only the suffix;
-- this preserves WTF8 normalization and arbitrary malformed byte boundaries.
emit :: TranslateExpr -> ExprContext -> Int -> CodePointFoldDescriptor -> TcoExpr -> Maybe ExprResult
emit translate context nextId descriptor producer = do
  guard (Array.null context.recVars && CodePointFold.validDescriptor descriptor && ArrayProducerFold.onlyCurried producer)
  guard (CodePointFold.compatibleT String producer)
  let Tuple function spine = collectGoSpine producer
  guard (named "Data.Foldable" "foldlArray" function)
  guard (CodePointFold.compatibleT (Func [ Func [ String, String ] String, String, Array String ] String) function)
  { concat, initial, input } <- case getGoSpineArgs spine of
    [ concat, initial, input ] -> Just { concat, initial, input }
    _ -> Nothing
  guard (named "Data.Semigroup" "concatString" concat)
  guard (CodePointFold.compatibleT (Func [ String, String ] String) concat)
  guard (CodePointFold.compatibleT String initial && CodePointFold.compatibleT (Array String) input)
  let
    start = max (maximumLevel (neutral producer)) (maximumLevel descriptor.callback) + nextId + 1
    callback = Tco.analyze [] (evalState (freshen descriptor.callback) start)
    seed = Tco.analyze [] (NeutralExpr (Typed Int (NeutralExpr (Lit (LitInt descriptor.seed)))))
  grouped <- CodePointFold.primitiveCallback callback
  guard (totalInt (map (\(Tuple _ level) -> level) (NEA.toArray grouped.args)) grouped.body)
  let
    params = map (\(Tuple name level) -> Tuple (localId name level) TypeInt64) (NEA.toArray grouped.args)
    callbackContext = context
      { depth = context.depth + 1, bound = bindParameters params context.bound
      , tcoIdent = Nothing, loopCtx = [], options = { isTail: true, inEffectBlock: false }
      , mbExpectedExprType = Just Int, mbExpectedGoType = Just TypeInt64 }
    coerce expression from to = coerceGoExpr context.codegenStateRef context.modNameStr expression from to
    translatedBody = translate callbackContext nextId grouped.body
    bodyStmts = flattenStmts translatedBody.stmts
    bodyExpr = coerce translatedBody.expr translatedBody.exprType TypeInt64
    accParam = fromMaybe "" (map (\(Tuple name _) -> name) (Array.index params 0))
    cpParam = fromMaybe "" (map (\(Tuple name _) -> name) (Array.index params 1))
    directAdd = Array.null bodyStmts && case bodyExpr of
      GoCall (GoSelector (GoVar "gopurs_runtime") "IntAdd") [ GoVar a, GoVar b ] ->
        (a == accParam && b == cpParam) || (a == cpParam && b == accParam)
      _ -> false
    -- A constant seed already in signed Int32 range satisfies int64(int32(seed)) == seed
    -- even on empty input, so the hot ASCII loop needs no per-string non-empty tracking.
    narrowSeed = descriptor.seed >= (bottom :: Int) && descriptor.seed <= (top :: Int)
    trackAny = directAdd && not narrowSeed
    capture index currentId value = if index == 0 then
      { stmts: StmtEmpty, expr: GoFuncLit params bodyStmts bodyExpr TypeInt64
      , exprType: TypeFunc [ TypeInt64, TypeInt64 ] TypeInt64, nextId: translatedBody.nextId }
      else translate ((childContext context Nothing) { mbExpectedGoType = if index == 4 then Just TypeValue else Nothing }) currentId value
    captured = CallArguments.capture { prefix: "concatFold_arg_", markUsed: true } capture nextId [ callback, seed, concat, initial, input ]
    arg index = fromMaybe (rawGo "nil") (Array.index captured.exprs index)
    argType index = fromMaybe TypeValue (Array.index captured.exprTypes index)
    source = arraySource "concatFold_source" "concatFold_buffer" (argType 4)
    projection = GoFuncLit [ Tuple "concatFold_value" source.elementType ] []
      (coerce (GoVar "concatFold_value") source.elementType TypeString) TypeString
    -- Pass acc (and wide seed if needed) by value when directAdd holds so
    -- concatFold_finish captures no mutable locals and concatFold_acc stays in a register.
    finishParams =
      [ Tuple "concatFold_tail" TypeString, Tuple "concatFold_remaining" (TypeNativeArray source.elementType) ]
        <> (if directAdd then [ Tuple "concatFold_acc_in" TypeInt64 ] else [])
        <> (if trackAny then [ Tuple "concatFold_seed_in" TypeInt64 ] else [])
    finishStep value =
      if directAdd then
        GoMutate "concatFold_acc_in" (GoBinOp "+" (GoVar "concatFold_acc_in") value)
      else
        GoMutate "concatFold_result" (GoCall (arg 0) [ GoVar "concatFold_result", value ])
    finishResult =
      if trackAny then
        GoCall (GoVar "int64") [ GoCall (GoVar "int32") [ GoBinOp "+" (GoVar "concatFold_seed_in") (GoVar "concatFold_acc_in") ] ]
      else if directAdd then
        GoCall (GoVar "int64") [ GoCall (GoVar "int32") [ GoVar "concatFold_acc_in" ] ]
      else
        GoVar "concatFold_result"
    finish = GoFuncLit finishParams
      [ GoAssign "concatFold_joined" (GoCall (GoSelector (GoVar "gopurs_runtime") "ConcatStrings")
          [ GoVar "concatFold_tail", GoVar "concatFold_remaining", projection ])
      , GoForRange "concatFold_offset := 0; concatFold_offset < len(concatFold_joined);"
          [ GoAssign "concatFold_point" (rawGo "rune(concatFold_joined[concatFold_offset])")
          , GoAssign "concatFold_width" (GoInt 1)
          , GoIfElse (rawGo "concatFold_point >= 128")
              [ GoMutate "concatFold_point, concatFold_width" (GoCall (GoSelector (GoVar "gopurs_runtime") "DecodeStringCodePoint")
                  [ rawGo "concatFold_joined[concatFold_offset:]" ]) ] []
          , finishStep (rawGo "int64(concatFold_point)")
          , GoMutate "concatFold_offset" (rawGo "concatFold_offset + concatFold_width")
          ]
      ] finishResult TypeInt64
    callFinish tailExpr remaining =
      GoCall (GoVar "concatFold_finish")
        ([ tailExpr, remaining ]
          <> (if directAdd then [ GoVar "concatFold_acc" ] else [])
          <> (if trackAny then [ GoVar "concatFold_result" ] else []))
    asciiLoop text remaining =
      GoForRange ("concatFold_offset := 0; concatFold_offset < len(" <> text <> "); concatFold_offset++")
        [ GoAssign "concatFold_b" (rawGo (text <> "[concatFold_offset]"))
        , GoIfElse (rawGo "concatFold_b >= 128")
            [ GoReturn (callFinish (rawGo (text <> "[concatFold_offset:]")) remaining) ] []
        , GoMutate "concatFold_acc" (rawGo "concatFold_acc + int64(concatFold_b)")
        ]
    ascii text remaining =
      if trackAny then
        GoIfElse (rawGo ("len(" <> text <> ") > 0"))
          [ GoMutate "concatFold_any" (rawGo "true")
          , asciiLoop text remaining
          ] []
      else if directAdd then
        asciiLoop text remaining
      else
        GoForRange ("concatFold_offset := 0; concatFold_offset < len(" <> text <> "); concatFold_offset++")
          [ GoIfElse (rawGo (text <> "[concatFold_offset] >= 128"))
              [ GoReturn (callFinish (rawGo (text <> "[concatFold_offset:]")) remaining) ] []
          , GoMutate "concatFold_result" (GoCall (arg 0) [ GoVar "concatFold_result", rawGo ("int64(" <> text <> "[concatFold_offset])") ])
          ]
    accInit =
      if trackAny then [ GoAssign "concatFold_acc" (rawGo "int64(0)"), GoAssign "concatFold_any" (rawGo "false") ]
      else if directAdd then [ GoAssign "concatFold_acc" (GoVar "concatFold_result") ]
      else []
    accFinish =
      if trackAny then
        [ GoIfElse (GoVar "concatFold_any")
            [ GoMutate "concatFold_result" (GoCall (GoVar "int64") [ GoCall (GoVar "int32") [ GoBinOp "+" (GoVar "concatFold_result") (GoVar "concatFold_acc") ] ]) ] []
        ]
      else if directAdd then
        [ GoMutate "concatFold_result" (GoCall (GoVar "int64") [ GoCall (GoVar "int32") [ GoVar "concatFold_acc" ] ]) ]
      else []
    body =
      [ GoAssign "concatFold_source" (arg 4), source.assignment
      , GoAssign "concatFold_initial" (coerce (arg 3) (argType 3) TypeString)
      , GoAssign "concatFold_values" (rawGo source.target)
      , GoAssign "concatFold_result" (coerce (arg 1) (argType 1) TypeInt64)
      ] <> accInit <>
      [ GoAssign "concatFold_finish" finish
      , ascii "concatFold_initial" (GoVar "concatFold_values")
      , GoForRange "concatFold_index, concatFold_value := range concatFold_values"
          [ GoAssign "concatFold_text" (coerce (GoVar "concatFold_value") source.elementType TypeString)
          , ascii "concatFold_text" (rawGo "concatFold_values[concatFold_index+1:]")
          ]
      ] <> accFinish
  pure { stmts: captured.stmts, expr: GoCall (GoFuncLit [] body (GoVar "concatFold_result") TypeInt64) []
       , exprType: TypeInt64, nextId: captured.nextId }

named :: String -> String -> TcoExpr -> Boolean
named moduleName name expression = case qualifiedTarget expression of
  Just target -> target.mbMod == Just (ModuleName moduleName) && target.name == name
  Nothing -> false

neutral :: TcoExpr -> NeutralExpr
neutral (TcoExpr _ syntax) = NeutralExpr (map neutral syntax)
