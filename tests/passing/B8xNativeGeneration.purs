-- @inline Main.sumEleven never
module Main where

-- @dependencies: assert prelude effect console refs unsafe-coerce

import Prelude
import Effect (Effect)
import Effect.Console (log)
import Effect.Ref as Ref
import Test.Assert (assertEqual)
import Unsafe.Coerce (unsafeCoerce)

class Marker a
instance markerInt :: Marker Int

sumEleven :: Int -> Int -> Int -> Int -> Int -> Int -> Int -> Int -> Int -> Int -> Int -> Int
sumEleven a b c d e f g h i j k = a + b + c + d + e + f + g + h + i + j + k

data Packed a b = Packed (a -> b) a
foreign import data Erased :: Type -> Type

pack :: forall a b. Packed a b -> Erased b
pack = unsafeCoerce

unpack :: forall b r. (forall a. Packed a b -> r) -> Erased b -> r
unpack = unsafeCoerce

mapErased :: forall b c. (b -> c) -> Erased b -> Erased c
mapErased f = unpack \(Packed g a) -> pack (Packed (f <<< g) a)

runErased :: forall b. Erased b -> b
runErased = unpack \(Packed f a) -> f a

main :: Effect Unit
main = do
  input <- Ref.new 1
  first <- Ref.read input
  assertEqual { expected: 66, actual: sumEleven first 2 3 4 5 6 7 8 9 10 11 }
  fnRef <- Ref.new sumEleven
  fn <- Ref.read fnRef
  assertEqual { expected: 66, actual: fn first 2 3 4 5 6 7 8 9 10 11 }
  assertEqual { expected: 44, actual: runErased (mapErased (_ * 2) (pack (Packed (_ + 1) 21))) }
  log "Done"
