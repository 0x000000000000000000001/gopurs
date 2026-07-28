module Gopurs.FBIP (extractProjections) where

import Prelude

import Control.Monad.State (State, evalState, runState, get, put)
import Data.Array as Array
import Data.Maybe (Maybe(..))
import Data.Traversable (traverse)
import Data.Tuple (Tuple(..))
import PureScript.Backend.Optimizer.Codegen.Tco (TcoExpr(..))
import PureScript.Backend.Optimizer.CoreFn (Ident(..), Qualified(..))
import PureScript.Backend.Optimizer.Syntax (BackendSyntax(..), Level(..), Pair(..))

type FBIPState = { nextId :: Int, bindings :: Array { ident :: String, expr :: TcoExpr } }

extractProjections :: TcoExpr -> TcoExpr
extractProjections expr = evalState (processExpr expr) 0

processExpr :: TcoExpr -> State Int TcoExpr
processExpr (TcoExpr a syn) = case syn of
  Branch branches def -> do
    newBranches <- traverse (\(Pair cond body) -> do
      cond' <- processExpr cond
      body' <- processBranchBody body
      pure (Pair cond' body')
    ) branches
    def' <- processBranchBody def
    pure (TcoExpr a (Branch newBranches def'))
  
  _ -> do
    syn' <- traverse processExpr syn
    pure (TcoExpr a syn')

processBranchBody :: TcoExpr -> State Int TcoExpr
processBranchBody body = do
  startId <- get
  let Tuple expr finalState = runState (extractAccessorsTraverser body) { nextId: startId, bindings: [] }
  put finalState.nextId
  let wrapped = Array.foldr (\b acc -> TcoExpr mempty (Let (Just (Ident b.ident)) (Level 0) b.expr acc)) expr finalState.bindings
  -- recursively process the new body which might contain nested Branches
  processExpr wrapped

extractAccessorsTraverser :: TcoExpr -> State FBIPState TcoExpr
extractAccessorsTraverser tco@(TcoExpr a syn) = case syn of
  Accessor obj accessor -> do
    obj' <- extractAccessorsTraverser obj
    st <- get
    let newId = st.nextId
    let varName = "__fbip_proj_" <> show newId
    let newBinding = { ident: varName, expr: TcoExpr a (Accessor obj' accessor) }
    put { nextId: newId + 1, bindings: Array.snoc st.bindings newBinding }
    pure (TcoExpr mempty (Local (Just (Ident varName)) (Level 0)))
    
  Branch _ _ -> pure tco
  
  _ -> do
    syn' <- traverse extractAccessorsTraverser syn
    pure (TcoExpr a syn')
