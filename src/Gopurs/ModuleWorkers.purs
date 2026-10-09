module Gopurs.ModuleWorkers (declaration) where

import Prelude
import Data.Array as Array
import Data.Map as Map
import Data.Maybe (Maybe(..))
import Data.Tuple (Tuple(..), fst, snd)
import Effect.Ref as Ref
import Effect.Unsafe (unsafePerformEffect)
import Gopurs.ExprAnalysis (extractExprFuncType, functionResultAfter, getExprType)
import Gopurs.ExprContext (ExprContext, TranslateExpr, bindParameters, flattenStmts)
import Gopurs.GoAst (GoDecl(..), GoExpr(..), GoType(..))
import Gopurs.GoConversions (boxGoExpr, coerceGoExpr)
import Gopurs.GoFunctions (Parameters, curriedFunction, iterationBindings, loopParameters, namedParameters)
import Gopurs.Int32Loops as Int32Loops
import PureScript.Backend.Optimizer.Codegen.Tco (TcoExpr)

type Worker =
  { ident :: String
  , args :: Array String
  , body :: TcoExpr
  , val :: TcoExpr
  }

-- Use precisely the ABI published by ModuleBindings, including native record
-- projections. A self-recursive singleton gets a loop; peers in larger groups
-- remain ordinary calls. Module bodies begin with no captured local bindings.
declaration :: TranslateExpr -> ExprContext -> Boolean -> Worker -> GoDecl
declaration translate context@{ codegenStateRef, modNameStr, moduleFunctions } isSelfRecursive fn =
  let
    info = Map.lookup fn.ident moduleFunctions
    params = namedParameters fn.args (case info of
      Just signature -> signature.fArgs
      Nothing -> [])
    expectedBodyType = map (functionResultAfter (Array.length fn.args)) (extractExprFuncType (getExprType fn.val))
    loop = if isSelfRecursive then
      [ { ident: fn.ident, loopParams: map fst (loopParameters params), goTypes: map snd params } ]
      else []
    result = translate (context
      { depth = 0
      , bound = bindParameters params Map.empty
      , tcoIdent = Just fn.ident
      , loopCtx = loop
      , options = { isTail: isSelfRecursive, inEffectBlock: false }
      , mbExpectedExprType = expectedBodyType
      }) 0 fn.body
    resultType = case info of
      Just signature | Array.length fn.args < Array.length signature.fArgs -> TypeValue
      Just signature -> signature.fRet
      Nothing -> TypeValue
    body returning =
      let statements = flattenStmts result.stmts <> [ GoReturn returning ]
      in if isSelfRecursive then Int32Loops.loop fn.ident params resultType statements
        else GoBlock (iterationBindings params <> statements)
    value = if Array.null fn.args then
      -- A zero-argument abstraction still defers its whole body until called.
      curriedFunction [ Tuple "_" TypeValue ] TypeValue
        (body (boxGoExpr codegenStateRef modNameStr result.expr result.exprType))
      else
        let
          worker = GoFunctionDecl
            { name: "Call_" <> modNameStr <> "_" <> fn.ident
            , params: loopParameters params
            , result: resultType
            , body: body (coerceGoExpr codegenStateRef modNameStr result.expr result.exprType resultType)
            }
        in
          unsafePerformEffect do
            Ref.modify_ (\state -> state { declarations = Array.snoc state.declarations worker }) codegenStateRef
            pure (boxedWrapper context fn.ident params resultType)
  in
    GoCachedValue { identifier: modNameStr <> "_" <> fn.ident, expression: value, goType: TypeValue }

-- Module wrappers group up to ten arguments; larger arities keep unary layers.
-- LocalWorkers and anonymous closures retain their own grouping policies.
boxedWrapper :: ExprContext -> String -> Parameters -> GoType -> GoExpr
boxedWrapper { codegenStateRef, modNameStr } name params resultType =
  let
    wrapperParams = map (\(Tuple param _) -> param <> "_box") params
    call = GoCall (GoVar ("Call_" <> modNameStr <> "_" <> name))
      (map (\(Tuple param ty) -> coerceGoExpr codegenStateRef modNameStr (GoVar (param <> "_box")) TypeValue ty) params)
    boxed = boxGoExpr codegenStateRef modNameStr call resultType
    arity = Array.length params
    wrapperName = if arity == 1 then "gopurs_runtime.Func" else "gopurs_runtime.Func" <> show arity
  in
    if arity <= 10 then
      GoCall (GoVar wrapperName) [ GoFuncLit (map (\param -> Tuple param TypeValue) wrapperParams) [] boxed TypeValue ]
    else Array.foldr
      (\param inner -> GoCall (GoSelector (GoVar "gopurs_runtime") "Func")
        [ GoFuncLit [ Tuple param TypeValue ] [] inner TypeValue ])
      boxed wrapperParams
