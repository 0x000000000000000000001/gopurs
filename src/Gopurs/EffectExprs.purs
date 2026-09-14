module Gopurs.EffectExprs
  ( bindEffect
  , pureValue
  , defer
  , primitive
  , wrap
  ) where

import Prelude
import Data.Map as Map
import Data.Maybe (Maybe(..))
import Data.Tuple (Tuple(..))
import Gopurs.ExprAnalysis (executeIfOpaque, unwrapTcoExpr)
import Gopurs.ExprContext (ExprContext, ExprResult, TranslateExpr, StmtTree(..), flattenStmts)
import Gopurs.GoAst (GoExpr(..), GoType(..))
import Gopurs.GoConversions (boxGoExpr)
import Gopurs.Printer (printGoExpr)
import PureScript.Backend.Optimizer.Codegen.Tco (TcoExpr(..))
import PureScript.Backend.Optimizer.CoreFn (Ident)
import PureScript.Backend.Optimizer.FreeVars (localId)
import PureScript.Backend.Optimizer.Syntax (BackendEffect(..), BackendSyntax(Abs, EffectDefer, Let, LetRec), Level)

bindEffect :: TranslateExpr -> ExprContext -> Int -> Maybe Ident -> Level -> TcoExpr -> TcoExpr -> ExprResult
bindEffect translate context@{ codegenStateRef, depth, modNameStr, bound, options: { isTail } } nextId mbIdent lvl binding body =
  let
    stripEffectDefer (TcoExpr a syn) = case unwrapTcoExpr (TcoExpr a syn) of
      EffectDefer inner -> let Tuple _ i = stripEffectDefer inner in Tuple true i
      Abs _ inner -> let Tuple _ i = stripEffectDefer inner in Tuple true i
      Let ident lvl_ val body_ -> let Tuple stripped i = stripEffectDefer body_ in Tuple stripped (TcoExpr a (Let ident lvl_ val i))
      LetRec lvl_ bindings_ body_ -> let Tuple stripped i = stripEffectDefer body_ in Tuple stripped (TcoExpr a (LetRec lvl_ bindings_ i))
      _ -> Tuple false (TcoExpr a syn)
    Tuple wasStripped realBinding = stripEffectDefer binding
    originalName = localId mbIdent lvl
    name = originalName <> "_" <> show nextId
    newBound = Map.insert originalName { name, goType: TypeValue } bound
    resBinding = translate (context { depth = (depth + 1), tcoIdent = Nothing, loopCtx = [], options = { isTail: false, inEffectBlock: true }, mbExpectedExprType = Nothing }) (nextId + 1) realBinding
    resBody = translate (context { depth = (depth + 1), bound = newBound, tcoIdent = Nothing, options = { isTail, inEffectBlock: true } }) resBinding.nextId body
    bindingExpr = if wasStripped then boxGoExpr codegenStateRef modNameStr resBinding.expr resBinding.exprType else executeIfOpaque realBinding (boxGoExpr codegenStateRef modNameStr resBinding.expr resBinding.exprType)
    bodyExpr = executeIfOpaque body resBody.expr
  in
    { stmts: resBinding.stmts <> StmtLeaf (GoAssign name bindingExpr) <> resBody.stmts, expr: bodyExpr, exprType: resBody.exprType, nextId: resBody.nextId }

pureValue :: TranslateExpr -> ExprContext -> Int -> TcoExpr -> ExprResult
pureValue translate context@{ depth } nextId binding =
  translate (context { depth = (depth + 1), tcoIdent = Nothing, loopCtx = [], options = { isTail: false, inEffectBlock: false }, mbExpectedExprType = Nothing }) nextId binding

defer :: TranslateExpr -> ExprContext -> Int -> TcoExpr -> ExprResult
defer translate context@{ codegenStateRef, depth, modNameStr } nextId binding =
  let
    resBinding = translate (context { depth = (depth + 1), tcoIdent = Nothing, loopCtx = [], options = { isTail: false, inEffectBlock: true }, mbExpectedExprType = Nothing }) nextId binding
    funcExpr = GoCall (GoSelector (GoVar "gopurs_runtime") "Func")
      [ GoFuncBlock [ Tuple "_" TypeValue ]
          (flattenStmts resBinding.stmts <> [ GoReturn (boxGoExpr codegenStateRef modNameStr resBinding.expr resBinding.exprType) ])
          TypeValue
      ]
  in
    { stmts: StmtEmpty, expr: funcExpr, exprType: TypeValue, nextId: resBinding.nextId }

primitive :: TranslateExpr -> ExprContext -> Int -> BackendEffect TcoExpr -> ExprResult
primitive translate context@{ codegenStateRef, depth, modNameStr } nextId eff =
  case eff of
  EffectRefNew a ->
    let
      resA = translate (context { depth = (depth + 1), tcoIdent = Nothing, loopCtx = [], options = { isTail: false, inEffectBlock: false }, mbExpectedExprType = Nothing }) nextId a
      refIdent = "__local_ref_" <> show resA.nextId
      declStmt = GoAssign refIdent (boxGoExpr codegenStateRef modNameStr resA.expr resA.exprType)
      ifaceIdent = "__local_iface_" <> show resA.nextId
      ifaceStmt = GoRaw ("var " <> ifaceIdent <> " interface{} = " <> refIdent)
    in
      { stmts: resA.stmts <> StmtLeaf declStmt <> StmtLeaf ifaceStmt
      , expr: GoRaw ("gopurs_runtime.Any(&" <> ifaceIdent <> ")")
      , exprType: TypeValue
      , nextId: resA.nextId + 1
      }
  EffectRefRead a ->
    let
      resA = translate (context { depth = (depth + 1), tcoIdent = Nothing, loopCtx = [], options = { isTail: false, inEffectBlock: false }, mbExpectedExprType = Nothing }) nextId a
    in
      { stmts: resA.stmts
      , expr: GoRaw ("(*(" <> printGoExpr resA.expr <> ".PtrVal().(*interface{}))).(gopurs_runtime.Value)")
      , exprType: TypeValue
      , nextId: resA.nextId
      }
  EffectRefWrite ref val ->
    let
      resRef = translate (context { depth = (depth + 1), tcoIdent = Nothing, loopCtx = [], options = { isTail: false, inEffectBlock: false }, mbExpectedExprType = Nothing }) nextId ref
      resVal = translate (context { depth = (depth + 1), tcoIdent = Nothing, loopCtx = [], options = { isTail: false, inEffectBlock: false }, mbExpectedExprType = Nothing }) resRef.nextId val
      writeStmt = GoRaw ("*(" <> printGoExpr resRef.expr <> ".PtrVal().(*interface{})) = " <> printGoExpr (boxGoExpr codegenStateRef modNameStr resVal.expr resVal.exprType))
    in
      { stmts: resRef.stmts <> resVal.stmts <> StmtLeaf writeStmt
      , expr: boxGoExpr codegenStateRef modNameStr resVal.expr resVal.exprType
      , exprType: TypeValue
      , nextId: resVal.nextId
      }

wrap :: TranslateExpr -> ExprContext -> Int -> TcoExpr -> ExprResult
wrap translate context@{ codegenStateRef, modNameStr } nextId tcoExpr =
  let
    res = translate (context { options = { isTail: false, inEffectBlock: true }, mbExpectedExprType = Nothing }) nextId tcoExpr
    funcExpr = GoCall (GoSelector (GoVar "gopurs_runtime") "Func")
      [ GoFuncBlock [ Tuple "_" TypeValue ]
          (flattenStmts res.stmts <> [ GoReturn (boxGoExpr codegenStateRef modNameStr res.expr res.exprType) ])
          TypeValue
      ]
  in
    { stmts: StmtEmpty, expr: funcExpr, exprType: TypeValue, nextId: res.nextId }
