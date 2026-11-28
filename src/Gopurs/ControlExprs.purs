module Gopurs.ControlExprs
  ( failure
  , branch
  , branchWith
  , booleanAnd
  , booleanOr
  ) where

import Prelude
import Data.Array as Array
import Data.Array.NonEmpty (NonEmptyArray, toArray)
import Data.Foldable (any, foldMap, foldl)
import Data.Map as Map
import Data.Maybe (Maybe(..), fromMaybe, isNothing)
import Data.Set (Set)
import Data.Set as Set
import Data.Tuple (Tuple(..), snd)
import Effect.Ref as Ref
import Effect.Unsafe (unsafePerformEffect)
import Gopurs.ExprAnalysis (getExprType, unwrapTcoExpr)
import Gopurs.ExprContext (ExprContext, ExprResult, TranslateExpr, StmtTree(..), flattenStmts)
import Gopurs.GoAst (rawGo, GoDecl(..), GoExpr(..), GoType(..), goTypeToStr)
import Gopurs.GoConversions (coerceGoExpr, unboxGoExpr)
import Gopurs.GoSymbol (GoReference(..), generated)
import Gopurs.GoTypes (exprTypeToGoType)
import Gopurs.Printer (printGoExpr)
import PureScript.Backend.Optimizer.Codegen.Tco (TcoExpr(..))
import PureScript.Backend.Optimizer.CoreFn (ExprType(..))
import PureScript.Backend.Optimizer.FreeVars (localId)
import PureScript.Backend.Optimizer.Syntax (BackendSyntax(..), Pair(..))

failure :: ExprContext -> Int -> TcoExpr -> String -> ExprResult
failure { metadata, modNameStr, mbExpectedExprType, mbExpectedGoType } nextId tcoExpr msg =
  let
    expectedGoType = fromMaybe inferredGoType mbExpectedGoType
    inferredGoType = exprTypeToGoType metadata.sharedRecords metadata.pointerAdtPaths metadata.enumAdts metadata.elidedCtors modNameStr
      ( case getExprType tcoExpr of
          Any -> fromMaybe Any mbExpectedExprType
          ty -> ty
      )
    expectedGoTypeStr = goTypeToStr expectedGoType
  in
    { stmts: StmtEmpty, expr: rawGo ("func() " <> expectedGoTypeStr <> " { panic(" <> printGoExpr (GoString msg) <> ") }()"), exprType: expectedGoType, nextId }

branch :: TranslateExpr -> ExprContext -> Int -> NonEmptyArray (Pair TcoExpr) -> TcoExpr -> ExprResult
branch translate = branchWith translate translate

-- Conditions always use the ordinary translator. A consumer which is already
-- executing the selected result may supply a distinct translator for bodies.
branchWith :: TranslateExpr -> TranslateExpr -> ExprContext -> Int -> NonEmptyArray (Pair TcoExpr) -> TcoExpr -> ExprResult
branchWith translate translateBody context@{ codegenStateRef, depth, modNameStr, mbExpectedGoType, options: { isTail } } nextId branches def =
  let
    resDef = translateBody (context { depth = (depth + 1), tcoIdent = Nothing, options = { isTail, inEffectBlock: false } }) nextId def

    computedBranches = foldl
      ( \acc (Pair condExpr bodyExpr) ->
          let
            resCond = translate (context { depth = (depth + 1), tcoIdent = Nothing, loopCtx = [], options = { isTail: false, inEffectBlock: false }, mbExpectedExprType = Nothing, mbExpectedGoType = Nothing }) acc.nextId condExpr
            resBody = translateBody (context { depth = (depth + 1), tcoIdent = Nothing, options = { isTail, inEffectBlock: false } }) resCond.nextId bodyExpr
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

    -- The consumer fixes the join representation when its ABI is known.
    -- Generic branches adapt at their exit; native branches then avoid a
    -- structural box/unbox round trip through the shared result slot.
    expectedGoType = fromMaybe inferredGoType mbExpectedGoType

    inferredGoType =
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

    outlinedDef = case toArray branches of
      [ Pair _ fastBody ]
        | not isFailNode
        , isNothing context.tcoIdent
        , Array.null context.loopCtx
        , Array.null context.recVars
        , isConcreteGoType expectedGoType
        , hasHeavyFallback def
        , not (hasHeavyFallback fastBody) ->
            outlineColdBranch translateBody context def expectedGoType
      _ -> Nothing

    defStmts = case outlinedDef of
      Just callExpr ->
        StmtLeaf (GoMutate tmpVar callExpr)
      Nothing ->
        StmtLeaf (rawGo "{") <> resDef.stmts <> StmtLeaf (GoMutate tmpVar (coerceGoExpr codegenStateRef modNameStr resDef.expr resDef.exprType expectedGoType)) <> StmtLeaf (rawGo "}")
  in
    { stmts: declTmp <> buildIfs <> defStmts <> StmtLeaf (rawGo (labelName <> ":")), expr: GoVar tmpVar, exprType: expectedGoType, nextId: computedBranches.nextId + 1 }

isConcreteGoType :: GoType -> Boolean
isConcreteGoType = case _ of
  TypeResults _ -> false
  TypeGenericParam _ -> false
  TypePointer inner -> isConcreteGoType inner
  TypeNativeArray inner -> isConcreteGoType inner
  TypeRecord fields -> Array.all (isConcreteGoType <<< snd) fields
  TypeFunc args ret -> Array.all isConcreteGoType args && isConcreteGoType ret
  TypeStructValue _ args -> Array.all isConcreteGoType args
  _ -> true

hasHeavyFallback :: TcoExpr -> Boolean
hasHeavyFallback (TcoExpr _ syntax) = case syntax of
  App _ _ -> true
  UncurriedApp _ _ -> true
  UncurriedEffectApp _ _ -> true
  EffectBind _ _ _ _ -> true
  EffectDefer _ -> true
  PrimEffect _ -> true
  Abs _ _ -> true
  UncurriedAbs _ _ -> true
  UncurriedEffectAbs _ _ -> true
  LetRec _ _ _ -> true
  _ -> any hasHeavyFallback syntax

freeLocals :: Set String -> TcoExpr -> Set String
freeLocals bound (TcoExpr _ syntax) = case syntax of
  Local ident level ->
    let key = localId ident level
    in if Set.member key bound then Set.empty else Set.singleton key
  Let ident level val body ->
    freeLocals bound val <> freeLocals (Set.insert (localId ident level) bound) body
  EffectBind ident level val body ->
    freeLocals bound val <> freeLocals (Set.insert (localId ident level) bound) body
  Abs params body ->
    freeLocals (foldl (\acc (Tuple ident level) -> Set.insert (localId ident level) acc) bound params) body
  UncurriedAbs params body ->
    freeLocals (foldl (\acc (Tuple ident level) -> Set.insert (localId ident level) acc) bound params) body
  UncurriedEffectAbs params body ->
    freeLocals (foldl (\acc (Tuple ident level) -> Set.insert (localId ident level) acc) bound params) body
  LetRec level bindings body ->
    let
      bound' = foldl (\acc (Tuple ident _) -> Set.insert (localId (Just ident) level) acc) bound bindings
    in
      foldMap (\(Tuple _ val) -> freeLocals bound' val) bindings <> freeLocals bound' body
  _ -> foldMap (freeLocals bound) syntax

outlineColdBranch :: TranslateExpr -> ExprContext -> TcoExpr -> GoType -> Maybe GoExpr
outlineColdBranch translateBody context@{ codegenStateRef, modNameStr, bound, moduleFunctions } def expectedGoType =
  let
    keys = Array.fromFoldable (freeLocals Set.empty def)
    hasThunkBinding key =
      Map.member ("$gopursBoundThunk$" <> key) bound
        || Map.member ("$gopursStatelessThunk$" <> key) bound
    resolve key =
      if hasThunkBinding key || Map.member key moduleFunctions then Nothing
      else case Map.lookup key bound of
        Just binding
          | isConcreteGoType binding.goType
          , not (Map.member binding.name moduleFunctions) ->
              Just { key, binding }
        _ -> Nothing
    resolved = Array.mapMaybe resolve keys
  in
    if Array.length keys > 6 || Array.length resolved /= Array.length keys then Nothing
    else
      let
        params = Array.mapWithIndex (\index entry -> Tuple ("cold_arg_" <> show index) entry.binding.goType) resolved
        coldBound = Array.foldl
          (\acc (Tuple index entry) -> Map.insert entry.key { name: "cold_arg_" <> show index, goType: entry.binding.goType } acc)
          Map.empty
          (Array.mapWithIndex Tuple resolved)
        coldModuleFunctions = Map.filter (\info -> case info.reference of
          LocalReference _ -> false
          _ -> true) moduleFunctions
        coldContext = context
          { bound = coldBound
          , moduleFunctions = coldModuleFunctions
          , depth = 0
          , tcoIdent = Nothing
          , loopCtx = []
          , recVars = []
          , options = { isTail: false, inEffectBlock: false }
          }
        coldResult = translateBody coldContext 0 def
        coldReturn = coerceGoExpr codegenStateRef modNameStr coldResult.expr coldResult.exprType expectedGoType
        args = map (\entry -> GoVar entry.binding.name) resolved
      in
        Just (unsafePerformEffect do
          state <- Ref.read codegenStateRef
          let
            coldName = "Call_" <> modNameStr <> "___gopurs_cold_branch_" <> show state.globalId
            decl = GoFunctionDecl
              { name: generated modNameStr coldName
              , params
              , result: expectedGoType
              , body: GoBlock (flattenStmts coldResult.stmts <> [ GoReturn coldReturn ])
              }
          Ref.modify_ (\current -> current
            { globalId = current.globalId + 1
            , declarations = current.declarations <> [ decl ]
            }) codegenStateRef
          pure (GoCall (GoRef (SymbolReference (generated modNameStr coldName))) args))

booleanAnd :: TranslateExpr -> ExprContext -> Int -> TcoExpr -> TcoExpr -> ExprResult
booleanAnd translate context@{ codegenStateRef, depth, modNameStr } nextId e1 e2 =
  let
    res1 = translate (context { depth = (depth + 1), tcoIdent = Nothing, loopCtx = [], options = { isTail: false, inEffectBlock: false }, mbExpectedExprType = Nothing, mbExpectedGoType = Nothing }) nextId e1
    res2 = translate (context { depth = (depth + 1), tcoIdent = Nothing, loopCtx = [], options = { isTail: false, inEffectBlock: false }, mbExpectedExprType = Nothing, mbExpectedGoType = Nothing }) res1.nextId e2
  in
    if isEmptyStmts res2.stmts then
      { expr: GoBinOp "&&" (unboxGoExpr codegenStateRef modNameStr res1.expr res1.exprType TypeBool) (unboxGoExpr codegenStateRef modNameStr res2.expr res2.exprType TypeBool), exprType: TypeBool, stmts: res1.stmts <> res2.stmts, nextId: res2.nextId }
    else
      let
        tmpVar = "__t_and_" <> show res2.nextId
        declTmp = StmtLeaf (rawGo ("var " <> tmpVar <> " bool = false"))
        guardIf = StmtLeaf (GoIfElse (unboxGoExpr codegenStateRef modNameStr res1.expr res1.exprType TypeBool)
          (flattenStmts res2.stmts <> [ GoMutate tmpVar (unboxGoExpr codegenStateRef modNameStr res2.expr res2.exprType TypeBool) ]) [])
      in
        { expr: GoVar tmpVar, exprType: TypeBool, stmts: res1.stmts <> declTmp <> guardIf, nextId: res2.nextId + 1 }

booleanOr :: TranslateExpr -> ExprContext -> Int -> TcoExpr -> TcoExpr -> ExprResult
booleanOr translate context@{ codegenStateRef, depth, modNameStr } nextId e1 e2 =
  let
    res1 = translate (context { depth = (depth + 1), tcoIdent = Nothing, loopCtx = [], options = { isTail: false, inEffectBlock: false }, mbExpectedExprType = Nothing, mbExpectedGoType = Nothing }) nextId e1
    res2 = translate (context { depth = (depth + 1), tcoIdent = Nothing, loopCtx = [], options = { isTail: false, inEffectBlock: false }, mbExpectedExprType = Nothing, mbExpectedGoType = Nothing }) res1.nextId e2
  in
    if isEmptyStmts res2.stmts then
      { expr: GoBinOp "||" (unboxGoExpr codegenStateRef modNameStr res1.expr res1.exprType TypeBool) (unboxGoExpr codegenStateRef modNameStr res2.expr res2.exprType TypeBool), exprType: TypeBool, stmts: res1.stmts <> res2.stmts, nextId: res2.nextId }
    else
      let
        tmpVar = "__t_or_" <> show res2.nextId
        declTmp = StmtLeaf (rawGo ("var " <> tmpVar <> " bool = true"))
        guardIf = StmtLeaf (GoIfElse (GoPrefixOp "!" (unboxGoExpr codegenStateRef modNameStr res1.expr res1.exprType TypeBool))
          (flattenStmts res2.stmts <> [ GoMutate tmpVar (unboxGoExpr codegenStateRef modNameStr res2.expr res2.exprType TypeBool) ]) [])
      in
        { expr: GoVar tmpVar, exprType: TypeBool, stmts: res1.stmts <> declTmp <> guardIf, nextId: res2.nextId + 1 }

isEmptyStmts :: StmtTree -> Boolean
isEmptyStmts StmtEmpty = true
isEmptyStmts (StmtAppend s1 s2) = isEmptyStmts s1 && isEmptyStmts s2
isEmptyStmts _ = false
