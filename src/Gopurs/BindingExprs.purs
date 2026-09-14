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
import Data.String as String
import Data.Traversable (traverse)
import Data.Tuple (Tuple(..), snd)
import Effect.Ref as Ref
import Effect.Unsafe (unsafePerformEffect)
import Gopurs.CallAnalysis (extractUncurriedAbs)
import Gopurs.ExprAnalysis (extractExprFuncType, getExprType, printTcoExprShape)
import Gopurs.ExprContext (ExprContext, ExprResult, TranslateExpr, StmtTree(..), flattenStmts)
import Gopurs.GoAst (GoExpr(..), GoType(..), goTypeToStr, sanitizeName)
import Gopurs.GoConversions (boxGoExpr, coerceGoExpr)
import Gopurs.GoTypes (exprTypeToGoType, printExprType)
import Gopurs.Printer (printGoExpr)
import PureScript.Backend.Optimizer.Codegen.Tco (TcoExpr(..))
import PureScript.Backend.Optimizer.CoreFn (ExprType(..), Ident(..))
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
          fArgsAst = case extractExprFuncType (getExprType binding) of
            Just { fArgs } -> fArgs
            Nothing -> []
          paramsWithTypes = Array.zipWith
            ( \idStr ty ->
                Tuple idStr (exprTypeToGoType metadata.pointerAdtPaths metadata.enumAdts metadata.elidedCtors modNameStr ty)
            )
            abs.args
            (fArgsAst <> Array.replicate (max 0 (Array.length abs.args - Array.length fArgsAst)) Any)
          goTypes = map snd paramsWithTypes

          localModuleFunctions = Map.insert name { fullName: "Call_local_" <> modNameStr <> "_" <> name, fArgs: goTypes, fRet: TypeValue, arity: Array.length abs.args } moduleFunctions
          declStmts = [ GoRaw ("var Call_local_" <> modNameStr <> "_" <> name <> " func(" <> String.joinWith ", " (map goTypeToStr goTypes) <> ") gopurs_runtime.Value"), GoRaw ("_ = Call_local_" <> modNameStr <> "_" <> name), GoRaw ("var " <> name <> " gopurs_runtime.Value"), GoRaw ("_ = " <> name) ]

          loopBound = foldl (\acc (Tuple idStr goT) -> Map.insert idStr { name: idStr, goType: goT } acc) bound paramsWithTypes
          resBodyMut = translate (context { depth = (depth + 1), moduleFunctions = localModuleFunctions, bound = loopBound, tcoIdent = (Just name), loopCtx = [], options = { isTail: true, inEffectBlock: false }, mbExpectedExprType = Nothing }) (nextId + 1) abs.body

          goParamsNative = map (\(Tuple p goT) -> Tuple (p <> "_loop") goT) paramsWithTypes
          initVars = Array.concatMap (\(Tuple p goT) -> [ GoRaw ("var " <> p <> " " <> goTypeToStr goT <> " = " <> p <> "_loop"), GoRaw ("_ = " <> p) ]) paramsWithTypes
          funcBody = GoBlock (initVars <> flattenStmts resBodyMut.stmts <> [ GoReturn (boxGoExpr codegenStateRef modNameStr resBodyMut.expr resBodyMut.exprType) ])
          nativeAssignment = GoMutate ("Call_local_" <> modNameStr <> "_" <> name) (GoFuncBlock goParamsNative [ funcBody ] TypeValue)

          nativeCallExpr = GoCall (GoVar ("Call_local_" <> modNameStr <> "_" <> name)) (map (\(Tuple p goT) -> coerceGoExpr codegenStateRef modNameStr (GoVar (p <> "_loop_val")) TypeValue goT) paramsWithTypes)
          funcExpr = Array.foldr (\(Tuple p _) acc -> GoCall (GoSelector (GoVar "gopurs_runtime") "Func") [ GoFuncLit [ Tuple (p <> "_loop_val") TypeValue ] [] acc TypeValue ]) nativeCallExpr paramsWithTypes

          newBound = Map.insert originalName { name, goType: TypeValue } bound
          resBodyOuter = translate (context { depth = (depth + 1), moduleFunctions = localModuleFunctions, bound = newBound, tcoIdent = Nothing, mbExpectedExprType = Nothing }) resBodyMut.nextId body
        in
          { stmts: foldMap StmtLeaf declStmts <> StmtLeaf nativeAssignment <> StmtLeaf (GoMutate name funcExpr) <> resBodyOuter.stmts, expr: resBodyOuter.expr, exprType: resBodyOuter.exprType, nextId: resBodyOuter.nextId }

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
              StmtLeaf (GoRaw ("var " <> name <> " " <> goTypeToStr actualGoType <> " = " <> printGoExpr (coerceGoExpr codegenStateRef modNameStr resBinding.expr resBinding.exprType actualGoType)))
        in
          { stmts: resBinding.stmts <> StmtLeaf (GoRaw ("// TAST (Let): " <> name <> " shape=" <> printTcoExprShape binding <> " bindingType=" <> printExprType (getExprType binding))) <> letStmt <> resBody.stmts, expr: resBody.expr, exprType: resBody.exprType, nextId: resBody.nextId }

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
            { newBound: Map.insert oldName { name: newName, goType: expectedGoTypeFromAst } acc.newBound, newNames: Array.snoc acc.newNames { oldName, newName }, exprType: TypeValue, nextId: acc.nextId + 1 }
      )
      { newBound: bound, newNames: [], exprType: TypeValue, nextId }
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
          prepopulatedFunctions = foldl
            ( \accCtx fn ->
                let
                  oldName = localId (Just (Ident fn.ident)) lvl
                  boundInfo = fromMaybe { name: oldName, goType: TypeValue } (Map.lookup oldName allocRes.newBound)
                  newName = boundInfo.name
                  fArgs = case extractExprFuncType (getExprType fn.val) of
                    Just { fArgs: a } -> map (exprTypeToGoType metadata.pointerAdtPaths metadata.enumAdts metadata.elidedCtors modNameStr) a
                    Nothing -> []
                  fRet = case extractExprFuncType (getExprType fn.val) of
                    Just { fRet: r } -> exprTypeToGoType metadata.pointerAdtPaths metadata.enumAdts metadata.elidedCtors modNameStr r
                    Nothing -> TypeValue
                  paramsWithTypes = Array.zipWith (\idStr goT -> Tuple idStr goT) fn.args (fArgs <> Array.replicate (max 0 (Array.length fn.args - Array.length fArgs)) TypeValue)
                in
                  Map.insert newName { fullName: "Call_local_" <> modNameStr <> "_" <> newName, fArgs: map snd paramsWithTypes, fRet: fRet, arity: Array.length fn.args } accCtx
            )
            moduleFunctions
            fns

          prepopulatedBound = foldl
            ( \accCtx fn ->
                let
                  oldName = localId (Just (Ident fn.ident)) lvl
                  boundInfo = fromMaybe { name: oldName, goType: TypeValue } (Map.lookup oldName allocRes.newBound)
                  newName = boundInfo.name
                  fArgs = case extractExprFuncType (getExprType fn.val) of
                    Just { fArgs: a } -> map (exprTypeToGoType metadata.pointerAdtPaths metadata.enumAdts metadata.elidedCtors modNameStr) a
                    Nothing -> []
                  fRet = case extractExprFuncType (getExprType fn.val) of
                    Just { fRet: r } -> exprTypeToGoType metadata.pointerAdtPaths metadata.enumAdts metadata.elidedCtors modNameStr r
                    Nothing -> TypeValue
                  paramsWithTypes = Array.zipWith (\idStr goT -> Tuple idStr goT) fn.args (fArgs <> Array.replicate (max 0 (Array.length fn.args - Array.length fArgs)) TypeValue)
                in
                  Map.insert oldName { name: newName, goType: TypeFunc (map snd paramsWithTypes) fRet } accCtx
            )
            allocRes.newBound
            fns

          resData = foldl
            ( \acc fn ->
                let
                  oldName = localId (Just (Ident fn.ident)) lvl
                  boundInfo = fromMaybe { name: oldName, goType: TypeValue } (Map.lookup oldName prepopulatedBound)
                  newName = boundInfo.name
                  fArgs = case extractExprFuncType (getExprType fn.val) of
                    Just { fArgs: a } -> map (exprTypeToGoType metadata.pointerAdtPaths metadata.enumAdts metadata.elidedCtors modNameStr) a
                    Nothing -> []
                  paramsWithTypes = Array.zipWith (\idStr goT -> Tuple idStr goT) fn.args (fArgs <> Array.replicate (max 0 (Array.length fn.args - Array.length fArgs)) TypeValue)
                  currentLoopCtx = [ { ident: newName, params: fn.args, loopParams: map (\p -> p <> "_loop") fn.args, goTypes: map snd paramsWithTypes, fRet: TypeValue } ]
                  loopBound = foldl (\acc2 (Tuple idStr goT) -> Map.insert idStr { name: idStr, goType: goT } acc2) prepopulatedBound paramsWithTypes
                  mbExpectedRet = case extractExprFuncType (getExprType fn.val) of
                    Just { fRet: r } -> Just r
                    Nothing -> Nothing
                  resBodyMut = translate (context { depth = (depth + 1), recVars = combinedRecVars, moduleFunctions = acc.moduleFunctions, bound = loopBound, tcoIdent = (Just newName), loopCtx = currentLoopCtx, options = { isTail: true, inEffectBlock: false }, mbExpectedExprType = mbExpectedRet }) acc.nextId fn.body
                  trueFRet = resBodyMut.exprType

                  initVars = Array.concatMap (\(Tuple p goT) -> [ GoRaw ("var " <> p <> " " <> goTypeToStr goT <> " = " <> p <> "_loop"), GoRaw ("_ = " <> p) ]) paramsWithTypes

                  funcBody = GoFor newName (initVars <> flattenStmts resBodyMut.stmts <> [ GoReturn resBodyMut.expr ])

                  goParamsNative = map (\(Tuple p goT) -> Tuple (p <> "_loop") goT) paramsWithTypes
                  nativeAssignment = GoMutate ("Call_local_" <> modNameStr <> "_" <> newName) (GoFuncBlock goParamsNative [ funcBody ] trueFRet)

                  nativeCallExpr = GoCall (GoVar ("Call_local_" <> modNameStr <> "_" <> newName)) (map (\(Tuple p goT) -> coerceGoExpr codegenStateRef modNameStr (GoVar (p <> "_loop_val")) TypeValue goT) paramsWithTypes)
                  funcExpr =
                    if Array.null paramsWithTypes then
                      GoFunc "_" TypeValue TypeValue (GoBlock [ GoReturn (boxGoExpr codegenStateRef modNameStr nativeCallExpr trueFRet) ])
                    else
                      Array.foldr (\(Tuple p _) accExpr -> GoCall (GoSelector (GoVar "gopurs_runtime") "Func") [ GoFuncLit [ Tuple (p <> "_loop_val") TypeValue ] [] accExpr TypeValue ]) (boxGoExpr codegenStateRef modNameStr nativeCallExpr trueFRet) paramsWithTypes

                  newFunctions = Map.insert newName { fullName: "Call_local_" <> modNameStr <> "_" <> newName, fArgs: map snd paramsWithTypes, fRet: trueFRet, arity: Array.length fn.args } acc.moduleFunctions
                  newBound2 = Map.insert oldName { name: newName, goType: TypeFunc (map snd paramsWithTypes) trueFRet } acc.newBound
                  declStmtsLocal = [ GoRaw ("var Call_local_" <> modNameStr <> "_" <> newName <> " func(" <> String.joinWith ", " (map goTypeToStr (map snd paramsWithTypes)) <> ") " <> goTypeToStr trueFRet), GoRaw ("_ = Call_local_" <> modNameStr <> "_" <> newName), GoRaw ("var " <> newName <> " gopurs_runtime.Value"), GoRaw ("_ = " <> newName) ]
                in
                  { declarations: acc.declarations <> declStmtsLocal, stmts: acc.stmts <> [ nativeAssignment, GoMutate newName funcExpr ], nextId: resBodyMut.nextId, moduleFunctions: newFunctions, newBound: newBound2 }
            )
            { declarations: [], stmts: [], nextId: allocRes.nextId, moduleFunctions: prepopulatedFunctions, newBound: prepopulatedBound }
            fns

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
                      <> StmtLeaf (GoMutate (alloc.newName <> "_cell") (GoRaw ("&" <> alloc.newName)))
                  , exprs: Array.snoc acc.exprs { key: alloc.newName, goType: expectedGoType }
                  , exprType: TypeValue
                  , nextId: res.nextId
                  }
            )
            { stmts: StmtEmpty, exprs: [], exprType: TypeValue, nextId: allocRes.nextId }
            (Array.zip (toArray bindings) allocRes.newNames)

          declStmts = map (\b -> GoRaw ("var " <> b.key <> " " <> goTypeToStr b.goType <> "\n_ = " <> b.key
            <> "\nvar " <> b.key <> "_cell *" <> goTypeToStr b.goType <> "\n_ = " <> b.key <> "_cell"
            <> "\n// FALLBACK TCO: isLoop=" <> show isLoop <> " len=" <> show (Array.length (toArray bindings)))) accBindings.exprs

          resBody = translate (context { depth = (depth + 1), recVars = combinedRecVars, bound = allocRes.newBound, tcoIdent = Nothing }) accBindings.nextId body
        in
          { stmts: foldMap StmtLeaf declStmts <> accBindings.stmts <> resBody.stmts, expr: resBody.expr, exprType: resBody.exprType, nextId: resBody.nextId }
