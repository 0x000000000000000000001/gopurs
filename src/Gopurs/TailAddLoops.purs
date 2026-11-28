-- Turn recursion in the right operand of IntAdd into an accumulator loop.
-- The accumulator is private and IntAdd is total/associative modulo 2^32;
-- every observable operand and argument retains its original evaluation order.
-- Pure order-2 linear additive recurrences (f(n) = f(n-1) + f(n-2)) also
-- receive a guarded O(n) int32 iterative fast path for 2 <= n <= MaxInt32.
module Gopurs.TailAddLoops (worker, hot) where

import Prelude hiding (add)

import Control.Alternative (guard)
import Data.Array as Array
import Data.Array.NonEmpty as NEA
import Data.Foldable (foldMap, foldl, foldr)
import Data.Map as Map
import Data.Maybe (Maybe(..))
import Data.Newtype (unwrap)
import Data.String as String
import Data.String.Pattern (Pattern(..), Replacement(..))
import Data.Tuple (Tuple(..), fst, snd)
import Gopurs.CallAnalysis (collectGoSpine, getGoSpineArgs)
import Gopurs.ExprAnalysis (unwrapTcoExpr)
import Gopurs.ExprContext (ExprContext, TranslateExpr, childContext, flattenStmts)
import Gopurs.GoAst (GoExpr(..), GoType(..), rawGo)
import Gopurs.GoConversions (coerceGoExpr)
import Gopurs.GoFunctions (Parameters, iterationBindings)
import Gopurs.GoSymbol (GoReference)
import PureScript.Backend.Optimizer.Codegen.Tco (TcoExpr(..))
import PureScript.Backend.Optimizer.CoreFn (ExprType(Int), Ident(..), Literal(..), Qualified(..))
import PureScript.Backend.Optimizer.FreeVars (localId)
import PureScript.Backend.Optimizer.Syntax (BackendOperator(..), BackendOperator1(..), BackendOperator2(..), BackendOperatorNum(..), BackendOperatorOrd(..), BackendSyntax(..), Pair(..))

data Plan
  = Finish TcoExpr
  | Step TcoExpr (Array TcoExpr)
  | Choice (Array (Tuple TcoExpr Plan)) Plan

type LinearOrder2 =
  { nParam :: String
  , c0 :: TcoExpr
  , c1 :: TcoExpr
  }

worker :: TranslateExpr -> ExprContext -> String -> Parameters -> GoType -> TcoExpr -> Maybe GoExpr
worker translate context name params resultType source = do
  guard (resultType == TypeInt64 && not (Array.null params))
  let plan = recognize context.modNameStr name (Array.length params) source
  guard (hasStep plan)
  let
    prefix = freshPrefix (Array.cons name (map fst params) <> bindingNames source)
    acc = prefix <> "acc"
    used = prefix <> "used"
    label = prefix <> "loop"
    baseCases = extractBaseCases (map fst params) plan
    unfoldStep = extractPureStep context.modNameStr name (map fst params) plan
    linearFast = case extractLinearOrder2 context.modNameStr name params plan of
      Just lin -> [ emitLinearOrder2Loop translate context params prefix lin ]
      Nothing -> []
    guardInfo =
      if not (Array.null baseCases) && hasLeftSelfCall context.modNameStr name (Array.length params) plan then
        case Map.lookup name context.moduleFunctions of
          Just info | info.arity == Array.length params && info.fArgs == map snd params && info.fRet == TypeInt64 ->
            Just { owner: context.modNameStr, name, ref: info.reference, baseCases, unfoldStep }
          _ -> Nothing
      else Nothing
    emitted = emit translate context params { prefix, acc, used, label, guardInfo } 0 plan
  pure $ GoBlock
    (linearFast <>
      [ GoAssign acc (GoCall (GoVar "int64") [ GoInt 0 ])
      , GoAssign used (GoVar "false")
      , GoFor label (iterationBindings params <> emitted.statements)
      ])

-- Outlined fast wrapper for pure order-2 linear additive recurrences so the
-- O(n) iterative path stays compact and inlines into call sites while cold
-- inputs (n < 0 or wide int64) delegate to the full fallback worker.
hot :: TranslateExpr -> ExprContext -> String -> Parameters -> GoType -> TcoExpr -> GoExpr -> Maybe GoExpr
hot translate context name params resultType source coldCall = do
  guard (resultType == TypeInt64 && not (Array.null params))
  let plan = recognize context.modNameStr name (Array.length params) source
  guard (hasStep plan)
  lin <- extractLinearOrder2 context.modNameStr name params plan
  let
    prefix = freshPrefix (Array.cons name (map fst params) <> bindingNames source)
    nLoop = lin.nParam <> "_loop"
    evalCtx = bindLoopParams context params
    c0Res = translate (childContext evalCtx (Just Int)) 0 lin.c0
    c1Res = translate (childContext evalCtx (Just Int)) c0Res.nextId lin.c1
    c0Val = native TypeInt64 (coerceGoExpr context.codegenStateRef context.modNameStr c0Res.expr c0Res.exprType TypeInt64)
    c1Val = native TypeInt64 (coerceGoExpr context.codegenStateRef context.modNameStr c1Res.expr c1Res.exprType TypeInt64)
  pure $ GoBlock
    [ emitLinearOrder2Loop translate context params prefix lin
    , GoIfElse (GoBinOp "==" (GoVar nLoop) (GoCall (GoVar "int64") [ GoInt 0 ]))
        (flattenStmts c0Res.stmts <> [ GoReturn c0Val ]) []
    , GoIfElse (GoBinOp "==" (GoVar nLoop) (GoCall (GoVar "int64") [ GoInt 1 ]))
        (flattenStmts c1Res.stmts <> [ GoReturn c1Val ]) []
    , GoReturn coldCall
    ]

bindLoopParams :: ExprContext -> Parameters -> ExprContext
bindLoopParams context params =
  context { bound = foldl (\b (Tuple p ty) -> Map.insert p { name: p <> "_loop", goType: ty } b) context.bound params }

emitLinearOrder2Loop :: TranslateExpr -> ExprContext -> Parameters -> String -> LinearOrder2 -> GoExpr
emitLinearOrder2Loop translate context params prefix lin =
  let
    nLoop = lin.nParam <> "_loop"
    aName = prefix <> "lin_a"
    bName = prefix <> "lin_b"
    remName = prefix <> "lin_rem"
    evalCtx = bindLoopParams context params
    c0Res = translate (childContext evalCtx (Just Int)) 0 lin.c0
    c1Res = translate (childContext evalCtx (Just Int)) c0Res.nextId lin.c1
    c0Val = coerceGoExpr context.codegenStateRef context.modNameStr c0Res.expr c0Res.exprType TypeInt64
    c1Val = coerceGoExpr context.codegenStateRef context.modNameStr c1Res.expr c1Res.exprType TypeInt64
    cond = GoBinOp "&&"
      (GoBinOp ">=" (GoVar nLoop) (GoCall (GoVar "int64") [ GoInt 2 ]))
      (GoBinOp "<=" (GoVar nLoop) (GoCall (GoVar "int64") [ GoInt 2147483647 ]))
  in
    GoIfElse cond
      (flattenStmts c0Res.stmts <> flattenStmts c1Res.stmts <>
        [ GoAssign aName (GoCall (GoVar "int32") [ c0Val ])
        , GoAssign bName (GoCall (GoVar "int32") [ c1Val ])
        , GoForRange (remName <> " := uint32(" <> nLoop <> " - 1); " <> remName <> " != 0; " <> remName <> "--")
            [ rawGo (aName <> ", " <> bName <> " = " <> bName <> ", " <> aName <> " + " <> bName) ]
        , GoReturn (GoCall (GoVar "int64") [ GoVar bName ])
        ])
      []

extractLinearOrder2 :: String -> String -> Parameters -> Plan -> Maybe LinearOrder2
extractLinearOrder2 owner name params plan = do
  let paramNames = map fst params
  { branches, leftArgs, rightArgs } <- flattenLinearPlan owner name (Array.length params) plan
  idx <- findVaryingIndex params leftArgs rightArgs
  Tuple nParam nTy <- Array.index params idx
  guard (nTy == TypeInt64)
  { c0, c1 } <- extractOrder2Bases nParam paramNames branches
  pure { nParam, c0, c1 }

flattenLinearPlan
  :: String
  -> String
  -> Int
  -> Plan
  -> Maybe { branches :: Array BaseCase, leftArgs :: Array TcoExpr, rightArgs :: Array TcoExpr }
flattenLinearPlan owner name arity = go []
  where
  go acc (Choice branches fallback) =
    if Array.length acc + Array.length branches > 2 then Nothing
    else case collectFinish branches of
      Just found -> go (acc <> found) fallback
      Nothing -> Nothing
  go acc (Step left rightArgs) | not (Array.null acc) = do
    leftArgs <- selfCallArgs owner name arity left
    pure { branches: acc, leftArgs, rightArgs }
  go _ _ = Nothing

  collectFinish arr = case Array.uncons arr of
    Nothing -> Just []
    Just { head: Tuple cond (Finish leaf), tail } -> map (Array.cons { cond, leaf }) (collectFinish tail)
    _ -> Nothing

findVaryingIndex :: Parameters -> Array TcoExpr -> Array TcoExpr -> Maybe Int
findVaryingIndex params leftArgs rightArgs =
  let
    indices = Array.range 0 (Array.length params - 1)
    isInvariant j = case Array.index params j, Array.index leftArgs j, Array.index rightArgs j of
      Just (Tuple p _), Just l, Just r -> isLocalName p l && isLocalName p r
      _, _, _ -> false
    isOrder2Step j = case Array.index params j, Array.index leftArgs j, Array.index rightArgs j of
      Just (Tuple p TypeInt64), Just l, Just r ->
        (isSubConst p 1 l && isSubConst p 2 r) || (isSubConst p 2 l && isSubConst p 1 r)
      _, _, _ -> false
  in
    Array.find (\i -> isOrder2Step i && Array.all (\j -> j == i || isInvariant j) indices) indices

isLocalName :: String -> TcoExpr -> Boolean
isLocalName expected expr = case unwrapTcoExpr expr of
  Local mbIdent level -> localId mbIdent level == expected
  _ -> false

isSubConst :: String -> Int -> TcoExpr -> Boolean
isSubConst param k expr = case unwrapTcoExpr expr of
  PrimOp (Op2 (OpIntNum OpSubtract) left right) -> isLocalName param left && isLitInt k right
  PrimOp (Op2 (OpIntNum OpAdd) left right) ->
    (isLocalName param left && isLitInt (-k) right) || (isLitInt (-k) left && isLocalName param right)
  _ -> false

isLitInt :: Int -> TcoExpr -> Boolean
isLitInt expected expr = case unwrapTcoExpr expr of
  Lit (LitInt n) -> n == expected
  _ -> false

isEqConst :: String -> Int -> TcoExpr -> Boolean
isEqConst param k expr = case unwrapTcoExpr expr of
  PrimOp (Op2 (OpIntOrd OpEq) left right) ->
    (isLocalName param left && isLitInt k right) || (isLitInt k left && isLocalName param right)
  _ -> false

isLeOne :: String -> TcoExpr -> Boolean
isLeOne param expr = case unwrapTcoExpr expr of
  PrimOp (Op2 (OpIntOrd OpLt) left right) -> isLocalName param left && isLitInt 2 right
  PrimOp (Op2 (OpIntOrd OpGt) left right) -> isLitInt 2 left && isLocalName param right
  PrimOp (Op2 (OpIntOrd OpLte) left right) -> isLocalName param left && isLitInt 1 right
  PrimOp (Op2 (OpIntOrd OpGte) left right) -> isLitInt 1 left && isLocalName param right
  _ -> false

substituteLocalInt :: String -> Int -> TcoExpr -> TcoExpr
substituteLocalInt target val (TcoExpr analysis syntax) = case syntax of
  Local mbIdent level | localId mbIdent level == target -> TcoExpr analysis (Lit (LitInt val))
  _ -> TcoExpr analysis (map (substituteLocalInt target val) syntax)

extractOrder2Bases :: String -> Array String -> Array BaseCase -> Maybe { c0 :: TcoExpr, c1 :: TcoExpr }
extractOrder2Bases nParam paramNames = case _ of
  [ { cond, leaf } ] | isLeOne nParam cond && pureScalarExpr paramNames leaf ->
    Just { c0: substituteLocalInt nParam 0 leaf, c1: substituteLocalInt nParam 1 leaf }
  [ a, b ] | pureScalarExpr paramNames a.leaf && pureScalarExpr paramNames b.leaf ->
    if isEqConst nParam 0 a.cond && (isEqConst nParam 1 b.cond || isLeOne nParam b.cond) then
      Just { c0: substituteLocalInt nParam 0 a.leaf, c1: substituteLocalInt nParam 1 b.leaf }
    else if isEqConst nParam 1 a.cond && (isEqConst nParam 0 b.cond || isLeOne nParam b.cond) then
      Just { c0: substituteLocalInt nParam 0 b.leaf, c1: substituteLocalInt nParam 1 a.leaf }
    else Nothing
  _ -> Nothing

recognize :: String -> String -> Int -> TcoExpr -> Plan
recognize owner name arity source = case unwrapTcoExpr source of
  Branch branches fallback -> Choice
    (map (\(Pair condition value) -> Tuple condition (recognize owner name arity value)) (NEA.toArray branches))
    (recognize owner name arity fallback)
  PrimOp (Op2 (OpIntNum OpAdd) left right) ->
    let Tuple callee spine = collectGoSpine right
        args = getGoSpineArgs spine
    in case unwrapTcoExpr callee of
      Var (Qualified moduleName (Ident ident))
        | ident == name && Array.length args == arity
          && (moduleName == Nothing || map (String.replaceAll (Pattern ".") (Replacement "_") <<< unwrap) moduleName == Just owner) ->
            Step left args
      _ -> Finish source
  _ -> Finish source

hasStep :: Plan -> Boolean
hasStep = case _ of
  Step _ _ -> true
  Choice branches fallback -> Array.any (hasStep <<< \(Tuple _ body) -> body) branches || hasStep fallback
  Finish _ -> false

hasLeftSelfCall :: String -> String -> Int -> Plan -> Boolean
hasLeftSelfCall owner name arity = go
  where
  go = case _ of
    Step left _ -> case selfCallArgs owner name arity left of
      Just _ -> true
      Nothing -> false
    Choice branches fallback -> Array.any (go <<< \(Tuple _ body) -> body) branches || go fallback
    Finish _ -> false

type BaseCase = { cond :: TcoExpr, leaf :: TcoExpr }
type PureStep = { leftArgs :: Array TcoExpr, rightArgs :: Array TcoExpr }
type GuardInfo = { owner :: String, name :: String, ref :: GoReference, baseCases :: Array BaseCase, unfoldStep :: Maybe PureStep }
type Names = { prefix :: String, acc :: String, used :: String, label :: String, guardInfo :: Maybe GuardInfo }
type Emitted = { statements :: Array GoExpr, nextId :: Int }

extractPureStep :: String -> String -> Array String -> Plan -> Maybe PureStep
extractPureStep owner name paramNames = go 0
  where
  go count (Choice branches fallback) =
    if count + Array.length branches > 4 then Nothing
    else if Array.all (\(Tuple cond body) -> case body of
      Finish leaf -> pureScalarExpr paramNames cond && pureScalarExpr paramNames leaf
      _ -> false) branches
    then go (count + Array.length branches) fallback
    else Nothing
  go count (Step left rightArgs) | count > 0 && count <= 4 =
    case selfCallArgs owner name (Array.length paramNames) left of
      Just leftArgs | Array.all (pureScalarExpr paramNames) leftArgs && Array.all (pureScalarExpr paramNames) rightArgs ->
        Just { leftArgs, rightArgs }
      _ -> Nothing
  go _ _ = Nothing

extractBaseCases :: Array String -> Plan -> Array BaseCase
extractBaseCases paramNames = go
  where
  go (Choice branches fallback) =
    let taken = takePure branches
    in if taken.allBase then taken.cases <> go fallback else taken.cases
  go _ = []

  takePure arr = case Array.uncons arr of
    Nothing -> { cases: [], allBase: true }
    Just { head: Tuple cond (Finish leaf), tail } | pureScalarExpr paramNames cond && pureScalarExpr paramNames leaf ->
      let rest = takePure tail
      in { cases: Array.cons { cond, leaf } rest.cases, allBase: rest.allBase }
    _ -> { cases: [], allBase: false }

pureScalarExpr :: Array String -> TcoExpr -> Boolean
pureScalarExpr paramNames = go
  where
  go expr = case unwrapTcoExpr expr of
    Lit (LitInt _) -> true
    Lit (LitBoolean _) -> true
    Local mbIdent level -> Array.elem (localId mbIdent level) paramNames
    PrimOp (Op1 op inner) ->
      (op == OpIntNegate || op == OpIntBitNot || op == OpBooleanNot) && go inner
    PrimOp (Op2 op a b) -> totalOp op && go a && go b
    _ -> false

  totalOp = case _ of
    OpIntNum OpAdd -> true
    OpIntNum OpSubtract -> true
    OpIntNum OpMultiply -> true
    OpIntOrd _ -> true
    OpBooleanOrd _ -> true
    OpBooleanAnd -> true
    OpBooleanOr -> true
    OpIntBitAnd -> true
    OpIntBitOr -> true
    OpIntBitXor -> true
    OpIntBitShiftLeft -> true
    OpIntBitShiftRight -> true
    OpIntBitZeroFillShiftRight -> true
    _ -> false

selfCallArgs :: String -> String -> Int -> TcoExpr -> Maybe (Array TcoExpr)
selfCallArgs owner name arity expr =
  let Tuple callee spine = collectGoSpine expr
      args = getGoSpineArgs spine
  in case unwrapTcoExpr callee of
    Var (Qualified moduleName (Ident ident))
      | ident == name && Array.length args == arity
        && (moduleName == Nothing || map (String.replaceAll (Pattern ".") (Replacement "_") <<< unwrap) moduleName == Just owner) ->
          Just args
    _ -> Nothing

emit :: TranslateExpr -> ExprContext -> Parameters -> Names -> Int -> Plan -> Emitted
emit translate context@{ codegenStateRef, modNameStr } params names nextId = case _ of
  Choice branches fallback ->
    let
      choices = foldl (\current (Tuple condition body) ->
        let cond = translate (childContext context Nothing) current.nextId condition
            selected = emit translate context params names cond.nextId body
        in { statements: current.statements <> flattenStmts cond.stmts <>
              [ GoIfElse (coerceGoExpr codegenStateRef modNameStr cond.expr cond.exprType TypeBool) selected.statements [] ]
           , nextId: selected.nextId
           }) { statements: [], nextId } branches
      final = emit translate context params names choices.nextId fallback
    in { statements: choices.statements <> final.statements, nextId: final.nextId }
  Finish source ->
    let result = translate (childContext context (Just Int)) nextId source
        valueName = names.prefix <> "result_" <> show result.nextId
        value = GoVar valueName
    in
      { statements: flattenStmts result.stmts <>
          [ GoAssign valueName (native TypeInt64 (coerceGoExpr codegenStateRef modNameStr result.expr result.exprType TypeInt64))
          , GoIfElse (GoVar names.used) [ GoReturn (add (GoVar names.acc) value) ] []
          , GoReturn value
          ]
      , nextId: result.nextId + 1
      }
  Step left args ->
    let
      emitGuardedSelf evalCtx startId info argExprs buildFallback =
        let
          evalArgs = foldl (\current (Tuple (Tuple _ ty) argument) ->
            let result = translate (childContext evalCtx Nothing) current.nextId argument
                argName = names.prefix <> "larg_" <> show result.nextId
            in { statements: current.statements <> flattenStmts result.stmts <>
                   [ GoAssign argName (native ty (coerceGoExpr codegenStateRef modNameStr result.expr result.exprType ty)) ]
               , argNames: Array.snoc current.argNames argName
               , nextId: result.nextId + 1
               })
            { statements: [], argNames: [], nextId: startId } (Array.zip params argExprs)
          targetName = names.prefix <> "left_" <> show evalArgs.nextId
          rebound = foldl (\b (Tuple (Tuple p ty) argName) -> Map.insert p { name: argName, goType: ty } b)
            evalCtx.bound (Array.zip params evalArgs.argNames)
          innerCtx = evalCtx { bound = rebound }
          elseBranch = buildFallback innerCtx evalArgs.argNames targetName (evalArgs.nextId + 1)
          chain = foldr (\bc accBranch ->
            let
              cRes = translate (childContext innerCtx Nothing) accBranch.nextId bc.cond
              lRes = translate (childContext innerCtx (Just Int)) cRes.nextId bc.leaf
              condExpr = coerceGoExpr codegenStateRef modNameStr cRes.expr cRes.exprType TypeBool
              leafExpr = native TypeInt64 (coerceGoExpr codegenStateRef modNameStr lRes.expr lRes.exprType TypeInt64)
            in { stmts: flattenStmts cRes.stmts <>
                   [ GoIfElse condExpr (flattenStmts lRes.stmts <> [ GoMutate targetName leafExpr ]) accBranch.stmts ]
               , nextId: lRes.nextId
               }) elseBranch info.baseCases
        in { statements: evalArgs.statements <> [ rawGo ("var " <> targetName <> " int64") ] <> chain.stmts
           , leftName: targetName
           , nextId: chain.nextId
           }
      directCallFallback info _ argNames targetName id =
        { stmts: [ GoMutate targetName (GoCall (GoRef info.ref) (map GoVar argNames)) ], nextId: id }
      emitLeft = case names.guardInfo of
        Just info -> case selfCallArgs info.owner info.name (Array.length params) left of
          Just leftArgs ->
            -- Capture left self-call arguments in source order, check leading
            -- pure base cases, and for pure binary recurrences inline one step
            -- (with base-case guards on both subcalls) before recursing.
            let
              fallback = case info.unfoldStep of
                Just step -> \innerCtx _ targetName id ->
                  let
                    firstSub = emitGuardedSelf innerCtx id info step.leftArgs (directCallFallback info)
                    secondSub = emitGuardedSelf innerCtx firstSub.nextId info step.rightArgs (directCallFallback info)
                  in { stmts: firstSub.statements <> secondSub.statements <>
                         [ GoMutate targetName (GoBinOp "+" (GoVar firstSub.leftName) (GoVar secondSub.leftName)) ]
                     , nextId: secondSub.nextId
                     }
                Nothing -> directCallFallback info
            in emitGuardedSelf context nextId info leftArgs fallback
          Nothing -> ordinaryLeft
        Nothing -> ordinaryLeft
      ordinaryLeft =
        let
          first = translate (childContext context (Just Int)) nextId left
          leftName = names.prefix <> "left_" <> show first.nextId
        in { statements: flattenStmts first.stmts <>
               [ GoAssign leftName (native TypeInt64 (coerceGoExpr codegenStateRef modNameStr first.expr first.exprType TypeInt64)) ]
           , leftName
           , nextId: first.nextId + 1
           }
      captured = foldl (\current (Tuple (Tuple _ ty) argument) ->
        let result = translate (childContext context Nothing) current.nextId argument
            argName = names.prefix <> "arg_" <> show result.nextId
        in { statements: current.statements <> flattenStmts result.stmts <>
               [ GoAssign argName (native ty (coerceGoExpr codegenStateRef modNameStr result.expr result.exprType ty)) ]
           , values: Array.snoc current.values (GoVar argName)
           , nextId: result.nextId + 1
           })
        { statements: emitLeft.statements
        , values: []
        , nextId: emitLeft.nextId
        } (Array.zip params args)
      updates = Array.zipWith (\(Tuple param _) value -> GoMutate (param <> "_loop") value) params captured.values
    in
      { statements: captured.statements <>
          -- Accumulate in raw wrapping int64; Final IntAdd narrows once mod 2^32.
          [ GoMutate names.acc (GoBinOp "+" (GoVar names.acc) (GoVar emitLeft.leftName))
          , GoMutate names.used (GoVar "true")
          ] <> updates <> [ GoContinue names.label ]
      , nextId: captured.nextId
      }

add :: GoExpr -> GoExpr -> GoExpr
add left right = GoCall (GoSelector (GoVar "gopurs_runtime") "IntAdd") [ left, right ]

-- A captured untyped literal must keep the worker ABI, rather than defaulting
-- to Go's machine int in a short declaration.
native :: GoType -> GoExpr -> GoExpr
native TypeInt64 value = GoCall (GoVar "int64") [ value ]
native TypeFloat64 value = GoCall (GoVar "float64") [ value ]
native TypeUint32 value = GoCall (GoVar "uint32") [ value ]
native _ value = value

-- Reserve a complete prefix, including temporary slots and the loop label.
-- The suffix chosen by another emitter cannot collide with this family.
freshPrefix :: Array String -> String
freshPrefix used = choose 0
  where
  choose index =
    let prefix = "__gopurs_tail_add_" <> show index <> "_"
    in if Array.any (String.contains (Pattern prefix)) used then choose (index + 1) else prefix

bindingNames :: TcoExpr -> Array String
bindingNames (TcoExpr _ syntax) = own <> foldMap bindingNames syntax
  where
  refs = map (\(Tuple name level) -> localId name level)
  own = case syntax of
    Local name level -> [ localId name level ]
    Let name level _ _ -> [ localId name level ]
    LetRec level bindings _ -> map (\(Tuple name _) -> localId (Just name) level) (NEA.toArray bindings)
    Abs args _ -> refs (NEA.toArray args)
    UncurriedAbs args _ -> refs args
    UncurriedEffectAbs args _ -> refs args
    EffectBind name level _ _ -> [ localId name level ]
    _ -> []
