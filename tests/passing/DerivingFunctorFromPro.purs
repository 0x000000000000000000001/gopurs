module Main where

-- The function Profunctor instance must be in the import closure when this
-- fixture is compiled on its own; having the package installed is not enough.

import Prelude

import Data.Profunctor (class Profunctor)
import Effect (Effect)
import Effect.Console (log)
import Effect.Ref as Ref
import Test.Assert (assert')

data Test a
  = Test1 ((Array a -> Int) -> Int)
  | Test2 { f :: ({ a :: a } -> Int) -> Int }
derive instance Functor Test

adapt :: Test Int -> Test { value :: Int }
adapt = map \n -> { value: n + 1 }

main :: Effect Unit
main = do
  n <- Ref.new 7 >>= Ref.read
  assert' "map - nested function arguments" $
    case adapt (Test1 \consume -> consume [n, n + 1]) of
      Test1 f -> f (\xs -> if xs == [{ value: 8 }, { value: 9 }] then 42 else 0) == 42
      _ -> false
  assert' "map - nested record function arguments" $
    case adapt (Test2 { f: \consume -> consume { a: n } }) of
      Test2 r -> r.f (\x -> x.a.value * 2) == 16
      _ -> false
  log "Done"
