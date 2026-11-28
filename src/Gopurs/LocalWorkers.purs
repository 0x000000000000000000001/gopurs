module Gopurs.LocalWorkers
  ( parameters
  , signature
  , needsWrapper
  , emit
  , emitZeroSeed
  ) where

import Prelude
import Data.Array as Array
import Data.Foldable (any, foldMap, foldl)
import Data.Map as Map
import Data.Maybe (Maybe(..))
import Data.Set (Set)
import Data.Set as Set
import Data.String as String
import Data.String.Pattern (Pattern(..))
import Data.Tuple (Tuple(..), fst, snd)
import Effect.Ref as Ref
import Effect.Unsafe (unsafePerformEffect)
import Gopurs.CodegenState (FunctionInfo)
import Gopurs.CallAnalysis (collectGoSpine, getGoSpineArgs)
import Gopurs.ExprAnalysis (extractExprFuncType, getExprType, unwrapTcoExpr)
import Gopurs.ExprContext (ExprContext, ExprResult, TranslateExpr, flattenStmts)
import Gopurs.FilteredSumLoops as FilteredSumLoops
import Gopurs.GoAst (GoDecl(..), GoExpr(..), GoType(..), goTypeToStr, rawGo)
import Gopurs.GoSymbol (GoReference(..), generated)
import Gopurs.GoConversions (boxGoExpr, coerceGoExpr)
import Gopurs.GoFunctions (Parameters, curriedFunction, iterationBindings, loopParameters, namedParameters)
import Gopurs.Int32Loops as Int32Loops
import Gopurs.AffineLoops as AffineLoops
import Gopurs.GoTypes (exprTypeToGoType)
import PureScript.Backend.Optimizer.Codegen.Tco (TcoExpr(..))
import PureScript.Backend.Optimizer.CoreFn (Literal(..))
import PureScript.Backend.Optimizer.FreeVars (localId)
import PureScript.Backend.Optimizer.Syntax (BackendSyntax(..))

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
emit :: TranslateExpr -> ExprContext -> String -> Parameters -> Boolean -> Boolean -> Maybe TcoExpr -> ExprResult -> { declarations :: Array GoExpr, assignments :: Array GoExpr, hoisted :: Maybe GoReference }
emit translate context@{ codegenStateRef, modNameStr, bound } name params keepWrapper standalone loopBody result =
  let
    nativeName = workerName modNameStr name
    resultType = result.exprType
    body = flattenStmts result.stmts <> [ GoReturn result.expr ]
    looped = Int32Loops.loop name params resultType body
    closedLoop source =
      let
        allowed = Set.fromFoldable (map fst params)
        free = freeLocalsIn allowed source
        isSelfKey k = case Map.lookup k bound of
          Just b -> b.name == name
          Nothing -> false
      in
        Array.all (\k -> isSelfKey k || Set.member k allowed) (Array.fromFoldable free)
    hoistedHot = case loopBody of
      Just source | standalone && not keepWrapper && closedLoop source ->
        let
          selfTarget = AffineLoops.localTarget bound name
          coldArgs = map (GoVar <<< fst) (loopParameters params)
          makeSplit hotFn =
            let
              coldBody = AffineLoops.wrap translate context selfTarget params resultType source looped
              coldDecl coldName = GoFunctionDecl
                { name: generated modNameStr coldName
                , params: loopParameters params
                , result: resultType
                , body: coldBody
                }
              hotDecl hotName fastBody = GoFunctionDecl
                { name: generated modNameStr hotName
                , params: loopParameters params
                , result: resultType
                , body: fastBody
                }
            in
              unsafePerformEffect do
                state <- Ref.read codegenStateRef
                let
                  coldName = "Call_" <> modNameStr <> "___gopurs_cold_local_" <> show state.globalId
                  hotName = "Call_" <> modNameStr <> "___gopurs_hot_local_" <> show (state.globalId + 1)
                  coldCall = GoCall (GoRef (SymbolReference (generated modNameStr coldName))) coldArgs
                case hotFn selfTarget params resultType source coldCall of
                  Just fastBody -> do
                    Ref.modify_ (\current -> current
                      { globalId = current.globalId + 2
                      , declarations = current.declarations <> [ coldDecl coldName, hotDecl hotName fastBody ]
                      }) codegenStateRef
                    pure (Just (SymbolReference (generated modNameStr hotName)))
                  Nothing -> pure Nothing
        in
          case makeSplit FilteredSumLoops.hot of
            Just ref -> Just ref
            Nothing -> makeSplit (AffineLoops.hot translate context)
      _ -> Nothing
    functionBody = case loopBody of
      Just source -> AffineLoops.wrap translate context (AffineLoops.localTarget bound name) params resultType source looped
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
    declarations = case hoistedHot of
      Just _ -> []
      Nothing ->
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
    assignments = case hoistedHot of
      Just _ -> []
      Nothing ->
        (if inlineInit then [] else [ GoMutate nativeName funcLiteral ])
          <> if keepWrapper then [ GoMutate name wrapper ] else []
  in
    { declarations, assignments, hoisted: hoistedHot }

-- Specialize a standalone closed FilteredSumLoops local worker when the entire
-- outer LetRec body is a single saturated call passing literal 0 as the seed.
-- Emits a 1-arg cold loop helper and returns an inlineable fast block (Go
-- inline cost <= 79 when unwrapped by ModuleWorkers.flattenTailReturn).
emitZeroSeed
  :: TranslateExpr
  -> ExprContext
  -> String
  -> String
  -> Parameters
  -> TcoExpr
  -> TcoExpr
  -> ExprResult
  -> Maybe { coldRef :: GoReference, counterArg :: TcoExpr }
emitZeroSeed translate context@{ codegenStateRef, modNameStr, bound } oldName newName params loopBody outerBody result =
  let
    resultType = result.exprType
    allowed = Set.fromFoldable (map fst params)
    free = freeLocalsIn allowed loopBody
    isSelfKey k = case Map.lookup k bound of
      Just b -> b.name == newName
      Nothing -> false
    isClosed = Array.all (\k -> isSelfKey k || Set.member k allowed) (Array.fromFoldable free)
    callArgs = case unwrapTcoExpr outerBody of
      App _ _ ->
        let Tuple fn spine = collectGoSpine outerBody
        in case unwrapTcoExpr fn of
          Local ident level | localId ident level == oldName -> Just (getGoSpineArgs spine)
          _ -> Nothing
      UncurriedApp fn args -> case unwrapTcoExpr fn of
        Local ident level | localId ident level == oldName -> Just args
        _ -> Nothing
      _ -> Nothing
  in
    if not isClosed then Nothing
    else case callArgs of
      Just [ a0, a1 ] ->
        let
          selfTarget = AffineLoops.localTarget bound newName
          body = flattenStmts result.stmts <> [ GoReturn result.expr ]
          looped = Int32Loops.loop newName params resultType body
          coldInner = AffineLoops.wrap translate context selfTarget params resultType loopBody looped
        in
          case FilteredSumLoops.hotZero selfTarget params resultType loopBody of
            Just hz ->
              let
                accArg = if hz.accIdx == 0 then a0 else a1
                counterArg = if hz.counterIdx == 0 then a0 else a1
                accIsZero = case unwrapTcoExpr accArg of
                  Lit (LitInt 0) -> true
                  _ -> false
                counterIsLocal = case unwrapTcoExpr counterArg of
                  Local _ _ -> true
                  _ -> false
                counterMentionsSelf = Set.member oldName (freeLocalsIn Set.empty counterArg)
              in
                if accIsZero && counterIsLocal && not counterMentionsSelf then
                  unsafePerformEffect do
                    state <- Ref.read codegenStateRef
                    let
                      coldName = "Call_" <> modNameStr <> "___gopurs_cold_local_" <> show state.globalId
                      coldDecl = GoFunctionDecl
                        { name: generated modNameStr coldName
                        , params: [ Tuple hz.counterParam TypeInt64 ]
                        , result: resultType
                        , body: GoBlock
                            [ GoAssign hz.accParam (GoCall (GoVar "int64") [ GoInt 0 ])
                            , rawGo ("_ = " <> hz.accParam)
                            , coldInner
                            ]
                        }
                    Ref.modify_ (\current -> current
                      { globalId = current.globalId + 1
                      , declarations = current.declarations <> [ coldDecl ]
                      }) codegenStateRef
                    pure (Just { coldRef: SymbolReference (generated modNameStr coldName), counterArg })
                else Nothing
            Nothing -> Nothing
      _ -> Nothing

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

freeLocalsIn :: Set String -> TcoExpr -> Set String
freeLocalsIn bound (TcoExpr _ syntax) = case syntax of
  Local ident level ->
    let key = localId ident level
    in if Set.member key bound then Set.empty else Set.singleton key
  Let ident level val body ->
    freeLocalsIn bound val <> freeLocalsIn (Set.insert (localId ident level) bound) body
  EffectBind ident level val body ->
    freeLocalsIn bound val <> freeLocalsIn (Set.insert (localId ident level) bound) body
  Abs params body ->
    freeLocalsIn (foldl (\acc (Tuple ident level) -> Set.insert (localId ident level) acc) bound params) body
  UncurriedAbs params body ->
    freeLocalsIn (foldl (\acc (Tuple ident level) -> Set.insert (localId ident level) acc) bound params) body
  UncurriedEffectAbs params body ->
    freeLocalsIn (foldl (\acc (Tuple ident level) -> Set.insert (localId ident level) acc) bound params) body
  LetRec level bindings body ->
    let bound' = foldl (\acc (Tuple ident _) -> Set.insert (localId (Just ident) level) acc) bound bindings
    in foldMap (\(Tuple _ val) -> freeLocalsIn bound' val) bindings <> freeLocalsIn bound' body
  _ -> foldMap (freeLocalsIn bound) syntax
