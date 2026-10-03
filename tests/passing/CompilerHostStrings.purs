module Main where

-- @dependencies: assert prelude effect console refs

import Prelude

import Effect (Effect)
import Effect.Console (log)
import Effect.Ref as Ref
import Test.Assert (assertEqual)

-- Keep the values live across an effect so this checks literal emission and
-- application execution, rather than constant-folding an equality in PBO.
check :: String -> Effect Unit
check expected = do
  reference <- Ref.new expected
  actual <- Ref.read reference
  assertEqual { expected, actual }

main :: Effect Unit
main = do
  check ""
  check "quotes\" slash\\"
  check "\x0\x8\t\n\r\x1f\x7f"
  check "é漢€💻𝄞\x2028\x2029"
  check "\xD834\xDF06"
  check "\xD800"
  check "\xDFFF"
  check "a\xD800z"
  check "\xD800\xD800\xDC00\xDFFF"
  check "math.Abs(1); unsafe.Pointer(nil); sync.Once{}; gopurs_runtime.Value{}"
  check "/* sync.Once */ // unsafe.Pointer\n\"math.Abs\""
  log "Done"
