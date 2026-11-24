module Main where

import Prelude

import Data.Array as Array
import Effect (Effect)
import Effect.Console (logShow)
import Effect.Ref as Ref

data Grove = Seed Int | Join Grove Grove | Weighted Int Grove

measure :: Grove -> Int
measure (Seed value) = value
measure (Join left right) = measure left + measure right
measure (Weighted weight value) = weight * measure value

difference :: Grove -> Int
difference (Seed value) = value
difference (Join left right) = difference left - difference right
difference (Weighted weight value) = weight + difference value

plant :: Int -> Int -> Grove
plant depth seed
  | depth <= 0 = Seed seed
  | depth == 1 = Weighted seed (Seed 3)
  | otherwise = Join (Weighted depth (plant (depth - 1) seed)) (Seed (depth + seed))

observe :: Int -> Int -> Int
observe depth seed = measure (plant depth seed)

observeDifference :: Int -> Int -> Int
observeDifference depth seed = difference (plant depth seed)

-- This caller retains and consumes the same immutable tree twice.
shared :: Grove -> Int
shared tree = measure tree + difference tree

-- Discarding a subtree is deliberately outside the fusion contract.
leftmost :: Grove -> Int
leftmost (Seed value) = value
leftmost (Join left _) = leftmost left
leftmost (Weighted _ value) = leftmost value

-- Opaque callbacks remain ordinary calls, including in constructor producers.
mapMeasure :: (Int -> Int) -> Grove -> Int
mapMeasure callback (Seed value) = callback value
mapMeasure callback (Join left right) = mapMeasure callback left + mapMeasure callback right
mapMeasure callback (Weighted weight value) = callback weight * mapMeasure callback value

plantWith :: (Int -> Int) -> Int -> Int -> Grove
plantWith callback depth seed
  | depth <= 0 = Seed (callback seed)
  | otherwise = Join (Seed (callback depth)) (plantWith callback (depth - 1) seed)

main :: Effect Unit
main = do
  traces <- Ref.new []
  Ref.modify_ (_ <> [ 1 ]) traces
  seedRef <- Ref.new 7
  seed <- Ref.read seedRef
  Ref.modify_ (_ <> [ 2 ]) traces
  depthRef <- Ref.new 8
  depth <- Ref.read depthRef
  Ref.modify_ (_ <> [ 3 ]) traces
  logShow (map (\n -> observe n seed) (Array.range (-2) depth))
  logShow (map (\n -> observeDifference n seed) (Array.range (-2) depth))
  logShow (map (observe depth) [ 0, -1, 2147483647, -2147483648 ])
  logShow (shared (plant depth seed))
  logShow (leftmost (plant depth seed))
  callbackRef <- Ref.new (\x -> x + 11)
  callback <- Ref.read callbackRef
  logShow (mapMeasure callback (plant 3 seed))
  logShow (measure (plantWith callback 4 seed))
  Ref.modify_ (_ <> [ 4 ]) traces
  logShow =<< Ref.read traces
