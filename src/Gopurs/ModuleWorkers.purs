module Gopurs.ModuleWorkers (declaration) where

import Prelude
import Data.Array as Array
import Data.Foldable (any)
import Data.Map as Map
import Data.Maybe (Maybe(..))
import Data.String as String
import Data.String.Pattern (Pattern(..))
import Data.Tuple (Tuple(..), fst, snd)
import Effect.Ref as Ref
import Effect.Unsafe (unsafePerformEffect)
import Gopurs.ExprAnalysis (extractExprFuncType, functionResultAfter, getExprType)
import Gopurs.ExprContext (ExprContext, TranslateExpr, bindParameters, flattenStmts)
import Gopurs.ArrayIndexedFold as ArrayIndexedFold
import Gopurs.CallAnalysis (collectGoSpine, getGoSpineArgs)
import Gopurs.GoAst (GoDecl(..), GoExpr(..), GoType(..), isPureConstantExpr, rawGo)
import Gopurs.GoSymbol (GoReference(..), generated, referenceName)
import Gopurs.GoConversions (boxGoExpr, coerceGoExpr)
import Gopurs.GoFunctions (Parameters, curriedFunction, loopParameters, namedParameters)
import Gopurs.Printer (printGoExpr)
import Gopurs.Int32Loops as Int32Loops
import Gopurs.AffineLoops as AffineLoops
import Gopurs.RecordAffineLoops as RecordAffineLoops
import Gopurs.ArraySnoc as ArraySnoc
import Gopurs.FunctionFusion as FunctionFusion
import Gopurs.ListFilter as ListFilter
import Gopurs.Int32Recursion as Int32Recursion
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
          workerName = "Call_" <> modNameStr <> "_" <> fn.ident
          guardedStatements = if isSelfRecursive && not boxed then
            guardNullableSelfCalls workerName params statements
            else statements
        in if isSelfRecursive then
          let
            looped = AffineLoops.wrap translate workerContext (AffineLoops.globalTarget modNameStr fn.ident) params resultType fn.body
              (Int32Loops.loop fn.ident params resultType guardedStatements)
          in case if boxed then Nothing else FunctionFusion.iteratorWorkerGuard fn.ident params resultType fn.body of
            Just guardStmt -> GoBlock [ guardStmt, looped ]
            Nothing -> looped
          else GoBlock (flattenTailReturn statements)
    value = if Array.null fn.args then
      -- A zero-argument abstraction still defers its whole body until called.
      curriedFunction [ Tuple "_" TypeValue ] TypeValue
        (body true)
      else
        let
          workerName = "Call_" <> modNameStr <> "_" <> fn.ident
          target = AffineLoops.globalTarget modNameStr fn.ident
          isReserved ident name = Map.member ident moduleFunctions
            || any (\signature -> referenceName signature.reference == name) context.metadata.ffiFunctions
          coldArgs = map (GoVar <<< fst) (loopParameters params)
          makeColdCall name = GoCall (GoRef (SymbolReference (generated modNameStr name))) coldArgs
          recordColdIdent = "__gopurs_affine_cold_" <> fn.ident
          recordColdName = "Call_" <> modNameStr <> "_" <> recordColdIdent
          scalarColdIdent = "__gopurs_cold_affine_" <> fn.ident
          scalarColdName = "Call_" <> modNameStr <> "_" <> scalarColdIdent
          i32RecColdIdent = "__gopurs_cold_i32_" <> fn.ident
          i32RecColdName = "Call_" <> modNameStr <> "_" <> i32RecColdIdent
          i32RecWorkerIdent = "__gopurs_i32_" <> fn.ident
          i32RecWorkerName = "Call_" <> modNameStr <> "_" <> i32RecWorkerIdent
          hot = if isSelfRecursive then
            case (if isReserved recordColdIdent recordColdName then Nothing
                  else RecordAffineLoops.hot target params resultType fn.body (makeColdCall recordColdName)) of
              Just fast -> Just { coldName: recordColdName, fast, extra: [] }
              Nothing -> case (if isReserved scalarColdIdent scalarColdName then Nothing
                else map (\fast -> { coldName: scalarColdName, fast, extra: [] })
                  (AffineLoops.hot translate workerContext target params resultType fn.body (makeColdCall scalarColdName))) of
                Just found -> Just found
                Nothing -> if isReserved i32RecColdIdent i32RecColdName || isReserved i32RecWorkerIdent i32RecWorkerName then Nothing
                  else map (\p -> { coldName: i32RecColdName, fast: p.fast, extra: [ p.worker ] })
                    (Int32Recursion.plan modNameStr fn.ident params resultType fn.body (generated modNameStr i32RecWorkerName) (makeColdCall i32RecColdName))
            else Nothing
          original = body false
          workerParams = if isSelfRecursive then loopParameters params else params
          declare name workerBody = GoFunctionDecl
            { name: generated modNameStr name, params: workerParams, result: resultType, body: workerBody }
          workers = case hot of
            Just split -> split.extra <> [ declare split.coldName original, declare workerName split.fast ]
            Nothing -> [ declare workerName original ]
        in
          unsafePerformEffect do
            Ref.modify_ (\state -> state { declarations = state.declarations <> workers }) codegenStateRef
            pure (case indexedWrapper translate context params fn.body of
              Just wrapper -> wrapper
              Nothing -> boxedWrapper context fn.ident params resultType)
  in
    GoCachedValue { identifier: generated modNameStr (modNameStr <> "_" <> fn.ident), expression: value, goType: TypeValue }

-- When a 1-argument self-recursive worker over a nullable pointer ADT returns
-- a pure scalar constant immediately on `param == nil`, guard direct self-calls
-- on pure pointer reads (`v.Fi` or local vars) with a local nil check so empty
-- leaf children skip function-call and stack-check overhead entirely.
guardNullableSelfCalls :: String -> Parameters -> Array GoExpr -> Array GoExpr
guardNullableSelfCalls workerName params stmts = case params of
  [ Tuple paramName (TypeStructPointer _) ] ->
    case findNullConstReturn paramName stmts of
      Just constVal -> rewriteStmts constVal stmts
      Nothing -> stmts
  _ -> stmts
  where
  findNullConstReturn paramName xs = do
    decl <- Array.index xs 0
    open <- Array.index xs 1
    firstIf <- Array.index xs 2
    endLabel <- Array.index xs (Array.length xs - 2)
    ret <- Array.last xs
    tmpVar <- case ret of
      GoReturn (GoVar v) -> Just v
      _ -> Nothing
    label <- case endLabel of
      GoRaw { text } -> String.stripSuffix (Pattern ":") text
      _ -> Nothing
    case decl, open, firstIf of
      GoRaw { text: declText }, GoRaw { text: "{" }, GoIfElse cond [ GoMutate lhs constVal, GoRaw { text: gotoText } ] []
        | String.contains (Pattern ("var " <> tmpVar <> " ")) declText
        , lhs == tmpVar
        , gotoText == "goto " <> label
        , isNullParamCond paramName cond
        , isPureConstantExpr constVal ->
            Just constVal
      _, _, _ -> Nothing

  isNullParamCond paramName = case _ of
    GoRaw { text } -> text == "(" <> paramName <> " == nil)"
    GoBinOp "==" (GoVar v) (GoRaw { text: "nil" }) -> v == paramName
    _ -> false

  isPureCheapPtrArg = case _ of
    GoVar _ -> true
    GoStructAccess (GoVar _) _ -> true
    GoConstructorAccess (GoVar _) _ _ _ true -> true
    _ -> false

  rewriteStmts constVal = Array.concatMap (rewriteStmt constVal)

  rewriteStmt constVal = case _ of
    GoAssign lhs (GoCall (GoRef (SymbolReference sym)) [ arg ])
      | sym.name == workerName && isPureCheapPtrArg arg ->
          let tmp = "__nc_" <> lhs
          in
            [ GoAssign lhs constVal
            , GoIfElse (rawGo (tmp <> " := " <> printGoExpr arg <> "; " <> tmp <> " != nil"))
                [ GoMutate lhs (GoCall (GoRef (SymbolReference sym)) [ GoVar tmp ]) ] []
            ]
    GoIfElse cond t f -> [ GoIfElse cond (rewriteStmts constVal t) (rewriteStmts constVal f) ]
    GoBlock inner -> [ GoBlock (rewriteStmts constVal inner) ]
    other -> [ other ]

-- Unwrap a trailing zero-argument IIFE `return func() T { stmts }()` in a
-- non-recursive worker when its statements introduce no local bindings or loop
-- labels. Go's inliner penalizes functions containing closure literals (+57
-- cost), so unwrapping restores direct inlining of tiny guarded fast paths.
flattenTailReturn :: Array GoExpr -> Array GoExpr
flattenTailReturn stmts = case Array.unsnoc stmts of
  Just { init, last: GoReturn (GoCall (GoFuncBlock [] innerStmts _) []) }
    | Array.all isDeclFreeStmt innerStmts -> init <> innerStmts
  Just { init, last: GoReturn (GoCall (GoFuncLit [] innerStmts innerRet _) []) }
    | Array.all isDeclFreeStmt innerStmts -> init <> innerStmts <> [ GoReturn innerRet ]
  _ -> stmts
  where
  isDeclFreeStmt = case _ of
    GoReturn _ -> true
    GoMutate _ _ -> true
    GoRaw { text } -> not (String.contains (Pattern ":=") text || String.contains (Pattern "var ") text)
    GoIfElse _ t f -> Array.all isDeclFreeStmt t && Array.all isDeclFreeStmt f
    GoBlock inner -> Array.all isDeclFreeStmt inner
    _ -> false

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
