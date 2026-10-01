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
import Data.Tuple (Tuple(..), snd)
import Effect.Ref as Ref
import Effect.Unsafe (unsafePerformEffect)
import Gopurs.CallAnalysis (extractUncurriedAbs)
import Gopurs.ExprAnalysis (extractExprFuncType, getExprType, printTcoExprShape)
import Gopurs.ExprContext (ExprContext, ExprResult, TranslateExpr, StmtTree(..), bindParameters)
import Gopurs.GoAst (rawGo, GoExpr(..), GoType(..), goTypeToStr, sanitizeName)
import Gopurs.GoConversions (coerceGoExpr)
import Gopurs.GoTypes (exprTypeToGoType, printExprType)
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
recursive translate context@{ metadata, codegenStateRef, depth, modNameStr, recVars, moduleFunctions, bound } nextId (TcoExpr tcoAnalysis _) lvl bindings body =
  let
    allocRes = foldl
      ( \acc (Tuple (Ident ident) val) ->
          let
            oldName = localId (Just (Ident ident)) lvl
            gId = unsafePerformEffect do
              curr <- Ref.read codegenStateRef
              Ref.modify_ (\r -> r { globalId = r.globalId + 1 }) codegenStateRef
              pure curr.globalId
            newName = oldName <> "_" <> show acc.nextId <> "_" <> show gId
            expectedGoTypeFromAst = exprTypeToGoType metadata.pointerAdtPaths metadata.enumAdts metadata.elidedCtors modNameStr (getExprType val)
          in
            { newBound: Map.insert oldName { name: newName, goType: expectedGoTypeFromAst } acc.newBound, newNames: Array.snoc acc.newNames { oldName, newName }, nextId: acc.nextId + 1 }
      )
      { newBound: bound, newNames: [], nextId }
      (toArray bindings)

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
                  boundInfo = fromMaybe { name: oldName, goType: TypeValue } (Map.lookup oldName allocRes.newBound)
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
            { functions: moduleFunctions, bound: allocRes.newBound }
            workers

          resData = foldl
            ( \acc worker ->
                let
                  { oldName, newName, params } = worker
                  currentLoopCtx = [ { ident: newName, loopParams: map (\(Tuple p _) -> p <> "_loop") params, goTypes: map snd params } ]
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
            { declarations: [], stmts: [], nextId: allocRes.nextId, moduleFunctions: published.functions, newBound: published.bound }
            workers

          resBodyOuter = translate (context { depth = (depth + 1), recVars = combinedRecVars, moduleFunctions = resData.moduleFunctions, bound = resData.newBound, tcoIdent = Nothing, mbExpectedExprType = Nothing }) resData.nextId body
        in
          -- Every function in the recursive group must be in scope
          -- before emitting bodies that can refer to its peers.
          { stmts: foldMap StmtLeaf (resData.declarations <> resData.stmts) <> resBodyOuter.stmts, expr: resBodyOuter.expr, exprType: resBodyOuter.exprType, nextId: resBodyOuter.nextId }

      Nothing ->
        let
          -- An initializer must not observe Go's zero value for an
          -- uninitialized recursive binding. Publish a pointer only
          -- after its complete initializer has run; closures capture
          -- that cell and can safely read it once initialization ends.
          initializingBound = foldl
            ( \acc alloc -> Map.update
                (\binding -> Just (binding { name = "(*" <> alloc.newName <> "_cell)" }))
                alloc.oldName
                acc
            )
            allocRes.newBound
            allocRes.newNames

          accBindings = foldl
            ( \acc (Tuple (Tuple _ val) alloc) ->
                let
                  res = translate (context { depth = (depth + 1), recVars = combinedRecVars, bound = initializingBound, tcoIdent = Nothing, loopCtx = [], options = { isTail: false, inEffectBlock: false }, mbExpectedExprType = Nothing }) acc.nextId val
                  expectedGoType = (fromMaybe { name: alloc.newName, goType: TypeValue } (Map.lookup alloc.oldName allocRes.newBound)).goType
                  assignedVal = coerceGoExpr codegenStateRef modNameStr res.expr res.exprType expectedGoType
                in
                  { stmts: acc.stmts <> res.stmts
                      <> StmtLeaf (GoMutate alloc.newName assignedVal)
                      <> StmtLeaf (GoMutate (alloc.newName <> "_cell") (rawGo ("&" <> alloc.newName)))
                  , exprs: Array.snoc acc.exprs { key: alloc.newName, goType: expectedGoType }
                  , nextId: res.nextId
                  }
            )
            { stmts: StmtEmpty, exprs: [], nextId: allocRes.nextId }
            (Array.zip (toArray bindings) allocRes.newNames)

          declStmts = map (\b -> rawGo ("var " <> b.key <> " " <> goTypeToStr b.goType <> "\n_ = " <> b.key
            <> "\nvar " <> b.key <> "_cell *" <> goTypeToStr b.goType <> "\n_ = " <> b.key <> "_cell"
            <> "\n// FALLBACK TCO: isLoop=" <> show isLoop <> " len=" <> show (Array.length (toArray bindings)))) accBindings.exprs

          resBody = translate (context { depth = (depth + 1), recVars = combinedRecVars, bound = allocRes.newBound, tcoIdent = Nothing }) accBindings.nextId body
        in
          { stmts: foldMap StmtLeaf declStmts <> accBindings.stmts <> resBody.stmts, expr: resBody.expr, exprType: resBody.exprType, nextId: resBody.nextId }
