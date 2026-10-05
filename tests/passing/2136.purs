module Main where

import Prelude
import Data.Int.Bits as Bits
import Effect (Effect)
import Effect.Console (log)
import Effect.Ref as Ref
import Test.Assert (assert')

-- Literal-derived operands exercise the optimizer as well as the runtime checks.
literalBottom :: Int
literalBottom = (-2147483647) - 1
foldedNegate :: Int
foldedNegate = negate literalBottom
foldedAdd :: Int
foldedAdd = 2147483647 + 1
foldedSubtract :: Int
foldedSubtract = literalBottom - 1
foldedMultiply :: Int
foldedMultiply = 2147483647 * 2147483647
foldedShift :: Int
foldedShift = Bits.shl 1 32

check1 :: String -> (Int -> Int) -> Int -> Int -> Effect Unit
check1 label op value expected = do
  ref <- Ref.new value
  a <- Ref.read ref
  assert' label (op a == expected)

check2 :: String -> (Int -> Int -> Int) -> Int -> Int -> Int -> Effect Unit
check2 label op left right expected = do
  leftRef <- Ref.new left
  rightRef <- Ref.new right
  a <- Ref.read leftRef
  b <- Ref.read rightRef
  assert' label (op a b == expected)

main = do
  assert' "2136: negating bottom does not exceed top" (not (negate (bottom :: Int) > top))
  assert' "folded negation wraps" (foldedNegate == bottom)
  assert' "folded addition wraps" (foldedAdd == bottom)
  assert' "folded subtraction wraps" (foldedSubtract == top)
  assert' "multiplication follows the JS Number then Int conversion" (foldedMultiply == 0)
  assert' "folded shift masks the count" (foldedShift == 1)
  check1 "runtime negation of bottom" negate bottom bottom
  check1 "runtime negation of top" negate top (-2147483647)
  check1 "runtime complement" Bits.complement bottom top
  check2 "runtime addition" (+) top 1 bottom
  check2 "runtime subtraction" (-) bottom 1 top
  check2 "runtime multiplication" (*) top 2 (-2)
  check2 "runtime multiplication with Number rounding" (*) top top 0
  check2 "runtime shift overflow" Bits.shl 1 31 bottom
  check2 "runtime shift count 32" Bits.shl 1 32 1
  check2 "runtime negative shift count" Bits.shl 1 (-1) bottom
  check2 "runtime arithmetic shift" Bits.shr bottom 32 bottom
  check2 "runtime unsigned shift" Bits.zshr bottom 1 1073741824
  check2 "runtime unsigned shift with count 32" Bits.zshr bottom 32 (Bits.zshr bottom 0)
  check2 "runtime Euclidean division boundary" div bottom (-1) (Bits.zshr bottom 0)
  check2 "runtime modulo boundary" mod bottom top 2147483646
  log "Done"
