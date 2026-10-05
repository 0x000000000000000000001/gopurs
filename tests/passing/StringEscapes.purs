module Main where

import Prelude ((==), (/=), (<>), discard, bind, pure, Unit)
import Effect (Effect)
import Effect.Ref as Ref
import Test.Assert (assert, assert')
import Effect.Console (log)

singleCharacter = "\t\n\r\"\\" == "\x9\xA\xD\x22\x5C"
hex = "\x1D306\x2603\x3C6\xE0" == "𝌆☃φà"
surrogatePair = "\xD834\xDF06" == "\x1D306"
highSurrogate = "\xD834"
lowSurrogate = "\xDF06"
loneSurrogates = (highSurrogate <> lowSurrogate) == "\x1D306"
outOfOrderSurrogates = (lowSurrogate <> highSurrogate) == "\xDF06\xD834"
replacement = "\xFFFD"
notReplacing = replacement /= highSurrogate

-- Reading each operand through an effect keeps concatenation live at runtime.
checkConcat :: String -> String -> String -> Effect Unit
checkConcat left right expected = do
  leftRef <- Ref.new left
  rightRef <- Ref.new right
  a <- Ref.read leftRef
  b <- Ref.read rightRef
  assert' "runtime UTF-16 concatenation" ((a <> b) == expected)

main = do
  assert' "single-character escape sequences" singleCharacter
  assert' "hex escape sequences" hex
  assert' "astral code points are represented as a UTF-16 surrogate pair" surrogatePair
  assert' "lone surrogates may be combined into a surrogate pair" loneSurrogates
  assert' "lone surrogates may be combined out of order to remain lone surrogates" outOfOrderSurrogates
  assert' "lone surrogates are not replaced with the Unicode replacement character U+FFFD" notReplacing
  checkConcat highSurrogate lowSurrogate "\x1D306"
  checkConcat ("é" <> highSurrogate) (lowSurrogate <> "z") "é\x1D306z"
  checkConcat lowSurrogate highSurrogate "\xDF06\xD834"
  checkConcat highSurrogate highSurrogate "\xD834\xD834"
  checkConcat "\xD800" "\xDC00" "\x10000"
  checkConcat "\xDBFF" "\xDFFF" "\x10FFFF"
  checkConcat highSurrogate "" highSurrogate
  checkConcat "" lowSurrogate lowSurrogate
  log "Done"
