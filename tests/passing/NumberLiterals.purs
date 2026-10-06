module Main where

-- See issue #2115. Expectations follow the current JS Prelude, preserving the
-- shortest round-tripping digits rather than the old 14-digit fixture oracle.

import Prelude
import Effect (Effect)
import Effect.Ref as Ref
import Test.Assert (assert')
import Effect.Console (log)

main :: Effect Unit
main = do
  test "0.17" 0.17
  test "0.25996181067141905" 0.25996181067141905
  test "0.3572019862807257" 0.3572019862807257
  test "0.46817723004874223" 0.46817723004874223
  test "0.9640035681058178" 0.9640035681058178
  test "4.23808622486133" 4.23808622486133
  test "4.540362294799751" 4.540362294799751
  test "5.212384849884261" 5.212384849884261
  test "13.958257048123212" 13.958257048123212
  test "32.96176575630599" 32.96176575630599
  test "38.47735512322269" 38.47735512322269

  test "10000000000.0" 1e10
  test "10000000000.0" 1.0e10
  test "0.00001" 1e-5
  test "0.00001" 1.0e-5
  test "1.5339794352098402e-118" 1.5339794352098402e-118
  test "2.108934760892056e-59" 2.108934760892056e-59
  test "2.250634744599241e-19" 2.250634744599241e-19
  test "5.960464477539063e-8" 5.960464477539063e-8
  test "5e-324" 5e-324
  test "5e-324" 5.0e-324

  test "0.0" 0.0
  test "0.0" (-0.0)
  test "-5e-324" (-5e-324)
  test "1e-7" 1e-7
  test "-1e-7" (-1e-7)
  test "1e-8" 1e-8
  test "1e-9" 1e-9
  test "0.000001" 1e-6
  test "9.999999999999997e-7" 9.999999999999997e-7
  test "100000000000000000000.0" 1e20
  test "999999999999999900000.0" 9.999999999999999e20
  test "1e+21" 1e21
  test "-1e+21" (-1e21)
  test "9007199254740991.0" 9007199254740991.0
  test "1000000000000000100.0" 1000000000000000128.0
  test "2.2250738585072014e-308" 2.2250738585072014e-308
  test "1.7976931348623157e+308" 1.7976931348623157e308
  test "Infinity" (1.0 / 0.0)
  test "-Infinity" (-1.0 / 0.0)
  test "NaN" (0.0 / 0.0)

  log "Done"

test :: String -> Number -> Effect Unit
test str num = do
  check str num
  ref <- Ref.new num
  dynamic <- Ref.read ref
  check str dynamic

check :: String -> Number -> Effect Unit
check str num =
  if (show num == str)
    then pure unit
    else flip assert' false $
      "Expected " <> show str <> ", got " <> show (show num) <> "."
