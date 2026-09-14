module Gopurs.ControlExprs
  ( failure
  , branch
  , booleanAnd
  , booleanOr
  ) where

import Prelude
import Data.Array as Array
import Data.Array.NonEmpty (NonEmptyArray, toArray)
import Data.Foldable (foldl)
import Data.Maybe (Maybe(..), fromMaybe)
import Gopurs.ExprAnalysis (getExprType, unwrapTcoExpr)
import Gopurs.ExprContext (ExprContext, ExprResult, TranslateExpr, StmtTree(..), flattenStmts)
import Gopurs.GoAst (rawGo, GoExpr(..), GoType(..), goTypeToStr)
import Gopurs.GoConversions (coerceGoExpr, unboxGoExpr)
import Gopurs.GoTypes (exprTypeToGoType)
import Gopurs.Printer (printGoExpr)
import PureScript.Backend.Optimizer.Codegen.Tco (TcoExpr)
import PureScript.Backend.Optimizer.CoreFn (ExprType(..))
import PureScript.Backend.Optimizer.Syntax (BackendSyntax(Fail), Pair(..))

failure :: ExprContext -> Int -> TcoExpr -> String -> ExprResult
failure { metadata, modNameStr, mbExpectedExprType } nextId tcoExpr msg =
  let
    expectedGoType = exprTypeToGoType metadata.pointerAdtPaths metadata.enumAdts metadata.elidedCtors modNameStr
      ( case getExprType tcoExpr of
          Any -> fromMaybe Any mbExpectedExprType
          ty -> ty
      )
    expectedGoTypeStr = goTypeToStr expectedGoType
  in
    { stmts: StmtEmpty, expr: rawGo ("func() " <> expectedGoTypeStr <> " { panic(" <> printGoExpr (GoString msg) <> ") }()"), exprType: expectedGoType, nextId }

branch :: TranslateExpr -> ExprContext -> Int -> NonEmptyArray (Pair TcoExpr) -> TcoExpr -> ExprResult
branch translate context@{ codegenStateRef, depth, modNameStr, options: { isTail } } nextId branches def =
  let
    resDef = translate (context { depth = (depth + 1), tcoIdent = Nothing, options = { isTail, inEffectBlock: false } }) nextId def

    computedBranches = foldl
      ( \acc (Pair condExpr bodyExpr) ->
          let
            resCond = translate (context { depth = (depth + 1), tcoIdent = Nothing, loopCtx = [], options = { isTail: false, inEffectBlock: false }, mbExpectedExprType = Nothing }) acc.nextId condExpr
            resBody = translate (context { depth = (depth + 1), tcoIdent = Nothing, options = { isTail, inEffectBlock: false } }) resCond.nextId bodyExpr
          in
            { nextId: resBody.nextId, results: acc.results <> [ { cond: resCond, body: resBody } ] }
      )
      { nextId: resDef.nextId, results: [] }
      (toArray branches)

    isFailNode = case unwrapTcoExpr def of
      Fail _ -> true
      _ -> false

    allTypes = (if isFailNode then [] else [ resDef.exprType ]) <> map (\r -> r.body.exprType) computedBranches.results

    hasTypeValue = Array.any
      ( \t -> case t of
          TypeValue -> true
          _ -> false
      )
      allTypes

    expectedGoType =
      if hasTypeValue then TypeValue
      else
        let
          nubbed = Array.nub (map goTypeToStr allTypes)
        in
          if Array.length nubbed == 1 then fromMaybe TypeValue (Array.head allTypes)
          else TypeValue

    tmpVar = "__t" <> show computedBranches.nextId
    declTmp = StmtLeaf (rawGo ("var " <> tmpVar <> " " <> goTypeToStr expectedGoType))
    labelName = "end_branch_" <> show computedBranches.nextId

    buildIfs = foldl
      ( \acc r ->
          let
            goIf = GoIfElse (unboxGoExpr codegenStateRef modNameStr r.cond.expr r.cond.exprType TypeBool) (flattenStmts r.body.stmts <> [ GoMutate tmpVar (coerceGoExpr codegenStateRef modNameStr r.body.expr r.body.exprType expectedGoType), rawGo ("goto " <> labelName) ]) []
          in
            acc <> StmtLeaf (rawGo "{") <> r.cond.stmts <> StmtLeaf goIf <> StmtLeaf (rawGo "}")
      )
      StmtEmpty
      computedBranches.results
  in
    { stmts: declTmp <> buildIfs <> StmtLeaf (rawGo "{") <> resDef.stmts <> StmtLeaf (GoMutate tmpVar (coerceGoExpr codegenStateRef modNameStr resDef.expr resDef.exprType expectedGoType)) <> StmtLeaf (rawGo "}") <> StmtLeaf (rawGo (labelName <> ":")), expr: GoVar tmpVar, exprType: expectedGoType, nextId: computedBranches.nextId + 1 }

booleanAnd :: TranslateExpr -> ExprContext -> Int -> TcoExpr -> TcoExpr -> ExprResult
booleanAnd translate context@{ codegenStateRef, depth, modNameStr } nextId e1 e2 =
  let
    res1 = translate (context { depth = (depth + 1), tcoIdent = Nothing, loopCtx = [], options = { isTail: false, inEffectBlock: false }, mbExpectedExprType = Nothing }) nextId e1
    res2 = translate (context { depth = (depth + 1), tcoIdent = Nothing, loopCtx = [], options = { isTail: false, inEffectBlock: false }, mbExpectedExprType = Nothing }) res1.nextId e2
  in
    if isEmptyStmts res2.stmts then
      { expr: GoBinOp "&&" (unboxGoExpr codegenStateRef modNameStr res1.expr res1.exprType TypeBool) (unboxGoExpr codegenStateRef modNameStr res2.expr res2.exprType TypeBool), exprType: TypeBool, stmts: res1.stmts <> res2.stmts, nextId: res2.nextId }
    else
      let
        tmpVar = "__t_and_" <> show res2.nextId
        declTmp = StmtLeaf (rawGo ("var " <> tmpVar <> " bool = false\nif " <> printGoExpr (unboxGoExpr codegenStateRef modNameStr res1.expr res1.exprType TypeBool) <> " {\n"))
        assignTmp = StmtLeaf (rawGo (tmpVar <> " = " <> printGoExpr (unboxGoExpr codegenStateRef modNameStr res2.expr res2.exprType TypeBool) <> "\n}"))
      in
        { expr: rawGo tmpVar, exprType: TypeBool, stmts: res1.stmts <> declTmp <> res2.stmts <> assignTmp, nextId: res2.nextId + 1 }

booleanOr :: TranslateExpr -> ExprContext -> Int -> TcoExpr -> TcoExpr -> ExprResult
booleanOr translate context@{ codegenStateRef, depth, modNameStr } nextId e1 e2 =
  let
    res1 = translate (context { depth = (depth + 1), tcoIdent = Nothing, loopCtx = [], options = { isTail: false, inEffectBlock: false }, mbExpectedExprType = Nothing }) nextId e1
    res2 = translate (context { depth = (depth + 1), tcoIdent = Nothing, loopCtx = [], options = { isTail: false, inEffectBlock: false }, mbExpectedExprType = Nothing }) res1.nextId e2
  in
    if isEmptyStmts res2.stmts then
      { expr: GoBinOp "||" (unboxGoExpr codegenStateRef modNameStr res1.expr res1.exprType TypeBool) (unboxGoExpr codegenStateRef modNameStr res2.expr res2.exprType TypeBool), exprType: TypeBool, stmts: res1.stmts <> res2.stmts, nextId: res2.nextId }
    else
      let
        tmpVar = "__t_or_" <> show res2.nextId
        declTmp = StmtLeaf (rawGo ("var " <> tmpVar <> " bool = true\nif !(" <> printGoExpr (unboxGoExpr codegenStateRef modNameStr res1.expr res1.exprType TypeBool) <> ") {\n"))
        assignTmp = StmtLeaf (rawGo (tmpVar <> " = " <> printGoExpr (unboxGoExpr codegenStateRef modNameStr res2.expr res2.exprType TypeBool) <> "\n}"))
      in
        { expr: rawGo tmpVar, exprType: TypeBool, stmts: res1.stmts <> declTmp <> res2.stmts <> assignTmp, nextId: res2.nextId + 1 }

isEmptyStmts :: StmtTree -> Boolean
isEmptyStmts StmtEmpty = true
isEmptyStmts (StmtAppend s1 s2) = isEmptyStmts s1 && isEmptyStmts s2
isEmptyStmts _ = false
