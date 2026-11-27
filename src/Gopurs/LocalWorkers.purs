module Gopurs.LocalWorkers
  ( parameters
  , signature
  , needsWrapper
  , emit
  ) where

import Prelude
import Data.Array as Array
import Data.Foldable (any)
import Data.Maybe (Maybe(..))
import Data.String as String
import Data.String.Pattern (Pattern(..))
import Data.Tuple (Tuple(..), snd)
import Gopurs.CodegenState (FunctionInfo)
import Gopurs.CallAnalysis (collectGoSpine, getGoSpineArgs)
import Gopurs.ExprAnalysis (extractExprFuncType, getExprType, unwrapTcoExpr)
import Gopurs.ExprContext (ExprContext, ExprResult, TranslateExpr, flattenStmts)
import Gopurs.GoAst (GoExpr(..), GoType(..), goTypeToStr, rawGo)
import Gopurs.GoSymbol (GoReference(..))
import Gopurs.GoConversions (boxGoExpr, coerceGoExpr)
import Gopurs.GoFunctions (Parameters, curriedFunction, iterationBindings, loopParameters, namedParameters)
import Gopurs.Int32Loops as Int32Loops
import Gopurs.AffineLoops as AffineLoops
import Gopurs.GoTypes (exprTypeToGoType)
import PureScript.Backend.Optimizer.Codegen.Tco (TcoExpr)
import PureScript.Backend.Optimizer.FreeVars (localId)
import PureScript.Backend.Optimizer.Syntax (BackendSyntax(App, Local))

-- Local workers use annotation slots directly. Module workers have a separate
-- ABI policy, including read-only record projections and native sum results.
parameters :: ExprContext -> Array String -> TcoExpr -> Parameters
parameters { metadata, modNameStr } names value =
  let
    types = case extractExprFuncType (getExprType value) of
      Just { fArgs } -> map (exprTypeToGoType metadata.sharedRecords metadata.pointerAdtPaths metadata.enumAdts metadata.elidedCtors modNameStr) fArgs
      Nothing -> []
  in
    namedParameters names types

workerName :: String -> String -> String
workerName modName name = "Call_local_" <> modName <> "_" <> name

signature :: ExprContext -> String -> Parameters -> GoType -> FunctionInfo
signature { modNameStr } name params result =
  { arrayFold: Nothing, codePointFold: Nothing, reference: LocalReference (workerName modNameStr name), fArgs: map snd params, fRet: result, arity: Array.length params }

-- Inspect the current TCO expression, not usage annotations from an earlier
-- IR. Saturated curried calls use the published native worker directly.
-- Every other occurrence keeps the wrapper, including partial applications,
-- returned functions, callback arguments and uses in peer closures.
needsWrapper :: String -> Int -> TcoExpr -> Boolean
needsWrapper target arity expression =
  let
    visit = needsWrapper target arity
    syntax = unwrapTcoExpr expression
  in case syntax of
    Local ident level -> localId ident level == target
    App _ _ ->
      let
        Tuple fn spine = collectGoSpine expression
        args = getGoSpineArgs spine
      in case unwrapTcoExpr fn of
        Local ident level | localId ident level == target && arity > 0 && Array.length args >= arity && onlyCurriedApplications expression -> any visit args
        _ -> any visit syntax
    _ -> any visit syntax

-- collectGoSpine also flattens uncurried applications. Keep their wrappers,
-- including when an uncurried application is nested under a curried one.
onlyCurriedApplications :: TcoExpr -> Boolean
onlyCurriedApplications expression = case unwrapTcoExpr expression of
  App fn _ -> onlyCurriedApplications fn
  Local _ _ -> true
  _ -> false

-- The translated body determines the native result. Only the curried wrapper
-- crosses Value. Declarations stay separate so recursive callers can place the
-- whole group's declarations before any assignment captures a peer worker.
-- When a standalone worker does not reference itself (e.g. non-recursive or
-- TCO-rewritten loop), initializing `var name func(...) R = func(...) R { ... }`
-- lets Go statically devirtualize and inline direct calls.
emit :: TranslateExpr -> ExprContext -> String -> Parameters -> Boolean -> Boolean -> Maybe TcoExpr -> ExprResult -> { declarations :: Array GoExpr, assignments :: Array GoExpr }
emit translate context@{ codegenStateRef, modNameStr, bound } name params keepWrapper standalone loopBody result =
  let
    nativeName = workerName modNameStr name
    resultType = result.exprType
    body = flattenStmts result.stmts <> [ GoReturn result.expr ]
    functionBody = case loopBody of
      Just source -> AffineLoops.wrap translate context (AffineLoops.localTarget bound name) params resultType source
        (Int32Loops.loop name params resultType body)
      Nothing -> GoBlock (iterationBindings params <> body)
    funcTypeStr = "func(" <> String.joinWith ", " (map (goTypeToStr <<< snd) params) <> ") " <> goTypeToStr resultType
    funcLiteral = GoFuncBlock (loopParameters params) [ functionBody ] resultType
    inlineInit = standalone && not (mentionsName nativeName functionBody)
    wrapperDeclarations =
      if keepWrapper then
        [ rawGo ("var " <> name <> " gopurs_runtime.Value")
        , rawGo ("_ = " <> name)
        ]
      else []
    declarations =
      (if inlineInit then
        [ GoMutate ("var " <> nativeName <> " " <> funcTypeStr) funcLiteral
        , rawGo ("_ = " <> nativeName)
        ]
      else
        [ rawGo ("var " <> nativeName <> " " <> funcTypeStr)
        , rawGo ("_ = " <> nativeName)
        ]) <> wrapperDeclarations
    call = GoCall (GoVar nativeName) (map (\(Tuple param ty) ->
      coerceGoExpr codegenStateRef modNameStr (GoVar (param <> "_loop_val")) TypeValue ty) params)
    boxedResult = boxGoExpr codegenStateRef modNameStr call resultType
    wrapper = if Array.null params then
      curriedFunction [ Tuple "_" TypeValue ] TypeValue (GoBlock [ GoReturn boxedResult ])
      else Array.foldr (\(Tuple param _) inner ->
        GoCall (GoSelector (GoVar "gopurs_runtime") "Func")
          [ GoFuncLit [ Tuple (param <> "_loop_val") TypeValue ] [] inner TypeValue ]) boxedResult params
    assignments =
      (if inlineInit then [] else [ GoMutate nativeName funcLiteral ])
        <> if keepWrapper then [ GoMutate name wrapper ] else []
  in
    { declarations, assignments }

mentionsName :: String -> GoExpr -> Boolean
mentionsName target = go
  where
  many exprs = Array.any go exprs
  go = case _ of
    GoVar v -> v == target
    GoRef (LocalReference v) -> v == target
    GoRef _ -> false
    GoString _ -> false
    GoInt _ -> false
    GoCall fn args -> go fn || many args
    GoTypeApply fn _ -> go fn
    GoSelector obj _ -> go obj
    GoBlock stmts -> many stmts
    GoReturn e -> go e
    GoAssign _ e -> go e
    GoRecordDict _ props -> Array.any (go <<< snd) props
    GoRecordUpdateDict obj props -> go obj || Array.any (go <<< snd) props
    GoRecordUpdateStatic obj _ updates fallback -> go obj || Array.any (go <<< snd) updates || Array.any (go <<< snd) fallback
    GoRecordUpdateNative _ obj props -> go obj || Array.any (go <<< snd) props
    GoIIFE _ val b -> go val || go b
    GoLetRec binds b -> Array.any (go <<< snd) binds || go b
    GoRecordAccess obj _ -> go obj
    GoStructAccess obj _ -> go obj
    GoRecordFieldAccess obj _ -> go obj
    GoRecordAccessStatic obj _ _ -> go obj
    GoConstructor _ _ _ args -> many args
    GoConstructorDict _ args -> many args
    GoConstructorAccess obj _ _ _ _ -> go obj
    GoBranch branches def -> Array.any (\(Tuple c t) -> go c || go t) branches || go def
    GoBinOp _ l r -> go l || go r
    GoPrefixOp _ e -> go e
    GoTypeAssertion e _ -> go e
    GoIndex e i -> go e || go i
    GoBoxStructPointer _ e -> go e
    GoBoxIntArray e -> go e
    GoUnboxIntArray e -> go e
    GoFreshFilterArray e -> go e
    GoRaw code -> String.contains (Pattern target) code.text
    GoFor _ stmts -> many stmts
    GoForRange range stmts -> String.contains (Pattern target) range || many stmts
    GoContinue _ -> false
    GoMutate lhs rhs -> String.contains (Pattern target) lhs || go rhs
    GoIfElse c t f -> go c || many t || many f
    GoFuncBlock _ stmts _ -> many stmts
    GoFuncLit _ stmts e _ -> many stmts || go e
    GoStructValue _ _ args -> many args
