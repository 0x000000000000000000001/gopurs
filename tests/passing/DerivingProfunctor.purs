module Main where

import Prelude

-- Test4 maps the first tuple argument through this instance module.
import Data.Bifunctor (class Bifunctor)
import Data.Predicate (Predicate(..))
import Data.Profunctor (class Profunctor, dimap)
import Data.Tuple (Tuple(..))
import Effect (Effect)
import Effect.Console (log)
import Effect.Ref as Ref
import Test.Assert (assert')

data Test f a b
  = Test0
  | Test1 (Predicate a) b
  | Test2 Int (forall a. Array a -> Array a)
  | Test3 Int (f a b) (f a Int) (f Int b)
  | Test4 (Array (a -> Int)) (Tuple b Int)
  | Test5 { nested :: Array { x :: f { a :: a } { b :: b } } }
derive instance Profunctor f => Profunctor (Test f)

adapt :: Test (->) Int Int -> Test (->) { value :: Int } String
adapt = dimap _.value show

main :: Effect Unit
main = do
  n <- Ref.new 7 >>= Ref.read
  assert' "dimap - empty constructor" $ case adapt Test0 of
    Test0 -> true
    _ -> false
  assert' "dimap - predicate and result" $ case adapt (Test1 (Predicate (_ == n)) (n + 1)) of
    Test1 (Predicate p) result -> p { value: n } && not (p { value: 0 }) && result == "8"
    _ -> false
  assert' "dimap - quantified argument stays unchanged" $ case adapt (Test2 42 identity) of
    Test2 tag f -> tag == 42 && f [n] == [n] && f ["x"] == ["x"]
    _ -> false
  assert' "dimap - both, left-only and right-only arguments" $
    case adapt (Test3 42 (_ + 1) (_ + 2) (_ + 3)) of
      Test3 tag both left right ->
        tag == 42 && both { value: n } == "8" && left { value: n } == 9 && right 4 == "7"
      _ -> false
  assert' "dimap - function and first tuple argument" $
    case adapt (Test4 [(_ + 1)] (Tuple n 42)) of
      Test4 [f] (Tuple result tag) -> f { value: n } == 8 && result == "7" && tag == 42
      _ -> false
  assert' "dimap - nested records" $
    case adapt (Test5 { nested: [{ x: \r -> { b: r.a + 1 } }] }) of
      Test5 { nested: [{ x: f }] } -> f { a: { value: n } } == { b: "8" }
      _ -> false
  log "Done"
