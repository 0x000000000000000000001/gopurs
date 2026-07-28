module Gopurs.Monomorphize.Substitute where

import Prelude

import Data.Map (Map)
import Data.Map as Map
import Data.Maybe (Maybe(..), fromMaybe)
import Data.Tuple (Tuple(..))
import Data.Array as Array
import Data.Set (Set)
import Data.Set as Set
import PureScript.Backend.Optimizer.Codegen.Tco (TcoExpr(..))
import PureScript.Backend.Optimizer.CoreFn (Ident(..), ExprType(..), ModuleName)
import PureScript.Backend.Optimizer.Convert (BackendBindingGroup)
import PureScript.Backend.Optimizer.Semantics (NeutralExpr)
import PureScript.Backend.Optimizer.Codegen.Tco as Tco
import PureScript.Backend.Optimizer.CoreFn (ModuleName)
import PureScript.Backend.Optimizer.CoreFn (ExprType(..))
import PureScript.Backend.Optimizer.Substitute (unify, substituteExprType, mapTcoExprTypes, substituteAst, setTcoExprType)
import PureScript.Backend.Optimizer.Syntax (BackendSyntax(..))
import Effect.Console as Console
import Effect.Unsafe (unsafePerformEffect)
import Data.Array.NonEmpty (fromArray)
getExprType (TcoExpr _ (Typed ty _)) = ty
getExprType _ = Any

specializeBindingGroups :: ModuleName -> Map String String -> String -> Map String ExprType -> Map String (Set ExprType) -> (ExprType -> String) -> Array (BackendBindingGroup Ident NeutralExpr) -> Array { recursive :: Boolean, bindings :: Array (Tuple Ident TcoExpr) }
specializeBindingGroups modName pointerAdtPaths modNameStrOrig globalTypes instantiations mangleTypeFn bindingsGroups =
  let
    Tuple _ tcoBindings = Array.foldl
      ( \(Tuple env acc) group ->
          let
            neBindings = fromArray group.bindings
            _ = unsafePerformEffect (Console.log ("Analyzing binding group for monomorphization"))
            env' = case neBindings of
              Just ne | group.recursive -> Tco.topLevelTcoEnvGroup modName ne <> env
              _ -> env
            tcoBinds = map (\(Tuple k v) -> Tuple k (Tco.analyze env' v)) group.bindings
          in
            Tuple env' (Array.snoc acc { recursive: group.recursive, bindings: tcoBinds })
      )
      (Tuple [] [])
      bindingsGroups
  in
  map (\group ->
    group { bindings = Array.concatMap expandBind group.bindings }
  ) tcoBindings
  where
    instsMap = map Set.toUnfoldable instantiations

    expandBind :: Tuple Ident TcoExpr -> Array (Tuple Ident TcoExpr)
    expandBind (Tuple id@(Ident name) val) =
      let qual = modNameStrOrig <> "." <> name
          baseVal = substituteAst instsMap mangleTypeFn val
      in case Map.lookup qual instantiations of
           Just concretes ->
             let genericType = fromMaybe (getExprType val) (Map.lookup qual globalTypes)
                 concreteArr = Set.toUnfoldable concretes :: Array ExprType
                 res = Tuple id baseVal `Array.cons` Array.mapMaybe (\concrete ->
                  let subst = unify genericType concrete Map.empty
                  in if Map.isEmpty subst then Nothing else Just $
                       let mangledName = name <> "__" <> mangleTypeFn concrete
                           mangledVal = mapTcoExprTypes (substituteExprType subst) val
                           mangledVal' = substituteAst instsMap mangleTypeFn mangledVal
                       in Tuple (Ident mangledName) (setTcoExprType concrete mangledVal')
                ) concreteArr
             in res
           Nothing -> [ Tuple id baseVal ]
