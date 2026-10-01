module Gopurs.BindingExprs
  ( nonRecursive
  , recursive
  ) where

import Prelude
import Data.Array as Array
import Data.Array.NonEmpty (NonEmptyArray, toArray)
import Data.Foldable (foldMap, foldl)
import Data.Map as Map
import Data.Maybe (Maybe(..), fromMaybe)
import Data.Newtype (unwrap)
import Data.Traversable (traverse)
import Data.Tuple (Tuple(..), fst, snd)
import Effect.Ref as Ref
import Effect.Unsafe (unsafePerformEffect)
import Gopurs.CallAnalysis (extractUncurriedAbs)
import Gopurs.ExprAnalysis (extractExprFuncType, getExprType, printTcoExprShape)
import Gopurs.ExprContext (ExprContext, ExprResult, LocalEnv, TranslateExpr, StmtTree(..), bindParameters, childContext)
import Gopurs.GoAst (rawGo, GoExpr(..), GoType(..), goTypeToStr, sanitizeName)
import Gopurs.GoConversions (coerceGoExpr)
import Gopurs.GoTypes (exprTypeToGoType, printExprType)
import Gopurs.GoFunctions (loopParameters)
import Gopurs.LocalWorkers as LocalWorkers
import Gopurs.Printer (printGoExpr)
import PureScript.Backend.Optimizer.Codegen.Tco (TcoExpr(..))
import PureScript.Backend.Optimizer.CoreFn (Ident(..))
import PureScript.Backend.Optimizer.FreeVars (localId)
import PureScript.Backend.Optimizer.Syntax (Level)

nonRecursive :: TranslateExpr -> ExprContext -> Int -> Maybe Ident -> Level -> TcoExpr -> TcoExpr -> ExprResult
nonRecursive translate context@{ metadata, codegenStateRef, depth, modNameStr, moduleFunctions, bound } nextId mbIdent lvl binding body =
  let
    originalName = localId mbIdent lvl
    name = originalName <> "_" <> show nextId
    expectedGoTypeFromAst = exprTypeToGoType metadata.pointerAdtPaths metadata.enumAdts metadata.elidedCtors modNameStr (getExprType binding)

    mbFunc = extractUncurriedAbs binding
  in
    case mbFunc of
      Just abs | Array.length abs.args > 0 ->
        let
          params = LocalWorkers.parameters context abs.args binding
          loopBound = bindParameters params bound
          resBodyMut = translate (context { depth = (depth + 1), bound = loopBound, tcoIdent = (Just name), loopCtx = [], options = { isTail: true, inEffectBlock: false }, mbExpectedExprType = Nothing }) (nextId + 1) abs.body
          -- The binding is non-recursive: infer its native result before
          -- publishing the signature to callers. Only its curried wrapper
          -- crosses the Value boundary, which can otherwise copy whole lists.
          trueFRet = resBodyMut.exprType
          localModuleFunctions = Map.insert name (LocalWorkers.signature context name params trueFRet) moduleFunctions
          worker = LocalWorkers.emit context name params false resBodyMut

          newBound = Map.insert originalName { name, goType: TypeValue } bound
          resBodyOuter = translate (context { depth = (depth + 1), moduleFunctions = localModuleFunctions, bound = newBound, tcoIdent = Nothing, mbExpectedExprType = Nothing }) resBodyMut.nextId body
        in
          { stmts: foldMap StmtLeaf (worker.declarations <> worker.assignments) <> resBodyOuter.stmts, expr: resBodyOuter.expr, exprType: resBodyOuter.exprType, nextId: resBodyOuter.nextId }

      _ ->
        let
          resBinding = translate (context { depth = (depth + 1), tcoIdent = Nothing, loopCtx = [], options = { isTail: false, inEffectBlock: false }, mbExpectedExprType = Nothing }) (nextId + 1) binding
          actualGoType = if expectedGoTypeFromAst == TypeValue then resBinding.exprType else expectedGoTypeFromAst
          newBound = Map.insert originalName { name, goType: actualGoType } bound
          resBody = translate (context { depth = (depth + 1), bound = newBound, tcoIdent = Nothing }) resBinding.nextId body
          letStmt =
            if actualGoType == resBinding.exprType then
              StmtLeaf (GoAssign name resBinding.expr)
            else
              StmtLeaf (rawGo ("var " <> name <> " " <> goTypeToStr actualGoType <> " = " <> printGoExpr (coerceGoExpr codegenStateRef modNameStr resBinding.expr resBinding.exprType actualGoType)))
        in
          { stmts: resBinding.stmts <> StmtLeaf (rawGo ("// TAST (Let): " <> name <> " shape=" <> printTcoExprShape binding <> " bindingType=" <> printExprType (getExprType binding))) <> letStmt <> resBody.stmts, expr: resBody.expr, exprType: resBody.exprType, nextId: resBody.nextId }

recursive :: TranslateExpr -> ExprContext -> Int -> TcoExpr -> Level -> NonEmptyArray (Tuple Ident TcoExpr) -> TcoExpr -> ExprResult
recursive translate context@{ metadata, depth, modNameStr, recVars, moduleFunctions } nextId (TcoExpr tcoAnalysis _) lvl bindings body =
  let
    allocated = allocateRecursive context nextId lvl (toArray bindings)
    combinedRecVars = recVars <> map (\(Tuple (Ident i) _) -> sanitizeName i) (toArray bindings)

    isLoop = (unwrap tcoAnalysis).role.isLoop
    mutRecBinds =
      if isLoop then
        traverse (\(Tuple (Ident name) val) -> map (\abs -> { ident: sanitizeName name, args: abs.args, body: abs.body, val }) (extractUncurriedAbs val)) (toArray bindings)
      else Nothing
  in
    case mutRecBinds of
      Just fns ->
        let
          -- Resolve parameters once, then publish every provisional signature
          -- before translating any body. The annotated return is provisional:
          -- bodies refine the function table in source order.
          workers = map
            ( \fn ->
                let
                  oldName = localId (Just (Ident fn.ident)) lvl
                  boundInfo = fromMaybe { name: oldName, goType: TypeValue } (Map.lookup oldName allocated.newBound)
                  newName = boundInfo.name
                  params = LocalWorkers.parameters context fn.args fn.val
                  expectedReturn = map _.fRet (extractExprFuncType (getExprType fn.val))
                  resultType = case expectedReturn of
                    Just ty -> exprTypeToGoType metadata.pointerAdtPaths metadata.enumAdts metadata.elidedCtors modNameStr ty
                    Nothing -> TypeValue
                in
                  { oldName, newName, params, expectedReturn, resultType, body: fn.body }
            )
            fns
          published = foldl
            (\acc worker ->
              { functions: Map.insert worker.newName (LocalWorkers.signature context worker.newName worker.params worker.resultType) acc.functions
              , bound: Map.insert worker.oldName { name: worker.newName, goType: TypeFunc (map snd worker.params) worker.resultType } acc.bound
              })
            { functions: moduleFunctions, bound: allocated.newBound }
            workers

          translatedWorkers = foldl
            ( \acc worker ->
                let
                  { oldName, newName, params } = worker
                  currentLoopCtx = [ { ident: newName, loopParams: map fst (loopParameters params), goTypes: map snd params } ]
                  -- Peer bindings keep the provisional view; only the function
                  -- table sees results inferred for earlier bodies in the group.
                  loopBound = bindParameters params published.bound
                  resBodyMut = translate (context { depth = (depth + 1), recVars = combinedRecVars, moduleFunctions = acc.moduleFunctions, bound = loopBound, tcoIdent = (Just newName), loopCtx = currentLoopCtx, options = { isTail: true, inEffectBlock: false }, mbExpectedExprType = worker.expectedReturn }) acc.nextId worker.body
                  trueFRet = resBodyMut.exprType
                  emitted = LocalWorkers.emit context newName params true resBodyMut
                  newFunctions = Map.insert newName (LocalWorkers.signature context newName params trueFRet) acc.moduleFunctions
                  newBound2 = Map.insert oldName { name: newName, goType: TypeFunc (map snd params) trueFRet } acc.newBound
                in
                  { declarations: acc.declarations <> emitted.declarations, stmts: acc.stmts <> emitted.assignments, nextId: resBodyMut.nextId, moduleFunctions: newFunctions, newBound: newBound2 }
            )
            { declarations: [], stmts: [], nextId: allocated.nextId, moduleFunctions: published.functions, newBound: published.bound }
            workers

          resBodyOuter = translate (context { depth = (depth + 1), recVars = combinedRecVars, moduleFunctions = translatedWorkers.moduleFunctions, bound = translatedWorkers.newBound, tcoIdent = Nothing, mbExpectedExprType = Nothing }) translatedWorkers.nextId body
        in
          -- Every function in the recursive group must be in scope
          -- before emitting bodies that can refer to its peers.
          { stmts: foldMap StmtLeaf (translatedWorkers.declarations <> translatedWorkers.stmts) <> resBodyOuter.stmts, expr: resBodyOuter.expr, exprType: resBodyOuter.exprType, nextId: resBodyOuter.nextId }

      Nothing ->
        initializeRecursiveValues translate (context { recVars = combinedRecVars }) isLoop allocated (toArray bindings) body

type RecursiveAllocation =
  { newBound :: LocalEnv
  , newNames :: Array { oldName :: String, newName :: String }
  , nextId :: Int
  }

-- Reserve both local and module-wide identifiers before inspecting the group.
-- Every initializer receives the same complete set of renamed bindings.
allocateRecursive :: ExprContext -> Int -> Level -> Array (Tuple Ident TcoExpr) -> RecursiveAllocation
allocateRecursive { metadata, codegenStateRef, modNameStr, bound } nextId lvl =
  foldl
    (\acc (Tuple ident value) ->
      let
        oldName = localId (Just ident) lvl
        globalId = unsafePerformEffect do
          state <- Ref.read codegenStateRef
          Ref.modify_ (\current -> current { globalId = current.globalId + 1 }) codegenStateRef
          pure state.globalId
        newName = oldName <> "_" <> show acc.nextId <> "_" <> show globalId
        goType = exprTypeToGoType metadata.pointerAdtPaths metadata.enumAdts metadata.elidedCtors modNameStr (getExprType value)
      in
        { newBound: Map.insert oldName { name: newName, goType } acc.newBound
        , newNames: Array.snoc acc.newNames { oldName, newName }
        , nextId: acc.nextId + 1
        })
    { newBound: bound, newNames: [], nextId }

-- An early read must not observe Go's zero value. Initializers and the closures
-- they create read through cells; publish each pointer only after its complete
-- initializer and coercion. The enclosing body uses the initialized values.
initializeRecursiveValues :: TranslateExpr -> ExprContext -> Boolean -> RecursiveAllocation -> Array (Tuple Ident TcoExpr) -> TcoExpr -> ExprResult
initializeRecursiveValues translate context@{ codegenStateRef, modNameStr, depth } isLoop allocated bindings body =
  let
    initializingBound = foldl
      (\acc name -> Map.update (\binding -> Just (binding { name = "(*" <> name.newName <> "_cell)" })) name.oldName acc)
      allocated.newBound allocated.newNames
    initialized = foldl
      (\acc (Tuple (Tuple _ value) name) ->
        let
          valueResult = translate ((childContext context Nothing) { bound = initializingBound }) acc.nextId value
          expected = (fromMaybe { name: name.newName, goType: TypeValue } (Map.lookup name.oldName allocated.newBound)).goType
          assigned = coerceGoExpr codegenStateRef modNameStr valueResult.expr valueResult.exprType expected
        in
          { stmts: acc.stmts <> valueResult.stmts
              <> StmtLeaf (GoMutate name.newName assigned)
              <> StmtLeaf (GoMutate (name.newName <> "_cell") (rawGo ("&" <> name.newName)))
          , values: Array.snoc acc.values { name: name.newName, goType: expected }
          , nextId: valueResult.nextId
          })
      { stmts: StmtEmpty, values: [], nextId: allocated.nextId }
      (Array.zip bindings allocated.newNames)
    declarations = map (\value -> rawGo ("var " <> value.name <> " " <> goTypeToStr value.goType <> "\n_ = " <> value.name
      <> "\nvar " <> value.name <> "_cell *" <> goTypeToStr value.goType <> "\n_ = " <> value.name <> "_cell"
      <> "\n// FALLBACK TCO: isLoop=" <> show isLoop <> " len=" <> show (Array.length bindings))) initialized.values
    result = translate (context { depth = depth + 1, bound = allocated.newBound, tcoIdent = Nothing }) initialized.nextId body
  in
    { stmts: foldMap StmtLeaf declarations <> initialized.stmts <> result.stmts, expr: result.expr, exprType: result.exprType, nextId: result.nextId }
