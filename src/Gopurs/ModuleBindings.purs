module Gopurs.ModuleBindings
  ( TcoBindingGroup
  , prepare
  , declarations
  ) where

import Prelude
import Data.Array as Array
import Data.Array.NonEmpty (fromArray)
import Data.Foldable (foldl)
import Data.Map as Map
import Data.Maybe (Maybe(..))
import Data.String as String
import Data.Traversable (traverse)
import Data.Tuple (Tuple(..), fst, snd)
import Effect.Ref (Ref)
import Effect.Ref as Ref
import Effect.Unsafe (unsafePerformEffect)
import Gopurs.CallAnalysis (extractUncurriedAbs)
import Gopurs.CodegenState (CodegenState, FunctionInfo)
import Gopurs.ExprAnalysis (extractExprFuncType, extractFuncType)
import Gopurs.ExprContext (LoopContext, ModuleFunctions, TranslateExpr, flattenStmts, wrapInStmts)
import Gopurs.GoAst (GoDecl, GoExpr(..), GoType(..), goTypeToStr, sanitizeName)
import Gopurs.GoConversions (boxGoExpr, coerceGoExpr, getUnboxedADT)
import Gopurs.GoTypes (exprTypeToGoType)
import Gopurs.Printer (printGoExpr)
import PureScript.Backend.Optimizer.Codegen.Tco (TcoExpr(..))
import PureScript.Backend.Optimizer.Codegen.Tco as Tco
import PureScript.Backend.Optimizer.Convert (BackendModule)
import PureScript.Backend.Optimizer.CoreFn (ExprType(..), Ident(..))
import PureScript.Backend.Optimizer.Syntax (BackendSyntax(Typed))

type TcoBindingGroup =
  { recursive :: Boolean
  , bindings :: Array (Tuple Ident TcoExpr)
  }

-- Analyze recursion before publishing the native signatures used by calls.
prepare :: Ref CodegenState -> String -> BackendModule -> { bindings :: Array TcoBindingGroup, functions :: ModuleFunctions }
prepare codegenStateRef modNameStr mod =
  let
    { pointerAdtPaths, enumAdts, elidedCtors } = unsafePerformEffect (Ref.read codegenStateRef)
    Tuple _ tcoBindings = foldl
      ( \(Tuple env acc) group ->
          let
            neBindings = fromArray group.bindings

            env' = case neBindings of
              Just ne | group.recursive -> Tco.topLevelTcoEnvGroup mod.name ne <> env
              _ -> env
            tcoBinds = map
              (\(Tuple id val) -> Tuple id (Tco.analyze env' val))
              group.bindings
          in
            Tuple env' (Array.snoc acc { recursive: group.recursive, bindings: tcoBinds })
      )
      (Tuple [] [])
      mod.bindings

    unwrapFunc :: Array (Tuple Ident TcoExpr) -> Array (Tuple String FunctionInfo)
    unwrapFunc binds =
      Array.concatMap
        ( \(Tuple (Ident name) val) ->
            case extractUncurriedAbs val of
              Just { args } ->
                let
                  typeSig = extractFuncType val
                  fArgsGo = case typeSig of
                    Just { fArgs } -> map (exprTypeToGoType pointerAdtPaths enumAdts elidedCtors modNameStr) (Array.take (Array.length args) fArgs) <> Array.replicate (Array.length args - Array.length fArgs) TypeValue
                    Nothing -> Array.replicate (Array.length args) TypeValue
                  fRetGo = case typeSig of
                    Just { fArgs, fRet } ->
                      if Array.length args < Array.length fArgs then TypeValue
                      else
                        case getUnboxedADT fRet of
                          Just (Tuple adtName adt) -> TypeStructValue adtName adt.signature
                          Nothing -> exprTypeToGoType pointerAdtPaths enumAdts elidedCtors modNameStr fRet
                    Nothing -> TypeValue
                  fullName = "Call_" <> modNameStr <> "_" <> sanitizeName name
                in
                  [ Tuple (sanitizeName name) { fullName, fArgs: fArgsGo, fRet: fRetGo, arity: Array.length args } ]
              Nothing ->
                let
                  fullName = "Call_" <> modNameStr <> "_" <> sanitizeName name
                in
                  [ Tuple (sanitizeName name) { fullName, fArgs: [], fRet: TypeValue, arity: 0 } ]
        )
        binds

    moduleFunctions :: ModuleFunctions
    moduleFunctions = Map.fromFoldable $ Array.concatMap
      ( \group ->
          if group.recursive then
            let
              mutRecBinds = traverse (\(Tuple _ val) -> extractUncurriedAbs val) group.bindings
            in
              case mutRecBinds of
                Just _ -> unwrapFunc group.bindings
                Nothing -> map (\(Tuple (Ident name) _) -> Tuple (sanitizeName name) { fullName: "Call_" <> modNameStr <> "_" <> sanitizeName name, fArgs: [], fRet: TypeValue, arity: 0 }) group.bindings
          else
            unwrapFunc group.bindings
      )
      tcoBindings
  in
    { bindings: tcoBindings, functions: moduleFunctions }


-- Declarations and loop bodies share the same child translator as local
-- bindings. Each recursive group publishes its function declarations together.
declarations :: TranslateExpr -> Ref CodegenState -> String -> ModuleFunctions -> Array TcoBindingGroup -> Array GoDecl
declarations translate codegenStateRef modNameStr moduleFunctions groups =
  Array.concatMap
    ( \group ->
        let
          recVars = if group.recursive then map (\(Tuple (Ident name) _) -> sanitizeName name) group.bindings else []

          context = { codegenStateRef, depth: 0, modNameStr, recVars, moduleFunctions, bound: Map.empty, tcoIdent: Nothing, loopCtx: [], options: { isTail: false, inEffectBlock: false }, mbExpectedExprType: Nothing }

          processBindingGroup :: Array (Tuple Ident TcoExpr) -> Array GoDecl
          processBindingGroup binds =
            let
              mutRecBinds = traverse (\(Tuple (Ident name) val) -> map (\abs -> { ident: sanitizeName name, args: abs.args, body: abs.body, val: val }) (extractUncurriedAbs val)) binds
            in
              case mutRecBinds of
                Just fns ->
                  let
                    fnWrapperStmts = map
                      ( \fn ->
                          let
                            paramsWithTypes = case extractExprFuncType (getExprType fn.val) of
                              Just { fArgs } -> Array.zipWith (\p goType -> Tuple p goType) fn.args (map (exprTypeToGoType (unsafePerformEffect (Ref.read codegenStateRef)).pointerAdtPaths (unsafePerformEffect (Ref.read codegenStateRef)).enumAdts (unsafePerformEffect (Ref.read codegenStateRef)).elidedCtors modNameStr) fArgs <> Array.replicate (Array.length fn.args - Array.length fArgs) TypeValue)
                              Nothing -> map (\p -> Tuple p TypeValue) fn.args

                            newBound = foldl (\acc (Tuple idStr goType) -> Map.insert idStr { name: idStr, goType } acc) Map.empty paramsWithTypes

                            isSelfRecursiveLoop = group.recursive && Array.length group.bindings == 1
                            mbExpectedRet = case extractExprFuncType (getExprType fn.val) of
                              Just { fArgs, fRet: rt } -> Just case Array.drop (Array.length fn.args) fArgs of
                                [] -> rt
                                remaining -> Func remaining rt
                              Nothing -> Nothing
                            fRet = case mbExpectedRet of
                              Just rt -> exprTypeToGoType (unsafePerformEffect (Ref.read codegenStateRef)).pointerAdtPaths (unsafePerformEffect (Ref.read codegenStateRef)).enumAdts (unsafePerformEffect (Ref.read codegenStateRef)).elidedCtors modNameStr rt
                              Nothing -> TypeValue

                            currentLoopCtx :: LoopContext
                            currentLoopCtx = if isSelfRecursiveLoop then [ { ident: fn.ident, params: map fst paramsWithTypes, loopParams: map (\p -> fst p <> "_loop") paramsWithTypes, goTypes: map snd paramsWithTypes, fRet } ] else []
                            resBodyMut = translate (context { depth = 0, bound = newBound, tcoIdent = (Just fn.ident), loopCtx = currentLoopCtx, options = { isTail: isSelfRecursiveLoop, inEffectBlock: false }, mbExpectedExprType = mbExpectedRet }) 0 fn.body

                            goName = fn.ident
                            initVars = Array.concatMap (\(Tuple p goT) -> [ GoRaw ("var " <> p <> " " <> goTypeToStr goT <> " = " <> p <> "_loop"), GoRaw ("_ = " <> p) ]) paramsWithTypes

                            arity = Array.length fn.args

                            expectedRetType = case Map.lookup goName moduleFunctions of
                              Just { fArgs } | arity < Array.length fArgs -> TypeValue
                              Just { fRet: resultType } -> resultType
                              Nothing -> TypeValue

                            goParams = String.joinWith ", " (map (\(Tuple p goT) -> p <> "_loop " <> goTypeToStr goT) paramsWithTypes)

                            funcExpr =
                              if arity >= 1 && arity <= 10 then
                                let
                                  coercedExpr = coerceGoExpr codegenStateRef modNameStr resBodyMut.expr resBodyMut.exprType expectedRetType
                                  bodyStmts = initVars <> flattenStmts resBodyMut.stmts <> [ GoReturn coercedExpr ]
                                  funcBody = if isSelfRecursiveLoop then GoFor goName bodyStmts else GoBlock bodyStmts
                                in
                                  unsafePerformEffect do
                                    let callFuncDecl = "func Call_" <> modNameStr <> "_" <> goName <> "(" <> goParams <> ") " <> goTypeToStr expectedRetType <> " {\n" <> printGoExpr funcBody <> "\n}"
                                    Ref.modify_ (\r -> r { rawDecls = Array.snoc r.rawDecls callFuncDecl }) codegenStateRef
                                    let wrapperParams = map (\(Tuple p _) -> p <> "_box") paramsWithTypes
                                    let callExpr = GoCall (GoVar ("Call_" <> modNameStr <> "_" <> goName)) (map (\(Tuple p goT) -> coerceGoExpr codegenStateRef modNameStr (GoVar (p <> "_box")) TypeValue goT) paramsWithTypes)
                                    let boxedRes = boxGoExpr codegenStateRef modNameStr callExpr expectedRetType
                                    let wrapperFunc = GoRaw ("func(" <> String.joinWith ", " (map (\p -> p <> " gopurs_runtime.Value") wrapperParams) <> ") gopurs_runtime.Value {\nreturn " <> printGoExpr boxedRes <> "\n}")
                                    let funcWrapperName = if arity == 1 then "gopurs_runtime.Func" else "gopurs_runtime.Func" <> show arity
                                    pure $ GoRaw (funcWrapperName <> "(" <> printGoExpr wrapperFunc <> ")")
                              else
                                let
                                  bodyStmts = initVars <> flattenStmts resBodyMut.stmts <> [ GoReturn (boxGoExpr codegenStateRef modNameStr resBodyMut.expr resBodyMut.exprType) ]
                                  funcBody = if isSelfRecursiveLoop then GoFor goName bodyStmts else GoBlock bodyStmts
                                  iife = GoRaw ("func() gopurs_runtime.Value {\n" <> printGoExpr funcBody <> "\n}()")
                                in
                                  if arity == 0 then
                                    GoFunc "_" TypeValue TypeValue funcBody
                                  else
                                    Array.foldr (\(Tuple p goT) acc -> GoCall (GoSelector (GoVar "gopurs_runtime") "Func") [ GoRaw ("func(" <> p <> "_box gopurs_runtime.Value) gopurs_runtime.Value {\nvar " <> p <> "_loop " <> goTypeToStr goT <> " = " <> printGoExpr (coerceGoExpr codegenStateRef modNameStr (GoVar (p <> "_box")) TypeValue goT) <> "\nreturn " <> printGoExpr acc <> "\n}") ]) iife paramsWithTypes
                          in
                            { identifier: modNameStr <> "_" <> goName, expression: funcExpr, goType: TypeValue }
                      )
                      fns
                  in
                    fnWrapperStmts
                Nothing ->
                  Array.concatMap
                    ( \(Tuple (Ident name) expr) ->
                        let
                          res = translate (context { depth = 0, bound = Map.empty, tcoIdent = (Just (sanitizeName name)), loopCtx = [], options = { isTail: false, inEffectBlock: false }, mbExpectedExprType = (Just (getExprType expr)) }) 0 expr
                        in
                          [ { identifier: modNameStr <> "_" <> sanitizeName name, expression: wrapInStmts [] res.stmts TypeValue (boxGoExpr codegenStateRef modNameStr res.expr res.exprType), goType: TypeValue } ]
                    )
                    binds
        in
          if group.recursive then
            processBindingGroup group.bindings
          else
            Array.concatMap (\b -> processBindingGroup [ b ]) group.bindings
    )
    groups

getExprType :: TcoExpr -> ExprType
getExprType (TcoExpr _ (Typed ty _)) = ty
getExprType _ = Any -- fallback
