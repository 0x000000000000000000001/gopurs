module Gopurs.ArrayIntrinsics.Callback (nativeLambda, worker) where

import Prelude

import Control.Alternative (guard)
import Data.Array as Array
import Data.Array.NonEmpty as NEA
import Data.Map as Map
import Data.Maybe (Maybe(..))
import Data.Newtype (unwrap)
import Data.String as String
import Data.String.Pattern (Pattern(..), Replacement(..))
import Data.Tuple (Tuple(..))
import Gopurs.CallAnalysis (collectCurriedAbs, qualifiedTarget)
import Gopurs.CodegenState (FunctionInfo)
import Gopurs.ExprAnalysis (extractFuncType, unwrapTcoExpr)
import Gopurs.ExprContext (ExprContext, ExprResult, StmtTree(..), TranslateExpr, bindParameters, flattenStmts)
import Gopurs.GoAst (GoExpr(..), GoType(..))
import Gopurs.GoTypes (exprTypeToGoType)
import PureScript.Backend.Optimizer.Codegen.Tco (TcoExpr)
import PureScript.Backend.Optimizer.CoreFn (ModuleName(..))
import PureScript.Backend.Optimizer.FreeVars (localId)
import PureScript.Backend.Optimizer.Syntax (BackendSyntax(Abs))

-- Only adjacent lambdas are fused. A let/call producing the next lambda keeps
-- its original application stage through the ordinary boxed callback path.
nativeLambda :: TranslateExpr -> ExprContext -> Int -> Int -> TcoExpr -> Maybe ExprResult
nativeLambda translate context@{ metadata, modNameStr } nextId arity expression = case unwrapTcoExpr expression of
  Abs args body -> do
    let grouped = collectCurriedAbs args body
    signature <- extractFuncType expression
    guard (NEA.length grouped.args == arity && Array.length signature.fArgs == arity)
    let
      types = map (exprTypeToGoType metadata.pointerAdtPaths metadata.enumAdts metadata.elidedCtors modNameStr) signature.fArgs
      names = map (\(Tuple ident level) -> localId ident level) (NEA.toArray grouped.args)
      params = Array.zip names types
      result = translate
        (context { depth = context.depth + 1, bound = bindParameters params context.bound
          , tcoIdent = Nothing, loopCtx = [], options = { isTail: true, inEffectBlock: false }
          , mbExpectedExprType = Just signature.fRet, mbExpectedGoType = Nothing }) nextId grouped.body
    pure { stmts: StmtEmpty, expr: GoFuncLit params (flattenStmts result.stmts) result.expr result.exprType
      , exprType: TypeFunc types result.exprType, nextId: result.nextId }
  _ -> Nothing

-- Imported workers live in globalFunctions; moduleFunctions contains only
-- the current module. Preserve the worker's symbol reference for package output.
worker :: ExprContext -> Int -> TcoExpr -> Maybe FunctionInfo
worker { metadata, moduleFunctions, modNameStr } arity expression = do
  target <- qualifiedTarget expression
  let local = target.mbMod == Nothing ||
        map (String.replaceAll (Pattern ".") (Replacement "_") <<< unwrap) target.mbMod == Just modNameStr
  info <- if local then Map.lookup target.name moduleFunctions
    else do
      ModuleName owner <- target.mbMod
      Map.lookup (owner <> "." <> target.name) metadata.globalFunctions
  guard (info.arity == arity)
  pure info
