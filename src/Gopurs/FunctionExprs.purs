module Gopurs.FunctionExprs
  ( abstraction
  , uncurriedAbstraction
  , effectAbstraction
  ) where

import Prelude
import Data.Array as Array
import Data.Array.NonEmpty (NonEmptyArray, toArray)
import Data.Array.NonEmpty as NonEmptyArray
import Data.Foldable (foldl)
import Data.Map as Map
import Data.Maybe (Maybe(..), fromMaybe)
import Data.String as String
import Data.Tuple (Tuple(..), fst)
import Effect.Ref as Ref
import Effect.Unsafe (unsafePerformEffect)
import Gopurs.CallAnalysis (collectCurriedAbs)
import Gopurs.ExprAnalysis (extractExprFuncType, extractFuncType, getExprType)
import Gopurs.ExprContext (ExprContext, ExprResult, TranslateExpr, StmtTree(..), flattenStmts)
import Gopurs.GoAst (GoExpr(..), GoType(..), goTypeToStr)
import Gopurs.GoConversions (boxGoExpr)
import Gopurs.Printer (printGoExpr)
import PureScript.Backend.Optimizer.Codegen.Tco (TcoExpr)
import PureScript.Backend.Optimizer.CoreFn (ExprType(..), Ident)
import PureScript.Backend.Optimizer.FreeVars (localId)
import PureScript.Backend.Optimizer.Syntax (Level)

abstraction :: TranslateExpr -> ExprContext -> Int -> TcoExpr -> NonEmptyArray (Tuple (Maybe Ident) Level) -> TcoExpr -> ExprResult
abstraction translate context@{ codegenStateRef, depth, modNameStr, bound, mbExpectedExprType, options: { isTail } } nextId tcoExpr args body =
  let
    grouped = collectCurriedAbs args body
    consumed = NonEmptyArray.length grouped.args
    mbFuncTy = case extractFuncType tcoExpr of
      Just r -> Just r
      Nothing -> case mbExpectedExprType of
        Just ty -> extractExprFuncType ty
        Nothing -> Nothing

    -- A computation may still separate us from further lambdas.
    -- Consume only the parameters actually collected, not the full
    -- flattened function type's argument list.
    mbBodyType = case mbFuncTy of
      Just { fArgs, fRet } | consumed <= Array.length fArgs ->
        Just case Array.drop consumed fArgs of
          [] -> fRet
          remaining -> Func remaining fRet
      _ -> Nothing

    paramsWithTypes = map (\(Tuple mbI lvl) -> Tuple (localId mbI lvl) TypeValue) (toArray grouped.args)

    newBound = foldl (\acc (Tuple idStr goType) -> Map.insert idStr { name: idStr, goType } acc) bound paramsWithTypes
    params = map fst paramsWithTypes
    resBody = translate (context { depth = (depth + 1), bound = newBound, tcoIdent = Nothing, loopCtx = [], options = { isTail, inEffectBlock: false }, mbExpectedExprType = mbBodyType }) nextId grouped.body

    buildFunc :: Array String -> GoExpr -> GoExpr
    buildFunc ps innerExpr =
      let
        len = Array.length ps
        bodyStr = case innerExpr of
          GoBlock _ -> printGoExpr innerExpr
          _ -> "return " <> printGoExpr innerExpr
      in
        if len == 1 then
          GoCall (GoSelector (GoVar "gopurs_runtime") "Func") [ GoRaw ("func(" <> fromMaybe "" (Array.index ps 0) <> " gopurs_runtime.Value) gopurs_runtime.Value {\n" <> bodyStr <> "\n}") ]
        else if len >= 2 && len <= 5 then
          let
            goParams = String.joinWith ", " (map (\p -> p <> " gopurs_runtime.Value") ps)
          in
            GoCall (GoSelector (GoVar "gopurs_runtime") ("Func" <> show len)) [ GoRaw ("func(" <> goParams <> ") gopurs_runtime.Value {\n" <> bodyStr <> "\n}") ]
        else
          let
            chunk = Array.take 5 ps
            rest = Array.drop 5 ps
          in
            buildFunc chunk (buildFunc rest innerExpr)

    funcExpr = buildFunc params (GoBlock (flattenStmts resBody.stmts <> [ GoReturn (boxGoExpr codegenStateRef modNameStr resBody.expr resBody.exprType) ]))
  in
    { stmts: StmtEmpty, expr: funcExpr, exprType: TypeValue, nextId: resBody.nextId }

uncurriedAbstraction :: TranslateExpr -> ExprContext -> Int -> TcoExpr -> Array (Tuple (Maybe Ident) Level) -> TcoExpr -> ExprResult
uncurriedAbstraction translate context@{ codegenStateRef, depth, modNameStr, bound, tcoIdent, mbExpectedExprType, options: { isTail } } nextId tcoExpr args body =
  let
    mbFuncTy = case extractExprFuncType (getExprType tcoExpr) of
      Just r -> Just r
      Nothing -> case mbExpectedExprType of
        Just ty -> extractExprFuncType ty
        Nothing -> Nothing

    paramsWithTypes = map (\(Tuple mbI lvl) -> Tuple (localId mbI lvl) TypeValue) args

    newBound = foldl (\acc (Tuple idStr goType) -> Map.insert idStr { name: idStr, goType } acc) bound paramsWithTypes

    goParams = String.joinWith ", " (map (\(Tuple p goT) -> p <> " " <> goTypeToStr goT) paramsWithTypes)
    resBody = translate (context { depth = (depth + 1), bound = newBound, tcoIdent = Nothing, loopCtx = [], options = { isTail, inEffectBlock: false }, mbExpectedExprType = ( case mbFuncTy of
          Just { fRet } -> Just fRet
          Nothing -> Nothing
      ) }) nextId body
    arity = Array.length args
  in
    if arity >= 2 && arity <= 10 then
      case tcoIdent of
        Just topName ->
          let
            callFuncDecl = "func Call_" <> modNameStr <> "_" <> topName <> "(" <> goParams <> ") gopurs_runtime.Value {\n" <> printGoExpr (GoBlock (flattenStmts resBody.stmts <> [ GoReturn (boxGoExpr codegenStateRef modNameStr resBody.expr resBody.exprType) ])) <> "\n}"
            funcExpr = unsafePerformEffect do
              Ref.modify_ (\r -> r { rawDecls = Array.snoc r.rawDecls callFuncDecl }) codegenStateRef
              pure $ GoRaw ("gopurs_runtime.Func" <> show arity <> "(Call_" <> modNameStr <> "_" <> topName <> ")")
          in
            { stmts: StmtEmpty, expr: funcExpr, exprType: TypeValue, nextId: resBody.nextId }
        Nothing ->
          let
            funcExpr = GoRaw ("gopurs_runtime.Func" <> show arity <> "(func(" <> goParams <> ") gopurs_runtime.Value {\n" <> printGoExpr (GoBlock (flattenStmts resBody.stmts <> [ GoReturn (boxGoExpr codegenStateRef modNameStr resBody.expr resBody.exprType) ])) <> "\n})")
          in
            { stmts: StmtEmpty, expr: funcExpr, exprType: TypeValue, nextId: resBody.nextId }
    else
      let
        params = map fst paramsWithTypes
        makeCurried [] = GoFunc "_" TypeValue TypeValue (GoBlock (flattenStmts resBody.stmts <> [ GoReturn (boxGoExpr codegenStateRef modNameStr resBody.expr resBody.exprType) ]))
        makeCurried [ p ] = GoFunc p TypeValue TypeValue (GoBlock (flattenStmts resBody.stmts <> [ GoReturn (boxGoExpr codegenStateRef modNameStr resBody.expr resBody.exprType) ]))
        makeCurried ps = case Array.uncons ps of
          Just { head: p, tail: rest } -> GoFunc p TypeValue TypeValue (makeCurried rest)
          Nothing -> resBody.expr
      in
        { stmts: StmtEmpty, expr: makeCurried params, exprType: TypeValue, nextId: resBody.nextId }

effectAbstraction :: TranslateExpr -> ExprContext -> Int -> TcoExpr -> Array (Tuple (Maybe Ident) Level) -> TcoExpr -> ExprResult
effectAbstraction translate context@{ codegenStateRef, depth, modNameStr, bound, mbExpectedExprType, options: { isTail } } nextId tcoExpr args body =
  let
    mbFuncTy = case extractExprFuncType (getExprType tcoExpr) of
      Just r -> Just r
      Nothing -> case mbExpectedExprType of
        Just ty -> extractExprFuncType ty
        Nothing -> Nothing

    paramsWithTypes = map (\(Tuple mbI lvl) -> Tuple (localId mbI lvl) TypeValue) args
    newBound = foldl (\acc (Tuple idStr goType) -> Map.insert idStr { name: idStr, goType } acc) bound paramsWithTypes
    goParams = String.joinWith ", " (map (\(Tuple p goT) -> p <> " " <> goTypeToStr goT) paramsWithTypes)
    resBody = translate (context { depth = (depth + 1), bound = newBound, tcoIdent = Nothing, loopCtx = [], options = { isTail, inEffectBlock: false }, mbExpectedExprType = ( case mbFuncTy of
          Just { fRet } -> Just fRet
          Nothing -> Nothing
      ) }) nextId body
    arity = Array.length args
  in
    if arity >= 2 && arity <= 5 then
      let
        funcExpr = GoRaw ("gopurs_runtime.Func" <> show arity <> "(func(" <> goParams <> ") gopurs_runtime.Value {\n" <> printGoExpr (GoBlock (flattenStmts resBody.stmts <> [ GoReturn (GoRaw ("gopurs_runtime.Apply(" <> printGoExpr (boxGoExpr codegenStateRef modNameStr resBody.expr resBody.exprType) <> ", gopurs_runtime.Value{})")) ])) <> "\n})")
      in
        { stmts: StmtEmpty, expr: funcExpr, exprType: TypeValue, nextId: resBody.nextId }
    else
      let
        params = map fst paramsWithTypes
        makeCurried [] = GoFunc "_" TypeValue TypeValue (GoBlock (flattenStmts resBody.stmts <> [ GoReturn (GoRaw ("gopurs_runtime.Apply(" <> printGoExpr (boxGoExpr codegenStateRef modNameStr resBody.expr resBody.exprType) <> ", gopurs_runtime.Value{})")) ]))
        makeCurried [ p ] = GoFunc p TypeValue TypeValue (GoBlock (flattenStmts resBody.stmts <> [ GoReturn (GoRaw ("gopurs_runtime.Apply(" <> printGoExpr (boxGoExpr codegenStateRef modNameStr resBody.expr resBody.exprType) <> ", gopurs_runtime.Value{})")) ]))
        makeCurried ps = case Array.uncons ps of
          Just { head: p, tail: rest } -> GoFunc p TypeValue TypeValue (makeCurried rest)
          Nothing -> resBody.expr
      in
        { stmts: StmtEmpty, expr: makeCurried params, exprType: TypeValue, nextId: resBody.nextId }
