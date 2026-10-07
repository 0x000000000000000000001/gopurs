module Main where

import Prelude

-- The corresponding tuple instances live in these modules, not Data.Tuple.
import Data.Bifoldable (class Bifoldable)
import Data.Bifunctor (class Bifunctor)
import Data.Bitraversable (class Bitraversable)
import Data.Foldable (class Foldable, foldl, foldr, foldMap)
import Data.Traversable (class Traversable, traverse, sequence)
import Data.Tuple (Tuple(..))
import Effect (Effect)
import Effect.Console (log)
import Effect.Ref as Ref
import Test.Assert (assert')

data Test a
  = Test1 (Tuple a Int)
  | Test2 (Tuple (Array a) a)
  | Test3 { x :: Tuple { a :: a } Int, y :: Tuple { a :: Array a } { a :: a } }
derive instance Functor Test
derive instance Foldable Test
derive instance Traversable Test
derive instance Eq a => Eq (Test a)

check :: String -> Test Int -> Array Int -> Test String -> Effect Unit
check name value expected mapped = do
  assert' (name <> " - map") $ map show value == mapped
  assert' (name <> " - foldl") $ foldl (\xs x -> xs <> [x]) [] value == expected
  assert' (name <> " - foldr") $ foldr (\x xs -> [x] <> xs) [] value == expected
  assert' (name <> " - foldMap") $ foldMap (\x -> [x]) value == expected
  seen <- Ref.new []
  let visit x = Ref.modify_ (_ <> [x]) seen *> pure (show x)
  traversed <- traverse visit value
  assert' (name <> " - traverse") $ traversed == mapped
  visited <- Ref.read seen
  assert' (name <> " - traverse order") $ visited == expected
  Ref.write [] seen
  sequenced <- sequence (map visit value)
  assert' (name <> " - sequence") $ sequenced == mapped
  sequencedOrder <- Ref.read seen
  assert' (name <> " - sequence order") $ sequencedOrder == expected

main :: Effect Unit
main = do
  n <- Ref.new 7 >>= Ref.read
  check "left argument" (Test1 (Tuple n 42)) [7] (Test1 (Tuple "7" 42))
  check "both arguments" (Test2 (Tuple [n, n + 1] (n + 2))) [7, 8, 9]
    (Test2 (Tuple ["7", "8"] "9"))
  check "nested records"
    (Test3 { x: Tuple { a: n } 42, y: Tuple { a: [n + 1, n + 2] } { a: n + 3 } })
    [7, 8, 9, 10]
    (Test3 { x: Tuple { a: "7" } 42, y: Tuple { a: ["8", "9"] } { a: "10" } })
  log "Done"
