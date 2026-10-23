module Gopurs.DependencyGraph
  ( directImports
  , transitiveDependencies
  , transitiveDependents
  ) where

import Prelude

import Data.Array as Array
import Data.Foldable (foldl)
import Data.List (List)
import Data.Map (Map)
import Data.Map as Map
import Data.Maybe (Maybe(..), fromMaybe)
import Data.Newtype (unwrap)
import Data.Set (Set)
import Data.Set as Set
import Data.Tuple (Tuple(..), snd)
import PureScript.Backend.Optimizer.CoreFn (Module(..))
import PureScript.Backend.Optimizer.CoreFn as CoreFn

directImports :: forall a. Module a -> Set String
directImports (Module value) =
  Set.fromFoldable (map (\(CoreFn.Import _ name) -> unwrap name) value.imports)

transitiveDependencies :: forall a. List (Module a) -> Map String (Set String)
transitiveDependencies modules =
  let
    byName = Map.fromFoldable (map (\m@(Module value) -> Tuple (unwrap value.name) m) (Array.fromFoldable modules))
    closure memo visiting name =
      case Map.lookup name memo of
        Just deps -> Tuple memo deps
        Nothing ->
          if Set.member name visiting then Tuple memo Set.empty
          else case Map.lookup name byName of
            Nothing -> Tuple (Map.insert name Set.empty memo) Set.empty
            Just m ->
              let
                Tuple memo' deps = foldl
                  (\(Tuple acc dependencies) dependency ->
                    let
                      Tuple next nested = closure acc (Set.insert name visiting) dependency
                    in
                      Tuple next (Set.insert dependency (Set.union dependencies nested)))
                  (Tuple memo Set.empty)
                  (Array.fromFoldable (directImports m))
              in
                Tuple (Map.insert name deps memo') deps
    build memo result (Module value) =
      let
        name = unwrap value.name
        Tuple next deps = closure memo Set.empty name
      in
        Tuple next (Map.insert name deps result)
    built = foldl
      (\(Tuple memo found) m -> build memo found m)
      (Tuple Map.empty Map.empty)
      (Array.fromFoldable modules)
  in
    snd built

transitiveDependents :: forall a. String -> List (Module a) -> Set String
transitiveDependents changed modules =
  let
    direct = Map.fromFoldableWith Set.union
      ( Array.concatMap
          (\(Module value) ->
            let name = unwrap value.name
            in map (\dependency -> Tuple dependency (Set.singleton name)) (Array.fromFoldable (directImports (Module value))))
          (Array.fromFoldable modules)
      )
    go pending seen = case Array.uncons (Set.toUnfoldable pending :: Array String) of
      Nothing -> seen
      Just { head: name, tail } ->
        let rest = Set.fromFoldable tail
        in if Set.member name seen then go rest seen
           else
             let next = fromMaybe Set.empty (Map.lookup name direct)
             in go (Set.union next rest) (Set.insert name seen)
  in go (fromMaybe Set.empty (Map.lookup changed direct)) Set.empty
