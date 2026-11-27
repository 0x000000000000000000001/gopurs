module Gopurs.ModuleWorkers (declaration) where

import Prelude
import Data.Array as Array
import Data.Foldable (any)
import Data.Map as Map
import Data.Maybe (Maybe(..))
import Data.Tuple (Tuple(..), fst, snd)
import Effect.Ref as Ref
import Effect.Unsafe (unsafePerformEffect)
import Gopurs.ExprAnalysis (extractExprFuncType, functionResultAfter, getExprType)
import Gopurs.ExprContext (ExprContext, TranslateExpr, bindParameters, flattenStmts)
import Gopurs.ArrayIndexedFold as ArrayIndexedFold
import Gopurs.CallAnalysis (collectGoSpine, getGoSpineArgs)
import Gopurs.GoAst (GoDecl(..), GoExpr(..), GoType(..))
import Gopurs.GoSymbol (GoReference(..), generated, referenceName)
import Gopurs.GoConversions (boxGoExpr, coerceGoExpr)
import Gopurs.GoFunctions (Parameters, curriedFunction, iterationBindings, loopParameters, namedParameters)
import Gopurs.Int32Loops as Int32Loops
import Gopurs.AffineLoops as AffineLoops
import Gopurs.RecordAffineLoops as RecordAffineLoops
import Gopurs.ArraySnoc as ArraySnoc
import Gopurs.ListFilter as ListFilter
import Gopurs.TailAddLoops as TailAddLoops
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
    workerContext = context
      { depth = 0
      , bound = bindParameters params Map.empty
      , tcoIdent = Just fn.ident
      , loopCtx = loop
      , options = { isTail: isSelfRecursive, inEffectBlock: false }
      , mbExpectedExprType = expectedBodyType
      , mbExpectedGoType = if Array.null fn.args || resultType == TypeValue then Nothing else Just resultType
      }
    resultType = case info of
      Just signature | Array.length fn.args < Array.length signature.fArgs -> TypeValue
      Just signature -> signature.fRet
      Nothing -> TypeValue
    tailAdd = if isSelfRecursive then
      TailAddLoops.worker translate workerContext fn.ident params resultType fn.body
      else Nothing
    body boxed = case tailAdd of
      Just loopBody -> loopBody
      Nothing ->
        let
          native = case ListFilter.worker context.metadata params resultType fn.val of
            Just filtered -> Just filtered
            Nothing -> ArraySnoc.worker modNameStr fn.ident params resultType fn.body
          result = case native of
            Just lowered -> lowered
            Nothing -> translate workerContext 0 fn.body
          returning = if boxed then boxGoExpr codegenStateRef modNameStr result.expr result.exprType
            else coerceGoExpr codegenStateRef modNameStr result.expr result.exprType resultType
          statements = flattenStmts result.stmts <> [ GoReturn returning ]
        in if isSelfRecursive then
          AffineLoops.wrap translate workerContext (AffineLoops.globalTarget modNameStr fn.ident) params resultType fn.body
            (Int32Loops.loop fn.ident params resultType statements)
          else GoBlock (iterationBindings params <> statements)
    value = if Array.null fn.args then
      -- A zero-argument abstraction still defers its whole body until called.
      curriedFunction [ Tuple "_" TypeValue ] TypeValue
        (body true)
      else
        let
          workerName = "Call_" <> modNameStr <> "_" <> fn.ident
          coldIdent = "__gopurs_affine_cold_" <> fn.ident
          coldName = "Call_" <> modNameStr <> "_" <> coldIdent
          reserved = Map.member coldIdent moduleFunctions
            || any (\signature -> referenceName signature.reference == coldName) context.metadata.ffiFunctions
          coldCall = GoCall (GoRef (SymbolReference (generated modNameStr coldName)))
            (map (GoVar <<< fst) (loopParameters params))
          hot = if isSelfRecursive && not reserved then
            RecordAffineLoops.hot (AffineLoops.globalTarget modNameStr fn.ident) params resultType fn.body coldCall
            else Nothing
          original = body false
          declare name workerBody = GoFunctionDecl
            { name: generated modNameStr name, params: loopParameters params, result: resultType, body: workerBody }
          workers = case hot of
            Just fast -> [ declare coldName original, declare workerName fast ]
            Nothing -> [ declare workerName original ]
        in
          unsafePerformEffect do
            Ref.modify_ (\state -> state { declarations = state.declarations <> workers }) codegenStateRef
            pure (case indexedWrapper translate context params fn.body of
              Just wrapper -> wrapper
              Nothing -> boxedWrapper context fn.ident params resultType)
  in
    GoCachedValue { identifier: generated modNameStr (modNameStr <> "_" <> fn.ident), expression: value, goType: TypeValue }

-- A scalar indexed fold can read either representation directly. Keep the
-- public Value wrapper on its original Value buffer instead of first copying
-- it into the native worker's array. Admission is shared with the fold emitter.
indexedWrapper :: TranslateExpr -> ExprContext -> Parameters -> TcoExpr -> Maybe GoExpr
indexedWrapper translate context params body = case params of
  [ Tuple name (TypeNativeArray TypeInt64) ] -> do
    let
      boxedContext = context { bound = bindParameters [ Tuple name TypeValue ] Map.empty }
      Tuple function spine = collectGoSpine body
    result <- ArrayIndexedFold.emit translate boxedContext 0 function (getGoSpineArgs spine)
    pure (GoCall (GoVar "gopurs_runtime.Func")
      [ GoFuncLit [ Tuple name TypeValue ] (flattenStmts result.stmts)
          (boxGoExpr context.codegenStateRef context.modNameStr result.expr result.exprType) TypeValue ])
  _ -> Nothing

-- Module wrappers group up to ten arguments; larger arities keep unary layers.
-- LocalWorkers and anonymous closures retain their own grouping policies.
boxedWrapper :: ExprContext -> String -> Parameters -> GoType -> GoExpr
boxedWrapper { codegenStateRef, modNameStr } name params resultType =
  let
    wrapperParams = map (\(Tuple param _) -> param <> "_box") params
    call = GoCall (GoRef (SymbolReference (generated modNameStr ("Call_" <> modNameStr <> "_" <> name))))
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
