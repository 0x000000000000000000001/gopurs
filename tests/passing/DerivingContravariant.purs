module Main where

import Prelude

-- Deriving needs these instance modules in the import closure, even though
-- their classes are only used by the generated implementation.
import Data.Bifunctor (class Bifunctor)
import Data.Profunctor (class Profunctor)
import Data.Functor.Contravariant (class Contravariant, cmap)
import Data.Predicate (Predicate(..))
import Data.Tuple (Tuple(..))
import Effect (Effect)
import Effect.Console (log)
import Effect.Ref as Ref
import Test.Assert (assert')

data Test f a
  = Test0
  | Test1 (Predicate a)
  | Test2 (Predicate (Predicate (Predicate a)))
  | Test3 Int (forall a. Array a -> Array a)
  | Test4 Int (f a)
  | Test5 (Array (a -> Int)) (Tuple (Predicate a) Int)
  | Test6 { nested :: Array { x :: f { a :: a } } }
derive instance Contravariant f => Contravariant (Test f)

adapt :: Test Predicate Int -> Test Predicate { value :: Int }
adapt = cmap _.value

main :: Effect Unit
main = do
  n <- Ref.new 7 >>= Ref.read
  assert' "cmap - empty constructor" $ case adapt Test0 of
    Test0 -> true
    _ -> false
  assert' "cmap - predicate" $ case adapt (Test1 (Predicate (_ == n))) of
    Test1 (Predicate p) -> p { value: n } && not (p { value: 0 })
    _ -> false
  assert' "cmap - triple contravariance" $
    case adapt (Test2 (Predicate \(Predicate p) -> p (Predicate (_ == n)))) of
      Test2 (Predicate p) -> p (Predicate \(Predicate q) -> q { value: n })
      _ -> false
  assert' "cmap - quantified argument stays unchanged" $ case adapt (Test3 42 identity) of
    Test3 tag f -> tag == 42 && f [n] == [n] && f ["x"] == ["x"]
    _ -> false
  assert' "cmap - constrained constructor" $ case adapt (Test4 42 (Predicate (_ == n))) of
    Test4 tag (Predicate p) -> tag == 42 && p { value: n } && not (p { value: 0 })
    _ -> false
  assert' "cmap - function and first tuple argument" $
    case adapt (Test5 [(_ + 2)] (Tuple (Predicate (_ == n)) 42)) of
      Test5 [f] (Tuple (Predicate p) tag) ->
        f { value: n } == 9 && p { value: n } && not (p { value: 0 }) && tag == 42
      _ -> false
  assert' "cmap - nested records" $
    case adapt (Test6 { nested: [{ x: Predicate \r -> r.a == n }] }) of
      Test6 { nested: [{ x: Predicate p }] } ->
        p { a: { value: n } } && not (p { a: { value: 0 } })
      _ -> false
  log "Done"
