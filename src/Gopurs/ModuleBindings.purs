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
import Data.Traversable (traverse)
import Data.Tuple (Tuple(..))
import Effect.Ref (Ref)
import Gopurs.CallAnalysis (extractUncurriedAbs)
import Gopurs.CodegenState (CodegenMetadata, CodegenState, FunctionInfo)
import Gopurs.ExprAnalysis (extractFuncType)
import Gopurs.ExprContext (ModuleFunctions, TranslateExpr, wrapInStmts)
import Gopurs.GoAst (GoDecl(..), GoType(..), sanitizeName)
import Gopurs.GoConversions (adtPayloadTypes, boxGoExpr, getUnboxedADT)
import Gopurs.GoTypes (exprTypeToGoType)
import Gopurs.ModuleWorkers as ModuleWorkers
import Gopurs.NativeRecordArgs (workerArguments)
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
prepare :: CodegenMetadata -> String -> BackendModule -> { bindings :: Array TcoBindingGroup, functions :: ModuleFunctions }
prepare metadata modNameStr mod =
  let
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

    moduleFunctions :: ModuleFunctions
    moduleFunctions = Map.fromFoldable $ Array.concatMap
      ( \group ->
          if group.recursive then
            let
              mutRecBinds = traverse (\(Tuple _ val) -> extractUncurriedAbs val) group.bindings
            in
              case mutRecBinds of
                Just _ -> map (bindingSignature metadata modNameStr) group.bindings
                Nothing -> map (\(Tuple (Ident name) _) -> Tuple (sanitizeName name) (boxedSignature modNameStr name)) group.bindings
          else
            map (bindingSignature metadata modNameStr) group.bindings
      )
      tcoBindings
  in
    { bindings: tcoBindings, functions: moduleFunctions }

boxedSignature :: String -> String -> FunctionInfo
boxedSignature modNameStr name =
  { fullName: "Call_" <> modNameStr <> "_" <> sanitizeName name, fArgs: [], fRet: TypeValue, arity: 0 }

-- Syntax fixes worker arity. A computation between lambdas leaves a boxed
-- function result; fully consumed native sums use their specialized ABI.
-- Read-only record projections are selected here and reused by ModuleWorkers.
bindingSignature :: CodegenMetadata -> String -> Tuple Ident TcoExpr -> Tuple String FunctionInfo
bindingSignature metadata modNameStr (Tuple (Ident name) value) =
  let
    fallback = boxedSignature modNameStr name
    toGoType = exprTypeToGoType metadata.pointerAdtPaths metadata.enumAdts metadata.elidedCtors modNameStr
    signature = case extractUncurriedAbs value of
      Just { args, body } ->
        let
          annotation = extractFuncType value
          parameters = case annotation of
            Just { fArgs, fRet } -> workerArguments toGoType args body
              (Array.take (Array.length args) fArgs)
              (if Array.length args < Array.length fArgs then Any else fRet)
              <> Array.replicate (Array.length args - Array.length fArgs) TypeValue
            Nothing -> Array.replicate (Array.length args) TypeValue
          result = case annotation of
            Just { fArgs, fRet } | Array.length args >= Array.length fArgs ->
              case getUnboxedADT fRet of
                Just (Tuple adtName adt) -> TypeStructValue adtName (adt.signature (adtPayloadTypes toGoType fRet))
                Nothing -> toGoType fRet
            _ -> TypeValue
        in
          fallback { fArgs = parameters, fRet = result, arity = Array.length args }
      Nothing -> fallback
  in
    Tuple (sanitizeName name) signature

-- Declarations and loop bodies share the same child translator as local
-- bindings. Each recursive group publishes its function declarations together.
declarations :: TranslateExpr -> CodegenMetadata -> Ref CodegenState -> String -> ModuleFunctions -> Array TcoBindingGroup -> Array GoDecl
declarations translate metadata codegenStateRef modNameStr moduleFunctions groups =
  Array.concatMap
    ( \group ->
        let
          recVars = if group.recursive then map (\(Tuple (Ident name) _) -> sanitizeName name) group.bindings else []

          context = { metadata, codegenStateRef, depth: 0, modNameStr, recVars, moduleFunctions, bound: Map.empty, tcoIdent: Nothing, loopCtx: [], options: { isTail: false, inEffectBlock: false }, mbExpectedExprType: Nothing }

          processBindingGroup :: Array (Tuple Ident TcoExpr) -> Array GoDecl
          processBindingGroup binds =
            let
              mutRecBinds = traverse (\(Tuple (Ident name) val) -> map (\abs -> { ident: sanitizeName name, args: abs.args, body: abs.body, val: val }) (extractUncurriedAbs val)) binds
            in
              case mutRecBinds of
                Just fns ->
                  map (ModuleWorkers.declaration translate context (group.recursive && Array.length group.bindings == 1)) fns
                Nothing ->
                  Array.concatMap
                    ( \(Tuple (Ident name) expr) ->
                        let
                          res = translate (context { depth = 0, bound = Map.empty, tcoIdent = (Just (sanitizeName name)), loopCtx = [], options = { isTail: false, inEffectBlock: false }, mbExpectedExprType = (Just (annotationType expr)) }) 0 expr
                        in
                          [ GoCachedValue { identifier: modNameStr <> "_" <> sanitizeName name, expression: wrapInStmts [] res.stmts TypeValue (boxGoExpr codegenStateRef modNameStr res.expr res.exprType), goType: TypeValue } ]
                    )
                    binds
        in
          if group.recursive then
            processBindingGroup group.bindings
          else
            Array.concatMap (\b -> processBindingGroup [ b ]) group.bindings
    )
    groups

-- Cached values use only an outer annotation, not primitive-result inference.
annotationType :: TcoExpr -> ExprType
annotationType (TcoExpr _ (Typed ty _)) = ty
annotationType _ = Any
