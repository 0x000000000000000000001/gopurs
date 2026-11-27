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
    capture index currentId value = if index == 0 then
      let body = translate callbackContext currentId grouped.body
      in { stmts: StmtEmpty, expr: GoFuncLit params (flattenStmts body.stmts) (coerce body.expr body.exprType TypeInt64) TypeInt64
         , exprType: TypeFunc [ TypeInt64, TypeInt64 ] TypeInt64, nextId: body.nextId }
      else translate ((childContext context Nothing) { mbExpectedGoType = if index == 4 then Just TypeValue else Nothing }) currentId value
    captured = CallArguments.capture { prefix: "concatFold_arg_", markUsed: true } capture nextId [ callback, seed, concat, initial, input ]
    arg index = fromMaybe (rawGo "nil") (Array.index captured.exprs index)
    argType index = fromMaybe TypeValue (Array.index captured.exprTypes index)
    source = arraySource "concatFold_source" "concatFold_buffer" (argType 4)
    step value = GoMutate "concatFold_result" (GoCall (arg 0) [ GoVar "concatFold_result", value ])
    projection = GoFuncLit [ Tuple "concatFold_value" source.elementType ] []
      (coerce (GoVar "concatFold_value") source.elementType TypeString) TypeString
    finish = GoFuncLit [ Tuple "concatFold_tail" TypeString, Tuple "concatFold_remaining" (TypeNativeArray source.elementType) ]
      [ GoAssign "concatFold_joined" (GoCall (GoSelector (GoVar "gopurs_runtime") "ConcatStrings")
          [ GoVar "concatFold_tail", GoVar "concatFold_remaining", projection ])
      , GoForRange "concatFold_offset := 0; concatFold_offset < len(concatFold_joined);"
          [ GoAssign "concatFold_point" (rawGo "rune(concatFold_joined[concatFold_offset])")
          , GoAssign "concatFold_width" (GoInt 1)
          , GoIfElse (rawGo "concatFold_point >= 128")
              [ GoMutate "concatFold_point, concatFold_width" (GoCall (GoSelector (GoVar "gopurs_runtime") "DecodeStringCodePoint")
                  [ rawGo "concatFold_joined[concatFold_offset:]" ]) ] []
          , step (rawGo "int64(concatFold_point)")
          , GoMutate "concatFold_offset" (rawGo "concatFold_offset + concatFold_width")
          ]
      ] (GoVar "concatFold_result") TypeInt64
    ascii text remaining = GoForRange ("concatFold_offset := 0; concatFold_offset < len(" <> text <> "); concatFold_offset++")
      [ GoIfElse (rawGo (text <> "[concatFold_offset] >= 128"))
          [ GoReturn (GoCall (GoVar "concatFold_finish") [ rawGo (text <> "[concatFold_offset:]"), remaining ]) ] []
      , step (rawGo ("int64(" <> text <> "[concatFold_offset])"))
      ]
    body =
      [ GoAssign "concatFold_source" (arg 4), source.assignment
      , GoAssign "concatFold_initial" (coerce (arg 3) (argType 3) TypeString)
      , GoAssign "concatFold_values" (rawGo source.target)
      , GoAssign "concatFold_result" (coerce (arg 1) (argType 1) TypeInt64)
      , GoAssign "concatFold_finish" finish
      , ascii "concatFold_initial" (GoVar "concatFold_values")
      , GoForRange "concatFold_index, concatFold_value := range concatFold_values"
          [ GoAssign "concatFold_text" (coerce (GoVar "concatFold_value") source.elementType TypeString)
          , ascii "concatFold_text" (rawGo "concatFold_values[concatFold_index+1:]")
          ]
      ]
  pure { stmts: captured.stmts, expr: GoCall (GoFuncLit [] body (GoVar "concatFold_result") TypeInt64) []
       , exprType: TypeInt64, nextId: captured.nextId }

named :: String -> String -> TcoExpr -> Boolean
named moduleName name expression = case qualifiedTarget expression of
  Just target -> target.mbMod == Just (ModuleName moduleName) && target.name == name
  Nothing -> false

neutral :: TcoExpr -> NeutralExpr
neutral (TcoExpr _ syntax) = NeutralExpr (map neutral syntax)
