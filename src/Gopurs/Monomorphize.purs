module Gopurs.Monomorphize where

import Prelude

import Data.Map (Map)
import Data.Map as Map
import Data.Set (Set)
import Data.Set as Set
import Data.Foldable (foldl)
import Data.Maybe (Maybe(..), fromMaybe)
import PureScript.Backend.Optimizer.CoreFn (Ann(..), Bind(..), Binding(..), CaseAlternative(..), CaseGuard(..), Expr(..), ExprType(..), Guard(..), Module(..), Prop(..), Ident(..), Qualified(..), unQualified)
import PureScript.Backend.Optimizer.Substitute (unify, substituteExprType)
import Data.Tuple (Tuple(..))
import Data.Array as Array
import Data.String as String
import PureScript.Backend.Optimizer.Monomorphize (mangleType, getExprAnn)
import Data.Newtype (unwrap)
import Debug (traceM)

type WorkItem = Tuple (Qualified Ident) ExprType

type Env =
  { globals :: Map (Qualified Ident) { type :: ExprType, expr :: Expr Ann }
  , processed :: Set String
  , instantiations :: Map (Qualified Ident) (Set ExprType)
  , adtInstantiations :: Set ExprType
  }

formatWorkItem :: WorkItem -> String
formatWorkItem (Tuple qIdent ty) =
  let
    modStr = case qIdent of
      Qualified (Just m) _ -> unwrap m <> "."
      _ -> ""
    identStr = case qIdent of
      Qualified _ (Ident i) -> i
  in modStr <> identStr <> "___" <> mangleType ty

buildGlobals :: Array (Module Ann) -> Map (Qualified Ident) { type :: ExprType, expr :: Expr Ann }
buildGlobals modules = foldl processModule Map.empty modules
  where
  processModule acc (Module m) =
    let modName = m.name
        processBind a (NonRec binding) = processBinding modName a binding
        processBind a (Rec bindings) = foldl (processBinding modName) a bindings
    in foldl processBind acc m.decls

  processBinding modName acc (Binding _ ident expr) =
    let qIdent = Qualified (Just modName) ident
    in case getExprAnn expr of
      Ann ann -> case ann.type of
        Just t -> Map.insert qIdent { type: t, expr } acc
        Nothing -> acc

processQueue :: Env -> Array WorkItem -> Env
processQueue env [] = env
processQueue env queue =
  let
    Tuple env' nextQueue = foldl processItem (Tuple env []) queue
  in processQueue env' nextQueue

processItem :: Tuple Env (Array WorkItem) -> WorkItem -> Tuple Env (Array WorkItem)
processItem (Tuple env nextQueue) item@(Tuple qIdent instType) =
  let itemId = formatWorkItem item
  in if Set.member itemId env.processed then
       Tuple env nextQueue
     else
       let
         env1 = env { processed = Set.insert itemId env.processed
                    , instantiations = Map.insertWith Set.union qIdent (Set.singleton instType) env.instantiations }
       in case Map.lookup qIdent env.globals of
         Nothing -> Tuple env1 nextQueue -- Foreign or missing
         Just def ->
           let
             subst = unify def.type instType Map.empty
             Tuple env2 newItems = collectExpr env1 subst def.expr
           in Tuple env2 (nextQueue <> newItems)

collectExpr :: Env -> Map String ExprType -> Expr Ann -> Tuple Env (Array WorkItem)
collectExpr env subst expr = case expr of
  ExprVar (Ann ann) qIdent ->
    case ann.type of
      Just t ->
        let concreteType = substituteExprType subst t
        in Tuple env [Tuple qIdent concreteType]
      Nothing -> Tuple env []
  ExprLit _ lit -> foldl (\(Tuple e q) ex -> let Tuple e' q' = collectExpr e subst ex in Tuple e' (q <> q')) (Tuple env []) lit
  ExprApp _ f arg ->
    let Tuple env1 q1 = collectExpr env subst f
        Tuple env2 q2 = collectExpr env1 subst arg
    in Tuple env2 (q1 <> q2)
  ExprAbs _ _ e -> collectExpr env subst e
  ExprLet _ binds e ->
    let Tuple env1 q1 = foldl (\(Tuple en qu) b ->
          case b of
            NonRec (Binding _ _ ex) -> let Tuple e' q' = collectExpr en subst ex in Tuple e' (qu <> q')
            Rec bindings -> foldl (\(Tuple en' qu') (Binding _ _ ex) -> let Tuple e'' q'' = collectExpr en' subst ex in Tuple e'' (qu' <> q'')) (Tuple en qu) bindings
        ) (Tuple env []) binds
        Tuple env2 q2 = collectExpr env1 subst e
    in Tuple env2 (q1 <> q2)
  ExprCase _ exprs alts ->
    let Tuple env1 q1 = foldl (\(Tuple en qu) ex -> let Tuple e' q' = collectExpr en subst ex in Tuple e' (qu <> q')) (Tuple env []) exprs
        Tuple env2 q2 = foldl (\(Tuple en qu) alt ->
          case alt of
            CaseAlternative _ cg -> case cg of
              Unconditional ex -> let Tuple e' q' = collectExpr en subst ex in Tuple e' (qu <> q')
              Guarded guards -> foldl (\(Tuple en' qu') (Guard e1 e2) ->
                  let Tuple ea qa = collectExpr en' subst e1
                      Tuple eb qb = collectExpr ea subst e2
                  in Tuple eb (qu' <> qa <> qb)
                ) (Tuple en qu) guards
        ) (Tuple env1 []) alts
    in Tuple env2 (q1 <> q2)
  ExprConstructor (Ann ann) _ _ _ ->
    case ann.type of
      Just t ->
        let concreteType = substituteExprType subst t
        in Tuple (env { adtInstantiations = Set.insert concreteType env.adtInstantiations }) []
      Nothing -> Tuple env []
  ExprAccessor _ e _ -> collectExpr env subst e
  ExprUpdate _ e props ->
    let Tuple env1 q1 = collectExpr env subst e
        Tuple env2 q2 = foldl (\(Tuple en qu) (Prop _ ex) -> let Tuple e' q' = collectExpr en subst ex in Tuple e' (qu <> q')) (Tuple env1 []) props
    in Tuple env2 (q1 <> q2)

analyzeReachability :: Array (Module Ann) -> Env
analyzeReachability modules =
  let
    globals = buildGlobals modules
    
    hasTypeVars :: ExprType -> Boolean
    hasTypeVars (TypeVar v) = String.take 1 v == String.toLower (String.take 1 v) && v /= "gopurs_runtime.Value"
    hasTypeVars (Array t) = hasTypeVars t
    hasTypeVars (Func args ret) = Array.any hasTypeVars args || hasTypeVars ret
    hasTypeVars (Record props) = Array.any (\(Tuple _ t) -> hasTypeVars t) props
    hasTypeVars (ADT _ args) = Array.any hasTypeVars args
    hasTypeVars _ = false
    
    entryPoints = Array.mapMaybe (\(Tuple qIdent def) ->
      if not (hasTypeVars def.type) then
        Just (Tuple qIdent def.type)
      else
        Nothing
    ) (Map.toUnfoldable globals :: Array (Tuple (Qualified Ident) { type :: ExprType, expr :: Expr Ann }))
    
    env = { globals, processed: Set.empty, instantiations: Map.empty, adtInstantiations: Set.empty }
  in processQueue env entryPoints

