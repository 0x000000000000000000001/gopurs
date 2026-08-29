module Gopurs.OptimizeTAST where

import Prelude
import Data.Maybe (Maybe(..))
import Data.Map as Map
import Data.Set as Set
import Data.Tuple (Tuple(..))
import Data.Array as Array
import Partial.Unsafe (unsafePartial)

import PureScript.Backend.Optimizer.CoreFn (Ann(..), Bind(..), Binding(..), Module(..))
import PureScript.Backend.Optimizer.Semantics (BackendExpr, Ctx(..), NeutralExpr(..), build)
import PureScript.Backend.Optimizer.Convert (toBackendModule)
import Gopurs.QuoteTAST (quoteTAST)

optimizeModuleTAST :: Module Ann -> Module Ann
optimizeModuleTAST coreFnMod@(Module m) =
  let
    -- 1. Convert to BackendModule (runs Semantics.eval implicitly)
    -- We ignore the OptimizationSteps returned as the first element of the Tuple.
    backendMod = case toBackendModule coreFnMod
      { analyzeCustom: \_ _ -> Nothing
      , currentModule: m.name
      , currentLevel: 0
      , toLevel: Map.empty
      , implementations: Map.empty
      , moduleImplementations: Map.empty
      , directives: Map.empty
      , dataTypes: Map.empty
      , foreignSemantics: Map.empty
      , rewriteLimit: 10000
      , traceIdents: Set.empty
      , optimizationSteps: []
      } of
      Tuple _ bMod -> bMod

    -- 2. Build a basic context for quoting
    defaultCtx = Ctx
      { currentLevel: 0
      , lookupExtern: \_ _ -> Nothing
      , analyze: \_ _ -> mempty
      , effect: false
      }

    dummyAnn = Ann { span: { path: "", start: { line: 0, column: 0 }, end: { line: 0, column: 0 } }, meta: Nothing, type: Nothing }

    unfreeze :: Ctx -> NeutralExpr -> BackendExpr
    unfreeze ctx (NeutralExpr syn) = build ctx (map (unfreeze ctx) syn)

    -- 3. Map over bindings and quote them back to Expr Ann
    newBinds = map (\group ->
      if group.recursive then
        Rec (map (\(Tuple ident neutSem) ->
          Binding dummyAnn ident (quoteTAST (unfreeze defaultCtx neutSem))
        ) group.bindings)
      else
        unsafePartial $
          let Just (Tuple ident neutSem) = Array.head group.bindings
          in NonRec (Binding dummyAnn ident (quoteTAST (unfreeze defaultCtx neutSem)))
    ) backendMod.bindings

  in
    Module (m { decls = newBinds })
