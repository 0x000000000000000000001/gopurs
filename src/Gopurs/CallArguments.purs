module Gopurs.CallArguments
  ( Arguments
  , childContext
  , translate
  , translateBoxed
  , coercePrefix
  , boxRemaining
  , applyBoxed
  ) where

import Prelude
import Data.Array as Array
import Data.Foldable (foldl)
import Data.Maybe (Maybe(..), fromMaybe)
import Gopurs.ExprContext (ExprContext, ExprResult, TranslateExpr, StmtTree)
import Gopurs.GoAst (GoExpr(..), GoType(..))
import Gopurs.GoConversions (boxGoExpr, coerceGoExpr)
import PureScript.Backend.Optimizer.Codegen.Tco (TcoExpr)
import PureScript.Backend.Optimizer.CoreFn (ExprType)

type Arguments =
  { stmts :: StmtTree
  , exprs :: Array GoExpr
  , exprTypes :: Array GoType
  , nextId :: Int
  }

childContext :: ExprContext -> Maybe ExprType -> ExprContext
childContext context expected = context
  { depth = context.depth + 1
  , tcoIdent = Nothing
  , loopCtx = []
  , options = { isTail: false, inEffectBlock: false }
  , mbExpectedExprType = expected
  }

translate :: TranslateExpr -> ExprContext -> Maybe ExprType -> { stmts :: StmtTree, nextId :: Int } -> Array TcoExpr -> Arguments
translate translateExpr context expected = translateWith translateExpr (childContext context expected) _.expr

-- Curried intrinsics box each argument immediately, before translating the
-- next one. Keep its source type for the existing surplus-argument adaptation.
translateBoxed :: TranslateExpr -> ExprContext -> { stmts :: StmtTree, nextId :: Int } -> Array TcoExpr -> Arguments
translateBoxed translateExpr context@{ codegenStateRef, modNameStr } =
  translateWith translateExpr (childContext context Nothing)
    (\result -> boxGoExpr codegenStateRef modNameStr result.expr result.exprType)

translateWith :: TranslateExpr -> ExprContext -> (ExprResult -> GoExpr) -> { stmts :: StmtTree, nextId :: Int } -> Array TcoExpr -> Arguments
translateWith translateExpr context adapt initial =
  foldl
    (\acc arg ->
      let result = translateExpr context acc.nextId arg
      in
        { stmts: acc.stmts <> result.stmts
        , exprs: Array.snoc acc.exprs (adapt result)
        , exprTypes: Array.snoc acc.exprTypes result.exprType
        , nextId: result.nextId
        })
    { stmts: initial.stmts, exprs: [], exprTypes: [], nextId: initial.nextId }

coercePrefix :: ExprContext -> Int -> Array GoType -> Arguments -> Array GoExpr
coercePrefix { codegenStateRef, modNameStr } arity expected args =
  Array.mapWithIndex
    (\index expr -> coerceGoExpr codegenStateRef modNameStr expr
      (fromMaybe TypeValue (Array.index args.exprTypes index))
      (fromMaybe TypeValue (Array.index expected index)))
    (Array.take arity args.exprs)

boxRemaining :: ExprContext -> Int -> Arguments -> Array GoExpr
boxRemaining { codegenStateRef, modNameStr } arity args =
  Array.zipWith (boxGoExpr codegenStateRef modNameStr)
    (Array.drop arity args.exprs) (Array.drop arity args.exprTypes)

applyBoxed :: GoExpr -> Array GoExpr -> GoExpr
applyBoxed fExpr argExprs =
  let len = Array.length argExprs
  in
    if len == 0 then fExpr
    else if len == 1 then GoCall (GoSelector (GoVar "gopurs_runtime") "Apply") [ fExpr, fromMaybe (GoRaw "nil") (Array.index argExprs 0) ]
    else if len <= 10 then GoCall (GoSelector (GoVar "gopurs_runtime") ("Apply" <> show len)) (Array.cons fExpr argExprs)
    else applyBoxed (applyBoxed fExpr (Array.take 10 argExprs)) (Array.drop 10 argExprs)
